-- Add auto_range and date_range fields to plot configurtion settings
-- Add constraints and backfill existing data for other setting options
-- Regenerate view

set search_path = "$user", midas, public, topology;

BEGIN;

ALTER TABLE plot_configuration_settings
    ADD auto_range BOOLEAN DEFAULT true,
    ADD date_range VARCHAR DEFAULT '1 year';

UPDATE plot_configuration_settings
    SET auto_range = true
    WHERE auto_range IS NULL;

UPDATE plot_configuration_settings
    SET date_range = '1 year'
    WHERE date_range IS NULL;

UPDATE plot_configuration_settings
    SET show_masked = true
    WHERE show_masked IS NULL;

UPDATE plot_configuration_settings
    SET show_nonvalidated = true
    WHERE show_nonvalidated IS NULL;

UPDATE plot_configuration_settings
    SET show_comments = true
    WHERE show_comments IS NULL;

ALTER TABLE plot_configuration_settings
    ALTER COLUMN auto_range SET NOT NULL,
    ALTER COLUMN date_range SET NOT NULL,
    ALTER COLUMN show_masked SET DEFAULT true,
    ALTER COLUMN show_nonvalidated SET DEFAULT true,
    ALTER COLUMN show_comments SET DEFAULT true,
    ALTER COLUMN show_masked SET NOT NULL,
    ALTER COLUMN show_nonvalidated SET NOT NULL,
    ALTER COLUMN show_comments SET NOT NULL;

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
           COALESCE(k.show_comments, 'true')      AS show_comments,
           COALESCE(k.auto_range, 'true')         AS auto_range,
           COALESCE(k.date_range, '1 year')       AS date_range
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
               date_range        AS date_range
        FROM plot_configuration_settings
        GROUP BY id
    ) as k ON pc.id = k.id
);

COMMIT;
