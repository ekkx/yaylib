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
import type { CampaignDTO } from './CampaignDTO';
import {
    CampaignDTOFromJSON,
    CampaignDTOFromJSONTyped,
    CampaignDTOToJSON,
    CampaignDTOToJSONTyped,
} from './CampaignDTO';

/**
 * 
 * @export
 * @interface CampaignResponse
 */
export interface CampaignResponse {
    /**
     * 
     * @type {CampaignDTO}
     * @memberof CampaignResponse
     */
    campaign?: CampaignDTO | null;
}

/**
 * Check if a given object implements the CampaignResponse interface.
 */
export function instanceOfCampaignResponse(value: object): value is CampaignResponse {
    return true;
}

export function CampaignResponseFromJSON(json: any): CampaignResponse {
    return CampaignResponseFromJSONTyped(json, false);
}

export function CampaignResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'campaign': json['campaign'] == null ? undefined : CampaignDTOFromJSON(json['campaign']),
    };
}

export function CampaignResponseToJSON(json: any): CampaignResponse {
    return CampaignResponseToJSONTyped(json, false);
}

export function CampaignResponseToJSONTyped(value?: CampaignResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'campaign': CampaignDTOToJSON(value['campaign']),
    };
}

