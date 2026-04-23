
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateGroupQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupQuota{}

// CreateGroupQuota struct for CreateGroupQuota
type CreateGroupQuota struct {
	RemainingQuota NullableInt32 `json:"remaining_quota,omitempty"`
	UsedQuota NullableInt32 `json:"used_quota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateGroupQuota CreateGroupQuota

// NewCreateGroupQuota instantiates a new CreateGroupQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupQuota() *CreateGroupQuota {
	this := CreateGroupQuota{}
	return &this
}

// NewCreateGroupQuotaWithDefaults instantiates a new CreateGroupQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupQuotaWithDefaults() *CreateGroupQuota {
	this := CreateGroupQuota{}
	return &this
}

// GetRemainingQuota returns the RemainingQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupQuota) GetRemainingQuota() int32 {
	if o == nil || IsNil(o.RemainingQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.RemainingQuota.Get()
}

// GetRemainingQuotaOk returns a tuple with the RemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupQuota) GetRemainingQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemainingQuota.Get(), o.RemainingQuota.IsSet()
}

// HasRemainingQuota returns a boolean if a field has been set.
func (o *CreateGroupQuota) HasRemainingQuota() bool {
	if o != nil && o.RemainingQuota.IsSet() {
		return true
	}

	return false
}

// SetRemainingQuota gets a reference to the given NullableInt32 and assigns it to the RemainingQuota field.
func (o *CreateGroupQuota) SetRemainingQuota(v int32) {
	o.RemainingQuota.Set(&v)
}
// SetRemainingQuotaNil sets the value for RemainingQuota to be an explicit nil
func (o *CreateGroupQuota) SetRemainingQuotaNil() {
	o.RemainingQuota.Set(nil)
}

// UnsetRemainingQuota ensures that no value is present for RemainingQuota, not even an explicit nil
func (o *CreateGroupQuota) UnsetRemainingQuota() {
	o.RemainingQuota.Unset()
}

// GetUsedQuota returns the UsedQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupQuota) GetUsedQuota() int32 {
	if o == nil || IsNil(o.UsedQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.UsedQuota.Get()
}

// GetUsedQuotaOk returns a tuple with the UsedQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupQuota) GetUsedQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedQuota.Get(), o.UsedQuota.IsSet()
}

// HasUsedQuota returns a boolean if a field has been set.
func (o *CreateGroupQuota) HasUsedQuota() bool {
	if o != nil && o.UsedQuota.IsSet() {
		return true
	}

	return false
}

// SetUsedQuota gets a reference to the given NullableInt32 and assigns it to the UsedQuota field.
func (o *CreateGroupQuota) SetUsedQuota(v int32) {
	o.UsedQuota.Set(&v)
}
// SetUsedQuotaNil sets the value for UsedQuota to be an explicit nil
func (o *CreateGroupQuota) SetUsedQuotaNil() {
	o.UsedQuota.Set(nil)
}

// UnsetUsedQuota ensures that no value is present for UsedQuota, not even an explicit nil
func (o *CreateGroupQuota) UnsetUsedQuota() {
	o.UsedQuota.Unset()
}

func (o CreateGroupQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupQuota) ToMap() (map[string]interface{}, error) {
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

func (o *CreateGroupQuota) UnmarshalJSON(data []byte) (err error) {
	varCreateGroupQuota := _CreateGroupQuota{}

	err = json.Unmarshal(data, &varCreateGroupQuota)

	if err != nil {
		return err
	}

	*o = CreateGroupQuota(varCreateGroupQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "remaining_quota")
		delete(additionalProperties, "used_quota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateGroupQuota struct {
	value *CreateGroupQuota
	isSet bool
}

func (v NullableCreateGroupQuota) Get() *CreateGroupQuota {
	return v.value
}

func (v *NullableCreateGroupQuota) Set(val *CreateGroupQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupQuota(val *CreateGroupQuota) *NullableCreateGroupQuota {
	return &NullableCreateGroupQuota{value: val, isSet: true}
}

func (v NullableCreateGroupQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


