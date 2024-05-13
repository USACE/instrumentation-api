/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Instrument } from '../models/Instrument';
import type { InstrumentGroup } from '../models/InstrumentGroup';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentGroupService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all instrument groups
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public getInstrumentGroups(): CancelablePromise<Array<InstrumentGroup>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instrument_groups',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creats an instrument group from an array of instruments
     * @param instrumentGroup instrument group payload
     * @param key api key
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public postInstrumentGroups(
        instrumentGroup: InstrumentGroup,
        key?: string,
    ): CancelablePromise<InstrumentGroup> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/instrument_groups',
            query: {
                'key': key,
            },
            body: instrumentGroup,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single instrument group
     * @param instrumentGroupId instrument group uuid
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public getInstrumentGroups1(
        instrumentGroupId: string,
    ): CancelablePromise<InstrumentGroup> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instrument_groups/{instrument_group_id}',
            path: {
                'instrument_group_id': instrumentGroupId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an existing instrument group
     * @param instrumentGroupId instrument group uuid
     * @param instrumentGroup instrument group payload
     * @param key api key
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public putInstrumentGroups(
        instrumentGroupId: string,
        instrumentGroup: InstrumentGroup,
        key?: string,
    ): CancelablePromise<InstrumentGroup> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/instrument_groups/{instrument_group_id}',
            path: {
                'instrument_group_id': instrumentGroupId,
            },
            query: {
                'key': key,
            },
            body: instrumentGroup,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * soft deletes an instrument
     * @param instrumentGroupId instrument group uuid
     * @param key api key
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public deleteInstrumentGroups(
        instrumentGroupId: string,
        key?: string,
    ): CancelablePromise<Array<InstrumentGroup>> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/instrument_groups/{instrument_group_id}',
            path: {
                'instrument_group_id': instrumentGroupId,
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
     * lists instruments in an instrument group
     * @param instrumentGroupId instrument group uuid
     * @returns Instrument OK
     * @throws ApiError
     */
    public getInstrumentGroupsInstruments(
        instrumentGroupId: string,
    ): CancelablePromise<Array<Instrument>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instrument_groups/{instrument_group_id}/instruments',
            path: {
                'instrument_group_id': instrumentGroupId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * adds an instrument to an instrument group
     * @param instrumentGroupId instrument group uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postInstrumentGroupsInstruments(
        instrumentGroupId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/instrument_groups/{instrument_group_id}/instruments',
            path: {
                'instrument_group_id': instrumentGroupId,
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
     * removes an instrument from an instrument group
     * @param instrumentGroupId instrument group uuid
     * @param instrumentId instrument uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteInstrumentGroupsInstruments(
        instrumentGroupId: string,
        instrumentId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/instrument_groups/{instrument_group_id}/instruments/{instrument_id}',
            path: {
                'instrument_group_id': instrumentGroupId,
                'instrument_id': instrumentId,
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
