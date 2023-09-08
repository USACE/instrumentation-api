CREATE OR REPLACE VIEW v_project AS (
    SELECT  p.id,
            p.federal_id,
            CASE WHEN p.image IS NOT NULL
                THEN cfg.static_host || '/projects/' || p.slug || '/images/' || p.image
                ELSE NULL
            END AS image,
            p.office_id,
            p.deleted,
            p.slug,
            p.name,
            p.creator,
            p.create_date,
            p.updater,
            p.update_date,
            COALESCE(t.timeseries, '{}') AS timeseries,
            COALESCE(i.count, 0) AS instrument_count,
            COALESCE(g.count, 0) AS instrument_group_count
        FROM project p
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument) as count
                FROM instrument
                WHERE NOT instrument.deleted
                GROUP BY project_id
            ) i ON i.project_id = p.id
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument_group) as count
                FROM instrument_group
                WHERE NOT instrument_group.deleted
                GROUP BY project_id
            ) g ON g.project_id = p.id
            LEFT JOIN (
                SELECT array_agg(timeseries_id) as timeseries,
                    project_id
                FROM project_timeseries
                GROUP BY project_id
            ) t on t.project_id = p.id
			CROSS JOIN config cfg
);

CREATE OR REPLACE VIEW v_district AS (
    SELECT
        dis.id          AS id,
        dis.name        AS name,
        dis.initials    AS initials,
        div.name        AS division_name,
        div.initials    AS division_initials,
        dis.office_id   AS office_id
    FROM district dis
    INNER JOIN division div ON dis.division_id = div.id
);

GRANT SELECT ON
    v_project,
    v_district
TO instrumentation_reader;