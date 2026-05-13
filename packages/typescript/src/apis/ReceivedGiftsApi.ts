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
  GiftReceivedResponse,
  GiftReceivedTransactionResponse,
  GiftSendersResponse,
} from '../models/index';
import {
    GiftReceivedResponseFromJSON,
    GiftReceivedResponseToJSON,
    GiftReceivedTransactionResponseFromJSON,
    GiftReceivedTransactionResponseToJSON,
    GiftSendersResponseFromJSON,
    GiftSendersResponseToJSON,
} from '../models/index';

export interface GetReceivedGiftSendersRequest {
    giftId: number;
    number?: number | null;
}

export interface GetReceivedGiftTransactionRequest {
    giftTransactionUuid: string;
}

export interface ListReceivedGiftsV1Request {
    from?: string | null;
    number?: number | null;
}

/**
 * 
 */
export class ReceivedGiftsApi extends runtime.BaseAPI {

    /**
     */
    async getReceivedGiftSendersRaw(requestParameters: GetReceivedGiftSendersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftSendersResponse>> {
        if (requestParameters['giftId'] == null) {
            throw new runtime.RequiredError(
                'giftId',
                'Required parameter "giftId" was null or undefined when calling getReceivedGiftSenders().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/received_gifts/{gift_id}/senders`.replace(`{${"gift_id"}}`, encodeURIComponent(String(requestParameters['giftId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftSendersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getReceivedGiftSenders(requestParameters: GetReceivedGiftSendersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftSendersResponse> {
        const response = await this.getReceivedGiftSendersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getReceivedGiftTransactionRaw(requestParameters: GetReceivedGiftTransactionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftReceivedTransactionResponse>> {
        if (requestParameters['giftTransactionUuid'] == null) {
            throw new runtime.RequiredError(
                'giftTransactionUuid',
                'Required parameter "giftTransactionUuid" was null or undefined when calling getReceivedGiftTransaction().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/received_gifts/{gift_transaction_uuid}`.replace(`{${"gift_transaction_uuid"}}`, encodeURIComponent(String(requestParameters['giftTransactionUuid']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftReceivedTransactionResponseFromJSON(jsonValue));
    }

    /**
     */
    async getReceivedGiftTransaction(requestParameters: GetReceivedGiftTransactionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftReceivedTransactionResponse> {
        const response = await this.getReceivedGiftTransactionRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listReceivedGiftsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftReceivedResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/received_gifts`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftReceivedResponseFromJSON(jsonValue));
    }

    /**
     */
    async listReceivedGifts(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftReceivedResponse> {
        const response = await this.listReceivedGiftsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async listReceivedGiftsV1Raw(requestParameters: ListReceivedGiftsV1Request, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftReceivedResponse>> {
        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/received_gifts`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftReceivedResponseFromJSON(jsonValue));
    }

    /**
     */
    async listReceivedGiftsV1(requestParameters: ListReceivedGiftsV1Request = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftReceivedResponse> {
        const response = await this.listReceivedGiftsV1Raw(requestParameters, initOverrides);
        return await response.value();
    }

}
