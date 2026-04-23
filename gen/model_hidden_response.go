
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the HiddenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HiddenResponse{}

// HiddenResponse struct for HiddenResponse
type HiddenResponse struct {
	HiddenUsers []RealmUser `json:"hidden_users,omitempty"`
	Limit NullableInt32 `json:"limit,omitempty"`
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HiddenResponse HiddenResponse

// NewHiddenResponse instantiates a new HiddenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiddenResponse() *HiddenResponse {
	this := HiddenResponse{}
	return &this
}

// NewHiddenResponseWithDefaults instantiates a new HiddenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiddenResponseWithDefaults() *HiddenResponse {
	this := HiddenResponse{}
	return &this
}

// GetHiddenUsers returns the HiddenUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiddenResponse) GetHiddenUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.HiddenUsers
}

// GetHiddenUsersOk returns a tuple with the HiddenUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiddenResponse) GetHiddenUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.HiddenUsers) {
		return nil, false
	}
	return o.HiddenUsers, true
}

// HasHiddenUsers returns a boolean if a field has been set.
func (o *HiddenResponse) HasHiddenUsers() bool {
	if o != nil && !IsNil(o.HiddenUsers) {
		return true
	}

	return false
}

// SetHiddenUsers gets a reference to the given []RealmUser and assigns it to the HiddenUsers field.
func (o *HiddenResponse) SetHiddenUsers(v []RealmUser) {
	o.HiddenUsers = v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiddenResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiddenResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *HiddenResponse) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *HiddenResponse) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *HiddenResponse) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *HiddenResponse) UnsetLimit() {
	o.Limit.Unset()
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiddenResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiddenResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *HiddenResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *HiddenResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *HiddenResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *HiddenResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiddenResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiddenResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *HiddenResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *HiddenResponse) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *HiddenResponse) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *HiddenResponse) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o HiddenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HiddenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HiddenUsers != nil {
		toSerialize["hidden_users"] = o.HiddenUsers
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HiddenResponse) UnmarshalJSON(data []byte) (err error) {
	varHiddenResponse := _HiddenResponse{}

	err = json.Unmarshal(data, &varHiddenResponse)

	if err != nil {
		return err
	}

	*o = HiddenResponse(varHiddenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hidden_users")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHiddenResponse struct {
	value *HiddenResponse
	isSet bool
}

func (v NullableHiddenResponse) Get() *HiddenResponse {
	return v.value
}

func (v *NullableHiddenResponse) Set(val *HiddenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHiddenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHiddenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiddenResponse(val *HiddenResponse) *NullableHiddenResponse {
	return &NullableHiddenResponse{value: val, isSet: true}
}

func (v NullableHiddenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiddenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


