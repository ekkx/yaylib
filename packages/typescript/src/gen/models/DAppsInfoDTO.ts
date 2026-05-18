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
import type { DAppDTO } from './DAppDTO.js';
import {
    DAppDTOFromJSON,
    DAppDTOFromJSONTyped,
    DAppDTOToJSON,
    DAppDTOToJSONTyped,
} from './DAppDTO.js';

/**
 * 
 * @export
 * @interface DAppsInfoDTO
 */
export interface DAppsInfoDTO {
    /**
     * 
     * @type {Array<DAppDTO>}
     * @memberof DAppsInfoDTO
     */
    dapps?: Array<DAppDTO> | null;
    /**
     * 
     * @type {Array<DAppDTO>}
     * @memberof DAppsInfoDTO
     */
    educationalLinks?: Array<DAppDTO> | null;
}

/**
 * Check if a given object implements the DAppsInfoDTO interface.
 */
export function instanceOfDAppsInfoDTO(value: object): value is DAppsInfoDTO {
    return true;
}

export function DAppsInfoDTOFromJSON(json: any): DAppsInfoDTO {
    return DAppsInfoDTOFromJSONTyped(json, false);
}

export function DAppsInfoDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): DAppsInfoDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'dapps': json['dapps'] == null ? undefined : ((json['dapps'] as Array<any>).map(DAppDTOFromJSON)),
        'educationalLinks': json['educational_links'] == null ? undefined : ((json['educational_links'] as Array<any>).map(DAppDTOFromJSON)),
    };
}

export function DAppsInfoDTOToJSON(json: any): DAppsInfoDTO {
    return DAppsInfoDTOToJSONTyped(json, false);
}

export function DAppsInfoDTOToJSONTyped(value?: DAppsInfoDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'dapps': value['dapps'] == null ? undefined : ((value['dapps'] as Array<any>).map(DAppDTOToJSON)),
        'educational_links': value['educationalLinks'] == null ? undefined : ((value['educationalLinks'] as Array<any>).map(DAppDTOToJSON)),
    };
}

