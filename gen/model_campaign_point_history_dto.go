
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignPointHistoryDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignPointHistoryDTO{}

// CampaignPointHistoryDTO struct for CampaignPointHistoryDTO
type CampaignPointHistoryDTO struct {
	Breakdown NullableBreakdown `json:"breakdown,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Mission NullableMission `json:"mission,omitempty"`
	Multiplier NullableFloat64 `json:"multiplier,omitempty"`
	PointType NullableString `json:"point_type,omitempty"`
	TotalPoints NullableInt32 `json:"total_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignPointHistoryDTO CampaignPointHistoryDTO

// NewCampaignPointHistoryDTO instantiates a new CampaignPointHistoryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignPointHistoryDTO() *CampaignPointHistoryDTO {
	this := CampaignPointHistoryDTO{}
	return &this
}

// NewCampaignPointHistoryDTOWithDefaults instantiates a new CampaignPointHistoryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignPointHistoryDTOWithDefaults() *CampaignPointHistoryDTO {
	this := CampaignPointHistoryDTO{}
	return &this
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetBreakdown() Breakdown {
	if o == nil || IsNil(o.Breakdown.Get()) {
		var ret Breakdown
		return ret
	}
	return *o.Breakdown.Get()
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetBreakdownOk() (*Breakdown, bool) {
	if o == nil {
		return nil, false
	}
	return o.Breakdown.Get(), o.Breakdown.IsSet()
}

// HasBreakdown returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasBreakdown() bool {
	if o != nil && o.Breakdown.IsSet() {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given NullableBreakdown and assigns it to the Breakdown field.
func (o *CampaignPointHistoryDTO) SetBreakdown(v Breakdown) {
	o.Breakdown.Set(&v)
}
// SetBreakdownNil sets the value for Breakdown to be an explicit nil
func (o *CampaignPointHistoryDTO) SetBreakdownNil() {
	o.Breakdown.Set(nil)
}

// UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetBreakdown() {
	o.Breakdown.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *CampaignPointHistoryDTO) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CampaignPointHistoryDTO) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CampaignPointHistoryDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CampaignPointHistoryDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetId() {
	o.Id.Unset()
}

// GetMission returns the Mission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetMission() Mission {
	if o == nil || IsNil(o.Mission.Get()) {
		var ret Mission
		return ret
	}
	return *o.Mission.Get()
}

// GetMissionOk returns a tuple with the Mission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetMissionOk() (*Mission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mission.Get(), o.Mission.IsSet()
}

// HasMission returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasMission() bool {
	if o != nil && o.Mission.IsSet() {
		return true
	}

	return false
}

// SetMission gets a reference to the given NullableMission and assigns it to the Mission field.
func (o *CampaignPointHistoryDTO) SetMission(v Mission) {
	o.Mission.Set(&v)
}
// SetMissionNil sets the value for Mission to be an explicit nil
func (o *CampaignPointHistoryDTO) SetMissionNil() {
	o.Mission.Set(nil)
}

// UnsetMission ensures that no value is present for Mission, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetMission() {
	o.Mission.Unset()
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetMultiplier() float64 {
	if o == nil || IsNil(o.Multiplier.Get()) {
		var ret float64
		return ret
	}
	return *o.Multiplier.Get()
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetMultiplierOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Multiplier.Get(), o.Multiplier.IsSet()
}

// HasMultiplier returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasMultiplier() bool {
	if o != nil && o.Multiplier.IsSet() {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given NullableFloat64 and assigns it to the Multiplier field.
func (o *CampaignPointHistoryDTO) SetMultiplier(v float64) {
	o.Multiplier.Set(&v)
}
// SetMultiplierNil sets the value for Multiplier to be an explicit nil
func (o *CampaignPointHistoryDTO) SetMultiplierNil() {
	o.Multiplier.Set(nil)
}

// UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetMultiplier() {
	o.Multiplier.Unset()
}

// GetPointType returns the PointType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetPointType() string {
	if o == nil || IsNil(o.PointType.Get()) {
		var ret string
		return ret
	}
	return *o.PointType.Get()
}

// GetPointTypeOk returns a tuple with the PointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetPointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointType.Get(), o.PointType.IsSet()
}

// HasPointType returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasPointType() bool {
	if o != nil && o.PointType.IsSet() {
		return true
	}

	return false
}

// SetPointType gets a reference to the given NullableString and assigns it to the PointType field.
func (o *CampaignPointHistoryDTO) SetPointType(v string) {
	o.PointType.Set(&v)
}
// SetPointTypeNil sets the value for PointType to be an explicit nil
func (o *CampaignPointHistoryDTO) SetPointTypeNil() {
	o.PointType.Set(nil)
}

// UnsetPointType ensures that no value is present for PointType, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetPointType() {
	o.PointType.Unset()
}

// GetTotalPoints returns the TotalPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryDTO) GetTotalPoints() int32 {
	if o == nil || IsNil(o.TotalPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalPoints.Get()
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryDTO) GetTotalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPoints.Get(), o.TotalPoints.IsSet()
}

// HasTotalPoints returns a boolean if a field has been set.
func (o *CampaignPointHistoryDTO) HasTotalPoints() bool {
	if o != nil && o.TotalPoints.IsSet() {
		return true
	}

	return false
}

// SetTotalPoints gets a reference to the given NullableInt32 and assigns it to the TotalPoints field.
func (o *CampaignPointHistoryDTO) SetTotalPoints(v int32) {
	o.TotalPoints.Set(&v)
}
// SetTotalPointsNil sets the value for TotalPoints to be an explicit nil
func (o *CampaignPointHistoryDTO) SetTotalPointsNil() {
	o.TotalPoints.Set(nil)
}

// UnsetTotalPoints ensures that no value is present for TotalPoints, not even an explicit nil
func (o *CampaignPointHistoryDTO) UnsetTotalPoints() {
	o.TotalPoints.Unset()
}

func (o CampaignPointHistoryDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignPointHistoryDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Breakdown.IsSet() {
		toSerialize["breakdown"] = o.Breakdown.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Mission.IsSet() {
		toSerialize["mission"] = o.Mission.Get()
	}
	if o.Multiplier.IsSet() {
		toSerialize["multiplier"] = o.Multiplier.Get()
	}
	if o.PointType.IsSet() {
		toSerialize["point_type"] = o.PointType.Get()
	}
	if o.TotalPoints.IsSet() {
		toSerialize["total_points"] = o.TotalPoints.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignPointHistoryDTO) UnmarshalJSON(data []byte) (err error) {
	varCampaignPointHistoryDTO := _CampaignPointHistoryDTO{}

	err = json.Unmarshal(data, &varCampaignPointHistoryDTO)

	if err != nil {
		return err
	}

	*o = CampaignPointHistoryDTO(varCampaignPointHistoryDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "breakdown")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mission")
		delete(additionalProperties, "multiplier")
		delete(additionalProperties, "point_type")
		delete(additionalProperties, "total_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignPointHistoryDTO struct {
	value *CampaignPointHistoryDTO
	isSet bool
}

func (v NullableCampaignPointHistoryDTO) Get() *CampaignPointHistoryDTO {
	return v.value
}

func (v *NullableCampaignPointHistoryDTO) Set(val *CampaignPointHistoryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignPointHistoryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignPointHistoryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignPointHistoryDTO(val *CampaignPointHistoryDTO) *NullableCampaignPointHistoryDTO {
	return &NullableCampaignPointHistoryDTO{value: val, isSet: true}
}

func (v NullableCampaignPointHistoryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignPointHistoryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


