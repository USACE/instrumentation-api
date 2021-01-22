-- NOTE: Creation of instrumentation_reader and instrumentation_writer roles
--       happens in 20-roles.sql, typically run prior to running this file

-- Role instrumentation_reader
-- Tables specific to instrumentation app
GRANT SELECT ON
    aware_platform_parameter_enabled,
    aware_platform,
    aware_parameter
TO instrumentation_reader;

-- Role instrumentation_writer
-- Tables specific to instrumentation app
GRANT INSERT,UPDATE,DELETE ON
    aware_platform_parameter_enabled,
    aware_platform,
    aware_parameter
TO instrumentation_writer;
