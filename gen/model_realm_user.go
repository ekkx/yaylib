
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmUser{}

// RealmUser struct for RealmUser
type RealmUser struct {
	AgeVerified NullableBool `json:"age_verified,omitempty"`
	AppleConnected NullableBool `json:"apple_connected,omitempty"`
	Biography NullableString `json:"biography,omitempty"`
	BirthDate NullableString `json:"birth_date,omitempty"`
	BlockingLimit NullableInt32 `json:"blocking_limit,omitempty"`
	ChatRequest NullableBool `json:"chat_request,omitempty"`
	ConnectedBy []string `json:"connected_by,omitempty"`
	ContactPhones []string `json:"contact_phones,omitempty"`
	CountryCode NullableString `json:"country_code,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	CoverImageThumbnail NullableString `json:"cover_image_thumbnail,omitempty"`
	CreatedAtSeconds NullableInt64 `json:"created_at_seconds,omitempty"`
	DangerousUser NullableBool `json:"dangerous_user,omitempty"`
	EmailConfirmed NullableBool `json:"email_confirmed,omitempty"`
	FacebookConnected NullableBool `json:"facebook_connected,omitempty"`
	FollowPending NullableBool `json:"follow_pending,omitempty"`
	FollowedBy NullableBool `json:"followed_by,omitempty"`
	FollowedByPending NullableBool `json:"followed_by_pending,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	Following NullableBool `json:"following,omitempty"`
	FollowingsCount NullableInt32 `json:"followings_count,omitempty"`
	Frame NullableString `json:"frame,omitempty"`
	FrameThumbnail NullableString `json:"frame_thumbnail,omitempty"`
	FromDifferentGenerationAndTrusted NullableBool `json:"from_different_generation_and_trusted,omitempty"`
	Gender NullableInt32 `json:"gender,omitempty"`
	Generation NullableInt32 `json:"generation,omitempty"`
	GiftingAbility NullableRealmGiftingAbility `json:"gifting_ability,omitempty"`
	GoogleConnected NullableBool `json:"google_connected,omitempty"`
	GroupPhoneOn NullableBool `json:"group_phone_on,omitempty"`
	GroupUser NullableGroupUser `json:"group_user,omitempty"`
	GroupVideoOn NullableBool `json:"group_video_on,omitempty"`
	GroupsUsersCount NullableInt32 `json:"groups_users_count,omitempty"`
	Hidden NullableBool `json:"hidden,omitempty"`
	HideVip NullableBool `json:"hide_vip,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InterestsCount NullableInt32 `json:"interests_count,omitempty"`
	InterestsSelected NullableBool `json:"interests_selected,omitempty"`
	IsPrivate NullableBool `json:"is_private,omitempty"`
	LastLoggedInAtSeconds NullableInt64 `json:"last_logged_in_at_seconds,omitempty"`
	LineConnected NullableBool `json:"line_connected,omitempty"`
	LoginStreakCount NullableInt32 `json:"login_streak_count,omitempty"`
	MaskedEmail NullableString `json:"masked_email,omitempty"`
	MediaCount NullableInt32 `json:"media_count,omitempty"`
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	NewUser NullableBool `json:"new_user,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	OnlineStatus NullableOnlineStatusEnum `json:"online_status,omitempty"`
	PhoneOn NullableBool `json:"phone_on,omitempty"`
	PostsCount NullableInt32 `json:"posts_count,omitempty"`
	Prefecture NullableString `json:"prefecture,omitempty"`
	ProfileIcon NullableString `json:"profile_icon,omitempty"`
	ProfileIconThumbnail NullableString `json:"profile_icon_thumbnail,omitempty"`
	PushNotification NullableBool `json:"push_notification,omitempty"`
	RecentlyKenta NullableBool `json:"recently_kenta,omitempty"`
	RestrictedReviewBy NullableString `json:"restricted_review_by,omitempty"`
	ReviewsCount NullableInt32 `json:"reviews_count,omitempty"`
	Title NullableString `json:"title,omitempty"`
	TotalGiftsReceivedCount NullableInt32 `json:"total_gifts_received_count,omitempty"`
	TwitterId NullableString `json:"twitter_id,omitempty"`
	TwoFaEnabled NullableBool `json:"two_fa_enabled,omitempty"`
	UpdatedTimeMillis NullableInt64 `json:"updated_time_millis,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	VideoOn NullableBool `json:"video_on,omitempty"`
	Vip NullableBool `json:"vip,omitempty"`
	VipUntilSeconds NullableInt64 `json:"vip_until_seconds,omitempty"`
	WorldIdConnected NullableBool `json:"world_id_connected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmUser RealmUser

// NewRealmUser instantiates a new RealmUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmUser() *RealmUser {
	this := RealmUser{}
	return &this
}

// NewRealmUserWithDefaults instantiates a new RealmUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmUserWithDefaults() *RealmUser {
	this := RealmUser{}
	return &this
}

// GetAgeVerified returns the AgeVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetAgeVerified() bool {
	if o == nil || IsNil(o.AgeVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.AgeVerified.Get()
}

// GetAgeVerifiedOk returns a tuple with the AgeVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetAgeVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgeVerified.Get(), o.AgeVerified.IsSet()
}

// HasAgeVerified returns a boolean if a field has been set.
func (o *RealmUser) HasAgeVerified() bool {
	if o != nil && o.AgeVerified.IsSet() {
		return true
	}

	return false
}

// SetAgeVerified gets a reference to the given NullableBool and assigns it to the AgeVerified field.
func (o *RealmUser) SetAgeVerified(v bool) {
	o.AgeVerified.Set(&v)
}
// SetAgeVerifiedNil sets the value for AgeVerified to be an explicit nil
func (o *RealmUser) SetAgeVerifiedNil() {
	o.AgeVerified.Set(nil)
}

// UnsetAgeVerified ensures that no value is present for AgeVerified, not even an explicit nil
func (o *RealmUser) UnsetAgeVerified() {
	o.AgeVerified.Unset()
}

// GetAppleConnected returns the AppleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetAppleConnected() bool {
	if o == nil || IsNil(o.AppleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.AppleConnected.Get()
}

// GetAppleConnectedOk returns a tuple with the AppleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetAppleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppleConnected.Get(), o.AppleConnected.IsSet()
}

// HasAppleConnected returns a boolean if a field has been set.
func (o *RealmUser) HasAppleConnected() bool {
	if o != nil && o.AppleConnected.IsSet() {
		return true
	}

	return false
}

// SetAppleConnected gets a reference to the given NullableBool and assigns it to the AppleConnected field.
func (o *RealmUser) SetAppleConnected(v bool) {
	o.AppleConnected.Set(&v)
}
// SetAppleConnectedNil sets the value for AppleConnected to be an explicit nil
func (o *RealmUser) SetAppleConnectedNil() {
	o.AppleConnected.Set(nil)
}

// UnsetAppleConnected ensures that no value is present for AppleConnected, not even an explicit nil
func (o *RealmUser) UnsetAppleConnected() {
	o.AppleConnected.Unset()
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *RealmUser) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *RealmUser) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *RealmUser) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *RealmUser) UnsetBiography() {
	o.Biography.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *RealmUser) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *RealmUser) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *RealmUser) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *RealmUser) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBlockingLimit returns the BlockingLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetBlockingLimit() int32 {
	if o == nil || IsNil(o.BlockingLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.BlockingLimit.Get()
}

// GetBlockingLimitOk returns a tuple with the BlockingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetBlockingLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockingLimit.Get(), o.BlockingLimit.IsSet()
}

// HasBlockingLimit returns a boolean if a field has been set.
func (o *RealmUser) HasBlockingLimit() bool {
	if o != nil && o.BlockingLimit.IsSet() {
		return true
	}

	return false
}

// SetBlockingLimit gets a reference to the given NullableInt32 and assigns it to the BlockingLimit field.
func (o *RealmUser) SetBlockingLimit(v int32) {
	o.BlockingLimit.Set(&v)
}
// SetBlockingLimitNil sets the value for BlockingLimit to be an explicit nil
func (o *RealmUser) SetBlockingLimitNil() {
	o.BlockingLimit.Set(nil)
}

// UnsetBlockingLimit ensures that no value is present for BlockingLimit, not even an explicit nil
func (o *RealmUser) UnsetBlockingLimit() {
	o.BlockingLimit.Unset()
}

// GetChatRequest returns the ChatRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetChatRequest() bool {
	if o == nil || IsNil(o.ChatRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.ChatRequest.Get()
}

// GetChatRequestOk returns a tuple with the ChatRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetChatRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChatRequest.Get(), o.ChatRequest.IsSet()
}

// HasChatRequest returns a boolean if a field has been set.
func (o *RealmUser) HasChatRequest() bool {
	if o != nil && o.ChatRequest.IsSet() {
		return true
	}

	return false
}

// SetChatRequest gets a reference to the given NullableBool and assigns it to the ChatRequest field.
func (o *RealmUser) SetChatRequest(v bool) {
	o.ChatRequest.Set(&v)
}
// SetChatRequestNil sets the value for ChatRequest to be an explicit nil
func (o *RealmUser) SetChatRequestNil() {
	o.ChatRequest.Set(nil)
}

// UnsetChatRequest ensures that no value is present for ChatRequest, not even an explicit nil
func (o *RealmUser) UnsetChatRequest() {
	o.ChatRequest.Unset()
}

// GetConnectedBy returns the ConnectedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetConnectedBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConnectedBy
}

// GetConnectedByOk returns a tuple with the ConnectedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetConnectedByOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectedBy) {
		return nil, false
	}
	return o.ConnectedBy, true
}

// HasConnectedBy returns a boolean if a field has been set.
func (o *RealmUser) HasConnectedBy() bool {
	if o != nil && !IsNil(o.ConnectedBy) {
		return true
	}

	return false
}

// SetConnectedBy gets a reference to the given []string and assigns it to the ConnectedBy field.
func (o *RealmUser) SetConnectedBy(v []string) {
	o.ConnectedBy = v
}

// GetContactPhones returns the ContactPhones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetContactPhones() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ContactPhones
}

// GetContactPhonesOk returns a tuple with the ContactPhones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetContactPhonesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactPhones) {
		return nil, false
	}
	return o.ContactPhones, true
}

// HasContactPhones returns a boolean if a field has been set.
func (o *RealmUser) HasContactPhones() bool {
	if o != nil && !IsNil(o.ContactPhones) {
		return true
	}

	return false
}

// SetContactPhones gets a reference to the given []string and assigns it to the ContactPhones field.
func (o *RealmUser) SetContactPhones(v []string) {
	o.ContactPhones = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RealmUser) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *RealmUser) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *RealmUser) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *RealmUser) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *RealmUser) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *RealmUser) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *RealmUser) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *RealmUser) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetCoverImageThumbnail returns the CoverImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetCoverImageThumbnail() string {
	if o == nil || IsNil(o.CoverImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImageThumbnail.Get()
}

// GetCoverImageThumbnailOk returns a tuple with the CoverImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetCoverImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImageThumbnail.Get(), o.CoverImageThumbnail.IsSet()
}

// HasCoverImageThumbnail returns a boolean if a field has been set.
func (o *RealmUser) HasCoverImageThumbnail() bool {
	if o != nil && o.CoverImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetCoverImageThumbnail gets a reference to the given NullableString and assigns it to the CoverImageThumbnail field.
func (o *RealmUser) SetCoverImageThumbnail(v string) {
	o.CoverImageThumbnail.Set(&v)
}
// SetCoverImageThumbnailNil sets the value for CoverImageThumbnail to be an explicit nil
func (o *RealmUser) SetCoverImageThumbnailNil() {
	o.CoverImageThumbnail.Set(nil)
}

// UnsetCoverImageThumbnail ensures that no value is present for CoverImageThumbnail, not even an explicit nil
func (o *RealmUser) UnsetCoverImageThumbnail() {
	o.CoverImageThumbnail.Unset()
}

// GetCreatedAtSeconds returns the CreatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetCreatedAtSeconds() int64 {
	if o == nil || IsNil(o.CreatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtSeconds.Get()
}

// GetCreatedAtSecondsOk returns a tuple with the CreatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetCreatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtSeconds.Get(), o.CreatedAtSeconds.IsSet()
}

// HasCreatedAtSeconds returns a boolean if a field has been set.
func (o *RealmUser) HasCreatedAtSeconds() bool {
	if o != nil && o.CreatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the CreatedAtSeconds field.
func (o *RealmUser) SetCreatedAtSeconds(v int64) {
	o.CreatedAtSeconds.Set(&v)
}
// SetCreatedAtSecondsNil sets the value for CreatedAtSeconds to be an explicit nil
func (o *RealmUser) SetCreatedAtSecondsNil() {
	o.CreatedAtSeconds.Set(nil)
}

// UnsetCreatedAtSeconds ensures that no value is present for CreatedAtSeconds, not even an explicit nil
func (o *RealmUser) UnsetCreatedAtSeconds() {
	o.CreatedAtSeconds.Unset()
}

// GetDangerousUser returns the DangerousUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetDangerousUser() bool {
	if o == nil || IsNil(o.DangerousUser.Get()) {
		var ret bool
		return ret
	}
	return *o.DangerousUser.Get()
}

// GetDangerousUserOk returns a tuple with the DangerousUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetDangerousUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DangerousUser.Get(), o.DangerousUser.IsSet()
}

// HasDangerousUser returns a boolean if a field has been set.
func (o *RealmUser) HasDangerousUser() bool {
	if o != nil && o.DangerousUser.IsSet() {
		return true
	}

	return false
}

// SetDangerousUser gets a reference to the given NullableBool and assigns it to the DangerousUser field.
func (o *RealmUser) SetDangerousUser(v bool) {
	o.DangerousUser.Set(&v)
}
// SetDangerousUserNil sets the value for DangerousUser to be an explicit nil
func (o *RealmUser) SetDangerousUserNil() {
	o.DangerousUser.Set(nil)
}

// UnsetDangerousUser ensures that no value is present for DangerousUser, not even an explicit nil
func (o *RealmUser) UnsetDangerousUser() {
	o.DangerousUser.Unset()
}

// GetEmailConfirmed returns the EmailConfirmed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetEmailConfirmed() bool {
	if o == nil || IsNil(o.EmailConfirmed.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailConfirmed.Get()
}

// GetEmailConfirmedOk returns a tuple with the EmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetEmailConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailConfirmed.Get(), o.EmailConfirmed.IsSet()
}

// HasEmailConfirmed returns a boolean if a field has been set.
func (o *RealmUser) HasEmailConfirmed() bool {
	if o != nil && o.EmailConfirmed.IsSet() {
		return true
	}

	return false
}

// SetEmailConfirmed gets a reference to the given NullableBool and assigns it to the EmailConfirmed field.
func (o *RealmUser) SetEmailConfirmed(v bool) {
	o.EmailConfirmed.Set(&v)
}
// SetEmailConfirmedNil sets the value for EmailConfirmed to be an explicit nil
func (o *RealmUser) SetEmailConfirmedNil() {
	o.EmailConfirmed.Set(nil)
}

// UnsetEmailConfirmed ensures that no value is present for EmailConfirmed, not even an explicit nil
func (o *RealmUser) UnsetEmailConfirmed() {
	o.EmailConfirmed.Unset()
}

// GetFacebookConnected returns the FacebookConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFacebookConnected() bool {
	if o == nil || IsNil(o.FacebookConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.FacebookConnected.Get()
}

// GetFacebookConnectedOk returns a tuple with the FacebookConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFacebookConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacebookConnected.Get(), o.FacebookConnected.IsSet()
}

// HasFacebookConnected returns a boolean if a field has been set.
func (o *RealmUser) HasFacebookConnected() bool {
	if o != nil && o.FacebookConnected.IsSet() {
		return true
	}

	return false
}

// SetFacebookConnected gets a reference to the given NullableBool and assigns it to the FacebookConnected field.
func (o *RealmUser) SetFacebookConnected(v bool) {
	o.FacebookConnected.Set(&v)
}
// SetFacebookConnectedNil sets the value for FacebookConnected to be an explicit nil
func (o *RealmUser) SetFacebookConnectedNil() {
	o.FacebookConnected.Set(nil)
}

// UnsetFacebookConnected ensures that no value is present for FacebookConnected, not even an explicit nil
func (o *RealmUser) UnsetFacebookConnected() {
	o.FacebookConnected.Unset()
}

// GetFollowPending returns the FollowPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowPending() bool {
	if o == nil || IsNil(o.FollowPending.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowPending.Get()
}

// GetFollowPendingOk returns a tuple with the FollowPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowPending.Get(), o.FollowPending.IsSet()
}

// HasFollowPending returns a boolean if a field has been set.
func (o *RealmUser) HasFollowPending() bool {
	if o != nil && o.FollowPending.IsSet() {
		return true
	}

	return false
}

// SetFollowPending gets a reference to the given NullableBool and assigns it to the FollowPending field.
func (o *RealmUser) SetFollowPending(v bool) {
	o.FollowPending.Set(&v)
}
// SetFollowPendingNil sets the value for FollowPending to be an explicit nil
func (o *RealmUser) SetFollowPendingNil() {
	o.FollowPending.Set(nil)
}

// UnsetFollowPending ensures that no value is present for FollowPending, not even an explicit nil
func (o *RealmUser) UnsetFollowPending() {
	o.FollowPending.Unset()
}

// GetFollowedBy returns the FollowedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowedBy() bool {
	if o == nil || IsNil(o.FollowedBy.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowedBy.Get()
}

// GetFollowedByOk returns a tuple with the FollowedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowedByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowedBy.Get(), o.FollowedBy.IsSet()
}

// HasFollowedBy returns a boolean if a field has been set.
func (o *RealmUser) HasFollowedBy() bool {
	if o != nil && o.FollowedBy.IsSet() {
		return true
	}

	return false
}

// SetFollowedBy gets a reference to the given NullableBool and assigns it to the FollowedBy field.
func (o *RealmUser) SetFollowedBy(v bool) {
	o.FollowedBy.Set(&v)
}
// SetFollowedByNil sets the value for FollowedBy to be an explicit nil
func (o *RealmUser) SetFollowedByNil() {
	o.FollowedBy.Set(nil)
}

// UnsetFollowedBy ensures that no value is present for FollowedBy, not even an explicit nil
func (o *RealmUser) UnsetFollowedBy() {
	o.FollowedBy.Unset()
}

// GetFollowedByPending returns the FollowedByPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowedByPending() bool {
	if o == nil || IsNil(o.FollowedByPending.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowedByPending.Get()
}

// GetFollowedByPendingOk returns a tuple with the FollowedByPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowedByPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowedByPending.Get(), o.FollowedByPending.IsSet()
}

// HasFollowedByPending returns a boolean if a field has been set.
func (o *RealmUser) HasFollowedByPending() bool {
	if o != nil && o.FollowedByPending.IsSet() {
		return true
	}

	return false
}

// SetFollowedByPending gets a reference to the given NullableBool and assigns it to the FollowedByPending field.
func (o *RealmUser) SetFollowedByPending(v bool) {
	o.FollowedByPending.Set(&v)
}
// SetFollowedByPendingNil sets the value for FollowedByPending to be an explicit nil
func (o *RealmUser) SetFollowedByPendingNil() {
	o.FollowedByPending.Set(nil)
}

// UnsetFollowedByPending ensures that no value is present for FollowedByPending, not even an explicit nil
func (o *RealmUser) UnsetFollowedByPending() {
	o.FollowedByPending.Unset()
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *RealmUser) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *RealmUser) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *RealmUser) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *RealmUser) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFollowing returns the Following field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowing() bool {
	if o == nil || IsNil(o.Following.Get()) {
		var ret bool
		return ret
	}
	return *o.Following.Get()
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Following.Get(), o.Following.IsSet()
}

// HasFollowing returns a boolean if a field has been set.
func (o *RealmUser) HasFollowing() bool {
	if o != nil && o.Following.IsSet() {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given NullableBool and assigns it to the Following field.
func (o *RealmUser) SetFollowing(v bool) {
	o.Following.Set(&v)
}
// SetFollowingNil sets the value for Following to be an explicit nil
func (o *RealmUser) SetFollowingNil() {
	o.Following.Set(nil)
}

// UnsetFollowing ensures that no value is present for Following, not even an explicit nil
func (o *RealmUser) UnsetFollowing() {
	o.Following.Unset()
}

// GetFollowingsCount returns the FollowingsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFollowingsCount() int32 {
	if o == nil || IsNil(o.FollowingsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowingsCount.Get()
}

// GetFollowingsCountOk returns a tuple with the FollowingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFollowingsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingsCount.Get(), o.FollowingsCount.IsSet()
}

// HasFollowingsCount returns a boolean if a field has been set.
func (o *RealmUser) HasFollowingsCount() bool {
	if o != nil && o.FollowingsCount.IsSet() {
		return true
	}

	return false
}

// SetFollowingsCount gets a reference to the given NullableInt32 and assigns it to the FollowingsCount field.
func (o *RealmUser) SetFollowingsCount(v int32) {
	o.FollowingsCount.Set(&v)
}
// SetFollowingsCountNil sets the value for FollowingsCount to be an explicit nil
func (o *RealmUser) SetFollowingsCountNil() {
	o.FollowingsCount.Set(nil)
}

// UnsetFollowingsCount ensures that no value is present for FollowingsCount, not even an explicit nil
func (o *RealmUser) UnsetFollowingsCount() {
	o.FollowingsCount.Unset()
}

// GetFrame returns the Frame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFrame() string {
	if o == nil || IsNil(o.Frame.Get()) {
		var ret string
		return ret
	}
	return *o.Frame.Get()
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFrameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frame.Get(), o.Frame.IsSet()
}

// HasFrame returns a boolean if a field has been set.
func (o *RealmUser) HasFrame() bool {
	if o != nil && o.Frame.IsSet() {
		return true
	}

	return false
}

// SetFrame gets a reference to the given NullableString and assigns it to the Frame field.
func (o *RealmUser) SetFrame(v string) {
	o.Frame.Set(&v)
}
// SetFrameNil sets the value for Frame to be an explicit nil
func (o *RealmUser) SetFrameNil() {
	o.Frame.Set(nil)
}

// UnsetFrame ensures that no value is present for Frame, not even an explicit nil
func (o *RealmUser) UnsetFrame() {
	o.Frame.Unset()
}

// GetFrameThumbnail returns the FrameThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFrameThumbnail() string {
	if o == nil || IsNil(o.FrameThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.FrameThumbnail.Get()
}

// GetFrameThumbnailOk returns a tuple with the FrameThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFrameThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameThumbnail.Get(), o.FrameThumbnail.IsSet()
}

// HasFrameThumbnail returns a boolean if a field has been set.
func (o *RealmUser) HasFrameThumbnail() bool {
	if o != nil && o.FrameThumbnail.IsSet() {
		return true
	}

	return false
}

// SetFrameThumbnail gets a reference to the given NullableString and assigns it to the FrameThumbnail field.
func (o *RealmUser) SetFrameThumbnail(v string) {
	o.FrameThumbnail.Set(&v)
}
// SetFrameThumbnailNil sets the value for FrameThumbnail to be an explicit nil
func (o *RealmUser) SetFrameThumbnailNil() {
	o.FrameThumbnail.Set(nil)
}

// UnsetFrameThumbnail ensures that no value is present for FrameThumbnail, not even an explicit nil
func (o *RealmUser) UnsetFrameThumbnail() {
	o.FrameThumbnail.Unset()
}

// GetFromDifferentGenerationAndTrusted returns the FromDifferentGenerationAndTrusted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetFromDifferentGenerationAndTrusted() bool {
	if o == nil || IsNil(o.FromDifferentGenerationAndTrusted.Get()) {
		var ret bool
		return ret
	}
	return *o.FromDifferentGenerationAndTrusted.Get()
}

// GetFromDifferentGenerationAndTrustedOk returns a tuple with the FromDifferentGenerationAndTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetFromDifferentGenerationAndTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromDifferentGenerationAndTrusted.Get(), o.FromDifferentGenerationAndTrusted.IsSet()
}

// HasFromDifferentGenerationAndTrusted returns a boolean if a field has been set.
func (o *RealmUser) HasFromDifferentGenerationAndTrusted() bool {
	if o != nil && o.FromDifferentGenerationAndTrusted.IsSet() {
		return true
	}

	return false
}

// SetFromDifferentGenerationAndTrusted gets a reference to the given NullableBool and assigns it to the FromDifferentGenerationAndTrusted field.
func (o *RealmUser) SetFromDifferentGenerationAndTrusted(v bool) {
	o.FromDifferentGenerationAndTrusted.Set(&v)
}
// SetFromDifferentGenerationAndTrustedNil sets the value for FromDifferentGenerationAndTrusted to be an explicit nil
func (o *RealmUser) SetFromDifferentGenerationAndTrustedNil() {
	o.FromDifferentGenerationAndTrusted.Set(nil)
}

// UnsetFromDifferentGenerationAndTrusted ensures that no value is present for FromDifferentGenerationAndTrusted, not even an explicit nil
func (o *RealmUser) UnsetFromDifferentGenerationAndTrusted() {
	o.FromDifferentGenerationAndTrusted.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGender() int32 {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret int32
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGenderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *RealmUser) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableInt32 and assigns it to the Gender field.
func (o *RealmUser) SetGender(v int32) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *RealmUser) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *RealmUser) UnsetGender() {
	o.Gender.Unset()
}

// GetGeneration returns the Generation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGeneration() int32 {
	if o == nil || IsNil(o.Generation.Get()) {
		var ret int32
		return ret
	}
	return *o.Generation.Get()
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGenerationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generation.Get(), o.Generation.IsSet()
}

// HasGeneration returns a boolean if a field has been set.
func (o *RealmUser) HasGeneration() bool {
	if o != nil && o.Generation.IsSet() {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given NullableInt32 and assigns it to the Generation field.
func (o *RealmUser) SetGeneration(v int32) {
	o.Generation.Set(&v)
}
// SetGenerationNil sets the value for Generation to be an explicit nil
func (o *RealmUser) SetGenerationNil() {
	o.Generation.Set(nil)
}

// UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
func (o *RealmUser) UnsetGeneration() {
	o.Generation.Unset()
}

// GetGiftingAbility returns the GiftingAbility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGiftingAbility() RealmGiftingAbility {
	if o == nil || IsNil(o.GiftingAbility.Get()) {
		var ret RealmGiftingAbility
		return ret
	}
	return *o.GiftingAbility.Get()
}

// GetGiftingAbilityOk returns a tuple with the GiftingAbility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGiftingAbilityOk() (*RealmGiftingAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftingAbility.Get(), o.GiftingAbility.IsSet()
}

// HasGiftingAbility returns a boolean if a field has been set.
func (o *RealmUser) HasGiftingAbility() bool {
	if o != nil && o.GiftingAbility.IsSet() {
		return true
	}

	return false
}

// SetGiftingAbility gets a reference to the given NullableRealmGiftingAbility and assigns it to the GiftingAbility field.
func (o *RealmUser) SetGiftingAbility(v RealmGiftingAbility) {
	o.GiftingAbility.Set(&v)
}
// SetGiftingAbilityNil sets the value for GiftingAbility to be an explicit nil
func (o *RealmUser) SetGiftingAbilityNil() {
	o.GiftingAbility.Set(nil)
}

// UnsetGiftingAbility ensures that no value is present for GiftingAbility, not even an explicit nil
func (o *RealmUser) UnsetGiftingAbility() {
	o.GiftingAbility.Unset()
}

// GetGoogleConnected returns the GoogleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGoogleConnected() bool {
	if o == nil || IsNil(o.GoogleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.GoogleConnected.Get()
}

// GetGoogleConnectedOk returns a tuple with the GoogleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGoogleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoogleConnected.Get(), o.GoogleConnected.IsSet()
}

// HasGoogleConnected returns a boolean if a field has been set.
func (o *RealmUser) HasGoogleConnected() bool {
	if o != nil && o.GoogleConnected.IsSet() {
		return true
	}

	return false
}

// SetGoogleConnected gets a reference to the given NullableBool and assigns it to the GoogleConnected field.
func (o *RealmUser) SetGoogleConnected(v bool) {
	o.GoogleConnected.Set(&v)
}
// SetGoogleConnectedNil sets the value for GoogleConnected to be an explicit nil
func (o *RealmUser) SetGoogleConnectedNil() {
	o.GoogleConnected.Set(nil)
}

// UnsetGoogleConnected ensures that no value is present for GoogleConnected, not even an explicit nil
func (o *RealmUser) UnsetGoogleConnected() {
	o.GoogleConnected.Unset()
}

// GetGroupPhoneOn returns the GroupPhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGroupPhoneOn() bool {
	if o == nil || IsNil(o.GroupPhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.GroupPhoneOn.Get()
}

// GetGroupPhoneOnOk returns a tuple with the GroupPhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGroupPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupPhoneOn.Get(), o.GroupPhoneOn.IsSet()
}

// HasGroupPhoneOn returns a boolean if a field has been set.
func (o *RealmUser) HasGroupPhoneOn() bool {
	if o != nil && o.GroupPhoneOn.IsSet() {
		return true
	}

	return false
}

// SetGroupPhoneOn gets a reference to the given NullableBool and assigns it to the GroupPhoneOn field.
func (o *RealmUser) SetGroupPhoneOn(v bool) {
	o.GroupPhoneOn.Set(&v)
}
// SetGroupPhoneOnNil sets the value for GroupPhoneOn to be an explicit nil
func (o *RealmUser) SetGroupPhoneOnNil() {
	o.GroupPhoneOn.Set(nil)
}

// UnsetGroupPhoneOn ensures that no value is present for GroupPhoneOn, not even an explicit nil
func (o *RealmUser) UnsetGroupPhoneOn() {
	o.GroupPhoneOn.Unset()
}

// GetGroupUser returns the GroupUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGroupUser() GroupUser {
	if o == nil || IsNil(o.GroupUser.Get()) {
		var ret GroupUser
		return ret
	}
	return *o.GroupUser.Get()
}

// GetGroupUserOk returns a tuple with the GroupUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGroupUserOk() (*GroupUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupUser.Get(), o.GroupUser.IsSet()
}

// HasGroupUser returns a boolean if a field has been set.
func (o *RealmUser) HasGroupUser() bool {
	if o != nil && o.GroupUser.IsSet() {
		return true
	}

	return false
}

// SetGroupUser gets a reference to the given NullableGroupUser and assigns it to the GroupUser field.
func (o *RealmUser) SetGroupUser(v GroupUser) {
	o.GroupUser.Set(&v)
}
// SetGroupUserNil sets the value for GroupUser to be an explicit nil
func (o *RealmUser) SetGroupUserNil() {
	o.GroupUser.Set(nil)
}

// UnsetGroupUser ensures that no value is present for GroupUser, not even an explicit nil
func (o *RealmUser) UnsetGroupUser() {
	o.GroupUser.Unset()
}

// GetGroupVideoOn returns the GroupVideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGroupVideoOn() bool {
	if o == nil || IsNil(o.GroupVideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.GroupVideoOn.Get()
}

// GetGroupVideoOnOk returns a tuple with the GroupVideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGroupVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupVideoOn.Get(), o.GroupVideoOn.IsSet()
}

// HasGroupVideoOn returns a boolean if a field has been set.
func (o *RealmUser) HasGroupVideoOn() bool {
	if o != nil && o.GroupVideoOn.IsSet() {
		return true
	}

	return false
}

// SetGroupVideoOn gets a reference to the given NullableBool and assigns it to the GroupVideoOn field.
func (o *RealmUser) SetGroupVideoOn(v bool) {
	o.GroupVideoOn.Set(&v)
}
// SetGroupVideoOnNil sets the value for GroupVideoOn to be an explicit nil
func (o *RealmUser) SetGroupVideoOnNil() {
	o.GroupVideoOn.Set(nil)
}

// UnsetGroupVideoOn ensures that no value is present for GroupVideoOn, not even an explicit nil
func (o *RealmUser) UnsetGroupVideoOn() {
	o.GroupVideoOn.Unset()
}

// GetGroupsUsersCount returns the GroupsUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetGroupsUsersCount() int32 {
	if o == nil || IsNil(o.GroupsUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.GroupsUsersCount.Get()
}

// GetGroupsUsersCountOk returns a tuple with the GroupsUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetGroupsUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsUsersCount.Get(), o.GroupsUsersCount.IsSet()
}

// HasGroupsUsersCount returns a boolean if a field has been set.
func (o *RealmUser) HasGroupsUsersCount() bool {
	if o != nil && o.GroupsUsersCount.IsSet() {
		return true
	}

	return false
}

// SetGroupsUsersCount gets a reference to the given NullableInt32 and assigns it to the GroupsUsersCount field.
func (o *RealmUser) SetGroupsUsersCount(v int32) {
	o.GroupsUsersCount.Set(&v)
}
// SetGroupsUsersCountNil sets the value for GroupsUsersCount to be an explicit nil
func (o *RealmUser) SetGroupsUsersCountNil() {
	o.GroupsUsersCount.Set(nil)
}

// UnsetGroupsUsersCount ensures that no value is present for GroupsUsersCount, not even an explicit nil
func (o *RealmUser) UnsetGroupsUsersCount() {
	o.GroupsUsersCount.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetHidden() bool {
	if o == nil || IsNil(o.Hidden.Get()) {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *RealmUser) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *RealmUser) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *RealmUser) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *RealmUser) UnsetHidden() {
	o.Hidden.Unset()
}

// GetHideVip returns the HideVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetHideVip() bool {
	if o == nil || IsNil(o.HideVip.Get()) {
		var ret bool
		return ret
	}
	return *o.HideVip.Get()
}

// GetHideVipOk returns a tuple with the HideVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetHideVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideVip.Get(), o.HideVip.IsSet()
}

// HasHideVip returns a boolean if a field has been set.
func (o *RealmUser) HasHideVip() bool {
	if o != nil && o.HideVip.IsSet() {
		return true
	}

	return false
}

// SetHideVip gets a reference to the given NullableBool and assigns it to the HideVip field.
func (o *RealmUser) SetHideVip(v bool) {
	o.HideVip.Set(&v)
}
// SetHideVipNil sets the value for HideVip to be an explicit nil
func (o *RealmUser) SetHideVipNil() {
	o.HideVip.Set(nil)
}

// UnsetHideVip ensures that no value is present for HideVip, not even an explicit nil
func (o *RealmUser) UnsetHideVip() {
	o.HideVip.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmUser) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmUser) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmUser) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmUser) UnsetId() {
	o.Id.Unset()
}

// GetInterestsCount returns the InterestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetInterestsCount() int32 {
	if o == nil || IsNil(o.InterestsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InterestsCount.Get()
}

// GetInterestsCountOk returns a tuple with the InterestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetInterestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterestsCount.Get(), o.InterestsCount.IsSet()
}

// HasInterestsCount returns a boolean if a field has been set.
func (o *RealmUser) HasInterestsCount() bool {
	if o != nil && o.InterestsCount.IsSet() {
		return true
	}

	return false
}

// SetInterestsCount gets a reference to the given NullableInt32 and assigns it to the InterestsCount field.
func (o *RealmUser) SetInterestsCount(v int32) {
	o.InterestsCount.Set(&v)
}
// SetInterestsCountNil sets the value for InterestsCount to be an explicit nil
func (o *RealmUser) SetInterestsCountNil() {
	o.InterestsCount.Set(nil)
}

// UnsetInterestsCount ensures that no value is present for InterestsCount, not even an explicit nil
func (o *RealmUser) UnsetInterestsCount() {
	o.InterestsCount.Unset()
}

// GetInterestsSelected returns the InterestsSelected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetInterestsSelected() bool {
	if o == nil || IsNil(o.InterestsSelected.Get()) {
		var ret bool
		return ret
	}
	return *o.InterestsSelected.Get()
}

// GetInterestsSelectedOk returns a tuple with the InterestsSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetInterestsSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterestsSelected.Get(), o.InterestsSelected.IsSet()
}

// HasInterestsSelected returns a boolean if a field has been set.
func (o *RealmUser) HasInterestsSelected() bool {
	if o != nil && o.InterestsSelected.IsSet() {
		return true
	}

	return false
}

// SetInterestsSelected gets a reference to the given NullableBool and assigns it to the InterestsSelected field.
func (o *RealmUser) SetInterestsSelected(v bool) {
	o.InterestsSelected.Set(&v)
}
// SetInterestsSelectedNil sets the value for InterestsSelected to be an explicit nil
func (o *RealmUser) SetInterestsSelectedNil() {
	o.InterestsSelected.Set(nil)
}

// UnsetInterestsSelected ensures that no value is present for InterestsSelected, not even an explicit nil
func (o *RealmUser) UnsetInterestsSelected() {
	o.InterestsSelected.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *RealmUser) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *RealmUser) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *RealmUser) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *RealmUser) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetLastLoggedInAtSeconds returns the LastLoggedInAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetLastLoggedInAtSeconds() int64 {
	if o == nil || IsNil(o.LastLoggedInAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.LastLoggedInAtSeconds.Get()
}

// GetLastLoggedInAtSecondsOk returns a tuple with the LastLoggedInAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetLastLoggedInAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoggedInAtSeconds.Get(), o.LastLoggedInAtSeconds.IsSet()
}

// HasLastLoggedInAtSeconds returns a boolean if a field has been set.
func (o *RealmUser) HasLastLoggedInAtSeconds() bool {
	if o != nil && o.LastLoggedInAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetLastLoggedInAtSeconds gets a reference to the given NullableInt64 and assigns it to the LastLoggedInAtSeconds field.
func (o *RealmUser) SetLastLoggedInAtSeconds(v int64) {
	o.LastLoggedInAtSeconds.Set(&v)
}
// SetLastLoggedInAtSecondsNil sets the value for LastLoggedInAtSeconds to be an explicit nil
func (o *RealmUser) SetLastLoggedInAtSecondsNil() {
	o.LastLoggedInAtSeconds.Set(nil)
}

// UnsetLastLoggedInAtSeconds ensures that no value is present for LastLoggedInAtSeconds, not even an explicit nil
func (o *RealmUser) UnsetLastLoggedInAtSeconds() {
	o.LastLoggedInAtSeconds.Unset()
}

// GetLineConnected returns the LineConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetLineConnected() bool {
	if o == nil || IsNil(o.LineConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.LineConnected.Get()
}

// GetLineConnectedOk returns a tuple with the LineConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetLineConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineConnected.Get(), o.LineConnected.IsSet()
}

// HasLineConnected returns a boolean if a field has been set.
func (o *RealmUser) HasLineConnected() bool {
	if o != nil && o.LineConnected.IsSet() {
		return true
	}

	return false
}

// SetLineConnected gets a reference to the given NullableBool and assigns it to the LineConnected field.
func (o *RealmUser) SetLineConnected(v bool) {
	o.LineConnected.Set(&v)
}
// SetLineConnectedNil sets the value for LineConnected to be an explicit nil
func (o *RealmUser) SetLineConnectedNil() {
	o.LineConnected.Set(nil)
}

// UnsetLineConnected ensures that no value is present for LineConnected, not even an explicit nil
func (o *RealmUser) UnsetLineConnected() {
	o.LineConnected.Unset()
}

// GetLoginStreakCount returns the LoginStreakCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetLoginStreakCount() int32 {
	if o == nil || IsNil(o.LoginStreakCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LoginStreakCount.Get()
}

// GetLoginStreakCountOk returns a tuple with the LoginStreakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetLoginStreakCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginStreakCount.Get(), o.LoginStreakCount.IsSet()
}

// HasLoginStreakCount returns a boolean if a field has been set.
func (o *RealmUser) HasLoginStreakCount() bool {
	if o != nil && o.LoginStreakCount.IsSet() {
		return true
	}

	return false
}

// SetLoginStreakCount gets a reference to the given NullableInt32 and assigns it to the LoginStreakCount field.
func (o *RealmUser) SetLoginStreakCount(v int32) {
	o.LoginStreakCount.Set(&v)
}
// SetLoginStreakCountNil sets the value for LoginStreakCount to be an explicit nil
func (o *RealmUser) SetLoginStreakCountNil() {
	o.LoginStreakCount.Set(nil)
}

// UnsetLoginStreakCount ensures that no value is present for LoginStreakCount, not even an explicit nil
func (o *RealmUser) UnsetLoginStreakCount() {
	o.LoginStreakCount.Unset()
}

// GetMaskedEmail returns the MaskedEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetMaskedEmail() string {
	if o == nil || IsNil(o.MaskedEmail.Get()) {
		var ret string
		return ret
	}
	return *o.MaskedEmail.Get()
}

// GetMaskedEmailOk returns a tuple with the MaskedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetMaskedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedEmail.Get(), o.MaskedEmail.IsSet()
}

// HasMaskedEmail returns a boolean if a field has been set.
func (o *RealmUser) HasMaskedEmail() bool {
	if o != nil && o.MaskedEmail.IsSet() {
		return true
	}

	return false
}

// SetMaskedEmail gets a reference to the given NullableString and assigns it to the MaskedEmail field.
func (o *RealmUser) SetMaskedEmail(v string) {
	o.MaskedEmail.Set(&v)
}
// SetMaskedEmailNil sets the value for MaskedEmail to be an explicit nil
func (o *RealmUser) SetMaskedEmailNil() {
	o.MaskedEmail.Set(nil)
}

// UnsetMaskedEmail ensures that no value is present for MaskedEmail, not even an explicit nil
func (o *RealmUser) UnsetMaskedEmail() {
	o.MaskedEmail.Unset()
}

// GetMediaCount returns the MediaCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetMediaCount() int32 {
	if o == nil || IsNil(o.MediaCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MediaCount.Get()
}

// GetMediaCountOk returns a tuple with the MediaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetMediaCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaCount.Get(), o.MediaCount.IsSet()
}

// HasMediaCount returns a boolean if a field has been set.
func (o *RealmUser) HasMediaCount() bool {
	if o != nil && o.MediaCount.IsSet() {
		return true
	}

	return false
}

// SetMediaCount gets a reference to the given NullableInt32 and assigns it to the MediaCount field.
func (o *RealmUser) SetMediaCount(v int32) {
	o.MediaCount.Set(&v)
}
// SetMediaCountNil sets the value for MediaCount to be an explicit nil
func (o *RealmUser) SetMediaCountNil() {
	o.MediaCount.Set(nil)
}

// UnsetMediaCount ensures that no value is present for MediaCount, not even an explicit nil
func (o *RealmUser) UnsetMediaCount() {
	o.MediaCount.Unset()
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetMobileNumber() string {
	if o == nil || IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *RealmUser) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *RealmUser) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *RealmUser) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *RealmUser) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetNewUser returns the NewUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetNewUser() bool {
	if o == nil || IsNil(o.NewUser.Get()) {
		var ret bool
		return ret
	}
	return *o.NewUser.Get()
}

// GetNewUserOk returns a tuple with the NewUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetNewUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewUser.Get(), o.NewUser.IsSet()
}

// HasNewUser returns a boolean if a field has been set.
func (o *RealmUser) HasNewUser() bool {
	if o != nil && o.NewUser.IsSet() {
		return true
	}

	return false
}

// SetNewUser gets a reference to the given NullableBool and assigns it to the NewUser field.
func (o *RealmUser) SetNewUser(v bool) {
	o.NewUser.Set(&v)
}
// SetNewUserNil sets the value for NewUser to be an explicit nil
func (o *RealmUser) SetNewUserNil() {
	o.NewUser.Set(nil)
}

// UnsetNewUser ensures that no value is present for NewUser, not even an explicit nil
func (o *RealmUser) UnsetNewUser() {
	o.NewUser.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *RealmUser) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *RealmUser) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *RealmUser) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *RealmUser) UnsetNickname() {
	o.Nickname.Unset()
}

// GetOnlineStatus returns the OnlineStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetOnlineStatus() OnlineStatusEnum {
	if o == nil || IsNil(o.OnlineStatus.Get()) {
		var ret OnlineStatusEnum
		return ret
	}
	return *o.OnlineStatus.Get()
}

// GetOnlineStatusOk returns a tuple with the OnlineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetOnlineStatusOk() (*OnlineStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineStatus.Get(), o.OnlineStatus.IsSet()
}

// HasOnlineStatus returns a boolean if a field has been set.
func (o *RealmUser) HasOnlineStatus() bool {
	if o != nil && o.OnlineStatus.IsSet() {
		return true
	}

	return false
}

// SetOnlineStatus gets a reference to the given NullableOnlineStatusEnum and assigns it to the OnlineStatus field.
func (o *RealmUser) SetOnlineStatus(v OnlineStatusEnum) {
	o.OnlineStatus.Set(&v)
}
// SetOnlineStatusNil sets the value for OnlineStatus to be an explicit nil
func (o *RealmUser) SetOnlineStatusNil() {
	o.OnlineStatus.Set(nil)
}

// UnsetOnlineStatus ensures that no value is present for OnlineStatus, not even an explicit nil
func (o *RealmUser) UnsetOnlineStatus() {
	o.OnlineStatus.Unset()
}

// GetPhoneOn returns the PhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetPhoneOn() bool {
	if o == nil || IsNil(o.PhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.PhoneOn.Get()
}

// GetPhoneOnOk returns a tuple with the PhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneOn.Get(), o.PhoneOn.IsSet()
}

// HasPhoneOn returns a boolean if a field has been set.
func (o *RealmUser) HasPhoneOn() bool {
	if o != nil && o.PhoneOn.IsSet() {
		return true
	}

	return false
}

// SetPhoneOn gets a reference to the given NullableBool and assigns it to the PhoneOn field.
func (o *RealmUser) SetPhoneOn(v bool) {
	o.PhoneOn.Set(&v)
}
// SetPhoneOnNil sets the value for PhoneOn to be an explicit nil
func (o *RealmUser) SetPhoneOnNil() {
	o.PhoneOn.Set(nil)
}

// UnsetPhoneOn ensures that no value is present for PhoneOn, not even an explicit nil
func (o *RealmUser) UnsetPhoneOn() {
	o.PhoneOn.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetPostsCount() int32 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetPostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *RealmUser) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt32 and assigns it to the PostsCount field.
func (o *RealmUser) SetPostsCount(v int32) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *RealmUser) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *RealmUser) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetPrefecture returns the Prefecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetPrefecture() string {
	if o == nil || IsNil(o.Prefecture.Get()) {
		var ret string
		return ret
	}
	return *o.Prefecture.Get()
}

// GetPrefectureOk returns a tuple with the Prefecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetPrefectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefecture.Get(), o.Prefecture.IsSet()
}

// HasPrefecture returns a boolean if a field has been set.
func (o *RealmUser) HasPrefecture() bool {
	if o != nil && o.Prefecture.IsSet() {
		return true
	}

	return false
}

// SetPrefecture gets a reference to the given NullableString and assigns it to the Prefecture field.
func (o *RealmUser) SetPrefecture(v string) {
	o.Prefecture.Set(&v)
}
// SetPrefectureNil sets the value for Prefecture to be an explicit nil
func (o *RealmUser) SetPrefectureNil() {
	o.Prefecture.Set(nil)
}

// UnsetPrefecture ensures that no value is present for Prefecture, not even an explicit nil
func (o *RealmUser) UnsetPrefecture() {
	o.Prefecture.Unset()
}

// GetProfileIcon returns the ProfileIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetProfileIcon() string {
	if o == nil || IsNil(o.ProfileIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIcon.Get()
}

// GetProfileIconOk returns a tuple with the ProfileIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetProfileIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIcon.Get(), o.ProfileIcon.IsSet()
}

// HasProfileIcon returns a boolean if a field has been set.
func (o *RealmUser) HasProfileIcon() bool {
	if o != nil && o.ProfileIcon.IsSet() {
		return true
	}

	return false
}

// SetProfileIcon gets a reference to the given NullableString and assigns it to the ProfileIcon field.
func (o *RealmUser) SetProfileIcon(v string) {
	o.ProfileIcon.Set(&v)
}
// SetProfileIconNil sets the value for ProfileIcon to be an explicit nil
func (o *RealmUser) SetProfileIconNil() {
	o.ProfileIcon.Set(nil)
}

// UnsetProfileIcon ensures that no value is present for ProfileIcon, not even an explicit nil
func (o *RealmUser) UnsetProfileIcon() {
	o.ProfileIcon.Unset()
}

// GetProfileIconThumbnail returns the ProfileIconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetProfileIconThumbnail() string {
	if o == nil || IsNil(o.ProfileIconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIconThumbnail.Get()
}

// GetProfileIconThumbnailOk returns a tuple with the ProfileIconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetProfileIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIconThumbnail.Get(), o.ProfileIconThumbnail.IsSet()
}

// HasProfileIconThumbnail returns a boolean if a field has been set.
func (o *RealmUser) HasProfileIconThumbnail() bool {
	if o != nil && o.ProfileIconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetProfileIconThumbnail gets a reference to the given NullableString and assigns it to the ProfileIconThumbnail field.
func (o *RealmUser) SetProfileIconThumbnail(v string) {
	o.ProfileIconThumbnail.Set(&v)
}
// SetProfileIconThumbnailNil sets the value for ProfileIconThumbnail to be an explicit nil
func (o *RealmUser) SetProfileIconThumbnailNil() {
	o.ProfileIconThumbnail.Set(nil)
}

// UnsetProfileIconThumbnail ensures that no value is present for ProfileIconThumbnail, not even an explicit nil
func (o *RealmUser) UnsetProfileIconThumbnail() {
	o.ProfileIconThumbnail.Unset()
}

// GetPushNotification returns the PushNotification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetPushNotification() bool {
	if o == nil || IsNil(o.PushNotification.Get()) {
		var ret bool
		return ret
	}
	return *o.PushNotification.Get()
}

// GetPushNotificationOk returns a tuple with the PushNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetPushNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushNotification.Get(), o.PushNotification.IsSet()
}

// HasPushNotification returns a boolean if a field has been set.
func (o *RealmUser) HasPushNotification() bool {
	if o != nil && o.PushNotification.IsSet() {
		return true
	}

	return false
}

// SetPushNotification gets a reference to the given NullableBool and assigns it to the PushNotification field.
func (o *RealmUser) SetPushNotification(v bool) {
	o.PushNotification.Set(&v)
}
// SetPushNotificationNil sets the value for PushNotification to be an explicit nil
func (o *RealmUser) SetPushNotificationNil() {
	o.PushNotification.Set(nil)
}

// UnsetPushNotification ensures that no value is present for PushNotification, not even an explicit nil
func (o *RealmUser) UnsetPushNotification() {
	o.PushNotification.Unset()
}

// GetRecentlyKenta returns the RecentlyKenta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetRecentlyKenta() bool {
	if o == nil || IsNil(o.RecentlyKenta.Get()) {
		var ret bool
		return ret
	}
	return *o.RecentlyKenta.Get()
}

// GetRecentlyKentaOk returns a tuple with the RecentlyKenta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetRecentlyKentaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecentlyKenta.Get(), o.RecentlyKenta.IsSet()
}

// HasRecentlyKenta returns a boolean if a field has been set.
func (o *RealmUser) HasRecentlyKenta() bool {
	if o != nil && o.RecentlyKenta.IsSet() {
		return true
	}

	return false
}

// SetRecentlyKenta gets a reference to the given NullableBool and assigns it to the RecentlyKenta field.
func (o *RealmUser) SetRecentlyKenta(v bool) {
	o.RecentlyKenta.Set(&v)
}
// SetRecentlyKentaNil sets the value for RecentlyKenta to be an explicit nil
func (o *RealmUser) SetRecentlyKentaNil() {
	o.RecentlyKenta.Set(nil)
}

// UnsetRecentlyKenta ensures that no value is present for RecentlyKenta, not even an explicit nil
func (o *RealmUser) UnsetRecentlyKenta() {
	o.RecentlyKenta.Unset()
}

// GetRestrictedReviewBy returns the RestrictedReviewBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetRestrictedReviewBy() string {
	if o == nil || IsNil(o.RestrictedReviewBy.Get()) {
		var ret string
		return ret
	}
	return *o.RestrictedReviewBy.Get()
}

// GetRestrictedReviewByOk returns a tuple with the RestrictedReviewBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetRestrictedReviewByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestrictedReviewBy.Get(), o.RestrictedReviewBy.IsSet()
}

// HasRestrictedReviewBy returns a boolean if a field has been set.
func (o *RealmUser) HasRestrictedReviewBy() bool {
	if o != nil && o.RestrictedReviewBy.IsSet() {
		return true
	}

	return false
}

// SetRestrictedReviewBy gets a reference to the given NullableString and assigns it to the RestrictedReviewBy field.
func (o *RealmUser) SetRestrictedReviewBy(v string) {
	o.RestrictedReviewBy.Set(&v)
}
// SetRestrictedReviewByNil sets the value for RestrictedReviewBy to be an explicit nil
func (o *RealmUser) SetRestrictedReviewByNil() {
	o.RestrictedReviewBy.Set(nil)
}

// UnsetRestrictedReviewBy ensures that no value is present for RestrictedReviewBy, not even an explicit nil
func (o *RealmUser) UnsetRestrictedReviewBy() {
	o.RestrictedReviewBy.Unset()
}

// GetReviewsCount returns the ReviewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetReviewsCount() int32 {
	if o == nil || IsNil(o.ReviewsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReviewsCount.Get()
}

// GetReviewsCountOk returns a tuple with the ReviewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetReviewsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewsCount.Get(), o.ReviewsCount.IsSet()
}

// HasReviewsCount returns a boolean if a field has been set.
func (o *RealmUser) HasReviewsCount() bool {
	if o != nil && o.ReviewsCount.IsSet() {
		return true
	}

	return false
}

// SetReviewsCount gets a reference to the given NullableInt32 and assigns it to the ReviewsCount field.
func (o *RealmUser) SetReviewsCount(v int32) {
	o.ReviewsCount.Set(&v)
}
// SetReviewsCountNil sets the value for ReviewsCount to be an explicit nil
func (o *RealmUser) SetReviewsCountNil() {
	o.ReviewsCount.Set(nil)
}

// UnsetReviewsCount ensures that no value is present for ReviewsCount, not even an explicit nil
func (o *RealmUser) UnsetReviewsCount() {
	o.ReviewsCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RealmUser) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RealmUser) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RealmUser) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RealmUser) UnsetTitle() {
	o.Title.Unset()
}

// GetTotalGiftsReceivedCount returns the TotalGiftsReceivedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetTotalGiftsReceivedCount() int32 {
	if o == nil || IsNil(o.TotalGiftsReceivedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalGiftsReceivedCount.Get()
}

// GetTotalGiftsReceivedCountOk returns a tuple with the TotalGiftsReceivedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetTotalGiftsReceivedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalGiftsReceivedCount.Get(), o.TotalGiftsReceivedCount.IsSet()
}

// HasTotalGiftsReceivedCount returns a boolean if a field has been set.
func (o *RealmUser) HasTotalGiftsReceivedCount() bool {
	if o != nil && o.TotalGiftsReceivedCount.IsSet() {
		return true
	}

	return false
}

// SetTotalGiftsReceivedCount gets a reference to the given NullableInt32 and assigns it to the TotalGiftsReceivedCount field.
func (o *RealmUser) SetTotalGiftsReceivedCount(v int32) {
	o.TotalGiftsReceivedCount.Set(&v)
}
// SetTotalGiftsReceivedCountNil sets the value for TotalGiftsReceivedCount to be an explicit nil
func (o *RealmUser) SetTotalGiftsReceivedCountNil() {
	o.TotalGiftsReceivedCount.Set(nil)
}

// UnsetTotalGiftsReceivedCount ensures that no value is present for TotalGiftsReceivedCount, not even an explicit nil
func (o *RealmUser) UnsetTotalGiftsReceivedCount() {
	o.TotalGiftsReceivedCount.Unset()
}

// GetTwitterId returns the TwitterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetTwitterId() string {
	if o == nil || IsNil(o.TwitterId.Get()) {
		var ret string
		return ret
	}
	return *o.TwitterId.Get()
}

// GetTwitterIdOk returns a tuple with the TwitterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetTwitterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterId.Get(), o.TwitterId.IsSet()
}

// HasTwitterId returns a boolean if a field has been set.
func (o *RealmUser) HasTwitterId() bool {
	if o != nil && o.TwitterId.IsSet() {
		return true
	}

	return false
}

// SetTwitterId gets a reference to the given NullableString and assigns it to the TwitterId field.
func (o *RealmUser) SetTwitterId(v string) {
	o.TwitterId.Set(&v)
}
// SetTwitterIdNil sets the value for TwitterId to be an explicit nil
func (o *RealmUser) SetTwitterIdNil() {
	o.TwitterId.Set(nil)
}

// UnsetTwitterId ensures that no value is present for TwitterId, not even an explicit nil
func (o *RealmUser) UnsetTwitterId() {
	o.TwitterId.Unset()
}

// GetTwoFaEnabled returns the TwoFaEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetTwoFaEnabled() bool {
	if o == nil || IsNil(o.TwoFaEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TwoFaEnabled.Get()
}

// GetTwoFaEnabledOk returns a tuple with the TwoFaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetTwoFaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFaEnabled.Get(), o.TwoFaEnabled.IsSet()
}

// HasTwoFaEnabled returns a boolean if a field has been set.
func (o *RealmUser) HasTwoFaEnabled() bool {
	if o != nil && o.TwoFaEnabled.IsSet() {
		return true
	}

	return false
}

// SetTwoFaEnabled gets a reference to the given NullableBool and assigns it to the TwoFaEnabled field.
func (o *RealmUser) SetTwoFaEnabled(v bool) {
	o.TwoFaEnabled.Set(&v)
}
// SetTwoFaEnabledNil sets the value for TwoFaEnabled to be an explicit nil
func (o *RealmUser) SetTwoFaEnabledNil() {
	o.TwoFaEnabled.Set(nil)
}

// UnsetTwoFaEnabled ensures that no value is present for TwoFaEnabled, not even an explicit nil
func (o *RealmUser) UnsetTwoFaEnabled() {
	o.TwoFaEnabled.Unset()
}

// GetUpdatedTimeMillis returns the UpdatedTimeMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetUpdatedTimeMillis() int64 {
	if o == nil || IsNil(o.UpdatedTimeMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedTimeMillis.Get()
}

// GetUpdatedTimeMillisOk returns a tuple with the UpdatedTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetUpdatedTimeMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedTimeMillis.Get(), o.UpdatedTimeMillis.IsSet()
}

// HasUpdatedTimeMillis returns a boolean if a field has been set.
func (o *RealmUser) HasUpdatedTimeMillis() bool {
	if o != nil && o.UpdatedTimeMillis.IsSet() {
		return true
	}

	return false
}

// SetUpdatedTimeMillis gets a reference to the given NullableInt64 and assigns it to the UpdatedTimeMillis field.
func (o *RealmUser) SetUpdatedTimeMillis(v int64) {
	o.UpdatedTimeMillis.Set(&v)
}
// SetUpdatedTimeMillisNil sets the value for UpdatedTimeMillis to be an explicit nil
func (o *RealmUser) SetUpdatedTimeMillisNil() {
	o.UpdatedTimeMillis.Set(nil)
}

// UnsetUpdatedTimeMillis ensures that no value is present for UpdatedTimeMillis, not even an explicit nil
func (o *RealmUser) UnsetUpdatedTimeMillis() {
	o.UpdatedTimeMillis.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *RealmUser) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *RealmUser) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *RealmUser) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *RealmUser) UnsetUsername() {
	o.Username.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *RealmUser) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *RealmUser) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *RealmUser) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *RealmUser) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVideoOn returns the VideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetVideoOn() bool {
	if o == nil || IsNil(o.VideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.VideoOn.Get()
}

// GetVideoOnOk returns a tuple with the VideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoOn.Get(), o.VideoOn.IsSet()
}

// HasVideoOn returns a boolean if a field has been set.
func (o *RealmUser) HasVideoOn() bool {
	if o != nil && o.VideoOn.IsSet() {
		return true
	}

	return false
}

// SetVideoOn gets a reference to the given NullableBool and assigns it to the VideoOn field.
func (o *RealmUser) SetVideoOn(v bool) {
	o.VideoOn.Set(&v)
}
// SetVideoOnNil sets the value for VideoOn to be an explicit nil
func (o *RealmUser) SetVideoOnNil() {
	o.VideoOn.Set(nil)
}

// UnsetVideoOn ensures that no value is present for VideoOn, not even an explicit nil
func (o *RealmUser) UnsetVideoOn() {
	o.VideoOn.Unset()
}

// GetVip returns the Vip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetVip() bool {
	if o == nil || IsNil(o.Vip.Get()) {
		var ret bool
		return ret
	}
	return *o.Vip.Get()
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vip.Get(), o.Vip.IsSet()
}

// HasVip returns a boolean if a field has been set.
func (o *RealmUser) HasVip() bool {
	if o != nil && o.Vip.IsSet() {
		return true
	}

	return false
}

// SetVip gets a reference to the given NullableBool and assigns it to the Vip field.
func (o *RealmUser) SetVip(v bool) {
	o.Vip.Set(&v)
}
// SetVipNil sets the value for Vip to be an explicit nil
func (o *RealmUser) SetVipNil() {
	o.Vip.Set(nil)
}

// UnsetVip ensures that no value is present for Vip, not even an explicit nil
func (o *RealmUser) UnsetVip() {
	o.Vip.Unset()
}

// GetVipUntilSeconds returns the VipUntilSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetVipUntilSeconds() int64 {
	if o == nil || IsNil(o.VipUntilSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.VipUntilSeconds.Get()
}

// GetVipUntilSecondsOk returns a tuple with the VipUntilSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetVipUntilSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipUntilSeconds.Get(), o.VipUntilSeconds.IsSet()
}

// HasVipUntilSeconds returns a boolean if a field has been set.
func (o *RealmUser) HasVipUntilSeconds() bool {
	if o != nil && o.VipUntilSeconds.IsSet() {
		return true
	}

	return false
}

// SetVipUntilSeconds gets a reference to the given NullableInt64 and assigns it to the VipUntilSeconds field.
func (o *RealmUser) SetVipUntilSeconds(v int64) {
	o.VipUntilSeconds.Set(&v)
}
// SetVipUntilSecondsNil sets the value for VipUntilSeconds to be an explicit nil
func (o *RealmUser) SetVipUntilSecondsNil() {
	o.VipUntilSeconds.Set(nil)
}

// UnsetVipUntilSeconds ensures that no value is present for VipUntilSeconds, not even an explicit nil
func (o *RealmUser) UnsetVipUntilSeconds() {
	o.VipUntilSeconds.Unset()
}

// GetWorldIdConnected returns the WorldIdConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmUser) GetWorldIdConnected() bool {
	if o == nil || IsNil(o.WorldIdConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.WorldIdConnected.Get()
}

// GetWorldIdConnectedOk returns a tuple with the WorldIdConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmUser) GetWorldIdConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorldIdConnected.Get(), o.WorldIdConnected.IsSet()
}

// HasWorldIdConnected returns a boolean if a field has been set.
func (o *RealmUser) HasWorldIdConnected() bool {
	if o != nil && o.WorldIdConnected.IsSet() {
		return true
	}

	return false
}

// SetWorldIdConnected gets a reference to the given NullableBool and assigns it to the WorldIdConnected field.
func (o *RealmUser) SetWorldIdConnected(v bool) {
	o.WorldIdConnected.Set(&v)
}
// SetWorldIdConnectedNil sets the value for WorldIdConnected to be an explicit nil
func (o *RealmUser) SetWorldIdConnectedNil() {
	o.WorldIdConnected.Set(nil)
}

// UnsetWorldIdConnected ensures that no value is present for WorldIdConnected, not even an explicit nil
func (o *RealmUser) UnsetWorldIdConnected() {
	o.WorldIdConnected.Unset()
}

func (o RealmUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgeVerified.IsSet() {
		toSerialize["age_verified"] = o.AgeVerified.Get()
	}
	if o.AppleConnected.IsSet() {
		toSerialize["apple_connected"] = o.AppleConnected.Get()
	}
	if o.Biography.IsSet() {
		toSerialize["biography"] = o.Biography.Get()
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if o.BlockingLimit.IsSet() {
		toSerialize["blocking_limit"] = o.BlockingLimit.Get()
	}
	if o.ChatRequest.IsSet() {
		toSerialize["chat_request"] = o.ChatRequest.Get()
	}
	if o.ConnectedBy != nil {
		toSerialize["connected_by"] = o.ConnectedBy
	}
	if o.ContactPhones != nil {
		toSerialize["contact_phones"] = o.ContactPhones
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.CoverImageThumbnail.IsSet() {
		toSerialize["cover_image_thumbnail"] = o.CoverImageThumbnail.Get()
	}
	if o.CreatedAtSeconds.IsSet() {
		toSerialize["created_at_seconds"] = o.CreatedAtSeconds.Get()
	}
	if o.DangerousUser.IsSet() {
		toSerialize["dangerous_user"] = o.DangerousUser.Get()
	}
	if o.EmailConfirmed.IsSet() {
		toSerialize["email_confirmed"] = o.EmailConfirmed.Get()
	}
	if o.FacebookConnected.IsSet() {
		toSerialize["facebook_connected"] = o.FacebookConnected.Get()
	}
	if o.FollowPending.IsSet() {
		toSerialize["follow_pending"] = o.FollowPending.Get()
	}
	if o.FollowedBy.IsSet() {
		toSerialize["followed_by"] = o.FollowedBy.Get()
	}
	if o.FollowedByPending.IsSet() {
		toSerialize["followed_by_pending"] = o.FollowedByPending.Get()
	}
	if o.FollowersCount.IsSet() {
		toSerialize["followers_count"] = o.FollowersCount.Get()
	}
	if o.Following.IsSet() {
		toSerialize["following"] = o.Following.Get()
	}
	if o.FollowingsCount.IsSet() {
		toSerialize["followings_count"] = o.FollowingsCount.Get()
	}
	if o.Frame.IsSet() {
		toSerialize["frame"] = o.Frame.Get()
	}
	if o.FrameThumbnail.IsSet() {
		toSerialize["frame_thumbnail"] = o.FrameThumbnail.Get()
	}
	if o.FromDifferentGenerationAndTrusted.IsSet() {
		toSerialize["from_different_generation_and_trusted"] = o.FromDifferentGenerationAndTrusted.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Generation.IsSet() {
		toSerialize["generation"] = o.Generation.Get()
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
	if o.GroupUser.IsSet() {
		toSerialize["group_user"] = o.GroupUser.Get()
	}
	if o.GroupVideoOn.IsSet() {
		toSerialize["group_video_on"] = o.GroupVideoOn.Get()
	}
	if o.GroupsUsersCount.IsSet() {
		toSerialize["groups_users_count"] = o.GroupsUsersCount.Get()
	}
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	if o.HideVip.IsSet() {
		toSerialize["hide_vip"] = o.HideVip.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InterestsCount.IsSet() {
		toSerialize["interests_count"] = o.InterestsCount.Get()
	}
	if o.InterestsSelected.IsSet() {
		toSerialize["interests_selected"] = o.InterestsSelected.Get()
	}
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.LastLoggedInAtSeconds.IsSet() {
		toSerialize["last_logged_in_at_seconds"] = o.LastLoggedInAtSeconds.Get()
	}
	if o.LineConnected.IsSet() {
		toSerialize["line_connected"] = o.LineConnected.Get()
	}
	if o.LoginStreakCount.IsSet() {
		toSerialize["login_streak_count"] = o.LoginStreakCount.Get()
	}
	if o.MaskedEmail.IsSet() {
		toSerialize["masked_email"] = o.MaskedEmail.Get()
	}
	if o.MediaCount.IsSet() {
		toSerialize["media_count"] = o.MediaCount.Get()
	}
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
	}
	if o.NewUser.IsSet() {
		toSerialize["new_user"] = o.NewUser.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.OnlineStatus.IsSet() {
		toSerialize["online_status"] = o.OnlineStatus.Get()
	}
	if o.PhoneOn.IsSet() {
		toSerialize["phone_on"] = o.PhoneOn.Get()
	}
	if o.PostsCount.IsSet() {
		toSerialize["posts_count"] = o.PostsCount.Get()
	}
	if o.Prefecture.IsSet() {
		toSerialize["prefecture"] = o.Prefecture.Get()
	}
	if o.ProfileIcon.IsSet() {
		toSerialize["profile_icon"] = o.ProfileIcon.Get()
	}
	if o.ProfileIconThumbnail.IsSet() {
		toSerialize["profile_icon_thumbnail"] = o.ProfileIconThumbnail.Get()
	}
	if o.PushNotification.IsSet() {
		toSerialize["push_notification"] = o.PushNotification.Get()
	}
	if o.RecentlyKenta.IsSet() {
		toSerialize["recently_kenta"] = o.RecentlyKenta.Get()
	}
	if o.RestrictedReviewBy.IsSet() {
		toSerialize["restricted_review_by"] = o.RestrictedReviewBy.Get()
	}
	if o.ReviewsCount.IsSet() {
		toSerialize["reviews_count"] = o.ReviewsCount.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TotalGiftsReceivedCount.IsSet() {
		toSerialize["total_gifts_received_count"] = o.TotalGiftsReceivedCount.Get()
	}
	if o.TwitterId.IsSet() {
		toSerialize["twitter_id"] = o.TwitterId.Get()
	}
	if o.TwoFaEnabled.IsSet() {
		toSerialize["two_fa_enabled"] = o.TwoFaEnabled.Get()
	}
	if o.UpdatedTimeMillis.IsSet() {
		toSerialize["updated_time_millis"] = o.UpdatedTimeMillis.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.VideoOn.IsSet() {
		toSerialize["video_on"] = o.VideoOn.Get()
	}
	if o.Vip.IsSet() {
		toSerialize["vip"] = o.Vip.Get()
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

func (o *RealmUser) UnmarshalJSON(data []byte) (err error) {
	varRealmUser := _RealmUser{}

	err = json.Unmarshal(data, &varRealmUser)

	if err != nil {
		return err
	}

	*o = RealmUser(varRealmUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "age_verified")
		delete(additionalProperties, "apple_connected")
		delete(additionalProperties, "biography")
		delete(additionalProperties, "birth_date")
		delete(additionalProperties, "blocking_limit")
		delete(additionalProperties, "chat_request")
		delete(additionalProperties, "connected_by")
		delete(additionalProperties, "contact_phones")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "cover_image_thumbnail")
		delete(additionalProperties, "created_at_seconds")
		delete(additionalProperties, "dangerous_user")
		delete(additionalProperties, "email_confirmed")
		delete(additionalProperties, "facebook_connected")
		delete(additionalProperties, "follow_pending")
		delete(additionalProperties, "followed_by")
		delete(additionalProperties, "followed_by_pending")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "following")
		delete(additionalProperties, "followings_count")
		delete(additionalProperties, "frame")
		delete(additionalProperties, "frame_thumbnail")
		delete(additionalProperties, "from_different_generation_and_trusted")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "generation")
		delete(additionalProperties, "gifting_ability")
		delete(additionalProperties, "google_connected")
		delete(additionalProperties, "group_phone_on")
		delete(additionalProperties, "group_user")
		delete(additionalProperties, "group_video_on")
		delete(additionalProperties, "groups_users_count")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "hide_vip")
		delete(additionalProperties, "id")
		delete(additionalProperties, "interests_count")
		delete(additionalProperties, "interests_selected")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "last_logged_in_at_seconds")
		delete(additionalProperties, "line_connected")
		delete(additionalProperties, "login_streak_count")
		delete(additionalProperties, "masked_email")
		delete(additionalProperties, "media_count")
		delete(additionalProperties, "mobile_number")
		delete(additionalProperties, "new_user")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "online_status")
		delete(additionalProperties, "phone_on")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "prefecture")
		delete(additionalProperties, "profile_icon")
		delete(additionalProperties, "profile_icon_thumbnail")
		delete(additionalProperties, "push_notification")
		delete(additionalProperties, "recently_kenta")
		delete(additionalProperties, "restricted_review_by")
		delete(additionalProperties, "reviews_count")
		delete(additionalProperties, "title")
		delete(additionalProperties, "total_gifts_received_count")
		delete(additionalProperties, "twitter_id")
		delete(additionalProperties, "two_fa_enabled")
		delete(additionalProperties, "updated_time_millis")
		delete(additionalProperties, "username")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "video_on")
		delete(additionalProperties, "vip")
		delete(additionalProperties, "vip_until_seconds")
		delete(additionalProperties, "world_id_connected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmUser struct {
	value *RealmUser
	isSet bool
}

func (v NullableRealmUser) Get() *RealmUser {
	return v.value
}

func (v *NullableRealmUser) Set(val *RealmUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmUser(val *RealmUser) *NullableRealmUser {
	return &NullableRealmUser{value: val, isSet: true}
}

func (v NullableRealmUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


