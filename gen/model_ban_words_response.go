
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BanWordsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BanWordsResponse{}

// BanWordsResponse struct for BanWordsResponse
type BanWordsResponse struct {
	BanWords []BanWord `json:"ban_words,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BanWordsResponse BanWordsResponse

// NewBanWordsResponse instantiates a new BanWordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanWordsResponse() *BanWordsResponse {
	this := BanWordsResponse{}
	return &this
}

// NewBanWordsResponseWithDefaults instantiates a new BanWordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanWordsResponseWithDefaults() *BanWordsResponse {
	this := BanWordsResponse{}
	return &this
}

// GetBanWords returns the BanWords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanWordsResponse) GetBanWords() []BanWord {
	if o == nil {
		var ret []BanWord
		return ret
	}
	return o.BanWords
}

// GetBanWordsOk returns a tuple with the BanWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanWordsResponse) GetBanWordsOk() ([]BanWord, bool) {
	if o == nil || IsNil(o.BanWords) {
		return nil, false
	}
	return o.BanWords, true
}

// HasBanWords returns a boolean if a field has been set.
func (o *BanWordsResponse) HasBanWords() bool {
	if o != nil && !IsNil(o.BanWords) {
		return true
	}

	return false
}

// SetBanWords gets a reference to the given []BanWord and assigns it to the BanWords field.
func (o *BanWordsResponse) SetBanWords(v []BanWord) {
	o.BanWords = v
}

func (o BanWordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BanWordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BanWords != nil {
		toSerialize["ban_words"] = o.BanWords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BanWordsResponse) UnmarshalJSON(data []byte) (err error) {
	varBanWordsResponse := _BanWordsResponse{}

	err = json.Unmarshal(data, &varBanWordsResponse)

	if err != nil {
		return err
	}

	*o = BanWordsResponse(varBanWordsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ban_words")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBanWordsResponse struct {
	value *BanWordsResponse
	isSet bool
}

func (v NullableBanWordsResponse) Get() *BanWordsResponse {
	return v.value
}

func (v *NullableBanWordsResponse) Set(val *BanWordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBanWordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBanWordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanWordsResponse(val *BanWordsResponse) *NullableBanWordsResponse {
	return &NullableBanWordsResponse{value: val, isSet: true}
}

func (v NullableBanWordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanWordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


