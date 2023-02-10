-- Datalogger
CREATE TABLE IF NOT EXISTS datalogger (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    sn TEXT NOT NULL,
    project_id UUID NOT NULL REFERENCES project(id),
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    update_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    model TEXT NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT false,
    CONSTRAINT datalogger_deleted_uni UNIQUE (id, deleted)
);

CREATE TABLE IF NOT EXISTS datalogger_hash (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    "hash" TEXT NOT NULL,
    CONSTRAINT unique_hash UNIQUE("hash"),
    CONSTRAINT unique_datalogger_hash UNIQUE(datalogger_id, "hash")
);

-- Datalogger Preview
CREATE TABLE IF NOT EXISTS datalogger_preview (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    payload JSON
);

-- Datalogger Field Instrument Timeseries Mapper
CREATE TABLE IF NOT EXISTS datalogger_equivalency_table (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    datalogger_id UUID NOT NULL,
    datalogger_deleted BOOLEAN,
    field_name TEXT NOT NULL,
    display_name TEXT,
    instrument_id UUID REFERENCES instrument (id) ON DELETE CASCADE,
    timeseries_id UUID REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT unique_datalogger_field UNIQUE(datalogger_id, field_name),
    CONSTRAINT fk2 FOREIGN KEY (datalogger_id, datalogger_deleted) REFERENCES datalogger (id, deleted)
);

-- timeseries must be unique, but only for dataloggers that are not flagged as "deleted".
-- To reference the "deleted" column of the datalogger table when creating this unique partial index,
-- a multi-column foreign key is needed: https://stackoverflow.com/a/35570262
CREATE UNIQUE INDEX datalogger_timeseries_uni_idx ON datalogger_equivalency_table (timeseries_id)
WHERE NOT datalogger_deleted;
