// Code generated; DO NOT EDIT.
//
// ErrorCode mirrors the Yay! API error discriminator carried in the
// JSON "error_code" field of a non-2xx response body. Use codeOf(err)
// to extract it and compare against these constants for switch /
// dispatch (PORTING.md §9).

export type ErrorCode = number;

export const ErrCodeUnknown: ErrorCode = 0;
export const ErrCodeInvalidParameter: ErrorCode = -1;
export const ErrCodeRegisteredUser: ErrorCode = -2;
export const ErrCodeAccessTokenExpired: ErrorCode = -3;
export const ErrCodeScreenNameAlreadyBeenTaken: ErrorCode = -4;
export const ErrCodeUserNotFound: ErrorCode = -5;
export const ErrCodePostNotFound: ErrorCode = -6;
export const ErrCodeChatRoomNotFound: ErrorCode = -7;
export const ErrCodeChatMessageNotFound: ErrorCode = -8;
export const ErrCodeUserNotFoundAtChatRoom: ErrorCode = -9;
export const ErrCodeUserMustBeOverTwoAtChatRoom: ErrorCode = -10;
export const ErrCodeIncorrectPassword: ErrorCode = -11;
export const ErrCodeUserBlocked: ErrorCode = -12;
export const ErrCodePrivateUser: ErrorCode = -13;
export const ErrCodeApplicationNotFound: ErrorCode = -14;
export const ErrCodeBadSNSCredentials: ErrorCode = -15;
export const ErrCodeSNSAlreadyConnected: ErrorCode = -16;
export const ErrCodeCannotDisconnectSNS: ErrorCode = -17;
export const ErrCodeAccessTokenInvalid: ErrorCode = -18;
export const ErrCodeSpotNotFound: ErrorCode = -19;
export const ErrCodeUserBanned: ErrorCode = -20;
export const ErrCodeUserTemporaryBanned: ErrorCode = -21;
export const ErrCodeSchoolInfoChange: ErrorCode = -22;
export const ErrCodeCannotDeleteNewUser: ErrorCode = -26;
export const ErrCodeCaptchaRequired: ErrorCode = -29;
export const ErrCodeFailedToVerifyCaptcha: ErrorCode = -30;
export const ErrCodeRequired2FA: ErrorCode = -31;
export const ErrCodeIncorrect2FA: ErrorCode = -32;
export const ErrCodeGroupIsFull: ErrorCode = -100;
export const ErrCodeBannedFromGroup: ErrorCode = -103;
export const ErrCodeInvalidCurrentPassword: ErrorCode = -201;
export const ErrCodeInvalidPassword: ErrorCode = -202;
export const ErrCodeExistEmail: ErrorCode = -203;
export const ErrCodeBadEmailReputation: ErrorCode = -204;
export const ErrCodeChatRoomIsFull: ErrorCode = -308;
export const ErrCodeConferenceIsFull: ErrorCode = -309;
export const ErrCodeConferenceInactive: ErrorCode = -310;
export const ErrCodeGroupOwnerBlockedYou: ErrorCode = -312;
export const ErrCodeChatNeedMutualFollowed: ErrorCode = -313;
export const ErrCodeConferenceCallIsLocked: ErrorCode = -315;
export const ErrCodeConferenceCallIsForFollowersOnly: ErrorCode = -317;
export const ErrCodeBannedFromCall: ErrorCode = -321;
export const ErrCodeNotCallOwner: ErrorCode = -322;
export const ErrCodeNotVipUser: ErrorCode = -326;
export const ErrCodeBlockingLimitExceeded: ErrorCode = -331;
export const ErrCodeVerificationCodeWrong: ErrorCode = -332;
export const ErrCodeVerificationCodeExpired: ErrorCode = -333;
export const ErrCodeFollowLimitation: ErrorCode = -336;
export const ErrCodeAgeGapNotAllowed: ErrorCode = -338;
export const ErrCodeGroupOwnerOrGroupModeratorOnly: ErrorCode = -339;
export const ErrCodeSnsShareRewardAlreadyBeenClaimed: ErrorCode = -342;
export const ErrCodeQuotaLimitExceeded: ErrorCode = -343;
export const ErrCodeChatNeedAgeVerified: ErrorCode = -346;
export const ErrCodeOnlyAgeVerifiedUserCanJoinGroup: ErrorCode = -347;
export const ErrCodeRequirePhoneVerificationToChat: ErrorCode = -348;
export const ErrCodeNotPostOwner: ErrorCode = -350;
export const ErrCodeGroupGenerationNotMatched: ErrorCode = -352;
export const ErrCodePhoneNumberCheckVerificationCodeSubmitQuotaExceeded: ErrorCode = -355;
export const ErrCodePhoneNumberCheckVerificationCodeRequestQuotaExceeded: ErrorCode = -356;
export const ErrCodeGroupOfferHasBeenAccepted: ErrorCode = -357;
export const ErrCodeGroupOfferHasBeenWithdrawn: ErrorCode = -358;
export const ErrCodeIpBanned: ErrorCode = -360;
export const ErrCodeNotConnectedToTwitter: ErrorCode = -361;
export const ErrCodePrivateUserTimeline: ErrorCode = -363;
export const ErrCodeCounterRefreshLimitExceeded: ErrorCode = -364;
export const ErrCodeNotFollowedByOpponent: ErrorCode = -367;
export const ErrCodeExceedChangeCountryQuota: ErrorCode = -369;
export const ErrCodeNotGroupMember: ErrorCode = -370;
export const ErrCodeGroupPendingTransfer: ErrorCode = -371;
export const ErrCodeGroupPendingDeputization: ErrorCode = -372;
export const ErrCodeUserRestrictedChatWithCautionUsers: ErrorCode = -373;
export const ErrCodeRestrictedCreateChatWithNewUsers: ErrorCode = -374;
export const ErrCodeRepostPostNotRepostable: ErrorCode = -375;
export const ErrCodeTooManyAccountsCreated: ErrorCode = -376;
export const ErrCodeOnlySpecificGenderCanJoinGroup: ErrorCode = -377;
export const ErrCodeCreateSpecificGenderGroupRequiredGender: ErrorCode = -378;
export const ErrCodeGroupRelatedExceededNumberOfRelatedGroups: ErrorCode = -382;
export const ErrCodeExceededPinnedLimit: ErrorCode = -383;
export const ErrCodeGroupShareOnTwitterLimitExceeded: ErrorCode = -384;
export const ErrCodeReportedContent: ErrorCode = -385;
export const ErrCodeInsufficientCoins: ErrorCode = -400;
export const ErrCodeConferenceCallIsForMutualFollowsOnly: ErrorCode = -402;
export const ErrCodeExceededLimit: ErrorCode = -403;
export const ErrCodeGroupInviteExceeded: ErrorCode = -404;
export const ErrCodePhoneVerificationRequired: ErrorCode = -405;
export const ErrCodeContentTooOld: ErrorCode = -406;
export const ErrCodePasswordTooShort: ErrorCode = -407;
export const ErrCodePasswordTooLong: ErrorCode = -408;
export const ErrCodePasswordNotAllowed: ErrorCode = -409;
export const ErrCodeCommonPassword: ErrorCode = -410;
export const ErrCodeUnableToMovePostToThread: ErrorCode = -412;
export const ErrCodeUnableToPostUrl: ErrorCode = -413;
export const ErrCodeReferralAlreadyRegistered: ErrorCode = -415;
export const ErrCodeMuteUserOverLimit: ErrorCode = -416;
export const ErrCodeTextTranslationLimitExceeded: ErrorCode = -419;
export const ErrCodeTextTranslation: ErrorCode = -420;
export const ErrCodeCardExchangeLimit: ErrorCode = -424;
export const ErrCodeInvalidAppVersion: ErrorCode = -800;
export const ErrCodeDynamicErrorMessage: ErrorCode = -999;
export const ErrCodePhoneNumberBanned: ErrorCode = -1000;
export const ErrCodeBadRequest: ErrorCode = 400;
export const ErrCodeUnauthorized: ErrorCode = 401;
export const ErrCodeAccessForbidden: ErrorCode = 403;
export const ErrCodeNotFound: ErrorCode = 404;
export const ErrCodeConflict: ErrorCode = 409;
export const ErrCodeTooManyRequests: ErrorCode = 429;
export const ErrCodeInternalServerError: ErrorCode = 500;
export const ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet: ErrorCode = 4002;
export const ErrCodeWeb3WalletAlreadyLinkedToAnotherAccount: ErrorCode = 4003;
export const ErrCodeWeb3AccountAlreadyLinkedToAnotherWallet2: ErrorCode = 4004;
export const ErrCodeWeb3PalLevelUpBattlesRequired: ErrorCode = 4005;
export const ErrCodeWeb3PalLevelUpMaximumLevelReached: ErrorCode = 4006;
export const ErrCodeWeb3PalPoolCooldown: ErrorCode = 4010;
export const ErrCodeWeb3PalPoolEmpty: ErrorCode = 4011;
export const ErrCodeWeb3PalAlreadyBattle: ErrorCode = 4012;
export const ErrCodeWeb3WalletHasPendingTransactions: ErrorCode = 4017;
export const ErrCodeWeb3WalletNetworkError: ErrorCode = 5003;
export const ErrCodeWeb3EMPLInsufficientFunds: ErrorCode = 6001;
export const ErrCodeWeb3EMPLFeeExceedsBalance: ErrorCode = 6002;
export const ErrCodeWeb3EMPLAmountInsufficientToCoverGasFee: ErrorCode = 6004;
export const ErrCodeWeb3EMPLInsufficientNativeBalance: ErrorCode = 6007;
export const ErrCodeWeb3EMPLExchangeExceedsRemainingQuota: ErrorCode = 6008;
export const ErrCodeOffline: ErrorCode = 10000;
export const ErrCodeCallOpponentDisabled: ErrorCode = 10001;

const NAMES: Record<number, string> = {
  [0]: "Unknown",
  [-1]: "InvalidParameter",
  [-2]: "RegisteredUser",
  [-3]: "AccessTokenExpired",
  [-4]: "ScreenNameAlreadyBeenTaken",
  [-5]: "UserNotFound",
  [-6]: "PostNotFound",
  [-7]: "ChatRoomNotFound",
  [-8]: "ChatMessageNotFound",
  [-9]: "UserNotFoundAtChatRoom",
  [-10]: "UserMustBeOverTwoAtChatRoom",
  [-11]: "IncorrectPassword",
  [-12]: "UserBlocked",
  [-13]: "PrivateUser",
  [-14]: "ApplicationNotFound",
  [-15]: "BadSNSCredentials",
  [-16]: "SNSAlreadyConnected",
  [-17]: "CannotDisconnectSNS",
  [-18]: "AccessTokenInvalid",
  [-19]: "SpotNotFound",
  [-20]: "UserBanned",
  [-21]: "UserTemporaryBanned",
  [-22]: "SchoolInfoChange",
  [-26]: "CannotDeleteNewUser",
  [-29]: "CaptchaRequired",
  [-30]: "FailedToVerifyCaptcha",
  [-31]: "Required2FA",
  [-32]: "Incorrect2FA",
  [-100]: "GroupIsFull",
  [-103]: "BannedFromGroup",
  [-201]: "InvalidCurrentPassword",
  [-202]: "InvalidPassword",
  [-203]: "ExistEmail",
  [-204]: "BadEmailReputation",
  [-308]: "ChatRoomIsFull",
  [-309]: "ConferenceIsFull",
  [-310]: "ConferenceInactive",
  [-312]: "GroupOwnerBlockedYou",
  [-313]: "ChatNeedMutualFollowed",
  [-315]: "ConferenceCallIsLocked",
  [-317]: "ConferenceCallIsForFollowersOnly",
  [-321]: "BannedFromCall",
  [-322]: "NotCallOwner",
  [-326]: "NotVipUser",
  [-331]: "BlockingLimitExceeded",
  [-332]: "VerificationCodeWrong",
  [-333]: "VerificationCodeExpired",
  [-336]: "FollowLimitation",
  [-338]: "AgeGapNotAllowed",
  [-339]: "GroupOwnerOrGroupModeratorOnly",
  [-342]: "SnsShareRewardAlreadyBeenClaimed",
  [-343]: "QuotaLimitExceeded",
  [-346]: "ChatNeedAgeVerified",
  [-347]: "OnlyAgeVerifiedUserCanJoinGroup",
  [-348]: "RequirePhoneVerificationToChat",
  [-350]: "NotPostOwner",
  [-352]: "GroupGenerationNotMatched",
  [-355]: "PhoneNumberCheckVerificationCodeSubmitQuotaExceeded",
  [-356]: "PhoneNumberCheckVerificationCodeRequestQuotaExceeded",
  [-357]: "GroupOfferHasBeenAccepted",
  [-358]: "GroupOfferHasBeenWithdrawn",
  [-360]: "IpBanned",
  [-361]: "NotConnectedToTwitter",
  [-363]: "PrivateUserTimeline",
  [-364]: "CounterRefreshLimitExceeded",
  [-367]: "NotFollowedByOpponent",
  [-369]: "ExceedChangeCountryQuota",
  [-370]: "NotGroupMember",
  [-371]: "GroupPendingTransfer",
  [-372]: "GroupPendingDeputization",
  [-373]: "UserRestrictedChatWithCautionUsers",
  [-374]: "RestrictedCreateChatWithNewUsers",
  [-375]: "RepostPostNotRepostable",
  [-376]: "TooManyAccountsCreated",
  [-377]: "OnlySpecificGenderCanJoinGroup",
  [-378]: "CreateSpecificGenderGroupRequiredGender",
  [-382]: "GroupRelatedExceededNumberOfRelatedGroups",
  [-383]: "ExceededPinnedLimit",
  [-384]: "GroupShareOnTwitterLimitExceeded",
  [-385]: "ReportedContent",
  [-400]: "InsufficientCoins",
  [-402]: "ConferenceCallIsForMutualFollowsOnly",
  [-403]: "ExceededLimit",
  [-404]: "GroupInviteExceeded",
  [-405]: "PhoneVerificationRequired",
  [-406]: "ContentTooOld",
  [-407]: "PasswordTooShort",
  [-408]: "PasswordTooLong",
  [-409]: "PasswordNotAllowed",
  [-410]: "CommonPassword",
  [-412]: "UnableToMovePostToThread",
  [-413]: "UnableToPostUrl",
  [-415]: "ReferralAlreadyRegistered",
  [-416]: "MuteUserOverLimit",
  [-419]: "TextTranslationLimitExceeded",
  [-420]: "TextTranslation",
  [-424]: "CardExchangeLimit",
  [-800]: "InvalidAppVersion",
  [-999]: "DynamicErrorMessage",
  [-1000]: "PhoneNumberBanned",
  [400]: "BadRequest",
  [401]: "Unauthorized",
  [403]: "AccessForbidden",
  [404]: "NotFound",
  [409]: "Conflict",
  [429]: "TooManyRequests",
  [500]: "InternalServerError",
  [4002]: "Web3AccountAlreadyLinkedToAnotherWallet",
  [4003]: "Web3WalletAlreadyLinkedToAnotherAccount",
  [4004]: "Web3AccountAlreadyLinkedToAnotherWallet2",
  [4005]: "Web3PalLevelUpBattlesRequired",
  [4006]: "Web3PalLevelUpMaximumLevelReached",
  [4010]: "Web3PalPoolCooldown",
  [4011]: "Web3PalPoolEmpty",
  [4012]: "Web3PalAlreadyBattle",
  [4017]: "Web3WalletHasPendingTransactions",
  [5003]: "Web3WalletNetworkError",
  [6001]: "Web3EMPLInsufficientFunds",
  [6002]: "Web3EMPLFeeExceedsBalance",
  [6004]: "Web3EMPLAmountInsufficientToCoverGasFee",
  [6007]: "Web3EMPLInsufficientNativeBalance",
  [6008]: "Web3EMPLExchangeExceedsRemainingQuota",
  [10000]: "Offline",
  [10001]: "CallOpponentDisabled",
};

// errorCodeName returns the canonical name (e.g. "UserNotFound");
// an unrecognized code renders as "ErrorCode(<n>)".
export function errorCodeName(c: ErrorCode): string {
  return NAMES[c] ?? `ErrorCode(${c})`;
}
