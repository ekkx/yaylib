
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ContactStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactStatusResponse{}

// ContactStatusResponse struct for ContactStatusResponse
type ContactStatusResponse struct {
	Contacts map[string]ContactStatus `json:"contacts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactStatusResponse ContactStatusResponse

// NewContactStatusResponse instantiates a new ContactStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactStatusResponse() *ContactStatusResponse {
	this := ContactStatusResponse{}
	return &this
}

// NewContactStatusResponseWithDefaults instantiates a new ContactStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactStatusResponseWithDefaults() *ContactStatusResponse {
	this := ContactStatusResponse{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ContactStatusResponse) GetContacts() map[string]ContactStatus {
	if o == nil || IsNil(o.Contacts) {
		var ret map[string]ContactStatus
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactStatusResponse) GetContactsOk() (map[string]ContactStatus, bool) {
	if o == nil || IsNil(o.Contacts) {
		return map[string]ContactStatus{}, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ContactStatusResponse) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given map[string]ContactStatus and assigns it to the Contacts field.
func (o *ContactStatusResponse) SetContacts(v map[string]ContactStatus) {
	o.Contacts = v
}

func (o ContactStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContactStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varContactStatusResponse := _ContactStatusResponse{}

	err = json.Unmarshal(data, &varContactStatusResponse)

	if err != nil {
		return err
	}

	*o = ContactStatusResponse(varContactStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contacts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactStatusResponse struct {
	value *ContactStatusResponse
	isSet bool
}

func (v NullableContactStatusResponse) Get() *ContactStatusResponse {
	return v.value
}

func (v *NullableContactStatusResponse) Set(val *ContactStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactStatusResponse(val *ContactStatusResponse) *NullableContactStatusResponse {
	return &NullableContactStatusResponse{value: val, isSet: true}
}

func (v NullableContactStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


