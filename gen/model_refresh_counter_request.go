
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RefreshCounterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshCounterRequest{}

// RefreshCounterRequest struct for RefreshCounterRequest
type RefreshCounterRequest struct {
	Counter NullableString `json:"counter,omitempty"`
	LastRequestedAt NullableInt64 `json:"last_requested_at,omitempty"`
	Status NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RefreshCounterRequest RefreshCounterRequest

// NewRefreshCounterRequest instantiates a new RefreshCounterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshCounterRequest() *RefreshCounterRequest {
	this := RefreshCounterRequest{}
	return &this
}

// NewRefreshCounterRequestWithDefaults instantiates a new RefreshCounterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshCounterRequestWithDefaults() *RefreshCounterRequest {
	this := RefreshCounterRequest{}
	return &this
}

// GetCounter returns the Counter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshCounterRequest) GetCounter() string {
	if o == nil || IsNil(o.Counter.Get()) {
		var ret string
		return ret
	}
	return *o.Counter.Get()
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshCounterRequest) GetCounterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Counter.Get(), o.Counter.IsSet()
}

// HasCounter returns a boolean if a field has been set.
func (o *RefreshCounterRequest) HasCounter() bool {
	if o != nil && o.Counter.IsSet() {
		return true
	}

	return false
}

// SetCounter gets a reference to the given NullableString and assigns it to the Counter field.
func (o *RefreshCounterRequest) SetCounter(v string) {
	o.Counter.Set(&v)
}
// SetCounterNil sets the value for Counter to be an explicit nil
func (o *RefreshCounterRequest) SetCounterNil() {
	o.Counter.Set(nil)
}

// UnsetCounter ensures that no value is present for Counter, not even an explicit nil
func (o *RefreshCounterRequest) UnsetCounter() {
	o.Counter.Unset()
}

// GetLastRequestedAt returns the LastRequestedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshCounterRequest) GetLastRequestedAt() int64 {
	if o == nil || IsNil(o.LastRequestedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LastRequestedAt.Get()
}

// GetLastRequestedAtOk returns a tuple with the LastRequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshCounterRequest) GetLastRequestedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRequestedAt.Get(), o.LastRequestedAt.IsSet()
}

// HasLastRequestedAt returns a boolean if a field has been set.
func (o *RefreshCounterRequest) HasLastRequestedAt() bool {
	if o != nil && o.LastRequestedAt.IsSet() {
		return true
	}

	return false
}

// SetLastRequestedAt gets a reference to the given NullableInt64 and assigns it to the LastRequestedAt field.
func (o *RefreshCounterRequest) SetLastRequestedAt(v int64) {
	o.LastRequestedAt.Set(&v)
}
// SetLastRequestedAtNil sets the value for LastRequestedAt to be an explicit nil
func (o *RefreshCounterRequest) SetLastRequestedAtNil() {
	o.LastRequestedAt.Set(nil)
}

// UnsetLastRequestedAt ensures that no value is present for LastRequestedAt, not even an explicit nil
func (o *RefreshCounterRequest) UnsetLastRequestedAt() {
	o.LastRequestedAt.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshCounterRequest) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshCounterRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RefreshCounterRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *RefreshCounterRequest) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RefreshCounterRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RefreshCounterRequest) UnsetStatus() {
	o.Status.Unset()
}

func (o RefreshCounterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshCounterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Counter.IsSet() {
		toSerialize["counter"] = o.Counter.Get()
	}
	if o.LastRequestedAt.IsSet() {
		toSerialize["last_requested_at"] = o.LastRequestedAt.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshCounterRequest) UnmarshalJSON(data []byte) (err error) {
	varRefreshCounterRequest := _RefreshCounterRequest{}

	err = json.Unmarshal(data, &varRefreshCounterRequest)

	if err != nil {
		return err
	}

	*o = RefreshCounterRequest(varRefreshCounterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "counter")
		delete(additionalProperties, "last_requested_at")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshCounterRequest struct {
	value *RefreshCounterRequest
	isSet bool
}

func (v NullableRefreshCounterRequest) Get() *RefreshCounterRequest {
	return v.value
}

func (v *NullableRefreshCounterRequest) Set(val *RefreshCounterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshCounterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshCounterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshCounterRequest(val *RefreshCounterRequest) *NullableRefreshCounterRequest {
	return &NullableRefreshCounterRequest{value: val, isSet: true}
}

func (v NullableRefreshCounterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshCounterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


