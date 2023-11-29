INSERT INTO instrument (id, project_id, slug, name, geometry, type_id) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'ipi-1', 'Demo IPI 1', ST_GeomFromText('POINT(-80.8 26.7)',4326), 'c81f3a5d-fc5f-47fd-b545-401fe6ee63bb'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'ipi-2', 'Demo IPI 2', ST_GeomFromText('POINT(-80.8 26.7)',4326), 'c81f3a5d-fc5f-47fd-b545-401fe6ee63bb');

INSERT INTO instrument_status (instrument_id, status_id) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');

INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
('5842c707-b4be-4d10-a89c-1064e282e555', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-1-bottom-elevation', 'ipi-1-bottom-elevation'),
('f7fa0d85-c684-4315-a7c6-e18e60667969', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-1-tilt', 'ipi-1-segment-1-tilt'),
('1bf787e9-8363-4047-8b03-fbaf9ff03eaf', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-1-inc-dev', 'ipi-1-segment-1-inc-dev'),
('8d10fbd9-2669-4727-b4c1-746361691388', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-1-temp', 'ipi-1-segment-1-temp'),
('258a5834-20bf-45fc-a60c-f245b2822592', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-2-tilt', 'ipi-1-segment-2-tilt'),
('4ffcb98f-962a-46ea-8923-8f992ef07c58', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-2-inc-dev', 'ipi-1-segment-2-inc-dev'),
('6044cffb-c241-4b66-9873-068c2bbac451', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-2-temp', 'ipi-1-segment-2-temp'),
('3bd67db5-abd6-4b35-a649-427791f9eeb7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-3-tilt', 'ipi-1-segment-3-tilt'),
('1db6717b-6cde-4f46-b7fb-bc82b75051d7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-3-inc-dev', 'ipi-1-segment-3-inc-dev'),
('98385e5a-c5d8-4441-aa2e-0f6120414352', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-3-temp', 'ipi-1-segment-3-temp'),
('a3c4254b-1448-4f70-a1b6-d7f5e5c66eb7', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-4-tilt', 'ipi-1-segment-4-tilt'),
('6d90eb76-f292-461e-a82b-0faee9999778', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-4-inc-dev', 'ipi-1-segment-4-inc-dev'),
('c488fc08-18ff-4e3d-851f-46cfd1257b6c', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-1-segment-4-temp', 'ipi-1-segment-4-temp'),
('bce99683-59bd-4e4b-ad79-64a03553cfdc', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-1-segment-1-length', 'ipi-1-segment-1-length'),
('e891ca7c-59b2-41bc-9d4a-43995e35b855', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-1-segment-2-length', 'ipi-1-segment-2-length'),
('18f17db2-4bc8-44cb-a9fa-ba84d13b8444', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-1-segment-3-length', 'ipi-1-segment-3-length'),
('d5c236cf-dca5-4a35-bc59-a9ecac4d572b', 'e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-1-segment-4-length', 'ipi-1-segment-4-length'),
('7d515571-d6a2-4990-a1e2-d6d42049d864', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-2-bottom-elevation', 'ipi-2-bottom-elevation'),
('b2968456-b26a-4bbb-b8d9-f1217a6147ff', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-2-segment-1-tilt', 'ipi-2-segment-1-tilt'),
('afcc8471-c91b-466e-833d-f173cc58797f', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-2-segment-2-tilt', 'ipi-2-segment-2-tilt'),
('26cb2cfa-910a-46c3-b03f-9dbcf823f8d8', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-2-segment-3-tilt', 'ipi-2-segment-3-tilt'),
('3a297a4e-093a-4f9b-b201-1a994e2f4da7', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'ipi-2-segment-4-tilt', 'ipi-2-segment-4-tilt'),
('88accf78-6f41-4342-86b5-026a8880cbb4', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-2-segment-1-length', 'ipi-2-segment-1-length'),
('fc332ef5-55a8-4657-9d6d-b0abeeb985f2', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-2-segment-2-length', 'ipi-2-segment-2-length'),
('a86c7468-09a7-4090-98e0-f7979103bbcd', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-2-segment-3-length', 'ipi-2-segment-3-length'),
('d28efb95-962d-4233-9002-827154bd76ad', '01ac435f-fe3c-4af1-9979-f5e00467e7f5', 'a9a5ad45-b2e5-4744-816e-d3184f2c08bd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'ipi-2-segment-4-length', 'ipi-2-segment-4-length');

INSERT INTO ipi_opts (instrument_id, num_segments, bottom_elevation_timeseries_id, initial_time) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929', 4, '5842c707-b4be-4d10-a89c-1064e282e555', NOW() - INTERVAL '1 month'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5', 4, '7d515571-d6a2-4990-a1e2-d6d42049d864', NOW() - INTERVAL '1 month');


INSERT INTO ipi_segment (instrument_id, id, length_timeseries_id, tilt_timeseries_id, inc_dev_timeseries_id, temp_timeseries_id) VALUES
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929',1,'bce99683-59bd-4e4b-ad79-64a03553cfdc','f7fa0d85-c684-4315-a7c6-e18e60667969','1bf787e9-8363-4047-8b03-fbaf9ff03eaf', '8d10fbd9-2669-4727-b4c1-746361691388'),
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929',2,'e891ca7c-59b2-41bc-9d4a-43995e35b855','258a5834-20bf-45fc-a60c-f245b2822592','4ffcb98f-962a-46ea-8923-8f992ef07c58', '6044cffb-c241-4b66-9873-068c2bbac451'),
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929',3,'18f17db2-4bc8-44cb-a9fa-ba84d13b8444','3bd67db5-abd6-4b35-a649-427791f9eeb7','1db6717b-6cde-4f46-b7fb-bc82b75051d7', '98385e5a-c5d8-4441-aa2e-0f6120414352'),
('e29a8c6d-c5a4-4fcc-b269-3a60bd48f929',4,'d5c236cf-dca5-4a35-bc59-a9ecac4d572b','a3c4254b-1448-4f70-a1b6-d7f5e5c66eb7','6d90eb76-f292-461e-a82b-0faee9999778', 'c488fc08-18ff-4e3d-851f-46cfd1257b6c'),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',1,'88accf78-6f41-4342-86b5-026a8880cbb4','b2968456-b26a-4bbb-b8d9-f1217a6147ff', NULL, NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',2,'fc332ef5-55a8-4657-9d6d-b0abeeb985f2','afcc8471-c91b-466e-833d-f173cc58797f', NULL, NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',3,'a86c7468-09a7-4090-98e0-f7979103bbcd','26cb2cfa-910a-46c3-b03f-9dbcf823f8d8', NULL, NULL),
('01ac435f-fe3c-4af1-9979-f5e00467e7f5',4,'d28efb95-962d-4233-9002-827154bd76ad','3a297a4e-093a-4f9b-b201-1a994e2f4da7', NULL, NULL);


INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES
('5842c707-b4be-4d10-a89c-1064e282e555', NOW() - INTERVAL '1 month', 0),
('7d515571-d6a2-4990-a1e2-d6d42049d864', NOW() - INTERVAL '1 month', 50),
('bce99683-59bd-4e4b-ad79-64a03553cfdc', NOW() - INTERVAL '1 month', 012),
('e891ca7c-59b2-41bc-9d4a-43995e35b855', NOW() - INTERVAL '1 month', 123),
('18f17db2-4bc8-44cb-a9fa-ba84d13b8444', NOW() - INTERVAL '1 month', 234),
('d5c236cf-dca5-4a35-bc59-a9ecac4d572b', NOW() - INTERVAL '1 month', 345),
('88accf78-6f41-4342-86b5-026a8880cbb4', NOW() - INTERVAL '1 month', 100),
('fc332ef5-55a8-4657-9d6d-b0abeeb985f2', NOW() - INTERVAL '1 month', 200),
('a86c7468-09a7-4090-98e0-f7979103bbcd', NOW() - INTERVAL '1 month', 150),
('d28efb95-962d-4233-9002-827154bd76ad', NOW() - INTERVAL '1 month', 050);


INSERT INTO instrument_constants (timeseries_id, instrument_id) VALUES
('5842c707-b4be-4d10-a89c-1064e282e555','e29a8c6d-c5a4-4fcc-b269-3a60bd48f929'),
('7d515571-d6a2-4990-a1e2-d6d42049d864','01ac435f-fe3c-4af1-9979-f5e00467e7f5'),
('bce99683-59bd-4e4b-ad79-64a03553cfdc','e29a8c6d-c5a4-4fcc-b269-3a60bd48f929'),
('e891ca7c-59b2-41bc-9d4a-43995e35b855','e29a8c6d-c5a4-4fcc-b269-3a60bd48f929'),
('18f17db2-4bc8-44cb-a9fa-ba84d13b8444','e29a8c6d-c5a4-4fcc-b269-3a60bd48f929'),
('d5c236cf-dca5-4a35-bc59-a9ecac4d572b','e29a8c6d-c5a4-4fcc-b269-3a60bd48f929'),
('88accf78-6f41-4342-86b5-026a8880cbb4','01ac435f-fe3c-4af1-9979-f5e00467e7f5'),
('fc332ef5-55a8-4657-9d6d-b0abeeb985f2','01ac435f-fe3c-4af1-9979-f5e00467e7f5'),
('a86c7468-09a7-4090-98e0-f7979103bbcd','01ac435f-fe3c-4af1-9979-f5e00467e7f5'),
('d28efb95-962d-4233-9002-827154bd76ad','01ac435f-fe3c-4af1-9979-f5e00467e7f5');


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
        'b2968456-b26a-4bbb-b8d9-f1217a6147ff'::UUID,
        'afcc8471-c91b-466e-833d-f173cc58797f'::UUID,
        '26cb2cfa-910a-46c3-b03f-9dbcf823f8d8'::UUID,
        '3a297a4e-093a-4f9b-b201-1a994e2f4da7'::UUID,
        '8d10fbd9-2669-4727-b4c1-746361691388'::UUID,
        '6044cffb-c241-4b66-9873-068c2bbac451'::UUID,
        '98385e5a-c5d8-4441-aa2e-0f6120414352'::UUID,
        'c488fc08-18ff-4e3d-851f-46cfd1257b6c'::UUID
]) AS timeseries_id,
    generate_series(
        now() - INTERVAL '1 month',
        now(),
        INTERVAL '1 hour'
    ) AS time;
