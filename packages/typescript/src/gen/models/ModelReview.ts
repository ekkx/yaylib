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
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';

/**
 * 
 * @export
 * @interface ModelReview
 */
export interface ModelReview {
    /**
     * 
     * @type {string}
     * @memberof ModelReview
     */
    comment?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelReview
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelReview
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelReview
     */
    isMutual?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelReview
     */
    isPinned?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof ModelReview
     */
    reportedCount?: number | null;
    /**
     * 
     * @type {User}
     * @memberof ModelReview
     */
    reviewer?: User | null;
    /**
     * 
     * @type {User}
     * @memberof ModelReview
     */
    user?: User | null;
}

/**
 * Check if a given object implements the ModelReview interface.
 */
export function instanceOfModelReview(value: object): value is ModelReview {
    return true;
}

export function ModelReviewFromJSON(json: any): ModelReview {
    return ModelReviewFromJSONTyped(json, false);
}

export function ModelReviewFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelReview {
    if (json == null) {
        return json;
    }
    return {
        
        'comment': json['comment'] == null ? undefined : json['comment'],
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'id': json['id'] == null ? undefined : json['id'],
        'isMutual': json['is_mutual'] == null ? undefined : json['is_mutual'],
        'isPinned': json['is_pinned'] == null ? undefined : json['is_pinned'],
        'reportedCount': json['reported_count'] == null ? undefined : json['reported_count'],
        'reviewer': json['reviewer'] == null ? undefined : UserFromJSON(json['reviewer']),
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
    };
}

export function ModelReviewToJSON(json: any): ModelReview {
    return ModelReviewToJSONTyped(json, false);
}

export function ModelReviewToJSONTyped(value?: ModelReview | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'comment': value['comment'],
        'created_at_millis': value['createdAtMillis'],
        'id': value['id'],
        'is_mutual': value['isMutual'],
        'is_pinned': value['isPinned'],
        'reported_count': value['reportedCount'],
        'reviewer': UserToJSON(value['reviewer']),
        'user': UserToJSON(value['user']),
    };
}

