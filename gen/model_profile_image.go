
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ProfileImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileImage{}

// ProfileImage struct for ProfileImage
type ProfileImage struct {
	Image NullableString `json:"image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileImage ProfileImage

// NewProfileImage instantiates a new ProfileImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileImage() *ProfileImage {
	this := ProfileImage{}
	return &this
}

// NewProfileImageWithDefaults instantiates a new ProfileImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileImageWithDefaults() *ProfileImage {
	this := ProfileImage{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileImage) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileImage) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ProfileImage) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ProfileImage) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ProfileImage) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ProfileImage) UnsetImage() {
	o.Image.Unset()
}

func (o ProfileImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileImage) UnmarshalJSON(data []byte) (err error) {
	varProfileImage := _ProfileImage{}

	err = json.Unmarshal(data, &varProfileImage)

	if err != nil {
		return err
	}

	*o = ProfileImage(varProfileImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileImage struct {
	value *ProfileImage
	isSet bool
}

func (v NullableProfileImage) Get() *ProfileImage {
	return v.value
}

func (v *NullableProfileImage) Set(val *ProfileImage) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileImage) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileImage(val *ProfileImage) *NullableProfileImage {
	return &NullableProfileImage{value: val, isSet: true}
}

func (v NullableProfileImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


