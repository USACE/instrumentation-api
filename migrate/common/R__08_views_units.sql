-- v_unit
DROP VIEW IF EXISTS v_unit;
CREATE VIEW v_unit AS (
    SELECT u.id AS id,
           u.name AS name,
           u.abbreviation AS abbreviation,
           u.unit_family_id AS unit_family_id,
           f.name           AS unit_family,
           u.measure_id     AS measure_id,
           m.name           AS measure
    FROM unit u
    INNER JOIN unit_family f ON f.id = u.unit_family_id
    INNER JOIN measure m ON m.id = u.measure_id
);

GRANT SELECT ON
    v_unit
TO instrumentation_reader;
