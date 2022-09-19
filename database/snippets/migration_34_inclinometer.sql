-- New table was created.
CREATE TABLE IF NOT EXISTS timeseries_notes (
    masked boolean NOT NULL DEFAULT false,
    validated boolean NOT NULL DEFAULT false,
    annotation varchar(400) NOT NULL DEFAULT '',
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT timeseries_unique_time UNIQUE(timeseries_id,time),
    time TIMESTAMPTZ NOT NULL,
    CONSTRAINT notes_unique_time UNIQUE(timeseries_id, time),
    PRIMARY KEY (timeseries_id, time)
);

-- A new column, timeseries_id, was added.
ALTER TABLE timeseries_measurement ADD COLUMN timeseries_id UUID;
ALTER TABLE timeseries_measurement ADD CONSTRAINT timeseries_unique_time UNIQUE(timeseries_id, time);

-- Primary key was changed from just "timeseries_id" to "(timeseries_id, time)".
ALTER TABLE timeseries_measurement DROP CONSTRAINT timeseries_measurement_pkey;
ALTER TABLE timeseries_measurement ADD PRIMARY KEY (timeseries_id, time);

-- Permissions
GRANT SELECT ON timeseries_notes TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON timeseries_notes TO instrumentation_writer;
