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
import type { ProgressDTO } from './ProgressDTO.js';
import {
    ProgressDTOFromJSON,
    ProgressDTOFromJSONTyped,
    ProgressDTOToJSON,
    ProgressDTOToJSONTyped,
} from './ProgressDTO.js';
import type { MissionType } from './MissionType.js';
import {
    MissionTypeFromJSON,
    MissionTypeFromJSONTyped,
    MissionTypeToJSON,
    MissionTypeToJSONTyped,
} from './MissionType.js';

/**
 * 
 * @export
 * @interface MissionDTO
 */
export interface MissionDTO {
    /**
     * 
     * @type {string}
     * @memberof MissionDTO
     */
    action?: string | null;
    /**
     * 
     * @type {number}
     * @memberof MissionDTO
     */
    id?: number | null;
    /**
     * 
     * @type {MissionType}
     * @memberof MissionDTO
     */
    missionType?: MissionType | null;
    /**
     * 
     * @type {ProgressDTO}
     * @memberof MissionDTO
     */
    progress?: ProgressDTO | null;
    /**
     * 
     * @type {number}
     * @memberof MissionDTO
     */
    requiredActionsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionDTO
     */
    rewardPoints?: number | null;
}



/**
 * Check if a given object implements the MissionDTO interface.
 */
export function instanceOfMissionDTO(value: object): value is MissionDTO {
    return true;
}

export function MissionDTOFromJSON(json: any): MissionDTO {
    return MissionDTOFromJSONTyped(json, false);
}

export function MissionDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'] == null ? undefined : json['action'],
        'id': json['id'] == null ? undefined : json['id'],
        'missionType': json['mission_type'] == null ? undefined : MissionTypeFromJSON(json['mission_type']),
        'progress': json['progress'] == null ? undefined : ProgressDTOFromJSON(json['progress']),
        'requiredActionsCount': json['required_actions_count'] == null ? undefined : json['required_actions_count'],
        'rewardPoints': json['reward_points'] == null ? undefined : json['reward_points'],
    };
}

export function MissionDTOToJSON(json: any): MissionDTO {
    return MissionDTOToJSONTyped(json, false);
}

export function MissionDTOToJSONTyped(value?: MissionDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'action': value['action'],
        'id': value['id'],
        'mission_type': MissionTypeToJSON(value['missionType']),
        'progress': ProgressDTOToJSON(value['progress']),
        'required_actions_count': value['requiredActionsCount'],
        'reward_points': value['rewardPoints'],
    };
}

