
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelUserRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUserRank{}

// ModelUserRank struct for ModelUserRank
type ModelUserRank struct {
	DisabledIconId NullableInt32 `json:"disabled_icon_id,omitempty"`
	EnabledIconId NullableInt32 `json:"enabled_icon_id,omitempty"`
	Value NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelUserRank ModelUserRank

// NewModelUserRank instantiates a new ModelUserRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUserRank() *ModelUserRank {
	this := ModelUserRank{}
	return &this
}

// NewModelUserRankWithDefaults instantiates a new ModelUserRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserRankWithDefaults() *ModelUserRank {
	this := ModelUserRank{}
	return &this
}

// GetDisabledIconId returns the DisabledIconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserRank) GetDisabledIconId() int32 {
	if o == nil || IsNil(o.DisabledIconId.Get()) {
		var ret int32
		return ret
	}
	return *o.DisabledIconId.Get()
}

// GetDisabledIconIdOk returns a tuple with the DisabledIconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserRank) GetDisabledIconIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisabledIconId.Get(), o.DisabledIconId.IsSet()
}

// HasDisabledIconId returns a boolean if a field has been set.
func (o *ModelUserRank) HasDisabledIconId() bool {
	if o != nil && o.DisabledIconId.IsSet() {
		return true
	}

	return false
}

// SetDisabledIconId gets a reference to the given NullableInt32 and assigns it to the DisabledIconId field.
func (o *ModelUserRank) SetDisabledIconId(v int32) {
	o.DisabledIconId.Set(&v)
}
// SetDisabledIconIdNil sets the value for DisabledIconId to be an explicit nil
func (o *ModelUserRank) SetDisabledIconIdNil() {
	o.DisabledIconId.Set(nil)
}

// UnsetDisabledIconId ensures that no value is present for DisabledIconId, not even an explicit nil
func (o *ModelUserRank) UnsetDisabledIconId() {
	o.DisabledIconId.Unset()
}

// GetEnabledIconId returns the EnabledIconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserRank) GetEnabledIconId() int32 {
	if o == nil || IsNil(o.EnabledIconId.Get()) {
		var ret int32
		return ret
	}
	return *o.EnabledIconId.Get()
}

// GetEnabledIconIdOk returns a tuple with the EnabledIconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserRank) GetEnabledIconIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledIconId.Get(), o.EnabledIconId.IsSet()
}

// HasEnabledIconId returns a boolean if a field has been set.
func (o *ModelUserRank) HasEnabledIconId() bool {
	if o != nil && o.EnabledIconId.IsSet() {
		return true
	}

	return false
}

// SetEnabledIconId gets a reference to the given NullableInt32 and assigns it to the EnabledIconId field.
func (o *ModelUserRank) SetEnabledIconId(v int32) {
	o.EnabledIconId.Set(&v)
}
// SetEnabledIconIdNil sets the value for EnabledIconId to be an explicit nil
func (o *ModelUserRank) SetEnabledIconIdNil() {
	o.EnabledIconId.Set(nil)
}

// UnsetEnabledIconId ensures that no value is present for EnabledIconId, not even an explicit nil
func (o *ModelUserRank) UnsetEnabledIconId() {
	o.EnabledIconId.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserRank) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserRank) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ModelUserRank) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ModelUserRank) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ModelUserRank) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ModelUserRank) UnsetValue() {
	o.Value.Unset()
}

func (o ModelUserRank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUserRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisabledIconId.IsSet() {
		toSerialize["disabled_icon_id"] = o.DisabledIconId.Get()
	}
	if o.EnabledIconId.IsSet() {
		toSerialize["enabled_icon_id"] = o.EnabledIconId.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelUserRank) UnmarshalJSON(data []byte) (err error) {
	varModelUserRank := _ModelUserRank{}

	err = json.Unmarshal(data, &varModelUserRank)

	if err != nil {
		return err
	}

	*o = ModelUserRank(varModelUserRank)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled_icon_id")
		delete(additionalProperties, "enabled_icon_id")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelUserRank struct {
	value *ModelUserRank
	isSet bool
}

func (v NullableModelUserRank) Get() *ModelUserRank {
	return v.value
}

func (v *NullableModelUserRank) Set(val *ModelUserRank) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUserRank) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUserRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUserRank(val *ModelUserRank) *NullableModelUserRank {
	return &NullableModelUserRank{value: val, isSet: true}
}

func (v NullableModelUserRank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUserRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


