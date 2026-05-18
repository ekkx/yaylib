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
import type { CampaignDTO } from './CampaignDTO.js';
import {
    CampaignDTOFromJSON,
    CampaignDTOFromJSONTyped,
    CampaignDTOToJSON,
    CampaignDTOToJSONTyped,
} from './CampaignDTO.js';

/**
 * 
 * @export
 * @interface CampaignsResponse
 */
export interface CampaignsResponse {
    /**
     * 
     * @type {Array<CampaignDTO>}
     * @memberof CampaignsResponse
     */
    campaigns?: Array<CampaignDTO> | null;
}

/**
 * Check if a given object implements the CampaignsResponse interface.
 */
export function instanceOfCampaignsResponse(value: object): value is CampaignsResponse {
    return true;
}

export function CampaignsResponseFromJSON(json: any): CampaignsResponse {
    return CampaignsResponseFromJSONTyped(json, false);
}

export function CampaignsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'campaigns': json['campaigns'] == null ? undefined : ((json['campaigns'] as Array<any>).map(CampaignDTOFromJSON)),
    };
}

export function CampaignsResponseToJSON(json: any): CampaignsResponse {
    return CampaignsResponseToJSONTyped(json, false);
}

export function CampaignsResponseToJSONTyped(value?: CampaignsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'campaigns': value['campaigns'] == null ? undefined : ((value['campaigns'] as Array<any>).map(CampaignDTOToJSON)),
    };
}

