
SET search_path TO midas,topology,public;

CREATE SCHEMA midas;

-- RUN ALL CREATE TABLE SCRIPTS CAREFULLY


INSERT INTO midas.project (id, image, office_id, deleted, slug, name, creator, create_date, updater, update_date) (
	SELECT id, image, office_id, deleted, slug, name, creator, create_date, updater, update_date FROM public.project WHERE NAME ILIKE '%streamgages%' AND NAME NOT ILIKE '%northwestern division%'
);

-- Profiles
INSERT INTO midas.profile (SELECT * FROM public.profile) ON CONFLICT DO NOTHING;
INSERT INTO midas.profile_project_roles (SELECT * FROM public.profile_project_roles WHERE project_id IN (SELECT ID FROM midas.project)) ON CONFLICT DO NOTHING;
INSERT INTO midas.profile_token (SELECT * FROM public.profile_token WHERE profile_id IN (SELECT ID FROM midas.profile)) ON CONFLICT DO NOTHING;
-- Instrument
INSERT INTO midas.INSTRUMENT (id, deleted, slug, name, formula_id, formula, formula_parameter_id, formula_unit_id, geometry, station, station_offset, creator, create_date, updater, update_date, type_id, project_id, nid_id, usgs_id) (
	SELECT id, deleted, slug, name, formula_id, formula, formula_parameter_id, formula_unit_id, geometry, station,
	       station_offset, creator, create_date, updater, update_date, type_id, project_id, nid_id, usgs_id
	FROM public.INSTRUMENT WHERE project_id IN (SELECT id FROM midas.project)
) ON CONFLICT DO NOTHING;
INSERT INTO midas.instrument_constants ( SELECT * FROM public.instrument_constants);
INSERT INTO midas.instrument_group (SELECT * FROM public.instrument_group WHERE PROJECT_ID IN (SELECT ID FROM MIDAS.PROJECT)) ON CONFLICT DO NOTHING;
INSERT INTO midas.instrument_group_instruments (SELECT * FROM public.instrument_group_instruments WHERE instrument_id IN (SELECT id FROM midas.instrument)) ON CONFLICT DO NOTHING;
INSERT INTO midas.instrument_note (SELECT * FROM public.instrument_note WHERE instrument_id IN (SELECT id FROM midas.instrument)) ON CONFLICT DO NOTHING;
INSERT INTO midas.instrument_status (SELECT * FROM public.instrument_status WHERE instrument_id IN (SELECT id FROM midas.instrument)) ON CONFLICT DO NOTHING;
-- Telemetry
INSERT INTO midas.instrument_telemetry (SELECT * FROM public.instrument_telemetry WHERE instrument_id IN (SELECT id FROM midas.instrument)) ON CONFLICT DO NOTHING;
INSERT INTO midas.telemetry_goes (SELECT * FROM public.telemetry_goes) ON CONFLICT DO NOTHING;
INSERT INTO midas.telemetry_iridium (SELECT * FROM public.telemetry_iridium) ON CONFLICT DO NOTHING;

-- Timeseries
INSERT INTO midas.timeseries (SELECT * FROM public.timeseries WHERE instrument_id IN (SELECT id FROM midas.instrument)) ON CONFLICT DO NOTHING;
INSERT INTO midas.timeseries_measurement (SELECT * FROM public.timeseries_measurement WHERE timeseries_id IN (SELECT id FROM midas.timeseries)) ON CONFLICT DO NOTHING;
-- Collection Groups
INSERT INTO midas.collection_group (SELECT * FROM public.collection_group WHERE project_id IN (SELECT ID from midas.project)) ON CONFLICT DO NOTHING;
INSERT INTO midas.collection_group_timeseries (SELECT * FROM public.collection_group_timeseries WHERE timeseries_id IN (SELECT ID from midas.timeseries)) ON CONFLICT DO NOTHING;

-- Plot Configurations
INSERT INTO midas.plot_configuration (SELECT * FROM public.plot_configuration WHERE project_id IN (SELECT ID FROM MIDAS.PROJECT)) ON CONFLICT DO NOTHING;
INSERT INTO midas.plot_configuration_timeseries (SELECT * FROM public.plot_configuration_timeseries WHERE timeseries_id IN (SELECT ID FROM MIDAS.timeseries)) ON CONFLICT DO NOTHING;

