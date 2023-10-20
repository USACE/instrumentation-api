DROP VIEW IF EXISTS v_plot_configuration;
CREATE VIEW v_plot_configuration AS (
    SELECT pc.id                                  AS id,
           pc.slug                                AS slug,
           pc.name                                AS name,
           pc.project_id                          AS project_id,
           t.timeseries_id                        AS timeseries_id,
           pc.creator                             AS creator,
           pc.create_date                         AS create_date,
           pc.updater                             AS updater,
           pc.update_date                         AS update_date,
           COALESCE(k.show_masked, 'true')        AS show_masked,
           COALESCE(k.show_nonvalidated, 'true')  AS show_nonvalidated,
           COALESCE(k.show_comments, 'true')      AS show_comments,
           COALESCE(k.auto_range, 'true')         AS auto_range,
           COALESCE(k.date_range, '1 year')       AS date_range,
           COALESCE(k.threshold, 3000)            AS threshold
    FROM plot_configuration pc
    LEFT JOIN (
        SELECT plot_configuration_id    AS plot_configuration_id,
               array_agg(timeseries_id) AS timeseries_id
        FROM plot_configuration_timeseries
        GROUP BY plot_configuration_id
    ) as t ON pc.id = t.plot_configuration_id
    LEFT JOIN (
        SELECT id                AS id,
               show_masked       AS show_masked,
               show_nonvalidated AS show_nonvalidated,
               show_comments     AS show_comments,
               auto_range        AS auto_range,
               date_range        AS date_range,
               threshold         AS threshold
        FROM plot_configuration_settings
        GROUP BY id
    ) AS k ON pc.id = k.id
);

GRANT SELECT ON
    v_plot_configuration
TO instrumentation_reader;
