
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplActivityDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplActivityDTO{}

// EmplActivityDTO struct for EmplActivityDTO
type EmplActivityDTO struct {
	Amount NullableString `json:"amount,omitempty"`
	AvatarFrameUserUuid NullableString `json:"avatar_frame_user_uuid,omitempty"`
	BattleHistoryId NullableInt64 `json:"battle_history_id,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	Details NullablePalDetailsDTO `json:"details,omitempty"`
	Empl NullableEmplDTO `json:"empl,omitempty"`
	GiftExchangeUuid NullableString `json:"gift_exchange_uuid,omitempty"`
	GiftTransactionUuid NullableString `json:"gift_transaction_uuid,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MoneyAmount NullableFloat64 `json:"money_amount,omitempty"`
	OccurredAt NullableInt64 `json:"occurred_at,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	RaceId NullableString `json:"race_id,omitempty"`
	Status NullableString `json:"status,omitempty"`
	TokenAddress NullableString `json:"token_address,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	TransactionCode NullableString `json:"transaction_code,omitempty"`
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplActivityDTO EmplActivityDTO

// NewEmplActivityDTO instantiates a new EmplActivityDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplActivityDTO() *EmplActivityDTO {
	this := EmplActivityDTO{}
	return &this
}

// NewEmplActivityDTOWithDefaults instantiates a new EmplActivityDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplActivityDTOWithDefaults() *EmplActivityDTO {
	this := EmplActivityDTO{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *EmplActivityDTO) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *EmplActivityDTO) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *EmplActivityDTO) UnsetAmount() {
	o.Amount.Unset()
}

// GetAvatarFrameUserUuid returns the AvatarFrameUserUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetAvatarFrameUserUuid() string {
	if o == nil || IsNil(o.AvatarFrameUserUuid.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarFrameUserUuid.Get()
}

// GetAvatarFrameUserUuidOk returns a tuple with the AvatarFrameUserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetAvatarFrameUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarFrameUserUuid.Get(), o.AvatarFrameUserUuid.IsSet()
}

// HasAvatarFrameUserUuid returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasAvatarFrameUserUuid() bool {
	if o != nil && o.AvatarFrameUserUuid.IsSet() {
		return true
	}

	return false
}

// SetAvatarFrameUserUuid gets a reference to the given NullableString and assigns it to the AvatarFrameUserUuid field.
func (o *EmplActivityDTO) SetAvatarFrameUserUuid(v string) {
	o.AvatarFrameUserUuid.Set(&v)
}
// SetAvatarFrameUserUuidNil sets the value for AvatarFrameUserUuid to be an explicit nil
func (o *EmplActivityDTO) SetAvatarFrameUserUuidNil() {
	o.AvatarFrameUserUuid.Set(nil)
}

// UnsetAvatarFrameUserUuid ensures that no value is present for AvatarFrameUserUuid, not even an explicit nil
func (o *EmplActivityDTO) UnsetAvatarFrameUserUuid() {
	o.AvatarFrameUserUuid.Unset()
}

// GetBattleHistoryId returns the BattleHistoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetBattleHistoryId() int64 {
	if o == nil || IsNil(o.BattleHistoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.BattleHistoryId.Get()
}

// GetBattleHistoryIdOk returns a tuple with the BattleHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetBattleHistoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BattleHistoryId.Get(), o.BattleHistoryId.IsSet()
}

// HasBattleHistoryId returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasBattleHistoryId() bool {
	if o != nil && o.BattleHistoryId.IsSet() {
		return true
	}

	return false
}

// SetBattleHistoryId gets a reference to the given NullableInt64 and assigns it to the BattleHistoryId field.
func (o *EmplActivityDTO) SetBattleHistoryId(v int64) {
	o.BattleHistoryId.Set(&v)
}
// SetBattleHistoryIdNil sets the value for BattleHistoryId to be an explicit nil
func (o *EmplActivityDTO) SetBattleHistoryIdNil() {
	o.BattleHistoryId.Set(nil)
}

// UnsetBattleHistoryId ensures that no value is present for BattleHistoryId, not even an explicit nil
func (o *EmplActivityDTO) UnsetBattleHistoryId() {
	o.BattleHistoryId.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *EmplActivityDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *EmplActivityDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *EmplActivityDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetDetails() PalDetailsDTO {
	if o == nil || IsNil(o.Details.Get()) {
		var ret PalDetailsDTO
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetDetailsOk() (*PalDetailsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullablePalDetailsDTO and assigns it to the Details field.
func (o *EmplActivityDTO) SetDetails(v PalDetailsDTO) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *EmplActivityDTO) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *EmplActivityDTO) UnsetDetails() {
	o.Details.Unset()
}

// GetEmpl returns the Empl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetEmpl() EmplDTO {
	if o == nil || IsNil(o.Empl.Get()) {
		var ret EmplDTO
		return ret
	}
	return *o.Empl.Get()
}

// GetEmplOk returns a tuple with the Empl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetEmplOk() (*EmplDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Empl.Get(), o.Empl.IsSet()
}

// HasEmpl returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasEmpl() bool {
	if o != nil && o.Empl.IsSet() {
		return true
	}

	return false
}

// SetEmpl gets a reference to the given NullableEmplDTO and assigns it to the Empl field.
func (o *EmplActivityDTO) SetEmpl(v EmplDTO) {
	o.Empl.Set(&v)
}
// SetEmplNil sets the value for Empl to be an explicit nil
func (o *EmplActivityDTO) SetEmplNil() {
	o.Empl.Set(nil)
}

// UnsetEmpl ensures that no value is present for Empl, not even an explicit nil
func (o *EmplActivityDTO) UnsetEmpl() {
	o.Empl.Unset()
}

// GetGiftExchangeUuid returns the GiftExchangeUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetGiftExchangeUuid() string {
	if o == nil || IsNil(o.GiftExchangeUuid.Get()) {
		var ret string
		return ret
	}
	return *o.GiftExchangeUuid.Get()
}

// GetGiftExchangeUuidOk returns a tuple with the GiftExchangeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetGiftExchangeUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftExchangeUuid.Get(), o.GiftExchangeUuid.IsSet()
}

// HasGiftExchangeUuid returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasGiftExchangeUuid() bool {
	if o != nil && o.GiftExchangeUuid.IsSet() {
		return true
	}

	return false
}

// SetGiftExchangeUuid gets a reference to the given NullableString and assigns it to the GiftExchangeUuid field.
func (o *EmplActivityDTO) SetGiftExchangeUuid(v string) {
	o.GiftExchangeUuid.Set(&v)
}
// SetGiftExchangeUuidNil sets the value for GiftExchangeUuid to be an explicit nil
func (o *EmplActivityDTO) SetGiftExchangeUuidNil() {
	o.GiftExchangeUuid.Set(nil)
}

// UnsetGiftExchangeUuid ensures that no value is present for GiftExchangeUuid, not even an explicit nil
func (o *EmplActivityDTO) UnsetGiftExchangeUuid() {
	o.GiftExchangeUuid.Unset()
}

// GetGiftTransactionUuid returns the GiftTransactionUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetGiftTransactionUuid() string {
	if o == nil || IsNil(o.GiftTransactionUuid.Get()) {
		var ret string
		return ret
	}
	return *o.GiftTransactionUuid.Get()
}

// GetGiftTransactionUuidOk returns a tuple with the GiftTransactionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetGiftTransactionUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftTransactionUuid.Get(), o.GiftTransactionUuid.IsSet()
}

// HasGiftTransactionUuid returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasGiftTransactionUuid() bool {
	if o != nil && o.GiftTransactionUuid.IsSet() {
		return true
	}

	return false
}

// SetGiftTransactionUuid gets a reference to the given NullableString and assigns it to the GiftTransactionUuid field.
func (o *EmplActivityDTO) SetGiftTransactionUuid(v string) {
	o.GiftTransactionUuid.Set(&v)
}
// SetGiftTransactionUuidNil sets the value for GiftTransactionUuid to be an explicit nil
func (o *EmplActivityDTO) SetGiftTransactionUuidNil() {
	o.GiftTransactionUuid.Set(nil)
}

// UnsetGiftTransactionUuid ensures that no value is present for GiftTransactionUuid, not even an explicit nil
func (o *EmplActivityDTO) UnsetGiftTransactionUuid() {
	o.GiftTransactionUuid.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *EmplActivityDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *EmplActivityDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EmplActivityDTO) UnsetId() {
	o.Id.Unset()
}

// GetMoneyAmount returns the MoneyAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetMoneyAmount() float64 {
	if o == nil || IsNil(o.MoneyAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MoneyAmount.Get()
}

// GetMoneyAmountOk returns a tuple with the MoneyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetMoneyAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoneyAmount.Get(), o.MoneyAmount.IsSet()
}

// HasMoneyAmount returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasMoneyAmount() bool {
	if o != nil && o.MoneyAmount.IsSet() {
		return true
	}

	return false
}

// SetMoneyAmount gets a reference to the given NullableFloat64 and assigns it to the MoneyAmount field.
func (o *EmplActivityDTO) SetMoneyAmount(v float64) {
	o.MoneyAmount.Set(&v)
}
// SetMoneyAmountNil sets the value for MoneyAmount to be an explicit nil
func (o *EmplActivityDTO) SetMoneyAmountNil() {
	o.MoneyAmount.Set(nil)
}

// UnsetMoneyAmount ensures that no value is present for MoneyAmount, not even an explicit nil
func (o *EmplActivityDTO) UnsetMoneyAmount() {
	o.MoneyAmount.Unset()
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetOccurredAt() int64 {
	if o == nil || IsNil(o.OccurredAt.Get()) {
		var ret int64
		return ret
	}
	return *o.OccurredAt.Get()
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetOccurredAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OccurredAt.Get(), o.OccurredAt.IsSet()
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasOccurredAt() bool {
	if o != nil && o.OccurredAt.IsSet() {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given NullableInt64 and assigns it to the OccurredAt field.
func (o *EmplActivityDTO) SetOccurredAt(v int64) {
	o.OccurredAt.Set(&v)
}
// SetOccurredAtNil sets the value for OccurredAt to be an explicit nil
func (o *EmplActivityDTO) SetOccurredAtNil() {
	o.OccurredAt.Set(nil)
}

// UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
func (o *EmplActivityDTO) UnsetOccurredAt() {
	o.OccurredAt.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *EmplActivityDTO) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *EmplActivityDTO) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *EmplActivityDTO) UnsetPrice() {
	o.Price.Unset()
}

// GetRaceId returns the RaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetRaceId() string {
	if o == nil || IsNil(o.RaceId.Get()) {
		var ret string
		return ret
	}
	return *o.RaceId.Get()
}

// GetRaceIdOk returns a tuple with the RaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetRaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceId.Get(), o.RaceId.IsSet()
}

// HasRaceId returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasRaceId() bool {
	if o != nil && o.RaceId.IsSet() {
		return true
	}

	return false
}

// SetRaceId gets a reference to the given NullableString and assigns it to the RaceId field.
func (o *EmplActivityDTO) SetRaceId(v string) {
	o.RaceId.Set(&v)
}
// SetRaceIdNil sets the value for RaceId to be an explicit nil
func (o *EmplActivityDTO) SetRaceIdNil() {
	o.RaceId.Set(nil)
}

// UnsetRaceId ensures that no value is present for RaceId, not even an explicit nil
func (o *EmplActivityDTO) UnsetRaceId() {
	o.RaceId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *EmplActivityDTO) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *EmplActivityDTO) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *EmplActivityDTO) UnsetStatus() {
	o.Status.Unset()
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetTokenAddress() string {
	if o == nil || IsNil(o.TokenAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenAddress.Get()
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenAddress.Get(), o.TokenAddress.IsSet()
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasTokenAddress() bool {
	if o != nil && o.TokenAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given NullableString and assigns it to the TokenAddress field.
func (o *EmplActivityDTO) SetTokenAddress(v string) {
	o.TokenAddress.Set(&v)
}
// SetTokenAddressNil sets the value for TokenAddress to be an explicit nil
func (o *EmplActivityDTO) SetTokenAddressNil() {
	o.TokenAddress.Set(nil)
}

// UnsetTokenAddress ensures that no value is present for TokenAddress, not even an explicit nil
func (o *EmplActivityDTO) UnsetTokenAddress() {
	o.TokenAddress.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *EmplActivityDTO) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *EmplActivityDTO) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *EmplActivityDTO) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasTransactionCode() bool {
	if o != nil && o.TransactionCode.IsSet() {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given NullableString and assigns it to the TransactionCode field.
func (o *EmplActivityDTO) SetTransactionCode(v string) {
	o.TransactionCode.Set(&v)
}
// SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil
func (o *EmplActivityDTO) SetTransactionCodeNil() {
	o.TransactionCode.Set(nil)
}

// UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
func (o *EmplActivityDTO) UnsetTransactionCode() {
	o.TransactionCode.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDTO) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDTO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *EmplActivityDTO) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *EmplActivityDTO) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *EmplActivityDTO) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *EmplActivityDTO) UnsetType() {
	o.Type.Unset()
}

func (o EmplActivityDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplActivityDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.AvatarFrameUserUuid.IsSet() {
		toSerialize["avatar_frame_user_uuid"] = o.AvatarFrameUserUuid.Get()
	}
	if o.BattleHistoryId.IsSet() {
		toSerialize["battle_history_id"] = o.BattleHistoryId.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.Empl.IsSet() {
		toSerialize["empl"] = o.Empl.Get()
	}
	if o.GiftExchangeUuid.IsSet() {
		toSerialize["gift_exchange_uuid"] = o.GiftExchangeUuid.Get()
	}
	if o.GiftTransactionUuid.IsSet() {
		toSerialize["gift_transaction_uuid"] = o.GiftTransactionUuid.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MoneyAmount.IsSet() {
		toSerialize["money_amount"] = o.MoneyAmount.Get()
	}
	if o.OccurredAt.IsSet() {
		toSerialize["occurred_at"] = o.OccurredAt.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.RaceId.IsSet() {
		toSerialize["race_id"] = o.RaceId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.TokenAddress.IsSet() {
		toSerialize["token_address"] = o.TokenAddress.Get()
	}
	if o.TokenSymbol.IsSet() {
		toSerialize["token_symbol"] = o.TokenSymbol.Get()
	}
	if o.TransactionCode.IsSet() {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplActivityDTO) UnmarshalJSON(data []byte) (err error) {
	varEmplActivityDTO := _EmplActivityDTO{}

	err = json.Unmarshal(data, &varEmplActivityDTO)

	if err != nil {
		return err
	}

	*o = EmplActivityDTO(varEmplActivityDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "avatar_frame_user_uuid")
		delete(additionalProperties, "battle_history_id")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "details")
		delete(additionalProperties, "empl")
		delete(additionalProperties, "gift_exchange_uuid")
		delete(additionalProperties, "gift_transaction_uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "money_amount")
		delete(additionalProperties, "occurred_at")
		delete(additionalProperties, "price")
		delete(additionalProperties, "race_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "token_address")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "transaction_code")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplActivityDTO struct {
	value *EmplActivityDTO
	isSet bool
}

func (v NullableEmplActivityDTO) Get() *EmplActivityDTO {
	return v.value
}

func (v *NullableEmplActivityDTO) Set(val *EmplActivityDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplActivityDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplActivityDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplActivityDTO(val *EmplActivityDTO) *NullableEmplActivityDTO {
	return &NullableEmplActivityDTO{value: val, isSet: true}
}

func (v NullableEmplActivityDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplActivityDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


