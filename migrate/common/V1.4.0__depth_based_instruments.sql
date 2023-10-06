CREATE TABLE saa_instrument (
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    num_segments INT NOT NULL,
    bottom_elevation REAL NOT NULL,
    initial_time TIMESTAMPTZ
);

CREATE TABLE saa_segment ( 
    instrument_id UUID NOT NULL REFERENCES instrument (id) ON DELETE CASCADE,
    id INT NOT NULL,
    length REAL NOT NULL,
    x_timeseries_id UUID NOT NULL REFERENCES timeseries (id),
    y_timeseries_id UUID NOT NULL REFERENCES timeseries (id),
    z_timeseries_id UUID NOT NULL REFERENCES timeseries (id),
    temp_timeseries_id UUID NOT NULL REFERENCES timeseries (id),
    PRIMARY KEY (instrument_id, id)
);

INSERT INTO instrument_type (id, name) VALUES ('07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9', 'SAA');
