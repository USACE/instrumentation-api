-- Add auto_range and date_range fields to plot configurtion settings

set search_path = "$user", midas, public, topology;

BEGIN;

ALTER TABLE plot_configuration_settings
    ADD auto_range BOOLEAN DEFAULT true,
    ADD date_range VARCHAR(23) DEFAULT '1 year';

UPDATE plot_configuration_settings
    SET auto_range = true
    WHERE auto_range IS NULL;

UPDATE plot_configuration_settings
    SET date_range = '1 year'
    WHERE date_range IS NULL;

-- Update other settings to store values rather than COALESCE'd in v_plot_configuration
UPDATE plot_configuration_settings
    SET show_masked = true
    WHERE show_masked IS NULL;

UPDATE plot_configuration_settings
    SET show_nonvalidated = true
    WHERE show_nonvalidated IS NULL;

UPDATE plot_configuration_settings
    SET show_comments = true
    WHERE show_comments IS NULL;

ALTER TABLE plot_configuration_settings
    ALTER COLUMN auto_range SET NOT NULL,
    ALTER COLUMN date_range SET NOT NULL,
    ALTER COLUMN show_masked SET DEFAULT true,
    ALTER COLUMN show_nonvalidated SET DEFAULT true,
    ALTER COLUMN show_comments SET DEFAULT true,
    ALTER COLUMN show_masked SET NOT NULL,
    ALTER COLUMN show_nonvalidated SET NOT NULL,
    ALTER COLUMN show_comments SET NOT NULL;

COMMIT;
