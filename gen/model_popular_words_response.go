
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PopularWordsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopularWordsResponse{}

// PopularWordsResponse struct for PopularWordsResponse
type PopularWordsResponse struct {
	PopularWords []PopularWord `json:"popular_words,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopularWordsResponse PopularWordsResponse

// NewPopularWordsResponse instantiates a new PopularWordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopularWordsResponse() *PopularWordsResponse {
	this := PopularWordsResponse{}
	return &this
}

// NewPopularWordsResponseWithDefaults instantiates a new PopularWordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopularWordsResponseWithDefaults() *PopularWordsResponse {
	this := PopularWordsResponse{}
	return &this
}

// GetPopularWords returns the PopularWords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PopularWordsResponse) GetPopularWords() []PopularWord {
	if o == nil {
		var ret []PopularWord
		return ret
	}
	return o.PopularWords
}

// GetPopularWordsOk returns a tuple with the PopularWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PopularWordsResponse) GetPopularWordsOk() ([]PopularWord, bool) {
	if o == nil || IsNil(o.PopularWords) {
		return nil, false
	}
	return o.PopularWords, true
}

// HasPopularWords returns a boolean if a field has been set.
func (o *PopularWordsResponse) HasPopularWords() bool {
	if o != nil && !IsNil(o.PopularWords) {
		return true
	}

	return false
}

// SetPopularWords gets a reference to the given []PopularWord and assigns it to the PopularWords field.
func (o *PopularWordsResponse) SetPopularWords(v []PopularWord) {
	o.PopularWords = v
}

func (o PopularWordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopularWordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PopularWords != nil {
		toSerialize["popular_words"] = o.PopularWords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PopularWordsResponse) UnmarshalJSON(data []byte) (err error) {
	varPopularWordsResponse := _PopularWordsResponse{}

	err = json.Unmarshal(data, &varPopularWordsResponse)

	if err != nil {
		return err
	}

	*o = PopularWordsResponse(varPopularWordsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "popular_words")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopularWordsResponse struct {
	value *PopularWordsResponse
	isSet bool
}

func (v NullablePopularWordsResponse) Get() *PopularWordsResponse {
	return v.value
}

func (v *NullablePopularWordsResponse) Set(val *PopularWordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePopularWordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePopularWordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopularWordsResponse(val *PopularWordsResponse) *NullablePopularWordsResponse {
	return &NullablePopularWordsResponse{value: val, isSet: true}
}

func (v NullablePopularWordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopularWordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


