-- v_datalogger
CREATE OR REPLACE VIEW v_datalogger AS (
    SELECT
        dl.id          AS id,
        dl.sn          AS sn,
        dl.project_id  AS project_id,
        p1.id          AS creator,
        p1.username    AS creator_username,
        dl.create_date AS create_date,
        p2.id          AS updater,
        p2.username    AS updater_username,
        dl.update_date AS update_date,
        dl.name        AS name,
        dl.slug        AS slug,
        m.id           AS model_id,
        m.model        AS model,
        dl.deleted     AS deleted
    FROM datalogger dl
    INNER JOIN profile p1         ON dl.creator = p1.id
    INNER JOIN profile p2         ON dl.creator = p2.id
    INNER JOIN datalogger_model m ON dl.model_id = m.id
    WHERE NOT dl.deleted
);

-- v_datalogger_preview
CREATE OR REPLACE VIEW v_datalogger_preview AS (
    SELECT
        p.datalogger_id                  AS datalogger_id,
        p.preview                        AS preview,
        p.update_date                    AS update_date,
        m.model                          AS model,
        dl.sn                            AS sn,
        COALESCE(e.errors, '{}'::TEXT[]) AS errors
    FROM datalogger_preview p
    INNER JOIN datalogger dl      ON p.datalogger_id = dl.id
    INNER JOIN datalogger_model m ON dl.model_id = m.id
    LEFT JOIN (
        SELECT
            datalogger_id,
            array_agg(error_message) AS errors
        FROM datalogger_error
        GROUP BY datalogger_id
    ) e ON p.datalogger_id = e.datalogger_id
    WHERE NOT dl.deleted
);

-- v_datalogger_equivalency_table
CREATE OR REPLACE VIEW v_datalogger_equivalency_table AS (
    SELECT
        t.datalogger_id AS datalogger_id,
        t.field_name    AS field_name,
        t.display_name  AS display_name,
        t.instrument_id AS instrument_id,
        t.timeseries_id AS timeseries_id,
        m.model         AS model,
        dl.sn           AS sn
    FROM datalogger_equivalency_table t
    INNER JOIN datalogger dl      ON t.datalogger_id = dl.id
    INNER JOIN datalogger_model m ON dl.model_id = m.id
    WHERE NOT t.datalogger_deleted
);

-- v_datalogger_hash
CREATE OR REPLACE VIEW v_datalogger_hash AS (
    SELECT
        dh.datalogger_id AS datalogger_id,
        dh.hash          AS "hash",
        m.model          AS model,
        dl.sn            AS sn
    FROM datalogger_hash dh
    INNER JOIN datalogger dl      ON dh.datalogger_id = dl.id
    INNER JOIN datalogger_model m ON dl.model_id = m.id
    WHERE NOT dl.deleted
);

GRANT SELECT ON
    v_datalogger,
    v_datalogger_preview,
    v_datalogger_equivalency_table,
    v_datalogger_hash
TO instrumentation_reader;
