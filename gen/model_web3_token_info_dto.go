
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3TokenInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3TokenInfoDTO{}

// Web3TokenInfoDTO struct for Web3TokenInfoDTO
type Web3TokenInfoDTO struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	LocalizedUrls NullableLocalizedStringDTO `json:"localized_urls,omitempty"`
	Symbol NullableString `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3TokenInfoDTO Web3TokenInfoDTO

// NewWeb3TokenInfoDTO instantiates a new Web3TokenInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3TokenInfoDTO() *Web3TokenInfoDTO {
	this := Web3TokenInfoDTO{}
	return &this
}

// NewWeb3TokenInfoDTOWithDefaults instantiates a new Web3TokenInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TokenInfoDTOWithDefaults() *Web3TokenInfoDTO {
	this := Web3TokenInfoDTO{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TokenInfoDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TokenInfoDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3TokenInfoDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3TokenInfoDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3TokenInfoDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3TokenInfoDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetLocalizedUrls returns the LocalizedUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TokenInfoDTO) GetLocalizedUrls() LocalizedStringDTO {
	if o == nil || IsNil(o.LocalizedUrls.Get()) {
		var ret LocalizedStringDTO
		return ret
	}
	return *o.LocalizedUrls.Get()
}

// GetLocalizedUrlsOk returns a tuple with the LocalizedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TokenInfoDTO) GetLocalizedUrlsOk() (*LocalizedStringDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedUrls.Get(), o.LocalizedUrls.IsSet()
}

// HasLocalizedUrls returns a boolean if a field has been set.
func (o *Web3TokenInfoDTO) HasLocalizedUrls() bool {
	if o != nil && o.LocalizedUrls.IsSet() {
		return true
	}

	return false
}

// SetLocalizedUrls gets a reference to the given NullableLocalizedStringDTO and assigns it to the LocalizedUrls field.
func (o *Web3TokenInfoDTO) SetLocalizedUrls(v LocalizedStringDTO) {
	o.LocalizedUrls.Set(&v)
}
// SetLocalizedUrlsNil sets the value for LocalizedUrls to be an explicit nil
func (o *Web3TokenInfoDTO) SetLocalizedUrlsNil() {
	o.LocalizedUrls.Set(nil)
}

// UnsetLocalizedUrls ensures that no value is present for LocalizedUrls, not even an explicit nil
func (o *Web3TokenInfoDTO) UnsetLocalizedUrls() {
	o.LocalizedUrls.Unset()
}

// GetSymbol returns the Symbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TokenInfoDTO) GetSymbol() string {
	if o == nil || IsNil(o.Symbol.Get()) {
		var ret string
		return ret
	}
	return *o.Symbol.Get()
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TokenInfoDTO) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Symbol.Get(), o.Symbol.IsSet()
}

// HasSymbol returns a boolean if a field has been set.
func (o *Web3TokenInfoDTO) HasSymbol() bool {
	if o != nil && o.Symbol.IsSet() {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given NullableString and assigns it to the Symbol field.
func (o *Web3TokenInfoDTO) SetSymbol(v string) {
	o.Symbol.Set(&v)
}
// SetSymbolNil sets the value for Symbol to be an explicit nil
func (o *Web3TokenInfoDTO) SetSymbolNil() {
	o.Symbol.Set(nil)
}

// UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
func (o *Web3TokenInfoDTO) UnsetSymbol() {
	o.Symbol.Unset()
}

func (o Web3TokenInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3TokenInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.LocalizedUrls.IsSet() {
		toSerialize["localized_urls"] = o.LocalizedUrls.Get()
	}
	if o.Symbol.IsSet() {
		toSerialize["symbol"] = o.Symbol.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3TokenInfoDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3TokenInfoDTO := _Web3TokenInfoDTO{}

	err = json.Unmarshal(data, &varWeb3TokenInfoDTO)

	if err != nil {
		return err
	}

	*o = Web3TokenInfoDTO(varWeb3TokenInfoDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "localized_urls")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3TokenInfoDTO struct {
	value *Web3TokenInfoDTO
	isSet bool
}

func (v NullableWeb3TokenInfoDTO) Get() *Web3TokenInfoDTO {
	return v.value
}

func (v *NullableWeb3TokenInfoDTO) Set(val *Web3TokenInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3TokenInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3TokenInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3TokenInfoDTO(val *Web3TokenInfoDTO) *NullableWeb3TokenInfoDTO {
	return &NullableWeb3TokenInfoDTO{value: val, isSet: true}
}

func (v NullableWeb3TokenInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3TokenInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


