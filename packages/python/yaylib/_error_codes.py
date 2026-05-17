# Code generated; DO NOT EDIT.
#
# ErrorCode mirrors the Yay! API error discriminator carried in the
# JSON "error_code" field of a non-2xx response body. Use code_of(err)
# to extract it and compare against these constants (PORTING.md §9).

__all__ = [
    "ErrCodeUnknown",
    "ErrCodeInvalidParameter",
    "ErrCodeRegisteredUser",
    "ErrCodeAccessTokenExpired",
    "ErrCodeScreenNameAlreadyBeenTaken",
    "ErrCodeUserNotFound",
    "ErrCodePostNotFound",
    "ErrCodeChatRoomNotFound",
    "ErrCodeChatMessageNotFound",
    "ErrCodeUserNotFoundAtChatRoom",
    "ErrCodeUserMustBeOverTwoAtChatRoom",
    "ErrCodeIncorrectPassword",
    "ErrCodeUserBlocked",
    "ErrCodePrivateUser",
    "ErrCodeApplicationNotFound",
    "ErrCodeBadSNSCredentials",
    "ErrCodeSNSAlreadyConnected",
    "ErrCodeCannotDisconnectSNS",
    "ErrCodeAccessTokenInvalid",
    "ErrCodeSpotNotFound",
    "ErrCodeUserBanned",
    "ErrCodeUserTemporaryBanned",
    "ErrCodeSchoolInfoChange",
    "ErrCodeCannotDeleteNewUser",
    "ErrCodeCaptchaRequired",
    "ErrCodeFailedToVerifyCaptcha",
    "ErrCodeRequired2FA",
    "ErrCodeIncorrect2FA",
    "ErrCodeGroupIsFull",
    "ErrCodeBannedFromGroup",
    "ErrCodeInvalidCurrentPassword",
    "ErrCodeInvalidPassword",
    "ErrCodeExistEmail",
    "ErrCodeBadEmailReputation",
    "ErrCodeChatRoomIsFull",
    "ErrCodeConferenceIsFull",
    "ErrCodeConferenceInactive",
    "ErrCodeGroupOwnerBlockedYou",
    "ErrCodeChatNeedMutualFollowed",
    "ErrCodeConferenceCallIsLocked",
    "ErrCodeConferenceCallIsForFollowersOnly",
    "ErrCodeBannedFromCall",
    "ErrCodeNotCallOwner",
    "ErrCodeNotVipUser",
    "ErrCodeBlockingLimitExceeded",
    "ErrCodeVerificationCodeWrong",
    "ErrCodeVerificationCodeExpired",
    "ErrCodeFollowLimitation",
    "ErrCodeAgeGapNotAllowed",
    "ErrCodeGroupOwnerOrGroupModeratorOnly",
    "ErrCodeSnsShareRewardAlreadyBeenClaimed",
    "ErrCodeQuotaLimitExceeded",
    "ErrCodeChatNeedAgeVerified",
    "ErrCodeOnlyAgeVerifiedUserCanJoinGroup",
    "ErrCodeRequirePhoneVerificationToChat",
    "ErrCodeNotPostOwner",
    "ErrCodeGroupGenerationNotMatched",
    "ErrCodePhoneNumberCheckVerificationCodeSubmitQuotaExceeded",
    "ErrCodePhoneNumberCheckVerificationCodeRequestQuotaExceeded",
    "ErrCodeGroupOfferHasBeenAccepted",
    "ErrCodeGroupOfferHasBeenWithdrawn",
    "ErrCodeIpBanned",
    "ErrCodeNotConnectedToTwitter",
    "ErrCodePrivateUserTimeline",
    "ErrCodeCounterRefreshLimitExceeded",
    "ErrCodeNotFollowedByOpponent",
    "ErrCodeExceedChangeCountryQuota",
    "ErrCodeNotGroupMember",
    "ErrCodeGroupPendingTransfer",
    "ErrCodeGroupPendingDeputization",
    "ErrCodeUserRestrictedChatWithCautionUsers",
    "ErrCodeRestrictedCreateChatWithNewUsers",
    "ErrCodeRepostPostNotRepostable",
    "ErrCodeTooManyAccountsCreated",
    "ErrCodeOnlySpecificGenderCanJoinGroup",
    "ErrCodeCreateSpecificGenderGroupRequiredGender",
    "ErrCodeGroupRelatedExceededNumberOfRelatedGroups",
    "ErrCodeExceededPinnedLimit",
    "ErrCodeGroupShareOnTwitterLimitExceeded",
    "ErrCodeReportedContent",
    "ErrCodeInsufficientCoins",
    "ErrCodeConferenceCallIsForMutualFollowsOnly",
    "ErrCodeExceededLimit",
    "ErrCodeGroupInviteExceeded",
    "ErrCodePhoneVerificationRequired",
    "ErrCodeContentTooOld",
    "ErrCodePasswordTooShort",
    "ErrCodePasswordTooLong",
    "ErrCodePasswordNotAllowed",
    "ErrCodeCommonPassword",
    "ErrCodeUnableToMovePostToThread",
    "ErrCodeUnableToPostUrl",
    "ErrCodeReferralAlreadyRegistered",
    "ErrCodeMuteUserOverLimit",
    "ErrCodeTextTranslationLimitExceeded",
    "ErrCodeTextTranslation",
    "ErrCodeCardExchangeLimit",
    "ErrCodeInvalidAppVersion",
    "ErrCodeDynamicErrorMessage",
    "ErrCodePhoneNumberBanned",
    "ErrCodeBadRequest",
    "ErrCodeUnauthorized",
    "ErrCodeAccessForbidden",
    "ErrCodeNotFound",
    "ErrCodeConflict",
    "ErrCodeTooManyRequests",
    "ErrCodeInternalServerError",
    "ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet",
    "ErrCodeWeb3WalletAlreadyLinkedToAnotherAccount",
    "ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet2",
    "ErrCodeWeb3PalLevelUpBattlesRequired",
    "ErrCodeWeb3PalLevelUpMaximumLevelReached",
    "ErrCodeWeb3PalPoolCooldown",
    "ErrCodeWeb3PalPoolEmpty",
    "ErrCodeWeb3PalAlreadyBattle",
    "ErrCodeWeb3WalletHasPendingTransactions",
    "ErrCodeWeb3WalletNetworkError",
    "ErrCodeWeb3EMPLInsufficientFunds",
    "ErrCodeWeb3EMPLFeeExceedsBalance",
    "ErrCodeWeb3EMPLAmountInsufficientToCoverGasFee",
    "ErrCodeWeb3EMPLInsufficientNativeBalance",
    "ErrCodeWeb3EMPLExchangeExceedsRemainingQuota",
    "ErrCodeOffline",
    "ErrCodeCallOpponentDisabled",
    "error_code_name",
]

ErrCodeUnknown: int = 0
ErrCodeInvalidParameter: int = -1
ErrCodeRegisteredUser: int = -2
ErrCodeAccessTokenExpired: int = -3
ErrCodeScreenNameAlreadyBeenTaken: int = -4
ErrCodeUserNotFound: int = -5
ErrCodePostNotFound: int = -6
ErrCodeChatRoomNotFound: int = -7
ErrCodeChatMessageNotFound: int = -8
ErrCodeUserNotFoundAtChatRoom: int = -9
ErrCodeUserMustBeOverTwoAtChatRoom: int = -10
ErrCodeIncorrectPassword: int = -11
ErrCodeUserBlocked: int = -12
ErrCodePrivateUser: int = -13
ErrCodeApplicationNotFound: int = -14
ErrCodeBadSNSCredentials: int = -15
ErrCodeSNSAlreadyConnected: int = -16
ErrCodeCannotDisconnectSNS: int = -17
ErrCodeAccessTokenInvalid: int = -18
ErrCodeSpotNotFound: int = -19
ErrCodeUserBanned: int = -20
ErrCodeUserTemporaryBanned: int = -21
ErrCodeSchoolInfoChange: int = -22
ErrCodeCannotDeleteNewUser: int = -26
ErrCodeCaptchaRequired: int = -29
ErrCodeFailedToVerifyCaptcha: int = -30
ErrCodeRequired2FA: int = -31
ErrCodeIncorrect2FA: int = -32
ErrCodeGroupIsFull: int = -100
ErrCodeBannedFromGroup: int = -103
ErrCodeInvalidCurrentPassword: int = -201
ErrCodeInvalidPassword: int = -202
ErrCodeExistEmail: int = -203
ErrCodeBadEmailReputation: int = -204
ErrCodeChatRoomIsFull: int = -308
ErrCodeConferenceIsFull: int = -309
ErrCodeConferenceInactive: int = -310
ErrCodeGroupOwnerBlockedYou: int = -312
ErrCodeChatNeedMutualFollowed: int = -313
ErrCodeConferenceCallIsLocked: int = -315
ErrCodeConferenceCallIsForFollowersOnly: int = -317
ErrCodeBannedFromCall: int = -321
ErrCodeNotCallOwner: int = -322
ErrCodeNotVipUser: int = -326
ErrCodeBlockingLimitExceeded: int = -331
ErrCodeVerificationCodeWrong: int = -332
ErrCodeVerificationCodeExpired: int = -333
ErrCodeFollowLimitation: int = -336
ErrCodeAgeGapNotAllowed: int = -338
ErrCodeGroupOwnerOrGroupModeratorOnly: int = -339
ErrCodeSnsShareRewardAlreadyBeenClaimed: int = -342
ErrCodeQuotaLimitExceeded: int = -343
ErrCodeChatNeedAgeVerified: int = -346
ErrCodeOnlyAgeVerifiedUserCanJoinGroup: int = -347
ErrCodeRequirePhoneVerificationToChat: int = -348
ErrCodeNotPostOwner: int = -350
ErrCodeGroupGenerationNotMatched: int = -352
ErrCodePhoneNumberCheckVerificationCodeSubmitQuotaExceeded: int = -355
ErrCodePhoneNumberCheckVerificationCodeRequestQuotaExceeded: int = -356
ErrCodeGroupOfferHasBeenAccepted: int = -357
ErrCodeGroupOfferHasBeenWithdrawn: int = -358
ErrCodeIpBanned: int = -360
ErrCodeNotConnectedToTwitter: int = -361
ErrCodePrivateUserTimeline: int = -363
ErrCodeCounterRefreshLimitExceeded: int = -364
ErrCodeNotFollowedByOpponent: int = -367
ErrCodeExceedChangeCountryQuota: int = -369
ErrCodeNotGroupMember: int = -370
ErrCodeGroupPendingTransfer: int = -371
ErrCodeGroupPendingDeputization: int = -372
ErrCodeUserRestrictedChatWithCautionUsers: int = -373
ErrCodeRestrictedCreateChatWithNewUsers: int = -374
ErrCodeRepostPostNotRepostable: int = -375
ErrCodeTooManyAccountsCreated: int = -376
ErrCodeOnlySpecificGenderCanJoinGroup: int = -377
ErrCodeCreateSpecificGenderGroupRequiredGender: int = -378
ErrCodeGroupRelatedExceededNumberOfRelatedGroups: int = -382
ErrCodeExceededPinnedLimit: int = -383
ErrCodeGroupShareOnTwitterLimitExceeded: int = -384
ErrCodeReportedContent: int = -385
ErrCodeInsufficientCoins: int = -400
ErrCodeConferenceCallIsForMutualFollowsOnly: int = -402
ErrCodeExceededLimit: int = -403
ErrCodeGroupInviteExceeded: int = -404
ErrCodePhoneVerificationRequired: int = -405
ErrCodeContentTooOld: int = -406
ErrCodePasswordTooShort: int = -407
ErrCodePasswordTooLong: int = -408
ErrCodePasswordNotAllowed: int = -409
ErrCodeCommonPassword: int = -410
ErrCodeUnableToMovePostToThread: int = -412
ErrCodeUnableToPostUrl: int = -413
ErrCodeReferralAlreadyRegistered: int = -415
ErrCodeMuteUserOverLimit: int = -416
ErrCodeTextTranslationLimitExceeded: int = -419
ErrCodeTextTranslation: int = -420
ErrCodeCardExchangeLimit: int = -424
ErrCodeInvalidAppVersion: int = -800
ErrCodeDynamicErrorMessage: int = -999
ErrCodePhoneNumberBanned: int = -1000
ErrCodeBadRequest: int = 400
ErrCodeUnauthorized: int = 401
ErrCodeAccessForbidden: int = 403
ErrCodeNotFound: int = 404
ErrCodeConflict: int = 409
ErrCodeTooManyRequests: int = 429
ErrCodeInternalServerError: int = 500
ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet: int = 4002
ErrCodeWeb3WalletAlreadyLinkedToAnotherAccount: int = 4003
ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet2: int = 4004
ErrCodeWeb3PalLevelUpBattlesRequired: int = 4005
ErrCodeWeb3PalLevelUpMaximumLevelReached: int = 4006
ErrCodeWeb3PalPoolCooldown: int = 4010
ErrCodeWeb3PalPoolEmpty: int = 4011
ErrCodeWeb3PalAlreadyBattle: int = 4012
ErrCodeWeb3WalletHasPendingTransactions: int = 4017
ErrCodeWeb3WalletNetworkError: int = 5003
ErrCodeWeb3EMPLInsufficientFunds: int = 6001
ErrCodeWeb3EMPLFeeExceedsBalance: int = 6002
ErrCodeWeb3EMPLAmountInsufficientToCoverGasFee: int = 6004
ErrCodeWeb3EMPLInsufficientNativeBalance: int = 6007
ErrCodeWeb3EMPLExchangeExceedsRemainingQuota: int = 6008
ErrCodeOffline: int = 10000
ErrCodeCallOpponentDisabled: int = 10001

_NAMES = {
    0: "Unknown",
    -1: "InvalidParameter",
    -2: "RegisteredUser",
    -3: "AccessTokenExpired",
    -4: "ScreenNameAlreadyBeenTaken",
    -5: "UserNotFound",
    -6: "PostNotFound",
    -7: "ChatRoomNotFound",
    -8: "ChatMessageNotFound",
    -9: "UserNotFoundAtChatRoom",
    -10: "UserMustBeOverTwoAtChatRoom",
    -11: "IncorrectPassword",
    -12: "UserBlocked",
    -13: "PrivateUser",
    -14: "ApplicationNotFound",
    -15: "BadSNSCredentials",
    -16: "SNSAlreadyConnected",
    -17: "CannotDisconnectSNS",
    -18: "AccessTokenInvalid",
    -19: "SpotNotFound",
    -20: "UserBanned",
    -21: "UserTemporaryBanned",
    -22: "SchoolInfoChange",
    -26: "CannotDeleteNewUser",
    -29: "CaptchaRequired",
    -30: "FailedToVerifyCaptcha",
    -31: "Required2FA",
    -32: "Incorrect2FA",
    -100: "GroupIsFull",
    -103: "BannedFromGroup",
    -201: "InvalidCurrentPassword",
    -202: "InvalidPassword",
    -203: "ExistEmail",
    -204: "BadEmailReputation",
    -308: "ChatRoomIsFull",
    -309: "ConferenceIsFull",
    -310: "ConferenceInactive",
    -312: "GroupOwnerBlockedYou",
    -313: "ChatNeedMutualFollowed",
    -315: "ConferenceCallIsLocked",
    -317: "ConferenceCallIsForFollowersOnly",
    -321: "BannedFromCall",
    -322: "NotCallOwner",
    -326: "NotVipUser",
    -331: "BlockingLimitExceeded",
    -332: "VerificationCodeWrong",
    -333: "VerificationCodeExpired",
    -336: "FollowLimitation",
    -338: "AgeGapNotAllowed",
    -339: "GroupOwnerOrGroupModeratorOnly",
    -342: "SnsShareRewardAlreadyBeenClaimed",
    -343: "QuotaLimitExceeded",
    -346: "ChatNeedAgeVerified",
    -347: "OnlyAgeVerifiedUserCanJoinGroup",
    -348: "RequirePhoneVerificationToChat",
    -350: "NotPostOwner",
    -352: "GroupGenerationNotMatched",
    -355: "PhoneNumberCheckVerificationCodeSubmitQuotaExceeded",
    -356: "PhoneNumberCheckVerificationCodeRequestQuotaExceeded",
    -357: "GroupOfferHasBeenAccepted",
    -358: "GroupOfferHasBeenWithdrawn",
    -360: "IpBanned",
    -361: "NotConnectedToTwitter",
    -363: "PrivateUserTimeline",
    -364: "CounterRefreshLimitExceeded",
    -367: "NotFollowedByOpponent",
    -369: "ExceedChangeCountryQuota",
    -370: "NotGroupMember",
    -371: "GroupPendingTransfer",
    -372: "GroupPendingDeputization",
    -373: "UserRestrictedChatWithCautionUsers",
    -374: "RestrictedCreateChatWithNewUsers",
    -375: "RepostPostNotRepostable",
    -376: "TooManyAccountsCreated",
    -377: "OnlySpecificGenderCanJoinGroup",
    -378: "CreateSpecificGenderGroupRequiredGender",
    -382: "GroupRelatedExceededNumberOfRelatedGroups",
    -383: "ExceededPinnedLimit",
    -384: "GroupShareOnTwitterLimitExceeded",
    -385: "ReportedContent",
    -400: "InsufficientCoins",
    -402: "ConferenceCallIsForMutualFollowsOnly",
    -403: "ExceededLimit",
    -404: "GroupInviteExceeded",
    -405: "PhoneVerificationRequired",
    -406: "ContentTooOld",
    -407: "PasswordTooShort",
    -408: "PasswordTooLong",
    -409: "PasswordNotAllowed",
    -410: "CommonPassword",
    -412: "UnableToMovePostToThread",
    -413: "UnableToPostUrl",
    -415: "ReferralAlreadyRegistered",
    -416: "MuteUserOverLimit",
    -419: "TextTranslationLimitExceeded",
    -420: "TextTranslation",
    -424: "CardExchangeLimit",
    -800: "InvalidAppVersion",
    -999: "DynamicErrorMessage",
    -1000: "PhoneNumberBanned",
    400: "BadRequest",
    401: "Unauthorized",
    403: "AccessForbidden",
    404: "NotFound",
    409: "Conflict",
    429: "TooManyRequests",
    500: "InternalServerError",
    4002: "Web3AccountAlreadyLinkedToAnotherWallet",
    4003: "Web3WalletAlreadyLinkedToAnotherAccount",
    4004: "Web3AccountAlreadyLinkedToAnotherWallet2",
    4005: "Web3PalLevelUpBattlesRequired",
    4006: "Web3PalLevelUpMaximumLevelReached",
    4010: "Web3PalPoolCooldown",
    4011: "Web3PalPoolEmpty",
    4012: "Web3PalAlreadyBattle",
    4017: "Web3WalletHasPendingTransactions",
    5003: "Web3WalletNetworkError",
    6001: "Web3EMPLInsufficientFunds",
    6002: "Web3EMPLFeeExceedsBalance",
    6004: "Web3EMPLAmountInsufficientToCoverGasFee",
    6007: "Web3EMPLInsufficientNativeBalance",
    6008: "Web3EMPLExchangeExceedsRemainingQuota",
    10000: "Offline",
    10001: "CallOpponentDisabled",
}


def error_code_name(c: int) -> str:
    """Canonical name (e.g. \"UserNotFound\"); an unrecognized code
    renders as \"ErrorCode(<n>)\"."""
    return _NAMES.get(c, f"ErrorCode({c})")
