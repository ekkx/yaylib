
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PresignedUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresignedUrl{}

// PresignedUrl struct for PresignedUrl
type PresignedUrl struct {
	Filename NullableString `json:"filename,omitempty"`
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PresignedUrl PresignedUrl

// NewPresignedUrl instantiates a new PresignedUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignedUrl() *PresignedUrl {
	this := PresignedUrl{}
	return &this
}

// NewPresignedUrlWithDefaults instantiates a new PresignedUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignedUrlWithDefaults() *PresignedUrl {
	this := PresignedUrl{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresignedUrl) GetFilename() string {
	if o == nil || IsNil(o.Filename.Get()) {
		var ret string
		return ret
	}
	return *o.Filename.Get()
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresignedUrl) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filename.Get(), o.Filename.IsSet()
}

// HasFilename returns a boolean if a field has been set.
func (o *PresignedUrl) HasFilename() bool {
	if o != nil && o.Filename.IsSet() {
		return true
	}

	return false
}

// SetFilename gets a reference to the given NullableString and assigns it to the Filename field.
func (o *PresignedUrl) SetFilename(v string) {
	o.Filename.Set(&v)
}
// SetFilenameNil sets the value for Filename to be an explicit nil
func (o *PresignedUrl) SetFilenameNil() {
	o.Filename.Set(nil)
}

// UnsetFilename ensures that no value is present for Filename, not even an explicit nil
func (o *PresignedUrl) UnsetFilename() {
	o.Filename.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresignedUrl) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresignedUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PresignedUrl) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PresignedUrl) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PresignedUrl) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PresignedUrl) UnsetUrl() {
	o.Url.Unset()
}

func (o PresignedUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresignedUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filename.IsSet() {
		toSerialize["filename"] = o.Filename.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PresignedUrl) UnmarshalJSON(data []byte) (err error) {
	varPresignedUrl := _PresignedUrl{}

	err = json.Unmarshal(data, &varPresignedUrl)

	if err != nil {
		return err
	}

	*o = PresignedUrl(varPresignedUrl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filename")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePresignedUrl struct {
	value *PresignedUrl
	isSet bool
}

func (v NullablePresignedUrl) Get() *PresignedUrl {
	return v.value
}

func (v *NullablePresignedUrl) Set(val *PresignedUrl) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignedUrl) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignedUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignedUrl(val *PresignedUrl) *NullablePresignedUrl {
	return &NullablePresignedUrl{value: val, isSet: true}
}

func (v NullablePresignedUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignedUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


