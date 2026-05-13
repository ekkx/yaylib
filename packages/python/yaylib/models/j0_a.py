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


class J0A(str, Enum):
    """
    J0A
    """

    """
    allowed enum values
    """
    PUBLICTIMELINE = 'PublicTimeline'
    FOLLOWINGTIMELINE = 'FollowingTimeline'
    GROUPTIMELINE = 'GroupTimeline'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of J0A from a JSON string"""
        return cls(json.loads(json_str))


