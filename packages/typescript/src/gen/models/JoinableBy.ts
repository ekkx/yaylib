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


/**
 * 
 * @export
 */
export const JoinableBy = {
    Anyone: 'anyone',
    Users: 'users',
    FollowersOnly: 'followers_only',
    MutualFollowsOnly: 'mutual_follows_only',
    InviteOnly: 'invite_only',
    NoOne: 'noone'
} as const;
export type JoinableBy = typeof JoinableBy[keyof typeof JoinableBy];


export function instanceOfJoinableBy(value: any): boolean {
    for (const key in JoinableBy) {
        if (Object.prototype.hasOwnProperty.call(JoinableBy, key)) {
            if (JoinableBy[key as keyof typeof JoinableBy] === value) {
                return true;
            }
        }
    }
    return false;
}

export function JoinableByFromJSON(json: any): JoinableBy {
    return JoinableByFromJSONTyped(json, false);
}

export function JoinableByFromJSONTyped(json: any, ignoreDiscriminator: boolean): JoinableBy {
    return json as JoinableBy;
}

export function JoinableByToJSON(value?: JoinableBy | null): any {
    return value as any;
}

export function JoinableByToJSONTyped(value: any, ignoreDiscriminator: boolean): JoinableBy {
    return value as JoinableBy;
}

