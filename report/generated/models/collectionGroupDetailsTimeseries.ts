/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Measurement } from './Measurement';
export type collectionGroupDetailsTimeseries = {
    id?: string;
    instrument?: string;
    instrument_id?: string;
    instrument_slug?: string;
    is_computed?: boolean;
    latest_time?: string;
    latest_value?: number;
    name?: string;
    parameter?: string;
    parameter_id?: string;
    slug?: string;
    unit?: string;
    unit_id?: string;
    values?: Array<Measurement>;
    variable?: string;
};

