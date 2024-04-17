/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Datalogger } from '../models/Datalogger';
import type { DataloggerTablePreview } from '../models/DataloggerTablePreview';
import type { DataloggerWithKey } from '../models/DataloggerWithKey';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DataloggerService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * creates a datalogger
     * @param datalogger datalogger payload
     * @param key api key
     * @returns DataloggerWithKey OK
     * @throws ApiError
     */
    public postDatalogger(
        datalogger: Datalogger,
        key?: string,
    ): CancelablePromise<Array<DataloggerWithKey>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/datalogger',
            query: {
                'key': key,
            },
            body: datalogger,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a datalogger by id
     * @param dataloggerId datalogger uuid
     * @param key api key
     * @returns Datalogger OK
     * @throws ApiError
     */
    public getDatalogger(
        dataloggerId: string,
        key?: string,
    ): CancelablePromise<Datalogger> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/datalogger/{datalogger_id}',
            path: {
                'datalogger_id': dataloggerId,
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
     * updates a datalogger
     * @param dataloggerId datalogger uuid
     * @param datalogger datalogger payload
     * @param key api key
     * @returns Datalogger OK
     * @throws ApiError
     */
    public putDatalogger(
        dataloggerId: string,
        datalogger: Datalogger,
        key?: string,
    ): CancelablePromise<Datalogger> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/datalogger/{datalogger_id}',
            path: {
                'datalogger_id': dataloggerId,
            },
            query: {
                'key': key,
            },
            body: datalogger,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a datalogger by id
     * @param dataloggerId datalogger uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteDatalogger(
        dataloggerId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/datalogger/{datalogger_id}',
            path: {
                'datalogger_id': dataloggerId,
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
     * deletes and recreates a datalogger api key
     * @param dataloggerId datalogger uuid
     * @param key api key
     * @returns DataloggerWithKey OK
     * @throws ApiError
     */
    public putDataloggerKey(
        dataloggerId: string,
        key?: string,
    ): CancelablePromise<DataloggerWithKey> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/datalogger/{datalogger_id}/key',
            path: {
                'datalogger_id': dataloggerId,
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
     * resets a datalogger table name to be renamed by incoming telemetry
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param key api key
     * @returns DataloggerTablePreview OK
     * @throws ApiError
     */
    public putDataloggerTablesName(
        dataloggerId: string,
        dataloggerTableId: string,
        key?: string,
    ): CancelablePromise<DataloggerTablePreview> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/name',
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
     * gets the most recent datalogger preview by by datalogger id
     * @param dataloggerId datalogger uuid
     * @param dataloggerTableId datalogger table uuid
     * @param key api key
     * @returns DataloggerTablePreview OK
     * @throws ApiError
     */
    public getDataloggerTablesPreview(
        dataloggerId: string,
        dataloggerTableId: string,
        key?: string,
    ): CancelablePromise<DataloggerTablePreview> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/datalogger/{datalogger_id}/tables/{datalogger_table_id}/preview',
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
     * lists dataloggers for a project
     * @param key api key
     * @returns Datalogger OK
     * @throws ApiError
     */
    public getDataloggers(
        key?: string,
    ): CancelablePromise<Array<Datalogger>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/dataloggers',
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
