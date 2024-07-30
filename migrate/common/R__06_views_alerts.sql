-- ${flyway:timestamp}
CREATE OR REPLACE VIEW v_alert AS (
    SELECT a.id AS id,
       a.alert_config_id AS alert_config_id,
       a.create_date AS create_date,
       p.id AS project_id,
       p.name AS project_name,
       ac.name AS name,
       ac.body AS body,
       (
            SELECT COALESCE(json_agg(json_build_object(
                'instrument_id',   id,
                'instrument_name', name
            ))::text, '[]'::text)
            FROM instrument
            WHERE id = ANY(
                SELECT iac.instrument_id
                FROM   alert_config_instrument iac
                WHERE  iac.alert_config_id = ac.id
            )
        ) AS instruments
    FROM alert a
    INNER JOIN alert_config ac ON a.alert_config_id = ac.id
    INNER JOIN project p ON ac.project_id = p.id
);

CREATE OR REPLACE VIEW v_alert_config AS (
    SELECT
        ac.id,
        ac.name,
        ac.body,
        prf1.id AS creator,
        COALESCE(prf1.username, 'midas') AS creator_username,
        ac.create_date,
        prf2.id AS updater,
        prf2.username AS updater_username,
        ac.update_date,
        prj.id AS project_id,
        prj.name AS project_name,
        ac.last_checked,
        atype.id AS alert_type_id,
        atype.name AS alert_type,
        CASE
            -- measurement-submittal and evaluation-submittal
            WHEN atype.id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::uuid OR atype.id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::uuid THEN json_build_object(
                'last_reminded', to_char(acs.last_reminded, 'YYYY-MM-DD"T"HH24:MI:SS.US') || 'Z',
                'mute_consecutive_alerts', acs.mute_consecutive_alerts,
                'start_date', to_char(acs.start_date, 'YYYY-MM-DD"T"HH24:MI:SS.US') || 'Z',
                'schedule_interval', acs.schedule_interval::text,
                'remind_interval', remind_interval::text,
                'warning_interval', acs.warning_interval::text
            )::text
            -- threshold
            WHEN atype.id = 'bb15e7c2-8eae-452c-92f7-e720dc5c9432'::uuid THEN json_build_object(
            )::text
            -- rate of change
            WHEN atype.id = 'c37effee-6b48-4436-8d72-737ed78c1fb7'::uuid THEN json_build_object(
            )::text
        END AS opts,
        (
            SELECT COALESCE(json_agg(json_build_object(
                'instrument_id', ii.id,
                'instrument_name', ii.name
            ))::text, '[]'::text)
            FROM instrument ii
            WHERE ii.id = ANY(
                SELECT iac.instrument_id
                FROM   alert_config_instrument iac
                WHERE  iac.alert_config_id = ac.id
            )
        ) AS instruments,
        (
            SELECT COALESCE(json_agg(json_build_object(
                'id', ae.id,
                'user_type', ae.user_type,
                'username', ae.username,
                'email', ae.email
            ))::text, '[]'::text)
            FROM (
                SELECT
                    ie.id,
                    'email' AS user_type,
                    null AS username,
                    ie.email AS email
                FROM email ie
                WHERE ie.id IN (
                    SELECT aes.email_id FROM alert_email_subscription aes
                    WHERE aes.alert_config_id = ac.id
                )
                UNION
                SELECT
                    ip.id,
                    'profile' AS user_type,
                    ip.username AS username,
                    ip.email AS email
                FROM profile ip
                WHERE ip.id IN (
                    SELECT aps.profile_id FROM alert_profile_subscription aps
                    WHERE aps.alert_config_id = ac.id
                )
            ) ae
        ) AS alert_email_subscriptions
    FROM alert_config ac
    LEFT JOIN alert_config_scheduler acs ON acs.alert_config_id = ac.id
    INNER JOIN project prj ON ac.project_id = prj.id
    INNER JOIN alert_type atype ON ac.alert_type_id = atype.id
    LEFT JOIN profile prf1 ON ac.creator = prf1.id
    LEFT JOIN profile prf2 ON ac.updater = prf2.id
    WHERE NOT ac.deleted
);

CREATE OR REPLACE VIEW v_submittal AS (
    SELECT
        sub.id,
        ac.id AS alert_config_id,
        ac.name AS alert_config_name,
        aty.id AS alert_type_id,
        aty.name AS alert_type_name,
        ac.project_id,
        sst.id AS submittal_status_id,
        sst.name AS submittal_status_name,
        sub.completion_date,
        sub.create_date,
        sub.due_date,
        sub.marked_as_missing,
        sub.warning_sent
    FROM submittal sub
    INNER JOIN alert_config ac ON sub.alert_config_id = ac.id
    INNER JOIN submittal_status sst ON sub.submittal_status_id = sst.id
    INNER JOIN alert_type aty ON ac.alert_type_id = aty.id
);

GRANT SELECT ON
    v_alert,
    v_alert_config,
    v_submittal
TO instrumentation_reader;
