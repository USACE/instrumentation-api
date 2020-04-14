
-- Instrument Type
INSERT INTO instrument_type (id, name) VALUES
    ('0fd1f9ba-2731-4ff9-96dd-3c03215ab06f', 'Staff Gage'),
    ('1bb4bf7c-f5f8-44eb-9805-43b07ffadbef', 'Piezometer');

-- Instrument Groups
INSERT INTO instrument_group (id, slug, name, description) VALUES
    ('0beff703-07a6-4220-aa1d-dec8746bd4ca', 'r3-2', 'R3-2', '1.7 Miles East of Gate at S-354'),
    ('1362a4e4-0dc7-4bc2-a279-868787e16012', 'r3-3', 'R3-3', '1.0 Mile West of C-4'),
    ('97cadd04-74a5-482c-bdee-f7e7fa217e66', 'r3-4', 'R3-4', '3.5 Miles East of Gate at S-354'),
    ('460c0de9-bd88-4141-a7ca-d7d69b0627a5', 'r3-5', 'R3-5', '0.8 Mile East of C-4A'),
    ('1e57130d-7dcd-4b01-8adf-77180f855d40', 'r3-6', 'R3-6', '0.6 Mile North East of South Bay Park Access Gate'),
    ('9f1f6ed2-5712-461e-bee2-e40979f464a5', 'r3-7', 'R3-7', '1.0 Mile North East of South Bay Park Access Gate'),
    ('5be75d58-2d16-4605-9652-72211e80b742', 'r3-1', 'R3-1', '0.7 Miles East of Gate at S-354');

-- Instruments
INSERT INTO instrument (id, slug, name, height, geometry, instrument_type_id) VALUES
    ('8a82bac3-61c5-4498-a2be-0cbb7e9f368c', 'pz-r3-g-b1', 'PZ-R3-G-B1', 43.40,ST_GeomFromText('POINT(-80.72044 26.69512)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('dbcaa68b-29ff-45bf-befa-43ee159881ba', 'pz-r3-g-b2', 'PZ-R3-G-B2', 43.52,ST_GeomFromText('POINT(-80.72042 26.69515)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('e6772734-05fb-4359-81e9-5ff2a164ba96', 'pz-r3-g-a1', 'PZ-R3-G-A1', 18.62,ST_GeomFromText('POINT(-80.72069 26.69536)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('8bd4e628-88a6-4ce7-8438-b50f3a6b3fd0', 'pz-r3-g-a2', 'PZ-R3-G-A2', 20.24,ST_GeomFromText('POINT(-80.72068 26.69534)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('e040bf6b-774d-4887-81b4-043155ff216a', 'pz-hhd16-r3-1c', 'PZ-HHD16-R3-1C', 13.17,ST_GeomFromText('POINT(-80.72016 26.69498)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('4cbd19fc-e123-4c32-a35a-8fc6b71b42f7', 'r3-7-tw', 'R3-7-TW', 1.41,ST_GeomFromText('POINT(-80.72011 26.69493)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('02688f1c-163e-4b1f-b3ae-906cb38e86f9', 'pz-er3p4-b1', 'PZ-ER3P4-B1', 42.25,ST_GeomFromText('POINT(-80.72472 26.68982)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('36528cb7-98b1-45fc-919b-d1086c55d4a5', 'pz-er3p4-b2', 'PZ-ER3P4-B2', 42.17,ST_GeomFromText('POINT(-80.72472 26.68982)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('a4c77d91-ca3d-4a5a-b75d-872bb10c314c', 'pz-r3-f-a1', 'PZ-R3-F-A1', 18.22,ST_GeomFromText('POINT(-80.72505 26.69005)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('8b3afa68-4b6a-491a-aff4-99ee06208d6b', 'pz-r3-f-a2', 'PZ-R3-F-A2', 18.23,ST_GeomFromText('POINT(-80.72507 26.69003)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('d32f6e60-2503-4b2a-8f0c-4ca475f46885', 'pz-er3p4-c', 'PZ-ER3P4-C', 11.48,ST_GeomFromText('POINT(-80.72443 26.68963)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('2350badc-4653-48e0-8280-5d53cf8d1c14', 'pz-hhd16-r3-2c', 'PZ-HHD16-R3-2C', 13.20,ST_GeomFromText('POINT(-80.72447 26.68967)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('f7a756f0-2e5f-448c-a0ed-2a2cd96a9cc3', 'r3-6-tw', 'R3-6-TW', 1.41,ST_GeomFromText('POINT(-80.72441 26.68956)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('e33653df-ec53-4126-91b3-af5774cb23eb', 'pz-r3-b-b1', 'PZ-R3-B-B1', 42.26,ST_GeomFromText('POINT(-80.73823 26.68107)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('69ad96ec-4696-407b-9c58-b8ec27a5bdfc', 'pz-r3-b-b2', 'PZ-R3-B-B2', 42.27,ST_GeomFromText('POINT(-80.73826 26.68107)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('af2e3327-120f-4d36-a52c-b22f99c72433', 'pz-hhd16-r3-3c', 'PZ-HHD16-R3-3C', 14.76,ST_GeomFromText('POINT(-80.73827 26.68082)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('ea1b950d-5d45-445f-9cfd-a126bb2968a5', 'r3-5-tw', 'R3-5-TW', 1.41,ST_GeomFromText('POINT(-80.73828 26.68071)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('7647df7e-88a4-4c1b-9881-2ab95fbad02c', 'pz-er3e-b', 'PZ-ER3E-B', 42.12,ST_GeomFromText('POINT(-80.75375 26.68420)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('171b776b-224b-49e9-9f91-070d22d99075', 'pz-r3-e-a1', 'PZ-R3-E-A1', 20.60,ST_GeomFromText('POINT(-80.75356 26.68457)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('5940a676-a4c5-4114-897e-de874857d8d1', 'pz-r3-e-a2', 'PZ-R3-E-A2', 20.55,ST_GeomFromText('POINT(-80.75357 26.68458)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('0105cc43-5cb6-4f06-977c-da3da5abe556', 'pz-r3-e-b2', 'PZ-R3-E-B2', 42.04,ST_GeomFromText('POINT(-80.75375 26.68419)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('667ccd7b-4f4c-4126-9953-cc0c0e7eb35b', 'pz-r3-e-c1', 'PZ-R3-E-C1', 18.05,ST_GeomFromText('POINT(-80.75389 26.68397)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('68851dd7-b466-4831-b0e4-c8ce8a45ea43', 'pz-r3-e-c2', 'PZ-R3-E-C2', 18.00,ST_GeomFromText('POINT(-80.75390 26.68398)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('60b68e17-a151-4568-8d22-02e94ead82a7', 'r3-4-tw', 'R3-4-TW', 1.41,ST_GeomFromText('POINT(-80.75390 26.68395)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('ed54e540-3953-44ae-aca9-a95ea668d349', 'pz-r3-c-a1', 'PZ-R3-C-A1', 23.54,ST_GeomFromText('POINT(-80.76456 26.69042)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('43c72cd2-8e0d-4876-bf0d-959642896ac6', 'pz-r3-c-a2', 'PZ-R3-C-A2', 23.59,ST_GeomFromText('POINT(-80.76454 26.69041)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('13ba47e5-5fcc-4e1e-8e1f-d2c76c839fb0', 'pz-r3-c-b', 'PZ-R3-C-B', 42.06,ST_GeomFromText('POINT(-80.76474 26.69017)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('3f07285b-79bc-4cb6-8939-bb182ca939ce', 'pz-r3-c-c1', 'PZ-R3-C-C1', 16.39,ST_GeomFromText('POINT(-80.76492 26.68999)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('9fdd0dd0-0e47-4aa5-8d53-34014ba67c4d', 'pz-r3-c-c2', 'PZ-R3-C-C2', 16.37,ST_GeomFromText('POINT(-80.76495 26.69000)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('a9f4e8d1-c208-40cd-bc66-55a02c5b9e08', 'r3-3-tw', 'R3-3-TW', 1.41,ST_GeomFromText('POINT(-80.76497 26.68993)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('9a30d375-0800-4297-891e-699fa8878899', 'pz-mrr3-sh-b1', 'PZ-MRR3-SH-B1', 41.77,ST_GeomFromText('POINT(-80.77763 26.69678)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('4e81c5a1-a25d-43e7-a1be-3a24930385f8', 'pz-mrr3-sh-b2', 'PZ-MRR3-SH-B2', 41.79,ST_GeomFromText('POINT(-80.77763 26.69678)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('77f0e427-e387-47b5-b707-b8dc5acad643', 'pz-r3-sh-b1a', 'PZ-R3-SH-B1A', 41.77,ST_GeomFromText('POINT(-80.77759 26.69677)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('27a43e2d-6b8d-4a9d-92c3-7e7b6de96d4a', 'pz-r3-sh-c1', 'PZ-R3-SH-C1', 19.56,ST_GeomFromText('POINT(-80.77775 26.69659)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('b0eccef9-ef09-44b4-b821-6cab400c0959', 'pz-r3-sh-c2', 'PZ-R3-SH-C2', 19.61,ST_GeomFromText('POINT(-80.77774 26.69659)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('277a786a-8512-4648-b8cf-9d610cdf8409', 'r3-2-tw', 'R3-2-TW', 1.41,ST_GeomFromText('POINT(-80.77779 26.69650)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}'),
    ('b1e3537c-6ef8-4c57-84b6-8d66f896472b', 'pz-er3d-b', 'PZ-ER3D-B', 40.96,ST_GeomFromText('POINT(-80.79336 26.69869)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('0e2cb2b9-bf18-4ed9-8792-85c286b085ba', 'pz-r3-d-a1', 'PZ-R3-D-A1', 21.24,ST_GeomFromText('POINT(-80.79334 26.69905)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('8b1fc50a-ce54-4ccf-b3c3-c9182e343647', 'pz-r3-d-a2', 'PZ-R3-D-A2', 21.03,ST_GeomFromText('POINT(-80.79332 26.69905)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('ba42dd69-bbe0-411d-9a20-12fa5ec0ebf3', 'pz-r3-d-c1', 'PZ-R3-D-C1', 14.07,ST_GeomFromText('POINT(-80.79336 26.69843)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('3629cbf6-1256-49ed-aebe-df2f661872a7', 'pz-r3-d-c2', 'PZ-R3-D-C2', 14.10,ST_GeomFromText('POINT(-80.79338 26.69843)',4326),'{1bb4bf7c-f5f8-44eb-9805-43b07ffadbef}'),
    ('9048e7c1-eba2-4888-9048-db63e71e85ef', 'r3-1-tw', 'R3-1-TW', 1.40,ST_GeomFromText('POINT(-80.79336 26.69835)',4326),'{0fd1f9ba-2731-4ff9-96dd-3c03215ab06f}');


-- Instrument Group Instruments
INSERT INTO instrument_group_instruments (instrument_id, instrument_group_id) VALUES
    ('8a82bac3-61c5-4498-a2be-0cbb7e9f368c', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('dbcaa68b-29ff-45bf-befa-43ee159881ba', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('e6772734-05fb-4359-81e9-5ff2a164ba96', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('8bd4e628-88a6-4ce7-8438-b50f3a6b3fd0', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('e040bf6b-774d-4887-81b4-043155ff216a', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('4cbd19fc-e123-4c32-a35a-8fc6b71b42f7', '9f1f6ed2-5712-461e-bee2-e40979f464a5'),
    ('02688f1c-163e-4b1f-b3ae-906cb38e86f9', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('36528cb7-98b1-45fc-919b-d1086c55d4a5', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('a4c77d91-ca3d-4a5a-b75d-872bb10c314c', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('8b3afa68-4b6a-491a-aff4-99ee06208d6b', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('d32f6e60-2503-4b2a-8f0c-4ca475f46885', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('2350badc-4653-48e0-8280-5d53cf8d1c14', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('f7a756f0-2e5f-448c-a0ed-2a2cd96a9cc3', '1e57130d-7dcd-4b01-8adf-77180f855d40'),
    ('e33653df-ec53-4126-91b3-af5774cb23eb', '460c0de9-bd88-4141-a7ca-d7d69b0627a5'),
    ('69ad96ec-4696-407b-9c58-b8ec27a5bdfc', '460c0de9-bd88-4141-a7ca-d7d69b0627a5'),
    ('af2e3327-120f-4d36-a52c-b22f99c72433', '460c0de9-bd88-4141-a7ca-d7d69b0627a5'),
    ('ea1b950d-5d45-445f-9cfd-a126bb2968a5', '460c0de9-bd88-4141-a7ca-d7d69b0627a5'),
    ('7647df7e-88a4-4c1b-9881-2ab95fbad02c', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('171b776b-224b-49e9-9f91-070d22d99075', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('5940a676-a4c5-4114-897e-de874857d8d1', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('0105cc43-5cb6-4f06-977c-da3da5abe556', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('667ccd7b-4f4c-4126-9953-cc0c0e7eb35b', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('68851dd7-b466-4831-b0e4-c8ce8a45ea43', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('60b68e17-a151-4568-8d22-02e94ead82a7', '97cadd04-74a5-482c-bdee-f7e7fa217e66'),
    ('ed54e540-3953-44ae-aca9-a95ea668d349', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('43c72cd2-8e0d-4876-bf0d-959642896ac6', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('13ba47e5-5fcc-4e1e-8e1f-d2c76c839fb0', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('3f07285b-79bc-4cb6-8939-bb182ca939ce', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('9fdd0dd0-0e47-4aa5-8d53-34014ba67c4d', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('a9f4e8d1-c208-40cd-bc66-55a02c5b9e08', '1362a4e4-0dc7-4bc2-a279-868787e16012'),
    ('9a30d375-0800-4297-891e-699fa8878899', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('4e81c5a1-a25d-43e7-a1be-3a24930385f8', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('77f0e427-e387-47b5-b707-b8dc5acad643', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('27a43e2d-6b8d-4a9d-92c3-7e7b6de96d4a', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('b0eccef9-ef09-44b4-b821-6cab400c0959', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('277a786a-8512-4648-b8cf-9d610cdf8409', '0beff703-07a6-4220-aa1d-dec8746bd4ca'),
    ('b1e3537c-6ef8-4c57-84b6-8d66f896472b', '5be75d58-2d16-4605-9652-72211e80b742'),
    ('0e2cb2b9-bf18-4ed9-8792-85c286b085ba', '5be75d58-2d16-4605-9652-72211e80b742'),
    ('8b1fc50a-ce54-4ccf-b3c3-c9182e343647', '5be75d58-2d16-4605-9652-72211e80b742'),
    ('ba42dd69-bbe0-411d-9a20-12fa5ec0ebf3', '5be75d58-2d16-4605-9652-72211e80b742'),
    ('3629cbf6-1256-49ed-aebe-df2f661872a7', '5be75d58-2d16-4605-9652-72211e80b742'),
    ('9048e7c1-eba2-4888-9048-db63e71e85ef', '5be75d58-2d16-4605-9652-72211e80b742');


-- Units
INSERT INTO unit (id, name) VALUES
('7b924ec2-c488-401d-a503-ca734b1ab804', 'pound-inch'),
('e0f8ecb4-02da-46f6-b10b-37b0bf09b43c', 'feet'),
('cbba941e-14cb-4d17-bc7f-b7156626a64a', 'unit');

-- Parameters
INSERT INTO parameter (id, name) VALUES
('068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'parameterexample');

-- Time Series
INSERT INTO timeseries (id, name, instrument_id, parameter_id, unit_id) VALUES
('cbba941e-14cb-4d17-bc7f-b7156626a64a', 'LAKEOKEE', '9048e7c1-eba2-4888-9048-db63e71e85ef', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac' , 'e0f8ecb4-02da-46f6-b10b-37b0bf09b43c');

-- Time Series Measurements
INSERT INTO timeseries_measurement (id, time, value, timeseries_id) VALUES
('1a5f08ea-de55-466a-b758-a9e2d8f15e31', '1/1/2020', '13.16', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('acfb521e-26e4-44f4-afc8-89e0e5cd1fb6', '1/2/2020', '13.16', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('56ca1fad-40a8-4e28-9135-b74897f85ce7', '1/3/2020', '13.17', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('41657ca9-ae5c-4870-b026-6fb7da36251c', '1/4/2020', '13.17', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('d92f2b07-aea1-437e-9cef-17e640bb1b1f', '1/5/2020', '13.13', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('861f3af7-c588-4ee9-91c6-e2499804d690', '1/6/2020', '13.12', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('0eda3e40-9e52-422b-b325-7ec285427259', '1/7/2020', '13.1', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('e1539122-da8a-4094-97a3-46bdd576f6bf', '1/8/2020', '13.08', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('9d316c28-ffde-4a61-8955-b27015944b49', '1/9/2020', '13.07', 'cbba941e-14cb-4d17-bc7f-b7156626a64a'),
('ba4c2963-71fc-4023-adc6-1a053b33e461', '1/10/2020', '13.05', 'cbba941e-14cb-4d17-bc7f-b7156626a64a');
