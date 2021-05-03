SELECT a.instrument_id,
       a.parsed_variable,
	   ts.id,
	   ts.variable
FROM (
	SELECT id AS instrument_id,
	 (regexp_matches(formula, '\[(.*?)\]', 'g')) AS parsed_variable
	FROM instrument
) a
LEFT JOIN v_timeseries ts ON a.parsed_variable = ts.variable

-- Invalid Instrument Formulas
-- https://stackoverflow.com/questions/42849929/display-all-results-queried-on-postgresql-where-the-joining-value-is-missing

-- All Instruments with Formulas
SELECT id, formula FROM instrument WHERE formula IS NOT NULL


-- Instruments whose formulas reference other timeseries
SELECT id AS instrument_id,
	 (regexp_matches(formula, '\[(.*?)\]', 'g'))[1] AS parsed_variable
FROM instrument


-- Figure out which parsed variables are legitimate timeseries in MIDAS
-- NOTE: Null timeseries_id columns mean the variable string used in the instrument
--       formula does not correspond to an actual factual timeseries in MIDAS
--       this could be caused by a typo, outdated formula (ts since deleted), etc..
SELECT a.instrument_id   AS instrument_id,
       a.parsed_variable AS parsed_variable,
	   ts.id             AS timeseries_id
FROM (
	--instruments whoes formulas reference other timeseries 	
	SELECT id AS instrument_id,
		 (regexp_matches(formula, '\[(.*?)\]', 'g'))[1] AS parsed_variable
	FROM instrument
) a
LEFT JOIN v_timeseries ts ON a.parsed_variable = ts.variable


-- Get Timeseries and Dependencies for Computations
-- timeseries required based on requested instrument
WITH required_timeseries AS (
-- 	Timeseries for Instrument
	SELECT id FROM timeseries WHERE instrument_id = 'a7540f69-c41e-43b3-b655-6e44097edb7e'
	UNION
-- Dependencies for Instrument Timeseries
	SELECT dependency_timeseries_id AS id FROM v_timeseries_dependency WHERE instrument_id = 'a7540f69-c41e-43b3-b655-6e44097edb7e'
),
-- Next Timeseries Measurement Outside Time Window (Earlier); Needed for Calculation Interpolation
next_low AS (
	SELECT nlm.timeseries_id AS timeseries_id, json_build_object(nlm.time, m1.value) AS measurement
	FROM (
		SELECT timeseries_id, MAX(time) AS time
		FROM timeseries_measurement
		WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time < '2016-01-01'
		GROUP BY timeseries_id
	) nlm
	INNER JOIN timeseries_measurement m1 ON m1.time = nlm.time AND m1.timeseries_id = nlm.timeseries_id
),
-- Next Timeseries Measurement Outside Time Window (Later); Needed For Calculation Interpolation
next_high AS (
	SELECT nhm.timeseries_id AS timeseries_id, json_build_object(nhm.time, m2.value) AS measurement
	FROM (
		SELECT timeseries_id, MIN(time) AS time
		FROM timeseries_measurement
		WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time > '2021-09-01'
		GROUP BY timeseries_id
	) nhm
	INNER JOIN timeseries_measurement m2 ON m2.time = nhm.time AND m2.timeseries_id = nhm.timeseries_id
),
-- Measurements Within Time Window by timeseries_id
measurements AS (
	SELECT timeseries_id,
	       json_agg(json_build_object(time, value))::text AS measurements
	FROM timeseries_measurement
	WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND
	      time >= '2016-01-01' AND time <= '2021-09-01'
	GROUP BY timeseries_id
)
-- Stored Timeseries
SELECT r.id                     AS timeseries_id,
       ts.instrument_id         AS instrument_id,
	   i.slug || '.' || ts.slug AS variable,
	   false                    AS is_computed,
	   null                     AS formula,
	   m.measurements           AS measurements,
	   nl.measurement::text     AS next_measurement_low,
	   nh.measurement::text     AS next_measurement_high
FROM required_timeseries r
INNER JOIN timeseries ts ON ts.id = r.id
INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id = 'a7540f69-c41e-43b3-b655-6e44097edb7e'
INNER JOIN measurements m ON m.timeseries_id = r.id
LEFT JOIN next_low nl ON nl.timeseries_id = r.id
LEFT JOIN next_high nh ON nh.timeseries_id = r.id
UNION
-- Computed Timeseries
SELECT i.formula_id            AS timeseries_id,
	   i.id                    AS instrument_id,
	   i.slug || '.formula'    AS variable,
	   true                    AS is_computed,
	   i.formula               AS formula,
	   '[]'::text              AS measurements,
	   null                    AS next_measurement_low,
	   null                    AS next_measurement_high
FROM instrument i
WHERE i.id = 'a7540f69-c41e-43b3-b655-6e44097edb7e'
