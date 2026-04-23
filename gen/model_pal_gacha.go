
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalGacha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalGacha{}

// PalGacha struct for PalGacha
type PalGacha struct {
	BannerCollectionUrl NullableString `json:"banner_collection_url,omitempty"`
	BannerUrl NullableString `json:"banner_url,omitempty"`
	DetailUrl NullableString `json:"detail_url,omitempty"`
	GachaType NullableString `json:"gacha_type,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Price NullableInt32 `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalGacha PalGacha

// NewPalGacha instantiates a new PalGacha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalGacha() *PalGacha {
	this := PalGacha{}
	return &this
}

// NewPalGachaWithDefaults instantiates a new PalGacha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalGachaWithDefaults() *PalGacha {
	this := PalGacha{}
	return &this
}

// GetBannerCollectionUrl returns the BannerCollectionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetBannerCollectionUrl() string {
	if o == nil || IsNil(o.BannerCollectionUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BannerCollectionUrl.Get()
}

// GetBannerCollectionUrlOk returns a tuple with the BannerCollectionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetBannerCollectionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerCollectionUrl.Get(), o.BannerCollectionUrl.IsSet()
}

// HasBannerCollectionUrl returns a boolean if a field has been set.
func (o *PalGacha) HasBannerCollectionUrl() bool {
	if o != nil && o.BannerCollectionUrl.IsSet() {
		return true
	}

	return false
}

// SetBannerCollectionUrl gets a reference to the given NullableString and assigns it to the BannerCollectionUrl field.
func (o *PalGacha) SetBannerCollectionUrl(v string) {
	o.BannerCollectionUrl.Set(&v)
}
// SetBannerCollectionUrlNil sets the value for BannerCollectionUrl to be an explicit nil
func (o *PalGacha) SetBannerCollectionUrlNil() {
	o.BannerCollectionUrl.Set(nil)
}

// UnsetBannerCollectionUrl ensures that no value is present for BannerCollectionUrl, not even an explicit nil
func (o *PalGacha) UnsetBannerCollectionUrl() {
	o.BannerCollectionUrl.Unset()
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetBannerUrl() string {
	if o == nil || IsNil(o.BannerUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BannerUrl.Get()
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetBannerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerUrl.Get(), o.BannerUrl.IsSet()
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *PalGacha) HasBannerUrl() bool {
	if o != nil && o.BannerUrl.IsSet() {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given NullableString and assigns it to the BannerUrl field.
func (o *PalGacha) SetBannerUrl(v string) {
	o.BannerUrl.Set(&v)
}
// SetBannerUrlNil sets the value for BannerUrl to be an explicit nil
func (o *PalGacha) SetBannerUrlNil() {
	o.BannerUrl.Set(nil)
}

// UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
func (o *PalGacha) UnsetBannerUrl() {
	o.BannerUrl.Unset()
}

// GetDetailUrl returns the DetailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetDetailUrl() string {
	if o == nil || IsNil(o.DetailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DetailUrl.Get()
}

// GetDetailUrlOk returns a tuple with the DetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetDetailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetailUrl.Get(), o.DetailUrl.IsSet()
}

// HasDetailUrl returns a boolean if a field has been set.
func (o *PalGacha) HasDetailUrl() bool {
	if o != nil && o.DetailUrl.IsSet() {
		return true
	}

	return false
}

// SetDetailUrl gets a reference to the given NullableString and assigns it to the DetailUrl field.
func (o *PalGacha) SetDetailUrl(v string) {
	o.DetailUrl.Set(&v)
}
// SetDetailUrlNil sets the value for DetailUrl to be an explicit nil
func (o *PalGacha) SetDetailUrlNil() {
	o.DetailUrl.Set(nil)
}

// UnsetDetailUrl ensures that no value is present for DetailUrl, not even an explicit nil
func (o *PalGacha) UnsetDetailUrl() {
	o.DetailUrl.Unset()
}

// GetGachaType returns the GachaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetGachaType() string {
	if o == nil || IsNil(o.GachaType.Get()) {
		var ret string
		return ret
	}
	return *o.GachaType.Get()
}

// GetGachaTypeOk returns a tuple with the GachaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetGachaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GachaType.Get(), o.GachaType.IsSet()
}

// HasGachaType returns a boolean if a field has been set.
func (o *PalGacha) HasGachaType() bool {
	if o != nil && o.GachaType.IsSet() {
		return true
	}

	return false
}

// SetGachaType gets a reference to the given NullableString and assigns it to the GachaType field.
func (o *PalGacha) SetGachaType(v string) {
	o.GachaType.Set(&v)
}
// SetGachaTypeNil sets the value for GachaType to be an explicit nil
func (o *PalGacha) SetGachaTypeNil() {
	o.GachaType.Set(nil)
}

// UnsetGachaType ensures that no value is present for GachaType, not even an explicit nil
func (o *PalGacha) UnsetGachaType() {
	o.GachaType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PalGacha) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PalGacha) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PalGacha) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PalGacha) UnsetId() {
	o.Id.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGacha) GetPrice() int32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret int32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGacha) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *PalGacha) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt32 and assigns it to the Price field.
func (o *PalGacha) SetPrice(v int32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *PalGacha) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *PalGacha) UnsetPrice() {
	o.Price.Unset()
}

func (o PalGacha) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalGacha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BannerCollectionUrl.IsSet() {
		toSerialize["banner_collection_url"] = o.BannerCollectionUrl.Get()
	}
	if o.BannerUrl.IsSet() {
		toSerialize["banner_url"] = o.BannerUrl.Get()
	}
	if o.DetailUrl.IsSet() {
		toSerialize["detail_url"] = o.DetailUrl.Get()
	}
	if o.GachaType.IsSet() {
		toSerialize["gacha_type"] = o.GachaType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalGacha) UnmarshalJSON(data []byte) (err error) {
	varPalGacha := _PalGacha{}

	err = json.Unmarshal(data, &varPalGacha)

	if err != nil {
		return err
	}

	*o = PalGacha(varPalGacha)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "banner_collection_url")
		delete(additionalProperties, "banner_url")
		delete(additionalProperties, "detail_url")
		delete(additionalProperties, "gacha_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalGacha struct {
	value *PalGacha
	isSet bool
}

func (v NullablePalGacha) Get() *PalGacha {
	return v.value
}

func (v *NullablePalGacha) Set(val *PalGacha) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGacha) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGacha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGacha(val *PalGacha) *NullablePalGacha {
	return &NullablePalGacha{value: val, isSet: true}
}

func (v NullablePalGacha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGacha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


