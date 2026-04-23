
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletLocalHistoryTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletLocalHistoryTransaction{}

// Web3WalletLocalHistoryTransaction struct for Web3WalletLocalHistoryTransaction
type Web3WalletLocalHistoryTransaction struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	AssetTypeName NullableString `json:"asset_type_name,omitempty"`
	Chain NullableString `json:"chain,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ChainUrl NullableString `json:"chain_url,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	GasLimit NullableInt64 `json:"gas_limit,omitempty"`
	GasPercentage NullableInt32 `json:"gas_percentage,omitempty"`
	GasPrice NullableFloat64 `json:"gas_price,omitempty"`
	InputsHex NullableString `json:"inputs_hex,omitempty"`
	Nonce NullableInt64 `json:"nonce,omitempty"`
	PalId NullableString `json:"pal_id,omitempty"`
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
	TransactionHashHex NullableString `json:"transaction_hash_hex,omitempty"`
	TransactionHistoryStatusName NullableString `json:"transaction_history_status_name,omitempty"`
	TransferModeKey NullableString `json:"transfer_mode_key,omitempty"`
	TxUniqueId NullableString `json:"tx_unique_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletLocalHistoryTransaction Web3WalletLocalHistoryTransaction

// NewWeb3WalletLocalHistoryTransaction instantiates a new Web3WalletLocalHistoryTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletLocalHistoryTransaction() *Web3WalletLocalHistoryTransaction {
	this := Web3WalletLocalHistoryTransaction{}
	return &this
}

// NewWeb3WalletLocalHistoryTransactionWithDefaults instantiates a new Web3WalletLocalHistoryTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletLocalHistoryTransactionWithDefaults() *Web3WalletLocalHistoryTransaction {
	this := Web3WalletLocalHistoryTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *Web3WalletLocalHistoryTransaction) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetAmount() {
	o.Amount.Unset()
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetAssetTypeName() string {
	if o == nil || IsNil(o.AssetTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTypeName.Get()
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetAssetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTypeName.Get(), o.AssetTypeName.IsSet()
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName.IsSet() {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given NullableString and assigns it to the AssetTypeName field.
func (o *Web3WalletLocalHistoryTransaction) SetAssetTypeName(v string) {
	o.AssetTypeName.Set(&v)
}
// SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetAssetTypeNameNil() {
	o.AssetTypeName.Set(nil)
}

// UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetAssetTypeName() {
	o.AssetTypeName.Unset()
}

// GetChain returns the Chain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetChain() string {
	if o == nil || IsNil(o.Chain.Get()) {
		var ret string
		return ret
	}
	return *o.Chain.Get()
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chain.Get(), o.Chain.IsSet()
}

// HasChain returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasChain() bool {
	if o != nil && o.Chain.IsSet() {
		return true
	}

	return false
}

// SetChain gets a reference to the given NullableString and assigns it to the Chain field.
func (o *Web3WalletLocalHistoryTransaction) SetChain(v string) {
	o.Chain.Set(&v)
}
// SetChainNil sets the value for Chain to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetChainNil() {
	o.Chain.Set(nil)
}

// UnsetChain ensures that no value is present for Chain, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetChain() {
	o.Chain.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3WalletLocalHistoryTransaction) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetChainId() {
	o.ChainId.Unset()
}

// GetChainUrl returns the ChainUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetChainUrl() string {
	if o == nil || IsNil(o.ChainUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ChainUrl.Get()
}

// GetChainUrlOk returns a tuple with the ChainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetChainUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainUrl.Get(), o.ChainUrl.IsSet()
}

// HasChainUrl returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasChainUrl() bool {
	if o != nil && o.ChainUrl.IsSet() {
		return true
	}

	return false
}

// SetChainUrl gets a reference to the given NullableString and assigns it to the ChainUrl field.
func (o *Web3WalletLocalHistoryTransaction) SetChainUrl(v string) {
	o.ChainUrl.Set(&v)
}
// SetChainUrlNil sets the value for ChainUrl to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetChainUrlNil() {
	o.ChainUrl.Set(nil)
}

// UnsetChainUrl ensures that no value is present for ChainUrl, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetChainUrl() {
	o.ChainUrl.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3WalletLocalHistoryTransaction) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetGasLimit() int64 {
	if o == nil || IsNil(o.GasLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.GasLimit.Get()
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetGasLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasLimit.Get(), o.GasLimit.IsSet()
}

// HasGasLimit returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasGasLimit() bool {
	if o != nil && o.GasLimit.IsSet() {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given NullableInt64 and assigns it to the GasLimit field.
func (o *Web3WalletLocalHistoryTransaction) SetGasLimit(v int64) {
	o.GasLimit.Set(&v)
}
// SetGasLimitNil sets the value for GasLimit to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetGasLimitNil() {
	o.GasLimit.Set(nil)
}

// UnsetGasLimit ensures that no value is present for GasLimit, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetGasLimit() {
	o.GasLimit.Unset()
}

// GetGasPercentage returns the GasPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetGasPercentage() int32 {
	if o == nil || IsNil(o.GasPercentage.Get()) {
		var ret int32
		return ret
	}
	return *o.GasPercentage.Get()
}

// GetGasPercentageOk returns a tuple with the GasPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetGasPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasPercentage.Get(), o.GasPercentage.IsSet()
}

// HasGasPercentage returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasGasPercentage() bool {
	if o != nil && o.GasPercentage.IsSet() {
		return true
	}

	return false
}

// SetGasPercentage gets a reference to the given NullableInt32 and assigns it to the GasPercentage field.
func (o *Web3WalletLocalHistoryTransaction) SetGasPercentage(v int32) {
	o.GasPercentage.Set(&v)
}
// SetGasPercentageNil sets the value for GasPercentage to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetGasPercentageNil() {
	o.GasPercentage.Set(nil)
}

// UnsetGasPercentage ensures that no value is present for GasPercentage, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetGasPercentage() {
	o.GasPercentage.Unset()
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetGasPrice() float64 {
	if o == nil || IsNil(o.GasPrice.Get()) {
		var ret float64
		return ret
	}
	return *o.GasPrice.Get()
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetGasPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasPrice.Get(), o.GasPrice.IsSet()
}

// HasGasPrice returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasGasPrice() bool {
	if o != nil && o.GasPrice.IsSet() {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given NullableFloat64 and assigns it to the GasPrice field.
func (o *Web3WalletLocalHistoryTransaction) SetGasPrice(v float64) {
	o.GasPrice.Set(&v)
}
// SetGasPriceNil sets the value for GasPrice to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetGasPriceNil() {
	o.GasPrice.Set(nil)
}

// UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetGasPrice() {
	o.GasPrice.Unset()
}

// GetInputsHex returns the InputsHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetInputsHex() string {
	if o == nil || IsNil(o.InputsHex.Get()) {
		var ret string
		return ret
	}
	return *o.InputsHex.Get()
}

// GetInputsHexOk returns a tuple with the InputsHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetInputsHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputsHex.Get(), o.InputsHex.IsSet()
}

// HasInputsHex returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasInputsHex() bool {
	if o != nil && o.InputsHex.IsSet() {
		return true
	}

	return false
}

// SetInputsHex gets a reference to the given NullableString and assigns it to the InputsHex field.
func (o *Web3WalletLocalHistoryTransaction) SetInputsHex(v string) {
	o.InputsHex.Set(&v)
}
// SetInputsHexNil sets the value for InputsHex to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetInputsHexNil() {
	o.InputsHex.Set(nil)
}

// UnsetInputsHex ensures that no value is present for InputsHex, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetInputsHex() {
	o.InputsHex.Unset()
}

// GetNonce returns the Nonce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetNonce() int64 {
	if o == nil || IsNil(o.Nonce.Get()) {
		var ret int64
		return ret
	}
	return *o.Nonce.Get()
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetNonceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nonce.Get(), o.Nonce.IsSet()
}

// HasNonce returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasNonce() bool {
	if o != nil && o.Nonce.IsSet() {
		return true
	}

	return false
}

// SetNonce gets a reference to the given NullableInt64 and assigns it to the Nonce field.
func (o *Web3WalletLocalHistoryTransaction) SetNonce(v int64) {
	o.Nonce.Set(&v)
}
// SetNonceNil sets the value for Nonce to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetNonceNil() {
	o.Nonce.Set(nil)
}

// UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetNonce() {
	o.Nonce.Unset()
}

// GetPalId returns the PalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetPalId() string {
	if o == nil || IsNil(o.PalId.Get()) {
		var ret string
		return ret
	}
	return *o.PalId.Get()
}

// GetPalIdOk returns a tuple with the PalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetPalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalId.Get(), o.PalId.IsSet()
}

// HasPalId returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasPalId() bool {
	if o != nil && o.PalId.IsSet() {
		return true
	}

	return false
}

// SetPalId gets a reference to the given NullableString and assigns it to the PalId field.
func (o *Web3WalletLocalHistoryTransaction) SetPalId(v string) {
	o.PalId.Set(&v)
}
// SetPalIdNil sets the value for PalId to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetPalIdNil() {
	o.PalId.Set(nil)
}

// UnsetPalId ensures that no value is present for PalId, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetPalId() {
	o.PalId.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *Web3WalletLocalHistoryTransaction) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetTransactionHashHex returns the TransactionHashHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetTransactionHashHex() string {
	if o == nil || IsNil(o.TransactionHashHex.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionHashHex.Get()
}

// GetTransactionHashHexOk returns a tuple with the TransactionHashHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetTransactionHashHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionHashHex.Get(), o.TransactionHashHex.IsSet()
}

// HasTransactionHashHex returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasTransactionHashHex() bool {
	if o != nil && o.TransactionHashHex.IsSet() {
		return true
	}

	return false
}

// SetTransactionHashHex gets a reference to the given NullableString and assigns it to the TransactionHashHex field.
func (o *Web3WalletLocalHistoryTransaction) SetTransactionHashHex(v string) {
	o.TransactionHashHex.Set(&v)
}
// SetTransactionHashHexNil sets the value for TransactionHashHex to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetTransactionHashHexNil() {
	o.TransactionHashHex.Set(nil)
}

// UnsetTransactionHashHex ensures that no value is present for TransactionHashHex, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetTransactionHashHex() {
	o.TransactionHashHex.Unset()
}

// GetTransactionHistoryStatusName returns the TransactionHistoryStatusName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetTransactionHistoryStatusName() string {
	if o == nil || IsNil(o.TransactionHistoryStatusName.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionHistoryStatusName.Get()
}

// GetTransactionHistoryStatusNameOk returns a tuple with the TransactionHistoryStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetTransactionHistoryStatusNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionHistoryStatusName.Get(), o.TransactionHistoryStatusName.IsSet()
}

// HasTransactionHistoryStatusName returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasTransactionHistoryStatusName() bool {
	if o != nil && o.TransactionHistoryStatusName.IsSet() {
		return true
	}

	return false
}

// SetTransactionHistoryStatusName gets a reference to the given NullableString and assigns it to the TransactionHistoryStatusName field.
func (o *Web3WalletLocalHistoryTransaction) SetTransactionHistoryStatusName(v string) {
	o.TransactionHistoryStatusName.Set(&v)
}
// SetTransactionHistoryStatusNameNil sets the value for TransactionHistoryStatusName to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetTransactionHistoryStatusNameNil() {
	o.TransactionHistoryStatusName.Set(nil)
}

// UnsetTransactionHistoryStatusName ensures that no value is present for TransactionHistoryStatusName, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetTransactionHistoryStatusName() {
	o.TransactionHistoryStatusName.Unset()
}

// GetTransferModeKey returns the TransferModeKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetTransferModeKey() string {
	if o == nil || IsNil(o.TransferModeKey.Get()) {
		var ret string
		return ret
	}
	return *o.TransferModeKey.Get()
}

// GetTransferModeKeyOk returns a tuple with the TransferModeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetTransferModeKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransferModeKey.Get(), o.TransferModeKey.IsSet()
}

// HasTransferModeKey returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasTransferModeKey() bool {
	if o != nil && o.TransferModeKey.IsSet() {
		return true
	}

	return false
}

// SetTransferModeKey gets a reference to the given NullableString and assigns it to the TransferModeKey field.
func (o *Web3WalletLocalHistoryTransaction) SetTransferModeKey(v string) {
	o.TransferModeKey.Set(&v)
}
// SetTransferModeKeyNil sets the value for TransferModeKey to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetTransferModeKeyNil() {
	o.TransferModeKey.Set(nil)
}

// UnsetTransferModeKey ensures that no value is present for TransferModeKey, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetTransferModeKey() {
	o.TransferModeKey.Unset()
}

// GetTxUniqueId returns the TxUniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletLocalHistoryTransaction) GetTxUniqueId() string {
	if o == nil || IsNil(o.TxUniqueId.Get()) {
		var ret string
		return ret
	}
	return *o.TxUniqueId.Get()
}

// GetTxUniqueIdOk returns a tuple with the TxUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletLocalHistoryTransaction) GetTxUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxUniqueId.Get(), o.TxUniqueId.IsSet()
}

// HasTxUniqueId returns a boolean if a field has been set.
func (o *Web3WalletLocalHistoryTransaction) HasTxUniqueId() bool {
	if o != nil && o.TxUniqueId.IsSet() {
		return true
	}

	return false
}

// SetTxUniqueId gets a reference to the given NullableString and assigns it to the TxUniqueId field.
func (o *Web3WalletLocalHistoryTransaction) SetTxUniqueId(v string) {
	o.TxUniqueId.Set(&v)
}
// SetTxUniqueIdNil sets the value for TxUniqueId to be an explicit nil
func (o *Web3WalletLocalHistoryTransaction) SetTxUniqueIdNil() {
	o.TxUniqueId.Set(nil)
}

// UnsetTxUniqueId ensures that no value is present for TxUniqueId, not even an explicit nil
func (o *Web3WalletLocalHistoryTransaction) UnsetTxUniqueId() {
	o.TxUniqueId.Unset()
}

func (o Web3WalletLocalHistoryTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletLocalHistoryTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.AssetTypeName.IsSet() {
		toSerialize["asset_type_name"] = o.AssetTypeName.Get()
	}
	if o.Chain.IsSet() {
		toSerialize["chain"] = o.Chain.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ChainUrl.IsSet() {
		toSerialize["chain_url"] = o.ChainUrl.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.GasLimit.IsSet() {
		toSerialize["gas_limit"] = o.GasLimit.Get()
	}
	if o.GasPercentage.IsSet() {
		toSerialize["gas_percentage"] = o.GasPercentage.Get()
	}
	if o.GasPrice.IsSet() {
		toSerialize["gas_price"] = o.GasPrice.Get()
	}
	if o.InputsHex.IsSet() {
		toSerialize["inputs_hex"] = o.InputsHex.Get()
	}
	if o.Nonce.IsSet() {
		toSerialize["nonce"] = o.Nonce.Get()
	}
	if o.PalId.IsSet() {
		toSerialize["pal_id"] = o.PalId.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.TransactionHashHex.IsSet() {
		toSerialize["transaction_hash_hex"] = o.TransactionHashHex.Get()
	}
	if o.TransactionHistoryStatusName.IsSet() {
		toSerialize["transaction_history_status_name"] = o.TransactionHistoryStatusName.Get()
	}
	if o.TransferModeKey.IsSet() {
		toSerialize["transfer_mode_key"] = o.TransferModeKey.Get()
	}
	if o.TxUniqueId.IsSet() {
		toSerialize["tx_unique_id"] = o.TxUniqueId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletLocalHistoryTransaction) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletLocalHistoryTransaction := _Web3WalletLocalHistoryTransaction{}

	err = json.Unmarshal(data, &varWeb3WalletLocalHistoryTransaction)

	if err != nil {
		return err
	}

	*o = Web3WalletLocalHistoryTransaction(varWeb3WalletLocalHistoryTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset_type_name")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_url")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "gas_limit")
		delete(additionalProperties, "gas_percentage")
		delete(additionalProperties, "gas_price")
		delete(additionalProperties, "inputs_hex")
		delete(additionalProperties, "nonce")
		delete(additionalProperties, "pal_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "transaction_hash_hex")
		delete(additionalProperties, "transaction_history_status_name")
		delete(additionalProperties, "transfer_mode_key")
		delete(additionalProperties, "tx_unique_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletLocalHistoryTransaction struct {
	value *Web3WalletLocalHistoryTransaction
	isSet bool
}

func (v NullableWeb3WalletLocalHistoryTransaction) Get() *Web3WalletLocalHistoryTransaction {
	return v.value
}

func (v *NullableWeb3WalletLocalHistoryTransaction) Set(val *Web3WalletLocalHistoryTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletLocalHistoryTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletLocalHistoryTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletLocalHistoryTransaction(val *Web3WalletLocalHistoryTransaction) *NullableWeb3WalletLocalHistoryTransaction {
	return &NullableWeb3WalletLocalHistoryTransaction{value: val, isSet: true}
}

func (v NullableWeb3WalletLocalHistoryTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletLocalHistoryTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


