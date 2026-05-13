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
/**
 * 
 * @export
 * @interface CreateGroupResponse
 */
export interface CreateGroupResponse {
    /**
     * 
     * @type {number}
     * @memberof CreateGroupResponse
     */
    groupId?: number | null;
}

/**
 * Check if a given object implements the CreateGroupResponse interface.
 */
export function instanceOfCreateGroupResponse(value: object): value is CreateGroupResponse {
    return true;
}

export function CreateGroupResponseFromJSON(json: any): CreateGroupResponse {
    return CreateGroupResponseFromJSONTyped(json, false);
}

export function CreateGroupResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateGroupResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
    };
}

export function CreateGroupResponseToJSON(json: any): CreateGroupResponse {
    return CreateGroupResponseToJSONTyped(json, false);
}

export function CreateGroupResponseToJSONTyped(value?: CreateGroupResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_id': value['groupId'],
    };
}

