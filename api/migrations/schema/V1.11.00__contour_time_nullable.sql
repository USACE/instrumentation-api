ALTER TABLE plot_contour_config ALTER COLUMN time DROP NOT NULL;

DROP VIEW IF EXISTS v_instrument;
SELECT UpdateGeometrySRID('midas', 'instrument', 'geometry', 4326);
