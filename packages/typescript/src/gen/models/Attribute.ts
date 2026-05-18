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
 * @interface Attribute
 */
export interface Attribute {
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    birthday?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    body?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    climbing?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    eligibleLeague?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    fatigue?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    generation?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    grade?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    head?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    item?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    level?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    luck?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    momentum?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    power?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    racingstyle?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    rank?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    running?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    skin?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    speed?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    stamina?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    swimming?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Attribute
     */
    technique?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    tribe?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Attribute
     */
    type?: string | null;
}

/**
 * Check if a given object implements the Attribute interface.
 */
export function instanceOfAttribute(value: object): value is Attribute {
    return true;
}

export function AttributeFromJSON(json: any): Attribute {
    return AttributeFromJSONTyped(json, false);
}

export function AttributeFromJSONTyped(json: any, ignoreDiscriminator: boolean): Attribute {
    if (json == null) {
        return json;
    }
    return {
        
        'birthday': json['birthday'] == null ? undefined : json['birthday'],
        'body': json['body'] == null ? undefined : json['body'],
        'climbing': json['climbing'] == null ? undefined : json['climbing'],
        'eligibleLeague': json['eligible_league'] == null ? undefined : json['eligible_league'],
        'fatigue': json['fatigue'] == null ? undefined : json['fatigue'],
        'generation': json['generation'] == null ? undefined : json['generation'],
        'grade': json['grade'] == null ? undefined : json['grade'],
        'head': json['head'] == null ? undefined : json['head'],
        'item': json['item'] == null ? undefined : json['item'],
        'level': json['level'] == null ? undefined : json['level'],
        'luck': json['luck'] == null ? undefined : json['luck'],
        'momentum': json['momentum'] == null ? undefined : json['momentum'],
        'name': json['name'] == null ? undefined : json['name'],
        'power': json['power'] == null ? undefined : json['power'],
        'racingstyle': json['racingstyle'] == null ? undefined : json['racingstyle'],
        'rank': json['rank'] == null ? undefined : json['rank'],
        'running': json['running'] == null ? undefined : json['running'],
        'skin': json['skin'] == null ? undefined : json['skin'],
        'speed': json['speed'] == null ? undefined : json['speed'],
        'stamina': json['stamina'] == null ? undefined : json['stamina'],
        'status': json['status'] == null ? undefined : json['status'],
        'swimming': json['swimming'] == null ? undefined : json['swimming'],
        'technique': json['technique'] == null ? undefined : json['technique'],
        'tribe': json['tribe'] == null ? undefined : json['tribe'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function AttributeToJSON(json: any): Attribute {
    return AttributeToJSONTyped(json, false);
}

export function AttributeToJSONTyped(value?: Attribute | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'birthday': value['birthday'],
        'body': value['body'],
        'climbing': value['climbing'],
        'eligible_league': value['eligibleLeague'],
        'fatigue': value['fatigue'],
        'generation': value['generation'],
        'grade': value['grade'],
        'head': value['head'],
        'item': value['item'],
        'level': value['level'],
        'luck': value['luck'],
        'momentum': value['momentum'],
        'name': value['name'],
        'power': value['power'],
        'racingstyle': value['racingstyle'],
        'rank': value['rank'],
        'running': value['running'],
        'skin': value['skin'],
        'speed': value['speed'],
        'stamina': value['stamina'],
        'status': value['status'],
        'swimming': value['swimming'],
        'technique': value['technique'],
        'tribe': value['tribe'],
        'type': value['type'],
    };
}

