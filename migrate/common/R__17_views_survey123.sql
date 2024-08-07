-- ${flyway:timestamp}
CREATE OR REPLACE VIEW v_survey123 AS (
  SELECT
    sv.id, 
    sv.project_id,
    sv.name,
    sv.slug,
    sv.create_date,
    sv.update_date,
    sv.creator,
    p1.username AS creator_username,
    sv.updater,
    p2.username AS updater_username,
    COALESCE(f.fields, '[]'::json)::text AS fields,
    COALESCE(er.errors, '{}') AS errors
  FROM survey123 sv
  LEFT JOIN profile p1 ON p1.id = sv.creator
  LEFT JOIN profile p2 ON p2.id = sv.updater
  LEFT JOIN LATERAL (
    SELECT json_agg(json_build_object(
      'field_name', eq.field_name,
      'display_name', eq.display_name,
      'instrument_id', eq.instrument_id,
      'timeseries_id', eq.timeseries_id
    ) ORDER BY eq.field_name) AS fields
    FROM survey123_equivalency_table eq
    WHERE eq.survey123_id = sv.id
  ) f ON true
  LEFT JOIN LATERAL (
    SELECT array_agg(ier.error_message) AS errors
    FROM survey123_payload_error ier
    WHERE ier.survey123_id = sv.id
  ) er ON true
);

GRANT SELECT ON v_survey123 TO instrumentation_reader;
