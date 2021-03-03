
INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('1fa28968-d0b1-4832-b4b3-e58206056819', 'd8f8934d-e414-499d-bd51-bc93bbde6345', 'chicago-district-streamgages', 'Chicago District Streamgages', 'chicago-district-streamgages.jpg');



--INSERT INSTRUMENTS--COUNT:40
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('f612a243-1673-4451-899d-5fd44c4a4f77', False, 'portland_55de', 'Portland', null, ST_GeomFromText('POINT(-85.039 40.4277)',4326), null, null, '2021-03-03T14:02:23.737872Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324200'),
('38c604e4-dd58-4ac1-9e4d-d456318b1705', False, 'linn-grove', 'Linn Grove', null, ST_GeomFromText('POINT(-85.0309 40.6439)',4326), null, null, '2021-03-03T14:02:23.738087Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03322900'),
('14e312f7-bb7e-4f88-afb1-8da275148025', False, 'berlin', 'Berlin', null, ST_GeomFromText('POINT(-88.95 43.9539)',4326), null, null, '2021-03-03T14:02:23.738226Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04073500'),
('edb24776-da77-4962-98ed-15515ad8b7ed', False, 'obrien', 'Obrien', null, ST_GeomFromText('POINT(-87.5611 41.65)',4326), null, null, '2021-03-03T14:02:23.738402Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('67ec7930-57ba-4d08-9280-d428e995cd59', False, 'fonddulac', 'FondDuLac', null, ST_GeomFromText('POINT(-88.4561 43.8)',4326), null, null, '2021-03-03T14:02:23.738785Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('52178424-b9be-46c7-9f60-305ab9d2b27d', False, 'menasha', 'Menasha', null, ST_GeomFromText('POINT(-88.4472 44.1994)',4326), null, null, '2021-03-03T14:02:23.738983Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('9a09251e-fff3-4d02-83e9-e7a539185376', False, 'lockport', 'Lockport', null, ST_GeomFromText('POINT(-88.0789 41.5697)',4326), null, null, '2021-03-03T14:02:23.739182Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', False, 'mississinewa-tailwater', 'Mississinewa-Tailwater', null, ST_GeomFromText('POINT(-85.9575 40.7233)',4326), null, null, '2021-03-03T14:02:23.739549Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327000'),
('800fc3e6-a171-4e99-955c-5e0d264ffd29', False, 'marion', 'Marion', null, ST_GeomFromText('POINT(-85.6595 40.5764)',4326), null, null, '2021-03-03T14:02:23.739722Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326500'),
('38ef8580-6278-4942-bb35-b1d3ac24412c', False, 'mississinewa-pool', 'Mississinewa-Pool', null, ST_GeomFromText('POINT(-85.9572 40.7144)',4326), null, null, '2021-03-03T14:02:23.739804Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03326950'),
('969c94eb-70bd-45c1-b23e-17399faa4605', False, 'bluffton', 'Bluffton', null, ST_GeomFromText('POINT(-85.1714 40.7424)',4326), null, null, '2021-03-03T14:02:23.739877Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04099510'),
('3631051d-ead4-46e8-a09b-42f5cee9ef0d', False, 'jerome', 'Jerome', null, ST_GeomFromText('POINT(-85.9188 40.4413)',4326), null, null, '2021-03-03T14:02:23.739985Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333450'),
('62ba81af-0197-4a1d-b0f8-cc7d8e5ba1e5', False, 'kokomo', 'Kokomo', null, ST_GeomFromText('POINT(-86.1529 40.4709)',4326), null, null, '2021-03-03T14:02:23.740057Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333700'),
('d5b41092-230b-414e-b2f2-57f7a792c6cc', False, 'owasco', 'Owasco', null, ST_GeomFromText('POINT(-86.6366 40.4648)',4326), null, null, '2021-03-03T14:02:23.740127Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03334000'),
('48fe98a9-38ae-487e-9f26-e0cc37b9699d', False, 'deer-creek', 'Deer Creek', null, ST_GeomFromText('POINT(-86.6214 40.5903)',4326), null, null, '2021-03-03T14:02:23.740201Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329700'),
('1668b0f0-ce59-4ca3-82cc-e2392a983c41', False, 'salamonie-tailwater', 'Salamonie-Tailwater', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-03T14:02:23.740273Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324500'),
('6cea522c-9e43-43a0-8e0e-02f82d41ae06', False, 'warren_ddf3', 'Warren', null, ST_GeomFromText('POINT(-85.4536 40.7125)',4326), null, null, '2021-03-03T14:02:23.740449Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324300'),
('bfc6ab1f-801a-4039-81e0-2a3755e0bb23', False, 'j-edward-roush-pool', 'J Edward Roush-Pool', null, ST_GeomFromText('POINT(-85.4686 40.8461)',4326), null, null, '2021-03-03T14:02:23.740609Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323450'),
('ce3729fc-a175-49e3-938b-f2bd0b92e61c', False, 'j-edward-roush-tailwater', 'J Edward Roush-Tailwater', null, ST_GeomFromText('POINT(-85.4898 40.8533)',4326), null, null, '2021-03-03T14:02:23.740686Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03323500'),
('3ead0f70-dab6-45f4-9878-bb80b45191ad', False, 'wabash', 'Wabash', null, ST_GeomFromText('POINT(-85.8203 40.7908)',4326), null, null, '2021-03-03T14:02:23.740788Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03325000'),
('ede9302f-6dbf-4493-97b2-884846f74570', False, 'lafayettewildcat', 'LafayetteWildcat', null, ST_GeomFromText('POINT(-86.8292 40.4406)',4326), null, null, '2021-03-03T14:02:23.740919Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335000'),
('d083e4cb-9269-4c73-8a4f-27bf87f675c9', False, 'delphi', 'Delphi', null, ST_GeomFromText('POINT(-86.7703 40.5939)',4326), null, null, '2021-03-03T14:02:23.740992Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03333050'),
('dbeeba6e-b039-4a07-9d2e-a1988698324e', False, 'oswego', 'Oswego', null, ST_GeomFromText('POINT(-85.7892 41.3206)',4326), null, null, '2021-03-03T14:02:23.741095Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330500'),
('7438be47-82ea-4e31-9a6a-ebe60c308d8f', False, 'ora', 'Ora', null, ST_GeomFromText('POINT(-86.5636 41.1572)',4326), null, null, '2021-03-03T14:02:23.741168Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03331500'),
('d37eb7f9-1453-47cd-b8bb-505d32beb899', False, 'north-manchester', 'North Manchester', null, ST_GeomFromText('POINT(-85.7825 40.9944)',4326), null, null, '2021-03-03T14:02:23.741271Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03328000'),
('f89c9bc9-8f33-4dbe-b225-a58e17176f5f', False, 'north-webster', 'North Webster', null, ST_GeomFromText('POINT(-85.6922 41.3164)',4326), null, null, '2021-03-03T14:02:23.741344Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03330241'),
('72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', False, 'newlondon', 'NewLondon', null, ST_GeomFromText('POINT(-88.7403 44.3922)',4326), null, null, '2021-03-03T14:02:23.741418Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04079000'),
('0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', False, 'oshkosh', 'Oshkosh', null, ST_GeomFromText('POINT(-88.5272 44.0097)',4326), null, null, '2021-03-03T14:02:23.741546Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04082500'),
('254210ea-46fd-4bd6-bf8e-9886b628bde0', False, 'royalton', 'Royalton', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-03T14:02:23.741674Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04080000'),
('22c0eb9c-ccf5-4267-9280-3e5f36497a60', False, 'stockbridge', 'Stockbridge', null, ST_GeomFromText('POINT(-88.8653 44.4125)',4326), null, null, '2021-03-03T14:02:23.741806Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04084255'),
('4fa4f306-95b4-4ec2-b4b1-4f3f87576804', False, 'waupaca', 'Waupaca', null, ST_GeomFromText('POINT(-88.9961 44.3292)',4326), null, null, '2021-03-03T14:02:23.741933Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '04081000'),
('1be8618a-27cc-4f01-9028-779daa13ec2c', False, 'fritsepark', 'FritsePark', null, ST_GeomFromText('POINT(-88.4703 44.205)',4326), null, null, '2021-03-03T14:02:23.742060Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('4b45639e-54b7-4cd7-aff5-8ad04fe244cf', False, 'poygan', 'Poygan', null, ST_GeomFromText('POINT(-88.7125 44.1108)',4326), null, null, '2021-03-03T14:02:23.742190Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, null),
('efe1a571-a8c6-4101-bd58-ed596f7fea0d', False, 'salamonie-pool', 'Salamonie-Pool', null, ST_GeomFromText('POINT(-85.6772 40.8072)',4326), null, null, '2021-03-03T14:02:23.742518Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324450'),
('511e0e84-6c04-4280-b5c1-e2039faacb96', False, 'peru', 'Peru', null, ST_GeomFromText('POINT(-86.0667 40.75)',4326), null, null, '2021-03-03T14:02:23.742607Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03327500'),
('df019cda-8877-4357-b0fd-a3561bc886e2', False, 'logansport', 'Logansport', null, ST_GeomFromText('POINT(-86.3775 40.7464)',4326), null, null, '2021-03-03T14:02:23.742741Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03329000'),
('fc21da72-2171-45a3-832b-17074b81b8f4', False, 'lafayette', 'Lafayette', null, ST_GeomFromText('POINT(-86.8969 40.4219)',4326), null, null, '2021-03-03T14:02:23.742815Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03335500'),
('1521a313-cf26-400e-b58d-676fbe1bc033', False, 'littleriver', 'LittleRiver', null, ST_GeomFromText('POINT(-85.4132 40.8986)',4326), null, null, '2021-03-03T14:02:23.742888Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03324000'),
('d235816d-4546-4261-9db8-e1d0568bf6b7', False, 'covington', 'Covington', null, ST_GeomFromText('POINT(-87.4067 40.14)',4326), null, null, '2021-03-03T14:02:23.742988Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03336000'),
('f9db6d6b-2cea-4bd7-b374-623cb66ff657', False, 'montezuma', 'Montezuma', null, ST_GeomFromText('POINT(-87.3739 39.7925)',4326), null, null, '2021-03-03T14:02:23.743090Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '1fa28968-d0b1-4832-b4b3-e58206056819', '00000000-0000-0000-0000-000000000000', null, '03340500');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('848da38f-ff00-4303-a449-6eaa1c0f734d', 'f612a243-1673-4451-899d-5fd44c4a4f77', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.737872Z'),
('09a2a9b5-e7c4-4231-ad1f-780478b2ee6e', '38c604e4-dd58-4ac1-9e4d-d456318b1705', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.738087Z'),
('c76ef327-7b0a-4de7-a3c8-5dcd3b74dfc5', '14e312f7-bb7e-4f88-afb1-8da275148025', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.738226Z'),
('44b3627d-dab3-4ba3-8b46-4629c78d41dd', 'edb24776-da77-4962-98ed-15515ad8b7ed', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.738402Z'),
('f02e0af1-eece-41d7-b6ab-2e1e6ca4dabd', '67ec7930-57ba-4d08-9280-d428e995cd59', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.738785Z'),
('92fd8f3f-197c-4986-a29f-3837bec5767e', '52178424-b9be-46c7-9f60-305ab9d2b27d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.738983Z'),
('b8912a7d-deab-4fe6-8217-5cc656ed08a1', '9a09251e-fff3-4d02-83e9-e7a539185376', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739182Z'),
('537654e7-c167-455e-b6c5-c7eb53af5c0c', '6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739549Z'),
('ef6ccef6-e75f-4ec4-adff-7cba8a2a79f0', '800fc3e6-a171-4e99-955c-5e0d264ffd29', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739722Z'),
('2530ee2a-2122-4caa-a14d-10bace1b70c1', '38ef8580-6278-4942-bb35-b1d3ac24412c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739804Z'),
('462ef251-9a5e-49ae-ad65-ce3916b40f9d', '969c94eb-70bd-45c1-b23e-17399faa4605', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739877Z'),
('d34cb6b9-8416-4e8c-8055-5453b0542a4a', '3631051d-ead4-46e8-a09b-42f5cee9ef0d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.739985Z'),
('bb16abd7-fbdc-403a-88e7-7cf62ded1cf4', '62ba81af-0197-4a1d-b0f8-cc7d8e5ba1e5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740057Z'),
('43042bf9-18a1-4a84-9a22-34b7866a49d6', 'd5b41092-230b-414e-b2f2-57f7a792c6cc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740127Z'),
('7378b758-515f-43a7-ac63-d7b2cfb442cb', '48fe98a9-38ae-487e-9f26-e0cc37b9699d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740201Z'),
('e3a72707-8507-4cd0-866c-998337023510', '1668b0f0-ce59-4ca3-82cc-e2392a983c41', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740273Z'),
('4bf878ba-40a0-4ddc-b1c6-2abc663920f9', '6cea522c-9e43-43a0-8e0e-02f82d41ae06', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740449Z'),
('5bb4981c-c019-45e4-9879-48cade7fad40', 'bfc6ab1f-801a-4039-81e0-2a3755e0bb23', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740609Z'),
('60b1779c-d0a4-4e6d-96dd-a9acb754b992', 'ce3729fc-a175-49e3-938b-f2bd0b92e61c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740686Z'),
('ebbe6123-7196-40a4-9141-6242c4eca1fa', '3ead0f70-dab6-45f4-9878-bb80b45191ad', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740788Z'),
('18ca0fb4-7cab-4ea2-a3d7-db8f5ea893f1', 'ede9302f-6dbf-4493-97b2-884846f74570', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740919Z'),
('fd78f61e-5b66-40a2-8c2f-40561706cf6f', 'd083e4cb-9269-4c73-8a4f-27bf87f675c9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.740992Z'),
('b6109c2a-d2af-4807-b79e-123b1ca64d46', 'dbeeba6e-b039-4a07-9d2e-a1988698324e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741095Z'),
('6e82d5ba-6d77-4d91-9dbc-70bcc68e0891', '7438be47-82ea-4e31-9a6a-ebe60c308d8f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741168Z'),
('6f44b0d5-1b9a-4fb6-9cb3-e20da99d5f59', 'd37eb7f9-1453-47cd-b8bb-505d32beb899', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741271Z'),
('5e1f8f6b-52b9-41b7-ae69-dbd62252a242', 'f89c9bc9-8f33-4dbe-b225-a58e17176f5f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741344Z'),
('71c4aa04-2439-4cc1-bd9e-421f5ec4327c', '72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741418Z'),
('7c0d7e28-158e-48f2-b4c4-d4a0e40dcadc', '0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741546Z'),
('7199fd7c-dc45-4d6d-a3e7-2e597e7170c9', '254210ea-46fd-4bd6-bf8e-9886b628bde0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741674Z'),
('b1b3d980-3507-4cf3-9365-62a0af931ebc', '22c0eb9c-ccf5-4267-9280-3e5f36497a60', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741806Z'),
('e8c0765d-a604-42fc-a3d7-39df6f14528a', '4fa4f306-95b4-4ec2-b4b1-4f3f87576804', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.741933Z'),
('214f4ec1-d3a6-482e-9030-4e501056e871', '1be8618a-27cc-4f01-9028-779daa13ec2c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742060Z'),
('0e99adfe-a032-4ffa-9eeb-3690b4942e01', '4b45639e-54b7-4cd7-aff5-8ad04fe244cf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742190Z'),
('8db06880-9ef1-4113-aa3a-404d0159074b', 'efe1a571-a8c6-4101-bd58-ed596f7fea0d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742518Z'),
('caf6a829-1156-4233-978c-56e626d7b132', '511e0e84-6c04-4280-b5c1-e2039faacb96', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742607Z'),
('2f9c4bbd-579f-491c-bd64-5e7ee5711a28', 'df019cda-8877-4357-b0fd-a3561bc886e2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742741Z'),
('a4f87b4a-0090-461f-8746-fd20f3f35796', 'fc21da72-2171-45a3-832b-17074b81b8f4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742815Z'),
('9076a5ae-4109-4665-924c-6e9b26b035f5', '1521a313-cf26-400e-b58d-676fbe1bc033', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742888Z'),
('af806175-10be-4adc-ba4c-ddb1b469e7ac', 'd235816d-4546-4261-9db8-e1d0568bf6b7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.742988Z'),
('f2bee67c-61c6-481f-bb88-51591c57789b', 'f9db6d6b-2cea-4bd7-b374-623cb66ff657', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-03T14:02:23.743090Z');

--INSERT TELEMETRY_GOES--COUNT:40
INSERT INTO public.telemetry_goes (id, nesdis_id) select '47b3ff98-b099-4d32-bd02-7600492ce672', 'DD46A45A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD46A45A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a2028e88-de82-48e3-b3da-289530fb539b', 'CE6B687C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B687C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '590b2173-a413-446a-93b8-c1e56ad6490d', 'CE72E178' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72E178');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '3cc9fc5d-7f9f-45a2-9b7a-4dc7b1083163', 'CE2A970C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE2A970C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5925ee9e-48e7-4f23-9fb1-d482f8e8dfd2', 'CE72DA30' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72DA30');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd71ca25f-283e-4757-bbd9-6ed89871ba08', 'CE72D4E2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72D4E2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6002de5c-8440-4ead-af9c-6e65aa00d684', 'CE720C58' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE720C58');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '22be7e6a-1ba3-48d2-b1cf-1934c18f2c6a', 'CE7723A6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7723A6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9bc0821f-0e44-455d-b8b8-6eae7a507930', 'CE6B5DE6' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B5DE6');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2441acda-89b8-4a0b-99a0-fc40dfc2c3ca', 'CE108358' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE108358');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'eb8b2bbd-5828-4de2-9be4-30cdc025cb0d', '17D2F560' where not exists (select 1 from telemetry_goes where nesdis_id = '17D2F560');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9af928eb-5cde-4cc7-a4e7-8955fca75d47', 'D11EE2FE' where not exists (select 1 from telemetry_goes where nesdis_id = 'D11EE2FE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b9aace04-d5c1-4522-85f0-0ac99e1a3e4e', '173623B8' where not exists (select 1 from telemetry_goes where nesdis_id = '173623B8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7cc7da96-e7c6-4de6-84fa-c47b52ea7723', '167526C8' where not exists (select 1 from telemetry_goes where nesdis_id = '167526C8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '13653d3f-ca1e-4cd9-b370-1c21a9fc81eb', 'DE4B8336' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE4B8336');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '15a36268-933f-44ba-85d7-246e508efa5f', 'CE7718EE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE7718EE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '139f1a29-6c65-4fe1-8145-a928b107b7c2', 'CE6B66AE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B66AE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '566d9138-31ff-4b7b-8ed9-1a4067105e7c', 'CE10C052' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE10C052');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'bac58d7d-d0c0-4cac-bed5-4f5f7f8327b6', 'CE77163C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE77163C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '38cc6559-8184-46e4-97c5-55d8f6866519', 'CE16B60C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16B60C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f175520f-b6a5-4d46-9943-e7236be3363e', 'CE6B0D9A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6B0D9A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'cffa8cf5-41d9-42d9-835e-a00cfc0b46a1', 'DDA6F1AC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDA6F1AC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'dc4d176c-6fba-4cb4-9206-e05877c64c3a', '17853608' where not exists (select 1 from telemetry_goes where nesdis_id = '17853608');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b51cdb72-174c-4589-8eb1-b83282c0b542', 'DD8F700A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8F700A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a80e3c92-4616-4182-a6a1-50f84d213758', '163D923C' where not exists (select 1 from telemetry_goes where nesdis_id = '163D923C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd3d718f0-25e0-4ed0-b2cb-b570ef3c9090', 'DD3E3032' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD3E3032');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5430f023-ed8f-4bad-95a2-2de0289902d8', 'CE72FCDC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72FCDC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9e3c97bd-ed48-40c1-944f-8c730dc0ed01', 'CE58F2B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE58F2B2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8fac147e-b3ec-445b-803e-dfffbf37f21b', 'CE72EFAA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72EFAA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1d4145b2-0f5c-4435-867b-6ee403021156', 'CE72F20E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72F20E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'dbc45de0-5a54-4dba-88b1-d1cc779063fb', 'CE730070' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE730070');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd9b96750-04c1-454b-a448-67ad259da8d1', 'CE26C6EC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE26C6EC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8976ffb3-cd7b-4d5e-ba43-914a02687ff1', 'CE72C946' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE72C946');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '673b19c2-e6a8-4ddc-a970-9379cfc56cb9', 'CE6D2BB8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D2BB8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6238ca98-7b3f-4da1-be00-724b03b7ddfa', 'CE777D08' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE777D08');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'cdc53cef-414d-464c-9e19-eb0fae666649', 'CE6D256A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D256A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f9bf7b6e-d550-4f9d-b481-b4a0460a2712', 'CE6D361C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE6D361C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '69b9498b-05ba-4587-abd0-35889ae50a59', 'CE14B3F8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14B3F8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'debf45c1-07d5-4318-9a18-5f242db69d43', 'CE16C09C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE16C09C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd2c2f4c4-e153-46bd-b15f-fd1e00df9d6a', 'CE14E384' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE14E384');

--INSERT INSTRUMENT_TELEMETRY--COUNT:40
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('f612a243-1673-4451-899d-5fd44c4a4f77', '10a32652-af43-4451-bd52-4980c5690cc9', '47b3ff98-b099-4d32-bd02-7600492ce672'),
('38c604e4-dd58-4ac1-9e4d-d456318b1705', '10a32652-af43-4451-bd52-4980c5690cc9', 'a2028e88-de82-48e3-b3da-289530fb539b'),
('14e312f7-bb7e-4f88-afb1-8da275148025', '10a32652-af43-4451-bd52-4980c5690cc9', '590b2173-a413-446a-93b8-c1e56ad6490d'),
('edb24776-da77-4962-98ed-15515ad8b7ed', '10a32652-af43-4451-bd52-4980c5690cc9', '3cc9fc5d-7f9f-45a2-9b7a-4dc7b1083163'),
('67ec7930-57ba-4d08-9280-d428e995cd59', '10a32652-af43-4451-bd52-4980c5690cc9', '5925ee9e-48e7-4f23-9fb1-d482f8e8dfd2'),
('52178424-b9be-46c7-9f60-305ab9d2b27d', '10a32652-af43-4451-bd52-4980c5690cc9', 'd71ca25f-283e-4757-bbd9-6ed89871ba08'),
('9a09251e-fff3-4d02-83e9-e7a539185376', '10a32652-af43-4451-bd52-4980c5690cc9', '6002de5c-8440-4ead-af9c-6e65aa00d684'),
('6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', '10a32652-af43-4451-bd52-4980c5690cc9', '22be7e6a-1ba3-48d2-b1cf-1934c18f2c6a'),
('800fc3e6-a171-4e99-955c-5e0d264ffd29', '10a32652-af43-4451-bd52-4980c5690cc9', '9bc0821f-0e44-455d-b8b8-6eae7a507930'),
('38ef8580-6278-4942-bb35-b1d3ac24412c', '10a32652-af43-4451-bd52-4980c5690cc9', '2441acda-89b8-4a0b-99a0-fc40dfc2c3ca'),
('969c94eb-70bd-45c1-b23e-17399faa4605', '10a32652-af43-4451-bd52-4980c5690cc9', 'eb8b2bbd-5828-4de2-9be4-30cdc025cb0d'),
('3631051d-ead4-46e8-a09b-42f5cee9ef0d', '10a32652-af43-4451-bd52-4980c5690cc9', '9af928eb-5cde-4cc7-a4e7-8955fca75d47'),
('62ba81af-0197-4a1d-b0f8-cc7d8e5ba1e5', '10a32652-af43-4451-bd52-4980c5690cc9', 'b9aace04-d5c1-4522-85f0-0ac99e1a3e4e'),
('d5b41092-230b-414e-b2f2-57f7a792c6cc', '10a32652-af43-4451-bd52-4980c5690cc9', '7cc7da96-e7c6-4de6-84fa-c47b52ea7723'),
('48fe98a9-38ae-487e-9f26-e0cc37b9699d', '10a32652-af43-4451-bd52-4980c5690cc9', '13653d3f-ca1e-4cd9-b370-1c21a9fc81eb'),
('1668b0f0-ce59-4ca3-82cc-e2392a983c41', '10a32652-af43-4451-bd52-4980c5690cc9', '15a36268-933f-44ba-85d7-246e508efa5f'),
('6cea522c-9e43-43a0-8e0e-02f82d41ae06', '10a32652-af43-4451-bd52-4980c5690cc9', '139f1a29-6c65-4fe1-8145-a928b107b7c2'),
('bfc6ab1f-801a-4039-81e0-2a3755e0bb23', '10a32652-af43-4451-bd52-4980c5690cc9', '566d9138-31ff-4b7b-8ed9-1a4067105e7c'),
('ce3729fc-a175-49e3-938b-f2bd0b92e61c', '10a32652-af43-4451-bd52-4980c5690cc9', 'bac58d7d-d0c0-4cac-bed5-4f5f7f8327b6'),
('3ead0f70-dab6-45f4-9878-bb80b45191ad', '10a32652-af43-4451-bd52-4980c5690cc9', '38cc6559-8184-46e4-97c5-55d8f6866519'),
('ede9302f-6dbf-4493-97b2-884846f74570', '10a32652-af43-4451-bd52-4980c5690cc9', 'f175520f-b6a5-4d46-9943-e7236be3363e'),
('d083e4cb-9269-4c73-8a4f-27bf87f675c9', '10a32652-af43-4451-bd52-4980c5690cc9', 'cffa8cf5-41d9-42d9-835e-a00cfc0b46a1'),
('dbeeba6e-b039-4a07-9d2e-a1988698324e', '10a32652-af43-4451-bd52-4980c5690cc9', 'dc4d176c-6fba-4cb4-9206-e05877c64c3a'),
('7438be47-82ea-4e31-9a6a-ebe60c308d8f', '10a32652-af43-4451-bd52-4980c5690cc9', 'b51cdb72-174c-4589-8eb1-b83282c0b542'),
('d37eb7f9-1453-47cd-b8bb-505d32beb899', '10a32652-af43-4451-bd52-4980c5690cc9', 'a80e3c92-4616-4182-a6a1-50f84d213758'),
('f89c9bc9-8f33-4dbe-b225-a58e17176f5f', '10a32652-af43-4451-bd52-4980c5690cc9', 'd3d718f0-25e0-4ed0-b2cb-b570ef3c9090'),
('72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', '10a32652-af43-4451-bd52-4980c5690cc9', '5430f023-ed8f-4bad-95a2-2de0289902d8'),
('0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', '10a32652-af43-4451-bd52-4980c5690cc9', '9e3c97bd-ed48-40c1-944f-8c730dc0ed01'),
('254210ea-46fd-4bd6-bf8e-9886b628bde0', '10a32652-af43-4451-bd52-4980c5690cc9', '8fac147e-b3ec-445b-803e-dfffbf37f21b'),
('22c0eb9c-ccf5-4267-9280-3e5f36497a60', '10a32652-af43-4451-bd52-4980c5690cc9', '1d4145b2-0f5c-4435-867b-6ee403021156'),
('4fa4f306-95b4-4ec2-b4b1-4f3f87576804', '10a32652-af43-4451-bd52-4980c5690cc9', 'dbc45de0-5a54-4dba-88b1-d1cc779063fb'),
('1be8618a-27cc-4f01-9028-779daa13ec2c', '10a32652-af43-4451-bd52-4980c5690cc9', 'd9b96750-04c1-454b-a448-67ad259da8d1'),
('4b45639e-54b7-4cd7-aff5-8ad04fe244cf', '10a32652-af43-4451-bd52-4980c5690cc9', '8976ffb3-cd7b-4d5e-ba43-914a02687ff1'),
('efe1a571-a8c6-4101-bd58-ed596f7fea0d', '10a32652-af43-4451-bd52-4980c5690cc9', '673b19c2-e6a8-4ddc-a970-9379cfc56cb9'),
('511e0e84-6c04-4280-b5c1-e2039faacb96', '10a32652-af43-4451-bd52-4980c5690cc9', '6238ca98-7b3f-4da1-be00-724b03b7ddfa'),
('df019cda-8877-4357-b0fd-a3561bc886e2', '10a32652-af43-4451-bd52-4980c5690cc9', 'cdc53cef-414d-464c-9e19-eb0fae666649'),
('fc21da72-2171-45a3-832b-17074b81b8f4', '10a32652-af43-4451-bd52-4980c5690cc9', 'f9bf7b6e-d550-4f9d-b481-b4a0460a2712'),
('1521a313-cf26-400e-b58d-676fbe1bc033', '10a32652-af43-4451-bd52-4980c5690cc9', '69b9498b-05ba-4587-abd0-35889ae50a59'),
('d235816d-4546-4261-9db8-e1d0568bf6b7', '10a32652-af43-4451-bd52-4980c5690cc9', 'debf45c1-07d5-4318-9a18-5f242db69d43'),
('f9db6d6b-2cea-4bd7-b374-623cb66ff657', '10a32652-af43-4451-bd52-4980c5690cc9', 'd2c2f4c4-e153-46bd-b15f-fd1e00df9d6a');

--INSERT TIMESERIES--COUNT:40
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('7995d7ae-44f5-4b5a-9e5b-9a06e3c1dca5','stage','Stage','f612a243-1673-4451-899d-5fd44c4a4f77', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c543fe70-530d-42a8-b5ea-1f851b6147b2','stage','Stage','38c604e4-dd58-4ac1-9e4d-d456318b1705', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c61778ed-b0f9-465f-82dd-da949c848d69','precipitation','Precipitation','38c604e4-dd58-4ac1-9e4d-d456318b1705', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9ad42ce6-ddb7-44ce-a4ec-06d554c4e31a','stage','Stage','14e312f7-bb7e-4f88-afb1-8da275148025', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9abf2d0b-6ce0-4f4d-8f2f-f7afcaca0e0a','precipitation','Precipitation','14e312f7-bb7e-4f88-afb1-8da275148025', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('cd2b55a7-247c-4935-ad9c-41e3d4bd8b7a','precipitation','Precipitation','edb24776-da77-4962-98ed-15515ad8b7ed', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('8b7f4978-61d6-47db-8f10-af4220df40a5','elevation','Elevation','67ec7930-57ba-4d08-9280-d428e995cd59', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ed7e153a-be3b-4197-82d2-f86c2ed74c57','precipitation','Precipitation','67ec7930-57ba-4d08-9280-d428e995cd59', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('5dd3e53f-5dc3-455f-a408-94f94565ab53','voltage','Voltage','67ec7930-57ba-4d08-9280-d428e995cd59', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3d75dd6c-ec07-4225-b17d-63d569e306b7','elevation','Elevation','52178424-b9be-46c7-9f60-305ab9d2b27d', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d70b2894-5f6b-4a26-867f-b3cd3439db69','precipitation','Precipitation','52178424-b9be-46c7-9f60-305ab9d2b27d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('cf1108f6-2289-426c-a015-5af47e5ecf0c','voltage','Voltage','52178424-b9be-46c7-9f60-305ab9d2b27d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('111f31de-896f-4f11-b8f1-a8086878e512','elevation','Elevation','9a09251e-fff3-4d02-83e9-e7a539185376', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('944154d5-5cf0-4bd5-a2c0-b95e7fe628cc','precipitation','Precipitation','9a09251e-fff3-4d02-83e9-e7a539185376', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c0ab8f27-bc64-4d34-974b-3fa7b135fd23','voltage','Voltage','9a09251e-fff3-4d02-83e9-e7a539185376', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('63a39657-ead1-46c1-8e09-22d89832c337','stage','Stage','6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bdeacb12-a664-46f4-a9f7-22468105833a','water-temperature','Water-Temperature','6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('314e1c4e-cda7-42b9-9d69-4f46dd631300','voltage','Voltage','6cf5ad41-8a7f-4b6e-8b89-9468b8d724e1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('17dd1144-1f2f-4dfc-881e-fac1747bd420','stage','Stage','800fc3e6-a171-4e99-955c-5e0d264ffd29', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('045f43f0-b653-475c-96a1-09e77a21af55','elevation','Elevation','38ef8580-6278-4942-bb35-b1d3ac24412c', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2600c26e-01c9-40c8-ada4-9b4a7331b9a6','stage','Stage','969c94eb-70bd-45c1-b23e-17399faa4605', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b9fd3f13-55c4-47ea-b5ba-4d2dd1fab6b3','precipitation','Precipitation','969c94eb-70bd-45c1-b23e-17399faa4605', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('7f8dbc9c-73c9-473c-add5-7f21973eb65a','stage','Stage','3631051d-ead4-46e8-a09b-42f5cee9ef0d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1b475140-66ec-4f93-8c02-58e63cf80479','stage','Stage','62ba81af-0197-4a1d-b0f8-cc7d8e5ba1e5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b2eeb995-242b-41fa-942d-eaaecbbebbed','stage','Stage','d5b41092-230b-414e-b2f2-57f7a792c6cc', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3be2cc4f-79b4-443f-b976-7102a3aba4ba','stage','Stage','48fe98a9-38ae-487e-9f26-e0cc37b9699d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('08c7bac3-3654-4f3d-8ca4-dad4fee0cd9c','stage','Stage','1668b0f0-ce59-4ca3-82cc-e2392a983c41', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ad4d542f-25c6-4d65-81ab-2355c5eddce3','water-temperature','Water-Temperature','1668b0f0-ce59-4ca3-82cc-e2392a983c41', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('273f7796-de37-4136-8eae-21f24de0f8f2','precipitation','Precipitation','1668b0f0-ce59-4ca3-82cc-e2392a983c41', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('f9437c5b-bf85-4c6a-bda4-e3396a966a93','stage','Stage','6cea522c-9e43-43a0-8e0e-02f82d41ae06', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('977fc8c5-5eb5-49cd-9e40-836ddcb3e4de','precipitation','Precipitation','6cea522c-9e43-43a0-8e0e-02f82d41ae06', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a14128bf-cda0-4b19-8a26-87b4a05e4548','voltage','Voltage','6cea522c-9e43-43a0-8e0e-02f82d41ae06', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8186421e-849d-4dfa-a26f-52e0ab00bca0','elevation','Elevation','bfc6ab1f-801a-4039-81e0-2a3755e0bb23', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('079d2efa-ef2e-4f88-b7ab-7495addf7e41','stage','Stage','ce3729fc-a175-49e3-938b-f2bd0b92e61c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5130ccb1-35ee-432b-9add-68ddf6e08169','water-temperature','Water-Temperature','ce3729fc-a175-49e3-938b-f2bd0b92e61c', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc'),
('c3209723-61a6-4329-8bfb-a9405252e0c6','stage','Stage','3ead0f70-dab6-45f4-9878-bb80b45191ad', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a4f252d6-fa9a-4cc4-97aa-e5d380eecaca','precipitation','Precipitation','3ead0f70-dab6-45f4-9878-bb80b45191ad', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d7ab1307-852c-47af-aa05-26fcf37dec9b','voltage','Voltage','3ead0f70-dab6-45f4-9878-bb80b45191ad', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c6311836-0d11-4469-984e-192193bf2935','stage','Stage','ede9302f-6dbf-4493-97b2-884846f74570', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('63cc4d20-bb9b-40ce-b9b4-73560c2e9590','stage','Stage','d083e4cb-9269-4c73-8a4f-27bf87f675c9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('df6155b1-a891-4603-ba7b-cd234b5f07d8','precipitation','Precipitation','d083e4cb-9269-4c73-8a4f-27bf87f675c9', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('68319924-85fa-442b-a797-3ad57742ce1c','stage','Stage','dbeeba6e-b039-4a07-9d2e-a1988698324e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ad6cac4e-8684-4344-87d1-0935cf894ff0','stage','Stage','7438be47-82ea-4e31-9a6a-ebe60c308d8f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cc710254-1f64-47ad-b25e-ed87e3a29e2d','precipitation','Precipitation','7438be47-82ea-4e31-9a6a-ebe60c308d8f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9139fa91-489d-4235-9b59-6dee17be1169','stage','Stage','d37eb7f9-1453-47cd-b8bb-505d32beb899', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f5950f94-a490-42bd-9e8b-e41101b57f64','stage','Stage','f89c9bc9-8f33-4dbe-b225-a58e17176f5f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8dddf703-4066-411c-ad97-9de510033a03','stage','Stage','72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ef044d7f-e0b9-4b3b-8a56-19f1e739a433','precipitation','Precipitation','72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e98cd4d6-cfca-4d0c-9512-0f0f8150dfad','voltage','Voltage','72b78850-7c4b-43f2-ab4e-0d6d4f4b6dc4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('cade9107-1ed2-4bff-93a9-93e7aeda40fd','elevation','Elevation','0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e195b56d-ac60-49c9-85c5-543e079a60c5','precipitation','Precipitation','0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a04128e1-7832-414c-b720-c1f4a2a3d297','voltage','Voltage','0bf054e8-28b2-44d1-82a5-cdfea7bd24fe', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('81eb3fa6-7ee2-4849-a939-e1b2545b2dd5','stage','Stage','254210ea-46fd-4bd6-bf8e-9886b628bde0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6d1a6e1a-dc35-42cd-ae5a-0084cc896294','precipitation','Precipitation','254210ea-46fd-4bd6-bf8e-9886b628bde0', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('1643f2ff-3c28-4471-8a03-dc7a583b316f','voltage','Voltage','254210ea-46fd-4bd6-bf8e-9886b628bde0', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('51c47d1d-a001-42f3-8a84-7ff38cd5ca3e','elevation','Elevation','22c0eb9c-ccf5-4267-9280-3e5f36497a60', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d1389b2e-7936-44b7-b41c-89f0a1eaa20c','precipitation','Precipitation','22c0eb9c-ccf5-4267-9280-3e5f36497a60', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('610f21bc-9e44-4f96-b30f-81e636ea8442','voltage','Voltage','22c0eb9c-ccf5-4267-9280-3e5f36497a60', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2a121abf-0781-4825-b720-e072f26aaa92','stage','Stage','4fa4f306-95b4-4ec2-b4b1-4f3f87576804', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('67220d25-5ac1-4f74-bf60-d61c8906cbf4','precipitation','Precipitation','4fa4f306-95b4-4ec2-b4b1-4f3f87576804', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('3cdaaad3-66d6-4e4e-93d0-4d99445f6c43','voltage','Voltage','4fa4f306-95b4-4ec2-b4b1-4f3f87576804', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('df373656-cdc0-436c-b333-eb9b86566242','elevation','Elevation','1be8618a-27cc-4f01-9028-779daa13ec2c', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('caa2fbfe-ebdf-4ff2-a4c5-85dc9b27410f','precipitation','Precipitation','1be8618a-27cc-4f01-9028-779daa13ec2c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ab3acad4-9951-4218-a151-60683cd295e7','voltage','Voltage','1be8618a-27cc-4f01-9028-779daa13ec2c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f4b80745-d2b0-4883-84bf-67bdfefc48a7','voltage','Voltage','4b45639e-54b7-4cd7-aff5-8ad04fe244cf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8eececc3-b6a9-4633-823a-0c89b68426e6','elevation','Elevation','efe1a571-a8c6-4101-bd58-ed596f7fea0d', '83b5a1f7-948b-4373-a47c-d73ff622aafd', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0f2d1335-3f64-4332-8a42-b58d053bd368','stage','Stage','511e0e84-6c04-4280-b5c1-e2039faacb96', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3e3a1e86-f0aa-4311-bbc1-c5e08a7c1680','precipitation','Precipitation','511e0e84-6c04-4280-b5c1-e2039faacb96', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('378064a2-972b-4bdd-9c14-6541a8268996','voltage','Voltage','511e0e84-6c04-4280-b5c1-e2039faacb96', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('453256e0-690b-45af-9b60-e407b1d20fc6','stage','Stage','df019cda-8877-4357-b0fd-a3561bc886e2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('620f2a85-f2c4-4724-9b2f-d643c9487402','stage','Stage','fc21da72-2171-45a3-832b-17074b81b8f4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('7fb22a4b-1c30-44e7-bdfc-457681cc544a','stage','Stage','1521a313-cf26-400e-b58d-676fbe1bc033', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('126d7486-f429-4c8e-a43d-2222f4b71193','precipitation','Precipitation','1521a313-cf26-400e-b58d-676fbe1bc033', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6e8a2389-2f70-40e4-b64e-760ed3440fa4','stage','Stage','d235816d-4546-4261-9db8-e1d0568bf6b7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d9ab6370-ad17-4e22-b480-11d66e87151b','precipitation','Precipitation','d235816d-4546-4261-9db8-e1d0568bf6b7', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('01d7079b-fee8-474d-8c2d-52934fb078f6','stage','Stage','f9db6d6b-2cea-4bd7-b374-623cb66ff657', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2203c29d-c2bf-479a-ad00-e4a24758ffb0','precipitation','Precipitation','f9db6d6b-2cea-4bd7-b374-623cb66ff657', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a');
