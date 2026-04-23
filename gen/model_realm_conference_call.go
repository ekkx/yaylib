
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmConferenceCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmConferenceCall{}

// RealmConferenceCall struct for RealmConferenceCall
type RealmConferenceCall struct {
	Active NullableBool `json:"active,omitempty"`
	AgoraChannel NullableString `json:"agora_channel,omitempty"`
	AgoraToken NullableString `json:"agora_token,omitempty"`
	AnonymousCallUsersCount NullableInt32 `json:"anonymous_call_users_count,omitempty"`
	BumpParams NullableConferenceCallBumpParams `json:"bump_params,omitempty"`
	CallType NullableString `json:"call_type,omitempty"`
	ConferenceCallUserRoles []ConferenceCallUserRole `json:"conference_call_user_roles,omitempty"`
	ConferenceCallUsers []RealmUser `json:"conference_call_users,omitempty"`
	ConferenceCallUsersCount NullableInt32 `json:"conference_call_users_count,omitempty"`
	DurationSeconds NullableInt64 `json:"duration_seconds,omitempty"`
	Game NullableRealmGame `json:"game,omitempty"`
	Genre NullableRealmGenre `json:"genre,omitempty"`
	GroupId NullableInt64 `json:"group_id,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	JoinableBy NullableString `json:"joinable_by,omitempty"`
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	PostId NullableInt64 `json:"post_id,omitempty"`
	Server NullableString `json:"server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmConferenceCall RealmConferenceCall

// NewRealmConferenceCall instantiates a new RealmConferenceCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmConferenceCall() *RealmConferenceCall {
	this := RealmConferenceCall{}
	return &this
}

// NewRealmConferenceCallWithDefaults instantiates a new RealmConferenceCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmConferenceCallWithDefaults() *RealmConferenceCall {
	this := RealmConferenceCall{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *RealmConferenceCall) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *RealmConferenceCall) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *RealmConferenceCall) UnsetActive() {
	o.Active.Unset()
}

// GetAgoraChannel returns the AgoraChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetAgoraChannel() string {
	if o == nil || IsNil(o.AgoraChannel.Get()) {
		var ret string
		return ret
	}
	return *o.AgoraChannel.Get()
}

// GetAgoraChannelOk returns a tuple with the AgoraChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetAgoraChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgoraChannel.Get(), o.AgoraChannel.IsSet()
}

// HasAgoraChannel returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasAgoraChannel() bool {
	if o != nil && o.AgoraChannel.IsSet() {
		return true
	}

	return false
}

// SetAgoraChannel gets a reference to the given NullableString and assigns it to the AgoraChannel field.
func (o *RealmConferenceCall) SetAgoraChannel(v string) {
	o.AgoraChannel.Set(&v)
}
// SetAgoraChannelNil sets the value for AgoraChannel to be an explicit nil
func (o *RealmConferenceCall) SetAgoraChannelNil() {
	o.AgoraChannel.Set(nil)
}

// UnsetAgoraChannel ensures that no value is present for AgoraChannel, not even an explicit nil
func (o *RealmConferenceCall) UnsetAgoraChannel() {
	o.AgoraChannel.Unset()
}

// GetAgoraToken returns the AgoraToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetAgoraToken() string {
	if o == nil || IsNil(o.AgoraToken.Get()) {
		var ret string
		return ret
	}
	return *o.AgoraToken.Get()
}

// GetAgoraTokenOk returns a tuple with the AgoraToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetAgoraTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgoraToken.Get(), o.AgoraToken.IsSet()
}

// HasAgoraToken returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasAgoraToken() bool {
	if o != nil && o.AgoraToken.IsSet() {
		return true
	}

	return false
}

// SetAgoraToken gets a reference to the given NullableString and assigns it to the AgoraToken field.
func (o *RealmConferenceCall) SetAgoraToken(v string) {
	o.AgoraToken.Set(&v)
}
// SetAgoraTokenNil sets the value for AgoraToken to be an explicit nil
func (o *RealmConferenceCall) SetAgoraTokenNil() {
	o.AgoraToken.Set(nil)
}

// UnsetAgoraToken ensures that no value is present for AgoraToken, not even an explicit nil
func (o *RealmConferenceCall) UnsetAgoraToken() {
	o.AgoraToken.Unset()
}

// GetAnonymousCallUsersCount returns the AnonymousCallUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetAnonymousCallUsersCount() int32 {
	if o == nil || IsNil(o.AnonymousCallUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AnonymousCallUsersCount.Get()
}

// GetAnonymousCallUsersCountOk returns a tuple with the AnonymousCallUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetAnonymousCallUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnonymousCallUsersCount.Get(), o.AnonymousCallUsersCount.IsSet()
}

// HasAnonymousCallUsersCount returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasAnonymousCallUsersCount() bool {
	if o != nil && o.AnonymousCallUsersCount.IsSet() {
		return true
	}

	return false
}

// SetAnonymousCallUsersCount gets a reference to the given NullableInt32 and assigns it to the AnonymousCallUsersCount field.
func (o *RealmConferenceCall) SetAnonymousCallUsersCount(v int32) {
	o.AnonymousCallUsersCount.Set(&v)
}
// SetAnonymousCallUsersCountNil sets the value for AnonymousCallUsersCount to be an explicit nil
func (o *RealmConferenceCall) SetAnonymousCallUsersCountNil() {
	o.AnonymousCallUsersCount.Set(nil)
}

// UnsetAnonymousCallUsersCount ensures that no value is present for AnonymousCallUsersCount, not even an explicit nil
func (o *RealmConferenceCall) UnsetAnonymousCallUsersCount() {
	o.AnonymousCallUsersCount.Unset()
}

// GetBumpParams returns the BumpParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetBumpParams() ConferenceCallBumpParams {
	if o == nil || IsNil(o.BumpParams.Get()) {
		var ret ConferenceCallBumpParams
		return ret
	}
	return *o.BumpParams.Get()
}

// GetBumpParamsOk returns a tuple with the BumpParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetBumpParamsOk() (*ConferenceCallBumpParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.BumpParams.Get(), o.BumpParams.IsSet()
}

// HasBumpParams returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasBumpParams() bool {
	if o != nil && o.BumpParams.IsSet() {
		return true
	}

	return false
}

// SetBumpParams gets a reference to the given NullableConferenceCallBumpParams and assigns it to the BumpParams field.
func (o *RealmConferenceCall) SetBumpParams(v ConferenceCallBumpParams) {
	o.BumpParams.Set(&v)
}
// SetBumpParamsNil sets the value for BumpParams to be an explicit nil
func (o *RealmConferenceCall) SetBumpParamsNil() {
	o.BumpParams.Set(nil)
}

// UnsetBumpParams ensures that no value is present for BumpParams, not even an explicit nil
func (o *RealmConferenceCall) UnsetBumpParams() {
	o.BumpParams.Unset()
}

// GetCallType returns the CallType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetCallType() string {
	if o == nil || IsNil(o.CallType.Get()) {
		var ret string
		return ret
	}
	return *o.CallType.Get()
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetCallTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallType.Get(), o.CallType.IsSet()
}

// HasCallType returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasCallType() bool {
	if o != nil && o.CallType.IsSet() {
		return true
	}

	return false
}

// SetCallType gets a reference to the given NullableString and assigns it to the CallType field.
func (o *RealmConferenceCall) SetCallType(v string) {
	o.CallType.Set(&v)
}
// SetCallTypeNil sets the value for CallType to be an explicit nil
func (o *RealmConferenceCall) SetCallTypeNil() {
	o.CallType.Set(nil)
}

// UnsetCallType ensures that no value is present for CallType, not even an explicit nil
func (o *RealmConferenceCall) UnsetCallType() {
	o.CallType.Unset()
}

// GetConferenceCallUserRoles returns the ConferenceCallUserRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetConferenceCallUserRoles() []ConferenceCallUserRole {
	if o == nil {
		var ret []ConferenceCallUserRole
		return ret
	}
	return o.ConferenceCallUserRoles
}

// GetConferenceCallUserRolesOk returns a tuple with the ConferenceCallUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetConferenceCallUserRolesOk() ([]ConferenceCallUserRole, bool) {
	if o == nil || IsNil(o.ConferenceCallUserRoles) {
		return nil, false
	}
	return o.ConferenceCallUserRoles, true
}

// HasConferenceCallUserRoles returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasConferenceCallUserRoles() bool {
	if o != nil && !IsNil(o.ConferenceCallUserRoles) {
		return true
	}

	return false
}

// SetConferenceCallUserRoles gets a reference to the given []ConferenceCallUserRole and assigns it to the ConferenceCallUserRoles field.
func (o *RealmConferenceCall) SetConferenceCallUserRoles(v []ConferenceCallUserRole) {
	o.ConferenceCallUserRoles = v
}

// GetConferenceCallUsers returns the ConferenceCallUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetConferenceCallUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.ConferenceCallUsers
}

// GetConferenceCallUsersOk returns a tuple with the ConferenceCallUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetConferenceCallUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.ConferenceCallUsers) {
		return nil, false
	}
	return o.ConferenceCallUsers, true
}

// HasConferenceCallUsers returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasConferenceCallUsers() bool {
	if o != nil && !IsNil(o.ConferenceCallUsers) {
		return true
	}

	return false
}

// SetConferenceCallUsers gets a reference to the given []RealmUser and assigns it to the ConferenceCallUsers field.
func (o *RealmConferenceCall) SetConferenceCallUsers(v []RealmUser) {
	o.ConferenceCallUsers = v
}

// GetConferenceCallUsersCount returns the ConferenceCallUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetConferenceCallUsersCount() int32 {
	if o == nil || IsNil(o.ConferenceCallUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ConferenceCallUsersCount.Get()
}

// GetConferenceCallUsersCountOk returns a tuple with the ConferenceCallUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetConferenceCallUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCallUsersCount.Get(), o.ConferenceCallUsersCount.IsSet()
}

// HasConferenceCallUsersCount returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasConferenceCallUsersCount() bool {
	if o != nil && o.ConferenceCallUsersCount.IsSet() {
		return true
	}

	return false
}

// SetConferenceCallUsersCount gets a reference to the given NullableInt32 and assigns it to the ConferenceCallUsersCount field.
func (o *RealmConferenceCall) SetConferenceCallUsersCount(v int32) {
	o.ConferenceCallUsersCount.Set(&v)
}
// SetConferenceCallUsersCountNil sets the value for ConferenceCallUsersCount to be an explicit nil
func (o *RealmConferenceCall) SetConferenceCallUsersCountNil() {
	o.ConferenceCallUsersCount.Set(nil)
}

// UnsetConferenceCallUsersCount ensures that no value is present for ConferenceCallUsersCount, not even an explicit nil
func (o *RealmConferenceCall) UnsetConferenceCallUsersCount() {
	o.ConferenceCallUsersCount.Unset()
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetDurationSeconds() int64 {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableInt64 and assigns it to the DurationSeconds field.
func (o *RealmConferenceCall) SetDurationSeconds(v int64) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *RealmConferenceCall) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *RealmConferenceCall) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

// GetGame returns the Game field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetGame() RealmGame {
	if o == nil || IsNil(o.Game.Get()) {
		var ret RealmGame
		return ret
	}
	return *o.Game.Get()
}

// GetGameOk returns a tuple with the Game field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetGameOk() (*RealmGame, bool) {
	if o == nil {
		return nil, false
	}
	return o.Game.Get(), o.Game.IsSet()
}

// HasGame returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasGame() bool {
	if o != nil && o.Game.IsSet() {
		return true
	}

	return false
}

// SetGame gets a reference to the given NullableRealmGame and assigns it to the Game field.
func (o *RealmConferenceCall) SetGame(v RealmGame) {
	o.Game.Set(&v)
}
// SetGameNil sets the value for Game to be an explicit nil
func (o *RealmConferenceCall) SetGameNil() {
	o.Game.Set(nil)
}

// UnsetGame ensures that no value is present for Game, not even an explicit nil
func (o *RealmConferenceCall) UnsetGame() {
	o.Game.Unset()
}

// GetGenre returns the Genre field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetGenre() RealmGenre {
	if o == nil || IsNil(o.Genre.Get()) {
		var ret RealmGenre
		return ret
	}
	return *o.Genre.Get()
}

// GetGenreOk returns a tuple with the Genre field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetGenreOk() (*RealmGenre, bool) {
	if o == nil {
		return nil, false
	}
	return o.Genre.Get(), o.Genre.IsSet()
}

// HasGenre returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasGenre() bool {
	if o != nil && o.Genre.IsSet() {
		return true
	}

	return false
}

// SetGenre gets a reference to the given NullableRealmGenre and assigns it to the Genre field.
func (o *RealmConferenceCall) SetGenre(v RealmGenre) {
	o.Genre.Set(&v)
}
// SetGenreNil sets the value for Genre to be an explicit nil
func (o *RealmConferenceCall) SetGenreNil() {
	o.Genre.Set(nil)
}

// UnsetGenre ensures that no value is present for Genre, not even an explicit nil
func (o *RealmConferenceCall) UnsetGenre() {
	o.Genre.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *RealmConferenceCall) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *RealmConferenceCall) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *RealmConferenceCall) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmConferenceCall) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmConferenceCall) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmConferenceCall) UnsetId() {
	o.Id.Unset()
}

// GetJoinableBy returns the JoinableBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetJoinableBy() string {
	if o == nil || IsNil(o.JoinableBy.Get()) {
		var ret string
		return ret
	}
	return *o.JoinableBy.Get()
}

// GetJoinableByOk returns a tuple with the JoinableBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetJoinableByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinableBy.Get(), o.JoinableBy.IsSet()
}

// HasJoinableBy returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasJoinableBy() bool {
	if o != nil && o.JoinableBy.IsSet() {
		return true
	}

	return false
}

// SetJoinableBy gets a reference to the given NullableString and assigns it to the JoinableBy field.
func (o *RealmConferenceCall) SetJoinableBy(v string) {
	o.JoinableBy.Set(&v)
}
// SetJoinableByNil sets the value for JoinableBy to be an explicit nil
func (o *RealmConferenceCall) SetJoinableByNil() {
	o.JoinableBy.Set(nil)
}

// UnsetJoinableBy ensures that no value is present for JoinableBy, not even an explicit nil
func (o *RealmConferenceCall) UnsetJoinableBy() {
	o.JoinableBy.Unset()
}

// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *RealmConferenceCall) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}
// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *RealmConferenceCall) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *RealmConferenceCall) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetPostId returns the PostId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetPostId() int64 {
	if o == nil || IsNil(o.PostId.Get()) {
		var ret int64
		return ret
	}
	return *o.PostId.Get()
}

// GetPostIdOk returns a tuple with the PostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetPostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostId.Get(), o.PostId.IsSet()
}

// HasPostId returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasPostId() bool {
	if o != nil && o.PostId.IsSet() {
		return true
	}

	return false
}

// SetPostId gets a reference to the given NullableInt64 and assigns it to the PostId field.
func (o *RealmConferenceCall) SetPostId(v int64) {
	o.PostId.Set(&v)
}
// SetPostIdNil sets the value for PostId to be an explicit nil
func (o *RealmConferenceCall) SetPostIdNil() {
	o.PostId.Set(nil)
}

// UnsetPostId ensures that no value is present for PostId, not even an explicit nil
func (o *RealmConferenceCall) UnsetPostId() {
	o.PostId.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmConferenceCall) GetServer() string {
	if o == nil || IsNil(o.Server.Get()) {
		var ret string
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmConferenceCall) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *RealmConferenceCall) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableString and assigns it to the Server field.
func (o *RealmConferenceCall) SetServer(v string) {
	o.Server.Set(&v)
}
// SetServerNil sets the value for Server to be an explicit nil
func (o *RealmConferenceCall) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *RealmConferenceCall) UnsetServer() {
	o.Server.Unset()
}

func (o RealmConferenceCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmConferenceCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
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
	if o.ConferenceCallUserRoles != nil {
		toSerialize["conference_call_user_roles"] = o.ConferenceCallUserRoles
	}
	if o.ConferenceCallUsers != nil {
		toSerialize["conference_call_users"] = o.ConferenceCallUsers
	}
	if o.ConferenceCallUsersCount.IsSet() {
		toSerialize["conference_call_users_count"] = o.ConferenceCallUsersCount.Get()
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
	if o.JoinableBy.IsSet() {
		toSerialize["joinable_by"] = o.JoinableBy.Get()
	}
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if o.PostId.IsSet() {
		toSerialize["post_id"] = o.PostId.Get()
	}
	if o.Server.IsSet() {
		toSerialize["server"] = o.Server.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmConferenceCall) UnmarshalJSON(data []byte) (err error) {
	varRealmConferenceCall := _RealmConferenceCall{}

	err = json.Unmarshal(data, &varRealmConferenceCall)

	if err != nil {
		return err
	}

	*o = RealmConferenceCall(varRealmConferenceCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "agora_channel")
		delete(additionalProperties, "agora_token")
		delete(additionalProperties, "anonymous_call_users_count")
		delete(additionalProperties, "bump_params")
		delete(additionalProperties, "call_type")
		delete(additionalProperties, "conference_call_user_roles")
		delete(additionalProperties, "conference_call_users")
		delete(additionalProperties, "conference_call_users_count")
		delete(additionalProperties, "duration_seconds")
		delete(additionalProperties, "game")
		delete(additionalProperties, "genre")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "joinable_by")
		delete(additionalProperties, "max_participants")
		delete(additionalProperties, "post_id")
		delete(additionalProperties, "server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmConferenceCall struct {
	value *RealmConferenceCall
	isSet bool
}

func (v NullableRealmConferenceCall) Get() *RealmConferenceCall {
	return v.value
}

func (v *NullableRealmConferenceCall) Set(val *RealmConferenceCall) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmConferenceCall) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmConferenceCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmConferenceCall(val *RealmConferenceCall) *NullableRealmConferenceCall {
	return &NullableRealmConferenceCall{value: val, isSet: true}
}

func (v NullableRealmConferenceCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmConferenceCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


