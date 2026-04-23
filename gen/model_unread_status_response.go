
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UnreadStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnreadStatusResponse{}

// UnreadStatusResponse struct for UnreadStatusResponse
type UnreadStatusResponse struct {
	IsUnread NullableBool `json:"is_unread,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnreadStatusResponse UnreadStatusResponse

// NewUnreadStatusResponse instantiates a new UnreadStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnreadStatusResponse() *UnreadStatusResponse {
	this := UnreadStatusResponse{}
	return &this
}

// NewUnreadStatusResponseWithDefaults instantiates a new UnreadStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnreadStatusResponseWithDefaults() *UnreadStatusResponse {
	this := UnreadStatusResponse{}
	return &this
}

// GetIsUnread returns the IsUnread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnreadStatusResponse) GetIsUnread() bool {
	if o == nil || IsNil(o.IsUnread.Get()) {
		var ret bool
		return ret
	}
	return *o.IsUnread.Get()
}

// GetIsUnreadOk returns a tuple with the IsUnread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnreadStatusResponse) GetIsUnreadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsUnread.Get(), o.IsUnread.IsSet()
}

// HasIsUnread returns a boolean if a field has been set.
func (o *UnreadStatusResponse) HasIsUnread() bool {
	if o != nil && o.IsUnread.IsSet() {
		return true
	}

	return false
}

// SetIsUnread gets a reference to the given NullableBool and assigns it to the IsUnread field.
func (o *UnreadStatusResponse) SetIsUnread(v bool) {
	o.IsUnread.Set(&v)
}
// SetIsUnreadNil sets the value for IsUnread to be an explicit nil
func (o *UnreadStatusResponse) SetIsUnreadNil() {
	o.IsUnread.Set(nil)
}

// UnsetIsUnread ensures that no value is present for IsUnread, not even an explicit nil
func (o *UnreadStatusResponse) UnsetIsUnread() {
	o.IsUnread.Unset()
}

func (o UnreadStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnreadStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsUnread.IsSet() {
		toSerialize["is_unread"] = o.IsUnread.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnreadStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varUnreadStatusResponse := _UnreadStatusResponse{}

	err = json.Unmarshal(data, &varUnreadStatusResponse)

	if err != nil {
		return err
	}

	*o = UnreadStatusResponse(varUnreadStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_unread")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnreadStatusResponse struct {
	value *UnreadStatusResponse
	isSet bool
}

func (v NullableUnreadStatusResponse) Get() *UnreadStatusResponse {
	return v.value
}

func (v *NullableUnreadStatusResponse) Set(val *UnreadStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnreadStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnreadStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnreadStatusResponse(val *UnreadStatusResponse) *NullableUnreadStatusResponse {
	return &NullableUnreadStatusResponse{value: val, isSet: true}
}

func (v NullableUnreadStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnreadStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


