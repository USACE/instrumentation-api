CREATE TABLE IF NOT EXISTS datalogger_model (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    model TEXT
);

INSERT INTO datalogger_model (id, model) VALUES
    ('6a10ef5f-b9d9-4fa0-8b1e-ea1bcc81748c', 'CR6'),
    ('f0d4effa-50dc-44e4-9a9b-cb8181c8e7e0', 'CR1000X');

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
    model_id UUID NOT NULL REFERENCES datalogger_model (id),
    deleted BOOLEAN NOT NULL DEFAULT false,
    CONSTRAINT unique_datalogger_deleted UNIQUE (id, deleted)
);

CREATE UNIQUE INDEX unique_idx_datalogger_sn_model ON datalogger (sn, model_id)
WHERE NOT deleted;

-- Datalogger Hash
CREATE TABLE IF NOT EXISTS datalogger_hash (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    "hash" TEXT NOT NULL,
    CONSTRAINT unique_datalogger_hash UNIQUE(datalogger_id, "hash")
);

-- Datalogger Preview
CREATE TABLE IF NOT EXISTS datalogger_preview (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    preview JSON,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Datalogger Field Instrument Timeseries Mapper
CREATE TABLE IF NOT EXISTS datalogger_equivalency_table (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    datalogger_id UUID NOT NULL,
    datalogger_deleted BOOLEAN NOT NULL DEFAULT false,
    field_name TEXT NOT NULL,
    display_name TEXT,
    instrument_id UUID REFERENCES instrument (id) ON DELETE CASCADE,
    timeseries_id UUID REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT unique_datalogger_field UNIQUE(datalogger_id, field_name),
    CONSTRAINT unique_active_datalogger FOREIGN KEY (datalogger_id, datalogger_deleted)
        REFERENCES datalogger (id, deleted) ON UPDATE CASCADE
);

-- timeseries must be unique, but only for dataloggers that are not flagged as "deleted".
-- To reference the "deleted" column of the datalogger table when creating this unique partial index,
-- a multi-column foreign key is needed: https://stackoverflow.com/a/35570262
CREATE UNIQUE INDEX unique_idx_datalogger_timeseries ON datalogger_equivalency_table (timeseries_id)
WHERE NOT datalogger_deleted;

CREATE TABLE IF NOT EXISTS datalogger_error (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id) ON DELETE CASCADE,
    error_message TEXT
);
