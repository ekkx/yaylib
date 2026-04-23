
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TextTranslationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextTranslationResponse{}

// TextTranslationResponse struct for TextTranslationResponse
type TextTranslationResponse struct {
	TranslatedText NullableString `json:"translated_text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TextTranslationResponse TextTranslationResponse

// NewTextTranslationResponse instantiates a new TextTranslationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextTranslationResponse() *TextTranslationResponse {
	this := TextTranslationResponse{}
	return &this
}

// NewTextTranslationResponseWithDefaults instantiates a new TextTranslationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextTranslationResponseWithDefaults() *TextTranslationResponse {
	this := TextTranslationResponse{}
	return &this
}

// GetTranslatedText returns the TranslatedText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextTranslationResponse) GetTranslatedText() string {
	if o == nil || IsNil(o.TranslatedText.Get()) {
		var ret string
		return ret
	}
	return *o.TranslatedText.Get()
}

// GetTranslatedTextOk returns a tuple with the TranslatedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextTranslationResponse) GetTranslatedTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TranslatedText.Get(), o.TranslatedText.IsSet()
}

// HasTranslatedText returns a boolean if a field has been set.
func (o *TextTranslationResponse) HasTranslatedText() bool {
	if o != nil && o.TranslatedText.IsSet() {
		return true
	}

	return false
}

// SetTranslatedText gets a reference to the given NullableString and assigns it to the TranslatedText field.
func (o *TextTranslationResponse) SetTranslatedText(v string) {
	o.TranslatedText.Set(&v)
}
// SetTranslatedTextNil sets the value for TranslatedText to be an explicit nil
func (o *TextTranslationResponse) SetTranslatedTextNil() {
	o.TranslatedText.Set(nil)
}

// UnsetTranslatedText ensures that no value is present for TranslatedText, not even an explicit nil
func (o *TextTranslationResponse) UnsetTranslatedText() {
	o.TranslatedText.Unset()
}

func (o TextTranslationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextTranslationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TranslatedText.IsSet() {
		toSerialize["translated_text"] = o.TranslatedText.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TextTranslationResponse) UnmarshalJSON(data []byte) (err error) {
	varTextTranslationResponse := _TextTranslationResponse{}

	err = json.Unmarshal(data, &varTextTranslationResponse)

	if err != nil {
		return err
	}

	*o = TextTranslationResponse(varTextTranslationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "translated_text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextTranslationResponse struct {
	value *TextTranslationResponse
	isSet bool
}

func (v NullableTextTranslationResponse) Get() *TextTranslationResponse {
	return v.value
}

func (v *NullableTextTranslationResponse) Set(val *TextTranslationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTextTranslationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTextTranslationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextTranslationResponse(val *TextTranslationResponse) *NullableTextTranslationResponse {
	return &NullableTextTranslationResponse{value: val, isSet: true}
}

func (v NullableTextTranslationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextTranslationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


