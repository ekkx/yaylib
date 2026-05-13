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
 * @interface CreateFriendShipsResponse
 */
export interface CreateFriendShipsResponse {
    /**
     * 
     * @type {boolean}
     * @memberof CreateFriendShipsResponse
     */
    following?: boolean | null;
}

/**
 * Check if a given object implements the CreateFriendShipsResponse interface.
 */
export function instanceOfCreateFriendShipsResponse(value: object): value is CreateFriendShipsResponse {
    return true;
}

export function CreateFriendShipsResponseFromJSON(json: any): CreateFriendShipsResponse {
    return CreateFriendShipsResponseFromJSONTyped(json, false);
}

export function CreateFriendShipsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateFriendShipsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'following': json['following'] == null ? undefined : json['following'],
    };
}

export function CreateFriendShipsResponseToJSON(json: any): CreateFriendShipsResponse {
    return CreateFriendShipsResponseToJSONTyped(json, false);
}

export function CreateFriendShipsResponseToJSONTyped(value?: CreateFriendShipsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'following': value['following'],
    };
}

