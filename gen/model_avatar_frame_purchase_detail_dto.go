
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AvatarFramePurchaseDetailDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvatarFramePurchaseDetailDTO{}

// AvatarFramePurchaseDetailDTO struct for AvatarFramePurchaseDetailDTO
type AvatarFramePurchaseDetailDTO struct {
	Frame NullableString `json:"frame,omitempty"`
	FrameThumbnail NullableString `json:"frame_thumbnail,omitempty"`
	FrameType NullableString `json:"frame_type,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	VipOnly NullableBool `json:"vip_only,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvatarFramePurchaseDetailDTO AvatarFramePurchaseDetailDTO

// NewAvatarFramePurchaseDetailDTO instantiates a new AvatarFramePurchaseDetailDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvatarFramePurchaseDetailDTO() *AvatarFramePurchaseDetailDTO {
	this := AvatarFramePurchaseDetailDTO{}
	return &this
}

// NewAvatarFramePurchaseDetailDTOWithDefaults instantiates a new AvatarFramePurchaseDetailDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvatarFramePurchaseDetailDTOWithDefaults() *AvatarFramePurchaseDetailDTO {
	this := AvatarFramePurchaseDetailDTO{}
	return &this
}

// GetFrame returns the Frame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvatarFramePurchaseDetailDTO) GetFrame() string {
	if o == nil || IsNil(o.Frame.Get()) {
		var ret string
		return ret
	}
	return *o.Frame.Get()
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvatarFramePurchaseDetailDTO) GetFrameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frame.Get(), o.Frame.IsSet()
}

// HasFrame returns a boolean if a field has been set.
func (o *AvatarFramePurchaseDetailDTO) HasFrame() bool {
	if o != nil && o.Frame.IsSet() {
		return true
	}

	return false
}

// SetFrame gets a reference to the given NullableString and assigns it to the Frame field.
func (o *AvatarFramePurchaseDetailDTO) SetFrame(v string) {
	o.Frame.Set(&v)
}
// SetFrameNil sets the value for Frame to be an explicit nil
func (o *AvatarFramePurchaseDetailDTO) SetFrameNil() {
	o.Frame.Set(nil)
}

// UnsetFrame ensures that no value is present for Frame, not even an explicit nil
func (o *AvatarFramePurchaseDetailDTO) UnsetFrame() {
	o.Frame.Unset()
}

// GetFrameThumbnail returns the FrameThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvatarFramePurchaseDetailDTO) GetFrameThumbnail() string {
	if o == nil || IsNil(o.FrameThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.FrameThumbnail.Get()
}

// GetFrameThumbnailOk returns a tuple with the FrameThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvatarFramePurchaseDetailDTO) GetFrameThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameThumbnail.Get(), o.FrameThumbnail.IsSet()
}

// HasFrameThumbnail returns a boolean if a field has been set.
func (o *AvatarFramePurchaseDetailDTO) HasFrameThumbnail() bool {
	if o != nil && o.FrameThumbnail.IsSet() {
		return true
	}

	return false
}

// SetFrameThumbnail gets a reference to the given NullableString and assigns it to the FrameThumbnail field.
func (o *AvatarFramePurchaseDetailDTO) SetFrameThumbnail(v string) {
	o.FrameThumbnail.Set(&v)
}
// SetFrameThumbnailNil sets the value for FrameThumbnail to be an explicit nil
func (o *AvatarFramePurchaseDetailDTO) SetFrameThumbnailNil() {
	o.FrameThumbnail.Set(nil)
}

// UnsetFrameThumbnail ensures that no value is present for FrameThumbnail, not even an explicit nil
func (o *AvatarFramePurchaseDetailDTO) UnsetFrameThumbnail() {
	o.FrameThumbnail.Unset()
}

// GetFrameType returns the FrameType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvatarFramePurchaseDetailDTO) GetFrameType() string {
	if o == nil || IsNil(o.FrameType.Get()) {
		var ret string
		return ret
	}
	return *o.FrameType.Get()
}

// GetFrameTypeOk returns a tuple with the FrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvatarFramePurchaseDetailDTO) GetFrameTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameType.Get(), o.FrameType.IsSet()
}

// HasFrameType returns a boolean if a field has been set.
func (o *AvatarFramePurchaseDetailDTO) HasFrameType() bool {
	if o != nil && o.FrameType.IsSet() {
		return true
	}

	return false
}

// SetFrameType gets a reference to the given NullableString and assigns it to the FrameType field.
func (o *AvatarFramePurchaseDetailDTO) SetFrameType(v string) {
	o.FrameType.Set(&v)
}
// SetFrameTypeNil sets the value for FrameType to be an explicit nil
func (o *AvatarFramePurchaseDetailDTO) SetFrameTypeNil() {
	o.FrameType.Set(nil)
}

// UnsetFrameType ensures that no value is present for FrameType, not even an explicit nil
func (o *AvatarFramePurchaseDetailDTO) UnsetFrameType() {
	o.FrameType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvatarFramePurchaseDetailDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvatarFramePurchaseDetailDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AvatarFramePurchaseDetailDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *AvatarFramePurchaseDetailDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AvatarFramePurchaseDetailDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AvatarFramePurchaseDetailDTO) UnsetId() {
	o.Id.Unset()
}

// GetVipOnly returns the VipOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvatarFramePurchaseDetailDTO) GetVipOnly() bool {
	if o == nil || IsNil(o.VipOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.VipOnly.Get()
}

// GetVipOnlyOk returns a tuple with the VipOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvatarFramePurchaseDetailDTO) GetVipOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipOnly.Get(), o.VipOnly.IsSet()
}

// HasVipOnly returns a boolean if a field has been set.
func (o *AvatarFramePurchaseDetailDTO) HasVipOnly() bool {
	if o != nil && o.VipOnly.IsSet() {
		return true
	}

	return false
}

// SetVipOnly gets a reference to the given NullableBool and assigns it to the VipOnly field.
func (o *AvatarFramePurchaseDetailDTO) SetVipOnly(v bool) {
	o.VipOnly.Set(&v)
}
// SetVipOnlyNil sets the value for VipOnly to be an explicit nil
func (o *AvatarFramePurchaseDetailDTO) SetVipOnlyNil() {
	o.VipOnly.Set(nil)
}

// UnsetVipOnly ensures that no value is present for VipOnly, not even an explicit nil
func (o *AvatarFramePurchaseDetailDTO) UnsetVipOnly() {
	o.VipOnly.Unset()
}

func (o AvatarFramePurchaseDetailDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvatarFramePurchaseDetailDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Frame.IsSet() {
		toSerialize["frame"] = o.Frame.Get()
	}
	if o.FrameThumbnail.IsSet() {
		toSerialize["frame_thumbnail"] = o.FrameThumbnail.Get()
	}
	if o.FrameType.IsSet() {
		toSerialize["frame_type"] = o.FrameType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.VipOnly.IsSet() {
		toSerialize["vip_only"] = o.VipOnly.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvatarFramePurchaseDetailDTO) UnmarshalJSON(data []byte) (err error) {
	varAvatarFramePurchaseDetailDTO := _AvatarFramePurchaseDetailDTO{}

	err = json.Unmarshal(data, &varAvatarFramePurchaseDetailDTO)

	if err != nil {
		return err
	}

	*o = AvatarFramePurchaseDetailDTO(varAvatarFramePurchaseDetailDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "frame")
		delete(additionalProperties, "frame_thumbnail")
		delete(additionalProperties, "frame_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "vip_only")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvatarFramePurchaseDetailDTO struct {
	value *AvatarFramePurchaseDetailDTO
	isSet bool
}

func (v NullableAvatarFramePurchaseDetailDTO) Get() *AvatarFramePurchaseDetailDTO {
	return v.value
}

func (v *NullableAvatarFramePurchaseDetailDTO) Set(val *AvatarFramePurchaseDetailDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAvatarFramePurchaseDetailDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAvatarFramePurchaseDetailDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvatarFramePurchaseDetailDTO(val *AvatarFramePurchaseDetailDTO) *NullableAvatarFramePurchaseDetailDTO {
	return &NullableAvatarFramePurchaseDetailDTO{value: val, isSet: true}
}

func (v NullableAvatarFramePurchaseDetailDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvatarFramePurchaseDetailDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


