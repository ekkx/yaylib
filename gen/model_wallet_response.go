
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the WalletResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletResponse{}

// WalletResponse struct for WalletResponse
type WalletResponse struct {
	Coins NullableCoinAmount `json:"coins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletResponse WalletResponse

// NewWalletResponse instantiates a new WalletResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletResponse() *WalletResponse {
	this := WalletResponse{}
	return &this
}

// NewWalletResponseWithDefaults instantiates a new WalletResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletResponseWithDefaults() *WalletResponse {
	this := WalletResponse{}
	return &this
}

// GetCoins returns the Coins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletResponse) GetCoins() CoinAmount {
	if o == nil || IsNil(o.Coins.Get()) {
		var ret CoinAmount
		return ret
	}
	return *o.Coins.Get()
}

// GetCoinsOk returns a tuple with the Coins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletResponse) GetCoinsOk() (*CoinAmount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coins.Get(), o.Coins.IsSet()
}

// HasCoins returns a boolean if a field has been set.
func (o *WalletResponse) HasCoins() bool {
	if o != nil && o.Coins.IsSet() {
		return true
	}

	return false
}

// SetCoins gets a reference to the given NullableCoinAmount and assigns it to the Coins field.
func (o *WalletResponse) SetCoins(v CoinAmount) {
	o.Coins.Set(&v)
}
// SetCoinsNil sets the value for Coins to be an explicit nil
func (o *WalletResponse) SetCoinsNil() {
	o.Coins.Set(nil)
}

// UnsetCoins ensures that no value is present for Coins, not even an explicit nil
func (o *WalletResponse) UnsetCoins() {
	o.Coins.Unset()
}

func (o WalletResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Coins.IsSet() {
		toSerialize["coins"] = o.Coins.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WalletResponse) UnmarshalJSON(data []byte) (err error) {
	varWalletResponse := _WalletResponse{}

	err = json.Unmarshal(data, &varWalletResponse)

	if err != nil {
		return err
	}

	*o = WalletResponse(varWalletResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletResponse struct {
	value *WalletResponse
	isSet bool
}

func (v NullableWalletResponse) Get() *WalletResponse {
	return v.value
}

func (v *NullableWalletResponse) Set(val *WalletResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletResponse(val *WalletResponse) *NullableWalletResponse {
	return &NullableWalletResponse{value: val, isSet: true}
}

func (v NullableWalletResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


