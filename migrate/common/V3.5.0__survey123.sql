CREATE TABLE survey123 (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    project_id uuid NOT NULL REFERENCES project(id),
    name text UNIQUE NOT NULL,
    slug text UNIQUE NOT NULL,
    create_date timestamptz NOT NULL DEFAULT now(),
    update_date timestamptz,
    creator uuid NOT NULL REFERENCES profile(id),
    updater uuid REFERENCES profile(id),
    deleted boolean NOT NULL DEFAULT false,
    CONSTRAINT survey123_id_deleted_key UNIQUE (id, deleted)
);

CREATE TABLE survey123_equivalency_table (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    survey123_deleted boolean NOT NULL DEFAULT false,
    field_name text NOT NULL,
    display_name text,
    instrument_id uuid REFERENCES instrument(id) ON DELETE CASCADE,
    timeseries_id uuid REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT survey123_equivalency_table_survey123_id_survey123_deleted_field_name_key UNIQUE (survey123_id, survey123_deleted, field_name),
    CONSTRAINT unique_active_survey123 FOREIGN KEY (survey123_id, survey123_deleted)
        REFERENCES survey123(id, deleted) ON UPDATE CASCADE
);

CREATE TABLE survey123_preview (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    preview json NOT NULL,
    update_date timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE survey123_payload_error (
    survey123_id uuid NOT NULL REFERENCES survey123(id),
    error_message text
);
