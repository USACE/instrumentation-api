/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IDSlugName } from './IDSlugName';
import type { PlotConfigSettings } from './PlotConfigSettings';
export type ReportConfig = {
    after?: string;
    before?: string;
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    description?: string;
    id?: string;
    name?: string;
    override_plot_config_settings?: PlotConfigSettings;
    plot_configs?: Array<IDSlugName>;
    project_id?: string;
    project_name?: string;
    slug?: string;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
};
