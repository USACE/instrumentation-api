CREATE TYPE plot_type AS ENUM ('scatter-line', 'profile', 'contour', 'bullseye');

ALTER TABLE plot_configuration_settings ADD COLUMN plot_type plot_type;

UPDATE plot_configuration_settings SET plot_type = 'scatter-line';

ALTER TABLE plot_configuration_settings
ALTER COLUMN plot_type SET NOT NULL;

CREATE TABLE plot_contour_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_config(id),
  date timestamptz NOT NULL,
  locf_backfill interval NOT NULL,
  gradient_smoothing boolean NOT NULL DEFAULT false,
  contour_smoothing boolean NOT NULL DEFAULT false,
  show_labels boolean NOT NULL DEFAULT false
);

CREATE TABLE plot_contour_config_timeseries (
  plot_contour_config_id uuid NOT NULL REFERENCES plot_contour_config(plot_config_id),
  timeseries_id uuid NOT NULL REFERENCES timeseries(id),
  CONSTRAINT UNIQUE(plot_contour_config_id, timeseries_id)
);

CREATE TABLE plot_bullseye_config (
  plot_config_id uuid UNIQUE NOT NULL REFERENCES plot_config(id),
  x_axis_timeseries_id uuid NOT NULL REFERENCES timeseries(id),
  y_axis_timeseries_id uuid NOT NULL REFERENCES timeseries(id),
  CONSTRAINT UNIQUE(x_axis_timeseries_id, y_axis_timeseries_id)
);
