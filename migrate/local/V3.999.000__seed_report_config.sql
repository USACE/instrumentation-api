INSERT INTO report_config (id, name, slug, project_id, after, before, creator, description) VALUES
('a625f801-66e7-4d10-8d96-eb9dc55d1376', 'Test Report Config 1', 'test-report-config-1', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', now() - INTERVAL '1 year', now(), '57329df6-9f7a-4dad-9383-4633b452efab', 'this is a test report config'),
('a6254bce-9235-4ada-afe7-8ffc3ad867e2', 'Test Report Config 2', 'test-report-config-2', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', NULL, NULL, '57329df6-9f7a-4dad-9383-4633b452efab', 'this is a test report config');

INSERT INTO report_config_plot_config (report_config_id, plot_config_id) VALUES
('a625f801-66e7-4d10-8d96-eb9dc55d1376', 'cc28ca81-f125-46c6-a5cd-cc055a003c19');