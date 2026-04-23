
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceAvailablePalListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceAvailablePalListResponse{}

// PalRaceAvailablePalListResponse struct for PalRaceAvailablePalListResponse
type PalRaceAvailablePalListResponse struct {
	Pals []PalRaceItemDTO `json:"pals,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceAvailablePalListResponse PalRaceAvailablePalListResponse

// NewPalRaceAvailablePalListResponse instantiates a new PalRaceAvailablePalListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceAvailablePalListResponse() *PalRaceAvailablePalListResponse {
	this := PalRaceAvailablePalListResponse{}
	return &this
}

// NewPalRaceAvailablePalListResponseWithDefaults instantiates a new PalRaceAvailablePalListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceAvailablePalListResponseWithDefaults() *PalRaceAvailablePalListResponse {
	this := PalRaceAvailablePalListResponse{}
	return &this
}

// GetPals returns the Pals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceAvailablePalListResponse) GetPals() []PalRaceItemDTO {
	if o == nil {
		var ret []PalRaceItemDTO
		return ret
	}
	return o.Pals
}

// GetPalsOk returns a tuple with the Pals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceAvailablePalListResponse) GetPalsOk() ([]PalRaceItemDTO, bool) {
	if o == nil || IsNil(o.Pals) {
		return nil, false
	}
	return o.Pals, true
}

// HasPals returns a boolean if a field has been set.
func (o *PalRaceAvailablePalListResponse) HasPals() bool {
	if o != nil && !IsNil(o.Pals) {
		return true
	}

	return false
}

// SetPals gets a reference to the given []PalRaceItemDTO and assigns it to the Pals field.
func (o *PalRaceAvailablePalListResponse) SetPals(v []PalRaceItemDTO) {
	o.Pals = v
}

func (o PalRaceAvailablePalListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceAvailablePalListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Pals != nil {
		toSerialize["pals"] = o.Pals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceAvailablePalListResponse) UnmarshalJSON(data []byte) (err error) {
	varPalRaceAvailablePalListResponse := _PalRaceAvailablePalListResponse{}

	err = json.Unmarshal(data, &varPalRaceAvailablePalListResponse)

	if err != nil {
		return err
	}

	*o = PalRaceAvailablePalListResponse(varPalRaceAvailablePalListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceAvailablePalListResponse struct {
	value *PalRaceAvailablePalListResponse
	isSet bool
}

func (v NullablePalRaceAvailablePalListResponse) Get() *PalRaceAvailablePalListResponse {
	return v.value
}

func (v *NullablePalRaceAvailablePalListResponse) Set(val *PalRaceAvailablePalListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceAvailablePalListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceAvailablePalListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceAvailablePalListResponse(val *PalRaceAvailablePalListResponse) *NullablePalRaceAvailablePalListResponse {
	return &NullablePalRaceAvailablePalListResponse{value: val, isSet: true}
}

func (v NullablePalRaceAvailablePalListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceAvailablePalListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


