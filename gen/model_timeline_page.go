
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TimelinePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelinePage{}

// TimelinePage struct for TimelinePage
type TimelinePage struct {
	Id NullableInt64 `json:"id,omitempty"`
	Items []map[string]interface{} `json:"items,omitempty"`
	NextPageValue NullableCursor `json:"next_page_value,omitempty"`
	PinnedItems []map[string]interface{} `json:"pinned_items,omitempty"`
	ShowMoreHotPostsButton NullableBool `json:"show_more_hot_posts_button,omitempty"`
	TotalItemCount NullableInt32 `json:"total_item_count,omitempty"`
	TotalItemLimit NullableInt32 `json:"total_item_limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelinePage TimelinePage

// NewTimelinePage instantiates a new TimelinePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelinePage() *TimelinePage {
	this := TimelinePage{}
	return &this
}

// NewTimelinePageWithDefaults instantiates a new TimelinePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelinePageWithDefaults() *TimelinePage {
	this := TimelinePage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TimelinePage) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *TimelinePage) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TimelinePage) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TimelinePage) UnsetId() {
	o.Id.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetItems() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TimelinePage) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *TimelinePage) SetItems(v []map[string]interface{}) {
	o.Items = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *TimelinePage) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *TimelinePage) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *TimelinePage) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *TimelinePage) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetPinnedItems returns the PinnedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetPinnedItems() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.PinnedItems
}

// GetPinnedItemsOk returns a tuple with the PinnedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetPinnedItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PinnedItems) {
		return nil, false
	}
	return o.PinnedItems, true
}

// HasPinnedItems returns a boolean if a field has been set.
func (o *TimelinePage) HasPinnedItems() bool {
	if o != nil && !IsNil(o.PinnedItems) {
		return true
	}

	return false
}

// SetPinnedItems gets a reference to the given []map[string]interface{} and assigns it to the PinnedItems field.
func (o *TimelinePage) SetPinnedItems(v []map[string]interface{}) {
	o.PinnedItems = v
}

// GetShowMoreHotPostsButton returns the ShowMoreHotPostsButton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetShowMoreHotPostsButton() bool {
	if o == nil || IsNil(o.ShowMoreHotPostsButton.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowMoreHotPostsButton.Get()
}

// GetShowMoreHotPostsButtonOk returns a tuple with the ShowMoreHotPostsButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetShowMoreHotPostsButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowMoreHotPostsButton.Get(), o.ShowMoreHotPostsButton.IsSet()
}

// HasShowMoreHotPostsButton returns a boolean if a field has been set.
func (o *TimelinePage) HasShowMoreHotPostsButton() bool {
	if o != nil && o.ShowMoreHotPostsButton.IsSet() {
		return true
	}

	return false
}

// SetShowMoreHotPostsButton gets a reference to the given NullableBool and assigns it to the ShowMoreHotPostsButton field.
func (o *TimelinePage) SetShowMoreHotPostsButton(v bool) {
	o.ShowMoreHotPostsButton.Set(&v)
}
// SetShowMoreHotPostsButtonNil sets the value for ShowMoreHotPostsButton to be an explicit nil
func (o *TimelinePage) SetShowMoreHotPostsButtonNil() {
	o.ShowMoreHotPostsButton.Set(nil)
}

// UnsetShowMoreHotPostsButton ensures that no value is present for ShowMoreHotPostsButton, not even an explicit nil
func (o *TimelinePage) UnsetShowMoreHotPostsButton() {
	o.ShowMoreHotPostsButton.Unset()
}

// GetTotalItemCount returns the TotalItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetTotalItemCount() int32 {
	if o == nil || IsNil(o.TotalItemCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItemCount.Get()
}

// GetTotalItemCountOk returns a tuple with the TotalItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetTotalItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItemCount.Get(), o.TotalItemCount.IsSet()
}

// HasTotalItemCount returns a boolean if a field has been set.
func (o *TimelinePage) HasTotalItemCount() bool {
	if o != nil && o.TotalItemCount.IsSet() {
		return true
	}

	return false
}

// SetTotalItemCount gets a reference to the given NullableInt32 and assigns it to the TotalItemCount field.
func (o *TimelinePage) SetTotalItemCount(v int32) {
	o.TotalItemCount.Set(&v)
}
// SetTotalItemCountNil sets the value for TotalItemCount to be an explicit nil
func (o *TimelinePage) SetTotalItemCountNil() {
	o.TotalItemCount.Set(nil)
}

// UnsetTotalItemCount ensures that no value is present for TotalItemCount, not even an explicit nil
func (o *TimelinePage) UnsetTotalItemCount() {
	o.TotalItemCount.Unset()
}

// GetTotalItemLimit returns the TotalItemLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePage) GetTotalItemLimit() int32 {
	if o == nil || IsNil(o.TotalItemLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItemLimit.Get()
}

// GetTotalItemLimitOk returns a tuple with the TotalItemLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePage) GetTotalItemLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItemLimit.Get(), o.TotalItemLimit.IsSet()
}

// HasTotalItemLimit returns a boolean if a field has been set.
func (o *TimelinePage) HasTotalItemLimit() bool {
	if o != nil && o.TotalItemLimit.IsSet() {
		return true
	}

	return false
}

// SetTotalItemLimit gets a reference to the given NullableInt32 and assigns it to the TotalItemLimit field.
func (o *TimelinePage) SetTotalItemLimit(v int32) {
	o.TotalItemLimit.Set(&v)
}
// SetTotalItemLimitNil sets the value for TotalItemLimit to be an explicit nil
func (o *TimelinePage) SetTotalItemLimitNil() {
	o.TotalItemLimit.Set(nil)
}

// UnsetTotalItemLimit ensures that no value is present for TotalItemLimit, not even an explicit nil
func (o *TimelinePage) UnsetTotalItemLimit() {
	o.TotalItemLimit.Unset()
}

func (o TimelinePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelinePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.PinnedItems != nil {
		toSerialize["pinned_items"] = o.PinnedItems
	}
	if o.ShowMoreHotPostsButton.IsSet() {
		toSerialize["show_more_hot_posts_button"] = o.ShowMoreHotPostsButton.Get()
	}
	if o.TotalItemCount.IsSet() {
		toSerialize["total_item_count"] = o.TotalItemCount.Get()
	}
	if o.TotalItemLimit.IsSet() {
		toSerialize["total_item_limit"] = o.TotalItemLimit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelinePage) UnmarshalJSON(data []byte) (err error) {
	varTimelinePage := _TimelinePage{}

	err = json.Unmarshal(data, &varTimelinePage)

	if err != nil {
		return err
	}

	*o = TimelinePage(varTimelinePage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "pinned_items")
		delete(additionalProperties, "show_more_hot_posts_button")
		delete(additionalProperties, "total_item_count")
		delete(additionalProperties, "total_item_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelinePage struct {
	value *TimelinePage
	isSet bool
}

func (v NullableTimelinePage) Get() *TimelinePage {
	return v.value
}

func (v *NullableTimelinePage) Set(val *TimelinePage) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelinePage) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelinePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelinePage(val *TimelinePage) *NullableTimelinePage {
	return &NullableTimelinePage{value: val, isSet: true}
}

func (v NullableTimelinePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelinePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


