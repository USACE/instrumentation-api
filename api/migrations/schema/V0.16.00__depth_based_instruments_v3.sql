ALTER TABLE ipi_segment ADD COLUMN temp_timeseries_id UUID REFERENCES timeseries (id);
ALTER TABLE ipi_segment RENAME COLUMN cum_dev_timeseries_id TO inc_dev_timeseries_id;
