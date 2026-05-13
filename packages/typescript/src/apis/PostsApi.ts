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
  CreatePostResponse,
  GiftTransactionsResponse,
  LikePostsResponse,
  NoreplyMode,
  Post,
  PostLikersResponse,
  PostResponse,
  PostTagsResponse,
  PostType,
  PostsResponse,
  SharedUrl,
  ThreadInfo,
  ValidationPostResponse,
} from '../models/index';
import {
    CreatePostResponseFromJSON,
    CreatePostResponseToJSON,
    GiftTransactionsResponseFromJSON,
    GiftTransactionsResponseToJSON,
    LikePostsResponseFromJSON,
    LikePostsResponseToJSON,
    NoreplyModeFromJSON,
    NoreplyModeToJSON,
    PostFromJSON,
    PostToJSON,
    PostLikersResponseFromJSON,
    PostLikersResponseToJSON,
    PostResponseFromJSON,
    PostResponseToJSON,
    PostTagsResponseFromJSON,
    PostTagsResponseToJSON,
    PostTypeFromJSON,
    PostTypeToJSON,
    PostsResponseFromJSON,
    PostsResponseToJSON,
    SharedUrlFromJSON,
    SharedUrlToJSON,
    ThreadInfoFromJSON,
    ThreadInfoToJSON,
    ValidationPostResponseFromJSON,
    ValidationPostResponseToJSON,
} from '../models/index';

export interface CreateConferenceCallPostRequest {
    apiKey?: string;
    attachment2Filename?: string | null;
    attachment3Filename?: string | null;
    attachment4Filename?: string | null;
    attachment5Filename?: string | null;
    attachment6Filename?: string | null;
    attachment7Filename?: string | null;
    attachment8Filename?: string | null;
    attachment9Filename?: string | null;
    attachmentFilename?: string | null;
    callType?: string | null;
    categoryId?: number | null;
    color?: number | null;
    fontSize?: number | null;
    gameTitle?: string | null;
    groupId?: number | null;
    joinableBy?: string | null;
    language?: string | null;
    messageTags?: object | null;
    signedInfo?: string;
    text?: string | null;
    timestamp?: number;
    uuid?: string;
}

export interface CreatePostRequest {
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

export interface CreateSharePostRequest {
    apiKey?: string;
    color?: number | null;
    fontSize?: number | null;
    groupId?: number | null;
    language?: string | null;
    messageTags?: object | null;
    shareableId?: number;
    shareableType?: string;
    signedInfo?: string;
    text?: string | null;
    timestamp?: number;
    uuid?: string;
}

export interface DeletePostsRequest {
    postsIds?: Array<number>;
}

export interface GetActiveCallPostRequest {
    userId?: number;
}

export interface GetCallFollowersTimelineRequest {
    fromTimestamp?: number | null;
    number?: number | null;
    categoryId?: number | null;
    callType?: string | null;
    includeCircleCall?: boolean | null;
    excludeRecentGomimushi?: boolean | null;
}

export interface GetCallTimelineRequest {
    groupId?: number | null;
    fromTimestamp?: number | null;
    number?: number | null;
    categoryId?: number | null;
    callType?: string | null;
    includeCircleCall?: boolean | null;
    crossGeneration?: boolean | null;
    excludeRecentGomimushi?: boolean | null;
}

export interface GetFollowingTimelineRequest {
    from?: string | null;
    fromPostId?: number | null;
    onlyRoot?: boolean | null;
    orderBy?: string;
    number?: number | null;
    mxn?: number | null;
    reduceSelfie?: boolean | null;
    customGenerationRange?: boolean | null;
}

export interface GetGroupTimelineRequest {
    groupId?: number;
    fromPostId?: number | null;
    reverse?: boolean | null;
    postType?: PostType | null;
    number?: number | null;
    onlyRoot?: boolean | null;
}

export interface GetMyPostsRequest {
    fromPostId?: number | null;
    number?: number;
    includeGroupPost?: boolean | null;
}

export interface GetPostRequest {
    id: number;
    cacheControl?: string | null;
}

export interface GetPostGiftTransactionsRequest {
    postId: number;
    number?: number | null;
    from?: string | null;
}

export interface GetPostLikersRequest {
    id: number;
    fromId?: number | null;
}

export interface GetPostRepostsRequest {
    id: number;
    fromPostId?: number | null;
}

export interface GetPostUrlMetadataRequest {
    url?: string;
}

export interface GetPostsByIdsRequest {
    postIds?: Array<number>;
}

export interface GetPostsByTagRequest {
    tag: string;
    fromPostId?: number | null;
    number?: number | null;
}

export interface GetRecentEngagementPostsRequest {
    number?: number;
}

export interface GetRecommendedPostTagsRequest {
    saveRecentSearch?: boolean;
    tag?: string;
}

export interface GetRecommendedTimelineRequest {
    experimentNum?: number;
    variantNum?: number;
    number?: number;
}

export interface GetTimelineRequest {
    noreplyMode: NoreplyMode;
    orderBy?: string;
    experimentOlderAgeRules?: boolean | null;
    from?: string | null;
    fromPostId?: number | null;
    number?: number | null;
    mxn?: number | null;
    en?: number | null;
    vn?: number | null;
    reduceSelfie?: boolean | null;
    customGenerationRange?: boolean | null;
}

export interface GetUserTimelineRequest {
    userId?: number;
    fromPostId?: number | null;
    postType?: PostType | null;
    number?: number | null;
}

export interface LikePostsRequest {
    postIds?: Array<number>;
}

export interface MovePostToSpecificThreadRequest {
    id: number;
    threadId: number;
}

export interface MovePostToThreadRequest {
    id: number;
    threadIconFilename?: string | null;
    title?: string | null;
}

export interface PinGroupPostRequest {
    groupId?: number;
    postId?: number;
}

export interface ReportPostRequest {
    postId: number;
    categoryId?: number;
    opponentId?: number;
    reason?: string | null;
    screenshot2Filename?: string | null;
    screenshot3Filename?: string | null;
    screenshot4Filename?: string | null;
    screenshotFilename?: string | null;
}

export interface RepostRequest {
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
    postId?: number;
    postType?: PostType | null;
    sharedUrl?: object | null;
    text?: string | null;
    videoFileName?: string | null;
}

export interface SearchPostsRequest {
    keyword?: string;
    postOwnerScope?: string;
    onlyMedia?: boolean;
    fromPostId?: number | null;
    number?: number | null;
}

export interface UnlikePostRequest {
    id: number;
}

export interface UnpinGroupPostRequest {
    groupId?: number;
}

export interface UpdatePostRequest {
    id: number;
    apiKey?: string;
    color?: number | null;
    fontSize?: number | null;
    language?: string | null;
    messageTags?: object | null;
    signedInfo?: string;
    text?: string | null;
    timestamp?: number;
    uuid?: string;
}

export interface ValidatePostRequest {
    groupId?: number | null;
    text?: string;
    threadId?: number | null;
}

export interface ViewPostVideoRequest {
    videoId: number;
}

/**
 * 
 */
export class PostsApi extends runtime.BaseAPI {

    /**
     */
    async createConferenceCallPostRaw(requestParameters: CreateConferenceCallPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreatePostResponse>> {
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

        if (requestParameters['apiKey'] != null) {
            formParams.append('api_key', requestParameters['apiKey'] as any);
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

        if (requestParameters['callType'] != null) {
            formParams.append('call_type', requestParameters['callType'] as any);
        }

        if (requestParameters['categoryId'] != null) {
            formParams.append('category_id', requestParameters['categoryId'] as any);
        }

        if (requestParameters['color'] != null) {
            formParams.append('color', requestParameters['color'] as any);
        }

        if (requestParameters['fontSize'] != null) {
            formParams.append('font_size', requestParameters['fontSize'] as any);
        }

        if (requestParameters['gameTitle'] != null) {
            formParams.append('game_title', requestParameters['gameTitle'] as any);
        }

        if (requestParameters['groupId'] != null) {
            formParams.append('group_id', requestParameters['groupId'] as any);
        }

        if (requestParameters['joinableBy'] != null) {
            formParams.append('joinable_by', requestParameters['joinableBy'] as any);
        }

        if (requestParameters['language'] != null) {
            formParams.append('language', requestParameters['language'] as any);
        }

        if (requestParameters['messageTags'] != null) {
            formParams.append('message_tags', new Blob([JSON.stringify(objectToJSON(requestParameters['messageTags']))], { type: "application/json", }));
                    }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v2/posts/new_conference_call`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreatePostResponseFromJSON(jsonValue));
    }

    /**
     */
    async createConferenceCallPost(requestParameters: CreateConferenceCallPostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreatePostResponse> {
        const response = await this.createConferenceCallPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createPostRaw(requestParameters: CreatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Post>> {
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
            path: `/v3/posts/new`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostFromJSON(jsonValue));
    }

    /**
     */
    async createPost(requestParameters: CreatePostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Post> {
        const response = await this.createPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createSharePostRaw(requestParameters: CreateSharePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Post>> {
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

        if (requestParameters['apiKey'] != null) {
            formParams.append('api_key', requestParameters['apiKey'] as any);
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

        if (requestParameters['language'] != null) {
            formParams.append('language', requestParameters['language'] as any);
        }

        if (requestParameters['messageTags'] != null) {
            formParams.append('message_tags', new Blob([JSON.stringify(objectToJSON(requestParameters['messageTags']))], { type: "application/json", }));
                    }

        if (requestParameters['shareableId'] != null) {
            formParams.append('shareable_id', requestParameters['shareableId'] as any);
        }

        if (requestParameters['shareableType'] != null) {
            formParams.append('shareable_type', requestParameters['shareableType'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v2/posts/new_share_post`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostFromJSON(jsonValue));
    }

    /**
     */
    async createSharePost(requestParameters: CreateSharePostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Post> {
        const response = await this.createSharePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async deleteAllPostsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/delete_all_post`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteAllPosts(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteAllPostsRaw(initOverrides);
    }

    /**
     */
    async deletePostsRaw(requestParameters: DeletePostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['postsIds'] != null) {
            formParams.append('posts_ids[]', requestParameters['postsIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v2/posts/mass_destroy`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deletePosts(requestParameters: DeletePostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deletePostsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async getActiveCallPostRaw(requestParameters: GetActiveCallPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostResponse>> {
        const queryParameters: any = {};

        if (requestParameters['userId'] != null) {
            queryParameters['user_id'] = requestParameters['userId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/active_call`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostResponseFromJSON(jsonValue));
    }

    /**
     */
    async getActiveCallPost(requestParameters: GetActiveCallPostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostResponse> {
        const response = await this.getActiveCallPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getCallFollowersTimelineRaw(requestParameters: GetCallFollowersTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['categoryId'] != null) {
            queryParameters['category_id'] = requestParameters['categoryId'];
        }

        if (requestParameters['callType'] != null) {
            queryParameters['call_type'] = requestParameters['callType'];
        }

        if (requestParameters['includeCircleCall'] != null) {
            queryParameters['include_circle_call'] = requestParameters['includeCircleCall'];
        }

        if (requestParameters['excludeRecentGomimushi'] != null) {
            queryParameters['exclude_recent_gomimushi'] = requestParameters['excludeRecentGomimushi'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/call_followers_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getCallFollowersTimeline(requestParameters: GetCallFollowersTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getCallFollowersTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getCallTimelineRaw(requestParameters: GetCallTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['categoryId'] != null) {
            queryParameters['category_id'] = requestParameters['categoryId'];
        }

        if (requestParameters['callType'] != null) {
            queryParameters['call_type'] = requestParameters['callType'];
        }

        if (requestParameters['includeCircleCall'] != null) {
            queryParameters['include_circle_call'] = requestParameters['includeCircleCall'];
        }

        if (requestParameters['crossGeneration'] != null) {
            queryParameters['cross_generation'] = requestParameters['crossGeneration'];
        }

        if (requestParameters['excludeRecentGomimushi'] != null) {
            queryParameters['exclude_recent_gomimushi'] = requestParameters['excludeRecentGomimushi'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/call_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getCallTimeline(requestParameters: GetCallTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getCallTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getFollowingTimelineRaw(requestParameters: GetFollowingTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['onlyRoot'] != null) {
            queryParameters['only_root'] = requestParameters['onlyRoot'];
        }

        if (requestParameters['orderBy'] != null) {
            queryParameters['order_by'] = requestParameters['orderBy'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['mxn'] != null) {
            queryParameters['mxn'] = requestParameters['mxn'];
        }

        if (requestParameters['reduceSelfie'] != null) {
            queryParameters['reduce_selfie'] = requestParameters['reduceSelfie'];
        }

        if (requestParameters['customGenerationRange'] != null) {
            queryParameters['custom_generation_range'] = requestParameters['customGenerationRange'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/following_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFollowingTimeline(requestParameters: GetFollowingTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getFollowingTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupTimelineRaw(requestParameters: GetGroupTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['reverse'] != null) {
            queryParameters['reverse'] = requestParameters['reverse'];
        }

        if (requestParameters['postType'] != null) {
            queryParameters['post_type'] = requestParameters['postType'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['onlyRoot'] != null) {
            queryParameters['only_root'] = requestParameters['onlyRoot'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/group_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupTimeline(requestParameters: GetGroupTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getGroupTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getMyPostsRaw(requestParameters: GetMyPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['includeGroupPost'] != null) {
            queryParameters['include_group_post'] = requestParameters['includeGroupPost'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/mine`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getMyPosts(requestParameters: GetMyPostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getMyPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostRaw(requestParameters: GetPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters['cacheControl'] != null) {
            headerParameters['Cache-Control'] = String(requestParameters['cacheControl']);
        }

        const response = await this.request({
            path: `/v2/posts/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPost(requestParameters: GetPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostResponse> {
        const response = await this.getPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostGiftTransactionsRaw(requestParameters: GetPostGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftTransactionsResponse>> {
        if (requestParameters['postId'] == null) {
            throw new runtime.RequiredError(
                'postId',
                'Required parameter "postId" was null or undefined when calling getPostGiftTransactions().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/{post_id}/gift_transactions`.replace(`{${"post_id"}}`, encodeURIComponent(String(requestParameters['postId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftTransactionsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPostGiftTransactions(requestParameters: GetPostGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftTransactionsResponse> {
        const response = await this.getPostGiftTransactionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostLikersRaw(requestParameters: GetPostLikersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostLikersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getPostLikers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/{id}/likers`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostLikersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPostLikers(requestParameters: GetPostLikersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostLikersResponse> {
        const response = await this.getPostLikersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostRepostsRaw(requestParameters: GetPostRepostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getPostReposts().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/{id}/reposts`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPostReposts(requestParameters: GetPostRepostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getPostRepostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostUrlMetadataRaw(requestParameters: GetPostUrlMetadataRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SharedUrl>> {
        const queryParameters: any = {};

        if (requestParameters['url'] != null) {
            queryParameters['url'] = requestParameters['url'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/url_metadata`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SharedUrlFromJSON(jsonValue));
    }

    /**
     */
    async getPostUrlMetadata(requestParameters: GetPostUrlMetadataRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SharedUrl> {
        const response = await this.getPostUrlMetadataRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostsByIdsRaw(requestParameters: GetPostsByIdsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['postIds'] != null) {
            queryParameters['post_ids[]'] = requestParameters['postIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/multiple`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPostsByIds(requestParameters: GetPostsByIdsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getPostsByIdsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPostsByTagRaw(requestParameters: GetPostsByTagRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['tag'] == null) {
            throw new runtime.RequiredError(
                'tag',
                'Required parameter "tag" was null or undefined when calling getPostsByTag().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/tags/{tag}`.replace(`{${"tag"}}`, encodeURIComponent(String(requestParameters['tag']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPostsByTag(requestParameters: GetPostsByTagRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getPostsByTagRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRecentEngagementPostsRaw(requestParameters: GetRecentEngagementPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/recent_engagement`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRecentEngagementPosts(requestParameters: GetRecentEngagementPostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getRecentEngagementPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRecommendedPostTagsRaw(requestParameters: GetRecommendedPostTagsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostTagsResponse>> {
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

        if (requestParameters['saveRecentSearch'] != null) {
            formParams.append('save_recent_search', requestParameters['saveRecentSearch'] as any);
        }

        if (requestParameters['tag'] != null) {
            formParams.append('tag', requestParameters['tag'] as any);
        }

        const response = await this.request({
            path: `/v1/posts/recommended_tag`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostTagsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRecommendedPostTags(requestParameters: GetRecommendedPostTagsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostTagsResponse> {
        const response = await this.getRecommendedPostTagsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRecommendedTimelineRaw(requestParameters: GetRecommendedTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['experimentNum'] != null) {
            queryParameters['experiment_num'] = requestParameters['experimentNum'];
        }

        if (requestParameters['variantNum'] != null) {
            queryParameters['variant_num'] = requestParameters['variantNum'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/recommended_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRecommendedTimeline(requestParameters: GetRecommendedTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getRecommendedTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getTimelineRaw(requestParameters: GetTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['noreplyMode'] == null) {
            throw new runtime.RequiredError(
                'noreplyMode',
                'Required parameter "noreplyMode" was null or undefined when calling getTimeline().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['orderBy'] != null) {
            queryParameters['order_by'] = requestParameters['orderBy'];
        }

        if (requestParameters['experimentOlderAgeRules'] != null) {
            queryParameters['experiment_older_age_rules'] = requestParameters['experimentOlderAgeRules'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['mxn'] != null) {
            queryParameters['mxn'] = requestParameters['mxn'];
        }

        if (requestParameters['en'] != null) {
            queryParameters['en'] = requestParameters['en'];
        }

        if (requestParameters['vn'] != null) {
            queryParameters['vn'] = requestParameters['vn'];
        }

        if (requestParameters['reduceSelfie'] != null) {
            queryParameters['reduce_selfie'] = requestParameters['reduceSelfie'];
        }

        if (requestParameters['customGenerationRange'] != null) {
            queryParameters['custom_generation_range'] = requestParameters['customGenerationRange'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/{noreply_mode}timeline`.replace(`{${"noreply_mode"}}`, encodeURIComponent(String(requestParameters['noreplyMode']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getTimeline(requestParameters: GetTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserTimelineRaw(requestParameters: GetUserTimelineRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['userId'] != null) {
            queryParameters['user_id'] = requestParameters['userId'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['postType'] != null) {
            queryParameters['post_type'] = requestParameters['postType'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/user_timeline`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserTimeline(requestParameters: GetUserTimelineRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getUserTimelineRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async likePostsRaw(requestParameters: LikePostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LikePostsResponse>> {
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

        if (requestParameters['postIds'] != null) {
            formParams.append('post_ids[]', requestParameters['postIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v2/posts/like`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LikePostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async likePosts(requestParameters: LikePostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LikePostsResponse> {
        const response = await this.likePostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async movePostToSpecificThreadRaw(requestParameters: MovePostToSpecificThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ThreadInfo>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling movePostToSpecificThread().'
            );
        }

        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling movePostToSpecificThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/posts/{id}/move_to_thread/{thread_id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"thread_id"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ThreadInfoFromJSON(jsonValue));
    }

    /**
     */
    async movePostToSpecificThread(requestParameters: MovePostToSpecificThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ThreadInfo> {
        const response = await this.movePostToSpecificThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async movePostToThreadRaw(requestParameters: MovePostToThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ThreadInfo>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling movePostToThread().'
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
            path: `/v3/posts/{id}/move_to_thread`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ThreadInfoFromJSON(jsonValue));
    }

    /**
     */
    async movePostToThread(requestParameters: MovePostToThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ThreadInfo> {
        const response = await this.movePostToThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async pinGroupPostRaw(requestParameters: PinGroupPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['groupId'] != null) {
            formParams.append('group_id', requestParameters['groupId'] as any);
        }

        if (requestParameters['postId'] != null) {
            formParams.append('post_id', requestParameters['postId'] as any);
        }

        const response = await this.request({
            path: `/v2/posts/group_pinned_post`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async pinGroupPost(requestParameters: PinGroupPostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.pinGroupPostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async reportPostRaw(requestParameters: ReportPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['postId'] == null) {
            throw new runtime.RequiredError(
                'postId',
                'Required parameter "postId" was null or undefined when calling reportPost().'
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

        if (requestParameters['categoryId'] != null) {
            formParams.append('category_id', requestParameters['categoryId'] as any);
        }

        if (requestParameters['opponentId'] != null) {
            formParams.append('opponent_id', requestParameters['opponentId'] as any);
        }

        if (requestParameters['reason'] != null) {
            formParams.append('reason', requestParameters['reason'] as any);
        }

        if (requestParameters['screenshot2Filename'] != null) {
            formParams.append('screenshot_2_filename', requestParameters['screenshot2Filename'] as any);
        }

        if (requestParameters['screenshot3Filename'] != null) {
            formParams.append('screenshot_3_filename', requestParameters['screenshot3Filename'] as any);
        }

        if (requestParameters['screenshot4Filename'] != null) {
            formParams.append('screenshot_4_filename', requestParameters['screenshot4Filename'] as any);
        }

        if (requestParameters['screenshotFilename'] != null) {
            formParams.append('screenshot_filename', requestParameters['screenshotFilename'] as any);
        }

        const response = await this.request({
            path: `/v3/posts/{post_id}/report`.replace(`{${"post_id"}}`, encodeURIComponent(String(requestParameters['postId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async reportPost(requestParameters: ReportPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.reportPostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async repostRaw(requestParameters: RepostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreatePostResponse>> {
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

        if (requestParameters['postId'] != null) {
            formParams.append('post_id', requestParameters['postId'] as any);
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
            path: `/v3/posts/repost`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreatePostResponseFromJSON(jsonValue));
    }

    /**
     */
    async repost(requestParameters: RepostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreatePostResponse> {
        const response = await this.repostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async searchPostsRaw(requestParameters: SearchPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['postOwnerScope'] != null) {
            queryParameters['post_owner_scope'] = requestParameters['postOwnerScope'];
        }

        if (requestParameters['onlyMedia'] != null) {
            queryParameters['only_media'] = requestParameters['onlyMedia'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/search`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async searchPosts(requestParameters: SearchPostsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.searchPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async unlikePostRaw(requestParameters: UnlikePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unlikePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/{id}/unlike`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unlikePost(requestParameters: UnlikePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unlikePostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unpinGroupPostRaw(requestParameters: UnpinGroupPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/posts/group_pinned_post`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unpinGroupPost(requestParameters: UnpinGroupPostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unpinGroupPostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updatePostRaw(requestParameters: UpdatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Post>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updatePost().'
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

        if (requestParameters['apiKey'] != null) {
            formParams.append('api_key', requestParameters['apiKey'] as any);
        }

        if (requestParameters['color'] != null) {
            formParams.append('color', requestParameters['color'] as any);
        }

        if (requestParameters['fontSize'] != null) {
            formParams.append('font_size', requestParameters['fontSize'] as any);
        }

        if (requestParameters['language'] != null) {
            formParams.append('language', requestParameters['language'] as any);
        }

        if (requestParameters['messageTags'] != null) {
            formParams.append('message_tags', new Blob([JSON.stringify(objectToJSON(requestParameters['messageTags']))], { type: "application/json", }));
                    }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/posts/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostFromJSON(jsonValue));
    }

    /**
     */
    async updatePost(requestParameters: UpdatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Post> {
        const response = await this.updatePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async validatePostRaw(requestParameters: ValidatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ValidationPostResponse>> {
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

        if (requestParameters['groupId'] != null) {
            formParams.append('group_id', requestParameters['groupId'] as any);
        }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['threadId'] != null) {
            formParams.append('thread_id', requestParameters['threadId'] as any);
        }

        const response = await this.request({
            path: `/v1/posts/validate`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ValidationPostResponseFromJSON(jsonValue));
    }

    /**
     */
    async validatePost(requestParameters: ValidatePostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ValidationPostResponse> {
        const response = await this.validatePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async viewPostVideoRaw(requestParameters: ViewPostVideoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['videoId'] == null) {
            throw new runtime.RequiredError(
                'videoId',
                'Required parameter "videoId" was null or undefined when calling viewPostVideo().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/posts/videos/{videoId}/view`.replace(`{${"videoId"}}`, encodeURIComponent(String(requestParameters['videoId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async viewPostVideo(requestParameters: ViewPostVideoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.viewPostVideoRaw(requestParameters, initOverrides);
    }

}
