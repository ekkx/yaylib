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
import type { Setting } from './Setting';
import {
    SettingFromJSON,
    SettingFromJSONTyped,
    SettingToJSON,
    SettingToJSONTyped,
} from './Setting';

/**
 * 
 * @export
 * @interface GroupNotificationSettingsResponse
 */
export interface GroupNotificationSettingsResponse {
    /**
     * 
     * @type {Setting}
     * @memberof GroupNotificationSettingsResponse
     */
    setting?: Setting | null;
}

/**
 * Check if a given object implements the GroupNotificationSettingsResponse interface.
 */
export function instanceOfGroupNotificationSettingsResponse(value: object): value is GroupNotificationSettingsResponse {
    return true;
}

export function GroupNotificationSettingsResponseFromJSON(json: any): GroupNotificationSettingsResponse {
    return GroupNotificationSettingsResponseFromJSONTyped(json, false);
}

export function GroupNotificationSettingsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupNotificationSettingsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'setting': json['setting'] == null ? undefined : SettingFromJSON(json['setting']),
    };
}

export function GroupNotificationSettingsResponseToJSON(json: any): GroupNotificationSettingsResponse {
    return GroupNotificationSettingsResponseToJSONTyped(json, false);
}

export function GroupNotificationSettingsResponseToJSONTyped(value?: GroupNotificationSettingsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'setting': SettingToJSON(value['setting']),
    };
}

