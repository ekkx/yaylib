
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LevelUpDetailsPal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LevelUpDetailsPal{}

// LevelUpDetailsPal struct for LevelUpDetailsPal
type LevelUpDetailsPal struct {
	Image NullableString `json:"image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LevelUpDetailsPal LevelUpDetailsPal

// NewLevelUpDetailsPal instantiates a new LevelUpDetailsPal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevelUpDetailsPal() *LevelUpDetailsPal {
	this := LevelUpDetailsPal{}
	return &this
}

// NewLevelUpDetailsPalWithDefaults instantiates a new LevelUpDetailsPal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelUpDetailsPalWithDefaults() *LevelUpDetailsPal {
	this := LevelUpDetailsPal{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetailsPal) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetailsPal) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *LevelUpDetailsPal) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *LevelUpDetailsPal) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *LevelUpDetailsPal) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *LevelUpDetailsPal) UnsetImage() {
	o.Image.Unset()
}

func (o LevelUpDetailsPal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LevelUpDetailsPal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LevelUpDetailsPal) UnmarshalJSON(data []byte) (err error) {
	varLevelUpDetailsPal := _LevelUpDetailsPal{}

	err = json.Unmarshal(data, &varLevelUpDetailsPal)

	if err != nil {
		return err
	}

	*o = LevelUpDetailsPal(varLevelUpDetailsPal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLevelUpDetailsPal struct {
	value *LevelUpDetailsPal
	isSet bool
}

func (v NullableLevelUpDetailsPal) Get() *LevelUpDetailsPal {
	return v.value
}

func (v *NullableLevelUpDetailsPal) Set(val *LevelUpDetailsPal) {
	v.value = val
	v.isSet = true
}

func (v NullableLevelUpDetailsPal) IsSet() bool {
	return v.isSet
}

func (v *NullableLevelUpDetailsPal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevelUpDetailsPal(val *LevelUpDetailsPal) *NullableLevelUpDetailsPal {
	return &NullableLevelUpDetailsPal{value: val, isSet: true}
}

func (v NullableLevelUpDetailsPal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevelUpDetailsPal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


