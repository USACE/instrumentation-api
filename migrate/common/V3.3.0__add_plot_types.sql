CREATE TYPE plot_type AS ENUM ('scatter-line', 'profile', 'contour', 'bullseye');

ALTER TABLE plot_configuration ADD COLUMN plot_type plot_type;

UPDATE plot_configuration SET plot_type = 'scatter-line';

ALTER TABLE plot_configuration
ALTER COLUMN plot_type SET NOT NULL;

CREATE TABLE plot_contour_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
  time timestamptz NOT NULL,
  locf_backfill interval NOT NULL,
  gradient_smoothing boolean NOT NULL DEFAULT false,
  contour_smoothing boolean NOT NULL DEFAULT false,
  show_labels boolean NOT NULL DEFAULT false
);

CREATE TABLE plot_contour_config_timeseries (
  plot_contour_config_id uuid NOT NULL REFERENCES plot_contour_config(plot_config_id) ON DELETE CASCADE,
  timeseries_id uuid NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
  CONSTRAINT plot_contour_config_timeseries_plot_contour_config_id_timeseries_id_key
  UNIQUE(plot_contour_config_id, timeseries_id)
);

CREATE TABLE plot_bullseye_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
  x_axis_timeseries_id uuid REFERENCES timeseries(id) ON DELETE SET NULL,
  y_axis_timeseries_id uuid REFERENCES timeseries(id) ON DELETE SET NULL
);

CREATE TABLE plot_profile_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
  instrument_id uuid NOT NULL REFERENCES instrument(id) ON DELETE CASCADE
);

CREATE TABLE plot_scatter_line_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
  y_axis_title text,
  y2_axis_title text
);

INSERT INTO plot_scatter_line_config (plot_config_id, y_axis_title, y2_axis_title)
SELECT id, yaxis_title, secondary_axis_title FROM plot_configuration_settings;

ALTER TABLE plot_configuration_settings
DROP COLUMN yaxis_title,
DROP COLUMN secondary_axis_title;
