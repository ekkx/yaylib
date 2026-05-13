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
 * @interface SearchCriteria
 */
export interface SearchCriteria {
    /**
     * 
     * @type {string}
     * @memberof SearchCriteria
     */
    biography?: string | null;
    /**
     * 
     * @type {number}
     * @memberof SearchCriteria
     */
    gender?: number | null;
    /**
     * 
     * @type {string}
     * @memberof SearchCriteria
     */
    nickname?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SearchCriteria
     */
    prefecture?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SearchCriteria
     */
    username?: string | null;
}

/**
 * Check if a given object implements the SearchCriteria interface.
 */
export function instanceOfSearchCriteria(value: object): value is SearchCriteria {
    return true;
}

export function SearchCriteriaFromJSON(json: any): SearchCriteria {
    return SearchCriteriaFromJSONTyped(json, false);
}

export function SearchCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchCriteria {
    if (json == null) {
        return json;
    }
    return {
        
        'biography': json['biography'] == null ? undefined : json['biography'],
        'gender': json['gender'] == null ? undefined : json['gender'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'prefecture': json['prefecture'] == null ? undefined : json['prefecture'],
        'username': json['username'] == null ? undefined : json['username'],
    };
}

export function SearchCriteriaToJSON(json: any): SearchCriteria {
    return SearchCriteriaToJSONTyped(json, false);
}

export function SearchCriteriaToJSONTyped(value?: SearchCriteria | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'biography': value['biography'],
        'gender': value['gender'],
        'nickname': value['nickname'],
        'prefecture': value['prefecture'],
        'username': value['username'],
    };
}

