
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ExpiredEmplResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpiredEmplResponse{}

// ExpiredEmplResponse struct for ExpiredEmplResponse
type ExpiredEmplResponse struct {
	Expirations []ExpiredEmpl `json:"expirations,omitempty"`
	TotalAmount NullableString `json:"total_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpiredEmplResponse ExpiredEmplResponse

// NewExpiredEmplResponse instantiates a new ExpiredEmplResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiredEmplResponse() *ExpiredEmplResponse {
	this := ExpiredEmplResponse{}
	return &this
}

// NewExpiredEmplResponseWithDefaults instantiates a new ExpiredEmplResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiredEmplResponseWithDefaults() *ExpiredEmplResponse {
	this := ExpiredEmplResponse{}
	return &this
}

// GetExpirations returns the Expirations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiredEmplResponse) GetExpirations() []ExpiredEmpl {
	if o == nil {
		var ret []ExpiredEmpl
		return ret
	}
	return o.Expirations
}

// GetExpirationsOk returns a tuple with the Expirations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiredEmplResponse) GetExpirationsOk() ([]ExpiredEmpl, bool) {
	if o == nil || IsNil(o.Expirations) {
		return nil, false
	}
	return o.Expirations, true
}

// HasExpirations returns a boolean if a field has been set.
func (o *ExpiredEmplResponse) HasExpirations() bool {
	if o != nil && !IsNil(o.Expirations) {
		return true
	}

	return false
}

// SetExpirations gets a reference to the given []ExpiredEmpl and assigns it to the Expirations field.
func (o *ExpiredEmplResponse) SetExpirations(v []ExpiredEmpl) {
	o.Expirations = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiredEmplResponse) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount.Get()) {
		var ret string
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiredEmplResponse) GetTotalAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *ExpiredEmplResponse) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableString and assigns it to the TotalAmount field.
func (o *ExpiredEmplResponse) SetTotalAmount(v string) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *ExpiredEmplResponse) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *ExpiredEmplResponse) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

func (o ExpiredEmplResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpiredEmplResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Expirations != nil {
		toSerialize["expirations"] = o.Expirations
	}
	if o.TotalAmount.IsSet() {
		toSerialize["total_amount"] = o.TotalAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpiredEmplResponse) UnmarshalJSON(data []byte) (err error) {
	varExpiredEmplResponse := _ExpiredEmplResponse{}

	err = json.Unmarshal(data, &varExpiredEmplResponse)

	if err != nil {
		return err
	}

	*o = ExpiredEmplResponse(varExpiredEmplResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirations")
		delete(additionalProperties, "total_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpiredEmplResponse struct {
	value *ExpiredEmplResponse
	isSet bool
}

func (v NullableExpiredEmplResponse) Get() *ExpiredEmplResponse {
	return v.value
}

func (v *NullableExpiredEmplResponse) Set(val *ExpiredEmplResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiredEmplResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiredEmplResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiredEmplResponse(val *ExpiredEmplResponse) *NullableExpiredEmplResponse {
	return &NullableExpiredEmplResponse{value: val, isSet: true}
}

func (v NullableExpiredEmplResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiredEmplResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


