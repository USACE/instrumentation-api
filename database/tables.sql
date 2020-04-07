-- instrument_type
CREATE TABLE IF NOT EXISTS public.instrument_type (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- unit
CREATE TABLE IF NOT EXISTS public.unit (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- parameter
CREATE TABLE IF NOT EXISTS public.parameter (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- instrument_group
CREATE TABLE IF NOT EXISTS public.instrument_group (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL,
    description VARCHAR(360)
	);

-- instrument
CREATE TABLE IF NOT EXISTS public.instrument (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL,
    height REAL,
    geometry geometry,
    instrument_type_id UUID REFERENCES instrument_type (id)
);

-- instrument_group_instruments
CREATE TABLE IF NOT EXISTS public.instrument_group_instruments (
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    instrument_group_id UUID NOT NULL REFERENCES instrument_group (id),
    UNIQUE (instrument_id, instrument_group_id)
);

-- timeseries
CREATE TABLE IF NOT EXISTS public.timeseries (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(240) UNIQUE NOT NULL,
    instrument_id UUID REFERENCES instrument (id),
    parameter_id UUID NOT NULL REFERENCES parameter (id),
    unit_id UUID NOT NULL REFERENCES unit (id)
);

-- timeseries_measurement
CREATE TABLE IF NOT EXISTS public.timeseries_measurement (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    time TIMESTAMPTZ NOT NULL,
    value REAL NOT NULL,
    timeseries_id UUID NOT NULL REFERENCES unit (id)
);
