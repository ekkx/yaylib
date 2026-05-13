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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class ModelInAppPurchaseProduct(BaseModel):
    """
    ModelInAppPurchaseProduct
    """ # noqa: E501
    amount: Optional[Union[StrictFloat, StrictInt]] = None
    bonus: Optional[Union[StrictFloat, StrictInt]] = None
    currency_code: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    price: Optional[Union[StrictFloat, StrictInt]] = None
    purchasable: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["amount", "bonus", "currency_code", "id", "price", "purchasable"]

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
        """Create an instance of ModelInAppPurchaseProduct from a JSON string"""
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
        # set to None if amount (nullable) is None
        # and model_fields_set contains the field
        if self.amount is None and "amount" in self.model_fields_set:
            _dict['amount'] = None

        # set to None if bonus (nullable) is None
        # and model_fields_set contains the field
        if self.bonus is None and "bonus" in self.model_fields_set:
            _dict['bonus'] = None

        # set to None if currency_code (nullable) is None
        # and model_fields_set contains the field
        if self.currency_code is None and "currency_code" in self.model_fields_set:
            _dict['currency_code'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        # set to None if purchasable (nullable) is None
        # and model_fields_set contains the field
        if self.purchasable is None and "purchasable" in self.model_fields_set:
            _dict['purchasable'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelInAppPurchaseProduct from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "bonus": obj.get("bonus"),
            "currency_code": obj.get("currency_code"),
            "id": obj.get("id"),
            "price": obj.get("price"),
            "purchasable": obj.get("purchasable")
        })
        return _obj


