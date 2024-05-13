INSERT INTO report_config (id, name, slug, project_id, creator, description, date_range, date_range_enabled, show_masked, show_masked_enabled, show_nonvalidated, show_nonvalidated_enabled) VALUES
('a625f801-66e7-4d10-8d96-eb9dc55d1376', 'Test Report Config 1', 'test-report-config-1', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', 'this is a test report config', '2022-01-01 2023-01-01', true, true, true, true, true),
('a6254bce-9235-4ada-afe7-8ffc3ad867e2', 'Test Report Config 2', 'test-report-config-2', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', 'this is a test report config', NULL, false, false, false, false, false);

INSERT INTO report_config_plot_config (report_config_id, plot_config_id) VALUES
('a625f801-66e7-4d10-8d96-eb9dc55d1376', 'cc28ca81-f125-46c6-a5cd-cc055a003c19');

INSERT INTO report_download_job (id, report_config_id, creator, status, file_key, file_expiry, progress) VALUES
('e90dbcc9-7bf4-4402-80ea-c0cdbbb91c6d', 'a625f801-66e7-4d10-8d96-eb9dc55d1376', '57329df6-9f7a-4dad-9383-4633b452efab', 'SUCCESS', 'test_report.pdf', now() + INTERVAL '24 hours', 100);

INSERT INTO report_download_job (id, report_config_id, creator, status, progress) VALUES
('61b69ef2-2c73-4143-930d-3832400ba8f2', 'a625f801-66e7-4d10-8d96-eb9dc55d1376', '57329df6-9f7a-4dad-9383-4633b452efab', 'INIT', 0);
