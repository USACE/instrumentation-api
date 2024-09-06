-- +goose up
GRANT SELECT ON ALL TABLES IN SCHEMA midas TO instrumentation_reader;
