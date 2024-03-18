-- ${flyway:timestamp}
CREATE VIEW v_project AS (
    SELECT
        p.id,
        p.federal_id,
        CASE WHEN p.image IS NOT NULL
            THEN cfg.static_host || '/projects/' || p.slug || '/images/' || p.image
            ELSE NULL
        END AS image,
        p.district_id,
        d.office_id,
        p.deleted,
        p.slug,
        p.name,
        p.creator,
        u.username AS creator_username,
        p.create_date,
        p.updater,
        u.username AS updater_username,
        p.update_date,
        COALESCE(i.count, 0) AS instrument_count,
        COALESCE(g.count, 0) AS instrument_group_count
    FROM project p
    LEFT JOIN profile c ON p.creator = c.id
    LEFT JOIN profile u ON p.updater = c.id
    LEFT JOIN (
        SELECT pi.project_id, COUNT(pi.*) as count
        FROM project_instrument pi
        INNER JOIN instrument i ON i.id = pi.instrument_id
        WHERE NOT i.deleted
        GROUP BY pi.project_id
    ) i ON i.project_id = p.id
    LEFT JOIN (
        SELECT project_id, COUNT(*) as count
        FROM instrument_group
        WHERE NOT deleted
        GROUP BY project_id
    ) g ON g.project_id = p.id
    LEFT JOIN (
        SELECT id, office_id FROM district
    ) d ON d.id = p.district_id
    CROSS JOIN config cfg
);

CREATE VIEW v_district AS (
    SELECT
        ag.name         AS agency,
        dis.id          AS id,
        dis.name        AS name,
        dis.initials    AS initials,
        div.name        AS division_name,
        div.initials    AS division_initials,
        dis.office_id   AS office_id
    FROM district dis
    INNER JOIN division div ON dis.division_id = div.id
    INNER JOIN agency ag ON ag.id = div.agency_id
);

GRANT SELECT ON
    v_project,
    v_district
TO instrumentation_reader;
