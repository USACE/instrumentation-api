/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Site } from '../models/Site';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class OpendcsService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all instruments, represented as opendcs sites
     * @returns Site OK
     * @throws ApiError
     */
    public getOpendcsSites(): CancelablePromise<Array<Site>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/opendcs/sites',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
