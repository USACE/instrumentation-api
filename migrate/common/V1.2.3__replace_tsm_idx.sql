DROP INDEX IF EXISTS timeseries_measurement_btree_idx;
CREATE INDEX timeseries_measurement_btree_idx ON timeseries_measurement (timeseries_id, time) INCLUDE (value);
