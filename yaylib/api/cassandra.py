from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def get_user_activities(self, **params) -> ActivitiesResponse:
    """

    Parameters:
    ---------------
        - important: bool - (required)
        - from_timestamp: int - (optional)
        - number: int - (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{self.host}/api/user_activities",
        params=params, data_type=ActivitiesResponse
    )


def get_user_merged_activities(self, from_timestamp: int = None) -> ActivitiesResponse:
    self._check_authorization()
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    return self._make_request(
        "GET", endpoint=f"{self.host}/api/v2/user_activities",
        params=params, data_type=ActivitiesResponse
    )


def received_notification(self, pid: str, type: str, opened_at: int = None):
    # TODO: opened_atはnullalbeか確認する
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{self.host}/api/received_push_notifications",
        payload={
            "pid": pid,
            "type": type,
            "opened_at": opened_at
        }
    )
