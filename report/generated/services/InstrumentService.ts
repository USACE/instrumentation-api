/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IDSlugName } from '../models/IDSlugName';
import type { Instrument } from '../models/Instrument';
import type { InstrumentCount } from '../models/InstrumentCount';
import type { InstrumentProjectAssignments } from '../models/InstrumentProjectAssignments';
import type { InstrumentsValidation } from '../models/InstrumentsValidation';
import type { ProjectInstrumentAssignments } from '../models/ProjectInstrumentAssignments';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * lists all instruments
     * @returns Instrument OK
     * @throws ApiError
     */
    public getInstruments(): CancelablePromise<Array<Instrument>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets the total number of non deleted instruments in the system
     * @returns InstrumentCount OK
     * @throws ApiError
     */
    public getInstrumentsCount(): CancelablePromise<InstrumentCount> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/count',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single instrument by id
     * @param instrumentId instrument uuid
     * @returns Instrument OK
     * @throws ApiError
     */
    public getInstruments1(
        instrumentId: string,
    ): CancelablePromise<Instrument> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}',
            path: {
                'instrument_id': instrumentId,
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
     * @param projectId project id
     * @param instrumentId instrument id
     * @param instrument instrument collection payload
     * @param key api key
     * @returns IDSlugName OK
     * @throws ApiError
     */
    public postProjectsInstruments(
        projectId: string,
        instrumentId: string,
        instrument: Array<Instrument>,
        key?: string,
    ): CancelablePromise<Array<IDSlugName>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/instruments',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: instrument,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates multiple instrument assigments for a project
     * must be Project (or Application) Admin of all existing instrument projects and project to be assigned
     * @param projectId project uuid
     * @param instrumentIds instrument uuids
     * @param action valid values are 'assign' or 'unassign'
     * @param dryRun validate request without performing action
     * @returns InstrumentsValidation OK
     * @throws ApiError
     */
    public putProjectsInstrumentsAssignments(
        projectId: string,
        instrumentIds: ProjectInstrumentAssignments,
        action: string,
        dryRun?: string,
    ): CancelablePromise<InstrumentsValidation> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/instruments/assignments',
            path: {
                'project_id': projectId,
            },
            query: {
                'action': action,
                'dry_run': dryRun,
            },
            body: instrumentIds,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an existing instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param instrument instrument payload
     * @param key api key
     * @returns Instrument OK
     * @throws ApiError
     */
    public putProjectsInstruments(
        projectId: string,
        instrumentId: string,
        instrument: Instrument,
        key?: string,
    ): CancelablePromise<Instrument> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/instruments/{instrument_id}',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: instrument,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * soft deletes an instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteProjectsInstruments(
        projectId: string,
        instrumentId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/instruments/{instrument_id}',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
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
     * updates multiple project assignments for an instrument
     * must be Project (or Application) Admin of all existing instrument projects and project to be assigned
     * @param instrumentId instrument uuid
     * @param projectIds project uuids
     * @param action valid values are 'assign' or 'unassign'
     * @param dryRun validate request without performing action
     * @returns InstrumentsValidation OK
     * @throws ApiError
     */
    public putProjectsInstrumentsAssignments1(
        instrumentId: string,
        projectIds: InstrumentProjectAssignments,
        action: string,
        dryRun?: string,
    ): CancelablePromise<InstrumentsValidation> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/instruments/{instrument_id}/assignments',
            path: {
                'instrument_id': instrumentId,
            },
            query: {
                'action': action,
                'dry_run': dryRun,
            },
            body: projectIds,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * assigns an instrument to a project.
     * must be Project (or Application) Admin of all existing instrument projects and project to be assigned
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param dryRun validate request without performing action
     * @returns InstrumentsValidation OK
     * @throws ApiError
     */
    public postProjectsInstrumentsAssignments(
        projectId: string,
        instrumentId: string,
        dryRun?: string,
    ): CancelablePromise<InstrumentsValidation> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/projects/{project_id}/instruments/{instrument_id}/assignments',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'dry_run': dryRun,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * unassigns an instrument from a project.
     * must be Project Admin of project to be unassigned
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param action valid values are 'assign' or 'unassign'
     * @param dryRun validate request without performing action
     * @returns InstrumentsValidation OK
     * @throws ApiError
     */
    public deleteProjectsInstrumentsAssignments(
        projectId: string,
        instrumentId: string,
        action: string,
        dryRun?: string,
    ): CancelablePromise<InstrumentsValidation> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/projects/{project_id}/instruments/{instrument_id}/assignments',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'action': action,
                'dry_run': dryRun,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates the geometry of an instrument
     * @param projectId project uuid
     * @param instrumentId instrument uuid
     * @param instrument instrument payload
     * @param key api key
     * @returns Instrument OK
     * @throws ApiError
     */
    public putProjectsInstrumentsGeometry(
        projectId: string,
        instrumentId: string,
        instrument: Instrument,
        key?: string,
    ): CancelablePromise<Instrument> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/projects/{project_id}/instruments/{instrument_id}/geometry',
            path: {
                'project_id': projectId,
                'instrument_id': instrumentId,
            },
            query: {
                'key': key,
            },
            body: instrument,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
}
