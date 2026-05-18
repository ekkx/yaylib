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
import type { RealmConferenceCall } from './RealmConferenceCall.js';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall.js';
import type { Post } from './Post.js';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
    PostToJSONTyped,
} from './Post.js';

/**
 * 
 * @export
 * @interface CreatePostResponse
 */
export interface CreatePostResponse {
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof CreatePostResponse
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {Post}
     * @memberof CreatePostResponse
     */
    post?: Post | null;
}

/**
 * Check if a given object implements the CreatePostResponse interface.
 */
export function instanceOfCreatePostResponse(value: object): value is CreatePostResponse {
    return true;
}

export function CreatePostResponseFromJSON(json: any): CreatePostResponse {
    return CreatePostResponseFromJSONTyped(json, false);
}

export function CreatePostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreatePostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'post': json['post'] == null ? undefined : PostFromJSON(json['post']),
    };
}

export function CreatePostResponseToJSON(json: any): CreatePostResponse {
    return CreatePostResponseToJSONTyped(json, false);
}

export function CreatePostResponseToJSONTyped(value?: CreatePostResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'post': PostToJSON(value['post']),
    };
}

