CREATE OR REPLACE VIEW v_saa_segment AS (
    SELECT
        seg.id,
        seg.instrument_id,
        seg.length_timeseries_id,
        sub.length,
        seg.x_timeseries_id,
        seg.y_timeseries_id,
        seg.z_timeseries_id,
        seg.temp_timeseries_id
    FROM saa_segment seg
    LEFT JOIN LATERAL (
        SELECT value AS length FROM timeseries_measurement
        WHERE timeseries_id = seg.length_timeseries_id
        ORDER BY time DESC
        LIMIT 1
    ) sub ON true
);

CREATE OR REPLACE VIEW v_saa_measurement AS (
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
            'temp_cum_dev',     r.temp_cum_dev,
            'elevation',        r.elevation
        ) ORDER BY r.segment_id)::TEXT AS measurements
    FROM (SELECT DISTINCT
        seg.instrument_id,
        seg.id AS segment_id,
        q.time,
        q.x,
        q.y,
        q.z,
        q.t,
        q.initial_x - q.x x_increment,
        q.initial_y - q.y y_increment,
        q.initial_z - q.z z_increment,
        q.initial_t - q.t temp_increment,
        SUM(q.initial_x - q.x) FILTER (WHERE q.time >= q.initial_time) OVER (ORDER BY seg.id ASC) x_cum_dev,
        SUM(q.initial_y - q.y) FILTER (WHERE q.time >= q.initial_time) OVER (ORDER BY seg.id ASC) y_cum_dev,
        SUM(q.initial_z - q.z) FILTER (WHERE q.time >= q.initial_time) OVER (ORDER BY seg.id ASC) z_cum_dev,
        SUM(q.initial_t - q.t) FILTER (WHERE q.time >= q.initial_time) OVER (ORDER BY seg.id ASC) temp_cum_dev,
        SUM(q.bottom + q.seg_length) OVER (ORDER BY seg.id ASC) elevation
    FROM saa_segment seg
    INNER JOIN saa_opts opts ON opts.instrument_id = seg.instrument_id
    LEFT JOIN LATERAL (
        SELECT
            a.time,
            x.value AS x,
            y.value AS y,
            z.value AS z,
            t.value AS t,
            ia.time AS initial_time,
            ix.value AS initial_x,
            iy.value AS initial_y,
            iz.value AS initial_z,
            it.value AS initial_t,
            locf(b.value) OVER (ORDER BY a.time ASC) AS bottom,
            locf(l.value) OVER (ORDER BY a.time ASC) AS seg_length
        FROM (
            SELECT DISTINCT time FROM timeseries_measurement
            WHERE timeseries_id IN (SELECT id FROM timeseries WHERE instrument_id = seg.instrument_id)
            UNION
            SELECT time FROM timeseries_measurement WHERE time = opts.initial_time
        ) a
        LEFT JOIN LATERAL (SELECT time FROM timeseries_measurement WHERE time = opts.initial_time) ia ON true
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.x_timeseries_id) x ON x.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.y_timeseries_id) y ON y.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.z_timeseries_id) z ON z.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id) t ON t.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.x_timeseries_id) ix ON ix.time = ia.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.y_timeseries_id) iy ON iy.time = ia.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.z_timeseries_id) iz ON iz.time = ia.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id) it ON it.time = ia.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = opts.bottom_elevation_timeseries_id) b ON b.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.length_timeseries_id) l ON l.time = a.time
    ) q ON true) r
    GROUP BY r.instrument_id, r.time
);

CREATE OR REPLACE VIEW v_ipi_segment AS (
    SELECT
        seg.id,
        seg.instrument_id,
        seg.length_timeseries_id,
        sub.length,
        seg.tilt_timeseries_id,
        seg.inc_dev_timeseries_id
    FROM ipi_segment seg
    LEFT JOIN LATERAL (
        SELECT value AS length FROM timeseries_measurement
        WHERE timeseries_id = seg.length_timeseries_id
        ORDER BY time DESC
        LIMIT 1
    ) sub ON true
);

CREATE OR REPLACE VIEW v_ipi_measurement AS (
    SELECT
        r.instrument_id,
        r.time,
        JSON_AGG(JSON_BUILD_OBJECT(
            'segment_id',   r.segment_id,
            'tilt',         r.tilt,
            'inc_dev',      r.inc_dev,
            'cum_dev',      r.cum_dev,
            'temp',         r.temp,
            'elevation',    r.elevation
        ) ORDER BY r.segment_id)::TEXT AS measurements
    FROM (SELECT DISTINCT
        seg.instrument_id,
        seg.id AS segment_id,
        q.seg_length,
        q.time,
        q.tilt,
        q.inc_dev,
        COALESCE(q.cum_dev, SIN(q.tilt * PI() / 180) * q.seg_length) cum_dev,
        q.temp,
        SUM(q.bottom + q.seg_length) OVER (ORDER BY seg.id ASC) elevation
    FROM ipi_segment seg
    INNER JOIN ipi_opts opts ON opts.instrument_id = seg.instrument_id
    LEFT JOIN LATERAL (
        SELECT
            a.time,
            t.value AS tilt,
            d.value AS inc_dev,
            SUM(d.value) OVER (ORDER BY seg.id ASC) AS cum_dev,
            temp.value AS temp,
            locf(b.value) OVER (ORDER BY a.time ASC) AS bottom,
            locf(l.value) OVER (ORDER BY a.time ASC) AS seg_length
        FROM (
            SELECT DISTINCT time FROM timeseries_measurement
            WHERE timeseries_id IN (SELECT id FROM timeseries WHERE instrument_id = seg.instrument_id)
            UNION
            SELECT time FROM timeseries_measurement WHERE time = opts.initial_time
        ) a
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.tilt_timeseries_id) t ON t.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.inc_dev_timeseries_id) d ON d.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.temp_timeseries_id) temp ON temp.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = opts.bottom_elevation_timeseries_id) b ON b.time = a.time
        LEFT JOIN (SELECT time, value FROM timeseries_measurement WHERE timeseries_id = seg.length_timeseries_id) l ON l.time = a.time
    ) q ON true) r
    GROUP BY r.instrument_id, r.time
);
