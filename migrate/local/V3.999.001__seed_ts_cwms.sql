INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name, type) VALUES
('47afea78-4169-499c-be51-013ca3b53cba', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'test-cwms-timeseries', 'Test CWMS Timeseries', 'cwms');

INSERT INTO timeseries_cwms (timeseries_id, cwms_timeseries_id, cwms_office_id) VALUES
('47afea78-4169-499c-be51-013ca3b53cba', 'LRD', '03049500.Flow-Holdout.Inst.1Day.0.Raw-LRP');
