/* tslint:disable */
/* eslint-disable */
/**
 * Yay! API
 * No description provided
 *
 * Schema version: 4.26.1
 * 
 *
 * NOTE: Code generated; DO NOT EDIT.
 * 
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Settings
 */
export interface Settings {
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    ageRestrictedOnReview?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    allowReposts?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    cautionUserChat?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    followingOnlyCallInvite?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    followingOnlyGroupInvite?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    followingRestrictedOnReview?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideActiveCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideGiftsReceived?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideHotPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideOnInvitable?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideOnlineStatus?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideReplyFollowingTimeline?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideReplyPublicTimeline?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    hideVip?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    invisibleOnUserSearch?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    noReplyGroupTimeline?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationBirthdayToFollowers?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationBulkCallInvite?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationCallInvite?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationChat?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationChatDelete?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationContactFriend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationDailyQuest?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationDailySummary?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollow?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowAccept?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowRequest?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowerConferenceCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowerCreateGroup?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowingBirthdateOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowingPostAfterBreak?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFollowingsInCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationFootprint?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGiftReceive?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupAccept?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupConferenceCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupInvite?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupJoin?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupMessageTagAll?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupModerator?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationGroupRequest?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationHimaNow?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationInviteeCreateAccount?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationLatestNews?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationLike?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationMessageTag?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationPopularPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationProfileScreenshot?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationReply?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationRepost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationReview?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationSecurityWarning?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    notificationTwitterFriend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    privacyMode?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    privatePost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    privateUserTimeline?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    showTotalGiftsReceivedCount?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    vipInvisibleFootprintMode?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Settings
     */
    visibleOnSnsFriendRecommendation?: boolean | null;
}

/**
 * Check if a given object implements the Settings interface.
 */
export function instanceOfSettings(value: object): value is Settings {
    return true;
}

export function SettingsFromJSON(json: any): Settings {
    return SettingsFromJSONTyped(json, false);
}

export function SettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): Settings {
    if (json == null) {
        return json;
    }
    return {
        
        'ageRestrictedOnReview': json['age_restricted_on_review'] == null ? undefined : json['age_restricted_on_review'],
        'allowReposts': json['allow_reposts'] == null ? undefined : json['allow_reposts'],
        'cautionUserChat': json['caution_user_chat'] == null ? undefined : json['caution_user_chat'],
        'followingOnlyCallInvite': json['following_only_call_invite'] == null ? undefined : json['following_only_call_invite'],
        'followingOnlyGroupInvite': json['following_only_group_invite'] == null ? undefined : json['following_only_group_invite'],
        'followingRestrictedOnReview': json['following_restricted_on_review'] == null ? undefined : json['following_restricted_on_review'],
        'hideActiveCall': json['hide_active_call'] == null ? undefined : json['hide_active_call'],
        'hideGiftsReceived': json['hide_gifts_received'] == null ? undefined : json['hide_gifts_received'],
        'hideHotPost': json['hide_hot_post'] == null ? undefined : json['hide_hot_post'],
        'hideOnInvitable': json['hide_on_invitable'] == null ? undefined : json['hide_on_invitable'],
        'hideOnlineStatus': json['hide_online_status'] == null ? undefined : json['hide_online_status'],
        'hideReplyFollowingTimeline': json['hide_reply_following_timeline'] == null ? undefined : json['hide_reply_following_timeline'],
        'hideReplyPublicTimeline': json['hide_reply_public_timeline'] == null ? undefined : json['hide_reply_public_timeline'],
        'hideVip': json['hide_vip'] == null ? undefined : json['hide_vip'],
        'invisibleOnUserSearch': json['invisible_on_user_search'] == null ? undefined : json['invisible_on_user_search'],
        'noReplyGroupTimeline': json['no_reply_group_timeline'] == null ? undefined : json['no_reply_group_timeline'],
        'notificationBirthdayToFollowers': json['notification_birthday_to_followers'] == null ? undefined : json['notification_birthday_to_followers'],
        'notificationBulkCallInvite': json['notification_bulk_call_invite'] == null ? undefined : json['notification_bulk_call_invite'],
        'notificationCallInvite': json['notification_call_invite'] == null ? undefined : json['notification_call_invite'],
        'notificationChat': json['notification_chat'] == null ? undefined : json['notification_chat'],
        'notificationChatDelete': json['notification_chat_delete'] == null ? undefined : json['notification_chat_delete'],
        'notificationContactFriend': json['notification_contact_friend'] == null ? undefined : json['notification_contact_friend'],
        'notificationDailyQuest': json['notification_daily_quest'] == null ? undefined : json['notification_daily_quest'],
        'notificationDailySummary': json['notification_daily_summary'] == null ? undefined : json['notification_daily_summary'],
        'notificationFollow': json['notification_follow'] == null ? undefined : json['notification_follow'],
        'notificationFollowAccept': json['notification_follow_accept'] == null ? undefined : json['notification_follow_accept'],
        'notificationFollowRequest': json['notification_follow_request'] == null ? undefined : json['notification_follow_request'],
        'notificationFollowerConferenceCall': json['notification_follower_conference_call'] == null ? undefined : json['notification_follower_conference_call'],
        'notificationFollowerCreateGroup': json['notification_follower_create_group'] == null ? undefined : json['notification_follower_create_group'],
        'notificationFollowingBirthdateOn': json['notification_following_birthdate_on'] == null ? undefined : json['notification_following_birthdate_on'],
        'notificationFollowingPostAfterBreak': json['notification_following_post_after_break'] == null ? undefined : json['notification_following_post_after_break'],
        'notificationFollowingsInCall': json['notification_followings_in_call'] == null ? undefined : json['notification_followings_in_call'],
        'notificationFootprint': json['notification_footprint'] == null ? undefined : json['notification_footprint'],
        'notificationGiftReceive': json['notification_gift_receive'] == null ? undefined : json['notification_gift_receive'],
        'notificationGroupAccept': json['notification_group_accept'] == null ? undefined : json['notification_group_accept'],
        'notificationGroupConferenceCall': json['notification_group_conference_call'] == null ? undefined : json['notification_group_conference_call'],
        'notificationGroupInvite': json['notification_group_invite'] == null ? undefined : json['notification_group_invite'],
        'notificationGroupJoin': json['notification_group_join'] == null ? undefined : json['notification_group_join'],
        'notificationGroupMessageTagAll': json['notification_group_message_tag_all'] == null ? undefined : json['notification_group_message_tag_all'],
        'notificationGroupModerator': json['notification_group_moderator'] == null ? undefined : json['notification_group_moderator'],
        'notificationGroupPost': json['notification_group_post'] == null ? undefined : json['notification_group_post'],
        'notificationGroupRequest': json['notification_group_request'] == null ? undefined : json['notification_group_request'],
        'notificationHimaNow': json['notification_hima_now'] == null ? undefined : json['notification_hima_now'],
        'notificationInviteeCreateAccount': json['notification_invitee_create_account'] == null ? undefined : json['notification_invitee_create_account'],
        'notificationLatestNews': json['notification_latest_news'] == null ? undefined : json['notification_latest_news'],
        'notificationLike': json['notification_like'] == null ? undefined : json['notification_like'],
        'notificationMessageTag': json['notification_message_tag'] == null ? undefined : json['notification_message_tag'],
        'notificationPopularPost': json['notification_popular_post'] == null ? undefined : json['notification_popular_post'],
        'notificationProfileScreenshot': json['notification_profile_screenshot'] == null ? undefined : json['notification_profile_screenshot'],
        'notificationReply': json['notification_reply'] == null ? undefined : json['notification_reply'],
        'notificationRepost': json['notification_repost'] == null ? undefined : json['notification_repost'],
        'notificationReview': json['notification_review'] == null ? undefined : json['notification_review'],
        'notificationSecurityWarning': json['notification_security_warning'] == null ? undefined : json['notification_security_warning'],
        'notificationTwitterFriend': json['notification_twitter_friend'] == null ? undefined : json['notification_twitter_friend'],
        'privacyMode': json['privacy_mode'] == null ? undefined : json['privacy_mode'],
        'privatePost': json['private_post'] == null ? undefined : json['private_post'],
        'privateUserTimeline': json['private_user_timeline'] == null ? undefined : json['private_user_timeline'],
        'showTotalGiftsReceivedCount': json['show_total_gifts_received_count'] == null ? undefined : json['show_total_gifts_received_count'],
        'vipInvisibleFootprintMode': json['vip_invisible_footprint_mode'] == null ? undefined : json['vip_invisible_footprint_mode'],
        'visibleOnSnsFriendRecommendation': json['visible_on_sns_friend_recommendation'] == null ? undefined : json['visible_on_sns_friend_recommendation'],
    };
}

export function SettingsToJSON(json: any): Settings {
    return SettingsToJSONTyped(json, false);
}

export function SettingsToJSONTyped(value?: Settings | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'age_restricted_on_review': value['ageRestrictedOnReview'],
        'allow_reposts': value['allowReposts'],
        'caution_user_chat': value['cautionUserChat'],
        'following_only_call_invite': value['followingOnlyCallInvite'],
        'following_only_group_invite': value['followingOnlyGroupInvite'],
        'following_restricted_on_review': value['followingRestrictedOnReview'],
        'hide_active_call': value['hideActiveCall'],
        'hide_gifts_received': value['hideGiftsReceived'],
        'hide_hot_post': value['hideHotPost'],
        'hide_on_invitable': value['hideOnInvitable'],
        'hide_online_status': value['hideOnlineStatus'],
        'hide_reply_following_timeline': value['hideReplyFollowingTimeline'],
        'hide_reply_public_timeline': value['hideReplyPublicTimeline'],
        'hide_vip': value['hideVip'],
        'invisible_on_user_search': value['invisibleOnUserSearch'],
        'no_reply_group_timeline': value['noReplyGroupTimeline'],
        'notification_birthday_to_followers': value['notificationBirthdayToFollowers'],
        'notification_bulk_call_invite': value['notificationBulkCallInvite'],
        'notification_call_invite': value['notificationCallInvite'],
        'notification_chat': value['notificationChat'],
        'notification_chat_delete': value['notificationChatDelete'],
        'notification_contact_friend': value['notificationContactFriend'],
        'notification_daily_quest': value['notificationDailyQuest'],
        'notification_daily_summary': value['notificationDailySummary'],
        'notification_follow': value['notificationFollow'],
        'notification_follow_accept': value['notificationFollowAccept'],
        'notification_follow_request': value['notificationFollowRequest'],
        'notification_follower_conference_call': value['notificationFollowerConferenceCall'],
        'notification_follower_create_group': value['notificationFollowerCreateGroup'],
        'notification_following_birthdate_on': value['notificationFollowingBirthdateOn'],
        'notification_following_post_after_break': value['notificationFollowingPostAfterBreak'],
        'notification_followings_in_call': value['notificationFollowingsInCall'],
        'notification_footprint': value['notificationFootprint'],
        'notification_gift_receive': value['notificationGiftReceive'],
        'notification_group_accept': value['notificationGroupAccept'],
        'notification_group_conference_call': value['notificationGroupConferenceCall'],
        'notification_group_invite': value['notificationGroupInvite'],
        'notification_group_join': value['notificationGroupJoin'],
        'notification_group_message_tag_all': value['notificationGroupMessageTagAll'],
        'notification_group_moderator': value['notificationGroupModerator'],
        'notification_group_post': value['notificationGroupPost'],
        'notification_group_request': value['notificationGroupRequest'],
        'notification_hima_now': value['notificationHimaNow'],
        'notification_invitee_create_account': value['notificationInviteeCreateAccount'],
        'notification_latest_news': value['notificationLatestNews'],
        'notification_like': value['notificationLike'],
        'notification_message_tag': value['notificationMessageTag'],
        'notification_popular_post': value['notificationPopularPost'],
        'notification_profile_screenshot': value['notificationProfileScreenshot'],
        'notification_reply': value['notificationReply'],
        'notification_repost': value['notificationRepost'],
        'notification_review': value['notificationReview'],
        'notification_security_warning': value['notificationSecurityWarning'],
        'notification_twitter_friend': value['notificationTwitterFriend'],
        'privacy_mode': value['privacyMode'],
        'private_post': value['privatePost'],
        'private_user_timeline': value['privateUserTimeline'],
        'show_total_gifts_received_count': value['showTotalGiftsReceivedCount'],
        'vip_invisible_footprint_mode': value['vipInvisibleFootprintMode'],
        'visible_on_sns_friend_recommendation': value['visibleOnSnsFriendRecommendation'],
    };
}

