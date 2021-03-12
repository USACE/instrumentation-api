
INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('1fa28968-d0b1-4832-b4b3-e58206056819', 'd8f8934d-e414-499d-bd51-bc93bbde6345', 
    'chicago-district-streamgages', 'Chicago District Streamgages', 'chicago-district-streamgages.jpg');

--#################################################
--DELETE existing data (mostly for dev, test, prod)
--#################################################

-- Delete Timeseries Measurements
delete from timeseries_measurement where timeseries_id in (
	select t.id from instrument i 
	join timeseries t on i.id = t.instrument_id
	where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819');

-- Delete Timeseries
delete from timeseries where instrument_id in (
	select i.id from instrument i
	where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819');
	
-- Delete Telemetry GOES
delete from telemetry_goes where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='10a32652-af43-4451-bd52-4980c5690cc9'
	and instrument_id in (
		select i.id from instrument i
		where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819')
	);

-- Delete Telemetry Iridium
delete from telemetry_iridium where id in (
	select telemetry_id from instrument_telemetry
	where telemetry_type_id='c0b03b0d-bfce-453a-b5a9-636118940449'
	and	instrument_id in (
		select i.id from instrument i
		where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819')
	);
	
-- Delete Instrument Telemetry
delete from instrument_telemetry where instrument_id in (
	select i.id from instrument i
	where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819'
    );
	
-- Delete Instrument Status
delete from instrument_status where instrument_id in (
	select i.id from instrument i
	where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819'
    );

-- Delete Instrument Group Instruments
delete from instrument_group_instruments where instrument_id in (
	select id from instrument i
	where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819'
	);

-- Delete Instrument Groups
delete from instrument_group 
where project_id = '1fa28968-d0b1-4832-b4b3-e58206056819';
	
--Delete Collection Groups

--Delete collection_group_timeseries

-- Delete Instruments
delete from instrument i 
where i.project_id = '1fa28968-d0b1-4832-b4b3-e58206056819';

-- Delete Alert Config

-- Delete Alert

--########################################
-- INSERT new data (built by script)
--########################################


--INSERT INSTRUMENTS--COUNT:40
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('6fbeb371-98db-4b70-93bb-831b14f72c7b', False, 'portland_f32d', 'Portland', null, ST_GeomFromText('POINT(-85.039 40.4277)',4326), null, null, '2021-03-12T16:51:13.861956Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324200'),
('1a47c87e-50d1-48ee-bb78-c246580263cb', False, 'linn-grove', 'Linn Grove', null, ST_GeomFromText('POINT(-85.0309 40.6439)',4326), null, null, '2021-03-12T16:51:13.862190Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03322900'),
('4deb6eee-81ed-4edd-ac38-0b081eb1a081', False, 'berlin', 'Berlin', null, ST_GeomFromText('POINT(-88.95 43.9539)',4326), null, null, '2021-03-12T16:51:13.862339Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04073500'),
('b3f286b8-60a3-43e0-9d8a-950f342388a4', False, 'obrien', 'Obrien', null, ST_GeomFromText('POINT(-87.5611 41.65)',4326), null, null, '2021-03-12T16:51:13.862479Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', False, 'fonddulac', 'FondDuLac', null, ST_GeomFromText('POINT(-88.4561 43.8)',4326), null, null, '2021-03-12T16:51:13.862806Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('005e3595-446b-40e1-b07c-604d2a808acc', False, 'menasha', 'Menasha', null, ST_GeomFromText('POINT(-88.4472 44.1994)',4326), null, null, '2021-03-12T16:51:13.863032Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('2155fea5-633e-4556-8c57-c49c8cf34cc1', False, 'lockport', 'Lockport', null, ST_GeomFromText('POINT(-88.0789 41.5697)',4326), null, null, '2021-03-12T16:51:13.863208Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('a597b201-1aa7-4965-8ff2-a1a26fef4763', False, 'mississinewa-tailwater', 'Mississinewa-Tailwater', null, ST_GeomFromText('POINT(-85.9575 40.7233)',4326), null, null, '2021-03-12T16:51:13.863521Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327000'),
('5c797d3a-f2f6-4149-86b3-e87b3b20618e', False, 'marion', 'Marion', null, ST_GeomFromText('POINT(-85.6595 40.5764)',4326), null, null, '2021-03-12T16:51:13.863666Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326500'),
('992677cd-af48-4ee9-82df-e05f17c8724d', False, 'mississinewa-pool', 'Mississinewa-Pool', null, ST_GeomFromText('POINT(-85.9572 40.7144)',4326), null, null, '2021-03-12T16:51:13.863742Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326950'),
('c844e536-5a97-4fa9-9e2c-d6502219c115', False, 'bluffton', 'Bluffton', null, ST_GeomFromText('POINT(-85.1714 40.7424)',4326), null, null, '2021-03-12T16:51:13.863815Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04099510'),
('653e2c8f-eac0-4d0a-882c-abe28355962e', False, 'jerome', 'Jerome', null, ST_GeomFromText('POINT(-85.9188 40.4413)',4326), null, null, '2021-03-12T16:51:13.863916Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333450'),
('e9e90579-80a4-4d53-ac11-689a920b887a', False, 'kokomo', 'Kokomo', null, ST_GeomFromText('POINT(-86.1529 40.4709)',4326), null, null, '2021-03-12T16:51:13.863986Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333700'),
('b5564b22-e521-49b7-b398-2bf9bc1c5398', False, 'owasco', 'Owasco', null, ST_GeomFromText('POINT(-86.6366 40.4648)',4326), null, null, '2021-03-12T16:51:13.864056Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03334000'),
('c090102e-82b3-4a01-9bc0-ce73395d3226', False, 'deer-creek', 'Deer Creek', null, ST_GeomFromText('POINT(-86.6214 40.5903)',4326), null, null, '2021-03-12T16:51:13.864130Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329700'),
('b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', False, 'salamonie-tailwater', 'Salamonie-Tailwater', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-12T16:51:13.864203Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324500'),
('622ba7b6-0bd2-4c26-9275-17b6462c525e', False, 'warren_efc6', 'Warren', null, ST_GeomFromText('POINT(-85.4536 40.7125)',4326), null, null, '2021-03-12T16:51:13.864398Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324300'),
('b48f1167-ef1a-469e-89d4-143263120130', False, 'j-edward-roush-pool', 'J Edward Roush-Pool', null, ST_GeomFromText('POINT(-85.4686 40.8461)',4326), null, null, '2021-03-12T16:51:13.864581Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323450'),
('b89d5ce3-7d71-4958-8809-2f7e701a5d99', False, 'j-edward-roush-tailwater', 'J Edward Roush-Tailwater', null, ST_GeomFromText('POINT(-85.4898 40.8533)',4326), null, null, '2021-03-12T16:51:13.864701Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323500'),
('28b4878c-3e78-4898-84d2-c0d9a47d6a54', False, 'wabash', 'Wabash', null, ST_GeomFromText('POINT(-85.8203 40.7908)',4326), null, null, '2021-03-12T16:51:13.864871Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03325000'),
('913eb762-d4be-491f-9fde-cbd054b1dfc0', False, 'lafayettewildcat', 'LafayetteWildcat', null, ST_GeomFromText('POINT(-86.8292 40.4406)',4326), null, null, '2021-03-12T16:51:13.865023Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335000'),
('c8877ff5-6a74-479a-8b8b-f2bc0a103a47', False, 'delphi', 'Delphi', null, ST_GeomFromText('POINT(-86.7703 40.5939)',4326), null, null, '2021-03-12T16:51:13.865104Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333050'),
('11258ce5-b0f0-475f-b06e-f3f056d7e7b8', False, 'oswego', 'Oswego', null, ST_GeomFromText('POINT(-85.7892 41.3206)',4326), null, null, '2021-03-12T16:51:13.865208Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330500'),
('0464c8ba-0e8c-4621-9d43-c91e90b011e9', False, 'ora', 'Ora', null, ST_GeomFromText('POINT(-86.5636 41.1572)',4326), null, null, '2021-03-12T16:51:13.865281Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03331500'),
('ebcdfa47-b1b6-4f49-87f4-85ac03221578', False, 'north-manchester', 'North Manchester', null, ST_GeomFromText('POINT(-85.7825 40.9944)',4326), null, null, '2021-03-12T16:51:13.865385Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03328000'),
('8c23bf4f-9ac0-4022-bdc2-a730229b87a9', False, 'north-webster', 'North Webster', null, ST_GeomFromText('POINT(-85.6922 41.3164)',4326), null, null, '2021-03-12T16:51:13.865458Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330241'),
('7e8a59bd-7e19-45f7-8e75-b336a3bccad2', False, 'newlondon', 'NewLondon', null, ST_GeomFromText('POINT(-88.7403 44.3922)',4326), null, null, '2021-03-12T16:51:13.865531Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04079000'),
('28927945-3f8f-4dbe-8055-c92baa9df5a7', False, 'oshkosh', 'Oshkosh', null, ST_GeomFromText('POINT(-88.5272 44.0097)',4326), null, null, '2021-03-12T16:51:13.865659Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04082500'),
('6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', False, 'royalton', 'Royalton', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-12T16:51:13.865791Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04080000'),
('a3de6a31-67af-43e7-9ac2-d6042e6823c9', False, 'stockbridge', 'Stockbridge', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-12T16:51:13.865964Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04084255'),
('3e8f793f-9e7b-482d-b08d-b2354e727644', False, 'waupaca', 'Waupaca', null, ST_GeomFromText('POINT(-88.9961 44.3292)',4326), null, null, '2021-03-12T16:51:13.866105Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04081000'),
('186ea37e-004c-40fd-8527-b5a288dc5764', False, 'fritsepark', 'FritsePark', null, ST_GeomFromText('POINT(-88.4703 44.205)',4326), null, null, '2021-03-12T16:51:13.866236Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('1eab7035-942e-4e0c-87c5-b27e2e353a35', False, 'poygan', 'Poygan', null, ST_GeomFromText('POINT(-88.7125 44.1108)',4326), null, null, '2021-03-12T16:51:13.866363Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', False, 'salamonie-pool', 'Salamonie-Pool', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-12T16:51:13.866552Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324450'),
('4a7d6948-c69d-4401-8413-8576229a2fa4', False, 'peru', 'Peru', null, ST_GeomFromText('POINT(-86.0667 40.75)',4326), null, null, '2021-03-12T16:51:13.866625Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327500'),
('4347b406-e66e-4bd1-92d8-b322ab2e9254', False, 'logansport', 'Logansport', null, ST_GeomFromText('POINT(-86.3775 40.7464)',4326), null, null, '2021-03-12T16:51:13.866755Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329000'),
('f79a66b0-6584-4d2b-8340-e014b0fa66e0', False, 'lafayette', 'Lafayette', null, ST_GeomFromText('POINT(-86.8969 40.4219)',4326), null, null, '2021-03-12T16:51:13.866865Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335500'),
('cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', False, 'littleriver', 'LittleRiver', null, ST_GeomFromText('POINT(-85.4132 40.8986)',4326), null, null, '2021-03-12T16:51:13.866954Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324000'),
('ce302849-2d7d-499c-a862-6b077c42a0ee', False, 'covington', 'Covington', null, ST_GeomFromText('POINT(-87.4067 40.14)',4326), null, null, '2021-03-12T16:51:13.867056Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03336000'),
('45948095-3b78-4123-91f8-b86e95f28c42', False, 'montezuma', 'Montezuma', null, ST_GeomFromText('POINT(-87.3739 39.7925)',4326), null, null, '2021-03-12T16:51:13.867162Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03340500');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('3cc5a75e-da66-4dbb-acf2-ccaa1b5afa6c', '6fbeb371-98db-4b70-93bb-831b14f72c7b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.861956Z'),
('f953a85f-2d62-4adf-a0db-272986d2bbbd', '1a47c87e-50d1-48ee-bb78-c246580263cb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.862190Z'),
('19877e9e-9beb-4e5d-bd19-1f9d9ee5cc99', '4deb6eee-81ed-4edd-ac38-0b081eb1a081', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.862339Z'),
('101d959e-31eb-4f92-8c38-62b9a51d6335', 'b3f286b8-60a3-43e0-9d8a-950f342388a4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.862479Z'),
('a92bc85e-2dcf-47d7-90a2-04476f0bed8c', '3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.862806Z'),
('b9bcea2e-789a-4d6a-99cb-2fd9ca2f1b7a', '005e3595-446b-40e1-b07c-604d2a808acc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863032Z'),
('b014373c-cc1a-4db4-b747-3bf73cabb892', '2155fea5-633e-4556-8c57-c49c8cf34cc1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863208Z'),
('678f7e32-760c-4bf3-ae3c-9a6cb4d5cb49', 'a597b201-1aa7-4965-8ff2-a1a26fef4763', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863521Z'),
('fabff80e-99a9-4f06-b271-5998bc8e0dd3', '5c797d3a-f2f6-4149-86b3-e87b3b20618e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863666Z'),
('e17e556e-b68a-4127-8927-e6a0b88d9425', '992677cd-af48-4ee9-82df-e05f17c8724d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863742Z'),
('6dd42635-11c6-4f32-a2b4-ccd3f0c45ae7', 'c844e536-5a97-4fa9-9e2c-d6502219c115', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863815Z'),
('7cf4368e-d0fa-42fb-b995-dddff1dd4c07', '653e2c8f-eac0-4d0a-882c-abe28355962e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863916Z'),
('976ac389-93d0-499e-ac0d-f4c55020e8ec', 'e9e90579-80a4-4d53-ac11-689a920b887a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.863986Z'),
('1145d3ef-17fc-466e-ab72-435dad2473ee', 'b5564b22-e521-49b7-b398-2bf9bc1c5398', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864056Z'),
('70901903-98c3-4db1-bfa9-81816325dc09', 'c090102e-82b3-4a01-9bc0-ce73395d3226', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864130Z'),
('a89f2a46-313f-48bd-a057-83976a879e59', 'b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864203Z'),
('21d69093-1e97-4fb6-92b6-a5d3898b88fb', '622ba7b6-0bd2-4c26-9275-17b6462c525e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864398Z'),
('3cdac32a-2ff3-4d4b-a5e1-4e07c5b88022', 'b48f1167-ef1a-469e-89d4-143263120130', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864581Z'),
('d33a9345-36e3-41d1-910e-33f50b48e088', 'b89d5ce3-7d71-4958-8809-2f7e701a5d99', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864701Z'),
('723fc953-1719-4d7b-b536-11c4ea4a4b0a', '28b4878c-3e78-4898-84d2-c0d9a47d6a54', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.864871Z'),
('1e524843-4907-4a22-81f6-4aaa869fb88e', '913eb762-d4be-491f-9fde-cbd054b1dfc0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865023Z'),
('4d18d156-eedb-4ec1-9c45-d06f26a5dccd', 'c8877ff5-6a74-479a-8b8b-f2bc0a103a47', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865104Z'),
('2fbbe16a-c425-4254-a164-53ed15df5ec3', '11258ce5-b0f0-475f-b06e-f3f056d7e7b8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865208Z'),
('ec29634e-7d5b-450d-9074-c6f3869d6d68', '0464c8ba-0e8c-4621-9d43-c91e90b011e9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865281Z'),
('d071266e-0cf6-43ab-a9ca-b4fb9258ae75', 'ebcdfa47-b1b6-4f49-87f4-85ac03221578', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865385Z'),
('e69ee703-9eb1-44f1-aaab-bcaf9115d3aa', '8c23bf4f-9ac0-4022-bdc2-a730229b87a9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865458Z'),
('13cdf92b-baeb-483a-ae17-271076aeac62', '7e8a59bd-7e19-45f7-8e75-b336a3bccad2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865531Z'),
('7860cc42-7cdd-40d6-a8c1-7be5789a6d40', '28927945-3f8f-4dbe-8055-c92baa9df5a7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865659Z'),
('f47b8f2a-ddba-4036-bb55-caa744aea59d', '6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865791Z'),
('a4de0f71-9495-4b7f-940a-cc6943240a3a', 'a3de6a31-67af-43e7-9ac2-d6042e6823c9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.865964Z'),
('1c4d1186-1543-49e4-96ce-2a40dcf09237', '3e8f793f-9e7b-482d-b08d-b2354e727644', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866105Z'),
('7de84a49-09b4-4165-afea-bd4674f0cf8c', '186ea37e-004c-40fd-8527-b5a288dc5764', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866236Z'),
('dbd006a1-9397-456b-90c2-32601a0f4db5', '1eab7035-942e-4e0c-87c5-b27e2e353a35', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866363Z'),
('ebf3eeff-271a-42cb-9399-7a33f8bd32a2', 'a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866552Z'),
('7ebe45fd-53b7-4ecc-b390-3fe4b5cba94b', '4a7d6948-c69d-4401-8413-8576229a2fa4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866625Z'),
('57206db5-9721-4332-9e86-2482e9020602', '4347b406-e66e-4bd1-92d8-b322ab2e9254', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866755Z'),
('05c40587-704d-4ded-8be9-4d6c0bc32074', 'f79a66b0-6584-4d2b-8340-e014b0fa66e0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866865Z'),
('361f02d3-23ca-4d54-b392-d77328f9b8ff', 'cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.866954Z'),
('9587676e-5a7d-4503-882b-cb55c66ef393', 'ce302849-2d7d-499c-a862-6b077c42a0ee', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.867056Z'),
('288c7498-a4fb-4a80-bfe9-508cc89cc7c3', '45948095-3b78-4123-91f8-b86e95f28c42', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-12T16:51:13.867162Z');

--INSERT TELEMETRY_GOES--COUNT:40
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5f802ce1-f868-441e-84dd-392efd799d20', 'DD46A45A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD46A45A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c08c654f-fc1b-4abf-aa5a-59db7d6843ef', 'CE6B687C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B687C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '51dd91e3-38b0-41bb-b5b9-5f2f1efaf6e3', 'CE72E178' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72E178');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2afecdd3-0594-4211-ae8e-35af5b415ef4', 'CE2A970C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE2A970C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6bcc382d-7b73-4752-aa07-d600c02a9c44', 'CE72DA30' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72DA30');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5040cee1-c8f9-4da2-8844-db69eaa29385', 'CE72D4E2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72D4E2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8d7246a4-c839-4c6b-bc7f-0eb939faf04f', 'CE720C58' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE720C58');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7f207386-cb80-4b12-b986-80adc3aa075b', 'CE7723A6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7723A6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd0e8e430-ae58-402f-90be-20b0e1c595d7', 'CE6B5DE6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B5DE6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'bd2f9788-68b9-4733-b662-52fbaa33ccbd', 'CE108358' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE108358');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5af7d3f9-1b10-4eac-8a21-bee8ad23456b', '17D2F560' where not exists (select 1 from telemetry_goes where nesdis_id = '17D2F560');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd80e2645-cd69-434c-8589-8d0540db9a5c', 'D11EE2FE' where not exists (select 1 from telemetry_goes where nesdis_id = 'D11EE2FE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c8d62f29-e8dc-4975-86be-e39fac580bf9', '173623B8' where not exists (select 1 from telemetry_goes where nesdis_id = '173623B8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e808f952-cd71-4e63-a93c-e5513faa7013', '167526C8' where not exists (select 1 from telemetry_goes where nesdis_id = '167526C8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b22dfd31-1e58-44fd-8d7a-a58c8a7cb344', 'DE4B8336' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE4B8336');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '60046334-76ce-4db8-bfc6-6996e924447c', 'CE7718EE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7718EE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd6763f1d-27d0-4043-89f4-3633f3a1c18f', 'CE6B66AE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B66AE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'af233cde-6b31-4783-bed0-82bbc6e65043', 'CE10C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE10C052');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f3ec11b6-a6f3-4ab2-940f-ba3cb8189a81', 'CE77163C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE77163C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7e83dd3e-ad16-4edb-ac63-5f0434377acf', 'CE16B60C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16B60C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '647c17c5-e62e-49b7-8cd8-e013a904402c', 'CE6B0D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B0D9A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'bf04eb31-bbe8-4878-ba21-c49a29ca281b', 'DDA6F1AC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA6F1AC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7677a225-7671-4349-87f4-ffb2ddeddce1', '17853608' where not exists (select 1 from telemetry_goes where nesdis_id = '17853608');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ef823893-e9e8-4563-be69-b73e4ad2173e', 'DD8F700A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8F700A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1d7d69c1-8734-48d2-89f1-6f4f38574c71', '163D923C' where not exists (select 1 from telemetry_goes where nesdis_id = '163D923C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8b1c2150-3943-4029-9b94-bcca96e5aa99', 'DD3E3032' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD3E3032');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a3d87a07-8719-4512-9cec-d47edf2e0cae', 'CE72FCDC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72FCDC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7545dd61-1b77-43dc-8f8d-28a3585817be', 'CE58F2B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE58F2B2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd514837b-4d59-49a4-b180-a3074ae6543c', 'CE72EFAA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72EFAA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e0e82613-fc55-479f-9be8-32a13f8acfb7', 'CE72F20E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72F20E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8e6ef8dc-3364-4ba1-ada5-d0d789c8d364', 'CE730070' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE730070');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f40dc2b0-357e-41fd-bbac-57715f64122e', 'CE26C6EC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE26C6EC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e9a3a492-d632-4e3d-b85e-fa34ebf8009b', 'CE72C946' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72C946');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '4f8d4fe2-44e1-4d15-aa36-fa90b4959e86', 'CE6D2BB8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D2BB8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '10652b92-253f-451c-8e6b-2a7ef7629558', 'CE777D08' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE777D08');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '3379946e-12c4-4522-815d-c3b6c39b507b', 'CE6D256A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D256A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b0606aff-98e3-423b-a3a2-f6ff9fcae996', 'CE6D361C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D361C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0872bc11-bcf6-4c7a-b5a1-167e6c7be74a', 'CE14B3F8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14B3F8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '489f3ba0-5732-49ca-8f82-a31ea32b4b41', 'CE16C09C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16C09C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '49e9e102-b6ee-4050-b5ce-09e4597b5093', 'CE14E384' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14E384');

--INSERT INSTRUMENT_TELEMETRY--COUNT:40
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('6fbeb371-98db-4b70-93bb-831b14f72c7b', '10a32652-af43-4451-bd52-4980c5690cc9', '5f802ce1-f868-441e-84dd-392efd799d20'),
('1a47c87e-50d1-48ee-bb78-c246580263cb', '10a32652-af43-4451-bd52-4980c5690cc9', 'c08c654f-fc1b-4abf-aa5a-59db7d6843ef'),
('4deb6eee-81ed-4edd-ac38-0b081eb1a081', '10a32652-af43-4451-bd52-4980c5690cc9', '51dd91e3-38b0-41bb-b5b9-5f2f1efaf6e3'),
('b3f286b8-60a3-43e0-9d8a-950f342388a4', '10a32652-af43-4451-bd52-4980c5690cc9', '2afecdd3-0594-4211-ae8e-35af5b415ef4'),
('3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '10a32652-af43-4451-bd52-4980c5690cc9', '6bcc382d-7b73-4752-aa07-d600c02a9c44'),
('005e3595-446b-40e1-b07c-604d2a808acc', '10a32652-af43-4451-bd52-4980c5690cc9', '5040cee1-c8f9-4da2-8844-db69eaa29385'),
('2155fea5-633e-4556-8c57-c49c8cf34cc1', '10a32652-af43-4451-bd52-4980c5690cc9', '8d7246a4-c839-4c6b-bc7f-0eb939faf04f'),
('a597b201-1aa7-4965-8ff2-a1a26fef4763', '10a32652-af43-4451-bd52-4980c5690cc9', '7f207386-cb80-4b12-b986-80adc3aa075b'),
('5c797d3a-f2f6-4149-86b3-e87b3b20618e', '10a32652-af43-4451-bd52-4980c5690cc9', 'd0e8e430-ae58-402f-90be-20b0e1c595d7'),
('992677cd-af48-4ee9-82df-e05f17c8724d', '10a32652-af43-4451-bd52-4980c5690cc9', 'bd2f9788-68b9-4733-b662-52fbaa33ccbd'),
('c844e536-5a97-4fa9-9e2c-d6502219c115', '10a32652-af43-4451-bd52-4980c5690cc9', '5af7d3f9-1b10-4eac-8a21-bee8ad23456b'),
('653e2c8f-eac0-4d0a-882c-abe28355962e', '10a32652-af43-4451-bd52-4980c5690cc9', 'd80e2645-cd69-434c-8589-8d0540db9a5c'),
('e9e90579-80a4-4d53-ac11-689a920b887a', '10a32652-af43-4451-bd52-4980c5690cc9', 'c8d62f29-e8dc-4975-86be-e39fac580bf9'),
('b5564b22-e521-49b7-b398-2bf9bc1c5398', '10a32652-af43-4451-bd52-4980c5690cc9', 'e808f952-cd71-4e63-a93c-e5513faa7013'),
('c090102e-82b3-4a01-9bc0-ce73395d3226', '10a32652-af43-4451-bd52-4980c5690cc9', 'b22dfd31-1e58-44fd-8d7a-a58c8a7cb344'),
('b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', '10a32652-af43-4451-bd52-4980c5690cc9', '60046334-76ce-4db8-bfc6-6996e924447c'),
('622ba7b6-0bd2-4c26-9275-17b6462c525e', '10a32652-af43-4451-bd52-4980c5690cc9', 'd6763f1d-27d0-4043-89f4-3633f3a1c18f'),
('b48f1167-ef1a-469e-89d4-143263120130', '10a32652-af43-4451-bd52-4980c5690cc9', 'af233cde-6b31-4783-bed0-82bbc6e65043'),
('b89d5ce3-7d71-4958-8809-2f7e701a5d99', '10a32652-af43-4451-bd52-4980c5690cc9', 'f3ec11b6-a6f3-4ab2-940f-ba3cb8189a81'),
('28b4878c-3e78-4898-84d2-c0d9a47d6a54', '10a32652-af43-4451-bd52-4980c5690cc9', '7e83dd3e-ad16-4edb-ac63-5f0434377acf'),
('913eb762-d4be-491f-9fde-cbd054b1dfc0', '10a32652-af43-4451-bd52-4980c5690cc9', '647c17c5-e62e-49b7-8cd8-e013a904402c'),
('c8877ff5-6a74-479a-8b8b-f2bc0a103a47', '10a32652-af43-4451-bd52-4980c5690cc9', 'bf04eb31-bbe8-4878-ba21-c49a29ca281b'),
('11258ce5-b0f0-475f-b06e-f3f056d7e7b8', '10a32652-af43-4451-bd52-4980c5690cc9', '7677a225-7671-4349-87f4-ffb2ddeddce1'),
('0464c8ba-0e8c-4621-9d43-c91e90b011e9', '10a32652-af43-4451-bd52-4980c5690cc9', 'ef823893-e9e8-4563-be69-b73e4ad2173e'),
('ebcdfa47-b1b6-4f49-87f4-85ac03221578', '10a32652-af43-4451-bd52-4980c5690cc9', '1d7d69c1-8734-48d2-89f1-6f4f38574c71'),
('8c23bf4f-9ac0-4022-bdc2-a730229b87a9', '10a32652-af43-4451-bd52-4980c5690cc9', '8b1c2150-3943-4029-9b94-bcca96e5aa99'),
('7e8a59bd-7e19-45f7-8e75-b336a3bccad2', '10a32652-af43-4451-bd52-4980c5690cc9', 'a3d87a07-8719-4512-9cec-d47edf2e0cae'),
('28927945-3f8f-4dbe-8055-c92baa9df5a7', '10a32652-af43-4451-bd52-4980c5690cc9', '7545dd61-1b77-43dc-8f8d-28a3585817be'),
('6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', '10a32652-af43-4451-bd52-4980c5690cc9', 'd514837b-4d59-49a4-b180-a3074ae6543c'),
('a3de6a31-67af-43e7-9ac2-d6042e6823c9', '10a32652-af43-4451-bd52-4980c5690cc9', 'e0e82613-fc55-479f-9be8-32a13f8acfb7'),
('3e8f793f-9e7b-482d-b08d-b2354e727644', '10a32652-af43-4451-bd52-4980c5690cc9', '8e6ef8dc-3364-4ba1-ada5-d0d789c8d364'),
('186ea37e-004c-40fd-8527-b5a288dc5764', '10a32652-af43-4451-bd52-4980c5690cc9', 'f40dc2b0-357e-41fd-bbac-57715f64122e'),
('1eab7035-942e-4e0c-87c5-b27e2e353a35', '10a32652-af43-4451-bd52-4980c5690cc9', 'e9a3a492-d632-4e3d-b85e-fa34ebf8009b'),
('a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', '10a32652-af43-4451-bd52-4980c5690cc9', '4f8d4fe2-44e1-4d15-aa36-fa90b4959e86'),
('4a7d6948-c69d-4401-8413-8576229a2fa4', '10a32652-af43-4451-bd52-4980c5690cc9', '10652b92-253f-451c-8e6b-2a7ef7629558'),
('4347b406-e66e-4bd1-92d8-b322ab2e9254', '10a32652-af43-4451-bd52-4980c5690cc9', '3379946e-12c4-4522-815d-c3b6c39b507b'),
('f79a66b0-6584-4d2b-8340-e014b0fa66e0', '10a32652-af43-4451-bd52-4980c5690cc9', 'b0606aff-98e3-423b-a3a2-f6ff9fcae996'),
('cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', '10a32652-af43-4451-bd52-4980c5690cc9', '0872bc11-bcf6-4c7a-b5a1-167e6c7be74a'),
('ce302849-2d7d-499c-a862-6b077c42a0ee', '10a32652-af43-4451-bd52-4980c5690cc9', '489f3ba0-5732-49ca-8f82-a31ea32b4b41'),
('45948095-3b78-4123-91f8-b86e95f28c42', '10a32652-af43-4451-bd52-4980c5690cc9', '49e9e102-b6ee-4050-b5ce-09e4597b5093');

--INSERT TIMESERIES--COUNT:40
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('19ee54b3-1eaa-43d2-b209-ed3f250077f9','stage','Stage','6fbeb371-98db-4b70-93bb-831b14f72c7b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f2ca7fc2-a66d-470a-a5d5-374cb719d7ba','stage','Stage','1a47c87e-50d1-48ee-bb78-c246580263cb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d57653ee-b3c7-4a6b-8b98-71781a41e7ce','precipitation','Precipitation','1a47c87e-50d1-48ee-bb78-c246580263cb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d6b0a480-c2b3-4a2d-a8bc-6de0dc3dd51b','stage','Stage','4deb6eee-81ed-4edd-ac38-0b081eb1a081', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('480c0509-7df7-4206-986e-ca2de854c2c1','precipitation','Precipitation','4deb6eee-81ed-4edd-ac38-0b081eb1a081', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('fc8c123e-745f-42b4-87df-fafd72bfcad2','unknown-volts','Unknown Volts','4deb6eee-81ed-4edd-ac38-0b081eb1a081', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f31fad51-b09d-4bf0-8e49-10e45c44e4c3','unknown-elev-pool','Unknown Elev-Pool','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('20cbd4de-5681-4011-b915-3511905bb215','unknown-elev-tailwater','Unknown Elev-Tailwater','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('f16d84ab-cb78-4ac9-941f-52c8f9f4ff33','precipitation','Precipitation','b3f286b8-60a3-43e0-9d8a-950f342388a4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a4c4832d-af61-4a79-9732-14c06cc12b39','unknown-wind-speed','Unknown Wind Speed','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('1d7742a1-dd6f-4b3b-92c1-b23025029bce','unknown-wind-direction','Unknown Wind Direction','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('abc5ad8e-8c4e-42ad-b55f-a1da041338ae','unknown-air-temp','Unknown Air Temp','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('05e075e8-7f12-46c4-bf3e-d7e40e09af2d','unknown-water-temp','Unknown Water Temp','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('5f82ea81-a900-46f4-9074-9c87264e0f2b','unknown-volt','Unknown Volt','b3f286b8-60a3-43e0-9d8a-950f342388a4', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('46462ef5-b7fa-4741-a0bb-f642f59d524b','elevation','Elevation','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e31b80eb-702e-4cc7-9f92-8b3d3a0f57b7','precipitation','Precipitation','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('742ffc7a-f514-4033-88da-e951605b737d','unknown-temp','Unknown Temp','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('660cd9f0-7110-4b16-9207-3e950ae04528','voltage','Voltage','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2b5f5152-8e91-4ea6-9e22-685b8280f53b','elevation','Elevation','005e3595-446b-40e1-b07c-604d2a808acc', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0fad1ae4-b422-444b-8733-9c0bbd97d914','precipitation','Precipitation','005e3595-446b-40e1-b07c-604d2a808acc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('2a52ac56-ef61-4362-a051-99669591b228','unknown-temp','Unknown Temp','005e3595-446b-40e1-b07c-604d2a808acc', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('cc38d3d8-662d-43fd-9746-2f4d33f80774','voltage','Voltage','005e3595-446b-40e1-b07c-604d2a808acc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cc3cd3aa-07cd-4e30-8d86-96513ab21901','elevation','Elevation','2155fea5-633e-4556-8c57-c49c8cf34cc1', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3fb8a852-5056-41b7-92d6-e2ae0d6fbaea','unknown-elev-tailwater','Unknown Elev-Tailwater','2155fea5-633e-4556-8c57-c49c8cf34cc1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('2b778059-22bd-48af-847b-8769603aead1','precipitation','Precipitation','2155fea5-633e-4556-8c57-c49c8cf34cc1', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('90d8db7f-f403-4997-a29d-5d2bf684185d','unknown-wind-speed','Unknown Wind Speed','2155fea5-633e-4556-8c57-c49c8cf34cc1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('8aa09517-db0f-43bc-88ac-12de30a54982','unknown-wind-dir','Unknown Wind Dir','2155fea5-633e-4556-8c57-c49c8cf34cc1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ae03b7a2-60f7-4ebf-9d39-54309c859316','unknown-air-temp','Unknown Air Temp','2155fea5-633e-4556-8c57-c49c8cf34cc1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ad71b4e1-e136-418d-b8e2-f1e481823592','unknown-water-temp','Unknown Water Temp','2155fea5-633e-4556-8c57-c49c8cf34cc1', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('ed085880-fab7-40bc-948c-16d6d5793e00','voltage','Voltage','2155fea5-633e-4556-8c57-c49c8cf34cc1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cccb71e3-ff41-4207-b73a-bbc0b824ac80','stage','Stage','a597b201-1aa7-4965-8ff2-a1a26fef4763', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fb959ca2-30b7-4bf7-a94f-9e16770317ce','water-temperature','Water-Temperature','a597b201-1aa7-4965-8ff2-a1a26fef4763', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('da8687dc-c8a0-4689-bfd7-e25255cf5c62','voltage','Voltage','a597b201-1aa7-4965-8ff2-a1a26fef4763', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2ea409c2-27b6-427c-aed2-c5ee2bc1cf65','stage','Stage','5c797d3a-f2f6-4149-86b3-e87b3b20618e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0d4a4091-8a3c-4a50-9686-945897cb738d','elevation','Elevation','992677cd-af48-4ee9-82df-e05f17c8724d', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4c6dfbac-7218-4068-a5a1-6b0ffd56279e','stage','Stage','c844e536-5a97-4fa9-9e2c-d6502219c115', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('78a13713-0cac-4981-8a6e-5520c962ca88','precipitation','Precipitation','c844e536-5a97-4fa9-9e2c-d6502219c115', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('719dd307-102c-4d9a-ba3a-18edf39b1780','stage','Stage','653e2c8f-eac0-4d0a-882c-abe28355962e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4dc8cce4-26b2-4305-9a07-12c0dba8ef88','stage','Stage','e9e90579-80a4-4d53-ac11-689a920b887a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1ca12df7-3349-4ed5-a83c-08111c93ecc2','stage','Stage','b5564b22-e521-49b7-b398-2bf9bc1c5398', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4b4fafd1-38e4-4e18-882d-49515a582557','stage','Stage','c090102e-82b3-4a01-9bc0-ce73395d3226', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d3d2fb20-5596-476f-a775-76d19a556c65','stage','Stage','b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('84b7cde7-18ea-481d-b1e9-de3bdc34bd94','water-temperature','Water-Temperature','b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('ff9c9973-840b-4b59-a501-ca83fd5e6114','stage','Stage','622ba7b6-0bd2-4c26-9275-17b6462c525e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4ddc5ff9-b2e0-431e-8ff3-9c9500f5034c','precipitation','Precipitation','622ba7b6-0bd2-4c26-9275-17b6462c525e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('22ae188c-8adc-4148-be12-06085b97fe11','voltage','Voltage','622ba7b6-0bd2-4c26-9275-17b6462c525e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('33ef6b8e-d53b-4501-9e17-9dc0bb8e5d9f','elevation','Elevation','b48f1167-ef1a-469e-89d4-143263120130', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('da20d325-790c-4508-ba88-b202870d648c','stage','Stage','b89d5ce3-7d71-4958-8809-2f7e701a5d99', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('deea54db-52fe-45ca-a9b3-477d9b010d1e','water-temperature','Water-Temperature','b89d5ce3-7d71-4958-8809-2f7e701a5d99', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('af09de00-c68c-4808-b188-fac1580085de','stage','Stage','28b4878c-3e78-4898-84d2-c0d9a47d6a54', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('701d76c5-1103-4568-9c55-ccb68ee5d03c','precipitation','Precipitation','28b4878c-3e78-4898-84d2-c0d9a47d6a54', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('32686823-6f57-4311-bb5a-c60995ac0d2c','voltage','Voltage','28b4878c-3e78-4898-84d2-c0d9a47d6a54', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('266cc3ee-e65b-4261-a743-ac2e513cec08','stage','Stage','913eb762-d4be-491f-9fde-cbd054b1dfc0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('faa64cd7-5736-4bc8-b5e9-43c1be5e31dc','stage','Stage','c8877ff5-6a74-479a-8b8b-f2bc0a103a47', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('974a9d9b-d7cd-45ea-9531-2492c49ebc49','precipitation','Precipitation','c8877ff5-6a74-479a-8b8b-f2bc0a103a47', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f0562a73-cceb-40b9-adba-0d910d72218b','stage','Stage','11258ce5-b0f0-475f-b06e-f3f056d7e7b8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('feafe3b7-7308-4c2e-9ebc-ec3c25b3c30b','stage','Stage','0464c8ba-0e8c-4621-9d43-c91e90b011e9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4c718302-164e-4431-a1b5-c337223ba8dd','precipitation','Precipitation','0464c8ba-0e8c-4621-9d43-c91e90b011e9', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('cee45901-4406-40fe-8f5d-f1c4eac63708','stage','Stage','ebcdfa47-b1b6-4f49-87f4-85ac03221578', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1c2fa41c-9dfc-4232-8aa8-ea4f2c1d28e1','stage','Stage','8c23bf4f-9ac0-4022-bdc2-a730229b87a9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('23447b48-2434-4c5e-b1a3-202a113a1dd7','stage','Stage','7e8a59bd-7e19-45f7-8e75-b336a3bccad2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c9657b96-be49-4bff-a566-2ebb619a2cba','precipitation','Precipitation','7e8a59bd-7e19-45f7-8e75-b336a3bccad2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('2b34ea1b-29ad-46ff-ab61-10b01ca1e0ae','voltage','Voltage','7e8a59bd-7e19-45f7-8e75-b336a3bccad2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('9b87b8cb-8645-4c61-b01e-45cd536f3793','elevation','Elevation','28927945-3f8f-4dbe-8055-c92baa9df5a7', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('933f9369-afba-443c-96ed-157a0df9a542','precipitation','Precipitation','28927945-3f8f-4dbe-8055-c92baa9df5a7', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('73b75f80-d39b-43ad-9126-39692293c0ab','voltage','Voltage','28927945-3f8f-4dbe-8055-c92baa9df5a7', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('eb360a96-b601-460b-9c41-f495bb8fb23d','stage','Stage','6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('97e3452f-9d0f-4613-9692-85cccb0c2533','precipitation','Precipitation','6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f696678a-8570-464c-8dc7-52a2d74be3a1','voltage','Voltage','6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('41c4b9c8-4b83-4126-a433-71cfb802e9f0','elevation','Elevation','a3de6a31-67af-43e7-9ac2-d6042e6823c9', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b458ba85-5508-4f6c-9eba-ac880a47c484','precipitation','Precipitation','a3de6a31-67af-43e7-9ac2-d6042e6823c9', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('aedf7e8a-5e73-4caf-b3bb-81bfcb1fe6cc','voltage','Voltage','a3de6a31-67af-43e7-9ac2-d6042e6823c9', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0def5cff-909c-4094-aae6-a3ad718d93a1','stage','Stage','3e8f793f-9e7b-482d-b08d-b2354e727644', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('961d061f-c03e-4d76-a2e4-67bc591e07a2','precipitation','Precipitation','3e8f793f-9e7b-482d-b08d-b2354e727644', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d75b944f-111d-4230-85f3-1975e6d438de','voltage','Voltage','3e8f793f-9e7b-482d-b08d-b2354e727644', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('49fd5136-1bf2-4f82-9451-5b854c5fad53','elevation','Elevation','186ea37e-004c-40fd-8527-b5a288dc5764', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d034f5ce-48d0-4446-95b5-fa9474c3db46','precipitation','Precipitation','186ea37e-004c-40fd-8527-b5a288dc5764', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('82bb6990-050a-4949-885a-db9b4212fbc4','voltage','Voltage','186ea37e-004c-40fd-8527-b5a288dc5764', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('388e2809-dde1-48a5-8339-00357e83070a','unknown-elev-encoder','Unknown Elev-Encoder','1eab7035-942e-4e0c-87c5-b27e2e353a35', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('0e4df939-4b3e-4ed5-9c37-b3f393c5bcd2','unknown-elev-transducer','Unknown Elev-Transducer','1eab7035-942e-4e0c-87c5-b27e2e353a35', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b9dca5e8-5903-4fa1-88ed-5dc37412bbab','unknown-precip-inc','Unknown Precip-Inc','1eab7035-942e-4e0c-87c5-b27e2e353a35', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('98ca86d9-a336-4a2b-90fa-1a43812f5d69','unknown-temp','Unknown Temp','1eab7035-942e-4e0c-87c5-b27e2e353a35', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8'),
('b00f69f9-9d9d-4495-aa67-e3c75ac963bd','voltage','Voltage','1eab7035-942e-4e0c-87c5-b27e2e353a35', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3572e882-3ee7-465b-aad0-e2d8bd722c3f','elevation','Elevation','a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('13679bd5-7a54-4443-bf54-a3353ace7d5b','stage','Stage','4a7d6948-c69d-4401-8413-8576229a2fa4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c0a157c1-0a2b-4c5a-9538-0c2ee0071f8c','precipitation','Precipitation','4a7d6948-c69d-4401-8413-8576229a2fa4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f9ce95e6-0d24-4f3f-89db-89184263c17d','voltage','Voltage','4a7d6948-c69d-4401-8413-8576229a2fa4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('807ed049-cd11-4452-b7c6-cc272af362f6','stage','Stage','4347b406-e66e-4bd1-92d8-b322ab2e9254', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('71968001-d556-4520-959b-993cc0f9d5f7','stage','Stage','f79a66b0-6584-4d2b-8340-e014b0fa66e0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cb20e093-7186-4662-ba43-9e06e0e5c146','stage','Stage','cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('70809b44-522b-4688-b89a-23ec359b1e57','precipitation','Precipitation','cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('981bf21e-6b4c-49cb-b721-93b29a07f801','stage','Stage','ce302849-2d7d-499c-a862-6b077c42a0ee', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e1a21416-d21e-467a-8be1-389072c8fa31','precipitation','Precipitation','ce302849-2d7d-499c-a862-6b077c42a0ee', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('3966de59-a649-4248-b9db-6179e2100d3d','stage','Stage','45948095-3b78-4123-91f8-b86e95f28c42', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8bdb5e6b-a62e-4a2d-a1c8-58101b84847c','precipitation','Precipitation','45948095-3b78-4123-91f8-b86e95f28c42', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a');
