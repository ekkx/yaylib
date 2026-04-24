// Code generated; DO NOT EDIT.

package yaylib

import "fmt"

// ErrorCode is a Yay! API error discriminator carried in the JSON
// "error_code" field of a non-2xx response body. Use CodeOf(err) to
// extract it from an API error.
type ErrorCode int

const (
	ErrCodeUnknown                                              ErrorCode = 0
	ErrCodeInvalidParameter                                     ErrorCode = -1
	ErrCodeRegisteredUser                                       ErrorCode = -2
	ErrCodeAccessTokenExpired                                   ErrorCode = -3
	ErrCodeScreenNameAlreadyBeenTaken                           ErrorCode = -4
	ErrCodeUserNotFound                                         ErrorCode = -5
	ErrCodePostNotFound                                         ErrorCode = -6
	ErrCodeChatRoomNotFound                                     ErrorCode = -7
	ErrCodeChatMessageNotFound                                  ErrorCode = -8
	ErrCodeUserNotFoundAtChatRoom                               ErrorCode = -9
	ErrCodeUserMustBeOverTwoAtChatRoom                          ErrorCode = -10
	ErrCodeIncorrectPassword                                    ErrorCode = -11
	ErrCodeUserBlocked                                          ErrorCode = -12
	ErrCodePrivateUser                                          ErrorCode = -13
	ErrCodeApplicationNotFound                                  ErrorCode = -14
	ErrCodeBadSNSCredentials                                    ErrorCode = -15
	ErrCodeSNSAlreadyConnected                                  ErrorCode = -16
	ErrCodeCannotDisconnectSNS                                  ErrorCode = -17
	ErrCodeAccessTokenInvalid                                   ErrorCode = -18
	ErrCodeSpotNotFound                                         ErrorCode = -19
	ErrCodeUserBanned                                           ErrorCode = -20
	ErrCodeUserTemporaryBanned                                  ErrorCode = -21
	ErrCodeSchoolInfoChange                                     ErrorCode = -22
	ErrCodeCannotDeleteNewUser                                  ErrorCode = -26
	ErrCodeCaptchaRequired                                      ErrorCode = -29
	ErrCodeFailedToVerifyCaptcha                                ErrorCode = -30
	ErrCodeRequired2FA                                          ErrorCode = -31
	ErrCodeIncorrect2FA                                         ErrorCode = -32
	ErrCodeGroupIsFull                                          ErrorCode = -100
	ErrCodeBannedFromGroup                                      ErrorCode = -103
	ErrCodeInvalidCurrentPassword                               ErrorCode = -201
	ErrCodeInvalidPassword                                      ErrorCode = -202
	ErrCodeExistEmail                                           ErrorCode = -203
	ErrCodeBadEmailReputation                                   ErrorCode = -204
	ErrCodeChatRoomIsFull                                       ErrorCode = -308
	ErrCodeConferenceIsFull                                     ErrorCode = -309
	ErrCodeConferenceInactive                                   ErrorCode = -310
	ErrCodeGroupOwnerBlockedYou                                 ErrorCode = -312
	ErrCodeChatNeedMutualFollowed                               ErrorCode = -313
	ErrCodeConferenceCallIsLocked                               ErrorCode = -315
	ErrCodeConferenceCallIsForFollowersOnly                     ErrorCode = -317
	ErrCodeBannedFromCall                                       ErrorCode = -321
	ErrCodeNotCallOwner                                         ErrorCode = -322
	ErrCodeNotVipUser                                           ErrorCode = -326
	ErrCodeBlockingLimitExceeded                                ErrorCode = -331
	ErrCodeVerificationCodeWrong                                ErrorCode = -332
	ErrCodeVerificationCodeExpired                              ErrorCode = -333
	ErrCodeFollowLimitation                                     ErrorCode = -336
	ErrCodeAgeGapNotAllowed                                     ErrorCode = -338
	ErrCodeGroupOwnerOrGroupModeratorOnly                       ErrorCode = -339
	ErrCodeSnsShareRewardAlreadyBeenClaimed                     ErrorCode = -342
	ErrCodeQuotaLimitExceeded                                   ErrorCode = -343
	ErrCodeChatNeedAgeVerified                                  ErrorCode = -346
	ErrCodeOnlyAgeVerifiedUserCanJoinGroup                      ErrorCode = -347
	ErrCodeRequirePhoneVerificationToChat                       ErrorCode = -348
	ErrCodeNotPostOwner                                         ErrorCode = -350
	ErrCodeGroupGenerationNotMatched                            ErrorCode = -352
	ErrCodePhoneNumberCheckVerificationCodeSubmitQuotaExceeded  ErrorCode = -355
	ErrCodePhoneNumberCheckVerificationCodeRequestQuotaExceeded ErrorCode = -356
	ErrCodeGroupOfferHasBeenAccepted                            ErrorCode = -357
	ErrCodeGroupOfferHasBeenWithdrawn                           ErrorCode = -358
	ErrCodeIpBanned                                             ErrorCode = -360
	ErrCodeNotConnectedToTwitter                                ErrorCode = -361
	ErrCodePrivateUserTimeline                                  ErrorCode = -363
	ErrCodeCounterRefreshLimitExceeded                          ErrorCode = -364
	ErrCodeNotFollowedByOpponent                                ErrorCode = -367
	ErrCodeExceedChangeCountryQuota                             ErrorCode = -369
	ErrCodeNotGroupMember                                       ErrorCode = -370
	ErrCodeGroupPendingTransfer                                 ErrorCode = -371
	ErrCodeGroupPendingDeputization                             ErrorCode = -372
	ErrCodeUserRestrictedChatWithCautionUsers                   ErrorCode = -373
	ErrCodeRestrictedCreateChatWithNewUsers                     ErrorCode = -374
	ErrCodeRepostPostNotRepostable                              ErrorCode = -375
	ErrCodeTooManyAccountsCreated                               ErrorCode = -376
	ErrCodeOnlySpecificGenderCanJoinGroup                       ErrorCode = -377
	ErrCodeCreateSpecificGenderGroupRequiredGender              ErrorCode = -378
	ErrCodeGroupRelatedExceededNumberOfRelatedGroups            ErrorCode = -382
	ErrCodeExceededPinnedLimit                                  ErrorCode = -383
	ErrCodeGroupShareOnTwitterLimitExceeded                     ErrorCode = -384
	ErrCodeReportedContent                                      ErrorCode = -385
	ErrCodeInsufficientCoins                                    ErrorCode = -400
	ErrCodeConferenceCallIsForMutualFollowsOnly                 ErrorCode = -402
	ErrCodeExceededLimit                                        ErrorCode = -403
	ErrCodeGroupInviteExceeded                                  ErrorCode = -404
	ErrCodePhoneVerificationRequired                            ErrorCode = -405
	ErrCodeContentTooOld                                        ErrorCode = -406
	ErrCodePasswordTooShort                                     ErrorCode = -407
	ErrCodePasswordTooLong                                      ErrorCode = -408
	ErrCodePasswordNotAllowed                                   ErrorCode = -409
	ErrCodeCommonPassword                                       ErrorCode = -410
	ErrCodeUnableToMovePostToThread                             ErrorCode = -412
	ErrCodeUnableToPostUrl                                      ErrorCode = -413
	ErrCodeReferralAlreadyRegistered                            ErrorCode = -415
	ErrCodeMuteUserOverLimit                                    ErrorCode = -416
	ErrCodeTextTranslationLimitExceeded                         ErrorCode = -419
	ErrCodeTextTranslation                                      ErrorCode = -420
	ErrCodeCardExchangeLimit                                    ErrorCode = -424
	ErrCodeInvalidAppVersion                                    ErrorCode = -800
	ErrCodeDynamicErrorMessage                                  ErrorCode = -999
	ErrCodePhoneNumberBanned                                    ErrorCode = -1000
	ErrCodeBadRequest                                           ErrorCode = 400
	ErrCodeUnauthorized                                         ErrorCode = 401
	ErrCodeAccessForbidden                                      ErrorCode = 403
	ErrCodeNotFound                                             ErrorCode = 404
	ErrCodeConflict                                             ErrorCode = 409
	ErrCodeTooManyRequests                                      ErrorCode = 429
	ErrCodeInternalServerError                                  ErrorCode = 500
	ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet              ErrorCode = 4002
	ErrCodeWeb3WalletAlreadyLinkedToAnotherAccount              ErrorCode = 4003
	ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet2             ErrorCode = 4004
	ErrCodeWeb3PalLevelUpBattlesRequired                        ErrorCode = 4005
	ErrCodeWeb3PalLevelUpMaximumLevelReached                    ErrorCode = 4006
	ErrCodeWeb3PalPoolCooldown                                  ErrorCode = 4010
	ErrCodeWeb3PalPoolEmpty                                     ErrorCode = 4011
	ErrCodeWeb3PalAlreadyBattle                                 ErrorCode = 4012
	ErrCodeWeb3WalletHasPendingTransactions                     ErrorCode = 4017
	ErrCodeWeb3WalletNetworkError                               ErrorCode = 5003
	ErrCodeWeb3EMPLInsufficientFunds                            ErrorCode = 6001
	ErrCodeWeb3EMPLFeeExceedsBalance                            ErrorCode = 6002
	ErrCodeWeb3EMPLAmountInsufficientToCoverGasFee              ErrorCode = 6004
	ErrCodeWeb3EMPLInsufficientNativeBalance                    ErrorCode = 6007
	ErrCodeWeb3EMPLExchangeExceedsRemainingQuota                ErrorCode = 6008
	ErrCodeOffline                                              ErrorCode = 10000
	ErrCodeCallOpponentDisabled                                 ErrorCode = 10001
)

// String returns the canonical name of the error code (e.g. "UserNotFound"
// for ErrCodeUserNotFound). Unrecognized codes render as "ErrorCode(<n>)".
func (c ErrorCode) String() string {
	switch c {
	case ErrCodeUnknown:
		return "Unknown"
	case ErrCodeInvalidParameter:
		return "InvalidParameter"
	case ErrCodeRegisteredUser:
		return "RegisteredUser"
	case ErrCodeAccessTokenExpired:
		return "AccessTokenExpired"
	case ErrCodeScreenNameAlreadyBeenTaken:
		return "ScreenNameAlreadyBeenTaken"
	case ErrCodeUserNotFound:
		return "UserNotFound"
	case ErrCodePostNotFound:
		return "PostNotFound"
	case ErrCodeChatRoomNotFound:
		return "ChatRoomNotFound"
	case ErrCodeChatMessageNotFound:
		return "ChatMessageNotFound"
	case ErrCodeUserNotFoundAtChatRoom:
		return "UserNotFoundAtChatRoom"
	case ErrCodeUserMustBeOverTwoAtChatRoom:
		return "UserMustBeOverTwoAtChatRoom"
	case ErrCodeIncorrectPassword:
		return "IncorrectPassword"
	case ErrCodeUserBlocked:
		return "UserBlocked"
	case ErrCodePrivateUser:
		return "PrivateUser"
	case ErrCodeApplicationNotFound:
		return "ApplicationNotFound"
	case ErrCodeBadSNSCredentials:
		return "BadSNSCredentials"
	case ErrCodeSNSAlreadyConnected:
		return "SNSAlreadyConnected"
	case ErrCodeCannotDisconnectSNS:
		return "CannotDisconnectSNS"
	case ErrCodeAccessTokenInvalid:
		return "AccessTokenInvalid"
	case ErrCodeSpotNotFound:
		return "SpotNotFound"
	case ErrCodeUserBanned:
		return "UserBanned"
	case ErrCodeUserTemporaryBanned:
		return "UserTemporaryBanned"
	case ErrCodeSchoolInfoChange:
		return "SchoolInfoChange"
	case ErrCodeCannotDeleteNewUser:
		return "CannotDeleteNewUser"
	case ErrCodeCaptchaRequired:
		return "CaptchaRequired"
	case ErrCodeFailedToVerifyCaptcha:
		return "FailedToVerifyCaptcha"
	case ErrCodeRequired2FA:
		return "Required2FA"
	case ErrCodeIncorrect2FA:
		return "Incorrect2FA"
	case ErrCodeGroupIsFull:
		return "GroupIsFull"
	case ErrCodeBannedFromGroup:
		return "BannedFromGroup"
	case ErrCodeInvalidCurrentPassword:
		return "InvalidCurrentPassword"
	case ErrCodeInvalidPassword:
		return "InvalidPassword"
	case ErrCodeExistEmail:
		return "ExistEmail"
	case ErrCodeBadEmailReputation:
		return "BadEmailReputation"
	case ErrCodeChatRoomIsFull:
		return "ChatRoomIsFull"
	case ErrCodeConferenceIsFull:
		return "ConferenceIsFull"
	case ErrCodeConferenceInactive:
		return "ConferenceInactive"
	case ErrCodeGroupOwnerBlockedYou:
		return "GroupOwnerBlockedYou"
	case ErrCodeChatNeedMutualFollowed:
		return "ChatNeedMutualFollowed"
	case ErrCodeConferenceCallIsLocked:
		return "ConferenceCallIsLocked"
	case ErrCodeConferenceCallIsForFollowersOnly:
		return "ConferenceCallIsForFollowersOnly"
	case ErrCodeBannedFromCall:
		return "BannedFromCall"
	case ErrCodeNotCallOwner:
		return "NotCallOwner"
	case ErrCodeNotVipUser:
		return "NotVipUser"
	case ErrCodeBlockingLimitExceeded:
		return "BlockingLimitExceeded"
	case ErrCodeVerificationCodeWrong:
		return "VerificationCodeWrong"
	case ErrCodeVerificationCodeExpired:
		return "VerificationCodeExpired"
	case ErrCodeFollowLimitation:
		return "FollowLimitation"
	case ErrCodeAgeGapNotAllowed:
		return "AgeGapNotAllowed"
	case ErrCodeGroupOwnerOrGroupModeratorOnly:
		return "GroupOwnerOrGroupModeratorOnly"
	case ErrCodeSnsShareRewardAlreadyBeenClaimed:
		return "SnsShareRewardAlreadyBeenClaimed"
	case ErrCodeQuotaLimitExceeded:
		return "QuotaLimitExceeded"
	case ErrCodeChatNeedAgeVerified:
		return "ChatNeedAgeVerified"
	case ErrCodeOnlyAgeVerifiedUserCanJoinGroup:
		return "OnlyAgeVerifiedUserCanJoinGroup"
	case ErrCodeRequirePhoneVerificationToChat:
		return "RequirePhoneVerificationToChat"
	case ErrCodeNotPostOwner:
		return "NotPostOwner"
	case ErrCodeGroupGenerationNotMatched:
		return "GroupGenerationNotMatched"
	case ErrCodePhoneNumberCheckVerificationCodeSubmitQuotaExceeded:
		return "PhoneNumberCheckVerificationCodeSubmitQuotaExceeded"
	case ErrCodePhoneNumberCheckVerificationCodeRequestQuotaExceeded:
		return "PhoneNumberCheckVerificationCodeRequestQuotaExceeded"
	case ErrCodeGroupOfferHasBeenAccepted:
		return "GroupOfferHasBeenAccepted"
	case ErrCodeGroupOfferHasBeenWithdrawn:
		return "GroupOfferHasBeenWithdrawn"
	case ErrCodeIpBanned:
		return "IpBanned"
	case ErrCodeNotConnectedToTwitter:
		return "NotConnectedToTwitter"
	case ErrCodePrivateUserTimeline:
		return "PrivateUserTimeline"
	case ErrCodeCounterRefreshLimitExceeded:
		return "CounterRefreshLimitExceeded"
	case ErrCodeNotFollowedByOpponent:
		return "NotFollowedByOpponent"
	case ErrCodeExceedChangeCountryQuota:
		return "ExceedChangeCountryQuota"
	case ErrCodeNotGroupMember:
		return "NotGroupMember"
	case ErrCodeGroupPendingTransfer:
		return "GroupPendingTransfer"
	case ErrCodeGroupPendingDeputization:
		return "GroupPendingDeputization"
	case ErrCodeUserRestrictedChatWithCautionUsers:
		return "UserRestrictedChatWithCautionUsers"
	case ErrCodeRestrictedCreateChatWithNewUsers:
		return "RestrictedCreateChatWithNewUsers"
	case ErrCodeRepostPostNotRepostable:
		return "RepostPostNotRepostable"
	case ErrCodeTooManyAccountsCreated:
		return "TooManyAccountsCreated"
	case ErrCodeOnlySpecificGenderCanJoinGroup:
		return "OnlySpecificGenderCanJoinGroup"
	case ErrCodeCreateSpecificGenderGroupRequiredGender:
		return "CreateSpecificGenderGroupRequiredGender"
	case ErrCodeGroupRelatedExceededNumberOfRelatedGroups:
		return "GroupRelatedExceededNumberOfRelatedGroups"
	case ErrCodeExceededPinnedLimit:
		return "ExceededPinnedLimit"
	case ErrCodeGroupShareOnTwitterLimitExceeded:
		return "GroupShareOnTwitterLimitExceeded"
	case ErrCodeReportedContent:
		return "ReportedContent"
	case ErrCodeInsufficientCoins:
		return "InsufficientCoins"
	case ErrCodeConferenceCallIsForMutualFollowsOnly:
		return "ConferenceCallIsForMutualFollowsOnly"
	case ErrCodeExceededLimit:
		return "ExceededLimit"
	case ErrCodeGroupInviteExceeded:
		return "GroupInviteExceeded"
	case ErrCodePhoneVerificationRequired:
		return "PhoneVerificationRequired"
	case ErrCodeContentTooOld:
		return "ContentTooOld"
	case ErrCodePasswordTooShort:
		return "PasswordTooShort"
	case ErrCodePasswordTooLong:
		return "PasswordTooLong"
	case ErrCodePasswordNotAllowed:
		return "PasswordNotAllowed"
	case ErrCodeCommonPassword:
		return "CommonPassword"
	case ErrCodeUnableToMovePostToThread:
		return "UnableToMovePostToThread"
	case ErrCodeUnableToPostUrl:
		return "UnableToPostUrl"
	case ErrCodeReferralAlreadyRegistered:
		return "ReferralAlreadyRegistered"
	case ErrCodeMuteUserOverLimit:
		return "MuteUserOverLimit"
	case ErrCodeTextTranslationLimitExceeded:
		return "TextTranslationLimitExceeded"
	case ErrCodeTextTranslation:
		return "TextTranslation"
	case ErrCodeCardExchangeLimit:
		return "CardExchangeLimit"
	case ErrCodeInvalidAppVersion:
		return "InvalidAppVersion"
	case ErrCodeDynamicErrorMessage:
		return "DynamicErrorMessage"
	case ErrCodePhoneNumberBanned:
		return "PhoneNumberBanned"
	case ErrCodeBadRequest:
		return "BadRequest"
	case ErrCodeUnauthorized:
		return "Unauthorized"
	case ErrCodeAccessForbidden:
		return "AccessForbidden"
	case ErrCodeNotFound:
		return "NotFound"
	case ErrCodeConflict:
		return "Conflict"
	case ErrCodeTooManyRequests:
		return "TooManyRequests"
	case ErrCodeInternalServerError:
		return "InternalServerError"
	case ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet:
		return "Web3AccountAlreadyLinkedToAnotherWallet"
	case ErrCodeWeb3WalletAlreadyLinkedToAnotherAccount:
		return "Web3WalletAlreadyLinkedToAnotherAccount"
	case ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet2:
		return "Web3AccountAlreadyLinkedToAnotherWallet2"
	case ErrCodeWeb3PalLevelUpBattlesRequired:
		return "Web3PalLevelUpBattlesRequired"
	case ErrCodeWeb3PalLevelUpMaximumLevelReached:
		return "Web3PalLevelUpMaximumLevelReached"
	case ErrCodeWeb3PalPoolCooldown:
		return "Web3PalPoolCooldown"
	case ErrCodeWeb3PalPoolEmpty:
		return "Web3PalPoolEmpty"
	case ErrCodeWeb3PalAlreadyBattle:
		return "Web3PalAlreadyBattle"
	case ErrCodeWeb3WalletHasPendingTransactions:
		return "Web3WalletHasPendingTransactions"
	case ErrCodeWeb3WalletNetworkError:
		return "Web3WalletNetworkError"
	case ErrCodeWeb3EMPLInsufficientFunds:
		return "Web3EMPLInsufficientFunds"
	case ErrCodeWeb3EMPLFeeExceedsBalance:
		return "Web3EMPLFeeExceedsBalance"
	case ErrCodeWeb3EMPLAmountInsufficientToCoverGasFee:
		return "Web3EMPLAmountInsufficientToCoverGasFee"
	case ErrCodeWeb3EMPLInsufficientNativeBalance:
		return "Web3EMPLInsufficientNativeBalance"
	case ErrCodeWeb3EMPLExchangeExceedsRemainingQuota:
		return "Web3EMPLExchangeExceedsRemainingQuota"
	case ErrCodeOffline:
		return "Offline"
	case ErrCodeCallOpponentDisabled:
		return "CallOpponentDisabled"
	}
	return fmt.Sprintf("ErrorCode(%d)", int(c))
}
