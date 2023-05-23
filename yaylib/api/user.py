from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..utils import *


def contact_friends(self):
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V1}/contact_friends",
    )


def get_user(self, user_id: int):
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V2}/{user_id}",
        data_type=User
    )


def follow(self, user_id: int):
    return self._make_request("GET", endpoint=f"{Endpoints.USER_V2}/{str(user_id)}/follow", data_type=User)


# BlockApi, HiddenApi
