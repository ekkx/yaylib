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
import type { RealmUser } from './RealmUser';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser';

/**
 * 
 * @export
 * @interface Review
 */
export interface Review {
    /**
     * 
     * @type {string}
     * @memberof Review
     */
    comment?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Review
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Review
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Review
     */
    mutualReview?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Review
     */
    reportedCount?: number | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof Review
     */
    reviewer?: RealmUser | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof Review
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the Review interface.
 */
export function instanceOfReview(value: object): value is Review {
    return true;
}

export function ReviewFromJSON(json: any): Review {
    return ReviewFromJSONTyped(json, false);
}

export function ReviewFromJSONTyped(json: any, ignoreDiscriminator: boolean): Review {
    if (json == null) {
        return json;
    }
    return {
        
        'comment': json['comment'] == null ? undefined : json['comment'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'mutualReview': json['mutual_review'] == null ? undefined : json['mutual_review'],
        'reportedCount': json['reported_count'] == null ? undefined : json['reported_count'],
        'reviewer': json['reviewer'] == null ? undefined : RealmUserFromJSON(json['reviewer']),
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function ReviewToJSON(json: any): Review {
    return ReviewToJSONTyped(json, false);
}

export function ReviewToJSONTyped(value?: Review | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'comment': value['comment'],
        'created_at': value['createdAt'],
        'id': value['id'],
        'mutual_review': value['mutualReview'],
        'reported_count': value['reportedCount'],
        'reviewer': RealmUserToJSON(value['reviewer']),
        'user': RealmUserToJSON(value['user']),
    };
}

