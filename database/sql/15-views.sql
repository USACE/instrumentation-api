-- -----
-- Views
-- -----

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
                ORDER BY instrument_id,
                    a.time DESC
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
    );

-- v_project
CREATE OR REPLACE VIEW v_project AS (
    SELECT  p.id,
            CASE WHEN p.image IS NOT NULL
                THEN cfg.static_host || cfg.static_prefix || '/projects/' || p.slug || '/images/' || p.image
                ELSE NULL
            END AS image,
            p.office_id,
            p.deleted,
            p.slug,
            p.federal_id,
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

-- v_telemetry
CREATE OR REPLACE VIEW v_instrument_telemetry AS (
    SELECT a.id,
           b.id AS telemetry_type_id,
           b.slug AS telemetry_type_slug,
           b.name AS telemetry_type_name
    FROM instrument_telemetry a
    INNER JOIN telemetry_type b ON b.id = a.telemetry_type_id
    LEFT JOIN telemetry_goes tg ON a.telemetry_id = tg.id
    LEFT JOIN telemetry_iridium ti ON a.telemetry_id = ti.id
);
