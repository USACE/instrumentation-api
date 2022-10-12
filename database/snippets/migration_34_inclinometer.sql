-- New table was created.
CREATE TABLE IF NOT EXISTS timeseries_notes (
    masked boolean NOT NULL DEFAULT false,
    validated boolean NOT NULL DEFAULT false,
    annotation varchar(400) NOT NULL DEFAULT '',
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    time TIMESTAMPTZ NOT NULL,
    CONSTRAINT notes_unique_time UNIQUE(timeseries_id, time),
    PRIMARY KEY (timeseries_id, time)
);

-- Permissions
GRANT SELECT ON timeseries_notes TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON timeseries_notes TO instrumentation_writer;
