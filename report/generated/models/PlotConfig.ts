/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IDSlugName } from './IDSlugName';
import type { PlotConfigDisplay } from './PlotConfigDisplay';
export type PlotConfig = {
    auto_range?: boolean;
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    date_range?: string;
    display?: PlotConfigDisplay;
    id?: string;
    name?: string;
    project_id?: string;
    report_configs?: Array<IDSlugName>;
    show_comments?: boolean;
    show_masked?: boolean;
    show_nonvalidated?: boolean;
    slug?: string;
    threshold?: number;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
};

