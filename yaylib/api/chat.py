from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def accept_request(self, chat_room_ids: List[int]):
    pass


def check_unread_status(self, from_time: int) -> bool:
    # this.isUnread = bool;
    pass


def create_group_chat(
        self,
        name: str,
        with_user_ids: List[int],
        icon_filename: str = None,
        background_filename: str = None
) -> int:
    # this.roomId = j2;
    pass


def create_private_chat(
        self,
        with_user_id: int,
        matching_id: int = None,
        hima_chat: bool = False
) -> int:
    # this.roomId = j2;
    pass


def delete_background(self, chat_room_id: int):
    pass


def delete_message(self, room_id: int, message_id: int):
    pass


def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: str = None,
        background_filename: str = None
):
    pass


def get_chatable_users(
        self,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None
) -> FollowUsersResponse:
    pass


def get_gifs_data(self) -> List[GifImageCategory]:
    pass


def get_hidden_chat_rooms(
        self,
        number: int,
        from_timestamp: int = None
) -> ChatRoomsResponse:
    pass


def get_main_chat_rooms(self, from_timestamp: int = None) -> ChatRoomsResponse:
    pass


def get_messages(
        self,
        chat_room_id: int,
        from_message_id: int = None,
        to_message_id: int = None
) -> List[Message]:
    pass


def get_notification_settings(self, chat_room_id: int) -> Settings:
    pass


def get_request_chat_rooms(self, from_timestamp: int = None) -> ChatRoomsResponse:
    pass


def get_chat_room(self, chat_room_id: int) -> ChatRoom:
    pass


def get_sticker_packs(self) -> List[StickerPack]:
    pass


def get_total_chat_requests(self) -> int:
    # this.total = i2;
    pass


def hide_chat(self, chat_room_id: int):
    pass


def invite_to_chat(self, chat_room_id: int, with_user_ids: List[int]):
    pass


def kick_users_from_chat(self, chat_room_id: int, user_ids: List[int]):
    pass


def pin_chat(self, chat_room_id: int):
    pass


def read_attachment(
        self,
        chat_room_id: int,
        attachment_msg_ids: List[int]
        # ↑ @Body @NotNull ReadAttachmentRequest readAttachmentRequest
):
    pass


def read_message(self, chat_room_id: int, message_id: int):
    pass


def read_video_message(
        self,
        chat_room_id: int,
        video_msg_ids: List[int]
):
    pass


def refresh_chat_rooms(self, from_time: int = None) -> ChatRoomsResponse:
    pass


def remove_chat_rooms(self, chat_room_ids: List[int]):
    pass


def report_chat_room(
        self,
        chat_room_id: int,
        opponent_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None
):
    pass


def send_media_screenshot_notification(self, chat_room_id: int):
    pass


def send_message(
        self,
        chat_room_id: int,
        message_type: str,
        call_type: str = None,
        text: str = None,
        font_size: int = None,
        gif_image_id: int = None,
        attachment_file_name: str = None,
        sticker_pack_id: int = None,
        video_file_name: str = None
) -> MessageResponse:
    pass


def set_notification_settings(
        chat_room_id: int,
        notification_chat: int
) -> Settings:
    # NotificationSettingResponse
    pass


def un_hide_chat(self, chat_room_id: int):
    pass


def unpin_chat(self, chat_room_id: int):
    pass
