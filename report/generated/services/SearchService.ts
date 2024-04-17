/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { SearchResult } from '../models/SearchResult';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class SearchService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * allows searching using a string on different entities
     * @param entity entity to search (i.e. projects, etc.)
     * @param q search string
     * @returns SearchResult OK
     * @throws ApiError
     */
    public getSearch(
        entity: string,
        q?: string,
    ): CancelablePromise<Array<SearchResult>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/search/{entity}',
            path: {
                'entity': entity,
            },
            query: {
                'q': q,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
