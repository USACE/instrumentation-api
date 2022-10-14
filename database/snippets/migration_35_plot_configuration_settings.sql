CREATE TABLE IF NOT EXISTS plot_configuration_settings (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    show_masked BOOLEAN DEFAULT 'false',
    show_nonvalidated BOOLEAN DEFAULT 'false',
    show_comments BOOLEAN DEFAULT 'false',

    FOREIGN KEY id REFERENCES plot_configuration (id)
);

CREATE OR REPLACE VIEW v_plot_configuration AS (
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
           COALESCE(k.show_comments, 'true')      AS show_comments
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
               show_comments     AS show_comments
        FROM plot_configuration_settings
        GROUP BY id
    ) as k ON pc.id = k.id
);

GRANT INSERT,UPDATE,DELETE
    ON plot_configuration_settings,
       v_plot_configuration
    TO instrumentation_writer;
GRANT SELECT
    ON plot_configuration_settings,
       v_plot_configuration
    TO instrumentation_reader;
