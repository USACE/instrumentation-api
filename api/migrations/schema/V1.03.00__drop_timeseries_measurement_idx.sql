-- large index created during troubleshooting
-- this was for the explorer / calculated timeseries endpoints
-- other (smaller) indexes were created afterwards that solve
-- those quesry performance issues
DROP INDEX IF EXISTS timeseries_measurement_btree_idx;

-- this has little benefit due to the amout of manually uploaded
-- and backfilled timeseries data, only works well if there is
-- correlation between insertion order and index (time)
DROP INDEX IF EXISTS timeseries_measurement_time_brin_idx;
