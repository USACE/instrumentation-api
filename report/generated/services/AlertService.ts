/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Alert } from '../models/Alert';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AlertService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists subscribed alerts for a single user
     * list all alerts a profile is subscribed to
     * @param key api key
     * @returns Alert OK
     * @throws ApiError
     */
    public getMyAlerts(
        key?: string,
    ): CancelablePromise<Array<Alert>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/my_alerts',
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
     * marks an alert as read
     * marks an alert as read for a profile
     * returning the updated alert
     * @param alertId alert uuid
     * @param key api key
     * @returns Alert OK
     * @throws ApiError
     */
    public postMyAlertsRead(
        alertId: string,
        key?: string,
    ): CancelablePromise<Alert> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/my_alerts/{alert_id}/read',
            path: {
                'alert_id': alertId,
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
     * marks an alert as unread for a profile
     * marks an alert as unread based on provided profile ID and alert ID.
     * returning the updated alert
     * @param alertId alert uuid
     * @param key api key
     * @returns Alert OK
     * @throws ApiError
     */
    public postMyAlertsUnread(
        alertId: string,
        key?: string,
    ): CancelablePromise<Alert> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/my_alerts/{alert_id}/unread',
            path: {
                'alert_id': alertId,
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
     * list all alerts associated an instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @returns Alert OK
     * @throws ApiError
     */
    public getProjectsInstrumentsAlerts(
        projectId: string,
        instrumentId: string,
    ): CancelablePromise<Array<Alert>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments/{instrument_id}/alerts',
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
