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
import type { CampaignVipEmplBonusDTO } from './CampaignVipEmplBonusDTO';
import {
    CampaignVipEmplBonusDTOFromJSON,
    CampaignVipEmplBonusDTOFromJSONTyped,
    CampaignVipEmplBonusDTOToJSON,
    CampaignVipEmplBonusDTOToJSONTyped,
} from './CampaignVipEmplBonusDTO';

/**
 * 
 * @export
 * @interface CampaignVipEmplBonusResponse
 */
export interface CampaignVipEmplBonusResponse {
    /**
     * 
     * @type {CampaignVipEmplBonusDTO}
     * @memberof CampaignVipEmplBonusResponse
     */
    appCampaign?: CampaignVipEmplBonusDTO | null;
}

/**
 * Check if a given object implements the CampaignVipEmplBonusResponse interface.
 */
export function instanceOfCampaignVipEmplBonusResponse(value: object): value is CampaignVipEmplBonusResponse {
    return true;
}

export function CampaignVipEmplBonusResponseFromJSON(json: any): CampaignVipEmplBonusResponse {
    return CampaignVipEmplBonusResponseFromJSONTyped(json, false);
}

export function CampaignVipEmplBonusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignVipEmplBonusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'appCampaign': json['app_campaign'] == null ? undefined : CampaignVipEmplBonusDTOFromJSON(json['app_campaign']),
    };
}

export function CampaignVipEmplBonusResponseToJSON(json: any): CampaignVipEmplBonusResponse {
    return CampaignVipEmplBonusResponseToJSONTyped(json, false);
}

export function CampaignVipEmplBonusResponseToJSONTyped(value?: CampaignVipEmplBonusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'app_campaign': CampaignVipEmplBonusDTOToJSON(value['appCampaign']),
    };
}

