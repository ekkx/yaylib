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
 * @interface StateChanges
 */
export interface StateChanges {
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    attack?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    defense?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    energy?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    luck?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    power?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    speed?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    stamina?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StateChanges
     */
    technique?: number | null;
}

/**
 * Check if a given object implements the StateChanges interface.
 */
export function instanceOfStateChanges(value: object): value is StateChanges {
    return true;
}

export function StateChangesFromJSON(json: any): StateChanges {
    return StateChangesFromJSONTyped(json, false);
}

export function StateChangesFromJSONTyped(json: any, ignoreDiscriminator: boolean): StateChanges {
    if (json == null) {
        return json;
    }
    return {
        
        'attack': json['attack'] == null ? undefined : json['attack'],
        'defense': json['defense'] == null ? undefined : json['defense'],
        'energy': json['energy'] == null ? undefined : json['energy'],
        'luck': json['luck'] == null ? undefined : json['luck'],
        'power': json['power'] == null ? undefined : json['power'],
        'speed': json['speed'] == null ? undefined : json['speed'],
        'stamina': json['stamina'] == null ? undefined : json['stamina'],
        'technique': json['technique'] == null ? undefined : json['technique'],
    };
}

export function StateChangesToJSON(json: any): StateChanges {
    return StateChangesToJSONTyped(json, false);
}

export function StateChangesToJSONTyped(value?: StateChanges | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attack': value['attack'],
        'defense': value['defense'],
        'energy': value['energy'],
        'luck': value['luck'],
        'power': value['power'],
        'speed': value['speed'],
        'stamina': value['stamina'],
        'technique': value['technique'],
    };
}

