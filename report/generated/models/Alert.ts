/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfigInstrument } from './AlertConfigInstrument';
export type Alert = {
    alert_config_id?: string;
    body?: string;
    create_date?: string;
    id?: string;
    instruments?: Array<AlertConfigInstrument>;
    name?: string;
    project_id?: string;
    project_name?: string;
    read?: boolean;
};

