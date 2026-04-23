
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupResponse{}

// CreateGroupResponse struct for CreateGroupResponse
type CreateGroupResponse struct {
	GroupId NullableInt64 `json:"group_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateGroupResponse CreateGroupResponse

// NewCreateGroupResponse instantiates a new CreateGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupResponse() *CreateGroupResponse {
	this := CreateGroupResponse{}
	return &this
}

// NewCreateGroupResponseWithDefaults instantiates a new CreateGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupResponseWithDefaults() *CreateGroupResponse {
	this := CreateGroupResponse{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupResponse) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupResponse) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *CreateGroupResponse) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *CreateGroupResponse) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *CreateGroupResponse) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *CreateGroupResponse) UnsetGroupId() {
	o.GroupId.Unset()
}

func (o CreateGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateGroupResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateGroupResponse := _CreateGroupResponse{}

	err = json.Unmarshal(data, &varCreateGroupResponse)

	if err != nil {
		return err
	}

	*o = CreateGroupResponse(varCreateGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateGroupResponse struct {
	value *CreateGroupResponse
	isSet bool
}

func (v NullableCreateGroupResponse) Get() *CreateGroupResponse {
	return v.value
}

func (v *NullableCreateGroupResponse) Set(val *CreateGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupResponse(val *CreateGroupResponse) *NullableCreateGroupResponse {
	return &NullableCreateGroupResponse{value: val, isSet: true}
}

func (v NullableCreateGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


