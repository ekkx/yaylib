import hmac
import hashlib
import base64
import uuid

from datetime import datetime


def generate_uuid() -> tuple:
    generated_uuid = str(uuid.uuid4())
    url_uuid = generated_uuid.replace("-", "")
    return generated_uuid, url_uuid


def parse_datetime(timestamp: int) -> str:
    if timestamp is not None:
        return str(datetime.fromtimestamp(timestamp))
    return timestamp


def signed_info_calculating(api_key: str, device_uuid: str, timestamp: int) -> str:
    return str(hashlib.md5(str(api_key + device_uuid + str(timestamp)).encode()).hexdigest())


def signed_version_calculating(yay_version_message: str, api_version_key: str) -> str:
    return base64.b64encode(hmac.new(api_version_key.encode("utf-8"), yay_version_message.encode("utf-8"), hashlib.sha256).digest()).decode("utf-8")
