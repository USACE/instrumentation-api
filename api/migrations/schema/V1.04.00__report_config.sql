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

CREATE TYPE job_status AS ENUM ('SUCCESS', 'FAIL', 'INIT');

CREATE TABLE IF NOT EXISTS report_download_job (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  report_config_id uuid REFERENCES report_config (id) ON DELETE CASCADE,
  creator uuid NOT NULL REFERENCES profile (id) ON DELETE CASCADE,
  create_date timestamptz NOT NULL DEFAULT now(),
  status job_status NOT NULL DEFAULT 'INIT',
  file_key text,
  file_expiry timestamptz,
  progress int NOT NULL DEFAULT 0 CHECK (progress >= 0 AND progress <= 100),
  progress_update_date timestamptz NOT NULL DEFAULT now()
);
