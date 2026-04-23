
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Withdraw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Withdraw{}

// Withdraw struct for Withdraw
type Withdraw struct {
	AssetType NullableAssetType `json:"asset_type,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Withdraw Withdraw

// NewWithdraw instantiates a new Withdraw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdraw() *Withdraw {
	this := Withdraw{}
	return &this
}

// NewWithdrawWithDefaults instantiates a new Withdraw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawWithDefaults() *Withdraw {
	this := Withdraw{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Withdraw) GetAssetType() AssetType {
	if o == nil || IsNil(o.AssetType.Get()) {
		var ret AssetType
		return ret
	}
	return *o.AssetType.Get()
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Withdraw) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetType.Get(), o.AssetType.IsSet()
}

// HasAssetType returns a boolean if a field has been set.
func (o *Withdraw) HasAssetType() bool {
	if o != nil && o.AssetType.IsSet() {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given NullableAssetType and assigns it to the AssetType field.
func (o *Withdraw) SetAssetType(v AssetType) {
	o.AssetType.Set(&v)
}
// SetAssetTypeNil sets the value for AssetType to be an explicit nil
func (o *Withdraw) SetAssetTypeNil() {
	o.AssetType.Set(nil)
}

// UnsetAssetType ensures that no value is present for AssetType, not even an explicit nil
func (o *Withdraw) UnsetAssetType() {
	o.AssetType.Unset()
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Withdraw) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Withdraw) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *Withdraw) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *Withdraw) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *Withdraw) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *Withdraw) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Withdraw) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Withdraw) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *Withdraw) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *Withdraw) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *Withdraw) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *Withdraw) UnsetPalName() {
	o.PalName.Unset()
}

func (o Withdraw) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Withdraw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType.IsSet() {
		toSerialize["asset_type"] = o.AssetType.Get()
	}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Withdraw) UnmarshalJSON(data []byte) (err error) {
	varWithdraw := _Withdraw{}

	err = json.Unmarshal(data, &varWithdraw)

	if err != nil {
		return err
	}

	*o = Withdraw(varWithdraw)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_type")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWithdraw struct {
	value *Withdraw
	isSet bool
}

func (v NullableWithdraw) Get() *Withdraw {
	return v.value
}

func (v *NullableWithdraw) Set(val *Withdraw) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdraw) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdraw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdraw(val *Withdraw) *NullableWithdraw {
	return &NullableWithdraw{value: val, isSet: true}
}

func (v NullableWithdraw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdraw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


