/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ReportConfigWithPlotConfigs } from '../models/ReportConfigWithPlotConfigs';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ReportConfigService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
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
