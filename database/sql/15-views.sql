-- -----
-- Views
-- -----
-- v_telemetry

CREATE OR REPLACE VIEW v_instrument_telemetry AS (
    SELECT a.id,
           a.instrument_id AS instrument_id,
           b.id AS telemetry_type_id,
           b.slug AS telemetry_type_slug,
           b.name AS telemetry_type_name
    FROM instrument_telemetry a
    INNER JOIN telemetry_type b ON b.id = a.telemetry_type_id
    LEFT JOIN telemetry_goes tg ON a.telemetry_id = tg.id
    LEFT JOIN telemetry_iridium ti ON a.telemetry_id = ti.id
);

-- v_instrument
CREATE OR REPLACE VIEW v_instrument AS (
    SELECT I.id,
        I.deleted,
        S.status_id,
        S.status,
        S.status_time,
        I.slug,
        I.name,
        I.type_id,
        I.formula,
        T.name AS type,
        ST_AsBinary(I.geometry) AS geometry,
        I.station,
        I.station_offset,
        I.creator,
        I.create_date,
        I.updater,
        I.update_date,
        I.project_id,
        I.nid_id,
        I.usgs_id,
        TEL.telemetry AS telemetry,
        COALESCE(C.constants, '{}') AS constants,
        COALESCE(G.groups, '{}') AS groups,
        COALESCE(A.alert_configs, '{}') AS alert_configs
    FROM instrument I
    INNER JOIN instrument_type T ON T.id = I.type_id
    INNER JOIN (
        SELECT DISTINCT ON (instrument_id) instrument_id,
            a.time AS status_time,
            a.status_id AS status_id,
            d.name AS status
        FROM instrument_status a
        INNER JOIN status d ON d.id = a.status_id
        WHERE a.time <= now()
        ORDER BY instrument_id, a.time DESC
    ) S ON S.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(timeseries_id) as constants,
            instrument_id
        FROM instrument_constants
        GROUP BY instrument_id
    ) C on C.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(instrument_group_id) as groups,
            instrument_id
        FROM instrument_group_instruments
        GROUP BY instrument_id
    ) G on G.instrument_id = I.id
    LEFT JOIN (
        SELECT array_agg(id) as alert_configs,
            instrument_id
        FROM alert_config
        GROUP BY instrument_id
    ) A on A.instrument_id = I.id
    LEFT JOIN (
        SELECT instrument_id,
                json_agg(
                    json_build_object(
                        'id', v.id,
                        'slug', v.telemetry_type_slug,
                        'name', v.telemetry_type_name
                    )
                ) AS telemetry
        FROM v_instrument_telemetry v
        GROUP BY instrument_id
    ) TEL ON TEL.instrument_id = I.id
);

-- v_project
CREATE OR REPLACE VIEW v_project AS (
    SELECT  p.id,
            CASE WHEN p.image IS NOT NULL
                THEN cfg.static_host || '/projects/' || p.slug || '/images/' || p.image
                ELSE NULL
            END AS image,
            p.office_id,
            p.deleted,
            p.slug,
            p.name,
            p.creator,
            p.create_date,
            p.updater,
            p.update_date,
            COALESCE(t.timeseries, '{}') AS timeseries,
            COALESCE(i.count, 0) AS instrument_count,
            COALESCE(g.count, 0) AS instrument_group_count
        FROM project p
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument) as count
                FROM instrument
                WHERE NOT instrument.deleted
                GROUP BY project_id
            ) i ON i.project_id = p.id
            LEFT JOIN (
                SELECT project_id,
                    COUNT(instrument_group) as count
                FROM instrument_group
                WHERE NOT instrument_group.deleted
                GROUP BY project_id
            ) g ON g.project_id = p.id
            LEFT JOIN (
                SELECT array_agg(timeseries_id) as timeseries,
                    project_id
                FROM project_timeseries
                GROUP BY project_id
            ) t on t.project_id = p.id
			CROSS JOIN config cfg
);

-- v_timeseries
CREATE OR REPLACE VIEW v_timeseries AS (
        SELECT t.id AS id,
            t.slug AS slug,
            t.name AS name,
            i.slug || '.' || t.slug AS variable,
            j.id AS project_id,
            j.slug AS project_slug,
            j.name AS project,
            i.id AS instrument_id,
            i.slug AS instrument_slug,
            i.name AS instrument,
            p.id AS parameter_id,
            p.name AS parameter,
            u.id AS unit_id,
            u.name AS unit
        FROM timeseries t
            LEFT JOIN instrument i ON i.id = t.instrument_id
            LEFT JOIN project j ON j.id = i.project_id
            INNER JOIN parameter p ON p.id = t.parameter_id
            INNER JOIN unit U ON u.id = t.unit_id
    );

-- v_timeseries_project_map
CREATE OR REPLACE VIEW v_timeseries_project_map AS (
    SELECT t.id AS timeseries_id,
           p.id AS project_id
    FROM timeseries t
    LEFT JOIN instrument n ON t.instrument_id = n.id
    LEFT JOIN project p ON p.id = n.project_id
);

-- v_timeseries_latest; same as v_timeseries, joined with latest times and values
CREATE OR REPLACE VIEW v_timeseries_latest AS (
    SELECT t.*,
       m.time AS latest_time,
	   m.value AS latest_value
    FROM v_timeseries t
    LEFT JOIN (
	    SELECT DISTINCT ON (timeseries_id) timeseries_id, time, value
	    FROM timeseries_measurement
	    ORDER BY timeseries_id, time DESC
    ) m ON t.id = m.timeseries_id
);

-- v_email_autocomplete
CREATE OR REPLACE VIEW v_email_autocomplete AS (
    SELECT id,
           'email' AS user_type,
	       null AS username,
	       email AS email,
           email AS username_email
    FROM email
    UNION
    SELECT id,
           'profile' AS user_type,
           username,
           email,
           username||email AS username_email
    FROM profile
);

-- v_alert
CREATE OR REPLACE VIEW v_alert AS (
    SELECT a.id AS id,
       a.alert_config_id AS alert_config_id,
       a.create_date AS create_date,
       p.id AS project_id,
       p.name AS project_name,
	   i.id AS instrument_id,
	   i.name AS instrument_name,
	   ac.name AS name,
	   ac.body AS body
FROM alert a
INNER JOIN alert_config ac ON a.alert_config_id = ac.id
INNER JOIN instrument i ON ac.instrument_id = i.id
INNER JOIN project p ON i.project_id = p.id
);

CREATE OR REPLACE VIEW v_unit AS (
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

CREATE OR REPLACE VIEW v_plot_configuration AS (
    SELECT pc.id            AS id,
           pc.slug          AS slug,
           pc.name          AS name,
           pc.project_id    AS project_id,
           t.timeseries_id     AS timeseries_id,
           pc.creator       AS creator,
           pc.create_date   AS create_date,
           pc.updater       AS updater,
           pc.update_date   AS update_date
    FROM plot_configuration pc
    LEFT JOIN (
        SELECT plot_configuration_id    as plot_configuration_id,
               array_agg(timeseries_id) as timeseries_id
        FROM plot_configuration_timeseries
        GROUP BY plot_configuration_id
    ) as t ON pc.id = t.plot_configuration_id
);

-- v_instrument_groups
CREATE OR REPLACE VIEW v_instrument_group AS (
    WITH instrument_count AS (
        SELECT 
        igi.instrument_group_id,
        count(igi.instrument_group_id) as i_count 
        FROM instrument_group_instruments igi
        JOIN instrument i on igi.instrument_id = i.id and not i.deleted
        GROUP BY igi.instrument_group_id
        )
        ,
        timeseries_instruments as (
            SELECT t.id, t.instrument_id, igi.instrument_group_id from timeseries t 
            JOIN instrument i on i.id = t.instrument_id and not i.deleted
            JOIN instrument_group_instruments igi on igi.instrument_id = i.id
        )

        SELECT  ig.id,
                ig.slug,
                ig.name,
                ig.description,
                ig.creator,
                ig.create_date,
                ig.updater,
                ig.update_date,
                ig.project_id,
                ig.deleted,
                COALESCE(ic.i_count,0) as instrument_count,
                COALESCE(count(ti.id),0) as timeseries_count
                --,
                --COALESCE(count(tm.id),0) as timeseries_measurements_count
                
        FROM instrument_group ig
        LEFT JOIN instrument_count ic on ic.instrument_group_id = ig.id
        LEFT JOIN timeseries_instruments ti on ig.id = ti.instrument_group_id
        --left join timeseries_measurement tm on tm.timeseries_id = ti.id
        GROUP BY ig.id, ic.i_count
        ORDER BY ig.name
);

-- v_profile_project_roles
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