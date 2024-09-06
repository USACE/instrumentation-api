-- +goose up
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

ALTER ROLE instrumentation_user SET search_path TO midas,topology,public;

ALTER ROLE instrumentation_user SET intervalstyle TO 'iso_8601';

ALTER ROLE instrumentation_user SET statement_timeout TO '55s';

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

DO $$
DECLARE drop_views_query text;
BEGIN
    SELECT 'DROP VIEW ' || string_agg (table_name, ', ') || ' CASCADE;'
    FROM information_schema.views
    INTO drop_views_query
    WHERE table_schema = 'midas'
    AND table_name LIKE 'v_%';

    IF (drop_views_query IS NULL) THEN
        RAISE NOTICE 'not dropping views on schema midas -- no views found to drop';
    ELSE
        EXECUTE drop_views_query;
    END IF;
END
$$;
