
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelCreateGroupQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCreateGroupQuota{}

// ModelCreateGroupQuota struct for ModelCreateGroupQuota
type ModelCreateGroupQuota struct {
	RemainingQuota NullableInt32 `json:"remaining_quota,omitempty"`
	UsedQuota NullableInt32 `json:"used_quota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelCreateGroupQuota ModelCreateGroupQuota

// NewModelCreateGroupQuota instantiates a new ModelCreateGroupQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCreateGroupQuota() *ModelCreateGroupQuota {
	this := ModelCreateGroupQuota{}
	return &this
}

// NewModelCreateGroupQuotaWithDefaults instantiates a new ModelCreateGroupQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCreateGroupQuotaWithDefaults() *ModelCreateGroupQuota {
	this := ModelCreateGroupQuota{}
	return &this
}

// GetRemainingQuota returns the RemainingQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreateGroupQuota) GetRemainingQuota() int32 {
	if o == nil || IsNil(o.RemainingQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.RemainingQuota.Get()
}

// GetRemainingQuotaOk returns a tuple with the RemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreateGroupQuota) GetRemainingQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemainingQuota.Get(), o.RemainingQuota.IsSet()
}

// HasRemainingQuota returns a boolean if a field has been set.
func (o *ModelCreateGroupQuota) HasRemainingQuota() bool {
	if o != nil && o.RemainingQuota.IsSet() {
		return true
	}

	return false
}

// SetRemainingQuota gets a reference to the given NullableInt32 and assigns it to the RemainingQuota field.
func (o *ModelCreateGroupQuota) SetRemainingQuota(v int32) {
	o.RemainingQuota.Set(&v)
}
// SetRemainingQuotaNil sets the value for RemainingQuota to be an explicit nil
func (o *ModelCreateGroupQuota) SetRemainingQuotaNil() {
	o.RemainingQuota.Set(nil)
}

// UnsetRemainingQuota ensures that no value is present for RemainingQuota, not even an explicit nil
func (o *ModelCreateGroupQuota) UnsetRemainingQuota() {
	o.RemainingQuota.Unset()
}

// GetUsedQuota returns the UsedQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreateGroupQuota) GetUsedQuota() int32 {
	if o == nil || IsNil(o.UsedQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.UsedQuota.Get()
}

// GetUsedQuotaOk returns a tuple with the UsedQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreateGroupQuota) GetUsedQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedQuota.Get(), o.UsedQuota.IsSet()
}

// HasUsedQuota returns a boolean if a field has been set.
func (o *ModelCreateGroupQuota) HasUsedQuota() bool {
	if o != nil && o.UsedQuota.IsSet() {
		return true
	}

	return false
}

// SetUsedQuota gets a reference to the given NullableInt32 and assigns it to the UsedQuota field.
func (o *ModelCreateGroupQuota) SetUsedQuota(v int32) {
	o.UsedQuota.Set(&v)
}
// SetUsedQuotaNil sets the value for UsedQuota to be an explicit nil
func (o *ModelCreateGroupQuota) SetUsedQuotaNil() {
	o.UsedQuota.Set(nil)
}

// UnsetUsedQuota ensures that no value is present for UsedQuota, not even an explicit nil
func (o *ModelCreateGroupQuota) UnsetUsedQuota() {
	o.UsedQuota.Unset()
}

func (o ModelCreateGroupQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCreateGroupQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RemainingQuota.IsSet() {
		toSerialize["remaining_quota"] = o.RemainingQuota.Get()
	}
	if o.UsedQuota.IsSet() {
		toSerialize["used_quota"] = o.UsedQuota.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelCreateGroupQuota) UnmarshalJSON(data []byte) (err error) {
	varModelCreateGroupQuota := _ModelCreateGroupQuota{}

	err = json.Unmarshal(data, &varModelCreateGroupQuota)

	if err != nil {
		return err
	}

	*o = ModelCreateGroupQuota(varModelCreateGroupQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "remaining_quota")
		delete(additionalProperties, "used_quota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelCreateGroupQuota struct {
	value *ModelCreateGroupQuota
	isSet bool
}

func (v NullableModelCreateGroupQuota) Get() *ModelCreateGroupQuota {
	return v.value
}

func (v *NullableModelCreateGroupQuota) Set(val *ModelCreateGroupQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCreateGroupQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCreateGroupQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCreateGroupQuota(val *ModelCreateGroupQuota) *NullableModelCreateGroupQuota {
	return &NullableModelCreateGroupQuota{value: val, isSet: true}
}

func (v NullableModelCreateGroupQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCreateGroupQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


