
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplWithdrawalTokenDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplWithdrawalTokenDTO{}

// Web3EmplWithdrawalTokenDTO struct for Web3EmplWithdrawalTokenDTO
type Web3EmplWithdrawalTokenDTO struct {
	ChainId NullableString `json:"chain_id,omitempty"`
	Id NullableString `json:"id,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	TokenAddress NullableString `json:"token_address,omitempty"`
	TokenName NullableString `json:"token_name,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	VeloPoolVersion NullableString `json:"velo_pool_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplWithdrawalTokenDTO Web3EmplWithdrawalTokenDTO

// NewWeb3EmplWithdrawalTokenDTO instantiates a new Web3EmplWithdrawalTokenDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplWithdrawalTokenDTO() *Web3EmplWithdrawalTokenDTO {
	this := Web3EmplWithdrawalTokenDTO{}
	return &this
}

// NewWeb3EmplWithdrawalTokenDTOWithDefaults instantiates a new Web3EmplWithdrawalTokenDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplWithdrawalTokenDTOWithDefaults() *Web3EmplWithdrawalTokenDTO {
	this := Web3EmplWithdrawalTokenDTO{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetChainId() string {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret string
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableString and assigns it to the ChainId field.
func (o *Web3EmplWithdrawalTokenDTO) SetChainId(v string) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Web3EmplWithdrawalTokenDTO) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetId() {
	o.Id.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Web3EmplWithdrawalTokenDTO) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetTokenAddress() string {
	if o == nil || IsNil(o.TokenAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenAddress.Get()
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenAddress.Get(), o.TokenAddress.IsSet()
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasTokenAddress() bool {
	if o != nil && o.TokenAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given NullableString and assigns it to the TokenAddress field.
func (o *Web3EmplWithdrawalTokenDTO) SetTokenAddress(v string) {
	o.TokenAddress.Set(&v)
}
// SetTokenAddressNil sets the value for TokenAddress to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetTokenAddressNil() {
	o.TokenAddress.Set(nil)
}

// UnsetTokenAddress ensures that no value is present for TokenAddress, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetTokenAddress() {
	o.TokenAddress.Unset()
}

// GetTokenName returns the TokenName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetTokenName() string {
	if o == nil || IsNil(o.TokenName.Get()) {
		var ret string
		return ret
	}
	return *o.TokenName.Get()
}

// GetTokenNameOk returns a tuple with the TokenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenName.Get(), o.TokenName.IsSet()
}

// HasTokenName returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasTokenName() bool {
	if o != nil && o.TokenName.IsSet() {
		return true
	}

	return false
}

// SetTokenName gets a reference to the given NullableString and assigns it to the TokenName field.
func (o *Web3EmplWithdrawalTokenDTO) SetTokenName(v string) {
	o.TokenName.Set(&v)
}
// SetTokenNameNil sets the value for TokenName to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetTokenNameNil() {
	o.TokenName.Set(nil)
}

// UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetTokenName() {
	o.TokenName.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *Web3EmplWithdrawalTokenDTO) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetVeloPoolVersion returns the VeloPoolVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokenDTO) GetVeloPoolVersion() string {
	if o == nil || IsNil(o.VeloPoolVersion.Get()) {
		var ret string
		return ret
	}
	return *o.VeloPoolVersion.Get()
}

// GetVeloPoolVersionOk returns a tuple with the VeloPoolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokenDTO) GetVeloPoolVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VeloPoolVersion.Get(), o.VeloPoolVersion.IsSet()
}

// HasVeloPoolVersion returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokenDTO) HasVeloPoolVersion() bool {
	if o != nil && o.VeloPoolVersion.IsSet() {
		return true
	}

	return false
}

// SetVeloPoolVersion gets a reference to the given NullableString and assigns it to the VeloPoolVersion field.
func (o *Web3EmplWithdrawalTokenDTO) SetVeloPoolVersion(v string) {
	o.VeloPoolVersion.Set(&v)
}
// SetVeloPoolVersionNil sets the value for VeloPoolVersion to be an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) SetVeloPoolVersionNil() {
	o.VeloPoolVersion.Set(nil)
}

// UnsetVeloPoolVersion ensures that no value is present for VeloPoolVersion, not even an explicit nil
func (o *Web3EmplWithdrawalTokenDTO) UnsetVeloPoolVersion() {
	o.VeloPoolVersion.Unset()
}

func (o Web3EmplWithdrawalTokenDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplWithdrawalTokenDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.TokenAddress.IsSet() {
		toSerialize["token_address"] = o.TokenAddress.Get()
	}
	if o.TokenName.IsSet() {
		toSerialize["token_name"] = o.TokenName.Get()
	}
	if o.TokenSymbol.IsSet() {
		toSerialize["token_symbol"] = o.TokenSymbol.Get()
	}
	if o.VeloPoolVersion.IsSet() {
		toSerialize["velo_pool_version"] = o.VeloPoolVersion.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplWithdrawalTokenDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplWithdrawalTokenDTO := _Web3EmplWithdrawalTokenDTO{}

	err = json.Unmarshal(data, &varWeb3EmplWithdrawalTokenDTO)

	if err != nil {
		return err
	}

	*o = Web3EmplWithdrawalTokenDTO(varWeb3EmplWithdrawalTokenDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "token_address")
		delete(additionalProperties, "token_name")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "velo_pool_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplWithdrawalTokenDTO struct {
	value *Web3EmplWithdrawalTokenDTO
	isSet bool
}

func (v NullableWeb3EmplWithdrawalTokenDTO) Get() *Web3EmplWithdrawalTokenDTO {
	return v.value
}

func (v *NullableWeb3EmplWithdrawalTokenDTO) Set(val *Web3EmplWithdrawalTokenDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplWithdrawalTokenDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplWithdrawalTokenDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplWithdrawalTokenDTO(val *Web3EmplWithdrawalTokenDTO) *NullableWeb3EmplWithdrawalTokenDTO {
	return &NullableWeb3EmplWithdrawalTokenDTO{value: val, isSet: true}
}

func (v NullableWeb3EmplWithdrawalTokenDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplWithdrawalTokenDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


