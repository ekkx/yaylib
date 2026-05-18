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
 * @interface ProgressDTO
 */
export interface ProgressDTO {
    /**
     * 
     * @type {number}
     * @memberof ProgressDTO
     */
    actionsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ProgressDTO
     */
    completedAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ProgressDTO
     */
    totalEarnedPoints?: number | null;
}

/**
 * Check if a given object implements the ProgressDTO interface.
 */
export function instanceOfProgressDTO(value: object): value is ProgressDTO {
    return true;
}

export function ProgressDTOFromJSON(json: any): ProgressDTO {
    return ProgressDTOFromJSONTyped(json, false);
}

export function ProgressDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProgressDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'actionsCount': json['actions_count'] == null ? undefined : json['actions_count'],
        'completedAt': json['completed_at'] == null ? undefined : json['completed_at'],
        'totalEarnedPoints': json['total_earned_points'] == null ? undefined : json['total_earned_points'],
    };
}

export function ProgressDTOToJSON(json: any): ProgressDTO {
    return ProgressDTOToJSONTyped(json, false);
}

export function ProgressDTOToJSONTyped(value?: ProgressDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'actions_count': value['actionsCount'],
        'completed_at': value['completedAt'],
        'total_earned_points': value['totalEarnedPoints'],
    };
}

