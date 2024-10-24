ALTER TABLE plot_configuration_settings
ADD COLUMN secondary_axis_title text;

CREATE TYPE line_style AS ENUM ('solid', 'dot', 'dash', 'longdash', 'dashdot', 'longdashdot');
CREATE TYPE y_axis AS ENUM ('y1', 'y2');
CREATE TYPE trace_type AS ENUM ('bar', 'scattergl');

CREATE TABLE plot_configuration_timeseries_trace (
  plot_configuration_id uuid REFERENCES plot_configuration (id) ON DELETE CASCADE,
  timeseries_id uuid REFERENCES timeseries (id) ON DELETE CASCADE,
  trace_order int NOT NULL,
  trace_type trace_type NOT NULL DEFAULT 'scattergl',
  color text NOT NULL,
  line_style line_style NOT NULL DEFAULT 'solid',
  width real NOT NULL DEFAULT 1,
  show_markers boolean NOT NULL DEFAULT false,
  y_axis y_axis NOT NULL DEFAULT 'y1',
  UNIQUE (plot_configuration_id, timeseries_id)
);

CREATE TABLE plot_configuration_custom_shape (
  plot_configuration_id uuid REFERENCES plot_configuration (id) ON DELETE CASCADE,
  enabled boolean NOT NULL DEFAULT false,
  name text NOT NULL,
  data_point real NOT NULL,
  color text NOT NULL
);

INSERT INTO plot_configuration_timeseries_trace (plot_configuration_id, timeseries_id, trace_order, color)
SELECT plot_configuration_id, timeseries_id, 0, '#' || lpad(to_hex(round(random() * 10000000)::int4),6,'0')
FROM plot_configuration_timeseries;

DROP VIEW IF EXISTS v_plot_configuration;
DROP TABLE plot_configuration_timeseries;
