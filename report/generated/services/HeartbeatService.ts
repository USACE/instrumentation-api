/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Heartbeat } from '../models/Heartbeat';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class HeartbeatService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * checks the health of the api server
     * @returns any OK
     * @throws ApiError
     */
    public getHealth(): CancelablePromise<Array<any>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/health',
        });
    }
    /**
     * creates a heartbeat entry at regular intervals
     * @param key api key
     * @returns Heartbeat OK
     * @throws ApiError
     */
    public postHeartbeat(
        key: string,
    ): CancelablePromise<Heartbeat> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/heartbeat',
            query: {
                'key': key,
            },
        });
    }
    /**
     * gets the latest heartbeat
     * @returns Heartbeat OK
     * @throws ApiError
     */
    public getHeartbeatLatest(): CancelablePromise<Heartbeat> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/heartbeat/latest',
        });
    }
    /**
     * returns all heartbeats
     * @returns Heartbeat OK
     * @throws ApiError
     */
    public getHeartbeats(): CancelablePromise<Array<Heartbeat>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/heartbeats',
        });
    }
}
