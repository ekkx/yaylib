
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletUserStatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletUserStatusDetails{}

// Web3WalletUserStatusDetails struct for Web3WalletUserStatusDetails
type Web3WalletUserStatusDetails struct {
	ActivityRank NullableModelUserRank `json:"activity_rank,omitempty"`
	IsRankCalculated NullableBool `json:"is_rank_calculated,omitempty"`
	IsScoreCalculated NullableBool `json:"is_score_calculated,omitempty"`
	Score NullableInt32 `json:"score,omitempty"`
	ScoreMaxFormat NullableString `json:"score_max_format,omitempty"`
	UserRank NullableModelUserRank `json:"user_rank,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletUserStatusDetails Web3WalletUserStatusDetails

// NewWeb3WalletUserStatusDetails instantiates a new Web3WalletUserStatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletUserStatusDetails() *Web3WalletUserStatusDetails {
	this := Web3WalletUserStatusDetails{}
	return &this
}

// NewWeb3WalletUserStatusDetailsWithDefaults instantiates a new Web3WalletUserStatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletUserStatusDetailsWithDefaults() *Web3WalletUserStatusDetails {
	this := Web3WalletUserStatusDetails{}
	return &this
}

// GetActivityRank returns the ActivityRank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetActivityRank() ModelUserRank {
	if o == nil || IsNil(o.ActivityRank.Get()) {
		var ret ModelUserRank
		return ret
	}
	return *o.ActivityRank.Get()
}

// GetActivityRankOk returns a tuple with the ActivityRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetActivityRankOk() (*ModelUserRank, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityRank.Get(), o.ActivityRank.IsSet()
}

// HasActivityRank returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasActivityRank() bool {
	if o != nil && o.ActivityRank.IsSet() {
		return true
	}

	return false
}

// SetActivityRank gets a reference to the given NullableModelUserRank and assigns it to the ActivityRank field.
func (o *Web3WalletUserStatusDetails) SetActivityRank(v ModelUserRank) {
	o.ActivityRank.Set(&v)
}
// SetActivityRankNil sets the value for ActivityRank to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetActivityRankNil() {
	o.ActivityRank.Set(nil)
}

// UnsetActivityRank ensures that no value is present for ActivityRank, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetActivityRank() {
	o.ActivityRank.Unset()
}

// GetIsRankCalculated returns the IsRankCalculated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetIsRankCalculated() bool {
	if o == nil || IsNil(o.IsRankCalculated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRankCalculated.Get()
}

// GetIsRankCalculatedOk returns a tuple with the IsRankCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetIsRankCalculatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRankCalculated.Get(), o.IsRankCalculated.IsSet()
}

// HasIsRankCalculated returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasIsRankCalculated() bool {
	if o != nil && o.IsRankCalculated.IsSet() {
		return true
	}

	return false
}

// SetIsRankCalculated gets a reference to the given NullableBool and assigns it to the IsRankCalculated field.
func (o *Web3WalletUserStatusDetails) SetIsRankCalculated(v bool) {
	o.IsRankCalculated.Set(&v)
}
// SetIsRankCalculatedNil sets the value for IsRankCalculated to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetIsRankCalculatedNil() {
	o.IsRankCalculated.Set(nil)
}

// UnsetIsRankCalculated ensures that no value is present for IsRankCalculated, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetIsRankCalculated() {
	o.IsRankCalculated.Unset()
}

// GetIsScoreCalculated returns the IsScoreCalculated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetIsScoreCalculated() bool {
	if o == nil || IsNil(o.IsScoreCalculated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsScoreCalculated.Get()
}

// GetIsScoreCalculatedOk returns a tuple with the IsScoreCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetIsScoreCalculatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsScoreCalculated.Get(), o.IsScoreCalculated.IsSet()
}

// HasIsScoreCalculated returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasIsScoreCalculated() bool {
	if o != nil && o.IsScoreCalculated.IsSet() {
		return true
	}

	return false
}

// SetIsScoreCalculated gets a reference to the given NullableBool and assigns it to the IsScoreCalculated field.
func (o *Web3WalletUserStatusDetails) SetIsScoreCalculated(v bool) {
	o.IsScoreCalculated.Set(&v)
}
// SetIsScoreCalculatedNil sets the value for IsScoreCalculated to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetIsScoreCalculatedNil() {
	o.IsScoreCalculated.Set(nil)
}

// UnsetIsScoreCalculated ensures that no value is present for IsScoreCalculated, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetIsScoreCalculated() {
	o.IsScoreCalculated.Unset()
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetScore() int32 {
	if o == nil || IsNil(o.Score.Get()) {
		var ret int32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableInt32 and assigns it to the Score field.
func (o *Web3WalletUserStatusDetails) SetScore(v int32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetScore() {
	o.Score.Unset()
}

// GetScoreMaxFormat returns the ScoreMaxFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetScoreMaxFormat() string {
	if o == nil || IsNil(o.ScoreMaxFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ScoreMaxFormat.Get()
}

// GetScoreMaxFormatOk returns a tuple with the ScoreMaxFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetScoreMaxFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScoreMaxFormat.Get(), o.ScoreMaxFormat.IsSet()
}

// HasScoreMaxFormat returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasScoreMaxFormat() bool {
	if o != nil && o.ScoreMaxFormat.IsSet() {
		return true
	}

	return false
}

// SetScoreMaxFormat gets a reference to the given NullableString and assigns it to the ScoreMaxFormat field.
func (o *Web3WalletUserStatusDetails) SetScoreMaxFormat(v string) {
	o.ScoreMaxFormat.Set(&v)
}
// SetScoreMaxFormatNil sets the value for ScoreMaxFormat to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetScoreMaxFormatNil() {
	o.ScoreMaxFormat.Set(nil)
}

// UnsetScoreMaxFormat ensures that no value is present for ScoreMaxFormat, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetScoreMaxFormat() {
	o.ScoreMaxFormat.Unset()
}

// GetUserRank returns the UserRank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletUserStatusDetails) GetUserRank() ModelUserRank {
	if o == nil || IsNil(o.UserRank.Get()) {
		var ret ModelUserRank
		return ret
	}
	return *o.UserRank.Get()
}

// GetUserRankOk returns a tuple with the UserRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletUserStatusDetails) GetUserRankOk() (*ModelUserRank, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserRank.Get(), o.UserRank.IsSet()
}

// HasUserRank returns a boolean if a field has been set.
func (o *Web3WalletUserStatusDetails) HasUserRank() bool {
	if o != nil && o.UserRank.IsSet() {
		return true
	}

	return false
}

// SetUserRank gets a reference to the given NullableModelUserRank and assigns it to the UserRank field.
func (o *Web3WalletUserStatusDetails) SetUserRank(v ModelUserRank) {
	o.UserRank.Set(&v)
}
// SetUserRankNil sets the value for UserRank to be an explicit nil
func (o *Web3WalletUserStatusDetails) SetUserRankNil() {
	o.UserRank.Set(nil)
}

// UnsetUserRank ensures that no value is present for UserRank, not even an explicit nil
func (o *Web3WalletUserStatusDetails) UnsetUserRank() {
	o.UserRank.Unset()
}

func (o Web3WalletUserStatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletUserStatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivityRank.IsSet() {
		toSerialize["activity_rank"] = o.ActivityRank.Get()
	}
	if o.IsRankCalculated.IsSet() {
		toSerialize["is_rank_calculated"] = o.IsRankCalculated.Get()
	}
	if o.IsScoreCalculated.IsSet() {
		toSerialize["is_score_calculated"] = o.IsScoreCalculated.Get()
	}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.ScoreMaxFormat.IsSet() {
		toSerialize["score_max_format"] = o.ScoreMaxFormat.Get()
	}
	if o.UserRank.IsSet() {
		toSerialize["user_rank"] = o.UserRank.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletUserStatusDetails) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletUserStatusDetails := _Web3WalletUserStatusDetails{}

	err = json.Unmarshal(data, &varWeb3WalletUserStatusDetails)

	if err != nil {
		return err
	}

	*o = Web3WalletUserStatusDetails(varWeb3WalletUserStatusDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity_rank")
		delete(additionalProperties, "is_rank_calculated")
		delete(additionalProperties, "is_score_calculated")
		delete(additionalProperties, "score")
		delete(additionalProperties, "score_max_format")
		delete(additionalProperties, "user_rank")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletUserStatusDetails struct {
	value *Web3WalletUserStatusDetails
	isSet bool
}

func (v NullableWeb3WalletUserStatusDetails) Get() *Web3WalletUserStatusDetails {
	return v.value
}

func (v *NullableWeb3WalletUserStatusDetails) Set(val *Web3WalletUserStatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletUserStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletUserStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletUserStatusDetails(val *Web3WalletUserStatusDetails) *NullableWeb3WalletUserStatusDetails {
	return &NullableWeb3WalletUserStatusDetails{value: val, isSet: true}
}

func (v NullableWeb3WalletUserStatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletUserStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


