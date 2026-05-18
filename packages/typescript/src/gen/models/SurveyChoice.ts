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
 * @interface SurveyChoice
 */
export interface SurveyChoice {
    /**
     * 
     * @type {number}
     * @memberof SurveyChoice
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof SurveyChoice
     */
    label?: string | null;
    /**
     * 
     * @type {number}
     * @memberof SurveyChoice
     */
    votesCount?: number | null;
}

/**
 * Check if a given object implements the SurveyChoice interface.
 */
export function instanceOfSurveyChoice(value: object): value is SurveyChoice {
    return true;
}

export function SurveyChoiceFromJSON(json: any): SurveyChoice {
    return SurveyChoiceFromJSONTyped(json, false);
}

export function SurveyChoiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): SurveyChoice {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'label': json['label'] == null ? undefined : json['label'],
        'votesCount': json['votes_count'] == null ? undefined : json['votes_count'],
    };
}

export function SurveyChoiceToJSON(json: any): SurveyChoice {
    return SurveyChoiceToJSONTyped(json, false);
}

export function SurveyChoiceToJSONTyped(value?: SurveyChoice | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'label': value['label'],
        'votes_count': value['votesCount'],
    };
}

