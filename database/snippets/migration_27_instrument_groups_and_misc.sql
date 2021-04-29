
-- timeseries_measurements - add primary key which should also index
ALTER TABLE timeseries_measurement ADD PRIMARY KEY (timeseries_id, time);

--add new view for instrument_groups
-- v_instrument_group
CREATE OR REPLACE VIEW v_instrument_group AS (
    WITH instrument_count AS (
        SELECT 
        igi.instrument_group_id,
        count(igi.instrument_group_id) as i_count 
        FROM instrument_group_instruments igi
        JOIN instrument i on igi.instrument_id = i.id and not i.deleted
        GROUP BY igi.instrument_group_id
        )
        ,
        timeseries_instruments as (
            SELECT t.id, t.instrument_id, igi.instrument_group_id from timeseries t 
            JOIN instrument i on i.id = t.instrument_id and not i.deleted
            JOIN instrument_group_instruments igi on igi.instrument_id = i.id
        )

        SELECT  ig.id,
                ig.slug,
                ig.name,
                ig.description,
                ig.creator,
                ig.create_date,
                ig.updater,
                ig.update_date,
                ig.project_id,
                ig.deleted,
                COALESCE(ic.i_count,0) as instrument_count,
                COALESCE(count(ti.id),0) as timeseries_count
                --,
                --COALESCE(count(tm.id),0) as timeseries_measurements_count
                
        FROM instrument_group ig
        LEFT JOIN instrument_count ic on ic.instrument_group_id = ig.id
        LEFT JOIN timeseries_instruments ti on ig.id = ti.instrument_group_id
        --left join timeseries_measurement tm on tm.timeseries_id = ti.id
        GROUP BY ig.id, ic.i_count
        ORDER BY ig.name
);

-- grant select on new view

GRANT SELECT ON v_instrument_group TO instrumentation_reader;