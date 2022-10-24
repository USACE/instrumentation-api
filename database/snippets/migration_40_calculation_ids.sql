--
-- Adds formula names to instruments with a calculated representation.
-- Useful for changing the display name of the instrument when being
-- batch-plotted, as well as adding multiple formulas associated to
-- a single instrument.
--

BEGIN;

CREATE TABLE IF NOT EXISTS calculation (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),

    instrument_id UUID NOT NULL REFERENCES instrument (id),
    parameter_id UUID REFERENCES parameter (id),
    unit_id UUID REFERENCES unit (id),

    name VARCHAR(255),
    contents VARCHAR
);

GRANT SELECT ON calculation TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON calculation TO instrumentation_writer;

INSERT INTO calculation (id, instrument_id, parameter_id, unit_id, name, contents)
SELECT
    I.formula_id,
    I.id,
    I.formula_parameter_id,
    I.formula_unit_id,
    I.name || '.formula',
    I.formula
FROM instrument I;

ALTER TABLE instrument
    DROP COLUMN IF EXISTS formula_id
    DROP COLUMN IF EXISTS formula
    DROP COLUMN IF EXISTS formula_parameter_id
    DROP COLUMN IF EXISTS formula_unit_id;

COMMIT;
