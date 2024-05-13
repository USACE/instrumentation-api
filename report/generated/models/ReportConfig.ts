/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IDSlugName } from './IDSlugName';
import type { ReportConfigGlobalOverrides } from './ReportConfigGlobalOverrides';
export type ReportConfig = {
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    description?: string;
    global_overrides?: ReportConfigGlobalOverrides;
    id?: string;
    name?: string;
    plot_configs?: Array<IDSlugName>;
    project_id?: string;
    project_name?: string;
    slug?: string;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
};

