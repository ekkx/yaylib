
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FootprintDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FootprintDTO{}

// FootprintDTO struct for FootprintDTO
type FootprintDTO struct {
	Id NullableInt64 `json:"id,omitempty"`
	User NullableUserUserDTO `json:"user,omitempty"`
	VisitedAt NullableInt64 `json:"visited_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FootprintDTO FootprintDTO

// NewFootprintDTO instantiates a new FootprintDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFootprintDTO() *FootprintDTO {
	this := FootprintDTO{}
	return &this
}

// NewFootprintDTOWithDefaults instantiates a new FootprintDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFootprintDTOWithDefaults() *FootprintDTO {
	this := FootprintDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FootprintDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FootprintDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FootprintDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *FootprintDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *FootprintDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FootprintDTO) UnsetId() {
	o.Id.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FootprintDTO) GetUser() UserUserDTO {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserUserDTO
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FootprintDTO) GetUserOk() (*UserUserDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *FootprintDTO) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserUserDTO and assigns it to the User field.
func (o *FootprintDTO) SetUser(v UserUserDTO) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *FootprintDTO) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *FootprintDTO) UnsetUser() {
	o.User.Unset()
}

// GetVisitedAt returns the VisitedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FootprintDTO) GetVisitedAt() int64 {
	if o == nil || IsNil(o.VisitedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.VisitedAt.Get()
}

// GetVisitedAtOk returns a tuple with the VisitedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FootprintDTO) GetVisitedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisitedAt.Get(), o.VisitedAt.IsSet()
}

// HasVisitedAt returns a boolean if a field has been set.
func (o *FootprintDTO) HasVisitedAt() bool {
	if o != nil && o.VisitedAt.IsSet() {
		return true
	}

	return false
}

// SetVisitedAt gets a reference to the given NullableInt64 and assigns it to the VisitedAt field.
func (o *FootprintDTO) SetVisitedAt(v int64) {
	o.VisitedAt.Set(&v)
}
// SetVisitedAtNil sets the value for VisitedAt to be an explicit nil
func (o *FootprintDTO) SetVisitedAtNil() {
	o.VisitedAt.Set(nil)
}

// UnsetVisitedAt ensures that no value is present for VisitedAt, not even an explicit nil
func (o *FootprintDTO) UnsetVisitedAt() {
	o.VisitedAt.Unset()
}

func (o FootprintDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FootprintDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.VisitedAt.IsSet() {
		toSerialize["visited_at"] = o.VisitedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FootprintDTO) UnmarshalJSON(data []byte) (err error) {
	varFootprintDTO := _FootprintDTO{}

	err = json.Unmarshal(data, &varFootprintDTO)

	if err != nil {
		return err
	}

	*o = FootprintDTO(varFootprintDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user")
		delete(additionalProperties, "visited_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFootprintDTO struct {
	value *FootprintDTO
	isSet bool
}

func (v NullableFootprintDTO) Get() *FootprintDTO {
	return v.value
}

func (v *NullableFootprintDTO) Set(val *FootprintDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFootprintDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFootprintDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFootprintDTO(val *FootprintDTO) *NullableFootprintDTO {
	return &NullableFootprintDTO{value: val, isSet: true}
}

func (v NullableFootprintDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFootprintDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


