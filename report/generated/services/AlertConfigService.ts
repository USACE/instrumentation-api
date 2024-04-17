/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfig } from '../models/AlertConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AlertConfigService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists alert configs for a project
     * @param projectId project uuid
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public getProjectsAlertConfigs(
        projectId: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/alert_configs',
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
     * creates one alert config
     * @param projectId project uuid
     * @param alertConfig alert config payload
     * @param key api key
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public postProjectsAlertConfigs(
        projectId: string,
        alertConfig: AlertConfig,
        key?: string,
    ): CancelablePromise<AlertConfig> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/alert_configs',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: alertConfig,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single alert
     * @param projectId project uuid
     * @param alertConfigId alert config uuid
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public getProjectsAlertConfigs1(
        projectId: string,
        alertConfigId: string,
    ): CancelablePromise<AlertConfig> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/alert_configs/{alert_config_id}',
            path: {
                'project_id': projectId,
                'alert_config_id': alertConfigId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an existing alert config
     * @param projectId project uuid
     * @param alertConfigId alert config uuid
     * @param alertConfig alert config payload
     * @param key api key
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public putProjectsAlertConfigs(
        projectId: string,
        alertConfigId: string,
        alertConfig: AlertConfig,
        key?: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/alert_configs/{alert_config_id}',
            path: {
                'project_id': projectId,
                'alert_config_id': alertConfigId,
            },
            query: {
                'key': key,
            },
            body: alertConfig,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes an alert config
     * @param projectId Project ID
     * @param alertConfigId instrument uuid
     * @param key api key
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public deleteProjectsAlertConfigs(
        projectId: string,
        alertConfigId: string,
        key?: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/alert_configs/{alert_config_id}',
            path: {
                'project_id': projectId,
                'alert_config_id': alertConfigId,
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
     * lists alerts for a single instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public getProjectsInstrumentsAlertConfigs(
        projectId: string,
        instrumentId: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments/{instrument_id}/alert_configs',
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
}
