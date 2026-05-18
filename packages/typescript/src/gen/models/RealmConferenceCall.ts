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
import type { RealmGenre } from './RealmGenre.js';
import {
    RealmGenreFromJSON,
    RealmGenreFromJSONTyped,
    RealmGenreToJSON,
    RealmGenreToJSONTyped,
} from './RealmGenre.js';
import type { ConferenceCallBumpParams } from './ConferenceCallBumpParams.js';
import {
    ConferenceCallBumpParamsFromJSON,
    ConferenceCallBumpParamsFromJSONTyped,
    ConferenceCallBumpParamsToJSON,
    ConferenceCallBumpParamsToJSONTyped,
} from './ConferenceCallBumpParams.js';
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';
import type { RealmGame } from './RealmGame.js';
import {
    RealmGameFromJSON,
    RealmGameFromJSONTyped,
    RealmGameToJSON,
    RealmGameToJSONTyped,
} from './RealmGame.js';
import type { ConferenceCallUserRole } from './ConferenceCallUserRole.js';
import {
    ConferenceCallUserRoleFromJSON,
    ConferenceCallUserRoleFromJSONTyped,
    ConferenceCallUserRoleToJSON,
    ConferenceCallUserRoleToJSONTyped,
} from './ConferenceCallUserRole.js';

/**
 * 
 * @export
 * @interface RealmConferenceCall
 */
export interface RealmConferenceCall {
    /**
     * 
     * @type {boolean}
     * @memberof RealmConferenceCall
     */
    active?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RealmConferenceCall
     */
    agoraChannel?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmConferenceCall
     */
    agoraToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    anonymousCallUsersCount?: number | null;
    /**
     * 
     * @type {ConferenceCallBumpParams}
     * @memberof RealmConferenceCall
     */
    bumpParams?: ConferenceCallBumpParams | null;
    /**
     * 
     * @type {string}
     * @memberof RealmConferenceCall
     */
    callType?: string | null;
    /**
     * 
     * @type {Array<ConferenceCallUserRole>}
     * @memberof RealmConferenceCall
     */
    conferenceCallUserRoles?: Array<ConferenceCallUserRole> | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof RealmConferenceCall
     */
    conferenceCallUsers?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    conferenceCallUsersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    durationSeconds?: number | null;
    /**
     * 
     * @type {RealmGame}
     * @memberof RealmConferenceCall
     */
    game?: RealmGame | null;
    /**
     * 
     * @type {RealmGenre}
     * @memberof RealmConferenceCall
     */
    genre?: RealmGenre | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    groupId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmConferenceCall
     */
    joinableBy?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    maxParticipants?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmConferenceCall
     */
    postId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmConferenceCall
     */
    server?: string | null;
}

/**
 * Check if a given object implements the RealmConferenceCall interface.
 */
export function instanceOfRealmConferenceCall(value: object): value is RealmConferenceCall {
    return true;
}

export function RealmConferenceCallFromJSON(json: any): RealmConferenceCall {
    return RealmConferenceCallFromJSONTyped(json, false);
}

export function RealmConferenceCallFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmConferenceCall {
    if (json == null) {
        return json;
    }
    return {
        
        'active': json['active'] == null ? undefined : json['active'],
        'agoraChannel': json['agora_channel'] == null ? undefined : json['agora_channel'],
        'agoraToken': json['agora_token'] == null ? undefined : json['agora_token'],
        'anonymousCallUsersCount': json['anonymous_call_users_count'] == null ? undefined : json['anonymous_call_users_count'],
        'bumpParams': json['bump_params'] == null ? undefined : ConferenceCallBumpParamsFromJSON(json['bump_params']),
        'callType': json['call_type'] == null ? undefined : json['call_type'],
        'conferenceCallUserRoles': json['conference_call_user_roles'] == null ? undefined : ((json['conference_call_user_roles'] as Array<any>).map(ConferenceCallUserRoleFromJSON)),
        'conferenceCallUsers': json['conference_call_users'] == null ? undefined : ((json['conference_call_users'] as Array<any>).map(RealmUserFromJSON)),
        'conferenceCallUsersCount': json['conference_call_users_count'] == null ? undefined : json['conference_call_users_count'],
        'durationSeconds': json['duration_seconds'] == null ? undefined : json['duration_seconds'],
        'game': json['game'] == null ? undefined : RealmGameFromJSON(json['game']),
        'genre': json['genre'] == null ? undefined : RealmGenreFromJSON(json['genre']),
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
        'id': json['id'] == null ? undefined : json['id'],
        'joinableBy': json['joinable_by'] == null ? undefined : json['joinable_by'],
        'maxParticipants': json['max_participants'] == null ? undefined : json['max_participants'],
        'postId': json['post_id'] == null ? undefined : json['post_id'],
        'server': json['server'] == null ? undefined : json['server'],
    };
}

export function RealmConferenceCallToJSON(json: any): RealmConferenceCall {
    return RealmConferenceCallToJSONTyped(json, false);
}

export function RealmConferenceCallToJSONTyped(value?: RealmConferenceCall | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'agora_channel': value['agoraChannel'],
        'agora_token': value['agoraToken'],
        'anonymous_call_users_count': value['anonymousCallUsersCount'],
        'bump_params': ConferenceCallBumpParamsToJSON(value['bumpParams']),
        'call_type': value['callType'],
        'conference_call_user_roles': value['conferenceCallUserRoles'] == null ? undefined : ((value['conferenceCallUserRoles'] as Array<any>).map(ConferenceCallUserRoleToJSON)),
        'conference_call_users': value['conferenceCallUsers'] == null ? undefined : ((value['conferenceCallUsers'] as Array<any>).map(RealmUserToJSON)),
        'conference_call_users_count': value['conferenceCallUsersCount'],
        'duration_seconds': value['durationSeconds'],
        'game': RealmGameToJSON(value['game']),
        'genre': RealmGenreToJSON(value['genre']),
        'group_id': value['groupId'],
        'id': value['id'],
        'joinable_by': value['joinableBy'],
        'max_participants': value['maxParticipants'],
        'post_id': value['postId'],
        'server': value['server'],
    };
}

