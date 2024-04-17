/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { PlotConfig } from '../models/PlotConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class PlotConfigService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists plot configs
     * @param projectId project uuid
     * @returns PlotConfig OK
     * @throws ApiError
     */
    public getProjectsPlotConfigurations(
        projectId: string,
    ): CancelablePromise<Array<PlotConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/plot_configurations',
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
     * @param projectId project uuid
     * @param plotConfig plot config payload
     * @param key api key
     * @returns PlotConfig OK
     * @throws ApiError
     */
    public postProjectsPlotConfigurations(
        projectId: string,
        plotConfig: PlotConfig,
        key?: string,
    ): CancelablePromise<PlotConfig> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/plot_configurations',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: plotConfig,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * @param projectId project uuid
     * @param plotConfigurationId plot config uuid
     * @returns PlotConfig OK
     * @throws ApiError
     */
    public getProjectsPlotConfigurations1(
        projectId: string,
        plotConfigurationId: string,
    ): CancelablePromise<PlotConfig> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/plot_configurations/{plot_configuration_id}',
            path: {
                'project_id': projectId,
                'plot_configuration_id': plotConfigurationId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * @param projectId project uuid
     * @param plotConfigurationId plot config uuid
     * @param plotConfig plot config payload
     * @param key api key
     * @returns PlotConfig OK
     * @throws ApiError
     */
    public putProjectsPlotConfigurations(
        projectId: string,
        plotConfigurationId: string,
        plotConfig: PlotConfig,
        key?: string,
    ): CancelablePromise<PlotConfig> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/plot_configurations/{plot_configuration_id}',
            path: {
                'project_id': projectId,
                'plot_configuration_id': plotConfigurationId,
            },
            query: {
                'key': key,
            },
            body: plotConfig,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * @param projectId project uuid
     * @param plotConfigurationId plot config uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsPlotConfigurations(
        projectId: string,
        plotConfigurationId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/plot_configurations/{plot_configuration_id}',
            path: {
                'project_id': projectId,
                'plot_configuration_id': plotConfigurationId,
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
     * @param projectId project uuid
     * @param plotConfigurationId plot config uuid
     * @returns binary OK
     * @throws ApiError
     */
    public getProjectsReportConfigsDownloads(
        projectId: string,
        plotConfigurationId: string,
    ): CancelablePromise<Blob> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/report_configs/{report_config_id}/downloads',
            path: {
                'project_id': projectId,
                'plot_configuration_id': plotConfigurationId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
