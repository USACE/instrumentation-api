/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Home } from '../models/Home';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class HomeService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * gets information for the homepage
     * @returns Home OK
     * @throws ApiError
     */
    public getHome(): CancelablePromise<Home> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/home',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }
}
