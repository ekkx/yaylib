
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftSlugItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftSlugItem{}

// GiftSlugItem struct for GiftSlugItem
type GiftSlugItem struct {
	Icon NullableString `json:"icon,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Quantity NullableInt32 `json:"quantity,omitempty"`
	Slug NullableString `json:"slug,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftSlugItem GiftSlugItem

// NewGiftSlugItem instantiates a new GiftSlugItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftSlugItem() *GiftSlugItem {
	this := GiftSlugItem{}
	return &this
}

// NewGiftSlugItemWithDefaults instantiates a new GiftSlugItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftSlugItemWithDefaults() *GiftSlugItem {
	this := GiftSlugItem{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSlugItem) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSlugItem) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GiftSlugItem) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GiftSlugItem) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *GiftSlugItem) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GiftSlugItem) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSlugItem) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSlugItem) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftSlugItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *GiftSlugItem) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftSlugItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftSlugItem) UnsetId() {
	o.Id.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSlugItem) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSlugItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *GiftSlugItem) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *GiftSlugItem) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *GiftSlugItem) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *GiftSlugItem) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSlugItem) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSlugItem) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *GiftSlugItem) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *GiftSlugItem) SetSlug(v string) {
	o.Slug.Set(&v)
}
// SetSlugNil sets the value for Slug to be an explicit nil
func (o *GiftSlugItem) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *GiftSlugItem) UnsetSlug() {
	o.Slug.Unset()
}

func (o GiftSlugItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftSlugItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftSlugItem) UnmarshalJSON(data []byte) (err error) {
	varGiftSlugItem := _GiftSlugItem{}

	err = json.Unmarshal(data, &varGiftSlugItem)

	if err != nil {
		return err
	}

	*o = GiftSlugItem(varGiftSlugItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftSlugItem struct {
	value *GiftSlugItem
	isSet bool
}

func (v NullableGiftSlugItem) Get() *GiftSlugItem {
	return v.value
}

func (v *NullableGiftSlugItem) Set(val *GiftSlugItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftSlugItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftSlugItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftSlugItem(val *GiftSlugItem) *NullableGiftSlugItem {
	return &NullableGiftSlugItem{value: val, isSet: true}
}

func (v NullableGiftSlugItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftSlugItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


