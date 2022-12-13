-- Drop foreign key constraint to accomidate calculation_id || timeseries_id
-- This is a temporary fix until the timeseries / calculation migrations are applied

ALTER TABLE midas.plot_configuration_timeseries
DROP CONSTRAINT plot_configuration_timeseries_timeseries_id_fkey;
