-- ${flyway:timestamp}
CREATE OR REPLACE VIEW v_alert_check_measurement_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        sub.id AS submittal_id,
        COALESCE(
            ac.warning_interval != INTERVAL '0'
            AND sub.completion_date IS NULL
            AND NOW() >= sub.due_date - ac.warning_interval
            AND NOW() < sub.due_date
            AND true = ANY(SELECT UNNEST(ARRAY_AGG(lm.time)) IS NULL),
            true
        ) AS should_warn,
        COALESCE(
            sub.completion_date IS NULL
            AND NOT sub.marked_as_missing
            AND NOW() >= sub.due_date
            AND true = ANY(SELECT UNNEST(ARRAY_AGG(lm.time)) IS NULL),
            true
        ) AS should_alert,
        COALESCE(
            ac.remind_interval != INTERVAL '0'
            AND ac.last_reminded IS NOT NULL
            AND sub.completion_date IS NULL
            AND NOT sub.marked_as_missing
            AND NOW() >= sub.due_date
            -- subtract 10 second constant to account for ticker accuracy/execution time
            AND NOW() >= ac.last_reminded + ac.remind_interval - INTERVAL '10 seconds',
            true
        ) AS should_remind,
        COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
            'instrument_name', inst.name,
            'timeseries_name', COALESCE(ts.name, 'No timeseries for instrument'),
            'status', CASE
                WHEN NOW() >= sub.due_date THEN 'missing'
                WHEN NOW() < sub.due_date  THEN 'warning'
                ELSE 'N/A'
            END
        )) FILTER (WHERE lm.time IS NULL), '[]')::text AS affected_timeseries
    FROM alert_config ac
    INNER JOIN submittal sub ON sub.alert_config_id = ac.id
    INNER JOIN alert_config_instrument aci ON aci.alert_config_id = ac.id
    INNER JOIN instrument inst ON aci.instrument_id = inst.id
    -- forces the query planner to use a loose index scan, which Postgres does not do automatically yet
    -- https://stackoverflow.com/questions/25536422/optimize-group-by-query-to-retrieve-latest-row-per-user/25536748#25536748
    LEFT JOIN LATERAL (
        SELECT
            timeseries_id,
            MAX(time) FILTER (WHERE time > sub.create_date AND time <= sub.due_date) AS time
        FROM timeseries_measurement
        WHERE timeseries_id = ANY(SELECT id FROM timeseries WHERE instrument_id = inst.id)
        AND NOT timeseries_id = ANY(SELECT timeseries_id FROM instrument_constants)
        GROUP BY timeseries_id
    ) lm ON true
    LEFT JOIN timeseries ts ON ts.id = lm.timeseries_id
    WHERE ac.alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID
    AND NOT ac.deleted
    GROUP BY ac.id, sub.id
);

CREATE OR REPLACE VIEW v_alert_check_evaluation_submittal AS (
    SELECT
        ac.id AS alert_config_id,
        sub.id AS submittal_id,
        COALESCE(
            ac.warning_interval != INTERVAL '0'
            AND sub.completion_date IS NULL
            AND NOW() >= sub.due_date - ac.warning_interval
            AND NOW() < sub.due_date,
            true
        ) AS should_warn,
        COALESCE(
            sub.completion_date IS NULL
            AND NOW() >= sub.due_date
            AND NOT sub.marked_as_missing,
            true
        ) AS should_alert,
        COALESCE(
            ac.remind_interval != INTERVAL '0'
            AND ac.last_reminded IS NOT NULL
            AND sub.completion_date IS NULL
            AND NOW() >= sub.due_date
            -- subtract 10 second constant to account for ticker accuracy/execution time
            AND NOW() >= ac.last_reminded + ac.remind_interval - INTERVAL '10 seconds'
            AND NOT sub.marked_as_missing,
            true
        ) AS should_remind
    FROM submittal sub
    INNER JOIN alert_config ac ON sub.alert_config_id = ac.id
    WHERE ac.alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID
    AND NOT ac.deleted
);

GRANT SELECT ON
    v_alert_check_measurement_submittal,
    v_alert_check_evaluation_submittal
TO instrumentation_reader;
