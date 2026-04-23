
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRacesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRacesResponse{}

// PalRacesResponse struct for PalRacesResponse
type PalRacesResponse struct {
	Races []RaceHistory `json:"races,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRacesResponse PalRacesResponse

// NewPalRacesResponse instantiates a new PalRacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRacesResponse() *PalRacesResponse {
	this := PalRacesResponse{}
	return &this
}

// NewPalRacesResponseWithDefaults instantiates a new PalRacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRacesResponseWithDefaults() *PalRacesResponse {
	this := PalRacesResponse{}
	return &this
}

// GetRaces returns the Races field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRacesResponse) GetRaces() []RaceHistory {
	if o == nil {
		var ret []RaceHistory
		return ret
	}
	return o.Races
}

// GetRacesOk returns a tuple with the Races field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRacesResponse) GetRacesOk() ([]RaceHistory, bool) {
	if o == nil || IsNil(o.Races) {
		return nil, false
	}
	return o.Races, true
}

// HasRaces returns a boolean if a field has been set.
func (o *PalRacesResponse) HasRaces() bool {
	if o != nil && !IsNil(o.Races) {
		return true
	}

	return false
}

// SetRaces gets a reference to the given []RaceHistory and assigns it to the Races field.
func (o *PalRacesResponse) SetRaces(v []RaceHistory) {
	o.Races = v
}

func (o PalRacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRacesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Races != nil {
		toSerialize["races"] = o.Races
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRacesResponse) UnmarshalJSON(data []byte) (err error) {
	varPalRacesResponse := _PalRacesResponse{}

	err = json.Unmarshal(data, &varPalRacesResponse)

	if err != nil {
		return err
	}

	*o = PalRacesResponse(varPalRacesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "races")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRacesResponse struct {
	value *PalRacesResponse
	isSet bool
}

func (v NullablePalRacesResponse) Get() *PalRacesResponse {
	return v.value
}

func (v *NullablePalRacesResponse) Set(val *PalRacesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRacesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRacesResponse(val *PalRacesResponse) *NullablePalRacesResponse {
	return &NullablePalRacesResponse{value: val, isSet: true}
}

func (v NullablePalRacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


