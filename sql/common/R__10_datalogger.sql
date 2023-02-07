-- v_datalogger
CREATE OR REPLACE VIEW v_datalogger AS (
    SELECT * FROM datalogger WHERE NOT deleted
);

-- v_datalogger_preview
CREATE OR REPLACE VIEW v_datalogger_preview AS (
    SELECT * FROM datalogger_preview
);

-- v_datalogger_field_instrument_timeseries
CREATE OR REPLACE VIEW v_datalogger_field_instrument_timeseries AS (
    SELECT
        datalogger_id,
        json_agg(json_build_object(
                'field_name', field_name,
                'display_name', display_name,
                'instrument_id', instrument_id,
                'timeseries_id', timeseries_id
        ) ORDER BY field_name ASC)::text AS "rows"
    FROM datalogger_field_instrument_timeseries
    GROUP BY datalogger_id
);

-- v_datalogger_hash
CREATE OR REPLACE VIEW v_datalogger_hash AS (
    SELECT
        dh.datalogger_id AS datalogger_id,
        dh.hash AS "hash",
        dl.sn AS sn
    FROM datalogger_hash dh
    INNER JOIN datalogger dl ON dh.datalogger_id = dl.id
);

GRANT SELECT ON
    v_datalogger,
    v_datalogger_preview,
    v_datalogger_field_instrument_timeseries
TO instrumentation_reader;
