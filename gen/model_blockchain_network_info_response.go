
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BlockchainNetworkInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainNetworkInfoResponse{}

// BlockchainNetworkInfoResponse struct for BlockchainNetworkInfoResponse
type BlockchainNetworkInfoResponse struct {
	Chains []BlockchainDTO `json:"chains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockchainNetworkInfoResponse BlockchainNetworkInfoResponse

// NewBlockchainNetworkInfoResponse instantiates a new BlockchainNetworkInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainNetworkInfoResponse() *BlockchainNetworkInfoResponse {
	this := BlockchainNetworkInfoResponse{}
	return &this
}

// NewBlockchainNetworkInfoResponseWithDefaults instantiates a new BlockchainNetworkInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainNetworkInfoResponseWithDefaults() *BlockchainNetworkInfoResponse {
	this := BlockchainNetworkInfoResponse{}
	return &this
}

// GetChains returns the Chains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainNetworkInfoResponse) GetChains() []BlockchainDTO {
	if o == nil {
		var ret []BlockchainDTO
		return ret
	}
	return o.Chains
}

// GetChainsOk returns a tuple with the Chains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainNetworkInfoResponse) GetChainsOk() ([]BlockchainDTO, bool) {
	if o == nil || IsNil(o.Chains) {
		return nil, false
	}
	return o.Chains, true
}

// HasChains returns a boolean if a field has been set.
func (o *BlockchainNetworkInfoResponse) HasChains() bool {
	if o != nil && !IsNil(o.Chains) {
		return true
	}

	return false
}

// SetChains gets a reference to the given []BlockchainDTO and assigns it to the Chains field.
func (o *BlockchainNetworkInfoResponse) SetChains(v []BlockchainDTO) {
	o.Chains = v
}

func (o BlockchainNetworkInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainNetworkInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Chains != nil {
		toSerialize["chains"] = o.Chains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockchainNetworkInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varBlockchainNetworkInfoResponse := _BlockchainNetworkInfoResponse{}

	err = json.Unmarshal(data, &varBlockchainNetworkInfoResponse)

	if err != nil {
		return err
	}

	*o = BlockchainNetworkInfoResponse(varBlockchainNetworkInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockchainNetworkInfoResponse struct {
	value *BlockchainNetworkInfoResponse
	isSet bool
}

func (v NullableBlockchainNetworkInfoResponse) Get() *BlockchainNetworkInfoResponse {
	return v.value
}

func (v *NullableBlockchainNetworkInfoResponse) Set(val *BlockchainNetworkInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainNetworkInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainNetworkInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainNetworkInfoResponse(val *BlockchainNetworkInfoResponse) *NullableBlockchainNetworkInfoResponse {
	return &NullableBlockchainNetworkInfoResponse{value: val, isSet: true}
}

func (v NullableBlockchainNetworkInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainNetworkInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


