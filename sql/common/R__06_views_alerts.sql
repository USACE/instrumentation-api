-- v_alert
CREATE OR REPLACE VIEW v_alert AS (
    SELECT a.id AS id,
       a.alert_config_id AS alert_config_id,
       a.create_date AS create_date,
       p.id AS project_id,
       p.name AS project_name,
	   i.id AS instrument_id,
	   i.name AS instrument_name,
	   ac.name AS name,
	   ac.body AS body
FROM alert a
INNER JOIN alert_config ac ON a.alert_config_id = ac.id
INNER JOIN instrument i ON ac.instrument_id = i.id
INNER JOIN project p ON i.project_id = p.id
);

GRANT SELECT ON
    v_alert
TO instrumentation_reader;
