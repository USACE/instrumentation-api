-- Datalogger
CREATE TABLE IF NOT EXISTS datalogger (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    sn TEXT NOT NULL,
    project_id UUID NOT NULL REFERENCES project(id),
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    model TEXT NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT false,
    CONSTRAINT datalogger_unique_sn UNIQUE(sn)
);

-- Datalogger Preview
CREATE TABLE IF NOT EXISTS datalogger_preview (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    payload JSON
);

-- Datalogger Field Instrument Timeseries Mapper
CREATE TABLE IF NOT EXISTS datalogger_field_instrument_timeseries (
    datalogger_id UUID NOT NULL REFERENCES datalogger (id),
    field_name TEXT NOT NULL,
    display_name TEXT,
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    timeseries_id UUID NOT NULL REFERENCES timeseries (id)
);
