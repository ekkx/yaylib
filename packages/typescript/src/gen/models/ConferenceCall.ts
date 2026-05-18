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
import type { BumpParams } from './BumpParams.js';
import {
    BumpParamsFromJSON,
    BumpParamsFromJSONTyped,
    BumpParamsToJSON,
    BumpParamsToJSONTyped,
} from './BumpParams.js';
import type { Game } from './Game.js';
import {
    GameFromJSON,
    GameFromJSONTyped,
    GameToJSON,
    GameToJSONTyped,
} from './Game.js';
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';
import type { Genre } from './Genre.js';
import {
    GenreFromJSON,
    GenreFromJSONTyped,
    GenreToJSON,
    GenreToJSONTyped,
} from './Genre.js';
import type { ModelConferenceCallUserRole } from './ModelConferenceCallUserRole.js';
import {
    ModelConferenceCallUserRoleFromJSON,
    ModelConferenceCallUserRoleFromJSONTyped,
    ModelConferenceCallUserRoleToJSON,
    ModelConferenceCallUserRoleToJSONTyped,
} from './ModelConferenceCallUserRole.js';

/**
 * 
 * @export
 * @interface ConferenceCall
 */
export interface ConferenceCall {
    /**
     * 
     * @type {string}
     * @memberof ConferenceCall
     */
    agoraChannel?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCall
     */
    agoraToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    anonymousCallUsersCount?: number | null;
    /**
     * 
     * @type {BumpParams}
     * @memberof ConferenceCall
     */
    bumpParams?: BumpParams | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCall
     */
    callType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCall
     */
    conferenceCallUserUuid?: string | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ConferenceCall
     */
    conferenceCallUsers?: Array<User> | null;
    /**
     * 
     * @type {Array<ModelConferenceCallUserRole>}
     * @memberof ConferenceCall
     */
    conferenceCallUsersRole?: Array<ModelConferenceCallUserRole> | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    durationSeconds?: number | null;
    /**
     * 
     * @type {Game}
     * @memberof ConferenceCall
     */
    game?: Game | null;
    /**
     * 
     * @type {Genre}
     * @memberof ConferenceCall
     */
    genre?: Genre | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    groupId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ConferenceCall
     */
    isActive?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCall
     */
    joinableBy?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    maxParticipants?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCall
     */
    postId?: number | null;
}

/**
 * Check if a given object implements the ConferenceCall interface.
 */
export function instanceOfConferenceCall(value: object): value is ConferenceCall {
    return true;
}

export function ConferenceCallFromJSON(json: any): ConferenceCall {
    return ConferenceCallFromJSONTyped(json, false);
}

export function ConferenceCallFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConferenceCall {
    if (json == null) {
        return json;
    }
    return {
        
        'agoraChannel': json['agora_channel'] == null ? undefined : json['agora_channel'],
        'agoraToken': json['agora_token'] == null ? undefined : json['agora_token'],
        'anonymousCallUsersCount': json['anonymous_call_users_count'] == null ? undefined : json['anonymous_call_users_count'],
        'bumpParams': json['bump_params'] == null ? undefined : BumpParamsFromJSON(json['bump_params']),
        'callType': json['call_type'] == null ? undefined : json['call_type'],
        'conferenceCallUserUuid': json['conference_call_user_uuid'] == null ? undefined : json['conference_call_user_uuid'],
        'conferenceCallUsers': json['conference_call_users'] == null ? undefined : ((json['conference_call_users'] as Array<any>).map(UserFromJSON)),
        'conferenceCallUsersRole': json['conference_call_users_role'] == null ? undefined : ((json['conference_call_users_role'] as Array<any>).map(ModelConferenceCallUserRoleFromJSON)),
        'durationSeconds': json['duration_seconds'] == null ? undefined : json['duration_seconds'],
        'game': json['game'] == null ? undefined : GameFromJSON(json['game']),
        'genre': json['genre'] == null ? undefined : GenreFromJSON(json['genre']),
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
        'id': json['id'] == null ? undefined : json['id'],
        'isActive': json['is_active'] == null ? undefined : json['is_active'],
        'joinableBy': json['joinable_by'] == null ? undefined : json['joinable_by'],
        'maxParticipants': json['max_participants'] == null ? undefined : json['max_participants'],
        'postId': json['post_id'] == null ? undefined : json['post_id'],
    };
}

export function ConferenceCallToJSON(json: any): ConferenceCall {
    return ConferenceCallToJSONTyped(json, false);
}

export function ConferenceCallToJSONTyped(value?: ConferenceCall | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'agora_channel': value['agoraChannel'],
        'agora_token': value['agoraToken'],
        'anonymous_call_users_count': value['anonymousCallUsersCount'],
        'bump_params': BumpParamsToJSON(value['bumpParams']),
        'call_type': value['callType'],
        'conference_call_user_uuid': value['conferenceCallUserUuid'],
        'conference_call_users': value['conferenceCallUsers'] == null ? undefined : ((value['conferenceCallUsers'] as Array<any>).map(UserToJSON)),
        'conference_call_users_role': value['conferenceCallUsersRole'] == null ? undefined : ((value['conferenceCallUsersRole'] as Array<any>).map(ModelConferenceCallUserRoleToJSON)),
        'duration_seconds': value['durationSeconds'],
        'game': GameToJSON(value['game']),
        'genre': GenreToJSON(value['genre']),
        'group_id': value['groupId'],
        'id': value['id'],
        'is_active': value['isActive'],
        'joinable_by': value['joinableBy'],
        'max_participants': value['maxParticipants'],
        'post_id': value['postId'],
    };
}

