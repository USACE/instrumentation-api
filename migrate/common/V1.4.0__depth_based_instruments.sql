CREATE TABLE saa_opts (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    num_segments INT NOT NULL,
    bottom_elevation REAL NOT NULL,
    initial_time TIMESTAMPTZ
);

CREATE TABLE saa_segment ( 
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    id INT NOT NULL,
    length REAL,
    x_timeseries_id UUID REFERENCES timeseries (id),
    y_timeseries_id UUID REFERENCES timeseries (id),
    z_timeseries_id UUID REFERENCES timeseries (id),
    temp_timeseries_id UUID REFERENCES timeseries (id),
    PRIMARY KEY (instrument_id, id)
);

INSERT INTO instrument_type (id, name) VALUES ('07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9', 'SAA');

CREATE TABLE ipi_opts (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    num_segments INT NOT NULL,
    bottom_elevation REAL NOT NULL,
    initial_time TIMESTAMPTZ
);

CREATE TABLE ipi_segment (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    id INT NOT NULL,
    length REAL,
    tilt_timeseries_id UUID REFERENCES timeseries (id),
    cum_dev_timeseries_id UUID REFERENCES timeseries (id),
    PRIMARY KEY (instrument_id, id)
);

INSERT INTO instrument_type (id, name) VALUES ('c81f3a5d-fc5f-47fd-b545-401fe6ee63bb', 'IPI');
