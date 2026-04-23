
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TokenDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenDTO{}

// TokenDTO struct for TokenDTO
type TokenDTO struct {
	Balance NullableString `json:"balance,omitempty"`
	Bridge NullableBridgeDTO `json:"bridge,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	Decimals NullableInt32 `json:"decimals,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	Price NullablePriceDTO `json:"price,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Swap NullableSwapDTO `json:"swap,omitempty"`
	Symbol NullableString `json:"symbol,omitempty"`
	TokenType NullableString `json:"token_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenDTO TokenDTO

// NewTokenDTO instantiates a new TokenDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenDTO() *TokenDTO {
	this := TokenDTO{}
	return &this
}

// NewTokenDTOWithDefaults instantiates a new TokenDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDTOWithDefaults() *TokenDTO {
	this := TokenDTO{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetBalance() string {
	if o == nil || IsNil(o.Balance.Get()) {
		var ret string
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *TokenDTO) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableString and assigns it to the Balance field.
func (o *TokenDTO) SetBalance(v string) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *TokenDTO) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *TokenDTO) UnsetBalance() {
	o.Balance.Unset()
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetBridge() BridgeDTO {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret BridgeDTO
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetBridgeOk() (*BridgeDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *TokenDTO) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableBridgeDTO and assigns it to the Bridge field.
func (o *TokenDTO) SetBridge(v BridgeDTO) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *TokenDTO) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *TokenDTO) UnsetBridge() {
	o.Bridge.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *TokenDTO) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *TokenDTO) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *TokenDTO) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *TokenDTO) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetDecimals returns the Decimals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetDecimals() int32 {
	if o == nil || IsNil(o.Decimals.Get()) {
		var ret int32
		return ret
	}
	return *o.Decimals.Get()
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Decimals.Get(), o.Decimals.IsSet()
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenDTO) HasDecimals() bool {
	if o != nil && o.Decimals.IsSet() {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given NullableInt32 and assigns it to the Decimals field.
func (o *TokenDTO) SetDecimals(v int32) {
	o.Decimals.Set(&v)
}
// SetDecimalsNil sets the value for Decimals to be an explicit nil
func (o *TokenDTO) SetDecimalsNil() {
	o.Decimals.Set(nil)
}

// UnsetDecimals ensures that no value is present for Decimals, not even an explicit nil
func (o *TokenDTO) UnsetDecimals() {
	o.Decimals.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *TokenDTO) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *TokenDTO) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *TokenDTO) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *TokenDTO) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetPrice() PriceDTO {
	if o == nil || IsNil(o.Price.Get()) {
		var ret PriceDTO
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetPriceOk() (*PriceDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *TokenDTO) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullablePriceDTO and assigns it to the Price field.
func (o *TokenDTO) SetPrice(v PriceDTO) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *TokenDTO) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *TokenDTO) UnsetPrice() {
	o.Price.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *TokenDTO) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *TokenDTO) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *TokenDTO) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *TokenDTO) UnsetPriority() {
	o.Priority.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TokenDTO) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TokenDTO) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TokenDTO) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TokenDTO) UnsetStatus() {
	o.Status.Unset()
}

// GetSwap returns the Swap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetSwap() SwapDTO {
	if o == nil || IsNil(o.Swap.Get()) {
		var ret SwapDTO
		return ret
	}
	return *o.Swap.Get()
}

// GetSwapOk returns a tuple with the Swap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetSwapOk() (*SwapDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Swap.Get(), o.Swap.IsSet()
}

// HasSwap returns a boolean if a field has been set.
func (o *TokenDTO) HasSwap() bool {
	if o != nil && o.Swap.IsSet() {
		return true
	}

	return false
}

// SetSwap gets a reference to the given NullableSwapDTO and assigns it to the Swap field.
func (o *TokenDTO) SetSwap(v SwapDTO) {
	o.Swap.Set(&v)
}
// SetSwapNil sets the value for Swap to be an explicit nil
func (o *TokenDTO) SetSwapNil() {
	o.Swap.Set(nil)
}

// UnsetSwap ensures that no value is present for Swap, not even an explicit nil
func (o *TokenDTO) UnsetSwap() {
	o.Swap.Unset()
}

// GetSymbol returns the Symbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetSymbol() string {
	if o == nil || IsNil(o.Symbol.Get()) {
		var ret string
		return ret
	}
	return *o.Symbol.Get()
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Symbol.Get(), o.Symbol.IsSet()
}

// HasSymbol returns a boolean if a field has been set.
func (o *TokenDTO) HasSymbol() bool {
	if o != nil && o.Symbol.IsSet() {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given NullableString and assigns it to the Symbol field.
func (o *TokenDTO) SetSymbol(v string) {
	o.Symbol.Set(&v)
}
// SetSymbolNil sets the value for Symbol to be an explicit nil
func (o *TokenDTO) SetSymbolNil() {
	o.Symbol.Set(nil)
}

// UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
func (o *TokenDTO) UnsetSymbol() {
	o.Symbol.Unset()
}

// GetTokenType returns the TokenType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenDTO) GetTokenType() string {
	if o == nil || IsNil(o.TokenType.Get()) {
		var ret string
		return ret
	}
	return *o.TokenType.Get()
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenDTO) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenType.Get(), o.TokenType.IsSet()
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenDTO) HasTokenType() bool {
	if o != nil && o.TokenType.IsSet() {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given NullableString and assigns it to the TokenType field.
func (o *TokenDTO) SetTokenType(v string) {
	o.TokenType.Set(&v)
}
// SetTokenTypeNil sets the value for TokenType to be an explicit nil
func (o *TokenDTO) SetTokenTypeNil() {
	o.TokenType.Set(nil)
}

// UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
func (o *TokenDTO) UnsetTokenType() {
	o.TokenType.Unset()
}

func (o TokenDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.Decimals.IsSet() {
		toSerialize["decimals"] = o.Decimals.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Swap.IsSet() {
		toSerialize["swap"] = o.Swap.Get()
	}
	if o.Symbol.IsSet() {
		toSerialize["symbol"] = o.Symbol.Get()
	}
	if o.TokenType.IsSet() {
		toSerialize["token_type"] = o.TokenType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenDTO) UnmarshalJSON(data []byte) (err error) {
	varTokenDTO := _TokenDTO{}

	err = json.Unmarshal(data, &varTokenDTO)

	if err != nil {
		return err
	}

	*o = TokenDTO(varTokenDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "decimals")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "price")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "status")
		delete(additionalProperties, "swap")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenDTO struct {
	value *TokenDTO
	isSet bool
}

func (v NullableTokenDTO) Get() *TokenDTO {
	return v.value
}

func (v *NullableTokenDTO) Set(val *TokenDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenDTO(val *TokenDTO) *NullableTokenDTO {
	return &NullableTokenDTO{value: val, isSet: true}
}

func (v NullableTokenDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


