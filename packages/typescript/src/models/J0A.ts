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
export const J0A = {
    PublicTimeline: 'PublicTimeline',
    FollowingTimeline: 'FollowingTimeline',
    GroupTimeline: 'GroupTimeline'
} as const;
export type J0A = typeof J0A[keyof typeof J0A];


export function instanceOfJ0A(value: any): boolean {
    for (const key in J0A) {
        if (Object.prototype.hasOwnProperty.call(J0A, key)) {
            if (J0A[key as keyof typeof J0A] === value) {
                return true;
            }
        }
    }
    return false;
}

export function J0AFromJSON(json: any): J0A {
    return J0AFromJSONTyped(json, false);
}

export function J0AFromJSONTyped(json: any, ignoreDiscriminator: boolean): J0A {
    return json as J0A;
}

export function J0AToJSON(value?: J0A | null): any {
    return value as any;
}

export function J0AToJSONTyped(value: any, ignoreDiscriminator: boolean): J0A {
    return value as J0A;
}

