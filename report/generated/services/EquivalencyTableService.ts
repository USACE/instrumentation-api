/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { EquivalencyTable } from '../models/EquivalencyTable';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class EquivalencyTableService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * creates an equivalency table for a datalogger and auto create data logger table if not exists
     * @param dataloggerId datalogger uuid
     * @param equivalencyTable equivalency table payload
     * @param key api key
     * @returns EquivalencyTable OK
     * @throws ApiError
     */
    public postDataloggerEquivalencyTable(
        dataloggerId: string,
        equivalencyTable: EquivalencyTable,
        key?: string,
    ): CancelablePromise<EquivalencyTable> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/datalogger/{datalogger_id}/equivalency_table',
            path: {
                'datalogger_id': dataloggerId,
            },
            query: {
                'key': key,
            },
            body: equivalencyTable,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets an equivalency table for a datalogger
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param key api key
     * @returns EquivalencyTable OK
     * @throws ApiError
     */
    public getDataloggerTablesEquivalencyTable(
        dataloggerId: string,
        dataloggerTableId: string,
        key?: string,
    ): CancelablePromise<Array<EquivalencyTable>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table',
            path: {
                'datalogger_id': dataloggerId,
                'datalogger_table_id': dataloggerTableId,
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
     * updates an equivalency table for a datalogger
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param equivalencyTable equivalency table payload
     * @param key api key
     * @returns EquivalencyTable OK
     * @throws ApiError
     */
    public putDataloggerTablesEquivalencyTable(
        dataloggerId: string,
        dataloggerTableId: string,
        equivalencyTable: EquivalencyTable,
        key?: string,
    ): CancelablePromise<EquivalencyTable> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table',
            path: {
                'datalogger_id': dataloggerId,
                'datalogger_table_id': dataloggerTableId,
            },
            query: {
                'key': key,
            },
            body: equivalencyTable,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates an equivalency table for a datalogger and auto create data logger table if not exists
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param equivalencyTable equivalency table payload
     * @param key api key
     * @returns EquivalencyTable OK
     * @throws ApiError
     */
    public postDataloggerTablesEquivalencyTable(
        dataloggerId: string,
        dataloggerTableId: string,
        equivalencyTable: EquivalencyTable,
        key?: string,
    ): CancelablePromise<EquivalencyTable> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table',
            path: {
                'datalogger_id': dataloggerId,
                'datalogger_table_id': dataloggerTableId,
            },
            query: {
                'key': key,
            },
            body: equivalencyTable,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes an equivalency table and corresponding datalogger table
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteDataloggerTablesEquivalencyTable(
        dataloggerId: string,
        dataloggerTableId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table',
            path: {
                'datalogger_id': dataloggerId,
                'datalogger_table_id': dataloggerTableId,
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
     * deletes an equivalency table row
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param rowId equivalency table row uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteDataloggerTablesEquivalencyTableRow(
        dataloggerId: string,
        dataloggerTableId: string,
        rowId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table/row/{row_id}',
            path: {
                'datalogger_id': dataloggerId,
                'datalogger_table_id': dataloggerTableId,
                'row_id': rowId,
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
