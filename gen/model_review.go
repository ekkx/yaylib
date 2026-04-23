
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Review type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Review{}

// Review struct for Review
type Review struct {
	Comment NullableString `json:"comment,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MutualReview NullableBool `json:"mutual_review,omitempty"`
	ReportedCount NullableInt32 `json:"reported_count,omitempty"`
	Reviewer NullableRealmUser `json:"reviewer,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Review Review

// NewReview instantiates a new Review object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReview() *Review {
	this := Review{}
	return &this
}

// NewReviewWithDefaults instantiates a new Review object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewWithDefaults() *Review {
	this := Review{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Review) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Review) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Review) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Review) UnsetComment() {
	o.Comment.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Review) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *Review) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Review) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Review) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Review) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Review) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Review) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Review) UnsetId() {
	o.Id.Unset()
}

// GetMutualReview returns the MutualReview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetMutualReview() bool {
	if o == nil || IsNil(o.MutualReview.Get()) {
		var ret bool
		return ret
	}
	return *o.MutualReview.Get()
}

// GetMutualReviewOk returns a tuple with the MutualReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetMutualReviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MutualReview.Get(), o.MutualReview.IsSet()
}

// HasMutualReview returns a boolean if a field has been set.
func (o *Review) HasMutualReview() bool {
	if o != nil && o.MutualReview.IsSet() {
		return true
	}

	return false
}

// SetMutualReview gets a reference to the given NullableBool and assigns it to the MutualReview field.
func (o *Review) SetMutualReview(v bool) {
	o.MutualReview.Set(&v)
}
// SetMutualReviewNil sets the value for MutualReview to be an explicit nil
func (o *Review) SetMutualReviewNil() {
	o.MutualReview.Set(nil)
}

// UnsetMutualReview ensures that no value is present for MutualReview, not even an explicit nil
func (o *Review) UnsetMutualReview() {
	o.MutualReview.Unset()
}

// GetReportedCount returns the ReportedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetReportedCount() int32 {
	if o == nil || IsNil(o.ReportedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReportedCount.Get()
}

// GetReportedCountOk returns a tuple with the ReportedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetReportedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedCount.Get(), o.ReportedCount.IsSet()
}

// HasReportedCount returns a boolean if a field has been set.
func (o *Review) HasReportedCount() bool {
	if o != nil && o.ReportedCount.IsSet() {
		return true
	}

	return false
}

// SetReportedCount gets a reference to the given NullableInt32 and assigns it to the ReportedCount field.
func (o *Review) SetReportedCount(v int32) {
	o.ReportedCount.Set(&v)
}
// SetReportedCountNil sets the value for ReportedCount to be an explicit nil
func (o *Review) SetReportedCountNil() {
	o.ReportedCount.Set(nil)
}

// UnsetReportedCount ensures that no value is present for ReportedCount, not even an explicit nil
func (o *Review) UnsetReportedCount() {
	o.ReportedCount.Unset()
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetReviewer() RealmUser {
	if o == nil || IsNil(o.Reviewer.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Reviewer.Get()
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetReviewerOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reviewer.Get(), o.Reviewer.IsSet()
}

// HasReviewer returns a boolean if a field has been set.
func (o *Review) HasReviewer() bool {
	if o != nil && o.Reviewer.IsSet() {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given NullableRealmUser and assigns it to the Reviewer field.
func (o *Review) SetReviewer(v RealmUser) {
	o.Reviewer.Set(&v)
}
// SetReviewerNil sets the value for Reviewer to be an explicit nil
func (o *Review) SetReviewerNil() {
	o.Reviewer.Set(nil)
}

// UnsetReviewer ensures that no value is present for Reviewer, not even an explicit nil
func (o *Review) UnsetReviewer() {
	o.Reviewer.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Review) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Review) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Review) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *Review) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Review) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Review) UnsetUser() {
	o.User.Unset()
}

func (o Review) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Review) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MutualReview.IsSet() {
		toSerialize["mutual_review"] = o.MutualReview.Get()
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

func (o *Review) UnmarshalJSON(data []byte) (err error) {
	varReview := _Review{}

	err = json.Unmarshal(data, &varReview)

	if err != nil {
		return err
	}

	*o = Review(varReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mutual_review")
		delete(additionalProperties, "reported_count")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReview struct {
	value *Review
	isSet bool
}

func (v NullableReview) Get() *Review {
	return v.value
}

func (v *NullableReview) Set(val *Review) {
	v.value = val
	v.isSet = true
}

func (v NullableReview) IsSet() bool {
	return v.isSet
}

func (v *NullableReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReview(val *Review) *NullableReview {
	return &NullableReview{value: val, isSet: true}
}

func (v NullableReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


