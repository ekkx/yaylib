import hmac
import hashlib
import base64
import uuid

from datetime import datetime
from .config import Configs


class colors:

    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    RESET = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'


def generate_uuid() -> tuple:
    generated_uuid = str(uuid.uuid4())
    url_uuid = generated_uuid.replace("-", "")
    return generated_uuid, url_uuid


def parse_datetime(timestamp: int) -> str:
    if timestamp is not None:
        return str(datetime.fromtimestamp(timestamp))
    return timestamp


def signed_info_calculating(uuid: str, timestamp: int, shared_key: bool = False) -> str:
    """
    Pass the device_uuid when shared_key is False.
    """
    shared_key = Configs.YAY_SHARED_KEY if shared_key is True else ""
    return hashlib.md5((
        Configs.YAY_API_KEY + uuid + str(timestamp) + shared_key
    ).encode()).hexdigest()


def signed_version_calculating() -> str:
    hash_object = hmac.new(
        Configs.YAY_API_VERSION_KEY.encode(),
        "yay_android/{}".format(Configs.YAY_API_VERSION).encode(),
        hashlib.sha256
    )
    return base64.b64encode(hash_object.digest()).decode("utf-8")
