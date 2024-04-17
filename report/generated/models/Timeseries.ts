/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Measurement } from './Measurement';
export type Timeseries = {
    id?: string;
    instrument?: string;
    instrument_id?: string;
    instrument_slug?: string;
    is_computed?: boolean;
    name?: string;
    parameter?: string;
    parameter_id?: string;
    slug?: string;
    unit?: string;
    unit_id?: string;
    values?: Array<Measurement>;
    variable?: string;
};

