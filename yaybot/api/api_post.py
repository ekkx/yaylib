from ..config import Endpoints as ep
from ..utils import console_print


def create_post(self, post_type, text, image, choices=None, color=0, font_size=0):
    valid_types = ['text', 'survey', 'image']

    if post_type not in valid_types:
        raise ValueError(f'Invalid post type. Must be one of {valid_types}')

    data = {
        'text': text,
        'color': color,
        'font_size': font_size,
        'post_type': post_type,
        'uuid': self.UUID,
        'message_tags': '[]',
    }

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

    resp = self._post(url, data)
    if resp['result'] == 'success':
        self.logger.info(
            'Post Created [https://yay.space/post/{}]'.format(resp['id']))
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
