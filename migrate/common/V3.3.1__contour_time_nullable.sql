ALTER TABLE plot_contour_config ALTER COLUMN time DROP NOT NULL;

SELECT UpdateGeometrySRID('midas', 'instrument', 'geometry', 4326);
