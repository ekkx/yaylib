
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelWeb3WalletTransactionHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelWeb3WalletTransactionHistory{}

// ModelWeb3WalletTransactionHistory struct for ModelWeb3WalletTransactionHistory
type ModelWeb3WalletTransactionHistory struct {
	AssetInfo NullableWeb3WalletTransactionHistoryAssetInfo `json:"asset_info,omitempty"`
	AssetType NullableAssetType `json:"asset_type,omitempty"`
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	DotMoneyAmount NullableFloat64 `json:"dot_money_amount,omitempty"`
	EmplDetails NullableEmplDetails `json:"empl_details,omitempty"`
	FromAddress NullableString `json:"from_address,omitempty"`
	GiftExchangeUuid NullableString `json:"gift_exchange_uuid,omitempty"`
	GiftTransactionDetail NullableGiftTransactionDetail `json:"gift_transaction_detail,omitempty"`
	GiftTransactionUuid NullableString `json:"gift_transaction_uuid,omitempty"`
	IsPendingTransaction NullableBool `json:"is_pending_transaction,omitempty"`
	NftTokenId NullableString `json:"nft_token_id,omitempty"`
	PalTransactionDetail map[string]interface{} `json:"pal_transaction_detail,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	RaceId NullableString `json:"race_id,omitempty"`
	RaceRanks []int32 `json:"race_ranks,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
	ToAddress NullableString `json:"to_address,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	TransactionCode NullableString `json:"transaction_code,omitempty"`
	TransactionHashHex NullableString `json:"transaction_hash_hex,omitempty"`
	TransactionHistoryType map[string]interface{} `json:"transaction_history_type,omitempty"`
	TransactionType map[string]interface{} `json:"transaction_type,omitempty"`
	Value NullableFloat64 `json:"value,omitempty"`
	ValueSymbol NullableString `json:"value_symbol,omitempty"`
	ValueText NullableString `json:"value_text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelWeb3WalletTransactionHistory ModelWeb3WalletTransactionHistory

// NewModelWeb3WalletTransactionHistory instantiates a new ModelWeb3WalletTransactionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelWeb3WalletTransactionHistory() *ModelWeb3WalletTransactionHistory {
	this := ModelWeb3WalletTransactionHistory{}
	return &this
}

// NewModelWeb3WalletTransactionHistoryWithDefaults instantiates a new ModelWeb3WalletTransactionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWeb3WalletTransactionHistoryWithDefaults() *ModelWeb3WalletTransactionHistory {
	this := ModelWeb3WalletTransactionHistory{}
	return &this
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetAssetInfo() Web3WalletTransactionHistoryAssetInfo {
	if o == nil || IsNil(o.AssetInfo.Get()) {
		var ret Web3WalletTransactionHistoryAssetInfo
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetAssetInfoOk() (*Web3WalletTransactionHistoryAssetInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableWeb3WalletTransactionHistoryAssetInfo and assigns it to the AssetInfo field.
func (o *ModelWeb3WalletTransactionHistory) SetAssetInfo(v Web3WalletTransactionHistoryAssetInfo) {
	o.AssetInfo.Set(&v)
}
// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

// GetAssetType returns the AssetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetAssetType() AssetType {
	if o == nil || IsNil(o.AssetType.Get()) {
		var ret AssetType
		return ret
	}
	return *o.AssetType.Get()
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetType.Get(), o.AssetType.IsSet()
}

// HasAssetType returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasAssetType() bool {
	if o != nil && o.AssetType.IsSet() {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given NullableAssetType and assigns it to the AssetType field.
func (o *ModelWeb3WalletTransactionHistory) SetAssetType(v AssetType) {
	o.AssetType.Set(&v)
}
// SetAssetTypeNil sets the value for AssetType to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetAssetTypeNil() {
	o.AssetType.Set(nil)
}

// UnsetAssetType ensures that no value is present for AssetType, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetAssetType() {
	o.AssetType.Unset()
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *ModelWeb3WalletTransactionHistory) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *ModelWeb3WalletTransactionHistory) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDotMoneyAmount returns the DotMoneyAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetDotMoneyAmount() float64 {
	if o == nil || IsNil(o.DotMoneyAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.DotMoneyAmount.Get()
}

// GetDotMoneyAmountOk returns a tuple with the DotMoneyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetDotMoneyAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DotMoneyAmount.Get(), o.DotMoneyAmount.IsSet()
}

// HasDotMoneyAmount returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasDotMoneyAmount() bool {
	if o != nil && o.DotMoneyAmount.IsSet() {
		return true
	}

	return false
}

// SetDotMoneyAmount gets a reference to the given NullableFloat64 and assigns it to the DotMoneyAmount field.
func (o *ModelWeb3WalletTransactionHistory) SetDotMoneyAmount(v float64) {
	o.DotMoneyAmount.Set(&v)
}
// SetDotMoneyAmountNil sets the value for DotMoneyAmount to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetDotMoneyAmountNil() {
	o.DotMoneyAmount.Set(nil)
}

// UnsetDotMoneyAmount ensures that no value is present for DotMoneyAmount, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetDotMoneyAmount() {
	o.DotMoneyAmount.Unset()
}

// GetEmplDetails returns the EmplDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetEmplDetails() EmplDetails {
	if o == nil || IsNil(o.EmplDetails.Get()) {
		var ret EmplDetails
		return ret
	}
	return *o.EmplDetails.Get()
}

// GetEmplDetailsOk returns a tuple with the EmplDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetEmplDetailsOk() (*EmplDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplDetails.Get(), o.EmplDetails.IsSet()
}

// HasEmplDetails returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasEmplDetails() bool {
	if o != nil && o.EmplDetails.IsSet() {
		return true
	}

	return false
}

// SetEmplDetails gets a reference to the given NullableEmplDetails and assigns it to the EmplDetails field.
func (o *ModelWeb3WalletTransactionHistory) SetEmplDetails(v EmplDetails) {
	o.EmplDetails.Set(&v)
}
// SetEmplDetailsNil sets the value for EmplDetails to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetEmplDetailsNil() {
	o.EmplDetails.Set(nil)
}

// UnsetEmplDetails ensures that no value is present for EmplDetails, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetEmplDetails() {
	o.EmplDetails.Unset()
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress.Get()) {
		var ret string
		return ret
	}
	return *o.FromAddress.Get()
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromAddress.Get(), o.FromAddress.IsSet()
}

// HasFromAddress returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasFromAddress() bool {
	if o != nil && o.FromAddress.IsSet() {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given NullableString and assigns it to the FromAddress field.
func (o *ModelWeb3WalletTransactionHistory) SetFromAddress(v string) {
	o.FromAddress.Set(&v)
}
// SetFromAddressNil sets the value for FromAddress to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetFromAddressNil() {
	o.FromAddress.Set(nil)
}

// UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetFromAddress() {
	o.FromAddress.Unset()
}

// GetGiftExchangeUuid returns the GiftExchangeUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetGiftExchangeUuid() string {
	if o == nil || IsNil(o.GiftExchangeUuid.Get()) {
		var ret string
		return ret
	}
	return *o.GiftExchangeUuid.Get()
}

// GetGiftExchangeUuidOk returns a tuple with the GiftExchangeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetGiftExchangeUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftExchangeUuid.Get(), o.GiftExchangeUuid.IsSet()
}

// HasGiftExchangeUuid returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasGiftExchangeUuid() bool {
	if o != nil && o.GiftExchangeUuid.IsSet() {
		return true
	}

	return false
}

// SetGiftExchangeUuid gets a reference to the given NullableString and assigns it to the GiftExchangeUuid field.
func (o *ModelWeb3WalletTransactionHistory) SetGiftExchangeUuid(v string) {
	o.GiftExchangeUuid.Set(&v)
}
// SetGiftExchangeUuidNil sets the value for GiftExchangeUuid to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetGiftExchangeUuidNil() {
	o.GiftExchangeUuid.Set(nil)
}

// UnsetGiftExchangeUuid ensures that no value is present for GiftExchangeUuid, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetGiftExchangeUuid() {
	o.GiftExchangeUuid.Unset()
}

// GetGiftTransactionDetail returns the GiftTransactionDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetGiftTransactionDetail() GiftTransactionDetail {
	if o == nil || IsNil(o.GiftTransactionDetail.Get()) {
		var ret GiftTransactionDetail
		return ret
	}
	return *o.GiftTransactionDetail.Get()
}

// GetGiftTransactionDetailOk returns a tuple with the GiftTransactionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetGiftTransactionDetailOk() (*GiftTransactionDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftTransactionDetail.Get(), o.GiftTransactionDetail.IsSet()
}

// HasGiftTransactionDetail returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasGiftTransactionDetail() bool {
	if o != nil && o.GiftTransactionDetail.IsSet() {
		return true
	}

	return false
}

// SetGiftTransactionDetail gets a reference to the given NullableGiftTransactionDetail and assigns it to the GiftTransactionDetail field.
func (o *ModelWeb3WalletTransactionHistory) SetGiftTransactionDetail(v GiftTransactionDetail) {
	o.GiftTransactionDetail.Set(&v)
}
// SetGiftTransactionDetailNil sets the value for GiftTransactionDetail to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetGiftTransactionDetailNil() {
	o.GiftTransactionDetail.Set(nil)
}

// UnsetGiftTransactionDetail ensures that no value is present for GiftTransactionDetail, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetGiftTransactionDetail() {
	o.GiftTransactionDetail.Unset()
}

// GetGiftTransactionUuid returns the GiftTransactionUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetGiftTransactionUuid() string {
	if o == nil || IsNil(o.GiftTransactionUuid.Get()) {
		var ret string
		return ret
	}
	return *o.GiftTransactionUuid.Get()
}

// GetGiftTransactionUuidOk returns a tuple with the GiftTransactionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetGiftTransactionUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftTransactionUuid.Get(), o.GiftTransactionUuid.IsSet()
}

// HasGiftTransactionUuid returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasGiftTransactionUuid() bool {
	if o != nil && o.GiftTransactionUuid.IsSet() {
		return true
	}

	return false
}

// SetGiftTransactionUuid gets a reference to the given NullableString and assigns it to the GiftTransactionUuid field.
func (o *ModelWeb3WalletTransactionHistory) SetGiftTransactionUuid(v string) {
	o.GiftTransactionUuid.Set(&v)
}
// SetGiftTransactionUuidNil sets the value for GiftTransactionUuid to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetGiftTransactionUuidNil() {
	o.GiftTransactionUuid.Set(nil)
}

// UnsetGiftTransactionUuid ensures that no value is present for GiftTransactionUuid, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetGiftTransactionUuid() {
	o.GiftTransactionUuid.Unset()
}

// GetIsPendingTransaction returns the IsPendingTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetIsPendingTransaction() bool {
	if o == nil || IsNil(o.IsPendingTransaction.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPendingTransaction.Get()
}

// GetIsPendingTransactionOk returns a tuple with the IsPendingTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetIsPendingTransactionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPendingTransaction.Get(), o.IsPendingTransaction.IsSet()
}

// HasIsPendingTransaction returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasIsPendingTransaction() bool {
	if o != nil && o.IsPendingTransaction.IsSet() {
		return true
	}

	return false
}

// SetIsPendingTransaction gets a reference to the given NullableBool and assigns it to the IsPendingTransaction field.
func (o *ModelWeb3WalletTransactionHistory) SetIsPendingTransaction(v bool) {
	o.IsPendingTransaction.Set(&v)
}
// SetIsPendingTransactionNil sets the value for IsPendingTransaction to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetIsPendingTransactionNil() {
	o.IsPendingTransaction.Set(nil)
}

// UnsetIsPendingTransaction ensures that no value is present for IsPendingTransaction, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetIsPendingTransaction() {
	o.IsPendingTransaction.Unset()
}

// GetNftTokenId returns the NftTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetNftTokenId() string {
	if o == nil || IsNil(o.NftTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.NftTokenId.Get()
}

// GetNftTokenIdOk returns a tuple with the NftTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetNftTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftTokenId.Get(), o.NftTokenId.IsSet()
}

// HasNftTokenId returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasNftTokenId() bool {
	if o != nil && o.NftTokenId.IsSet() {
		return true
	}

	return false
}

// SetNftTokenId gets a reference to the given NullableString and assigns it to the NftTokenId field.
func (o *ModelWeb3WalletTransactionHistory) SetNftTokenId(v string) {
	o.NftTokenId.Set(&v)
}
// SetNftTokenIdNil sets the value for NftTokenId to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetNftTokenIdNil() {
	o.NftTokenId.Set(nil)
}

// UnsetNftTokenId ensures that no value is present for NftTokenId, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetNftTokenId() {
	o.NftTokenId.Unset()
}

// GetPalTransactionDetail returns the PalTransactionDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetPalTransactionDetail() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PalTransactionDetail
}

// GetPalTransactionDetailOk returns a tuple with the PalTransactionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetPalTransactionDetailOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PalTransactionDetail) {
		return map[string]interface{}{}, false
	}
	return o.PalTransactionDetail, true
}

// HasPalTransactionDetail returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasPalTransactionDetail() bool {
	if o != nil && !IsNil(o.PalTransactionDetail) {
		return true
	}

	return false
}

// SetPalTransactionDetail gets a reference to the given map[string]interface{} and assigns it to the PalTransactionDetail field.
func (o *ModelWeb3WalletTransactionHistory) SetPalTransactionDetail(v map[string]interface{}) {
	o.PalTransactionDetail = v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *ModelWeb3WalletTransactionHistory) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetPrice() {
	o.Price.Unset()
}

// GetRaceId returns the RaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetRaceId() string {
	if o == nil || IsNil(o.RaceId.Get()) {
		var ret string
		return ret
	}
	return *o.RaceId.Get()
}

// GetRaceIdOk returns a tuple with the RaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetRaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceId.Get(), o.RaceId.IsSet()
}

// HasRaceId returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasRaceId() bool {
	if o != nil && o.RaceId.IsSet() {
		return true
	}

	return false
}

// SetRaceId gets a reference to the given NullableString and assigns it to the RaceId field.
func (o *ModelWeb3WalletTransactionHistory) SetRaceId(v string) {
	o.RaceId.Set(&v)
}
// SetRaceIdNil sets the value for RaceId to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetRaceIdNil() {
	o.RaceId.Set(nil)
}

// UnsetRaceId ensures that no value is present for RaceId, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetRaceId() {
	o.RaceId.Unset()
}

// GetRaceRanks returns the RaceRanks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetRaceRanks() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.RaceRanks
}

// GetRaceRanksOk returns a tuple with the RaceRanks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetRaceRanksOk() ([]int32, bool) {
	if o == nil || IsNil(o.RaceRanks) {
		return nil, false
	}
	return o.RaceRanks, true
}

// HasRaceRanks returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasRaceRanks() bool {
	if o != nil && !IsNil(o.RaceRanks) {
		return true
	}

	return false
}

// SetRaceRanks gets a reference to the given []int32 and assigns it to the RaceRanks field.
func (o *ModelWeb3WalletTransactionHistory) SetRaceRanks(v []int32) {
	o.RaceRanks = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *ModelWeb3WalletTransactionHistory) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetToAddress() string {
	if o == nil || IsNil(o.ToAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ToAddress.Get()
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToAddress.Get(), o.ToAddress.IsSet()
}

// HasToAddress returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasToAddress() bool {
	if o != nil && o.ToAddress.IsSet() {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given NullableString and assigns it to the ToAddress field.
func (o *ModelWeb3WalletTransactionHistory) SetToAddress(v string) {
	o.ToAddress.Set(&v)
}
// SetToAddressNil sets the value for ToAddress to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetToAddressNil() {
	o.ToAddress.Set(nil)
}

// UnsetToAddress ensures that no value is present for ToAddress, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetToAddress() {
	o.ToAddress.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *ModelWeb3WalletTransactionHistory) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasTransactionCode() bool {
	if o != nil && o.TransactionCode.IsSet() {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given NullableString and assigns it to the TransactionCode field.
func (o *ModelWeb3WalletTransactionHistory) SetTransactionCode(v string) {
	o.TransactionCode.Set(&v)
}
// SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetTransactionCodeNil() {
	o.TransactionCode.Set(nil)
}

// UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetTransactionCode() {
	o.TransactionCode.Unset()
}

// GetTransactionHashHex returns the TransactionHashHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetTransactionHashHex() string {
	if o == nil || IsNil(o.TransactionHashHex.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionHashHex.Get()
}

// GetTransactionHashHexOk returns a tuple with the TransactionHashHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetTransactionHashHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionHashHex.Get(), o.TransactionHashHex.IsSet()
}

// HasTransactionHashHex returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasTransactionHashHex() bool {
	if o != nil && o.TransactionHashHex.IsSet() {
		return true
	}

	return false
}

// SetTransactionHashHex gets a reference to the given NullableString and assigns it to the TransactionHashHex field.
func (o *ModelWeb3WalletTransactionHistory) SetTransactionHashHex(v string) {
	o.TransactionHashHex.Set(&v)
}
// SetTransactionHashHexNil sets the value for TransactionHashHex to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetTransactionHashHexNil() {
	o.TransactionHashHex.Set(nil)
}

// UnsetTransactionHashHex ensures that no value is present for TransactionHashHex, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetTransactionHashHex() {
	o.TransactionHashHex.Unset()
}

// GetTransactionHistoryType returns the TransactionHistoryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetTransactionHistoryType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransactionHistoryType
}

// GetTransactionHistoryTypeOk returns a tuple with the TransactionHistoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetTransactionHistoryTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransactionHistoryType) {
		return map[string]interface{}{}, false
	}
	return o.TransactionHistoryType, true
}

// HasTransactionHistoryType returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasTransactionHistoryType() bool {
	if o != nil && !IsNil(o.TransactionHistoryType) {
		return true
	}

	return false
}

// SetTransactionHistoryType gets a reference to the given map[string]interface{} and assigns it to the TransactionHistoryType field.
func (o *ModelWeb3WalletTransactionHistory) SetTransactionHistoryType(v map[string]interface{}) {
	o.TransactionHistoryType = v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetTransactionType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetTransactionTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return map[string]interface{}{}, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given map[string]interface{} and assigns it to the TransactionType field.
func (o *ModelWeb3WalletTransactionHistory) SetTransactionType(v map[string]interface{}) {
	o.TransactionType = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetValue() float64 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat64 and assigns it to the Value field.
func (o *ModelWeb3WalletTransactionHistory) SetValue(v float64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetValue() {
	o.Value.Unset()
}

// GetValueSymbol returns the ValueSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetValueSymbol() string {
	if o == nil || IsNil(o.ValueSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.ValueSymbol.Get()
}

// GetValueSymbolOk returns a tuple with the ValueSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetValueSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueSymbol.Get(), o.ValueSymbol.IsSet()
}

// HasValueSymbol returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasValueSymbol() bool {
	if o != nil && o.ValueSymbol.IsSet() {
		return true
	}

	return false
}

// SetValueSymbol gets a reference to the given NullableString and assigns it to the ValueSymbol field.
func (o *ModelWeb3WalletTransactionHistory) SetValueSymbol(v string) {
	o.ValueSymbol.Set(&v)
}
// SetValueSymbolNil sets the value for ValueSymbol to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetValueSymbolNil() {
	o.ValueSymbol.Set(nil)
}

// UnsetValueSymbol ensures that no value is present for ValueSymbol, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetValueSymbol() {
	o.ValueSymbol.Unset()
}

// GetValueText returns the ValueText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletTransactionHistory) GetValueText() string {
	if o == nil || IsNil(o.ValueText.Get()) {
		var ret string
		return ret
	}
	return *o.ValueText.Get()
}

// GetValueTextOk returns a tuple with the ValueText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletTransactionHistory) GetValueTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueText.Get(), o.ValueText.IsSet()
}

// HasValueText returns a boolean if a field has been set.
func (o *ModelWeb3WalletTransactionHistory) HasValueText() bool {
	if o != nil && o.ValueText.IsSet() {
		return true
	}

	return false
}

// SetValueText gets a reference to the given NullableString and assigns it to the ValueText field.
func (o *ModelWeb3WalletTransactionHistory) SetValueText(v string) {
	o.ValueText.Set(&v)
}
// SetValueTextNil sets the value for ValueText to be an explicit nil
func (o *ModelWeb3WalletTransactionHistory) SetValueTextNil() {
	o.ValueText.Set(nil)
}

// UnsetValueText ensures that no value is present for ValueText, not even an explicit nil
func (o *ModelWeb3WalletTransactionHistory) UnsetValueText() {
	o.ValueText.Unset()
}

func (o ModelWeb3WalletTransactionHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelWeb3WalletTransactionHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetInfo.IsSet() {
		toSerialize["asset_info"] = o.AssetInfo.Get()
	}
	if o.AssetType.IsSet() {
		toSerialize["asset_type"] = o.AssetType.Get()
	}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.DotMoneyAmount.IsSet() {
		toSerialize["dot_money_amount"] = o.DotMoneyAmount.Get()
	}
	if o.EmplDetails.IsSet() {
		toSerialize["empl_details"] = o.EmplDetails.Get()
	}
	if o.FromAddress.IsSet() {
		toSerialize["from_address"] = o.FromAddress.Get()
	}
	if o.GiftExchangeUuid.IsSet() {
		toSerialize["gift_exchange_uuid"] = o.GiftExchangeUuid.Get()
	}
	if o.GiftTransactionDetail.IsSet() {
		toSerialize["gift_transaction_detail"] = o.GiftTransactionDetail.Get()
	}
	if o.GiftTransactionUuid.IsSet() {
		toSerialize["gift_transaction_uuid"] = o.GiftTransactionUuid.Get()
	}
	if o.IsPendingTransaction.IsSet() {
		toSerialize["is_pending_transaction"] = o.IsPendingTransaction.Get()
	}
	if o.NftTokenId.IsSet() {
		toSerialize["nft_token_id"] = o.NftTokenId.Get()
	}
	if o.PalTransactionDetail != nil {
		toSerialize["pal_transaction_detail"] = o.PalTransactionDetail
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.RaceId.IsSet() {
		toSerialize["race_id"] = o.RaceId.Get()
	}
	if o.RaceRanks != nil {
		toSerialize["race_ranks"] = o.RaceRanks
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ToAddress.IsSet() {
		toSerialize["to_address"] = o.ToAddress.Get()
	}
	if o.TokenSymbol.IsSet() {
		toSerialize["token_symbol"] = o.TokenSymbol.Get()
	}
	if o.TransactionCode.IsSet() {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.TransactionHashHex.IsSet() {
		toSerialize["transaction_hash_hex"] = o.TransactionHashHex.Get()
	}
	if o.TransactionHistoryType != nil {
		toSerialize["transaction_history_type"] = o.TransactionHistoryType
	}
	if o.TransactionType != nil {
		toSerialize["transaction_type"] = o.TransactionType
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.ValueSymbol.IsSet() {
		toSerialize["value_symbol"] = o.ValueSymbol.Get()
	}
	if o.ValueText.IsSet() {
		toSerialize["value_text"] = o.ValueText.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelWeb3WalletTransactionHistory) UnmarshalJSON(data []byte) (err error) {
	varModelWeb3WalletTransactionHistory := _ModelWeb3WalletTransactionHistory{}

	err = json.Unmarshal(data, &varModelWeb3WalletTransactionHistory)

	if err != nil {
		return err
	}

	*o = ModelWeb3WalletTransactionHistory(varModelWeb3WalletTransactionHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_info")
		delete(additionalProperties, "asset_type")
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "dot_money_amount")
		delete(additionalProperties, "empl_details")
		delete(additionalProperties, "from_address")
		delete(additionalProperties, "gift_exchange_uuid")
		delete(additionalProperties, "gift_transaction_detail")
		delete(additionalProperties, "gift_transaction_uuid")
		delete(additionalProperties, "is_pending_transaction")
		delete(additionalProperties, "nft_token_id")
		delete(additionalProperties, "pal_transaction_detail")
		delete(additionalProperties, "price")
		delete(additionalProperties, "race_id")
		delete(additionalProperties, "race_ranks")
		delete(additionalProperties, "status")
		delete(additionalProperties, "to_address")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "transaction_code")
		delete(additionalProperties, "transaction_hash_hex")
		delete(additionalProperties, "transaction_history_type")
		delete(additionalProperties, "transaction_type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "value_symbol")
		delete(additionalProperties, "value_text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelWeb3WalletTransactionHistory struct {
	value *ModelWeb3WalletTransactionHistory
	isSet bool
}

func (v NullableModelWeb3WalletTransactionHistory) Get() *ModelWeb3WalletTransactionHistory {
	return v.value
}

func (v *NullableModelWeb3WalletTransactionHistory) Set(val *ModelWeb3WalletTransactionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableModelWeb3WalletTransactionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableModelWeb3WalletTransactionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelWeb3WalletTransactionHistory(val *ModelWeb3WalletTransactionHistory) *NullableModelWeb3WalletTransactionHistory {
	return &NullableModelWeb3WalletTransactionHistory{value: val, isSet: true}
}

func (v NullableModelWeb3WalletTransactionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelWeb3WalletTransactionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


