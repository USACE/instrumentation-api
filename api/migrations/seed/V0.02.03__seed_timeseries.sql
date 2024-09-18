INSERT INTO timeseries_measurement (timeseries_id, time, value)
SELECT
    timeseries_id,
    time,
    round((random() * (100-3) + 3)::numeric, 4) AS value
FROM
    (SELECT timeseries_id FROM timeseries_measurement) AS timeseries_id,
    generate_series(now()-'1 year'::interval,now(),'1 hour'::interval) AS time;
