
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Promotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Promotion{}

// Promotion struct for Promotion
type Promotion struct {
	Id NullableInt64 `json:"id,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	Order NullableInt32 `json:"order,omitempty"`
	PromotionUrl NullableString `json:"promotion_url,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Promotion Promotion

// NewPromotion instantiates a new Promotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion() *Promotion {
	this := Promotion{}
	return &this
}

// NewPromotionWithDefaults instantiates a new Promotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionWithDefaults() *Promotion {
	this := Promotion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Promotion) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Promotion) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Promotion) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Promotion) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Promotion) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Promotion) UnsetId() {
	o.Id.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Promotion) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Promotion) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Promotion) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Promotion) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Promotion) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Promotion) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Promotion) GetOrder() int32 {
	if o == nil || IsNil(o.Order.Get()) {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Promotion) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *Promotion) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *Promotion) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *Promotion) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *Promotion) UnsetOrder() {
	o.Order.Unset()
}

// GetPromotionUrl returns the PromotionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Promotion) GetPromotionUrl() string {
	if o == nil || IsNil(o.PromotionUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PromotionUrl.Get()
}

// GetPromotionUrlOk returns a tuple with the PromotionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Promotion) GetPromotionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromotionUrl.Get(), o.PromotionUrl.IsSet()
}

// HasPromotionUrl returns a boolean if a field has been set.
func (o *Promotion) HasPromotionUrl() bool {
	if o != nil && o.PromotionUrl.IsSet() {
		return true
	}

	return false
}

// SetPromotionUrl gets a reference to the given NullableString and assigns it to the PromotionUrl field.
func (o *Promotion) SetPromotionUrl(v string) {
	o.PromotionUrl.Set(&v)
}
// SetPromotionUrlNil sets the value for PromotionUrl to be an explicit nil
func (o *Promotion) SetPromotionUrlNil() {
	o.PromotionUrl.Set(nil)
}

// UnsetPromotionUrl ensures that no value is present for PromotionUrl, not even an explicit nil
func (o *Promotion) UnsetPromotionUrl() {
	o.PromotionUrl.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Promotion) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Promotion) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Promotion) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Promotion) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Promotion) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Promotion) UnsetTitle() {
	o.Title.Unset()
}

func (o Promotion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Promotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.PromotionUrl.IsSet() {
		toSerialize["promotion_url"] = o.PromotionUrl.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Promotion) UnmarshalJSON(data []byte) (err error) {
	varPromotion := _Promotion{}

	err = json.Unmarshal(data, &varPromotion)

	if err != nil {
		return err
	}

	*o = Promotion(varPromotion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "order")
		delete(additionalProperties, "promotion_url")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePromotion struct {
	value *Promotion
	isSet bool
}

func (v NullablePromotion) Get() *Promotion {
	return v.value
}

func (v *NullablePromotion) Set(val *Promotion) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion(val *Promotion) *NullablePromotion {
	return &NullablePromotion{value: val, isSet: true}
}

func (v NullablePromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


