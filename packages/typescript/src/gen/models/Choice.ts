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
/**
 * 
 * @export
 * @interface Choice
 */
export interface Choice {
    /**
     * 
     * @type {number}
     * @memberof Choice
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Choice
     */
    label?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Choice
     */
    votesCount?: number | null;
}

/**
 * Check if a given object implements the Choice interface.
 */
export function instanceOfChoice(value: object): value is Choice {
    return true;
}

export function ChoiceFromJSON(json: any): Choice {
    return ChoiceFromJSONTyped(json, false);
}

export function ChoiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): Choice {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'label': json['label'] == null ? undefined : json['label'],
        'votesCount': json['votes_count'] == null ? undefined : json['votes_count'],
    };
}

export function ChoiceToJSON(json: any): Choice {
    return ChoiceToJSONTyped(json, false);
}

export function ChoiceToJSONTyped(value?: Choice | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'label': value['label'],
        'votes_count': value['votesCount'],
    };
}

