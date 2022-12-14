-- fix error when attempting to delete plot config; foreign key constraint

set search_path = "$user", midas, public, topology;

BEGIN;

ALTER TABLE IF EXISTS midas.plot_configuration_settings
DROP CONSTRAINT plot_configuration_settings_id_fkey,
ADD CONSTRAINT plot_configuration_settings_id_fkey
	FOREIGN KEY (id)
    REFERENCES midas.plot_configuration (id) MATCH SIMPLE
	ON DELETE CASCADE;
	
COMMIT;
