-- The below statements are designed to quickly remove project instruments
-- and associated tables.  It may not be completely comprehensive yet.
-- Statements are using SELECT on purpose.  You have to change to DELETE
-- when ready (safety first :-) ).
-- *******************************************************************
-- Replace the project_id with YOUR target project
-- *******************************************************************

-- Delete Timeseries Measurements
select * from timeseries_measurement
where timeseries_id in (
	select t.id 
	from instrument i 
	join timeseries t on i.id = t.instrument_id
	where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')

-- Delete Timeseries
select * from timeseries
where instrument_id in (
	select i.id 
	from instrument i
	where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')
	
-- Delete Telemetry GOES
select * from telemetry_goes
where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='10a32652-af43-4451-bd52-4980c5690cc9'
	and instrument_id in (
		select i.id 
		from instrument i
		where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')
	)

-- Delete Telemetry Iridium
select * from telemetry_iridium
where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='c0b03b0d-bfce-453a-b5a9-636118940449'
	and	instrument_id in (
		select i.id 
		from instrument i
		where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')
	)
	
-- Delete Instrument Telemetry
select * from instrument_telemetry
where instrument_id in (
	select i.id 
	from instrument i
	where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')
	
-- Delete Instrument Status
select * from instrument_status
where instrument_id in (
	select i.id 
	from instrument i
	where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3')

-- Delete Instrument Group Instruments
select * from instrument_group_instruments
where instrument_id in (
	select id from instrument i
	where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3'
	)

-- Delete Instrument Groups
select * from instrument_group 
where project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3'
	
--Delete Collection Groups

--Delete collection_group_timeseries

-- Delete Instruments
select * 
from instrument i
where i.project_id = 'a6e542eb-41bc-45b3-aab7-7f45004ad8d3'

-- Delete Alert Config

-- Delete Alert