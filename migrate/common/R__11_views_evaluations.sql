DROP VIEW IF EXISTS v_evaluation;

CREATE VIEW v_evaluation AS (
    SELECT
        ev.id                               AS id,
        ev.name                             AS name,
	    ev.body                             AS body,
        prf1.id                             AS creator,
        COALESCE(prf1.username, 'midas')    AS creator_username,
        ev.create_date                      AS create_date,
        prf2.id                             AS updater,
        prf2.username                       AS updater_username,
        ev.update_date                      AS update_date,
        prj.id                              AS project_id,
        prj.name                            AS project_name,
        ac.id                               AS alert_config_id,
        ac.name                             AS alert_config_name,
        ev.submittal_id                     AS submittal_id,
        ev.start_date                       AS start_date,
        ev.end_date                         AS end_date,
        (
            SELECT COALESCE(JSON_AGG(JSON_BUILD_OBJECT(
                'instrument_id',   id,
                'instrument_name', name
            ))::text, '[]'::text)
            FROM   instrument
            WHERE  id = ANY(
                SELECT evi.instrument_id
                FROM   evaluation_instrument evi
                WHERE  evi.evaluation_id = ev.id
            )
        )                                   AS instruments
    FROM evaluation ev
    INNER JOIN project prj  ON ev.project_id = prj.id
    LEFT  JOIN profile prf1 ON ev.creator = prf1.id
    LEFT  JOIN profile prf2 ON ev.updater = prf2.id
    LEFT  JOIN submittal sub ON sub.id = ev.submittal_id
    LEFT  JOIN alert_config ac ON ac.id = sub.alert_config_id
);

GRANT SELECT ON
    v_evaluation
TO instrumentation_reader;
