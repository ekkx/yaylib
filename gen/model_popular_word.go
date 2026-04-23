
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PopularWord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopularWord{}

// PopularWord struct for PopularWord
type PopularWord struct {
	Id NullableInt64 `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Word NullableString `json:"word,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopularWord PopularWord

// NewPopularWord instantiates a new PopularWord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopularWord() *PopularWord {
	this := PopularWord{}
	return &this
}

// NewPopularWordWithDefaults instantiates a new PopularWord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopularWordWithDefaults() *PopularWord {
	this := PopularWord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PopularWord) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PopularWord) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PopularWord) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *PopularWord) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PopularWord) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PopularWord) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PopularWord) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PopularWord) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PopularWord) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PopularWord) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PopularWord) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PopularWord) UnsetType() {
	o.Type.Unset()
}

// GetWord returns the Word field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PopularWord) GetWord() string {
	if o == nil || IsNil(o.Word.Get()) {
		var ret string
		return ret
	}
	return *o.Word.Get()
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PopularWord) GetWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Word.Get(), o.Word.IsSet()
}

// HasWord returns a boolean if a field has been set.
func (o *PopularWord) HasWord() bool {
	if o != nil && o.Word.IsSet() {
		return true
	}

	return false
}

// SetWord gets a reference to the given NullableString and assigns it to the Word field.
func (o *PopularWord) SetWord(v string) {
	o.Word.Set(&v)
}
// SetWordNil sets the value for Word to be an explicit nil
func (o *PopularWord) SetWordNil() {
	o.Word.Set(nil)
}

// UnsetWord ensures that no value is present for Word, not even an explicit nil
func (o *PopularWord) UnsetWord() {
	o.Word.Unset()
}

func (o PopularWord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopularWord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Word.IsSet() {
		toSerialize["word"] = o.Word.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PopularWord) UnmarshalJSON(data []byte) (err error) {
	varPopularWord := _PopularWord{}

	err = json.Unmarshal(data, &varPopularWord)

	if err != nil {
		return err
	}

	*o = PopularWord(varPopularWord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "word")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopularWord struct {
	value *PopularWord
	isSet bool
}

func (v NullablePopularWord) Get() *PopularWord {
	return v.value
}

func (v *NullablePopularWord) Set(val *PopularWord) {
	v.value = val
	v.isSet = true
}

func (v NullablePopularWord) IsSet() bool {
	return v.isSet
}

func (v *NullablePopularWord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopularWord(val *PopularWord) *NullablePopularWord {
	return &NullablePopularWord{value: val, isSet: true}
}

func (v NullablePopularWord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopularWord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


