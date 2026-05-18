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
 * @interface ReviewRestriction
 */
export interface ReviewRestriction {
    /**
     * 
     * @type {string}
     * @memberof ReviewRestriction
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the ReviewRestriction interface.
 */
export function instanceOfReviewRestriction(value: object): value is ReviewRestriction {
    return true;
}

export function ReviewRestrictionFromJSON(json: any): ReviewRestriction {
    return ReviewRestrictionFromJSONTyped(json, false);
}

export function ReviewRestrictionFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReviewRestriction {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function ReviewRestrictionToJSON(json: any): ReviewRestriction {
    return ReviewRestrictionToJSONTyped(json, false);
}

export function ReviewRestrictionToJSONTyped(value?: ReviewRestriction | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

