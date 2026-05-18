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
  ActiveFollowingsResponse,
  AdditionalSettingsResponse,
  BlockedUserIdsResponse,
  BlockedUsersResponse,
  BookmarkPostResponse,
  CommonIdsRequest,
  DefaultSettingsResponse,
  FollowRequestCountResponse,
  FollowUsersResponse,
  FootprintsResponse,
  GiftTransactionsResponse,
  HimaUsersResponse,
  LoginEmailUserRequest,
  LoginUpdateResponse,
  LoginUserResponse,
  PolicyAgreementsResponse,
  PostsResponse,
  PresignedUrlResponse,
  RefreshCounterRequestsResponse,
  ReviewsResponse,
  SearchUsersRequest,
  TwoFAStatusResponse,
  TwoStepAuthEnabledResponse,
  TwoStepAuthRequestInfoResponse,
  UserCustomDefinitionsResponse,
  UserInterestsResponse,
  UserResponse,
  UserTimestampResponse,
  UsersByTimestampResponse,
  UsersResponse,
  WebSocketTokenResponse,
} from '../models/index.js';
import {
    ActiveFollowingsResponseFromJSON,
    ActiveFollowingsResponseToJSON,
    AdditionalSettingsResponseFromJSON,
    AdditionalSettingsResponseToJSON,
    BlockedUserIdsResponseFromJSON,
    BlockedUserIdsResponseToJSON,
    BlockedUsersResponseFromJSON,
    BlockedUsersResponseToJSON,
    BookmarkPostResponseFromJSON,
    BookmarkPostResponseToJSON,
    CommonIdsRequestFromJSON,
    CommonIdsRequestToJSON,
    DefaultSettingsResponseFromJSON,
    DefaultSettingsResponseToJSON,
    FollowRequestCountResponseFromJSON,
    FollowRequestCountResponseToJSON,
    FollowUsersResponseFromJSON,
    FollowUsersResponseToJSON,
    FootprintsResponseFromJSON,
    FootprintsResponseToJSON,
    GiftTransactionsResponseFromJSON,
    GiftTransactionsResponseToJSON,
    HimaUsersResponseFromJSON,
    HimaUsersResponseToJSON,
    LoginEmailUserRequestFromJSON,
    LoginEmailUserRequestToJSON,
    LoginUpdateResponseFromJSON,
    LoginUpdateResponseToJSON,
    LoginUserResponseFromJSON,
    LoginUserResponseToJSON,
    PolicyAgreementsResponseFromJSON,
    PolicyAgreementsResponseToJSON,
    PostsResponseFromJSON,
    PostsResponseToJSON,
    PresignedUrlResponseFromJSON,
    PresignedUrlResponseToJSON,
    RefreshCounterRequestsResponseFromJSON,
    RefreshCounterRequestsResponseToJSON,
    ReviewsResponseFromJSON,
    ReviewsResponseToJSON,
    SearchUsersRequestFromJSON,
    SearchUsersRequestToJSON,
    TwoFAStatusResponseFromJSON,
    TwoFAStatusResponseToJSON,
    TwoStepAuthEnabledResponseFromJSON,
    TwoStepAuthEnabledResponseToJSON,
    TwoStepAuthRequestInfoResponseFromJSON,
    TwoStepAuthRequestInfoResponseToJSON,
    UserCustomDefinitionsResponseFromJSON,
    UserCustomDefinitionsResponseToJSON,
    UserInterestsResponseFromJSON,
    UserInterestsResponseToJSON,
    UserResponseFromJSON,
    UserResponseToJSON,
    UserTimestampResponseFromJSON,
    UserTimestampResponseToJSON,
    UsersByTimestampResponseFromJSON,
    UsersByTimestampResponseToJSON,
    UsersResponseFromJSON,
    UsersResponseToJSON,
    WebSocketTokenResponseFromJSON,
    WebSocketTokenResponseToJSON,
} from '../models/index.js';

export interface AgreePolicyRequest {
    type: string;
}

export interface BlockUserRequest {
    id: number;
}

export interface ChangeEmailRequest {
    apiKey?: string;
    email?: string;
    emailGrantToken?: string | null;
    password?: string;
}

export interface ChangePasswordRequest {
    apiKey?: string;
    currentPassword?: string;
    password?: string;
}

export interface CreateBookmarkRequest {
    userId: number;
    id: number;
}

export interface CreateReviewRequest {
    apiKey?: string;
    comment?: string;
    signedInfo?: string;
    timestamp?: number;
    userIds?: Array<number>;
    uuid?: string;
}

export interface DeleteBookmarkRequest {
    userId: number;
    id: number;
}

export interface DeleteFootprintRequest {
    userId: number;
    footprintId: number;
}

export interface DeleteMyReviewsRequest {
    reviewIds?: Array<number>;
}

export interface DisableTwoFactorAuthRequest {
    code?: string;
}

export interface EditUserRequest {
    apiKey?: string;
    biography?: string | null;
    countryCode?: string | null;
    coverImageFilename?: string | null;
    gender?: number | null;
    nickname?: string;
    prefecture?: string | null;
    profileIconFilename?: string | null;
    signedInfo?: string;
    timestamp?: number;
    username?: string | null;
    uuid?: string;
}

export interface EditUserV2Request {
    apiKey?: string;
    isPrivate?: boolean;
    nickname?: string;
    signedInfo?: string;
    timestamp?: number;
    uuid?: string;
}

export interface EnableTwoFactorAuthRequest {
    code?: string;
    type?: string;
}

export interface FollowUserRequest {
    id: number;
}

export interface FollowUsersRequest {
    userIds?: Array<number>;
}

export interface GetActiveFollowingsRequest {
    onlyOnline?: boolean;
    fromLoggedinAt?: number | null;
}

export interface GetBlockedUsersRequest {
    searchUsersRequest: SearchUsersRequest;
    fromId?: number | null;
}

export interface GetBookmarkedPostsRequest {
    userId: number;
    from?: string | null;
}

export interface GetChatableFollowingsRequest {
    searchUsersRequest: SearchUsersRequest;
    fromFollowId?: number | null;
    fromTimestamp?: number | null;
    orderBy?: string | null;
}

export interface GetFollowRequestsRequest {
    fromTimestamp?: number | null;
}

export interface GetFollowingsBornTodayRequest {
    birthdate?: number | null;
}

export interface GetFootprintsRequest {
    mode?: string | null;
    number?: number;
    from?: string | null;
}

export interface GetFreshUserRequest {
    id: number;
}

export interface GetHimaUsersRequest {
    fromHimaId?: number | null;
    number?: number | null;
}

export interface GetMyReviewsRequest {
    fromId?: number | null;
}

export interface GetRecommendedFollowUsersRequest {
    id: number;
    number?: number | null;
    page?: number | null;
}

export interface GetUserRequest {
    id: number;
}

export interface GetUserByQrRequest {
    qr: string;
}

export interface GetUserFollowersRequest {
    id: number;
    from?: number | null;
    followedByMe?: boolean | null;
    userNickname?: string | null;
}

export interface GetUserFollowingsRequest {
    id: number;
    from?: number | null;
    fromTimestamp?: number | null;
    orderBy?: string | null;
    userNickname?: string | null;
}

export interface GetUserGiftTransactionsRequest {
    userId: number;
    number?: number | null;
    from?: string | null;
}

export interface GetUserInfoRequest {
    id: number;
}

export interface GetUserPresignedUrlRequest {
    videoFileName?: string;
}

export interface GetUserReviewsRequest {
    id: number;
    fromId?: number | null;
}

export interface GetUsersByIdsRequest {
    xJwt?: string;
    userIds?: Array<number>;
}

export interface LoginWithEmailRequest {
    loginEmailUserRequest: LoginEmailUserRequest;
}

export interface LogoutRequest {
    uuid?: string;
}

export interface ReplyToReviewRequest {
    id: number;
    apiKey?: string;
    comment?: string;
    context?: string | null;
    signedInfo?: string;
    timestamp?: number;
    uuid?: string;
}

export interface ReportUserRequest {
    userId: number;
    categoryId?: number;
    reason?: string | null;
    screenshot2Filename?: string | null;
    screenshot3Filename?: string | null;
    screenshot4Filename?: string | null;
    screenshotFilename?: string | null;
}

export interface RequestFollowRequest {
    targetId: number;
    action?: string;
}

export interface ResetCountersRequest {
    counter?: string;
}

export interface ResetPasswordRequest {
    email?: string;
    emailGrantToken?: string;
    password?: string;
}

export interface SearchUsersRequest {
    gender?: number | null;
    nickname?: string | null;
    title?: string | null;
    biography?: string | null;
    fromTimestamp?: number | null;
    similarAge?: boolean | null;
    notRecentGomimushi?: boolean | null;
    recentlyCreated?: boolean | null;
    samePrefecture?: boolean | null;
    saveRecentSearch?: boolean | null;
}

export interface UnblockUserRequest {
    id: number;
}

export interface UnfollowUserRequest {
    id: number;
}

export interface UpdateAdditionalNotificationSettingRequest {
    mode?: string;
    on?: number;
}

export interface UpdateLanguageRequest {
    apiKey?: string;
    language?: string;
    signedInfo?: string;
    timestamp?: number;
    uuid?: string;
}

export interface UpdateLoginRequest {
    apiKey?: string;
    currentPassword?: string | null;
    email?: string;
    emailGrantToken?: string | null;
    password?: string | null;
}

export interface UpdateUserInterestsRequest {
    commonIdsRequest: CommonIdsRequest;
}

/**
 * 
 */
export class UsersApi extends runtime.BaseAPI {

    /**
     */
    async agreePolicyRaw(requestParameters: AgreePolicyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['type'] == null) {
            throw new runtime.RequiredError(
                'type',
                'Required parameter "type" was null or undefined when calling agreePolicy().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/policy_agreements/{type}`.replace(`{${"type"}}`, encodeURIComponent(String(requestParameters['type']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async agreePolicy(requestParameters: AgreePolicyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.agreePolicyRaw(requestParameters, initOverrides);
    }

    /**
     */
    async blockUserRaw(requestParameters: BlockUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling blockUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/{id}/block`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async blockUser(requestParameters: BlockUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.blockUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async changeEmailRaw(requestParameters: ChangeEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LoginUpdateResponse>> {
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

        if (requestParameters['email'] != null) {
            formParams.append('email', requestParameters['email'] as any);
        }

        if (requestParameters['emailGrantToken'] != null) {
            formParams.append('email_grant_token', requestParameters['emailGrantToken'] as any);
        }

        if (requestParameters['password'] != null) {
            formParams.append('password', requestParameters['password'] as any);
        }

        const response = await this.request({
            path: `/v1/users/change_email`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LoginUpdateResponseFromJSON(jsonValue));
    }

    /**
     */
    async changeEmail(requestParameters: ChangeEmailRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginUpdateResponse> {
        const response = await this.changeEmailRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async changePasswordRaw(requestParameters: ChangePasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LoginUpdateResponse>> {
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

        if (requestParameters['currentPassword'] != null) {
            formParams.append('current_password', requestParameters['currentPassword'] as any);
        }

        if (requestParameters['password'] != null) {
            formParams.append('password', requestParameters['password'] as any);
        }

        const response = await this.request({
            path: `/v1/users/change_password`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LoginUpdateResponseFromJSON(jsonValue));
    }

    /**
     */
    async changePassword(requestParameters: ChangePasswordRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginUpdateResponse> {
        const response = await this.changePasswordRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createBookmarkRaw(requestParameters: CreateBookmarkRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BookmarkPostResponse>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling createBookmark().'
            );
        }

        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling createBookmark().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/{user_id}/bookmarks/{id}`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BookmarkPostResponseFromJSON(jsonValue));
    }

    /**
     */
    async createBookmark(requestParameters: CreateBookmarkRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BookmarkPostResponse> {
        const response = await this.createBookmarkRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createReviewRaw(requestParameters: CreateReviewRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['comment'] != null) {
            formParams.append('comment', requestParameters['comment'] as any);
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
            path: `/v1/users/reviews`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async createReview(requestParameters: CreateReviewRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.createReviewRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deleteBookmarkRaw(requestParameters: DeleteBookmarkRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling deleteBookmark().'
            );
        }

        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling deleteBookmark().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/{user_id}/bookmarks/{id}`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteBookmark(requestParameters: DeleteBookmarkRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteBookmarkRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deleteFootprintRaw(requestParameters: DeleteFootprintRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling deleteFootprint().'
            );
        }

        if (requestParameters['footprintId'] == null) {
            throw new runtime.RequiredError(
                'footprintId',
                'Required parameter "footprintId" was null or undefined when calling deleteFootprint().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/{user_id}/footprints/{footprint_id}`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))).replace(`{${"footprint_id"}}`, encodeURIComponent(String(requestParameters['footprintId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteFootprint(requestParameters: DeleteFootprintRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteFootprintRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deleteMyReviewsRaw(requestParameters: DeleteMyReviewsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['reviewIds'] != null) {
            queryParameters['review_ids[]'] = requestParameters['reviewIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/reviews`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteMyReviews(requestParameters: DeleteMyReviewsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteMyReviewsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async disableTwoFactorAuthRaw(requestParameters: DisableTwoFactorAuthRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['code'] != null) {
            formParams.append('code', requestParameters['code'] as any);
        }

        const response = await this.request({
            path: `/v1/users/secret/disable`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async disableTwoFactorAuth(requestParameters: DisableTwoFactorAuthRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.disableTwoFactorAuthRaw(requestParameters, initOverrides);
    }

    /**
     */
    async editUserRaw(requestParameters: EditUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['biography'] != null) {
            formParams.append('biography', requestParameters['biography'] as any);
        }

        if (requestParameters['countryCode'] != null) {
            formParams.append('country_code', requestParameters['countryCode'] as any);
        }

        if (requestParameters['coverImageFilename'] != null) {
            formParams.append('cover_image_filename', requestParameters['coverImageFilename'] as any);
        }

        if (requestParameters['gender'] != null) {
            formParams.append('gender', requestParameters['gender'] as any);
        }

        if (requestParameters['nickname'] != null) {
            formParams.append('nickname', requestParameters['nickname'] as any);
        }

        if (requestParameters['prefecture'] != null) {
            formParams.append('prefecture', requestParameters['prefecture'] as any);
        }

        if (requestParameters['profileIconFilename'] != null) {
            formParams.append('profile_icon_filename', requestParameters['profileIconFilename'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['username'] != null) {
            formParams.append('username', requestParameters['username'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/users/edit`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async editUser(requestParameters: EditUserRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.editUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async editUserV2Raw(requestParameters: EditUserV2Request, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['isPrivate'] != null) {
            formParams.append('is_private', requestParameters['isPrivate'] as any);
        }

        if (requestParameters['nickname'] != null) {
            formParams.append('nickname', requestParameters['nickname'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v2/users/edit`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async editUserV2(requestParameters: EditUserV2Request = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.editUserV2Raw(requestParameters, initOverrides);
    }

    /**
     */
    async enableTwoFactorAuthRaw(requestParameters: EnableTwoFactorAuthRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<TwoStepAuthEnabledResponse>> {
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

        if (requestParameters['code'] != null) {
            formParams.append('code', requestParameters['code'] as any);
        }

        if (requestParameters['type'] != null) {
            formParams.append('type', requestParameters['type'] as any);
        }

        const response = await this.request({
            path: `/v1/users/secret/enable`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TwoStepAuthEnabledResponseFromJSON(jsonValue));
    }

    /**
     */
    async enableTwoFactorAuth(requestParameters: EnableTwoFactorAuthRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<TwoStepAuthEnabledResponse> {
        const response = await this.enableTwoFactorAuthRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async followUserRaw(requestParameters: FollowUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling followUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/{id}/follow`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async followUser(requestParameters: FollowUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.followUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async followUsersRaw(requestParameters: FollowUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['userIds'] != null) {
            queryParameters['user_ids[]'] = requestParameters['userIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/follow`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async followUsers(requestParameters: FollowUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.followUsersRaw(requestParameters, initOverrides);
    }

    /**
     */
    async getActiveFollowingsRaw(requestParameters: GetActiveFollowingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ActiveFollowingsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['onlyOnline'] != null) {
            queryParameters['only_online'] = requestParameters['onlyOnline'];
        }

        if (requestParameters['fromLoggedinAt'] != null) {
            queryParameters['from_loggedin_at'] = requestParameters['fromLoggedinAt'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/active_followings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ActiveFollowingsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getActiveFollowings(requestParameters: GetActiveFollowingsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ActiveFollowingsResponse> {
        const response = await this.getActiveFollowingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getAdditionalNotificationSettingRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AdditionalSettingsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/additonal_notification_setting`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AdditionalSettingsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getAdditionalNotificationSetting(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AdditionalSettingsResponse> {
        const response = await this.getAdditionalNotificationSettingRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getBlockedUserIdsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BlockedUserIdsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/block_ids`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BlockedUserIdsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getBlockedUserIds(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BlockedUserIdsResponse> {
        const response = await this.getBlockedUserIdsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getBlockedUsersRaw(requestParameters: GetBlockedUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BlockedUsersResponse>> {
        if (requestParameters['searchUsersRequest'] == null) {
            throw new runtime.RequiredError(
                'searchUsersRequest',
                'Required parameter "searchUsersRequest" was null or undefined when calling getBlockedUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v2/users/blocked`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SearchUsersRequestToJSON(requestParameters['searchUsersRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BlockedUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getBlockedUsers(requestParameters: GetBlockedUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BlockedUsersResponse> {
        const response = await this.getBlockedUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getBookmarkedPostsRaw(requestParameters: GetBookmarkedPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostsResponse>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling getBookmarkedPosts().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/{user_id}/bookmarks`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getBookmarkedPosts(requestParameters: GetBookmarkedPostsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostsResponse> {
        const response = await this.getBookmarkedPostsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getChatableFollowingsRaw(requestParameters: GetChatableFollowingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FollowUsersResponse>> {
        if (requestParameters['searchUsersRequest'] == null) {
            throw new runtime.RequiredError(
                'searchUsersRequest',
                'Required parameter "searchUsersRequest" was null or undefined when calling getChatableFollowings().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromFollowId'] != null) {
            queryParameters['from_follow_id'] = requestParameters['fromFollowId'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['orderBy'] != null) {
            queryParameters['order_by'] = requestParameters['orderBy'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/users/followings/chatable`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SearchUsersRequestToJSON(requestParameters['searchUsersRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FollowUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatableFollowings(requestParameters: GetChatableFollowingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FollowUsersResponse> {
        const response = await this.getChatableFollowingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getDefaultSettingsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DefaultSettingsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/default_settings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DefaultSettingsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getDefaultSettings(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DefaultSettingsResponse> {
        const response = await this.getDefaultSettingsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getFollowRequestCountRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FollowRequestCountResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/follow_requests_count`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FollowRequestCountResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFollowRequestCount(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FollowRequestCountResponse> {
        const response = await this.getFollowRequestCountRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getFollowRequestsRaw(requestParameters: GetFollowRequestsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersByTimestampResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/follow_requests`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersByTimestampResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFollowRequests(requestParameters: GetFollowRequestsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersByTimestampResponse> {
        const response = await this.getFollowRequestsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getFollowingsBornTodayRaw(requestParameters: GetFollowingsBornTodayRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersResponse>> {
        const queryParameters: any = {};

        if (requestParameters['birthdate'] != null) {
            queryParameters['birthdate'] = requestParameters['birthdate'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/following_born_today`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFollowingsBornToday(requestParameters: GetFollowingsBornTodayRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersResponse> {
        const response = await this.getFollowingsBornTodayRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getFootprintsRaw(requestParameters: GetFootprintsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FootprintsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['mode'] != null) {
            queryParameters['mode'] = requestParameters['mode'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/users/footprints`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FootprintsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFootprints(requestParameters: GetFootprintsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FootprintsResponse> {
        const response = await this.getFootprintsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getFreshUserRaw(requestParameters: GetFreshUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getFreshUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/fresh/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserResponseFromJSON(jsonValue));
    }

    /**
     */
    async getFreshUser(requestParameters: GetFreshUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserResponse> {
        const response = await this.getFreshUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getHimaUsersRaw(requestParameters: GetHimaUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<HimaUsersResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromHimaId'] != null) {
            queryParameters['from_hima_id'] = requestParameters['fromHimaId'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/hima_users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HimaUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getHimaUsers(requestParameters: GetHimaUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<HimaUsersResponse> {
        const response = await this.getHimaUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getMyReviewsRaw(requestParameters: GetMyReviewsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReviewsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/reviews/mine`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReviewsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getMyReviews(requestParameters: GetMyReviewsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReviewsResponse> {
        const response = await this.getMyReviewsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPolicyAgreementsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PolicyAgreementsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/policy_agreements`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PolicyAgreementsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPolicyAgreements(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PolicyAgreementsResponse> {
        const response = await this.getPolicyAgreementsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getRecommendedFollowUsersRaw(requestParameters: GetRecommendedFollowUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getRecommendedFollowUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['page'] != null) {
            queryParameters['page'] = requestParameters['page'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/{id}/follow_recommended`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRecommendedFollowUsers(requestParameters: GetRecommendedFollowUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersResponse> {
        const response = await this.getRecommendedFollowUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getResetCountersRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RefreshCounterRequestsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/reset_counters`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RefreshCounterRequestsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getResetCounters(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RefreshCounterRequestsResponse> {
        const response = await this.getResetCountersRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getTwoFactorAuthRequestInfoRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<TwoStepAuthRequestInfoResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/secret`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TwoStepAuthRequestInfoResponseFromJSON(jsonValue));
    }

    /**
     */
    async getTwoFactorAuthRequestInfo(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<TwoStepAuthRequestInfoResponse> {
        const response = await this.getTwoFactorAuthRequestInfoRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getTwoFactorAuthStatusRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<TwoFAStatusResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/secret/status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TwoFAStatusResponseFromJSON(jsonValue));
    }

    /**
     */
    async getTwoFactorAuthStatus(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<TwoFAStatusResponse> {
        const response = await this.getTwoFactorAuthStatusRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserRaw(requestParameters: GetUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUser(requestParameters: GetUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserResponse> {
        const response = await this.getUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserByQrRaw(requestParameters: GetUserByQrRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserResponse>> {
        if (requestParameters['qr'] == null) {
            throw new runtime.RequiredError(
                'qr',
                'Required parameter "qr" was null or undefined when calling getUserByQr().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/qr_codes/{qr}`.replace(`{${"qr"}}`, encodeURIComponent(String(requestParameters['qr']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserByQr(requestParameters: GetUserByQrRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserResponse> {
        const response = await this.getUserByQrRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserCustomDefinitionsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserCustomDefinitionsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/custom_definitions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserCustomDefinitionsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserCustomDefinitions(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserCustomDefinitionsResponse> {
        const response = await this.getUserCustomDefinitionsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserFollowersRaw(requestParameters: GetUserFollowersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FollowUsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getUserFollowers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['followedByMe'] != null) {
            queryParameters['followed_by_me'] = requestParameters['followedByMe'];
        }

        if (requestParameters['userNickname'] != null) {
            queryParameters['user[nickname]'] = requestParameters['userNickname'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/users/{id}/followers`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FollowUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserFollowers(requestParameters: GetUserFollowersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FollowUsersResponse> {
        const response = await this.getUserFollowersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserFollowingsRaw(requestParameters: GetUserFollowingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FollowUsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getUserFollowings().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['orderBy'] != null) {
            queryParameters['order_by'] = requestParameters['orderBy'];
        }

        if (requestParameters['userNickname'] != null) {
            queryParameters['user[nickname]'] = requestParameters['userNickname'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/users/{id}/followings`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FollowUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserFollowings(requestParameters: GetUserFollowingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FollowUsersResponse> {
        const response = await this.getUserFollowingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserGiftTransactionsRaw(requestParameters: GetUserGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftTransactionsResponse>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling getUserGiftTransactions().'
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
            path: `/v1/users/{user_id}/gift_transactions`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftTransactionsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserGiftTransactions(requestParameters: GetUserGiftTransactionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftTransactionsResponse> {
        const response = await this.getUserGiftTransactionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserInfoRaw(requestParameters: GetUserInfoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getUserInfo().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/info/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserInfo(requestParameters: GetUserInfoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserResponse> {
        const response = await this.getUserInfoRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserInterestsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserInterestsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/interests`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserInterestsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserInterests(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserInterestsResponse> {
        const response = await this.getUserInterestsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserPresignedUrlRaw(requestParameters: GetUserPresignedUrlRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PresignedUrlResponse>> {
        const queryParameters: any = {};

        if (requestParameters['videoFileName'] != null) {
            queryParameters['video_file_name'] = requestParameters['videoFileName'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/presigned_url`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PresignedUrlResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserPresignedUrl(requestParameters: GetUserPresignedUrlRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PresignedUrlResponse> {
        const response = await this.getUserPresignedUrlRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserReviewsRaw(requestParameters: GetUserReviewsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReviewsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getUserReviews().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/reviews/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReviewsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserReviews(requestParameters: GetUserReviewsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReviewsResponse> {
        const response = await this.getUserReviewsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserTimestampRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserTimestampResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/timestamp`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserTimestampResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserTimestamp(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserTimestampResponse> {
        const response = await this.getUserTimestampRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getUsersByIdsRaw(requestParameters: GetUsersByIdsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersResponse>> {
        const queryParameters: any = {};

        if (requestParameters['userIds'] != null) {
            queryParameters['user_ids[]'] = requestParameters['userIds'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters['xJwt'] != null) {
            headerParameters['X-Jwt'] = String(requestParameters['xJwt']);
        }

        const response = await this.request({
            path: `/v1/users/list_id`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUsersByIds(requestParameters: GetUsersByIdsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersResponse> {
        const response = await this.getUsersByIdsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getWebSocketTokenRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<WebSocketTokenResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/ws_token`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => WebSocketTokenResponseFromJSON(jsonValue));
    }

    /**
     */
    async getWebSocketToken(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<WebSocketTokenResponse> {
        const response = await this.getWebSocketTokenRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async loginWithEmailRaw(requestParameters: LoginWithEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LoginUserResponse>> {
        if (requestParameters['loginEmailUserRequest'] == null) {
            throw new runtime.RequiredError(
                'loginEmailUserRequest',
                'Required parameter "loginEmailUserRequest" was null or undefined when calling loginWithEmail().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v3/users/login_with_email`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: LoginEmailUserRequestToJSON(requestParameters['loginEmailUserRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LoginUserResponseFromJSON(jsonValue));
    }

    /**
     */
    async loginWithEmail(requestParameters: LoginWithEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginUserResponse> {
        const response = await this.loginWithEmailRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async logoutRaw(requestParameters: LogoutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v1/users/logout`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async logout(requestParameters: LogoutRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.logoutRaw(requestParameters, initOverrides);
    }

    /**
     */
    async pingAliveRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/alive`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async pingAlive(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.pingAliveRaw(initOverrides);
    }

    /**
     */
    async removeCoverImageRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/remove_cover_image`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeCoverImage(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeCoverImageRaw(initOverrides);
    }

    /**
     */
    async removeProfilePhotoRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/remove_profile_photo`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeProfilePhoto(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeProfilePhotoRaw(initOverrides);
    }

    /**
     */
    async replyToReviewRaw(requestParameters: ReplyToReviewRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling replyToReview().'
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

        if (requestParameters['comment'] != null) {
            formParams.append('comment', requestParameters['comment'] as any);
        }

        if (requestParameters['context'] != null) {
            formParams.append('context', requestParameters['context'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v2/users/reviews/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async replyToReview(requestParameters: ReplyToReviewRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.replyToReviewRaw(requestParameters, initOverrides);
    }

    /**
     */
    async reportUserRaw(requestParameters: ReportUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling reportUser().'
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
            path: `/v3/users/{user_id}/report`.replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async reportUser(requestParameters: ReportUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.reportUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async requestFollowRaw(requestParameters: RequestFollowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['targetId'] == null) {
            throw new runtime.RequiredError(
                'targetId',
                'Required parameter "targetId" was null or undefined when calling requestFollow().'
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

        if (requestParameters['action'] != null) {
            formParams.append('action', requestParameters['action'] as any);
        }

        const response = await this.request({
            path: `/v2/users/{target_id}/follow_request`.replace(`{${"target_id"}}`, encodeURIComponent(String(requestParameters['targetId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async requestFollow(requestParameters: RequestFollowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.requestFollowRaw(requestParameters, initOverrides);
    }

    /**
     */
    async resendConfirmEmailRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/resend_confirm_email`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async resendConfirmEmail(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.resendConfirmEmailRaw(initOverrides);
    }

    /**
     */
    async resetCountersRaw(requestParameters: ResetCountersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['counter'] != null) {
            formParams.append('counter', requestParameters['counter'] as any);
        }

        const response = await this.request({
            path: `/v1/users/reset_counters`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async resetCounters(requestParameters: ResetCountersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.resetCountersRaw(requestParameters, initOverrides);
    }

    /**
     */
    async resetPasswordRaw(requestParameters: ResetPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['email'] != null) {
            formParams.append('email', requestParameters['email'] as any);
        }

        if (requestParameters['emailGrantToken'] != null) {
            formParams.append('email_grant_token', requestParameters['emailGrantToken'] as any);
        }

        if (requestParameters['password'] != null) {
            formParams.append('password', requestParameters['password'] as any);
        }

        const response = await this.request({
            path: `/v1/users/reset_password`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async resetPassword(requestParameters: ResetPasswordRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.resetPasswordRaw(requestParameters, initOverrides);
    }

    /**
     */
    async searchUsersRaw(requestParameters: SearchUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersResponse>> {
        const queryParameters: any = {};

        if (requestParameters['gender'] != null) {
            queryParameters['gender'] = requestParameters['gender'];
        }

        if (requestParameters['nickname'] != null) {
            queryParameters['nickname'] = requestParameters['nickname'];
        }

        if (requestParameters['title'] != null) {
            queryParameters['title'] = requestParameters['title'];
        }

        if (requestParameters['biography'] != null) {
            queryParameters['biography'] = requestParameters['biography'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['similarAge'] != null) {
            queryParameters['similar_age'] = requestParameters['similarAge'];
        }

        if (requestParameters['notRecentGomimushi'] != null) {
            queryParameters['not_recent_gomimushi'] = requestParameters['notRecentGomimushi'];
        }

        if (requestParameters['recentlyCreated'] != null) {
            queryParameters['recently_created'] = requestParameters['recentlyCreated'];
        }

        if (requestParameters['samePrefecture'] != null) {
            queryParameters['same_prefecture'] = requestParameters['samePrefecture'];
        }

        if (requestParameters['saveRecentSearch'] != null) {
            queryParameters['save_recent_search'] = requestParameters['saveRecentSearch'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/search`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async searchUsers(requestParameters: SearchUsersRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersResponse> {
        const response = await this.searchUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async setHimaRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/users/hima`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async setHima(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.setHimaRaw(initOverrides);
    }

    /**
     */
    async unblockUserRaw(requestParameters: UnblockUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unblockUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/{id}/unblock`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unblockUser(requestParameters: UnblockUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unblockUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unfollowUserRaw(requestParameters: UnfollowUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unfollowUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/users/{id}/unfollow`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unfollowUser(requestParameters: UnfollowUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unfollowUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateAdditionalNotificationSettingRaw(requestParameters: UpdateAdditionalNotificationSettingRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['mode'] != null) {
            formParams.append('mode', requestParameters['mode'] as any);
        }

        if (requestParameters['on'] != null) {
            formParams.append('on', requestParameters['on'] as any);
        }

        const response = await this.request({
            path: `/v1/users/additonal_notification_setting`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateAdditionalNotificationSetting(requestParameters: UpdateAdditionalNotificationSettingRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateAdditionalNotificationSettingRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateLanguageRaw(requestParameters: UpdateLanguageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['language'] != null) {
            formParams.append('language', requestParameters['language'] as any);
        }

        if (requestParameters['signedInfo'] != null) {
            formParams.append('signed_info', requestParameters['signedInfo'] as any);
        }

        if (requestParameters['timestamp'] != null) {
            formParams.append('timestamp', requestParameters['timestamp'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v1/users/language`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateLanguage(requestParameters: UpdateLanguageRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateLanguageRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateLoginRaw(requestParameters: UpdateLoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LoginUpdateResponse>> {
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

        if (requestParameters['currentPassword'] != null) {
            formParams.append('current_password', requestParameters['currentPassword'] as any);
        }

        if (requestParameters['email'] != null) {
            formParams.append('email', requestParameters['email'] as any);
        }

        if (requestParameters['emailGrantToken'] != null) {
            formParams.append('email_grant_token', requestParameters['emailGrantToken'] as any);
        }

        if (requestParameters['password'] != null) {
            formParams.append('password', requestParameters['password'] as any);
        }

        const response = await this.request({
            path: `/v3/users/login_update`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LoginUpdateResponseFromJSON(jsonValue));
    }

    /**
     */
    async updateLogin(requestParameters: UpdateLoginRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginUpdateResponse> {
        const response = await this.updateLoginRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async updateUserInterestsRaw(requestParameters: UpdateUserInterestsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['commonIdsRequest'] == null) {
            throw new runtime.RequiredError(
                'commonIdsRequest',
                'Required parameter "commonIdsRequest" was null or undefined when calling updateUserInterests().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/users/interests`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: CommonIdsRequestToJSON(requestParameters['commonIdsRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateUserInterests(requestParameters: UpdateUserInterestsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateUserInterestsRaw(requestParameters, initOverrides);
    }

}