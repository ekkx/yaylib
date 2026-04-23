
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionTypeX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionTypeX{}

// MissionTypeX struct for MissionTypeX
type MissionTypeX struct {
	SectionInfoId NullableInt32 `json:"section_info_id,omitempty"`
	SectionTitleId NullableInt32 `json:"section_title_id,omitempty"`
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionTypeX MissionTypeX

// NewMissionTypeX instantiates a new MissionTypeX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionTypeX() *MissionTypeX {
	this := MissionTypeX{}
	return &this
}

// NewMissionTypeXWithDefaults instantiates a new MissionTypeX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionTypeXWithDefaults() *MissionTypeX {
	this := MissionTypeX{}
	return &this
}

// GetSectionInfoId returns the SectionInfoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionTypeX) GetSectionInfoId() int32 {
	if o == nil || IsNil(o.SectionInfoId.Get()) {
		var ret int32
		return ret
	}
	return *o.SectionInfoId.Get()
}

// GetSectionInfoIdOk returns a tuple with the SectionInfoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionTypeX) GetSectionInfoIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionInfoId.Get(), o.SectionInfoId.IsSet()
}

// HasSectionInfoId returns a boolean if a field has been set.
func (o *MissionTypeX) HasSectionInfoId() bool {
	if o != nil && o.SectionInfoId.IsSet() {
		return true
	}

	return false
}

// SetSectionInfoId gets a reference to the given NullableInt32 and assigns it to the SectionInfoId field.
func (o *MissionTypeX) SetSectionInfoId(v int32) {
	o.SectionInfoId.Set(&v)
}
// SetSectionInfoIdNil sets the value for SectionInfoId to be an explicit nil
func (o *MissionTypeX) SetSectionInfoIdNil() {
	o.SectionInfoId.Set(nil)
}

// UnsetSectionInfoId ensures that no value is present for SectionInfoId, not even an explicit nil
func (o *MissionTypeX) UnsetSectionInfoId() {
	o.SectionInfoId.Unset()
}

// GetSectionTitleId returns the SectionTitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionTypeX) GetSectionTitleId() int32 {
	if o == nil || IsNil(o.SectionTitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.SectionTitleId.Get()
}

// GetSectionTitleIdOk returns a tuple with the SectionTitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionTypeX) GetSectionTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionTitleId.Get(), o.SectionTitleId.IsSet()
}

// HasSectionTitleId returns a boolean if a field has been set.
func (o *MissionTypeX) HasSectionTitleId() bool {
	if o != nil && o.SectionTitleId.IsSet() {
		return true
	}

	return false
}

// SetSectionTitleId gets a reference to the given NullableInt32 and assigns it to the SectionTitleId field.
func (o *MissionTypeX) SetSectionTitleId(v int32) {
	o.SectionTitleId.Set(&v)
}
// SetSectionTitleIdNil sets the value for SectionTitleId to be an explicit nil
func (o *MissionTypeX) SetSectionTitleIdNil() {
	o.SectionTitleId.Set(nil)
}

// UnsetSectionTitleId ensures that no value is present for SectionTitleId, not even an explicit nil
func (o *MissionTypeX) UnsetSectionTitleId() {
	o.SectionTitleId.Unset()
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionTypeX) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionTypeX) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *MissionTypeX) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *MissionTypeX) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *MissionTypeX) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *MissionTypeX) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o MissionTypeX) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionTypeX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SectionInfoId.IsSet() {
		toSerialize["section_info_id"] = o.SectionInfoId.Get()
	}
	if o.SectionTitleId.IsSet() {
		toSerialize["section_title_id"] = o.SectionTitleId.Get()
	}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionTypeX) UnmarshalJSON(data []byte) (err error) {
	varMissionTypeX := _MissionTypeX{}

	err = json.Unmarshal(data, &varMissionTypeX)

	if err != nil {
		return err
	}

	*o = MissionTypeX(varMissionTypeX)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "section_info_id")
		delete(additionalProperties, "section_title_id")
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionTypeX struct {
	value *MissionTypeX
	isSet bool
}

func (v NullableMissionTypeX) Get() *MissionTypeX {
	return v.value
}

func (v *NullableMissionTypeX) Set(val *MissionTypeX) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionTypeX) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionTypeX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionTypeX(val *MissionTypeX) *NullableMissionTypeX {
	return &NullableMissionTypeX{value: val, isSet: true}
}

func (v NullableMissionTypeX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionTypeX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


