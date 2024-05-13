/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { DataloggerTable } from './DataloggerTable';
export type Datalogger = {
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    errors?: Array<string>;
    id?: string;
    model?: string;
    model_id?: string;
    name?: string;
    project_id?: string;
    slug?: string;
    sn?: string;
    tables?: Array<DataloggerTable>;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
};

