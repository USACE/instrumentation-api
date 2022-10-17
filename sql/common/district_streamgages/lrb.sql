-- Project
INSERT INTO project (id, slug, name, image) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'buffalo-district-streamgages', 
    'Buffalo District Streamgages', 'buffalo-district-streamgages.jpg');

--#################################################
--DELETE existing data (mostly for dev, test, prod)
--#################################################

-- Delete Timeseries Measurements
delete from timeseries_measurement where timeseries_id in (
	select t.id from instrument i 
	join timeseries t on i.id = t.instrument_id
	where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980');

-- Delete Timeseries
delete from timeseries where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980');
	
-- Delete Telemetry GOES
delete from telemetry_goes where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='10a32652-af43-4451-bd52-4980c5690cc9'
	and instrument_id in (
		select i.id from instrument i
		where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980')
	);

-- Delete Telemetry Iridium
delete from telemetry_iridium where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='c0b03b0d-bfce-453a-b5a9-636118940449'
	and	instrument_id in (
		select i.id from instrument i
		where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980')
	);
	
-- Delete Instrument Telemetry
delete from instrument_telemetry where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980'
    );
	
-- Delete Instrument Status
delete from instrument_status where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980'
    );

-- Delete Instrument Group Instruments
delete from instrument_group_instruments where instrument_id in (
	select id from instrument i
	where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980'
	);

-- Delete Instrument Groups
delete from instrument_group 
where project_id = 'a012e753-9eff-426d-b0ee-090b430d1980';
	
--Delete Collection Groups

--Delete collection_group_timeseries

-- Delete Instruments
delete from instrument i 
where i.project_id = 'a012e753-9eff-426d-b0ee-090b430d1980';

-- Delete Alert Config

-- Delete Alert

--########################################
-- INSERT new data (built by script)
--########################################


--INSERT INSTRUMENTS--COUNT:29
INSERT INTO instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('089fbda3-5a6a-408d-aec2-4e108910d94b', False, 'avnn6', 'AVNN6', null, ST_GeomFromText('POINT(-77.7566 42.9184)',4326), null, null, '2021-03-12T16:45:21.630469Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', False, 'blbn6', 'BLBN6', null, ST_GeomFromText('POINT(-77.6806 43.0922)',4326), null, null, '2021-03-12T16:45:21.630825Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', False, 'chcn6', 'CHCN6', null, ST_GeomFromText('POINT(-77.8822 43.1008)',4326), null, null, '2021-03-12T16:45:21.630985Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', False, 'dsvn6', 'DSVN6', null, ST_GeomFromText('POINT(-77.7064 42.5322)',4326), null, null, '2021-03-12T16:45:21.631193Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224775'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', False, 'garn6', 'GARN6', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-12T16:45:21.631427Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', False, 'hnyn6', 'HNYN6', null, ST_GeomFromText('POINT(-77.5869 42.9567)',4326), null, null, '2021-03-12T16:45:21.631609Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04229500'),
('08735594-14f0-4594-aa38-fb0805968918', False, 'blackcr-churchvl', 'BlackCr Churchvl', null, ST_GeomFromText('POINT(-77.8822 43.1006)',4326), null, null, '2021-03-12T16:45:21.631836Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231000'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', False, 'genr-portagevill', 'GenR Portagevill', null, ST_GeomFromText('POINT(-78.0422 42.5703)',4326), null, null, '2021-03-12T16:45:21.632074Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', False, 'oatkacr-garbutt', 'OatkaCr Garbutt', null, ST_GeomFromText('POINT(-77.7914 43.01)',4326), null, null, '2021-03-12T16:45:21.632371Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230500'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', False, 'knvn6', 'KNVN6', null, ST_GeomFromText('POINT(-78.3103 43.3011)',4326), null, null, '2021-03-12T16:45:21.632569Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '0422016550'),
('3b260872-3454-473b-bfb7-087ed0e809b0', False, 'mbyp1', 'MBYP1', null, ST_GeomFromText('POINT(-77.2736 41.8425)',4326), null, null, '2021-03-12T16:45:21.632863Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01518420'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', False, 'olnn6', 'OLNN6', null, ST_GeomFromText('POINT(-78.4511 42.0731)',4326), null, null, '2021-03-12T16:45:21.633203Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010820'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', False, 'rohn6', 'ROHN6', null, ST_GeomFromText('POINT(-77.6163 43.1417)',4326), null, null, '2021-03-12T16:45:21.633494Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04231600'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', False, 'shnp1', 'SHNP1', null, ST_GeomFromText('POINT(-78.1983 41.9617)',4326), null, null, '2021-03-12T16:45:21.634043Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '03010655'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', False, 'genr-wellsville', 'GenR Wellsville', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-12T16:45:21.634387Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', False, 'jonn6', 'JONN6', null, ST_GeomFromText('POINT(-77.8386 42.7667)',4326), null, null, '2021-03-12T16:45:21.634674Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227500'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', False, 'mount-morris', 'Mount Morris', null, ST_GeomFromText('POINT(-77.9071 42.7333)',4326), null, null, '2021-03-12T16:45:21.634946Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04224000'),
('37989e13-f192-4d1f-aa24-62591f8d731d', False, 'mount-morris-tailwater', 'Mount Morris-Tailwater', null, ST_GeomFromText('POINT(-77.9109 42.7332)',4326), null, null, '2021-03-12T16:45:21.634946Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('90664d51-0503-45cf-85bb-5d46fa579d0f', False, 'ptgn6', 'PTGN6', null, ST_GeomFromText('POINT(-78.0431 42.5697)',4326), null, null, '2021-03-12T16:45:21.635399Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04223000'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', False, 'weln6', 'WELN6', null, ST_GeomFromText('POINT(-77.9572 42.1222)',4326), null, null, '2021-03-12T16:45:21.635707Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04221000'),
('f063108f-9c3b-49fd-b981-1be021278ebf', False, 'wrsn6', 'WRSN6', null, ST_GeomFromText('POINT(-78.1375 42.7447)',4326), null, null, '2021-03-12T16:45:21.636012Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('15614ca3-c115-4cff-a021-c43ef352fda8', False, 'rcrn6', 'RCRN6', null, ST_GeomFromText('POINT(-77.6025 43.258)',4326), null, null, '2021-03-12T16:45:21.636321Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, null),
('31be5853-f839-4cc2-80cb-071210651059', False, 'elkp1', 'ELKP1', null, ST_GeomFromText('POINT(-77.3025 41.9875)',4326), null, null, '2021-03-12T16:45:21.636578Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01519200'),
('8df04db9-392f-4216-9776-66defd9d8d32', False, 'frkn6', 'FRKN6', null, ST_GeomFromText('POINT(-78.4636 42.3294)',4326), null, null, '2021-03-12T16:45:21.636856Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '421946078274901'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', False, 'hrln6', 'HRLN6', null, ST_GeomFromText('POINT(-77.7044 42.3489)',4326), null, null, '2021-03-12T16:45:21.637182Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01523000'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', False, 'canaseragashaker', 'CanaseragaShaker', null, ST_GeomFromText('POINT(-77.8414 42.7361)',4326), null, null, '2021-03-12T16:45:21.637556Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04227000'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', False, 'oatkacr-warsaw', 'OatkaCr Warsaw', null, ST_GeomFromText('POINT(-78.1375 42.7442)',4326), null, null, '2021-03-12T16:45:21.637935Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04230380'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', False, 'genr-avon', 'GenR Avon', null, ST_GeomFromText('POINT(-77.7572 42.9178)',4326), null, null, '2021-03-12T16:45:21.638262Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '04228500'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', False, 'akln6', 'AKLN6', null, ST_GeomFromText('POINT(-77.7167 42.3958)',4326), null, null, '2021-03-12T16:45:21.638580Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'a012e753-9eff-426d-b0ee-090b430d1980', '00000000-0000-0000-0000-000000000000', null, '01521000');

--INSERT INSTRUMENT STATUS--
INSERT INTO instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('78435764-f217-48ba-84a8-c7b3f08630e5', '089fbda3-5a6a-408d-aec2-4e108910d94b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.630469Z'),
('7d4f2bca-411c-4a83-93e2-ca5c2f233ecc', '209980f3-ab26-42d6-9fe7-13c0b6221f88', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.630825Z'),
('4426c51f-dfe5-4885-a5af-0b6c5685f1f4', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.630985Z'),
('d6476cad-be3f-4ff0-ac16-97f49f39321a', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.631193Z'),
('bb111ea6-769c-4fd2-a4de-06b560e24db7', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.631427Z'),
('ef3a15e8-6baf-4faf-bfa9-10f5c3e504be', '1dc2a261-8b00-4039-8d3f-f48666e97729', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.631609Z'),
('37a523da-9c71-4108-9de1-1068773a3eda', '08735594-14f0-4594-aa38-fb0805968918', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.631836Z'),
('55794c4d-65f4-4476-989e-d35e51fdb152', '871485de-0116-4f4a-afcd-30f3f8ec02c5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.632074Z'),
('cfb99de5-74a9-4fe5-838e-b27228a13c59', '4e06209a-9e0b-4430-87fe-a7ac49d92160', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.632371Z'),
('cf637df8-4851-41e4-b13b-0f63d6ed551e', 'a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.632569Z'),
('2b94dcac-cab5-45fa-be8e-52ddd2682673', '3b260872-3454-473b-bfb7-087ed0e809b0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.632863Z'),
('17a54b21-21eb-43bb-ac61-e1257204b28b', '606dcda7-8287-4435-9d66-deb4c03b4bf6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.633203Z'),
('b536efc2-5a0f-446b-b44f-315451402c0c', '0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.633494Z'),
('8160a965-1abd-4ad0-a61d-bc1374ae4076', '2673f0e2-16a4-443a-ac86-7c337d960f4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.634043Z'),
('ad5c3a1d-599b-4b1f-8593-473c50f98dcf', 'c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.634387Z'),
('92a16ea2-734e-4bf4-bad0-c8cd5e1ca178', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.634674Z'),
('55468288-d170-49f1-aec5-4827164120da', 'ef991379-a625-4de7-bc6a-2902fea1bc79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.634946Z'),
('58e5cf2a-73ab-4d60-b219-d8318a94da96', '37989e13-f192-4d1f-aa24-62591f8d731d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.634946Z'),
('f1b4e5e1-3500-4590-b5c5-db15987542e2', '90664d51-0503-45cf-85bb-5d46fa579d0f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.635399Z'),
('473efbb2-b1ed-4246-8cb7-519250982cbf', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.635707Z'),
('d10804a2-0033-4890-a1e2-eb4800f8c548', 'f063108f-9c3b-49fd-b981-1be021278ebf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.636012Z'),
('40dfaf56-eb73-48f5-b430-5c27f0f0aef9', '15614ca3-c115-4cff-a021-c43ef352fda8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.636321Z'),
('6a997ac4-8711-46be-ae8b-bf867f825b00', '31be5853-f839-4cc2-80cb-071210651059', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.636578Z'),
('bed2acd4-cca2-48e4-bcbf-516dc46d9f98', '8df04db9-392f-4216-9776-66defd9d8d32', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.636856Z'),
('547e0540-4532-4eaa-9f42-f8b1177c299a', 'a6e1707c-08f9-4484-ba65-cd054e9561bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.637182Z'),
('f05fff01-75a0-4075-bcbe-7a92d7bc5e8b', '4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.637556Z'),
('bc1d6e30-58da-4cf9-a2e8-36a2cc241cf8', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.637935Z'),
('2b460edc-5347-4c18-8931-11e0ac1cb82d', 'c559d145-3351-4e6e-bf8d-df05f69814ea', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.638262Z'),
('ec65457b-9a28-48a1-98b3-b10115cf80c8', 'b695967d-c050-4428-ab1e-db9407fe9d2f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:45:21.638580Z');

--INSERT TELEMETRY_GOES--COUNT:28
INSERT INTO telemetry_goes (id, nesdis_id) select '7f01f60b-cdcb-44aa-ad55-dfd2b0060bc6', 'CE7EB098' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EB098');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f62da7a6-5c32-44a0-b77a-b412b7db7c08', 'CE7EBE4A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EBE4A');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fe367321-f1a5-4d1f-85f1-53c7b6330cb3', 'CE7EC608' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC608');
INSERT INTO telemetry_goes (id, nesdis_id) select '7622ec04-61ba-41e5-8b56-d7975f78ebf7', 'CE7EC8DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EC8DA');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd446ce1f-aec3-4202-9a5a-c89160eceb66', 'CE7ED57E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7ED57E');
INSERT INTO telemetry_goes (id, nesdis_id) select '6fc7fab3-fbda-4dcc-ada3-4653f3ef8406', 'CE7EDBAC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EDBAC');
INSERT INTO telemetry_goes (id, nesdis_id) select '033ef124-f7a3-44aa-a4e0-31bba7564559', '1715F330' where not exists (select 1 from telemetry_goes where nesdis_id = '1715F330');
INSERT INTO telemetry_goes (id, nesdis_id) select 'bd287b31-4aff-430f-b7c9-5097701568a2', '1716615C' where not exists (select 1 from telemetry_goes where nesdis_id = '1716615C');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd4c5c3ea-12dd-4926-b363-3b9ab6867e31', 'DD8362E0' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8362E0');
INSERT INTO telemetry_goes (id, nesdis_id) select 'c3baf528-4f3c-4d28-aa5a-428d265da83c', '172024E8' where not exists (select 1 from telemetry_goes where nesdis_id = '172024E8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'aa4000d5-2448-43a4-996b-dec888cd4da3', 'CE5D2DC8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D2DC8');
INSERT INTO telemetry_goes (id, nesdis_id) select '8efc2e04-c09e-4312-aebb-9ff1234e2200', 'CE6B8B8E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B8B8E');
INSERT INTO telemetry_goes (id, nesdis_id) select '1bbe0df7-653e-48b2-8d02-5d78409c5cd1', 'DDA2D27A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA2D27A');
INSERT INTO telemetry_goes (id, nesdis_id) select 'a033a47f-0d10-41e1-9b6e-c0dd524b444e', 'CE216420' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE216420');
INSERT INTO telemetry_goes (id, nesdis_id) select 'af812385-a995-4481-bcb8-a3b9f715e3b8', 'DD6E777A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD6E777A');
INSERT INTO telemetry_goes (id, nesdis_id) select '1d5e9b80-61fc-4740-912e-49d27da4be7a', 'CE7EE0E4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EE0E4');
INSERT INTO telemetry_goes (id, nesdis_id) select '71cad6e8-1bd4-4f36-9176-04cdc490a385', 'CE1FF35A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE1FF35A');
INSERT INTO telemetry_goes (id, nesdis_id) select '940cd9af-b717-4cf5-b2e9-11c9489c9c8f', 'CE7EEE36' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EEE36');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd2e79b0c-255c-43bc-befc-9b5900ed4b10', 'CE7EFD40' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EFD40');
INSERT INTO telemetry_goes (id, nesdis_id) select '4b657dc2-5af2-43a7-8712-43d427e1c6bf', 'CE7EF392' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7EF392');
INSERT INTO telemetry_goes (id, nesdis_id) select 'eb4c51c8-a962-4b3a-b520-982c9dc596cb', '33655136' where not exists (select 1 from telemetry_goes where nesdis_id = '33655136');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd9395a29-2c60-4b52-9619-73aaf12aa2c4', 'CE5D231A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5D231A');
INSERT INTO telemetry_goes (id, nesdis_id) select '05216805-1451-4e61-821d-5c6bbad2757e', 'CE19401A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE19401A');
INSERT INTO telemetry_goes (id, nesdis_id) select '9db8b32f-9a4d-45a6-aa10-01f2ea1f1c63', 'CE59652A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE59652A');
INSERT INTO telemetry_goes (id, nesdis_id) select '15ac2a82-3c93-4121-8ff0-6fb2561f6135', 'DF02AF2A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF02AF2A');
INSERT INTO telemetry_goes (id, nesdis_id) select '617c1ba3-075d-437d-be6f-9c5fffd10946', 'DD66C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD66C052');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e79dd5ad-578f-4b60-b331-0e4d0537dc28', 'DF055D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DF055D9A');
INSERT INTO telemetry_goes (id, nesdis_id) select 'df14f769-e737-4d98-8cf4-24492b725f8c', 'CE437E42' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE437E42');

--INSERT INSTRUMENT_TELEMETRY--COUNT:28
INSERT INTO instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('089fbda3-5a6a-408d-aec2-4e108910d94b', '10a32652-af43-4451-bd52-4980c5690cc9', '7f01f60b-cdcb-44aa-ad55-dfd2b0060bc6'),
('209980f3-ab26-42d6-9fe7-13c0b6221f88', '10a32652-af43-4451-bd52-4980c5690cc9', 'f62da7a6-5c32-44a0-b77a-b412b7db7c08'),
('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '10a32652-af43-4451-bd52-4980c5690cc9', 'fe367321-f1a5-4d1f-85f1-53c7b6330cb3'),
('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '10a32652-af43-4451-bd52-4980c5690cc9', '7622ec04-61ba-41e5-8b56-d7975f78ebf7'),
('891c7d48-fa1b-4f4a-a7dd-7393258bb575', '10a32652-af43-4451-bd52-4980c5690cc9', 'd446ce1f-aec3-4202-9a5a-c89160eceb66'),
('1dc2a261-8b00-4039-8d3f-f48666e97729', '10a32652-af43-4451-bd52-4980c5690cc9', '6fc7fab3-fbda-4dcc-ada3-4653f3ef8406'),
('08735594-14f0-4594-aa38-fb0805968918', '10a32652-af43-4451-bd52-4980c5690cc9', '033ef124-f7a3-44aa-a4e0-31bba7564559'),
('871485de-0116-4f4a-afcd-30f3f8ec02c5', '10a32652-af43-4451-bd52-4980c5690cc9', 'bd287b31-4aff-430f-b7c9-5097701568a2'),
('4e06209a-9e0b-4430-87fe-a7ac49d92160', '10a32652-af43-4451-bd52-4980c5690cc9', 'd4c5c3ea-12dd-4926-b363-3b9ab6867e31'),
('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', '10a32652-af43-4451-bd52-4980c5690cc9', 'c3baf528-4f3c-4d28-aa5a-428d265da83c'),
('3b260872-3454-473b-bfb7-087ed0e809b0', '10a32652-af43-4451-bd52-4980c5690cc9', 'aa4000d5-2448-43a4-996b-dec888cd4da3'),
('606dcda7-8287-4435-9d66-deb4c03b4bf6', '10a32652-af43-4451-bd52-4980c5690cc9', '8efc2e04-c09e-4312-aebb-9ff1234e2200'),
('0db970b5-3c46-4f54-b17b-8e861c0a5d65', '10a32652-af43-4451-bd52-4980c5690cc9', '1bbe0df7-653e-48b2-8d02-5d78409c5cd1'),
('2673f0e2-16a4-443a-ac86-7c337d960f4f', '10a32652-af43-4451-bd52-4980c5690cc9', 'a033a47f-0d10-41e1-9b6e-c0dd524b444e'),
('c05efa0c-fa2c-4776-a49b-dd68f70f5104', '10a32652-af43-4451-bd52-4980c5690cc9', 'af812385-a995-4481-bcb8-a3b9f715e3b8'),
('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '10a32652-af43-4451-bd52-4980c5690cc9', '1d5e9b80-61fc-4740-912e-49d27da4be7a'),
('ef991379-a625-4de7-bc6a-2902fea1bc79', '10a32652-af43-4451-bd52-4980c5690cc9', '71cad6e8-1bd4-4f36-9176-04cdc490a385'),
('90664d51-0503-45cf-85bb-5d46fa579d0f', '10a32652-af43-4451-bd52-4980c5690cc9', '940cd9af-b717-4cf5-b2e9-11c9489c9c8f'),
('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '10a32652-af43-4451-bd52-4980c5690cc9', 'd2e79b0c-255c-43bc-befc-9b5900ed4b10'),
('f063108f-9c3b-49fd-b981-1be021278ebf', '10a32652-af43-4451-bd52-4980c5690cc9', '4b657dc2-5af2-43a7-8712-43d427e1c6bf'),
('15614ca3-c115-4cff-a021-c43ef352fda8', '10a32652-af43-4451-bd52-4980c5690cc9', 'eb4c51c8-a962-4b3a-b520-982c9dc596cb'),
('31be5853-f839-4cc2-80cb-071210651059', '10a32652-af43-4451-bd52-4980c5690cc9', 'd9395a29-2c60-4b52-9619-73aaf12aa2c4'),
('8df04db9-392f-4216-9776-66defd9d8d32', '10a32652-af43-4451-bd52-4980c5690cc9', '05216805-1451-4e61-821d-5c6bbad2757e'),
('a6e1707c-08f9-4484-ba65-cd054e9561bd', '10a32652-af43-4451-bd52-4980c5690cc9', '9db8b32f-9a4d-45a6-aa10-01f2ea1f1c63'),
('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '10a32652-af43-4451-bd52-4980c5690cc9', '15ac2a82-3c93-4121-8ff0-6fb2561f6135'),
('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '10a32652-af43-4451-bd52-4980c5690cc9', '617c1ba3-075d-437d-be6f-9c5fffd10946'),
('c559d145-3351-4e6e-bf8d-df05f69814ea', '10a32652-af43-4451-bd52-4980c5690cc9', 'e79dd5ad-578f-4b60-b331-0e4d0537dc28'),
('b695967d-c050-4428-ab1e-db9407fe9d2f', '10a32652-af43-4451-bd52-4980c5690cc9', 'df14f769-e737-4d98-8cf4-24492b725f8c');

--INSERT TIMESERIES--COUNT:28
INSERT INTO timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
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
('f4811c7b-1f20-4de5-a5ad-6116d4ba8eaa','unknown-wv','Unknown WV','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('c5dc016f-9692-4271-88b9-01db31a4c453','stage','Stage','0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b3a538af-293e-4d31-99df-4887f639e224','unknown-ta','Unknown TA','0db970b5-3c46-4f54-b17b-8e861c0a5d65', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
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
('fa9ea163-afde-4d80-a6d4-1d915f91923f','unknown-hb','Unknown HB','15614ca3-c115-4cff-a021-c43ef352fda8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b42dca04-13f6-439e-90a4-4fd14de746d4','unknown-ta','Unknown TA','15614ca3-c115-4cff-a021-c43ef352fda8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('4c3d8edd-ea01-4ead-937e-25ee78446350','unknown-pa','Unknown PA','15614ca3-c115-4cff-a021-c43ef352fda8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
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
('337a16aa-0c40-487a-b41f-c40c7d2adbd3','unknown-wv','Unknown WV','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('afa5818c-95de-44b2-9521-90480aff6c93','unknown-tw','Unknown TW','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('cbf7c063-dec1-417b-abde-56302cf5ffac','voltage','Voltage','4f8c4f01-5e19-4a60-b55b-e7013e8977bb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dfdee47b-f6d3-4c79-979a-115a7dc8aa90','stage','Stage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1fe6104e-a3b2-4683-ad37-9e7c86aa4b84','precipitation','Precipitation','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('011270a6-76d7-4e20-941c-2dc6d8d18da4','voltage','Voltage','abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2aaf99a5-00d7-409c-b5a9-3abeea28791a','stage','Stage','c559d145-3351-4e6e-bf8d-df05f69814ea', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1863e041-6363-48ce-b6f6-b684a7162e3e','unknown-wv','Unknown WV','c559d145-3351-4e6e-bf8d-df05f69814ea', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('8de721a8-4599-4b7f-9ddd-a4916234c9be','voltage','Voltage','c559d145-3351-4e6e-bf8d-df05f69814ea', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8555f21e-8847-4117-aebc-55f91cd48c29','stage','Stage','b695967d-c050-4428-ab1e-db9407fe9d2f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cc8aba30-7ee8-4cc2-91b0-95c91c2a919c','precipitation','Precipitation','b695967d-c050-4428-ab1e-db9407fe9d2f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c83c55ec-21b4-4b62-8440-b12c130dcfad','voltage','Voltage','b695967d-c050-4428-ab1e-db9407fe9d2f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64');
