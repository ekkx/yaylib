
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftCardActivityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardActivityDetails{}

// GiftCardActivityDetails struct for GiftCardActivityDetails
type GiftCardActivityDetails struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Email NullableString `json:"email,omitempty"`
	EmplAmount NullableInt32 `json:"empl_amount,omitempty"`
	GiftCardAmount NullableInt32 `json:"gift_card_amount,omitempty"`
	// Unresolved Java type: ao.h
	Status map[string]interface{} `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftCardActivityDetails GiftCardActivityDetails

// NewGiftCardActivityDetails instantiates a new GiftCardActivityDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardActivityDetails() *GiftCardActivityDetails {
	this := GiftCardActivityDetails{}
	return &this
}

// NewGiftCardActivityDetailsWithDefaults instantiates a new GiftCardActivityDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardActivityDetailsWithDefaults() *GiftCardActivityDetails {
	this := GiftCardActivityDetails{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardActivityDetails) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardActivityDetails) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GiftCardActivityDetails) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *GiftCardActivityDetails) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GiftCardActivityDetails) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GiftCardActivityDetails) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardActivityDetails) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardActivityDetails) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *GiftCardActivityDetails) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *GiftCardActivityDetails) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *GiftCardActivityDetails) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *GiftCardActivityDetails) UnsetEmail() {
	o.Email.Unset()
}

// GetEmplAmount returns the EmplAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardActivityDetails) GetEmplAmount() int32 {
	if o == nil || IsNil(o.EmplAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.EmplAmount.Get()
}

// GetEmplAmountOk returns a tuple with the EmplAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardActivityDetails) GetEmplAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplAmount.Get(), o.EmplAmount.IsSet()
}

// HasEmplAmount returns a boolean if a field has been set.
func (o *GiftCardActivityDetails) HasEmplAmount() bool {
	if o != nil && o.EmplAmount.IsSet() {
		return true
	}

	return false
}

// SetEmplAmount gets a reference to the given NullableInt32 and assigns it to the EmplAmount field.
func (o *GiftCardActivityDetails) SetEmplAmount(v int32) {
	o.EmplAmount.Set(&v)
}
// SetEmplAmountNil sets the value for EmplAmount to be an explicit nil
func (o *GiftCardActivityDetails) SetEmplAmountNil() {
	o.EmplAmount.Set(nil)
}

// UnsetEmplAmount ensures that no value is present for EmplAmount, not even an explicit nil
func (o *GiftCardActivityDetails) UnsetEmplAmount() {
	o.EmplAmount.Unset()
}

// GetGiftCardAmount returns the GiftCardAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardActivityDetails) GetGiftCardAmount() int32 {
	if o == nil || IsNil(o.GiftCardAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.GiftCardAmount.Get()
}

// GetGiftCardAmountOk returns a tuple with the GiftCardAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardActivityDetails) GetGiftCardAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftCardAmount.Get(), o.GiftCardAmount.IsSet()
}

// HasGiftCardAmount returns a boolean if a field has been set.
func (o *GiftCardActivityDetails) HasGiftCardAmount() bool {
	if o != nil && o.GiftCardAmount.IsSet() {
		return true
	}

	return false
}

// SetGiftCardAmount gets a reference to the given NullableInt32 and assigns it to the GiftCardAmount field.
func (o *GiftCardActivityDetails) SetGiftCardAmount(v int32) {
	o.GiftCardAmount.Set(&v)
}
// SetGiftCardAmountNil sets the value for GiftCardAmount to be an explicit nil
func (o *GiftCardActivityDetails) SetGiftCardAmountNil() {
	o.GiftCardAmount.Set(nil)
}

// UnsetGiftCardAmount ensures that no value is present for GiftCardAmount, not even an explicit nil
func (o *GiftCardActivityDetails) UnsetGiftCardAmount() {
	o.GiftCardAmount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardActivityDetails) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardActivityDetails) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GiftCardActivityDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *GiftCardActivityDetails) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o GiftCardActivityDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardActivityDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.EmplAmount.IsSet() {
		toSerialize["empl_amount"] = o.EmplAmount.Get()
	}
	if o.GiftCardAmount.IsSet() {
		toSerialize["gift_card_amount"] = o.GiftCardAmount.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftCardActivityDetails) UnmarshalJSON(data []byte) (err error) {
	varGiftCardActivityDetails := _GiftCardActivityDetails{}

	err = json.Unmarshal(data, &varGiftCardActivityDetails)

	if err != nil {
		return err
	}

	*o = GiftCardActivityDetails(varGiftCardActivityDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "email")
		delete(additionalProperties, "empl_amount")
		delete(additionalProperties, "gift_card_amount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftCardActivityDetails struct {
	value *GiftCardActivityDetails
	isSet bool
}

func (v NullableGiftCardActivityDetails) Get() *GiftCardActivityDetails {
	return v.value
}

func (v *NullableGiftCardActivityDetails) Set(val *GiftCardActivityDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardActivityDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardActivityDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardActivityDetails(val *GiftCardActivityDetails) *NullableGiftCardActivityDetails {
	return &NullableGiftCardActivityDetails{value: val, isSet: true}
}

func (v NullableGiftCardActivityDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardActivityDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


