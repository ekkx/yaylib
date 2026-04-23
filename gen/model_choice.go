
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Choice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Choice{}

// Choice struct for Choice
type Choice struct {
	Id NullableInt64 `json:"id,omitempty"`
	Label NullableString `json:"label,omitempty"`
	VotesCount NullableInt32 `json:"votes_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Choice Choice

// NewChoice instantiates a new Choice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChoice() *Choice {
	this := Choice{}
	return &this
}

// NewChoiceWithDefaults instantiates a new Choice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChoiceWithDefaults() *Choice {
	this := Choice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Choice) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Choice) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Choice) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Choice) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Choice) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Choice) UnsetId() {
	o.Id.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Choice) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Choice) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *Choice) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *Choice) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *Choice) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *Choice) UnsetLabel() {
	o.Label.Unset()
}

// GetVotesCount returns the VotesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Choice) GetVotesCount() int32 {
	if o == nil || IsNil(o.VotesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VotesCount.Get()
}

// GetVotesCountOk returns a tuple with the VotesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Choice) GetVotesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VotesCount.Get(), o.VotesCount.IsSet()
}

// HasVotesCount returns a boolean if a field has been set.
func (o *Choice) HasVotesCount() bool {
	if o != nil && o.VotesCount.IsSet() {
		return true
	}

	return false
}

// SetVotesCount gets a reference to the given NullableInt32 and assigns it to the VotesCount field.
func (o *Choice) SetVotesCount(v int32) {
	o.VotesCount.Set(&v)
}
// SetVotesCountNil sets the value for VotesCount to be an explicit nil
func (o *Choice) SetVotesCountNil() {
	o.VotesCount.Set(nil)
}

// UnsetVotesCount ensures that no value is present for VotesCount, not even an explicit nil
func (o *Choice) UnsetVotesCount() {
	o.VotesCount.Unset()
}

func (o Choice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Choice) ToMap() (map[string]interface{}, error) {
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

func (o *Choice) UnmarshalJSON(data []byte) (err error) {
	varChoice := _Choice{}

	err = json.Unmarshal(data, &varChoice)

	if err != nil {
		return err
	}

	*o = Choice(varChoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "votes_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChoice struct {
	value *Choice
	isSet bool
}

func (v NullableChoice) Get() *Choice {
	return v.value
}

func (v *NullableChoice) Set(val *Choice) {
	v.value = val
	v.isSet = true
}

func (v NullableChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChoice(val *Choice) *NullableChoice {
	return &NullableChoice{value: val, isSet: true}
}

func (v NullableChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


