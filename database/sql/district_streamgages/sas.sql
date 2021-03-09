INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('6b60b4e6-ff3a-4b8a-8426-93347824588f', '0154184e-2509-4485-b449-8eff4ab52eef', 'savannah-district-streamgages', 'Savannah District Streamgages', 'savannah-district-streamgages.jpg');





--INSERT INSTRUMENTS--COUNT:49
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('f15bbc65-b81b-4f2d-b953-8f407770488a', False, 'chattooga-river-at-burrells-ford-nr-pine-mtn-ga', 'CHATTOOGA RIVER AT BURRELLS FORD, NR PINE MTN, GA', null, ST_GeomFromText('POINT(-83.1194 34.9689)',4326), null, null, '2021-03-08T19:38:58.163478Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02176930'),
('ac22b3f0-391d-456b-9121-0915e9128395', False, 'chattooga-river-near-clayton-ga', 'CHATTOOGA RIVER NEAR CLAYTON, GA', null, ST_GeomFromText('POINT(-83.306 34.814)',4326), null, null, '2021-03-08T19:38:58.477096Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02177000'),
('5e07bfde-37a7-4647-add3-4967b97be84d', False, 'tallulah-river-near-clayton-ga', 'TALLULAH RIVER NEAR CLAYTON, GA', null, ST_GeomFromText('POINT(-83.5304 34.8904)',4326), null, null, '2021-03-08T19:38:58.784246Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02178400'),
('1f730421-b1c6-496a-be70-0eac7bbd0569', False, 'tallulah-river-ab-powerhouse-nr-tallulah-falls-ga', 'TALLULAH RIVER AB POWERHOUSE, NR TALLULAH FALLS,GA', null, ST_GeomFromText('POINT(-83.3757 34.732)',4326), null, null, '2021-03-08T19:38:59.142420Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02181580'),
('4ff4097b-41e4-41e7-b162-7e6f450d04e7', False, '02182000', '02182000', null, ST_GeomFromText('POINT(-83.3452 34.6779)',4326), null, null, '2021-03-08T19:38:59.451074Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02182000'),
('eaacc1af-594d-4402-9a9c-14e02091463c', False, 'beaverdam-creek-above-elberton-ga', 'BEAVERDAM CREEK ABOVE ELBERTON, GA', null, ST_GeomFromText('POINT(-82.8965 34.1687)',4326), null, null, '2021-03-08T19:38:59.785261Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02188600'),
('6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', False, 'broad-river-above-carlton-ga', 'BROAD RIVER ABOVE CARLTON, GA', null, ST_GeomFromText('POINT(-83.0033 34.0733)',4326), null, null, '2021-03-08T19:39:00.108872Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02191300'),
('9bb46afe-2234-4595-85df-3027ab01e462', False, 'broad-river-near-bell-ga', 'BROAD RIVER NEAR BELL, GA', null, ST_GeomFromText('POINT(-82.77 33.9742)',4326), null, null, '2021-03-08T19:39:01.176192Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02192000'),
('3c237d5a-d700-4092-9f01-cbd6314979a4', False, 'kettle-creek-near-washington-ga', 'KETTLE CREEK NEAR WASHINGTON, GA', null, ST_GeomFromText('POINT(-82.8579 33.6826)',4326), null, null, '2021-03-08T19:39:03.684370Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193340'),
('50a3d15a-9408-465b-a143-4655e5ed345d', False, 'little-river-near-washington-ga', 'LITTLE RIVER NEAR WASHINGTON, GA', null, ST_GeomFromText('POINT(-82.7425 33.6128)',4326), null, null, '2021-03-08T19:39:05.652339Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193500'),
('87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', False, 'kiokee-creek-at-ga-104-near-evans-ga', 'KIOKEE CREEK AT GA 104, NEAR EVANS, GA', null, ST_GeomFromText('POINT(-82.2326 33.601)',4326), null, null, '2021-03-08T19:39:06.748977Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02195320'),
('9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', False, 'butler-creek-below-7th-avenue-at-ft-gordon-ga', 'BUTLER CREEK BELOW 7TH AVENUE, AT FT. GORDON, GA', null, ST_GeomFromText('POINT(-82.1161 33.4386)',4326), null, null, '2021-03-08T19:39:07.168194Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196835'),
('952595d3-c5e9-40f4-a22c-ac4c0e8dee33', False, 'butler-creek-reservoir-at-fort-gordon-ga', 'BUTLER CREEK RESERVOIR AT FORT GORDON, GA', null, ST_GeomFromText('POINT(-82.0992 33.4258)',4326), null, null, '2021-03-08T19:39:07.506950Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196838'),
('005234ba-1cb4-4b79-ad5b-2b9982b7544b', False, 'spirit-creek-at-us-1-near-augusta-ga', 'SPIRIT CREEK AT US 1, NEAR AUGUSTA, GA', null, ST_GeomFromText('POINT(-82.139 33.3735)',4326), null, null, '2021-03-08T19:39:07.867176Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197020'),
('e08da097-1833-4ec5-b271-0c6ba47dcc4a', False, 'savannah-river-near-waynesboro-ga', 'SAVANNAH RIVER NEAR WAYNESBORO, GA', null, ST_GeomFromText('POINT(-81.7548 33.1499)',4326), null, null, '2021-03-08T19:39:08.182921Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021973269'),
('cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', False, 'brushy-creek-at-campground-road-near-wrens-ga', 'BRUSHY CREEK AT CAMPGROUND ROAD, NEAR WRENS, GA', null, ST_GeomFromText('POINT(-82.3343 33.1807)',4326), null, null, '2021-03-08T19:39:08.581997Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197598'),
('1afe57a9-fc1a-45bb-bfa8-e0a578a1ee64', False, 'brier-creek-near-waynesboro-ga', 'BRIER CREEK NEAR WAYNESBORO, GA', null, ST_GeomFromText('POINT(-81.9637 33.1182)',4326), null, null, '2021-03-08T19:39:08.926314Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197830'),
('97a23ce8-23bb-45bc-9ab3-16d57dce197c', False, 'brier-creek-at-millhaven-ga', 'BRIER CREEK AT MILLHAVEN, GA', null, ST_GeomFromText('POINT(-81.6512 32.9335)',4326), null, null, '2021-03-08T19:39:09.245560Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198000'),
('84cfa98f-6a00-4a3d-97b4-442b9212b6a7', False, 'beaverdam-creek-near-sardis-ga', 'BEAVERDAM CREEK NEAR SARDIS, GA', null, ST_GeomFromText('POINT(-81.8154 32.9377)',4326), null, null, '2021-03-08T19:39:09.575962Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198100'),
('971212d2-be5e-4b80-b2d7-0d19acc01cfa', False, 'ebenezer-creek-at-springfield-ga', 'EBENEZER CREEK AT SPRINGFIELD, GA', null, ST_GeomFromText('POINT(-81.2973 32.3657)',4326), null, null, '2021-03-08T19:39:09.894247Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198690'),
('3d69499f-9a7e-4efe-81b5-fe0f84df3d61', False, 'withlacoochee-river-at-us-41-near-valdosta-ga', 'WITHLACOOCHEE RIVER AT US 41 NEAR VALDOSTA GA', null, ST_GeomFromText('POINT(-81.1615 32.3532)',4326), null, null, '2021-03-08T19:39:10.236308Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198745'),
('6a157044-3688-45cb-98cf-c4515c024b4f', False, 'yellow-river-at-milstead-ga', 'YELLOW RIVER AT MILSTEAD GA', null, ST_GeomFromText('POINT(-81.1464 32.3003)',4326), null, null, '2021-03-08T19:39:10.558973Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198768'),
('a350805c-0d10-4fa0-9351-43424b2beca2', False, 'abercorn-creek-near-savannah-ga', 'ABERCORN CREEK NEAR SAVANNAH,GA', null, ST_GeomFromText('POINT(-81.1782 32.2558)',4326), null, null, '2021-03-08T19:39:10.913953Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198810'),
('31509bea-a87b-4d92-af78-aa825ff3ae2a', False, 'savannah-river-i-95-near-port-wentworth-ga', 'SAVANNAH RIVER (I-95) NEAR PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1512 32.2358)',4326), null, null, '2021-03-08T19:39:11.238463Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198840'),
('4efebec3-04ac-4f9f-b49a-451770117fac', False, 'savannah-river-at-ga-25-at-port-wentworth-ga', 'SAVANNAH RIVER AT GA 25, AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1537 32.166)',4326), null, null, '2021-03-08T19:39:11.610115Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198920'),
('6453d62e-e970-4592-a696-74f73bc09901', False, 'middle-river-at-ga-25-at-port-wentworth-ga', 'MIDDLE RIVER AT GA 25 AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1383 32.1656)',4326), null, null, '2021-03-08T19:39:11.946529Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198950'),
('6492879d-184d-40e0-902f-65a26f121a83', False, 'savannah-river-at-usace-dock-at-savannah-ga', 'SAVANNAH RIVER AT USACE DOCK, AT SAVANNAH, GA', null, ST_GeomFromText('POINT(-81.0812 32.081)',4326), null, null, '2021-03-08T19:39:12.275451Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989773'),
('413b15a2-4666-4c43-a911-7e02548a3841', False, 'l-back-river-above-lucknow-canal-nr-limehouse-sc', 'L BACK RIVER ABOVE LUCKNOW CANAL, NR LIMEHOUSE, SC', null, ST_GeomFromText('POINT(-81.1179 32.1858)',4326), null, null, '2021-03-08T19:39:12.605072Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989784'),
('dead0d24-43b5-4383-9538-aa0769de9206', False, 'little-back-river-at-f&w-dock-near-limehouse-sc', 'LITTLE BACK RIVER AT F&W DOCK, NEAR LIMEHOUSE, SC', null, ST_GeomFromText('POINT(-81.1182 32.1708)',4326), null, null, '2021-03-08T19:39:12.929612Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989791'),
('cda5713a-6f72-4d10-9831-2ae368a7a909', False, 'little-back-river-at-ga-25-at-port-wentworth-ga', 'LITTLE BACK RIVER AT GA 25 AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.13 32.1658)',4326), null, null, '2021-03-08T19:39:13.257414Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989792'),
('093180ba-2e1d-4bbb-9dd4-feab113ce564', False, 'savannah-river-at-fort-pulaski-ga', 'SAVANNAH RIVER AT FORT PULASKI, GA', null, ST_GeomFromText('POINT(-80.9032 32.0341)',4326), null, null, '2021-03-08T19:39:13.661734Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198980'),
('924a4b3d-aa54-452d-955d-d5763b916c03', False, 'south-channel-savannah-river-near-savannah-ga', 'SOUTH CHANNEL (SAVANNAH RIVER) NEAR SAVANNAH, GA', null, ST_GeomFromText('POINT(-81.0023 32.0827)',4326), null, null, '2021-03-08T19:39:13.982066Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02199000'),
('0f2a5b22-363e-472a-ab11-8aabd9779af2', False, 'twelvemile-creek-near-liberty-sc', 'TWELVEMILE CREEK NEAR LIBERTY, SC', null, ST_GeomFromText('POINT(-82.7485 34.8015)',4326), null, null, '2021-03-08T19:39:14.326776Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02186000'),
('420d5ecc-20f5-46cb-bbfc-f9ec08874283', False, 'eighteenmile-creek-above-pendleton-sc', 'EIGHTEENMILE CREEK ABOVE PENDLETON, SC', null, ST_GeomFromText('POINT(-82.7988 34.659)',4326), null, null, '2021-03-08T19:39:14.709093Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02186699'),
('8ae11ab9-2efa-410f-ac1c-ce1eb4b9e33f', False, 'hartwell-lake-near-anderson-sc', 'HARTWELL LAKE NEAR ANDERSON, SC', null, ST_GeomFromText('POINT(-82.8161 34.475)',4326), null, null, '2021-03-08T19:39:15.030117Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02187010'),
('fd46c733-b662-4195-b96d-0369df362a11', False, 'rocky-river-nr-starr-sc', 'ROCKY RIVER NR STARR, SC', null, ST_GeomFromText('POINT(-82.5774 34.3832)',4326), null, null, '2021-03-08T19:39:15.370186Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02187910'),
('8f7dc916-885e-43dd-850a-509741483b80', False, 'russell-lake-above-calhoun-falls-sc', 'RUSSELL LAKE ABOVE CALHOUN FALLS, SC', null, ST_GeomFromText('POINT(-82.6181 34.1011)',4326), null, null, '2021-03-08T19:39:15.706762Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02188100'),
('f956ed57-4c11-45aa-ba97-137fc08f26f2', False, 'little-river-near-mt-carmel-sc', 'LITTLE RIVER NEAR MT. CARMEL, SC', null, ST_GeomFromText('POINT(-82.5007 34.0715)',4326), null, null, '2021-03-08T19:39:16.032828Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02192500'),
('7a1d3531-440f-4a13-bc32-f43101ebac4d', False, 'thurmond-lake-near-plum-branch-sc', 'THURMOND LAKE NEAR PLUM BRANCH, SC', null, ST_GeomFromText('POINT(-82.3506 33.8403)',4326), null, null, '2021-03-08T19:39:16.376662Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193900'),
('0247a7d6-fb7f-4968-b6dc-c2e8f277dab6', False, 'savannah-river-near-evans-ga', 'SAVANNAH RIVER NEAR EVANS, GA', null, ST_GeomFromText('POINT(-82.1232 33.5929)',4326), null, null, '2021-03-08T19:39:16.718617Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02195520'),
('195be647-588b-4973-bf02-8717a31d02e1', False, 'stevens-creek-near-modoc-sc', 'STEVENS CREEK NEAR MODOC, SC', null, ST_GeomFromText('POINT(-82.1818 33.7293)',4326), null, null, '2021-03-08T19:39:17.038625Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196000'),
('c2c43566-82eb-4901-a1d3-81f2200e193e', False, 'savannah-rvr-at-stevens-creek-dam-nr-morgana-sc', 'SAVANNAH RVR AT STEVENS CREEK DAM NR MORGANA, SC', null, ST_GeomFromText('POINT(-82.051 33.5629)',4326), null, null, '2021-03-08T19:39:17.443320Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196483'),
('f8a51a89-74ce-4c03-bc93-594c97b7de00', False, 'augusta-canal-nr-augusta-ga-upper', 'AUGUSTA CANAL NR AUGUSTA, GA (UPPER)', null, ST_GeomFromText('POINT(-82.0379 33.5493)',4326), null, null, '2021-03-08T19:39:17.787735Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196485'),
('34b223e2-5394-466b-b7ac-9d84788a4db3', False, 'horse-creek-at-clearwater-sc', 'HORSE CREEK AT CLEARWATER, SC', null, ST_GeomFromText('POINT(-81.8971 33.4849)',4326), null, null, '2021-03-08T19:39:18.117363Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196690'),
('01de28df-bf75-428c-ae1d-f4feb975db63', False, '021970161', '021970161', null, ST_GeomFromText('POINT(-82.1572 33.39)',4326), null, null, '2021-03-08T19:39:18.436712Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021970161'),
('ace2e8fb-6357-4bc2-90be-b47e426caf1d', False, 'savannah-r-at-burtons-ferry-br-nr-millhaven-ga', 'SAVANNAH R AT BURTONS FERRY BR NR MILLHAVEN, GA', null, ST_GeomFromText('POINT(-81.5026 32.939)',4326), null, null, '2021-03-08T19:39:18.754620Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197500'),
('02675aaf-22fc-426a-bcab-228f285857c1', False, 'savannah-river-near-estill-sc', 'SAVANNAH RIVER NEAR ESTILL, SC', null, ST_GeomFromText('POINT(-81.4283 32.7033)',4326), null, null, '2021-03-08T19:39:19.077534Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198375'),
('7c73c2ba-9e4f-4ac4-885b-d51aea328d5f', False, 'savannah-river-near-clyo-ga', 'SAVANNAH RIVER NEAR CLYO, GA', null, ST_GeomFromText('POINT(-81.2687 32.5282)',4326), null, null, '2021-03-08T19:39:19.398274Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198500'),
('cd09e10e-ab4d-49ba-830d-31e34f8faaa3', False, 'savannah-river-above-hardeeville-sc', 'SAVANNAH RIVER ABOVE HARDEEVILLE, SC', null, ST_GeomFromText('POINT(-81.1284 32.3394)',4326), null, null, '2021-03-08T19:39:19.711752Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198760');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('56fc9ce4-1cab-4432-9c41-354957de2fe1', 'f15bbc65-b81b-4f2d-b953-8f407770488a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:58.163478Z'),
('4ceca331-7e9d-4830-b43d-f043a7f99226', 'ac22b3f0-391d-456b-9121-0915e9128395', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:58.477096Z'),
('ffd94648-d052-4506-8856-91ddccf5be70', '5e07bfde-37a7-4647-add3-4967b97be84d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:58.784246Z'),
('533ce1b8-2de7-4bb4-ab1e-86ed823ec8fe', '1f730421-b1c6-496a-be70-0eac7bbd0569', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:59.142420Z'),
('19080927-0209-4edd-89fd-bf8f23df1e27', '4ff4097b-41e4-41e7-b162-7e6f450d04e7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:59.451074Z'),
('3a83e1fd-04da-4bd1-9ae7-702a19d5074b', 'eaacc1af-594d-4402-9a9c-14e02091463c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:38:59.785261Z'),
('d75be866-82f0-480a-8866-5d107182613a', '6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:00.108872Z'),
('c77134b8-9578-49f4-be5d-4b0227333070', '9bb46afe-2234-4595-85df-3027ab01e462', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:01.176192Z'),
('026fbe19-391f-4c9d-a930-32d09cd382cf', '3c237d5a-d700-4092-9f01-cbd6314979a4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:03.684370Z'),
('b90c8647-5654-45ef-b1ca-f08515aca959', '50a3d15a-9408-465b-a143-4655e5ed345d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:05.652339Z'),
('9820f737-5aab-4e35-8710-c29483c94316', '87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:06.748977Z'),
('ba826039-16be-4433-b76e-df31a1864de2', '9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:07.168194Z'),
('8c7b8518-fad6-40d9-875f-3aff882c43fc', '952595d3-c5e9-40f4-a22c-ac4c0e8dee33', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:07.506950Z'),
('f0fb8102-3966-4867-8178-c97d0f118040', '005234ba-1cb4-4b79-ad5b-2b9982b7544b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:07.867176Z'),
('de5f9f66-1482-4a10-91b4-093535ae7451', 'e08da097-1833-4ec5-b271-0c6ba47dcc4a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:08.182921Z'),
('16c49fd6-c2c4-497a-8e93-d958dcde4add', 'cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:08.581997Z'),
('c07ff517-6e3c-4b23-86c9-1215b23de2da', '1afe57a9-fc1a-45bb-bfa8-e0a578a1ee64', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:08.926314Z'),
('53410755-9f0a-4366-93dd-fab75c595c73', '97a23ce8-23bb-45bc-9ab3-16d57dce197c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:09.245560Z'),
('8c23f477-fa17-4206-b507-b0ca47200747', '84cfa98f-6a00-4a3d-97b4-442b9212b6a7', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:09.575962Z'),
('f8bd14ae-21d5-4b4f-a4a8-ce1e82fc223b', '971212d2-be5e-4b80-b2d7-0d19acc01cfa', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:09.894247Z'),
('a96889d9-6832-4392-95f6-996fbab11a66', '3d69499f-9a7e-4efe-81b5-fe0f84df3d61', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:10.236308Z'),
('f40337b7-09de-4bd5-9acc-6abc59726d9b', '6a157044-3688-45cb-98cf-c4515c024b4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:10.558973Z'),
('5d4228a3-6e38-4acb-b2bb-416e033dc52d', 'a350805c-0d10-4fa0-9351-43424b2beca2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:10.913953Z'),
('95a053b5-c60d-4325-ba43-b0ce249f2d82', '31509bea-a87b-4d92-af78-aa825ff3ae2a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:11.238463Z'),
('959ba2e4-14a2-4b42-981f-183c30aabb03', '4efebec3-04ac-4f9f-b49a-451770117fac', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:11.610115Z'),
('1c9fe4ab-a44c-4f11-b8aa-3eb48aaf97c3', '6453d62e-e970-4592-a696-74f73bc09901', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:11.946529Z'),
('7b6b69aa-0b20-42ca-a1e7-a91c56f7fee7', '6492879d-184d-40e0-902f-65a26f121a83', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:12.275451Z'),
('c51227cf-c41f-46d6-a6f7-78ce702016a6', '413b15a2-4666-4c43-a911-7e02548a3841', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:12.605072Z'),
('aefc5dfc-ec40-452e-b6f9-b4c6c16fe01e', 'dead0d24-43b5-4383-9538-aa0769de9206', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:12.929612Z'),
('809f63a0-bfea-4949-aaf8-c1161d964793', 'cda5713a-6f72-4d10-9831-2ae368a7a909', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:13.257414Z'),
('99495ef0-032d-46d0-b05e-94b205bae353', '093180ba-2e1d-4bbb-9dd4-feab113ce564', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:13.661734Z'),
('3724861e-27c5-4f69-a4c9-54ac2d0b7c0f', '924a4b3d-aa54-452d-955d-d5763b916c03', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:13.982066Z'),
('d11a96ad-fa16-4a00-9dd4-ca9bb4aeed3d', '0f2a5b22-363e-472a-ab11-8aabd9779af2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:14.326776Z'),
('7abacbfa-0f38-4177-b93a-752fef8610ed', '420d5ecc-20f5-46cb-bbfc-f9ec08874283', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:14.709093Z'),
('0e10eaf9-3040-45e2-a93b-4944f05df31e', '8ae11ab9-2efa-410f-ac1c-ce1eb4b9e33f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:15.030117Z'),
('8b15b0e1-925e-4724-b46f-48b5e319f40f', 'fd46c733-b662-4195-b96d-0369df362a11', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:15.370186Z'),
('6fbfd889-f040-4b7b-b212-428d65bf5c4f', '8f7dc916-885e-43dd-850a-509741483b80', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:15.706762Z'),
('6d04ba8b-7d42-44dc-bb45-bfd710bcebe9', 'f956ed57-4c11-45aa-ba97-137fc08f26f2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:16.032828Z'),
('a6e116ae-c12f-466c-af2a-a8d527c0f1d4', '7a1d3531-440f-4a13-bc32-f43101ebac4d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:16.376662Z'),
('46799b05-8918-4e29-ae8b-71f85eda077d', '0247a7d6-fb7f-4968-b6dc-c2e8f277dab6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:16.718617Z'),
('4af25777-3b53-4e2d-89dd-dcd13681c7e6', '195be647-588b-4973-bf02-8717a31d02e1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:17.038625Z'),
('f5bbd3b7-e5df-4976-8276-443b31156010', 'c2c43566-82eb-4901-a1d3-81f2200e193e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:17.443320Z'),
('7daa037b-4973-43c2-afbc-2e971f53d7a7', 'f8a51a89-74ce-4c03-bc93-594c97b7de00', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:17.787735Z'),
('d6cc65b9-fede-4d39-989e-e507ef7ba973', '34b223e2-5394-466b-b7ac-9d84788a4db3', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:18.117363Z'),
('f9d3730a-e18d-4a72-b56a-8aa0435261f6', '01de28df-bf75-428c-ae1d-f4feb975db63', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:18.436712Z'),
('8daf1611-896f-4e1f-9b75-793c0a50d720', 'ace2e8fb-6357-4bc2-90be-b47e426caf1d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:18.754620Z'),
('ed174605-d3fd-4cb4-a4ea-986a30612ac0', '02675aaf-22fc-426a-bcab-228f285857c1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:19.077534Z'),
('d2e7532a-75ca-4055-a453-271ae4632c6c', '7c73c2ba-9e4f-4ac4-885b-d51aea328d5f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:19.398274Z'),
('00e56ff3-a439-4289-a391-f9c076f2200e', 'cd09e10e-ab4d-49ba-830d-31e34f8faaa3', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-08T19:39:19.711752Z');

--INSERT TELEMETRY_GOES--COUNT:43
INSERT INTO public.telemetry_goes (id, nesdis_id) select '4d3fbc0f-f319-4d4d-a247-c49603fa4210', 'DE203122' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE203122');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '68fbef21-425e-447c-8d91-5e83018c2c23', '17ACC13E' where not exists (select 1 from telemetry_goes where nesdis_id = '17ACC13E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f104b569-ae4f-417f-9111-b9a2150390fc', 'DE25E08A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE25E08A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '57631351-b9a1-4945-b486-4f3a7447b41f', 'CE8771CE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8771CE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '2e0705c6-a21a-4f86-aaaa-76df2f86e587', 'DE15F58C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE15F58C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '14147a70-1d65-422d-9a15-59d021fa919c', 'DD8D036E' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8D036E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '54854d6b-60c2-4cbc-b487-ff6b41b011a6', 'DE160206' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE160206');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'fbb6d711-9fe3-4c7d-8243-6f51a47d5eae', '17E731CE' where not exists (select 1 from telemetry_goes where nesdis_id = '17E731CE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '393733c9-df3b-42a3-8fba-fa0099ff421a', 'CE8721B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8721B2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'dd430e84-d742-42ae-a0e8-11cab542b28b', '17F2177C' where not exists (select 1 from telemetry_goes where nesdis_id = '17F2177C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '877892a9-a8a7-4325-8e4d-be4b853ebef5', 'DD22D5C2' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD22D5C2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '16451f1f-205b-4472-ace6-ec33e49e2ded', '17E7475E' where not exists (select 1 from telemetry_goes where nesdis_id = '17E7475E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd8d75b04-4d73-41cc-b2a8-b8f08b0df1fe', 'DD64F03C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD64F03C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'db6e2920-37d6-424c-935e-be18c18a8ad3', 'DDB645BC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDB645BC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd0c235ca-e8c8-4096-84fb-d16762e4640f', 'DDB0A680' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDB0A680');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '3393cca4-af89-4d1e-903c-b844fe8d81e5', 'DD49F2B4' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD49F2B4');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '31accc32-e5fb-47ae-b84d-9c73247024ad', 'DD4A1648' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD4A1648');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '10f81da1-4150-4337-b778-db48980e122b', '17B5F566' where not exists (select 1 from telemetry_goes where nesdis_id = '17B5F566');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'abf469c4-de53-48d7-8e80-98acd6f75226', '17B6670A' where not exists (select 1 from telemetry_goes where nesdis_id = '17B6670A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'cbd3aee5-200e-45da-9b6e-a3859eb22118', '17456210' where not exists (select 1 from telemetry_goes where nesdis_id = '17456210');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a6d7abac-c32e-41e6-8eaf-ea1f8f1fd3b1', '17D33284' where not exists (select 1 from telemetry_goes where nesdis_id = '17D33284');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '99c93727-907d-484f-a117-e31367c93cfa', 'DD24C67A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD24C67A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '020a8fb4-6376-4f97-9e81-e1a26ca2ce5f', '17B71360' where not exists (select 1 from telemetry_goes where nesdis_id = '17B71360');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '23dbf912-8c0a-4f3e-a183-b40857357d50', '17CDA410' where not exists (select 1 from telemetry_goes where nesdis_id = '17CDA410');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '94f84201-326d-4b67-a055-0d4e388a009c', '17B6E11E' where not exists (select 1 from telemetry_goes where nesdis_id = '17B6E11E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1848cc49-c439-47e2-8399-d4b24e9005b4', '17CD610E' where not exists (select 1 from telemetry_goes where nesdis_id = '17CD610E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b77d6d69-be7a-4eb4-a177-885a9886d65c', 'CE87923C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87923C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a93bfae2-3a55-45d1-811a-32999d85b57c', '17055456' where not exists (select 1 from telemetry_goes where nesdis_id = '17055456');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'cc215595-2858-46ab-83a4-8e912bbbd1f0', 'CE87814A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87814A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '438a8090-a689-4880-a2b5-11be7aabc7be', 'CE8762B8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8762B8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0543f508-765a-4304-8d3e-da9f1e91f8bd', 'CE875722' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE875722');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'afad3965-4d17-425f-878c-5f18f5cf65a9', 'CE874454' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE874454');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6f00276b-7ca4-483a-acc6-025cdeb6933b', 'CE871428' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE871428');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '9220ea0c-fde3-4f4b-98dd-f1a6b341175a', 'CE87075E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87075E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e19759d4-0733-4e61-bd3a-221df25fe839', 'DD2CC2DC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD2CC2DC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6e5d40b6-54f0-4d54-aa07-3bc6e3f7c1d8', 'CE87E4AC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87E4AC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ec756864-f512-4050-8cac-0e9effa346c6', '17EAE4C0' where not exists (select 1 from telemetry_goes where nesdis_id = '17EAE4C0');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0df578f5-361d-44ba-b9ff-563ddc178547', 'CE87F7DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87F7DA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '3b8cb8d1-ff7b-4f84-b035-179a9a2e0ac2', 'DE25D510' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE25D510');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ee706ea7-0395-485b-b58b-7a952faa5931', 'CE86B62A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE86B62A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '796f62c7-14c1-4da7-b38f-fa04874f8c7d', 'DE0CA432' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE0CA432');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6398c493-5816-4967-9503-80c0ed3d0df8', 'D11C04F8' where not exists (select 1 from telemetry_goes where nesdis_id = 'D11C04F8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '644ebca1-2760-4c80-bbd5-f268e4aef72b', 'DE0CD2A2' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE0CD2A2');

--INSERT INSTRUMENT_TELEMETRY--COUNT:43
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('f15bbc65-b81b-4f2d-b953-8f407770488a', '10a32652-af43-4451-bd52-4980c5690cc9', '4d3fbc0f-f319-4d4d-a247-c49603fa4210'),
('ac22b3f0-391d-456b-9121-0915e9128395', '10a32652-af43-4451-bd52-4980c5690cc9', '68fbef21-425e-447c-8d91-5e83018c2c23'),
('4ff4097b-41e4-41e7-b162-7e6f450d04e7', '10a32652-af43-4451-bd52-4980c5690cc9', 'f104b569-ae4f-417f-9111-b9a2150390fc'),
('eaacc1af-594d-4402-9a9c-14e02091463c', '10a32652-af43-4451-bd52-4980c5690cc9', '57631351-b9a1-4945-b486-4f3a7447b41f'),
('6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', '10a32652-af43-4451-bd52-4980c5690cc9', '2e0705c6-a21a-4f86-aaaa-76df2f86e587'),
('9bb46afe-2234-4595-85df-3027ab01e462', '10a32652-af43-4451-bd52-4980c5690cc9', '14147a70-1d65-422d-9a15-59d021fa919c'),
('3c237d5a-d700-4092-9f01-cbd6314979a4', '10a32652-af43-4451-bd52-4980c5690cc9', '54854d6b-60c2-4cbc-b487-ff6b41b011a6'),
('50a3d15a-9408-465b-a143-4655e5ed345d', '10a32652-af43-4451-bd52-4980c5690cc9', 'fbb6d711-9fe3-4c7d-8243-6f51a47d5eae'),
('87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', '10a32652-af43-4451-bd52-4980c5690cc9', '393733c9-df3b-42a3-8fba-fa0099ff421a'),
('9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', '10a32652-af43-4451-bd52-4980c5690cc9', 'dd430e84-d742-42ae-a0e8-11cab542b28b'),
('952595d3-c5e9-40f4-a22c-ac4c0e8dee33', '10a32652-af43-4451-bd52-4980c5690cc9', '877892a9-a8a7-4325-8e4d-be4b853ebef5'),
('005234ba-1cb4-4b79-ad5b-2b9982b7544b', '10a32652-af43-4451-bd52-4980c5690cc9', '16451f1f-205b-4472-ace6-ec33e49e2ded'),
('e08da097-1833-4ec5-b271-0c6ba47dcc4a', '10a32652-af43-4451-bd52-4980c5690cc9', 'd8d75b04-4d73-41cc-b2a8-b8f08b0df1fe'),
('cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', '10a32652-af43-4451-bd52-4980c5690cc9', 'db6e2920-37d6-424c-935e-be18c18a8ad3'),
('97a23ce8-23bb-45bc-9ab3-16d57dce197c', '10a32652-af43-4451-bd52-4980c5690cc9', 'd0c235ca-e8c8-4096-84fb-d16762e4640f'),
('84cfa98f-6a00-4a3d-97b4-442b9212b6a7', '10a32652-af43-4451-bd52-4980c5690cc9', '3393cca4-af89-4d1e-903c-b844fe8d81e5'),
('971212d2-be5e-4b80-b2d7-0d19acc01cfa', '10a32652-af43-4451-bd52-4980c5690cc9', '31accc32-e5fb-47ae-b84d-9c73247024ad'),
('3d69499f-9a7e-4efe-81b5-fe0f84df3d61', '10a32652-af43-4451-bd52-4980c5690cc9', '10f81da1-4150-4337-b778-db48980e122b'),
('6a157044-3688-45cb-98cf-c4515c024b4f', '10a32652-af43-4451-bd52-4980c5690cc9', 'abf469c4-de53-48d7-8e80-98acd6f75226'),
('a350805c-0d10-4fa0-9351-43424b2beca2', '10a32652-af43-4451-bd52-4980c5690cc9', 'cbd3aee5-200e-45da-9b6e-a3859eb22118'),
('31509bea-a87b-4d92-af78-aa825ff3ae2a', '10a32652-af43-4451-bd52-4980c5690cc9', 'a6d7abac-c32e-41e6-8eaf-ea1f8f1fd3b1'),
('4efebec3-04ac-4f9f-b49a-451770117fac', '10a32652-af43-4451-bd52-4980c5690cc9', '99c93727-907d-484f-a117-e31367c93cfa'),
('6453d62e-e970-4592-a696-74f73bc09901', '10a32652-af43-4451-bd52-4980c5690cc9', '020a8fb4-6376-4f97-9e81-e1a26ca2ce5f'),
('6492879d-184d-40e0-902f-65a26f121a83', '10a32652-af43-4451-bd52-4980c5690cc9', '23dbf912-8c0a-4f3e-a183-b40857357d50'),
('cda5713a-6f72-4d10-9831-2ae368a7a909', '10a32652-af43-4451-bd52-4980c5690cc9', '94f84201-326d-4b67-a055-0d4e388a009c'),
('924a4b3d-aa54-452d-955d-d5763b916c03', '10a32652-af43-4451-bd52-4980c5690cc9', '1848cc49-c439-47e2-8399-d4b24e9005b4'),
('0f2a5b22-363e-472a-ab11-8aabd9779af2', '10a32652-af43-4451-bd52-4980c5690cc9', 'b77d6d69-be7a-4eb4-a177-885a9886d65c'),
('420d5ecc-20f5-46cb-bbfc-f9ec08874283', '10a32652-af43-4451-bd52-4980c5690cc9', 'a93bfae2-3a55-45d1-811a-32999d85b57c'),
('8ae11ab9-2efa-410f-ac1c-ce1eb4b9e33f', '10a32652-af43-4451-bd52-4980c5690cc9', 'cc215595-2858-46ab-83a4-8e912bbbd1f0'),
('fd46c733-b662-4195-b96d-0369df362a11', '10a32652-af43-4451-bd52-4980c5690cc9', '438a8090-a689-4880-a2b5-11be7aabc7be'),
('8f7dc916-885e-43dd-850a-509741483b80', '10a32652-af43-4451-bd52-4980c5690cc9', '0543f508-765a-4304-8d3e-da9f1e91f8bd'),
('f956ed57-4c11-45aa-ba97-137fc08f26f2', '10a32652-af43-4451-bd52-4980c5690cc9', 'afad3965-4d17-425f-878c-5f18f5cf65a9'),
('7a1d3531-440f-4a13-bc32-f43101ebac4d', '10a32652-af43-4451-bd52-4980c5690cc9', '6f00276b-7ca4-483a-acc6-025cdeb6933b'),
('0247a7d6-fb7f-4968-b6dc-c2e8f277dab6', '10a32652-af43-4451-bd52-4980c5690cc9', '9220ea0c-fde3-4f4b-98dd-f1a6b341175a'),
('195be647-588b-4973-bf02-8717a31d02e1', '10a32652-af43-4451-bd52-4980c5690cc9', 'e19759d4-0733-4e61-bd3a-221df25fe839'),
('c2c43566-82eb-4901-a1d3-81f2200e193e', '10a32652-af43-4451-bd52-4980c5690cc9', '6e5d40b6-54f0-4d54-aa07-3bc6e3f7c1d8'),
('f8a51a89-74ce-4c03-bc93-594c97b7de00', '10a32652-af43-4451-bd52-4980c5690cc9', 'ec756864-f512-4050-8cac-0e9effa346c6'),
('34b223e2-5394-466b-b7ac-9d84788a4db3', '10a32652-af43-4451-bd52-4980c5690cc9', '0df578f5-361d-44ba-b9ff-563ddc178547'),
('01de28df-bf75-428c-ae1d-f4feb975db63', '10a32652-af43-4451-bd52-4980c5690cc9', '3b8cb8d1-ff7b-4f84-b035-179a9a2e0ac2'),
('ace2e8fb-6357-4bc2-90be-b47e426caf1d', '10a32652-af43-4451-bd52-4980c5690cc9', 'ee706ea7-0395-485b-b58b-7a952faa5931'),
('02675aaf-22fc-426a-bcab-228f285857c1', '10a32652-af43-4451-bd52-4980c5690cc9', '796f62c7-14c1-4da7-b38f-fa04874f8c7d'),
('7c73c2ba-9e4f-4ac4-885b-d51aea328d5f', '10a32652-af43-4451-bd52-4980c5690cc9', '6398c493-5816-4967-9503-80c0ed3d0df8'),
('cd09e10e-ab4d-49ba-830d-31e34f8faaa3', '10a32652-af43-4451-bd52-4980c5690cc9', '644ebca1-2760-4c80-bbd5-f268e4aef72b');

--INSERT TIMESERIES--COUNT:49
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('0df6f857-4922-4e6d-94b4-5d8df4fb4145','stage','Stage','f15bbc65-b81b-4f2d-b953-8f407770488a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('76c6dd56-29a0-4f6c-afe4-3024acb123d7','precipitation','Precipitation','f15bbc65-b81b-4f2d-b953-8f407770488a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('4aae8615-1c73-4855-88f7-27b583570282','voltage','Voltage','f15bbc65-b81b-4f2d-b953-8f407770488a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e7e9a46f-27eb-4c7e-86e4-a8112e79dcb8','stage','Stage','ac22b3f0-391d-456b-9121-0915e9128395', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fbdadc52-73cf-4fc1-af62-e3323bf85f70','precipitation','Precipitation','ac22b3f0-391d-456b-9121-0915e9128395', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('16f71512-7afa-4bc8-9631-7927975d9244','voltage','Voltage','ac22b3f0-391d-456b-9121-0915e9128395', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4c6df559-4598-4105-80a1-f64d51a4b239','stage','Stage','5e07bfde-37a7-4647-add3-4967b97be84d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('06709b44-e6b6-4ecf-8ad7-79b61418260f','precipitation','Precipitation','5e07bfde-37a7-4647-add3-4967b97be84d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('7d2c508e-5f4d-42a9-91c3-bbc98be6ca74','voltage','Voltage','5e07bfde-37a7-4647-add3-4967b97be84d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1404f468-720c-47ed-ab13-d3c3dffbf909','stage','Stage','1f730421-b1c6-496a-be70-0eac7bbd0569', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('123ac942-25ee-4604-92e7-a8f1bd0f3752','precipitation','Precipitation','1f730421-b1c6-496a-be70-0eac7bbd0569', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('37d6ed74-eea5-406e-abc7-e753b908954f','voltage','Voltage','1f730421-b1c6-496a-be70-0eac7bbd0569', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('18d33bea-4b3e-4451-94c6-e73ccf5f883f','stage','Stage','4ff4097b-41e4-41e7-b162-7e6f450d04e7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('28520460-fe3c-4a42-9e70-cb5ba1c3f6f4','precipitation','Precipitation','4ff4097b-41e4-41e7-b162-7e6f450d04e7', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('843bd39e-ac99-4037-8641-6fd009f89319','voltage','Voltage','4ff4097b-41e4-41e7-b162-7e6f450d04e7', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a0f44093-b1c1-4f30-8594-3d219ba3846e','stage','Stage','eaacc1af-594d-4402-9a9c-14e02091463c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('8eb40909-e155-4aa2-9f2e-9e582789c1f0','precipitation','Precipitation','eaacc1af-594d-4402-9a9c-14e02091463c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('500649a4-68ee-47a8-8350-6a1be120e0b4','voltage','Voltage','eaacc1af-594d-4402-9a9c-14e02091463c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('951f13c3-51f1-44fc-b744-40c46330da5e','stage','Stage','6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('7b0d73b8-101f-4b3b-897e-555d6fb99baf','precipitation','Precipitation','6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6e9a1bfa-4263-4faa-a3f5-03e791c8d2a1','voltage','Voltage','6c8bee7f-bdfd-4ed3-82f0-f6f228efc516', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f70cf112-3ba7-4232-a125-433bd101a8d1','stage','Stage','9bb46afe-2234-4595-85df-3027ab01e462', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('68532f9e-19b4-4de0-acfe-ba8fb52a459e','voltage','Voltage','9bb46afe-2234-4595-85df-3027ab01e462', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('06ee17c8-393a-462d-ac7f-b43565a5630a','stage','Stage','3c237d5a-d700-4092-9f01-cbd6314979a4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('71f2453e-9df2-429e-ae71-e16cc6363ba4','voltage','Voltage','3c237d5a-d700-4092-9f01-cbd6314979a4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('11bcc924-ad2b-4ece-ace9-186153a0cd4e','stage','Stage','50a3d15a-9408-465b-a143-4655e5ed345d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c742b520-6966-4002-8249-bf9e5790a7d6','precipitation','Precipitation','50a3d15a-9408-465b-a143-4655e5ed345d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a4bfb7eb-f91a-47f4-8245-16a9e19ed8fb','voltage','Voltage','50a3d15a-9408-465b-a143-4655e5ed345d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1594b94b-791a-4c47-8583-33b39eae2a45','stage','Stage','87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b0d73b25-b848-4d78-879e-bdd098bc3025','precipitation','Precipitation','87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9663b72b-dddc-48fa-8428-3f0815ff2890','voltage','Voltage','87a81cbd-8cc8-4f29-aab3-d1ebedb78b24', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('46f2b11e-dd44-4aa7-834a-9a159691079e','stage','Stage','9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('fd95bf71-c270-4d8b-b8d2-beef2eac9077','precipitation','Precipitation','9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('bff924c1-9550-4fbe-8647-98dfda56d838','voltage','Voltage','9634693a-1d5f-4e95-8e78-c2ae40ac5a0c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f74a22ff-612d-42e2-9c47-10764887793e','precipitation','Precipitation','952595d3-c5e9-40f4-a22c-ac4c0e8dee33', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('41888fd1-467e-47a1-ac85-94acafc5003a','voltage','Voltage','952595d3-c5e9-40f4-a22c-ac4c0e8dee33', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1a15c1de-6077-4f67-8e64-fc3b8c0cb78e','stage','Stage','005234ba-1cb4-4b79-ad5b-2b9982b7544b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('12cdde05-3161-409d-9b0d-20d957e4c596','precipitation','Precipitation','005234ba-1cb4-4b79-ad5b-2b9982b7544b', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9b9821c5-0aee-4da9-a45c-f34a86b3efe8','voltage','Voltage','005234ba-1cb4-4b79-ad5b-2b9982b7544b', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('15739537-5752-4ed1-a826-d5af590555bc','stage','Stage','e08da097-1833-4ec5-b271-0c6ba47dcc4a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5a814edb-5e0e-45dd-857c-796f3108f69c','precipitation','Precipitation','e08da097-1833-4ec5-b271-0c6ba47dcc4a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e4482d86-f777-4b4e-ab33-23e504506ac9','voltage','Voltage','e08da097-1833-4ec5-b271-0c6ba47dcc4a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('7770dc71-7a76-483f-89a0-2df722f604ce','stage','Stage','cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1e6a1e6a-8822-441c-b50b-6c76cada5d26','precipitation','Precipitation','cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b7a131b3-4ebd-48c1-82ca-0f813e36dd56','voltage','Voltage','cefbe623-b8de-4e6f-87e2-992b0dbf4fe9', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('7eaaf37d-3f00-42b4-aa90-899cb175f160','stage','Stage','1afe57a9-fc1a-45bb-bfa8-e0a578a1ee64', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1a73c09c-5dec-492c-988d-fc9a0c208156','precipitation','Precipitation','1afe57a9-fc1a-45bb-bfa8-e0a578a1ee64', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('5a2d676a-a609-4bad-b753-6041c1e7f626','voltage','Voltage','1afe57a9-fc1a-45bb-bfa8-e0a578a1ee64', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('08b27001-cdec-4fb2-a5dd-abe6250eb9f6','stage','Stage','97a23ce8-23bb-45bc-9ab3-16d57dce197c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d4781458-1ceb-46f3-bd2e-5bf4eee70f0f','precipitation','Precipitation','97a23ce8-23bb-45bc-9ab3-16d57dce197c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('2043d9c8-728b-4e2c-a659-af324c989b86','voltage','Voltage','97a23ce8-23bb-45bc-9ab3-16d57dce197c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('489550f9-6b11-420e-8a58-dcff07ce145e','stage','Stage','84cfa98f-6a00-4a3d-97b4-442b9212b6a7', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e65c7026-4dcf-4478-acc9-1efcb11b3910','precipitation','Precipitation','84cfa98f-6a00-4a3d-97b4-442b9212b6a7', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a1cb1459-d8ac-45d7-98ca-0210fb24234e','voltage','Voltage','84cfa98f-6a00-4a3d-97b4-442b9212b6a7', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('dd02fba2-60cb-4dfd-b4dc-9c46474c0a01','stage','Stage','971212d2-be5e-4b80-b2d7-0d19acc01cfa', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d5d5c25f-1b24-4299-9771-cd5c2adfa5dd','voltage','Voltage','971212d2-be5e-4b80-b2d7-0d19acc01cfa', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('221e52fd-cdf9-4f7e-abe1-5f606e033922','stage','Stage','3d69499f-9a7e-4efe-81b5-fe0f84df3d61', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('434c348c-4841-4e96-bf87-61fa6e0fe21b','precipitation','Precipitation','3d69499f-9a7e-4efe-81b5-fe0f84df3d61', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b01432df-e4ce-494b-b900-1744e30c2f05','voltage','Voltage','3d69499f-9a7e-4efe-81b5-fe0f84df3d61', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8025c379-fcd0-4097-b0f3-b8f218d77bbb','stage','Stage','a350805c-0d10-4fa0-9351-43424b2beca2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('010a7f9f-dab4-4144-b0d8-ec6c6b51b986','precipitation','Precipitation','a350805c-0d10-4fa0-9351-43424b2beca2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('82ed9081-bb23-46c4-a8cf-71881b1f9035','voltage','Voltage','a350805c-0d10-4fa0-9351-43424b2beca2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('255b4153-6450-4ac5-9fb8-7be4a8786900','stage','Stage','31509bea-a87b-4d92-af78-aa825ff3ae2a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ea5687e8-abf6-400e-afe3-130b152e32c7','precipitation','Precipitation','31509bea-a87b-4d92-af78-aa825ff3ae2a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a02f2951-db24-4493-8f7d-b7eaf669e99b','voltage','Voltage','31509bea-a87b-4d92-af78-aa825ff3ae2a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a9a7c6dd-4586-4481-8669-910ece078db3','stage','Stage','4efebec3-04ac-4f9f-b49a-451770117fac', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a17b0aa8-f6ee-4f48-ac98-81c17d92de4b','precipitation','Precipitation','4efebec3-04ac-4f9f-b49a-451770117fac', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('d68182df-04a5-4f78-adbe-c1729d92b8f4','voltage','Voltage','4efebec3-04ac-4f9f-b49a-451770117fac', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('469d23a1-058f-474a-8bc5-096697517447','stage','Stage','6492879d-184d-40e0-902f-65a26f121a83', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9b6aef75-0032-4db4-91fd-be2676b8ce79','precipitation','Precipitation','6492879d-184d-40e0-902f-65a26f121a83', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('5903d1fd-7836-4922-be67-d098a031a9c7','dissolved-oxygen','Dissolved-Oxygen','6492879d-184d-40e0-902f-65a26f121a83', '98007857-d027-4524-9a63-d07ae93e5fa2', '67d75ccd-6bf7-4086-a970-5ed65a5c30f3'),
('051c9899-6032-4aea-959f-c3e032993130','voltage','Voltage','6492879d-184d-40e0-902f-65a26f121a83', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('22280276-4702-4b76-a510-51871b29bd8e','stage','Stage','413b15a2-4666-4c43-a911-7e02548a3841', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('250dfb96-5759-4196-a728-cdfb617a1ce9','precipitation','Precipitation','413b15a2-4666-4c43-a911-7e02548a3841', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('be1fecdc-9ff9-4371-bb8e-2428d0f1fa85','voltage','Voltage','413b15a2-4666-4c43-a911-7e02548a3841', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('815341e9-a9d4-4c1f-9c4f-b2a851205a1b','precipitation','Precipitation','dead0d24-43b5-4383-9538-aa0769de9206', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('bb6c05ad-5224-45f3-910d-5be182faf905','voltage','Voltage','dead0d24-43b5-4383-9538-aa0769de9206', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('26e8b86e-172d-41d4-95ea-d49f919bf2a7','stage','Stage','093180ba-2e1d-4bbb-9dd4-feab113ce564', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9488fd0b-79b1-4d3a-b76f-dac99e1bc1f9','precipitation','Precipitation','093180ba-2e1d-4bbb-9dd4-feab113ce564', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('bfef3b13-f709-48c8-ab5e-3a838946165c','voltage','Voltage','093180ba-2e1d-4bbb-9dd4-feab113ce564', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('fe5c8b1e-6ab0-4d00-b7a5-710190a609d9','stage','Stage','924a4b3d-aa54-452d-955d-d5763b916c03', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9cf0ae8d-fb80-4717-81d8-ec19d32b64c4','precipitation','Precipitation','924a4b3d-aa54-452d-955d-d5763b916c03', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6efc5e51-6648-4652-b022-99f8535ef4b4','voltage','Voltage','924a4b3d-aa54-452d-955d-d5763b916c03', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('982b65eb-0d31-46c2-9c23-396f2105ee70','stage','Stage','0f2a5b22-363e-472a-ab11-8aabd9779af2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('123cc246-d8c0-4673-8a99-cf5279011c9f','precipitation','Precipitation','0f2a5b22-363e-472a-ab11-8aabd9779af2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('44e20efd-5a32-486a-8b23-6083e9435e0b','voltage','Voltage','0f2a5b22-363e-472a-ab11-8aabd9779af2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('57891e19-43c1-4387-a791-e34a1075e782','stage','Stage','420d5ecc-20f5-46cb-bbfc-f9ec08874283', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('77456a59-13b5-4500-bad4-a75fa4924b30','voltage','Voltage','420d5ecc-20f5-46cb-bbfc-f9ec08874283', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('f59b59f0-5422-4be9-8118-a463c4ffb178','precipitation','Precipitation','8ae11ab9-2efa-410f-ac1c-ce1eb4b9e33f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('5fa4285b-e491-4caa-9877-b75a275a8a9c','voltage','Voltage','8ae11ab9-2efa-410f-ac1c-ce1eb4b9e33f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('b6f5197e-b7cd-443b-b62d-361eaf588545','stage','Stage','fd46c733-b662-4195-b96d-0369df362a11', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3bf04445-d993-474a-836a-06dcbd12550f','voltage','Voltage','fd46c733-b662-4195-b96d-0369df362a11', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('b555d7b1-785a-4860-924b-0f741bcb2b8b','precipitation','Precipitation','8f7dc916-885e-43dd-850a-509741483b80', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('6945ea70-0792-4c60-b3a6-293c474a69ba','voltage','Voltage','8f7dc916-885e-43dd-850a-509741483b80', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ef98dfd0-ee9d-4a0f-bde2-5f60ebfc65e7','precipitation','Precipitation','f956ed57-4c11-45aa-ba97-137fc08f26f2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('86356704-c6c1-447f-8aaa-7b3c18d2fc99','stage','Stage','f956ed57-4c11-45aa-ba97-137fc08f26f2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ce621e35-4463-4f62-bb0f-a5487f810670','voltage','Voltage','f956ed57-4c11-45aa-ba97-137fc08f26f2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d9eb0611-1642-4679-a9c9-c184807fe43f','precipitation','Precipitation','7a1d3531-440f-4a13-bc32-f43101ebac4d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('280f43db-be3b-46e0-970c-7c5c3d6b5da3','voltage','Voltage','7a1d3531-440f-4a13-bc32-f43101ebac4d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6cc461bf-1d21-4934-a6ea-b9f4e51dc2a4','stage','Stage','0247a7d6-fb7f-4968-b6dc-c2e8f277dab6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('06eae943-4b5c-4f39-82c2-77df6301c8cb','voltage','Voltage','0247a7d6-fb7f-4968-b6dc-c2e8f277dab6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1935b5ce-437d-4a3c-b703-44fd21b6fe5e','stage','Stage','195be647-588b-4973-bf02-8717a31d02e1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('cb67d36d-8ced-4e8c-9db1-bcbdd6ead608','voltage','Voltage','195be647-588b-4973-bf02-8717a31d02e1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('08451627-ed1b-4f0d-93f4-50c22fc161e5','stage','Stage','c2c43566-82eb-4901-a1d3-81f2200e193e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3591ba30-0a2a-4dfa-ba27-175dd1fe5a15','voltage','Voltage','c2c43566-82eb-4901-a1d3-81f2200e193e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0d6c156d-9c59-4308-9199-b6c9870a55fd','stage','Stage','f8a51a89-74ce-4c03-bc93-594c97b7de00', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dc561bbf-2cd3-4b27-9715-5e7f33a8bf85','voltage','Voltage','f8a51a89-74ce-4c03-bc93-594c97b7de00', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('932c86f4-7c86-4761-b206-bc426f3ff67a','stage','Stage','34b223e2-5394-466b-b7ac-9d84788a4db3', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4a771a27-7f2c-4dab-ba64-6eaa1c85aa76','precipitation','Precipitation','34b223e2-5394-466b-b7ac-9d84788a4db3', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c274f372-df57-4fc3-854e-ed1c297911f8','voltage','Voltage','34b223e2-5394-466b-b7ac-9d84788a4db3', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6b973991-c731-4a24-8551-19fc62f69828','stage','Stage','01de28df-bf75-428c-ae1d-f4feb975db63', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1671d089-9fde-49dd-a70a-8cd8ff607e28','voltage','Voltage','01de28df-bf75-428c-ae1d-f4feb975db63', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('16cd87a0-95ed-4459-84ae-aa724c3a0cb6','stage','Stage','ace2e8fb-6357-4bc2-90be-b47e426caf1d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('abf6f54a-c26a-4572-8ff1-09020a46f50a','precipitation','Precipitation','ace2e8fb-6357-4bc2-90be-b47e426caf1d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('c8678751-b40d-4c12-a354-3fe7ac85ae7b','voltage','Voltage','ace2e8fb-6357-4bc2-90be-b47e426caf1d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4136ff4f-9ed3-4151-83a4-b4f12dbaee13','stage','Stage','02675aaf-22fc-426a-bcab-228f285857c1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a1f03172-cc37-402d-8237-934d1db1bf27','voltage','Voltage','02675aaf-22fc-426a-bcab-228f285857c1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('72f667df-6a18-4f8d-a0bf-808f37e753a1','stage','Stage','7c73c2ba-9e4f-4ac4-885b-d51aea328d5f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('96122e8f-a2f2-4210-aecc-cdc04a84274f','voltage','Voltage','7c73c2ba-9e4f-4ac4-885b-d51aea328d5f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('eb93013b-5f9e-4599-9966-cb216cf2922f','stage','Stage','cd09e10e-ab4d-49ba-830d-31e34f8faaa3', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bdcef576-edc7-4d3b-b3e9-e5c162235ac6','voltage','Voltage','cd09e10e-ab4d-49ba-830d-31e34f8faaa3', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64');
