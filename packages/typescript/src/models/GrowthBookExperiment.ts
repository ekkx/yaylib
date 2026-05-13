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
 * @interface GrowthBookExperiment
 */
export interface GrowthBookExperiment {
    /**
     * 
     * @type {number}
     * @memberof GrowthBookExperiment
     */
    experimentNumber?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GrowthBookExperiment
     */
    number?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GrowthBookExperiment
     */
    variantNumber?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof GrowthBookExperiment
     */
    visible?: boolean | null;
}

/**
 * Check if a given object implements the GrowthBookExperiment interface.
 */
export function instanceOfGrowthBookExperiment(value: object): value is GrowthBookExperiment {
    return true;
}

export function GrowthBookExperimentFromJSON(json: any): GrowthBookExperiment {
    return GrowthBookExperimentFromJSONTyped(json, false);
}

export function GrowthBookExperimentFromJSONTyped(json: any, ignoreDiscriminator: boolean): GrowthBookExperiment {
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

export function GrowthBookExperimentToJSON(json: any): GrowthBookExperiment {
    return GrowthBookExperimentToJSONTyped(json, false);
}

export function GrowthBookExperimentToJSONTyped(value?: GrowthBookExperiment | null, ignoreDiscriminator: boolean = false): any {
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

