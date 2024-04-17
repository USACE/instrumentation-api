/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ReportConfig } from '../models/ReportConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class PlotReportService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all report configs for a project
     * @param projectId project uuid
     * @param key api key
     * @returns ReportConfig OK
     * @throws ApiError
     */
    public getProjectsReportConfigs(
        projectId: string,
        key?: string,
    ): CancelablePromise<ReportConfig> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/report_configs',
            path: {
                'project_id': projectId,
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
     * creates a report config
     * @param projectId project uuid
     * @param key api key
     * @returns ReportConfig OK
     * @throws ApiError
     */
    public postProjectsReportConfigs(
        projectId: string,
        key?: string,
    ): CancelablePromise<ReportConfig> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/report_configs',
            path: {
                'project_id': projectId,
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
     * updates a report config
     * @param projectId project uuid
     * @param reportConfigId report config uuid
     * @param key api key
     * @returns ReportConfig OK
     * @throws ApiError
     */
    public putProjectsReportConfigs(
        projectId: string,
        reportConfigId: string,
        key?: string,
    ): CancelablePromise<ReportConfig> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/report_configs/{report_config_id}',
            path: {
                'project_id': projectId,
                'report_config_id': reportConfigId,
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
     * updates a report config
     * @param projectId project uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsReportConfigs(
        projectId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/report_configs/{report_config_id}',
            path: {
                'project_id': projectId,
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
