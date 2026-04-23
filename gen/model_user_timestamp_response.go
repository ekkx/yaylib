
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserTimestampResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTimestampResponse{}

// UserTimestampResponse struct for UserTimestampResponse
type UserTimestampResponse struct {
	Country NullableString `json:"country,omitempty"`
	IpAddress NullableString `json:"ip_address,omitempty"`
	Time NullableInt64 `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserTimestampResponse UserTimestampResponse

// NewUserTimestampResponse instantiates a new UserTimestampResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTimestampResponse() *UserTimestampResponse {
	this := UserTimestampResponse{}
	return &this
}

// NewUserTimestampResponseWithDefaults instantiates a new UserTimestampResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTimestampResponseWithDefaults() *UserTimestampResponse {
	this := UserTimestampResponse{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTimestampResponse) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTimestampResponse) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *UserTimestampResponse) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *UserTimestampResponse) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *UserTimestampResponse) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *UserTimestampResponse) UnsetCountry() {
	o.Country.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTimestampResponse) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTimestampResponse) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *UserTimestampResponse) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *UserTimestampResponse) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *UserTimestampResponse) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *UserTimestampResponse) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTimestampResponse) GetTime() int64 {
	if o == nil || IsNil(o.Time.Get()) {
		var ret int64
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTimestampResponse) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *UserTimestampResponse) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableInt64 and assigns it to the Time field.
func (o *UserTimestampResponse) SetTime(v int64) {
	o.Time.Set(&v)
}
// SetTimeNil sets the value for Time to be an explicit nil
func (o *UserTimestampResponse) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *UserTimestampResponse) UnsetTime() {
	o.Time.Unset()
}

func (o UserTimestampResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTimestampResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserTimestampResponse) UnmarshalJSON(data []byte) (err error) {
	varUserTimestampResponse := _UserTimestampResponse{}

	err = json.Unmarshal(data, &varUserTimestampResponse)

	if err != nil {
		return err
	}

	*o = UserTimestampResponse(varUserTimestampResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "ip_address")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserTimestampResponse struct {
	value *UserTimestampResponse
	isSet bool
}

func (v NullableUserTimestampResponse) Get() *UserTimestampResponse {
	return v.value
}

func (v *NullableUserTimestampResponse) Set(val *UserTimestampResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTimestampResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTimestampResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTimestampResponse(val *UserTimestampResponse) *NullableUserTimestampResponse {
	return &NullableUserTimestampResponse{value: val, isSet: true}
}

func (v NullableUserTimestampResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTimestampResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


