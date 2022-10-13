-- role
CREATE TABLE IF NOT EXISTS public.role (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    deleted boolean NOT NULL DEFAULT false
);

INSERT INTO role (id, name) VALUES
    ('37f14863-8f3b-44ca-8deb-4b74ce8a8a69', 'ADMIN'),
    ('2962bdde-7007-4ba0-943f-cb8e72e90704', 'MEMBER');

CREATE TABLE IF NOT EXISTS public.profile_project_roles (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    profile_id UUID NOT NULL REFERENCES profile(id),
    role_id UUID NOT NULL REFERENCES role(id),
    project_id UUID NOT NULL REFERENCES project(id),
    granted_by UUID REFERENCES profile(id),
    granted_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_profile_project_role UNIQUE(profile_id,project_id,role_id)
);

ALTER TABLE public.profile ADD COLUMN is_admin boolean NOT NULL DEFAULT false;

CREATE OR REPLACE VIEW v_profile_project_roles AS (
    SELECT a.id,
           a.profile_id,
           b.edipi,
           b.username,
           b.email,
           b.is_admin,
           a.id AS project_id,
           UPPER(c.slug || '.' || r.name) AS role
    FROM profile_project_roles a
    INNER JOIN profile b ON b.id = a.profile_id
    INNER JOIN project c ON c.id = a.project_id
    INNER JOIN role    r ON r.id = a.role_id
    ORDER BY username, role
);

CREATE OR REPLACE VIEW v_profile AS (
    WITH roles_by_profile AS (
        SELECT profile_id,
               array_agg(UPPER(b.slug || '.' || c.name)) AS roles
        FROM profile_project_roles a
        LEFT JOIN project b ON a.project_id = b.id
        LEFT JOIN role    c ON a.role_id    = c.id
        GROUP BY profile_id
    )
    SELECT p.id,
           p.edipi,
           p.username,
           p.email,
           p.is_admin,
           COALESCE(r.roles,'{}') AS roles
    FROM profile p
    LEFT JOIN roles_by_profile r ON r.profile_id = p.id
);

GRANT SELECT ON
    profile,
    role,
    profile_project_roles,
    v_profile,
    v_profile_project_roles
TO instrumentation_reader;

GRANT INSERT,UPDATE,DELETE ON
    profile,
    profile_project_roles,
    role
TO instrumentation_writer;