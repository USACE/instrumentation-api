-- Project
INSERT INTO project (id, slug, name, image) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'buffalo-district-streamgages', 'Buffalo District Streamgages', 'buffalo-district-streamgages.jpg');




--INSERT INSTRUMENTS--COUNT:29
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('089fbda3-5a6a-408d-aec2-4e108910d94b', False, 'avnn6', 'AVNN6', null, ST_GeomFromText('POINT(-77.7566 42.9184)',4326), null, null, '2021-03-02T13:03:50.008075Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', False, 'blbn6', 'BLBN6', null, ST_GeomFromText('POINT(-77.6806 43.0922)',4326), null, null, '2021-03-02T13:03:50.008416Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', False, 'chcn6', 'CHCN6', null, ST_GeomFromText('POINT(-77.8822 43.1008)',4326), null, null, '2021-03-02T13:03:50.008576Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', False, 'dsvn6', 'DSVN6', null, ST_GeomFromText('POINT(-77.7064 42.5322)',4326), null, null, '2021-03-02T13:03:50.008752Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224775'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', False, 'garn6', 'GARN6', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-02T13:03:50.008929Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', False, 'hnyn6', 'HNYN6', null, ST_GeomFromText('POINT(-77.5869 42.9567)',4326), null, null, '2021-03-02T13:03:50.009105Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04229500'),
('08735594-14f0-4594-aa38-fb0805968918', False, 'blackcr-churchvl', 'BlackCr Churchvl', null, ST_GeomFromText('POINT(-77.8822 43.1006)',4326), null, null, '2021-03-02T13:03:50.009283Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', False, 'genr-portagevill', 'GenR Portagevill', null, ST_GeomFromText('POINT(-78.0422 42.5703)',4326), null, null, '2021-03-02T13:03:50.009431Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', False, 'oatkacr-garbutt', 'OatkaCr Garbutt', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-02T13:03:50.009606Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', False, 'knvn6', 'KNVN6', null, ST_GeomFromText('POINT(-78.3103 43.3011)',4326), null, null, '2021-03-02T13:03:50.009749Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '0422016550'),
('3b260872-3454-473b-bfb7-087ed0e809b0', False, 'mbyp1', 'MBYP1', null, ST_GeomFromText('POINT(-77.2736 41.8425)',4326), null, null, '2021-03-02T13:03:50.009941Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01518420'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', False, 'olnn6', 'OLNN6', null, ST_GeomFromText('POINT(-78.4511 42.0731)',4326), null, null, '2021-03-02T13:03:50.010146Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010820'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', False, 'rohn6', 'ROHN6', null, ST_GeomFromText('POINT(-77.6163 43.1417)',4326), null, null, '2021-03-02T13:03:50.010325Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231600'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', False, 'shnp1', 'SHNP1', null, ST_GeomFromText('POINT(-78.1983 41.9617)',4326), null, null, '2021-03-02T13:03:50.010683Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010655'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', False, 'genr-wellsville', 'GenR Wellsville', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-02T13:03:50.010861Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', False, 'jonn6', 'JONN6', null, ST_GeomFromText('POINT(-77.8386 42.7667)',4326), null, null, '2021-03-02T13:03:50.011037Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227500'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', False, 'mount-morris', 'Mount Morris', null, ST_GeomFromText('POINT(-77.9071 42.7333)',4326), null, null, '2021-03-02T13:03:50.011210Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224000'),
('0bac5d43-fab7-48b3-9fbc-33c5d565c1a9', False, 'mount-morris-tailwater', 'Mount Morris-Tailwater', null, ST_GeomFromText('POINT(-77.9109 42.7332)',4326), null, null, '2021-03-02T13:03:50.011210Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('90664d51-0503-45cf-85bb-5d46fa579d0f', False, 'ptgn6', 'PTGN6', null, ST_GeomFromText('POINT(-78.0431 42.5697)',4326), null, null, '2021-03-02T13:03:50.011449Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', False, 'weln6', 'WELN6', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-02T13:03:50.011619Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('f063108f-9c3b-49fd-b981-1be021278ebf', False, 'wrsn6', 'WRSN6', null, ST_GeomFromText('POINT(-78.1375 42.7447)',4326), null, null, '2021-03-02T13:03:50.011794Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('15614ca3-c115-4cff-a021-c43ef352fda8', False, 'rcrn6', 'RCRN6', null, ST_GeomFromText('POINT(-77.6025 43.258)',4326), null, null, '2021-03-02T13:03:50.011929Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('31be5853-f839-4cc2-80cb-071210651059', False, 'elkp1', 'ELKP1', null, ST_GeomFromText('POINT(-77.3025 41.9875)',4326), null, null, '2021-03-02T13:03:50.012131Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01519200'),
('8df04db9-392f-4216-9776-66defd9d8d32', False, 'frkn6', 'FRKN6', null, ST_GeomFromText('POINT(-78.4636 42.3294)',4326), null, null, '2021-03-02T13:03:50.012308Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '421946078274901'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', False, 'hrln6', 'HRLN6', null, ST_GeomFromText('POINT(-77.7044 42.3489)',4326), null, null, '2021-03-02T13:03:50.012476Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01523000'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', False, 'canaseragashaker', 'CanaseragaShaker', null, ST_GeomFromText('POINT(-77.8414 42.7361)',4326), null, null, '2021-03-02T13:03:50.012668Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227000'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', False, 'oatkacr-warsaw', 'OatkaCr Warsaw', null, ST_GeomFromText('POINT(-78.1375 42.7442)',4326), null, null, '2021-03-02T13:03:50.012919Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230380'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', False, 'genr-avon', 'GenR Avon', null, ST_GeomFromText('POINT(-77.7572 42.9178)',4326), null, null, '2021-03-02T13:03:50.013093Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', False, 'akln6', 'AKLN6', null, ST_GeomFromText('POINT(-77.7167 42.3958)',4326), null, null, '2021-03-02T13:03:50.013291Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01521000');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('64651b76-ae43-490f-8ab3-5e0324e35163', '089fbda3-5a6a-408d-aec2-4e108910d94b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.008075Z'),
('9deb54cc-2afe-487b-b531-9c35af3e2d38', '209980f3-ab26-42d6-9fe7-13c0b6221f88', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.008416Z'),
('5e96109f-9db4-4924-876c-78791a08df8e', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.008576Z'),
('f9344b50-3a97-4db9-b691-85181706c3df', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.008752Z'),
('f225288c-7d24-4cc4-93c2-b2c20633a519', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.008929Z'),
('d14a2432-6a04-4efa-b050-59a826361052', '1dc2a261-8b00-4039-8d3f-f48666e97729', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009105Z'),
('1ef17f38-fd36-4801-a6ba-ffbe626c33ca', '08735594-14f0-4594-aa38-fb0805968918', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009283Z'),
('08320265-9347-4a80-9331-c910419ca3d8', '871485de-0116-4f4a-afcd-30f3f8ec02c5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009431Z'),
('2c03766d-9482-4090-a44c-b4c9745adeeb', '4e06209a-9e0b-4430-87fe-a7ac49d92160', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009606Z'),
('d15e224c-bd35-4411-b464-672db6740928', 'a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009749Z'),
('655f735a-c093-4198-88c3-9b1a538e967c', '3b260872-3454-473b-bfb7-087ed0e809b0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.009941Z'),
('9604b5fd-3d0f-44a4-90c3-7e26327bc851', '606dcda7-8287-4435-9d66-deb4c03b4bf6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.010146Z'),
('f6e0cd34-6a85-4287-85f1-26e96cc13b8a', '0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.010325Z'),
('f2997b21-0afd-4aa8-9a70-4b88af72a31b', '2673f0e2-16a4-443a-ac86-7c337d960f4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.010683Z'),
('91ecc2bb-396a-4ffb-af47-b64e2364441a', 'c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.010861Z'),
('92d3e157-5c88-4ce3-acf8-5dae729f8dcb', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011037Z'),
('2fd906a9-0581-4327-a3e2-f4bad553cfbb', 'ef991379-a625-4de7-bc6a-2902fea1bc79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011210Z'),
('467d2d13-4fb1-42b8-b575-506908cb6b7a', '0bac5d43-fab7-48b3-9fbc-33c5d565c1a9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011210Z'),
('f6ebd1a0-57e0-4df2-b6eb-df01873cbc0b', '90664d51-0503-45cf-85bb-5d46fa579d0f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011449Z'),
('8eece4cf-8d6d-43c4-87c5-6c46f89a2f29', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011619Z'),
('37a318f2-df3c-459e-bcff-ed8860fc447b', 'f063108f-9c3b-49fd-b981-1be021278ebf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011794Z'),
('41f666b6-2577-4ba8-b650-5ffe81ce18f3', '15614ca3-c115-4cff-a021-c43ef352fda8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.011929Z'),
('59a60947-9665-4205-be95-ef048cba1ca0', '31be5853-f839-4cc2-80cb-071210651059', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.012131Z'),
('123ae902-ddee-47b6-9756-4e1d2349559a', '8df04db9-392f-4216-9776-66defd9d8d32', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.012308Z'),
('b6b89950-bc09-44c9-88db-06bb0978a524', 'a6e1707c-08f9-4484-ba65-cd054e9561bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.012476Z'),
('d41d15b4-478b-4cf1-8cbc-1f43d8dc2f76', '4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.012668Z'),
('63f6a65d-b9c6-42c8-abfe-269b822ae6e7', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.012919Z'),
('0e6a533e-d2cd-4038-8204-a28da281f6ef', 'c559d145-3351-4e6e-bf8d-df05f69814ea', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.013093Z'),
('40671ca3-2447-49c1-8890-89e21ece9a83', 'b695967d-c050-4428-ab1e-db9407fe9d2f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T13:03:50.013291Z');

--INSERT TELEMETRY_GOES--COUNT:28
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8f87d488-5967-430f-8017-6b0ae982d4a6', 'CE7EB098' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EB098');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '57193da5-b7d8-47bf-95e7-19d9bc14d190', 'CE7EBE4A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EBE4A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '82001ff9-0ecb-434c-817d-d08da056a70c', 'CE7EC608' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC608');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd4fe88bb-9f58-4e15-b031-98b4b557cf0a', 'CE7EC8DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC8DA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '47ba9379-b660-437f-ab2d-efaa8f901b4b', 'CE7ED57E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7ED57E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'fdbaf4fb-3056-4e3c-95af-74281c834f89', 'CE7EDBAC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EDBAC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6731e215-1222-4bae-843d-91a966d08d8b', '1715F330' where not exists (select 1 from telemetry_goes where nesdis_id = '1715F330');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '01eb2ec1-8f33-42d5-9067-b8ce3aee9c49', '1716615C' where not exists (select 1 from telemetry_goes where nesdis_id = '1716615C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a39fc95b-3844-49f4-a95a-070449b71fd8', 'DD8362E0' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8362E0');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '067049ba-7f1a-42f7-9de7-216251bdaf3f', '172024E8' where not exists (select 1 from telemetry_goes where nesdis_id = '172024E8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a23a7f32-3c46-4f4b-a62d-cd35189f8bd7', 'CE5D2DC8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D2DC8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b5944040-5f19-4863-8979-02599421ed20', 'CE6B8B8E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B8B8E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9527ccff-cc1f-4dd6-a3f9-f54df90c9582', 'DDA2D27A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA2D27A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6587a1af-a687-430e-b204-daf2d15ea644', 'CE216420' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE216420');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '418b83f4-9ef1-408e-bc5b-0b6e5b9bf764', 'DD6E777A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD6E777A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6d28db6c-9c70-4d3d-bf6b-b44856eef759', 'CE7EE0E4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EE0E4');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6289f345-ed03-4c33-bf51-3b53073eea5c', 'CE1FF35A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE1FF35A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '92ff6255-1ae7-4e30-855f-bd594dd53f75', 'CE7EEE36' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EEE36');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0d7008c3-2c8c-4b81-8c3a-1206806847c1', 'CE7EFD40' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EFD40');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'dd808d2f-3d28-4234-ad30-450e22470a0b', 'CE7EF392' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EF392');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8119cd62-0586-40c3-86ba-d45855c5e7e0', '33655136' where not exists (select 1 from telemetry_goes where nesdis_id = '33655136');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '916b6e5e-f3ae-4897-bd18-6bc87e33832d', 'CE5D231A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D231A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'fb0bc190-ff2e-45e3-a6e8-6d9af0a0d316', 'CE19401A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE19401A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '39a2db16-f654-4956-a2f8-d06f9158fb71', 'CE59652A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE59652A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'efb54edb-10e2-420f-a1fa-d1043c5b4fed', 'DF02AF2A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF02AF2A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ce2c5095-7c04-4629-8838-037481a30f88', 'DD66C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD66C052');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '78c30182-76b1-45f4-a140-2c2059930cf5', 'DF055D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF055D9A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '326c34db-c19c-416d-b21c-cf61fa9ab24c', 'CE437E42' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE437E42');

--INSERT INSTRUMENT_TELEMETRY--COUNT:28
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('089fbda3-5a6a-408d-aec2-4e108910d94b', '10a32652-af43-4451-bd52-4980c5690cc9', '8f87d488-5967-430f-8017-6b0ae982d4a6'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', '10a32652-af43-4451-bd52-4980c5690cc9', '57193da5-b7d8-47bf-95e7-19d9bc14d190'),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '10a32652-af43-4451-bd52-4980c5690cc9', '82001ff9-0ecb-434c-817d-d08da056a70c'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '10a32652-af43-4451-bd52-4980c5690cc9', 'd4fe88bb-9f58-4e15-b031-98b4b557cf0a'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', '10a32652-af43-4451-bd52-4980c5690cc9', '47ba9379-b660-437f-ab2d-efaa8f901b4b'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', '10a32652-af43-4451-bd52-4980c5690cc9', 'fdbaf4fb-3056-4e3c-95af-74281c834f89'),
('08735594-14f0-4594-aa38-fb0805968918', '10a32652-af43-4451-bd52-4980c5690cc9', '6731e215-1222-4bae-843d-91a966d08d8b'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', '10a32652-af43-4451-bd52-4980c5690cc9', '01eb2ec1-8f33-42d5-9067-b8ce3aee9c49'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', '10a32652-af43-4451-bd52-4980c5690cc9', 'a39fc95b-3844-49f4-a95a-070449b71fd8'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '10a32652-af43-4451-bd52-4980c5690cc9', '067049ba-7f1a-42f7-9de7-216251bdaf3f'),
('3b260872-3454-473b-bfb7-087ed0e809b0', '10a32652-af43-4451-bd52-4980c5690cc9', 'a23a7f32-3c46-4f4b-a62d-cd35189f8bd7'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', '10a32652-af43-4451-bd52-4980c5690cc9', 'b5944040-5f19-4863-8979-02599421ed20'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', '10a32652-af43-4451-bd52-4980c5690cc9', '9527ccff-cc1f-4dd6-a3f9-f54df90c9582'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', '10a32652-af43-4451-bd52-4980c5690cc9', '6587a1af-a687-430e-b204-daf2d15ea644'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', '10a32652-af43-4451-bd52-4980c5690cc9', '418b83f4-9ef1-408e-bc5b-0b6e5b9bf764'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '10a32652-af43-4451-bd52-4980c5690cc9', '6d28db6c-9c70-4d3d-bf6b-b44856eef759'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', '10a32652-af43-4451-bd52-4980c5690cc9', '6289f345-ed03-4c33-bf51-3b53073eea5c'),
('90664d51-0503-45cf-85bb-5d46fa579d0f', '10a32652-af43-4451-bd52-4980c5690cc9', '92ff6255-1ae7-4e30-855f-bd594dd53f75'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '10a32652-af43-4451-bd52-4980c5690cc9', '0d7008c3-2c8c-4b81-8c3a-1206806847c1'),
('f063108f-9c3b-49fd-b981-1be021278ebf', '10a32652-af43-4451-bd52-4980c5690cc9', 'dd808d2f-3d28-4234-ad30-450e22470a0b'),
('15614ca3-c115-4cff-a021-c43ef352fda8', '10a32652-af43-4451-bd52-4980c5690cc9', '8119cd62-0586-40c3-86ba-d45855c5e7e0'),
('31be5853-f839-4cc2-80cb-071210651059', '10a32652-af43-4451-bd52-4980c5690cc9', '916b6e5e-f3ae-4897-bd18-6bc87e33832d'),
('8df04db9-392f-4216-9776-66defd9d8d32', '10a32652-af43-4451-bd52-4980c5690cc9', 'fb0bc190-ff2e-45e3-a6e8-6d9af0a0d316'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', '10a32652-af43-4451-bd52-4980c5690cc9', '39a2db16-f654-4956-a2f8-d06f9158fb71'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '10a32652-af43-4451-bd52-4980c5690cc9', 'efb54edb-10e2-420f-a1fa-d1043c5b4fed'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '10a32652-af43-4451-bd52-4980c5690cc9', 'ce2c5095-7c04-4629-8838-037481a30f88'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', '10a32652-af43-4451-bd52-4980c5690cc9', '78c30182-76b1-45f4-a140-2c2059930cf5'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', '10a32652-af43-4451-bd52-4980c5690cc9', '326c34db-c19c-416d-b21c-cf61fa9ab24c');

--INSERT TIMESERIES--COUNT:28
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('d64b6f89-35d3-41ed-b7a8-5f6b1e442b18','stage','Stage','089fbda3-5a6a-408d-aec2-4e108910d94b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('70f10f0c-60ce-41be-94bf-4e8af67e5d61','precipitation','Precipitation','089fbda3-5a6a-408d-aec2-4e108910d94b', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('8ffb6eb8-ce23-469d-8afc-68481850bf4f','voltage','Voltage','089fbda3-5a6a-408d-aec2-4e108910d94b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8bb4bbc1-6691-453b-86c4-1e3f702bd6d3','stage','Stage','209980f3-ab26-42d6-9fe7-13c0b6221f88', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ac3e3607-294b-4446-9d09-b0c42f57f5f0','precipitation','Precipitation','209980f3-ab26-42d6-9fe7-13c0b6221f88', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('894b5bf7-0542-4218-a4c3-770151d60760','voltage','Voltage','209980f3-ab26-42d6-9fe7-13c0b6221f88', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cb84744f-add7-41f8-befb-a28d3889007b','stage','Stage','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1b31efbf-18b5-4395-af72-eeec2940ec4a','precipitation','Precipitation','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c3d94b72-91eb-48c1-9835-e4bb1b96e250','voltage','Voltage','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('da574818-07c4-41ec-8946-2ce2afa3dffa','stage','Stage','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('de4558a9-a2dc-44bc-936f-d992922a2608','precipitation','Precipitation','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6e3455fb-c1f1-495d-b6eb-88efb93210a3','voltage','Voltage','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cf42fece-3a14-4767-8d05-37b65a748c9d','stage','Stage','891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bf2cd19a-e1ee-4f8c-9297-9007938210f8','precipitation','Precipitation','891c7d48-fa1b-4f4a-a7dd-7393258bb575', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('dd43d481-1e82-4834-8358-1c3a6c4398a1','voltage','Voltage','891c7d48-fa1b-4f4a-a7dd-7393258bb575', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('aab1c280-1342-4f95-b8ba-d440b4772b10','stage','Stage','1dc2a261-8b00-4039-8d3f-f48666e97729', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b045f6b5-a09e-4986-aaa6-2092406cd94c','precipitation','Precipitation','1dc2a261-8b00-4039-8d3f-f48666e97729', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('5dafcc36-e894-482e-b850-9bf68d4d5f22','voltage','Voltage','1dc2a261-8b00-4039-8d3f-f48666e97729', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6d610e49-9a76-4a78-bfe8-23a997220d2c','stage','Stage','08735594-14f0-4594-aa38-fb0805968918', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cca5407e-939c-403e-bf2d-acf9efab52e7','voltage','Voltage','08735594-14f0-4594-aa38-fb0805968918', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e46df6e5-d0f5-4286-b2f5-b5c044f7038c','stage','Stage','871485de-0116-4f4a-afcd-30f3f8ec02c5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('62744304-0dd3-4c54-bd49-a696052cbf2e','precipitation','Precipitation','871485de-0116-4f4a-afcd-30f3f8ec02c5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('472ca572-0896-4bfa-9a44-1910171997a5','voltage','Voltage','871485de-0116-4f4a-afcd-30f3f8ec02c5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('383aef78-09d0-421a-b4ec-8fc62d1aec5f','stage','Stage','4e06209a-9e0b-4430-87fe-a7ac49d92160', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('219c7f59-b0c9-4c5f-b6bc-b3a30b07f187','voltage','Voltage','4e06209a-9e0b-4430-87fe-a7ac49d92160', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e2b6283b-1511-4719-96cf-ba2f1aef8bba','stage','Stage','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fa0dbbab-4933-4b57-b243-f38e9efa990b','precipitation','Precipitation','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('8200c349-8d60-4f6b-a730-8c94645b9dce','voltage','Voltage','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1f1d144f-520f-40bf-a93f-0c1e1e5fbe8b','stage','Stage','3b260872-3454-473b-bfb7-087ed0e809b0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5de3bf56-e393-4279-9358-67b9b927e53a','precipitation','Precipitation','3b260872-3454-473b-bfb7-087ed0e809b0', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('3fd27bfd-91f3-42e3-8295-653eb1cb9e98','voltage','Voltage','3b260872-3454-473b-bfb7-087ed0e809b0', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('376312cf-9464-4073-9c4e-6b92b1fa3591','stage','Stage','606dcda7-8287-4435-9d66-deb4c03b4bf6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('51dab6cc-3d2f-40c2-8c7d-74b5f8d0b1cf','precipitation','Precipitation','606dcda7-8287-4435-9d66-deb4c03b4bf6', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('22dfb117-e59e-4b8b-ad0c-a8633f9dab2c','voltage','Voltage','606dcda7-8287-4435-9d66-deb4c03b4bf6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2fcad72d-9b50-4d18-857f-05f2dd34d999','stage','Stage','0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('40a93802-6485-460e-91d3-d118b0836748','conductivity','Conductivity','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '377ecec0-f785-46ab-b0e2-5fd8c682dfea', '633bd96c-5bdb-436f-b464-f18d90b7d736'),
('fe1ff27b-514c-492e-a444-ed14477a9887','ph','ph','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '5d0b2c85-6a4c-4d82-aed3-193b066349f1', '4484c18a-61aa-48b4-8cf5-63d3b8c6d200'),
('8a1a623a-a43d-4b42-ad5c-d35d8cc7847e','turbidity','Turbidity','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '3676df6a-37c2-4a81-9072-ddcd4ab93702', '7d8e5bb9-b9ea-4920-9def-0589160ea412'),
('2fb84b15-9338-4f45-bb69-9ab4699f7ad2','dissolved-oxygen','Dissolved-Oxygen','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '98007857-d027-4524-9a63-d07ae93e5fa2', '67d75ccd-6bf7-4086-a970-5ed65a5c30f3'),
('ed1907de-9438-4ee6-ad84-81fcfe20b257','stage','Stage','2673f0e2-16a4-443a-ac86-7c337d960f4f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6135692a-0974-4eaa-a612-75b257c889ab','precipitation','Precipitation','2673f0e2-16a4-443a-ac86-7c337d960f4f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6396794d-1298-41bb-9738-f25917724e9d','voltage','Voltage','2673f0e2-16a4-443a-ac86-7c337d960f4f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('92dbe666-07c2-4148-a598-74ab93774431','stage','Stage','c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('37d12da7-ebfa-4d5e-9c37-0e3d12f491f4','precipitation','Precipitation','c05efa0c-fa2c-4776-a49b-dd68f70f5104', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6df0b12b-edbf-45af-95de-6547a77b39cf','voltage','Voltage','c05efa0c-fa2c-4776-a49b-dd68f70f5104', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f298d244-ad11-4dc5-8619-aeaa596628d8','stage','Stage','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d3e7c175-c8bf-4368-84bc-35ef64fa6640','precipitation','Precipitation','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d9e45426-6b48-4014-ac3f-ab91d03b8d39','voltage','Voltage','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('14d9136e-528c-4f35-8270-6db15894f307','elevation','Elevation','ef991379-a625-4de7-bc6a-2902fea1bc79', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b7546e1b-0652-47c2-abb9-e4d812087cfd','precipitation','Precipitation','ef991379-a625-4de7-bc6a-2902fea1bc79', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('8773557b-53c5-4113-b7de-d61b4f78fe50','voltage','Voltage','ef991379-a625-4de7-bc6a-2902fea1bc79', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5119578c-2f1a-47cc-9f2e-c1063349fdac','stage','Stage','90664d51-0503-45cf-85bb-5d46fa579d0f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ed4a9313-0c56-4618-9f01-74632df20d09','precipitation','Precipitation','90664d51-0503-45cf-85bb-5d46fa579d0f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('809976c2-46c5-4f9b-834d-1ed6930391e1','voltage','Voltage','90664d51-0503-45cf-85bb-5d46fa579d0f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dc718eb3-297c-4876-9fa5-332e31466f7a','stage','Stage','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fa26f254-9e21-4a5a-ae85-11ee294ecf81','precipitation','Precipitation','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('41f6b51e-c65b-4dff-a240-8bf886ee9b0b','voltage','Voltage','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('95003251-4935-4ef6-baa1-80c0a9af9755','stage','Stage','f063108f-9c3b-49fd-b981-1be021278ebf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('683e2968-014c-442e-8bf4-278f7ba03657','precipitation','Precipitation','f063108f-9c3b-49fd-b981-1be021278ebf', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('182e4042-6d79-4998-89e5-8c15cba34fb2','voltage','Voltage','f063108f-9c3b-49fd-b981-1be021278ebf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('19226124-41a3-448f-9e62-37f4eca4275f','stage','Stage','31be5853-f839-4cc2-80cb-071210651059', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5be216e0-8921-403e-afc0-8a8c80f08bcd','precipitation','Precipitation','31be5853-f839-4cc2-80cb-071210651059', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b588cedd-9964-4330-b77c-8dbb90a12f6d','voltage','Voltage','31be5853-f839-4cc2-80cb-071210651059', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cbeae195-0d22-4f5f-a8e4-867b574e9835','stage','Stage','8df04db9-392f-4216-9776-66defd9d8d32', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('391b5b6a-1c76-409f-9a35-6bf5a646a453','precipitation','Precipitation','8df04db9-392f-4216-9776-66defd9d8d32', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('51802c53-23ca-4eb3-9d19-bf6076cdbb4e','voltage','Voltage','8df04db9-392f-4216-9776-66defd9d8d32', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('534960e5-f961-4b4a-a0ba-2f4768fa4ece','stage','Stage','a6e1707c-08f9-4484-ba65-cd054e9561bd', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ae827c86-1fed-4eea-be5c-4341323514ca','precipitation','Precipitation','a6e1707c-08f9-4484-ba65-cd054e9561bd', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ab8be5fb-f186-4f3d-95de-3fd4cd8c5729','voltage','Voltage','a6e1707c-08f9-4484-ba65-cd054e9561bd', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('240e4efc-9a29-4984-94b0-50d57fa181d2','stage','Stage','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e25c541b-c573-42a2-8045-993f020b950a','voltage','Voltage','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1ab25ef2-00e6-42bb-acbf-e249945f127d','stage','Stage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4b775ca3-8898-47fc-861c-05c956a83c9d','precipitation','Precipitation','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('1375ae73-248b-4475-a6ee-f49c69aab987','voltage','Voltage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('52053f6a-8671-409b-840d-3a3450ef4911','stage','Stage','c559d145-3351-4e6e-bf8d-df05f69814ea', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('20c49b79-1736-4c63-8dce-25fcfef7e4ad','voltage','Voltage','c559d145-3351-4e6e-bf8d-df05f69814ea', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dcbbd66d-f7e2-4cee-92c6-ab34e2c46da4','stage','Stage','b695967d-c050-4428-ab1e-db9407fe9d2f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('eefe2cdb-0457-41bd-aba3-27fe22b3c2e9','precipitation','Precipitation','b695967d-c050-4428-ab1e-db9407fe9d2f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f3b75a66-4cf8-4d3e-b54f-47bc847b0bf1','voltage','Voltage','b695967d-c050-4428-ab1e-db9407fe9d2f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64');
