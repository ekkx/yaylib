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
 * @interface DetailsDTO
 */
export interface DetailsDTO {
    /**
     * 
     * @type {number}
     * @memberof DetailsDTO
     */
    rewardAmountAdult?: number | null;
    /**
     * 
     * @type {number}
     * @memberof DetailsDTO
     */
    rewardAmountMinor?: number | null;
}

/**
 * Check if a given object implements the DetailsDTO interface.
 */
export function instanceOfDetailsDTO(value: object): value is DetailsDTO {
    return true;
}

export function DetailsDTOFromJSON(json: any): DetailsDTO {
    return DetailsDTOFromJSONTyped(json, false);
}

export function DetailsDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): DetailsDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'rewardAmountAdult': json['reward_amount_adult'] == null ? undefined : json['reward_amount_adult'],
        'rewardAmountMinor': json['reward_amount_minor'] == null ? undefined : json['reward_amount_minor'],
    };
}

export function DetailsDTOToJSON(json: any): DetailsDTO {
    return DetailsDTOToJSONTyped(json, false);
}

export function DetailsDTOToJSONTyped(value?: DetailsDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'reward_amount_adult': value['rewardAmountAdult'],
        'reward_amount_minor': value['rewardAmountMinor'],
    };
}

