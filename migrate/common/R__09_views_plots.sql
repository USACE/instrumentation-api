-- ${flyway:timestamp}
CREATE OR REPLACE VIEW v_plot_configuration AS (
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
        pc.plot_type,
        CASE
            WHEN pc.plot_type = 'scatter-line' THEN json_build_object(
                'traces', COALESCE(traces.items, '[]'),
                'layout', json_build_object(
                    'y_axis_title', pcl.y_axis_title,
                    'y2_axis_title', pcl.y2_axis_title,
                    'custom_shapes', COALESCE(cs.items, '[]')
                )
            )::text
            WHEN pc.plot_type = 'profile' THEN json_build_object(
                'instrument_id', ppc.instrument_id,
                'instrument_type', it.name
            )::text
            WHEN pc.plot_type = 'contour' THEN json_build_object(
                'timeseries_ids', COALESCE(pcct.timeseries_ids, '{}'),
                'time', to_char(time, 'YYYY-MM-DD"T"HH24:MI:SS.US') || 'Z',
                'locf_backfill', pcc.locf_backfill,
                'gradient_smoothing', pcc.gradient_smoothing,
                'contour_smoothing', pcc.contour_smoothing,
                'show_labels', pcc.show_labels
            )::text
            WHEN pc.plot_type = 'bullseye' THEN json_build_object(
                'x_axis_timeseries_id', pbc.x_axis_timeseries_id,
                'y_axis_timeseries_id', pbc.y_axis_timeseries_id
            )::text
            ELSE NULL
        END AS display
    FROM plot_configuration pc
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
    LEFT JOIN LATERAL (
        SELECT json_agg(
            to_jsonb(tr) || jsonb_build_object(
                'name', i.name || ' - ' || ts.name || ' (' || u.name || ')',
                'parameter', p.name
            )
            ORDER BY tr.trace_order ASC
        ) as items
        FROM plot_configuration_timeseries_trace tr
        INNER JOIN timeseries ts ON tr.timeseries_id = ts.id
        INNER JOIN instrument i ON ts.instrument_id = i.id
        INNER JOIN unit u ON ts.unit_id = u.id
        INNER JOIN parameter p ON ts.parameter_id = p.id
        WHERE tr.plot_configuration_id = pc.id
    ) traces ON true
    LEFT JOIN LATERAL (
        SELECT json_agg(to_json(ccs)) AS items
        FROM plot_configuration_custom_shape ccs
        WHERE pc.id = ccs.plot_configuration_id
    ) cs on true
    LEFT JOIN plot_bullseye_config pbc ON pbc.plot_config_id = pc.id
    LEFT JOIN plot_profile_config ppc ON ppc.plot_config_id = pc.id
    LEFT JOIN instrument ii ON ii.id = ppc.instrument_id
    LEFT JOIN instrument_type it ON it.id = ii.type_id
    LEFT JOIN plot_contour_config pcc ON pcc.plot_config_id = pc.id
    LEFT JOIN LATERAL (
        SELECT array_agg(ipcct.timeseries_id) as timeseries_ids
        FROM plot_contour_config_timeseries ipcct
        WHERE ipcct.plot_contour_config_id = pc.id
    ) pcct ON true
    LEFT JOIN plot_scatter_line_config pcl ON pcl.plot_config_id = pc.id
    ORDER BY pc.name
);

GRANT SELECT ON v_plot_configuration TO instrumentation_reader;
