-- Project
INSERT INTO project (id, slug, name) VALUES
    ('278c353d-bd42-4906-b60a-152c9efafd2b', 'black-rock-dam', 'Black Rock Dam'),
    ('9e129231-8445-4e10-8e7c-989c1535b3b2', 'hancock-brook-dam', 'Hancock Brook Dam');

-- Instruments
INSERT INTO instrument (project_id, id, slug, name, geometry, type_id) VALUES
('278c353d-bd42-4906-b60a-152c9efafd2b', '6de29534-ca51-403a-944f-aa164e74b116', 'pz-1a', 'PZ-1A', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('278c353d-bd42-4906-b60a-152c9efafd2b', 'd8c5d568-b557-4076-abe3-b7d33ffa6765', 'pz-1b', 'PZ-1B', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('278c353d-bd42-4906-b60a-152c9efafd2b', '18cbf7c0-fccb-4e23-8281-709f9b2e3d4d', 'pz-2a', 'PZ-2A', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('278c353d-bd42-4906-b60a-152c9efafd2b', 'af8e8345-5447-4af2-9ec1-fb4c0740b446', 'pz-2b', 'PZ-2B', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('278c353d-bd42-4906-b60a-152c9efafd2b', '3d78a3c5-6b2e-44e4-8063-7374e43d5a18', 'pz-3a', 'PZ-3A', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('278c353d-bd42-4906-b60a-152c9efafd2b', 'a80fb98c-a062-4c9d-a219-b42d694d4d02', 'pz-3b', 'PZ-3B', ST_GeomFromText('POINT(-73.0381 41.6219)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'fcd82a2d-0269-487b-9e71-91a99649d4e1', 'pz-1', 'PZ-1', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'b4f19fb8-1f10-4f66-9d49-dd74ad766fc7', 'pz-2a-1', 'PZ-2A', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', '0827f14d-9d22-4dd6-918a-eb27fe469d88', 'pz-2b-1', 'PZ-2B', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'f1895968-1639-4575-9635-9d263e638cbf', 'pz-3-1', 'PZ-3', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'b96a34bf-9137-47f9-a3a2-97adfbde4e05', 'pz-4', 'PZ-4', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', '89797655-0682-415e-b714-257a7c1911c1', 'pz-5a', 'PZ-5A', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', '416b4ea0-2ae0-4fa3-8f2e-2ca7a4c2a3d1', 'pz-5b', 'PZ-5B', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'd3e59119-fe47-44b1-afec-d5815a22ac0f', 'pz-6a', 'PZ-6A', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'e85c6f6c-2fdf-4918-98c8-4fc26c0428b1', 'pz-6b', 'PZ-6B', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', '7c200e13-5ed6-40f2-9706-068842020009', 'pz-7a', 'PZ-7A', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'c4f79e18-97c9-4677-90c1-dd11ef925f4e', 'pz-7b', 'PZ-7B', ST_GeomFromText('POINT(-73.1056 41.6575)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
('9e129231-8445-4e10-8e7c-989c1535b3b2', 'c9a67340-9699-4616-8f01-52299b5994f3', 'barometer', 'Barometer', ST_GeomFromText('POINT(-73.1056 41.6575)',4326), '3350b1d1-a946-49a8-bf19-587d7163e0f7');

-- instrument_status
-- status 'active' for all instruments
INSERT INTO instrument_status (instrument_id, status_id) VALUES
('6de29534-ca51-403a-944f-aa164e74b116', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('d8c5d568-b557-4076-abe3-b7d33ffa6765', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('18cbf7c0-fccb-4e23-8281-709f9b2e3d4d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('af8e8345-5447-4af2-9ec1-fb4c0740b446', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('3d78a3c5-6b2e-44e4-8063-7374e43d5a18', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('a80fb98c-a062-4c9d-a219-b42d694d4d02', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('fcd82a2d-0269-487b-9e71-91a99649d4e1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('b4f19fb8-1f10-4f66-9d49-dd74ad766fc7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('0827f14d-9d22-4dd6-918a-eb27fe469d88', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('f1895968-1639-4575-9635-9d263e638cbf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('b96a34bf-9137-47f9-a3a2-97adfbde4e05', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('89797655-0682-415e-b714-257a7c1911c1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('416b4ea0-2ae0-4fa3-8f2e-2ca7a4c2a3d1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('d3e59119-fe47-44b1-afec-d5815a22ac0f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('e85c6f6c-2fdf-4918-98c8-4fc26c0428b1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('7c200e13-5ed6-40f2-9706-068842020009', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
('c4f79e18-97c9-4677-90c1-dd11ef925f4e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');
