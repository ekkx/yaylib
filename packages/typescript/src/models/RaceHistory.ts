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
import type { RaceReward } from './RaceReward';
import {
    RaceRewardFromJSON,
    RaceRewardFromJSONTyped,
    RaceRewardToJSON,
    RaceRewardToJSONTyped,
} from './RaceReward';
import type { PalRank } from './PalRank';
import {
    PalRankFromJSON,
    PalRankFromJSONTyped,
    PalRankToJSON,
    PalRankToJSONTyped,
} from './PalRank';
import type { PalRaceLeagueDTO } from './PalRaceLeagueDTO';
import {
    PalRaceLeagueDTOFromJSON,
    PalRaceLeagueDTOFromJSONTyped,
    PalRaceLeagueDTOToJSON,
    PalRaceLeagueDTOToJSONTyped,
} from './PalRaceLeagueDTO';
import type { Pal } from './Pal';
import {
    PalFromJSON,
    PalFromJSONTyped,
    PalToJSON,
    PalToJSONTyped,
} from './Pal';
import type { Reward } from './Reward';
import {
    RewardFromJSON,
    RewardFromJSONTyped,
    RewardToJSON,
    RewardToJSONTyped,
} from './Reward';

/**
 * 
 * @export
 * @interface RaceHistory
 */
export interface RaceHistory {
    /**
     * 
     * @type {string}
     * @memberof RaceHistory
     */
    animationUrl?: string | null;
    /**
     * 
     * @type {PalRaceLeagueDTO}
     * @memberof RaceHistory
     */
    league?: PalRaceLeagueDTO | null;
    /**
     * 
     * @type {string}
     * @memberof RaceHistory
     */
    leagueName?: string | null;
    /**
     * 
     * @type {Array<Pal>}
     * @memberof RaceHistory
     */
    pals?: Array<Pal> | null;
    /**
     * 
     * @type {number}
     * @memberof RaceHistory
     */
    raceDate?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RaceHistory
     */
    raceEntryFee?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RaceHistory
     */
    raceId?: string | null;
    /**
     * 
     * @type {RaceReward}
     * @memberof RaceHistory
     */
    raceReward?: RaceReward | null;
    /**
     * 
     * @type {number}
     * @memberof RaceHistory
     */
    raceStartAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RaceHistory
     */
    raceStatus?: string | null;
    /**
     * 
     * @type {Array<PalRank>}
     * @memberof RaceHistory
     */
    results?: Array<PalRank> | null;
    /**
     * 
     * @type {Reward}
     * @memberof RaceHistory
     */
    reward?: Reward | null;
}

/**
 * Check if a given object implements the RaceHistory interface.
 */
export function instanceOfRaceHistory(value: object): value is RaceHistory {
    return true;
}

export function RaceHistoryFromJSON(json: any): RaceHistory {
    return RaceHistoryFromJSONTyped(json, false);
}

export function RaceHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): RaceHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'animationUrl': json['animation_url'] == null ? undefined : json['animation_url'],
        'league': json['league'] == null ? undefined : PalRaceLeagueDTOFromJSON(json['league']),
        'leagueName': json['league_name'] == null ? undefined : json['league_name'],
        'pals': json['pals'] == null ? undefined : ((json['pals'] as Array<any>).map(PalFromJSON)),
        'raceDate': json['race_date'] == null ? undefined : json['race_date'],
        'raceEntryFee': json['race_entry_fee'] == null ? undefined : json['race_entry_fee'],
        'raceId': json['race_id'] == null ? undefined : json['race_id'],
        'raceReward': json['race_reward'] == null ? undefined : RaceRewardFromJSON(json['race_reward']),
        'raceStartAt': json['race_start_at'] == null ? undefined : json['race_start_at'],
        'raceStatus': json['race_status'] == null ? undefined : json['race_status'],
        'results': json['results'] == null ? undefined : ((json['results'] as Array<any>).map(PalRankFromJSON)),
        'reward': json['reward'] == null ? undefined : RewardFromJSON(json['reward']),
    };
}

export function RaceHistoryToJSON(json: any): RaceHistory {
    return RaceHistoryToJSONTyped(json, false);
}

export function RaceHistoryToJSONTyped(value?: RaceHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'animation_url': value['animationUrl'],
        'league': PalRaceLeagueDTOToJSON(value['league']),
        'league_name': value['leagueName'],
        'pals': value['pals'] == null ? undefined : ((value['pals'] as Array<any>).map(PalToJSON)),
        'race_date': value['raceDate'],
        'race_entry_fee': value['raceEntryFee'],
        'race_id': value['raceId'],
        'race_reward': RaceRewardToJSON(value['raceReward']),
        'race_start_at': value['raceStartAt'],
        'race_status': value['raceStatus'],
        'results': value['results'] == null ? undefined : ((value['results'] as Array<any>).map(PalRankToJSON)),
        'reward': RewardToJSON(value['reward']),
    };
}

