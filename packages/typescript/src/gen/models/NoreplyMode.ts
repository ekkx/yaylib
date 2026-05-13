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
export const NoreplyMode = {
    False: '',
    True: 'noreply_'
} as const;
export type NoreplyMode = typeof NoreplyMode[keyof typeof NoreplyMode];


export function instanceOfNoreplyMode(value: any): boolean {
    for (const key in NoreplyMode) {
        if (Object.prototype.hasOwnProperty.call(NoreplyMode, key)) {
            if (NoreplyMode[key as keyof typeof NoreplyMode] === value) {
                return true;
            }
        }
    }
    return false;
}

export function NoreplyModeFromJSON(json: any): NoreplyMode {
    return NoreplyModeFromJSONTyped(json, false);
}

export function NoreplyModeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NoreplyMode {
    return json as NoreplyMode;
}

export function NoreplyModeToJSON(value?: NoreplyMode | null): any {
    return value as any;
}

export function NoreplyModeToJSONTyped(value: any, ignoreDiscriminator: boolean): NoreplyMode {
    return value as NoreplyMode;
}

