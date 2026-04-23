
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserCustomDefinitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCustomDefinitionsResponse{}

// UserCustomDefinitionsResponse struct for UserCustomDefinitionsResponse
type UserCustomDefinitionsResponse struct {
	Age NullableInt32 `json:"age,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	FollowingsCount NullableInt32 `json:"followings_count,omitempty"`
	LastLoggedinAt NullableInt64 `json:"last_loggedin_at,omitempty"`
	ReportedCount NullableInt32 `json:"reported_count,omitempty"`
	Status NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCustomDefinitionsResponse UserCustomDefinitionsResponse

// NewUserCustomDefinitionsResponse instantiates a new UserCustomDefinitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCustomDefinitionsResponse() *UserCustomDefinitionsResponse {
	this := UserCustomDefinitionsResponse{}
	return &this
}

// NewUserCustomDefinitionsResponseWithDefaults instantiates a new UserCustomDefinitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCustomDefinitionsResponseWithDefaults() *UserCustomDefinitionsResponse {
	this := UserCustomDefinitionsResponse{}
	return &this
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetAge() int32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret int32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableInt32 and assigns it to the Age field.
func (o *UserCustomDefinitionsResponse) SetAge(v int32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetAge() {
	o.Age.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *UserCustomDefinitionsResponse) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *UserCustomDefinitionsResponse) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFollowingsCount returns the FollowingsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetFollowingsCount() int32 {
	if o == nil || IsNil(o.FollowingsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowingsCount.Get()
}

// GetFollowingsCountOk returns a tuple with the FollowingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetFollowingsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingsCount.Get(), o.FollowingsCount.IsSet()
}

// HasFollowingsCount returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasFollowingsCount() bool {
	if o != nil && o.FollowingsCount.IsSet() {
		return true
	}

	return false
}

// SetFollowingsCount gets a reference to the given NullableInt32 and assigns it to the FollowingsCount field.
func (o *UserCustomDefinitionsResponse) SetFollowingsCount(v int32) {
	o.FollowingsCount.Set(&v)
}
// SetFollowingsCountNil sets the value for FollowingsCount to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetFollowingsCountNil() {
	o.FollowingsCount.Set(nil)
}

// UnsetFollowingsCount ensures that no value is present for FollowingsCount, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetFollowingsCount() {
	o.FollowingsCount.Unset()
}

// GetLastLoggedinAt returns the LastLoggedinAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetLastLoggedinAt() int64 {
	if o == nil || IsNil(o.LastLoggedinAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LastLoggedinAt.Get()
}

// GetLastLoggedinAtOk returns a tuple with the LastLoggedinAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetLastLoggedinAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoggedinAt.Get(), o.LastLoggedinAt.IsSet()
}

// HasLastLoggedinAt returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasLastLoggedinAt() bool {
	if o != nil && o.LastLoggedinAt.IsSet() {
		return true
	}

	return false
}

// SetLastLoggedinAt gets a reference to the given NullableInt64 and assigns it to the LastLoggedinAt field.
func (o *UserCustomDefinitionsResponse) SetLastLoggedinAt(v int64) {
	o.LastLoggedinAt.Set(&v)
}
// SetLastLoggedinAtNil sets the value for LastLoggedinAt to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetLastLoggedinAtNil() {
	o.LastLoggedinAt.Set(nil)
}

// UnsetLastLoggedinAt ensures that no value is present for LastLoggedinAt, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetLastLoggedinAt() {
	o.LastLoggedinAt.Unset()
}

// GetReportedCount returns the ReportedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetReportedCount() int32 {
	if o == nil || IsNil(o.ReportedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReportedCount.Get()
}

// GetReportedCountOk returns a tuple with the ReportedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetReportedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedCount.Get(), o.ReportedCount.IsSet()
}

// HasReportedCount returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasReportedCount() bool {
	if o != nil && o.ReportedCount.IsSet() {
		return true
	}

	return false
}

// SetReportedCount gets a reference to the given NullableInt32 and assigns it to the ReportedCount field.
func (o *UserCustomDefinitionsResponse) SetReportedCount(v int32) {
	o.ReportedCount.Set(&v)
}
// SetReportedCountNil sets the value for ReportedCount to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetReportedCountNil() {
	o.ReportedCount.Set(nil)
}

// UnsetReportedCount ensures that no value is present for ReportedCount, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetReportedCount() {
	o.ReportedCount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomDefinitionsResponse) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomDefinitionsResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *UserCustomDefinitionsResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *UserCustomDefinitionsResponse) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *UserCustomDefinitionsResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *UserCustomDefinitionsResponse) UnsetStatus() {
	o.Status.Unset()
}

func (o UserCustomDefinitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCustomDefinitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.FollowersCount.IsSet() {
		toSerialize["followers_count"] = o.FollowersCount.Get()
	}
	if o.FollowingsCount.IsSet() {
		toSerialize["followings_count"] = o.FollowingsCount.Get()
	}
	if o.LastLoggedinAt.IsSet() {
		toSerialize["last_loggedin_at"] = o.LastLoggedinAt.Get()
	}
	if o.ReportedCount.IsSet() {
		toSerialize["reported_count"] = o.ReportedCount.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCustomDefinitionsResponse) UnmarshalJSON(data []byte) (err error) {
	varUserCustomDefinitionsResponse := _UserCustomDefinitionsResponse{}

	err = json.Unmarshal(data, &varUserCustomDefinitionsResponse)

	if err != nil {
		return err
	}

	*o = UserCustomDefinitionsResponse(varUserCustomDefinitionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "age")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "followings_count")
		delete(additionalProperties, "last_loggedin_at")
		delete(additionalProperties, "reported_count")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCustomDefinitionsResponse struct {
	value *UserCustomDefinitionsResponse
	isSet bool
}

func (v NullableUserCustomDefinitionsResponse) Get() *UserCustomDefinitionsResponse {
	return v.value
}

func (v *NullableUserCustomDefinitionsResponse) Set(val *UserCustomDefinitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCustomDefinitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCustomDefinitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCustomDefinitionsResponse(val *UserCustomDefinitionsResponse) *NullableUserCustomDefinitionsResponse {
	return &NullableUserCustomDefinitionsResponse{value: val, isSet: true}
}

func (v NullableUserCustomDefinitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCustomDefinitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


