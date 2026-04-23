
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Survey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Survey{}

// Survey struct for Survey
type Survey struct {
	Choices []Choice `json:"choices,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Voted NullableBool `json:"voted,omitempty"`
	VotesCount NullableInt32 `json:"votes_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Survey Survey

// NewSurvey instantiates a new Survey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurvey() *Survey {
	this := Survey{}
	return &this
}

// NewSurveyWithDefaults instantiates a new Survey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyWithDefaults() *Survey {
	this := Survey{}
	return &this
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetChoices() []Choice {
	if o == nil {
		var ret []Choice
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetChoicesOk() ([]Choice, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *Survey) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []Choice and assigns it to the Choices field.
func (o *Survey) SetChoices(v []Choice) {
	o.Choices = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Survey) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Survey) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Survey) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Survey) UnsetId() {
	o.Id.Unset()
}

// GetVoted returns the Voted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetVoted() bool {
	if o == nil || IsNil(o.Voted.Get()) {
		var ret bool
		return ret
	}
	return *o.Voted.Get()
}

// GetVotedOk returns a tuple with the Voted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetVotedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Voted.Get(), o.Voted.IsSet()
}

// HasVoted returns a boolean if a field has been set.
func (o *Survey) HasVoted() bool {
	if o != nil && o.Voted.IsSet() {
		return true
	}

	return false
}

// SetVoted gets a reference to the given NullableBool and assigns it to the Voted field.
func (o *Survey) SetVoted(v bool) {
	o.Voted.Set(&v)
}
// SetVotedNil sets the value for Voted to be an explicit nil
func (o *Survey) SetVotedNil() {
	o.Voted.Set(nil)
}

// UnsetVoted ensures that no value is present for Voted, not even an explicit nil
func (o *Survey) UnsetVoted() {
	o.Voted.Unset()
}

// GetVotesCount returns the VotesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetVotesCount() int32 {
	if o == nil || IsNil(o.VotesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VotesCount.Get()
}

// GetVotesCountOk returns a tuple with the VotesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetVotesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VotesCount.Get(), o.VotesCount.IsSet()
}

// HasVotesCount returns a boolean if a field has been set.
func (o *Survey) HasVotesCount() bool {
	if o != nil && o.VotesCount.IsSet() {
		return true
	}

	return false
}

// SetVotesCount gets a reference to the given NullableInt32 and assigns it to the VotesCount field.
func (o *Survey) SetVotesCount(v int32) {
	o.VotesCount.Set(&v)
}
// SetVotesCountNil sets the value for VotesCount to be an explicit nil
func (o *Survey) SetVotesCountNil() {
	o.VotesCount.Set(nil)
}

// UnsetVotesCount ensures that no value is present for VotesCount, not even an explicit nil
func (o *Survey) UnsetVotesCount() {
	o.VotesCount.Unset()
}

func (o Survey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Survey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Voted.IsSet() {
		toSerialize["voted"] = o.Voted.Get()
	}
	if o.VotesCount.IsSet() {
		toSerialize["votes_count"] = o.VotesCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Survey) UnmarshalJSON(data []byte) (err error) {
	varSurvey := _Survey{}

	err = json.Unmarshal(data, &varSurvey)

	if err != nil {
		return err
	}

	*o = Survey(varSurvey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "choices")
		delete(additionalProperties, "id")
		delete(additionalProperties, "voted")
		delete(additionalProperties, "votes_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSurvey struct {
	value *Survey
	isSet bool
}

func (v NullableSurvey) Get() *Survey {
	return v.value
}

func (v *NullableSurvey) Set(val *Survey) {
	v.value = val
	v.isSet = true
}

func (v NullableSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurvey(val *Survey) *NullableSurvey {
	return &NullableSurvey{value: val, isSet: true}
}

func (v NullableSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


