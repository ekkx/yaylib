
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network struct for Network
type Network struct {
	AdditionGasPercent NullableModelWeb3WalletGasPercent `json:"addition_gas_percent,omitempty"`
	AmmAddress NullableString `json:"amm_address,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	QuoterAddress NullableString `json:"quoter_address,omitempty"`
	Url NullableString `json:"url,omitempty"`
	WethAddress NullableString `json:"weth_address,omitempty"`
	YayAddress NullableString `json:"yay_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetAdditionGasPercent returns the AdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetAdditionGasPercent() ModelWeb3WalletGasPercent {
	if o == nil || IsNil(o.AdditionGasPercent.Get()) {
		var ret ModelWeb3WalletGasPercent
		return ret
	}
	return *o.AdditionGasPercent.Get()
}

// GetAdditionGasPercentOk returns a tuple with the AdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetAdditionGasPercentOk() (*ModelWeb3WalletGasPercent, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionGasPercent.Get(), o.AdditionGasPercent.IsSet()
}

// HasAdditionGasPercent returns a boolean if a field has been set.
func (o *Network) HasAdditionGasPercent() bool {
	if o != nil && o.AdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetAdditionGasPercent gets a reference to the given NullableModelWeb3WalletGasPercent and assigns it to the AdditionGasPercent field.
func (o *Network) SetAdditionGasPercent(v ModelWeb3WalletGasPercent) {
	o.AdditionGasPercent.Set(&v)
}
// SetAdditionGasPercentNil sets the value for AdditionGasPercent to be an explicit nil
func (o *Network) SetAdditionGasPercentNil() {
	o.AdditionGasPercent.Set(nil)
}

// UnsetAdditionGasPercent ensures that no value is present for AdditionGasPercent, not even an explicit nil
func (o *Network) UnsetAdditionGasPercent() {
	o.AdditionGasPercent.Unset()
}

// GetAmmAddress returns the AmmAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetAmmAddress() string {
	if o == nil || IsNil(o.AmmAddress.Get()) {
		var ret string
		return ret
	}
	return *o.AmmAddress.Get()
}

// GetAmmAddressOk returns a tuple with the AmmAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetAmmAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmmAddress.Get(), o.AmmAddress.IsSet()
}

// HasAmmAddress returns a boolean if a field has been set.
func (o *Network) HasAmmAddress() bool {
	if o != nil && o.AmmAddress.IsSet() {
		return true
	}

	return false
}

// SetAmmAddress gets a reference to the given NullableString and assigns it to the AmmAddress field.
func (o *Network) SetAmmAddress(v string) {
	o.AmmAddress.Set(&v)
}
// SetAmmAddressNil sets the value for AmmAddress to be an explicit nil
func (o *Network) SetAmmAddressNil() {
	o.AmmAddress.Set(nil)
}

// UnsetAmmAddress ensures that no value is present for AmmAddress, not even an explicit nil
func (o *Network) UnsetAmmAddress() {
	o.AmmAddress.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Network) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Network) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Network) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Network) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Network) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Network) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Network) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Network) UnsetName() {
	o.Name.Unset()
}

// GetQuoterAddress returns the QuoterAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetQuoterAddress() string {
	if o == nil || IsNil(o.QuoterAddress.Get()) {
		var ret string
		return ret
	}
	return *o.QuoterAddress.Get()
}

// GetQuoterAddressOk returns a tuple with the QuoterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetQuoterAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuoterAddress.Get(), o.QuoterAddress.IsSet()
}

// HasQuoterAddress returns a boolean if a field has been set.
func (o *Network) HasQuoterAddress() bool {
	if o != nil && o.QuoterAddress.IsSet() {
		return true
	}

	return false
}

// SetQuoterAddress gets a reference to the given NullableString and assigns it to the QuoterAddress field.
func (o *Network) SetQuoterAddress(v string) {
	o.QuoterAddress.Set(&v)
}
// SetQuoterAddressNil sets the value for QuoterAddress to be an explicit nil
func (o *Network) SetQuoterAddressNil() {
	o.QuoterAddress.Set(nil)
}

// UnsetQuoterAddress ensures that no value is present for QuoterAddress, not even an explicit nil
func (o *Network) UnsetQuoterAddress() {
	o.QuoterAddress.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Network) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Network) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Network) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Network) UnsetUrl() {
	o.Url.Unset()
}

// GetWethAddress returns the WethAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetWethAddress() string {
	if o == nil || IsNil(o.WethAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WethAddress.Get()
}

// GetWethAddressOk returns a tuple with the WethAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetWethAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WethAddress.Get(), o.WethAddress.IsSet()
}

// HasWethAddress returns a boolean if a field has been set.
func (o *Network) HasWethAddress() bool {
	if o != nil && o.WethAddress.IsSet() {
		return true
	}

	return false
}

// SetWethAddress gets a reference to the given NullableString and assigns it to the WethAddress field.
func (o *Network) SetWethAddress(v string) {
	o.WethAddress.Set(&v)
}
// SetWethAddressNil sets the value for WethAddress to be an explicit nil
func (o *Network) SetWethAddressNil() {
	o.WethAddress.Set(nil)
}

// UnsetWethAddress ensures that no value is present for WethAddress, not even an explicit nil
func (o *Network) UnsetWethAddress() {
	o.WethAddress.Unset()
}

// GetYayAddress returns the YayAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetYayAddress() string {
	if o == nil || IsNil(o.YayAddress.Get()) {
		var ret string
		return ret
	}
	return *o.YayAddress.Get()
}

// GetYayAddressOk returns a tuple with the YayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetYayAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YayAddress.Get(), o.YayAddress.IsSet()
}

// HasYayAddress returns a boolean if a field has been set.
func (o *Network) HasYayAddress() bool {
	if o != nil && o.YayAddress.IsSet() {
		return true
	}

	return false
}

// SetYayAddress gets a reference to the given NullableString and assigns it to the YayAddress field.
func (o *Network) SetYayAddress(v string) {
	o.YayAddress.Set(&v)
}
// SetYayAddressNil sets the value for YayAddress to be an explicit nil
func (o *Network) SetYayAddressNil() {
	o.YayAddress.Set(nil)
}

// UnsetYayAddress ensures that no value is present for YayAddress, not even an explicit nil
func (o *Network) UnsetYayAddress() {
	o.YayAddress.Unset()
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionGasPercent.IsSet() {
		toSerialize["addition_gas_percent"] = o.AdditionGasPercent.Get()
	}
	if o.AmmAddress.IsSet() {
		toSerialize["amm_address"] = o.AmmAddress.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.QuoterAddress.IsSet() {
		toSerialize["quoter_address"] = o.QuoterAddress.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.WethAddress.IsSet() {
		toSerialize["weth_address"] = o.WethAddress.Get()
	}
	if o.YayAddress.IsSet() {
		toSerialize["yay_address"] = o.YayAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Network) UnmarshalJSON(data []byte) (err error) {
	varNetwork := _Network{}

	err = json.Unmarshal(data, &varNetwork)

	if err != nil {
		return err
	}

	*o = Network(varNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addition_gas_percent")
		delete(additionalProperties, "amm_address")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "quoter_address")
		delete(additionalProperties, "url")
		delete(additionalProperties, "weth_address")
		delete(additionalProperties, "yay_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


