/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Profile } from '../models/Profile';
import type { Token } from '../models/Token';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ProfileService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * gets profile for current authenticated user
     * @returns Profile OK
     * @throws ApiError
     */
    public getMyProfile(): CancelablePromise<Profile> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/my_profile',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates token for a profile
     * @returns Token OK
     * @throws ApiError
     */
    public postMyTokens(): CancelablePromise<Token> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/my_tokens',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a token for a profile
     * @param tokenId token uuid
     * @returns any OK
     * @throws ApiError
     */
    public deleteMyTokens(
        tokenId: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/my_tokens/{token_id}',
            path: {
                'token_id': tokenId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates a user profile
     * @returns Profile OK
     * @throws ApiError
     */
    public postProfiles(): CancelablePromise<Profile> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/profiles',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
