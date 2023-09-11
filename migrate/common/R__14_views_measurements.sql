DROP VIEW IF EXISTS v_inclinometer_measurement;

CREATE VIEW v_inclinometer_measurement AS (
    SELECT
        ts.id AS timeseries_id,
        mmt.time AS "time",
        JSON_AGG(JSON_BUILD_OBJECT(
            'depth', mmt.depth,
            'a0', mmt.a0,
            'a180', mmt.a180,
            'a_checksum', mmt.a0 + mmt.a180,
            'a_comb', (mmt.a0 - mmt.a180) / 2,
            'a_increment', CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
                THEN 0
                ELSE (mmt.a0 - mmt.a180) / 2 / const_mmt.value * 24,
            'a_cum_dev', CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
                THEN 0
                ELSE SUM((mmt.a0 - mmt.a180) / 2 / const_mmt.value * 24) OVER (ORDER BY mmt.depth DESC),
            'b0', mmt.b0,
            'b180', mmt.b180,
            'b_checksum', mmt.b0 + mmt.b180,
            'b_comb', (mmt.b0 - mmt.b180) / 2,
            'b_increment', CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
                THEN 0
                ELSE (mmt.b0 - mmt.b180) / 2 / const_mmt.value * 24,
            'b_cum_dev', CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
                THEN 0
                ELSE SUM((mmt.b0 - mmt.b180) / 2 / const_mmt.value * 24) OVER (ORDER BY mmt.depth DESC)
        ) ORDER BY mmt.depth ASC)::TEXT AS "value"
    FROM timeseries ts
    INNER JOIN (
        SELECT
            ts.id AS timeseries_id,
            mmt1.time,
            mmt1.value,
        FROM timeseries ts
        JOIN LATERAL (
            SELECT * from timeseries_measurement mmt2
            WHERE ts.id = mmt2.timeseries_id
            ORDER BY mmt2.time DESC
            LIMIT 1
        ) mmt1 ON true
    ) const_mmt ON const_mmt.timeseries_id = ts.id
    INNER JOIN inclnometer_measurement mmt ON mmt.timeseries_id = ts.id
    INNER JOIN parameter pm ON ts.parameter_id = pm.id
    WHERE pm.name = 'inclinometer-constant'
    GROUP BY mmt.timeseries_id, mmt.time
    ORDER BY mmt.timeseries_id, mmt.time DESC
);

DROP VIEW IF EXISTS v_saa_measurement;

CREATE VIEW v_saa_measurement AS (
    SELECT
        ts.id AS timeseries_id,
        mmt.time,
        mmt.depth,
        mmt.x,
        CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
            THEN 0
            ELSE x / 2 / const_mmt.value * 24
        AS x_increment,
        CASE WHEN const_mmt.value IS NULL OR const_mmt.value = 0
            THEN 0
            ELSE SUM((items.a0 - items.a180) / 2 / const_mmt.value * 24) OVER (ORDER BY depth DESC)
        AS a_cum_dev,
        mmt.y
    FROM timeseries ts
    INNER JOIN (
        SELECT
            ts.id AS timeseries_id,
            mmt1.time,
            mmt1.value,
        FROM timeseries ts
        JOIN LATERAL (
            SELECT * from timeseries_measurement mmt2
            WHERE ts.id = mmt2.timeseries_id
            ORDER BY mmt2.time DESC
            LIMIT 1
        ) mmt1 ON true
        INNER JOIN timeseries ts ON ts.id = mmt1.timeseries_id
    ) const_mmt ON const_mmt.timeseries_id = ts.id
    INNER JOIN saa_measurement mmt ON mmt.timeseries_id = ts.id
);
