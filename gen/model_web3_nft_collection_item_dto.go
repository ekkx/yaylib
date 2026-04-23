
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3NftCollectionItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3NftCollectionItemDTO{}

// Web3NftCollectionItemDTO struct for Web3NftCollectionItemDTO
type Web3NftCollectionItemDTO struct {
	Balance NullableString `json:"balance,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	ExtraMetadata NullableNFTMetadataDTO `json:"extra_metadata,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	TokenType NullableString `json:"token_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3NftCollectionItemDTO Web3NftCollectionItemDTO

// NewWeb3NftCollectionItemDTO instantiates a new Web3NftCollectionItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3NftCollectionItemDTO() *Web3NftCollectionItemDTO {
	this := Web3NftCollectionItemDTO{}
	return &this
}

// NewWeb3NftCollectionItemDTOWithDefaults instantiates a new Web3NftCollectionItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3NftCollectionItemDTOWithDefaults() *Web3NftCollectionItemDTO {
	this := Web3NftCollectionItemDTO{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetBalance() string {
	if o == nil || IsNil(o.Balance.Get()) {
		var ret string
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableString and assigns it to the Balance field.
func (o *Web3NftCollectionItemDTO) SetBalance(v string) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetBalance() {
	o.Balance.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3NftCollectionItemDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3NftCollectionItemDTO) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetExtraMetadata returns the ExtraMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetExtraMetadata() NFTMetadataDTO {
	if o == nil || IsNil(o.ExtraMetadata.Get()) {
		var ret NFTMetadataDTO
		return ret
	}
	return *o.ExtraMetadata.Get()
}

// GetExtraMetadataOk returns a tuple with the ExtraMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetExtraMetadataOk() (*NFTMetadataDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraMetadata.Get(), o.ExtraMetadata.IsSet()
}

// HasExtraMetadata returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasExtraMetadata() bool {
	if o != nil && o.ExtraMetadata.IsSet() {
		return true
	}

	return false
}

// SetExtraMetadata gets a reference to the given NullableNFTMetadataDTO and assigns it to the ExtraMetadata field.
func (o *Web3NftCollectionItemDTO) SetExtraMetadata(v NFTMetadataDTO) {
	o.ExtraMetadata.Set(&v)
}
// SetExtraMetadataNil sets the value for ExtraMetadata to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetExtraMetadataNil() {
	o.ExtraMetadata.Set(nil)
}

// UnsetExtraMetadata ensures that no value is present for ExtraMetadata, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetExtraMetadata() {
	o.ExtraMetadata.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Web3NftCollectionItemDTO) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Web3NftCollectionItemDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetName() {
	o.Name.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *Web3NftCollectionItemDTO) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetTokenType returns the TokenType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemDTO) GetTokenType() string {
	if o == nil || IsNil(o.TokenType.Get()) {
		var ret string
		return ret
	}
	return *o.TokenType.Get()
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemDTO) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenType.Get(), o.TokenType.IsSet()
}

// HasTokenType returns a boolean if a field has been set.
func (o *Web3NftCollectionItemDTO) HasTokenType() bool {
	if o != nil && o.TokenType.IsSet() {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given NullableString and assigns it to the TokenType field.
func (o *Web3NftCollectionItemDTO) SetTokenType(v string) {
	o.TokenType.Set(&v)
}
// SetTokenTypeNil sets the value for TokenType to be an explicit nil
func (o *Web3NftCollectionItemDTO) SetTokenTypeNil() {
	o.TokenType.Set(nil)
}

// UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
func (o *Web3NftCollectionItemDTO) UnsetTokenType() {
	o.TokenType.Unset()
}

func (o Web3NftCollectionItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3NftCollectionItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.ExtraMetadata.IsSet() {
		toSerialize["extra_metadata"] = o.ExtraMetadata.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.TokenType.IsSet() {
		toSerialize["token_type"] = o.TokenType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3NftCollectionItemDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3NftCollectionItemDTO := _Web3NftCollectionItemDTO{}

	err = json.Unmarshal(data, &varWeb3NftCollectionItemDTO)

	if err != nil {
		return err
	}

	*o = Web3NftCollectionItemDTO(varWeb3NftCollectionItemDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "extra_metadata")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3NftCollectionItemDTO struct {
	value *Web3NftCollectionItemDTO
	isSet bool
}

func (v NullableWeb3NftCollectionItemDTO) Get() *Web3NftCollectionItemDTO {
	return v.value
}

func (v *NullableWeb3NftCollectionItemDTO) Set(val *Web3NftCollectionItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3NftCollectionItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3NftCollectionItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3NftCollectionItemDTO(val *Web3NftCollectionItemDTO) *NullableWeb3NftCollectionItemDTO {
	return &NullableWeb3NftCollectionItemDTO{value: val, isSet: true}
}

func (v NullableWeb3NftCollectionItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3NftCollectionItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


