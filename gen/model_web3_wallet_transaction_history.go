
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletTransactionHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletTransactionHistory{}

// Web3WalletTransactionHistory struct for Web3WalletTransactionHistory
type Web3WalletTransactionHistory struct {
	Amount NullableString `json:"amount,omitempty"`
	AssetInfo NullableAssetInfo `json:"asset_info,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Details NullableEmplTokenExchangeDetails `json:"details,omitempty"`
	FinalStatusUpdatedAt NullableInt64 `json:"final_status_updated_at,omitempty"`
	FromAddress NullableString `json:"from_address,omitempty"`
	GiftTransactionUuid NullableString `json:"gift_transaction_uuid,omitempty"`
	Hash NullableString `json:"hash,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Status NullableString `json:"status,omitempty"`
	ToAddress NullableString `json:"to_address,omitempty"`
	TokenContract NullableString `json:"token_contract,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	TransactionCode NullableString `json:"transaction_code,omitempty"`
	// Unresolved Java type: ao.i
	Type map[string]interface{} `json:"type,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletTransactionHistory Web3WalletTransactionHistory

// NewWeb3WalletTransactionHistory instantiates a new Web3WalletTransactionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletTransactionHistory() *Web3WalletTransactionHistory {
	this := Web3WalletTransactionHistory{}
	return &this
}

// NewWeb3WalletTransactionHistoryWithDefaults instantiates a new Web3WalletTransactionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletTransactionHistoryWithDefaults() *Web3WalletTransactionHistory {
	this := Web3WalletTransactionHistory{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *Web3WalletTransactionHistory) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Web3WalletTransactionHistory) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetAmount() {
	o.Amount.Unset()
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetAssetInfo() AssetInfo {
	if o == nil || IsNil(o.AssetInfo.Get()) {
		var ret AssetInfo
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetAssetInfoOk() (*AssetInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableAssetInfo and assigns it to the AssetInfo field.
func (o *Web3WalletTransactionHistory) SetAssetInfo(v AssetInfo) {
	o.AssetInfo.Set(&v)
}
// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *Web3WalletTransactionHistory) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *Web3WalletTransactionHistory) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Web3WalletTransactionHistory) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Web3WalletTransactionHistory) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Web3WalletTransactionHistory) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Web3WalletTransactionHistory) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Web3WalletTransactionHistory) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetDescription() {
	o.Description.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetDetails() EmplTokenExchangeDetails {
	if o == nil || IsNil(o.Details.Get()) {
		var ret EmplTokenExchangeDetails
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetDetailsOk() (*EmplTokenExchangeDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableEmplTokenExchangeDetails and assigns it to the Details field.
func (o *Web3WalletTransactionHistory) SetDetails(v EmplTokenExchangeDetails) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *Web3WalletTransactionHistory) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetDetails() {
	o.Details.Unset()
}

// GetFinalStatusUpdatedAt returns the FinalStatusUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetFinalStatusUpdatedAt() int64 {
	if o == nil || IsNil(o.FinalStatusUpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.FinalStatusUpdatedAt.Get()
}

// GetFinalStatusUpdatedAtOk returns a tuple with the FinalStatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetFinalStatusUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinalStatusUpdatedAt.Get(), o.FinalStatusUpdatedAt.IsSet()
}

// HasFinalStatusUpdatedAt returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasFinalStatusUpdatedAt() bool {
	if o != nil && o.FinalStatusUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetFinalStatusUpdatedAt gets a reference to the given NullableInt64 and assigns it to the FinalStatusUpdatedAt field.
func (o *Web3WalletTransactionHistory) SetFinalStatusUpdatedAt(v int64) {
	o.FinalStatusUpdatedAt.Set(&v)
}
// SetFinalStatusUpdatedAtNil sets the value for FinalStatusUpdatedAt to be an explicit nil
func (o *Web3WalletTransactionHistory) SetFinalStatusUpdatedAtNil() {
	o.FinalStatusUpdatedAt.Set(nil)
}

// UnsetFinalStatusUpdatedAt ensures that no value is present for FinalStatusUpdatedAt, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetFinalStatusUpdatedAt() {
	o.FinalStatusUpdatedAt.Unset()
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress.Get()) {
		var ret string
		return ret
	}
	return *o.FromAddress.Get()
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromAddress.Get(), o.FromAddress.IsSet()
}

// HasFromAddress returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasFromAddress() bool {
	if o != nil && o.FromAddress.IsSet() {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given NullableString and assigns it to the FromAddress field.
func (o *Web3WalletTransactionHistory) SetFromAddress(v string) {
	o.FromAddress.Set(&v)
}
// SetFromAddressNil sets the value for FromAddress to be an explicit nil
func (o *Web3WalletTransactionHistory) SetFromAddressNil() {
	o.FromAddress.Set(nil)
}

// UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetFromAddress() {
	o.FromAddress.Unset()
}

// GetGiftTransactionUuid returns the GiftTransactionUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetGiftTransactionUuid() string {
	if o == nil || IsNil(o.GiftTransactionUuid.Get()) {
		var ret string
		return ret
	}
	return *o.GiftTransactionUuid.Get()
}

// GetGiftTransactionUuidOk returns a tuple with the GiftTransactionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetGiftTransactionUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftTransactionUuid.Get(), o.GiftTransactionUuid.IsSet()
}

// HasGiftTransactionUuid returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasGiftTransactionUuid() bool {
	if o != nil && o.GiftTransactionUuid.IsSet() {
		return true
	}

	return false
}

// SetGiftTransactionUuid gets a reference to the given NullableString and assigns it to the GiftTransactionUuid field.
func (o *Web3WalletTransactionHistory) SetGiftTransactionUuid(v string) {
	o.GiftTransactionUuid.Set(&v)
}
// SetGiftTransactionUuidNil sets the value for GiftTransactionUuid to be an explicit nil
func (o *Web3WalletTransactionHistory) SetGiftTransactionUuidNil() {
	o.GiftTransactionUuid.Set(nil)
}

// UnsetGiftTransactionUuid ensures that no value is present for GiftTransactionUuid, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetGiftTransactionUuid() {
	o.GiftTransactionUuid.Unset()
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *Web3WalletTransactionHistory) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *Web3WalletTransactionHistory) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetHash() {
	o.Hash.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Web3WalletTransactionHistory) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Web3WalletTransactionHistory) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetId() {
	o.Id.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Web3WalletTransactionHistory) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Web3WalletTransactionHistory) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetStatus() {
	o.Status.Unset()
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetToAddress() string {
	if o == nil || IsNil(o.ToAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ToAddress.Get()
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToAddress.Get(), o.ToAddress.IsSet()
}

// HasToAddress returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasToAddress() bool {
	if o != nil && o.ToAddress.IsSet() {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given NullableString and assigns it to the ToAddress field.
func (o *Web3WalletTransactionHistory) SetToAddress(v string) {
	o.ToAddress.Set(&v)
}
// SetToAddressNil sets the value for ToAddress to be an explicit nil
func (o *Web3WalletTransactionHistory) SetToAddressNil() {
	o.ToAddress.Set(nil)
}

// UnsetToAddress ensures that no value is present for ToAddress, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetToAddress() {
	o.ToAddress.Unset()
}

// GetTokenContract returns the TokenContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetTokenContract() string {
	if o == nil || IsNil(o.TokenContract.Get()) {
		var ret string
		return ret
	}
	return *o.TokenContract.Get()
}

// GetTokenContractOk returns a tuple with the TokenContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetTokenContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenContract.Get(), o.TokenContract.IsSet()
}

// HasTokenContract returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasTokenContract() bool {
	if o != nil && o.TokenContract.IsSet() {
		return true
	}

	return false
}

// SetTokenContract gets a reference to the given NullableString and assigns it to the TokenContract field.
func (o *Web3WalletTransactionHistory) SetTokenContract(v string) {
	o.TokenContract.Set(&v)
}
// SetTokenContractNil sets the value for TokenContract to be an explicit nil
func (o *Web3WalletTransactionHistory) SetTokenContractNil() {
	o.TokenContract.Set(nil)
}

// UnsetTokenContract ensures that no value is present for TokenContract, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetTokenContract() {
	o.TokenContract.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *Web3WalletTransactionHistory) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *Web3WalletTransactionHistory) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasTransactionCode() bool {
	if o != nil && o.TransactionCode.IsSet() {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given NullableString and assigns it to the TransactionCode field.
func (o *Web3WalletTransactionHistory) SetTransactionCode(v string) {
	o.TransactionCode.Set(&v)
}
// SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil
func (o *Web3WalletTransactionHistory) SetTransactionCodeNil() {
	o.TransactionCode.Set(nil)
}

// UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetTransactionCode() {
	o.TransactionCode.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *Web3WalletTransactionHistory) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistory) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistory) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistory) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *Web3WalletTransactionHistory) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Web3WalletTransactionHistory) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Web3WalletTransactionHistory) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Web3WalletTransactionHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletTransactionHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.AssetInfo.IsSet() {
		toSerialize["asset_info"] = o.AssetInfo.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.FinalStatusUpdatedAt.IsSet() {
		toSerialize["final_status_updated_at"] = o.FinalStatusUpdatedAt.Get()
	}
	if o.FromAddress.IsSet() {
		toSerialize["from_address"] = o.FromAddress.Get()
	}
	if o.GiftTransactionUuid.IsSet() {
		toSerialize["gift_transaction_uuid"] = o.GiftTransactionUuid.Get()
	}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.ToAddress.IsSet() {
		toSerialize["to_address"] = o.ToAddress.Get()
	}
	if o.TokenContract.IsSet() {
		toSerialize["token_contract"] = o.TokenContract.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.TransactionCode.IsSet() {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletTransactionHistory) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletTransactionHistory := _Web3WalletTransactionHistory{}

	err = json.Unmarshal(data, &varWeb3WalletTransactionHistory)

	if err != nil {
		return err
	}

	*o = Web3WalletTransactionHistory(varWeb3WalletTransactionHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset_info")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "description")
		delete(additionalProperties, "details")
		delete(additionalProperties, "final_status_updated_at")
		delete(additionalProperties, "from_address")
		delete(additionalProperties, "gift_transaction_uuid")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "to_address")
		delete(additionalProperties, "token_contract")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "transaction_code")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletTransactionHistory struct {
	value *Web3WalletTransactionHistory
	isSet bool
}

func (v NullableWeb3WalletTransactionHistory) Get() *Web3WalletTransactionHistory {
	return v.value
}

func (v *NullableWeb3WalletTransactionHistory) Set(val *Web3WalletTransactionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletTransactionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletTransactionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletTransactionHistory(val *Web3WalletTransactionHistory) *NullableWeb3WalletTransactionHistory {
	return &NullableWeb3WalletTransactionHistory{value: val, isSet: true}
}

func (v NullableWeb3WalletTransactionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletTransactionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


