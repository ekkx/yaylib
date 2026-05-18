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
import type { CampaignPointHistoryDTO } from './CampaignPointHistoryDTO.js';
import {
    CampaignPointHistoryDTOFromJSON,
    CampaignPointHistoryDTOFromJSONTyped,
    CampaignPointHistoryDTOToJSON,
    CampaignPointHistoryDTOToJSONTyped,
} from './CampaignPointHistoryDTO.js';

/**
 * 
 * @export
 * @interface CampaignPointHistoryResponse
 */
export interface CampaignPointHistoryResponse {
    /**
     * 
     * @type {Array<CampaignPointHistoryDTO>}
     * @memberof CampaignPointHistoryResponse
     */
    pointHistories?: Array<CampaignPointHistoryDTO> | null;
}

/**
 * Check if a given object implements the CampaignPointHistoryResponse interface.
 */
export function instanceOfCampaignPointHistoryResponse(value: object): value is CampaignPointHistoryResponse {
    return true;
}

export function CampaignPointHistoryResponseFromJSON(json: any): CampaignPointHistoryResponse {
    return CampaignPointHistoryResponseFromJSONTyped(json, false);
}

export function CampaignPointHistoryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignPointHistoryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'pointHistories': json['point_histories'] == null ? undefined : ((json['point_histories'] as Array<any>).map(CampaignPointHistoryDTOFromJSON)),
    };
}

export function CampaignPointHistoryResponseToJSON(json: any): CampaignPointHistoryResponse {
    return CampaignPointHistoryResponseToJSONTyped(json, false);
}

export function CampaignPointHistoryResponseToJSONTyped(value?: CampaignPointHistoryResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'point_histories': value['pointHistories'] == null ? undefined : ((value['pointHistories'] as Array<any>).map(CampaignPointHistoryDTOToJSON)),
    };
}

