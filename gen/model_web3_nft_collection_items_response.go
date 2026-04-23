
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3NftCollectionItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3NftCollectionItemsResponse{}

// Web3NftCollectionItemsResponse struct for Web3NftCollectionItemsResponse
type Web3NftCollectionItemsResponse struct {
	Data []Web3NftCollectionItemDTO `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3NftCollectionItemsResponse Web3NftCollectionItemsResponse

// NewWeb3NftCollectionItemsResponse instantiates a new Web3NftCollectionItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3NftCollectionItemsResponse() *Web3NftCollectionItemsResponse {
	this := Web3NftCollectionItemsResponse{}
	return &this
}

// NewWeb3NftCollectionItemsResponseWithDefaults instantiates a new Web3NftCollectionItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3NftCollectionItemsResponseWithDefaults() *Web3NftCollectionItemsResponse {
	this := Web3NftCollectionItemsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftCollectionItemsResponse) GetData() []Web3NftCollectionItemDTO {
	if o == nil {
		var ret []Web3NftCollectionItemDTO
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftCollectionItemsResponse) GetDataOk() ([]Web3NftCollectionItemDTO, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Web3NftCollectionItemsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Web3NftCollectionItemDTO and assigns it to the Data field.
func (o *Web3NftCollectionItemsResponse) SetData(v []Web3NftCollectionItemDTO) {
	o.Data = v
}

func (o Web3NftCollectionItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3NftCollectionItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3NftCollectionItemsResponse) UnmarshalJSON(data []byte) (err error) {
	varWeb3NftCollectionItemsResponse := _Web3NftCollectionItemsResponse{}

	err = json.Unmarshal(data, &varWeb3NftCollectionItemsResponse)

	if err != nil {
		return err
	}

	*o = Web3NftCollectionItemsResponse(varWeb3NftCollectionItemsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3NftCollectionItemsResponse struct {
	value *Web3NftCollectionItemsResponse
	isSet bool
}

func (v NullableWeb3NftCollectionItemsResponse) Get() *Web3NftCollectionItemsResponse {
	return v.value
}

func (v *NullableWeb3NftCollectionItemsResponse) Set(val *Web3NftCollectionItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3NftCollectionItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3NftCollectionItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3NftCollectionItemsResponse(val *Web3NftCollectionItemsResponse) *NullableWeb3NftCollectionItemsResponse {
	return &NullableWeb3NftCollectionItemsResponse{value: val, isSet: true}
}

func (v NullableWeb3NftCollectionItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3NftCollectionItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


