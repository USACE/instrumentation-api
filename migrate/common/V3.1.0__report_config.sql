CREATE TABLE IF NOT EXISTS report_config (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  project_id uuid NOT NULL REFERENCES project (id),
  slug text UNIQUE NOT NULL,
  name text NOT NULL,
  description text NOT NULL,
  creator uuid NOT NULL REFERENCES profile (id),
  create_date timestamptz NOT NULL DEFAULT now(),
  updater uuid REFERENCES profile (id),
  update_date timestamptz,
  date_range text DEFAULT '1 year',
  date_range_enabled boolean DEFAULT false,
  show_masked boolean DEFAULT false,
  show_masked_enabled boolean DEFAULT false,
  show_nonvalidated boolean DEFAULT false,
  show_nonvalidated_enabled boolean DEFAULT false
);

CREATE TABLE IF NOT EXISTS report_config_plot_config (
  report_config_id uuid NOT NULL REFERENCES report_config (id) ON DELETE CASCADE,
  plot_config_id uuid NOT NULL REFERENCES plot_configuration (id) ON DELETE CASCADE,
  CONSTRAINT report_config_plot_config_report_config_id_plot_config_id_key UNIQUE(report_config_id,plot_config_id)
);

CREATE TABLE IF NOT EXISTS job_status (
  id int PRIMARY KEY NOT NULL,
  status text NOT NULL
);

INSERT INTO job_status (id, status) VALUES (0, 'SUCCESS'), (1, 'FAIL'), (2, 'INIT');

CREATE TABLE IF NOT EXISTS report_download_job (
  job_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  report_config_id uuid REFERENCES report_config (id),
  create_date timestamptz NOT NULL DEFAULT now(),
  update_date timestamptz,
  status int REFERENCES job_status (id) DEFAULT 2,
  file_key text
);
