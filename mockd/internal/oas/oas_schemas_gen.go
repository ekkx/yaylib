// Code generated; DO NOT EDIT.

package oas

import (
	"github.com/go-faster/errors"
)

// AcceptChatRequestNoContent is response for AcceptChatRequest operation.
type AcceptChatRequestNoContent struct{}

type AcceptChatRequestReq struct {
	ChatRoomIds []int64 `json:"chat_room_ids[]"`
}

// GetChatRoomIds returns the value of ChatRoomIds.
func (s *AcceptChatRequestReq) GetChatRoomIds() []int64 {
	return s.ChatRoomIds
}

// SetChatRoomIds sets the value of ChatRoomIds.
func (s *AcceptChatRequestReq) SetChatRoomIds(val []int64) {
	s.ChatRoomIds = val
}

// AcceptGroupJoinRequestNoContent is response for AcceptGroupJoinRequest operation.
type AcceptGroupJoinRequestNoContent struct{}

// AcceptGroupTransferNoContent is response for AcceptGroupTransfer operation.
type AcceptGroupTransferNoContent struct{}

// Ref: #/components/schemas/ActiveFollowingsResponse
type ActiveFollowingsResponse struct {
	LastLoggedinAt OptNilInt64          `json:"last_loggedin_at"`
	Users          OptNilRealmUserArray `json:"users"`
}

// GetLastLoggedinAt returns the value of LastLoggedinAt.
func (s *ActiveFollowingsResponse) GetLastLoggedinAt() OptNilInt64 {
	return s.LastLoggedinAt
}

// GetUsers returns the value of Users.
func (s *ActiveFollowingsResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetLastLoggedinAt sets the value of LastLoggedinAt.
func (s *ActiveFollowingsResponse) SetLastLoggedinAt(val OptNilInt64) {
	s.LastLoggedinAt = val
}

// SetUsers sets the value of Users.
func (s *ActiveFollowingsResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// Ref: #/components/schemas/ActivitiesResponse
type ActivitiesResponse struct {
	Activities    OptNilActivityArray `json:"activities"`
	LastTimestamp OptNilInt64         `json:"last_timestamp"`
}

// GetActivities returns the value of Activities.
func (s *ActivitiesResponse) GetActivities() OptNilActivityArray {
	return s.Activities
}

// GetLastTimestamp returns the value of LastTimestamp.
func (s *ActivitiesResponse) GetLastTimestamp() OptNilInt64 {
	return s.LastTimestamp
}

// SetActivities sets the value of Activities.
func (s *ActivitiesResponse) SetActivities(val OptNilActivityArray) {
	s.Activities = val
}

// SetLastTimestamp sets the value of LastTimestamp.
func (s *ActivitiesResponse) SetLastTimestamp(val OptNilInt64) {
	s.LastTimestamp = val
}

// Ref: #/components/schemas/Activity
type Activity struct {
	BirthdayUsers      OptNilRealmUserArray `json:"birthday_users"`
	BirthdayUsersCount OptNilInt32          `json:"birthday_users_count"`
	CreatedAt          OptNilInt64          `json:"created_at"`
	EmplAmount         OptNilInt32          `json:"empl_amount"`
	Followers          OptNilRealmUserArray `json:"followers"`
	FollowersCount     OptNilInt32          `json:"followers_count"`
	FromPost           *Post                `json:"from_post"`
	FromPostIds        OptNilInt64Array     `json:"from_post_ids"`
	GiftItem           OptRealmGift         `json:"gift_item"`
	Group              *Group               `json:"group"`
	ID                 OptNilInt64          `json:"id"`
	Metadata           OptMetadata          `json:"metadata"`
	ToPost             *Post                `json:"to_post"`
	Type               OptNilString         `json:"type"`
	User               *RealmUser           `json:"user"`
	VipReward          OptNilInt32          `json:"vip_reward"`
}

// GetBirthdayUsers returns the value of BirthdayUsers.
func (s *Activity) GetBirthdayUsers() OptNilRealmUserArray {
	return s.BirthdayUsers
}

// GetBirthdayUsersCount returns the value of BirthdayUsersCount.
func (s *Activity) GetBirthdayUsersCount() OptNilInt32 {
	return s.BirthdayUsersCount
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Activity) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetEmplAmount returns the value of EmplAmount.
func (s *Activity) GetEmplAmount() OptNilInt32 {
	return s.EmplAmount
}

// GetFollowers returns the value of Followers.
func (s *Activity) GetFollowers() OptNilRealmUserArray {
	return s.Followers
}

// GetFollowersCount returns the value of FollowersCount.
func (s *Activity) GetFollowersCount() OptNilInt32 {
	return s.FollowersCount
}

// GetFromPost returns the value of FromPost.
func (s *Activity) GetFromPost() *Post {
	return s.FromPost
}

// GetFromPostIds returns the value of FromPostIds.
func (s *Activity) GetFromPostIds() OptNilInt64Array {
	return s.FromPostIds
}

// GetGiftItem returns the value of GiftItem.
func (s *Activity) GetGiftItem() OptRealmGift {
	return s.GiftItem
}

// GetGroup returns the value of Group.
func (s *Activity) GetGroup() *Group {
	return s.Group
}

// GetID returns the value of ID.
func (s *Activity) GetID() OptNilInt64 {
	return s.ID
}

// GetMetadata returns the value of Metadata.
func (s *Activity) GetMetadata() OptMetadata {
	return s.Metadata
}

// GetToPost returns the value of ToPost.
func (s *Activity) GetToPost() *Post {
	return s.ToPost
}

// GetType returns the value of Type.
func (s *Activity) GetType() OptNilString {
	return s.Type
}

// GetUser returns the value of User.
func (s *Activity) GetUser() *RealmUser {
	return s.User
}

// GetVipReward returns the value of VipReward.
func (s *Activity) GetVipReward() OptNilInt32 {
	return s.VipReward
}

// SetBirthdayUsers sets the value of BirthdayUsers.
func (s *Activity) SetBirthdayUsers(val OptNilRealmUserArray) {
	s.BirthdayUsers = val
}

// SetBirthdayUsersCount sets the value of BirthdayUsersCount.
func (s *Activity) SetBirthdayUsersCount(val OptNilInt32) {
	s.BirthdayUsersCount = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Activity) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetEmplAmount sets the value of EmplAmount.
func (s *Activity) SetEmplAmount(val OptNilInt32) {
	s.EmplAmount = val
}

// SetFollowers sets the value of Followers.
func (s *Activity) SetFollowers(val OptNilRealmUserArray) {
	s.Followers = val
}

// SetFollowersCount sets the value of FollowersCount.
func (s *Activity) SetFollowersCount(val OptNilInt32) {
	s.FollowersCount = val
}

// SetFromPost sets the value of FromPost.
func (s *Activity) SetFromPost(val *Post) {
	s.FromPost = val
}

// SetFromPostIds sets the value of FromPostIds.
func (s *Activity) SetFromPostIds(val OptNilInt64Array) {
	s.FromPostIds = val
}

// SetGiftItem sets the value of GiftItem.
func (s *Activity) SetGiftItem(val OptRealmGift) {
	s.GiftItem = val
}

// SetGroup sets the value of Group.
func (s *Activity) SetGroup(val *Group) {
	s.Group = val
}

// SetID sets the value of ID.
func (s *Activity) SetID(val OptNilInt64) {
	s.ID = val
}

// SetMetadata sets the value of Metadata.
func (s *Activity) SetMetadata(val OptMetadata) {
	s.Metadata = val
}

// SetToPost sets the value of ToPost.
func (s *Activity) SetToPost(val *Post) {
	s.ToPost = val
}

// SetType sets the value of Type.
func (s *Activity) SetType(val OptNilString) {
	s.Type = val
}

// SetUser sets the value of User.
func (s *Activity) SetUser(val *RealmUser) {
	s.User = val
}

// SetVipReward sets the value of VipReward.
func (s *Activity) SetVipReward(val OptNilInt32) {
	s.VipReward = val
}

// AddThreadMemberNoContent is response for AddThreadMember operation.
type AddThreadMemberNoContent struct{}

// Ref: #/components/schemas/AdditionalSettingsResponse
type AdditionalSettingsResponse struct {
	Settings OptSettings `json:"settings"`
}

// GetSettings returns the value of Settings.
func (s *AdditionalSettingsResponse) GetSettings() OptSettings {
	return s.Settings
}

// SetSettings sets the value of Settings.
func (s *AdditionalSettingsResponse) SetSettings(val OptSettings) {
	s.Settings = val
}

// AgreePolicyNoContent is response for AgreePolicy operation.
type AgreePolicyNoContent struct{}

// Ref: #/components/schemas/Application
type Application struct {
	Settings OptConfig `json:"settings"`
}

// GetSettings returns the value of Settings.
func (s *Application) GetSettings() OptConfig {
	return s.Settings
}

// SetSettings sets the value of Settings.
func (s *Application) SetSettings(val OptConfig) {
	s.Settings = val
}

// Ref: #/components/schemas/ApplicationConfigResponse
type ApplicationConfigResponse struct {
	App OptApplication `json:"app"`
}

// GetApp returns the value of App.
func (s *ApplicationConfigResponse) GetApp() OptApplication {
	return s.App
}

// SetApp sets the value of App.
func (s *ApplicationConfigResponse) SetApp(val OptApplication) {
	s.App = val
}

// BanGroupUserNoContent is response for BanGroupUser operation.
type BanGroupUserNoContent struct{}

// Ref: #/components/schemas/BanWord
type BanWord struct {
	ID   OptNilInt64  `json:"id"`
	Type OptNilString `json:"type"`
	Word OptNilString `json:"word"`
}

// GetID returns the value of ID.
func (s *BanWord) GetID() OptNilInt64 {
	return s.ID
}

// GetType returns the value of Type.
func (s *BanWord) GetType() OptNilString {
	return s.Type
}

// GetWord returns the value of Word.
func (s *BanWord) GetWord() OptNilString {
	return s.Word
}

// SetID sets the value of ID.
func (s *BanWord) SetID(val OptNilInt64) {
	s.ID = val
}

// SetType sets the value of Type.
func (s *BanWord) SetType(val OptNilString) {
	s.Type = val
}

// SetWord sets the value of Word.
func (s *BanWord) SetWord(val OptNilString) {
	s.Word = val
}

// Ref: #/components/schemas/BanWordsResponse
type BanWordsResponse struct {
	BanWords OptNilBanWordArray `json:"ban_words"`
}

// GetBanWords returns the value of BanWords.
func (s *BanWordsResponse) GetBanWords() OptNilBanWordArray {
	return s.BanWords
}

// SetBanWords sets the value of BanWords.
func (s *BanWordsResponse) SetBanWords(val OptNilBanWordArray) {
	s.BanWords = val
}

// Ref: #/components/schemas/Bgm
type Bgm struct {
	ID       OptNilInt64  `json:"id"`
	MusicURL OptNilString `json:"music_url"`
	Order    OptNilInt32  `json:"order"`
	Title    OptNilString `json:"title"`
}

// GetID returns the value of ID.
func (s *Bgm) GetID() OptNilInt64 {
	return s.ID
}

// GetMusicURL returns the value of MusicURL.
func (s *Bgm) GetMusicURL() OptNilString {
	return s.MusicURL
}

// GetOrder returns the value of Order.
func (s *Bgm) GetOrder() OptNilInt32 {
	return s.Order
}

// GetTitle returns the value of Title.
func (s *Bgm) GetTitle() OptNilString {
	return s.Title
}

// SetID sets the value of ID.
func (s *Bgm) SetID(val OptNilInt64) {
	s.ID = val
}

// SetMusicURL sets the value of MusicURL.
func (s *Bgm) SetMusicURL(val OptNilString) {
	s.MusicURL = val
}

// SetOrder sets the value of Order.
func (s *Bgm) SetOrder(val OptNilInt32) {
	s.Order = val
}

// SetTitle sets the value of Title.
func (s *Bgm) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/BgmsResponse
type BgmsResponse struct {
	Bgm OptNilBgmArray `json:"bgm"`
}

// GetBgm returns the value of Bgm.
func (s *BgmsResponse) GetBgm() OptNilBgmArray {
	return s.Bgm
}

// SetBgm sets the value of Bgm.
func (s *BgmsResponse) SetBgm(val OptNilBgmArray) {
	s.Bgm = val
}

// BlockUserNoContent is response for BlockUser operation.
type BlockUserNoContent struct{}

// Ref: #/components/schemas/BlockedUserIdsResponse
type BlockedUserIdsResponse struct {
	BlockIds OptNilInt64Array `json:"block_ids"`
}

// GetBlockIds returns the value of BlockIds.
func (s *BlockedUserIdsResponse) GetBlockIds() OptNilInt64Array {
	return s.BlockIds
}

// SetBlockIds sets the value of BlockIds.
func (s *BlockedUserIdsResponse) SetBlockIds(val OptNilInt64Array) {
	s.BlockIds = val
}

// Ref: #/components/schemas/BlockedUsersResponse
type BlockedUsersResponse struct {
	BlockedCount OptNilInt32          `json:"blocked_count"`
	LastID       OptNilInt64          `json:"last_id"`
	Users        OptNilRealmUserArray `json:"users"`
}

// GetBlockedCount returns the value of BlockedCount.
func (s *BlockedUsersResponse) GetBlockedCount() OptNilInt32 {
	return s.BlockedCount
}

// GetLastID returns the value of LastID.
func (s *BlockedUsersResponse) GetLastID() OptNilInt64 {
	return s.LastID
}

// GetUsers returns the value of Users.
func (s *BlockedUsersResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetBlockedCount sets the value of BlockedCount.
func (s *BlockedUsersResponse) SetBlockedCount(val OptNilInt32) {
	s.BlockedCount = val
}

// SetLastID sets the value of LastID.
func (s *BlockedUsersResponse) SetLastID(val OptNilInt64) {
	s.LastID = val
}

// SetUsers sets the value of Users.
func (s *BlockedUsersResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// Ref: #/components/schemas/BookmarkPostResponse
type BookmarkPostResponse struct {
	IsBookmarked OptNilBool `json:"is_bookmarked"`
}

// GetIsBookmarked returns the value of IsBookmarked.
func (s *BookmarkPostResponse) GetIsBookmarked() OptNilBool {
	return s.IsBookmarked
}

// SetIsBookmarked sets the value of IsBookmarked.
func (s *BookmarkPostResponse) SetIsBookmarked(val OptNilBool) {
	s.IsBookmarked = val
}

// BulkInviteToCallNoContent is response for BulkInviteToCall operation.
type BulkInviteToCallNoContent struct{}

// BumpCallNoContent is response for BumpCall operation.
type BumpCallNoContent struct{}

// Ref: #/components/schemas/CallActionSignatureResponse
type CallActionSignatureResponse struct {
	SignaturePayload OptSignaturePayload `json:"signature_payload"`
}

// GetSignaturePayload returns the value of SignaturePayload.
func (s *CallActionSignatureResponse) GetSignaturePayload() OptSignaturePayload {
	return s.SignaturePayload
}

// SetSignaturePayload sets the value of SignaturePayload.
func (s *CallActionSignatureResponse) SetSignaturePayload(val OptSignaturePayload) {
	s.SignaturePayload = val
}

// Ref: #/components/schemas/CallGiftHistoryResponse
type CallGiftHistoryResponse struct {
	NextPageValue OptNilString           `json:"next_page_value"`
	SentGifts     OptNilGiftHistoryArray `json:"sent_gifts"`
}

// GetNextPageValue returns the value of NextPageValue.
func (s *CallGiftHistoryResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetSentGifts returns the value of SentGifts.
func (s *CallGiftHistoryResponse) GetSentGifts() OptNilGiftHistoryArray {
	return s.SentGifts
}

// SetNextPageValue sets the value of NextPageValue.
func (s *CallGiftHistoryResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetSentGifts sets the value of SentGifts.
func (s *CallGiftHistoryResponse) SetSentGifts(val OptNilGiftHistoryArray) {
	s.SentGifts = val
}

// Ref: #/components/schemas/CallStatusResponse
type CallStatusResponse struct {
	PhoneStatus OptNilBool   `json:"phone_status"`
	RoomURL     OptNilString `json:"room_url"`
	VideoStatus OptNilBool   `json:"video_status"`
}

// GetPhoneStatus returns the value of PhoneStatus.
func (s *CallStatusResponse) GetPhoneStatus() OptNilBool {
	return s.PhoneStatus
}

// GetRoomURL returns the value of RoomURL.
func (s *CallStatusResponse) GetRoomURL() OptNilString {
	return s.RoomURL
}

// GetVideoStatus returns the value of VideoStatus.
func (s *CallStatusResponse) GetVideoStatus() OptNilBool {
	return s.VideoStatus
}

// SetPhoneStatus sets the value of PhoneStatus.
func (s *CallStatusResponse) SetPhoneStatus(val OptNilBool) {
	s.PhoneStatus = val
}

// SetRoomURL sets the value of RoomURL.
func (s *CallStatusResponse) SetRoomURL(val OptNilString) {
	s.RoomURL = val
}

// SetVideoStatus sets the value of VideoStatus.
func (s *CallStatusResponse) SetVideoStatus(val OptNilBool) {
	s.VideoStatus = val
}

// CancelGroupTransferNoContent is response for CancelGroupTransfer operation.
type CancelGroupTransferNoContent struct{}

type ChangeEmailReq struct {
	APIKey          OptString    `json:"api_key"`
	Email           OptString    `json:"email"`
	EmailGrantToken OptNilString `json:"email_grant_token"`
	Password        OptString    `json:"password"`
}

// GetAPIKey returns the value of APIKey.
func (s *ChangeEmailReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetEmail returns the value of Email.
func (s *ChangeEmailReq) GetEmail() OptString {
	return s.Email
}

// GetEmailGrantToken returns the value of EmailGrantToken.
func (s *ChangeEmailReq) GetEmailGrantToken() OptNilString {
	return s.EmailGrantToken
}

// GetPassword returns the value of Password.
func (s *ChangeEmailReq) GetPassword() OptString {
	return s.Password
}

// SetAPIKey sets the value of APIKey.
func (s *ChangeEmailReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetEmail sets the value of Email.
func (s *ChangeEmailReq) SetEmail(val OptString) {
	s.Email = val
}

// SetEmailGrantToken sets the value of EmailGrantToken.
func (s *ChangeEmailReq) SetEmailGrantToken(val OptNilString) {
	s.EmailGrantToken = val
}

// SetPassword sets the value of Password.
func (s *ChangeEmailReq) SetPassword(val OptString) {
	s.Password = val
}

type ChangePasswordReq struct {
	APIKey          OptString `json:"api_key"`
	CurrentPassword OptString `json:"current_password"`
	Password        OptString `json:"password"`
}

// GetAPIKey returns the value of APIKey.
func (s *ChangePasswordReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetCurrentPassword returns the value of CurrentPassword.
func (s *ChangePasswordReq) GetCurrentPassword() OptString {
	return s.CurrentPassword
}

// GetPassword returns the value of Password.
func (s *ChangePasswordReq) GetPassword() OptString {
	return s.Password
}

// SetAPIKey sets the value of APIKey.
func (s *ChangePasswordReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetCurrentPassword sets the value of CurrentPassword.
func (s *ChangePasswordReq) SetCurrentPassword(val OptString) {
	s.CurrentPassword = val
}

// SetPassword sets the value of Password.
func (s *ChangePasswordReq) SetPassword(val OptString) {
	s.Password = val
}

// Ref: #/components/schemas/ChatInvitation
type ChatInvitation struct {
	FirstUser  OptNilString `json:"first_user"`
	Inviter    OptNilString `json:"inviter"`
	UsersCount OptNilInt32  `json:"users_count"`
}

// GetFirstUser returns the value of FirstUser.
func (s *ChatInvitation) GetFirstUser() OptNilString {
	return s.FirstUser
}

// GetInviter returns the value of Inviter.
func (s *ChatInvitation) GetInviter() OptNilString {
	return s.Inviter
}

// GetUsersCount returns the value of UsersCount.
func (s *ChatInvitation) GetUsersCount() OptNilInt32 {
	return s.UsersCount
}

// SetFirstUser sets the value of FirstUser.
func (s *ChatInvitation) SetFirstUser(val OptNilString) {
	s.FirstUser = val
}

// SetInviter sets the value of Inviter.
func (s *ChatInvitation) SetInviter(val OptNilString) {
	s.Inviter = val
}

// SetUsersCount sets the value of UsersCount.
func (s *ChatInvitation) SetUsersCount(val OptNilInt32) {
	s.UsersCount = val
}

// Ref: #/components/schemas/ChatRoomLastMessage
type ChatRoomLastMessage struct {
	ConferenceCall OptRealmConferenceCall `json:"conference_call"`
	CreatedAt      OptNilInt64            `json:"created_at"`
	ID             OptNilInt64            `json:"id"`
	MessageType    OptMessageType         `json:"message_type"`
	RoomID         OptNilInt64            `json:"room_id"`
	Text           OptNilString           `json:"text"`
	UserID         OptNilInt64            `json:"user_id"`
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *ChatRoomLastMessage) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetCreatedAt returns the value of CreatedAt.
func (s *ChatRoomLastMessage) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetID returns the value of ID.
func (s *ChatRoomLastMessage) GetID() OptNilInt64 {
	return s.ID
}

// GetMessageType returns the value of MessageType.
func (s *ChatRoomLastMessage) GetMessageType() OptMessageType {
	return s.MessageType
}

// GetRoomID returns the value of RoomID.
func (s *ChatRoomLastMessage) GetRoomID() OptNilInt64 {
	return s.RoomID
}

// GetText returns the value of Text.
func (s *ChatRoomLastMessage) GetText() OptNilString {
	return s.Text
}

// GetUserID returns the value of UserID.
func (s *ChatRoomLastMessage) GetUserID() OptNilInt64 {
	return s.UserID
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *ChatRoomLastMessage) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *ChatRoomLastMessage) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetID sets the value of ID.
func (s *ChatRoomLastMessage) SetID(val OptNilInt64) {
	s.ID = val
}

// SetMessageType sets the value of MessageType.
func (s *ChatRoomLastMessage) SetMessageType(val OptMessageType) {
	s.MessageType = val
}

// SetRoomID sets the value of RoomID.
func (s *ChatRoomLastMessage) SetRoomID(val OptNilInt64) {
	s.RoomID = val
}

// SetText sets the value of Text.
func (s *ChatRoomLastMessage) SetText(val OptNilString) {
	s.Text = val
}

// SetUserID sets the value of UserID.
func (s *ChatRoomLastMessage) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/ChatRoomResponse
type ChatRoomResponse struct {
	Chat *RealmChatRoom `json:"chat"`
}

// GetChat returns the value of Chat.
func (s *ChatRoomResponse) GetChat() *RealmChatRoom {
	return s.Chat
}

// SetChat sets the value of Chat.
func (s *ChatRoomResponse) SetChat(val *RealmChatRoom) {
	s.Chat = val
}

// Ref: #/components/schemas/ChatRoomsResponse
type ChatRoomsResponse struct {
	ChatRooms       OptNilRealmChatRoomArray `json:"chat_rooms"`
	NextPageValue   OptNilString             `json:"next_page_value"`
	PinnedChatRooms OptNilRealmChatRoomArray `json:"pinned_chat_rooms"`
}

// GetChatRooms returns the value of ChatRooms.
func (s *ChatRoomsResponse) GetChatRooms() OptNilRealmChatRoomArray {
	return s.ChatRooms
}

// GetNextPageValue returns the value of NextPageValue.
func (s *ChatRoomsResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetPinnedChatRooms returns the value of PinnedChatRooms.
func (s *ChatRoomsResponse) GetPinnedChatRooms() OptNilRealmChatRoomArray {
	return s.PinnedChatRooms
}

// SetChatRooms sets the value of ChatRooms.
func (s *ChatRoomsResponse) SetChatRooms(val OptNilRealmChatRoomArray) {
	s.ChatRooms = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *ChatRoomsResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetPinnedChatRooms sets the value of PinnedChatRooms.
func (s *ChatRoomsResponse) SetPinnedChatRooms(val OptNilRealmChatRoomArray) {
	s.PinnedChatRooms = val
}

// Ref: #/components/schemas/Choice
type Choice struct {
	ID         OptNilInt64  `json:"id"`
	Label      OptNilString `json:"label"`
	VotesCount OptNilInt32  `json:"votes_count"`
}

// GetID returns the value of ID.
func (s *Choice) GetID() OptNilInt64 {
	return s.ID
}

// GetLabel returns the value of Label.
func (s *Choice) GetLabel() OptNilString {
	return s.Label
}

// GetVotesCount returns the value of VotesCount.
func (s *Choice) GetVotesCount() OptNilInt32 {
	return s.VotesCount
}

// SetID sets the value of ID.
func (s *Choice) SetID(val OptNilInt64) {
	s.ID = val
}

// SetLabel sets the value of Label.
func (s *Choice) SetLabel(val OptNilString) {
	s.Label = val
}

// SetVotesCount sets the value of VotesCount.
func (s *Choice) SetVotesCount(val OptNilInt32) {
	s.VotesCount = val
}

// Ref: #/components/schemas/CommonIdsRequest
type CommonIdsRequest struct {
	Ids OptNilInt64Array `json:"ids"`
}

// GetIds returns the value of Ids.
func (s *CommonIdsRequest) GetIds() OptNilInt64Array {
	return s.Ids
}

// SetIds sets the value of Ids.
func (s *CommonIdsRequest) SetIds(val OptNilInt64Array) {
	s.Ids = val
}

// Ref: #/components/schemas/CommonUrlResponse
type CommonUrlResponse struct {
	URL OptNilString `json:"url"`
}

// GetURL returns the value of URL.
func (s *CommonUrlResponse) GetURL() OptNilString {
	return s.URL
}

// SetURL sets the value of URL.
func (s *CommonUrlResponse) SetURL(val OptNilString) {
	s.URL = val
}

// Ref: #/components/schemas/ConferenceCall_BumpParams
type ConferenceCallBumpParams struct {
	ParticipantLimit OptNilInt32 `json:"participant_limit"`
}

// GetParticipantLimit returns the value of ParticipantLimit.
func (s *ConferenceCallBumpParams) GetParticipantLimit() OptNilInt32 {
	return s.ParticipantLimit
}

// SetParticipantLimit sets the value of ParticipantLimit.
func (s *ConferenceCallBumpParams) SetParticipantLimit(val OptNilInt32) {
	s.ParticipantLimit = val
}

// Ref: #/components/schemas/ConferenceCallResponse
type ConferenceCallResponse struct {
	ConferenceCall         OptRealmConferenceCall `json:"conference_call"`
	ConferenceCallUserUUID OptNilString           `json:"conference_call_user_uuid"`
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *ConferenceCallResponse) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetConferenceCallUserUUID returns the value of ConferenceCallUserUUID.
func (s *ConferenceCallResponse) GetConferenceCallUserUUID() OptNilString {
	return s.ConferenceCallUserUUID
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *ConferenceCallResponse) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetConferenceCallUserUUID sets the value of ConferenceCallUserUUID.
func (s *ConferenceCallResponse) SetConferenceCallUserUUID(val OptNilString) {
	s.ConferenceCallUserUUID = val
}

// Ref: #/components/schemas/ConferenceCallUserRole
type ConferenceCallUserRole struct {
	ID     OptNilInt64  `json:"id"`
	Role   OptNilString `json:"role"`
	UserID OptNilInt64  `json:"user_id"`
}

// GetID returns the value of ID.
func (s *ConferenceCallUserRole) GetID() OptNilInt64 {
	return s.ID
}

// GetRole returns the value of Role.
func (s *ConferenceCallUserRole) GetRole() OptNilString {
	return s.Role
}

// GetUserID returns the value of UserID.
func (s *ConferenceCallUserRole) GetUserID() OptNilInt64 {
	return s.UserID
}

// SetID sets the value of ID.
func (s *ConferenceCallUserRole) SetID(val OptNilInt64) {
	s.ID = val
}

// SetRole sets the value of Role.
func (s *ConferenceCallUserRole) SetRole(val OptNilString) {
	s.Role = val
}

// SetUserID sets the value of UserID.
func (s *ConferenceCallUserRole) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/Config
type Config struct {
	AdTesterUserIds                                   OptNilString `json:"ad_tester_user_ids"`
	AndroidNewsVersion                                OptNilString `json:"android_news_version"`
	Dapps                                             OptNilString `json:"dapps"`
	IsChatWebsocketEnabled                            OptNilString `json:"is_chat_websocket_enabled"`
	IsDirectVipPurchaseEnabled                        OptNilString `json:"is_direct_vip_purchase_enabled"`
	IsGiftFeaturesEnabled                             OptNilString `json:"is_gift_features_enabled"`
	IsIDCardAndPhoneVerificationCheckForReviewEnabled OptNilString `json:"is_id_card_and_phone_verification_check_for_review_enabled"`
	IsIDCardCheckRequiredBlockerViewEnabled           OptNilString `json:"is_id_card_check_required_blocker_view_enabled"`
	IsMaintenanced                                    OptNilString `json:"is_maintenanced"`
	IsPhoneVerificationRequiredBlockerViewEnabled     OptNilString `json:"is_phone_verification_required_blocker_view_enabled"`
	IsSpeedTestEnabled                                OptNilString `json:"is_speed_test_enabled"`
	IsSslPinningEnabled                               OptNilString `json:"is_ssl_pinning_enabled"`
	IsStarEnabled                                     OptNilString `json:"is_star_enabled"`
	LatestAndroidAppVersion                           OptNilString `json:"latest_android_app_version"`
	LineOfficialAccountID                             OptNilString `json:"line_official_account_id"`
	LocalizedAndroidNewsTitle                         OptNilString `json:"localized_android_news_title"`
	LocalizedAndroidNewsURL                           OptNilString `json:"localized_android_news_url"`
	LocalizedMaintenanceURL                           OptNilString `json:"localized_maintenance_url"`
	LocalizedNewsTitle                                OptNilString `json:"localized_news_title"`
	LocalizedNewsURL                                  OptNilString `json:"localized_news_url"`
	MaintenanceFeatures                               OptNilString `json:"maintenance_features"`
	MaxImageFrameCount                                OptNilString `json:"max_image_frame_count"`
	MinimumAndroidAppVersionRequired                  OptNilString `json:"minimum_android_app_version_required"`
	MinimumAndroidVersionSupported                    OptNilString `json:"minimum_android_version_supported"`
	NewsVersion                                       OptNilString `json:"news_version"`
	NftInfos                                          OptNilString `json:"nft_infos"`
	PromotionStickerPackIds                           OptNilString `json:"promotion_sticker_pack_ids"`
	ShouldAppendUserIDToNewsURL                       OptNilString `json:"should_append_user_id_to_news_url"`
	SupportEmailAddress                               OptNilString `json:"support_email_address"`
	TokenInfos                                        OptNilString `json:"token_infos"`
	UseRandomMessageRefreshRate                       OptNilString `json:"use_random_message_refresh_rate"`
	Web3IsMaintenanced                                OptNilString `json:"web3_is_maintenanced"`
	Web3LocalizedMaintenanceURL                       OptNilString `json:"web3_localized_maintenance_url"`
	Web3MaintenanceFeatures                           OptNilString `json:"web3_maintenance_features"`
	Web3MaintenanceURL                                OptNilString `json:"web3_maintenance_url"`
	XYayGlobalID                                      OptNilString `json:"x_yay_global_id"`
	XYayJpID                                          OptNilString `json:"x_yay_jp_id"`
	XYayUpdateID                                      OptNilString `json:"x_yay_update_id"`
}

// GetAdTesterUserIds returns the value of AdTesterUserIds.
func (s *Config) GetAdTesterUserIds() OptNilString {
	return s.AdTesterUserIds
}

// GetAndroidNewsVersion returns the value of AndroidNewsVersion.
func (s *Config) GetAndroidNewsVersion() OptNilString {
	return s.AndroidNewsVersion
}

// GetDapps returns the value of Dapps.
func (s *Config) GetDapps() OptNilString {
	return s.Dapps
}

// GetIsChatWebsocketEnabled returns the value of IsChatWebsocketEnabled.
func (s *Config) GetIsChatWebsocketEnabled() OptNilString {
	return s.IsChatWebsocketEnabled
}

// GetIsDirectVipPurchaseEnabled returns the value of IsDirectVipPurchaseEnabled.
func (s *Config) GetIsDirectVipPurchaseEnabled() OptNilString {
	return s.IsDirectVipPurchaseEnabled
}

// GetIsGiftFeaturesEnabled returns the value of IsGiftFeaturesEnabled.
func (s *Config) GetIsGiftFeaturesEnabled() OptNilString {
	return s.IsGiftFeaturesEnabled
}

// GetIsIDCardAndPhoneVerificationCheckForReviewEnabled returns the value of IsIDCardAndPhoneVerificationCheckForReviewEnabled.
func (s *Config) GetIsIDCardAndPhoneVerificationCheckForReviewEnabled() OptNilString {
	return s.IsIDCardAndPhoneVerificationCheckForReviewEnabled
}

// GetIsIDCardCheckRequiredBlockerViewEnabled returns the value of IsIDCardCheckRequiredBlockerViewEnabled.
func (s *Config) GetIsIDCardCheckRequiredBlockerViewEnabled() OptNilString {
	return s.IsIDCardCheckRequiredBlockerViewEnabled
}

// GetIsMaintenanced returns the value of IsMaintenanced.
func (s *Config) GetIsMaintenanced() OptNilString {
	return s.IsMaintenanced
}

// GetIsPhoneVerificationRequiredBlockerViewEnabled returns the value of IsPhoneVerificationRequiredBlockerViewEnabled.
func (s *Config) GetIsPhoneVerificationRequiredBlockerViewEnabled() OptNilString {
	return s.IsPhoneVerificationRequiredBlockerViewEnabled
}

// GetIsSpeedTestEnabled returns the value of IsSpeedTestEnabled.
func (s *Config) GetIsSpeedTestEnabled() OptNilString {
	return s.IsSpeedTestEnabled
}

// GetIsSslPinningEnabled returns the value of IsSslPinningEnabled.
func (s *Config) GetIsSslPinningEnabled() OptNilString {
	return s.IsSslPinningEnabled
}

// GetIsStarEnabled returns the value of IsStarEnabled.
func (s *Config) GetIsStarEnabled() OptNilString {
	return s.IsStarEnabled
}

// GetLatestAndroidAppVersion returns the value of LatestAndroidAppVersion.
func (s *Config) GetLatestAndroidAppVersion() OptNilString {
	return s.LatestAndroidAppVersion
}

// GetLineOfficialAccountID returns the value of LineOfficialAccountID.
func (s *Config) GetLineOfficialAccountID() OptNilString {
	return s.LineOfficialAccountID
}

// GetLocalizedAndroidNewsTitle returns the value of LocalizedAndroidNewsTitle.
func (s *Config) GetLocalizedAndroidNewsTitle() OptNilString {
	return s.LocalizedAndroidNewsTitle
}

// GetLocalizedAndroidNewsURL returns the value of LocalizedAndroidNewsURL.
func (s *Config) GetLocalizedAndroidNewsURL() OptNilString {
	return s.LocalizedAndroidNewsURL
}

// GetLocalizedMaintenanceURL returns the value of LocalizedMaintenanceURL.
func (s *Config) GetLocalizedMaintenanceURL() OptNilString {
	return s.LocalizedMaintenanceURL
}

// GetLocalizedNewsTitle returns the value of LocalizedNewsTitle.
func (s *Config) GetLocalizedNewsTitle() OptNilString {
	return s.LocalizedNewsTitle
}

// GetLocalizedNewsURL returns the value of LocalizedNewsURL.
func (s *Config) GetLocalizedNewsURL() OptNilString {
	return s.LocalizedNewsURL
}

// GetMaintenanceFeatures returns the value of MaintenanceFeatures.
func (s *Config) GetMaintenanceFeatures() OptNilString {
	return s.MaintenanceFeatures
}

// GetMaxImageFrameCount returns the value of MaxImageFrameCount.
func (s *Config) GetMaxImageFrameCount() OptNilString {
	return s.MaxImageFrameCount
}

// GetMinimumAndroidAppVersionRequired returns the value of MinimumAndroidAppVersionRequired.
func (s *Config) GetMinimumAndroidAppVersionRequired() OptNilString {
	return s.MinimumAndroidAppVersionRequired
}

// GetMinimumAndroidVersionSupported returns the value of MinimumAndroidVersionSupported.
func (s *Config) GetMinimumAndroidVersionSupported() OptNilString {
	return s.MinimumAndroidVersionSupported
}

// GetNewsVersion returns the value of NewsVersion.
func (s *Config) GetNewsVersion() OptNilString {
	return s.NewsVersion
}

// GetNftInfos returns the value of NftInfos.
func (s *Config) GetNftInfos() OptNilString {
	return s.NftInfos
}

// GetPromotionStickerPackIds returns the value of PromotionStickerPackIds.
func (s *Config) GetPromotionStickerPackIds() OptNilString {
	return s.PromotionStickerPackIds
}

// GetShouldAppendUserIDToNewsURL returns the value of ShouldAppendUserIDToNewsURL.
func (s *Config) GetShouldAppendUserIDToNewsURL() OptNilString {
	return s.ShouldAppendUserIDToNewsURL
}

// GetSupportEmailAddress returns the value of SupportEmailAddress.
func (s *Config) GetSupportEmailAddress() OptNilString {
	return s.SupportEmailAddress
}

// GetTokenInfos returns the value of TokenInfos.
func (s *Config) GetTokenInfos() OptNilString {
	return s.TokenInfos
}

// GetUseRandomMessageRefreshRate returns the value of UseRandomMessageRefreshRate.
func (s *Config) GetUseRandomMessageRefreshRate() OptNilString {
	return s.UseRandomMessageRefreshRate
}

// GetWeb3IsMaintenanced returns the value of Web3IsMaintenanced.
func (s *Config) GetWeb3IsMaintenanced() OptNilString {
	return s.Web3IsMaintenanced
}

// GetWeb3LocalizedMaintenanceURL returns the value of Web3LocalizedMaintenanceURL.
func (s *Config) GetWeb3LocalizedMaintenanceURL() OptNilString {
	return s.Web3LocalizedMaintenanceURL
}

// GetWeb3MaintenanceFeatures returns the value of Web3MaintenanceFeatures.
func (s *Config) GetWeb3MaintenanceFeatures() OptNilString {
	return s.Web3MaintenanceFeatures
}

// GetWeb3MaintenanceURL returns the value of Web3MaintenanceURL.
func (s *Config) GetWeb3MaintenanceURL() OptNilString {
	return s.Web3MaintenanceURL
}

// GetXYayGlobalID returns the value of XYayGlobalID.
func (s *Config) GetXYayGlobalID() OptNilString {
	return s.XYayGlobalID
}

// GetXYayJpID returns the value of XYayJpID.
func (s *Config) GetXYayJpID() OptNilString {
	return s.XYayJpID
}

// GetXYayUpdateID returns the value of XYayUpdateID.
func (s *Config) GetXYayUpdateID() OptNilString {
	return s.XYayUpdateID
}

// SetAdTesterUserIds sets the value of AdTesterUserIds.
func (s *Config) SetAdTesterUserIds(val OptNilString) {
	s.AdTesterUserIds = val
}

// SetAndroidNewsVersion sets the value of AndroidNewsVersion.
func (s *Config) SetAndroidNewsVersion(val OptNilString) {
	s.AndroidNewsVersion = val
}

// SetDapps sets the value of Dapps.
func (s *Config) SetDapps(val OptNilString) {
	s.Dapps = val
}

// SetIsChatWebsocketEnabled sets the value of IsChatWebsocketEnabled.
func (s *Config) SetIsChatWebsocketEnabled(val OptNilString) {
	s.IsChatWebsocketEnabled = val
}

// SetIsDirectVipPurchaseEnabled sets the value of IsDirectVipPurchaseEnabled.
func (s *Config) SetIsDirectVipPurchaseEnabled(val OptNilString) {
	s.IsDirectVipPurchaseEnabled = val
}

// SetIsGiftFeaturesEnabled sets the value of IsGiftFeaturesEnabled.
func (s *Config) SetIsGiftFeaturesEnabled(val OptNilString) {
	s.IsGiftFeaturesEnabled = val
}

// SetIsIDCardAndPhoneVerificationCheckForReviewEnabled sets the value of IsIDCardAndPhoneVerificationCheckForReviewEnabled.
func (s *Config) SetIsIDCardAndPhoneVerificationCheckForReviewEnabled(val OptNilString) {
	s.IsIDCardAndPhoneVerificationCheckForReviewEnabled = val
}

// SetIsIDCardCheckRequiredBlockerViewEnabled sets the value of IsIDCardCheckRequiredBlockerViewEnabled.
func (s *Config) SetIsIDCardCheckRequiredBlockerViewEnabled(val OptNilString) {
	s.IsIDCardCheckRequiredBlockerViewEnabled = val
}

// SetIsMaintenanced sets the value of IsMaintenanced.
func (s *Config) SetIsMaintenanced(val OptNilString) {
	s.IsMaintenanced = val
}

// SetIsPhoneVerificationRequiredBlockerViewEnabled sets the value of IsPhoneVerificationRequiredBlockerViewEnabled.
func (s *Config) SetIsPhoneVerificationRequiredBlockerViewEnabled(val OptNilString) {
	s.IsPhoneVerificationRequiredBlockerViewEnabled = val
}

// SetIsSpeedTestEnabled sets the value of IsSpeedTestEnabled.
func (s *Config) SetIsSpeedTestEnabled(val OptNilString) {
	s.IsSpeedTestEnabled = val
}

// SetIsSslPinningEnabled sets the value of IsSslPinningEnabled.
func (s *Config) SetIsSslPinningEnabled(val OptNilString) {
	s.IsSslPinningEnabled = val
}

// SetIsStarEnabled sets the value of IsStarEnabled.
func (s *Config) SetIsStarEnabled(val OptNilString) {
	s.IsStarEnabled = val
}

// SetLatestAndroidAppVersion sets the value of LatestAndroidAppVersion.
func (s *Config) SetLatestAndroidAppVersion(val OptNilString) {
	s.LatestAndroidAppVersion = val
}

// SetLineOfficialAccountID sets the value of LineOfficialAccountID.
func (s *Config) SetLineOfficialAccountID(val OptNilString) {
	s.LineOfficialAccountID = val
}

// SetLocalizedAndroidNewsTitle sets the value of LocalizedAndroidNewsTitle.
func (s *Config) SetLocalizedAndroidNewsTitle(val OptNilString) {
	s.LocalizedAndroidNewsTitle = val
}

// SetLocalizedAndroidNewsURL sets the value of LocalizedAndroidNewsURL.
func (s *Config) SetLocalizedAndroidNewsURL(val OptNilString) {
	s.LocalizedAndroidNewsURL = val
}

// SetLocalizedMaintenanceURL sets the value of LocalizedMaintenanceURL.
func (s *Config) SetLocalizedMaintenanceURL(val OptNilString) {
	s.LocalizedMaintenanceURL = val
}

// SetLocalizedNewsTitle sets the value of LocalizedNewsTitle.
func (s *Config) SetLocalizedNewsTitle(val OptNilString) {
	s.LocalizedNewsTitle = val
}

// SetLocalizedNewsURL sets the value of LocalizedNewsURL.
func (s *Config) SetLocalizedNewsURL(val OptNilString) {
	s.LocalizedNewsURL = val
}

// SetMaintenanceFeatures sets the value of MaintenanceFeatures.
func (s *Config) SetMaintenanceFeatures(val OptNilString) {
	s.MaintenanceFeatures = val
}

// SetMaxImageFrameCount sets the value of MaxImageFrameCount.
func (s *Config) SetMaxImageFrameCount(val OptNilString) {
	s.MaxImageFrameCount = val
}

// SetMinimumAndroidAppVersionRequired sets the value of MinimumAndroidAppVersionRequired.
func (s *Config) SetMinimumAndroidAppVersionRequired(val OptNilString) {
	s.MinimumAndroidAppVersionRequired = val
}

// SetMinimumAndroidVersionSupported sets the value of MinimumAndroidVersionSupported.
func (s *Config) SetMinimumAndroidVersionSupported(val OptNilString) {
	s.MinimumAndroidVersionSupported = val
}

// SetNewsVersion sets the value of NewsVersion.
func (s *Config) SetNewsVersion(val OptNilString) {
	s.NewsVersion = val
}

// SetNftInfos sets the value of NftInfos.
func (s *Config) SetNftInfos(val OptNilString) {
	s.NftInfos = val
}

// SetPromotionStickerPackIds sets the value of PromotionStickerPackIds.
func (s *Config) SetPromotionStickerPackIds(val OptNilString) {
	s.PromotionStickerPackIds = val
}

// SetShouldAppendUserIDToNewsURL sets the value of ShouldAppendUserIDToNewsURL.
func (s *Config) SetShouldAppendUserIDToNewsURL(val OptNilString) {
	s.ShouldAppendUserIDToNewsURL = val
}

// SetSupportEmailAddress sets the value of SupportEmailAddress.
func (s *Config) SetSupportEmailAddress(val OptNilString) {
	s.SupportEmailAddress = val
}

// SetTokenInfos sets the value of TokenInfos.
func (s *Config) SetTokenInfos(val OptNilString) {
	s.TokenInfos = val
}

// SetUseRandomMessageRefreshRate sets the value of UseRandomMessageRefreshRate.
func (s *Config) SetUseRandomMessageRefreshRate(val OptNilString) {
	s.UseRandomMessageRefreshRate = val
}

// SetWeb3IsMaintenanced sets the value of Web3IsMaintenanced.
func (s *Config) SetWeb3IsMaintenanced(val OptNilString) {
	s.Web3IsMaintenanced = val
}

// SetWeb3LocalizedMaintenanceURL sets the value of Web3LocalizedMaintenanceURL.
func (s *Config) SetWeb3LocalizedMaintenanceURL(val OptNilString) {
	s.Web3LocalizedMaintenanceURL = val
}

// SetWeb3MaintenanceFeatures sets the value of Web3MaintenanceFeatures.
func (s *Config) SetWeb3MaintenanceFeatures(val OptNilString) {
	s.Web3MaintenanceFeatures = val
}

// SetWeb3MaintenanceURL sets the value of Web3MaintenanceURL.
func (s *Config) SetWeb3MaintenanceURL(val OptNilString) {
	s.Web3MaintenanceURL = val
}

// SetXYayGlobalID sets the value of XYayGlobalID.
func (s *Config) SetXYayGlobalID(val OptNilString) {
	s.XYayGlobalID = val
}

// SetXYayJpID sets the value of XYayJpID.
func (s *Config) SetXYayJpID(val OptNilString) {
	s.XYayJpID = val
}

// SetXYayUpdateID sets the value of XYayUpdateID.
func (s *Config) SetXYayUpdateID(val OptNilString) {
	s.XYayUpdateID = val
}

type CreateChatRoomReq struct {
	BackgroundFilename OptNilString `json:"background_filename"`
	IconFilename       OptNilString `json:"icon_filename"`
	Name               OptString    `json:"name"`
	WithUserIds        []int64      `json:"with_user_ids[]"`
}

// GetBackgroundFilename returns the value of BackgroundFilename.
func (s *CreateChatRoomReq) GetBackgroundFilename() OptNilString {
	return s.BackgroundFilename
}

// GetIconFilename returns the value of IconFilename.
func (s *CreateChatRoomReq) GetIconFilename() OptNilString {
	return s.IconFilename
}

// GetName returns the value of Name.
func (s *CreateChatRoomReq) GetName() OptString {
	return s.Name
}

// GetWithUserIds returns the value of WithUserIds.
func (s *CreateChatRoomReq) GetWithUserIds() []int64 {
	return s.WithUserIds
}

// SetBackgroundFilename sets the value of BackgroundFilename.
func (s *CreateChatRoomReq) SetBackgroundFilename(val OptNilString) {
	s.BackgroundFilename = val
}

// SetIconFilename sets the value of IconFilename.
func (s *CreateChatRoomReq) SetIconFilename(val OptNilString) {
	s.IconFilename = val
}

// SetName sets the value of Name.
func (s *CreateChatRoomReq) SetName(val OptString) {
	s.Name = val
}

// SetWithUserIds sets the value of WithUserIds.
func (s *CreateChatRoomReq) SetWithUserIds(val []int64) {
	s.WithUserIds = val
}

// Ref: #/components/schemas/CreateChatRoomResponse
type CreateChatRoomResponse struct {
	RoomID OptNilInt64 `json:"room_id"`
}

// GetRoomID returns the value of RoomID.
func (s *CreateChatRoomResponse) GetRoomID() OptNilInt64 {
	return s.RoomID
}

// SetRoomID sets the value of RoomID.
func (s *CreateChatRoomResponse) SetRoomID(val OptNilInt64) {
	s.RoomID = val
}

type CreateChatRoomV1Req struct {
	HimaChat   OptBool     `json:"hima_chat"`
	MatchingID OptNilInt64 `json:"matching_id"`
	WithUserID OptInt64    `json:"with_user_id"`
}

// GetHimaChat returns the value of HimaChat.
func (s *CreateChatRoomV1Req) GetHimaChat() OptBool {
	return s.HimaChat
}

// GetMatchingID returns the value of MatchingID.
func (s *CreateChatRoomV1Req) GetMatchingID() OptNilInt64 {
	return s.MatchingID
}

// GetWithUserID returns the value of WithUserID.
func (s *CreateChatRoomV1Req) GetWithUserID() OptInt64 {
	return s.WithUserID
}

// SetHimaChat sets the value of HimaChat.
func (s *CreateChatRoomV1Req) SetHimaChat(val OptBool) {
	s.HimaChat = val
}

// SetMatchingID sets the value of MatchingID.
func (s *CreateChatRoomV1Req) SetMatchingID(val OptNilInt64) {
	s.MatchingID = val
}

// SetWithUserID sets the value of WithUserID.
func (s *CreateChatRoomV1Req) SetWithUserID(val OptInt64) {
	s.WithUserID = val
}

type CreateConferenceCallPostReq struct {
	APIKey              OptString                                 `json:"api_key"`
	Attachment2Filename OptNilString                              `json:"attachment_2_filename"`
	Attachment3Filename OptNilString                              `json:"attachment_3_filename"`
	Attachment4Filename OptNilString                              `json:"attachment_4_filename"`
	Attachment5Filename OptNilString                              `json:"attachment_5_filename"`
	Attachment6Filename OptNilString                              `json:"attachment_6_filename"`
	Attachment7Filename OptNilString                              `json:"attachment_7_filename"`
	Attachment8Filename OptNilString                              `json:"attachment_8_filename"`
	Attachment9Filename OptNilString                              `json:"attachment_9_filename"`
	AttachmentFilename  OptNilString                              `json:"attachment_filename"`
	CallType            OptNilString                              `json:"call_type"`
	CategoryID          OptNilInt64                               `json:"category_id"`
	Color               OptNilInt32                               `json:"color"`
	FontSize            OptNilInt32                               `json:"font_size"`
	GameTitle           OptNilString                              `json:"game_title"`
	GroupID             OptNilInt64                               `json:"group_id"`
	JoinableBy          OptNilString                              `json:"joinable_by"`
	Language            OptNilString                              `json:"language"`
	MessageTags         OptCreateConferenceCallPostReqMessageTags `json:"message_tags"`
	SignedInfo          OptString                                 `json:"signed_info"`
	Text                OptNilString                              `json:"text"`
	Timestamp           OptInt64                                  `json:"timestamp"`
	UUID                OptString                                 `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *CreateConferenceCallPostReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetAttachment2Filename returns the value of Attachment2Filename.
func (s *CreateConferenceCallPostReq) GetAttachment2Filename() OptNilString {
	return s.Attachment2Filename
}

// GetAttachment3Filename returns the value of Attachment3Filename.
func (s *CreateConferenceCallPostReq) GetAttachment3Filename() OptNilString {
	return s.Attachment3Filename
}

// GetAttachment4Filename returns the value of Attachment4Filename.
func (s *CreateConferenceCallPostReq) GetAttachment4Filename() OptNilString {
	return s.Attachment4Filename
}

// GetAttachment5Filename returns the value of Attachment5Filename.
func (s *CreateConferenceCallPostReq) GetAttachment5Filename() OptNilString {
	return s.Attachment5Filename
}

// GetAttachment6Filename returns the value of Attachment6Filename.
func (s *CreateConferenceCallPostReq) GetAttachment6Filename() OptNilString {
	return s.Attachment6Filename
}

// GetAttachment7Filename returns the value of Attachment7Filename.
func (s *CreateConferenceCallPostReq) GetAttachment7Filename() OptNilString {
	return s.Attachment7Filename
}

// GetAttachment8Filename returns the value of Attachment8Filename.
func (s *CreateConferenceCallPostReq) GetAttachment8Filename() OptNilString {
	return s.Attachment8Filename
}

// GetAttachment9Filename returns the value of Attachment9Filename.
func (s *CreateConferenceCallPostReq) GetAttachment9Filename() OptNilString {
	return s.Attachment9Filename
}

// GetAttachmentFilename returns the value of AttachmentFilename.
func (s *CreateConferenceCallPostReq) GetAttachmentFilename() OptNilString {
	return s.AttachmentFilename
}

// GetCallType returns the value of CallType.
func (s *CreateConferenceCallPostReq) GetCallType() OptNilString {
	return s.CallType
}

// GetCategoryID returns the value of CategoryID.
func (s *CreateConferenceCallPostReq) GetCategoryID() OptNilInt64 {
	return s.CategoryID
}

// GetColor returns the value of Color.
func (s *CreateConferenceCallPostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *CreateConferenceCallPostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGameTitle returns the value of GameTitle.
func (s *CreateConferenceCallPostReq) GetGameTitle() OptNilString {
	return s.GameTitle
}

// GetGroupID returns the value of GroupID.
func (s *CreateConferenceCallPostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetJoinableBy returns the value of JoinableBy.
func (s *CreateConferenceCallPostReq) GetJoinableBy() OptNilString {
	return s.JoinableBy
}

// GetLanguage returns the value of Language.
func (s *CreateConferenceCallPostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMessageTags returns the value of MessageTags.
func (s *CreateConferenceCallPostReq) GetMessageTags() OptCreateConferenceCallPostReqMessageTags {
	return s.MessageTags
}

// GetSignedInfo returns the value of SignedInfo.
func (s *CreateConferenceCallPostReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetText returns the value of Text.
func (s *CreateConferenceCallPostReq) GetText() OptNilString {
	return s.Text
}

// GetTimestamp returns the value of Timestamp.
func (s *CreateConferenceCallPostReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *CreateConferenceCallPostReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *CreateConferenceCallPostReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetAttachment2Filename sets the value of Attachment2Filename.
func (s *CreateConferenceCallPostReq) SetAttachment2Filename(val OptNilString) {
	s.Attachment2Filename = val
}

// SetAttachment3Filename sets the value of Attachment3Filename.
func (s *CreateConferenceCallPostReq) SetAttachment3Filename(val OptNilString) {
	s.Attachment3Filename = val
}

// SetAttachment4Filename sets the value of Attachment4Filename.
func (s *CreateConferenceCallPostReq) SetAttachment4Filename(val OptNilString) {
	s.Attachment4Filename = val
}

// SetAttachment5Filename sets the value of Attachment5Filename.
func (s *CreateConferenceCallPostReq) SetAttachment5Filename(val OptNilString) {
	s.Attachment5Filename = val
}

// SetAttachment6Filename sets the value of Attachment6Filename.
func (s *CreateConferenceCallPostReq) SetAttachment6Filename(val OptNilString) {
	s.Attachment6Filename = val
}

// SetAttachment7Filename sets the value of Attachment7Filename.
func (s *CreateConferenceCallPostReq) SetAttachment7Filename(val OptNilString) {
	s.Attachment7Filename = val
}

// SetAttachment8Filename sets the value of Attachment8Filename.
func (s *CreateConferenceCallPostReq) SetAttachment8Filename(val OptNilString) {
	s.Attachment8Filename = val
}

// SetAttachment9Filename sets the value of Attachment9Filename.
func (s *CreateConferenceCallPostReq) SetAttachment9Filename(val OptNilString) {
	s.Attachment9Filename = val
}

// SetAttachmentFilename sets the value of AttachmentFilename.
func (s *CreateConferenceCallPostReq) SetAttachmentFilename(val OptNilString) {
	s.AttachmentFilename = val
}

// SetCallType sets the value of CallType.
func (s *CreateConferenceCallPostReq) SetCallType(val OptNilString) {
	s.CallType = val
}

// SetCategoryID sets the value of CategoryID.
func (s *CreateConferenceCallPostReq) SetCategoryID(val OptNilInt64) {
	s.CategoryID = val
}

// SetColor sets the value of Color.
func (s *CreateConferenceCallPostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *CreateConferenceCallPostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGameTitle sets the value of GameTitle.
func (s *CreateConferenceCallPostReq) SetGameTitle(val OptNilString) {
	s.GameTitle = val
}

// SetGroupID sets the value of GroupID.
func (s *CreateConferenceCallPostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetJoinableBy sets the value of JoinableBy.
func (s *CreateConferenceCallPostReq) SetJoinableBy(val OptNilString) {
	s.JoinableBy = val
}

// SetLanguage sets the value of Language.
func (s *CreateConferenceCallPostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMessageTags sets the value of MessageTags.
func (s *CreateConferenceCallPostReq) SetMessageTags(val OptCreateConferenceCallPostReqMessageTags) {
	s.MessageTags = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *CreateConferenceCallPostReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetText sets the value of Text.
func (s *CreateConferenceCallPostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetTimestamp sets the value of Timestamp.
func (s *CreateConferenceCallPostReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *CreateConferenceCallPostReq) SetUUID(val OptString) {
	s.UUID = val
}

type CreateConferenceCallPostReqMessageTags struct{}

// Ref: #/components/schemas/CreateGroupQuota
type CreateGroupQuota struct {
	RemainingQuota OptNilInt32 `json:"remaining_quota"`
	UsedQuota      OptNilInt32 `json:"used_quota"`
}

// GetRemainingQuota returns the value of RemainingQuota.
func (s *CreateGroupQuota) GetRemainingQuota() OptNilInt32 {
	return s.RemainingQuota
}

// GetUsedQuota returns the value of UsedQuota.
func (s *CreateGroupQuota) GetUsedQuota() OptNilInt32 {
	return s.UsedQuota
}

// SetRemainingQuota sets the value of RemainingQuota.
func (s *CreateGroupQuota) SetRemainingQuota(val OptNilInt32) {
	s.RemainingQuota = val
}

// SetUsedQuota sets the value of UsedQuota.
func (s *CreateGroupQuota) SetUsedQuota(val OptNilInt32) {
	s.UsedQuota = val
}

type CreateGroupReq struct {
	AllowMembersToPostImageAndVideo OptNilBool   `json:"allow_members_to_post_image_and_video"`
	AllowMembersToPostURL           OptNilBool   `json:"allow_members_to_post_url"`
	AllowOwnershipTransfer          OptNilBool   `json:"allow_ownership_transfer"`
	AllowThreadCreationBy           OptNilString `json:"allow_thread_creation_by"`
	APIKey                          OptString    `json:"api_key"`
	CallTimelineDisplay             OptNilBool   `json:"call_timeline_display"`
	CoverImageFilename              OptNilString `json:"cover_image_filename"`
	Description                     OptNilString `json:"description"`
	Gender                          OptNilInt32  `json:"gender"`
	GenerationGroupsLimit           OptNilInt32  `json:"generation_groups_limit"`
	GroupCategoryID                 OptNilInt64  `json:"group_category_id"`
	GroupIconFilename               OptNilString `json:"group_icon_filename"`
	Guidelines                      OptNilString `json:"guidelines"`
	HideConferenceCall              OptNilBool   `json:"hide_conference_call"`
	HideFromGameEight               OptBool      `json:"hide_from_game_eight"`
	HideReportedPosts               OptNilBool   `json:"hide_reported_posts"`
	IsPrivate                       OptNilBool   `json:"is_private"`
	OnlyMobileVerified              OptNilBool   `json:"only_mobile_verified"`
	OnlyVerifiedAge                 OptNilBool   `json:"only_verified_age"`
	Secret                          OptNilBool   `json:"secret"`
	SignedInfo                      OptString    `json:"signed_info"`
	SubCategoryID                   OptNilString `json:"sub_category_id"`
	Timestamp                       OptInt64     `json:"timestamp"`
	Topic                           OptString    `json:"topic"`
	UUID                            OptString    `json:"uuid"`
}

// GetAllowMembersToPostImageAndVideo returns the value of AllowMembersToPostImageAndVideo.
func (s *CreateGroupReq) GetAllowMembersToPostImageAndVideo() OptNilBool {
	return s.AllowMembersToPostImageAndVideo
}

// GetAllowMembersToPostURL returns the value of AllowMembersToPostURL.
func (s *CreateGroupReq) GetAllowMembersToPostURL() OptNilBool {
	return s.AllowMembersToPostURL
}

// GetAllowOwnershipTransfer returns the value of AllowOwnershipTransfer.
func (s *CreateGroupReq) GetAllowOwnershipTransfer() OptNilBool {
	return s.AllowOwnershipTransfer
}

// GetAllowThreadCreationBy returns the value of AllowThreadCreationBy.
func (s *CreateGroupReq) GetAllowThreadCreationBy() OptNilString {
	return s.AllowThreadCreationBy
}

// GetAPIKey returns the value of APIKey.
func (s *CreateGroupReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetCallTimelineDisplay returns the value of CallTimelineDisplay.
func (s *CreateGroupReq) GetCallTimelineDisplay() OptNilBool {
	return s.CallTimelineDisplay
}

// GetCoverImageFilename returns the value of CoverImageFilename.
func (s *CreateGroupReq) GetCoverImageFilename() OptNilString {
	return s.CoverImageFilename
}

// GetDescription returns the value of Description.
func (s *CreateGroupReq) GetDescription() OptNilString {
	return s.Description
}

// GetGender returns the value of Gender.
func (s *CreateGroupReq) GetGender() OptNilInt32 {
	return s.Gender
}

// GetGenerationGroupsLimit returns the value of GenerationGroupsLimit.
func (s *CreateGroupReq) GetGenerationGroupsLimit() OptNilInt32 {
	return s.GenerationGroupsLimit
}

// GetGroupCategoryID returns the value of GroupCategoryID.
func (s *CreateGroupReq) GetGroupCategoryID() OptNilInt64 {
	return s.GroupCategoryID
}

// GetGroupIconFilename returns the value of GroupIconFilename.
func (s *CreateGroupReq) GetGroupIconFilename() OptNilString {
	return s.GroupIconFilename
}

// GetGuidelines returns the value of Guidelines.
func (s *CreateGroupReq) GetGuidelines() OptNilString {
	return s.Guidelines
}

// GetHideConferenceCall returns the value of HideConferenceCall.
func (s *CreateGroupReq) GetHideConferenceCall() OptNilBool {
	return s.HideConferenceCall
}

// GetHideFromGameEight returns the value of HideFromGameEight.
func (s *CreateGroupReq) GetHideFromGameEight() OptBool {
	return s.HideFromGameEight
}

// GetHideReportedPosts returns the value of HideReportedPosts.
func (s *CreateGroupReq) GetHideReportedPosts() OptNilBool {
	return s.HideReportedPosts
}

// GetIsPrivate returns the value of IsPrivate.
func (s *CreateGroupReq) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetOnlyMobileVerified returns the value of OnlyMobileVerified.
func (s *CreateGroupReq) GetOnlyMobileVerified() OptNilBool {
	return s.OnlyMobileVerified
}

// GetOnlyVerifiedAge returns the value of OnlyVerifiedAge.
func (s *CreateGroupReq) GetOnlyVerifiedAge() OptNilBool {
	return s.OnlyVerifiedAge
}

// GetSecret returns the value of Secret.
func (s *CreateGroupReq) GetSecret() OptNilBool {
	return s.Secret
}

// GetSignedInfo returns the value of SignedInfo.
func (s *CreateGroupReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetSubCategoryID returns the value of SubCategoryID.
func (s *CreateGroupReq) GetSubCategoryID() OptNilString {
	return s.SubCategoryID
}

// GetTimestamp returns the value of Timestamp.
func (s *CreateGroupReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetTopic returns the value of Topic.
func (s *CreateGroupReq) GetTopic() OptString {
	return s.Topic
}

// GetUUID returns the value of UUID.
func (s *CreateGroupReq) GetUUID() OptString {
	return s.UUID
}

// SetAllowMembersToPostImageAndVideo sets the value of AllowMembersToPostImageAndVideo.
func (s *CreateGroupReq) SetAllowMembersToPostImageAndVideo(val OptNilBool) {
	s.AllowMembersToPostImageAndVideo = val
}

// SetAllowMembersToPostURL sets the value of AllowMembersToPostURL.
func (s *CreateGroupReq) SetAllowMembersToPostURL(val OptNilBool) {
	s.AllowMembersToPostURL = val
}

// SetAllowOwnershipTransfer sets the value of AllowOwnershipTransfer.
func (s *CreateGroupReq) SetAllowOwnershipTransfer(val OptNilBool) {
	s.AllowOwnershipTransfer = val
}

// SetAllowThreadCreationBy sets the value of AllowThreadCreationBy.
func (s *CreateGroupReq) SetAllowThreadCreationBy(val OptNilString) {
	s.AllowThreadCreationBy = val
}

// SetAPIKey sets the value of APIKey.
func (s *CreateGroupReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetCallTimelineDisplay sets the value of CallTimelineDisplay.
func (s *CreateGroupReq) SetCallTimelineDisplay(val OptNilBool) {
	s.CallTimelineDisplay = val
}

// SetCoverImageFilename sets the value of CoverImageFilename.
func (s *CreateGroupReq) SetCoverImageFilename(val OptNilString) {
	s.CoverImageFilename = val
}

// SetDescription sets the value of Description.
func (s *CreateGroupReq) SetDescription(val OptNilString) {
	s.Description = val
}

// SetGender sets the value of Gender.
func (s *CreateGroupReq) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetGenerationGroupsLimit sets the value of GenerationGroupsLimit.
func (s *CreateGroupReq) SetGenerationGroupsLimit(val OptNilInt32) {
	s.GenerationGroupsLimit = val
}

// SetGroupCategoryID sets the value of GroupCategoryID.
func (s *CreateGroupReq) SetGroupCategoryID(val OptNilInt64) {
	s.GroupCategoryID = val
}

// SetGroupIconFilename sets the value of GroupIconFilename.
func (s *CreateGroupReq) SetGroupIconFilename(val OptNilString) {
	s.GroupIconFilename = val
}

// SetGuidelines sets the value of Guidelines.
func (s *CreateGroupReq) SetGuidelines(val OptNilString) {
	s.Guidelines = val
}

// SetHideConferenceCall sets the value of HideConferenceCall.
func (s *CreateGroupReq) SetHideConferenceCall(val OptNilBool) {
	s.HideConferenceCall = val
}

// SetHideFromGameEight sets the value of HideFromGameEight.
func (s *CreateGroupReq) SetHideFromGameEight(val OptBool) {
	s.HideFromGameEight = val
}

// SetHideReportedPosts sets the value of HideReportedPosts.
func (s *CreateGroupReq) SetHideReportedPosts(val OptNilBool) {
	s.HideReportedPosts = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *CreateGroupReq) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetOnlyMobileVerified sets the value of OnlyMobileVerified.
func (s *CreateGroupReq) SetOnlyMobileVerified(val OptNilBool) {
	s.OnlyMobileVerified = val
}

// SetOnlyVerifiedAge sets the value of OnlyVerifiedAge.
func (s *CreateGroupReq) SetOnlyVerifiedAge(val OptNilBool) {
	s.OnlyVerifiedAge = val
}

// SetSecret sets the value of Secret.
func (s *CreateGroupReq) SetSecret(val OptNilBool) {
	s.Secret = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *CreateGroupReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetSubCategoryID sets the value of SubCategoryID.
func (s *CreateGroupReq) SetSubCategoryID(val OptNilString) {
	s.SubCategoryID = val
}

// SetTimestamp sets the value of Timestamp.
func (s *CreateGroupReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetTopic sets the value of Topic.
func (s *CreateGroupReq) SetTopic(val OptString) {
	s.Topic = val
}

// SetUUID sets the value of UUID.
func (s *CreateGroupReq) SetUUID(val OptString) {
	s.UUID = val
}

// Ref: #/components/schemas/CreateGroupResponse
type CreateGroupResponse struct {
	GroupID OptNilInt64 `json:"group_id"`
}

// GetGroupID returns the value of GroupID.
func (s *CreateGroupResponse) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// SetGroupID sets the value of GroupID.
func (s *CreateGroupResponse) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// Ref: #/components/schemas/CreateGroupThreadRequest
type CreateGroupThreadRequest struct {
	GroupID            OptNilInt64  `json:"group_id"`
	ThreadIconFilename OptNilString `json:"thread_icon_filename"`
	Title              OptNilString `json:"title"`
}

// GetGroupID returns the value of GroupID.
func (s *CreateGroupThreadRequest) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetThreadIconFilename returns the value of ThreadIconFilename.
func (s *CreateGroupThreadRequest) GetThreadIconFilename() OptNilString {
	return s.ThreadIconFilename
}

// GetTitle returns the value of Title.
func (s *CreateGroupThreadRequest) GetTitle() OptNilString {
	return s.Title
}

// SetGroupID sets the value of GroupID.
func (s *CreateGroupThreadRequest) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetThreadIconFilename sets the value of ThreadIconFilename.
func (s *CreateGroupThreadRequest) SetThreadIconFilename(val OptNilString) {
	s.ThreadIconFilename = val
}

// SetTitle sets the value of Title.
func (s *CreateGroupThreadRequest) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/CreateMuteKeywordResponse
type CreateMuteKeywordResponse struct {
	HiddenWord OptMuteKeyword `json:"hidden_word"`
}

// GetHiddenWord returns the value of HiddenWord.
func (s *CreateMuteKeywordResponse) GetHiddenWord() OptMuteKeyword {
	return s.HiddenWord
}

// SetHiddenWord sets the value of HiddenWord.
func (s *CreateMuteKeywordResponse) SetHiddenWord(val OptMuteKeyword) {
	s.HiddenWord = val
}

type CreatePostReq struct {
	Attachment2Filename OptNilString                `json:"attachment_2_filename"`
	Attachment3Filename OptNilString                `json:"attachment_3_filename"`
	Attachment4Filename OptNilString                `json:"attachment_4_filename"`
	Attachment5Filename OptNilString                `json:"attachment_5_filename"`
	Attachment6Filename OptNilString                `json:"attachment_6_filename"`
	Attachment7Filename OptNilString                `json:"attachment_7_filename"`
	Attachment8Filename OptNilString                `json:"attachment_8_filename"`
	Attachment9Filename OptNilString                `json:"attachment_9_filename"`
	AttachmentFilename  OptNilString                `json:"attachment_filename"`
	Choices             OptNilStringArray           `json:"choices[]"`
	Color               OptNilInt32                 `json:"color"`
	FontSize            OptNilInt32                 `json:"font_size"`
	GroupID             OptNilInt64                 `json:"group_id"`
	InReplyTo           OptNilInt64                 `json:"in_reply_to"`
	Language            OptNilString                `json:"language"`
	MentionIds          OptNilInt64Array            `json:"mention_ids[]"`
	MessageTags         OptCreatePostReqMessageTags `json:"message_tags"`
	PostType            OptPostType                 `json:"post_type"`
	SharedURL           OptCreatePostReqSharedURL   `json:"shared_url"`
	Text                OptNilString                `json:"text"`
	VideoFileName       OptNilString                `json:"video_file_name"`
}

// GetAttachment2Filename returns the value of Attachment2Filename.
func (s *CreatePostReq) GetAttachment2Filename() OptNilString {
	return s.Attachment2Filename
}

// GetAttachment3Filename returns the value of Attachment3Filename.
func (s *CreatePostReq) GetAttachment3Filename() OptNilString {
	return s.Attachment3Filename
}

// GetAttachment4Filename returns the value of Attachment4Filename.
func (s *CreatePostReq) GetAttachment4Filename() OptNilString {
	return s.Attachment4Filename
}

// GetAttachment5Filename returns the value of Attachment5Filename.
func (s *CreatePostReq) GetAttachment5Filename() OptNilString {
	return s.Attachment5Filename
}

// GetAttachment6Filename returns the value of Attachment6Filename.
func (s *CreatePostReq) GetAttachment6Filename() OptNilString {
	return s.Attachment6Filename
}

// GetAttachment7Filename returns the value of Attachment7Filename.
func (s *CreatePostReq) GetAttachment7Filename() OptNilString {
	return s.Attachment7Filename
}

// GetAttachment8Filename returns the value of Attachment8Filename.
func (s *CreatePostReq) GetAttachment8Filename() OptNilString {
	return s.Attachment8Filename
}

// GetAttachment9Filename returns the value of Attachment9Filename.
func (s *CreatePostReq) GetAttachment9Filename() OptNilString {
	return s.Attachment9Filename
}

// GetAttachmentFilename returns the value of AttachmentFilename.
func (s *CreatePostReq) GetAttachmentFilename() OptNilString {
	return s.AttachmentFilename
}

// GetChoices returns the value of Choices.
func (s *CreatePostReq) GetChoices() OptNilStringArray {
	return s.Choices
}

// GetColor returns the value of Color.
func (s *CreatePostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *CreatePostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGroupID returns the value of GroupID.
func (s *CreatePostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetInReplyTo returns the value of InReplyTo.
func (s *CreatePostReq) GetInReplyTo() OptNilInt64 {
	return s.InReplyTo
}

// GetLanguage returns the value of Language.
func (s *CreatePostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMentionIds returns the value of MentionIds.
func (s *CreatePostReq) GetMentionIds() OptNilInt64Array {
	return s.MentionIds
}

// GetMessageTags returns the value of MessageTags.
func (s *CreatePostReq) GetMessageTags() OptCreatePostReqMessageTags {
	return s.MessageTags
}

// GetPostType returns the value of PostType.
func (s *CreatePostReq) GetPostType() OptPostType {
	return s.PostType
}

// GetSharedURL returns the value of SharedURL.
func (s *CreatePostReq) GetSharedURL() OptCreatePostReqSharedURL {
	return s.SharedURL
}

// GetText returns the value of Text.
func (s *CreatePostReq) GetText() OptNilString {
	return s.Text
}

// GetVideoFileName returns the value of VideoFileName.
func (s *CreatePostReq) GetVideoFileName() OptNilString {
	return s.VideoFileName
}

// SetAttachment2Filename sets the value of Attachment2Filename.
func (s *CreatePostReq) SetAttachment2Filename(val OptNilString) {
	s.Attachment2Filename = val
}

// SetAttachment3Filename sets the value of Attachment3Filename.
func (s *CreatePostReq) SetAttachment3Filename(val OptNilString) {
	s.Attachment3Filename = val
}

// SetAttachment4Filename sets the value of Attachment4Filename.
func (s *CreatePostReq) SetAttachment4Filename(val OptNilString) {
	s.Attachment4Filename = val
}

// SetAttachment5Filename sets the value of Attachment5Filename.
func (s *CreatePostReq) SetAttachment5Filename(val OptNilString) {
	s.Attachment5Filename = val
}

// SetAttachment6Filename sets the value of Attachment6Filename.
func (s *CreatePostReq) SetAttachment6Filename(val OptNilString) {
	s.Attachment6Filename = val
}

// SetAttachment7Filename sets the value of Attachment7Filename.
func (s *CreatePostReq) SetAttachment7Filename(val OptNilString) {
	s.Attachment7Filename = val
}

// SetAttachment8Filename sets the value of Attachment8Filename.
func (s *CreatePostReq) SetAttachment8Filename(val OptNilString) {
	s.Attachment8Filename = val
}

// SetAttachment9Filename sets the value of Attachment9Filename.
func (s *CreatePostReq) SetAttachment9Filename(val OptNilString) {
	s.Attachment9Filename = val
}

// SetAttachmentFilename sets the value of AttachmentFilename.
func (s *CreatePostReq) SetAttachmentFilename(val OptNilString) {
	s.AttachmentFilename = val
}

// SetChoices sets the value of Choices.
func (s *CreatePostReq) SetChoices(val OptNilStringArray) {
	s.Choices = val
}

// SetColor sets the value of Color.
func (s *CreatePostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *CreatePostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGroupID sets the value of GroupID.
func (s *CreatePostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetInReplyTo sets the value of InReplyTo.
func (s *CreatePostReq) SetInReplyTo(val OptNilInt64) {
	s.InReplyTo = val
}

// SetLanguage sets the value of Language.
func (s *CreatePostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMentionIds sets the value of MentionIds.
func (s *CreatePostReq) SetMentionIds(val OptNilInt64Array) {
	s.MentionIds = val
}

// SetMessageTags sets the value of MessageTags.
func (s *CreatePostReq) SetMessageTags(val OptCreatePostReqMessageTags) {
	s.MessageTags = val
}

// SetPostType sets the value of PostType.
func (s *CreatePostReq) SetPostType(val OptPostType) {
	s.PostType = val
}

// SetSharedURL sets the value of SharedURL.
func (s *CreatePostReq) SetSharedURL(val OptCreatePostReqSharedURL) {
	s.SharedURL = val
}

// SetText sets the value of Text.
func (s *CreatePostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetVideoFileName sets the value of VideoFileName.
func (s *CreatePostReq) SetVideoFileName(val OptNilString) {
	s.VideoFileName = val
}

type CreatePostReqMessageTags struct{}

type CreatePostReqSharedURL struct{}

// Ref: #/components/schemas/CreatePostResponse
type CreatePostResponse struct {
	ConferenceCall OptRealmConferenceCall `json:"conference_call"`
	Post           *Post                  `json:"post"`
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *CreatePostResponse) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetPost returns the value of Post.
func (s *CreatePostResponse) GetPost() *Post {
	return s.Post
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *CreatePostResponse) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetPost sets the value of Post.
func (s *CreatePostResponse) SetPost(val *Post) {
	s.Post = val
}

// Ref: #/components/schemas/CreateQuotaResponse
type CreateQuotaResponse struct {
	Create OptCreateGroupQuota `json:"create"`
}

// GetCreate returns the value of Create.
func (s *CreateQuotaResponse) GetCreate() OptCreateGroupQuota {
	return s.Create
}

// SetCreate sets the value of Create.
func (s *CreateQuotaResponse) SetCreate(val OptCreateGroupQuota) {
	s.Create = val
}

// CreateReviewNoContent is response for CreateReview operation.
type CreateReviewNoContent struct{}

type CreateReviewReq struct {
	APIKey     OptString `json:"api_key"`
	Comment    OptString `json:"comment"`
	SignedInfo OptString `json:"signed_info"`
	Timestamp  OptInt64  `json:"timestamp"`
	UserIds    []int64   `json:"user_ids[]"`
	UUID       OptString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *CreateReviewReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetComment returns the value of Comment.
func (s *CreateReviewReq) GetComment() OptString {
	return s.Comment
}

// GetSignedInfo returns the value of SignedInfo.
func (s *CreateReviewReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *CreateReviewReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUserIds returns the value of UserIds.
func (s *CreateReviewReq) GetUserIds() []int64 {
	return s.UserIds
}

// GetUUID returns the value of UUID.
func (s *CreateReviewReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *CreateReviewReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetComment sets the value of Comment.
func (s *CreateReviewReq) SetComment(val OptString) {
	s.Comment = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *CreateReviewReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *CreateReviewReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUserIds sets the value of UserIds.
func (s *CreateReviewReq) SetUserIds(val []int64) {
	s.UserIds = val
}

// SetUUID sets the value of UUID.
func (s *CreateReviewReq) SetUUID(val OptString) {
	s.UUID = val
}

type CreateSharePostReq struct {
	APIKey        OptString                        `json:"api_key"`
	Color         OptNilInt32                      `json:"color"`
	FontSize      OptNilInt32                      `json:"font_size"`
	GroupID       OptNilInt64                      `json:"group_id"`
	Language      OptNilString                     `json:"language"`
	MessageTags   OptCreateSharePostReqMessageTags `json:"message_tags"`
	ShareableID   OptInt64                         `json:"shareable_id"`
	ShareableType OptString                        `json:"shareable_type"`
	SignedInfo    OptString                        `json:"signed_info"`
	Text          OptNilString                     `json:"text"`
	Timestamp     OptInt64                         `json:"timestamp"`
	UUID          OptString                        `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *CreateSharePostReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetColor returns the value of Color.
func (s *CreateSharePostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *CreateSharePostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGroupID returns the value of GroupID.
func (s *CreateSharePostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetLanguage returns the value of Language.
func (s *CreateSharePostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMessageTags returns the value of MessageTags.
func (s *CreateSharePostReq) GetMessageTags() OptCreateSharePostReqMessageTags {
	return s.MessageTags
}

// GetShareableID returns the value of ShareableID.
func (s *CreateSharePostReq) GetShareableID() OptInt64 {
	return s.ShareableID
}

// GetShareableType returns the value of ShareableType.
func (s *CreateSharePostReq) GetShareableType() OptString {
	return s.ShareableType
}

// GetSignedInfo returns the value of SignedInfo.
func (s *CreateSharePostReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetText returns the value of Text.
func (s *CreateSharePostReq) GetText() OptNilString {
	return s.Text
}

// GetTimestamp returns the value of Timestamp.
func (s *CreateSharePostReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *CreateSharePostReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *CreateSharePostReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetColor sets the value of Color.
func (s *CreateSharePostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *CreateSharePostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGroupID sets the value of GroupID.
func (s *CreateSharePostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetLanguage sets the value of Language.
func (s *CreateSharePostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMessageTags sets the value of MessageTags.
func (s *CreateSharePostReq) SetMessageTags(val OptCreateSharePostReqMessageTags) {
	s.MessageTags = val
}

// SetShareableID sets the value of ShareableID.
func (s *CreateSharePostReq) SetShareableID(val OptInt64) {
	s.ShareableID = val
}

// SetShareableType sets the value of ShareableType.
func (s *CreateSharePostReq) SetShareableType(val OptString) {
	s.ShareableType = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *CreateSharePostReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetText sets the value of Text.
func (s *CreateSharePostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetTimestamp sets the value of Timestamp.
func (s *CreateSharePostReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *CreateSharePostReq) SetUUID(val OptString) {
	s.UUID = val
}

type CreateSharePostReqMessageTags struct{}

type CreateThreadPostReq struct {
	Attachment2Filename OptNilString                      `json:"attachment_2_filename"`
	Attachment3Filename OptNilString                      `json:"attachment_3_filename"`
	Attachment4Filename OptNilString                      `json:"attachment_4_filename"`
	Attachment5Filename OptNilString                      `json:"attachment_5_filename"`
	Attachment6Filename OptNilString                      `json:"attachment_6_filename"`
	Attachment7Filename OptNilString                      `json:"attachment_7_filename"`
	Attachment8Filename OptNilString                      `json:"attachment_8_filename"`
	Attachment9Filename OptNilString                      `json:"attachment_9_filename"`
	AttachmentFilename  OptNilString                      `json:"attachment_filename"`
	Choices             OptNilStringArray                 `json:"choices[]"`
	Color               OptNilInt32                       `json:"color"`
	FontSize            OptNilInt32                       `json:"font_size"`
	GroupID             OptNilInt64                       `json:"group_id"`
	InReplyTo           OptNilInt64                       `json:"in_reply_to"`
	Language            OptNilString                      `json:"language"`
	MentionIds          OptNilInt64Array                  `json:"mention_ids[]"`
	MessageTags         OptCreateThreadPostReqMessageTags `json:"message_tags"`
	PostType            OptPostType                       `json:"post_type"`
	SharedURL           OptCreateThreadPostReqSharedURL   `json:"shared_url"`
	Text                OptNilString                      `json:"text"`
	VideoFileName       OptNilString                      `json:"video_file_name"`
}

// GetAttachment2Filename returns the value of Attachment2Filename.
func (s *CreateThreadPostReq) GetAttachment2Filename() OptNilString {
	return s.Attachment2Filename
}

// GetAttachment3Filename returns the value of Attachment3Filename.
func (s *CreateThreadPostReq) GetAttachment3Filename() OptNilString {
	return s.Attachment3Filename
}

// GetAttachment4Filename returns the value of Attachment4Filename.
func (s *CreateThreadPostReq) GetAttachment4Filename() OptNilString {
	return s.Attachment4Filename
}

// GetAttachment5Filename returns the value of Attachment5Filename.
func (s *CreateThreadPostReq) GetAttachment5Filename() OptNilString {
	return s.Attachment5Filename
}

// GetAttachment6Filename returns the value of Attachment6Filename.
func (s *CreateThreadPostReq) GetAttachment6Filename() OptNilString {
	return s.Attachment6Filename
}

// GetAttachment7Filename returns the value of Attachment7Filename.
func (s *CreateThreadPostReq) GetAttachment7Filename() OptNilString {
	return s.Attachment7Filename
}

// GetAttachment8Filename returns the value of Attachment8Filename.
func (s *CreateThreadPostReq) GetAttachment8Filename() OptNilString {
	return s.Attachment8Filename
}

// GetAttachment9Filename returns the value of Attachment9Filename.
func (s *CreateThreadPostReq) GetAttachment9Filename() OptNilString {
	return s.Attachment9Filename
}

// GetAttachmentFilename returns the value of AttachmentFilename.
func (s *CreateThreadPostReq) GetAttachmentFilename() OptNilString {
	return s.AttachmentFilename
}

// GetChoices returns the value of Choices.
func (s *CreateThreadPostReq) GetChoices() OptNilStringArray {
	return s.Choices
}

// GetColor returns the value of Color.
func (s *CreateThreadPostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *CreateThreadPostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGroupID returns the value of GroupID.
func (s *CreateThreadPostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetInReplyTo returns the value of InReplyTo.
func (s *CreateThreadPostReq) GetInReplyTo() OptNilInt64 {
	return s.InReplyTo
}

// GetLanguage returns the value of Language.
func (s *CreateThreadPostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMentionIds returns the value of MentionIds.
func (s *CreateThreadPostReq) GetMentionIds() OptNilInt64Array {
	return s.MentionIds
}

// GetMessageTags returns the value of MessageTags.
func (s *CreateThreadPostReq) GetMessageTags() OptCreateThreadPostReqMessageTags {
	return s.MessageTags
}

// GetPostType returns the value of PostType.
func (s *CreateThreadPostReq) GetPostType() OptPostType {
	return s.PostType
}

// GetSharedURL returns the value of SharedURL.
func (s *CreateThreadPostReq) GetSharedURL() OptCreateThreadPostReqSharedURL {
	return s.SharedURL
}

// GetText returns the value of Text.
func (s *CreateThreadPostReq) GetText() OptNilString {
	return s.Text
}

// GetVideoFileName returns the value of VideoFileName.
func (s *CreateThreadPostReq) GetVideoFileName() OptNilString {
	return s.VideoFileName
}

// SetAttachment2Filename sets the value of Attachment2Filename.
func (s *CreateThreadPostReq) SetAttachment2Filename(val OptNilString) {
	s.Attachment2Filename = val
}

// SetAttachment3Filename sets the value of Attachment3Filename.
func (s *CreateThreadPostReq) SetAttachment3Filename(val OptNilString) {
	s.Attachment3Filename = val
}

// SetAttachment4Filename sets the value of Attachment4Filename.
func (s *CreateThreadPostReq) SetAttachment4Filename(val OptNilString) {
	s.Attachment4Filename = val
}

// SetAttachment5Filename sets the value of Attachment5Filename.
func (s *CreateThreadPostReq) SetAttachment5Filename(val OptNilString) {
	s.Attachment5Filename = val
}

// SetAttachment6Filename sets the value of Attachment6Filename.
func (s *CreateThreadPostReq) SetAttachment6Filename(val OptNilString) {
	s.Attachment6Filename = val
}

// SetAttachment7Filename sets the value of Attachment7Filename.
func (s *CreateThreadPostReq) SetAttachment7Filename(val OptNilString) {
	s.Attachment7Filename = val
}

// SetAttachment8Filename sets the value of Attachment8Filename.
func (s *CreateThreadPostReq) SetAttachment8Filename(val OptNilString) {
	s.Attachment8Filename = val
}

// SetAttachment9Filename sets the value of Attachment9Filename.
func (s *CreateThreadPostReq) SetAttachment9Filename(val OptNilString) {
	s.Attachment9Filename = val
}

// SetAttachmentFilename sets the value of AttachmentFilename.
func (s *CreateThreadPostReq) SetAttachmentFilename(val OptNilString) {
	s.AttachmentFilename = val
}

// SetChoices sets the value of Choices.
func (s *CreateThreadPostReq) SetChoices(val OptNilStringArray) {
	s.Choices = val
}

// SetColor sets the value of Color.
func (s *CreateThreadPostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *CreateThreadPostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGroupID sets the value of GroupID.
func (s *CreateThreadPostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetInReplyTo sets the value of InReplyTo.
func (s *CreateThreadPostReq) SetInReplyTo(val OptNilInt64) {
	s.InReplyTo = val
}

// SetLanguage sets the value of Language.
func (s *CreateThreadPostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMentionIds sets the value of MentionIds.
func (s *CreateThreadPostReq) SetMentionIds(val OptNilInt64Array) {
	s.MentionIds = val
}

// SetMessageTags sets the value of MessageTags.
func (s *CreateThreadPostReq) SetMessageTags(val OptCreateThreadPostReqMessageTags) {
	s.MessageTags = val
}

// SetPostType sets the value of PostType.
func (s *CreateThreadPostReq) SetPostType(val OptPostType) {
	s.PostType = val
}

// SetSharedURL sets the value of SharedURL.
func (s *CreateThreadPostReq) SetSharedURL(val OptCreateThreadPostReqSharedURL) {
	s.SharedURL = val
}

// SetText sets the value of Text.
func (s *CreateThreadPostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetVideoFileName sets the value of VideoFileName.
func (s *CreateThreadPostReq) SetVideoFileName(val OptNilString) {
	s.VideoFileName = val
}

type CreateThreadPostReqMessageTags struct{}

type CreateThreadPostReqSharedURL struct{}

// DeclineGroupJoinRequestNoContent is response for DeclineGroupJoinRequest operation.
type DeclineGroupJoinRequestNoContent struct{}

// Ref: #/components/schemas/DefaultSettingsResponse
type DefaultSettingsResponse struct {
	TimelineSettings OptTimelineSettings `json:"timeline_settings"`
}

// GetTimelineSettings returns the value of TimelineSettings.
func (s *DefaultSettingsResponse) GetTimelineSettings() OptTimelineSettings {
	return s.TimelineSettings
}

// SetTimelineSettings sets the value of TimelineSettings.
func (s *DefaultSettingsResponse) SetTimelineSettings(val OptTimelineSettings) {
	s.TimelineSettings = val
}

// DeleteAllPostsNoContent is response for DeleteAllPosts operation.
type DeleteAllPostsNoContent struct{}

// DeleteBookmarkNoContent is response for DeleteBookmark operation.
type DeleteBookmarkNoContent struct{}

// DeleteChatMessageNoContent is response for DeleteChatMessage operation.
type DeleteChatMessageNoContent struct{}

// DeleteChatRoomsNoContent is response for DeleteChatRooms operation.
type DeleteChatRoomsNoContent struct{}

type DeleteChatRoomsReq struct {
	ChatRoomIds []int64 `json:"chat_room_ids[]"`
}

// GetChatRoomIds returns the value of ChatRoomIds.
func (s *DeleteChatRoomsReq) GetChatRoomIds() []int64 {
	return s.ChatRoomIds
}

// SetChatRoomIds sets the value of ChatRoomIds.
func (s *DeleteChatRoomsReq) SetChatRoomIds(val []int64) {
	s.ChatRoomIds = val
}

// DeleteFootprintNoContent is response for DeleteFootprint operation.
type DeleteFootprintNoContent struct{}

// DeleteMuteKeywordNoContent is response for DeleteMuteKeyword operation.
type DeleteMuteKeywordNoContent struct{}

// DeleteMyReviewsNoContent is response for DeleteMyReviews operation.
type DeleteMyReviewsNoContent struct{}

// DeletePostsNoContent is response for DeletePosts operation.
type DeletePostsNoContent struct{}

type DeletePostsReq struct {
	PostsIds []int64 `json:"posts_ids[]"`
}

// GetPostsIds returns the value of PostsIds.
func (s *DeletePostsReq) GetPostsIds() []int64 {
	return s.PostsIds
}

// SetPostsIds sets the value of PostsIds.
func (s *DeletePostsReq) SetPostsIds(val []int64) {
	s.PostsIds = val
}

// DeleteThreadNoContent is response for DeleteThread operation.
type DeleteThreadNoContent struct{}

// DeputizeGroupUsersMassNoContent is response for DeputizeGroupUsersMass operation.
type DeputizeGroupUsersMassNoContent struct{}

type DeputizeGroupUsersMassReq struct {
	APIKey     OptString `json:"api_key"`
	SignedInfo OptString `json:"signed_info"`
	Timestamp  OptInt64  `json:"timestamp"`
	UserIds    []int64   `json:"user_ids[]"`
	UUID       OptString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *DeputizeGroupUsersMassReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetSignedInfo returns the value of SignedInfo.
func (s *DeputizeGroupUsersMassReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *DeputizeGroupUsersMassReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUserIds returns the value of UserIds.
func (s *DeputizeGroupUsersMassReq) GetUserIds() []int64 {
	return s.UserIds
}

// GetUUID returns the value of UUID.
func (s *DeputizeGroupUsersMassReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *DeputizeGroupUsersMassReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *DeputizeGroupUsersMassReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *DeputizeGroupUsersMassReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUserIds sets the value of UserIds.
func (s *DeputizeGroupUsersMassReq) SetUserIds(val []int64) {
	s.UserIds = val
}

// SetUUID sets the value of UUID.
func (s *DeputizeGroupUsersMassReq) SetUUID(val OptString) {
	s.UUID = val
}

// DeputizeGroupUsersNoContent is response for DeputizeGroupUsers operation.
type DeputizeGroupUsersNoContent struct{}

// DisableTwoFactorAuthNoContent is response for DisableTwoFactorAuth operation.
type DisableTwoFactorAuthNoContent struct{}

type DisableTwoFactorAuthReq struct {
	Code OptString `json:"code"`
}

// GetCode returns the value of Code.
func (s *DisableTwoFactorAuthReq) GetCode() OptString {
	return s.Code
}

// SetCode sets the value of Code.
func (s *DisableTwoFactorAuthReq) SetCode(val OptString) {
	s.Code = val
}

// EditUserNoContent is response for EditUser operation.
type EditUserNoContent struct{}

type EditUserReq struct {
	APIKey              OptString    `json:"api_key"`
	Biography           OptNilString `json:"biography"`
	CountryCode         OptNilString `json:"country_code"`
	CoverImageFilename  OptNilString `json:"cover_image_filename"`
	Gender              OptNilInt32  `json:"gender"`
	Nickname            OptString    `json:"nickname"`
	Prefecture          OptNilString `json:"prefecture"`
	ProfileIconFilename OptNilString `json:"profile_icon_filename"`
	SignedInfo          OptString    `json:"signed_info"`
	Timestamp           OptInt64     `json:"timestamp"`
	Username            OptNilString `json:"username"`
	UUID                OptString    `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *EditUserReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetBiography returns the value of Biography.
func (s *EditUserReq) GetBiography() OptNilString {
	return s.Biography
}

// GetCountryCode returns the value of CountryCode.
func (s *EditUserReq) GetCountryCode() OptNilString {
	return s.CountryCode
}

// GetCoverImageFilename returns the value of CoverImageFilename.
func (s *EditUserReq) GetCoverImageFilename() OptNilString {
	return s.CoverImageFilename
}

// GetGender returns the value of Gender.
func (s *EditUserReq) GetGender() OptNilInt32 {
	return s.Gender
}

// GetNickname returns the value of Nickname.
func (s *EditUserReq) GetNickname() OptString {
	return s.Nickname
}

// GetPrefecture returns the value of Prefecture.
func (s *EditUserReq) GetPrefecture() OptNilString {
	return s.Prefecture
}

// GetProfileIconFilename returns the value of ProfileIconFilename.
func (s *EditUserReq) GetProfileIconFilename() OptNilString {
	return s.ProfileIconFilename
}

// GetSignedInfo returns the value of SignedInfo.
func (s *EditUserReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *EditUserReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUsername returns the value of Username.
func (s *EditUserReq) GetUsername() OptNilString {
	return s.Username
}

// GetUUID returns the value of UUID.
func (s *EditUserReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *EditUserReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetBiography sets the value of Biography.
func (s *EditUserReq) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetCountryCode sets the value of CountryCode.
func (s *EditUserReq) SetCountryCode(val OptNilString) {
	s.CountryCode = val
}

// SetCoverImageFilename sets the value of CoverImageFilename.
func (s *EditUserReq) SetCoverImageFilename(val OptNilString) {
	s.CoverImageFilename = val
}

// SetGender sets the value of Gender.
func (s *EditUserReq) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetNickname sets the value of Nickname.
func (s *EditUserReq) SetNickname(val OptString) {
	s.Nickname = val
}

// SetPrefecture sets the value of Prefecture.
func (s *EditUserReq) SetPrefecture(val OptNilString) {
	s.Prefecture = val
}

// SetProfileIconFilename sets the value of ProfileIconFilename.
func (s *EditUserReq) SetProfileIconFilename(val OptNilString) {
	s.ProfileIconFilename = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *EditUserReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *EditUserReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUsername sets the value of Username.
func (s *EditUserReq) SetUsername(val OptNilString) {
	s.Username = val
}

// SetUUID sets the value of UUID.
func (s *EditUserReq) SetUUID(val OptString) {
	s.UUID = val
}

// EditUserV2NoContent is response for EditUserV2 operation.
type EditUserV2NoContent struct{}

type EditUserV2Req struct {
	APIKey     OptString `json:"api_key"`
	IsPrivate  OptBool   `json:"is_private"`
	Nickname   OptString `json:"nickname"`
	SignedInfo OptString `json:"signed_info"`
	Timestamp  OptInt64  `json:"timestamp"`
	UUID       OptString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *EditUserV2Req) GetAPIKey() OptString {
	return s.APIKey
}

// GetIsPrivate returns the value of IsPrivate.
func (s *EditUserV2Req) GetIsPrivate() OptBool {
	return s.IsPrivate
}

// GetNickname returns the value of Nickname.
func (s *EditUserV2Req) GetNickname() OptString {
	return s.Nickname
}

// GetSignedInfo returns the value of SignedInfo.
func (s *EditUserV2Req) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *EditUserV2Req) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *EditUserV2Req) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *EditUserV2Req) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *EditUserV2Req) SetIsPrivate(val OptBool) {
	s.IsPrivate = val
}

// SetNickname sets the value of Nickname.
func (s *EditUserV2Req) SetNickname(val OptString) {
	s.Nickname = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *EditUserV2Req) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *EditUserV2Req) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *EditUserV2Req) SetUUID(val OptString) {
	s.UUID = val
}

type EnableTwoFactorAuthReq struct {
	Code OptString `json:"code"`
	Type OptString `json:"type"`
}

// GetCode returns the value of Code.
func (s *EnableTwoFactorAuthReq) GetCode() OptString {
	return s.Code
}

// GetType returns the value of Type.
func (s *EnableTwoFactorAuthReq) GetType() OptString {
	return s.Type
}

// SetCode sets the value of Code.
func (s *EnableTwoFactorAuthReq) SetCode(val OptString) {
	s.Code = val
}

// SetType sets the value of Type.
func (s *EnableTwoFactorAuthReq) SetType(val OptString) {
	s.Type = val
}

// FireGroupUserNoContent is response for FireGroupUser operation.
type FireGroupUserNoContent struct{}

// Ref: #/components/schemas/FollowRequestCountResponse
type FollowRequestCountResponse struct {
	UsersCount OptNilInt32 `json:"users_count"`
}

// GetUsersCount returns the value of UsersCount.
func (s *FollowRequestCountResponse) GetUsersCount() OptNilInt32 {
	return s.UsersCount
}

// SetUsersCount sets the value of UsersCount.
func (s *FollowRequestCountResponse) SetUsersCount(val OptNilInt32) {
	s.UsersCount = val
}

// FollowUserNoContent is response for FollowUser operation.
type FollowUserNoContent struct{}

// FollowUsersNoContent is response for FollowUsers operation.
type FollowUsersNoContent struct{}

// Ref: #/components/schemas/FollowUsersResponse
type FollowUsersResponse struct {
	NextPageValue OptNilString         `json:"next_page_value"`
	Users         OptNilRealmUserArray `json:"users"`
}

// GetNextPageValue returns the value of NextPageValue.
func (s *FollowUsersResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetUsers returns the value of Users.
func (s *FollowUsersResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetNextPageValue sets the value of NextPageValue.
func (s *FollowUsersResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetUsers sets the value of Users.
func (s *FollowUsersResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// Ref: #/components/schemas/FootprintDTO
type FootprintDTO struct {
	ID        OptNilInt64    `json:"id"`
	User      OptUserUserDTO `json:"user"`
	VisitedAt OptNilInt64    `json:"visited_at"`
}

// GetID returns the value of ID.
func (s *FootprintDTO) GetID() OptNilInt64 {
	return s.ID
}

// GetUser returns the value of User.
func (s *FootprintDTO) GetUser() OptUserUserDTO {
	return s.User
}

// GetVisitedAt returns the value of VisitedAt.
func (s *FootprintDTO) GetVisitedAt() OptNilInt64 {
	return s.VisitedAt
}

// SetID sets the value of ID.
func (s *FootprintDTO) SetID(val OptNilInt64) {
	s.ID = val
}

// SetUser sets the value of User.
func (s *FootprintDTO) SetUser(val OptUserUserDTO) {
	s.User = val
}

// SetVisitedAt sets the value of VisitedAt.
func (s *FootprintDTO) SetVisitedAt(val OptNilInt64) {
	s.VisitedAt = val
}

// Ref: #/components/schemas/FootprintsResponse
type FootprintsResponse struct {
	Footprints    OptNilFootprintDTOArray `json:"footprints"`
	NextPageValue OptNilString            `json:"next_page_value"`
}

// GetFootprints returns the value of Footprints.
func (s *FootprintsResponse) GetFootprints() OptNilFootprintDTOArray {
	return s.Footprints
}

// GetNextPageValue returns the value of NextPageValue.
func (s *FootprintsResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// SetFootprints sets the value of Footprints.
func (s *FootprintsResponse) SetFootprints(val OptNilFootprintDTOArray) {
	s.Footprints = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *FootprintsResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// Ref: #/components/schemas/GamesResponse
type GamesResponse struct {
	FromID OptNilInt64          `json:"from_id"`
	Games  OptNilRealmGameArray `json:"games"`
}

// GetFromID returns the value of FromID.
func (s *GamesResponse) GetFromID() OptNilInt64 {
	return s.FromID
}

// GetGames returns the value of Games.
func (s *GamesResponse) GetGames() OptNilRealmGameArray {
	return s.Games
}

// SetFromID sets the value of FromID.
func (s *GamesResponse) SetFromID(val OptNilInt64) {
	s.FromID = val
}

// SetGames sets the value of Games.
func (s *GamesResponse) SetGames(val OptNilRealmGameArray) {
	s.Games = val
}

// Ref: #/components/schemas/Gender
type Gender struct {
	APIIntValue    OptNilInt32  `json:"api_int_value"`
	APIStringValue OptNilString `json:"api_string_value"`
}

// GetAPIIntValue returns the value of APIIntValue.
func (s *Gender) GetAPIIntValue() OptNilInt32 {
	return s.APIIntValue
}

// GetAPIStringValue returns the value of APIStringValue.
func (s *Gender) GetAPIStringValue() OptNilString {
	return s.APIStringValue
}

// SetAPIIntValue sets the value of APIIntValue.
func (s *Gender) SetAPIIntValue(val OptNilInt32) {
	s.APIIntValue = val
}

// SetAPIStringValue sets the value of APIStringValue.
func (s *Gender) SetAPIStringValue(val OptNilString) {
	s.APIStringValue = val
}

type GenerateCallActionSignatureReq struct {
	Action         OptString `json:"action"`
	ConferenceID   OptInt64  `json:"conference_id"`
	TargetUserUUID OptString `json:"target_user_uuid"`
}

// GetAction returns the value of Action.
func (s *GenerateCallActionSignatureReq) GetAction() OptString {
	return s.Action
}

// GetConferenceID returns the value of ConferenceID.
func (s *GenerateCallActionSignatureReq) GetConferenceID() OptInt64 {
	return s.ConferenceID
}

// GetTargetUserUUID returns the value of TargetUserUUID.
func (s *GenerateCallActionSignatureReq) GetTargetUserUUID() OptString {
	return s.TargetUserUUID
}

// SetAction sets the value of Action.
func (s *GenerateCallActionSignatureReq) SetAction(val OptString) {
	s.Action = val
}

// SetConferenceID sets the value of ConferenceID.
func (s *GenerateCallActionSignatureReq) SetConferenceID(val OptInt64) {
	s.ConferenceID = val
}

// SetTargetUserUUID sets the value of TargetUserUUID.
func (s *GenerateCallActionSignatureReq) SetTargetUserUUID(val OptString) {
	s.TargetUserUUID = val
}

// Ref: #/components/schemas/GenresResponse
type GenresResponse struct {
	Genres        OptNilRealmGenreArray `json:"genres"`
	NextPageValue OptNilString          `json:"next_page_value"`
}

// GetGenres returns the value of Genres.
func (s *GenresResponse) GetGenres() OptNilRealmGenreArray {
	return s.Genres
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GenresResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// SetGenres sets the value of Genres.
func (s *GenresResponse) SetGenres(val OptNilRealmGenreArray) {
	s.Genres = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GenresResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

type GetJoinedGroupStatusesOK map[string]string

func (s *GetJoinedGroupStatusesOK) init() GetJoinedGroupStatusesOK {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

type GetJoinedThreadStatusesOK map[string]string

func (s *GetJoinedThreadStatusesOK) init() GetJoinedThreadStatusesOK {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

type GetRecommendedPostTagsReq struct {
	SaveRecentSearch OptBool   `json:"save_recent_search"`
	Tag              OptString `json:"tag"`
}

// GetSaveRecentSearch returns the value of SaveRecentSearch.
func (s *GetRecommendedPostTagsReq) GetSaveRecentSearch() OptBool {
	return s.SaveRecentSearch
}

// GetTag returns the value of Tag.
func (s *GetRecommendedPostTagsReq) GetTag() OptString {
	return s.Tag
}

// SetSaveRecentSearch sets the value of SaveRecentSearch.
func (s *GetRecommendedPostTagsReq) SetSaveRecentSearch(val OptBool) {
	s.SaveRecentSearch = val
}

// SetTag sets the value of Tag.
func (s *GetRecommendedPostTagsReq) SetTag(val OptString) {
	s.Tag = val
}

// Ref: #/components/schemas/GifImage
type GifImage struct {
	Height OptNilInt32  `json:"height"`
	ID     OptNilInt64  `json:"id"`
	URL    OptNilString `json:"url"`
	Width  OptNilInt32  `json:"width"`
}

// GetHeight returns the value of Height.
func (s *GifImage) GetHeight() OptNilInt32 {
	return s.Height
}

// GetID returns the value of ID.
func (s *GifImage) GetID() OptNilInt64 {
	return s.ID
}

// GetURL returns the value of URL.
func (s *GifImage) GetURL() OptNilString {
	return s.URL
}

// GetWidth returns the value of Width.
func (s *GifImage) GetWidth() OptNilInt32 {
	return s.Width
}

// SetHeight sets the value of Height.
func (s *GifImage) SetHeight(val OptNilInt32) {
	s.Height = val
}

// SetID sets the value of ID.
func (s *GifImage) SetID(val OptNilInt64) {
	s.ID = val
}

// SetURL sets the value of URL.
func (s *GifImage) SetURL(val OptNilString) {
	s.URL = val
}

// SetWidth sets the value of Width.
func (s *GifImage) SetWidth(val OptNilInt32) {
	s.Width = val
}

// Ref: #/components/schemas/Gift
type Gift struct {
	Currency      OptNilString  `json:"currency"`
	Icon          OptNilString  `json:"icon"`
	IconThumbnail OptNilString  `json:"icon_thumbnail"`
	ID            OptNilInt64   `json:"id"`
	Price         OptNilFloat64 `json:"price"`
	Title         OptNilString  `json:"title"`
}

// GetCurrency returns the value of Currency.
func (s *Gift) GetCurrency() OptNilString {
	return s.Currency
}

// GetIcon returns the value of Icon.
func (s *Gift) GetIcon() OptNilString {
	return s.Icon
}

// GetIconThumbnail returns the value of IconThumbnail.
func (s *Gift) GetIconThumbnail() OptNilString {
	return s.IconThumbnail
}

// GetID returns the value of ID.
func (s *Gift) GetID() OptNilInt64 {
	return s.ID
}

// GetPrice returns the value of Price.
func (s *Gift) GetPrice() OptNilFloat64 {
	return s.Price
}

// GetTitle returns the value of Title.
func (s *Gift) GetTitle() OptNilString {
	return s.Title
}

// SetCurrency sets the value of Currency.
func (s *Gift) SetCurrency(val OptNilString) {
	s.Currency = val
}

// SetIcon sets the value of Icon.
func (s *Gift) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetIconThumbnail sets the value of IconThumbnail.
func (s *Gift) SetIconThumbnail(val OptNilString) {
	s.IconThumbnail = val
}

// SetID sets the value of ID.
func (s *Gift) SetID(val OptNilInt64) {
	s.ID = val
}

// SetPrice sets the value of Price.
func (s *Gift) SetPrice(val OptNilFloat64) {
	s.Price = val
}

// SetTitle sets the value of Title.
func (s *Gift) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/GiftCount
type GiftCount struct {
	ID       OptNilInt64 `json:"id"`
	Quantity OptNilInt32 `json:"quantity"`
}

// GetID returns the value of ID.
func (s *GiftCount) GetID() OptNilInt64 {
	return s.ID
}

// GetQuantity returns the value of Quantity.
func (s *GiftCount) GetQuantity() OptNilInt32 {
	return s.Quantity
}

// SetID sets the value of ID.
func (s *GiftCount) SetID(val OptNilInt64) {
	s.ID = val
}

// SetQuantity sets the value of Quantity.
func (s *GiftCount) SetQuantity(val OptNilInt32) {
	s.Quantity = val
}

// Ref: #/components/schemas/GiftHistory
type GiftHistory struct {
	Anonymous  OptNilBool           `json:"anonymous"`
	GiftItem   OptGiftSlugItem      `json:"gift_item"`
	GiftsCount OptNilGiftCountArray `json:"gifts_count"`
	Receiver   *RealmUser           `json:"receiver"`
	Sender     *RealmUser           `json:"sender"`
	SentAt     OptNilInt64          `json:"sent_at"`
}

// GetAnonymous returns the value of Anonymous.
func (s *GiftHistory) GetAnonymous() OptNilBool {
	return s.Anonymous
}

// GetGiftItem returns the value of GiftItem.
func (s *GiftHistory) GetGiftItem() OptGiftSlugItem {
	return s.GiftItem
}

// GetGiftsCount returns the value of GiftsCount.
func (s *GiftHistory) GetGiftsCount() OptNilGiftCountArray {
	return s.GiftsCount
}

// GetReceiver returns the value of Receiver.
func (s *GiftHistory) GetReceiver() *RealmUser {
	return s.Receiver
}

// GetSender returns the value of Sender.
func (s *GiftHistory) GetSender() *RealmUser {
	return s.Sender
}

// GetSentAt returns the value of SentAt.
func (s *GiftHistory) GetSentAt() OptNilInt64 {
	return s.SentAt
}

// SetAnonymous sets the value of Anonymous.
func (s *GiftHistory) SetAnonymous(val OptNilBool) {
	s.Anonymous = val
}

// SetGiftItem sets the value of GiftItem.
func (s *GiftHistory) SetGiftItem(val OptGiftSlugItem) {
	s.GiftItem = val
}

// SetGiftsCount sets the value of GiftsCount.
func (s *GiftHistory) SetGiftsCount(val OptNilGiftCountArray) {
	s.GiftsCount = val
}

// SetReceiver sets the value of Receiver.
func (s *GiftHistory) SetReceiver(val *RealmUser) {
	s.Receiver = val
}

// SetSender sets the value of Sender.
func (s *GiftHistory) SetSender(val *RealmUser) {
	s.Sender = val
}

// SetSentAt sets the value of SentAt.
func (s *GiftHistory) SetSentAt(val OptNilInt64) {
	s.SentAt = val
}

// Ref: #/components/schemas/GiftReceivedResponse
type GiftReceivedResponse struct {
	NextPageValue OptNilString            `json:"next_page_value"`
	ReceivedGifts OptNilReceivedGiftArray `json:"received_gifts"`
	TotalCount    OptNilInt32             `json:"total_count"`
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GiftReceivedResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetReceivedGifts returns the value of ReceivedGifts.
func (s *GiftReceivedResponse) GetReceivedGifts() OptNilReceivedGiftArray {
	return s.ReceivedGifts
}

// GetTotalCount returns the value of TotalCount.
func (s *GiftReceivedResponse) GetTotalCount() OptNilInt32 {
	return s.TotalCount
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GiftReceivedResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetReceivedGifts sets the value of ReceivedGifts.
func (s *GiftReceivedResponse) SetReceivedGifts(val OptNilReceivedGiftArray) {
	s.ReceivedGifts = val
}

// SetTotalCount sets the value of TotalCount.
func (s *GiftReceivedResponse) SetTotalCount(val OptNilInt32) {
	s.TotalCount = val
}

// Ref: #/components/schemas/GiftReceivedTransactionResponse
type GiftReceivedTransactionResponse struct {
	CreatedAt OptNilInt64                        `json:"created_at"`
	Gifts     OptNilTransactionGiftReceivedArray `json:"gifts"`
	Receiver  *RealmUser                         `json:"receiver"`
}

// GetCreatedAt returns the value of CreatedAt.
func (s *GiftReceivedTransactionResponse) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetGifts returns the value of Gifts.
func (s *GiftReceivedTransactionResponse) GetGifts() OptNilTransactionGiftReceivedArray {
	return s.Gifts
}

// GetReceiver returns the value of Receiver.
func (s *GiftReceivedTransactionResponse) GetReceiver() *RealmUser {
	return s.Receiver
}

// SetCreatedAt sets the value of CreatedAt.
func (s *GiftReceivedTransactionResponse) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetGifts sets the value of Gifts.
func (s *GiftReceivedTransactionResponse) SetGifts(val OptNilTransactionGiftReceivedArray) {
	s.Gifts = val
}

// SetReceiver sets the value of Receiver.
func (s *GiftReceivedTransactionResponse) SetReceiver(val *RealmUser) {
	s.Receiver = val
}

// Ref: #/components/schemas/GiftSendersResponse
type GiftSendersResponse struct {
	Senders           OptNilRealmUserArray `json:"senders"`
	TotalSendersCount OptNilInt32          `json:"total_senders_count"`
}

// GetSenders returns the value of Senders.
func (s *GiftSendersResponse) GetSenders() OptNilRealmUserArray {
	return s.Senders
}

// GetTotalSendersCount returns the value of TotalSendersCount.
func (s *GiftSendersResponse) GetTotalSendersCount() OptNilInt32 {
	return s.TotalSendersCount
}

// SetSenders sets the value of Senders.
func (s *GiftSendersResponse) SetSenders(val OptNilRealmUserArray) {
	s.Senders = val
}

// SetTotalSendersCount sets the value of TotalSendersCount.
func (s *GiftSendersResponse) SetTotalSendersCount(val OptNilInt32) {
	s.TotalSendersCount = val
}

// Ref: #/components/schemas/GiftSlugItem
type GiftSlugItem struct {
	Icon     OptNilString `json:"icon"`
	ID       OptNilInt64  `json:"id"`
	Quantity OptNilInt32  `json:"quantity"`
	Slug     OptNilString `json:"slug"`
}

// GetIcon returns the value of Icon.
func (s *GiftSlugItem) GetIcon() OptNilString {
	return s.Icon
}

// GetID returns the value of ID.
func (s *GiftSlugItem) GetID() OptNilInt64 {
	return s.ID
}

// GetQuantity returns the value of Quantity.
func (s *GiftSlugItem) GetQuantity() OptNilInt32 {
	return s.Quantity
}

// GetSlug returns the value of Slug.
func (s *GiftSlugItem) GetSlug() OptNilString {
	return s.Slug
}

// SetIcon sets the value of Icon.
func (s *GiftSlugItem) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetID sets the value of ID.
func (s *GiftSlugItem) SetID(val OptNilInt64) {
	s.ID = val
}

// SetQuantity sets the value of Quantity.
func (s *GiftSlugItem) SetQuantity(val OptNilInt32) {
	s.Quantity = val
}

// SetSlug sets the value of Slug.
func (s *GiftSlugItem) SetSlug(val OptNilString) {
	s.Slug = val
}

// Ref: #/components/schemas/GiftTransactionsResponse
type GiftTransactionsResponse struct {
	HideGiftsReceived OptNilBool             `json:"hide_gifts_received"`
	NextPageValue     OptNilString           `json:"next_page_value"`
	SentGifts         OptNilGiftHistoryArray `json:"sent_gifts"`
}

// GetHideGiftsReceived returns the value of HideGiftsReceived.
func (s *GiftTransactionsResponse) GetHideGiftsReceived() OptNilBool {
	return s.HideGiftsReceived
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GiftTransactionsResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetSentGifts returns the value of SentGifts.
func (s *GiftTransactionsResponse) GetSentGifts() OptNilGiftHistoryArray {
	return s.SentGifts
}

// SetHideGiftsReceived sets the value of HideGiftsReceived.
func (s *GiftTransactionsResponse) SetHideGiftsReceived(val OptNilBool) {
	s.HideGiftsReceived = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GiftTransactionsResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetSentGifts sets the value of SentGifts.
func (s *GiftTransactionsResponse) SetSentGifts(val OptNilGiftHistoryArray) {
	s.SentGifts = val
}

// Ref: #/components/schemas/GiftingAbility
type GiftingAbility struct {
	CanReceive OptNilBool  `json:"can_receive"`
	CanSend    OptNilBool  `json:"can_send"`
	Enabled    OptNilBool  `json:"enabled"`
	UserID     OptNilInt64 `json:"user_id"`
}

// GetCanReceive returns the value of CanReceive.
func (s *GiftingAbility) GetCanReceive() OptNilBool {
	return s.CanReceive
}

// GetCanSend returns the value of CanSend.
func (s *GiftingAbility) GetCanSend() OptNilBool {
	return s.CanSend
}

// GetEnabled returns the value of Enabled.
func (s *GiftingAbility) GetEnabled() OptNilBool {
	return s.Enabled
}

// GetUserID returns the value of UserID.
func (s *GiftingAbility) GetUserID() OptNilInt64 {
	return s.UserID
}

// SetCanReceive sets the value of CanReceive.
func (s *GiftingAbility) SetCanReceive(val OptNilBool) {
	s.CanReceive = val
}

// SetCanSend sets the value of CanSend.
func (s *GiftingAbility) SetCanSend(val OptNilBool) {
	s.CanSend = val
}

// SetEnabled sets the value of Enabled.
func (s *GiftingAbility) SetEnabled(val OptNilBool) {
	s.Enabled = val
}

// SetUserID sets the value of UserID.
func (s *GiftingAbility) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/GiftsResponse
type GiftsResponse struct {
	Gifts         OptNilRealmGiftArray `json:"gifts"`
	NextPageValue OptNilString         `json:"next_page_value"`
	TotalCount    OptNilInt32          `json:"total_count"`
}

// GetGifts returns the value of Gifts.
func (s *GiftsResponse) GetGifts() OptNilRealmGiftArray {
	return s.Gifts
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GiftsResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetTotalCount returns the value of TotalCount.
func (s *GiftsResponse) GetTotalCount() OptNilInt32 {
	return s.TotalCount
}

// SetGifts sets the value of Gifts.
func (s *GiftsResponse) SetGifts(val OptNilRealmGiftArray) {
	s.Gifts = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GiftsResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetTotalCount sets the value of TotalCount.
func (s *GiftsResponse) SetTotalCount(val OptNilInt32) {
	s.TotalCount = val
}

// Ref: #/components/schemas/Group
type Group struct {
	AllowMembersToPostImageAndVideo OptNilBool       `json:"allow_members_to_post_image_and_video"`
	AllowMembersToPostURL           OptNilBool       `json:"allow_members_to_post_url"`
	AllowOwnershipTransfer          OptNilBool       `json:"allow_ownership_transfer"`
	AllowThreadCreationBy           OptNilString     `json:"allow_thread_creation_by"`
	CallTimelineDisplay             OptNilBool       `json:"call_timeline_display"`
	CoverImage                      OptNilString     `json:"cover_image"`
	CoverImageThumbnail             OptNilString     `json:"cover_image_thumbnail"`
	Description                     OptNilString     `json:"description"`
	Gender                          OptNilInt32      `json:"gender"`
	GenerationGroupsLimit           OptNilInt32      `json:"generation_groups_limit"`
	GroupCategoryID                 OptNilInt64      `json:"group_category_id"`
	GroupIcon                       OptNilString     `json:"group_icon"`
	GroupIconThumbnail              OptNilString     `json:"group_icon_thumbnail"`
	GroupsUsersCount                OptNilInt64      `json:"groups_users_count"`
	Guidelines                      OptNilString     `json:"guidelines"`
	HideConferenceCall              OptNilBool       `json:"hide_conference_call"`
	HideFromGameEight               OptNilBool       `json:"hide_from_game_eight"`
	HideReportedPosts               OptNilBool       `json:"hide_reported_posts"`
	HighlightedCount                OptNilInt64      `json:"highlighted_count"`
	ID                              OptNilInt64      `json:"id"`
	InvitedToJoin                   OptNilBool       `json:"invited_to_join"`
	IsJoined                        OptNilBool       `json:"is_joined"`
	IsPending                       OptNilBool       `json:"is_pending"`
	IsPrivate                       OptNilBool       `json:"is_private"`
	IsRelated                       OptNilBool       `json:"is_related"`
	JoinedCommunityCampaign         OptNilBool       `json:"joined_community_campaign"`
	ModeratorIds                    OptNilInt64Array `json:"moderator_ids"`
	OnlyMobileVerified              OptNilBool       `json:"only_mobile_verified"`
	OnlyVerifiedAge                 OptNilBool       `json:"only_verified_age"`
	Owner                           *RealmUser       `json:"owner"`
	PendingCount                    OptNilInt64      `json:"pending_count"`
	PendingDeputizeIds              OptNilInt64Array `json:"pending_deputize_ids"`
	PendingTransferID               OptNilInt64      `json:"pending_transfer_id"`
	PostsCount                      OptNilInt64      `json:"posts_count"`
	RelatedCount                    OptNilInt64      `json:"related_count"`
	SafeMode                        OptNilBool       `json:"safe_mode"`
	Secret                          OptNilBool       `json:"secret"`
	Seizable                        OptNilBool       `json:"seizable"`
	SeizableBefore                  OptNilInt64      `json:"seizable_before"`
	SubCategoryID                   OptNilInt64      `json:"sub_category_id"`
	ThreadsCount                    OptNilInt64      `json:"threads_count"`
	Title                           OptNilString     `json:"title"`
	Topic                           OptNilString     `json:"topic"`
	UnreadCounts                    OptNilInt64      `json:"unread_counts"`
	UnreadThreadsCount              OptNilInt64      `json:"unread_threads_count"`
	UpdatedAt                       OptNilInt64      `json:"updated_at"`
	UserID                          OptNilInt64      `json:"user_id"`
	ViewsCount                      OptNilInt64      `json:"views_count"`
	WalkthroughRequested            OptNilBool       `json:"walkthrough_requested"`
}

// GetAllowMembersToPostImageAndVideo returns the value of AllowMembersToPostImageAndVideo.
func (s *Group) GetAllowMembersToPostImageAndVideo() OptNilBool {
	return s.AllowMembersToPostImageAndVideo
}

// GetAllowMembersToPostURL returns the value of AllowMembersToPostURL.
func (s *Group) GetAllowMembersToPostURL() OptNilBool {
	return s.AllowMembersToPostURL
}

// GetAllowOwnershipTransfer returns the value of AllowOwnershipTransfer.
func (s *Group) GetAllowOwnershipTransfer() OptNilBool {
	return s.AllowOwnershipTransfer
}

// GetAllowThreadCreationBy returns the value of AllowThreadCreationBy.
func (s *Group) GetAllowThreadCreationBy() OptNilString {
	return s.AllowThreadCreationBy
}

// GetCallTimelineDisplay returns the value of CallTimelineDisplay.
func (s *Group) GetCallTimelineDisplay() OptNilBool {
	return s.CallTimelineDisplay
}

// GetCoverImage returns the value of CoverImage.
func (s *Group) GetCoverImage() OptNilString {
	return s.CoverImage
}

// GetCoverImageThumbnail returns the value of CoverImageThumbnail.
func (s *Group) GetCoverImageThumbnail() OptNilString {
	return s.CoverImageThumbnail
}

// GetDescription returns the value of Description.
func (s *Group) GetDescription() OptNilString {
	return s.Description
}

// GetGender returns the value of Gender.
func (s *Group) GetGender() OptNilInt32 {
	return s.Gender
}

// GetGenerationGroupsLimit returns the value of GenerationGroupsLimit.
func (s *Group) GetGenerationGroupsLimit() OptNilInt32 {
	return s.GenerationGroupsLimit
}

// GetGroupCategoryID returns the value of GroupCategoryID.
func (s *Group) GetGroupCategoryID() OptNilInt64 {
	return s.GroupCategoryID
}

// GetGroupIcon returns the value of GroupIcon.
func (s *Group) GetGroupIcon() OptNilString {
	return s.GroupIcon
}

// GetGroupIconThumbnail returns the value of GroupIconThumbnail.
func (s *Group) GetGroupIconThumbnail() OptNilString {
	return s.GroupIconThumbnail
}

// GetGroupsUsersCount returns the value of GroupsUsersCount.
func (s *Group) GetGroupsUsersCount() OptNilInt64 {
	return s.GroupsUsersCount
}

// GetGuidelines returns the value of Guidelines.
func (s *Group) GetGuidelines() OptNilString {
	return s.Guidelines
}

// GetHideConferenceCall returns the value of HideConferenceCall.
func (s *Group) GetHideConferenceCall() OptNilBool {
	return s.HideConferenceCall
}

// GetHideFromGameEight returns the value of HideFromGameEight.
func (s *Group) GetHideFromGameEight() OptNilBool {
	return s.HideFromGameEight
}

// GetHideReportedPosts returns the value of HideReportedPosts.
func (s *Group) GetHideReportedPosts() OptNilBool {
	return s.HideReportedPosts
}

// GetHighlightedCount returns the value of HighlightedCount.
func (s *Group) GetHighlightedCount() OptNilInt64 {
	return s.HighlightedCount
}

// GetID returns the value of ID.
func (s *Group) GetID() OptNilInt64 {
	return s.ID
}

// GetInvitedToJoin returns the value of InvitedToJoin.
func (s *Group) GetInvitedToJoin() OptNilBool {
	return s.InvitedToJoin
}

// GetIsJoined returns the value of IsJoined.
func (s *Group) GetIsJoined() OptNilBool {
	return s.IsJoined
}

// GetIsPending returns the value of IsPending.
func (s *Group) GetIsPending() OptNilBool {
	return s.IsPending
}

// GetIsPrivate returns the value of IsPrivate.
func (s *Group) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetIsRelated returns the value of IsRelated.
func (s *Group) GetIsRelated() OptNilBool {
	return s.IsRelated
}

// GetJoinedCommunityCampaign returns the value of JoinedCommunityCampaign.
func (s *Group) GetJoinedCommunityCampaign() OptNilBool {
	return s.JoinedCommunityCampaign
}

// GetModeratorIds returns the value of ModeratorIds.
func (s *Group) GetModeratorIds() OptNilInt64Array {
	return s.ModeratorIds
}

// GetOnlyMobileVerified returns the value of OnlyMobileVerified.
func (s *Group) GetOnlyMobileVerified() OptNilBool {
	return s.OnlyMobileVerified
}

// GetOnlyVerifiedAge returns the value of OnlyVerifiedAge.
func (s *Group) GetOnlyVerifiedAge() OptNilBool {
	return s.OnlyVerifiedAge
}

// GetOwner returns the value of Owner.
func (s *Group) GetOwner() *RealmUser {
	return s.Owner
}

// GetPendingCount returns the value of PendingCount.
func (s *Group) GetPendingCount() OptNilInt64 {
	return s.PendingCount
}

// GetPendingDeputizeIds returns the value of PendingDeputizeIds.
func (s *Group) GetPendingDeputizeIds() OptNilInt64Array {
	return s.PendingDeputizeIds
}

// GetPendingTransferID returns the value of PendingTransferID.
func (s *Group) GetPendingTransferID() OptNilInt64 {
	return s.PendingTransferID
}

// GetPostsCount returns the value of PostsCount.
func (s *Group) GetPostsCount() OptNilInt64 {
	return s.PostsCount
}

// GetRelatedCount returns the value of RelatedCount.
func (s *Group) GetRelatedCount() OptNilInt64 {
	return s.RelatedCount
}

// GetSafeMode returns the value of SafeMode.
func (s *Group) GetSafeMode() OptNilBool {
	return s.SafeMode
}

// GetSecret returns the value of Secret.
func (s *Group) GetSecret() OptNilBool {
	return s.Secret
}

// GetSeizable returns the value of Seizable.
func (s *Group) GetSeizable() OptNilBool {
	return s.Seizable
}

// GetSeizableBefore returns the value of SeizableBefore.
func (s *Group) GetSeizableBefore() OptNilInt64 {
	return s.SeizableBefore
}

// GetSubCategoryID returns the value of SubCategoryID.
func (s *Group) GetSubCategoryID() OptNilInt64 {
	return s.SubCategoryID
}

// GetThreadsCount returns the value of ThreadsCount.
func (s *Group) GetThreadsCount() OptNilInt64 {
	return s.ThreadsCount
}

// GetTitle returns the value of Title.
func (s *Group) GetTitle() OptNilString {
	return s.Title
}

// GetTopic returns the value of Topic.
func (s *Group) GetTopic() OptNilString {
	return s.Topic
}

// GetUnreadCounts returns the value of UnreadCounts.
func (s *Group) GetUnreadCounts() OptNilInt64 {
	return s.UnreadCounts
}

// GetUnreadThreadsCount returns the value of UnreadThreadsCount.
func (s *Group) GetUnreadThreadsCount() OptNilInt64 {
	return s.UnreadThreadsCount
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Group) GetUpdatedAt() OptNilInt64 {
	return s.UpdatedAt
}

// GetUserID returns the value of UserID.
func (s *Group) GetUserID() OptNilInt64 {
	return s.UserID
}

// GetViewsCount returns the value of ViewsCount.
func (s *Group) GetViewsCount() OptNilInt64 {
	return s.ViewsCount
}

// GetWalkthroughRequested returns the value of WalkthroughRequested.
func (s *Group) GetWalkthroughRequested() OptNilBool {
	return s.WalkthroughRequested
}

// SetAllowMembersToPostImageAndVideo sets the value of AllowMembersToPostImageAndVideo.
func (s *Group) SetAllowMembersToPostImageAndVideo(val OptNilBool) {
	s.AllowMembersToPostImageAndVideo = val
}

// SetAllowMembersToPostURL sets the value of AllowMembersToPostURL.
func (s *Group) SetAllowMembersToPostURL(val OptNilBool) {
	s.AllowMembersToPostURL = val
}

// SetAllowOwnershipTransfer sets the value of AllowOwnershipTransfer.
func (s *Group) SetAllowOwnershipTransfer(val OptNilBool) {
	s.AllowOwnershipTransfer = val
}

// SetAllowThreadCreationBy sets the value of AllowThreadCreationBy.
func (s *Group) SetAllowThreadCreationBy(val OptNilString) {
	s.AllowThreadCreationBy = val
}

// SetCallTimelineDisplay sets the value of CallTimelineDisplay.
func (s *Group) SetCallTimelineDisplay(val OptNilBool) {
	s.CallTimelineDisplay = val
}

// SetCoverImage sets the value of CoverImage.
func (s *Group) SetCoverImage(val OptNilString) {
	s.CoverImage = val
}

// SetCoverImageThumbnail sets the value of CoverImageThumbnail.
func (s *Group) SetCoverImageThumbnail(val OptNilString) {
	s.CoverImageThumbnail = val
}

// SetDescription sets the value of Description.
func (s *Group) SetDescription(val OptNilString) {
	s.Description = val
}

// SetGender sets the value of Gender.
func (s *Group) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetGenerationGroupsLimit sets the value of GenerationGroupsLimit.
func (s *Group) SetGenerationGroupsLimit(val OptNilInt32) {
	s.GenerationGroupsLimit = val
}

// SetGroupCategoryID sets the value of GroupCategoryID.
func (s *Group) SetGroupCategoryID(val OptNilInt64) {
	s.GroupCategoryID = val
}

// SetGroupIcon sets the value of GroupIcon.
func (s *Group) SetGroupIcon(val OptNilString) {
	s.GroupIcon = val
}

// SetGroupIconThumbnail sets the value of GroupIconThumbnail.
func (s *Group) SetGroupIconThumbnail(val OptNilString) {
	s.GroupIconThumbnail = val
}

// SetGroupsUsersCount sets the value of GroupsUsersCount.
func (s *Group) SetGroupsUsersCount(val OptNilInt64) {
	s.GroupsUsersCount = val
}

// SetGuidelines sets the value of Guidelines.
func (s *Group) SetGuidelines(val OptNilString) {
	s.Guidelines = val
}

// SetHideConferenceCall sets the value of HideConferenceCall.
func (s *Group) SetHideConferenceCall(val OptNilBool) {
	s.HideConferenceCall = val
}

// SetHideFromGameEight sets the value of HideFromGameEight.
func (s *Group) SetHideFromGameEight(val OptNilBool) {
	s.HideFromGameEight = val
}

// SetHideReportedPosts sets the value of HideReportedPosts.
func (s *Group) SetHideReportedPosts(val OptNilBool) {
	s.HideReportedPosts = val
}

// SetHighlightedCount sets the value of HighlightedCount.
func (s *Group) SetHighlightedCount(val OptNilInt64) {
	s.HighlightedCount = val
}

// SetID sets the value of ID.
func (s *Group) SetID(val OptNilInt64) {
	s.ID = val
}

// SetInvitedToJoin sets the value of InvitedToJoin.
func (s *Group) SetInvitedToJoin(val OptNilBool) {
	s.InvitedToJoin = val
}

// SetIsJoined sets the value of IsJoined.
func (s *Group) SetIsJoined(val OptNilBool) {
	s.IsJoined = val
}

// SetIsPending sets the value of IsPending.
func (s *Group) SetIsPending(val OptNilBool) {
	s.IsPending = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *Group) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetIsRelated sets the value of IsRelated.
func (s *Group) SetIsRelated(val OptNilBool) {
	s.IsRelated = val
}

// SetJoinedCommunityCampaign sets the value of JoinedCommunityCampaign.
func (s *Group) SetJoinedCommunityCampaign(val OptNilBool) {
	s.JoinedCommunityCampaign = val
}

// SetModeratorIds sets the value of ModeratorIds.
func (s *Group) SetModeratorIds(val OptNilInt64Array) {
	s.ModeratorIds = val
}

// SetOnlyMobileVerified sets the value of OnlyMobileVerified.
func (s *Group) SetOnlyMobileVerified(val OptNilBool) {
	s.OnlyMobileVerified = val
}

// SetOnlyVerifiedAge sets the value of OnlyVerifiedAge.
func (s *Group) SetOnlyVerifiedAge(val OptNilBool) {
	s.OnlyVerifiedAge = val
}

// SetOwner sets the value of Owner.
func (s *Group) SetOwner(val *RealmUser) {
	s.Owner = val
}

// SetPendingCount sets the value of PendingCount.
func (s *Group) SetPendingCount(val OptNilInt64) {
	s.PendingCount = val
}

// SetPendingDeputizeIds sets the value of PendingDeputizeIds.
func (s *Group) SetPendingDeputizeIds(val OptNilInt64Array) {
	s.PendingDeputizeIds = val
}

// SetPendingTransferID sets the value of PendingTransferID.
func (s *Group) SetPendingTransferID(val OptNilInt64) {
	s.PendingTransferID = val
}

// SetPostsCount sets the value of PostsCount.
func (s *Group) SetPostsCount(val OptNilInt64) {
	s.PostsCount = val
}

// SetRelatedCount sets the value of RelatedCount.
func (s *Group) SetRelatedCount(val OptNilInt64) {
	s.RelatedCount = val
}

// SetSafeMode sets the value of SafeMode.
func (s *Group) SetSafeMode(val OptNilBool) {
	s.SafeMode = val
}

// SetSecret sets the value of Secret.
func (s *Group) SetSecret(val OptNilBool) {
	s.Secret = val
}

// SetSeizable sets the value of Seizable.
func (s *Group) SetSeizable(val OptNilBool) {
	s.Seizable = val
}

// SetSeizableBefore sets the value of SeizableBefore.
func (s *Group) SetSeizableBefore(val OptNilInt64) {
	s.SeizableBefore = val
}

// SetSubCategoryID sets the value of SubCategoryID.
func (s *Group) SetSubCategoryID(val OptNilInt64) {
	s.SubCategoryID = val
}

// SetThreadsCount sets the value of ThreadsCount.
func (s *Group) SetThreadsCount(val OptNilInt64) {
	s.ThreadsCount = val
}

// SetTitle sets the value of Title.
func (s *Group) SetTitle(val OptNilString) {
	s.Title = val
}

// SetTopic sets the value of Topic.
func (s *Group) SetTopic(val OptNilString) {
	s.Topic = val
}

// SetUnreadCounts sets the value of UnreadCounts.
func (s *Group) SetUnreadCounts(val OptNilInt64) {
	s.UnreadCounts = val
}

// SetUnreadThreadsCount sets the value of UnreadThreadsCount.
func (s *Group) SetUnreadThreadsCount(val OptNilInt64) {
	s.UnreadThreadsCount = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Group) SetUpdatedAt(val OptNilInt64) {
	s.UpdatedAt = val
}

// SetUserID sets the value of UserID.
func (s *Group) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// SetViewsCount sets the value of ViewsCount.
func (s *Group) SetViewsCount(val OptNilInt64) {
	s.ViewsCount = val
}

// SetWalkthroughRequested sets the value of WalkthroughRequested.
func (s *Group) SetWalkthroughRequested(val OptNilBool) {
	s.WalkthroughRequested = val
}

// Ref: #/components/schemas/GroupCategoriesResponse
type GroupCategoriesResponse struct {
	GroupCategories OptNilGroupCategoryArray `json:"group_categories"`
}

// GetGroupCategories returns the value of GroupCategories.
func (s *GroupCategoriesResponse) GetGroupCategories() OptNilGroupCategoryArray {
	return s.GroupCategories
}

// SetGroupCategories sets the value of GroupCategories.
func (s *GroupCategoriesResponse) SetGroupCategories(val OptNilGroupCategoryArray) {
	s.GroupCategories = val
}

// Ref: #/components/schemas/GroupCategory
type GroupCategory struct {
	Icon OptNilString `json:"icon"`
	ID   OptNilInt64  `json:"id"`
	Name OptNilString `json:"name"`
	Rank OptNilInt32  `json:"rank"`
}

// GetIcon returns the value of Icon.
func (s *GroupCategory) GetIcon() OptNilString {
	return s.Icon
}

// GetID returns the value of ID.
func (s *GroupCategory) GetID() OptNilInt64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *GroupCategory) GetName() OptNilString {
	return s.Name
}

// GetRank returns the value of Rank.
func (s *GroupCategory) GetRank() OptNilInt32 {
	return s.Rank
}

// SetIcon sets the value of Icon.
func (s *GroupCategory) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetID sets the value of ID.
func (s *GroupCategory) SetID(val OptNilInt64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *GroupCategory) SetName(val OptNilString) {
	s.Name = val
}

// SetRank sets the value of Rank.
func (s *GroupCategory) SetRank(val OptNilInt32) {
	s.Rank = val
}

// Ref: #/components/schemas/GroupGiftHistory
type GroupGiftHistory struct {
	GiftsCount   OptNilGiftCountArray `json:"gifts_count"`
	ReceivedDate OptNilInt64          `json:"received_date"`
	User         *RealmUser           `json:"user"`
}

// GetGiftsCount returns the value of GiftsCount.
func (s *GroupGiftHistory) GetGiftsCount() OptNilGiftCountArray {
	return s.GiftsCount
}

// GetReceivedDate returns the value of ReceivedDate.
func (s *GroupGiftHistory) GetReceivedDate() OptNilInt64 {
	return s.ReceivedDate
}

// GetUser returns the value of User.
func (s *GroupGiftHistory) GetUser() *RealmUser {
	return s.User
}

// SetGiftsCount sets the value of GiftsCount.
func (s *GroupGiftHistory) SetGiftsCount(val OptNilGiftCountArray) {
	s.GiftsCount = val
}

// SetReceivedDate sets the value of ReceivedDate.
func (s *GroupGiftHistory) SetReceivedDate(val OptNilInt64) {
	s.ReceivedDate = val
}

// SetUser sets the value of User.
func (s *GroupGiftHistory) SetUser(val *RealmUser) {
	s.User = val
}

// Ref: #/components/schemas/GroupGiftHistoryResponse
type GroupGiftHistoryResponse struct {
	GiftHistory   OptNilGroupGiftHistoryArray `json:"gift_history"`
	NextPageValue OptNilString                `json:"next_page_value"`
}

// GetGiftHistory returns the value of GiftHistory.
func (s *GroupGiftHistoryResponse) GetGiftHistory() OptNilGroupGiftHistoryArray {
	return s.GiftHistory
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GroupGiftHistoryResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// SetGiftHistory sets the value of GiftHistory.
func (s *GroupGiftHistoryResponse) SetGiftHistory(val OptNilGroupGiftHistoryArray) {
	s.GiftHistory = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GroupGiftHistoryResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// Ref: #/components/schemas/GroupMuteUsersResponse
type GroupMuteUsersResponse struct {
	GroupUsers OptNilGroupUserArray `json:"group_users"`
	NextCursor OptNilInt64          `json:"next_cursor"`
	Total      OptNilInt32          `json:"total"`
}

// GetGroupUsers returns the value of GroupUsers.
func (s *GroupMuteUsersResponse) GetGroupUsers() OptNilGroupUserArray {
	return s.GroupUsers
}

// GetNextCursor returns the value of NextCursor.
func (s *GroupMuteUsersResponse) GetNextCursor() OptNilInt64 {
	return s.NextCursor
}

// GetTotal returns the value of Total.
func (s *GroupMuteUsersResponse) GetTotal() OptNilInt32 {
	return s.Total
}

// SetGroupUsers sets the value of GroupUsers.
func (s *GroupMuteUsersResponse) SetGroupUsers(val OptNilGroupUserArray) {
	s.GroupUsers = val
}

// SetNextCursor sets the value of NextCursor.
func (s *GroupMuteUsersResponse) SetNextCursor(val OptNilInt64) {
	s.NextCursor = val
}

// SetTotal sets the value of Total.
func (s *GroupMuteUsersResponse) SetTotal(val OptNilInt32) {
	s.Total = val
}

// Ref: #/components/schemas/GroupNotificationSettingsResponse
type GroupNotificationSettingsResponse struct {
	Setting OptSetting `json:"setting"`
}

// GetSetting returns the value of Setting.
func (s *GroupNotificationSettingsResponse) GetSetting() OptSetting {
	return s.Setting
}

// SetSetting sets the value of Setting.
func (s *GroupNotificationSettingsResponse) SetSetting(val OptSetting) {
	s.Setting = val
}

// Ref: #/components/schemas/GroupResponse
type GroupResponse struct {
	Group OptGroup `json:"group"`
}

// GetGroup returns the value of Group.
func (s *GroupResponse) GetGroup() OptGroup {
	return s.Group
}

// SetGroup sets the value of Group.
func (s *GroupResponse) SetGroup(val OptGroup) {
	s.Group = val
}

// Ref: #/components/schemas/GroupThreadListResponse
type GroupThreadListResponse struct {
	NextPageValue OptNilString          `json:"next_page_value"`
	Threads       OptNilThreadInfoArray `json:"threads"`
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GroupThreadListResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetThreads returns the value of Threads.
func (s *GroupThreadListResponse) GetThreads() OptNilThreadInfoArray {
	return s.Threads
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GroupThreadListResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetThreads sets the value of Threads.
func (s *GroupThreadListResponse) SetThreads(val OptNilThreadInfoArray) {
	s.Threads = val
}

// Ref: #/components/schemas/GroupUser
type GroupUser struct {
	IsModerator     OptNilBool   `json:"is_moderator"`
	IsMute          OptNilBool   `json:"is_mute"`
	PendingDeputize OptNilBool   `json:"pending_deputize"`
	PendingTransfer OptNilBool   `json:"pending_transfer"`
	Title           OptNilString `json:"title"`
	User            *RealmUser   `json:"user"`
}

// GetIsModerator returns the value of IsModerator.
func (s *GroupUser) GetIsModerator() OptNilBool {
	return s.IsModerator
}

// GetIsMute returns the value of IsMute.
func (s *GroupUser) GetIsMute() OptNilBool {
	return s.IsMute
}

// GetPendingDeputize returns the value of PendingDeputize.
func (s *GroupUser) GetPendingDeputize() OptNilBool {
	return s.PendingDeputize
}

// GetPendingTransfer returns the value of PendingTransfer.
func (s *GroupUser) GetPendingTransfer() OptNilBool {
	return s.PendingTransfer
}

// GetTitle returns the value of Title.
func (s *GroupUser) GetTitle() OptNilString {
	return s.Title
}

// GetUser returns the value of User.
func (s *GroupUser) GetUser() *RealmUser {
	return s.User
}

// SetIsModerator sets the value of IsModerator.
func (s *GroupUser) SetIsModerator(val OptNilBool) {
	s.IsModerator = val
}

// SetIsMute sets the value of IsMute.
func (s *GroupUser) SetIsMute(val OptNilBool) {
	s.IsMute = val
}

// SetPendingDeputize sets the value of PendingDeputize.
func (s *GroupUser) SetPendingDeputize(val OptNilBool) {
	s.PendingDeputize = val
}

// SetPendingTransfer sets the value of PendingTransfer.
func (s *GroupUser) SetPendingTransfer(val OptNilBool) {
	s.PendingTransfer = val
}

// SetTitle sets the value of Title.
func (s *GroupUser) SetTitle(val OptNilString) {
	s.Title = val
}

// SetUser sets the value of User.
func (s *GroupUser) SetUser(val *RealmUser) {
	s.User = val
}

// Ref: #/components/schemas/GroupUserResponse
type GroupUserResponse struct {
	GroupUser OptGroupUser `json:"group_user"`
}

// GetGroupUser returns the value of GroupUser.
func (s *GroupUserResponse) GetGroupUser() OptGroupUser {
	return s.GroupUser
}

// SetGroupUser sets the value of GroupUser.
func (s *GroupUserResponse) SetGroupUser(val OptGroupUser) {
	s.GroupUser = val
}

// Ref: #/components/schemas/GroupUsersResponse
type GroupUsersResponse struct {
	GroupUsers OptNilGroupUserArray `json:"group_users"`
}

// GetGroupUsers returns the value of GroupUsers.
func (s *GroupUsersResponse) GetGroupUsers() OptNilGroupUserArray {
	return s.GroupUsers
}

// SetGroupUsers sets the value of GroupUsers.
func (s *GroupUsersResponse) SetGroupUsers(val OptNilGroupUserArray) {
	s.GroupUsers = val
}

// Ref: #/components/schemas/GroupsRelatedResponse
type GroupsRelatedResponse struct {
	Groups        OptNilGroupArray `json:"groups"`
	NextPageValue OptNilString     `json:"next_page_value"`
}

// GetGroups returns the value of Groups.
func (s *GroupsRelatedResponse) GetGroups() OptNilGroupArray {
	return s.Groups
}

// GetNextPageValue returns the value of NextPageValue.
func (s *GroupsRelatedResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// SetGroups sets the value of Groups.
func (s *GroupsRelatedResponse) SetGroups(val OptNilGroupArray) {
	s.Groups = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *GroupsRelatedResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// Ref: #/components/schemas/GroupsResponse
type GroupsResponse struct {
	Groups       OptNilGroupArray `json:"groups"`
	PinnedGroups OptNilGroupArray `json:"pinned_groups"`
}

// GetGroups returns the value of Groups.
func (s *GroupsResponse) GetGroups() OptNilGroupArray {
	return s.Groups
}

// GetPinnedGroups returns the value of PinnedGroups.
func (s *GroupsResponse) GetPinnedGroups() OptNilGroupArray {
	return s.PinnedGroups
}

// SetGroups sets the value of Groups.
func (s *GroupsResponse) SetGroups(val OptNilGroupArray) {
	s.Groups = val
}

// SetPinnedGroups sets the value of PinnedGroups.
func (s *GroupsResponse) SetPinnedGroups(val OptNilGroupArray) {
	s.PinnedGroups = val
}

// Ref: #/components/schemas/HiddenResponse
type HiddenResponse struct {
	HiddenUsers   OptNilRealmUserArray `json:"hidden_users"`
	Limit         OptNilInt32          `json:"limit"`
	NextPageValue OptNilString         `json:"next_page_value"`
	TotalCount    OptNilInt32          `json:"total_count"`
}

// GetHiddenUsers returns the value of HiddenUsers.
func (s *HiddenResponse) GetHiddenUsers() OptNilRealmUserArray {
	return s.HiddenUsers
}

// GetLimit returns the value of Limit.
func (s *HiddenResponse) GetLimit() OptNilInt32 {
	return s.Limit
}

// GetNextPageValue returns the value of NextPageValue.
func (s *HiddenResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetTotalCount returns the value of TotalCount.
func (s *HiddenResponse) GetTotalCount() OptNilInt32 {
	return s.TotalCount
}

// SetHiddenUsers sets the value of HiddenUsers.
func (s *HiddenResponse) SetHiddenUsers(val OptNilRealmUserArray) {
	s.HiddenUsers = val
}

// SetLimit sets the value of Limit.
func (s *HiddenResponse) SetLimit(val OptNilInt32) {
	s.Limit = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *HiddenResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetTotalCount sets the value of TotalCount.
func (s *HiddenResponse) SetTotalCount(val OptNilInt32) {
	s.TotalCount = val
}

// HideChatsNoContent is response for HideChats operation.
type HideChatsNoContent struct{}

type HideChatsReq struct {
	ChatRoomID OptInt64 `json:"chat_room_id"`
}

// GetChatRoomID returns the value of ChatRoomID.
func (s *HideChatsReq) GetChatRoomID() OptInt64 {
	return s.ChatRoomID
}

// SetChatRoomID sets the value of ChatRoomID.
func (s *HideChatsReq) SetChatRoomID(val OptInt64) {
	s.ChatRoomID = val
}

// HideUsersNoContent is response for HideUsers operation.
type HideUsersNoContent struct{}

type HideUsersReq struct {
	UserID OptInt64 `json:"user_id"`
}

// GetUserID returns the value of UserID.
func (s *HideUsersReq) GetUserID() OptInt64 {
	return s.UserID
}

// SetUserID sets the value of UserID.
func (s *HideUsersReq) SetUserID(val OptInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/HimaUsersResponse
type HimaUsersResponse struct {
	HimaUsers OptNilUserWrapperArray `json:"hima_users"`
}

// GetHimaUsers returns the value of HimaUsers.
func (s *HimaUsersResponse) GetHimaUsers() OptNilUserWrapperArray {
	return s.HimaUsers
}

// SetHimaUsers sets the value of HimaUsers.
func (s *HimaUsersResponse) SetHimaUsers(val OptNilUserWrapperArray) {
	s.HimaUsers = val
}

// Ref: #/components/schemas/Interest
type Interest struct {
	Icon     OptNilString `json:"icon"`
	ID       OptNilInt64  `json:"id"`
	Name     OptNilString `json:"name"`
	Selected OptNilBool   `json:"selected"`
}

// GetIcon returns the value of Icon.
func (s *Interest) GetIcon() OptNilString {
	return s.Icon
}

// GetID returns the value of ID.
func (s *Interest) GetID() OptNilInt64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Interest) GetName() OptNilString {
	return s.Name
}

// GetSelected returns the value of Selected.
func (s *Interest) GetSelected() OptNilBool {
	return s.Selected
}

// SetIcon sets the value of Icon.
func (s *Interest) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetID sets the value of ID.
func (s *Interest) SetID(val OptNilInt64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Interest) SetName(val OptNilString) {
	s.Name = val
}

// SetSelected sets the value of Selected.
func (s *Interest) SetSelected(val OptNilBool) {
	s.Selected = val
}

// InviteToCallNoContent is response for InviteToCall operation.
type InviteToCallNoContent struct{}

type InviteToCallReq struct {
	ChatRoomID OptInt64  `json:"chat_room_id"`
	RoomID     OptInt64  `json:"room_id"`
	RoomURL    OptString `json:"room_url"`
}

// GetChatRoomID returns the value of ChatRoomID.
func (s *InviteToCallReq) GetChatRoomID() OptInt64 {
	return s.ChatRoomID
}

// GetRoomID returns the value of RoomID.
func (s *InviteToCallReq) GetRoomID() OptInt64 {
	return s.RoomID
}

// GetRoomURL returns the value of RoomURL.
func (s *InviteToCallReq) GetRoomURL() OptString {
	return s.RoomURL
}

// SetChatRoomID sets the value of ChatRoomID.
func (s *InviteToCallReq) SetChatRoomID(val OptInt64) {
	s.ChatRoomID = val
}

// SetRoomID sets the value of RoomID.
func (s *InviteToCallReq) SetRoomID(val OptInt64) {
	s.RoomID = val
}

// SetRoomURL sets the value of RoomURL.
func (s *InviteToCallReq) SetRoomURL(val OptString) {
	s.RoomURL = val
}

// InviteToChatRoomNoContent is response for InviteToChatRoom operation.
type InviteToChatRoomNoContent struct{}

type InviteToChatRoomReq struct {
	WithUserIds []int64 `json:"with_user_ids[]"`
}

// GetWithUserIds returns the value of WithUserIds.
func (s *InviteToChatRoomReq) GetWithUserIds() []int64 {
	return s.WithUserIds
}

// SetWithUserIds sets the value of WithUserIds.
func (s *InviteToChatRoomReq) SetWithUserIds(val []int64) {
	s.WithUserIds = val
}

// InviteToConferenceCallNoContent is response for InviteToConferenceCall operation.
type InviteToConferenceCallNoContent struct{}

type InviteToConferenceCallReq struct {
	UserIds []int64 `json:"user_ids[]"`
}

// GetUserIds returns the value of UserIds.
func (s *InviteToConferenceCallReq) GetUserIds() []int64 {
	return s.UserIds
}

// SetUserIds sets the value of UserIds.
func (s *InviteToConferenceCallReq) SetUserIds(val []int64) {
	s.UserIds = val
}

// InviteToGroupNoContent is response for InviteToGroup operation.
type InviteToGroupNoContent struct{}

type InviteToGroupReq struct {
	UserIds []int64 `json:"user_ids[]"`
}

// GetUserIds returns the value of UserIds.
func (s *InviteToGroupReq) GetUserIds() []int64 {
	return s.UserIds
}

// SetUserIds sets the value of UserIds.
func (s *InviteToGroupReq) SetUserIds(val []int64) {
	s.UserIds = val
}

// JoinGroupNoContent is response for JoinGroup operation.
type JoinGroupNoContent struct{}

// KickFromChatRoomNoContent is response for KickFromChatRoom operation.
type KickFromChatRoomNoContent struct{}

type KickFromChatRoomReq struct {
	WithUserIds []int64 `json:"with_user_ids[]"`
}

// GetWithUserIds returns the value of WithUserIds.
func (s *KickFromChatRoomReq) GetWithUserIds() []int64 {
	return s.WithUserIds
}

// SetWithUserIds sets the value of WithUserIds.
func (s *KickFromChatRoomReq) SetWithUserIds(val []int64) {
	s.WithUserIds = val
}

type KickFromConferenceCallReq struct {
	Ban  OptBool   `json:"ban"`
	UUID OptString `json:"uuid"`
}

// GetBan returns the value of Ban.
func (s *KickFromConferenceCallReq) GetBan() OptBool {
	return s.Ban
}

// GetUUID returns the value of UUID.
func (s *KickFromConferenceCallReq) GetUUID() OptString {
	return s.UUID
}

// SetBan sets the value of Ban.
func (s *KickFromConferenceCallReq) SetBan(val OptBool) {
	s.Ban = val
}

// SetUUID sets the value of UUID.
func (s *KickFromConferenceCallReq) SetUUID(val OptString) {
	s.UUID = val
}

// LeaveAgoraChannelNoContent is response for LeaveAgoraChannel operation.
type LeaveAgoraChannelNoContent struct{}

type LeaveAgoraChannelReq struct {
	ConferenceID OptInt64 `json:"conference_id"`
	UserID       OptInt64 `json:"user_id"`
}

// GetConferenceID returns the value of ConferenceID.
func (s *LeaveAgoraChannelReq) GetConferenceID() OptInt64 {
	return s.ConferenceID
}

// GetUserID returns the value of UserID.
func (s *LeaveAgoraChannelReq) GetUserID() OptInt64 {
	return s.UserID
}

// SetConferenceID sets the value of ConferenceID.
func (s *LeaveAgoraChannelReq) SetConferenceID(val OptInt64) {
	s.ConferenceID = val
}

// SetUserID sets the value of UserID.
func (s *LeaveAgoraChannelReq) SetUserID(val OptInt64) {
	s.UserID = val
}

// LeaveConferenceCallNoContent is response for LeaveConferenceCall operation.
type LeaveConferenceCallNoContent struct{}

type LeaveConferenceCallReq struct {
	CallSid          OptString   `json:"call_sid"`
	ConferenceID     OptInt64    `json:"conference_id"`
	Duration         OptNilInt64 `json:"duration"`
	TotalUsersInCall OptNilInt32 `json:"total_users_in_call"`
}

// GetCallSid returns the value of CallSid.
func (s *LeaveConferenceCallReq) GetCallSid() OptString {
	return s.CallSid
}

// GetConferenceID returns the value of ConferenceID.
func (s *LeaveConferenceCallReq) GetConferenceID() OptInt64 {
	return s.ConferenceID
}

// GetDuration returns the value of Duration.
func (s *LeaveConferenceCallReq) GetDuration() OptNilInt64 {
	return s.Duration
}

// GetTotalUsersInCall returns the value of TotalUsersInCall.
func (s *LeaveConferenceCallReq) GetTotalUsersInCall() OptNilInt32 {
	return s.TotalUsersInCall
}

// SetCallSid sets the value of CallSid.
func (s *LeaveConferenceCallReq) SetCallSid(val OptString) {
	s.CallSid = val
}

// SetConferenceID sets the value of ConferenceID.
func (s *LeaveConferenceCallReq) SetConferenceID(val OptInt64) {
	s.ConferenceID = val
}

// SetDuration sets the value of Duration.
func (s *LeaveConferenceCallReq) SetDuration(val OptNilInt64) {
	s.Duration = val
}

// SetTotalUsersInCall sets the value of TotalUsersInCall.
func (s *LeaveConferenceCallReq) SetTotalUsersInCall(val OptNilInt32) {
	s.TotalUsersInCall = val
}

// LeaveGroupNoContent is response for LeaveGroup operation.
type LeaveGroupNoContent struct{}

type LikePostsReq struct {
	PostIds []int64 `json:"post_ids[]"`
}

// GetPostIds returns the value of PostIds.
func (s *LikePostsReq) GetPostIds() []int64 {
	return s.PostIds
}

// SetPostIds sets the value of PostIds.
func (s *LikePostsReq) SetPostIds(val []int64) {
	s.PostIds = val
}

// Ref: #/components/schemas/LikePostsResponse
type LikePostsResponse struct {
	LikeIds OptNilInt64Array `json:"like_ids"`
}

// GetLikeIds returns the value of LikeIds.
func (s *LikePostsResponse) GetLikeIds() OptNilInt64Array {
	return s.LikeIds
}

// SetLikeIds sets the value of LikeIds.
func (s *LikePostsResponse) SetLikeIds(val OptNilInt64Array) {
	s.LikeIds = val
}

// Ref: #/components/schemas/LoginEmailUserRequest
type LoginEmailUserRequest struct {
	APIKey    OptNilString `json:"api_key"`
	Email     OptNilString `json:"email"`
	Password  OptNilString `json:"password"`
	TwoFACode OptNilString `json:"two_f_a_code"`
	UUID      OptNilString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *LoginEmailUserRequest) GetAPIKey() OptNilString {
	return s.APIKey
}

// GetEmail returns the value of Email.
func (s *LoginEmailUserRequest) GetEmail() OptNilString {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *LoginEmailUserRequest) GetPassword() OptNilString {
	return s.Password
}

// GetTwoFACode returns the value of TwoFACode.
func (s *LoginEmailUserRequest) GetTwoFACode() OptNilString {
	return s.TwoFACode
}

// GetUUID returns the value of UUID.
func (s *LoginEmailUserRequest) GetUUID() OptNilString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *LoginEmailUserRequest) SetAPIKey(val OptNilString) {
	s.APIKey = val
}

// SetEmail sets the value of Email.
func (s *LoginEmailUserRequest) SetEmail(val OptNilString) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *LoginEmailUserRequest) SetPassword(val OptNilString) {
	s.Password = val
}

// SetTwoFACode sets the value of TwoFACode.
func (s *LoginEmailUserRequest) SetTwoFACode(val OptNilString) {
	s.TwoFACode = val
}

// SetUUID sets the value of UUID.
func (s *LoginEmailUserRequest) SetUUID(val OptNilString) {
	s.UUID = val
}

// Ref: #/components/schemas/LoginUpdateResponse
type LoginUpdateResponse struct {
	AccessToken  OptNilString `json:"access_token"`
	ExpiresIn    OptNilInt64  `json:"expires_in"`
	RefreshToken OptNilString `json:"refresh_token"`
	UserID       OptNilInt64  `json:"user_id"`
	Username     OptNilString `json:"username"`
}

// GetAccessToken returns the value of AccessToken.
func (s *LoginUpdateResponse) GetAccessToken() OptNilString {
	return s.AccessToken
}

// GetExpiresIn returns the value of ExpiresIn.
func (s *LoginUpdateResponse) GetExpiresIn() OptNilInt64 {
	return s.ExpiresIn
}

// GetRefreshToken returns the value of RefreshToken.
func (s *LoginUpdateResponse) GetRefreshToken() OptNilString {
	return s.RefreshToken
}

// GetUserID returns the value of UserID.
func (s *LoginUpdateResponse) GetUserID() OptNilInt64 {
	return s.UserID
}

// GetUsername returns the value of Username.
func (s *LoginUpdateResponse) GetUsername() OptNilString {
	return s.Username
}

// SetAccessToken sets the value of AccessToken.
func (s *LoginUpdateResponse) SetAccessToken(val OptNilString) {
	s.AccessToken = val
}

// SetExpiresIn sets the value of ExpiresIn.
func (s *LoginUpdateResponse) SetExpiresIn(val OptNilInt64) {
	s.ExpiresIn = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *LoginUpdateResponse) SetRefreshToken(val OptNilString) {
	s.RefreshToken = val
}

// SetUserID sets the value of UserID.
func (s *LoginUpdateResponse) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// SetUsername sets the value of Username.
func (s *LoginUpdateResponse) SetUsername(val OptNilString) {
	s.Username = val
}

// Ref: #/components/schemas/LoginUserResponse
type LoginUserResponse struct {
	AccessToken  OptNilString `json:"access_token"`
	ExpiresIn    OptNilInt64  `json:"expires_in"`
	IsNew        OptNilBool   `json:"is_new"`
	RefreshToken OptNilString `json:"refresh_token"`
	SnsInfo      OptSnsInfo   `json:"sns_info"`
	UserID       OptNilInt64  `json:"user_id"`
	Username     OptNilString `json:"username"`
}

// GetAccessToken returns the value of AccessToken.
func (s *LoginUserResponse) GetAccessToken() OptNilString {
	return s.AccessToken
}

// GetExpiresIn returns the value of ExpiresIn.
func (s *LoginUserResponse) GetExpiresIn() OptNilInt64 {
	return s.ExpiresIn
}

// GetIsNew returns the value of IsNew.
func (s *LoginUserResponse) GetIsNew() OptNilBool {
	return s.IsNew
}

// GetRefreshToken returns the value of RefreshToken.
func (s *LoginUserResponse) GetRefreshToken() OptNilString {
	return s.RefreshToken
}

// GetSnsInfo returns the value of SnsInfo.
func (s *LoginUserResponse) GetSnsInfo() OptSnsInfo {
	return s.SnsInfo
}

// GetUserID returns the value of UserID.
func (s *LoginUserResponse) GetUserID() OptNilInt64 {
	return s.UserID
}

// GetUsername returns the value of Username.
func (s *LoginUserResponse) GetUsername() OptNilString {
	return s.Username
}

// SetAccessToken sets the value of AccessToken.
func (s *LoginUserResponse) SetAccessToken(val OptNilString) {
	s.AccessToken = val
}

// SetExpiresIn sets the value of ExpiresIn.
func (s *LoginUserResponse) SetExpiresIn(val OptNilInt64) {
	s.ExpiresIn = val
}

// SetIsNew sets the value of IsNew.
func (s *LoginUserResponse) SetIsNew(val OptNilBool) {
	s.IsNew = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *LoginUserResponse) SetRefreshToken(val OptNilString) {
	s.RefreshToken = val
}

// SetSnsInfo sets the value of SnsInfo.
func (s *LoginUserResponse) SetSnsInfo(val OptSnsInfo) {
	s.SnsInfo = val
}

// SetUserID sets the value of UserID.
func (s *LoginUserResponse) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// SetUsername sets the value of Username.
func (s *LoginUserResponse) SetUsername(val OptNilString) {
	s.Username = val
}

// LogoutNoContent is response for Logout operation.
type LogoutNoContent struct{}

type LogoutReq struct {
	UUID OptString `json:"uuid"`
}

// GetUUID returns the value of UUID.
func (s *LogoutReq) GetUUID() OptString {
	return s.UUID
}

// SetUUID sets the value of UUID.
func (s *LogoutReq) SetUUID(val OptString) {
	s.UUID = val
}

// Ref: #/components/schemas/MessageResponse
type MessageResponse struct {
	ConferenceCall OptRealmConferenceCall `json:"conference_call"`
	ID             OptNilInt64            `json:"id"`
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *MessageResponse) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetID returns the value of ID.
func (s *MessageResponse) GetID() OptNilInt64 {
	return s.ID
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *MessageResponse) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetID sets the value of ID.
func (s *MessageResponse) SetID(val OptNilInt64) {
	s.ID = val
}

// Ref: #/components/schemas/MessageTag
type MessageTag struct {
	Length OptNilInt32  `json:"length"`
	Offset OptNilInt32  `json:"offset"`
	Type   OptNilString `json:"type"`
	UserID OptNilInt64  `json:"user_id"`
}

// GetLength returns the value of Length.
func (s *MessageTag) GetLength() OptNilInt32 {
	return s.Length
}

// GetOffset returns the value of Offset.
func (s *MessageTag) GetOffset() OptNilInt32 {
	return s.Offset
}

// GetType returns the value of Type.
func (s *MessageTag) GetType() OptNilString {
	return s.Type
}

// GetUserID returns the value of UserID.
func (s *MessageTag) GetUserID() OptNilInt64 {
	return s.UserID
}

// SetLength sets the value of Length.
func (s *MessageTag) SetLength(val OptNilInt32) {
	s.Length = val
}

// SetOffset sets the value of Offset.
func (s *MessageTag) SetOffset(val OptNilInt32) {
	s.Offset = val
}

// SetType sets the value of Type.
func (s *MessageTag) SetType(val OptNilString) {
	s.Type = val
}

// SetUserID sets the value of UserID.
func (s *MessageTag) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/MessageType
type MessageType string

const (
	MessageTypeText              MessageType = "text"
	MessageTypeImage             MessageType = "image"
	MessageTypeEternalImage      MessageType = "eternal_image"
	MessageTypeVideo             MessageType = "video"
	MessageTypeEternalVideo      MessageType = "eternal_video"
	MessageTypeScreenshotWarning MessageType = "screenshot_warning"
	MessageTypeGIF               MessageType = "gif"
	MessageTypeSticker           MessageType = "sticker"
	MessageTypeIndividualCall    MessageType = "individual_call"
	MessageTypeDeleted           MessageType = "deleted"
	MessageTypeUserJoined        MessageType = "user_joined"
	MessageTypeUserLeave         MessageType = "user_leave"
	MessageTypeOwnerKick         MessageType = "owner_kick"
)

// AllValues returns all MessageType values.
func (MessageType) AllValues() []MessageType {
	return []MessageType{
		MessageTypeText,
		MessageTypeImage,
		MessageTypeEternalImage,
		MessageTypeVideo,
		MessageTypeEternalVideo,
		MessageTypeScreenshotWarning,
		MessageTypeGIF,
		MessageTypeSticker,
		MessageTypeIndividualCall,
		MessageTypeDeleted,
		MessageTypeUserJoined,
		MessageTypeUserLeave,
		MessageTypeOwnerKick,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s MessageType) MarshalText() ([]byte, error) {
	switch s {
	case MessageTypeText:
		return []byte(s), nil
	case MessageTypeImage:
		return []byte(s), nil
	case MessageTypeEternalImage:
		return []byte(s), nil
	case MessageTypeVideo:
		return []byte(s), nil
	case MessageTypeEternalVideo:
		return []byte(s), nil
	case MessageTypeScreenshotWarning:
		return []byte(s), nil
	case MessageTypeGIF:
		return []byte(s), nil
	case MessageTypeSticker:
		return []byte(s), nil
	case MessageTypeIndividualCall:
		return []byte(s), nil
	case MessageTypeDeleted:
		return []byte(s), nil
	case MessageTypeUserJoined:
		return []byte(s), nil
	case MessageTypeUserLeave:
		return []byte(s), nil
	case MessageTypeOwnerKick:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *MessageType) UnmarshalText(data []byte) error {
	switch MessageType(data) {
	case MessageTypeText:
		*s = MessageTypeText
		return nil
	case MessageTypeImage:
		*s = MessageTypeImage
		return nil
	case MessageTypeEternalImage:
		*s = MessageTypeEternalImage
		return nil
	case MessageTypeVideo:
		*s = MessageTypeVideo
		return nil
	case MessageTypeEternalVideo:
		*s = MessageTypeEternalVideo
		return nil
	case MessageTypeScreenshotWarning:
		*s = MessageTypeScreenshotWarning
		return nil
	case MessageTypeGIF:
		*s = MessageTypeGIF
		return nil
	case MessageTypeSticker:
		*s = MessageTypeSticker
		return nil
	case MessageTypeIndividualCall:
		*s = MessageTypeIndividualCall
		return nil
	case MessageTypeDeleted:
		*s = MessageTypeDeleted
		return nil
	case MessageTypeUserJoined:
		*s = MessageTypeUserJoined
		return nil
	case MessageTypeUserLeave:
		*s = MessageTypeUserLeave
		return nil
	case MessageTypeOwnerKick:
		*s = MessageTypeOwnerKick
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/MessagesResponse
type MessagesResponse struct {
	Messages OptNilRealmMessageArray `json:"messages"`
}

// GetMessages returns the value of Messages.
func (s *MessagesResponse) GetMessages() OptNilRealmMessageArray {
	return s.Messages
}

// SetMessages sets the value of Messages.
func (s *MessagesResponse) SetMessages(val OptNilRealmMessageArray) {
	s.Messages = val
}

// Ref: #/components/schemas/Metadata
type Metadata struct {
	Anonymous      OptNilBool   `json:"anonymous"`
	Body           OptNilString `json:"body"`
	BulkInvitation OptNilBool   `json:"bulk_invitation"`
	ContentPreview OptNilString `json:"content_preview"`
	GroupTopic     OptNilString `json:"group_topic"`
	Title          OptNilString `json:"title"`
	URL            OptNilString `json:"url"`
}

// GetAnonymous returns the value of Anonymous.
func (s *Metadata) GetAnonymous() OptNilBool {
	return s.Anonymous
}

// GetBody returns the value of Body.
func (s *Metadata) GetBody() OptNilString {
	return s.Body
}

// GetBulkInvitation returns the value of BulkInvitation.
func (s *Metadata) GetBulkInvitation() OptNilBool {
	return s.BulkInvitation
}

// GetContentPreview returns the value of ContentPreview.
func (s *Metadata) GetContentPreview() OptNilString {
	return s.ContentPreview
}

// GetGroupTopic returns the value of GroupTopic.
func (s *Metadata) GetGroupTopic() OptNilString {
	return s.GroupTopic
}

// GetTitle returns the value of Title.
func (s *Metadata) GetTitle() OptNilString {
	return s.Title
}

// GetURL returns the value of URL.
func (s *Metadata) GetURL() OptNilString {
	return s.URL
}

// SetAnonymous sets the value of Anonymous.
func (s *Metadata) SetAnonymous(val OptNilBool) {
	s.Anonymous = val
}

// SetBody sets the value of Body.
func (s *Metadata) SetBody(val OptNilString) {
	s.Body = val
}

// SetBulkInvitation sets the value of BulkInvitation.
func (s *Metadata) SetBulkInvitation(val OptNilBool) {
	s.BulkInvitation = val
}

// SetContentPreview sets the value of ContentPreview.
func (s *Metadata) SetContentPreview(val OptNilString) {
	s.ContentPreview = val
}

// SetGroupTopic sets the value of GroupTopic.
func (s *Metadata) SetGroupTopic(val OptNilString) {
	s.GroupTopic = val
}

// SetTitle sets the value of Title.
func (s *Metadata) SetTitle(val OptNilString) {
	s.Title = val
}

// SetURL sets the value of URL.
func (s *Metadata) SetURL(val OptNilString) {
	s.URL = val
}

type MovePostToThreadReq struct {
	ThreadIconFilename OptNilString `json:"thread_icon_filename"`
	Title              OptNilString `json:"title"`
}

// GetThreadIconFilename returns the value of ThreadIconFilename.
func (s *MovePostToThreadReq) GetThreadIconFilename() OptNilString {
	return s.ThreadIconFilename
}

// GetTitle returns the value of Title.
func (s *MovePostToThreadReq) GetTitle() OptNilString {
	return s.Title
}

// SetThreadIconFilename sets the value of ThreadIconFilename.
func (s *MovePostToThreadReq) SetThreadIconFilename(val OptNilString) {
	s.ThreadIconFilename = val
}

// SetTitle sets the value of Title.
func (s *MovePostToThreadReq) SetTitle(val OptNilString) {
	s.Title = val
}

// MuteGroupUserNoContent is response for MuteGroupUser operation.
type MuteGroupUserNoContent struct{}

// Ref: #/components/schemas/MuteKeyword
type MuteKeyword struct {
	Context OptNilStringArray `json:"context"`
	ID      OptNilInt64       `json:"id"`
	Word    OptNilString      `json:"word"`
}

// GetContext returns the value of Context.
func (s *MuteKeyword) GetContext() OptNilStringArray {
	return s.Context
}

// GetID returns the value of ID.
func (s *MuteKeyword) GetID() OptNilInt64 {
	return s.ID
}

// GetWord returns the value of Word.
func (s *MuteKeyword) GetWord() OptNilString {
	return s.Word
}

// SetContext sets the value of Context.
func (s *MuteKeyword) SetContext(val OptNilStringArray) {
	s.Context = val
}

// SetID sets the value of ID.
func (s *MuteKeyword) SetID(val OptNilInt64) {
	s.ID = val
}

// SetWord sets the value of Word.
func (s *MuteKeyword) SetWord(val OptNilString) {
	s.Word = val
}

// Ref: #/components/schemas/MuteKeywordRequest
type MuteKeywordRequest struct {
	Context OptNilStringArray `json:"context"`
	Word    OptNilString      `json:"word"`
}

// GetContext returns the value of Context.
func (s *MuteKeywordRequest) GetContext() OptNilStringArray {
	return s.Context
}

// GetWord returns the value of Word.
func (s *MuteKeywordRequest) GetWord() OptNilString {
	return s.Word
}

// SetContext sets the value of Context.
func (s *MuteKeywordRequest) SetContext(val OptNilStringArray) {
	s.Context = val
}

// SetWord sets the value of Word.
func (s *MuteKeywordRequest) SetWord(val OptNilString) {
	s.Word = val
}

// Ref: #/components/schemas/MuteKeywordResponse
type MuteKeywordResponse struct {
	HiddenWords OptNilMuteKeywordArray `json:"hidden_words"`
}

// GetHiddenWords returns the value of HiddenWords.
func (s *MuteKeywordResponse) GetHiddenWords() OptNilMuteKeywordArray {
	return s.HiddenWords
}

// SetHiddenWords sets the value of HiddenWords.
func (s *MuteKeywordResponse) SetHiddenWords(val OptNilMuteKeywordArray) {
	s.HiddenWords = val
}

// Ref: #/components/schemas/NoreplyMode
type NoreplyMode string

const (
	NoreplyModeEmpty   NoreplyMode = ""
	NoreplyModeNoreply NoreplyMode = "noreply_"
)

// AllValues returns all NoreplyMode values.
func (NoreplyMode) AllValues() []NoreplyMode {
	return []NoreplyMode{
		NoreplyModeEmpty,
		NoreplyModeNoreply,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s NoreplyMode) MarshalText() ([]byte, error) {
	switch s {
	case NoreplyModeEmpty:
		return []byte(s), nil
	case NoreplyModeNoreply:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *NoreplyMode) UnmarshalText(data []byte) error {
	switch NoreplyMode(data) {
	case NoreplyModeEmpty:
		*s = NoreplyModeEmpty
		return nil
	case NoreplyModeNoreply:
		*s = NoreplyModeNoreply
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/NotificationSettingResponse
type NotificationSettingResponse struct {
	Setting OptUserSetting `json:"setting"`
}

// GetSetting returns the value of Setting.
func (s *NotificationSettingResponse) GetSetting() OptUserSetting {
	return s.Setting
}

// SetSetting sets the value of Setting.
func (s *NotificationSettingResponse) SetSetting(val OptUserSetting) {
	s.Setting = val
}

type OauthTokenReq struct {
	Email        OptNilString `json:"email"`
	GrantType    OptString    `json:"grant_type"`
	Password     OptNilString `json:"password"`
	RefreshToken OptNilString `json:"refresh_token"`
}

// GetEmail returns the value of Email.
func (s *OauthTokenReq) GetEmail() OptNilString {
	return s.Email
}

// GetGrantType returns the value of GrantType.
func (s *OauthTokenReq) GetGrantType() OptString {
	return s.GrantType
}

// GetPassword returns the value of Password.
func (s *OauthTokenReq) GetPassword() OptNilString {
	return s.Password
}

// GetRefreshToken returns the value of RefreshToken.
func (s *OauthTokenReq) GetRefreshToken() OptNilString {
	return s.RefreshToken
}

// SetEmail sets the value of Email.
func (s *OauthTokenReq) SetEmail(val OptNilString) {
	s.Email = val
}

// SetGrantType sets the value of GrantType.
func (s *OauthTokenReq) SetGrantType(val OptString) {
	s.GrantType = val
}

// SetPassword sets the value of Password.
func (s *OauthTokenReq) SetPassword(val OptNilString) {
	s.Password = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *OauthTokenReq) SetRefreshToken(val OptNilString) {
	s.RefreshToken = val
}

// Ref: #/components/schemas/OnlineStatus
type OnlineStatus struct {
	APIValue OptNilString `json:"api_value"`
}

// GetAPIValue returns the value of APIValue.
func (s *OnlineStatus) GetAPIValue() OptNilString {
	return s.APIValue
}

// SetAPIValue sets the value of APIValue.
func (s *OnlineStatus) SetAPIValue(val OptNilString) {
	s.APIValue = val
}

// Ref: #/components/schemas/OnlineStatusEnum
type OnlineStatusEnum string

const (
	OnlineStatusEnumOffline OnlineStatusEnum = "offline"
	OnlineStatusEnumHidden  OnlineStatusEnum = "hidden"
)

// AllValues returns all OnlineStatusEnum values.
func (OnlineStatusEnum) AllValues() []OnlineStatusEnum {
	return []OnlineStatusEnum{
		OnlineStatusEnumOffline,
		OnlineStatusEnumHidden,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OnlineStatusEnum) MarshalText() ([]byte, error) {
	switch s {
	case OnlineStatusEnumOffline:
		return []byte(s), nil
	case OnlineStatusEnumHidden:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OnlineStatusEnum) UnmarshalText(data []byte) error {
	switch OnlineStatusEnum(data) {
	case OnlineStatusEnumOffline:
		*s = OnlineStatusEnumOffline
		return nil
	case OnlineStatusEnumHidden:
		*s = OnlineStatusEnumHidden
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptApplication returns new OptApplication with value set to v.
func NewOptApplication(v Application) OptApplication {
	return OptApplication{
		Value: v,
		Set:   true,
	}
}

// OptApplication is optional Application.
type OptApplication struct {
	Value Application
	Set   bool
}

// IsSet returns true if OptApplication was set.
func (o OptApplication) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptApplication) Reset() {
	var v Application
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptApplication) SetTo(v Application) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptApplication) Get() (v Application, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptApplication) Or(d Application) Application {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChatInvitation returns new OptChatInvitation with value set to v.
func NewOptChatInvitation(v ChatInvitation) OptChatInvitation {
	return OptChatInvitation{
		Value: v,
		Set:   true,
	}
}

// OptChatInvitation is optional ChatInvitation.
type OptChatInvitation struct {
	Value ChatInvitation
	Set   bool
}

// IsSet returns true if OptChatInvitation was set.
func (o OptChatInvitation) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChatInvitation) Reset() {
	var v ChatInvitation
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChatInvitation) SetTo(v ChatInvitation) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChatInvitation) Get() (v ChatInvitation, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChatInvitation) Or(d ChatInvitation) ChatInvitation {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChatRoomLastMessage returns new OptChatRoomLastMessage with value set to v.
func NewOptChatRoomLastMessage(v ChatRoomLastMessage) OptChatRoomLastMessage {
	return OptChatRoomLastMessage{
		Value: v,
		Set:   true,
	}
}

// OptChatRoomLastMessage is optional ChatRoomLastMessage.
type OptChatRoomLastMessage struct {
	Value ChatRoomLastMessage
	Set   bool
}

// IsSet returns true if OptChatRoomLastMessage was set.
func (o OptChatRoomLastMessage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChatRoomLastMessage) Reset() {
	var v ChatRoomLastMessage
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChatRoomLastMessage) SetTo(v ChatRoomLastMessage) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChatRoomLastMessage) Get() (v ChatRoomLastMessage, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChatRoomLastMessage) Or(d ChatRoomLastMessage) ChatRoomLastMessage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptConferenceCallBumpParams returns new OptConferenceCallBumpParams with value set to v.
func NewOptConferenceCallBumpParams(v ConferenceCallBumpParams) OptConferenceCallBumpParams {
	return OptConferenceCallBumpParams{
		Value: v,
		Set:   true,
	}
}

// OptConferenceCallBumpParams is optional ConferenceCallBumpParams.
type OptConferenceCallBumpParams struct {
	Value ConferenceCallBumpParams
	Set   bool
}

// IsSet returns true if OptConferenceCallBumpParams was set.
func (o OptConferenceCallBumpParams) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptConferenceCallBumpParams) Reset() {
	var v ConferenceCallBumpParams
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptConferenceCallBumpParams) SetTo(v ConferenceCallBumpParams) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptConferenceCallBumpParams) Get() (v ConferenceCallBumpParams, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptConferenceCallBumpParams) Or(d ConferenceCallBumpParams) ConferenceCallBumpParams {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptConfig returns new OptConfig with value set to v.
func NewOptConfig(v Config) OptConfig {
	return OptConfig{
		Value: v,
		Set:   true,
	}
}

// OptConfig is optional Config.
type OptConfig struct {
	Value Config
	Set   bool
}

// IsSet returns true if OptConfig was set.
func (o OptConfig) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptConfig) Reset() {
	var v Config
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptConfig) SetTo(v Config) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptConfig) Get() (v Config, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptConfig) Or(d Config) Config {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateConferenceCallPostReqMessageTags returns new OptCreateConferenceCallPostReqMessageTags with value set to v.
func NewOptCreateConferenceCallPostReqMessageTags(v *CreateConferenceCallPostReqMessageTags) OptCreateConferenceCallPostReqMessageTags {
	return OptCreateConferenceCallPostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptCreateConferenceCallPostReqMessageTags is optional *CreateConferenceCallPostReqMessageTags.
type OptCreateConferenceCallPostReqMessageTags struct {
	Value *CreateConferenceCallPostReqMessageTags
	Set   bool
}

// IsSet returns true if OptCreateConferenceCallPostReqMessageTags was set.
func (o OptCreateConferenceCallPostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateConferenceCallPostReqMessageTags) Reset() {
	var v *CreateConferenceCallPostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateConferenceCallPostReqMessageTags) SetTo(v *CreateConferenceCallPostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateConferenceCallPostReqMessageTags) Get() (v *CreateConferenceCallPostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateConferenceCallPostReqMessageTags) Or(d *CreateConferenceCallPostReqMessageTags) *CreateConferenceCallPostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateGroupQuota returns new OptCreateGroupQuota with value set to v.
func NewOptCreateGroupQuota(v CreateGroupQuota) OptCreateGroupQuota {
	return OptCreateGroupQuota{
		Value: v,
		Set:   true,
	}
}

// OptCreateGroupQuota is optional CreateGroupQuota.
type OptCreateGroupQuota struct {
	Value CreateGroupQuota
	Set   bool
}

// IsSet returns true if OptCreateGroupQuota was set.
func (o OptCreateGroupQuota) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateGroupQuota) Reset() {
	var v CreateGroupQuota
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateGroupQuota) SetTo(v CreateGroupQuota) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateGroupQuota) Get() (v CreateGroupQuota, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateGroupQuota) Or(d CreateGroupQuota) CreateGroupQuota {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreatePostReqMessageTags returns new OptCreatePostReqMessageTags with value set to v.
func NewOptCreatePostReqMessageTags(v *CreatePostReqMessageTags) OptCreatePostReqMessageTags {
	return OptCreatePostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptCreatePostReqMessageTags is optional *CreatePostReqMessageTags.
type OptCreatePostReqMessageTags struct {
	Value *CreatePostReqMessageTags
	Set   bool
}

// IsSet returns true if OptCreatePostReqMessageTags was set.
func (o OptCreatePostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreatePostReqMessageTags) Reset() {
	var v *CreatePostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreatePostReqMessageTags) SetTo(v *CreatePostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreatePostReqMessageTags) Get() (v *CreatePostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreatePostReqMessageTags) Or(d *CreatePostReqMessageTags) *CreatePostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreatePostReqSharedURL returns new OptCreatePostReqSharedURL with value set to v.
func NewOptCreatePostReqSharedURL(v *CreatePostReqSharedURL) OptCreatePostReqSharedURL {
	return OptCreatePostReqSharedURL{
		Value: v,
		Set:   true,
	}
}

// OptCreatePostReqSharedURL is optional *CreatePostReqSharedURL.
type OptCreatePostReqSharedURL struct {
	Value *CreatePostReqSharedURL
	Set   bool
}

// IsSet returns true if OptCreatePostReqSharedURL was set.
func (o OptCreatePostReqSharedURL) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreatePostReqSharedURL) Reset() {
	var v *CreatePostReqSharedURL
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreatePostReqSharedURL) SetTo(v *CreatePostReqSharedURL) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreatePostReqSharedURL) Get() (v *CreatePostReqSharedURL, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreatePostReqSharedURL) Or(d *CreatePostReqSharedURL) *CreatePostReqSharedURL {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateSharePostReqMessageTags returns new OptCreateSharePostReqMessageTags with value set to v.
func NewOptCreateSharePostReqMessageTags(v *CreateSharePostReqMessageTags) OptCreateSharePostReqMessageTags {
	return OptCreateSharePostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptCreateSharePostReqMessageTags is optional *CreateSharePostReqMessageTags.
type OptCreateSharePostReqMessageTags struct {
	Value *CreateSharePostReqMessageTags
	Set   bool
}

// IsSet returns true if OptCreateSharePostReqMessageTags was set.
func (o OptCreateSharePostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateSharePostReqMessageTags) Reset() {
	var v *CreateSharePostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateSharePostReqMessageTags) SetTo(v *CreateSharePostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateSharePostReqMessageTags) Get() (v *CreateSharePostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateSharePostReqMessageTags) Or(d *CreateSharePostReqMessageTags) *CreateSharePostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateThreadPostReqMessageTags returns new OptCreateThreadPostReqMessageTags with value set to v.
func NewOptCreateThreadPostReqMessageTags(v *CreateThreadPostReqMessageTags) OptCreateThreadPostReqMessageTags {
	return OptCreateThreadPostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptCreateThreadPostReqMessageTags is optional *CreateThreadPostReqMessageTags.
type OptCreateThreadPostReqMessageTags struct {
	Value *CreateThreadPostReqMessageTags
	Set   bool
}

// IsSet returns true if OptCreateThreadPostReqMessageTags was set.
func (o OptCreateThreadPostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateThreadPostReqMessageTags) Reset() {
	var v *CreateThreadPostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateThreadPostReqMessageTags) SetTo(v *CreateThreadPostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateThreadPostReqMessageTags) Get() (v *CreateThreadPostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateThreadPostReqMessageTags) Or(d *CreateThreadPostReqMessageTags) *CreateThreadPostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateThreadPostReqSharedURL returns new OptCreateThreadPostReqSharedURL with value set to v.
func NewOptCreateThreadPostReqSharedURL(v *CreateThreadPostReqSharedURL) OptCreateThreadPostReqSharedURL {
	return OptCreateThreadPostReqSharedURL{
		Value: v,
		Set:   true,
	}
}

// OptCreateThreadPostReqSharedURL is optional *CreateThreadPostReqSharedURL.
type OptCreateThreadPostReqSharedURL struct {
	Value *CreateThreadPostReqSharedURL
	Set   bool
}

// IsSet returns true if OptCreateThreadPostReqSharedURL was set.
func (o OptCreateThreadPostReqSharedURL) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateThreadPostReqSharedURL) Reset() {
	var v *CreateThreadPostReqSharedURL
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateThreadPostReqSharedURL) SetTo(v *CreateThreadPostReqSharedURL) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateThreadPostReqSharedURL) Get() (v *CreateThreadPostReqSharedURL, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateThreadPostReqSharedURL) Or(d *CreateThreadPostReqSharedURL) *CreateThreadPostReqSharedURL {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGender returns new OptGender with value set to v.
func NewOptGender(v Gender) OptGender {
	return OptGender{
		Value: v,
		Set:   true,
	}
}

// OptGender is optional Gender.
type OptGender struct {
	Value Gender
	Set   bool
}

// IsSet returns true if OptGender was set.
func (o OptGender) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGender) Reset() {
	var v Gender
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGender) SetTo(v Gender) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGender) Get() (v Gender, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGender) Or(d Gender) Gender {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGifImage returns new OptGifImage with value set to v.
func NewOptGifImage(v GifImage) OptGifImage {
	return OptGifImage{
		Value: v,
		Set:   true,
	}
}

// OptGifImage is optional GifImage.
type OptGifImage struct {
	Value GifImage
	Set   bool
}

// IsSet returns true if OptGifImage was set.
func (o OptGifImage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGifImage) Reset() {
	var v GifImage
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGifImage) SetTo(v GifImage) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGifImage) Get() (v GifImage, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGifImage) Or(d GifImage) GifImage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGift returns new OptGift with value set to v.
func NewOptGift(v Gift) OptGift {
	return OptGift{
		Value: v,
		Set:   true,
	}
}

// OptGift is optional Gift.
type OptGift struct {
	Value Gift
	Set   bool
}

// IsSet returns true if OptGift was set.
func (o OptGift) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGift) Reset() {
	var v Gift
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGift) SetTo(v Gift) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGift) Get() (v Gift, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGift) Or(d Gift) Gift {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGiftSlugItem returns new OptGiftSlugItem with value set to v.
func NewOptGiftSlugItem(v GiftSlugItem) OptGiftSlugItem {
	return OptGiftSlugItem{
		Value: v,
		Set:   true,
	}
}

// OptGiftSlugItem is optional GiftSlugItem.
type OptGiftSlugItem struct {
	Value GiftSlugItem
	Set   bool
}

// IsSet returns true if OptGiftSlugItem was set.
func (o OptGiftSlugItem) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGiftSlugItem) Reset() {
	var v GiftSlugItem
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGiftSlugItem) SetTo(v GiftSlugItem) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGiftSlugItem) Get() (v GiftSlugItem, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGiftSlugItem) Or(d GiftSlugItem) GiftSlugItem {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGiftingAbility returns new OptGiftingAbility with value set to v.
func NewOptGiftingAbility(v GiftingAbility) OptGiftingAbility {
	return OptGiftingAbility{
		Value: v,
		Set:   true,
	}
}

// OptGiftingAbility is optional GiftingAbility.
type OptGiftingAbility struct {
	Value GiftingAbility
	Set   bool
}

// IsSet returns true if OptGiftingAbility was set.
func (o OptGiftingAbility) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGiftingAbility) Reset() {
	var v GiftingAbility
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGiftingAbility) SetTo(v GiftingAbility) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGiftingAbility) Get() (v GiftingAbility, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGiftingAbility) Or(d GiftingAbility) GiftingAbility {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGroup returns new OptGroup with value set to v.
func NewOptGroup(v Group) OptGroup {
	return OptGroup{
		Value: v,
		Set:   true,
	}
}

// OptGroup is optional Group.
type OptGroup struct {
	Value Group
	Set   bool
}

// IsSet returns true if OptGroup was set.
func (o OptGroup) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGroup) Reset() {
	var v Group
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGroup) SetTo(v Group) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGroup) Get() (v Group, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGroup) Or(d Group) Group {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGroupUser returns new OptGroupUser with value set to v.
func NewOptGroupUser(v GroupUser) OptGroupUser {
	return OptGroupUser{
		Value: v,
		Set:   true,
	}
}

// OptGroupUser is optional GroupUser.
type OptGroupUser struct {
	Value GroupUser
	Set   bool
}

// IsSet returns true if OptGroupUser was set.
func (o OptGroupUser) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGroupUser) Reset() {
	var v GroupUser
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGroupUser) SetTo(v GroupUser) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGroupUser) Get() (v GroupUser, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGroupUser) Or(d GroupUser) GroupUser {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMessageType returns new OptMessageType with value set to v.
func NewOptMessageType(v MessageType) OptMessageType {
	return OptMessageType{
		Value: v,
		Set:   true,
	}
}

// OptMessageType is optional MessageType.
type OptMessageType struct {
	Value MessageType
	Set   bool
}

// IsSet returns true if OptMessageType was set.
func (o OptMessageType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMessageType) Reset() {
	var v MessageType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMessageType) SetTo(v MessageType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMessageType) Get() (v MessageType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMessageType) Or(d MessageType) MessageType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMetadata returns new OptMetadata with value set to v.
func NewOptMetadata(v Metadata) OptMetadata {
	return OptMetadata{
		Value: v,
		Set:   true,
	}
}

// OptMetadata is optional Metadata.
type OptMetadata struct {
	Value Metadata
	Set   bool
}

// IsSet returns true if OptMetadata was set.
func (o OptMetadata) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMetadata) Reset() {
	var v Metadata
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMetadata) SetTo(v Metadata) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMetadata) Get() (v Metadata, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMetadata) Or(d Metadata) Metadata {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMuteKeyword returns new OptMuteKeyword with value set to v.
func NewOptMuteKeyword(v MuteKeyword) OptMuteKeyword {
	return OptMuteKeyword{
		Value: v,
		Set:   true,
	}
}

// OptMuteKeyword is optional MuteKeyword.
type OptMuteKeyword struct {
	Value MuteKeyword
	Set   bool
}

// IsSet returns true if OptMuteKeyword was set.
func (o OptMuteKeyword) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMuteKeyword) Reset() {
	var v MuteKeyword
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMuteKeyword) SetTo(v MuteKeyword) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMuteKeyword) Get() (v MuteKeyword, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMuteKeyword) Or(d MuteKeyword) MuteKeyword {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilActivityArray returns new OptNilActivityArray with value set to v.
func NewOptNilActivityArray(v []Activity) OptNilActivityArray {
	return OptNilActivityArray{
		Value: v,
		Set:   true,
	}
}

// OptNilActivityArray is optional nullable []Activity.
type OptNilActivityArray struct {
	Value []Activity
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilActivityArray was set.
func (o OptNilActivityArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilActivityArray) Reset() {
	var v []Activity
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilActivityArray) SetTo(v []Activity) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilActivityArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilActivityArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Activity
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilActivityArray) Get() (v []Activity, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilActivityArray) Or(d []Activity) []Activity {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilBanWordArray returns new OptNilBanWordArray with value set to v.
func NewOptNilBanWordArray(v []BanWord) OptNilBanWordArray {
	return OptNilBanWordArray{
		Value: v,
		Set:   true,
	}
}

// OptNilBanWordArray is optional nullable []BanWord.
type OptNilBanWordArray struct {
	Value []BanWord
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilBanWordArray was set.
func (o OptNilBanWordArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilBanWordArray) Reset() {
	var v []BanWord
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilBanWordArray) SetTo(v []BanWord) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilBanWordArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilBanWordArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []BanWord
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilBanWordArray) Get() (v []BanWord, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilBanWordArray) Or(d []BanWord) []BanWord {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilBgmArray returns new OptNilBgmArray with value set to v.
func NewOptNilBgmArray(v []Bgm) OptNilBgmArray {
	return OptNilBgmArray{
		Value: v,
		Set:   true,
	}
}

// OptNilBgmArray is optional nullable []Bgm.
type OptNilBgmArray struct {
	Value []Bgm
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilBgmArray was set.
func (o OptNilBgmArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilBgmArray) Reset() {
	var v []Bgm
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilBgmArray) SetTo(v []Bgm) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilBgmArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilBgmArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Bgm
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilBgmArray) Get() (v []Bgm, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilBgmArray) Or(d []Bgm) []Bgm {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilBool returns new OptNilBool with value set to v.
func NewOptNilBool(v bool) OptNilBool {
	return OptNilBool{
		Value: v,
		Set:   true,
	}
}

// OptNilBool is optional nullable bool.
type OptNilBool struct {
	Value bool
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilBool was set.
func (o OptNilBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilBool) SetTo(v bool) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilBool) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilBool) SetToNull() {
	o.Set = true
	o.Null = true
	var v bool
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilBool) Get() (v bool, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilChoiceArray returns new OptNilChoiceArray with value set to v.
func NewOptNilChoiceArray(v []Choice) OptNilChoiceArray {
	return OptNilChoiceArray{
		Value: v,
		Set:   true,
	}
}

// OptNilChoiceArray is optional nullable []Choice.
type OptNilChoiceArray struct {
	Value []Choice
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilChoiceArray was set.
func (o OptNilChoiceArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilChoiceArray) Reset() {
	var v []Choice
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilChoiceArray) SetTo(v []Choice) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilChoiceArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilChoiceArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Choice
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilChoiceArray) Get() (v []Choice, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilChoiceArray) Or(d []Choice) []Choice {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilConferenceCallUserRoleArray returns new OptNilConferenceCallUserRoleArray with value set to v.
func NewOptNilConferenceCallUserRoleArray(v []ConferenceCallUserRole) OptNilConferenceCallUserRoleArray {
	return OptNilConferenceCallUserRoleArray{
		Value: v,
		Set:   true,
	}
}

// OptNilConferenceCallUserRoleArray is optional nullable []ConferenceCallUserRole.
type OptNilConferenceCallUserRoleArray struct {
	Value []ConferenceCallUserRole
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilConferenceCallUserRoleArray was set.
func (o OptNilConferenceCallUserRoleArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilConferenceCallUserRoleArray) Reset() {
	var v []ConferenceCallUserRole
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilConferenceCallUserRoleArray) SetTo(v []ConferenceCallUserRole) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilConferenceCallUserRoleArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilConferenceCallUserRoleArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []ConferenceCallUserRole
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilConferenceCallUserRoleArray) Get() (v []ConferenceCallUserRole, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilConferenceCallUserRoleArray) Or(d []ConferenceCallUserRole) []ConferenceCallUserRole {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilFloat64 returns new OptNilFloat64 with value set to v.
func NewOptNilFloat64(v float64) OptNilFloat64 {
	return OptNilFloat64{
		Value: v,
		Set:   true,
	}
}

// OptNilFloat64 is optional nullable float64.
type OptNilFloat64 struct {
	Value float64
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilFloat64 was set.
func (o OptNilFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilFloat64) SetTo(v float64) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilFloat64) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilFloat64) SetToNull() {
	o.Set = true
	o.Null = true
	var v float64
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilFloat64) Get() (v float64, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilFootprintDTOArray returns new OptNilFootprintDTOArray with value set to v.
func NewOptNilFootprintDTOArray(v []FootprintDTO) OptNilFootprintDTOArray {
	return OptNilFootprintDTOArray{
		Value: v,
		Set:   true,
	}
}

// OptNilFootprintDTOArray is optional nullable []FootprintDTO.
type OptNilFootprintDTOArray struct {
	Value []FootprintDTO
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilFootprintDTOArray was set.
func (o OptNilFootprintDTOArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilFootprintDTOArray) Reset() {
	var v []FootprintDTO
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilFootprintDTOArray) SetTo(v []FootprintDTO) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilFootprintDTOArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilFootprintDTOArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []FootprintDTO
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilFootprintDTOArray) Get() (v []FootprintDTO, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilFootprintDTOArray) Or(d []FootprintDTO) []FootprintDTO {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGiftCountArray returns new OptNilGiftCountArray with value set to v.
func NewOptNilGiftCountArray(v []GiftCount) OptNilGiftCountArray {
	return OptNilGiftCountArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGiftCountArray is optional nullable []GiftCount.
type OptNilGiftCountArray struct {
	Value []GiftCount
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGiftCountArray was set.
func (o OptNilGiftCountArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGiftCountArray) Reset() {
	var v []GiftCount
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGiftCountArray) SetTo(v []GiftCount) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGiftCountArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGiftCountArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []GiftCount
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGiftCountArray) Get() (v []GiftCount, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGiftCountArray) Or(d []GiftCount) []GiftCount {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGiftHistoryArray returns new OptNilGiftHistoryArray with value set to v.
func NewOptNilGiftHistoryArray(v []GiftHistory) OptNilGiftHistoryArray {
	return OptNilGiftHistoryArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGiftHistoryArray is optional nullable []GiftHistory.
type OptNilGiftHistoryArray struct {
	Value []GiftHistory
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGiftHistoryArray was set.
func (o OptNilGiftHistoryArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGiftHistoryArray) Reset() {
	var v []GiftHistory
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGiftHistoryArray) SetTo(v []GiftHistory) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGiftHistoryArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGiftHistoryArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []GiftHistory
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGiftHistoryArray) Get() (v []GiftHistory, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGiftHistoryArray) Or(d []GiftHistory) []GiftHistory {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGroupArray returns new OptNilGroupArray with value set to v.
func NewOptNilGroupArray(v []Group) OptNilGroupArray {
	return OptNilGroupArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGroupArray is optional nullable []Group.
type OptNilGroupArray struct {
	Value []Group
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGroupArray was set.
func (o OptNilGroupArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGroupArray) Reset() {
	var v []Group
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGroupArray) SetTo(v []Group) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGroupArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGroupArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Group
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGroupArray) Get() (v []Group, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGroupArray) Or(d []Group) []Group {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGroupCategoryArray returns new OptNilGroupCategoryArray with value set to v.
func NewOptNilGroupCategoryArray(v []GroupCategory) OptNilGroupCategoryArray {
	return OptNilGroupCategoryArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGroupCategoryArray is optional nullable []GroupCategory.
type OptNilGroupCategoryArray struct {
	Value []GroupCategory
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGroupCategoryArray was set.
func (o OptNilGroupCategoryArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGroupCategoryArray) Reset() {
	var v []GroupCategory
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGroupCategoryArray) SetTo(v []GroupCategory) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGroupCategoryArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGroupCategoryArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []GroupCategory
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGroupCategoryArray) Get() (v []GroupCategory, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGroupCategoryArray) Or(d []GroupCategory) []GroupCategory {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGroupGiftHistoryArray returns new OptNilGroupGiftHistoryArray with value set to v.
func NewOptNilGroupGiftHistoryArray(v []GroupGiftHistory) OptNilGroupGiftHistoryArray {
	return OptNilGroupGiftHistoryArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGroupGiftHistoryArray is optional nullable []GroupGiftHistory.
type OptNilGroupGiftHistoryArray struct {
	Value []GroupGiftHistory
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGroupGiftHistoryArray was set.
func (o OptNilGroupGiftHistoryArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGroupGiftHistoryArray) Reset() {
	var v []GroupGiftHistory
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGroupGiftHistoryArray) SetTo(v []GroupGiftHistory) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGroupGiftHistoryArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGroupGiftHistoryArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []GroupGiftHistory
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGroupGiftHistoryArray) Get() (v []GroupGiftHistory, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGroupGiftHistoryArray) Or(d []GroupGiftHistory) []GroupGiftHistory {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGroupUserArray returns new OptNilGroupUserArray with value set to v.
func NewOptNilGroupUserArray(v []GroupUser) OptNilGroupUserArray {
	return OptNilGroupUserArray{
		Value: v,
		Set:   true,
	}
}

// OptNilGroupUserArray is optional nullable []GroupUser.
type OptNilGroupUserArray struct {
	Value []GroupUser
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGroupUserArray was set.
func (o OptNilGroupUserArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGroupUserArray) Reset() {
	var v []GroupUser
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGroupUserArray) SetTo(v []GroupUser) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilGroupUserArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilGroupUserArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []GroupUser
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGroupUserArray) Get() (v []GroupUser, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGroupUserArray) Or(d []GroupUser) []GroupUser {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilInt32 returns new OptNilInt32 with value set to v.
func NewOptNilInt32(v int32) OptNilInt32 {
	return OptNilInt32{
		Value: v,
		Set:   true,
	}
}

// OptNilInt32 is optional nullable int32.
type OptNilInt32 struct {
	Value int32
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilInt32 was set.
func (o OptNilInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilInt32) SetTo(v int32) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilInt32) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilInt32) SetToNull() {
	o.Set = true
	o.Null = true
	var v int32
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilInt32) Get() (v int32, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilInt64 returns new OptNilInt64 with value set to v.
func NewOptNilInt64(v int64) OptNilInt64 {
	return OptNilInt64{
		Value: v,
		Set:   true,
	}
}

// OptNilInt64 is optional nullable int64.
type OptNilInt64 struct {
	Value int64
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilInt64 was set.
func (o OptNilInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilInt64) SetTo(v int64) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilInt64) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilInt64) SetToNull() {
	o.Set = true
	o.Null = true
	var v int64
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilInt64) Get() (v int64, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilInt64Array returns new OptNilInt64Array with value set to v.
func NewOptNilInt64Array(v []int64) OptNilInt64Array {
	return OptNilInt64Array{
		Value: v,
		Set:   true,
	}
}

// OptNilInt64Array is optional nullable []int64.
type OptNilInt64Array struct {
	Value []int64
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilInt64Array was set.
func (o OptNilInt64Array) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilInt64Array) Reset() {
	var v []int64
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilInt64Array) SetTo(v []int64) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilInt64Array) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilInt64Array) SetToNull() {
	o.Set = true
	o.Null = true
	var v []int64
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilInt64Array) Get() (v []int64, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilInt64Array) Or(d []int64) []int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilInterestArray returns new OptNilInterestArray with value set to v.
func NewOptNilInterestArray(v []Interest) OptNilInterestArray {
	return OptNilInterestArray{
		Value: v,
		Set:   true,
	}
}

// OptNilInterestArray is optional nullable []Interest.
type OptNilInterestArray struct {
	Value []Interest
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilInterestArray was set.
func (o OptNilInterestArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilInterestArray) Reset() {
	var v []Interest
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilInterestArray) SetTo(v []Interest) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilInterestArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilInterestArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Interest
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilInterestArray) Get() (v []Interest, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilInterestArray) Or(d []Interest) []Interest {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilMessageTagArray returns new OptNilMessageTagArray with value set to v.
func NewOptNilMessageTagArray(v []MessageTag) OptNilMessageTagArray {
	return OptNilMessageTagArray{
		Value: v,
		Set:   true,
	}
}

// OptNilMessageTagArray is optional nullable []MessageTag.
type OptNilMessageTagArray struct {
	Value []MessageTag
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilMessageTagArray was set.
func (o OptNilMessageTagArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilMessageTagArray) Reset() {
	var v []MessageTag
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilMessageTagArray) SetTo(v []MessageTag) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilMessageTagArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilMessageTagArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []MessageTag
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilMessageTagArray) Get() (v []MessageTag, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilMessageTagArray) Or(d []MessageTag) []MessageTag {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilMuteKeywordArray returns new OptNilMuteKeywordArray with value set to v.
func NewOptNilMuteKeywordArray(v []MuteKeyword) OptNilMuteKeywordArray {
	return OptNilMuteKeywordArray{
		Value: v,
		Set:   true,
	}
}

// OptNilMuteKeywordArray is optional nullable []MuteKeyword.
type OptNilMuteKeywordArray struct {
	Value []MuteKeyword
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilMuteKeywordArray was set.
func (o OptNilMuteKeywordArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilMuteKeywordArray) Reset() {
	var v []MuteKeyword
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilMuteKeywordArray) SetTo(v []MuteKeyword) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilMuteKeywordArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilMuteKeywordArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []MuteKeyword
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilMuteKeywordArray) Get() (v []MuteKeyword, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilMuteKeywordArray) Or(d []MuteKeyword) []MuteKeyword {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilPopularWordArray returns new OptNilPopularWordArray with value set to v.
func NewOptNilPopularWordArray(v []PopularWord) OptNilPopularWordArray {
	return OptNilPopularWordArray{
		Value: v,
		Set:   true,
	}
}

// OptNilPopularWordArray is optional nullable []PopularWord.
type OptNilPopularWordArray struct {
	Value []PopularWord
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilPopularWordArray was set.
func (o OptNilPopularWordArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilPopularWordArray) Reset() {
	var v []PopularWord
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilPopularWordArray) SetTo(v []PopularWord) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilPopularWordArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilPopularWordArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []PopularWord
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilPopularWordArray) Get() (v []PopularWord, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilPopularWordArray) Or(d []PopularWord) []PopularWord {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilPostArray returns new OptNilPostArray with value set to v.
func NewOptNilPostArray(v []Post) OptNilPostArray {
	return OptNilPostArray{
		Value: v,
		Set:   true,
	}
}

// OptNilPostArray is optional nullable []Post.
type OptNilPostArray struct {
	Value []Post
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilPostArray was set.
func (o OptNilPostArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilPostArray) Reset() {
	var v []Post
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilPostArray) SetTo(v []Post) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilPostArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilPostArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Post
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilPostArray) Get() (v []Post, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilPostArray) Or(d []Post) []Post {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilPostTagArray returns new OptNilPostTagArray with value set to v.
func NewOptNilPostTagArray(v []PostTag) OptNilPostTagArray {
	return OptNilPostTagArray{
		Value: v,
		Set:   true,
	}
}

// OptNilPostTagArray is optional nullable []PostTag.
type OptNilPostTagArray struct {
	Value []PostTag
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilPostTagArray was set.
func (o OptNilPostTagArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilPostTagArray) Reset() {
	var v []PostTag
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilPostTagArray) SetTo(v []PostTag) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilPostTagArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilPostTagArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []PostTag
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilPostTagArray) Get() (v []PostTag, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilPostTagArray) Or(d []PostTag) []PostTag {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilPresignedUrlArray returns new OptNilPresignedUrlArray with value set to v.
func NewOptNilPresignedUrlArray(v []PresignedUrl) OptNilPresignedUrlArray {
	return OptNilPresignedUrlArray{
		Value: v,
		Set:   true,
	}
}

// OptNilPresignedUrlArray is optional nullable []PresignedUrl.
type OptNilPresignedUrlArray struct {
	Value []PresignedUrl
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilPresignedUrlArray was set.
func (o OptNilPresignedUrlArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilPresignedUrlArray) Reset() {
	var v []PresignedUrl
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilPresignedUrlArray) SetTo(v []PresignedUrl) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilPresignedUrlArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilPresignedUrlArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []PresignedUrl
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilPresignedUrlArray) Get() (v []PresignedUrl, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilPresignedUrlArray) Or(d []PresignedUrl) []PresignedUrl {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmChatRoomArray returns new OptNilRealmChatRoomArray with value set to v.
func NewOptNilRealmChatRoomArray(v []RealmChatRoom) OptNilRealmChatRoomArray {
	return OptNilRealmChatRoomArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmChatRoomArray is optional nullable []RealmChatRoom.
type OptNilRealmChatRoomArray struct {
	Value []RealmChatRoom
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmChatRoomArray was set.
func (o OptNilRealmChatRoomArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmChatRoomArray) Reset() {
	var v []RealmChatRoom
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmChatRoomArray) SetTo(v []RealmChatRoom) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmChatRoomArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmChatRoomArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmChatRoom
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmChatRoomArray) Get() (v []RealmChatRoom, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmChatRoomArray) Or(d []RealmChatRoom) []RealmChatRoom {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmGameArray returns new OptNilRealmGameArray with value set to v.
func NewOptNilRealmGameArray(v []RealmGame) OptNilRealmGameArray {
	return OptNilRealmGameArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmGameArray is optional nullable []RealmGame.
type OptNilRealmGameArray struct {
	Value []RealmGame
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmGameArray was set.
func (o OptNilRealmGameArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmGameArray) Reset() {
	var v []RealmGame
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmGameArray) SetTo(v []RealmGame) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmGameArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmGameArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmGame
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmGameArray) Get() (v []RealmGame, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmGameArray) Or(d []RealmGame) []RealmGame {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmGenreArray returns new OptNilRealmGenreArray with value set to v.
func NewOptNilRealmGenreArray(v []RealmGenre) OptNilRealmGenreArray {
	return OptNilRealmGenreArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmGenreArray is optional nullable []RealmGenre.
type OptNilRealmGenreArray struct {
	Value []RealmGenre
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmGenreArray was set.
func (o OptNilRealmGenreArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmGenreArray) Reset() {
	var v []RealmGenre
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmGenreArray) SetTo(v []RealmGenre) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmGenreArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmGenreArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmGenre
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmGenreArray) Get() (v []RealmGenre, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmGenreArray) Or(d []RealmGenre) []RealmGenre {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmGiftArray returns new OptNilRealmGiftArray with value set to v.
func NewOptNilRealmGiftArray(v []RealmGift) OptNilRealmGiftArray {
	return OptNilRealmGiftArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmGiftArray is optional nullable []RealmGift.
type OptNilRealmGiftArray struct {
	Value []RealmGift
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmGiftArray was set.
func (o OptNilRealmGiftArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmGiftArray) Reset() {
	var v []RealmGift
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmGiftArray) SetTo(v []RealmGift) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmGiftArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmGiftArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmGift
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmGiftArray) Get() (v []RealmGift, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmGiftArray) Or(d []RealmGift) []RealmGift {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmMessageArray returns new OptNilRealmMessageArray with value set to v.
func NewOptNilRealmMessageArray(v []RealmMessage) OptNilRealmMessageArray {
	return OptNilRealmMessageArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmMessageArray is optional nullable []RealmMessage.
type OptNilRealmMessageArray struct {
	Value []RealmMessage
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmMessageArray was set.
func (o OptNilRealmMessageArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmMessageArray) Reset() {
	var v []RealmMessage
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmMessageArray) SetTo(v []RealmMessage) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmMessageArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmMessageArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmMessage
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmMessageArray) Get() (v []RealmMessage, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmMessageArray) Or(d []RealmMessage) []RealmMessage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRealmUserArray returns new OptNilRealmUserArray with value set to v.
func NewOptNilRealmUserArray(v []RealmUser) OptNilRealmUserArray {
	return OptNilRealmUserArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRealmUserArray is optional nullable []RealmUser.
type OptNilRealmUserArray struct {
	Value []RealmUser
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRealmUserArray was set.
func (o OptNilRealmUserArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRealmUserArray) Reset() {
	var v []RealmUser
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRealmUserArray) SetTo(v []RealmUser) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRealmUserArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRealmUserArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RealmUser
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRealmUserArray) Get() (v []RealmUser, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRealmUserArray) Or(d []RealmUser) []RealmUser {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilReceivedGiftArray returns new OptNilReceivedGiftArray with value set to v.
func NewOptNilReceivedGiftArray(v []ReceivedGift) OptNilReceivedGiftArray {
	return OptNilReceivedGiftArray{
		Value: v,
		Set:   true,
	}
}

// OptNilReceivedGiftArray is optional nullable []ReceivedGift.
type OptNilReceivedGiftArray struct {
	Value []ReceivedGift
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilReceivedGiftArray was set.
func (o OptNilReceivedGiftArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilReceivedGiftArray) Reset() {
	var v []ReceivedGift
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilReceivedGiftArray) SetTo(v []ReceivedGift) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilReceivedGiftArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilReceivedGiftArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []ReceivedGift
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilReceivedGiftArray) Get() (v []ReceivedGift, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilReceivedGiftArray) Or(d []ReceivedGift) []ReceivedGift {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilRefreshCounterRequestArray returns new OptNilRefreshCounterRequestArray with value set to v.
func NewOptNilRefreshCounterRequestArray(v []RefreshCounterRequest) OptNilRefreshCounterRequestArray {
	return OptNilRefreshCounterRequestArray{
		Value: v,
		Set:   true,
	}
}

// OptNilRefreshCounterRequestArray is optional nullable []RefreshCounterRequest.
type OptNilRefreshCounterRequestArray struct {
	Value []RefreshCounterRequest
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilRefreshCounterRequestArray was set.
func (o OptNilRefreshCounterRequestArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilRefreshCounterRequestArray) Reset() {
	var v []RefreshCounterRequest
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilRefreshCounterRequestArray) SetTo(v []RefreshCounterRequest) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilRefreshCounterRequestArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilRefreshCounterRequestArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []RefreshCounterRequest
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilRefreshCounterRequestArray) Get() (v []RefreshCounterRequest, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilRefreshCounterRequestArray) Or(d []RefreshCounterRequest) []RefreshCounterRequest {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilReviewArray returns new OptNilReviewArray with value set to v.
func NewOptNilReviewArray(v []Review) OptNilReviewArray {
	return OptNilReviewArray{
		Value: v,
		Set:   true,
	}
}

// OptNilReviewArray is optional nullable []Review.
type OptNilReviewArray struct {
	Value []Review
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilReviewArray was set.
func (o OptNilReviewArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilReviewArray) Reset() {
	var v []Review
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilReviewArray) SetTo(v []Review) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilReviewArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilReviewArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Review
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilReviewArray) Get() (v []Review, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilReviewArray) Or(d []Review) []Review {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilStickerArray returns new OptNilStickerArray with value set to v.
func NewOptNilStickerArray(v []Sticker) OptNilStickerArray {
	return OptNilStickerArray{
		Value: v,
		Set:   true,
	}
}

// OptNilStickerArray is optional nullable []Sticker.
type OptNilStickerArray struct {
	Value []Sticker
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilStickerArray was set.
func (o OptNilStickerArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilStickerArray) Reset() {
	var v []Sticker
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilStickerArray) SetTo(v []Sticker) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilStickerArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilStickerArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Sticker
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilStickerArray) Get() (v []Sticker, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilStickerArray) Or(d []Sticker) []Sticker {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilStickerPackArray returns new OptNilStickerPackArray with value set to v.
func NewOptNilStickerPackArray(v []StickerPack) OptNilStickerPackArray {
	return OptNilStickerPackArray{
		Value: v,
		Set:   true,
	}
}

// OptNilStickerPackArray is optional nullable []StickerPack.
type OptNilStickerPackArray struct {
	Value []StickerPack
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilStickerPackArray was set.
func (o OptNilStickerPackArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilStickerPackArray) Reset() {
	var v []StickerPack
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilStickerPackArray) SetTo(v []StickerPack) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilStickerPackArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilStickerPackArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []StickerPack
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilStickerPackArray) Get() (v []StickerPack, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilStickerPackArray) Or(d []StickerPack) []StickerPack {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilString returns new OptNilString with value set to v.
func NewOptNilString(v string) OptNilString {
	return OptNilString{
		Value: v,
		Set:   true,
	}
}

// OptNilString is optional nullable string.
type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilString was set.
func (o OptNilString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilString) SetTo(v string) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilString) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilString) SetToNull() {
	o.Set = true
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilStringArray returns new OptNilStringArray with value set to v.
func NewOptNilStringArray(v []string) OptNilStringArray {
	return OptNilStringArray{
		Value: v,
		Set:   true,
	}
}

// OptNilStringArray is optional nullable []string.
type OptNilStringArray struct {
	Value []string
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilStringArray was set.
func (o OptNilStringArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilStringArray) Reset() {
	var v []string
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilStringArray) SetTo(v []string) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilStringArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilStringArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilStringArray) Get() (v []string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilStringArray) Or(d []string) []string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilThreadInfoArray returns new OptNilThreadInfoArray with value set to v.
func NewOptNilThreadInfoArray(v []ThreadInfo) OptNilThreadInfoArray {
	return OptNilThreadInfoArray{
		Value: v,
		Set:   true,
	}
}

// OptNilThreadInfoArray is optional nullable []ThreadInfo.
type OptNilThreadInfoArray struct {
	Value []ThreadInfo
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilThreadInfoArray was set.
func (o OptNilThreadInfoArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilThreadInfoArray) Reset() {
	var v []ThreadInfo
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilThreadInfoArray) SetTo(v []ThreadInfo) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilThreadInfoArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilThreadInfoArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []ThreadInfo
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilThreadInfoArray) Get() (v []ThreadInfo, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilThreadInfoArray) Or(d []ThreadInfo) []ThreadInfo {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilTransactionGiftReceivedArray returns new OptNilTransactionGiftReceivedArray with value set to v.
func NewOptNilTransactionGiftReceivedArray(v []TransactionGiftReceived) OptNilTransactionGiftReceivedArray {
	return OptNilTransactionGiftReceivedArray{
		Value: v,
		Set:   true,
	}
}

// OptNilTransactionGiftReceivedArray is optional nullable []TransactionGiftReceived.
type OptNilTransactionGiftReceivedArray struct {
	Value []TransactionGiftReceived
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilTransactionGiftReceivedArray was set.
func (o OptNilTransactionGiftReceivedArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilTransactionGiftReceivedArray) Reset() {
	var v []TransactionGiftReceived
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilTransactionGiftReceivedArray) SetTo(v []TransactionGiftReceived) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilTransactionGiftReceivedArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilTransactionGiftReceivedArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []TransactionGiftReceived
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilTransactionGiftReceivedArray) Get() (v []TransactionGiftReceived, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilTransactionGiftReceivedArray) Or(d []TransactionGiftReceived) []TransactionGiftReceived {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilUserArray returns new OptNilUserArray with value set to v.
func NewOptNilUserArray(v []User) OptNilUserArray {
	return OptNilUserArray{
		Value: v,
		Set:   true,
	}
}

// OptNilUserArray is optional nullable []User.
type OptNilUserArray struct {
	Value []User
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilUserArray was set.
func (o OptNilUserArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilUserArray) Reset() {
	var v []User
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilUserArray) SetTo(v []User) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilUserArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilUserArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []User
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilUserArray) Get() (v []User, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilUserArray) Or(d []User) []User {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilUserWrapperArray returns new OptNilUserWrapperArray with value set to v.
func NewOptNilUserWrapperArray(v []UserWrapper) OptNilUserWrapperArray {
	return OptNilUserWrapperArray{
		Value: v,
		Set:   true,
	}
}

// OptNilUserWrapperArray is optional nullable []UserWrapper.
type OptNilUserWrapperArray struct {
	Value []UserWrapper
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilUserWrapperArray was set.
func (o OptNilUserWrapperArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilUserWrapperArray) Reset() {
	var v []UserWrapper
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilUserWrapperArray) SetTo(v []UserWrapper) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilUserWrapperArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilUserWrapperArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []UserWrapper
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilUserWrapperArray) Get() (v []UserWrapper, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilUserWrapperArray) Or(d []UserWrapper) []UserWrapper {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilVideoArray returns new OptNilVideoArray with value set to v.
func NewOptNilVideoArray(v []Video) OptNilVideoArray {
	return OptNilVideoArray{
		Value: v,
		Set:   true,
	}
}

// OptNilVideoArray is optional nullable []Video.
type OptNilVideoArray struct {
	Value []Video
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilVideoArray was set.
func (o OptNilVideoArray) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilVideoArray) Reset() {
	var v []Video
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilVideoArray) SetTo(v []Video) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNilVideoArray) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNilVideoArray) SetToNull() {
	o.Set = true
	o.Null = true
	var v []Video
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilVideoArray) Get() (v []Video, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilVideoArray) Or(d []Video) []Video {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOnlineStatus returns new OptOnlineStatus with value set to v.
func NewOptOnlineStatus(v OnlineStatus) OptOnlineStatus {
	return OptOnlineStatus{
		Value: v,
		Set:   true,
	}
}

// OptOnlineStatus is optional OnlineStatus.
type OptOnlineStatus struct {
	Value OnlineStatus
	Set   bool
}

// IsSet returns true if OptOnlineStatus was set.
func (o OptOnlineStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOnlineStatus) Reset() {
	var v OnlineStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOnlineStatus) SetTo(v OnlineStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOnlineStatus) Get() (v OnlineStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOnlineStatus) Or(d OnlineStatus) OnlineStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOnlineStatusEnum returns new OptOnlineStatusEnum with value set to v.
func NewOptOnlineStatusEnum(v OnlineStatusEnum) OptOnlineStatusEnum {
	return OptOnlineStatusEnum{
		Value: v,
		Set:   true,
	}
}

// OptOnlineStatusEnum is optional OnlineStatusEnum.
type OptOnlineStatusEnum struct {
	Value OnlineStatusEnum
	Set   bool
}

// IsSet returns true if OptOnlineStatusEnum was set.
func (o OptOnlineStatusEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOnlineStatusEnum) Reset() {
	var v OnlineStatusEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOnlineStatusEnum) SetTo(v OnlineStatusEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOnlineStatusEnum) Get() (v OnlineStatusEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOnlineStatusEnum) Or(d OnlineStatusEnum) OnlineStatusEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptParentMessage returns new OptParentMessage with value set to v.
func NewOptParentMessage(v ParentMessage) OptParentMessage {
	return OptParentMessage{
		Value: v,
		Set:   true,
	}
}

// OptParentMessage is optional ParentMessage.
type OptParentMessage struct {
	Value ParentMessage
	Set   bool
}

// IsSet returns true if OptParentMessage was set.
func (o OptParentMessage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptParentMessage) Reset() {
	var v ParentMessage
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptParentMessage) SetTo(v ParentMessage) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptParentMessage) Get() (v ParentMessage, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptParentMessage) Or(d ParentMessage) ParentMessage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPost returns new OptPost with value set to v.
func NewOptPost(v Post) OptPost {
	return OptPost{
		Value: v,
		Set:   true,
	}
}

// OptPost is optional Post.
type OptPost struct {
	Value Post
	Set   bool
}

// IsSet returns true if OptPost was set.
func (o OptPost) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPost) Reset() {
	var v Post
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPost) SetTo(v Post) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPost) Get() (v Post, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPost) Or(d Post) Post {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPostType returns new OptPostType with value set to v.
func NewOptPostType(v PostType) OptPostType {
	return OptPostType{
		Value: v,
		Set:   true,
	}
}

// OptPostType is optional PostType.
type OptPostType struct {
	Value PostType
	Set   bool
}

// IsSet returns true if OptPostType was set.
func (o OptPostType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPostType) Reset() {
	var v PostType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPostType) SetTo(v PostType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPostType) Get() (v PostType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPostType) Or(d PostType) PostType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmChatRoom returns new OptRealmChatRoom with value set to v.
func NewOptRealmChatRoom(v RealmChatRoom) OptRealmChatRoom {
	return OptRealmChatRoom{
		Value: v,
		Set:   true,
	}
}

// OptRealmChatRoom is optional RealmChatRoom.
type OptRealmChatRoom struct {
	Value RealmChatRoom
	Set   bool
}

// IsSet returns true if OptRealmChatRoom was set.
func (o OptRealmChatRoom) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmChatRoom) Reset() {
	var v RealmChatRoom
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmChatRoom) SetTo(v RealmChatRoom) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmChatRoom) Get() (v RealmChatRoom, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmChatRoom) Or(d RealmChatRoom) RealmChatRoom {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmConferenceCall returns new OptRealmConferenceCall with value set to v.
func NewOptRealmConferenceCall(v RealmConferenceCall) OptRealmConferenceCall {
	return OptRealmConferenceCall{
		Value: v,
		Set:   true,
	}
}

// OptRealmConferenceCall is optional RealmConferenceCall.
type OptRealmConferenceCall struct {
	Value RealmConferenceCall
	Set   bool
}

// IsSet returns true if OptRealmConferenceCall was set.
func (o OptRealmConferenceCall) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmConferenceCall) Reset() {
	var v RealmConferenceCall
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmConferenceCall) SetTo(v RealmConferenceCall) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmConferenceCall) Get() (v RealmConferenceCall, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmConferenceCall) Or(d RealmConferenceCall) RealmConferenceCall {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmGame returns new OptRealmGame with value set to v.
func NewOptRealmGame(v RealmGame) OptRealmGame {
	return OptRealmGame{
		Value: v,
		Set:   true,
	}
}

// OptRealmGame is optional RealmGame.
type OptRealmGame struct {
	Value RealmGame
	Set   bool
}

// IsSet returns true if OptRealmGame was set.
func (o OptRealmGame) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmGame) Reset() {
	var v RealmGame
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmGame) SetTo(v RealmGame) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmGame) Get() (v RealmGame, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmGame) Or(d RealmGame) RealmGame {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmGenre returns new OptRealmGenre with value set to v.
func NewOptRealmGenre(v RealmGenre) OptRealmGenre {
	return OptRealmGenre{
		Value: v,
		Set:   true,
	}
}

// OptRealmGenre is optional RealmGenre.
type OptRealmGenre struct {
	Value RealmGenre
	Set   bool
}

// IsSet returns true if OptRealmGenre was set.
func (o OptRealmGenre) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmGenre) Reset() {
	var v RealmGenre
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmGenre) SetTo(v RealmGenre) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmGenre) Get() (v RealmGenre, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmGenre) Or(d RealmGenre) RealmGenre {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmGift returns new OptRealmGift with value set to v.
func NewOptRealmGift(v RealmGift) OptRealmGift {
	return OptRealmGift{
		Value: v,
		Set:   true,
	}
}

// OptRealmGift is optional RealmGift.
type OptRealmGift struct {
	Value RealmGift
	Set   bool
}

// IsSet returns true if OptRealmGift was set.
func (o OptRealmGift) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmGift) Reset() {
	var v RealmGift
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmGift) SetTo(v RealmGift) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmGift) Get() (v RealmGift, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmGift) Or(d RealmGift) RealmGift {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmGiftingAbility returns new OptRealmGiftingAbility with value set to v.
func NewOptRealmGiftingAbility(v RealmGiftingAbility) OptRealmGiftingAbility {
	return OptRealmGiftingAbility{
		Value: v,
		Set:   true,
	}
}

// OptRealmGiftingAbility is optional RealmGiftingAbility.
type OptRealmGiftingAbility struct {
	Value RealmGiftingAbility
	Set   bool
}

// IsSet returns true if OptRealmGiftingAbility was set.
func (o OptRealmGiftingAbility) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmGiftingAbility) Reset() {
	var v RealmGiftingAbility
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmGiftingAbility) SetTo(v RealmGiftingAbility) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmGiftingAbility) Get() (v RealmGiftingAbility, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmGiftingAbility) Or(d RealmGiftingAbility) RealmGiftingAbility {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmPlatformDetails returns new OptRealmPlatformDetails with value set to v.
func NewOptRealmPlatformDetails(v RealmPlatformDetails) OptRealmPlatformDetails {
	return OptRealmPlatformDetails{
		Value: v,
		Set:   true,
	}
}

// OptRealmPlatformDetails is optional RealmPlatformDetails.
type OptRealmPlatformDetails struct {
	Value RealmPlatformDetails
	Set   bool
}

// IsSet returns true if OptRealmPlatformDetails was set.
func (o OptRealmPlatformDetails) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmPlatformDetails) Reset() {
	var v RealmPlatformDetails
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmPlatformDetails) SetTo(v RealmPlatformDetails) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmPlatformDetails) Get() (v RealmPlatformDetails, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmPlatformDetails) Or(d RealmPlatformDetails) RealmPlatformDetails {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRealmUser returns new OptRealmUser with value set to v.
func NewOptRealmUser(v RealmUser) OptRealmUser {
	return OptRealmUser{
		Value: v,
		Set:   true,
	}
}

// OptRealmUser is optional RealmUser.
type OptRealmUser struct {
	Value RealmUser
	Set   bool
}

// IsSet returns true if OptRealmUser was set.
func (o OptRealmUser) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRealmUser) Reset() {
	var v RealmUser
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRealmUser) SetTo(v RealmUser) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRealmUser) Get() (v RealmUser, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRealmUser) Or(d RealmUser) RealmUser {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRepostReqMessageTags returns new OptRepostReqMessageTags with value set to v.
func NewOptRepostReqMessageTags(v *RepostReqMessageTags) OptRepostReqMessageTags {
	return OptRepostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptRepostReqMessageTags is optional *RepostReqMessageTags.
type OptRepostReqMessageTags struct {
	Value *RepostReqMessageTags
	Set   bool
}

// IsSet returns true if OptRepostReqMessageTags was set.
func (o OptRepostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRepostReqMessageTags) Reset() {
	var v *RepostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRepostReqMessageTags) SetTo(v *RepostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRepostReqMessageTags) Get() (v *RepostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRepostReqMessageTags) Or(d *RepostReqMessageTags) *RepostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRepostReqSharedURL returns new OptRepostReqSharedURL with value set to v.
func NewOptRepostReqSharedURL(v *RepostReqSharedURL) OptRepostReqSharedURL {
	return OptRepostReqSharedURL{
		Value: v,
		Set:   true,
	}
}

// OptRepostReqSharedURL is optional *RepostReqSharedURL.
type OptRepostReqSharedURL struct {
	Value *RepostReqSharedURL
	Set   bool
}

// IsSet returns true if OptRepostReqSharedURL was set.
func (o OptRepostReqSharedURL) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRepostReqSharedURL) Reset() {
	var v *RepostReqSharedURL
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRepostReqSharedURL) SetTo(v *RepostReqSharedURL) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRepostReqSharedURL) Get() (v *RepostReqSharedURL, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRepostReqSharedURL) Or(d *RepostReqSharedURL) *RepostReqSharedURL {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptReviewRestriction returns new OptReviewRestriction with value set to v.
func NewOptReviewRestriction(v ReviewRestriction) OptReviewRestriction {
	return OptReviewRestriction{
		Value: v,
		Set:   true,
	}
}

// OptReviewRestriction is optional ReviewRestriction.
type OptReviewRestriction struct {
	Value ReviewRestriction
	Set   bool
}

// IsSet returns true if OptReviewRestriction was set.
func (o OptReviewRestriction) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptReviewRestriction) Reset() {
	var v ReviewRestriction
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptReviewRestriction) SetTo(v ReviewRestriction) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptReviewRestriction) Get() (v ReviewRestriction, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptReviewRestriction) Or(d ReviewRestriction) ReviewRestriction {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSearchCriteria returns new OptSearchCriteria with value set to v.
func NewOptSearchCriteria(v SearchCriteria) OptSearchCriteria {
	return OptSearchCriteria{
		Value: v,
		Set:   true,
	}
}

// OptSearchCriteria is optional SearchCriteria.
type OptSearchCriteria struct {
	Value SearchCriteria
	Set   bool
}

// IsSet returns true if OptSearchCriteria was set.
func (o OptSearchCriteria) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSearchCriteria) Reset() {
	var v SearchCriteria
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSearchCriteria) SetTo(v SearchCriteria) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSearchCriteria) Get() (v SearchCriteria, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSearchCriteria) Or(d SearchCriteria) SearchCriteria {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSetting returns new OptSetting with value set to v.
func NewOptSetting(v Setting) OptSetting {
	return OptSetting{
		Value: v,
		Set:   true,
	}
}

// OptSetting is optional Setting.
type OptSetting struct {
	Value Setting
	Set   bool
}

// IsSet returns true if OptSetting was set.
func (o OptSetting) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSetting) Reset() {
	var v Setting
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSetting) SetTo(v Setting) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSetting) Get() (v Setting, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSetting) Or(d Setting) Setting {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSettings returns new OptSettings with value set to v.
func NewOptSettings(v Settings) OptSettings {
	return OptSettings{
		Value: v,
		Set:   true,
	}
}

// OptSettings is optional Settings.
type OptSettings struct {
	Value Settings
	Set   bool
}

// IsSet returns true if OptSettings was set.
func (o OptSettings) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSettings) Reset() {
	var v Settings
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSettings) SetTo(v Settings) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSettings) Get() (v Settings, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSettings) Or(d Settings) Settings {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptShareable returns new OptShareable with value set to v.
func NewOptShareable(v Shareable) OptShareable {
	return OptShareable{
		Value: v,
		Set:   true,
	}
}

// OptShareable is optional Shareable.
type OptShareable struct {
	Value Shareable
	Set   bool
}

// IsSet returns true if OptShareable was set.
func (o OptShareable) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptShareable) Reset() {
	var v Shareable
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptShareable) SetTo(v Shareable) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptShareable) Get() (v Shareable, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptShareable) Or(d Shareable) Shareable {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSharedUrl returns new OptSharedUrl with value set to v.
func NewOptSharedUrl(v SharedUrl) OptSharedUrl {
	return OptSharedUrl{
		Value: v,
		Set:   true,
	}
}

// OptSharedUrl is optional SharedUrl.
type OptSharedUrl struct {
	Value SharedUrl
	Set   bool
}

// IsSet returns true if OptSharedUrl was set.
func (o OptSharedUrl) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSharedUrl) Reset() {
	var v SharedUrl
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSharedUrl) SetTo(v SharedUrl) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSharedUrl) Get() (v SharedUrl, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSharedUrl) Or(d SharedUrl) SharedUrl {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSignaturePayload returns new OptSignaturePayload with value set to v.
func NewOptSignaturePayload(v SignaturePayload) OptSignaturePayload {
	return OptSignaturePayload{
		Value: v,
		Set:   true,
	}
}

// OptSignaturePayload is optional SignaturePayload.
type OptSignaturePayload struct {
	Value SignaturePayload
	Set   bool
}

// IsSet returns true if OptSignaturePayload was set.
func (o OptSignaturePayload) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSignaturePayload) Reset() {
	var v SignaturePayload
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSignaturePayload) SetTo(v SignaturePayload) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSignaturePayload) Get() (v SignaturePayload, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSignaturePayload) Or(d SignaturePayload) SignaturePayload {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSnsInfo returns new OptSnsInfo with value set to v.
func NewOptSnsInfo(v SnsInfo) OptSnsInfo {
	return OptSnsInfo{
		Value: v,
		Set:   true,
	}
}

// OptSnsInfo is optional SnsInfo.
type OptSnsInfo struct {
	Value SnsInfo
	Set   bool
}

// IsSet returns true if OptSnsInfo was set.
func (o OptSnsInfo) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSnsInfo) Reset() {
	var v SnsInfo
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSnsInfo) SetTo(v SnsInfo) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSnsInfo) Get() (v SnsInfo, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSnsInfo) Or(d SnsInfo) SnsInfo {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSticker returns new OptSticker with value set to v.
func NewOptSticker(v Sticker) OptSticker {
	return OptSticker{
		Value: v,
		Set:   true,
	}
}

// OptSticker is optional Sticker.
type OptSticker struct {
	Value Sticker
	Set   bool
}

// IsSet returns true if OptSticker was set.
func (o OptSticker) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSticker) Reset() {
	var v Sticker
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSticker) SetTo(v Sticker) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSticker) Get() (v Sticker, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSticker) Or(d Sticker) Sticker {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSurvey returns new OptSurvey with value set to v.
func NewOptSurvey(v Survey) OptSurvey {
	return OptSurvey{
		Value: v,
		Set:   true,
	}
}

// OptSurvey is optional Survey.
type OptSurvey struct {
	Value Survey
	Set   bool
}

// IsSet returns true if OptSurvey was set.
func (o OptSurvey) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSurvey) Reset() {
	var v Survey
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSurvey) SetTo(v Survey) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSurvey) Get() (v Survey, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSurvey) Or(d Survey) Survey {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptThreadInfo returns new OptThreadInfo with value set to v.
func NewOptThreadInfo(v ThreadInfo) OptThreadInfo {
	return OptThreadInfo{
		Value: v,
		Set:   true,
	}
}

// OptThreadInfo is optional ThreadInfo.
type OptThreadInfo struct {
	Value ThreadInfo
	Set   bool
}

// IsSet returns true if OptThreadInfo was set.
func (o OptThreadInfo) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptThreadInfo) Reset() {
	var v ThreadInfo
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptThreadInfo) SetTo(v ThreadInfo) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptThreadInfo) Get() (v ThreadInfo, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptThreadInfo) Or(d ThreadInfo) ThreadInfo {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTimelineSettings returns new OptTimelineSettings with value set to v.
func NewOptTimelineSettings(v TimelineSettings) OptTimelineSettings {
	return OptTimelineSettings{
		Value: v,
		Set:   true,
	}
}

// OptTimelineSettings is optional TimelineSettings.
type OptTimelineSettings struct {
	Value TimelineSettings
	Set   bool
}

// IsSet returns true if OptTimelineSettings was set.
func (o OptTimelineSettings) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTimelineSettings) Reset() {
	var v TimelineSettings
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTimelineSettings) SetTo(v TimelineSettings) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTimelineSettings) Get() (v TimelineSettings, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTimelineSettings) Or(d TimelineSettings) TimelineSettings {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTitle returns new OptTitle with value set to v.
func NewOptTitle(v Title) OptTitle {
	return OptTitle{
		Value: v,
		Set:   true,
	}
}

// OptTitle is optional Title.
type OptTitle struct {
	Value Title
	Set   bool
}

// IsSet returns true if OptTitle was set.
func (o OptTitle) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTitle) Reset() {
	var v Title
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTitle) SetTo(v Title) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTitle) Get() (v Title, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTitle) Or(d Title) Title {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTransactionGiftReceivedItem returns new OptTransactionGiftReceivedItem with value set to v.
func NewOptTransactionGiftReceivedItem(v TransactionGiftReceivedItem) OptTransactionGiftReceivedItem {
	return OptTransactionGiftReceivedItem{
		Value: v,
		Set:   true,
	}
}

// OptTransactionGiftReceivedItem is optional TransactionGiftReceivedItem.
type OptTransactionGiftReceivedItem struct {
	Value TransactionGiftReceivedItem
	Set   bool
}

// IsSet returns true if OptTransactionGiftReceivedItem was set.
func (o OptTransactionGiftReceivedItem) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTransactionGiftReceivedItem) Reset() {
	var v TransactionGiftReceivedItem
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTransactionGiftReceivedItem) SetTo(v TransactionGiftReceivedItem) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTransactionGiftReceivedItem) Get() (v TransactionGiftReceivedItem, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTransactionGiftReceivedItem) Or(d TransactionGiftReceivedItem) TransactionGiftReceivedItem {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTwoFaAuthRequiredDTO returns new OptTwoFaAuthRequiredDTO with value set to v.
func NewOptTwoFaAuthRequiredDTO(v TwoFaAuthRequiredDTO) OptTwoFaAuthRequiredDTO {
	return OptTwoFaAuthRequiredDTO{
		Value: v,
		Set:   true,
	}
}

// OptTwoFaAuthRequiredDTO is optional TwoFaAuthRequiredDTO.
type OptTwoFaAuthRequiredDTO struct {
	Value TwoFaAuthRequiredDTO
	Set   bool
}

// IsSet returns true if OptTwoFaAuthRequiredDTO was set.
func (o OptTwoFaAuthRequiredDTO) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTwoFaAuthRequiredDTO) Reset() {
	var v TwoFaAuthRequiredDTO
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTwoFaAuthRequiredDTO) SetTo(v TwoFaAuthRequiredDTO) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTwoFaAuthRequiredDTO) Get() (v TwoFaAuthRequiredDTO, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTwoFaAuthRequiredDTO) Or(d TwoFaAuthRequiredDTO) TwoFaAuthRequiredDTO {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUpdatePostReqMessageTags returns new OptUpdatePostReqMessageTags with value set to v.
func NewOptUpdatePostReqMessageTags(v *UpdatePostReqMessageTags) OptUpdatePostReqMessageTags {
	return OptUpdatePostReqMessageTags{
		Value: v,
		Set:   true,
	}
}

// OptUpdatePostReqMessageTags is optional *UpdatePostReqMessageTags.
type OptUpdatePostReqMessageTags struct {
	Value *UpdatePostReqMessageTags
	Set   bool
}

// IsSet returns true if OptUpdatePostReqMessageTags was set.
func (o OptUpdatePostReqMessageTags) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUpdatePostReqMessageTags) Reset() {
	var v *UpdatePostReqMessageTags
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUpdatePostReqMessageTags) SetTo(v *UpdatePostReqMessageTags) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUpdatePostReqMessageTags) Get() (v *UpdatePostReqMessageTags, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUpdatePostReqMessageTags) Or(d *UpdatePostReqMessageTags) *UpdatePostReqMessageTags {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserConnectedBy returns new OptUserConnectedBy with value set to v.
func NewOptUserConnectedBy(v *UserConnectedBy) OptUserConnectedBy {
	return OptUserConnectedBy{
		Value: v,
		Set:   true,
	}
}

// OptUserConnectedBy is optional *UserConnectedBy.
type OptUserConnectedBy struct {
	Value *UserConnectedBy
	Set   bool
}

// IsSet returns true if OptUserConnectedBy was set.
func (o OptUserConnectedBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserConnectedBy) Reset() {
	var v *UserConnectedBy
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserConnectedBy) SetTo(v *UserConnectedBy) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserConnectedBy) Get() (v *UserConnectedBy, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserConnectedBy) Or(d *UserConnectedBy) *UserConnectedBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserCountry returns new OptUserCountry with value set to v.
func NewOptUserCountry(v *UserCountry) OptUserCountry {
	return OptUserCountry{
		Value: v,
		Set:   true,
	}
}

// OptUserCountry is optional *UserCountry.
type OptUserCountry struct {
	Value *UserCountry
	Set   bool
}

// IsSet returns true if OptUserCountry was set.
func (o OptUserCountry) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserCountry) Reset() {
	var v *UserCountry
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserCountry) SetTo(v *UserCountry) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserCountry) Get() (v *UserCountry, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserCountry) Or(d *UserCountry) *UserCountry {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserSetting returns new OptUserSetting with value set to v.
func NewOptUserSetting(v UserSetting) OptUserSetting {
	return OptUserSetting{
		Value: v,
		Set:   true,
	}
}

// OptUserSetting is optional UserSetting.
type OptUserSetting struct {
	Value UserSetting
	Set   bool
}

// IsSet returns true if OptUserSetting was set.
func (o OptUserSetting) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserSetting) Reset() {
	var v UserSetting
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserSetting) SetTo(v UserSetting) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserSetting) Get() (v UserSetting, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserSetting) Or(d UserSetting) UserSetting {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserUserDTO returns new OptUserUserDTO with value set to v.
func NewOptUserUserDTO(v UserUserDTO) OptUserUserDTO {
	return OptUserUserDTO{
		Value: v,
		Set:   true,
	}
}

// OptUserUserDTO is optional UserUserDTO.
type OptUserUserDTO struct {
	Value UserUserDTO
	Set   bool
}

// IsSet returns true if OptUserUserDTO was set.
func (o OptUserUserDTO) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserUserDTO) Reset() {
	var v UserUserDTO
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserUserDTO) SetTo(v UserUserDTO) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserUserDTO) Get() (v UserUserDTO, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserUserDTO) Or(d UserUserDTO) UserUserDTO {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/ParentMessage
type ParentMessage struct {
	Attachment          OptNilString   `json:"attachment"`
	AttachmentThumbnail OptNilString   `json:"attachment_thumbnail"`
	CreatedAt           OptNilInt64    `json:"created_at"`
	FontSize            OptNilInt32    `json:"font_size"`
	GIF                 OptGifImage    `json:"gif"`
	ID                  OptNilInt64    `json:"id"`
	MessageType         OptMessageType `json:"message_type"`
	RoomID              OptNilInt64    `json:"room_id"`
	Sticker             OptSticker     `json:"sticker"`
	Text                OptNilString   `json:"text"`
	UserID              OptNilInt64    `json:"user_id"`
	VideoThumbnailURL   OptNilString   `json:"video_thumbnail_url"`
	VideoURL            OptNilString   `json:"video_url"`
}

// GetAttachment returns the value of Attachment.
func (s *ParentMessage) GetAttachment() OptNilString {
	return s.Attachment
}

// GetAttachmentThumbnail returns the value of AttachmentThumbnail.
func (s *ParentMessage) GetAttachmentThumbnail() OptNilString {
	return s.AttachmentThumbnail
}

// GetCreatedAt returns the value of CreatedAt.
func (s *ParentMessage) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetFontSize returns the value of FontSize.
func (s *ParentMessage) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGIF returns the value of GIF.
func (s *ParentMessage) GetGIF() OptGifImage {
	return s.GIF
}

// GetID returns the value of ID.
func (s *ParentMessage) GetID() OptNilInt64 {
	return s.ID
}

// GetMessageType returns the value of MessageType.
func (s *ParentMessage) GetMessageType() OptMessageType {
	return s.MessageType
}

// GetRoomID returns the value of RoomID.
func (s *ParentMessage) GetRoomID() OptNilInt64 {
	return s.RoomID
}

// GetSticker returns the value of Sticker.
func (s *ParentMessage) GetSticker() OptSticker {
	return s.Sticker
}

// GetText returns the value of Text.
func (s *ParentMessage) GetText() OptNilString {
	return s.Text
}

// GetUserID returns the value of UserID.
func (s *ParentMessage) GetUserID() OptNilInt64 {
	return s.UserID
}

// GetVideoThumbnailURL returns the value of VideoThumbnailURL.
func (s *ParentMessage) GetVideoThumbnailURL() OptNilString {
	return s.VideoThumbnailURL
}

// GetVideoURL returns the value of VideoURL.
func (s *ParentMessage) GetVideoURL() OptNilString {
	return s.VideoURL
}

// SetAttachment sets the value of Attachment.
func (s *ParentMessage) SetAttachment(val OptNilString) {
	s.Attachment = val
}

// SetAttachmentThumbnail sets the value of AttachmentThumbnail.
func (s *ParentMessage) SetAttachmentThumbnail(val OptNilString) {
	s.AttachmentThumbnail = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *ParentMessage) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetFontSize sets the value of FontSize.
func (s *ParentMessage) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGIF sets the value of GIF.
func (s *ParentMessage) SetGIF(val OptGifImage) {
	s.GIF = val
}

// SetID sets the value of ID.
func (s *ParentMessage) SetID(val OptNilInt64) {
	s.ID = val
}

// SetMessageType sets the value of MessageType.
func (s *ParentMessage) SetMessageType(val OptMessageType) {
	s.MessageType = val
}

// SetRoomID sets the value of RoomID.
func (s *ParentMessage) SetRoomID(val OptNilInt64) {
	s.RoomID = val
}

// SetSticker sets the value of Sticker.
func (s *ParentMessage) SetSticker(val OptSticker) {
	s.Sticker = val
}

// SetText sets the value of Text.
func (s *ParentMessage) SetText(val OptNilString) {
	s.Text = val
}

// SetUserID sets the value of UserID.
func (s *ParentMessage) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// SetVideoThumbnailURL sets the value of VideoThumbnailURL.
func (s *ParentMessage) SetVideoThumbnailURL(val OptNilString) {
	s.VideoThumbnailURL = val
}

// SetVideoURL sets the value of VideoURL.
func (s *ParentMessage) SetVideoURL(val OptNilString) {
	s.VideoURL = val
}

// PinChatRoomNoContent is response for PinChatRoom operation.
type PinChatRoomNoContent struct{}

// PinGroupHighlightPostNoContent is response for PinGroupHighlightPost operation.
type PinGroupHighlightPostNoContent struct{}

// PinGroupNoContent is response for PinGroup operation.
type PinGroupNoContent struct{}

// PinGroupPostNoContent is response for PinGroupPost operation.
type PinGroupPostNoContent struct{}

type PinGroupPostReq struct {
	GroupID OptInt64 `json:"group_id"`
	PostID  OptInt64 `json:"post_id"`
}

// GetGroupID returns the value of GroupID.
func (s *PinGroupPostReq) GetGroupID() OptInt64 {
	return s.GroupID
}

// GetPostID returns the value of PostID.
func (s *PinGroupPostReq) GetPostID() OptInt64 {
	return s.PostID
}

// SetGroupID sets the value of GroupID.
func (s *PinGroupPostReq) SetGroupID(val OptInt64) {
	s.GroupID = val
}

// SetPostID sets the value of PostID.
func (s *PinGroupPostReq) SetPostID(val OptInt64) {
	s.PostID = val
}

type PinGroupReq struct {
	ID OptInt64 `json:"id"`
}

// GetID returns the value of ID.
func (s *PinGroupReq) GetID() OptInt64 {
	return s.ID
}

// SetID sets the value of ID.
func (s *PinGroupReq) SetID(val OptInt64) {
	s.ID = val
}

// PinPostNoContent is response for PinPost operation.
type PinPostNoContent struct{}

type PinPostReq struct {
	ID OptInt64 `json:"id"`
}

// GetID returns the value of ID.
func (s *PinPostReq) GetID() OptInt64 {
	return s.ID
}

// SetID sets the value of ID.
func (s *PinPostReq) SetID(val OptInt64) {
	s.ID = val
}

// PinReviewNoContent is response for PinReview operation.
type PinReviewNoContent struct{}

type PinReviewReq struct {
	ID OptInt64 `json:"id"`
}

// GetID returns the value of ID.
func (s *PinReviewReq) GetID() OptInt64 {
	return s.ID
}

// SetID sets the value of ID.
func (s *PinReviewReq) SetID(val OptInt64) {
	s.ID = val
}

// PingAliveNoContent is response for PingAlive operation.
type PingAliveNoContent struct{}

// Ref: #/components/schemas/PolicyAgreementsResponse
type PolicyAgreementsResponse struct {
	LatestDotmoneyAgreed      OptNilBool `json:"latest_dotmoney_agreed"`
	LatestEmpl2Agreed         OptNilBool `json:"latest_empl2_agreed"`
	LatestPrivacyPolicyAgreed OptNilBool `json:"latest_privacy_policy_agreed"`
	LatestTermsOfUseAgreed    OptNilBool `json:"latest_terms_of_use_agreed"`
}

// GetLatestDotmoneyAgreed returns the value of LatestDotmoneyAgreed.
func (s *PolicyAgreementsResponse) GetLatestDotmoneyAgreed() OptNilBool {
	return s.LatestDotmoneyAgreed
}

// GetLatestEmpl2Agreed returns the value of LatestEmpl2Agreed.
func (s *PolicyAgreementsResponse) GetLatestEmpl2Agreed() OptNilBool {
	return s.LatestEmpl2Agreed
}

// GetLatestPrivacyPolicyAgreed returns the value of LatestPrivacyPolicyAgreed.
func (s *PolicyAgreementsResponse) GetLatestPrivacyPolicyAgreed() OptNilBool {
	return s.LatestPrivacyPolicyAgreed
}

// GetLatestTermsOfUseAgreed returns the value of LatestTermsOfUseAgreed.
func (s *PolicyAgreementsResponse) GetLatestTermsOfUseAgreed() OptNilBool {
	return s.LatestTermsOfUseAgreed
}

// SetLatestDotmoneyAgreed sets the value of LatestDotmoneyAgreed.
func (s *PolicyAgreementsResponse) SetLatestDotmoneyAgreed(val OptNilBool) {
	s.LatestDotmoneyAgreed = val
}

// SetLatestEmpl2Agreed sets the value of LatestEmpl2Agreed.
func (s *PolicyAgreementsResponse) SetLatestEmpl2Agreed(val OptNilBool) {
	s.LatestEmpl2Agreed = val
}

// SetLatestPrivacyPolicyAgreed sets the value of LatestPrivacyPolicyAgreed.
func (s *PolicyAgreementsResponse) SetLatestPrivacyPolicyAgreed(val OptNilBool) {
	s.LatestPrivacyPolicyAgreed = val
}

// SetLatestTermsOfUseAgreed sets the value of LatestTermsOfUseAgreed.
func (s *PolicyAgreementsResponse) SetLatestTermsOfUseAgreed(val OptNilBool) {
	s.LatestTermsOfUseAgreed = val
}

// Ref: #/components/schemas/PopularWord
type PopularWord struct {
	ID   OptNilInt64  `json:"id"`
	Type OptNilString `json:"type"`
	Word OptNilString `json:"word"`
}

// GetID returns the value of ID.
func (s *PopularWord) GetID() OptNilInt64 {
	return s.ID
}

// GetType returns the value of Type.
func (s *PopularWord) GetType() OptNilString {
	return s.Type
}

// GetWord returns the value of Word.
func (s *PopularWord) GetWord() OptNilString {
	return s.Word
}

// SetID sets the value of ID.
func (s *PopularWord) SetID(val OptNilInt64) {
	s.ID = val
}

// SetType sets the value of Type.
func (s *PopularWord) SetType(val OptNilString) {
	s.Type = val
}

// SetWord sets the value of Word.
func (s *PopularWord) SetWord(val OptNilString) {
	s.Word = val
}

// Ref: #/components/schemas/PopularWordsResponse
type PopularWordsResponse struct {
	PopularWords OptNilPopularWordArray `json:"popular_words"`
}

// GetPopularWords returns the value of PopularWords.
func (s *PopularWordsResponse) GetPopularWords() OptNilPopularWordArray {
	return s.PopularWords
}

// SetPopularWords sets the value of PopularWords.
func (s *PopularWordsResponse) SetPopularWords(val OptNilPopularWordArray) {
	s.PopularWords = val
}

// Ref: #/components/schemas/Post
type Post struct {
	Attachment           OptNilString           `json:"attachment"`
	Attachment2          OptNilString           `json:"attachment_2"`
	Attachment2Thumbnail OptNilString           `json:"attachment_2_thumbnail"`
	Attachment3          OptNilString           `json:"attachment_3"`
	Attachment3Thumbnail OptNilString           `json:"attachment_3_thumbnail"`
	Attachment4          OptNilString           `json:"attachment_4"`
	Attachment4Thumbnail OptNilString           `json:"attachment_4_thumbnail"`
	Attachment5          OptNilString           `json:"attachment_5"`
	Attachment5Thumbnail OptNilString           `json:"attachment_5_thumbnail"`
	Attachment6          OptNilString           `json:"attachment_6"`
	Attachment6Thumbnail OptNilString           `json:"attachment_6_thumbnail"`
	Attachment7          OptNilString           `json:"attachment_7"`
	Attachment7Thumbnail OptNilString           `json:"attachment_7_thumbnail"`
	Attachment8          OptNilString           `json:"attachment_8"`
	Attachment8Thumbnail OptNilString           `json:"attachment_8_thumbnail"`
	Attachment9          OptNilString           `json:"attachment_9"`
	Attachment9Thumbnail OptNilString           `json:"attachment_9_thumbnail"`
	AttachmentThumbnail  OptNilString           `json:"attachment_thumbnail"`
	Color                OptNilInt32            `json:"color"`
	ConferenceCall       OptRealmConferenceCall `json:"conference_call"`
	ConversationID       OptNilInt64            `json:"conversation_id"`
	CreatedAt            OptNilInt64            `json:"created_at"`
	EditedAt             OptNilInt64            `json:"edited_at"`
	FontSize             OptNilInt32            `json:"font_size"`
	GiftsCount           OptNilGiftCountArray   `json:"gifts_count"`
	Group                OptGroup               `json:"group"`
	GroupID              OptNilInt64            `json:"group_id"`
	Highlighted          OptNilBool             `json:"highlighted"`
	ID                   OptNilInt64            `json:"id"`
	InReplyTo            OptNilInt64            `json:"in_reply_to"`
	InReplyToPost        *Post                  `json:"in_reply_to_post"`
	InReplyToPostCount   OptNilInt32            `json:"in_reply_to_post_count"`
	IsFailToSend         OptNilBool             `json:"is_fail_to_send"`
	IsMuted              OptNilBool             `json:"is_muted"`
	Liked                OptNilBool             `json:"liked"`
	Likers               OptNilRealmUserArray   `json:"likers"`
	LikersCount          OptNilInt32            `json:"likers_count"`
	LikesCount           OptNilInt32            `json:"likes_count"`
	Mentions             OptNilRealmUserArray   `json:"mentions"`
	MessageTags          OptNilMessageTagArray  `json:"message_tags"`
	PostType             OptPostType            `json:"post_type"`
	ReportedCount        OptNilInt32            `json:"reported_count"`
	Repostable           OptNilBool             `json:"repostable"`
	Reposted             OptNilBool             `json:"reposted"`
	RepostsCount         OptNilInt32            `json:"reposts_count"`
	Shareable            *Shareable             `json:"shareable"`
	SharedThread         *ThreadInfo            `json:"shared_thread"`
	SharedURL            OptSharedUrl           `json:"shared_url"`
	Survey               OptSurvey              `json:"survey"`
	Tag                  OptNilString           `json:"tag"`
	Text                 OptNilString           `json:"text"`
	Thread               *ThreadInfo            `json:"thread"`
	ThreadID             OptNilInt64            `json:"thread_id"`
	TotalGiftsCount      OptNilInt32            `json:"total_gifts_count"`
	UpdatedAt            OptNilInt64            `json:"updated_at"`
	User                 OptRealmUser           `json:"user"`
	Videos               OptNilVideoArray       `json:"videos"`
}

// GetAttachment returns the value of Attachment.
func (s *Post) GetAttachment() OptNilString {
	return s.Attachment
}

// GetAttachment2 returns the value of Attachment2.
func (s *Post) GetAttachment2() OptNilString {
	return s.Attachment2
}

// GetAttachment2Thumbnail returns the value of Attachment2Thumbnail.
func (s *Post) GetAttachment2Thumbnail() OptNilString {
	return s.Attachment2Thumbnail
}

// GetAttachment3 returns the value of Attachment3.
func (s *Post) GetAttachment3() OptNilString {
	return s.Attachment3
}

// GetAttachment3Thumbnail returns the value of Attachment3Thumbnail.
func (s *Post) GetAttachment3Thumbnail() OptNilString {
	return s.Attachment3Thumbnail
}

// GetAttachment4 returns the value of Attachment4.
func (s *Post) GetAttachment4() OptNilString {
	return s.Attachment4
}

// GetAttachment4Thumbnail returns the value of Attachment4Thumbnail.
func (s *Post) GetAttachment4Thumbnail() OptNilString {
	return s.Attachment4Thumbnail
}

// GetAttachment5 returns the value of Attachment5.
func (s *Post) GetAttachment5() OptNilString {
	return s.Attachment5
}

// GetAttachment5Thumbnail returns the value of Attachment5Thumbnail.
func (s *Post) GetAttachment5Thumbnail() OptNilString {
	return s.Attachment5Thumbnail
}

// GetAttachment6 returns the value of Attachment6.
func (s *Post) GetAttachment6() OptNilString {
	return s.Attachment6
}

// GetAttachment6Thumbnail returns the value of Attachment6Thumbnail.
func (s *Post) GetAttachment6Thumbnail() OptNilString {
	return s.Attachment6Thumbnail
}

// GetAttachment7 returns the value of Attachment7.
func (s *Post) GetAttachment7() OptNilString {
	return s.Attachment7
}

// GetAttachment7Thumbnail returns the value of Attachment7Thumbnail.
func (s *Post) GetAttachment7Thumbnail() OptNilString {
	return s.Attachment7Thumbnail
}

// GetAttachment8 returns the value of Attachment8.
func (s *Post) GetAttachment8() OptNilString {
	return s.Attachment8
}

// GetAttachment8Thumbnail returns the value of Attachment8Thumbnail.
func (s *Post) GetAttachment8Thumbnail() OptNilString {
	return s.Attachment8Thumbnail
}

// GetAttachment9 returns the value of Attachment9.
func (s *Post) GetAttachment9() OptNilString {
	return s.Attachment9
}

// GetAttachment9Thumbnail returns the value of Attachment9Thumbnail.
func (s *Post) GetAttachment9Thumbnail() OptNilString {
	return s.Attachment9Thumbnail
}

// GetAttachmentThumbnail returns the value of AttachmentThumbnail.
func (s *Post) GetAttachmentThumbnail() OptNilString {
	return s.AttachmentThumbnail
}

// GetColor returns the value of Color.
func (s *Post) GetColor() OptNilInt32 {
	return s.Color
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *Post) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetConversationID returns the value of ConversationID.
func (s *Post) GetConversationID() OptNilInt64 {
	return s.ConversationID
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Post) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetEditedAt returns the value of EditedAt.
func (s *Post) GetEditedAt() OptNilInt64 {
	return s.EditedAt
}

// GetFontSize returns the value of FontSize.
func (s *Post) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGiftsCount returns the value of GiftsCount.
func (s *Post) GetGiftsCount() OptNilGiftCountArray {
	return s.GiftsCount
}

// GetGroup returns the value of Group.
func (s *Post) GetGroup() OptGroup {
	return s.Group
}

// GetGroupID returns the value of GroupID.
func (s *Post) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetHighlighted returns the value of Highlighted.
func (s *Post) GetHighlighted() OptNilBool {
	return s.Highlighted
}

// GetID returns the value of ID.
func (s *Post) GetID() OptNilInt64 {
	return s.ID
}

// GetInReplyTo returns the value of InReplyTo.
func (s *Post) GetInReplyTo() OptNilInt64 {
	return s.InReplyTo
}

// GetInReplyToPost returns the value of InReplyToPost.
func (s *Post) GetInReplyToPost() *Post {
	return s.InReplyToPost
}

// GetInReplyToPostCount returns the value of InReplyToPostCount.
func (s *Post) GetInReplyToPostCount() OptNilInt32 {
	return s.InReplyToPostCount
}

// GetIsFailToSend returns the value of IsFailToSend.
func (s *Post) GetIsFailToSend() OptNilBool {
	return s.IsFailToSend
}

// GetIsMuted returns the value of IsMuted.
func (s *Post) GetIsMuted() OptNilBool {
	return s.IsMuted
}

// GetLiked returns the value of Liked.
func (s *Post) GetLiked() OptNilBool {
	return s.Liked
}

// GetLikers returns the value of Likers.
func (s *Post) GetLikers() OptNilRealmUserArray {
	return s.Likers
}

// GetLikersCount returns the value of LikersCount.
func (s *Post) GetLikersCount() OptNilInt32 {
	return s.LikersCount
}

// GetLikesCount returns the value of LikesCount.
func (s *Post) GetLikesCount() OptNilInt32 {
	return s.LikesCount
}

// GetMentions returns the value of Mentions.
func (s *Post) GetMentions() OptNilRealmUserArray {
	return s.Mentions
}

// GetMessageTags returns the value of MessageTags.
func (s *Post) GetMessageTags() OptNilMessageTagArray {
	return s.MessageTags
}

// GetPostType returns the value of PostType.
func (s *Post) GetPostType() OptPostType {
	return s.PostType
}

// GetReportedCount returns the value of ReportedCount.
func (s *Post) GetReportedCount() OptNilInt32 {
	return s.ReportedCount
}

// GetRepostable returns the value of Repostable.
func (s *Post) GetRepostable() OptNilBool {
	return s.Repostable
}

// GetReposted returns the value of Reposted.
func (s *Post) GetReposted() OptNilBool {
	return s.Reposted
}

// GetRepostsCount returns the value of RepostsCount.
func (s *Post) GetRepostsCount() OptNilInt32 {
	return s.RepostsCount
}

// GetShareable returns the value of Shareable.
func (s *Post) GetShareable() *Shareable {
	return s.Shareable
}

// GetSharedThread returns the value of SharedThread.
func (s *Post) GetSharedThread() *ThreadInfo {
	return s.SharedThread
}

// GetSharedURL returns the value of SharedURL.
func (s *Post) GetSharedURL() OptSharedUrl {
	return s.SharedURL
}

// GetSurvey returns the value of Survey.
func (s *Post) GetSurvey() OptSurvey {
	return s.Survey
}

// GetTag returns the value of Tag.
func (s *Post) GetTag() OptNilString {
	return s.Tag
}

// GetText returns the value of Text.
func (s *Post) GetText() OptNilString {
	return s.Text
}

// GetThread returns the value of Thread.
func (s *Post) GetThread() *ThreadInfo {
	return s.Thread
}

// GetThreadID returns the value of ThreadID.
func (s *Post) GetThreadID() OptNilInt64 {
	return s.ThreadID
}

// GetTotalGiftsCount returns the value of TotalGiftsCount.
func (s *Post) GetTotalGiftsCount() OptNilInt32 {
	return s.TotalGiftsCount
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Post) GetUpdatedAt() OptNilInt64 {
	return s.UpdatedAt
}

// GetUser returns the value of User.
func (s *Post) GetUser() OptRealmUser {
	return s.User
}

// GetVideos returns the value of Videos.
func (s *Post) GetVideos() OptNilVideoArray {
	return s.Videos
}

// SetAttachment sets the value of Attachment.
func (s *Post) SetAttachment(val OptNilString) {
	s.Attachment = val
}

// SetAttachment2 sets the value of Attachment2.
func (s *Post) SetAttachment2(val OptNilString) {
	s.Attachment2 = val
}

// SetAttachment2Thumbnail sets the value of Attachment2Thumbnail.
func (s *Post) SetAttachment2Thumbnail(val OptNilString) {
	s.Attachment2Thumbnail = val
}

// SetAttachment3 sets the value of Attachment3.
func (s *Post) SetAttachment3(val OptNilString) {
	s.Attachment3 = val
}

// SetAttachment3Thumbnail sets the value of Attachment3Thumbnail.
func (s *Post) SetAttachment3Thumbnail(val OptNilString) {
	s.Attachment3Thumbnail = val
}

// SetAttachment4 sets the value of Attachment4.
func (s *Post) SetAttachment4(val OptNilString) {
	s.Attachment4 = val
}

// SetAttachment4Thumbnail sets the value of Attachment4Thumbnail.
func (s *Post) SetAttachment4Thumbnail(val OptNilString) {
	s.Attachment4Thumbnail = val
}

// SetAttachment5 sets the value of Attachment5.
func (s *Post) SetAttachment5(val OptNilString) {
	s.Attachment5 = val
}

// SetAttachment5Thumbnail sets the value of Attachment5Thumbnail.
func (s *Post) SetAttachment5Thumbnail(val OptNilString) {
	s.Attachment5Thumbnail = val
}

// SetAttachment6 sets the value of Attachment6.
func (s *Post) SetAttachment6(val OptNilString) {
	s.Attachment6 = val
}

// SetAttachment6Thumbnail sets the value of Attachment6Thumbnail.
func (s *Post) SetAttachment6Thumbnail(val OptNilString) {
	s.Attachment6Thumbnail = val
}

// SetAttachment7 sets the value of Attachment7.
func (s *Post) SetAttachment7(val OptNilString) {
	s.Attachment7 = val
}

// SetAttachment7Thumbnail sets the value of Attachment7Thumbnail.
func (s *Post) SetAttachment7Thumbnail(val OptNilString) {
	s.Attachment7Thumbnail = val
}

// SetAttachment8 sets the value of Attachment8.
func (s *Post) SetAttachment8(val OptNilString) {
	s.Attachment8 = val
}

// SetAttachment8Thumbnail sets the value of Attachment8Thumbnail.
func (s *Post) SetAttachment8Thumbnail(val OptNilString) {
	s.Attachment8Thumbnail = val
}

// SetAttachment9 sets the value of Attachment9.
func (s *Post) SetAttachment9(val OptNilString) {
	s.Attachment9 = val
}

// SetAttachment9Thumbnail sets the value of Attachment9Thumbnail.
func (s *Post) SetAttachment9Thumbnail(val OptNilString) {
	s.Attachment9Thumbnail = val
}

// SetAttachmentThumbnail sets the value of AttachmentThumbnail.
func (s *Post) SetAttachmentThumbnail(val OptNilString) {
	s.AttachmentThumbnail = val
}

// SetColor sets the value of Color.
func (s *Post) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *Post) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetConversationID sets the value of ConversationID.
func (s *Post) SetConversationID(val OptNilInt64) {
	s.ConversationID = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Post) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetEditedAt sets the value of EditedAt.
func (s *Post) SetEditedAt(val OptNilInt64) {
	s.EditedAt = val
}

// SetFontSize sets the value of FontSize.
func (s *Post) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGiftsCount sets the value of GiftsCount.
func (s *Post) SetGiftsCount(val OptNilGiftCountArray) {
	s.GiftsCount = val
}

// SetGroup sets the value of Group.
func (s *Post) SetGroup(val OptGroup) {
	s.Group = val
}

// SetGroupID sets the value of GroupID.
func (s *Post) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetHighlighted sets the value of Highlighted.
func (s *Post) SetHighlighted(val OptNilBool) {
	s.Highlighted = val
}

// SetID sets the value of ID.
func (s *Post) SetID(val OptNilInt64) {
	s.ID = val
}

// SetInReplyTo sets the value of InReplyTo.
func (s *Post) SetInReplyTo(val OptNilInt64) {
	s.InReplyTo = val
}

// SetInReplyToPost sets the value of InReplyToPost.
func (s *Post) SetInReplyToPost(val *Post) {
	s.InReplyToPost = val
}

// SetInReplyToPostCount sets the value of InReplyToPostCount.
func (s *Post) SetInReplyToPostCount(val OptNilInt32) {
	s.InReplyToPostCount = val
}

// SetIsFailToSend sets the value of IsFailToSend.
func (s *Post) SetIsFailToSend(val OptNilBool) {
	s.IsFailToSend = val
}

// SetIsMuted sets the value of IsMuted.
func (s *Post) SetIsMuted(val OptNilBool) {
	s.IsMuted = val
}

// SetLiked sets the value of Liked.
func (s *Post) SetLiked(val OptNilBool) {
	s.Liked = val
}

// SetLikers sets the value of Likers.
func (s *Post) SetLikers(val OptNilRealmUserArray) {
	s.Likers = val
}

// SetLikersCount sets the value of LikersCount.
func (s *Post) SetLikersCount(val OptNilInt32) {
	s.LikersCount = val
}

// SetLikesCount sets the value of LikesCount.
func (s *Post) SetLikesCount(val OptNilInt32) {
	s.LikesCount = val
}

// SetMentions sets the value of Mentions.
func (s *Post) SetMentions(val OptNilRealmUserArray) {
	s.Mentions = val
}

// SetMessageTags sets the value of MessageTags.
func (s *Post) SetMessageTags(val OptNilMessageTagArray) {
	s.MessageTags = val
}

// SetPostType sets the value of PostType.
func (s *Post) SetPostType(val OptPostType) {
	s.PostType = val
}

// SetReportedCount sets the value of ReportedCount.
func (s *Post) SetReportedCount(val OptNilInt32) {
	s.ReportedCount = val
}

// SetRepostable sets the value of Repostable.
func (s *Post) SetRepostable(val OptNilBool) {
	s.Repostable = val
}

// SetReposted sets the value of Reposted.
func (s *Post) SetReposted(val OptNilBool) {
	s.Reposted = val
}

// SetRepostsCount sets the value of RepostsCount.
func (s *Post) SetRepostsCount(val OptNilInt32) {
	s.RepostsCount = val
}

// SetShareable sets the value of Shareable.
func (s *Post) SetShareable(val *Shareable) {
	s.Shareable = val
}

// SetSharedThread sets the value of SharedThread.
func (s *Post) SetSharedThread(val *ThreadInfo) {
	s.SharedThread = val
}

// SetSharedURL sets the value of SharedURL.
func (s *Post) SetSharedURL(val OptSharedUrl) {
	s.SharedURL = val
}

// SetSurvey sets the value of Survey.
func (s *Post) SetSurvey(val OptSurvey) {
	s.Survey = val
}

// SetTag sets the value of Tag.
func (s *Post) SetTag(val OptNilString) {
	s.Tag = val
}

// SetText sets the value of Text.
func (s *Post) SetText(val OptNilString) {
	s.Text = val
}

// SetThread sets the value of Thread.
func (s *Post) SetThread(val *ThreadInfo) {
	s.Thread = val
}

// SetThreadID sets the value of ThreadID.
func (s *Post) SetThreadID(val OptNilInt64) {
	s.ThreadID = val
}

// SetTotalGiftsCount sets the value of TotalGiftsCount.
func (s *Post) SetTotalGiftsCount(val OptNilInt32) {
	s.TotalGiftsCount = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Post) SetUpdatedAt(val OptNilInt64) {
	s.UpdatedAt = val
}

// SetUser sets the value of User.
func (s *Post) SetUser(val OptRealmUser) {
	s.User = val
}

// SetVideos sets the value of Videos.
func (s *Post) SetVideos(val OptNilVideoArray) {
	s.Videos = val
}

// Ref: #/components/schemas/PostLikersResponse
type PostLikersResponse struct {
	LastID OptNilInt64          `json:"last_id"`
	Users  OptNilRealmUserArray `json:"users"`
}

// GetLastID returns the value of LastID.
func (s *PostLikersResponse) GetLastID() OptNilInt64 {
	return s.LastID
}

// GetUsers returns the value of Users.
func (s *PostLikersResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetLastID sets the value of LastID.
func (s *PostLikersResponse) SetLastID(val OptNilInt64) {
	s.LastID = val
}

// SetUsers sets the value of Users.
func (s *PostLikersResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// Ref: #/components/schemas/PostResponse
type PostResponse struct {
	Post OptPost `json:"post"`
}

// GetPost returns the value of Post.
func (s *PostResponse) GetPost() OptPost {
	return s.Post
}

// SetPost sets the value of Post.
func (s *PostResponse) SetPost(val OptPost) {
	s.Post = val
}

// Ref: #/components/schemas/PostTag
type PostTag struct {
	ID                OptNilInt64  `json:"id"`
	PostHashtagsCount OptNilInt32  `json:"post_hashtags_count"`
	Tag               OptNilString `json:"tag"`
}

// GetID returns the value of ID.
func (s *PostTag) GetID() OptNilInt64 {
	return s.ID
}

// GetPostHashtagsCount returns the value of PostHashtagsCount.
func (s *PostTag) GetPostHashtagsCount() OptNilInt32 {
	return s.PostHashtagsCount
}

// GetTag returns the value of Tag.
func (s *PostTag) GetTag() OptNilString {
	return s.Tag
}

// SetID sets the value of ID.
func (s *PostTag) SetID(val OptNilInt64) {
	s.ID = val
}

// SetPostHashtagsCount sets the value of PostHashtagsCount.
func (s *PostTag) SetPostHashtagsCount(val OptNilInt32) {
	s.PostHashtagsCount = val
}

// SetTag sets the value of Tag.
func (s *PostTag) SetTag(val OptNilString) {
	s.Tag = val
}

// Ref: #/components/schemas/PostTagsResponse
type PostTagsResponse struct {
	Tags OptNilPostTagArray `json:"tags"`
}

// GetTags returns the value of Tags.
func (s *PostTagsResponse) GetTags() OptNilPostTagArray {
	return s.Tags
}

// SetTags sets the value of Tags.
func (s *PostTagsResponse) SetTags(val OptNilPostTagArray) {
	s.Tags = val
}

// Ref: #/components/schemas/PostType
type PostType string

const (
	PostTypeCall            PostType = "call"
	PostTypeImage           PostType = "image"
	PostTypeVideo           PostType = "video"
	PostTypeSurvey          PostType = "survey"
	PostTypeRepost          PostType = "repost"
	PostTypeThread          PostType = "thread"
	PostTypeShareableGroup  PostType = "shareable_group"
	PostTypeShareableURL    PostType = "shareable_url"
	PostTypeYoutube         PostType = "youtube"
	PostTypeShareableThread PostType = "shareable_thread"
	PostTypeSharepal        PostType = "sharepal"
	PostTypeUndefined       PostType = "undefined"
	PostTypeText            PostType = "text"
)

// AllValues returns all PostType values.
func (PostType) AllValues() []PostType {
	return []PostType{
		PostTypeCall,
		PostTypeImage,
		PostTypeVideo,
		PostTypeSurvey,
		PostTypeRepost,
		PostTypeThread,
		PostTypeShareableGroup,
		PostTypeShareableURL,
		PostTypeYoutube,
		PostTypeShareableThread,
		PostTypeSharepal,
		PostTypeUndefined,
		PostTypeText,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s PostType) MarshalText() ([]byte, error) {
	switch s {
	case PostTypeCall:
		return []byte(s), nil
	case PostTypeImage:
		return []byte(s), nil
	case PostTypeVideo:
		return []byte(s), nil
	case PostTypeSurvey:
		return []byte(s), nil
	case PostTypeRepost:
		return []byte(s), nil
	case PostTypeThread:
		return []byte(s), nil
	case PostTypeShareableGroup:
		return []byte(s), nil
	case PostTypeShareableURL:
		return []byte(s), nil
	case PostTypeYoutube:
		return []byte(s), nil
	case PostTypeShareableThread:
		return []byte(s), nil
	case PostTypeSharepal:
		return []byte(s), nil
	case PostTypeUndefined:
		return []byte(s), nil
	case PostTypeText:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *PostType) UnmarshalText(data []byte) error {
	switch PostType(data) {
	case PostTypeCall:
		*s = PostTypeCall
		return nil
	case PostTypeImage:
		*s = PostTypeImage
		return nil
	case PostTypeVideo:
		*s = PostTypeVideo
		return nil
	case PostTypeSurvey:
		*s = PostTypeSurvey
		return nil
	case PostTypeRepost:
		*s = PostTypeRepost
		return nil
	case PostTypeThread:
		*s = PostTypeThread
		return nil
	case PostTypeShareableGroup:
		*s = PostTypeShareableGroup
		return nil
	case PostTypeShareableURL:
		*s = PostTypeShareableURL
		return nil
	case PostTypeYoutube:
		*s = PostTypeYoutube
		return nil
	case PostTypeShareableThread:
		*s = PostTypeShareableThread
		return nil
	case PostTypeSharepal:
		*s = PostTypeSharepal
		return nil
	case PostTypeUndefined:
		*s = PostTypeUndefined
		return nil
	case PostTypeText:
		*s = PostTypeText
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/PostsResponse
type PostsResponse struct {
	HasMoreHotPosts OptNilBool      `json:"has_more_hot_posts"`
	NextPageValue   OptNilString    `json:"next_page_value"`
	PinnedPosts     OptNilPostArray `json:"pinned_posts"`
	Posts           OptNilPostArray `json:"posts"`
}

// GetHasMoreHotPosts returns the value of HasMoreHotPosts.
func (s *PostsResponse) GetHasMoreHotPosts() OptNilBool {
	return s.HasMoreHotPosts
}

// GetNextPageValue returns the value of NextPageValue.
func (s *PostsResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetPinnedPosts returns the value of PinnedPosts.
func (s *PostsResponse) GetPinnedPosts() OptNilPostArray {
	return s.PinnedPosts
}

// GetPosts returns the value of Posts.
func (s *PostsResponse) GetPosts() OptNilPostArray {
	return s.Posts
}

// SetHasMoreHotPosts sets the value of HasMoreHotPosts.
func (s *PostsResponse) SetHasMoreHotPosts(val OptNilBool) {
	s.HasMoreHotPosts = val
}

// SetNextPageValue sets the value of NextPageValue.
func (s *PostsResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetPinnedPosts sets the value of PinnedPosts.
func (s *PostsResponse) SetPinnedPosts(val OptNilPostArray) {
	s.PinnedPosts = val
}

// SetPosts sets the value of Posts.
func (s *PostsResponse) SetPosts(val OptNilPostArray) {
	s.Posts = val
}

// Ref: #/components/schemas/PresignedUrl
type PresignedUrl struct {
	Filename OptNilString `json:"filename"`
	URL      OptNilString `json:"url"`
}

// GetFilename returns the value of Filename.
func (s *PresignedUrl) GetFilename() OptNilString {
	return s.Filename
}

// GetURL returns the value of URL.
func (s *PresignedUrl) GetURL() OptNilString {
	return s.URL
}

// SetFilename sets the value of Filename.
func (s *PresignedUrl) SetFilename(val OptNilString) {
	s.Filename = val
}

// SetURL sets the value of URL.
func (s *PresignedUrl) SetURL(val OptNilString) {
	s.URL = val
}

// Ref: #/components/schemas/PresignedUrlResponse
type PresignedUrlResponse struct {
	PresignedURL OptNilString `json:"presigned_url"`
}

// GetPresignedURL returns the value of PresignedURL.
func (s *PresignedUrlResponse) GetPresignedURL() OptNilString {
	return s.PresignedURL
}

// SetPresignedURL sets the value of PresignedURL.
func (s *PresignedUrlResponse) SetPresignedURL(val OptNilString) {
	s.PresignedURL = val
}

// Ref: #/components/schemas/PresignedUrlsResponse
type PresignedUrlsResponse struct {
	PresignedUrls OptNilPresignedUrlArray `json:"presigned_urls"`
}

// GetPresignedUrls returns the value of PresignedUrls.
func (s *PresignedUrlsResponse) GetPresignedUrls() OptNilPresignedUrlArray {
	return s.PresignedUrls
}

// SetPresignedUrls sets the value of PresignedUrls.
func (s *PresignedUrlsResponse) SetPresignedUrls(val OptNilPresignedUrlArray) {
	s.PresignedUrls = val
}

// Ref: #/components/schemas/ReadAttachmentRequest
type ReadAttachmentRequest struct {
	AttachmentMsgIds OptNilInt64Array `json:"attachment_msg_ids"`
}

// GetAttachmentMsgIds returns the value of AttachmentMsgIds.
func (s *ReadAttachmentRequest) GetAttachmentMsgIds() OptNilInt64Array {
	return s.AttachmentMsgIds
}

// SetAttachmentMsgIds sets the value of AttachmentMsgIds.
func (s *ReadAttachmentRequest) SetAttachmentMsgIds(val OptNilInt64Array) {
	s.AttachmentMsgIds = val
}

// ReadChatAttachmentsNoContent is response for ReadChatAttachments operation.
type ReadChatAttachmentsNoContent struct{}

// ReadChatMessageNoContent is response for ReadChatMessage operation.
type ReadChatMessageNoContent struct{}

// ReadChatVideosNoContent is response for ReadChatVideos operation.
type ReadChatVideosNoContent struct{}

type ReadChatVideosReq struct {
	VideoMsgIds []int64 `json:"video_msg_ids[]"`
}

// GetVideoMsgIds returns the value of VideoMsgIds.
func (s *ReadChatVideosReq) GetVideoMsgIds() []int64 {
	return s.VideoMsgIds
}

// SetVideoMsgIds sets the value of VideoMsgIds.
func (s *ReadChatVideosReq) SetVideoMsgIds(val []int64) {
	s.VideoMsgIds = val
}

// Ref: #/components/schemas/realm_ChatRoom
type RealmChatRoom struct {
	Background  OptNilString           `json:"background"`
	ID          OptNilInt64            `json:"id"`
	IsGroup     OptNilBool             `json:"is_group"`
	IsRequest   OptNilBool             `json:"is_request"`
	LastMessage OptChatRoomLastMessage `json:"last_message"`
	Members     OptNilRealmUserArray   `json:"members"`
	Name        OptNilString           `json:"name"`
	Owner       OptRealmUser           `json:"owner"`
	UnreadCount OptNilInt32            `json:"unread_count"`
	UpdatedAt   OptNilInt64            `json:"updated_at"`
	UserSetting OptUserSetting         `json:"user_setting"`
}

// GetBackground returns the value of Background.
func (s *RealmChatRoom) GetBackground() OptNilString {
	return s.Background
}

// GetID returns the value of ID.
func (s *RealmChatRoom) GetID() OptNilInt64 {
	return s.ID
}

// GetIsGroup returns the value of IsGroup.
func (s *RealmChatRoom) GetIsGroup() OptNilBool {
	return s.IsGroup
}

// GetIsRequest returns the value of IsRequest.
func (s *RealmChatRoom) GetIsRequest() OptNilBool {
	return s.IsRequest
}

// GetLastMessage returns the value of LastMessage.
func (s *RealmChatRoom) GetLastMessage() OptChatRoomLastMessage {
	return s.LastMessage
}

// GetMembers returns the value of Members.
func (s *RealmChatRoom) GetMembers() OptNilRealmUserArray {
	return s.Members
}

// GetName returns the value of Name.
func (s *RealmChatRoom) GetName() OptNilString {
	return s.Name
}

// GetOwner returns the value of Owner.
func (s *RealmChatRoom) GetOwner() OptRealmUser {
	return s.Owner
}

// GetUnreadCount returns the value of UnreadCount.
func (s *RealmChatRoom) GetUnreadCount() OptNilInt32 {
	return s.UnreadCount
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *RealmChatRoom) GetUpdatedAt() OptNilInt64 {
	return s.UpdatedAt
}

// GetUserSetting returns the value of UserSetting.
func (s *RealmChatRoom) GetUserSetting() OptUserSetting {
	return s.UserSetting
}

// SetBackground sets the value of Background.
func (s *RealmChatRoom) SetBackground(val OptNilString) {
	s.Background = val
}

// SetID sets the value of ID.
func (s *RealmChatRoom) SetID(val OptNilInt64) {
	s.ID = val
}

// SetIsGroup sets the value of IsGroup.
func (s *RealmChatRoom) SetIsGroup(val OptNilBool) {
	s.IsGroup = val
}

// SetIsRequest sets the value of IsRequest.
func (s *RealmChatRoom) SetIsRequest(val OptNilBool) {
	s.IsRequest = val
}

// SetLastMessage sets the value of LastMessage.
func (s *RealmChatRoom) SetLastMessage(val OptChatRoomLastMessage) {
	s.LastMessage = val
}

// SetMembers sets the value of Members.
func (s *RealmChatRoom) SetMembers(val OptNilRealmUserArray) {
	s.Members = val
}

// SetName sets the value of Name.
func (s *RealmChatRoom) SetName(val OptNilString) {
	s.Name = val
}

// SetOwner sets the value of Owner.
func (s *RealmChatRoom) SetOwner(val OptRealmUser) {
	s.Owner = val
}

// SetUnreadCount sets the value of UnreadCount.
func (s *RealmChatRoom) SetUnreadCount(val OptNilInt32) {
	s.UnreadCount = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *RealmChatRoom) SetUpdatedAt(val OptNilInt64) {
	s.UpdatedAt = val
}

// SetUserSetting sets the value of UserSetting.
func (s *RealmChatRoom) SetUserSetting(val OptUserSetting) {
	s.UserSetting = val
}

// Ref: #/components/schemas/realm_ConferenceCall
type RealmConferenceCall struct {
	Active                   OptNilBool                        `json:"active"`
	AgoraChannel             OptNilString                      `json:"agora_channel"`
	AgoraToken               OptNilString                      `json:"agora_token"`
	AnonymousCallUsersCount  OptNilInt32                       `json:"anonymous_call_users_count"`
	BumpParams               OptConferenceCallBumpParams       `json:"bump_params"`
	CallType                 OptNilString                      `json:"call_type"`
	ConferenceCallUserRoles  OptNilConferenceCallUserRoleArray `json:"conference_call_user_roles"`
	ConferenceCallUsers      OptNilRealmUserArray              `json:"conference_call_users"`
	ConferenceCallUsersCount OptNilInt32                       `json:"conference_call_users_count"`
	DurationSeconds          OptNilInt64                       `json:"duration_seconds"`
	Game                     OptRealmGame                      `json:"game"`
	Genre                    OptRealmGenre                     `json:"genre"`
	GroupID                  OptNilInt64                       `json:"group_id"`
	ID                       OptNilInt64                       `json:"id"`
	JoinableBy               OptNilString                      `json:"joinable_by"`
	MaxParticipants          OptNilInt32                       `json:"max_participants"`
	PostID                   OptNilInt64                       `json:"post_id"`
	Server                   OptNilString                      `json:"server"`
}

// GetActive returns the value of Active.
func (s *RealmConferenceCall) GetActive() OptNilBool {
	return s.Active
}

// GetAgoraChannel returns the value of AgoraChannel.
func (s *RealmConferenceCall) GetAgoraChannel() OptNilString {
	return s.AgoraChannel
}

// GetAgoraToken returns the value of AgoraToken.
func (s *RealmConferenceCall) GetAgoraToken() OptNilString {
	return s.AgoraToken
}

// GetAnonymousCallUsersCount returns the value of AnonymousCallUsersCount.
func (s *RealmConferenceCall) GetAnonymousCallUsersCount() OptNilInt32 {
	return s.AnonymousCallUsersCount
}

// GetBumpParams returns the value of BumpParams.
func (s *RealmConferenceCall) GetBumpParams() OptConferenceCallBumpParams {
	return s.BumpParams
}

// GetCallType returns the value of CallType.
func (s *RealmConferenceCall) GetCallType() OptNilString {
	return s.CallType
}

// GetConferenceCallUserRoles returns the value of ConferenceCallUserRoles.
func (s *RealmConferenceCall) GetConferenceCallUserRoles() OptNilConferenceCallUserRoleArray {
	return s.ConferenceCallUserRoles
}

// GetConferenceCallUsers returns the value of ConferenceCallUsers.
func (s *RealmConferenceCall) GetConferenceCallUsers() OptNilRealmUserArray {
	return s.ConferenceCallUsers
}

// GetConferenceCallUsersCount returns the value of ConferenceCallUsersCount.
func (s *RealmConferenceCall) GetConferenceCallUsersCount() OptNilInt32 {
	return s.ConferenceCallUsersCount
}

// GetDurationSeconds returns the value of DurationSeconds.
func (s *RealmConferenceCall) GetDurationSeconds() OptNilInt64 {
	return s.DurationSeconds
}

// GetGame returns the value of Game.
func (s *RealmConferenceCall) GetGame() OptRealmGame {
	return s.Game
}

// GetGenre returns the value of Genre.
func (s *RealmConferenceCall) GetGenre() OptRealmGenre {
	return s.Genre
}

// GetGroupID returns the value of GroupID.
func (s *RealmConferenceCall) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetID returns the value of ID.
func (s *RealmConferenceCall) GetID() OptNilInt64 {
	return s.ID
}

// GetJoinableBy returns the value of JoinableBy.
func (s *RealmConferenceCall) GetJoinableBy() OptNilString {
	return s.JoinableBy
}

// GetMaxParticipants returns the value of MaxParticipants.
func (s *RealmConferenceCall) GetMaxParticipants() OptNilInt32 {
	return s.MaxParticipants
}

// GetPostID returns the value of PostID.
func (s *RealmConferenceCall) GetPostID() OptNilInt64 {
	return s.PostID
}

// GetServer returns the value of Server.
func (s *RealmConferenceCall) GetServer() OptNilString {
	return s.Server
}

// SetActive sets the value of Active.
func (s *RealmConferenceCall) SetActive(val OptNilBool) {
	s.Active = val
}

// SetAgoraChannel sets the value of AgoraChannel.
func (s *RealmConferenceCall) SetAgoraChannel(val OptNilString) {
	s.AgoraChannel = val
}

// SetAgoraToken sets the value of AgoraToken.
func (s *RealmConferenceCall) SetAgoraToken(val OptNilString) {
	s.AgoraToken = val
}

// SetAnonymousCallUsersCount sets the value of AnonymousCallUsersCount.
func (s *RealmConferenceCall) SetAnonymousCallUsersCount(val OptNilInt32) {
	s.AnonymousCallUsersCount = val
}

// SetBumpParams sets the value of BumpParams.
func (s *RealmConferenceCall) SetBumpParams(val OptConferenceCallBumpParams) {
	s.BumpParams = val
}

// SetCallType sets the value of CallType.
func (s *RealmConferenceCall) SetCallType(val OptNilString) {
	s.CallType = val
}

// SetConferenceCallUserRoles sets the value of ConferenceCallUserRoles.
func (s *RealmConferenceCall) SetConferenceCallUserRoles(val OptNilConferenceCallUserRoleArray) {
	s.ConferenceCallUserRoles = val
}

// SetConferenceCallUsers sets the value of ConferenceCallUsers.
func (s *RealmConferenceCall) SetConferenceCallUsers(val OptNilRealmUserArray) {
	s.ConferenceCallUsers = val
}

// SetConferenceCallUsersCount sets the value of ConferenceCallUsersCount.
func (s *RealmConferenceCall) SetConferenceCallUsersCount(val OptNilInt32) {
	s.ConferenceCallUsersCount = val
}

// SetDurationSeconds sets the value of DurationSeconds.
func (s *RealmConferenceCall) SetDurationSeconds(val OptNilInt64) {
	s.DurationSeconds = val
}

// SetGame sets the value of Game.
func (s *RealmConferenceCall) SetGame(val OptRealmGame) {
	s.Game = val
}

// SetGenre sets the value of Genre.
func (s *RealmConferenceCall) SetGenre(val OptRealmGenre) {
	s.Genre = val
}

// SetGroupID sets the value of GroupID.
func (s *RealmConferenceCall) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetID sets the value of ID.
func (s *RealmConferenceCall) SetID(val OptNilInt64) {
	s.ID = val
}

// SetJoinableBy sets the value of JoinableBy.
func (s *RealmConferenceCall) SetJoinableBy(val OptNilString) {
	s.JoinableBy = val
}

// SetMaxParticipants sets the value of MaxParticipants.
func (s *RealmConferenceCall) SetMaxParticipants(val OptNilInt32) {
	s.MaxParticipants = val
}

// SetPostID sets the value of PostID.
func (s *RealmConferenceCall) SetPostID(val OptNilInt64) {
	s.PostID = val
}

// SetServer sets the value of Server.
func (s *RealmConferenceCall) SetServer(val OptNilString) {
	s.Server = val
}

// Ref: #/components/schemas/realm_Game
type RealmGame struct {
	IconURL         OptNilString            `json:"icon_url"`
	ID              OptNilInt64             `json:"id"`
	PlatformDetails OptRealmPlatformDetails `json:"platform_details"`
	Title           OptNilString            `json:"title"`
}

// GetIconURL returns the value of IconURL.
func (s *RealmGame) GetIconURL() OptNilString {
	return s.IconURL
}

// GetID returns the value of ID.
func (s *RealmGame) GetID() OptNilInt64 {
	return s.ID
}

// GetPlatformDetails returns the value of PlatformDetails.
func (s *RealmGame) GetPlatformDetails() OptRealmPlatformDetails {
	return s.PlatformDetails
}

// GetTitle returns the value of Title.
func (s *RealmGame) GetTitle() OptNilString {
	return s.Title
}

// SetIconURL sets the value of IconURL.
func (s *RealmGame) SetIconURL(val OptNilString) {
	s.IconURL = val
}

// SetID sets the value of ID.
func (s *RealmGame) SetID(val OptNilInt64) {
	s.ID = val
}

// SetPlatformDetails sets the value of PlatformDetails.
func (s *RealmGame) SetPlatformDetails(val OptRealmPlatformDetails) {
	s.PlatformDetails = val
}

// SetTitle sets the value of Title.
func (s *RealmGame) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/realm_Genre
type RealmGenre struct {
	IconURL OptNilString `json:"icon_url"`
	ID      OptNilInt64  `json:"id"`
	Title   OptNilString `json:"title"`
}

// GetIconURL returns the value of IconURL.
func (s *RealmGenre) GetIconURL() OptNilString {
	return s.IconURL
}

// GetID returns the value of ID.
func (s *RealmGenre) GetID() OptNilInt64 {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *RealmGenre) GetTitle() OptNilString {
	return s.Title
}

// SetIconURL sets the value of IconURL.
func (s *RealmGenre) SetIconURL(val OptNilString) {
	s.IconURL = val
}

// SetID sets the value of ID.
func (s *RealmGenre) SetID(val OptNilInt64) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *RealmGenre) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/realm_Gift
type RealmGift struct {
	Currency      OptNilString  `json:"currency"`
	Icon          OptNilString  `json:"icon"`
	IconThumbnail OptNilString  `json:"icon_thumbnail"`
	ID            OptNilInt64   `json:"id"`
	Price         OptNilFloat64 `json:"price"`
	Slug          OptNilString  `json:"slug"`
	Title         OptNilString  `json:"title"`
}

// GetCurrency returns the value of Currency.
func (s *RealmGift) GetCurrency() OptNilString {
	return s.Currency
}

// GetIcon returns the value of Icon.
func (s *RealmGift) GetIcon() OptNilString {
	return s.Icon
}

// GetIconThumbnail returns the value of IconThumbnail.
func (s *RealmGift) GetIconThumbnail() OptNilString {
	return s.IconThumbnail
}

// GetID returns the value of ID.
func (s *RealmGift) GetID() OptNilInt64 {
	return s.ID
}

// GetPrice returns the value of Price.
func (s *RealmGift) GetPrice() OptNilFloat64 {
	return s.Price
}

// GetSlug returns the value of Slug.
func (s *RealmGift) GetSlug() OptNilString {
	return s.Slug
}

// GetTitle returns the value of Title.
func (s *RealmGift) GetTitle() OptNilString {
	return s.Title
}

// SetCurrency sets the value of Currency.
func (s *RealmGift) SetCurrency(val OptNilString) {
	s.Currency = val
}

// SetIcon sets the value of Icon.
func (s *RealmGift) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetIconThumbnail sets the value of IconThumbnail.
func (s *RealmGift) SetIconThumbnail(val OptNilString) {
	s.IconThumbnail = val
}

// SetID sets the value of ID.
func (s *RealmGift) SetID(val OptNilInt64) {
	s.ID = val
}

// SetPrice sets the value of Price.
func (s *RealmGift) SetPrice(val OptNilFloat64) {
	s.Price = val
}

// SetSlug sets the value of Slug.
func (s *RealmGift) SetSlug(val OptNilString) {
	s.Slug = val
}

// SetTitle sets the value of Title.
func (s *RealmGift) SetTitle(val OptNilString) {
	s.Title = val
}

// Ref: #/components/schemas/realm_GiftingAbility
type RealmGiftingAbility struct {
	CanReceive OptNilBool  `json:"can_receive"`
	CanSend    OptNilBool  `json:"can_send"`
	Enabled    OptNilBool  `json:"enabled"`
	UserID     OptNilInt64 `json:"user_id"`
}

// GetCanReceive returns the value of CanReceive.
func (s *RealmGiftingAbility) GetCanReceive() OptNilBool {
	return s.CanReceive
}

// GetCanSend returns the value of CanSend.
func (s *RealmGiftingAbility) GetCanSend() OptNilBool {
	return s.CanSend
}

// GetEnabled returns the value of Enabled.
func (s *RealmGiftingAbility) GetEnabled() OptNilBool {
	return s.Enabled
}

// GetUserID returns the value of UserID.
func (s *RealmGiftingAbility) GetUserID() OptNilInt64 {
	return s.UserID
}

// SetCanReceive sets the value of CanReceive.
func (s *RealmGiftingAbility) SetCanReceive(val OptNilBool) {
	s.CanReceive = val
}

// SetCanSend sets the value of CanSend.
func (s *RealmGiftingAbility) SetCanSend(val OptNilBool) {
	s.CanSend = val
}

// SetEnabled sets the value of Enabled.
func (s *RealmGiftingAbility) SetEnabled(val OptNilBool) {
	s.Enabled = val
}

// SetUserID sets the value of UserID.
func (s *RealmGiftingAbility) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// Ref: #/components/schemas/realm_Message
type RealmMessage struct {
	Attachment          OptNilString           `json:"attachment"`
	AttachmentAndroid   OptNilString           `json:"attachment_android"`
	AttachmentReadCount OptNilInt32            `json:"attachment_read_count"`
	AttachmentThumbnail OptNilString           `json:"attachment_thumbnail"`
	ConferenceCall      OptRealmConferenceCall `json:"conference_call"`
	CreatedAt           OptNilInt64            `json:"created_at"`
	FontSize            OptNilInt32            `json:"font_size"`
	GIF                 OptGifImage            `json:"gif"`
	ID                  OptNilInt64            `json:"id"`
	Invitation          OptChatInvitation      `json:"invitation"`
	IsError             OptNilBool             `json:"is_error"`
	IsSent              OptNilBool             `json:"is_sent"`
	MessageType         OptMessageType         `json:"message_type"`
	Parent              OptParentMessage       `json:"parent"`
	RefreshRetryCount   OptNilInt32            `json:"refresh_retry_count"`
	RoomID              OptNilInt64            `json:"room_id"`
	Sticker             OptSticker             `json:"sticker"`
	Text                OptNilString           `json:"text"`
	UserID              OptNilInt64            `json:"user_id"`
	VideoProcessed      OptNilBool             `json:"video_processed"`
	VideoThumbnailURL   OptNilString           `json:"video_thumbnail_url"`
	VideoURL            OptNilString           `json:"video_url"`
}

// GetAttachment returns the value of Attachment.
func (s *RealmMessage) GetAttachment() OptNilString {
	return s.Attachment
}

// GetAttachmentAndroid returns the value of AttachmentAndroid.
func (s *RealmMessage) GetAttachmentAndroid() OptNilString {
	return s.AttachmentAndroid
}

// GetAttachmentReadCount returns the value of AttachmentReadCount.
func (s *RealmMessage) GetAttachmentReadCount() OptNilInt32 {
	return s.AttachmentReadCount
}

// GetAttachmentThumbnail returns the value of AttachmentThumbnail.
func (s *RealmMessage) GetAttachmentThumbnail() OptNilString {
	return s.AttachmentThumbnail
}

// GetConferenceCall returns the value of ConferenceCall.
func (s *RealmMessage) GetConferenceCall() OptRealmConferenceCall {
	return s.ConferenceCall
}

// GetCreatedAt returns the value of CreatedAt.
func (s *RealmMessage) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetFontSize returns the value of FontSize.
func (s *RealmMessage) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGIF returns the value of GIF.
func (s *RealmMessage) GetGIF() OptGifImage {
	return s.GIF
}

// GetID returns the value of ID.
func (s *RealmMessage) GetID() OptNilInt64 {
	return s.ID
}

// GetInvitation returns the value of Invitation.
func (s *RealmMessage) GetInvitation() OptChatInvitation {
	return s.Invitation
}

// GetIsError returns the value of IsError.
func (s *RealmMessage) GetIsError() OptNilBool {
	return s.IsError
}

// GetIsSent returns the value of IsSent.
func (s *RealmMessage) GetIsSent() OptNilBool {
	return s.IsSent
}

// GetMessageType returns the value of MessageType.
func (s *RealmMessage) GetMessageType() OptMessageType {
	return s.MessageType
}

// GetParent returns the value of Parent.
func (s *RealmMessage) GetParent() OptParentMessage {
	return s.Parent
}

// GetRefreshRetryCount returns the value of RefreshRetryCount.
func (s *RealmMessage) GetRefreshRetryCount() OptNilInt32 {
	return s.RefreshRetryCount
}

// GetRoomID returns the value of RoomID.
func (s *RealmMessage) GetRoomID() OptNilInt64 {
	return s.RoomID
}

// GetSticker returns the value of Sticker.
func (s *RealmMessage) GetSticker() OptSticker {
	return s.Sticker
}

// GetText returns the value of Text.
func (s *RealmMessage) GetText() OptNilString {
	return s.Text
}

// GetUserID returns the value of UserID.
func (s *RealmMessage) GetUserID() OptNilInt64 {
	return s.UserID
}

// GetVideoProcessed returns the value of VideoProcessed.
func (s *RealmMessage) GetVideoProcessed() OptNilBool {
	return s.VideoProcessed
}

// GetVideoThumbnailURL returns the value of VideoThumbnailURL.
func (s *RealmMessage) GetVideoThumbnailURL() OptNilString {
	return s.VideoThumbnailURL
}

// GetVideoURL returns the value of VideoURL.
func (s *RealmMessage) GetVideoURL() OptNilString {
	return s.VideoURL
}

// SetAttachment sets the value of Attachment.
func (s *RealmMessage) SetAttachment(val OptNilString) {
	s.Attachment = val
}

// SetAttachmentAndroid sets the value of AttachmentAndroid.
func (s *RealmMessage) SetAttachmentAndroid(val OptNilString) {
	s.AttachmentAndroid = val
}

// SetAttachmentReadCount sets the value of AttachmentReadCount.
func (s *RealmMessage) SetAttachmentReadCount(val OptNilInt32) {
	s.AttachmentReadCount = val
}

// SetAttachmentThumbnail sets the value of AttachmentThumbnail.
func (s *RealmMessage) SetAttachmentThumbnail(val OptNilString) {
	s.AttachmentThumbnail = val
}

// SetConferenceCall sets the value of ConferenceCall.
func (s *RealmMessage) SetConferenceCall(val OptRealmConferenceCall) {
	s.ConferenceCall = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *RealmMessage) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetFontSize sets the value of FontSize.
func (s *RealmMessage) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGIF sets the value of GIF.
func (s *RealmMessage) SetGIF(val OptGifImage) {
	s.GIF = val
}

// SetID sets the value of ID.
func (s *RealmMessage) SetID(val OptNilInt64) {
	s.ID = val
}

// SetInvitation sets the value of Invitation.
func (s *RealmMessage) SetInvitation(val OptChatInvitation) {
	s.Invitation = val
}

// SetIsError sets the value of IsError.
func (s *RealmMessage) SetIsError(val OptNilBool) {
	s.IsError = val
}

// SetIsSent sets the value of IsSent.
func (s *RealmMessage) SetIsSent(val OptNilBool) {
	s.IsSent = val
}

// SetMessageType sets the value of MessageType.
func (s *RealmMessage) SetMessageType(val OptMessageType) {
	s.MessageType = val
}

// SetParent sets the value of Parent.
func (s *RealmMessage) SetParent(val OptParentMessage) {
	s.Parent = val
}

// SetRefreshRetryCount sets the value of RefreshRetryCount.
func (s *RealmMessage) SetRefreshRetryCount(val OptNilInt32) {
	s.RefreshRetryCount = val
}

// SetRoomID sets the value of RoomID.
func (s *RealmMessage) SetRoomID(val OptNilInt64) {
	s.RoomID = val
}

// SetSticker sets the value of Sticker.
func (s *RealmMessage) SetSticker(val OptSticker) {
	s.Sticker = val
}

// SetText sets the value of Text.
func (s *RealmMessage) SetText(val OptNilString) {
	s.Text = val
}

// SetUserID sets the value of UserID.
func (s *RealmMessage) SetUserID(val OptNilInt64) {
	s.UserID = val
}

// SetVideoProcessed sets the value of VideoProcessed.
func (s *RealmMessage) SetVideoProcessed(val OptNilBool) {
	s.VideoProcessed = val
}

// SetVideoThumbnailURL sets the value of VideoThumbnailURL.
func (s *RealmMessage) SetVideoThumbnailURL(val OptNilString) {
	s.VideoThumbnailURL = val
}

// SetVideoURL sets the value of VideoURL.
func (s *RealmMessage) SetVideoURL(val OptNilString) {
	s.VideoURL = val
}

// Ref: #/components/schemas/realm_PlatformDetails
type RealmPlatformDetails struct {
	AffiliateURL OptNilString `json:"affiliate_url"`
	PackageID    OptNilString `json:"package_id"`
}

// GetAffiliateURL returns the value of AffiliateURL.
func (s *RealmPlatformDetails) GetAffiliateURL() OptNilString {
	return s.AffiliateURL
}

// GetPackageID returns the value of PackageID.
func (s *RealmPlatformDetails) GetPackageID() OptNilString {
	return s.PackageID
}

// SetAffiliateURL sets the value of AffiliateURL.
func (s *RealmPlatformDetails) SetAffiliateURL(val OptNilString) {
	s.AffiliateURL = val
}

// SetPackageID sets the value of PackageID.
func (s *RealmPlatformDetails) SetPackageID(val OptNilString) {
	s.PackageID = val
}

// Ref: #/components/schemas/realm_User
type RealmUser struct {
	AgeVerified                       OptNilBool             `json:"age_verified"`
	AppleConnected                    OptNilBool             `json:"apple_connected"`
	Biography                         OptNilString           `json:"biography"`
	BirthDate                         OptNilString           `json:"birth_date"`
	BlockingLimit                     OptNilInt32            `json:"blocking_limit"`
	ChatRequest                       OptNilBool             `json:"chat_request"`
	ConnectedBy                       OptNilStringArray      `json:"connected_by"`
	ContactPhones                     OptNilStringArray      `json:"contact_phones"`
	CountryCode                       OptNilString           `json:"country_code"`
	CoverImage                        OptNilString           `json:"cover_image"`
	CoverImageThumbnail               OptNilString           `json:"cover_image_thumbnail"`
	CreatedAtSeconds                  OptNilInt64            `json:"created_at_seconds"`
	DangerousUser                     OptNilBool             `json:"dangerous_user"`
	EmailConfirmed                    OptNilBool             `json:"email_confirmed"`
	FacebookConnected                 OptNilBool             `json:"facebook_connected"`
	FollowPending                     OptNilBool             `json:"follow_pending"`
	FollowedBy                        OptNilBool             `json:"followed_by"`
	FollowedByPending                 OptNilBool             `json:"followed_by_pending"`
	FollowersCount                    OptNilInt32            `json:"followers_count"`
	Following                         OptNilBool             `json:"following"`
	FollowingsCount                   OptNilInt32            `json:"followings_count"`
	Frame                             OptNilString           `json:"frame"`
	FrameThumbnail                    OptNilString           `json:"frame_thumbnail"`
	FromDifferentGenerationAndTrusted OptNilBool             `json:"from_different_generation_and_trusted"`
	Gender                            OptNilInt32            `json:"gender"`
	Generation                        OptNilInt32            `json:"generation"`
	GiftingAbility                    OptRealmGiftingAbility `json:"gifting_ability"`
	GoogleConnected                   OptNilBool             `json:"google_connected"`
	GroupPhoneOn                      OptNilBool             `json:"group_phone_on"`
	GroupUser                         OptGroupUser           `json:"group_user"`
	GroupVideoOn                      OptNilBool             `json:"group_video_on"`
	GroupsUsersCount                  OptNilInt32            `json:"groups_users_count"`
	Hidden                            OptNilBool             `json:"hidden"`
	HideVip                           OptNilBool             `json:"hide_vip"`
	ID                                OptNilInt64            `json:"id"`
	InterestsCount                    OptNilInt32            `json:"interests_count"`
	InterestsSelected                 OptNilBool             `json:"interests_selected"`
	IsPrivate                         OptNilBool             `json:"is_private"`
	LastLoggedInAtSeconds             OptNilInt64            `json:"last_logged_in_at_seconds"`
	LineConnected                     OptNilBool             `json:"line_connected"`
	LoginStreakCount                  OptNilInt32            `json:"login_streak_count"`
	MaskedEmail                       OptNilString           `json:"masked_email"`
	MediaCount                        OptNilInt32            `json:"media_count"`
	MobileNumber                      OptNilString           `json:"mobile_number"`
	NewUser                           OptNilBool             `json:"new_user"`
	Nickname                          OptNilString           `json:"nickname"`
	OnlineStatus                      OptOnlineStatusEnum    `json:"online_status"`
	PhoneOn                           OptNilBool             `json:"phone_on"`
	PostsCount                        OptNilInt32            `json:"posts_count"`
	Prefecture                        OptNilString           `json:"prefecture"`
	ProfileIcon                       OptNilString           `json:"profile_icon"`
	ProfileIconThumbnail              OptNilString           `json:"profile_icon_thumbnail"`
	PushNotification                  OptNilBool             `json:"push_notification"`
	RecentlyKenta                     OptNilBool             `json:"recently_kenta"`
	RestrictedReviewBy                OptNilString           `json:"restricted_review_by"`
	ReviewsCount                      OptNilInt32            `json:"reviews_count"`
	Title                             OptNilString           `json:"title"`
	TotalGiftsReceivedCount           OptNilInt32            `json:"total_gifts_received_count"`
	TwitterID                         OptNilString           `json:"twitter_id"`
	TwoFaEnabled                      OptNilBool             `json:"two_fa_enabled"`
	UpdatedTimeMillis                 OptNilInt64            `json:"updated_time_millis"`
	Username                          OptNilString           `json:"username"`
	UUID                              OptNilString           `json:"uuid"`
	VideoOn                           OptNilBool             `json:"video_on"`
	Vip                               OptNilBool             `json:"vip"`
	VipUntilSeconds                   OptNilInt64            `json:"vip_until_seconds"`
	WorldIDConnected                  OptNilBool             `json:"world_id_connected"`
}

// GetAgeVerified returns the value of AgeVerified.
func (s *RealmUser) GetAgeVerified() OptNilBool {
	return s.AgeVerified
}

// GetAppleConnected returns the value of AppleConnected.
func (s *RealmUser) GetAppleConnected() OptNilBool {
	return s.AppleConnected
}

// GetBiography returns the value of Biography.
func (s *RealmUser) GetBiography() OptNilString {
	return s.Biography
}

// GetBirthDate returns the value of BirthDate.
func (s *RealmUser) GetBirthDate() OptNilString {
	return s.BirthDate
}

// GetBlockingLimit returns the value of BlockingLimit.
func (s *RealmUser) GetBlockingLimit() OptNilInt32 {
	return s.BlockingLimit
}

// GetChatRequest returns the value of ChatRequest.
func (s *RealmUser) GetChatRequest() OptNilBool {
	return s.ChatRequest
}

// GetConnectedBy returns the value of ConnectedBy.
func (s *RealmUser) GetConnectedBy() OptNilStringArray {
	return s.ConnectedBy
}

// GetContactPhones returns the value of ContactPhones.
func (s *RealmUser) GetContactPhones() OptNilStringArray {
	return s.ContactPhones
}

// GetCountryCode returns the value of CountryCode.
func (s *RealmUser) GetCountryCode() OptNilString {
	return s.CountryCode
}

// GetCoverImage returns the value of CoverImage.
func (s *RealmUser) GetCoverImage() OptNilString {
	return s.CoverImage
}

// GetCoverImageThumbnail returns the value of CoverImageThumbnail.
func (s *RealmUser) GetCoverImageThumbnail() OptNilString {
	return s.CoverImageThumbnail
}

// GetCreatedAtSeconds returns the value of CreatedAtSeconds.
func (s *RealmUser) GetCreatedAtSeconds() OptNilInt64 {
	return s.CreatedAtSeconds
}

// GetDangerousUser returns the value of DangerousUser.
func (s *RealmUser) GetDangerousUser() OptNilBool {
	return s.DangerousUser
}

// GetEmailConfirmed returns the value of EmailConfirmed.
func (s *RealmUser) GetEmailConfirmed() OptNilBool {
	return s.EmailConfirmed
}

// GetFacebookConnected returns the value of FacebookConnected.
func (s *RealmUser) GetFacebookConnected() OptNilBool {
	return s.FacebookConnected
}

// GetFollowPending returns the value of FollowPending.
func (s *RealmUser) GetFollowPending() OptNilBool {
	return s.FollowPending
}

// GetFollowedBy returns the value of FollowedBy.
func (s *RealmUser) GetFollowedBy() OptNilBool {
	return s.FollowedBy
}

// GetFollowedByPending returns the value of FollowedByPending.
func (s *RealmUser) GetFollowedByPending() OptNilBool {
	return s.FollowedByPending
}

// GetFollowersCount returns the value of FollowersCount.
func (s *RealmUser) GetFollowersCount() OptNilInt32 {
	return s.FollowersCount
}

// GetFollowing returns the value of Following.
func (s *RealmUser) GetFollowing() OptNilBool {
	return s.Following
}

// GetFollowingsCount returns the value of FollowingsCount.
func (s *RealmUser) GetFollowingsCount() OptNilInt32 {
	return s.FollowingsCount
}

// GetFrame returns the value of Frame.
func (s *RealmUser) GetFrame() OptNilString {
	return s.Frame
}

// GetFrameThumbnail returns the value of FrameThumbnail.
func (s *RealmUser) GetFrameThumbnail() OptNilString {
	return s.FrameThumbnail
}

// GetFromDifferentGenerationAndTrusted returns the value of FromDifferentGenerationAndTrusted.
func (s *RealmUser) GetFromDifferentGenerationAndTrusted() OptNilBool {
	return s.FromDifferentGenerationAndTrusted
}

// GetGender returns the value of Gender.
func (s *RealmUser) GetGender() OptNilInt32 {
	return s.Gender
}

// GetGeneration returns the value of Generation.
func (s *RealmUser) GetGeneration() OptNilInt32 {
	return s.Generation
}

// GetGiftingAbility returns the value of GiftingAbility.
func (s *RealmUser) GetGiftingAbility() OptRealmGiftingAbility {
	return s.GiftingAbility
}

// GetGoogleConnected returns the value of GoogleConnected.
func (s *RealmUser) GetGoogleConnected() OptNilBool {
	return s.GoogleConnected
}

// GetGroupPhoneOn returns the value of GroupPhoneOn.
func (s *RealmUser) GetGroupPhoneOn() OptNilBool {
	return s.GroupPhoneOn
}

// GetGroupUser returns the value of GroupUser.
func (s *RealmUser) GetGroupUser() OptGroupUser {
	return s.GroupUser
}

// GetGroupVideoOn returns the value of GroupVideoOn.
func (s *RealmUser) GetGroupVideoOn() OptNilBool {
	return s.GroupVideoOn
}

// GetGroupsUsersCount returns the value of GroupsUsersCount.
func (s *RealmUser) GetGroupsUsersCount() OptNilInt32 {
	return s.GroupsUsersCount
}

// GetHidden returns the value of Hidden.
func (s *RealmUser) GetHidden() OptNilBool {
	return s.Hidden
}

// GetHideVip returns the value of HideVip.
func (s *RealmUser) GetHideVip() OptNilBool {
	return s.HideVip
}

// GetID returns the value of ID.
func (s *RealmUser) GetID() OptNilInt64 {
	return s.ID
}

// GetInterestsCount returns the value of InterestsCount.
func (s *RealmUser) GetInterestsCount() OptNilInt32 {
	return s.InterestsCount
}

// GetInterestsSelected returns the value of InterestsSelected.
func (s *RealmUser) GetInterestsSelected() OptNilBool {
	return s.InterestsSelected
}

// GetIsPrivate returns the value of IsPrivate.
func (s *RealmUser) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetLastLoggedInAtSeconds returns the value of LastLoggedInAtSeconds.
func (s *RealmUser) GetLastLoggedInAtSeconds() OptNilInt64 {
	return s.LastLoggedInAtSeconds
}

// GetLineConnected returns the value of LineConnected.
func (s *RealmUser) GetLineConnected() OptNilBool {
	return s.LineConnected
}

// GetLoginStreakCount returns the value of LoginStreakCount.
func (s *RealmUser) GetLoginStreakCount() OptNilInt32 {
	return s.LoginStreakCount
}

// GetMaskedEmail returns the value of MaskedEmail.
func (s *RealmUser) GetMaskedEmail() OptNilString {
	return s.MaskedEmail
}

// GetMediaCount returns the value of MediaCount.
func (s *RealmUser) GetMediaCount() OptNilInt32 {
	return s.MediaCount
}

// GetMobileNumber returns the value of MobileNumber.
func (s *RealmUser) GetMobileNumber() OptNilString {
	return s.MobileNumber
}

// GetNewUser returns the value of NewUser.
func (s *RealmUser) GetNewUser() OptNilBool {
	return s.NewUser
}

// GetNickname returns the value of Nickname.
func (s *RealmUser) GetNickname() OptNilString {
	return s.Nickname
}

// GetOnlineStatus returns the value of OnlineStatus.
func (s *RealmUser) GetOnlineStatus() OptOnlineStatusEnum {
	return s.OnlineStatus
}

// GetPhoneOn returns the value of PhoneOn.
func (s *RealmUser) GetPhoneOn() OptNilBool {
	return s.PhoneOn
}

// GetPostsCount returns the value of PostsCount.
func (s *RealmUser) GetPostsCount() OptNilInt32 {
	return s.PostsCount
}

// GetPrefecture returns the value of Prefecture.
func (s *RealmUser) GetPrefecture() OptNilString {
	return s.Prefecture
}

// GetProfileIcon returns the value of ProfileIcon.
func (s *RealmUser) GetProfileIcon() OptNilString {
	return s.ProfileIcon
}

// GetProfileIconThumbnail returns the value of ProfileIconThumbnail.
func (s *RealmUser) GetProfileIconThumbnail() OptNilString {
	return s.ProfileIconThumbnail
}

// GetPushNotification returns the value of PushNotification.
func (s *RealmUser) GetPushNotification() OptNilBool {
	return s.PushNotification
}

// GetRecentlyKenta returns the value of RecentlyKenta.
func (s *RealmUser) GetRecentlyKenta() OptNilBool {
	return s.RecentlyKenta
}

// GetRestrictedReviewBy returns the value of RestrictedReviewBy.
func (s *RealmUser) GetRestrictedReviewBy() OptNilString {
	return s.RestrictedReviewBy
}

// GetReviewsCount returns the value of ReviewsCount.
func (s *RealmUser) GetReviewsCount() OptNilInt32 {
	return s.ReviewsCount
}

// GetTitle returns the value of Title.
func (s *RealmUser) GetTitle() OptNilString {
	return s.Title
}

// GetTotalGiftsReceivedCount returns the value of TotalGiftsReceivedCount.
func (s *RealmUser) GetTotalGiftsReceivedCount() OptNilInt32 {
	return s.TotalGiftsReceivedCount
}

// GetTwitterID returns the value of TwitterID.
func (s *RealmUser) GetTwitterID() OptNilString {
	return s.TwitterID
}

// GetTwoFaEnabled returns the value of TwoFaEnabled.
func (s *RealmUser) GetTwoFaEnabled() OptNilBool {
	return s.TwoFaEnabled
}

// GetUpdatedTimeMillis returns the value of UpdatedTimeMillis.
func (s *RealmUser) GetUpdatedTimeMillis() OptNilInt64 {
	return s.UpdatedTimeMillis
}

// GetUsername returns the value of Username.
func (s *RealmUser) GetUsername() OptNilString {
	return s.Username
}

// GetUUID returns the value of UUID.
func (s *RealmUser) GetUUID() OptNilString {
	return s.UUID
}

// GetVideoOn returns the value of VideoOn.
func (s *RealmUser) GetVideoOn() OptNilBool {
	return s.VideoOn
}

// GetVip returns the value of Vip.
func (s *RealmUser) GetVip() OptNilBool {
	return s.Vip
}

// GetVipUntilSeconds returns the value of VipUntilSeconds.
func (s *RealmUser) GetVipUntilSeconds() OptNilInt64 {
	return s.VipUntilSeconds
}

// GetWorldIDConnected returns the value of WorldIDConnected.
func (s *RealmUser) GetWorldIDConnected() OptNilBool {
	return s.WorldIDConnected
}

// SetAgeVerified sets the value of AgeVerified.
func (s *RealmUser) SetAgeVerified(val OptNilBool) {
	s.AgeVerified = val
}

// SetAppleConnected sets the value of AppleConnected.
func (s *RealmUser) SetAppleConnected(val OptNilBool) {
	s.AppleConnected = val
}

// SetBiography sets the value of Biography.
func (s *RealmUser) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetBirthDate sets the value of BirthDate.
func (s *RealmUser) SetBirthDate(val OptNilString) {
	s.BirthDate = val
}

// SetBlockingLimit sets the value of BlockingLimit.
func (s *RealmUser) SetBlockingLimit(val OptNilInt32) {
	s.BlockingLimit = val
}

// SetChatRequest sets the value of ChatRequest.
func (s *RealmUser) SetChatRequest(val OptNilBool) {
	s.ChatRequest = val
}

// SetConnectedBy sets the value of ConnectedBy.
func (s *RealmUser) SetConnectedBy(val OptNilStringArray) {
	s.ConnectedBy = val
}

// SetContactPhones sets the value of ContactPhones.
func (s *RealmUser) SetContactPhones(val OptNilStringArray) {
	s.ContactPhones = val
}

// SetCountryCode sets the value of CountryCode.
func (s *RealmUser) SetCountryCode(val OptNilString) {
	s.CountryCode = val
}

// SetCoverImage sets the value of CoverImage.
func (s *RealmUser) SetCoverImage(val OptNilString) {
	s.CoverImage = val
}

// SetCoverImageThumbnail sets the value of CoverImageThumbnail.
func (s *RealmUser) SetCoverImageThumbnail(val OptNilString) {
	s.CoverImageThumbnail = val
}

// SetCreatedAtSeconds sets the value of CreatedAtSeconds.
func (s *RealmUser) SetCreatedAtSeconds(val OptNilInt64) {
	s.CreatedAtSeconds = val
}

// SetDangerousUser sets the value of DangerousUser.
func (s *RealmUser) SetDangerousUser(val OptNilBool) {
	s.DangerousUser = val
}

// SetEmailConfirmed sets the value of EmailConfirmed.
func (s *RealmUser) SetEmailConfirmed(val OptNilBool) {
	s.EmailConfirmed = val
}

// SetFacebookConnected sets the value of FacebookConnected.
func (s *RealmUser) SetFacebookConnected(val OptNilBool) {
	s.FacebookConnected = val
}

// SetFollowPending sets the value of FollowPending.
func (s *RealmUser) SetFollowPending(val OptNilBool) {
	s.FollowPending = val
}

// SetFollowedBy sets the value of FollowedBy.
func (s *RealmUser) SetFollowedBy(val OptNilBool) {
	s.FollowedBy = val
}

// SetFollowedByPending sets the value of FollowedByPending.
func (s *RealmUser) SetFollowedByPending(val OptNilBool) {
	s.FollowedByPending = val
}

// SetFollowersCount sets the value of FollowersCount.
func (s *RealmUser) SetFollowersCount(val OptNilInt32) {
	s.FollowersCount = val
}

// SetFollowing sets the value of Following.
func (s *RealmUser) SetFollowing(val OptNilBool) {
	s.Following = val
}

// SetFollowingsCount sets the value of FollowingsCount.
func (s *RealmUser) SetFollowingsCount(val OptNilInt32) {
	s.FollowingsCount = val
}

// SetFrame sets the value of Frame.
func (s *RealmUser) SetFrame(val OptNilString) {
	s.Frame = val
}

// SetFrameThumbnail sets the value of FrameThumbnail.
func (s *RealmUser) SetFrameThumbnail(val OptNilString) {
	s.FrameThumbnail = val
}

// SetFromDifferentGenerationAndTrusted sets the value of FromDifferentGenerationAndTrusted.
func (s *RealmUser) SetFromDifferentGenerationAndTrusted(val OptNilBool) {
	s.FromDifferentGenerationAndTrusted = val
}

// SetGender sets the value of Gender.
func (s *RealmUser) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetGeneration sets the value of Generation.
func (s *RealmUser) SetGeneration(val OptNilInt32) {
	s.Generation = val
}

// SetGiftingAbility sets the value of GiftingAbility.
func (s *RealmUser) SetGiftingAbility(val OptRealmGiftingAbility) {
	s.GiftingAbility = val
}

// SetGoogleConnected sets the value of GoogleConnected.
func (s *RealmUser) SetGoogleConnected(val OptNilBool) {
	s.GoogleConnected = val
}

// SetGroupPhoneOn sets the value of GroupPhoneOn.
func (s *RealmUser) SetGroupPhoneOn(val OptNilBool) {
	s.GroupPhoneOn = val
}

// SetGroupUser sets the value of GroupUser.
func (s *RealmUser) SetGroupUser(val OptGroupUser) {
	s.GroupUser = val
}

// SetGroupVideoOn sets the value of GroupVideoOn.
func (s *RealmUser) SetGroupVideoOn(val OptNilBool) {
	s.GroupVideoOn = val
}

// SetGroupsUsersCount sets the value of GroupsUsersCount.
func (s *RealmUser) SetGroupsUsersCount(val OptNilInt32) {
	s.GroupsUsersCount = val
}

// SetHidden sets the value of Hidden.
func (s *RealmUser) SetHidden(val OptNilBool) {
	s.Hidden = val
}

// SetHideVip sets the value of HideVip.
func (s *RealmUser) SetHideVip(val OptNilBool) {
	s.HideVip = val
}

// SetID sets the value of ID.
func (s *RealmUser) SetID(val OptNilInt64) {
	s.ID = val
}

// SetInterestsCount sets the value of InterestsCount.
func (s *RealmUser) SetInterestsCount(val OptNilInt32) {
	s.InterestsCount = val
}

// SetInterestsSelected sets the value of InterestsSelected.
func (s *RealmUser) SetInterestsSelected(val OptNilBool) {
	s.InterestsSelected = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *RealmUser) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetLastLoggedInAtSeconds sets the value of LastLoggedInAtSeconds.
func (s *RealmUser) SetLastLoggedInAtSeconds(val OptNilInt64) {
	s.LastLoggedInAtSeconds = val
}

// SetLineConnected sets the value of LineConnected.
func (s *RealmUser) SetLineConnected(val OptNilBool) {
	s.LineConnected = val
}

// SetLoginStreakCount sets the value of LoginStreakCount.
func (s *RealmUser) SetLoginStreakCount(val OptNilInt32) {
	s.LoginStreakCount = val
}

// SetMaskedEmail sets the value of MaskedEmail.
func (s *RealmUser) SetMaskedEmail(val OptNilString) {
	s.MaskedEmail = val
}

// SetMediaCount sets the value of MediaCount.
func (s *RealmUser) SetMediaCount(val OptNilInt32) {
	s.MediaCount = val
}

// SetMobileNumber sets the value of MobileNumber.
func (s *RealmUser) SetMobileNumber(val OptNilString) {
	s.MobileNumber = val
}

// SetNewUser sets the value of NewUser.
func (s *RealmUser) SetNewUser(val OptNilBool) {
	s.NewUser = val
}

// SetNickname sets the value of Nickname.
func (s *RealmUser) SetNickname(val OptNilString) {
	s.Nickname = val
}

// SetOnlineStatus sets the value of OnlineStatus.
func (s *RealmUser) SetOnlineStatus(val OptOnlineStatusEnum) {
	s.OnlineStatus = val
}

// SetPhoneOn sets the value of PhoneOn.
func (s *RealmUser) SetPhoneOn(val OptNilBool) {
	s.PhoneOn = val
}

// SetPostsCount sets the value of PostsCount.
func (s *RealmUser) SetPostsCount(val OptNilInt32) {
	s.PostsCount = val
}

// SetPrefecture sets the value of Prefecture.
func (s *RealmUser) SetPrefecture(val OptNilString) {
	s.Prefecture = val
}

// SetProfileIcon sets the value of ProfileIcon.
func (s *RealmUser) SetProfileIcon(val OptNilString) {
	s.ProfileIcon = val
}

// SetProfileIconThumbnail sets the value of ProfileIconThumbnail.
func (s *RealmUser) SetProfileIconThumbnail(val OptNilString) {
	s.ProfileIconThumbnail = val
}

// SetPushNotification sets the value of PushNotification.
func (s *RealmUser) SetPushNotification(val OptNilBool) {
	s.PushNotification = val
}

// SetRecentlyKenta sets the value of RecentlyKenta.
func (s *RealmUser) SetRecentlyKenta(val OptNilBool) {
	s.RecentlyKenta = val
}

// SetRestrictedReviewBy sets the value of RestrictedReviewBy.
func (s *RealmUser) SetRestrictedReviewBy(val OptNilString) {
	s.RestrictedReviewBy = val
}

// SetReviewsCount sets the value of ReviewsCount.
func (s *RealmUser) SetReviewsCount(val OptNilInt32) {
	s.ReviewsCount = val
}

// SetTitle sets the value of Title.
func (s *RealmUser) SetTitle(val OptNilString) {
	s.Title = val
}

// SetTotalGiftsReceivedCount sets the value of TotalGiftsReceivedCount.
func (s *RealmUser) SetTotalGiftsReceivedCount(val OptNilInt32) {
	s.TotalGiftsReceivedCount = val
}

// SetTwitterID sets the value of TwitterID.
func (s *RealmUser) SetTwitterID(val OptNilString) {
	s.TwitterID = val
}

// SetTwoFaEnabled sets the value of TwoFaEnabled.
func (s *RealmUser) SetTwoFaEnabled(val OptNilBool) {
	s.TwoFaEnabled = val
}

// SetUpdatedTimeMillis sets the value of UpdatedTimeMillis.
func (s *RealmUser) SetUpdatedTimeMillis(val OptNilInt64) {
	s.UpdatedTimeMillis = val
}

// SetUsername sets the value of Username.
func (s *RealmUser) SetUsername(val OptNilString) {
	s.Username = val
}

// SetUUID sets the value of UUID.
func (s *RealmUser) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetVideoOn sets the value of VideoOn.
func (s *RealmUser) SetVideoOn(val OptNilBool) {
	s.VideoOn = val
}

// SetVip sets the value of Vip.
func (s *RealmUser) SetVip(val OptNilBool) {
	s.Vip = val
}

// SetVipUntilSeconds sets the value of VipUntilSeconds.
func (s *RealmUser) SetVipUntilSeconds(val OptNilInt64) {
	s.VipUntilSeconds = val
}

// SetWorldIDConnected sets the value of WorldIDConnected.
func (s *RealmUser) SetWorldIDConnected(val OptNilBool) {
	s.WorldIDConnected = val
}

// Ref: #/components/schemas/ReceivedGift
type ReceivedGift struct {
	Gift              OptGift         `json:"gift"`
	ReceivedCount     OptNilInt32     `json:"received_count"`
	Senders           OptNilUserArray `json:"senders"`
	TotalSendersCount OptNilInt32     `json:"total_senders_count"`
}

// GetGift returns the value of Gift.
func (s *ReceivedGift) GetGift() OptGift {
	return s.Gift
}

// GetReceivedCount returns the value of ReceivedCount.
func (s *ReceivedGift) GetReceivedCount() OptNilInt32 {
	return s.ReceivedCount
}

// GetSenders returns the value of Senders.
func (s *ReceivedGift) GetSenders() OptNilUserArray {
	return s.Senders
}

// GetTotalSendersCount returns the value of TotalSendersCount.
func (s *ReceivedGift) GetTotalSendersCount() OptNilInt32 {
	return s.TotalSendersCount
}

// SetGift sets the value of Gift.
func (s *ReceivedGift) SetGift(val OptGift) {
	s.Gift = val
}

// SetReceivedCount sets the value of ReceivedCount.
func (s *ReceivedGift) SetReceivedCount(val OptNilInt32) {
	s.ReceivedCount = val
}

// SetSenders sets the value of Senders.
func (s *ReceivedGift) SetSenders(val OptNilUserArray) {
	s.Senders = val
}

// SetTotalSendersCount sets the value of TotalSendersCount.
func (s *ReceivedGift) SetTotalSendersCount(val OptNilInt32) {
	s.TotalSendersCount = val
}

// Ref: #/components/schemas/RefreshCounterRequest
type RefreshCounterRequest struct {
	Counter         OptNilString `json:"counter"`
	LastRequestedAt OptNilInt64  `json:"last_requested_at"`
	Status          OptNilString `json:"status"`
}

// GetCounter returns the value of Counter.
func (s *RefreshCounterRequest) GetCounter() OptNilString {
	return s.Counter
}

// GetLastRequestedAt returns the value of LastRequestedAt.
func (s *RefreshCounterRequest) GetLastRequestedAt() OptNilInt64 {
	return s.LastRequestedAt
}

// GetStatus returns the value of Status.
func (s *RefreshCounterRequest) GetStatus() OptNilString {
	return s.Status
}

// SetCounter sets the value of Counter.
func (s *RefreshCounterRequest) SetCounter(val OptNilString) {
	s.Counter = val
}

// SetLastRequestedAt sets the value of LastRequestedAt.
func (s *RefreshCounterRequest) SetLastRequestedAt(val OptNilInt64) {
	s.LastRequestedAt = val
}

// SetStatus sets the value of Status.
func (s *RefreshCounterRequest) SetStatus(val OptNilString) {
	s.Status = val
}

// Ref: #/components/schemas/RefreshCounterRequestsResponse
type RefreshCounterRequestsResponse struct {
	ResetCounterRequests OptNilRefreshCounterRequestArray `json:"reset_counter_requests"`
}

// GetResetCounterRequests returns the value of ResetCounterRequests.
func (s *RefreshCounterRequestsResponse) GetResetCounterRequests() OptNilRefreshCounterRequestArray {
	return s.ResetCounterRequests
}

// SetResetCounterRequests sets the value of ResetCounterRequests.
func (s *RefreshCounterRequestsResponse) SetResetCounterRequests(val OptNilRefreshCounterRequestArray) {
	s.ResetCounterRequests = val
}

// RemoveChatRoomBackgroundNoContent is response for RemoveChatRoomBackground operation.
type RemoveChatRoomBackgroundNoContent struct{}

// RemoveCoverImageNoContent is response for RemoveCoverImage operation.
type RemoveCoverImageNoContent struct{}

// RemoveGroupCoverNoContent is response for RemoveGroupCover operation.
type RemoveGroupCoverNoContent struct{}

// RemoveGroupDeputiesNoContent is response for RemoveGroupDeputies operation.
type RemoveGroupDeputiesNoContent struct{}

// RemoveGroupIconNoContent is response for RemoveGroupIcon operation.
type RemoveGroupIconNoContent struct{}

// RemoveProfilePhotoNoContent is response for RemoveProfilePhoto operation.
type RemoveProfilePhotoNoContent struct{}

// RemoveRelatedGroupsNoContent is response for RemoveRelatedGroups operation.
type RemoveRelatedGroupsNoContent struct{}

// RemoveThreadMemberNoContent is response for RemoveThreadMember operation.
type RemoveThreadMemberNoContent struct{}

// ReplyToReviewNoContent is response for ReplyToReview operation.
type ReplyToReviewNoContent struct{}

type ReplyToReviewReq struct {
	APIKey     OptString    `json:"api_key"`
	Comment    OptString    `json:"comment"`
	Context    OptNilString `json:"context"`
	SignedInfo OptString    `json:"signed_info"`
	Timestamp  OptInt64     `json:"timestamp"`
	UUID       OptString    `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *ReplyToReviewReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetComment returns the value of Comment.
func (s *ReplyToReviewReq) GetComment() OptString {
	return s.Comment
}

// GetContext returns the value of Context.
func (s *ReplyToReviewReq) GetContext() OptNilString {
	return s.Context
}

// GetSignedInfo returns the value of SignedInfo.
func (s *ReplyToReviewReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *ReplyToReviewReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *ReplyToReviewReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *ReplyToReviewReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetComment sets the value of Comment.
func (s *ReplyToReviewReq) SetComment(val OptString) {
	s.Comment = val
}

// SetContext sets the value of Context.
func (s *ReplyToReviewReq) SetContext(val OptNilString) {
	s.Context = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *ReplyToReviewReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *ReplyToReviewReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *ReplyToReviewReq) SetUUID(val OptString) {
	s.UUID = val
}

// ReportChatRoomNoContent is response for ReportChatRoom operation.
type ReportChatRoomNoContent struct{}

type ReportChatRoomReq struct {
	CategoryID          OptInt64     `json:"category_id"`
	OpponentID          OptInt64     `json:"opponent_id"`
	Reason              OptNilString `json:"reason"`
	Screenshot2Filename OptNilString `json:"screenshot_2_filename"`
	Screenshot3Filename OptNilString `json:"screenshot_3_filename"`
	Screenshot4Filename OptNilString `json:"screenshot_4_filename"`
	ScreenshotFilename  OptNilString `json:"screenshot_filename"`
}

// GetCategoryID returns the value of CategoryID.
func (s *ReportChatRoomReq) GetCategoryID() OptInt64 {
	return s.CategoryID
}

// GetOpponentID returns the value of OpponentID.
func (s *ReportChatRoomReq) GetOpponentID() OptInt64 {
	return s.OpponentID
}

// GetReason returns the value of Reason.
func (s *ReportChatRoomReq) GetReason() OptNilString {
	return s.Reason
}

// GetScreenshot2Filename returns the value of Screenshot2Filename.
func (s *ReportChatRoomReq) GetScreenshot2Filename() OptNilString {
	return s.Screenshot2Filename
}

// GetScreenshot3Filename returns the value of Screenshot3Filename.
func (s *ReportChatRoomReq) GetScreenshot3Filename() OptNilString {
	return s.Screenshot3Filename
}

// GetScreenshot4Filename returns the value of Screenshot4Filename.
func (s *ReportChatRoomReq) GetScreenshot4Filename() OptNilString {
	return s.Screenshot4Filename
}

// GetScreenshotFilename returns the value of ScreenshotFilename.
func (s *ReportChatRoomReq) GetScreenshotFilename() OptNilString {
	return s.ScreenshotFilename
}

// SetCategoryID sets the value of CategoryID.
func (s *ReportChatRoomReq) SetCategoryID(val OptInt64) {
	s.CategoryID = val
}

// SetOpponentID sets the value of OpponentID.
func (s *ReportChatRoomReq) SetOpponentID(val OptInt64) {
	s.OpponentID = val
}

// SetReason sets the value of Reason.
func (s *ReportChatRoomReq) SetReason(val OptNilString) {
	s.Reason = val
}

// SetScreenshot2Filename sets the value of Screenshot2Filename.
func (s *ReportChatRoomReq) SetScreenshot2Filename(val OptNilString) {
	s.Screenshot2Filename = val
}

// SetScreenshot3Filename sets the value of Screenshot3Filename.
func (s *ReportChatRoomReq) SetScreenshot3Filename(val OptNilString) {
	s.Screenshot3Filename = val
}

// SetScreenshot4Filename sets the value of Screenshot4Filename.
func (s *ReportChatRoomReq) SetScreenshot4Filename(val OptNilString) {
	s.Screenshot4Filename = val
}

// SetScreenshotFilename sets the value of ScreenshotFilename.
func (s *ReportChatRoomReq) SetScreenshotFilename(val OptNilString) {
	s.ScreenshotFilename = val
}

// ReportGroupNoContent is response for ReportGroup operation.
type ReportGroupNoContent struct{}

type ReportGroupReq struct {
	CategoryID          OptInt64     `json:"category_id"`
	OpponentID          OptInt64     `json:"opponent_id"`
	Reason              OptNilString `json:"reason"`
	Screenshot2Filename OptNilString `json:"screenshot_2_filename"`
	Screenshot3Filename OptNilString `json:"screenshot_3_filename"`
	Screenshot4Filename OptNilString `json:"screenshot_4_filename"`
	ScreenshotFilename  OptNilString `json:"screenshot_filename"`
}

// GetCategoryID returns the value of CategoryID.
func (s *ReportGroupReq) GetCategoryID() OptInt64 {
	return s.CategoryID
}

// GetOpponentID returns the value of OpponentID.
func (s *ReportGroupReq) GetOpponentID() OptInt64 {
	return s.OpponentID
}

// GetReason returns the value of Reason.
func (s *ReportGroupReq) GetReason() OptNilString {
	return s.Reason
}

// GetScreenshot2Filename returns the value of Screenshot2Filename.
func (s *ReportGroupReq) GetScreenshot2Filename() OptNilString {
	return s.Screenshot2Filename
}

// GetScreenshot3Filename returns the value of Screenshot3Filename.
func (s *ReportGroupReq) GetScreenshot3Filename() OptNilString {
	return s.Screenshot3Filename
}

// GetScreenshot4Filename returns the value of Screenshot4Filename.
func (s *ReportGroupReq) GetScreenshot4Filename() OptNilString {
	return s.Screenshot4Filename
}

// GetScreenshotFilename returns the value of ScreenshotFilename.
func (s *ReportGroupReq) GetScreenshotFilename() OptNilString {
	return s.ScreenshotFilename
}

// SetCategoryID sets the value of CategoryID.
func (s *ReportGroupReq) SetCategoryID(val OptInt64) {
	s.CategoryID = val
}

// SetOpponentID sets the value of OpponentID.
func (s *ReportGroupReq) SetOpponentID(val OptInt64) {
	s.OpponentID = val
}

// SetReason sets the value of Reason.
func (s *ReportGroupReq) SetReason(val OptNilString) {
	s.Reason = val
}

// SetScreenshot2Filename sets the value of Screenshot2Filename.
func (s *ReportGroupReq) SetScreenshot2Filename(val OptNilString) {
	s.Screenshot2Filename = val
}

// SetScreenshot3Filename sets the value of Screenshot3Filename.
func (s *ReportGroupReq) SetScreenshot3Filename(val OptNilString) {
	s.Screenshot3Filename = val
}

// SetScreenshot4Filename sets the value of Screenshot4Filename.
func (s *ReportGroupReq) SetScreenshot4Filename(val OptNilString) {
	s.Screenshot4Filename = val
}

// SetScreenshotFilename sets the value of ScreenshotFilename.
func (s *ReportGroupReq) SetScreenshotFilename(val OptNilString) {
	s.ScreenshotFilename = val
}

// ReportPostNoContent is response for ReportPost operation.
type ReportPostNoContent struct{}

type ReportPostReq struct {
	CategoryID          OptInt64     `json:"category_id"`
	OpponentID          OptInt64     `json:"opponent_id"`
	Reason              OptNilString `json:"reason"`
	Screenshot2Filename OptNilString `json:"screenshot_2_filename"`
	Screenshot3Filename OptNilString `json:"screenshot_3_filename"`
	Screenshot4Filename OptNilString `json:"screenshot_4_filename"`
	ScreenshotFilename  OptNilString `json:"screenshot_filename"`
}

// GetCategoryID returns the value of CategoryID.
func (s *ReportPostReq) GetCategoryID() OptInt64 {
	return s.CategoryID
}

// GetOpponentID returns the value of OpponentID.
func (s *ReportPostReq) GetOpponentID() OptInt64 {
	return s.OpponentID
}

// GetReason returns the value of Reason.
func (s *ReportPostReq) GetReason() OptNilString {
	return s.Reason
}

// GetScreenshot2Filename returns the value of Screenshot2Filename.
func (s *ReportPostReq) GetScreenshot2Filename() OptNilString {
	return s.Screenshot2Filename
}

// GetScreenshot3Filename returns the value of Screenshot3Filename.
func (s *ReportPostReq) GetScreenshot3Filename() OptNilString {
	return s.Screenshot3Filename
}

// GetScreenshot4Filename returns the value of Screenshot4Filename.
func (s *ReportPostReq) GetScreenshot4Filename() OptNilString {
	return s.Screenshot4Filename
}

// GetScreenshotFilename returns the value of ScreenshotFilename.
func (s *ReportPostReq) GetScreenshotFilename() OptNilString {
	return s.ScreenshotFilename
}

// SetCategoryID sets the value of CategoryID.
func (s *ReportPostReq) SetCategoryID(val OptInt64) {
	s.CategoryID = val
}

// SetOpponentID sets the value of OpponentID.
func (s *ReportPostReq) SetOpponentID(val OptInt64) {
	s.OpponentID = val
}

// SetReason sets the value of Reason.
func (s *ReportPostReq) SetReason(val OptNilString) {
	s.Reason = val
}

// SetScreenshot2Filename sets the value of Screenshot2Filename.
func (s *ReportPostReq) SetScreenshot2Filename(val OptNilString) {
	s.Screenshot2Filename = val
}

// SetScreenshot3Filename sets the value of Screenshot3Filename.
func (s *ReportPostReq) SetScreenshot3Filename(val OptNilString) {
	s.Screenshot3Filename = val
}

// SetScreenshot4Filename sets the value of Screenshot4Filename.
func (s *ReportPostReq) SetScreenshot4Filename(val OptNilString) {
	s.Screenshot4Filename = val
}

// SetScreenshotFilename sets the value of ScreenshotFilename.
func (s *ReportPostReq) SetScreenshotFilename(val OptNilString) {
	s.ScreenshotFilename = val
}

// ReportUserNoContent is response for ReportUser operation.
type ReportUserNoContent struct{}

type ReportUserReq struct {
	CategoryID          OptInt64     `json:"category_id"`
	Reason              OptNilString `json:"reason"`
	Screenshot2Filename OptNilString `json:"screenshot_2_filename"`
	Screenshot3Filename OptNilString `json:"screenshot_3_filename"`
	Screenshot4Filename OptNilString `json:"screenshot_4_filename"`
	ScreenshotFilename  OptNilString `json:"screenshot_filename"`
}

// GetCategoryID returns the value of CategoryID.
func (s *ReportUserReq) GetCategoryID() OptInt64 {
	return s.CategoryID
}

// GetReason returns the value of Reason.
func (s *ReportUserReq) GetReason() OptNilString {
	return s.Reason
}

// GetScreenshot2Filename returns the value of Screenshot2Filename.
func (s *ReportUserReq) GetScreenshot2Filename() OptNilString {
	return s.Screenshot2Filename
}

// GetScreenshot3Filename returns the value of Screenshot3Filename.
func (s *ReportUserReq) GetScreenshot3Filename() OptNilString {
	return s.Screenshot3Filename
}

// GetScreenshot4Filename returns the value of Screenshot4Filename.
func (s *ReportUserReq) GetScreenshot4Filename() OptNilString {
	return s.Screenshot4Filename
}

// GetScreenshotFilename returns the value of ScreenshotFilename.
func (s *ReportUserReq) GetScreenshotFilename() OptNilString {
	return s.ScreenshotFilename
}

// SetCategoryID sets the value of CategoryID.
func (s *ReportUserReq) SetCategoryID(val OptInt64) {
	s.CategoryID = val
}

// SetReason sets the value of Reason.
func (s *ReportUserReq) SetReason(val OptNilString) {
	s.Reason = val
}

// SetScreenshot2Filename sets the value of Screenshot2Filename.
func (s *ReportUserReq) SetScreenshot2Filename(val OptNilString) {
	s.Screenshot2Filename = val
}

// SetScreenshot3Filename sets the value of Screenshot3Filename.
func (s *ReportUserReq) SetScreenshot3Filename(val OptNilString) {
	s.Screenshot3Filename = val
}

// SetScreenshot4Filename sets the value of Screenshot4Filename.
func (s *ReportUserReq) SetScreenshot4Filename(val OptNilString) {
	s.Screenshot4Filename = val
}

// SetScreenshotFilename sets the value of ScreenshotFilename.
func (s *ReportUserReq) SetScreenshotFilename(val OptNilString) {
	s.ScreenshotFilename = val
}

type RepostReq struct {
	Attachment2Filename OptNilString            `json:"attachment_2_filename"`
	Attachment3Filename OptNilString            `json:"attachment_3_filename"`
	Attachment4Filename OptNilString            `json:"attachment_4_filename"`
	Attachment5Filename OptNilString            `json:"attachment_5_filename"`
	Attachment6Filename OptNilString            `json:"attachment_6_filename"`
	Attachment7Filename OptNilString            `json:"attachment_7_filename"`
	Attachment8Filename OptNilString            `json:"attachment_8_filename"`
	Attachment9Filename OptNilString            `json:"attachment_9_filename"`
	AttachmentFilename  OptNilString            `json:"attachment_filename"`
	Choices             OptNilStringArray       `json:"choices[]"`
	Color               OptNilInt32             `json:"color"`
	FontSize            OptNilInt32             `json:"font_size"`
	GroupID             OptNilInt64             `json:"group_id"`
	InReplyTo           OptNilInt64             `json:"in_reply_to"`
	Language            OptNilString            `json:"language"`
	MentionIds          OptNilInt64Array        `json:"mention_ids[]"`
	MessageTags         OptRepostReqMessageTags `json:"message_tags"`
	PostID              OptInt64                `json:"post_id"`
	PostType            OptPostType             `json:"post_type"`
	SharedURL           OptRepostReqSharedURL   `json:"shared_url"`
	Text                OptNilString            `json:"text"`
	VideoFileName       OptNilString            `json:"video_file_name"`
}

// GetAttachment2Filename returns the value of Attachment2Filename.
func (s *RepostReq) GetAttachment2Filename() OptNilString {
	return s.Attachment2Filename
}

// GetAttachment3Filename returns the value of Attachment3Filename.
func (s *RepostReq) GetAttachment3Filename() OptNilString {
	return s.Attachment3Filename
}

// GetAttachment4Filename returns the value of Attachment4Filename.
func (s *RepostReq) GetAttachment4Filename() OptNilString {
	return s.Attachment4Filename
}

// GetAttachment5Filename returns the value of Attachment5Filename.
func (s *RepostReq) GetAttachment5Filename() OptNilString {
	return s.Attachment5Filename
}

// GetAttachment6Filename returns the value of Attachment6Filename.
func (s *RepostReq) GetAttachment6Filename() OptNilString {
	return s.Attachment6Filename
}

// GetAttachment7Filename returns the value of Attachment7Filename.
func (s *RepostReq) GetAttachment7Filename() OptNilString {
	return s.Attachment7Filename
}

// GetAttachment8Filename returns the value of Attachment8Filename.
func (s *RepostReq) GetAttachment8Filename() OptNilString {
	return s.Attachment8Filename
}

// GetAttachment9Filename returns the value of Attachment9Filename.
func (s *RepostReq) GetAttachment9Filename() OptNilString {
	return s.Attachment9Filename
}

// GetAttachmentFilename returns the value of AttachmentFilename.
func (s *RepostReq) GetAttachmentFilename() OptNilString {
	return s.AttachmentFilename
}

// GetChoices returns the value of Choices.
func (s *RepostReq) GetChoices() OptNilStringArray {
	return s.Choices
}

// GetColor returns the value of Color.
func (s *RepostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *RepostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGroupID returns the value of GroupID.
func (s *RepostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetInReplyTo returns the value of InReplyTo.
func (s *RepostReq) GetInReplyTo() OptNilInt64 {
	return s.InReplyTo
}

// GetLanguage returns the value of Language.
func (s *RepostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMentionIds returns the value of MentionIds.
func (s *RepostReq) GetMentionIds() OptNilInt64Array {
	return s.MentionIds
}

// GetMessageTags returns the value of MessageTags.
func (s *RepostReq) GetMessageTags() OptRepostReqMessageTags {
	return s.MessageTags
}

// GetPostID returns the value of PostID.
func (s *RepostReq) GetPostID() OptInt64 {
	return s.PostID
}

// GetPostType returns the value of PostType.
func (s *RepostReq) GetPostType() OptPostType {
	return s.PostType
}

// GetSharedURL returns the value of SharedURL.
func (s *RepostReq) GetSharedURL() OptRepostReqSharedURL {
	return s.SharedURL
}

// GetText returns the value of Text.
func (s *RepostReq) GetText() OptNilString {
	return s.Text
}

// GetVideoFileName returns the value of VideoFileName.
func (s *RepostReq) GetVideoFileName() OptNilString {
	return s.VideoFileName
}

// SetAttachment2Filename sets the value of Attachment2Filename.
func (s *RepostReq) SetAttachment2Filename(val OptNilString) {
	s.Attachment2Filename = val
}

// SetAttachment3Filename sets the value of Attachment3Filename.
func (s *RepostReq) SetAttachment3Filename(val OptNilString) {
	s.Attachment3Filename = val
}

// SetAttachment4Filename sets the value of Attachment4Filename.
func (s *RepostReq) SetAttachment4Filename(val OptNilString) {
	s.Attachment4Filename = val
}

// SetAttachment5Filename sets the value of Attachment5Filename.
func (s *RepostReq) SetAttachment5Filename(val OptNilString) {
	s.Attachment5Filename = val
}

// SetAttachment6Filename sets the value of Attachment6Filename.
func (s *RepostReq) SetAttachment6Filename(val OptNilString) {
	s.Attachment6Filename = val
}

// SetAttachment7Filename sets the value of Attachment7Filename.
func (s *RepostReq) SetAttachment7Filename(val OptNilString) {
	s.Attachment7Filename = val
}

// SetAttachment8Filename sets the value of Attachment8Filename.
func (s *RepostReq) SetAttachment8Filename(val OptNilString) {
	s.Attachment8Filename = val
}

// SetAttachment9Filename sets the value of Attachment9Filename.
func (s *RepostReq) SetAttachment9Filename(val OptNilString) {
	s.Attachment9Filename = val
}

// SetAttachmentFilename sets the value of AttachmentFilename.
func (s *RepostReq) SetAttachmentFilename(val OptNilString) {
	s.AttachmentFilename = val
}

// SetChoices sets the value of Choices.
func (s *RepostReq) SetChoices(val OptNilStringArray) {
	s.Choices = val
}

// SetColor sets the value of Color.
func (s *RepostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *RepostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGroupID sets the value of GroupID.
func (s *RepostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetInReplyTo sets the value of InReplyTo.
func (s *RepostReq) SetInReplyTo(val OptNilInt64) {
	s.InReplyTo = val
}

// SetLanguage sets the value of Language.
func (s *RepostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMentionIds sets the value of MentionIds.
func (s *RepostReq) SetMentionIds(val OptNilInt64Array) {
	s.MentionIds = val
}

// SetMessageTags sets the value of MessageTags.
func (s *RepostReq) SetMessageTags(val OptRepostReqMessageTags) {
	s.MessageTags = val
}

// SetPostID sets the value of PostID.
func (s *RepostReq) SetPostID(val OptInt64) {
	s.PostID = val
}

// SetPostType sets the value of PostType.
func (s *RepostReq) SetPostType(val OptPostType) {
	s.PostType = val
}

// SetSharedURL sets the value of SharedURL.
func (s *RepostReq) SetSharedURL(val OptRepostReqSharedURL) {
	s.SharedURL = val
}

// SetText sets the value of Text.
func (s *RepostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetVideoFileName sets the value of VideoFileName.
func (s *RepostReq) SetVideoFileName(val OptNilString) {
	s.VideoFileName = val
}

type RepostReqMessageTags struct{}

type RepostReqSharedURL struct{}

type RequestEmailVerificationReq struct {
	DeviceUUID OptString    `json:"device_uuid"`
	Email      OptString    `json:"email"`
	Intent     OptNilString `json:"intent"`
	Locale     OptString    `json:"locale"`
}

// GetDeviceUUID returns the value of DeviceUUID.
func (s *RequestEmailVerificationReq) GetDeviceUUID() OptString {
	return s.DeviceUUID
}

// GetEmail returns the value of Email.
func (s *RequestEmailVerificationReq) GetEmail() OptString {
	return s.Email
}

// GetIntent returns the value of Intent.
func (s *RequestEmailVerificationReq) GetIntent() OptNilString {
	return s.Intent
}

// GetLocale returns the value of Locale.
func (s *RequestEmailVerificationReq) GetLocale() OptString {
	return s.Locale
}

// SetDeviceUUID sets the value of DeviceUUID.
func (s *RequestEmailVerificationReq) SetDeviceUUID(val OptString) {
	s.DeviceUUID = val
}

// SetEmail sets the value of Email.
func (s *RequestEmailVerificationReq) SetEmail(val OptString) {
	s.Email = val
}

// SetIntent sets the value of Intent.
func (s *RequestEmailVerificationReq) SetIntent(val OptNilString) {
	s.Intent = val
}

// SetLocale sets the value of Locale.
func (s *RequestEmailVerificationReq) SetLocale(val OptString) {
	s.Locale = val
}

// RequestFollowNoContent is response for RequestFollow operation.
type RequestFollowNoContent struct{}

type RequestFollowReq struct {
	Action OptString `json:"action"`
}

// GetAction returns the value of Action.
func (s *RequestFollowReq) GetAction() OptString {
	return s.Action
}

// SetAction sets the value of Action.
func (s *RequestFollowReq) SetAction(val OptString) {
	s.Action = val
}

// RequestGroupWalkthroughNoContent is response for RequestGroupWalkthrough operation.
type RequestGroupWalkthroughNoContent struct{}

// ResendConfirmEmailNoContent is response for ResendConfirmEmail operation.
type ResendConfirmEmailNoContent struct{}

// ResetCountersNoContent is response for ResetCounters operation.
type ResetCountersNoContent struct{}

type ResetCountersReq struct {
	Counter OptString `json:"counter"`
}

// GetCounter returns the value of Counter.
func (s *ResetCountersReq) GetCounter() OptString {
	return s.Counter
}

// SetCounter sets the value of Counter.
func (s *ResetCountersReq) SetCounter(val OptString) {
	s.Counter = val
}

// ResetPasswordNoContent is response for ResetPassword operation.
type ResetPasswordNoContent struct{}

type ResetPasswordReq struct {
	Email           OptString `json:"email"`
	EmailGrantToken OptString `json:"email_grant_token"`
	Password        OptString `json:"password"`
}

// GetEmail returns the value of Email.
func (s *ResetPasswordReq) GetEmail() OptString {
	return s.Email
}

// GetEmailGrantToken returns the value of EmailGrantToken.
func (s *ResetPasswordReq) GetEmailGrantToken() OptString {
	return s.EmailGrantToken
}

// GetPassword returns the value of Password.
func (s *ResetPasswordReq) GetPassword() OptString {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *ResetPasswordReq) SetEmail(val OptString) {
	s.Email = val
}

// SetEmailGrantToken sets the value of EmailGrantToken.
func (s *ResetPasswordReq) SetEmailGrantToken(val OptString) {
	s.EmailGrantToken = val
}

// SetPassword sets the value of Password.
func (s *ResetPasswordReq) SetPassword(val OptString) {
	s.Password = val
}

// Ref: #/components/schemas/Review
type Review struct {
	Comment       OptNilString `json:"comment"`
	CreatedAt     OptNilInt64  `json:"created_at"`
	ID            OptNilInt64  `json:"id"`
	MutualReview  OptNilBool   `json:"mutual_review"`
	ReportedCount OptNilInt32  `json:"reported_count"`
	Reviewer      OptRealmUser `json:"reviewer"`
	User          OptRealmUser `json:"user"`
}

// GetComment returns the value of Comment.
func (s *Review) GetComment() OptNilString {
	return s.Comment
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Review) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetID returns the value of ID.
func (s *Review) GetID() OptNilInt64 {
	return s.ID
}

// GetMutualReview returns the value of MutualReview.
func (s *Review) GetMutualReview() OptNilBool {
	return s.MutualReview
}

// GetReportedCount returns the value of ReportedCount.
func (s *Review) GetReportedCount() OptNilInt32 {
	return s.ReportedCount
}

// GetReviewer returns the value of Reviewer.
func (s *Review) GetReviewer() OptRealmUser {
	return s.Reviewer
}

// GetUser returns the value of User.
func (s *Review) GetUser() OptRealmUser {
	return s.User
}

// SetComment sets the value of Comment.
func (s *Review) SetComment(val OptNilString) {
	s.Comment = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Review) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetID sets the value of ID.
func (s *Review) SetID(val OptNilInt64) {
	s.ID = val
}

// SetMutualReview sets the value of MutualReview.
func (s *Review) SetMutualReview(val OptNilBool) {
	s.MutualReview = val
}

// SetReportedCount sets the value of ReportedCount.
func (s *Review) SetReportedCount(val OptNilInt32) {
	s.ReportedCount = val
}

// SetReviewer sets the value of Reviewer.
func (s *Review) SetReviewer(val OptRealmUser) {
	s.Reviewer = val
}

// SetUser sets the value of User.
func (s *Review) SetUser(val OptRealmUser) {
	s.User = val
}

// Ref: #/components/schemas/ReviewRestriction
type ReviewRestriction struct {
	APIValue OptNilString `json:"api_value"`
}

// GetAPIValue returns the value of APIValue.
func (s *ReviewRestriction) GetAPIValue() OptNilString {
	return s.APIValue
}

// SetAPIValue sets the value of APIValue.
func (s *ReviewRestriction) SetAPIValue(val OptNilString) {
	s.APIValue = val
}

// Ref: #/components/schemas/ReviewsResponse
type ReviewsResponse struct {
	PinnedReviews OptNilReviewArray `json:"pinned_reviews"`
	Reviews       OptNilReviewArray `json:"reviews"`
}

// GetPinnedReviews returns the value of PinnedReviews.
func (s *ReviewsResponse) GetPinnedReviews() OptNilReviewArray {
	return s.PinnedReviews
}

// GetReviews returns the value of Reviews.
func (s *ReviewsResponse) GetReviews() OptNilReviewArray {
	return s.Reviews
}

// SetPinnedReviews sets the value of PinnedReviews.
func (s *ReviewsResponse) SetPinnedReviews(val OptNilReviewArray) {
	s.PinnedReviews = val
}

// SetReviews sets the value of Reviews.
func (s *ReviewsResponse) SetReviews(val OptNilReviewArray) {
	s.Reviews = val
}

// Ref: #/components/schemas/RtmTokenResponse
type RtmTokenResponse struct {
	Token OptNilString `json:"token"`
}

// GetToken returns the value of Token.
func (s *RtmTokenResponse) GetToken() OptNilString {
	return s.Token
}

// SetToken sets the value of Token.
func (s *RtmTokenResponse) SetToken(val OptNilString) {
	s.Token = val
}

// Ref: #/components/schemas/SearchCriteria
type SearchCriteria struct {
	Biography  OptNilString `json:"biography"`
	Gender     OptNilInt32  `json:"gender"`
	Nickname   OptNilString `json:"nickname"`
	Prefecture OptNilString `json:"prefecture"`
	Username   OptNilString `json:"username"`
}

// GetBiography returns the value of Biography.
func (s *SearchCriteria) GetBiography() OptNilString {
	return s.Biography
}

// GetGender returns the value of Gender.
func (s *SearchCriteria) GetGender() OptNilInt32 {
	return s.Gender
}

// GetNickname returns the value of Nickname.
func (s *SearchCriteria) GetNickname() OptNilString {
	return s.Nickname
}

// GetPrefecture returns the value of Prefecture.
func (s *SearchCriteria) GetPrefecture() OptNilString {
	return s.Prefecture
}

// GetUsername returns the value of Username.
func (s *SearchCriteria) GetUsername() OptNilString {
	return s.Username
}

// SetBiography sets the value of Biography.
func (s *SearchCriteria) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetGender sets the value of Gender.
func (s *SearchCriteria) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetNickname sets the value of Nickname.
func (s *SearchCriteria) SetNickname(val OptNilString) {
	s.Nickname = val
}

// SetPrefecture sets the value of Prefecture.
func (s *SearchCriteria) SetPrefecture(val OptNilString) {
	s.Prefecture = val
}

// SetUsername sets the value of Username.
func (s *SearchCriteria) SetUsername(val OptNilString) {
	s.Username = val
}

// Ref: #/components/schemas/SearchUsersRequest
type SearchUsersRequest struct {
	User OptSearchCriteria `json:"user"`
}

// GetUser returns the value of User.
func (s *SearchUsersRequest) GetUser() OptSearchCriteria {
	return s.User
}

// SetUser sets the value of User.
func (s *SearchUsersRequest) SetUser(val OptSearchCriteria) {
	s.User = val
}

type SendChatMessageReq struct {
	AttachmentFileName OptNilString   `json:"attachment_file_name"`
	CallType           OptNilString   `json:"call_type"`
	FontSize           OptNilInt32    `json:"font_size"`
	GIFImageID         OptNilInt64    `json:"gif_image_id"`
	MessageType        OptMessageType `json:"message_type"`
	ParentID           OptNilInt64    `json:"parent_id"`
	StickerID          OptNilInt64    `json:"sticker_id"`
	StickerPackID      OptNilInt64    `json:"sticker_pack_id"`
	Text               OptNilString   `json:"text"`
	VideoFileName      OptNilString   `json:"video_file_name"`
}

// GetAttachmentFileName returns the value of AttachmentFileName.
func (s *SendChatMessageReq) GetAttachmentFileName() OptNilString {
	return s.AttachmentFileName
}

// GetCallType returns the value of CallType.
func (s *SendChatMessageReq) GetCallType() OptNilString {
	return s.CallType
}

// GetFontSize returns the value of FontSize.
func (s *SendChatMessageReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetGIFImageID returns the value of GIFImageID.
func (s *SendChatMessageReq) GetGIFImageID() OptNilInt64 {
	return s.GIFImageID
}

// GetMessageType returns the value of MessageType.
func (s *SendChatMessageReq) GetMessageType() OptMessageType {
	return s.MessageType
}

// GetParentID returns the value of ParentID.
func (s *SendChatMessageReq) GetParentID() OptNilInt64 {
	return s.ParentID
}

// GetStickerID returns the value of StickerID.
func (s *SendChatMessageReq) GetStickerID() OptNilInt64 {
	return s.StickerID
}

// GetStickerPackID returns the value of StickerPackID.
func (s *SendChatMessageReq) GetStickerPackID() OptNilInt64 {
	return s.StickerPackID
}

// GetText returns the value of Text.
func (s *SendChatMessageReq) GetText() OptNilString {
	return s.Text
}

// GetVideoFileName returns the value of VideoFileName.
func (s *SendChatMessageReq) GetVideoFileName() OptNilString {
	return s.VideoFileName
}

// SetAttachmentFileName sets the value of AttachmentFileName.
func (s *SendChatMessageReq) SetAttachmentFileName(val OptNilString) {
	s.AttachmentFileName = val
}

// SetCallType sets the value of CallType.
func (s *SendChatMessageReq) SetCallType(val OptNilString) {
	s.CallType = val
}

// SetFontSize sets the value of FontSize.
func (s *SendChatMessageReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetGIFImageID sets the value of GIFImageID.
func (s *SendChatMessageReq) SetGIFImageID(val OptNilInt64) {
	s.GIFImageID = val
}

// SetMessageType sets the value of MessageType.
func (s *SendChatMessageReq) SetMessageType(val OptMessageType) {
	s.MessageType = val
}

// SetParentID sets the value of ParentID.
func (s *SendChatMessageReq) SetParentID(val OptNilInt64) {
	s.ParentID = val
}

// SetStickerID sets the value of StickerID.
func (s *SendChatMessageReq) SetStickerID(val OptNilInt64) {
	s.StickerID = val
}

// SetStickerPackID sets the value of StickerPackID.
func (s *SendChatMessageReq) SetStickerPackID(val OptNilInt64) {
	s.StickerPackID = val
}

// SetText sets the value of Text.
func (s *SendChatMessageReq) SetText(val OptNilString) {
	s.Text = val
}

// SetVideoFileName sets the value of VideoFileName.
func (s *SendChatMessageReq) SetVideoFileName(val OptNilString) {
	s.VideoFileName = val
}

// SetGroupTitleNoContent is response for SetGroupTitle operation.
type SetGroupTitleNoContent struct{}

type SetGroupTitleReq struct {
	Title OptString `json:"title"`
}

// GetTitle returns the value of Title.
func (s *SetGroupTitleReq) GetTitle() OptString {
	return s.Title
}

// SetTitle sets the value of Title.
func (s *SetGroupTitleReq) SetTitle(val OptString) {
	s.Title = val
}

// SetHimaNoContent is response for SetHima operation.
type SetHimaNoContent struct{}

// Ref: #/components/schemas/Setting
type Setting struct {
	NotificationGroupJoin          OptNilBool `json:"notification_group_join"`
	NotificationGroupMessageTagAll OptNilBool `json:"notification_group_message_tag_all"`
	NotificationGroupPost          OptNilBool `json:"notification_group_post"`
	NotificationGroupRequest       OptNilBool `json:"notification_group_request"`
}

// GetNotificationGroupJoin returns the value of NotificationGroupJoin.
func (s *Setting) GetNotificationGroupJoin() OptNilBool {
	return s.NotificationGroupJoin
}

// GetNotificationGroupMessageTagAll returns the value of NotificationGroupMessageTagAll.
func (s *Setting) GetNotificationGroupMessageTagAll() OptNilBool {
	return s.NotificationGroupMessageTagAll
}

// GetNotificationGroupPost returns the value of NotificationGroupPost.
func (s *Setting) GetNotificationGroupPost() OptNilBool {
	return s.NotificationGroupPost
}

// GetNotificationGroupRequest returns the value of NotificationGroupRequest.
func (s *Setting) GetNotificationGroupRequest() OptNilBool {
	return s.NotificationGroupRequest
}

// SetNotificationGroupJoin sets the value of NotificationGroupJoin.
func (s *Setting) SetNotificationGroupJoin(val OptNilBool) {
	s.NotificationGroupJoin = val
}

// SetNotificationGroupMessageTagAll sets the value of NotificationGroupMessageTagAll.
func (s *Setting) SetNotificationGroupMessageTagAll(val OptNilBool) {
	s.NotificationGroupMessageTagAll = val
}

// SetNotificationGroupPost sets the value of NotificationGroupPost.
func (s *Setting) SetNotificationGroupPost(val OptNilBool) {
	s.NotificationGroupPost = val
}

// SetNotificationGroupRequest sets the value of NotificationGroupRequest.
func (s *Setting) SetNotificationGroupRequest(val OptNilBool) {
	s.NotificationGroupRequest = val
}

// Ref: #/components/schemas/Settings
type Settings struct {
	AgeRestrictedOnReview               OptNilBool `json:"age_restricted_on_review"`
	AllowReposts                        OptNilBool `json:"allow_reposts"`
	CautionUserChat                     OptNilBool `json:"caution_user_chat"`
	FollowingOnlyCallInvite             OptNilBool `json:"following_only_call_invite"`
	FollowingOnlyGroupInvite            OptNilBool `json:"following_only_group_invite"`
	FollowingRestrictedOnReview         OptNilBool `json:"following_restricted_on_review"`
	HideActiveCall                      OptNilBool `json:"hide_active_call"`
	HideGiftsReceived                   OptNilBool `json:"hide_gifts_received"`
	HideHotPost                         OptNilBool `json:"hide_hot_post"`
	HideOnInvitable                     OptNilBool `json:"hide_on_invitable"`
	HideOnlineStatus                    OptNilBool `json:"hide_online_status"`
	HideReplyFollowingTimeline          OptNilBool `json:"hide_reply_following_timeline"`
	HideReplyPublicTimeline             OptNilBool `json:"hide_reply_public_timeline"`
	HideVip                             OptNilBool `json:"hide_vip"`
	InvisibleOnUserSearch               OptNilBool `json:"invisible_on_user_search"`
	NoReplyGroupTimeline                OptNilBool `json:"no_reply_group_timeline"`
	NotificationBirthdayToFollowers     OptNilBool `json:"notification_birthday_to_followers"`
	NotificationBulkCallInvite          OptNilBool `json:"notification_bulk_call_invite"`
	NotificationCallInvite              OptNilBool `json:"notification_call_invite"`
	NotificationChat                    OptNilBool `json:"notification_chat"`
	NotificationChatDelete              OptNilBool `json:"notification_chat_delete"`
	NotificationContactFriend           OptNilBool `json:"notification_contact_friend"`
	NotificationDailyQuest              OptNilBool `json:"notification_daily_quest"`
	NotificationDailySummary            OptNilBool `json:"notification_daily_summary"`
	NotificationFollow                  OptNilBool `json:"notification_follow"`
	NotificationFollowAccept            OptNilBool `json:"notification_follow_accept"`
	NotificationFollowRequest           OptNilBool `json:"notification_follow_request"`
	NotificationFollowerConferenceCall  OptNilBool `json:"notification_follower_conference_call"`
	NotificationFollowerCreateGroup     OptNilBool `json:"notification_follower_create_group"`
	NotificationFollowingBirthdateOn    OptNilBool `json:"notification_following_birthdate_on"`
	NotificationFollowingPostAfterBreak OptNilBool `json:"notification_following_post_after_break"`
	NotificationFollowingsInCall        OptNilBool `json:"notification_followings_in_call"`
	NotificationFootprint               OptNilBool `json:"notification_footprint"`
	NotificationGiftReceive             OptNilBool `json:"notification_gift_receive"`
	NotificationGroupAccept             OptNilBool `json:"notification_group_accept"`
	NotificationGroupConferenceCall     OptNilBool `json:"notification_group_conference_call"`
	NotificationGroupInvite             OptNilBool `json:"notification_group_invite"`
	NotificationGroupJoin               OptNilBool `json:"notification_group_join"`
	NotificationGroupMessageTagAll      OptNilBool `json:"notification_group_message_tag_all"`
	NotificationGroupModerator          OptNilBool `json:"notification_group_moderator"`
	NotificationGroupPost               OptNilBool `json:"notification_group_post"`
	NotificationGroupRequest            OptNilBool `json:"notification_group_request"`
	NotificationHimaNow                 OptNilBool `json:"notification_hima_now"`
	NotificationInviteeCreateAccount    OptNilBool `json:"notification_invitee_create_account"`
	NotificationLatestNews              OptNilBool `json:"notification_latest_news"`
	NotificationLike                    OptNilBool `json:"notification_like"`
	NotificationMessageTag              OptNilBool `json:"notification_message_tag"`
	NotificationPopularPost             OptNilBool `json:"notification_popular_post"`
	NotificationProfileScreenshot       OptNilBool `json:"notification_profile_screenshot"`
	NotificationReply                   OptNilBool `json:"notification_reply"`
	NotificationRepost                  OptNilBool `json:"notification_repost"`
	NotificationReview                  OptNilBool `json:"notification_review"`
	NotificationSecurityWarning         OptNilBool `json:"notification_security_warning"`
	NotificationTwitterFriend           OptNilBool `json:"notification_twitter_friend"`
	PrivacyMode                         OptNilBool `json:"privacy_mode"`
	PrivatePost                         OptNilBool `json:"private_post"`
	PrivateUserTimeline                 OptNilBool `json:"private_user_timeline"`
	ShowTotalGiftsReceivedCount         OptNilBool `json:"show_total_gifts_received_count"`
	VipInvisibleFootprintMode           OptNilBool `json:"vip_invisible_footprint_mode"`
	VisibleOnSnsFriendRecommendation    OptNilBool `json:"visible_on_sns_friend_recommendation"`
}

// GetAgeRestrictedOnReview returns the value of AgeRestrictedOnReview.
func (s *Settings) GetAgeRestrictedOnReview() OptNilBool {
	return s.AgeRestrictedOnReview
}

// GetAllowReposts returns the value of AllowReposts.
func (s *Settings) GetAllowReposts() OptNilBool {
	return s.AllowReposts
}

// GetCautionUserChat returns the value of CautionUserChat.
func (s *Settings) GetCautionUserChat() OptNilBool {
	return s.CautionUserChat
}

// GetFollowingOnlyCallInvite returns the value of FollowingOnlyCallInvite.
func (s *Settings) GetFollowingOnlyCallInvite() OptNilBool {
	return s.FollowingOnlyCallInvite
}

// GetFollowingOnlyGroupInvite returns the value of FollowingOnlyGroupInvite.
func (s *Settings) GetFollowingOnlyGroupInvite() OptNilBool {
	return s.FollowingOnlyGroupInvite
}

// GetFollowingRestrictedOnReview returns the value of FollowingRestrictedOnReview.
func (s *Settings) GetFollowingRestrictedOnReview() OptNilBool {
	return s.FollowingRestrictedOnReview
}

// GetHideActiveCall returns the value of HideActiveCall.
func (s *Settings) GetHideActiveCall() OptNilBool {
	return s.HideActiveCall
}

// GetHideGiftsReceived returns the value of HideGiftsReceived.
func (s *Settings) GetHideGiftsReceived() OptNilBool {
	return s.HideGiftsReceived
}

// GetHideHotPost returns the value of HideHotPost.
func (s *Settings) GetHideHotPost() OptNilBool {
	return s.HideHotPost
}

// GetHideOnInvitable returns the value of HideOnInvitable.
func (s *Settings) GetHideOnInvitable() OptNilBool {
	return s.HideOnInvitable
}

// GetHideOnlineStatus returns the value of HideOnlineStatus.
func (s *Settings) GetHideOnlineStatus() OptNilBool {
	return s.HideOnlineStatus
}

// GetHideReplyFollowingTimeline returns the value of HideReplyFollowingTimeline.
func (s *Settings) GetHideReplyFollowingTimeline() OptNilBool {
	return s.HideReplyFollowingTimeline
}

// GetHideReplyPublicTimeline returns the value of HideReplyPublicTimeline.
func (s *Settings) GetHideReplyPublicTimeline() OptNilBool {
	return s.HideReplyPublicTimeline
}

// GetHideVip returns the value of HideVip.
func (s *Settings) GetHideVip() OptNilBool {
	return s.HideVip
}

// GetInvisibleOnUserSearch returns the value of InvisibleOnUserSearch.
func (s *Settings) GetInvisibleOnUserSearch() OptNilBool {
	return s.InvisibleOnUserSearch
}

// GetNoReplyGroupTimeline returns the value of NoReplyGroupTimeline.
func (s *Settings) GetNoReplyGroupTimeline() OptNilBool {
	return s.NoReplyGroupTimeline
}

// GetNotificationBirthdayToFollowers returns the value of NotificationBirthdayToFollowers.
func (s *Settings) GetNotificationBirthdayToFollowers() OptNilBool {
	return s.NotificationBirthdayToFollowers
}

// GetNotificationBulkCallInvite returns the value of NotificationBulkCallInvite.
func (s *Settings) GetNotificationBulkCallInvite() OptNilBool {
	return s.NotificationBulkCallInvite
}

// GetNotificationCallInvite returns the value of NotificationCallInvite.
func (s *Settings) GetNotificationCallInvite() OptNilBool {
	return s.NotificationCallInvite
}

// GetNotificationChat returns the value of NotificationChat.
func (s *Settings) GetNotificationChat() OptNilBool {
	return s.NotificationChat
}

// GetNotificationChatDelete returns the value of NotificationChatDelete.
func (s *Settings) GetNotificationChatDelete() OptNilBool {
	return s.NotificationChatDelete
}

// GetNotificationContactFriend returns the value of NotificationContactFriend.
func (s *Settings) GetNotificationContactFriend() OptNilBool {
	return s.NotificationContactFriend
}

// GetNotificationDailyQuest returns the value of NotificationDailyQuest.
func (s *Settings) GetNotificationDailyQuest() OptNilBool {
	return s.NotificationDailyQuest
}

// GetNotificationDailySummary returns the value of NotificationDailySummary.
func (s *Settings) GetNotificationDailySummary() OptNilBool {
	return s.NotificationDailySummary
}

// GetNotificationFollow returns the value of NotificationFollow.
func (s *Settings) GetNotificationFollow() OptNilBool {
	return s.NotificationFollow
}

// GetNotificationFollowAccept returns the value of NotificationFollowAccept.
func (s *Settings) GetNotificationFollowAccept() OptNilBool {
	return s.NotificationFollowAccept
}

// GetNotificationFollowRequest returns the value of NotificationFollowRequest.
func (s *Settings) GetNotificationFollowRequest() OptNilBool {
	return s.NotificationFollowRequest
}

// GetNotificationFollowerConferenceCall returns the value of NotificationFollowerConferenceCall.
func (s *Settings) GetNotificationFollowerConferenceCall() OptNilBool {
	return s.NotificationFollowerConferenceCall
}

// GetNotificationFollowerCreateGroup returns the value of NotificationFollowerCreateGroup.
func (s *Settings) GetNotificationFollowerCreateGroup() OptNilBool {
	return s.NotificationFollowerCreateGroup
}

// GetNotificationFollowingBirthdateOn returns the value of NotificationFollowingBirthdateOn.
func (s *Settings) GetNotificationFollowingBirthdateOn() OptNilBool {
	return s.NotificationFollowingBirthdateOn
}

// GetNotificationFollowingPostAfterBreak returns the value of NotificationFollowingPostAfterBreak.
func (s *Settings) GetNotificationFollowingPostAfterBreak() OptNilBool {
	return s.NotificationFollowingPostAfterBreak
}

// GetNotificationFollowingsInCall returns the value of NotificationFollowingsInCall.
func (s *Settings) GetNotificationFollowingsInCall() OptNilBool {
	return s.NotificationFollowingsInCall
}

// GetNotificationFootprint returns the value of NotificationFootprint.
func (s *Settings) GetNotificationFootprint() OptNilBool {
	return s.NotificationFootprint
}

// GetNotificationGiftReceive returns the value of NotificationGiftReceive.
func (s *Settings) GetNotificationGiftReceive() OptNilBool {
	return s.NotificationGiftReceive
}

// GetNotificationGroupAccept returns the value of NotificationGroupAccept.
func (s *Settings) GetNotificationGroupAccept() OptNilBool {
	return s.NotificationGroupAccept
}

// GetNotificationGroupConferenceCall returns the value of NotificationGroupConferenceCall.
func (s *Settings) GetNotificationGroupConferenceCall() OptNilBool {
	return s.NotificationGroupConferenceCall
}

// GetNotificationGroupInvite returns the value of NotificationGroupInvite.
func (s *Settings) GetNotificationGroupInvite() OptNilBool {
	return s.NotificationGroupInvite
}

// GetNotificationGroupJoin returns the value of NotificationGroupJoin.
func (s *Settings) GetNotificationGroupJoin() OptNilBool {
	return s.NotificationGroupJoin
}

// GetNotificationGroupMessageTagAll returns the value of NotificationGroupMessageTagAll.
func (s *Settings) GetNotificationGroupMessageTagAll() OptNilBool {
	return s.NotificationGroupMessageTagAll
}

// GetNotificationGroupModerator returns the value of NotificationGroupModerator.
func (s *Settings) GetNotificationGroupModerator() OptNilBool {
	return s.NotificationGroupModerator
}

// GetNotificationGroupPost returns the value of NotificationGroupPost.
func (s *Settings) GetNotificationGroupPost() OptNilBool {
	return s.NotificationGroupPost
}

// GetNotificationGroupRequest returns the value of NotificationGroupRequest.
func (s *Settings) GetNotificationGroupRequest() OptNilBool {
	return s.NotificationGroupRequest
}

// GetNotificationHimaNow returns the value of NotificationHimaNow.
func (s *Settings) GetNotificationHimaNow() OptNilBool {
	return s.NotificationHimaNow
}

// GetNotificationInviteeCreateAccount returns the value of NotificationInviteeCreateAccount.
func (s *Settings) GetNotificationInviteeCreateAccount() OptNilBool {
	return s.NotificationInviteeCreateAccount
}

// GetNotificationLatestNews returns the value of NotificationLatestNews.
func (s *Settings) GetNotificationLatestNews() OptNilBool {
	return s.NotificationLatestNews
}

// GetNotificationLike returns the value of NotificationLike.
func (s *Settings) GetNotificationLike() OptNilBool {
	return s.NotificationLike
}

// GetNotificationMessageTag returns the value of NotificationMessageTag.
func (s *Settings) GetNotificationMessageTag() OptNilBool {
	return s.NotificationMessageTag
}

// GetNotificationPopularPost returns the value of NotificationPopularPost.
func (s *Settings) GetNotificationPopularPost() OptNilBool {
	return s.NotificationPopularPost
}

// GetNotificationProfileScreenshot returns the value of NotificationProfileScreenshot.
func (s *Settings) GetNotificationProfileScreenshot() OptNilBool {
	return s.NotificationProfileScreenshot
}

// GetNotificationReply returns the value of NotificationReply.
func (s *Settings) GetNotificationReply() OptNilBool {
	return s.NotificationReply
}

// GetNotificationRepost returns the value of NotificationRepost.
func (s *Settings) GetNotificationRepost() OptNilBool {
	return s.NotificationRepost
}

// GetNotificationReview returns the value of NotificationReview.
func (s *Settings) GetNotificationReview() OptNilBool {
	return s.NotificationReview
}

// GetNotificationSecurityWarning returns the value of NotificationSecurityWarning.
func (s *Settings) GetNotificationSecurityWarning() OptNilBool {
	return s.NotificationSecurityWarning
}

// GetNotificationTwitterFriend returns the value of NotificationTwitterFriend.
func (s *Settings) GetNotificationTwitterFriend() OptNilBool {
	return s.NotificationTwitterFriend
}

// GetPrivacyMode returns the value of PrivacyMode.
func (s *Settings) GetPrivacyMode() OptNilBool {
	return s.PrivacyMode
}

// GetPrivatePost returns the value of PrivatePost.
func (s *Settings) GetPrivatePost() OptNilBool {
	return s.PrivatePost
}

// GetPrivateUserTimeline returns the value of PrivateUserTimeline.
func (s *Settings) GetPrivateUserTimeline() OptNilBool {
	return s.PrivateUserTimeline
}

// GetShowTotalGiftsReceivedCount returns the value of ShowTotalGiftsReceivedCount.
func (s *Settings) GetShowTotalGiftsReceivedCount() OptNilBool {
	return s.ShowTotalGiftsReceivedCount
}

// GetVipInvisibleFootprintMode returns the value of VipInvisibleFootprintMode.
func (s *Settings) GetVipInvisibleFootprintMode() OptNilBool {
	return s.VipInvisibleFootprintMode
}

// GetVisibleOnSnsFriendRecommendation returns the value of VisibleOnSnsFriendRecommendation.
func (s *Settings) GetVisibleOnSnsFriendRecommendation() OptNilBool {
	return s.VisibleOnSnsFriendRecommendation
}

// SetAgeRestrictedOnReview sets the value of AgeRestrictedOnReview.
func (s *Settings) SetAgeRestrictedOnReview(val OptNilBool) {
	s.AgeRestrictedOnReview = val
}

// SetAllowReposts sets the value of AllowReposts.
func (s *Settings) SetAllowReposts(val OptNilBool) {
	s.AllowReposts = val
}

// SetCautionUserChat sets the value of CautionUserChat.
func (s *Settings) SetCautionUserChat(val OptNilBool) {
	s.CautionUserChat = val
}

// SetFollowingOnlyCallInvite sets the value of FollowingOnlyCallInvite.
func (s *Settings) SetFollowingOnlyCallInvite(val OptNilBool) {
	s.FollowingOnlyCallInvite = val
}

// SetFollowingOnlyGroupInvite sets the value of FollowingOnlyGroupInvite.
func (s *Settings) SetFollowingOnlyGroupInvite(val OptNilBool) {
	s.FollowingOnlyGroupInvite = val
}

// SetFollowingRestrictedOnReview sets the value of FollowingRestrictedOnReview.
func (s *Settings) SetFollowingRestrictedOnReview(val OptNilBool) {
	s.FollowingRestrictedOnReview = val
}

// SetHideActiveCall sets the value of HideActiveCall.
func (s *Settings) SetHideActiveCall(val OptNilBool) {
	s.HideActiveCall = val
}

// SetHideGiftsReceived sets the value of HideGiftsReceived.
func (s *Settings) SetHideGiftsReceived(val OptNilBool) {
	s.HideGiftsReceived = val
}

// SetHideHotPost sets the value of HideHotPost.
func (s *Settings) SetHideHotPost(val OptNilBool) {
	s.HideHotPost = val
}

// SetHideOnInvitable sets the value of HideOnInvitable.
func (s *Settings) SetHideOnInvitable(val OptNilBool) {
	s.HideOnInvitable = val
}

// SetHideOnlineStatus sets the value of HideOnlineStatus.
func (s *Settings) SetHideOnlineStatus(val OptNilBool) {
	s.HideOnlineStatus = val
}

// SetHideReplyFollowingTimeline sets the value of HideReplyFollowingTimeline.
func (s *Settings) SetHideReplyFollowingTimeline(val OptNilBool) {
	s.HideReplyFollowingTimeline = val
}

// SetHideReplyPublicTimeline sets the value of HideReplyPublicTimeline.
func (s *Settings) SetHideReplyPublicTimeline(val OptNilBool) {
	s.HideReplyPublicTimeline = val
}

// SetHideVip sets the value of HideVip.
func (s *Settings) SetHideVip(val OptNilBool) {
	s.HideVip = val
}

// SetInvisibleOnUserSearch sets the value of InvisibleOnUserSearch.
func (s *Settings) SetInvisibleOnUserSearch(val OptNilBool) {
	s.InvisibleOnUserSearch = val
}

// SetNoReplyGroupTimeline sets the value of NoReplyGroupTimeline.
func (s *Settings) SetNoReplyGroupTimeline(val OptNilBool) {
	s.NoReplyGroupTimeline = val
}

// SetNotificationBirthdayToFollowers sets the value of NotificationBirthdayToFollowers.
func (s *Settings) SetNotificationBirthdayToFollowers(val OptNilBool) {
	s.NotificationBirthdayToFollowers = val
}

// SetNotificationBulkCallInvite sets the value of NotificationBulkCallInvite.
func (s *Settings) SetNotificationBulkCallInvite(val OptNilBool) {
	s.NotificationBulkCallInvite = val
}

// SetNotificationCallInvite sets the value of NotificationCallInvite.
func (s *Settings) SetNotificationCallInvite(val OptNilBool) {
	s.NotificationCallInvite = val
}

// SetNotificationChat sets the value of NotificationChat.
func (s *Settings) SetNotificationChat(val OptNilBool) {
	s.NotificationChat = val
}

// SetNotificationChatDelete sets the value of NotificationChatDelete.
func (s *Settings) SetNotificationChatDelete(val OptNilBool) {
	s.NotificationChatDelete = val
}

// SetNotificationContactFriend sets the value of NotificationContactFriend.
func (s *Settings) SetNotificationContactFriend(val OptNilBool) {
	s.NotificationContactFriend = val
}

// SetNotificationDailyQuest sets the value of NotificationDailyQuest.
func (s *Settings) SetNotificationDailyQuest(val OptNilBool) {
	s.NotificationDailyQuest = val
}

// SetNotificationDailySummary sets the value of NotificationDailySummary.
func (s *Settings) SetNotificationDailySummary(val OptNilBool) {
	s.NotificationDailySummary = val
}

// SetNotificationFollow sets the value of NotificationFollow.
func (s *Settings) SetNotificationFollow(val OptNilBool) {
	s.NotificationFollow = val
}

// SetNotificationFollowAccept sets the value of NotificationFollowAccept.
func (s *Settings) SetNotificationFollowAccept(val OptNilBool) {
	s.NotificationFollowAccept = val
}

// SetNotificationFollowRequest sets the value of NotificationFollowRequest.
func (s *Settings) SetNotificationFollowRequest(val OptNilBool) {
	s.NotificationFollowRequest = val
}

// SetNotificationFollowerConferenceCall sets the value of NotificationFollowerConferenceCall.
func (s *Settings) SetNotificationFollowerConferenceCall(val OptNilBool) {
	s.NotificationFollowerConferenceCall = val
}

// SetNotificationFollowerCreateGroup sets the value of NotificationFollowerCreateGroup.
func (s *Settings) SetNotificationFollowerCreateGroup(val OptNilBool) {
	s.NotificationFollowerCreateGroup = val
}

// SetNotificationFollowingBirthdateOn sets the value of NotificationFollowingBirthdateOn.
func (s *Settings) SetNotificationFollowingBirthdateOn(val OptNilBool) {
	s.NotificationFollowingBirthdateOn = val
}

// SetNotificationFollowingPostAfterBreak sets the value of NotificationFollowingPostAfterBreak.
func (s *Settings) SetNotificationFollowingPostAfterBreak(val OptNilBool) {
	s.NotificationFollowingPostAfterBreak = val
}

// SetNotificationFollowingsInCall sets the value of NotificationFollowingsInCall.
func (s *Settings) SetNotificationFollowingsInCall(val OptNilBool) {
	s.NotificationFollowingsInCall = val
}

// SetNotificationFootprint sets the value of NotificationFootprint.
func (s *Settings) SetNotificationFootprint(val OptNilBool) {
	s.NotificationFootprint = val
}

// SetNotificationGiftReceive sets the value of NotificationGiftReceive.
func (s *Settings) SetNotificationGiftReceive(val OptNilBool) {
	s.NotificationGiftReceive = val
}

// SetNotificationGroupAccept sets the value of NotificationGroupAccept.
func (s *Settings) SetNotificationGroupAccept(val OptNilBool) {
	s.NotificationGroupAccept = val
}

// SetNotificationGroupConferenceCall sets the value of NotificationGroupConferenceCall.
func (s *Settings) SetNotificationGroupConferenceCall(val OptNilBool) {
	s.NotificationGroupConferenceCall = val
}

// SetNotificationGroupInvite sets the value of NotificationGroupInvite.
func (s *Settings) SetNotificationGroupInvite(val OptNilBool) {
	s.NotificationGroupInvite = val
}

// SetNotificationGroupJoin sets the value of NotificationGroupJoin.
func (s *Settings) SetNotificationGroupJoin(val OptNilBool) {
	s.NotificationGroupJoin = val
}

// SetNotificationGroupMessageTagAll sets the value of NotificationGroupMessageTagAll.
func (s *Settings) SetNotificationGroupMessageTagAll(val OptNilBool) {
	s.NotificationGroupMessageTagAll = val
}

// SetNotificationGroupModerator sets the value of NotificationGroupModerator.
func (s *Settings) SetNotificationGroupModerator(val OptNilBool) {
	s.NotificationGroupModerator = val
}

// SetNotificationGroupPost sets the value of NotificationGroupPost.
func (s *Settings) SetNotificationGroupPost(val OptNilBool) {
	s.NotificationGroupPost = val
}

// SetNotificationGroupRequest sets the value of NotificationGroupRequest.
func (s *Settings) SetNotificationGroupRequest(val OptNilBool) {
	s.NotificationGroupRequest = val
}

// SetNotificationHimaNow sets the value of NotificationHimaNow.
func (s *Settings) SetNotificationHimaNow(val OptNilBool) {
	s.NotificationHimaNow = val
}

// SetNotificationInviteeCreateAccount sets the value of NotificationInviteeCreateAccount.
func (s *Settings) SetNotificationInviteeCreateAccount(val OptNilBool) {
	s.NotificationInviteeCreateAccount = val
}

// SetNotificationLatestNews sets the value of NotificationLatestNews.
func (s *Settings) SetNotificationLatestNews(val OptNilBool) {
	s.NotificationLatestNews = val
}

// SetNotificationLike sets the value of NotificationLike.
func (s *Settings) SetNotificationLike(val OptNilBool) {
	s.NotificationLike = val
}

// SetNotificationMessageTag sets the value of NotificationMessageTag.
func (s *Settings) SetNotificationMessageTag(val OptNilBool) {
	s.NotificationMessageTag = val
}

// SetNotificationPopularPost sets the value of NotificationPopularPost.
func (s *Settings) SetNotificationPopularPost(val OptNilBool) {
	s.NotificationPopularPost = val
}

// SetNotificationProfileScreenshot sets the value of NotificationProfileScreenshot.
func (s *Settings) SetNotificationProfileScreenshot(val OptNilBool) {
	s.NotificationProfileScreenshot = val
}

// SetNotificationReply sets the value of NotificationReply.
func (s *Settings) SetNotificationReply(val OptNilBool) {
	s.NotificationReply = val
}

// SetNotificationRepost sets the value of NotificationRepost.
func (s *Settings) SetNotificationRepost(val OptNilBool) {
	s.NotificationRepost = val
}

// SetNotificationReview sets the value of NotificationReview.
func (s *Settings) SetNotificationReview(val OptNilBool) {
	s.NotificationReview = val
}

// SetNotificationSecurityWarning sets the value of NotificationSecurityWarning.
func (s *Settings) SetNotificationSecurityWarning(val OptNilBool) {
	s.NotificationSecurityWarning = val
}

// SetNotificationTwitterFriend sets the value of NotificationTwitterFriend.
func (s *Settings) SetNotificationTwitterFriend(val OptNilBool) {
	s.NotificationTwitterFriend = val
}

// SetPrivacyMode sets the value of PrivacyMode.
func (s *Settings) SetPrivacyMode(val OptNilBool) {
	s.PrivacyMode = val
}

// SetPrivatePost sets the value of PrivatePost.
func (s *Settings) SetPrivatePost(val OptNilBool) {
	s.PrivatePost = val
}

// SetPrivateUserTimeline sets the value of PrivateUserTimeline.
func (s *Settings) SetPrivateUserTimeline(val OptNilBool) {
	s.PrivateUserTimeline = val
}

// SetShowTotalGiftsReceivedCount sets the value of ShowTotalGiftsReceivedCount.
func (s *Settings) SetShowTotalGiftsReceivedCount(val OptNilBool) {
	s.ShowTotalGiftsReceivedCount = val
}

// SetVipInvisibleFootprintMode sets the value of VipInvisibleFootprintMode.
func (s *Settings) SetVipInvisibleFootprintMode(val OptNilBool) {
	s.VipInvisibleFootprintMode = val
}

// SetVisibleOnSnsFriendRecommendation sets the value of VisibleOnSnsFriendRecommendation.
func (s *Settings) SetVisibleOnSnsFriendRecommendation(val OptNilBool) {
	s.VisibleOnSnsFriendRecommendation = val
}

// Ref: #/components/schemas/Shareable
type Shareable struct {
	Group  OptGroup      `json:"group"`
	Post   OptPost       `json:"post"`
	Thread OptThreadInfo `json:"thread"`
}

// GetGroup returns the value of Group.
func (s *Shareable) GetGroup() OptGroup {
	return s.Group
}

// GetPost returns the value of Post.
func (s *Shareable) GetPost() OptPost {
	return s.Post
}

// GetThread returns the value of Thread.
func (s *Shareable) GetThread() OptThreadInfo {
	return s.Thread
}

// SetGroup sets the value of Group.
func (s *Shareable) SetGroup(val OptGroup) {
	s.Group = val
}

// SetPost sets the value of Post.
func (s *Shareable) SetPost(val OptPost) {
	s.Post = val
}

// SetThread sets the value of Thread.
func (s *Shareable) SetThread(val OptThreadInfo) {
	s.Thread = val
}

// Ref: #/components/schemas/SharedUrl
type SharedUrl struct {
	Description OptNilString `json:"description"`
	ImageURL    OptNilString `json:"image_url"`
	Title       OptNilString `json:"title"`
	URL         OptNilString `json:"url"`
}

// GetDescription returns the value of Description.
func (s *SharedUrl) GetDescription() OptNilString {
	return s.Description
}

// GetImageURL returns the value of ImageURL.
func (s *SharedUrl) GetImageURL() OptNilString {
	return s.ImageURL
}

// GetTitle returns the value of Title.
func (s *SharedUrl) GetTitle() OptNilString {
	return s.Title
}

// GetURL returns the value of URL.
func (s *SharedUrl) GetURL() OptNilString {
	return s.URL
}

// SetDescription sets the value of Description.
func (s *SharedUrl) SetDescription(val OptNilString) {
	s.Description = val
}

// SetImageURL sets the value of ImageURL.
func (s *SharedUrl) SetImageURL(val OptNilString) {
	s.ImageURL = val
}

// SetTitle sets the value of Title.
func (s *SharedUrl) SetTitle(val OptNilString) {
	s.Title = val
}

// SetURL sets the value of URL.
func (s *SharedUrl) SetURL(val OptNilString) {
	s.URL = val
}

// Ref: #/components/schemas/SignaturePayload
type SignaturePayload struct {
	Action       OptNilString `json:"action"`
	CallID       OptNilInt64  `json:"call_id"`
	ReceiverUUID OptNilString `json:"receiver_uuid"`
	SenderUUID   OptNilString `json:"sender_uuid"`
	Signature    OptNilString `json:"signature"`
	Timestamp    OptNilInt64  `json:"timestamp"`
}

// GetAction returns the value of Action.
func (s *SignaturePayload) GetAction() OptNilString {
	return s.Action
}

// GetCallID returns the value of CallID.
func (s *SignaturePayload) GetCallID() OptNilInt64 {
	return s.CallID
}

// GetReceiverUUID returns the value of ReceiverUUID.
func (s *SignaturePayload) GetReceiverUUID() OptNilString {
	return s.ReceiverUUID
}

// GetSenderUUID returns the value of SenderUUID.
func (s *SignaturePayload) GetSenderUUID() OptNilString {
	return s.SenderUUID
}

// GetSignature returns the value of Signature.
func (s *SignaturePayload) GetSignature() OptNilString {
	return s.Signature
}

// GetTimestamp returns the value of Timestamp.
func (s *SignaturePayload) GetTimestamp() OptNilInt64 {
	return s.Timestamp
}

// SetAction sets the value of Action.
func (s *SignaturePayload) SetAction(val OptNilString) {
	s.Action = val
}

// SetCallID sets the value of CallID.
func (s *SignaturePayload) SetCallID(val OptNilInt64) {
	s.CallID = val
}

// SetReceiverUUID sets the value of ReceiverUUID.
func (s *SignaturePayload) SetReceiverUUID(val OptNilString) {
	s.ReceiverUUID = val
}

// SetSenderUUID sets the value of SenderUUID.
func (s *SignaturePayload) SetSenderUUID(val OptNilString) {
	s.SenderUUID = val
}

// SetSignature sets the value of Signature.
func (s *SignaturePayload) SetSignature(val OptNilString) {
	s.Signature = val
}

// SetTimestamp sets the value of Timestamp.
func (s *SignaturePayload) SetTimestamp(val OptNilInt64) {
	s.Timestamp = val
}

// Ref: #/components/schemas/SnsInfo
type SnsInfo struct {
	Biography    OptNilString `json:"biography"`
	Gender       OptNilString `json:"gender"`
	Nickname     OptNilString `json:"nickname"`
	ProfileImage OptNilString `json:"profile_image"`
	Type         OptNilString `json:"type"`
	UID          OptNilString `json:"uid"`
}

// GetBiography returns the value of Biography.
func (s *SnsInfo) GetBiography() OptNilString {
	return s.Biography
}

// GetGender returns the value of Gender.
func (s *SnsInfo) GetGender() OptNilString {
	return s.Gender
}

// GetNickname returns the value of Nickname.
func (s *SnsInfo) GetNickname() OptNilString {
	return s.Nickname
}

// GetProfileImage returns the value of ProfileImage.
func (s *SnsInfo) GetProfileImage() OptNilString {
	return s.ProfileImage
}

// GetType returns the value of Type.
func (s *SnsInfo) GetType() OptNilString {
	return s.Type
}

// GetUID returns the value of UID.
func (s *SnsInfo) GetUID() OptNilString {
	return s.UID
}

// SetBiography sets the value of Biography.
func (s *SnsInfo) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetGender sets the value of Gender.
func (s *SnsInfo) SetGender(val OptNilString) {
	s.Gender = val
}

// SetNickname sets the value of Nickname.
func (s *SnsInfo) SetNickname(val OptNilString) {
	s.Nickname = val
}

// SetProfileImage sets the value of ProfileImage.
func (s *SnsInfo) SetProfileImage(val OptNilString) {
	s.ProfileImage = val
}

// SetType sets the value of Type.
func (s *SnsInfo) SetType(val OptNilString) {
	s.Type = val
}

// SetUID sets the value of UID.
func (s *SnsInfo) SetUID(val OptNilString) {
	s.UID = val
}

type StartConferenceCallReq struct {
	CallSid      OptString `json:"call_sid"`
	ConferenceID OptInt64  `json:"conference_id"`
}

// GetCallSid returns the value of CallSid.
func (s *StartConferenceCallReq) GetCallSid() OptString {
	return s.CallSid
}

// GetConferenceID returns the value of ConferenceID.
func (s *StartConferenceCallReq) GetConferenceID() OptInt64 {
	return s.ConferenceID
}

// SetCallSid sets the value of CallSid.
func (s *StartConferenceCallReq) SetCallSid(val OptString) {
	s.CallSid = val
}

// SetConferenceID sets the value of ConferenceID.
func (s *StartConferenceCallReq) SetConferenceID(val OptInt64) {
	s.ConferenceID = val
}

// Ref: #/components/schemas/Sticker
type Sticker struct {
	Extension     OptNilString `json:"extension"`
	Height        OptNilInt32  `json:"height"`
	ID            OptNilInt64  `json:"id"`
	StickerPackID OptNilInt64  `json:"sticker_pack_id"`
	URL           OptNilString `json:"url"`
	Width         OptNilInt32  `json:"width"`
}

// GetExtension returns the value of Extension.
func (s *Sticker) GetExtension() OptNilString {
	return s.Extension
}

// GetHeight returns the value of Height.
func (s *Sticker) GetHeight() OptNilInt32 {
	return s.Height
}

// GetID returns the value of ID.
func (s *Sticker) GetID() OptNilInt64 {
	return s.ID
}

// GetStickerPackID returns the value of StickerPackID.
func (s *Sticker) GetStickerPackID() OptNilInt64 {
	return s.StickerPackID
}

// GetURL returns the value of URL.
func (s *Sticker) GetURL() OptNilString {
	return s.URL
}

// GetWidth returns the value of Width.
func (s *Sticker) GetWidth() OptNilInt32 {
	return s.Width
}

// SetExtension sets the value of Extension.
func (s *Sticker) SetExtension(val OptNilString) {
	s.Extension = val
}

// SetHeight sets the value of Height.
func (s *Sticker) SetHeight(val OptNilInt32) {
	s.Height = val
}

// SetID sets the value of ID.
func (s *Sticker) SetID(val OptNilInt64) {
	s.ID = val
}

// SetStickerPackID sets the value of StickerPackID.
func (s *Sticker) SetStickerPackID(val OptNilInt64) {
	s.StickerPackID = val
}

// SetURL sets the value of URL.
func (s *Sticker) SetURL(val OptNilString) {
	s.URL = val
}

// SetWidth sets the value of Width.
func (s *Sticker) SetWidth(val OptNilInt32) {
	s.Width = val
}

// Ref: #/components/schemas/StickerPack
type StickerPack struct {
	Cover       OptNilString       `json:"cover"`
	Description OptNilString       `json:"description"`
	ID          OptNilInt64        `json:"id"`
	Name        OptNilString       `json:"name"`
	Order       OptNilInt32        `json:"order"`
	Stickers    OptNilStickerArray `json:"stickers"`
}

// GetCover returns the value of Cover.
func (s *StickerPack) GetCover() OptNilString {
	return s.Cover
}

// GetDescription returns the value of Description.
func (s *StickerPack) GetDescription() OptNilString {
	return s.Description
}

// GetID returns the value of ID.
func (s *StickerPack) GetID() OptNilInt64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *StickerPack) GetName() OptNilString {
	return s.Name
}

// GetOrder returns the value of Order.
func (s *StickerPack) GetOrder() OptNilInt32 {
	return s.Order
}

// GetStickers returns the value of Stickers.
func (s *StickerPack) GetStickers() OptNilStickerArray {
	return s.Stickers
}

// SetCover sets the value of Cover.
func (s *StickerPack) SetCover(val OptNilString) {
	s.Cover = val
}

// SetDescription sets the value of Description.
func (s *StickerPack) SetDescription(val OptNilString) {
	s.Description = val
}

// SetID sets the value of ID.
func (s *StickerPack) SetID(val OptNilInt64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *StickerPack) SetName(val OptNilString) {
	s.Name = val
}

// SetOrder sets the value of Order.
func (s *StickerPack) SetOrder(val OptNilInt32) {
	s.Order = val
}

// SetStickers sets the value of Stickers.
func (s *StickerPack) SetStickers(val OptNilStickerArray) {
	s.Stickers = val
}

// Ref: #/components/schemas/StickerPacksResponse
type StickerPacksResponse struct {
	StickerPacks OptNilStickerPackArray `json:"sticker_packs"`
}

// GetStickerPacks returns the value of StickerPacks.
func (s *StickerPacksResponse) GetStickerPacks() OptNilStickerPackArray {
	return s.StickerPacks
}

// SetStickerPacks sets the value of StickerPacks.
func (s *StickerPacksResponse) SetStickerPacks(val OptNilStickerPackArray) {
	s.StickerPacks = val
}

// Ref: #/components/schemas/Survey
type Survey struct {
	Choices    OptNilChoiceArray `json:"choices"`
	ID         OptNilInt64       `json:"id"`
	Voted      OptNilBool        `json:"voted"`
	VotesCount OptNilInt32       `json:"votes_count"`
}

// GetChoices returns the value of Choices.
func (s *Survey) GetChoices() OptNilChoiceArray {
	return s.Choices
}

// GetID returns the value of ID.
func (s *Survey) GetID() OptNilInt64 {
	return s.ID
}

// GetVoted returns the value of Voted.
func (s *Survey) GetVoted() OptNilBool {
	return s.Voted
}

// GetVotesCount returns the value of VotesCount.
func (s *Survey) GetVotesCount() OptNilInt32 {
	return s.VotesCount
}

// SetChoices sets the value of Choices.
func (s *Survey) SetChoices(val OptNilChoiceArray) {
	s.Choices = val
}

// SetID sets the value of ID.
func (s *Survey) SetID(val OptNilInt64) {
	s.ID = val
}

// SetVoted sets the value of Voted.
func (s *Survey) SetVoted(val OptNilBool) {
	s.Voted = val
}

// SetVotesCount sets the value of VotesCount.
func (s *Survey) SetVotesCount(val OptNilInt32) {
	s.VotesCount = val
}

// TakeOverGroupNoContent is response for TakeOverGroup operation.
type TakeOverGroupNoContent struct{}

// Ref: #/components/schemas/ThreadInfo
type ThreadInfo struct {
	CreatedAt   OptNilInt64  `json:"created_at"`
	ID          OptNilInt64  `json:"id"`
	IsJoined    OptNilBool   `json:"is_joined"`
	LastPost    OptPost      `json:"last_post"`
	NewUpdates  OptNilBool   `json:"new_updates"`
	Owner       OptRealmUser `json:"owner"`
	PostsCount  OptNilInt32  `json:"posts_count"`
	ThreadIcon  OptNilString `json:"thread_icon"`
	Title       OptNilString `json:"title"`
	UnreadCount OptNilInt32  `json:"unread_count"`
	UpdatedAt   OptNilInt64  `json:"updated_at"`
}

// GetCreatedAt returns the value of CreatedAt.
func (s *ThreadInfo) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetID returns the value of ID.
func (s *ThreadInfo) GetID() OptNilInt64 {
	return s.ID
}

// GetIsJoined returns the value of IsJoined.
func (s *ThreadInfo) GetIsJoined() OptNilBool {
	return s.IsJoined
}

// GetLastPost returns the value of LastPost.
func (s *ThreadInfo) GetLastPost() OptPost {
	return s.LastPost
}

// GetNewUpdates returns the value of NewUpdates.
func (s *ThreadInfo) GetNewUpdates() OptNilBool {
	return s.NewUpdates
}

// GetOwner returns the value of Owner.
func (s *ThreadInfo) GetOwner() OptRealmUser {
	return s.Owner
}

// GetPostsCount returns the value of PostsCount.
func (s *ThreadInfo) GetPostsCount() OptNilInt32 {
	return s.PostsCount
}

// GetThreadIcon returns the value of ThreadIcon.
func (s *ThreadInfo) GetThreadIcon() OptNilString {
	return s.ThreadIcon
}

// GetTitle returns the value of Title.
func (s *ThreadInfo) GetTitle() OptNilString {
	return s.Title
}

// GetUnreadCount returns the value of UnreadCount.
func (s *ThreadInfo) GetUnreadCount() OptNilInt32 {
	return s.UnreadCount
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *ThreadInfo) GetUpdatedAt() OptNilInt64 {
	return s.UpdatedAt
}

// SetCreatedAt sets the value of CreatedAt.
func (s *ThreadInfo) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetID sets the value of ID.
func (s *ThreadInfo) SetID(val OptNilInt64) {
	s.ID = val
}

// SetIsJoined sets the value of IsJoined.
func (s *ThreadInfo) SetIsJoined(val OptNilBool) {
	s.IsJoined = val
}

// SetLastPost sets the value of LastPost.
func (s *ThreadInfo) SetLastPost(val OptPost) {
	s.LastPost = val
}

// SetNewUpdates sets the value of NewUpdates.
func (s *ThreadInfo) SetNewUpdates(val OptNilBool) {
	s.NewUpdates = val
}

// SetOwner sets the value of Owner.
func (s *ThreadInfo) SetOwner(val OptRealmUser) {
	s.Owner = val
}

// SetPostsCount sets the value of PostsCount.
func (s *ThreadInfo) SetPostsCount(val OptNilInt32) {
	s.PostsCount = val
}

// SetThreadIcon sets the value of ThreadIcon.
func (s *ThreadInfo) SetThreadIcon(val OptNilString) {
	s.ThreadIcon = val
}

// SetTitle sets the value of Title.
func (s *ThreadInfo) SetTitle(val OptNilString) {
	s.Title = val
}

// SetUnreadCount sets the value of UnreadCount.
func (s *ThreadInfo) SetUnreadCount(val OptNilInt32) {
	s.UnreadCount = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *ThreadInfo) SetUpdatedAt(val OptNilInt64) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/TimelineSettings
type TimelineSettings struct {
	FavesFilter                OptNilBool `json:"faves_filter"`
	HideHotPost                OptNilBool `json:"hide_hot_post"`
	HideReplyFollowingTimeline OptNilBool `json:"hide_reply_following_timeline"`
	HideReplyPublicTimeline    OptNilBool `json:"hide_reply_public_timeline"`
}

// GetFavesFilter returns the value of FavesFilter.
func (s *TimelineSettings) GetFavesFilter() OptNilBool {
	return s.FavesFilter
}

// GetHideHotPost returns the value of HideHotPost.
func (s *TimelineSettings) GetHideHotPost() OptNilBool {
	return s.HideHotPost
}

// GetHideReplyFollowingTimeline returns the value of HideReplyFollowingTimeline.
func (s *TimelineSettings) GetHideReplyFollowingTimeline() OptNilBool {
	return s.HideReplyFollowingTimeline
}

// GetHideReplyPublicTimeline returns the value of HideReplyPublicTimeline.
func (s *TimelineSettings) GetHideReplyPublicTimeline() OptNilBool {
	return s.HideReplyPublicTimeline
}

// SetFavesFilter sets the value of FavesFilter.
func (s *TimelineSettings) SetFavesFilter(val OptNilBool) {
	s.FavesFilter = val
}

// SetHideHotPost sets the value of HideHotPost.
func (s *TimelineSettings) SetHideHotPost(val OptNilBool) {
	s.HideHotPost = val
}

// SetHideReplyFollowingTimeline sets the value of HideReplyFollowingTimeline.
func (s *TimelineSettings) SetHideReplyFollowingTimeline(val OptNilBool) {
	s.HideReplyFollowingTimeline = val
}

// SetHideReplyPublicTimeline sets the value of HideReplyPublicTimeline.
func (s *TimelineSettings) SetHideReplyPublicTimeline(val OptNilBool) {
	s.HideReplyPublicTimeline = val
}

// Ref: #/components/schemas/Title
type Title struct {
	APIValue OptNilString `json:"api_value"`
}

// GetAPIValue returns the value of APIValue.
func (s *Title) GetAPIValue() OptNilString {
	return s.APIValue
}

// SetAPIValue sets the value of APIValue.
func (s *Title) SetAPIValue(val OptNilString) {
	s.APIValue = val
}

// Ref: #/components/schemas/TokenResponse
type TokenResponse struct {
	AccessToken  OptNilString `json:"access_token"`
	CreatedAt    OptNilInt64  `json:"created_at"`
	ExpiresIn    OptNilInt64  `json:"expires_in"`
	ID           OptNilInt64  `json:"id"`
	RefreshToken OptNilString `json:"refresh_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *TokenResponse) GetAccessToken() OptNilString {
	return s.AccessToken
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TokenResponse) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetExpiresIn returns the value of ExpiresIn.
func (s *TokenResponse) GetExpiresIn() OptNilInt64 {
	return s.ExpiresIn
}

// GetID returns the value of ID.
func (s *TokenResponse) GetID() OptNilInt64 {
	return s.ID
}

// GetRefreshToken returns the value of RefreshToken.
func (s *TokenResponse) GetRefreshToken() OptNilString {
	return s.RefreshToken
}

// SetAccessToken sets the value of AccessToken.
func (s *TokenResponse) SetAccessToken(val OptNilString) {
	s.AccessToken = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TokenResponse) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetExpiresIn sets the value of ExpiresIn.
func (s *TokenResponse) SetExpiresIn(val OptNilInt64) {
	s.ExpiresIn = val
}

// SetID sets the value of ID.
func (s *TokenResponse) SetID(val OptNilInt64) {
	s.ID = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *TokenResponse) SetRefreshToken(val OptNilString) {
	s.RefreshToken = val
}

// Ref: #/components/schemas/TotalChatRequestResponse
type TotalChatRequestResponse struct {
	Total OptNilInt32 `json:"total"`
}

// GetTotal returns the value of Total.
func (s *TotalChatRequestResponse) GetTotal() OptNilInt32 {
	return s.Total
}

// SetTotal sets the value of Total.
func (s *TotalChatRequestResponse) SetTotal(val OptNilInt32) {
	s.Total = val
}

// Ref: #/components/schemas/TransactionGiftReceived
type TransactionGiftReceived struct {
	Icon     OptNilString                   `json:"icon"`
	Item     OptTransactionGiftReceivedItem `json:"item"`
	ItemID   OptNilInt64                    `json:"item_id"`
	Quantity OptNilInt32                    `json:"quantity"`
}

// GetIcon returns the value of Icon.
func (s *TransactionGiftReceived) GetIcon() OptNilString {
	return s.Icon
}

// GetItem returns the value of Item.
func (s *TransactionGiftReceived) GetItem() OptTransactionGiftReceivedItem {
	return s.Item
}

// GetItemID returns the value of ItemID.
func (s *TransactionGiftReceived) GetItemID() OptNilInt64 {
	return s.ItemID
}

// GetQuantity returns the value of Quantity.
func (s *TransactionGiftReceived) GetQuantity() OptNilInt32 {
	return s.Quantity
}

// SetIcon sets the value of Icon.
func (s *TransactionGiftReceived) SetIcon(val OptNilString) {
	s.Icon = val
}

// SetItem sets the value of Item.
func (s *TransactionGiftReceived) SetItem(val OptTransactionGiftReceivedItem) {
	s.Item = val
}

// SetItemID sets the value of ItemID.
func (s *TransactionGiftReceived) SetItemID(val OptNilInt64) {
	s.ItemID = val
}

// SetQuantity sets the value of Quantity.
func (s *TransactionGiftReceived) SetQuantity(val OptNilInt32) {
	s.Quantity = val
}

// Ref: #/components/schemas/TransactionGiftReceivedItem
type TransactionGiftReceivedItem struct {
	Currency OptNilString `json:"currency"`
	ID       OptNilInt64  `json:"id"`
	Price    OptNilInt32  `json:"price"`
	Title    OptNilString `json:"title"`
}

// GetCurrency returns the value of Currency.
func (s *TransactionGiftReceivedItem) GetCurrency() OptNilString {
	return s.Currency
}

// GetID returns the value of ID.
func (s *TransactionGiftReceivedItem) GetID() OptNilInt64 {
	return s.ID
}

// GetPrice returns the value of Price.
func (s *TransactionGiftReceivedItem) GetPrice() OptNilInt32 {
	return s.Price
}

// GetTitle returns the value of Title.
func (s *TransactionGiftReceivedItem) GetTitle() OptNilString {
	return s.Title
}

// SetCurrency sets the value of Currency.
func (s *TransactionGiftReceivedItem) SetCurrency(val OptNilString) {
	s.Currency = val
}

// SetID sets the value of ID.
func (s *TransactionGiftReceivedItem) SetID(val OptNilInt64) {
	s.ID = val
}

// SetPrice sets the value of Price.
func (s *TransactionGiftReceivedItem) SetPrice(val OptNilInt32) {
	s.Price = val
}

// SetTitle sets the value of Title.
func (s *TransactionGiftReceivedItem) SetTitle(val OptNilString) {
	s.Title = val
}

// TransferGroupNoContent is response for TransferGroup operation.
type TransferGroupNoContent struct{}

type TransferGroupReq struct {
	APIKey     OptString `json:"api_key"`
	SignedInfo OptString `json:"signed_info"`
	Timestamp  OptInt64  `json:"timestamp"`
	UserID     OptInt64  `json:"user_id"`
	UUID       OptString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *TransferGroupReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetSignedInfo returns the value of SignedInfo.
func (s *TransferGroupReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *TransferGroupReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUserID returns the value of UserID.
func (s *TransferGroupReq) GetUserID() OptInt64 {
	return s.UserID
}

// GetUUID returns the value of UUID.
func (s *TransferGroupReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *TransferGroupReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *TransferGroupReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *TransferGroupReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUserID sets the value of UserID.
func (s *TransferGroupReq) SetUserID(val OptInt64) {
	s.UserID = val
}

// SetUUID sets the value of UUID.
func (s *TransferGroupReq) SetUUID(val OptString) {
	s.UUID = val
}

// Ref: #/components/schemas/TwoFAStatusResponse
type TwoFAStatusResponse struct {
	TwoFaAuthRequired OptTwoFaAuthRequiredDTO `json:"two_fa_auth_required"`
	TwoFaEnabled      OptNilBool              `json:"two_fa_enabled"`
}

// GetTwoFaAuthRequired returns the value of TwoFaAuthRequired.
func (s *TwoFAStatusResponse) GetTwoFaAuthRequired() OptTwoFaAuthRequiredDTO {
	return s.TwoFaAuthRequired
}

// GetTwoFaEnabled returns the value of TwoFaEnabled.
func (s *TwoFAStatusResponse) GetTwoFaEnabled() OptNilBool {
	return s.TwoFaEnabled
}

// SetTwoFaAuthRequired sets the value of TwoFaAuthRequired.
func (s *TwoFAStatusResponse) SetTwoFaAuthRequired(val OptTwoFaAuthRequiredDTO) {
	s.TwoFaAuthRequired = val
}

// SetTwoFaEnabled sets the value of TwoFaEnabled.
func (s *TwoFAStatusResponse) SetTwoFaEnabled(val OptNilBool) {
	s.TwoFaEnabled = val
}

// Ref: #/components/schemas/TwoFaAuthRequiredDTO
type TwoFaAuthRequiredDTO struct {
	Login  OptNilBool `json:"login"`
	Wallet OptNilBool `json:"wallet"`
}

// GetLogin returns the value of Login.
func (s *TwoFaAuthRequiredDTO) GetLogin() OptNilBool {
	return s.Login
}

// GetWallet returns the value of Wallet.
func (s *TwoFaAuthRequiredDTO) GetWallet() OptNilBool {
	return s.Wallet
}

// SetLogin sets the value of Login.
func (s *TwoFaAuthRequiredDTO) SetLogin(val OptNilBool) {
	s.Login = val
}

// SetWallet sets the value of Wallet.
func (s *TwoFaAuthRequiredDTO) SetWallet(val OptNilBool) {
	s.Wallet = val
}

// Ref: #/components/schemas/TwoStepAuthEnabledResponse
type TwoStepAuthEnabledResponse struct {
	RecoveryCodes OptNilStringArray `json:"recovery_codes"`
}

// GetRecoveryCodes returns the value of RecoveryCodes.
func (s *TwoStepAuthEnabledResponse) GetRecoveryCodes() OptNilStringArray {
	return s.RecoveryCodes
}

// SetRecoveryCodes sets the value of RecoveryCodes.
func (s *TwoStepAuthEnabledResponse) SetRecoveryCodes(val OptNilStringArray) {
	s.RecoveryCodes = val
}

// Ref: #/components/schemas/TwoStepAuthRequestInfoResponse
type TwoStepAuthRequestInfoResponse struct {
	QrCode OptNilString `json:"qr_code"`
	Secret OptNilString `json:"secret"`
}

// GetQrCode returns the value of QrCode.
func (s *TwoStepAuthRequestInfoResponse) GetQrCode() OptNilString {
	return s.QrCode
}

// GetSecret returns the value of Secret.
func (s *TwoStepAuthRequestInfoResponse) GetSecret() OptNilString {
	return s.Secret
}

// SetQrCode sets the value of QrCode.
func (s *TwoStepAuthRequestInfoResponse) SetQrCode(val OptNilString) {
	s.QrCode = val
}

// SetSecret sets the value of Secret.
func (s *TwoStepAuthRequestInfoResponse) SetSecret(val OptNilString) {
	s.Secret = val
}

// UnbanGroupUserNoContent is response for UnbanGroupUser operation.
type UnbanGroupUserNoContent struct{}

// UnblockUserNoContent is response for UnblockUser operation.
type UnblockUserNoContent struct{}

// UnfollowUserNoContent is response for UnfollowUser operation.
type UnfollowUserNoContent struct{}

// UnhideChatsNoContent is response for UnhideChats operation.
type UnhideChatsNoContent struct{}

// UnhideUsersNoContent is response for UnhideUsers operation.
type UnhideUsersNoContent struct{}

// UnlikePostNoContent is response for UnlikePost operation.
type UnlikePostNoContent struct{}

// UnmuteGroupUserNoContent is response for UnmuteGroupUser operation.
type UnmuteGroupUserNoContent struct{}

// UnpinChatRoomNoContent is response for UnpinChatRoom operation.
type UnpinChatRoomNoContent struct{}

// UnpinGroupHighlightPostNoContent is response for UnpinGroupHighlightPost operation.
type UnpinGroupHighlightPostNoContent struct{}

// UnpinGroupNoContent is response for UnpinGroup operation.
type UnpinGroupNoContent struct{}

// UnpinGroupPostNoContent is response for UnpinGroupPost operation.
type UnpinGroupPostNoContent struct{}

// UnpinReviewNoContent is response for UnpinReview operation.
type UnpinReviewNoContent struct{}

// Ref: #/components/schemas/UnreadStatusResponse
type UnreadStatusResponse struct {
	IsUnread OptNilBool `json:"is_unread"`
}

// GetIsUnread returns the value of IsUnread.
func (s *UnreadStatusResponse) GetIsUnread() OptNilBool {
	return s.IsUnread
}

// SetIsUnread sets the value of IsUnread.
func (s *UnreadStatusResponse) SetIsUnread(val OptNilBool) {
	s.IsUnread = val
}

// UpdateAdditionalNotificationSettingNoContent is response for UpdateAdditionalNotificationSetting operation.
type UpdateAdditionalNotificationSettingNoContent struct{}

type UpdateAdditionalNotificationSettingReq struct {
	Mode OptString `json:"mode"`
	On   OptInt32  `json:"on"`
}

// GetMode returns the value of Mode.
func (s *UpdateAdditionalNotificationSettingReq) GetMode() OptString {
	return s.Mode
}

// GetOn returns the value of On.
func (s *UpdateAdditionalNotificationSettingReq) GetOn() OptInt32 {
	return s.On
}

// SetMode sets the value of Mode.
func (s *UpdateAdditionalNotificationSettingReq) SetMode(val OptString) {
	s.Mode = val
}

// SetOn sets the value of On.
func (s *UpdateAdditionalNotificationSettingReq) SetOn(val OptInt32) {
	s.On = val
}

// UpdateCallNoContent is response for UpdateCall operation.
type UpdateCallNoContent struct{}

type UpdateCallReq struct {
	CategoryID OptNilString `json:"category_id"`
	GameTitle  OptNilString `json:"game_title"`
	JoinableBy OptString    `json:"joinable_by"`
}

// GetCategoryID returns the value of CategoryID.
func (s *UpdateCallReq) GetCategoryID() OptNilString {
	return s.CategoryID
}

// GetGameTitle returns the value of GameTitle.
func (s *UpdateCallReq) GetGameTitle() OptNilString {
	return s.GameTitle
}

// GetJoinableBy returns the value of JoinableBy.
func (s *UpdateCallReq) GetJoinableBy() OptString {
	return s.JoinableBy
}

// SetCategoryID sets the value of CategoryID.
func (s *UpdateCallReq) SetCategoryID(val OptNilString) {
	s.CategoryID = val
}

// SetGameTitle sets the value of GameTitle.
func (s *UpdateCallReq) SetGameTitle(val OptNilString) {
	s.GameTitle = val
}

// SetJoinableBy sets the value of JoinableBy.
func (s *UpdateCallReq) SetJoinableBy(val OptString) {
	s.JoinableBy = val
}

// UpdateCallUserNoContent is response for UpdateCallUser operation.
type UpdateCallUserNoContent struct{}

type UpdateCallUserReq struct {
	Role OptString `json:"role"`
}

// GetRole returns the value of Role.
func (s *UpdateCallUserReq) GetRole() OptString {
	return s.Role
}

// SetRole sets the value of Role.
func (s *UpdateCallUserReq) SetRole(val OptString) {
	s.Role = val
}

// UpdateChatRoomNoContent is response for UpdateChatRoom operation.
type UpdateChatRoomNoContent struct{}

type UpdateChatRoomNotificationSettingsReq struct {
	NotificationChat OptInt32 `json:"notification_chat"`
}

// GetNotificationChat returns the value of NotificationChat.
func (s *UpdateChatRoomNotificationSettingsReq) GetNotificationChat() OptInt32 {
	return s.NotificationChat
}

// SetNotificationChat sets the value of NotificationChat.
func (s *UpdateChatRoomNotificationSettingsReq) SetNotificationChat(val OptInt32) {
	s.NotificationChat = val
}

type UpdateChatRoomReq struct {
	BackgroundFilename OptNilString `json:"background_filename"`
	IconFilename       OptNilString `json:"icon_filename"`
	Name               OptNilString `json:"name"`
}

// GetBackgroundFilename returns the value of BackgroundFilename.
func (s *UpdateChatRoomReq) GetBackgroundFilename() OptNilString {
	return s.BackgroundFilename
}

// GetIconFilename returns the value of IconFilename.
func (s *UpdateChatRoomReq) GetIconFilename() OptNilString {
	return s.IconFilename
}

// GetName returns the value of Name.
func (s *UpdateChatRoomReq) GetName() OptNilString {
	return s.Name
}

// SetBackgroundFilename sets the value of BackgroundFilename.
func (s *UpdateChatRoomReq) SetBackgroundFilename(val OptNilString) {
	s.BackgroundFilename = val
}

// SetIconFilename sets the value of IconFilename.
func (s *UpdateChatRoomReq) SetIconFilename(val OptNilString) {
	s.IconFilename = val
}

// SetName sets the value of Name.
func (s *UpdateChatRoomReq) SetName(val OptNilString) {
	s.Name = val
}

type UpdateGroupNotificationSettingsReq struct {
	NotificationGroupJoin          OptNilInt32 `json:"notification_group_join"`
	NotificationGroupMessageTagAll OptNilInt32 `json:"notification_group_message_tag_all"`
	NotificationGroupPost          OptNilInt32 `json:"notification_group_post"`
	NotificationGroupRequest       OptNilInt32 `json:"notification_group_request"`
}

// GetNotificationGroupJoin returns the value of NotificationGroupJoin.
func (s *UpdateGroupNotificationSettingsReq) GetNotificationGroupJoin() OptNilInt32 {
	return s.NotificationGroupJoin
}

// GetNotificationGroupMessageTagAll returns the value of NotificationGroupMessageTagAll.
func (s *UpdateGroupNotificationSettingsReq) GetNotificationGroupMessageTagAll() OptNilInt32 {
	return s.NotificationGroupMessageTagAll
}

// GetNotificationGroupPost returns the value of NotificationGroupPost.
func (s *UpdateGroupNotificationSettingsReq) GetNotificationGroupPost() OptNilInt32 {
	return s.NotificationGroupPost
}

// GetNotificationGroupRequest returns the value of NotificationGroupRequest.
func (s *UpdateGroupNotificationSettingsReq) GetNotificationGroupRequest() OptNilInt32 {
	return s.NotificationGroupRequest
}

// SetNotificationGroupJoin sets the value of NotificationGroupJoin.
func (s *UpdateGroupNotificationSettingsReq) SetNotificationGroupJoin(val OptNilInt32) {
	s.NotificationGroupJoin = val
}

// SetNotificationGroupMessageTagAll sets the value of NotificationGroupMessageTagAll.
func (s *UpdateGroupNotificationSettingsReq) SetNotificationGroupMessageTagAll(val OptNilInt32) {
	s.NotificationGroupMessageTagAll = val
}

// SetNotificationGroupPost sets the value of NotificationGroupPost.
func (s *UpdateGroupNotificationSettingsReq) SetNotificationGroupPost(val OptNilInt32) {
	s.NotificationGroupPost = val
}

// SetNotificationGroupRequest sets the value of NotificationGroupRequest.
func (s *UpdateGroupNotificationSettingsReq) SetNotificationGroupRequest(val OptNilInt32) {
	s.NotificationGroupRequest = val
}

type UpdateGroupReq struct {
	AllowMembersToPostImageAndVideo OptNilBool   `json:"allow_members_to_post_image_and_video"`
	AllowMembersToPostURL           OptNilBool   `json:"allow_members_to_post_url"`
	AllowOwnershipTransfer          OptNilBool   `json:"allow_ownership_transfer"`
	AllowThreadCreationBy           OptNilString `json:"allow_thread_creation_by"`
	APIKey                          OptString    `json:"api_key"`
	CallTimelineDisplay             OptNilBool   `json:"call_timeline_display"`
	CoverImageFilename              OptNilString `json:"cover_image_filename"`
	Description                     OptNilString `json:"description"`
	Gender                          OptNilInt32  `json:"gender"`
	GenerationGroupsLimit           OptNilInt32  `json:"generation_groups_limit"`
	GroupCategoryID                 OptNilInt64  `json:"group_category_id"`
	GroupIconFilename               OptNilString `json:"group_icon_filename"`
	Guidelines                      OptNilString `json:"guidelines"`
	HideConferenceCall              OptNilBool   `json:"hide_conference_call"`
	HideFromGameEight               OptNilBool   `json:"hide_from_game_eight"`
	HideReportedPosts               OptNilBool   `json:"hide_reported_posts"`
	IsPrivate                       OptNilBool   `json:"is_private"`
	OnlyMobileVerified              OptNilBool   `json:"only_mobile_verified"`
	OnlyVerifiedAge                 OptNilBool   `json:"only_verified_age"`
	Secret                          OptNilBool   `json:"secret"`
	SignedInfo                      OptString    `json:"signed_info"`
	SubCategoryID                   OptNilString `json:"sub_category_id"`
	Timestamp                       OptInt64     `json:"timestamp"`
	Topic                           OptNilString `json:"topic"`
	UUID                            OptString    `json:"uuid"`
}

// GetAllowMembersToPostImageAndVideo returns the value of AllowMembersToPostImageAndVideo.
func (s *UpdateGroupReq) GetAllowMembersToPostImageAndVideo() OptNilBool {
	return s.AllowMembersToPostImageAndVideo
}

// GetAllowMembersToPostURL returns the value of AllowMembersToPostURL.
func (s *UpdateGroupReq) GetAllowMembersToPostURL() OptNilBool {
	return s.AllowMembersToPostURL
}

// GetAllowOwnershipTransfer returns the value of AllowOwnershipTransfer.
func (s *UpdateGroupReq) GetAllowOwnershipTransfer() OptNilBool {
	return s.AllowOwnershipTransfer
}

// GetAllowThreadCreationBy returns the value of AllowThreadCreationBy.
func (s *UpdateGroupReq) GetAllowThreadCreationBy() OptNilString {
	return s.AllowThreadCreationBy
}

// GetAPIKey returns the value of APIKey.
func (s *UpdateGroupReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetCallTimelineDisplay returns the value of CallTimelineDisplay.
func (s *UpdateGroupReq) GetCallTimelineDisplay() OptNilBool {
	return s.CallTimelineDisplay
}

// GetCoverImageFilename returns the value of CoverImageFilename.
func (s *UpdateGroupReq) GetCoverImageFilename() OptNilString {
	return s.CoverImageFilename
}

// GetDescription returns the value of Description.
func (s *UpdateGroupReq) GetDescription() OptNilString {
	return s.Description
}

// GetGender returns the value of Gender.
func (s *UpdateGroupReq) GetGender() OptNilInt32 {
	return s.Gender
}

// GetGenerationGroupsLimit returns the value of GenerationGroupsLimit.
func (s *UpdateGroupReq) GetGenerationGroupsLimit() OptNilInt32 {
	return s.GenerationGroupsLimit
}

// GetGroupCategoryID returns the value of GroupCategoryID.
func (s *UpdateGroupReq) GetGroupCategoryID() OptNilInt64 {
	return s.GroupCategoryID
}

// GetGroupIconFilename returns the value of GroupIconFilename.
func (s *UpdateGroupReq) GetGroupIconFilename() OptNilString {
	return s.GroupIconFilename
}

// GetGuidelines returns the value of Guidelines.
func (s *UpdateGroupReq) GetGuidelines() OptNilString {
	return s.Guidelines
}

// GetHideConferenceCall returns the value of HideConferenceCall.
func (s *UpdateGroupReq) GetHideConferenceCall() OptNilBool {
	return s.HideConferenceCall
}

// GetHideFromGameEight returns the value of HideFromGameEight.
func (s *UpdateGroupReq) GetHideFromGameEight() OptNilBool {
	return s.HideFromGameEight
}

// GetHideReportedPosts returns the value of HideReportedPosts.
func (s *UpdateGroupReq) GetHideReportedPosts() OptNilBool {
	return s.HideReportedPosts
}

// GetIsPrivate returns the value of IsPrivate.
func (s *UpdateGroupReq) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetOnlyMobileVerified returns the value of OnlyMobileVerified.
func (s *UpdateGroupReq) GetOnlyMobileVerified() OptNilBool {
	return s.OnlyMobileVerified
}

// GetOnlyVerifiedAge returns the value of OnlyVerifiedAge.
func (s *UpdateGroupReq) GetOnlyVerifiedAge() OptNilBool {
	return s.OnlyVerifiedAge
}

// GetSecret returns the value of Secret.
func (s *UpdateGroupReq) GetSecret() OptNilBool {
	return s.Secret
}

// GetSignedInfo returns the value of SignedInfo.
func (s *UpdateGroupReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetSubCategoryID returns the value of SubCategoryID.
func (s *UpdateGroupReq) GetSubCategoryID() OptNilString {
	return s.SubCategoryID
}

// GetTimestamp returns the value of Timestamp.
func (s *UpdateGroupReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetTopic returns the value of Topic.
func (s *UpdateGroupReq) GetTopic() OptNilString {
	return s.Topic
}

// GetUUID returns the value of UUID.
func (s *UpdateGroupReq) GetUUID() OptString {
	return s.UUID
}

// SetAllowMembersToPostImageAndVideo sets the value of AllowMembersToPostImageAndVideo.
func (s *UpdateGroupReq) SetAllowMembersToPostImageAndVideo(val OptNilBool) {
	s.AllowMembersToPostImageAndVideo = val
}

// SetAllowMembersToPostURL sets the value of AllowMembersToPostURL.
func (s *UpdateGroupReq) SetAllowMembersToPostURL(val OptNilBool) {
	s.AllowMembersToPostURL = val
}

// SetAllowOwnershipTransfer sets the value of AllowOwnershipTransfer.
func (s *UpdateGroupReq) SetAllowOwnershipTransfer(val OptNilBool) {
	s.AllowOwnershipTransfer = val
}

// SetAllowThreadCreationBy sets the value of AllowThreadCreationBy.
func (s *UpdateGroupReq) SetAllowThreadCreationBy(val OptNilString) {
	s.AllowThreadCreationBy = val
}

// SetAPIKey sets the value of APIKey.
func (s *UpdateGroupReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetCallTimelineDisplay sets the value of CallTimelineDisplay.
func (s *UpdateGroupReq) SetCallTimelineDisplay(val OptNilBool) {
	s.CallTimelineDisplay = val
}

// SetCoverImageFilename sets the value of CoverImageFilename.
func (s *UpdateGroupReq) SetCoverImageFilename(val OptNilString) {
	s.CoverImageFilename = val
}

// SetDescription sets the value of Description.
func (s *UpdateGroupReq) SetDescription(val OptNilString) {
	s.Description = val
}

// SetGender sets the value of Gender.
func (s *UpdateGroupReq) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetGenerationGroupsLimit sets the value of GenerationGroupsLimit.
func (s *UpdateGroupReq) SetGenerationGroupsLimit(val OptNilInt32) {
	s.GenerationGroupsLimit = val
}

// SetGroupCategoryID sets the value of GroupCategoryID.
func (s *UpdateGroupReq) SetGroupCategoryID(val OptNilInt64) {
	s.GroupCategoryID = val
}

// SetGroupIconFilename sets the value of GroupIconFilename.
func (s *UpdateGroupReq) SetGroupIconFilename(val OptNilString) {
	s.GroupIconFilename = val
}

// SetGuidelines sets the value of Guidelines.
func (s *UpdateGroupReq) SetGuidelines(val OptNilString) {
	s.Guidelines = val
}

// SetHideConferenceCall sets the value of HideConferenceCall.
func (s *UpdateGroupReq) SetHideConferenceCall(val OptNilBool) {
	s.HideConferenceCall = val
}

// SetHideFromGameEight sets the value of HideFromGameEight.
func (s *UpdateGroupReq) SetHideFromGameEight(val OptNilBool) {
	s.HideFromGameEight = val
}

// SetHideReportedPosts sets the value of HideReportedPosts.
func (s *UpdateGroupReq) SetHideReportedPosts(val OptNilBool) {
	s.HideReportedPosts = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *UpdateGroupReq) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetOnlyMobileVerified sets the value of OnlyMobileVerified.
func (s *UpdateGroupReq) SetOnlyMobileVerified(val OptNilBool) {
	s.OnlyMobileVerified = val
}

// SetOnlyVerifiedAge sets the value of OnlyVerifiedAge.
func (s *UpdateGroupReq) SetOnlyVerifiedAge(val OptNilBool) {
	s.OnlyVerifiedAge = val
}

// SetSecret sets the value of Secret.
func (s *UpdateGroupReq) SetSecret(val OptNilBool) {
	s.Secret = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *UpdateGroupReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetSubCategoryID sets the value of SubCategoryID.
func (s *UpdateGroupReq) SetSubCategoryID(val OptNilString) {
	s.SubCategoryID = val
}

// SetTimestamp sets the value of Timestamp.
func (s *UpdateGroupReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetTopic sets the value of Topic.
func (s *UpdateGroupReq) SetTopic(val OptNilString) {
	s.Topic = val
}

// SetUUID sets the value of UUID.
func (s *UpdateGroupReq) SetUUID(val OptString) {
	s.UUID = val
}

// UpdateLanguageNoContent is response for UpdateLanguage operation.
type UpdateLanguageNoContent struct{}

type UpdateLanguageReq struct {
	APIKey     OptString `json:"api_key"`
	Language   OptString `json:"language"`
	SignedInfo OptString `json:"signed_info"`
	Timestamp  OptInt64  `json:"timestamp"`
	UUID       OptString `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *UpdateLanguageReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetLanguage returns the value of Language.
func (s *UpdateLanguageReq) GetLanguage() OptString {
	return s.Language
}

// GetSignedInfo returns the value of SignedInfo.
func (s *UpdateLanguageReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetTimestamp returns the value of Timestamp.
func (s *UpdateLanguageReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *UpdateLanguageReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *UpdateLanguageReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetLanguage sets the value of Language.
func (s *UpdateLanguageReq) SetLanguage(val OptString) {
	s.Language = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *UpdateLanguageReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetTimestamp sets the value of Timestamp.
func (s *UpdateLanguageReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *UpdateLanguageReq) SetUUID(val OptString) {
	s.UUID = val
}

type UpdateLoginReq struct {
	APIKey          OptString    `json:"api_key"`
	CurrentPassword OptNilString `json:"current_password"`
	Email           OptString    `json:"email"`
	EmailGrantToken OptNilString `json:"email_grant_token"`
	Password        OptNilString `json:"password"`
}

// GetAPIKey returns the value of APIKey.
func (s *UpdateLoginReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetCurrentPassword returns the value of CurrentPassword.
func (s *UpdateLoginReq) GetCurrentPassword() OptNilString {
	return s.CurrentPassword
}

// GetEmail returns the value of Email.
func (s *UpdateLoginReq) GetEmail() OptString {
	return s.Email
}

// GetEmailGrantToken returns the value of EmailGrantToken.
func (s *UpdateLoginReq) GetEmailGrantToken() OptNilString {
	return s.EmailGrantToken
}

// GetPassword returns the value of Password.
func (s *UpdateLoginReq) GetPassword() OptNilString {
	return s.Password
}

// SetAPIKey sets the value of APIKey.
func (s *UpdateLoginReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetCurrentPassword sets the value of CurrentPassword.
func (s *UpdateLoginReq) SetCurrentPassword(val OptNilString) {
	s.CurrentPassword = val
}

// SetEmail sets the value of Email.
func (s *UpdateLoginReq) SetEmail(val OptString) {
	s.Email = val
}

// SetEmailGrantToken sets the value of EmailGrantToken.
func (s *UpdateLoginReq) SetEmailGrantToken(val OptNilString) {
	s.EmailGrantToken = val
}

// SetPassword sets the value of Password.
func (s *UpdateLoginReq) SetPassword(val OptNilString) {
	s.Password = val
}

type UpdatePostReq struct {
	APIKey      OptString                   `json:"api_key"`
	Color       OptNilInt32                 `json:"color"`
	FontSize    OptNilInt32                 `json:"font_size"`
	Language    OptNilString                `json:"language"`
	MessageTags OptUpdatePostReqMessageTags `json:"message_tags"`
	SignedInfo  OptString                   `json:"signed_info"`
	Text        OptNilString                `json:"text"`
	Timestamp   OptInt64                    `json:"timestamp"`
	UUID        OptString                   `json:"uuid"`
}

// GetAPIKey returns the value of APIKey.
func (s *UpdatePostReq) GetAPIKey() OptString {
	return s.APIKey
}

// GetColor returns the value of Color.
func (s *UpdatePostReq) GetColor() OptNilInt32 {
	return s.Color
}

// GetFontSize returns the value of FontSize.
func (s *UpdatePostReq) GetFontSize() OptNilInt32 {
	return s.FontSize
}

// GetLanguage returns the value of Language.
func (s *UpdatePostReq) GetLanguage() OptNilString {
	return s.Language
}

// GetMessageTags returns the value of MessageTags.
func (s *UpdatePostReq) GetMessageTags() OptUpdatePostReqMessageTags {
	return s.MessageTags
}

// GetSignedInfo returns the value of SignedInfo.
func (s *UpdatePostReq) GetSignedInfo() OptString {
	return s.SignedInfo
}

// GetText returns the value of Text.
func (s *UpdatePostReq) GetText() OptNilString {
	return s.Text
}

// GetTimestamp returns the value of Timestamp.
func (s *UpdatePostReq) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// GetUUID returns the value of UUID.
func (s *UpdatePostReq) GetUUID() OptString {
	return s.UUID
}

// SetAPIKey sets the value of APIKey.
func (s *UpdatePostReq) SetAPIKey(val OptString) {
	s.APIKey = val
}

// SetColor sets the value of Color.
func (s *UpdatePostReq) SetColor(val OptNilInt32) {
	s.Color = val
}

// SetFontSize sets the value of FontSize.
func (s *UpdatePostReq) SetFontSize(val OptNilInt32) {
	s.FontSize = val
}

// SetLanguage sets the value of Language.
func (s *UpdatePostReq) SetLanguage(val OptNilString) {
	s.Language = val
}

// SetMessageTags sets the value of MessageTags.
func (s *UpdatePostReq) SetMessageTags(val OptUpdatePostReqMessageTags) {
	s.MessageTags = val
}

// SetSignedInfo sets the value of SignedInfo.
func (s *UpdatePostReq) SetSignedInfo(val OptString) {
	s.SignedInfo = val
}

// SetText sets the value of Text.
func (s *UpdatePostReq) SetText(val OptNilString) {
	s.Text = val
}

// SetTimestamp sets the value of Timestamp.
func (s *UpdatePostReq) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

// SetUUID sets the value of UUID.
func (s *UpdatePostReq) SetUUID(val OptString) {
	s.UUID = val
}

type UpdatePostReqMessageTags struct{}

// UpdateRelatedGroupsNoContent is response for UpdateRelatedGroups operation.
type UpdateRelatedGroupsNoContent struct{}

// UpdateThreadNoContent is response for UpdateThread operation.
type UpdateThreadNoContent struct{}

type UpdateThreadReq struct {
	ThreadIconFilename OptNilString `json:"thread_icon_filename"`
	Title              OptString    `json:"title"`
}

// GetThreadIconFilename returns the value of ThreadIconFilename.
func (s *UpdateThreadReq) GetThreadIconFilename() OptNilString {
	return s.ThreadIconFilename
}

// GetTitle returns the value of Title.
func (s *UpdateThreadReq) GetTitle() OptString {
	return s.Title
}

// SetThreadIconFilename sets the value of ThreadIconFilename.
func (s *UpdateThreadReq) SetThreadIconFilename(val OptNilString) {
	s.ThreadIconFilename = val
}

// SetTitle sets the value of Title.
func (s *UpdateThreadReq) SetTitle(val OptString) {
	s.Title = val
}

// UpdateUserInterestsNoContent is response for UpdateUserInterests operation.
type UpdateUserInterestsNoContent struct{}

// Ref: #/components/schemas/User
type User struct {
	Biography                           OptNilString         `json:"biography"`
	BirthDate                           OptNilString         `json:"birth_date"`
	BlockingLimit                       OptNilInt32          `json:"blocking_limit"`
	ConnectedBy                         OptUserConnectedBy   `json:"connected_by"`
	ContactPhones                       OptNilStringArray    `json:"contact_phones"`
	Country                             OptUserCountry       `json:"country"`
	CoverImage                          OptNilString         `json:"cover_image"`
	CoverImageThumbnail                 OptNilString         `json:"cover_image_thumbnail"`
	CreatedAtMillis                     OptNilInt64          `json:"created_at_millis"`
	FollowersCount                      OptNilInt32          `json:"followers_count"`
	FollowingsCount                     OptNilInt32          `json:"followings_count"`
	Gender                              OptGender            `json:"gender"`
	GiftReceivedCounter                 OptNilInt32          `json:"gift_received_counter"`
	GiftingAbility                      OptGiftingAbility    `json:"gifting_ability"`
	GroupsUsersCount                    OptNilInt32          `json:"groups_users_count"`
	ID                                  OptNilInt64          `json:"id"`
	InterestsCount                      OptNilInt32          `json:"interests_count"`
	IsAgeVerified                       OptNilBool           `json:"is_age_verified"`
	IsAppleConnected                    OptNilBool           `json:"is_apple_connected"`
	IsChatRequestOn                     OptNilBool           `json:"is_chat_request_on"`
	IsDangerous                         OptNilBool           `json:"is_dangerous"`
	IsEmailConfirmed                    OptNilBool           `json:"is_email_confirmed"`
	IsFacebookConnected                 OptNilBool           `json:"is_facebook_connected"`
	IsFollowPending                     OptNilBool           `json:"is_follow_pending"`
	IsFollowedBy                        OptNilBool           `json:"is_followed_by"`
	IsFollowedByPending                 OptNilBool           `json:"is_followed_by_pending"`
	IsFollowing                         OptNilBool           `json:"is_following"`
	IsFromDifferentGenerationAndTrusted OptNilBool           `json:"is_from_different_generation_and_trusted"`
	IsGoogleConnected                   OptNilBool           `json:"is_google_connected"`
	IsGroupPhoneOn                      OptNilBool           `json:"is_group_phone_on"`
	IsGroupVideoOn                      OptNilBool           `json:"is_group_video_on"`
	IsHidden                            OptNilBool           `json:"is_hidden"`
	IsHideVip                           OptNilBool           `json:"is_hide_vip"`
	IsInterestsSelected                 OptNilBool           `json:"is_interests_selected"`
	IsLineConnected                     OptNilBool           `json:"is_line_connected"`
	IsMuted                             OptNilBool           `json:"is_muted"`
	IsNew                               OptNilBool           `json:"is_new"`
	IsPhoneOn                           OptNilBool           `json:"is_phone_on"`
	IsPrivate                           OptNilBool           `json:"is_private"`
	IsRecentlyPenalized                 OptNilBool           `json:"is_recently_penalized"`
	IsRegistered                        OptNilBool           `json:"is_registered"`
	IsTwoFaEnabled                      OptNilBool           `json:"is_two_fa_enabled"`
	IsVideoOn                           OptNilBool           `json:"is_video_on"`
	IsVip                               OptNilBool           `json:"is_vip"`
	IsWorldIDConnected                  OptNilBool           `json:"is_world_id_connected"`
	LastLoggedInAtMillis                OptNilInt64          `json:"last_logged_in_at_millis"`
	LoginStreakCount                    OptNilInt32          `json:"login_streak_count"`
	MaskedEmail                         OptNilString         `json:"masked_email"`
	MatchingID                          OptNilInt64          `json:"matching_id"`
	MediaCount                          OptNilInt32          `json:"media_count"`
	MobileNumber                        OptNilString         `json:"mobile_number"`
	Nickname                            OptNilString         `json:"nickname"`
	OnlineStatus                        OptOnlineStatus      `json:"online_status"`
	PostsCount                          OptNilInt32          `json:"posts_count"`
	Prefecture                          OptNilString         `json:"prefecture"`
	ProfileIcon                         OptNilString         `json:"profile_icon"`
	ProfileIconFrame                    OptNilString         `json:"profile_icon_frame"`
	ProfileIconFrameThumbnail           OptNilString         `json:"profile_icon_frame_thumbnail"`
	ProfileIconThumbnail                OptNilString         `json:"profile_icon_thumbnail"`
	ReviewRestriction                   OptReviewRestriction `json:"review_restriction"`
	ReviewsCount                        OptNilInt32          `json:"reviews_count"`
	Title                               OptTitle             `json:"title"`
	TwitterID                           OptNilString         `json:"twitter_id"`
	Username                            OptNilString         `json:"username"`
	UUID                                OptNilString         `json:"uuid"`
	VipUntilMillis                      OptNilInt64          `json:"vip_until_millis"`
}

// GetBiography returns the value of Biography.
func (s *User) GetBiography() OptNilString {
	return s.Biography
}

// GetBirthDate returns the value of BirthDate.
func (s *User) GetBirthDate() OptNilString {
	return s.BirthDate
}

// GetBlockingLimit returns the value of BlockingLimit.
func (s *User) GetBlockingLimit() OptNilInt32 {
	return s.BlockingLimit
}

// GetConnectedBy returns the value of ConnectedBy.
func (s *User) GetConnectedBy() OptUserConnectedBy {
	return s.ConnectedBy
}

// GetContactPhones returns the value of ContactPhones.
func (s *User) GetContactPhones() OptNilStringArray {
	return s.ContactPhones
}

// GetCountry returns the value of Country.
func (s *User) GetCountry() OptUserCountry {
	return s.Country
}

// GetCoverImage returns the value of CoverImage.
func (s *User) GetCoverImage() OptNilString {
	return s.CoverImage
}

// GetCoverImageThumbnail returns the value of CoverImageThumbnail.
func (s *User) GetCoverImageThumbnail() OptNilString {
	return s.CoverImageThumbnail
}

// GetCreatedAtMillis returns the value of CreatedAtMillis.
func (s *User) GetCreatedAtMillis() OptNilInt64 {
	return s.CreatedAtMillis
}

// GetFollowersCount returns the value of FollowersCount.
func (s *User) GetFollowersCount() OptNilInt32 {
	return s.FollowersCount
}

// GetFollowingsCount returns the value of FollowingsCount.
func (s *User) GetFollowingsCount() OptNilInt32 {
	return s.FollowingsCount
}

// GetGender returns the value of Gender.
func (s *User) GetGender() OptGender {
	return s.Gender
}

// GetGiftReceivedCounter returns the value of GiftReceivedCounter.
func (s *User) GetGiftReceivedCounter() OptNilInt32 {
	return s.GiftReceivedCounter
}

// GetGiftingAbility returns the value of GiftingAbility.
func (s *User) GetGiftingAbility() OptGiftingAbility {
	return s.GiftingAbility
}

// GetGroupsUsersCount returns the value of GroupsUsersCount.
func (s *User) GetGroupsUsersCount() OptNilInt32 {
	return s.GroupsUsersCount
}

// GetID returns the value of ID.
func (s *User) GetID() OptNilInt64 {
	return s.ID
}

// GetInterestsCount returns the value of InterestsCount.
func (s *User) GetInterestsCount() OptNilInt32 {
	return s.InterestsCount
}

// GetIsAgeVerified returns the value of IsAgeVerified.
func (s *User) GetIsAgeVerified() OptNilBool {
	return s.IsAgeVerified
}

// GetIsAppleConnected returns the value of IsAppleConnected.
func (s *User) GetIsAppleConnected() OptNilBool {
	return s.IsAppleConnected
}

// GetIsChatRequestOn returns the value of IsChatRequestOn.
func (s *User) GetIsChatRequestOn() OptNilBool {
	return s.IsChatRequestOn
}

// GetIsDangerous returns the value of IsDangerous.
func (s *User) GetIsDangerous() OptNilBool {
	return s.IsDangerous
}

// GetIsEmailConfirmed returns the value of IsEmailConfirmed.
func (s *User) GetIsEmailConfirmed() OptNilBool {
	return s.IsEmailConfirmed
}

// GetIsFacebookConnected returns the value of IsFacebookConnected.
func (s *User) GetIsFacebookConnected() OptNilBool {
	return s.IsFacebookConnected
}

// GetIsFollowPending returns the value of IsFollowPending.
func (s *User) GetIsFollowPending() OptNilBool {
	return s.IsFollowPending
}

// GetIsFollowedBy returns the value of IsFollowedBy.
func (s *User) GetIsFollowedBy() OptNilBool {
	return s.IsFollowedBy
}

// GetIsFollowedByPending returns the value of IsFollowedByPending.
func (s *User) GetIsFollowedByPending() OptNilBool {
	return s.IsFollowedByPending
}

// GetIsFollowing returns the value of IsFollowing.
func (s *User) GetIsFollowing() OptNilBool {
	return s.IsFollowing
}

// GetIsFromDifferentGenerationAndTrusted returns the value of IsFromDifferentGenerationAndTrusted.
func (s *User) GetIsFromDifferentGenerationAndTrusted() OptNilBool {
	return s.IsFromDifferentGenerationAndTrusted
}

// GetIsGoogleConnected returns the value of IsGoogleConnected.
func (s *User) GetIsGoogleConnected() OptNilBool {
	return s.IsGoogleConnected
}

// GetIsGroupPhoneOn returns the value of IsGroupPhoneOn.
func (s *User) GetIsGroupPhoneOn() OptNilBool {
	return s.IsGroupPhoneOn
}

// GetIsGroupVideoOn returns the value of IsGroupVideoOn.
func (s *User) GetIsGroupVideoOn() OptNilBool {
	return s.IsGroupVideoOn
}

// GetIsHidden returns the value of IsHidden.
func (s *User) GetIsHidden() OptNilBool {
	return s.IsHidden
}

// GetIsHideVip returns the value of IsHideVip.
func (s *User) GetIsHideVip() OptNilBool {
	return s.IsHideVip
}

// GetIsInterestsSelected returns the value of IsInterestsSelected.
func (s *User) GetIsInterestsSelected() OptNilBool {
	return s.IsInterestsSelected
}

// GetIsLineConnected returns the value of IsLineConnected.
func (s *User) GetIsLineConnected() OptNilBool {
	return s.IsLineConnected
}

// GetIsMuted returns the value of IsMuted.
func (s *User) GetIsMuted() OptNilBool {
	return s.IsMuted
}

// GetIsNew returns the value of IsNew.
func (s *User) GetIsNew() OptNilBool {
	return s.IsNew
}

// GetIsPhoneOn returns the value of IsPhoneOn.
func (s *User) GetIsPhoneOn() OptNilBool {
	return s.IsPhoneOn
}

// GetIsPrivate returns the value of IsPrivate.
func (s *User) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetIsRecentlyPenalized returns the value of IsRecentlyPenalized.
func (s *User) GetIsRecentlyPenalized() OptNilBool {
	return s.IsRecentlyPenalized
}

// GetIsRegistered returns the value of IsRegistered.
func (s *User) GetIsRegistered() OptNilBool {
	return s.IsRegistered
}

// GetIsTwoFaEnabled returns the value of IsTwoFaEnabled.
func (s *User) GetIsTwoFaEnabled() OptNilBool {
	return s.IsTwoFaEnabled
}

// GetIsVideoOn returns the value of IsVideoOn.
func (s *User) GetIsVideoOn() OptNilBool {
	return s.IsVideoOn
}

// GetIsVip returns the value of IsVip.
func (s *User) GetIsVip() OptNilBool {
	return s.IsVip
}

// GetIsWorldIDConnected returns the value of IsWorldIDConnected.
func (s *User) GetIsWorldIDConnected() OptNilBool {
	return s.IsWorldIDConnected
}

// GetLastLoggedInAtMillis returns the value of LastLoggedInAtMillis.
func (s *User) GetLastLoggedInAtMillis() OptNilInt64 {
	return s.LastLoggedInAtMillis
}

// GetLoginStreakCount returns the value of LoginStreakCount.
func (s *User) GetLoginStreakCount() OptNilInt32 {
	return s.LoginStreakCount
}

// GetMaskedEmail returns the value of MaskedEmail.
func (s *User) GetMaskedEmail() OptNilString {
	return s.MaskedEmail
}

// GetMatchingID returns the value of MatchingID.
func (s *User) GetMatchingID() OptNilInt64 {
	return s.MatchingID
}

// GetMediaCount returns the value of MediaCount.
func (s *User) GetMediaCount() OptNilInt32 {
	return s.MediaCount
}

// GetMobileNumber returns the value of MobileNumber.
func (s *User) GetMobileNumber() OptNilString {
	return s.MobileNumber
}

// GetNickname returns the value of Nickname.
func (s *User) GetNickname() OptNilString {
	return s.Nickname
}

// GetOnlineStatus returns the value of OnlineStatus.
func (s *User) GetOnlineStatus() OptOnlineStatus {
	return s.OnlineStatus
}

// GetPostsCount returns the value of PostsCount.
func (s *User) GetPostsCount() OptNilInt32 {
	return s.PostsCount
}

// GetPrefecture returns the value of Prefecture.
func (s *User) GetPrefecture() OptNilString {
	return s.Prefecture
}

// GetProfileIcon returns the value of ProfileIcon.
func (s *User) GetProfileIcon() OptNilString {
	return s.ProfileIcon
}

// GetProfileIconFrame returns the value of ProfileIconFrame.
func (s *User) GetProfileIconFrame() OptNilString {
	return s.ProfileIconFrame
}

// GetProfileIconFrameThumbnail returns the value of ProfileIconFrameThumbnail.
func (s *User) GetProfileIconFrameThumbnail() OptNilString {
	return s.ProfileIconFrameThumbnail
}

// GetProfileIconThumbnail returns the value of ProfileIconThumbnail.
func (s *User) GetProfileIconThumbnail() OptNilString {
	return s.ProfileIconThumbnail
}

// GetReviewRestriction returns the value of ReviewRestriction.
func (s *User) GetReviewRestriction() OptReviewRestriction {
	return s.ReviewRestriction
}

// GetReviewsCount returns the value of ReviewsCount.
func (s *User) GetReviewsCount() OptNilInt32 {
	return s.ReviewsCount
}

// GetTitle returns the value of Title.
func (s *User) GetTitle() OptTitle {
	return s.Title
}

// GetTwitterID returns the value of TwitterID.
func (s *User) GetTwitterID() OptNilString {
	return s.TwitterID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() OptNilString {
	return s.Username
}

// GetUUID returns the value of UUID.
func (s *User) GetUUID() OptNilString {
	return s.UUID
}

// GetVipUntilMillis returns the value of VipUntilMillis.
func (s *User) GetVipUntilMillis() OptNilInt64 {
	return s.VipUntilMillis
}

// SetBiography sets the value of Biography.
func (s *User) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetBirthDate sets the value of BirthDate.
func (s *User) SetBirthDate(val OptNilString) {
	s.BirthDate = val
}

// SetBlockingLimit sets the value of BlockingLimit.
func (s *User) SetBlockingLimit(val OptNilInt32) {
	s.BlockingLimit = val
}

// SetConnectedBy sets the value of ConnectedBy.
func (s *User) SetConnectedBy(val OptUserConnectedBy) {
	s.ConnectedBy = val
}

// SetContactPhones sets the value of ContactPhones.
func (s *User) SetContactPhones(val OptNilStringArray) {
	s.ContactPhones = val
}

// SetCountry sets the value of Country.
func (s *User) SetCountry(val OptUserCountry) {
	s.Country = val
}

// SetCoverImage sets the value of CoverImage.
func (s *User) SetCoverImage(val OptNilString) {
	s.CoverImage = val
}

// SetCoverImageThumbnail sets the value of CoverImageThumbnail.
func (s *User) SetCoverImageThumbnail(val OptNilString) {
	s.CoverImageThumbnail = val
}

// SetCreatedAtMillis sets the value of CreatedAtMillis.
func (s *User) SetCreatedAtMillis(val OptNilInt64) {
	s.CreatedAtMillis = val
}

// SetFollowersCount sets the value of FollowersCount.
func (s *User) SetFollowersCount(val OptNilInt32) {
	s.FollowersCount = val
}

// SetFollowingsCount sets the value of FollowingsCount.
func (s *User) SetFollowingsCount(val OptNilInt32) {
	s.FollowingsCount = val
}

// SetGender sets the value of Gender.
func (s *User) SetGender(val OptGender) {
	s.Gender = val
}

// SetGiftReceivedCounter sets the value of GiftReceivedCounter.
func (s *User) SetGiftReceivedCounter(val OptNilInt32) {
	s.GiftReceivedCounter = val
}

// SetGiftingAbility sets the value of GiftingAbility.
func (s *User) SetGiftingAbility(val OptGiftingAbility) {
	s.GiftingAbility = val
}

// SetGroupsUsersCount sets the value of GroupsUsersCount.
func (s *User) SetGroupsUsersCount(val OptNilInt32) {
	s.GroupsUsersCount = val
}

// SetID sets the value of ID.
func (s *User) SetID(val OptNilInt64) {
	s.ID = val
}

// SetInterestsCount sets the value of InterestsCount.
func (s *User) SetInterestsCount(val OptNilInt32) {
	s.InterestsCount = val
}

// SetIsAgeVerified sets the value of IsAgeVerified.
func (s *User) SetIsAgeVerified(val OptNilBool) {
	s.IsAgeVerified = val
}

// SetIsAppleConnected sets the value of IsAppleConnected.
func (s *User) SetIsAppleConnected(val OptNilBool) {
	s.IsAppleConnected = val
}

// SetIsChatRequestOn sets the value of IsChatRequestOn.
func (s *User) SetIsChatRequestOn(val OptNilBool) {
	s.IsChatRequestOn = val
}

// SetIsDangerous sets the value of IsDangerous.
func (s *User) SetIsDangerous(val OptNilBool) {
	s.IsDangerous = val
}

// SetIsEmailConfirmed sets the value of IsEmailConfirmed.
func (s *User) SetIsEmailConfirmed(val OptNilBool) {
	s.IsEmailConfirmed = val
}

// SetIsFacebookConnected sets the value of IsFacebookConnected.
func (s *User) SetIsFacebookConnected(val OptNilBool) {
	s.IsFacebookConnected = val
}

// SetIsFollowPending sets the value of IsFollowPending.
func (s *User) SetIsFollowPending(val OptNilBool) {
	s.IsFollowPending = val
}

// SetIsFollowedBy sets the value of IsFollowedBy.
func (s *User) SetIsFollowedBy(val OptNilBool) {
	s.IsFollowedBy = val
}

// SetIsFollowedByPending sets the value of IsFollowedByPending.
func (s *User) SetIsFollowedByPending(val OptNilBool) {
	s.IsFollowedByPending = val
}

// SetIsFollowing sets the value of IsFollowing.
func (s *User) SetIsFollowing(val OptNilBool) {
	s.IsFollowing = val
}

// SetIsFromDifferentGenerationAndTrusted sets the value of IsFromDifferentGenerationAndTrusted.
func (s *User) SetIsFromDifferentGenerationAndTrusted(val OptNilBool) {
	s.IsFromDifferentGenerationAndTrusted = val
}

// SetIsGoogleConnected sets the value of IsGoogleConnected.
func (s *User) SetIsGoogleConnected(val OptNilBool) {
	s.IsGoogleConnected = val
}

// SetIsGroupPhoneOn sets the value of IsGroupPhoneOn.
func (s *User) SetIsGroupPhoneOn(val OptNilBool) {
	s.IsGroupPhoneOn = val
}

// SetIsGroupVideoOn sets the value of IsGroupVideoOn.
func (s *User) SetIsGroupVideoOn(val OptNilBool) {
	s.IsGroupVideoOn = val
}

// SetIsHidden sets the value of IsHidden.
func (s *User) SetIsHidden(val OptNilBool) {
	s.IsHidden = val
}

// SetIsHideVip sets the value of IsHideVip.
func (s *User) SetIsHideVip(val OptNilBool) {
	s.IsHideVip = val
}

// SetIsInterestsSelected sets the value of IsInterestsSelected.
func (s *User) SetIsInterestsSelected(val OptNilBool) {
	s.IsInterestsSelected = val
}

// SetIsLineConnected sets the value of IsLineConnected.
func (s *User) SetIsLineConnected(val OptNilBool) {
	s.IsLineConnected = val
}

// SetIsMuted sets the value of IsMuted.
func (s *User) SetIsMuted(val OptNilBool) {
	s.IsMuted = val
}

// SetIsNew sets the value of IsNew.
func (s *User) SetIsNew(val OptNilBool) {
	s.IsNew = val
}

// SetIsPhoneOn sets the value of IsPhoneOn.
func (s *User) SetIsPhoneOn(val OptNilBool) {
	s.IsPhoneOn = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *User) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetIsRecentlyPenalized sets the value of IsRecentlyPenalized.
func (s *User) SetIsRecentlyPenalized(val OptNilBool) {
	s.IsRecentlyPenalized = val
}

// SetIsRegistered sets the value of IsRegistered.
func (s *User) SetIsRegistered(val OptNilBool) {
	s.IsRegistered = val
}

// SetIsTwoFaEnabled sets the value of IsTwoFaEnabled.
func (s *User) SetIsTwoFaEnabled(val OptNilBool) {
	s.IsTwoFaEnabled = val
}

// SetIsVideoOn sets the value of IsVideoOn.
func (s *User) SetIsVideoOn(val OptNilBool) {
	s.IsVideoOn = val
}

// SetIsVip sets the value of IsVip.
func (s *User) SetIsVip(val OptNilBool) {
	s.IsVip = val
}

// SetIsWorldIDConnected sets the value of IsWorldIDConnected.
func (s *User) SetIsWorldIDConnected(val OptNilBool) {
	s.IsWorldIDConnected = val
}

// SetLastLoggedInAtMillis sets the value of LastLoggedInAtMillis.
func (s *User) SetLastLoggedInAtMillis(val OptNilInt64) {
	s.LastLoggedInAtMillis = val
}

// SetLoginStreakCount sets the value of LoginStreakCount.
func (s *User) SetLoginStreakCount(val OptNilInt32) {
	s.LoginStreakCount = val
}

// SetMaskedEmail sets the value of MaskedEmail.
func (s *User) SetMaskedEmail(val OptNilString) {
	s.MaskedEmail = val
}

// SetMatchingID sets the value of MatchingID.
func (s *User) SetMatchingID(val OptNilInt64) {
	s.MatchingID = val
}

// SetMediaCount sets the value of MediaCount.
func (s *User) SetMediaCount(val OptNilInt32) {
	s.MediaCount = val
}

// SetMobileNumber sets the value of MobileNumber.
func (s *User) SetMobileNumber(val OptNilString) {
	s.MobileNumber = val
}

// SetNickname sets the value of Nickname.
func (s *User) SetNickname(val OptNilString) {
	s.Nickname = val
}

// SetOnlineStatus sets the value of OnlineStatus.
func (s *User) SetOnlineStatus(val OptOnlineStatus) {
	s.OnlineStatus = val
}

// SetPostsCount sets the value of PostsCount.
func (s *User) SetPostsCount(val OptNilInt32) {
	s.PostsCount = val
}

// SetPrefecture sets the value of Prefecture.
func (s *User) SetPrefecture(val OptNilString) {
	s.Prefecture = val
}

// SetProfileIcon sets the value of ProfileIcon.
func (s *User) SetProfileIcon(val OptNilString) {
	s.ProfileIcon = val
}

// SetProfileIconFrame sets the value of ProfileIconFrame.
func (s *User) SetProfileIconFrame(val OptNilString) {
	s.ProfileIconFrame = val
}

// SetProfileIconFrameThumbnail sets the value of ProfileIconFrameThumbnail.
func (s *User) SetProfileIconFrameThumbnail(val OptNilString) {
	s.ProfileIconFrameThumbnail = val
}

// SetProfileIconThumbnail sets the value of ProfileIconThumbnail.
func (s *User) SetProfileIconThumbnail(val OptNilString) {
	s.ProfileIconThumbnail = val
}

// SetReviewRestriction sets the value of ReviewRestriction.
func (s *User) SetReviewRestriction(val OptReviewRestriction) {
	s.ReviewRestriction = val
}

// SetReviewsCount sets the value of ReviewsCount.
func (s *User) SetReviewsCount(val OptNilInt32) {
	s.ReviewsCount = val
}

// SetTitle sets the value of Title.
func (s *User) SetTitle(val OptTitle) {
	s.Title = val
}

// SetTwitterID sets the value of TwitterID.
func (s *User) SetTwitterID(val OptNilString) {
	s.TwitterID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val OptNilString) {
	s.Username = val
}

// SetUUID sets the value of UUID.
func (s *User) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetVipUntilMillis sets the value of VipUntilMillis.
func (s *User) SetVipUntilMillis(val OptNilInt64) {
	s.VipUntilMillis = val
}

type UserConnectedBy struct{}

type UserCountry struct{}

// Ref: #/components/schemas/UserCustomDefinitionsResponse
type UserCustomDefinitionsResponse struct {
	Age             OptNilInt32  `json:"age"`
	CreatedAt       OptNilInt64  `json:"created_at"`
	FollowersCount  OptNilInt32  `json:"followers_count"`
	FollowingsCount OptNilInt32  `json:"followings_count"`
	LastLoggedinAt  OptNilInt64  `json:"last_loggedin_at"`
	ReportedCount   OptNilInt32  `json:"reported_count"`
	Status          OptNilString `json:"status"`
}

// GetAge returns the value of Age.
func (s *UserCustomDefinitionsResponse) GetAge() OptNilInt32 {
	return s.Age
}

// GetCreatedAt returns the value of CreatedAt.
func (s *UserCustomDefinitionsResponse) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetFollowersCount returns the value of FollowersCount.
func (s *UserCustomDefinitionsResponse) GetFollowersCount() OptNilInt32 {
	return s.FollowersCount
}

// GetFollowingsCount returns the value of FollowingsCount.
func (s *UserCustomDefinitionsResponse) GetFollowingsCount() OptNilInt32 {
	return s.FollowingsCount
}

// GetLastLoggedinAt returns the value of LastLoggedinAt.
func (s *UserCustomDefinitionsResponse) GetLastLoggedinAt() OptNilInt64 {
	return s.LastLoggedinAt
}

// GetReportedCount returns the value of ReportedCount.
func (s *UserCustomDefinitionsResponse) GetReportedCount() OptNilInt32 {
	return s.ReportedCount
}

// GetStatus returns the value of Status.
func (s *UserCustomDefinitionsResponse) GetStatus() OptNilString {
	return s.Status
}

// SetAge sets the value of Age.
func (s *UserCustomDefinitionsResponse) SetAge(val OptNilInt32) {
	s.Age = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *UserCustomDefinitionsResponse) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetFollowersCount sets the value of FollowersCount.
func (s *UserCustomDefinitionsResponse) SetFollowersCount(val OptNilInt32) {
	s.FollowersCount = val
}

// SetFollowingsCount sets the value of FollowingsCount.
func (s *UserCustomDefinitionsResponse) SetFollowingsCount(val OptNilInt32) {
	s.FollowingsCount = val
}

// SetLastLoggedinAt sets the value of LastLoggedinAt.
func (s *UserCustomDefinitionsResponse) SetLastLoggedinAt(val OptNilInt64) {
	s.LastLoggedinAt = val
}

// SetReportedCount sets the value of ReportedCount.
func (s *UserCustomDefinitionsResponse) SetReportedCount(val OptNilInt32) {
	s.ReportedCount = val
}

// SetStatus sets the value of Status.
func (s *UserCustomDefinitionsResponse) SetStatus(val OptNilString) {
	s.Status = val
}

// Ref: #/components/schemas/UserInterestsResponse
type UserInterestsResponse struct {
	Interests OptNilInterestArray `json:"interests"`
}

// GetInterests returns the value of Interests.
func (s *UserInterestsResponse) GetInterests() OptNilInterestArray {
	return s.Interests
}

// SetInterests sets the value of Interests.
func (s *UserInterestsResponse) SetInterests(val OptNilInterestArray) {
	s.Interests = val
}

// Ref: #/components/schemas/UserResponse
type UserResponse struct {
	AppleConnected    OptNilBool             `json:"apple_connected"`
	BirthDate         OptNilString           `json:"birth_date"`
	BlockingLimit     OptNilInt32            `json:"blocking_limit"`
	EmailConfirmed    OptNilBool             `json:"email_confirmed"`
	FacebookConnected OptNilBool             `json:"facebook_connected"`
	GiftingAbility    OptRealmGiftingAbility `json:"gifting_ability"`
	GoogleConnected   OptNilBool             `json:"google_connected"`
	GroupPhoneOn      OptNilBool             `json:"group_phone_on"`
	GroupVideoOn      OptNilBool             `json:"group_video_on"`
	InterestsCount    OptNilInt32            `json:"interests_count"`
	LineConnected     OptNilBool             `json:"line_connected"`
	MaskedEmail       OptNilString           `json:"masked_email"`
	PhoneOn           OptNilBool             `json:"phone_on"`
	PushNotification  OptNilBool             `json:"push_notification"`
	TwitterID         OptNilString           `json:"twitter_id"`
	User              OptRealmUser           `json:"user"`
	UUID              OptNilString           `json:"uuid"`
	VideoOn           OptNilBool             `json:"video_on"`
	VipUntilSeconds   OptNilInt64            `json:"vip_until_seconds"`
	WorldIDConnected  OptNilBool             `json:"world_id_connected"`
}

// GetAppleConnected returns the value of AppleConnected.
func (s *UserResponse) GetAppleConnected() OptNilBool {
	return s.AppleConnected
}

// GetBirthDate returns the value of BirthDate.
func (s *UserResponse) GetBirthDate() OptNilString {
	return s.BirthDate
}

// GetBlockingLimit returns the value of BlockingLimit.
func (s *UserResponse) GetBlockingLimit() OptNilInt32 {
	return s.BlockingLimit
}

// GetEmailConfirmed returns the value of EmailConfirmed.
func (s *UserResponse) GetEmailConfirmed() OptNilBool {
	return s.EmailConfirmed
}

// GetFacebookConnected returns the value of FacebookConnected.
func (s *UserResponse) GetFacebookConnected() OptNilBool {
	return s.FacebookConnected
}

// GetGiftingAbility returns the value of GiftingAbility.
func (s *UserResponse) GetGiftingAbility() OptRealmGiftingAbility {
	return s.GiftingAbility
}

// GetGoogleConnected returns the value of GoogleConnected.
func (s *UserResponse) GetGoogleConnected() OptNilBool {
	return s.GoogleConnected
}

// GetGroupPhoneOn returns the value of GroupPhoneOn.
func (s *UserResponse) GetGroupPhoneOn() OptNilBool {
	return s.GroupPhoneOn
}

// GetGroupVideoOn returns the value of GroupVideoOn.
func (s *UserResponse) GetGroupVideoOn() OptNilBool {
	return s.GroupVideoOn
}

// GetInterestsCount returns the value of InterestsCount.
func (s *UserResponse) GetInterestsCount() OptNilInt32 {
	return s.InterestsCount
}

// GetLineConnected returns the value of LineConnected.
func (s *UserResponse) GetLineConnected() OptNilBool {
	return s.LineConnected
}

// GetMaskedEmail returns the value of MaskedEmail.
func (s *UserResponse) GetMaskedEmail() OptNilString {
	return s.MaskedEmail
}

// GetPhoneOn returns the value of PhoneOn.
func (s *UserResponse) GetPhoneOn() OptNilBool {
	return s.PhoneOn
}

// GetPushNotification returns the value of PushNotification.
func (s *UserResponse) GetPushNotification() OptNilBool {
	return s.PushNotification
}

// GetTwitterID returns the value of TwitterID.
func (s *UserResponse) GetTwitterID() OptNilString {
	return s.TwitterID
}

// GetUser returns the value of User.
func (s *UserResponse) GetUser() OptRealmUser {
	return s.User
}

// GetUUID returns the value of UUID.
func (s *UserResponse) GetUUID() OptNilString {
	return s.UUID
}

// GetVideoOn returns the value of VideoOn.
func (s *UserResponse) GetVideoOn() OptNilBool {
	return s.VideoOn
}

// GetVipUntilSeconds returns the value of VipUntilSeconds.
func (s *UserResponse) GetVipUntilSeconds() OptNilInt64 {
	return s.VipUntilSeconds
}

// GetWorldIDConnected returns the value of WorldIDConnected.
func (s *UserResponse) GetWorldIDConnected() OptNilBool {
	return s.WorldIDConnected
}

// SetAppleConnected sets the value of AppleConnected.
func (s *UserResponse) SetAppleConnected(val OptNilBool) {
	s.AppleConnected = val
}

// SetBirthDate sets the value of BirthDate.
func (s *UserResponse) SetBirthDate(val OptNilString) {
	s.BirthDate = val
}

// SetBlockingLimit sets the value of BlockingLimit.
func (s *UserResponse) SetBlockingLimit(val OptNilInt32) {
	s.BlockingLimit = val
}

// SetEmailConfirmed sets the value of EmailConfirmed.
func (s *UserResponse) SetEmailConfirmed(val OptNilBool) {
	s.EmailConfirmed = val
}

// SetFacebookConnected sets the value of FacebookConnected.
func (s *UserResponse) SetFacebookConnected(val OptNilBool) {
	s.FacebookConnected = val
}

// SetGiftingAbility sets the value of GiftingAbility.
func (s *UserResponse) SetGiftingAbility(val OptRealmGiftingAbility) {
	s.GiftingAbility = val
}

// SetGoogleConnected sets the value of GoogleConnected.
func (s *UserResponse) SetGoogleConnected(val OptNilBool) {
	s.GoogleConnected = val
}

// SetGroupPhoneOn sets the value of GroupPhoneOn.
func (s *UserResponse) SetGroupPhoneOn(val OptNilBool) {
	s.GroupPhoneOn = val
}

// SetGroupVideoOn sets the value of GroupVideoOn.
func (s *UserResponse) SetGroupVideoOn(val OptNilBool) {
	s.GroupVideoOn = val
}

// SetInterestsCount sets the value of InterestsCount.
func (s *UserResponse) SetInterestsCount(val OptNilInt32) {
	s.InterestsCount = val
}

// SetLineConnected sets the value of LineConnected.
func (s *UserResponse) SetLineConnected(val OptNilBool) {
	s.LineConnected = val
}

// SetMaskedEmail sets the value of MaskedEmail.
func (s *UserResponse) SetMaskedEmail(val OptNilString) {
	s.MaskedEmail = val
}

// SetPhoneOn sets the value of PhoneOn.
func (s *UserResponse) SetPhoneOn(val OptNilBool) {
	s.PhoneOn = val
}

// SetPushNotification sets the value of PushNotification.
func (s *UserResponse) SetPushNotification(val OptNilBool) {
	s.PushNotification = val
}

// SetTwitterID sets the value of TwitterID.
func (s *UserResponse) SetTwitterID(val OptNilString) {
	s.TwitterID = val
}

// SetUser sets the value of User.
func (s *UserResponse) SetUser(val OptRealmUser) {
	s.User = val
}

// SetUUID sets the value of UUID.
func (s *UserResponse) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetVideoOn sets the value of VideoOn.
func (s *UserResponse) SetVideoOn(val OptNilBool) {
	s.VideoOn = val
}

// SetVipUntilSeconds sets the value of VipUntilSeconds.
func (s *UserResponse) SetVipUntilSeconds(val OptNilInt64) {
	s.VipUntilSeconds = val
}

// SetWorldIDConnected sets the value of WorldIDConnected.
func (s *UserResponse) SetWorldIDConnected(val OptNilBool) {
	s.WorldIDConnected = val
}

// Ref: #/components/schemas/UserSetting
type UserSetting struct {
	NotificationChat OptNilBool `json:"notification_chat"`
}

// GetNotificationChat returns the value of NotificationChat.
func (s *UserSetting) GetNotificationChat() OptNilBool {
	return s.NotificationChat
}

// SetNotificationChat sets the value of NotificationChat.
func (s *UserSetting) SetNotificationChat(val OptNilBool) {
	s.NotificationChat = val
}

// Ref: #/components/schemas/UserTimestampResponse
type UserTimestampResponse struct {
	Country   OptNilString `json:"country"`
	IPAddress OptNilString `json:"ip_address"`
	Time      OptNilInt64  `json:"time"`
}

// GetCountry returns the value of Country.
func (s *UserTimestampResponse) GetCountry() OptNilString {
	return s.Country
}

// GetIPAddress returns the value of IPAddress.
func (s *UserTimestampResponse) GetIPAddress() OptNilString {
	return s.IPAddress
}

// GetTime returns the value of Time.
func (s *UserTimestampResponse) GetTime() OptNilInt64 {
	return s.Time
}

// SetCountry sets the value of Country.
func (s *UserTimestampResponse) SetCountry(val OptNilString) {
	s.Country = val
}

// SetIPAddress sets the value of IPAddress.
func (s *UserTimestampResponse) SetIPAddress(val OptNilString) {
	s.IPAddress = val
}

// SetTime sets the value of Time.
func (s *UserTimestampResponse) SetTime(val OptNilInt64) {
	s.Time = val
}

// Ref: #/components/schemas/user_UserDTO
type UserUserDTO struct {
	AgeVerified          OptNilBool   `json:"age_verified"`
	Biography            OptNilString `json:"biography"`
	CoverImage           OptNilString `json:"cover_image"`
	CoverImageThumbnail  OptNilString `json:"cover_image_thumbnail"`
	CreatedAt            OptNilInt64  `json:"created_at"`
	DangerousUser        OptNilBool   `json:"dangerous_user"`
	FollowPending        OptNilBool   `json:"follow_pending"`
	FollowedBy           OptNilBool   `json:"followed_by"`
	FollowersCount       OptNilInt32  `json:"followers_count"`
	Following            OptNilBool   `json:"following"`
	FollowingsCount      OptNilInt32  `json:"followings_count"`
	Gender               OptNilInt32  `json:"gender"`
	GroupsUsersCount     OptNilInt32  `json:"groups_users_count"`
	Hidden               OptNilBool   `json:"hidden"`
	HideVip              OptNilBool   `json:"hide_vip"`
	ID                   OptNilInt64  `json:"id"`
	IsPrivate            OptNilBool   `json:"is_private"`
	NewUser              OptNilBool   `json:"new_user"`
	Nickname             OptNilString `json:"nickname"`
	PostsCount           OptNilInt32  `json:"posts_count"`
	Prefecture           OptNilString `json:"prefecture"`
	ProfileIcon          OptNilString `json:"profile_icon"`
	ProfileIconThumbnail OptNilString `json:"profile_icon_thumbnail"`
	RecentlyKenta        OptNilBool   `json:"recently_kenta"`
	ReviewsCount         OptNilInt32  `json:"reviews_count"`
	Title                OptNilString `json:"title"`
	Vip                  OptNilBool   `json:"vip"`
}

// GetAgeVerified returns the value of AgeVerified.
func (s *UserUserDTO) GetAgeVerified() OptNilBool {
	return s.AgeVerified
}

// GetBiography returns the value of Biography.
func (s *UserUserDTO) GetBiography() OptNilString {
	return s.Biography
}

// GetCoverImage returns the value of CoverImage.
func (s *UserUserDTO) GetCoverImage() OptNilString {
	return s.CoverImage
}

// GetCoverImageThumbnail returns the value of CoverImageThumbnail.
func (s *UserUserDTO) GetCoverImageThumbnail() OptNilString {
	return s.CoverImageThumbnail
}

// GetCreatedAt returns the value of CreatedAt.
func (s *UserUserDTO) GetCreatedAt() OptNilInt64 {
	return s.CreatedAt
}

// GetDangerousUser returns the value of DangerousUser.
func (s *UserUserDTO) GetDangerousUser() OptNilBool {
	return s.DangerousUser
}

// GetFollowPending returns the value of FollowPending.
func (s *UserUserDTO) GetFollowPending() OptNilBool {
	return s.FollowPending
}

// GetFollowedBy returns the value of FollowedBy.
func (s *UserUserDTO) GetFollowedBy() OptNilBool {
	return s.FollowedBy
}

// GetFollowersCount returns the value of FollowersCount.
func (s *UserUserDTO) GetFollowersCount() OptNilInt32 {
	return s.FollowersCount
}

// GetFollowing returns the value of Following.
func (s *UserUserDTO) GetFollowing() OptNilBool {
	return s.Following
}

// GetFollowingsCount returns the value of FollowingsCount.
func (s *UserUserDTO) GetFollowingsCount() OptNilInt32 {
	return s.FollowingsCount
}

// GetGender returns the value of Gender.
func (s *UserUserDTO) GetGender() OptNilInt32 {
	return s.Gender
}

// GetGroupsUsersCount returns the value of GroupsUsersCount.
func (s *UserUserDTO) GetGroupsUsersCount() OptNilInt32 {
	return s.GroupsUsersCount
}

// GetHidden returns the value of Hidden.
func (s *UserUserDTO) GetHidden() OptNilBool {
	return s.Hidden
}

// GetHideVip returns the value of HideVip.
func (s *UserUserDTO) GetHideVip() OptNilBool {
	return s.HideVip
}

// GetID returns the value of ID.
func (s *UserUserDTO) GetID() OptNilInt64 {
	return s.ID
}

// GetIsPrivate returns the value of IsPrivate.
func (s *UserUserDTO) GetIsPrivate() OptNilBool {
	return s.IsPrivate
}

// GetNewUser returns the value of NewUser.
func (s *UserUserDTO) GetNewUser() OptNilBool {
	return s.NewUser
}

// GetNickname returns the value of Nickname.
func (s *UserUserDTO) GetNickname() OptNilString {
	return s.Nickname
}

// GetPostsCount returns the value of PostsCount.
func (s *UserUserDTO) GetPostsCount() OptNilInt32 {
	return s.PostsCount
}

// GetPrefecture returns the value of Prefecture.
func (s *UserUserDTO) GetPrefecture() OptNilString {
	return s.Prefecture
}

// GetProfileIcon returns the value of ProfileIcon.
func (s *UserUserDTO) GetProfileIcon() OptNilString {
	return s.ProfileIcon
}

// GetProfileIconThumbnail returns the value of ProfileIconThumbnail.
func (s *UserUserDTO) GetProfileIconThumbnail() OptNilString {
	return s.ProfileIconThumbnail
}

// GetRecentlyKenta returns the value of RecentlyKenta.
func (s *UserUserDTO) GetRecentlyKenta() OptNilBool {
	return s.RecentlyKenta
}

// GetReviewsCount returns the value of ReviewsCount.
func (s *UserUserDTO) GetReviewsCount() OptNilInt32 {
	return s.ReviewsCount
}

// GetTitle returns the value of Title.
func (s *UserUserDTO) GetTitle() OptNilString {
	return s.Title
}

// GetVip returns the value of Vip.
func (s *UserUserDTO) GetVip() OptNilBool {
	return s.Vip
}

// SetAgeVerified sets the value of AgeVerified.
func (s *UserUserDTO) SetAgeVerified(val OptNilBool) {
	s.AgeVerified = val
}

// SetBiography sets the value of Biography.
func (s *UserUserDTO) SetBiography(val OptNilString) {
	s.Biography = val
}

// SetCoverImage sets the value of CoverImage.
func (s *UserUserDTO) SetCoverImage(val OptNilString) {
	s.CoverImage = val
}

// SetCoverImageThumbnail sets the value of CoverImageThumbnail.
func (s *UserUserDTO) SetCoverImageThumbnail(val OptNilString) {
	s.CoverImageThumbnail = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *UserUserDTO) SetCreatedAt(val OptNilInt64) {
	s.CreatedAt = val
}

// SetDangerousUser sets the value of DangerousUser.
func (s *UserUserDTO) SetDangerousUser(val OptNilBool) {
	s.DangerousUser = val
}

// SetFollowPending sets the value of FollowPending.
func (s *UserUserDTO) SetFollowPending(val OptNilBool) {
	s.FollowPending = val
}

// SetFollowedBy sets the value of FollowedBy.
func (s *UserUserDTO) SetFollowedBy(val OptNilBool) {
	s.FollowedBy = val
}

// SetFollowersCount sets the value of FollowersCount.
func (s *UserUserDTO) SetFollowersCount(val OptNilInt32) {
	s.FollowersCount = val
}

// SetFollowing sets the value of Following.
func (s *UserUserDTO) SetFollowing(val OptNilBool) {
	s.Following = val
}

// SetFollowingsCount sets the value of FollowingsCount.
func (s *UserUserDTO) SetFollowingsCount(val OptNilInt32) {
	s.FollowingsCount = val
}

// SetGender sets the value of Gender.
func (s *UserUserDTO) SetGender(val OptNilInt32) {
	s.Gender = val
}

// SetGroupsUsersCount sets the value of GroupsUsersCount.
func (s *UserUserDTO) SetGroupsUsersCount(val OptNilInt32) {
	s.GroupsUsersCount = val
}

// SetHidden sets the value of Hidden.
func (s *UserUserDTO) SetHidden(val OptNilBool) {
	s.Hidden = val
}

// SetHideVip sets the value of HideVip.
func (s *UserUserDTO) SetHideVip(val OptNilBool) {
	s.HideVip = val
}

// SetID sets the value of ID.
func (s *UserUserDTO) SetID(val OptNilInt64) {
	s.ID = val
}

// SetIsPrivate sets the value of IsPrivate.
func (s *UserUserDTO) SetIsPrivate(val OptNilBool) {
	s.IsPrivate = val
}

// SetNewUser sets the value of NewUser.
func (s *UserUserDTO) SetNewUser(val OptNilBool) {
	s.NewUser = val
}

// SetNickname sets the value of Nickname.
func (s *UserUserDTO) SetNickname(val OptNilString) {
	s.Nickname = val
}

// SetPostsCount sets the value of PostsCount.
func (s *UserUserDTO) SetPostsCount(val OptNilInt32) {
	s.PostsCount = val
}

// SetPrefecture sets the value of Prefecture.
func (s *UserUserDTO) SetPrefecture(val OptNilString) {
	s.Prefecture = val
}

// SetProfileIcon sets the value of ProfileIcon.
func (s *UserUserDTO) SetProfileIcon(val OptNilString) {
	s.ProfileIcon = val
}

// SetProfileIconThumbnail sets the value of ProfileIconThumbnail.
func (s *UserUserDTO) SetProfileIconThumbnail(val OptNilString) {
	s.ProfileIconThumbnail = val
}

// SetRecentlyKenta sets the value of RecentlyKenta.
func (s *UserUserDTO) SetRecentlyKenta(val OptNilBool) {
	s.RecentlyKenta = val
}

// SetReviewsCount sets the value of ReviewsCount.
func (s *UserUserDTO) SetReviewsCount(val OptNilInt32) {
	s.ReviewsCount = val
}

// SetTitle sets the value of Title.
func (s *UserUserDTO) SetTitle(val OptNilString) {
	s.Title = val
}

// SetVip sets the value of Vip.
func (s *UserUserDTO) SetVip(val OptNilBool) {
	s.Vip = val
}

// Ref: #/components/schemas/UserWrapper
type UserWrapper struct {
	ID   OptNilInt64  `json:"id"`
	User OptRealmUser `json:"user"`
}

// GetID returns the value of ID.
func (s *UserWrapper) GetID() OptNilInt64 {
	return s.ID
}

// GetUser returns the value of User.
func (s *UserWrapper) GetUser() OptRealmUser {
	return s.User
}

// SetID sets the value of ID.
func (s *UserWrapper) SetID(val OptNilInt64) {
	s.ID = val
}

// SetUser sets the value of User.
func (s *UserWrapper) SetUser(val OptRealmUser) {
	s.User = val
}

// Ref: #/components/schemas/UsersByTimestampResponse
type UsersByTimestampResponse struct {
	LastTimestamp OptNilInt64          `json:"last_timestamp"`
	Users         OptNilRealmUserArray `json:"users"`
}

// GetLastTimestamp returns the value of LastTimestamp.
func (s *UsersByTimestampResponse) GetLastTimestamp() OptNilInt64 {
	return s.LastTimestamp
}

// GetUsers returns the value of Users.
func (s *UsersByTimestampResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetLastTimestamp sets the value of LastTimestamp.
func (s *UsersByTimestampResponse) SetLastTimestamp(val OptNilInt64) {
	s.LastTimestamp = val
}

// SetUsers sets the value of Users.
func (s *UsersByTimestampResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// Ref: #/components/schemas/UsersResponse
type UsersResponse struct {
	NextPageValue OptNilString         `json:"next_page_value"`
	Users         OptNilRealmUserArray `json:"users"`
}

// GetNextPageValue returns the value of NextPageValue.
func (s *UsersResponse) GetNextPageValue() OptNilString {
	return s.NextPageValue
}

// GetUsers returns the value of Users.
func (s *UsersResponse) GetUsers() OptNilRealmUserArray {
	return s.Users
}

// SetNextPageValue sets the value of NextPageValue.
func (s *UsersResponse) SetNextPageValue(val OptNilString) {
	s.NextPageValue = val
}

// SetUsers sets the value of Users.
func (s *UsersResponse) SetUsers(val OptNilRealmUserArray) {
	s.Users = val
}

// ValidateCallActionSignatureNoContent is response for ValidateCallActionSignature operation.
type ValidateCallActionSignatureNoContent struct{}

type ValidatePostReq struct {
	GroupID  OptNilInt64 `json:"group_id"`
	Text     OptString   `json:"text"`
	ThreadID OptNilInt64 `json:"thread_id"`
}

// GetGroupID returns the value of GroupID.
func (s *ValidatePostReq) GetGroupID() OptNilInt64 {
	return s.GroupID
}

// GetText returns the value of Text.
func (s *ValidatePostReq) GetText() OptString {
	return s.Text
}

// GetThreadID returns the value of ThreadID.
func (s *ValidatePostReq) GetThreadID() OptNilInt64 {
	return s.ThreadID
}

// SetGroupID sets the value of GroupID.
func (s *ValidatePostReq) SetGroupID(val OptNilInt64) {
	s.GroupID = val
}

// SetText sets the value of Text.
func (s *ValidatePostReq) SetText(val OptString) {
	s.Text = val
}

// SetThreadID sets the value of ThreadID.
func (s *ValidatePostReq) SetThreadID(val OptNilInt64) {
	s.ThreadID = val
}

// Ref: #/components/schemas/ValidationPostResponse
type ValidationPostResponse struct {
	IsAllowToPost OptNilBool `json:"is_allow_to_post"`
}

// GetIsAllowToPost returns the value of IsAllowToPost.
func (s *ValidationPostResponse) GetIsAllowToPost() OptNilBool {
	return s.IsAllowToPost
}

// SetIsAllowToPost sets the value of IsAllowToPost.
func (s *ValidationPostResponse) SetIsAllowToPost(val OptNilBool) {
	s.IsAllowToPost = val
}

// Ref: #/components/schemas/Video
type Video struct {
	Bitrate         OptNilInt32  `json:"bitrate"`
	Completed       OptNilBool   `json:"completed"`
	Height          OptNilInt32  `json:"height"`
	ID              OptNilInt64  `json:"id"`
	ThumbnailBigURL OptNilString `json:"thumbnail_big_url"`
	ThumbnailURL    OptNilString `json:"thumbnail_url"`
	VideoURL        OptNilString `json:"video_url"`
	ViewsCount      OptNilInt32  `json:"views_count"`
	Width           OptNilInt32  `json:"width"`
}

// GetBitrate returns the value of Bitrate.
func (s *Video) GetBitrate() OptNilInt32 {
	return s.Bitrate
}

// GetCompleted returns the value of Completed.
func (s *Video) GetCompleted() OptNilBool {
	return s.Completed
}

// GetHeight returns the value of Height.
func (s *Video) GetHeight() OptNilInt32 {
	return s.Height
}

// GetID returns the value of ID.
func (s *Video) GetID() OptNilInt64 {
	return s.ID
}

// GetThumbnailBigURL returns the value of ThumbnailBigURL.
func (s *Video) GetThumbnailBigURL() OptNilString {
	return s.ThumbnailBigURL
}

// GetThumbnailURL returns the value of ThumbnailURL.
func (s *Video) GetThumbnailURL() OptNilString {
	return s.ThumbnailURL
}

// GetVideoURL returns the value of VideoURL.
func (s *Video) GetVideoURL() OptNilString {
	return s.VideoURL
}

// GetViewsCount returns the value of ViewsCount.
func (s *Video) GetViewsCount() OptNilInt32 {
	return s.ViewsCount
}

// GetWidth returns the value of Width.
func (s *Video) GetWidth() OptNilInt32 {
	return s.Width
}

// SetBitrate sets the value of Bitrate.
func (s *Video) SetBitrate(val OptNilInt32) {
	s.Bitrate = val
}

// SetCompleted sets the value of Completed.
func (s *Video) SetCompleted(val OptNilBool) {
	s.Completed = val
}

// SetHeight sets the value of Height.
func (s *Video) SetHeight(val OptNilInt32) {
	s.Height = val
}

// SetID sets the value of ID.
func (s *Video) SetID(val OptNilInt64) {
	s.ID = val
}

// SetThumbnailBigURL sets the value of ThumbnailBigURL.
func (s *Video) SetThumbnailBigURL(val OptNilString) {
	s.ThumbnailBigURL = val
}

// SetThumbnailURL sets the value of ThumbnailURL.
func (s *Video) SetThumbnailURL(val OptNilString) {
	s.ThumbnailURL = val
}

// SetVideoURL sets the value of VideoURL.
func (s *Video) SetVideoURL(val OptNilString) {
	s.VideoURL = val
}

// SetViewsCount sets the value of ViewsCount.
func (s *Video) SetViewsCount(val OptNilInt32) {
	s.ViewsCount = val
}

// SetWidth sets the value of Width.
func (s *Video) SetWidth(val OptNilInt32) {
	s.Width = val
}

// ViewPostVideoNoContent is response for ViewPostVideo operation.
type ViewPostVideoNoContent struct{}

// VisitGroupNoContent is response for VisitGroup operation.
type VisitGroupNoContent struct{}

type VoteSurveyReq struct {
	ChoiceID OptInt64 `json:"choice_id"`
}

// GetChoiceID returns the value of ChoiceID.
func (s *VoteSurveyReq) GetChoiceID() OptInt64 {
	return s.ChoiceID
}

// SetChoiceID sets the value of ChoiceID.
func (s *VoteSurveyReq) SetChoiceID(val OptInt64) {
	s.ChoiceID = val
}

// Ref: #/components/schemas/VoteSurveyResponse
type VoteSurveyResponse struct {
	Survey OptSurvey `json:"survey"`
}

// GetSurvey returns the value of Survey.
func (s *VoteSurveyResponse) GetSurvey() OptSurvey {
	return s.Survey
}

// SetSurvey sets the value of Survey.
func (s *VoteSurveyResponse) SetSurvey(val OptSurvey) {
	s.Survey = val
}

// Ref: #/components/schemas/Walkthrough
type Walkthrough struct {
	Title OptNilString `json:"title"`
	URL   OptNilString `json:"url"`
}

// GetTitle returns the value of Title.
func (s *Walkthrough) GetTitle() OptNilString {
	return s.Title
}

// GetURL returns the value of URL.
func (s *Walkthrough) GetURL() OptNilString {
	return s.URL
}

// SetTitle sets the value of Title.
func (s *Walkthrough) SetTitle(val OptNilString) {
	s.Title = val
}

// SetURL sets the value of URL.
func (s *Walkthrough) SetURL(val OptNilString) {
	s.URL = val
}

// Ref: #/components/schemas/WebSocketTokenResponse
type WebSocketTokenResponse struct {
	Token OptNilString `json:"token"`
}

// GetToken returns the value of Token.
func (s *WebSocketTokenResponse) GetToken() OptNilString {
	return s.Token
}

// SetToken sets the value of Token.
func (s *WebSocketTokenResponse) SetToken(val OptNilString) {
	s.Token = val
}

// WithdrawGroupDeputyNoContent is response for WithdrawGroupDeputy operation.
type WithdrawGroupDeputyNoContent struct{}

// WithdrawGroupTransferNoContent is response for WithdrawGroupTransfer operation.
type WithdrawGroupTransferNoContent struct{}

type WithdrawGroupTransferReq struct {
	UserID OptInt64 `json:"user_id"`
}

// GetUserID returns the value of UserID.
func (s *WithdrawGroupTransferReq) GetUserID() OptInt64 {
	return s.UserID
}

// SetUserID sets the value of UserID.
func (s *WithdrawGroupTransferReq) SetUserID(val OptInt64) {
	s.UserID = val
}
