CREATE TABLE IF NOT EXISTS public.config (
    static_host VARCHAR NOT NULL DEFAULT 'http://minio:9000',
    static_prefix VARCHAR NOT NULL DEFAULT '/instrumentation'
);

INSERT INTO public.config (static_host, static_prefix) VALUES
    ('https://api.rsgis.dev', '/instrumentation');

-- Add image field to project table
ALTER TABLE public.project ADD COLUMN image VARCHAR;

-- Update Sample Data
UPDATE PROJECT SET image = 'site_photo.jpg' WHERE id = '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984';

-- Update View v_project
DROP VIEW v_project;
CREATE OR REPLACE VIEW v_project AS (
    SELECT  p.id,
            CASE WHEN p.image IS NOT NULL
                THEN cfg.static_host || cfg.static_prefix || '/projects/' || p.id || '/images/' || p.image
                ELSE NULL
            END AS image,
            p.office_id,
            p.deleted,
            p.slug,
            p.federal_id,
            p.name,
            p.creator,
            p.create_date,
            p.updater,
            p.update_date,
            COALESCE(t.timeseries, '{}') AS timeseries,
            COALESCE(i.count, 0) AS instrument_count,
            COALESCE(g.count, 0) AS instrument_group_count
        FROM project p
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument) as count
                FROM instrument
                WHERE NOT instrument.deleted
                GROUP BY project_id
            ) i ON i.project_id = p.id
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument_group) as count
                FROM instrument_group
                WHERE NOT instrument_group.deleted
                GROUP BY project_id
            ) g ON g.project_id = p.id
            LEFT JOIN (
                SELECT array_agg(timeseries_id) as timeseries,
                    project_id
                FROM project_timeseries
                GROUP BY project_id
            ) t on t.project_id = p.id
			CROSS JOIN config cfg
);

-- Role Changes
GRANT SELECT ON config, v_project TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON config TO instrumentation_writer;
