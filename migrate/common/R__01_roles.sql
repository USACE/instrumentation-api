-- ${flyway:timestamp}
-- Users and Roles for HHD Instrumentation Webapp

-- User instrumentation_user
-- Note: Substitute real password for 'password'
DO $$
BEGIN
    CREATE USER instrumentation_user WITH ENCRYPTED PASSWORD 'password';
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating role instrumentation_user -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE instrumentation_reader;
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating role instrumentation_reader -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE instrumentation_writer;
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating role instrumentation_writer -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE postgis_reader;
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating role postgis_reader -- it already exists';
END
$$;

-- Set Search Path
ALTER ROLE instrumentation_user SET search_path TO midas,topology,public;

-- Set intervalstyle
ALTER ROLE instrumentation_user SET intervalstyle TO 'iso_8601';

-- Set statement timeout
ALTER ROLE instrumentation_user SET statement_timeout TO '55s';

-- Grant Schema Usage to instrumentation_user
GRANT USAGE ON SCHEMA midas TO instrumentation_user;

--------------------------------------------------------------------------
-- NOTE: IF USERS ALREADY EXIST ON DATABASE, JUST RUN FROM THIS POINT DOWN
--------------------------------------------------------------------------

-- Role instrumentation_reader
-- Tables specific to instrumentation app
GRANT SELECT ON
    instrument,
    instrument_telemetry,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type,
    alert,
    alert_read,
    alert_config,
    alert_email_subscription,
    alert_profile_subscription,
    calculation,
    collection_group,
    collection_group_timeseries,
    config,
    profile,
    profile_project_roles,
    profile_token,
    role,
    email,
    heartbeat,
    instrument_constants,
    instrument_group,
    instrument_group_instruments,
    instrument_note,
    instrument_status,
    instrument_type,
    measure,
    parameter,
    plot_configuration,
    plot_configuration_settings,
    project,
    project_instrument,
    status,
    timeseries,
    timeseries_measurement,
    timeseries_notes,
    unit,
    unit_family,
    inclinometer_measurement,
    aware_platform_parameter_enabled,
    aware_platform,
    aware_parameter,
    datalogger,
    datalogger_hash,
    datalogger_table,
    datalogger_preview,
    datalogger_equivalency_table,
    datalogger_model,
    datalogger_error,
    evaluation,
    evaluation_instrument,
    submittal_status,
    alert_type,
    alert_config_instrument,
    submittal,
    saa_opts,
    saa_segment,
    ipi_opts,
    ipi_segment,
    report_config,
    report_config_plot_config,
    plot_configuration_timeseries_trace,
    plot_configuration_custom_shape
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    alert,
    alert_read,
    alert_config,
    alert_email_subscription,
    alert_profile_subscription,
    calculation,
    collection_group,
    collection_group_timeseries,
    config,
    instrument,
    instrument_telemetry,
    plot_configuration,
    plot_configuration_settings,
    profile,
    profile_project_roles,
    profile_token,
    email,
    heartbeat,
    instrument_constants,
    instrument_group,
    instrument_group_instruments,
    instrument_note,
    instrument_status,
    instrument_type,
    measure,
    parameter,
    project,
    project_instrument,
    role,
    status,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type,
    timeseries,
    timeseries_measurement,
    timeseries_notes,
    unit,
    unit_family,
    inclinometer_measurement,
    aware_platform_parameter_enabled,
    aware_platform,
    aware_parameter,
    datalogger,
    datalogger_hash,
    datalogger_table,
    datalogger_preview,
    datalogger_equivalency_table,
    datalogger_model,
    datalogger_error,
    evaluation,
    evaluation_instrument,
    alert_config_instrument,
    submittal,
    saa_opts,
    saa_segment,
    ipi_opts,
    ipi_segment,
    report_config,
    report_config_plot_config,
    plot_configuration_timeseries_trace,
    plot_configuration_custom_shape
TO instrumentation_writer;

-- Role postgis_reader
GRANT SELECT ON geometry_columns TO postgis_reader;
GRANT SELECT ON geography_columns TO postgis_reader;
GRANT SELECT ON spatial_ref_sys TO postgis_reader;

-- Grant Permissions to instrument_user
GRANT postgis_reader TO instrumentation_user;
GRANT instrumentation_reader TO instrumentation_user;
GRANT instrumentation_writer TO instrumentation_user;

-- Drop all views to be recreated
DROP VIEW IF EXISTS v_alert CASCADE;
DROP VIEW IF EXISTS v_alert_check_evaluation_submittal CASCADE;
DROP VIEW IF EXISTS v_alert_check_measurement_submittal CASCADE;
DROP VIEW IF EXISTS v_alert_config CASCADE;
DROP VIEW IF EXISTS v_aware_platform_parameter_enabled CASCADE;
DROP VIEW IF EXISTS v_datalogger CASCADE;
DROP VIEW IF EXISTS v_datalogger_equivalency_table CASCADE;
DROP VIEW IF EXISTS v_datalogger_hash CASCADE;
DROP VIEW IF EXISTS v_datalogger_preview CASCADE;
DROP VIEW IF EXISTS v_district CASCADE;
DROP VIEW IF EXISTS v_district_rollup CASCADE;
DROP VIEW IF EXISTS v_domain CASCADE;
DROP VIEW IF EXISTS v_domain_group CASCADE;
DROP VIEW IF EXISTS v_email_autocomplete CASCADE;
DROP VIEW IF EXISTS v_evaluation CASCADE;
DROP VIEW IF EXISTS v_instrument CASCADE;
DROP VIEW IF EXISTS v_instrument_group CASCADE;
DROP VIEW IF EXISTS v_instrument_telemetry CASCADE;
DROP VIEW IF EXISTS v_ipi_measurement CASCADE;
DROP VIEW IF EXISTS v_ipi_segment CASCADE;
DROP VIEW IF EXISTS v_plot_configuration CASCADE;
DROP VIEW IF EXISTS v_profile CASCADE;
DROP VIEW IF EXISTS v_profile_project_roles CASCADE;
DROP VIEW IF EXISTS v_project CASCADE;
DROP VIEW IF EXISTS v_saa_measurement CASCADE;
DROP VIEW IF EXISTS v_saa_segment CASCADE;
DROP VIEW IF EXISTS v_submittal CASCADE;
DROP VIEW IF EXISTS v_timeseries CASCADE;
DROP VIEW IF EXISTS v_timeseries_computed CASCADE;
DROP VIEW IF EXISTS v_timeseries_dependency CASCADE;
DROP VIEW IF EXISTS v_timeseries_project_map CASCADE;
DROP VIEW IF EXISTS v_timeseries_stored CASCADE;
DROP VIEW IF EXISTS v_unit CASCADE;
DROP VIEW IF EXISTS v_report_config CASCADE;
