-- ${flyway:timestamp}
CREATE VIEW v_timeseries AS (
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
        i.id                    AS instrument_id,
        i.slug                  AS instrument_slug,
        i.name                  AS instrument,
        p.id                    AS parameter_id,
        p.name                  AS parameter,
        u.id                    AS unit_id,
        u.name                  AS unit
    FROM ts_stored_and_computed t
    INNER JOIN instrument i ON i.id = t.instrument_id
    INNER JOIN parameter p ON p.id = t.parameter_id
    INNER JOIN unit U ON u.id = t.unit_id
);

CREATE VIEW v_timeseries_dependency AS (
    WITH variable_tsid_map AS (
	    SELECT
            a.id AS timeseries_id,
            b.slug || '.' || a.slug AS variable
	    FROM timeseries a
	    LEFT JOIN instrument b ON b.id = a.instrument_id
        WHERE a.id NOT IN (SELECT timeseries_id FROM calculation)
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

CREATE VIEW v_timeseries_project_map AS (
    SELECT
        t.id AS timeseries_id,
        pi.project_id AS project_id
    FROM timeseries t
    LEFT JOIN instrument i ON t.instrument_id = i.id
    LEFT JOIN project_instrument pi ON pi.instrument_id = i.id
);

CREATE VIEW v_timeseries_stored AS (
    SELECT * FROM timeseries WHERE id NOT IN (SELECT timeseries_id FROM calculation)
);

CREATE VIEW v_timeseries_computed AS (
    SELECT
        ts.*,
        cc.contents AS contents
    FROM timeseries ts
    LEFT JOIN calculation cc ON ts.id = cc.timeseries_id
    WHERE id IN (SELECT timeseries_id FROM calculation)
);

GRANT SELECT ON
    v_timeseries,
    v_timeseries_dependency,
    v_timeseries_stored,
    v_timeseries_computed,
    v_timeseries_project_map
TO instrumentation_reader;
