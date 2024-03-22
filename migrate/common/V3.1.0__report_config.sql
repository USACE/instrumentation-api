CREATE TABLE IF NOT EXISTS report_config (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  project_id uuid NOT NULL REFERENCES project (id),
  slug text UNIQUE NOT NULL,
  name text NOT NULL,
  description text NOT NULL,
  after timestamptz,
  before timestamptz,
  creator uuid NOT NULL REFERENCES profile (id),
  create_date timestamptz NOT NULL DEFAULT now(),
  updater uuid REFERENCES profile (id),
  update_date timestamptz
);

CREATE TABLE IF NOT EXISTS report_config_plot_config (
  report_config_id uuid NOT NULL REFERENCES report_config (id) ON DELETE CASCADE,
  plot_config_id uuid NOT NULL REFERENCES plot_configuration (id) ON DELETE CASCADE,
  CONSTRAINT report_config_plot_config_report_config_id_plot_config_id_key UNIQUE(report_config_id,plot_config_id)
);
