
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelWeb3WalletExternalWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelWeb3WalletExternalWallet{}

// ModelWeb3WalletExternalWallet struct for ModelWeb3WalletExternalWallet
type ModelWeb3WalletExternalWallet struct {
	EggsCount NullableInt32 `json:"eggs_count,omitempty"`
	PalsCount NullableInt32 `json:"pals_count,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelWeb3WalletExternalWallet ModelWeb3WalletExternalWallet

// NewModelWeb3WalletExternalWallet instantiates a new ModelWeb3WalletExternalWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelWeb3WalletExternalWallet() *ModelWeb3WalletExternalWallet {
	this := ModelWeb3WalletExternalWallet{}
	return &this
}

// NewModelWeb3WalletExternalWalletWithDefaults instantiates a new ModelWeb3WalletExternalWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWeb3WalletExternalWalletWithDefaults() *ModelWeb3WalletExternalWallet {
	this := ModelWeb3WalletExternalWallet{}
	return &this
}

// GetEggsCount returns the EggsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletExternalWallet) GetEggsCount() int32 {
	if o == nil || IsNil(o.EggsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.EggsCount.Get()
}

// GetEggsCountOk returns a tuple with the EggsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletExternalWallet) GetEggsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EggsCount.Get(), o.EggsCount.IsSet()
}

// HasEggsCount returns a boolean if a field has been set.
func (o *ModelWeb3WalletExternalWallet) HasEggsCount() bool {
	if o != nil && o.EggsCount.IsSet() {
		return true
	}

	return false
}

// SetEggsCount gets a reference to the given NullableInt32 and assigns it to the EggsCount field.
func (o *ModelWeb3WalletExternalWallet) SetEggsCount(v int32) {
	o.EggsCount.Set(&v)
}
// SetEggsCountNil sets the value for EggsCount to be an explicit nil
func (o *ModelWeb3WalletExternalWallet) SetEggsCountNil() {
	o.EggsCount.Set(nil)
}

// UnsetEggsCount ensures that no value is present for EggsCount, not even an explicit nil
func (o *ModelWeb3WalletExternalWallet) UnsetEggsCount() {
	o.EggsCount.Unset()
}

// GetPalsCount returns the PalsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletExternalWallet) GetPalsCount() int32 {
	if o == nil || IsNil(o.PalsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PalsCount.Get()
}

// GetPalsCountOk returns a tuple with the PalsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletExternalWallet) GetPalsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalsCount.Get(), o.PalsCount.IsSet()
}

// HasPalsCount returns a boolean if a field has been set.
func (o *ModelWeb3WalletExternalWallet) HasPalsCount() bool {
	if o != nil && o.PalsCount.IsSet() {
		return true
	}

	return false
}

// SetPalsCount gets a reference to the given NullableInt32 and assigns it to the PalsCount field.
func (o *ModelWeb3WalletExternalWallet) SetPalsCount(v int32) {
	o.PalsCount.Set(&v)
}
// SetPalsCountNil sets the value for PalsCount to be an explicit nil
func (o *ModelWeb3WalletExternalWallet) SetPalsCountNil() {
	o.PalsCount.Set(nil)
}

// UnsetPalsCount ensures that no value is present for PalsCount, not even an explicit nil
func (o *ModelWeb3WalletExternalWallet) UnsetPalsCount() {
	o.PalsCount.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletExternalWallet) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletExternalWallet) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *ModelWeb3WalletExternalWallet) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *ModelWeb3WalletExternalWallet) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *ModelWeb3WalletExternalWallet) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *ModelWeb3WalletExternalWallet) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o ModelWeb3WalletExternalWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelWeb3WalletExternalWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EggsCount.IsSet() {
		toSerialize["eggs_count"] = o.EggsCount.Get()
	}
	if o.PalsCount.IsSet() {
		toSerialize["pals_count"] = o.PalsCount.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelWeb3WalletExternalWallet) UnmarshalJSON(data []byte) (err error) {
	varModelWeb3WalletExternalWallet := _ModelWeb3WalletExternalWallet{}

	err = json.Unmarshal(data, &varModelWeb3WalletExternalWallet)

	if err != nil {
		return err
	}

	*o = ModelWeb3WalletExternalWallet(varModelWeb3WalletExternalWallet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eggs_count")
		delete(additionalProperties, "pals_count")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelWeb3WalletExternalWallet struct {
	value *ModelWeb3WalletExternalWallet
	isSet bool
}

func (v NullableModelWeb3WalletExternalWallet) Get() *ModelWeb3WalletExternalWallet {
	return v.value
}

func (v *NullableModelWeb3WalletExternalWallet) Set(val *ModelWeb3WalletExternalWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableModelWeb3WalletExternalWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableModelWeb3WalletExternalWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelWeb3WalletExternalWallet(val *ModelWeb3WalletExternalWallet) *NullableModelWeb3WalletExternalWallet {
	return &NullableModelWeb3WalletExternalWallet{value: val, isSet: true}
}

func (v NullableModelWeb3WalletExternalWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelWeb3WalletExternalWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


