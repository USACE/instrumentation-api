-- v_datalogger
CREATE OR REPLACE VIEW v_datalogger AS (
    SELECT * FROM datalogger WHERE NOT deleted
);

-- v_datalogger_preview
CREATE OR REPLACE VIEW v_datalogger_preview AS (
    SELECT * FROM datalogger_preview
);

-- v_datalogger_equivalency_table
CREATE OR REPLACE VIEW v_datalogger_equivalency_table AS (
    SELECT
        datalogger_id,
        field_name,
        display_name,
        instrument_id,
        timeseries_id
    FROM datalogger_equivalency_table
    WHERE NOT datalogger_deleted
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
    v_datalogger_equivalency_table,
    v_datalogger_hash
TO instrumentation_reader;
