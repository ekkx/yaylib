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
 * @interface AdditionGasPercentDTO
 */
export interface AdditionGasPercentDTO {
    /**
     * 
     * @type {number}
     * @memberof AdditionGasPercentDTO
     */
    fast?: number | null;
    /**
     * 
     * @type {number}
     * @memberof AdditionGasPercentDTO
     */
    normal?: number | null;
    /**
     * 
     * @type {number}
     * @memberof AdditionGasPercentDTO
     */
    slow?: number | null;
}

/**
 * Check if a given object implements the AdditionGasPercentDTO interface.
 */
export function instanceOfAdditionGasPercentDTO(value: object): value is AdditionGasPercentDTO {
    return true;
}

export function AdditionGasPercentDTOFromJSON(json: any): AdditionGasPercentDTO {
    return AdditionGasPercentDTOFromJSONTyped(json, false);
}

export function AdditionGasPercentDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): AdditionGasPercentDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'fast': json['fast'] == null ? undefined : json['fast'],
        'normal': json['normal'] == null ? undefined : json['normal'],
        'slow': json['slow'] == null ? undefined : json['slow'],
    };
}

export function AdditionGasPercentDTOToJSON(json: any): AdditionGasPercentDTO {
    return AdditionGasPercentDTOToJSONTyped(json, false);
}

export function AdditionGasPercentDTOToJSONTyped(value?: AdditionGasPercentDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'fast': value['fast'],
        'normal': value['normal'],
        'slow': value['slow'],
    };
}

