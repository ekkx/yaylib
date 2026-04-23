
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TransactionGiftReceivedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionGiftReceivedItem{}

// TransactionGiftReceivedItem struct for TransactionGiftReceivedItem
type TransactionGiftReceivedItem struct {
	Currency NullableString `json:"currency,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Price NullableInt32 `json:"price,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionGiftReceivedItem TransactionGiftReceivedItem

// NewTransactionGiftReceivedItem instantiates a new TransactionGiftReceivedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionGiftReceivedItem() *TransactionGiftReceivedItem {
	this := TransactionGiftReceivedItem{}
	return &this
}

// NewTransactionGiftReceivedItemWithDefaults instantiates a new TransactionGiftReceivedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionGiftReceivedItemWithDefaults() *TransactionGiftReceivedItem {
	this := TransactionGiftReceivedItem{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceivedItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceivedItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransactionGiftReceivedItem) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *TransactionGiftReceivedItem) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *TransactionGiftReceivedItem) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *TransactionGiftReceivedItem) UnsetCurrency() {
	o.Currency.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceivedItem) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceivedItem) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TransactionGiftReceivedItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *TransactionGiftReceivedItem) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TransactionGiftReceivedItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TransactionGiftReceivedItem) UnsetId() {
	o.Id.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceivedItem) GetPrice() int32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret int32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceivedItem) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *TransactionGiftReceivedItem) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt32 and assigns it to the Price field.
func (o *TransactionGiftReceivedItem) SetPrice(v int32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *TransactionGiftReceivedItem) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *TransactionGiftReceivedItem) UnsetPrice() {
	o.Price.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceivedItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceivedItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TransactionGiftReceivedItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TransactionGiftReceivedItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TransactionGiftReceivedItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TransactionGiftReceivedItem) UnsetTitle() {
	o.Title.Unset()
}

func (o TransactionGiftReceivedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionGiftReceivedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionGiftReceivedItem) UnmarshalJSON(data []byte) (err error) {
	varTransactionGiftReceivedItem := _TransactionGiftReceivedItem{}

	err = json.Unmarshal(data, &varTransactionGiftReceivedItem)

	if err != nil {
		return err
	}

	*o = TransactionGiftReceivedItem(varTransactionGiftReceivedItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionGiftReceivedItem struct {
	value *TransactionGiftReceivedItem
	isSet bool
}

func (v NullableTransactionGiftReceivedItem) Get() *TransactionGiftReceivedItem {
	return v.value
}

func (v *NullableTransactionGiftReceivedItem) Set(val *TransactionGiftReceivedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionGiftReceivedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionGiftReceivedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionGiftReceivedItem(val *TransactionGiftReceivedItem) *NullableTransactionGiftReceivedItem {
	return &NullableTransactionGiftReceivedItem{value: val, isSet: true}
}

func (v NullableTransactionGiftReceivedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionGiftReceivedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


