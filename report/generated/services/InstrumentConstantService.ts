/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Timeseries } from '../models/Timeseries';
import type { TimeseriesCollectionItems } from '../models/TimeseriesCollectionItems';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentConstantService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists constants for a given instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getProjectsInstrumentsConstants(
        projectId: string,
        instrumentId: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments/{instrument_id}/constants',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates instrument constants (i.e. timeseries)
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param timeseriesCollectionItems timeseries collection items payload
     * @param key api key
     * @returns Timeseries OK
     * @throws ApiError
     */
    public postProjectsInstrumentsConstants(
        projectId: string,
        instrumentId: string,
        timeseriesCollectionItems: TimeseriesCollectionItems,
        key?: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/instruments/{instrument_id}/constants',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: timeseriesCollectionItems,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * removes a timeseries as an instrument constant
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param timeseriesId timeseries uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsInstrumentsConstants(
        projectId: string,
        instrumentId: string,
        timeseriesId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/instruments/{instrument_id}/constants/{timeseries_id}',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
                'timeseries_id': timeseriesId,
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
