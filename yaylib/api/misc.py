from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..utils import *


def accept_policy_agreement(self, type: str):
    pass


def generate_sns_thumbnail(
        self,
        resource_type: str,
        resource_id: int
):
    pass


def get_email_verification_presigned_url(
        self,
        device_uuid: str,
        email: str,
        locale: str,
        intent: str = None
) -> str:
    # EmailVerificationPresignedUrlResponse.url
    pass


def get_web_socket_token(self) -> str:
    self._check_authorization()
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V1}/ws_token",
        data_type=WebSocketTokenResponse
    )
    return response.token
