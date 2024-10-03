INSERT INTO plot_configuration (id, project_id, slug, name, plot_type) VALUES
('9a9c6a73-b0c0-48d2-bb96-5aac9f4aa7a5', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-contour-plot-config', 'Test Contour Plot Config', 'contour'),
('94df34f5-ba00-4c3d-bfa7-f128a00166be', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-contour-plot-config-2', 'Test Contour Plot Config 2', 'contour'),
('9ea45347-5a70-47d4-bb0d-d31885a836ca', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-bullseye-plot-config', 'Test Bullseye Plot Config', 'bullseye'),
('871e34da-c911-4d8f-ab68-e29ba17f8937', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-bullseye-plot-config-2', 'Test Bullseye Plot Config 2', 'bullseye'),
('af0c637e-76dc-4af8-8844-7f5b98e91248', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-profile-plot-config-saa', 'Test Profile Plot Config SAA', 'profile'),
('f08a15fa-448c-4066-a81f-124d54712c62', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'test-profile-plot-config-ipi', 'Test Profile Plot Config IPI', 'profile');

INSERT INTO plot_contour_config (plot_config_id, "time", locf_backfill) VALUES
('9a9c6a73-b0c0-48d2-bb96-5aac9f4aa7a5', now(), '1 month'::interval),
('94df34f5-ba00-4c3d-bfa7-f128a00166be', now(), '1 week'::interval);

INSERT INTO plot_contour_config_timeseries (plot_contour_config_id, timeseries_id) VALUES
('9a9c6a73-b0c0-48d2-bb96-5aac9f4aa7a5', '00ae950d-5bdd-455e-a72a-56da67dafb85'),
('9a9c6a73-b0c0-48d2-bb96-5aac9f4aa7a5', '5842c707-b4be-4d10-a89c-1064e282e555');

INSERT INTO plot_bullseye_config (plot_config_id, x_axis_timeseries_id, y_axis_timeseries_id) VALUES
('9ea45347-5a70-47d4-bb0d-d31885a836ca', '21cfe121-d29d-40a2-b04f-6be71ba479fe', '23bda2f6-c479-48e0-a0c2-db48c3b08c3c'),
('871e34da-c911-4d8f-ab68-e29ba17f8937', '2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70', '4759bdac-656e-47c3-b403-d3118cf57342');

INSERT INTO plot_profile_config (plot_config_id, instrument_id) VALUES
('af0c637e-76dc-4af8-8844-7f5b98e91248', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6'),
('f08a15fa-448c-4066-a81f-124d54712c62', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929');
