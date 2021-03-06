DROP VIEW v_instrument;
DROP VIEW v_alert;
DROP VIEW v_timeseries_latest;
DROP VIEW v_timeseries;

ALTER TABLE INSTRUMENT ALTER COLUMN NAME TYPE VARCHAR(360);
ALTER TABLE INSTRUMENT ALTER COLUMN SLUG TYPE VARCHAR;
ALTER TABLE INSTRUMENT DROP CONSTRAINT project_unique_instrument_name;

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

GRANT SELECT ON
	v_instrument, 
	v_alert,
	v_timeseries,
	v_timeseries_latest
TO instrumentation_reader;