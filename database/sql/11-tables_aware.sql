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
    unit_id UUID NOT NULL REFERENCES unit(id),
    timeseries_slug VARCHAR NOT NULL,
    timeseries_name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS public.aware_platform_parameter_enabled (
    aware_platform_id UUID NOT NULL REFERENCES aware_platform(id),
    aware_parameter_id UUID NOT NULL REFERENCES aware_parameter(id),
    CONSTRAINT aware_platform_unique_parameter UNIQUE(aware_platform_id, aware_parameter_id)
);

CREATE OR REPLACE VIEW v_aware_platform_parameter_enabled AS (
    SELECT i.project_id  AS project_id,
	       i.id          AS instrument_id,
		   a.aware_id    AS aware_id,
		   b.key         AS aware_parameter_key,
		   t.id          AS timeseries_id
    FROM aware_platform_parameter_enabled e
    INNER JOIN aware_platform a ON a.id = e.aware_platform_id
    INNER JOIN instrument i ON i.id = a.instrument_id
    INNER JOIN aware_parameter b ON b.id = e.aware_parameter_id
	LEFT JOIN timeseries t ON t.instrument_id=i.id AND t.parameter_id=b.parameter_id AND t.unit_id=b.unit_id
	ORDER BY i.project_id, a.aware_id
);

-- Seed Data
-- aware_parameter
INSERT INTO aware_parameter (id, key, parameter_id, unit_id, timeseries_slug, timeseries_name) VALUES
    ('1d9f9d06-6fcb-41dd-9fe4-e513a2575e74', 'depth1', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'stage', 'Stage'),
    ('c5f2842d-a5a9-4f53-9583-f613080a9c36', 'battery', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64', 'battery-voltage', 'Battery Voltage'),
    ('78d32638-5137-481c-aa9d-a48c2d57824a', 'baro', '1de79e29-fb70-45c3-ae7d-4695517ced90', '55dda9ef-7ba6-4432-b64d-8ef0e65154f4', 'barometric-pressure', 'Barometric Pressure'),
    ('53ce89a7-1db8-45bd-baf9-9536d75d7046', 'h2oTemp', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc', 'water-temperature', 'Water Temperature'),
    ('3ca7c1da-7124-42c0-b92c-f76b5c318b0c', 'rssi', 'b23b141d-ce7b-4e0d-82ab-c8beb39c8325', 'be854f6e-e36e-4bba-9e06-6d5aa09485be', 'signal-strength', 'Signal Strength');

-- Trigger Function; Create Timeseries when aware_platform_parameter_enabled
CREATE OR REPLACE FUNCTION public.aware_create_timeseries()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
BEGIN

INSERT INTO timeseries(instrument_id, parameter_id, unit_id, slug, name) (
	SELECT a.instrument_id AS instrument_id,
		   ap.parameter_id AS parameter_id,
		   ap.unit_id AS unit_id,
		   ap.timeseries_slug AS slug,
		   ap.timeseries_name AS name
	FROM aware_parameter ap
	CROSS JOIN aware_platform a
	WHERE ap.id = NEW.aware_parameter_id AND a.id = NEW.aware_platform_id
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;

-- Trigger; Create Timeseries when aware_platform_parameter_enabled
CREATE TRIGGER aware_create_timeseries
AFTER INSERT ON public.aware_platform_parameter_enabled
FOR EACH ROW
EXECUTE PROCEDURE public.aware_create_timeseries();


-- Trigger Function; Enable all AWARE parameters when new record insert into aware_platform
CREATE OR REPLACE FUNCTION public.aware_enable_params()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
BEGIN

INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b	
	where a.id = NEW.id
	ORDER BY aware_platform_id
	
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;

-- Trigger; Enable all AWARE parameters when new record insert into aware_platform
CREATE TRIGGER aware_enable_params
AFTER INSERT ON public.aware_platform
FOR EACH ROW 
EXECUTE PROCEDURE public.aware_enable_params();



-- Seed Data

-- aware_platform
-- aware_id used below is not an actual aware_id; generated for testing
INSERT INTO aware_platform (id, aware_id, instrument_id) VALUES
    ('b896ce34-2bd4-436c-9f28-7a1eefb5744a', '6df213c4-a582-4735-a916-6f4065082872', 'a7540f69-c41e-43b3-b655-6e44097edb7e');

-- enable all parameters for sample aware platform
INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b
	ORDER BY aware_platform_id
    
)
ON CONFLICT DO NOTHING;