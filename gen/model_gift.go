
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Gift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gift{}

// Gift struct for Gift
type Gift struct {
	Currency NullableString `json:"currency,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	IconThumbnail NullableString `json:"icon_thumbnail,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Gift Gift

// NewGift instantiates a new Gift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGift() *Gift {
	this := Gift{}
	return &this
}

// NewGiftWithDefaults instantiates a new Gift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftWithDefaults() *Gift {
	this := Gift{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Gift) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Gift) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Gift) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Gift) UnsetCurrency() {
	o.Currency.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *Gift) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *Gift) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *Gift) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *Gift) UnsetIcon() {
	o.Icon.Unset()
}

// GetIconThumbnail returns the IconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetIconThumbnail() string {
	if o == nil || IsNil(o.IconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.IconThumbnail.Get()
}

// GetIconThumbnailOk returns a tuple with the IconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconThumbnail.Get(), o.IconThumbnail.IsSet()
}

// HasIconThumbnail returns a boolean if a field has been set.
func (o *Gift) HasIconThumbnail() bool {
	if o != nil && o.IconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetIconThumbnail gets a reference to the given NullableString and assigns it to the IconThumbnail field.
func (o *Gift) SetIconThumbnail(v string) {
	o.IconThumbnail.Set(&v)
}
// SetIconThumbnailNil sets the value for IconThumbnail to be an explicit nil
func (o *Gift) SetIconThumbnailNil() {
	o.IconThumbnail.Set(nil)
}

// UnsetIconThumbnail ensures that no value is present for IconThumbnail, not even an explicit nil
func (o *Gift) UnsetIconThumbnail() {
	o.IconThumbnail.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Gift) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Gift) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Gift) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Gift) UnsetId() {
	o.Id.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *Gift) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *Gift) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *Gift) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *Gift) UnsetPrice() {
	o.Price.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gift) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gift) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Gift) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Gift) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Gift) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Gift) UnsetTitle() {
	o.Title.Unset()
}

func (o Gift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.IconThumbnail.IsSet() {
		toSerialize["icon_thumbnail"] = o.IconThumbnail.Get()
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

func (o *Gift) UnmarshalJSON(data []byte) (err error) {
	varGift := _Gift{}

	err = json.Unmarshal(data, &varGift)

	if err != nil {
		return err
	}

	*o = Gift(varGift)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "icon_thumbnail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGift struct {
	value *Gift
	isSet bool
}

func (v NullableGift) Get() *Gift {
	return v.value
}

func (v *NullableGift) Set(val *Gift) {
	v.value = val
	v.isSet = true
}

func (v NullableGift) IsSet() bool {
	return v.isSet
}

func (v *NullableGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGift(val *Gift) *NullableGift {
	return &NullableGift{value: val, isSet: true}
}

func (v NullableGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


