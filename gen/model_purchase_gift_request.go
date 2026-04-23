
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PurchaseGiftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseGiftRequest{}

// PurchaseGiftRequest struct for PurchaseGiftRequest
type PurchaseGiftRequest struct {
	Gifts map[string]int32 `json:"gifts,omitempty"`
	IsAnonymously NullableBool `json:"is_anonymously,omitempty"`
	ReceiverId NullableInt64 `json:"receiver_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PurchaseGiftRequest PurchaseGiftRequest

// NewPurchaseGiftRequest instantiates a new PurchaseGiftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseGiftRequest() *PurchaseGiftRequest {
	this := PurchaseGiftRequest{}
	return &this
}

// NewPurchaseGiftRequestWithDefaults instantiates a new PurchaseGiftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseGiftRequestWithDefaults() *PurchaseGiftRequest {
	this := PurchaseGiftRequest{}
	return &this
}

// GetGifts returns the Gifts field value if set, zero value otherwise.
func (o *PurchaseGiftRequest) GetGifts() map[string]int32 {
	if o == nil || IsNil(o.Gifts) {
		var ret map[string]int32
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseGiftRequest) GetGiftsOk() (map[string]int32, bool) {
	if o == nil || IsNil(o.Gifts) {
		return map[string]int32{}, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *PurchaseGiftRequest) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given map[string]int32 and assigns it to the Gifts field.
func (o *PurchaseGiftRequest) SetGifts(v map[string]int32) {
	o.Gifts = v
}

// GetIsAnonymously returns the IsAnonymously field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseGiftRequest) GetIsAnonymously() bool {
	if o == nil || IsNil(o.IsAnonymously.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAnonymously.Get()
}

// GetIsAnonymouslyOk returns a tuple with the IsAnonymously field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseGiftRequest) GetIsAnonymouslyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAnonymously.Get(), o.IsAnonymously.IsSet()
}

// HasIsAnonymously returns a boolean if a field has been set.
func (o *PurchaseGiftRequest) HasIsAnonymously() bool {
	if o != nil && o.IsAnonymously.IsSet() {
		return true
	}

	return false
}

// SetIsAnonymously gets a reference to the given NullableBool and assigns it to the IsAnonymously field.
func (o *PurchaseGiftRequest) SetIsAnonymously(v bool) {
	o.IsAnonymously.Set(&v)
}
// SetIsAnonymouslyNil sets the value for IsAnonymously to be an explicit nil
func (o *PurchaseGiftRequest) SetIsAnonymouslyNil() {
	o.IsAnonymously.Set(nil)
}

// UnsetIsAnonymously ensures that no value is present for IsAnonymously, not even an explicit nil
func (o *PurchaseGiftRequest) UnsetIsAnonymously() {
	o.IsAnonymously.Unset()
}

// GetReceiverId returns the ReceiverId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseGiftRequest) GetReceiverId() int64 {
	if o == nil || IsNil(o.ReceiverId.Get()) {
		var ret int64
		return ret
	}
	return *o.ReceiverId.Get()
}

// GetReceiverIdOk returns a tuple with the ReceiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseGiftRequest) GetReceiverIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverId.Get(), o.ReceiverId.IsSet()
}

// HasReceiverId returns a boolean if a field has been set.
func (o *PurchaseGiftRequest) HasReceiverId() bool {
	if o != nil && o.ReceiverId.IsSet() {
		return true
	}

	return false
}

// SetReceiverId gets a reference to the given NullableInt64 and assigns it to the ReceiverId field.
func (o *PurchaseGiftRequest) SetReceiverId(v int64) {
	o.ReceiverId.Set(&v)
}
// SetReceiverIdNil sets the value for ReceiverId to be an explicit nil
func (o *PurchaseGiftRequest) SetReceiverIdNil() {
	o.ReceiverId.Set(nil)
}

// UnsetReceiverId ensures that no value is present for ReceiverId, not even an explicit nil
func (o *PurchaseGiftRequest) UnsetReceiverId() {
	o.ReceiverId.Unset()
}

func (o PurchaseGiftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseGiftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gifts) {
		toSerialize["gifts"] = o.Gifts
	}
	if o.IsAnonymously.IsSet() {
		toSerialize["is_anonymously"] = o.IsAnonymously.Get()
	}
	if o.ReceiverId.IsSet() {
		toSerialize["receiver_id"] = o.ReceiverId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PurchaseGiftRequest) UnmarshalJSON(data []byte) (err error) {
	varPurchaseGiftRequest := _PurchaseGiftRequest{}

	err = json.Unmarshal(data, &varPurchaseGiftRequest)

	if err != nil {
		return err
	}

	*o = PurchaseGiftRequest(varPurchaseGiftRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "is_anonymously")
		delete(additionalProperties, "receiver_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePurchaseGiftRequest struct {
	value *PurchaseGiftRequest
	isSet bool
}

func (v NullablePurchaseGiftRequest) Get() *PurchaseGiftRequest {
	return v.value
}

func (v *NullablePurchaseGiftRequest) Set(val *PurchaseGiftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseGiftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseGiftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseGiftRequest(val *PurchaseGiftRequest) *NullablePurchaseGiftRequest {
	return &NullablePurchaseGiftRequest{value: val, isSet: true}
}

func (v NullablePurchaseGiftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseGiftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


