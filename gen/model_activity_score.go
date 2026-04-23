
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ActivityScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityScore{}

// ActivityScore struct for ActivityScore
type ActivityScore struct {
	Calculated NullableBool `json:"calculated,omitempty"`
	Grade NullableModelUserRank `json:"grade,omitempty"`
	Score NullableFloat64 `json:"score,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActivityScore ActivityScore

// NewActivityScore instantiates a new ActivityScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityScore() *ActivityScore {
	this := ActivityScore{}
	return &this
}

// NewActivityScoreWithDefaults instantiates a new ActivityScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityScoreWithDefaults() *ActivityScore {
	this := ActivityScore{}
	return &this
}

// GetCalculated returns the Calculated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityScore) GetCalculated() bool {
	if o == nil || IsNil(o.Calculated.Get()) {
		var ret bool
		return ret
	}
	return *o.Calculated.Get()
}

// GetCalculatedOk returns a tuple with the Calculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityScore) GetCalculatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Calculated.Get(), o.Calculated.IsSet()
}

// HasCalculated returns a boolean if a field has been set.
func (o *ActivityScore) HasCalculated() bool {
	if o != nil && o.Calculated.IsSet() {
		return true
	}

	return false
}

// SetCalculated gets a reference to the given NullableBool and assigns it to the Calculated field.
func (o *ActivityScore) SetCalculated(v bool) {
	o.Calculated.Set(&v)
}
// SetCalculatedNil sets the value for Calculated to be an explicit nil
func (o *ActivityScore) SetCalculatedNil() {
	o.Calculated.Set(nil)
}

// UnsetCalculated ensures that no value is present for Calculated, not even an explicit nil
func (o *ActivityScore) UnsetCalculated() {
	o.Calculated.Unset()
}

// GetGrade returns the Grade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityScore) GetGrade() ModelUserRank {
	if o == nil || IsNil(o.Grade.Get()) {
		var ret ModelUserRank
		return ret
	}
	return *o.Grade.Get()
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityScore) GetGradeOk() (*ModelUserRank, bool) {
	if o == nil {
		return nil, false
	}
	return o.Grade.Get(), o.Grade.IsSet()
}

// HasGrade returns a boolean if a field has been set.
func (o *ActivityScore) HasGrade() bool {
	if o != nil && o.Grade.IsSet() {
		return true
	}

	return false
}

// SetGrade gets a reference to the given NullableModelUserRank and assigns it to the Grade field.
func (o *ActivityScore) SetGrade(v ModelUserRank) {
	o.Grade.Set(&v)
}
// SetGradeNil sets the value for Grade to be an explicit nil
func (o *ActivityScore) SetGradeNil() {
	o.Grade.Set(nil)
}

// UnsetGrade ensures that no value is present for Grade, not even an explicit nil
func (o *ActivityScore) UnsetGrade() {
	o.Grade.Unset()
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityScore) GetScore() float64 {
	if o == nil || IsNil(o.Score.Get()) {
		var ret float64
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityScore) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *ActivityScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat64 and assigns it to the Score field.
func (o *ActivityScore) SetScore(v float64) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *ActivityScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *ActivityScore) UnsetScore() {
	o.Score.Unset()
}

func (o ActivityScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Calculated.IsSet() {
		toSerialize["calculated"] = o.Calculated.Get()
	}
	if o.Grade.IsSet() {
		toSerialize["grade"] = o.Grade.Get()
	}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ActivityScore) UnmarshalJSON(data []byte) (err error) {
	varActivityScore := _ActivityScore{}

	err = json.Unmarshal(data, &varActivityScore)

	if err != nil {
		return err
	}

	*o = ActivityScore(varActivityScore)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "calculated")
		delete(additionalProperties, "grade")
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActivityScore struct {
	value *ActivityScore
	isSet bool
}

func (v NullableActivityScore) Get() *ActivityScore {
	return v.value
}

func (v *NullableActivityScore) Set(val *ActivityScore) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityScore) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityScore(val *ActivityScore) *NullableActivityScore {
	return &NullableActivityScore{value: val, isSet: true}
}

func (v NullableActivityScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


