
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftReceivedTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftReceivedTransactionResponse{}

// GiftReceivedTransactionResponse struct for GiftReceivedTransactionResponse
type GiftReceivedTransactionResponse struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Gifts []TransactionGiftReceived `json:"gifts,omitempty"`
	Receiver NullableRealmUser `json:"receiver,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftReceivedTransactionResponse GiftReceivedTransactionResponse

// NewGiftReceivedTransactionResponse instantiates a new GiftReceivedTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftReceivedTransactionResponse() *GiftReceivedTransactionResponse {
	this := GiftReceivedTransactionResponse{}
	return &this
}

// NewGiftReceivedTransactionResponseWithDefaults instantiates a new GiftReceivedTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftReceivedTransactionResponseWithDefaults() *GiftReceivedTransactionResponse {
	this := GiftReceivedTransactionResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedTransactionResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedTransactionResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GiftReceivedTransactionResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *GiftReceivedTransactionResponse) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GiftReceivedTransactionResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GiftReceivedTransactionResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedTransactionResponse) GetGifts() []TransactionGiftReceived {
	if o == nil {
		var ret []TransactionGiftReceived
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedTransactionResponse) GetGiftsOk() ([]TransactionGiftReceived, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *GiftReceivedTransactionResponse) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []TransactionGiftReceived and assigns it to the Gifts field.
func (o *GiftReceivedTransactionResponse) SetGifts(v []TransactionGiftReceived) {
	o.Gifts = v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedTransactionResponse) GetReceiver() RealmUser {
	if o == nil || IsNil(o.Receiver.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Receiver.Get()
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedTransactionResponse) GetReceiverOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Receiver.Get(), o.Receiver.IsSet()
}

// HasReceiver returns a boolean if a field has been set.
func (o *GiftReceivedTransactionResponse) HasReceiver() bool {
	if o != nil && o.Receiver.IsSet() {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given NullableRealmUser and assigns it to the Receiver field.
func (o *GiftReceivedTransactionResponse) SetReceiver(v RealmUser) {
	o.Receiver.Set(&v)
}
// SetReceiverNil sets the value for Receiver to be an explicit nil
func (o *GiftReceivedTransactionResponse) SetReceiverNil() {
	o.Receiver.Set(nil)
}

// UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
func (o *GiftReceivedTransactionResponse) UnsetReceiver() {
	o.Receiver.Unset()
}

func (o GiftReceivedTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftReceivedTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.Receiver.IsSet() {
		toSerialize["receiver"] = o.Receiver.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftReceivedTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftReceivedTransactionResponse := _GiftReceivedTransactionResponse{}

	err = json.Unmarshal(data, &varGiftReceivedTransactionResponse)

	if err != nil {
		return err
	}

	*o = GiftReceivedTransactionResponse(varGiftReceivedTransactionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "receiver")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftReceivedTransactionResponse struct {
	value *GiftReceivedTransactionResponse
	isSet bool
}

func (v NullableGiftReceivedTransactionResponse) Get() *GiftReceivedTransactionResponse {
	return v.value
}

func (v *NullableGiftReceivedTransactionResponse) Set(val *GiftReceivedTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftReceivedTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftReceivedTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftReceivedTransactionResponse(val *GiftReceivedTransactionResponse) *NullableGiftReceivedTransactionResponse {
	return &NullableGiftReceivedTransactionResponse{value: val, isSet: true}
}

func (v NullableGiftReceivedTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftReceivedTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


