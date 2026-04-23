
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Deposit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deposit{}

// Deposit struct for Deposit
type Deposit struct {
	AssetType NullableAssetType `json:"asset_type,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Deposit Deposit

// NewDeposit instantiates a new Deposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeposit() *Deposit {
	this := Deposit{}
	return &this
}

// NewDepositWithDefaults instantiates a new Deposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositWithDefaults() *Deposit {
	this := Deposit{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deposit) GetAssetType() AssetType {
	if o == nil || IsNil(o.AssetType.Get()) {
		var ret AssetType
		return ret
	}
	return *o.AssetType.Get()
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deposit) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetType.Get(), o.AssetType.IsSet()
}

// HasAssetType returns a boolean if a field has been set.
func (o *Deposit) HasAssetType() bool {
	if o != nil && o.AssetType.IsSet() {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given NullableAssetType and assigns it to the AssetType field.
func (o *Deposit) SetAssetType(v AssetType) {
	o.AssetType.Set(&v)
}
// SetAssetTypeNil sets the value for AssetType to be an explicit nil
func (o *Deposit) SetAssetTypeNil() {
	o.AssetType.Set(nil)
}

// UnsetAssetType ensures that no value is present for AssetType, not even an explicit nil
func (o *Deposit) UnsetAssetType() {
	o.AssetType.Unset()
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deposit) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deposit) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *Deposit) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *Deposit) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *Deposit) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *Deposit) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deposit) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deposit) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *Deposit) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *Deposit) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *Deposit) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *Deposit) UnsetPalName() {
	o.PalName.Unset()
}

func (o Deposit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deposit) ToMap() (map[string]interface{}, error) {
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

func (o *Deposit) UnmarshalJSON(data []byte) (err error) {
	varDeposit := _Deposit{}

	err = json.Unmarshal(data, &varDeposit)

	if err != nil {
		return err
	}

	*o = Deposit(varDeposit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_type")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeposit struct {
	value *Deposit
	isSet bool
}

func (v NullableDeposit) Get() *Deposit {
	return v.value
}

func (v *NullableDeposit) Set(val *Deposit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeposit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeposit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeposit(val *Deposit) *NullableDeposit {
	return &NullableDeposit{value: val, isSet: true}
}

func (v NullableDeposit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeposit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


