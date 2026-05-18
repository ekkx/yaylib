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
import type { RecentSearch } from './RecentSearch.js';
import {
    RecentSearchFromJSON,
    RecentSearchFromJSONTyped,
    RecentSearchToJSON,
    RecentSearchToJSONTyped,
} from './RecentSearch.js';

/**
 * 
 * @export
 * @interface RecentSearchesResponse
 */
export interface RecentSearchesResponse {
    /**
     * 
     * @type {Array<RecentSearch>}
     * @memberof RecentSearchesResponse
     */
    recentSearches?: Array<RecentSearch> | null;
}

/**
 * Check if a given object implements the RecentSearchesResponse interface.
 */
export function instanceOfRecentSearchesResponse(value: object): value is RecentSearchesResponse {
    return true;
}

export function RecentSearchesResponseFromJSON(json: any): RecentSearchesResponse {
    return RecentSearchesResponseFromJSONTyped(json, false);
}

export function RecentSearchesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecentSearchesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'recentSearches': json['recent_searches'] == null ? undefined : ((json['recent_searches'] as Array<any>).map(RecentSearchFromJSON)),
    };
}

export function RecentSearchesResponseToJSON(json: any): RecentSearchesResponse {
    return RecentSearchesResponseToJSONTyped(json, false);
}

export function RecentSearchesResponseToJSONTyped(value?: RecentSearchesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'recent_searches': value['recentSearches'] == null ? undefined : ((value['recentSearches'] as Array<any>).map(RecentSearchToJSON)),
    };
}

