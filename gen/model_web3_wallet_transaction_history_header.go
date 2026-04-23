
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletTransactionHistoryHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletTransactionHistoryHeader{}

// Web3WalletTransactionHistoryHeader struct for Web3WalletTransactionHistoryHeader
type Web3WalletTransactionHistoryHeader struct {
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletTransactionHistoryHeader Web3WalletTransactionHistoryHeader

// NewWeb3WalletTransactionHistoryHeader instantiates a new Web3WalletTransactionHistoryHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletTransactionHistoryHeader() *Web3WalletTransactionHistoryHeader {
	this := Web3WalletTransactionHistoryHeader{}
	return &this
}

// NewWeb3WalletTransactionHistoryHeaderWithDefaults instantiates a new Web3WalletTransactionHistoryHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletTransactionHistoryHeaderWithDefaults() *Web3WalletTransactionHistoryHeader {
	this := Web3WalletTransactionHistoryHeader{}
	return &this
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistoryHeader) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistoryHeader) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistoryHeader) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *Web3WalletTransactionHistoryHeader) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *Web3WalletTransactionHistoryHeader) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *Web3WalletTransactionHistoryHeader) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o Web3WalletTransactionHistoryHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletTransactionHistoryHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletTransactionHistoryHeader) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletTransactionHistoryHeader := _Web3WalletTransactionHistoryHeader{}

	err = json.Unmarshal(data, &varWeb3WalletTransactionHistoryHeader)

	if err != nil {
		return err
	}

	*o = Web3WalletTransactionHistoryHeader(varWeb3WalletTransactionHistoryHeader)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletTransactionHistoryHeader struct {
	value *Web3WalletTransactionHistoryHeader
	isSet bool
}

func (v NullableWeb3WalletTransactionHistoryHeader) Get() *Web3WalletTransactionHistoryHeader {
	return v.value
}

func (v *NullableWeb3WalletTransactionHistoryHeader) Set(val *Web3WalletTransactionHistoryHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletTransactionHistoryHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletTransactionHistoryHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletTransactionHistoryHeader(val *Web3WalletTransactionHistoryHeader) *NullableWeb3WalletTransactionHistoryHeader {
	return &NullableWeb3WalletTransactionHistoryHeader{value: val, isSet: true}
}

func (v NullableWeb3WalletTransactionHistoryHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletTransactionHistoryHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


