
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Bgm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bgm{}

// Bgm struct for Bgm
type Bgm struct {
	Id NullableInt64 `json:"id,omitempty"`
	MusicUrl NullableString `json:"music_url,omitempty"`
	Order NullableInt32 `json:"order,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Bgm Bgm

// NewBgm instantiates a new Bgm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgm() *Bgm {
	this := Bgm{}
	return &this
}

// NewBgmWithDefaults instantiates a new Bgm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgmWithDefaults() *Bgm {
	this := Bgm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bgm) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bgm) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Bgm) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Bgm) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Bgm) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Bgm) UnsetId() {
	o.Id.Unset()
}

// GetMusicUrl returns the MusicUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bgm) GetMusicUrl() string {
	if o == nil || IsNil(o.MusicUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MusicUrl.Get()
}

// GetMusicUrlOk returns a tuple with the MusicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bgm) GetMusicUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MusicUrl.Get(), o.MusicUrl.IsSet()
}

// HasMusicUrl returns a boolean if a field has been set.
func (o *Bgm) HasMusicUrl() bool {
	if o != nil && o.MusicUrl.IsSet() {
		return true
	}

	return false
}

// SetMusicUrl gets a reference to the given NullableString and assigns it to the MusicUrl field.
func (o *Bgm) SetMusicUrl(v string) {
	o.MusicUrl.Set(&v)
}
// SetMusicUrlNil sets the value for MusicUrl to be an explicit nil
func (o *Bgm) SetMusicUrlNil() {
	o.MusicUrl.Set(nil)
}

// UnsetMusicUrl ensures that no value is present for MusicUrl, not even an explicit nil
func (o *Bgm) UnsetMusicUrl() {
	o.MusicUrl.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bgm) GetOrder() int32 {
	if o == nil || IsNil(o.Order.Get()) {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bgm) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *Bgm) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *Bgm) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *Bgm) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *Bgm) UnsetOrder() {
	o.Order.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bgm) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bgm) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Bgm) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Bgm) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Bgm) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Bgm) UnsetTitle() {
	o.Title.Unset()
}

func (o Bgm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bgm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MusicUrl.IsSet() {
		toSerialize["music_url"] = o.MusicUrl.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Bgm) UnmarshalJSON(data []byte) (err error) {
	varBgm := _Bgm{}

	err = json.Unmarshal(data, &varBgm)

	if err != nil {
		return err
	}

	*o = Bgm(varBgm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "music_url")
		delete(additionalProperties, "order")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgm struct {
	value *Bgm
	isSet bool
}

func (v NullableBgm) Get() *Bgm {
	return v.value
}

func (v *NullableBgm) Set(val *Bgm) {
	v.value = val
	v.isSet = true
}

func (v NullableBgm) IsSet() bool {
	return v.isSet
}

func (v *NullableBgm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgm(val *Bgm) *NullableBgm {
	return &NullableBgm{value: val, isSet: true}
}

func (v NullableBgm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


