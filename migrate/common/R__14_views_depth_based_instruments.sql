DROP VIEW IF EXISTS v_saa_instrument;
CREATE VIEW v_saa_instrument AS (
    SELECT
        inst.*,
        saa.num_segments,
        saa.bottom_elevation
    FROM v_instrument inst
    INNER JOIN saa_instrument saa ON inst.id = saa.instrument_id
);

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
                timeseries_id,
                time,
                value                  AS x,
                NULL::DOUBLE PRECISION AS y,
                NULL::DOUBLE PRECISION AS z,
                NULL::DOUBLE PRECISION AS "temp"
            FROM timeseries_measurement
            WHERE timeseries_id = seg.x_timeseries_id
            UNION ALL
            SELECT
                timeseries_id,
                time,
                NULL::DOUBLE PRECISION AS x,
                value                  AS y,
                NULL::DOUBLE PRECISION AS z,
                NULL::DOUBLE PRECISION AS "temp"
            FROM timeseries_measurement
            WHERE timeseries_id = seg.y_timeseries_id
            UNION ALL
            SELECT
                timeseries_id,
                time,
                NULL::DOUBLE PRECISION AS x,
                NULL::DOUBLE PRECISION AS y,
                value                  AS z,
                NULL::DOUBLE PRECISION AS "temp"
            FROM timeseries_measurement
            WHERE timeseries_id = seg.z_timeseries_id
            UNION ALL
            SELECT
                timeseries_id,
                time,
                NULL::DOUBLE PRECISION AS x,
                NULL::DOUBLE PRECISION AS y,
                NULL::DOUBLE PRECISION AS z,
                value                  AS "temp"
            FROM timeseries_measurement
            WHERE timeseries_id = seg.temp_timeseries_id
        ) sq
        WHERE sq.timeseries_id IN (
            SELECT id FROM timeseries WHERE instrument_id = seg.instrument_id
        )
    ) q ON true
    GROUP BY seg.instrument_id, q.time
);
