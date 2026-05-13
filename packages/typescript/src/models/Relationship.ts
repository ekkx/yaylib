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
import type { Target } from './Target';
import {
    TargetFromJSON,
    TargetFromJSONTyped,
    TargetToJSON,
    TargetToJSONTyped,
} from './Target';

/**
 * 
 * @export
 * @interface Relationship
 */
export interface Relationship {
    /**
     * 
     * @type {Target}
     * @memberof Relationship
     */
    target?: Target | null;
}

/**
 * Check if a given object implements the Relationship interface.
 */
export function instanceOfRelationship(value: object): value is Relationship {
    return true;
}

export function RelationshipFromJSON(json: any): Relationship {
    return RelationshipFromJSONTyped(json, false);
}

export function RelationshipFromJSONTyped(json: any, ignoreDiscriminator: boolean): Relationship {
    if (json == null) {
        return json;
    }
    return {
        
        'target': json['target'] == null ? undefined : TargetFromJSON(json['target']),
    };
}

export function RelationshipToJSON(json: any): Relationship {
    return RelationshipToJSONTyped(json, false);
}

export function RelationshipToJSONTyped(value?: Relationship | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'target': TargetToJSON(value['target']),
    };
}

