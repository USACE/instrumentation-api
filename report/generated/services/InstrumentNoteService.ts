/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfig } from '../models/AlertConfig';
import type { InstrumentNote } from '../models/InstrumentNote';
import type { InstrumentNoteCollection } from '../models/InstrumentNoteCollection';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class InstrumentNoteService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * gets all instrument notes
     * @returns InstrumentNote OK
     * @throws ApiError
     */
    public getInstrumentsNotes(): CancelablePromise<Array<InstrumentNote>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/notes',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * creates instrument notes
     * @param instrumentNote instrument note collection payload
     * @param key api key
     * @returns InstrumentNote OK
     * @throws ApiError
     */
    public postInstrumentsNotes(
        instrumentNote: InstrumentNoteCollection,
        key?: string,
    ): CancelablePromise<Array<InstrumentNote>> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/instruments/notes',
            query: {
                'key': key,
            },
            body: instrumentNote,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets a single instrument note by id
     * @param noteId note uuid
     * @returns InstrumentNote OK
     * @throws ApiError
     */
    public getInstrumentsNotes1(
        noteId: string,
    ): CancelablePromise<InstrumentNote> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/notes/{note_id}',
            path: {
                'note_id': noteId,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * updates an instrument note by id
     * @param noteId note uuid
     * @param instrumentNote instrument note collection payload
     * @param key api key
     * @returns AlertConfig OK
     * @throws ApiError
     */
    public putInstrumentsNotes(
        noteId: string,
        instrumentNote: InstrumentNote,
        key?: string,
    ): CancelablePromise<Array<AlertConfig>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/instruments/notes/{note_id}',
            path: {
                'note_id': noteId,
            },
            query: {
                'key': key,
            },
            body: instrumentNote,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * gets instrument notes for a single instrument
     * @param instrumentId instrument uuid
     * @returns InstrumentNote OK
     * @throws ApiError
     */
    public getInstrumentsNotes2(
        instrumentId: string,
    ): CancelablePromise<Array<InstrumentNote>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/instruments/{instrument_id}/notes',
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
     * deletes an instrument note
     * @param instrumentId instrument uuid
     * @param noteId note uuid
     * @param key api key
     * @returns any OK
     * @throws ApiError
     */
    public deleteInstrumentsNotes(
        instrumentId: string,
        noteId: string,
        key?: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/instruments/{instrument_id}/notes/{note_id}',
            path: {
                'instrument_id': instrumentId,
                'note_id': noteId,
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
