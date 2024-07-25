CREATE TYPE timeseries_type AS ENUM ('standard', 'constant', 'computed', 'cwms');

ALTER TABLE timeseries ADD COLUMN type timeseries_type;

UPDATE timeseries SET type = 'standard';

UPDATE timeseries SET type = 'constant'
WHERE id = ANY(SELECT ic.timeseries_id FROM instrument_constants ic)
OR id = ANY(SELECT so.bottom_elevation_timeseries_id FROM saa_opts so)
OR id = ANY(SELECT io.bottom_elevation_timeseries_id FROM ipi_opts io);

UPDATE timeseries SET type = 'computed'
WHERE id = ANY(SELECT ca.timeseries_id FROM calculation ca);

CREATE TABLE timeseries_cwms (
  timeseries_id uuid NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
  cwms_timeseries_id text NOT NULL,
  cwms_office_id text NOT NULL
);
