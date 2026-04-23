
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InAppPurchaseProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseProductsResponse{}

// InAppPurchaseProductsResponse struct for InAppPurchaseProductsResponse
type InAppPurchaseProductsResponse struct {
	IapProducts []InAppPurchaseProduct `json:"iap_products,omitempty"`
	Quota NullableProductQuota `json:"quota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InAppPurchaseProductsResponse InAppPurchaseProductsResponse

// NewInAppPurchaseProductsResponse instantiates a new InAppPurchaseProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseProductsResponse() *InAppPurchaseProductsResponse {
	this := InAppPurchaseProductsResponse{}
	return &this
}

// NewInAppPurchaseProductsResponseWithDefaults instantiates a new InAppPurchaseProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseProductsResponseWithDefaults() *InAppPurchaseProductsResponse {
	this := InAppPurchaseProductsResponse{}
	return &this
}

// GetIapProducts returns the IapProducts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProductsResponse) GetIapProducts() []InAppPurchaseProduct {
	if o == nil {
		var ret []InAppPurchaseProduct
		return ret
	}
	return o.IapProducts
}

// GetIapProductsOk returns a tuple with the IapProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProductsResponse) GetIapProductsOk() ([]InAppPurchaseProduct, bool) {
	if o == nil || IsNil(o.IapProducts) {
		return nil, false
	}
	return o.IapProducts, true
}

// HasIapProducts returns a boolean if a field has been set.
func (o *InAppPurchaseProductsResponse) HasIapProducts() bool {
	if o != nil && !IsNil(o.IapProducts) {
		return true
	}

	return false
}

// SetIapProducts gets a reference to the given []InAppPurchaseProduct and assigns it to the IapProducts field.
func (o *InAppPurchaseProductsResponse) SetIapProducts(v []InAppPurchaseProduct) {
	o.IapProducts = v
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProductsResponse) GetQuota() ProductQuota {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret ProductQuota
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProductsResponse) GetQuotaOk() (*ProductQuota, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *InAppPurchaseProductsResponse) HasQuota() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableProductQuota and assigns it to the Quota field.
func (o *InAppPurchaseProductsResponse) SetQuota(v ProductQuota) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *InAppPurchaseProductsResponse) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *InAppPurchaseProductsResponse) UnsetQuota() {
	o.Quota.Unset()
}

func (o InAppPurchaseProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IapProducts != nil {
		toSerialize["iap_products"] = o.IapProducts
	}
	if o.Quota.IsSet() {
		toSerialize["quota"] = o.Quota.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InAppPurchaseProductsResponse) UnmarshalJSON(data []byte) (err error) {
	varInAppPurchaseProductsResponse := _InAppPurchaseProductsResponse{}

	err = json.Unmarshal(data, &varInAppPurchaseProductsResponse)

	if err != nil {
		return err
	}

	*o = InAppPurchaseProductsResponse(varInAppPurchaseProductsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "iap_products")
		delete(additionalProperties, "quota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInAppPurchaseProductsResponse struct {
	value *InAppPurchaseProductsResponse
	isSet bool
}

func (v NullableInAppPurchaseProductsResponse) Get() *InAppPurchaseProductsResponse {
	return v.value
}

func (v *NullableInAppPurchaseProductsResponse) Set(val *InAppPurchaseProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseProductsResponse(val *InAppPurchaseProductsResponse) *NullableInAppPurchaseProductsResponse {
	return &NullableInAppPurchaseProductsResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


