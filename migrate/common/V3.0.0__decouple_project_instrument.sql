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
  CONSTRAINT project_instrument_project_id_instrument_id_key UNIQUE (project_id, instrument_id)
);

INSERT INTO project_instrument (project_id, instrument_id)
SELECT project_id, id FROM instrument;

ALTER TABLE instrument DROP COLUMN project_id;

-- This table is old cruft, completely empty
-- In practice, all timeseries are tied to instruments not projects
DROP TABLE project_timeseries;

---

CREATE TABLE IF NOT EXISTS agency (
  id UUID UNIQUE NOT NULL,
  name TEXT UNIQUE NOT NULL
);

INSERT INTO agency (id, name)
  VALUES ('db5a9cf6-860f-49d9-be25-34eb069478ef', 'USACE');

ALTER TABLE division
  ADD COLUMN agency_id UUID REFERENCES agency (id);

UPDATE division
  SET agency_id = 'db5a9cf6-860f-49d9-be25-34eb069478ef';

ALTER TABLE division
  ALTER COLUMN agency_id SET NOT NULL;

---

ALTER TABLE project
  ADD COLUMN district_id UUID REFERENCES district (id);

UPDATE project p SET (district_id) = (
  SELECT id
  FROM district d
  WHERE d.office_id = p.office_id
  AND p.office_id IS NOT NULL
);

ALTER TABLE project
  DROP COLUMN office_id;
