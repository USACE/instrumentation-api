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

