CREATE TABLE submittal_status_history (
    alert_config_id UUID REFERENCES alert_config (id) ON DELETE SET NULL,
    alert_status_id UUID NOT NULL REFERENCES alert_status (id),
    submitted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
