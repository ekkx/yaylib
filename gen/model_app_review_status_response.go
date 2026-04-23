
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AppReviewStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppReviewStatusResponse{}

// AppReviewStatusResponse struct for AppReviewStatusResponse
type AppReviewStatusResponse struct {
	IsAppReviewed NullableBool `json:"is_app_reviewed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppReviewStatusResponse AppReviewStatusResponse

// NewAppReviewStatusResponse instantiates a new AppReviewStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppReviewStatusResponse() *AppReviewStatusResponse {
	this := AppReviewStatusResponse{}
	return &this
}

// NewAppReviewStatusResponseWithDefaults instantiates a new AppReviewStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppReviewStatusResponseWithDefaults() *AppReviewStatusResponse {
	this := AppReviewStatusResponse{}
	return &this
}

// GetIsAppReviewed returns the IsAppReviewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppReviewStatusResponse) GetIsAppReviewed() bool {
	if o == nil || IsNil(o.IsAppReviewed.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAppReviewed.Get()
}

// GetIsAppReviewedOk returns a tuple with the IsAppReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppReviewStatusResponse) GetIsAppReviewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAppReviewed.Get(), o.IsAppReviewed.IsSet()
}

// HasIsAppReviewed returns a boolean if a field has been set.
func (o *AppReviewStatusResponse) HasIsAppReviewed() bool {
	if o != nil && o.IsAppReviewed.IsSet() {
		return true
	}

	return false
}

// SetIsAppReviewed gets a reference to the given NullableBool and assigns it to the IsAppReviewed field.
func (o *AppReviewStatusResponse) SetIsAppReviewed(v bool) {
	o.IsAppReviewed.Set(&v)
}
// SetIsAppReviewedNil sets the value for IsAppReviewed to be an explicit nil
func (o *AppReviewStatusResponse) SetIsAppReviewedNil() {
	o.IsAppReviewed.Set(nil)
}

// UnsetIsAppReviewed ensures that no value is present for IsAppReviewed, not even an explicit nil
func (o *AppReviewStatusResponse) UnsetIsAppReviewed() {
	o.IsAppReviewed.Unset()
}

func (o AppReviewStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppReviewStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAppReviewed.IsSet() {
		toSerialize["is_app_reviewed"] = o.IsAppReviewed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppReviewStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varAppReviewStatusResponse := _AppReviewStatusResponse{}

	err = json.Unmarshal(data, &varAppReviewStatusResponse)

	if err != nil {
		return err
	}

	*o = AppReviewStatusResponse(varAppReviewStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_app_reviewed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppReviewStatusResponse struct {
	value *AppReviewStatusResponse
	isSet bool
}

func (v NullableAppReviewStatusResponse) Get() *AppReviewStatusResponse {
	return v.value
}

func (v *NullableAppReviewStatusResponse) Set(val *AppReviewStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppReviewStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppReviewStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppReviewStatusResponse(val *AppReviewStatusResponse) *NullableAppReviewStatusResponse {
	return &NullableAppReviewStatusResponse{value: val, isSet: true}
}

func (v NullableAppReviewStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppReviewStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


