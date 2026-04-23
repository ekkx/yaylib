
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelVideo{}

// ModelVideo struct for ModelVideo
type ModelVideo struct {
	Bitrate NullableInt32 `json:"bitrate,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsCompleted NullableBool `json:"is_completed,omitempty"`
	ThumbnailBigUrl NullableString `json:"thumbnail_big_url,omitempty"`
	ThumbnailUrl NullableString `json:"thumbnail_url,omitempty"`
	VideoUrl NullableString `json:"video_url,omitempty"`
	ViewsCount NullableInt32 `json:"views_count,omitempty"`
	Width NullableInt32 `json:"width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelVideo ModelVideo

// NewModelVideo instantiates a new ModelVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelVideo() *ModelVideo {
	this := ModelVideo{}
	return &this
}

// NewModelVideoWithDefaults instantiates a new ModelVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelVideoWithDefaults() *ModelVideo {
	this := ModelVideo{}
	return &this
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *ModelVideo) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *ModelVideo) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}
// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *ModelVideo) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *ModelVideo) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *ModelVideo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *ModelVideo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *ModelVideo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *ModelVideo) UnsetHeight() {
	o.Height.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelVideo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelVideo) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelVideo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelVideo) UnsetId() {
	o.Id.Unset()
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCompleted.Get()
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCompleted.Get(), o.IsCompleted.IsSet()
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *ModelVideo) HasIsCompleted() bool {
	if o != nil && o.IsCompleted.IsSet() {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given NullableBool and assigns it to the IsCompleted field.
func (o *ModelVideo) SetIsCompleted(v bool) {
	o.IsCompleted.Set(&v)
}
// SetIsCompletedNil sets the value for IsCompleted to be an explicit nil
func (o *ModelVideo) SetIsCompletedNil() {
	o.IsCompleted.Set(nil)
}

// UnsetIsCompleted ensures that no value is present for IsCompleted, not even an explicit nil
func (o *ModelVideo) UnsetIsCompleted() {
	o.IsCompleted.Unset()
}

// GetThumbnailBigUrl returns the ThumbnailBigUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetThumbnailBigUrl() string {
	if o == nil || IsNil(o.ThumbnailBigUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailBigUrl.Get()
}

// GetThumbnailBigUrlOk returns a tuple with the ThumbnailBigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetThumbnailBigUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailBigUrl.Get(), o.ThumbnailBigUrl.IsSet()
}

// HasThumbnailBigUrl returns a boolean if a field has been set.
func (o *ModelVideo) HasThumbnailBigUrl() bool {
	if o != nil && o.ThumbnailBigUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailBigUrl gets a reference to the given NullableString and assigns it to the ThumbnailBigUrl field.
func (o *ModelVideo) SetThumbnailBigUrl(v string) {
	o.ThumbnailBigUrl.Set(&v)
}
// SetThumbnailBigUrlNil sets the value for ThumbnailBigUrl to be an explicit nil
func (o *ModelVideo) SetThumbnailBigUrlNil() {
	o.ThumbnailBigUrl.Set(nil)
}

// UnsetThumbnailBigUrl ensures that no value is present for ThumbnailBigUrl, not even an explicit nil
func (o *ModelVideo) UnsetThumbnailBigUrl() {
	o.ThumbnailBigUrl.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *ModelVideo) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *ModelVideo) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *ModelVideo) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *ModelVideo) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *ModelVideo) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *ModelVideo) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *ModelVideo) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *ModelVideo) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

// GetViewsCount returns the ViewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetViewsCount() int32 {
	if o == nil || IsNil(o.ViewsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ViewsCount.Get()
}

// GetViewsCountOk returns a tuple with the ViewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetViewsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewsCount.Get(), o.ViewsCount.IsSet()
}

// HasViewsCount returns a boolean if a field has been set.
func (o *ModelVideo) HasViewsCount() bool {
	if o != nil && o.ViewsCount.IsSet() {
		return true
	}

	return false
}

// SetViewsCount gets a reference to the given NullableInt32 and assigns it to the ViewsCount field.
func (o *ModelVideo) SetViewsCount(v int32) {
	o.ViewsCount.Set(&v)
}
// SetViewsCountNil sets the value for ViewsCount to be an explicit nil
func (o *ModelVideo) SetViewsCountNil() {
	o.ViewsCount.Set(nil)
}

// UnsetViewsCount ensures that no value is present for ViewsCount, not even an explicit nil
func (o *ModelVideo) UnsetViewsCount() {
	o.ViewsCount.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVideo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVideo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *ModelVideo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *ModelVideo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *ModelVideo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *ModelVideo) UnsetWidth() {
	o.Width.Unset()
}

func (o ModelVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitrate.IsSet() {
		toSerialize["bitrate"] = o.Bitrate.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsCompleted.IsSet() {
		toSerialize["is_completed"] = o.IsCompleted.Get()
	}
	if o.ThumbnailBigUrl.IsSet() {
		toSerialize["thumbnail_big_url"] = o.ThumbnailBigUrl.Get()
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl.Get()
	}
	if o.VideoUrl.IsSet() {
		toSerialize["video_url"] = o.VideoUrl.Get()
	}
	if o.ViewsCount.IsSet() {
		toSerialize["views_count"] = o.ViewsCount.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelVideo) UnmarshalJSON(data []byte) (err error) {
	varModelVideo := _ModelVideo{}

	err = json.Unmarshal(data, &varModelVideo)

	if err != nil {
		return err
	}

	*o = ModelVideo(varModelVideo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bitrate")
		delete(additionalProperties, "height")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_completed")
		delete(additionalProperties, "thumbnail_big_url")
		delete(additionalProperties, "thumbnail_url")
		delete(additionalProperties, "video_url")
		delete(additionalProperties, "views_count")
		delete(additionalProperties, "width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelVideo struct {
	value *ModelVideo
	isSet bool
}

func (v NullableModelVideo) Get() *ModelVideo {
	return v.value
}

func (v *NullableModelVideo) Set(val *ModelVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelVideo(val *ModelVideo) *NullableModelVideo {
	return &NullableModelVideo{value: val, isSet: true}
}

func (v NullableModelVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


