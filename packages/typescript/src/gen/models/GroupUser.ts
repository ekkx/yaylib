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
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface GroupUser
 */
export interface GroupUser {
    /**
     * 
     * @type {boolean}
     * @memberof GroupUser
     */
    isModerator?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GroupUser
     */
    isMute?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GroupUser
     */
    pendingDeputize?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GroupUser
     */
    pendingTransfer?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof GroupUser
     */
    title?: string | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof GroupUser
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the GroupUser interface.
 */
export function instanceOfGroupUser(value: object): value is GroupUser {
    return true;
}

export function GroupUserFromJSON(json: any): GroupUser {
    return GroupUserFromJSONTyped(json, false);
}

export function GroupUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupUser {
    if (json == null) {
        return json;
    }
    return {
        
        'isModerator': json['is_moderator'] == null ? undefined : json['is_moderator'],
        'isMute': json['is_mute'] == null ? undefined : json['is_mute'],
        'pendingDeputize': json['pending_deputize'] == null ? undefined : json['pending_deputize'],
        'pendingTransfer': json['pending_transfer'] == null ? undefined : json['pending_transfer'],
        'title': json['title'] == null ? undefined : json['title'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function GroupUserToJSON(json: any): GroupUser {
    return GroupUserToJSONTyped(json, false);
}

export function GroupUserToJSONTyped(value?: GroupUser | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_moderator': value['isModerator'],
        'is_mute': value['isMute'],
        'pending_deputize': value['pendingDeputize'],
        'pending_transfer': value['pendingTransfer'],
        'title': value['title'],
        'user': RealmUserToJSON(value['user']),
    };
}

