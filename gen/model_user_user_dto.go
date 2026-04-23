
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserUserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserDTO{}

// UserUserDTO struct for UserUserDTO
type UserUserDTO struct {
	AgeVerified NullableBool `json:"age_verified,omitempty"`
	Biography NullableString `json:"biography,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	CoverImageThumbnail NullableString `json:"cover_image_thumbnail,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	DangerousUser NullableBool `json:"dangerous_user,omitempty"`
	FollowPending NullableBool `json:"follow_pending,omitempty"`
	FollowedBy NullableBool `json:"followed_by,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	Following NullableBool `json:"following,omitempty"`
	FollowingsCount NullableInt32 `json:"followings_count,omitempty"`
	Gender NullableInt32 `json:"gender,omitempty"`
	GroupsUsersCount NullableInt32 `json:"groups_users_count,omitempty"`
	Hidden NullableBool `json:"hidden,omitempty"`
	HideVip NullableBool `json:"hide_vip,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsPrivate NullableBool `json:"is_private,omitempty"`
	NewUser NullableBool `json:"new_user,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	PostsCount NullableInt32 `json:"posts_count,omitempty"`
	Prefecture NullableString `json:"prefecture,omitempty"`
	ProfileIcon NullableString `json:"profile_icon,omitempty"`
	ProfileIconThumbnail NullableString `json:"profile_icon_thumbnail,omitempty"`
	RecentlyKenta NullableBool `json:"recently_kenta,omitempty"`
	ReviewsCount NullableInt32 `json:"reviews_count,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Vip NullableBool `json:"vip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserUserDTO UserUserDTO

// NewUserUserDTO instantiates a new UserUserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserDTO() *UserUserDTO {
	this := UserUserDTO{}
	return &this
}

// NewUserUserDTOWithDefaults instantiates a new UserUserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserDTOWithDefaults() *UserUserDTO {
	this := UserUserDTO{}
	return &this
}

// GetAgeVerified returns the AgeVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetAgeVerified() bool {
	if o == nil || IsNil(o.AgeVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.AgeVerified.Get()
}

// GetAgeVerifiedOk returns a tuple with the AgeVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetAgeVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgeVerified.Get(), o.AgeVerified.IsSet()
}

// HasAgeVerified returns a boolean if a field has been set.
func (o *UserUserDTO) HasAgeVerified() bool {
	if o != nil && o.AgeVerified.IsSet() {
		return true
	}

	return false
}

// SetAgeVerified gets a reference to the given NullableBool and assigns it to the AgeVerified field.
func (o *UserUserDTO) SetAgeVerified(v bool) {
	o.AgeVerified.Set(&v)
}
// SetAgeVerifiedNil sets the value for AgeVerified to be an explicit nil
func (o *UserUserDTO) SetAgeVerifiedNil() {
	o.AgeVerified.Set(nil)
}

// UnsetAgeVerified ensures that no value is present for AgeVerified, not even an explicit nil
func (o *UserUserDTO) UnsetAgeVerified() {
	o.AgeVerified.Unset()
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *UserUserDTO) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *UserUserDTO) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *UserUserDTO) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *UserUserDTO) UnsetBiography() {
	o.Biography.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *UserUserDTO) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *UserUserDTO) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *UserUserDTO) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *UserUserDTO) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetCoverImageThumbnail returns the CoverImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetCoverImageThumbnail() string {
	if o == nil || IsNil(o.CoverImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImageThumbnail.Get()
}

// GetCoverImageThumbnailOk returns a tuple with the CoverImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetCoverImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImageThumbnail.Get(), o.CoverImageThumbnail.IsSet()
}

// HasCoverImageThumbnail returns a boolean if a field has been set.
func (o *UserUserDTO) HasCoverImageThumbnail() bool {
	if o != nil && o.CoverImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetCoverImageThumbnail gets a reference to the given NullableString and assigns it to the CoverImageThumbnail field.
func (o *UserUserDTO) SetCoverImageThumbnail(v string) {
	o.CoverImageThumbnail.Set(&v)
}
// SetCoverImageThumbnailNil sets the value for CoverImageThumbnail to be an explicit nil
func (o *UserUserDTO) SetCoverImageThumbnailNil() {
	o.CoverImageThumbnail.Set(nil)
}

// UnsetCoverImageThumbnail ensures that no value is present for CoverImageThumbnail, not even an explicit nil
func (o *UserUserDTO) UnsetCoverImageThumbnail() {
	o.CoverImageThumbnail.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserUserDTO) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *UserUserDTO) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *UserUserDTO) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *UserUserDTO) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDangerousUser returns the DangerousUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetDangerousUser() bool {
	if o == nil || IsNil(o.DangerousUser.Get()) {
		var ret bool
		return ret
	}
	return *o.DangerousUser.Get()
}

// GetDangerousUserOk returns a tuple with the DangerousUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetDangerousUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DangerousUser.Get(), o.DangerousUser.IsSet()
}

// HasDangerousUser returns a boolean if a field has been set.
func (o *UserUserDTO) HasDangerousUser() bool {
	if o != nil && o.DangerousUser.IsSet() {
		return true
	}

	return false
}

// SetDangerousUser gets a reference to the given NullableBool and assigns it to the DangerousUser field.
func (o *UserUserDTO) SetDangerousUser(v bool) {
	o.DangerousUser.Set(&v)
}
// SetDangerousUserNil sets the value for DangerousUser to be an explicit nil
func (o *UserUserDTO) SetDangerousUserNil() {
	o.DangerousUser.Set(nil)
}

// UnsetDangerousUser ensures that no value is present for DangerousUser, not even an explicit nil
func (o *UserUserDTO) UnsetDangerousUser() {
	o.DangerousUser.Unset()
}

// GetFollowPending returns the FollowPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetFollowPending() bool {
	if o == nil || IsNil(o.FollowPending.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowPending.Get()
}

// GetFollowPendingOk returns a tuple with the FollowPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetFollowPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowPending.Get(), o.FollowPending.IsSet()
}

// HasFollowPending returns a boolean if a field has been set.
func (o *UserUserDTO) HasFollowPending() bool {
	if o != nil && o.FollowPending.IsSet() {
		return true
	}

	return false
}

// SetFollowPending gets a reference to the given NullableBool and assigns it to the FollowPending field.
func (o *UserUserDTO) SetFollowPending(v bool) {
	o.FollowPending.Set(&v)
}
// SetFollowPendingNil sets the value for FollowPending to be an explicit nil
func (o *UserUserDTO) SetFollowPendingNil() {
	o.FollowPending.Set(nil)
}

// UnsetFollowPending ensures that no value is present for FollowPending, not even an explicit nil
func (o *UserUserDTO) UnsetFollowPending() {
	o.FollowPending.Unset()
}

// GetFollowedBy returns the FollowedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetFollowedBy() bool {
	if o == nil || IsNil(o.FollowedBy.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowedBy.Get()
}

// GetFollowedByOk returns a tuple with the FollowedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetFollowedByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowedBy.Get(), o.FollowedBy.IsSet()
}

// HasFollowedBy returns a boolean if a field has been set.
func (o *UserUserDTO) HasFollowedBy() bool {
	if o != nil && o.FollowedBy.IsSet() {
		return true
	}

	return false
}

// SetFollowedBy gets a reference to the given NullableBool and assigns it to the FollowedBy field.
func (o *UserUserDTO) SetFollowedBy(v bool) {
	o.FollowedBy.Set(&v)
}
// SetFollowedByNil sets the value for FollowedBy to be an explicit nil
func (o *UserUserDTO) SetFollowedByNil() {
	o.FollowedBy.Set(nil)
}

// UnsetFollowedBy ensures that no value is present for FollowedBy, not even an explicit nil
func (o *UserUserDTO) UnsetFollowedBy() {
	o.FollowedBy.Unset()
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *UserUserDTO) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *UserUserDTO) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *UserUserDTO) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *UserUserDTO) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFollowing returns the Following field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetFollowing() bool {
	if o == nil || IsNil(o.Following.Get()) {
		var ret bool
		return ret
	}
	return *o.Following.Get()
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Following.Get(), o.Following.IsSet()
}

// HasFollowing returns a boolean if a field has been set.
func (o *UserUserDTO) HasFollowing() bool {
	if o != nil && o.Following.IsSet() {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given NullableBool and assigns it to the Following field.
func (o *UserUserDTO) SetFollowing(v bool) {
	o.Following.Set(&v)
}
// SetFollowingNil sets the value for Following to be an explicit nil
func (o *UserUserDTO) SetFollowingNil() {
	o.Following.Set(nil)
}

// UnsetFollowing ensures that no value is present for Following, not even an explicit nil
func (o *UserUserDTO) UnsetFollowing() {
	o.Following.Unset()
}

// GetFollowingsCount returns the FollowingsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetFollowingsCount() int32 {
	if o == nil || IsNil(o.FollowingsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowingsCount.Get()
}

// GetFollowingsCountOk returns a tuple with the FollowingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetFollowingsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingsCount.Get(), o.FollowingsCount.IsSet()
}

// HasFollowingsCount returns a boolean if a field has been set.
func (o *UserUserDTO) HasFollowingsCount() bool {
	if o != nil && o.FollowingsCount.IsSet() {
		return true
	}

	return false
}

// SetFollowingsCount gets a reference to the given NullableInt32 and assigns it to the FollowingsCount field.
func (o *UserUserDTO) SetFollowingsCount(v int32) {
	o.FollowingsCount.Set(&v)
}
// SetFollowingsCountNil sets the value for FollowingsCount to be an explicit nil
func (o *UserUserDTO) SetFollowingsCountNil() {
	o.FollowingsCount.Set(nil)
}

// UnsetFollowingsCount ensures that no value is present for FollowingsCount, not even an explicit nil
func (o *UserUserDTO) UnsetFollowingsCount() {
	o.FollowingsCount.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetGender() int32 {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret int32
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetGenderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *UserUserDTO) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableInt32 and assigns it to the Gender field.
func (o *UserUserDTO) SetGender(v int32) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *UserUserDTO) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *UserUserDTO) UnsetGender() {
	o.Gender.Unset()
}

// GetGroupsUsersCount returns the GroupsUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetGroupsUsersCount() int32 {
	if o == nil || IsNil(o.GroupsUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.GroupsUsersCount.Get()
}

// GetGroupsUsersCountOk returns a tuple with the GroupsUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetGroupsUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsUsersCount.Get(), o.GroupsUsersCount.IsSet()
}

// HasGroupsUsersCount returns a boolean if a field has been set.
func (o *UserUserDTO) HasGroupsUsersCount() bool {
	if o != nil && o.GroupsUsersCount.IsSet() {
		return true
	}

	return false
}

// SetGroupsUsersCount gets a reference to the given NullableInt32 and assigns it to the GroupsUsersCount field.
func (o *UserUserDTO) SetGroupsUsersCount(v int32) {
	o.GroupsUsersCount.Set(&v)
}
// SetGroupsUsersCountNil sets the value for GroupsUsersCount to be an explicit nil
func (o *UserUserDTO) SetGroupsUsersCountNil() {
	o.GroupsUsersCount.Set(nil)
}

// UnsetGroupsUsersCount ensures that no value is present for GroupsUsersCount, not even an explicit nil
func (o *UserUserDTO) UnsetGroupsUsersCount() {
	o.GroupsUsersCount.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetHidden() bool {
	if o == nil || IsNil(o.Hidden.Get()) {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *UserUserDTO) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *UserUserDTO) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *UserUserDTO) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *UserUserDTO) UnsetHidden() {
	o.Hidden.Unset()
}

// GetHideVip returns the HideVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetHideVip() bool {
	if o == nil || IsNil(o.HideVip.Get()) {
		var ret bool
		return ret
	}
	return *o.HideVip.Get()
}

// GetHideVipOk returns a tuple with the HideVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetHideVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideVip.Get(), o.HideVip.IsSet()
}

// HasHideVip returns a boolean if a field has been set.
func (o *UserUserDTO) HasHideVip() bool {
	if o != nil && o.HideVip.IsSet() {
		return true
	}

	return false
}

// SetHideVip gets a reference to the given NullableBool and assigns it to the HideVip field.
func (o *UserUserDTO) SetHideVip(v bool) {
	o.HideVip.Set(&v)
}
// SetHideVipNil sets the value for HideVip to be an explicit nil
func (o *UserUserDTO) SetHideVipNil() {
	o.HideVip.Set(nil)
}

// UnsetHideVip ensures that no value is present for HideVip, not even an explicit nil
func (o *UserUserDTO) UnsetHideVip() {
	o.HideVip.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UserUserDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *UserUserDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UserUserDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UserUserDTO) UnsetId() {
	o.Id.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *UserUserDTO) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *UserUserDTO) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *UserUserDTO) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *UserUserDTO) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetNewUser returns the NewUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetNewUser() bool {
	if o == nil || IsNil(o.NewUser.Get()) {
		var ret bool
		return ret
	}
	return *o.NewUser.Get()
}

// GetNewUserOk returns a tuple with the NewUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetNewUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewUser.Get(), o.NewUser.IsSet()
}

// HasNewUser returns a boolean if a field has been set.
func (o *UserUserDTO) HasNewUser() bool {
	if o != nil && o.NewUser.IsSet() {
		return true
	}

	return false
}

// SetNewUser gets a reference to the given NullableBool and assigns it to the NewUser field.
func (o *UserUserDTO) SetNewUser(v bool) {
	o.NewUser.Set(&v)
}
// SetNewUserNil sets the value for NewUser to be an explicit nil
func (o *UserUserDTO) SetNewUserNil() {
	o.NewUser.Set(nil)
}

// UnsetNewUser ensures that no value is present for NewUser, not even an explicit nil
func (o *UserUserDTO) UnsetNewUser() {
	o.NewUser.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *UserUserDTO) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *UserUserDTO) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *UserUserDTO) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *UserUserDTO) UnsetNickname() {
	o.Nickname.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetPostsCount() int32 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetPostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *UserUserDTO) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt32 and assigns it to the PostsCount field.
func (o *UserUserDTO) SetPostsCount(v int32) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *UserUserDTO) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *UserUserDTO) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetPrefecture returns the Prefecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetPrefecture() string {
	if o == nil || IsNil(o.Prefecture.Get()) {
		var ret string
		return ret
	}
	return *o.Prefecture.Get()
}

// GetPrefectureOk returns a tuple with the Prefecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetPrefectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefecture.Get(), o.Prefecture.IsSet()
}

// HasPrefecture returns a boolean if a field has been set.
func (o *UserUserDTO) HasPrefecture() bool {
	if o != nil && o.Prefecture.IsSet() {
		return true
	}

	return false
}

// SetPrefecture gets a reference to the given NullableString and assigns it to the Prefecture field.
func (o *UserUserDTO) SetPrefecture(v string) {
	o.Prefecture.Set(&v)
}
// SetPrefectureNil sets the value for Prefecture to be an explicit nil
func (o *UserUserDTO) SetPrefectureNil() {
	o.Prefecture.Set(nil)
}

// UnsetPrefecture ensures that no value is present for Prefecture, not even an explicit nil
func (o *UserUserDTO) UnsetPrefecture() {
	o.Prefecture.Unset()
}

// GetProfileIcon returns the ProfileIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetProfileIcon() string {
	if o == nil || IsNil(o.ProfileIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIcon.Get()
}

// GetProfileIconOk returns a tuple with the ProfileIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetProfileIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIcon.Get(), o.ProfileIcon.IsSet()
}

// HasProfileIcon returns a boolean if a field has been set.
func (o *UserUserDTO) HasProfileIcon() bool {
	if o != nil && o.ProfileIcon.IsSet() {
		return true
	}

	return false
}

// SetProfileIcon gets a reference to the given NullableString and assigns it to the ProfileIcon field.
func (o *UserUserDTO) SetProfileIcon(v string) {
	o.ProfileIcon.Set(&v)
}
// SetProfileIconNil sets the value for ProfileIcon to be an explicit nil
func (o *UserUserDTO) SetProfileIconNil() {
	o.ProfileIcon.Set(nil)
}

// UnsetProfileIcon ensures that no value is present for ProfileIcon, not even an explicit nil
func (o *UserUserDTO) UnsetProfileIcon() {
	o.ProfileIcon.Unset()
}

// GetProfileIconThumbnail returns the ProfileIconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetProfileIconThumbnail() string {
	if o == nil || IsNil(o.ProfileIconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIconThumbnail.Get()
}

// GetProfileIconThumbnailOk returns a tuple with the ProfileIconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetProfileIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIconThumbnail.Get(), o.ProfileIconThumbnail.IsSet()
}

// HasProfileIconThumbnail returns a boolean if a field has been set.
func (o *UserUserDTO) HasProfileIconThumbnail() bool {
	if o != nil && o.ProfileIconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetProfileIconThumbnail gets a reference to the given NullableString and assigns it to the ProfileIconThumbnail field.
func (o *UserUserDTO) SetProfileIconThumbnail(v string) {
	o.ProfileIconThumbnail.Set(&v)
}
// SetProfileIconThumbnailNil sets the value for ProfileIconThumbnail to be an explicit nil
func (o *UserUserDTO) SetProfileIconThumbnailNil() {
	o.ProfileIconThumbnail.Set(nil)
}

// UnsetProfileIconThumbnail ensures that no value is present for ProfileIconThumbnail, not even an explicit nil
func (o *UserUserDTO) UnsetProfileIconThumbnail() {
	o.ProfileIconThumbnail.Unset()
}

// GetRecentlyKenta returns the RecentlyKenta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetRecentlyKenta() bool {
	if o == nil || IsNil(o.RecentlyKenta.Get()) {
		var ret bool
		return ret
	}
	return *o.RecentlyKenta.Get()
}

// GetRecentlyKentaOk returns a tuple with the RecentlyKenta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetRecentlyKentaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecentlyKenta.Get(), o.RecentlyKenta.IsSet()
}

// HasRecentlyKenta returns a boolean if a field has been set.
func (o *UserUserDTO) HasRecentlyKenta() bool {
	if o != nil && o.RecentlyKenta.IsSet() {
		return true
	}

	return false
}

// SetRecentlyKenta gets a reference to the given NullableBool and assigns it to the RecentlyKenta field.
func (o *UserUserDTO) SetRecentlyKenta(v bool) {
	o.RecentlyKenta.Set(&v)
}
// SetRecentlyKentaNil sets the value for RecentlyKenta to be an explicit nil
func (o *UserUserDTO) SetRecentlyKentaNil() {
	o.RecentlyKenta.Set(nil)
}

// UnsetRecentlyKenta ensures that no value is present for RecentlyKenta, not even an explicit nil
func (o *UserUserDTO) UnsetRecentlyKenta() {
	o.RecentlyKenta.Unset()
}

// GetReviewsCount returns the ReviewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetReviewsCount() int32 {
	if o == nil || IsNil(o.ReviewsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReviewsCount.Get()
}

// GetReviewsCountOk returns a tuple with the ReviewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetReviewsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewsCount.Get(), o.ReviewsCount.IsSet()
}

// HasReviewsCount returns a boolean if a field has been set.
func (o *UserUserDTO) HasReviewsCount() bool {
	if o != nil && o.ReviewsCount.IsSet() {
		return true
	}

	return false
}

// SetReviewsCount gets a reference to the given NullableInt32 and assigns it to the ReviewsCount field.
func (o *UserUserDTO) SetReviewsCount(v int32) {
	o.ReviewsCount.Set(&v)
}
// SetReviewsCountNil sets the value for ReviewsCount to be an explicit nil
func (o *UserUserDTO) SetReviewsCountNil() {
	o.ReviewsCount.Set(nil)
}

// UnsetReviewsCount ensures that no value is present for ReviewsCount, not even an explicit nil
func (o *UserUserDTO) UnsetReviewsCount() {
	o.ReviewsCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UserUserDTO) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UserUserDTO) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UserUserDTO) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UserUserDTO) UnsetTitle() {
	o.Title.Unset()
}

// GetVip returns the Vip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUserDTO) GetVip() bool {
	if o == nil || IsNil(o.Vip.Get()) {
		var ret bool
		return ret
	}
	return *o.Vip.Get()
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUserDTO) GetVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vip.Get(), o.Vip.IsSet()
}

// HasVip returns a boolean if a field has been set.
func (o *UserUserDTO) HasVip() bool {
	if o != nil && o.Vip.IsSet() {
		return true
	}

	return false
}

// SetVip gets a reference to the given NullableBool and assigns it to the Vip field.
func (o *UserUserDTO) SetVip(v bool) {
	o.Vip.Set(&v)
}
// SetVipNil sets the value for Vip to be an explicit nil
func (o *UserUserDTO) SetVipNil() {
	o.Vip.Set(nil)
}

// UnsetVip ensures that no value is present for Vip, not even an explicit nil
func (o *UserUserDTO) UnsetVip() {
	o.Vip.Unset()
}

func (o UserUserDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgeVerified.IsSet() {
		toSerialize["age_verified"] = o.AgeVerified.Get()
	}
	if o.Biography.IsSet() {
		toSerialize["biography"] = o.Biography.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.CoverImageThumbnail.IsSet() {
		toSerialize["cover_image_thumbnail"] = o.CoverImageThumbnail.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DangerousUser.IsSet() {
		toSerialize["dangerous_user"] = o.DangerousUser.Get()
	}
	if o.FollowPending.IsSet() {
		toSerialize["follow_pending"] = o.FollowPending.Get()
	}
	if o.FollowedBy.IsSet() {
		toSerialize["followed_by"] = o.FollowedBy.Get()
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
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
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
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.NewUser.IsSet() {
		toSerialize["new_user"] = o.NewUser.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
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
	if o.RecentlyKenta.IsSet() {
		toSerialize["recently_kenta"] = o.RecentlyKenta.Get()
	}
	if o.ReviewsCount.IsSet() {
		toSerialize["reviews_count"] = o.ReviewsCount.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Vip.IsSet() {
		toSerialize["vip"] = o.Vip.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserUserDTO) UnmarshalJSON(data []byte) (err error) {
	varUserUserDTO := _UserUserDTO{}

	err = json.Unmarshal(data, &varUserUserDTO)

	if err != nil {
		return err
	}

	*o = UserUserDTO(varUserUserDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "age_verified")
		delete(additionalProperties, "biography")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "cover_image_thumbnail")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "dangerous_user")
		delete(additionalProperties, "follow_pending")
		delete(additionalProperties, "followed_by")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "following")
		delete(additionalProperties, "followings_count")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "groups_users_count")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "hide_vip")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "new_user")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "prefecture")
		delete(additionalProperties, "profile_icon")
		delete(additionalProperties, "profile_icon_thumbnail")
		delete(additionalProperties, "recently_kenta")
		delete(additionalProperties, "reviews_count")
		delete(additionalProperties, "title")
		delete(additionalProperties, "vip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserUserDTO struct {
	value *UserUserDTO
	isSet bool
}

func (v NullableUserUserDTO) Get() *UserUserDTO {
	return v.value
}

func (v *NullableUserUserDTO) Set(val *UserUserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserDTO(val *UserUserDTO) *NullableUserUserDTO {
	return &NullableUserUserDTO{value: val, isSet: true}
}

func (v NullableUserUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


