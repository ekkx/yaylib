"""
MIT License

Copyright (c) 2023 ekkx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from enum import Enum


__all__ = (
    "YaylibError",
    "HTTPError",
    "BadRequestError",
    "AuthenticationError",
    "ForbiddenError",
    "NotFoundError",
    "RateLimitError",
    "YayServerError",
    "ErrorCode",
)


class YaylibError(Exception):
    """Base exception for yaylib"""

    pass


class HTTPError(YaylibError):
    """HTTPError()

    Exception raised when an HTTP request fails
    """

    pass


class BadRequestError(HTTPError):
    """BadRequestError()

    Exception raised for a 400 HTTP status code
    """

    pass


class AuthenticationError(HTTPError):
    """AuthenticationError()

    Exception raised for a 401 HTTP status code
    """

    pass


class ForbiddenError(HTTPError):
    """ForbiddenError()

    Exception raised for a 403 HTTP status code
    """

    pass


class NotFoundError(HTTPError):
    """NotFoundError()

    Exception raised for a 404 HTTP status code
    """

    pass


class RateLimitError(HTTPError):
    """RateLimitError()

    Exception raised for a 429 HTTP status code
    """

    pass


class YayServerError(HTTPError):
    """YayServerError()

    Exception raised for a 5xx HTTP status code
    """

    pass


class ErrorCode(Enum):
    Unknown = "unknown"
    InvalidParameter = -1
    RegisteredUser = -2
    AccessTokenExpired = -3
    ScreenNameAlreadyBeenTaken = -4
    UserNotFound = -5
    PostNotFound = -6
    ChatRoomNotFound = -7
    ChatMessageNotFound = -8
    UserNotFoundAtChatRoom = -9
    UserMustBeOverTwoAtChatRoom = -10
    IncorrectPassword = -11
    UserBlocked = -12
    PrivateUser = -13
    ApplicationNotFound = -14
    BadSNSCredentials = -15
    SNSAlreadyConnected = -16
    CannotDisconnectSNS = -17
    AccessTokenInvalid = -18
    SpotNotFound = -19
    UserBanned = -20
    UserTemporaryBanned = -21
    SchoolInfoChange = -22
    CannotDeleteNewUser = -26
    CaptchaRequired = -29
    FailedToVerifyCaptcha = -30
    GroupIsFull = -100
    BannedFromGroup = -103
    InvalidCurrentPassword = -200
    InvalidPassword = -201
    InvalidEmailOrPassword = -202
    ExistEmail = -203
    BadEmailReputation = -204
    ChatRoomIsFull = -308
    ConferenceIsFull = -309
    ConferenceInactive = -310
    GroupOwnerBlockedYou = -312
    ChatNeedMutualFollowed = -313
    ConferenceCallIsLocked = -315
    ConferenceCallIsForFollowersOnly = -317
    InvalidEmail = -319
    RegisteredEmail = -320
    BannedFromCall = -321
    NotCallOwner = -322
    NotVipUser = -326
    BlockingLimitExceeded = -331
    VerificationCodeWrong = -332
    VerificationCodeExpired = -333
    InvalidAppVersion = -334
    InvalidPhoneNumber = -335
    FollowLimitation = -336
    AgeGapNotAllowed = -338
    GroupOwnerOrGroupModeratorOnly = -339
    UnableToRegisterUserDueToPolicy = -340
    SnsShareRewardAlreadyBeenClaimed = -342
    QuotaLimitExceeded = -343
    ChatNeedAgeVerified = -346
    OnlyAgeVerifiedUserCanJoinGroup = -347
    RequirePhoneVerificationToChat = -348
    NotPostOwner = -350
    GroupGenerationNotMatched = -352
    PhoneNumberCheckVerificationCodeSubmitQuotaExceeded = -355
    PhoneNumberCheckVerificationCodeRequestQuotaExceeded = -356
    GroupOfferHasBeenAccepted = -357
    GroupOfferHasBeenWithdrawn = -358
    IpBanned = -360
    NotConnectedToTwitter = -361
    PrivateUserTimeline = -363
    CounterRefreshLimitExceeded = -364
    NotFollowedByOpponent = -367
    ExceedChangeCountryQuota = -369
    NotGroupMember = -370
    GroupPendingTransfer = -371
    GroupPendingDeputization = -372
    UserRestrictedChatWithCautionUsers = -373
    RestrictedCreateChatWithNewUsers = -374
    RepostPostNotRepostable = -375
    TooManyAccountsCreated = -376
    OnlySpecificGenderCanJoinGroup = -377
    CreateSpecificGenderGroupRequiredGender = -378
    GroupRelatedExceededNumberOfRelatedGroups = -382
    ExceededPinnedLimit = -383
    GroupShareOnTwitterLimitExceeded = -384
    ReportedContent = -385
    ConferenceCallIsForMutualFollowsOnly = -402
    ExceededLimit = -403
    GroupInviteExceeded = -404
    PhoneVerificationRequired = -405
    ContentTooOld = -406
    PasswordTooShort = -407
    PasswordTooLong = -408
    PasswordNotAllowed = -409
    CommonPassword = -410
    EmailNotAuthorized = -411
    UnableToMovePostToThread = -412
    UnableToPostUrl = -413
    UnableToSetCall = -977
    PhoneNumberBanned = -1000
    TooManyRequests = -5302


class ErrorMessage(Enum):
    Unknown = "原因不明"
    RegisteredUser = "このアカウントはすでに登録されています"
    AccessTokenExpired = "アクセストークンの有効期限切れ"
    ScreenNameAlreadyBeenTaken = "このIDはすでに使われています"
    UserNotFound = "ユーザーが見つかりません"
    PostNotFound = "投稿が見つかりません"
    ChatRoomNotFound = "チャットルームが見つかりません"
    ChatMessageNotFound = "メッセージが見つかりません"
    UserNotFoundAtChatRoom = "チャットルームに特定のユーザーが見つかりません"
    UserMustBeOverTwoAtChatRoom = "チャットルーム内のユーザーは2人以上の必要があります"
    IncorrectPassword = "正しいパスワードを確認してください"
    UserBlocked = "このユーザーはブロック中です"
    PrivateUser = "プライベートユーザーです"
    ApplicationNotFound = "アプリケーションが見つかりません"
    BadSNSCredentials = "アカウント認証に失敗しました"
    SNSAlreadyConnected = "このSNSアカウントはすでに使われています"
    CannotDisconnectSNS = "SNSの連携を解除できません"
    AccessTokenInvalid = "アクセストークンが不正です"
    SpotNotFound = "スポットが見つかりません"
    UserBanned = "このアカウントは永久的に凍結されました"
    UserTemporaryBanned = "サービスの規約に抵触したため、アカウントを停止しています"
    SchoolInfoChange = "学校の情報が変更されています"
    CannotDeleteNewUser = "アカウントを作成して3日以内に削除はできません"
    CaptchaRequired = "Captcha認証が必要です"
    FailedToVerifyCaptcha = "Captcha認証に失敗しました"
    GroupIsFull = "サークルが満員です"
    BannedFromGroup = "このサークルから追放されています"
    InvalidCurrentPassword = "現在のパスワードを入力してください"
    InvalidPassword = "現在のパスワードを確認してください"
    InvalidEmailOrPassword = "メールアドレス、もしくはパスワードが不正です"
    ExistEmail = "このメールアドレスはすでに使われています"
    BadEmailReputation = "スパムメールの可能性があり、操作を完了できません"
    ChatRoomIsFull = "グルチャの最大人数は50人です"
    ConferenceIsFull = "通話が満員です"
    ConferenceInactive = "通話は終了しました"
    GroupOwnerBlockedYou = "サークルの管理人にブロックされています"
    ChatNeedMutualFollowed = "チャットするにはフォローされる必要があります"
    ConferenceCallIsLocked = "通話はロックされています"
    ConferenceCallIsForFollowersOnly = "枠主をフォローすることで参加できます"
    InvalidEmail = "正しいメールアドレスを確認してください"
    RegisteredEmail = "このメールアドレスはすでに登録されています"
    BannedFromCall = "通話に参加できません。永久退出処分を受けています。"
    NotCallOwner = "通話のオーナーではありません"
    NotVipUser = "VIPユーザーではありません"
    BlockingLimitExceeded = "ブロック数の上限に達しました"
    VerificationCodeWrong = "認証コードが不正です"
    VerificationCodeExpired = "認証コードの有効期限切れ"
    InvalidAppVersion = "アプリをアップデートしてください"
    InvalidPhoneNumber = "この機能を使用するためには電話番号認証が必要です"
    FollowLimitation = "制限に達したので、これ以上フォローすることができません。この制限はフォローワーを増やすことで、増やすことができます。"
    AgeGapNotAllowed = "レター、もしくはチャットを送れません。年齢が離れすぎています"
    GroupOwnerOrGroupModeratorOnly = "サークル管理人か副管理人のみ可能な操作です"
    UnableToRegisterUserDueToPolicy = "利用規約に基づき、現在アカウントを登録できません"
    SnsShareRewardAlreadyBeenClaimed = "SNS共有の特典は既に取得済みです"
    QuotaLimitExceeded = "この機能の上限回数に達しました。1時間ほど時間を置いて再度お試しください。"
    ChatNeedAgeVerified = "チャット送信に年齢確認が必要です"
    OnlyAgeVerifiedUserCanJoinGroup = "このサークルに参加するには年齢確認をする必要があります"
    RequirePhoneVerificationToChat = "チャットをするには電話番号認証が必要です"
    NotPostOwner = "編集は投稿の作成者のみ可能です"
    GroupGenerationNotMatched = "特定の年齢層のみ参加が許可されているサークルです"
    PhoneNumberCheckVerificationCodeSubmitQuotaExceeded = (
        "認証コードの送信回数が上限を越えました。1時間ほど時間をおいて再度お試しください"
    )
    PhoneNumberCheckVerificationCodeRequestQuotaExceeded = "チャットをするには電話番号認証が必要です"
    GroupOfferHasBeenAccepted = "サークルの招待は承諾されています"
    GroupOfferHasBeenWithdrawn = "サークルの招待は拒否されています"
    IpBanned = "IPがBanされました"
    NotConnectedToTwitter = "Twitterに接続されていません"
    PrivateUserTimeline = "フォロワーにのみ投稿を公開しています"
    CounterRefreshLimitExceeded = "カウンター更新の回数制限に達しました"
    NotFollowedByOpponent = "このユーザーがあなたをフォローする必要があります"
    ExceedChangeCountryQuota = "国設定は一度設定すると8時間は変更できません。時間をおいて再度お試しください"
    NotGroupMember = "サークルメンバーでないため、この通話に参加できません"
    GroupPendingTransfer = "グループ保留中の移動"
    GroupPendingDeputization = "グループ保留中の代理指名"
    UserRestrictedChatWithCautionUsers = "相手は危険なユーザーとのチャットを許可していません"
    RestrictedCreateChatWithNewUsers = "近日ペナルティなどを受けているため、新規のユーザーとのチャットルーム作成が制限されています"
    RepostPostNotRepostable = "このユーザーの投稿は(´∀｀∩)↑age↑できません"
    TooManyAccountsCreated = "これ以上アカウントを作成することはできません"
    OnlySpecificGenderCanJoinGroup = "特定の性別のみ参加が許可されているサークルです"
    CreateSpecificGenderGroupRequiredGender = "このサークルを作成するには性別の設定が必要です"
    GroupRelatedExceededNumberOfRelatedGroups = "サークルの追加上限に到達しました"
    ExceededPinnedLimit = "これ以上ピン留めできません。解除してからもう一度お試しください"
    GroupShareOnTwitterLimitExceeded = "Twitterでのグループ共有の制限を超えています"
    ReportedContent = "通報を受けているため、この投稿は(´∀｀∩)↑age↑できません"
    ConferenceCallIsForMutualFollowsOnly = "この通話に参加するには枠主と相互にフォローする必要があります"
    ExceededLimit = "投稿の編集上限に達しています。これ以上の編集はできません"
    GroupInviteExceeded = "サークルにおける招待の許容数を超えました。時間をあけてから再度行ってください。"
    PhoneVerificationRequired = "参加するには電話番号認証が必要です。Yay! アプリから電話番号認証を行なってください。"
    ContentTooOld = "一定期間の経過後、投稿の編集はできません"
    PasswordTooShort = "パスワードが短すぎます"
    PasswordTooLong = "パスワードが長すぎます"
    PasswordNotAllowed = "他のパスワードを使用してください。無効な文字列が含まれています"
    CommonPassword = "他のパスワードを使用してください。文字列や数字、または記号などの組み合わせをお試しください"
    EmailNotAuthorized = "このメールアドレスを使用するには承認される必要があります"
    UnableToMovePostToThread = "投稿をスレッドに移動できません"
    UnableToPostUrl = "URLを投稿できません"
    UnableToSetCall = "通話を開始できませんでした"
    PhoneNumberBanned = "電話番号がBanされています"
    TooManyRequests = "リクエストが多すぎます"
