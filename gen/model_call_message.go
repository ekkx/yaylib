
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CallMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallMessage{}

// CallMessage struct for CallMessage
type CallMessage struct {
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	Id NullableString `json:"id,omitempty"`
	IsMessageDeleted NullableBool `json:"is_message_deleted,omitempty"`
	Text NullableString `json:"text,omitempty"`
	UserId NullableString `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallMessage CallMessage

// NewCallMessage instantiates a new CallMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallMessage() *CallMessage {
	this := CallMessage{}
	return &this
}

// NewCallMessageWithDefaults instantiates a new CallMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallMessageWithDefaults() *CallMessage {
	this := CallMessage{}
	return &this
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallMessage) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallMessage) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *CallMessage) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *CallMessage) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *CallMessage) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *CallMessage) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallMessage) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CallMessage) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CallMessage) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CallMessage) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CallMessage) UnsetId() {
	o.Id.Unset()
}

// GetIsMessageDeleted returns the IsMessageDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallMessage) GetIsMessageDeleted() bool {
	if o == nil || IsNil(o.IsMessageDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMessageDeleted.Get()
}

// GetIsMessageDeletedOk returns a tuple with the IsMessageDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallMessage) GetIsMessageDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMessageDeleted.Get(), o.IsMessageDeleted.IsSet()
}

// HasIsMessageDeleted returns a boolean if a field has been set.
func (o *CallMessage) HasIsMessageDeleted() bool {
	if o != nil && o.IsMessageDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsMessageDeleted gets a reference to the given NullableBool and assigns it to the IsMessageDeleted field.
func (o *CallMessage) SetIsMessageDeleted(v bool) {
	o.IsMessageDeleted.Set(&v)
}
// SetIsMessageDeletedNil sets the value for IsMessageDeleted to be an explicit nil
func (o *CallMessage) SetIsMessageDeletedNil() {
	o.IsMessageDeleted.Set(nil)
}

// UnsetIsMessageDeleted ensures that no value is present for IsMessageDeleted, not even an explicit nil
func (o *CallMessage) UnsetIsMessageDeleted() {
	o.IsMessageDeleted.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallMessage) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *CallMessage) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *CallMessage) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *CallMessage) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *CallMessage) UnsetText() {
	o.Text.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallMessage) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallMessage) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CallMessage) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *CallMessage) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CallMessage) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CallMessage) UnsetUserId() {
	o.UserId.Unset()
}

func (o CallMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsMessageDeleted.IsSet() {
		toSerialize["is_message_deleted"] = o.IsMessageDeleted.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallMessage) UnmarshalJSON(data []byte) (err error) {
	varCallMessage := _CallMessage{}

	err = json.Unmarshal(data, &varCallMessage)

	if err != nil {
		return err
	}

	*o = CallMessage(varCallMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_message_deleted")
		delete(additionalProperties, "text")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallMessage struct {
	value *CallMessage
	isSet bool
}

func (v NullableCallMessage) Get() *CallMessage {
	return v.value
}

func (v *NullableCallMessage) Set(val *CallMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCallMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCallMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallMessage(val *CallMessage) *NullableCallMessage {
	return &NullableCallMessage{value: val, isSet: true}
}

func (v NullableCallMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


