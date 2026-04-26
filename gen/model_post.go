
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Post type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Post{}

// Post struct for Post
type Post struct {
	Attachment NullableString `json:"attachment,omitempty"`
	Attachment2 NullableString `json:"attachment_2,omitempty"`
	Attachment2Thumbnail NullableString `json:"attachment_2_thumbnail,omitempty"`
	Attachment3 NullableString `json:"attachment_3,omitempty"`
	Attachment3Thumbnail NullableString `json:"attachment_3_thumbnail,omitempty"`
	Attachment4 NullableString `json:"attachment_4,omitempty"`
	Attachment4Thumbnail NullableString `json:"attachment_4_thumbnail,omitempty"`
	Attachment5 NullableString `json:"attachment_5,omitempty"`
	Attachment5Thumbnail NullableString `json:"attachment_5_thumbnail,omitempty"`
	Attachment6 NullableString `json:"attachment_6,omitempty"`
	Attachment6Thumbnail NullableString `json:"attachment_6_thumbnail,omitempty"`
	Attachment7 NullableString `json:"attachment_7,omitempty"`
	Attachment7Thumbnail NullableString `json:"attachment_7_thumbnail,omitempty"`
	Attachment8 NullableString `json:"attachment_8,omitempty"`
	Attachment8Thumbnail NullableString `json:"attachment_8_thumbnail,omitempty"`
	Attachment9 NullableString `json:"attachment_9,omitempty"`
	Attachment9Thumbnail NullableString `json:"attachment_9_thumbnail,omitempty"`
	AttachmentThumbnail NullableString `json:"attachment_thumbnail,omitempty"`
	Color NullableInt32 `json:"color,omitempty"`
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	ConversationId NullableInt64 `json:"conversation_id,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	EditedAt NullableInt64 `json:"edited_at,omitempty"`
	FontSize NullableInt32 `json:"font_size,omitempty"`
	GiftsCount []GiftCount `json:"gifts_count,omitempty"`
	Group NullableGroup `json:"group,omitempty"`
	GroupId NullableInt64 `json:"group_id,omitempty"`
	Highlighted NullableBool `json:"highlighted,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InReplyTo NullableInt64 `json:"in_reply_to,omitempty"`
	InReplyToPost NullablePost `json:"in_reply_to_post,omitempty"`
	InReplyToPostCount NullableInt32 `json:"in_reply_to_post_count,omitempty"`
	IsFailToSend NullableBool `json:"is_fail_to_send,omitempty"`
	IsMuted NullableBool `json:"is_muted,omitempty"`
	Liked NullableBool `json:"liked,omitempty"`
	Likers []RealmUser `json:"likers,omitempty"`
	LikersCount NullableInt32 `json:"likers_count,omitempty"`
	LikesCount NullableInt32 `json:"likes_count,omitempty"`
	Mentions []RealmUser `json:"mentions,omitempty"`
	MessageTags []MessageTag `json:"message_tags,omitempty"`
	PostType NullablePostType `json:"post_type,omitempty"`
	ReportedCount NullableInt32 `json:"reported_count,omitempty"`
	Repostable NullableBool `json:"repostable,omitempty"`
	Reposted NullableBool `json:"reposted,omitempty"`
	RepostsCount NullableInt32 `json:"reposts_count,omitempty"`
	Shareable NullableShareable `json:"shareable,omitempty"`
	SharedThread NullableThreadInfo `json:"shared_thread,omitempty"`
	SharedUrl NullableSharedUrl `json:"shared_url,omitempty"`
	Survey NullableSurvey `json:"survey,omitempty"`
	Tag NullableString `json:"tag,omitempty"`
	Text NullableString `json:"text,omitempty"`
	Thread NullableThreadInfo `json:"thread,omitempty"`
	ThreadId NullableInt64 `json:"thread_id,omitempty"`
	TotalGiftsCount NullableInt32 `json:"total_gifts_count,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	Videos []Video `json:"videos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Post Post

// NewPost instantiates a new Post object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPost() *Post {
	this := Post{}
	return &this
}

// NewPostWithDefaults instantiates a new Post object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostWithDefaults() *Post {
	this := Post{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment() string {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *Post) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableString and assigns it to the Attachment field.
func (o *Post) SetAttachment(v string) {
	o.Attachment.Set(&v)
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *Post) SetAttachmentNil() {
	o.Attachment.Set(nil)
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *Post) UnsetAttachment() {
	o.Attachment.Unset()
}

// GetAttachment2 returns the Attachment2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment2() string {
	if o == nil || IsNil(o.Attachment2.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment2.Get()
}

// GetAttachment2Ok returns a tuple with the Attachment2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment2.Get(), o.Attachment2.IsSet()
}

// HasAttachment2 returns a boolean if a field has been set.
func (o *Post) HasAttachment2() bool {
	if o != nil && o.Attachment2.IsSet() {
		return true
	}

	return false
}

// SetAttachment2 gets a reference to the given NullableString and assigns it to the Attachment2 field.
func (o *Post) SetAttachment2(v string) {
	o.Attachment2.Set(&v)
}
// SetAttachment2Nil sets the value for Attachment2 to be an explicit nil
func (o *Post) SetAttachment2Nil() {
	o.Attachment2.Set(nil)
}

// UnsetAttachment2 ensures that no value is present for Attachment2, not even an explicit nil
func (o *Post) UnsetAttachment2() {
	o.Attachment2.Unset()
}

// GetAttachment2Thumbnail returns the Attachment2Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment2Thumbnail() string {
	if o == nil || IsNil(o.Attachment2Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment2Thumbnail.Get()
}

// GetAttachment2ThumbnailOk returns a tuple with the Attachment2Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment2ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment2Thumbnail.Get(), o.Attachment2Thumbnail.IsSet()
}

// HasAttachment2Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment2Thumbnail() bool {
	if o != nil && o.Attachment2Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment2Thumbnail gets a reference to the given NullableString and assigns it to the Attachment2Thumbnail field.
func (o *Post) SetAttachment2Thumbnail(v string) {
	o.Attachment2Thumbnail.Set(&v)
}
// SetAttachment2ThumbnailNil sets the value for Attachment2Thumbnail to be an explicit nil
func (o *Post) SetAttachment2ThumbnailNil() {
	o.Attachment2Thumbnail.Set(nil)
}

// UnsetAttachment2Thumbnail ensures that no value is present for Attachment2Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment2Thumbnail() {
	o.Attachment2Thumbnail.Unset()
}

// GetAttachment3 returns the Attachment3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment3() string {
	if o == nil || IsNil(o.Attachment3.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment3.Get()
}

// GetAttachment3Ok returns a tuple with the Attachment3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment3.Get(), o.Attachment3.IsSet()
}

// HasAttachment3 returns a boolean if a field has been set.
func (o *Post) HasAttachment3() bool {
	if o != nil && o.Attachment3.IsSet() {
		return true
	}

	return false
}

// SetAttachment3 gets a reference to the given NullableString and assigns it to the Attachment3 field.
func (o *Post) SetAttachment3(v string) {
	o.Attachment3.Set(&v)
}
// SetAttachment3Nil sets the value for Attachment3 to be an explicit nil
func (o *Post) SetAttachment3Nil() {
	o.Attachment3.Set(nil)
}

// UnsetAttachment3 ensures that no value is present for Attachment3, not even an explicit nil
func (o *Post) UnsetAttachment3() {
	o.Attachment3.Unset()
}

// GetAttachment3Thumbnail returns the Attachment3Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment3Thumbnail() string {
	if o == nil || IsNil(o.Attachment3Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment3Thumbnail.Get()
}

// GetAttachment3ThumbnailOk returns a tuple with the Attachment3Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment3ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment3Thumbnail.Get(), o.Attachment3Thumbnail.IsSet()
}

// HasAttachment3Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment3Thumbnail() bool {
	if o != nil && o.Attachment3Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment3Thumbnail gets a reference to the given NullableString and assigns it to the Attachment3Thumbnail field.
func (o *Post) SetAttachment3Thumbnail(v string) {
	o.Attachment3Thumbnail.Set(&v)
}
// SetAttachment3ThumbnailNil sets the value for Attachment3Thumbnail to be an explicit nil
func (o *Post) SetAttachment3ThumbnailNil() {
	o.Attachment3Thumbnail.Set(nil)
}

// UnsetAttachment3Thumbnail ensures that no value is present for Attachment3Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment3Thumbnail() {
	o.Attachment3Thumbnail.Unset()
}

// GetAttachment4 returns the Attachment4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment4() string {
	if o == nil || IsNil(o.Attachment4.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment4.Get()
}

// GetAttachment4Ok returns a tuple with the Attachment4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment4.Get(), o.Attachment4.IsSet()
}

// HasAttachment4 returns a boolean if a field has been set.
func (o *Post) HasAttachment4() bool {
	if o != nil && o.Attachment4.IsSet() {
		return true
	}

	return false
}

// SetAttachment4 gets a reference to the given NullableString and assigns it to the Attachment4 field.
func (o *Post) SetAttachment4(v string) {
	o.Attachment4.Set(&v)
}
// SetAttachment4Nil sets the value for Attachment4 to be an explicit nil
func (o *Post) SetAttachment4Nil() {
	o.Attachment4.Set(nil)
}

// UnsetAttachment4 ensures that no value is present for Attachment4, not even an explicit nil
func (o *Post) UnsetAttachment4() {
	o.Attachment4.Unset()
}

// GetAttachment4Thumbnail returns the Attachment4Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment4Thumbnail() string {
	if o == nil || IsNil(o.Attachment4Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment4Thumbnail.Get()
}

// GetAttachment4ThumbnailOk returns a tuple with the Attachment4Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment4ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment4Thumbnail.Get(), o.Attachment4Thumbnail.IsSet()
}

// HasAttachment4Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment4Thumbnail() bool {
	if o != nil && o.Attachment4Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment4Thumbnail gets a reference to the given NullableString and assigns it to the Attachment4Thumbnail field.
func (o *Post) SetAttachment4Thumbnail(v string) {
	o.Attachment4Thumbnail.Set(&v)
}
// SetAttachment4ThumbnailNil sets the value for Attachment4Thumbnail to be an explicit nil
func (o *Post) SetAttachment4ThumbnailNil() {
	o.Attachment4Thumbnail.Set(nil)
}

// UnsetAttachment4Thumbnail ensures that no value is present for Attachment4Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment4Thumbnail() {
	o.Attachment4Thumbnail.Unset()
}

// GetAttachment5 returns the Attachment5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment5() string {
	if o == nil || IsNil(o.Attachment5.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment5.Get()
}

// GetAttachment5Ok returns a tuple with the Attachment5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment5.Get(), o.Attachment5.IsSet()
}

// HasAttachment5 returns a boolean if a field has been set.
func (o *Post) HasAttachment5() bool {
	if o != nil && o.Attachment5.IsSet() {
		return true
	}

	return false
}

// SetAttachment5 gets a reference to the given NullableString and assigns it to the Attachment5 field.
func (o *Post) SetAttachment5(v string) {
	o.Attachment5.Set(&v)
}
// SetAttachment5Nil sets the value for Attachment5 to be an explicit nil
func (o *Post) SetAttachment5Nil() {
	o.Attachment5.Set(nil)
}

// UnsetAttachment5 ensures that no value is present for Attachment5, not even an explicit nil
func (o *Post) UnsetAttachment5() {
	o.Attachment5.Unset()
}

// GetAttachment5Thumbnail returns the Attachment5Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment5Thumbnail() string {
	if o == nil || IsNil(o.Attachment5Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment5Thumbnail.Get()
}

// GetAttachment5ThumbnailOk returns a tuple with the Attachment5Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment5ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment5Thumbnail.Get(), o.Attachment5Thumbnail.IsSet()
}

// HasAttachment5Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment5Thumbnail() bool {
	if o != nil && o.Attachment5Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment5Thumbnail gets a reference to the given NullableString and assigns it to the Attachment5Thumbnail field.
func (o *Post) SetAttachment5Thumbnail(v string) {
	o.Attachment5Thumbnail.Set(&v)
}
// SetAttachment5ThumbnailNil sets the value for Attachment5Thumbnail to be an explicit nil
func (o *Post) SetAttachment5ThumbnailNil() {
	o.Attachment5Thumbnail.Set(nil)
}

// UnsetAttachment5Thumbnail ensures that no value is present for Attachment5Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment5Thumbnail() {
	o.Attachment5Thumbnail.Unset()
}

// GetAttachment6 returns the Attachment6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment6() string {
	if o == nil || IsNil(o.Attachment6.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment6.Get()
}

// GetAttachment6Ok returns a tuple with the Attachment6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment6.Get(), o.Attachment6.IsSet()
}

// HasAttachment6 returns a boolean if a field has been set.
func (o *Post) HasAttachment6() bool {
	if o != nil && o.Attachment6.IsSet() {
		return true
	}

	return false
}

// SetAttachment6 gets a reference to the given NullableString and assigns it to the Attachment6 field.
func (o *Post) SetAttachment6(v string) {
	o.Attachment6.Set(&v)
}
// SetAttachment6Nil sets the value for Attachment6 to be an explicit nil
func (o *Post) SetAttachment6Nil() {
	o.Attachment6.Set(nil)
}

// UnsetAttachment6 ensures that no value is present for Attachment6, not even an explicit nil
func (o *Post) UnsetAttachment6() {
	o.Attachment6.Unset()
}

// GetAttachment6Thumbnail returns the Attachment6Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment6Thumbnail() string {
	if o == nil || IsNil(o.Attachment6Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment6Thumbnail.Get()
}

// GetAttachment6ThumbnailOk returns a tuple with the Attachment6Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment6ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment6Thumbnail.Get(), o.Attachment6Thumbnail.IsSet()
}

// HasAttachment6Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment6Thumbnail() bool {
	if o != nil && o.Attachment6Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment6Thumbnail gets a reference to the given NullableString and assigns it to the Attachment6Thumbnail field.
func (o *Post) SetAttachment6Thumbnail(v string) {
	o.Attachment6Thumbnail.Set(&v)
}
// SetAttachment6ThumbnailNil sets the value for Attachment6Thumbnail to be an explicit nil
func (o *Post) SetAttachment6ThumbnailNil() {
	o.Attachment6Thumbnail.Set(nil)
}

// UnsetAttachment6Thumbnail ensures that no value is present for Attachment6Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment6Thumbnail() {
	o.Attachment6Thumbnail.Unset()
}

// GetAttachment7 returns the Attachment7 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment7() string {
	if o == nil || IsNil(o.Attachment7.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment7.Get()
}

// GetAttachment7Ok returns a tuple with the Attachment7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment7Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment7.Get(), o.Attachment7.IsSet()
}

// HasAttachment7 returns a boolean if a field has been set.
func (o *Post) HasAttachment7() bool {
	if o != nil && o.Attachment7.IsSet() {
		return true
	}

	return false
}

// SetAttachment7 gets a reference to the given NullableString and assigns it to the Attachment7 field.
func (o *Post) SetAttachment7(v string) {
	o.Attachment7.Set(&v)
}
// SetAttachment7Nil sets the value for Attachment7 to be an explicit nil
func (o *Post) SetAttachment7Nil() {
	o.Attachment7.Set(nil)
}

// UnsetAttachment7 ensures that no value is present for Attachment7, not even an explicit nil
func (o *Post) UnsetAttachment7() {
	o.Attachment7.Unset()
}

// GetAttachment7Thumbnail returns the Attachment7Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment7Thumbnail() string {
	if o == nil || IsNil(o.Attachment7Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment7Thumbnail.Get()
}

// GetAttachment7ThumbnailOk returns a tuple with the Attachment7Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment7ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment7Thumbnail.Get(), o.Attachment7Thumbnail.IsSet()
}

// HasAttachment7Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment7Thumbnail() bool {
	if o != nil && o.Attachment7Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment7Thumbnail gets a reference to the given NullableString and assigns it to the Attachment7Thumbnail field.
func (o *Post) SetAttachment7Thumbnail(v string) {
	o.Attachment7Thumbnail.Set(&v)
}
// SetAttachment7ThumbnailNil sets the value for Attachment7Thumbnail to be an explicit nil
func (o *Post) SetAttachment7ThumbnailNil() {
	o.Attachment7Thumbnail.Set(nil)
}

// UnsetAttachment7Thumbnail ensures that no value is present for Attachment7Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment7Thumbnail() {
	o.Attachment7Thumbnail.Unset()
}

// GetAttachment8 returns the Attachment8 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment8() string {
	if o == nil || IsNil(o.Attachment8.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment8.Get()
}

// GetAttachment8Ok returns a tuple with the Attachment8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment8Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment8.Get(), o.Attachment8.IsSet()
}

// HasAttachment8 returns a boolean if a field has been set.
func (o *Post) HasAttachment8() bool {
	if o != nil && o.Attachment8.IsSet() {
		return true
	}

	return false
}

// SetAttachment8 gets a reference to the given NullableString and assigns it to the Attachment8 field.
func (o *Post) SetAttachment8(v string) {
	o.Attachment8.Set(&v)
}
// SetAttachment8Nil sets the value for Attachment8 to be an explicit nil
func (o *Post) SetAttachment8Nil() {
	o.Attachment8.Set(nil)
}

// UnsetAttachment8 ensures that no value is present for Attachment8, not even an explicit nil
func (o *Post) UnsetAttachment8() {
	o.Attachment8.Unset()
}

// GetAttachment8Thumbnail returns the Attachment8Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment8Thumbnail() string {
	if o == nil || IsNil(o.Attachment8Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment8Thumbnail.Get()
}

// GetAttachment8ThumbnailOk returns a tuple with the Attachment8Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment8ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment8Thumbnail.Get(), o.Attachment8Thumbnail.IsSet()
}

// HasAttachment8Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment8Thumbnail() bool {
	if o != nil && o.Attachment8Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment8Thumbnail gets a reference to the given NullableString and assigns it to the Attachment8Thumbnail field.
func (o *Post) SetAttachment8Thumbnail(v string) {
	o.Attachment8Thumbnail.Set(&v)
}
// SetAttachment8ThumbnailNil sets the value for Attachment8Thumbnail to be an explicit nil
func (o *Post) SetAttachment8ThumbnailNil() {
	o.Attachment8Thumbnail.Set(nil)
}

// UnsetAttachment8Thumbnail ensures that no value is present for Attachment8Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment8Thumbnail() {
	o.Attachment8Thumbnail.Unset()
}

// GetAttachment9 returns the Attachment9 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment9() string {
	if o == nil || IsNil(o.Attachment9.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment9.Get()
}

// GetAttachment9Ok returns a tuple with the Attachment9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment9Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment9.Get(), o.Attachment9.IsSet()
}

// HasAttachment9 returns a boolean if a field has been set.
func (o *Post) HasAttachment9() bool {
	if o != nil && o.Attachment9.IsSet() {
		return true
	}

	return false
}

// SetAttachment9 gets a reference to the given NullableString and assigns it to the Attachment9 field.
func (o *Post) SetAttachment9(v string) {
	o.Attachment9.Set(&v)
}
// SetAttachment9Nil sets the value for Attachment9 to be an explicit nil
func (o *Post) SetAttachment9Nil() {
	o.Attachment9.Set(nil)
}

// UnsetAttachment9 ensures that no value is present for Attachment9, not even an explicit nil
func (o *Post) UnsetAttachment9() {
	o.Attachment9.Unset()
}

// GetAttachment9Thumbnail returns the Attachment9Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachment9Thumbnail() string {
	if o == nil || IsNil(o.Attachment9Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment9Thumbnail.Get()
}

// GetAttachment9ThumbnailOk returns a tuple with the Attachment9Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachment9ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment9Thumbnail.Get(), o.Attachment9Thumbnail.IsSet()
}

// HasAttachment9Thumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachment9Thumbnail() bool {
	if o != nil && o.Attachment9Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment9Thumbnail gets a reference to the given NullableString and assigns it to the Attachment9Thumbnail field.
func (o *Post) SetAttachment9Thumbnail(v string) {
	o.Attachment9Thumbnail.Set(&v)
}
// SetAttachment9ThumbnailNil sets the value for Attachment9Thumbnail to be an explicit nil
func (o *Post) SetAttachment9ThumbnailNil() {
	o.Attachment9Thumbnail.Set(nil)
}

// UnsetAttachment9Thumbnail ensures that no value is present for Attachment9Thumbnail, not even an explicit nil
func (o *Post) UnsetAttachment9Thumbnail() {
	o.Attachment9Thumbnail.Unset()
}

// GetAttachmentThumbnail returns the AttachmentThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetAttachmentThumbnail() string {
	if o == nil || IsNil(o.AttachmentThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentThumbnail.Get()
}

// GetAttachmentThumbnailOk returns a tuple with the AttachmentThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetAttachmentThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentThumbnail.Get(), o.AttachmentThumbnail.IsSet()
}

// HasAttachmentThumbnail returns a boolean if a field has been set.
func (o *Post) HasAttachmentThumbnail() bool {
	if o != nil && o.AttachmentThumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachmentThumbnail gets a reference to the given NullableString and assigns it to the AttachmentThumbnail field.
func (o *Post) SetAttachmentThumbnail(v string) {
	o.AttachmentThumbnail.Set(&v)
}
// SetAttachmentThumbnailNil sets the value for AttachmentThumbnail to be an explicit nil
func (o *Post) SetAttachmentThumbnailNil() {
	o.AttachmentThumbnail.Set(nil)
}

// UnsetAttachmentThumbnail ensures that no value is present for AttachmentThumbnail, not even an explicit nil
func (o *Post) UnsetAttachmentThumbnail() {
	o.AttachmentThumbnail.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetColor() int32 {
	if o == nil || IsNil(o.Color.Get()) {
		var ret int32
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *Post) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableInt32 and assigns it to the Color field.
func (o *Post) SetColor(v int32) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *Post) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *Post) UnsetColor() {
	o.Color.Unset()
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *Post) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *Post) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *Post) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *Post) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetConversationId() int64 {
	if o == nil || IsNil(o.ConversationId.Get()) {
		var ret int64
		return ret
	}
	return *o.ConversationId.Get()
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetConversationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConversationId.Get(), o.ConversationId.IsSet()
}

// HasConversationId returns a boolean if a field has been set.
func (o *Post) HasConversationId() bool {
	if o != nil && o.ConversationId.IsSet() {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given NullableInt64 and assigns it to the ConversationId field.
func (o *Post) SetConversationId(v int64) {
	o.ConversationId.Set(&v)
}
// SetConversationIdNil sets the value for ConversationId to be an explicit nil
func (o *Post) SetConversationIdNil() {
	o.ConversationId.Set(nil)
}

// UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
func (o *Post) UnsetConversationId() {
	o.ConversationId.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Post) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *Post) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Post) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Post) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetEditedAt returns the EditedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetEditedAt() int64 {
	if o == nil || IsNil(o.EditedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.EditedAt.Get()
}

// GetEditedAtOk returns a tuple with the EditedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetEditedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditedAt.Get(), o.EditedAt.IsSet()
}

// HasEditedAt returns a boolean if a field has been set.
func (o *Post) HasEditedAt() bool {
	if o != nil && o.EditedAt.IsSet() {
		return true
	}

	return false
}

// SetEditedAt gets a reference to the given NullableInt64 and assigns it to the EditedAt field.
func (o *Post) SetEditedAt(v int64) {
	o.EditedAt.Set(&v)
}
// SetEditedAtNil sets the value for EditedAt to be an explicit nil
func (o *Post) SetEditedAtNil() {
	o.EditedAt.Set(nil)
}

// UnsetEditedAt ensures that no value is present for EditedAt, not even an explicit nil
func (o *Post) UnsetEditedAt() {
	o.EditedAt.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret int32
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetFontSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *Post) HasFontSize() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableInt32 and assigns it to the FontSize field.
func (o *Post) SetFontSize(v int32) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *Post) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *Post) UnsetFontSize() {
	o.FontSize.Unset()
}

// GetGiftsCount returns the GiftsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetGiftsCount() []GiftCount {
	if o == nil {
		var ret []GiftCount
		return ret
	}
	return o.GiftsCount
}

// GetGiftsCountOk returns a tuple with the GiftsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetGiftsCountOk() ([]GiftCount, bool) {
	if o == nil || IsNil(o.GiftsCount) {
		return nil, false
	}
	return o.GiftsCount, true
}

// HasGiftsCount returns a boolean if a field has been set.
func (o *Post) HasGiftsCount() bool {
	if o != nil && !IsNil(o.GiftsCount) {
		return true
	}

	return false
}

// SetGiftsCount gets a reference to the given []GiftCount and assigns it to the GiftsCount field.
func (o *Post) SetGiftsCount(v []GiftCount) {
	o.GiftsCount = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetGroup() Group {
	if o == nil || IsNil(o.Group.Get()) {
		var ret Group
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Post) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableGroup and assigns it to the Group field.
func (o *Post) SetGroup(v Group) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Post) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Post) UnsetGroup() {
	o.Group.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *Post) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *Post) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *Post) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *Post) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetHighlighted returns the Highlighted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetHighlighted() bool {
	if o == nil || IsNil(o.Highlighted.Get()) {
		var ret bool
		return ret
	}
	return *o.Highlighted.Get()
}

// GetHighlightedOk returns a tuple with the Highlighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetHighlightedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Highlighted.Get(), o.Highlighted.IsSet()
}

// HasHighlighted returns a boolean if a field has been set.
func (o *Post) HasHighlighted() bool {
	if o != nil && o.Highlighted.IsSet() {
		return true
	}

	return false
}

// SetHighlighted gets a reference to the given NullableBool and assigns it to the Highlighted field.
func (o *Post) SetHighlighted(v bool) {
	o.Highlighted.Set(&v)
}
// SetHighlightedNil sets the value for Highlighted to be an explicit nil
func (o *Post) SetHighlightedNil() {
	o.Highlighted.Set(nil)
}

// UnsetHighlighted ensures that no value is present for Highlighted, not even an explicit nil
func (o *Post) UnsetHighlighted() {
	o.Highlighted.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Post) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Post) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Post) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Post) UnsetId() {
	o.Id.Unset()
}

// GetInReplyTo returns the InReplyTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetInReplyTo() int64 {
	if o == nil || IsNil(o.InReplyTo.Get()) {
		var ret int64
		return ret
	}
	return *o.InReplyTo.Get()
}

// GetInReplyToOk returns a tuple with the InReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetInReplyToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyTo.Get(), o.InReplyTo.IsSet()
}

// HasInReplyTo returns a boolean if a field has been set.
func (o *Post) HasInReplyTo() bool {
	if o != nil && o.InReplyTo.IsSet() {
		return true
	}

	return false
}

// SetInReplyTo gets a reference to the given NullableInt64 and assigns it to the InReplyTo field.
func (o *Post) SetInReplyTo(v int64) {
	o.InReplyTo.Set(&v)
}
// SetInReplyToNil sets the value for InReplyTo to be an explicit nil
func (o *Post) SetInReplyToNil() {
	o.InReplyTo.Set(nil)
}

// UnsetInReplyTo ensures that no value is present for InReplyTo, not even an explicit nil
func (o *Post) UnsetInReplyTo() {
	o.InReplyTo.Unset()
}

// GetInReplyToPost returns the InReplyToPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetInReplyToPost() Post {
	if o == nil || IsNil(o.InReplyToPost.Get()) {
		var ret Post
		return ret
	}
	return *o.InReplyToPost.Get()
}

// GetInReplyToPostOk returns a tuple with the InReplyToPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetInReplyToPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyToPost.Get(), o.InReplyToPost.IsSet()
}

// HasInReplyToPost returns a boolean if a field has been set.
func (o *Post) HasInReplyToPost() bool {
	if o != nil && o.InReplyToPost.IsSet() {
		return true
	}

	return false
}

// SetInReplyToPost gets a reference to the given NullablePost and assigns it to the InReplyToPost field.
func (o *Post) SetInReplyToPost(v Post) {
	o.InReplyToPost.Set(&v)
}
// SetInReplyToPostNil sets the value for InReplyToPost to be an explicit nil
func (o *Post) SetInReplyToPostNil() {
	o.InReplyToPost.Set(nil)
}

// UnsetInReplyToPost ensures that no value is present for InReplyToPost, not even an explicit nil
func (o *Post) UnsetInReplyToPost() {
	o.InReplyToPost.Unset()
}

// GetInReplyToPostCount returns the InReplyToPostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetInReplyToPostCount() int32 {
	if o == nil || IsNil(o.InReplyToPostCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InReplyToPostCount.Get()
}

// GetInReplyToPostCountOk returns a tuple with the InReplyToPostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetInReplyToPostCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyToPostCount.Get(), o.InReplyToPostCount.IsSet()
}

// HasInReplyToPostCount returns a boolean if a field has been set.
func (o *Post) HasInReplyToPostCount() bool {
	if o != nil && o.InReplyToPostCount.IsSet() {
		return true
	}

	return false
}

// SetInReplyToPostCount gets a reference to the given NullableInt32 and assigns it to the InReplyToPostCount field.
func (o *Post) SetInReplyToPostCount(v int32) {
	o.InReplyToPostCount.Set(&v)
}
// SetInReplyToPostCountNil sets the value for InReplyToPostCount to be an explicit nil
func (o *Post) SetInReplyToPostCountNil() {
	o.InReplyToPostCount.Set(nil)
}

// UnsetInReplyToPostCount ensures that no value is present for InReplyToPostCount, not even an explicit nil
func (o *Post) UnsetInReplyToPostCount() {
	o.InReplyToPostCount.Unset()
}

// GetIsFailToSend returns the IsFailToSend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetIsFailToSend() bool {
	if o == nil || IsNil(o.IsFailToSend.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFailToSend.Get()
}

// GetIsFailToSendOk returns a tuple with the IsFailToSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetIsFailToSendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFailToSend.Get(), o.IsFailToSend.IsSet()
}

// HasIsFailToSend returns a boolean if a field has been set.
func (o *Post) HasIsFailToSend() bool {
	if o != nil && o.IsFailToSend.IsSet() {
		return true
	}

	return false
}

// SetIsFailToSend gets a reference to the given NullableBool and assigns it to the IsFailToSend field.
func (o *Post) SetIsFailToSend(v bool) {
	o.IsFailToSend.Set(&v)
}
// SetIsFailToSendNil sets the value for IsFailToSend to be an explicit nil
func (o *Post) SetIsFailToSendNil() {
	o.IsFailToSend.Set(nil)
}

// UnsetIsFailToSend ensures that no value is present for IsFailToSend, not even an explicit nil
func (o *Post) UnsetIsFailToSend() {
	o.IsFailToSend.Unset()
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMuted.Get()
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMuted.Get(), o.IsMuted.IsSet()
}

// HasIsMuted returns a boolean if a field has been set.
func (o *Post) HasIsMuted() bool {
	if o != nil && o.IsMuted.IsSet() {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given NullableBool and assigns it to the IsMuted field.
func (o *Post) SetIsMuted(v bool) {
	o.IsMuted.Set(&v)
}
// SetIsMutedNil sets the value for IsMuted to be an explicit nil
func (o *Post) SetIsMutedNil() {
	o.IsMuted.Set(nil)
}

// UnsetIsMuted ensures that no value is present for IsMuted, not even an explicit nil
func (o *Post) UnsetIsMuted() {
	o.IsMuted.Unset()
}

// GetLiked returns the Liked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetLiked() bool {
	if o == nil || IsNil(o.Liked.Get()) {
		var ret bool
		return ret
	}
	return *o.Liked.Get()
}

// GetLikedOk returns a tuple with the Liked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetLikedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Liked.Get(), o.Liked.IsSet()
}

// HasLiked returns a boolean if a field has been set.
func (o *Post) HasLiked() bool {
	if o != nil && o.Liked.IsSet() {
		return true
	}

	return false
}

// SetLiked gets a reference to the given NullableBool and assigns it to the Liked field.
func (o *Post) SetLiked(v bool) {
	o.Liked.Set(&v)
}
// SetLikedNil sets the value for Liked to be an explicit nil
func (o *Post) SetLikedNil() {
	o.Liked.Set(nil)
}

// UnsetLiked ensures that no value is present for Liked, not even an explicit nil
func (o *Post) UnsetLiked() {
	o.Liked.Unset()
}

// GetLikers returns the Likers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetLikers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Likers
}

// GetLikersOk returns a tuple with the Likers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetLikersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Likers) {
		return nil, false
	}
	return o.Likers, true
}

// HasLikers returns a boolean if a field has been set.
func (o *Post) HasLikers() bool {
	if o != nil && !IsNil(o.Likers) {
		return true
	}

	return false
}

// SetLikers gets a reference to the given []RealmUser and assigns it to the Likers field.
func (o *Post) SetLikers(v []RealmUser) {
	o.Likers = v
}

// GetLikersCount returns the LikersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetLikersCount() int32 {
	if o == nil || IsNil(o.LikersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LikersCount.Get()
}

// GetLikersCountOk returns a tuple with the LikersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetLikersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LikersCount.Get(), o.LikersCount.IsSet()
}

// HasLikersCount returns a boolean if a field has been set.
func (o *Post) HasLikersCount() bool {
	if o != nil && o.LikersCount.IsSet() {
		return true
	}

	return false
}

// SetLikersCount gets a reference to the given NullableInt32 and assigns it to the LikersCount field.
func (o *Post) SetLikersCount(v int32) {
	o.LikersCount.Set(&v)
}
// SetLikersCountNil sets the value for LikersCount to be an explicit nil
func (o *Post) SetLikersCountNil() {
	o.LikersCount.Set(nil)
}

// UnsetLikersCount ensures that no value is present for LikersCount, not even an explicit nil
func (o *Post) UnsetLikersCount() {
	o.LikersCount.Unset()
}

// GetLikesCount returns the LikesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetLikesCount() int32 {
	if o == nil || IsNil(o.LikesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LikesCount.Get()
}

// GetLikesCountOk returns a tuple with the LikesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetLikesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LikesCount.Get(), o.LikesCount.IsSet()
}

// HasLikesCount returns a boolean if a field has been set.
func (o *Post) HasLikesCount() bool {
	if o != nil && o.LikesCount.IsSet() {
		return true
	}

	return false
}

// SetLikesCount gets a reference to the given NullableInt32 and assigns it to the LikesCount field.
func (o *Post) SetLikesCount(v int32) {
	o.LikesCount.Set(&v)
}
// SetLikesCountNil sets the value for LikesCount to be an explicit nil
func (o *Post) SetLikesCountNil() {
	o.LikesCount.Set(nil)
}

// UnsetLikesCount ensures that no value is present for LikesCount, not even an explicit nil
func (o *Post) UnsetLikesCount() {
	o.LikesCount.Unset()
}

// GetMentions returns the Mentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetMentions() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetMentionsOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *Post) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given []RealmUser and assigns it to the Mentions field.
func (o *Post) SetMentions(v []RealmUser) {
	o.Mentions = v
}

// GetMessageTags returns the MessageTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetMessageTags() []MessageTag {
	if o == nil {
		var ret []MessageTag
		return ret
	}
	return o.MessageTags
}

// GetMessageTagsOk returns a tuple with the MessageTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetMessageTagsOk() ([]MessageTag, bool) {
	if o == nil || IsNil(o.MessageTags) {
		return nil, false
	}
	return o.MessageTags, true
}

// HasMessageTags returns a boolean if a field has been set.
func (o *Post) HasMessageTags() bool {
	if o != nil && !IsNil(o.MessageTags) {
		return true
	}

	return false
}

// SetMessageTags gets a reference to the given []MessageTag and assigns it to the MessageTags field.
func (o *Post) SetMessageTags(v []MessageTag) {
	o.MessageTags = v
}

// GetPostType returns the PostType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetPostType() PostType {
	if o == nil || IsNil(o.PostType.Get()) {
		var ret PostType
		return ret
	}
	return *o.PostType.Get()
}

// GetPostTypeOk returns a tuple with the PostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetPostTypeOk() (*PostType, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostType.Get(), o.PostType.IsSet()
}

// HasPostType returns a boolean if a field has been set.
func (o *Post) HasPostType() bool {
	if o != nil && o.PostType.IsSet() {
		return true
	}

	return false
}

// SetPostType gets a reference to the given NullablePostType and assigns it to the PostType field.
func (o *Post) SetPostType(v PostType) {
	o.PostType.Set(&v)
}
// SetPostTypeNil sets the value for PostType to be an explicit nil
func (o *Post) SetPostTypeNil() {
	o.PostType.Set(nil)
}

// UnsetPostType ensures that no value is present for PostType, not even an explicit nil
func (o *Post) UnsetPostType() {
	o.PostType.Unset()
}

// GetReportedCount returns the ReportedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetReportedCount() int32 {
	if o == nil || IsNil(o.ReportedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReportedCount.Get()
}

// GetReportedCountOk returns a tuple with the ReportedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetReportedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedCount.Get(), o.ReportedCount.IsSet()
}

// HasReportedCount returns a boolean if a field has been set.
func (o *Post) HasReportedCount() bool {
	if o != nil && o.ReportedCount.IsSet() {
		return true
	}

	return false
}

// SetReportedCount gets a reference to the given NullableInt32 and assigns it to the ReportedCount field.
func (o *Post) SetReportedCount(v int32) {
	o.ReportedCount.Set(&v)
}
// SetReportedCountNil sets the value for ReportedCount to be an explicit nil
func (o *Post) SetReportedCountNil() {
	o.ReportedCount.Set(nil)
}

// UnsetReportedCount ensures that no value is present for ReportedCount, not even an explicit nil
func (o *Post) UnsetReportedCount() {
	o.ReportedCount.Unset()
}

// GetRepostable returns the Repostable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetRepostable() bool {
	if o == nil || IsNil(o.Repostable.Get()) {
		var ret bool
		return ret
	}
	return *o.Repostable.Get()
}

// GetRepostableOk returns a tuple with the Repostable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetRepostableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repostable.Get(), o.Repostable.IsSet()
}

// HasRepostable returns a boolean if a field has been set.
func (o *Post) HasRepostable() bool {
	if o != nil && o.Repostable.IsSet() {
		return true
	}

	return false
}

// SetRepostable gets a reference to the given NullableBool and assigns it to the Repostable field.
func (o *Post) SetRepostable(v bool) {
	o.Repostable.Set(&v)
}
// SetRepostableNil sets the value for Repostable to be an explicit nil
func (o *Post) SetRepostableNil() {
	o.Repostable.Set(nil)
}

// UnsetRepostable ensures that no value is present for Repostable, not even an explicit nil
func (o *Post) UnsetRepostable() {
	o.Repostable.Unset()
}

// GetReposted returns the Reposted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetReposted() bool {
	if o == nil || IsNil(o.Reposted.Get()) {
		var ret bool
		return ret
	}
	return *o.Reposted.Get()
}

// GetRepostedOk returns a tuple with the Reposted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetRepostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reposted.Get(), o.Reposted.IsSet()
}

// HasReposted returns a boolean if a field has been set.
func (o *Post) HasReposted() bool {
	if o != nil && o.Reposted.IsSet() {
		return true
	}

	return false
}

// SetReposted gets a reference to the given NullableBool and assigns it to the Reposted field.
func (o *Post) SetReposted(v bool) {
	o.Reposted.Set(&v)
}
// SetRepostedNil sets the value for Reposted to be an explicit nil
func (o *Post) SetRepostedNil() {
	o.Reposted.Set(nil)
}

// UnsetReposted ensures that no value is present for Reposted, not even an explicit nil
func (o *Post) UnsetReposted() {
	o.Reposted.Unset()
}

// GetRepostsCount returns the RepostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetRepostsCount() int32 {
	if o == nil || IsNil(o.RepostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RepostsCount.Get()
}

// GetRepostsCountOk returns a tuple with the RepostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetRepostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepostsCount.Get(), o.RepostsCount.IsSet()
}

// HasRepostsCount returns a boolean if a field has been set.
func (o *Post) HasRepostsCount() bool {
	if o != nil && o.RepostsCount.IsSet() {
		return true
	}

	return false
}

// SetRepostsCount gets a reference to the given NullableInt32 and assigns it to the RepostsCount field.
func (o *Post) SetRepostsCount(v int32) {
	o.RepostsCount.Set(&v)
}
// SetRepostsCountNil sets the value for RepostsCount to be an explicit nil
func (o *Post) SetRepostsCountNil() {
	o.RepostsCount.Set(nil)
}

// UnsetRepostsCount ensures that no value is present for RepostsCount, not even an explicit nil
func (o *Post) UnsetRepostsCount() {
	o.RepostsCount.Unset()
}

// GetShareable returns the Shareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetShareable() Shareable {
	if o == nil || IsNil(o.Shareable.Get()) {
		var ret Shareable
		return ret
	}
	return *o.Shareable.Get()
}

// GetShareableOk returns a tuple with the Shareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetShareableOk() (*Shareable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shareable.Get(), o.Shareable.IsSet()
}

// HasShareable returns a boolean if a field has been set.
func (o *Post) HasShareable() bool {
	if o != nil && o.Shareable.IsSet() {
		return true
	}

	return false
}

// SetShareable gets a reference to the given NullableShareable and assigns it to the Shareable field.
func (o *Post) SetShareable(v Shareable) {
	o.Shareable.Set(&v)
}
// SetShareableNil sets the value for Shareable to be an explicit nil
func (o *Post) SetShareableNil() {
	o.Shareable.Set(nil)
}

// UnsetShareable ensures that no value is present for Shareable, not even an explicit nil
func (o *Post) UnsetShareable() {
	o.Shareable.Unset()
}

// GetSharedThread returns the SharedThread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetSharedThread() ThreadInfo {
	if o == nil || IsNil(o.SharedThread.Get()) {
		var ret ThreadInfo
		return ret
	}
	return *o.SharedThread.Get()
}

// GetSharedThreadOk returns a tuple with the SharedThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetSharedThreadOk() (*ThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedThread.Get(), o.SharedThread.IsSet()
}

// HasSharedThread returns a boolean if a field has been set.
func (o *Post) HasSharedThread() bool {
	if o != nil && o.SharedThread.IsSet() {
		return true
	}

	return false
}

// SetSharedThread gets a reference to the given NullableThreadInfo and assigns it to the SharedThread field.
func (o *Post) SetSharedThread(v ThreadInfo) {
	o.SharedThread.Set(&v)
}
// SetSharedThreadNil sets the value for SharedThread to be an explicit nil
func (o *Post) SetSharedThreadNil() {
	o.SharedThread.Set(nil)
}

// UnsetSharedThread ensures that no value is present for SharedThread, not even an explicit nil
func (o *Post) UnsetSharedThread() {
	o.SharedThread.Unset()
}

// GetSharedUrl returns the SharedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetSharedUrl() SharedUrl {
	if o == nil || IsNil(o.SharedUrl.Get()) {
		var ret SharedUrl
		return ret
	}
	return *o.SharedUrl.Get()
}

// GetSharedUrlOk returns a tuple with the SharedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetSharedUrlOk() (*SharedUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUrl.Get(), o.SharedUrl.IsSet()
}

// HasSharedUrl returns a boolean if a field has been set.
func (o *Post) HasSharedUrl() bool {
	if o != nil && o.SharedUrl.IsSet() {
		return true
	}

	return false
}

// SetSharedUrl gets a reference to the given NullableSharedUrl and assigns it to the SharedUrl field.
func (o *Post) SetSharedUrl(v SharedUrl) {
	o.SharedUrl.Set(&v)
}
// SetSharedUrlNil sets the value for SharedUrl to be an explicit nil
func (o *Post) SetSharedUrlNil() {
	o.SharedUrl.Set(nil)
}

// UnsetSharedUrl ensures that no value is present for SharedUrl, not even an explicit nil
func (o *Post) UnsetSharedUrl() {
	o.SharedUrl.Unset()
}

// GetSurvey returns the Survey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetSurvey() Survey {
	if o == nil || IsNil(o.Survey.Get()) {
		var ret Survey
		return ret
	}
	return *o.Survey.Get()
}

// GetSurveyOk returns a tuple with the Survey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetSurveyOk() (*Survey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Survey.Get(), o.Survey.IsSet()
}

// HasSurvey returns a boolean if a field has been set.
func (o *Post) HasSurvey() bool {
	if o != nil && o.Survey.IsSet() {
		return true
	}

	return false
}

// SetSurvey gets a reference to the given NullableSurvey and assigns it to the Survey field.
func (o *Post) SetSurvey(v Survey) {
	o.Survey.Set(&v)
}
// SetSurveyNil sets the value for Survey to be an explicit nil
func (o *Post) SetSurveyNil() {
	o.Survey.Set(nil)
}

// UnsetSurvey ensures that no value is present for Survey, not even an explicit nil
func (o *Post) UnsetSurvey() {
	o.Survey.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *Post) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *Post) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *Post) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *Post) UnsetTag() {
	o.Tag.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *Post) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *Post) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *Post) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *Post) UnsetText() {
	o.Text.Unset()
}

// GetThread returns the Thread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetThread() ThreadInfo {
	if o == nil || IsNil(o.Thread.Get()) {
		var ret ThreadInfo
		return ret
	}
	return *o.Thread.Get()
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetThreadOk() (*ThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thread.Get(), o.Thread.IsSet()
}

// HasThread returns a boolean if a field has been set.
func (o *Post) HasThread() bool {
	if o != nil && o.Thread.IsSet() {
		return true
	}

	return false
}

// SetThread gets a reference to the given NullableThreadInfo and assigns it to the Thread field.
func (o *Post) SetThread(v ThreadInfo) {
	o.Thread.Set(&v)
}
// SetThreadNil sets the value for Thread to be an explicit nil
func (o *Post) SetThreadNil() {
	o.Thread.Set(nil)
}

// UnsetThread ensures that no value is present for Thread, not even an explicit nil
func (o *Post) UnsetThread() {
	o.Thread.Unset()
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetThreadId() int64 {
	if o == nil || IsNil(o.ThreadId.Get()) {
		var ret int64
		return ret
	}
	return *o.ThreadId.Get()
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetThreadIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadId.Get(), o.ThreadId.IsSet()
}

// HasThreadId returns a boolean if a field has been set.
func (o *Post) HasThreadId() bool {
	if o != nil && o.ThreadId.IsSet() {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given NullableInt64 and assigns it to the ThreadId field.
func (o *Post) SetThreadId(v int64) {
	o.ThreadId.Set(&v)
}
// SetThreadIdNil sets the value for ThreadId to be an explicit nil
func (o *Post) SetThreadIdNil() {
	o.ThreadId.Set(nil)
}

// UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
func (o *Post) UnsetThreadId() {
	o.ThreadId.Unset()
}

// GetTotalGiftsCount returns the TotalGiftsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetTotalGiftsCount() int32 {
	if o == nil || IsNil(o.TotalGiftsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalGiftsCount.Get()
}

// GetTotalGiftsCountOk returns a tuple with the TotalGiftsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetTotalGiftsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalGiftsCount.Get(), o.TotalGiftsCount.IsSet()
}

// HasTotalGiftsCount returns a boolean if a field has been set.
func (o *Post) HasTotalGiftsCount() bool {
	if o != nil && o.TotalGiftsCount.IsSet() {
		return true
	}

	return false
}

// SetTotalGiftsCount gets a reference to the given NullableInt32 and assigns it to the TotalGiftsCount field.
func (o *Post) SetTotalGiftsCount(v int32) {
	o.TotalGiftsCount.Set(&v)
}
// SetTotalGiftsCountNil sets the value for TotalGiftsCount to be an explicit nil
func (o *Post) SetTotalGiftsCountNil() {
	o.TotalGiftsCount.Set(nil)
}

// UnsetTotalGiftsCount ensures that no value is present for TotalGiftsCount, not even an explicit nil
func (o *Post) UnsetTotalGiftsCount() {
	o.TotalGiftsCount.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Post) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *Post) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Post) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Post) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Post) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *Post) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Post) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Post) UnsetUser() {
	o.User.Unset()
}

// GetVideos returns the Videos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Post) GetVideos() []Video {
	if o == nil {
		var ret []Video
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Post) GetVideosOk() ([]Video, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *Post) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []Video and assigns it to the Videos field.
func (o *Post) SetVideos(v []Video) {
	o.Videos = v
}

func (o Post) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Post) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachment.IsSet() {
		toSerialize["attachment"] = o.Attachment.Get()
	}
	if o.Attachment2.IsSet() {
		toSerialize["attachment_2"] = o.Attachment2.Get()
	}
	if o.Attachment2Thumbnail.IsSet() {
		toSerialize["attachment_2_thumbnail"] = o.Attachment2Thumbnail.Get()
	}
	if o.Attachment3.IsSet() {
		toSerialize["attachment_3"] = o.Attachment3.Get()
	}
	if o.Attachment3Thumbnail.IsSet() {
		toSerialize["attachment_3_thumbnail"] = o.Attachment3Thumbnail.Get()
	}
	if o.Attachment4.IsSet() {
		toSerialize["attachment_4"] = o.Attachment4.Get()
	}
	if o.Attachment4Thumbnail.IsSet() {
		toSerialize["attachment_4_thumbnail"] = o.Attachment4Thumbnail.Get()
	}
	if o.Attachment5.IsSet() {
		toSerialize["attachment_5"] = o.Attachment5.Get()
	}
	if o.Attachment5Thumbnail.IsSet() {
		toSerialize["attachment_5_thumbnail"] = o.Attachment5Thumbnail.Get()
	}
	if o.Attachment6.IsSet() {
		toSerialize["attachment_6"] = o.Attachment6.Get()
	}
	if o.Attachment6Thumbnail.IsSet() {
		toSerialize["attachment_6_thumbnail"] = o.Attachment6Thumbnail.Get()
	}
	if o.Attachment7.IsSet() {
		toSerialize["attachment_7"] = o.Attachment7.Get()
	}
	if o.Attachment7Thumbnail.IsSet() {
		toSerialize["attachment_7_thumbnail"] = o.Attachment7Thumbnail.Get()
	}
	if o.Attachment8.IsSet() {
		toSerialize["attachment_8"] = o.Attachment8.Get()
	}
	if o.Attachment8Thumbnail.IsSet() {
		toSerialize["attachment_8_thumbnail"] = o.Attachment8Thumbnail.Get()
	}
	if o.Attachment9.IsSet() {
		toSerialize["attachment_9"] = o.Attachment9.Get()
	}
	if o.Attachment9Thumbnail.IsSet() {
		toSerialize["attachment_9_thumbnail"] = o.Attachment9Thumbnail.Get()
	}
	if o.AttachmentThumbnail.IsSet() {
		toSerialize["attachment_thumbnail"] = o.AttachmentThumbnail.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.ConversationId.IsSet() {
		toSerialize["conversation_id"] = o.ConversationId.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.EditedAt.IsSet() {
		toSerialize["edited_at"] = o.EditedAt.Get()
	}
	if o.FontSize.IsSet() {
		toSerialize["font_size"] = o.FontSize.Get()
	}
	if o.GiftsCount != nil {
		toSerialize["gifts_count"] = o.GiftsCount
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if o.Highlighted.IsSet() {
		toSerialize["highlighted"] = o.Highlighted.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InReplyTo.IsSet() {
		toSerialize["in_reply_to"] = o.InReplyTo.Get()
	}
	if o.InReplyToPost.IsSet() {
		toSerialize["in_reply_to_post"] = o.InReplyToPost.Get()
	}
	if o.InReplyToPostCount.IsSet() {
		toSerialize["in_reply_to_post_count"] = o.InReplyToPostCount.Get()
	}
	if o.IsFailToSend.IsSet() {
		toSerialize["is_fail_to_send"] = o.IsFailToSend.Get()
	}
	if o.IsMuted.IsSet() {
		toSerialize["is_muted"] = o.IsMuted.Get()
	}
	if o.Liked.IsSet() {
		toSerialize["liked"] = o.Liked.Get()
	}
	if o.Likers != nil {
		toSerialize["likers"] = o.Likers
	}
	if o.LikersCount.IsSet() {
		toSerialize["likers_count"] = o.LikersCount.Get()
	}
	if o.LikesCount.IsSet() {
		toSerialize["likes_count"] = o.LikesCount.Get()
	}
	if o.Mentions != nil {
		toSerialize["mentions"] = o.Mentions
	}
	if o.MessageTags != nil {
		toSerialize["message_tags"] = o.MessageTags
	}
	if o.PostType.IsSet() {
		toSerialize["post_type"] = o.PostType.Get()
	}
	if o.ReportedCount.IsSet() {
		toSerialize["reported_count"] = o.ReportedCount.Get()
	}
	if o.Repostable.IsSet() {
		toSerialize["repostable"] = o.Repostable.Get()
	}
	if o.Reposted.IsSet() {
		toSerialize["reposted"] = o.Reposted.Get()
	}
	if o.RepostsCount.IsSet() {
		toSerialize["reposts_count"] = o.RepostsCount.Get()
	}
	if o.Shareable.IsSet() {
		toSerialize["shareable"] = o.Shareable.Get()
	}
	if o.SharedThread.IsSet() {
		toSerialize["shared_thread"] = o.SharedThread.Get()
	}
	if o.SharedUrl.IsSet() {
		toSerialize["shared_url"] = o.SharedUrl.Get()
	}
	if o.Survey.IsSet() {
		toSerialize["survey"] = o.Survey.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Thread.IsSet() {
		toSerialize["thread"] = o.Thread.Get()
	}
	if o.ThreadId.IsSet() {
		toSerialize["thread_id"] = o.ThreadId.Get()
	}
	if o.TotalGiftsCount.IsSet() {
		toSerialize["total_gifts_count"] = o.TotalGiftsCount.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Videos != nil {
		toSerialize["videos"] = o.Videos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Post) UnmarshalJSON(data []byte) (err error) {
	varPost := _Post{}

	err = json.Unmarshal(data, &varPost)

	if err != nil {
		return err
	}

	*o = Post(varPost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attachment")
		delete(additionalProperties, "attachment_2")
		delete(additionalProperties, "attachment_2_thumbnail")
		delete(additionalProperties, "attachment_3")
		delete(additionalProperties, "attachment_3_thumbnail")
		delete(additionalProperties, "attachment_4")
		delete(additionalProperties, "attachment_4_thumbnail")
		delete(additionalProperties, "attachment_5")
		delete(additionalProperties, "attachment_5_thumbnail")
		delete(additionalProperties, "attachment_6")
		delete(additionalProperties, "attachment_6_thumbnail")
		delete(additionalProperties, "attachment_7")
		delete(additionalProperties, "attachment_7_thumbnail")
		delete(additionalProperties, "attachment_8")
		delete(additionalProperties, "attachment_8_thumbnail")
		delete(additionalProperties, "attachment_9")
		delete(additionalProperties, "attachment_9_thumbnail")
		delete(additionalProperties, "attachment_thumbnail")
		delete(additionalProperties, "color")
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "conversation_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "edited_at")
		delete(additionalProperties, "font_size")
		delete(additionalProperties, "gifts_count")
		delete(additionalProperties, "group")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "highlighted")
		delete(additionalProperties, "id")
		delete(additionalProperties, "in_reply_to")
		delete(additionalProperties, "in_reply_to_post")
		delete(additionalProperties, "in_reply_to_post_count")
		delete(additionalProperties, "is_fail_to_send")
		delete(additionalProperties, "is_muted")
		delete(additionalProperties, "liked")
		delete(additionalProperties, "likers")
		delete(additionalProperties, "likers_count")
		delete(additionalProperties, "likes_count")
		delete(additionalProperties, "mentions")
		delete(additionalProperties, "message_tags")
		delete(additionalProperties, "post_type")
		delete(additionalProperties, "reported_count")
		delete(additionalProperties, "repostable")
		delete(additionalProperties, "reposted")
		delete(additionalProperties, "reposts_count")
		delete(additionalProperties, "shareable")
		delete(additionalProperties, "shared_thread")
		delete(additionalProperties, "shared_url")
		delete(additionalProperties, "survey")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "text")
		delete(additionalProperties, "thread")
		delete(additionalProperties, "thread_id")
		delete(additionalProperties, "total_gifts_count")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user")
		delete(additionalProperties, "videos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePost struct {
	value *Post
	isSet bool
}

func (v NullablePost) Get() *Post {
	return v.value
}

func (v *NullablePost) Set(val *Post) {
	v.value = val
	v.isSet = true
}

func (v NullablePost) IsSet() bool {
	return v.isSet
}

func (v *NullablePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePost(val *Post) *NullablePost {
	return &NullablePost{value: val, isSet: true}
}

func (v NullablePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


