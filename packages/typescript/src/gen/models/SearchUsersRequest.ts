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
import type { SearchCriteria } from './SearchCriteria.js';
import {
    SearchCriteriaFromJSON,
    SearchCriteriaFromJSONTyped,
    SearchCriteriaToJSON,
    SearchCriteriaToJSONTyped,
} from './SearchCriteria.js';

/**
 * 
 * @export
 * @interface SearchUsersRequest
 */
export interface SearchUsersRequest {
    /**
     * 
     * @type {SearchCriteria}
     * @memberof SearchUsersRequest
     */
    user?: SearchCriteria | null;
}

/**
 * Check if a given object implements the SearchUsersRequest interface.
 */
export function instanceOfSearchUsersRequest(value: object): value is SearchUsersRequest {
    return true;
}

export function SearchUsersRequestFromJSON(json: any): SearchUsersRequest {
    return SearchUsersRequestFromJSONTyped(json, false);
}

export function SearchUsersRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchUsersRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'user': json['user'] == null ? undefined : SearchCriteriaFromJSON(json['user']),
    };
}

export function SearchUsersRequestToJSON(json: any): SearchUsersRequest {
    return SearchUsersRequestToJSONTyped(json, false);
}

export function SearchUsersRequestToJSONTyped(value?: SearchUsersRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'user': SearchCriteriaToJSON(value['user']),
    };
}

