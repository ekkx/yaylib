
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPost{}

// ModelPost struct for ModelPost
type ModelPost struct {
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
	CategoryId NullableInt64 `json:"category_id,omitempty"`
	Color NullableInt32 `json:"color,omitempty"`
	ConferenceCall NullableConferenceCall `json:"conference_call,omitempty"`
	ConversationId NullableInt64 `json:"conversation_id,omitempty"`
	CreatedAtSeconds NullableInt64 `json:"created_at_seconds,omitempty"`
	EditedAtSeconds NullableInt64 `json:"edited_at_seconds,omitempty"`
	FontSize NullableInt32 `json:"font_size,omitempty"`
	GameTitle NullableString `json:"game_title,omitempty"`
	Gifts []PostGift `json:"gifts,omitempty"`
	Group NullableModelGroup `json:"group,omitempty"`
	GroupId NullableInt64 `json:"group_id,omitempty"`
	GroupUserTitle NullableString `json:"group_user_title,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InReplyTo NullableInt64 `json:"in_reply_to,omitempty"`
	InReplyToPost NullableModelPost `json:"in_reply_to_post,omitempty"`
	InReplyToPostCount NullableInt32 `json:"in_reply_to_post_count,omitempty"`
	IsEdited NullableBool `json:"is_edited,omitempty"`
	IsFailToSend NullableBool `json:"is_fail_to_send,omitempty"`
	IsGroup NullableBool `json:"is_group,omitempty"`
	IsHighlighted NullableBool `json:"is_highlighted,omitempty"`
	IsHot NullableBool `json:"is_hot,omitempty"`
	IsIdFake NullableBool `json:"is_id_fake,omitempty"`
	IsLiked NullableBool `json:"is_liked,omitempty"`
	IsMuted NullableBool `json:"is_muted,omitempty"`
	IsPinned NullableBool `json:"is_pinned,omitempty"`
	IsRecommended NullableBool `json:"is_recommended,omitempty"`
	IsRepostable NullableBool `json:"is_repostable,omitempty"`
	IsReposted NullableBool `json:"is_reposted,omitempty"`
	IsRootOfConversation NullableBool `json:"is_root_of_conversation,omitempty"`
	IsRootOfConversationComment NullableBool `json:"is_root_of_conversation_comment,omitempty"`
	Likers []User `json:"likers,omitempty"`
	LikersCount NullableInt32 `json:"likers_count,omitempty"`
	LikesCount NullableInt32 `json:"likes_count,omitempty"`
	Mentions []User `json:"mentions,omitempty"`
	MessageTags []ModelMessageTag `json:"message_tags,omitempty"`
	// Unresolved Java type: jp.nanameue.yay.data.model.Post.Companion.PostType
	PostType map[string]interface{} `json:"post_type,omitempty"`
	ReportedCount NullableInt32 `json:"reported_count,omitempty"`
	RepostsCount NullableInt32 `json:"reposts_count,omitempty"`
	Shareable NullableModelShareable `json:"shareable,omitempty"`
	SharedThread NullableModelThreadInfo `json:"shared_thread,omitempty"`
	SharedUrl NullableModelSharedUrl `json:"shared_url,omitempty"`
	ShouldShowOriginalSource NullableBool `json:"should_show_original_source,omitempty"`
	ShouldShowSeeMoreConversation NullableBool `json:"should_show_see_more_conversation,omitempty"`
	Survey NullableModelSurvey `json:"survey,omitempty"`
	Tag NullableString `json:"tag,omitempty"`
	Text NullableString `json:"text,omitempty"`
	Thread NullableModelThreadInfo `json:"thread,omitempty"`
	ThreadId NullableInt64 `json:"thread_id,omitempty"`
	TotalGiftsCount NullableInt32 `json:"total_gifts_count,omitempty"`
	UpdatedAtSeconds NullableInt64 `json:"updated_at_seconds,omitempty"`
	User NullableUser `json:"user,omitempty"`
	Videos []ModelVideo `json:"videos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelPost ModelPost

// NewModelPost instantiates a new ModelPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPost() *ModelPost {
	this := ModelPost{}
	return &this
}

// NewModelPostWithDefaults instantiates a new ModelPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPostWithDefaults() *ModelPost {
	this := ModelPost{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment() string {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableString and assigns it to the Attachment field.
func (o *ModelPost) SetAttachment(v string) {
	o.Attachment.Set(&v)
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *ModelPost) SetAttachmentNil() {
	o.Attachment.Set(nil)
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *ModelPost) UnsetAttachment() {
	o.Attachment.Unset()
}

// GetAttachment2 returns the Attachment2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment2() string {
	if o == nil || IsNil(o.Attachment2.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment2.Get()
}

// GetAttachment2Ok returns a tuple with the Attachment2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment2.Get(), o.Attachment2.IsSet()
}

// HasAttachment2 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment2() bool {
	if o != nil && o.Attachment2.IsSet() {
		return true
	}

	return false
}

// SetAttachment2 gets a reference to the given NullableString and assigns it to the Attachment2 field.
func (o *ModelPost) SetAttachment2(v string) {
	o.Attachment2.Set(&v)
}
// SetAttachment2Nil sets the value for Attachment2 to be an explicit nil
func (o *ModelPost) SetAttachment2Nil() {
	o.Attachment2.Set(nil)
}

// UnsetAttachment2 ensures that no value is present for Attachment2, not even an explicit nil
func (o *ModelPost) UnsetAttachment2() {
	o.Attachment2.Unset()
}

// GetAttachment2Thumbnail returns the Attachment2Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment2Thumbnail() string {
	if o == nil || IsNil(o.Attachment2Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment2Thumbnail.Get()
}

// GetAttachment2ThumbnailOk returns a tuple with the Attachment2Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment2ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment2Thumbnail.Get(), o.Attachment2Thumbnail.IsSet()
}

// HasAttachment2Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment2Thumbnail() bool {
	if o != nil && o.Attachment2Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment2Thumbnail gets a reference to the given NullableString and assigns it to the Attachment2Thumbnail field.
func (o *ModelPost) SetAttachment2Thumbnail(v string) {
	o.Attachment2Thumbnail.Set(&v)
}
// SetAttachment2ThumbnailNil sets the value for Attachment2Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment2ThumbnailNil() {
	o.Attachment2Thumbnail.Set(nil)
}

// UnsetAttachment2Thumbnail ensures that no value is present for Attachment2Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment2Thumbnail() {
	o.Attachment2Thumbnail.Unset()
}

// GetAttachment3 returns the Attachment3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment3() string {
	if o == nil || IsNil(o.Attachment3.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment3.Get()
}

// GetAttachment3Ok returns a tuple with the Attachment3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment3.Get(), o.Attachment3.IsSet()
}

// HasAttachment3 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment3() bool {
	if o != nil && o.Attachment3.IsSet() {
		return true
	}

	return false
}

// SetAttachment3 gets a reference to the given NullableString and assigns it to the Attachment3 field.
func (o *ModelPost) SetAttachment3(v string) {
	o.Attachment3.Set(&v)
}
// SetAttachment3Nil sets the value for Attachment3 to be an explicit nil
func (o *ModelPost) SetAttachment3Nil() {
	o.Attachment3.Set(nil)
}

// UnsetAttachment3 ensures that no value is present for Attachment3, not even an explicit nil
func (o *ModelPost) UnsetAttachment3() {
	o.Attachment3.Unset()
}

// GetAttachment3Thumbnail returns the Attachment3Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment3Thumbnail() string {
	if o == nil || IsNil(o.Attachment3Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment3Thumbnail.Get()
}

// GetAttachment3ThumbnailOk returns a tuple with the Attachment3Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment3ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment3Thumbnail.Get(), o.Attachment3Thumbnail.IsSet()
}

// HasAttachment3Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment3Thumbnail() bool {
	if o != nil && o.Attachment3Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment3Thumbnail gets a reference to the given NullableString and assigns it to the Attachment3Thumbnail field.
func (o *ModelPost) SetAttachment3Thumbnail(v string) {
	o.Attachment3Thumbnail.Set(&v)
}
// SetAttachment3ThumbnailNil sets the value for Attachment3Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment3ThumbnailNil() {
	o.Attachment3Thumbnail.Set(nil)
}

// UnsetAttachment3Thumbnail ensures that no value is present for Attachment3Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment3Thumbnail() {
	o.Attachment3Thumbnail.Unset()
}

// GetAttachment4 returns the Attachment4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment4() string {
	if o == nil || IsNil(o.Attachment4.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment4.Get()
}

// GetAttachment4Ok returns a tuple with the Attachment4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment4.Get(), o.Attachment4.IsSet()
}

// HasAttachment4 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment4() bool {
	if o != nil && o.Attachment4.IsSet() {
		return true
	}

	return false
}

// SetAttachment4 gets a reference to the given NullableString and assigns it to the Attachment4 field.
func (o *ModelPost) SetAttachment4(v string) {
	o.Attachment4.Set(&v)
}
// SetAttachment4Nil sets the value for Attachment4 to be an explicit nil
func (o *ModelPost) SetAttachment4Nil() {
	o.Attachment4.Set(nil)
}

// UnsetAttachment4 ensures that no value is present for Attachment4, not even an explicit nil
func (o *ModelPost) UnsetAttachment4() {
	o.Attachment4.Unset()
}

// GetAttachment4Thumbnail returns the Attachment4Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment4Thumbnail() string {
	if o == nil || IsNil(o.Attachment4Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment4Thumbnail.Get()
}

// GetAttachment4ThumbnailOk returns a tuple with the Attachment4Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment4ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment4Thumbnail.Get(), o.Attachment4Thumbnail.IsSet()
}

// HasAttachment4Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment4Thumbnail() bool {
	if o != nil && o.Attachment4Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment4Thumbnail gets a reference to the given NullableString and assigns it to the Attachment4Thumbnail field.
func (o *ModelPost) SetAttachment4Thumbnail(v string) {
	o.Attachment4Thumbnail.Set(&v)
}
// SetAttachment4ThumbnailNil sets the value for Attachment4Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment4ThumbnailNil() {
	o.Attachment4Thumbnail.Set(nil)
}

// UnsetAttachment4Thumbnail ensures that no value is present for Attachment4Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment4Thumbnail() {
	o.Attachment4Thumbnail.Unset()
}

// GetAttachment5 returns the Attachment5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment5() string {
	if o == nil || IsNil(o.Attachment5.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment5.Get()
}

// GetAttachment5Ok returns a tuple with the Attachment5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment5.Get(), o.Attachment5.IsSet()
}

// HasAttachment5 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment5() bool {
	if o != nil && o.Attachment5.IsSet() {
		return true
	}

	return false
}

// SetAttachment5 gets a reference to the given NullableString and assigns it to the Attachment5 field.
func (o *ModelPost) SetAttachment5(v string) {
	o.Attachment5.Set(&v)
}
// SetAttachment5Nil sets the value for Attachment5 to be an explicit nil
func (o *ModelPost) SetAttachment5Nil() {
	o.Attachment5.Set(nil)
}

// UnsetAttachment5 ensures that no value is present for Attachment5, not even an explicit nil
func (o *ModelPost) UnsetAttachment5() {
	o.Attachment5.Unset()
}

// GetAttachment5Thumbnail returns the Attachment5Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment5Thumbnail() string {
	if o == nil || IsNil(o.Attachment5Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment5Thumbnail.Get()
}

// GetAttachment5ThumbnailOk returns a tuple with the Attachment5Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment5ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment5Thumbnail.Get(), o.Attachment5Thumbnail.IsSet()
}

// HasAttachment5Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment5Thumbnail() bool {
	if o != nil && o.Attachment5Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment5Thumbnail gets a reference to the given NullableString and assigns it to the Attachment5Thumbnail field.
func (o *ModelPost) SetAttachment5Thumbnail(v string) {
	o.Attachment5Thumbnail.Set(&v)
}
// SetAttachment5ThumbnailNil sets the value for Attachment5Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment5ThumbnailNil() {
	o.Attachment5Thumbnail.Set(nil)
}

// UnsetAttachment5Thumbnail ensures that no value is present for Attachment5Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment5Thumbnail() {
	o.Attachment5Thumbnail.Unset()
}

// GetAttachment6 returns the Attachment6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment6() string {
	if o == nil || IsNil(o.Attachment6.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment6.Get()
}

// GetAttachment6Ok returns a tuple with the Attachment6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment6.Get(), o.Attachment6.IsSet()
}

// HasAttachment6 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment6() bool {
	if o != nil && o.Attachment6.IsSet() {
		return true
	}

	return false
}

// SetAttachment6 gets a reference to the given NullableString and assigns it to the Attachment6 field.
func (o *ModelPost) SetAttachment6(v string) {
	o.Attachment6.Set(&v)
}
// SetAttachment6Nil sets the value for Attachment6 to be an explicit nil
func (o *ModelPost) SetAttachment6Nil() {
	o.Attachment6.Set(nil)
}

// UnsetAttachment6 ensures that no value is present for Attachment6, not even an explicit nil
func (o *ModelPost) UnsetAttachment6() {
	o.Attachment6.Unset()
}

// GetAttachment6Thumbnail returns the Attachment6Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment6Thumbnail() string {
	if o == nil || IsNil(o.Attachment6Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment6Thumbnail.Get()
}

// GetAttachment6ThumbnailOk returns a tuple with the Attachment6Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment6ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment6Thumbnail.Get(), o.Attachment6Thumbnail.IsSet()
}

// HasAttachment6Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment6Thumbnail() bool {
	if o != nil && o.Attachment6Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment6Thumbnail gets a reference to the given NullableString and assigns it to the Attachment6Thumbnail field.
func (o *ModelPost) SetAttachment6Thumbnail(v string) {
	o.Attachment6Thumbnail.Set(&v)
}
// SetAttachment6ThumbnailNil sets the value for Attachment6Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment6ThumbnailNil() {
	o.Attachment6Thumbnail.Set(nil)
}

// UnsetAttachment6Thumbnail ensures that no value is present for Attachment6Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment6Thumbnail() {
	o.Attachment6Thumbnail.Unset()
}

// GetAttachment7 returns the Attachment7 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment7() string {
	if o == nil || IsNil(o.Attachment7.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment7.Get()
}

// GetAttachment7Ok returns a tuple with the Attachment7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment7Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment7.Get(), o.Attachment7.IsSet()
}

// HasAttachment7 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment7() bool {
	if o != nil && o.Attachment7.IsSet() {
		return true
	}

	return false
}

// SetAttachment7 gets a reference to the given NullableString and assigns it to the Attachment7 field.
func (o *ModelPost) SetAttachment7(v string) {
	o.Attachment7.Set(&v)
}
// SetAttachment7Nil sets the value for Attachment7 to be an explicit nil
func (o *ModelPost) SetAttachment7Nil() {
	o.Attachment7.Set(nil)
}

// UnsetAttachment7 ensures that no value is present for Attachment7, not even an explicit nil
func (o *ModelPost) UnsetAttachment7() {
	o.Attachment7.Unset()
}

// GetAttachment7Thumbnail returns the Attachment7Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment7Thumbnail() string {
	if o == nil || IsNil(o.Attachment7Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment7Thumbnail.Get()
}

// GetAttachment7ThumbnailOk returns a tuple with the Attachment7Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment7ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment7Thumbnail.Get(), o.Attachment7Thumbnail.IsSet()
}

// HasAttachment7Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment7Thumbnail() bool {
	if o != nil && o.Attachment7Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment7Thumbnail gets a reference to the given NullableString and assigns it to the Attachment7Thumbnail field.
func (o *ModelPost) SetAttachment7Thumbnail(v string) {
	o.Attachment7Thumbnail.Set(&v)
}
// SetAttachment7ThumbnailNil sets the value for Attachment7Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment7ThumbnailNil() {
	o.Attachment7Thumbnail.Set(nil)
}

// UnsetAttachment7Thumbnail ensures that no value is present for Attachment7Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment7Thumbnail() {
	o.Attachment7Thumbnail.Unset()
}

// GetAttachment8 returns the Attachment8 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment8() string {
	if o == nil || IsNil(o.Attachment8.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment8.Get()
}

// GetAttachment8Ok returns a tuple with the Attachment8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment8Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment8.Get(), o.Attachment8.IsSet()
}

// HasAttachment8 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment8() bool {
	if o != nil && o.Attachment8.IsSet() {
		return true
	}

	return false
}

// SetAttachment8 gets a reference to the given NullableString and assigns it to the Attachment8 field.
func (o *ModelPost) SetAttachment8(v string) {
	o.Attachment8.Set(&v)
}
// SetAttachment8Nil sets the value for Attachment8 to be an explicit nil
func (o *ModelPost) SetAttachment8Nil() {
	o.Attachment8.Set(nil)
}

// UnsetAttachment8 ensures that no value is present for Attachment8, not even an explicit nil
func (o *ModelPost) UnsetAttachment8() {
	o.Attachment8.Unset()
}

// GetAttachment8Thumbnail returns the Attachment8Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment8Thumbnail() string {
	if o == nil || IsNil(o.Attachment8Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment8Thumbnail.Get()
}

// GetAttachment8ThumbnailOk returns a tuple with the Attachment8Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment8ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment8Thumbnail.Get(), o.Attachment8Thumbnail.IsSet()
}

// HasAttachment8Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment8Thumbnail() bool {
	if o != nil && o.Attachment8Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment8Thumbnail gets a reference to the given NullableString and assigns it to the Attachment8Thumbnail field.
func (o *ModelPost) SetAttachment8Thumbnail(v string) {
	o.Attachment8Thumbnail.Set(&v)
}
// SetAttachment8ThumbnailNil sets the value for Attachment8Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment8ThumbnailNil() {
	o.Attachment8Thumbnail.Set(nil)
}

// UnsetAttachment8Thumbnail ensures that no value is present for Attachment8Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment8Thumbnail() {
	o.Attachment8Thumbnail.Unset()
}

// GetAttachment9 returns the Attachment9 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment9() string {
	if o == nil || IsNil(o.Attachment9.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment9.Get()
}

// GetAttachment9Ok returns a tuple with the Attachment9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment9Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment9.Get(), o.Attachment9.IsSet()
}

// HasAttachment9 returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment9() bool {
	if o != nil && o.Attachment9.IsSet() {
		return true
	}

	return false
}

// SetAttachment9 gets a reference to the given NullableString and assigns it to the Attachment9 field.
func (o *ModelPost) SetAttachment9(v string) {
	o.Attachment9.Set(&v)
}
// SetAttachment9Nil sets the value for Attachment9 to be an explicit nil
func (o *ModelPost) SetAttachment9Nil() {
	o.Attachment9.Set(nil)
}

// UnsetAttachment9 ensures that no value is present for Attachment9, not even an explicit nil
func (o *ModelPost) UnsetAttachment9() {
	o.Attachment9.Unset()
}

// GetAttachment9Thumbnail returns the Attachment9Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachment9Thumbnail() string {
	if o == nil || IsNil(o.Attachment9Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment9Thumbnail.Get()
}

// GetAttachment9ThumbnailOk returns a tuple with the Attachment9Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachment9ThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment9Thumbnail.Get(), o.Attachment9Thumbnail.IsSet()
}

// HasAttachment9Thumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachment9Thumbnail() bool {
	if o != nil && o.Attachment9Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachment9Thumbnail gets a reference to the given NullableString and assigns it to the Attachment9Thumbnail field.
func (o *ModelPost) SetAttachment9Thumbnail(v string) {
	o.Attachment9Thumbnail.Set(&v)
}
// SetAttachment9ThumbnailNil sets the value for Attachment9Thumbnail to be an explicit nil
func (o *ModelPost) SetAttachment9ThumbnailNil() {
	o.Attachment9Thumbnail.Set(nil)
}

// UnsetAttachment9Thumbnail ensures that no value is present for Attachment9Thumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachment9Thumbnail() {
	o.Attachment9Thumbnail.Unset()
}

// GetAttachmentThumbnail returns the AttachmentThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetAttachmentThumbnail() string {
	if o == nil || IsNil(o.AttachmentThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentThumbnail.Get()
}

// GetAttachmentThumbnailOk returns a tuple with the AttachmentThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetAttachmentThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentThumbnail.Get(), o.AttachmentThumbnail.IsSet()
}

// HasAttachmentThumbnail returns a boolean if a field has been set.
func (o *ModelPost) HasAttachmentThumbnail() bool {
	if o != nil && o.AttachmentThumbnail.IsSet() {
		return true
	}

	return false
}

// SetAttachmentThumbnail gets a reference to the given NullableString and assigns it to the AttachmentThumbnail field.
func (o *ModelPost) SetAttachmentThumbnail(v string) {
	o.AttachmentThumbnail.Set(&v)
}
// SetAttachmentThumbnailNil sets the value for AttachmentThumbnail to be an explicit nil
func (o *ModelPost) SetAttachmentThumbnailNil() {
	o.AttachmentThumbnail.Set(nil)
}

// UnsetAttachmentThumbnail ensures that no value is present for AttachmentThumbnail, not even an explicit nil
func (o *ModelPost) UnsetAttachmentThumbnail() {
	o.AttachmentThumbnail.Unset()
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetCategoryId() int64 {
	if o == nil || IsNil(o.CategoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.CategoryId.Get()
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryId.Get(), o.CategoryId.IsSet()
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ModelPost) HasCategoryId() bool {
	if o != nil && o.CategoryId.IsSet() {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given NullableInt64 and assigns it to the CategoryId field.
func (o *ModelPost) SetCategoryId(v int64) {
	o.CategoryId.Set(&v)
}
// SetCategoryIdNil sets the value for CategoryId to be an explicit nil
func (o *ModelPost) SetCategoryIdNil() {
	o.CategoryId.Set(nil)
}

// UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
func (o *ModelPost) UnsetCategoryId() {
	o.CategoryId.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetColor() int32 {
	if o == nil || IsNil(o.Color.Get()) {
		var ret int32
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ModelPost) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableInt32 and assigns it to the Color field.
func (o *ModelPost) SetColor(v int32) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *ModelPost) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ModelPost) UnsetColor() {
	o.Color.Unset()
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetConferenceCall() ConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret ConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetConferenceCallOk() (*ConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *ModelPost) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableConferenceCall and assigns it to the ConferenceCall field.
func (o *ModelPost) SetConferenceCall(v ConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *ModelPost) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *ModelPost) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetConversationId() int64 {
	if o == nil || IsNil(o.ConversationId.Get()) {
		var ret int64
		return ret
	}
	return *o.ConversationId.Get()
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetConversationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConversationId.Get(), o.ConversationId.IsSet()
}

// HasConversationId returns a boolean if a field has been set.
func (o *ModelPost) HasConversationId() bool {
	if o != nil && o.ConversationId.IsSet() {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given NullableInt64 and assigns it to the ConversationId field.
func (o *ModelPost) SetConversationId(v int64) {
	o.ConversationId.Set(&v)
}
// SetConversationIdNil sets the value for ConversationId to be an explicit nil
func (o *ModelPost) SetConversationIdNil() {
	o.ConversationId.Set(nil)
}

// UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
func (o *ModelPost) UnsetConversationId() {
	o.ConversationId.Unset()
}

// GetCreatedAtSeconds returns the CreatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetCreatedAtSeconds() int64 {
	if o == nil || IsNil(o.CreatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtSeconds.Get()
}

// GetCreatedAtSecondsOk returns a tuple with the CreatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetCreatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtSeconds.Get(), o.CreatedAtSeconds.IsSet()
}

// HasCreatedAtSeconds returns a boolean if a field has been set.
func (o *ModelPost) HasCreatedAtSeconds() bool {
	if o != nil && o.CreatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the CreatedAtSeconds field.
func (o *ModelPost) SetCreatedAtSeconds(v int64) {
	o.CreatedAtSeconds.Set(&v)
}
// SetCreatedAtSecondsNil sets the value for CreatedAtSeconds to be an explicit nil
func (o *ModelPost) SetCreatedAtSecondsNil() {
	o.CreatedAtSeconds.Set(nil)
}

// UnsetCreatedAtSeconds ensures that no value is present for CreatedAtSeconds, not even an explicit nil
func (o *ModelPost) UnsetCreatedAtSeconds() {
	o.CreatedAtSeconds.Unset()
}

// GetEditedAtSeconds returns the EditedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetEditedAtSeconds() int64 {
	if o == nil || IsNil(o.EditedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.EditedAtSeconds.Get()
}

// GetEditedAtSecondsOk returns a tuple with the EditedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetEditedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditedAtSeconds.Get(), o.EditedAtSeconds.IsSet()
}

// HasEditedAtSeconds returns a boolean if a field has been set.
func (o *ModelPost) HasEditedAtSeconds() bool {
	if o != nil && o.EditedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetEditedAtSeconds gets a reference to the given NullableInt64 and assigns it to the EditedAtSeconds field.
func (o *ModelPost) SetEditedAtSeconds(v int64) {
	o.EditedAtSeconds.Set(&v)
}
// SetEditedAtSecondsNil sets the value for EditedAtSeconds to be an explicit nil
func (o *ModelPost) SetEditedAtSecondsNil() {
	o.EditedAtSeconds.Set(nil)
}

// UnsetEditedAtSeconds ensures that no value is present for EditedAtSeconds, not even an explicit nil
func (o *ModelPost) UnsetEditedAtSeconds() {
	o.EditedAtSeconds.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret int32
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetFontSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *ModelPost) HasFontSize() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableInt32 and assigns it to the FontSize field.
func (o *ModelPost) SetFontSize(v int32) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *ModelPost) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *ModelPost) UnsetFontSize() {
	o.FontSize.Unset()
}

// GetGameTitle returns the GameTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetGameTitle() string {
	if o == nil || IsNil(o.GameTitle.Get()) {
		var ret string
		return ret
	}
	return *o.GameTitle.Get()
}

// GetGameTitleOk returns a tuple with the GameTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetGameTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GameTitle.Get(), o.GameTitle.IsSet()
}

// HasGameTitle returns a boolean if a field has been set.
func (o *ModelPost) HasGameTitle() bool {
	if o != nil && o.GameTitle.IsSet() {
		return true
	}

	return false
}

// SetGameTitle gets a reference to the given NullableString and assigns it to the GameTitle field.
func (o *ModelPost) SetGameTitle(v string) {
	o.GameTitle.Set(&v)
}
// SetGameTitleNil sets the value for GameTitle to be an explicit nil
func (o *ModelPost) SetGameTitleNil() {
	o.GameTitle.Set(nil)
}

// UnsetGameTitle ensures that no value is present for GameTitle, not even an explicit nil
func (o *ModelPost) UnsetGameTitle() {
	o.GameTitle.Unset()
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetGifts() []PostGift {
	if o == nil {
		var ret []PostGift
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetGiftsOk() ([]PostGift, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *ModelPost) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []PostGift and assigns it to the Gifts field.
func (o *ModelPost) SetGifts(v []PostGift) {
	o.Gifts = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetGroup() ModelGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret ModelGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetGroupOk() (*ModelGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ModelPost) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableModelGroup and assigns it to the Group field.
func (o *ModelPost) SetGroup(v ModelGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ModelPost) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ModelPost) UnsetGroup() {
	o.Group.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *ModelPost) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *ModelPost) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *ModelPost) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *ModelPost) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetGroupUserTitle returns the GroupUserTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetGroupUserTitle() string {
	if o == nil || IsNil(o.GroupUserTitle.Get()) {
		var ret string
		return ret
	}
	return *o.GroupUserTitle.Get()
}

// GetGroupUserTitleOk returns a tuple with the GroupUserTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetGroupUserTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupUserTitle.Get(), o.GroupUserTitle.IsSet()
}

// HasGroupUserTitle returns a boolean if a field has been set.
func (o *ModelPost) HasGroupUserTitle() bool {
	if o != nil && o.GroupUserTitle.IsSet() {
		return true
	}

	return false
}

// SetGroupUserTitle gets a reference to the given NullableString and assigns it to the GroupUserTitle field.
func (o *ModelPost) SetGroupUserTitle(v string) {
	o.GroupUserTitle.Set(&v)
}
// SetGroupUserTitleNil sets the value for GroupUserTitle to be an explicit nil
func (o *ModelPost) SetGroupUserTitleNil() {
	o.GroupUserTitle.Set(nil)
}

// UnsetGroupUserTitle ensures that no value is present for GroupUserTitle, not even an explicit nil
func (o *ModelPost) UnsetGroupUserTitle() {
	o.GroupUserTitle.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelPost) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelPost) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelPost) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelPost) UnsetId() {
	o.Id.Unset()
}

// GetInReplyTo returns the InReplyTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetInReplyTo() int64 {
	if o == nil || IsNil(o.InReplyTo.Get()) {
		var ret int64
		return ret
	}
	return *o.InReplyTo.Get()
}

// GetInReplyToOk returns a tuple with the InReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetInReplyToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyTo.Get(), o.InReplyTo.IsSet()
}

// HasInReplyTo returns a boolean if a field has been set.
func (o *ModelPost) HasInReplyTo() bool {
	if o != nil && o.InReplyTo.IsSet() {
		return true
	}

	return false
}

// SetInReplyTo gets a reference to the given NullableInt64 and assigns it to the InReplyTo field.
func (o *ModelPost) SetInReplyTo(v int64) {
	o.InReplyTo.Set(&v)
}
// SetInReplyToNil sets the value for InReplyTo to be an explicit nil
func (o *ModelPost) SetInReplyToNil() {
	o.InReplyTo.Set(nil)
}

// UnsetInReplyTo ensures that no value is present for InReplyTo, not even an explicit nil
func (o *ModelPost) UnsetInReplyTo() {
	o.InReplyTo.Unset()
}

// GetInReplyToPost returns the InReplyToPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetInReplyToPost() ModelPost {
	if o == nil || IsNil(o.InReplyToPost.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.InReplyToPost.Get()
}

// GetInReplyToPostOk returns a tuple with the InReplyToPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetInReplyToPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyToPost.Get(), o.InReplyToPost.IsSet()
}

// HasInReplyToPost returns a boolean if a field has been set.
func (o *ModelPost) HasInReplyToPost() bool {
	if o != nil && o.InReplyToPost.IsSet() {
		return true
	}

	return false
}

// SetInReplyToPost gets a reference to the given NullableModelPost and assigns it to the InReplyToPost field.
func (o *ModelPost) SetInReplyToPost(v ModelPost) {
	o.InReplyToPost.Set(&v)
}
// SetInReplyToPostNil sets the value for InReplyToPost to be an explicit nil
func (o *ModelPost) SetInReplyToPostNil() {
	o.InReplyToPost.Set(nil)
}

// UnsetInReplyToPost ensures that no value is present for InReplyToPost, not even an explicit nil
func (o *ModelPost) UnsetInReplyToPost() {
	o.InReplyToPost.Unset()
}

// GetInReplyToPostCount returns the InReplyToPostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetInReplyToPostCount() int32 {
	if o == nil || IsNil(o.InReplyToPostCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InReplyToPostCount.Get()
}

// GetInReplyToPostCountOk returns a tuple with the InReplyToPostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetInReplyToPostCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyToPostCount.Get(), o.InReplyToPostCount.IsSet()
}

// HasInReplyToPostCount returns a boolean if a field has been set.
func (o *ModelPost) HasInReplyToPostCount() bool {
	if o != nil && o.InReplyToPostCount.IsSet() {
		return true
	}

	return false
}

// SetInReplyToPostCount gets a reference to the given NullableInt32 and assigns it to the InReplyToPostCount field.
func (o *ModelPost) SetInReplyToPostCount(v int32) {
	o.InReplyToPostCount.Set(&v)
}
// SetInReplyToPostCountNil sets the value for InReplyToPostCount to be an explicit nil
func (o *ModelPost) SetInReplyToPostCountNil() {
	o.InReplyToPostCount.Set(nil)
}

// UnsetInReplyToPostCount ensures that no value is present for InReplyToPostCount, not even an explicit nil
func (o *ModelPost) UnsetInReplyToPostCount() {
	o.InReplyToPostCount.Unset()
}

// GetIsEdited returns the IsEdited field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsEdited() bool {
	if o == nil || IsNil(o.IsEdited.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEdited.Get()
}

// GetIsEditedOk returns a tuple with the IsEdited field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsEditedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEdited.Get(), o.IsEdited.IsSet()
}

// HasIsEdited returns a boolean if a field has been set.
func (o *ModelPost) HasIsEdited() bool {
	if o != nil && o.IsEdited.IsSet() {
		return true
	}

	return false
}

// SetIsEdited gets a reference to the given NullableBool and assigns it to the IsEdited field.
func (o *ModelPost) SetIsEdited(v bool) {
	o.IsEdited.Set(&v)
}
// SetIsEditedNil sets the value for IsEdited to be an explicit nil
func (o *ModelPost) SetIsEditedNil() {
	o.IsEdited.Set(nil)
}

// UnsetIsEdited ensures that no value is present for IsEdited, not even an explicit nil
func (o *ModelPost) UnsetIsEdited() {
	o.IsEdited.Unset()
}

// GetIsFailToSend returns the IsFailToSend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsFailToSend() bool {
	if o == nil || IsNil(o.IsFailToSend.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFailToSend.Get()
}

// GetIsFailToSendOk returns a tuple with the IsFailToSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsFailToSendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFailToSend.Get(), o.IsFailToSend.IsSet()
}

// HasIsFailToSend returns a boolean if a field has been set.
func (o *ModelPost) HasIsFailToSend() bool {
	if o != nil && o.IsFailToSend.IsSet() {
		return true
	}

	return false
}

// SetIsFailToSend gets a reference to the given NullableBool and assigns it to the IsFailToSend field.
func (o *ModelPost) SetIsFailToSend(v bool) {
	o.IsFailToSend.Set(&v)
}
// SetIsFailToSendNil sets the value for IsFailToSend to be an explicit nil
func (o *ModelPost) SetIsFailToSendNil() {
	o.IsFailToSend.Set(nil)
}

// UnsetIsFailToSend ensures that no value is present for IsFailToSend, not even an explicit nil
func (o *ModelPost) UnsetIsFailToSend() {
	o.IsFailToSend.Unset()
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsGroup() bool {
	if o == nil || IsNil(o.IsGroup.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGroup.Get()
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGroup.Get(), o.IsGroup.IsSet()
}

// HasIsGroup returns a boolean if a field has been set.
func (o *ModelPost) HasIsGroup() bool {
	if o != nil && o.IsGroup.IsSet() {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given NullableBool and assigns it to the IsGroup field.
func (o *ModelPost) SetIsGroup(v bool) {
	o.IsGroup.Set(&v)
}
// SetIsGroupNil sets the value for IsGroup to be an explicit nil
func (o *ModelPost) SetIsGroupNil() {
	o.IsGroup.Set(nil)
}

// UnsetIsGroup ensures that no value is present for IsGroup, not even an explicit nil
func (o *ModelPost) UnsetIsGroup() {
	o.IsGroup.Unset()
}

// GetIsHighlighted returns the IsHighlighted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsHighlighted() bool {
	if o == nil || IsNil(o.IsHighlighted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHighlighted.Get()
}

// GetIsHighlightedOk returns a tuple with the IsHighlighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsHighlightedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHighlighted.Get(), o.IsHighlighted.IsSet()
}

// HasIsHighlighted returns a boolean if a field has been set.
func (o *ModelPost) HasIsHighlighted() bool {
	if o != nil && o.IsHighlighted.IsSet() {
		return true
	}

	return false
}

// SetIsHighlighted gets a reference to the given NullableBool and assigns it to the IsHighlighted field.
func (o *ModelPost) SetIsHighlighted(v bool) {
	o.IsHighlighted.Set(&v)
}
// SetIsHighlightedNil sets the value for IsHighlighted to be an explicit nil
func (o *ModelPost) SetIsHighlightedNil() {
	o.IsHighlighted.Set(nil)
}

// UnsetIsHighlighted ensures that no value is present for IsHighlighted, not even an explicit nil
func (o *ModelPost) UnsetIsHighlighted() {
	o.IsHighlighted.Unset()
}

// GetIsHot returns the IsHot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsHot() bool {
	if o == nil || IsNil(o.IsHot.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHot.Get()
}

// GetIsHotOk returns a tuple with the IsHot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsHotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHot.Get(), o.IsHot.IsSet()
}

// HasIsHot returns a boolean if a field has been set.
func (o *ModelPost) HasIsHot() bool {
	if o != nil && o.IsHot.IsSet() {
		return true
	}

	return false
}

// SetIsHot gets a reference to the given NullableBool and assigns it to the IsHot field.
func (o *ModelPost) SetIsHot(v bool) {
	o.IsHot.Set(&v)
}
// SetIsHotNil sets the value for IsHot to be an explicit nil
func (o *ModelPost) SetIsHotNil() {
	o.IsHot.Set(nil)
}

// UnsetIsHot ensures that no value is present for IsHot, not even an explicit nil
func (o *ModelPost) UnsetIsHot() {
	o.IsHot.Unset()
}

// GetIsIdFake returns the IsIdFake field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsIdFake() bool {
	if o == nil || IsNil(o.IsIdFake.Get()) {
		var ret bool
		return ret
	}
	return *o.IsIdFake.Get()
}

// GetIsIdFakeOk returns a tuple with the IsIdFake field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsIdFakeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsIdFake.Get(), o.IsIdFake.IsSet()
}

// HasIsIdFake returns a boolean if a field has been set.
func (o *ModelPost) HasIsIdFake() bool {
	if o != nil && o.IsIdFake.IsSet() {
		return true
	}

	return false
}

// SetIsIdFake gets a reference to the given NullableBool and assigns it to the IsIdFake field.
func (o *ModelPost) SetIsIdFake(v bool) {
	o.IsIdFake.Set(&v)
}
// SetIsIdFakeNil sets the value for IsIdFake to be an explicit nil
func (o *ModelPost) SetIsIdFakeNil() {
	o.IsIdFake.Set(nil)
}

// UnsetIsIdFake ensures that no value is present for IsIdFake, not even an explicit nil
func (o *ModelPost) UnsetIsIdFake() {
	o.IsIdFake.Unset()
}

// GetIsLiked returns the IsLiked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsLiked() bool {
	if o == nil || IsNil(o.IsLiked.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLiked.Get()
}

// GetIsLikedOk returns a tuple with the IsLiked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsLikedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLiked.Get(), o.IsLiked.IsSet()
}

// HasIsLiked returns a boolean if a field has been set.
func (o *ModelPost) HasIsLiked() bool {
	if o != nil && o.IsLiked.IsSet() {
		return true
	}

	return false
}

// SetIsLiked gets a reference to the given NullableBool and assigns it to the IsLiked field.
func (o *ModelPost) SetIsLiked(v bool) {
	o.IsLiked.Set(&v)
}
// SetIsLikedNil sets the value for IsLiked to be an explicit nil
func (o *ModelPost) SetIsLikedNil() {
	o.IsLiked.Set(nil)
}

// UnsetIsLiked ensures that no value is present for IsLiked, not even an explicit nil
func (o *ModelPost) UnsetIsLiked() {
	o.IsLiked.Unset()
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMuted.Get()
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMuted.Get(), o.IsMuted.IsSet()
}

// HasIsMuted returns a boolean if a field has been set.
func (o *ModelPost) HasIsMuted() bool {
	if o != nil && o.IsMuted.IsSet() {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given NullableBool and assigns it to the IsMuted field.
func (o *ModelPost) SetIsMuted(v bool) {
	o.IsMuted.Set(&v)
}
// SetIsMutedNil sets the value for IsMuted to be an explicit nil
func (o *ModelPost) SetIsMutedNil() {
	o.IsMuted.Set(nil)
}

// UnsetIsMuted ensures that no value is present for IsMuted, not even an explicit nil
func (o *ModelPost) UnsetIsMuted() {
	o.IsMuted.Unset()
}

// GetIsPinned returns the IsPinned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsPinned() bool {
	if o == nil || IsNil(o.IsPinned.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPinned.Get()
}

// GetIsPinnedOk returns a tuple with the IsPinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPinned.Get(), o.IsPinned.IsSet()
}

// HasIsPinned returns a boolean if a field has been set.
func (o *ModelPost) HasIsPinned() bool {
	if o != nil && o.IsPinned.IsSet() {
		return true
	}

	return false
}

// SetIsPinned gets a reference to the given NullableBool and assigns it to the IsPinned field.
func (o *ModelPost) SetIsPinned(v bool) {
	o.IsPinned.Set(&v)
}
// SetIsPinnedNil sets the value for IsPinned to be an explicit nil
func (o *ModelPost) SetIsPinnedNil() {
	o.IsPinned.Set(nil)
}

// UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
func (o *ModelPost) UnsetIsPinned() {
	o.IsPinned.Unset()
}

// GetIsRecommended returns the IsRecommended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsRecommended() bool {
	if o == nil || IsNil(o.IsRecommended.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRecommended.Get()
}

// GetIsRecommendedOk returns a tuple with the IsRecommended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsRecommendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRecommended.Get(), o.IsRecommended.IsSet()
}

// HasIsRecommended returns a boolean if a field has been set.
func (o *ModelPost) HasIsRecommended() bool {
	if o != nil && o.IsRecommended.IsSet() {
		return true
	}

	return false
}

// SetIsRecommended gets a reference to the given NullableBool and assigns it to the IsRecommended field.
func (o *ModelPost) SetIsRecommended(v bool) {
	o.IsRecommended.Set(&v)
}
// SetIsRecommendedNil sets the value for IsRecommended to be an explicit nil
func (o *ModelPost) SetIsRecommendedNil() {
	o.IsRecommended.Set(nil)
}

// UnsetIsRecommended ensures that no value is present for IsRecommended, not even an explicit nil
func (o *ModelPost) UnsetIsRecommended() {
	o.IsRecommended.Unset()
}

// GetIsRepostable returns the IsRepostable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsRepostable() bool {
	if o == nil || IsNil(o.IsRepostable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRepostable.Get()
}

// GetIsRepostableOk returns a tuple with the IsRepostable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsRepostableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRepostable.Get(), o.IsRepostable.IsSet()
}

// HasIsRepostable returns a boolean if a field has been set.
func (o *ModelPost) HasIsRepostable() bool {
	if o != nil && o.IsRepostable.IsSet() {
		return true
	}

	return false
}

// SetIsRepostable gets a reference to the given NullableBool and assigns it to the IsRepostable field.
func (o *ModelPost) SetIsRepostable(v bool) {
	o.IsRepostable.Set(&v)
}
// SetIsRepostableNil sets the value for IsRepostable to be an explicit nil
func (o *ModelPost) SetIsRepostableNil() {
	o.IsRepostable.Set(nil)
}

// UnsetIsRepostable ensures that no value is present for IsRepostable, not even an explicit nil
func (o *ModelPost) UnsetIsRepostable() {
	o.IsRepostable.Unset()
}

// GetIsReposted returns the IsReposted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsReposted() bool {
	if o == nil || IsNil(o.IsReposted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsReposted.Get()
}

// GetIsRepostedOk returns a tuple with the IsReposted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsRepostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsReposted.Get(), o.IsReposted.IsSet()
}

// HasIsReposted returns a boolean if a field has been set.
func (o *ModelPost) HasIsReposted() bool {
	if o != nil && o.IsReposted.IsSet() {
		return true
	}

	return false
}

// SetIsReposted gets a reference to the given NullableBool and assigns it to the IsReposted field.
func (o *ModelPost) SetIsReposted(v bool) {
	o.IsReposted.Set(&v)
}
// SetIsRepostedNil sets the value for IsReposted to be an explicit nil
func (o *ModelPost) SetIsRepostedNil() {
	o.IsReposted.Set(nil)
}

// UnsetIsReposted ensures that no value is present for IsReposted, not even an explicit nil
func (o *ModelPost) UnsetIsReposted() {
	o.IsReposted.Unset()
}

// GetIsRootOfConversation returns the IsRootOfConversation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsRootOfConversation() bool {
	if o == nil || IsNil(o.IsRootOfConversation.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRootOfConversation.Get()
}

// GetIsRootOfConversationOk returns a tuple with the IsRootOfConversation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsRootOfConversationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRootOfConversation.Get(), o.IsRootOfConversation.IsSet()
}

// HasIsRootOfConversation returns a boolean if a field has been set.
func (o *ModelPost) HasIsRootOfConversation() bool {
	if o != nil && o.IsRootOfConversation.IsSet() {
		return true
	}

	return false
}

// SetIsRootOfConversation gets a reference to the given NullableBool and assigns it to the IsRootOfConversation field.
func (o *ModelPost) SetIsRootOfConversation(v bool) {
	o.IsRootOfConversation.Set(&v)
}
// SetIsRootOfConversationNil sets the value for IsRootOfConversation to be an explicit nil
func (o *ModelPost) SetIsRootOfConversationNil() {
	o.IsRootOfConversation.Set(nil)
}

// UnsetIsRootOfConversation ensures that no value is present for IsRootOfConversation, not even an explicit nil
func (o *ModelPost) UnsetIsRootOfConversation() {
	o.IsRootOfConversation.Unset()
}

// GetIsRootOfConversationComment returns the IsRootOfConversationComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetIsRootOfConversationComment() bool {
	if o == nil || IsNil(o.IsRootOfConversationComment.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRootOfConversationComment.Get()
}

// GetIsRootOfConversationCommentOk returns a tuple with the IsRootOfConversationComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetIsRootOfConversationCommentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRootOfConversationComment.Get(), o.IsRootOfConversationComment.IsSet()
}

// HasIsRootOfConversationComment returns a boolean if a field has been set.
func (o *ModelPost) HasIsRootOfConversationComment() bool {
	if o != nil && o.IsRootOfConversationComment.IsSet() {
		return true
	}

	return false
}

// SetIsRootOfConversationComment gets a reference to the given NullableBool and assigns it to the IsRootOfConversationComment field.
func (o *ModelPost) SetIsRootOfConversationComment(v bool) {
	o.IsRootOfConversationComment.Set(&v)
}
// SetIsRootOfConversationCommentNil sets the value for IsRootOfConversationComment to be an explicit nil
func (o *ModelPost) SetIsRootOfConversationCommentNil() {
	o.IsRootOfConversationComment.Set(nil)
}

// UnsetIsRootOfConversationComment ensures that no value is present for IsRootOfConversationComment, not even an explicit nil
func (o *ModelPost) UnsetIsRootOfConversationComment() {
	o.IsRootOfConversationComment.Unset()
}

// GetLikers returns the Likers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetLikers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Likers
}

// GetLikersOk returns a tuple with the Likers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetLikersOk() ([]User, bool) {
	if o == nil || IsNil(o.Likers) {
		return nil, false
	}
	return o.Likers, true
}

// HasLikers returns a boolean if a field has been set.
func (o *ModelPost) HasLikers() bool {
	if o != nil && !IsNil(o.Likers) {
		return true
	}

	return false
}

// SetLikers gets a reference to the given []User and assigns it to the Likers field.
func (o *ModelPost) SetLikers(v []User) {
	o.Likers = v
}

// GetLikersCount returns the LikersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetLikersCount() int32 {
	if o == nil || IsNil(o.LikersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LikersCount.Get()
}

// GetLikersCountOk returns a tuple with the LikersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetLikersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LikersCount.Get(), o.LikersCount.IsSet()
}

// HasLikersCount returns a boolean if a field has been set.
func (o *ModelPost) HasLikersCount() bool {
	if o != nil && o.LikersCount.IsSet() {
		return true
	}

	return false
}

// SetLikersCount gets a reference to the given NullableInt32 and assigns it to the LikersCount field.
func (o *ModelPost) SetLikersCount(v int32) {
	o.LikersCount.Set(&v)
}
// SetLikersCountNil sets the value for LikersCount to be an explicit nil
func (o *ModelPost) SetLikersCountNil() {
	o.LikersCount.Set(nil)
}

// UnsetLikersCount ensures that no value is present for LikersCount, not even an explicit nil
func (o *ModelPost) UnsetLikersCount() {
	o.LikersCount.Unset()
}

// GetLikesCount returns the LikesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetLikesCount() int32 {
	if o == nil || IsNil(o.LikesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LikesCount.Get()
}

// GetLikesCountOk returns a tuple with the LikesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetLikesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LikesCount.Get(), o.LikesCount.IsSet()
}

// HasLikesCount returns a boolean if a field has been set.
func (o *ModelPost) HasLikesCount() bool {
	if o != nil && o.LikesCount.IsSet() {
		return true
	}

	return false
}

// SetLikesCount gets a reference to the given NullableInt32 and assigns it to the LikesCount field.
func (o *ModelPost) SetLikesCount(v int32) {
	o.LikesCount.Set(&v)
}
// SetLikesCountNil sets the value for LikesCount to be an explicit nil
func (o *ModelPost) SetLikesCountNil() {
	o.LikesCount.Set(nil)
}

// UnsetLikesCount ensures that no value is present for LikesCount, not even an explicit nil
func (o *ModelPost) UnsetLikesCount() {
	o.LikesCount.Unset()
}

// GetMentions returns the Mentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetMentions() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetMentionsOk() ([]User, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *ModelPost) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given []User and assigns it to the Mentions field.
func (o *ModelPost) SetMentions(v []User) {
	o.Mentions = v
}

// GetMessageTags returns the MessageTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetMessageTags() []ModelMessageTag {
	if o == nil {
		var ret []ModelMessageTag
		return ret
	}
	return o.MessageTags
}

// GetMessageTagsOk returns a tuple with the MessageTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetMessageTagsOk() ([]ModelMessageTag, bool) {
	if o == nil || IsNil(o.MessageTags) {
		return nil, false
	}
	return o.MessageTags, true
}

// HasMessageTags returns a boolean if a field has been set.
func (o *ModelPost) HasMessageTags() bool {
	if o != nil && !IsNil(o.MessageTags) {
		return true
	}

	return false
}

// SetMessageTags gets a reference to the given []ModelMessageTag and assigns it to the MessageTags field.
func (o *ModelPost) SetMessageTags(v []ModelMessageTag) {
	o.MessageTags = v
}

// GetPostType returns the PostType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetPostType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PostType
}

// GetPostTypeOk returns a tuple with the PostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetPostTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PostType) {
		return map[string]interface{}{}, false
	}
	return o.PostType, true
}

// HasPostType returns a boolean if a field has been set.
func (o *ModelPost) HasPostType() bool {
	if o != nil && !IsNil(o.PostType) {
		return true
	}

	return false
}

// SetPostType gets a reference to the given map[string]interface{} and assigns it to the PostType field.
func (o *ModelPost) SetPostType(v map[string]interface{}) {
	o.PostType = v
}

// GetReportedCount returns the ReportedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetReportedCount() int32 {
	if o == nil || IsNil(o.ReportedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReportedCount.Get()
}

// GetReportedCountOk returns a tuple with the ReportedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetReportedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedCount.Get(), o.ReportedCount.IsSet()
}

// HasReportedCount returns a boolean if a field has been set.
func (o *ModelPost) HasReportedCount() bool {
	if o != nil && o.ReportedCount.IsSet() {
		return true
	}

	return false
}

// SetReportedCount gets a reference to the given NullableInt32 and assigns it to the ReportedCount field.
func (o *ModelPost) SetReportedCount(v int32) {
	o.ReportedCount.Set(&v)
}
// SetReportedCountNil sets the value for ReportedCount to be an explicit nil
func (o *ModelPost) SetReportedCountNil() {
	o.ReportedCount.Set(nil)
}

// UnsetReportedCount ensures that no value is present for ReportedCount, not even an explicit nil
func (o *ModelPost) UnsetReportedCount() {
	o.ReportedCount.Unset()
}

// GetRepostsCount returns the RepostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetRepostsCount() int32 {
	if o == nil || IsNil(o.RepostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RepostsCount.Get()
}

// GetRepostsCountOk returns a tuple with the RepostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetRepostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepostsCount.Get(), o.RepostsCount.IsSet()
}

// HasRepostsCount returns a boolean if a field has been set.
func (o *ModelPost) HasRepostsCount() bool {
	if o != nil && o.RepostsCount.IsSet() {
		return true
	}

	return false
}

// SetRepostsCount gets a reference to the given NullableInt32 and assigns it to the RepostsCount field.
func (o *ModelPost) SetRepostsCount(v int32) {
	o.RepostsCount.Set(&v)
}
// SetRepostsCountNil sets the value for RepostsCount to be an explicit nil
func (o *ModelPost) SetRepostsCountNil() {
	o.RepostsCount.Set(nil)
}

// UnsetRepostsCount ensures that no value is present for RepostsCount, not even an explicit nil
func (o *ModelPost) UnsetRepostsCount() {
	o.RepostsCount.Unset()
}

// GetShareable returns the Shareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetShareable() ModelShareable {
	if o == nil || IsNil(o.Shareable.Get()) {
		var ret ModelShareable
		return ret
	}
	return *o.Shareable.Get()
}

// GetShareableOk returns a tuple with the Shareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetShareableOk() (*ModelShareable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shareable.Get(), o.Shareable.IsSet()
}

// HasShareable returns a boolean if a field has been set.
func (o *ModelPost) HasShareable() bool {
	if o != nil && o.Shareable.IsSet() {
		return true
	}

	return false
}

// SetShareable gets a reference to the given NullableModelShareable and assigns it to the Shareable field.
func (o *ModelPost) SetShareable(v ModelShareable) {
	o.Shareable.Set(&v)
}
// SetShareableNil sets the value for Shareable to be an explicit nil
func (o *ModelPost) SetShareableNil() {
	o.Shareable.Set(nil)
}

// UnsetShareable ensures that no value is present for Shareable, not even an explicit nil
func (o *ModelPost) UnsetShareable() {
	o.Shareable.Unset()
}

// GetSharedThread returns the SharedThread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetSharedThread() ModelThreadInfo {
	if o == nil || IsNil(o.SharedThread.Get()) {
		var ret ModelThreadInfo
		return ret
	}
	return *o.SharedThread.Get()
}

// GetSharedThreadOk returns a tuple with the SharedThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetSharedThreadOk() (*ModelThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedThread.Get(), o.SharedThread.IsSet()
}

// HasSharedThread returns a boolean if a field has been set.
func (o *ModelPost) HasSharedThread() bool {
	if o != nil && o.SharedThread.IsSet() {
		return true
	}

	return false
}

// SetSharedThread gets a reference to the given NullableModelThreadInfo and assigns it to the SharedThread field.
func (o *ModelPost) SetSharedThread(v ModelThreadInfo) {
	o.SharedThread.Set(&v)
}
// SetSharedThreadNil sets the value for SharedThread to be an explicit nil
func (o *ModelPost) SetSharedThreadNil() {
	o.SharedThread.Set(nil)
}

// UnsetSharedThread ensures that no value is present for SharedThread, not even an explicit nil
func (o *ModelPost) UnsetSharedThread() {
	o.SharedThread.Unset()
}

// GetSharedUrl returns the SharedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetSharedUrl() ModelSharedUrl {
	if o == nil || IsNil(o.SharedUrl.Get()) {
		var ret ModelSharedUrl
		return ret
	}
	return *o.SharedUrl.Get()
}

// GetSharedUrlOk returns a tuple with the SharedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetSharedUrlOk() (*ModelSharedUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUrl.Get(), o.SharedUrl.IsSet()
}

// HasSharedUrl returns a boolean if a field has been set.
func (o *ModelPost) HasSharedUrl() bool {
	if o != nil && o.SharedUrl.IsSet() {
		return true
	}

	return false
}

// SetSharedUrl gets a reference to the given NullableModelSharedUrl and assigns it to the SharedUrl field.
func (o *ModelPost) SetSharedUrl(v ModelSharedUrl) {
	o.SharedUrl.Set(&v)
}
// SetSharedUrlNil sets the value for SharedUrl to be an explicit nil
func (o *ModelPost) SetSharedUrlNil() {
	o.SharedUrl.Set(nil)
}

// UnsetSharedUrl ensures that no value is present for SharedUrl, not even an explicit nil
func (o *ModelPost) UnsetSharedUrl() {
	o.SharedUrl.Unset()
}

// GetShouldShowOriginalSource returns the ShouldShowOriginalSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetShouldShowOriginalSource() bool {
	if o == nil || IsNil(o.ShouldShowOriginalSource.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowOriginalSource.Get()
}

// GetShouldShowOriginalSourceOk returns a tuple with the ShouldShowOriginalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetShouldShowOriginalSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowOriginalSource.Get(), o.ShouldShowOriginalSource.IsSet()
}

// HasShouldShowOriginalSource returns a boolean if a field has been set.
func (o *ModelPost) HasShouldShowOriginalSource() bool {
	if o != nil && o.ShouldShowOriginalSource.IsSet() {
		return true
	}

	return false
}

// SetShouldShowOriginalSource gets a reference to the given NullableBool and assigns it to the ShouldShowOriginalSource field.
func (o *ModelPost) SetShouldShowOriginalSource(v bool) {
	o.ShouldShowOriginalSource.Set(&v)
}
// SetShouldShowOriginalSourceNil sets the value for ShouldShowOriginalSource to be an explicit nil
func (o *ModelPost) SetShouldShowOriginalSourceNil() {
	o.ShouldShowOriginalSource.Set(nil)
}

// UnsetShouldShowOriginalSource ensures that no value is present for ShouldShowOriginalSource, not even an explicit nil
func (o *ModelPost) UnsetShouldShowOriginalSource() {
	o.ShouldShowOriginalSource.Unset()
}

// GetShouldShowSeeMoreConversation returns the ShouldShowSeeMoreConversation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetShouldShowSeeMoreConversation() bool {
	if o == nil || IsNil(o.ShouldShowSeeMoreConversation.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowSeeMoreConversation.Get()
}

// GetShouldShowSeeMoreConversationOk returns a tuple with the ShouldShowSeeMoreConversation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetShouldShowSeeMoreConversationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowSeeMoreConversation.Get(), o.ShouldShowSeeMoreConversation.IsSet()
}

// HasShouldShowSeeMoreConversation returns a boolean if a field has been set.
func (o *ModelPost) HasShouldShowSeeMoreConversation() bool {
	if o != nil && o.ShouldShowSeeMoreConversation.IsSet() {
		return true
	}

	return false
}

// SetShouldShowSeeMoreConversation gets a reference to the given NullableBool and assigns it to the ShouldShowSeeMoreConversation field.
func (o *ModelPost) SetShouldShowSeeMoreConversation(v bool) {
	o.ShouldShowSeeMoreConversation.Set(&v)
}
// SetShouldShowSeeMoreConversationNil sets the value for ShouldShowSeeMoreConversation to be an explicit nil
func (o *ModelPost) SetShouldShowSeeMoreConversationNil() {
	o.ShouldShowSeeMoreConversation.Set(nil)
}

// UnsetShouldShowSeeMoreConversation ensures that no value is present for ShouldShowSeeMoreConversation, not even an explicit nil
func (o *ModelPost) UnsetShouldShowSeeMoreConversation() {
	o.ShouldShowSeeMoreConversation.Unset()
}

// GetSurvey returns the Survey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetSurvey() ModelSurvey {
	if o == nil || IsNil(o.Survey.Get()) {
		var ret ModelSurvey
		return ret
	}
	return *o.Survey.Get()
}

// GetSurveyOk returns a tuple with the Survey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetSurveyOk() (*ModelSurvey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Survey.Get(), o.Survey.IsSet()
}

// HasSurvey returns a boolean if a field has been set.
func (o *ModelPost) HasSurvey() bool {
	if o != nil && o.Survey.IsSet() {
		return true
	}

	return false
}

// SetSurvey gets a reference to the given NullableModelSurvey and assigns it to the Survey field.
func (o *ModelPost) SetSurvey(v ModelSurvey) {
	o.Survey.Set(&v)
}
// SetSurveyNil sets the value for Survey to be an explicit nil
func (o *ModelPost) SetSurveyNil() {
	o.Survey.Set(nil)
}

// UnsetSurvey ensures that no value is present for Survey, not even an explicit nil
func (o *ModelPost) UnsetSurvey() {
	o.Survey.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *ModelPost) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *ModelPost) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *ModelPost) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *ModelPost) UnsetTag() {
	o.Tag.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ModelPost) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ModelPost) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ModelPost) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ModelPost) UnsetText() {
	o.Text.Unset()
}

// GetThread returns the Thread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetThread() ModelThreadInfo {
	if o == nil || IsNil(o.Thread.Get()) {
		var ret ModelThreadInfo
		return ret
	}
	return *o.Thread.Get()
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetThreadOk() (*ModelThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thread.Get(), o.Thread.IsSet()
}

// HasThread returns a boolean if a field has been set.
func (o *ModelPost) HasThread() bool {
	if o != nil && o.Thread.IsSet() {
		return true
	}

	return false
}

// SetThread gets a reference to the given NullableModelThreadInfo and assigns it to the Thread field.
func (o *ModelPost) SetThread(v ModelThreadInfo) {
	o.Thread.Set(&v)
}
// SetThreadNil sets the value for Thread to be an explicit nil
func (o *ModelPost) SetThreadNil() {
	o.Thread.Set(nil)
}

// UnsetThread ensures that no value is present for Thread, not even an explicit nil
func (o *ModelPost) UnsetThread() {
	o.Thread.Unset()
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetThreadId() int64 {
	if o == nil || IsNil(o.ThreadId.Get()) {
		var ret int64
		return ret
	}
	return *o.ThreadId.Get()
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetThreadIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadId.Get(), o.ThreadId.IsSet()
}

// HasThreadId returns a boolean if a field has been set.
func (o *ModelPost) HasThreadId() bool {
	if o != nil && o.ThreadId.IsSet() {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given NullableInt64 and assigns it to the ThreadId field.
func (o *ModelPost) SetThreadId(v int64) {
	o.ThreadId.Set(&v)
}
// SetThreadIdNil sets the value for ThreadId to be an explicit nil
func (o *ModelPost) SetThreadIdNil() {
	o.ThreadId.Set(nil)
}

// UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
func (o *ModelPost) UnsetThreadId() {
	o.ThreadId.Unset()
}

// GetTotalGiftsCount returns the TotalGiftsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetTotalGiftsCount() int32 {
	if o == nil || IsNil(o.TotalGiftsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalGiftsCount.Get()
}

// GetTotalGiftsCountOk returns a tuple with the TotalGiftsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetTotalGiftsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalGiftsCount.Get(), o.TotalGiftsCount.IsSet()
}

// HasTotalGiftsCount returns a boolean if a field has been set.
func (o *ModelPost) HasTotalGiftsCount() bool {
	if o != nil && o.TotalGiftsCount.IsSet() {
		return true
	}

	return false
}

// SetTotalGiftsCount gets a reference to the given NullableInt32 and assigns it to the TotalGiftsCount field.
func (o *ModelPost) SetTotalGiftsCount(v int32) {
	o.TotalGiftsCount.Set(&v)
}
// SetTotalGiftsCountNil sets the value for TotalGiftsCount to be an explicit nil
func (o *ModelPost) SetTotalGiftsCountNil() {
	o.TotalGiftsCount.Set(nil)
}

// UnsetTotalGiftsCount ensures that no value is present for TotalGiftsCount, not even an explicit nil
func (o *ModelPost) UnsetTotalGiftsCount() {
	o.TotalGiftsCount.Unset()
}

// GetUpdatedAtSeconds returns the UpdatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetUpdatedAtSeconds() int64 {
	if o == nil || IsNil(o.UpdatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtSeconds.Get()
}

// GetUpdatedAtSecondsOk returns a tuple with the UpdatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetUpdatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtSeconds.Get(), o.UpdatedAtSeconds.IsSet()
}

// HasUpdatedAtSeconds returns a boolean if a field has been set.
func (o *ModelPost) HasUpdatedAtSeconds() bool {
	if o != nil && o.UpdatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the UpdatedAtSeconds field.
func (o *ModelPost) SetUpdatedAtSeconds(v int64) {
	o.UpdatedAtSeconds.Set(&v)
}
// SetUpdatedAtSecondsNil sets the value for UpdatedAtSeconds to be an explicit nil
func (o *ModelPost) SetUpdatedAtSecondsNil() {
	o.UpdatedAtSeconds.Set(nil)
}

// UnsetUpdatedAtSeconds ensures that no value is present for UpdatedAtSeconds, not even an explicit nil
func (o *ModelPost) UnsetUpdatedAtSeconds() {
	o.UpdatedAtSeconds.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetUser() User {
	if o == nil || IsNil(o.User.Get()) {
		var ret User
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ModelPost) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUser and assigns it to the User field.
func (o *ModelPost) SetUser(v User) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ModelPost) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ModelPost) UnsetUser() {
	o.User.Unset()
}

// GetVideos returns the Videos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPost) GetVideos() []ModelVideo {
	if o == nil {
		var ret []ModelVideo
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPost) GetVideosOk() ([]ModelVideo, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *ModelPost) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []ModelVideo and assigns it to the Videos field.
func (o *ModelPost) SetVideos(v []ModelVideo) {
	o.Videos = v
}

func (o ModelPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPost) ToMap() (map[string]interface{}, error) {
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
	if o.CategoryId.IsSet() {
		toSerialize["category_id"] = o.CategoryId.Get()
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
	if o.CreatedAtSeconds.IsSet() {
		toSerialize["created_at_seconds"] = o.CreatedAtSeconds.Get()
	}
	if o.EditedAtSeconds.IsSet() {
		toSerialize["edited_at_seconds"] = o.EditedAtSeconds.Get()
	}
	if o.FontSize.IsSet() {
		toSerialize["font_size"] = o.FontSize.Get()
	}
	if o.GameTitle.IsSet() {
		toSerialize["game_title"] = o.GameTitle.Get()
	}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if o.GroupUserTitle.IsSet() {
		toSerialize["group_user_title"] = o.GroupUserTitle.Get()
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
	if o.IsEdited.IsSet() {
		toSerialize["is_edited"] = o.IsEdited.Get()
	}
	if o.IsFailToSend.IsSet() {
		toSerialize["is_fail_to_send"] = o.IsFailToSend.Get()
	}
	if o.IsGroup.IsSet() {
		toSerialize["is_group"] = o.IsGroup.Get()
	}
	if o.IsHighlighted.IsSet() {
		toSerialize["is_highlighted"] = o.IsHighlighted.Get()
	}
	if o.IsHot.IsSet() {
		toSerialize["is_hot"] = o.IsHot.Get()
	}
	if o.IsIdFake.IsSet() {
		toSerialize["is_id_fake"] = o.IsIdFake.Get()
	}
	if o.IsLiked.IsSet() {
		toSerialize["is_liked"] = o.IsLiked.Get()
	}
	if o.IsMuted.IsSet() {
		toSerialize["is_muted"] = o.IsMuted.Get()
	}
	if o.IsPinned.IsSet() {
		toSerialize["is_pinned"] = o.IsPinned.Get()
	}
	if o.IsRecommended.IsSet() {
		toSerialize["is_recommended"] = o.IsRecommended.Get()
	}
	if o.IsRepostable.IsSet() {
		toSerialize["is_repostable"] = o.IsRepostable.Get()
	}
	if o.IsReposted.IsSet() {
		toSerialize["is_reposted"] = o.IsReposted.Get()
	}
	if o.IsRootOfConversation.IsSet() {
		toSerialize["is_root_of_conversation"] = o.IsRootOfConversation.Get()
	}
	if o.IsRootOfConversationComment.IsSet() {
		toSerialize["is_root_of_conversation_comment"] = o.IsRootOfConversationComment.Get()
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
	if o.PostType != nil {
		toSerialize["post_type"] = o.PostType
	}
	if o.ReportedCount.IsSet() {
		toSerialize["reported_count"] = o.ReportedCount.Get()
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
	if o.ShouldShowOriginalSource.IsSet() {
		toSerialize["should_show_original_source"] = o.ShouldShowOriginalSource.Get()
	}
	if o.ShouldShowSeeMoreConversation.IsSet() {
		toSerialize["should_show_see_more_conversation"] = o.ShouldShowSeeMoreConversation.Get()
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
	if o.UpdatedAtSeconds.IsSet() {
		toSerialize["updated_at_seconds"] = o.UpdatedAtSeconds.Get()
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

func (o *ModelPost) UnmarshalJSON(data []byte) (err error) {
	varModelPost := _ModelPost{}

	err = json.Unmarshal(data, &varModelPost)

	if err != nil {
		return err
	}

	*o = ModelPost(varModelPost)

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
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "color")
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "conversation_id")
		delete(additionalProperties, "created_at_seconds")
		delete(additionalProperties, "edited_at_seconds")
		delete(additionalProperties, "font_size")
		delete(additionalProperties, "game_title")
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "group")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "group_user_title")
		delete(additionalProperties, "id")
		delete(additionalProperties, "in_reply_to")
		delete(additionalProperties, "in_reply_to_post")
		delete(additionalProperties, "in_reply_to_post_count")
		delete(additionalProperties, "is_edited")
		delete(additionalProperties, "is_fail_to_send")
		delete(additionalProperties, "is_group")
		delete(additionalProperties, "is_highlighted")
		delete(additionalProperties, "is_hot")
		delete(additionalProperties, "is_id_fake")
		delete(additionalProperties, "is_liked")
		delete(additionalProperties, "is_muted")
		delete(additionalProperties, "is_pinned")
		delete(additionalProperties, "is_recommended")
		delete(additionalProperties, "is_repostable")
		delete(additionalProperties, "is_reposted")
		delete(additionalProperties, "is_root_of_conversation")
		delete(additionalProperties, "is_root_of_conversation_comment")
		delete(additionalProperties, "likers")
		delete(additionalProperties, "likers_count")
		delete(additionalProperties, "likes_count")
		delete(additionalProperties, "mentions")
		delete(additionalProperties, "message_tags")
		delete(additionalProperties, "post_type")
		delete(additionalProperties, "reported_count")
		delete(additionalProperties, "reposts_count")
		delete(additionalProperties, "shareable")
		delete(additionalProperties, "shared_thread")
		delete(additionalProperties, "shared_url")
		delete(additionalProperties, "should_show_original_source")
		delete(additionalProperties, "should_show_see_more_conversation")
		delete(additionalProperties, "survey")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "text")
		delete(additionalProperties, "thread")
		delete(additionalProperties, "thread_id")
		delete(additionalProperties, "total_gifts_count")
		delete(additionalProperties, "updated_at_seconds")
		delete(additionalProperties, "user")
		delete(additionalProperties, "videos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelPost struct {
	value *ModelPost
	isSet bool
}

func (v NullableModelPost) Get() *ModelPost {
	return v.value
}

func (v *NullableModelPost) Set(val *ModelPost) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPost) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPost(val *ModelPost) *NullableModelPost {
	return &NullableModelPost{value: val, isSet: true}
}

func (v NullableModelPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


