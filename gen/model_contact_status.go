
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ContactStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactStatus{}

// ContactStatus struct for ContactStatus
type ContactStatus struct {
	Status NullableString `json:"status,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactStatus ContactStatus

// NewContactStatus instantiates a new ContactStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactStatus() *ContactStatus {
	this := ContactStatus{}
	return &this
}

// NewContactStatusWithDefaults instantiates a new ContactStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactStatusWithDefaults() *ContactStatus {
	this := ContactStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactStatus) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ContactStatus) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ContactStatus) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ContactStatus) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ContactStatus) UnsetStatus() {
	o.Status.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactStatus) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactStatus) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ContactStatus) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *ContactStatus) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ContactStatus) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ContactStatus) UnsetUserId() {
	o.UserId.Unset()
}

func (o ContactStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContactStatus) UnmarshalJSON(data []byte) (err error) {
	varContactStatus := _ContactStatus{}

	err = json.Unmarshal(data, &varContactStatus)

	if err != nil {
		return err
	}

	*o = ContactStatus(varContactStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactStatus struct {
	value *ContactStatus
	isSet bool
}

func (v NullableContactStatus) Get() *ContactStatus {
	return v.value
}

func (v *NullableContactStatus) Set(val *ContactStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableContactStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableContactStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactStatus(val *ContactStatus) *NullableContactStatus {
	return &NullableContactStatus{value: val, isSet: true}
}

func (v NullableContactStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


