drop table if exists 
    public.aware_platform_parameter_enabled,
    public.aware_platform,
    public.aware_parameter;

CREATE TABLE IF NOT EXISTS public.aware_platform (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    platform_id VARCHAR NOT NULL,
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
    ('1d9f9d06-6fcb-41dd-9fe4-e513a2575e74', 'depth1', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
    ('c5f2842d-a5a9-4f53-9583-f613080a9c36', 'battery', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
    ('78d32638-5137-481c-aa9d-a48c2d57824a', 'baro', '1de79e29-fb70-45c3-ae7d-4695517ced90', '55dda9ef-7ba6-4432-b64d-8ef0e65154f4'),
    ('53ce89a7-1db8-45bd-baf9-9536d75d7046', 'h2oTemp', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc');
