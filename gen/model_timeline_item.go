
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TimelineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineItem{}

// TimelineItem struct for TimelineItem
type TimelineItem struct {
	Categories []GroupCategory `json:"categories,omitempty"`
	IsRepliedByNextPost NullableBool `json:"is_replied_by_next_post,omitempty"`
	IsReplyingToPrevPost NullableBool `json:"is_replying_to_prev_post,omitempty"`
	Post NullableModelPost `json:"post,omitempty"`
	ShouldShowConversationLine NullableBool `json:"should_show_conversation_line,omitempty"`
	ShouldShowMentionBox NullableBool `json:"should_show_mention_box,omitempty"`
	ShouldShowRepostButton NullableBool `json:"should_show_repost_button,omitempty"`
	ShouldShowTranslateButton NullableBool `json:"should_show_translate_button,omitempty"`
	ShouldShowViewReposts NullableBool `json:"should_show_view_reposts,omitempty"`
	TranslatedText NullableString `json:"translated_text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelineItem TimelineItem

// NewTimelineItem instantiates a new TimelineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineItem() *TimelineItem {
	this := TimelineItem{}
	return &this
}

// NewTimelineItemWithDefaults instantiates a new TimelineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineItemWithDefaults() *TimelineItem {
	this := TimelineItem{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetCategories() []GroupCategory {
	if o == nil {
		var ret []GroupCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetCategoriesOk() ([]GroupCategory, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *TimelineItem) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []GroupCategory and assigns it to the Categories field.
func (o *TimelineItem) SetCategories(v []GroupCategory) {
	o.Categories = v
}

// GetIsRepliedByNextPost returns the IsRepliedByNextPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetIsRepliedByNextPost() bool {
	if o == nil || IsNil(o.IsRepliedByNextPost.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRepliedByNextPost.Get()
}

// GetIsRepliedByNextPostOk returns a tuple with the IsRepliedByNextPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetIsRepliedByNextPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRepliedByNextPost.Get(), o.IsRepliedByNextPost.IsSet()
}

// HasIsRepliedByNextPost returns a boolean if a field has been set.
func (o *TimelineItem) HasIsRepliedByNextPost() bool {
	if o != nil && o.IsRepliedByNextPost.IsSet() {
		return true
	}

	return false
}

// SetIsRepliedByNextPost gets a reference to the given NullableBool and assigns it to the IsRepliedByNextPost field.
func (o *TimelineItem) SetIsRepliedByNextPost(v bool) {
	o.IsRepliedByNextPost.Set(&v)
}
// SetIsRepliedByNextPostNil sets the value for IsRepliedByNextPost to be an explicit nil
func (o *TimelineItem) SetIsRepliedByNextPostNil() {
	o.IsRepliedByNextPost.Set(nil)
}

// UnsetIsRepliedByNextPost ensures that no value is present for IsRepliedByNextPost, not even an explicit nil
func (o *TimelineItem) UnsetIsRepliedByNextPost() {
	o.IsRepliedByNextPost.Unset()
}

// GetIsReplyingToPrevPost returns the IsReplyingToPrevPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetIsReplyingToPrevPost() bool {
	if o == nil || IsNil(o.IsReplyingToPrevPost.Get()) {
		var ret bool
		return ret
	}
	return *o.IsReplyingToPrevPost.Get()
}

// GetIsReplyingToPrevPostOk returns a tuple with the IsReplyingToPrevPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetIsReplyingToPrevPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsReplyingToPrevPost.Get(), o.IsReplyingToPrevPost.IsSet()
}

// HasIsReplyingToPrevPost returns a boolean if a field has been set.
func (o *TimelineItem) HasIsReplyingToPrevPost() bool {
	if o != nil && o.IsReplyingToPrevPost.IsSet() {
		return true
	}

	return false
}

// SetIsReplyingToPrevPost gets a reference to the given NullableBool and assigns it to the IsReplyingToPrevPost field.
func (o *TimelineItem) SetIsReplyingToPrevPost(v bool) {
	o.IsReplyingToPrevPost.Set(&v)
}
// SetIsReplyingToPrevPostNil sets the value for IsReplyingToPrevPost to be an explicit nil
func (o *TimelineItem) SetIsReplyingToPrevPostNil() {
	o.IsReplyingToPrevPost.Set(nil)
}

// UnsetIsReplyingToPrevPost ensures that no value is present for IsReplyingToPrevPost, not even an explicit nil
func (o *TimelineItem) UnsetIsReplyingToPrevPost() {
	o.IsReplyingToPrevPost.Unset()
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetPost() ModelPost {
	if o == nil || IsNil(o.Post.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *TimelineItem) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullableModelPost and assigns it to the Post field.
func (o *TimelineItem) SetPost(v ModelPost) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *TimelineItem) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *TimelineItem) UnsetPost() {
	o.Post.Unset()
}

// GetShouldShowConversationLine returns the ShouldShowConversationLine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetShouldShowConversationLine() bool {
	if o == nil || IsNil(o.ShouldShowConversationLine.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowConversationLine.Get()
}

// GetShouldShowConversationLineOk returns a tuple with the ShouldShowConversationLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetShouldShowConversationLineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowConversationLine.Get(), o.ShouldShowConversationLine.IsSet()
}

// HasShouldShowConversationLine returns a boolean if a field has been set.
func (o *TimelineItem) HasShouldShowConversationLine() bool {
	if o != nil && o.ShouldShowConversationLine.IsSet() {
		return true
	}

	return false
}

// SetShouldShowConversationLine gets a reference to the given NullableBool and assigns it to the ShouldShowConversationLine field.
func (o *TimelineItem) SetShouldShowConversationLine(v bool) {
	o.ShouldShowConversationLine.Set(&v)
}
// SetShouldShowConversationLineNil sets the value for ShouldShowConversationLine to be an explicit nil
func (o *TimelineItem) SetShouldShowConversationLineNil() {
	o.ShouldShowConversationLine.Set(nil)
}

// UnsetShouldShowConversationLine ensures that no value is present for ShouldShowConversationLine, not even an explicit nil
func (o *TimelineItem) UnsetShouldShowConversationLine() {
	o.ShouldShowConversationLine.Unset()
}

// GetShouldShowMentionBox returns the ShouldShowMentionBox field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetShouldShowMentionBox() bool {
	if o == nil || IsNil(o.ShouldShowMentionBox.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowMentionBox.Get()
}

// GetShouldShowMentionBoxOk returns a tuple with the ShouldShowMentionBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetShouldShowMentionBoxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowMentionBox.Get(), o.ShouldShowMentionBox.IsSet()
}

// HasShouldShowMentionBox returns a boolean if a field has been set.
func (o *TimelineItem) HasShouldShowMentionBox() bool {
	if o != nil && o.ShouldShowMentionBox.IsSet() {
		return true
	}

	return false
}

// SetShouldShowMentionBox gets a reference to the given NullableBool and assigns it to the ShouldShowMentionBox field.
func (o *TimelineItem) SetShouldShowMentionBox(v bool) {
	o.ShouldShowMentionBox.Set(&v)
}
// SetShouldShowMentionBoxNil sets the value for ShouldShowMentionBox to be an explicit nil
func (o *TimelineItem) SetShouldShowMentionBoxNil() {
	o.ShouldShowMentionBox.Set(nil)
}

// UnsetShouldShowMentionBox ensures that no value is present for ShouldShowMentionBox, not even an explicit nil
func (o *TimelineItem) UnsetShouldShowMentionBox() {
	o.ShouldShowMentionBox.Unset()
}

// GetShouldShowRepostButton returns the ShouldShowRepostButton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetShouldShowRepostButton() bool {
	if o == nil || IsNil(o.ShouldShowRepostButton.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowRepostButton.Get()
}

// GetShouldShowRepostButtonOk returns a tuple with the ShouldShowRepostButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetShouldShowRepostButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowRepostButton.Get(), o.ShouldShowRepostButton.IsSet()
}

// HasShouldShowRepostButton returns a boolean if a field has been set.
func (o *TimelineItem) HasShouldShowRepostButton() bool {
	if o != nil && o.ShouldShowRepostButton.IsSet() {
		return true
	}

	return false
}

// SetShouldShowRepostButton gets a reference to the given NullableBool and assigns it to the ShouldShowRepostButton field.
func (o *TimelineItem) SetShouldShowRepostButton(v bool) {
	o.ShouldShowRepostButton.Set(&v)
}
// SetShouldShowRepostButtonNil sets the value for ShouldShowRepostButton to be an explicit nil
func (o *TimelineItem) SetShouldShowRepostButtonNil() {
	o.ShouldShowRepostButton.Set(nil)
}

// UnsetShouldShowRepostButton ensures that no value is present for ShouldShowRepostButton, not even an explicit nil
func (o *TimelineItem) UnsetShouldShowRepostButton() {
	o.ShouldShowRepostButton.Unset()
}

// GetShouldShowTranslateButton returns the ShouldShowTranslateButton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetShouldShowTranslateButton() bool {
	if o == nil || IsNil(o.ShouldShowTranslateButton.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowTranslateButton.Get()
}

// GetShouldShowTranslateButtonOk returns a tuple with the ShouldShowTranslateButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetShouldShowTranslateButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowTranslateButton.Get(), o.ShouldShowTranslateButton.IsSet()
}

// HasShouldShowTranslateButton returns a boolean if a field has been set.
func (o *TimelineItem) HasShouldShowTranslateButton() bool {
	if o != nil && o.ShouldShowTranslateButton.IsSet() {
		return true
	}

	return false
}

// SetShouldShowTranslateButton gets a reference to the given NullableBool and assigns it to the ShouldShowTranslateButton field.
func (o *TimelineItem) SetShouldShowTranslateButton(v bool) {
	o.ShouldShowTranslateButton.Set(&v)
}
// SetShouldShowTranslateButtonNil sets the value for ShouldShowTranslateButton to be an explicit nil
func (o *TimelineItem) SetShouldShowTranslateButtonNil() {
	o.ShouldShowTranslateButton.Set(nil)
}

// UnsetShouldShowTranslateButton ensures that no value is present for ShouldShowTranslateButton, not even an explicit nil
func (o *TimelineItem) UnsetShouldShowTranslateButton() {
	o.ShouldShowTranslateButton.Unset()
}

// GetShouldShowViewReposts returns the ShouldShowViewReposts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetShouldShowViewReposts() bool {
	if o == nil || IsNil(o.ShouldShowViewReposts.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldShowViewReposts.Get()
}

// GetShouldShowViewRepostsOk returns a tuple with the ShouldShowViewReposts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetShouldShowViewRepostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldShowViewReposts.Get(), o.ShouldShowViewReposts.IsSet()
}

// HasShouldShowViewReposts returns a boolean if a field has been set.
func (o *TimelineItem) HasShouldShowViewReposts() bool {
	if o != nil && o.ShouldShowViewReposts.IsSet() {
		return true
	}

	return false
}

// SetShouldShowViewReposts gets a reference to the given NullableBool and assigns it to the ShouldShowViewReposts field.
func (o *TimelineItem) SetShouldShowViewReposts(v bool) {
	o.ShouldShowViewReposts.Set(&v)
}
// SetShouldShowViewRepostsNil sets the value for ShouldShowViewReposts to be an explicit nil
func (o *TimelineItem) SetShouldShowViewRepostsNil() {
	o.ShouldShowViewReposts.Set(nil)
}

// UnsetShouldShowViewReposts ensures that no value is present for ShouldShowViewReposts, not even an explicit nil
func (o *TimelineItem) UnsetShouldShowViewReposts() {
	o.ShouldShowViewReposts.Unset()
}

// GetTranslatedText returns the TranslatedText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineItem) GetTranslatedText() string {
	if o == nil || IsNil(o.TranslatedText.Get()) {
		var ret string
		return ret
	}
	return *o.TranslatedText.Get()
}

// GetTranslatedTextOk returns a tuple with the TranslatedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineItem) GetTranslatedTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TranslatedText.Get(), o.TranslatedText.IsSet()
}

// HasTranslatedText returns a boolean if a field has been set.
func (o *TimelineItem) HasTranslatedText() bool {
	if o != nil && o.TranslatedText.IsSet() {
		return true
	}

	return false
}

// SetTranslatedText gets a reference to the given NullableString and assigns it to the TranslatedText field.
func (o *TimelineItem) SetTranslatedText(v string) {
	o.TranslatedText.Set(&v)
}
// SetTranslatedTextNil sets the value for TranslatedText to be an explicit nil
func (o *TimelineItem) SetTranslatedTextNil() {
	o.TranslatedText.Set(nil)
}

// UnsetTranslatedText ensures that no value is present for TranslatedText, not even an explicit nil
func (o *TimelineItem) UnsetTranslatedText() {
	o.TranslatedText.Unset()
}

func (o TimelineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.IsRepliedByNextPost.IsSet() {
		toSerialize["is_replied_by_next_post"] = o.IsRepliedByNextPost.Get()
	}
	if o.IsReplyingToPrevPost.IsSet() {
		toSerialize["is_replying_to_prev_post"] = o.IsReplyingToPrevPost.Get()
	}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}
	if o.ShouldShowConversationLine.IsSet() {
		toSerialize["should_show_conversation_line"] = o.ShouldShowConversationLine.Get()
	}
	if o.ShouldShowMentionBox.IsSet() {
		toSerialize["should_show_mention_box"] = o.ShouldShowMentionBox.Get()
	}
	if o.ShouldShowRepostButton.IsSet() {
		toSerialize["should_show_repost_button"] = o.ShouldShowRepostButton.Get()
	}
	if o.ShouldShowTranslateButton.IsSet() {
		toSerialize["should_show_translate_button"] = o.ShouldShowTranslateButton.Get()
	}
	if o.ShouldShowViewReposts.IsSet() {
		toSerialize["should_show_view_reposts"] = o.ShouldShowViewReposts.Get()
	}
	if o.TranslatedText.IsSet() {
		toSerialize["translated_text"] = o.TranslatedText.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelineItem) UnmarshalJSON(data []byte) (err error) {
	varTimelineItem := _TimelineItem{}

	err = json.Unmarshal(data, &varTimelineItem)

	if err != nil {
		return err
	}

	*o = TimelineItem(varTimelineItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "categories")
		delete(additionalProperties, "is_replied_by_next_post")
		delete(additionalProperties, "is_replying_to_prev_post")
		delete(additionalProperties, "post")
		delete(additionalProperties, "should_show_conversation_line")
		delete(additionalProperties, "should_show_mention_box")
		delete(additionalProperties, "should_show_repost_button")
		delete(additionalProperties, "should_show_translate_button")
		delete(additionalProperties, "should_show_view_reposts")
		delete(additionalProperties, "translated_text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelineItem struct {
	value *TimelineItem
	isSet bool
}

func (v NullableTimelineItem) Get() *TimelineItem {
	return v.value
}

func (v *NullableTimelineItem) Set(val *TimelineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineItem(val *TimelineItem) *NullableTimelineItem {
	return &NullableTimelineItem{value: val, isSet: true}
}

func (v NullableTimelineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


