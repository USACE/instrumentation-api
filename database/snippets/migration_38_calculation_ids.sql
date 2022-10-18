--
-- Adds formula names to instruments with a calculated representation.
-- Useful for changing the display name of the instrument when being
-- batch-plotted.
--

ALTER TABLE instrument ADD COLUMN formula_name VARCHAR;
