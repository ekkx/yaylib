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

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.in_app_purchase_product import InAppPurchaseProduct
from yaylib.models.product_quota import ProductQuota
from typing import Optional, Set
from typing_extensions import Self

class InAppPurchaseProductsResponse(BaseModel):
    """
    InAppPurchaseProductsResponse
    """ # noqa: E501
    iap_products: Optional[List[InAppPurchaseProduct]] = None
    quota: Optional[ProductQuota] = None
    __properties: ClassVar[List[str]] = ["iap_products", "quota"]

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
        """Create an instance of InAppPurchaseProductsResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in iap_products (list)
        _items = []
        if self.iap_products:
            for _item_iap_products in self.iap_products:
                if _item_iap_products:
                    _items.append(_item_iap_products.to_dict())
            _dict['iap_products'] = _items
        # override the default output from pydantic by calling `to_dict()` of quota
        if self.quota:
            _dict['quota'] = self.quota.to_dict()
        # set to None if iap_products (nullable) is None
        # and model_fields_set contains the field
        if self.iap_products is None and "iap_products" in self.model_fields_set:
            _dict['iap_products'] = None

        # set to None if quota (nullable) is None
        # and model_fields_set contains the field
        if self.quota is None and "quota" in self.model_fields_set:
            _dict['quota'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of InAppPurchaseProductsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "iap_products": [InAppPurchaseProduct.from_dict(_item) for _item in obj["iap_products"]] if obj.get("iap_products") is not None else None,
            "quota": ProductQuota.from_dict(obj["quota"]) if obj.get("quota") is not None else None
        })
        return _obj


