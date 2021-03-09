-- Project
INSERT INTO project (id, slug, name, image) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'buffalo-district-streamgages', 'Buffalo District Streamgages', 'buffalo-district-streamgages.jpg');





--INSERT INSTRUMENTS--COUNT:29
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('089fbda3-5a6a-408d-aec2-4e108910d94b', False, 'avnn6', 'AVNN6', null, ST_GeomFromText('POINT(-77.7566 42.9184)',4326), null, null, '2021-03-08T19:47:26.196355Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', False, 'blbn6', 'BLBN6', null, ST_GeomFromText('POINT(-77.6806 43.0922)',4326), null, null, '2021-03-08T19:47:26.196768Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', False, 'chcn6', 'CHCN6', null, ST_GeomFromText('POINT(-77.8822 43.1008)',4326), null, null, '2021-03-08T19:47:26.197020Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', False, 'dsvn6', 'DSVN6', null, ST_GeomFromText('POINT(-77.7064 42.5322)',4326), null, null, '2021-03-08T19:47:26.197299Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224775'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', False, 'garn6', 'GARN6', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-08T19:47:26.197569Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', False, 'hnyn6', 'HNYN6', null, ST_GeomFromText('POINT(-77.5869 42.9567)',4326), null, null, '2021-03-08T19:47:26.197803Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04229500'),
('08735594-14f0-4594-aa38-fb0805968918', False, 'blackcr-churchvl', 'BlackCr Churchvl', null, ST_GeomFromText('POINT(-77.8822 43.1006)',4326), null, null, '2021-03-08T19:47:26.198040Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', False, 'genr-portagevill', 'GenR Portagevill', null, ST_GeomFromText('POINT(-78.0422 42.5703)',4326), null, null, '2021-03-08T19:47:26.198232Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', False, 'oatkacr-garbutt', 'OatkaCr Garbutt', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-08T19:47:26.198414Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', False, 'knvn6', 'KNVN6', null, ST_GeomFromText('POINT(-78.3103 43.3011)',4326), null, null, '2021-03-08T19:47:26.198603Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '0422016550'),
('3b260872-3454-473b-bfb7-087ed0e809b0', False, 'mbyp1', 'MBYP1', null, ST_GeomFromText('POINT(-77.2736 41.8425)',4326), null, null, '2021-03-08T19:47:26.199223Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01518420'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', False, 'olnn6', 'OLNN6', null, ST_GeomFromText('POINT(-78.4511 42.0731)',4326), null, null, '2021-03-08T19:47:26.199658Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010820'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', False, 'rohn6', 'ROHN6', null, ST_GeomFromText('POINT(-77.6163 43.1417)',4326), null, null, '2021-03-08T19:47:26.200101Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231600'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', False, 'shnp1', 'SHNP1', null, ST_GeomFromText('POINT(-78.1983 41.9617)',4326), null, null, '2021-03-08T19:47:26.201010Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010655'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', False, 'genr-wellsville', 'GenR Wellsville', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-08T19:47:26.201443Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', False, 'jonn6', 'JONN6', null, ST_GeomFromText('POINT(-77.8386 42.7667)',4326), null, null, '2021-03-08T19:47:26.201804Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227500'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', False, 'mount-morris', 'Mount Morris', null, ST_GeomFromText('POINT(-77.9071 42.7333)',4326), null, null, '2021-03-08T19:47:26.202193Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224000'),
('37989e13-f192-4d1f-aa24-62591f8d731d', False, 'mount-morris-tailwater', 'Mount Morris-Tailwater', null, ST_GeomFromText('POINT(-77.9109 42.7332)',4326), null, null, '2021-03-08T19:47:26.202193Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('90664d51-0503-45cf-85bb-5d46fa579d0f', False, 'ptgn6', 'PTGN6', null, ST_GeomFromText('POINT(-78.0431 42.5697)',4326), null, null, '2021-03-08T19:47:26.202615Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', False, 'weln6', 'WELN6', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-08T19:47:26.202981Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('f063108f-9c3b-49fd-b981-1be021278ebf', False, 'wrsn6', 'WRSN6', null, ST_GeomFromText('POINT(-78.1375 42.7447)',4326), null, null, '2021-03-08T19:47:26.203327Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('15614ca3-c115-4cff-a021-c43ef352fda8', False, 'rcrn6', 'RCRN6', null, ST_GeomFromText('POINT(-77.6025 43.258)',4326), null, null, '2021-03-08T19:47:26.203504Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('31be5853-f839-4cc2-80cb-071210651059', False, 'elkp1', 'ELKP1', null, ST_GeomFromText('POINT(-77.3025 41.9875)',4326), null, null, '2021-03-08T19:47:26.204127Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01519200'),
('8df04db9-392f-4216-9776-66defd9d8d32', False, 'frkn6', 'FRKN6', null, ST_GeomFromText('POINT(-78.4636 42.3294)',4326), null, null, '2021-03-08T19:47:26.204499Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '421946078274901'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', False, 'hrln6', 'HRLN6', null, ST_GeomFromText('POINT(-77.7044 42.3489)',4326), null, null, '2021-03-08T19:47:26.204890Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01523000'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', False, 'canaseragashaker', 'CanaseragaShaker', null, ST_GeomFromText('POINT(-77.8414 42.7361)',4326), null, null, '2021-03-08T19:47:26.205554Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227000'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', False, 'oatkacr-warsaw', 'OatkaCr Warsaw', null, ST_GeomFromText('POINT(-78.1375 42.7442)',4326), null, null, '2021-03-08T19:47:26.206187Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230380'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', False, 'genr-avon', 'GenR Avon', null, ST_GeomFromText('POINT(-77.7572 42.9178)',4326), null, null, '2021-03-08T19:47:26.206532Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', False, 'akln6', 'AKLN6', null, ST_GeomFromText('POINT(-77.7167 42.3958)',4326), null, null, '2021-03-08T19:47:26.207036Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01521000');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('93ee043c-5776-4df9-afb0-0517ea00158c', '089fbda3-5a6a-408d-aec2-4e108910d94b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.196355Z'),
('0964a1fd-18ba-46a0-bc7c-f60fd9ba45ea', '209980f3-ab26-42d6-9fe7-13c0b6221f88', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.196768Z'),
('5f902dff-f198-4e9d-9328-7330bb5bf1bf', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.197020Z'),
('bb761417-2763-43f6-aac8-33fedc0c8d0b', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.197299Z'),
('9c27922f-5f13-4ee4-8fa8-01726dcfbb43', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.197569Z'),
('e764e0da-6e35-4582-b811-d288d7d87b1e', '1dc2a261-8b00-4039-8d3f-f48666e97729', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.197803Z'),
('20f3b18c-0993-4701-b3d9-6b01c1f0b2aa', '08735594-14f0-4594-aa38-fb0805968918', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.198040Z'),
('be8055bf-6951-448c-a4c8-07d715d95871', '871485de-0116-4f4a-afcd-30f3f8ec02c5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.198232Z'),
('739fa2fc-936f-49de-aad7-3f9e8cd6d328', '4e06209a-9e0b-4430-87fe-a7ac49d92160', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.198414Z'),
('214c0d6b-2806-401d-94a1-fc2306461e76', 'a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.198603Z'),
('82626743-83a4-4d71-8589-4a6ae0a71ba1', '3b260872-3454-473b-bfb7-087ed0e809b0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.199223Z'),
('fa41467f-55f3-4340-a847-5553c18f6c18', '606dcda7-8287-4435-9d66-deb4c03b4bf6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.199658Z'),
('92b1a2ca-b6a6-42ea-893e-dac7cdb4e7eb', '0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.200101Z'),
('d9710cde-196f-4dd0-9c90-069d094fc3be', '2673f0e2-16a4-443a-ac86-7c337d960f4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.201010Z'),
('3802d399-dbda-4448-971a-055f15999087', 'c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.201443Z'),
('41ca9ad3-efa7-4d4d-8f36-11af11b7e9b7', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.201804Z'),
('ddca5583-893e-471f-9138-381dd01dc854', 'ef991379-a625-4de7-bc6a-2902fea1bc79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.202193Z'),
('972bda71-cbf7-4ce9-886c-d587cec0a0e1', '37989e13-f192-4d1f-aa24-62591f8d731d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.202193Z'),
('fb52d05f-05b9-4b6e-918d-7a75574f4dcc', '90664d51-0503-45cf-85bb-5d46fa579d0f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.202615Z'),
('abb4d6e9-75f1-4206-9f6d-6297631ad90b', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.202981Z'),
('eb6c1d5d-2035-4bd1-870b-945cc90dd14c', 'f063108f-9c3b-49fd-b981-1be021278ebf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.203327Z'),
('9767a221-5eda-4872-9ed4-46a6927bb359', '15614ca3-c115-4cff-a021-c43ef352fda8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.203504Z'),
('478f5e5b-2cca-4a49-ba41-2ea71041a6ac', '31be5853-f839-4cc2-80cb-071210651059', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.204127Z'),
('9051cfea-e8ef-403a-bdaa-de297ad3cedd', '8df04db9-392f-4216-9776-66defd9d8d32', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.204499Z'),
('e12f4c3f-5251-4106-ae22-9a18a828c922', 'a6e1707c-08f9-4484-ba65-cd054e9561bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.204890Z'),
('872f6f98-7f60-43a3-847e-71d756b9cbab', '4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.205554Z'),
('609a7fe6-cf5c-404c-8342-e0f806f2f526', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.206187Z'),
('b315088f-c92e-4618-ab80-e19f740bbcbd', 'c559d145-3351-4e6e-bf8d-df05f69814ea', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.206532Z'),
('01e32708-b9bc-4cda-8f62-e10d33ca8e54', 'b695967d-c050-4428-ab1e-db9407fe9d2f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:47:26.207036Z');

--INSERT TELEMETRY_GOES--COUNT:28
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e1fb5518-ae6d-4345-ad3d-b90a08274dae', 'CE7EB098' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EB098');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '120c0a3e-b690-45a5-a5c1-bd4a64576d29', 'CE7EBE4A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EBE4A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '76480260-7ec1-4293-b94c-667b0a1471fc', 'CE7EC608' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC608');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '58e3f6d5-686c-4779-bf26-c8a9da5dc12c', 'CE7EC8DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC8DA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd70f905e-85da-454f-a013-43b865d53be4', 'CE7ED57E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7ED57E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f01d0833-7c16-4eb7-b541-9815685c9750', 'CE7EDBAC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EDBAC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f1fa3bca-5b88-4ca4-9bf0-d2661b527d90', '1715F330' where not exists (select 1 from telemetry_goes where nesdis_id = '1715F330');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0e981015-a770-4cbb-8634-47ea6317aa38', '1716615C' where not exists (select 1 from telemetry_goes where nesdis_id = '1716615C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '4fe4918d-3ac5-4ebc-a58a-61e52acaea97', 'DD8362E0' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8362E0');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0d44a5eb-6ebc-445f-9b13-602b90dd63b4', '172024E8' where not exists (select 1 from telemetry_goes where nesdis_id = '172024E8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9ced5af8-f85d-4133-ae9c-2b6266607914', 'CE5D2DC8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D2DC8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b2bb5188-aed0-45f0-b0e6-be3ea0db4a74', 'CE6B8B8E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B8B8E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'df47b8af-393a-46cd-81d9-df74f8a3b730', 'DDA2D27A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA2D27A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a7328a09-e7dc-4325-9db0-4b7e8813f1fe', 'CE216420' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE216420');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6667c519-a5e2-4767-b54e-e25955bded94', 'DD6E777A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD6E777A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2060d810-28a7-4983-b0c8-f74354883838', 'CE7EE0E4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EE0E4');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f6c28bed-5ec8-4735-aff5-c1bbb2c8b0de', 'CE1FF35A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE1FF35A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c4538c36-1b27-4ac3-8f99-190348cc76df', 'CE7EEE36' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EEE36');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1f9d87bc-adac-4187-a2d3-dabb14d60e39', 'CE7EFD40' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EFD40');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'dcadadd9-3057-4624-820b-034aee6403cc', 'CE7EF392' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EF392');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '292b8d1e-89aa-4240-915b-80481197dd82', '33655136' where not exists (select 1 from telemetry_goes where nesdis_id = '33655136');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2b664259-6342-4a7a-a07c-49f6fe4f0506', 'CE5D231A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D231A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1ea10049-a422-479a-bc57-006362e9cd94', 'CE19401A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE19401A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '67abb326-0229-42f3-bf47-ffd346891840', 'CE59652A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE59652A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '418bb7dd-f63a-4123-a307-3c4bfc76822d', 'DF02AF2A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF02AF2A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a0991b73-a866-4582-8d25-348cda3489b9', 'DD66C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD66C052');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd9680838-def3-4e47-8bd8-3be640c44474', 'DF055D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF055D9A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '00339057-ea1c-43e3-93da-057cbf235de5', 'CE437E42' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE437E42');

--INSERT INSTRUMENT_TELEMETRY--COUNT:28
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('089fbda3-5a6a-408d-aec2-4e108910d94b', '10a32652-af43-4451-bd52-4980c5690cc9', 'e1fb5518-ae6d-4345-ad3d-b90a08274dae'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', '10a32652-af43-4451-bd52-4980c5690cc9', '120c0a3e-b690-45a5-a5c1-bd4a64576d29'),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '10a32652-af43-4451-bd52-4980c5690cc9', '76480260-7ec1-4293-b94c-667b0a1471fc'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '10a32652-af43-4451-bd52-4980c5690cc9', '58e3f6d5-686c-4779-bf26-c8a9da5dc12c'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', '10a32652-af43-4451-bd52-4980c5690cc9', 'd70f905e-85da-454f-a013-43b865d53be4'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', '10a32652-af43-4451-bd52-4980c5690cc9', 'f01d0833-7c16-4eb7-b541-9815685c9750'),
('08735594-14f0-4594-aa38-fb0805968918', '10a32652-af43-4451-bd52-4980c5690cc9', 'f1fa3bca-5b88-4ca4-9bf0-d2661b527d90'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', '10a32652-af43-4451-bd52-4980c5690cc9', '0e981015-a770-4cbb-8634-47ea6317aa38'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', '10a32652-af43-4451-bd52-4980c5690cc9', '4fe4918d-3ac5-4ebc-a58a-61e52acaea97'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '10a32652-af43-4451-bd52-4980c5690cc9', '0d44a5eb-6ebc-445f-9b13-602b90dd63b4'),
('3b260872-3454-473b-bfb7-087ed0e809b0', '10a32652-af43-4451-bd52-4980c5690cc9', '9ced5af8-f85d-4133-ae9c-2b6266607914'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', '10a32652-af43-4451-bd52-4980c5690cc9', 'b2bb5188-aed0-45f0-b0e6-be3ea0db4a74'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', '10a32652-af43-4451-bd52-4980c5690cc9', 'df47b8af-393a-46cd-81d9-df74f8a3b730'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', '10a32652-af43-4451-bd52-4980c5690cc9', 'a7328a09-e7dc-4325-9db0-4b7e8813f1fe'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', '10a32652-af43-4451-bd52-4980c5690cc9', '6667c519-a5e2-4767-b54e-e25955bded94'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '10a32652-af43-4451-bd52-4980c5690cc9', '2060d810-28a7-4983-b0c8-f74354883838'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', '10a32652-af43-4451-bd52-4980c5690cc9', 'f6c28bed-5ec8-4735-aff5-c1bbb2c8b0de'),
('90664d51-0503-45cf-85bb-5d46fa579d0f', '10a32652-af43-4451-bd52-4980c5690cc9', 'c4538c36-1b27-4ac3-8f99-190348cc76df'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '10a32652-af43-4451-bd52-4980c5690cc9', '1f9d87bc-adac-4187-a2d3-dabb14d60e39'),
('f063108f-9c3b-49fd-b981-1be021278ebf', '10a32652-af43-4451-bd52-4980c5690cc9', 'dcadadd9-3057-4624-820b-034aee6403cc'),
('15614ca3-c115-4cff-a021-c43ef352fda8', '10a32652-af43-4451-bd52-4980c5690cc9', '292b8d1e-89aa-4240-915b-80481197dd82'),
('31be5853-f839-4cc2-80cb-071210651059', '10a32652-af43-4451-bd52-4980c5690cc9', '2b664259-6342-4a7a-a07c-49f6fe4f0506'),
('8df04db9-392f-4216-9776-66defd9d8d32', '10a32652-af43-4451-bd52-4980c5690cc9', '1ea10049-a422-479a-bc57-006362e9cd94'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', '10a32652-af43-4451-bd52-4980c5690cc9', '67abb326-0229-42f3-bf47-ffd346891840'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '10a32652-af43-4451-bd52-4980c5690cc9', '418bb7dd-f63a-4123-a307-3c4bfc76822d'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '10a32652-af43-4451-bd52-4980c5690cc9', 'a0991b73-a866-4582-8d25-348cda3489b9'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', '10a32652-af43-4451-bd52-4980c5690cc9', 'd9680838-def3-4e47-8bd8-3be640c44474'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', '10a32652-af43-4451-bd52-4980c5690cc9', '00339057-ea1c-43e3-93da-057cbf235de5');

--INSERT TIMESERIES--COUNT:28
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('0ba3e490-f6dd-4843-980f-a054b6d7484b','stage','Stage','089fbda3-5a6a-408d-aec2-4e108910d94b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('83d56dd6-0618-4b6e-b42c-154eb57c0192','precipitation','Precipitation','089fbda3-5a6a-408d-aec2-4e108910d94b', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('da5d6bd9-1d60-472e-9fbc-a3808fe0d52f','voltage','Voltage','089fbda3-5a6a-408d-aec2-4e108910d94b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4aac3294-ab34-4b03-8f57-615045901811','stage','Stage','209980f3-ab26-42d6-9fe7-13c0b6221f88', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0ad46a96-5bc6-4c03-b242-66b7413b2fa2','precipitation','Precipitation','209980f3-ab26-42d6-9fe7-13c0b6221f88', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('90182a05-126b-4b52-a154-3ca0bbd04715','voltage','Voltage','209980f3-ab26-42d6-9fe7-13c0b6221f88', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('633381a2-7ced-4329-bf3d-98b2853d1756','stage','Stage','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('059905c0-07c6-40fa-be05-87ab2cf8d6a3','precipitation','Precipitation','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('4fa40edb-2462-4f48-aa96-a0251f6c4831','voltage','Voltage','8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c004968d-d710-4e50-a149-a35b6b016827','stage','Stage','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2c4c58fc-18a1-493b-8cf9-2a5ecde97981','precipitation','Precipitation','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('0dc4bedf-3c75-4811-8f45-cdd28827989f','voltage','Voltage','b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('194224c2-30c6-430c-9f76-2e0179c8f927','stage','Stage','891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('34173142-6250-47bf-ab88-e21572870afb','precipitation','Precipitation','891c7d48-fa1b-4f4a-a7dd-7393258bb575', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a7cdc419-8d9b-4e98-90b1-16662010f464','voltage','Voltage','891c7d48-fa1b-4f4a-a7dd-7393258bb575', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3b111eff-a2bd-4091-b44e-741f44169093','stage','Stage','1dc2a261-8b00-4039-8d3f-f48666e97729', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e4364c76-475b-4383-8259-da0f7476285f','precipitation','Precipitation','1dc2a261-8b00-4039-8d3f-f48666e97729', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('69c10676-30af-493d-bf8a-e9e36d2e932f','voltage','Voltage','1dc2a261-8b00-4039-8d3f-f48666e97729', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f7702d1d-c3c4-42f0-8a20-4117f53d5c1c','stage','Stage','08735594-14f0-4594-aa38-fb0805968918', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('42bc1b6f-34b4-44d8-a091-12306a03c1d0','voltage','Voltage','08735594-14f0-4594-aa38-fb0805968918', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ba6987f4-69e7-4222-bd2f-23ac7efbd98f','stage','Stage','871485de-0116-4f4a-afcd-30f3f8ec02c5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1c037111-0162-4142-815b-c499eadb1049','precipitation','Precipitation','871485de-0116-4f4a-afcd-30f3f8ec02c5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e9afbf86-70c1-457c-aefa-100c8bd782da','voltage','Voltage','871485de-0116-4f4a-afcd-30f3f8ec02c5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f62378c6-89bb-41a0-a8ac-21872e938715','stage','Stage','4e06209a-9e0b-4430-87fe-a7ac49d92160', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('557007fd-d43a-4202-82a0-dab8f9db27e4','voltage','Voltage','4e06209a-9e0b-4430-87fe-a7ac49d92160', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2af24b01-5a73-40d0-adc1-5d8165f918f6','stage','Stage','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('450b34a8-f998-4d9f-803b-375f3535485b','precipitation','Precipitation','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('74d1db87-ec08-4956-b1e5-7149c81d6871','voltage','Voltage','a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('80849133-26cb-4e17-b551-6a19023d360a','stage','Stage','3b260872-3454-473b-bfb7-087ed0e809b0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0899efd1-9aeb-40f6-be40-5be70161eb60','precipitation','Precipitation','3b260872-3454-473b-bfb7-087ed0e809b0', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('3aaee076-f0a8-4890-ba27-74e876d31c91','voltage','Voltage','3b260872-3454-473b-bfb7-087ed0e809b0', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2578e054-ff69-4175-ac3b-9cc192f5bbdb','stage','Stage','606dcda7-8287-4435-9d66-deb4c03b4bf6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('318dafc8-087b-439d-bd54-ffcd77768b13','precipitation','Precipitation','606dcda7-8287-4435-9d66-deb4c03b4bf6', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('03177a36-2279-430e-9fc7-3f8614c64f14','voltage','Voltage','606dcda7-8287-4435-9d66-deb4c03b4bf6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c5dc016f-9692-4271-88b9-01db31a4c453','stage','Stage','0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ae246707-427c-4fe9-8336-64f5abb7a3c6','conductivity','Conductivity','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '377ecec0-f785-46ab-b0e2-5fd8c682dfea', '633bd96c-5bdb-436f-b464-f18d90b7d736'),
('775a8c1a-5b06-47ec-8c0d-3ca987df817f','ph','ph','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '5d0b2c85-6a4c-4d82-aed3-193b066349f1', '4484c18a-61aa-48b4-8cf5-63d3b8c6d200'),
('e55a5bf7-ac62-43dd-ae28-478f1938a883','turbidity','Turbidity','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '3676df6a-37c2-4a81-9072-ddcd4ab93702', '7d8e5bb9-b9ea-4920-9def-0589160ea412'),
('767ceda9-2dc9-488f-a8c2-be636e6d22bf','dissolved-oxygen','Dissolved-Oxygen','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '98007857-d027-4524-9a63-d07ae93e5fa2', '67d75ccd-6bf7-4086-a970-5ed65a5c30f3'),
('9a112fa8-e716-4568-9f9c-2d6a09c83d24','stage','Stage','2673f0e2-16a4-443a-ac86-7c337d960f4f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('24da98ef-c332-4e78-b483-ad5016db506d','precipitation','Precipitation','2673f0e2-16a4-443a-ac86-7c337d960f4f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('95a80bbb-3972-4999-b493-ab53f38b3d17','voltage','Voltage','2673f0e2-16a4-443a-ac86-7c337d960f4f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('32cfc7f6-bf5f-4c5b-84b6-6e4141c343f2','stage','Stage','c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5b1a7d09-6012-44e4-b686-eb7894ee23de','precipitation','Precipitation','c05efa0c-fa2c-4776-a49b-dd68f70f5104', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d9eaca4a-78a0-4138-8fd9-930fa88c1d5c','voltage','Voltage','c05efa0c-fa2c-4776-a49b-dd68f70f5104', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a1a3d119-0d64-446f-be95-2a8c0bc0c7a3','stage','Stage','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('401529b3-8d16-4181-949b-c4c694216bd7','precipitation','Precipitation','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ac74f5c9-8e20-402a-80fc-47b5aa8e862d','voltage','Voltage','b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('62adf1cb-f342-48f0-ad71-60abe7d2f850','elevation','Elevation','ef991379-a625-4de7-bc6a-2902fea1bc79', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9ecb18a4-41e5-4160-ae98-9914a007c424','precipitation','Precipitation','ef991379-a625-4de7-bc6a-2902fea1bc79', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c9d050b8-db13-4a17-9400-b79f477aef1e','voltage','Voltage','ef991379-a625-4de7-bc6a-2902fea1bc79', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('603e0176-c729-42bd-b394-ecf50d7cf458','stage','Stage','90664d51-0503-45cf-85bb-5d46fa579d0f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('24410a4a-2313-4544-8320-5f9ca08f4915','precipitation','Precipitation','90664d51-0503-45cf-85bb-5d46fa579d0f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f25273ce-b87b-441e-b886-268c1312ea06','voltage','Voltage','90664d51-0503-45cf-85bb-5d46fa579d0f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8b6b0f67-a918-456b-9950-ad4336d7570c','stage','Stage','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3f0e8645-1a34-4575-beb1-df2f74a209a4','precipitation','Precipitation','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e57970dd-b81b-4471-8be6-3e1cafc183ea','voltage','Voltage','2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a6af065e-7191-4738-95c9-4d700617cac2','stage','Stage','f063108f-9c3b-49fd-b981-1be021278ebf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c8e7ec54-168a-400b-853e-4b360056ceeb','precipitation','Precipitation','f063108f-9c3b-49fd-b981-1be021278ebf', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9b0b0bba-43e3-45be-a682-3509a74e8f4b','voltage','Voltage','f063108f-9c3b-49fd-b981-1be021278ebf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4c7f0765-c442-4bf3-9d40-51d2213cb6a6','stage','Stage','31be5853-f839-4cc2-80cb-071210651059', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('72905be3-58d0-40d1-b460-f137cd4c0f61','precipitation','Precipitation','31be5853-f839-4cc2-80cb-071210651059', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('697dbc1f-eb48-46eb-9c07-7f15418f6f3e','voltage','Voltage','31be5853-f839-4cc2-80cb-071210651059', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('71c5005d-ddd7-4ffc-bc20-57501ac7f51e','stage','Stage','8df04db9-392f-4216-9776-66defd9d8d32', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('570b8c34-3eed-4ede-b2a5-ba5eea9848ad','precipitation','Precipitation','8df04db9-392f-4216-9776-66defd9d8d32', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('86c24b27-b491-4bd3-823e-f6afc1c79639','voltage','Voltage','8df04db9-392f-4216-9776-66defd9d8d32', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('9e83e2c2-655c-4de0-80a9-f116128292cb','stage','Stage','a6e1707c-08f9-4484-ba65-cd054e9561bd', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fdec0e13-ef18-44dc-aa1b-3ea35c7e3a24','precipitation','Precipitation','a6e1707c-08f9-4484-ba65-cd054e9561bd', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('add5b1aa-35a4-44d5-8b02-ac88f4489ffe','voltage','Voltage','a6e1707c-08f9-4484-ba65-cd054e9561bd', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('984d8f65-e3c6-436d-8068-7304c6cb4f9d','stage','Stage','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cbf7c063-dec1-417b-abde-56302cf5ffac','voltage','Voltage','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dfdee47b-f6d3-4c79-979a-115a7dc8aa90','stage','Stage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1fe6104e-a3b2-4683-ad37-9e7c86aa4b84','precipitation','Precipitation','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('011270a6-76d7-4e20-941c-2dc6d8d18da4','voltage','Voltage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2aaf99a5-00d7-409c-b5a9-3abeea28791a','stage','Stage','c559d145-3351-4e6e-bf8d-df05f69814ea', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8de721a8-4599-4b7f-9ddd-a4916234c9be','voltage','Voltage','c559d145-3351-4e6e-bf8d-df05f69814ea', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8555f21e-8847-4117-aebc-55f91cd48c29','stage','Stage','b695967d-c050-4428-ab1e-db9407fe9d2f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cc8aba30-7ee8-4cc2-91b0-95c91c2a919c','precipitation','Precipitation','b695967d-c050-4428-ab1e-db9407fe9d2f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c83c55ec-21b4-4b62-8440-b12c130dcfad','voltage','Voltage','b695967d-c050-4428-ab1e-db9407fe9d2f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64');
