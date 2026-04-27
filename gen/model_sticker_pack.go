
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the StickerPack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StickerPack{}

// StickerPack struct for StickerPack
type StickerPack struct {
	Cover NullableString `json:"cover,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Order NullableInt32 `json:"order,omitempty"`
	Stickers map[string]interface{} `json:"stickers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StickerPack StickerPack

// NewStickerPack instantiates a new StickerPack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStickerPack() *StickerPack {
	this := StickerPack{}
	return &this
}

// NewStickerPackWithDefaults instantiates a new StickerPack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerPackWithDefaults() *StickerPack {
	this := StickerPack{}
	return &this
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *StickerPack) HasCover() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *StickerPack) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *StickerPack) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *StickerPack) UnsetCover() {
	o.Cover.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StickerPack) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StickerPack) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StickerPack) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StickerPack) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *StickerPack) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *StickerPack) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *StickerPack) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *StickerPack) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StickerPack) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StickerPack) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *StickerPack) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StickerPack) UnsetName() {
	o.Name.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetOrder() int32 {
	if o == nil || IsNil(o.Order.Get()) {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *StickerPack) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *StickerPack) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *StickerPack) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *StickerPack) UnsetOrder() {
	o.Order.Unset()
}

// GetStickers returns the Stickers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPack) GetStickers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPack) GetStickersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Stickers) {
		return map[string]interface{}{}, false
	}
	return o.Stickers, true
}

// HasStickers returns a boolean if a field has been set.
func (o *StickerPack) HasStickers() bool {
	if o != nil && !IsNil(o.Stickers) {
		return true
	}

	return false
}

// SetStickers gets a reference to the given map[string]interface{} and assigns it to the Stickers field.
func (o *StickerPack) SetStickers(v map[string]interface{}) {
	o.Stickers = v
}

func (o StickerPack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StickerPack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cover.IsSet() {
		toSerialize["cover"] = o.Cover.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.Stickers != nil {
		toSerialize["stickers"] = o.Stickers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StickerPack) UnmarshalJSON(data []byte) (err error) {
	varStickerPack := _StickerPack{}

	err = json.Unmarshal(data, &varStickerPack)

	if err != nil {
		return err
	}

	*o = StickerPack(varStickerPack)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cover")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "order")
		delete(additionalProperties, "stickers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStickerPack struct {
	value *StickerPack
	isSet bool
}

func (v NullableStickerPack) Get() *StickerPack {
	return v.value
}

func (v *NullableStickerPack) Set(val *StickerPack) {
	v.value = val
	v.isSet = true
}

func (v NullableStickerPack) IsSet() bool {
	return v.isSet
}

func (v *NullableStickerPack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStickerPack(val *StickerPack) *NullableStickerPack {
	return &NullableStickerPack{value: val, isSet: true}
}

func (v NullableStickerPack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStickerPack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


