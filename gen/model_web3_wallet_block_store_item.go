
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletBlockStoreItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletBlockStoreItem{}

// Web3WalletBlockStoreItem struct for Web3WalletBlockStoreItem
type Web3WalletBlockStoreItem struct {
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	Name NullableString `json:"name,omitempty"`
	PrivateKey NullableString `json:"private_key,omitempty"`
	SeedPhrase NullableString `json:"seed_phrase,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletBlockStoreItem Web3WalletBlockStoreItem

// NewWeb3WalletBlockStoreItem instantiates a new Web3WalletBlockStoreItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletBlockStoreItem() *Web3WalletBlockStoreItem {
	this := Web3WalletBlockStoreItem{}
	return &this
}

// NewWeb3WalletBlockStoreItemWithDefaults instantiates a new Web3WalletBlockStoreItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletBlockStoreItemWithDefaults() *Web3WalletBlockStoreItem {
	this := Web3WalletBlockStoreItem{}
	return &this
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreItem) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreItem) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreItem) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *Web3WalletBlockStoreItem) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *Web3WalletBlockStoreItem) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *Web3WalletBlockStoreItem) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Web3WalletBlockStoreItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Web3WalletBlockStoreItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Web3WalletBlockStoreItem) UnsetName() {
	o.Name.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreItem) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreItem) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreItem) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableString and assigns it to the PrivateKey field.
func (o *Web3WalletBlockStoreItem) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}
// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *Web3WalletBlockStoreItem) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *Web3WalletBlockStoreItem) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

// GetSeedPhrase returns the SeedPhrase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreItem) GetSeedPhrase() string {
	if o == nil || IsNil(o.SeedPhrase.Get()) {
		var ret string
		return ret
	}
	return *o.SeedPhrase.Get()
}

// GetSeedPhraseOk returns a tuple with the SeedPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreItem) GetSeedPhraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeedPhrase.Get(), o.SeedPhrase.IsSet()
}

// HasSeedPhrase returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreItem) HasSeedPhrase() bool {
	if o != nil && o.SeedPhrase.IsSet() {
		return true
	}

	return false
}

// SetSeedPhrase gets a reference to the given NullableString and assigns it to the SeedPhrase field.
func (o *Web3WalletBlockStoreItem) SetSeedPhrase(v string) {
	o.SeedPhrase.Set(&v)
}
// SetSeedPhraseNil sets the value for SeedPhrase to be an explicit nil
func (o *Web3WalletBlockStoreItem) SetSeedPhraseNil() {
	o.SeedPhrase.Set(nil)
}

// UnsetSeedPhrase ensures that no value is present for SeedPhrase, not even an explicit nil
func (o *Web3WalletBlockStoreItem) UnsetSeedPhrase() {
	o.SeedPhrase.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreItem) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreItem) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreItem) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *Web3WalletBlockStoreItem) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *Web3WalletBlockStoreItem) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *Web3WalletBlockStoreItem) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o Web3WalletBlockStoreItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletBlockStoreItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["private_key"] = o.PrivateKey.Get()
	}
	if o.SeedPhrase.IsSet() {
		toSerialize["seed_phrase"] = o.SeedPhrase.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletBlockStoreItem) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletBlockStoreItem := _Web3WalletBlockStoreItem{}

	err = json.Unmarshal(data, &varWeb3WalletBlockStoreItem)

	if err != nil {
		return err
	}

	*o = Web3WalletBlockStoreItem(varWeb3WalletBlockStoreItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "name")
		delete(additionalProperties, "private_key")
		delete(additionalProperties, "seed_phrase")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletBlockStoreItem struct {
	value *Web3WalletBlockStoreItem
	isSet bool
}

func (v NullableWeb3WalletBlockStoreItem) Get() *Web3WalletBlockStoreItem {
	return v.value
}

func (v *NullableWeb3WalletBlockStoreItem) Set(val *Web3WalletBlockStoreItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletBlockStoreItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletBlockStoreItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletBlockStoreItem(val *Web3WalletBlockStoreItem) *NullableWeb3WalletBlockStoreItem {
	return &NullableWeb3WalletBlockStoreItem{value: val, isSet: true}
}

func (v NullableWeb3WalletBlockStoreItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletBlockStoreItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


