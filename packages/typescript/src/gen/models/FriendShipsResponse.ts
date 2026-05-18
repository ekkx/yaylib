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
import type { Relationship } from './Relationship.js';
import {
    RelationshipFromJSON,
    RelationshipFromJSONTyped,
    RelationshipToJSON,
    RelationshipToJSONTyped,
} from './Relationship.js';

/**
 * 
 * @export
 * @interface FriendShipsResponse
 */
export interface FriendShipsResponse {
    /**
     * 
     * @type {Relationship}
     * @memberof FriendShipsResponse
     */
    relationship?: Relationship | null;
}

/**
 * Check if a given object implements the FriendShipsResponse interface.
 */
export function instanceOfFriendShipsResponse(value: object): value is FriendShipsResponse {
    return true;
}

export function FriendShipsResponseFromJSON(json: any): FriendShipsResponse {
    return FriendShipsResponseFromJSONTyped(json, false);
}

export function FriendShipsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FriendShipsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'relationship': json['relationship'] == null ? undefined : RelationshipFromJSON(json['relationship']),
    };
}

export function FriendShipsResponseToJSON(json: any): FriendShipsResponse {
    return FriendShipsResponseToJSONTyped(json, false);
}

export function FriendShipsResponseToJSONTyped(value?: FriendShipsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'relationship': RelationshipToJSON(value['relationship']),
    };
}

