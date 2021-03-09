
INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('1fa28968-d0b1-4832-b4b3-e58206056819', 'd8f8934d-e414-499d-bd51-bc93bbde6345', 'chicago-district-streamgages', 'Chicago District Streamgages', 'chicago-district-streamgages.jpg');




--INSERT INSTRUMENTS--COUNT:40
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('6fbeb371-98db-4b70-93bb-831b14f72c7b', False, 'portland_74f4', 'Portland', null, ST_GeomFromText('POINT(-85.039 40.4277)',4326), null, null, '2021-03-08T19:49:06.533465Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324200'),
('1a47c87e-50d1-48ee-bb78-c246580263cb', False, 'linn-grove', 'Linn Grove', null, ST_GeomFromText('POINT(-85.0309 40.6439)',4326), null, null, '2021-03-08T19:49:06.533750Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03322900'),
('4deb6eee-81ed-4edd-ac38-0b081eb1a081', False, 'berlin', 'Berlin', null, ST_GeomFromText('POINT(-88.95 43.9539)',4326), null, null, '2021-03-08T19:49:06.533923Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04073500'),
('b3f286b8-60a3-43e0-9d8a-950f342388a4', False, 'obrien', 'Obrien', null, ST_GeomFromText('POINT(-87.5611 41.65)',4326), null, null, '2021-03-08T19:49:06.534147Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', False, 'fonddulac', 'FondDuLac', null, ST_GeomFromText('POINT(-88.4561 43.8)',4326), null, null, '2021-03-08T19:49:06.534702Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('005e3595-446b-40e1-b07c-604d2a808acc', False, 'menasha', 'Menasha', null, ST_GeomFromText('POINT(-88.4472 44.1994)',4326), null, null, '2021-03-08T19:49:06.535021Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('2155fea5-633e-4556-8c57-c49c8cf34cc1', False, 'lockport', 'Lockport', null, ST_GeomFromText('POINT(-88.0789 41.5697)',4326), null, null, '2021-03-08T19:49:06.535321Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('a597b201-1aa7-4965-8ff2-a1a26fef4763', False, 'mississinewa-tailwater', 'Mississinewa-Tailwater', null, ST_GeomFromText('POINT(-85.9575 40.7233)',4326), null, null, '2021-03-08T19:49:06.535718Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327000'),
('5c797d3a-f2f6-4149-86b3-e87b3b20618e', False, 'marion', 'Marion', null, ST_GeomFromText('POINT(-85.6595 40.5764)',4326), null, null, '2021-03-08T19:49:06.535923Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326500'),
('992677cd-af48-4ee9-82df-e05f17c8724d', False, 'mississinewa-pool', 'Mississinewa-Pool', null, ST_GeomFromText('POINT(-85.9572 40.7144)',4326), null, null, '2021-03-08T19:49:06.536023Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326950'),
('c844e536-5a97-4fa9-9e2c-d6502219c115', False, 'bluffton', 'Bluffton', null, ST_GeomFromText('POINT(-85.1714 40.7424)',4326), null, null, '2021-03-08T19:49:06.536111Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04099510'),
('653e2c8f-eac0-4d0a-882c-abe28355962e', False, 'jerome', 'Jerome', null, ST_GeomFromText('POINT(-85.9188 40.4413)',4326), null, null, '2021-03-08T19:49:06.536242Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333450'),
('e9e90579-80a4-4d53-ac11-689a920b887a', False, 'kokomo', 'Kokomo', null, ST_GeomFromText('POINT(-86.1529 40.4709)',4326), null, null, '2021-03-08T19:49:06.536328Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333700'),
('b5564b22-e521-49b7-b398-2bf9bc1c5398', False, 'owasco', 'Owasco', null, ST_GeomFromText('POINT(-86.6366 40.4648)',4326), null, null, '2021-03-08T19:49:06.536411Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03334000'),
('c090102e-82b3-4a01-9bc0-ce73395d3226', False, 'deer-creek', 'Deer Creek', null, ST_GeomFromText('POINT(-86.6214 40.5903)',4326), null, null, '2021-03-08T19:49:06.536500Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329700'),
('b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', False, 'salamonie-tailwater', 'Salamonie-Tailwater', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-08T19:49:06.536585Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324500'),
('622ba7b6-0bd2-4c26-9275-17b6462c525e', False, 'warren_0955', 'Warren', null, ST_GeomFromText('POINT(-85.4536 40.7125)',4326), null, null, '2021-03-08T19:49:06.536751Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324300'),
('b48f1167-ef1a-469e-89d4-143263120130', False, 'j-edward-roush-pool', 'J Edward Roush-Pool', null, ST_GeomFromText('POINT(-85.4686 40.8461)',4326), null, null, '2021-03-08T19:49:06.536975Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323450'),
('b89d5ce3-7d71-4958-8809-2f7e701a5d99', False, 'j-edward-roush-tailwater', 'J Edward Roush-Tailwater', null, ST_GeomFromText('POINT(-85.4898 40.8533)',4326), null, null, '2021-03-08T19:49:06.537079Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323500'),
('28b4878c-3e78-4898-84d2-c0d9a47d6a54', False, 'wabash', 'Wabash', null, ST_GeomFromText('POINT(-85.8203 40.7908)',4326), null, null, '2021-03-08T19:49:06.537205Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03325000'),
('913eb762-d4be-491f-9fde-cbd054b1dfc0', False, 'lafayettewildcat', 'LafayetteWildcat', null, ST_GeomFromText('POINT(-86.8292 40.4406)',4326), null, null, '2021-03-08T19:49:06.537360Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335000'),
('c8877ff5-6a74-479a-8b8b-f2bc0a103a47', False, 'delphi', 'Delphi', null, ST_GeomFromText('POINT(-86.7703 40.5939)',4326), null, null, '2021-03-08T19:49:06.537445Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333050'),
('11258ce5-b0f0-475f-b06e-f3f056d7e7b8', False, 'oswego', 'Oswego', null, ST_GeomFromText('POINT(-85.7892 41.3206)',4326), null, null, '2021-03-08T19:49:06.537563Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330500'),
('0464c8ba-0e8c-4621-9d43-c91e90b011e9', False, 'ora', 'Ora', null, ST_GeomFromText('POINT(-86.5636 41.1572)',4326), null, null, '2021-03-08T19:49:06.537647Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03331500'),
('ebcdfa47-b1b6-4f49-87f4-85ac03221578', False, 'north-manchester', 'North Manchester', null, ST_GeomFromText('POINT(-85.7825 40.9944)',4326), null, null, '2021-03-08T19:49:06.537766Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03328000'),
('8c23bf4f-9ac0-4022-bdc2-a730229b87a9', False, 'north-webster', 'North Webster', null, ST_GeomFromText('POINT(-85.6922 41.3164)',4326), null, null, '2021-03-08T19:49:06.537851Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330241'),
('7e8a59bd-7e19-45f7-8e75-b336a3bccad2', False, 'newlondon', 'NewLondon', null, ST_GeomFromText('POINT(-88.7403 44.3922)',4326), null, null, '2021-03-08T19:49:06.537939Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04079000'),
('28927945-3f8f-4dbe-8055-c92baa9df5a7', False, 'oshkosh', 'Oshkosh', null, ST_GeomFromText('POINT(-88.5272 44.0097)',4326), null, null, '2021-03-08T19:49:06.538091Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04082500'),
('6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', False, 'royalton', 'Royalton', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-08T19:49:06.538241Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04080000'),
('a3de6a31-67af-43e7-9ac2-d6042e6823c9', False, 'stockbridge', 'Stockbridge', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-08T19:49:06.538392Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04084255'),
('3e8f793f-9e7b-482d-b08d-b2354e727644', False, 'waupaca', 'Waupaca', null, ST_GeomFromText('POINT(-88.9961 44.3292)',4326), null, null, '2021-03-08T19:49:06.538541Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04081000'),
('186ea37e-004c-40fd-8527-b5a288dc5764', False, 'fritsepark', 'FritsePark', null, ST_GeomFromText('POINT(-88.4703 44.205)',4326), null, null, '2021-03-08T19:49:06.538764Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('1eab7035-942e-4e0c-87c5-b27e2e353a35', False, 'poygan', 'Poygan', null, ST_GeomFromText('POINT(-88.7125 44.1108)',4326), null, null, '2021-03-08T19:49:06.538948Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', False, 'salamonie-pool', 'Salamonie-Pool', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-08T19:49:06.539182Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324450'),
('4a7d6948-c69d-4401-8413-8576229a2fa4', False, 'peru', 'Peru', null, ST_GeomFromText('POINT(-86.0667 40.75)',4326), null, null, '2021-03-08T19:49:06.539296Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327500'),
('4347b406-e66e-4bd1-92d8-b322ab2e9254', False, 'logansport', 'Logansport', null, ST_GeomFromText('POINT(-86.3775 40.7464)',4326), null, null, '2021-03-08T19:49:06.539448Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329000'),
('f79a66b0-6584-4d2b-8340-e014b0fa66e0', False, 'lafayette', 'Lafayette', null, ST_GeomFromText('POINT(-86.8969 40.4219)',4326), null, null, '2021-03-08T19:49:06.539532Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335500'),
('cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', False, 'littleriver', 'LittleRiver', null, ST_GeomFromText('POINT(-85.4132 40.8986)',4326), null, null, '2021-03-08T19:49:06.539616Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324000'),
('ce302849-2d7d-499c-a862-6b077c42a0ee', False, 'covington', 'Covington', null, ST_GeomFromText('POINT(-87.4067 40.14)',4326), null, null, '2021-03-08T19:49:06.539733Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03336000'),
('45948095-3b78-4123-91f8-b86e95f28c42', False, 'montezuma', 'Montezuma', null, ST_GeomFromText('POINT(-87.3739 39.7925)',4326), null, null, '2021-03-08T19:49:06.539855Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03340500');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('7331e3f5-8d04-4652-8e9c-104805259095', '6fbeb371-98db-4b70-93bb-831b14f72c7b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.533465Z'),
('b337351e-8f0d-46b8-a814-a742bca3fd4b', '1a47c87e-50d1-48ee-bb78-c246580263cb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.533750Z'),
('c193f6e3-7e6d-4973-bfee-80976b46c36b', '4deb6eee-81ed-4edd-ac38-0b081eb1a081', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.533923Z'),
('a7d10d38-c214-4674-9b0a-bd3e53f0f33d', 'b3f286b8-60a3-43e0-9d8a-950f342388a4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.534147Z'),
('f75da3a0-7e4b-4eee-8c50-f9a95c1910a4', '3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.534702Z'),
('8bc011e6-b5a2-48e3-bc17-786c23b459f4', '005e3595-446b-40e1-b07c-604d2a808acc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.535021Z'),
('12c82a09-ae54-4936-8547-f789885c3289', '2155fea5-633e-4556-8c57-c49c8cf34cc1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.535321Z'),
('4ee76e26-331a-4c0f-ab4f-eda94c345baa', 'a597b201-1aa7-4965-8ff2-a1a26fef4763', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.535718Z'),
('91f67238-bc31-44e6-a5aa-08b1a7d0c07a', '5c797d3a-f2f6-4149-86b3-e87b3b20618e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.535923Z'),
('972eb2a1-1041-45c2-8ff2-b953e1130053', '992677cd-af48-4ee9-82df-e05f17c8724d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536023Z'),
('a92c9e06-85b2-4eec-aa90-12646a5f143d', 'c844e536-5a97-4fa9-9e2c-d6502219c115', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536111Z'),
('a300f1ae-cd81-4889-a4ab-cb97339ae7af', '653e2c8f-eac0-4d0a-882c-abe28355962e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536242Z'),
('61a05171-b6e3-48d4-ab34-dda96a35d909', 'e9e90579-80a4-4d53-ac11-689a920b887a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536328Z'),
('3bbc28a7-b0d0-4485-ab05-ba54318803f5', 'b5564b22-e521-49b7-b398-2bf9bc1c5398', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536411Z'),
('7738ec90-e553-4ed4-878e-c88a4b3bbaba', 'c090102e-82b3-4a01-9bc0-ce73395d3226', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536500Z'),
('9a7787bc-5b08-434b-a360-ef0a0fccfa0a', 'b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536585Z'),
('dc418d23-5ad6-445b-99c0-7fe314865abd', '622ba7b6-0bd2-4c26-9275-17b6462c525e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536751Z'),
('51ee7cbe-54b8-4ee9-9246-fac826beecbf', 'b48f1167-ef1a-469e-89d4-143263120130', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.536975Z'),
('5e1bf0e2-931d-41f0-b4c8-e07487d0d830', 'b89d5ce3-7d71-4958-8809-2f7e701a5d99', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537079Z'),
('c429c56b-b2af-438e-8d9f-b6c58758de00', '28b4878c-3e78-4898-84d2-c0d9a47d6a54', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537205Z'),
('e55cf6c8-68a4-4929-87c6-df7b9ad97925', '913eb762-d4be-491f-9fde-cbd054b1dfc0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537360Z'),
('6676f490-2869-456a-b308-782d729ca639', 'c8877ff5-6a74-479a-8b8b-f2bc0a103a47', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537445Z'),
('3b8dba86-a394-4ba6-ae6e-3f7437f728a7', '11258ce5-b0f0-475f-b06e-f3f056d7e7b8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537563Z'),
('89161402-cc92-43ad-999e-2360cc429fa0', '0464c8ba-0e8c-4621-9d43-c91e90b011e9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537647Z'),
('1ac31007-65db-442a-9754-ad36071d53e7', 'ebcdfa47-b1b6-4f49-87f4-85ac03221578', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537766Z'),
('de0a6ead-876a-466e-858a-8c55c6bb97e9', '8c23bf4f-9ac0-4022-bdc2-a730229b87a9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537851Z'),
('c2398c03-98a9-4494-97e7-1911996d0803', '7e8a59bd-7e19-45f7-8e75-b336a3bccad2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.537939Z'),
('d54a91c4-9aa9-4dea-ae3e-2137966ebffb', '28927945-3f8f-4dbe-8055-c92baa9df5a7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538091Z'),
('2f79e84e-e0bc-4938-9843-fab560dd0312', '6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538241Z'),
('7c9a6625-cfec-4088-99fa-ee259ea32086', 'a3de6a31-67af-43e7-9ac2-d6042e6823c9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538392Z'),
('8e12855c-0047-4fe0-a187-501b42fef466', '3e8f793f-9e7b-482d-b08d-b2354e727644', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538541Z'),
('14f81f5f-a914-4481-946f-6c392150c1ef', '186ea37e-004c-40fd-8527-b5a288dc5764', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538764Z'),
('73027d7d-4b12-4cc4-9f09-041a1ea2eb02', '1eab7035-942e-4e0c-87c5-b27e2e353a35', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.538948Z'),
('f94109f6-6d7a-47fc-99f8-490b0c4ac71d', 'a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539182Z'),
('d4e0dccf-33e4-4222-8020-5bce5839a035', '4a7d6948-c69d-4401-8413-8576229a2fa4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539296Z'),
('72a1b5d0-b65c-4240-8708-82e757328052', '4347b406-e66e-4bd1-92d8-b322ab2e9254', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539448Z'),
('8557c9d0-6324-45cd-b0f3-fca9cba70765', 'f79a66b0-6584-4d2b-8340-e014b0fa66e0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539532Z'),
('e60deb2e-d939-4c1b-9821-0d108bc7c5fe', 'cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539616Z'),
('f0a04d82-9719-49d5-b75f-34980c3e5a73', 'ce302849-2d7d-499c-a862-6b077c42a0ee', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539733Z'),
('0de35989-4b77-4458-bbb2-eb2a61ebefd7', '45948095-3b78-4123-91f8-b86e95f28c42', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:49:06.539855Z');

--INSERT TELEMETRY_GOES--COUNT:40
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7752a1c3-347e-4dc9-a866-f6908a315097', 'DD46A45A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD46A45A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '14259230-aaa0-4c6a-945a-87b8f98a260a', 'CE6B687C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B687C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '12caab07-a9b7-468e-85fc-b7f978d255d4', 'CE72E178' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72E178');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd33bcbc8-98ef-4544-b508-acb6f7462c34', 'CE2A970C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE2A970C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '304aca92-bd0b-4d84-aa62-70e2d287291b', 'CE72DA30' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72DA30');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '51150d98-8f90-4b48-be6f-c99baf706849', 'CE72D4E2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72D4E2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9ee51fdf-dead-4b98-9835-b94a65519981', 'CE720C58' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE720C58');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd17618fc-89af-419f-8fb2-6e8e0c5377ad', 'CE7723A6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7723A6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '57dff2ba-5aa0-4a5f-8591-cf6d1aa8914c', 'CE6B5DE6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B5DE6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd7d3cfd0-9022-4322-a728-a110260ea86e', 'CE108358' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE108358');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '805093c2-6b0e-4065-bfee-a31c4b11ec3c', '17D2F560' where not exists (select 1 from telemetry_goes where nesdis_id = '17D2F560');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '28d1460a-40b2-4546-ac18-df1337ce347f', 'D11EE2FE' where not exists (select 1 from telemetry_goes where nesdis_id = 'D11EE2FE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '01348d10-738d-4d81-b2f4-6c23906dd47d', '173623B8' where not exists (select 1 from telemetry_goes where nesdis_id = '173623B8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0438f5c6-f0b2-43b3-b16e-1a9edcb4638f', '167526C8' where not exists (select 1 from telemetry_goes where nesdis_id = '167526C8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd8224ed1-4a67-44ca-b527-dbbc7c4017c6', 'DE4B8336' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE4B8336');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd2ee1a8b-c848-4095-8dbc-295385639a55', 'CE7718EE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7718EE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b701fa86-7246-4169-b7a3-573995032674', 'CE6B66AE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B66AE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '4d14cc99-a16a-4cbb-8a9b-85f0939cda5c', 'CE10C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE10C052');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2290d4da-69c2-49fb-b771-9f6e9de7c8ac', 'CE77163C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE77163C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7f8b1972-4297-4597-a2c0-6383705dbc01', 'CE16B60C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16B60C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'fb3cbf48-5dde-49d2-935c-da73461f9246', 'CE6B0D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B0D9A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ec04fc0d-c7bb-4eef-9460-506f53c4ec88', 'DDA6F1AC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA6F1AC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '06075bb1-76fd-476c-aaf4-f9a8f8ee13d2', '17853608' where not exists (select 1 from telemetry_goes where nesdis_id = '17853608');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c6c6beec-6c3a-4c48-adc3-a3e170fe742f', 'DD8F700A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8F700A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5c02081e-c12d-4675-a77b-8c9462cee12e', '163D923C' where not exists (select 1 from telemetry_goes where nesdis_id = '163D923C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f8e42aca-8bb9-40b9-b39a-d8fc3d89bf29', 'DD3E3032' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD3E3032');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6bfc7304-1a9c-4fa5-b377-830929110bf1', 'CE72FCDC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72FCDC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1d21fffe-f3dc-4bc8-b34b-bfd7e9f70211', 'CE58F2B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE58F2B2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0c1dcb7b-2f04-4337-ade7-f2ecaba8216a', 'CE72EFAA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72EFAA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c5c44e4e-cdba-4ee6-b178-333da04f93e8', 'CE72F20E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72F20E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ccb26ca2-4964-4a54-acd4-a8a03552ba29', 'CE730070' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE730070');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'aa4dc500-0b89-4f69-83d6-110847fe9ee6', 'CE26C6EC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE26C6EC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0eb1ea04-1b3c-4a43-a1df-e2aeac508558', 'CE72C946' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72C946');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd505852a-d004-4be2-af73-28ff9ed97925', 'CE6D2BB8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D2BB8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8c912f7b-7eb6-46d2-b06c-59e2cfe7f421', 'CE777D08' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE777D08');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'efc2ef68-d3b0-4450-90b6-3cc8963950de', 'CE6D256A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D256A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'cac73c04-575d-49a4-b2b7-eab162def815', 'CE6D361C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D361C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '689785ba-d5b8-4bb8-85cd-11a9a643886e', 'CE14B3F8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14B3F8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6169136a-e059-4db2-ae32-e6859bb6d540', 'CE16C09C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16C09C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd1568af4-1293-47e0-8d7e-39523cca2acc', 'CE14E384' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14E384');

--INSERT INSTRUMENT_TELEMETRY--COUNT:40
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('6fbeb371-98db-4b70-93bb-831b14f72c7b', '10a32652-af43-4451-bd52-4980c5690cc9', '7752a1c3-347e-4dc9-a866-f6908a315097'),
('1a47c87e-50d1-48ee-bb78-c246580263cb', '10a32652-af43-4451-bd52-4980c5690cc9', '14259230-aaa0-4c6a-945a-87b8f98a260a'),
('4deb6eee-81ed-4edd-ac38-0b081eb1a081', '10a32652-af43-4451-bd52-4980c5690cc9', '12caab07-a9b7-468e-85fc-b7f978d255d4'),
('b3f286b8-60a3-43e0-9d8a-950f342388a4', '10a32652-af43-4451-bd52-4980c5690cc9', 'd33bcbc8-98ef-4544-b508-acb6f7462c34'),
('3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '10a32652-af43-4451-bd52-4980c5690cc9', '304aca92-bd0b-4d84-aa62-70e2d287291b'),
('005e3595-446b-40e1-b07c-604d2a808acc', '10a32652-af43-4451-bd52-4980c5690cc9', '51150d98-8f90-4b48-be6f-c99baf706849'),
('2155fea5-633e-4556-8c57-c49c8cf34cc1', '10a32652-af43-4451-bd52-4980c5690cc9', '9ee51fdf-dead-4b98-9835-b94a65519981'),
('a597b201-1aa7-4965-8ff2-a1a26fef4763', '10a32652-af43-4451-bd52-4980c5690cc9', 'd17618fc-89af-419f-8fb2-6e8e0c5377ad'),
('5c797d3a-f2f6-4149-86b3-e87b3b20618e', '10a32652-af43-4451-bd52-4980c5690cc9', '57dff2ba-5aa0-4a5f-8591-cf6d1aa8914c'),
('992677cd-af48-4ee9-82df-e05f17c8724d', '10a32652-af43-4451-bd52-4980c5690cc9', 'd7d3cfd0-9022-4322-a728-a110260ea86e'),
('c844e536-5a97-4fa9-9e2c-d6502219c115', '10a32652-af43-4451-bd52-4980c5690cc9', '805093c2-6b0e-4065-bfee-a31c4b11ec3c'),
('653e2c8f-eac0-4d0a-882c-abe28355962e', '10a32652-af43-4451-bd52-4980c5690cc9', '28d1460a-40b2-4546-ac18-df1337ce347f'),
('e9e90579-80a4-4d53-ac11-689a920b887a', '10a32652-af43-4451-bd52-4980c5690cc9', '01348d10-738d-4d81-b2f4-6c23906dd47d'),
('b5564b22-e521-49b7-b398-2bf9bc1c5398', '10a32652-af43-4451-bd52-4980c5690cc9', '0438f5c6-f0b2-43b3-b16e-1a9edcb4638f'),
('c090102e-82b3-4a01-9bc0-ce73395d3226', '10a32652-af43-4451-bd52-4980c5690cc9', 'd8224ed1-4a67-44ca-b527-dbbc7c4017c6'),
('b1c0bd15-00a4-4dc4-9445-ef9aed4068f0', '10a32652-af43-4451-bd52-4980c5690cc9', 'd2ee1a8b-c848-4095-8dbc-295385639a55'),
('622ba7b6-0bd2-4c26-9275-17b6462c525e', '10a32652-af43-4451-bd52-4980c5690cc9', 'b701fa86-7246-4169-b7a3-573995032674'),
('b48f1167-ef1a-469e-89d4-143263120130', '10a32652-af43-4451-bd52-4980c5690cc9', '4d14cc99-a16a-4cbb-8a9b-85f0939cda5c'),
('b89d5ce3-7d71-4958-8809-2f7e701a5d99', '10a32652-af43-4451-bd52-4980c5690cc9', '2290d4da-69c2-49fb-b771-9f6e9de7c8ac'),
('28b4878c-3e78-4898-84d2-c0d9a47d6a54', '10a32652-af43-4451-bd52-4980c5690cc9', '7f8b1972-4297-4597-a2c0-6383705dbc01'),
('913eb762-d4be-491f-9fde-cbd054b1dfc0', '10a32652-af43-4451-bd52-4980c5690cc9', 'fb3cbf48-5dde-49d2-935c-da73461f9246'),
('c8877ff5-6a74-479a-8b8b-f2bc0a103a47', '10a32652-af43-4451-bd52-4980c5690cc9', 'ec04fc0d-c7bb-4eef-9460-506f53c4ec88'),
('11258ce5-b0f0-475f-b06e-f3f056d7e7b8', '10a32652-af43-4451-bd52-4980c5690cc9', '06075bb1-76fd-476c-aaf4-f9a8f8ee13d2'),
('0464c8ba-0e8c-4621-9d43-c91e90b011e9', '10a32652-af43-4451-bd52-4980c5690cc9', 'c6c6beec-6c3a-4c48-adc3-a3e170fe742f'),
('ebcdfa47-b1b6-4f49-87f4-85ac03221578', '10a32652-af43-4451-bd52-4980c5690cc9', '5c02081e-c12d-4675-a77b-8c9462cee12e'),
('8c23bf4f-9ac0-4022-bdc2-a730229b87a9', '10a32652-af43-4451-bd52-4980c5690cc9', 'f8e42aca-8bb9-40b9-b39a-d8fc3d89bf29'),
('7e8a59bd-7e19-45f7-8e75-b336a3bccad2', '10a32652-af43-4451-bd52-4980c5690cc9', '6bfc7304-1a9c-4fa5-b377-830929110bf1'),
('28927945-3f8f-4dbe-8055-c92baa9df5a7', '10a32652-af43-4451-bd52-4980c5690cc9', '1d21fffe-f3dc-4bc8-b34b-bfd7e9f70211'),
('6c106d17-d0c1-4a10-a5ef-59f20dbbfe2c', '10a32652-af43-4451-bd52-4980c5690cc9', '0c1dcb7b-2f04-4337-ade7-f2ecaba8216a'),
('a3de6a31-67af-43e7-9ac2-d6042e6823c9', '10a32652-af43-4451-bd52-4980c5690cc9', 'c5c44e4e-cdba-4ee6-b178-333da04f93e8'),
('3e8f793f-9e7b-482d-b08d-b2354e727644', '10a32652-af43-4451-bd52-4980c5690cc9', 'ccb26ca2-4964-4a54-acd4-a8a03552ba29'),
('186ea37e-004c-40fd-8527-b5a288dc5764', '10a32652-af43-4451-bd52-4980c5690cc9', 'aa4dc500-0b89-4f69-83d6-110847fe9ee6'),
('1eab7035-942e-4e0c-87c5-b27e2e353a35', '10a32652-af43-4451-bd52-4980c5690cc9', '0eb1ea04-1b3c-4a43-a1df-e2aeac508558'),
('a4e8f09d-08b8-4fa7-8ccc-30757e0962f4', '10a32652-af43-4451-bd52-4980c5690cc9', 'd505852a-d004-4be2-af73-28ff9ed97925'),
('4a7d6948-c69d-4401-8413-8576229a2fa4', '10a32652-af43-4451-bd52-4980c5690cc9', '8c912f7b-7eb6-46d2-b06c-59e2cfe7f421'),
('4347b406-e66e-4bd1-92d8-b322ab2e9254', '10a32652-af43-4451-bd52-4980c5690cc9', 'efc2ef68-d3b0-4450-90b6-3cc8963950de'),
('f79a66b0-6584-4d2b-8340-e014b0fa66e0', '10a32652-af43-4451-bd52-4980c5690cc9', 'cac73c04-575d-49a4-b2b7-eab162def815'),
('cf8614a1-726f-4cf3-a576-8c4a8cf65cfb', '10a32652-af43-4451-bd52-4980c5690cc9', '689785ba-d5b8-4bb8-85cd-11a9a643886e'),
('ce302849-2d7d-499c-a862-6b077c42a0ee', '10a32652-af43-4451-bd52-4980c5690cc9', '6169136a-e059-4db2-ae32-e6859bb6d540'),
('45948095-3b78-4123-91f8-b86e95f28c42', '10a32652-af43-4451-bd52-4980c5690cc9', 'd1568af4-1293-47e0-8d7e-39523cca2acc');

--INSERT TIMESERIES--COUNT:40
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('19ee54b3-1eaa-43d2-b209-ed3f250077f9','stage','Stage','6fbeb371-98db-4b70-93bb-831b14f72c7b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f2ca7fc2-a66d-470a-a5d5-374cb719d7ba','stage','Stage','1a47c87e-50d1-48ee-bb78-c246580263cb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d57653ee-b3c7-4a6b-8b98-71781a41e7ce','precipitation','Precipitation','1a47c87e-50d1-48ee-bb78-c246580263cb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d6b0a480-c2b3-4a2d-a8bc-6de0dc3dd51b','stage','Stage','4deb6eee-81ed-4edd-ac38-0b081eb1a081', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('480c0509-7df7-4206-986e-ca2de854c2c1','precipitation','Precipitation','4deb6eee-81ed-4edd-ac38-0b081eb1a081', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f16d84ab-cb78-4ac9-941f-52c8f9f4ff33','precipitation','Precipitation','b3f286b8-60a3-43e0-9d8a-950f342388a4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('46462ef5-b7fa-4741-a0bb-f642f59d524b','elevation','Elevation','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e31b80eb-702e-4cc7-9f92-8b3d3a0f57b7','precipitation','Precipitation','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('660cd9f0-7110-4b16-9207-3e950ae04528','voltage','Voltage','3e29ecea-8f9e-439f-ad99-aaff1b7e1cbc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2b5f5152-8e91-4ea6-9e22-685b8280f53b','elevation','Elevation','005e3595-446b-40e1-b07c-604d2a808acc', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0fad1ae4-b422-444b-8733-9c0bbd97d914','precipitation','Precipitation','005e3595-446b-40e1-b07c-604d2a808acc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('cc38d3d8-662d-43fd-9746-2f4d33f80774','voltage','Voltage','005e3595-446b-40e1-b07c-604d2a808acc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cc3cd3aa-07cd-4e30-8d86-96513ab21901','elevation','Elevation','2155fea5-633e-4556-8c57-c49c8cf34cc1', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2b778059-22bd-48af-847b-8769603aead1','precipitation','Precipitation','2155fea5-633e-4556-8c57-c49c8cf34cc1', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
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
