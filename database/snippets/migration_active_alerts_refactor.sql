-- drop affected views
drop view v_instrument;

-- drop affected tables
drop table if exists
    public.alert,
    public.email_alerts,
    public.profile_alerts;

-- alert_config
CREATE TABLE IF NOT EXISTS public.alert_config (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    name VARCHAR(480),
    body TEXT,
    formula TEXT,
    schedule TEXT,
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT instrument_unique_alert_config_name UNIQUE(name,instrument_id)
);

-- alert
CREATE TABLE IF NOT EXISTS public.alert (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    create_date TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- profile alert subscription (subscribe profiles to alerts)
CREATE TABLE IF NOT EXISTS public.alert_profile_subscription (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    profile_id UUID NOT NULL REFERENCES profile (id),
    mute_ui boolean NOT NULL DEFAULT false,
    mute_notify boolean NOT NULL DEFAULT false,
    CONSTRAINT profile_unique_alert_config UNIQUE(profile_id, alert_config_id)
);

-- email alert subscription (subscribe emails to alerts)
CREATE TABLE IF NOT EXISTS public.alert_email_subscription (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    email_id UUID NOT NULL REFERENCES email (id),
    mute_notify boolean NOT NULL DEFAULT false,
    CONSTRAINT email_unique_alert_config UNIQUE(email_id, alert_config_id)
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

-- Grants
GRANT SELECT ON
    alert,
    alert_config,
    alert_email_subscription,
    alert_profile_subscription,
    v_instrument
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    alert,
    alert_config,
    alert_email_subscription,
    alert_profile_subscription
TO instrumentation_writer;