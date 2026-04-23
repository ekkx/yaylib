
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RecentSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentSearch{}

// RecentSearch struct for RecentSearch
type RecentSearch struct {
	Hashtag NullablePostTag `json:"hashtag,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Keyword NullableString `json:"keyword,omitempty"`
	Type NullableString `json:"type,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentSearch RecentSearch

// NewRecentSearch instantiates a new RecentSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentSearch() *RecentSearch {
	this := RecentSearch{}
	return &this
}

// NewRecentSearchWithDefaults instantiates a new RecentSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentSearchWithDefaults() *RecentSearch {
	this := RecentSearch{}
	return &this
}

// GetHashtag returns the Hashtag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearch) GetHashtag() PostTag {
	if o == nil || IsNil(o.Hashtag.Get()) {
		var ret PostTag
		return ret
	}
	return *o.Hashtag.Get()
}

// GetHashtagOk returns a tuple with the Hashtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearch) GetHashtagOk() (*PostTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hashtag.Get(), o.Hashtag.IsSet()
}

// HasHashtag returns a boolean if a field has been set.
func (o *RecentSearch) HasHashtag() bool {
	if o != nil && o.Hashtag.IsSet() {
		return true
	}

	return false
}

// SetHashtag gets a reference to the given NullablePostTag and assigns it to the Hashtag field.
func (o *RecentSearch) SetHashtag(v PostTag) {
	o.Hashtag.Set(&v)
}
// SetHashtagNil sets the value for Hashtag to be an explicit nil
func (o *RecentSearch) SetHashtagNil() {
	o.Hashtag.Set(nil)
}

// UnsetHashtag ensures that no value is present for Hashtag, not even an explicit nil
func (o *RecentSearch) UnsetHashtag() {
	o.Hashtag.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearch) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearch) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RecentSearch) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RecentSearch) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RecentSearch) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RecentSearch) UnsetId() {
	o.Id.Unset()
}

// GetKeyword returns the Keyword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearch) GetKeyword() string {
	if o == nil || IsNil(o.Keyword.Get()) {
		var ret string
		return ret
	}
	return *o.Keyword.Get()
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearch) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keyword.Get(), o.Keyword.IsSet()
}

// HasKeyword returns a boolean if a field has been set.
func (o *RecentSearch) HasKeyword() bool {
	if o != nil && o.Keyword.IsSet() {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given NullableString and assigns it to the Keyword field.
func (o *RecentSearch) SetKeyword(v string) {
	o.Keyword.Set(&v)
}
// SetKeywordNil sets the value for Keyword to be an explicit nil
func (o *RecentSearch) SetKeywordNil() {
	o.Keyword.Set(nil)
}

// UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
func (o *RecentSearch) UnsetKeyword() {
	o.Keyword.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearch) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearch) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *RecentSearch) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *RecentSearch) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *RecentSearch) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *RecentSearch) UnsetType() {
	o.Type.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearch) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearch) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *RecentSearch) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *RecentSearch) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *RecentSearch) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *RecentSearch) UnsetUser() {
	o.User.Unset()
}

func (o RecentSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Hashtag.IsSet() {
		toSerialize["hashtag"] = o.Hashtag.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Keyword.IsSet() {
		toSerialize["keyword"] = o.Keyword.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentSearch) UnmarshalJSON(data []byte) (err error) {
	varRecentSearch := _RecentSearch{}

	err = json.Unmarshal(data, &varRecentSearch)

	if err != nil {
		return err
	}

	*o = RecentSearch(varRecentSearch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hashtag")
		delete(additionalProperties, "id")
		delete(additionalProperties, "keyword")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentSearch struct {
	value *RecentSearch
	isSet bool
}

func (v NullableRecentSearch) Get() *RecentSearch {
	return v.value
}

func (v *NullableRecentSearch) Set(val *RecentSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentSearch(val *RecentSearch) *NullableRecentSearch {
	return &NullableRecentSearch{value: val, isSet: true}
}

func (v NullableRecentSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


