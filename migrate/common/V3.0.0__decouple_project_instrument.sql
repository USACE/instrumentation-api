-- jump to v3.0 to sync with API release

DROP VIEW IF EXISTS v_timeseries_latest;
DROP VIEW IF EXISTS v_timeseries;
DROP VIEW IF EXISTS v_instrument;
DROP VIEW IF EXISTS v_project;
DROP VIEW IF EXISTS v_timeseries_project_map;
DROP VIEW IF EXISTS v_aware_platform_parameter_enabled;

CREATE TABLE IF NOT EXISTS project_instrument (
  project_id UUID NOT NULL REFERENCES project (id),
  instrument_id UUID NOT NULL REFERENCES instrument (id),
  CONSTRAINT project_instrument_project_id_instrument_id UNIQUE (project_id, instrument_id)
);

INSERT INTO project_instrument (project_id, instrument_id)
SELECT project_id, id FROM instrument;

ALTER TABLE instrument DROP COLUMN project_id;

-- This table is old cruft, completely empty
-- In practice, all timeseries are tied to instruments not projects
DROP TABLE project_timeseries;
