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
 * @interface UserSetting
 */
export interface UserSetting {
    /**
     * 
     * @type {boolean}
     * @memberof UserSetting
     */
    notificationChat?: boolean | null;
}

/**
 * Check if a given object implements the UserSetting interface.
 */
export function instanceOfUserSetting(value: object): value is UserSetting {
    return true;
}

export function UserSettingFromJSON(json: any): UserSetting {
    return UserSettingFromJSONTyped(json, false);
}

export function UserSettingFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSetting {
    if (json == null) {
        return json;
    }
    return {
        
        'notificationChat': json['notification_chat'] == null ? undefined : json['notification_chat'],
    };
}

export function UserSettingToJSON(json: any): UserSetting {
    return UserSettingToJSONTyped(json, false);
}

export function UserSettingToJSONTyped(value?: UserSetting | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'notification_chat': value['notificationChat'],
    };
}

