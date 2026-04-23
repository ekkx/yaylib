
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserReputations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserReputations{}

// UserReputations struct for UserReputations
type UserReputations struct {
	ActivityScore NullableActivityScore `json:"activity_score,omitempty"`
	TrustScore NullableTrustScore `json:"trust_score,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserReputations UserReputations

// NewUserReputations instantiates a new UserReputations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserReputations() *UserReputations {
	this := UserReputations{}
	return &this
}

// NewUserReputationsWithDefaults instantiates a new UserReputations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserReputationsWithDefaults() *UserReputations {
	this := UserReputations{}
	return &this
}

// GetActivityScore returns the ActivityScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserReputations) GetActivityScore() ActivityScore {
	if o == nil || IsNil(o.ActivityScore.Get()) {
		var ret ActivityScore
		return ret
	}
	return *o.ActivityScore.Get()
}

// GetActivityScoreOk returns a tuple with the ActivityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReputations) GetActivityScoreOk() (*ActivityScore, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityScore.Get(), o.ActivityScore.IsSet()
}

// HasActivityScore returns a boolean if a field has been set.
func (o *UserReputations) HasActivityScore() bool {
	if o != nil && o.ActivityScore.IsSet() {
		return true
	}

	return false
}

// SetActivityScore gets a reference to the given NullableActivityScore and assigns it to the ActivityScore field.
func (o *UserReputations) SetActivityScore(v ActivityScore) {
	o.ActivityScore.Set(&v)
}
// SetActivityScoreNil sets the value for ActivityScore to be an explicit nil
func (o *UserReputations) SetActivityScoreNil() {
	o.ActivityScore.Set(nil)
}

// UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil
func (o *UserReputations) UnsetActivityScore() {
	o.ActivityScore.Unset()
}

// GetTrustScore returns the TrustScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserReputations) GetTrustScore() TrustScore {
	if o == nil || IsNil(o.TrustScore.Get()) {
		var ret TrustScore
		return ret
	}
	return *o.TrustScore.Get()
}

// GetTrustScoreOk returns a tuple with the TrustScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReputations) GetTrustScoreOk() (*TrustScore, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustScore.Get(), o.TrustScore.IsSet()
}

// HasTrustScore returns a boolean if a field has been set.
func (o *UserReputations) HasTrustScore() bool {
	if o != nil && o.TrustScore.IsSet() {
		return true
	}

	return false
}

// SetTrustScore gets a reference to the given NullableTrustScore and assigns it to the TrustScore field.
func (o *UserReputations) SetTrustScore(v TrustScore) {
	o.TrustScore.Set(&v)
}
// SetTrustScoreNil sets the value for TrustScore to be an explicit nil
func (o *UserReputations) SetTrustScoreNil() {
	o.TrustScore.Set(nil)
}

// UnsetTrustScore ensures that no value is present for TrustScore, not even an explicit nil
func (o *UserReputations) UnsetTrustScore() {
	o.TrustScore.Unset()
}

func (o UserReputations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserReputations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivityScore.IsSet() {
		toSerialize["activity_score"] = o.ActivityScore.Get()
	}
	if o.TrustScore.IsSet() {
		toSerialize["trust_score"] = o.TrustScore.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserReputations) UnmarshalJSON(data []byte) (err error) {
	varUserReputations := _UserReputations{}

	err = json.Unmarshal(data, &varUserReputations)

	if err != nil {
		return err
	}

	*o = UserReputations(varUserReputations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity_score")
		delete(additionalProperties, "trust_score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserReputations struct {
	value *UserReputations
	isSet bool
}

func (v NullableUserReputations) Get() *UserReputations {
	return v.value
}

func (v *NullableUserReputations) Set(val *UserReputations) {
	v.value = val
	v.isSet = true
}

func (v NullableUserReputations) IsSet() bool {
	return v.isSet
}

func (v *NullableUserReputations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserReputations(val *UserReputations) *NullableUserReputations {
	return &NullableUserReputations{value: val, isSet: true}
}

func (v NullableUserReputations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserReputations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


