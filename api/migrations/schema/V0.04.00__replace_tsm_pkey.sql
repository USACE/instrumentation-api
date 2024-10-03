-- Adding brin index due to natural correlation of time and order of columns inserted
CREATE INDEX timeseries_measurement_time_brin_idx ON timeseries_measurement 
    USING BRIN (time);

-- Adding value to timeseries index allows for faster index-only scans
CREATE INDEX timeseries_measurement_btree_idx ON timeseries_measurement (timeseries_id, time, value);
