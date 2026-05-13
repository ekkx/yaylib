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
import type { SurveyChoice } from './SurveyChoice';
import {
    SurveyChoiceFromJSON,
    SurveyChoiceFromJSONTyped,
    SurveyChoiceToJSON,
    SurveyChoiceToJSONTyped,
} from './SurveyChoice';

/**
 * 
 * @export
 * @interface ModelSurvey
 */
export interface ModelSurvey {
    /**
     * 
     * @type {Array<SurveyChoice>}
     * @memberof ModelSurvey
     */
    choices?: Array<SurveyChoice> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelSurvey
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelSurvey
     */
    isVoted?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof ModelSurvey
     */
    votesCount?: number | null;
}

/**
 * Check if a given object implements the ModelSurvey interface.
 */
export function instanceOfModelSurvey(value: object): value is ModelSurvey {
    return true;
}

export function ModelSurveyFromJSON(json: any): ModelSurvey {
    return ModelSurveyFromJSONTyped(json, false);
}

export function ModelSurveyFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSurvey {
    if (json == null) {
        return json;
    }
    return {
        
        'choices': json['choices'] == null ? undefined : ((json['choices'] as Array<any>).map(SurveyChoiceFromJSON)),
        'id': json['id'] == null ? undefined : json['id'],
        'isVoted': json['is_voted'] == null ? undefined : json['is_voted'],
        'votesCount': json['votes_count'] == null ? undefined : json['votes_count'],
    };
}

export function ModelSurveyToJSON(json: any): ModelSurvey {
    return ModelSurveyToJSONTyped(json, false);
}

export function ModelSurveyToJSONTyped(value?: ModelSurvey | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'choices': value['choices'] == null ? undefined : ((value['choices'] as Array<any>).map(SurveyChoiceToJSON)),
        'id': value['id'],
        'is_voted': value['isVoted'],
        'votes_count': value['votesCount'],
    };
}

