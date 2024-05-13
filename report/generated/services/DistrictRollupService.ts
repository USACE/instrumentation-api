/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { DistrictRollup } from '../models/DistrictRollup';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DistrictRollupService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists monthly evaluation statistics for a district by project id
     * @param projectId project id
     * @returns DistrictRollup OK
     * @throws ApiError
     */
    public getProjectsDistrictRollupEvaluationSubmittals(
        projectId: string,
    ): CancelablePromise<Array<DistrictRollup>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/district_rollup/evaluation_submittals',
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
     * lists monthly measurement statistics for a district by project id
     * @param projectId project id
     * @returns DistrictRollup OK
     * @throws ApiError
     */
    public getProjectsDistrictRollupMeasurementSubmittals(
        projectId: string,
    ): CancelablePromise<Array<DistrictRollup>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/district_rollup/measurement_submittals',
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
}
