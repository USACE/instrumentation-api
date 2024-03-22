-- ${flyway:timestamp}
CREATE VIEW v_report_config AS (
    SELECT
        rc.id,
        rc.slug,
        rc.name,
        rc.description,
        rc.project_id,
        p.name AS project_name,
        rc.after,
        rc.before,
        rc.creator,
        cp.username AS creator_username,
        rc.create_date,
        rc.updater,
        up.username AS updater_username,
        rc.update_date,
        COALESCE(pc.configs, '[]')::text AS plot_configs
    FROM report_config rc
    INNER JOIN project p ON rc.project_id = p.id
    INNER JOIN profile cp ON cp.id = rc.creator
    LEFT JOIN profile up ON up.id = rc.updater
    LEFT JOIN LATERAL (
        SELECT json_agg(json_build_object(
            'id', pc.id,
            'slug', pc.slug,
            'name', pc.name
        )) AS configs
        FROM plot_configuration pc
        WHERE pc.id = ANY(SELECT id FROM report_config_plot_config WHERE report_config_id = rc.id)
    ) pc ON true
);

GRANT SELECT ON v_report_config TO instrumentation_reader;
