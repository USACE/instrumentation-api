CREATE OR REPLACE FUNCTION slugify(rawname TEXT, _tbl REGCLASS)
RETURNS TEXT AS $func$
DECLARE
  newslug TEXT := '';
  n_slugs_taken INTEGER := 0;
BEGIN
  -- removes accents (diacritic signs) from a given string --
  WITH unaccented AS (
    SELECT unaccent(rawname) AS val
  ),
  -- lowercases the string
  lowercased AS (
    SELECT lower(val) AS val
    FROM unaccented
  ),
  -- remove single and double quotes
  removed_quotes AS (
    SELECT regexp_replace(val, '[''"]+', '', 'gi') AS val
    FROM lowercased
  ),
  -- replaces anything that's not a letter, number, hyphen('-'), or underscore('_') with a hyphen('-')
  hyphenated AS (
    SELECT regexp_replace(val, '[^a-z0-9\-_]+', '-', 'gi') AS val
    FROM removed_quotes
  ),
  -- trims hyphens('-') if they exist on the head or tail of the string
  trimmed AS (
    SELECT trim(BOTH '-' FROM val) AS val
    FROM hyphenated
  )
  SELECT val INTO newslug FROM trimmed;

  EXECUTE 'SELECT COUNT(*) FROM ' || _tbl  || ' t WHERE t.slug LIKE ''' || newslug || '%''' INTO n_slugs_taken;

  CASE WHEN n_slugs_taken > 0
  THEN
    RETURN newslug || '-' || n_slugs_taken::TEXT;
  ELSE
    RETURN newslug;
  END CASE;
END;
$func$ LANGUAGE plpgsql;
