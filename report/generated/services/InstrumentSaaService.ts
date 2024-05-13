/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { SaaMeasurements } from '../models/SaaMeasurements';
import type { SaaSegment } from '../models/SaaSegment';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentSaaService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * creates instrument notes
     * @param instrumentId instrument uuid
     * @param before before time
     * @param after after time
     * @returns SaaMeasurements OK
     * @throws ApiError
     */
    public getInstrumentsSaaMeasurements(
        instrumentId: string,
        before: string,
        after?: string,
    ): CancelablePromise<Array<SaaMeasurements>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/saa/{instrument_id}/measurements',
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
     * gets all saa segments for an instrument
     * @param instrumentId instrument uuid
     * @returns SaaSegment OK
     * @throws ApiError
     */
    public getInstrumentsSaaSegments(
        instrumentId: string,
    ): CancelablePromise<Array<SaaSegment>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/saa/{instrument_id}/segments',
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
     * updates multiple segments for an saa instrument
     * @param instrumentId instrument uuid
     * @param instrumentSegments saa instrument segments payload
     * @param key api key
     * @returns SaaSegment OK
     * @throws ApiError
     */
    public putInstrumentsSaaSegments(
        instrumentId: string,
        instrumentSegments: Array<SaaSegment>,
        key?: string,
    ): CancelablePromise<Array<SaaSegment>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/instruments/saa/{instrument_id}/segments',
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
