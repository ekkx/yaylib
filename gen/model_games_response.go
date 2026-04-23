
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GamesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GamesResponse{}

// GamesResponse struct for GamesResponse
type GamesResponse struct {
	FromId NullableInt64 `json:"from_id,omitempty"`
	Games []RealmGame `json:"games,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GamesResponse GamesResponse

// NewGamesResponse instantiates a new GamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGamesResponse() *GamesResponse {
	this := GamesResponse{}
	return &this
}

// NewGamesResponseWithDefaults instantiates a new GamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGamesResponseWithDefaults() *GamesResponse {
	this := GamesResponse{}
	return &this
}

// GetFromId returns the FromId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GamesResponse) GetFromId() int64 {
	if o == nil || IsNil(o.FromId.Get()) {
		var ret int64
		return ret
	}
	return *o.FromId.Get()
}

// GetFromIdOk returns a tuple with the FromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GamesResponse) GetFromIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromId.Get(), o.FromId.IsSet()
}

// HasFromId returns a boolean if a field has been set.
func (o *GamesResponse) HasFromId() bool {
	if o != nil && o.FromId.IsSet() {
		return true
	}

	return false
}

// SetFromId gets a reference to the given NullableInt64 and assigns it to the FromId field.
func (o *GamesResponse) SetFromId(v int64) {
	o.FromId.Set(&v)
}
// SetFromIdNil sets the value for FromId to be an explicit nil
func (o *GamesResponse) SetFromIdNil() {
	o.FromId.Set(nil)
}

// UnsetFromId ensures that no value is present for FromId, not even an explicit nil
func (o *GamesResponse) UnsetFromId() {
	o.FromId.Unset()
}

// GetGames returns the Games field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GamesResponse) GetGames() []RealmGame {
	if o == nil {
		var ret []RealmGame
		return ret
	}
	return o.Games
}

// GetGamesOk returns a tuple with the Games field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GamesResponse) GetGamesOk() ([]RealmGame, bool) {
	if o == nil || IsNil(o.Games) {
		return nil, false
	}
	return o.Games, true
}

// HasGames returns a boolean if a field has been set.
func (o *GamesResponse) HasGames() bool {
	if o != nil && !IsNil(o.Games) {
		return true
	}

	return false
}

// SetGames gets a reference to the given []RealmGame and assigns it to the Games field.
func (o *GamesResponse) SetGames(v []RealmGame) {
	o.Games = v
}

func (o GamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GamesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FromId.IsSet() {
		toSerialize["from_id"] = o.FromId.Get()
	}
	if o.Games != nil {
		toSerialize["games"] = o.Games
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GamesResponse) UnmarshalJSON(data []byte) (err error) {
	varGamesResponse := _GamesResponse{}

	err = json.Unmarshal(data, &varGamesResponse)

	if err != nil {
		return err
	}

	*o = GamesResponse(varGamesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from_id")
		delete(additionalProperties, "games")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGamesResponse struct {
	value *GamesResponse
	isSet bool
}

func (v NullableGamesResponse) Get() *GamesResponse {
	return v.value
}

func (v *NullableGamesResponse) Set(val *GamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGamesResponse(val *GamesResponse) *NullableGamesResponse {
	return &NullableGamesResponse{value: val, isSet: true}
}

func (v NullableGamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


