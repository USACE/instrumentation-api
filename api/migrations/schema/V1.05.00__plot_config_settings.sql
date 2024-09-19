UPDATE plot_configuration_settings s SET date_range=ss.utc_format
FROM (
  WITH ab AS (
    SELECT
      split_part(date_range, ' - ', 1) AS a,
      split_part(date_range, ' - ', 2) AS b
    FROM plot_configuration_settings
    WHERE date_range LIKE '__/__/____ - __/__/____'
  ),
  formatted AS (
    SELECT
      split_part(a, '/', 1) AS a_month,
      split_part(a, '/', 2) AS a_day,
      split_part(a, '/', 3) AS a_year,
      split_part(b, '/', 1) AS b_month,
      split_part(b, '/', 2) AS b_day,
      split_part(b, '/', 3) AS b_year
    FROM ab
  )
  SELECT concat(
    a_year, '-', a_month, '-', a_day,
    ' ',
    b_year, '-', b_month, '-', b_day
  ) AS utc_format
  FROM formatted
) ss
WHERE date_range LIKE '__/__/____ - __/__/____';
