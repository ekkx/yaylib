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

import { mapValues } from '../runtime.js';
/**
 * 
 * @export
 * @interface TimelineSettings
 */
export interface TimelineSettings {
    /**
     * 
     * @type {boolean}
     * @memberof TimelineSettings
     */
    favesFilter?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineSettings
     */
    hideHotPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineSettings
     */
    hideReplyFollowingTimeline?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineSettings
     */
    hideReplyPublicTimeline?: boolean | null;
}

/**
 * Check if a given object implements the TimelineSettings interface.
 */
export function instanceOfTimelineSettings(value: object): value is TimelineSettings {
    return true;
}

export function TimelineSettingsFromJSON(json: any): TimelineSettings {
    return TimelineSettingsFromJSONTyped(json, false);
}

export function TimelineSettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): TimelineSettings {
    if (json == null) {
        return json;
    }
    return {
        
        'favesFilter': json['faves_filter'] == null ? undefined : json['faves_filter'],
        'hideHotPost': json['hide_hot_post'] == null ? undefined : json['hide_hot_post'],
        'hideReplyFollowingTimeline': json['hide_reply_following_timeline'] == null ? undefined : json['hide_reply_following_timeline'],
        'hideReplyPublicTimeline': json['hide_reply_public_timeline'] == null ? undefined : json['hide_reply_public_timeline'],
    };
}

export function TimelineSettingsToJSON(json: any): TimelineSettings {
    return TimelineSettingsToJSONTyped(json, false);
}

export function TimelineSettingsToJSONTyped(value?: TimelineSettings | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'faves_filter': value['favesFilter'],
        'hide_hot_post': value['hideHotPost'],
        'hide_reply_following_timeline': value['hideReplyFollowingTimeline'],
        'hide_reply_public_timeline': value['hideReplyPublicTimeline'],
    };
}

