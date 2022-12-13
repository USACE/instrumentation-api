-- midas.timeseries and midas.calculation are identical schemas in the database, except that the latter has a contents field.
-- all timeseries should be stored in the same table (computed and stored).

set search_path = "$user", midas, public, topology;

BEGIN;

-- migrate existing calculation table data to timeseries table
CREATE TABLE IF NOT EXISTS tmp_calculation (
    id UUID,
    instrument_id UUID,
    parameter_id UUID,
    unit_id UUID,
    slug VARCHAR(255),
    name VARCHAR(255),
    contents VARCHAR
);

INSERT INTO tmp_calculation (id, instrument_id, parameter_id, unit_id, slug, name, contents)
SELECT id, instrument_id, parameter_id, unit_id, slug, name, contents FROM calculation;

-- remove calculation table and all depending views:
--      v_timeseries
--      v_timeseries_latest
--      v_timeseries_dependency
DROP TABLE calculation CASCADE;

-- derive timeseries fom tmp_calculation
INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name)
SELECT
    id,
    instrument_id,
    COALESCE(parameter_id, '2b7f96e1-820f-4f61-ba8f-861640af6232'::uuid) AS parameter_id,
    COALESCE(unit_id, '4a999277-4cf5-4282-93ce-23b33c65e2c8'::uuid) AS unit_id,
    slug,
    name
FROM tmp_calculation;

-- create new calculation table referencing computed timeseries
-- calculation
CREATE TABLE IF NOT EXISTS calculation (
	timeseries_id UUID UNIQUE NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
	contents VARCHAR
);

-- derive calculation form tmp_calculation
INSERT INTO calculation (timeseries_id, contents)
SELECT id, contents FROM tmp_calculation;

-- recreate views
-- stored and computed timeseries
CREATE OR REPLACE VIEW v_timeseries AS (
    WITH ts_stored_and_computed AS (
        SELECT
            id,
            slug,
            name,
            instrument_id,
            parameter_id,
            unit_id,
            (SELECT id IN (SELECT timeseries_id FROM calculation)) AS is_computed
        FROM timeseries
    )
    SELECT t.id                 AS id,
        t.slug                  AS slug,
        t.name                  AS name,
        t.is_computed           AS is_computed,
        i.slug || '.' || t.slug AS variable,
        j.id                    AS project_id,
        j.slug                  AS project_slug,
        j.name                  AS project,
        i.id                    AS instrument_id,
        i.slug                  AS instrument_slug,
        i.name                  AS instrument,
        p.id                    AS parameter_id,
        p.name                  AS parameter,
        u.id                    AS unit_id,
        u.name                  AS unit
    FROM ts_stored_and_computed t
    INNER JOIN instrument i ON i.id = t.instrument_id
    INNER JOIN project j ON j.id = i.project_id
    INNER JOIN parameter p ON p.id = t.parameter_id
    INNER JOIN unit U ON u.id = t.unit_id
);

-- stored timeseries only
CREATE OR REPLACE VIEW v_timeseries_latest AS (
    SELECT t.*,
       m.time AS latest_time,
	   m.value AS latest_value
    FROM v_timeseries t
    LEFT JOIN (
	    SELECT DISTINCT ON (timeseries_id) timeseries_id, time, value
	    FROM timeseries_measurement
	    ORDER BY timeseries_id, time DESC
    ) m ON t.id = m.timeseries_id
    WHERE NOT t.is_computed
);

-- computed timeseries and stored dependency timeseries
CREATE OR REPLACE VIEW v_timeseries_dependency AS (
    WITH variable_tsid_map AS (
	    SELECT a.id AS timeseries_id,
               b.slug || '.' || a.slug AS variable
	    FROM timeseries a
	    LEFT JOIN instrument b ON b.id = a.instrument_id
        WHERE a.id NOT IN (SELECT timeseries_id FROM calculation)
    )
    SELECT i.instrument_id   AS instrument_id,
           i.formula_id      AS timeseries_id,
           i.parsed_variable AS parsed_variable,
           m.timeseries_id   AS dependency_timeseries_id
    FROM (
        SELECT instrument_id,
            id AS formula_id,
            (regexp_matches(contents, '\[(.*?)\]', 'g'))[1] AS parsed_variable
        FROM timeseries t
        INNER JOIN calculation cc ON cc.timeseries_id = t.id
    ) i
    LEFT JOIN variable_tsid_map m ON m.variable = i.parsed_variable
);

CREATE OR REPLACE VIEW v_timeseries_stored AS (
    SELECT * FROM timeseries WHERE id NOT IN (SELECT timeseries_id FROM calculation)
);

CREATE OR REPLACE VIEW v_timeseries_computed AS (
    SELECT
        ts.*,
        cc.contents AS contents
    FROM timeseries ts
    LEFT JOIN calculation cc ON ts.id = cc.timeseries_id
    WHERE id IN (SELECT timeseries_id FROM calculation)
);

-- cleanup
DROP TABLE tmp_calculation;

-- roles
GRANT SELECT ON
    calculation,
    v_timeseries,
    v_timeseries_latest,
    v_timeseries_dependency,
    v_timeseries_stored,
    v_timeseries_computed
TO instrumentation_reader;

GRANT INSERT,UPDATE,DELETE ON calculation TO instrumentation_writer;

COMMIT;
