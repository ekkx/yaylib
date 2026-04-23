
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Game type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Game{}

// Game struct for Game
type Game struct {
	IconUrl NullableString `json:"icon_url,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	PlatformDetails NullablePlatformDetails `json:"platform_details,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Game Game

// NewGame instantiates a new Game object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGame() *Game {
	this := Game{}
	return &this
}

// NewGameWithDefaults instantiates a new Game object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameWithDefaults() *Game {
	this := Game{}
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Game) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Game) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *Game) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *Game) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *Game) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *Game) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Game) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Game) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Game) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Game) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Game) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Game) UnsetId() {
	o.Id.Unset()
}

// GetPlatformDetails returns the PlatformDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Game) GetPlatformDetails() PlatformDetails {
	if o == nil || IsNil(o.PlatformDetails.Get()) {
		var ret PlatformDetails
		return ret
	}
	return *o.PlatformDetails.Get()
}

// GetPlatformDetailsOk returns a tuple with the PlatformDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Game) GetPlatformDetailsOk() (*PlatformDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlatformDetails.Get(), o.PlatformDetails.IsSet()
}

// HasPlatformDetails returns a boolean if a field has been set.
func (o *Game) HasPlatformDetails() bool {
	if o != nil && o.PlatformDetails.IsSet() {
		return true
	}

	return false
}

// SetPlatformDetails gets a reference to the given NullablePlatformDetails and assigns it to the PlatformDetails field.
func (o *Game) SetPlatformDetails(v PlatformDetails) {
	o.PlatformDetails.Set(&v)
}
// SetPlatformDetailsNil sets the value for PlatformDetails to be an explicit nil
func (o *Game) SetPlatformDetailsNil() {
	o.PlatformDetails.Set(nil)
}

// UnsetPlatformDetails ensures that no value is present for PlatformDetails, not even an explicit nil
func (o *Game) UnsetPlatformDetails() {
	o.PlatformDetails.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Game) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Game) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Game) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Game) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Game) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Game) UnsetTitle() {
	o.Title.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Game) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Game) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Game) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Game) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Game) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Game) UnsetType() {
	o.Type.Unset()
}

func (o Game) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Game) ToMap() (map[string]interface{}, error) {
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
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Game) UnmarshalJSON(data []byte) (err error) {
	varGame := _Game{}

	err = json.Unmarshal(data, &varGame)

	if err != nil {
		return err
	}

	*o = Game(varGame)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "platform_details")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGame struct {
	value *Game
	isSet bool
}

func (v NullableGame) Get() *Game {
	return v.value
}

func (v *NullableGame) Set(val *Game) {
	v.value = val
	v.isSet = true
}

func (v NullableGame) IsSet() bool {
	return v.isSet
}

func (v *NullableGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGame(val *Game) *NullableGame {
	return &NullableGame{value: val, isSet: true}
}

func (v NullableGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


