
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelInAppPurchaseProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInAppPurchaseProduct{}

// ModelInAppPurchaseProduct struct for ModelInAppPurchaseProduct
type ModelInAppPurchaseProduct struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	Bonus NullableFloat64 `json:"bonus,omitempty"`
	CurrencyCode NullableString `json:"currency_code,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	Purchasable NullableBool `json:"purchasable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelInAppPurchaseProduct ModelInAppPurchaseProduct

// NewModelInAppPurchaseProduct instantiates a new ModelInAppPurchaseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInAppPurchaseProduct() *ModelInAppPurchaseProduct {
	this := ModelInAppPurchaseProduct{}
	return &this
}

// NewModelInAppPurchaseProductWithDefaults instantiates a new ModelInAppPurchaseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInAppPurchaseProductWithDefaults() *ModelInAppPurchaseProduct {
	this := ModelInAppPurchaseProduct{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *ModelInAppPurchaseProduct) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetAmount() {
	o.Amount.Unset()
}

// GetBonus returns the Bonus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetBonus() float64 {
	if o == nil || IsNil(o.Bonus.Get()) {
		var ret float64
		return ret
	}
	return *o.Bonus.Get()
}

// GetBonusOk returns a tuple with the Bonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetBonusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bonus.Get(), o.Bonus.IsSet()
}

// HasBonus returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasBonus() bool {
	if o != nil && o.Bonus.IsSet() {
		return true
	}

	return false
}

// SetBonus gets a reference to the given NullableFloat64 and assigns it to the Bonus field.
func (o *ModelInAppPurchaseProduct) SetBonus(v float64) {
	o.Bonus.Set(&v)
}
// SetBonusNil sets the value for Bonus to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetBonusNil() {
	o.Bonus.Set(nil)
}

// UnsetBonus ensures that no value is present for Bonus, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetBonus() {
	o.Bonus.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *ModelInAppPurchaseProduct) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}
// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ModelInAppPurchaseProduct) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetId() {
	o.Id.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *ModelInAppPurchaseProduct) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetPrice() {
	o.Price.Unset()
}

// GetPurchasable returns the Purchasable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInAppPurchaseProduct) GetPurchasable() bool {
	if o == nil || IsNil(o.Purchasable.Get()) {
		var ret bool
		return ret
	}
	return *o.Purchasable.Get()
}

// GetPurchasableOk returns a tuple with the Purchasable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInAppPurchaseProduct) GetPurchasableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purchasable.Get(), o.Purchasable.IsSet()
}

// HasPurchasable returns a boolean if a field has been set.
func (o *ModelInAppPurchaseProduct) HasPurchasable() bool {
	if o != nil && o.Purchasable.IsSet() {
		return true
	}

	return false
}

// SetPurchasable gets a reference to the given NullableBool and assigns it to the Purchasable field.
func (o *ModelInAppPurchaseProduct) SetPurchasable(v bool) {
	o.Purchasable.Set(&v)
}
// SetPurchasableNil sets the value for Purchasable to be an explicit nil
func (o *ModelInAppPurchaseProduct) SetPurchasableNil() {
	o.Purchasable.Set(nil)
}

// UnsetPurchasable ensures that no value is present for Purchasable, not even an explicit nil
func (o *ModelInAppPurchaseProduct) UnsetPurchasable() {
	o.Purchasable.Unset()
}

func (o ModelInAppPurchaseProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInAppPurchaseProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Bonus.IsSet() {
		toSerialize["bonus"] = o.Bonus.Get()
	}
	if o.CurrencyCode.IsSet() {
		toSerialize["currency_code"] = o.CurrencyCode.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Purchasable.IsSet() {
		toSerialize["purchasable"] = o.Purchasable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelInAppPurchaseProduct) UnmarshalJSON(data []byte) (err error) {
	varModelInAppPurchaseProduct := _ModelInAppPurchaseProduct{}

	err = json.Unmarshal(data, &varModelInAppPurchaseProduct)

	if err != nil {
		return err
	}

	*o = ModelInAppPurchaseProduct(varModelInAppPurchaseProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "bonus")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "purchasable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelInAppPurchaseProduct struct {
	value *ModelInAppPurchaseProduct
	isSet bool
}

func (v NullableModelInAppPurchaseProduct) Get() *ModelInAppPurchaseProduct {
	return v.value
}

func (v *NullableModelInAppPurchaseProduct) Set(val *ModelInAppPurchaseProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInAppPurchaseProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInAppPurchaseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInAppPurchaseProduct(val *ModelInAppPurchaseProduct) *NullableModelInAppPurchaseProduct {
	return &NullableModelInAppPurchaseProduct{value: val, isSet: true}
}

func (v NullableModelInAppPurchaseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInAppPurchaseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


