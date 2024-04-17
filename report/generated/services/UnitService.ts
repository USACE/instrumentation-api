/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Unit } from '../models/Unit';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class UnitService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists the available units
     * @returns Unit OK
     * @throws ApiError
     */
    public getUnits(): CancelablePromise<Array<Unit>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/units',
            errors: {
                400: `Bad Request`,
            },
        });
    }
}
