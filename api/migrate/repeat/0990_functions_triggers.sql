-- +goose up
CREATE OR REPLACE FUNCTION aware_create_timeseries()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
BEGIN

INSERT INTO timeseries(instrument_id, parameter_id, unit_id, slug, name) (
	SELECT a.instrument_id AS instrument_id,
		   ap.parameter_id AS parameter_id,
		   ap.unit_id AS unit_id,
		   ap.timeseries_slug AS slug,
		   ap.timeseries_name AS name
	FROM aware_parameter ap
	CROSS JOIN aware_platform a
	WHERE ap.id = NEW.aware_parameter_id AND a.id = NEW.aware_platform_id
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;

-- Trigger; Create Timeseries when aware_platform_parameter_enabled
DROP TRIGGER IF EXISTS aware_create_timeseries ON aware_platform_parameter_enabled;
CREATE TRIGGER aware_create_timeseries
AFTER INSERT ON aware_platform_parameter_enabled
FOR EACH ROW
EXECUTE PROCEDURE aware_create_timeseries();

-- Trigger Function; Enable all AWARE parameters when new record insert into aware_platform
CREATE OR REPLACE FUNCTION aware_enable_params()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
BEGIN

INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b	
	where a.id = NEW.id
	ORDER BY aware_platform_id
	
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;

-- Trigger; Enable all AWARE parameters when new record insert into aware_platform
DROP TRIGGER IF EXISTS aware_enable_params ON aware_platform;
CREATE TRIGGER aware_enable_params
AFTER INSERT ON aware_platform
FOR EACH ROW 
EXECUTE PROCEDURE aware_enable_params();

/* 
######################################################################
Trigger to create group when instrument is inserted/updated if:
1) Instrument has a federal id assigned
2) Federal ID group doesn't already exist for the project
3) Instrument is not already assigned to the Federal ID group
*/
/*
CREATE OR REPLACE FUNCTION create_federal_id_instrument_group()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
    declare groupId uuid;
    declare oldGroupId uuid;

    BEGIN
        
            -- NEW instrument INSERT
            -- If federal id set on new instrument
            IF TG_OP = 'INSERT' and NEW.nid_id is not NULL THEN
        
                -- Group may or may not exist
                -- Create the group, gracefully fail if already exists
                INSERT INTO instrument_group(slug, name, description, project_id) 
                VALUES (            
                    lower(NEW.nid_id),
                    NEW.nid_id,
                    CONCAT(NEW.nid_id, ' automated group'),
                    NEW.project_id
                ) ON CONFLICT DO nothing;
                --returning id into groupId;

                -- Regardless if new group was created or it already existed,
                -- get the group id
                SELECT id INTO groupId from instrument_group
                WHERE project_id = NEW.project_id and slug = lower(NEW.nid_id);
        
            
                -- If new groupId found, insert record in instrument_group_instruments
                if groupId is not NULL THEN
                    INSERT INTO instrument_group_instruments(instrument_id, instrument_group_id) 
                    VALUES (NEW.id, groupId)
                        ON CONFLICT DO NOTHING;
                end if;  
        
            END IF;	   
        
        
        -- Instrument UPDATE
        -- If federal id has changed.  Could be old vs new or null vs new
        IF TG_OP = 'UPDATE' and NEW.nid_id is not NULL and OLD.nid_id != NEW.nid_id then
        
        		-- get the old group id (if it exists)
               	SELECT id INTO oldGroupId from instrument_group
               	WHERE project_id = NEW.project_id and slug = lower(OLD.nid_id);                
        
        		-- Group may or may not exist
                -- Create the group, gracefully fail if already exists
                INSERT INTO instrument_group(slug, name, description, project_id) 
                VALUES (            
                    lower(NEW.nid_id),
                    NEW.nid_id,
                    CONCAT(NEW.nid_id, ' automated group'),
                    NEW.project_id
                ) ON CONFLICT DO nothing;
                --returning id into groupId;

                -- Regardless if new group was created or it already existed,
                -- get the new group id
                SELECT id INTO groupId from instrument_group
                WHERE project_id = NEW.project_id and slug = lower(NEW.nid_id);
            
                 -- If new groupId found, insert record in instrument_group_instruments
                if groupId is not NULL THEN
                    INSERT INTO instrument_group_instruments(instrument_id, instrument_group_id) 
                    VALUES (NEW.id, groupId)
                    ON CONFLICT DO NOTHING;
                end if;               

                -- if the previous federal id was set on the instrument was not null,
                -- remove the instrument from the group
                if OLD.nid_id is not null and oldGroupId is not null THEN
                    DELETE from instrument_group_instruments
                    WHERE instrument_id = OLD.id
                    AND instrument_group_id = oldGroupId;
                end if;
        
            END IF;
        
        RETURN NEW;
        
    END;	
    $$;

-- Trigger; Create instrument_group when
DROP TRIGGER IF EXISTS create_federal_id_instrument_group ON instrument;
CREATE TRIGGER create_federal_id_instrument_group
AFTER insert or UPDATE ON instrument
FOR EACH ROW
EXECUTE PROCEDURE create_federal_id_instrument_group();
*/
/*
######################################################################
*/
