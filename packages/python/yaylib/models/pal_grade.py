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


class PalGrade(str, Enum):
    """
    PalGrade
    """

    """
    allowed enum values
    """
    EGG = 'Egg'
    PAL = 'Pal'
    SUPER_PAL = 'Super Pal'
    ULTIMATE_PAL = 'Ultimate Pal'

    @classmethod
    def __get_pydantic_core_schema__(cls, source_type, handler):
        # Accept any string the server sends, not only the values
        # enumerated above, so a value added server-side still
        # decodes. A typed constant passed in stays a member (so
        # callers and request serializers keep working); an unknown
        # string passes through unchanged.
        from pydantic_core import core_schema
        def _v(value):
            if isinstance(value, cls):
                return value
            try:
                return cls(value)
            except ValueError:
                return value
        return core_schema.no_info_plain_validator_function(
            _v,
            serialization=core_schema.plain_serializer_function_ser_schema(
                lambda value: value.value if isinstance(value, cls) else value
            ),
        )

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of PalGrade from a JSON string"""
        return cls(json.loads(json_str))


