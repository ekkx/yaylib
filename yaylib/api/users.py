from typing import Dict

from ..config import (Endpoints, Configs)
from ..errors import (
    YayError,
    ForbiddenError,
    RateLimitError,
    AuthenticationError,
)
from .api import (
    _get,
    _post,
    _put,
    _delete,
    _check_authorization,
    _handle_response,
)


def contact_friends(self, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.USER_V1}/contact_friends",
        headers=headers
    )
