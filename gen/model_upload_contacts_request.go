
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UploadContactsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadContactsRequest{}

// UploadContactsRequest struct for UploadContactsRequest
type UploadContactsRequest struct {
	Contacts []Contact `json:"contacts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UploadContactsRequest UploadContactsRequest

// NewUploadContactsRequest instantiates a new UploadContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadContactsRequest() *UploadContactsRequest {
	this := UploadContactsRequest{}
	return &this
}

// NewUploadContactsRequestWithDefaults instantiates a new UploadContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadContactsRequestWithDefaults() *UploadContactsRequest {
	this := UploadContactsRequest{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadContactsRequest) GetContacts() []Contact {
	if o == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadContactsRequest) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *UploadContactsRequest) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *UploadContactsRequest) SetContacts(v []Contact) {
	o.Contacts = v
}

func (o UploadContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadContactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UploadContactsRequest) UnmarshalJSON(data []byte) (err error) {
	varUploadContactsRequest := _UploadContactsRequest{}

	err = json.Unmarshal(data, &varUploadContactsRequest)

	if err != nil {
		return err
	}

	*o = UploadContactsRequest(varUploadContactsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contacts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadContactsRequest struct {
	value *UploadContactsRequest
	isSet bool
}

func (v NullableUploadContactsRequest) Get() *UploadContactsRequest {
	return v.value
}

func (v *NullableUploadContactsRequest) Set(val *UploadContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadContactsRequest(val *UploadContactsRequest) *NullableUploadContactsRequest {
	return &NullableUploadContactsRequest{value: val, isSet: true}
}

func (v NullableUploadContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


