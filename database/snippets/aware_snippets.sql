
-- NOTE: UUID in this file are random, auto generated as an example only

-- Turn On Data Acquisition for All AWARE Platforms; All AWARE Parameters
INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b
	ORDER BY aware_platform_id
);

-- Turn off all Aware Data Acquisition for All Parameters, All Platforms
DELETE FROM aware_platform_parameter_enabled;

-- Confirm all aware parameters turned off for all gages
SELECT * FROM aware_platform_parameter_enabled;

-- Aware Platform by Aware ID (Used in FlashFloodInfo)
SELECT * FROM aware_platform WHERE aware_id = '4112fd82-d910-46ee-ae08-720742d028c2';

-- Timeseries for Instrument Associated with Aware Gage
SELECT * FROM timeseries WHERE instrument_id IN (
	SELECT instrument_id from aware_platform WHERE aware_id = '4112fd82-d910-46ee-ae08-720742d028c2'
);
SELECT * FROM TIMESERIES

-- Aware Parameters in system
SELECT * FROM aware_parameter;

-- Turn on One Parameter Data Acquisition for Aware Gage
INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b
	WHERE a.aware_id = '4112fd82-d910-46ee-ae08-720742d028c2'
	ORDER BY aware_platform_id
);
