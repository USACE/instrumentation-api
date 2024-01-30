INSERT INTO datalogger (id, sn, project_id, creator, create_date, updater, update_date, name, slug, model_id, deleted) VALUES
    ('83a7345c-62d8-4e29-84db-c2e36f8bc40d', '12345', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:20:18.223202+00', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:20:18.223202+00', 'Test Data Logger 1', 'test-data-logger-1', '6a10ef5f-b9d9-4fa0-8b1e-ea1bcc81748c', false),
    ('c0b65315-f802-4ca5-a4dd-7e0cfcffd057', '56789', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:31:40.678486+00', '57329df6-9f7a-4dad-9383-4633b452efab', '2023-02-10 00:31:40.678486+00', 'Test Data Logger 2', 'test-data-logger-2', 'f0d4effa-50dc-44e4-9a9b-cb8181c8e7e0', false);

-- mock data logger hash for local
INSERT INTO datalogger_hash (datalogger_id, "hash") VALUES
    ('83a7345c-62d8-4e29-84db-c2e36f8bc40d', '$argon2id$v=19$m=65536,t=3,p=2$Ud3epc4MAsAAN6pCUQxCjh$5mXpCpcc3nc46XUC7pJiAKvK9UzK1qfP6GoYDFH66zv4pjQxt88ybQJ'),
    ('c0b65315-f802-4ca5-a4dd-7e0cfcffd057', '$argon2id$v=19$m=65536,t=3,p=2$Ud3epc4MAsAAN6pCUQxCjh$5mXpCpcc3nc46XUC7pJiAKvK9UzK1qfP6GoYDFH66zv4pjQxt88ybQJ');

INSERT INTO datalogger_table (id, datalogger_id, table_name) VALUES
    ('98a77c65-e5c4-49ed-8fb4-b0ffd06add4c', '83a7345c-62d8-4e29-84db-c2e36f8bc40d', 'Demo Datalogger Table'),
    ('5b47be95-6ba9-488c-bdda-be7bdbca2909', 'c0b65315-f802-4ca5-a4dd-7e0cfcffd057', '');

INSERT INTO datalogger_preview (datalogger_table_id, update_date, preview) VALUES
    ('98a77c65-e5c4-49ed-8fb4-b0ffd06add4c', '2023-02-16 18:53:00.582812+00', '{"data":[{"no":158,"time":"2023-02-16T18:53:00","vals":[12.13,22.37]},{"no":158,"time":"2023-02-16T18:52:55","vals":[12.25,20.32]}],"head":{"environment":{"model":"CR6","os_version":"CR6.Std.12.01","prog_name":"CPU:Updated_CR6_Sample_Template.CR6","serial_no":"11111","station_name":"6239","table_name":"Test"},"fields":[{"name":"batt_volt_Min","process":"Min","settable":false,"type":"xsd:float","units":"Volts"},{"name":"PanelT","process":"Smp","settable":false,"type":"xsd:float","units":"Deg_C"}],"signature":20883,"transaction":0}}'),
    ('5b47be95-6ba9-488c-bdda-be7bdbca2909', '2023-02-16 18:53:00.582812+00', NULL);

INSERT INTO datalogger_equivalency_table (id, datalogger_table_id, datalogger_id, datalogger_deleted, field_name, display_name, instrument_id, timeseries_id) VALUES
    ('40ceff10-cdc3-4715-a4ca-c1e570fe25de', '98a77c65-e5c4-49ed-8fb4-b0ffd06add4c', '83a7345c-62d8-4e29-84db-c2e36f8bc40d', false, 'batt_volt_Min', 'Battery Voltage', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'),
    ('2f1f7c3d-8b6f-4b11-917e-8f049eb6c62b', '98a77c65-e5c4-49ed-8fb4-b0ffd06add4c', '83a7345c-62d8-4e29-84db-c2e36f8bc40d', false, 'PanelT', 'Panel Temperature', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'da79bdb9-ded4-4f4a-8982-33e09b136815');
