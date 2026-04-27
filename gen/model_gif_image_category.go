
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GifImageCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GifImageCategory{}

// GifImageCategory struct for GifImageCategory
type GifImageCategory struct {
	Gifs map[string]interface{} `json:"gifs,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Language NullableString `json:"language,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GifImageCategory GifImageCategory

// NewGifImageCategory instantiates a new GifImageCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGifImageCategory() *GifImageCategory {
	this := GifImageCategory{}
	return &this
}

// NewGifImageCategoryWithDefaults instantiates a new GifImageCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGifImageCategoryWithDefaults() *GifImageCategory {
	this := GifImageCategory{}
	return &this
}

// GetGifs returns the Gifs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GifImageCategory) GetGifs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Gifs
}

// GetGifsOk returns a tuple with the Gifs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GifImageCategory) GetGifsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Gifs) {
		return map[string]interface{}{}, false
	}
	return o.Gifs, true
}

// HasGifs returns a boolean if a field has been set.
func (o *GifImageCategory) HasGifs() bool {
	if o != nil && !IsNil(o.Gifs) {
		return true
	}

	return false
}

// SetGifs gets a reference to the given map[string]interface{} and assigns it to the Gifs field.
func (o *GifImageCategory) SetGifs(v map[string]interface{}) {
	o.Gifs = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GifImageCategory) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GifImageCategory) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GifImageCategory) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *GifImageCategory) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GifImageCategory) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GifImageCategory) UnsetId() {
	o.Id.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GifImageCategory) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GifImageCategory) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *GifImageCategory) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *GifImageCategory) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *GifImageCategory) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *GifImageCategory) UnsetLanguage() {
	o.Language.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GifImageCategory) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GifImageCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GifImageCategory) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GifImageCategory) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GifImageCategory) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GifImageCategory) UnsetName() {
	o.Name.Unset()
}

func (o GifImageCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GifImageCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gifs != nil {
		toSerialize["gifs"] = o.Gifs
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GifImageCategory) UnmarshalJSON(data []byte) (err error) {
	varGifImageCategory := _GifImageCategory{}

	err = json.Unmarshal(data, &varGifImageCategory)

	if err != nil {
		return err
	}

	*o = GifImageCategory(varGifImageCategory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifs")
		delete(additionalProperties, "id")
		delete(additionalProperties, "language")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGifImageCategory struct {
	value *GifImageCategory
	isSet bool
}

func (v NullableGifImageCategory) Get() *GifImageCategory {
	return v.value
}

func (v *NullableGifImageCategory) Set(val *GifImageCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableGifImageCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableGifImageCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGifImageCategory(val *GifImageCategory) *NullableGifImageCategory {
	return &NullableGifImageCategory{value: val, isSet: true}
}

func (v NullableGifImageCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGifImageCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


