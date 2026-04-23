
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ReviewsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewsResponse{}

// ReviewsResponse struct for ReviewsResponse
type ReviewsResponse struct {
	PinnedReviews []Review `json:"pinned_reviews,omitempty"`
	Reviews []Review `json:"reviews,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewsResponse ReviewsResponse

// NewReviewsResponse instantiates a new ReviewsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewsResponse() *ReviewsResponse {
	this := ReviewsResponse{}
	return &this
}

// NewReviewsResponseWithDefaults instantiates a new ReviewsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewsResponseWithDefaults() *ReviewsResponse {
	this := ReviewsResponse{}
	return &this
}

// GetPinnedReviews returns the PinnedReviews field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewsResponse) GetPinnedReviews() []Review {
	if o == nil {
		var ret []Review
		return ret
	}
	return o.PinnedReviews
}

// GetPinnedReviewsOk returns a tuple with the PinnedReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewsResponse) GetPinnedReviewsOk() ([]Review, bool) {
	if o == nil || IsNil(o.PinnedReviews) {
		return nil, false
	}
	return o.PinnedReviews, true
}

// HasPinnedReviews returns a boolean if a field has been set.
func (o *ReviewsResponse) HasPinnedReviews() bool {
	if o != nil && !IsNil(o.PinnedReviews) {
		return true
	}

	return false
}

// SetPinnedReviews gets a reference to the given []Review and assigns it to the PinnedReviews field.
func (o *ReviewsResponse) SetPinnedReviews(v []Review) {
	o.PinnedReviews = v
}

// GetReviews returns the Reviews field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewsResponse) GetReviews() []Review {
	if o == nil {
		var ret []Review
		return ret
	}
	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewsResponse) GetReviewsOk() ([]Review, bool) {
	if o == nil || IsNil(o.Reviews) {
		return nil, false
	}
	return o.Reviews, true
}

// HasReviews returns a boolean if a field has been set.
func (o *ReviewsResponse) HasReviews() bool {
	if o != nil && !IsNil(o.Reviews) {
		return true
	}

	return false
}

// SetReviews gets a reference to the given []Review and assigns it to the Reviews field.
func (o *ReviewsResponse) SetReviews(v []Review) {
	o.Reviews = v
}

func (o ReviewsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PinnedReviews != nil {
		toSerialize["pinned_reviews"] = o.PinnedReviews
	}
	if o.Reviews != nil {
		toSerialize["reviews"] = o.Reviews
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewsResponse) UnmarshalJSON(data []byte) (err error) {
	varReviewsResponse := _ReviewsResponse{}

	err = json.Unmarshal(data, &varReviewsResponse)

	if err != nil {
		return err
	}

	*o = ReviewsResponse(varReviewsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pinned_reviews")
		delete(additionalProperties, "reviews")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewsResponse struct {
	value *ReviewsResponse
	isSet bool
}

func (v NullableReviewsResponse) Get() *ReviewsResponse {
	return v.value
}

func (v *NullableReviewsResponse) Set(val *ReviewsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewsResponse(val *ReviewsResponse) *NullableReviewsResponse {
	return &NullableReviewsResponse{value: val, isSet: true}
}

func (v NullableReviewsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


