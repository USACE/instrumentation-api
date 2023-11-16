DROP TABLE IF EXISTS saa_opts;
CREATE TABLE IF NOT EXISTS saa_opts (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    num_segments INT NOT NULL,
    bottom_elevation_timeseries_id UUID REFERENCES timeseries (id),
    initial_time TIMESTAMPTZ
);

DROP TABLE IF EXISTS saa_segment;
CREATE TABLE IF NOT EXISTS saa_segment ( 
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    id INT NOT NULL,
    length_timeseries_id UUID REFERENCES timeseries (id),
    x_timeseries_id UUID REFERENCES timeseries (id),
    y_timeseries_id UUID REFERENCES timeseries (id),
    z_timeseries_id UUID REFERENCES timeseries (id),
    temp_timeseries_id UUID REFERENCES timeseries (id),
    PRIMARY KEY (instrument_id, id)
);

INSERT INTO parameter (id, name) VALUES ('6d12ca4c-b618-41cd-87a2-a248980a0d69', 'saa-constant');

CREATE TABLE IF NOT EXISTS ipi_opts (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    num_segments INT NOT NULL,
    bottom_elevation_timeseries_id UUID REFERENCES timeseries (id),
    initial_time TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS ipi_segment (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    id INT NOT NULL,
    length_timeseries_id UUID REFERENCES timeseries (id),
    tilt_timeseries_id UUID REFERENCES timeseries (id),
    cum_dev_timeseries_id UUID REFERENCES timeseries (id),
    PRIMARY KEY (instrument_id, id)
);

INSERT INTO instrument_type (id, name) VALUES ('c81f3a5d-fc5f-47fd-b545-401fe6ee63bb', 'IPI');

INSERT INTO parameter (id, name) VALUES ('a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'ipi-constant');
