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
import type { TopCampaignUserDTO } from './TopCampaignUserDTO';
import {
    TopCampaignUserDTOFromJSON,
    TopCampaignUserDTOFromJSONTyped,
    TopCampaignUserDTOToJSON,
    TopCampaignUserDTOToJSONTyped,
} from './TopCampaignUserDTO';

/**
 * 
 * @export
 * @interface CampaignRankingResponse
 */
export interface CampaignRankingResponse {
    /**
     * 
     * @type {number}
     * @memberof CampaignRankingResponse
     */
    myRanking?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignRankingResponse
     */
    nextCursor?: string | null;
    /**
     * 
     * @type {Array<TopCampaignUserDTO>}
     * @memberof CampaignRankingResponse
     */
    topCampaignUsers?: Array<TopCampaignUserDTO> | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignRankingResponse
     */
    updatedAt?: number | null;
}

/**
 * Check if a given object implements the CampaignRankingResponse interface.
 */
export function instanceOfCampaignRankingResponse(value: object): value is CampaignRankingResponse {
    return true;
}

export function CampaignRankingResponseFromJSON(json: any): CampaignRankingResponse {
    return CampaignRankingResponseFromJSONTyped(json, false);
}

export function CampaignRankingResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignRankingResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'myRanking': json['my_ranking'] == null ? undefined : json['my_ranking'],
        'nextCursor': json['next_cursor'] == null ? undefined : json['next_cursor'],
        'topCampaignUsers': json['top_campaign_users'] == null ? undefined : ((json['top_campaign_users'] as Array<any>).map(TopCampaignUserDTOFromJSON)),
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
    };
}

export function CampaignRankingResponseToJSON(json: any): CampaignRankingResponse {
    return CampaignRankingResponseToJSONTyped(json, false);
}

export function CampaignRankingResponseToJSONTyped(value?: CampaignRankingResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'my_ranking': value['myRanking'],
        'next_cursor': value['nextCursor'],
        'top_campaign_users': value['topCampaignUsers'] == null ? undefined : ((value['topCampaignUsers'] as Array<any>).map(TopCampaignUserDTOToJSON)),
        'updated_at': value['updatedAt'],
    };
}

