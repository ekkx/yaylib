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
import type { Survey } from './Survey.js';
import {
    SurveyFromJSON,
    SurveyFromJSONTyped,
    SurveyToJSON,
    SurveyToJSONTyped,
} from './Survey.js';

/**
 * 
 * @export
 * @interface VoteSurveyResponse
 */
export interface VoteSurveyResponse {
    /**
     * 
     * @type {Survey}
     * @memberof VoteSurveyResponse
     */
    survey?: Survey | null;
}

/**
 * Check if a given object implements the VoteSurveyResponse interface.
 */
export function instanceOfVoteSurveyResponse(value: object): value is VoteSurveyResponse {
    return true;
}

export function VoteSurveyResponseFromJSON(json: any): VoteSurveyResponse {
    return VoteSurveyResponseFromJSONTyped(json, false);
}

export function VoteSurveyResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): VoteSurveyResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'survey': json['survey'] == null ? undefined : SurveyFromJSON(json['survey']),
    };
}

export function VoteSurveyResponseToJSON(json: any): VoteSurveyResponse {
    return VoteSurveyResponseToJSONTyped(json, false);
}

export function VoteSurveyResponseToJSONTyped(value?: VoteSurveyResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'survey': SurveyToJSON(value['survey']),
    };
}

