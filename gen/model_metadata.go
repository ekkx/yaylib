
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata struct for Metadata
type Metadata struct {
	Anonymous NullableBool `json:"anonymous,omitempty"`
	Body NullableString `json:"body,omitempty"`
	BulkInvitation NullableBool `json:"bulk_invitation,omitempty"`
	ContentPreview NullableString `json:"content_preview,omitempty"`
	GroupTopic NullableString `json:"group_topic,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Metadata Metadata

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetAnonymous() bool {
	if o == nil || IsNil(o.Anonymous.Get()) {
		var ret bool
		return ret
	}
	return *o.Anonymous.Get()
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Anonymous.Get(), o.Anonymous.IsSet()
}

// HasAnonymous returns a boolean if a field has been set.
func (o *Metadata) HasAnonymous() bool {
	if o != nil && o.Anonymous.IsSet() {
		return true
	}

	return false
}

// SetAnonymous gets a reference to the given NullableBool and assigns it to the Anonymous field.
func (o *Metadata) SetAnonymous(v bool) {
	o.Anonymous.Set(&v)
}
// SetAnonymousNil sets the value for Anonymous to be an explicit nil
func (o *Metadata) SetAnonymousNil() {
	o.Anonymous.Set(nil)
}

// UnsetAnonymous ensures that no value is present for Anonymous, not even an explicit nil
func (o *Metadata) UnsetAnonymous() {
	o.Anonymous.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *Metadata) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *Metadata) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *Metadata) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *Metadata) UnsetBody() {
	o.Body.Unset()
}

// GetBulkInvitation returns the BulkInvitation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetBulkInvitation() bool {
	if o == nil || IsNil(o.BulkInvitation.Get()) {
		var ret bool
		return ret
	}
	return *o.BulkInvitation.Get()
}

// GetBulkInvitationOk returns a tuple with the BulkInvitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetBulkInvitationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BulkInvitation.Get(), o.BulkInvitation.IsSet()
}

// HasBulkInvitation returns a boolean if a field has been set.
func (o *Metadata) HasBulkInvitation() bool {
	if o != nil && o.BulkInvitation.IsSet() {
		return true
	}

	return false
}

// SetBulkInvitation gets a reference to the given NullableBool and assigns it to the BulkInvitation field.
func (o *Metadata) SetBulkInvitation(v bool) {
	o.BulkInvitation.Set(&v)
}
// SetBulkInvitationNil sets the value for BulkInvitation to be an explicit nil
func (o *Metadata) SetBulkInvitationNil() {
	o.BulkInvitation.Set(nil)
}

// UnsetBulkInvitation ensures that no value is present for BulkInvitation, not even an explicit nil
func (o *Metadata) UnsetBulkInvitation() {
	o.BulkInvitation.Unset()
}

// GetContentPreview returns the ContentPreview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetContentPreview() string {
	if o == nil || IsNil(o.ContentPreview.Get()) {
		var ret string
		return ret
	}
	return *o.ContentPreview.Get()
}

// GetContentPreviewOk returns a tuple with the ContentPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetContentPreviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentPreview.Get(), o.ContentPreview.IsSet()
}

// HasContentPreview returns a boolean if a field has been set.
func (o *Metadata) HasContentPreview() bool {
	if o != nil && o.ContentPreview.IsSet() {
		return true
	}

	return false
}

// SetContentPreview gets a reference to the given NullableString and assigns it to the ContentPreview field.
func (o *Metadata) SetContentPreview(v string) {
	o.ContentPreview.Set(&v)
}
// SetContentPreviewNil sets the value for ContentPreview to be an explicit nil
func (o *Metadata) SetContentPreviewNil() {
	o.ContentPreview.Set(nil)
}

// UnsetContentPreview ensures that no value is present for ContentPreview, not even an explicit nil
func (o *Metadata) UnsetContentPreview() {
	o.ContentPreview.Unset()
}

// GetGroupTopic returns the GroupTopic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetGroupTopic() string {
	if o == nil || IsNil(o.GroupTopic.Get()) {
		var ret string
		return ret
	}
	return *o.GroupTopic.Get()
}

// GetGroupTopicOk returns a tuple with the GroupTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetGroupTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupTopic.Get(), o.GroupTopic.IsSet()
}

// HasGroupTopic returns a boolean if a field has been set.
func (o *Metadata) HasGroupTopic() bool {
	if o != nil && o.GroupTopic.IsSet() {
		return true
	}

	return false
}

// SetGroupTopic gets a reference to the given NullableString and assigns it to the GroupTopic field.
func (o *Metadata) SetGroupTopic(v string) {
	o.GroupTopic.Set(&v)
}
// SetGroupTopicNil sets the value for GroupTopic to be an explicit nil
func (o *Metadata) SetGroupTopicNil() {
	o.GroupTopic.Set(nil)
}

// UnsetGroupTopic ensures that no value is present for GroupTopic, not even an explicit nil
func (o *Metadata) UnsetGroupTopic() {
	o.GroupTopic.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Metadata) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Metadata) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Metadata) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Metadata) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Metadata) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Metadata) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Metadata) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Metadata) UnsetUrl() {
	o.Url.Unset()
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Anonymous.IsSet() {
		toSerialize["anonymous"] = o.Anonymous.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.BulkInvitation.IsSet() {
		toSerialize["bulk_invitation"] = o.BulkInvitation.Get()
	}
	if o.ContentPreview.IsSet() {
		toSerialize["content_preview"] = o.ContentPreview.Get()
	}
	if o.GroupTopic.IsSet() {
		toSerialize["group_topic"] = o.GroupTopic.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Metadata) UnmarshalJSON(data []byte) (err error) {
	varMetadata := _Metadata{}

	err = json.Unmarshal(data, &varMetadata)

	if err != nil {
		return err
	}

	*o = Metadata(varMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "anonymous")
		delete(additionalProperties, "body")
		delete(additionalProperties, "bulk_invitation")
		delete(additionalProperties, "content_preview")
		delete(additionalProperties, "group_topic")
		delete(additionalProperties, "title")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


