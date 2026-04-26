
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmMessage{}

// RealmMessage struct for RealmMessage
type RealmMessage struct {
	Attachment NullableString `json:"attachment,omitempty"`
	AttachmentAndroid NullableString `json:"attachment_android,omitempty"`
	AttachmentReadCount NullableInt32 `json:"attachment_read_count,omitempty"`
	AttachmentThumbnail NullableString `json:"attachment_thumbnail,omitempty"`
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	FontSize NullableInt32 `json:"font_size,omitempty"`
	Gif NullableGifImage `json:"gif,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Invitation NullableChatInvitation `json:"invitation,omitempty"`
	IsError NullableBool `json:"is_error,omitempty"`
	IsSent NullableBool `json:"is_sent,omitempty"`
	MessageType NullableMessageType `json:"message_type,omitempty"`
	Parent NullableParentMessage `json:"parent,omitempty"`
	RefreshRetryCount NullableInt32 `json:"refresh_retry_count,omitempty"`
	RoomId NullableInt64 `json:"room_id,omitempty"`
	Sticker NullableSticker `json:"sticker,omitempty"`
	Text NullableString `json:"text,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	VideoProcessed NullableBool `json:"video_processed,omitempty"`
	VideoThumbnailUrl NullableString `json:"video_thumbnail_url,omitempty"`
	VideoUrl NullableString `json:"video_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmMessage RealmMessage

// NewRealmMessage instantiates a new RealmMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmMessage() *RealmMessage {
	this := RealmMessage{}
	return &this
}

// NewRealmMessageWithDefaults instantiates a new RealmMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmMessageWithDefaults() *RealmMessage {
	this := RealmMessage{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetAttachment() string {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *RealmMessage) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableString and assigns it to the Attachment field.
func (o *RealmMessage) SetAttachment(v string) {
	o.Attachment.Set(&v)
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *RealmMessage) SetAttachmentNil() {
	o.Attachment.Set(nil)
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *RealmMessage) UnsetAttachment() {
	o.Attachment.Unset()
}

// GetAttachmentAndroid returns the AttachmentAndroid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetAttachmentAndroid() string {
	if o == nil || IsNil(o.AttachmentAndroid.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentAndroid.Get()
}

// GetAttachmentAndroidOk returns a tuple with the AttachmentAndroid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetAttachmentAndroidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentAndroid.Get(), o.AttachmentAndroid.IsSet()
}

// HasAttachmentAndroid returns a boolean if a field has been set.
func (o *RealmMessage) HasAttachmentAndroid() bool {
	if o != nil && o.AttachmentAndroid.IsSet() {
		return true
	}

	return false
}

// SetAttachmentAndroid gets a reference to the given NullableString and assigns it to the AttachmentAndroid field.
func (o *RealmMessage) SetAttachmentAndroid(v string) {
	o.AttachmentAndroid.Set(&v)
}
// SetAttachmentAndroidNil sets the value for AttachmentAndroid to be an explicit nil
func (o *RealmMessage) SetAttachmentAndroidNil() {
	o.AttachmentAndroid.Set(nil)
}

// UnsetAttachmentAndroid ensures that no value is present for AttachmentAndroid, not even an explicit nil
func (o *RealmMessage) UnsetAttachmentAndroid() {
	o.AttachmentAndroid.Unset()
}

// GetAttachmentReadCount returns the AttachmentReadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetAttachmentReadCount() int32 {
	if o == nil || IsNil(o.AttachmentReadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AttachmentReadCount.Get()
}

// GetAttachmentReadCountOk returns a tuple with the AttachmentReadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetAttachmentReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentReadCount.Get(), o.AttachmentReadCount.IsSet()
}

// HasAttachmentReadCount returns a boolean if a field has been set.
func (o *RealmMessage) HasAttachmentReadCount() bool {
	if o != nil && o.AttachmentReadCount.IsSet() {
		return true
	}

	return false
}

// SetAttachmentReadCount gets a reference to the given NullableInt32 and assigns it to the AttachmentReadCount field.
func (o *RealmMessage) SetAttachmentReadCount(v int32) {
	o.AttachmentReadCount.Set(&v)
}
// SetAttachmentReadCountNil sets the value for AttachmentReadCount to be an explicit nil
func (o *RealmMessage) SetAttachmentReadCountNil() {
	o.AttachmentReadCount.Set(nil)
}

// UnsetAttachmentReadCount ensures that no value is present for AttachmentReadCount, not even an explicit nil
func (o *RealmMessage) UnsetAttachmentReadCount() {
	o.AttachmentReadCount.Unset()
}

// GetAttachmentThumbnail returns the AttachmentThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetAttachmentThumbnail() string {
	if o == nil || IsNil(o.AttachmentThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentThumbnail.Get()
}

// GetAttachmentThumbnailOk returns a tuple with the AttachmentThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetAttachmentThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentThumbnail.Get(), o.AttachmentThumbnail.IsSet()
}

// HasAttachmentThumbnail returns a boolean if a field has been set.
func (o *RealmMessage) HasAttachmentThumbnail() bool {
	if o != nil && o.AttachmentThumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachmentThumbnail gets a reference to the given NullableString and assigns it to the AttachmentThumbnail field.
func (o *RealmMessage) SetAttachmentThumbnail(v string) {
	o.AttachmentThumbnail.Set(&v)
}
// SetAttachmentThumbnailNil sets the value for AttachmentThumbnail to be an explicit nil
func (o *RealmMessage) SetAttachmentThumbnailNil() {
	o.AttachmentThumbnail.Set(nil)
}

// UnsetAttachmentThumbnail ensures that no value is present for AttachmentThumbnail, not even an explicit nil
func (o *RealmMessage) UnsetAttachmentThumbnail() {
	o.AttachmentThumbnail.Unset()
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *RealmMessage) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *RealmMessage) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *RealmMessage) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *RealmMessage) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RealmMessage) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *RealmMessage) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RealmMessage) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RealmMessage) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret int32
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetFontSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *RealmMessage) HasFontSize() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableInt32 and assigns it to the FontSize field.
func (o *RealmMessage) SetFontSize(v int32) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *RealmMessage) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *RealmMessage) UnsetFontSize() {
	o.FontSize.Unset()
}

// GetGif returns the Gif field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetGif() GifImage {
	if o == nil || IsNil(o.Gif.Get()) {
		var ret GifImage
		return ret
	}
	return *o.Gif.Get()
}

// GetGifOk returns a tuple with the Gif field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetGifOk() (*GifImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gif.Get(), o.Gif.IsSet()
}

// HasGif returns a boolean if a field has been set.
func (o *RealmMessage) HasGif() bool {
	if o != nil && o.Gif.IsSet() {
		return true
	}

	return false
}

// SetGif gets a reference to the given NullableGifImage and assigns it to the Gif field.
func (o *RealmMessage) SetGif(v GifImage) {
	o.Gif.Set(&v)
}
// SetGifNil sets the value for Gif to be an explicit nil
func (o *RealmMessage) SetGifNil() {
	o.Gif.Set(nil)
}

// UnsetGif ensures that no value is present for Gif, not even an explicit nil
func (o *RealmMessage) UnsetGif() {
	o.Gif.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmMessage) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmMessage) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmMessage) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmMessage) UnsetId() {
	o.Id.Unset()
}

// GetInvitation returns the Invitation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetInvitation() ChatInvitation {
	if o == nil || IsNil(o.Invitation.Get()) {
		var ret ChatInvitation
		return ret
	}
	return *o.Invitation.Get()
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetInvitationOk() (*ChatInvitation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invitation.Get(), o.Invitation.IsSet()
}

// HasInvitation returns a boolean if a field has been set.
func (o *RealmMessage) HasInvitation() bool {
	if o != nil && o.Invitation.IsSet() {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given NullableChatInvitation and assigns it to the Invitation field.
func (o *RealmMessage) SetInvitation(v ChatInvitation) {
	o.Invitation.Set(&v)
}
// SetInvitationNil sets the value for Invitation to be an explicit nil
func (o *RealmMessage) SetInvitationNil() {
	o.Invitation.Set(nil)
}

// UnsetInvitation ensures that no value is present for Invitation, not even an explicit nil
func (o *RealmMessage) UnsetInvitation() {
	o.Invitation.Unset()
}

// GetIsError returns the IsError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetIsError() bool {
	if o == nil || IsNil(o.IsError.Get()) {
		var ret bool
		return ret
	}
	return *o.IsError.Get()
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetIsErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsError.Get(), o.IsError.IsSet()
}

// HasIsError returns a boolean if a field has been set.
func (o *RealmMessage) HasIsError() bool {
	if o != nil && o.IsError.IsSet() {
		return true
	}

	return false
}

// SetIsError gets a reference to the given NullableBool and assigns it to the IsError field.
func (o *RealmMessage) SetIsError(v bool) {
	o.IsError.Set(&v)
}
// SetIsErrorNil sets the value for IsError to be an explicit nil
func (o *RealmMessage) SetIsErrorNil() {
	o.IsError.Set(nil)
}

// UnsetIsError ensures that no value is present for IsError, not even an explicit nil
func (o *RealmMessage) UnsetIsError() {
	o.IsError.Unset()
}

// GetIsSent returns the IsSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetIsSent() bool {
	if o == nil || IsNil(o.IsSent.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSent.Get()
}

// GetIsSentOk returns a tuple with the IsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetIsSentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSent.Get(), o.IsSent.IsSet()
}

// HasIsSent returns a boolean if a field has been set.
func (o *RealmMessage) HasIsSent() bool {
	if o != nil && o.IsSent.IsSet() {
		return true
	}

	return false
}

// SetIsSent gets a reference to the given NullableBool and assigns it to the IsSent field.
func (o *RealmMessage) SetIsSent(v bool) {
	o.IsSent.Set(&v)
}
// SetIsSentNil sets the value for IsSent to be an explicit nil
func (o *RealmMessage) SetIsSentNil() {
	o.IsSent.Set(nil)
}

// UnsetIsSent ensures that no value is present for IsSent, not even an explicit nil
func (o *RealmMessage) UnsetIsSent() {
	o.IsSent.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetMessageType() MessageType {
	if o == nil || IsNil(o.MessageType.Get()) {
		var ret MessageType
		return ret
	}
	return *o.MessageType.Get()
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetMessageTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageType.Get(), o.MessageType.IsSet()
}

// HasMessageType returns a boolean if a field has been set.
func (o *RealmMessage) HasMessageType() bool {
	if o != nil && o.MessageType.IsSet() {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given NullableMessageType and assigns it to the MessageType field.
func (o *RealmMessage) SetMessageType(v MessageType) {
	o.MessageType.Set(&v)
}
// SetMessageTypeNil sets the value for MessageType to be an explicit nil
func (o *RealmMessage) SetMessageTypeNil() {
	o.MessageType.Set(nil)
}

// UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
func (o *RealmMessage) UnsetMessageType() {
	o.MessageType.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetParent() ParentMessage {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret ParentMessage
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetParentOk() (*ParentMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *RealmMessage) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableParentMessage and assigns it to the Parent field.
func (o *RealmMessage) SetParent(v ParentMessage) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *RealmMessage) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *RealmMessage) UnsetParent() {
	o.Parent.Unset()
}

// GetRefreshRetryCount returns the RefreshRetryCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetRefreshRetryCount() int32 {
	if o == nil || IsNil(o.RefreshRetryCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RefreshRetryCount.Get()
}

// GetRefreshRetryCountOk returns a tuple with the RefreshRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetRefreshRetryCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshRetryCount.Get(), o.RefreshRetryCount.IsSet()
}

// HasRefreshRetryCount returns a boolean if a field has been set.
func (o *RealmMessage) HasRefreshRetryCount() bool {
	if o != nil && o.RefreshRetryCount.IsSet() {
		return true
	}

	return false
}

// SetRefreshRetryCount gets a reference to the given NullableInt32 and assigns it to the RefreshRetryCount field.
func (o *RealmMessage) SetRefreshRetryCount(v int32) {
	o.RefreshRetryCount.Set(&v)
}
// SetRefreshRetryCountNil sets the value for RefreshRetryCount to be an explicit nil
func (o *RealmMessage) SetRefreshRetryCountNil() {
	o.RefreshRetryCount.Set(nil)
}

// UnsetRefreshRetryCount ensures that no value is present for RefreshRetryCount, not even an explicit nil
func (o *RealmMessage) UnsetRefreshRetryCount() {
	o.RefreshRetryCount.Unset()
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetRoomId() int64 {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetRoomIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *RealmMessage) HasRoomId() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableInt64 and assigns it to the RoomId field.
func (o *RealmMessage) SetRoomId(v int64) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *RealmMessage) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *RealmMessage) UnsetRoomId() {
	o.RoomId.Unset()
}

// GetSticker returns the Sticker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetSticker() Sticker {
	if o == nil || IsNil(o.Sticker.Get()) {
		var ret Sticker
		return ret
	}
	return *o.Sticker.Get()
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetStickerOk() (*Sticker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sticker.Get(), o.Sticker.IsSet()
}

// HasSticker returns a boolean if a field has been set.
func (o *RealmMessage) HasSticker() bool {
	if o != nil && o.Sticker.IsSet() {
		return true
	}

	return false
}

// SetSticker gets a reference to the given NullableSticker and assigns it to the Sticker field.
func (o *RealmMessage) SetSticker(v Sticker) {
	o.Sticker.Set(&v)
}
// SetStickerNil sets the value for Sticker to be an explicit nil
func (o *RealmMessage) SetStickerNil() {
	o.Sticker.Set(nil)
}

// UnsetSticker ensures that no value is present for Sticker, not even an explicit nil
func (o *RealmMessage) UnsetSticker() {
	o.Sticker.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *RealmMessage) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *RealmMessage) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *RealmMessage) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *RealmMessage) UnsetText() {
	o.Text.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *RealmMessage) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *RealmMessage) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *RealmMessage) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *RealmMessage) UnsetUserId() {
	o.UserId.Unset()
}

// GetVideoProcessed returns the VideoProcessed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetVideoProcessed() bool {
	if o == nil || IsNil(o.VideoProcessed.Get()) {
		var ret bool
		return ret
	}
	return *o.VideoProcessed.Get()
}

// GetVideoProcessedOk returns a tuple with the VideoProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetVideoProcessedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoProcessed.Get(), o.VideoProcessed.IsSet()
}

// HasVideoProcessed returns a boolean if a field has been set.
func (o *RealmMessage) HasVideoProcessed() bool {
	if o != nil && o.VideoProcessed.IsSet() {
		return true
	}

	return false
}

// SetVideoProcessed gets a reference to the given NullableBool and assigns it to the VideoProcessed field.
func (o *RealmMessage) SetVideoProcessed(v bool) {
	o.VideoProcessed.Set(&v)
}
// SetVideoProcessedNil sets the value for VideoProcessed to be an explicit nil
func (o *RealmMessage) SetVideoProcessedNil() {
	o.VideoProcessed.Set(nil)
}

// UnsetVideoProcessed ensures that no value is present for VideoProcessed, not even an explicit nil
func (o *RealmMessage) UnsetVideoProcessed() {
	o.VideoProcessed.Unset()
}

// GetVideoThumbnailUrl returns the VideoThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetVideoThumbnailUrl() string {
	if o == nil || IsNil(o.VideoThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoThumbnailUrl.Get()
}

// GetVideoThumbnailUrlOk returns a tuple with the VideoThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetVideoThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoThumbnailUrl.Get(), o.VideoThumbnailUrl.IsSet()
}

// HasVideoThumbnailUrl returns a boolean if a field has been set.
func (o *RealmMessage) HasVideoThumbnailUrl() bool {
	if o != nil && o.VideoThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoThumbnailUrl gets a reference to the given NullableString and assigns it to the VideoThumbnailUrl field.
func (o *RealmMessage) SetVideoThumbnailUrl(v string) {
	o.VideoThumbnailUrl.Set(&v)
}
// SetVideoThumbnailUrlNil sets the value for VideoThumbnailUrl to be an explicit nil
func (o *RealmMessage) SetVideoThumbnailUrlNil() {
	o.VideoThumbnailUrl.Set(nil)
}

// UnsetVideoThumbnailUrl ensures that no value is present for VideoThumbnailUrl, not even an explicit nil
func (o *RealmMessage) UnsetVideoThumbnailUrl() {
	o.VideoThumbnailUrl.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmMessage) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmMessage) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *RealmMessage) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *RealmMessage) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *RealmMessage) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *RealmMessage) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

func (o RealmMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachment.IsSet() {
		toSerialize["attachment"] = o.Attachment.Get()
	}
	if o.AttachmentAndroid.IsSet() {
		toSerialize["attachment_android"] = o.AttachmentAndroid.Get()
	}
	if o.AttachmentReadCount.IsSet() {
		toSerialize["attachment_read_count"] = o.AttachmentReadCount.Get()
	}
	if o.AttachmentThumbnail.IsSet() {
		toSerialize["attachment_thumbnail"] = o.AttachmentThumbnail.Get()
	}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
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
	if o.Invitation.IsSet() {
		toSerialize["invitation"] = o.Invitation.Get()
	}
	if o.IsError.IsSet() {
		toSerialize["is_error"] = o.IsError.Get()
	}
	if o.IsSent.IsSet() {
		toSerialize["is_sent"] = o.IsSent.Get()
	}
	if o.MessageType.IsSet() {
		toSerialize["message_type"] = o.MessageType.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.RefreshRetryCount.IsSet() {
		toSerialize["refresh_retry_count"] = o.RefreshRetryCount.Get()
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
	if o.VideoProcessed.IsSet() {
		toSerialize["video_processed"] = o.VideoProcessed.Get()
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

func (o *RealmMessage) UnmarshalJSON(data []byte) (err error) {
	varRealmMessage := _RealmMessage{}

	err = json.Unmarshal(data, &varRealmMessage)

	if err != nil {
		return err
	}

	*o = RealmMessage(varRealmMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attachment")
		delete(additionalProperties, "attachment_android")
		delete(additionalProperties, "attachment_read_count")
		delete(additionalProperties, "attachment_thumbnail")
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "font_size")
		delete(additionalProperties, "gif")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invitation")
		delete(additionalProperties, "is_error")
		delete(additionalProperties, "is_sent")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "refresh_retry_count")
		delete(additionalProperties, "room_id")
		delete(additionalProperties, "sticker")
		delete(additionalProperties, "text")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "video_processed")
		delete(additionalProperties, "video_thumbnail_url")
		delete(additionalProperties, "video_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmMessage struct {
	value *RealmMessage
	isSet bool
}

func (v NullableRealmMessage) Get() *RealmMessage {
	return v.value
}

func (v *NullableRealmMessage) Set(val *RealmMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmMessage(val *RealmMessage) *NullableRealmMessage {
	return &NullableRealmMessage{value: val, isSet: true}
}

func (v NullableRealmMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


