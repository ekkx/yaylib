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
 * @interface PlatformDetails
 */
export interface PlatformDetails {
    /**
     * 
     * @type {string}
     * @memberof PlatformDetails
     */
    affiliateUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PlatformDetails
     */
    packageId?: string | null;
}

/**
 * Check if a given object implements the PlatformDetails interface.
 */
export function instanceOfPlatformDetails(value: object): value is PlatformDetails {
    return true;
}

export function PlatformDetailsFromJSON(json: any): PlatformDetails {
    return PlatformDetailsFromJSONTyped(json, false);
}

export function PlatformDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlatformDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'affiliateUrl': json['affiliate_url'] == null ? undefined : json['affiliate_url'],
        'packageId': json['package_id'] == null ? undefined : json['package_id'],
    };
}

export function PlatformDetailsToJSON(json: any): PlatformDetails {
    return PlatformDetailsToJSONTyped(json, false);
}

export function PlatformDetailsToJSONTyped(value?: PlatformDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'affiliate_url': value['affiliateUrl'],
        'package_id': value['packageId'],
    };
}

