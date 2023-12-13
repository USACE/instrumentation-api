CREATE TABLE IF NOT EXISTS datalogger_table (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    table_name TEXT NOT NULL,
    CONSTRAINT datalogger_table_datalogger_id_table_name_key UNIQUE (datalogger_id, table_name)
);

-- create datalogger table for pre and post parse
INSERT INTO datalogger_table (datalogger_id, table_name)
SELECT id, '' FROM datalogger
UNION ALL
SELECT id, 'preparse' FROM datalogger;

ALTER TABLE datalogger_preview
ADD COLUMN datalogger_table_id UUID REFERENCES datalogger_table (id) ON DELETE CASCADE;

UPDATE datalogger_preview dp SET datalogger_table_id = dt.id
FROM (SELECT id, datalogger_id FROM datalogger_table) dt
WHERE dp.datalogger_id = dt.datalogger_id;

ALTER TABLE datalogger_preview
ALTER COLUMN datalogger_table_id SET NOT NULL,
ADD CONSTRAINT datalogger_preview_datalogger_table_id_key UNIQUE (datalogger_table_id),
DROP COLUMN datalogger_id;

ALTER TABLE datalogger_equivalency_table
ADD COLUMN datalogger_table_id UUID REFERENCES datalogger_table (id) ON DELETE CASCADE;

UPDATE datalogger_equivalency_table deq SET datalogger_table_id = dt.id
FROM (SELECT id, datalogger_id FROM datalogger_table WHERE table_name = '') dt
WHERE deq.datalogger_id = dt.datalogger_id;

ALTER TABLE datalogger_equivalency_table
DROP CONSTRAINT unique_datalogger_field,
ADD CONSTRAINT datalogger_equivalency_table_datalogger_table_id_field_name_key UNIQUE (datalogger_table_id, field_name);

ALTER TABLE datalogger_error
ADD COLUMN datalogger_table_id UUID REFERENCES datalogger_table (id) ON DELETE CASCADE;

UPDATE datalogger_error de SET datalogger_table_id = dt.id
FROM (SELECT id, datalogger_id FROM datalogger_table) dt
WHERE de.datalogger_id = dt.datalogger_id;
