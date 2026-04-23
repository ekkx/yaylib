
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignMissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMissionsResponse{}

// CampaignMissionsResponse struct for CampaignMissionsResponse
type CampaignMissionsResponse struct {
	DailyMissions []MissionDTO `json:"daily_missions,omitempty"`
	MultiplierMissions []MissionDTO `json:"multiplier_missions,omitempty"`
	NormalMissions []MissionDTO `json:"normal_missions,omitempty"`
	UnlimitedMissions []MissionDTO `json:"unlimited_missions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignMissionsResponse CampaignMissionsResponse

// NewCampaignMissionsResponse instantiates a new CampaignMissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMissionsResponse() *CampaignMissionsResponse {
	this := CampaignMissionsResponse{}
	return &this
}

// NewCampaignMissionsResponseWithDefaults instantiates a new CampaignMissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMissionsResponseWithDefaults() *CampaignMissionsResponse {
	this := CampaignMissionsResponse{}
	return &this
}

// GetDailyMissions returns the DailyMissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMissionsResponse) GetDailyMissions() []MissionDTO {
	if o == nil {
		var ret []MissionDTO
		return ret
	}
	return o.DailyMissions
}

// GetDailyMissionsOk returns a tuple with the DailyMissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMissionsResponse) GetDailyMissionsOk() ([]MissionDTO, bool) {
	if o == nil || IsNil(o.DailyMissions) {
		return nil, false
	}
	return o.DailyMissions, true
}

// HasDailyMissions returns a boolean if a field has been set.
func (o *CampaignMissionsResponse) HasDailyMissions() bool {
	if o != nil && !IsNil(o.DailyMissions) {
		return true
	}

	return false
}

// SetDailyMissions gets a reference to the given []MissionDTO and assigns it to the DailyMissions field.
func (o *CampaignMissionsResponse) SetDailyMissions(v []MissionDTO) {
	o.DailyMissions = v
}

// GetMultiplierMissions returns the MultiplierMissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMissionsResponse) GetMultiplierMissions() []MissionDTO {
	if o == nil {
		var ret []MissionDTO
		return ret
	}
	return o.MultiplierMissions
}

// GetMultiplierMissionsOk returns a tuple with the MultiplierMissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMissionsResponse) GetMultiplierMissionsOk() ([]MissionDTO, bool) {
	if o == nil || IsNil(o.MultiplierMissions) {
		return nil, false
	}
	return o.MultiplierMissions, true
}

// HasMultiplierMissions returns a boolean if a field has been set.
func (o *CampaignMissionsResponse) HasMultiplierMissions() bool {
	if o != nil && !IsNil(o.MultiplierMissions) {
		return true
	}

	return false
}

// SetMultiplierMissions gets a reference to the given []MissionDTO and assigns it to the MultiplierMissions field.
func (o *CampaignMissionsResponse) SetMultiplierMissions(v []MissionDTO) {
	o.MultiplierMissions = v
}

// GetNormalMissions returns the NormalMissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMissionsResponse) GetNormalMissions() []MissionDTO {
	if o == nil {
		var ret []MissionDTO
		return ret
	}
	return o.NormalMissions
}

// GetNormalMissionsOk returns a tuple with the NormalMissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMissionsResponse) GetNormalMissionsOk() ([]MissionDTO, bool) {
	if o == nil || IsNil(o.NormalMissions) {
		return nil, false
	}
	return o.NormalMissions, true
}

// HasNormalMissions returns a boolean if a field has been set.
func (o *CampaignMissionsResponse) HasNormalMissions() bool {
	if o != nil && !IsNil(o.NormalMissions) {
		return true
	}

	return false
}

// SetNormalMissions gets a reference to the given []MissionDTO and assigns it to the NormalMissions field.
func (o *CampaignMissionsResponse) SetNormalMissions(v []MissionDTO) {
	o.NormalMissions = v
}

// GetUnlimitedMissions returns the UnlimitedMissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMissionsResponse) GetUnlimitedMissions() []MissionDTO {
	if o == nil {
		var ret []MissionDTO
		return ret
	}
	return o.UnlimitedMissions
}

// GetUnlimitedMissionsOk returns a tuple with the UnlimitedMissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMissionsResponse) GetUnlimitedMissionsOk() ([]MissionDTO, bool) {
	if o == nil || IsNil(o.UnlimitedMissions) {
		return nil, false
	}
	return o.UnlimitedMissions, true
}

// HasUnlimitedMissions returns a boolean if a field has been set.
func (o *CampaignMissionsResponse) HasUnlimitedMissions() bool {
	if o != nil && !IsNil(o.UnlimitedMissions) {
		return true
	}

	return false
}

// SetUnlimitedMissions gets a reference to the given []MissionDTO and assigns it to the UnlimitedMissions field.
func (o *CampaignMissionsResponse) SetUnlimitedMissions(v []MissionDTO) {
	o.UnlimitedMissions = v
}

func (o CampaignMissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignMissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DailyMissions != nil {
		toSerialize["daily_missions"] = o.DailyMissions
	}
	if o.MultiplierMissions != nil {
		toSerialize["multiplier_missions"] = o.MultiplierMissions
	}
	if o.NormalMissions != nil {
		toSerialize["normal_missions"] = o.NormalMissions
	}
	if o.UnlimitedMissions != nil {
		toSerialize["unlimited_missions"] = o.UnlimitedMissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignMissionsResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignMissionsResponse := _CampaignMissionsResponse{}

	err = json.Unmarshal(data, &varCampaignMissionsResponse)

	if err != nil {
		return err
	}

	*o = CampaignMissionsResponse(varCampaignMissionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "daily_missions")
		delete(additionalProperties, "multiplier_missions")
		delete(additionalProperties, "normal_missions")
		delete(additionalProperties, "unlimited_missions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignMissionsResponse struct {
	value *CampaignMissionsResponse
	isSet bool
}

func (v NullableCampaignMissionsResponse) Get() *CampaignMissionsResponse {
	return v.value
}

func (v *NullableCampaignMissionsResponse) Set(val *CampaignMissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMissionsResponse(val *CampaignMissionsResponse) *NullableCampaignMissionsResponse {
	return &NullableCampaignMissionsResponse{value: val, isSet: true}
}

func (v NullableCampaignMissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


