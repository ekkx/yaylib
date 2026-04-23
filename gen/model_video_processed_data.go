
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the VideoProcessedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoProcessedData{}

// VideoProcessedData struct for VideoProcessedData
type VideoProcessedData struct {
	Id NullableInt64 `json:"id,omitempty"`
	VideoProcessed NullableBool `json:"video_processed,omitempty"`
	VideoThumbnailBigUrl NullableString `json:"video_thumbnail_big_url,omitempty"`
	VideoThumbnailUrl NullableString `json:"video_thumbnail_url,omitempty"`
	VideoUrl NullableString `json:"video_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VideoProcessedData VideoProcessedData

// NewVideoProcessedData instantiates a new VideoProcessedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoProcessedData() *VideoProcessedData {
	this := VideoProcessedData{}
	return &this
}

// NewVideoProcessedDataWithDefaults instantiates a new VideoProcessedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoProcessedDataWithDefaults() *VideoProcessedData {
	this := VideoProcessedData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoProcessedData) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoProcessedData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *VideoProcessedData) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *VideoProcessedData) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *VideoProcessedData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *VideoProcessedData) UnsetId() {
	o.Id.Unset()
}

// GetVideoProcessed returns the VideoProcessed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoProcessedData) GetVideoProcessed() bool {
	if o == nil || IsNil(o.VideoProcessed.Get()) {
		var ret bool
		return ret
	}
	return *o.VideoProcessed.Get()
}

// GetVideoProcessedOk returns a tuple with the VideoProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoProcessedData) GetVideoProcessedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoProcessed.Get(), o.VideoProcessed.IsSet()
}

// HasVideoProcessed returns a boolean if a field has been set.
func (o *VideoProcessedData) HasVideoProcessed() bool {
	if o != nil && o.VideoProcessed.IsSet() {
		return true
	}

	return false
}

// SetVideoProcessed gets a reference to the given NullableBool and assigns it to the VideoProcessed field.
func (o *VideoProcessedData) SetVideoProcessed(v bool) {
	o.VideoProcessed.Set(&v)
}
// SetVideoProcessedNil sets the value for VideoProcessed to be an explicit nil
func (o *VideoProcessedData) SetVideoProcessedNil() {
	o.VideoProcessed.Set(nil)
}

// UnsetVideoProcessed ensures that no value is present for VideoProcessed, not even an explicit nil
func (o *VideoProcessedData) UnsetVideoProcessed() {
	o.VideoProcessed.Unset()
}

// GetVideoThumbnailBigUrl returns the VideoThumbnailBigUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoProcessedData) GetVideoThumbnailBigUrl() string {
	if o == nil || IsNil(o.VideoThumbnailBigUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoThumbnailBigUrl.Get()
}

// GetVideoThumbnailBigUrlOk returns a tuple with the VideoThumbnailBigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoProcessedData) GetVideoThumbnailBigUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoThumbnailBigUrl.Get(), o.VideoThumbnailBigUrl.IsSet()
}

// HasVideoThumbnailBigUrl returns a boolean if a field has been set.
func (o *VideoProcessedData) HasVideoThumbnailBigUrl() bool {
	if o != nil && o.VideoThumbnailBigUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoThumbnailBigUrl gets a reference to the given NullableString and assigns it to the VideoThumbnailBigUrl field.
func (o *VideoProcessedData) SetVideoThumbnailBigUrl(v string) {
	o.VideoThumbnailBigUrl.Set(&v)
}
// SetVideoThumbnailBigUrlNil sets the value for VideoThumbnailBigUrl to be an explicit nil
func (o *VideoProcessedData) SetVideoThumbnailBigUrlNil() {
	o.VideoThumbnailBigUrl.Set(nil)
}

// UnsetVideoThumbnailBigUrl ensures that no value is present for VideoThumbnailBigUrl, not even an explicit nil
func (o *VideoProcessedData) UnsetVideoThumbnailBigUrl() {
	o.VideoThumbnailBigUrl.Unset()
}

// GetVideoThumbnailUrl returns the VideoThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoProcessedData) GetVideoThumbnailUrl() string {
	if o == nil || IsNil(o.VideoThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoThumbnailUrl.Get()
}

// GetVideoThumbnailUrlOk returns a tuple with the VideoThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoProcessedData) GetVideoThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoThumbnailUrl.Get(), o.VideoThumbnailUrl.IsSet()
}

// HasVideoThumbnailUrl returns a boolean if a field has been set.
func (o *VideoProcessedData) HasVideoThumbnailUrl() bool {
	if o != nil && o.VideoThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoThumbnailUrl gets a reference to the given NullableString and assigns it to the VideoThumbnailUrl field.
func (o *VideoProcessedData) SetVideoThumbnailUrl(v string) {
	o.VideoThumbnailUrl.Set(&v)
}
// SetVideoThumbnailUrlNil sets the value for VideoThumbnailUrl to be an explicit nil
func (o *VideoProcessedData) SetVideoThumbnailUrlNil() {
	o.VideoThumbnailUrl.Set(nil)
}

// UnsetVideoThumbnailUrl ensures that no value is present for VideoThumbnailUrl, not even an explicit nil
func (o *VideoProcessedData) UnsetVideoThumbnailUrl() {
	o.VideoThumbnailUrl.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoProcessedData) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoProcessedData) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *VideoProcessedData) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *VideoProcessedData) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *VideoProcessedData) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *VideoProcessedData) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

func (o VideoProcessedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoProcessedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.VideoProcessed.IsSet() {
		toSerialize["video_processed"] = o.VideoProcessed.Get()
	}
	if o.VideoThumbnailBigUrl.IsSet() {
		toSerialize["video_thumbnail_big_url"] = o.VideoThumbnailBigUrl.Get()
	}
	if o.VideoThumbnailUrl.IsSet() {
		toSerialize["video_thumbnail_url"] = o.VideoThumbnailUrl.Get()
	}
	if o.VideoUrl.IsSet() {
		toSerialize["video_url"] = o.VideoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VideoProcessedData) UnmarshalJSON(data []byte) (err error) {
	varVideoProcessedData := _VideoProcessedData{}

	err = json.Unmarshal(data, &varVideoProcessedData)

	if err != nil {
		return err
	}

	*o = VideoProcessedData(varVideoProcessedData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "video_processed")
		delete(additionalProperties, "video_thumbnail_big_url")
		delete(additionalProperties, "video_thumbnail_url")
		delete(additionalProperties, "video_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVideoProcessedData struct {
	value *VideoProcessedData
	isSet bool
}

func (v NullableVideoProcessedData) Get() *VideoProcessedData {
	return v.value
}

func (v *NullableVideoProcessedData) Set(val *VideoProcessedData) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoProcessedData) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoProcessedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoProcessedData(val *VideoProcessedData) *NullableVideoProcessedData {
	return &NullableVideoProcessedData{value: val, isSet: true}
}

func (v NullableVideoProcessedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoProcessedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


