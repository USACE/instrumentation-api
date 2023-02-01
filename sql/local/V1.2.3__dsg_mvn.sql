INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('e38f3d27-21fe-4e35-8de5-3aea8580e909', 'f81f5659-ce57-4c87-9c7a-0d685a983bfd', 
    'new-orleans-district-streamgages', 'New Orleans District Streamgages', 'new-orleans-district-streamgages.jpg');

--#################################################
--DELETE existing data (mostly for dev, test, prod)
--#################################################

-- Delete Timeseries Measurements
delete from timeseries_measurement where timeseries_id in (
	select t.id from instrument i 
	join timeseries t on i.id = t.instrument_id
	where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909');

-- Delete Timeseries
delete from timeseries where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909');
	
-- Delete Telemetry GOES
delete from telemetry_goes where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='10a32652-af43-4451-bd52-4980c5690cc9'
	and instrument_id in (
		select i.id from instrument i
		where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909')
	);

-- Delete Telemetry Iridium
delete from telemetry_iridium where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='c0b03b0d-bfce-453a-b5a9-636118940449'
	and	instrument_id in (
		select i.id from instrument i
		where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909')
	);
	
-- Delete Instrument Telemetry
delete from instrument_telemetry where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909'
    );
	
-- Delete Instrument Status
delete from instrument_status where instrument_id in (
	select i.id from instrument i
	where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909'
    );

-- Delete Instrument Group Instruments
delete from instrument_group_instruments where instrument_id in (
	select id from instrument i
	where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909'
	);

-- Delete Instrument Groups
delete from instrument_group 
where project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909';
	
--Delete Collection Groups

--Delete collection_group_timeseries

-- Delete Instruments
delete from instrument i 
where i.project_id = 'e38f3d27-21fe-4e35-8de5-3aea8580e909';

-- Delete Alert Config

-- Delete Alert

--########################################
-- INSERT new data (built by script)
--########################################

--Ignoring 76010 platform sensor site/instrument, already in instruments unique list
--Ignoring 76005 platform sensor site/instrument, already in instruments unique list
--Ignoring 76005 platform sensor site/instrument, already in instruments unique list
--Ignoring 76025 platform sensor site/instrument, already in instruments unique list
--Ignoring 76025 platform sensor site/instrument, already in instruments unique list
--Ignoring 76025 platform sensor site/instrument, already in instruments unique list
--Ignoring 76025 platform sensor site/instrument, already in instruments unique list
--Ignoring 76025 platform sensor site/instrument, already in instruments unique list
--Ignoring 76065 platform sensor site/instrument, already in instruments unique list
--Ignoring 76065 platform sensor site/instrument, already in instruments unique list
--Ignoring 76065 platform sensor site/instrument, already in instruments unique list
--Ignoring 76260 platform sensor site/instrument, already in instruments unique list
--Ignoring 76260 platform sensor site/instrument, already in instruments unique list
--Ignoring 76360 platform sensor site/instrument, already in instruments unique list
--Ignoring 76360 platform sensor site/instrument, already in instruments unique list
--Ignoring 76592 platform sensor site/instrument, already in instruments unique list
--Ignoring 76592 platform sensor site/instrument, already in instruments unique list
--Ignoring 76680 platform sensor site/instrument, already in instruments unique list
--Ignoring 76680 platform sensor site/instrument, already in instruments unique list
--Ignoring 76680 platform sensor site/instrument, already in instruments unique list
--Ignoring 76600 platform sensor site/instrument, already in instruments unique list
--Ignoring 76600 platform sensor site/instrument, already in instruments unique list
--Ignoring 76600 platform sensor site/instrument, already in instruments unique list
--Ignoring 76720 platform sensor site/instrument, already in instruments unique list
--Ignoring 76720 platform sensor site/instrument, already in instruments unique list
--Ignoring 76880 platform sensor site/instrument, already in instruments unique list
--Ignoring 76880 platform sensor site/instrument, already in instruments unique list
--Ignoring 82250 platform sensor site/instrument, already in instruments unique list
--Ignoring 82250 platform sensor site/instrument, already in instruments unique list
--Ignoring 82250 platform sensor site/instrument, already in instruments unique list
--Ignoring 82718 platform sensor site/instrument, already in instruments unique list
--Ignoring 82718 platform sensor site/instrument, already in instruments unique list
--Ignoring 82718 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82745 platform sensor site/instrument, already in instruments unique list
--Ignoring 82760 platform sensor site/instrument, already in instruments unique list
--Ignoring 82760 platform sensor site/instrument, already in instruments unique list
--Ignoring 82760 platform sensor site/instrument, already in instruments unique list
--Ignoring 82770 platform sensor site/instrument, already in instruments unique list
--Ignoring 82772 platform sensor site/instrument, already in instruments unique list
--Ignoring 82772 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85627 platform sensor site/instrument, already in instruments unique list
--Ignoring 85631 platform sensor site/instrument, already in instruments unique list
--Ignoring 85631 platform sensor site/instrument, already in instruments unique list
--Ignoring 85633 platform sensor site/instrument, already in instruments unique list
--Ignoring 85633 platform sensor site/instrument, already in instruments unique list
--Ignoring 85636 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85635 platform sensor site/instrument, already in instruments unique list
--Ignoring 85637 platform sensor site/instrument, already in instruments unique list
--Ignoring 85637 platform sensor site/instrument, already in instruments unique list
--Ignoring 85639 platform sensor site/instrument, already in instruments unique list
--Ignoring 85639 platform sensor site/instrument, already in instruments unique list
--Ignoring 85639 platform sensor site/instrument, already in instruments unique list
--Ignoring 85641 platform sensor site/instrument, already in instruments unique list
--Ignoring 85641 platform sensor site/instrument, already in instruments unique list
--Ignoring 85641 platform sensor site/instrument, already in instruments unique list
--Ignoring 85653 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85652 platform sensor site/instrument, already in instruments unique list
--Ignoring 85655 platform sensor site/instrument, already in instruments unique list
--Ignoring 85655 platform sensor site/instrument, already in instruments unique list
--Ignoring 85657 platform sensor site/instrument, already in instruments unique list
--Ignoring 85657 platform sensor site/instrument, already in instruments unique list
--Ignoring 85659 platform sensor site/instrument, already in instruments unique list
--Ignoring 85659 platform sensor site/instrument, already in instruments unique list
--Ignoring 85660 platform sensor site/instrument, already in instruments unique list
--Ignoring 85667 platform sensor site/instrument, already in instruments unique list
--Ignoring 85667 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85670 platform sensor site/instrument, already in instruments unique list
--Ignoring 85765 platform sensor site/instrument, already in instruments unique list
--Ignoring 85765 platform sensor site/instrument, already in instruments unique list
--Ignoring 85765 platform sensor site/instrument, already in instruments unique list
--Ignoring 01143 platform sensor site/instrument, already in instruments unique list
--Ignoring 01143 platform sensor site/instrument, already in instruments unique list
--Ignoring 01143 platform sensor site/instrument, already in instruments unique list
--Ignoring 01142 platform sensor site/instrument, already in instruments unique list
--Ignoring 01143 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01160 platform sensor site/instrument, already in instruments unique list
--Ignoring 01320 platform sensor site/instrument, already in instruments unique list
--Ignoring 01320 platform sensor site/instrument, already in instruments unique list
--Ignoring 01340 platform sensor site/instrument, already in instruments unique list
--Ignoring 01340 platform sensor site/instrument, already in instruments unique list
--Ignoring 01380 platform sensor site/instrument, already in instruments unique list
--Ignoring 01380 platform sensor site/instrument, already in instruments unique list
--Ignoring 01400 platform sensor site/instrument, already in instruments unique list
--Ignoring 01400 platform sensor site/instrument, already in instruments unique list
--Ignoring 02050 platform sensor site/instrument, already in instruments unique list
--Ignoring 02200 platform sensor site/instrument, already in instruments unique list
--Ignoring 02200 platform sensor site/instrument, already in instruments unique list
--Ignoring 03060 platform sensor site/instrument, already in instruments unique list
--Ignoring 03550 platform sensor site/instrument, already in instruments unique list
--Ignoring 03550 platform sensor site/instrument, already in instruments unique list
--Ignoring 43500 platform sensor site/instrument, already in instruments unique list
--Ignoring 43500 platform sensor site/instrument, already in instruments unique list
--Ignoring 52560 platform sensor site/instrument, already in instruments unique list
--Ignoring 52560 platform sensor site/instrument, already in instruments unique list
--Ignoring 52560 platform sensor site/instrument, already in instruments unique list
--Ignoring 52560 platform sensor site/instrument, already in instruments unique list
--Ignoring 52560 platform sensor site/instrument, already in instruments unique list
--Ignoring 70675 platform sensor site/instrument, already in instruments unique list
--Ignoring 70675 platform sensor site/instrument, already in instruments unique list
--Ignoring 70675 platform sensor site/instrument, already in instruments unique list
--Ignoring 73472 platform sensor site/instrument, already in instruments unique list
--Ignoring 73472 platform sensor site/instrument, already in instruments unique list
--Ignoring 73472 platform sensor site/instrument, already in instruments unique list
--Ignoring 73472 platform sensor site/instrument, already in instruments unique list
--Ignoring 73473 platform sensor site/instrument, already in instruments unique list
--Ignoring 73473 platform sensor site/instrument, already in instruments unique list
--Ignoring 73473 platform sensor site/instrument, already in instruments unique list

--INSERT INSTRUMENTS--COUNT:215
INSERT INTO instrument(id, deleted, slug, name, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('15dacad3-a7e5-4378-91b4-34208e7039cb', False, '73600', '73600', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.948076Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('35896e62-b69f-4bc9-9fc6-78ea7c5fa92a', False, '73650', '73650', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.948289Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bbeff60a-9198-4956-9ef2-021fc851cdbd', False, 'barataria-pass-at-grand-isle-la', 'BARATARIA PASS AT GRAND ISLE LA', ST_GeomFromText('POINT(-89.9468 29.2728)',4326), null, null, '2021-03-12T16:10:12.948450Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073802516'),
('64d4fe14-952c-4f0a-b256-1ae56e687fd2', False, '76010', '76010', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.948600Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9e13d62c-b9e4-49eb-aee5-59b4d1bbd46c', False, '76005', '76005', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.948600Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d469455f-6ce9-4faf-b2fc-c321f3365a6f', False, '76025', '76025', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.949259Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b44d2de5-fa95-46fd-9254-93078405d35d', False, '76024', '76024', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.949259Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('41e9219e-2f1f-439f-a724-a3e5c7c7cf73', False, 'giww-east-storm-surge-barrier-at-new-orleans-la', 'GIWW EAST STORM SURGE BARRIER AT NEW ORLEANS LA', ST_GeomFromText('POINT(-89.9019 30.0144)',4326), null, null, '2021-03-12T16:10:12.950027Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073802339'),
('6ee253df-10ae-493c-b307-de8e85464c07', False, '76032', '76032', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.950027Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('72ca06c0-2710-4135-be3a-bf93d8c51dfd', False, '76030', '76030', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.950027Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('95b436bc-842a-4ad8-83e7-2410a56c439c', False, '76040', '76040', ST_GeomFromText('POINT(-89.9374 30.0067)',4326), null, null, '2021-03-12T16:10:12.950279Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('016d2287-c4a8-4544-9588-dfb961726c26', False, '76060', '76060', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.950439Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3f1daace-2b8e-4d6c-a21c-1a23359ae00b', False, '76065', '76065', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.950616Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('a757599f-b2e5-4fbe-b40d-8522c3b04af0', False, '76062', '76062', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.950616Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('38fa06c8-6457-4523-b4ef-3ebffa520151', False, '76220', '76220', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.951160Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b9e94249-d8e9-41ac-b683-bde39fa79ed4', False, 'coe-freshwater-c-at-freshwater-b-lock-south', '(COE) FRESHWATER C. AT FRESHWATER B. LOCK (SOUTH)', ST_GeomFromText('POINT(-92.1956 29.7833)',4326), null, null, '2021-03-12T16:10:12.951444Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07389700'),
('616fa910-eabd-4718-940f-c7ec2afeffd8', False, '76260', '76260', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.951621Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c9e023a7-9cd7-4443-9661-c0f629bc6c8a', False, '76265', '76265', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.951621Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6d781aa1-93e5-4aa5-85fe-9d057ead4247', False, '76305', '76305', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952221Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4adb6c37-4ae1-43e1-9540-e5595bf86d1b', False, '76320', '76320', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952424Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('67cf3e31-f3de-4413-930d-1cf4903c454e', False, '76360', '76360', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952571Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('eafbba9f-6875-462c-a8ec-99e224c37765', False, '76362', '76362', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952571Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('0b4fb382-4721-4c93-8c92-eb47812e0bac', False, '76400', '76400', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952571Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bacf256e-eb96-4c36-8954-bd92a306f4d8', False, '76480', '76480', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.952872Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5524e00c-1f89-418d-9740-6de8645b8333', False, 'giww-at-bayou-sale-ridge-near-franklin-la', 'GIWW AT BAYOU SALE RIDGE NEAR FRANKLIN LA', ST_GeomFromText('POINT(-91.4706 29.6808)',4326), null, null, '2021-03-12T16:10:12.953035Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381670'),
('d7281dd7-778d-4c3a-90a4-b21bb641e850', False, '76592', '76592', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.953219Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('17968170-2ff2-4cd8-9809-818f88732879', False, '76595', '76595', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.953219Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('cabcd7a9-db9e-415e-ba84-cbe5f1f5c768', False, '76593', '76593', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.953219Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f6e8916f-296a-49e9-aca1-82665748ca62', False, '76680', '76680', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.953537Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ee622c1e-70e8-4620-8f61-2261c0ed13ac', False, '76600', '76600', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.953537Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('70ba9edc-40ea-42ad-a8a3-ceabb25b05e2', False, '76720', '76720', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954216Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b226f922-c336-4929-8bb1-9e9ddf4b59e6', False, '76724', '76724', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954216Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7c82c421-5294-40a7-87ff-7af6e42a18a5', False, '76800', '76800', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954216Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('694b13bf-449a-4c44-8aa2-c38469c07ed5', False, '76880', '76880', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954530Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f15dbad6-ccb7-4a72-9298-5cd33a8b4797', False, '76884', '76884', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954530Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c42b64ce-0460-41f3-8bbf-86a8da1d4c00', False, '76960', '76960', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954530Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('a392c3e8-738d-4feb-afc1-00d064e4298e', False, '82250', '82250', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954845Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e0ecd779-e73c-47f4-a906-55c8002dcba7', False, '82260', '82260', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.954845Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8add8896-c24d-44b3-9853-c189d2d83695', False, '82700', '82700', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.955289Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('1592e55f-ff35-40a9-80c4-7c81e4f620b3', False, '82718', '82718', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.955475Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4d62ac6e-35b9-41cd-80de-7468d2bad87c', False, '82715', '82715', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.955475Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3604b451-24dc-423f-9e8a-f0bbf9d7d8f2', False, '82720', '82720', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.955929Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9d3ea29e-59a3-46dd-a091-d8fed56f03ec', False, '82725', '82725', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.956037Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4c7ce601-f232-4634-936c-d7604eb54e3b', False, '82740', '82740', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.956209Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e213714d-2656-4719-84f5-61e208032038', False, '82745', '82745', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.956353Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('965c672f-d3d2-41a9-a537-7626ee1da6d1', False, '82742', '82742', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.956353Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6bdca601-8044-4830-85e5-7f34dd5d5865', False, '82760', '82760', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.957042Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c7e40f37-2390-404f-82fe-c9e7496b8801', False, '82762', '82762', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.957042Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', False, '82770', '82770', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.957437Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('59bc5f1d-63a4-455c-80d0-fa6ad8f641fc', False, '82772', '82772', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.957437Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8091654d-2afd-46a6-a2a7-26447b8c0df4', False, '82875', '82875', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.957812Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2cce0e9d-c608-4424-a8f8-0f76e0691338', False, 'amite-river-near-denham-springs-la', 'AMITE RIVER NEAR DENHAM SPRINGS LA', ST_GeomFromText('POINT(-90.9903 30.4639)',4326), null, null, '2021-03-12T16:10:12.958013Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07378500'),
('82d2991c-71db-4699-9e7f-5506f9614d58', False, '85300', '85300', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.958172Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('21bdce69-d2a7-49f0-807b-d29c747f2802', False, 'tickfaw-river-at-holden-la', 'TICKFAW RIVER AT HOLDEN LA', ST_GeomFromText('POINT(-90.6772 30.5036)',4326), null, null, '2021-03-12T16:10:12.958428Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07376000'),
('28bc2a6e-72a9-4e21-9da0-29ce3dda495d', False, 'natalbany-river-at-baptist-la', 'NATALBANY RIVER AT BAPTIST LA', ST_GeomFromText('POINT(-90.5458 30.5042)',4326), null, null, '2021-03-12T16:10:12.958696Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07376500'),
('00fb2016-699b-45cf-a44f-2944e0860c26', False, 'tangipahoa-river-at-robert-la', 'TANGIPAHOA RIVER AT ROBERT LA', ST_GeomFromText('POINT(-90.3617 30.5064)',4326), null, null, '2021-03-12T16:10:12.958900Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375500'),
('df1cf782-4230-4f79-8437-4e361ee67dfc', False, '85420', '85420', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.959049Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f0658ffe-4b1f-4b6c-b17b-7b8cc43ad302', False, 'tchefuncte-river-near-folsom-la', 'TCHEFUNCTE RIVER NEAR FOLSOM LA', ST_GeomFromText('POINT(-90.2486 30.6158)',4326), null, null, '2021-03-12T16:10:12.959245Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375000'),
('581ba9dc-3836-4fea-9598-fc0d9255da47', False, '85555', '85555', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.959449Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9ce6e98b-a26b-450e-b859-41bdc05bab06', False, '85575', '85575', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.959650Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('1cc92566-d593-4c0f-abd5-fd39aabf9fb1', False, '85625', '85625', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.960032Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5c3f1f19-1639-477f-81e0-8edf1b85bda7', False, '85627', '85627', ST_GeomFromText('POINT(-90.0742 30.0308)',4326), null, null, '2021-03-12T16:10:12.960348Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9a60c8bd-475f-44a9-8e14-7a0c794f0c1c', False, '85626', '85626', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.960348Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('449c9725-605c-4d0d-a0c6-8e434054b26d', False, '85631', '85631', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.961250Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c37feb2a-559b-4f6d-8b3c-75230e5adeae', False, '85630', '85630', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.961250Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('a70a73d7-47fa-4f92-9b00-6bf9f02c52df', False, '85633', '85633', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.961584Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3416688d-bd11-4ac3-86c7-f75f2d311498', False, '85632', '85632', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.961584Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d56068c3-a3c4-4553-b30a-a885570cd7d8', False, '85634', '85634', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.961964Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', False, '85636', '85636', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.962150Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5d870e7d-cf49-4b9a-8352-10a75948e282', False, '85635', '85635', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.962150Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('572ce69b-8b32-40ab-ac00-e783dc6ceb4f', False, '85637', '85637', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.963175Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('252a79a9-2538-4016-a9bd-c5edcfb701b8', False, '85638', '85638', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.963175Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bcfb5c23-300d-4d78-b8b7-f7cf40bfbd09', False, '85639', '85639', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.963531Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('0d2b2c4a-e8ab-464d-aa5e-5d4fb3145dc5', False, '85641', '85641', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.963904Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d0803aa7-bbee-48ed-b23f-7dea74b2ac59', False, '85653', '85653', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.964324Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('40f1872b-b305-44ab-bc8d-06351fd29d3a', False, '85652', '85652', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.964324Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5e19aa7a-1611-49e9-bec3-75c3421c7b5e', False, '85655', '85655', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965240Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e7665d40-d030-4b35-b91a-741546df5fb0', False, '85654', '85654', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965240Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e1993ab5-ec08-4e07-bc3c-c4db92731d1c', False, '85657', '85657', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965538Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4429e603-e648-4a97-a2ce-42d1f94347fd', False, '85656', '85656', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965538Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2db45595-fefc-4eb8-9237-8fc72561b47e', False, '85659', '85659', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965946Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('19c0ac2f-09ab-4c68-b509-71d8a3c07a99', False, '85658', '85658', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.965946Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ec40c087-d75b-4580-9ee5-337eeff32f95', False, '85660', '85660', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.966295Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4009c409-974c-46f8-86af-992ee43d5ac7', False, '85661', '85661', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.966295Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f3b99e3b-1856-4226-a73c-1c46b200fb66', False, '85667', '85667', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.966596Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f58051b1-180a-4c5d-a602-e4dbaaac3367', False, '85666', '85666', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.966596Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('72d33561-aaaa-4e90-91b4-eb73eaeaaac6', False, '85670', '85670', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.966959Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6b78716c-4975-46e4-b635-9ff79e7f0627', False, '85700', '85700', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.967531Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('a88ece8d-d3ba-4a77-a413-5dd03ebfb623', False, '85750', '85750', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.967792Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f6629ae2-98e2-4e06-b685-331189cff3a7', False, '85765', '85765', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968003Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('eb4969b9-30ae-45ea-ad68-d19948a84020', False, '85760', '85760', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968003Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('59db968a-43d1-4590-b4f7-6f7f6423e735', False, '85780', '85780', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968359Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ac71a2f1-74bb-49de-8dd4-d2b15b974978', False, '88450', '88450', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968467Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', False, 'mouth-of-atchafalaya-river-at-atchafalaya-bay', 'MOUTH OF ATCHAFALAYA RIVER AT ATCHAFALAYA BAY', ST_GeomFromText('POINT(-91.3339 29.4303)',4326), null, null, '2021-03-12T16:10:12.968642Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381654'),
('4916c4af-8019-4492-a734-fc3ff8c49074', False, '88800', '88800', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968781Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5dfc826b-a94c-4cee-b057-6fb858afc95d', False, 'acme_7350', 'ACME', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.968932Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7e772c75-64b6-485e-8294-d1d3bc6da0fe', False, 'arkcity', 'ARKCITY', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.969013Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c4c87cf8-a155-4abe-bf89-4f7af607db31', False, 'bayou-des-glaises-diversion-ch-at-moreauville-la', 'BAYOU DES GLAISES DIVERSION CH. AT MOREAUVILLE LA', ST_GeomFromText('POINT(-91.9825 31.0331)',4326), null, null, '2021-03-12T16:10:12.969105Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07383500'),
('3fcc1834-0263-4dc5-a7c8-9157c767bb4f', False, 'cairo_ca30', 'CAIRO', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.969820Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6b5d1a5d-0861-45a1-b56c-ec37c501e993', False, 'dupage-county-airport-near-st-charles-il', 'DUPAGE COUNTY AIRPORT NEAR ST CHARLES IL', ST_GeomFromText('POINT(-88.2517 41.9158)',4326), null, null, '2021-03-12T16:10:12.969954Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375175'),
('6890212a-c7f1-499e-a3fe-e7d973eb7cc1', False, 'gndl1', 'GNDL1', ST_GeomFromText('POINT(-91.4456 29.8925)',4326), null, null, '2021-03-12T16:10:12.970341Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073815450'),
('bb1389a0-bfad-430d-a66d-2f041a738cb1', False, 'green', 'GREEN', ST_GeomFromText('POINT(-91.1606 33.2925)',4326), null, null, '2021-03-12T16:10:12.970452Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07265455'),
('38ea9864-0181-475e-a8c7-b34edbf31b6b', False, 'big-shoe-heel-creek-nr-laurinburg-nc', 'BIG SHOE HEEL CREEK NR LAURINBURG NC', ST_GeomFromText('POINT(-79.3867 34.7506)',4326), null, null, '2021-03-12T16:10:12.970593Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073816202'),
('b417e9f4-ce46-47a7-8bd2-0b26506395e4', False, 'tangipahoa-river-near-kentwood-la', 'TANGIPAHOA RIVER NEAR KENTWOOD LA', ST_GeomFromText('POINT(-90.4903 30.9375)',4326), null, null, '2021-03-12T16:10:12.971045Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375300'),
('04e220cf-128b-4115-8a7b-76ac2625a81b', False, 'lower-atchafalaya-river-at-morgan-city-la', 'LOWER ATCHAFALAYA RIVER AT MORGAN CITY LA', ST_GeomFromText('POINT(-91.2118 29.6926)',4326), null, null, '2021-03-12T16:10:12.971160Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381600'),
('7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', False, 'natch', 'NATCH', ST_GeomFromText('POINT(-91.4186 31.5603)',4326), null, null, '2021-03-12T16:10:12.971257Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07290880'),
('892ca8ff-08f2-4e1b-af40-32a77baab497', False, 'rio-humacao-at-las-piedras-pr', 'RIO HUMACAO AT LAS PIEDRAS PR', ST_GeomFromText('POINT(-65.8694 18.1741)',4326), null, null, '2021-03-12T16:10:12.971452Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07376420'),
('8250ac38-3d9f-4039-9f08-149912301977', False, 'atchafalaya-river-at-simmesport-la', 'ATCHAFALAYA RIVER AT SIMMESPORT LA', ST_GeomFromText('POINT(-91.7983 30.9825)',4326), null, null, '2021-03-12T16:10:12.971601Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381490'),
('9dcf334f-2982-49c9-941c-38c91565f019', False, 'vicks', 'VICKS', ST_GeomFromText('POINT(-90.9058 32.315)',4326), null, null, '2021-03-12T16:10:12.971709Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07289000'),
('df1b2707-3dea-4b91-be0d-2f64f21544e2', False, '40700', '40700', ST_GeomFromText('POINT(-91.7389 30.7258)',4326), null, null, '2021-03-12T16:10:12.971900Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('fcf546da-e8c6-471b-b894-98a691c71251', False, '01440', '01440', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.971970Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', False, '01080', '01080', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.972095Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3eba9840-cebb-45ee-8e9f-a62d8279c86a', False, '01120', '01120', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.972249Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bc7d2d9a-8e02-4a78-b838-cadd444ed756', False, '01143', '01143', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.972384Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3a43504e-d72c-43e0-9849-d67e016fbc7a', False, '01142', '01142', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.972384Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', False, '01145', '01145', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.972880Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5171ef3b-9ef9-4984-8ab7-57b76db13621', False, '01160', '01160', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973019Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d02cd42c-3cf0-4b41-a670-1d35708b97cd', False, '52415', '52415', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973019Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7c317f0e-f87f-482e-9776-e6c8ee96170a', False, '52417', '52417', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973019Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b72006fc-5809-4985-8f3c-45baed7cfcb9', False, '01220', '01220', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973681Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4811e8d0-4c72-4c86-8bec-37e491f5ab63', False, '01260', '01260', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973802Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3f9a0fcf-4e4f-4b0c-ac88-011455d6957f', False, '01275', '01275', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.973944Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('553f696a-9292-41b1-a4bf-67236068a1d6', False, '01280', '01280', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974052Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('394d6441-43b2-4745-b97d-c6462c7bc908', False, '01300', '01300', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974176Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2711323d-5ed2-44b5-a7a6-f0614c0b422b', False, '01320', '01320', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974426Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b04e2a92-11d5-416a-a50f-d91a161456bd', False, '76205', '76205', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974426Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3838e35a-efaa-493a-bd48-52561971ce52', False, '76200', '76200', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974426Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3bd76812-033f-4367-b8c4-eceba262f393', False, '01340', '01340', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974687Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('0eb3be0e-dd34-4e25-bf96-c5443b2cb400', False, '76160', '76160', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974687Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('89cdff86-ab37-4149-aa95-5f6e43454b39', False, '76165', '76165', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974687Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f8d6ccd6-fb48-4228-b777-e423612b1f27', False, '01380', '01380', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974947Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('cf64b9de-4b12-462e-955c-669ca84e2716', False, '76240', '76240', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974947Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6c4ee0c2-6972-4e69-9555-ff05f155ad5d', False, '76245', '76245', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.974947Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7babc4e1-950e-41e8-874f-49559dd947bf', False, '01390', '01390', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975208Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ab7ee39b-8397-484a-ac29-e2e813ed6adb', False, '01400', '01400', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975345Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d93718c4-08ae-4065-8855-674c5206c1ed', False, '01480', '01480', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975574Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('de20f637-d3c7-4ea4-8ca6-604f6a127532', False, '01515', '01515', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975677Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8a7739c0-ccae-4f3b-83d7-e2f6f506e6da', False, '01516', '01516', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975800Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2006cc18-de31-4ae0-8aaa-81ba6cf1ba95', False, '01545', '01545', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975863Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e29fcc6e-a0a8-4fe0-b450-099142997846', False, '01575', '01575', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.975954Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', False, '01670', '01670', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976045Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bcc35574-bf23-4952-9760-049ba8843bd7', False, '01850', '01850', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976167Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4bbb3d28-b816-4f71-98c1-de647d8f5e38', False, '02050', '02050', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976293Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e0e57af0-7523-4bb1-b327-eff3dca3e508', False, '02100', '02100', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976293Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b3524459-4cbe-4802-835e-b415cc6657b2', False, '02200', '02200', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976487Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5d20b46f-1fbe-48c2-875e-1894f4aeb348', False, '02210', '02210', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976487Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ed2c1d1c-7bb2-4374-8e90-878ab6b70d9c', False, 'old-river-outflow-channel-below-hydropower-channel', 'OLD RIVER OUTFLOW CHANNEL BELOW HYDROPOWER CHANNEL', ST_GeomFromText('POINT(-91.6483 31.0667)',4326), null, null, '2021-03-12T16:10:12.976771Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381482'),
('eaa406e4-560c-4fa3-9c28-8f265999f17b', False, '03045', '03045', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.976965Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('bd3d7701-94b0-4390-b2f8-40e03c48508c', False, '03060', '03060', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977098Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8e5ed58a-febf-4b91-86e3-4b39d798a592', False, '03075', '03075', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977321Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('33564f35-da99-4db5-9d9d-2c96698f99e6', False, 'atchafalaya-river-at-butte-la-rose-la', 'ATCHAFALAYA RIVER AT BUTTE LA ROSE LA', ST_GeomFromText('POINT(-91.6867 30.2814)',4326), null, null, '2021-03-12T16:10:12.977451Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381515'),
('228d8881-b162-4842-92c9-75422535a5ad', False, '03210', '03210', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977561Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8f32236d-1e18-47f4-bf36-383eb8c82fbe', False, '03240', '03240', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977655Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('dccf96c3-4dff-4b3b-a7b7-446a97db654b', False, '03315', '03315', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977749Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8527c379-3f95-4ef6-aae7-4fcf94ad8691', False, '03465', '03465', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977842Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8bd798a4-1c7b-4684-ae3b-eb75a62a09cf', False, '03550', '03550', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977950Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5fbab803-eb25-480c-9d30-1064b32d9f87', False, '64400', '64400', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.977950Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2eed3135-b0ee-423d-b8a6-5336072ac52a', False, '03615', '03615', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.978175Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('3277950f-fca8-4bcd-8799-64570ea99bd6', False, '03645', '03645', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.978278Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9bda847a-a23e-430c-bbbe-27da4e67d463', False, 'wax-lake-outlet-at-calumet-la', 'WAX LAKE OUTLET AT CALUMET LA', ST_GeomFromText('POINT(-91.3728 29.6978)',4326), null, null, '2021-03-12T16:10:12.978431Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381590'),
('9b19c51c-b11f-4d74-9eab-2228dc25508b', False, '03750', '03750', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.978545Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('8965cb62-506a-4f0e-b285-48fe8ddfe35e', False, '03780', '03780', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.978723Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2432c199-8996-4ae7-8a30-50c52c9f746e', False, 'thompson-creek-at-chesterfield-sc', 'THOMPSON CREEK AT CHESTERFIELD SC', ST_GeomFromText('POINT(-80.0641 34.7286)',4326), null, null, '2021-03-12T16:10:12.978940Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '02130450'),
('93960641-3b74-456c-8fcf-7aadbf48c6d8', False, '03830', '03830', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.979092Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('1e6a3e9c-3b71-4256-a4b7-75a9615ba8b1', False, '07374525', '07374525', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.979230Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d56eb333-2749-4797-b5c8-f282ecf8fd4a', False, 'northeast-bay-gardene-near-point-a-la-hache-la', 'NORTHEAST BAY GARDENE NEAR POINT-A-LA-HACHE LA', ST_GeomFromText('POINT(-89.606 29.5857)',4326), null, null, '2021-03-12T16:10:12.979362Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07374527'),
('3ba8a61e-8920-46d2-9831-8cd0713508ef', False, 'bartlett-wwtf-near-bartlett-il', 'BARTLETT WWTF NEAR BARTLETT IL', ST_GeomFromText('POINT(-88.1658 41.9669)',4326), null, null, '2021-03-12T16:10:12.979509Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '415801088095700'),
('7e52d2f4-c471-400e-9f1f-8d2f50b4d6e4', False, 'tangipahoa-river-near-amite-la', 'TANGIPAHOA RIVER NEAR AMITE LA', ST_GeomFromText('POINT(-90.4842 30.7289)',4326), null, null, '2021-03-12T16:10:12.979659Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375430'),
('b72a87bd-94c0-48af-b773-d9bdc5419859', False, 'tickfaw-river-at-liverpool-la', 'TICKFAW RIVER AT LIVERPOOL LA', ST_GeomFromText('POINT(-90.6733 30.9306)',4326), null, null, '2021-03-12T16:10:12.979783Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07375800'),
('69d3de52-e320-46e4-8c5b-f29cb8cbb4bb', False, 'barataria-bay-n-of-grand-isle-la', 'BARATARIA BAY N OF GRAND ISLE LA', ST_GeomFromText('POINT(-89.9506 29.4225)',4326), null, null, '2021-03-12T16:10:12.979904Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07380251'),
('b4cd105a-dd1c-485f-bbdb-b80269084676', False, 'caillou-lake-sister-lake-sw-of-dulac-la', 'CAILLOU LAKE (SISTER LAKE) SW OF DULAC LA', ST_GeomFromText('POINT(-90.9211 29.2492)',4326), null, null, '2021-03-12T16:10:12.980056Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381349'),
('d4682aed-6cd6-498c-a610-3d565b03c971', False, 'vermilion-bay-near-cypremort-point-la', 'VERMILION BAY NEAR CYPREMORT POINT LA', ST_GeomFromText('POINT(-91.8803 29.7131)',4326), null, null, '2021-03-12T16:10:12.980177Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07387040'),
('017d8b2c-2234-412e-a88c-d5949ce2fd02', False, 'mississippi-sound-near-grand-pass', 'MISSISSIPPI SOUND NEAR GRAND PASS', ST_GeomFromText('POINT(-89.2503 30.1228)',4326), null, null, '2021-03-12T16:10:12.980297Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '300722089150100'),
('333d17db-11e8-47b9-8621-af2fccce7190', False, '43500', '43500', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.980428Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('53e8a131-3739-42d5-9189-fd5aec548a56', False, '40900', '40900', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.980428Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('1cff71c6-56ec-4baa-8524-ddcc2b292f25', False, '49235', '49235', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.980750Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07381567'),
('afda57df-9971-4d0d-8d7f-7092d94d1791', False, '49255', '49255', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.980833Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('7b94b33a-0922-42ab-9928-2d4dfe40aa22', False, '49355', '49355', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.980931Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d2beb71a-5aeb-4ba3-ae0e-346e81943d97', False, '49365', '49365', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981027Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('704bded7-49c2-43d6-9c4f-9b6b1512b9f1', False, '49400', '49400', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981153Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('707c1f3f-cf12-4768-b826-419478701450', False, '49415', '49415', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981246Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('60a37953-139e-4d66-b000-06f5b16ece51', False, '49542', '49542', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981376Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('908dcb66-f39e-415a-af5f-fa8272a67a42', False, '49570', '49570', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981473Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('5478ea0d-9117-48fb-9694-02515d516579', False, '49615', '49615', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981598Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('59062009-9c5c-413e-9f0a-077fe4d1a015', False, 'mctier-creek-rd-209-near-monetta-sc', 'MCTIER CREEK (RD 209) NEAR MONETTA SC', ST_GeomFromText('POINT(-81.6019 33.7533)',4326), null, null, '2021-03-12T16:10:12.981748Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '02172300'),
('a4b91146-0108-4d0d-8cba-21e9cc6281eb', False, '49725', '49725', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.981950Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4f556761-b926-4086-bdcd-19bcc311ff1e', False, '52280', '52280', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982052Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('4369132f-f2b1-4161-a857-95e6bcc2de59', False, '52330', '52330', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982150Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('76862e2d-87cb-4605-804b-e4d9cb4ce00c', False, '52390', '52390', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982277Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d4198ec8-52ae-4296-a5ce-3763c3165d79', False, '52441', '52441', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982375Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('68eaa243-1fe5-4615-9d87-eb0779971a30', False, '52560', '52560', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982471Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ba4e3eea-f71b-4116-883e-0ab07fada8e4', False, '49630', '49630', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982471Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b3087995-8550-4b0f-a201-dbbf1e52698b', False, '52750', '52750', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.982977Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('47c17779-6fd6-4620-a176-4112878f687a', False, '52800', '52800', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983098Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('081d900a-ef53-442b-940f-d7cee74a26d8', False, '52840', '52840', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983197Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('514648d9-4d3b-4bc4-a57e-edd091b9f3ce', False, '52880', '52880', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983292Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('15101a32-a069-4b0b-a29b-f86a409053b3', False, '52900', '52900', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983387Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('ce4e2f40-708a-434e-9d71-2b441ce8611b', False, '58050', '58050', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983452Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2725f343-d52f-43b2-b41d-0f9f671aa84e', False, '58350', '58350', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983546Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c0b1ecd2-af0a-450c-8a51-f7b44130d612', False, 'coe-lake-pontchartrain-at-frenier-la', '(COE) LAKE PONTCHARTRAIN AT FRENIER LA', ST_GeomFromText('POINT(-90.4214 30.1061)',4326), null, null, '2021-03-12T16:10:12.983687Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073802305'),
('94e20e3a-7c6b-4eae-ac1c-fea8a22e4662', False, '58600', '58600', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983832Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('615d5750-5d23-4dfc-98d1-5fa69b3fc8ca', False, '58700', '58700', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.983903Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('f8f8fd49-cdf8-47ec-bfaa-a406ccbb12c2', False, '61720', '61720', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984001Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('e0806b46-1b66-460e-97d5-0c98bff2edd9', False, '61760', '61760', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984097Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('b7425815-d58a-4422-9907-e115f41c77dd', False, '64050', '64050', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984193Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('6207bc92-5cc7-46ed-a93f-1ff32121a68a', False, 'charenton-drainage-canal-at-baldwin-la', 'CHARENTON DRAINAGE CANAL AT BALDWIN LA', ST_GeomFromText('POINT(-91.5417 29.8231)',4326), null, null, '2021-03-12T16:10:12.984329Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07385790'),
('5f0e95fe-84b1-4dee-9584-408d2b79b4e1', False, 'bayou-teche-w-of-calumet-flood-gate-at-calumet-la', 'BAYOU TECHE W OF CALUMET FLOOD GATE AT CALUMET LA', ST_GeomFromText('POINT(-91.3728 29.7039)',4326), null, null, '2021-03-12T16:10:12.984463Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '07385820'),
('34d6ecd5-a584-4de1-90a6-dacc386988bd', False, 'bayou-teche-e-of-calumet-flood-gate-at-calumet-la', 'BAYOU TECHE E OF CALUMET FLOOD GATE AT CALUMET LA', ST_GeomFromText('POINT(-91.3739 29.7042)',4326), null, null, '2021-03-12T16:10:12.984668Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, '073815945'),
('5f4a3ce2-f359-4f7b-b622-9db903ca8f30', False, '70600', '70600', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984836Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('d0c6b11b-44a4-4041-9cd7-1ff725f042a4', False, '70675', '70675', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984937Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('26ff06dc-5aa1-4928-9595-9fd5713dcf02', False, '70750', '70750', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.984937Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('9c4eec28-485c-405f-93fe-a1671ba1e85e', False, '70900', '70900', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.985270Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', False, '73472', '73472', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.985381Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('77576b15-85c2-42ca-82b8-98cbc8244122', False, '73473', '73473', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.985381Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null),
('c8197081-957a-451d-ab4d-41805e0ecf5f', False, '73550', '73550', ST_GeomFromText('POINT(0.0 0.0)',4326), null, null, '2021-03-12T16:10:12.986013Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', 'e38f3d27-21fe-4e35-8de5-3aea8580e909', '00000000-0000-0000-0000-000000000000', null, null);

--INSERT INSTRUMENT STATUS--
INSERT INTO instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('6575aad0-1ba9-45f7-97c3-7dbf287f00b7', '15dacad3-a7e5-4378-91b4-34208e7039cb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.948076Z'),
('3ce3ffc2-8738-4d07-ab7b-0fdccbd78c0c', '35896e62-b69f-4bc9-9fc6-78ea7c5fa92a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.948289Z'),
('aab82874-bc65-4174-aec3-9b670d7ddf41', 'bbeff60a-9198-4956-9ef2-021fc851cdbd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.948450Z'),
('53a8136b-0de9-442c-b895-c15e65b1a8aa', '64d4fe14-952c-4f0a-b256-1ae56e687fd2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.948600Z'),
('eddb1531-3127-4c14-be64-3c9ed4e491e3', '9e13d62c-b9e4-49eb-aee5-59b4d1bbd46c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.948600Z'),
('daf1c37e-c098-4c14-8813-37555d33c2bb', 'd469455f-6ce9-4faf-b2fc-c321f3365a6f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.949259Z'),
('199b9461-a736-4f21-9d49-755a6e010793', 'b44d2de5-fa95-46fd-9254-93078405d35d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.949259Z'),
('cc99dc2f-8ed4-4cd6-bba8-df4cbe20c009', '41e9219e-2f1f-439f-a724-a3e5c7c7cf73', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950027Z'),
('82a8fd72-ec68-4706-838b-a5b31283d339', '6ee253df-10ae-493c-b307-de8e85464c07', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950027Z'),
('85a18a43-8e38-4b60-858d-1919f3c26c38', '72ca06c0-2710-4135-be3a-bf93d8c51dfd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950027Z'),
('1b4bf1a0-4d78-47bb-8a70-e70d767622c5', '95b436bc-842a-4ad8-83e7-2410a56c439c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950279Z'),
('32ecbd9c-8aa8-4221-a770-c87f0b26e9f9', '016d2287-c4a8-4544-9588-dfb961726c26', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950439Z'),
('52c4b79c-2871-4c5e-9cfa-02407161c44c', '3f1daace-2b8e-4d6c-a21c-1a23359ae00b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950616Z'),
('566c102b-e923-4809-83af-27d2254d5435', 'a757599f-b2e5-4fbe-b40d-8522c3b04af0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.950616Z'),
('26b84d88-cc92-4064-901a-cdefec7f3cf0', '38fa06c8-6457-4523-b4ef-3ebffa520151', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.951160Z'),
('d747034a-eea8-4fcc-bbe7-535a3b0ecef5', 'b9e94249-d8e9-41ac-b683-bde39fa79ed4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.951444Z'),
('fa964a9c-e008-4730-a980-6584f85b12b8', '616fa910-eabd-4718-940f-c7ec2afeffd8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.951621Z'),
('c3d4dbac-f48f-40cd-aa20-e839fdea107e', 'c9e023a7-9cd7-4443-9661-c0f629bc6c8a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.951621Z'),
('f936d6a6-4140-4ef2-a0cd-91a8d5dbf637', '6d781aa1-93e5-4aa5-85fe-9d057ead4247', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952221Z'),
('28a9ed90-89b1-4299-b58d-7d4476a9f7ea', '4adb6c37-4ae1-43e1-9540-e5595bf86d1b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952424Z'),
('aac7d81e-f730-48a6-a704-38b55974fe76', '67cf3e31-f3de-4413-930d-1cf4903c454e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952571Z'),
('f0b9fe11-570f-4f92-a2be-c9c8b0731fb1', 'eafbba9f-6875-462c-a8ec-99e224c37765', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952571Z'),
('ddfae5cf-bc30-453f-9a16-b54c49810e9e', '0b4fb382-4721-4c93-8c92-eb47812e0bac', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952571Z'),
('1ff34031-1970-49ac-aeb8-c0dd177a64b7', 'bacf256e-eb96-4c36-8954-bd92a306f4d8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.952872Z'),
('903f0df7-fe62-46e5-a600-28e62e10d25b', '5524e00c-1f89-418d-9740-6de8645b8333', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953035Z'),
('7c5133d0-3df0-42b3-95cf-a12ac998bf2b', 'd7281dd7-778d-4c3a-90a4-b21bb641e850', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953219Z'),
('254a5773-9965-4f7b-8861-3f727a6d14f5', '17968170-2ff2-4cd8-9809-818f88732879', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953219Z'),
('60b2125c-d208-41cc-80bc-a214f8df5c26', 'cabcd7a9-db9e-415e-ba84-cbe5f1f5c768', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953219Z'),
('ea898048-5adb-4cde-a6cb-40a24348fad3', 'f6e8916f-296a-49e9-aca1-82665748ca62', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953537Z'),
('3ac4aa13-fdb5-4e66-a908-b9b693c899f1', 'ee622c1e-70e8-4620-8f61-2261c0ed13ac', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.953537Z'),
('24a29657-acc5-4827-8eb4-b286c307cb2c', '70ba9edc-40ea-42ad-a8a3-ceabb25b05e2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954216Z'),
('bf346c95-2532-4862-97a4-0f2f688f360f', 'b226f922-c336-4929-8bb1-9e9ddf4b59e6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954216Z'),
('18dc3467-644f-474c-83f2-88e0abf577e9', '7c82c421-5294-40a7-87ff-7af6e42a18a5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954216Z'),
('55ff73c6-78dc-49aa-8351-4b14bc042a64', '694b13bf-449a-4c44-8aa2-c38469c07ed5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954530Z'),
('9048525f-8519-43e3-9bbc-cb244fa2ca8f', 'f15dbad6-ccb7-4a72-9298-5cd33a8b4797', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954530Z'),
('cc6a0227-a447-42d9-9d10-8762eb84a34b', 'c42b64ce-0460-41f3-8bbf-86a8da1d4c00', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954530Z'),
('e1f597e2-b70a-4be9-ba44-49e394cc343d', 'a392c3e8-738d-4feb-afc1-00d064e4298e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954845Z'),
('74977c38-4773-4ac4-b5ee-6189dc066cd9', 'e0ecd779-e73c-47f4-a906-55c8002dcba7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.954845Z'),
('88ebede4-060b-428d-b3cb-a866b4106d6a', '8add8896-c24d-44b3-9853-c189d2d83695', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.955289Z'),
('2451af3c-85ea-4df5-8f17-3ee0d522b893', '1592e55f-ff35-40a9-80c4-7c81e4f620b3', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.955475Z'),
('41c34294-3680-427e-a15b-ad635d1ed412', '4d62ac6e-35b9-41cd-80de-7468d2bad87c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.955475Z'),
('dd0e2687-f764-4352-8271-7395b66fdd1c', '3604b451-24dc-423f-9e8a-f0bbf9d7d8f2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.955929Z'),
('896bd174-4dc8-4a08-a255-aa75b2922cbd', '9d3ea29e-59a3-46dd-a091-d8fed56f03ec', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.956037Z'),
('73e83b1b-ba76-49b8-8c8c-cdc239dc4b49', '4c7ce601-f232-4634-936c-d7604eb54e3b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.956209Z'),
('9150f384-f5a2-4456-ada0-6ecbf3f0b0e9', 'e213714d-2656-4719-84f5-61e208032038', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.956353Z'),
('3ec0056f-3e99-4b43-8c18-49ba4ecb7c02', '965c672f-d3d2-41a9-a537-7626ee1da6d1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.956353Z'),
('b9f21296-99a9-4a84-a9c0-d45c29b80bce', '6bdca601-8044-4830-85e5-7f34dd5d5865', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.957042Z'),
('4c9606e4-afb9-4451-ab5c-61e1568b8045', 'c7e40f37-2390-404f-82fe-c9e7496b8801', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.957042Z'),
('a011722a-0f75-4693-9daa-abc33d83d8b1', '75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.957437Z'),
('6be03791-6c0c-4bd7-859a-1940dc9dfd4b', '59bc5f1d-63a4-455c-80d0-fa6ad8f641fc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.957437Z'),
('87427e07-a3ff-4f56-b4ea-d040be62b36b', '8091654d-2afd-46a6-a2a7-26447b8c0df4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.957812Z'),
('6adf25a4-cd2e-45e2-8b5f-54e6a9bf9ddf', '2cce0e9d-c608-4424-a8f8-0f76e0691338', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.958013Z'),
('0706843c-02d3-42b8-a496-0c07be1f1dcb', '82d2991c-71db-4699-9e7f-5506f9614d58', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.958172Z'),
('3af53879-978c-4a8c-85b7-e10c2dfdbac1', '21bdce69-d2a7-49f0-807b-d29c747f2802', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.958428Z'),
('941535b1-b018-424e-8094-f857af9e25ce', '28bc2a6e-72a9-4e21-9da0-29ce3dda495d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.958696Z'),
('659903b9-b97b-40eb-843b-389bda7efd9f', '00fb2016-699b-45cf-a44f-2944e0860c26', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.958900Z'),
('69f6c552-7ce0-409f-8fec-abd06cb9a4e9', 'df1cf782-4230-4f79-8437-4e361ee67dfc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.959049Z'),
('e98e9118-547a-4a8c-9a39-6eec2e6a841f', 'f0658ffe-4b1f-4b6c-b17b-7b8cc43ad302', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.959245Z'),
('473fa281-4cd2-4f6d-a0f2-8e902db841a7', '581ba9dc-3836-4fea-9598-fc0d9255da47', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.959449Z'),
('4fea1fb0-1aaa-4d71-a8bf-46aa1ad17937', '9ce6e98b-a26b-450e-b859-41bdc05bab06', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.959650Z'),
('9afd89cd-a738-473d-a70d-c658883b53b6', '1cc92566-d593-4c0f-abd5-fd39aabf9fb1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.960032Z'),
('33bc7edf-61e8-4087-89bf-7254fdd32447', '5c3f1f19-1639-477f-81e0-8edf1b85bda7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.960348Z'),
('b5d83da1-2d5e-41c6-8ac0-74c55c1b0f5b', '9a60c8bd-475f-44a9-8e14-7a0c794f0c1c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.960348Z'),
('324447ff-9e08-4d86-9304-82582c8bccf5', '449c9725-605c-4d0d-a0c6-8e434054b26d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.961250Z'),
('d31fc6db-182c-44ab-914f-0f814e3811cd', 'c37feb2a-559b-4f6d-8b3c-75230e5adeae', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.961250Z'),
('b8e960f2-7dc1-4e71-81c5-239db94e291f', 'a70a73d7-47fa-4f92-9b00-6bf9f02c52df', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.961584Z'),
('0bcc8e52-fbf8-4344-ac68-6267c767cded', '3416688d-bd11-4ac3-86c7-f75f2d311498', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.961584Z'),
('63d5fe54-a1af-418a-aa28-f97ab5dbfd04', 'd56068c3-a3c4-4553-b30a-a885570cd7d8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.961964Z'),
('0fd8d469-ac44-41cf-91f5-9ffcf95ca41d', '6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.962150Z'),
('8409363d-852d-4ceb-880e-6e846d33bc40', '5d870e7d-cf49-4b9a-8352-10a75948e282', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.962150Z'),
('d8ae8abf-6da9-49a9-a30b-625f7707a82d', '572ce69b-8b32-40ab-ac00-e783dc6ceb4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.963175Z'),
('7ff55cb4-26ab-4b1b-84d5-424908c78e75', '252a79a9-2538-4016-a9bd-c5edcfb701b8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.963175Z'),
('4b2ae651-863b-4282-9600-a30e41c8523b', 'bcfb5c23-300d-4d78-b8b7-f7cf40bfbd09', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.963531Z'),
('989d4500-7d46-45cf-963a-519efe00b1e7', '0d2b2c4a-e8ab-464d-aa5e-5d4fb3145dc5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.963904Z'),
('9079b4e4-c362-44fd-9df1-20056785dd82', 'd0803aa7-bbee-48ed-b23f-7dea74b2ac59', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.964324Z'),
('8a560b92-c50f-44c9-86d8-6d720ab782a1', '40f1872b-b305-44ab-bc8d-06351fd29d3a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.964324Z'),
('151b8a28-30ab-4ed1-818d-e8108a14bb69', '5e19aa7a-1611-49e9-bec3-75c3421c7b5e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965240Z'),
('5657db92-b3c2-40bd-90ab-fb22526e5723', 'e7665d40-d030-4b35-b91a-741546df5fb0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965240Z'),
('37fff46c-2e9e-4c72-945b-581c718c3469', 'e1993ab5-ec08-4e07-bc3c-c4db92731d1c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965538Z'),
('9e9ddc4b-4a93-4ef2-a154-9290d4e79ebe', '4429e603-e648-4a97-a2ce-42d1f94347fd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965538Z'),
('e983da56-e711-4e52-b505-1c3cca86de2b', '2db45595-fefc-4eb8-9237-8fc72561b47e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965946Z'),
('c8115ac9-9360-4d9f-a7a1-29f0a8e01ea3', '19c0ac2f-09ab-4c68-b509-71d8a3c07a99', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.965946Z'),
('f208e66c-af1b-4579-a35d-fb55183de5fb', 'ec40c087-d75b-4580-9ee5-337eeff32f95', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.966295Z'),
('41da6de3-852e-4dc6-95c5-9f0f1d5adb4f', '4009c409-974c-46f8-86af-992ee43d5ac7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.966295Z'),
('a32847be-e95e-4d94-8561-f487da32024e', 'f3b99e3b-1856-4226-a73c-1c46b200fb66', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.966596Z'),
('01f24f18-b9ad-45ac-bddf-4ad2dc231d19', 'f58051b1-180a-4c5d-a602-e4dbaaac3367', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.966596Z'),
('9a3fd868-f722-46f2-ab5c-a7295df1b9fa', '72d33561-aaaa-4e90-91b4-eb73eaeaaac6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.966959Z'),
('43e681f4-0e9a-4ca9-b510-31a3a515a620', '6b78716c-4975-46e4-b635-9ff79e7f0627', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.967531Z'),
('1e75b07c-34e1-4d98-b051-6573250774a4', 'a88ece8d-d3ba-4a77-a413-5dd03ebfb623', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.967792Z'),
('a431910d-4b57-4ce5-970e-2abfc0ab3aeb', 'f6629ae2-98e2-4e06-b685-331189cff3a7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968003Z'),
('941dda1c-786b-4f5f-86b9-870eda6c1a4a', 'eb4969b9-30ae-45ea-ad68-d19948a84020', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968003Z'),
('923aa590-a7ca-466f-b474-9119300082f7', '59db968a-43d1-4590-b4f7-6f7f6423e735', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968359Z'),
('8ef1c694-c396-4d34-9f70-0c5b940b3dc1', 'ac71a2f1-74bb-49de-8dd4-d2b15b974978', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968467Z'),
('5bdd5d26-771a-4d0d-85b3-bc60b9695e44', 'ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968642Z'),
('7a2d2a11-e668-4241-a715-f23fd014d45b', '4916c4af-8019-4492-a734-fc3ff8c49074', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968781Z'),
('fcb43b83-0211-487a-bc4c-943f461012e1', '5dfc826b-a94c-4cee-b057-6fb858afc95d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.968932Z'),
('42dc5ed2-496a-4809-865e-ed58caef9231', '7e772c75-64b6-485e-8294-d1d3bc6da0fe', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.969013Z'),
('b497e7b8-5c74-4911-990e-d58d9e4e3937', 'c4c87cf8-a155-4abe-bf89-4f7af607db31', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.969105Z'),
('96cbd5ab-bc12-4387-8269-5e0f9a31888a', '3fcc1834-0263-4dc5-a7c8-9157c767bb4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.969820Z'),
('5c6c6695-4cb4-4b47-833c-b3dbca2b5496', '6b5d1a5d-0861-45a1-b56c-ec37c501e993', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.969954Z'),
('ae70b88b-62b7-4de8-ab6c-27f02530268b', '6890212a-c7f1-499e-a3fe-e7d973eb7cc1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.970341Z'),
('ba69002a-88cc-4985-9383-ce12a3871290', 'bb1389a0-bfad-430d-a66d-2f041a738cb1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.970452Z'),
('ffff7844-26a9-4e8f-865f-fb2ed639aee1', '38ea9864-0181-475e-a8c7-b34edbf31b6b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.970593Z'),
('58e251e3-33f1-4cfb-aae6-f6a2280d3997', 'b417e9f4-ce46-47a7-8bd2-0b26506395e4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971045Z'),
('318ba6a6-ca2f-4dc7-b726-d3ed51cafbec', '04e220cf-128b-4115-8a7b-76ac2625a81b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971160Z'),
('e7fef8a2-1774-4a37-864d-c2e3b796341d', '7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971257Z'),
('1adb50af-8b5d-4c33-9669-4cb702fbb3b8', '892ca8ff-08f2-4e1b-af40-32a77baab497', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971452Z'),
('60c546ca-6032-4d44-8f9f-4543f47bc444', '8250ac38-3d9f-4039-9f08-149912301977', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971601Z'),
('6cf81f08-8ac5-4b53-9217-4203bc0fa5c4', '9dcf334f-2982-49c9-941c-38c91565f019', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971709Z'),
('3854ceeb-35b3-4b59-88f4-182be2b08359', 'df1b2707-3dea-4b91-be0d-2f64f21544e2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971900Z'),
('c584c7dd-fe63-4068-82a5-5c3a505a8082', 'fcf546da-e8c6-471b-b894-98a691c71251', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.971970Z'),
('28d06558-b7c8-47a8-ad8c-0b1c90268fbc', '1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.972095Z'),
('9ee1dbaa-a592-4fe6-8f8f-41a511fe3100', '3eba9840-cebb-45ee-8e9f-a62d8279c86a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.972249Z'),
('577e9bf4-0b3c-4319-a702-4267d235a907', 'bc7d2d9a-8e02-4a78-b838-cadd444ed756', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.972384Z'),
('b744ec16-8bb0-418c-9b12-f3b4bd802e8f', '3a43504e-d72c-43e0-9849-d67e016fbc7a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.972384Z'),
('4d720fe5-5cca-4662-b0ba-e04910737d6e', '3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.972880Z'),
('69d1eaf5-c7d9-4ce6-baeb-f2069c070367', '5171ef3b-9ef9-4984-8ab7-57b76db13621', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973019Z'),
('ccadc953-86cd-4a5c-9592-ddeb58552e1b', 'd02cd42c-3cf0-4b41-a670-1d35708b97cd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973019Z'),
('9fec5ef0-da13-4f50-947c-87704d5dfa0a', '7c317f0e-f87f-482e-9776-e6c8ee96170a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973019Z'),
('1c7dfd23-1ecd-4748-bcb5-e1992aa74c3e', 'b72006fc-5809-4985-8f3c-45baed7cfcb9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973681Z'),
('c420d538-d1af-48c1-bb90-b0b6f1228d92', '4811e8d0-4c72-4c86-8bec-37e491f5ab63', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973802Z'),
('9f1e5e3e-2b0e-4623-8bb2-3fb97b1daec9', '3f9a0fcf-4e4f-4b0c-ac88-011455d6957f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.973944Z'),
('7e1cb663-a608-467b-8188-cf054248a7a0', '553f696a-9292-41b1-a4bf-67236068a1d6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974052Z'),
('46330728-5966-42e8-8472-c7b88d7d5a46', '394d6441-43b2-4745-b97d-c6462c7bc908', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974176Z'),
('dc0abdd9-c467-4d5f-8f7d-7d580f31a6d6', '2711323d-5ed2-44b5-a7a6-f0614c0b422b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974426Z'),
('6fcf2569-20a7-47fd-a4fa-34f794058d52', 'b04e2a92-11d5-416a-a50f-d91a161456bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974426Z'),
('59361f46-6760-4f40-bdd2-520f06a61c03', '3838e35a-efaa-493a-bd48-52561971ce52', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974426Z'),
('2dcb12db-f0b1-4fde-b7a7-aea97c0035b5', '3bd76812-033f-4367-b8c4-eceba262f393', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974687Z'),
('e24f2d0e-c570-4b5d-a519-74d6b86a6c62', '0eb3be0e-dd34-4e25-bf96-c5443b2cb400', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974687Z'),
('7b490a0e-bf69-44fa-b2b5-b4b9a28828c5', '89cdff86-ab37-4149-aa95-5f6e43454b39', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974687Z'),
('9216156a-110d-4e91-bf31-a99717884184', 'f8d6ccd6-fb48-4228-b777-e423612b1f27', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974947Z'),
('7dfc2ea8-3578-445c-81fc-e18ccb59714d', 'cf64b9de-4b12-462e-955c-669ca84e2716', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974947Z'),
('be0e66f8-4ebe-4e6a-9ac7-7065449d7541', '6c4ee0c2-6972-4e69-9555-ff05f155ad5d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.974947Z'),
('c00bd941-44a1-40a8-b28e-fbbd175c0ed1', '7babc4e1-950e-41e8-874f-49559dd947bf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975208Z'),
('24cbccbc-f13e-4744-bcfd-caf18f692f11', 'ab7ee39b-8397-484a-ac29-e2e813ed6adb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975345Z'),
('5d5b92f6-8d7c-48a9-af00-b1d7f4c42646', 'd93718c4-08ae-4065-8855-674c5206c1ed', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975574Z'),
('31d2f6bf-4f7b-4dc4-a0f6-1925490e8da6', 'de20f637-d3c7-4ea4-8ca6-604f6a127532', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975677Z'),
('ad30313a-c785-4e9b-a865-254b7d1baa59', '8a7739c0-ccae-4f3b-83d7-e2f6f506e6da', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975800Z'),
('a4be06a5-41d3-4fd3-9f36-052bb4aeb4d3', '2006cc18-de31-4ae0-8aaa-81ba6cf1ba95', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975863Z'),
('5f52bcc9-46b4-44bd-83de-3f9f670dccb1', 'e29fcc6e-a0a8-4fe0-b450-099142997846', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.975954Z'),
('33fd1683-503e-46be-a4d7-66f70c85733e', '7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976045Z'),
('f69fee77-8798-4c7f-96ee-9563d7d5fb91', 'bcc35574-bf23-4952-9760-049ba8843bd7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976167Z'),
('978fa5a3-8527-4631-9527-0ad4ecd49a54', '4bbb3d28-b816-4f71-98c1-de647d8f5e38', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976293Z'),
('6045cc19-f80e-4fd2-a0d0-c1709604f3b3', 'e0e57af0-7523-4bb1-b327-eff3dca3e508', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976293Z'),
('414a08af-2653-4479-83cd-d26314fb598b', 'b3524459-4cbe-4802-835e-b415cc6657b2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976487Z'),
('00a16548-757b-4eb8-ba80-7ef401f2bc88', '5d20b46f-1fbe-48c2-875e-1894f4aeb348', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976487Z'),
('8616ae66-c317-4768-8e96-9a985aa96abf', 'ed2c1d1c-7bb2-4374-8e90-878ab6b70d9c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976771Z'),
('2d94abaf-b853-402c-b975-1e4a9e97b98a', 'eaa406e4-560c-4fa3-9c28-8f265999f17b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.976965Z'),
('38492de5-5f14-42cd-8b07-24b6b87bc464', 'bd3d7701-94b0-4390-b2f8-40e03c48508c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977098Z'),
('85bcc192-00d0-4c19-866b-488f716282f0', '8e5ed58a-febf-4b91-86e3-4b39d798a592', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977321Z'),
('709db487-a7cf-4a67-a308-4d372ea59132', '33564f35-da99-4db5-9d9d-2c96698f99e6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977451Z'),
('8d6a1514-c745-4c23-9e80-210c99675fa3', '228d8881-b162-4842-92c9-75422535a5ad', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977561Z'),
('ef9a6fc8-e835-4af9-8dd0-66eb08fe31c3', '8f32236d-1e18-47f4-bf36-383eb8c82fbe', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977655Z'),
('7e233a28-b5ff-4040-b5d1-76a010983211', 'dccf96c3-4dff-4b3b-a7b7-446a97db654b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977749Z'),
('cf10ae3f-eabb-4a54-afe6-2a9915e687ca', '8527c379-3f95-4ef6-aae7-4fcf94ad8691', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977842Z'),
('ac146f39-42e0-4f7e-a01b-0939d7300c5b', '8bd798a4-1c7b-4684-ae3b-eb75a62a09cf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977950Z'),
('4d9aca18-fe0d-431f-8a50-374e98221022', '5fbab803-eb25-480c-9d30-1064b32d9f87', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.977950Z'),
('4a40c32c-a0cd-4923-ace6-dec5f35eef21', '2eed3135-b0ee-423d-b8a6-5336072ac52a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978175Z'),
('c2a91202-b554-4161-8c2e-1f2ca965e72b', '3277950f-fca8-4bcd-8799-64570ea99bd6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978278Z'),
('69c51a49-01b4-42f7-8361-adc381e4dca2', '9bda847a-a23e-430c-bbbe-27da4e67d463', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978431Z'),
('4bfca19a-6715-4c35-a48f-f457b44b1e22', '9b19c51c-b11f-4d74-9eab-2228dc25508b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978545Z'),
('b2db1651-6f17-466f-9004-62ae417064f0', '8965cb62-506a-4f0e-b285-48fe8ddfe35e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978723Z'),
('6dd26849-09d9-4608-b673-471f4671714a', '2432c199-8996-4ae7-8a30-50c52c9f746e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.978940Z'),
('12406600-2e3d-4278-aaf5-d1d00d9c7822', '93960641-3b74-456c-8fcf-7aadbf48c6d8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979092Z'),
('c0a01059-21b6-4811-8d18-3df79e264055', '1e6a3e9c-3b71-4256-a4b7-75a9615ba8b1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979230Z'),
('7e355e21-173a-4aaf-8066-f215384a49d1', 'd56eb333-2749-4797-b5c8-f282ecf8fd4a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979362Z'),
('226070ed-4d17-4bd3-96ab-c10c50366783', '3ba8a61e-8920-46d2-9831-8cd0713508ef', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979509Z'),
('b646d28f-2ea6-4050-96a1-70cb9062d31e', '7e52d2f4-c471-400e-9f1f-8d2f50b4d6e4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979659Z'),
('e728c581-bc22-4a54-8d24-620183a82c7a', 'b72a87bd-94c0-48af-b773-d9bdc5419859', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979783Z'),
('a5c335fb-2286-42b9-b4fc-ff821e060158', '69d3de52-e320-46e4-8c5b-f29cb8cbb4bb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.979904Z'),
('23a2e8b4-0f65-4225-877c-3eeaa52e3c82', 'b4cd105a-dd1c-485f-bbdb-b80269084676', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980056Z'),
('add56905-dea7-4d9c-beb2-0a5bf2d8a49c', 'd4682aed-6cd6-498c-a610-3d565b03c971', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980177Z'),
('48c1d146-ab02-4066-80ce-4a277da2a6ae', '017d8b2c-2234-412e-a88c-d5949ce2fd02', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980297Z'),
('9419f868-9e0f-4679-9fa0-414c087e198f', '333d17db-11e8-47b9-8621-af2fccce7190', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980428Z'),
('7e5d9263-37f9-421a-9395-41406e47ce71', '53e8a131-3739-42d5-9189-fd5aec548a56', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980428Z'),
('87cd2fae-89da-4061-939e-1dc4688e8ade', '1cff71c6-56ec-4baa-8524-ddcc2b292f25', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980750Z'),
('14526309-38e7-478f-a544-47735aaaa20d', 'afda57df-9971-4d0d-8d7f-7092d94d1791', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980833Z'),
('2bdda7d4-a848-4e13-b26a-1bd08e204cc9', '7b94b33a-0922-42ab-9928-2d4dfe40aa22', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.980931Z'),
('18a153e8-6da9-4173-a399-b3c504e3b696', 'd2beb71a-5aeb-4ba3-ae0e-346e81943d97', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981027Z'),
('c31356c0-14ec-46c9-a724-f5b7d8269dfe', '704bded7-49c2-43d6-9c4f-9b6b1512b9f1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981153Z'),
('0477947a-0216-4816-be69-76413990e0a7', '707c1f3f-cf12-4768-b826-419478701450', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981246Z'),
('e8ed0f23-c807-464c-bc2b-fe6a816c0fdd', '60a37953-139e-4d66-b000-06f5b16ece51', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981376Z'),
('08196120-d108-490c-9a57-58680c46c6a6', '908dcb66-f39e-415a-af5f-fa8272a67a42', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981473Z'),
('74a43fe6-5d45-4706-afc1-bd2b94e837a5', '5478ea0d-9117-48fb-9694-02515d516579', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981598Z'),
('467124a9-10e3-4bb0-bcc4-b3faf519b291', '59062009-9c5c-413e-9f0a-077fe4d1a015', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981748Z'),
('c1291813-7d90-448e-b08f-422d19d9dd6b', 'a4b91146-0108-4d0d-8cba-21e9cc6281eb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.981950Z'),
('a82e64f7-8e76-4816-a85c-9d7b7ec97feb', '4f556761-b926-4086-bdcd-19bcc311ff1e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982052Z'),
('032bfc49-0107-43a4-adbc-e887e1ee735a', '4369132f-f2b1-4161-a857-95e6bcc2de59', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982150Z'),
('a6c1a817-6156-4b4b-b16c-b95224f7cede', '76862e2d-87cb-4605-804b-e4d9cb4ce00c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982277Z'),
('b0793449-2187-4ee5-a278-d1800f33b974', 'd4198ec8-52ae-4296-a5ce-3763c3165d79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982375Z'),
('6b4607dd-a778-4e17-8e6e-d11b2fca87f0', '68eaa243-1fe5-4615-9d87-eb0779971a30', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982471Z'),
('7c78acff-3a1d-47f4-8020-0f2d7f62d1db', 'ba4e3eea-f71b-4116-883e-0ab07fada8e4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982471Z'),
('4ebf8bb1-7c8d-4a3b-a550-160014cfee3e', 'b3087995-8550-4b0f-a201-dbbf1e52698b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.982977Z'),
('efbee740-8954-4e2f-b4d3-a69152c1d97b', '47c17779-6fd6-4620-a176-4112878f687a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983098Z'),
('f0a00863-94f8-4dd2-8aa9-b511eb46c458', '081d900a-ef53-442b-940f-d7cee74a26d8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983197Z'),
('3616f37b-05bd-49e3-ba17-e982bbadbdbd', '514648d9-4d3b-4bc4-a57e-edd091b9f3ce', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983292Z'),
('a1b9b193-9060-4931-928f-d017ac9bee13', '15101a32-a069-4b0b-a29b-f86a409053b3', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983387Z'),
('fcc6ee10-24a7-4d64-8fc7-d28263521aa2', 'ce4e2f40-708a-434e-9d71-2b441ce8611b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983452Z'),
('1df2f0ea-9274-48e0-acae-401174be5335', '2725f343-d52f-43b2-b41d-0f9f671aa84e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983546Z'),
('35240ce6-e095-46de-a64e-077141085af2', 'c0b1ecd2-af0a-450c-8a51-f7b44130d612', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983687Z'),
('c4c02066-3bdb-43de-9e7d-422083e18d31', '94e20e3a-7c6b-4eae-ac1c-fea8a22e4662', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983832Z'),
('60e1e5f0-ff5f-4095-980c-89c1c302e114', '615d5750-5d23-4dfc-98d1-5fa69b3fc8ca', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.983903Z'),
('67c65a39-aede-4ec3-8dca-185a7e69d8d3', 'f8f8fd49-cdf8-47ec-bfaa-a406ccbb12c2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984001Z'),
('01ab86e0-c1b0-48a6-b06d-eea39c3e2d16', 'e0806b46-1b66-460e-97d5-0c98bff2edd9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984097Z'),
('500a48a2-ed55-426a-b5e9-bfc79a57f3b6', 'b7425815-d58a-4422-9907-e115f41c77dd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984193Z'),
('1d2341f4-b57f-4441-8db6-1c79f59f0b06', '6207bc92-5cc7-46ed-a93f-1ff32121a68a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984329Z'),
('0c7b1982-9854-4c02-8844-4ec86f57dbe4', '5f0e95fe-84b1-4dee-9584-408d2b79b4e1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984463Z'),
('5cd19f3d-2272-4d15-bd54-aee304534776', '34d6ecd5-a584-4de1-90a6-dacc386988bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984668Z'),
('f38043da-e645-4e0f-83cb-7b5969105142', '5f4a3ce2-f359-4f7b-b622-9db903ca8f30', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984836Z'),
('30590c6c-3413-4538-a518-62e512aaebf5', 'd0c6b11b-44a4-4041-9cd7-1ff725f042a4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984937Z'),
('a0cdb9b3-f6e2-4c2c-b42d-d8afa9070112', '26ff06dc-5aa1-4928-9595-9fd5713dcf02', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.984937Z'),
('4ce77f05-0c8d-4b9c-b612-59c608112fa2', '9c4eec28-485c-405f-93fe-a1671ba1e85e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.985270Z'),
('99c9b295-e6ca-43d6-9f2f-171913202868', '2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.985381Z'),
('e88eb093-3e8f-462d-a60f-0bebab80437b', '77576b15-85c2-42ca-82b8-98cbc8244122', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.985381Z'),
('b41879d8-54da-4e2b-902c-cbd6292117bb', 'c8197081-957a-451d-ab4d-41805e0ecf5f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:10:12.986013Z');

--INSERT TELEMETRY_GOES--COUNT:159
INSERT INTO telemetry_goes (id, nesdis_id) select '03d73df2-3db2-4003-95e5-7bce29a52b0e', 'CE672AEA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE672AEA');
INSERT INTO telemetry_goes (id, nesdis_id) select '6edd26bd-bb45-46a5-93ec-1d9bb87fd52f', 'CE5AD1AA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5AD1AA');
INSERT INTO telemetry_goes (id, nesdis_id) select '79412361-14d1-4b65-b5bb-eb2338539e20', '1783C642' where not exists (select 1 from telemetry_goes where nesdis_id = '1783C642');
INSERT INTO telemetry_goes (id, nesdis_id) select '8887b5c8-de97-4c32-b63e-58b0ef975aaa', 'CE6A94D0' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6A94D0');
INSERT INTO telemetry_goes (id, nesdis_id) select '1ba3a311-a34b-47e8-8103-9e1eac4e2ce1', 'CE96C724' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96C724');
INSERT INTO telemetry_goes (id, nesdis_id) select '08e43040-2a98-40fd-98c0-699384f5f48b', 'DE6C3362' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE6C3362');
INSERT INTO telemetry_goes (id, nesdis_id) select '712ed4ff-dc49-45a5-a0e2-1487c4581f8a', 'CE3D325E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D325E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f8cfd6bb-7a6c-4d66-9983-82cf388728e4', 'CE61F29E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61F29E');
INSERT INTO telemetry_goes (id, nesdis_id) select '38847c6d-58af-4848-bee3-fde7e7e4f40a', 'CE03A524' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE03A524');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f15b511b-5e42-46f3-b7cf-a8e112c92191', 'CE61EF3A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61EF3A');
INSERT INTO telemetry_goes (id, nesdis_id) select '6a629e80-0fdb-4418-afd9-58c20cb289a0', 'CE61C9D6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61C9D6');
INSERT INTO telemetry_goes (id, nesdis_id) select 'bd0507cb-ef3c-45df-a129-961463a7f22b', 'CE3D2128' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D2128');
INSERT INTO telemetry_goes (id, nesdis_id) select '5c4da467-cf0f-444e-9b8a-3a7398a02c25', 'CE670C06' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE670C06');
INSERT INTO telemetry_goes (id, nesdis_id) select '48196b78-149e-449a-a623-791d18496637', '168341F4' where not exists (select 1 from telemetry_goes where nesdis_id = '168341F4');
INSERT INTO telemetry_goes (id, nesdis_id) select '9d60b511-b8cc-41ac-99ef-048e01e889d2', 'CE96842E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96842E');
INSERT INTO telemetry_goes (id, nesdis_id) select '02dda7b8-e77a-4ee8-981f-109310db9602', 'CE5F4308' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5F4308');
INSERT INTO telemetry_goes (id, nesdis_id) select '73956d33-2c35-41df-81aa-1d5436f06df3', '17CC20FE' where not exists (select 1 from telemetry_goes where nesdis_id = '17CC20FE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f2ee2a9a-3e24-415c-a339-875f6c42bd52', 'CE96E1C8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96E1C8');
INSERT INTO telemetry_goes (id, nesdis_id) select '458b72bc-c2d2-4ab9-a600-4e7105015ca2', 'CE61DAA0' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61DAA0');
INSERT INTO telemetry_goes (id, nesdis_id) select 'a5c821c6-3a1a-4dd1-8028-d05cf9294e25', 'CE9713B6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9713B6');
INSERT INTO telemetry_goes (id, nesdis_id) select 'c79f68db-ef3d-4b00-a4de-0ad99ce0e20c', 'CE96B1B4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96B1B4');
INSERT INTO telemetry_goes (id, nesdis_id) select 'acd11117-6a70-4e3f-9add-763d0e2c0720', 'CE65CCEC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65CCEC');
INSERT INTO telemetry_goes (id, nesdis_id) select '04d2c33d-331a-4141-ac31-080a3a877d59', 'CE65D148' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65D148');
INSERT INTO telemetry_goes (id, nesdis_id) select '22791f7f-8f05-42ff-9ad2-f3cd3672074b', 'CE0383C8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0383C8');
INSERT INTO telemetry_goes (id, nesdis_id) select '7ddcbd7b-8571-4a5f-9b2b-f5eada8cb5c6', 'CE3D57B8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D57B8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e83b9db7-488b-47fc-bd14-935ae625556c', 'CE43BB5C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE43BB5C');
INSERT INTO telemetry_goes (id, nesdis_id) select '5dcf465d-70bf-4163-902c-c128bbf10ad3', 'CE3D14B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D14B2');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd45c5dca-a28f-4bdc-bf73-b5bf5a93bf02', 'CE06A2E4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE06A2E4');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fc73977c-e4da-4190-98f3-12134d2b78e8', 'CE03B652' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE03B652');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e458550d-eb9e-47b7-90c0-cc8e7b3a6ee5', 'CE0221CA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0221CA');
INSERT INTO telemetry_goes (id, nesdis_id) select '059e3288-618f-4221-863b-cc46b7aa1295', 'CE3D44CE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D44CE');
INSERT INTO telemetry_goes (id, nesdis_id) select '5542dad4-dc5a-4b72-8178-d921282ebcf2', '17CBC138' where not exists (select 1 from telemetry_goes where nesdis_id = '17CBC138');
INSERT INTO telemetry_goes (id, nesdis_id) select '5c422eac-3dcf-4416-b5ff-3575beb673bf', 'CE7E6822' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7E6822');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ddddf336-7c67-447c-bd41-fe14f13b0333', '1783963E' where not exists (select 1 from telemetry_goes where nesdis_id = '1783963E');
INSERT INTO telemetry_goes (id, nesdis_id) select '883be004-c07c-47d3-aca9-e44afa3c316b', '172AC648' where not exists (select 1 from telemetry_goes where nesdis_id = '172AC648');
INSERT INTO telemetry_goes (id, nesdis_id) select '9b75e35f-ba64-42aa-95d9-602f5db81268', '17CA704C' where not exists (select 1 from telemetry_goes where nesdis_id = '17CA704C');
INSERT INTO telemetry_goes (id, nesdis_id) select '8bdb09dc-9886-4101-ac1f-d414644b9ec4', 'CE5F88C4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5F88C4');
INSERT INTO telemetry_goes (id, nesdis_id) select '87296849-5609-4114-8761-48904cc647c1', '172AB0D8' where not exists (select 1 from telemetry_goes where nesdis_id = '172AB0D8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e583fe22-47a2-44d8-bdfd-eee0dd0ad05c', 'CE976526' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE976526');
INSERT INTO telemetry_goes (id, nesdis_id) select '3e489698-2d15-4264-b484-ec3883959f9e', 'CE9750BC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9750BC');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ab95547c-07e2-45fd-9917-345ec4f0fe48', 'CE96A2C2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96A2C2');
INSERT INTO telemetry_goes (id, nesdis_id) select '12952bdc-8f8e-43be-a619-4e70c6005c51', 'CE977650' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE977650');
INSERT INTO telemetry_goes (id, nesdis_id) select '3cda0e26-9d89-45bb-960c-ed4ce3f22893', 'CE9795A2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9795A2');
INSERT INTO telemetry_goes (id, nesdis_id) select '61d18727-7980-4218-bc9a-7459c92e5fc6', 'CE97A038' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97A038');
INSERT INTO telemetry_goes (id, nesdis_id) select '6a10892d-82e9-4e14-9977-c096c2cf30ac', 'CE020726' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE020726');
INSERT INTO telemetry_goes (id, nesdis_id) select '608de20c-bfa0-4ac6-a9ec-3bbc3b9ee7e3', 'CE97B34E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97B34E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f6c85772-21a5-40ab-bc23-65d606facd26', 'CE97C5DE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97C5DE');
INSERT INTO telemetry_goes (id, nesdis_id) select '9aaf8a00-a76c-4505-89db-2f6a3017f9d5', 'CE97D6A8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97D6A8');
INSERT INTO telemetry_goes (id, nesdis_id) select '4fffe379-5431-4069-8185-9b4716a802a3', 'CE97E332' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97E332');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f833d903-8b98-4071-838c-86a71acb43a9', 'CE97F044' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97F044');
INSERT INTO telemetry_goes (id, nesdis_id) select '822b81ea-967f-4bc8-99a1-1e87776b8a53', 'CE980652' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE980652');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e9f318a5-5200-495f-9c4a-50849474d05d', 'CE981524' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE981524');
INSERT INTO telemetry_goes (id, nesdis_id) select '6dcc781b-afda-40d3-89ea-ddfbf658dde0', 'CE9820BE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9820BE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'bbea5db8-f4ad-41d0-919b-b1ba78361763', 'CE659C90' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE659C90');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f4a69857-87c1-4170-8f21-4fbb5b238d8c', 'CE969758' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE969758');
INSERT INTO telemetry_goes (id, nesdis_id) select '9436171f-b2a5-4ab0-9816-159170e5dec4', 'CE96D452' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96D452');
INSERT INTO telemetry_goes (id, nesdis_id) select '34c74351-2122-448b-89c1-cb321cc50876', 'CE9743CA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9743CA');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd2c00c91-b9d8-46e0-8a71-b1195bfb5cd6', 'CE0390BE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0390BE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ac4d0beb-88c0-4f31-b613-486bfa08d5d3', 'CE41D59C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41D59C');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e65f9673-aede-4199-99b3-2a94f87fe771', 'CE9667DC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9667DC');
INSERT INTO telemetry_goes (id, nesdis_id) select '4c5fdc7f-0545-44df-9785-d625d72c60dd', 'DDAB972C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDAB972C');
INSERT INTO telemetry_goes (id, nesdis_id) select 'b06cee02-eb20-474e-a861-5cf42e03a396', 'DD17D472' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD17D472');
INSERT INTO telemetry_goes (id, nesdis_id) select '43a7ed8d-8f4b-4284-b0c2-324c93418644', 'CE41ADDE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41ADDE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e5b47b00-45be-4c16-a590-e8b9df8c8d18', 'CE40EC2E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE40EC2E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fefd961a-a953-492c-af37-e736cc5bfced', '16EF0150' where not exists (select 1 from telemetry_goes where nesdis_id = '16EF0150');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd0038cdb-aed7-4aac-b043-961d84923e2c', 'CE4027E2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE4027E2');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ae215b11-24d5-4a5b-9819-8f41c91b3439', '16283C48' where not exists (select 1 from telemetry_goes where nesdis_id = '16283C48');
INSERT INTO telemetry_goes (id, nesdis_id) select '8cd3406b-a43d-4080-9850-e10d5e74d7b4', 'CE40F18A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE40F18A');
INSERT INTO telemetry_goes (id, nesdis_id) select '78fd0eb2-ff24-4db8-90a9-6380c8e595ce', 'DD2F00CC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD2F00CC');
INSERT INTO telemetry_goes (id, nesdis_id) select '761c9eb2-a9e5-4b45-822f-379a3cc19d5b', 'CE4103F4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE4103F4');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd34c00b4-db45-43ec-a20d-58d6705cd240', '170B76D2' where not exists (select 1 from telemetry_goes where nesdis_id = '170B76D2');
INSERT INTO telemetry_goes (id, nesdis_id) select 'dfc714e0-c5aa-4d6d-b92c-bc49ee8e9a9c', 'CE40FF58' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE40FF58');
INSERT INTO telemetry_goes (id, nesdis_id) select '839aeca4-d480-4cb8-a1da-9bda8e0d00ad', 'CE7F948E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7F948E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'a1081cd1-23a5-4e5f-87b8-4d3523760f8c', 'CE411082' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE411082');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ca7bc2a4-5283-457a-8382-b3d45c091dbc', 'CE658FE6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE658FE6');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f58f0217-a629-46d8-853b-79baa9f6604b', 'CE5F1374' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5F1374');
INSERT INTO telemetry_goes (id, nesdis_id) select '96a9674d-4f64-4552-9b5f-ad353d68bc04', 'CE97355A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97355A');
INSERT INTO telemetry_goes (id, nesdis_id) select '60d9864b-8140-4527-871b-28c015f36cf3', 'CE65C23E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65C23E');
INSERT INTO telemetry_goes (id, nesdis_id) select '6a65e3a2-2af6-4aa0-9cfb-ab8801507e28', 'CE41C6EA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41C6EA');
INSERT INTO telemetry_goes (id, nesdis_id) select '4d809aeb-d269-468d-8c90-b0f62bafb8a8', 'CE45705E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE45705E');
INSERT INTO telemetry_goes (id, nesdis_id) select '4f1f6779-d931-45f7-ac57-88a7d158816b', 'CE41EED4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41EED4');
INSERT INTO telemetry_goes (id, nesdis_id) select '6bc93d1a-fa95-45b8-81ea-bdd7b5ad8ba6', 'CE412518' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE412518');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ac681a2e-2577-4e1b-9628-97fe3e4cbf70', 'CE96F2BE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE96F2BE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f186e2c0-295b-4bb5-aa65-8245c9bb3c78', 'CE9700C0' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9700C0');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd4359bc1-0563-4ec2-a1b8-b45c75b3b393', 'CE9674AA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE9674AA');
INSERT INTO telemetry_goes (id, nesdis_id) select '8c44b370-114b-4e09-8634-8fc4ad5ee599', 'CE3D6222' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3D6222');
INSERT INTO telemetry_goes (id, nesdis_id) select '57c8c07b-621d-4cbb-8efc-857f0aeb0b45', 'CE965246' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE965246');
INSERT INTO telemetry_goes (id, nesdis_id) select '7e96b629-0c92-4fe0-b096-523e24d9664c', 'CE41F370' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41F370');
INSERT INTO telemetry_goes (id, nesdis_id) select '8280848b-f4e6-4ab6-99c9-6070e802f81a', 'CE964130' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE964130');
INSERT INTO telemetry_goes (id, nesdis_id) select '1a7a69cc-4b69-4350-9d7d-0aeece773a0c', 'CE61E1E8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61E1E8');
INSERT INTO telemetry_goes (id, nesdis_id) select '7d59b71e-79eb-4161-9cd8-17cfe78352b5', 'CE65A90A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65A90A');
INSERT INTO telemetry_goes (id, nesdis_id) select '2169217d-63b9-437f-96df-d8717ece6f5c', 'CE65B4AE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65B4AE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'c71b5df9-9544-4dd5-803a-86356e44434a', 'CE410D26' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE410D26');
INSERT INTO telemetry_goes (id, nesdis_id) select '709e06c8-f35c-447b-a8f3-e478c5777239', 'CE97262C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE97262C');
INSERT INTO telemetry_goes (id, nesdis_id) select 'de10d184-a255-42c4-befc-d1a36b0e8704', 'CE659242' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE659242');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e8d8ef0b-447b-4a7c-b9bb-5bf860588420', 'CE7FF168' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7FF168');
INSERT INTO telemetry_goes (id, nesdis_id) select '94ce4d8d-336b-4b66-947e-33a31e2e81e8', '173EF176' where not exists (select 1 from telemetry_goes where nesdis_id = '173EF176');
INSERT INTO telemetry_goes (id, nesdis_id) select '44d4349f-9ebf-4837-8011-0c2b34f6ec21', 'CE4805A8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE4805A8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'efcb9937-061d-46c9-bb46-2669dcc8985e', 'CE41C838' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE41C838');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd7fe401b-5ba1-451d-968f-9f4a6c0113ab', 'CE4CD6FA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE4CD6FA');
INSERT INTO telemetry_goes (id, nesdis_id) select '675ab579-32d3-4b50-b05b-327754156e33', '1728C3BC' where not exists (select 1 from telemetry_goes where nesdis_id = '1728C3BC');
INSERT INTO telemetry_goes (id, nesdis_id) select '335a07af-341e-483c-a02b-4e1084a66935', 'CE1157CA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE1157CA');
INSERT INTO telemetry_goes (id, nesdis_id) select '9c70fa76-d499-4c1c-99ad-c3a99399e930', 'CE11215A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE11215A');
INSERT INTO telemetry_goes (id, nesdis_id) select '8bedf047-92aa-44f3-adc5-4396192e9fa8', 'CE5A8F04' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5A8F04');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ff6541a7-e553-461e-aec7-4016c0ff31a9', 'CE5B0538' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE5B0538');
INSERT INTO telemetry_goes (id, nesdis_id) select 'bc7bbb78-5c84-491b-800f-6b10ecb26470', 'CE06C702' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE06C702');
INSERT INTO telemetry_goes (id, nesdis_id) select '8b881d38-3f0d-4738-ae22-0cf2b6232acf', 'CE10F5C8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE10F5C8');
INSERT INTO telemetry_goes (id, nesdis_id) select '31efc4a0-30b8-461c-bf97-a76d26cd804f', 'CE021450' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE021450');
INSERT INTO telemetry_goes (id, nesdis_id) select '6c35806d-335e-41e2-8f1d-9f9c2a31c40c', '16F7970A' where not exists (select 1 from telemetry_goes where nesdis_id = '16F7970A');
INSERT INTO telemetry_goes (id, nesdis_id) select '6a6741d7-74af-4159-8c51-08c9889a76f1', 'CE09C190' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE09C190');
INSERT INTO telemetry_goes (id, nesdis_id) select '35112d47-25b4-43c1-bdf8-bd4f487f506c', 'CE40E2FC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE40E2FC');
INSERT INTO telemetry_goes (id, nesdis_id) select '8a954371-3e4a-4214-bdd4-1941cbc09f00', '17758792' where not exists (select 1 from telemetry_goes where nesdis_id = '17758792');
INSERT INTO telemetry_goes (id, nesdis_id) select 'f2896a07-2908-4f35-a4a7-8c9e2500a812', 'DD17E1E8' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD17E1E8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'b6b72cee-df96-4436-99a7-2464731f92f0', '1782274A' where not exists (select 1 from telemetry_goes where nesdis_id = '1782274A');
INSERT INTO telemetry_goes (id, nesdis_id) select '3b5474b6-0a71-455f-a8b5-1234a621dac9', 'DD05B3FE' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD05B3FE');
INSERT INTO telemetry_goes (id, nesdis_id) select '0cf1ad87-df49-4ad9-8c02-92f203e2ddd7', '171E136A' where not exists (select 1 from telemetry_goes where nesdis_id = '171E136A');
INSERT INTO telemetry_goes (id, nesdis_id) select '90156d1a-9c3c-4a3e-918d-a55bc4ae4614', '17CBD24E' where not exists (select 1 from telemetry_goes where nesdis_id = '17CBD24E');
INSERT INTO telemetry_goes (id, nesdis_id) select '1edd27c1-c9c3-4238-8095-11f92d9adc33', '178325B0' where not exists (select 1 from telemetry_goes where nesdis_id = '178325B0');
INSERT INTO telemetry_goes (id, nesdis_id) select '50b5ed80-2c4d-47ea-a212-f9c73b193082', '17CC3388' where not exists (select 1 from telemetry_goes where nesdis_id = '17CC3388');
INSERT INTO telemetry_goes (id, nesdis_id) select '19ef1234-ecad-40f7-8637-a0866b8684b3', '17834056' where not exists (select 1 from telemetry_goes where nesdis_id = '17834056');
INSERT INTO telemetry_goes (id, nesdis_id) select '3a218697-212b-4f26-af91-c9f32f5bcae5', '1783E0AE' where not exists (select 1 from telemetry_goes where nesdis_id = '1783E0AE');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fb4f6747-45fa-4fcb-ba3e-1f2a746cfd88', 'CE65A7D8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65A7D8');
INSERT INTO telemetry_goes (id, nesdis_id) select 'beadb18f-1e4e-4dc8-bd01-4ff4149004cb', '1653B28A' where not exists (select 1 from telemetry_goes where nesdis_id = '1653B28A');
INSERT INTO telemetry_goes (id, nesdis_id) select '75ffc306-0604-424a-8088-c859c6779cce', 'CE0D342E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0D342E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fb8496bd-f94f-4823-87b3-5cc687afe06b', 'CE660EFC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE660EFC');
INSERT INTO telemetry_goes (id, nesdis_id) select '5a418eb5-364f-4acd-9025-3dc381d8dd2c', 'CE67374E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE67374E');
INSERT INTO telemetry_goes (id, nesdis_id) select '5d29793d-25f1-4107-ab5d-ffc7cef5feb9', 'CE0EA642' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0EA642');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e1388504-75e6-45b0-8290-6538b4f744db', 'CE2FF52A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE2FF52A');
INSERT INTO telemetry_goes (id, nesdis_id) select 'fc487932-151e-46d8-92ec-04e96f4be9a5', 'CE03C0C2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE03C0C2');
INSERT INTO telemetry_goes (id, nesdis_id) select 'a1783378-cc48-4b22-901e-7abfe11ffa83', 'CE0232BC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0232BC');
INSERT INTO telemetry_goes (id, nesdis_id) select '448402ec-0cb2-4fb2-bb69-3b2c43e7a369', 'CE0E3320' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E3320');
INSERT INTO telemetry_goes (id, nesdis_id) select '03ff08c1-4883-45cf-afd3-5c40bebd5981', 'DD17116C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD17116C');
INSERT INTO telemetry_goes (id, nesdis_id) select '5d6df7f6-40c2-4cde-8ded-ada50a1b7ac2', 'CE3004A2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE3004A2');
INSERT INTO telemetry_goes (id, nesdis_id) select '083b8121-c2d4-4be9-8376-cc88a1e8eec5', 'CE0E80AE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E80AE');
INSERT INTO telemetry_goes (id, nesdis_id) select '2b36d6f3-44fa-4100-b79a-997acb037f62', 'CE0D12C2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0D12C2');
INSERT INTO telemetry_goes (id, nesdis_id) select '2d3370cd-02d3-4cd3-ac31-f99862028358', 'CE0D6452' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0D6452');
INSERT INTO telemetry_goes (id, nesdis_id) select '53777f35-fc62-4a53-9a96-148ddbbf2f65', 'CE617A58' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE617A58');
INSERT INTO telemetry_goes (id, nesdis_id) select '6554106d-b3ce-4d79-b714-ee9ac1befa11', 'CE411E50' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE411E50');
INSERT INTO telemetry_goes (id, nesdis_id) select '485b78c3-c795-44be-9f08-2a74483a4293', 'CE28710A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE28710A');
INSERT INTO telemetry_goes (id, nesdis_id) select '8f5ec9b1-6dde-4c82-aa1c-b5a1d073c637', 'CE65BA7C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE65BA7C');
INSERT INTO telemetry_goes (id, nesdis_id) select '2c4bfab6-a3d3-4d3f-8a88-6704cccd0b9a', 'CE303138' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE303138');
INSERT INTO telemetry_goes (id, nesdis_id) select '883033e2-0e63-4cc0-818c-65c066afad3d', 'CE80B6E4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE80B6E4');
INSERT INTO telemetry_goes (id, nesdis_id) select 'e924c8f6-e97e-4255-905b-96adb52cf23b', 'CE09D2E6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE09D2E6');
INSERT INTO telemetry_goes (id, nesdis_id) select 'a67ec75d-d175-4da7-9c59-ded81a883e3c', 'CE0E702A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E702A');
INSERT INTO telemetry_goes (id, nesdis_id) select '16ae3ea4-6589-4c63-8e74-53fb9b21cab0', 'CE0EB534' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0EB534');
INSERT INTO telemetry_goes (id, nesdis_id) select '9e61af73-bf8a-42ce-837f-e89468cb655c', 'CE658134' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE658134');
INSERT INTO telemetry_goes (id, nesdis_id) select '58e38f6e-f616-4965-9bc7-3c5b4740043d', 'CE09E77C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE09E77C');
INSERT INTO telemetry_goes (id, nesdis_id) select 'ac3c5a96-f622-4372-a4e0-34642ac44465', 'CE06B192' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE06B192');
INSERT INTO telemetry_goes (id, nesdis_id) select '95662ab3-7282-4deb-aee7-b4f731d773da', 'CE0E56C6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E56C6');
INSERT INTO telemetry_goes (id, nesdis_id) select '82c8cfa6-9571-4ba5-a6f7-0542fd273f4d', 'CE0E635C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E635C');
INSERT INTO telemetry_goes (id, nesdis_id) select '6b8306b8-a142-46b8-8e16-2adc7e58749d', 'CE0E93D8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE0E93D8');
INSERT INTO telemetry_goes (id, nesdis_id) select '25f45619-107a-42ee-9458-3395603d57a3', '16832412' where not exists (select 1 from telemetry_goes where nesdis_id = '16832412');
INSERT INTO telemetry_goes (id, nesdis_id) select '6ae0f672-7b78-4d6e-bad5-3cbb1a699f91', '1774C662' where not exists (select 1 from telemetry_goes where nesdis_id = '1774C662');
INSERT INTO telemetry_goes (id, nesdis_id) select '2e728675-0b18-446f-abd2-3944aa7b3be1', 'DD17840E' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD17840E');
INSERT INTO telemetry_goes (id, nesdis_id) select 'd88973fc-3ef9-4dac-bf22-9b283a459368', 'CE4580DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE4580DA');
INSERT INTO telemetry_goes (id, nesdis_id) select '8ca92104-beb1-40e8-b742-6cb689b1b42b', 'CE61D472' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61D472');
INSERT INTO telemetry_goes (id, nesdis_id) select '735f7426-cb80-450f-bb37-a41c17d5a023', 'CE03D3B4' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE03D3B4');
INSERT INTO telemetry_goes (id, nesdis_id) select 'eafeba3c-660b-49a4-8078-59f8c899fa1b', 'CE61C704' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE61C704');
INSERT INTO telemetry_goes (id, nesdis_id) select 'dd3f8909-43e1-47c9-905a-e7ab75f6e97d', 'CE577234' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE577234');

--INSERT INSTRUMENT_TELEMETRY--COUNT:159
INSERT INTO instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('15dacad3-a7e5-4378-91b4-34208e7039cb', '10a32652-af43-4451-bd52-4980c5690cc9', '03d73df2-3db2-4003-95e5-7bce29a52b0e'),
('35896e62-b69f-4bc9-9fc6-78ea7c5fa92a', '10a32652-af43-4451-bd52-4980c5690cc9', '6edd26bd-bb45-46a5-93ec-1d9bb87fd52f'),
('bbeff60a-9198-4956-9ef2-021fc851cdbd', '10a32652-af43-4451-bd52-4980c5690cc9', '79412361-14d1-4b65-b5bb-eb2338539e20'),
('64d4fe14-952c-4f0a-b256-1ae56e687fd2', '10a32652-af43-4451-bd52-4980c5690cc9', '8887b5c8-de97-4c32-b63e-58b0ef975aaa'),
('d469455f-6ce9-4faf-b2fc-c321f3365a6f', '10a32652-af43-4451-bd52-4980c5690cc9', '1ba3a311-a34b-47e8-8103-9e1eac4e2ce1'),
('41e9219e-2f1f-439f-a724-a3e5c7c7cf73', '10a32652-af43-4451-bd52-4980c5690cc9', '08e43040-2a98-40fd-98c0-699384f5f48b'),
('95b436bc-842a-4ad8-83e7-2410a56c439c', '10a32652-af43-4451-bd52-4980c5690cc9', '712ed4ff-dc49-45a5-a0e2-1487c4581f8a'),
('016d2287-c4a8-4544-9588-dfb961726c26', '10a32652-af43-4451-bd52-4980c5690cc9', 'f8cfd6bb-7a6c-4d66-9983-82cf388728e4'),
('3f1daace-2b8e-4d6c-a21c-1a23359ae00b', '10a32652-af43-4451-bd52-4980c5690cc9', '38847c6d-58af-4848-bee3-fde7e7e4f40a'),
('38fa06c8-6457-4523-b4ef-3ebffa520151', '10a32652-af43-4451-bd52-4980c5690cc9', 'f15b511b-5e42-46f3-b7cf-a8e112c92191'),
('b9e94249-d8e9-41ac-b683-bde39fa79ed4', '10a32652-af43-4451-bd52-4980c5690cc9', '6a629e80-0fdb-4418-afd9-58c20cb289a0'),
('616fa910-eabd-4718-940f-c7ec2afeffd8', '10a32652-af43-4451-bd52-4980c5690cc9', 'bd0507cb-ef3c-45df-a129-961463a7f22b'),
('6d781aa1-93e5-4aa5-85fe-9d057ead4247', '10a32652-af43-4451-bd52-4980c5690cc9', '5c4da467-cf0f-444e-9b8a-3a7398a02c25'),
('4adb6c37-4ae1-43e1-9540-e5595bf86d1b', '10a32652-af43-4451-bd52-4980c5690cc9', '48196b78-149e-449a-a623-791d18496637'),
('67cf3e31-f3de-4413-930d-1cf4903c454e', '10a32652-af43-4451-bd52-4980c5690cc9', '9d60b511-b8cc-41ac-99ef-048e01e889d2'),
('bacf256e-eb96-4c36-8954-bd92a306f4d8', '10a32652-af43-4451-bd52-4980c5690cc9', '02dda7b8-e77a-4ee8-981f-109310db9602'),
('5524e00c-1f89-418d-9740-6de8645b8333', '10a32652-af43-4451-bd52-4980c5690cc9', '73956d33-2c35-41df-81aa-1d5436f06df3'),
('d7281dd7-778d-4c3a-90a4-b21bb641e850', '10a32652-af43-4451-bd52-4980c5690cc9', 'f2ee2a9a-3e24-415c-a339-875f6c42bd52'),
('f6e8916f-296a-49e9-aca1-82665748ca62', '10a32652-af43-4451-bd52-4980c5690cc9', '458b72bc-c2d2-4ab9-a600-4e7105015ca2'),
('70ba9edc-40ea-42ad-a8a3-ceabb25b05e2', '10a32652-af43-4451-bd52-4980c5690cc9', 'a5c821c6-3a1a-4dd1-8028-d05cf9294e25'),
('694b13bf-449a-4c44-8aa2-c38469c07ed5', '10a32652-af43-4451-bd52-4980c5690cc9', 'c79f68db-ef3d-4b00-a4de-0ad99ce0e20c'),
('a392c3e8-738d-4feb-afc1-00d064e4298e', '10a32652-af43-4451-bd52-4980c5690cc9', 'acd11117-6a70-4e3f-9add-763d0e2c0720'),
('8add8896-c24d-44b3-9853-c189d2d83695', '10a32652-af43-4451-bd52-4980c5690cc9', '04d2c33d-331a-4141-ac31-080a3a877d59'),
('1592e55f-ff35-40a9-80c4-7c81e4f620b3', '10a32652-af43-4451-bd52-4980c5690cc9', '22791f7f-8f05-42ff-9ad2-f3cd3672074b'),
('3604b451-24dc-423f-9e8a-f0bbf9d7d8f2', '10a32652-af43-4451-bd52-4980c5690cc9', '7ddcbd7b-8571-4a5f-9b2b-f5eada8cb5c6'),
('9d3ea29e-59a3-46dd-a091-d8fed56f03ec', '10a32652-af43-4451-bd52-4980c5690cc9', 'e83b9db7-488b-47fc-bd14-935ae625556c'),
('4c7ce601-f232-4634-936c-d7604eb54e3b', '10a32652-af43-4451-bd52-4980c5690cc9', '5dcf465d-70bf-4163-902c-c128bbf10ad3'),
('e213714d-2656-4719-84f5-61e208032038', '10a32652-af43-4451-bd52-4980c5690cc9', 'd45c5dca-a28f-4bdc-bf73-b5bf5a93bf02'),
('6bdca601-8044-4830-85e5-7f34dd5d5865', '10a32652-af43-4451-bd52-4980c5690cc9', 'fc73977c-e4da-4190-98f3-12134d2b78e8'),
('75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', '10a32652-af43-4451-bd52-4980c5690cc9', 'e458550d-eb9e-47b7-90c0-cc8e7b3a6ee5'),
('8091654d-2afd-46a6-a2a7-26447b8c0df4', '10a32652-af43-4451-bd52-4980c5690cc9', '059e3288-618f-4221-863b-cc46b7aa1295'),
('2cce0e9d-c608-4424-a8f8-0f76e0691338', '10a32652-af43-4451-bd52-4980c5690cc9', '5542dad4-dc5a-4b72-8178-d921282ebcf2'),
('82d2991c-71db-4699-9e7f-5506f9614d58', '10a32652-af43-4451-bd52-4980c5690cc9', '5c422eac-3dcf-4416-b5ff-3575beb673bf'),
('21bdce69-d2a7-49f0-807b-d29c747f2802', '10a32652-af43-4451-bd52-4980c5690cc9', 'ddddf336-7c67-447c-bd41-fe14f13b0333'),
('28bc2a6e-72a9-4e21-9da0-29ce3dda495d', '10a32652-af43-4451-bd52-4980c5690cc9', '883be004-c07c-47d3-aca9-e44afa3c316b'),
('00fb2016-699b-45cf-a44f-2944e0860c26', '10a32652-af43-4451-bd52-4980c5690cc9', '9b75e35f-ba64-42aa-95d9-602f5db81268'),
('df1cf782-4230-4f79-8437-4e361ee67dfc', '10a32652-af43-4451-bd52-4980c5690cc9', '8bdb09dc-9886-4101-ac1f-d414644b9ec4'),
('f0658ffe-4b1f-4b6c-b17b-7b8cc43ad302', '10a32652-af43-4451-bd52-4980c5690cc9', '87296849-5609-4114-8761-48904cc647c1'),
('581ba9dc-3836-4fea-9598-fc0d9255da47', '10a32652-af43-4451-bd52-4980c5690cc9', 'e583fe22-47a2-44d8-bdfd-eee0dd0ad05c'),
('9ce6e98b-a26b-450e-b859-41bdc05bab06', '10a32652-af43-4451-bd52-4980c5690cc9', '3e489698-2d15-4264-b484-ec3883959f9e'),
('1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '10a32652-af43-4451-bd52-4980c5690cc9', 'ab95547c-07e2-45fd-9917-345ec4f0fe48'),
('5c3f1f19-1639-477f-81e0-8edf1b85bda7', '10a32652-af43-4451-bd52-4980c5690cc9', '12952bdc-8f8e-43be-a619-4e70c6005c51'),
('449c9725-605c-4d0d-a0c6-8e434054b26d', '10a32652-af43-4451-bd52-4980c5690cc9', '3cda0e26-9d89-45bb-960c-ed4ce3f22893'),
('a70a73d7-47fa-4f92-9b00-6bf9f02c52df', '10a32652-af43-4451-bd52-4980c5690cc9', '61d18727-7980-4218-bc9a-7459c92e5fc6'),
('d56068c3-a3c4-4553-b30a-a885570cd7d8', '10a32652-af43-4451-bd52-4980c5690cc9', '6a10892d-82e9-4e14-9977-c096c2cf30ac'),
('6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '10a32652-af43-4451-bd52-4980c5690cc9', '608de20c-bfa0-4ac6-a9ec-3bbc3b9ee7e3'),
('572ce69b-8b32-40ab-ac00-e783dc6ceb4f', '10a32652-af43-4451-bd52-4980c5690cc9', 'f6c85772-21a5-40ab-bc23-65d606facd26'),
('bcfb5c23-300d-4d78-b8b7-f7cf40bfbd09', '10a32652-af43-4451-bd52-4980c5690cc9', '9aaf8a00-a76c-4505-89db-2f6a3017f9d5'),
('0d2b2c4a-e8ab-464d-aa5e-5d4fb3145dc5', '10a32652-af43-4451-bd52-4980c5690cc9', '4fffe379-5431-4069-8185-9b4716a802a3'),
('d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '10a32652-af43-4451-bd52-4980c5690cc9', 'f833d903-8b98-4071-838c-86a71acb43a9'),
('5e19aa7a-1611-49e9-bec3-75c3421c7b5e', '10a32652-af43-4451-bd52-4980c5690cc9', '822b81ea-967f-4bc8-99a1-1e87776b8a53'),
('e1993ab5-ec08-4e07-bc3c-c4db92731d1c', '10a32652-af43-4451-bd52-4980c5690cc9', 'e9f318a5-5200-495f-9c4a-50849474d05d'),
('2db45595-fefc-4eb8-9237-8fc72561b47e', '10a32652-af43-4451-bd52-4980c5690cc9', '6dcc781b-afda-40d3-89ea-ddfbf658dde0'),
('f3b99e3b-1856-4226-a73c-1c46b200fb66', '10a32652-af43-4451-bd52-4980c5690cc9', 'bbea5db8-f4ad-41d0-919b-b1ba78361763'),
('72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '10a32652-af43-4451-bd52-4980c5690cc9', 'f4a69857-87c1-4170-8f21-4fbb5b238d8c'),
('6b78716c-4975-46e4-b635-9ff79e7f0627', '10a32652-af43-4451-bd52-4980c5690cc9', '9436171f-b2a5-4ab0-9816-159170e5dec4'),
('a88ece8d-d3ba-4a77-a413-5dd03ebfb623', '10a32652-af43-4451-bd52-4980c5690cc9', '34c74351-2122-448b-89c1-cb321cc50876'),
('f6629ae2-98e2-4e06-b685-331189cff3a7', '10a32652-af43-4451-bd52-4980c5690cc9', 'd2c00c91-b9d8-46e0-8a71-b1195bfb5cd6'),
('59db968a-43d1-4590-b4f7-6f7f6423e735', '10a32652-af43-4451-bd52-4980c5690cc9', 'ac4d0beb-88c0-4f31-b613-486bfa08d5d3'),
('ac71a2f1-74bb-49de-8dd4-d2b15b974978', '10a32652-af43-4451-bd52-4980c5690cc9', 'e65f9673-aede-4199-99b3-2a94f87fe771'),
('ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', '10a32652-af43-4451-bd52-4980c5690cc9', '4c5fdc7f-0545-44df-9785-d625d72c60dd'),
('4916c4af-8019-4492-a734-fc3ff8c49074', '10a32652-af43-4451-bd52-4980c5690cc9', 'b06cee02-eb20-474e-a861-5cf42e03a396'),
('5dfc826b-a94c-4cee-b057-6fb858afc95d', '10a32652-af43-4451-bd52-4980c5690cc9', '43a7ed8d-8f4b-4284-b0c2-324c93418644'),
('7e772c75-64b6-485e-8294-d1d3bc6da0fe', '10a32652-af43-4451-bd52-4980c5690cc9', 'e5b47b00-45be-4c16-a590-e8b9df8c8d18'),
('c4c87cf8-a155-4abe-bf89-4f7af607db31', '10a32652-af43-4451-bd52-4980c5690cc9', 'fefd961a-a953-492c-af37-e736cc5bfced'),
('3fcc1834-0263-4dc5-a7c8-9157c767bb4f', '10a32652-af43-4451-bd52-4980c5690cc9', 'd0038cdb-aed7-4aac-b043-961d84923e2c'),
('6890212a-c7f1-499e-a3fe-e7d973eb7cc1', '10a32652-af43-4451-bd52-4980c5690cc9', 'ae215b11-24d5-4a5b-9819-8f41c91b3439'),
('bb1389a0-bfad-430d-a66d-2f041a738cb1', '10a32652-af43-4451-bd52-4980c5690cc9', '8cd3406b-a43d-4080-9850-e10d5e74d7b4'),
('04e220cf-128b-4115-8a7b-76ac2625a81b', '10a32652-af43-4451-bd52-4980c5690cc9', '78fd0eb2-ff24-4db8-90a9-6380c8e595ce'),
('7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', '10a32652-af43-4451-bd52-4980c5690cc9', '761c9eb2-a9e5-4b45-822f-379a3cc19d5b'),
('8250ac38-3d9f-4039-9f08-149912301977', '10a32652-af43-4451-bd52-4980c5690cc9', 'd34c00b4-db45-43ec-a20d-58d6705cd240'),
('9dcf334f-2982-49c9-941c-38c91565f019', '10a32652-af43-4451-bd52-4980c5690cc9', 'dfc714e0-c5aa-4d6d-b92c-bc49ee8e9a9c'),
('1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', '10a32652-af43-4451-bd52-4980c5690cc9', '839aeca4-d480-4cb8-a1da-9bda8e0d00ad'),
('3eba9840-cebb-45ee-8e9f-a62d8279c86a', '10a32652-af43-4451-bd52-4980c5690cc9', 'a1081cd1-23a5-4e5f-87b8-4d3523760f8c'),
('bc7d2d9a-8e02-4a78-b838-cadd444ed756', '10a32652-af43-4451-bd52-4980c5690cc9', 'ca7bc2a4-5283-457a-8382-b3d45c091dbc'),
('3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', '10a32652-af43-4451-bd52-4980c5690cc9', 'f58f0217-a629-46d8-853b-79baa9f6604b'),
('5171ef3b-9ef9-4984-8ab7-57b76db13621', '10a32652-af43-4451-bd52-4980c5690cc9', '96a9674d-4f64-4552-9b5f-ad353d68bc04'),
('b72006fc-5809-4985-8f3c-45baed7cfcb9', '10a32652-af43-4451-bd52-4980c5690cc9', '60d9864b-8140-4527-871b-28c015f36cf3'),
('4811e8d0-4c72-4c86-8bec-37e491f5ab63', '10a32652-af43-4451-bd52-4980c5690cc9', '6a65e3a2-2af6-4aa0-9cfb-ab8801507e28'),
('3f9a0fcf-4e4f-4b0c-ac88-011455d6957f', '10a32652-af43-4451-bd52-4980c5690cc9', '4d809aeb-d269-468d-8c90-b0f62bafb8a8'),
('553f696a-9292-41b1-a4bf-67236068a1d6', '10a32652-af43-4451-bd52-4980c5690cc9', '4f1f6779-d931-45f7-ac57-88a7d158816b'),
('394d6441-43b2-4745-b97d-c6462c7bc908', '10a32652-af43-4451-bd52-4980c5690cc9', '6bc93d1a-fa95-45b8-81ea-bdd7b5ad8ba6'),
('2711323d-5ed2-44b5-a7a6-f0614c0b422b', '10a32652-af43-4451-bd52-4980c5690cc9', 'ac681a2e-2577-4e1b-9628-97fe3e4cbf70'),
('3bd76812-033f-4367-b8c4-eceba262f393', '10a32652-af43-4451-bd52-4980c5690cc9', 'f186e2c0-295b-4bb5-aa65-8245c9bb3c78'),
('f8d6ccd6-fb48-4228-b777-e423612b1f27', '10a32652-af43-4451-bd52-4980c5690cc9', 'd4359bc1-0563-4ec2-a1b8-b45c75b3b393'),
('7babc4e1-950e-41e8-874f-49559dd947bf', '10a32652-af43-4451-bd52-4980c5690cc9', '8c44b370-114b-4e09-8634-8fc4ad5ee599'),
('ab7ee39b-8397-484a-ac29-e2e813ed6adb', '10a32652-af43-4451-bd52-4980c5690cc9', '57c8c07b-621d-4cbb-8efc-857f0aeb0b45'),
('d93718c4-08ae-4065-8855-674c5206c1ed', '10a32652-af43-4451-bd52-4980c5690cc9', '7e96b629-0c92-4fe0-b096-523e24d9664c'),
('de20f637-d3c7-4ea4-8ca6-604f6a127532', '10a32652-af43-4451-bd52-4980c5690cc9', '8280848b-f4e6-4ab6-99c9-6070e802f81a'),
('8a7739c0-ccae-4f3b-83d7-e2f6f506e6da', '10a32652-af43-4451-bd52-4980c5690cc9', '1a7a69cc-4b69-4350-9d7d-0aeece773a0c'),
('2006cc18-de31-4ae0-8aaa-81ba6cf1ba95', '10a32652-af43-4451-bd52-4980c5690cc9', '7d59b71e-79eb-4161-9cd8-17cfe78352b5'),
('e29fcc6e-a0a8-4fe0-b450-099142997846', '10a32652-af43-4451-bd52-4980c5690cc9', '2169217d-63b9-437f-96df-d8717ece6f5c'),
('7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', '10a32652-af43-4451-bd52-4980c5690cc9', 'c71b5df9-9544-4dd5-803a-86356e44434a'),
('bcc35574-bf23-4952-9760-049ba8843bd7', '10a32652-af43-4451-bd52-4980c5690cc9', '709e06c8-f35c-447b-a8f3-e478c5777239'),
('4bbb3d28-b816-4f71-98c1-de647d8f5e38', '10a32652-af43-4451-bd52-4980c5690cc9', 'de10d184-a255-42c4-befc-d1a36b0e8704'),
('b3524459-4cbe-4802-835e-b415cc6657b2', '10a32652-af43-4451-bd52-4980c5690cc9', 'e8d8ef0b-447b-4a7c-b9bb-5bf860588420'),
('ed2c1d1c-7bb2-4374-8e90-878ab6b70d9c', '10a32652-af43-4451-bd52-4980c5690cc9', '94ce4d8d-336b-4b66-947e-33a31e2e81e8'),
('eaa406e4-560c-4fa3-9c28-8f265999f17b', '10a32652-af43-4451-bd52-4980c5690cc9', '44d4349f-9ebf-4837-8011-0c2b34f6ec21'),
('bd3d7701-94b0-4390-b2f8-40e03c48508c', '10a32652-af43-4451-bd52-4980c5690cc9', 'efcb9937-061d-46c9-bb46-2669dcc8985e'),
('8e5ed58a-febf-4b91-86e3-4b39d798a592', '10a32652-af43-4451-bd52-4980c5690cc9', 'd7fe401b-5ba1-451d-968f-9f4a6c0113ab'),
('33564f35-da99-4db5-9d9d-2c96698f99e6', '10a32652-af43-4451-bd52-4980c5690cc9', '675ab579-32d3-4b50-b05b-327754156e33'),
('228d8881-b162-4842-92c9-75422535a5ad', '10a32652-af43-4451-bd52-4980c5690cc9', '335a07af-341e-483c-a02b-4e1084a66935'),
('8f32236d-1e18-47f4-bf36-383eb8c82fbe', '10a32652-af43-4451-bd52-4980c5690cc9', '9c70fa76-d499-4c1c-99ad-c3a99399e930'),
('dccf96c3-4dff-4b3b-a7b7-446a97db654b', '10a32652-af43-4451-bd52-4980c5690cc9', '8bedf047-92aa-44f3-adc5-4396192e9fa8'),
('8527c379-3f95-4ef6-aae7-4fcf94ad8691', '10a32652-af43-4451-bd52-4980c5690cc9', 'ff6541a7-e553-461e-aec7-4016c0ff31a9'),
('8bd798a4-1c7b-4684-ae3b-eb75a62a09cf', '10a32652-af43-4451-bd52-4980c5690cc9', 'bc7bbb78-5c84-491b-800f-6b10ecb26470'),
('2eed3135-b0ee-423d-b8a6-5336072ac52a', '10a32652-af43-4451-bd52-4980c5690cc9', '8b881d38-3f0d-4738-ae22-0cf2b6232acf'),
('3277950f-fca8-4bcd-8799-64570ea99bd6', '10a32652-af43-4451-bd52-4980c5690cc9', '31efc4a0-30b8-461c-bf97-a76d26cd804f'),
('9bda847a-a23e-430c-bbbe-27da4e67d463', '10a32652-af43-4451-bd52-4980c5690cc9', '6c35806d-335e-41e2-8f1d-9f9c2a31c40c'),
('9b19c51c-b11f-4d74-9eab-2228dc25508b', '10a32652-af43-4451-bd52-4980c5690cc9', '6a6741d7-74af-4159-8c51-08c9889a76f1'),
('8965cb62-506a-4f0e-b285-48fe8ddfe35e', '10a32652-af43-4451-bd52-4980c5690cc9', '35112d47-25b4-43c1-bdf8-bd4f487f506c'),
('2432c199-8996-4ae7-8a30-50c52c9f746e', '10a32652-af43-4451-bd52-4980c5690cc9', '8a954371-3e4a-4214-bdd4-1941cbc09f00'),
('93960641-3b74-456c-8fcf-7aadbf48c6d8', '10a32652-af43-4451-bd52-4980c5690cc9', 'f2896a07-2908-4f35-a4a7-8c9e2500a812'),
('1e6a3e9c-3b71-4256-a4b7-75a9615ba8b1', '10a32652-af43-4451-bd52-4980c5690cc9', 'b6b72cee-df96-4436-99a7-2464731f92f0'),
('d56eb333-2749-4797-b5c8-f282ecf8fd4a', '10a32652-af43-4451-bd52-4980c5690cc9', '3b5474b6-0a71-455f-a8b5-1234a621dac9'),
('3ba8a61e-8920-46d2-9831-8cd0713508ef', '10a32652-af43-4451-bd52-4980c5690cc9', '0cf1ad87-df49-4ad9-8c02-92f203e2ddd7'),
('b72a87bd-94c0-48af-b773-d9bdc5419859', '10a32652-af43-4451-bd52-4980c5690cc9', '90156d1a-9c3c-4a3e-918d-a55bc4ae4614'),
('69d3de52-e320-46e4-8c5b-f29cb8cbb4bb', '10a32652-af43-4451-bd52-4980c5690cc9', '1edd27c1-c9c3-4238-8095-11f92d9adc33'),
('b4cd105a-dd1c-485f-bbdb-b80269084676', '10a32652-af43-4451-bd52-4980c5690cc9', '50b5ed80-2c4d-47ea-a212-f9c73b193082'),
('d4682aed-6cd6-498c-a610-3d565b03c971', '10a32652-af43-4451-bd52-4980c5690cc9', '19ef1234-ecad-40f7-8637-a0866b8684b3'),
('017d8b2c-2234-412e-a88c-d5949ce2fd02', '10a32652-af43-4451-bd52-4980c5690cc9', '3a218697-212b-4f26-af91-c9f32f5bcae5'),
('333d17db-11e8-47b9-8621-af2fccce7190', '10a32652-af43-4451-bd52-4980c5690cc9', 'fb4f6747-45fa-4fcb-ba3e-1f2a746cfd88'),
('1cff71c6-56ec-4baa-8524-ddcc2b292f25', '10a32652-af43-4451-bd52-4980c5690cc9', 'beadb18f-1e4e-4dc8-bd01-4ff4149004cb'),
('afda57df-9971-4d0d-8d7f-7092d94d1791', '10a32652-af43-4451-bd52-4980c5690cc9', '75ffc306-0604-424a-8088-c859c6779cce'),
('7b94b33a-0922-42ab-9928-2d4dfe40aa22', '10a32652-af43-4451-bd52-4980c5690cc9', 'fb8496bd-f94f-4823-87b3-5cc687afe06b'),
('d2beb71a-5aeb-4ba3-ae0e-346e81943d97', '10a32652-af43-4451-bd52-4980c5690cc9', '5a418eb5-364f-4acd-9025-3dc381d8dd2c'),
('704bded7-49c2-43d6-9c4f-9b6b1512b9f1', '10a32652-af43-4451-bd52-4980c5690cc9', '5d29793d-25f1-4107-ab5d-ffc7cef5feb9'),
('707c1f3f-cf12-4768-b826-419478701450', '10a32652-af43-4451-bd52-4980c5690cc9', 'e1388504-75e6-45b0-8290-6538b4f744db'),
('60a37953-139e-4d66-b000-06f5b16ece51', '10a32652-af43-4451-bd52-4980c5690cc9', 'fc487932-151e-46d8-92ec-04e96f4be9a5'),
('908dcb66-f39e-415a-af5f-fa8272a67a42', '10a32652-af43-4451-bd52-4980c5690cc9', 'a1783378-cc48-4b22-901e-7abfe11ffa83'),
('5478ea0d-9117-48fb-9694-02515d516579', '10a32652-af43-4451-bd52-4980c5690cc9', '448402ec-0cb2-4fb2-bb69-3b2c43e7a369'),
('59062009-9c5c-413e-9f0a-077fe4d1a015', '10a32652-af43-4451-bd52-4980c5690cc9', '03ff08c1-4883-45cf-afd3-5c40bebd5981'),
('a4b91146-0108-4d0d-8cba-21e9cc6281eb', '10a32652-af43-4451-bd52-4980c5690cc9', '5d6df7f6-40c2-4cde-8ded-ada50a1b7ac2'),
('4f556761-b926-4086-bdcd-19bcc311ff1e', '10a32652-af43-4451-bd52-4980c5690cc9', '083b8121-c2d4-4be9-8376-cc88a1e8eec5'),
('4369132f-f2b1-4161-a857-95e6bcc2de59', '10a32652-af43-4451-bd52-4980c5690cc9', '2b36d6f3-44fa-4100-b79a-997acb037f62'),
('76862e2d-87cb-4605-804b-e4d9cb4ce00c', '10a32652-af43-4451-bd52-4980c5690cc9', '2d3370cd-02d3-4cd3-ac31-f99862028358'),
('d4198ec8-52ae-4296-a5ce-3763c3165d79', '10a32652-af43-4451-bd52-4980c5690cc9', '53777f35-fc62-4a53-9a96-148ddbbf2f65'),
('68eaa243-1fe5-4615-9d87-eb0779971a30', '10a32652-af43-4451-bd52-4980c5690cc9', '6554106d-b3ce-4d79-b714-ee9ac1befa11'),
('b3087995-8550-4b0f-a201-dbbf1e52698b', '10a32652-af43-4451-bd52-4980c5690cc9', '485b78c3-c795-44be-9f08-2a74483a4293'),
('47c17779-6fd6-4620-a176-4112878f687a', '10a32652-af43-4451-bd52-4980c5690cc9', '8f5ec9b1-6dde-4c82-aa1c-b5a1d073c637'),
('081d900a-ef53-442b-940f-d7cee74a26d8', '10a32652-af43-4451-bd52-4980c5690cc9', '2c4bfab6-a3d3-4d3f-8a88-6704cccd0b9a'),
('514648d9-4d3b-4bc4-a57e-edd091b9f3ce', '10a32652-af43-4451-bd52-4980c5690cc9', '883033e2-0e63-4cc0-818c-65c066afad3d'),
('15101a32-a069-4b0b-a29b-f86a409053b3', '10a32652-af43-4451-bd52-4980c5690cc9', 'e924c8f6-e97e-4255-905b-96adb52cf23b'),
('ce4e2f40-708a-434e-9d71-2b441ce8611b', '10a32652-af43-4451-bd52-4980c5690cc9', 'a67ec75d-d175-4da7-9c59-ded81a883e3c'),
('2725f343-d52f-43b2-b41d-0f9f671aa84e', '10a32652-af43-4451-bd52-4980c5690cc9', '16ae3ea4-6589-4c63-8e74-53fb9b21cab0'),
('c0b1ecd2-af0a-450c-8a51-f7b44130d612', '10a32652-af43-4451-bd52-4980c5690cc9', '9e61af73-bf8a-42ce-837f-e89468cb655c'),
('94e20e3a-7c6b-4eae-ac1c-fea8a22e4662', '10a32652-af43-4451-bd52-4980c5690cc9', '58e38f6e-f616-4965-9bc7-3c5b4740043d'),
('615d5750-5d23-4dfc-98d1-5fa69b3fc8ca', '10a32652-af43-4451-bd52-4980c5690cc9', 'ac3c5a96-f622-4372-a4e0-34642ac44465'),
('f8f8fd49-cdf8-47ec-bfaa-a406ccbb12c2', '10a32652-af43-4451-bd52-4980c5690cc9', '95662ab3-7282-4deb-aee7-b4f731d773da'),
('e0806b46-1b66-460e-97d5-0c98bff2edd9', '10a32652-af43-4451-bd52-4980c5690cc9', '82c8cfa6-9571-4ba5-a6f7-0542fd273f4d'),
('b7425815-d58a-4422-9907-e115f41c77dd', '10a32652-af43-4451-bd52-4980c5690cc9', '6b8306b8-a142-46b8-8e16-2adc7e58749d'),
('6207bc92-5cc7-46ed-a93f-1ff32121a68a', '10a32652-af43-4451-bd52-4980c5690cc9', '25f45619-107a-42ee-9458-3395603d57a3'),
('5f0e95fe-84b1-4dee-9584-408d2b79b4e1', '10a32652-af43-4451-bd52-4980c5690cc9', '6ae0f672-7b78-4d6e-bad5-3cbb1a699f91'),
('34d6ecd5-a584-4de1-90a6-dacc386988bd', '10a32652-af43-4451-bd52-4980c5690cc9', '2e728675-0b18-446f-abd2-3944aa7b3be1'),
('5f4a3ce2-f359-4f7b-b622-9db903ca8f30', '10a32652-af43-4451-bd52-4980c5690cc9', 'd88973fc-3ef9-4dac-bf22-9b283a459368'),
('d0c6b11b-44a4-4041-9cd7-1ff725f042a4', '10a32652-af43-4451-bd52-4980c5690cc9', '8ca92104-beb1-40e8-b742-6cb689b1b42b'),
('9c4eec28-485c-405f-93fe-a1671ba1e85e', '10a32652-af43-4451-bd52-4980c5690cc9', '735f7426-cb80-450f-bb37-a41c17d5a023'),
('2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', '10a32652-af43-4451-bd52-4980c5690cc9', 'eafeba3c-660b-49a4-8078-59f8c899fa1b'),
('c8197081-957a-451d-ab4d-41805e0ecf5f', '10a32652-af43-4451-bd52-4980c5690cc9', 'dd3f8909-43e1-47c9-905a-e7ab75f6e97d');

--INSERT TIMESERIES--COUNT:167
INSERT INTO timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('27595ebe-3261-4abf-94f9-327a98d6b7a8','stage','Stage','15dacad3-a7e5-4378-91b4-34208e7039cb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('80b29a72-2cf7-469a-8943-4b3f0d3e3741','voltage','Voltage','15dacad3-a7e5-4378-91b4-34208e7039cb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('23c91c24-8f1d-4a45-bf6b-e644bb56b0a6','stage','Stage','35896e62-b69f-4bc9-9fc6-78ea7c5fa92a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('a3aa6008-841b-420b-bd7f-f538ab839a14','voltage','Voltage','35896e62-b69f-4bc9-9fc6-78ea7c5fa92a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('b42be20f-4497-4b72-9c75-3453ccfbb7f5','stage','Stage','bbeff60a-9198-4956-9ef2-021fc851cdbd', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('39cad445-3d55-4525-8f31-5890fd8fa9f8','voltage','Voltage','bbeff60a-9198-4956-9ef2-021fc851cdbd', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('88491bf4-f911-4d3b-8cd3-6784a9600549','stage','Stage','64d4fe14-952c-4f0a-b256-1ae56e687fd2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('adef9b54-dea4-465a-9663-7fdd7b57eef1','voltage','Voltage','64d4fe14-952c-4f0a-b256-1ae56e687fd2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('750bbaef-c3d7-4375-a224-b51ae24606f5','unknown-us','Unknown US','64d4fe14-952c-4f0a-b256-1ae56e687fd2', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('112630b2-230d-40cd-974c-4027c3f955c2','unknown-ud','Unknown UD','64d4fe14-952c-4f0a-b256-1ae56e687fd2', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('5faf53af-2165-40f0-b8b5-2b71da3f1e52','unknown-pa','Unknown PA','d469455f-6ce9-4faf-b2fc-c321f3365a6f', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('cb682fda-23a8-4094-ba05-79b48627ab79','unknown-us','Unknown US','d469455f-6ce9-4faf-b2fc-c321f3365a6f', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('35e35b3f-8179-4243-a44d-1fbc43fe9a59','unknown-ud','Unknown UD','d469455f-6ce9-4faf-b2fc-c321f3365a6f', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('2aaae673-02fe-495e-819d-2d324a86113f','stage','Stage','d469455f-6ce9-4faf-b2fc-c321f3365a6f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('388f6d64-7225-45cd-a472-c1b8034d1cf6','water-temperature','Water-Temperature','d469455f-6ce9-4faf-b2fc-c321f3365a6f', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('fc26fb2a-fc2a-4339-b204-800219a14db0','precipitation','Precipitation','d469455f-6ce9-4faf-b2fc-c321f3365a6f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f469e806-f20c-42cf-8891-0e4ad238388f','voltage','Voltage','d469455f-6ce9-4faf-b2fc-c321f3365a6f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('6b123227-2f52-400f-ba7b-d6875cb4a60e','stage','Stage','41e9219e-2f1f-439f-a724-a3e5c7c7cf73', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bb0a2af5-4200-40b5-a5cd-07f840bbd3c1','unknown-us','Unknown US','41e9219e-2f1f-439f-a724-a3e5c7c7cf73', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('0fdbf073-57e5-465c-8267-30840ea5f8ba','unknown-ud','Unknown UD','41e9219e-2f1f-439f-a724-a3e5c7c7cf73', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('40da9747-8ad8-4ea4-ae05-506f63f58782','voltage','Voltage','41e9219e-2f1f-439f-a724-a3e5c7c7cf73', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('7e0efc7d-0dae-423e-a9bd-cc73d87e2dc4','stage','Stage','95b436bc-842a-4ad8-83e7-2410a56c439c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8422c025-d5c6-49a1-97f6-47f82b485ff4','voltage','Voltage','95b436bc-842a-4ad8-83e7-2410a56c439c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('df201e91-eea5-491b-82a0-d41afd802c24','stage','Stage','016d2287-c4a8-4544-9588-dfb961726c26', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('91528922-14a4-4adc-b56e-9fb2cc41558d','voltage','Voltage','016d2287-c4a8-4544-9588-dfb961726c26', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('bb8baf9b-2658-4419-a4c2-c5ae2af25e19','stage','Stage','3f1daace-2b8e-4d6c-a21c-1a23359ae00b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b7f02293-bc23-4770-b8d1-d34b31dea63c','unknown-us','Unknown US','3f1daace-2b8e-4d6c-a21c-1a23359ae00b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f3e12b42-f82b-4908-bd39-d4d705fa25b1','unknown-ud','Unknown UD','3f1daace-2b8e-4d6c-a21c-1a23359ae00b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('0b16b5f0-d6f3-4f0e-bc24-69af07407b52','voltage','Voltage','3f1daace-2b8e-4d6c-a21c-1a23359ae00b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('52e946ec-7c9f-4e6e-b91d-427261c8a15f','stage','Stage','38fa06c8-6457-4523-b4ef-3ebffa520151', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d4a4c941-0de8-400c-9eec-f4b6e66283ff','stage','Stage','b9e94249-d8e9-41ac-b683-bde39fa79ed4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('104e1d23-e690-4eed-a250-ea610e3935aa','stage','Stage','616fa910-eabd-4718-940f-c7ec2afeffd8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4d1e443d-f6ef-4e0d-87f3-3e963a2eaeca','voltage','Voltage','616fa910-eabd-4718-940f-c7ec2afeffd8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dd625b27-2225-4eae-95e7-beac677e3ca2','unknown-us','Unknown US','616fa910-eabd-4718-940f-c7ec2afeffd8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('be23519a-7a5d-494a-b25b-25d2a672b552','unknown-ud','Unknown UD','616fa910-eabd-4718-940f-c7ec2afeffd8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b315787d-d06b-4cb8-adb2-f505b8258aaf','precipitation','Precipitation','616fa910-eabd-4718-940f-c7ec2afeffd8', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('2956dd3a-bb82-4669-9d66-65aae28c406a','unknown-ta','Unknown TA','616fa910-eabd-4718-940f-c7ec2afeffd8', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f8d8d6a5-1b08-49bb-bf6e-61027f1a75d2','stage','Stage','6d781aa1-93e5-4aa5-85fe-9d057ead4247', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9c4349e8-37dc-494b-8af7-5b74ec106d31','water-temperature','Water-Temperature','6d781aa1-93e5-4aa5-85fe-9d057ead4247', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('bb4da4c5-ebac-4e13-9f5d-33981d953794','voltage','Voltage','6d781aa1-93e5-4aa5-85fe-9d057ead4247', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('806b9958-4efd-4385-9356-0e5e1c96e102','stage','Stage','4adb6c37-4ae1-43e1-9540-e5595bf86d1b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('25666a77-a034-46d5-bef2-878530da2e67','stage','Stage','67cf3e31-f3de-4413-930d-1cf4903c454e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cb104026-aa96-47fa-b4ec-e722224e19df','unknown-pa','Unknown PA','67cf3e31-f3de-4413-930d-1cf4903c454e', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b73111c2-8e6f-45d8-adc4-ec28fe9ffba9','stage','Stage','bacf256e-eb96-4c36-8954-bd92a306f4d8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('77e9416d-ec76-4d8a-9fdd-6156a470863d','voltage','Voltage','bacf256e-eb96-4c36-8954-bd92a306f4d8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c397f2e8-4669-46ae-ac9d-532daa78d2f9','stage','Stage','d7281dd7-778d-4c3a-90a4-b21bb641e850', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2b21285d-2b43-4ddd-baf3-5a96c568fc3b','unknown-pa','Unknown PA','d7281dd7-778d-4c3a-90a4-b21bb641e850', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('03552083-0c40-4921-95c0-738ce4e2c276','stage','Stage','f6e8916f-296a-49e9-aca1-82665748ca62', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f14c5555-dd95-4217-8405-180c7659a686','turbidity','Turbidity','f6e8916f-296a-49e9-aca1-82665748ca62', '3676df6a-37c2-4a81-9072-ddcd4ab93702', 'daeee256-c762-43a2-8369-2d295525023c'),
('5e72dfd4-ab64-4cad-a427-bc0815b6efcd','unknown-ws','Unknown WS','f6e8916f-296a-49e9-aca1-82665748ca62', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('22876827-ad66-40cf-b8e1-b1d2943bcca4','precipitation','Precipitation','f6e8916f-296a-49e9-aca1-82665748ca62', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('14ee97fa-dff6-40fe-86a7-3a6c381b64ea','voltage','Voltage','f6e8916f-296a-49e9-aca1-82665748ca62', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('af8e441f-628c-487c-bf9a-53c791045755','stage','Stage','70ba9edc-40ea-42ad-a8a3-ceabb25b05e2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0865084e-96d9-43a2-ac91-bef460917077','unknown-pa','Unknown PA','70ba9edc-40ea-42ad-a8a3-ceabb25b05e2', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6a832fc5-ba47-4a01-8486-0b8483098afe','stage','Stage','694b13bf-449a-4c44-8aa2-c38469c07ed5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('80ca5f2f-4929-450e-a033-6caceced07ca','unknown-pa','Unknown PA','694b13bf-449a-4c44-8aa2-c38469c07ed5', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('edde0aa0-0afd-43c6-80c4-e3b8a3c8b823','stage','Stage','a392c3e8-738d-4feb-afc1-00d064e4298e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2ce1c3a5-e7f7-456a-a563-ef4e89f989b5','precipitation','Precipitation','a392c3e8-738d-4feb-afc1-00d064e4298e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('bef145ed-93d8-48d3-a2e6-0d8ff9421ed2','voltage','Voltage','a392c3e8-738d-4feb-afc1-00d064e4298e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('df44c2e2-3977-4958-bb12-9ea1a4124eca','stage','Stage','8add8896-c24d-44b3-9853-c189d2d83695', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cc80f76a-fd22-4cdf-8e69-d622b6156ec8','unknown-us','Unknown US','1592e55f-ff35-40a9-80c4-7c81e4f620b3', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e299b7d1-d0e4-4c15-bb9f-4ab22d64354d','unknown-ud','Unknown UD','1592e55f-ff35-40a9-80c4-7c81e4f620b3', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ddc3d449-c906-44e6-bee0-9c683bc438ca','stage','Stage','1592e55f-ff35-40a9-80c4-7c81e4f620b3', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0ca4efad-0f56-4198-a571-879ead4f35ea','stage','Stage','3604b451-24dc-423f-9e8a-f0bbf9d7d8f2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5fed354c-ca25-4468-bcb9-fb346ede63bc','voltage','Voltage','3604b451-24dc-423f-9e8a-f0bbf9d7d8f2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a41d5b55-66b7-476f-943e-b24012742f56','stage','Stage','9d3ea29e-59a3-46dd-a091-d8fed56f03ec', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8d70a401-75e4-4ca8-90dc-d03c3b7b6a11','stage','Stage','4c7ce601-f232-4634-936c-d7604eb54e3b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('54125d1f-38e5-46fb-8928-4be65ab7a51b','voltage','Voltage','4c7ce601-f232-4634-936c-d7604eb54e3b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ab6e912c-54a0-420b-86f8-a97a3d4a0412','stage','Stage','e213714d-2656-4719-84f5-61e208032038', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0bad177e-ae2a-4597-b0dc-0f80a25365d6','unknown-us','Unknown US','e213714d-2656-4719-84f5-61e208032038', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('c5fa4e39-3510-480b-8b8f-fe512d66e751','unknown-ud','Unknown UD','e213714d-2656-4719-84f5-61e208032038', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1b594c1d-1a8b-4f66-a76b-3b624ed56140','voltage','Voltage','e213714d-2656-4719-84f5-61e208032038', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2c9ea688-f7ad-4a1d-afb9-70402e23bac7','air-temperature','Air-Temperature','e213714d-2656-4719-84f5-61e208032038', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('7075d4f6-5ac8-4c65-a69f-b4bd1ca15ca8','unknown-xr','Unknown XR','e213714d-2656-4719-84f5-61e208032038', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('66149ab1-d598-4773-99dc-6939ba700fa8','unknown-pa','Unknown PA','e213714d-2656-4719-84f5-61e208032038', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('7ae34d90-00ae-4cd6-a1e6-f63fe525fd4c','stage','Stage','6bdca601-8044-4830-85e5-7f34dd5d5865', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3fa6553f-bbb1-4261-8c39-ce5bed0554eb','unknown-us','Unknown US','6bdca601-8044-4830-85e5-7f34dd5d5865', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('85947bac-0afa-4642-a4ad-cc9eaecaa007','unknown-ud','Unknown UD','6bdca601-8044-4830-85e5-7f34dd5d5865', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ab4c6022-ead9-4ac7-a8a2-cfbd63a3dfbd','voltage','Voltage','6bdca601-8044-4830-85e5-7f34dd5d5865', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('702d295c-4edd-4a99-a296-66f7fd55ea11','stage','Stage','75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fdbf525b-7e2e-493b-a963-e47f07862957','voltage','Voltage','75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('afa2f281-4f2c-4797-89e2-043027e5a99e','unknown-us','Unknown US','75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('c4d36c90-9625-4618-99fa-a8ecbe9e4bc4','unknown-ud','Unknown UD','75e2eee7-51ce-4c21-b0fc-e039c9ef38d2', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('9248fdcb-8351-4832-ac57-1198498260c8','stage','Stage','8091654d-2afd-46a6-a2a7-26447b8c0df4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fadd8eba-b2a2-4913-b17e-76ca542666ef','voltage','Voltage','8091654d-2afd-46a6-a2a7-26447b8c0df4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('bebe0d7f-1631-4ed9-aa8d-38187f84a898','stage','Stage','2cce0e9d-c608-4424-a8f8-0f76e0691338', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9417bd53-954f-4c20-b85b-8b3e490e0c6e','stage','Stage','82d2991c-71db-4699-9e7f-5506f9614d58', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c266495c-6f08-4ac2-b168-351e5d6332e9','water-temperature','Water-Temperature','82d2991c-71db-4699-9e7f-5506f9614d58', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('24fb5d59-7f69-47c4-a91d-cd2ff778270d','stage','Stage','21bdce69-d2a7-49f0-807b-d29c747f2802', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e5b1fb47-4e21-4715-a27f-c18279cdcd3a','stage','Stage','28bc2a6e-72a9-4e21-9da0-29ce3dda495d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4abf6619-5d67-4d39-987f-5f6faea57ad7','stage','Stage','00fb2016-699b-45cf-a44f-2944e0860c26', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8de63cba-9e8c-4c5a-b98a-4eede9cfc846','stage','Stage','df1cf782-4230-4f79-8437-4e361ee67dfc', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a106e1fe-9a63-44c3-8189-b2ec57598fc8','water-temperature','Water-Temperature','df1cf782-4230-4f79-8437-4e361ee67dfc', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('c5a57028-da2a-44a5-8e52-b17f6397d3ba','voltage','Voltage','df1cf782-4230-4f79-8437-4e361ee67dfc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ac70133e-8cf6-4067-9bef-9752839af8ce','stage','Stage','f0658ffe-4b1f-4b6c-b17b-7b8cc43ad302', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f45bc626-0468-4227-b448-4e9c4462bb87','voltage','Voltage','f0658ffe-4b1f-4b6c-b17b-7b8cc43ad302', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('81bfa5fc-0ecc-443c-86fb-509251ee5e94','stage','Stage','581ba9dc-3836-4fea-9598-fc0d9255da47', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f247f768-4ad6-41ec-85b0-721ba83a6a80','unknown-pa','Unknown PA','9ce6e98b-a26b-450e-b859-41bdc05bab06', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('c74b658e-c141-4c90-bc79-7e120d099a53','unknown-us','Unknown US','9ce6e98b-a26b-450e-b859-41bdc05bab06', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('7a396afc-bf9d-41a7-94c2-909e84d0998a','unknown-ud','Unknown UD','9ce6e98b-a26b-450e-b859-41bdc05bab06', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('fb24af9c-eaaa-42b8-845c-c30198599d6b','stage','Stage','9ce6e98b-a26b-450e-b859-41bdc05bab06', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8e228b17-4bed-4f88-9da2-fdefd8d6a69c','water-temperature','Water-Temperature','9ce6e98b-a26b-450e-b859-41bdc05bab06', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('46473d67-8d84-46d0-acee-e82975a1beb1','precipitation','Precipitation','9ce6e98b-a26b-450e-b859-41bdc05bab06', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ab581292-7ebf-4bae-8ff7-9685d48ad11d','voltage','Voltage','9ce6e98b-a26b-450e-b859-41bdc05bab06', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('344b5198-2a3d-416c-b7fe-3ff41516f56c','unknown-pa','Unknown PA','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('078e7890-3f9d-40ea-8a49-5c33396905c0','unknown-us','Unknown US','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('499e1f78-9c0d-4ed4-aa00-27cec5489d28','unknown-ud','Unknown UD','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('acc25b45-1f69-4f49-a51f-acd01f180dbd','stage','Stage','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ae1d506a-72e4-4ad5-ad9f-e2c259af4dd5','water-temperature','Water-Temperature','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('acdef086-cee2-4bf7-9603-a442f400fb73','precipitation','Precipitation','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ef1d330d-945f-4857-81ec-4f015427c337','voltage','Voltage','1cc92566-d593-4c0f-abd5-fd39aabf9fb1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4509340b-1f32-48fe-923f-9e03f9c28128','stage','Stage','5c3f1f19-1639-477f-81e0-8edf1b85bda7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('71c20ce5-9a79-4076-b136-c894f2fb8b36','air-temperature','Air-Temperature','5c3f1f19-1639-477f-81e0-8edf1b85bda7', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('bd468e98-4df0-43eb-9c16-251942ea24cd','unknown-xr','Unknown XR','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e90a09cd-941e-440f-9e9b-b732b71a3132','unknown-us','Unknown US','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('77a90bcd-b7b7-4c1e-8e6d-e71e91603648','unknown-ud','Unknown UD','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('21570d87-1cd4-424a-bcf4-9bee83622762','unknown-pa','Unknown PA','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b3ecb1a2-eefd-42c7-a690-2c79f1225d96','unknown-pr','Unknown PR','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6d31acdd-f48d-49c9-932b-9d495064f6a8','precipitation','Precipitation','5c3f1f19-1639-477f-81e0-8edf1b85bda7', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c914b9e1-56cf-4274-8010-a9525dacf939','stage','Stage','449c9725-605c-4d0d-a0c6-8e434054b26d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d3427489-ee63-45f8-95d3-107215b96536','voltage','Voltage','449c9725-605c-4d0d-a0c6-8e434054b26d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5b47d219-85d4-47a3-a79f-a53a6fbb2bc3','stage','Stage','a70a73d7-47fa-4f92-9b00-6bf9f02c52df', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('167d3ccf-6d5e-48a8-8611-cf382cdb97d4','voltage','Voltage','a70a73d7-47fa-4f92-9b00-6bf9f02c52df', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e86ca313-1fa5-44bf-948e-645c9c268148','stage','Stage','d56068c3-a3c4-4553-b30a-a885570cd7d8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('94db37c1-5bcc-4678-b0a7-4197fd255dff','voltage','Voltage','d56068c3-a3c4-4553-b30a-a885570cd7d8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c5d61c58-3494-47f5-90c9-d075ff91c129','stage','Stage','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e8c0964f-dff0-466c-9814-de4c973d3790','air-temperature','Air-Temperature','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('c0b5927d-ecd5-4bce-90e6-d089767afd27','unknown-xr','Unknown XR','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('cc8ac547-ed63-47c3-b3cf-e02a0516cba2','unknown-us','Unknown US','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ed3f66cf-3c80-4145-aaa9-ce9f08d0358f','unknown-ud','Unknown UD','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e748f1a1-3b4b-4ea6-b861-9ebd3f4260c7','unknown-pa','Unknown PA','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('29728886-bee0-4818-a9f2-d9f8d0f86771','unknown-pr','Unknown PR','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('bb56ee7f-aa9b-495f-a39e-375aa18f49ed','precipitation','Precipitation','6fb38d2f-9428-43ef-a45f-d9dc90c7bedc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d27345e6-56e1-4dfd-bc96-b7ea0c3473ea','stage','Stage','572ce69b-8b32-40ab-ac00-e783dc6ceb4f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5dd58d24-d23b-4cd3-ba3e-18be90be25c7','voltage','Voltage','572ce69b-8b32-40ab-ac00-e783dc6ceb4f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('df74e1c9-12d8-4263-a37e-9ac3682e1bab','stage','Stage','bcfb5c23-300d-4d78-b8b7-f7cf40bfbd09', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0b0b708a-5f48-4445-988b-4d49fc4850b3','voltage','Voltage','bcfb5c23-300d-4d78-b8b7-f7cf40bfbd09', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('536ebf5e-d3d3-460c-9ac9-474e8d17e96c','stage','Stage','0d2b2c4a-e8ab-464d-aa5e-5d4fb3145dc5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('daf0a700-f89b-48f1-8c2f-14d14d009f12','voltage','Voltage','0d2b2c4a-e8ab-464d-aa5e-5d4fb3145dc5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5054dbf4-7fba-4f32-b052-d9a48741c3d7','stage','Stage','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3f5a1066-228a-4d92-b3ac-b226b2d09512','air-temperature','Air-Temperature','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('0e449e31-27fa-4178-b6f7-b03545302cb8','unknown-xr','Unknown XR','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e985d553-b2ca-4b43-98f5-6b7c66fa6b4d','unknown-us','Unknown US','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b5506bc6-82bd-4200-95ee-a46a7c189b1f','unknown-ud','Unknown UD','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('04a8f91d-2d7d-425f-8b81-b09071d4178b','unknown-pa','Unknown PA','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('00fc1b03-a5f9-410b-9aff-4b44e398aff5','unknown-pr','Unknown PR','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6aa04c42-b24b-4d01-b7eb-d66d93c1656b','precipitation','Precipitation','d0803aa7-bbee-48ed-b23f-7dea74b2ac59', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f85cd014-7d19-4ccd-a498-a3458fab48e9','stage','Stage','5e19aa7a-1611-49e9-bec3-75c3421c7b5e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d29ffdb3-54ef-474e-b1de-e11468fcd7be','voltage','Voltage','5e19aa7a-1611-49e9-bec3-75c3421c7b5e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ea191ec7-e81c-448d-bed4-84404b31679c','stage','Stage','e1993ab5-ec08-4e07-bc3c-c4db92731d1c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f016e93c-4972-4675-baff-42f6aff78fb1','voltage','Voltage','e1993ab5-ec08-4e07-bc3c-c4db92731d1c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('aca090ae-3375-49da-9ac5-db3b55913114','stage','Stage','2db45595-fefc-4eb8-9237-8fc72561b47e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dacd1a2d-56c5-477d-9395-47a96dadc2e7','voltage','Voltage','2db45595-fefc-4eb8-9237-8fc72561b47e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('7da16127-aeb6-426e-a191-3d70f829ea72','stage','Stage','ec40c087-d75b-4580-9ee5-337eeff32f95', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3386f093-572e-4912-b5fe-251d62192683','unknown-us','Unknown US','ec40c087-d75b-4580-9ee5-337eeff32f95', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('28f9272a-07d7-4c4f-a89b-16f3df1a26d0','stage','Stage','f3b99e3b-1856-4226-a73c-1c46b200fb66', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6eaa7f4c-cd3f-4571-9ab3-e9701b209770','voltage','Voltage','f3b99e3b-1856-4226-a73c-1c46b200fb66', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f5483847-71fe-4181-a84d-048d813325e0','unknown-pa','Unknown PA','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('13378c5d-593d-4bf3-af32-72bb0df4c32f','unknown-us','Unknown US','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('75e519d8-359d-471f-a0b6-c45f5416c0ba','unknown-ud','Unknown UD','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1ab701ad-28a8-412d-9132-7efa06f4b3a5','stage','Stage','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('00884f48-3fa2-4a03-a703-d907aa9b04d5','water-temperature','Water-Temperature','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('c1019923-f5c3-4ee7-9e27-4fdf59714b9e','precipitation','Precipitation','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a3cec8bb-f4a8-4432-a6c7-ce3011066ad9','voltage','Voltage','72d33561-aaaa-4e90-91b4-eb73eaeaaac6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('323ed904-dc1b-4bd7-a0ad-94156665c0c7','unknown-pa','Unknown PA','6b78716c-4975-46e4-b635-9ff79e7f0627', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e7c4e557-5a17-4290-af5e-ce3fa4e08d26','unknown-us','Unknown US','6b78716c-4975-46e4-b635-9ff79e7f0627', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f62b049e-dae6-428a-b684-62a3d738fb22','unknown-ud','Unknown UD','6b78716c-4975-46e4-b635-9ff79e7f0627', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ff96d1f2-9f25-400c-b5b6-34909b133fd3','stage','Stage','6b78716c-4975-46e4-b635-9ff79e7f0627', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fd856b45-8e60-4b4d-8811-8a1521b627b0','water-temperature','Water-Temperature','6b78716c-4975-46e4-b635-9ff79e7f0627', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('b22d6482-db00-435b-919b-d3d9e63b3ebe','precipitation','Precipitation','6b78716c-4975-46e4-b635-9ff79e7f0627', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ece3d621-1b07-431e-81bf-a173eabe43c2','voltage','Voltage','6b78716c-4975-46e4-b635-9ff79e7f0627', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e6509eab-7d83-4c34-a201-57736cbd2b3f','unknown-pa','Unknown PA','a88ece8d-d3ba-4a77-a413-5dd03ebfb623', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6bb153d6-f225-4f2b-8bfe-0982eaf94529','stage','Stage','a88ece8d-d3ba-4a77-a413-5dd03ebfb623', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('092f58ad-4f9e-440a-bbf1-d273bf541929','unknown-us','Unknown US','f6629ae2-98e2-4e06-b685-331189cff3a7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6ba46a97-c962-48e1-a19c-ea205fd089e0','stage','Stage','f6629ae2-98e2-4e06-b685-331189cff3a7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('322dafe5-1ab6-4624-8737-22e6fd31ddf1','unknown-ud','Unknown UD','f6629ae2-98e2-4e06-b685-331189cff3a7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f813e32d-b967-4352-9a54-4bb07b1ce0a2','stage','Stage','59db968a-43d1-4590-b4f7-6f7f6423e735', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6bb36470-a055-4be5-9156-39f7f9aae837','voltage','Voltage','59db968a-43d1-4590-b4f7-6f7f6423e735', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2bee52bd-6385-4f50-bb3b-701a62ef8949','stage','Stage','ac71a2f1-74bb-49de-8dd4-d2b15b974978', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b627ab9e-b958-499f-9dea-157c8a5ad0eb','stage','Stage','ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3262460e-08b9-43ec-9af5-40ef96656204','unknown-us','Unknown US','ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1d4b18b6-a405-4935-93b3-4b802aa5934c','unknown-ud','Unknown UD','ae73a827-b0b9-4fa6-83fd-fa7d78e22f5c', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('6478f569-f7d1-42e8-8fa3-04c36b8644a9','stage','Stage','4916c4af-8019-4492-a734-fc3ff8c49074', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b5b009de-7ea9-4e05-9438-a2a74a78cbc1','water-temperature','Water-Temperature','4916c4af-8019-4492-a734-fc3ff8c49074', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('978a6a18-7273-42ae-a781-fbe52028ec37','voltage','Voltage','4916c4af-8019-4492-a734-fc3ff8c49074', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3607c1be-bebf-48c0-b0ba-b90b4e9b0f58','stage','Stage','5dfc826b-a94c-4cee-b057-6fb858afc95d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f8e8b98b-da6c-4abf-8b39-610041144a21','stage','Stage','7e772c75-64b6-485e-8294-d1d3bc6da0fe', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('781c071b-bf74-4976-97c5-8d597e0929f3','stage','Stage','3fcc1834-0263-4dc5-a7c8-9157c767bb4f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d5cd5dee-4ad9-4b81-9693-1ad374da9bd2','stage','Stage','6b5d1a5d-0861-45a1-b56c-ec37c501e993', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9cc39ed9-53fd-4d44-bdab-8348c7339f4b','precipitation','Precipitation','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('97753d46-3b72-4754-ae43-eee17e7ce6cc','voltage','Voltage','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('91a0f363-b7a3-4ee9-b9b8-e70cc1b181e4','unknown-72114','Unknown 72114','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('d396854a-cf75-40e9-afa0-08c514ec3af7','unknown-72113','Unknown 72113','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('02f0d09e-c320-4ee2-bc7a-f53fe826ecf3','unknown-72115','Unknown 72115','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1a5eadf9-8d8f-4f53-ac67-fc36f1ae8eee','unknown-72112','Unknown 72112','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('8446591a-8f83-47e5-8412-c677f31a9535','unknown-72111','Unknown 72111','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('12861a31-6fe2-4b84-8ec0-41cb4dd68358','unknown-82292','Unknown 82292','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('30990d93-19e4-4163-a29b-703db38ad746','unknown-72117','Unknown 72117','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('722a3c37-cad7-4a83-8097-2c3442c9f624','unknown-72116','Unknown 72116','6b5d1a5d-0861-45a1-b56c-ec37c501e993', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('96e1b66d-c006-4e5b-b977-9e5a11066703','stage','Stage','6890212a-c7f1-499e-a3fe-e7d973eb7cc1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8c9697f4-9228-4674-bdd3-4e65a3b08e93','stage','Stage','bb1389a0-bfad-430d-a66d-2f041a738cb1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b713b4dd-d090-459c-808d-2a6b221616ac','stage','Stage','38ea9864-0181-475e-a8c7-b34edbf31b6b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a7d33ef6-0101-4679-8146-6bdf01525f1b','water-temperature','Water-Temperature','38ea9864-0181-475e-a8c7-b34edbf31b6b', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('088d5a4f-ef9b-4a3e-bdea-c6d5133824ae','voltage','Voltage','38ea9864-0181-475e-a8c7-b34edbf31b6b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('174a7bf2-1925-4087-ac69-32227935f317','unknown-72020','Unknown 72020','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('973c7bd9-52a2-4047-9ae9-6131e5b3b6e5','unknown-hk','Unknown HK','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('73f30a96-dfac-4f24-af5e-3e50635a3888','unknown-72114','Unknown 72114','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('e039701a-00b9-4b2f-8c5f-6a98c81cfdd8','unknown-72113','Unknown 72113','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('82692453-c591-432a-9d47-7a9144d3d801','unknown-72115','Unknown 72115','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('16627fc3-caff-4704-805d-12e6e6aeaaa0','unknown-72112','Unknown 72112','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('d2329a53-9026-4b12-a6e3-563445e41402','unknown-72111','Unknown 72111','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('19dfa183-8538-4625-8663-dd3b5e66449a','unknown-82292','Unknown 82292','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('22c38628-f6cf-4a36-acaa-1701ef25d63f','unknown-72117','Unknown 72117','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b4cec86d-af92-4340-b40d-a525830a5ddf','unknown-72116','Unknown 72116','38ea9864-0181-475e-a8c7-b34edbf31b6b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b98ed7ea-93be-419e-8fe8-2248762146c8','stage','Stage','b417e9f4-ce46-47a7-8bd2-0b26506395e4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('35968a24-dff1-4bdd-8561-53ba8f2237fc','stage','Stage','04e220cf-128b-4115-8a7b-76ac2625a81b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dfbbb59a-ee32-4aa3-8fa1-ee306381b067','stage','Stage','7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('91dd1823-f7f5-4899-9759-7f5c18467cf3','precipitation','Precipitation','7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('0cf9ab90-d54b-4f5b-8a93-32ede486d9a3','voltage','Voltage','7a6b4f10-c1d7-4301-89f5-cbadcc1e255e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1842b351-e961-4353-9426-c863d9b4dfde','stage','Stage','892ca8ff-08f2-4e1b-af40-32a77baab497', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('63c680cc-1681-4bc4-9f3d-20c5b6002194','voltage','Voltage','892ca8ff-08f2-4e1b-af40-32a77baab497', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('84cb2361-e1a4-428f-80cb-034777fa1f7f','stage','Stage','8250ac38-3d9f-4039-9f08-149912301977', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('32ab5b47-2c1c-46e9-b2e4-e84cf3996133','stage','Stage','9dcf334f-2982-49c9-941c-38c91565f019', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ba002706-819f-4f5e-b3fe-62e23b91c5da','precipitation','Precipitation','9dcf334f-2982-49c9-941c-38c91565f019', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('82d6c95f-2c02-4d08-a219-92d118013524','unknown-vx','Unknown VX','9dcf334f-2982-49c9-941c-38c91565f019', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1258a8d9-e348-44fa-aba3-3d78522c7660','voltage','Voltage','9dcf334f-2982-49c9-941c-38c91565f019', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('bfaa0312-9206-43d4-a516-67e0dd2d9f4a','stage','Stage','df1b2707-3dea-4b91-be0d-2f64f21544e2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('07a3759d-3403-4a5d-bffb-6390ef6c1d12','stage','Stage','fcf546da-e8c6-471b-b894-98a691c71251', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dff305e1-cc3a-4117-b225-22fd14b24a87','water-temperature','Water-Temperature','fcf546da-e8c6-471b-b894-98a691c71251', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('69fbf02b-dac7-4ba4-8271-803b7fa5d094','voltage','Voltage','fcf546da-e8c6-471b-b894-98a691c71251', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('5a5298b5-1b69-4331-bb56-3549866e35d9','precipitation','Precipitation','1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('93217695-e52e-4b3d-9aca-86a030842ea5','voltage','Voltage','1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0bb9034e-27b4-400b-848c-25d644f904e6','stage','Stage','1aafd0d7-eca1-4237-8d5c-6bcf24b31bbf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c07f8f90-804f-4f6d-9b34-b074df75dd39','stage','Stage','3eba9840-cebb-45ee-8e9f-a62d8279c86a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('846c1b39-efa7-4b27-9cf2-8d7cd49b237b','voltage','Voltage','3eba9840-cebb-45ee-8e9f-a62d8279c86a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f4123631-b2f8-48f9-99e5-0571aa4d14f1','precipitation','Precipitation','3eba9840-cebb-45ee-8e9f-a62d8279c86a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('55a9e7f3-d03a-4310-967b-b134f6eec42e','precipitation','Precipitation','bc7d2d9a-8e02-4a78-b838-cadd444ed756', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d8031923-eefb-40f6-a4ab-62650af7c8b3','stage','Stage','bc7d2d9a-8e02-4a78-b838-cadd444ed756', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('22cb1064-a089-41c4-b5bc-b06de7b79589','unknown-qr','Unknown QR','bc7d2d9a-8e02-4a78-b838-cadd444ed756', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('eb724b60-1808-4178-adef-d5c369a5d31d','voltage','Voltage','bc7d2d9a-8e02-4a78-b838-cadd444ed756', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d114a8b3-f974-49ea-9b96-27778c71758a','stage','Stage','3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('11b67431-5bcd-42fb-b81a-4e84df58d467','water-temperature','Water-Temperature','3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('b2806dca-6a6b-4769-ba2e-fb5fdaa9fbe2','voltage','Voltage','3a4dc536-bbfb-4891-927f-c5cb4e2bb3e5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('00e09434-b7c0-4b9c-85c2-d26ea80d843e','stage','Stage','5171ef3b-9ef9-4984-8ab7-57b76db13621', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('af8359b2-0402-4ff0-a5b0-096f651b02d0','unknown-pa','Unknown PA','5171ef3b-9ef9-4984-8ab7-57b76db13621', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('38c4dbd5-c2da-4b3b-a967-d4d759de871b','unknown-us','Unknown US','5171ef3b-9ef9-4984-8ab7-57b76db13621', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('2ac1b729-c23f-4ef4-b7d0-a1a4b4cd4731','unknown-ud','Unknown UD','5171ef3b-9ef9-4984-8ab7-57b76db13621', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('321c1125-9a25-4918-8a19-306f83b2d732','air-temperature','Air-Temperature','5171ef3b-9ef9-4984-8ab7-57b76db13621', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('3714bafb-60f6-44ec-9b36-14c6fbbf9efa','unknown-xr','Unknown XR','5171ef3b-9ef9-4984-8ab7-57b76db13621', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('65a5374c-de46-437a-96d2-26fd2ff60b7e','precipitation','Precipitation','5171ef3b-9ef9-4984-8ab7-57b76db13621', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('37c2f358-8d57-467a-8970-67bca5e9dd80','voltage','Voltage','5171ef3b-9ef9-4984-8ab7-57b76db13621', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('eff22eb4-c99c-4c8d-8fdd-4a0130a32023','voltage','Voltage','b72006fc-5809-4985-8f3c-45baed7cfcb9', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('eef1d7d6-6a33-4371-b79d-9f67c108e7b0','stage','Stage','b72006fc-5809-4985-8f3c-45baed7cfcb9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cc762a94-e458-46fa-ab4b-d78cfb5fbd92','stage','Stage','4811e8d0-4c72-4c86-8bec-37e491f5ab63', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('91b535c5-9f56-4c56-90f8-2f25ddb12450','stage','Stage','3f9a0fcf-4e4f-4b0c-ac88-011455d6957f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1f4d0bc8-7549-467f-a553-8e40e06a0bb2','voltage','Voltage','3f9a0fcf-4e4f-4b0c-ac88-011455d6957f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3b424bd3-9982-46b5-afd6-9e3e0abfbba1','stage','Stage','553f696a-9292-41b1-a4bf-67236068a1d6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ab808845-2f37-4286-a4a2-db059d2736fb','unknown-us','Unknown US','394d6441-43b2-4745-b97d-c6462c7bc908', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f753d916-b87a-460e-bc45-712010736794','precipitation','Precipitation','394d6441-43b2-4745-b97d-c6462c7bc908', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('95f8f2ff-165a-443c-b458-5fba54f6c249','unknown-ud','Unknown UD','394d6441-43b2-4745-b97d-c6462c7bc908', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('4f0b52bd-3457-4bf3-9a85-c9f253b904e7','water-temperature','Water-Temperature','394d6441-43b2-4745-b97d-c6462c7bc908', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('7c39e9a1-49b7-4922-975b-405b36df4e4d','air-temperature','Air-Temperature','394d6441-43b2-4745-b97d-c6462c7bc908', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('01f56cf1-56e3-4e2d-af8b-593df54a9eb8','unknown-xr','Unknown XR','394d6441-43b2-4745-b97d-c6462c7bc908', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('755b91e6-aa19-4cfc-a739-11c72e214203','stage','Stage','394d6441-43b2-4745-b97d-c6462c7bc908', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('01d0f0fd-00e2-4393-be43-557be2669a4f','stage','Stage','2711323d-5ed2-44b5-a7a6-f0614c0b422b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9e6e581d-6d08-4c35-930c-8d4ab84297dd','unknown-pa','Unknown PA','2711323d-5ed2-44b5-a7a6-f0614c0b422b', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f3ccfdbe-19e9-404e-94a3-10fb55037152','stage','Stage','3bd76812-033f-4367-b8c4-eceba262f393', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('958543d3-2eb7-40b3-bda8-33492d2d9273','unknown-pa','Unknown PA','3bd76812-033f-4367-b8c4-eceba262f393', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('63b34e59-4676-4fe5-ac6a-44f420669798','stage','Stage','f8d6ccd6-fb48-4228-b777-e423612b1f27', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dc1ae570-d5bc-49d2-9a1e-f295c00bc2af','unknown-pa','Unknown PA','f8d6ccd6-fb48-4228-b777-e423612b1f27', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('34e9472c-8f44-44a0-a63b-1dd4e4bc4c1e','stage','Stage','7babc4e1-950e-41e8-874f-49559dd947bf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fc4d03f4-a4a5-462b-a8f1-6f085b09f27c','water-temperature','Water-Temperature','7babc4e1-950e-41e8-874f-49559dd947bf', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('ebd2c8ef-8a18-4e8b-a29e-3d481b66e661','voltage','Voltage','7babc4e1-950e-41e8-874f-49559dd947bf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0359a117-b2a8-48fe-8f8e-e7c6d58127d4','stage','Stage','ab7ee39b-8397-484a-ac29-e2e813ed6adb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fab565fc-17ea-4cc9-bd25-8304159fc908','water-temperature','Water-Temperature','ab7ee39b-8397-484a-ac29-e2e813ed6adb', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('14d7d427-7918-421c-8b33-57320c5ffe29','voltage','Voltage','ab7ee39b-8397-484a-ac29-e2e813ed6adb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('841e081d-f8f5-4884-b5d4-2060995ddab5','stage','Stage','d93718c4-08ae-4065-8855-674c5206c1ed', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('15b59556-9c16-4637-8a4e-feb26fbe3081','voltage','Voltage','d93718c4-08ae-4065-8855-674c5206c1ed', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('627de30d-afb7-4497-8ae2-c3be638a6ba5','stage','Stage','de20f637-d3c7-4ea4-8ca6-604f6a127532', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8ab0201b-b935-495d-87b6-f33d9eb66a2c','water-temperature','Water-Temperature','de20f637-d3c7-4ea4-8ca6-604f6a127532', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('45a3b3e4-5606-4f56-96d7-b5fc3b5a87a8','voltage','Voltage','de20f637-d3c7-4ea4-8ca6-604f6a127532', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('865ad98a-0e5a-4430-8290-511260e4e9da','stage','Stage','8a7739c0-ccae-4f3b-83d7-e2f6f506e6da', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('940c1bb5-1567-4190-9090-3de1cb2de60d','stage','Stage','2006cc18-de31-4ae0-8aaa-81ba6cf1ba95', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fd8f256c-21fa-4248-be17-a5560772cc51','voltage','Voltage','2006cc18-de31-4ae0-8aaa-81ba6cf1ba95', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('388d15f4-f71c-4c6f-b606-64e9b1acf467','stage','Stage','e29fcc6e-a0a8-4fe0-b450-099142997846', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3f52e10d-62ce-479a-9c68-6277b7f62b2d','voltage','Voltage','e29fcc6e-a0a8-4fe0-b450-099142997846', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f7a5623e-9318-4412-a6a1-3690850534c9','stage','Stage','7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3f49c928-59f0-4f40-b38e-13564b195ce5','water-temperature','Water-Temperature','7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('9f22e15d-e417-4a89-9896-3c8bacd3c919','voltage','Voltage','7a2db3f0-aea2-4f91-b6ec-23aa4d2a70d7', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4af23042-79b2-4884-bd02-0828811ac1ab','stage','Stage','bcc35574-bf23-4952-9760-049ba8843bd7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5428c796-de03-4f7b-b0a7-7c9c740795ef','water-temperature','Water-Temperature','bcc35574-bf23-4952-9760-049ba8843bd7', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('de78b46c-a1a2-4c72-897b-eff78a949640','unknown-vb','Unknown VB','bcc35574-bf23-4952-9760-049ba8843bd7', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('d1873492-a47c-44c9-bb27-115fe1e4e055','stage','Stage','4bbb3d28-b816-4f71-98c1-de647d8f5e38', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8f341737-52ea-4cb4-a8ab-38983107fc49','voltage','Voltage','4bbb3d28-b816-4f71-98c1-de647d8f5e38', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('76f0f3e5-0e44-43ec-b095-e887ecaabf6a','stage','Stage','b3524459-4cbe-4802-835e-b415cc6657b2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6f36923e-d277-4a5f-8f5e-4292ceddd2b9','voltage','Voltage','b3524459-4cbe-4802-835e-b415cc6657b2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4343b80f-8f2c-47ba-8844-44698ef9e774','stage','Stage','ed2c1d1c-7bb2-4374-8e90-878ab6b70d9c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d55baf27-5602-4dcf-853d-4445f3195698','precipitation','Precipitation','eaa406e4-560c-4fa3-9c28-8f265999f17b', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('0cf0e3ea-a705-4b90-83e5-1a3094041d51','stage','Stage','eaa406e4-560c-4fa3-9c28-8f265999f17b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cd36d05f-9808-4efd-950a-46f09b996da7','voltage','Voltage','eaa406e4-560c-4fa3-9c28-8f265999f17b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('9ce72c42-d003-48f6-a4e1-c30351915d5f','stage','Stage','bd3d7701-94b0-4390-b2f8-40e03c48508c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('13a9db55-1d9a-4fdf-bfc4-0229000ac00f','voltage','Voltage','8e5ed58a-febf-4b91-86e3-4b39d798a592', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('65d2194d-1125-46fb-96f6-b8c200f4dee4','stage','Stage','8e5ed58a-febf-4b91-86e3-4b39d798a592', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9625a180-64d0-40c5-8b08-ce30149adf3d','stage','Stage','33564f35-da99-4db5-9d9d-2c96698f99e6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('336bebb8-4bdf-47a6-9b7f-9f9222701484','voltage','Voltage','33564f35-da99-4db5-9d9d-2c96698f99e6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d4d1ca3d-ebc8-498b-afc0-2f758bb200be','stage','Stage','228d8881-b162-4842-92c9-75422535a5ad', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('caf194a7-2ece-4dec-b746-a19c6e1d5ca2','voltage','Voltage','228d8881-b162-4842-92c9-75422535a5ad', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('902d188a-3f29-4aab-8444-e6c4aaca0f73','stage','Stage','8f32236d-1e18-47f4-bf36-383eb8c82fbe', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('296e149e-c6ad-4959-abd3-ec33d6bab29b','voltage','Voltage','8f32236d-1e18-47f4-bf36-383eb8c82fbe', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('69130c04-1f1a-4155-8509-f9abfe0182c1','stage','Stage','dccf96c3-4dff-4b3b-a7b7-446a97db654b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('45729ce6-176a-457b-b8c4-54e57440673c','voltage','Voltage','dccf96c3-4dff-4b3b-a7b7-446a97db654b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f7ad1caf-dfb8-4ce7-b8b7-4b76aae7c3dc','stage','Stage','8527c379-3f95-4ef6-aae7-4fcf94ad8691', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a27bd242-2b5a-4dbe-b94f-f956f6a59da3','voltage','Voltage','8527c379-3f95-4ef6-aae7-4fcf94ad8691', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('7eecf618-1ea7-48b0-a70c-7bcff39d7264','stage','Stage','8bd798a4-1c7b-4684-ae3b-eb75a62a09cf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('50c093b2-334e-49aa-92cc-aabe4d566af3','voltage','Voltage','8bd798a4-1c7b-4684-ae3b-eb75a62a09cf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('35d176f4-01d0-4f9e-ba2f-abdf640b7543','stage','Stage','2eed3135-b0ee-423d-b8a6-5336072ac52a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8123d663-dccc-4501-82f1-66acc96f7a67','voltage','Voltage','2eed3135-b0ee-423d-b8a6-5336072ac52a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('90c41d18-cfd8-4a7e-938f-a28ef6982b55','stage','Stage','3277950f-fca8-4bcd-8799-64570ea99bd6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('65aaf529-2da0-410d-bc8e-00d79b49e37c','water-temperature','Water-Temperature','3277950f-fca8-4bcd-8799-64570ea99bd6', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('c3c966e6-4b9c-4759-a443-1dc4e41b9472','voltage','Voltage','3277950f-fca8-4bcd-8799-64570ea99bd6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2cd6f73d-6a32-4457-a37c-8da5469c52cb','stage','Stage','9bda847a-a23e-430c-bbbe-27da4e67d463', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5605e302-902e-4338-befc-31f0de0c0148','voltage','Voltage','9bda847a-a23e-430c-bbbe-27da4e67d463', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f2df7bf9-1f78-4687-a1b4-32206e517a30','stage','Stage','9b19c51c-b11f-4d74-9eab-2228dc25508b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bf6fe069-6134-41bc-9c1f-b0678d9a245b','stage','Stage','8965cb62-506a-4f0e-b285-48fe8ddfe35e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bff2794b-4cab-4313-99bd-207615868fe8','stage','Stage','2432c199-8996-4ae7-8a30-50c52c9f746e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9dabffa4-a077-49d7-8640-369281f9a17c','stage','Stage','93960641-3b74-456c-8fcf-7aadbf48c6d8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('17092127-980d-4e13-932a-2ded8e1a2d7e','water-temperature','Water-Temperature','93960641-3b74-456c-8fcf-7aadbf48c6d8', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('ca2407e3-02c7-4ff1-bbba-465cfbf71049','voltage','Voltage','93960641-3b74-456c-8fcf-7aadbf48c6d8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('b10fabdd-e88f-47f2-acc8-6e5eea30d24e','stage','Stage','1e6a3e9c-3b71-4256-a4b7-75a9615ba8b1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5d948f0f-d192-4aa1-b27d-fba062db2ab3','water-temperature','Water-Temperature','1e6a3e9c-3b71-4256-a4b7-75a9615ba8b1', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('a1bbc3e8-58e8-4d3b-b2d3-a990d0655404','stage','Stage','d56eb333-2749-4797-b5c8-f282ecf8fd4a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e07ffd2d-c566-4bb3-9b50-dbd0dd006a6b','stage','Stage','3ba8a61e-8920-46d2-9831-8cd0713508ef', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('967e940a-ce43-41bc-ad51-3b4591a0765b','stage','Stage','7e52d2f4-c471-400e-9f1f-8d2f50b4d6e4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('71655b54-91b8-4506-91ae-7086d1fc9aeb','stage','Stage','b72a87bd-94c0-48af-b773-d9bdc5419859', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('7524867b-2980-4cab-b2ba-13f208af4c9b','stage','Stage','69d3de52-e320-46e4-8c5b-f29cb8cbb4bb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4c7f92b8-3f06-4d4d-aef2-8e039be0da3f','voltage','Voltage','69d3de52-e320-46e4-8c5b-f29cb8cbb4bb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('af433098-45f4-4642-9221-1f3db1f7e41d','stage','Stage','b4cd105a-dd1c-485f-bbdb-b80269084676', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('634aa99a-5edd-47f9-8c17-d1f9edb8046b','stage','Stage','d4682aed-6cd6-498c-a610-3d565b03c971', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f53dfc3c-7700-4013-99b1-9993388e434a','stage','Stage','017d8b2c-2234-412e-a88c-d5949ce2fd02', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8fc8c361-9f14-4e06-8545-270f534a2899','water-temperature','Water-Temperature','017d8b2c-2234-412e-a88c-d5949ce2fd02', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('69330748-282b-4f2f-b932-99020bff94f7','stage','Stage','333d17db-11e8-47b9-8621-af2fccce7190', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('255b1a08-45cb-47c7-9ca4-ed43ed87986b','voltage','Voltage','333d17db-11e8-47b9-8621-af2fccce7190', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e4cb61bd-0d42-4008-b9ad-9b5bcc86c5f2','stage','Stage','1cff71c6-56ec-4baa-8524-ddcc2b292f25', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4dd72f31-2396-475e-87d3-c8ae1a5989b2','stage','Stage','afda57df-9971-4d0d-8d7f-7092d94d1791', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4f044706-80a2-4aa1-a4a3-1ebca34587e3','voltage','Voltage','afda57df-9971-4d0d-8d7f-7092d94d1791', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c74e72c2-d24e-4a3d-80cf-bc63809c0cd5','stage','Stage','7b94b33a-0922-42ab-9928-2d4dfe40aa22', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a6c5591b-8198-4038-abca-1e214968d800','voltage','Voltage','7b94b33a-0922-42ab-9928-2d4dfe40aa22', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('673623b6-f3aa-4110-8462-a8b061f89099','stage','Stage','d2beb71a-5aeb-4ba3-ae0e-346e81943d97', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5d08f88f-49b7-4de9-8e1d-818726180671','water-temperature','Water-Temperature','d2beb71a-5aeb-4ba3-ae0e-346e81943d97', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('8eac61a2-92b3-4fa1-bda7-3a2c62b41f20','voltage','Voltage','d2beb71a-5aeb-4ba3-ae0e-346e81943d97', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8f3ec568-2ba2-46df-a129-a0188e018f24','stage','Stage','704bded7-49c2-43d6-9c4f-9b6b1512b9f1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('930b864a-fe8e-4eb9-a329-3eaa87244a47','voltage','Voltage','704bded7-49c2-43d6-9c4f-9b6b1512b9f1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d7ffabcb-a116-4b08-bb53-2d831897fc55','stage','Stage','707c1f3f-cf12-4768-b826-419478701450', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f3e37ba2-db25-419a-8d09-358f14ae7596','water-temperature','Water-Temperature','707c1f3f-cf12-4768-b826-419478701450', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('987ba22a-4690-4220-bf6b-765f1dfd3ba3','unknown-vb','Unknown VB','707c1f3f-cf12-4768-b826-419478701450', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('8f8cd7e1-2b39-4d18-8445-74eb7392b7f9','stage','Stage','60a37953-139e-4d66-b000-06f5b16ece51', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f9f4c793-75bf-4472-8df7-9c1f7976f0aa','voltage','Voltage','60a37953-139e-4d66-b000-06f5b16ece51', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dc9ffd33-94aa-41ea-af42-2ddb0ffced7c','stage','Stage','908dcb66-f39e-415a-af5f-fa8272a67a42', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('de556448-20b6-44c9-8a0b-2bdddb92cc0f','water-temperature','Water-Temperature','908dcb66-f39e-415a-af5f-fa8272a67a42', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('839be486-1162-4975-aca8-9338881dea11','voltage','Voltage','908dcb66-f39e-415a-af5f-fa8272a67a42', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('fc31ee75-a769-436b-b376-13cc455a4793','stage','Stage','5478ea0d-9117-48fb-9694-02515d516579', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('aa87f62f-5e94-4703-87fc-22b78ba2672c','voltage','Voltage','5478ea0d-9117-48fb-9694-02515d516579', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('00276899-8c9d-48c6-b2f6-dfa9e84a77ad','stage','Stage','59062009-9c5c-413e-9f0a-077fe4d1a015', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fc36b738-9b3b-4605-b3f8-d88a77d8957d','water-temperature','Water-Temperature','59062009-9c5c-413e-9f0a-077fe4d1a015', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('82194f48-9af6-4167-847e-683c30886691','voltage','Voltage','59062009-9c5c-413e-9f0a-077fe4d1a015', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a12a33b6-76fa-4e23-af66-86d42d4ccee8','stage','Stage','a4b91146-0108-4d0d-8cba-21e9cc6281eb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('111aafbd-6c12-45ca-b552-747ce8d9c5b5','voltage','Voltage','a4b91146-0108-4d0d-8cba-21e9cc6281eb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('226616b3-ca58-4712-be2b-68a4bba0257a','stage','Stage','4f556761-b926-4086-bdcd-19bcc311ff1e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('26403f12-e246-4602-b542-742288492beb','voltage','Voltage','4f556761-b926-4086-bdcd-19bcc311ff1e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('bcea49dd-17de-406b-be77-fb1d90d5b29b','stage','Stage','4369132f-f2b1-4161-a857-95e6bcc2de59', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e1796f8f-f3f1-4ff3-bf52-6a06071fefa4','water-temperature','Water-Temperature','4369132f-f2b1-4161-a857-95e6bcc2de59', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('1938028c-6547-4258-a873-fbf02f4d79e3','voltage','Voltage','4369132f-f2b1-4161-a857-95e6bcc2de59', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dd32627d-aa72-4019-838e-6bcc05f3c85c','stage','Stage','76862e2d-87cb-4605-804b-e4d9cb4ce00c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('86212ddd-8725-4a92-854e-39db79b8a3dd','voltage','Voltage','76862e2d-87cb-4605-804b-e4d9cb4ce00c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0d5c2c49-2c0d-4975-9755-6a4f5045dae1','stage','Stage','d4198ec8-52ae-4296-a5ce-3763c3165d79', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b11ee4fa-ca84-447e-9f9e-01f217d781a2','voltage','Voltage','d4198ec8-52ae-4296-a5ce-3763c3165d79', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a36e3fb8-7e99-4918-afc8-8d7313035e7b','stage','Stage','68eaa243-1fe5-4615-9d87-eb0779971a30', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ca85284c-05d2-4e1d-adb9-80d58a5ad9f6','air-temperature','Air-Temperature','68eaa243-1fe5-4615-9d87-eb0779971a30', 'b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'daeee256-c762-43a2-8369-2d295525023c'),
('721ab01c-8c52-48cf-a6c7-058a19e3c172','unknown-us','Unknown US','68eaa243-1fe5-4615-9d87-eb0779971a30', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b6b4fe47-bc4b-46a9-a35b-32037266c033','unknown-ud','Unknown UD','68eaa243-1fe5-4615-9d87-eb0779971a30', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('862f76dc-946a-44e6-93e2-a9842cd478a7','precipitation','Precipitation','68eaa243-1fe5-4615-9d87-eb0779971a30', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c0d3f501-8d54-444b-862c-a41d195e8697','stage','Stage','b3087995-8550-4b0f-a201-dbbf1e52698b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e41d9335-31e8-4566-bcd1-b25c506d2929','voltage','Voltage','b3087995-8550-4b0f-a201-dbbf1e52698b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1b0cc86d-860a-4a41-bae1-d25eca9a51bb','stage','Stage','47c17779-6fd6-4620-a176-4112878f687a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d976de69-4f6d-4719-8371-a1b2f8e66efc','voltage','Voltage','47c17779-6fd6-4620-a176-4112878f687a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3ba3b45e-0c28-4b2b-a2c9-e1f676dbd70b','stage','Stage','081d900a-ef53-442b-940f-d7cee74a26d8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d9d4956e-6ff8-4f3d-8a67-691d46c1dc2d','voltage','Voltage','081d900a-ef53-442b-940f-d7cee74a26d8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('b7cceee6-0fbc-4e2b-9394-3daea12248aa','stage','Stage','514648d9-4d3b-4bc4-a57e-edd091b9f3ce', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bbd8d256-8cf0-4e84-ad15-6945a822564c','voltage','Voltage','514648d9-4d3b-4bc4-a57e-edd091b9f3ce', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2cb1a0a6-de9f-4686-b26f-218d95f64b44','stage','Stage','15101a32-a069-4b0b-a29b-f86a409053b3', 'b49f214e-f69f-43da-9ce3-ad96042268d0', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('f06af255-57a9-414d-b079-fba32ccad052','stage','Stage','ce4e2f40-708a-434e-9d71-2b441ce8611b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0ae2194c-2185-459b-a9b3-867ccd91229b','voltage','Voltage','ce4e2f40-708a-434e-9d71-2b441ce8611b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e1943c31-6359-48d2-b291-60bfdbcfbbdb','stage','Stage','2725f343-d52f-43b2-b41d-0f9f671aa84e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('19f3bdca-3701-435b-98c9-cc1fe3efa552','voltage','Voltage','2725f343-d52f-43b2-b41d-0f9f671aa84e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('af56547b-3965-4227-a5a7-9d4f457baf53','stage','Stage','c0b1ecd2-af0a-450c-8a51-f7b44130d612', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('79920faf-2b68-4cdf-95c7-8c85dd098677','voltage','Voltage','c0b1ecd2-af0a-450c-8a51-f7b44130d612', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('44d1fdf7-d8af-4955-ad59-17982ab2978e','stage','Stage','94e20e3a-7c6b-4eae-ac1c-fea8a22e4662', 'b49f214e-f69f-43da-9ce3-ad96042268d0', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('04cf257a-097e-4c2d-b5c8-05c524441727','stage','Stage','615d5750-5d23-4dfc-98d1-5fa69b3fc8ca', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('faba6e27-4744-4e12-a827-0573bd3e538c','voltage','Voltage','615d5750-5d23-4dfc-98d1-5fa69b3fc8ca', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('423fd5c0-091f-47fe-b28b-9c34983c5f74','stage','Stage','f8f8fd49-cdf8-47ec-bfaa-a406ccbb12c2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('49d59411-f70c-42b3-8d5f-77a6f5c56452','voltage','Voltage','f8f8fd49-cdf8-47ec-bfaa-a406ccbb12c2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6f0d5876-5c0e-4eec-baf9-84969e938037','stage','Stage','e0806b46-1b66-460e-97d5-0c98bff2edd9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b5b5abb4-b714-4691-b3ae-1aef0825cdda','voltage','Voltage','e0806b46-1b66-460e-97d5-0c98bff2edd9', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('efb64af7-061a-40db-9c3b-dc092483f350','stage','Stage','b7425815-d58a-4422-9907-e115f41c77dd', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('86215576-17fd-4751-8ce9-a460382e528b','voltage','Voltage','b7425815-d58a-4422-9907-e115f41c77dd', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('231fd808-acb7-4463-abef-dce6e4271b1b','stage','Stage','6207bc92-5cc7-46ed-a93f-1ff32121a68a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4a2aca4a-83e6-4327-8d38-5bf36e706264','stage','Stage','5f0e95fe-84b1-4dee-9584-408d2b79b4e1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ed7b5cc6-1224-482f-a1ec-1aee48d94094','water-temperature','Water-Temperature','5f0e95fe-84b1-4dee-9584-408d2b79b4e1', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('2fde7754-279a-49fc-a85a-500281dc4aaf','voltage','Voltage','5f0e95fe-84b1-4dee-9584-408d2b79b4e1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('9d4765f1-32b3-4313-a5d4-039c8187ef33','stage','Stage','34d6ecd5-a584-4de1-90a6-dacc386988bd', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0c3bf4b8-3e2c-4c2b-bb73-36825e2a5959','water-temperature','Water-Temperature','34d6ecd5-a584-4de1-90a6-dacc386988bd', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('17b11ffa-4792-476c-bc56-9edbcbdec3a3','voltage','Voltage','34d6ecd5-a584-4de1-90a6-dacc386988bd', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2995e2d9-acdb-42ee-a5bf-751b87ded8ff','stage','Stage','5f4a3ce2-f359-4f7b-b622-9db903ca8f30', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5345687d-5dd7-42c6-99bc-6a4dab4be5e0','voltage','Voltage','5f4a3ce2-f359-4f7b-b622-9db903ca8f30', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('35c1d7b5-e817-40ac-b925-974e2294ef6f','stage','Stage','d0c6b11b-44a4-4041-9cd7-1ff725f042a4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e5dcefb3-a1e6-4b72-9c68-d790f886cc16','precipitation','Precipitation','d0c6b11b-44a4-4041-9cd7-1ff725f042a4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('89e01ca7-bdbb-4858-9d95-d3a7af093c88','voltage','Voltage','d0c6b11b-44a4-4041-9cd7-1ff725f042a4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a09d3a47-d9b2-4c46-ad18-1ddeaae38486','stage','Stage','9c4eec28-485c-405f-93fe-a1671ba1e85e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('68bf4702-e9a9-4a7b-85f7-59d18b952941','voltage','Voltage','9c4eec28-485c-405f-93fe-a1671ba1e85e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1ff4db26-5120-4461-9e51-b8d52ca97a47','precipitation','Precipitation','2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('05afcc40-7265-4517-b621-ad93f21c4319','stage','Stage','2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('913aa8cf-4e83-48e7-8061-f16c5feecb4d','water-temperature','Water-Temperature','2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c'),
('0d2d566c-4978-4bca-9250-a17f6ec0f717','unknown-ws','Unknown WS','2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('fd5747f5-be7f-4cde-824f-dda541e5bd07','voltage','Voltage','2a0651b4-b2fa-4e91-a4f6-1412b6ce7893', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8c6123d2-3c1a-43a0-8878-235c4c290668','stage','Stage','c8197081-957a-451d-ab4d-41805e0ecf5f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('967bd426-9384-4280-92fc-efc6750e179f','water-temperature','Water-Temperature','c8197081-957a-451d-ab4d-41805e0ecf5f', 'de6112da-8489-4286-ae56-ec72aa09974d', 'daeee256-c762-43a2-8369-2d295525023c');
