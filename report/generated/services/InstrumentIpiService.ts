/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IpiMeasurements } from '../models/IpiMeasurements';
import type { IpiSegment } from '../models/IpiSegment';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentIpiService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * creates instrument notes
     * @param instrumentId instrument uuid
     * @param before before time
     * @param after after time
     * @returns IpiMeasurements OK
     * @throws ApiError
     */
    public getInstrumentsIpiMeasurements(
        instrumentId: string,
        before: string,
        after?: string,
    ): CancelablePromise<Array<IpiMeasurements>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/ipi/{instrument_id}/measurements',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'after': after,
                'before': before,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets all ipi segments for an instrument
     * @param instrumentId instrument uuid
     * @returns IpiSegment OK
     * @throws ApiError
     */
    public getInstrumentsIpiSegments(
        instrumentId: string,
    ): CancelablePromise<Array<IpiSegment>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/ipi/{instrument_id}/segments',
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
     * updates multiple segments for an ipi instrument
     * @param instrumentId instrument uuid
     * @param instrumentSegments ipi instrument segments payload
     * @param key api key
     * @returns IpiSegment OK
     * @throws ApiError
     */
    public putInstrumentsIpiSegments(
        instrumentId: string,
        instrumentSegments: Array<IpiSegment>,
        key?: string,
    ): CancelablePromise<Array<IpiSegment>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/instruments/ipi/{instrument_id}/segments',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: instrumentSegments,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
