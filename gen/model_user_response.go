
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponse{}

// UserResponse struct for UserResponse
type UserResponse struct {
	AppleConnected NullableBool `json:"apple_connected,omitempty"`
	BirthDate NullableString `json:"birth_date,omitempty"`
	BlockingLimit NullableInt32 `json:"blocking_limit,omitempty"`
	EmailConfirmed NullableBool `json:"email_confirmed,omitempty"`
	FacebookConnected NullableBool `json:"facebook_connected,omitempty"`
	GiftingAbility NullableRealmGiftingAbility `json:"gifting_ability,omitempty"`
	GoogleConnected NullableBool `json:"google_connected,omitempty"`
	GroupPhoneOn NullableBool `json:"group_phone_on,omitempty"`
	GroupVideoOn NullableBool `json:"group_video_on,omitempty"`
	InterestsCount NullableInt32 `json:"interests_count,omitempty"`
	LineConnected NullableBool `json:"line_connected,omitempty"`
	MaskedEmail NullableString `json:"masked_email,omitempty"`
	PhoneOn NullableBool `json:"phone_on,omitempty"`
	PushNotification NullableBool `json:"push_notification,omitempty"`
	TwitterId NullableString `json:"twitter_id,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	VideoOn NullableBool `json:"video_on,omitempty"`
	VipUntilSeconds NullableInt64 `json:"vip_until_seconds,omitempty"`
	WorldIdConnected NullableBool `json:"world_id_connected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserResponse UserResponse

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetAppleConnected returns the AppleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetAppleConnected() bool {
	if o == nil || IsNil(o.AppleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.AppleConnected.Get()
}

// GetAppleConnectedOk returns a tuple with the AppleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetAppleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppleConnected.Get(), o.AppleConnected.IsSet()
}

// HasAppleConnected returns a boolean if a field has been set.
func (o *UserResponse) HasAppleConnected() bool {
	if o != nil && o.AppleConnected.IsSet() {
		return true
	}

	return false
}

// SetAppleConnected gets a reference to the given NullableBool and assigns it to the AppleConnected field.
func (o *UserResponse) SetAppleConnected(v bool) {
	o.AppleConnected.Set(&v)
}
// SetAppleConnectedNil sets the value for AppleConnected to be an explicit nil
func (o *UserResponse) SetAppleConnectedNil() {
	o.AppleConnected.Set(nil)
}

// UnsetAppleConnected ensures that no value is present for AppleConnected, not even an explicit nil
func (o *UserResponse) UnsetAppleConnected() {
	o.AppleConnected.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *UserResponse) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *UserResponse) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *UserResponse) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *UserResponse) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBlockingLimit returns the BlockingLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetBlockingLimit() int32 {
	if o == nil || IsNil(o.BlockingLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.BlockingLimit.Get()
}

// GetBlockingLimitOk returns a tuple with the BlockingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetBlockingLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockingLimit.Get(), o.BlockingLimit.IsSet()
}

// HasBlockingLimit returns a boolean if a field has been set.
func (o *UserResponse) HasBlockingLimit() bool {
	if o != nil && o.BlockingLimit.IsSet() {
		return true
	}

	return false
}

// SetBlockingLimit gets a reference to the given NullableInt32 and assigns it to the BlockingLimit field.
func (o *UserResponse) SetBlockingLimit(v int32) {
	o.BlockingLimit.Set(&v)
}
// SetBlockingLimitNil sets the value for BlockingLimit to be an explicit nil
func (o *UserResponse) SetBlockingLimitNil() {
	o.BlockingLimit.Set(nil)
}

// UnsetBlockingLimit ensures that no value is present for BlockingLimit, not even an explicit nil
func (o *UserResponse) UnsetBlockingLimit() {
	o.BlockingLimit.Unset()
}

// GetEmailConfirmed returns the EmailConfirmed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetEmailConfirmed() bool {
	if o == nil || IsNil(o.EmailConfirmed.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailConfirmed.Get()
}

// GetEmailConfirmedOk returns a tuple with the EmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetEmailConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailConfirmed.Get(), o.EmailConfirmed.IsSet()
}

// HasEmailConfirmed returns a boolean if a field has been set.
func (o *UserResponse) HasEmailConfirmed() bool {
	if o != nil && o.EmailConfirmed.IsSet() {
		return true
	}

	return false
}

// SetEmailConfirmed gets a reference to the given NullableBool and assigns it to the EmailConfirmed field.
func (o *UserResponse) SetEmailConfirmed(v bool) {
	o.EmailConfirmed.Set(&v)
}
// SetEmailConfirmedNil sets the value for EmailConfirmed to be an explicit nil
func (o *UserResponse) SetEmailConfirmedNil() {
	o.EmailConfirmed.Set(nil)
}

// UnsetEmailConfirmed ensures that no value is present for EmailConfirmed, not even an explicit nil
func (o *UserResponse) UnsetEmailConfirmed() {
	o.EmailConfirmed.Unset()
}

// GetFacebookConnected returns the FacebookConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetFacebookConnected() bool {
	if o == nil || IsNil(o.FacebookConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.FacebookConnected.Get()
}

// GetFacebookConnectedOk returns a tuple with the FacebookConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetFacebookConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacebookConnected.Get(), o.FacebookConnected.IsSet()
}

// HasFacebookConnected returns a boolean if a field has been set.
func (o *UserResponse) HasFacebookConnected() bool {
	if o != nil && o.FacebookConnected.IsSet() {
		return true
	}

	return false
}

// SetFacebookConnected gets a reference to the given NullableBool and assigns it to the FacebookConnected field.
func (o *UserResponse) SetFacebookConnected(v bool) {
	o.FacebookConnected.Set(&v)
}
// SetFacebookConnectedNil sets the value for FacebookConnected to be an explicit nil
func (o *UserResponse) SetFacebookConnectedNil() {
	o.FacebookConnected.Set(nil)
}

// UnsetFacebookConnected ensures that no value is present for FacebookConnected, not even an explicit nil
func (o *UserResponse) UnsetFacebookConnected() {
	o.FacebookConnected.Unset()
}

// GetGiftingAbility returns the GiftingAbility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetGiftingAbility() RealmGiftingAbility {
	if o == nil || IsNil(o.GiftingAbility.Get()) {
		var ret RealmGiftingAbility
		return ret
	}
	return *o.GiftingAbility.Get()
}

// GetGiftingAbilityOk returns a tuple with the GiftingAbility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetGiftingAbilityOk() (*RealmGiftingAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftingAbility.Get(), o.GiftingAbility.IsSet()
}

// HasGiftingAbility returns a boolean if a field has been set.
func (o *UserResponse) HasGiftingAbility() bool {
	if o != nil && o.GiftingAbility.IsSet() {
		return true
	}

	return false
}

// SetGiftingAbility gets a reference to the given NullableRealmGiftingAbility and assigns it to the GiftingAbility field.
func (o *UserResponse) SetGiftingAbility(v RealmGiftingAbility) {
	o.GiftingAbility.Set(&v)
}
// SetGiftingAbilityNil sets the value for GiftingAbility to be an explicit nil
func (o *UserResponse) SetGiftingAbilityNil() {
	o.GiftingAbility.Set(nil)
}

// UnsetGiftingAbility ensures that no value is present for GiftingAbility, not even an explicit nil
func (o *UserResponse) UnsetGiftingAbility() {
	o.GiftingAbility.Unset()
}

// GetGoogleConnected returns the GoogleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetGoogleConnected() bool {
	if o == nil || IsNil(o.GoogleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.GoogleConnected.Get()
}

// GetGoogleConnectedOk returns a tuple with the GoogleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetGoogleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoogleConnected.Get(), o.GoogleConnected.IsSet()
}

// HasGoogleConnected returns a boolean if a field has been set.
func (o *UserResponse) HasGoogleConnected() bool {
	if o != nil && o.GoogleConnected.IsSet() {
		return true
	}

	return false
}

// SetGoogleConnected gets a reference to the given NullableBool and assigns it to the GoogleConnected field.
func (o *UserResponse) SetGoogleConnected(v bool) {
	o.GoogleConnected.Set(&v)
}
// SetGoogleConnectedNil sets the value for GoogleConnected to be an explicit nil
func (o *UserResponse) SetGoogleConnectedNil() {
	o.GoogleConnected.Set(nil)
}

// UnsetGoogleConnected ensures that no value is present for GoogleConnected, not even an explicit nil
func (o *UserResponse) UnsetGoogleConnected() {
	o.GoogleConnected.Unset()
}

// GetGroupPhoneOn returns the GroupPhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetGroupPhoneOn() bool {
	if o == nil || IsNil(o.GroupPhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.GroupPhoneOn.Get()
}

// GetGroupPhoneOnOk returns a tuple with the GroupPhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetGroupPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupPhoneOn.Get(), o.GroupPhoneOn.IsSet()
}

// HasGroupPhoneOn returns a boolean if a field has been set.
func (o *UserResponse) HasGroupPhoneOn() bool {
	if o != nil && o.GroupPhoneOn.IsSet() {
		return true
	}

	return false
}

// SetGroupPhoneOn gets a reference to the given NullableBool and assigns it to the GroupPhoneOn field.
func (o *UserResponse) SetGroupPhoneOn(v bool) {
	o.GroupPhoneOn.Set(&v)
}
// SetGroupPhoneOnNil sets the value for GroupPhoneOn to be an explicit nil
func (o *UserResponse) SetGroupPhoneOnNil() {
	o.GroupPhoneOn.Set(nil)
}

// UnsetGroupPhoneOn ensures that no value is present for GroupPhoneOn, not even an explicit nil
func (o *UserResponse) UnsetGroupPhoneOn() {
	o.GroupPhoneOn.Unset()
}

// GetGroupVideoOn returns the GroupVideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetGroupVideoOn() bool {
	if o == nil || IsNil(o.GroupVideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.GroupVideoOn.Get()
}

// GetGroupVideoOnOk returns a tuple with the GroupVideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetGroupVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupVideoOn.Get(), o.GroupVideoOn.IsSet()
}

// HasGroupVideoOn returns a boolean if a field has been set.
func (o *UserResponse) HasGroupVideoOn() bool {
	if o != nil && o.GroupVideoOn.IsSet() {
		return true
	}

	return false
}

// SetGroupVideoOn gets a reference to the given NullableBool and assigns it to the GroupVideoOn field.
func (o *UserResponse) SetGroupVideoOn(v bool) {
	o.GroupVideoOn.Set(&v)
}
// SetGroupVideoOnNil sets the value for GroupVideoOn to be an explicit nil
func (o *UserResponse) SetGroupVideoOnNil() {
	o.GroupVideoOn.Set(nil)
}

// UnsetGroupVideoOn ensures that no value is present for GroupVideoOn, not even an explicit nil
func (o *UserResponse) UnsetGroupVideoOn() {
	o.GroupVideoOn.Unset()
}

// GetInterestsCount returns the InterestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetInterestsCount() int32 {
	if o == nil || IsNil(o.InterestsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InterestsCount.Get()
}

// GetInterestsCountOk returns a tuple with the InterestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetInterestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterestsCount.Get(), o.InterestsCount.IsSet()
}

// HasInterestsCount returns a boolean if a field has been set.
func (o *UserResponse) HasInterestsCount() bool {
	if o != nil && o.InterestsCount.IsSet() {
		return true
	}

	return false
}

// SetInterestsCount gets a reference to the given NullableInt32 and assigns it to the InterestsCount field.
func (o *UserResponse) SetInterestsCount(v int32) {
	o.InterestsCount.Set(&v)
}
// SetInterestsCountNil sets the value for InterestsCount to be an explicit nil
func (o *UserResponse) SetInterestsCountNil() {
	o.InterestsCount.Set(nil)
}

// UnsetInterestsCount ensures that no value is present for InterestsCount, not even an explicit nil
func (o *UserResponse) UnsetInterestsCount() {
	o.InterestsCount.Unset()
}

// GetLineConnected returns the LineConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetLineConnected() bool {
	if o == nil || IsNil(o.LineConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.LineConnected.Get()
}

// GetLineConnectedOk returns a tuple with the LineConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetLineConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineConnected.Get(), o.LineConnected.IsSet()
}

// HasLineConnected returns a boolean if a field has been set.
func (o *UserResponse) HasLineConnected() bool {
	if o != nil && o.LineConnected.IsSet() {
		return true
	}

	return false
}

// SetLineConnected gets a reference to the given NullableBool and assigns it to the LineConnected field.
func (o *UserResponse) SetLineConnected(v bool) {
	o.LineConnected.Set(&v)
}
// SetLineConnectedNil sets the value for LineConnected to be an explicit nil
func (o *UserResponse) SetLineConnectedNil() {
	o.LineConnected.Set(nil)
}

// UnsetLineConnected ensures that no value is present for LineConnected, not even an explicit nil
func (o *UserResponse) UnsetLineConnected() {
	o.LineConnected.Unset()
}

// GetMaskedEmail returns the MaskedEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetMaskedEmail() string {
	if o == nil || IsNil(o.MaskedEmail.Get()) {
		var ret string
		return ret
	}
	return *o.MaskedEmail.Get()
}

// GetMaskedEmailOk returns a tuple with the MaskedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetMaskedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedEmail.Get(), o.MaskedEmail.IsSet()
}

// HasMaskedEmail returns a boolean if a field has been set.
func (o *UserResponse) HasMaskedEmail() bool {
	if o != nil && o.MaskedEmail.IsSet() {
		return true
	}

	return false
}

// SetMaskedEmail gets a reference to the given NullableString and assigns it to the MaskedEmail field.
func (o *UserResponse) SetMaskedEmail(v string) {
	o.MaskedEmail.Set(&v)
}
// SetMaskedEmailNil sets the value for MaskedEmail to be an explicit nil
func (o *UserResponse) SetMaskedEmailNil() {
	o.MaskedEmail.Set(nil)
}

// UnsetMaskedEmail ensures that no value is present for MaskedEmail, not even an explicit nil
func (o *UserResponse) UnsetMaskedEmail() {
	o.MaskedEmail.Unset()
}

// GetPhoneOn returns the PhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetPhoneOn() bool {
	if o == nil || IsNil(o.PhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.PhoneOn.Get()
}

// GetPhoneOnOk returns a tuple with the PhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneOn.Get(), o.PhoneOn.IsSet()
}

// HasPhoneOn returns a boolean if a field has been set.
func (o *UserResponse) HasPhoneOn() bool {
	if o != nil && o.PhoneOn.IsSet() {
		return true
	}

	return false
}

// SetPhoneOn gets a reference to the given NullableBool and assigns it to the PhoneOn field.
func (o *UserResponse) SetPhoneOn(v bool) {
	o.PhoneOn.Set(&v)
}
// SetPhoneOnNil sets the value for PhoneOn to be an explicit nil
func (o *UserResponse) SetPhoneOnNil() {
	o.PhoneOn.Set(nil)
}

// UnsetPhoneOn ensures that no value is present for PhoneOn, not even an explicit nil
func (o *UserResponse) UnsetPhoneOn() {
	o.PhoneOn.Unset()
}

// GetPushNotification returns the PushNotification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetPushNotification() bool {
	if o == nil || IsNil(o.PushNotification.Get()) {
		var ret bool
		return ret
	}
	return *o.PushNotification.Get()
}

// GetPushNotificationOk returns a tuple with the PushNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetPushNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushNotification.Get(), o.PushNotification.IsSet()
}

// HasPushNotification returns a boolean if a field has been set.
func (o *UserResponse) HasPushNotification() bool {
	if o != nil && o.PushNotification.IsSet() {
		return true
	}

	return false
}

// SetPushNotification gets a reference to the given NullableBool and assigns it to the PushNotification field.
func (o *UserResponse) SetPushNotification(v bool) {
	o.PushNotification.Set(&v)
}
// SetPushNotificationNil sets the value for PushNotification to be an explicit nil
func (o *UserResponse) SetPushNotificationNil() {
	o.PushNotification.Set(nil)
}

// UnsetPushNotification ensures that no value is present for PushNotification, not even an explicit nil
func (o *UserResponse) UnsetPushNotification() {
	o.PushNotification.Unset()
}

// GetTwitterId returns the TwitterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetTwitterId() string {
	if o == nil || IsNil(o.TwitterId.Get()) {
		var ret string
		return ret
	}
	return *o.TwitterId.Get()
}

// GetTwitterIdOk returns a tuple with the TwitterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetTwitterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterId.Get(), o.TwitterId.IsSet()
}

// HasTwitterId returns a boolean if a field has been set.
func (o *UserResponse) HasTwitterId() bool {
	if o != nil && o.TwitterId.IsSet() {
		return true
	}

	return false
}

// SetTwitterId gets a reference to the given NullableString and assigns it to the TwitterId field.
func (o *UserResponse) SetTwitterId(v string) {
	o.TwitterId.Set(&v)
}
// SetTwitterIdNil sets the value for TwitterId to be an explicit nil
func (o *UserResponse) SetTwitterIdNil() {
	o.TwitterId.Set(nil)
}

// UnsetTwitterId ensures that no value is present for TwitterId, not even an explicit nil
func (o *UserResponse) UnsetTwitterId() {
	o.TwitterId.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *UserResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *UserResponse) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *UserResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *UserResponse) UnsetUser() {
	o.User.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *UserResponse) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *UserResponse) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *UserResponse) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *UserResponse) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVideoOn returns the VideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetVideoOn() bool {
	if o == nil || IsNil(o.VideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.VideoOn.Get()
}

// GetVideoOnOk returns a tuple with the VideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoOn.Get(), o.VideoOn.IsSet()
}

// HasVideoOn returns a boolean if a field has been set.
func (o *UserResponse) HasVideoOn() bool {
	if o != nil && o.VideoOn.IsSet() {
		return true
	}

	return false
}

// SetVideoOn gets a reference to the given NullableBool and assigns it to the VideoOn field.
func (o *UserResponse) SetVideoOn(v bool) {
	o.VideoOn.Set(&v)
}
// SetVideoOnNil sets the value for VideoOn to be an explicit nil
func (o *UserResponse) SetVideoOnNil() {
	o.VideoOn.Set(nil)
}

// UnsetVideoOn ensures that no value is present for VideoOn, not even an explicit nil
func (o *UserResponse) UnsetVideoOn() {
	o.VideoOn.Unset()
}

// GetVipUntilSeconds returns the VipUntilSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetVipUntilSeconds() int64 {
	if o == nil || IsNil(o.VipUntilSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.VipUntilSeconds.Get()
}

// GetVipUntilSecondsOk returns a tuple with the VipUntilSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetVipUntilSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipUntilSeconds.Get(), o.VipUntilSeconds.IsSet()
}

// HasVipUntilSeconds returns a boolean if a field has been set.
func (o *UserResponse) HasVipUntilSeconds() bool {
	if o != nil && o.VipUntilSeconds.IsSet() {
		return true
	}

	return false
}

// SetVipUntilSeconds gets a reference to the given NullableInt64 and assigns it to the VipUntilSeconds field.
func (o *UserResponse) SetVipUntilSeconds(v int64) {
	o.VipUntilSeconds.Set(&v)
}
// SetVipUntilSecondsNil sets the value for VipUntilSeconds to be an explicit nil
func (o *UserResponse) SetVipUntilSecondsNil() {
	o.VipUntilSeconds.Set(nil)
}

// UnsetVipUntilSeconds ensures that no value is present for VipUntilSeconds, not even an explicit nil
func (o *UserResponse) UnsetVipUntilSeconds() {
	o.VipUntilSeconds.Unset()
}

// GetWorldIdConnected returns the WorldIdConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetWorldIdConnected() bool {
	if o == nil || IsNil(o.WorldIdConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.WorldIdConnected.Get()
}

// GetWorldIdConnectedOk returns a tuple with the WorldIdConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetWorldIdConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorldIdConnected.Get(), o.WorldIdConnected.IsSet()
}

// HasWorldIdConnected returns a boolean if a field has been set.
func (o *UserResponse) HasWorldIdConnected() bool {
	if o != nil && o.WorldIdConnected.IsSet() {
		return true
	}

	return false
}

// SetWorldIdConnected gets a reference to the given NullableBool and assigns it to the WorldIdConnected field.
func (o *UserResponse) SetWorldIdConnected(v bool) {
	o.WorldIdConnected.Set(&v)
}
// SetWorldIdConnectedNil sets the value for WorldIdConnected to be an explicit nil
func (o *UserResponse) SetWorldIdConnectedNil() {
	o.WorldIdConnected.Set(nil)
}

// UnsetWorldIdConnected ensures that no value is present for WorldIdConnected, not even an explicit nil
func (o *UserResponse) UnsetWorldIdConnected() {
	o.WorldIdConnected.Unset()
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppleConnected.IsSet() {
		toSerialize["apple_connected"] = o.AppleConnected.Get()
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if o.BlockingLimit.IsSet() {
		toSerialize["blocking_limit"] = o.BlockingLimit.Get()
	}
	if o.EmailConfirmed.IsSet() {
		toSerialize["email_confirmed"] = o.EmailConfirmed.Get()
	}
	if o.FacebookConnected.IsSet() {
		toSerialize["facebook_connected"] = o.FacebookConnected.Get()
	}
	if o.GiftingAbility.IsSet() {
		toSerialize["gifting_ability"] = o.GiftingAbility.Get()
	}
	if o.GoogleConnected.IsSet() {
		toSerialize["google_connected"] = o.GoogleConnected.Get()
	}
	if o.GroupPhoneOn.IsSet() {
		toSerialize["group_phone_on"] = o.GroupPhoneOn.Get()
	}
	if o.GroupVideoOn.IsSet() {
		toSerialize["group_video_on"] = o.GroupVideoOn.Get()
	}
	if o.InterestsCount.IsSet() {
		toSerialize["interests_count"] = o.InterestsCount.Get()
	}
	if o.LineConnected.IsSet() {
		toSerialize["line_connected"] = o.LineConnected.Get()
	}
	if o.MaskedEmail.IsSet() {
		toSerialize["masked_email"] = o.MaskedEmail.Get()
	}
	if o.PhoneOn.IsSet() {
		toSerialize["phone_on"] = o.PhoneOn.Get()
	}
	if o.PushNotification.IsSet() {
		toSerialize["push_notification"] = o.PushNotification.Get()
	}
	if o.TwitterId.IsSet() {
		toSerialize["twitter_id"] = o.TwitterId.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.VideoOn.IsSet() {
		toSerialize["video_on"] = o.VideoOn.Get()
	}
	if o.VipUntilSeconds.IsSet() {
		toSerialize["vip_until_seconds"] = o.VipUntilSeconds.Get()
	}
	if o.WorldIdConnected.IsSet() {
		toSerialize["world_id_connected"] = o.WorldIdConnected.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserResponse) UnmarshalJSON(data []byte) (err error) {
	varUserResponse := _UserResponse{}

	err = json.Unmarshal(data, &varUserResponse)

	if err != nil {
		return err
	}

	*o = UserResponse(varUserResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apple_connected")
		delete(additionalProperties, "birth_date")
		delete(additionalProperties, "blocking_limit")
		delete(additionalProperties, "email_confirmed")
		delete(additionalProperties, "facebook_connected")
		delete(additionalProperties, "gifting_ability")
		delete(additionalProperties, "google_connected")
		delete(additionalProperties, "group_phone_on")
		delete(additionalProperties, "group_video_on")
		delete(additionalProperties, "interests_count")
		delete(additionalProperties, "line_connected")
		delete(additionalProperties, "masked_email")
		delete(additionalProperties, "phone_on")
		delete(additionalProperties, "push_notification")
		delete(additionalProperties, "twitter_id")
		delete(additionalProperties, "user")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "video_on")
		delete(additionalProperties, "vip_until_seconds")
		delete(additionalProperties, "world_id_connected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


