/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AlertConfigInstrument } from './AlertConfigInstrument';
import type { EmailAutocompleteResult } from './EmailAutocompleteResult';
export type AlertConfig = {
    alert_email_subscriptions?: Array<EmailAutocompleteResult>;
    alert_type?: string;
    alert_type_id?: string;
    body?: string;
    create_date?: string;
    creator_id?: string;
    creator_username?: string;
    id?: string;
    instruments?: Array<AlertConfigInstrument>;
    last_checked?: string;
    last_reminded?: string;
    mute_consecutive_alerts?: boolean;
    name?: string;
    project_id?: string;
    project_name?: string;
    remind_interval?: string;
    schedule_interval?: string;
    start_date?: string;
    update_date?: string;
    updater_id?: string;
    updater_username?: string;
    warning_interval?: string;
};

