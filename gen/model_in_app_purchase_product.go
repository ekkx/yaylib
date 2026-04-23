
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InAppPurchaseProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseProduct{}

// InAppPurchaseProduct struct for InAppPurchaseProduct
type InAppPurchaseProduct struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	Bonus NullableFloat64 `json:"bonus,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Purchasable NullableBool `json:"purchasable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InAppPurchaseProduct InAppPurchaseProduct

// NewInAppPurchaseProduct instantiates a new InAppPurchaseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseProduct() *InAppPurchaseProduct {
	this := InAppPurchaseProduct{}
	return &this
}

// NewInAppPurchaseProductWithDefaults instantiates a new InAppPurchaseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseProductWithDefaults() *InAppPurchaseProduct {
	this := InAppPurchaseProduct{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProduct) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProduct) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *InAppPurchaseProduct) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *InAppPurchaseProduct) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *InAppPurchaseProduct) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *InAppPurchaseProduct) UnsetAmount() {
	o.Amount.Unset()
}

// GetBonus returns the Bonus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProduct) GetBonus() float64 {
	if o == nil || IsNil(o.Bonus.Get()) {
		var ret float64
		return ret
	}
	return *o.Bonus.Get()
}

// GetBonusOk returns a tuple with the Bonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProduct) GetBonusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bonus.Get(), o.Bonus.IsSet()
}

// HasBonus returns a boolean if a field has been set.
func (o *InAppPurchaseProduct) HasBonus() bool {
	if o != nil && o.Bonus.IsSet() {
		return true
	}

	return false
}

// SetBonus gets a reference to the given NullableFloat64 and assigns it to the Bonus field.
func (o *InAppPurchaseProduct) SetBonus(v float64) {
	o.Bonus.Set(&v)
}
// SetBonusNil sets the value for Bonus to be an explicit nil
func (o *InAppPurchaseProduct) SetBonusNil() {
	o.Bonus.Set(nil)
}

// UnsetBonus ensures that no value is present for Bonus, not even an explicit nil
func (o *InAppPurchaseProduct) UnsetBonus() {
	o.Bonus.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProduct) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProduct) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InAppPurchaseProduct) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *InAppPurchaseProduct) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InAppPurchaseProduct) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InAppPurchaseProduct) UnsetId() {
	o.Id.Unset()
}

// GetPurchasable returns the Purchasable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InAppPurchaseProduct) GetPurchasable() bool {
	if o == nil || IsNil(o.Purchasable.Get()) {
		var ret bool
		return ret
	}
	return *o.Purchasable.Get()
}

// GetPurchasableOk returns a tuple with the Purchasable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InAppPurchaseProduct) GetPurchasableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purchasable.Get(), o.Purchasable.IsSet()
}

// HasPurchasable returns a boolean if a field has been set.
func (o *InAppPurchaseProduct) HasPurchasable() bool {
	if o != nil && o.Purchasable.IsSet() {
		return true
	}

	return false
}

// SetPurchasable gets a reference to the given NullableBool and assigns it to the Purchasable field.
func (o *InAppPurchaseProduct) SetPurchasable(v bool) {
	o.Purchasable.Set(&v)
}
// SetPurchasableNil sets the value for Purchasable to be an explicit nil
func (o *InAppPurchaseProduct) SetPurchasableNil() {
	o.Purchasable.Set(nil)
}

// UnsetPurchasable ensures that no value is present for Purchasable, not even an explicit nil
func (o *InAppPurchaseProduct) UnsetPurchasable() {
	o.Purchasable.Unset()
}

func (o InAppPurchaseProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Bonus.IsSet() {
		toSerialize["bonus"] = o.Bonus.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Purchasable.IsSet() {
		toSerialize["purchasable"] = o.Purchasable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InAppPurchaseProduct) UnmarshalJSON(data []byte) (err error) {
	varInAppPurchaseProduct := _InAppPurchaseProduct{}

	err = json.Unmarshal(data, &varInAppPurchaseProduct)

	if err != nil {
		return err
	}

	*o = InAppPurchaseProduct(varInAppPurchaseProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "bonus")
		delete(additionalProperties, "id")
		delete(additionalProperties, "purchasable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInAppPurchaseProduct struct {
	value *InAppPurchaseProduct
	isSet bool
}

func (v NullableInAppPurchaseProduct) Get() *InAppPurchaseProduct {
	return v.value
}

func (v *NullableInAppPurchaseProduct) Set(val *InAppPurchaseProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseProduct(val *InAppPurchaseProduct) *NullableInAppPurchaseProduct {
	return &NullableInAppPurchaseProduct{value: val, isSet: true}
}

func (v NullableInAppPurchaseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


