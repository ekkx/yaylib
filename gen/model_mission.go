
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Mission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mission{}

// Mission struct for Mission
type Mission struct {
	Action NullableString `json:"action,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	IsMultiplier NullableBool `json:"is_multiplier,omitempty"`
	MissionType NullableString `json:"mission_type,omitempty"`
	RequiredActionsCount NullableInt32 `json:"required_actions_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mission Mission

// NewMission instantiates a new Mission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMission() *Mission {
	this := Mission{}
	return &this
}

// NewMissionWithDefaults instantiates a new Mission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionWithDefaults() *Mission {
	this := Mission{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mission) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mission) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *Mission) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *Mission) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *Mission) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *Mission) UnsetAction() {
	o.Action.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mission) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mission) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *Mission) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *Mission) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *Mission) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *Mission) UnsetDetail() {
	o.Detail.Unset()
}

// GetIsMultiplier returns the IsMultiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mission) GetIsMultiplier() bool {
	if o == nil || IsNil(o.IsMultiplier.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMultiplier.Get()
}

// GetIsMultiplierOk returns a tuple with the IsMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mission) GetIsMultiplierOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMultiplier.Get(), o.IsMultiplier.IsSet()
}

// HasIsMultiplier returns a boolean if a field has been set.
func (o *Mission) HasIsMultiplier() bool {
	if o != nil && o.IsMultiplier.IsSet() {
		return true
	}

	return false
}

// SetIsMultiplier gets a reference to the given NullableBool and assigns it to the IsMultiplier field.
func (o *Mission) SetIsMultiplier(v bool) {
	o.IsMultiplier.Set(&v)
}
// SetIsMultiplierNil sets the value for IsMultiplier to be an explicit nil
func (o *Mission) SetIsMultiplierNil() {
	o.IsMultiplier.Set(nil)
}

// UnsetIsMultiplier ensures that no value is present for IsMultiplier, not even an explicit nil
func (o *Mission) UnsetIsMultiplier() {
	o.IsMultiplier.Unset()
}

// GetMissionType returns the MissionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mission) GetMissionType() string {
	if o == nil || IsNil(o.MissionType.Get()) {
		var ret string
		return ret
	}
	return *o.MissionType.Get()
}

// GetMissionTypeOk returns a tuple with the MissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mission) GetMissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MissionType.Get(), o.MissionType.IsSet()
}

// HasMissionType returns a boolean if a field has been set.
func (o *Mission) HasMissionType() bool {
	if o != nil && o.MissionType.IsSet() {
		return true
	}

	return false
}

// SetMissionType gets a reference to the given NullableString and assigns it to the MissionType field.
func (o *Mission) SetMissionType(v string) {
	o.MissionType.Set(&v)
}
// SetMissionTypeNil sets the value for MissionType to be an explicit nil
func (o *Mission) SetMissionTypeNil() {
	o.MissionType.Set(nil)
}

// UnsetMissionType ensures that no value is present for MissionType, not even an explicit nil
func (o *Mission) UnsetMissionType() {
	o.MissionType.Unset()
}

// GetRequiredActionsCount returns the RequiredActionsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mission) GetRequiredActionsCount() int32 {
	if o == nil || IsNil(o.RequiredActionsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RequiredActionsCount.Get()
}

// GetRequiredActionsCountOk returns a tuple with the RequiredActionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mission) GetRequiredActionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredActionsCount.Get(), o.RequiredActionsCount.IsSet()
}

// HasRequiredActionsCount returns a boolean if a field has been set.
func (o *Mission) HasRequiredActionsCount() bool {
	if o != nil && o.RequiredActionsCount.IsSet() {
		return true
	}

	return false
}

// SetRequiredActionsCount gets a reference to the given NullableInt32 and assigns it to the RequiredActionsCount field.
func (o *Mission) SetRequiredActionsCount(v int32) {
	o.RequiredActionsCount.Set(&v)
}
// SetRequiredActionsCountNil sets the value for RequiredActionsCount to be an explicit nil
func (o *Mission) SetRequiredActionsCountNil() {
	o.RequiredActionsCount.Set(nil)
}

// UnsetRequiredActionsCount ensures that no value is present for RequiredActionsCount, not even an explicit nil
func (o *Mission) UnsetRequiredActionsCount() {
	o.RequiredActionsCount.Unset()
}

func (o Mission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.IsMultiplier.IsSet() {
		toSerialize["is_multiplier"] = o.IsMultiplier.Get()
	}
	if o.MissionType.IsSet() {
		toSerialize["mission_type"] = o.MissionType.Get()
	}
	if o.RequiredActionsCount.IsSet() {
		toSerialize["required_actions_count"] = o.RequiredActionsCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mission) UnmarshalJSON(data []byte) (err error) {
	varMission := _Mission{}

	err = json.Unmarshal(data, &varMission)

	if err != nil {
		return err
	}

	*o = Mission(varMission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "is_multiplier")
		delete(additionalProperties, "mission_type")
		delete(additionalProperties, "required_actions_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMission struct {
	value *Mission
	isSet bool
}

func (v NullableMission) Get() *Mission {
	return v.value
}

func (v *NullableMission) Set(val *Mission) {
	v.value = val
	v.isSet = true
}

func (v NullableMission) IsSet() bool {
	return v.isSet
}

func (v *NullableMission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMission(val *Mission) *NullableMission {
	return &NullableMission{value: val, isSet: true}
}

func (v NullableMission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


