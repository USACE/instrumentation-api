-- ${flyway:timestamp}
CREATE VIEW v_plot_configuration AS (
    SELECT
        pc.id,
        pc.slug,
        pc.name,
        pc.project_id,
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
        COALESCE(rc.configs, '[]')::text      AS report_configs,
        json_build_object(
            'traces', COALESCE(traces.items, '[]'),
            'layout', json_build_object(
                'secondary_axis_title', k.secondary_axis_title,
                'custom_shapes', COALESCE(cs.items, '[]')
            )
        )::text                               AS display
    FROM plot_configuration pc
    LEFT JOIN (
        SELECT
            id,
            show_masked,
            show_nonvalidated,
            show_comments,
            auto_range,
            date_range,
            threshold,
            secondary_axis_title
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
    LEFT JOIN LATERAL (
        SELECT json_agg(
            to_jsonb(tr) || jsonb_build_object('name', i.name || ' - ' || ts.name || ' (' || u.name || ')')
            ORDER BY tr.trace_order ASC
        ) as items
        FROM plot_configuration_timeseries_trace tr
        INNER JOIN timeseries ts ON tr.timeseries_id = ts.id
        INNER JOIN instrument i ON ts.instrument_id = i.id
        INNER JOIN unit u ON ts.unit_id = u.id
        WHERE tr.plot_configuration_id = pc.id
    ) traces ON true
    LEFT JOIN LATERAL (
        SELECT json_agg(to_json(ccs)) AS items
        FROM plot_configuration_custom_shape ccs
        WHERE pc.id = ccs.plot_configuration_id
    ) cs on true
);

GRANT SELECT ON v_plot_configuration TO instrumentation_reader;
