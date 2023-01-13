CREATE extension IF NOT EXISTS "postgis";
CREATE extension IF NOT EXISTS "fuzzystrmatch";
CREATE extension IF NOT EXISTS "postgis_tiger_geocoder";
CREATE extension IF NOT EXISTS "postgis_topology";

ALTER SCHEMA tiger OWNER TO rds_superuser;
ALTER SCHEMA tiger_data OWNER TO rds_superuser;
ALTER SCHEMA topology OWNER TO rds_superuser;

CREATE FUNCTION exec(text) returns text language plpgsql volatile AS $f$ BEGIN EXECUTE $1; RETURN $1; END; $f$;

SELECT exec('ALTER TABLE ' || quote_ident(s.nspname) || '.' || quote_ident(s.relname) || ' OWNER TO rds_superuser;')
FROM (
    SELECT nspname, relname
    FROM pg_class c
    JOIN pg_namespace n ON (c.relnamespace = n.oid)
    WHERE nspname in ('tiger','topology')
    AND relkind IN ('r','S','v')
    ORDER BY relkind = 'S'
) s;

ALTER DATABASE postgres SET search_path TO public,tiger;
