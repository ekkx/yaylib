
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MuteKeywordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteKeywordRequest{}

// MuteKeywordRequest struct for MuteKeywordRequest
type MuteKeywordRequest struct {
	Context []string `json:"context,omitempty"`
	Word NullableString `json:"word,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MuteKeywordRequest MuteKeywordRequest

// NewMuteKeywordRequest instantiates a new MuteKeywordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteKeywordRequest() *MuteKeywordRequest {
	this := MuteKeywordRequest{}
	return &this
}

// NewMuteKeywordRequestWithDefaults instantiates a new MuteKeywordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteKeywordRequestWithDefaults() *MuteKeywordRequest {
	this := MuteKeywordRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeywordRequest) GetContext() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeywordRequest) GetContextOk() ([]string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MuteKeywordRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []string and assigns it to the Context field.
func (o *MuteKeywordRequest) SetContext(v []string) {
	o.Context = v
}

// GetWord returns the Word field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeywordRequest) GetWord() string {
	if o == nil || IsNil(o.Word.Get()) {
		var ret string
		return ret
	}
	return *o.Word.Get()
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeywordRequest) GetWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Word.Get(), o.Word.IsSet()
}

// HasWord returns a boolean if a field has been set.
func (o *MuteKeywordRequest) HasWord() bool {
	if o != nil && o.Word.IsSet() {
		return true
	}

	return false
}

// SetWord gets a reference to the given NullableString and assigns it to the Word field.
func (o *MuteKeywordRequest) SetWord(v string) {
	o.Word.Set(&v)
}
// SetWordNil sets the value for Word to be an explicit nil
func (o *MuteKeywordRequest) SetWordNil() {
	o.Word.Set(nil)
}

// UnsetWord ensures that no value is present for Word, not even an explicit nil
func (o *MuteKeywordRequest) UnsetWord() {
	o.Word.Unset()
}

func (o MuteKeywordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteKeywordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Word.IsSet() {
		toSerialize["word"] = o.Word.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MuteKeywordRequest) UnmarshalJSON(data []byte) (err error) {
	varMuteKeywordRequest := _MuteKeywordRequest{}

	err = json.Unmarshal(data, &varMuteKeywordRequest)

	if err != nil {
		return err
	}

	*o = MuteKeywordRequest(varMuteKeywordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "word")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMuteKeywordRequest struct {
	value *MuteKeywordRequest
	isSet bool
}

func (v NullableMuteKeywordRequest) Get() *MuteKeywordRequest {
	return v.value
}

func (v *NullableMuteKeywordRequest) Set(val *MuteKeywordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteKeywordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteKeywordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteKeywordRequest(val *MuteKeywordRequest) *NullableMuteKeywordRequest {
	return &NullableMuteKeywordRequest{value: val, isSet: true}
}

func (v NullableMuteKeywordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteKeywordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


