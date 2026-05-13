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
import type { MissionDTO } from './MissionDTO';
import {
    MissionDTOFromJSON,
    MissionDTOFromJSONTyped,
    MissionDTOToJSON,
    MissionDTOToJSONTyped,
} from './MissionDTO';

/**
 * 
 * @export
 * @interface CampaignMissionsResponse
 */
export interface CampaignMissionsResponse {
    /**
     * 
     * @type {Array<MissionDTO>}
     * @memberof CampaignMissionsResponse
     */
    dailyMissions?: Array<MissionDTO> | null;
    /**
     * 
     * @type {Array<MissionDTO>}
     * @memberof CampaignMissionsResponse
     */
    multiplierMissions?: Array<MissionDTO> | null;
    /**
     * 
     * @type {Array<MissionDTO>}
     * @memberof CampaignMissionsResponse
     */
    normalMissions?: Array<MissionDTO> | null;
    /**
     * 
     * @type {Array<MissionDTO>}
     * @memberof CampaignMissionsResponse
     */
    unlimitedMissions?: Array<MissionDTO> | null;
}

/**
 * Check if a given object implements the CampaignMissionsResponse interface.
 */
export function instanceOfCampaignMissionsResponse(value: object): value is CampaignMissionsResponse {
    return true;
}

export function CampaignMissionsResponseFromJSON(json: any): CampaignMissionsResponse {
    return CampaignMissionsResponseFromJSONTyped(json, false);
}

export function CampaignMissionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignMissionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'dailyMissions': json['daily_missions'] == null ? undefined : ((json['daily_missions'] as Array<any>).map(MissionDTOFromJSON)),
        'multiplierMissions': json['multiplier_missions'] == null ? undefined : ((json['multiplier_missions'] as Array<any>).map(MissionDTOFromJSON)),
        'normalMissions': json['normal_missions'] == null ? undefined : ((json['normal_missions'] as Array<any>).map(MissionDTOFromJSON)),
        'unlimitedMissions': json['unlimited_missions'] == null ? undefined : ((json['unlimited_missions'] as Array<any>).map(MissionDTOFromJSON)),
    };
}

export function CampaignMissionsResponseToJSON(json: any): CampaignMissionsResponse {
    return CampaignMissionsResponseToJSONTyped(json, false);
}

export function CampaignMissionsResponseToJSONTyped(value?: CampaignMissionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'daily_missions': value['dailyMissions'] == null ? undefined : ((value['dailyMissions'] as Array<any>).map(MissionDTOToJSON)),
        'multiplier_missions': value['multiplierMissions'] == null ? undefined : ((value['multiplierMissions'] as Array<any>).map(MissionDTOToJSON)),
        'normal_missions': value['normalMissions'] == null ? undefined : ((value['normalMissions'] as Array<any>).map(MissionDTOToJSON)),
        'unlimited_missions': value['unlimitedMissions'] == null ? undefined : ((value['unlimitedMissions'] as Array<any>).map(MissionDTOToJSON)),
    };
}

