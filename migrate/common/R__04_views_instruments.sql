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
    SELECT
        i.id,
        i.deleted,
        s.status_id,
        s.status,
        s.status_time,
        i.slug,
        i.name,
        i.type_id,
        t.name AS type,
        ST_AsBinary(i.geometry) AS geometry,
        i.station,
        i.station_offset,
        i.creator,
        i.create_date,
        i.updater,
        i.update_date,
        i.project_id,
        i.nid_id,
        i.usgs_id,
        tel.telemetry AS telemetry,
        COALESCE(c.constants, '{}') AS constants,
        COALESCE(g.groups, '{}') AS groups,
        COALESCE(a.alert_configs, '{}') AS alert_configs,
        COALESCE(o.opts, '{}'::JSON)::TEXT AS opts
    FROM instrument i
    INNER JOIN instrument_type t ON t.id = i.type_id
    INNER JOIN (
        SELECT
            DISTINCT ON (instrument_id) instrument_id,
            ss.time AS status_time,
            ss.status_id AS status_id,
            d.name AS status
        FROM instrument_status ss
        INNER JOIN status d ON d.id = ss.status_id
        WHERE ss.time <= NOW()
        ORDER BY instrument_id, ss.time DESC
    ) s ON s.instrument_id = i.id
    LEFT JOIN (
        SELECT
            ARRAY_AGG(timeseries_id) as constants,
            instrument_id
        FROM instrument_constants
        GROUP BY instrument_id
    ) c ON c.instrument_id = i.id
    LEFT JOIN (
        SELECT
            ARRAY_AGG(instrument_group_id) as groups,
            instrument_id
        FROM instrument_group_instruments
        GROUP BY instrument_id
    ) g ON g.instrument_id = i.id
    LEFT JOIN (
        SELECT
            ARRAY_AGG(alert_config_id) as alert_configs,
            instrument_id
        FROM alert_config_instrument
        GROUP BY instrument_id
    ) a ON a.instrument_id = i.id
    LEFT JOIN (
        SELECT
            instrument_id,
            JSON_AGG(JSON_BUILD_OBJECT(
                'id', v.id,
                'slug', v.telemetry_type_slug,
                'name', v.telemetry_type_name
            )) AS telemetry
        FROM v_instrument_telemetry v
        GROUP BY instrument_id
    ) tel ON tel.instrument_id = i.id
    LEFT JOIN (
        -- optional properties that vary per
        -- instrument can be added here via union
    SELECT o1.instrument_id, (ROW_TO_JSON(o1)::JSONB || ROW_TO_JSON(b1)::JSONB)::JSON AS opts
        FROM saa_opts o1
        LEFT JOIN LATERAL (
            SELECT value AS bottom_elevation FROM timeseries_measurement m
            WHERE m.timeseries_id = o1.bottom_elevation_timeseries_id
            ORDER BY m.time DESC
            LIMIT 1
        ) b1 ON true
        UNION ALL
    SELECT o2.instrument_id, (ROW_TO_JSON(o2)::JSONB || ROW_TO_JSON(b2)::JSONB)::JSON AS opts
        FROM ipi_opts o2
        LEFT JOIN LATERAL (
            SELECT value AS bottom_elevation FROM timeseries_measurement m
            WHERE m.timeseries_id = o2.bottom_elevation_timeseries_id
            ORDER BY m.time DESC
            LIMIT 1
        ) b2 ON true
    ) o ON o.instrument_id = i.id
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
                --COALESCE(count(tm.timeseries_id),0) as timeseries_measurements_count
                
        FROM instrument_group ig
        LEFT JOIN instrument_count ic on ic.instrument_group_id = ig.id
        LEFT JOIN timeseries_instruments ti on ig.id = ti.instrument_group_id
        --left join timeseries_measurement tm on tm.timeseries_id = ti.id
        GROUP BY ig.id, ic.i_count
        ORDER BY ig.name
);

GRANT SELECT ON
    v_instrument_telemetry,
    v_instrument,
    v_instrument_group
TO instrumentation_reader;
