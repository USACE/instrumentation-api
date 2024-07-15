INSERT INTO plot_contour_config (plot_config_id, "time", locf_backfill) VALUES
('cc28ca81-f125-46c6-a5cd-cc055a003c19', now(), '10 years'::interval);

INSERT INTO plot_contour_config_timeseries (plot_contour_config_id, timeseries_id) VALUES
('cc28ca81-f125-46c6-a5cd-cc055a003c19', '869465fc-dc1e-445e-81f4-9979b5fadda9'),
('cc28ca81-f125-46c6-a5cd-cc055a003c19', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c');
