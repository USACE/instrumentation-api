DROP VIEW v_profile_project_roles;

CREATE OR REPLACE VIEW v_profile_project_roles AS (
    SELECT a.id,
           a.profile_id,
           b.edipi,
           b.username,
           b.email,
           b.is_admin,
           c.id AS project_id,
           r.id   AS role_id,
           r.name AS role,
           UPPER(c.slug || '.' || r.name) AS rolename
    FROM profile_project_roles a
    INNER JOIN profile b ON b.id = a.profile_id
    INNER JOIN project c ON c.id = a.project_id
    INNER JOIN role    r ON r.id = a.role_id
    ORDER BY username, role
);

GRANT SELECT ON v_profile_project_roles TO instrumentation_reader;