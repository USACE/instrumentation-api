/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { District } from '../models/District';
import type { IDSlugName } from '../models/IDSlugName';
import type { InstrumentGroup } from '../models/InstrumentGroup';
import type { Project } from '../models/Project';
import type { ProjectCount } from '../models/ProjectCount';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ProjectService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all districts
     * @returns District OK
     * @throws ApiError
     */
    public getDistricts(): CancelablePromise<Array<District>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/districts',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists projects where current profile is an admin or member with optional filter by project role
     * @param role role
     * @returns Project OK
     * @throws ApiError
     */
    public getMyProjects(
        role?: string,
    ): CancelablePromise<Array<Project>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/my_projects',
            query: {
                'role': role,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * lists all projects optionally filtered by federal id
     * @param federalId federal id
     * @returns Project OK
     * @throws ApiError
     */
    public getProjects(
        federalId?: string,
    ): CancelablePromise<Array<Project>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects',
            query: {
                'federal_id': federalId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * accepts an array of instruments for bulk upload to the database
     * @param projectCollection project collection payload
     * @param key api key
     * @returns IDSlugName OK
     * @throws ApiError
     */
    public postProjects(
        projectCollection: Array<Project>,
        key?: string,
    ): CancelablePromise<Array<IDSlugName>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects',
            query: {
                'key': key,
            },
            body: projectCollection,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets the total number of non-deleted projects in the system
     * @returns ProjectCount OK
     * @throws ApiError
     */
    public getProjectsCount(): CancelablePromise<ProjectCount> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/count',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single project
     * @param projectId project uuid
     * @returns Project OK
     * @throws ApiError
     */
    public getProjects1(
        projectId: string,
    ): CancelablePromise<Project> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}',
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
     * updates an existing project
     * @param projectId project uuid
     * @param project project payload
     * @param key api key
     * @returns Project OK
     * @throws ApiError
     */
    public putProjects(
        projectId: string,
        project: Project,
        key?: string,
    ): CancelablePromise<Project> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: project,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * soft deletes a project
     * @param projectId project id
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjects(
        projectId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}',
            path: {
                'project_id': projectId,
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
     * uploades a picture for a project
     * @param projectId project id
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public postProjectsImages(
        projectId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/images',
            path: {
                'project_id': projectId,
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
     * lists instrument groups associated with a project
     * @param projectId project uuid
     * @returns InstrumentGroup OK
     * @throws ApiError
     */
    public getProjectsInstrumentGroups(
        projectId: string,
    ): CancelablePromise<Array<InstrumentGroup>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instrument_groups',
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
     * lists instruments associated with a project
     * @param projectId project uuid
     * @returns Project OK
     * @throws ApiError
     */
    public getProjectsInstruments(
        projectId: string,
    ): CancelablePromise<Array<Project>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments',
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
}
