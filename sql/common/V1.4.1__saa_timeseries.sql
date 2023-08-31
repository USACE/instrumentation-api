CREATE TABLE IF NOT EXISTS saa_measurement (
    timeseries_id       UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    time                TIMESTAMPTZ NOT NULL,
    elevation_change    REAL NOT NULL,
    x                   REAL NOT NULL,
    y                   REAL NOT NULL,
    temperature         REAL NOT NULL,
    CONSTRAINT saa_measurement_time_depth_key UNIQUE(time, depth),
    PRIMARY KEY (timeseries_id, time)
);
