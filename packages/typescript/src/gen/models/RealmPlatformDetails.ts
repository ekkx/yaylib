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
 * @interface RealmPlatformDetails
 */
export interface RealmPlatformDetails {
    /**
     * 
     * @type {string}
     * @memberof RealmPlatformDetails
     */
    affiliateUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmPlatformDetails
     */
    packageId?: string | null;
}

/**
 * Check if a given object implements the RealmPlatformDetails interface.
 */
export function instanceOfRealmPlatformDetails(value: object): value is RealmPlatformDetails {
    return true;
}

export function RealmPlatformDetailsFromJSON(json: any): RealmPlatformDetails {
    return RealmPlatformDetailsFromJSONTyped(json, false);
}

export function RealmPlatformDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmPlatformDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'affiliateUrl': json['affiliate_url'] == null ? undefined : json['affiliate_url'],
        'packageId': json['package_id'] == null ? undefined : json['package_id'],
    };
}

export function RealmPlatformDetailsToJSON(json: any): RealmPlatformDetails {
    return RealmPlatformDetailsToJSONTyped(json, false);
}

export function RealmPlatformDetailsToJSONTyped(value?: RealmPlatformDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'affiliate_url': value['affiliateUrl'],
        'package_id': value['packageId'],
    };
}

