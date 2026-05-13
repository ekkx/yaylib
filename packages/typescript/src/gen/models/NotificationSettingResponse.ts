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
import type { UserSetting } from './UserSetting';
import {
    UserSettingFromJSON,
    UserSettingFromJSONTyped,
    UserSettingToJSON,
    UserSettingToJSONTyped,
} from './UserSetting';

/**
 * 
 * @export
 * @interface NotificationSettingResponse
 */
export interface NotificationSettingResponse {
    /**
     * 
     * @type {UserSetting}
     * @memberof NotificationSettingResponse
     */
    setting?: UserSetting | null;
}

/**
 * Check if a given object implements the NotificationSettingResponse interface.
 */
export function instanceOfNotificationSettingResponse(value: object): value is NotificationSettingResponse {
    return true;
}

export function NotificationSettingResponseFromJSON(json: any): NotificationSettingResponse {
    return NotificationSettingResponseFromJSONTyped(json, false);
}

export function NotificationSettingResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): NotificationSettingResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'setting': json['setting'] == null ? undefined : UserSettingFromJSON(json['setting']),
    };
}

export function NotificationSettingResponseToJSON(json: any): NotificationSettingResponse {
    return NotificationSettingResponseToJSONTyped(json, false);
}

export function NotificationSettingResponseToJSONTyped(value?: NotificationSettingResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'setting': UserSettingToJSON(value['setting']),
    };
}

