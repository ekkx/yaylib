
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MuteKeywordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteKeywordResponse{}

// MuteKeywordResponse struct for MuteKeywordResponse
type MuteKeywordResponse struct {
	HiddenWords []MuteKeyword `json:"hidden_words,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MuteKeywordResponse MuteKeywordResponse

// NewMuteKeywordResponse instantiates a new MuteKeywordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteKeywordResponse() *MuteKeywordResponse {
	this := MuteKeywordResponse{}
	return &this
}

// NewMuteKeywordResponseWithDefaults instantiates a new MuteKeywordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteKeywordResponseWithDefaults() *MuteKeywordResponse {
	this := MuteKeywordResponse{}
	return &this
}

// GetHiddenWords returns the HiddenWords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeywordResponse) GetHiddenWords() []MuteKeyword {
	if o == nil {
		var ret []MuteKeyword
		return ret
	}
	return o.HiddenWords
}

// GetHiddenWordsOk returns a tuple with the HiddenWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeywordResponse) GetHiddenWordsOk() ([]MuteKeyword, bool) {
	if o == nil || IsNil(o.HiddenWords) {
		return nil, false
	}
	return o.HiddenWords, true
}

// HasHiddenWords returns a boolean if a field has been set.
func (o *MuteKeywordResponse) HasHiddenWords() bool {
	if o != nil && !IsNil(o.HiddenWords) {
		return true
	}

	return false
}

// SetHiddenWords gets a reference to the given []MuteKeyword and assigns it to the HiddenWords field.
func (o *MuteKeywordResponse) SetHiddenWords(v []MuteKeyword) {
	o.HiddenWords = v
}

func (o MuteKeywordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteKeywordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HiddenWords != nil {
		toSerialize["hidden_words"] = o.HiddenWords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MuteKeywordResponse) UnmarshalJSON(data []byte) (err error) {
	varMuteKeywordResponse := _MuteKeywordResponse{}

	err = json.Unmarshal(data, &varMuteKeywordResponse)

	if err != nil {
		return err
	}

	*o = MuteKeywordResponse(varMuteKeywordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hidden_words")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMuteKeywordResponse struct {
	value *MuteKeywordResponse
	isSet bool
}

func (v NullableMuteKeywordResponse) Get() *MuteKeywordResponse {
	return v.value
}

func (v *NullableMuteKeywordResponse) Set(val *MuteKeywordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteKeywordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteKeywordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteKeywordResponse(val *MuteKeywordResponse) *NullableMuteKeywordResponse {
	return &NullableMuteKeywordResponse{value: val, isSet: true}
}

func (v NullableMuteKeywordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteKeywordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


