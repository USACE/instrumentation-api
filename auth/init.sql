CREATE SCHEMA IF NOT EXISTS keycloak;

DO $$
    BEGIN
    CREATE ROLE keycloak_user LOGIN PASSWORD 'keycloak_password';
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating role keycloak_user -- it already exists';
END
$$;

ALTER SCHEMA keycloak OWNER TO keycloak_user;
