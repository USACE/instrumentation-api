INSERT INTO alert_config (id, project_id, name, body, creator, create_date, updater, update_date, alert_type_id, last_checked, deleted) VALUES 
('5f9c9b14-dc4e-4a97-99e2-b6e9088a4f23', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Threshold Alert 1', 'Alert for threshold', '57329df6-9f7a-4dad-9383-4633b452efab', now(), NULL, NULL, 'bb15e7c2-8eae-452c-92f7-e720dc5c9432', NULL, false),
('7fae1097-4e96-453f-bf36-7a3427cfb0d7', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Threshold Alert 2', 'Alert for another threshold', '57329df6-9f7a-4dad-9383-4633b452efab', now(), NULL, NULL, 'bb15e7c2-8eae-452c-92f7-e720dc5c9432', NULL, false),
('d76df8f4-6f49-4f02-9bdf-cf55cfeede1f', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Rate of Change Alert 1', 'Alert for rate of change', '57329df6-9f7a-4dad-9383-4633b452efab', now(), NULL, NULL, 'c37effee-6b48-4436-8d72-737ed78c1fb7', NULL, false),
('8dc1d10e-938e-4b27-8ba8-2fd5bb957ccb', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Rate of Change Alert 2', 'Another alert for rate of change', '57329df6-9f7a-4dad-9383-4633b452efab', now(), NULL, NULL, 'c37effee-6b48-4436-8d72-737ed78c1fb7', NULL, false);

INSERT INTO alert_config_timeseries (alert_config_id, timeseries_id) VALUES 
('5f9c9b14-dc4e-4a97-99e2-b6e9088a4f23', '869465fc-dc1e-445e-81f4-9979b5fadda9'),
('5f9c9b14-dc4e-4a97-99e2-b6e9088a4f23', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'),
('7fae1097-4e96-453f-bf36-7a3427cfb0d7', '9a3864a8-8766-4bfa-bad1-0328b166f6a8'),
('7fae1097-4e96-453f-bf36-7a3427cfb0d7', '7ee902a3-56d0-4acf-8956-67ac82c03a96'),
('d76df8f4-6f49-4f02-9bdf-cf55cfeede1f', '869465fc-dc1e-445e-81f4-9979b5fadda9'),
('d76df8f4-6f49-4f02-9bdf-cf55cfeede1f', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'),
('8dc1d10e-938e-4b27-8ba8-2fd5bb957ccb', '9a3864a8-8766-4bfa-bad1-0328b166f6a8'),
('8dc1d10e-938e-4b27-8ba8-2fd5bb957ccb', '7ee902a3-56d0-4acf-8956-67ac82c03a96');

INSERT INTO alert_config_threshold (alert_config_id, alert_low_value, alert_high_value, warn_low_value, warn_high_value, ignore_low_value, ignore_high_value, variance) VALUES 
('5f9c9b14-dc4e-4a97-99e2-b6e9088a4f23', 10.0, 50.0, 15.0, 45.0, 5.0, 55.0, 1.0),
('7fae1097-4e96-453f-bf36-7a3427cfb0d7', 20.0, 60.0, 25.0, 55.0, 10.0, 65.0, 2.0);

INSERT INTO alert_config_change (alert_config_id, warn_rate_of_change, alert_rate_of_change, locf_backfill) VALUES 
('d76df8f4-6f49-4f02-9bdf-cf55cfeede1f', 5.0, 10.0, '1 hour'),
('8dc1d10e-938e-4b27-8ba8-2fd5bb957ccb', 8.0, 15.0, '30 minutes');
