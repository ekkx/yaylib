
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftTransactionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftTransactionDetail{}

// GiftTransactionDetail struct for GiftTransactionDetail
type GiftTransactionDetail struct {
	Gifts []GiftTransaction `json:"gifts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftTransactionDetail GiftTransactionDetail

// NewGiftTransactionDetail instantiates a new GiftTransactionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftTransactionDetail() *GiftTransactionDetail {
	this := GiftTransactionDetail{}
	return &this
}

// NewGiftTransactionDetailWithDefaults instantiates a new GiftTransactionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftTransactionDetailWithDefaults() *GiftTransactionDetail {
	this := GiftTransactionDetail{}
	return &this
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransactionDetail) GetGifts() []GiftTransaction {
	if o == nil {
		var ret []GiftTransaction
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransactionDetail) GetGiftsOk() ([]GiftTransaction, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *GiftTransactionDetail) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []GiftTransaction and assigns it to the Gifts field.
func (o *GiftTransactionDetail) SetGifts(v []GiftTransaction) {
	o.Gifts = v
}

func (o GiftTransactionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftTransactionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftTransactionDetail) UnmarshalJSON(data []byte) (err error) {
	varGiftTransactionDetail := _GiftTransactionDetail{}

	err = json.Unmarshal(data, &varGiftTransactionDetail)

	if err != nil {
		return err
	}

	*o = GiftTransactionDetail(varGiftTransactionDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftTransactionDetail struct {
	value *GiftTransactionDetail
	isSet bool
}

func (v NullableGiftTransactionDetail) Get() *GiftTransactionDetail {
	return v.value
}

func (v *NullableGiftTransactionDetail) Set(val *GiftTransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftTransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftTransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftTransactionDetail(val *GiftTransactionDetail) *NullableGiftTransactionDetail {
	return &NullableGiftTransactionDetail{value: val, isSet: true}
}

func (v NullableGiftTransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftTransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


