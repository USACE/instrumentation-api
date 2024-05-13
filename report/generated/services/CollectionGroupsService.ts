/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfig } from '../models/AlertConfig';
import type { CollectionGroup } from '../models/CollectionGroup';
import type { CollectionGroupDetails } from '../models/CollectionGroupDetails';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class CollectionGroupsService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists instrument groups
     * @param projectId project uuid
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public getProjectsCollectionGroups(
        projectId: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/collection_groups',
            path: {
                'project_id': projectId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates a new collection group
     * lists alert configs for a single project optionally filtered by alert_type_id
     * @param projectId project uuid
     * @param collectionGroup collection group payload
     * @param key api key
     * @returns CollectionGroup OK
     * @throws ApiError
     */
    public postProjectsCollectionGroups(
        projectId: string,
        collectionGroup: CollectionGroup,
        key?: string,
    ): CancelablePromise<Array<CollectionGroup>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/collection_groups',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: collectionGroup,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets all data needed to render collection group form
     * @param projectId project uuid
     * @param collectionGroupId collection group uuid
     * @returns CollectionGroupDetails OK
     * @throws ApiError
     */
    public getProjectsCollectionGroups1(
        projectId: string,
        collectionGroupId: string,
    ): CancelablePromise<CollectionGroupDetails> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/collection_groups/{collection_group_id}',
            path: {
                'project_id': projectId,
                'collection_group_id': collectionGroupId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an existing collection group
     * @param projectId project uuid
     * @param collectionGroupId collection group uuid
     * @param collectionGroup collection group payload
     * @param key api key
     * @returns CollectionGroup OK
     * @throws ApiError
     */
    public putProjectsCollectionGroups(
        projectId: string,
        collectionGroupId: string,
        collectionGroup: CollectionGroup,
        key?: string,
    ): CancelablePromise<CollectionGroup> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/collection_groups/{collection_group_id}',
            path: {
                'project_id': projectId,
                'collection_group_id': collectionGroupId,
            },
            query: {
                'key': key,
            },
            body: collectionGroup,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes a collection group using the id of the collection group
     * @param projectId project uuid
     * @param collectionGroupId collection group uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsCollectionGroups(
        projectId: string,
        collectionGroupId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/collection_groups/{collection_group_id}',
            path: {
                'project_id': projectId,
                'collection_group_id': collectionGroupId,
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
     * adds a timeseries to a collection group
     * @param projectId project uuid
     * @param collectionGroupId collection group uuid
     * @param timeseriesId timeseries uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postProjectsCollectionGroupsTimeseries(
        projectId: string,
        collectionGroupId: string,
        timeseriesId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id}',
            path: {
                'project_id': projectId,
                'collection_group_id': collectionGroupId,
                'timeseries_id': timeseriesId,
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
     * removes a timeseries from a collection group
     * @param projectId project uuid
     * @param collectionGroupId collection group uuid
     * @param timeseriesId timeseries uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsCollectionGroupsTimeseries(
        projectId: string,
        collectionGroupId: string,
        timeseriesId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id}',
            path: {
                'project_id': projectId,
                'collection_group_id': collectionGroupId,
                'timeseries_id': timeseriesId,
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
