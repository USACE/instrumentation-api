
-- v_datalogger
CREATE OR REPLACE VIEW v_datalogger AS (
    SELECT 
        id,
        sn,
        project_id,
        creator,
        create_date,
        name,
        slug,
        model
    FROM datalogger
    WHERE NOT deleted
);

CREATE OR REPLACE VIEW v_datalogger_field_instrument_timeseries AS (
    SELECT
        datalogger_id,
        json_agg(json_build_object(
                'field_name', field_name,
                'display_name', display_name,
                'instrument_id', instrument_id,
                'timeseries_id', timeseries_id
        ) ORDER BY field_name ASC)::text AS 'rows'
        GROUP BY datalogger_id
)

GRANT SELECT ON
    v_datalogger
TO instrumentation_reader;
