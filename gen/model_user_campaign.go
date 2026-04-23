
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCampaign{}

// UserCampaign struct for UserCampaign
type UserCampaign struct {
	Multiplier NullableFloat64 `json:"multiplier,omitempty"`
	MultiplierBreakdown NullableMultiplierBreakdown `json:"multiplier_breakdown,omitempty"`
	Points NullableInt32 `json:"points,omitempty"`
	Ranking NullableInt32 `json:"ranking,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCampaign UserCampaign

// NewUserCampaign instantiates a new UserCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCampaign() *UserCampaign {
	this := UserCampaign{}
	return &this
}

// NewUserCampaignWithDefaults instantiates a new UserCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCampaignWithDefaults() *UserCampaign {
	this := UserCampaign{}
	return &this
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCampaign) GetMultiplier() float64 {
	if o == nil || IsNil(o.Multiplier.Get()) {
		var ret float64
		return ret
	}
	return *o.Multiplier.Get()
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCampaign) GetMultiplierOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Multiplier.Get(), o.Multiplier.IsSet()
}

// HasMultiplier returns a boolean if a field has been set.
func (o *UserCampaign) HasMultiplier() bool {
	if o != nil && o.Multiplier.IsSet() {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given NullableFloat64 and assigns it to the Multiplier field.
func (o *UserCampaign) SetMultiplier(v float64) {
	o.Multiplier.Set(&v)
}
// SetMultiplierNil sets the value for Multiplier to be an explicit nil
func (o *UserCampaign) SetMultiplierNil() {
	o.Multiplier.Set(nil)
}

// UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
func (o *UserCampaign) UnsetMultiplier() {
	o.Multiplier.Unset()
}

// GetMultiplierBreakdown returns the MultiplierBreakdown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCampaign) GetMultiplierBreakdown() MultiplierBreakdown {
	if o == nil || IsNil(o.MultiplierBreakdown.Get()) {
		var ret MultiplierBreakdown
		return ret
	}
	return *o.MultiplierBreakdown.Get()
}

// GetMultiplierBreakdownOk returns a tuple with the MultiplierBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCampaign) GetMultiplierBreakdownOk() (*MultiplierBreakdown, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultiplierBreakdown.Get(), o.MultiplierBreakdown.IsSet()
}

// HasMultiplierBreakdown returns a boolean if a field has been set.
func (o *UserCampaign) HasMultiplierBreakdown() bool {
	if o != nil && o.MultiplierBreakdown.IsSet() {
		return true
	}

	return false
}

// SetMultiplierBreakdown gets a reference to the given NullableMultiplierBreakdown and assigns it to the MultiplierBreakdown field.
func (o *UserCampaign) SetMultiplierBreakdown(v MultiplierBreakdown) {
	o.MultiplierBreakdown.Set(&v)
}
// SetMultiplierBreakdownNil sets the value for MultiplierBreakdown to be an explicit nil
func (o *UserCampaign) SetMultiplierBreakdownNil() {
	o.MultiplierBreakdown.Set(nil)
}

// UnsetMultiplierBreakdown ensures that no value is present for MultiplierBreakdown, not even an explicit nil
func (o *UserCampaign) UnsetMultiplierBreakdown() {
	o.MultiplierBreakdown.Unset()
}

// GetPoints returns the Points field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCampaign) GetPoints() int32 {
	if o == nil || IsNil(o.Points.Get()) {
		var ret int32
		return ret
	}
	return *o.Points.Get()
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCampaign) GetPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Points.Get(), o.Points.IsSet()
}

// HasPoints returns a boolean if a field has been set.
func (o *UserCampaign) HasPoints() bool {
	if o != nil && o.Points.IsSet() {
		return true
	}

	return false
}

// SetPoints gets a reference to the given NullableInt32 and assigns it to the Points field.
func (o *UserCampaign) SetPoints(v int32) {
	o.Points.Set(&v)
}
// SetPointsNil sets the value for Points to be an explicit nil
func (o *UserCampaign) SetPointsNil() {
	o.Points.Set(nil)
}

// UnsetPoints ensures that no value is present for Points, not even an explicit nil
func (o *UserCampaign) UnsetPoints() {
	o.Points.Unset()
}

// GetRanking returns the Ranking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCampaign) GetRanking() int32 {
	if o == nil || IsNil(o.Ranking.Get()) {
		var ret int32
		return ret
	}
	return *o.Ranking.Get()
}

// GetRankingOk returns a tuple with the Ranking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCampaign) GetRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ranking.Get(), o.Ranking.IsSet()
}

// HasRanking returns a boolean if a field has been set.
func (o *UserCampaign) HasRanking() bool {
	if o != nil && o.Ranking.IsSet() {
		return true
	}

	return false
}

// SetRanking gets a reference to the given NullableInt32 and assigns it to the Ranking field.
func (o *UserCampaign) SetRanking(v int32) {
	o.Ranking.Set(&v)
}
// SetRankingNil sets the value for Ranking to be an explicit nil
func (o *UserCampaign) SetRankingNil() {
	o.Ranking.Set(nil)
}

// UnsetRanking ensures that no value is present for Ranking, not even an explicit nil
func (o *UserCampaign) UnsetRanking() {
	o.Ranking.Unset()
}

func (o UserCampaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Multiplier.IsSet() {
		toSerialize["multiplier"] = o.Multiplier.Get()
	}
	if o.MultiplierBreakdown.IsSet() {
		toSerialize["multiplier_breakdown"] = o.MultiplierBreakdown.Get()
	}
	if o.Points.IsSet() {
		toSerialize["points"] = o.Points.Get()
	}
	if o.Ranking.IsSet() {
		toSerialize["ranking"] = o.Ranking.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCampaign) UnmarshalJSON(data []byte) (err error) {
	varUserCampaign := _UserCampaign{}

	err = json.Unmarshal(data, &varUserCampaign)

	if err != nil {
		return err
	}

	*o = UserCampaign(varUserCampaign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "multiplier")
		delete(additionalProperties, "multiplier_breakdown")
		delete(additionalProperties, "points")
		delete(additionalProperties, "ranking")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCampaign struct {
	value *UserCampaign
	isSet bool
}

func (v NullableUserCampaign) Get() *UserCampaign {
	return v.value
}

func (v *NullableUserCampaign) Set(val *UserCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCampaign(val *UserCampaign) *NullableUserCampaign {
	return &NullableUserCampaign{value: val, isSet: true}
}

func (v NullableUserCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


