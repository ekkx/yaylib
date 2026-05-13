// @ts-nocheck
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
  ChatRoomsResponse,
  CreateMuteKeywordResponse,
  HiddenResponse,
  MuteKeywordRequest,
  MuteKeywordResponse,
} from '../models/index';
import {
    ChatRoomsResponseFromJSON,
    ChatRoomsResponseToJSON,
    CreateMuteKeywordResponseFromJSON,
    CreateMuteKeywordResponseToJSON,
    HiddenResponseFromJSON,
    HiddenResponseToJSON,
    MuteKeywordRequestFromJSON,
    MuteKeywordRequestToJSON,
    MuteKeywordResponseFromJSON,
    MuteKeywordResponseToJSON,
} from '../models/index';

export interface CreateMuteKeywordRequest {
    muteKeywordRequest: MuteKeywordRequest;
}

export interface DeleteMuteKeywordRequest {
    ids?: Array<number>;
}

export interface HideChatsRequest {
    chatRoomId?: number;
}

export interface HideUsersRequest {
    userId?: number;
}

export interface ListHiddenChatsRequest {
    number?: number;
    fromTimestamp?: number | null;
}

export interface ListHiddenUsersRequest {
    from?: string | null;
    number?: number | null;
}

export interface UnhideChatsRequest {
    chatRoomIds?: number;
}

export interface UnhideUsersRequest {
    userIds?: Array<number>;
}

/**
 * 
 */
export class HiddenApi extends runtime.BaseAPI {

    /**
     */
    async createMuteKeywordRaw(requestParameters: CreateMuteKeywordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateMuteKeywordResponse>> {
        if (requestParameters['muteKeywordRequest'] == null) {
            throw new runtime.RequiredError(
                'muteKeywordRequest',
                'Required parameter "muteKeywordRequest" was null or undefined when calling createMuteKeyword().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/hidden/words`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: MuteKeywordRequestToJSON(requestParameters['muteKeywordRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateMuteKeywordResponseFromJSON(jsonValue));
    }

    /**
     */
    async createMuteKeyword(requestParameters: CreateMuteKeywordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateMuteKeywordResponse> {
        const response = await this.createMuteKeywordRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async deleteMuteKeywordRaw(requestParameters: DeleteMuteKeywordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['ids'] != null) {
            queryParameters['ids[]'] = requestParameters['ids'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/words`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteMuteKeyword(requestParameters: DeleteMuteKeywordRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteMuteKeywordRaw(requestParameters, initOverrides);
    }

    /**
     */
    async hideChatsRaw(requestParameters: HideChatsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const consumes: runtime.Consume[] = [
            { contentType: 'multipart/form-data' },
        ];
        // @ts-ignore: canConsumeForm may be unused
        const canConsumeForm = runtime.canConsumeForm(consumes);

        let formParams: { append(param: string, value: any): any };
        let useForm = false;
        if (useForm) {
            formParams = new FormData();
        } else {
            formParams = new URLSearchParams();
        }

        if (requestParameters['chatRoomId'] != null) {
            formParams.append('chat_room_id', requestParameters['chatRoomId'] as any);
        }

        const response = await this.request({
            path: `/v1/hidden/chats`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async hideChats(requestParameters: HideChatsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.hideChatsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async hideUsersRaw(requestParameters: HideUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const consumes: runtime.Consume[] = [
            { contentType: 'multipart/form-data' },
        ];
        // @ts-ignore: canConsumeForm may be unused
        const canConsumeForm = runtime.canConsumeForm(consumes);

        let formParams: { append(param: string, value: any): any };
        let useForm = false;
        if (useForm) {
            formParams = new FormData();
        } else {
            formParams = new URLSearchParams();
        }

        if (requestParameters['userId'] != null) {
            formParams.append('user_id', requestParameters['userId'] as any);
        }

        const response = await this.request({
            path: `/v1/hidden/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async hideUsers(requestParameters: HideUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.hideUsersRaw(requestParameters, initOverrides);
    }

    /**
     */
    async listHiddenChatsRaw(requestParameters: ListHiddenChatsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatRoomsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/chats`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatRoomsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listHiddenChats(requestParameters: ListHiddenChatsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatRoomsResponse> {
        const response = await this.listHiddenChatsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listHiddenUsersRaw(requestParameters: ListHiddenUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<HiddenResponse>> {
        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HiddenResponseFromJSON(jsonValue));
    }

    /**
     */
    async listHiddenUsers(requestParameters: ListHiddenUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<HiddenResponse> {
        const response = await this.listHiddenUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listMuteKeywordsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MuteKeywordResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/words`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => MuteKeywordResponseFromJSON(jsonValue));
    }

    /**
     */
    async listMuteKeywords(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MuteKeywordResponse> {
        const response = await this.listMuteKeywordsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async unhideChatsRaw(requestParameters: UnhideChatsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['chatRoomIds'] != null) {
            queryParameters['chat_room_ids'] = requestParameters['chatRoomIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/chats`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unhideChats(requestParameters: UnhideChatsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unhideChatsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unhideUsersRaw(requestParameters: UnhideUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['userIds'] != null) {
            queryParameters['user_ids[]'] = requestParameters['userIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/hidden/users`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unhideUsers(requestParameters: UnhideUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unhideUsersRaw(requestParameters, initOverrides);
    }

}