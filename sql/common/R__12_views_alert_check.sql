CREATE OR REPLACE VIEW v_alert_check_measurement_submittal AS (
    WITH alert_interval AS (
        SELECT
            id AS alert_config_id,
            NOW() - (schedule_interval * n_missed_before_alert)                             AS last_schedule,
            NOW() - (schedule_interval * n_missed_before_alert) + warning_interval          AS last_warning,
            create_date + (schedule_interval * n_missed_before_alert)                       AS first_schedule,
            create_date + (schedule_interval * n_missed_before_alert) - warning_interval    AS first_warning,
            start_date + (schedule_interval * n_missed_before_alert)                        AS start_schedule,
            start_date + (schedule_interval * n_missed_before_alert) - warning_interval     AS start_warning
        FROM alert_config
    )
    SELECT
        ac.id AS alert_config_id,
        (ac.warning_interval != INTERVAL 'PT0S'
            AND NOW() >= ai.start_warning
            AND ((true = ANY(SELECT UNNEST(ARRAY_AGG(lm.time)) IS NULL) AND ai.last_warning >= ai.first_warning)
                OR (ai.last_warning >= ANY(ARRAY_AGG(lm.time))))
        ) AS should_warn,
        (NOW() >= ai.start_schedule
            AND ((true = ANY(SELECT UNNEST(ARRAY_AGG(lm.time)) IS NULL) AND ai.last_schedule >= ai.first_schedule)
                OR (ai.last_schedule >= ANY(ARRAY_AGG(lm.time))))
        ) AS should_alert,
        (ac.remind_interval != INTERVAL 'PT0S'
            AND NOW() >= COALESCE(ac.last_reminded, ai.start_schedule) + ac.remind_interval
        ) AS should_remind,
        (COALESCE(
            (SELECT MAX(lmt) FROM UNNEST(ARRAY_AGG(lm.time) FILTER (WHERE lm.time IS NULL OR lm.time <= ai.last_warning)) lmt),
            ac.create_date
        ) + (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
        COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
            'instrument_name', inst.name,
            'timeseries_name', ts.name,
            'last_measurement_time', lm.time,
            'status', CASE
                WHEN lm.time IS NULL
                    THEN CASE
                        WHEN ai.last_schedule >= ai.first_schedule THEN 'red'
                        WHEN ai.last_warning  >= ai.first_warning  THEN 'yellow'
                        ELSE 'green'
                    END
                WHEN lm.time <= ai.last_schedule THEN 'red'
                WHEN lm.time <= ai.last_warning THEN 'yellow'
                ELSE 'green'
            END
        )) FILTER (WHERE lm.time IS NULL OR lm.time <= ai.last_warning), '[]')::text AS affected_timeseries
    FROM alert_config ac
    INNER JOIN alert_interval ai ON ai.alert_config_id = ac.id
    INNER JOIN alert_config_instrument aci ON aci.alert_config_id = ac.id
    INNER JOIN instrument inst ON aci.instrument_id = inst.id
    -- forces the query planner to use a loose index scan, which Postgres does not do automatically yet
    -- https://stackoverflow.com/questions/25536422/optimize-group-by-query-to-retrieve-latest-row-per-user/25536748#25536748
    LEFT JOIN LATERAL (
        SELECT timeseries_id, time FROM timeseries_measurement
        WHERE timeseries_id = ANY(SELECT id FROM timeseries WHERE instrument_id = inst.id)
            AND NOT timeseries_id = ANY(SELECT timeseries_id FROM instrument_constants)
            AND time <= NOW()
        ORDER BY time DESC NULLS LAST
        LIMIT 1
    ) lm ON true
    INNER JOIN timeseries ts ON ts.id = lm.timeseries_id
    WHERE ac.alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID AND NOT ac.deleted
    GROUP BY ac.id, ai.last_schedule, ai.last_warning, ai.first_schedule, ai.first_warning, ai.start_schedule, ai.start_warning
);

CREATE OR REPLACE VIEW v_alert_check_evaluation_submittal AS (
    WITH alert_interval AS (
        SELECT
            id AS alert_config_id,
            NOW() - (schedule_interval * n_missed_before_alert)                             AS last_schedule,
            NOW() - (schedule_interval * n_missed_before_alert) + warning_interval          AS last_warning,
            create_date + (schedule_interval * n_missed_before_alert)                       AS first_schedule,
            create_date + (schedule_interval * n_missed_before_alert) - warning_interval    AS first_warning,
            start_date + (schedule_interval * n_missed_before_alert)                        AS start_schedule,
            start_date + (schedule_interval * n_missed_before_alert) - warning_interval     AS start_warning
        FROM alert_config
    )
    SELECT
        ac.id AS alert_config_id,
        (ac.warning_interval != INTERVAL 'PT0S'
            AND NOW() >= ai.start_warning
            AND COALESCE(le.time, ai.first_warning) < ai.last_warning
        ) AS should_warn,
        (NOW() >= ai.start_schedule
            AND COALESCE(le.time, ai.first_schedule) < ai.last_schedule
        ) AS should_alert,
        (ac.remind_interval != INTERVAL 'PT0S'
            AND NOW() >= COALESCE(ac.last_reminded, ai.start_schedule) + ac.remind_interval
        ) AS should_remind,
        (COALESCE(le.time, ac.create_date) + (ac.schedule_interval * ac.n_missed_before_alert)) AS expected_submittal,
        le.time AS last_evaluation_time
    FROM alert_config ac
    INNER JOIN alert_interval ai ON ai.alert_config_id = ac.id
    LEFT JOIN (
        SELECT
            alert_config_id,
            MAX(create_date) AS time
        FROM evaluation
        GROUP BY alert_config_id
    ) le ON le.alert_config_id = ac.id
    WHERE ac.alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID AND NOT ac.deleted
);

GRANT SELECT ON
    v_alert_check_measurement_submittal,
    v_alert_check_evaluation_submittal
TO instrumentation_reader;
