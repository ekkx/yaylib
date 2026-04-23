
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MuteKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteKeyword{}

// MuteKeyword struct for MuteKeyword
type MuteKeyword struct {
	// Unresolved Java type: io.realm.u0
	Context map[string]interface{} `json:"context,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Word NullableString `json:"word,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MuteKeyword MuteKeyword

// NewMuteKeyword instantiates a new MuteKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteKeyword() *MuteKeyword {
	this := MuteKeyword{}
	return &this
}

// NewMuteKeywordWithDefaults instantiates a new MuteKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteKeywordWithDefaults() *MuteKeyword {
	this := MuteKeyword{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeyword) GetContext() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeyword) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MuteKeyword) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *MuteKeyword) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeyword) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeyword) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MuteKeyword) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *MuteKeyword) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MuteKeyword) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MuteKeyword) UnsetId() {
	o.Id.Unset()
}

// GetWord returns the Word field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MuteKeyword) GetWord() string {
	if o == nil || IsNil(o.Word.Get()) {
		var ret string
		return ret
	}
	return *o.Word.Get()
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MuteKeyword) GetWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Word.Get(), o.Word.IsSet()
}

// HasWord returns a boolean if a field has been set.
func (o *MuteKeyword) HasWord() bool {
	if o != nil && o.Word.IsSet() {
		return true
	}

	return false
}

// SetWord gets a reference to the given NullableString and assigns it to the Word field.
func (o *MuteKeyword) SetWord(v string) {
	o.Word.Set(&v)
}
// SetWordNil sets the value for Word to be an explicit nil
func (o *MuteKeyword) SetWordNil() {
	o.Word.Set(nil)
}

// UnsetWord ensures that no value is present for Word, not even an explicit nil
func (o *MuteKeyword) UnsetWord() {
	o.Word.Unset()
}

func (o MuteKeyword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Word.IsSet() {
		toSerialize["word"] = o.Word.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MuteKeyword) UnmarshalJSON(data []byte) (err error) {
	varMuteKeyword := _MuteKeyword{}

	err = json.Unmarshal(data, &varMuteKeyword)

	if err != nil {
		return err
	}

	*o = MuteKeyword(varMuteKeyword)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "id")
		delete(additionalProperties, "word")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMuteKeyword struct {
	value *MuteKeyword
	isSet bool
}

func (v NullableMuteKeyword) Get() *MuteKeyword {
	return v.value
}

func (v *NullableMuteKeyword) Set(val *MuteKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteKeyword(val *MuteKeyword) *NullableMuteKeyword {
	return &NullableMuteKeyword{value: val, isSet: true}
}

func (v NullableMuteKeyword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


