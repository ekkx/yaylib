from enum import Enum


class Configs:

    YAYLIB_VERSION = "0.1.2"
    YAY_API_VERSION = "3.16"
    YAY_VERSION_NAME = "3.16.1"
    YAY_API_VERSION_KEY = "e83a1d2588918c2061280427c88e6f56"
    YAY_API_KEY = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
    YAY_SHARED_KEY = "yayZ1"
    YAY_STORE_KEY = "yayZ1payment"
    ID_CARD_CHECK_SECRET_KEY = "4aa6d1c301a97154bc1098c2"
    YAY_REVIEW_HOST_1 = "review.yay.space"
    YAY_REVIEW_HOST_2 = "cas-stg.yay.space"
    YAY_STAGING_HOST_1 = "stg.yay.space"
    YAY_STAGING_HOST_2 = "cas.yay.space"
    YAY_PRODUCTION_HOST = "api.yay.space"
    YAY_API_URL = "https://" + YAY_PRODUCTION_HOST
    ID_CARD_CHECK_HOST_PRODUCTION = "idcardcheck.com"
    ID_CARD_CHECK_HOST_STAGING = "stg.idcardcheck.com"
    USER_AGENT = "android 11 (3.5x 1440x2960 Galaxy S9)"
    REQUEST_HEADERS = {
        "Host": YAY_PRODUCTION_HOST,
        "X-App-Version": YAY_API_VERSION,
        "User-Agent": USER_AGENT,
        "X-Device-Info": f"yay {YAY_VERSION_NAME} {USER_AGENT}",
        "X-Device-Uuid": "",
        "X-Connection-Type": "wifi",
        "Accept-Language": "ja",
        "Content-Type": "application/json;charset=UTF-8"
    }


class Endpoints:

    # api v1 endpoints
    USERS_V1 = Configs.YAY_API_URL + "/v1/users"
    PAYMENTS_V1 = Configs.YAY_API_URL + "/v1/payments"
    THREADS_V1 = Configs.YAY_API_URL + "/v1/threads"
    PINNED_V1 = Configs.YAY_API_URL + "/v1/pinned"
    POSTS_V1 = Configs.YAY_API_URL + "/v1/posts"
    CONVERSATIONS_V1 = Configs.YAY_API_URL + "/v1/conversations"
    HIDDEN_V1 = Configs.YAY_API_URL + "/v1/hidden"
    GROUPS_V1 = Configs.YAY_API_URL + "/v1/groups"
    CHAT_ROOMS_V1 = Configs.YAY_API_URL + "/v1/chat_rooms"
    CALLS_V1 = Configs.YAY_API_URL + "/v1/calls"
    SURVEYS_V1 = Configs.YAY_API_URL + "/v1/surveys"
    FRIENDS_V1 = Configs.YAY_API_URL + "/v1/friends"
    GAMES_V1 = Configs.YAY_API_URL + "/v1/games"
    GENRES_V1 = Configs.YAY_API_URL + "/v1/genres"
    ANONYMOUS_CALLS_V1 = Configs.YAY_API_URL + "/v1/anonymous_calls"
    NOTIFICATION_SETTINGS_V1 = Configs.YAY_API_URL + "/v1/notification_settings"
    STICKER_PACKS_V1 = Configs.YAY_API_URL + "/v1/sticker_packs"
    SNS_THUMBNAIL_V1 = Configs.YAY_API_URL + "/v1/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V1 = Configs.YAY_API_URL + "/v1/email_verification_urls"
    BUCKETS_V1 = Configs.YAY_API_URL + "/v1/buckets"
    ID_CHECK_V1 = Configs.YAY_API_URL + "/v1/id_check"
    PROMOTIONS_V1 = Configs.YAY_API_URL + "/v1/promotions"
    SKYFALL_V1 = Configs.YAY_API_URL + "/v1/skyfall"
    GENUINE_DEVICES_V1 = Configs.YAY_API_URL + "/v1/genuine_devices"
    LOBI_FRIENDS_V1 = Configs.YAY_API_URL + "/v1/lobi_friends"

    # api v2 endpoints
    USERS_V2 = Configs.YAY_API_URL + "/v2/users"
    PAYMENTS_V2 = Configs.YAY_API_URL + "/v2/payments"
    THREADS_V2 = Configs.YAY_API_URL + "/v2/threads"
    PINNED_V2 = Configs.YAY_API_URL + "/v2/pinned"
    POSTS_V2 = Configs.YAY_API_URL + "/v2/posts"
    CONVERSATIONS_V2 = Configs.YAY_API_URL + "/v2/conversations"
    HIDDEN_V2 = Configs.YAY_API_URL + "/v2/hidden"
    GROUPS_V2 = Configs.YAY_API_URL + "/v2/groups"
    CHAT_ROOMS_V2 = Configs.YAY_API_URL + "/v2/chat_rooms"
    CALLS_V2 = Configs.YAY_API_URL + "/v2/calls"
    SURVEYS_V2 = Configs.YAY_API_URL + "/v2/surveys"
    FRIENDS_V2 = Configs.YAY_API_URL + "/v2/friends"
    GAMES_V2 = Configs.YAY_API_URL + "/v2/games"
    GENRES_V2 = Configs.YAY_API_URL + "/v2/genres"
    ANONYMOUS_CALLS_V2 = Configs.YAY_API_URL + "/v2/anonymous_calls"
    NOTIFICATION_SETTINGS_V2 = Configs.YAY_API_URL + "/v2/notification_settings"
    STICKER_PACKS_V2 = Configs.YAY_API_URL + "/v2/sticker_packs"
    SNS_THUMBNAIL_V2 = Configs.YAY_API_URL + "/v2/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V2 = Configs.YAY_API_URL + "/v2/email_verification_urls"
    BUCKETS_V2 = Configs.YAY_API_URL + "/v2/buckets"
    ID_CHECK_V2 = Configs.YAY_API_URL + "/v2/id_check"
    PROMOTIONS_V2 = Configs.YAY_API_URL + "/v2/promotions"
    SKYFALL_V2 = Configs.YAY_API_URL + "/v2/skyfall"
    GENUINE_DEVICES_V2 = Configs.YAY_API_URL + "/v2/genuine_devices"
    LOBI_FRIENDS_V2 = Configs.YAY_API_URL + "/v2/lobi_friends"

    # api v3 endpoints
    USERS_V3 = Configs.YAY_API_URL + "/v3/users"
    PAYMENTS_V3 = Configs.YAY_API_URL + "/v3/payments"
    THREADS_V3 = Configs.YAY_API_URL + "/v3/threads"
    PINNED_V3 = Configs.YAY_API_URL + "/v3/pinned"
    POSTS_V3 = Configs.YAY_API_URL + "/v3/posts"
    CONVERSATIONS_V3 = Configs.YAY_API_URL + "/v3/conversations"
    HIDDEN_V3 = Configs.YAY_API_URL + "/v3/hidden"
    GROUPS_V3 = Configs.YAY_API_URL + "/v3/groups"
    CHAT_ROOMS_V3 = Configs.YAY_API_URL + "/v3/chat_rooms"
    CALLS_V3 = Configs.YAY_API_URL + "/v3/calls"
    SURVEYS_V3 = Configs.YAY_API_URL + "/v3/surveys"
    FRIENDS_V3 = Configs.YAY_API_URL + "/v3/friends"
    GAMES_V3 = Configs.YAY_API_URL + "/v3/games"
    GENRES_V3 = Configs.YAY_API_URL + "/v3/genres"
    ANONYMOUS_CALLS_V3 = Configs.YAY_API_URL + "/v3/anonymous_calls"
    NOTIFICATION_SETTINGS_V3 = Configs.YAY_API_URL + "/v3/notification_settings"
    STICKER_PACKS_V3 = Configs.YAY_API_URL + "/v3/sticker_packs"
    SNS_THUMBNAIL_V3 = Configs.YAY_API_URL + "/v3/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V3 = Configs.YAY_API_URL + "/v3/email_verification_urls"
    BUCKETS_V3 = Configs.YAY_API_URL + "/v3/buckets"
    ID_CHECK_V3 = Configs.YAY_API_URL + "/v3/id_check"
    PROMOTIONS_V3 = Configs.YAY_API_URL + "/v3/promotions"
    SKYFALL_V3 = Configs.YAY_API_URL + "/v3/skyfall"
    GENUINE_DEVICES_V3 = Configs.YAY_API_URL + "/v3/genuine_devices"
    LOBI_FRIENDS_V3 = Configs.YAY_API_URL + "/v3/lobi_friends"


class ErrorType(Enum):

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
    UnableToMovePostToThread = -412
    UnableToPostUrl = -413
    UnableToSetCall = -977
    PhoneNumberBanned = -1000


class ErrorMessage(Enum):

    Unknown = "原因不明"
    # InvalidParameter = "引数が不正です"
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
    AgeGapNotAllowed = "レターを送れません。年齢が離れすぎています"
    GroupOwnerOrGroupModeratorOnly = "サークル管理人か副管理人のみ可能な操作です"
    UnableToRegisterUserDueToPolicy = "利用規約に基づき、現在アカウントを登録できません"
    SnsShareRewardAlreadyBeenClaimed = "SNS共有の特典は既に取得済みです"
    QuotaLimitExceeded = "この機能の上限回数に達しました。1時間ほど時間を置いて再度お試しください。"
    ChatNeedAgeVerified = "チャット送信に年齢確認が必要です"
    OnlyAgeVerifiedUserCanJoinGroup = "このサークルに参加するには年齢確認をする必要があります"
    RequirePhoneVerificationToChat = "チャットをするには電話番号認証が必要です"
    NotPostOwner = "編集は投稿の作成者のみ可能です"
    GroupGenerationNotMatched = "特定の年齢層のみ参加が許可されているサークルです"
    PhoneNumberCheckVerificationCodeSubmitQuotaExceeded = "認証コードの送信回数が上限を越えました。1時間ほど時間をおいて再度お試しください"
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
    UnableToMovePostToThread = "投稿をスレッドに移動できません"
    UnableToPostUrl = "URLを投稿できません"
    UnableToSetCall = "通話を開始できませんでした"
    PhoneNumberBanned = "電話番号がBanされています"


class LogMessage(Enum):

    # -CALL
    bump_call = {
        "en": "Call bumped.",
        "ja": "Call bumped."
    }
    invite_to_call_bulk = {
        "en": "Invited your online followings to the call",
        "ja": "オンラインを友達をまとめて通話に招待しました"
    }
    invite_users_to_call = {
        "en": "Invitation sent.",
        "ja": "ユーザーを通話に招待しました"
    }
    invite_users_to_chat_call = {
        "en": "Invitation of chat call sent.",
        "ja": "ユーザーをチャット通話に招待しました"
    }
    kick_and_ban_from_call = {
        "en": "User has been banned from the call.",
        "ja": "ユーザーを通話から追放しました"
    }
    notify_anonymous_user_leave_agora_channel = {
        "en": "Notified anonymous user left the call.",
        "ja": "無名くんが退出したことを通知しました"
    }
    notify_user_leave_agora_channel = {
        "en": "Notified user '{user_id}' left the call.",
        "ja": "ユーザーが通話から退出したことを通知しました"
    }
    send_call_screenshot = {
        "en": "Call screenshot sent.",
        "ja": "通話のスクリーンショットを送信しました"
    }
    set_call = {
        "en": "Call set.",
        "ja": "通話を開始しました"
    }
    set_user_role = {
        "en": "User '{user_id}' has been given a role.",
        "ja": "通話に参加中のユーザー「{user_id}」に役職を与えました"
    }
    start_call = {
        "en": "Joined the call.",
        "ja": "通話に参加しました"
    }
    stop_call = {
        "en": "Left the call.",
        "ja": "通話から退出しました"
    }

    # -CHAT
    accept_chat_request = {
        "en": "Chat request accepted.",
        "ja": "チャットリクエストを承認しました"
    }
    create_group_chat = {
        "en": "Group chat '{name}' created.",
        "ja": "グループチャット「{name}」を作成しました"
    }
    create_private_chat = {
        "en": "Private chatroom with '{with_user_id}' created.",
        "ja": "「{with_user_id}」との個人チャットを作成しました"
    }
    delete_background = {
        "en": "Background deleted.",
        "ja": "チャットの背景画像を削除しました"
    }
    delete_message = {
        "en": "Message deleted.",
        "ja": "チャットメッセージを取り消しました"
    }
    edit_chat_room = {
        "en": "Chatroom updated.",
        "ja": "チャットルームを編集しました"
    }
    invite_to_chat = {
        "en": "Chatroom invitation sent.",
        "ja": "チャットルームに招待しました"
    }
    kick_users_from_chat = {
        "en": "Users have been kicked from the chatroom.",
        "ja": "チャットルームからユーザーを追放しました"
    }
    pin_chat = {
        "en": "Pinned chatroom.",
        "ja": "チャットルームをピンしました"
    }
    read_attachment = {
        "en": "Attachment has been read.",
        "ja": "アタッチメントを既読にしました"
    }
    read_message = {
        "en": "Message has been read.",
        "ja": "メッセージを既読にしました"
    }
    read_video_message = {
        "en": "Video message has been read.",
        "ja": "動画のメッセージを既読にしました"
    }
    remove_chat_rooms = {
        "en": "Chat rooms removed.",
        "ja": "チャットルームを削除しました"
    }
    report_chat_room = {
        "en": "Chatroom '{chat_room_id}' has been reported.",
        "ja": "チャットルーム「{chat_room_id}」を通報しました"
    }
    send_media_screenshot_notification = {
        "en": "Media screenshot notification sent.",
        "ja": "メディアのスクリーンショットを通知しました"
    }
    send_message = {
        "en": "Your message has been sent.",
        "ja": "メッセージを送信しました"
    }
    unhide_chat = {
        "en": "Chatroom unhidden.",
        "ja": "チャットルームの非表示を解除しました"
    }
    unpin_chat = {
        "en": "Chatroom unpinned.",
        "ja": "チャットルームを非表示にしました"
    }

    # -GROUP
    accept_moderator_offer = {
        "en": "Moderator offer accepted.",
        "ja": "サークル副管理人の権限オファーを引き受けうけました"
    }
    accept_ownership_offer = {
        "en": "Accepted ownership offer.",
        "ja": "サークル管理人の権限オファーを引き受けました"
    }
    accept_group_join_request = {
        "en": "Accepted group join request.",
        "ja": "サークル参加リクエストを承認しました"
    }
    add_related_groups = {
        "en": "Group added to the related groups",
        "ja": "関連サークルを追加しました"
    }
    ban_group_user = {
        "en": "User has been banned from the group.",
        "ja": "サークルからユーザーを追放しました"
    }
    create_group = {
        "en": "Group created.",
        "ja": "サークルを作成しました"
    }
    create_pin_group = {
        "en": "Pinned post in the group.",
        "ja": "サークルの投稿をピンしました"
    }
    decline_moderator_offer = {
        "en": "Declined moderator offer.",
        "ja": "サークル副管理人の権限オファーを断りました"
    }
    decline_ownership_offer = {
        "en": "Declined ownership offer.",
        "ja": "サークル管理人の権限オファーを断りました"
    }
    decline_group_join_request = {
        "en": "Declined group join request",
        "ja": "サークル参加リクエストを断りました"
    }
    delete_pin_group = {
        "en": "Unpinned group post.",
        "ja": "サークルのピン投稿を解除しました"
    }
    invite_users_to_group = {
        "en": "Group invitation sent.",
        "ja": "サークルにユーザーを招待しました"
    }
    join_group = {
        "en": "Joined the group.",
        "ja": "サークルに参加しました"
    }
    leave_group = {
        "en": "Left the group.",
        "ja": "サークルから脱退しました"
    }
    post_gruop_social_shared = {
        "en": "Group social shared posted.",
        "ja": "サークルのソーシャルシェアを投稿しました"
    }
    remove_group_cover = {
        "en": "Group cover removed.",
        "ja": "サークルのカバー画像を削除しました"
    }
    remove_moderator = {
        "en": "Group moderator removed.",
        "ja": "サークルの副管理人を削除します"
    }
    remove_related_groups = {
        "en": "Related groups removed.",
        "ja": "関連のあるサークルを削除しました"
    }
    report_group = {
        "en": "Group reported.",
        "ja": "サークルを通報しました"
    }
    send_moderator_offers = {
        "en": "Group moderator offer sent.",
        "ja": "サークル副管理人の権限をオファーしました"
    }
    send_ownership_offer = {
        "en": "Group ownership offer sent.",
        "ja": "サークル管理人の権限をオファーしました"
    }
    set_group_title = {
        "en": "Group title set.",
        "ja": "サークルのタイトルを設定しました"
    }
    take_over_group_ownership = {
        "en": "Took over the group ownership of {group_id}",
        "ja": "サークル管理人の権限を引き継ぎました"
    }
    unban_group_member = {
        "en": "Group user has been unbanned.",
        "ja": "特定のサークルメンバーの追放を解除しました"
    }
    update_group = {
        "en": "Group details have been updated.",
        "ja": "サークルを編集しました"
    }
    visit_group = {
        "en": "Visited the group.",
        "ja": "サークルを訪問しました"
    }
    withdraw_moderator_offer = {
        "en": "Group moderator offer withdrawn",
        "ja": "サークル副管理人のオファーを取り消しました"
    }
    withdraw_ownership_offer = {
        "en": "Group ownershipt offer withdrawn",
        "ja": "サークル管理人のオファーを取り消しました"
    }

    # -LOGIN
    change_email = {
        "en": "Your email has been changed.",
        "ja": "メールアドレスを変更しました"
    }
    change_password = {
        "en": "Password has been changed.",
        "ja": "パスワードを変更しました"
    }
    login = {
        "en": "Successfully logged in as '{user_id}'",
        "ja": "ログインしました。ユーザーID: {user_id}"
    }
    logout = {
        "en": "User has Logged out.",
        "ja": "ログアウトしました"
    }
    restore_user = {
        "en": "User restored.",
        "ja": "ユーザーを復元しました"
    }
    revoke_tokens = {
        "en": "Token revoked.",
        "ja": "トークンを無効化しました"
    }
    save_account_with_email = {
        "en": "Account has been save with email.",
        "ja": "メールアドレスでアカウントを保存しました"
    }

    # -MISC
    accept_policy_agreement = {
        "en": "Accepted to policy agreement.",
        "ja": "利用規約を承諾しました"
    }
    generate_sns_thumbnail = {
        "en": "SNS thumbnail generated.",
        "ja": "SNSのサムネイルを生成しました"
    }
    verify_device = {
        "en": "Device verified.",
        "ja": "デバイスを検証しました"
    }
    upload_image = {
        "en": "Image '{filename}{ext}' is uploaded",
        "ja": "画像をアップロードしました。ファイル名: '{filename}{ext}'"
    }

    # -POST
    add_bookmark = {
        "en": "Post added to bookmarks.",
        "ja": "投稿をブックマークに追加しました"
    }
    add_group_highlight_post = {
        "en": "Post added to group highlight.",
        "ja": "投稿をサークルのハイライトに追加しました"
    }
    create_call_post = {
        "en": "Call post created.",
        "ja": "通話の投稿を作成しました"
    }
    create_group_pin_post = {
        "en": "Pinned post in a group.",
        "ja": "サークルの投稿をピンしました"
    }
    create_pin_post = {
        "en": "Pinned post.",
        "ja": "投稿をプロフィールにピンしました"
    }
    create_post = {
        "en": "Post has been created.",
        "ja": "投稿を作成しました"
    }
    create_repost = {
        "en": "Repost has been created.",
        "ja": "投稿を(´∀｀∩)↑age↑しました"
    }
    create_share_post = {
        "en": "Share post has been created.",
        "ja": "シェア投稿を作成しました"  # fix this
    }
    create_thread_post = {
        "en": "Thread post has been created.",
        "ja": "スレッドの投稿を作成しました"
    }
    delete_all_post = {
        "en": "Post deletion requested.",
        "ja": "投稿をすべて削除するリクエストを送信しました"
    }
    delete_all_post_not_found = {
        "en": "Post not found.",
        "ja": "削除する投稿が見つかりません"
    }
    delete_group_pin_post = {
        "en": "Unpinned post in a group.",
        "ja": "グループのピン投稿を解除しました"
    }
    delete_pin_post = {
        "en": "Unpinned post.",
        "ja": "ピン投稿を解除しました"
    }
    like_posts = {
        "en": "Posts liked.",
        "ja": "投稿にいいねしました"
    }
    remove_bookmark = {
        "en": "Bookmark has been removed.",
        "ja": "投稿をブックマークから削除しました"
    }
    remove_group_highlight_post = {
        "en": "Group hightlight post removed.",
        "ja": "投稿をサークルのハイライトから削除しました"
    }
    remove_posts = {
        "en": "Posts have been removed.",
        "ja": "投稿を削除しました"
    }
    report_post = {
        "en": "Post has been reported.",
        "ja": "投稿を通報しました"
    }
    unlike_post = {
        "en": "Post unliked.",
        "ja": "投稿をアンライクしました"
    }
    update_post = {
        "en": "Post has been updated.",
        "ja": "投稿を編集しました"
    }
    update_recommendation_feedback = {
        "en": "Recommendation feedback updated.",
        "ja": "おすすめのフィードバックを更新しました"
    }
    validate_post = {
        "en": "Post validated.",
        "ja": "投稿を検証しました"
    }
    view_video = {
        "en": "Viewed the video",
        "ja": "動画を視聴しました"
    }
    vote_survey = {
        "en": "Survey voted.",
        "ja": "投票しました"
    }

    # -REVIEW
    create_review = {
        "en": "Review has been sent to {user_id}.",
        "ja": "実行完了"
    }
    create_reviews = {
        "en": "Reviews have been sent to multiple users.",
        "ja": "実行完了"
    }
    delete_reviews = {
        "en": "Reviews have been deleted.",
        "ja": "実行完了"
    }
    pin_review = {
        "en": "Pinned the review.",
        "ja": "実行完了"
    }
    unpin_review = {
        "en": "Unpinned the review.",
        "ja": "実行完了"
    }

    # -THREAD
    add_post_to_thread = {
        "en": "Post '{post_id}' added to the thread '{thread_id}'.",
        "ja": "投稿「{post_id}」をスレッド「{thread_id}」に追加しました"
    }
    convert_post_to_thread = {
        "en": "Post has been converted to a thread.",
        "ja": "投稿をスレッドに変換しました"
    }
    create_thread = {
        "en": "A new thread has been created.",
        "ja": "新しいスレッドを作成しました"
    }
    join_thread = {
        "en": "Joined the thread '{thread_id}'.",
        "ja": "スレッド「{thread_id}」に参加しました"
    }
    leave_thread = {
        "en": "Left the thread.",
        "ja": "スレッドから脱退しました"
    }
    remove_thread = {
        "en": "Thread '{thread_id}' has been removed.",
        "ja": "スレッドを作成しました。スレッドID: {thread_id}"
    }
    update_thread = {
        "en": "Thread '{thread_id}' has been updated.",
        "ja": "スレッド「{thread_id}」を編集しました"
    }

    # -USERS
    create_user = {
        "en": "A new account has been created.",
        "ja": "アカウントを作成しました"
    }
    delete_contact_friends = {
        "en": "Contact friends deleted.",
        "ja": "連絡先の友人を削除しました"
    }
    delete_footprint = {
        "en": "Footprint deleted.",
        "ja": "足跡を削除しました"
    }
    destroy_user = {
        "en": "User deleted.",
        "ja": "ユーザーを削除しました"
    }
    follow_user = {
        "en": "Followed the user '{user_id}'.",
        "ja": "ユーザー「{user_id}」をフォローしました"
    }
    follow_users = {
        "en": "Followed multiple users.",
        "ja": "複数のユーザーをフォローしました"
    }
    post_social_shared = {
        "en": "Posted social shared post.",
        "ja": "ソーシャルシェア投稿を投稿しました"
    }
    record_app_review_status = {
        "en": "App review status recored.",
        "ja": "アプリのレビュー状況を記録しました"
    }
    reduce_kenta_penalty = {
        "en": "Penalty reduced.",
        "ja": "ペナルティーを軽減しました"
    }
    refresh_counter = {
        "en": "Refreshed the counter",
        "ja": "カウンターをリフレッシュしました"
    }
    remove_user_avatar = {
        "en": "User avatar has been removed.",
        "ja": "アイコンを削除しました"
    }
    remove_user_cover = {
        "en": "User cover has been removed.",
        "ja": "カバー画像を削除しました"
    }
    report_user = {
        "en": "Reported the user '{user_id}'",
        "ja": "ユーザー「{user_id}」を通報しました"
    }
    reset_password = {
        "en": "Reset the password.",
        "ja": "パスワードをリセットしました"
    }
    set_additional_setting_enabled = {
        "en": "Additional settings have been enabled.",
        "ja": "Additional settings have been enabled."
    }
    set_follow_permission_enabled = {
        "en": "Follow permission enabled.",
        "ja": "フォローを許可性にしました"
    }
    set_setting_follow_recommendation_enabled = {
        "en": "Follow recommendation enabled.",
        "ja": "Follow recommendation enabled."
    }
    take_action_follow_request = {
        "en": "Took action follow request.",
        "ja": "Took action follow request."
    }
    turn_on_hima = {
        "en": "Turned on hima now.",
        "ja": "「ひまなう」しました"
    }
    unfollow_user = {
        "en": "Unfollowed the user '{user_id}'",
        "ja": "ユーザー「{user_id}」をアンフォローしました"
    }
    update_invite_contact_status = {
        "en": "Invite contact status updated.",
        "ja": "Invite contact status updated."
    }
    update_language = {
        "en": "Language updated.",
        "ja": "言語を更新しました"
    }
    update_user = {
        "en": "User profile has been updated.",
        "ja": "プロフィールを編集しました"
    }
    upload_twitter_friend_ids = {
        "en": "Uploaded Twitter friend ids.",
        "ja": "TwitterのフレンドIDをアップロードしました"
    }
    block_user = {
        "en": "Blocked the user '{user_id}'",
        "ja": "ユーザー「{user_id}」をブロックしました"
    }
    unblock_user = {
        "en": "Unblocked the user '{user_id}'",
        "ja": "ユーザー「{user_id}」のブロックを解除しました"
    }
    hide_user = {
        "en": "User '{user_id}' is hidden",
        "ja": "ユーザーを非表示にしました"
    }
    unhide_users = {
        "en": "Unhidden users",
        "ja": "ユーザーの非表示を解除しました"
    }
