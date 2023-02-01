-- v_telemetry
CREATE OR REPLACE VIEW v_instrument_telemetry AS (
    SELECT a.id,
           a.instrument_id AS instrument_id,
           b.id AS telemetry_type_id,
           b.slug AS telemetry_type_slug,
           b.name AS telemetry_type_name
    FROM instrument_telemetry a
    INNER JOIN telemetry_type b ON b.id = a.telemetry_type_id
    LEFT JOIN telemetry_goes tg ON a.telemetry_id = tg.id
    LEFT JOIN telemetry_iridium ti ON a.telemetry_id = ti.id
);

GRANT SELECT ON
    v_instrument_telemetry
TO instrumentation_reader;
