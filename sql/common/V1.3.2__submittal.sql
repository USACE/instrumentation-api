CREATE EXTENSION btree_gist;

ALTER TABLE alert_config DROP COLUMN alert_status_id;
ALTER TABLE alert_status RENAME TO submittal_status;

CREATE TABLE submittal (
    id                      UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id         UUID REFERENCES alert_config (id) ON DELETE CASCADE,
    submittal_status_id     UUID REFERENCES submittal_status (id) DEFAULT '0c0d6487-3f71-4121-8575-19514c7b9f03'::UUID,
    completion_date         TIMESTAMPTZ,
    create_date             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    due_date                TIMESTAMPTZ NOT NULL,
    marked_as_missing       BOOLEAN NOT NULL DEFAULT false,
    warning_sent            BOOLEAN NOT NULL DEFAULT false,
    CHECK (create_date < due_date),
    EXCLUDE USING gist (alert_config_id WITH =, TSTZRANGE(create_date, due_date) WITH &&)
);

CREATE UNIQUE INDEX unique_alert_config_id_submittal_date ON submittal (alert_config_id,completion_date)
WHERE completion_date IS NOT NULL;

DROP VIEW IF EXISTS v_evaluation;
ALTER TABLE evaluation DROP COLUMN alert_config_id;
ALTER TABLE evaluation ADD COLUMN submittal_id UUID REFERENCES submittal (id) ON DELETE CASCADE;
