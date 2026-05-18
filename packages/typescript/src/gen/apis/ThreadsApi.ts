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


import * as runtime from '../runtime.js';
import type {
  CreateGroupThreadRequest,
  GroupThreadListResponse,
  Post,
  PostType,
  PostsResponse,
  ThreadInfo,
} from '../models/index.js';
import {
    CreateGroupThreadRequestFromJSON,
    CreateGroupThreadRequestToJSON,
    GroupThreadListResponseFromJSON,
    GroupThreadListResponseToJSON,
    PostFromJSON,
    PostToJSON,
    PostTypeFromJSON,
    PostTypeToJSON,
    PostsResponseFromJSON,
    PostsResponseToJSON,
    ThreadInfoFromJSON,
    ThreadInfoToJSON,
} from '../models/index.js';

export interface AddThreadMemberRequest {
    threadId: number;
    id: number;
}

export interface CreateThreadRequest {
    createGroupThreadRequest: CreateGroupThreadRequest;
}

export interface CreateThreadPostRequest {
    id: number;
    xJwt?: string;
    attachment2Filename?: string | null;
    attachment3Filename?: string | null;
    attachment4Filename?: string | null;
    attachment5Filename?: string | null;
    attachment6Filename?: string | null;
    attachment7Filename?: string | null;
    attachment8Filename?: string | null;
    attachment9Filename?: string | null;
    attachmentFilename?: string | null;
    choices?: Array<string> | null;
    color?: number | null;
    fontSize?: number | null;
    groupId?: number | null;
    inReplyTo?: number | null;
    language?: string | null;
    mentionIds?: Array<number> | null;
    messageTags?: object | null;
    postType?: PostType | null;
    sharedUrl?: object | null;
    text?: string | null;
    videoFileName?: string | null;
}

export interface DeleteThreadRequest {
    id: number;
}

export interface GetJoinedThreadStatusesRequest {
    ids?: Array<number>;
}

export interface GetThreadRequest {
    id: number;
}

export interface GetThreadPostsRequest {
    id: number;
    postType?: PostType | null;
    number?: number | null;
    from?: number | null;
}

export interface ListThreadsRequest {
    groupId?: number;
    from?: string | null;
    joinStatus?: string | null;
}

export interface RemoveThreadMemberRequest {
    threadId: number;
    id: number;
}

export interface UpdateThreadRequest {
    id: number;
    threadIconFilename?: string | null;
    title?: string;
}

/**
 * 
 */
export class ThreadsApi extends runtime.BaseAPI {

    /**
     */
    async addThreadMemberRaw(requestParameters: AddThreadMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling addThreadMember().'
            );
        }

        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling addThreadMember().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/{thread_id}/members/{id}`.replace(`{${"thread_id"}}`, encodeURIComponent(String(requestParameters['threadId']))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async addThreadMember(requestParameters: AddThreadMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.addThreadMemberRaw(requestParameters, initOverrides);
    }

    /**
     */
    async createThreadRaw(requestParameters: CreateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ThreadInfo>> {
        if (requestParameters['createGroupThreadRequest'] == null) {
            throw new runtime.RequiredError(
                'createGroupThreadRequest',
                'Required parameter "createGroupThreadRequest" was null or undefined when calling createThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/threads`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateGroupThreadRequestToJSON(requestParameters['createGroupThreadRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ThreadInfoFromJSON(jsonValue));
    }

    /**
     */
    async createThread(requestParameters: CreateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ThreadInfo> {
        const response = await this.createThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createThreadPostRaw(requestParameters: CreateThreadPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Post>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling createThreadPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters['xJwt'] != null) {
            headerParameters['X-Jwt'] = String(requestParameters['xJwt']);
        }

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

        if (requestParameters['attachment2Filename'] != null) {
            formParams.append('attachment_2_filename', requestParameters['attachment2Filename'] as any);
        }

        if (requestParameters['attachment3Filename'] != null) {
            formParams.append('attachment_3_filename', requestParameters['attachment3Filename'] as any);
        }

        if (requestParameters['attachment4Filename'] != null) {
            formParams.append('attachment_4_filename', requestParameters['attachment4Filename'] as any);
        }

        if (requestParameters['attachment5Filename'] != null) {
            formParams.append('attachment_5_filename', requestParameters['attachment5Filename'] as any);
        }

        if (requestParameters['attachment6Filename'] != null) {
            formParams.append('attachment_6_filename', requestParameters['attachment6Filename'] as any);
        }

        if (requestParameters['attachment7Filename'] != null) {
            formParams.append('attachment_7_filename', requestParameters['attachment7Filename'] as any);
        }

        if (requestParameters['attachment8Filename'] != null) {
            formParams.append('attachment_8_filename', requestParameters['attachment8Filename'] as any);
        }

        if (requestParameters['attachment9Filename'] != null) {
            formParams.append('attachment_9_filename', requestParameters['attachment9Filename'] as any);
        }

        if (requestParameters['attachmentFilename'] != null) {
            formParams.append('attachment_filename', requestParameters['attachmentFilename'] as any);
        }

        if (requestParameters['choices'] != null) {
            formParams.append('choices[]', requestParameters['choices']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        if (requestParameters['color'] != null) {
            formParams.append('color', requestParameters['color'] as any);
        }

        if (requestParameters['fontSize'] != null) {
            formParams.append('font_size', requestParameters['fontSize'] as any);
        }

        if (requestParameters['groupId'] != null) {
            formParams.append('group_id', requestParameters['groupId'] as any);
        }

        if (requestParameters['inReplyTo'] != null) {
            formParams.append('in_reply_to', requestParameters['inReplyTo'] as any);
        }

        if (requestParameters['language'] != null) {
            formParams.append('language', requestParameters['language'] as any);
        }

        if (requestParameters['mentionIds'] != null) {
            formParams.append('mention_ids[]', requestParameters['mentionIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        if (requestParameters['messageTags'] != null) {
            formParams.append('message_tags', new Blob([JSON.stringify(objectToJSON(requestParameters['messageTags']))], { type: "application/json", }));
                    }

        if (requestParameters['postType'] != null) {
            formParams.append('post_type', requestParameters['postType'] as any);
        }

        if (requestParameters['sharedUrl'] != null) {
            formParams.append('shared_url', new Blob([JSON.stringify(objectToJSON(requestParameters['sharedUrl']))], { type: "application/json", }));
                    }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['videoFileName'] != null) {
            formParams.append('video_file_name', requestParameters['videoFileName'] as any);
        }

        const response = await this.request({
            path: `/v1/threads/{id}/posts`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostFromJSON(jsonValue));
    }

    /**
     */
    async createThreadPost(requestParameters: CreateThreadPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Post> {
        const response = await this.createThreadPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async deleteThreadRaw(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling deleteThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteThread(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteThreadRaw(requestParameters, initOverrides);
    }

    /**
     */
    async getJoinedThreadStatusesRaw(requestParameters: GetJoinedThreadStatusesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: string; }>> {
        const queryParameters: any = {};

        if (requestParameters['ids'] != null) {
            queryParameters['ids[]'] = requestParameters['ids'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/joined_statuses`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async getJoinedThreadStatuses(requestParameters: GetJoinedThreadStatusesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: string; }> {
        const response = await this.getJoinedThreadStatusesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getThreadRaw(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ThreadInfo>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ThreadInfoFromJSON(jsonValue));
    }

    /**
     */
    async getThread(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ThreadInfo> {
        const response = await this.getThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getThreadPostsRaw(requestParameters: GetThreadPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getThreadPosts().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['postType'] != null) {
            queryParameters['post_type'] = requestParameters['postType'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/{id}/posts`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getThreadPosts(requestParameters: GetThreadPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getThreadPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listThreadsRaw(requestParameters: ListThreadsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupThreadListResponse>> {
        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['joinStatus'] != null) {
            queryParameters['join_status'] = requestParameters['joinStatus'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupThreadListResponseFromJSON(jsonValue));
    }

    /**
     */
    async listThreads(requestParameters: ListThreadsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupThreadListResponse> {
        const response = await this.listThreadsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async removeThreadMemberRaw(requestParameters: RemoveThreadMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling removeThreadMember().'
            );
        }

        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeThreadMember().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/threads/{thread_id}/members/{id}`.replace(`{${"thread_id"}}`, encodeURIComponent(String(requestParameters['threadId']))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeThreadMember(requestParameters: RemoveThreadMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeThreadMemberRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateThreadRaw(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateThread().'
            );
        }

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

        if (requestParameters['threadIconFilename'] != null) {
            formParams.append('thread_icon_filename', requestParameters['threadIconFilename'] as any);
        }

        if (requestParameters['title'] != null) {
            formParams.append('title', requestParameters['title'] as any);
        }

        const response = await this.request({
            path: `/v1/threads/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateThread(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateThreadRaw(requestParameters, initOverrides);
    }

}