/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AwareParameter } from '../models/AwareParameter';
import type { AwarePlatformParameterConfig } from '../models/AwarePlatformParameterConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AwareService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists alert configs for a project
     * @returns AwarePlatformParameterConfig OK
     * @throws ApiError
     */
    public getAwareDataAcquisitionConfig(): CancelablePromise<Array<AwarePlatformParameterConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/aware/data_acquisition_config',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists alert configs for a project
     * @returns AwareParameter OK
     * @throws ApiError
     */
    public getAwareParameters(): CancelablePromise<Array<AwareParameter>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/aware/parameters',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
