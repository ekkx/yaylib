import hmac
import hashlib
import base64
from datetime import datetime


def parse_datetime(timestamp: int) -> str:
    return str(datetime.fromtimestamp(timestamp))


def signed_info_calculating(api_key: str, device_uuid: str, timestamp: int) -> str:
    return str(hashlib.md5(str(api_key + device_uuid + str(timestamp)).encode()).hexdigest())


def signed_version_calculating(yay_vision_message: str, api_vision_key: str) -> str:
    return base64.b64encode(hmac.new(api_vision_key.encode("utf-8"), yay_vision_message.encode("utf-8"), hashlib.sha256).digest()).decode("utf-8")
