/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { EmailAutocompleteResult } from '../models/EmailAutocompleteResult';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AutocompleteService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists results of email autocomplete
     * @param q search query string
     * @returns EmailAutocompleteResult OK
     * @throws ApiError
     */
    public getEmailAutocomplete(
        q: string,
    ): CancelablePromise<Array<EmailAutocompleteResult>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/email_autocomplete',
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
