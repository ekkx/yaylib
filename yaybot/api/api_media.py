import datetime
import requests
import os

from PIL import Image
from ..config import Endpoints as ep


def upload_photo(self, photo):
    filename, ext = os.path.splitext(photo)
    image = Image.open(photo)

    width, height = image.size

    date = datetime.datetime.now()
    timestamp = int(date.timestamp() * 1000)

    post_url = f'post/{date.year}/{date.month}/{date.day}/{filename}_{timestamp}_0_size_{width}x{height}{ext}'
    post_url_thumb = f'post/{date.year}/{date.month}/{date.day}/thumb_{filename}_{timestamp}_0_size_{width}x{height}{ext}'

    data = self._get(
        f'{ep.API_URL}/v1/buckets/presigned_urls?file_names[]={post_url}&file_names[]={post_url_thumb}'
    )

    filename_original = data['presigned_urls'][0]['filename']
    filename_thumb = data['presigned_urls'][1]['filename']

    # upload original
    url = data['presigned_urls'][0]['url']
    with open(photo, "rb") as f:
        resp = requests.put(url, data=f.read())

    # upload thumbnail
    url = data['presigned_urls'][1]['url']
    with open(photo, "rb") as f:
        resp = requests.put(url, data=f.read())

    data = {
        'filename': filename_original,
        'filename_thumb': filename_thumb,
        'width': width,
        'height': height
    }

    return data
