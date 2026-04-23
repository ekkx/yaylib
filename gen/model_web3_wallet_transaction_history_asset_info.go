
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletTransactionHistoryAssetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletTransactionHistoryAssetInfo{}

// Web3WalletTransactionHistoryAssetInfo struct for Web3WalletTransactionHistoryAssetInfo
type Web3WalletTransactionHistoryAssetInfo struct {
	ImageUrl NullableString `json:"image_url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	VideoUrl NullableString `json:"video_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletTransactionHistoryAssetInfo Web3WalletTransactionHistoryAssetInfo

// NewWeb3WalletTransactionHistoryAssetInfo instantiates a new Web3WalletTransactionHistoryAssetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletTransactionHistoryAssetInfo() *Web3WalletTransactionHistoryAssetInfo {
	this := Web3WalletTransactionHistoryAssetInfo{}
	return &this
}

// NewWeb3WalletTransactionHistoryAssetInfoWithDefaults instantiates a new Web3WalletTransactionHistoryAssetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletTransactionHistoryAssetInfoWithDefaults() *Web3WalletTransactionHistoryAssetInfo {
	this := Web3WalletTransactionHistoryAssetInfo{}
	return &this
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistoryAssetInfo) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistoryAssetInfo) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistoryAssetInfo) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Web3WalletTransactionHistoryAssetInfo) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistoryAssetInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistoryAssetInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistoryAssetInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Web3WalletTransactionHistoryAssetInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) UnsetName() {
	o.Name.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistoryAssetInfo) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistoryAssetInfo) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistoryAssetInfo) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *Web3WalletTransactionHistoryAssetInfo) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletTransactionHistoryAssetInfo) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletTransactionHistoryAssetInfo) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *Web3WalletTransactionHistoryAssetInfo) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *Web3WalletTransactionHistoryAssetInfo) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *Web3WalletTransactionHistoryAssetInfo) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

func (o Web3WalletTransactionHistoryAssetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletTransactionHistoryAssetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.VideoUrl.IsSet() {
		toSerialize["video_url"] = o.VideoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletTransactionHistoryAssetInfo) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletTransactionHistoryAssetInfo := _Web3WalletTransactionHistoryAssetInfo{}

	err = json.Unmarshal(data, &varWeb3WalletTransactionHistoryAssetInfo)

	if err != nil {
		return err
	}

	*o = Web3WalletTransactionHistoryAssetInfo(varWeb3WalletTransactionHistoryAssetInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "video_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletTransactionHistoryAssetInfo struct {
	value *Web3WalletTransactionHistoryAssetInfo
	isSet bool
}

func (v NullableWeb3WalletTransactionHistoryAssetInfo) Get() *Web3WalletTransactionHistoryAssetInfo {
	return v.value
}

func (v *NullableWeb3WalletTransactionHistoryAssetInfo) Set(val *Web3WalletTransactionHistoryAssetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletTransactionHistoryAssetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletTransactionHistoryAssetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletTransactionHistoryAssetInfo(val *Web3WalletTransactionHistoryAssetInfo) *NullableWeb3WalletTransactionHistoryAssetInfo {
	return &NullableWeb3WalletTransactionHistoryAssetInfo{value: val, isSet: true}
}

func (v NullableWeb3WalletTransactionHistoryAssetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletTransactionHistoryAssetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


