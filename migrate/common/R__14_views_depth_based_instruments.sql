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
        r.instrument_id,
        r.time,
        JSON_AGG(JSON_BUILD_OBJECT(
            'segment_id',       r.segment_id,
            'x',                r.x,
            'y',                r.y,
            'z',                r.z,
            'temp',             r.t,
            'x_increment',      r.x_increment,
            'y_increment',      r.y_increment,
            'z_increment',      r.z_increment,
            'temp_increment',   r.temp_increment,
            'x_cum_dev',        r.x_cum_dev,
            'y_cum_dev',        r.y_cum_dev,
            'z_cum_dev',        r.z_cum_dev,
            'temp_cum_dev',     r.temp_cum_dev
        ) ORDER BY r.segment_id)::TEXT AS measurements
    FROM (SELECT DISTINCT
        seg.instrument_id,
        seg.id AS segment_id,
        q.time,
        q.x,
        q.y,
        q.z,
        q.t,
        CASE WHEN q.initial_x IS NULL THEN NULL ELSE (q.initial_x - q.x) END x_increment,
        CASE WHEN q.initial_y IS NULL THEN NULL ELSE (q.initial_y - q.y) END y_increment,
        CASE WHEN q.initial_z IS NULL THEN NULL ELSE (q.initial_z - q.z) END z_increment,
        CASE WHEN q.initial_t IS NULL THEN NULL ELSE (q.initial_t - q.t) END temp_increment,
        CASE WHEN q.initial_x IS NULL THEN NULL ELSE SUM(q.initial_x - q.x) OVER (ORDER BY seg.id DESC) END x_cum_dev,
        CASE WHEN q.initial_y IS NULL THEN NULL ELSE SUM(q.initial_y - q.y) OVER (ORDER BY seg.id DESC) END y_cum_dev,
        CASE WHEN q.initial_z IS NULL THEN NULL ELSE SUM(q.initial_z - q.z) OVER (ORDER BY seg.id DESC) END z_cum_dev,
        CASE WHEN q.initial_t IS NULL THEN NULL ELSE SUM(q.initial_t - q.t) OVER (ORDER BY seg.id DESC) END temp_cum_dev
    FROM saa_segment seg
    INNER JOIN saa_opts opts ON opts.instrument_id = seg.instrument_id
    LEFT JOIN LATERAL (
        SELECT
            a.time,
            x.value AS x,
            y.value AS y,
            z.value AS z,
            t.value AS t,
            ix.value AS initial_x,
            iy.value AS initial_y,
            iz.value AS initial_z,
            it.value AS initial_t
        FROM (SELECT DISTINCT time FROM timeseries_measurement WHERE timeseries_id IN (SELECT id FROM timeseries WHERE instrument_id = seg.instrument_id)) a
        LEFT JOIN LATERAL (SELECT time FROM timeseries_measurement WHERE time = opts.initial_time) ia ON true
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.x_timeseries_id) x ON x.time = a.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.y_timeseries_id) y ON y.time = a.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.z_timeseries_id) z ON z.time = a.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id) t ON t.time = a.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.x_timeseries_id) ix ON ix.time = ia.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.y_timeseries_id) iy ON iy.time = ia.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.z_timeseries_id) iz ON iz.time = ia.time
        LEFT JOIN (SELECT * FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id) it ON it.time = ia.time
        WHERE a.time >= ia.time
    ) q ON true) r
    GROUP BY r.instrument_id, r.time
);

GRANT SELECT ON
    v_saa_segment,
    v_saa_measurement
TO instrumentation_reader;
