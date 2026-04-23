
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RaceHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RaceHistory{}

// RaceHistory struct for RaceHistory
type RaceHistory struct {
	AnimationUrl NullableString `json:"animation_url,omitempty"`
	League NullablePalRaceLeagueDTO `json:"league,omitempty"`
	LeagueName NullableString `json:"league_name,omitempty"`
	Pals []Pal `json:"pals,omitempty"`
	RaceDate NullableInt64 `json:"race_date,omitempty"`
	RaceEntryFee NullableInt32 `json:"race_entry_fee,omitempty"`
	RaceId NullableString `json:"race_id,omitempty"`
	RaceReward NullableRaceReward `json:"race_reward,omitempty"`
	RaceStartAt NullableInt64 `json:"race_start_at,omitempty"`
	RaceStatus NullableString `json:"race_status,omitempty"`
	Results []PalRank `json:"results,omitempty"`
	Reward NullableReward `json:"reward,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RaceHistory RaceHistory

// NewRaceHistory instantiates a new RaceHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRaceHistory() *RaceHistory {
	this := RaceHistory{}
	return &this
}

// NewRaceHistoryWithDefaults instantiates a new RaceHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRaceHistoryWithDefaults() *RaceHistory {
	this := RaceHistory{}
	return &this
}

// GetAnimationUrl returns the AnimationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetAnimationUrl() string {
	if o == nil || IsNil(o.AnimationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AnimationUrl.Get()
}

// GetAnimationUrlOk returns a tuple with the AnimationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetAnimationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnimationUrl.Get(), o.AnimationUrl.IsSet()
}

// HasAnimationUrl returns a boolean if a field has been set.
func (o *RaceHistory) HasAnimationUrl() bool {
	if o != nil && o.AnimationUrl.IsSet() {
		return true
	}

	return false
}

// SetAnimationUrl gets a reference to the given NullableString and assigns it to the AnimationUrl field.
func (o *RaceHistory) SetAnimationUrl(v string) {
	o.AnimationUrl.Set(&v)
}
// SetAnimationUrlNil sets the value for AnimationUrl to be an explicit nil
func (o *RaceHistory) SetAnimationUrlNil() {
	o.AnimationUrl.Set(nil)
}

// UnsetAnimationUrl ensures that no value is present for AnimationUrl, not even an explicit nil
func (o *RaceHistory) UnsetAnimationUrl() {
	o.AnimationUrl.Unset()
}

// GetLeague returns the League field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetLeague() PalRaceLeagueDTO {
	if o == nil || IsNil(o.League.Get()) {
		var ret PalRaceLeagueDTO
		return ret
	}
	return *o.League.Get()
}

// GetLeagueOk returns a tuple with the League field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetLeagueOk() (*PalRaceLeagueDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.League.Get(), o.League.IsSet()
}

// HasLeague returns a boolean if a field has been set.
func (o *RaceHistory) HasLeague() bool {
	if o != nil && o.League.IsSet() {
		return true
	}

	return false
}

// SetLeague gets a reference to the given NullablePalRaceLeagueDTO and assigns it to the League field.
func (o *RaceHistory) SetLeague(v PalRaceLeagueDTO) {
	o.League.Set(&v)
}
// SetLeagueNil sets the value for League to be an explicit nil
func (o *RaceHistory) SetLeagueNil() {
	o.League.Set(nil)
}

// UnsetLeague ensures that no value is present for League, not even an explicit nil
func (o *RaceHistory) UnsetLeague() {
	o.League.Unset()
}

// GetLeagueName returns the LeagueName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetLeagueName() string {
	if o == nil || IsNil(o.LeagueName.Get()) {
		var ret string
		return ret
	}
	return *o.LeagueName.Get()
}

// GetLeagueNameOk returns a tuple with the LeagueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetLeagueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeagueName.Get(), o.LeagueName.IsSet()
}

// HasLeagueName returns a boolean if a field has been set.
func (o *RaceHistory) HasLeagueName() bool {
	if o != nil && o.LeagueName.IsSet() {
		return true
	}

	return false
}

// SetLeagueName gets a reference to the given NullableString and assigns it to the LeagueName field.
func (o *RaceHistory) SetLeagueName(v string) {
	o.LeagueName.Set(&v)
}
// SetLeagueNameNil sets the value for LeagueName to be an explicit nil
func (o *RaceHistory) SetLeagueNameNil() {
	o.LeagueName.Set(nil)
}

// UnsetLeagueName ensures that no value is present for LeagueName, not even an explicit nil
func (o *RaceHistory) UnsetLeagueName() {
	o.LeagueName.Unset()
}

// GetPals returns the Pals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetPals() []Pal {
	if o == nil {
		var ret []Pal
		return ret
	}
	return o.Pals
}

// GetPalsOk returns a tuple with the Pals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetPalsOk() ([]Pal, bool) {
	if o == nil || IsNil(o.Pals) {
		return nil, false
	}
	return o.Pals, true
}

// HasPals returns a boolean if a field has been set.
func (o *RaceHistory) HasPals() bool {
	if o != nil && !IsNil(o.Pals) {
		return true
	}

	return false
}

// SetPals gets a reference to the given []Pal and assigns it to the Pals field.
func (o *RaceHistory) SetPals(v []Pal) {
	o.Pals = v
}

// GetRaceDate returns the RaceDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceDate() int64 {
	if o == nil || IsNil(o.RaceDate.Get()) {
		var ret int64
		return ret
	}
	return *o.RaceDate.Get()
}

// GetRaceDateOk returns a tuple with the RaceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceDate.Get(), o.RaceDate.IsSet()
}

// HasRaceDate returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceDate() bool {
	if o != nil && o.RaceDate.IsSet() {
		return true
	}

	return false
}

// SetRaceDate gets a reference to the given NullableInt64 and assigns it to the RaceDate field.
func (o *RaceHistory) SetRaceDate(v int64) {
	o.RaceDate.Set(&v)
}
// SetRaceDateNil sets the value for RaceDate to be an explicit nil
func (o *RaceHistory) SetRaceDateNil() {
	o.RaceDate.Set(nil)
}

// UnsetRaceDate ensures that no value is present for RaceDate, not even an explicit nil
func (o *RaceHistory) UnsetRaceDate() {
	o.RaceDate.Unset()
}

// GetRaceEntryFee returns the RaceEntryFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceEntryFee() int32 {
	if o == nil || IsNil(o.RaceEntryFee.Get()) {
		var ret int32
		return ret
	}
	return *o.RaceEntryFee.Get()
}

// GetRaceEntryFeeOk returns a tuple with the RaceEntryFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceEntryFeeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceEntryFee.Get(), o.RaceEntryFee.IsSet()
}

// HasRaceEntryFee returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceEntryFee() bool {
	if o != nil && o.RaceEntryFee.IsSet() {
		return true
	}

	return false
}

// SetRaceEntryFee gets a reference to the given NullableInt32 and assigns it to the RaceEntryFee field.
func (o *RaceHistory) SetRaceEntryFee(v int32) {
	o.RaceEntryFee.Set(&v)
}
// SetRaceEntryFeeNil sets the value for RaceEntryFee to be an explicit nil
func (o *RaceHistory) SetRaceEntryFeeNil() {
	o.RaceEntryFee.Set(nil)
}

// UnsetRaceEntryFee ensures that no value is present for RaceEntryFee, not even an explicit nil
func (o *RaceHistory) UnsetRaceEntryFee() {
	o.RaceEntryFee.Unset()
}

// GetRaceId returns the RaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceId() string {
	if o == nil || IsNil(o.RaceId.Get()) {
		var ret string
		return ret
	}
	return *o.RaceId.Get()
}

// GetRaceIdOk returns a tuple with the RaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceId.Get(), o.RaceId.IsSet()
}

// HasRaceId returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceId() bool {
	if o != nil && o.RaceId.IsSet() {
		return true
	}

	return false
}

// SetRaceId gets a reference to the given NullableString and assigns it to the RaceId field.
func (o *RaceHistory) SetRaceId(v string) {
	o.RaceId.Set(&v)
}
// SetRaceIdNil sets the value for RaceId to be an explicit nil
func (o *RaceHistory) SetRaceIdNil() {
	o.RaceId.Set(nil)
}

// UnsetRaceId ensures that no value is present for RaceId, not even an explicit nil
func (o *RaceHistory) UnsetRaceId() {
	o.RaceId.Unset()
}

// GetRaceReward returns the RaceReward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceReward() RaceReward {
	if o == nil || IsNil(o.RaceReward.Get()) {
		var ret RaceReward
		return ret
	}
	return *o.RaceReward.Get()
}

// GetRaceRewardOk returns a tuple with the RaceReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceRewardOk() (*RaceReward, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceReward.Get(), o.RaceReward.IsSet()
}

// HasRaceReward returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceReward() bool {
	if o != nil && o.RaceReward.IsSet() {
		return true
	}

	return false
}

// SetRaceReward gets a reference to the given NullableRaceReward and assigns it to the RaceReward field.
func (o *RaceHistory) SetRaceReward(v RaceReward) {
	o.RaceReward.Set(&v)
}
// SetRaceRewardNil sets the value for RaceReward to be an explicit nil
func (o *RaceHistory) SetRaceRewardNil() {
	o.RaceReward.Set(nil)
}

// UnsetRaceReward ensures that no value is present for RaceReward, not even an explicit nil
func (o *RaceHistory) UnsetRaceReward() {
	o.RaceReward.Unset()
}

// GetRaceStartAt returns the RaceStartAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceStartAt() int64 {
	if o == nil || IsNil(o.RaceStartAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RaceStartAt.Get()
}

// GetRaceStartAtOk returns a tuple with the RaceStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceStartAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceStartAt.Get(), o.RaceStartAt.IsSet()
}

// HasRaceStartAt returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceStartAt() bool {
	if o != nil && o.RaceStartAt.IsSet() {
		return true
	}

	return false
}

// SetRaceStartAt gets a reference to the given NullableInt64 and assigns it to the RaceStartAt field.
func (o *RaceHistory) SetRaceStartAt(v int64) {
	o.RaceStartAt.Set(&v)
}
// SetRaceStartAtNil sets the value for RaceStartAt to be an explicit nil
func (o *RaceHistory) SetRaceStartAtNil() {
	o.RaceStartAt.Set(nil)
}

// UnsetRaceStartAt ensures that no value is present for RaceStartAt, not even an explicit nil
func (o *RaceHistory) UnsetRaceStartAt() {
	o.RaceStartAt.Unset()
}

// GetRaceStatus returns the RaceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetRaceStatus() string {
	if o == nil || IsNil(o.RaceStatus.Get()) {
		var ret string
		return ret
	}
	return *o.RaceStatus.Get()
}

// GetRaceStatusOk returns a tuple with the RaceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRaceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceStatus.Get(), o.RaceStatus.IsSet()
}

// HasRaceStatus returns a boolean if a field has been set.
func (o *RaceHistory) HasRaceStatus() bool {
	if o != nil && o.RaceStatus.IsSet() {
		return true
	}

	return false
}

// SetRaceStatus gets a reference to the given NullableString and assigns it to the RaceStatus field.
func (o *RaceHistory) SetRaceStatus(v string) {
	o.RaceStatus.Set(&v)
}
// SetRaceStatusNil sets the value for RaceStatus to be an explicit nil
func (o *RaceHistory) SetRaceStatusNil() {
	o.RaceStatus.Set(nil)
}

// UnsetRaceStatus ensures that no value is present for RaceStatus, not even an explicit nil
func (o *RaceHistory) UnsetRaceStatus() {
	o.RaceStatus.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetResults() []PalRank {
	if o == nil {
		var ret []PalRank
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetResultsOk() ([]PalRank, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RaceHistory) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PalRank and assigns it to the Results field.
func (o *RaceHistory) SetResults(v []PalRank) {
	o.Results = v
}

// GetReward returns the Reward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceHistory) GetReward() Reward {
	if o == nil || IsNil(o.Reward.Get()) {
		var ret Reward
		return ret
	}
	return *o.Reward.Get()
}

// GetRewardOk returns a tuple with the Reward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceHistory) GetRewardOk() (*Reward, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reward.Get(), o.Reward.IsSet()
}

// HasReward returns a boolean if a field has been set.
func (o *RaceHistory) HasReward() bool {
	if o != nil && o.Reward.IsSet() {
		return true
	}

	return false
}

// SetReward gets a reference to the given NullableReward and assigns it to the Reward field.
func (o *RaceHistory) SetReward(v Reward) {
	o.Reward.Set(&v)
}
// SetRewardNil sets the value for Reward to be an explicit nil
func (o *RaceHistory) SetRewardNil() {
	o.Reward.Set(nil)
}

// UnsetReward ensures that no value is present for Reward, not even an explicit nil
func (o *RaceHistory) UnsetReward() {
	o.Reward.Unset()
}

func (o RaceHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RaceHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AnimationUrl.IsSet() {
		toSerialize["animation_url"] = o.AnimationUrl.Get()
	}
	if o.League.IsSet() {
		toSerialize["league"] = o.League.Get()
	}
	if o.LeagueName.IsSet() {
		toSerialize["league_name"] = o.LeagueName.Get()
	}
	if o.Pals != nil {
		toSerialize["pals"] = o.Pals
	}
	if o.RaceDate.IsSet() {
		toSerialize["race_date"] = o.RaceDate.Get()
	}
	if o.RaceEntryFee.IsSet() {
		toSerialize["race_entry_fee"] = o.RaceEntryFee.Get()
	}
	if o.RaceId.IsSet() {
		toSerialize["race_id"] = o.RaceId.Get()
	}
	if o.RaceReward.IsSet() {
		toSerialize["race_reward"] = o.RaceReward.Get()
	}
	if o.RaceStartAt.IsSet() {
		toSerialize["race_start_at"] = o.RaceStartAt.Get()
	}
	if o.RaceStatus.IsSet() {
		toSerialize["race_status"] = o.RaceStatus.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Reward.IsSet() {
		toSerialize["reward"] = o.Reward.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RaceHistory) UnmarshalJSON(data []byte) (err error) {
	varRaceHistory := _RaceHistory{}

	err = json.Unmarshal(data, &varRaceHistory)

	if err != nil {
		return err
	}

	*o = RaceHistory(varRaceHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "animation_url")
		delete(additionalProperties, "league")
		delete(additionalProperties, "league_name")
		delete(additionalProperties, "pals")
		delete(additionalProperties, "race_date")
		delete(additionalProperties, "race_entry_fee")
		delete(additionalProperties, "race_id")
		delete(additionalProperties, "race_reward")
		delete(additionalProperties, "race_start_at")
		delete(additionalProperties, "race_status")
		delete(additionalProperties, "results")
		delete(additionalProperties, "reward")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRaceHistory struct {
	value *RaceHistory
	isSet bool
}

func (v NullableRaceHistory) Get() *RaceHistory {
	return v.value
}

func (v *NullableRaceHistory) Set(val *RaceHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableRaceHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableRaceHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaceHistory(val *RaceHistory) *NullableRaceHistory {
	return &NullableRaceHistory{value: val, isSet: true}
}

func (v NullableRaceHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaceHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


