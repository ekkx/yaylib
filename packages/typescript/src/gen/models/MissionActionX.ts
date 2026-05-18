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
 * @interface MissionActionX
 */
export interface MissionActionX {
    /**
     * 
     * @type {number}
     * @memberof MissionActionX
     */
    descriptionId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionActionX
     */
    redirectButtonId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionActionX
     */
    rewardDetailsId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionActionX
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the MissionActionX interface.
 */
export function instanceOfMissionActionX(value: object): value is MissionActionX {
    return true;
}

export function MissionActionXFromJSON(json: any): MissionActionX {
    return MissionActionXFromJSONTyped(json, false);
}

export function MissionActionXFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionActionX {
    if (json == null) {
        return json;
    }
    return {
        
        'descriptionId': json['description_id'] == null ? undefined : json['description_id'],
        'redirectButtonId': json['redirect_button_id'] == null ? undefined : json['redirect_button_id'],
        'rewardDetailsId': json['reward_details_id'] == null ? undefined : json['reward_details_id'],
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function MissionActionXToJSON(json: any): MissionActionX {
    return MissionActionXToJSONTyped(json, false);
}

export function MissionActionXToJSONTyped(value?: MissionActionX | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'description_id': value['descriptionId'],
        'redirect_button_id': value['redirectButtonId'],
        'reward_details_id': value['rewardDetailsId'],
        'title_id': value['titleId'],
    };
}

