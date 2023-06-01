from datetime import datetime
from typing import Union, Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def accept_request(self, chat_room_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/accept_chat_request",
        payload={"chat_room_ids[]": chat_room_ids}
    )


def check_unread_status(self, from_time: int) -> UnreadStatusResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/accept_chat_request",
        params={"from_time": from_time},
        data_type=UnreadStatusResponse
    )


def create_group_chat(
        self,
        name: str,
        with_user_ids: List[int],
        icon_filename: str = None,
        background_filename: str = None
) -> CreateChatRoomResponse:
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V3}/new",
        payload={
            "name": name,
            "with_user_ids[]": with_user_ids,
            "icon_filename": icon_filename,
            "background_filename": background_filename,
        },
        data_type=CreateChatRoomResponse
    )


def create_private_chat(
        self,
        with_user_id: int,
        matching_id: int = None,
        hima_chat: bool = False
) -> CreateChatRoomResponse:
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/new",
        payload={
            "with_user_id": with_user_id,
            "matching_id": matching_id,
            "hima_chat": hima_chat,
        },
        data_type=CreateChatRoomResponse
    )


def delete_background(self, room_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{room_id}/background",
    )


def delete_message(self, room_id: int, message_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{room_id}/messages/{message_id}/delete",
    )


def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: str = None,
        background_filename: str = None
):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{chat_room_id}/edit",
        payload={
            "name": name,
            "icon_filename": icon_filename,
            "background_filename": background_filename,
        }
    )


def get_chatable_users(
        self,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None
) -> FollowUsersResponse:
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/followings/chatable",
        payload={
            "from_follow_id": from_follow_id,
            "from_timestamp": from_timestamp,
            "order_by": order_by,
        }, data_type=FollowUsersResponse
    )


def get_gifs_data(self) -> List[GifImageCategory]:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.HIDDEN_V1}/chats", data_type=GifsDataResponse
    ).gif_categories


def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
    """

    Parameters:
    ---------------

        - from_timestamp: int - (optional)
        - number: int - (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.HIDDEN_V1}/chats",
        params=params, data_type=ChatRoomsResponse
    )


def get_main_chat_rooms(self, from_timestamp: int = None) -> ChatRoomsResponse:
    self._check_authorization()
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/main_list",
        params=params, data_type=ChatRoomsResponse
    )


def get_messages(self, chat_room_id: int, **params) -> List[Message]:
    """

    Parameters:
    ---------------
        - from_message_id: int - (optional)
        - to_message_id: int - (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{chat_room_id}/messages",
        params=params, data_type=MessagesResponse
    ).messages


# def get_chat_room_notification_settings(self, room_id: int) -> Settings:
#     self._check_authorization()
#     return self._make_request(
#         "GET", endpoint=f"{Endpoints.NOTIFICATION_SETTINGS_V2}/chat_rooms/{room_id}",
#         data_type=AdditionalSettingsResponse
#     ).settings


def get_request_chat_rooms(self, from_timestamp: int = None) -> ChatRoomsResponse:
    self._check_authorization()
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/request_list",
        params=params, data_type=ChatRoomsResponse
    )


def get_chat_room(self, chat_room_id: int) -> ChatRoom:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{chat_room_id}",
        data_type=ChatRoomResponse
    ).chat


def get_sticker_packs(self) -> List[StickerPack]:
    return self._make_request(
        "GET", endpoint=Endpoints.STICKER_PACKS_V2,
        data_type=StickerPacksResponse
    ).sticker_packs


def get_total_chat_requests(self) -> int:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/total_chat_request",
        data_type=TotalChatRequestResponse
    ).total


def hide_chat(self, chat_room_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.HIDDEN_V1}/chats",
        payload={"chat_room_id": chat_room_id},
    )


def invite_to_chat(self, chat_room_id: int, user_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{chat_room_id}/invite",
        payload={"with_user_ids[]": user_ids},
    )


def kick_users_from_chat(self, chat_room_id: int, user_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{chat_room_id}/kick",
        payload={"with_user_ids[]": user_ids},
    )


def pin_chat(self, room_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{room_id}/pinned"
    )


def read_attachment(
        self,
        chat_room_id: int,
        attachment_msg_ids: List[int]
        # ↑ @Body @NotNull ReadAttachmentRequest readAttachmentRequest
):
    pass


def read_message(self, chat_room_id: int, message_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/{chat_room_id}/messages/{message_id}/read",
    )


def read_video_message(
        self,
        room_id: int,
        video_msg_ids: List[int]
):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{room_id}/videos_read",
        payload={"video_msg_ids": video_msg_ids},
    )


def refresh_chat_rooms(self, from_time: int = None) -> ChatRoomsResponse:
    self._check_authorization()
    params = {}
    if from_time:
        params["from_time"] = from_time
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/update",
        params=params, data_type=ChatRoomsResponse
    )


def remove_chat_rooms(self, chat_room_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V2}/mass_destroy",
        payload={"chat_room_ids[]": chat_room_ids},
    )


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
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V3}/{chat_room_id}/report",
        payload={
            "chat_room_id": chat_room_id,
            "opponent_id": opponent_id,
            "category_id": category_id,
            "reason": reason,
            "screenshot_filename": screenshot_filename,
            "screenshot_2_filename": screenshot_2_filename,
            "screenshot_3_filename": screenshot_3_filename,
            "screenshot_4_filename": screenshot_4_filename,
        },
    )


def send_media_screenshot_notification(self, room_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{room_id}/screen_captured"
    )


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
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.CHAT_ROOMS_V3}/{chat_room_id}/messages/new",
        payload={
            "chat_room_id": chat_room_id,
            "message_type": message_type,
            "call_type": call_type,
            "text": text,
            "font_size": font_size,
            "gif_image_id": gif_image_id,
            "attachment_file_name": attachment_file_name,
            "sticker_pack_id": sticker_pack_id,
            "video_file_name": video_file_name,
        }, data_type=MessageResponse
    )


def set_notification_settings(
        self,
        chat_room_id: int,
        notification_chat: int
) -> Settings:
    # NotificationSettingResponse
    pass


def unhide_chat(self, chat_room_ids: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.HIDDEN_V1}/chats",
        params={"chat_room_ids[]": chat_room_ids},
    )


def unpin_chat(self, chat_room_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.CHAT_ROOMS_V1}/{chat_room_id}/pinned"
    )
