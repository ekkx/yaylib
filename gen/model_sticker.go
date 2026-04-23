
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Sticker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sticker{}

// Sticker struct for Sticker
type Sticker struct {
	Extension NullableString `json:"extension,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	StickerPackId NullableInt64 `json:"sticker_pack_id,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Width NullableInt32 `json:"width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sticker Sticker

// NewSticker instantiates a new Sticker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSticker() *Sticker {
	this := Sticker{}
	return &this
}

// NewStickerWithDefaults instantiates a new Sticker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerWithDefaults() *Sticker {
	this := Sticker{}
	return &this
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetExtension() string {
	if o == nil || IsNil(o.Extension.Get()) {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *Sticker) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *Sticker) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *Sticker) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *Sticker) UnsetExtension() {
	o.Extension.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *Sticker) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *Sticker) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *Sticker) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *Sticker) UnsetHeight() {
	o.Height.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Sticker) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Sticker) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Sticker) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Sticker) UnsetId() {
	o.Id.Unset()
}

// GetStickerPackId returns the StickerPackId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetStickerPackId() int64 {
	if o == nil || IsNil(o.StickerPackId.Get()) {
		var ret int64
		return ret
	}
	return *o.StickerPackId.Get()
}

// GetStickerPackIdOk returns a tuple with the StickerPackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetStickerPackIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StickerPackId.Get(), o.StickerPackId.IsSet()
}

// HasStickerPackId returns a boolean if a field has been set.
func (o *Sticker) HasStickerPackId() bool {
	if o != nil && o.StickerPackId.IsSet() {
		return true
	}

	return false
}

// SetStickerPackId gets a reference to the given NullableInt64 and assigns it to the StickerPackId field.
func (o *Sticker) SetStickerPackId(v int64) {
	o.StickerPackId.Set(&v)
}
// SetStickerPackIdNil sets the value for StickerPackId to be an explicit nil
func (o *Sticker) SetStickerPackIdNil() {
	o.StickerPackId.Set(nil)
}

// UnsetStickerPackId ensures that no value is present for StickerPackId, not even an explicit nil
func (o *Sticker) UnsetStickerPackId() {
	o.StickerPackId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Sticker) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Sticker) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Sticker) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Sticker) UnsetUrl() {
	o.Url.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sticker) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Sticker) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *Sticker) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *Sticker) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *Sticker) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *Sticker) UnsetWidth() {
	o.Width.Unset()
}

func (o Sticker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sticker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.StickerPackId.IsSet() {
		toSerialize["sticker_pack_id"] = o.StickerPackId.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Sticker) UnmarshalJSON(data []byte) (err error) {
	varSticker := _Sticker{}

	err = json.Unmarshal(data, &varSticker)

	if err != nil {
		return err
	}

	*o = Sticker(varSticker)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extension")
		delete(additionalProperties, "height")
		delete(additionalProperties, "id")
		delete(additionalProperties, "sticker_pack_id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSticker struct {
	value *Sticker
	isSet bool
}

func (v NullableSticker) Get() *Sticker {
	return v.value
}

func (v *NullableSticker) Set(val *Sticker) {
	v.value = val
	v.isSet = true
}

func (v NullableSticker) IsSet() bool {
	return v.isSet
}

func (v *NullableSticker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSticker(val *Sticker) *NullableSticker {
	return &NullableSticker{value: val, isSet: true}
}

func (v NullableSticker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSticker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


