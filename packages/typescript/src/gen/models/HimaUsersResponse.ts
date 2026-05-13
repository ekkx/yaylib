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
import type { UserWrapper } from './UserWrapper';
import {
    UserWrapperFromJSON,
    UserWrapperFromJSONTyped,
    UserWrapperToJSON,
    UserWrapperToJSONTyped,
} from './UserWrapper';

/**
 * 
 * @export
 * @interface HimaUsersResponse
 */
export interface HimaUsersResponse {
    /**
     * 
     * @type {Array<UserWrapper>}
     * @memberof HimaUsersResponse
     */
    himaUsers?: Array<UserWrapper> | null;
}

/**
 * Check if a given object implements the HimaUsersResponse interface.
 */
export function instanceOfHimaUsersResponse(value: object): value is HimaUsersResponse {
    return true;
}

export function HimaUsersResponseFromJSON(json: any): HimaUsersResponse {
    return HimaUsersResponseFromJSONTyped(json, false);
}

export function HimaUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): HimaUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'himaUsers': json['hima_users'] == null ? undefined : ((json['hima_users'] as Array<any>).map(UserWrapperFromJSON)),
    };
}

export function HimaUsersResponseToJSON(json: any): HimaUsersResponse {
    return HimaUsersResponseToJSONTyped(json, false);
}

export function HimaUsersResponseToJSONTyped(value?: HimaUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hima_users': value['himaUsers'] == null ? undefined : ((value['himaUsers'] as Array<any>).map(UserWrapperToJSON)),
    };
}

