drop table if exists 
    public.profile,
    public.email,
    public.alert,
    public.profile_alerts,
    public.email_alerts
	CASCADE;

CREATE TABLE IF NOT EXISTS public.profile (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    edipi VARCHAR(240) UNIQUE NOT NULL,
    username VARCHAR(240) UNIQUE NOT NULL,
    email VARCHAR(240) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS public.email (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(240) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS public.alert (
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
    CONSTRAINT instrument_unique_alert_name UNIQUE(name,instrument_id)
);

-- profile alerts (subscribe profiles to alerts)
CREATE TABLE IF NOT EXISTS public.profile_alerts (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_id UUID NOT NULL REFERENCES alert (id),
    profile_id UUID NOT NULL REFERENCES profile (id),
    mute_ui boolean NOT NULL DEFAULT false,
    mute_notify boolean NOT NULL DEFAULT false
);

-- email alerts (subscribe emails to alerts)
CREATE TABLE IF NOT EXISTS public.email_alerts (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_id UUID NOT NULL REFERENCES alert (id),
    email_id UUID NOT NULL REFERENCES email (id),
    mute_notify boolean NOT NULL DEFAULT false
);

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

-- Sample Data
INSERT INTO alert (id, instrument_id, name, body, formula, schedule) VALUES
    ('1efd2d85-d3ee-4388-85a0-f824a761ff8b', '9e8f2ca4-4037-45a4-aaca-d9e598877439','Above Target Height', 'The demo staff gage has exceeded the target height. Sincerely, Midas', '[stage] >= 10', '0,10,20,30,40,50 * * * *'),
    ('243e9d32-2cba-4f12-9abe-63adc09fc5dd', 'a7540f69-c41e-43b3-b655-6e44097edb7e','Below Target Height', 'Distance to water is near artesian conditions. Sincerely, Midas', '[distance-to-water] <= 2', '0,10,20,30,40,50 * * * *');

-- New Grants
GRANT SELECT ON
    alert,
    profile,
    profile_alerts,
    email,
    email_alerts,
    v_email_autocomplete
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    alert,
    profile,
    profile_alerts,
    email,
    email_alerts,
    v_email_autocomplete
TO instrumentation_writer;
