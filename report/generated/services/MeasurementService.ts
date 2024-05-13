/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { MeasurementCollection } from '../models/MeasurementCollection';
import type { TimeseriesMeasurementCollectionCollection } from '../models/TimeseriesMeasurementCollectionCollection';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class MeasurementService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * overwrites all measurements witin date range with the supplied payload
     * @param projectId project uuid
     * @param timeseriesMeasurementCollections array of timeseries measurement collections
     * @param after after timestamp
     * @param before before timestamp
     * @param key api key
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public putProjectsTimeseriesMeasurements(
        projectId: string,
        timeseriesMeasurementCollections: TimeseriesMeasurementCollectionCollection,
        after?: string,
        before?: string,
        key?: string,
    ): CancelablePromise<Array<MeasurementCollection>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/timeseries_measurements',
            path: {
                'project_id': projectId,
            },
            query: {
                'after': after,
                'before': before,
                'key': key,
            },
            body: timeseriesMeasurementCollections,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates or updates one or more timeseries measurements
     * @param projectId project uuid
     * @param timeseriesMeasurementCollections array of timeseries measurement collections
     * @param key api key
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public postProjectsTimeseriesMeasurements(
        projectId: string,
        timeseriesMeasurementCollections: TimeseriesMeasurementCollectionCollection,
        key?: string,
    ): CancelablePromise<Array<MeasurementCollection>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/timeseries_measurements',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: timeseriesMeasurementCollections,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a single timeseries measurement by timestamp
     * @param timeseriesId timeseries uuid
     * @param time timestamp of measurement to delete
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteTimeseriesMeasurements(
        timeseriesId: string,
        time: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/timeseries/{timeseries_id}/measurements',
            path: {
                'timeseries_id': timeseriesId,
            },
            query: {
                'time': time,
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
     * creates or updates one or more timeseries measurements
     * @param timeseriesMeasurementCollections array of timeseries measurement collections
     * @param key api key
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public postTimeseriesMeasurements(
        timeseriesMeasurementCollections: TimeseriesMeasurementCollectionCollection,
        key: string,
    ): CancelablePromise<Array<MeasurementCollection>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/timeseries_measurements',
            query: {
                'key': key,
            },
            body: timeseriesMeasurementCollections,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
