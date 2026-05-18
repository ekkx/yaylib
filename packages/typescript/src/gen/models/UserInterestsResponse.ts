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
import type { Interest } from './Interest.js';
import {
    InterestFromJSON,
    InterestFromJSONTyped,
    InterestToJSON,
    InterestToJSONTyped,
} from './Interest.js';

/**
 * 
 * @export
 * @interface UserInterestsResponse
 */
export interface UserInterestsResponse {
    /**
     * 
     * @type {Array<Interest>}
     * @memberof UserInterestsResponse
     */
    interests?: Array<Interest> | null;
}

/**
 * Check if a given object implements the UserInterestsResponse interface.
 */
export function instanceOfUserInterestsResponse(value: object): value is UserInterestsResponse {
    return true;
}

export function UserInterestsResponseFromJSON(json: any): UserInterestsResponse {
    return UserInterestsResponseFromJSONTyped(json, false);
}

export function UserInterestsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserInterestsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'interests': json['interests'] == null ? undefined : ((json['interests'] as Array<any>).map(InterestFromJSON)),
    };
}

export function UserInterestsResponseToJSON(json: any): UserInterestsResponse {
    return UserInterestsResponseToJSONTyped(json, false);
}

export function UserInterestsResponseToJSONTyped(value?: UserInterestsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'interests': value['interests'] == null ? undefined : ((value['interests'] as Array<any>).map(InterestToJSON)),
    };
}

