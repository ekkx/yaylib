
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ConferenceCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceCall{}

// ConferenceCall struct for ConferenceCall
type ConferenceCall struct {
	AgoraChannel NullableString `json:"agora_channel,omitempty"`
	AgoraToken NullableString `json:"agora_token,omitempty"`
	AnonymousCallUsersCount NullableInt32 `json:"anonymous_call_users_count,omitempty"`
	BumpParams NullableBumpParams `json:"bump_params,omitempty"`
	CallType NullableString `json:"call_type,omitempty"`
	ConferenceCallUserUuid NullableString `json:"conference_call_user_uuid,omitempty"`
	ConferenceCallUsers []User `json:"conference_call_users,omitempty"`
	ConferenceCallUsersRole []ModelConferenceCallUserRole `json:"conference_call_users_role,omitempty"`
	DurationSeconds NullableInt64 `json:"duration_seconds,omitempty"`
	Game NullableGame `json:"game,omitempty"`
	Genre NullableGenre `json:"genre,omitempty"`
	GroupId NullableInt64 `json:"group_id,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsActive NullableBool `json:"is_active,omitempty"`
	JoinableBy NullableString `json:"joinable_by,omitempty"`
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	PostId NullableInt64 `json:"post_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConferenceCall ConferenceCall

// NewConferenceCall instantiates a new ConferenceCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceCall() *ConferenceCall {
	this := ConferenceCall{}
	return &this
}

// NewConferenceCallWithDefaults instantiates a new ConferenceCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceCallWithDefaults() *ConferenceCall {
	this := ConferenceCall{}
	return &this
}

// GetAgoraChannel returns the AgoraChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetAgoraChannel() string {
	if o == nil || IsNil(o.AgoraChannel.Get()) {
		var ret string
		return ret
	}
	return *o.AgoraChannel.Get()
}

// GetAgoraChannelOk returns a tuple with the AgoraChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetAgoraChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgoraChannel.Get(), o.AgoraChannel.IsSet()
}

// HasAgoraChannel returns a boolean if a field has been set.
func (o *ConferenceCall) HasAgoraChannel() bool {
	if o != nil && o.AgoraChannel.IsSet() {
		return true
	}

	return false
}

// SetAgoraChannel gets a reference to the given NullableString and assigns it to the AgoraChannel field.
func (o *ConferenceCall) SetAgoraChannel(v string) {
	o.AgoraChannel.Set(&v)
}
// SetAgoraChannelNil sets the value for AgoraChannel to be an explicit nil
func (o *ConferenceCall) SetAgoraChannelNil() {
	o.AgoraChannel.Set(nil)
}

// UnsetAgoraChannel ensures that no value is present for AgoraChannel, not even an explicit nil
func (o *ConferenceCall) UnsetAgoraChannel() {
	o.AgoraChannel.Unset()
}

// GetAgoraToken returns the AgoraToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetAgoraToken() string {
	if o == nil || IsNil(o.AgoraToken.Get()) {
		var ret string
		return ret
	}
	return *o.AgoraToken.Get()
}

// GetAgoraTokenOk returns a tuple with the AgoraToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetAgoraTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgoraToken.Get(), o.AgoraToken.IsSet()
}

// HasAgoraToken returns a boolean if a field has been set.
func (o *ConferenceCall) HasAgoraToken() bool {
	if o != nil && o.AgoraToken.IsSet() {
		return true
	}

	return false
}

// SetAgoraToken gets a reference to the given NullableString and assigns it to the AgoraToken field.
func (o *ConferenceCall) SetAgoraToken(v string) {
	o.AgoraToken.Set(&v)
}
// SetAgoraTokenNil sets the value for AgoraToken to be an explicit nil
func (o *ConferenceCall) SetAgoraTokenNil() {
	o.AgoraToken.Set(nil)
}

// UnsetAgoraToken ensures that no value is present for AgoraToken, not even an explicit nil
func (o *ConferenceCall) UnsetAgoraToken() {
	o.AgoraToken.Unset()
}

// GetAnonymousCallUsersCount returns the AnonymousCallUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetAnonymousCallUsersCount() int32 {
	if o == nil || IsNil(o.AnonymousCallUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AnonymousCallUsersCount.Get()
}

// GetAnonymousCallUsersCountOk returns a tuple with the AnonymousCallUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetAnonymousCallUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnonymousCallUsersCount.Get(), o.AnonymousCallUsersCount.IsSet()
}

// HasAnonymousCallUsersCount returns a boolean if a field has been set.
func (o *ConferenceCall) HasAnonymousCallUsersCount() bool {
	if o != nil && o.AnonymousCallUsersCount.IsSet() {
		return true
	}

	return false
}

// SetAnonymousCallUsersCount gets a reference to the given NullableInt32 and assigns it to the AnonymousCallUsersCount field.
func (o *ConferenceCall) SetAnonymousCallUsersCount(v int32) {
	o.AnonymousCallUsersCount.Set(&v)
}
// SetAnonymousCallUsersCountNil sets the value for AnonymousCallUsersCount to be an explicit nil
func (o *ConferenceCall) SetAnonymousCallUsersCountNil() {
	o.AnonymousCallUsersCount.Set(nil)
}

// UnsetAnonymousCallUsersCount ensures that no value is present for AnonymousCallUsersCount, not even an explicit nil
func (o *ConferenceCall) UnsetAnonymousCallUsersCount() {
	o.AnonymousCallUsersCount.Unset()
}

// GetBumpParams returns the BumpParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetBumpParams() BumpParams {
	if o == nil || IsNil(o.BumpParams.Get()) {
		var ret BumpParams
		return ret
	}
	return *o.BumpParams.Get()
}

// GetBumpParamsOk returns a tuple with the BumpParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetBumpParamsOk() (*BumpParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.BumpParams.Get(), o.BumpParams.IsSet()
}

// HasBumpParams returns a boolean if a field has been set.
func (o *ConferenceCall) HasBumpParams() bool {
	if o != nil && o.BumpParams.IsSet() {
		return true
	}

	return false
}

// SetBumpParams gets a reference to the given NullableBumpParams and assigns it to the BumpParams field.
func (o *ConferenceCall) SetBumpParams(v BumpParams) {
	o.BumpParams.Set(&v)
}
// SetBumpParamsNil sets the value for BumpParams to be an explicit nil
func (o *ConferenceCall) SetBumpParamsNil() {
	o.BumpParams.Set(nil)
}

// UnsetBumpParams ensures that no value is present for BumpParams, not even an explicit nil
func (o *ConferenceCall) UnsetBumpParams() {
	o.BumpParams.Unset()
}

// GetCallType returns the CallType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetCallType() string {
	if o == nil || IsNil(o.CallType.Get()) {
		var ret string
		return ret
	}
	return *o.CallType.Get()
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetCallTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallType.Get(), o.CallType.IsSet()
}

// HasCallType returns a boolean if a field has been set.
func (o *ConferenceCall) HasCallType() bool {
	if o != nil && o.CallType.IsSet() {
		return true
	}

	return false
}

// SetCallType gets a reference to the given NullableString and assigns it to the CallType field.
func (o *ConferenceCall) SetCallType(v string) {
	o.CallType.Set(&v)
}
// SetCallTypeNil sets the value for CallType to be an explicit nil
func (o *ConferenceCall) SetCallTypeNil() {
	o.CallType.Set(nil)
}

// UnsetCallType ensures that no value is present for CallType, not even an explicit nil
func (o *ConferenceCall) UnsetCallType() {
	o.CallType.Unset()
}

// GetConferenceCallUserUuid returns the ConferenceCallUserUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetConferenceCallUserUuid() string {
	if o == nil || IsNil(o.ConferenceCallUserUuid.Get()) {
		var ret string
		return ret
	}
	return *o.ConferenceCallUserUuid.Get()
}

// GetConferenceCallUserUuidOk returns a tuple with the ConferenceCallUserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetConferenceCallUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCallUserUuid.Get(), o.ConferenceCallUserUuid.IsSet()
}

// HasConferenceCallUserUuid returns a boolean if a field has been set.
func (o *ConferenceCall) HasConferenceCallUserUuid() bool {
	if o != nil && o.ConferenceCallUserUuid.IsSet() {
		return true
	}

	return false
}

// SetConferenceCallUserUuid gets a reference to the given NullableString and assigns it to the ConferenceCallUserUuid field.
func (o *ConferenceCall) SetConferenceCallUserUuid(v string) {
	o.ConferenceCallUserUuid.Set(&v)
}
// SetConferenceCallUserUuidNil sets the value for ConferenceCallUserUuid to be an explicit nil
func (o *ConferenceCall) SetConferenceCallUserUuidNil() {
	o.ConferenceCallUserUuid.Set(nil)
}

// UnsetConferenceCallUserUuid ensures that no value is present for ConferenceCallUserUuid, not even an explicit nil
func (o *ConferenceCall) UnsetConferenceCallUserUuid() {
	o.ConferenceCallUserUuid.Unset()
}

// GetConferenceCallUsers returns the ConferenceCallUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetConferenceCallUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.ConferenceCallUsers
}

// GetConferenceCallUsersOk returns a tuple with the ConferenceCallUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetConferenceCallUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.ConferenceCallUsers) {
		return nil, false
	}
	return o.ConferenceCallUsers, true
}

// HasConferenceCallUsers returns a boolean if a field has been set.
func (o *ConferenceCall) HasConferenceCallUsers() bool {
	if o != nil && !IsNil(o.ConferenceCallUsers) {
		return true
	}

	return false
}

// SetConferenceCallUsers gets a reference to the given []User and assigns it to the ConferenceCallUsers field.
func (o *ConferenceCall) SetConferenceCallUsers(v []User) {
	o.ConferenceCallUsers = v
}

// GetConferenceCallUsersRole returns the ConferenceCallUsersRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetConferenceCallUsersRole() []ModelConferenceCallUserRole {
	if o == nil {
		var ret []ModelConferenceCallUserRole
		return ret
	}
	return o.ConferenceCallUsersRole
}

// GetConferenceCallUsersRoleOk returns a tuple with the ConferenceCallUsersRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetConferenceCallUsersRoleOk() ([]ModelConferenceCallUserRole, bool) {
	if o == nil || IsNil(o.ConferenceCallUsersRole) {
		return nil, false
	}
	return o.ConferenceCallUsersRole, true
}

// HasConferenceCallUsersRole returns a boolean if a field has been set.
func (o *ConferenceCall) HasConferenceCallUsersRole() bool {
	if o != nil && !IsNil(o.ConferenceCallUsersRole) {
		return true
	}

	return false
}

// SetConferenceCallUsersRole gets a reference to the given []ModelConferenceCallUserRole and assigns it to the ConferenceCallUsersRole field.
func (o *ConferenceCall) SetConferenceCallUsersRole(v []ModelConferenceCallUserRole) {
	o.ConferenceCallUsersRole = v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetDurationSeconds() int64 {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *ConferenceCall) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableInt64 and assigns it to the DurationSeconds field.
func (o *ConferenceCall) SetDurationSeconds(v int64) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *ConferenceCall) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *ConferenceCall) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

// GetGame returns the Game field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetGame() Game {
	if o == nil || IsNil(o.Game.Get()) {
		var ret Game
		return ret
	}
	return *o.Game.Get()
}

// GetGameOk returns a tuple with the Game field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetGameOk() (*Game, bool) {
	if o == nil {
		return nil, false
	}
	return o.Game.Get(), o.Game.IsSet()
}

// HasGame returns a boolean if a field has been set.
func (o *ConferenceCall) HasGame() bool {
	if o != nil && o.Game.IsSet() {
		return true
	}

	return false
}

// SetGame gets a reference to the given NullableGame and assigns it to the Game field.
func (o *ConferenceCall) SetGame(v Game) {
	o.Game.Set(&v)
}
// SetGameNil sets the value for Game to be an explicit nil
func (o *ConferenceCall) SetGameNil() {
	o.Game.Set(nil)
}

// UnsetGame ensures that no value is present for Game, not even an explicit nil
func (o *ConferenceCall) UnsetGame() {
	o.Game.Unset()
}

// GetGenre returns the Genre field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetGenre() Genre {
	if o == nil || IsNil(o.Genre.Get()) {
		var ret Genre
		return ret
	}
	return *o.Genre.Get()
}

// GetGenreOk returns a tuple with the Genre field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetGenreOk() (*Genre, bool) {
	if o == nil {
		return nil, false
	}
	return o.Genre.Get(), o.Genre.IsSet()
}

// HasGenre returns a boolean if a field has been set.
func (o *ConferenceCall) HasGenre() bool {
	if o != nil && o.Genre.IsSet() {
		return true
	}

	return false
}

// SetGenre gets a reference to the given NullableGenre and assigns it to the Genre field.
func (o *ConferenceCall) SetGenre(v Genre) {
	o.Genre.Set(&v)
}
// SetGenreNil sets the value for Genre to be an explicit nil
func (o *ConferenceCall) SetGenreNil() {
	o.Genre.Set(nil)
}

// UnsetGenre ensures that no value is present for Genre, not even an explicit nil
func (o *ConferenceCall) UnsetGenre() {
	o.Genre.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *ConferenceCall) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *ConferenceCall) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *ConferenceCall) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *ConferenceCall) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ConferenceCall) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ConferenceCall) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ConferenceCall) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ConferenceCall) UnsetId() {
	o.Id.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *ConferenceCall) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *ConferenceCall) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *ConferenceCall) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *ConferenceCall) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetJoinableBy returns the JoinableBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetJoinableBy() string {
	if o == nil || IsNil(o.JoinableBy.Get()) {
		var ret string
		return ret
	}
	return *o.JoinableBy.Get()
}

// GetJoinableByOk returns a tuple with the JoinableBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetJoinableByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinableBy.Get(), o.JoinableBy.IsSet()
}

// HasJoinableBy returns a boolean if a field has been set.
func (o *ConferenceCall) HasJoinableBy() bool {
	if o != nil && o.JoinableBy.IsSet() {
		return true
	}

	return false
}

// SetJoinableBy gets a reference to the given NullableString and assigns it to the JoinableBy field.
func (o *ConferenceCall) SetJoinableBy(v string) {
	o.JoinableBy.Set(&v)
}
// SetJoinableByNil sets the value for JoinableBy to be an explicit nil
func (o *ConferenceCall) SetJoinableByNil() {
	o.JoinableBy.Set(nil)
}

// UnsetJoinableBy ensures that no value is present for JoinableBy, not even an explicit nil
func (o *ConferenceCall) UnsetJoinableBy() {
	o.JoinableBy.Unset()
}

// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *ConferenceCall) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *ConferenceCall) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}
// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *ConferenceCall) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *ConferenceCall) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetPostId returns the PostId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCall) GetPostId() int64 {
	if o == nil || IsNil(o.PostId.Get()) {
		var ret int64
		return ret
	}
	return *o.PostId.Get()
}

// GetPostIdOk returns a tuple with the PostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCall) GetPostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostId.Get(), o.PostId.IsSet()
}

// HasPostId returns a boolean if a field has been set.
func (o *ConferenceCall) HasPostId() bool {
	if o != nil && o.PostId.IsSet() {
		return true
	}

	return false
}

// SetPostId gets a reference to the given NullableInt64 and assigns it to the PostId field.
func (o *ConferenceCall) SetPostId(v int64) {
	o.PostId.Set(&v)
}
// SetPostIdNil sets the value for PostId to be an explicit nil
func (o *ConferenceCall) SetPostIdNil() {
	o.PostId.Set(nil)
}

// UnsetPostId ensures that no value is present for PostId, not even an explicit nil
func (o *ConferenceCall) UnsetPostId() {
	o.PostId.Unset()
}

func (o ConferenceCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgoraChannel.IsSet() {
		toSerialize["agora_channel"] = o.AgoraChannel.Get()
	}
	if o.AgoraToken.IsSet() {
		toSerialize["agora_token"] = o.AgoraToken.Get()
	}
	if o.AnonymousCallUsersCount.IsSet() {
		toSerialize["anonymous_call_users_count"] = o.AnonymousCallUsersCount.Get()
	}
	if o.BumpParams.IsSet() {
		toSerialize["bump_params"] = o.BumpParams.Get()
	}
	if o.CallType.IsSet() {
		toSerialize["call_type"] = o.CallType.Get()
	}
	if o.ConferenceCallUserUuid.IsSet() {
		toSerialize["conference_call_user_uuid"] = o.ConferenceCallUserUuid.Get()
	}
	if o.ConferenceCallUsers != nil {
		toSerialize["conference_call_users"] = o.ConferenceCallUsers
	}
	if o.ConferenceCallUsersRole != nil {
		toSerialize["conference_call_users_role"] = o.ConferenceCallUsersRole
	}
	if o.DurationSeconds.IsSet() {
		toSerialize["duration_seconds"] = o.DurationSeconds.Get()
	}
	if o.Game.IsSet() {
		toSerialize["game"] = o.Game.Get()
	}
	if o.Genre.IsSet() {
		toSerialize["genre"] = o.Genre.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["is_active"] = o.IsActive.Get()
	}
	if o.JoinableBy.IsSet() {
		toSerialize["joinable_by"] = o.JoinableBy.Get()
	}
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if o.PostId.IsSet() {
		toSerialize["post_id"] = o.PostId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConferenceCall) UnmarshalJSON(data []byte) (err error) {
	varConferenceCall := _ConferenceCall{}

	err = json.Unmarshal(data, &varConferenceCall)

	if err != nil {
		return err
	}

	*o = ConferenceCall(varConferenceCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agora_channel")
		delete(additionalProperties, "agora_token")
		delete(additionalProperties, "anonymous_call_users_count")
		delete(additionalProperties, "bump_params")
		delete(additionalProperties, "call_type")
		delete(additionalProperties, "conference_call_user_uuid")
		delete(additionalProperties, "conference_call_users")
		delete(additionalProperties, "conference_call_users_role")
		delete(additionalProperties, "duration_seconds")
		delete(additionalProperties, "game")
		delete(additionalProperties, "genre")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "joinable_by")
		delete(additionalProperties, "max_participants")
		delete(additionalProperties, "post_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConferenceCall struct {
	value *ConferenceCall
	isSet bool
}

func (v NullableConferenceCall) Get() *ConferenceCall {
	return v.value
}

func (v *NullableConferenceCall) Set(val *ConferenceCall) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceCall) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceCall(val *ConferenceCall) *NullableConferenceCall {
	return &NullableConferenceCall{value: val, isSet: true}
}

func (v NullableConferenceCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


