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
import type { Breakdown } from './Breakdown';
import {
    BreakdownFromJSON,
    BreakdownFromJSONTyped,
    BreakdownToJSON,
    BreakdownToJSONTyped,
} from './Breakdown';
import type { Mission } from './Mission';
import {
    MissionFromJSON,
    MissionFromJSONTyped,
    MissionToJSON,
    MissionToJSONTyped,
} from './Mission';

/**
 * 
 * @export
 * @interface CampaignPointHistoryDTO
 */
export interface CampaignPointHistoryDTO {
    /**
     * 
     * @type {Breakdown}
     * @memberof CampaignPointHistoryDTO
     */
    breakdown?: Breakdown | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignPointHistoryDTO
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignPointHistoryDTO
     */
    id?: number | null;
    /**
     * 
     * @type {Mission}
     * @memberof CampaignPointHistoryDTO
     */
    mission?: Mission | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignPointHistoryDTO
     */
    multiplier?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignPointHistoryDTO
     */
    pointType?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignPointHistoryDTO
     */
    totalPoints?: number | null;
}

/**
 * Check if a given object implements the CampaignPointHistoryDTO interface.
 */
export function instanceOfCampaignPointHistoryDTO(value: object): value is CampaignPointHistoryDTO {
    return true;
}

export function CampaignPointHistoryDTOFromJSON(json: any): CampaignPointHistoryDTO {
    return CampaignPointHistoryDTOFromJSONTyped(json, false);
}

export function CampaignPointHistoryDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignPointHistoryDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'breakdown': json['breakdown'] == null ? undefined : BreakdownFromJSON(json['breakdown']),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'mission': json['mission'] == null ? undefined : MissionFromJSON(json['mission']),
        'multiplier': json['multiplier'] == null ? undefined : json['multiplier'],
        'pointType': json['point_type'] == null ? undefined : json['point_type'],
        'totalPoints': json['total_points'] == null ? undefined : json['total_points'],
    };
}

export function CampaignPointHistoryDTOToJSON(json: any): CampaignPointHistoryDTO {
    return CampaignPointHistoryDTOToJSONTyped(json, false);
}

export function CampaignPointHistoryDTOToJSONTyped(value?: CampaignPointHistoryDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'breakdown': BreakdownToJSON(value['breakdown']),
        'created_at': value['createdAt'],
        'id': value['id'],
        'mission': MissionToJSON(value['mission']),
        'multiplier': value['multiplier'],
        'point_type': value['pointType'],
        'total_points': value['totalPoints'],
    };
}

