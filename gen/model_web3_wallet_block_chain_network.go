
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletBlockChainNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletBlockChainNetwork{}

// Web3WalletBlockChainNetwork struct for Web3WalletBlockChainNetwork
type Web3WalletBlockChainNetwork struct {
	L1AmmAddress NullableString `json:"l1_amm_address,omitempty"`
	L1BlockChainNetworkAdditionGasPercent NullableModelWeb3WalletGasPercent `json:"l1_block_chain_network_addition_gas_percent,omitempty"`
	L1BlockChainNetworkId NullableInt64 `json:"l1_block_chain_network_id,omitempty"`
	L1BlockChainNetworkName NullableString `json:"l1_block_chain_network_name,omitempty"`
	L1BlockChainNetworkUrl NullableString `json:"l1_block_chain_network_url,omitempty"`
	L1QuoterAddress NullableString `json:"l1_quoter_address,omitempty"`
	L1WethAddress NullableString `json:"l1_weth_address,omitempty"`
	L2AmmAddress NullableString `json:"l2_amm_address,omitempty"`
	L2BlockChainNetworkAdditionGasPercent NullableModelWeb3WalletGasPercent `json:"l2_block_chain_network_addition_gas_percent,omitempty"`
	L2BlockChainNetworkId NullableInt64 `json:"l2_block_chain_network_id,omitempty"`
	L2BlockChainNetworkName NullableString `json:"l2_block_chain_network_name,omitempty"`
	L2BlockChainNetworkUrl NullableString `json:"l2_block_chain_network_url,omitempty"`
	L2QuoterAddress NullableString `json:"l2_quoter_address,omitempty"`
	L2WethAddress NullableString `json:"l2_weth_address,omitempty"`
	MiscPalBaseImageUri NullableString `json:"misc_pal_base_image_uri,omitempty"`
	MiscPalBaseJsonUri NullableString `json:"misc_pal_base_json_uri,omitempty"`
	SmartContractEmplAddress NullableString `json:"smart_contract_empl_address,omitempty"`
	SmartContractEmplDepositAddress NullableString `json:"smart_contract_empl_deposit_address,omitempty"`
	SmartContractEmplWithdrawAddress NullableString `json:"smart_contract_empl_withdraw_address,omitempty"`
	SmartContractL1YayAddress NullableString `json:"smart_contract_l1_yay_address,omitempty"`
	SmartContractL2YayAddress NullableString `json:"smart_contract_l2_yay_address,omitempty"`
	SmartContractPalAddress NullableString `json:"smart_contract_pal_address,omitempty"`
	SmartContractPalDepositWithdrawAddress NullableString `json:"smart_contract_pal_deposit_withdraw_address,omitempty"`
	SmartContractX2y2delegateerc721Address NullableString `json:"smart_contract_x2y2delegateerc721_address,omitempty"`
	SmartContractX2y2marketAddress NullableString `json:"smart_contract_x2y2market_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletBlockChainNetwork Web3WalletBlockChainNetwork

// NewWeb3WalletBlockChainNetwork instantiates a new Web3WalletBlockChainNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletBlockChainNetwork() *Web3WalletBlockChainNetwork {
	this := Web3WalletBlockChainNetwork{}
	return &this
}

// NewWeb3WalletBlockChainNetworkWithDefaults instantiates a new Web3WalletBlockChainNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletBlockChainNetworkWithDefaults() *Web3WalletBlockChainNetwork {
	this := Web3WalletBlockChainNetwork{}
	return &this
}

// GetL1AmmAddress returns the L1AmmAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1AmmAddress() string {
	if o == nil || IsNil(o.L1AmmAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L1AmmAddress.Get()
}

// GetL1AmmAddressOk returns a tuple with the L1AmmAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1AmmAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1AmmAddress.Get(), o.L1AmmAddress.IsSet()
}

// HasL1AmmAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1AmmAddress() bool {
	if o != nil && o.L1AmmAddress.IsSet() {
		return true
	}

	return false
}

// SetL1AmmAddress gets a reference to the given NullableString and assigns it to the L1AmmAddress field.
func (o *Web3WalletBlockChainNetwork) SetL1AmmAddress(v string) {
	o.L1AmmAddress.Set(&v)
}
// SetL1AmmAddressNil sets the value for L1AmmAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1AmmAddressNil() {
	o.L1AmmAddress.Set(nil)
}

// UnsetL1AmmAddress ensures that no value is present for L1AmmAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1AmmAddress() {
	o.L1AmmAddress.Unset()
}

// GetL1BlockChainNetworkAdditionGasPercent returns the L1BlockChainNetworkAdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkAdditionGasPercent() ModelWeb3WalletGasPercent {
	if o == nil || IsNil(o.L1BlockChainNetworkAdditionGasPercent.Get()) {
		var ret ModelWeb3WalletGasPercent
		return ret
	}
	return *o.L1BlockChainNetworkAdditionGasPercent.Get()
}

// GetL1BlockChainNetworkAdditionGasPercentOk returns a tuple with the L1BlockChainNetworkAdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkAdditionGasPercentOk() (*ModelWeb3WalletGasPercent, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1BlockChainNetworkAdditionGasPercent.Get(), o.L1BlockChainNetworkAdditionGasPercent.IsSet()
}

// HasL1BlockChainNetworkAdditionGasPercent returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1BlockChainNetworkAdditionGasPercent() bool {
	if o != nil && o.L1BlockChainNetworkAdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetL1BlockChainNetworkAdditionGasPercent gets a reference to the given NullableModelWeb3WalletGasPercent and assigns it to the L1BlockChainNetworkAdditionGasPercent field.
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkAdditionGasPercent(v ModelWeb3WalletGasPercent) {
	o.L1BlockChainNetworkAdditionGasPercent.Set(&v)
}
// SetL1BlockChainNetworkAdditionGasPercentNil sets the value for L1BlockChainNetworkAdditionGasPercent to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkAdditionGasPercentNil() {
	o.L1BlockChainNetworkAdditionGasPercent.Set(nil)
}

// UnsetL1BlockChainNetworkAdditionGasPercent ensures that no value is present for L1BlockChainNetworkAdditionGasPercent, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1BlockChainNetworkAdditionGasPercent() {
	o.L1BlockChainNetworkAdditionGasPercent.Unset()
}

// GetL1BlockChainNetworkId returns the L1BlockChainNetworkId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkId() int64 {
	if o == nil || IsNil(o.L1BlockChainNetworkId.Get()) {
		var ret int64
		return ret
	}
	return *o.L1BlockChainNetworkId.Get()
}

// GetL1BlockChainNetworkIdOk returns a tuple with the L1BlockChainNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1BlockChainNetworkId.Get(), o.L1BlockChainNetworkId.IsSet()
}

// HasL1BlockChainNetworkId returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1BlockChainNetworkId() bool {
	if o != nil && o.L1BlockChainNetworkId.IsSet() {
		return true
	}

	return false
}

// SetL1BlockChainNetworkId gets a reference to the given NullableInt64 and assigns it to the L1BlockChainNetworkId field.
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkId(v int64) {
	o.L1BlockChainNetworkId.Set(&v)
}
// SetL1BlockChainNetworkIdNil sets the value for L1BlockChainNetworkId to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkIdNil() {
	o.L1BlockChainNetworkId.Set(nil)
}

// UnsetL1BlockChainNetworkId ensures that no value is present for L1BlockChainNetworkId, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1BlockChainNetworkId() {
	o.L1BlockChainNetworkId.Unset()
}

// GetL1BlockChainNetworkName returns the L1BlockChainNetworkName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkName() string {
	if o == nil || IsNil(o.L1BlockChainNetworkName.Get()) {
		var ret string
		return ret
	}
	return *o.L1BlockChainNetworkName.Get()
}

// GetL1BlockChainNetworkNameOk returns a tuple with the L1BlockChainNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1BlockChainNetworkName.Get(), o.L1BlockChainNetworkName.IsSet()
}

// HasL1BlockChainNetworkName returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1BlockChainNetworkName() bool {
	if o != nil && o.L1BlockChainNetworkName.IsSet() {
		return true
	}

	return false
}

// SetL1BlockChainNetworkName gets a reference to the given NullableString and assigns it to the L1BlockChainNetworkName field.
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkName(v string) {
	o.L1BlockChainNetworkName.Set(&v)
}
// SetL1BlockChainNetworkNameNil sets the value for L1BlockChainNetworkName to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkNameNil() {
	o.L1BlockChainNetworkName.Set(nil)
}

// UnsetL1BlockChainNetworkName ensures that no value is present for L1BlockChainNetworkName, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1BlockChainNetworkName() {
	o.L1BlockChainNetworkName.Unset()
}

// GetL1BlockChainNetworkUrl returns the L1BlockChainNetworkUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkUrl() string {
	if o == nil || IsNil(o.L1BlockChainNetworkUrl.Get()) {
		var ret string
		return ret
	}
	return *o.L1BlockChainNetworkUrl.Get()
}

// GetL1BlockChainNetworkUrlOk returns a tuple with the L1BlockChainNetworkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1BlockChainNetworkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1BlockChainNetworkUrl.Get(), o.L1BlockChainNetworkUrl.IsSet()
}

// HasL1BlockChainNetworkUrl returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1BlockChainNetworkUrl() bool {
	if o != nil && o.L1BlockChainNetworkUrl.IsSet() {
		return true
	}

	return false
}

// SetL1BlockChainNetworkUrl gets a reference to the given NullableString and assigns it to the L1BlockChainNetworkUrl field.
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkUrl(v string) {
	o.L1BlockChainNetworkUrl.Set(&v)
}
// SetL1BlockChainNetworkUrlNil sets the value for L1BlockChainNetworkUrl to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1BlockChainNetworkUrlNil() {
	o.L1BlockChainNetworkUrl.Set(nil)
}

// UnsetL1BlockChainNetworkUrl ensures that no value is present for L1BlockChainNetworkUrl, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1BlockChainNetworkUrl() {
	o.L1BlockChainNetworkUrl.Unset()
}

// GetL1QuoterAddress returns the L1QuoterAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1QuoterAddress() string {
	if o == nil || IsNil(o.L1QuoterAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L1QuoterAddress.Get()
}

// GetL1QuoterAddressOk returns a tuple with the L1QuoterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1QuoterAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1QuoterAddress.Get(), o.L1QuoterAddress.IsSet()
}

// HasL1QuoterAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1QuoterAddress() bool {
	if o != nil && o.L1QuoterAddress.IsSet() {
		return true
	}

	return false
}

// SetL1QuoterAddress gets a reference to the given NullableString and assigns it to the L1QuoterAddress field.
func (o *Web3WalletBlockChainNetwork) SetL1QuoterAddress(v string) {
	o.L1QuoterAddress.Set(&v)
}
// SetL1QuoterAddressNil sets the value for L1QuoterAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1QuoterAddressNil() {
	o.L1QuoterAddress.Set(nil)
}

// UnsetL1QuoterAddress ensures that no value is present for L1QuoterAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1QuoterAddress() {
	o.L1QuoterAddress.Unset()
}

// GetL1WethAddress returns the L1WethAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL1WethAddress() string {
	if o == nil || IsNil(o.L1WethAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L1WethAddress.Get()
}

// GetL1WethAddressOk returns a tuple with the L1WethAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL1WethAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1WethAddress.Get(), o.L1WethAddress.IsSet()
}

// HasL1WethAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL1WethAddress() bool {
	if o != nil && o.L1WethAddress.IsSet() {
		return true
	}

	return false
}

// SetL1WethAddress gets a reference to the given NullableString and assigns it to the L1WethAddress field.
func (o *Web3WalletBlockChainNetwork) SetL1WethAddress(v string) {
	o.L1WethAddress.Set(&v)
}
// SetL1WethAddressNil sets the value for L1WethAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL1WethAddressNil() {
	o.L1WethAddress.Set(nil)
}

// UnsetL1WethAddress ensures that no value is present for L1WethAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL1WethAddress() {
	o.L1WethAddress.Unset()
}

// GetL2AmmAddress returns the L2AmmAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2AmmAddress() string {
	if o == nil || IsNil(o.L2AmmAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L2AmmAddress.Get()
}

// GetL2AmmAddressOk returns a tuple with the L2AmmAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2AmmAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2AmmAddress.Get(), o.L2AmmAddress.IsSet()
}

// HasL2AmmAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2AmmAddress() bool {
	if o != nil && o.L2AmmAddress.IsSet() {
		return true
	}

	return false
}

// SetL2AmmAddress gets a reference to the given NullableString and assigns it to the L2AmmAddress field.
func (o *Web3WalletBlockChainNetwork) SetL2AmmAddress(v string) {
	o.L2AmmAddress.Set(&v)
}
// SetL2AmmAddressNil sets the value for L2AmmAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2AmmAddressNil() {
	o.L2AmmAddress.Set(nil)
}

// UnsetL2AmmAddress ensures that no value is present for L2AmmAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2AmmAddress() {
	o.L2AmmAddress.Unset()
}

// GetL2BlockChainNetworkAdditionGasPercent returns the L2BlockChainNetworkAdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkAdditionGasPercent() ModelWeb3WalletGasPercent {
	if o == nil || IsNil(o.L2BlockChainNetworkAdditionGasPercent.Get()) {
		var ret ModelWeb3WalletGasPercent
		return ret
	}
	return *o.L2BlockChainNetworkAdditionGasPercent.Get()
}

// GetL2BlockChainNetworkAdditionGasPercentOk returns a tuple with the L2BlockChainNetworkAdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkAdditionGasPercentOk() (*ModelWeb3WalletGasPercent, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2BlockChainNetworkAdditionGasPercent.Get(), o.L2BlockChainNetworkAdditionGasPercent.IsSet()
}

// HasL2BlockChainNetworkAdditionGasPercent returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2BlockChainNetworkAdditionGasPercent() bool {
	if o != nil && o.L2BlockChainNetworkAdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetL2BlockChainNetworkAdditionGasPercent gets a reference to the given NullableModelWeb3WalletGasPercent and assigns it to the L2BlockChainNetworkAdditionGasPercent field.
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkAdditionGasPercent(v ModelWeb3WalletGasPercent) {
	o.L2BlockChainNetworkAdditionGasPercent.Set(&v)
}
// SetL2BlockChainNetworkAdditionGasPercentNil sets the value for L2BlockChainNetworkAdditionGasPercent to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkAdditionGasPercentNil() {
	o.L2BlockChainNetworkAdditionGasPercent.Set(nil)
}

// UnsetL2BlockChainNetworkAdditionGasPercent ensures that no value is present for L2BlockChainNetworkAdditionGasPercent, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2BlockChainNetworkAdditionGasPercent() {
	o.L2BlockChainNetworkAdditionGasPercent.Unset()
}

// GetL2BlockChainNetworkId returns the L2BlockChainNetworkId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkId() int64 {
	if o == nil || IsNil(o.L2BlockChainNetworkId.Get()) {
		var ret int64
		return ret
	}
	return *o.L2BlockChainNetworkId.Get()
}

// GetL2BlockChainNetworkIdOk returns a tuple with the L2BlockChainNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2BlockChainNetworkId.Get(), o.L2BlockChainNetworkId.IsSet()
}

// HasL2BlockChainNetworkId returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2BlockChainNetworkId() bool {
	if o != nil && o.L2BlockChainNetworkId.IsSet() {
		return true
	}

	return false
}

// SetL2BlockChainNetworkId gets a reference to the given NullableInt64 and assigns it to the L2BlockChainNetworkId field.
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkId(v int64) {
	o.L2BlockChainNetworkId.Set(&v)
}
// SetL2BlockChainNetworkIdNil sets the value for L2BlockChainNetworkId to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkIdNil() {
	o.L2BlockChainNetworkId.Set(nil)
}

// UnsetL2BlockChainNetworkId ensures that no value is present for L2BlockChainNetworkId, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2BlockChainNetworkId() {
	o.L2BlockChainNetworkId.Unset()
}

// GetL2BlockChainNetworkName returns the L2BlockChainNetworkName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkName() string {
	if o == nil || IsNil(o.L2BlockChainNetworkName.Get()) {
		var ret string
		return ret
	}
	return *o.L2BlockChainNetworkName.Get()
}

// GetL2BlockChainNetworkNameOk returns a tuple with the L2BlockChainNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2BlockChainNetworkName.Get(), o.L2BlockChainNetworkName.IsSet()
}

// HasL2BlockChainNetworkName returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2BlockChainNetworkName() bool {
	if o != nil && o.L2BlockChainNetworkName.IsSet() {
		return true
	}

	return false
}

// SetL2BlockChainNetworkName gets a reference to the given NullableString and assigns it to the L2BlockChainNetworkName field.
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkName(v string) {
	o.L2BlockChainNetworkName.Set(&v)
}
// SetL2BlockChainNetworkNameNil sets the value for L2BlockChainNetworkName to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkNameNil() {
	o.L2BlockChainNetworkName.Set(nil)
}

// UnsetL2BlockChainNetworkName ensures that no value is present for L2BlockChainNetworkName, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2BlockChainNetworkName() {
	o.L2BlockChainNetworkName.Unset()
}

// GetL2BlockChainNetworkUrl returns the L2BlockChainNetworkUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkUrl() string {
	if o == nil || IsNil(o.L2BlockChainNetworkUrl.Get()) {
		var ret string
		return ret
	}
	return *o.L2BlockChainNetworkUrl.Get()
}

// GetL2BlockChainNetworkUrlOk returns a tuple with the L2BlockChainNetworkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2BlockChainNetworkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2BlockChainNetworkUrl.Get(), o.L2BlockChainNetworkUrl.IsSet()
}

// HasL2BlockChainNetworkUrl returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2BlockChainNetworkUrl() bool {
	if o != nil && o.L2BlockChainNetworkUrl.IsSet() {
		return true
	}

	return false
}

// SetL2BlockChainNetworkUrl gets a reference to the given NullableString and assigns it to the L2BlockChainNetworkUrl field.
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkUrl(v string) {
	o.L2BlockChainNetworkUrl.Set(&v)
}
// SetL2BlockChainNetworkUrlNil sets the value for L2BlockChainNetworkUrl to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2BlockChainNetworkUrlNil() {
	o.L2BlockChainNetworkUrl.Set(nil)
}

// UnsetL2BlockChainNetworkUrl ensures that no value is present for L2BlockChainNetworkUrl, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2BlockChainNetworkUrl() {
	o.L2BlockChainNetworkUrl.Unset()
}

// GetL2QuoterAddress returns the L2QuoterAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2QuoterAddress() string {
	if o == nil || IsNil(o.L2QuoterAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L2QuoterAddress.Get()
}

// GetL2QuoterAddressOk returns a tuple with the L2QuoterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2QuoterAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2QuoterAddress.Get(), o.L2QuoterAddress.IsSet()
}

// HasL2QuoterAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2QuoterAddress() bool {
	if o != nil && o.L2QuoterAddress.IsSet() {
		return true
	}

	return false
}

// SetL2QuoterAddress gets a reference to the given NullableString and assigns it to the L2QuoterAddress field.
func (o *Web3WalletBlockChainNetwork) SetL2QuoterAddress(v string) {
	o.L2QuoterAddress.Set(&v)
}
// SetL2QuoterAddressNil sets the value for L2QuoterAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2QuoterAddressNil() {
	o.L2QuoterAddress.Set(nil)
}

// UnsetL2QuoterAddress ensures that no value is present for L2QuoterAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2QuoterAddress() {
	o.L2QuoterAddress.Unset()
}

// GetL2WethAddress returns the L2WethAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetL2WethAddress() string {
	if o == nil || IsNil(o.L2WethAddress.Get()) {
		var ret string
		return ret
	}
	return *o.L2WethAddress.Get()
}

// GetL2WethAddressOk returns a tuple with the L2WethAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetL2WethAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2WethAddress.Get(), o.L2WethAddress.IsSet()
}

// HasL2WethAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasL2WethAddress() bool {
	if o != nil && o.L2WethAddress.IsSet() {
		return true
	}

	return false
}

// SetL2WethAddress gets a reference to the given NullableString and assigns it to the L2WethAddress field.
func (o *Web3WalletBlockChainNetwork) SetL2WethAddress(v string) {
	o.L2WethAddress.Set(&v)
}
// SetL2WethAddressNil sets the value for L2WethAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetL2WethAddressNil() {
	o.L2WethAddress.Set(nil)
}

// UnsetL2WethAddress ensures that no value is present for L2WethAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetL2WethAddress() {
	o.L2WethAddress.Unset()
}

// GetMiscPalBaseImageUri returns the MiscPalBaseImageUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetMiscPalBaseImageUri() string {
	if o == nil || IsNil(o.MiscPalBaseImageUri.Get()) {
		var ret string
		return ret
	}
	return *o.MiscPalBaseImageUri.Get()
}

// GetMiscPalBaseImageUriOk returns a tuple with the MiscPalBaseImageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetMiscPalBaseImageUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiscPalBaseImageUri.Get(), o.MiscPalBaseImageUri.IsSet()
}

// HasMiscPalBaseImageUri returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasMiscPalBaseImageUri() bool {
	if o != nil && o.MiscPalBaseImageUri.IsSet() {
		return true
	}

	return false
}

// SetMiscPalBaseImageUri gets a reference to the given NullableString and assigns it to the MiscPalBaseImageUri field.
func (o *Web3WalletBlockChainNetwork) SetMiscPalBaseImageUri(v string) {
	o.MiscPalBaseImageUri.Set(&v)
}
// SetMiscPalBaseImageUriNil sets the value for MiscPalBaseImageUri to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetMiscPalBaseImageUriNil() {
	o.MiscPalBaseImageUri.Set(nil)
}

// UnsetMiscPalBaseImageUri ensures that no value is present for MiscPalBaseImageUri, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetMiscPalBaseImageUri() {
	o.MiscPalBaseImageUri.Unset()
}

// GetMiscPalBaseJsonUri returns the MiscPalBaseJsonUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetMiscPalBaseJsonUri() string {
	if o == nil || IsNil(o.MiscPalBaseJsonUri.Get()) {
		var ret string
		return ret
	}
	return *o.MiscPalBaseJsonUri.Get()
}

// GetMiscPalBaseJsonUriOk returns a tuple with the MiscPalBaseJsonUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetMiscPalBaseJsonUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiscPalBaseJsonUri.Get(), o.MiscPalBaseJsonUri.IsSet()
}

// HasMiscPalBaseJsonUri returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasMiscPalBaseJsonUri() bool {
	if o != nil && o.MiscPalBaseJsonUri.IsSet() {
		return true
	}

	return false
}

// SetMiscPalBaseJsonUri gets a reference to the given NullableString and assigns it to the MiscPalBaseJsonUri field.
func (o *Web3WalletBlockChainNetwork) SetMiscPalBaseJsonUri(v string) {
	o.MiscPalBaseJsonUri.Set(&v)
}
// SetMiscPalBaseJsonUriNil sets the value for MiscPalBaseJsonUri to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetMiscPalBaseJsonUriNil() {
	o.MiscPalBaseJsonUri.Set(nil)
}

// UnsetMiscPalBaseJsonUri ensures that no value is present for MiscPalBaseJsonUri, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetMiscPalBaseJsonUri() {
	o.MiscPalBaseJsonUri.Unset()
}

// GetSmartContractEmplAddress returns the SmartContractEmplAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplAddress() string {
	if o == nil || IsNil(o.SmartContractEmplAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractEmplAddress.Get()
}

// GetSmartContractEmplAddressOk returns a tuple with the SmartContractEmplAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractEmplAddress.Get(), o.SmartContractEmplAddress.IsSet()
}

// HasSmartContractEmplAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractEmplAddress() bool {
	if o != nil && o.SmartContractEmplAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractEmplAddress gets a reference to the given NullableString and assigns it to the SmartContractEmplAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplAddress(v string) {
	o.SmartContractEmplAddress.Set(&v)
}
// SetSmartContractEmplAddressNil sets the value for SmartContractEmplAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplAddressNil() {
	o.SmartContractEmplAddress.Set(nil)
}

// UnsetSmartContractEmplAddress ensures that no value is present for SmartContractEmplAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractEmplAddress() {
	o.SmartContractEmplAddress.Unset()
}

// GetSmartContractEmplDepositAddress returns the SmartContractEmplDepositAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplDepositAddress() string {
	if o == nil || IsNil(o.SmartContractEmplDepositAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractEmplDepositAddress.Get()
}

// GetSmartContractEmplDepositAddressOk returns a tuple with the SmartContractEmplDepositAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplDepositAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractEmplDepositAddress.Get(), o.SmartContractEmplDepositAddress.IsSet()
}

// HasSmartContractEmplDepositAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractEmplDepositAddress() bool {
	if o != nil && o.SmartContractEmplDepositAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractEmplDepositAddress gets a reference to the given NullableString and assigns it to the SmartContractEmplDepositAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplDepositAddress(v string) {
	o.SmartContractEmplDepositAddress.Set(&v)
}
// SetSmartContractEmplDepositAddressNil sets the value for SmartContractEmplDepositAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplDepositAddressNil() {
	o.SmartContractEmplDepositAddress.Set(nil)
}

// UnsetSmartContractEmplDepositAddress ensures that no value is present for SmartContractEmplDepositAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractEmplDepositAddress() {
	o.SmartContractEmplDepositAddress.Unset()
}

// GetSmartContractEmplWithdrawAddress returns the SmartContractEmplWithdrawAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplWithdrawAddress() string {
	if o == nil || IsNil(o.SmartContractEmplWithdrawAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractEmplWithdrawAddress.Get()
}

// GetSmartContractEmplWithdrawAddressOk returns a tuple with the SmartContractEmplWithdrawAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractEmplWithdrawAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractEmplWithdrawAddress.Get(), o.SmartContractEmplWithdrawAddress.IsSet()
}

// HasSmartContractEmplWithdrawAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractEmplWithdrawAddress() bool {
	if o != nil && o.SmartContractEmplWithdrawAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractEmplWithdrawAddress gets a reference to the given NullableString and assigns it to the SmartContractEmplWithdrawAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplWithdrawAddress(v string) {
	o.SmartContractEmplWithdrawAddress.Set(&v)
}
// SetSmartContractEmplWithdrawAddressNil sets the value for SmartContractEmplWithdrawAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractEmplWithdrawAddressNil() {
	o.SmartContractEmplWithdrawAddress.Set(nil)
}

// UnsetSmartContractEmplWithdrawAddress ensures that no value is present for SmartContractEmplWithdrawAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractEmplWithdrawAddress() {
	o.SmartContractEmplWithdrawAddress.Unset()
}

// GetSmartContractL1YayAddress returns the SmartContractL1YayAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractL1YayAddress() string {
	if o == nil || IsNil(o.SmartContractL1YayAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractL1YayAddress.Get()
}

// GetSmartContractL1YayAddressOk returns a tuple with the SmartContractL1YayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractL1YayAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractL1YayAddress.Get(), o.SmartContractL1YayAddress.IsSet()
}

// HasSmartContractL1YayAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractL1YayAddress() bool {
	if o != nil && o.SmartContractL1YayAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractL1YayAddress gets a reference to the given NullableString and assigns it to the SmartContractL1YayAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractL1YayAddress(v string) {
	o.SmartContractL1YayAddress.Set(&v)
}
// SetSmartContractL1YayAddressNil sets the value for SmartContractL1YayAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractL1YayAddressNil() {
	o.SmartContractL1YayAddress.Set(nil)
}

// UnsetSmartContractL1YayAddress ensures that no value is present for SmartContractL1YayAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractL1YayAddress() {
	o.SmartContractL1YayAddress.Unset()
}

// GetSmartContractL2YayAddress returns the SmartContractL2YayAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractL2YayAddress() string {
	if o == nil || IsNil(o.SmartContractL2YayAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractL2YayAddress.Get()
}

// GetSmartContractL2YayAddressOk returns a tuple with the SmartContractL2YayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractL2YayAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractL2YayAddress.Get(), o.SmartContractL2YayAddress.IsSet()
}

// HasSmartContractL2YayAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractL2YayAddress() bool {
	if o != nil && o.SmartContractL2YayAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractL2YayAddress gets a reference to the given NullableString and assigns it to the SmartContractL2YayAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractL2YayAddress(v string) {
	o.SmartContractL2YayAddress.Set(&v)
}
// SetSmartContractL2YayAddressNil sets the value for SmartContractL2YayAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractL2YayAddressNil() {
	o.SmartContractL2YayAddress.Set(nil)
}

// UnsetSmartContractL2YayAddress ensures that no value is present for SmartContractL2YayAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractL2YayAddress() {
	o.SmartContractL2YayAddress.Unset()
}

// GetSmartContractPalAddress returns the SmartContractPalAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractPalAddress() string {
	if o == nil || IsNil(o.SmartContractPalAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractPalAddress.Get()
}

// GetSmartContractPalAddressOk returns a tuple with the SmartContractPalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractPalAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractPalAddress.Get(), o.SmartContractPalAddress.IsSet()
}

// HasSmartContractPalAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractPalAddress() bool {
	if o != nil && o.SmartContractPalAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractPalAddress gets a reference to the given NullableString and assigns it to the SmartContractPalAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractPalAddress(v string) {
	o.SmartContractPalAddress.Set(&v)
}
// SetSmartContractPalAddressNil sets the value for SmartContractPalAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractPalAddressNil() {
	o.SmartContractPalAddress.Set(nil)
}

// UnsetSmartContractPalAddress ensures that no value is present for SmartContractPalAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractPalAddress() {
	o.SmartContractPalAddress.Unset()
}

// GetSmartContractPalDepositWithdrawAddress returns the SmartContractPalDepositWithdrawAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractPalDepositWithdrawAddress() string {
	if o == nil || IsNil(o.SmartContractPalDepositWithdrawAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractPalDepositWithdrawAddress.Get()
}

// GetSmartContractPalDepositWithdrawAddressOk returns a tuple with the SmartContractPalDepositWithdrawAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractPalDepositWithdrawAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractPalDepositWithdrawAddress.Get(), o.SmartContractPalDepositWithdrawAddress.IsSet()
}

// HasSmartContractPalDepositWithdrawAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractPalDepositWithdrawAddress() bool {
	if o != nil && o.SmartContractPalDepositWithdrawAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractPalDepositWithdrawAddress gets a reference to the given NullableString and assigns it to the SmartContractPalDepositWithdrawAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractPalDepositWithdrawAddress(v string) {
	o.SmartContractPalDepositWithdrawAddress.Set(&v)
}
// SetSmartContractPalDepositWithdrawAddressNil sets the value for SmartContractPalDepositWithdrawAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractPalDepositWithdrawAddressNil() {
	o.SmartContractPalDepositWithdrawAddress.Set(nil)
}

// UnsetSmartContractPalDepositWithdrawAddress ensures that no value is present for SmartContractPalDepositWithdrawAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractPalDepositWithdrawAddress() {
	o.SmartContractPalDepositWithdrawAddress.Unset()
}

// GetSmartContractX2y2delegateerc721Address returns the SmartContractX2y2delegateerc721Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractX2y2delegateerc721Address() string {
	if o == nil || IsNil(o.SmartContractX2y2delegateerc721Address.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractX2y2delegateerc721Address.Get()
}

// GetSmartContractX2y2delegateerc721AddressOk returns a tuple with the SmartContractX2y2delegateerc721Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractX2y2delegateerc721AddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractX2y2delegateerc721Address.Get(), o.SmartContractX2y2delegateerc721Address.IsSet()
}

// HasSmartContractX2y2delegateerc721Address returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractX2y2delegateerc721Address() bool {
	if o != nil && o.SmartContractX2y2delegateerc721Address.IsSet() {
		return true
	}

	return false
}

// SetSmartContractX2y2delegateerc721Address gets a reference to the given NullableString and assigns it to the SmartContractX2y2delegateerc721Address field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractX2y2delegateerc721Address(v string) {
	o.SmartContractX2y2delegateerc721Address.Set(&v)
}
// SetSmartContractX2y2delegateerc721AddressNil sets the value for SmartContractX2y2delegateerc721Address to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractX2y2delegateerc721AddressNil() {
	o.SmartContractX2y2delegateerc721Address.Set(nil)
}

// UnsetSmartContractX2y2delegateerc721Address ensures that no value is present for SmartContractX2y2delegateerc721Address, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractX2y2delegateerc721Address() {
	o.SmartContractX2y2delegateerc721Address.Unset()
}

// GetSmartContractX2y2marketAddress returns the SmartContractX2y2marketAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockChainNetwork) GetSmartContractX2y2marketAddress() string {
	if o == nil || IsNil(o.SmartContractX2y2marketAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SmartContractX2y2marketAddress.Get()
}

// GetSmartContractX2y2marketAddressOk returns a tuple with the SmartContractX2y2marketAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockChainNetwork) GetSmartContractX2y2marketAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartContractX2y2marketAddress.Get(), o.SmartContractX2y2marketAddress.IsSet()
}

// HasSmartContractX2y2marketAddress returns a boolean if a field has been set.
func (o *Web3WalletBlockChainNetwork) HasSmartContractX2y2marketAddress() bool {
	if o != nil && o.SmartContractX2y2marketAddress.IsSet() {
		return true
	}

	return false
}

// SetSmartContractX2y2marketAddress gets a reference to the given NullableString and assigns it to the SmartContractX2y2marketAddress field.
func (o *Web3WalletBlockChainNetwork) SetSmartContractX2y2marketAddress(v string) {
	o.SmartContractX2y2marketAddress.Set(&v)
}
// SetSmartContractX2y2marketAddressNil sets the value for SmartContractX2y2marketAddress to be an explicit nil
func (o *Web3WalletBlockChainNetwork) SetSmartContractX2y2marketAddressNil() {
	o.SmartContractX2y2marketAddress.Set(nil)
}

// UnsetSmartContractX2y2marketAddress ensures that no value is present for SmartContractX2y2marketAddress, not even an explicit nil
func (o *Web3WalletBlockChainNetwork) UnsetSmartContractX2y2marketAddress() {
	o.SmartContractX2y2marketAddress.Unset()
}

func (o Web3WalletBlockChainNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletBlockChainNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.L1AmmAddress.IsSet() {
		toSerialize["l1_amm_address"] = o.L1AmmAddress.Get()
	}
	if o.L1BlockChainNetworkAdditionGasPercent.IsSet() {
		toSerialize["l1_block_chain_network_addition_gas_percent"] = o.L1BlockChainNetworkAdditionGasPercent.Get()
	}
	if o.L1BlockChainNetworkId.IsSet() {
		toSerialize["l1_block_chain_network_id"] = o.L1BlockChainNetworkId.Get()
	}
	if o.L1BlockChainNetworkName.IsSet() {
		toSerialize["l1_block_chain_network_name"] = o.L1BlockChainNetworkName.Get()
	}
	if o.L1BlockChainNetworkUrl.IsSet() {
		toSerialize["l1_block_chain_network_url"] = o.L1BlockChainNetworkUrl.Get()
	}
	if o.L1QuoterAddress.IsSet() {
		toSerialize["l1_quoter_address"] = o.L1QuoterAddress.Get()
	}
	if o.L1WethAddress.IsSet() {
		toSerialize["l1_weth_address"] = o.L1WethAddress.Get()
	}
	if o.L2AmmAddress.IsSet() {
		toSerialize["l2_amm_address"] = o.L2AmmAddress.Get()
	}
	if o.L2BlockChainNetworkAdditionGasPercent.IsSet() {
		toSerialize["l2_block_chain_network_addition_gas_percent"] = o.L2BlockChainNetworkAdditionGasPercent.Get()
	}
	if o.L2BlockChainNetworkId.IsSet() {
		toSerialize["l2_block_chain_network_id"] = o.L2BlockChainNetworkId.Get()
	}
	if o.L2BlockChainNetworkName.IsSet() {
		toSerialize["l2_block_chain_network_name"] = o.L2BlockChainNetworkName.Get()
	}
	if o.L2BlockChainNetworkUrl.IsSet() {
		toSerialize["l2_block_chain_network_url"] = o.L2BlockChainNetworkUrl.Get()
	}
	if o.L2QuoterAddress.IsSet() {
		toSerialize["l2_quoter_address"] = o.L2QuoterAddress.Get()
	}
	if o.L2WethAddress.IsSet() {
		toSerialize["l2_weth_address"] = o.L2WethAddress.Get()
	}
	if o.MiscPalBaseImageUri.IsSet() {
		toSerialize["misc_pal_base_image_uri"] = o.MiscPalBaseImageUri.Get()
	}
	if o.MiscPalBaseJsonUri.IsSet() {
		toSerialize["misc_pal_base_json_uri"] = o.MiscPalBaseJsonUri.Get()
	}
	if o.SmartContractEmplAddress.IsSet() {
		toSerialize["smart_contract_empl_address"] = o.SmartContractEmplAddress.Get()
	}
	if o.SmartContractEmplDepositAddress.IsSet() {
		toSerialize["smart_contract_empl_deposit_address"] = o.SmartContractEmplDepositAddress.Get()
	}
	if o.SmartContractEmplWithdrawAddress.IsSet() {
		toSerialize["smart_contract_empl_withdraw_address"] = o.SmartContractEmplWithdrawAddress.Get()
	}
	if o.SmartContractL1YayAddress.IsSet() {
		toSerialize["smart_contract_l1_yay_address"] = o.SmartContractL1YayAddress.Get()
	}
	if o.SmartContractL2YayAddress.IsSet() {
		toSerialize["smart_contract_l2_yay_address"] = o.SmartContractL2YayAddress.Get()
	}
	if o.SmartContractPalAddress.IsSet() {
		toSerialize["smart_contract_pal_address"] = o.SmartContractPalAddress.Get()
	}
	if o.SmartContractPalDepositWithdrawAddress.IsSet() {
		toSerialize["smart_contract_pal_deposit_withdraw_address"] = o.SmartContractPalDepositWithdrawAddress.Get()
	}
	if o.SmartContractX2y2delegateerc721Address.IsSet() {
		toSerialize["smart_contract_x2y2delegateerc721_address"] = o.SmartContractX2y2delegateerc721Address.Get()
	}
	if o.SmartContractX2y2marketAddress.IsSet() {
		toSerialize["smart_contract_x2y2market_address"] = o.SmartContractX2y2marketAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletBlockChainNetwork) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletBlockChainNetwork := _Web3WalletBlockChainNetwork{}

	err = json.Unmarshal(data, &varWeb3WalletBlockChainNetwork)

	if err != nil {
		return err
	}

	*o = Web3WalletBlockChainNetwork(varWeb3WalletBlockChainNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "l1_amm_address")
		delete(additionalProperties, "l1_block_chain_network_addition_gas_percent")
		delete(additionalProperties, "l1_block_chain_network_id")
		delete(additionalProperties, "l1_block_chain_network_name")
		delete(additionalProperties, "l1_block_chain_network_url")
		delete(additionalProperties, "l1_quoter_address")
		delete(additionalProperties, "l1_weth_address")
		delete(additionalProperties, "l2_amm_address")
		delete(additionalProperties, "l2_block_chain_network_addition_gas_percent")
		delete(additionalProperties, "l2_block_chain_network_id")
		delete(additionalProperties, "l2_block_chain_network_name")
		delete(additionalProperties, "l2_block_chain_network_url")
		delete(additionalProperties, "l2_quoter_address")
		delete(additionalProperties, "l2_weth_address")
		delete(additionalProperties, "misc_pal_base_image_uri")
		delete(additionalProperties, "misc_pal_base_json_uri")
		delete(additionalProperties, "smart_contract_empl_address")
		delete(additionalProperties, "smart_contract_empl_deposit_address")
		delete(additionalProperties, "smart_contract_empl_withdraw_address")
		delete(additionalProperties, "smart_contract_l1_yay_address")
		delete(additionalProperties, "smart_contract_l2_yay_address")
		delete(additionalProperties, "smart_contract_pal_address")
		delete(additionalProperties, "smart_contract_pal_deposit_withdraw_address")
		delete(additionalProperties, "smart_contract_x2y2delegateerc721_address")
		delete(additionalProperties, "smart_contract_x2y2market_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletBlockChainNetwork struct {
	value *Web3WalletBlockChainNetwork
	isSet bool
}

func (v NullableWeb3WalletBlockChainNetwork) Get() *Web3WalletBlockChainNetwork {
	return v.value
}

func (v *NullableWeb3WalletBlockChainNetwork) Set(val *Web3WalletBlockChainNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletBlockChainNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletBlockChainNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletBlockChainNetwork(val *Web3WalletBlockChainNetwork) *NullableWeb3WalletBlockChainNetwork {
	return &NullableWeb3WalletBlockChainNetwork{value: val, isSet: true}
}

func (v NullableWeb3WalletBlockChainNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletBlockChainNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


