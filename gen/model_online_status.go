
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the OnlineStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnlineStatus{}

// OnlineStatus struct for OnlineStatus
type OnlineStatus struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnlineStatus OnlineStatus

// NewOnlineStatus instantiates a new OnlineStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnlineStatus() *OnlineStatus {
	this := OnlineStatus{}
	return &this
}

// NewOnlineStatusWithDefaults instantiates a new OnlineStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnlineStatusWithDefaults() *OnlineStatus {
	this := OnlineStatus{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineStatus) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineStatus) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *OnlineStatus) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *OnlineStatus) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *OnlineStatus) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *OnlineStatus) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o OnlineStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnlineStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnlineStatus) UnmarshalJSON(data []byte) (err error) {
	varOnlineStatus := _OnlineStatus{}

	err = json.Unmarshal(data, &varOnlineStatus)

	if err != nil {
		return err
	}

	*o = OnlineStatus(varOnlineStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnlineStatus struct {
	value *OnlineStatus
	isSet bool
}

func (v NullableOnlineStatus) Get() *OnlineStatus {
	return v.value
}

func (v *NullableOnlineStatus) Set(val *OnlineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineStatus(val *OnlineStatus) *NullableOnlineStatus {
	return &NullableOnlineStatus{value: val, isSet: true}
}

func (v NullableOnlineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


