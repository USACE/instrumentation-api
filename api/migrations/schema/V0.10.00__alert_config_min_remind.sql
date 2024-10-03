ALTER TABLE alert_config
DROP CONSTRAINT interval_not_negative,
ADD  CONSTRAINT interval_not_negative CHECK (
	schedule_interval >= INTERVAL 'PT0S'
	AND warning_interval >= INTERVAL 'PT0S'
	-- reminder interval 0 if NA, >= 1 day for valid values
	AND (remind_interval = INTERVAL 'PT0S' OR remind_interval >= INTERVAL 'P1D')
);
