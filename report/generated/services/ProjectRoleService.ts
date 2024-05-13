/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ProjectMembership } from '../models/ProjectMembership';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ProjectRoleService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists project members and their role information
     * @param projectId project uuid
     * @param key api key
     * @returns ProjectMembership OK
     * @throws ApiError
     */
    public getProjectsMembers(
        projectId: string,
        key?: string,
    ): CancelablePromise<Array<ProjectMembership>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/members',
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
     * adds project members and their role information
     * @param projectId project uuid
     * @param profileId profile uuid
     * @param roleId role uuid
     * @param key api key
     * @returns ProjectMembership OK
     * @throws ApiError
     */
    public postProjectsMembersRoles(
        projectId: string,
        profileId: string,
        roleId: string,
        key?: string,
    ): CancelablePromise<ProjectMembership> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/members/{profile_id}/roles/{role_id}',
            path: {
                'project_id': projectId,
                'profile_id': profileId,
                'role_id': roleId,
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
     * removes project members and their role information
     * @param projectId project uuid
     * @param profileId profile uuid
     * @param roleId role uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsMembersRoles(
        projectId: string,
        profileId: string,
        roleId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/members/{profile_id}/roles/{role_id}',
            path: {
                'project_id': projectId,
                'profile_id': profileId,
                'role_id': roleId,
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
