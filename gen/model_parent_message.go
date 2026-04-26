
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ParentMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentMessage{}

// ParentMessage struct for ParentMessage
type ParentMessage struct {
	Attachment NullableString `json:"attachment,omitempty"`
	AttachmentThumbnail NullableString `json:"attachment_thumbnail,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	FontSize NullableInt32 `json:"font_size,omitempty"`
	Gif NullableGifImage `json:"gif,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MessageType NullableMessageType `json:"message_type,omitempty"`
	RoomId NullableInt64 `json:"room_id,omitempty"`
	Sticker NullableSticker `json:"sticker,omitempty"`
	Text NullableString `json:"text,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	VideoThumbnailUrl NullableString `json:"video_thumbnail_url,omitempty"`
	VideoUrl NullableString `json:"video_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ParentMessage ParentMessage

// NewParentMessage instantiates a new ParentMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentMessage() *ParentMessage {
	this := ParentMessage{}
	return &this
}

// NewParentMessageWithDefaults instantiates a new ParentMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentMessageWithDefaults() *ParentMessage {
	this := ParentMessage{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetAttachment() string {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *ParentMessage) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableString and assigns it to the Attachment field.
func (o *ParentMessage) SetAttachment(v string) {
	o.Attachment.Set(&v)
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *ParentMessage) SetAttachmentNil() {
	o.Attachment.Set(nil)
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *ParentMessage) UnsetAttachment() {
	o.Attachment.Unset()
}

// GetAttachmentThumbnail returns the AttachmentThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetAttachmentThumbnail() string {
	if o == nil || IsNil(o.AttachmentThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentThumbnail.Get()
}

// GetAttachmentThumbnailOk returns a tuple with the AttachmentThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetAttachmentThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentThumbnail.Get(), o.AttachmentThumbnail.IsSet()
}

// HasAttachmentThumbnail returns a boolean if a field has been set.
func (o *ParentMessage) HasAttachmentThumbnail() bool {
	if o != nil && o.AttachmentThumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachmentThumbnail gets a reference to the given NullableString and assigns it to the AttachmentThumbnail field.
func (o *ParentMessage) SetAttachmentThumbnail(v string) {
	o.AttachmentThumbnail.Set(&v)
}
// SetAttachmentThumbnailNil sets the value for AttachmentThumbnail to be an explicit nil
func (o *ParentMessage) SetAttachmentThumbnailNil() {
	o.AttachmentThumbnail.Set(nil)
}

// UnsetAttachmentThumbnail ensures that no value is present for AttachmentThumbnail, not even an explicit nil
func (o *ParentMessage) UnsetAttachmentThumbnail() {
	o.AttachmentThumbnail.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ParentMessage) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *ParentMessage) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ParentMessage) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ParentMessage) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret int32
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetFontSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *ParentMessage) HasFontSize() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableInt32 and assigns it to the FontSize field.
func (o *ParentMessage) SetFontSize(v int32) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *ParentMessage) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *ParentMessage) UnsetFontSize() {
	o.FontSize.Unset()
}

// GetGif returns the Gif field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetGif() GifImage {
	if o == nil || IsNil(o.Gif.Get()) {
		var ret GifImage
		return ret
	}
	return *o.Gif.Get()
}

// GetGifOk returns a tuple with the Gif field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetGifOk() (*GifImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gif.Get(), o.Gif.IsSet()
}

// HasGif returns a boolean if a field has been set.
func (o *ParentMessage) HasGif() bool {
	if o != nil && o.Gif.IsSet() {
		return true
	}

	return false
}

// SetGif gets a reference to the given NullableGifImage and assigns it to the Gif field.
func (o *ParentMessage) SetGif(v GifImage) {
	o.Gif.Set(&v)
}
// SetGifNil sets the value for Gif to be an explicit nil
func (o *ParentMessage) SetGifNil() {
	o.Gif.Set(nil)
}

// UnsetGif ensures that no value is present for Gif, not even an explicit nil
func (o *ParentMessage) UnsetGif() {
	o.Gif.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ParentMessage) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ParentMessage) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ParentMessage) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ParentMessage) UnsetId() {
	o.Id.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetMessageType() MessageType {
	if o == nil || IsNil(o.MessageType.Get()) {
		var ret MessageType
		return ret
	}
	return *o.MessageType.Get()
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetMessageTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageType.Get(), o.MessageType.IsSet()
}

// HasMessageType returns a boolean if a field has been set.
func (o *ParentMessage) HasMessageType() bool {
	if o != nil && o.MessageType.IsSet() {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given NullableMessageType and assigns it to the MessageType field.
func (o *ParentMessage) SetMessageType(v MessageType) {
	o.MessageType.Set(&v)
}
// SetMessageTypeNil sets the value for MessageType to be an explicit nil
func (o *ParentMessage) SetMessageTypeNil() {
	o.MessageType.Set(nil)
}

// UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
func (o *ParentMessage) UnsetMessageType() {
	o.MessageType.Unset()
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetRoomId() int64 {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetRoomIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *ParentMessage) HasRoomId() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableInt64 and assigns it to the RoomId field.
func (o *ParentMessage) SetRoomId(v int64) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *ParentMessage) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *ParentMessage) UnsetRoomId() {
	o.RoomId.Unset()
}

// GetSticker returns the Sticker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetSticker() Sticker {
	if o == nil || IsNil(o.Sticker.Get()) {
		var ret Sticker
		return ret
	}
	return *o.Sticker.Get()
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetStickerOk() (*Sticker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sticker.Get(), o.Sticker.IsSet()
}

// HasSticker returns a boolean if a field has been set.
func (o *ParentMessage) HasSticker() bool {
	if o != nil && o.Sticker.IsSet() {
		return true
	}

	return false
}

// SetSticker gets a reference to the given NullableSticker and assigns it to the Sticker field.
func (o *ParentMessage) SetSticker(v Sticker) {
	o.Sticker.Set(&v)
}
// SetStickerNil sets the value for Sticker to be an explicit nil
func (o *ParentMessage) SetStickerNil() {
	o.Sticker.Set(nil)
}

// UnsetSticker ensures that no value is present for Sticker, not even an explicit nil
func (o *ParentMessage) UnsetSticker() {
	o.Sticker.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ParentMessage) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ParentMessage) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ParentMessage) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ParentMessage) UnsetText() {
	o.Text.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ParentMessage) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *ParentMessage) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ParentMessage) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ParentMessage) UnsetUserId() {
	o.UserId.Unset()
}

// GetVideoThumbnailUrl returns the VideoThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetVideoThumbnailUrl() string {
	if o == nil || IsNil(o.VideoThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoThumbnailUrl.Get()
}

// GetVideoThumbnailUrlOk returns a tuple with the VideoThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetVideoThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoThumbnailUrl.Get(), o.VideoThumbnailUrl.IsSet()
}

// HasVideoThumbnailUrl returns a boolean if a field has been set.
func (o *ParentMessage) HasVideoThumbnailUrl() bool {
	if o != nil && o.VideoThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoThumbnailUrl gets a reference to the given NullableString and assigns it to the VideoThumbnailUrl field.
func (o *ParentMessage) SetVideoThumbnailUrl(v string) {
	o.VideoThumbnailUrl.Set(&v)
}
// SetVideoThumbnailUrlNil sets the value for VideoThumbnailUrl to be an explicit nil
func (o *ParentMessage) SetVideoThumbnailUrlNil() {
	o.VideoThumbnailUrl.Set(nil)
}

// UnsetVideoThumbnailUrl ensures that no value is present for VideoThumbnailUrl, not even an explicit nil
func (o *ParentMessage) UnsetVideoThumbnailUrl() {
	o.VideoThumbnailUrl.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentMessage) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentMessage) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *ParentMessage) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *ParentMessage) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *ParentMessage) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *ParentMessage) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

func (o ParentMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachment.IsSet() {
		toSerialize["attachment"] = o.Attachment.Get()
	}
	if o.AttachmentThumbnail.IsSet() {
		toSerialize["attachment_thumbnail"] = o.AttachmentThumbnail.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.FontSize.IsSet() {
		toSerialize["font_size"] = o.FontSize.Get()
	}
	if o.Gif.IsSet() {
		toSerialize["gif"] = o.Gif.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MessageType.IsSet() {
		toSerialize["message_type"] = o.MessageType.Get()
	}
	if o.RoomId.IsSet() {
		toSerialize["room_id"] = o.RoomId.Get()
	}
	if o.Sticker.IsSet() {
		toSerialize["sticker"] = o.Sticker.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
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

func (o *ParentMessage) UnmarshalJSON(data []byte) (err error) {
	varParentMessage := _ParentMessage{}

	err = json.Unmarshal(data, &varParentMessage)

	if err != nil {
		return err
	}

	*o = ParentMessage(varParentMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attachment")
		delete(additionalProperties, "attachment_thumbnail")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "font_size")
		delete(additionalProperties, "gif")
		delete(additionalProperties, "id")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "room_id")
		delete(additionalProperties, "sticker")
		delete(additionalProperties, "text")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "video_thumbnail_url")
		delete(additionalProperties, "video_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableParentMessage struct {
	value *ParentMessage
	isSet bool
}

func (v NullableParentMessage) Get() *ParentMessage {
	return v.value
}

func (v *NullableParentMessage) Set(val *ParentMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableParentMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableParentMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentMessage(val *ParentMessage) *NullableParentMessage {
	return &NullableParentMessage{value: val, isSet: true}
}

func (v NullableParentMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


