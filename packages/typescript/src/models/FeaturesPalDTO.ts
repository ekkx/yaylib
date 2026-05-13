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
 * @interface FeaturesPalDTO
 */
export interface FeaturesPalDTO {
    /**
     * 
     * @type {number}
     * @memberof FeaturesPalDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof FeaturesPalDTO
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof FeaturesPalDTO
     */
    depositwithdrawAddress?: string | null;
}

/**
 * Check if a given object implements the FeaturesPalDTO interface.
 */
export function instanceOfFeaturesPalDTO(value: object): value is FeaturesPalDTO {
    return true;
}

export function FeaturesPalDTOFromJSON(json: any): FeaturesPalDTO {
    return FeaturesPalDTOFromJSONTyped(json, false);
}

export function FeaturesPalDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): FeaturesPalDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'depositwithdrawAddress': json['depositwithdraw_address'] == null ? undefined : json['depositwithdraw_address'],
    };
}

export function FeaturesPalDTOToJSON(json: any): FeaturesPalDTO {
    return FeaturesPalDTOToJSONTyped(json, false);
}

export function FeaturesPalDTOToJSONTyped(value?: FeaturesPalDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'contract_address': value['contractAddress'],
        'depositwithdraw_address': value['depositwithdrawAddress'],
    };
}

