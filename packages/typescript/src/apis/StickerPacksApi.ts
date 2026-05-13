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


import * as runtime from '../runtime';
import type {
  StickerPacksResponse,
} from '../models/index';
import {
    StickerPacksResponseFromJSON,
    StickerPacksResponseToJSON,
} from '../models/index';

/**
 * 
 */
export class StickerPacksApi extends runtime.BaseAPI {

    /**
     */
    async listStickerPacksRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<StickerPacksResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/sticker_packs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StickerPacksResponseFromJSON(jsonValue));
    }

    /**
     */
    async listStickerPacks(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<StickerPacksResponse> {
        const response = await this.listStickerPacksRaw(initOverrides);
        return await response.value();
    }

}
