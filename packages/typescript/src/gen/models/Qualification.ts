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
 * @interface Qualification
 */
export interface Qualification {
    /**
     * 
     * @type {boolean}
     * @memberof Qualification
     */
    qualified?: boolean | null;
    /**
     * 
     * @type {object}
     * @memberof Qualification
     */
    type?: object | null;
}

/**
 * Check if a given object implements the Qualification interface.
 */
export function instanceOfQualification(value: object): value is Qualification {
    return true;
}

export function QualificationFromJSON(json: any): Qualification {
    return QualificationFromJSONTyped(json, false);
}

export function QualificationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Qualification {
    if (json == null) {
        return json;
    }
    return {
        
        'qualified': json['qualified'] == null ? undefined : json['qualified'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function QualificationToJSON(json: any): Qualification {
    return QualificationToJSONTyped(json, false);
}

export function QualificationToJSONTyped(value?: Qualification | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'qualified': value['qualified'],
        'type': value['type'],
    };
}

