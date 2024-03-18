-- ${flyway:timestamp}
CREATE VIEW v_datalogger AS (
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
        COALESCE(e.errors, '{}'::TEXT[]) AS errors,
        COALESCE(t.tables, '[]'::JSON)::TEXT AS tables
    FROM datalogger dl
    INNER JOIN profile p1         ON dl.creator = p1.id
    INNER JOIN profile p2         ON dl.updater = p2.id
    INNER JOIN datalogger_model m ON dl.model_id = m.id
    LEFT JOIN (
        SELECT
            de.datalogger_id,
            ARRAY_AGG(de.error_message) AS errors
        FROM datalogger_error de
        INNER JOIN datalogger_table dt ON dt.id = de.datalogger_table_id
        WHERE dt.table_name = 'preparse'
        GROUP BY de.datalogger_id
    ) e ON dl.id = e.datalogger_id
    LEFT JOIN (
        SELECT
            dt.datalogger_id,
            JSON_AGG(JSON_BUILD_OBJECT(
                'id',         dt.id,
                'table_name', dt.table_name
            )) AS tables
        FROM datalogger_table dt
        GROUP BY dt.datalogger_id
    ) t ON dl.id = t.datalogger_id
    WHERE NOT dl.deleted
);

CREATE VIEW v_datalogger_preview AS (
    SELECT
        p.datalogger_table_id,
        p.preview,
        p.update_date
    FROM datalogger_preview p
    INNER JOIN datalogger_table dt ON dt.id = p.datalogger_table_id
    INNER JOIN datalogger dl ON dl.id = dt.datalogger_id
    WHERE NOT dl.deleted
);

CREATE VIEW v_datalogger_equivalency_table AS (
    SELECT
        dt.datalogger_id AS datalogger_id,
        dt.id AS datalogger_table_id,
        dt.table_name AS datalogger_table_name,
        COALESCE(JSON_AGG(ROW_TO_JSON(eq)) FILTER (WHERE eq.id IS NOT NULL), '[]'::JSON)::TEXT AS fields
    FROM datalogger_table dt
    INNER JOIN datalogger dl ON dt.datalogger_id = dl.id
    LEFT JOIN LATERAL (
        SELECT id, field_name, display_name, instrument_id, timeseries_id
        FROM datalogger_equivalency_table
        WHERE datalogger_table_id = dt.id
    ) eq ON true
    WHERE NOT dl.deleted
    GROUP BY dt.datalogger_id, dt.id
);

CREATE VIEW v_datalogger_hash AS (
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
