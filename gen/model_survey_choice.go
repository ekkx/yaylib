
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SurveyChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveyChoice{}

// SurveyChoice struct for SurveyChoice
type SurveyChoice struct {
	Id NullableInt64 `json:"id,omitempty"`
	Label NullableString `json:"label,omitempty"`
	VotesCount NullableInt32 `json:"votes_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SurveyChoice SurveyChoice

// NewSurveyChoice instantiates a new SurveyChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveyChoice() *SurveyChoice {
	this := SurveyChoice{}
	return &this
}

// NewSurveyChoiceWithDefaults instantiates a new SurveyChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyChoiceWithDefaults() *SurveyChoice {
	this := SurveyChoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveyChoice) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveyChoice) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SurveyChoice) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *SurveyChoice) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SurveyChoice) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SurveyChoice) UnsetId() {
	o.Id.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveyChoice) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveyChoice) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *SurveyChoice) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *SurveyChoice) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *SurveyChoice) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *SurveyChoice) UnsetLabel() {
	o.Label.Unset()
}

// GetVotesCount returns the VotesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveyChoice) GetVotesCount() int32 {
	if o == nil || IsNil(o.VotesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VotesCount.Get()
}

// GetVotesCountOk returns a tuple with the VotesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveyChoice) GetVotesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VotesCount.Get(), o.VotesCount.IsSet()
}

// HasVotesCount returns a boolean if a field has been set.
func (o *SurveyChoice) HasVotesCount() bool {
	if o != nil && o.VotesCount.IsSet() {
		return true
	}

	return false
}

// SetVotesCount gets a reference to the given NullableInt32 and assigns it to the VotesCount field.
func (o *SurveyChoice) SetVotesCount(v int32) {
	o.VotesCount.Set(&v)
}
// SetVotesCountNil sets the value for VotesCount to be an explicit nil
func (o *SurveyChoice) SetVotesCountNil() {
	o.VotesCount.Set(nil)
}

// UnsetVotesCount ensures that no value is present for VotesCount, not even an explicit nil
func (o *SurveyChoice) UnsetVotesCount() {
	o.VotesCount.Unset()
}

func (o SurveyChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveyChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.VotesCount.IsSet() {
		toSerialize["votes_count"] = o.VotesCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SurveyChoice) UnmarshalJSON(data []byte) (err error) {
	varSurveyChoice := _SurveyChoice{}

	err = json.Unmarshal(data, &varSurveyChoice)

	if err != nil {
		return err
	}

	*o = SurveyChoice(varSurveyChoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "votes_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSurveyChoice struct {
	value *SurveyChoice
	isSet bool
}

func (v NullableSurveyChoice) Get() *SurveyChoice {
	return v.value
}

func (v *NullableSurveyChoice) Set(val *SurveyChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveyChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveyChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveyChoice(val *SurveyChoice) *NullableSurveyChoice {
	return &NullableSurveyChoice{value: val, isSet: true}
}

func (v NullableSurveyChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveyChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


