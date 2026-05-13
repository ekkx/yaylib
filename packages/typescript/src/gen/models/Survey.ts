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
import type { Choice } from './Choice';
import {
    ChoiceFromJSON,
    ChoiceFromJSONTyped,
    ChoiceToJSON,
    ChoiceToJSONTyped,
} from './Choice';

/**
 * 
 * @export
 * @interface Survey
 */
export interface Survey {
    /**
     * 
     * @type {Array<Choice>}
     * @memberof Survey
     */
    choices?: Array<Choice> | null;
    /**
     * 
     * @type {number}
     * @memberof Survey
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Survey
     */
    voted?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Survey
     */
    votesCount?: number | null;
}

/**
 * Check if a given object implements the Survey interface.
 */
export function instanceOfSurvey(value: object): value is Survey {
    return true;
}

export function SurveyFromJSON(json: any): Survey {
    return SurveyFromJSONTyped(json, false);
}

export function SurveyFromJSONTyped(json: any, ignoreDiscriminator: boolean): Survey {
    if (json == null) {
        return json;
    }
    return {
        
        'choices': json['choices'] == null ? undefined : ((json['choices'] as Array<any>).map(ChoiceFromJSON)),
        'id': json['id'] == null ? undefined : json['id'],
        'voted': json['voted'] == null ? undefined : json['voted'],
        'votesCount': json['votes_count'] == null ? undefined : json['votes_count'],
    };
}

export function SurveyToJSON(json: any): Survey {
    return SurveyToJSONTyped(json, false);
}

export function SurveyToJSONTyped(value?: Survey | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'choices': value['choices'] == null ? undefined : ((value['choices'] as Array<any>).map(ChoiceToJSON)),
        'id': value['id'],
        'voted': value['voted'],
        'votes_count': value['votesCount'],
    };
}

