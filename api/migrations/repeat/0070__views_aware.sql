CREATE OR REPLACE VIEW v_aware_platform_parameter_enabled AS (
    SELECT
	i.id          AS instrument_id,
	a.aware_id    AS aware_id,
	b.key         AS aware_parameter_key,
	t.id          AS timeseries_id
    FROM aware_platform_parameter_enabled e
    INNER JOIN aware_platform a ON a.id = e.aware_platform_id
    INNER JOIN instrument i ON i.id = a.instrument_id
    INNER JOIN aware_parameter b ON b.id = e.aware_parameter_id
    LEFT JOIN timeseries t ON t.instrument_id=i.id AND t.parameter_id=b.parameter_id AND t.unit_id=b.unit_id
    ORDER BY a.aware_id
);
