/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Geometry } from './Geometry';
import type { IDSlugName } from './IDSlugName';
import type { Opts } from './Opts';
export type Instrument = {
    alert_configs?: Array<string>;
    aware_id?: string;
    constants?: Array<string>;
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    geometry?: Geometry;
    groups?: Array<string>;
    icon?: string;
    id?: string;
    name?: string;
    nid_id?: string;
    offset?: number;
    opts?: Opts;
    projects?: Array<IDSlugName>;
    slug?: string;
    station?: number;
    status?: string;
    status_id?: string;
    status_time?: string;
    type?: string;
    type_id?: string;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
    usgs_id?: string;
};

