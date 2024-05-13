/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertSubscription } from '../models/AlertSubscription';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AlertSubscriptionService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * updates settings for an alert subscription
     * @param alertSubscriptionId alert subscription id
     * @param alertSubscription alert subscription payload
     * @param key api key
     * @returns AlertSubscription OK
     * @throws ApiError
     */
    public putAlertSubscriptions(
        alertSubscriptionId: string,
        alertSubscription: AlertSubscription,
        key?: string,
    ): CancelablePromise<Array<AlertSubscription>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/alert_subscriptions/{alert_subscription_id}',
            path: {
                'alert_subscription_id': alertSubscriptionId,
            },
            query: {
                'key': key,
            },
            body: alertSubscription,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists all alerts subscribed to by the current profile
     * @param key api key
     * @returns AlertSubscription OK
     * @throws ApiError
     */
    public getMyAlertSubscriptions(
        key?: string,
    ): CancelablePromise<Array<AlertSubscription>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/my_alert_subscriptions',
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
     * subscribes a profile to an alert
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param alertConfigId alert config uuid
     * @param key api key
     * @returns AlertSubscription OK
     * @throws ApiError
     */
    public postProjectsInstrumentsAlertConfigsSubscribe(
        projectId: string,
        instrumentId: string,
        alertConfigId: string,
        key?: string,
    ): CancelablePromise<AlertSubscription> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/instruments/{instrument_id}/alert_configs/{alert_config_id}/subscribe',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
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
     * unsubscribes a profile to an alert
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param alertConfigId alert config uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postProjectsInstrumentsAlertConfigsUnsubscribe(
        projectId: string,
        instrumentId: string,
        alertConfigId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/instruments/{instrument_id}/alert_configs/{alert_config_id}/unsubscribe',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
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
}
