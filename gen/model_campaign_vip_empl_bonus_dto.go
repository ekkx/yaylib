
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignVipEmplBonusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignVipEmplBonusDTO{}

// CampaignVipEmplBonusDTO struct for CampaignVipEmplBonusDTO
type CampaignVipEmplBonusDTO struct {
	Claimed NullableBool `json:"claimed,omitempty"`
	Details NullableDetailsDTO `json:"details,omitempty"`
	EndAt NullableInt64 `json:"end_at,omitempty"`
	Key NullableString `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignVipEmplBonusDTO CampaignVipEmplBonusDTO

// NewCampaignVipEmplBonusDTO instantiates a new CampaignVipEmplBonusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignVipEmplBonusDTO() *CampaignVipEmplBonusDTO {
	this := CampaignVipEmplBonusDTO{}
	return &this
}

// NewCampaignVipEmplBonusDTOWithDefaults instantiates a new CampaignVipEmplBonusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignVipEmplBonusDTOWithDefaults() *CampaignVipEmplBonusDTO {
	this := CampaignVipEmplBonusDTO{}
	return &this
}

// GetClaimed returns the Claimed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignVipEmplBonusDTO) GetClaimed() bool {
	if o == nil || IsNil(o.Claimed.Get()) {
		var ret bool
		return ret
	}
	return *o.Claimed.Get()
}

// GetClaimedOk returns a tuple with the Claimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignVipEmplBonusDTO) GetClaimedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Claimed.Get(), o.Claimed.IsSet()
}

// HasClaimed returns a boolean if a field has been set.
func (o *CampaignVipEmplBonusDTO) HasClaimed() bool {
	if o != nil && o.Claimed.IsSet() {
		return true
	}

	return false
}

// SetClaimed gets a reference to the given NullableBool and assigns it to the Claimed field.
func (o *CampaignVipEmplBonusDTO) SetClaimed(v bool) {
	o.Claimed.Set(&v)
}
// SetClaimedNil sets the value for Claimed to be an explicit nil
func (o *CampaignVipEmplBonusDTO) SetClaimedNil() {
	o.Claimed.Set(nil)
}

// UnsetClaimed ensures that no value is present for Claimed, not even an explicit nil
func (o *CampaignVipEmplBonusDTO) UnsetClaimed() {
	o.Claimed.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignVipEmplBonusDTO) GetDetails() DetailsDTO {
	if o == nil || IsNil(o.Details.Get()) {
		var ret DetailsDTO
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignVipEmplBonusDTO) GetDetailsOk() (*DetailsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *CampaignVipEmplBonusDTO) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableDetailsDTO and assigns it to the Details field.
func (o *CampaignVipEmplBonusDTO) SetDetails(v DetailsDTO) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *CampaignVipEmplBonusDTO) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *CampaignVipEmplBonusDTO) UnsetDetails() {
	o.Details.Unset()
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignVipEmplBonusDTO) GetEndAt() int64 {
	if o == nil || IsNil(o.EndAt.Get()) {
		var ret int64
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignVipEmplBonusDTO) GetEndAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *CampaignVipEmplBonusDTO) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableInt64 and assigns it to the EndAt field.
func (o *CampaignVipEmplBonusDTO) SetEndAt(v int64) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *CampaignVipEmplBonusDTO) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *CampaignVipEmplBonusDTO) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignVipEmplBonusDTO) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignVipEmplBonusDTO) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CampaignVipEmplBonusDTO) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CampaignVipEmplBonusDTO) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CampaignVipEmplBonusDTO) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CampaignVipEmplBonusDTO) UnsetKey() {
	o.Key.Unset()
}

func (o CampaignVipEmplBonusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignVipEmplBonusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Claimed.IsSet() {
		toSerialize["claimed"] = o.Claimed.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.EndAt.IsSet() {
		toSerialize["end_at"] = o.EndAt.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignVipEmplBonusDTO) UnmarshalJSON(data []byte) (err error) {
	varCampaignVipEmplBonusDTO := _CampaignVipEmplBonusDTO{}

	err = json.Unmarshal(data, &varCampaignVipEmplBonusDTO)

	if err != nil {
		return err
	}

	*o = CampaignVipEmplBonusDTO(varCampaignVipEmplBonusDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "claimed")
		delete(additionalProperties, "details")
		delete(additionalProperties, "end_at")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignVipEmplBonusDTO struct {
	value *CampaignVipEmplBonusDTO
	isSet bool
}

func (v NullableCampaignVipEmplBonusDTO) Get() *CampaignVipEmplBonusDTO {
	return v.value
}

func (v *NullableCampaignVipEmplBonusDTO) Set(val *CampaignVipEmplBonusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignVipEmplBonusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignVipEmplBonusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignVipEmplBonusDTO(val *CampaignVipEmplBonusDTO) *NullableCampaignVipEmplBonusDTO {
	return &NullableCampaignVipEmplBonusDTO{value: val, isSet: true}
}

func (v NullableCampaignVipEmplBonusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignVipEmplBonusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


