
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AssetInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetInfoDTO{}

// AssetInfoDTO struct for AssetInfoDTO
type AssetInfoDTO struct {
	Image NullableString `json:"image,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetInfoDTO AssetInfoDTO

// NewAssetInfoDTO instantiates a new AssetInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetInfoDTO() *AssetInfoDTO {
	this := AssetInfoDTO{}
	return &this
}

// NewAssetInfoDTOWithDefaults instantiates a new AssetInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetInfoDTOWithDefaults() *AssetInfoDTO {
	this := AssetInfoDTO{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoDTO) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoDTO) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *AssetInfoDTO) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *AssetInfoDTO) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *AssetInfoDTO) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *AssetInfoDTO) UnsetImage() {
	o.Image.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AssetInfoDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AssetInfoDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AssetInfoDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AssetInfoDTO) UnsetName() {
	o.Name.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoDTO) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoDTO) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *AssetInfoDTO) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *AssetInfoDTO) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *AssetInfoDTO) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *AssetInfoDTO) UnsetTokenId() {
	o.TokenId.Unset()
}

func (o AssetInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetInfoDTO) UnmarshalJSON(data []byte) (err error) {
	varAssetInfoDTO := _AssetInfoDTO{}

	err = json.Unmarshal(data, &varAssetInfoDTO)

	if err != nil {
		return err
	}

	*o = AssetInfoDTO(varAssetInfoDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		delete(additionalProperties, "name")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetInfoDTO struct {
	value *AssetInfoDTO
	isSet bool
}

func (v NullableAssetInfoDTO) Get() *AssetInfoDTO {
	return v.value
}

func (v *NullableAssetInfoDTO) Set(val *AssetInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetInfoDTO(val *AssetInfoDTO) *NullableAssetInfoDTO {
	return &NullableAssetInfoDTO{value: val, isSet: true}
}

func (v NullableAssetInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


