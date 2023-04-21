ALTER TABLE timeseries_notes
    ALTER COLUMN "masked" DROP NOT NULL,
    ALTER COLUMN validated DROP NOT NULL,
    ALTER COLUMN annotation DROP NOT NULL;

CREATE INDEX timeseries_measurement_time_idx ON timeseries_measurement (time);
