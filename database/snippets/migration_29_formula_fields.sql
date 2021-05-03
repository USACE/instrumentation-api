    
ALTER TABLE public.instrument ADD COLUMN formula_id UUID NOT NULL DEFAULT uuid_generate_v4();
ALTER TABLE public.instrument ADD COLUMN formula_parameter_id UUID REFERENCES parameter(id) DEFAULT '2b7f96e1-820f-4f61-ba8f-861640af6232';
ALTER TABLE public.instrument ADD COLUMN formula_unit_id UUID REFERENCES unit(id) DEFAULT '4a999277-4cf5-4282-93ce-23b33c65e2c8';

CREATE OR REPLACE VIEW v_timeseries_dependency AS (
    WITH variable_tsid_map AS (
	    SELECT a.id AS timeseries_id,
               b.slug || '.' || a.slug AS variable
	    FROM timeseries a
	    LEFT JOIN instrument b ON b.id = a.instrument_id
    )
    SELECT i.instrument_id   AS instrument_id,
           i.formula_id      AS timeseries_id,
           i.parsed_variable AS parsed_variable,
           m.timeseries_id   AS dependency_timeseries_id
    FROM (
        SELECT id AS instrument_id,
            formula_id,
            (regexp_matches(formula, '\[(.*?)\]', 'g'))[1] AS parsed_variable
        FROM instrument
    ) i
    LEFT JOIN variable_tsid_map m ON m.variable = i.parsed_variable
);

GRANT SELECT ON v_timeseries_dependency TO instrumentation_reader;