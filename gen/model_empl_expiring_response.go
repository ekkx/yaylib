
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplExpiringResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplExpiringResponse{}

// EmplExpiringResponse struct for EmplExpiringResponse
type EmplExpiringResponse struct {
	ExpiringAmount NullableString `json:"expiring_amount,omitempty"`
	ExpiringDate NullableInt64 `json:"expiring_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplExpiringResponse EmplExpiringResponse

// NewEmplExpiringResponse instantiates a new EmplExpiringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplExpiringResponse() *EmplExpiringResponse {
	this := EmplExpiringResponse{}
	return &this
}

// NewEmplExpiringResponseWithDefaults instantiates a new EmplExpiringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplExpiringResponseWithDefaults() *EmplExpiringResponse {
	this := EmplExpiringResponse{}
	return &this
}

// GetExpiringAmount returns the ExpiringAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplExpiringResponse) GetExpiringAmount() string {
	if o == nil || IsNil(o.ExpiringAmount.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiringAmount.Get()
}

// GetExpiringAmountOk returns a tuple with the ExpiringAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplExpiringResponse) GetExpiringAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringAmount.Get(), o.ExpiringAmount.IsSet()
}

// HasExpiringAmount returns a boolean if a field has been set.
func (o *EmplExpiringResponse) HasExpiringAmount() bool {
	if o != nil && o.ExpiringAmount.IsSet() {
		return true
	}

	return false
}

// SetExpiringAmount gets a reference to the given NullableString and assigns it to the ExpiringAmount field.
func (o *EmplExpiringResponse) SetExpiringAmount(v string) {
	o.ExpiringAmount.Set(&v)
}
// SetExpiringAmountNil sets the value for ExpiringAmount to be an explicit nil
func (o *EmplExpiringResponse) SetExpiringAmountNil() {
	o.ExpiringAmount.Set(nil)
}

// UnsetExpiringAmount ensures that no value is present for ExpiringAmount, not even an explicit nil
func (o *EmplExpiringResponse) UnsetExpiringAmount() {
	o.ExpiringAmount.Unset()
}

// GetExpiringDate returns the ExpiringDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplExpiringResponse) GetExpiringDate() int64 {
	if o == nil || IsNil(o.ExpiringDate.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiringDate.Get()
}

// GetExpiringDateOk returns a tuple with the ExpiringDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplExpiringResponse) GetExpiringDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringDate.Get(), o.ExpiringDate.IsSet()
}

// HasExpiringDate returns a boolean if a field has been set.
func (o *EmplExpiringResponse) HasExpiringDate() bool {
	if o != nil && o.ExpiringDate.IsSet() {
		return true
	}

	return false
}

// SetExpiringDate gets a reference to the given NullableInt64 and assigns it to the ExpiringDate field.
func (o *EmplExpiringResponse) SetExpiringDate(v int64) {
	o.ExpiringDate.Set(&v)
}
// SetExpiringDateNil sets the value for ExpiringDate to be an explicit nil
func (o *EmplExpiringResponse) SetExpiringDateNil() {
	o.ExpiringDate.Set(nil)
}

// UnsetExpiringDate ensures that no value is present for ExpiringDate, not even an explicit nil
func (o *EmplExpiringResponse) UnsetExpiringDate() {
	o.ExpiringDate.Unset()
}

func (o EmplExpiringResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplExpiringResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiringAmount.IsSet() {
		toSerialize["expiring_amount"] = o.ExpiringAmount.Get()
	}
	if o.ExpiringDate.IsSet() {
		toSerialize["expiring_date"] = o.ExpiringDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplExpiringResponse) UnmarshalJSON(data []byte) (err error) {
	varEmplExpiringResponse := _EmplExpiringResponse{}

	err = json.Unmarshal(data, &varEmplExpiringResponse)

	if err != nil {
		return err
	}

	*o = EmplExpiringResponse(varEmplExpiringResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiring_amount")
		delete(additionalProperties, "expiring_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplExpiringResponse struct {
	value *EmplExpiringResponse
	isSet bool
}

func (v NullableEmplExpiringResponse) Get() *EmplExpiringResponse {
	return v.value
}

func (v *NullableEmplExpiringResponse) Set(val *EmplExpiringResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplExpiringResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplExpiringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplExpiringResponse(val *EmplExpiringResponse) *NullableEmplExpiringResponse {
	return &NullableEmplExpiringResponse{value: val, isSet: true}
}

func (v NullableEmplExpiringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplExpiringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


