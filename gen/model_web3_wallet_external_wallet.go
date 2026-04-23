
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletExternalWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletExternalWallet{}

// Web3WalletExternalWallet struct for Web3WalletExternalWallet
type Web3WalletExternalWallet struct {
	Balance NullableBalance `json:"balance,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletExternalWallet Web3WalletExternalWallet

// NewWeb3WalletExternalWallet instantiates a new Web3WalletExternalWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletExternalWallet() *Web3WalletExternalWallet {
	this := Web3WalletExternalWallet{}
	return &this
}

// NewWeb3WalletExternalWalletWithDefaults instantiates a new Web3WalletExternalWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletExternalWalletWithDefaults() *Web3WalletExternalWallet {
	this := Web3WalletExternalWallet{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletExternalWallet) GetBalance() Balance {
	if o == nil || IsNil(o.Balance.Get()) {
		var ret Balance
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletExternalWallet) GetBalanceOk() (*Balance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *Web3WalletExternalWallet) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableBalance and assigns it to the Balance field.
func (o *Web3WalletExternalWallet) SetBalance(v Balance) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *Web3WalletExternalWallet) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *Web3WalletExternalWallet) UnsetBalance() {
	o.Balance.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletExternalWallet) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletExternalWallet) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *Web3WalletExternalWallet) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *Web3WalletExternalWallet) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *Web3WalletExternalWallet) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *Web3WalletExternalWallet) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o Web3WalletExternalWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletExternalWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletExternalWallet) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletExternalWallet := _Web3WalletExternalWallet{}

	err = json.Unmarshal(data, &varWeb3WalletExternalWallet)

	if err != nil {
		return err
	}

	*o = Web3WalletExternalWallet(varWeb3WalletExternalWallet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletExternalWallet struct {
	value *Web3WalletExternalWallet
	isSet bool
}

func (v NullableWeb3WalletExternalWallet) Get() *Web3WalletExternalWallet {
	return v.value
}

func (v *NullableWeb3WalletExternalWallet) Set(val *Web3WalletExternalWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletExternalWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletExternalWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletExternalWallet(val *Web3WalletExternalWallet) *NullableWeb3WalletExternalWallet {
	return &NullableWeb3WalletExternalWallet{value: val, isSet: true}
}

func (v NullableWeb3WalletExternalWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletExternalWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


