CREATE TABLE IF NOT EXISTS public.telemetry_type (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR UNIQUE NOT NULL,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS public.instrument_telemetry (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument(id),
    telemetry_type_id UUID NOT NULL REFERENCES telemetry_type(id),
    telemetry_id UUID NOT NULL,
    CONSTRAINT instrument_unique_telemetry_id UNIQUE(instrument_id, telemetry_id)
);

CREATE TABLE IF NOT EXISTS public.telemetry_goes (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nesdis_id VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS public.telemetry_iridium (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    imei VARCHAR(15) UNIQUE NOT NULL
);


INSERT INTO public.telemetry_type (id, slug, name) VALUES
    ('10a32652-af43-4451-bd52-4980c5690cc9', 'goes-self-timed', 'GOES Self Timed'),
    ('c0b03b0d-bfce-453a-b5a9-636118940449', 'iridium', 'Iridium');

GRANT SELECT ON
    instrument_telemetry,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type
TO instrumentation_reader;

GRANT INSERT,UPDATE,DELETE ON
    instrument_telemetry,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type
TO instrumentation_writer;