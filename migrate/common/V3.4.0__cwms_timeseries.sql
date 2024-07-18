CREATE TABLE timeseries_cwms (
  timeseries_id uuid NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
  cwms_timeseries_id text NOT NULL,
  cwms_office_id text NOT NULL
);
