-- ${flyway:timestamp}
DROP FUNCTION IF EXISTS lttb;
DROP FUNCTION IF EXISTS largest_triangle_accum;
DROP AGGREGATE IF EXISTS largest_triangle(point_t,point_t,point_t);
DROP FUNCTION IF EXISTS triangle_surface;

DROP TYPE IF EXISTS ts_measurement;
DROP TYPE IF EXISTS triangle_t;
DROP TYPE IF EXISTS point_t;
