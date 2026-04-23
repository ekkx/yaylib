
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the StickerPacksResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StickerPacksResponse{}

// StickerPacksResponse struct for StickerPacksResponse
type StickerPacksResponse struct {
	StickerPacks []StickerPack `json:"sticker_packs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StickerPacksResponse StickerPacksResponse

// NewStickerPacksResponse instantiates a new StickerPacksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStickerPacksResponse() *StickerPacksResponse {
	this := StickerPacksResponse{}
	return &this
}

// NewStickerPacksResponseWithDefaults instantiates a new StickerPacksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerPacksResponseWithDefaults() *StickerPacksResponse {
	this := StickerPacksResponse{}
	return &this
}

// GetStickerPacks returns the StickerPacks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPacksResponse) GetStickerPacks() []StickerPack {
	if o == nil {
		var ret []StickerPack
		return ret
	}
	return o.StickerPacks
}

// GetStickerPacksOk returns a tuple with the StickerPacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPacksResponse) GetStickerPacksOk() ([]StickerPack, bool) {
	if o == nil || IsNil(o.StickerPacks) {
		return nil, false
	}
	return o.StickerPacks, true
}

// HasStickerPacks returns a boolean if a field has been set.
func (o *StickerPacksResponse) HasStickerPacks() bool {
	if o != nil && !IsNil(o.StickerPacks) {
		return true
	}

	return false
}

// SetStickerPacks gets a reference to the given []StickerPack and assigns it to the StickerPacks field.
func (o *StickerPacksResponse) SetStickerPacks(v []StickerPack) {
	o.StickerPacks = v
}

func (o StickerPacksResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StickerPacksResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StickerPacks != nil {
		toSerialize["sticker_packs"] = o.StickerPacks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StickerPacksResponse) UnmarshalJSON(data []byte) (err error) {
	varStickerPacksResponse := _StickerPacksResponse{}

	err = json.Unmarshal(data, &varStickerPacksResponse)

	if err != nil {
		return err
	}

	*o = StickerPacksResponse(varStickerPacksResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sticker_packs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStickerPacksResponse struct {
	value *StickerPacksResponse
	isSet bool
}

func (v NullableStickerPacksResponse) Get() *StickerPacksResponse {
	return v.value
}

func (v *NullableStickerPacksResponse) Set(val *StickerPacksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStickerPacksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStickerPacksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStickerPacksResponse(val *StickerPacksResponse) *NullableStickerPacksResponse {
	return &NullableStickerPacksResponse{value: val, isSet: true}
}

func (v NullableStickerPacksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStickerPacksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


