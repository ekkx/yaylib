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
 * @interface AppReviewStatusResponse
 */
export interface AppReviewStatusResponse {
    /**
     * 
     * @type {boolean}
     * @memberof AppReviewStatusResponse
     */
    isAppReviewed?: boolean | null;
}

/**
 * Check if a given object implements the AppReviewStatusResponse interface.
 */
export function instanceOfAppReviewStatusResponse(value: object): value is AppReviewStatusResponse {
    return true;
}

export function AppReviewStatusResponseFromJSON(json: any): AppReviewStatusResponse {
    return AppReviewStatusResponseFromJSONTyped(json, false);
}

export function AppReviewStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AppReviewStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isAppReviewed': json['is_app_reviewed'] == null ? undefined : json['is_app_reviewed'],
    };
}

export function AppReviewStatusResponseToJSON(json: any): AppReviewStatusResponse {
    return AppReviewStatusResponseToJSONTyped(json, false);
}

export function AppReviewStatusResponseToJSONTyped(value?: AppReviewStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_app_reviewed': value['isAppReviewed'],
    };
}

