
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceDetailsResponse{}

// PalRaceDetailsResponse struct for PalRaceDetailsResponse
type PalRaceDetailsResponse struct {
	Race NullableRaceHistory `json:"race,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceDetailsResponse PalRaceDetailsResponse

// NewPalRaceDetailsResponse instantiates a new PalRaceDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceDetailsResponse() *PalRaceDetailsResponse {
	this := PalRaceDetailsResponse{}
	return &this
}

// NewPalRaceDetailsResponseWithDefaults instantiates a new PalRaceDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceDetailsResponseWithDefaults() *PalRaceDetailsResponse {
	this := PalRaceDetailsResponse{}
	return &this
}

// GetRace returns the Race field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceDetailsResponse) GetRace() RaceHistory {
	if o == nil || IsNil(o.Race.Get()) {
		var ret RaceHistory
		return ret
	}
	return *o.Race.Get()
}

// GetRaceOk returns a tuple with the Race field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceDetailsResponse) GetRaceOk() (*RaceHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Race.Get(), o.Race.IsSet()
}

// HasRace returns a boolean if a field has been set.
func (o *PalRaceDetailsResponse) HasRace() bool {
	if o != nil && o.Race.IsSet() {
		return true
	}

	return false
}

// SetRace gets a reference to the given NullableRaceHistory and assigns it to the Race field.
func (o *PalRaceDetailsResponse) SetRace(v RaceHistory) {
	o.Race.Set(&v)
}
// SetRaceNil sets the value for Race to be an explicit nil
func (o *PalRaceDetailsResponse) SetRaceNil() {
	o.Race.Set(nil)
}

// UnsetRace ensures that no value is present for Race, not even an explicit nil
func (o *PalRaceDetailsResponse) UnsetRace() {
	o.Race.Unset()
}

func (o PalRaceDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Race.IsSet() {
		toSerialize["race"] = o.Race.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varPalRaceDetailsResponse := _PalRaceDetailsResponse{}

	err = json.Unmarshal(data, &varPalRaceDetailsResponse)

	if err != nil {
		return err
	}

	*o = PalRaceDetailsResponse(varPalRaceDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "race")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceDetailsResponse struct {
	value *PalRaceDetailsResponse
	isSet bool
}

func (v NullablePalRaceDetailsResponse) Get() *PalRaceDetailsResponse {
	return v.value
}

func (v *NullablePalRaceDetailsResponse) Set(val *PalRaceDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceDetailsResponse(val *PalRaceDetailsResponse) *NullablePalRaceDetailsResponse {
	return &NullablePalRaceDetailsResponse{value: val, isSet: true}
}

func (v NullablePalRaceDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


