-- plot_configuration
CREATE TABLE IF NOT EXISTS public.plot_configuration (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    project_id UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    CONSTRAINT project_unique_plot_configuration_name UNIQUE(project_id, name),
    CONSTRAINT project_unique_plot_configuration_slug UNIQUE(project_id, slug)
);

CREATE TABLE IF NOT EXISTS public.plot_configuration_timeseries (
    plot_configuration_id UUID NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT plot_configuration_unique_timeseries UNIQUE(plot_configuration_id, timeseries_id)
);

CREATE OR REPLACE VIEW v_plot_configuration AS (
    SELECT pc.id            AS id,
           pc.slug          AS slug,
           pc.name          AS name,
           pc.project_id    AS project_id,
           t.timeseries_id     AS timeseries_id,
           pc.creator       AS creator,
           pc.create_date   AS create_date,
           pc.updater       AS updater,
           pc.update_date   AS update_date
    FROM plot_configuration pc
    LEFT JOIN (
        SELECT plot_configuration_id    as plot_configuration_id,
               array_agg(timeseries_id) as timeseries_id
        FROM plot_configuration_timeseries
        GROUP BY plot_configuration_id
    ) as t ON pc.id = t.plot_configuration_id
);

GRANT SELECT ON
    plot_configuration,
    plot_configuration_timeseries,
    v_plot_configuration
TO instrumentation_reader;

GRANT INSERT,UPDATE,DELETE ON
    plot_configuration,
    plot_configuration_timeseries
TO instrumentation_writer;
