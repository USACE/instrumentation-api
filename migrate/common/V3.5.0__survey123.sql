CREATE TABLE survey123 (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    project_id uuid NOT NULL REFERENCES project(id),
    name text UNIQUE NOT NULL,
    slug text UNIQUE NOT NULL,
    create_date NOT NULL DEFAULT now(),
    update_date NOT NULL DEFAULT now(),
    creator uuid NOT NULL REFERENCES profile(id),
    updater uuid REFERENCES profile(id),
    deleted boolean NOT NULL DEFAULT false
);

CREATE TABLE survey123_equivalency_table (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    survey123_deleted boolean NOT NULL DEFAULT false,
    field_name text NOT NULL,
    display_name text,
    instrument_id uuid REFERENCES instrument(id) ON DELETE CASCADE,
    timeseries_id uuid REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT  UNIQUE(),
    CONSTRAINT unique_active_survey123 FOREIGN KEY (survey123_id, survey123_deleted)
        REFERENCES survey123(id, deleted) ON UPDATE CASCADE
);

CREATE UNIQUE INDEX survey123_equivalency_table_survey123_id_field_name_key ON survey123_equivalency_table(survey123_id, field_name)
WHERE NOT survey123_deleted;

CREATE TABLE survey123_preview (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    preview json NOT NULL,
    update_date timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE survey123_payload_error (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    error_message text
);
