
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SwapDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwapDTO{}

// SwapDTO struct for SwapDTO
type SwapDTO struct {
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwapDTO SwapDTO

// NewSwapDTO instantiates a new SwapDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwapDTO() *SwapDTO {
	this := SwapDTO{}
	return &this
}

// NewSwapDTOWithDefaults instantiates a new SwapDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwapDTOWithDefaults() *SwapDTO {
	this := SwapDTO{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwapDTO) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwapDTO) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *SwapDTO) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *SwapDTO) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *SwapDTO) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *SwapDTO) UnsetUrl() {
	o.Url.Unset()
}

func (o SwapDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwapDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwapDTO) UnmarshalJSON(data []byte) (err error) {
	varSwapDTO := _SwapDTO{}

	err = json.Unmarshal(data, &varSwapDTO)

	if err != nil {
		return err
	}

	*o = SwapDTO(varSwapDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwapDTO struct {
	value *SwapDTO
	isSet bool
}

func (v NullableSwapDTO) Get() *SwapDTO {
	return v.value
}

func (v *NullableSwapDTO) Set(val *SwapDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSwapDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSwapDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwapDTO(val *SwapDTO) *NullableSwapDTO {
	return &NullableSwapDTO{value: val, isSet: true}
}

func (v NullableSwapDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwapDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


