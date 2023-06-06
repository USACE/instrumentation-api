CREATE OR REPLACE VIEW v_alert_check_measurement_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        (SELECT (ac.warning_interval != INTERVAL 'PT0') AND NOT EXISTS (
            SELECT 1 WHERE MAX(lm.time) >= (now() - (ac.schedule_interval * ac.n_missed_before_alert) + ac.warning_interval)
        )) AS should_warn,
        (SELECT NOT EXISTS (
            SELECT 1 WHERE MAX(lm.time) >= (now() - (ac.schedule_interval * ac.n_missed_before_alert))
        )) AS should_alert,
        (
            (ac.remind_interval != INTERVAL 'PT0')
            AND (now() >= COALESCE(ac.last_reminded, (ac.start_date + ac.schedule_interval)) + ac.remind_interval)
        ) AS should_remind,
        (now() - (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
        (SELECT COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
            'instrument_name', name,
            'last_measurement_time', MAX(lm.time)
        )) FILTER (
            WHERE MAX(lm.time) < (now() - (ac.schedule_interval * ac.n_missed_before_alert))
        ), '[]')::text
        FROM   instrument
        WHERE  id = ANY(
            SELECT iac.instrument_id
            FROM   alert_config_instrument iac
            WHERE  iac.alert_config_id = ac.id
        )) AS affected_instruments
    FROM alert_config ac
    INNER JOIN alert_config_instrument aci ON aci.alert_config_id = ac.id
    INNER JOIN timeseries ts ON ts.instrument_id = aci.instrument_id
    INNER JOIN (
        SELECT
            timeseries_id,
            MAX(time) AS time
        FROM timeseries_measurement
        WHERE NOT timeseries_id = ANY(SELECT timeseries_id FROM instrument_constants)
        AND time <= now()
        GROUP BY timeseries_id
    ) lm ON lm.timeseries_id = ts.id
    WHERE ac.alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID
    GROUP BY ac.id
);

CREATE OR REPLACE VIEW v_alert_check_evaluation_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        (SELECT (ac.warning_interval != INTERVAL 'PT0') AND NOT EXISTS (
            SELECT 1 WHERE le.time >= (now() - (ac.schedule_interval * ac.n_missed_before_alert) + ac.warning_interval)
        )) AS should_warn,
        (SELECT NOT EXISTS (
            SELECT 1 WHERE le.time >= (now() - (ac.schedule_interval * ac.n_missed_before_alert))
        )) AS should_alert,
        (
            (ac.remind_interval != INTERVAL 'PT0')
            AND (now() >= COALESCE(ac.last_reminded, (ac.start_date + ac.schedule_interval)) + ac.remind_interval)
        ) AS should_remind,
        (now() - (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
        le.time AS last_evaluation_time
    FROM alert_config ac
    LEFT JOIN (
        SELECT
            alert_config_id,
            MAX(create_date) AS time
        FROM evaluation
        GROUP BY alert_config_id
    ) le ON le.alert_config_id = ac.id
    WHERE ac.alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID
);

GRANT SELECT ON
    v_alert_check_measurement_submittal,
    v_alert_check_evaluation_submittal
TO instrumentation_reader;
