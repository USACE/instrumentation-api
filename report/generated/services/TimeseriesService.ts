/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { MeasurementCollection } from '../models/MeasurementCollection';
import type { Timeseries } from '../models/Timeseries';
import type { TimeseriesCollectionItems } from '../models/TimeseriesCollectionItems';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class TimeseriesService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists timeseries for instruments in an instrument group
     * @param instrumentGroupId instrument group uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getInstrumentGroupsTimeseries(
        instrumentGroupId: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instrument_groups/{instrument_group_id}/timeseries',
            path: {
                'instrument_group_id': instrumentGroupId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists timeseries measurements by instrument group id
     * @param instrumentGroupId instrument group uuid
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public getInstrumentGroupsTimeseriesMeasurements(
        instrumentGroupId: string,
    ): CancelablePromise<MeasurementCollection> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instrument_groups/{instrument_group_id}/timeseries_measurements',
            path: {
                'instrument_group_id': instrumentGroupId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single timeseries by id
     * @param timeseriesId timeseries uuid
     * @param instrumentId instrument uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getInstrumentsTimeseries(
        timeseriesId: string,
        instrumentId: string,
    ): CancelablePromise<Timeseries> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/timeseries/{timeseries_id}',
            path: {
                'timeseries_id': timeseriesId,
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
     * lists timeseries by timeseries uuid
     * @param instrumentId instrument uuid
     * @param timeseriesId timeseries uuid
     * @param after after time
     * @param before before time
     * @param threshold downsample threshold
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public getInstrumentsTimeseriesMeasurements(
        instrumentId: string,
        timeseriesId: string,
        after?: string,
        before?: string,
        threshold?: number,
    ): CancelablePromise<MeasurementCollection> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/timeseries/{timeseries_id}/measurements',
            path: {
                'instrument_id': instrumentId,
                'timeseries_id': timeseriesId,
            },
            query: {
                'after': after,
                'before': before,
                'threshold': threshold,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists timeseries measurements by instrument id
     * @param instrumentId instrument uuid
     * @param after after time
     * @param before before time
     * @param threshold downsample threshold
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public getInstrumentsTimeseriesMeasurements1(
        instrumentId: string,
        after?: string,
        before?: string,
        threshold?: number,
    ): CancelablePromise<MeasurementCollection> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/timeseries_measurements',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'after': after,
                'before': before,
                'threshold': threshold,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists timeseries for an instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getProjectsInstrumentsTimeseries(
        projectId: string,
        instrumentId: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments/{instrument_id}/timeseries',
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
     * lists all timeseries for a single plot config
     * @param projectId project uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getProjectsPlotConfigurationsTimeseries(
        projectId: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/plot_configurations/{plot_config_id}/timeseries',
            path: {
                'project_id': projectId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists all timeseries for a single project
     * @param projectId project uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getProjectsTimeseries(
        projectId: string,
    ): CancelablePromise<Array<Timeseries>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/timeseries',
            path: {
                'project_id': projectId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates one or more timeseries
     * @param timeseriesCollectionItems timeseries collection items payload
     * @param key api key
     * @returns string OK
     * @throws ApiError
     */
    public postTimeseries(
        timeseriesCollectionItems: TimeseriesCollectionItems,
        key?: string,
    ): CancelablePromise<Array<Record<string, string>>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/timeseries',
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
     * gets a single timeseries by id
     * @param timeseriesId timeseries uuid
     * @returns Timeseries OK
     * @throws ApiError
     */
    public getTimeseries(
        timeseriesId: string,
    ): CancelablePromise<Timeseries> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/timeseries/{timeseries_id}',
            path: {
                'timeseries_id': timeseriesId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates a single timeseries by id
     * @param timeseriesId timeseries uuid
     * @param timeseries timeseries payload
     * @param key api key
     * @returns string OK
     * @throws ApiError
     */
    public putTimeseries(
        timeseriesId: string,
        timeseries: Timeseries,
        key?: string,
    ): CancelablePromise<Record<string, string>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/timeseries/{timeseries_id}',
            path: {
                'timeseries_id': timeseriesId,
            },
            query: {
                'key': key,
            },
            body: timeseries,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a single timeseries by id
     * @param timeseriesId timeseries uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteTimeseries(
        timeseriesId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/timeseries/{timeseries_id}',
            path: {
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
    /**
     * lists timeseries by timeseries uuid
     * @param timeseriesId timeseries uuid
     * @param after after time
     * @param before before time
     * @param threshold downsample threshold
     * @returns MeasurementCollection OK
     * @throws ApiError
     */
    public getTimeseriesMeasurements(
        timeseriesId: string,
        after?: string,
        before?: string,
        threshold?: number,
    ): CancelablePromise<MeasurementCollection> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/timeseries/{timeseries_id}/measurements',
            path: {
                'timeseries_id': timeseriesId,
            },
            query: {
                'after': after,
                'before': before,
                'threshold': threshold,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
