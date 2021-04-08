alter table timeseries_measurement drop column id;
alter table timeseries_measurement drop constraint timeseries_unique_time;

alter table timeseries_measurement add primary key (timeseries_id,time);