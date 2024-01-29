-- -------------------------------------------------
-- basic seed data to demo the app and run API tests
-- -------------------------------------------------

INSERT INTO heartbeat (time) VALUES (NOW());

-- Profile (Faked with: https://homepage.net/name_generator/)
-- NOTE: EDIPI 1 should not be used; test user with EDIPI = 1 created by integration tests
INSERT INTO profile (id, edipi, is_admin, username, email) VALUES
    -- Application Admin
    ('57329df6-9f7a-4dad-9383-4633b452efab',2,true,'AnthonyLambert','anthony.lambert@fake.usace.army.mil'),
    -- Blue Water Dam Project Admin
    ('f320df83-e2ea-4fe9-969a-4e0239b8da51',3,false,'MollyRutherford','molly.rutherford@fake.usace.army.mil'),
    -- Blue Water Dam Project Member
    ('89aa1e13-041a-4d15-9e45-f76eba3b0551',4,false,'DominicGlover','dominic.glover@fake.usace.army.mil'),
    ('405ab7e1-20fc-4d26-a074-eccad88bf0a9',5,false,'JoeQuinn','joe.quinn@fake.usace.army.mil'),
    ('81c77210-6244-46fe-bdf6-35da4f00934b',6,false,'TrevorDavidson','trevor.davidson@fake.usace.army.mil'),
    ('f056201a-ffec-4f5b-aec5-14b34bb5e3d8',7,false,'ClaireButler','claire.butler@fake.usace.army.mil'),
    ('9effda27-49f7-4745-8e55-fa819f550b09',8,false,'SophieBower','sophie.bower@fake.usace.army.mil'),
    ('37407aba-904a-42fa-af73-6ab748ee1f98',9,false,'NeilMcLean','neil.mclean@fake.usace.army.mil'),
    ('c0fd72ae-cccc-45c9-ba1d-4353170c352d',10,false,'JakeBurgess','jake.burgess@fake.usace.army.mil'),
    ('be549c16-3f65-4af4-afb6-e18c814c44dc',11,false,'DanQuinn','dan.quinn@fake.usace.army.mil');

-- project
INSERT INTO project (id, slug, name, image) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'blue-water-dam-example-project', 'Blue Water Dam Example Project', 'site_photo.jpg');

-- profile_project_role
INSERT INTO profile_project_roles (profile_id, role_id, project_id) VALUES
    -- Blue Water Dam Project Admin
    ('f320df83-e2ea-4fe9-969a-4e0239b8da51', '37f14863-8f3b-44ca-8deb-4b74ce8a8a69', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984'),
    -- Blue Water dam Project Member
    ('89aa1e13-041a-4d15-9e45-f76eba3b0551', '2962bdde-7007-4ba0-943f-cb8e72e90704', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984');

-- instrument_group
INSERT INTO instrument_group (id, project_id, slug, name, description) VALUES
    ('d0916e8a-39a6-4f2f-bd31-879881f8b40c', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'sample-instrument-group', 'Sample Instrument Group 1', 'This is an example instrument group');

-- instrument
INSERT INTO instrument (id, project_id, slug, name, geometry, type_id) VALUES
    ('a7540f69-c41e-43b3-b655-6e44097edb7e', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'demo-piezometer-1', 'Demo Piezometer 1', ST_GeomFromText('POINT(-80.8 26.7)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
    ('9e8f2ca4-4037-45a4-aaca-d9e598877439', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'demo-staffgage-1', 'Demo Staffgage 1', ST_GeomFromText('POINT(-80.85 26.75)',4326),'0fd1f9ba-2731-4ff9-96dd-3c03215ab06f'),
    ('d8c66ef9-06f0-4d52-9233-f3778e0624f0', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'inclinometer-1', 'inclinometer-1', ST_GeomFromText('POINT(-80.8 26.7)',4326),'98a61f29-18a8-430a-9d02-0f53486e0984');

-- instrument_group_instruments
INSERT INTO instrument_group_instruments (instrument_id, instrument_group_id) VALUES
    ('a7540f69-c41e-43b3-b655-6e44097edb7e', 'd0916e8a-39a6-4f2f-bd31-879881f8b40c');

-- instrument_status
-- (1) Active    in 1980 (sample, project construction)
-- (2) Destroyed in 2000
-- (3) Abandoned in 2001
INSERT INTO instrument_status (id, instrument_id, time, status_id) VALUES
    ('52ad0ce9-1034-448c-a5a9-8f6e9676ed1b', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '1980-01-01','e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('2fa6a965-73aa-463c-ac10-6c83a2a34f60', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2000-05-01','94578354-ffdf-4119-9663-6bd4323e58f5'),
    ('4ed5e9ac-40dc-4bca-b44f-7b837ec1b0fc', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2001-01-01', '3d9add10-4418-4cb2-bd31-20a639504539'),
    ('f98cf0c4-347b-4709-9c9a-4fde4c5726be', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '2020-01-01', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');

-- instrument_notes
INSERT INTO instrument_note (id, instrument_id, title, body) VALUES
('90a3f8de-de65-48a7-8286-024c13162958', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'Instrument Test Note 1',
'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
 labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
 nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
 cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
 officia deserunt mollit anim id est laborum.
'),
('d7a2bc43-551a-4ee4-8dd4-dc7e21079f43', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'Instrument Test Note 2',
'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
 labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
 nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
 cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
 officia deserunt mollit anim id est laborum.
'),
('29eacfc0-090d-4ed4-8dac-e492c76c305f', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'Instrument Test Note 3',
'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
 labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
 nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
 cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
 officia deserunt mollit anim id est laborum.
');

-- Time Series
INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
('869465fc-dc1e-445e-81f4-9979b5fadda9', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '1de79e29-fb70-45c3-ae7d-4695517ced90', '6407a23f-b5f8-4214-9343-50b6231e4bfe', 'atmospheric-pressure', 'Atmospheric Pressure'),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'distance-to-water', 'Distance to Water'),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'height', 'Height'),
('d9697351-3a38-4194-9ac4-41541927e475', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'top-of-riser', 'Top of Riser'),
('22a734d6-dc24-451d-a462-43a32f335ae8', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'tip-depth', 'Tip Depth'),
('14247bc8-b264-4857-836f-182d47ebb39d', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'constant-to-test-delete', 'Constant to Test Delete'),
('5985f20a-1e37-4add-823c-545cdca49b5e', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'inclinometer-1', 'Inclinometer-1'),
('479d90eb-3454-4f39-be9a-bfd23099a552', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '3ea5ed77-c926-4696-a580-a3fde0f9a556', 'ae06a7db-1e18-4994-be41-9d5a408d6cad', 'inclinometer-constant', 'inclinometer-constant'),
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'demo-piezometer-1.formula', 'demo-piezometer-1'),
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5985', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '2b7f96e1-820f-4f61-ba8f-861640af6232', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'demo-staffgage-1.formula', 'demo-staffgage-1'),
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5986', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'inclinometer-1.formula', 'inclinometer-1'),
('844fb688-e77c-481e-bff9-81a0fff9f3f2', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'test-create-datalogger-table-mapping', 'test-create-datalogger-table-mapping'),
('da79bdb9-ded4-4f4a-8982-33e09b136815', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'test-table-mapping', 'test-table-mapping');

INSERT INTO calculation (timeseries_id, contents) VALUES
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '[demo-piezometer-1.top-of-riser] - [demo-piezometer-1.distance-to-water]'),
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5985', null),
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5986', null);

-- instrument_constants
INSERT INTO instrument_constants (instrument_id, timeseries_id) VALUES
('a7540f69-c41e-43b3-b655-6e44097edb7e', 'd9697351-3a38-4194-9ac4-41541927e475'),
('a7540f69-c41e-43b3-b655-6e44097edb7e', '22a734d6-dc24-451d-a462-43a32f335ae8'),
('d8c66ef9-06f0-4d52-9233-f3778e0624f0', '479d90eb-3454-4f39-be9a-bfd23099a552'),
('a7540f69-c41e-43b3-b655-6e44097edb7e', '14247bc8-b264-4857-836f-182d47ebb39d');

-- Time Series Measurements
INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/1/2020' , 13.16),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/2/2020' , 13.16),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/3/2020' , 13.17),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/4/2020' , 13.17),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/5/2020' , 13.13),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/6/2020' , 13.12),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/7/2020' , 13.10),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/8/2020' , 13.08),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/9/2020' , 13.07),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/10/2020', 13.05),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/11/2020', 13.16),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/12/2020', 13.16),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/13/2020', 13.17),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/14/2020', 13.17),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/15/2020', 13.13),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/16/2020', 13.12),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/17/2020', 13.10),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/18/2020', 13.08),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/19/2020', 13.07),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/20/2020', 13.05),
('869465fc-dc1e-445e-81f4-9979b5fadda9', '1/21/2020', 13.05),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/1/2020' , 20.16),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/2/2020' , 20.16),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/3/2020' , 20.17),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/4/2020' , 20.17),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/5/2020' , 20.13),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/6/2020' , 20.12),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/7/2020' , 20.10),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/8/2020' , 20.08),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/9/2020' , 20.07),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', '1/10/2020', 20.05),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/1/2020' , 20.16),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/2/2020' , 20.16),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/3/2020' , 20.17),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/4/2020' , 20.17),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/5/2020' , 20.13),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/6/2020' , 20.12),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/7/2020' , 20.10),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/8/2020' , 20.08),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/9/2020' , 20.07),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', '3/10/2020', 20.05),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/1/2020' , 20.16),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/2/2020' , 20.16),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/3/2020' , 20.17),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/4/2020' , 20.17),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/5/2020' , 20.13),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/6/2020' , 20.12),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/7/2020' , 20.10),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/8/2020' , 20.08),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/9/2020' , 20.07),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/10/2020', 20.05),
('d9697351-3a38-4194-9ac4-41541927e475', '3/10/2015', 40.50),
('d9697351-3a38-4194-9ac4-41541927e475', '6/10/2020', 40.00),
('d9697351-3a38-4194-9ac4-41541927e475', '3/10/2020', 39.50),
('22a734d6-dc24-451d-a462-43a32f335ae8', '3/10/2015', 10.0),
('479d90eb-3454-4f39-be9a-bfd23099a552', '6/21/2021', 20000.0);

-- inclinometers
INSERT INTO inclinometer_measurement (timeseries_id, time, creator, create_date, values) VALUES 
('5985f20a-1e37-4add-823c-545cdca49b5e', '6/21/2021', '176704ad-829f-44fa-b71b-c112e80261fa', '6/1/2020', 
    '[
        {"depth": 106, "a0": 590, "a180": -562, "b0": -142, "b180": 176},
        {"depth": 108, "a0": 614, "a180": -586, "b0": 107, "b180": -149},
        {"depth": 110, "a0": 622, "a180": -592, "b0": -67, "b180": 107},
        {"depth": 112, "a0": 623, "a180": -598, "b0": 8, "b180": -48},
        {"depth": 114, "a0": 606, "a180": -577, "b0": 124, "b180": -72},
        {"depth": 116, "a0": 0, "a180": 0, "b0": 0, "b180": 0}
    ]'
);

-- collection_group
INSERT INTO collection_group (id, project_id, name, slug) VALUES
    ('1519eaea-1799-4375-aa37-0e35aa654643', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Manual Collection Route 1', 'manual-collection-route-1'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'High Water Inspection', 'high-water-inspection');

-- collection_group_timeseries
INSERT INTO collection_group_timeseries (collection_group_id, timeseries_id) VALUES
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '7ee902a3-56d0-4acf-8956-67ac82c03a96'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '9a3864a8-8766-4bfa-bad1-0328b166f6a8');

-- plot_configuration
INSERT INTO plot_configuration (id, project_id, slug, name) VALUES
    ('cc28ca81-f125-46c6-a5cd-cc055a003c19', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'all-plots', 'All Plots'),
    ('64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'pz-1a-plot', 'PZ-1A PLOT');


-- plot_configuration_timeseries
INSERT INTO plot_configuration_timeseries (plot_configuration_id, timeseries_id) VALUES
    ('cc28ca81-f125-46c6-a5cd-cc055a003c19', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'),
    ('cc28ca81-f125-46c6-a5cd-cc055a003c19', '9a3864a8-8766-4bfa-bad1-0328b166f6a8'),
    ('64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a', '8f4ca3a3-5971-4597-bd6f-332d1cf5af7c'),
    ('64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a', '9a3864a8-8766-4bfa-bad1-0328b166f6a8');

-- telemetry_type
INSERT INTO telemetry_type (id, slug, name) VALUES
    ('10a32652-af43-4451-bd52-4980c5690cc9', 'goes-self-timed', 'GOES Self Timed'),
    ('c0b03b0d-bfce-453a-b5a9-636118940449', 'iridium', 'Iridium');


-- THE FOLLOWING IS A 100% SAMPLE TELEMETRY CONFIGURATION;
-- THIS REPRESENTS A SINGLE INSTRUMENT WITH GOES AND IRIDIUM DATA TRANSMISSION
INSERT INTO telemetry_goes (id, nesdis_id) VALUES
    ('52fb5fbc-af7d-4a60-9fe3-3d1237091e6d', 'TEST123'),
    ('c6b18827-5841-49dd-a7f8-bfafc681e158', 'TEST456');

INSERT INTO telemetry_iridium (id, imei) VALUES
    ('a5e8df6c-554f-4312-a84a-3876c41b4b1a', '123456789098765'),
    ('1bda5844-1065-4bdb-8f49-d35c7a75b1de', '098765432123456');

INSERT INTO instrument_telemetry (id, instrument_id, telemetry_type_id, telemetry_id) VALUES
    ('8bb7c44f-7c72-4715-8337-457643b1a0d5', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '10a32652-af43-4451-bd52-4980c5690cc9', '52fb5fbc-af7d-4a60-9fe3-3d1237091e6d'),
    ('a7cab13d-f6d2-44ba-8e08-8550ac690427', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'c0b03b0d-bfce-453a-b5a9-636118940449', 'a5e8df6c-554f-4312-a84a-3876c41b4b1a');


-- aware_parameter
INSERT INTO aware_parameter (id, key, parameter_id, unit_id, timeseries_slug, timeseries_name) VALUES
    ('1d9f9d06-6fcb-41dd-9fe4-e513a2575e74', 'depth1', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'stage', 'Stage'),
    ('c5f2842d-a5a9-4f53-9583-f613080a9c36', 'battery', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64', 'battery-voltage', 'Battery Voltage'),
    ('78d32638-5137-481c-aa9d-a48c2d57824a', 'baro', '1de79e29-fb70-45c3-ae7d-4695517ced90', '55dda9ef-7ba6-4432-b64d-8ef0e65154f4', 'barometric-pressure', 'Barometric Pressure'),
    ('53ce89a7-1db8-45bd-baf9-9536d75d7046', 'h2oTemp', 'de6112da-8489-4286-ae56-ec72aa09974d', '6462733b-5b42-46a2-ad44-882a5332eafc', 'water-temperature', 'Water Temperature'),
    ('3ca7c1da-7124-42c0-b92c-f76b5c318b0c', 'rssi', 'b23b141d-ce7b-4e0d-82ab-c8beb39c8325', 'be854f6e-e36e-4bba-9e06-6d5aa09485be', 'signal-strength', 'Signal Strength');


-- aware_platform
-- aware_id used below is not an actual aware_id; generated for testing
INSERT INTO aware_platform (id, aware_id, instrument_id) VALUES
    ('b896ce34-2bd4-436c-9f28-7a1eefb5744a', '6df213c4-a582-4735-a916-6f4065082872', 'a7540f69-c41e-43b3-b655-6e44097edb7e');

-- enable all parameters for sample aware platform
INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b
	ORDER BY aware_platform_id
    
)
ON CONFLICT DO NOTHING;

INSERT INTO project (id, slug, name, image) VALUES
    ('6c56c7b0-2d9f-4a1b-9173-b969942dacb5','c-44','C-44', 'midas-project-default.jpg'),
    ('04d0a9a7-c170-4aba-a819-6f0b1e772a9d','alaska-district','Alaska District', 'midas-project-default.jpg'),
    ('f714ccdd-9695-4284-ab98-ace2daba2545','albuquerque-district','Albuquerque District', 'midas-project-default.jpg'),
    ('d559abfd-7ec7-4d0d-97bd-a04018f01e4c','baltimore-district','Baltimore District', 'midas-project-default.jpg'),
    ('bb2e0cfd-41ac-4c49-9408-a1e423d37995','buffalo-district','Buffalo District', 'midas-project-default.jpg'),
    ('7c45de1c-9ace-47ed-9ff8-ec31d0586dc4','charleston-district','Charleston District', 'midas-project-default.jpg'),
    ('ef444d8e-b5bc-42a3-b40d-f502312858e1','chicago-district','Chicago District', 'midas-project-default.jpg'),
    ('fb0cd440-846e-4f5a-a1bc-ed0866126378','detroit-district','Detroit District', 'midas-project-default.jpg'),
    ('0c31a8ec-cf31-4998-8c7d-79e72c993048','fort-worth-district','Fort Worth District', 'midas-project-default.jpg'),
    ('afbb602f-c0b4-463c-a718-568b263f69af','galveston-district','Galveston District', 'midas-project-default.jpg'),
    ('b09015da-eae4-4f96-ad67-745ac0e3ea5b','huntington-district','Huntington District', 'midas-project-default.jpg'),
    ('e0b13eab-6f52-4acd-b3b4-8332021db679','jacksonville-district','Jacksonville District', 'midas-project-default.jpg'),
    ('78b0538f-8fd5-4484-b38e-0ab6cc2c2104','kansas-city-district','Kansas City District', 'midas-project-default.jpg'),
    ('0356320d-c89b-45f0-9e77-faea8ee39d5c','little-rock-district','Little Rock District', 'midas-project-default.jpg'),
    ('63e96826-4dac-4a7c-9d15-6cddf27965c3','los-angeles-district','Los Angeles District', 'midas-project-default.jpg'),
    ('42869cd4-840f-4494-949e-6b5090ccfd41','louisville-district','Louisville District', 'midas-project-default.jpg'),
    ('8db17b8e-8e5e-45cb-9816-ae3d4b6e8c54','memphis-district','Memphis District', 'midas-project-default.jpg'),
    ('d8eba071-d9ce-4a9c-bfc0-79d55b602c8b','mobile-district','Mobile District', 'midas-project-default.jpg'),
    ('d12c79b2-1adb-44ae-910b-ba29ad1e96d8','nashville-district','Nashville District', 'midas-project-default.jpg'),
    ('e1b18976-12fc-4b6c-b615-3ffa5cde6a31','new-england-district','New England District', 'midas-project-default.jpg'),
    ('e227908a-26a6-46e9-bd45-a9554c46a690','norfolk-district','Norfolk District', 'midas-project-default.jpg'),
    ('b6474c93-c683-4080-8dd8-50b42a1676cd','omaha-district','Omaha District', 'midas-project-default.jpg'),
    ('a2c0f965-b3da-44b3-bde3-2e8b27cac1aa','philadelphia-district','Philadelphia District', 'midas-project-default.jpg'),
    ('7ae67773-1f5e-44ce-a9a9-7fcabf5d478d','pittsburgh-district','Pittsburgh District', 'midas-project-default.jpg'),
    ('92d3a8f8-7c59-44c9-997a-cafaf61ee4ff','portland-district','Portland District', 'midas-project-default.jpg'),
    ('40f931c7-3212-4d18-93da-2b1788eb18a9','rock-island-district','Rock Island District', 'midas-project-default.jpg'),
    ('34c7a4bf-9eb5-44cd-971f-95f6246766b5','sacramento-district','Sacramento District', 'midas-project-default.jpg'),
    ('5e2ec151-840b-43c0-a9e5-aeff3cbf08fa','san-francisco-district','San Francisco District', 'midas-project-default.jpg'),
    ('2103b415-475e-443b-b637-a17ca626b6d2','savannah-district','Savannah District', 'midas-project-default.jpg'),
    ('eb0ee6ca-9689-41a5-9131-5be0a753485e','seattle-district','Seattle District', 'midas-project-default.jpg'),
    ('39ff3886-7279-41b8-b62a-c1d89d4aacf4','st-louis-district','St. Louis District', 'midas-project-default.jpg'),
    ('30f68842-d8fb-42af-97f1-f677bd42f104','st-paul-district','St. Paul District', 'midas-project-default.jpg'),
    ('c069709f-24d9-4878-a10c-09246a351a38','tulsa-district','Tulsa District', 'midas-project-default.jpg'),
    ('be3320f5-977c-4923-a5ec-54d72f244d31','vicksburg-district','Vicksburg District', 'midas-project-default.jpg'),
    ('215bedb9-0486-481c-a927-d747257ab265','walla-walla-district','Walla Walla District', 'midas-project-default.jpg'),
    ('e0f8ab44-ae81-44fc-9e8a-d0d3f0b6b923','wilmington-district','Wilmington District', 'midas-project-default.jpg');

INSERT INTO project (id, slug, name, federal_id, image) VALUES
    ('57261121-e868-41d5-8046-20fabeb930e0','cts','CTS','NIST001','midas-project-default.jpg'),
    ('c50bee37-011f-4f02-a60b-3ff4baa9b8e2','chickamauga-lock','Chickamauga Lock','TN06504','midas-project-default.jpg'),
    ('676299fb-5c7a-4a96-95c7-fa81923920f6','kentucky-lock','Kentucky Lock','KY05017','midas-project-default.jpg'),
    ('f31c5a2c-347e-4534-a002-deae09c87ee6','sardis-dam-ok','Sardis Dam OK','OK22199','midas-project-default.jpg'),
    ('6756f41b-b5da-41ee-8a0d-d117d6f5b47a','caesar-creek-dam','Caesar Creek Dam','OH00927','midas-project-default.jpg'),
    ('0bb21182-acfc-4f42-ac0e-9ad8cfc26261','olmsted-locks-and-dam','Olmsted Locks and Dam','IL50745','midas-project-default.jpg');
