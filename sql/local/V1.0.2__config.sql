-- Config for Local Development (minio replicating S3 Storage)
INSERT INTO config (static_host, static_prefix) VALUES
    ('http://localhost', '/instrumentation');
