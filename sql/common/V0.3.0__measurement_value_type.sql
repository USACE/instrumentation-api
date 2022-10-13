
    			    
-- Drop the depedant view
drop view v_timeseries_latest;  

-- Drop index on timeseries_measurement table for faster processing
--
-- WARNING: This may or may not be the name of the index
--
DROP INDEX  timeseries_measurement_pkey;

-- Add temp column to store value
ALTER TABLE timeseries_measurement
  ADD COLUMN temp_val double precision;
 
-- Store value in temp field 
update timeseries_measurement set temp_val = round(value::numeric,2);

-- Change data type of value field
ALTER TABLE timeseries_measurement
  ALTER COLUMN value TYPE double precision;

-- Restore the value from temp back to value field
update timeseries_measurement set value = temp_val;

-- remove temp field
ALTER TABLE timeseries_measurement
  DROP COLUMN temp_val;

-- Restore index to timeseries_measurement table
CREATE INDEX timeseries_measurement_pkey ON timeseries_measurement(timeseries_id, time);


-- add the view back
CREATE OR REPLACE VIEW v_timeseries_latest AS (
    SELECT t.*,
       m.time AS latest_time,
	   m.value AS latest_value
    FROM v_timeseries t
    LEFT JOIN (
	    SELECT DISTINCT ON (timeseries_id) timeseries_id, time, value
	    FROM timeseries_measurement
	    ORDER BY timeseries_id, time DESC
    ) m ON t.id = m.timeseries_id
);

GRANT SELECT ON v_timeseries_latest TO instrumentation_reader;