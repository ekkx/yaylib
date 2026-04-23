
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CallStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallStatusResponse{}

// CallStatusResponse struct for CallStatusResponse
type CallStatusResponse struct {
	PhoneStatus NullableBool `json:"phone_status,omitempty"`
	RoomUrl NullableString `json:"room_url,omitempty"`
	VideoStatus NullableBool `json:"video_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallStatusResponse CallStatusResponse

// NewCallStatusResponse instantiates a new CallStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallStatusResponse() *CallStatusResponse {
	this := CallStatusResponse{}
	return &this
}

// NewCallStatusResponseWithDefaults instantiates a new CallStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallStatusResponseWithDefaults() *CallStatusResponse {
	this := CallStatusResponse{}
	return &this
}

// GetPhoneStatus returns the PhoneStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallStatusResponse) GetPhoneStatus() bool {
	if o == nil || IsNil(o.PhoneStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.PhoneStatus.Get()
}

// GetPhoneStatusOk returns a tuple with the PhoneStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallStatusResponse) GetPhoneStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneStatus.Get(), o.PhoneStatus.IsSet()
}

// HasPhoneStatus returns a boolean if a field has been set.
func (o *CallStatusResponse) HasPhoneStatus() bool {
	if o != nil && o.PhoneStatus.IsSet() {
		return true
	}

	return false
}

// SetPhoneStatus gets a reference to the given NullableBool and assigns it to the PhoneStatus field.
func (o *CallStatusResponse) SetPhoneStatus(v bool) {
	o.PhoneStatus.Set(&v)
}
// SetPhoneStatusNil sets the value for PhoneStatus to be an explicit nil
func (o *CallStatusResponse) SetPhoneStatusNil() {
	o.PhoneStatus.Set(nil)
}

// UnsetPhoneStatus ensures that no value is present for PhoneStatus, not even an explicit nil
func (o *CallStatusResponse) UnsetPhoneStatus() {
	o.PhoneStatus.Unset()
}

// GetRoomUrl returns the RoomUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallStatusResponse) GetRoomUrl() string {
	if o == nil || IsNil(o.RoomUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RoomUrl.Get()
}

// GetRoomUrlOk returns a tuple with the RoomUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallStatusResponse) GetRoomUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomUrl.Get(), o.RoomUrl.IsSet()
}

// HasRoomUrl returns a boolean if a field has been set.
func (o *CallStatusResponse) HasRoomUrl() bool {
	if o != nil && o.RoomUrl.IsSet() {
		return true
	}

	return false
}

// SetRoomUrl gets a reference to the given NullableString and assigns it to the RoomUrl field.
func (o *CallStatusResponse) SetRoomUrl(v string) {
	o.RoomUrl.Set(&v)
}
// SetRoomUrlNil sets the value for RoomUrl to be an explicit nil
func (o *CallStatusResponse) SetRoomUrlNil() {
	o.RoomUrl.Set(nil)
}

// UnsetRoomUrl ensures that no value is present for RoomUrl, not even an explicit nil
func (o *CallStatusResponse) UnsetRoomUrl() {
	o.RoomUrl.Unset()
}

// GetVideoStatus returns the VideoStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallStatusResponse) GetVideoStatus() bool {
	if o == nil || IsNil(o.VideoStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.VideoStatus.Get()
}

// GetVideoStatusOk returns a tuple with the VideoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallStatusResponse) GetVideoStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoStatus.Get(), o.VideoStatus.IsSet()
}

// HasVideoStatus returns a boolean if a field has been set.
func (o *CallStatusResponse) HasVideoStatus() bool {
	if o != nil && o.VideoStatus.IsSet() {
		return true
	}

	return false
}

// SetVideoStatus gets a reference to the given NullableBool and assigns it to the VideoStatus field.
func (o *CallStatusResponse) SetVideoStatus(v bool) {
	o.VideoStatus.Set(&v)
}
// SetVideoStatusNil sets the value for VideoStatus to be an explicit nil
func (o *CallStatusResponse) SetVideoStatusNil() {
	o.VideoStatus.Set(nil)
}

// UnsetVideoStatus ensures that no value is present for VideoStatus, not even an explicit nil
func (o *CallStatusResponse) UnsetVideoStatus() {
	o.VideoStatus.Unset()
}

func (o CallStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneStatus.IsSet() {
		toSerialize["phone_status"] = o.PhoneStatus.Get()
	}
	if o.RoomUrl.IsSet() {
		toSerialize["room_url"] = o.RoomUrl.Get()
	}
	if o.VideoStatus.IsSet() {
		toSerialize["video_status"] = o.VideoStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varCallStatusResponse := _CallStatusResponse{}

	err = json.Unmarshal(data, &varCallStatusResponse)

	if err != nil {
		return err
	}

	*o = CallStatusResponse(varCallStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "phone_status")
		delete(additionalProperties, "room_url")
		delete(additionalProperties, "video_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallStatusResponse struct {
	value *CallStatusResponse
	isSet bool
}

func (v NullableCallStatusResponse) Get() *CallStatusResponse {
	return v.value
}

func (v *NullableCallStatusResponse) Set(val *CallStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallStatusResponse(val *CallStatusResponse) *NullableCallStatusResponse {
	return &NullableCallStatusResponse{value: val, isSet: true}
}

func (v NullableCallStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


