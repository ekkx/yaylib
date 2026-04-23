
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the HimaUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HimaUsersResponse{}

// HimaUsersResponse struct for HimaUsersResponse
type HimaUsersResponse struct {
	HimaUsers []UserWrapper `json:"hima_users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HimaUsersResponse HimaUsersResponse

// NewHimaUsersResponse instantiates a new HimaUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHimaUsersResponse() *HimaUsersResponse {
	this := HimaUsersResponse{}
	return &this
}

// NewHimaUsersResponseWithDefaults instantiates a new HimaUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHimaUsersResponseWithDefaults() *HimaUsersResponse {
	this := HimaUsersResponse{}
	return &this
}

// GetHimaUsers returns the HimaUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HimaUsersResponse) GetHimaUsers() []UserWrapper {
	if o == nil {
		var ret []UserWrapper
		return ret
	}
	return o.HimaUsers
}

// GetHimaUsersOk returns a tuple with the HimaUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HimaUsersResponse) GetHimaUsersOk() ([]UserWrapper, bool) {
	if o == nil || IsNil(o.HimaUsers) {
		return nil, false
	}
	return o.HimaUsers, true
}

// HasHimaUsers returns a boolean if a field has been set.
func (o *HimaUsersResponse) HasHimaUsers() bool {
	if o != nil && !IsNil(o.HimaUsers) {
		return true
	}

	return false
}

// SetHimaUsers gets a reference to the given []UserWrapper and assigns it to the HimaUsers field.
func (o *HimaUsersResponse) SetHimaUsers(v []UserWrapper) {
	o.HimaUsers = v
}

func (o HimaUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HimaUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HimaUsers != nil {
		toSerialize["hima_users"] = o.HimaUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HimaUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varHimaUsersResponse := _HimaUsersResponse{}

	err = json.Unmarshal(data, &varHimaUsersResponse)

	if err != nil {
		return err
	}

	*o = HimaUsersResponse(varHimaUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hima_users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHimaUsersResponse struct {
	value *HimaUsersResponse
	isSet bool
}

func (v NullableHimaUsersResponse) Get() *HimaUsersResponse {
	return v.value
}

func (v *NullableHimaUsersResponse) Set(val *HimaUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHimaUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHimaUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHimaUsersResponse(val *HimaUsersResponse) *NullableHimaUsersResponse {
	return &NullableHimaUsersResponse{value: val, isSet: true}
}

func (v NullableHimaUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHimaUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


