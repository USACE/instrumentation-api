
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
);

GRANT SELECT ON
    v_datalogger
TO instrumentation_reader;
