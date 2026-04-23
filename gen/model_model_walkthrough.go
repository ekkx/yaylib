
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelWalkthrough type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelWalkthrough{}

// ModelWalkthrough struct for ModelWalkthrough
type ModelWalkthrough struct {
	ArticleId NullableInt64 `json:"article_id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelWalkthrough ModelWalkthrough

// NewModelWalkthrough instantiates a new ModelWalkthrough object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelWalkthrough() *ModelWalkthrough {
	this := ModelWalkthrough{}
	return &this
}

// NewModelWalkthroughWithDefaults instantiates a new ModelWalkthrough object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWalkthroughWithDefaults() *ModelWalkthrough {
	this := ModelWalkthrough{}
	return &this
}

// GetArticleId returns the ArticleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWalkthrough) GetArticleId() int64 {
	if o == nil || IsNil(o.ArticleId.Get()) {
		var ret int64
		return ret
	}
	return *o.ArticleId.Get()
}

// GetArticleIdOk returns a tuple with the ArticleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWalkthrough) GetArticleIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArticleId.Get(), o.ArticleId.IsSet()
}

// HasArticleId returns a boolean if a field has been set.
func (o *ModelWalkthrough) HasArticleId() bool {
	if o != nil && o.ArticleId.IsSet() {
		return true
	}

	return false
}

// SetArticleId gets a reference to the given NullableInt64 and assigns it to the ArticleId field.
func (o *ModelWalkthrough) SetArticleId(v int64) {
	o.ArticleId.Set(&v)
}
// SetArticleIdNil sets the value for ArticleId to be an explicit nil
func (o *ModelWalkthrough) SetArticleIdNil() {
	o.ArticleId.Set(nil)
}

// UnsetArticleId ensures that no value is present for ArticleId, not even an explicit nil
func (o *ModelWalkthrough) UnsetArticleId() {
	o.ArticleId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWalkthrough) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWalkthrough) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelWalkthrough) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ModelWalkthrough) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ModelWalkthrough) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ModelWalkthrough) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWalkthrough) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWalkthrough) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ModelWalkthrough) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ModelWalkthrough) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ModelWalkthrough) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ModelWalkthrough) UnsetUrl() {
	o.Url.Unset()
}

func (o ModelWalkthrough) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelWalkthrough) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ArticleId.IsSet() {
		toSerialize["article_id"] = o.ArticleId.Get()
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

func (o *ModelWalkthrough) UnmarshalJSON(data []byte) (err error) {
	varModelWalkthrough := _ModelWalkthrough{}

	err = json.Unmarshal(data, &varModelWalkthrough)

	if err != nil {
		return err
	}

	*o = ModelWalkthrough(varModelWalkthrough)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "article_id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelWalkthrough struct {
	value *ModelWalkthrough
	isSet bool
}

func (v NullableModelWalkthrough) Get() *ModelWalkthrough {
	return v.value
}

func (v *NullableModelWalkthrough) Set(val *ModelWalkthrough) {
	v.value = val
	v.isSet = true
}

func (v NullableModelWalkthrough) IsSet() bool {
	return v.isSet
}

func (v *NullableModelWalkthrough) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelWalkthrough(val *ModelWalkthrough) *NullableModelWalkthrough {
	return &NullableModelWalkthrough{value: val, isSet: true}
}

func (v NullableModelWalkthrough) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelWalkthrough) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


