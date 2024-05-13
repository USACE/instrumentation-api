/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfig } from '../models/AlertConfig';
import type { InstrumentStatus } from '../models/InstrumentStatus';
import type { InstrumentStatusCollection } from '../models/InstrumentStatusCollection';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentStatusService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all Status for an instrument
     * @param instrumentId instrument uuid
     * @returns InstrumentStatus OK
     * @throws ApiError
     */
    public getInstrumentsStatus(
        instrumentId: string,
    ): CancelablePromise<Array<InstrumentStatus>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/status',
            path: {
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
     * creates a status for an instrument
     * @param instrumentId instrument uuid
     * @param instrumentStatus instrument status collection paylaod
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postInstrumentsStatus(
        instrumentId: string,
        instrumentStatus: InstrumentStatusCollection,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/instruments/{instrument_id}/status',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: instrumentStatus,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single status
     * @param instrumentId instrument uuid
     * @param statusId status uuid
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public getInstrumentsStatus1(
        instrumentId: string,
        statusId: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/status/{status_id}',
            path: {
                'instrument_id': instrumentId,
                'status_id': statusId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a status for an instrument
     * @param instrumentId instrument uuid
     * @param statusId project uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteInstrumentsStatus(
        instrumentId: string,
        statusId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/instruments/{instrument_id}/status/{status_id}',
            path: {
                'instrument_id': instrumentId,
                'status_id': statusId,
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
