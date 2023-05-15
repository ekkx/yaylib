import json
from huepy import red

from ..config import Endpoints as ep
from ..utils import console_print


def create_post(self, text, image, choices=None, color=0, font_size=0, reply_to=None, group_id=None):
    post_type = 'survey' if choices else 'image' if image else 'text'

    data = {
        'post_type': post_type,
        'text': text,
        'color': color,
        'font_size': font_size,
        'message_tags': '[]',
        'uuid': self.UUID,
    }

    if reply_to:
        data['in_reply_to'] = reply_to

    if group_id:
        data['group_id'] = group_id

    if post_type == 'text':
        url = '{}/v1/web/posts/new'.format(ep.API_URL)
    else:
        url = 'https://yay.space/api/posts'

    if choices:
        data['choices[]'] = choices

    if image:
        image_data = self.upload_image('post', image)
        data['attachment_sizes'] = {
            'attachment': [image_data['width'], image_data['height']]
        }
        data['attachment_filename'] = image_data['filename']

    if '@:start:' in text and ':end:' in text:
        mention_format_info = parse_mention_format(self, text)
        data['text'] = mention_format_info['text']
        data['message_tags'] = json.dumps(mention_format_info['message_tags'])
        url = 'https://yay.space/api/posts'

    resp = self._post(url, data)
    self.logger.debug(resp)

    if resp['result'] == 'success':
        self.logger.info(
            'post created (https://yay.space/post/{})'.format(resp['id']))
        return resp
    self.logger.error(
        red('failed to create a post ({})'.format(resp['message'])))
    return resp


def create_repost(self, post_id, text, image, choices=None, color=0, font_size=0):
    # TODO: fix (attachment_sizes is invalid) error
    post_type = 'survey' if choices else 'image' if image else 'text'

    data = {
        'post_type': post_type,
        'text': text,
        'color': color,
        'font_size': font_size,
        'post_id': post_id,
        'message_tags': '[]',
        'uuid': self.UUID,
    }

    url = f'{ep.POST_v3}/repost'

    if choices:
        data['choices[]'] = choices

    if image:
        image_data = self.upload_image('post', image)
        data['attachment_sizes'] = {
            'attachment': [image_data['width'], image_data['height']]
        }
        data['attachment_filename'] = image_data['filename']

    if '@:start:' in text and ':end:' in text:
        mention_format_info = parse_mention_format(self, text)
        data['text'] = mention_format_info['text']
        data['message_tags'] = json.dumps(mention_format_info['message_tags'])

    resp = self._post(url, data)
    self.logger.debug(resp)

    if resp['result'] == 'success':
        self.logger.info(
            'repost created (https://yay.space/post/{})'.format(resp['post']['id']))
        return resp
    self.logger.error(
        red('failed to recreate a post ({})'.format(resp['message'])))
    return resp


def delete_post(self, post_id):
    data = {'posts_ids[]': post_id}
    resp = self._post(
        f'{ep.POST_v2}/mass_destroy', data)
    return resp


def pin_post(self, post_id):
    data = {'id': post_id}
    resp = self._post(
        f'{ep.PIN_v1}/posts', data)
    return resp


def unpin_post(self, post_id):
    resp = self._post(
        f'{ep.PIN_v1}/posts/{post_id}')
    return resp


def pin_group_post(self, group_id, post_id):
    data = {'group_id': group_id, 'post_id': post_id}
    resp = self._put(
        f'{ep.POST_v2}/group_pinned_post', data)
    return resp


def unpin_group_post(self, group_id):
    data = {'group_id': group_id}
    resp = self._delete(
        f'{ep.POST_v2}/group_pinned_post', data)
    return resp


def like_posts(self, post_ids: list):
    data = {'post_ids[]': post_ids}
    resp = self._post(
        f'{ep.POST_v2}/like', data)
    return resp


def unlike_post(self, post_id):
    resp = self._post(
        f'{ep.POST_v1}/{post_id}/unlike')
    return resp


def vote_post(self, post_id, vote_to_word, choice_id=None):
    post = self.get_post(post_id)
    survey_id = post.survey_id
    if survey_id is None:
        raise ValueError('the post is not a survey post')

    if choice_id is None:
        choices = post.survey_choices
        if not vote_to_word in choices:
            raise ValueError('the vote_to_word is not in the choices')
        choice_ids = post.survey_choice_ids
        choice_id = choice_ids[choices.index(vote_to_word)]

    data = {
        'post_id': post_id,
        'survey_id': survey_id,
        'choice_id': choice_id,
    }
    resp = self._post(
        f'{ep.SURVEY_v2}/{survey_id}/vote', data
    )
    return resp


def mention(self, user_id):
    if not isinstance(user_id, int):
        return user_id
    return '@:start:' + str(user_id) + ':end:'


def convert_mention_text_format(self, text):
    formatted_text = ''
    user_ids = []
    segments = text.split('@:start:')

    for i, segment in enumerate(segments):
        if i == 0:
            formatted_text += segment
            continue
        user_id, text = segment.split(':end:')
        username = self.get_user(user_id).username
        formatted_text += '@' + username + ' ' + text
        user_ids.append(user_id)

    return {'text': formatted_text, 'user_ids': user_ids}


def parse_mention_format(self, text):
    user_ids = {}
    message_tags = []
    offset = 0

    data = convert_mention_text_format(self, text)
    text = data['text']
    user_ids = data['user_ids']

    for user_id in user_ids:
        start = text.find('@', offset)
        if start == -1:
            break
        end = text.find(' ', start)
        if end == -1:
            end = len(text)
        username = text[start+1:end]
        message_tags.append({'type': 'user', 'user_id': int(
            user_id), 'offset': start, 'length': len(username)+1})
        offset = end

    return {'text': text, 'message_tags': message_tags}
