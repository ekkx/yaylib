
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplExchangeQuotaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplExchangeQuotaResponse{}

// Web3EmplExchangeQuotaResponse struct for Web3EmplExchangeQuotaResponse
type Web3EmplExchangeQuotaResponse struct {
	EmplExchangeLimit NullableInt64 `json:"empl_exchange_limit,omitempty"`
	EmplExchangeQuotaRemaining NullableInt64 `json:"empl_exchange_quota_remaining,omitempty"`
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplExchangeQuotaResponse Web3EmplExchangeQuotaResponse

// NewWeb3EmplExchangeQuotaResponse instantiates a new Web3EmplExchangeQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplExchangeQuotaResponse() *Web3EmplExchangeQuotaResponse {
	this := Web3EmplExchangeQuotaResponse{}
	return &this
}

// NewWeb3EmplExchangeQuotaResponseWithDefaults instantiates a new Web3EmplExchangeQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplExchangeQuotaResponseWithDefaults() *Web3EmplExchangeQuotaResponse {
	this := Web3EmplExchangeQuotaResponse{}
	return &this
}

// GetEmplExchangeLimit returns the EmplExchangeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplExchangeQuotaResponse) GetEmplExchangeLimit() int64 {
	if o == nil || IsNil(o.EmplExchangeLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.EmplExchangeLimit.Get()
}

// GetEmplExchangeLimitOk returns a tuple with the EmplExchangeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplExchangeQuotaResponse) GetEmplExchangeLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplExchangeLimit.Get(), o.EmplExchangeLimit.IsSet()
}

// HasEmplExchangeLimit returns a boolean if a field has been set.
func (o *Web3EmplExchangeQuotaResponse) HasEmplExchangeLimit() bool {
	if o != nil && o.EmplExchangeLimit.IsSet() {
		return true
	}

	return false
}

// SetEmplExchangeLimit gets a reference to the given NullableInt64 and assigns it to the EmplExchangeLimit field.
func (o *Web3EmplExchangeQuotaResponse) SetEmplExchangeLimit(v int64) {
	o.EmplExchangeLimit.Set(&v)
}
// SetEmplExchangeLimitNil sets the value for EmplExchangeLimit to be an explicit nil
func (o *Web3EmplExchangeQuotaResponse) SetEmplExchangeLimitNil() {
	o.EmplExchangeLimit.Set(nil)
}

// UnsetEmplExchangeLimit ensures that no value is present for EmplExchangeLimit, not even an explicit nil
func (o *Web3EmplExchangeQuotaResponse) UnsetEmplExchangeLimit() {
	o.EmplExchangeLimit.Unset()
}

// GetEmplExchangeQuotaRemaining returns the EmplExchangeQuotaRemaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplExchangeQuotaResponse) GetEmplExchangeQuotaRemaining() int64 {
	if o == nil || IsNil(o.EmplExchangeQuotaRemaining.Get()) {
		var ret int64
		return ret
	}
	return *o.EmplExchangeQuotaRemaining.Get()
}

// GetEmplExchangeQuotaRemainingOk returns a tuple with the EmplExchangeQuotaRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplExchangeQuotaResponse) GetEmplExchangeQuotaRemainingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplExchangeQuotaRemaining.Get(), o.EmplExchangeQuotaRemaining.IsSet()
}

// HasEmplExchangeQuotaRemaining returns a boolean if a field has been set.
func (o *Web3EmplExchangeQuotaResponse) HasEmplExchangeQuotaRemaining() bool {
	if o != nil && o.EmplExchangeQuotaRemaining.IsSet() {
		return true
	}

	return false
}

// SetEmplExchangeQuotaRemaining gets a reference to the given NullableInt64 and assigns it to the EmplExchangeQuotaRemaining field.
func (o *Web3EmplExchangeQuotaResponse) SetEmplExchangeQuotaRemaining(v int64) {
	o.EmplExchangeQuotaRemaining.Set(&v)
}
// SetEmplExchangeQuotaRemainingNil sets the value for EmplExchangeQuotaRemaining to be an explicit nil
func (o *Web3EmplExchangeQuotaResponse) SetEmplExchangeQuotaRemainingNil() {
	o.EmplExchangeQuotaRemaining.Set(nil)
}

// UnsetEmplExchangeQuotaRemaining ensures that no value is present for EmplExchangeQuotaRemaining, not even an explicit nil
func (o *Web3EmplExchangeQuotaResponse) UnsetEmplExchangeQuotaRemaining() {
	o.EmplExchangeQuotaRemaining.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplExchangeQuotaResponse) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplExchangeQuotaResponse) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Web3EmplExchangeQuotaResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *Web3EmplExchangeQuotaResponse) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *Web3EmplExchangeQuotaResponse) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *Web3EmplExchangeQuotaResponse) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o Web3EmplExchangeQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplExchangeQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmplExchangeLimit.IsSet() {
		toSerialize["empl_exchange_limit"] = o.EmplExchangeLimit.Get()
	}
	if o.EmplExchangeQuotaRemaining.IsSet() {
		toSerialize["empl_exchange_quota_remaining"] = o.EmplExchangeQuotaRemaining.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplExchangeQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplExchangeQuotaResponse := _Web3EmplExchangeQuotaResponse{}

	err = json.Unmarshal(data, &varWeb3EmplExchangeQuotaResponse)

	if err != nil {
		return err
	}

	*o = Web3EmplExchangeQuotaResponse(varWeb3EmplExchangeQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl_exchange_limit")
		delete(additionalProperties, "empl_exchange_quota_remaining")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplExchangeQuotaResponse struct {
	value *Web3EmplExchangeQuotaResponse
	isSet bool
}

func (v NullableWeb3EmplExchangeQuotaResponse) Get() *Web3EmplExchangeQuotaResponse {
	return v.value
}

func (v *NullableWeb3EmplExchangeQuotaResponse) Set(val *Web3EmplExchangeQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplExchangeQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplExchangeQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplExchangeQuotaResponse(val *Web3EmplExchangeQuotaResponse) *NullableWeb3EmplExchangeQuotaResponse {
	return &NullableWeb3EmplExchangeQuotaResponse{value: val, isSet: true}
}

func (v NullableWeb3EmplExchangeQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplExchangeQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


