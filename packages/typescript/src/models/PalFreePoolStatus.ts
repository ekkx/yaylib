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
import type { Qualification } from './Qualification';
import {
    QualificationFromJSON,
    QualificationFromJSONTyped,
    QualificationToJSON,
    QualificationToJSONTyped,
} from './Qualification';
import type { Cooldown } from './Cooldown';
import {
    CooldownFromJSON,
    CooldownFromJSONTyped,
    CooldownToJSON,
    CooldownToJSONTyped,
} from './Cooldown';

/**
 * 
 * @export
 * @interface PalFreePoolStatus
 */
export interface PalFreePoolStatus {
    /**
     * 
     * @type {Cooldown}
     * @memberof PalFreePoolStatus
     */
    cooldown?: Cooldown | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalFreePoolStatus
     */
    palsAvailable?: boolean | null;
    /**
     * 
     * @type {Array<Qualification>}
     * @memberof PalFreePoolStatus
     */
    qualifications?: Array<Qualification> | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalFreePoolStatus
     */
    qualified?: boolean | null;
}

/**
 * Check if a given object implements the PalFreePoolStatus interface.
 */
export function instanceOfPalFreePoolStatus(value: object): value is PalFreePoolStatus {
    return true;
}

export function PalFreePoolStatusFromJSON(json: any): PalFreePoolStatus {
    return PalFreePoolStatusFromJSONTyped(json, false);
}

export function PalFreePoolStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalFreePoolStatus {
    if (json == null) {
        return json;
    }
    return {
        
        'cooldown': json['cooldown'] == null ? undefined : CooldownFromJSON(json['cooldown']),
        'palsAvailable': json['pals_available'] == null ? undefined : json['pals_available'],
        'qualifications': json['qualifications'] == null ? undefined : ((json['qualifications'] as Array<any>).map(QualificationFromJSON)),
        'qualified': json['qualified'] == null ? undefined : json['qualified'],
    };
}

export function PalFreePoolStatusToJSON(json: any): PalFreePoolStatus {
    return PalFreePoolStatusToJSONTyped(json, false);
}

export function PalFreePoolStatusToJSONTyped(value?: PalFreePoolStatus | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cooldown': CooldownToJSON(value['cooldown']),
        'pals_available': value['palsAvailable'],
        'qualifications': value['qualifications'] == null ? undefined : ((value['qualifications'] as Array<any>).map(QualificationToJSON)),
        'qualified': value['qualified'],
    };
}

