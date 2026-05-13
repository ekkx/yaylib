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
import type { UserCampaign } from './UserCampaign';
import {
    UserCampaignFromJSON,
    UserCampaignFromJSONTyped,
    UserCampaignToJSON,
    UserCampaignToJSONTyped,
} from './UserCampaign';

/**
 * 
 * @export
 * @interface CampaignDTO
 */
export interface CampaignDTO {
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    bannerImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    descriptionUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignDTO
     */
    endAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignDTO
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    multiplierDetailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CampaignDTO
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignDTO
     */
    ranking?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignDTO
     */
    startAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignDTO
     */
    totalPoints?: number | null;
    /**
     * 
     * @type {UserCampaign}
     * @memberof CampaignDTO
     */
    userCampaign?: UserCampaign | null;
}

/**
 * Check if a given object implements the CampaignDTO interface.
 */
export function instanceOfCampaignDTO(value: object): value is CampaignDTO {
    return true;
}

export function CampaignDTOFromJSON(json: any): CampaignDTO {
    return CampaignDTOFromJSONTyped(json, false);
}

export function CampaignDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'bannerImage': json['banner_image'] == null ? undefined : json['banner_image'],
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'description': json['description'] == null ? undefined : json['description'],
        'descriptionUrl': json['description_url'] == null ? undefined : json['description_url'],
        'endAt': json['end_at'] == null ? undefined : json['end_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'multiplierDetailUrl': json['multiplier_detail_url'] == null ? undefined : json['multiplier_detail_url'],
        'name': json['name'] == null ? undefined : json['name'],
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
        'startAt': json['start_at'] == null ? undefined : json['start_at'],
        'totalPoints': json['total_points'] == null ? undefined : json['total_points'],
        'userCampaign': json['user_campaign'] == null ? undefined : UserCampaignFromJSON(json['user_campaign']),
    };
}

export function CampaignDTOToJSON(json: any): CampaignDTO {
    return CampaignDTOToJSONTyped(json, false);
}

export function CampaignDTOToJSONTyped(value?: CampaignDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'banner_image': value['bannerImage'],
        'cover_image': value['coverImage'],
        'description': value['description'],
        'description_url': value['descriptionUrl'],
        'end_at': value['endAt'],
        'id': value['id'],
        'multiplier_detail_url': value['multiplierDetailUrl'],
        'name': value['name'],
        'ranking': value['ranking'],
        'start_at': value['startAt'],
        'total_points': value['totalPoints'],
        'user_campaign': UserCampaignToJSON(value['userCampaign']),
    };
}

