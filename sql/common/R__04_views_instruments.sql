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
