
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRank{}

// UserRank struct for UserRank
type UserRank struct {
	Ranking NullableInt32 `json:"ranking,omitempty"`
	TopGifts []string `json:"top_gifts,omitempty"`
	TotalScores NullableInt32 `json:"total_scores,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRank UserRank

// NewUserRank instantiates a new UserRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRank() *UserRank {
	this := UserRank{}
	return &this
}

// NewUserRankWithDefaults instantiates a new UserRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRankWithDefaults() *UserRank {
	this := UserRank{}
	return &this
}

// GetRanking returns the Ranking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRank) GetRanking() int32 {
	if o == nil || IsNil(o.Ranking.Get()) {
		var ret int32
		return ret
	}
	return *o.Ranking.Get()
}

// GetRankingOk returns a tuple with the Ranking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRank) GetRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ranking.Get(), o.Ranking.IsSet()
}

// HasRanking returns a boolean if a field has been set.
func (o *UserRank) HasRanking() bool {
	if o != nil && o.Ranking.IsSet() {
		return true
	}

	return false
}

// SetRanking gets a reference to the given NullableInt32 and assigns it to the Ranking field.
func (o *UserRank) SetRanking(v int32) {
	o.Ranking.Set(&v)
}
// SetRankingNil sets the value for Ranking to be an explicit nil
func (o *UserRank) SetRankingNil() {
	o.Ranking.Set(nil)
}

// UnsetRanking ensures that no value is present for Ranking, not even an explicit nil
func (o *UserRank) UnsetRanking() {
	o.Ranking.Unset()
}

// GetTopGifts returns the TopGifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRank) GetTopGifts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TopGifts
}

// GetTopGiftsOk returns a tuple with the TopGifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRank) GetTopGiftsOk() ([]string, bool) {
	if o == nil || IsNil(o.TopGifts) {
		return nil, false
	}
	return o.TopGifts, true
}

// HasTopGifts returns a boolean if a field has been set.
func (o *UserRank) HasTopGifts() bool {
	if o != nil && !IsNil(o.TopGifts) {
		return true
	}

	return false
}

// SetTopGifts gets a reference to the given []string and assigns it to the TopGifts field.
func (o *UserRank) SetTopGifts(v []string) {
	o.TopGifts = v
}

// GetTotalScores returns the TotalScores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRank) GetTotalScores() int32 {
	if o == nil || IsNil(o.TotalScores.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalScores.Get()
}

// GetTotalScoresOk returns a tuple with the TotalScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRank) GetTotalScoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalScores.Get(), o.TotalScores.IsSet()
}

// HasTotalScores returns a boolean if a field has been set.
func (o *UserRank) HasTotalScores() bool {
	if o != nil && o.TotalScores.IsSet() {
		return true
	}

	return false
}

// SetTotalScores gets a reference to the given NullableInt32 and assigns it to the TotalScores field.
func (o *UserRank) SetTotalScores(v int32) {
	o.TotalScores.Set(&v)
}
// SetTotalScoresNil sets the value for TotalScores to be an explicit nil
func (o *UserRank) SetTotalScoresNil() {
	o.TotalScores.Set(nil)
}

// UnsetTotalScores ensures that no value is present for TotalScores, not even an explicit nil
func (o *UserRank) UnsetTotalScores() {
	o.TotalScores.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRank) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRank) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserRank) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *UserRank) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *UserRank) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *UserRank) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRank) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRank) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *UserRank) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *UserRank) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *UserRank) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *UserRank) UnsetUser() {
	o.User.Unset()
}

func (o UserRank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ranking.IsSet() {
		toSerialize["ranking"] = o.Ranking.Get()
	}
	if o.TopGifts != nil {
		toSerialize["top_gifts"] = o.TopGifts
	}
	if o.TotalScores.IsSet() {
		toSerialize["total_scores"] = o.TotalScores.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRank) UnmarshalJSON(data []byte) (err error) {
	varUserRank := _UserRank{}

	err = json.Unmarshal(data, &varUserRank)

	if err != nil {
		return err
	}

	*o = UserRank(varUserRank)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ranking")
		delete(additionalProperties, "top_gifts")
		delete(additionalProperties, "total_scores")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRank struct {
	value *UserRank
	isSet bool
}

func (v NullableUserRank) Get() *UserRank {
	return v.value
}

func (v *NullableUserRank) Set(val *UserRank) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRank) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRank(val *UserRank) *NullableUserRank {
	return &NullableUserRank{value: val, isSet: true}
}

func (v NullableUserRank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


