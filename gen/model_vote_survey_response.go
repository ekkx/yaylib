
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the VoteSurveyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoteSurveyResponse{}

// VoteSurveyResponse struct for VoteSurveyResponse
type VoteSurveyResponse struct {
	Survey NullableSurvey `json:"survey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VoteSurveyResponse VoteSurveyResponse

// NewVoteSurveyResponse instantiates a new VoteSurveyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoteSurveyResponse() *VoteSurveyResponse {
	this := VoteSurveyResponse{}
	return &this
}

// NewVoteSurveyResponseWithDefaults instantiates a new VoteSurveyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoteSurveyResponseWithDefaults() *VoteSurveyResponse {
	this := VoteSurveyResponse{}
	return &this
}

// GetSurvey returns the Survey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoteSurveyResponse) GetSurvey() Survey {
	if o == nil || IsNil(o.Survey.Get()) {
		var ret Survey
		return ret
	}
	return *o.Survey.Get()
}

// GetSurveyOk returns a tuple with the Survey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoteSurveyResponse) GetSurveyOk() (*Survey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Survey.Get(), o.Survey.IsSet()
}

// HasSurvey returns a boolean if a field has been set.
func (o *VoteSurveyResponse) HasSurvey() bool {
	if o != nil && o.Survey.IsSet() {
		return true
	}

	return false
}

// SetSurvey gets a reference to the given NullableSurvey and assigns it to the Survey field.
func (o *VoteSurveyResponse) SetSurvey(v Survey) {
	o.Survey.Set(&v)
}
// SetSurveyNil sets the value for Survey to be an explicit nil
func (o *VoteSurveyResponse) SetSurveyNil() {
	o.Survey.Set(nil)
}

// UnsetSurvey ensures that no value is present for Survey, not even an explicit nil
func (o *VoteSurveyResponse) UnsetSurvey() {
	o.Survey.Unset()
}

func (o VoteSurveyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoteSurveyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Survey.IsSet() {
		toSerialize["survey"] = o.Survey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VoteSurveyResponse) UnmarshalJSON(data []byte) (err error) {
	varVoteSurveyResponse := _VoteSurveyResponse{}

	err = json.Unmarshal(data, &varVoteSurveyResponse)

	if err != nil {
		return err
	}

	*o = VoteSurveyResponse(varVoteSurveyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "survey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVoteSurveyResponse struct {
	value *VoteSurveyResponse
	isSet bool
}

func (v NullableVoteSurveyResponse) Get() *VoteSurveyResponse {
	return v.value
}

func (v *NullableVoteSurveyResponse) Set(val *VoteSurveyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoteSurveyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoteSurveyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoteSurveyResponse(val *VoteSurveyResponse) *NullableVoteSurveyResponse {
	return &NullableVoteSurveyResponse{value: val, isSet: true}
}

func (v NullableVoteSurveyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoteSurveyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


