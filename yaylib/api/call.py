from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..utils import *


def bump_call(self, call_id: int, participant_limit: int = None):
    pass


def get_active_call(self, user_id: int) -> Post:
    pass


def get_bgms(self) -> List[Bgm]:
    pass


def get_call(self, call_id: int) -> ConferenceCall:
    pass


def get_call_invitable_users(
        self, call_id: int,
        from_timestamp: int = None,
        # @Nullable @Query("user[nickname]")
) -> UsersByTimestampResponse:
    pass


def get_call_status(self, opponent_id: int) -> CallStatusResponse:
    pass


def get_games(
        self,
        number: int,
        ids: List[int],
        from_id: int = None
) -> GamesResponse:
    # v1/games/apps?ids[]=11&ids[]=22
    pass


def get_genres(self, number: int, from_id: int = None) -> GenresResponse:
    pass


def get_group_calls(
        self,
        number: int = None,
        group_category_id: int = None,
        from_timestamp: int = None,
        scope: str = None
) -> PostsResponse:
    pass


def invite_to_call_bulk(self, call_id: int, group_id: int = None):
    pass


def invite_users_to_call(self, call_id: int, user_ids: List[int]):
    pass


def invite_users_to_chat_call(
        self,
        chat_room_id: int,
        room_id: int,
        room_url: str
):
    pass


def kick_and_ban_from_call(self, call_id: int, user_id: int):
    pass


def notify_anonymous_user_leave_agora_channel(
        self,
        conference_id: int,
        agora_uid: str
):
    pass


def notify_user_leave_agora_channel(
        self,
        conference_id: int,
        user_id: int
):
    pass


def send_call_screenshot(
        self,
        screenshot_filename: str,
        conference_id: int
):
    pass


def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: str = None,
        category_id: str = None
):
    pass


def set_user_role(
        self,
        call_id: int,
        user_id: int,
        role: str
):
    pass


def start_call(
        self,
        conference_id: int,
        call_sid: str
):
    pass


def stop_call(
        self,
        conference_id: int,
        call_sid: str
):
    pass
