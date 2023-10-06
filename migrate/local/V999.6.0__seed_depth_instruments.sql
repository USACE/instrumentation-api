INSERT INTO instrument (id, project_id, slug, name, geometry, type_id) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'saa-1', 'saa-1', ST_GeomFromText('POINT(-80.8 26.7)',4326), '07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9');


INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
-- seg 1
('8b3762ef-a852-4edc-8e87-746a92eaac9d', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-1', 'saa-x-1'),
('ecfa267b-339b-4bb8-b7ae-eda550257878', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-1', 'saa-y-1'),
('a31a24c4-aa8e-4e52-9895-43cdb69fe703', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-1', 'saa-z-1'),
('eec831d1-56a5-47ef-85eb-02c7622d6cb8', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-1', 'saa-t-1'),
-- seg 2                                                                                                                                                                            
('eb25ab9f-af8b-4383-839a-7d24899e02c4', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-2', 'saa-x-2'),
('8e641473-d7bf-433c-a24b-55fa065ca0c3', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-2', 'saa-y-2'),
('21cfe121-d29d-40a2-b04f-6be71ba479fe', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-2', 'saa-z-2'),
('23bda2f6-c479-48e0-a0c2-db48c3b08c3c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-2', 'saa-t-2'),
-- seg 3                                                                                                                                                                            
('2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-3', 'saa-x-3'),
('4759bdac-656e-47c3-b403-d3118cf57342', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-3', 'saa-y-3'),
('1f47a1b9-a2bb-4282-8618-42ba1341533e', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-3', 'saa-z-3'),
('d2dbac06-ad03-45d9-a7ad-1e7fb9d09ce2', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-3', 'saa-t-3'),
-- seg 4                                                                                                                                                                            
('c22ffd8a-eae3-41cb-a75b-faae36236465', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-4', 'saa-x-4'),
('d11a0e91-0125-46cc-a3fc-b0252361bd9c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-4', 'saa-y-4'),
('9fbf2061-cf73-45f3-9e6c-b745ae7f72a1', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-4', 'saa-z-4'),
('0503e693-bc58-49b5-a477-288174dc90ed', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-4', 'saa-t-4'),
-- seg 5                                                                                                                                                                            
('24ad9638-5c5e-48b6-9ad6-a2eb0b93f87c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-5', 'saa-x-5'),
('8cfaffb4-80b2-411b-be81-776385fc5862', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-5', 'saa-y-5'),
('ea0f561f-e3f4-4155-a360-17407a0884d4', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-5', 'saa-z-5'),
('a10e8627-621c-4aa7-8301-a2142a760e0c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-5', 'saa-t-5'),
-- seg 6                                                                                                                                                                            
('88e22274-021e-4e91-88bb-046b67171a36', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-6', 'saa-x-6'),
('f684bec8-9cc3-470f-a355-21d65f2be435', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-6', 'saa-y-6'),
('1a8c9bfc-0e65-4f76-aba9-fc32d643748f', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-6', 'saa-z-6'),
('2bf6aecd-3df0-4237-b28b-95731b7e333d', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-6', 'saa-t-6'),
-- seg 7                                                                                                                                                                            
('00f3e1f2-e7ff-4901-abfb-e9bf695802f6', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-7', 'saa-x-7'),
('2ef9b1d9-ee8f-4f2d-a482-2e0f0dd76f80', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-7', 'saa-y-7'),
('00ae950d-5bdd-455e-a72a-56da67dafb85', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-7', 'saa-z-7'),
('3d07cbc0-4aff-4efa-a162-ec1800801665', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-7', 'saa-t-7'),
-- seg 8                                                                                                                                                                            
('fb0795ba-9d80-4a41-abd7-5de140392454', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-x-8', 'saa-x-8'),
('32889a6d-93d0-49f9-b281-44e19e88474c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-y-8', 'saa-y-8'),
('bcb95c35-08f7-4c5a-83ff-b505b8d76481', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-z-8', 'saa-z-8'),
('54dcd1e1-e9da-4db5-95e5-3c28fab5c03c', 'eca4040e-aecb-4cd3-bcde-3e308f0356a6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'saa-t-8', 'saa-t-8');


INSERT INTO saa_instrument (instrument_id, num_segments, bottom_elevation, initial_time) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6', 8, 0, NULL);


INSERT INTO saa_segment (instrument_id, id, length, x_timeseries_id, y_timeseries_id, z_timeseries_id, temp_timeseries_id) VALUES
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',0,200,'8b3762ef-a852-4edc-8e87-746a92eaac9d','ecfa267b-339b-4bb8-b7ae-eda550257878','a31a24c4-aa8e-4e52-9895-43cdb69fe703','eec831d1-56a5-47ef-85eb-02c7622d6cb8'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',1,200,'eb25ab9f-af8b-4383-839a-7d24899e02c4','8e641473-d7bf-433c-a24b-55fa065ca0c3','21cfe121-d29d-40a2-b04f-6be71ba479fe','23bda2f6-c479-48e0-a0c2-db48c3b08c3c'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',2,200,'2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70','4759bdac-656e-47c3-b403-d3118cf57342','1f47a1b9-a2bb-4282-8618-42ba1341533e','d2dbac06-ad03-45d9-a7ad-1e7fb9d09ce2'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',3,200,'c22ffd8a-eae3-41cb-a75b-faae36236465','d11a0e91-0125-46cc-a3fc-b0252361bd9c','9fbf2061-cf73-45f3-9e6c-b745ae7f72a1','0503e693-bc58-49b5-a477-288174dc90ed'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',4,200,'24ad9638-5c5e-48b6-9ad6-a2eb0b93f87c','8cfaffb4-80b2-411b-be81-776385fc5862','ea0f561f-e3f4-4155-a360-17407a0884d4','a10e8627-621c-4aa7-8301-a2142a760e0c'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',5,200,'88e22274-021e-4e91-88bb-046b67171a36','f684bec8-9cc3-470f-a355-21d65f2be435','1a8c9bfc-0e65-4f76-aba9-fc32d643748f','2bf6aecd-3df0-4237-b28b-95731b7e333d'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',6,200,'00f3e1f2-e7ff-4901-abfb-e9bf695802f6','ecfa267b-339b-4bb8-b7ae-eda550257878','00ae950d-5bdd-455e-a72a-56da67dafb85','3d07cbc0-4aff-4efa-a162-ec1800801665'),
('eca4040e-aecb-4cd3-bcde-3e308f0356a6',7,200,'fb0795ba-9d80-4a41-abd7-5de140392454','32889a6d-93d0-49f9-b281-44e19e88474c','bcb95c35-08f7-4c5a-83ff-b505b8d76481','54dcd1e1-e9da-4db5-95e5-3c28fab5c03c');
