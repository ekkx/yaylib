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
import type { RealmPlatformDetails } from './RealmPlatformDetails';
import {
    RealmPlatformDetailsFromJSON,
    RealmPlatformDetailsFromJSONTyped,
    RealmPlatformDetailsToJSON,
    RealmPlatformDetailsToJSONTyped,
} from './RealmPlatformDetails';

/**
 * 
 * @export
 * @interface RealmGame
 */
export interface RealmGame {
    /**
     * 
     * @type {string}
     * @memberof RealmGame
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmGame
     */
    id?: number | null;
    /**
     * 
     * @type {RealmPlatformDetails}
     * @memberof RealmGame
     */
    platformDetails?: RealmPlatformDetails | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGame
     */
    title?: string | null;
}

/**
 * Check if a given object implements the RealmGame interface.
 */
export function instanceOfRealmGame(value: object): value is RealmGame {
    return true;
}

export function RealmGameFromJSON(json: any): RealmGame {
    return RealmGameFromJSONTyped(json, false);
}

export function RealmGameFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmGame {
    if (json == null) {
        return json;
    }
    return {
        
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'platformDetails': json['platform_details'] == null ? undefined : RealmPlatformDetailsFromJSON(json['platform_details']),
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function RealmGameToJSON(json: any): RealmGame {
    return RealmGameToJSONTyped(json, false);
}

export function RealmGameToJSONTyped(value?: RealmGame | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'platform_details': RealmPlatformDetailsToJSON(value['platformDetails']),
        'title': value['title'],
    };
}

