INSERT INTO instrument (id, project_id, slug, name, geometry, type_id) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'ipi-1', 'Demo IPI 1', ST_GeomFromText('POINT(-80.8 26.7)',4326), 'c81f3a5d-fc5f-47fd-b545-401fe6ee63bb'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'ipi-2', 'Demo IPI 2', ST_GeomFromText('POINT(-80.8 26.7)',4326), 'c81f3a5d-fc5f-47fd-b545-401fe6ee63bb');

INSERT INTO instrument_status (instrument_id, status_id) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');

INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
('f7fa0d85-c684-4315-a7c6-e18e60667969', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-1', 'ipi-1-tilt-1'),
('1bf787e9-8363-4047-8b03-fbaf9ff03eaf', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-y-1', 'ipi-1-cum-dev-1'),
('258a5834-20bf-45fc-a60c-f245b2822592', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-2', 'ipi-1-tilt-2'),
('4ffcb98f-962a-46ea-8923-8f992ef07c58', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-y-2', 'ipi-1-cum-dev-2'),
('3bd67db5-abd6-4b35-a649-427791f9eeb7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-3', 'ipi-1-tilt-3'),
('1db6717b-6cde-4f46-b7fb-bc82b75051d7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-y-3', 'ipi-1-cum-dev-3'),
('a3c4254b-1448-4f70-a1b6-d7f5e5c66eb7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-4', 'ipi-1-tilt-4'),
('6d90eb76-f292-461e-a82b-0faee9999778', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-y-4', 'ipi-1-cum-dev-4'),
('88accf78-6f41-4342-86b5-026a8880cbb4', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-1', 'ipi-2-tilt-1'),
('afcc8471-c91b-466e-833d-f173cc58797f', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-2', 'ipi-2-tilt-2'),
('26cb2cfa-910a-46c3-b03f-9dbcf823f8d8', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-3', 'ipi-2-tilt-3'),
('3a297a4e-093a-4f9b-b201-1a994e2f4da7', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-x-4', 'ipi-2-tilt-4');


INSERT INTO ipi_opts (instrument_id, num_segments, bottom_elevation, initial_time) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 4, 100, NOW() - INTERVAL '1 month'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', 4, 100, NOW() - INTERVAL '1 month');


INSERT INTO ipi_segment (instrument_id, id, length, tilt_timeseries_id, cum_dev_timeseries_id) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',1,012,'f7fa0d85-c684-4315-a7c6-e18e60667969','1bf787e9-8363-4047-8b03-fbaf9ff03eaf'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',2,123,'258a5834-20bf-45fc-a60c-f245b2822592','4ffcb98f-962a-46ea-8923-8f992ef07c58'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',3,234,'3bd67db5-abd6-4b35-a649-427791f9eeb7','1db6717b-6cde-4f46-b7fb-bc82b75051d7'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',4,345,'a3c4254b-1448-4f70-a1b6-d7f5e5c66eb7','6d90eb76-f292-461e-a82b-0faee9999778'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',1,100,'88accf78-6f41-4342-86b5-026a8880cbb4', NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',2,200,'afcc8471-c91b-466e-833d-f173cc58797f', NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',3,150,'26cb2cfa-910a-46c3-b03f-9dbcf823f8d8', NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',4,050,'3a297a4e-093a-4f9b-b201-1a994e2f4da7', NULL);


INSERT INTO timeseries_measurement (timeseries_id, time, value)
SELECT
    timeseries_id,
    time,
    round((random() * (100-3) + 3)::NUMERIC, 4) AS value
FROM
    unnest(ARRAY[
        'f7fa0d85-c684-4315-a7c6-e18e60667969'::UUID,
        '1bf787e9-8363-4047-8b03-fbaf9ff03eaf'::UUID,
        '258a5834-20bf-45fc-a60c-f245b2822592'::UUID,
        '4ffcb98f-962a-46ea-8923-8f992ef07c58'::UUID,
        '3bd67db5-abd6-4b35-a649-427791f9eeb7'::UUID,
        '1db6717b-6cde-4f46-b7fb-bc82b75051d7'::UUID,
        'a3c4254b-1448-4f70-a1b6-d7f5e5c66eb7'::UUID,
        '6d90eb76-f292-461e-a82b-0faee9999778'::UUID,
        '88accf78-6f41-4342-86b5-026a8880cbb4'::UUID,
        'afcc8471-c91b-466e-833d-f173cc58797f'::UUID,
        '26cb2cfa-910a-46c3-b03f-9dbcf823f8d8'::UUID,
        '3a297a4e-093a-4f9b-b201-1a994e2f4da7'::UUID
]) AS timeseries_id,
    generate_series(
        now() - INTERVAL '1 month',
        now(),
        INTERVAL '1 hour'
    ) AS time;
