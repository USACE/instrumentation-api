CREATE OR REPLACE VIEW v_alert_check_measurement_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        (ac.warning_interval != INTERVAL 'PT0'
            AND now() >= ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert) - ac.warning_interval
            AND (now() - (ac.schedule_interval * ac.n_missed_before_alert) + ac.warning_interval
            ) < ANY(ARRAY_AGG(COALESCE(lm.time, ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert) - ac.warning_interval)))
        ) AS should_warn,
        (now() >= ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert)
            AND (now() - (ac.schedule_interval * ac.n_missed_before_alert)
            ) < ANY(ARRAY_AGG(COALESCE(lm.time, ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert))))
        ) AS should_alert,
        (ac.remind_interval != INTERVAL 'PT0'
            AND now() >= COALESCE(ac.last_reminded, (ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert))) + ac.remind_interval
        ) AS should_remind,
        (COALESCE(MIN(lm.time), ac.start_date) + (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
        COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
            'instrument_name', inst.name,
            'last_measurement_time', lm.time
        )) FILTER (
            WHERE COALESCE(lm.time, ac.start_date) < (now() - (ac.schedule_interval * ac.n_missed_before_alert) + ac.warning_interval)
        ), '[]')::text AS affected_instruments
    FROM alert_config ac
    INNER JOIN alert_config_instrument aci ON aci.alert_config_id = ac.id
    INNER JOIN instrument inst ON aci.instrument_id = inst.id
    INNER JOIN (
        -- this subquery forces the query planner to use a loose index scan, which Postgres does not do automatically yet
        -- https://stackoverflow.com/questions/25536422/optimize-group-by-query-to-retrieve-latest-row-per-user/25536748#25536748
        SELECT inst2.id AS instrument_id, mmt.time AS time
        FROM (
            SELECT id FROM instrument
            ORDER BY id
        ) inst2
        LEFT JOIN LATERAL (
            SELECT time FROM timeseries_measurement
            WHERE timeseries_id = ANY(SELECT id FROM timeseries WHERE instrument_id = inst2.id)
            AND NOT timeseries_id = ANY(SELECT timeseries_id FROM instrument_constants)
            AND time <= now()
            ORDER BY time DESC NULLS LAST
            LIMIT 1
        ) mmt ON true
    ) lm ON lm.instrument_id = inst.id
    WHERE ac.alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID
    AND NOT ac.deleted
    GROUP BY ac.id
);

CREATE OR REPLACE VIEW v_alert_check_evaluation_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        ((ac.warning_interval != INTERVAL 'PT0')
            AND (now() >= ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert) - ac.warning_interval)
            AND COALESCE(le.time, ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert) - ac.warning_interval
            ) < (now() - (ac.schedule_interval * ac.n_missed_before_alert) + ac.warning_interval)
        ) AS should_warn,
        ((now() >= ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert))
            AND COALESCE(le.time, ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert)
            ) < (now() - (ac.schedule_interval * ac.n_missed_before_alert))
        ) AS should_alert,
        ((ac.remind_interval != INTERVAL 'PT0')
            AND (now() >= COALESCE(ac.last_reminded, (ac.start_date + (ac.schedule_interval * ac.n_missed_before_alert))) + ac.remind_interval)
        ) AS should_remind,
        (COALESCE(le.time, ac.start_date) + (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
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
    AND NOT ac.deleted
);

GRANT SELECT ON
    v_alert_check_measurement_submittal,
    v_alert_check_evaluation_submittal
TO instrumentation_reader;
