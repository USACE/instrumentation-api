-- ${flyway:timestamp}
CREATE VIEW v_plot_bullseye_measurement AS (
  SELECT
    pc.plot_config_id,
    COALESCE(xm.time, ym.time) AS time,
    locf(xm.value) OVER (ORDER BY xm.time) AS x,
    locf(ym.value) OVER (ORDER BY ym.time) AS y
  FROM plot_bullseye_config pc
  LEFT JOIN timeseries_measurement xm ON xm.timeseries_id = pc.x_timeseries_id
  LEFT JOIN timeseries_measurement ym ON ym.timeseries_id = pc.y_timeseries_id
);

CREATE VIEW v_plot_contour_measurement AS (
  SELECT
    pc.plot_config_id,
    mm.time AS time,
    ii.name AS instrument_name,
    ts.name AS timeseries_name,
    ST_X(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS x,
    ST_Y(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS y,
    mm.value AS z
  FROM plot_contour_config pc
  LEFT JOIN plot_contour_config_timeseries pcts ON pcts.plot_config_id = pc.plot_config_id
  LEFT JOIN timeseries ts ON ts.id = pcts.timeseries_id
  LEFT JOIN instrument ii ON ts.instrument_id = ii.id
  LEFT JOIN measurement mm ON mm.timeseries_id = ts.id
  GROUP BY pc.plot_config_id, mm.time
);
