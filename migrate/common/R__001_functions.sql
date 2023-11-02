CREATE OR REPLACE FUNCTION locf_s(a float, b float)
RETURNS float
LANGUAGE SQL
AS '
  SELECT COALESCE(b, a)
';

DROP AGGREGATE IF EXISTS locf(float);
CREATE AGGREGATE locf(float) (
  sfunc = locf_s,
  stype = float
);
