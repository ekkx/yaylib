import datetime
import requests
import os

from PIL import Image
from requests import HTTPError
from ..config import Endpoints as ep
from ..utils import console_print


def upload_image(self, image_type, image):
    valid_types = ['post', 'user_avatar']

    if image_type not in valid_types:
        raise ValueError(f'Invalid post type. Must be one of {valid_types}')

    filename, ext = os.path.splitext(image)
    image_data = Image.open(image)

    width, height = image_data.size

    date = datetime.datetime.now()
    timestamp = int(date.timestamp() * 1000)

    url_main = f'{image_type}/{date.year}/{date.month}/{date.day}/{filename}_{timestamp}_0_size_{width}x{height}{ext}'
    url_thumb = f'{image_type}/{date.year}/{date.month}/{date.day}/thumb_{filename}_{timestamp}_0_size_{width}x{height}{ext}'

    data = self._get(
        f'{ep.API_URL}/v1/buckets/presigned_urls?file_names[]={url_main}&file_names[]={url_thumb}'
    )

    filename_original = data['presigned_urls'][0]['filename']
    filename_thumb = data['presigned_urls'][1]['filename']

    # upload original
    url = data['presigned_urls'][0]['url']
    with open(image, "rb") as f:
        resp = requests.put(url, data=f.read())

    # upload thumbnail
    url = data['presigned_urls'][1]['url']
    with open(image, "rb") as f:
        resp = requests.put(url, data=f.read())

    data = {
        'filename': filename_original,
        'filename_thumb': filename_thumb,
        'width': width,
        'height': height
    }

    return data


def download_image(self, media_url, filename=None, folder='images'):
    resp = requests.get(media_url)

    if resp.status_code != 200:
        raise HTTPError(f'The URL is invalid')

    image_data = resp.content

    basename = os.path.basename(media_url)
    filename = filename or os.path.splitext(basename)[0]
    ext = os.path.splitext(basename)[1]
    fname = os.path.join(folder, filename + ext)

    if not os.path.exists(folder):
        os.makedirs(folder)

    if os.path.exists(fname):
        self.logger.info('File already exists, skipping...')
        return fname

    with open(fname, 'wb') as f:
        f.write(image_data)
    return fname


def download_video(self, path: str = None):
    return
