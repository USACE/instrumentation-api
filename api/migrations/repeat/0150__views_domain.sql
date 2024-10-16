CREATE OR REPLACE VIEW v_domain AS (
    SELECT
        id, 
        'instrument_type' AS group, 
        name AS value,
        icon AS description
    FROM instrument_type 
    UNION 
    SELECT
        id, 
        'parameter' AS group, 
        name AS value,
        null AS description
    FROM parameter 
    UNION 
    SELECT
        id, 
        'unit' AS group, 
        name AS value,
        null AS description
    FROM unit
    UNION
    SELECT
        id,
        'status' AS group,
        name AS value,
        description AS description
    FROM status
    UNION
    SELECT
        id,
        'role' AS group,
        name AS value,
        null AS description
    FROM role
    UNION
    SELECT
        id,
        'datalogger_model' AS group,
        model AS value,
        null AS description
    FROM datalogger_model
    UNION
    SELECT
        id,
        'submittal_status' AS group,
        name AS value,
        null AS description
    FROM submittal_status
    UNION
    SELECT
        id,
        'alert_type' AS group,
        name AS value,
        null AS description
    FROM alert_type
    ORDER BY "group", value
);

CREATE OR REPLACE VIEW v_domain_group AS (
  SELECT
    "group",
    json_agg(json_build_object(
      'id', id,
      'value', value,
      'description', description
    ))::text AS opts
  FROM v_domain
  GROUP BY "group"
);
