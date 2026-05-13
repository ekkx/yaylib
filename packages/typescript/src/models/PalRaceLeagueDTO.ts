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
/**
 * 
 * @export
 * @interface PalRaceLeagueDTO
 */
export interface PalRaceLeagueDTO {
    /**
     * 
     * @type {boolean}
     * @memberof PalRaceLeagueDTO
     */
    accessible?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceLeagueDTO
     */
    entryFee?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceLeagueDTO
     */
    gameEndAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceLeagueDTO
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceLeagueDTO
     */
    priceReward?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceLeagueDTO
     */
    registrationCloseAt?: number | null;
}

/**
 * Check if a given object implements the PalRaceLeagueDTO interface.
 */
export function instanceOfPalRaceLeagueDTO(value: object): value is PalRaceLeagueDTO {
    return true;
}

export function PalRaceLeagueDTOFromJSON(json: any): PalRaceLeagueDTO {
    return PalRaceLeagueDTOFromJSONTyped(json, false);
}

export function PalRaceLeagueDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceLeagueDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'accessible': json['accessible'] == null ? undefined : json['accessible'],
        'entryFee': json['entry_fee'] == null ? undefined : json['entry_fee'],
        'gameEndAt': json['game_end_at'] == null ? undefined : json['game_end_at'],
        'name': json['name'] == null ? undefined : json['name'],
        'priceReward': json['price_reward'] == null ? undefined : json['price_reward'],
        'registrationCloseAt': json['registration_close_at'] == null ? undefined : json['registration_close_at'],
    };
}

export function PalRaceLeagueDTOToJSON(json: any): PalRaceLeagueDTO {
    return PalRaceLeagueDTOToJSONTyped(json, false);
}

export function PalRaceLeagueDTOToJSONTyped(value?: PalRaceLeagueDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'accessible': value['accessible'],
        'entry_fee': value['entryFee'],
        'game_end_at': value['gameEndAt'],
        'name': value['name'],
        'price_reward': value['priceReward'],
        'registration_close_at': value['registrationCloseAt'],
    };
}

