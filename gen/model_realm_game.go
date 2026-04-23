
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmGame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmGame{}

// RealmGame struct for RealmGame
type RealmGame struct {
	IconUrl NullableString `json:"icon_url,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	PlatformDetails NullableRealmPlatformDetails `json:"platform_details,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmGame RealmGame

// NewRealmGame instantiates a new RealmGame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmGame() *RealmGame {
	this := RealmGame{}
	return &this
}

// NewRealmGameWithDefaults instantiates a new RealmGame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmGameWithDefaults() *RealmGame {
	this := RealmGame{}
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGame) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGame) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *RealmGame) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *RealmGame) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *RealmGame) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *RealmGame) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGame) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGame) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmGame) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmGame) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmGame) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmGame) UnsetId() {
	o.Id.Unset()
}

// GetPlatformDetails returns the PlatformDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGame) GetPlatformDetails() RealmPlatformDetails {
	if o == nil || IsNil(o.PlatformDetails.Get()) {
		var ret RealmPlatformDetails
		return ret
	}
	return *o.PlatformDetails.Get()
}

// GetPlatformDetailsOk returns a tuple with the PlatformDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGame) GetPlatformDetailsOk() (*RealmPlatformDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlatformDetails.Get(), o.PlatformDetails.IsSet()
}

// HasPlatformDetails returns a boolean if a field has been set.
func (o *RealmGame) HasPlatformDetails() bool {
	if o != nil && o.PlatformDetails.IsSet() {
		return true
	}

	return false
}

// SetPlatformDetails gets a reference to the given NullableRealmPlatformDetails and assigns it to the PlatformDetails field.
func (o *RealmGame) SetPlatformDetails(v RealmPlatformDetails) {
	o.PlatformDetails.Set(&v)
}
// SetPlatformDetailsNil sets the value for PlatformDetails to be an explicit nil
func (o *RealmGame) SetPlatformDetailsNil() {
	o.PlatformDetails.Set(nil)
}

// UnsetPlatformDetails ensures that no value is present for PlatformDetails, not even an explicit nil
func (o *RealmGame) UnsetPlatformDetails() {
	o.PlatformDetails.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGame) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGame) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RealmGame) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RealmGame) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RealmGame) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RealmGame) UnsetTitle() {
	o.Title.Unset()
}

func (o RealmGame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmGame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.PlatformDetails.IsSet() {
		toSerialize["platform_details"] = o.PlatformDetails.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmGame) UnmarshalJSON(data []byte) (err error) {
	varRealmGame := _RealmGame{}

	err = json.Unmarshal(data, &varRealmGame)

	if err != nil {
		return err
	}

	*o = RealmGame(varRealmGame)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "platform_details")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmGame struct {
	value *RealmGame
	isSet bool
}

func (v NullableRealmGame) Get() *RealmGame {
	return v.value
}

func (v *NullableRealmGame) Set(val *RealmGame) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmGame) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmGame(val *RealmGame) *NullableRealmGame {
	return &NullableRealmGame{value: val, isSet: true}
}

func (v NullableRealmGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


