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
import type { MultiplierBreakdown } from './MultiplierBreakdown.js';
import {
    MultiplierBreakdownFromJSON,
    MultiplierBreakdownFromJSONTyped,
    MultiplierBreakdownToJSON,
    MultiplierBreakdownToJSONTyped,
} from './MultiplierBreakdown.js';

/**
 * 
 * @export
 * @interface UserCampaign
 */
export interface UserCampaign {
    /**
     * 
     * @type {number}
     * @memberof UserCampaign
     */
    multiplier?: number | null;
    /**
     * 
     * @type {MultiplierBreakdown}
     * @memberof UserCampaign
     */
    multiplierBreakdown?: MultiplierBreakdown | null;
    /**
     * 
     * @type {number}
     * @memberof UserCampaign
     */
    points?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCampaign
     */
    ranking?: number | null;
}

/**
 * Check if a given object implements the UserCampaign interface.
 */
export function instanceOfUserCampaign(value: object): value is UserCampaign {
    return true;
}

export function UserCampaignFromJSON(json: any): UserCampaign {
    return UserCampaignFromJSONTyped(json, false);
}

export function UserCampaignFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserCampaign {
    if (json == null) {
        return json;
    }
    return {
        
        'multiplier': json['multiplier'] == null ? undefined : json['multiplier'],
        'multiplierBreakdown': json['multiplier_breakdown'] == null ? undefined : MultiplierBreakdownFromJSON(json['multiplier_breakdown']),
        'points': json['points'] == null ? undefined : json['points'],
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
    };
}

export function UserCampaignToJSON(json: any): UserCampaign {
    return UserCampaignToJSONTyped(json, false);
}

export function UserCampaignToJSONTyped(value?: UserCampaign | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'multiplier': value['multiplier'],
        'multiplier_breakdown': MultiplierBreakdownToJSON(value['multiplierBreakdown']),
        'points': value['points'],
        'ranking': value['ranking'],
    };
}

