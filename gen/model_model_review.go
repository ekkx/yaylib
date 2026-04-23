
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelReview{}

// ModelReview struct for ModelReview
type ModelReview struct {
	Comment NullableString `json:"comment,omitempty"`
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsMutual NullableBool `json:"is_mutual,omitempty"`
	IsPinned NullableBool `json:"is_pinned,omitempty"`
	ReportedCount NullableInt32 `json:"reported_count,omitempty"`
	Reviewer NullableUser `json:"reviewer,omitempty"`
	User NullableUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelReview ModelReview

// NewModelReview instantiates a new ModelReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelReview() *ModelReview {
	this := ModelReview{}
	return &this
}

// NewModelReviewWithDefaults instantiates a new ModelReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelReviewWithDefaults() *ModelReview {
	this := ModelReview{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ModelReview) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ModelReview) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ModelReview) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ModelReview) UnsetComment() {
	o.Comment.Unset()
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *ModelReview) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *ModelReview) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *ModelReview) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *ModelReview) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelReview) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelReview) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelReview) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelReview) UnsetId() {
	o.Id.Unset()
}

// GetIsMutual returns the IsMutual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetIsMutual() bool {
	if o == nil || IsNil(o.IsMutual.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMutual.Get()
}

// GetIsMutualOk returns a tuple with the IsMutual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetIsMutualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMutual.Get(), o.IsMutual.IsSet()
}

// HasIsMutual returns a boolean if a field has been set.
func (o *ModelReview) HasIsMutual() bool {
	if o != nil && o.IsMutual.IsSet() {
		return true
	}

	return false
}

// SetIsMutual gets a reference to the given NullableBool and assigns it to the IsMutual field.
func (o *ModelReview) SetIsMutual(v bool) {
	o.IsMutual.Set(&v)
}
// SetIsMutualNil sets the value for IsMutual to be an explicit nil
func (o *ModelReview) SetIsMutualNil() {
	o.IsMutual.Set(nil)
}

// UnsetIsMutual ensures that no value is present for IsMutual, not even an explicit nil
func (o *ModelReview) UnsetIsMutual() {
	o.IsMutual.Unset()
}

// GetIsPinned returns the IsPinned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetIsPinned() bool {
	if o == nil || IsNil(o.IsPinned.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPinned.Get()
}

// GetIsPinnedOk returns a tuple with the IsPinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPinned.Get(), o.IsPinned.IsSet()
}

// HasIsPinned returns a boolean if a field has been set.
func (o *ModelReview) HasIsPinned() bool {
	if o != nil && o.IsPinned.IsSet() {
		return true
	}

	return false
}

// SetIsPinned gets a reference to the given NullableBool and assigns it to the IsPinned field.
func (o *ModelReview) SetIsPinned(v bool) {
	o.IsPinned.Set(&v)
}
// SetIsPinnedNil sets the value for IsPinned to be an explicit nil
func (o *ModelReview) SetIsPinnedNil() {
	o.IsPinned.Set(nil)
}

// UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
func (o *ModelReview) UnsetIsPinned() {
	o.IsPinned.Unset()
}

// GetReportedCount returns the ReportedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetReportedCount() int32 {
	if o == nil || IsNil(o.ReportedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReportedCount.Get()
}

// GetReportedCountOk returns a tuple with the ReportedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetReportedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedCount.Get(), o.ReportedCount.IsSet()
}

// HasReportedCount returns a boolean if a field has been set.
func (o *ModelReview) HasReportedCount() bool {
	if o != nil && o.ReportedCount.IsSet() {
		return true
	}

	return false
}

// SetReportedCount gets a reference to the given NullableInt32 and assigns it to the ReportedCount field.
func (o *ModelReview) SetReportedCount(v int32) {
	o.ReportedCount.Set(&v)
}
// SetReportedCountNil sets the value for ReportedCount to be an explicit nil
func (o *ModelReview) SetReportedCountNil() {
	o.ReportedCount.Set(nil)
}

// UnsetReportedCount ensures that no value is present for ReportedCount, not even an explicit nil
func (o *ModelReview) UnsetReportedCount() {
	o.ReportedCount.Unset()
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetReviewer() User {
	if o == nil || IsNil(o.Reviewer.Get()) {
		var ret User
		return ret
	}
	return *o.Reviewer.Get()
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetReviewerOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reviewer.Get(), o.Reviewer.IsSet()
}

// HasReviewer returns a boolean if a field has been set.
func (o *ModelReview) HasReviewer() bool {
	if o != nil && o.Reviewer.IsSet() {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given NullableUser and assigns it to the Reviewer field.
func (o *ModelReview) SetReviewer(v User) {
	o.Reviewer.Set(&v)
}
// SetReviewerNil sets the value for Reviewer to be an explicit nil
func (o *ModelReview) SetReviewerNil() {
	o.Reviewer.Set(nil)
}

// UnsetReviewer ensures that no value is present for Reviewer, not even an explicit nil
func (o *ModelReview) UnsetReviewer() {
	o.Reviewer.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelReview) GetUser() User {
	if o == nil || IsNil(o.User.Get()) {
		var ret User
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelReview) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ModelReview) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUser and assigns it to the User field.
func (o *ModelReview) SetUser(v User) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ModelReview) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ModelReview) UnsetUser() {
	o.User.Unset()
}

func (o ModelReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsMutual.IsSet() {
		toSerialize["is_mutual"] = o.IsMutual.Get()
	}
	if o.IsPinned.IsSet() {
		toSerialize["is_pinned"] = o.IsPinned.Get()
	}
	if o.ReportedCount.IsSet() {
		toSerialize["reported_count"] = o.ReportedCount.Get()
	}
	if o.Reviewer.IsSet() {
		toSerialize["reviewer"] = o.Reviewer.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelReview) UnmarshalJSON(data []byte) (err error) {
	varModelReview := _ModelReview{}

	err = json.Unmarshal(data, &varModelReview)

	if err != nil {
		return err
	}

	*o = ModelReview(varModelReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_mutual")
		delete(additionalProperties, "is_pinned")
		delete(additionalProperties, "reported_count")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelReview struct {
	value *ModelReview
	isSet bool
}

func (v NullableModelReview) Get() *ModelReview {
	return v.value
}

func (v *NullableModelReview) Set(val *ModelReview) {
	v.value = val
	v.isSet = true
}

func (v NullableModelReview) IsSet() bool {
	return v.isSet
}

func (v *NullableModelReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelReview(val *ModelReview) *NullableModelReview {
	return &NullableModelReview{value: val, isSet: true}
}

func (v NullableModelReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


