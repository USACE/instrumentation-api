/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Domain } from '../models/Domain';
import type { DomainMap } from '../models/DomainMap';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DomainService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all domains
     * @returns Domain OK
     * @throws ApiError
     */
    public getDomains(): CancelablePromise<Array<Domain>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/domains',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * Get map with domain group as key
     * @returns DomainMap OK
     * @throws ApiError
     */
    public getDomainsMap(): CancelablePromise<DomainMap> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/domains/map',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
