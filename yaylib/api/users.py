from typing import Dict

from ..config import *
from ..errors import *
from .api import (
    _check_authorization,
    _get,
    _post,
    _put,
    _delete,
    _handle_response,
)


def contact_friends(self, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.USER_V1}/contact_friends",
        headers=headers
    )
