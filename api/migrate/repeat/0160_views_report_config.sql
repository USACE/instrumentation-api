-- +goose up
CREATE OR REPLACE VIEW v_report_config AS (
    SELECT
        rc.id,
        rc.slug,
        rc.name,
        rc.description,
        rc.project_id,
        p.name AS project_name,
        dt.name AS district_name,
        rc.creator,
        cp.username AS creator_username,
        rc.create_date,
        rc.updater,
        up.username AS updater_username,
        rc.update_date,
        COALESCE(pc.configs, '[]')::text AS plot_configs,
        json_build_object(
            'date_range', json_build_object(
                'enabled', rc.date_range_enabled,
                'value', rc.date_range
            ),
            'show_masked', json_build_object(
                'enabled', rc.show_masked_enabled,
                'value', rc.show_masked
            ),
            'show_nonvalidated', json_build_object(
                'enabled', rc.show_nonvalidated_enabled,
                'value', rc.show_nonvalidated
            )
        )::text AS global_overrides
    FROM report_config rc
    INNER JOIN project p ON rc.project_id = p.id
    LEFT JOIN district dt ON p.district_id = dt.id
    INNER JOIN profile cp ON cp.id = rc.creator
    LEFT JOIN profile up ON up.id = rc.updater
    LEFT JOIN LATERAL (
        SELECT json_agg(json_build_object(
            'id', pc.id,
            'slug', pc.slug,
            'name', pc.name
        )) AS configs
        FROM plot_configuration pc
        WHERE pc.id = ANY(SELECT plot_config_id FROM report_config_plot_config WHERE report_config_id = rc.id)
    ) pc ON true
);
