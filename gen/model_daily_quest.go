
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DailyQuest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyQuest{}

// DailyQuest struct for DailyQuest
type DailyQuest struct {
	Calculated NullableBool `json:"calculated,omitempty"`
	SpentDurationSeconds NullableInt32 `json:"spent_duration_seconds,omitempty"`
	TargetDurationSeconds NullableInt32 `json:"target_duration_seconds,omitempty"`
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyQuest DailyQuest

// NewDailyQuest instantiates a new DailyQuest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyQuest() *DailyQuest {
	this := DailyQuest{}
	return &this
}

// NewDailyQuestWithDefaults instantiates a new DailyQuest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyQuestWithDefaults() *DailyQuest {
	this := DailyQuest{}
	return &this
}

// GetCalculated returns the Calculated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DailyQuest) GetCalculated() bool {
	if o == nil || IsNil(o.Calculated.Get()) {
		var ret bool
		return ret
	}
	return *o.Calculated.Get()
}

// GetCalculatedOk returns a tuple with the Calculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyQuest) GetCalculatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Calculated.Get(), o.Calculated.IsSet()
}

// HasCalculated returns a boolean if a field has been set.
func (o *DailyQuest) HasCalculated() bool {
	if o != nil && o.Calculated.IsSet() {
		return true
	}

	return false
}

// SetCalculated gets a reference to the given NullableBool and assigns it to the Calculated field.
func (o *DailyQuest) SetCalculated(v bool) {
	o.Calculated.Set(&v)
}
// SetCalculatedNil sets the value for Calculated to be an explicit nil
func (o *DailyQuest) SetCalculatedNil() {
	o.Calculated.Set(nil)
}

// UnsetCalculated ensures that no value is present for Calculated, not even an explicit nil
func (o *DailyQuest) UnsetCalculated() {
	o.Calculated.Unset()
}

// GetSpentDurationSeconds returns the SpentDurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DailyQuest) GetSpentDurationSeconds() int32 {
	if o == nil || IsNil(o.SpentDurationSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.SpentDurationSeconds.Get()
}

// GetSpentDurationSecondsOk returns a tuple with the SpentDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyQuest) GetSpentDurationSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpentDurationSeconds.Get(), o.SpentDurationSeconds.IsSet()
}

// HasSpentDurationSeconds returns a boolean if a field has been set.
func (o *DailyQuest) HasSpentDurationSeconds() bool {
	if o != nil && o.SpentDurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetSpentDurationSeconds gets a reference to the given NullableInt32 and assigns it to the SpentDurationSeconds field.
func (o *DailyQuest) SetSpentDurationSeconds(v int32) {
	o.SpentDurationSeconds.Set(&v)
}
// SetSpentDurationSecondsNil sets the value for SpentDurationSeconds to be an explicit nil
func (o *DailyQuest) SetSpentDurationSecondsNil() {
	o.SpentDurationSeconds.Set(nil)
}

// UnsetSpentDurationSeconds ensures that no value is present for SpentDurationSeconds, not even an explicit nil
func (o *DailyQuest) UnsetSpentDurationSeconds() {
	o.SpentDurationSeconds.Unset()
}

// GetTargetDurationSeconds returns the TargetDurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DailyQuest) GetTargetDurationSeconds() int32 {
	if o == nil || IsNil(o.TargetDurationSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.TargetDurationSeconds.Get()
}

// GetTargetDurationSecondsOk returns a tuple with the TargetDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyQuest) GetTargetDurationSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDurationSeconds.Get(), o.TargetDurationSeconds.IsSet()
}

// HasTargetDurationSeconds returns a boolean if a field has been set.
func (o *DailyQuest) HasTargetDurationSeconds() bool {
	if o != nil && o.TargetDurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetTargetDurationSeconds gets a reference to the given NullableInt32 and assigns it to the TargetDurationSeconds field.
func (o *DailyQuest) SetTargetDurationSeconds(v int32) {
	o.TargetDurationSeconds.Set(&v)
}
// SetTargetDurationSecondsNil sets the value for TargetDurationSeconds to be an explicit nil
func (o *DailyQuest) SetTargetDurationSecondsNil() {
	o.TargetDurationSeconds.Set(nil)
}

// UnsetTargetDurationSeconds ensures that no value is present for TargetDurationSeconds, not even an explicit nil
func (o *DailyQuest) UnsetTargetDurationSeconds() {
	o.TargetDurationSeconds.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DailyQuest) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyQuest) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DailyQuest) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *DailyQuest) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DailyQuest) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DailyQuest) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o DailyQuest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyQuest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Calculated.IsSet() {
		toSerialize["calculated"] = o.Calculated.Get()
	}
	if o.SpentDurationSeconds.IsSet() {
		toSerialize["spent_duration_seconds"] = o.SpentDurationSeconds.Get()
	}
	if o.TargetDurationSeconds.IsSet() {
		toSerialize["target_duration_seconds"] = o.TargetDurationSeconds.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DailyQuest) UnmarshalJSON(data []byte) (err error) {
	varDailyQuest := _DailyQuest{}

	err = json.Unmarshal(data, &varDailyQuest)

	if err != nil {
		return err
	}

	*o = DailyQuest(varDailyQuest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "calculated")
		delete(additionalProperties, "spent_duration_seconds")
		delete(additionalProperties, "target_duration_seconds")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyQuest struct {
	value *DailyQuest
	isSet bool
}

func (v NullableDailyQuest) Get() *DailyQuest {
	return v.value
}

func (v *NullableDailyQuest) Set(val *DailyQuest) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyQuest) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyQuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyQuest(val *DailyQuest) *NullableDailyQuest {
	return &NullableDailyQuest{value: val, isSet: true}
}

func (v NullableDailyQuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyQuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


