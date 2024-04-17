/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { InclinometerMeasurementCollection } from '../models/InclinometerMeasurementCollection';
import type { InclinometerMeasurementCollectionCollection } from '../models/InclinometerMeasurementCollectionCollection';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class MeasurementInclinometerService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * creates or updates one or more inclinometer measurements
     * @param projectId project uuid
     * @param timeseriesMeasurementCollections inclinometer measurement collections
     * @param key api key
     * @returns InclinometerMeasurementCollection OK
     * @throws ApiError
     */
    public postProjectsInclinometerMeasurements(
        projectId: string,
        timeseriesMeasurementCollections: InclinometerMeasurementCollectionCollection,
        key?: string,
    ): CancelablePromise<Array<InclinometerMeasurementCollection>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/inclinometer_measurements',
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
     * lists all measurements for an inclinometer
     * @param timeseriesId timeseries uuid
     * @param after after timestamp
     * @param before before timestamp
     * @returns InclinometerMeasurementCollection OK
     * @throws ApiError
     */
    public getTimeseriesInclinometerMeasurements(
        timeseriesId: string,
        after?: string,
        before?: string,
    ): CancelablePromise<InclinometerMeasurementCollection> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/timeseries/{timeseries_id}/inclinometer_measurements',
            path: {
                'timeseries_id': timeseriesId,
            },
            query: {
                'after': after,
                'before': before,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a single inclinometer measurement by timestamp
     * @param timeseriesId timeseries uuid
     * @param time timestamp of measurement to delete
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteTimeseriesInclinometerMeasurements(
        timeseriesId: string,
        time: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/timeseries/{timeseries_id}/inclinometer_measurements',
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
}
