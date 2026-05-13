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
 * @interface SnsInfo
 */
export interface SnsInfo {
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    biography?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    gender?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    nickname?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    profileImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SnsInfo
     */
    uid?: string | null;
}

/**
 * Check if a given object implements the SnsInfo interface.
 */
export function instanceOfSnsInfo(value: object): value is SnsInfo {
    return true;
}

export function SnsInfoFromJSON(json: any): SnsInfo {
    return SnsInfoFromJSONTyped(json, false);
}

export function SnsInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): SnsInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'biography': json['biography'] == null ? undefined : json['biography'],
        'gender': json['gender'] == null ? undefined : json['gender'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'profileImage': json['profile_image'] == null ? undefined : json['profile_image'],
        'type': json['type'] == null ? undefined : json['type'],
        'uid': json['uid'] == null ? undefined : json['uid'],
    };
}

export function SnsInfoToJSON(json: any): SnsInfo {
    return SnsInfoToJSONTyped(json, false);
}

export function SnsInfoToJSONTyped(value?: SnsInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'biography': value['biography'],
        'gender': value['gender'],
        'nickname': value['nickname'],
        'profile_image': value['profileImage'],
        'type': value['type'],
        'uid': value['uid'],
    };
}

