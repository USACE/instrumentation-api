DROP TABLE public.profile_alerts;

-- profile alerts (subscribe profiles to alerts)
CREATE TABLE IF NOT EXISTS public.profile_alerts (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_id UUID NOT NULL REFERENCES alert (id),
    profile_id UUID NOT NULL REFERENCES profile (id),
    mute_ui boolean NOT NULL DEFAULT false,
    mute_notify boolean NOT NULL DEFAULT false,
    CONSTRAINT profile_unique_alert UNIQUE(profile_id, alert_id)
);

GRANT SELECT ON profile_alerts TO instrumentation_reader;
GRANT INSERT,UPDATE,DELETE ON profile_alerts TO instrumentation_writer;