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

-- grant permissions for instrumentation_user (api user)
GRANT USAGE ON SCHEMA midas TO instrumentation_user;
GRANT SELECT ON ALL TABLES IN SCHEMA midas TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON ALL TABLES IN SCHEMA midas TO instrumentation_writer;

REVOKE SELECT ON flyway_schema_history FROM instrumentation_reader;
REVOKE INSERT,UPDATE,DELETE ON flyway_schema_history FROM instrumentation_writer;

GRANT SELECT ON
    geometry_columns,
    geography_columns,
    spatial_ref_sys
TO postgis_reader;

GRANT
    postgis_reader,
    instrumentation_reader,
    instrumentation_writer
TO instrumentation_user;

-- Drop all views to be recreated
SELECT 'DROP VIEW ' || string_agg (table_name, ', ') || ' CASCADE;'
FROM information_schema.views
WHERE table_schema = 'midas'
AND table_name LIKE 'v_%';
