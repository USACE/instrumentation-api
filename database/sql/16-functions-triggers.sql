/* 
######################################################################
Trigger to create group when instrument is inserted/updated if:
1) Instrument has a federal id assigned
2) Federal ID group doesn't already exist for the project
3) Instrument is not already assigned to the Federal ID group
*/
CREATE OR REPLACE FUNCTION public.create_federal_id_instrument_group()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
    declare groupId uuid;

    BEGIN
        
            -- If federal id set on new instrument
            IF TG_OP = 'INSERT' and NEW.nid_id is not NULL THEN
        
                INSERT INTO instrument_group(slug, name, description, project_id) 
                VALUES (            
                    lower(NEW.nid_id),
                    NEW.nid_id,
                    CONCAT(NEW.nid_id, ' automated group'),
                    NEW.project_id
                ) ON CONFLICT DO nothing
            returning id into groupId;
        
            
            if groupId is not NULL THEN
                INSERT INTO instrument_group_instruments(instrument_id, instrument_group_id) 
                VALUES (NEW.id, groupId)
                    ON CONFLICT DO NOTHING;
            end if;  
        
            END IF;	   
        
        
        -- If federal id set on new instrument
        IF TG_OP = 'UPDATE' and NEW.nid_id is not NULL THEN
        
                INSERT INTO instrument_group(slug, name, description, project_id) 
                VALUES (            
                    lower(NEW.nid_id),
                    NEW.nid_id,
                    CONCAT(NEW.nid_id, ' automated group'),
                    NEW.project_id
                ) ON CONFLICT DO nothing
            returning id into groupId;
            
            
            if groupId is not NULL THEN
                INSERT INTO instrument_group_instruments(instrument_id, instrument_group_id) 
                VALUES (NEW.id, groupId)
                    ON CONFLICT DO NOTHING;
            end if;
        
            END IF;
        
        RETURN NEW;
        
    END;	
    $$;

-- Trigger; Create instrument_group when
CREATE TRIGGER create_federal_id_instrument_group
AFTER insert or UPDATE ON public.instrument
FOR EACH ROW
EXECUTE PROCEDURE public.create_federal_id_instrument_group();
/*
######################################################################
*/