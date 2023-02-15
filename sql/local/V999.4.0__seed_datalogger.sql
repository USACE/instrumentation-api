INSERT INTO datalogger (id, sn, project_id, creator, create_date, updater, update_date, name, slug, model_id, deleted) VALUES
    ('83a7345c-62d8-4e29-84db-c2e36f8bc40d', '12345', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:20:18.223202+00', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:20:18.223202+00', 'Test Data Logger 1', 'test-data-logger-1', '6a10ef5f-b9d9-4fa0-8b1e-ea1bcc81748c', false),
    ('c0b65315-f802-4ca5-a4dd-7e0cfcffd057', '56789', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:31:40.678486+00', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:31:40.678486+00', 'Test Data Logger 2', 'test-data-logger-2', 'f0d4effa-50dc-44e4-9a9b-cb8181c8e7e0', false);

INSERT INTO datalogger_equivalency_table (id, datalogger_id, datalogger_deleted, field_name, display_name, instrument_id, timeseries_id) VALUES
    ('40ceff10-cdc3-4715-a4ca-c1e570fe25de', '83a7345c-62d8-4e29-84db-c2e36f8bc40d', false, 'field name 1', 'test 1', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '7ee902a3-56d0-4acf-8956-67ac82c03a96'),
    ('2f1f7c3d-8b6f-4b11-917e-8f049eb6c62b', '83a7345c-62d8-4e29-84db-c2e36f8bc40d', false, 'field name 2', 'test 2', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'd9697351-3a38-4194-9ac4-41541927e475');
