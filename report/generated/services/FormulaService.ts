/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CalculatedTimeseries } from '../models/CalculatedTimeseries';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class FormulaService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists calculations associated with an instrument
     * @returns CalculatedTimeseries OK
     * @throws ApiError
     */
    public getFormulas(): CancelablePromise<Array<CalculatedTimeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/formulas',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates a calculation
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postFormulas(
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/formulas',
            query: {
                'key': key,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates a calculation
     * @param formulaId formula uuid
     * @param key api key
     * @returns CalculatedTimeseries OK
     * @throws ApiError
     */
    public putFormulas(
        formulaId: string,
        key?: string,
    ): CancelablePromise<Array<CalculatedTimeseries>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/formulas/{formula_id}',
            path: {
                'formula_id': formulaId,
            },
            query: {
                'key': key,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a calculation
     * @param formulaId formula uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteFormulas(
        formulaId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/formulas/{formula_id}',
            path: {
                'formula_id': formulaId,
            },
            query: {
                'key': key,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
