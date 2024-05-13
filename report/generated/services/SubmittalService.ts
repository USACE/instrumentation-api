/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Submittal } from '../models/Submittal';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class SubmittalService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all submittals for an instrument
     * @param alertConfigId alert config uuid
     * @returns Submittal OK
     * @throws ApiError
     */
    public getAlertConfigsSubmittals(
        alertConfigId: string,
    ): CancelablePromise<Array<Submittal>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/alert_configs/{alert_config_id}/submittals',
            path: {
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
     * verifies all current submittals for the alert config are "missing" and will not be completed
     * @param alertConfigId alert config uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public putAlertConfigsSubmittalsVerifyMissing(
        alertConfigId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/alert_configs/{alert_config_id}/submittals/verify_missing',
            path: {
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
     * lists all submittals for an instrument
     * @param instrumentId instrument uuid
     * @param missing filter by missing projects only
     * @returns Submittal OK
     * @throws ApiError
     */
    public getInstrumentsSubmittals(
        instrumentId: string,
        missing?: boolean,
    ): CancelablePromise<Array<Submittal>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/submittals',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'missing': missing,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists all submittals for a project
     * @param projectId project uuid
     * @param missing filter by missing projects only
     * @returns Submittal OK
     * @throws ApiError
     */
    public getProjectsSubmittals(
        projectId: string,
        missing?: boolean,
    ): CancelablePromise<Array<Submittal>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/submittals',
            path: {
                'project_id': projectId,
            },
            query: {
                'missing': missing,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * verifies the specified submittal is "missing" and will not be completed
     * @param submittalId submittal uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public putSubmittalsVerifyMissing(
        submittalId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/submittals/{submittal_id}/verify_missing',
            path: {
                'submittal_id': submittalId,
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
