-- ${flyway:timestamp}
CREATE VIEW v_timeseries AS (
    SELECT
        t.id,
        t.slug,
        t.name,
        t.type,
        -- is_computed should be deprecated now that timeseries type enum is included
        CASE WHEN t.type = 'computed' THEN true ELSE false END AS is_computed,
        i.slug || '.' || t.slug AS variable,
        i.id AS instrument_id,
        i.slug AS instrument_slug,
        i.name AS instrument,
        p.id AS parameter_id,
        p.name AS parameter,
        u.id AS unit_id,
        u.name AS unit
    FROM timeseries t
    INNER JOIN instrument i ON i.id = t.instrument_id
    INNER JOIN parameter p ON p.id = t.parameter_id
    INNER JOIN unit u ON u.id = t.unit_id
);

CREATE OR REPLACE VIEW v_timeseries_dependency AS (
    WITH variable_tsid_map AS (
	    SELECT
            a.id AS timeseries_id,
            b.slug || '.' || a.slug AS variable
	    FROM timeseries a
	    LEFT JOIN instrument b ON b.id = a.instrument_id
            WHERE a.id NOT IN (SELECT c.timeseries_id FROM calculation c)
    )
    -- id references computed timeseries
    -- dependency_timeseries_id references timeseries used to caclulate computed timeseries
    SELECT
        i.id              AS id,
        i.instrument_id   AS instrument_id,
        i.parsed_variable AS parsed_variable,
        m.timeseries_id   AS dependency_timeseries_id
    FROM (
        SELECT
            id,
            instrument_id,
            (regexp_matches(contents, '\[(.*?)\]', 'g'))[1] AS parsed_variable
        FROM timeseries t
        INNER JOIN calculation cc ON cc.timeseries_id = t.id
    ) i
    LEFT JOIN variable_tsid_map m ON m.variable = i.parsed_variable
);

CREATE OR REPLACE VIEW v_timeseries_project_map AS (
    SELECT
        t.id AS timeseries_id,
        pi.project_id AS project_id
    FROM timeseries t
    LEFT JOIN instrument i ON t.instrument_id = i.id
    LEFT JOIN project_instrument pi ON pi.instrument_id = i.id
);

CREATE VIEW v_timeseries_stored AS (
    SELECT * FROM timeseries
    WHERE type = 'standard'
    OR type = 'constant'
);

CREATE OR REPLACE VIEW v_timeseries_computed AS (
    SELECT
        ts.*,
        cc.contents AS contents
    FROM timeseries ts
    INNER JOIN calculation cc ON ts.id = cc.timeseries_id
);

CREATE OR REPLACE VIEW v_timeseries_cwms AS (
    SELECT
        ts.*,
        tc.cwms_timeseries_id,
        tc.cwms_office_id
    FROM v_timeseries ts
    INNER JOIN timeseries_cwms tc ON ts.id = tc.timeseries_id
);

GRANT SELECT ON
    v_timeseries,
    v_timeseries_dependency,
    v_timeseries_stored,
    v_timeseries_computed,
    v_timeseries_cwms,
    v_timeseries_project_map
TO instrumentation_reader;
