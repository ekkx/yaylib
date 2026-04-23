
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceLeaguesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceLeaguesResponse{}

// PalRaceLeaguesResponse struct for PalRaceLeaguesResponse
type PalRaceLeaguesResponse struct {
	Leagues []PalRaceLeagueDTO `json:"leagues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceLeaguesResponse PalRaceLeaguesResponse

// NewPalRaceLeaguesResponse instantiates a new PalRaceLeaguesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceLeaguesResponse() *PalRaceLeaguesResponse {
	this := PalRaceLeaguesResponse{}
	return &this
}

// NewPalRaceLeaguesResponseWithDefaults instantiates a new PalRaceLeaguesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceLeaguesResponseWithDefaults() *PalRaceLeaguesResponse {
	this := PalRaceLeaguesResponse{}
	return &this
}

// GetLeagues returns the Leagues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeaguesResponse) GetLeagues() []PalRaceLeagueDTO {
	if o == nil {
		var ret []PalRaceLeagueDTO
		return ret
	}
	return o.Leagues
}

// GetLeaguesOk returns a tuple with the Leagues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeaguesResponse) GetLeaguesOk() ([]PalRaceLeagueDTO, bool) {
	if o == nil || IsNil(o.Leagues) {
		return nil, false
	}
	return o.Leagues, true
}

// HasLeagues returns a boolean if a field has been set.
func (o *PalRaceLeaguesResponse) HasLeagues() bool {
	if o != nil && !IsNil(o.Leagues) {
		return true
	}

	return false
}

// SetLeagues gets a reference to the given []PalRaceLeagueDTO and assigns it to the Leagues field.
func (o *PalRaceLeaguesResponse) SetLeagues(v []PalRaceLeagueDTO) {
	o.Leagues = v
}

func (o PalRaceLeaguesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceLeaguesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Leagues != nil {
		toSerialize["leagues"] = o.Leagues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceLeaguesResponse) UnmarshalJSON(data []byte) (err error) {
	varPalRaceLeaguesResponse := _PalRaceLeaguesResponse{}

	err = json.Unmarshal(data, &varPalRaceLeaguesResponse)

	if err != nil {
		return err
	}

	*o = PalRaceLeaguesResponse(varPalRaceLeaguesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leagues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceLeaguesResponse struct {
	value *PalRaceLeaguesResponse
	isSet bool
}

func (v NullablePalRaceLeaguesResponse) Get() *PalRaceLeaguesResponse {
	return v.value
}

func (v *NullablePalRaceLeaguesResponse) Set(val *PalRaceLeaguesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceLeaguesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceLeaguesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceLeaguesResponse(val *PalRaceLeaguesResponse) *NullablePalRaceLeaguesResponse {
	return &NullablePalRaceLeaguesResponse{value: val, isSet: true}
}

func (v NullablePalRaceLeaguesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceLeaguesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


