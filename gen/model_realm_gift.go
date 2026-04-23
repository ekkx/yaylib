
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmGift{}

// RealmGift struct for RealmGift
type RealmGift struct {
	Currency NullableString `json:"currency,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	IconThumbnail NullableString `json:"icon_thumbnail,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	Slug NullableString `json:"slug,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmGift RealmGift

// NewRealmGift instantiates a new RealmGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmGift() *RealmGift {
	this := RealmGift{}
	return &this
}

// NewRealmGiftWithDefaults instantiates a new RealmGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmGiftWithDefaults() *RealmGift {
	this := RealmGift{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *RealmGift) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *RealmGift) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *RealmGift) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *RealmGift) UnsetCurrency() {
	o.Currency.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *RealmGift) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *RealmGift) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *RealmGift) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *RealmGift) UnsetIcon() {
	o.Icon.Unset()
}

// GetIconThumbnail returns the IconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetIconThumbnail() string {
	if o == nil || IsNil(o.IconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.IconThumbnail.Get()
}

// GetIconThumbnailOk returns a tuple with the IconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconThumbnail.Get(), o.IconThumbnail.IsSet()
}

// HasIconThumbnail returns a boolean if a field has been set.
func (o *RealmGift) HasIconThumbnail() bool {
	if o != nil && o.IconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetIconThumbnail gets a reference to the given NullableString and assigns it to the IconThumbnail field.
func (o *RealmGift) SetIconThumbnail(v string) {
	o.IconThumbnail.Set(&v)
}
// SetIconThumbnailNil sets the value for IconThumbnail to be an explicit nil
func (o *RealmGift) SetIconThumbnailNil() {
	o.IconThumbnail.Set(nil)
}

// UnsetIconThumbnail ensures that no value is present for IconThumbnail, not even an explicit nil
func (o *RealmGift) UnsetIconThumbnail() {
	o.IconThumbnail.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmGift) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmGift) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmGift) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmGift) UnsetId() {
	o.Id.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *RealmGift) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *RealmGift) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *RealmGift) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *RealmGift) UnsetPrice() {
	o.Price.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *RealmGift) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *RealmGift) SetSlug(v string) {
	o.Slug.Set(&v)
}
// SetSlugNil sets the value for Slug to be an explicit nil
func (o *RealmGift) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *RealmGift) UnsetSlug() {
	o.Slug.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGift) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGift) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RealmGift) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RealmGift) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RealmGift) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RealmGift) UnsetTitle() {
	o.Title.Unset()
}

func (o RealmGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmGift) ToMap() (map[string]interface{}, error) {
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
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmGift) UnmarshalJSON(data []byte) (err error) {
	varRealmGift := _RealmGift{}

	err = json.Unmarshal(data, &varRealmGift)

	if err != nil {
		return err
	}

	*o = RealmGift(varRealmGift)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "icon_thumbnail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmGift struct {
	value *RealmGift
	isSet bool
}

func (v NullableRealmGift) Get() *RealmGift {
	return v.value
}

func (v *NullableRealmGift) Set(val *RealmGift) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmGift) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmGift(val *RealmGift) *NullableRealmGift {
	return &NullableRealmGift{value: val, isSet: true}
}

func (v NullableRealmGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


