
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionDTO{}

// MissionDTO struct for MissionDTO
type MissionDTO struct {
	Action NullableString `json:"action,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MissionType NullableString `json:"mission_type,omitempty"`
	Progress NullableProgressDTO `json:"progress,omitempty"`
	RequiredActionsCount NullableInt32 `json:"required_actions_count,omitempty"`
	RewardPoints NullableInt32 `json:"reward_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionDTO MissionDTO

// NewMissionDTO instantiates a new MissionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionDTO() *MissionDTO {
	this := MissionDTO{}
	return &this
}

// NewMissionDTOWithDefaults instantiates a new MissionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionDTOWithDefaults() *MissionDTO {
	this := MissionDTO{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *MissionDTO) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *MissionDTO) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *MissionDTO) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *MissionDTO) UnsetAction() {
	o.Action.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MissionDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *MissionDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MissionDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MissionDTO) UnsetId() {
	o.Id.Unset()
}

// GetMissionType returns the MissionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetMissionType() string {
	if o == nil || IsNil(o.MissionType.Get()) {
		var ret string
		return ret
	}
	return *o.MissionType.Get()
}

// GetMissionTypeOk returns a tuple with the MissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetMissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MissionType.Get(), o.MissionType.IsSet()
}

// HasMissionType returns a boolean if a field has been set.
func (o *MissionDTO) HasMissionType() bool {
	if o != nil && o.MissionType.IsSet() {
		return true
	}

	return false
}

// SetMissionType gets a reference to the given NullableString and assigns it to the MissionType field.
func (o *MissionDTO) SetMissionType(v string) {
	o.MissionType.Set(&v)
}
// SetMissionTypeNil sets the value for MissionType to be an explicit nil
func (o *MissionDTO) SetMissionTypeNil() {
	o.MissionType.Set(nil)
}

// UnsetMissionType ensures that no value is present for MissionType, not even an explicit nil
func (o *MissionDTO) UnsetMissionType() {
	o.MissionType.Unset()
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetProgress() ProgressDTO {
	if o == nil || IsNil(o.Progress.Get()) {
		var ret ProgressDTO
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetProgressOk() (*ProgressDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *MissionDTO) HasProgress() bool {
	if o != nil && o.Progress.IsSet() {
		return true
	}

	return false
}

// SetProgress gets a reference to the given NullableProgressDTO and assigns it to the Progress field.
func (o *MissionDTO) SetProgress(v ProgressDTO) {
	o.Progress.Set(&v)
}
// SetProgressNil sets the value for Progress to be an explicit nil
func (o *MissionDTO) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil
func (o *MissionDTO) UnsetProgress() {
	o.Progress.Unset()
}

// GetRequiredActionsCount returns the RequiredActionsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetRequiredActionsCount() int32 {
	if o == nil || IsNil(o.RequiredActionsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RequiredActionsCount.Get()
}

// GetRequiredActionsCountOk returns a tuple with the RequiredActionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetRequiredActionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredActionsCount.Get(), o.RequiredActionsCount.IsSet()
}

// HasRequiredActionsCount returns a boolean if a field has been set.
func (o *MissionDTO) HasRequiredActionsCount() bool {
	if o != nil && o.RequiredActionsCount.IsSet() {
		return true
	}

	return false
}

// SetRequiredActionsCount gets a reference to the given NullableInt32 and assigns it to the RequiredActionsCount field.
func (o *MissionDTO) SetRequiredActionsCount(v int32) {
	o.RequiredActionsCount.Set(&v)
}
// SetRequiredActionsCountNil sets the value for RequiredActionsCount to be an explicit nil
func (o *MissionDTO) SetRequiredActionsCountNil() {
	o.RequiredActionsCount.Set(nil)
}

// UnsetRequiredActionsCount ensures that no value is present for RequiredActionsCount, not even an explicit nil
func (o *MissionDTO) UnsetRequiredActionsCount() {
	o.RequiredActionsCount.Unset()
}

// GetRewardPoints returns the RewardPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDTO) GetRewardPoints() int32 {
	if o == nil || IsNil(o.RewardPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardPoints.Get()
}

// GetRewardPointsOk returns a tuple with the RewardPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDTO) GetRewardPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardPoints.Get(), o.RewardPoints.IsSet()
}

// HasRewardPoints returns a boolean if a field has been set.
func (o *MissionDTO) HasRewardPoints() bool {
	if o != nil && o.RewardPoints.IsSet() {
		return true
	}

	return false
}

// SetRewardPoints gets a reference to the given NullableInt32 and assigns it to the RewardPoints field.
func (o *MissionDTO) SetRewardPoints(v int32) {
	o.RewardPoints.Set(&v)
}
// SetRewardPointsNil sets the value for RewardPoints to be an explicit nil
func (o *MissionDTO) SetRewardPointsNil() {
	o.RewardPoints.Set(nil)
}

// UnsetRewardPoints ensures that no value is present for RewardPoints, not even an explicit nil
func (o *MissionDTO) UnsetRewardPoints() {
	o.RewardPoints.Unset()
}

func (o MissionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MissionType.IsSet() {
		toSerialize["mission_type"] = o.MissionType.Get()
	}
	if o.Progress.IsSet() {
		toSerialize["progress"] = o.Progress.Get()
	}
	if o.RequiredActionsCount.IsSet() {
		toSerialize["required_actions_count"] = o.RequiredActionsCount.Get()
	}
	if o.RewardPoints.IsSet() {
		toSerialize["reward_points"] = o.RewardPoints.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionDTO) UnmarshalJSON(data []byte) (err error) {
	varMissionDTO := _MissionDTO{}

	err = json.Unmarshal(data, &varMissionDTO)

	if err != nil {
		return err
	}

	*o = MissionDTO(varMissionDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mission_type")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "required_actions_count")
		delete(additionalProperties, "reward_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionDTO struct {
	value *MissionDTO
	isSet bool
}

func (v NullableMissionDTO) Get() *MissionDTO {
	return v.value
}

func (v *NullableMissionDTO) Set(val *MissionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionDTO(val *MissionDTO) *NullableMissionDTO {
	return &NullableMissionDTO{value: val, isSet: true}
}

func (v NullableMissionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


