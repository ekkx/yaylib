
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchQuery{}

// UserSearchQuery struct for UserSearchQuery
type UserSearchQuery struct {
	Keyword NullableString `json:"keyword,omitempty"`
	NoRecentPenalty NullableBool `json:"no_recent_penalty,omitempty"`
	SamePrefecture NullableBool `json:"same_prefecture,omitempty"`
	Scope NullableUserSearchScope `json:"scope,omitempty"`
	SimilarAge NullableBool `json:"similar_age,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSearchQuery UserSearchQuery

// NewUserSearchQuery instantiates a new UserSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchQuery() *UserSearchQuery {
	this := UserSearchQuery{}
	return &this
}

// NewUserSearchQueryWithDefaults instantiates a new UserSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchQueryWithDefaults() *UserSearchQuery {
	this := UserSearchQuery{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchQuery) GetKeyword() string {
	if o == nil || IsNil(o.Keyword.Get()) {
		var ret string
		return ret
	}
	return *o.Keyword.Get()
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchQuery) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keyword.Get(), o.Keyword.IsSet()
}

// HasKeyword returns a boolean if a field has been set.
func (o *UserSearchQuery) HasKeyword() bool {
	if o != nil && o.Keyword.IsSet() {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given NullableString and assigns it to the Keyword field.
func (o *UserSearchQuery) SetKeyword(v string) {
	o.Keyword.Set(&v)
}
// SetKeywordNil sets the value for Keyword to be an explicit nil
func (o *UserSearchQuery) SetKeywordNil() {
	o.Keyword.Set(nil)
}

// UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
func (o *UserSearchQuery) UnsetKeyword() {
	o.Keyword.Unset()
}

// GetNoRecentPenalty returns the NoRecentPenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchQuery) GetNoRecentPenalty() bool {
	if o == nil || IsNil(o.NoRecentPenalty.Get()) {
		var ret bool
		return ret
	}
	return *o.NoRecentPenalty.Get()
}

// GetNoRecentPenaltyOk returns a tuple with the NoRecentPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchQuery) GetNoRecentPenaltyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoRecentPenalty.Get(), o.NoRecentPenalty.IsSet()
}

// HasNoRecentPenalty returns a boolean if a field has been set.
func (o *UserSearchQuery) HasNoRecentPenalty() bool {
	if o != nil && o.NoRecentPenalty.IsSet() {
		return true
	}

	return false
}

// SetNoRecentPenalty gets a reference to the given NullableBool and assigns it to the NoRecentPenalty field.
func (o *UserSearchQuery) SetNoRecentPenalty(v bool) {
	o.NoRecentPenalty.Set(&v)
}
// SetNoRecentPenaltyNil sets the value for NoRecentPenalty to be an explicit nil
func (o *UserSearchQuery) SetNoRecentPenaltyNil() {
	o.NoRecentPenalty.Set(nil)
}

// UnsetNoRecentPenalty ensures that no value is present for NoRecentPenalty, not even an explicit nil
func (o *UserSearchQuery) UnsetNoRecentPenalty() {
	o.NoRecentPenalty.Unset()
}

// GetSamePrefecture returns the SamePrefecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchQuery) GetSamePrefecture() bool {
	if o == nil || IsNil(o.SamePrefecture.Get()) {
		var ret bool
		return ret
	}
	return *o.SamePrefecture.Get()
}

// GetSamePrefectureOk returns a tuple with the SamePrefecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchQuery) GetSamePrefectureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamePrefecture.Get(), o.SamePrefecture.IsSet()
}

// HasSamePrefecture returns a boolean if a field has been set.
func (o *UserSearchQuery) HasSamePrefecture() bool {
	if o != nil && o.SamePrefecture.IsSet() {
		return true
	}

	return false
}

// SetSamePrefecture gets a reference to the given NullableBool and assigns it to the SamePrefecture field.
func (o *UserSearchQuery) SetSamePrefecture(v bool) {
	o.SamePrefecture.Set(&v)
}
// SetSamePrefectureNil sets the value for SamePrefecture to be an explicit nil
func (o *UserSearchQuery) SetSamePrefectureNil() {
	o.SamePrefecture.Set(nil)
}

// UnsetSamePrefecture ensures that no value is present for SamePrefecture, not even an explicit nil
func (o *UserSearchQuery) UnsetSamePrefecture() {
	o.SamePrefecture.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchQuery) GetScope() UserSearchScope {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret UserSearchScope
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchQuery) GetScopeOk() (*UserSearchScope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *UserSearchQuery) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableUserSearchScope and assigns it to the Scope field.
func (o *UserSearchQuery) SetScope(v UserSearchScope) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *UserSearchQuery) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *UserSearchQuery) UnsetScope() {
	o.Scope.Unset()
}

// GetSimilarAge returns the SimilarAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchQuery) GetSimilarAge() bool {
	if o == nil || IsNil(o.SimilarAge.Get()) {
		var ret bool
		return ret
	}
	return *o.SimilarAge.Get()
}

// GetSimilarAgeOk returns a tuple with the SimilarAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchQuery) GetSimilarAgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SimilarAge.Get(), o.SimilarAge.IsSet()
}

// HasSimilarAge returns a boolean if a field has been set.
func (o *UserSearchQuery) HasSimilarAge() bool {
	if o != nil && o.SimilarAge.IsSet() {
		return true
	}

	return false
}

// SetSimilarAge gets a reference to the given NullableBool and assigns it to the SimilarAge field.
func (o *UserSearchQuery) SetSimilarAge(v bool) {
	o.SimilarAge.Set(&v)
}
// SetSimilarAgeNil sets the value for SimilarAge to be an explicit nil
func (o *UserSearchQuery) SetSimilarAgeNil() {
	o.SimilarAge.Set(nil)
}

// UnsetSimilarAge ensures that no value is present for SimilarAge, not even an explicit nil
func (o *UserSearchQuery) UnsetSimilarAge() {
	o.SimilarAge.Unset()
}

func (o UserSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Keyword.IsSet() {
		toSerialize["keyword"] = o.Keyword.Get()
	}
	if o.NoRecentPenalty.IsSet() {
		toSerialize["no_recent_penalty"] = o.NoRecentPenalty.Get()
	}
	if o.SamePrefecture.IsSet() {
		toSerialize["same_prefecture"] = o.SamePrefecture.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.SimilarAge.IsSet() {
		toSerialize["similar_age"] = o.SimilarAge.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSearchQuery) UnmarshalJSON(data []byte) (err error) {
	varUserSearchQuery := _UserSearchQuery{}

	err = json.Unmarshal(data, &varUserSearchQuery)

	if err != nil {
		return err
	}

	*o = UserSearchQuery(varUserSearchQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keyword")
		delete(additionalProperties, "no_recent_penalty")
		delete(additionalProperties, "same_prefecture")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "similar_age")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSearchQuery struct {
	value *UserSearchQuery
	isSet bool
}

func (v NullableUserSearchQuery) Get() *UserSearchQuery {
	return v.value
}

func (v *NullableUserSearchQuery) Set(val *UserSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchQuery(val *UserSearchQuery) *NullableUserSearchQuery {
	return &NullableUserSearchQuery{value: val, isSet: true}
}

func (v NullableUserSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


