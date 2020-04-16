-- add slug field to instrument and instrument_group
ALTER TABLE instrument ADD slug VARCHAR(240) UNIQUE;
-- add slug field to instrument_group
ALTER TABLE instrument_group ADD slug VARCHAR(240) UNIQUE;

-- drop unique constraint on instrument name
ALTER TABLE instrument DROP CONSTRAINT instrument_name_key;
-- drop unique constraint on 
ALTER TABLE instrument_group DROP CONSTRAINT instrument_group_name_key;

