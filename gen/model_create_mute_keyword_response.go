
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateMuteKeywordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMuteKeywordResponse{}

// CreateMuteKeywordResponse struct for CreateMuteKeywordResponse
type CreateMuteKeywordResponse struct {
	HiddenWord NullableMuteKeyword `json:"hidden_word,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMuteKeywordResponse CreateMuteKeywordResponse

// NewCreateMuteKeywordResponse instantiates a new CreateMuteKeywordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMuteKeywordResponse() *CreateMuteKeywordResponse {
	this := CreateMuteKeywordResponse{}
	return &this
}

// NewCreateMuteKeywordResponseWithDefaults instantiates a new CreateMuteKeywordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMuteKeywordResponseWithDefaults() *CreateMuteKeywordResponse {
	this := CreateMuteKeywordResponse{}
	return &this
}

// GetHiddenWord returns the HiddenWord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMuteKeywordResponse) GetHiddenWord() MuteKeyword {
	if o == nil || IsNil(o.HiddenWord.Get()) {
		var ret MuteKeyword
		return ret
	}
	return *o.HiddenWord.Get()
}

// GetHiddenWordOk returns a tuple with the HiddenWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMuteKeywordResponse) GetHiddenWordOk() (*MuteKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenWord.Get(), o.HiddenWord.IsSet()
}

// HasHiddenWord returns a boolean if a field has been set.
func (o *CreateMuteKeywordResponse) HasHiddenWord() bool {
	if o != nil && o.HiddenWord.IsSet() {
		return true
	}

	return false
}

// SetHiddenWord gets a reference to the given NullableMuteKeyword and assigns it to the HiddenWord field.
func (o *CreateMuteKeywordResponse) SetHiddenWord(v MuteKeyword) {
	o.HiddenWord.Set(&v)
}
// SetHiddenWordNil sets the value for HiddenWord to be an explicit nil
func (o *CreateMuteKeywordResponse) SetHiddenWordNil() {
	o.HiddenWord.Set(nil)
}

// UnsetHiddenWord ensures that no value is present for HiddenWord, not even an explicit nil
func (o *CreateMuteKeywordResponse) UnsetHiddenWord() {
	o.HiddenWord.Unset()
}

func (o CreateMuteKeywordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMuteKeywordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HiddenWord.IsSet() {
		toSerialize["hidden_word"] = o.HiddenWord.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMuteKeywordResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateMuteKeywordResponse := _CreateMuteKeywordResponse{}

	err = json.Unmarshal(data, &varCreateMuteKeywordResponse)

	if err != nil {
		return err
	}

	*o = CreateMuteKeywordResponse(varCreateMuteKeywordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hidden_word")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMuteKeywordResponse struct {
	value *CreateMuteKeywordResponse
	isSet bool
}

func (v NullableCreateMuteKeywordResponse) Get() *CreateMuteKeywordResponse {
	return v.value
}

func (v *NullableCreateMuteKeywordResponse) Set(val *CreateMuteKeywordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMuteKeywordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMuteKeywordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMuteKeywordResponse(val *CreateMuteKeywordResponse) *NullableCreateMuteKeywordResponse {
	return &NullableCreateMuteKeywordResponse{value: val, isSet: true}
}

func (v NullableCreateMuteKeywordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMuteKeywordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


