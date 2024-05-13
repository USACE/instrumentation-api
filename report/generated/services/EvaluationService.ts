/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfig } from '../models/AlertConfig';
import type { Evaluation } from '../models/Evaluation';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class EvaluationService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists evaluations for a single project optionally filtered by alert_config_id
     * @param projectId project uuid
     * @returns Evaluation OK
     * @throws ApiError
     */
    public getProjectsEvaluations(
        projectId: string,
    ): CancelablePromise<Array<Evaluation>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/evaluations',
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
     * creates one evaluation
     * @param projectId project uuid
     * @param evaluation evaluation payload
     * @param key api key
     * @returns Evaluation OK
     * @throws ApiError
     */
    public postProjectsEvaluations(
        projectId: string,
        evaluation: Evaluation,
        key?: string,
    ): CancelablePromise<Evaluation> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/evaluations',
            path: {
                'project_id': projectId,
            },
            query: {
                'key': key,
            },
            body: evaluation,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single evaluation by id
     * @param projectId project uuid
     * @param evaluationId evaluation uuid
     * @returns Evaluation OK
     * @throws ApiError
     */
    public getProjectsEvaluations1(
        projectId: string,
        evaluationId: string,
    ): CancelablePromise<Evaluation> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/evaluations/{evaluation_id}',
            path: {
                'project_id': projectId,
                'evaluation_id': evaluationId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an existing evaluation
     * @param projectId project uuid
     * @param evaluationId evaluation uuid
     * @param evaluation evaluation payload
     * @param key api key
     * @returns Evaluation OK
     * @throws ApiError
     */
    public putProjectsEvaluations(
        projectId: string,
        evaluationId: string,
        evaluation: Evaluation,
        key?: string,
    ): CancelablePromise<Evaluation> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/evaluations/{evaluation_id}',
            path: {
                'project_id': projectId,
                'evaluation_id': evaluationId,
            },
            query: {
                'key': key,
            },
            body: evaluation,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * deletes an evaluation
     * @param projectId project uuid
     * @param evaluationId evaluation uuid
     * @param key api key
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public deleteProjectsEvaluations(
        projectId: string,
        evaluationId: string,
        key?: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/evaluations/{evaluation_id}',
            path: {
                'project_id': projectId,
                'evaluation_id': evaluationId,
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
     * lists evaluations for a single instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @returns Evaluation OK
     * @throws ApiError
     */
    public getProjectsInstrumentsEvaluations(
        projectId: string,
        instrumentId: string,
    ): CancelablePromise<Array<Evaluation>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/projects/{project_id}/instruments/{instrument_id}/evaluations',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
