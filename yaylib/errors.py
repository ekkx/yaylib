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

from json.decoder import JSONDecodeError

import aiohttp

from .responses import ErrorResponse


class YaylibError(Exception):
    """yaylib における例外基底クラス"""


class ClientError(YaylibError):
    """yaylib クライアントに関する例外クラス"""

    def __init__(self, response: ErrorResponse):
        super().__init__(
            f'Server returned error code: {response.error_code}, with message "{response.message}"'
        )

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


class HTTPError(YaylibError):
    """HTTP に関する例外クラス"""

    def __init__(self, response: aiohttp.ClientResponse):
        super().__init__(f"HTTP request failed with status: {response.status}")

        self.__response = response

    @property
    def response(self) -> aiohttp.ClientResponse:
        """HTTP レスポンス"""
        return self.__response


class HTTPBadRequestError(HTTPError):
    """Exception raised for a 400 HTTP status code"""


class HTTPAuthenticationError(HTTPError):
    """Exception raised for a 401 HTTP status code"""


class HTTPForbiddenError(HTTPError):
    """Exception raised for a 403 HTTP status code"""


class HTTPNotFoundError(HTTPError):
    """Exception raised for a 404 HTTP status code"""


class HTTPRateLimitError(HTTPError):
    """Exception raised for a 429 HTTP status code"""


class HTTPInternalServerError(HTTPError):
    """Exception raised for a 5xx HTTP status code"""


# pylint: disable=too-many-statements
async def raise_for_code(response: aiohttp.ClientResponse) -> None:
    """`error_code` によって例外を発生させる

    Args:
        response (aiohttp.ClientResponse):

    Raises:
        ClientError:
    """
    try:
        response_json = await response.json(content_type=None)
        if response_json is None:
            return None
    except JSONDecodeError:
        return None

    err = ErrorResponse(response_json)
    if err.result is None or err.result != "error":
        return

    match err.error_code:
        case 0:
            raise UnknownError(err)
        case -1:
            raise InvalidParameterError(err)
        case -2:
            raise RegisteredUserError(err)
        case -3:
            raise AccessTokenExpiredError(err)
        case -4:
            raise ScreenNameAlreadyBeenTakenError(err)
        case -5:
            raise UserNotFoundError(err)
        case -6:
            raise PostNotFoundError(err)
        case -7:
            raise ChatRoomNotFoundError(err)
        case -8:
            raise ChatMessageNotFoundError(err)
        case -9:
            raise UserNotFoundAtChatRoomError(err)
        case -10:
            raise UserMustBeOverTwoAtChatRoomError(err)
        case -11:
            raise IncorrectPasswordError(err)
        case -12:
            raise UserBlockedError(err)
        case -13:
            raise PrivateUserError(err)
        case -14:
            raise ApplicationNotFoundError(err)
        case -15:
            raise BadSNSCredentialsError(err)
        case -16:
            raise SNSAlreadyConnectedError(err)
        case -17:
            raise CannotDisconnectSNSError(err)
        case -18:
            raise AccessTokenInvalidError(err)
        case -19:
            raise SpotNotFoundError(err)
        case -20:
            raise UserBannedError(err)
        case -21:
            raise UserTemporaryBannedError(err)
        case -22:
            raise SchoolInfoChangeError(err)
        case -26:
            raise CannotDeleteNewUserError(err)
        case -29:
            raise CaptchaRequiredError(err)
        case -30:
            raise FailedToVerifyCaptchaError(err)
        case -31:
            raise Required2FAError(err)
        case -32:
            raise Incorrect2FAError(err)
        case -100:
            raise GroupIsFullError(err)
        case -103:
            raise BannedFromGroupError(err)
        case -200:
            raise InvalidCurrentPasswordError(err)
        case -201:
            raise InvalidPasswordError(err)
        case -202:
            raise InvalidEmailOrPasswordError(err)
        case -203:
            raise ExistEmailError(err)
        case -204:
            raise BadEmailReputationError(err)
        case -308:
            raise ChatRoomIsFullError(err)
        case -309:
            raise ConferenceIsFullError(err)
        case -310:
            raise ConferenceInactiveError(err)
        case -312:
            raise GroupOwnerBlockedYouError(err)
        case -313:
            raise ChatNeedMutualFollowedError(err)
        case -315:
            raise ConferenceCallIsLockedError(err)
        case -317:
            raise ConferenceCallIsForFollowersOnlyError(err)
        case -319:
            raise InvalidEmailError(err)
        case -320:
            raise RegisteredEmailError(err)
        case -321:
            raise BannedFromCallError(err)
        case -322:
            raise NotCallOwnerError(err)
        case -326:
            raise NotVipUserError(err)
        case -331:
            raise BlockingLimitExceededError(err)
        case -332:
            raise VerificationCodeWrongError(err)
        case -333:
            raise VerificationCodeExpiredError(err)
        case -335:
            raise InvalidPhoneNumberError(err)
        case -336:
            raise FollowLimitationError(err)
        case -338:
            raise AgeGapNotAllowedError(err)
        case -339:
            raise GroupOwnerOrGroupModeratorOnlyError(err)
        case -340:
            raise UnableToRegisterUserDueToPolicyError(err)
        case -342:
            raise SNSShareRewardAlreadyBeenClaimedError(err)
        case -343:
            raise QuotaLimitExceededError(err)
        case -346:
            raise ChatNeedAgeVerifiedError(err)
        case -347:
            raise OnlyAgeVerifiedUserCanJoinGroupError(err)
        case -348:
            raise RequirePhoneVerificationToChatError(err)
        case -350:
            raise NotPostOwnerError(err)
        case -352:
            raise GroupGenerationNotMatchedError(err)
        case -355:
            raise PhoneNumberCheckVerificationCodeSubmitQuotaExceededError(err)
        case -356:
            raise PhoneNumberCheckVerificationCodeRequestQuotaExceededError(err)
        case -357:
            raise GroupOfferHasBeenAcceptedError(err)
        case -358:
            raise GroupOfferHasBeenWithdrawnError(err)
        case -360:
            raise IpBannedError(err)
        case -361:
            raise NotConnectedToTwitterError(err)
        case -363:
            raise PrivateUserTimelineError(err)
        case -364:
            raise CounterRefreshLimitExceededError(err)
        case -367:
            raise NotFollowedByOpponentError(err)
        case -369:
            raise ExceedChangeCountryQuotaError(err)
        case -370:
            raise NotGroupMemberError(err)
        case -371:
            raise GroupPendingTransferError(err)
        case -372:
            raise GroupPendingDeputizationError(err)
        case -373:
            raise UserRestrictedChatWithCautionUsersError(err)
        case -374:
            raise RestrictedCreateChatWithNewUsersError(err)
        case -375:
            raise RepostPostNotRepostableError(err)
        case -376:
            raise TooManyAccountsCreatedError(err)
        case -377:
            raise OnlySpecificGenderCanJoinGroupError(err)
        case -378:
            raise CreateSpecificGenderGroupRequiredGenderError(err)
        case -382:
            raise GroupRelatedExceededNumberOfRelatedGroupsError(err)
        case -383:
            raise ExceededPinnedLimitError(err)
        case -384:
            raise GroupShareOnTwitterLimitExceededError(err)
        case -385:
            raise ReportedContentError(err)
        case -400:
            raise InsufficientCoinsError(err)
        case -402:
            raise ConferenceCallIsForMutualFollowsOnlyError(err)
        case -403:
            raise ExceededLimitError(err)
        case -404:
            raise GroupInviteExceededError(err)
        case -405:
            raise PhoneVerificationRequiredError(err)
        case -406:
            raise ContentTooOldError(err)
        case -407:
            raise PasswordTooShortError(err)
        case -408:
            raise PasswordTooLongError(err)
        case -409:
            raise PasswordNotAllowedError(err)
        case -410:
            raise CommonPasswordError(err)
        case -411:
            raise EmailNotAuthorizedError(err)
        case -412:
            raise UnableToMovePostToThreadError(err)
        case -413:
            raise UnableToPostUrlError(err)
        case -415:
            raise ReferralAlreadyRegisteredError(err)
        case -416:
            raise MuteUserOverLimitError(err)
        case -800:
            raise InvalidAppVersionError(err)
        case -977:
            raise UnableToSetCallError(err)
        case -999:
            raise DynamicErrorMessageError(err)
        case -1000:
            raise PhoneNumberBannedError(err)
        case 400:
            raise BadRequestError(err)
        case 401:
            raise UnauthorizedError(err)
        case 403:
            raise AccessForbiddenError(err)
        case 404:
            raise NotFoundError(err)
        case 409:
            raise ConflictError(err)
        case 429:
            raise TooManyRequestsError(err)
        case 500:
            raise InternalServerError(err)
        case 4002:
            raise Web3AccountAlreadyLinkedToAnotherWalletError(err)
        case 4003:
            raise Web3WalletAlreadyLinkedToAnotherAccountError(err)
        case 4005:
            raise Web3PalLevelUpBattlesRequiredError(err)
        case 4006:
            raise Web3PalLevelUpMaximumLevelReachedError(err)
        case 4010:
            raise Web3PalPoolCooldownError(err)
        case 4011:
            raise Web3PalPoolEmptyError(err)
        case 4012:
            raise Web3PalAlreadyBattleError(err)
        case 4017:
            raise Web3WalletHasPendingTransactionsError(err)
        case 5003:
            raise Web3WalletNetworkErrorError(err)
        case 6001:
            raise Web3EMPLInsufficientFundsError(err)
        case 6002:
            raise Web3EMPLFeeExceedsBalanceError(err)
        case _:
            raise ClientError(err)


async def raise_for_status(response: aiohttp.ClientResponse):
    """HTTP ステータスコードを元に例外を発生させる

    Args:
        response (aiohttp.ClientResponse):

    Raises:
        HTTPError:
    """
    match response.status:
        case 400:
            raise HTTPBadRequestError(response)
        case 401:
            raise HTTPAuthenticationError(response)
        case 403:
            raise HTTPForbiddenError(response)
        case 404:
            raise HTTPNotFoundError(response)
        case status if status >= 500:
            raise HTTPInternalServerError(response)
        case status if status and not 200 <= status < 300:
            raise HTTPError(response)
