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
import type { DetailsDTO } from './DetailsDTO';
import {
    DetailsDTOFromJSON,
    DetailsDTOFromJSONTyped,
    DetailsDTOToJSON,
    DetailsDTOToJSONTyped,
} from './DetailsDTO';

/**
 * 
 * @export
 * @interface CampaignVipEmplBonusDTO
 */
export interface CampaignVipEmplBonusDTO {
    /**
     * 
     * @type {boolean}
     * @memberof CampaignVipEmplBonusDTO
     */
    claimed?: boolean | null;
    /**
     * 
     * @type {DetailsDTO}
     * @memberof CampaignVipEmplBonusDTO
     */
    details?: DetailsDTO | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignVipEmplBonusDTO
     */
    endAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignVipEmplBonusDTO
     */
    key?: string | null;
}

/**
 * Check if a given object implements the CampaignVipEmplBonusDTO interface.
 */
export function instanceOfCampaignVipEmplBonusDTO(value: object): value is CampaignVipEmplBonusDTO {
    return true;
}

export function CampaignVipEmplBonusDTOFromJSON(json: any): CampaignVipEmplBonusDTO {
    return CampaignVipEmplBonusDTOFromJSONTyped(json, false);
}

export function CampaignVipEmplBonusDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignVipEmplBonusDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'claimed': json['claimed'] == null ? undefined : json['claimed'],
        'details': json['details'] == null ? undefined : DetailsDTOFromJSON(json['details']),
        'endAt': json['end_at'] == null ? undefined : json['end_at'],
        'key': json['key'] == null ? undefined : json['key'],
    };
}

export function CampaignVipEmplBonusDTOToJSON(json: any): CampaignVipEmplBonusDTO {
    return CampaignVipEmplBonusDTOToJSONTyped(json, false);
}

export function CampaignVipEmplBonusDTOToJSONTyped(value?: CampaignVipEmplBonusDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'claimed': value['claimed'],
        'details': DetailsDTOToJSON(value['details']),
        'end_at': value['endAt'],
        'key': value['key'],
    };
}

