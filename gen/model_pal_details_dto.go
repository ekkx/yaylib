
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalDetailsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalDetailsDTO{}

// PalDetailsDTO struct for PalDetailsDTO
type PalDetailsDTO struct {
	Cost NullableFloat64 `json:"cost,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	NftType NullableString `json:"nft_type,omitempty"`
	PalImage NullableString `json:"pal_image,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	Rank NullableString `json:"rank,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalDetailsDTO PalDetailsDTO

// NewPalDetailsDTO instantiates a new PalDetailsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalDetailsDTO() *PalDetailsDTO {
	this := PalDetailsDTO{}
	return &this
}

// NewPalDetailsDTOWithDefaults instantiates a new PalDetailsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalDetailsDTOWithDefaults() *PalDetailsDTO {
	this := PalDetailsDTO{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetCost() float64 {
	if o == nil || IsNil(o.Cost.Get()) {
		var ret float64
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableFloat64 and assigns it to the Cost field.
func (o *PalDetailsDTO) SetCost(v float64) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *PalDetailsDTO) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *PalDetailsDTO) UnsetCost() {
	o.Cost.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *PalDetailsDTO) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *PalDetailsDTO) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *PalDetailsDTO) UnsetLevel() {
	o.Level.Unset()
}

// GetNftType returns the NftType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetNftType() string {
	if o == nil || IsNil(o.NftType.Get()) {
		var ret string
		return ret
	}
	return *o.NftType.Get()
}

// GetNftTypeOk returns a tuple with the NftType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetNftTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftType.Get(), o.NftType.IsSet()
}

// HasNftType returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasNftType() bool {
	if o != nil && o.NftType.IsSet() {
		return true
	}

	return false
}

// SetNftType gets a reference to the given NullableString and assigns it to the NftType field.
func (o *PalDetailsDTO) SetNftType(v string) {
	o.NftType.Set(&v)
}
// SetNftTypeNil sets the value for NftType to be an explicit nil
func (o *PalDetailsDTO) SetNftTypeNil() {
	o.NftType.Set(nil)
}

// UnsetNftType ensures that no value is present for NftType, not even an explicit nil
func (o *PalDetailsDTO) UnsetNftType() {
	o.NftType.Unset()
}

// GetPalImage returns the PalImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetPalImage() string {
	if o == nil || IsNil(o.PalImage.Get()) {
		var ret string
		return ret
	}
	return *o.PalImage.Get()
}

// GetPalImageOk returns a tuple with the PalImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetPalImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImage.Get(), o.PalImage.IsSet()
}

// HasPalImage returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasPalImage() bool {
	if o != nil && o.PalImage.IsSet() {
		return true
	}

	return false
}

// SetPalImage gets a reference to the given NullableString and assigns it to the PalImage field.
func (o *PalDetailsDTO) SetPalImage(v string) {
	o.PalImage.Set(&v)
}
// SetPalImageNil sets the value for PalImage to be an explicit nil
func (o *PalDetailsDTO) SetPalImageNil() {
	o.PalImage.Set(nil)
}

// UnsetPalImage ensures that no value is present for PalImage, not even an explicit nil
func (o *PalDetailsDTO) UnsetPalImage() {
	o.PalImage.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *PalDetailsDTO) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *PalDetailsDTO) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *PalDetailsDTO) UnsetPalName() {
	o.PalName.Unset()
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetRank() string {
	if o == nil || IsNil(o.Rank.Get()) {
		var ret string
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetRankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableString and assigns it to the Rank field.
func (o *PalDetailsDTO) SetRank(v string) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *PalDetailsDTO) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *PalDetailsDTO) UnsetRank() {
	o.Rank.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsDTO) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsDTO) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *PalDetailsDTO) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *PalDetailsDTO) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *PalDetailsDTO) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *PalDetailsDTO) UnsetTokenId() {
	o.TokenId.Unset()
}

func (o PalDetailsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalDetailsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.NftType.IsSet() {
		toSerialize["nft_type"] = o.NftType.Get()
	}
	if o.PalImage.IsSet() {
		toSerialize["pal_image"] = o.PalImage.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalDetailsDTO) UnmarshalJSON(data []byte) (err error) {
	varPalDetailsDTO := _PalDetailsDTO{}

	err = json.Unmarshal(data, &varPalDetailsDTO)

	if err != nil {
		return err
	}

	*o = PalDetailsDTO(varPalDetailsDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost")
		delete(additionalProperties, "level")
		delete(additionalProperties, "nft_type")
		delete(additionalProperties, "pal_image")
		delete(additionalProperties, "pal_name")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalDetailsDTO struct {
	value *PalDetailsDTO
	isSet bool
}

func (v NullablePalDetailsDTO) Get() *PalDetailsDTO {
	return v.value
}

func (v *NullablePalDetailsDTO) Set(val *PalDetailsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePalDetailsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePalDetailsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalDetailsDTO(val *PalDetailsDTO) *NullablePalDetailsDTO {
	return &NullablePalDetailsDTO{value: val, isSet: true}
}

func (v NullablePalDetailsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalDetailsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


