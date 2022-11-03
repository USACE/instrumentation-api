--
-- Adds formula names to instruments with a calculated representation.
-- Useful for changing the display name of the instrument when being
-- batch-plotted, as well as adding multiple formulas associated to
-- a single instrument.
--

set search_path = "$user", midas, public, topology;

BEGIN;

CREATE TABLE IF NOT EXISTS calculation (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),

    instrument_id UUID NOT NULL REFERENCES instrument (id),
    parameter_id UUID REFERENCES parameter (id),
    unit_id UUID REFERENCES unit (id),

    slug VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    contents VARCHAR
);

GRANT SELECT ON calculation TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON calculation TO instrumentation_writer;

INSERT INTO calculation (id, instrument_id, parameter_id, unit_id, slug, name, contents)
SELECT
    I.formula_id,
    I.id,
    I.formula_parameter_id,
    I.formula_unit_id,
    I.slug || '.formula',
    I.name || ' formula',
    I.formula
FROM instrument I
ON CONFLICT (id) DO NOTHING;

DROP VIEW IF EXISTS v_instrument;
DROP VIEW IF EXISTS v_timeseries_latest;
DROP VIEW IF EXISTS v_timeseries_dependency;
DROP VIEW IF EXISTS v_timeseries;

CREATE OR REPLACE VIEW v_instrument AS (
    SELECT I.id,
        I.deleted,
        S.status_id,
        S.status,
        S.status_time,
        I.slug,
        I.name,
        I.type_id,
        T.name AS type,
        ST_AsBinary(I.geometry) AS geometry,
        I.station,
        I.station_offset,
        I.creator,
        I.create_date,
        I.updater,
        I.update_date,
        I.project_id,
        I.nid_id,
        I.usgs_id,
        TEL.telemetry AS telemetry,
        COALESCE(C.constants, '{}') AS constants,
        COALESCE(G.groups, '{}') AS groups,
        COALESCE(A.alert_configs, '{}') AS alert_configs
    FROM instrument I
    INNER JOIN instrument_type T ON T.id = I.type_id
    INNER JOIN (
        SELECT DISTINCT ON (instrument_id) instrument_id,
            a.time AS status_time,
            a.status_id AS status_id,
            d.name AS status
        FROM instrument_status a
        INNER JOIN status d ON d.id = a.status_id
        WHERE a.time <= now()
        ORDER BY instrument_id, a.time DESC
    ) S ON S.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(timeseries_id) as constants,
            instrument_id
        FROM instrument_constants
        GROUP BY instrument_id
    ) C on C.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(instrument_group_id) as groups,
            instrument_id
        FROM instrument_group_instruments
        GROUP BY instrument_id
    ) G on G.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(id) as alert_configs,
            instrument_id
        FROM alert_config
        GROUP BY instrument_id
    ) A on A.instrument_id = I.id
    LEFT JOIN (
        SELECT instrument_id,
                json_agg(
                    json_build_object(
                        'id', v.id,
                        'slug', v.telemetry_type_slug,
                        'name', v.telemetry_type_name
                    )
                ) AS telemetry
        FROM v_instrument_telemetry v
        GROUP BY instrument_id
    ) TEL ON TEL.instrument_id = I.id
);

CREATE OR REPLACE VIEW v_timeseries AS (
    WITH ts_stored_and_computed AS (
        SELECT id,
            slug,
            name,
            instrument_id,
            parameter_id,
            unit_id,
            false                AS is_computed
        FROM timeseries
        UNION
        SELECT CC.id,
            'formula_' || CC.name AS slug,
            CC.name,
            CC.instrument_id,
            CC.parameter_id,
            CC.unit_id,
            'true'                AS is_computed
        FROM instrument II
        LEFT JOIN calculation CC
        ON CC.instrument_id = II.id
        WHERE NOT II.deleted
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
);

CREATE OR REPLACE VIEW v_timeseries_dependency AS (
    WITH variable_tsid_map AS (
	    SELECT a.id AS timeseries_id,
               b.slug || '.' || a.slug AS variable
	    FROM timeseries a
	    LEFT JOIN instrument b ON b.id = a.instrument_id
    )
    SELECT i.instrument_id   AS instrument_id,
           i.formula_id      AS timeseries_id,
           i.parsed_variable AS parsed_variable,
           m.timeseries_id   AS dependency_timeseries_id
    FROM (
        SELECT instrument_id,
            id AS formula_id,
            (regexp_matches(contents, '\[(.*?)\]', 'g'))[1] AS parsed_variable
        FROM calculation
    ) i
    LEFT JOIN variable_tsid_map m ON m.variable = i.parsed_variable
);

ALTER TABLE instrument
    DROP COLUMN IF EXISTS formula_id,
    DROP COLUMN IF EXISTS formula,
    DROP COLUMN IF EXISTS formula_parameter_id,
    DROP COLUMN IF EXISTS formula_unit_id;

GRANT SELECT ON
    v_instrument,
    v_timeseries,
    v_timeseries_latest,
    v_timeseries_dependency
TO instrumentation_reader;

COMMIT;
