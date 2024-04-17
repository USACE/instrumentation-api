/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { InclinometerMeasurementCollectionLean } from '../models/InclinometerMeasurementCollectionLean';
import type { MeasurementCollectionLean } from '../models/MeasurementCollectionLean';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ExplorerService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * list timeseries measurements for explorer page
     * @param instrumentIds array of instrument uuids
     * @returns MeasurementCollectionLean OK
     * @throws ApiError
     */
    public postExplorer(
        instrumentIds: Array<string>,
    ): CancelablePromise<Array<Record<string, MeasurementCollectionLean>>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/explorer',
            body: instrumentIds,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * list inclinometer timeseries measurements for explorer page
     * @param instrumentIds array of inclinometer instrument uuids
     * @returns InclinometerMeasurementCollectionLean OK
     * @throws ApiError
     */
    public postInclinometerExplorer(
        instrumentIds: Array<string>,
    ): CancelablePromise<Array<Record<string, InclinometerMeasurementCollectionLean>>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/inclinometer_explorer',
            body: instrumentIds,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
