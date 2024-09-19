ALTER TABLE profile ADD COLUMN display_name text;

UPDATE profile SET display_name = username;

ALTER TABLE profile ALTER COLUMN display_name SET NOT NULL;
