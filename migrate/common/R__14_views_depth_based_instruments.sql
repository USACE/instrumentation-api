DROP VIEW IF EXISTS v_saa_segment;
CREATE VIEW v_saa_segment AS (
    SELECT
        id,
        instrument_id,
        length,
        x_timeseries_id,
        y_timeseries_id,
        z_timeseries_id,
        temp_timeseries_id
    FROM saa_segment
);

DROP VIEW IF EXISTS v_saa_measurement;
CREATE VIEW v_saa_measurement AS (
    SELECT
        seg.instrument_id,
        q.time,
        JSON_AGG(JSON_BUILD_OBJECT(
            'segment_id', seg.id,
            'x', q.x,
            'y', q.y,
            'z', q.z,
            'temp', q.temp
        ) ORDER BY seg.id)::TEXT AS measurements
    FROM saa_segment seg
    LEFT JOIN LATERAL (
        SELECT sq.time, sq.x, sq.y, sq.z, sq.temp
        FROM (
            SELECT
                a.time,
                x.value AS x,
                y.value AS y,
                z.value AS z,
                t.value AS "temp"
            FROM (SELECT DISTINCT time FROM timeseries_measurement WHERE timeseries_id IN (
                SELECT id FROM timeseries WHERE instrument_id = seg.instrument_id
            )) a
            LEFT JOIN (
                SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.x_timeseries_id
            ) x ON x.time = a.time
            LEFT JOIN (
                SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.y_timeseries_id
            ) y ON y.time = a.time
            LEFT JOIN (
                SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.z_timeseries_id
            ) z ON z.time = a.time
            LEFT JOIN (
                SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id
            ) t ON t.time = a.time
        ) sq
    ) q ON true
    GROUP BY seg.instrument_id, q.time
);

GRANT SELECT ON
    v_saa_segment,
    v_saa_measurement
TO instrumentation_reader;
