CREATE OR REPLACE VIEW v_alert AS (
    SELECT a.id AS id,
       a.alert_config_id AS alert_config_id,
       a.create_date AS create_date,
       p.id AS project_id,
       p.name AS project_name,
	   ac.name AS name,
	   ac.body AS body,
       (
            SELECT COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
                'instrument_id',   id,
                'instrument_name', name
            ))::text, '[]'::text)
            FROM   instrument
            WHERE  id = ANY(
                SELECT iac.instrument_id
                FROM   alert_config_instrument iac
                WHERE  iac.alert_config_id = ac.id
            )
        )                           AS instruments
    FROM alert a
    INNER JOIN alert_config ac ON a.alert_config_id = ac.id
    INNER JOIN project p ON ac.project_id = p.id
);

CREATE OR REPLACE VIEW v_alert_config AS (
    SELECT
        ac.id                               AS id,
        ac.name                             AS name,
	    ac.body                             AS body,
        prf1.id                             AS creator,
        COALESCE(prf1.username, 'midas')    AS creator_username,
        ac.create_date                      AS create_date,
        prf2.id                             AS updater,
        prf2.username                       AS updater_username,
        ac.update_date                      AS update_date,
        prj.id                              AS project_id,
        prj.name                            AS project_name,
        atype.id                            AS alert_type_id,
        atype.name                          AS alert_type,
        ac.start_date                       AS start_date,
        ac.schedule_interval::text          AS schedule_interval,
        ac.n_missed_before_alert            AS n_missed_before_alert,
        ac.remind_interval::text            AS remind_interval,
        ac.warning_interval::text           AS warning_interval,
        ac.last_checked                     AS last_checked,
        ac.last_reminded                    AS last_reminded,
        astatus.name                        AS alert_status,
        (
            SELECT COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
                'instrument_id',   id,
                'instrument_name', name
            ))::text, '[]'::text)
            FROM   instrument
            WHERE  id = ANY(
                SELECT iac.instrument_id
                FROM   alert_config_instrument iac
                WHERE  iac.alert_config_id = ac.id
            )
        )                           AS instruments,
        (
            SELECT COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
                'id',       id,
                'user_type', user_type,
                'username', username,
                'email',    email
            ))::text, '[]'::text)
            FROM (
                SELECT
                    id,
                    'email'     AS user_type,
                    null        AS username,
                    email       AS email
                FROM email
                WHERE id IN (
                    SELECT aes.email_id FROM alert_email_subscription aes
                    WHERE aes.alert_config_id = ac.id
                )
                UNION
                SELECT
                    id,
                    'profile'   AS user_type,
                    username    AS username,
                    email       AS email
                FROM profile
                WHERE id IN (
                    SELECT aps.profile_id FROM alert_profile_subscription aps
                    WHERE aps.alert_config_id = ac.id
                )
            ) all_emails
        )                           AS alert_email_subscriptions
    FROM alert_config ac
    INNER JOIN project prj          ON ac.project_id = prj.id
    INNER JOIN alert_type atype     ON ac.alert_type_id = atype.id
    INNER JOIN alert_status astatus ON ac.alert_status_id = astatus.id
    LEFT  JOIN profile prf1         ON ac.creator = prf1.id
    LEFT  JOIN profile prf2         ON ac.updater = prf2.id
);

GRANT SELECT ON
    v_alert,
    v_alert_config
TO instrumentation_reader;
