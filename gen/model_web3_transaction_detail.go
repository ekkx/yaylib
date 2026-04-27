
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3TransactionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3TransactionDetail{}

// Web3TransactionDetail struct for Web3TransactionDetail
type Web3TransactionDetail struct {
	Amount NullableString `json:"amount,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	EmplAmount NullableFloat64 `json:"empl_amount,omitempty"`
	Error map[string]interface{} `json:"error,omitempty"`
	FinalStatusUpdatedAtMillis NullableInt64 `json:"final_status_updated_at_millis,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MaxTokenAmount NullableString `json:"max_token_amount,omitempty"`
	MinTokenAmount NullableString `json:"min_token_amount,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
	TokenOutAddress NullableString `json:"token_out_address,omitempty"`
	TokenOutAmount NullableString `json:"token_out_amount,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	TransactionFee NullableFloat64 `json:"transaction_fee,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3TransactionDetail Web3TransactionDetail

// NewWeb3TransactionDetail instantiates a new Web3TransactionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3TransactionDetail() *Web3TransactionDetail {
	this := Web3TransactionDetail{}
	return &this
}

// NewWeb3TransactionDetailWithDefaults instantiates a new Web3TransactionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TransactionDetailWithDefaults() *Web3TransactionDetail {
	this := Web3TransactionDetail{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *Web3TransactionDetail) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Web3TransactionDetail) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Web3TransactionDetail) UnsetAmount() {
	o.Amount.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3TransactionDetail) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3TransactionDetail) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3TransactionDetail) UnsetChainId() {
	o.ChainId.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Web3TransactionDetail) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Web3TransactionDetail) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Web3TransactionDetail) UnsetCurrency() {
	o.Currency.Unset()
}

// GetEmplAmount returns the EmplAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetEmplAmount() float64 {
	if o == nil || IsNil(o.EmplAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.EmplAmount.Get()
}

// GetEmplAmountOk returns a tuple with the EmplAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetEmplAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplAmount.Get(), o.EmplAmount.IsSet()
}

// HasEmplAmount returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasEmplAmount() bool {
	if o != nil && o.EmplAmount.IsSet() {
		return true
	}

	return false
}

// SetEmplAmount gets a reference to the given NullableFloat64 and assigns it to the EmplAmount field.
func (o *Web3TransactionDetail) SetEmplAmount(v float64) {
	o.EmplAmount.Set(&v)
}
// SetEmplAmountNil sets the value for EmplAmount to be an explicit nil
func (o *Web3TransactionDetail) SetEmplAmountNil() {
	o.EmplAmount.Set(nil)
}

// UnsetEmplAmount ensures that no value is present for EmplAmount, not even an explicit nil
func (o *Web3TransactionDetail) UnsetEmplAmount() {
	o.EmplAmount.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetError() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *Web3TransactionDetail) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetFinalStatusUpdatedAtMillis returns the FinalStatusUpdatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetFinalStatusUpdatedAtMillis() int64 {
	if o == nil || IsNil(o.FinalStatusUpdatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.FinalStatusUpdatedAtMillis.Get()
}

// GetFinalStatusUpdatedAtMillisOk returns a tuple with the FinalStatusUpdatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetFinalStatusUpdatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinalStatusUpdatedAtMillis.Get(), o.FinalStatusUpdatedAtMillis.IsSet()
}

// HasFinalStatusUpdatedAtMillis returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasFinalStatusUpdatedAtMillis() bool {
	if o != nil && o.FinalStatusUpdatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetFinalStatusUpdatedAtMillis gets a reference to the given NullableInt64 and assigns it to the FinalStatusUpdatedAtMillis field.
func (o *Web3TransactionDetail) SetFinalStatusUpdatedAtMillis(v int64) {
	o.FinalStatusUpdatedAtMillis.Set(&v)
}
// SetFinalStatusUpdatedAtMillisNil sets the value for FinalStatusUpdatedAtMillis to be an explicit nil
func (o *Web3TransactionDetail) SetFinalStatusUpdatedAtMillisNil() {
	o.FinalStatusUpdatedAtMillis.Set(nil)
}

// UnsetFinalStatusUpdatedAtMillis ensures that no value is present for FinalStatusUpdatedAtMillis, not even an explicit nil
func (o *Web3TransactionDetail) UnsetFinalStatusUpdatedAtMillis() {
	o.FinalStatusUpdatedAtMillis.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Web3TransactionDetail) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Web3TransactionDetail) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Web3TransactionDetail) UnsetId() {
	o.Id.Unset()
}

// GetMaxTokenAmount returns the MaxTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetMaxTokenAmount() string {
	if o == nil || IsNil(o.MaxTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MaxTokenAmount.Get()
}

// GetMaxTokenAmountOk returns a tuple with the MaxTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetMaxTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokenAmount.Get(), o.MaxTokenAmount.IsSet()
}

// HasMaxTokenAmount returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasMaxTokenAmount() bool {
	if o != nil && o.MaxTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxTokenAmount gets a reference to the given NullableString and assigns it to the MaxTokenAmount field.
func (o *Web3TransactionDetail) SetMaxTokenAmount(v string) {
	o.MaxTokenAmount.Set(&v)
}
// SetMaxTokenAmountNil sets the value for MaxTokenAmount to be an explicit nil
func (o *Web3TransactionDetail) SetMaxTokenAmountNil() {
	o.MaxTokenAmount.Set(nil)
}

// UnsetMaxTokenAmount ensures that no value is present for MaxTokenAmount, not even an explicit nil
func (o *Web3TransactionDetail) UnsetMaxTokenAmount() {
	o.MaxTokenAmount.Unset()
}

// GetMinTokenAmount returns the MinTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetMinTokenAmount() string {
	if o == nil || IsNil(o.MinTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MinTokenAmount.Get()
}

// GetMinTokenAmountOk returns a tuple with the MinTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetMinTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTokenAmount.Get(), o.MinTokenAmount.IsSet()
}

// HasMinTokenAmount returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasMinTokenAmount() bool {
	if o != nil && o.MinTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMinTokenAmount gets a reference to the given NullableString and assigns it to the MinTokenAmount field.
func (o *Web3TransactionDetail) SetMinTokenAmount(v string) {
	o.MinTokenAmount.Set(&v)
}
// SetMinTokenAmountNil sets the value for MinTokenAmount to be an explicit nil
func (o *Web3TransactionDetail) SetMinTokenAmountNil() {
	o.MinTokenAmount.Set(nil)
}

// UnsetMinTokenAmount ensures that no value is present for MinTokenAmount, not even an explicit nil
func (o *Web3TransactionDetail) UnsetMinTokenAmount() {
	o.MinTokenAmount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *Web3TransactionDetail) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// GetTokenOutAddress returns the TokenOutAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetTokenOutAddress() string {
	if o == nil || IsNil(o.TokenOutAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAddress.Get()
}

// GetTokenOutAddressOk returns a tuple with the TokenOutAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetTokenOutAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAddress.Get(), o.TokenOutAddress.IsSet()
}

// HasTokenOutAddress returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasTokenOutAddress() bool {
	if o != nil && o.TokenOutAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAddress gets a reference to the given NullableString and assigns it to the TokenOutAddress field.
func (o *Web3TransactionDetail) SetTokenOutAddress(v string) {
	o.TokenOutAddress.Set(&v)
}
// SetTokenOutAddressNil sets the value for TokenOutAddress to be an explicit nil
func (o *Web3TransactionDetail) SetTokenOutAddressNil() {
	o.TokenOutAddress.Set(nil)
}

// UnsetTokenOutAddress ensures that no value is present for TokenOutAddress, not even an explicit nil
func (o *Web3TransactionDetail) UnsetTokenOutAddress() {
	o.TokenOutAddress.Unset()
}

// GetTokenOutAmount returns the TokenOutAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetTokenOutAmount() string {
	if o == nil || IsNil(o.TokenOutAmount.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAmount.Get()
}

// GetTokenOutAmountOk returns a tuple with the TokenOutAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetTokenOutAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAmount.Get(), o.TokenOutAmount.IsSet()
}

// HasTokenOutAmount returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasTokenOutAmount() bool {
	if o != nil && o.TokenOutAmount.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAmount gets a reference to the given NullableString and assigns it to the TokenOutAmount field.
func (o *Web3TransactionDetail) SetTokenOutAmount(v string) {
	o.TokenOutAmount.Set(&v)
}
// SetTokenOutAmountNil sets the value for TokenOutAmount to be an explicit nil
func (o *Web3TransactionDetail) SetTokenOutAmountNil() {
	o.TokenOutAmount.Set(nil)
}

// UnsetTokenOutAmount ensures that no value is present for TokenOutAmount, not even an explicit nil
func (o *Web3TransactionDetail) UnsetTokenOutAmount() {
	o.TokenOutAmount.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *Web3TransactionDetail) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *Web3TransactionDetail) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *Web3TransactionDetail) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetTransactionFee() float64 {
	if o == nil || IsNil(o.TransactionFee.Get()) {
		var ret float64
		return ret
	}
	return *o.TransactionFee.Get()
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetTransactionFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionFee.Get(), o.TransactionFee.IsSet()
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasTransactionFee() bool {
	if o != nil && o.TransactionFee.IsSet() {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given NullableFloat64 and assigns it to the TransactionFee field.
func (o *Web3TransactionDetail) SetTransactionFee(v float64) {
	o.TransactionFee.Set(&v)
}
// SetTransactionFeeNil sets the value for TransactionFee to be an explicit nil
func (o *Web3TransactionDetail) SetTransactionFeeNil() {
	o.TransactionFee.Set(nil)
}

// UnsetTransactionFee ensures that no value is present for TransactionFee, not even an explicit nil
func (o *Web3TransactionDetail) UnsetTransactionFee() {
	o.TransactionFee.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionDetail) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionDetail) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *Web3TransactionDetail) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *Web3TransactionDetail) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *Web3TransactionDetail) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *Web3TransactionDetail) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o Web3TransactionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3TransactionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.EmplAmount.IsSet() {
		toSerialize["empl_amount"] = o.EmplAmount.Get()
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.FinalStatusUpdatedAtMillis.IsSet() {
		toSerialize["final_status_updated_at_millis"] = o.FinalStatusUpdatedAtMillis.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MaxTokenAmount.IsSet() {
		toSerialize["max_token_amount"] = o.MaxTokenAmount.Get()
	}
	if o.MinTokenAmount.IsSet() {
		toSerialize["min_token_amount"] = o.MinTokenAmount.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TokenOutAddress.IsSet() {
		toSerialize["token_out_address"] = o.TokenOutAddress.Get()
	}
	if o.TokenOutAmount.IsSet() {
		toSerialize["token_out_amount"] = o.TokenOutAmount.Get()
	}
	if o.TokenSymbol.IsSet() {
		toSerialize["token_symbol"] = o.TokenSymbol.Get()
	}
	if o.TransactionFee.IsSet() {
		toSerialize["transaction_fee"] = o.TransactionFee.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3TransactionDetail) UnmarshalJSON(data []byte) (err error) {
	varWeb3TransactionDetail := _Web3TransactionDetail{}

	err = json.Unmarshal(data, &varWeb3TransactionDetail)

	if err != nil {
		return err
	}

	*o = Web3TransactionDetail(varWeb3TransactionDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "empl_amount")
		delete(additionalProperties, "error")
		delete(additionalProperties, "final_status_updated_at_millis")
		delete(additionalProperties, "id")
		delete(additionalProperties, "max_token_amount")
		delete(additionalProperties, "min_token_amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "token_out_address")
		delete(additionalProperties, "token_out_amount")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "transaction_fee")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3TransactionDetail struct {
	value *Web3TransactionDetail
	isSet bool
}

func (v NullableWeb3TransactionDetail) Get() *Web3TransactionDetail {
	return v.value
}

func (v *NullableWeb3TransactionDetail) Set(val *Web3TransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3TransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3TransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3TransactionDetail(val *Web3TransactionDetail) *NullableWeb3TransactionDetail {
	return &NullableWeb3TransactionDetail{value: val, isSet: true}
}

func (v NullableWeb3TransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3TransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


