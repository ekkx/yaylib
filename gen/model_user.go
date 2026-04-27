
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Biography NullableString `json:"biography,omitempty"`
	BirthDate NullableString `json:"birth_date,omitempty"`
	BlockingLimit NullableInt32 `json:"blocking_limit,omitempty"`
	ConnectedBy map[string]interface{} `json:"connected_by,omitempty"`
	ContactPhones []string `json:"contact_phones,omitempty"`
	Country map[string]interface{} `json:"country,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	CoverImageThumbnail NullableString `json:"cover_image_thumbnail,omitempty"`
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	FollowingsCount NullableInt32 `json:"followings_count,omitempty"`
	Gender NullableGender `json:"gender,omitempty"`
	GiftReceivedCounter NullableInt32 `json:"gift_received_counter,omitempty"`
	GiftingAbility NullableGiftingAbility `json:"gifting_ability,omitempty"`
	GroupsUsersCount NullableInt32 `json:"groups_users_count,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InterestsCount NullableInt32 `json:"interests_count,omitempty"`
	IsAgeVerified NullableBool `json:"is_age_verified,omitempty"`
	IsAppleConnected NullableBool `json:"is_apple_connected,omitempty"`
	IsChatRequestOn NullableBool `json:"is_chat_request_on,omitempty"`
	IsDangerous NullableBool `json:"is_dangerous,omitempty"`
	IsEmailConfirmed NullableBool `json:"is_email_confirmed,omitempty"`
	IsFacebookConnected NullableBool `json:"is_facebook_connected,omitempty"`
	IsFollowPending NullableBool `json:"is_follow_pending,omitempty"`
	IsFollowedBy NullableBool `json:"is_followed_by,omitempty"`
	IsFollowedByPending NullableBool `json:"is_followed_by_pending,omitempty"`
	IsFollowing NullableBool `json:"is_following,omitempty"`
	IsFromDifferentGenerationAndTrusted NullableBool `json:"is_from_different_generation_and_trusted,omitempty"`
	IsGoogleConnected NullableBool `json:"is_google_connected,omitempty"`
	IsGroupPhoneOn NullableBool `json:"is_group_phone_on,omitempty"`
	IsGroupVideoOn NullableBool `json:"is_group_video_on,omitempty"`
	IsHidden NullableBool `json:"is_hidden,omitempty"`
	IsHideVip NullableBool `json:"is_hide_vip,omitempty"`
	IsInterestsSelected NullableBool `json:"is_interests_selected,omitempty"`
	IsLineConnected NullableBool `json:"is_line_connected,omitempty"`
	IsMuted NullableBool `json:"is_muted,omitempty"`
	IsNew NullableBool `json:"is_new,omitempty"`
	IsPhoneOn NullableBool `json:"is_phone_on,omitempty"`
	IsPrivate NullableBool `json:"is_private,omitempty"`
	IsRecentlyPenalized NullableBool `json:"is_recently_penalized,omitempty"`
	IsRegistered NullableBool `json:"is_registered,omitempty"`
	IsTwoFaEnabled NullableBool `json:"is_two_fa_enabled,omitempty"`
	IsVideoOn NullableBool `json:"is_video_on,omitempty"`
	IsVip NullableBool `json:"is_vip,omitempty"`
	IsWorldIdConnected NullableBool `json:"is_world_id_connected,omitempty"`
	LastLoggedInAtMillis NullableInt64 `json:"last_logged_in_at_millis,omitempty"`
	LoginStreakCount NullableInt32 `json:"login_streak_count,omitempty"`
	MaskedEmail NullableString `json:"masked_email,omitempty"`
	MatchingId NullableInt64 `json:"matching_id,omitempty"`
	MediaCount NullableInt32 `json:"media_count,omitempty"`
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	OnlineStatus NullableOnlineStatus `json:"online_status,omitempty"`
	PostsCount NullableInt32 `json:"posts_count,omitempty"`
	Prefecture NullableString `json:"prefecture,omitempty"`
	ProfileIcon NullableString `json:"profile_icon,omitempty"`
	ProfileIconFrame NullableString `json:"profile_icon_frame,omitempty"`
	ProfileIconFrameThumbnail NullableString `json:"profile_icon_frame_thumbnail,omitempty"`
	ProfileIconThumbnail NullableString `json:"profile_icon_thumbnail,omitempty"`
	ReviewRestriction NullableReviewRestriction `json:"review_restriction,omitempty"`
	ReviewsCount NullableInt32 `json:"reviews_count,omitempty"`
	Title NullableTitle `json:"title,omitempty"`
	TwitterId NullableString `json:"twitter_id,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	VipUntilMillis NullableInt64 `json:"vip_until_millis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *User) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *User) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *User) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *User) UnsetBiography() {
	o.Biography.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *User) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *User) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *User) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *User) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBlockingLimit returns the BlockingLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetBlockingLimit() int32 {
	if o == nil || IsNil(o.BlockingLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.BlockingLimit.Get()
}

// GetBlockingLimitOk returns a tuple with the BlockingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetBlockingLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockingLimit.Get(), o.BlockingLimit.IsSet()
}

// HasBlockingLimit returns a boolean if a field has been set.
func (o *User) HasBlockingLimit() bool {
	if o != nil && o.BlockingLimit.IsSet() {
		return true
	}

	return false
}

// SetBlockingLimit gets a reference to the given NullableInt32 and assigns it to the BlockingLimit field.
func (o *User) SetBlockingLimit(v int32) {
	o.BlockingLimit.Set(&v)
}
// SetBlockingLimitNil sets the value for BlockingLimit to be an explicit nil
func (o *User) SetBlockingLimitNil() {
	o.BlockingLimit.Set(nil)
}

// UnsetBlockingLimit ensures that no value is present for BlockingLimit, not even an explicit nil
func (o *User) UnsetBlockingLimit() {
	o.BlockingLimit.Unset()
}

// GetConnectedBy returns the ConnectedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetConnectedBy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectedBy
}

// GetConnectedByOk returns a tuple with the ConnectedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetConnectedByOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConnectedBy) {
		return map[string]interface{}{}, false
	}
	return o.ConnectedBy, true
}

// HasConnectedBy returns a boolean if a field has been set.
func (o *User) HasConnectedBy() bool {
	if o != nil && !IsNil(o.ConnectedBy) {
		return true
	}

	return false
}

// SetConnectedBy gets a reference to the given map[string]interface{} and assigns it to the ConnectedBy field.
func (o *User) SetConnectedBy(v map[string]interface{}) {
	o.ConnectedBy = v
}

// GetContactPhones returns the ContactPhones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetContactPhones() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ContactPhones
}

// GetContactPhonesOk returns a tuple with the ContactPhones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetContactPhonesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactPhones) {
		return nil, false
	}
	return o.ContactPhones, true
}

// HasContactPhones returns a boolean if a field has been set.
func (o *User) HasContactPhones() bool {
	if o != nil && !IsNil(o.ContactPhones) {
		return true
	}

	return false
}

// SetContactPhones gets a reference to the given []string and assigns it to the ContactPhones field.
func (o *User) SetContactPhones(v []string) {
	o.ContactPhones = v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCountry() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCountryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Country) {
		return map[string]interface{}{}, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *User) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given map[string]interface{} and assigns it to the Country field.
func (o *User) SetCountry(v map[string]interface{}) {
	o.Country = v
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *User) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *User) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *User) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *User) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetCoverImageThumbnail returns the CoverImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCoverImageThumbnail() string {
	if o == nil || IsNil(o.CoverImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImageThumbnail.Get()
}

// GetCoverImageThumbnailOk returns a tuple with the CoverImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCoverImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImageThumbnail.Get(), o.CoverImageThumbnail.IsSet()
}

// HasCoverImageThumbnail returns a boolean if a field has been set.
func (o *User) HasCoverImageThumbnail() bool {
	if o != nil && o.CoverImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetCoverImageThumbnail gets a reference to the given NullableString and assigns it to the CoverImageThumbnail field.
func (o *User) SetCoverImageThumbnail(v string) {
	o.CoverImageThumbnail.Set(&v)
}
// SetCoverImageThumbnailNil sets the value for CoverImageThumbnail to be an explicit nil
func (o *User) SetCoverImageThumbnailNil() {
	o.CoverImageThumbnail.Set(nil)
}

// UnsetCoverImageThumbnail ensures that no value is present for CoverImageThumbnail, not even an explicit nil
func (o *User) UnsetCoverImageThumbnail() {
	o.CoverImageThumbnail.Unset()
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *User) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *User) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *User) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *User) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *User) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *User) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *User) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *User) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFollowingsCount returns the FollowingsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFollowingsCount() int32 {
	if o == nil || IsNil(o.FollowingsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowingsCount.Get()
}

// GetFollowingsCountOk returns a tuple with the FollowingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFollowingsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingsCount.Get(), o.FollowingsCount.IsSet()
}

// HasFollowingsCount returns a boolean if a field has been set.
func (o *User) HasFollowingsCount() bool {
	if o != nil && o.FollowingsCount.IsSet() {
		return true
	}

	return false
}

// SetFollowingsCount gets a reference to the given NullableInt32 and assigns it to the FollowingsCount field.
func (o *User) SetFollowingsCount(v int32) {
	o.FollowingsCount.Set(&v)
}
// SetFollowingsCountNil sets the value for FollowingsCount to be an explicit nil
func (o *User) SetFollowingsCountNil() {
	o.FollowingsCount.Set(nil)
}

// UnsetFollowingsCount ensures that no value is present for FollowingsCount, not even an explicit nil
func (o *User) UnsetFollowingsCount() {
	o.FollowingsCount.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetGender() Gender {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret Gender
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetGenderOk() (*Gender, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *User) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableGender and assigns it to the Gender field.
func (o *User) SetGender(v Gender) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *User) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *User) UnsetGender() {
	o.Gender.Unset()
}

// GetGiftReceivedCounter returns the GiftReceivedCounter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetGiftReceivedCounter() int32 {
	if o == nil || IsNil(o.GiftReceivedCounter.Get()) {
		var ret int32
		return ret
	}
	return *o.GiftReceivedCounter.Get()
}

// GetGiftReceivedCounterOk returns a tuple with the GiftReceivedCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetGiftReceivedCounterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftReceivedCounter.Get(), o.GiftReceivedCounter.IsSet()
}

// HasGiftReceivedCounter returns a boolean if a field has been set.
func (o *User) HasGiftReceivedCounter() bool {
	if o != nil && o.GiftReceivedCounter.IsSet() {
		return true
	}

	return false
}

// SetGiftReceivedCounter gets a reference to the given NullableInt32 and assigns it to the GiftReceivedCounter field.
func (o *User) SetGiftReceivedCounter(v int32) {
	o.GiftReceivedCounter.Set(&v)
}
// SetGiftReceivedCounterNil sets the value for GiftReceivedCounter to be an explicit nil
func (o *User) SetGiftReceivedCounterNil() {
	o.GiftReceivedCounter.Set(nil)
}

// UnsetGiftReceivedCounter ensures that no value is present for GiftReceivedCounter, not even an explicit nil
func (o *User) UnsetGiftReceivedCounter() {
	o.GiftReceivedCounter.Unset()
}

// GetGiftingAbility returns the GiftingAbility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetGiftingAbility() GiftingAbility {
	if o == nil || IsNil(o.GiftingAbility.Get()) {
		var ret GiftingAbility
		return ret
	}
	return *o.GiftingAbility.Get()
}

// GetGiftingAbilityOk returns a tuple with the GiftingAbility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetGiftingAbilityOk() (*GiftingAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftingAbility.Get(), o.GiftingAbility.IsSet()
}

// HasGiftingAbility returns a boolean if a field has been set.
func (o *User) HasGiftingAbility() bool {
	if o != nil && o.GiftingAbility.IsSet() {
		return true
	}

	return false
}

// SetGiftingAbility gets a reference to the given NullableGiftingAbility and assigns it to the GiftingAbility field.
func (o *User) SetGiftingAbility(v GiftingAbility) {
	o.GiftingAbility.Set(&v)
}
// SetGiftingAbilityNil sets the value for GiftingAbility to be an explicit nil
func (o *User) SetGiftingAbilityNil() {
	o.GiftingAbility.Set(nil)
}

// UnsetGiftingAbility ensures that no value is present for GiftingAbility, not even an explicit nil
func (o *User) UnsetGiftingAbility() {
	o.GiftingAbility.Unset()
}

// GetGroupsUsersCount returns the GroupsUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetGroupsUsersCount() int32 {
	if o == nil || IsNil(o.GroupsUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.GroupsUsersCount.Get()
}

// GetGroupsUsersCountOk returns a tuple with the GroupsUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetGroupsUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsUsersCount.Get(), o.GroupsUsersCount.IsSet()
}

// HasGroupsUsersCount returns a boolean if a field has been set.
func (o *User) HasGroupsUsersCount() bool {
	if o != nil && o.GroupsUsersCount.IsSet() {
		return true
	}

	return false
}

// SetGroupsUsersCount gets a reference to the given NullableInt32 and assigns it to the GroupsUsersCount field.
func (o *User) SetGroupsUsersCount(v int32) {
	o.GroupsUsersCount.Set(&v)
}
// SetGroupsUsersCountNil sets the value for GroupsUsersCount to be an explicit nil
func (o *User) SetGroupsUsersCountNil() {
	o.GroupsUsersCount.Set(nil)
}

// UnsetGroupsUsersCount ensures that no value is present for GroupsUsersCount, not even an explicit nil
func (o *User) UnsetGroupsUsersCount() {
	o.GroupsUsersCount.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *User) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *User) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *User) UnsetId() {
	o.Id.Unset()
}

// GetInterestsCount returns the InterestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetInterestsCount() int32 {
	if o == nil || IsNil(o.InterestsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InterestsCount.Get()
}

// GetInterestsCountOk returns a tuple with the InterestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetInterestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterestsCount.Get(), o.InterestsCount.IsSet()
}

// HasInterestsCount returns a boolean if a field has been set.
func (o *User) HasInterestsCount() bool {
	if o != nil && o.InterestsCount.IsSet() {
		return true
	}

	return false
}

// SetInterestsCount gets a reference to the given NullableInt32 and assigns it to the InterestsCount field.
func (o *User) SetInterestsCount(v int32) {
	o.InterestsCount.Set(&v)
}
// SetInterestsCountNil sets the value for InterestsCount to be an explicit nil
func (o *User) SetInterestsCountNil() {
	o.InterestsCount.Set(nil)
}

// UnsetInterestsCount ensures that no value is present for InterestsCount, not even an explicit nil
func (o *User) UnsetInterestsCount() {
	o.InterestsCount.Unset()
}

// GetIsAgeVerified returns the IsAgeVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsAgeVerified() bool {
	if o == nil || IsNil(o.IsAgeVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAgeVerified.Get()
}

// GetIsAgeVerifiedOk returns a tuple with the IsAgeVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsAgeVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAgeVerified.Get(), o.IsAgeVerified.IsSet()
}

// HasIsAgeVerified returns a boolean if a field has been set.
func (o *User) HasIsAgeVerified() bool {
	if o != nil && o.IsAgeVerified.IsSet() {
		return true
	}

	return false
}

// SetIsAgeVerified gets a reference to the given NullableBool and assigns it to the IsAgeVerified field.
func (o *User) SetIsAgeVerified(v bool) {
	o.IsAgeVerified.Set(&v)
}
// SetIsAgeVerifiedNil sets the value for IsAgeVerified to be an explicit nil
func (o *User) SetIsAgeVerifiedNil() {
	o.IsAgeVerified.Set(nil)
}

// UnsetIsAgeVerified ensures that no value is present for IsAgeVerified, not even an explicit nil
func (o *User) UnsetIsAgeVerified() {
	o.IsAgeVerified.Unset()
}

// GetIsAppleConnected returns the IsAppleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsAppleConnected() bool {
	if o == nil || IsNil(o.IsAppleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAppleConnected.Get()
}

// GetIsAppleConnectedOk returns a tuple with the IsAppleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsAppleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAppleConnected.Get(), o.IsAppleConnected.IsSet()
}

// HasIsAppleConnected returns a boolean if a field has been set.
func (o *User) HasIsAppleConnected() bool {
	if o != nil && o.IsAppleConnected.IsSet() {
		return true
	}

	return false
}

// SetIsAppleConnected gets a reference to the given NullableBool and assigns it to the IsAppleConnected field.
func (o *User) SetIsAppleConnected(v bool) {
	o.IsAppleConnected.Set(&v)
}
// SetIsAppleConnectedNil sets the value for IsAppleConnected to be an explicit nil
func (o *User) SetIsAppleConnectedNil() {
	o.IsAppleConnected.Set(nil)
}

// UnsetIsAppleConnected ensures that no value is present for IsAppleConnected, not even an explicit nil
func (o *User) UnsetIsAppleConnected() {
	o.IsAppleConnected.Unset()
}

// GetIsChatRequestOn returns the IsChatRequestOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsChatRequestOn() bool {
	if o == nil || IsNil(o.IsChatRequestOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsChatRequestOn.Get()
}

// GetIsChatRequestOnOk returns a tuple with the IsChatRequestOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsChatRequestOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsChatRequestOn.Get(), o.IsChatRequestOn.IsSet()
}

// HasIsChatRequestOn returns a boolean if a field has been set.
func (o *User) HasIsChatRequestOn() bool {
	if o != nil && o.IsChatRequestOn.IsSet() {
		return true
	}

	return false
}

// SetIsChatRequestOn gets a reference to the given NullableBool and assigns it to the IsChatRequestOn field.
func (o *User) SetIsChatRequestOn(v bool) {
	o.IsChatRequestOn.Set(&v)
}
// SetIsChatRequestOnNil sets the value for IsChatRequestOn to be an explicit nil
func (o *User) SetIsChatRequestOnNil() {
	o.IsChatRequestOn.Set(nil)
}

// UnsetIsChatRequestOn ensures that no value is present for IsChatRequestOn, not even an explicit nil
func (o *User) UnsetIsChatRequestOn() {
	o.IsChatRequestOn.Unset()
}

// GetIsDangerous returns the IsDangerous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsDangerous() bool {
	if o == nil || IsNil(o.IsDangerous.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDangerous.Get()
}

// GetIsDangerousOk returns a tuple with the IsDangerous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsDangerousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDangerous.Get(), o.IsDangerous.IsSet()
}

// HasIsDangerous returns a boolean if a field has been set.
func (o *User) HasIsDangerous() bool {
	if o != nil && o.IsDangerous.IsSet() {
		return true
	}

	return false
}

// SetIsDangerous gets a reference to the given NullableBool and assigns it to the IsDangerous field.
func (o *User) SetIsDangerous(v bool) {
	o.IsDangerous.Set(&v)
}
// SetIsDangerousNil sets the value for IsDangerous to be an explicit nil
func (o *User) SetIsDangerousNil() {
	o.IsDangerous.Set(nil)
}

// UnsetIsDangerous ensures that no value is present for IsDangerous, not even an explicit nil
func (o *User) UnsetIsDangerous() {
	o.IsDangerous.Unset()
}

// GetIsEmailConfirmed returns the IsEmailConfirmed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsEmailConfirmed() bool {
	if o == nil || IsNil(o.IsEmailConfirmed.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEmailConfirmed.Get()
}

// GetIsEmailConfirmedOk returns a tuple with the IsEmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsEmailConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEmailConfirmed.Get(), o.IsEmailConfirmed.IsSet()
}

// HasIsEmailConfirmed returns a boolean if a field has been set.
func (o *User) HasIsEmailConfirmed() bool {
	if o != nil && o.IsEmailConfirmed.IsSet() {
		return true
	}

	return false
}

// SetIsEmailConfirmed gets a reference to the given NullableBool and assigns it to the IsEmailConfirmed field.
func (o *User) SetIsEmailConfirmed(v bool) {
	o.IsEmailConfirmed.Set(&v)
}
// SetIsEmailConfirmedNil sets the value for IsEmailConfirmed to be an explicit nil
func (o *User) SetIsEmailConfirmedNil() {
	o.IsEmailConfirmed.Set(nil)
}

// UnsetIsEmailConfirmed ensures that no value is present for IsEmailConfirmed, not even an explicit nil
func (o *User) UnsetIsEmailConfirmed() {
	o.IsEmailConfirmed.Unset()
}

// GetIsFacebookConnected returns the IsFacebookConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFacebookConnected() bool {
	if o == nil || IsNil(o.IsFacebookConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFacebookConnected.Get()
}

// GetIsFacebookConnectedOk returns a tuple with the IsFacebookConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFacebookConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFacebookConnected.Get(), o.IsFacebookConnected.IsSet()
}

// HasIsFacebookConnected returns a boolean if a field has been set.
func (o *User) HasIsFacebookConnected() bool {
	if o != nil && o.IsFacebookConnected.IsSet() {
		return true
	}

	return false
}

// SetIsFacebookConnected gets a reference to the given NullableBool and assigns it to the IsFacebookConnected field.
func (o *User) SetIsFacebookConnected(v bool) {
	o.IsFacebookConnected.Set(&v)
}
// SetIsFacebookConnectedNil sets the value for IsFacebookConnected to be an explicit nil
func (o *User) SetIsFacebookConnectedNil() {
	o.IsFacebookConnected.Set(nil)
}

// UnsetIsFacebookConnected ensures that no value is present for IsFacebookConnected, not even an explicit nil
func (o *User) UnsetIsFacebookConnected() {
	o.IsFacebookConnected.Unset()
}

// GetIsFollowPending returns the IsFollowPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFollowPending() bool {
	if o == nil || IsNil(o.IsFollowPending.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFollowPending.Get()
}

// GetIsFollowPendingOk returns a tuple with the IsFollowPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFollowPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFollowPending.Get(), o.IsFollowPending.IsSet()
}

// HasIsFollowPending returns a boolean if a field has been set.
func (o *User) HasIsFollowPending() bool {
	if o != nil && o.IsFollowPending.IsSet() {
		return true
	}

	return false
}

// SetIsFollowPending gets a reference to the given NullableBool and assigns it to the IsFollowPending field.
func (o *User) SetIsFollowPending(v bool) {
	o.IsFollowPending.Set(&v)
}
// SetIsFollowPendingNil sets the value for IsFollowPending to be an explicit nil
func (o *User) SetIsFollowPendingNil() {
	o.IsFollowPending.Set(nil)
}

// UnsetIsFollowPending ensures that no value is present for IsFollowPending, not even an explicit nil
func (o *User) UnsetIsFollowPending() {
	o.IsFollowPending.Unset()
}

// GetIsFollowedBy returns the IsFollowedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFollowedBy() bool {
	if o == nil || IsNil(o.IsFollowedBy.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFollowedBy.Get()
}

// GetIsFollowedByOk returns a tuple with the IsFollowedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFollowedByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFollowedBy.Get(), o.IsFollowedBy.IsSet()
}

// HasIsFollowedBy returns a boolean if a field has been set.
func (o *User) HasIsFollowedBy() bool {
	if o != nil && o.IsFollowedBy.IsSet() {
		return true
	}

	return false
}

// SetIsFollowedBy gets a reference to the given NullableBool and assigns it to the IsFollowedBy field.
func (o *User) SetIsFollowedBy(v bool) {
	o.IsFollowedBy.Set(&v)
}
// SetIsFollowedByNil sets the value for IsFollowedBy to be an explicit nil
func (o *User) SetIsFollowedByNil() {
	o.IsFollowedBy.Set(nil)
}

// UnsetIsFollowedBy ensures that no value is present for IsFollowedBy, not even an explicit nil
func (o *User) UnsetIsFollowedBy() {
	o.IsFollowedBy.Unset()
}

// GetIsFollowedByPending returns the IsFollowedByPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFollowedByPending() bool {
	if o == nil || IsNil(o.IsFollowedByPending.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFollowedByPending.Get()
}

// GetIsFollowedByPendingOk returns a tuple with the IsFollowedByPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFollowedByPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFollowedByPending.Get(), o.IsFollowedByPending.IsSet()
}

// HasIsFollowedByPending returns a boolean if a field has been set.
func (o *User) HasIsFollowedByPending() bool {
	if o != nil && o.IsFollowedByPending.IsSet() {
		return true
	}

	return false
}

// SetIsFollowedByPending gets a reference to the given NullableBool and assigns it to the IsFollowedByPending field.
func (o *User) SetIsFollowedByPending(v bool) {
	o.IsFollowedByPending.Set(&v)
}
// SetIsFollowedByPendingNil sets the value for IsFollowedByPending to be an explicit nil
func (o *User) SetIsFollowedByPendingNil() {
	o.IsFollowedByPending.Set(nil)
}

// UnsetIsFollowedByPending ensures that no value is present for IsFollowedByPending, not even an explicit nil
func (o *User) UnsetIsFollowedByPending() {
	o.IsFollowedByPending.Unset()
}

// GetIsFollowing returns the IsFollowing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFollowing() bool {
	if o == nil || IsNil(o.IsFollowing.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFollowing.Get()
}

// GetIsFollowingOk returns a tuple with the IsFollowing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFollowing.Get(), o.IsFollowing.IsSet()
}

// HasIsFollowing returns a boolean if a field has been set.
func (o *User) HasIsFollowing() bool {
	if o != nil && o.IsFollowing.IsSet() {
		return true
	}

	return false
}

// SetIsFollowing gets a reference to the given NullableBool and assigns it to the IsFollowing field.
func (o *User) SetIsFollowing(v bool) {
	o.IsFollowing.Set(&v)
}
// SetIsFollowingNil sets the value for IsFollowing to be an explicit nil
func (o *User) SetIsFollowingNil() {
	o.IsFollowing.Set(nil)
}

// UnsetIsFollowing ensures that no value is present for IsFollowing, not even an explicit nil
func (o *User) UnsetIsFollowing() {
	o.IsFollowing.Unset()
}

// GetIsFromDifferentGenerationAndTrusted returns the IsFromDifferentGenerationAndTrusted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsFromDifferentGenerationAndTrusted() bool {
	if o == nil || IsNil(o.IsFromDifferentGenerationAndTrusted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFromDifferentGenerationAndTrusted.Get()
}

// GetIsFromDifferentGenerationAndTrustedOk returns a tuple with the IsFromDifferentGenerationAndTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsFromDifferentGenerationAndTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFromDifferentGenerationAndTrusted.Get(), o.IsFromDifferentGenerationAndTrusted.IsSet()
}

// HasIsFromDifferentGenerationAndTrusted returns a boolean if a field has been set.
func (o *User) HasIsFromDifferentGenerationAndTrusted() bool {
	if o != nil && o.IsFromDifferentGenerationAndTrusted.IsSet() {
		return true
	}

	return false
}

// SetIsFromDifferentGenerationAndTrusted gets a reference to the given NullableBool and assigns it to the IsFromDifferentGenerationAndTrusted field.
func (o *User) SetIsFromDifferentGenerationAndTrusted(v bool) {
	o.IsFromDifferentGenerationAndTrusted.Set(&v)
}
// SetIsFromDifferentGenerationAndTrustedNil sets the value for IsFromDifferentGenerationAndTrusted to be an explicit nil
func (o *User) SetIsFromDifferentGenerationAndTrustedNil() {
	o.IsFromDifferentGenerationAndTrusted.Set(nil)
}

// UnsetIsFromDifferentGenerationAndTrusted ensures that no value is present for IsFromDifferentGenerationAndTrusted, not even an explicit nil
func (o *User) UnsetIsFromDifferentGenerationAndTrusted() {
	o.IsFromDifferentGenerationAndTrusted.Unset()
}

// GetIsGoogleConnected returns the IsGoogleConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsGoogleConnected() bool {
	if o == nil || IsNil(o.IsGoogleConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGoogleConnected.Get()
}

// GetIsGoogleConnectedOk returns a tuple with the IsGoogleConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsGoogleConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGoogleConnected.Get(), o.IsGoogleConnected.IsSet()
}

// HasIsGoogleConnected returns a boolean if a field has been set.
func (o *User) HasIsGoogleConnected() bool {
	if o != nil && o.IsGoogleConnected.IsSet() {
		return true
	}

	return false
}

// SetIsGoogleConnected gets a reference to the given NullableBool and assigns it to the IsGoogleConnected field.
func (o *User) SetIsGoogleConnected(v bool) {
	o.IsGoogleConnected.Set(&v)
}
// SetIsGoogleConnectedNil sets the value for IsGoogleConnected to be an explicit nil
func (o *User) SetIsGoogleConnectedNil() {
	o.IsGoogleConnected.Set(nil)
}

// UnsetIsGoogleConnected ensures that no value is present for IsGoogleConnected, not even an explicit nil
func (o *User) UnsetIsGoogleConnected() {
	o.IsGoogleConnected.Unset()
}

// GetIsGroupPhoneOn returns the IsGroupPhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsGroupPhoneOn() bool {
	if o == nil || IsNil(o.IsGroupPhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGroupPhoneOn.Get()
}

// GetIsGroupPhoneOnOk returns a tuple with the IsGroupPhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsGroupPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGroupPhoneOn.Get(), o.IsGroupPhoneOn.IsSet()
}

// HasIsGroupPhoneOn returns a boolean if a field has been set.
func (o *User) HasIsGroupPhoneOn() bool {
	if o != nil && o.IsGroupPhoneOn.IsSet() {
		return true
	}

	return false
}

// SetIsGroupPhoneOn gets a reference to the given NullableBool and assigns it to the IsGroupPhoneOn field.
func (o *User) SetIsGroupPhoneOn(v bool) {
	o.IsGroupPhoneOn.Set(&v)
}
// SetIsGroupPhoneOnNil sets the value for IsGroupPhoneOn to be an explicit nil
func (o *User) SetIsGroupPhoneOnNil() {
	o.IsGroupPhoneOn.Set(nil)
}

// UnsetIsGroupPhoneOn ensures that no value is present for IsGroupPhoneOn, not even an explicit nil
func (o *User) UnsetIsGroupPhoneOn() {
	o.IsGroupPhoneOn.Unset()
}

// GetIsGroupVideoOn returns the IsGroupVideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsGroupVideoOn() bool {
	if o == nil || IsNil(o.IsGroupVideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGroupVideoOn.Get()
}

// GetIsGroupVideoOnOk returns a tuple with the IsGroupVideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsGroupVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGroupVideoOn.Get(), o.IsGroupVideoOn.IsSet()
}

// HasIsGroupVideoOn returns a boolean if a field has been set.
func (o *User) HasIsGroupVideoOn() bool {
	if o != nil && o.IsGroupVideoOn.IsSet() {
		return true
	}

	return false
}

// SetIsGroupVideoOn gets a reference to the given NullableBool and assigns it to the IsGroupVideoOn field.
func (o *User) SetIsGroupVideoOn(v bool) {
	o.IsGroupVideoOn.Set(&v)
}
// SetIsGroupVideoOnNil sets the value for IsGroupVideoOn to be an explicit nil
func (o *User) SetIsGroupVideoOnNil() {
	o.IsGroupVideoOn.Set(nil)
}

// UnsetIsGroupVideoOn ensures that no value is present for IsGroupVideoOn, not even an explicit nil
func (o *User) UnsetIsGroupVideoOn() {
	o.IsGroupVideoOn.Unset()
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsHidden() bool {
	if o == nil || IsNil(o.IsHidden.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHidden.Get()
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHidden.Get(), o.IsHidden.IsSet()
}

// HasIsHidden returns a boolean if a field has been set.
func (o *User) HasIsHidden() bool {
	if o != nil && o.IsHidden.IsSet() {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given NullableBool and assigns it to the IsHidden field.
func (o *User) SetIsHidden(v bool) {
	o.IsHidden.Set(&v)
}
// SetIsHiddenNil sets the value for IsHidden to be an explicit nil
func (o *User) SetIsHiddenNil() {
	o.IsHidden.Set(nil)
}

// UnsetIsHidden ensures that no value is present for IsHidden, not even an explicit nil
func (o *User) UnsetIsHidden() {
	o.IsHidden.Unset()
}

// GetIsHideVip returns the IsHideVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsHideVip() bool {
	if o == nil || IsNil(o.IsHideVip.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHideVip.Get()
}

// GetIsHideVipOk returns a tuple with the IsHideVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsHideVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHideVip.Get(), o.IsHideVip.IsSet()
}

// HasIsHideVip returns a boolean if a field has been set.
func (o *User) HasIsHideVip() bool {
	if o != nil && o.IsHideVip.IsSet() {
		return true
	}

	return false
}

// SetIsHideVip gets a reference to the given NullableBool and assigns it to the IsHideVip field.
func (o *User) SetIsHideVip(v bool) {
	o.IsHideVip.Set(&v)
}
// SetIsHideVipNil sets the value for IsHideVip to be an explicit nil
func (o *User) SetIsHideVipNil() {
	o.IsHideVip.Set(nil)
}

// UnsetIsHideVip ensures that no value is present for IsHideVip, not even an explicit nil
func (o *User) UnsetIsHideVip() {
	o.IsHideVip.Unset()
}

// GetIsInterestsSelected returns the IsInterestsSelected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsInterestsSelected() bool {
	if o == nil || IsNil(o.IsInterestsSelected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInterestsSelected.Get()
}

// GetIsInterestsSelectedOk returns a tuple with the IsInterestsSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsInterestsSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInterestsSelected.Get(), o.IsInterestsSelected.IsSet()
}

// HasIsInterestsSelected returns a boolean if a field has been set.
func (o *User) HasIsInterestsSelected() bool {
	if o != nil && o.IsInterestsSelected.IsSet() {
		return true
	}

	return false
}

// SetIsInterestsSelected gets a reference to the given NullableBool and assigns it to the IsInterestsSelected field.
func (o *User) SetIsInterestsSelected(v bool) {
	o.IsInterestsSelected.Set(&v)
}
// SetIsInterestsSelectedNil sets the value for IsInterestsSelected to be an explicit nil
func (o *User) SetIsInterestsSelectedNil() {
	o.IsInterestsSelected.Set(nil)
}

// UnsetIsInterestsSelected ensures that no value is present for IsInterestsSelected, not even an explicit nil
func (o *User) UnsetIsInterestsSelected() {
	o.IsInterestsSelected.Unset()
}

// GetIsLineConnected returns the IsLineConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsLineConnected() bool {
	if o == nil || IsNil(o.IsLineConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLineConnected.Get()
}

// GetIsLineConnectedOk returns a tuple with the IsLineConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsLineConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLineConnected.Get(), o.IsLineConnected.IsSet()
}

// HasIsLineConnected returns a boolean if a field has been set.
func (o *User) HasIsLineConnected() bool {
	if o != nil && o.IsLineConnected.IsSet() {
		return true
	}

	return false
}

// SetIsLineConnected gets a reference to the given NullableBool and assigns it to the IsLineConnected field.
func (o *User) SetIsLineConnected(v bool) {
	o.IsLineConnected.Set(&v)
}
// SetIsLineConnectedNil sets the value for IsLineConnected to be an explicit nil
func (o *User) SetIsLineConnectedNil() {
	o.IsLineConnected.Set(nil)
}

// UnsetIsLineConnected ensures that no value is present for IsLineConnected, not even an explicit nil
func (o *User) UnsetIsLineConnected() {
	o.IsLineConnected.Unset()
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMuted.Get()
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMuted.Get(), o.IsMuted.IsSet()
}

// HasIsMuted returns a boolean if a field has been set.
func (o *User) HasIsMuted() bool {
	if o != nil && o.IsMuted.IsSet() {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given NullableBool and assigns it to the IsMuted field.
func (o *User) SetIsMuted(v bool) {
	o.IsMuted.Set(&v)
}
// SetIsMutedNil sets the value for IsMuted to be an explicit nil
func (o *User) SetIsMutedNil() {
	o.IsMuted.Set(nil)
}

// UnsetIsMuted ensures that no value is present for IsMuted, not even an explicit nil
func (o *User) UnsetIsMuted() {
	o.IsMuted.Unset()
}

// GetIsNew returns the IsNew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsNew() bool {
	if o == nil || IsNil(o.IsNew.Get()) {
		var ret bool
		return ret
	}
	return *o.IsNew.Get()
}

// GetIsNewOk returns a tuple with the IsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsNewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsNew.Get(), o.IsNew.IsSet()
}

// HasIsNew returns a boolean if a field has been set.
func (o *User) HasIsNew() bool {
	if o != nil && o.IsNew.IsSet() {
		return true
	}

	return false
}

// SetIsNew gets a reference to the given NullableBool and assigns it to the IsNew field.
func (o *User) SetIsNew(v bool) {
	o.IsNew.Set(&v)
}
// SetIsNewNil sets the value for IsNew to be an explicit nil
func (o *User) SetIsNewNil() {
	o.IsNew.Set(nil)
}

// UnsetIsNew ensures that no value is present for IsNew, not even an explicit nil
func (o *User) UnsetIsNew() {
	o.IsNew.Unset()
}

// GetIsPhoneOn returns the IsPhoneOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsPhoneOn() bool {
	if o == nil || IsNil(o.IsPhoneOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPhoneOn.Get()
}

// GetIsPhoneOnOk returns a tuple with the IsPhoneOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsPhoneOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPhoneOn.Get(), o.IsPhoneOn.IsSet()
}

// HasIsPhoneOn returns a boolean if a field has been set.
func (o *User) HasIsPhoneOn() bool {
	if o != nil && o.IsPhoneOn.IsSet() {
		return true
	}

	return false
}

// SetIsPhoneOn gets a reference to the given NullableBool and assigns it to the IsPhoneOn field.
func (o *User) SetIsPhoneOn(v bool) {
	o.IsPhoneOn.Set(&v)
}
// SetIsPhoneOnNil sets the value for IsPhoneOn to be an explicit nil
func (o *User) SetIsPhoneOnNil() {
	o.IsPhoneOn.Set(nil)
}

// UnsetIsPhoneOn ensures that no value is present for IsPhoneOn, not even an explicit nil
func (o *User) UnsetIsPhoneOn() {
	o.IsPhoneOn.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *User) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *User) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *User) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *User) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetIsRecentlyPenalized returns the IsRecentlyPenalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsRecentlyPenalized() bool {
	if o == nil || IsNil(o.IsRecentlyPenalized.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRecentlyPenalized.Get()
}

// GetIsRecentlyPenalizedOk returns a tuple with the IsRecentlyPenalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsRecentlyPenalizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRecentlyPenalized.Get(), o.IsRecentlyPenalized.IsSet()
}

// HasIsRecentlyPenalized returns a boolean if a field has been set.
func (o *User) HasIsRecentlyPenalized() bool {
	if o != nil && o.IsRecentlyPenalized.IsSet() {
		return true
	}

	return false
}

// SetIsRecentlyPenalized gets a reference to the given NullableBool and assigns it to the IsRecentlyPenalized field.
func (o *User) SetIsRecentlyPenalized(v bool) {
	o.IsRecentlyPenalized.Set(&v)
}
// SetIsRecentlyPenalizedNil sets the value for IsRecentlyPenalized to be an explicit nil
func (o *User) SetIsRecentlyPenalizedNil() {
	o.IsRecentlyPenalized.Set(nil)
}

// UnsetIsRecentlyPenalized ensures that no value is present for IsRecentlyPenalized, not even an explicit nil
func (o *User) UnsetIsRecentlyPenalized() {
	o.IsRecentlyPenalized.Unset()
}

// GetIsRegistered returns the IsRegistered field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsRegistered() bool {
	if o == nil || IsNil(o.IsRegistered.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRegistered.Get()
}

// GetIsRegisteredOk returns a tuple with the IsRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRegistered.Get(), o.IsRegistered.IsSet()
}

// HasIsRegistered returns a boolean if a field has been set.
func (o *User) HasIsRegistered() bool {
	if o != nil && o.IsRegistered.IsSet() {
		return true
	}

	return false
}

// SetIsRegistered gets a reference to the given NullableBool and assigns it to the IsRegistered field.
func (o *User) SetIsRegistered(v bool) {
	o.IsRegistered.Set(&v)
}
// SetIsRegisteredNil sets the value for IsRegistered to be an explicit nil
func (o *User) SetIsRegisteredNil() {
	o.IsRegistered.Set(nil)
}

// UnsetIsRegistered ensures that no value is present for IsRegistered, not even an explicit nil
func (o *User) UnsetIsRegistered() {
	o.IsRegistered.Unset()
}

// GetIsTwoFaEnabled returns the IsTwoFaEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsTwoFaEnabled() bool {
	if o == nil || IsNil(o.IsTwoFaEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsTwoFaEnabled.Get()
}

// GetIsTwoFaEnabledOk returns a tuple with the IsTwoFaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsTwoFaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTwoFaEnabled.Get(), o.IsTwoFaEnabled.IsSet()
}

// HasIsTwoFaEnabled returns a boolean if a field has been set.
func (o *User) HasIsTwoFaEnabled() bool {
	if o != nil && o.IsTwoFaEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsTwoFaEnabled gets a reference to the given NullableBool and assigns it to the IsTwoFaEnabled field.
func (o *User) SetIsTwoFaEnabled(v bool) {
	o.IsTwoFaEnabled.Set(&v)
}
// SetIsTwoFaEnabledNil sets the value for IsTwoFaEnabled to be an explicit nil
func (o *User) SetIsTwoFaEnabledNil() {
	o.IsTwoFaEnabled.Set(nil)
}

// UnsetIsTwoFaEnabled ensures that no value is present for IsTwoFaEnabled, not even an explicit nil
func (o *User) UnsetIsTwoFaEnabled() {
	o.IsTwoFaEnabled.Unset()
}

// GetIsVideoOn returns the IsVideoOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsVideoOn() bool {
	if o == nil || IsNil(o.IsVideoOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsVideoOn.Get()
}

// GetIsVideoOnOk returns a tuple with the IsVideoOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsVideoOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsVideoOn.Get(), o.IsVideoOn.IsSet()
}

// HasIsVideoOn returns a boolean if a field has been set.
func (o *User) HasIsVideoOn() bool {
	if o != nil && o.IsVideoOn.IsSet() {
		return true
	}

	return false
}

// SetIsVideoOn gets a reference to the given NullableBool and assigns it to the IsVideoOn field.
func (o *User) SetIsVideoOn(v bool) {
	o.IsVideoOn.Set(&v)
}
// SetIsVideoOnNil sets the value for IsVideoOn to be an explicit nil
func (o *User) SetIsVideoOnNil() {
	o.IsVideoOn.Set(nil)
}

// UnsetIsVideoOn ensures that no value is present for IsVideoOn, not even an explicit nil
func (o *User) UnsetIsVideoOn() {
	o.IsVideoOn.Unset()
}

// GetIsVip returns the IsVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsVip() bool {
	if o == nil || IsNil(o.IsVip.Get()) {
		var ret bool
		return ret
	}
	return *o.IsVip.Get()
}

// GetIsVipOk returns a tuple with the IsVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsVip.Get(), o.IsVip.IsSet()
}

// HasIsVip returns a boolean if a field has been set.
func (o *User) HasIsVip() bool {
	if o != nil && o.IsVip.IsSet() {
		return true
	}

	return false
}

// SetIsVip gets a reference to the given NullableBool and assigns it to the IsVip field.
func (o *User) SetIsVip(v bool) {
	o.IsVip.Set(&v)
}
// SetIsVipNil sets the value for IsVip to be an explicit nil
func (o *User) SetIsVipNil() {
	o.IsVip.Set(nil)
}

// UnsetIsVip ensures that no value is present for IsVip, not even an explicit nil
func (o *User) UnsetIsVip() {
	o.IsVip.Unset()
}

// GetIsWorldIdConnected returns the IsWorldIdConnected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIsWorldIdConnected() bool {
	if o == nil || IsNil(o.IsWorldIdConnected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWorldIdConnected.Get()
}

// GetIsWorldIdConnectedOk returns a tuple with the IsWorldIdConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIsWorldIdConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWorldIdConnected.Get(), o.IsWorldIdConnected.IsSet()
}

// HasIsWorldIdConnected returns a boolean if a field has been set.
func (o *User) HasIsWorldIdConnected() bool {
	if o != nil && o.IsWorldIdConnected.IsSet() {
		return true
	}

	return false
}

// SetIsWorldIdConnected gets a reference to the given NullableBool and assigns it to the IsWorldIdConnected field.
func (o *User) SetIsWorldIdConnected(v bool) {
	o.IsWorldIdConnected.Set(&v)
}
// SetIsWorldIdConnectedNil sets the value for IsWorldIdConnected to be an explicit nil
func (o *User) SetIsWorldIdConnectedNil() {
	o.IsWorldIdConnected.Set(nil)
}

// UnsetIsWorldIdConnected ensures that no value is present for IsWorldIdConnected, not even an explicit nil
func (o *User) UnsetIsWorldIdConnected() {
	o.IsWorldIdConnected.Unset()
}

// GetLastLoggedInAtMillis returns the LastLoggedInAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLastLoggedInAtMillis() int64 {
	if o == nil || IsNil(o.LastLoggedInAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.LastLoggedInAtMillis.Get()
}

// GetLastLoggedInAtMillisOk returns a tuple with the LastLoggedInAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLastLoggedInAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoggedInAtMillis.Get(), o.LastLoggedInAtMillis.IsSet()
}

// HasLastLoggedInAtMillis returns a boolean if a field has been set.
func (o *User) HasLastLoggedInAtMillis() bool {
	if o != nil && o.LastLoggedInAtMillis.IsSet() {
		return true
	}

	return false
}

// SetLastLoggedInAtMillis gets a reference to the given NullableInt64 and assigns it to the LastLoggedInAtMillis field.
func (o *User) SetLastLoggedInAtMillis(v int64) {
	o.LastLoggedInAtMillis.Set(&v)
}
// SetLastLoggedInAtMillisNil sets the value for LastLoggedInAtMillis to be an explicit nil
func (o *User) SetLastLoggedInAtMillisNil() {
	o.LastLoggedInAtMillis.Set(nil)
}

// UnsetLastLoggedInAtMillis ensures that no value is present for LastLoggedInAtMillis, not even an explicit nil
func (o *User) UnsetLastLoggedInAtMillis() {
	o.LastLoggedInAtMillis.Unset()
}

// GetLoginStreakCount returns the LoginStreakCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLoginStreakCount() int32 {
	if o == nil || IsNil(o.LoginStreakCount.Get()) {
		var ret int32
		return ret
	}
	return *o.LoginStreakCount.Get()
}

// GetLoginStreakCountOk returns a tuple with the LoginStreakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLoginStreakCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginStreakCount.Get(), o.LoginStreakCount.IsSet()
}

// HasLoginStreakCount returns a boolean if a field has been set.
func (o *User) HasLoginStreakCount() bool {
	if o != nil && o.LoginStreakCount.IsSet() {
		return true
	}

	return false
}

// SetLoginStreakCount gets a reference to the given NullableInt32 and assigns it to the LoginStreakCount field.
func (o *User) SetLoginStreakCount(v int32) {
	o.LoginStreakCount.Set(&v)
}
// SetLoginStreakCountNil sets the value for LoginStreakCount to be an explicit nil
func (o *User) SetLoginStreakCountNil() {
	o.LoginStreakCount.Set(nil)
}

// UnsetLoginStreakCount ensures that no value is present for LoginStreakCount, not even an explicit nil
func (o *User) UnsetLoginStreakCount() {
	o.LoginStreakCount.Unset()
}

// GetMaskedEmail returns the MaskedEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetMaskedEmail() string {
	if o == nil || IsNil(o.MaskedEmail.Get()) {
		var ret string
		return ret
	}
	return *o.MaskedEmail.Get()
}

// GetMaskedEmailOk returns a tuple with the MaskedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMaskedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedEmail.Get(), o.MaskedEmail.IsSet()
}

// HasMaskedEmail returns a boolean if a field has been set.
func (o *User) HasMaskedEmail() bool {
	if o != nil && o.MaskedEmail.IsSet() {
		return true
	}

	return false
}

// SetMaskedEmail gets a reference to the given NullableString and assigns it to the MaskedEmail field.
func (o *User) SetMaskedEmail(v string) {
	o.MaskedEmail.Set(&v)
}
// SetMaskedEmailNil sets the value for MaskedEmail to be an explicit nil
func (o *User) SetMaskedEmailNil() {
	o.MaskedEmail.Set(nil)
}

// UnsetMaskedEmail ensures that no value is present for MaskedEmail, not even an explicit nil
func (o *User) UnsetMaskedEmail() {
	o.MaskedEmail.Unset()
}

// GetMatchingId returns the MatchingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetMatchingId() int64 {
	if o == nil || IsNil(o.MatchingId.Get()) {
		var ret int64
		return ret
	}
	return *o.MatchingId.Get()
}

// GetMatchingIdOk returns a tuple with the MatchingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMatchingIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchingId.Get(), o.MatchingId.IsSet()
}

// HasMatchingId returns a boolean if a field has been set.
func (o *User) HasMatchingId() bool {
	if o != nil && o.MatchingId.IsSet() {
		return true
	}

	return false
}

// SetMatchingId gets a reference to the given NullableInt64 and assigns it to the MatchingId field.
func (o *User) SetMatchingId(v int64) {
	o.MatchingId.Set(&v)
}
// SetMatchingIdNil sets the value for MatchingId to be an explicit nil
func (o *User) SetMatchingIdNil() {
	o.MatchingId.Set(nil)
}

// UnsetMatchingId ensures that no value is present for MatchingId, not even an explicit nil
func (o *User) UnsetMatchingId() {
	o.MatchingId.Unset()
}

// GetMediaCount returns the MediaCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetMediaCount() int32 {
	if o == nil || IsNil(o.MediaCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MediaCount.Get()
}

// GetMediaCountOk returns a tuple with the MediaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMediaCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaCount.Get(), o.MediaCount.IsSet()
}

// HasMediaCount returns a boolean if a field has been set.
func (o *User) HasMediaCount() bool {
	if o != nil && o.MediaCount.IsSet() {
		return true
	}

	return false
}

// SetMediaCount gets a reference to the given NullableInt32 and assigns it to the MediaCount field.
func (o *User) SetMediaCount(v int32) {
	o.MediaCount.Set(&v)
}
// SetMediaCountNil sets the value for MediaCount to be an explicit nil
func (o *User) SetMediaCountNil() {
	o.MediaCount.Set(nil)
}

// UnsetMediaCount ensures that no value is present for MediaCount, not even an explicit nil
func (o *User) UnsetMediaCount() {
	o.MediaCount.Unset()
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetMobileNumber() string {
	if o == nil || IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *User) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *User) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *User) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *User) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *User) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *User) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *User) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *User) UnsetNickname() {
	o.Nickname.Unset()
}

// GetOnlineStatus returns the OnlineStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetOnlineStatus() OnlineStatus {
	if o == nil || IsNil(o.OnlineStatus.Get()) {
		var ret OnlineStatus
		return ret
	}
	return *o.OnlineStatus.Get()
}

// GetOnlineStatusOk returns a tuple with the OnlineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetOnlineStatusOk() (*OnlineStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineStatus.Get(), o.OnlineStatus.IsSet()
}

// HasOnlineStatus returns a boolean if a field has been set.
func (o *User) HasOnlineStatus() bool {
	if o != nil && o.OnlineStatus.IsSet() {
		return true
	}

	return false
}

// SetOnlineStatus gets a reference to the given NullableOnlineStatus and assigns it to the OnlineStatus field.
func (o *User) SetOnlineStatus(v OnlineStatus) {
	o.OnlineStatus.Set(&v)
}
// SetOnlineStatusNil sets the value for OnlineStatus to be an explicit nil
func (o *User) SetOnlineStatusNil() {
	o.OnlineStatus.Set(nil)
}

// UnsetOnlineStatus ensures that no value is present for OnlineStatus, not even an explicit nil
func (o *User) UnsetOnlineStatus() {
	o.OnlineStatus.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPostsCount() int32 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *User) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt32 and assigns it to the PostsCount field.
func (o *User) SetPostsCount(v int32) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *User) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *User) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetPrefecture returns the Prefecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPrefecture() string {
	if o == nil || IsNil(o.Prefecture.Get()) {
		var ret string
		return ret
	}
	return *o.Prefecture.Get()
}

// GetPrefectureOk returns a tuple with the Prefecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPrefectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefecture.Get(), o.Prefecture.IsSet()
}

// HasPrefecture returns a boolean if a field has been set.
func (o *User) HasPrefecture() bool {
	if o != nil && o.Prefecture.IsSet() {
		return true
	}

	return false
}

// SetPrefecture gets a reference to the given NullableString and assigns it to the Prefecture field.
func (o *User) SetPrefecture(v string) {
	o.Prefecture.Set(&v)
}
// SetPrefectureNil sets the value for Prefecture to be an explicit nil
func (o *User) SetPrefectureNil() {
	o.Prefecture.Set(nil)
}

// UnsetPrefecture ensures that no value is present for Prefecture, not even an explicit nil
func (o *User) UnsetPrefecture() {
	o.Prefecture.Unset()
}

// GetProfileIcon returns the ProfileIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetProfileIcon() string {
	if o == nil || IsNil(o.ProfileIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIcon.Get()
}

// GetProfileIconOk returns a tuple with the ProfileIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetProfileIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIcon.Get(), o.ProfileIcon.IsSet()
}

// HasProfileIcon returns a boolean if a field has been set.
func (o *User) HasProfileIcon() bool {
	if o != nil && o.ProfileIcon.IsSet() {
		return true
	}

	return false
}

// SetProfileIcon gets a reference to the given NullableString and assigns it to the ProfileIcon field.
func (o *User) SetProfileIcon(v string) {
	o.ProfileIcon.Set(&v)
}
// SetProfileIconNil sets the value for ProfileIcon to be an explicit nil
func (o *User) SetProfileIconNil() {
	o.ProfileIcon.Set(nil)
}

// UnsetProfileIcon ensures that no value is present for ProfileIcon, not even an explicit nil
func (o *User) UnsetProfileIcon() {
	o.ProfileIcon.Unset()
}

// GetProfileIconFrame returns the ProfileIconFrame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetProfileIconFrame() string {
	if o == nil || IsNil(o.ProfileIconFrame.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIconFrame.Get()
}

// GetProfileIconFrameOk returns a tuple with the ProfileIconFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetProfileIconFrameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIconFrame.Get(), o.ProfileIconFrame.IsSet()
}

// HasProfileIconFrame returns a boolean if a field has been set.
func (o *User) HasProfileIconFrame() bool {
	if o != nil && o.ProfileIconFrame.IsSet() {
		return true
	}

	return false
}

// SetProfileIconFrame gets a reference to the given NullableString and assigns it to the ProfileIconFrame field.
func (o *User) SetProfileIconFrame(v string) {
	o.ProfileIconFrame.Set(&v)
}
// SetProfileIconFrameNil sets the value for ProfileIconFrame to be an explicit nil
func (o *User) SetProfileIconFrameNil() {
	o.ProfileIconFrame.Set(nil)
}

// UnsetProfileIconFrame ensures that no value is present for ProfileIconFrame, not even an explicit nil
func (o *User) UnsetProfileIconFrame() {
	o.ProfileIconFrame.Unset()
}

// GetProfileIconFrameThumbnail returns the ProfileIconFrameThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetProfileIconFrameThumbnail() string {
	if o == nil || IsNil(o.ProfileIconFrameThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIconFrameThumbnail.Get()
}

// GetProfileIconFrameThumbnailOk returns a tuple with the ProfileIconFrameThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetProfileIconFrameThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIconFrameThumbnail.Get(), o.ProfileIconFrameThumbnail.IsSet()
}

// HasProfileIconFrameThumbnail returns a boolean if a field has been set.
func (o *User) HasProfileIconFrameThumbnail() bool {
	if o != nil && o.ProfileIconFrameThumbnail.IsSet() {
		return true
	}

	return false
}

// SetProfileIconFrameThumbnail gets a reference to the given NullableString and assigns it to the ProfileIconFrameThumbnail field.
func (o *User) SetProfileIconFrameThumbnail(v string) {
	o.ProfileIconFrameThumbnail.Set(&v)
}
// SetProfileIconFrameThumbnailNil sets the value for ProfileIconFrameThumbnail to be an explicit nil
func (o *User) SetProfileIconFrameThumbnailNil() {
	o.ProfileIconFrameThumbnail.Set(nil)
}

// UnsetProfileIconFrameThumbnail ensures that no value is present for ProfileIconFrameThumbnail, not even an explicit nil
func (o *User) UnsetProfileIconFrameThumbnail() {
	o.ProfileIconFrameThumbnail.Unset()
}

// GetProfileIconThumbnail returns the ProfileIconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetProfileIconThumbnail() string {
	if o == nil || IsNil(o.ProfileIconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIconThumbnail.Get()
}

// GetProfileIconThumbnailOk returns a tuple with the ProfileIconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetProfileIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIconThumbnail.Get(), o.ProfileIconThumbnail.IsSet()
}

// HasProfileIconThumbnail returns a boolean if a field has been set.
func (o *User) HasProfileIconThumbnail() bool {
	if o != nil && o.ProfileIconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetProfileIconThumbnail gets a reference to the given NullableString and assigns it to the ProfileIconThumbnail field.
func (o *User) SetProfileIconThumbnail(v string) {
	o.ProfileIconThumbnail.Set(&v)
}
// SetProfileIconThumbnailNil sets the value for ProfileIconThumbnail to be an explicit nil
func (o *User) SetProfileIconThumbnailNil() {
	o.ProfileIconThumbnail.Set(nil)
}

// UnsetProfileIconThumbnail ensures that no value is present for ProfileIconThumbnail, not even an explicit nil
func (o *User) UnsetProfileIconThumbnail() {
	o.ProfileIconThumbnail.Unset()
}

// GetReviewRestriction returns the ReviewRestriction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetReviewRestriction() ReviewRestriction {
	if o == nil || IsNil(o.ReviewRestriction.Get()) {
		var ret ReviewRestriction
		return ret
	}
	return *o.ReviewRestriction.Get()
}

// GetReviewRestrictionOk returns a tuple with the ReviewRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetReviewRestrictionOk() (*ReviewRestriction, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewRestriction.Get(), o.ReviewRestriction.IsSet()
}

// HasReviewRestriction returns a boolean if a field has been set.
func (o *User) HasReviewRestriction() bool {
	if o != nil && o.ReviewRestriction.IsSet() {
		return true
	}

	return false
}

// SetReviewRestriction gets a reference to the given NullableReviewRestriction and assigns it to the ReviewRestriction field.
func (o *User) SetReviewRestriction(v ReviewRestriction) {
	o.ReviewRestriction.Set(&v)
}
// SetReviewRestrictionNil sets the value for ReviewRestriction to be an explicit nil
func (o *User) SetReviewRestrictionNil() {
	o.ReviewRestriction.Set(nil)
}

// UnsetReviewRestriction ensures that no value is present for ReviewRestriction, not even an explicit nil
func (o *User) UnsetReviewRestriction() {
	o.ReviewRestriction.Unset()
}

// GetReviewsCount returns the ReviewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetReviewsCount() int32 {
	if o == nil || IsNil(o.ReviewsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReviewsCount.Get()
}

// GetReviewsCountOk returns a tuple with the ReviewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetReviewsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewsCount.Get(), o.ReviewsCount.IsSet()
}

// HasReviewsCount returns a boolean if a field has been set.
func (o *User) HasReviewsCount() bool {
	if o != nil && o.ReviewsCount.IsSet() {
		return true
	}

	return false
}

// SetReviewsCount gets a reference to the given NullableInt32 and assigns it to the ReviewsCount field.
func (o *User) SetReviewsCount(v int32) {
	o.ReviewsCount.Set(&v)
}
// SetReviewsCountNil sets the value for ReviewsCount to be an explicit nil
func (o *User) SetReviewsCountNil() {
	o.ReviewsCount.Set(nil)
}

// UnsetReviewsCount ensures that no value is present for ReviewsCount, not even an explicit nil
func (o *User) UnsetReviewsCount() {
	o.ReviewsCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetTitle() Title {
	if o == nil || IsNil(o.Title.Get()) {
		var ret Title
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTitleOk() (*Title, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *User) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableTitle and assigns it to the Title field.
func (o *User) SetTitle(v Title) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *User) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *User) UnsetTitle() {
	o.Title.Unset()
}

// GetTwitterId returns the TwitterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetTwitterId() string {
	if o == nil || IsNil(o.TwitterId.Get()) {
		var ret string
		return ret
	}
	return *o.TwitterId.Get()
}

// GetTwitterIdOk returns a tuple with the TwitterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTwitterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterId.Get(), o.TwitterId.IsSet()
}

// HasTwitterId returns a boolean if a field has been set.
func (o *User) HasTwitterId() bool {
	if o != nil && o.TwitterId.IsSet() {
		return true
	}

	return false
}

// SetTwitterId gets a reference to the given NullableString and assigns it to the TwitterId field.
func (o *User) SetTwitterId(v string) {
	o.TwitterId.Set(&v)
}
// SetTwitterIdNil sets the value for TwitterId to be an explicit nil
func (o *User) SetTwitterIdNil() {
	o.TwitterId.Set(nil)
}

// UnsetTwitterId ensures that no value is present for TwitterId, not even an explicit nil
func (o *User) UnsetTwitterId() {
	o.TwitterId.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *User) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *User) UnsetUsername() {
	o.Username.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *User) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *User) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *User) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *User) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVipUntilMillis returns the VipUntilMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetVipUntilMillis() int64 {
	if o == nil || IsNil(o.VipUntilMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.VipUntilMillis.Get()
}

// GetVipUntilMillisOk returns a tuple with the VipUntilMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetVipUntilMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipUntilMillis.Get(), o.VipUntilMillis.IsSet()
}

// HasVipUntilMillis returns a boolean if a field has been set.
func (o *User) HasVipUntilMillis() bool {
	if o != nil && o.VipUntilMillis.IsSet() {
		return true
	}

	return false
}

// SetVipUntilMillis gets a reference to the given NullableInt64 and assigns it to the VipUntilMillis field.
func (o *User) SetVipUntilMillis(v int64) {
	o.VipUntilMillis.Set(&v)
}
// SetVipUntilMillisNil sets the value for VipUntilMillis to be an explicit nil
func (o *User) SetVipUntilMillisNil() {
	o.VipUntilMillis.Set(nil)
}

// UnsetVipUntilMillis ensures that no value is present for VipUntilMillis, not even an explicit nil
func (o *User) UnsetVipUntilMillis() {
	o.VipUntilMillis.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Biography.IsSet() {
		toSerialize["biography"] = o.Biography.Get()
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if o.BlockingLimit.IsSet() {
		toSerialize["blocking_limit"] = o.BlockingLimit.Get()
	}
	if o.ConnectedBy != nil {
		toSerialize["connected_by"] = o.ConnectedBy
	}
	if o.ContactPhones != nil {
		toSerialize["contact_phones"] = o.ContactPhones
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.CoverImageThumbnail.IsSet() {
		toSerialize["cover_image_thumbnail"] = o.CoverImageThumbnail.Get()
	}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.FollowersCount.IsSet() {
		toSerialize["followers_count"] = o.FollowersCount.Get()
	}
	if o.FollowingsCount.IsSet() {
		toSerialize["followings_count"] = o.FollowingsCount.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.GiftReceivedCounter.IsSet() {
		toSerialize["gift_received_counter"] = o.GiftReceivedCounter.Get()
	}
	if o.GiftingAbility.IsSet() {
		toSerialize["gifting_ability"] = o.GiftingAbility.Get()
	}
	if o.GroupsUsersCount.IsSet() {
		toSerialize["groups_users_count"] = o.GroupsUsersCount.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InterestsCount.IsSet() {
		toSerialize["interests_count"] = o.InterestsCount.Get()
	}
	if o.IsAgeVerified.IsSet() {
		toSerialize["is_age_verified"] = o.IsAgeVerified.Get()
	}
	if o.IsAppleConnected.IsSet() {
		toSerialize["is_apple_connected"] = o.IsAppleConnected.Get()
	}
	if o.IsChatRequestOn.IsSet() {
		toSerialize["is_chat_request_on"] = o.IsChatRequestOn.Get()
	}
	if o.IsDangerous.IsSet() {
		toSerialize["is_dangerous"] = o.IsDangerous.Get()
	}
	if o.IsEmailConfirmed.IsSet() {
		toSerialize["is_email_confirmed"] = o.IsEmailConfirmed.Get()
	}
	if o.IsFacebookConnected.IsSet() {
		toSerialize["is_facebook_connected"] = o.IsFacebookConnected.Get()
	}
	if o.IsFollowPending.IsSet() {
		toSerialize["is_follow_pending"] = o.IsFollowPending.Get()
	}
	if o.IsFollowedBy.IsSet() {
		toSerialize["is_followed_by"] = o.IsFollowedBy.Get()
	}
	if o.IsFollowedByPending.IsSet() {
		toSerialize["is_followed_by_pending"] = o.IsFollowedByPending.Get()
	}
	if o.IsFollowing.IsSet() {
		toSerialize["is_following"] = o.IsFollowing.Get()
	}
	if o.IsFromDifferentGenerationAndTrusted.IsSet() {
		toSerialize["is_from_different_generation_and_trusted"] = o.IsFromDifferentGenerationAndTrusted.Get()
	}
	if o.IsGoogleConnected.IsSet() {
		toSerialize["is_google_connected"] = o.IsGoogleConnected.Get()
	}
	if o.IsGroupPhoneOn.IsSet() {
		toSerialize["is_group_phone_on"] = o.IsGroupPhoneOn.Get()
	}
	if o.IsGroupVideoOn.IsSet() {
		toSerialize["is_group_video_on"] = o.IsGroupVideoOn.Get()
	}
	if o.IsHidden.IsSet() {
		toSerialize["is_hidden"] = o.IsHidden.Get()
	}
	if o.IsHideVip.IsSet() {
		toSerialize["is_hide_vip"] = o.IsHideVip.Get()
	}
	if o.IsInterestsSelected.IsSet() {
		toSerialize["is_interests_selected"] = o.IsInterestsSelected.Get()
	}
	if o.IsLineConnected.IsSet() {
		toSerialize["is_line_connected"] = o.IsLineConnected.Get()
	}
	if o.IsMuted.IsSet() {
		toSerialize["is_muted"] = o.IsMuted.Get()
	}
	if o.IsNew.IsSet() {
		toSerialize["is_new"] = o.IsNew.Get()
	}
	if o.IsPhoneOn.IsSet() {
		toSerialize["is_phone_on"] = o.IsPhoneOn.Get()
	}
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.IsRecentlyPenalized.IsSet() {
		toSerialize["is_recently_penalized"] = o.IsRecentlyPenalized.Get()
	}
	if o.IsRegistered.IsSet() {
		toSerialize["is_registered"] = o.IsRegistered.Get()
	}
	if o.IsTwoFaEnabled.IsSet() {
		toSerialize["is_two_fa_enabled"] = o.IsTwoFaEnabled.Get()
	}
	if o.IsVideoOn.IsSet() {
		toSerialize["is_video_on"] = o.IsVideoOn.Get()
	}
	if o.IsVip.IsSet() {
		toSerialize["is_vip"] = o.IsVip.Get()
	}
	if o.IsWorldIdConnected.IsSet() {
		toSerialize["is_world_id_connected"] = o.IsWorldIdConnected.Get()
	}
	if o.LastLoggedInAtMillis.IsSet() {
		toSerialize["last_logged_in_at_millis"] = o.LastLoggedInAtMillis.Get()
	}
	if o.LoginStreakCount.IsSet() {
		toSerialize["login_streak_count"] = o.LoginStreakCount.Get()
	}
	if o.MaskedEmail.IsSet() {
		toSerialize["masked_email"] = o.MaskedEmail.Get()
	}
	if o.MatchingId.IsSet() {
		toSerialize["matching_id"] = o.MatchingId.Get()
	}
	if o.MediaCount.IsSet() {
		toSerialize["media_count"] = o.MediaCount.Get()
	}
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.OnlineStatus.IsSet() {
		toSerialize["online_status"] = o.OnlineStatus.Get()
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
	if o.ProfileIconFrame.IsSet() {
		toSerialize["profile_icon_frame"] = o.ProfileIconFrame.Get()
	}
	if o.ProfileIconFrameThumbnail.IsSet() {
		toSerialize["profile_icon_frame_thumbnail"] = o.ProfileIconFrameThumbnail.Get()
	}
	if o.ProfileIconThumbnail.IsSet() {
		toSerialize["profile_icon_thumbnail"] = o.ProfileIconThumbnail.Get()
	}
	if o.ReviewRestriction.IsSet() {
		toSerialize["review_restriction"] = o.ReviewRestriction.Get()
	}
	if o.ReviewsCount.IsSet() {
		toSerialize["reviews_count"] = o.ReviewsCount.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TwitterId.IsSet() {
		toSerialize["twitter_id"] = o.TwitterId.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.VipUntilMillis.IsSet() {
		toSerialize["vip_until_millis"] = o.VipUntilMillis.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	varUser := _User{}

	err = json.Unmarshal(data, &varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "biography")
		delete(additionalProperties, "birth_date")
		delete(additionalProperties, "blocking_limit")
		delete(additionalProperties, "connected_by")
		delete(additionalProperties, "contact_phones")
		delete(additionalProperties, "country")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "cover_image_thumbnail")
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "followings_count")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "gift_received_counter")
		delete(additionalProperties, "gifting_ability")
		delete(additionalProperties, "groups_users_count")
		delete(additionalProperties, "id")
		delete(additionalProperties, "interests_count")
		delete(additionalProperties, "is_age_verified")
		delete(additionalProperties, "is_apple_connected")
		delete(additionalProperties, "is_chat_request_on")
		delete(additionalProperties, "is_dangerous")
		delete(additionalProperties, "is_email_confirmed")
		delete(additionalProperties, "is_facebook_connected")
		delete(additionalProperties, "is_follow_pending")
		delete(additionalProperties, "is_followed_by")
		delete(additionalProperties, "is_followed_by_pending")
		delete(additionalProperties, "is_following")
		delete(additionalProperties, "is_from_different_generation_and_trusted")
		delete(additionalProperties, "is_google_connected")
		delete(additionalProperties, "is_group_phone_on")
		delete(additionalProperties, "is_group_video_on")
		delete(additionalProperties, "is_hidden")
		delete(additionalProperties, "is_hide_vip")
		delete(additionalProperties, "is_interests_selected")
		delete(additionalProperties, "is_line_connected")
		delete(additionalProperties, "is_muted")
		delete(additionalProperties, "is_new")
		delete(additionalProperties, "is_phone_on")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "is_recently_penalized")
		delete(additionalProperties, "is_registered")
		delete(additionalProperties, "is_two_fa_enabled")
		delete(additionalProperties, "is_video_on")
		delete(additionalProperties, "is_vip")
		delete(additionalProperties, "is_world_id_connected")
		delete(additionalProperties, "last_logged_in_at_millis")
		delete(additionalProperties, "login_streak_count")
		delete(additionalProperties, "masked_email")
		delete(additionalProperties, "matching_id")
		delete(additionalProperties, "media_count")
		delete(additionalProperties, "mobile_number")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "online_status")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "prefecture")
		delete(additionalProperties, "profile_icon")
		delete(additionalProperties, "profile_icon_frame")
		delete(additionalProperties, "profile_icon_frame_thumbnail")
		delete(additionalProperties, "profile_icon_thumbnail")
		delete(additionalProperties, "review_restriction")
		delete(additionalProperties, "reviews_count")
		delete(additionalProperties, "title")
		delete(additionalProperties, "twitter_id")
		delete(additionalProperties, "username")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "vip_until_millis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


