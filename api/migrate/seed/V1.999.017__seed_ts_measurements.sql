INSERT INTO timeseries_measurement (timeseries_id, time, value)
SELECT
    timeseries_id,
    time,
    round((random() * (100-3) + 3)::NUMERIC, 4) AS value
FROM
    unnest(ARRAY[
        '869465fc-dc1e-445e-81f4-9979b5fadda9'::uuid,
        '9a3864a8-8766-4bfa-bad1-0328b166f6a8'::uuid,
        '7ee902a3-56d0-4acf-8956-67ac82c03a96'::uuid,
        '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'::uuid,
        'd9697351-3a38-4194-9ac4-41541927e475'::uuid
    ]) AS timeseries_id,
    generate_series(
        now() - INTERVAL '1 year',
        now(),
        INTERVAL '1 hour'
    ) AS time;
