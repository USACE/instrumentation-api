CREATE TABLE survey123 (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
  project_id uuid NOT NULL REFERENCES project(id),
  name text NOT NULL,
  deleted boolean NOT NULL DEFAULT false
);

CREATE TABLE survey123_equivalency_table (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    survey123_deleted boolean NOT NULL DEFAULT false,
    field_name text NOT NULL,
    display_name text,
    instrument_id uuid REFERENCES instrument (id) ON DELETE CASCADE,
    timeseries_id uuid REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT unique_survey123_field UNIQUE(survey123_id, field_name),
    CONSTRAINT unique_active_survey123 FOREIGN KEY (survey123_id, survey123_deleted)
        REFERENCES survey123(id, deleted) ON UPDATE CASCADE
);
