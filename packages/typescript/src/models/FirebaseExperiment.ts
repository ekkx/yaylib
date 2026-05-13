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
 * @interface FirebaseExperiment
 */
export interface FirebaseExperiment {
    /**
     * 
     * @type {number}
     * @memberof FirebaseExperiment
     */
    experimentNumber?: number | null;
    /**
     * 
     * @type {number}
     * @memberof FirebaseExperiment
     */
    number?: number | null;
    /**
     * 
     * @type {number}
     * @memberof FirebaseExperiment
     */
    variantNumber?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof FirebaseExperiment
     */
    visible?: boolean | null;
}

/**
 * Check if a given object implements the FirebaseExperiment interface.
 */
export function instanceOfFirebaseExperiment(value: object): value is FirebaseExperiment {
    return true;
}

export function FirebaseExperimentFromJSON(json: any): FirebaseExperiment {
    return FirebaseExperimentFromJSONTyped(json, false);
}

export function FirebaseExperimentFromJSONTyped(json: any, ignoreDiscriminator: boolean): FirebaseExperiment {
    if (json == null) {
        return json;
    }
    return {
        
        'experimentNumber': json['experiment_number'] == null ? undefined : json['experiment_number'],
        'number': json['number'] == null ? undefined : json['number'],
        'variantNumber': json['variant_number'] == null ? undefined : json['variant_number'],
        'visible': json['visible'] == null ? undefined : json['visible'],
    };
}

export function FirebaseExperimentToJSON(json: any): FirebaseExperiment {
    return FirebaseExperimentToJSONTyped(json, false);
}

export function FirebaseExperimentToJSONTyped(value?: FirebaseExperiment | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'experiment_number': value['experimentNumber'],
        'number': value['number'],
        'variant_number': value['variantNumber'],
        'visible': value['visible'],
    };
}

