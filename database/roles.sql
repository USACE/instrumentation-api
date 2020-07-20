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
    instrument_group,
    instrument_group_instruments,
    instrument_note,
    instrument_status,
    instrument_type,
    instrument_zreference,
    measure,
    parameter,
    project,
    status,
    timeseries,
    timeseries_measurement,
    unit,
    unit_family,
    zreference_datum
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    instrument,
    instrument_group,
    instrument_group_instruments,
    instrument_note,
    instrument_status,
    instrument_type,
    instrument_zreference,
    measure,
    parameter,
    project,
    status,
    timeseries,
    timeseries_measurement,
    unit,
    unit_family,
    zreference_datum
TO instrumentation_writer;

-- Role postgis_reader
GRANT SELECT ON geometry_columns TO postgis_reader;
GRANT SELECT ON geography_columns TO postgis_reader;
GRANT SELECT ON spatial_ref_sys TO postgis_reader;
-- Grant Permissions to instrument_user
GRANT postgis_reader TO instrumentation_user;
GRANT instrumentation_reader TO instrumentation_user;
GRANT instrumentation_writer TO instrumentation_user;
