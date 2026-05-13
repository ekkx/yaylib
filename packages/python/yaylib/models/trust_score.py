# coding: utf-8

"""
    Yay! API

    No description provided

    Schema version: 4.26.1
    Code generated; DO NOT EDIT

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, StrictBool, StrictFloat, StrictInt
from typing import Any, ClassVar, Dict, List, Optional, Union
from yaylib.models.model_user_rank import ModelUserRank
from typing import Optional, Set
from typing_extensions import Self

class TrustScore(BaseModel):
    """
    TrustScore
    """ # noqa: E501
    calculated: Optional[StrictBool] = None
    grade: Optional[ModelUserRank] = None
    score: Optional[Union[StrictFloat, StrictInt]] = None
    __properties: ClassVar[List[str]] = ["calculated", "grade", "score"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of TrustScore from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of grade
        if self.grade:
            _dict['grade'] = self.grade.to_dict()
        # set to None if calculated (nullable) is None
        # and model_fields_set contains the field
        if self.calculated is None and "calculated" in self.model_fields_set:
            _dict['calculated'] = None

        # set to None if grade (nullable) is None
        # and model_fields_set contains the field
        if self.grade is None and "grade" in self.model_fields_set:
            _dict['grade'] = None

        # set to None if score (nullable) is None
        # and model_fields_set contains the field
        if self.score is None and "score" in self.model_fields_set:
            _dict['score'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TrustScore from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "calculated": obj.get("calculated"),
            "grade": ModelUserRank.from_dict(obj["grade"]) if obj.get("grade") is not None else None,
            "score": obj.get("score")
        })
        return _obj


