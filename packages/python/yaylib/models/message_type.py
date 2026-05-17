# coding: utf-8

"""
    Yay! API

    No description provided

    Schema version: 4.26.1
    Code generated; DO NOT EDIT

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
from enum import Enum
from typing_extensions import Self


class MessageType(str, Enum):
    """
    MessageType
    """

    """
    allowed enum values
    """
    TEXT = 'text'
    IMAGE = 'image'
    ETERNAL_IMAGE = 'eternal_image'
    VIDEO = 'video'
    ETERNAL_VIDEO = 'eternal_video'
    SCREENSHOT_WARNING = 'screenshot_warning'
    GIF = 'gif'
    STICKER = 'sticker'
    INDIVIDUAL_CALL = 'individual_call'
    DELETED = 'deleted'
    USER_JOINED = 'user_joined'
    USER_LEAVE = 'user_leave'
    OWNER_KICK = 'owner_kick'

    @classmethod
    def __get_pydantic_core_schema__(cls, source_type, handler):
        # Accept any string the server sends, not only the values
        # enumerated above, so a value added server-side still
        # decodes. The typed constants stay usable for comparisons.
        from pydantic_core import core_schema
        return core_schema.str_schema()

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of MessageType from a JSON string"""
        return cls(json.loads(json_str))


