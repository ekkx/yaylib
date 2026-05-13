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
  CreateGroupResponse,
  CreateQuotaResponse,
  GiftSendersResponse,
  GiftTransactionsResponse,
  GroupCategoriesResponse,
  GroupGiftHistoryResponse,
  GroupResponse,
  GroupUserResponse,
  GroupUsersResponse,
  GroupsRelatedResponse,
  GroupsResponse,
  PostsResponse,
  UnreadStatusResponse,
  UsersByTimestampResponse,
  UsersResponse,
} from '../models/index';
import {
    CreateGroupResponseFromJSON,
    CreateGroupResponseToJSON,
    CreateQuotaResponseFromJSON,
    CreateQuotaResponseToJSON,
    GiftSendersResponseFromJSON,
    GiftSendersResponseToJSON,
    GiftTransactionsResponseFromJSON,
    GiftTransactionsResponseToJSON,
    GroupCategoriesResponseFromJSON,
    GroupCategoriesResponseToJSON,
    GroupGiftHistoryResponseFromJSON,
    GroupGiftHistoryResponseToJSON,
    GroupResponseFromJSON,
    GroupResponseToJSON,
    GroupUserResponseFromJSON,
    GroupUserResponseToJSON,
    GroupUsersResponseFromJSON,
    GroupUsersResponseToJSON,
    GroupsRelatedResponseFromJSON,
    GroupsRelatedResponseToJSON,
    GroupsResponseFromJSON,
    GroupsResponseToJSON,
    PostsResponseFromJSON,
    PostsResponseToJSON,
    UnreadStatusResponseFromJSON,
    UnreadStatusResponseToJSON,
    UsersByTimestampResponseFromJSON,
    UsersByTimestampResponseToJSON,
    UsersResponseFromJSON,
    UsersResponseToJSON,
} from '../models/index';

export interface AcceptGroupJoinRequestRequest {
    id: number;
    userId: number;
}

export interface AcceptGroupTransferRequest {
    id: number;
}

export interface BanGroupUserRequest {
    id: number;
    userId: number;
}

export interface CancelGroupTransferRequest {
    id: number;
}

export interface CreateGroupRequest {
    allowMembersToPostImageAndVideo?: boolean | null;
    allowMembersToPostUrl?: boolean | null;
    allowOwnershipTransfer?: boolean | null;
    allowThreadCreationBy?: string | null;
    apiKey?: string;
    callTimelineDisplay?: boolean | null;
    coverImageFilename?: string | null;
    description?: string | null;
    gender?: number | null;
    generationGroupsLimit?: number | null;
    groupCategoryId?: number | null;
    groupIconFilename?: string | null;
    guidelines?: string | null;
    hideConferenceCall?: boolean | null;
    hideFromGameEight?: boolean;
    hideReportedPosts?: boolean | null;
    isPrivate?: boolean | null;
    onlyMobileVerified?: boolean | null;
    onlyVerifiedAge?: boolean | null;
    secret?: boolean | null;
    signedInfo?: string;
    subCategoryId?: string | null;
    timestamp?: number;
    topic?: string;
    uuid?: string;
}

export interface DeclineGroupJoinRequestRequest {
    id: number;
    userId: number;
}

export interface DeputizeGroupUsersRequest {
    id: number;
}

export interface DeputizeGroupUsersMassRequest {
    groupId: number;
    apiKey?: string;
    signedInfo?: string;
    timestamp?: number;
    userIds?: Array<number>;
    uuid?: string;
}

export interface FireGroupUserRequest {
    groupId: number;
    userId: number;
}

export interface GetGroupRequest {
    id: number;
}

export interface GetGroupBanListRequest {
    id: number;
    page?: number;
}

export interface GetGroupGiftHistoryRequest {
    groupId: number;
    number?: number | null;
    from?: string | null;
}

export interface GetGroupGiftTransactionsRequest {
    groupId: number;
    number?: number | null;
    from?: string | null;
}

export interface GetGroupHighlightsRequest {
    groupId: number;
    from?: number | null;
    number?: number | null;
}

export interface GetGroupMemberRequest {
    id: number;
    userId: number;
}

export interface GetGroupMembersRequest {
    id: number;
    number?: number | null;
    fromId?: number | null;
}

export interface GetGroupReceivedGiftSendersRequest {
    groupId: number;
    giftId: number;
    userId?: number;
    timestamp?: number;
    number?: number | null;
}

export interface GetGroupUnreadStatusRequest {
    fromTime?: number | null;
}

export interface GetInvitableGroupUsersRequest {
    groupId: number;
    fromTimestamp?: number | null;
    userNickname?: string | null;
}

export interface GetJoinedGroupStatusesRequest {
    ids?: Array<number>;
}

export interface GetRelatableGroupsRequest {
    id: number;
    keyword?: string | null;
    from?: string | null;
}

export interface GetRelatedGroupsRequest {
    id: number;
    keyword?: string | null;
    from?: string | null;
}

export interface GetUserGroupListRequest {
    page?: number;
    userId?: number;
}

export interface InviteToGroupRequest {
    id: number;
    userIds?: Array<number>;
}

export interface JoinGroupRequest {
    id: number;
}

export interface LeaveGroupRequest {
    id: number;
}

export interface ListGroupCategoriesRequest {
    page?: number;
    number?: number | null;
}

export interface ListGroupsRequest {
    groupCategoryId?: number | null;
    keyword?: string | null;
    fromTimestamp?: number | null;
    subCategoryId?: number | null;
}

export interface ListMyGroupsRequest {
    fromTimestamp?: number | null;
}

export interface PinGroupHighlightPostRequest {
    groupId: number;
    postId: number;
}

export interface RemoveGroupCoverRequest {
    id: number;
}

export interface RemoveGroupDeputiesRequest {
    id: number;
}

export interface RemoveGroupIconRequest {
    id: number;
}

export interface RemoveRelatedGroupsRequest {
    id: number;
    relatedGroupId?: Array<number>;
}

export interface ReportGroupRequest {
    groupId: number;
    categoryId?: number;
    opponentId?: number;
    reason?: string | null;
    screenshot2Filename?: string | null;
    screenshot3Filename?: string | null;
    screenshot4Filename?: string | null;
    screenshotFilename?: string | null;
}

export interface RequestGroupWalkthroughRequest {
    id: number;
}

export interface SearchGroupPostsRequest {
    id: number;
    keyword?: string;
    fromPostId?: number | null;
    number?: number | null;
    onlyThreadPosts?: boolean;
}

export interface SetGroupTitleRequest {
    id: number;
    title?: string;
}

export interface TakeOverGroupRequest {
    id: number;
}

export interface TransferGroupRequest {
    id: number;
    apiKey?: string;
    signedInfo?: string;
    timestamp?: number;
    userId?: number;
    uuid?: string;
}

export interface UnbanGroupUserRequest {
    id: number;
    userId: number;
}

export interface UnpinGroupHighlightPostRequest {
    groupId: number;
    postId: number;
}

export interface UpdateGroupRequest {
    id: number;
    allowMembersToPostImageAndVideo?: boolean | null;
    allowMembersToPostUrl?: boolean | null;
    allowOwnershipTransfer?: boolean | null;
    allowThreadCreationBy?: string | null;
    apiKey?: string;
    callTimelineDisplay?: boolean | null;
    coverImageFilename?: string | null;
    description?: string | null;
    gender?: number | null;
    generationGroupsLimit?: number | null;
    groupCategoryId?: number | null;
    groupIconFilename?: string | null;
    guidelines?: string | null;
    hideConferenceCall?: boolean | null;
    hideFromGameEight?: boolean | null;
    hideReportedPosts?: boolean | null;
    isPrivate?: boolean | null;
    onlyMobileVerified?: boolean | null;
    onlyVerifiedAge?: boolean | null;
    secret?: boolean | null;
    signedInfo?: string;
    subCategoryId?: string | null;
    timestamp?: number;
    topic?: string | null;
    uuid?: string;
}

export interface UpdateRelatedGroupsRequest {
    id: number;
    relatedGroupId?: Array<number>;
}

export interface VisitGroupRequest {
    id: number;
}

export interface WithdrawGroupDeputyRequest {
    groupId: number;
    userId: number;
}

export interface WithdrawGroupTransferRequest {
    id: number;
    userId?: number;
}

/**
 * 
 */
export class GroupsApi extends runtime.BaseAPI {

    /**
     */
    async acceptGroupJoinRequestRaw(requestParameters: AcceptGroupJoinRequestRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling acceptGroupJoinRequest().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling acceptGroupJoinRequest().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/accept/{userId}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async acceptGroupJoinRequest(requestParameters: AcceptGroupJoinRequestRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.acceptGroupJoinRequestRaw(requestParameters, initOverrides);
    }

    /**
     */
    async acceptGroupTransferRaw(requestParameters: AcceptGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling acceptGroupTransfer().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/transfer`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async acceptGroupTransfer(requestParameters: AcceptGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.acceptGroupTransferRaw(requestParameters, initOverrides);
    }

    /**
     */
    async banGroupUserRaw(requestParameters: BanGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling banGroupUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling banGroupUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/ban/{userId}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async banGroupUser(requestParameters: BanGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.banGroupUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async cancelGroupTransferRaw(requestParameters: CancelGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling cancelGroupTransfer().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/transfer`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async cancelGroupTransfer(requestParameters: CancelGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.cancelGroupTransferRaw(requestParameters, initOverrides);
    }

    /**
     */
    async createGroupRaw(requestParameters: CreateGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateGroupResponse>> {
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

        if (requestParameters['allowMembersToPostImageAndVideo'] != null) {
            formParams.append('allow_members_to_post_image_and_video', requestParameters['allowMembersToPostImageAndVideo'] as any);
        }

        if (requestParameters['allowMembersToPostUrl'] != null) {
            formParams.append('allow_members_to_post_url', requestParameters['allowMembersToPostUrl'] as any);
        }

        if (requestParameters['allowOwnershipTransfer'] != null) {
            formParams.append('allow_ownership_transfer', requestParameters['allowOwnershipTransfer'] as any);
        }

        if (requestParameters['allowThreadCreationBy'] != null) {
            formParams.append('allow_thread_creation_by', requestParameters['allowThreadCreationBy'] as any);
        }

        if (requestParameters['apiKey'] != null) {
            formParams.append('api_key', requestParameters['apiKey'] as any);
        }

        if (requestParameters['callTimelineDisplay'] != null) {
            formParams.append('call_timeline_display', requestParameters['callTimelineDisplay'] as any);
        }

        if (requestParameters['coverImageFilename'] != null) {
            formParams.append('cover_image_filename', requestParameters['coverImageFilename'] as any);
        }

        if (requestParameters['description'] != null) {
            formParams.append('description', requestParameters['description'] as any);
        }

        if (requestParameters['gender'] != null) {
            formParams.append('gender', requestParameters['gender'] as any);
        }

        if (requestParameters['generationGroupsLimit'] != null) {
            formParams.append('generation_groups_limit', requestParameters['generationGroupsLimit'] as any);
        }

        if (requestParameters['groupCategoryId'] != null) {
            formParams.append('group_category_id', requestParameters['groupCategoryId'] as any);
        }

        if (requestParameters['groupIconFilename'] != null) {
            formParams.append('group_icon_filename', requestParameters['groupIconFilename'] as any);
        }

        if (requestParameters['guidelines'] != null) {
            formParams.append('guidelines', requestParameters['guidelines'] as any);
        }

        if (requestParameters['hideConferenceCall'] != null) {
            formParams.append('hide_conference_call', requestParameters['hideConferenceCall'] as any);
        }

        if (requestParameters['hideFromGameEight'] != null) {
            formParams.append('hide_from_game_eight', requestParameters['hideFromGameEight'] as any);
        }

        if (requestParameters['hideReportedPosts'] != null) {
            formParams.append('hide_reported_posts', requestParameters['hideReportedPosts'] as any);
        }

        if (requestParameters['isPrivate'] != null) {
            formParams.append('is_private', requestParameters['isPrivate'] as any);
        }

        if (requestParameters['onlyMobileVerified'] != null) {
            formParams.append('only_mobile_verified', requestParameters['onlyMobileVerified'] as any);
        }

        if (requestParameters['onlyVerifiedAge'] != null) {
            formParams.append('only_verified_age', requestParameters['onlyVerifiedAge'] as any);
        }

        if (requestParameters['secret'] != null) {
            formParams.append('secret', requestParameters['secret'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['subCategoryId'] != null) {
            formParams.append('sub_category_id', requestParameters['subCategoryId'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['topic'] != null) {
            formParams.append('topic', requestParameters['topic'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/groups/new`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateGroupResponseFromJSON(jsonValue));
    }

    /**
     */
    async createGroup(requestParameters: CreateGroupRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateGroupResponse> {
        const response = await this.createGroupRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async declineGroupJoinRequestRaw(requestParameters: DeclineGroupJoinRequestRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling declineGroupJoinRequest().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling declineGroupJoinRequest().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/decline/{userId}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async declineGroupJoinRequest(requestParameters: DeclineGroupJoinRequestRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.declineGroupJoinRequestRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deputizeGroupUsersRaw(requestParameters: DeputizeGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling deputizeGroupUsers().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/deputize`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deputizeGroupUsers(requestParameters: DeputizeGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deputizeGroupUsersRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deputizeGroupUsersMassRaw(requestParameters: DeputizeGroupUsersMassRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling deputizeGroupUsersMass().'
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

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['userIds'] != null) {
            formParams.append('user_ids[]', requestParameters['userIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/groups/{group_id}/deputize/mass`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deputizeGroupUsersMass(requestParameters: DeputizeGroupUsersMassRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deputizeGroupUsersMassRaw(requestParameters, initOverrides);
    }

    /**
     */
    async fireGroupUserRaw(requestParameters: FireGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling fireGroupUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling fireGroupUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/fire/{user_id}`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))).replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async fireGroupUser(requestParameters: FireGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.fireGroupUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async getGroupRaw(requestParameters: GetGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getGroup().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroup(requestParameters: GetGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupResponse> {
        const response = await this.getGroupRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupBanListRaw(requestParameters: GetGroupBanListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getGroupBanList().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['page'] != null) {
            queryParameters['page'] = requestParameters['page'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/ban_list`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupBanList(requestParameters: GetGroupBanListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersResponse> {
        const response = await this.getGroupBanListRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupCreateQuotaRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateQuotaResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/created_quota`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateQuotaResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupCreateQuota(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateQuotaResponse> {
        const response = await this.getGroupCreateQuotaRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupGiftHistoryRaw(requestParameters: GetGroupGiftHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupGiftHistoryResponse>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling getGroupGiftHistory().'
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
            path: `/v1/groups/{group_id}/gift_history`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupGiftHistoryResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupGiftHistory(requestParameters: GetGroupGiftHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupGiftHistoryResponse> {
        const response = await this.getGroupGiftHistoryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupGiftTransactionsRaw(requestParameters: GetGroupGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftTransactionsResponse>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling getGroupGiftTransactions().'
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
            path: `/v1/groups/{group_id}/gift_transactions`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftTransactionsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupGiftTransactions(requestParameters: GetGroupGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftTransactionsResponse> {
        const response = await this.getGroupGiftTransactionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupHighlightsRaw(requestParameters: GetGroupHighlightsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling getGroupHighlights().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/highlights`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupHighlights(requestParameters: GetGroupHighlightsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getGroupHighlightsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupMemberRaw(requestParameters: GetGroupMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupUserResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getGroupMember().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling getGroupMember().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/members/{userId}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupUserResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupMember(requestParameters: GetGroupMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupUserResponse> {
        const response = await this.getGroupMemberRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupMembersRaw(requestParameters: GetGroupMembersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupUsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getGroupMembers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/groups/{id}/members`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupMembers(requestParameters: GetGroupMembersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupUsersResponse> {
        const response = await this.getGroupMembersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupReceivedGiftSendersRaw(requestParameters: GetGroupReceivedGiftSendersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftSendersResponse>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling getGroupReceivedGiftSenders().'
            );
        }

        if (requestParameters['giftId'] == null) {
            throw new runtime.RequiredError(
                'giftId',
                'Required parameter "giftId" was null or undefined when calling getGroupReceivedGiftSenders().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['userId'] != null) {
            queryParameters['user_id'] = requestParameters['userId'];
        }

        if (requestParameters['timestamp'] != null) {
            queryParameters['timestamp'] = requestParameters['timestamp'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/received_gifts/{gift_id}/senders`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))).replace(`{${"gift_id"}}`, encodeURIComponent(String(requestParameters['giftId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftSendersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupReceivedGiftSenders(requestParameters: GetGroupReceivedGiftSendersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftSendersResponse> {
        const response = await this.getGroupReceivedGiftSendersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGroupUnreadStatusRaw(requestParameters: GetGroupUnreadStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UnreadStatusResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTime'] != null) {
            queryParameters['from_time'] = requestParameters['fromTime'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/unread_status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UnreadStatusResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupUnreadStatus(requestParameters: GetGroupUnreadStatusRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UnreadStatusResponse> {
        const response = await this.getGroupUnreadStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getInvitableGroupUsersRaw(requestParameters: GetInvitableGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersByTimestampResponse>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling getInvitableGroupUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['userNickname'] != null) {
            queryParameters['user[nickname]'] = requestParameters['userNickname'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/users/invitable`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersByTimestampResponseFromJSON(jsonValue));
    }

    /**
     */
    async getInvitableGroupUsers(requestParameters: GetInvitableGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersByTimestampResponse> {
        const response = await this.getInvitableGroupUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getJoinedGroupStatusesRaw(requestParameters: GetJoinedGroupStatusesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: string; }>> {
        const queryParameters: any = {};

        if (requestParameters['ids'] != null) {
            queryParameters['ids[]'] = requestParameters['ids'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/joined_statuses`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async getJoinedGroupStatuses(requestParameters: GetJoinedGroupStatusesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: string; }> {
        const response = await this.getJoinedGroupStatusesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRelatableGroupsRaw(requestParameters: GetRelatableGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupsRelatedResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getRelatableGroups().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/relatable`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupsRelatedResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRelatableGroups(requestParameters: GetRelatableGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupsRelatedResponse> {
        const response = await this.getRelatableGroupsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRelatedGroupsRaw(requestParameters: GetRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupsRelatedResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getRelatedGroups().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/related`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupsRelatedResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRelatedGroups(requestParameters: GetRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupsRelatedResponse> {
        const response = await this.getRelatedGroupsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserGroupListRaw(requestParameters: GetUserGroupListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['page'] != null) {
            queryParameters['page'] = requestParameters['page'];
        }

        if (requestParameters['userId'] != null) {
            queryParameters['user_id'] = requestParameters['userId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/user_group_list`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserGroupList(requestParameters: GetUserGroupListRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupsResponse> {
        const response = await this.getUserGroupListRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async inviteToGroupRaw(requestParameters: InviteToGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling inviteToGroup().'
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

        if (requestParameters['userIds'] != null) {
            formParams.append('user_ids[]', requestParameters['userIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v1/groups/{id}/invite`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async inviteToGroup(requestParameters: InviteToGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.inviteToGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async joinGroupRaw(requestParameters: JoinGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling joinGroup().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/join`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async joinGroup(requestParameters: JoinGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.joinGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async leaveGroupRaw(requestParameters: LeaveGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling leaveGroup().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/leave`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async leaveGroup(requestParameters: LeaveGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.leaveGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async listGroupCategoriesRaw(requestParameters: ListGroupCategoriesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupCategoriesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['page'] != null) {
            queryParameters['page'] = requestParameters['page'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/categories`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupCategoriesResponseFromJSON(jsonValue));
    }

    /**
     */
    async listGroupCategories(requestParameters: ListGroupCategoriesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupCategoriesResponse> {
        const response = await this.listGroupCategoriesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listGroupsRaw(requestParameters: ListGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['groupCategoryId'] != null) {
            queryParameters['group_category_id'] = requestParameters['groupCategoryId'];
        }

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['subCategoryId'] != null) {
            queryParameters['sub_category_id'] = requestParameters['subCategoryId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/groups`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listGroups(requestParameters: ListGroupsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupsResponse> {
        const response = await this.listGroupsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listMyGroupsRaw(requestParameters: ListMyGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/groups/mine`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listMyGroups(requestParameters: ListMyGroupsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupsResponse> {
        const response = await this.listMyGroupsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async pinGroupHighlightPostRaw(requestParameters: PinGroupHighlightPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling pinGroupHighlightPost().'
            );
        }

        if (requestParameters['postId'] == null) {
            throw new runtime.RequiredError(
                'postId',
                'Required parameter "postId" was null or undefined when calling pinGroupHighlightPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/highlights/{post_id}`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))).replace(`{${"post_id"}}`, encodeURIComponent(String(requestParameters['postId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async pinGroupHighlightPost(requestParameters: PinGroupHighlightPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.pinGroupHighlightPostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async removeGroupCoverRaw(requestParameters: RemoveGroupCoverRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeGroupCover().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/groups/{id}/cover`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeGroupCover(requestParameters: RemoveGroupCoverRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeGroupCoverRaw(requestParameters, initOverrides);
    }

    /**
     */
    async removeGroupDeputiesRaw(requestParameters: RemoveGroupDeputiesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeGroupDeputies().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/deputize`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeGroupDeputies(requestParameters: RemoveGroupDeputiesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeGroupDeputiesRaw(requestParameters, initOverrides);
    }

    /**
     */
    async removeGroupIconRaw(requestParameters: RemoveGroupIconRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeGroupIcon().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/groups/{id}/icon`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeGroupIcon(requestParameters: RemoveGroupIconRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeGroupIconRaw(requestParameters, initOverrides);
    }

    /**
     */
    async removeRelatedGroupsRaw(requestParameters: RemoveRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeRelatedGroups().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['relatedGroupId'] != null) {
            queryParameters['related_group_id[]'] = requestParameters['relatedGroupId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/related`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeRelatedGroups(requestParameters: RemoveRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeRelatedGroupsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async reportGroupRaw(requestParameters: ReportGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling reportGroup().'
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
            path: `/v3/groups/{group_id}/report`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async reportGroup(requestParameters: ReportGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.reportGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async requestGroupWalkthroughRaw(requestParameters: RequestGroupWalkthroughRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling requestGroupWalkthrough().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/request_walkthrough`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async requestGroupWalkthrough(requestParameters: RequestGroupWalkthroughRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.requestGroupWalkthroughRaw(requestParameters, initOverrides);
    }

    /**
     */
    async searchGroupPostsRaw(requestParameters: SearchGroupPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling searchGroupPosts().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['fromPostId'] != null) {
            queryParameters['from_post_id'] = requestParameters['fromPostId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['onlyThreadPosts'] != null) {
            queryParameters['only_thread_posts'] = requestParameters['onlyThreadPosts'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/groups/{id}/posts/search`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async searchGroupPosts(requestParameters: SearchGroupPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.searchGroupPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async setGroupTitleRaw(requestParameters: SetGroupTitleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling setGroupTitle().'
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

        if (requestParameters['title'] != null) {
            formParams.append('title', requestParameters['title'] as any);
        }

        const response = await this.request({
            path: `/v1/groups/{id}/set_title`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async setGroupTitle(requestParameters: SetGroupTitleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.setGroupTitleRaw(requestParameters, initOverrides);
    }

    /**
     */
    async takeOverGroupRaw(requestParameters: TakeOverGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling takeOverGroup().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/take_over`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async takeOverGroup(requestParameters: TakeOverGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.takeOverGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async transferGroupRaw(requestParameters: TransferGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling transferGroup().'
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

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['userId'] != null) {
            formParams.append('user_id', requestParameters['userId'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/groups/{id}/transfer`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async transferGroup(requestParameters: TransferGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.transferGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unbanGroupUserRaw(requestParameters: UnbanGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unbanGroupUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling unbanGroupUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/unban/{userId}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unbanGroupUser(requestParameters: UnbanGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unbanGroupUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unpinGroupHighlightPostRaw(requestParameters: UnpinGroupHighlightPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling unpinGroupHighlightPost().'
            );
        }

        if (requestParameters['postId'] == null) {
            throw new runtime.RequiredError(
                'postId',
                'Required parameter "postId" was null or undefined when calling unpinGroupHighlightPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/highlights/{post_id}`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))).replace(`{${"post_id"}}`, encodeURIComponent(String(requestParameters['postId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unpinGroupHighlightPost(requestParameters: UnpinGroupHighlightPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unpinGroupHighlightPostRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateGroupRaw(requestParameters: UpdateGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateGroup().'
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

        if (requestParameters['allowMembersToPostImageAndVideo'] != null) {
            formParams.append('allow_members_to_post_image_and_video', requestParameters['allowMembersToPostImageAndVideo'] as any);
        }

        if (requestParameters['allowMembersToPostUrl'] != null) {
            formParams.append('allow_members_to_post_url', requestParameters['allowMembersToPostUrl'] as any);
        }

        if (requestParameters['allowOwnershipTransfer'] != null) {
            formParams.append('allow_ownership_transfer', requestParameters['allowOwnershipTransfer'] as any);
        }

        if (requestParameters['allowThreadCreationBy'] != null) {
            formParams.append('allow_thread_creation_by', requestParameters['allowThreadCreationBy'] as any);
        }

        if (requestParameters['apiKey'] != null) {
            formParams.append('api_key', requestParameters['apiKey'] as any);
        }

        if (requestParameters['callTimelineDisplay'] != null) {
            formParams.append('call_timeline_display', requestParameters['callTimelineDisplay'] as any);
        }

        if (requestParameters['coverImageFilename'] != null) {
            formParams.append('cover_image_filename', requestParameters['coverImageFilename'] as any);
        }

        if (requestParameters['description'] != null) {
            formParams.append('description', requestParameters['description'] as any);
        }

        if (requestParameters['gender'] != null) {
            formParams.append('gender', requestParameters['gender'] as any);
        }

        if (requestParameters['generationGroupsLimit'] != null) {
            formParams.append('generation_groups_limit', requestParameters['generationGroupsLimit'] as any);
        }

        if (requestParameters['groupCategoryId'] != null) {
            formParams.append('group_category_id', requestParameters['groupCategoryId'] as any);
        }

        if (requestParameters['groupIconFilename'] != null) {
            formParams.append('group_icon_filename', requestParameters['groupIconFilename'] as any);
        }

        if (requestParameters['guidelines'] != null) {
            formParams.append('guidelines', requestParameters['guidelines'] as any);
        }

        if (requestParameters['hideConferenceCall'] != null) {
            formParams.append('hide_conference_call', requestParameters['hideConferenceCall'] as any);
        }

        if (requestParameters['hideFromGameEight'] != null) {
            formParams.append('hide_from_game_eight', requestParameters['hideFromGameEight'] as any);
        }

        if (requestParameters['hideReportedPosts'] != null) {
            formParams.append('hide_reported_posts', requestParameters['hideReportedPosts'] as any);
        }

        if (requestParameters['isPrivate'] != null) {
            formParams.append('is_private', requestParameters['isPrivate'] as any);
        }

        if (requestParameters['onlyMobileVerified'] != null) {
            formParams.append('only_mobile_verified', requestParameters['onlyMobileVerified'] as any);
        }

        if (requestParameters['onlyVerifiedAge'] != null) {
            formParams.append('only_verified_age', requestParameters['onlyVerifiedAge'] as any);
        }

        if (requestParameters['secret'] != null) {
            formParams.append('secret', requestParameters['secret'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['subCategoryId'] != null) {
            formParams.append('sub_category_id', requestParameters['subCategoryId'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['topic'] != null) {
            formParams.append('topic', requestParameters['topic'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/groups/{id}/update`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupResponseFromJSON(jsonValue));
    }

    /**
     */
    async updateGroup(requestParameters: UpdateGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupResponse> {
        const response = await this.updateGroupRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async updateRelatedGroupsRaw(requestParameters: UpdateRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateRelatedGroups().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['relatedGroupId'] != null) {
            queryParameters['related_group_id[]'] = requestParameters['relatedGroupId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/related`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateRelatedGroups(requestParameters: UpdateRelatedGroupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateRelatedGroupsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async visitGroupRaw(requestParameters: VisitGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling visitGroup().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{id}/visit`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async visitGroup(requestParameters: VisitGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.visitGroupRaw(requestParameters, initOverrides);
    }

    /**
     */
    async withdrawGroupDeputyRaw(requestParameters: WithdrawGroupDeputyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['groupId'] == null) {
            throw new runtime.RequiredError(
                'groupId',
                'Required parameter "groupId" was null or undefined when calling withdrawGroupDeputy().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling withdrawGroupDeputy().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/groups/{group_id}/deputize/{user_id}/withdraw`.replace(`{${"group_id"}}`, encodeURIComponent(String(requestParameters['groupId']))).replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async withdrawGroupDeputy(requestParameters: WithdrawGroupDeputyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.withdrawGroupDeputyRaw(requestParameters, initOverrides);
    }

    /**
     */
    async withdrawGroupTransferRaw(requestParameters: WithdrawGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling withdrawGroupTransfer().'
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

        if (requestParameters['userId'] != null) {
            formParams.append('user_id', requestParameters['userId'] as any);
        }

        const response = await this.request({
            path: `/v1/groups/{id}/transfer/withdraw`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async withdrawGroupTransfer(requestParameters: WithdrawGroupTransferRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.withdrawGroupTransferRaw(requestParameters, initOverrides);
    }

}
