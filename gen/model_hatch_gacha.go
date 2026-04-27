
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the HatchGacha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HatchGacha{}

// HatchGacha struct for HatchGacha
type HatchGacha struct {
	EmplCost NullableFloat64 `json:"empl_cost,omitempty"`
	GachaType map[string]interface{} `json:"gacha_type,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	PalRank map[string]interface{} `json:"pal_rank,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HatchGacha HatchGacha

// NewHatchGacha instantiates a new HatchGacha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHatchGacha() *HatchGacha {
	this := HatchGacha{}
	return &this
}

// NewHatchGachaWithDefaults instantiates a new HatchGacha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHatchGachaWithDefaults() *HatchGacha {
	this := HatchGacha{}
	return &this
}

// GetEmplCost returns the EmplCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetEmplCost() float64 {
	if o == nil || IsNil(o.EmplCost.Get()) {
		var ret float64
		return ret
	}
	return *o.EmplCost.Get()
}

// GetEmplCostOk returns a tuple with the EmplCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetEmplCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplCost.Get(), o.EmplCost.IsSet()
}

// HasEmplCost returns a boolean if a field has been set.
func (o *HatchGacha) HasEmplCost() bool {
	if o != nil && o.EmplCost.IsSet() {
		return true
	}

	return false
}

// SetEmplCost gets a reference to the given NullableFloat64 and assigns it to the EmplCost field.
func (o *HatchGacha) SetEmplCost(v float64) {
	o.EmplCost.Set(&v)
}
// SetEmplCostNil sets the value for EmplCost to be an explicit nil
func (o *HatchGacha) SetEmplCostNil() {
	o.EmplCost.Set(nil)
}

// UnsetEmplCost ensures that no value is present for EmplCost, not even an explicit nil
func (o *HatchGacha) UnsetEmplCost() {
	o.EmplCost.Unset()
}

// GetGachaType returns the GachaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetGachaType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GachaType
}

// GetGachaTypeOk returns a tuple with the GachaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetGachaTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GachaType) {
		return map[string]interface{}{}, false
	}
	return o.GachaType, true
}

// HasGachaType returns a boolean if a field has been set.
func (o *HatchGacha) HasGachaType() bool {
	if o != nil && !IsNil(o.GachaType) {
		return true
	}

	return false
}

// SetGachaType gets a reference to the given map[string]interface{} and assigns it to the GachaType field.
func (o *HatchGacha) SetGachaType(v map[string]interface{}) {
	o.GachaType = v
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *HatchGacha) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *HatchGacha) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *HatchGacha) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *HatchGacha) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *HatchGacha) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *HatchGacha) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *HatchGacha) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *HatchGacha) UnsetPalName() {
	o.PalName.Unset()
}

// GetPalRank returns the PalRank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetPalRank() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PalRank
}

// GetPalRankOk returns a tuple with the PalRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetPalRankOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PalRank) {
		return map[string]interface{}{}, false
	}
	return o.PalRank, true
}

// HasPalRank returns a boolean if a field has been set.
func (o *HatchGacha) HasPalRank() bool {
	if o != nil && !IsNil(o.PalRank) {
		return true
	}

	return false
}

// SetPalRank gets a reference to the given map[string]interface{} and assigns it to the PalRank field.
func (o *HatchGacha) SetPalRank(v map[string]interface{}) {
	o.PalRank = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HatchGacha) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HatchGacha) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *HatchGacha) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *HatchGacha) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *HatchGacha) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *HatchGacha) UnsetTokenId() {
	o.TokenId.Unset()
}

func (o HatchGacha) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HatchGacha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmplCost.IsSet() {
		toSerialize["empl_cost"] = o.EmplCost.Get()
	}
	if o.GachaType != nil {
		toSerialize["gacha_type"] = o.GachaType
	}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}
	if o.PalRank != nil {
		toSerialize["pal_rank"] = o.PalRank
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HatchGacha) UnmarshalJSON(data []byte) (err error) {
	varHatchGacha := _HatchGacha{}

	err = json.Unmarshal(data, &varHatchGacha)

	if err != nil {
		return err
	}

	*o = HatchGacha(varHatchGacha)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl_cost")
		delete(additionalProperties, "gacha_type")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		delete(additionalProperties, "pal_rank")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHatchGacha struct {
	value *HatchGacha
	isSet bool
}

func (v NullableHatchGacha) Get() *HatchGacha {
	return v.value
}

func (v *NullableHatchGacha) Set(val *HatchGacha) {
	v.value = val
	v.isSet = true
}

func (v NullableHatchGacha) IsSet() bool {
	return v.isSet
}

func (v *NullableHatchGacha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHatchGacha(val *HatchGacha) *NullableHatchGacha {
	return &NullableHatchGacha{value: val, isSet: true}
}

func (v NullableHatchGacha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHatchGacha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


