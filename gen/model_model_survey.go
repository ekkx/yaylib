
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelSurvey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSurvey{}

// ModelSurvey struct for ModelSurvey
type ModelSurvey struct {
	Choices []SurveyChoice `json:"choices,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsVoted NullableBool `json:"is_voted,omitempty"`
	VotesCount NullableInt32 `json:"votes_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelSurvey ModelSurvey

// NewModelSurvey instantiates a new ModelSurvey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSurvey() *ModelSurvey {
	this := ModelSurvey{}
	return &this
}

// NewModelSurveyWithDefaults instantiates a new ModelSurvey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSurveyWithDefaults() *ModelSurvey {
	this := ModelSurvey{}
	return &this
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSurvey) GetChoices() []SurveyChoice {
	if o == nil {
		var ret []SurveyChoice
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSurvey) GetChoicesOk() ([]SurveyChoice, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *ModelSurvey) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []SurveyChoice and assigns it to the Choices field.
func (o *ModelSurvey) SetChoices(v []SurveyChoice) {
	o.Choices = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSurvey) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSurvey) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelSurvey) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelSurvey) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelSurvey) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelSurvey) UnsetId() {
	o.Id.Unset()
}

// GetIsVoted returns the IsVoted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSurvey) GetIsVoted() bool {
	if o == nil || IsNil(o.IsVoted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsVoted.Get()
}

// GetIsVotedOk returns a tuple with the IsVoted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSurvey) GetIsVotedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsVoted.Get(), o.IsVoted.IsSet()
}

// HasIsVoted returns a boolean if a field has been set.
func (o *ModelSurvey) HasIsVoted() bool {
	if o != nil && o.IsVoted.IsSet() {
		return true
	}

	return false
}

// SetIsVoted gets a reference to the given NullableBool and assigns it to the IsVoted field.
func (o *ModelSurvey) SetIsVoted(v bool) {
	o.IsVoted.Set(&v)
}
// SetIsVotedNil sets the value for IsVoted to be an explicit nil
func (o *ModelSurvey) SetIsVotedNil() {
	o.IsVoted.Set(nil)
}

// UnsetIsVoted ensures that no value is present for IsVoted, not even an explicit nil
func (o *ModelSurvey) UnsetIsVoted() {
	o.IsVoted.Unset()
}

// GetVotesCount returns the VotesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSurvey) GetVotesCount() int32 {
	if o == nil || IsNil(o.VotesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VotesCount.Get()
}

// GetVotesCountOk returns a tuple with the VotesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSurvey) GetVotesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VotesCount.Get(), o.VotesCount.IsSet()
}

// HasVotesCount returns a boolean if a field has been set.
func (o *ModelSurvey) HasVotesCount() bool {
	if o != nil && o.VotesCount.IsSet() {
		return true
	}

	return false
}

// SetVotesCount gets a reference to the given NullableInt32 and assigns it to the VotesCount field.
func (o *ModelSurvey) SetVotesCount(v int32) {
	o.VotesCount.Set(&v)
}
// SetVotesCountNil sets the value for VotesCount to be an explicit nil
func (o *ModelSurvey) SetVotesCountNil() {
	o.VotesCount.Set(nil)
}

// UnsetVotesCount ensures that no value is present for VotesCount, not even an explicit nil
func (o *ModelSurvey) UnsetVotesCount() {
	o.VotesCount.Unset()
}

func (o ModelSurvey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSurvey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsVoted.IsSet() {
		toSerialize["is_voted"] = o.IsVoted.Get()
	}
	if o.VotesCount.IsSet() {
		toSerialize["votes_count"] = o.VotesCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelSurvey) UnmarshalJSON(data []byte) (err error) {
	varModelSurvey := _ModelSurvey{}

	err = json.Unmarshal(data, &varModelSurvey)

	if err != nil {
		return err
	}

	*o = ModelSurvey(varModelSurvey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "choices")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_voted")
		delete(additionalProperties, "votes_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelSurvey struct {
	value *ModelSurvey
	isSet bool
}

func (v NullableModelSurvey) Get() *ModelSurvey {
	return v.value
}

func (v *NullableModelSurvey) Set(val *ModelSurvey) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSurvey(val *ModelSurvey) *NullableModelSurvey {
	return &NullableModelSurvey{value: val, isSet: true}
}

func (v NullableModelSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


