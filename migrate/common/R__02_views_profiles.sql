-- ${flyway:timestamp}
CREATE VIEW v_profile AS (
    WITH roles_by_profile AS (
        SELECT
            profile_id,
            array_agg(upper(b.slug || '.' || c.name)) AS roles
        FROM profile_project_roles a
        LEFT JOIN project b ON a.project_id = b.id
        LEFT JOIN role c ON a.role_id = c.id
        GROUP BY profile_id
    )
    SELECT
        p.id,
        p.edipi,
        p.username,
        p.display_name,
        p.email,
        p.is_admin,
        COALESCE(r.roles,'{}') AS roles
    FROM profile p
    LEFT JOIN roles_by_profile r ON r.profile_id = p.id
);

CREATE VIEW v_profile_project_roles AS (
    SELECT
        a.id,
        a.profile_id,
        b.edipi,
        b.username,
        b.display_name,
        b.email,
        b.is_admin,
        c.id AS project_id,
        r.id AS role_id,
        r.name AS role,
        upper(c.slug || '.' || r.name) AS rolename
    FROM profile_project_roles a
    INNER JOIN profile b ON b.id = a.profile_id
    INNER JOIN project c ON c.id = a.project_id
    INNER JOIN role r ON r.id = a.role_id
    ORDER BY username, role
);

CREATE VIEW v_email_autocomplete AS (
    SELECT
        id,
        'email' AS user_type,
        null AS username,
        email AS email,
        email AS username_email
    FROM email
    UNION
    SELECT
        id,
        'profile' AS user_type,
        username,
        email,
        username || email AS username_email
    FROM profile
);

GRANT SELECT ON
    v_profile,
    v_profile_project_roles,
    v_email_autocomplete
TO instrumentation_reader;
