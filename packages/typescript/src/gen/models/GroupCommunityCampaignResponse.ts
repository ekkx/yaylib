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
/**
 * 
 * @export
 * @interface GroupCommunityCampaignResponse
 */
export interface GroupCommunityCampaignResponse {
    /**
     * 
     * @type {boolean}
     * @memberof GroupCommunityCampaignResponse
     */
    eligibleForApply?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GroupCommunityCampaignResponse
     */
    viewableGroupLeaderboard?: boolean | null;
}

/**
 * Check if a given object implements the GroupCommunityCampaignResponse interface.
 */
export function instanceOfGroupCommunityCampaignResponse(value: object): value is GroupCommunityCampaignResponse {
    return true;
}

export function GroupCommunityCampaignResponseFromJSON(json: any): GroupCommunityCampaignResponse {
    return GroupCommunityCampaignResponseFromJSONTyped(json, false);
}

export function GroupCommunityCampaignResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupCommunityCampaignResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'eligibleForApply': json['eligible_for_apply'] == null ? undefined : json['eligible_for_apply'],
        'viewableGroupLeaderboard': json['viewable_group_leaderboard'] == null ? undefined : json['viewable_group_leaderboard'],
    };
}

export function GroupCommunityCampaignResponseToJSON(json: any): GroupCommunityCampaignResponse {
    return GroupCommunityCampaignResponseToJSONTyped(json, false);
}

export function GroupCommunityCampaignResponseToJSONTyped(value?: GroupCommunityCampaignResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'eligible_for_apply': value['eligibleForApply'],
        'viewable_group_leaderboard': value['viewableGroupLeaderboard'],
    };
}

