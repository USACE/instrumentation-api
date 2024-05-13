/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class MediaService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * serves media, files, etc for a given project
     * @param projectSlug project abbr
     * @param uriPath uri path of requested resource
     * @returns any OK
     * @throws ApiError
     */
    public getProjectsImages(
        projectSlug: string,
        uriPath: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_slug}/images/{uri_path}',
            path: {
                'project_slug': projectSlug,
                'uri_path': uriPath,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
