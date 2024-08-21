ALTER TABLE alert_config
  DROP CONSTRAINT warning_before_schedule,
  DROP CONSTRAINT interval_not_negative;

CREATE TABLE alert_config_scheduler (
  alert_config_id uuid NOT NULL REFERENCES alert_config(id) ON DELETE CASCADE,
  start_date timestamptz NOT NULL DEFAULT now(),
  schedule_interval interval NOT NULL,
  warning_interval interval NOT NULL DEFAULT 'PT0',
  remind_interval interval NOT NULL DEFAULT 'PT0',
  last_reminded timestamptz,
  mute_consecutive_alerts boolean NOT NULL DEFAULT false
  CONSTRAINT warning_before_schedule CHECK (warning_interval < schedule_interval),
  CONSTRAINT interval_not_negative CHECK (
    schedule_interval >= INTERVAL 'PT0'
    AND warning_interval >= INTERVAL 'PT0'
    AND remind_interval >= INTERVAL 'PT0'
  )
);

INSERT INTO alert_config_scheduler (alert_config_id, start_date, schedule_interval, warning_interval, remind_interval, last_reminded, mute_consecutive_alerts)
SELECT id, start_date, schedule_interval, warning_interval, remind_interval, last_reminded, mute_consecutive_alerts
FROM alert_config;

DROP VIEW IF EXISTS v_alert_config;
DROP VIEW IF EXISTS v_alert_check_measurement_submittal;
DROP VIEW IF EXISTS v_alert_check_evaluation_submittal;

ALTER TABLE alert_config
  -- dropping because migrating to alert_config_scheduler table
  DROP start_date,
  DROP schedule_interval,
  DROP warning_interval,
  DROP remind_interval,
  DROP last_reminded,
  DROP mute_consecutive_alerts,
  -- dropping because ununsed
  DROP n_missed_before_alert,
  ADD COLUMN muted boolean DEFAULT false;

UPDATE alert_config SET muted = false;
ALTER TABLE alert_config ALTER COLUMN muted SET NOT NULL;

CREATE TABLE alert_config_timeseries (
  alert_config_id uuid NOT NULL REFERENCES alert_config(id) ON DELETE CASCADE,
  timeseries_id uuid NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE
);

INSERT INTO alert_type VALUES
('bb15e7c2-8eae-452c-92f7-e720dc5c9432', 'Threshold'),
('c37effee-6b48-4436-8d72-737ed78c1fb7', 'Rate of Change');

CREATE TABLE alert_config_threshold (
  alert_config_id uuid NOT NULL REFERENCES alert_config(id) ON DELETE CASCADE,
  alert_low_value double precision,
  alert_high_value double precision,
  warn_low_value double precision,
  warn_high_value double precision,
  ignore_low_value double precision,
  ignore_high_value double precision,
  variance double precision NOT NULL DEFAULT 0,
  CHECK (ignore_low_value < alert_low_value),
  CHECK (alert_low_value < warn_low_value),
  CHECK (warn_low_value < warn_high_value),
  CHECK (warn_high_value < alert_high_value),
  CHECK (alert_high_value < ignore_high_value)
);

CREATE TABLE alert_config_change (
  alert_config_id uuid NOT NULL REFERENCES alert_config(id) ON DELETE CASCADE,
  warn_rate_of_change double precision,
  alert_rate_of_change double precision NOT NULL,
  ignore_rate_of_change double precision,
  locf_backfill interval
);
