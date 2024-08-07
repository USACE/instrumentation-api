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
    COALESCE(json_agg(json_build_object(
      'field_name', eq.field_name,
      'display_name', eq.display_name,
      'instrument_id', eq.instrument_id,
      'timeseries_id', eq.timeseries_id
    ) ORDER BY eq.field_name), '[]'::json)::text AS fields,
    COALESCE(array_agg(er.error_message)), '{}') AS errors
  FROM survey123 sv
  LEFT JOIN survey123_equivalency_table eq ON eq.survey123_id = sv.id
  LEFT JOIN survey123_payload_errors er ON er.survey123_id = sv.id
  LEFT JOIN profile p1 ON p1.id = sv.creator
  LEFT JOIN profile p2 ON p2.id = sv.updater
);
