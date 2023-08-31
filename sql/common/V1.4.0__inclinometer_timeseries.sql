CREATE TABLE IF NOT EXISTS tmp_inclinometer_measurement (
    timeseries_id   UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    time            TIMESTAMPTZ NOT NULL,
    depth           REAL NOT NULL,
    a0              REAL NOT NULL,
    a180            REAL NOT NULL,
    b0              REAL NOT NULL,
    b180            REAL NOT NULL,
    CONSTRAINT inclinometer_measurement_timeseries_id_time_key UNIQUE(timeseries_id, time),
    CONSTRAINT inclinometer_measurement_time_depth_key UNIQUE(time, depth),
    PRIMARY KEY (timeseries_id, time)
);

INSERT INTO tmp_inclinometer_measurement (time, timeseries_id, depth, a0, a180, b0, b180)
SELECT
    mmt.timeseries_id   AS timeseries_id,
    mmt.time            AS time,
    nd.depth            AS depth,
    nd.a0               AS a0,
    nd.a180             AS a180,
    nd.b0               AS b0,
    nd.b180             AS b180
FROM inclinometer_measurement mmt,
JSONB_TO_RECORDSET(inclinometer_measurement.values) AS nd(depth REAL, a0 REAL, a180 REAL, b0 REAL, b180 REAL);

DROP TABLE inclinometer_measurement;

ALTER TABLE tmp_inclinometer_measurement RENAME TO inclinometer_measurement;
