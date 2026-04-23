
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Dead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dead{}

// Dead struct for Dead
type Dead struct {
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Dead Dead

// NewDead instantiates a new Dead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDead() *Dead {
	this := Dead{}
	return &this
}

// NewDeadWithDefaults instantiates a new Dead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeadWithDefaults() *Dead {
	this := Dead{}
	return &this
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dead) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dead) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *Dead) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *Dead) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *Dead) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *Dead) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dead) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dead) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *Dead) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *Dead) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *Dead) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *Dead) UnsetPalName() {
	o.PalName.Unset()
}

func (o Dead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Dead) UnmarshalJSON(data []byte) (err error) {
	varDead := _Dead{}

	err = json.Unmarshal(data, &varDead)

	if err != nil {
		return err
	}

	*o = Dead(varDead)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDead struct {
	value *Dead
	isSet bool
}

func (v NullableDead) Get() *Dead {
	return v.value
}

func (v *NullableDead) Set(val *Dead) {
	v.value = val
	v.isSet = true
}

func (v NullableDead) IsSet() bool {
	return v.isSet
}

func (v *NullableDead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDead(val *Dead) *NullableDead {
	return &NullableDead{value: val, isSet: true}
}

func (v NullableDead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


