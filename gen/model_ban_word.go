
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BanWord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BanWord{}

// BanWord struct for BanWord
type BanWord struct {
	Id NullableInt64 `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Word NullableString `json:"word,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BanWord BanWord

// NewBanWord instantiates a new BanWord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanWord() *BanWord {
	this := BanWord{}
	return &this
}

// NewBanWordWithDefaults instantiates a new BanWord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanWordWithDefaults() *BanWord {
	this := BanWord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanWord) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanWord) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *BanWord) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *BanWord) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *BanWord) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *BanWord) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanWord) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanWord) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *BanWord) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *BanWord) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *BanWord) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *BanWord) UnsetType() {
	o.Type.Unset()
}

// GetWord returns the Word field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanWord) GetWord() string {
	if o == nil || IsNil(o.Word.Get()) {
		var ret string
		return ret
	}
	return *o.Word.Get()
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanWord) GetWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Word.Get(), o.Word.IsSet()
}

// HasWord returns a boolean if a field has been set.
func (o *BanWord) HasWord() bool {
	if o != nil && o.Word.IsSet() {
		return true
	}

	return false
}

// SetWord gets a reference to the given NullableString and assigns it to the Word field.
func (o *BanWord) SetWord(v string) {
	o.Word.Set(&v)
}
// SetWordNil sets the value for Word to be an explicit nil
func (o *BanWord) SetWordNil() {
	o.Word.Set(nil)
}

// UnsetWord ensures that no value is present for Word, not even an explicit nil
func (o *BanWord) UnsetWord() {
	o.Word.Unset()
}

func (o BanWord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BanWord) ToMap() (map[string]interface{}, error) {
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

func (o *BanWord) UnmarshalJSON(data []byte) (err error) {
	varBanWord := _BanWord{}

	err = json.Unmarshal(data, &varBanWord)

	if err != nil {
		return err
	}

	*o = BanWord(varBanWord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "word")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBanWord struct {
	value *BanWord
	isSet bool
}

func (v NullableBanWord) Get() *BanWord {
	return v.value
}

func (v *NullableBanWord) Set(val *BanWord) {
	v.value = val
	v.isSet = true
}

func (v NullableBanWord) IsSet() bool {
	return v.isSet
}

func (v *NullableBanWord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanWord(val *BanWord) *NullableBanWord {
	return &NullableBanWord{value: val, isSet: true}
}

func (v NullableBanWord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanWord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


