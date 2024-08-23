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

import aiohttp

from .responses import ErrorResponse


class YaylibError(Exception):  # pylint: disable=used-before-assignment
    """yaylib における例外基底クラス"""


class ClientError(YaylibError):
    """yaylib クライアントに関する例外クラス"""

    def __init__(self, response: ErrorResponse):
        super().__init__(f"Server returned error code: {response.error_code}")

        self.__response = response

    @property
    def response(self) -> ErrorResponse:
        """エラーレスポンス"""
        return self.__response


class UnknownError(ClientError):
    """原因不明"""


class InvalidParameterError(ClientError):
    """パラメターが不正"""


class RegisteredUserError(ClientError):
    """このアカウントはすでに登録されています"""


class AccessTokenExpiredError(ClientError):
    """アクセストークンの有効期限切れ"""


class ScreenNameAlreadyBeenTakenError(ClientError):
    """このIDはすでに使われています"""


class UserNotFoundError(ClientError):
    """ユーザーが見つかりません"""


class PostNotFoundError(ClientError):
    """投稿が見つかりません"""


class ChatRoomNotFoundError(ClientError):
    """チャットルームが見つかりません"""


class ChatMessageNotFoundError(ClientError):
    """メッセージが見つかりません"""


class UserNotFoundAtChatRoomError(ClientError):
    """チャットルームに特定のユーザーが見つかりません"""


class UserMustBeOverTwoAtChatRoomError(ClientError):
    """チャットルーム内のユーザーは2人以上の必要があります"""


class IncorrectPasswordError(ClientError):
    """正しいパスワードを確認してください"""


class UserBlockedError(ClientError):
    """このユーザーはブロック中です"""


class PrivateUserError(ClientError):
    """プライベートユーザーです"""


class ApplicationNotFoundError(ClientError):
    """アプリケーションが見つかりません"""


class BadSNSCredentialsError(ClientError):
    """アカウント認証に失敗しました"""


class SNSAlreadyConnectedError(ClientError):
    """このSNSアカウントはすでに使われています"""


class CannotDisconnectSNSError(ClientError):
    """SNSの連携を解除できません"""


class AccessTokenInvalidError(ClientError):
    """アクセストークンが不正です"""


class SpotNotFoundError(ClientError):
    """スポットが見つかりません"""


class UserBannedError(ClientError):
    """このアカウントは永久的に凍結されました"""


class UserTemporaryBannedError(ClientError):
    """サービスの規約に抵触したため、アカウントを停止しています"""


class SchoolInfoChangeError(ClientError):
    """学校の情報が変更されています"""


class CannotDeleteNewUserError(ClientError):
    """アカウントを作成して3日以内に削除はできません"""


class CaptchaRequiredError(ClientError):
    """Captcha認証が必要です"""


class FailedToVerifyCaptchaError(ClientError):
    """Captcha認証に失敗しました"""


class Required2FAError(ClientError):
    """2段階認証が必要です"""


class Incorrect2FAError(ClientError):
    """2段階認証が不正です"""


class GroupIsFullError(ClientError):
    """サークルが満員です"""


class BannedFromGroupError(ClientError):
    """このサークルから追放されています"""


class InvalidCurrentPasswordError(ClientError):
    """現在のパスワードを入力してください"""


class InvalidPasswordError(ClientError):
    """現在のパスワードを確認してください"""


class InvalidEmailOrPasswordError(ClientError):
    """メールアドレス、もしくはパスワードが不正です"""


class ExistEmailError(ClientError):
    """このメールアドレスはすでに使われています"""


class BadEmailReputationError(ClientError):
    """スパムメールの可能性があり、操作を完了できません"""


class ChatRoomIsFullError(ClientError):
    """グループチャットの最大人数は50人です"""


class ConferenceIsFullError(ClientError):
    """通話が満員です"""


class ConferenceInactiveError(ClientError):
    """通話は終了しました"""


class GroupOwnerBlockedYouError(ClientError):
    """サークルの管理人にブロックされています"""


class ChatNeedMutualFollowedError(ClientError):
    """チャットするにはフォローされる必要があります"""


class ConferenceCallIsLockedError(ClientError):
    """通話はロックされています"""


class ConferenceCallIsForFollowersOnlyError(ClientError):
    """枠主をフォローすることで通話に参加できます"""


class InvalidEmailError(ClientError):
    """正しいメールアドレスを確認してください"""


class RegisteredEmailError(ClientError):
    """このメールアドレスはすでに登録されています"""


class BannedFromCallError(ClientError):
    """通話に参加できません。永久退出処分を受けています。"""


class NotCallOwnerError(ClientError):
    """通話のオーナーではありません"""


class NotVipUserError(ClientError):
    """VIPユーザーではありません"""


class BlockingLimitExceededError(ClientError):
    """ブロック数の上限に達しました"""


class VerificationCodeWrongError(ClientError):
    """認証コードが不正です"""


class VerificationCodeExpiredError(ClientError):
    """認証コードの有効期限切れ"""


class InvalidPhoneNumberError(ClientError):
    """この機能を使用するためには電話番号認証が必要です"""


class FollowLimitationError(ClientError):
    """制限に達したので、これ以上フォローすることができません。この制限はフォローワーを増やすことで、増やすことができます。"""


class AgeGapNotAllowedError(ClientError):
    """レター、もしくはチャットを送れません。年齢が離れすぎています"""


class GroupOwnerOrGroupModeratorOnlyError(ClientError):
    """サークル管理人か副管理人のみ可能な操作です"""


class UnableToRegisterUserDueToPolicyError(ClientError):
    """利用規約に基づき、現在アカウントを登録できません"""


class SNSShareRewardAlreadyBeenClaimedError(ClientError):
    """SNS共有の特典は既に取得済みです"""


class QuotaLimitExceededError(ClientError):
    """この機能の上限回数に達しました。1時間ほど時間を置いて再度お試しください。"""


class ChatNeedAgeVerifiedError(ClientError):
    """チャット送信に年齢確認が必要です"""


class OnlyAgeVerifiedUserCanJoinGroupError(ClientError):
    """このサークルに参加するには年齢確認をする必要があります"""


class RequirePhoneVerificationToChatError(ClientError):
    """チャットをするには電話番号認証が必要です"""


class NotPostOwnerError(ClientError):
    """編集は投稿の作成者のみ可能です"""


class GroupGenerationNotMatchedError(ClientError):
    """特定の年齢層のみ参加が許可されているサークルです"""


class PhoneNumberCheckVerificationCodeSubmitQuotaExceededError(ClientError):
    """SMS認証コードの送信回数が上限を越えました。1時間ほど時間をおいて再度お試しください"""


class PhoneNumberCheckVerificationCodeRequestQuotaExceededError(ClientError):
    """チャットをするには電話番号認証が必要です"""


class GroupOfferHasBeenAcceptedError(ClientError):
    """サークルの招待はすでに承諾されています"""


class GroupOfferHasBeenWithdrawnError(ClientError):
    """サークルの招待は拒否されています"""


class IpBannedError(ClientError):
    """IPアドレスが凍結されました"""


class NotConnectedToTwitterError(ClientError):
    """Twitterに接続されていません"""


class PrivateUserTimelineError(ClientError):
    """フォロワーにのみ投稿を公開しています"""


class CounterRefreshLimitExceededError(ClientError):
    """カウンター更新の回数制限に達しました"""


class NotFollowedByOpponentError(ClientError):
    """ユーザーがあなたをフォローする必要があります"""


class ExceedChangeCountryQuotaError(ClientError):
    """国設定は一度設定すると8時間は変更できません。時間をおいて再度お試しください"""


class NotGroupMemberError(ClientError):
    """サークルメンバーでないため、この通話に参加できません"""


class GroupPendingTransferError(ClientError):
    """グループ保留中の移動"""


class GroupPendingDeputizationError(ClientError):
    """グループ保留中の代理指名"""


class UserRestrictedChatWithCautionUsersError(ClientError):
    """相手は危険なユーザーとのチャットを許可していません"""


class RestrictedCreateChatWithNewUsersError(ClientError):
    """近日ペナルティなどを受けているため、新規のユーザーとのチャットルーム作成が制限されています"""


class RepostPostNotRepostableError(ClientError):
    """このユーザーの投稿は(´∀｀∩)↑age↑できません"""


class TooManyAccountsCreatedError(ClientError):
    """これ以上アカウントを作成することはできません"""


class OnlySpecificGenderCanJoinGroupError(ClientError):
    """特定の性別のみ参加が許可されているサークルです"""


class CreateSpecificGenderGroupRequiredGenderError(ClientError):
    """このサークルを作成するには性別の設定が必要です"""


class GroupRelatedExceededNumberOfRelatedGroupsError(ClientError):
    """サークルの追加上限に到達しました"""


class ExceededPinnedLimitError(ClientError):
    """これ以上ピン留めできません。解除してからもう一度お試しください"""


class GroupShareOnTwitterLimitExceededError(ClientError):
    """Twitterでのグループ共有の制限を超えています"""


class ReportedContentError(ClientError):
    """通報を受けているため、この投稿は(´∀｀∩)↑age↑できません"""


class InsufficientCoinsError(ClientError):
    """コインが不十分です"""


class ConferenceCallIsForMutualFollowsOnlyError(ClientError):
    """この通話に参加するには枠主と相互にフォローする必要があります"""


class ExceededLimitError(ClientError):
    """操作の上限に達しています"""


class GroupInviteExceededError(ClientError):
    """サークルにおける招待の許容数を超えました。時間をあけてから再度行ってください。"""


class PhoneVerificationRequiredError(ClientError):
    """参加するには電話番号認証が必要です。Yay! アプリから電話番号認証を行なってください。"""


class ContentTooOldError(ClientError):
    """一定期間の経過後、投稿の編集はできません"""


class PasswordTooShortError(ClientError):
    """パスワードが短すぎます"""


class PasswordTooLongError(ClientError):
    """パスワードが長すぎます"""


class PasswordNotAllowedError(ClientError):
    """他のパスワードを使用してください。無効な文字列が含まれています"""


class CommonPasswordError(ClientError):
    """他のパスワードを使用してください。文字列や数字、または記号などの組み合わせをお試しください"""


class EmailNotAuthorizedError(ClientError):
    """このメールアドレスを使用するには承認される必要があります"""


class UnableToMovePostToThreadError(ClientError):
    """投稿をスレッドに移動できません"""


class UnableToPostUrlError(ClientError):
    """URLを投稿できません"""


class ReferralAlreadyRegisteredError(ClientError):
    """description"""


class MuteUserOverLimitError(ClientError):
    """description"""


class InvalidAppVersionError(ClientError):
    """description"""


class UnableToSetCallError(ClientError):
    """通話を開始できませんでした"""


class DynamicErrorMessageError(ClientError):
    """description"""


class PhoneNumberBannedError(ClientError):
    """電話番号が凍結されています"""


class BadRequestError(ClientError):
    """リクエストが多すぎます"""


class UnauthorizedError(ClientError):
    """認証エラー"""


class AccessForbiddenError(ClientError):
    """アクセスが拒否されました"""


class NotFoundError(ClientError):
    """コンテンツは見つかりませんでした"""


class ConflictError(ClientError):
    """衝突が発生しました"""


class TooManyRequestsError(ClientError):
    """リクエストが多すぎます"""


class InternalServerError(ClientError):
    """サーバーエラー"""


class Web3AccountAlreadyLinkedToAnotherWalletError(ClientError):
    """この Web3 アカウントはすでに他のウォレットに連携されています"""


class Web3WalletAlreadyLinkedToAnotherAccountError(ClientError):
    """この Web3 アカウントはすでに他のアカウントに連携されています"""


class Web3PalLevelUpBattlesRequiredError(ClientError):
    """Web3PalLevelUpBattlesRequiredError"""


class Web3PalLevelUpMaximumLevelReachedError(ClientError):
    """Web3PalLevelUpMaximumLevelReachedError"""


class Web3PalPoolCooldownError(ClientError):
    """Web3PalPoolCooldownError"""


class Web3PalPoolEmptyError(ClientError):
    """Web3PalPoolEmptyError"""


class Web3PalAlreadyBattleError(ClientError):
    """Web3PalAlreadyBattleError"""


class Web3WalletHasPendingTransactionsError(ClientError):
    """Web3WalletHasPendingTransactionsError"""


class Web3WalletNetworkErrorError(ClientError):
    """Web3WalletNetworkErrorError"""


class Web3EMPLInsufficientFundsError(ClientError):
    """Web3EMPLInsufficientFundsError"""


class Web3EMPLFeeExceedsBalanceError(ClientError):
    """Web3EMPLFeeExceedsBalanceError"""


class ErrorMessage:
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
    QuotaLimitExceeded = (
        "この機能の上限回数に達しました。1時間ほど時間を置いて再度お試しください。"
    )
    ChatNeedAgeVerified = "チャット送信に年齢確認が必要です"
    OnlyAgeVerifiedUserCanJoinGroup = (
        "このサークルに参加するには年齢確認をする必要があります"
    )
    RequirePhoneVerificationToChat = "チャットをするには電話番号認証が必要です"
    NotPostOwner = "編集は投稿の作成者のみ可能です"
    GroupGenerationNotMatched = "特定の年齢層のみ参加が許可されているサークルです"
    PhoneNumberCheckVerificationCodeSubmitQuotaExceeded = "認証コードの送信回数が上限を越えました。1時間ほど時間をおいて再度お試しください"
    PhoneNumberCheckVerificationCodeRequestQuotaExceeded = (
        "チャットをするには電話番号認証が必要です"
    )
    GroupOfferHasBeenAccepted = "サークルの招待は承諾されています"
    GroupOfferHasBeenWithdrawn = "サークルの招待は拒否されています"
    IpBanned = "IPがBanされました"
    NotConnectedToTwitter = "Twitterに接続されていません"
    PrivateUserTimeline = "フォロワーにのみ投稿を公開しています"
    CounterRefreshLimitExceeded = "カウンター更新の回数制限に達しました"
    NotFollowedByOpponent = "このユーザーがあなたをフォローする必要があります"
    ExceedChangeCountryQuota = (
        "国設定は一度設定すると8時間は変更できません。時間をおいて再度お試しください"
    )
    NotGroupMember = "サークルメンバーでないため、この通話に参加できません"
    GroupPendingTransfer = "グループ保留中の移動"
    GroupPendingDeputization = "グループ保留中の代理指名"
    UserRestrictedChatWithCautionUsers = (
        "相手は危険なユーザーとのチャットを許可していません"
    )
    RestrictedCreateChatWithNewUsers = "近日ペナルティなどを受けているため、新規のユーザーとのチャットルーム作成が制限されています"
    RepostPostNotRepostable = "このユーザーの投稿は(´∀｀∩)↑age↑できません"
    TooManyAccountsCreated = "これ以上アカウントを作成することはできません"
    OnlySpecificGenderCanJoinGroup = "特定の性別のみ参加が許可されているサークルです"
    CreateSpecificGenderGroupRequiredGender = (
        "このサークルを作成するには性別の設定が必要です"
    )
    GroupRelatedExceededNumberOfRelatedGroups = "サークルの追加上限に到達しました"
    ExceededPinnedLimit = (
        "これ以上ピン留めできません。解除してからもう一度お試しください"
    )
    GroupShareOnTwitterLimitExceeded = "Twitterでのグループ共有の制限を超えています"
    ReportedContent = "通報を受けているため、この投稿は(´∀｀∩)↑age↑できません"
    ConferenceCallIsForMutualFollowsOnly = (
        "この通話に参加するには枠主と相互にフォローする必要があります"
    )
    ExceededLimit = "投稿の編集上限に達しています。これ以上の編集はできません"
    GroupInviteExceeded = (
        "サークルにおける招待の許容数を超えました。時間をあけてから再度行ってください。"
    )
    PhoneVerificationRequired = "参加するには電話番号認証が必要です。Yay! アプリから電話番号認証を行なってください。"
    ContentTooOld = "一定期間の経過後、投稿の編集はできません"
    PasswordTooShort = "パスワードが短すぎます"
    PasswordTooLong = "パスワードが長すぎます"
    PasswordNotAllowed = (
        "他のパスワードを使用してください。無効な文字列が含まれています"
    )
    CommonPassword = "他のパスワードを使用してください。文字列や数字、または記号などの組み合わせをお試しください"
    EmailNotAuthorized = "このメールアドレスを使用するには承認される必要があります"
    UnableToMovePostToThread = "投稿をスレッドに移動できません"
    UnableToPostUrl = "URLを投稿できません"
    UnableToSetCall = "通話を開始できませんでした"
    PhoneNumberBanned = "電話番号がBanされています"
    TooManyRequests = "リクエストが多すぎます"
