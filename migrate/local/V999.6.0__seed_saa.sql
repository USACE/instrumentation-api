INSERT INTO instrument (id, project_id, slug, name, geometry, type_id) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'saa-1', 'Demo SAA 1', ST_GeomFromText('POINT(-80.8 26.7)',4326), '07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9');

INSERT INTO instrument_status (instrument_id, status_id) VALUES ('eca4040e-aecb-4cd3-bcde-3e308f0356a6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');

INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
('4affc367-ea0f-41f5-a4bc-5f387b01d7a4', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-bottom-elevation', 'saa-1-bottom-elevation'),
('cf2f2304-d44e-4363-bc8d-95533222efd6', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-1-length', 'saa-1-segment-1-length'),
('ff2086ae-0eae-42a8-b598-2e97be2ab3b0', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-2-length', 'saa-1-segment-2-length'),
('669b63d7-87b2-4aed-9b15-e19ea39789b9', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-3-length', 'saa-1-segment-3-length'),
('e404e8f4-41c6-4355-9ddb-9d8c635525fc', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-4-length', 'saa-1-segment-4-length'),
('ccb80fd4-8902-450f-bb3b-cc1e6718b03c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-5-length', 'saa-1-segment-5-length'),
('7f98f239-ac1e-4651-9d69-c163b2dc06a6', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-6-length', 'saa-1-segment-6-length'),
('72bd19f1-23d3-4edb-b16f-9ebb121cf921', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-7-length', 'saa-1-segment-7-length'),
('df6a9cca-29fc-4ec3-9415-d497fbae1a58', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '6d12ca4c-b618-41cd-87a2-a248980a0d69', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'saa-1-segment-8-length', 'saa-1-segment-8-length'),
('8b3762ef-a852-4edc-8e87-746a92eaac9d', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-1-x', 'saa-1-segment-1-x'),
('ecfa267b-339b-4bb8-b7ae-eda550257878', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-1-y', 'saa-1-segment-1-y'),
('a31a24c4-aa8e-4e52-9895-43cdb69fe703', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-1-z', 'saa-1-segment-1-z'),
('eec831d1-56a5-47ef-85eb-02c7622d6cb8', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-1-t', 'saa-1-segment-1-t'),
('eb25ab9f-af8b-4383-839a-7d24899e02c4', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-2-x', 'saa-1-segment-2-x'),
('8e641473-d7bf-433c-a24b-55fa065ca0c3', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-2-y', 'saa-1-segment-2-y'),
('21cfe121-d29d-40a2-b04f-6be71ba479fe', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-2-z', 'saa-1-segment-2-z'),
('23bda2f6-c479-48e0-a0c2-db48c3b08c3c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-2-t', 'saa-1-segment-2-t'),
('2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-3-x', 'saa-1-segment-3-x'),
('4759bdac-656e-47c3-b403-d3118cf57342', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-3-y', 'saa-1-segment-3-y'),
('1f47a1b9-a2bb-4282-8618-42ba1341533e', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-3-z', 'saa-1-segment-3-z'),
('d2dbac06-ad03-45d9-a7ad-1e7fb9d09ce2', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-3-t', 'saa-1-segment-3-t'),
('c22ffd8a-eae3-41cb-a75b-faae36236465', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-4-x', 'saa-1-segment-4-x'),
('d11a0e91-0125-46cc-a3fc-b0252361bd9c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-4-y', 'saa-1-segment-4-y'),
('9fbf2061-cf73-45f3-9e6c-b745ae7f72a1', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-4-z', 'saa-1-segment-4-z'),
('0503e693-bc58-49b5-a477-288174dc90ed', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-4-t', 'saa-1-segment-4-t'),
('24ad9638-5c5e-48b6-9ad6-a2eb0b93f87c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-5-x', 'saa-1-segment-5-x'),
('8cfaffb4-80b2-411b-be81-776385fc5862', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-5-y', 'saa-1-segment-5-y'),
('ea0f561f-e3f4-4155-a360-17407a0884d4', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-5-z', 'saa-1-segment-5-z'),
('a10e8627-621c-4aa7-8301-a2142a760e0c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-5-t', 'saa-1-segment-5-t'),
('88e22274-021e-4e91-88bb-046b67171a36', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-6-x', 'saa-1-segment-6-x'),
('f684bec8-9cc3-470f-a355-21d65f2be435', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-6-y', 'saa-1-segment-6-y'),
('1a8c9bfc-0e65-4f76-aba9-fc32d643748f', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-6-z', 'saa-1-segment-6-z'),
('2bf6aecd-3df0-4237-b28b-95731b7e333d', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-6-t', 'saa-1-segment-6-t'),
('00f3e1f2-e7ff-4901-abfb-e9bf695802f6', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-7-x', 'saa-1-segment-7-x'),
('2ef9b1d9-ee8f-4f2d-a482-2e0f0dd76f80', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-7-y', 'saa-1-segment-7-y'),
('00ae950d-5bdd-455e-a72a-56da67dafb85', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-7-z', 'saa-1-segment-7-z'),
('3d07cbc0-4aff-4efa-a162-ec1800801665', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-7-t', 'saa-1-segment-7-t'),
('fb0795ba-9d80-4a41-abd7-5de140392454', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-8-x', 'saa-1-segment-8-x'),
('32889a6d-93d0-49f9-b281-44e19e88474c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-8-y', 'saa-1-segment-8-y'),
('bcb95c35-08f7-4c5a-83ff-b505b8d76481', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-8-z', 'saa-1-segment-8-z'),
('54dcd1e1-e9da-4db5-95e5-3c28fab5c03c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-1-segment-8-t', 'saa-1-segment-8-t');


INSERT INTO saa_opts (instrument_id, num_segments, bottom_elevation_timeseries_id, initial_time) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6', 8, '4affc367-ea0f-41f5-a4bc-5f387b01d7a4', NOW() - INTERVAL '1 month');


INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES
('4affc367-ea0f-41f5-a4bc-5f387b01d7a4', NOW() - INTERVAL '1 month', 0),
('cf2f2304-d44e-4363-bc8d-95533222efd6', NOW() - INTERVAL '1 month', 200),
('ff2086ae-0eae-42a8-b598-2e97be2ab3b0', NOW() - INTERVAL '1 month', 200),
('669b63d7-87b2-4aed-9b15-e19ea39789b9', NOW() - INTERVAL '1 month', 200),
('e404e8f4-41c6-4355-9ddb-9d8c635525fc', NOW() - INTERVAL '1 month', 200),
('ccb80fd4-8902-450f-bb3b-cc1e6718b03c', NOW() - INTERVAL '1 month', 200),
('7f98f239-ac1e-4651-9d69-c163b2dc06a6', NOW() - INTERVAL '1 month', 200),
('72bd19f1-23d3-4edb-b16f-9ebb121cf921', NOW() - INTERVAL '1 month', 200),
('df6a9cca-29fc-4ec3-9415-d497fbae1a58', NOW() - INTERVAL '1 month', 200);


INSERT INTO saa_segment (instrument_id, id, length_timeseries_id, x_timeseries_id, y_timeseries_id, z_timeseries_id, temp_timeseries_id) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',1,'cf2f2304-d44e-4363-bc8d-95533222efd6','8b3762ef-a852-4edc-8e87-746a92eaac9d','ecfa267b-339b-4bb8-b7ae-eda550257878','a31a24c4-aa8e-4e52-9895-43cdb69fe703','eec831d1-56a5-47ef-85eb-02c7622d6cb8'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',2,'ff2086ae-0eae-42a8-b598-2e97be2ab3b0','eb25ab9f-af8b-4383-839a-7d24899e02c4','8e641473-d7bf-433c-a24b-55fa065ca0c3','21cfe121-d29d-40a2-b04f-6be71ba479fe','23bda2f6-c479-48e0-a0c2-db48c3b08c3c'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',3,'669b63d7-87b2-4aed-9b15-e19ea39789b9','2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70','4759bdac-656e-47c3-b403-d3118cf57342','1f47a1b9-a2bb-4282-8618-42ba1341533e','d2dbac06-ad03-45d9-a7ad-1e7fb9d09ce2'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',4,'e404e8f4-41c6-4355-9ddb-9d8c635525fc','c22ffd8a-eae3-41cb-a75b-faae36236465','d11a0e91-0125-46cc-a3fc-b0252361bd9c','9fbf2061-cf73-45f3-9e6c-b745ae7f72a1','0503e693-bc58-49b5-a477-288174dc90ed'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',5,'ccb80fd4-8902-450f-bb3b-cc1e6718b03c','24ad9638-5c5e-48b6-9ad6-a2eb0b93f87c','8cfaffb4-80b2-411b-be81-776385fc5862','ea0f561f-e3f4-4155-a360-17407a0884d4','a10e8627-621c-4aa7-8301-a2142a760e0c'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',6,'7f98f239-ac1e-4651-9d69-c163b2dc06a6','88e22274-021e-4e91-88bb-046b67171a36','f684bec8-9cc3-470f-a355-21d65f2be435','1a8c9bfc-0e65-4f76-aba9-fc32d643748f','2bf6aecd-3df0-4237-b28b-95731b7e333d'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',7,'72bd19f1-23d3-4edb-b16f-9ebb121cf921','00f3e1f2-e7ff-4901-abfb-e9bf695802f6','ecfa267b-339b-4bb8-b7ae-eda550257878','00ae950d-5bdd-455e-a72a-56da67dafb85','3d07cbc0-4aff-4efa-a162-ec1800801665'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',8,'df6a9cca-29fc-4ec3-9415-d497fbae1a58','fb0795ba-9d80-4a41-abd7-5de140392454','32889a6d-93d0-49f9-b281-44e19e88474c','bcb95c35-08f7-4c5a-83ff-b505b8d76481','54dcd1e1-e9da-4db5-95e5-3c28fab5c03c');


INSERT INTO timeseries_measurement (timeseries_id, time, value)
SELECT
    timeseries_id,
    time,
    round((random() * (100-3) + 3)::NUMERIC, 4) AS value
FROM
    unnest(ARRAY[
        '8b3762ef-a852-4edc-8e87-746a92eaac9d'::uuid,
        'ecfa267b-339b-4bb8-b7ae-eda550257878'::uuid,
        'a31a24c4-aa8e-4e52-9895-43cdb69fe703'::uuid,
        'eec831d1-56a5-47ef-85eb-02c7622d6cb8'::uuid,
        'eb25ab9f-af8b-4383-839a-7d24899e02c4'::uuid,
        '8e641473-d7bf-433c-a24b-55fa065ca0c3'::uuid,
        '21cfe121-d29d-40a2-b04f-6be71ba479fe'::uuid,
        '23bda2f6-c479-48e0-a0c2-db48c3b08c3c'::uuid,
        '2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70'::uuid,
        '4759bdac-656e-47c3-b403-d3118cf57342'::uuid,
        '1f47a1b9-a2bb-4282-8618-42ba1341533e'::uuid,
        'd2dbac06-ad03-45d9-a7ad-1e7fb9d09ce2'::uuid,
        'c22ffd8a-eae3-41cb-a75b-faae36236465'::uuid,
        'd11a0e91-0125-46cc-a3fc-b0252361bd9c'::uuid,
        '9fbf2061-cf73-45f3-9e6c-b745ae7f72a1'::uuid,
        '0503e693-bc58-49b5-a477-288174dc90ed'::uuid,
        '24ad9638-5c5e-48b6-9ad6-a2eb0b93f87c'::uuid,
        '8cfaffb4-80b2-411b-be81-776385fc5862'::uuid,
        'ea0f561f-e3f4-4155-a360-17407a0884d4'::uuid,
        'a10e8627-621c-4aa7-8301-a2142a760e0c'::uuid,
        '88e22274-021e-4e91-88bb-046b67171a36'::uuid,
        'f684bec8-9cc3-470f-a355-21d65f2be435'::uuid,
        '1a8c9bfc-0e65-4f76-aba9-fc32d643748f'::uuid,
        '2bf6aecd-3df0-4237-b28b-95731b7e333d'::uuid,
        '00f3e1f2-e7ff-4901-abfb-e9bf695802f6'::uuid,
        '2ef9b1d9-ee8f-4f2d-a482-2e0f0dd76f80'::uuid,
        '00ae950d-5bdd-455e-a72a-56da67dafb85'::uuid,
        '3d07cbc0-4aff-4efa-a162-ec1800801665'::uuid,
        'fb0795ba-9d80-4a41-abd7-5de140392454'::uuid,
        '32889a6d-93d0-49f9-b281-44e19e88474c'::uuid,
        'bcb95c35-08f7-4c5a-83ff-b505b8d76481'::uuid,
        '54dcd1e1-e9da-4db5-95e5-3c28fab5c03c'::uuid
    ]) AS timeseries_id,
    generate_series(
        now() - INTERVAL '1 month',
        now(),
        INTERVAL '1 hour'
    ) AS time;
