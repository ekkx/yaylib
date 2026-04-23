
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SectionNavigation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionNavigation{}

// SectionNavigation struct for SectionNavigation
type SectionNavigation struct {
	IconId NullableInt32 `json:"icon_id,omitempty"`
	IsEnabled NullableBool `json:"is_enabled,omitempty"`
	SectionIndex NullableInt32 `json:"section_index,omitempty"`
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SectionNavigation SectionNavigation

// NewSectionNavigation instantiates a new SectionNavigation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionNavigation() *SectionNavigation {
	this := SectionNavigation{}
	return &this
}

// NewSectionNavigationWithDefaults instantiates a new SectionNavigation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionNavigationWithDefaults() *SectionNavigation {
	this := SectionNavigation{}
	return &this
}

// GetIconId returns the IconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionNavigation) GetIconId() int32 {
	if o == nil || IsNil(o.IconId.Get()) {
		var ret int32
		return ret
	}
	return *o.IconId.Get()
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionNavigation) GetIconIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconId.Get(), o.IconId.IsSet()
}

// HasIconId returns a boolean if a field has been set.
func (o *SectionNavigation) HasIconId() bool {
	if o != nil && o.IconId.IsSet() {
		return true
	}

	return false
}

// SetIconId gets a reference to the given NullableInt32 and assigns it to the IconId field.
func (o *SectionNavigation) SetIconId(v int32) {
	o.IconId.Set(&v)
}
// SetIconIdNil sets the value for IconId to be an explicit nil
func (o *SectionNavigation) SetIconIdNil() {
	o.IconId.Set(nil)
}

// UnsetIconId ensures that no value is present for IconId, not even an explicit nil
func (o *SectionNavigation) UnsetIconId() {
	o.IconId.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionNavigation) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionNavigation) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SectionNavigation) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *SectionNavigation) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *SectionNavigation) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *SectionNavigation) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetSectionIndex returns the SectionIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionNavigation) GetSectionIndex() int32 {
	if o == nil || IsNil(o.SectionIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SectionIndex.Get()
}

// GetSectionIndexOk returns a tuple with the SectionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionNavigation) GetSectionIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionIndex.Get(), o.SectionIndex.IsSet()
}

// HasSectionIndex returns a boolean if a field has been set.
func (o *SectionNavigation) HasSectionIndex() bool {
	if o != nil && o.SectionIndex.IsSet() {
		return true
	}

	return false
}

// SetSectionIndex gets a reference to the given NullableInt32 and assigns it to the SectionIndex field.
func (o *SectionNavigation) SetSectionIndex(v int32) {
	o.SectionIndex.Set(&v)
}
// SetSectionIndexNil sets the value for SectionIndex to be an explicit nil
func (o *SectionNavigation) SetSectionIndexNil() {
	o.SectionIndex.Set(nil)
}

// UnsetSectionIndex ensures that no value is present for SectionIndex, not even an explicit nil
func (o *SectionNavigation) UnsetSectionIndex() {
	o.SectionIndex.Unset()
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionNavigation) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionNavigation) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *SectionNavigation) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *SectionNavigation) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *SectionNavigation) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *SectionNavigation) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o SectionNavigation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectionNavigation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IconId.IsSet() {
		toSerialize["icon_id"] = o.IconId.Get()
	}
	if o.IsEnabled.IsSet() {
		toSerialize["is_enabled"] = o.IsEnabled.Get()
	}
	if o.SectionIndex.IsSet() {
		toSerialize["section_index"] = o.SectionIndex.Get()
	}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SectionNavigation) UnmarshalJSON(data []byte) (err error) {
	varSectionNavigation := _SectionNavigation{}

	err = json.Unmarshal(data, &varSectionNavigation)

	if err != nil {
		return err
	}

	*o = SectionNavigation(varSectionNavigation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon_id")
		delete(additionalProperties, "is_enabled")
		delete(additionalProperties, "section_index")
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSectionNavigation struct {
	value *SectionNavigation
	isSet bool
}

func (v NullableSectionNavigation) Get() *SectionNavigation {
	return v.value
}

func (v *NullableSectionNavigation) Set(val *SectionNavigation) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionNavigation) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionNavigation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionNavigation(val *SectionNavigation) *NullableSectionNavigation {
	return &NullableSectionNavigation{value: val, isSet: true}
}

func (v NullableSectionNavigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionNavigation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


