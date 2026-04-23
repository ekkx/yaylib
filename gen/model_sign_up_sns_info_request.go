
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SignUpSnsInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignUpSnsInfoRequest{}

// SignUpSnsInfoRequest struct for SignUpSnsInfoRequest
type SignUpSnsInfoRequest struct {
	Type NullableString `json:"type,omitempty"`
	Uid NullableString `json:"uid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignUpSnsInfoRequest SignUpSnsInfoRequest

// NewSignUpSnsInfoRequest instantiates a new SignUpSnsInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignUpSnsInfoRequest() *SignUpSnsInfoRequest {
	this := SignUpSnsInfoRequest{}
	return &this
}

// NewSignUpSnsInfoRequestWithDefaults instantiates a new SignUpSnsInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignUpSnsInfoRequestWithDefaults() *SignUpSnsInfoRequest {
	this := SignUpSnsInfoRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUpSnsInfoRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUpSnsInfoRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *SignUpSnsInfoRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *SignUpSnsInfoRequest) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *SignUpSnsInfoRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *SignUpSnsInfoRequest) UnsetType() {
	o.Type.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignUpSnsInfoRequest) GetUid() string {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret string
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignUpSnsInfoRequest) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *SignUpSnsInfoRequest) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableString and assigns it to the Uid field.
func (o *SignUpSnsInfoRequest) SetUid(v string) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *SignUpSnsInfoRequest) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *SignUpSnsInfoRequest) UnsetUid() {
	o.Uid.Unset()
}

func (o SignUpSnsInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignUpSnsInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignUpSnsInfoRequest) UnmarshalJSON(data []byte) (err error) {
	varSignUpSnsInfoRequest := _SignUpSnsInfoRequest{}

	err = json.Unmarshal(data, &varSignUpSnsInfoRequest)

	if err != nil {
		return err
	}

	*o = SignUpSnsInfoRequest(varSignUpSnsInfoRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "uid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignUpSnsInfoRequest struct {
	value *SignUpSnsInfoRequest
	isSet bool
}

func (v NullableSignUpSnsInfoRequest) Get() *SignUpSnsInfoRequest {
	return v.value
}

func (v *NullableSignUpSnsInfoRequest) Set(val *SignUpSnsInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignUpSnsInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignUpSnsInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignUpSnsInfoRequest(val *SignUpSnsInfoRequest) *NullableSignUpSnsInfoRequest {
	return &NullableSignUpSnsInfoRequest{value: val, isSet: true}
}

func (v NullableSignUpSnsInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignUpSnsInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


