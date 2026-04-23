
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ProfileImageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileImageDTO{}

// ProfileImageDTO struct for ProfileImageDTO
type ProfileImageDTO struct {
	Image NullableString `json:"image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileImageDTO ProfileImageDTO

// NewProfileImageDTO instantiates a new ProfileImageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileImageDTO() *ProfileImageDTO {
	this := ProfileImageDTO{}
	return &this
}

// NewProfileImageDTOWithDefaults instantiates a new ProfileImageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileImageDTOWithDefaults() *ProfileImageDTO {
	this := ProfileImageDTO{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileImageDTO) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileImageDTO) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ProfileImageDTO) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ProfileImageDTO) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ProfileImageDTO) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ProfileImageDTO) UnsetImage() {
	o.Image.Unset()
}

func (o ProfileImageDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileImageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileImageDTO) UnmarshalJSON(data []byte) (err error) {
	varProfileImageDTO := _ProfileImageDTO{}

	err = json.Unmarshal(data, &varProfileImageDTO)

	if err != nil {
		return err
	}

	*o = ProfileImageDTO(varProfileImageDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileImageDTO struct {
	value *ProfileImageDTO
	isSet bool
}

func (v NullableProfileImageDTO) Get() *ProfileImageDTO {
	return v.value
}

func (v *NullableProfileImageDTO) Set(val *ProfileImageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileImageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileImageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileImageDTO(val *ProfileImageDTO) *NullableProfileImageDTO {
	return &NullableProfileImageDTO{value: val, isSet: true}
}

func (v NullableProfileImageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileImageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


