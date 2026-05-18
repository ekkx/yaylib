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
import type { Review } from './Review.js';
import {
    ReviewFromJSON,
    ReviewFromJSONTyped,
    ReviewToJSON,
    ReviewToJSONTyped,
} from './Review.js';

/**
 * 
 * @export
 * @interface ReviewsResponse
 */
export interface ReviewsResponse {
    /**
     * 
     * @type {Array<Review>}
     * @memberof ReviewsResponse
     */
    pinnedReviews?: Array<Review> | null;
    /**
     * 
     * @type {Array<Review>}
     * @memberof ReviewsResponse
     */
    reviews?: Array<Review> | null;
}

/**
 * Check if a given object implements the ReviewsResponse interface.
 */
export function instanceOfReviewsResponse(value: object): value is ReviewsResponse {
    return true;
}

export function ReviewsResponseFromJSON(json: any): ReviewsResponse {
    return ReviewsResponseFromJSONTyped(json, false);
}

export function ReviewsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReviewsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'pinnedReviews': json['pinned_reviews'] == null ? undefined : ((json['pinned_reviews'] as Array<any>).map(ReviewFromJSON)),
        'reviews': json['reviews'] == null ? undefined : ((json['reviews'] as Array<any>).map(ReviewFromJSON)),
    };
}

export function ReviewsResponseToJSON(json: any): ReviewsResponse {
    return ReviewsResponseToJSONTyped(json, false);
}

export function ReviewsResponseToJSONTyped(value?: ReviewsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pinned_reviews': value['pinnedReviews'] == null ? undefined : ((value['pinnedReviews'] as Array<any>).map(ReviewToJSON)),
        'reviews': value['reviews'] == null ? undefined : ((value['reviews'] as Array<any>).map(ReviewToJSON)),
    };
}

