
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateGroupThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupThreadRequest{}

// CreateGroupThreadRequest struct for CreateGroupThreadRequest
type CreateGroupThreadRequest struct {
	GroupId NullableInt64 `json:"group_id,omitempty"`
	ThreadIconFilename NullableString `json:"thread_icon_filename,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateGroupThreadRequest CreateGroupThreadRequest

// NewCreateGroupThreadRequest instantiates a new CreateGroupThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupThreadRequest() *CreateGroupThreadRequest {
	this := CreateGroupThreadRequest{}
	return &this
}

// NewCreateGroupThreadRequestWithDefaults instantiates a new CreateGroupThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupThreadRequestWithDefaults() *CreateGroupThreadRequest {
	this := CreateGroupThreadRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupThreadRequest) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupThreadRequest) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *CreateGroupThreadRequest) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *CreateGroupThreadRequest) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *CreateGroupThreadRequest) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *CreateGroupThreadRequest) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetThreadIconFilename returns the ThreadIconFilename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupThreadRequest) GetThreadIconFilename() string {
	if o == nil || IsNil(o.ThreadIconFilename.Get()) {
		var ret string
		return ret
	}
	return *o.ThreadIconFilename.Get()
}

// GetThreadIconFilenameOk returns a tuple with the ThreadIconFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupThreadRequest) GetThreadIconFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadIconFilename.Get(), o.ThreadIconFilename.IsSet()
}

// HasThreadIconFilename returns a boolean if a field has been set.
func (o *CreateGroupThreadRequest) HasThreadIconFilename() bool {
	if o != nil && o.ThreadIconFilename.IsSet() {
		return true
	}

	return false
}

// SetThreadIconFilename gets a reference to the given NullableString and assigns it to the ThreadIconFilename field.
func (o *CreateGroupThreadRequest) SetThreadIconFilename(v string) {
	o.ThreadIconFilename.Set(&v)
}
// SetThreadIconFilenameNil sets the value for ThreadIconFilename to be an explicit nil
func (o *CreateGroupThreadRequest) SetThreadIconFilenameNil() {
	o.ThreadIconFilename.Set(nil)
}

// UnsetThreadIconFilename ensures that no value is present for ThreadIconFilename, not even an explicit nil
func (o *CreateGroupThreadRequest) UnsetThreadIconFilename() {
	o.ThreadIconFilename.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupThreadRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupThreadRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateGroupThreadRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CreateGroupThreadRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CreateGroupThreadRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CreateGroupThreadRequest) UnsetTitle() {
	o.Title.Unset()
}

func (o CreateGroupThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if o.ThreadIconFilename.IsSet() {
		toSerialize["thread_icon_filename"] = o.ThreadIconFilename.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateGroupThreadRequest) UnmarshalJSON(data []byte) (err error) {
	varCreateGroupThreadRequest := _CreateGroupThreadRequest{}

	err = json.Unmarshal(data, &varCreateGroupThreadRequest)

	if err != nil {
		return err
	}

	*o = CreateGroupThreadRequest(varCreateGroupThreadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "thread_icon_filename")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateGroupThreadRequest struct {
	value *CreateGroupThreadRequest
	isSet bool
}

func (v NullableCreateGroupThreadRequest) Get() *CreateGroupThreadRequest {
	return v.value
}

func (v *NullableCreateGroupThreadRequest) Set(val *CreateGroupThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupThreadRequest(val *CreateGroupThreadRequest) *NullableCreateGroupThreadRequest {
	return &NullableCreateGroupThreadRequest{value: val, isSet: true}
}

func (v NullableCreateGroupThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


