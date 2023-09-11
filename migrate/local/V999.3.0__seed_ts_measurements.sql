INSERT INTO timeseries_measurement (timeseries_id, time, value)
SELECT
    timeseries_id,
    time,
    ROUND((RANDOM() * (100-3) + 3)::NUMERIC, 4) AS value
FROM
    UNNEST(ARRAY[
        '869465fc-dc1e-445e-81f4-9979b5fadda9'::UUID,
        '9a3864a8-8766-4bfa-bad1-0328b166f6a8'::UUID,
        '7ee902a3-56d0-4acf-8956-67ac82c03a96'::UUID,
        '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'::UUID,
        'd9697351-3a38-4194-9ac4-41541927e475'::UUID
    ]) AS timeseries_id,
    GENERATE_SERIES(
        NOW() - INTERVAL '1 year',
        NOW(),
        INTERVAL '1 hour'
    ) AS time;

-- INSERT INTO inclinometer_measurement (timeseries_id, time, depth, a0, a180, b0, b180)
-- SELECT
--     timeseries_id,
--     time,
--     ROUND(raw_depth + (RANDOM() * 2))::NUMERIC) AS depth,
--     ROUND((RANDOM() * (100-3) + 3)::NUMERIC, 4) AS a0,
--     ROUND((RANDOM() * (100-3) + 3)::NUMERIC, 4) AS a180,
--     ROUND((RANDOM() * (100-3) + 3)::NUMERIC, 4) AS b0,
--     ROUND((RANDOM() * (100-3) + 3)::NUMERIC, 4) AS b180
-- FROM
--     UNNEST(ARRAY[
--         '869465fc-dc1e-445e-81f4-9979b5fadda9'::UUID,
--         '9a3864a8-8766-4bfa-bad1-0328b166f6a8'::UUID,
--         '7ee902a3-56d0-4acf-8956-67ac82c03a96'::UUID
--     ]) AS timeseries_id,
--     GENERATE_SERIES(
--         NOW() - INTERVAL '1 year',
--         NOW(),
--         INTERVAL '1 hour'
--     ) AS time,
--     GENERATE_SERIES(
--         500,
--         0,
--         5
--     ) AS raw_depth;
