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
 * @interface Setting
 */
export interface Setting {
    /**
     * 
     * @type {boolean}
     * @memberof Setting
     */
    notificationGroupJoin?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Setting
     */
    notificationGroupMessageTagAll?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Setting
     */
    notificationGroupPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Setting
     */
    notificationGroupRequest?: boolean | null;
}

/**
 * Check if a given object implements the Setting interface.
 */
export function instanceOfSetting(value: object): value is Setting {
    return true;
}

export function SettingFromJSON(json: any): Setting {
    return SettingFromJSONTyped(json, false);
}

export function SettingFromJSONTyped(json: any, ignoreDiscriminator: boolean): Setting {
    if (json == null) {
        return json;
    }
    return {
        
        'notificationGroupJoin': json['notification_group_join'] == null ? undefined : json['notification_group_join'],
        'notificationGroupMessageTagAll': json['notification_group_message_tag_all'] == null ? undefined : json['notification_group_message_tag_all'],
        'notificationGroupPost': json['notification_group_post'] == null ? undefined : json['notification_group_post'],
        'notificationGroupRequest': json['notification_group_request'] == null ? undefined : json['notification_group_request'],
    };
}

export function SettingToJSON(json: any): Setting {
    return SettingToJSONTyped(json, false);
}

export function SettingToJSONTyped(value?: Setting | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'notification_group_join': value['notificationGroupJoin'],
        'notification_group_message_tag_all': value['notificationGroupMessageTagAll'],
        'notification_group_post': value['notificationGroupPost'],
        'notification_group_request': value['notificationGroupRequest'],
    };
}

