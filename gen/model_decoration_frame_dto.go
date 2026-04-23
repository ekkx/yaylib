
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DecorationFrameDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecorationFrameDTO{}

// DecorationFrameDTO struct for DecorationFrameDTO
type DecorationFrameDTO struct {
	Current NullableBool `json:"current,omitempty"`
	Frame NullableString `json:"frame,omitempty"`
	FrameThumbnail NullableString `json:"frame_thumbnail,omitempty"`
	FrameType NullableString `json:"frame_type,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Owned NullableBool `json:"owned,omitempty"`
	Price NullableInt64 `json:"price,omitempty"`
	VipOnly NullableBool `json:"vip_only,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DecorationFrameDTO DecorationFrameDTO

// NewDecorationFrameDTO instantiates a new DecorationFrameDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecorationFrameDTO() *DecorationFrameDTO {
	this := DecorationFrameDTO{}
	return &this
}

// NewDecorationFrameDTOWithDefaults instantiates a new DecorationFrameDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecorationFrameDTOWithDefaults() *DecorationFrameDTO {
	this := DecorationFrameDTO{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetCurrent() bool {
	if o == nil || IsNil(o.Current.Get()) {
		var ret bool
		return ret
	}
	return *o.Current.Get()
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Current.Get(), o.Current.IsSet()
}

// HasCurrent returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasCurrent() bool {
	if o != nil && o.Current.IsSet() {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given NullableBool and assigns it to the Current field.
func (o *DecorationFrameDTO) SetCurrent(v bool) {
	o.Current.Set(&v)
}
// SetCurrentNil sets the value for Current to be an explicit nil
func (o *DecorationFrameDTO) SetCurrentNil() {
	o.Current.Set(nil)
}

// UnsetCurrent ensures that no value is present for Current, not even an explicit nil
func (o *DecorationFrameDTO) UnsetCurrent() {
	o.Current.Unset()
}

// GetFrame returns the Frame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetFrame() string {
	if o == nil || IsNil(o.Frame.Get()) {
		var ret string
		return ret
	}
	return *o.Frame.Get()
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetFrameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frame.Get(), o.Frame.IsSet()
}

// HasFrame returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasFrame() bool {
	if o != nil && o.Frame.IsSet() {
		return true
	}

	return false
}

// SetFrame gets a reference to the given NullableString and assigns it to the Frame field.
func (o *DecorationFrameDTO) SetFrame(v string) {
	o.Frame.Set(&v)
}
// SetFrameNil sets the value for Frame to be an explicit nil
func (o *DecorationFrameDTO) SetFrameNil() {
	o.Frame.Set(nil)
}

// UnsetFrame ensures that no value is present for Frame, not even an explicit nil
func (o *DecorationFrameDTO) UnsetFrame() {
	o.Frame.Unset()
}

// GetFrameThumbnail returns the FrameThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetFrameThumbnail() string {
	if o == nil || IsNil(o.FrameThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.FrameThumbnail.Get()
}

// GetFrameThumbnailOk returns a tuple with the FrameThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetFrameThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameThumbnail.Get(), o.FrameThumbnail.IsSet()
}

// HasFrameThumbnail returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasFrameThumbnail() bool {
	if o != nil && o.FrameThumbnail.IsSet() {
		return true
	}

	return false
}

// SetFrameThumbnail gets a reference to the given NullableString and assigns it to the FrameThumbnail field.
func (o *DecorationFrameDTO) SetFrameThumbnail(v string) {
	o.FrameThumbnail.Set(&v)
}
// SetFrameThumbnailNil sets the value for FrameThumbnail to be an explicit nil
func (o *DecorationFrameDTO) SetFrameThumbnailNil() {
	o.FrameThumbnail.Set(nil)
}

// UnsetFrameThumbnail ensures that no value is present for FrameThumbnail, not even an explicit nil
func (o *DecorationFrameDTO) UnsetFrameThumbnail() {
	o.FrameThumbnail.Unset()
}

// GetFrameType returns the FrameType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetFrameType() string {
	if o == nil || IsNil(o.FrameType.Get()) {
		var ret string
		return ret
	}
	return *o.FrameType.Get()
}

// GetFrameTypeOk returns a tuple with the FrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetFrameTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameType.Get(), o.FrameType.IsSet()
}

// HasFrameType returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasFrameType() bool {
	if o != nil && o.FrameType.IsSet() {
		return true
	}

	return false
}

// SetFrameType gets a reference to the given NullableString and assigns it to the FrameType field.
func (o *DecorationFrameDTO) SetFrameType(v string) {
	o.FrameType.Set(&v)
}
// SetFrameTypeNil sets the value for FrameType to be an explicit nil
func (o *DecorationFrameDTO) SetFrameTypeNil() {
	o.FrameType.Set(nil)
}

// UnsetFrameType ensures that no value is present for FrameType, not even an explicit nil
func (o *DecorationFrameDTO) UnsetFrameType() {
	o.FrameType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *DecorationFrameDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DecorationFrameDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DecorationFrameDTO) UnsetId() {
	o.Id.Unset()
}

// GetOwned returns the Owned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetOwned() bool {
	if o == nil || IsNil(o.Owned.Get()) {
		var ret bool
		return ret
	}
	return *o.Owned.Get()
}

// GetOwnedOk returns a tuple with the Owned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetOwnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owned.Get(), o.Owned.IsSet()
}

// HasOwned returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasOwned() bool {
	if o != nil && o.Owned.IsSet() {
		return true
	}

	return false
}

// SetOwned gets a reference to the given NullableBool and assigns it to the Owned field.
func (o *DecorationFrameDTO) SetOwned(v bool) {
	o.Owned.Set(&v)
}
// SetOwnedNil sets the value for Owned to be an explicit nil
func (o *DecorationFrameDTO) SetOwnedNil() {
	o.Owned.Set(nil)
}

// UnsetOwned ensures that no value is present for Owned, not even an explicit nil
func (o *DecorationFrameDTO) UnsetOwned() {
	o.Owned.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetPrice() int64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret int64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt64 and assigns it to the Price field.
func (o *DecorationFrameDTO) SetPrice(v int64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *DecorationFrameDTO) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *DecorationFrameDTO) UnsetPrice() {
	o.Price.Unset()
}

// GetVipOnly returns the VipOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecorationFrameDTO) GetVipOnly() bool {
	if o == nil || IsNil(o.VipOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.VipOnly.Get()
}

// GetVipOnlyOk returns a tuple with the VipOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecorationFrameDTO) GetVipOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipOnly.Get(), o.VipOnly.IsSet()
}

// HasVipOnly returns a boolean if a field has been set.
func (o *DecorationFrameDTO) HasVipOnly() bool {
	if o != nil && o.VipOnly.IsSet() {
		return true
	}

	return false
}

// SetVipOnly gets a reference to the given NullableBool and assigns it to the VipOnly field.
func (o *DecorationFrameDTO) SetVipOnly(v bool) {
	o.VipOnly.Set(&v)
}
// SetVipOnlyNil sets the value for VipOnly to be an explicit nil
func (o *DecorationFrameDTO) SetVipOnlyNil() {
	o.VipOnly.Set(nil)
}

// UnsetVipOnly ensures that no value is present for VipOnly, not even an explicit nil
func (o *DecorationFrameDTO) UnsetVipOnly() {
	o.VipOnly.Unset()
}

func (o DecorationFrameDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecorationFrameDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Current.IsSet() {
		toSerialize["current"] = o.Current.Get()
	}
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
	if o.Owned.IsSet() {
		toSerialize["owned"] = o.Owned.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.VipOnly.IsSet() {
		toSerialize["vip_only"] = o.VipOnly.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DecorationFrameDTO) UnmarshalJSON(data []byte) (err error) {
	varDecorationFrameDTO := _DecorationFrameDTO{}

	err = json.Unmarshal(data, &varDecorationFrameDTO)

	if err != nil {
		return err
	}

	*o = DecorationFrameDTO(varDecorationFrameDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current")
		delete(additionalProperties, "frame")
		delete(additionalProperties, "frame_thumbnail")
		delete(additionalProperties, "frame_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "owned")
		delete(additionalProperties, "price")
		delete(additionalProperties, "vip_only")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDecorationFrameDTO struct {
	value *DecorationFrameDTO
	isSet bool
}

func (v NullableDecorationFrameDTO) Get() *DecorationFrameDTO {
	return v.value
}

func (v *NullableDecorationFrameDTO) Set(val *DecorationFrameDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDecorationFrameDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDecorationFrameDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecorationFrameDTO(val *DecorationFrameDTO) *NullableDecorationFrameDTO {
	return &NullableDecorationFrameDTO{value: val, isSet: true}
}

func (v NullableDecorationFrameDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecorationFrameDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


