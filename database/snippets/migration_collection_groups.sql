DROP VIEW IF EXISTS v_timeseries;

-- collection_group
CREATE TABLE IF NOT EXISTS public.collection_group (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT project_unique_collection_group_name UNIQUE(project_id, name),
    CONSTRAINT project_unique_collection_group_slug UNIQUE(project_id, slug)
);

CREATE TABLE IF NOT EXISTS public.collection_group_timeseries (
    collection_group_id UUID NOT NULL REFERENCES collection_group(id) ON DELETE CASCADE,
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT collection_group_unique_timeseries UNIQUE(collection_group_id, timeseries_id)
);

-- v_timeseries
CREATE OR REPLACE VIEW v_timeseries AS (
        SELECT t.id AS id,
            t.slug AS slug,
            t.name AS name,
            i.slug || '.' || t.slug AS variable,
            j.id AS project_id,
            j.slug AS project_slug,
            j.name AS project,
            i.id AS instrument_id,
            i.slug AS instrument_slug,
            i.name AS instrument,
            p.id AS parameter_id,
            p.name AS parameter,
            u.id AS unit_id,
            u.name AS unit
        FROM timeseries t
            LEFT JOIN instrument i ON i.id = t.instrument_id
            LEFT JOIN project j ON j.id = i.project_id
            INNER JOIN parameter p ON p.id = t.parameter_id
            INNER JOIN unit U ON u.id = t.unit_id
    );

-- v_timeseries_project_map
CREATE OR REPLACE VIEW v_timeseries_project_map AS (
    SELECT t.id AS timeseries_id,
           p.id AS project_id
    FROM timeseries t
    LEFT JOIN instrument n ON t.instrument_id = n.id
    LEFT JOIN project p ON p.id = n.project_id
);

-- v_timeseries_latest; same as v_timeseries, joined with latest times and values
CREATE OR REPLACE VIEW v_timeseries_latest AS (
    SELECT t.*,
       m.time AS latest_time,
	   m.value AS latest_value
    FROM v_timeseries t
    LEFT JOIN (
	    SELECT DISTINCT ON (timeseries_id) timeseries_id, time, value
	    FROM timeseries_measurement
	    ORDER BY timeseries_id, time DESC
    ) m ON t.id = m.timeseries_id
);

-- Role instrumentation_reader
-- Tables specific to instrumentation app
GRANT SELECT ON
    collection_group,
    collection_group_timeseries,
    v_timeseries,
    v_timeseries_latest,
    v_timeseries_project_map
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    collection_group,
    collection_group_timeseries
TO instrumentation_writer;

-- collection_group
INSERT INTO collection_group (id, project_id, name, slug) VALUES
    ('1519eaea-1799-4375-aa37-0e35aa654643', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Manual Collection Route 1', 'manual-collection-route-1'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'High Water Inspection', 'high-water-inspection');

-- collection_group_timeseries
INSERT INTO collection_group_timeseries (collection_group_id, timeseries_id) VALUES
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '7ee902a3-56d0-4acf-8956-67ac82c03a96'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '9a3864a8-8766-4bfa-bad1-0328b166f6a8');