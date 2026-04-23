
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalGachaHistoryResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalGachaHistoryResponseItem{}

// PalGachaHistoryResponseItem struct for PalGachaHistoryResponseItem
type PalGachaHistoryResponseItem struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	GachaType NullableString `json:"gacha_type,omitempty"`
	Id NullableString `json:"id,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	PalState NullableString `json:"pal_state,omitempty"`
	PalTokenId NullableString `json:"pal_token_id,omitempty"`
	Price NullableInt32 `json:"price,omitempty"`
	Rarity NullableString `json:"rarity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalGachaHistoryResponseItem PalGachaHistoryResponseItem

// NewPalGachaHistoryResponseItem instantiates a new PalGachaHistoryResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalGachaHistoryResponseItem() *PalGachaHistoryResponseItem {
	this := PalGachaHistoryResponseItem{}
	return &this
}

// NewPalGachaHistoryResponseItemWithDefaults instantiates a new PalGachaHistoryResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalGachaHistoryResponseItemWithDefaults() *PalGachaHistoryResponseItem {
	this := PalGachaHistoryResponseItem{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *PalGachaHistoryResponseItem) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGachaType returns the GachaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetGachaType() string {
	if o == nil || IsNil(o.GachaType.Get()) {
		var ret string
		return ret
	}
	return *o.GachaType.Get()
}

// GetGachaTypeOk returns a tuple with the GachaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetGachaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GachaType.Get(), o.GachaType.IsSet()
}

// HasGachaType returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasGachaType() bool {
	if o != nil && o.GachaType.IsSet() {
		return true
	}

	return false
}

// SetGachaType gets a reference to the given NullableString and assigns it to the GachaType field.
func (o *PalGachaHistoryResponseItem) SetGachaType(v string) {
	o.GachaType.Set(&v)
}
// SetGachaTypeNil sets the value for GachaType to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetGachaTypeNil() {
	o.GachaType.Set(nil)
}

// UnsetGachaType ensures that no value is present for GachaType, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetGachaType() {
	o.GachaType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PalGachaHistoryResponseItem) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetId() {
	o.Id.Unset()
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *PalGachaHistoryResponseItem) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *PalGachaHistoryResponseItem) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetPalName() {
	o.PalName.Unset()
}

// GetPalState returns the PalState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetPalState() string {
	if o == nil || IsNil(o.PalState.Get()) {
		var ret string
		return ret
	}
	return *o.PalState.Get()
}

// GetPalStateOk returns a tuple with the PalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetPalStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalState.Get(), o.PalState.IsSet()
}

// HasPalState returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasPalState() bool {
	if o != nil && o.PalState.IsSet() {
		return true
	}

	return false
}

// SetPalState gets a reference to the given NullableString and assigns it to the PalState field.
func (o *PalGachaHistoryResponseItem) SetPalState(v string) {
	o.PalState.Set(&v)
}
// SetPalStateNil sets the value for PalState to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetPalStateNil() {
	o.PalState.Set(nil)
}

// UnsetPalState ensures that no value is present for PalState, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetPalState() {
	o.PalState.Unset()
}

// GetPalTokenId returns the PalTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetPalTokenId() string {
	if o == nil || IsNil(o.PalTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.PalTokenId.Get()
}

// GetPalTokenIdOk returns a tuple with the PalTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetPalTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalTokenId.Get(), o.PalTokenId.IsSet()
}

// HasPalTokenId returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasPalTokenId() bool {
	if o != nil && o.PalTokenId.IsSet() {
		return true
	}

	return false
}

// SetPalTokenId gets a reference to the given NullableString and assigns it to the PalTokenId field.
func (o *PalGachaHistoryResponseItem) SetPalTokenId(v string) {
	o.PalTokenId.Set(&v)
}
// SetPalTokenIdNil sets the value for PalTokenId to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetPalTokenIdNil() {
	o.PalTokenId.Set(nil)
}

// UnsetPalTokenId ensures that no value is present for PalTokenId, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetPalTokenId() {
	o.PalTokenId.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetPrice() int32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret int32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt32 and assigns it to the Price field.
func (o *PalGachaHistoryResponseItem) SetPrice(v int32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetPrice() {
	o.Price.Unset()
}

// GetRarity returns the Rarity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponseItem) GetRarity() string {
	if o == nil || IsNil(o.Rarity.Get()) {
		var ret string
		return ret
	}
	return *o.Rarity.Get()
}

// GetRarityOk returns a tuple with the Rarity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponseItem) GetRarityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rarity.Get(), o.Rarity.IsSet()
}

// HasRarity returns a boolean if a field has been set.
func (o *PalGachaHistoryResponseItem) HasRarity() bool {
	if o != nil && o.Rarity.IsSet() {
		return true
	}

	return false
}

// SetRarity gets a reference to the given NullableString and assigns it to the Rarity field.
func (o *PalGachaHistoryResponseItem) SetRarity(v string) {
	o.Rarity.Set(&v)
}
// SetRarityNil sets the value for Rarity to be an explicit nil
func (o *PalGachaHistoryResponseItem) SetRarityNil() {
	o.Rarity.Set(nil)
}

// UnsetRarity ensures that no value is present for Rarity, not even an explicit nil
func (o *PalGachaHistoryResponseItem) UnsetRarity() {
	o.Rarity.Unset()
}

func (o PalGachaHistoryResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalGachaHistoryResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.GachaType.IsSet() {
		toSerialize["gacha_type"] = o.GachaType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}
	if o.PalState.IsSet() {
		toSerialize["pal_state"] = o.PalState.Get()
	}
	if o.PalTokenId.IsSet() {
		toSerialize["pal_token_id"] = o.PalTokenId.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Rarity.IsSet() {
		toSerialize["rarity"] = o.Rarity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalGachaHistoryResponseItem) UnmarshalJSON(data []byte) (err error) {
	varPalGachaHistoryResponseItem := _PalGachaHistoryResponseItem{}

	err = json.Unmarshal(data, &varPalGachaHistoryResponseItem)

	if err != nil {
		return err
	}

	*o = PalGachaHistoryResponseItem(varPalGachaHistoryResponseItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "gacha_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		delete(additionalProperties, "pal_state")
		delete(additionalProperties, "pal_token_id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "rarity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalGachaHistoryResponseItem struct {
	value *PalGachaHistoryResponseItem
	isSet bool
}

func (v NullablePalGachaHistoryResponseItem) Get() *PalGachaHistoryResponseItem {
	return v.value
}

func (v *NullablePalGachaHistoryResponseItem) Set(val *PalGachaHistoryResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGachaHistoryResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGachaHistoryResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGachaHistoryResponseItem(val *PalGachaHistoryResponseItem) *NullablePalGachaHistoryResponseItem {
	return &NullablePalGachaHistoryResponseItem{value: val, isSet: true}
}

func (v NullablePalGachaHistoryResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGachaHistoryResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


