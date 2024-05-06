/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ReportConfig } from '../models/ReportConfig';
import type { ReportConfigWithPlotConfigs } from '../models/ReportConfigWithPlotConfigs';
import type { ReportDownloadJob } from '../models/ReportDownloadJob';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ReportConfigService {
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
     * @param reportConfig report config payload
     * @param key api key
     * @returns ReportConfig OK
     * @throws ApiError
     */
    public postProjectsReportConfigs(
        projectId: string,
        reportConfig: ReportConfig,
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
            body: reportConfig,
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
     * @param reportConfig report config payload
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public putProjectsReportConfigs(
        projectId: string,
        reportConfigId: string,
        reportConfig: ReportConfig,
        key?: string,
    ): CancelablePromise<any> {
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
            body: reportConfig,
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
    /**
     * @param projectId project uuid
     * @param reportConfigId report config uuid
     * @param key api key
     * @returns ReportDownloadJob Created
     * @throws ApiError
     */
    public postProjectsReportConfigsJobs(
        projectId: string,
        reportConfigId: string,
        key?: string,
    ): CancelablePromise<ReportDownloadJob> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/report_configs/{report_config_id}/jobs',
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
     * @param projectId project uuid
     * @param reportConfigId report config uuid
     * @param jobId download job uuid
     * @param key api key
     * @returns ReportDownloadJob OK
     * @throws ApiError
     */
    public getProjectsReportConfigsJobs(
        projectId: string,
        reportConfigId: string,
        jobId: string,
        key?: string,
    ): CancelablePromise<ReportDownloadJob> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/report_configs/{report_config_id}/jobs/{job_id}',
            path: {
                'project_id': projectId,
                'report_config_id': reportConfigId,
                'job_id': jobId,
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
     * @param jobId download job uuid
     * @param reportDownloadJob report download job payload
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public putProjectsReportConfigsJobs(
        jobId: string,
        reportDownloadJob: ReportDownloadJob,
        key: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/report_configs/{report_config_id}/jobs/{job_id}',
            path: {
                'job_id': jobId,
            },
            query: {
                'key': key,
            },
            body: reportDownloadJob,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * @param reportConfigId report config uuid
     * @param key api key
     * @returns ReportConfigWithPlotConfigs OK
     * @throws ApiError
     */
    public getReportConfigsPlotConfigs(
        reportConfigId: string,
        key: string,
    ): CancelablePromise<ReportConfigWithPlotConfigs> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/report_configs/{report_config_id}/plot_configs',
            path: {
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
}
