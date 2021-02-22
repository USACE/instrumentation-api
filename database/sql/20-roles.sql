-- Users and Roles for HHD Instrumentation Webapp

-- User instrumentation_user
-- Note: Substitute real password for 'password'
CREATE USER instrumentation_user WITH ENCRYPTED PASSWORD 'password';
CREATE ROLE instrumentation_reader;
CREATE ROLE instrumentation_writer;
CREATE ROLE postgis_reader;

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
    collection_group,
    collection_group_timeseries,
    config,
    profile,
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
    plot_configuration,
    plot_configuration_timeseries,
    project,
    project_timeseries,
    status,
    timeseries,
    timeseries_measurement,
    unit,
    unit_family,
    v_instrument,
    v_project,
    v_timeseries,
    v_email_autocomplete,
    v_alert,
    v_timeseries_latest,
    v_timeseries_project_map,
    v_instrument_telemetry
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    alert,
    alert_read,
    alert_config,
    alert_email_subscription,
    alert_profile_subscription,
    collection_group,
    collection_group_timeseries,
    config,
    instrument,
    instrument_telemetry,
    plot_configuration,
    plot_configuration_timeseries,
    profile,
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
    project_timeseries,
    status,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type,
    timeseries,
    timeseries_measurement,
    unit,
    unit_family
TO instrumentation_writer;

-- Role postgis_reader
GRANT SELECT ON geometry_columns TO postgis_reader;
GRANT SELECT ON geography_columns TO postgis_reader;
GRANT SELECT ON spatial_ref_sys TO postgis_reader;
-- Grant Permissions to instrument_user
GRANT postgis_reader TO instrumentation_user;
GRANT instrumentation_reader TO instrumentation_user;
GRANT instrumentation_writer TO instrumentation_user;
