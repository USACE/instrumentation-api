drop table if exists 
    public.aware_platform_parameter_enabled,
    public.aware_platform,
    public.aware_parameter;

CREATE TABLE IF NOT EXISTS public.aware_platform (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    aware_id UUID NOT NULL,
    instrument_id UUID REFERENCES instrument(id)
);

CREATE TABLE IF NOT EXISTS public.aware_parameter (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    key VARCHAR NOT NULL,
    parameter_id UUID NOT NULL REFERENCES parameter(id),
    unit_id UUID NOT NULL REFERENCES unit(id)
);

CREATE TABLE IF NOT EXISTS public.aware_platform_parameter_enabled (
    aware_platform_id UUID NOT NULL REFERENCES aware_platform(id),
    aware_parameter_id UUID NOT NULL REFERENCES aware_parameter(id),
    CONSTRAINT aware_platform_unique_parameter UNIQUE(aware_platform_id, aware_parameter_id)
);

-- Seed Data

INSERT INTO aware_parameter (id, key, parameter_id, unit_id) VALUES
    ('1d9f9d06-6fcb-41dd-9fe4-e513a2575e74', 'depth_1', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a');
