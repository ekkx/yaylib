from huepy import red

from ..config import Endpoints as ep
from ..utils import console_print


def create_post(self, post_type, text, image, choices=None, color=0, font_size=0):
    valid_types = ['text', 'survey', 'image']

    if post_type not in valid_types:
        raise ValueError(f'Invalid post type. Must be one of {valid_types}')

    data = {
        'post_type': post_type,
        'text': text,
        'color': color,
        'font_size': font_size,
        'message_tags': '[]',
        'uuid': self.UUID,
    }

    if '@:start:' in text and ':end:' in text:
        mention_format_info = parse_mention_format(self, text)
        data['text'] = mention_format_info['text']
        data['message_tags'] = mention_format_info['message_tags']

    if post_type == 'text':
        url = '{}/v1/web/posts/new'.format(ep.API_URL)

    elif post_type == 'survey':
        if choices is None:
            raise ValueError('Choices cannot be None for survey post type')

        data['choices[]'] = choices
        url = 'https://yay.space/api/posts'

    elif post_type == 'image':
        if image is None:
            raise ValueError('Image cannot be None for image post type')

        image_data = self.upload_image('post', image)
        data['attachment_sizes'] = {
            'attachment': [image_data['width'], image_data['height']]
        }
        data['attachment_filename'] = image_data['filename']
        url = 'https://yay.space/api/posts'

    print(data)

    resp = self._post(url, data)
    if resp['result'] == 'success':
        self.logger.info(
            'Post Created [https://yay.space/post/{}]'.format(resp['id']))
    else:
        self.logger.error(
            red(f"Failed to create a post [{resp['message']}]"))
    return resp


def create_post_in_group(self, group_id, post_type, text, image, choices=None, color=0, font_size=0):
    valid_types = ['text', 'survey', 'image']

    if post_type not in valid_types:
        raise ValueError(f'Invalid post type. Must be one of {valid_types}')

    data = {
        'text': text,
        'color': color,
        'font_size': font_size,
        'post_type': post_type,
        'group_id': group_id,
        'uuid': self.UUID,
        'message_tags': '[]',
    }
    url = 'https://yay.space/api/posts'

    if post_type == 'survey':
        if choices is None:
            raise ValueError('Choices cannot be None for survey post type')
        data['choices[]'] = choices

    elif post_type == 'image':
        if image is None:
            raise ValueError('Image cannot be None for image post type')

        image_data = self.upload_image(image)
        data['attachment_sizes'] = {
            'attachment': [image_data['width'], image_data['height']]
        }
        data['attachment_filename'] = image_data['filename']

    resp = self._post(url, data)
    return resp


def create_repost(self, text, post_id, color=0, font_size=0):
    data = {
        'text': text,
        'color': color,
        'font_size': font_size,
        'post_id': post_id,
        'message_tags': '[]',
        'post_type': 'text',
        'uuid': ''
    }
    resp = self._post(
        f'{ep.POST_v3}/repost', data)
    return resp


def create_reply(self, text, post_id, color=0, font_size=0):
    data = {
        'text': text,
        'color': color,
        'font_size': font_size,
        'in_reply_to': post_id
    }
    resp = self._post(
        f'{ep.API_URL}/v1/web/posts/new', data)
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


def like_posts(self, post_ids: list):
    data = {'post_ids[]': post_ids}
    resp = self._post(
        f'{ep.POST_v2}/like', data)
    return resp


def unlike_post(self, post_id):
    resp = self._post(
        f'{ep.POST_v1}/{post_id}/unlike')
    return resp


def mention(self, user_id):
    if not isinstance(user_id, int):
        return user_id
    return '@:start:' + str(user_id) + ':end:'


def convert_text_format(self, text):
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

    data = convert_text_format(self, text)
    text = data['text']
    user_ids = data['user_ids']

    for user_id in user_ids:
        start = text.find("@", offset)
        if start == -1:
            break
        end = text.find(" ", start)
        if end == -1:
            end = len(text)
        username = text[start+1:end]
        message_tags.append({"type": "user", "user_id": int(user_id), "offset": start, "length": len(username)+1})
        offset = end

    return {'text': text, 'message_tags': str(message_tags)}