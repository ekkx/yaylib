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
 * @interface AdditionalSetting
 */
export interface AdditionalSetting {
    /**
     * 
     * @type {string}
     * @memberof AdditionalSetting
     */
    apiId?: string | null;
    /**
     * 
     * @type {object}
     * @memberof AdditionalSetting
     */
    isOn?: object | null;
    /**
     * 
     * @type {object}
     * @memberof AdditionalSetting
     */
    isOnFromUser?: object | null;
}

/**
 * Check if a given object implements the AdditionalSetting interface.
 */
export function instanceOfAdditionalSetting(value: object): value is AdditionalSetting {
    return true;
}

export function AdditionalSettingFromJSON(json: any): AdditionalSetting {
    return AdditionalSettingFromJSONTyped(json, false);
}

export function AdditionalSettingFromJSONTyped(json: any, ignoreDiscriminator: boolean): AdditionalSetting {
    if (json == null) {
        return json;
    }
    return {
        
        'apiId': json['api_id'] == null ? undefined : json['api_id'],
        'isOn': json['is_on'] == null ? undefined : json['is_on'],
        'isOnFromUser': json['is_on_from_user'] == null ? undefined : json['is_on_from_user'],
    };
}

export function AdditionalSettingToJSON(json: any): AdditionalSetting {
    return AdditionalSettingToJSONTyped(json, false);
}

export function AdditionalSettingToJSONTyped(value?: AdditionalSetting | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_id': value['apiId'],
        'is_on': value['isOn'],
        'is_on_from_user': value['isOnFromUser'],
    };
}

