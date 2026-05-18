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
import type { UserUserDTO } from './UserUserDTO.js';
import {
    UserUserDTOFromJSON,
    UserUserDTOFromJSONTyped,
    UserUserDTOToJSON,
    UserUserDTOToJSONTyped,
} from './UserUserDTO.js';

/**
 * 
 * @export
 * @interface FootprintDTO
 */
export interface FootprintDTO {
    /**
     * 
     * @type {number}
     * @memberof FootprintDTO
     */
    id?: number | null;
    /**
     * 
     * @type {UserUserDTO}
     * @memberof FootprintDTO
     */
    user?: UserUserDTO | null;
    /**
     * 
     * @type {number}
     * @memberof FootprintDTO
     */
    visitedAt?: number | null;
}

/**
 * Check if a given object implements the FootprintDTO interface.
 */
export function instanceOfFootprintDTO(value: object): value is FootprintDTO {
    return true;
}

export function FootprintDTOFromJSON(json: any): FootprintDTO {
    return FootprintDTOFromJSONTyped(json, false);
}

export function FootprintDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): FootprintDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'user': json['user'] == null ? undefined : UserUserDTOFromJSON(json['user']),
        'visitedAt': json['visited_at'] == null ? undefined : json['visited_at'],
    };
}

export function FootprintDTOToJSON(json: any): FootprintDTO {
    return FootprintDTOToJSONTyped(json, false);
}

export function FootprintDTOToJSONTyped(value?: FootprintDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'user': UserUserDTOToJSON(value['user']),
        'visited_at': value['visitedAt'],
    };
}

