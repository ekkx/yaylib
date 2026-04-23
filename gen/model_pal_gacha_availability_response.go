
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalGachaAvailabilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalGachaAvailabilityResponse{}

// PalGachaAvailabilityResponse struct for PalGachaAvailabilityResponse
type PalGachaAvailabilityResponse struct {
	State NullableString `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalGachaAvailabilityResponse PalGachaAvailabilityResponse

// NewPalGachaAvailabilityResponse instantiates a new PalGachaAvailabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalGachaAvailabilityResponse() *PalGachaAvailabilityResponse {
	this := PalGachaAvailabilityResponse{}
	return &this
}

// NewPalGachaAvailabilityResponseWithDefaults instantiates a new PalGachaAvailabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalGachaAvailabilityResponseWithDefaults() *PalGachaAvailabilityResponse {
	this := PalGachaAvailabilityResponse{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaAvailabilityResponse) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaAvailabilityResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *PalGachaAvailabilityResponse) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *PalGachaAvailabilityResponse) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *PalGachaAvailabilityResponse) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *PalGachaAvailabilityResponse) UnsetState() {
	o.State.Unset()
}

func (o PalGachaAvailabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalGachaAvailabilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalGachaAvailabilityResponse) UnmarshalJSON(data []byte) (err error) {
	varPalGachaAvailabilityResponse := _PalGachaAvailabilityResponse{}

	err = json.Unmarshal(data, &varPalGachaAvailabilityResponse)

	if err != nil {
		return err
	}

	*o = PalGachaAvailabilityResponse(varPalGachaAvailabilityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalGachaAvailabilityResponse struct {
	value *PalGachaAvailabilityResponse
	isSet bool
}

func (v NullablePalGachaAvailabilityResponse) Get() *PalGachaAvailabilityResponse {
	return v.value
}

func (v *NullablePalGachaAvailabilityResponse) Set(val *PalGachaAvailabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGachaAvailabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGachaAvailabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGachaAvailabilityResponse(val *PalGachaAvailabilityResponse) *NullablePalGachaAvailabilityResponse {
	return &NullablePalGachaAvailabilityResponse{value: val, isSet: true}
}

func (v NullablePalGachaAvailabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGachaAvailabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


