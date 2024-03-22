-- ${flyway:timestamp}
CREATE VIEW v_plot_configuration AS (
    SELECT
        pc.id,
        pc.slug,
        pc.name,
        pc.project_id,
        t.timeseries_id,
        pc.creator,
        pc.create_date,
        pc.updater,
        pc.update_date,
        COALESCE(k.show_masked, 'true')       AS show_masked,
        COALESCE(k.show_nonvalidated, 'true') AS show_nonvalidated,
        COALESCE(k.show_comments, 'true')     AS show_comments,
        COALESCE(k.auto_range, 'true')        AS auto_range,
        COALESCE(k.date_range, '1 year')      AS date_range,
        COALESCE(k.threshold, 3000)           AS threshold,
        COALESCE(rc.configs, '[]')::text      AS report_configs
    FROM plot_configuration pc
    LEFT JOIN (
        SELECT
            plot_configuration_id,
            array_agg(timeseries_id) AS timeseries_id
        FROM plot_configuration_timeseries
        GROUP BY plot_configuration_id
    ) t ON pc.id = t.plot_configuration_id
    LEFT JOIN (
        SELECT
            id,
            show_masked,
            show_nonvalidated,
            show_comments,
            auto_range,
            date_range,
            threshold
        FROM plot_configuration_settings
        GROUP BY id
    ) k ON pc.id = k.id
    LEFT JOIN LATERAL (
        SELECT
            json_agg(json_build_object(
                'id', id,
                'slug', slug,
                'name', name
            )) AS configs
        FROM report_config
        WHERE id = ANY(SELECT report_config_id FROM report_config_plot_config WHERE plot_config_id = pc.id)
    ) rc ON true
);

GRANT SELECT ON v_plot_configuration TO instrumentation_reader;
