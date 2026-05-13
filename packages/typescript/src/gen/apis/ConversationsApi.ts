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
  PostsResponse,
} from '../models/index';
import {
    PostsResponseFromJSON,
    PostsResponseToJSON,
} from '../models/index';

export interface GetConversationRequest {
    id: number;
    groupId?: number | null;
    threadId?: number | null;
    number?: number | null;
    fromPostId?: number | null;
    reverse?: boolean;
}

export interface GetRootPostsRequest {
    ids?: Array<number>;
}

/**
 * 
 */
export class ConversationsApi extends runtime.BaseAPI {

    /**
     */
    async getConversationRaw(requestParameters: GetConversationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getConversation().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        if (requestParameters['threadId'] != null) {
            queryParameters['thread_id'] = requestParameters['threadId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['reverse'] != null) {
            queryParameters['reverse'] = requestParameters['reverse'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/conversations/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getConversation(requestParameters: GetConversationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getConversationRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRootPostsRaw(requestParameters: GetRootPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['ids'] != null) {
            queryParameters['ids[]'] = requestParameters['ids'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/conversations/root_posts`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRootPosts(requestParameters: GetRootPostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getRootPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}