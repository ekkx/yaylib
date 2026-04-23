
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Result type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Result{}

// Result struct for Result
type Result struct {
	Amount NullableString `json:"amount,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	Details NullableDetails `json:"details,omitempty"`
	EmplTransaction NullableEmplTransaction `json:"empl_transaction,omitempty"`
	FinalStatusUpdatedAt NullableInt64 `json:"final_status_updated_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Status NullableString `json:"status,omitempty"`
	TransactionCode NullableString `json:"transaction_code,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Result Result

// NewResult instantiates a new Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResult() *Result {
	this := Result{}
	return &this
}

// NewResultWithDefaults instantiates a new Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Result) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *Result) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Result) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Result) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Result) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Result) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Result) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Result) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetDetails() Details {
	if o == nil || IsNil(o.Details.Get()) {
		var ret Details
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetDetailsOk() (*Details, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *Result) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableDetails and assigns it to the Details field.
func (o *Result) SetDetails(v Details) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *Result) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *Result) UnsetDetails() {
	o.Details.Unset()
}

// GetEmplTransaction returns the EmplTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetEmplTransaction() EmplTransaction {
	if o == nil || IsNil(o.EmplTransaction.Get()) {
		var ret EmplTransaction
		return ret
	}
	return *o.EmplTransaction.Get()
}

// GetEmplTransactionOk returns a tuple with the EmplTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetEmplTransactionOk() (*EmplTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplTransaction.Get(), o.EmplTransaction.IsSet()
}

// HasEmplTransaction returns a boolean if a field has been set.
func (o *Result) HasEmplTransaction() bool {
	if o != nil && o.EmplTransaction.IsSet() {
		return true
	}

	return false
}

// SetEmplTransaction gets a reference to the given NullableEmplTransaction and assigns it to the EmplTransaction field.
func (o *Result) SetEmplTransaction(v EmplTransaction) {
	o.EmplTransaction.Set(&v)
}
// SetEmplTransactionNil sets the value for EmplTransaction to be an explicit nil
func (o *Result) SetEmplTransactionNil() {
	o.EmplTransaction.Set(nil)
}

// UnsetEmplTransaction ensures that no value is present for EmplTransaction, not even an explicit nil
func (o *Result) UnsetEmplTransaction() {
	o.EmplTransaction.Unset()
}

// GetFinalStatusUpdatedAt returns the FinalStatusUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetFinalStatusUpdatedAt() int64 {
	if o == nil || IsNil(o.FinalStatusUpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.FinalStatusUpdatedAt.Get()
}

// GetFinalStatusUpdatedAtOk returns a tuple with the FinalStatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetFinalStatusUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinalStatusUpdatedAt.Get(), o.FinalStatusUpdatedAt.IsSet()
}

// HasFinalStatusUpdatedAt returns a boolean if a field has been set.
func (o *Result) HasFinalStatusUpdatedAt() bool {
	if o != nil && o.FinalStatusUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetFinalStatusUpdatedAt gets a reference to the given NullableInt64 and assigns it to the FinalStatusUpdatedAt field.
func (o *Result) SetFinalStatusUpdatedAt(v int64) {
	o.FinalStatusUpdatedAt.Set(&v)
}
// SetFinalStatusUpdatedAtNil sets the value for FinalStatusUpdatedAt to be an explicit nil
func (o *Result) SetFinalStatusUpdatedAtNil() {
	o.FinalStatusUpdatedAt.Set(nil)
}

// UnsetFinalStatusUpdatedAt ensures that no value is present for FinalStatusUpdatedAt, not even an explicit nil
func (o *Result) UnsetFinalStatusUpdatedAt() {
	o.FinalStatusUpdatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Result) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Result) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Result) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Result) UnsetId() {
	o.Id.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Result) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Result) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Result) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Result) UnsetStatus() {
	o.Status.Unset()
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *Result) HasTransactionCode() bool {
	if o != nil && o.TransactionCode.IsSet() {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given NullableString and assigns it to the TransactionCode field.
func (o *Result) SetTransactionCode(v string) {
	o.TransactionCode.Set(&v)
}
// SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil
func (o *Result) SetTransactionCodeNil() {
	o.TransactionCode.Set(nil)
}

// UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
func (o *Result) UnsetTransactionCode() {
	o.TransactionCode.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Result) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Result) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *Result) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *Result) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *Result) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *Result) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.EmplTransaction.IsSet() {
		toSerialize["empl_transaction"] = o.EmplTransaction.Get()
	}
	if o.FinalStatusUpdatedAt.IsSet() {
		toSerialize["final_status_updated_at"] = o.FinalStatusUpdatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.TransactionCode.IsSet() {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Result) UnmarshalJSON(data []byte) (err error) {
	varResult := _Result{}

	err = json.Unmarshal(data, &varResult)

	if err != nil {
		return err
	}

	*o = Result(varResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "details")
		delete(additionalProperties, "empl_transaction")
		delete(additionalProperties, "final_status_updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "transaction_code")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResult struct {
	value *Result
	isSet bool
}

func (v NullableResult) Get() *Result {
	return v.value
}

func (v *NullableResult) Set(val *Result) {
	v.value = val
	v.isSet = true
}

func (v NullableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResult(val *Result) *NullableResult {
	return &NullableResult{value: val, isSet: true}
}

func (v NullableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


