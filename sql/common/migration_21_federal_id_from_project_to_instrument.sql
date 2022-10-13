
DROP VIEW v_instrument;
DROP VIEW v_project;

ALTER TABLE INSTRUMENT ADD COLUMN nid_id VARCHAR;
ALTER TABLE PROJECT DROP COLUMN federal_id;

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
                THEN cfg.static_host || cfg.static_prefix || '/projects/' || p.slug || '/images/' || p.image
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

GRANT SELECT ON
    v_instrument,
    v_instrument_telemetry,
    v_project
TO instrumentation_reader;
