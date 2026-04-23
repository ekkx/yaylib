
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceRegisterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceRegisterResponse{}

// PalRaceRegisterResponse struct for PalRaceRegisterResponse
type PalRaceRegisterResponse struct {
	RaceHistory NullableRaceHistory `json:"race_history,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceRegisterResponse PalRaceRegisterResponse

// NewPalRaceRegisterResponse instantiates a new PalRaceRegisterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceRegisterResponse() *PalRaceRegisterResponse {
	this := PalRaceRegisterResponse{}
	return &this
}

// NewPalRaceRegisterResponseWithDefaults instantiates a new PalRaceRegisterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceRegisterResponseWithDefaults() *PalRaceRegisterResponse {
	this := PalRaceRegisterResponse{}
	return &this
}

// GetRaceHistory returns the RaceHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceRegisterResponse) GetRaceHistory() RaceHistory {
	if o == nil || IsNil(o.RaceHistory.Get()) {
		var ret RaceHistory
		return ret
	}
	return *o.RaceHistory.Get()
}

// GetRaceHistoryOk returns a tuple with the RaceHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceRegisterResponse) GetRaceHistoryOk() (*RaceHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceHistory.Get(), o.RaceHistory.IsSet()
}

// HasRaceHistory returns a boolean if a field has been set.
func (o *PalRaceRegisterResponse) HasRaceHistory() bool {
	if o != nil && o.RaceHistory.IsSet() {
		return true
	}

	return false
}

// SetRaceHistory gets a reference to the given NullableRaceHistory and assigns it to the RaceHistory field.
func (o *PalRaceRegisterResponse) SetRaceHistory(v RaceHistory) {
	o.RaceHistory.Set(&v)
}
// SetRaceHistoryNil sets the value for RaceHistory to be an explicit nil
func (o *PalRaceRegisterResponse) SetRaceHistoryNil() {
	o.RaceHistory.Set(nil)
}

// UnsetRaceHistory ensures that no value is present for RaceHistory, not even an explicit nil
func (o *PalRaceRegisterResponse) UnsetRaceHistory() {
	o.RaceHistory.Unset()
}

func (o PalRaceRegisterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceRegisterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RaceHistory.IsSet() {
		toSerialize["race_history"] = o.RaceHistory.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceRegisterResponse) UnmarshalJSON(data []byte) (err error) {
	varPalRaceRegisterResponse := _PalRaceRegisterResponse{}

	err = json.Unmarshal(data, &varPalRaceRegisterResponse)

	if err != nil {
		return err
	}

	*o = PalRaceRegisterResponse(varPalRaceRegisterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "race_history")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceRegisterResponse struct {
	value *PalRaceRegisterResponse
	isSet bool
}

func (v NullablePalRaceRegisterResponse) Get() *PalRaceRegisterResponse {
	return v.value
}

func (v *NullablePalRaceRegisterResponse) Set(val *PalRaceRegisterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceRegisterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceRegisterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceRegisterResponse(val *PalRaceRegisterResponse) *NullablePalRaceRegisterResponse {
	return &NullablePalRaceRegisterResponse{value: val, isSet: true}
}

func (v NullablePalRaceRegisterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceRegisterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


