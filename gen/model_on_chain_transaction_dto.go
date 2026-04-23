
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the OnChainTransactionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnChainTransactionDTO{}

// OnChainTransactionDTO struct for OnChainTransactionDTO
type OnChainTransactionDTO struct {
	Amount NullableString `json:"amount,omitempty"`
	AssetInfo NullableAssetInfoDTO `json:"asset_info,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	Description NullableString `json:"description,omitempty"`
	FromAddress NullableString `json:"from_address,omitempty"`
	Hash NullableString `json:"hash,omitempty"`
	Status NullableString `json:"status,omitempty"`
	ToAddress NullableString `json:"to_address,omitempty"`
	TokenContract NullableString `json:"token_contract,omitempty"`
	TokenType NullableString `json:"token_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnChainTransactionDTO OnChainTransactionDTO

// NewOnChainTransactionDTO instantiates a new OnChainTransactionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainTransactionDTO() *OnChainTransactionDTO {
	this := OnChainTransactionDTO{}
	return &this
}

// NewOnChainTransactionDTOWithDefaults instantiates a new OnChainTransactionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainTransactionDTOWithDefaults() *OnChainTransactionDTO {
	this := OnChainTransactionDTO{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *OnChainTransactionDTO) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *OnChainTransactionDTO) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetAmount() {
	o.Amount.Unset()
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetAssetInfo() AssetInfoDTO {
	if o == nil || IsNil(o.AssetInfo.Get()) {
		var ret AssetInfoDTO
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetAssetInfoOk() (*AssetInfoDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableAssetInfoDTO and assigns it to the AssetInfo field.
func (o *OnChainTransactionDTO) SetAssetInfo(v AssetInfoDTO) {
	o.AssetInfo.Set(&v)
}
// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *OnChainTransactionDTO) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *OnChainTransactionDTO) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *OnChainTransactionDTO) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *OnChainTransactionDTO) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *OnChainTransactionDTO) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OnChainTransactionDTO) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OnChainTransactionDTO) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetDescription() {
	o.Description.Unset()
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress.Get()) {
		var ret string
		return ret
	}
	return *o.FromAddress.Get()
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromAddress.Get(), o.FromAddress.IsSet()
}

// HasFromAddress returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasFromAddress() bool {
	if o != nil && o.FromAddress.IsSet() {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given NullableString and assigns it to the FromAddress field.
func (o *OnChainTransactionDTO) SetFromAddress(v string) {
	o.FromAddress.Set(&v)
}
// SetFromAddressNil sets the value for FromAddress to be an explicit nil
func (o *OnChainTransactionDTO) SetFromAddressNil() {
	o.FromAddress.Set(nil)
}

// UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetFromAddress() {
	o.FromAddress.Unset()
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *OnChainTransactionDTO) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *OnChainTransactionDTO) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetHash() {
	o.Hash.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *OnChainTransactionDTO) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *OnChainTransactionDTO) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetStatus() {
	o.Status.Unset()
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetToAddress() string {
	if o == nil || IsNil(o.ToAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ToAddress.Get()
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToAddress.Get(), o.ToAddress.IsSet()
}

// HasToAddress returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasToAddress() bool {
	if o != nil && o.ToAddress.IsSet() {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given NullableString and assigns it to the ToAddress field.
func (o *OnChainTransactionDTO) SetToAddress(v string) {
	o.ToAddress.Set(&v)
}
// SetToAddressNil sets the value for ToAddress to be an explicit nil
func (o *OnChainTransactionDTO) SetToAddressNil() {
	o.ToAddress.Set(nil)
}

// UnsetToAddress ensures that no value is present for ToAddress, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetToAddress() {
	o.ToAddress.Unset()
}

// GetTokenContract returns the TokenContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetTokenContract() string {
	if o == nil || IsNil(o.TokenContract.Get()) {
		var ret string
		return ret
	}
	return *o.TokenContract.Get()
}

// GetTokenContractOk returns a tuple with the TokenContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetTokenContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenContract.Get(), o.TokenContract.IsSet()
}

// HasTokenContract returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasTokenContract() bool {
	if o != nil && o.TokenContract.IsSet() {
		return true
	}

	return false
}

// SetTokenContract gets a reference to the given NullableString and assigns it to the TokenContract field.
func (o *OnChainTransactionDTO) SetTokenContract(v string) {
	o.TokenContract.Set(&v)
}
// SetTokenContractNil sets the value for TokenContract to be an explicit nil
func (o *OnChainTransactionDTO) SetTokenContractNil() {
	o.TokenContract.Set(nil)
}

// UnsetTokenContract ensures that no value is present for TokenContract, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetTokenContract() {
	o.TokenContract.Unset()
}

// GetTokenType returns the TokenType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionDTO) GetTokenType() string {
	if o == nil || IsNil(o.TokenType.Get()) {
		var ret string
		return ret
	}
	return *o.TokenType.Get()
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionDTO) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenType.Get(), o.TokenType.IsSet()
}

// HasTokenType returns a boolean if a field has been set.
func (o *OnChainTransactionDTO) HasTokenType() bool {
	if o != nil && o.TokenType.IsSet() {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given NullableString and assigns it to the TokenType field.
func (o *OnChainTransactionDTO) SetTokenType(v string) {
	o.TokenType.Set(&v)
}
// SetTokenTypeNil sets the value for TokenType to be an explicit nil
func (o *OnChainTransactionDTO) SetTokenTypeNil() {
	o.TokenType.Set(nil)
}

// UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
func (o *OnChainTransactionDTO) UnsetTokenType() {
	o.TokenType.Unset()
}

func (o OnChainTransactionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainTransactionDTO) ToMap() (map[string]interface{}, error) {
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
	if o.FromAddress.IsSet() {
		toSerialize["from_address"] = o.FromAddress.Get()
	}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
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
	if o.TokenType.IsSet() {
		toSerialize["token_type"] = o.TokenType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnChainTransactionDTO) UnmarshalJSON(data []byte) (err error) {
	varOnChainTransactionDTO := _OnChainTransactionDTO{}

	err = json.Unmarshal(data, &varOnChainTransactionDTO)

	if err != nil {
		return err
	}

	*o = OnChainTransactionDTO(varOnChainTransactionDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset_info")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "description")
		delete(additionalProperties, "from_address")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "status")
		delete(additionalProperties, "to_address")
		delete(additionalProperties, "token_contract")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnChainTransactionDTO struct {
	value *OnChainTransactionDTO
	isSet bool
}

func (v NullableOnChainTransactionDTO) Get() *OnChainTransactionDTO {
	return v.value
}

func (v *NullableOnChainTransactionDTO) Set(val *OnChainTransactionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainTransactionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainTransactionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainTransactionDTO(val *OnChainTransactionDTO) *NullableOnChainTransactionDTO {
	return &NullableOnChainTransactionDTO{value: val, isSet: true}
}

func (v NullableOnChainTransactionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainTransactionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


