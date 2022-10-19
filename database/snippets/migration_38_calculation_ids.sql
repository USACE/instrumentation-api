--
-- Adds formula names to instruments with a calculated representation.
-- Useful for changing the display name of the instrument when being
-- batch-plotted.
--

ALTER TABLE instrument
    DROP COLUMN IF EXISTS formula_id
    DROP COLUMN IF EXISTS formula
    DROP COLUMN IF EXISTS formula_parameter_id
    DROP COLUMN IF EXISTS formula_unit_id;

CREATE TABLE IF NOT EXISTS calculation (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),

    instrument_id UUID NOT NULL REFERENCES instrument (id),
    parameter_id UUID REFERENCES parameter (id),
    unit_id UUID REFERENCES unit (id),

    name VARCHAR(255),
    contents VARCHAR
);
