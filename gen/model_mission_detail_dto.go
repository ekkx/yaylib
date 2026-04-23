
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionDetailDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionDetailDTO{}

// MissionDetailDTO struct for MissionDetailDTO
type MissionDetailDTO struct {
	Action NullableString `json:"action,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	MissionId NullableInt64 `json:"mission_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionDetailDTO MissionDetailDTO

// NewMissionDetailDTO instantiates a new MissionDetailDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionDetailDTO() *MissionDetailDTO {
	this := MissionDetailDTO{}
	return &this
}

// NewMissionDetailDTOWithDefaults instantiates a new MissionDetailDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionDetailDTOWithDefaults() *MissionDetailDTO {
	this := MissionDetailDTO{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDetailDTO) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDetailDTO) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *MissionDetailDTO) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *MissionDetailDTO) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *MissionDetailDTO) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *MissionDetailDTO) UnsetAction() {
	o.Action.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDetailDTO) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDetailDTO) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *MissionDetailDTO) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *MissionDetailDTO) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *MissionDetailDTO) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *MissionDetailDTO) UnsetDetail() {
	o.Detail.Unset()
}

// GetMissionId returns the MissionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionDetailDTO) GetMissionId() int64 {
	if o == nil || IsNil(o.MissionId.Get()) {
		var ret int64
		return ret
	}
	return *o.MissionId.Get()
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionDetailDTO) GetMissionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MissionId.Get(), o.MissionId.IsSet()
}

// HasMissionId returns a boolean if a field has been set.
func (o *MissionDetailDTO) HasMissionId() bool {
	if o != nil && o.MissionId.IsSet() {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given NullableInt64 and assigns it to the MissionId field.
func (o *MissionDetailDTO) SetMissionId(v int64) {
	o.MissionId.Set(&v)
}
// SetMissionIdNil sets the value for MissionId to be an explicit nil
func (o *MissionDetailDTO) SetMissionIdNil() {
	o.MissionId.Set(nil)
}

// UnsetMissionId ensures that no value is present for MissionId, not even an explicit nil
func (o *MissionDetailDTO) UnsetMissionId() {
	o.MissionId.Unset()
}

func (o MissionDetailDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionDetailDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.MissionId.IsSet() {
		toSerialize["mission_id"] = o.MissionId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionDetailDTO) UnmarshalJSON(data []byte) (err error) {
	varMissionDetailDTO := _MissionDetailDTO{}

	err = json.Unmarshal(data, &varMissionDetailDTO)

	if err != nil {
		return err
	}

	*o = MissionDetailDTO(varMissionDetailDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "mission_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionDetailDTO struct {
	value *MissionDetailDTO
	isSet bool
}

func (v NullableMissionDetailDTO) Get() *MissionDetailDTO {
	return v.value
}

func (v *NullableMissionDetailDTO) Set(val *MissionDetailDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionDetailDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionDetailDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionDetailDTO(val *MissionDetailDTO) *NullableMissionDetailDTO {
	return &NullableMissionDetailDTO{value: val, isSet: true}
}

func (v NullableMissionDetailDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionDetailDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


