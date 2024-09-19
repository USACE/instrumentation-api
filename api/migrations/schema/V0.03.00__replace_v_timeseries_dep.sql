-- drop view needed to update column name
-- create or replace only works if col names are the same
DROP VIEW IF EXISTS v_timeseries_dependency;
