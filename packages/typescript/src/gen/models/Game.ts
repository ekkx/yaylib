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
import type { PlatformDetails } from './PlatformDetails';
import {
    PlatformDetailsFromJSON,
    PlatformDetailsFromJSONTyped,
    PlatformDetailsToJSON,
    PlatformDetailsToJSONTyped,
} from './PlatformDetails';

/**
 * 
 * @export
 * @interface Game
 */
export interface Game {
    /**
     * 
     * @type {string}
     * @memberof Game
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Game
     */
    id?: number | null;
    /**
     * 
     * @type {PlatformDetails}
     * @memberof Game
     */
    platformDetails?: PlatformDetails | null;
    /**
     * 
     * @type {string}
     * @memberof Game
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Game
     */
    type?: string | null;
}

/**
 * Check if a given object implements the Game interface.
 */
export function instanceOfGame(value: object): value is Game {
    return true;
}

export function GameFromJSON(json: any): Game {
    return GameFromJSONTyped(json, false);
}

export function GameFromJSONTyped(json: any, ignoreDiscriminator: boolean): Game {
    if (json == null) {
        return json;
    }
    return {
        
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'platformDetails': json['platform_details'] == null ? undefined : PlatformDetailsFromJSON(json['platform_details']),
        'title': json['title'] == null ? undefined : json['title'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function GameToJSON(json: any): Game {
    return GameToJSONTyped(json, false);
}

export function GameToJSONTyped(value?: Game | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'platform_details': PlatformDetailsToJSON(value['platformDetails']),
        'title': value['title'],
        'type': value['type'],
    };
}

