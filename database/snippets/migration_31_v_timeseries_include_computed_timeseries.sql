DROP VIEW v_timeseries_latest;
DROP VIEW v_timeseries;

CREATE OR REPLACE VIEW v_timeseries AS (
    WITH ts_stored_and_computed AS (
        SELECT id,
            slug,
            name,
            instrument_id,
            parameter_id,
            unit_id,
            false AS is_computed
        FROM timeseries
        UNION
        SELECT formula_id AS id,
            'formula'  AS slug,
            'Formula'  AS name,
            id         AS instrument_id,
            formula_parameter_id AS parameter_id,
            formula_unit_id      AS unit_id,
            true AS is_computed
        FROM instrument
        WHERE NOT deleted AND formula IS NOT NULL
    )
    SELECT t.id AS id,
        t.slug AS slug,
        t.name AS name,
        t.is_computed AS is_computed,
        i.slug || '.' || t.slug AS variable,
        j.id AS project_id,
        j.slug AS project_slug,
        j.name AS project,
        i.id AS instrument_id,
        i.slug AS instrument_slug,
        i.name AS instrument,
        p.id AS parameter_id,
        p.name AS parameter,
        u.id AS unit_id,
        u.name AS unit
    FROM ts_stored_and_computed t
    INNER JOIN instrument i ON i.id = t.instrument_id AND NOT i.deleted
    INNER JOIN project j ON j.id = i.project_id
    INNER JOIN parameter p ON p.id = t.parameter_id
    INNER JOIN unit U ON u.id = t.unit_id
);

-- v_timeseries_latest; same as v_timeseries, joined with latest times and values
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

GRANT SELECT ON v_timeseries, v_timeseries_latest TO instrumentation_reader;