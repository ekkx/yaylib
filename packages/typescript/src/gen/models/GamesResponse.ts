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
import type { RealmGame } from './RealmGame';
import {
    RealmGameFromJSON,
    RealmGameFromJSONTyped,
    RealmGameToJSON,
    RealmGameToJSONTyped,
} from './RealmGame';

/**
 * 
 * @export
 * @interface GamesResponse
 */
export interface GamesResponse {
    /**
     * 
     * @type {number}
     * @memberof GamesResponse
     */
    fromId?: number | null;
    /**
     * 
     * @type {Array<RealmGame>}
     * @memberof GamesResponse
     */
    games?: Array<RealmGame> | null;
}

/**
 * Check if a given object implements the GamesResponse interface.
 */
export function instanceOfGamesResponse(value: object): value is GamesResponse {
    return true;
}

export function GamesResponseFromJSON(json: any): GamesResponse {
    return GamesResponseFromJSONTyped(json, false);
}

export function GamesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GamesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'fromId': json['from_id'] == null ? undefined : json['from_id'],
        'games': json['games'] == null ? undefined : ((json['games'] as Array<any>).map(RealmGameFromJSON)),
    };
}

export function GamesResponseToJSON(json: any): GamesResponse {
    return GamesResponseToJSONTyped(json, false);
}

export function GamesResponseToJSONTyped(value?: GamesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'from_id': value['fromId'],
        'games': value['games'] == null ? undefined : ((value['games'] as Array<any>).map(RealmGameToJSON)),
    };
}

