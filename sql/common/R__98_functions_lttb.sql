-- LTTB (Largest-Triangle-Three-Buckets) SQL implementation for downsampling without losing the shape of the plot
-- i.e. keeps important outliers like peaks and valleys. This can be applied conditionally depending on the density of
-- data and the area requested. For example, if a user requests 10 years worth of data sampled as 15 minute intervals,
-- this will need to be downsampled to something that the client can reasonably consume. Adusting the BucketSize based on the
-- estimated number of pixels can return an approximate number of desired points. Since counting all of the rows is expensive,
-- we can estimate using the sample rate (i.e. 15 minutes in this case) and min / max time range. Something like:
--      SELECT extract(epoch from max(time) - min(time)) / (60 * 15) AS n_samples
--
-- https://datylon.medium.com/sampling-time-series-data-sets-fc16caefff1b
--
DO $$
BEGIN
    CREATE TYPE point_t AS (x float, y float);
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating type point_t -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE TYPE triangle_t AS (p1 point_t, p2 point_t, p3 point_t, surface float);
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating type triangle_t -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE TYPE ts_measurement AS (id uuid, t timestamptz, v float);
    EXCEPTION WHEN DUPLICATE_OBJECT THEN
    RAISE NOTICE 'not creating type ts_measurement -- it already exists';
END
$$;

CREATE OR REPLACE FUNCTION triangle_surface(p1 point_t,p2 point_t,p3 point_t)
RETURNS float
LANGUAGE SQL
AS $$
    SELECT abs(p1.x * (p2.y - p3.y) + p2.x * (p3.y - p1.y) + p3.x * ( p1.y - p2.y)) / 2
$$;

CREATE OR REPLACE FUNCTION largest_triangle_accum(maxsurfacetriangle triangle_t, p1 point_t, p2 point_t, p3 point_t)
RETURNS triangle_t
LANGUAGE SQL
AS $$
    SELECT
      CASE
        WHEN maxsurfacetriangle IS NULL OR triangle_surface(p1,p2,p3) > maxsurfacetriangle.surface
          THEN (p1,p2,p3,triangle_surface(p1,p2,p3))::triangle_t
          ELSE maxsurfacetriangle
        END
$$;

CREATE OR REPLACE AGGREGATE largest_triangle(point_t,point_t,point_t) (
    stype = triangle_t,
    sfunc = largest_triangle_accum
);

CREATE OR REPLACE FUNCTION lttb(bucket_size int, dataset ts_measurement[])
RETURNS SETOF ts_measurement
LANGUAGE SQL
AS $$
  WITH RECURSIVE inputparams AS (
    SELECT bucket_size AS BucketSize
  ),
  timeseries AS (
    SELECT * FROM unnest(dataset)
  ),
  -- First and last timestamp, value of the timeseries
  tsrange AS (
      SELECT
          (SELECT
              (extract(epoch FROM t), v)::point_t
          FROM timeseries
          ORDER BY t ASC LIMIT 1) AS frst,
          (SELECT
              (extract(epoch from t), v)::point_t
          FROM timeseries
          ORDER BY t DESC LIMIT 1) AS lst
  ),
  -- Add bucket number (grp column) for all but the last bucket
  withgrptmp as (
      SELECT 1 AS grp, (tsr.frst::point_t).x, (tsr.frst::point_t).y FROM tsrange tsr
      UNION
      SELECT
          1 + dense_rank() OVER (ORDER BY i.BucketSize * cast(extract(epoch FROM ts.t) / i.BucketSize AS int)) AS grp,
          extract(epoch FROM ts.t),
          ts.v AS val
      FROM timeseries ts, tsrange tsr, inputparams i
      WHERE
          ts.t > to_timestamp((tsr.frst::point_t).x) AT TIME ZONE 'utc'
          AND ts.t < to_timestamp((tsr.lst::point_t).x) AT TIME ZONE 'utc'
  ),
  -- Add bucket number (grp column)
  withgrp AS (
    SELECT * FROM withgrptmp
    UNION
    SELECT 1 + (SELECT max(grp) FROM withgrptmp) AS grp, (tsr.lst::point_t).x, (tsr.lst::point_t).y FROM tsrange tsr
  ),
  -- Average per bucket
  withgrpavgtmp AS (
    SELECT grp,avg(x) AS xavg, avg(y) AS yavg FROM withgrp GROUP BY grp
  ),
  -- Join time series timestamp, value with average values of next bucket
  withgrpavg AS (
    SELECT
      withgrp.grp AS grp,
      withgrp.x,
      withgrp.y,
      withgrpavgtmp.xavg AS xavg3,
      withgrpavgtmp.yavg AS yavg3
    FROM withgrp LEFT OUTER JOIN withgrpavgtmp ON withgrp.grp=withgrpavgtmp.grp-1
  ),
  largesttriangle(grp,p) AS (
    SELECT
      wga.grp,
      ((0,0)::point_t,(wga.x,wga.y)::point_t,(0,0)::point_t,0.0::float)::triangle_t
    FROM withgrpavg wga
    WHERE grp = 1
    UNION ALL
    SELECT DISTINCT
      wga.grp,
      (largest_triangle(
        (ltt.p).p2::point_t,           -- the selected point of the previous bucket
        (wga.x,wga.y)::point_t,        -- current bucket
        (wga.xavg3,wga.yavg3)::point_t -- average of next bucket
      ) OVER (PARTITION BY wga.grp))::triangle_t
    FROM withgrpavg wga
    JOIN largesttriangle ltt ON wga.grp=ltt.grp+1
    WHERE wga.grp > 1
  ),
  postproc AS (
    SELECT
      to_timestamp(((t.p).p2::point_t).x) AT TIME ZONE 'utc' AS t,
      ((t.p).p2::point_t).y AS v
    FROM largesttriangle t
    WHERE ((t.p).p2::point_t).y != 0
    ORDER BY 1
  )
  SELECT
    src.id AS id,
    ppc.t AS t,
    ppc.v AS v
  FROM postproc ppc
  INNER JOIN timeseries src
    ON src.t = ppc.t
    AND src.v = ppc.v
$$;
