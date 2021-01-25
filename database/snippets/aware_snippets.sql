-- Turn On Data Acquisition for All AWARE Platforms; All AWARE Parameters
INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b
	ORDER BY aware_platform_id
);