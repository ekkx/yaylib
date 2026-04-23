
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionSectionHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionSectionHeader{}

// MissionSectionHeader struct for MissionSectionHeader
type MissionSectionHeader struct {
	SectionInfoId NullableInt32 `json:"section_info_id,omitempty"`
	SectionTitleId NullableInt32 `json:"section_title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionSectionHeader MissionSectionHeader

// NewMissionSectionHeader instantiates a new MissionSectionHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionSectionHeader() *MissionSectionHeader {
	this := MissionSectionHeader{}
	return &this
}

// NewMissionSectionHeaderWithDefaults instantiates a new MissionSectionHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionSectionHeaderWithDefaults() *MissionSectionHeader {
	this := MissionSectionHeader{}
	return &this
}

// GetSectionInfoId returns the SectionInfoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionSectionHeader) GetSectionInfoId() int32 {
	if o == nil || IsNil(o.SectionInfoId.Get()) {
		var ret int32
		return ret
	}
	return *o.SectionInfoId.Get()
}

// GetSectionInfoIdOk returns a tuple with the SectionInfoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionSectionHeader) GetSectionInfoIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionInfoId.Get(), o.SectionInfoId.IsSet()
}

// HasSectionInfoId returns a boolean if a field has been set.
func (o *MissionSectionHeader) HasSectionInfoId() bool {
	if o != nil && o.SectionInfoId.IsSet() {
		return true
	}

	return false
}

// SetSectionInfoId gets a reference to the given NullableInt32 and assigns it to the SectionInfoId field.
func (o *MissionSectionHeader) SetSectionInfoId(v int32) {
	o.SectionInfoId.Set(&v)
}
// SetSectionInfoIdNil sets the value for SectionInfoId to be an explicit nil
func (o *MissionSectionHeader) SetSectionInfoIdNil() {
	o.SectionInfoId.Set(nil)
}

// UnsetSectionInfoId ensures that no value is present for SectionInfoId, not even an explicit nil
func (o *MissionSectionHeader) UnsetSectionInfoId() {
	o.SectionInfoId.Unset()
}

// GetSectionTitleId returns the SectionTitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionSectionHeader) GetSectionTitleId() int32 {
	if o == nil || IsNil(o.SectionTitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.SectionTitleId.Get()
}

// GetSectionTitleIdOk returns a tuple with the SectionTitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionSectionHeader) GetSectionTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionTitleId.Get(), o.SectionTitleId.IsSet()
}

// HasSectionTitleId returns a boolean if a field has been set.
func (o *MissionSectionHeader) HasSectionTitleId() bool {
	if o != nil && o.SectionTitleId.IsSet() {
		return true
	}

	return false
}

// SetSectionTitleId gets a reference to the given NullableInt32 and assigns it to the SectionTitleId field.
func (o *MissionSectionHeader) SetSectionTitleId(v int32) {
	o.SectionTitleId.Set(&v)
}
// SetSectionTitleIdNil sets the value for SectionTitleId to be an explicit nil
func (o *MissionSectionHeader) SetSectionTitleIdNil() {
	o.SectionTitleId.Set(nil)
}

// UnsetSectionTitleId ensures that no value is present for SectionTitleId, not even an explicit nil
func (o *MissionSectionHeader) UnsetSectionTitleId() {
	o.SectionTitleId.Unset()
}

func (o MissionSectionHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionSectionHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SectionInfoId.IsSet() {
		toSerialize["section_info_id"] = o.SectionInfoId.Get()
	}
	if o.SectionTitleId.IsSet() {
		toSerialize["section_title_id"] = o.SectionTitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionSectionHeader) UnmarshalJSON(data []byte) (err error) {
	varMissionSectionHeader := _MissionSectionHeader{}

	err = json.Unmarshal(data, &varMissionSectionHeader)

	if err != nil {
		return err
	}

	*o = MissionSectionHeader(varMissionSectionHeader)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "section_info_id")
		delete(additionalProperties, "section_title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionSectionHeader struct {
	value *MissionSectionHeader
	isSet bool
}

func (v NullableMissionSectionHeader) Get() *MissionSectionHeader {
	return v.value
}

func (v *NullableMissionSectionHeader) Set(val *MissionSectionHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionSectionHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionSectionHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionSectionHeader(val *MissionSectionHeader) *NullableMissionSectionHeader {
	return &NullableMissionSectionHeader{value: val, isSet: true}
}

func (v NullableMissionSectionHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionSectionHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


