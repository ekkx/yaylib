
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionSection{}

// MissionSection struct for MissionSection
type MissionSection struct {
	MissionList []MissionItem `json:"mission_list,omitempty"`
	MissionType map[string]interface{} `json:"mission_type,omitempty"`
	Section NullableMissionSectionHeader `json:"section,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionSection MissionSection

// NewMissionSection instantiates a new MissionSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionSection() *MissionSection {
	this := MissionSection{}
	return &this
}

// NewMissionSectionWithDefaults instantiates a new MissionSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionSectionWithDefaults() *MissionSection {
	this := MissionSection{}
	return &this
}

// GetMissionList returns the MissionList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionSection) GetMissionList() []MissionItem {
	if o == nil {
		var ret []MissionItem
		return ret
	}
	return o.MissionList
}

// GetMissionListOk returns a tuple with the MissionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionSection) GetMissionListOk() ([]MissionItem, bool) {
	if o == nil || IsNil(o.MissionList) {
		return nil, false
	}
	return o.MissionList, true
}

// HasMissionList returns a boolean if a field has been set.
func (o *MissionSection) HasMissionList() bool {
	if o != nil && !IsNil(o.MissionList) {
		return true
	}

	return false
}

// SetMissionList gets a reference to the given []MissionItem and assigns it to the MissionList field.
func (o *MissionSection) SetMissionList(v []MissionItem) {
	o.MissionList = v
}

// GetMissionType returns the MissionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionSection) GetMissionType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MissionType
}

// GetMissionTypeOk returns a tuple with the MissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionSection) GetMissionTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MissionType) {
		return map[string]interface{}{}, false
	}
	return o.MissionType, true
}

// HasMissionType returns a boolean if a field has been set.
func (o *MissionSection) HasMissionType() bool {
	if o != nil && !IsNil(o.MissionType) {
		return true
	}

	return false
}

// SetMissionType gets a reference to the given map[string]interface{} and assigns it to the MissionType field.
func (o *MissionSection) SetMissionType(v map[string]interface{}) {
	o.MissionType = v
}

// GetSection returns the Section field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionSection) GetSection() MissionSectionHeader {
	if o == nil || IsNil(o.Section.Get()) {
		var ret MissionSectionHeader
		return ret
	}
	return *o.Section.Get()
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionSection) GetSectionOk() (*MissionSectionHeader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Section.Get(), o.Section.IsSet()
}

// HasSection returns a boolean if a field has been set.
func (o *MissionSection) HasSection() bool {
	if o != nil && o.Section.IsSet() {
		return true
	}

	return false
}

// SetSection gets a reference to the given NullableMissionSectionHeader and assigns it to the Section field.
func (o *MissionSection) SetSection(v MissionSectionHeader) {
	o.Section.Set(&v)
}
// SetSectionNil sets the value for Section to be an explicit nil
func (o *MissionSection) SetSectionNil() {
	o.Section.Set(nil)
}

// UnsetSection ensures that no value is present for Section, not even an explicit nil
func (o *MissionSection) UnsetSection() {
	o.Section.Unset()
}

func (o MissionSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MissionList != nil {
		toSerialize["mission_list"] = o.MissionList
	}
	if o.MissionType != nil {
		toSerialize["mission_type"] = o.MissionType
	}
	if o.Section.IsSet() {
		toSerialize["section"] = o.Section.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionSection) UnmarshalJSON(data []byte) (err error) {
	varMissionSection := _MissionSection{}

	err = json.Unmarshal(data, &varMissionSection)

	if err != nil {
		return err
	}

	*o = MissionSection(varMissionSection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mission_list")
		delete(additionalProperties, "mission_type")
		delete(additionalProperties, "section")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionSection struct {
	value *MissionSection
	isSet bool
}

func (v NullableMissionSection) Get() *MissionSection {
	return v.value
}

func (v *NullableMissionSection) Set(val *MissionSection) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionSection) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionSection(val *MissionSection) *NullableMissionSection {
	return &NullableMissionSection{value: val, isSet: true}
}

func (v NullableMissionSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


