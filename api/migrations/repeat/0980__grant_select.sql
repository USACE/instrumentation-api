GRANT SELECT ON ALL TABLES IN SCHEMA midas TO instrumentation_reader;
REVOKE SELECT ON schema_migration_history FROM instrumentation_reader;
