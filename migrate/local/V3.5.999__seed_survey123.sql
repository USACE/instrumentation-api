INSERT INTO survey123 (id, project_id, name, slug, create_date, update_date, creator, updater, deleted) VALUES
('f5e1f7d2-7b1d-4b1e-8e93-d50e55b0a6b6', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Survey 1', 'survey-1', now(), now(), '57329df6-9f7a-4dad-9383-4633b452efab', NULL, false),
('a2e19d85-4c64-4e99-b93a-4f4f56a718cf', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Survey 2', 'survey-2', now(), now(), '57329df6-9f7a-4dad-9383-4633b452efab', NULL, false);

INSERT INTO survey123_equivalency_table (survey123_id, survey123_deleted, field_name, display_name, instrument_id, timeseries_id) VALUES
('f5e1f7d2-7b1d-4b1e-8e93-d50e55b0a6b6', false, 'field1', 'Field 1', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', 'da79bdb9-ded4-4f4a-8982-33e09b136815'),
('f5e1f7d2-7b1d-4b1e-8e93-d50e55b0a6b6', false, 'field2', 'Field 2', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '359bd5df-d43e-491a-871d-4701dcbff136'),
('a2e19d85-4c64-4e99-b93a-4f4f56a718cf', false, 'field3', 'Field 3', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', 'c3c00251-12fb-42a1-9d49-cdb269bb3039'),
('a2e19d85-4c64-4e99-b93a-4f4f56a718cf', false, 'field4', 'Field 4', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', 'e45a9620-a431-4b70-af97-a4e185eb7311');

INSERT INTO survey123_preview (survey123_id, preview, update_date) VALUES
('f5e1f7d2-7b1d-4b1e-8e93-d50e55b0a6b6', '{"content": "Preview content for Survey 1"}', now()),
('a2e19d85-4c64-4e99-b93a-4f4f56a718cf', '{"content": "Preview content for Survey 2"}', now());

INSERT INTO survey123_payload_error (survey123_id, error_message) VALUES
('f5e1f7d2-7b1d-4b1e-8e93-d50e55b0a6b6', 'Error message for Survey 1'),
('a2e19d85-4c64-4e99-b93a-4f4f56a718cf', 'Error message for Survey 2');
