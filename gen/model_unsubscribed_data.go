
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UnsubscribedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnsubscribedData{}

// UnsubscribedData struct for UnsubscribedData
type UnsubscribedData struct {
	UserIds []int64 `json:"user_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnsubscribedData UnsubscribedData

// NewUnsubscribedData instantiates a new UnsubscribedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsubscribedData() *UnsubscribedData {
	this := UnsubscribedData{}
	return &this
}

// NewUnsubscribedDataWithDefaults instantiates a new UnsubscribedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsubscribedDataWithDefaults() *UnsubscribedData {
	this := UnsubscribedData{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnsubscribedData) GetUserIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnsubscribedData) GetUserIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *UnsubscribedData) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []int64 and assigns it to the UserIds field.
func (o *UnsubscribedData) SetUserIds(v []int64) {
	o.UserIds = v
}

func (o UnsubscribedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnsubscribedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserIds != nil {
		toSerialize["user_ids"] = o.UserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnsubscribedData) UnmarshalJSON(data []byte) (err error) {
	varUnsubscribedData := _UnsubscribedData{}

	err = json.Unmarshal(data, &varUnsubscribedData)

	if err != nil {
		return err
	}

	*o = UnsubscribedData(varUnsubscribedData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnsubscribedData struct {
	value *UnsubscribedData
	isSet bool
}

func (v NullableUnsubscribedData) Get() *UnsubscribedData {
	return v.value
}

func (v *NullableUnsubscribedData) Set(val *UnsubscribedData) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsubscribedData) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsubscribedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsubscribedData(val *UnsubscribedData) *NullableUnsubscribedData {
	return &NullableUnsubscribedData{value: val, isSet: true}
}

func (v NullableUnsubscribedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsubscribedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


