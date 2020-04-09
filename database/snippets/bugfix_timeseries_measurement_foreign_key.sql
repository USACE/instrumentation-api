-- fix foreign key to wrong table in tables.sql

-- drop constraint
ALTER TABLE timeseries_measurement
DROP CONSTRAINT timeseries_measurement_timeseries_id_fkey;

-- add constraint
ALTER TABLE timeseries_measurement
ADD CONSTRAINT timeseries_measurement_timeseries_id_fkey
FOREIGN KEY (timeseries_id) REFERENCES public.timeseries (id);
