-- inclinometer_measurement
CREATE TABLE IF NOT EXISTS inclinometer_measurement (
    time TIMESTAMPTZ NOT NULL,
    values JSONB NOT NULL,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT inclinometer_unique_time UNIQUE(timeseries_id,time),
    PRIMARY KEY (timeseries_id, time)
);

GRANT SELECT ON
    inclinometer_measurement
TO instrumentation_reader;

GRANT INSERT,UPDATE,DELETE ON
    inclinometer_measurement
TO instrumentation_writer;