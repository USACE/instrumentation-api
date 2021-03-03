INSERT INTO project (id, office_id, slug, name, image) VALUES
    ('6b60b4e6-ff3a-4b8a-8426-93347824588f', '0154184e-2509-4485-b449-8eff4ab52eef', 'savannah-district-streamgages', 'Savannah District Streamgages', 'savannah-district-streamgages.jpg');




--INSERT INSTRUMENTS--COUNT:49
INSERT INTO public.instrument(id, deleted, slug, name, formula, geometry, station, station_offset, create_date, update_date, type_id, project_id, creator, updater, usgs_id)
 VALUES 
('23655b91-27a8-4f46-90d2-fe92c12c321a', False, 'chattooga-river-at-burrells-ford-nr-pine-mtn-ga', 'CHATTOOGA RIVER AT BURRELLS FORD, NR PINE MTN, GA', null, ST_GeomFromText('POINT(-83.1194 34.9689)',4326), null, null, '2021-03-02T10:00:11.446927Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02176930'),
('d7e9e70e-4612-4904-891a-95899b1e9cbb', False, 'chattooga-river-near-clayton-ga', 'CHATTOOGA RIVER NEAR CLAYTON, GA', null, ST_GeomFromText('POINT(-83.306 34.814)',4326), null, null, '2021-03-02T10:00:17.905120Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02177000'),
('1d6ae3ce-7e9a-4cb3-8aff-efeefd4da85e', False, 'tallulah-river-near-clayton-ga', 'TALLULAH RIVER NEAR CLAYTON, GA', null, ST_GeomFromText('POINT(-83.5304 34.8904)',4326), null, null, '2021-03-02T10:00:22.679199Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02178400'),
('e2acbc5a-c694-4bf8-aa2d-4a6e5cd66a04', False, 'tallulah-river-ab-powerhouse-nr-tallulah-falls-ga', 'TALLULAH RIVER AB POWERHOUSE, NR TALLULAH FALLS,GA', null, ST_GeomFromText('POINT(-83.3757 34.732)',4326), null, null, '2021-03-02T10:00:26.777534Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02181580'),
('d8ee2a9b-baca-4ac2-840f-ee674abc4dd5', False, '02182000', '02182000', null, ST_GeomFromText('POINT(-83.3452 34.6779)',4326), null, null, '2021-03-02T10:00:31.327307Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02182000'),
('b31fcde3-71a6-4c96-9db4-f7b024a159a9', False, 'beaverdam-creek-above-elberton-ga', 'BEAVERDAM CREEK ABOVE ELBERTON, GA', null, ST_GeomFromText('POINT(-82.8965 34.1687)',4326), null, null, '2021-03-02T10:00:34.895412Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02188600'),
('ea5fb959-5791-4150-a7db-b01691f9bae2', False, 'broad-river-above-carlton-ga', 'BROAD RIVER ABOVE CARLTON, GA', null, ST_GeomFromText('POINT(-83.0033 34.0733)',4326), null, null, '2021-03-02T10:00:38.527215Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02191300'),
('4043173a-17bb-47d8-b1b9-507075443cb0', False, 'broad-river-near-bell-ga', 'BROAD RIVER NEAR BELL, GA', null, ST_GeomFromText('POINT(-82.77 33.9742)',4326), null, null, '2021-03-02T10:00:42.314445Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02192000'),
('b360d2a9-7314-42a3-be5d-a69662d69e55', False, 'kettle-creek-near-washington-ga', 'KETTLE CREEK NEAR WASHINGTON, GA', null, ST_GeomFromText('POINT(-82.8579 33.6826)',4326), null, null, '2021-03-02T10:00:45.918801Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193340'),
('461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', False, 'little-river-near-washington-ga', 'LITTLE RIVER NEAR WASHINGTON, GA', null, ST_GeomFromText('POINT(-82.7425 33.6128)',4326), null, null, '2021-03-02T10:00:49.918516Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193500'),
('9f1e159e-0e94-4232-a61e-a426294aac5d', False, 'kiokee-creek-at-ga-104-near-evans-ga', 'KIOKEE CREEK AT GA 104, NEAR EVANS, GA', null, ST_GeomFromText('POINT(-82.2326 33.601)',4326), null, null, '2021-03-02T10:00:53.757909Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02195320'),
('890a95de-c578-419d-849c-7ce1679ec6b8', False, 'butler-creek-below-7th-avenue-at-ft-gordon-ga', 'BUTLER CREEK BELOW 7TH AVENUE, AT FT. GORDON, GA', null, ST_GeomFromText('POINT(-82.1161 33.4386)',4326), null, null, '2021-03-02T10:00:56.790822Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196835'),
('b5b9f77f-94bc-4f48-b1c6-e47b208966fb', False, 'butler-creek-reservoir-at-fort-gordon-ga', 'BUTLER CREEK RESERVOIR AT FORT GORDON, GA', null, ST_GeomFromText('POINT(-82.0992 33.4258)',4326), null, null, '2021-03-02T10:01:00.151848Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196838'),
('69309767-e07c-474a-9ad6-d38d19599dae', False, 'spirit-creek-at-us-1-near-augusta-ga', 'SPIRIT CREEK AT US 1, NEAR AUGUSTA, GA', null, ST_GeomFromText('POINT(-82.139 33.3735)',4326), null, null, '2021-03-02T10:01:04.141360Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197020'),
('98c94154-b54a-44a0-b5ab-bc69127ff297', False, 'savannah-river-near-waynesboro-ga', 'SAVANNAH RIVER NEAR WAYNESBORO, GA', null, ST_GeomFromText('POINT(-81.7548 33.1499)',4326), null, null, '2021-03-02T10:01:06.001359Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021973269'),
('0b941abe-45b7-470f-8c64-1a90f770dc08', False, 'brushy-creek-at-campground-road-near-wrens-ga', 'BRUSHY CREEK AT CAMPGROUND ROAD, NEAR WRENS, GA', null, ST_GeomFromText('POINT(-82.3343 33.1807)',4326), null, null, '2021-03-02T10:01:06.923628Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197598'),
('cce37bfe-98a0-4e16-8b15-2ad8831ca912', False, 'brier-creek-near-waynesboro-ga', 'BRIER CREEK NEAR WAYNESBORO, GA', null, ST_GeomFromText('POINT(-81.9637 33.1182)',4326), null, null, '2021-03-02T10:01:07.800029Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197830'),
('eaea76f8-0ef4-44f5-ad78-7c7286facf98', False, 'brier-creek-at-millhaven-ga', 'BRIER CREEK AT MILLHAVEN, GA', null, ST_GeomFromText('POINT(-81.6512 32.9335)',4326), null, null, '2021-03-02T10:01:08.212182Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198000'),
('9e1b87a5-2c84-45fb-8c52-d46acaef525d', False, 'beaverdam-creek-near-sardis-ga', 'BEAVERDAM CREEK NEAR SARDIS, GA', null, ST_GeomFromText('POINT(-81.8154 32.9377)',4326), null, null, '2021-03-02T10:01:08.580577Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198100'),
('8ed22f3a-2490-4a4b-afb9-7afd853ff491', False, 'ebenezer-creek-at-springfield-ga', 'EBENEZER CREEK AT SPRINGFIELD, GA', null, ST_GeomFromText('POINT(-81.2973 32.3657)',4326), null, null, '2021-03-02T10:01:09.408861Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198690'),
('b6217f54-1741-4a6c-93d0-547f9cdec38c', False, 'withlacoochee-river-at-us-41-near-valdosta-ga', 'WITHLACOOCHEE RIVER AT US 41 NEAR VALDOSTA GA', null, ST_GeomFromText('POINT(-81.1615 32.3532)',4326), null, null, '2021-03-02T10:01:14.682948Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198745'),
('07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', False, 'yellow-river-at-milstead-ga', 'YELLOW RIVER AT MILSTEAD GA', null, ST_GeomFromText('POINT(-81.1464 32.3003)',4326), null, null, '2021-03-02T10:01:18.989543Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198768'),
('402a8b7e-688a-4fca-9b1d-31fc77391bd0', False, 'abercorn-creek-near-savannah-ga', 'ABERCORN CREEK NEAR SAVANNAH,GA', null, ST_GeomFromText('POINT(-81.1782 32.2558)',4326), null, null, '2021-03-02T10:01:22.639113Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198810'),
('62454df9-d25f-4d7a-8647-948f7c8393f4', False, 'savannah-river-i-95-near-port-wentworth-ga', 'SAVANNAH RIVER (I-95) NEAR PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1512 32.2358)',4326), null, null, '2021-03-02T10:01:26.889438Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198840'),
('47be997a-168a-4f2d-a860-8b95476d0f09', False, 'savannah-river-at-ga-25-at-port-wentworth-ga', 'SAVANNAH RIVER AT GA 25, AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1537 32.166)',4326), null, null, '2021-03-02T10:01:32.079248Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198920'),
('c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', False, 'middle-river-at-ga-25-at-port-wentworth-ga', 'MIDDLE RIVER AT GA 25 AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.1383 32.1656)',4326), null, null, '2021-03-02T10:01:38.163400Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198950'),
('a26f2105-23e4-4331-ba82-d3c5d9e777f5', False, 'savannah-river-at-usace-dock-at-savannah-ga', 'SAVANNAH RIVER AT USACE DOCK, AT SAVANNAH, GA', null, ST_GeomFromText('POINT(-81.0812 32.081)',4326), null, null, '2021-03-02T10:01:41.804256Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989773'),
('79b365ba-f543-456e-bf40-01dd7d515bcf', False, 'l-back-river-above-lucknow-canal-nr-limehouse-sc', 'L BACK RIVER ABOVE LUCKNOW CANAL, NR LIMEHOUSE, SC', null, ST_GeomFromText('POINT(-81.1179 32.1858)',4326), null, null, '2021-03-02T10:01:45.084800Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989784'),
('c2034ba0-a672-42ce-87ef-1f12b2faa79a', False, 'little-back-river-at-f&w-dock-near-limehouse-sc', 'LITTLE BACK RIVER AT F&W DOCK, NEAR LIMEHOUSE, SC', null, ST_GeomFromText('POINT(-81.1182 32.1708)',4326), null, null, '2021-03-02T10:01:48.993415Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989791'),
('f539acc2-c718-45e6-892a-d6d150317e6f', False, 'little-back-river-at-ga-25-at-port-wentworth-ga', 'LITTLE BACK RIVER AT GA 25 AT PORT WENTWORTH, GA', null, ST_GeomFromText('POINT(-81.13 32.1658)',4326), null, null, '2021-03-02T10:01:53.665702Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021989792'),
('d8e7ebe1-2e78-4f57-82b4-53a3432c81cc', False, 'savannah-river-at-fort-pulaski-ga', 'SAVANNAH RIVER AT FORT PULASKI, GA', null, ST_GeomFromText('POINT(-80.9032 32.0341)',4326), null, null, '2021-03-02T10:01:59.023388Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198980'),
('97fd6c25-af34-4579-b2a6-9141d03b13de', False, 'south-channel-savannah-river-near-savannah-ga', 'SOUTH CHANNEL (SAVANNAH RIVER) NEAR SAVANNAH, GA', null, ST_GeomFromText('POINT(-81.0023 32.0827)',4326), null, null, '2021-03-02T10:02:08.909190Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02199000'),
('f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', False, 'twelvemile-creek-near-liberty-sc', 'TWELVEMILE CREEK NEAR LIBERTY, SC', null, ST_GeomFromText('POINT(-82.7485 34.8015)',4326), null, null, '2021-03-02T10:02:09.306985Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02186000'),
('2a6b30d6-d2eb-4c9a-b35f-1285108455ef', False, 'eighteenmile-creek-above-pendleton-sc', 'EIGHTEENMILE CREEK ABOVE PENDLETON, SC', null, ST_GeomFromText('POINT(-82.7988 34.659)',4326), null, null, '2021-03-02T10:02:09.649935Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02186699'),
('91da2bf4-ffd7-43bf-a8e7-f0dd04311a58', False, 'hartwell-lake-near-anderson-sc', 'HARTWELL LAKE NEAR ANDERSON, SC', null, ST_GeomFromText('POINT(-82.8161 34.475)',4326), null, null, '2021-03-02T10:02:09.991547Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02187010'),
('4d6d85d1-c779-42ad-9cb1-3cafc0a0bcd5', False, 'rocky-river-nr-starr-sc', 'ROCKY RIVER NR STARR, SC', null, ST_GeomFromText('POINT(-82.5774 34.3832)',4326), null, null, '2021-03-02T10:02:10.322939Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02187910'),
('d2a58327-b9bf-4862-915c-dcc82878ea05', False, 'russell-lake-above-calhoun-falls-sc', 'RUSSELL LAKE ABOVE CALHOUN FALLS, SC', null, ST_GeomFromText('POINT(-82.6181 34.1011)',4326), null, null, '2021-03-02T10:02:10.661804Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02188100'),
('269bc11e-2bf1-4d98-86b2-977f75e4e1fb', False, 'little-river-near-mt-carmel-sc', 'LITTLE RIVER NEAR MT. CARMEL, SC', null, ST_GeomFromText('POINT(-82.5007 34.0715)',4326), null, null, '2021-03-02T10:02:11.012240Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02192500'),
('798a93e4-0fd7-402f-8c2b-65a8a0850ea2', False, 'thurmond-lake-near-plum-branch-sc', 'THURMOND LAKE NEAR PLUM BRANCH, SC', null, ST_GeomFromText('POINT(-82.3506 33.8403)',4326), null, null, '2021-03-02T10:02:11.360158Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02193900'),
('f2c66c58-3f34-4f92-a46d-848c62764499', False, 'savannah-river-near-evans-ga', 'SAVANNAH RIVER NEAR EVANS, GA', null, ST_GeomFromText('POINT(-82.1232 33.5929)',4326), null, null, '2021-03-02T10:02:11.696916Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02195520'),
('91b3ca6b-8fab-40b3-8911-35beb879f823', False, 'stevens-creek-near-modoc-sc', 'STEVENS CREEK NEAR MODOC, SC', null, ST_GeomFromText('POINT(-82.1818 33.7293)',4326), null, null, '2021-03-02T10:02:12.115395Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196000'),
('aa4f72e5-1fea-4ae6-b7cb-8a7b05579120', False, 'savannah-rvr-at-stevens-creek-dam-nr-morgana-sc', 'SAVANNAH RVR AT STEVENS CREEK DAM NR MORGANA, SC', null, ST_GeomFromText('POINT(-82.051 33.5629)',4326), null, null, '2021-03-02T10:02:12.453651Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196483'),
('fbcac5af-25a7-4e23-be64-ecf18c9fe23e', False, 'augusta-canal-nr-augusta-ga-upper', 'AUGUSTA CANAL NR AUGUSTA, GA (UPPER)', null, ST_GeomFromText('POINT(-82.0379 33.5493)',4326), null, null, '2021-03-02T10:02:12.806459Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196485'),
('abd86a4b-db70-4aac-b397-96fa6fee821a', False, 'horse-creek-at-clearwater-sc', 'HORSE CREEK AT CLEARWATER, SC', null, ST_GeomFromText('POINT(-81.8971 33.4849)',4326), null, null, '2021-03-02T10:02:13.167706Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02196690'),
('c2013f5c-ac88-44cb-a33e-f0ed8d2d8109', False, '021970161', '021970161', null, ST_GeomFromText('POINT(-82.1572 33.39)',4326), null, null, '2021-03-02T10:02:13.537459Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '021970161'),
('cc927dcf-5a23-460b-b7ec-dc731374edaa', False, 'savannah-r-at-burtons-ferry-br-nr-millhaven-ga', 'SAVANNAH R AT BURTONS FERRY BR NR MILLHAVEN, GA', null, ST_GeomFromText('POINT(-81.5026 32.939)',4326), null, null, '2021-03-02T10:02:13.905506Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02197500'),
('87d871a1-3a41-4442-b61d-bd65e3ac7b8d', False, 'savannah-river-near-estill-sc', 'SAVANNAH RIVER NEAR ESTILL, SC', null, ST_GeomFromText('POINT(-81.4283 32.7033)',4326), null, null, '2021-03-02T10:02:14.240268Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198375'),
('47b3ed3c-bee7-427f-9174-4c0bc84a5685', False, 'savannah-river-near-clyo-ga', 'SAVANNAH RIVER NEAR CLYO, GA', null, ST_GeomFromText('POINT(-81.2687 32.5282)',4326), null, null, '2021-03-02T10:02:20.831574Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198500'),
('82fc7100-670d-4076-9292-5e4020aa9478', False, 'savannah-river-above-hardeeville-sc', 'SAVANNAH RIVER ABOVE HARDEEVILLE, SC', null, ST_GeomFromText('POINT(-81.1284 32.3394)',4326), null, null, '2021-03-02T10:02:25.640795Z', null, '98a61f29-18a8-430a-9d02-0f53486e0984', '6b60b4e6-ff3a-4b8a-8426-93347824588f', '00000000-0000-0000-0000-000000000000', null, '02198760');

--INSERT INSTRUMENT STATUS--
INSERT INTO public.instrument_status(id, instrument_id, status_id, "time")
 VALUES 
('63e8256d-bfe6-4714-bbd9-ec42614191dc', '23655b91-27a8-4f46-90d2-fe92c12c321a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:11.446927Z'),
('a2623cb3-f847-4ade-b533-91792b2a3b8e', 'd7e9e70e-4612-4904-891a-95899b1e9cbb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:17.905120Z'),
('66efc2ad-6e71-46c9-9040-2349f86ee1c3', '1d6ae3ce-7e9a-4cb3-8aff-efeefd4da85e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:22.679199Z'),
('2e6123a4-ddab-47ae-9d51-375693ddc8e8', 'e2acbc5a-c694-4bf8-aa2d-4a6e5cd66a04', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:26.777534Z'),
('be01eb9a-2a0d-47d4-826a-49fd42946080', 'd8ee2a9b-baca-4ac2-840f-ee674abc4dd5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:31.327307Z'),
('0341c8ab-9ada-427f-97a5-f89847a32f24', 'b31fcde3-71a6-4c96-9db4-f7b024a159a9', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:34.895412Z'),
('a851575c-1f9c-4d95-a535-38ec37b42b6c', 'ea5fb959-5791-4150-a7db-b01691f9bae2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:38.527215Z'),
('1393bf25-0615-41cb-b4a3-9d074af8036d', '4043173a-17bb-47d8-b1b9-507075443cb0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:42.314445Z'),
('f5548011-dd8a-4bfd-8761-010b4f48075a', 'b360d2a9-7314-42a3-be5d-a69662d69e55', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:45.918801Z'),
('9e98d029-05ec-4858-abc9-3dcd8e81ce5a', '461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:49.918516Z'),
('5f5951e0-2368-4670-a28b-f05a7714ceda', '9f1e159e-0e94-4232-a61e-a426294aac5d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:53.757909Z'),
('acaa0e82-f01d-4edb-bd96-96939bca9615', '890a95de-c578-419d-849c-7ce1679ec6b8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:00:56.790822Z'),
('fd53303d-fe7b-4064-a73d-c037135a52cd', 'b5b9f77f-94bc-4f48-b1c6-e47b208966fb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:00.151848Z'),
('36b8d316-4314-4a42-92b8-2477874469c9', '69309767-e07c-474a-9ad6-d38d19599dae', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:04.141360Z'),
('f78ebb70-03be-4c2e-9eed-6613925829e7', '98c94154-b54a-44a0-b5ab-bc69127ff297', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:06.001359Z'),
('cf075470-2747-4694-9430-e3729453c1b8', '0b941abe-45b7-470f-8c64-1a90f770dc08', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:06.923628Z'),
('c78616da-c118-4482-9a42-9af889027b11', 'cce37bfe-98a0-4e16-8b15-2ad8831ca912', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:07.800029Z'),
('48f38d1a-f4c4-443f-a013-890061178717', 'eaea76f8-0ef4-44f5-ad78-7c7286facf98', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:08.212182Z'),
('3181932b-5d2d-4db4-a082-eaa712d5fb80', '9e1b87a5-2c84-45fb-8c52-d46acaef525d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:08.580577Z'),
('a72a05d2-8af2-4603-84e8-baf5caa73a75', '8ed22f3a-2490-4a4b-afb9-7afd853ff491', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:09.408861Z'),
('e259dfcb-6b8d-491f-86d6-8928aa1cf403', 'b6217f54-1741-4a6c-93d0-547f9cdec38c', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:14.682948Z'),
('216f5fa7-6307-4ac5-bb24-d44f821a2c5a', '07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:18.989543Z'),
('05a18940-49bd-4a0c-90e7-325c72608336', '402a8b7e-688a-4fca-9b1d-31fc77391bd0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:22.639113Z'),
('ca7ea8c1-e255-42a8-a5ba-8bbdd5c8c8c8', '62454df9-d25f-4d7a-8647-948f7c8393f4', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:26.889438Z'),
('c5d9adfb-c1a4-4925-a10b-6ecb315852d1', '47be997a-168a-4f2d-a860-8b95476d0f09', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:32.079248Z'),
('4c59d7fc-6630-4d37-b2e6-0d48daf3b1c5', 'c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:38.163400Z'),
('10986461-71d6-4283-bc14-aa5b9a16330b', 'a26f2105-23e4-4331-ba82-d3c5d9e777f5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:41.804256Z'),
('29035fc8-50b8-4982-94f0-e7d8a3a9fd16', '79b365ba-f543-456e-bf40-01dd7d515bcf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:45.084800Z'),
('486f16ae-8f76-46fd-8a75-77dae8d94d8b', 'c2034ba0-a672-42ce-87ef-1f12b2faa79a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:48.993415Z'),
('2144935b-9126-492c-8e14-4b286614dc98', 'f539acc2-c718-45e6-892a-d6d150317e6f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:53.665702Z'),
('db191246-83e5-4b67-add8-16dbbc00f049', 'd8e7ebe1-2e78-4f57-82b4-53a3432c81cc', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:01:59.023388Z'),
('b902dcc9-95c2-4faa-b8a9-e27d56cb80ee', '97fd6c25-af34-4579-b2a6-9141d03b13de', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:08.909190Z'),
('6a70b89a-ffb7-4f5e-9a1d-ee69e939c9d2', 'f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:09.306985Z'),
('32a8e1ac-8944-4bd9-9f79-997b8d0876d5', '2a6b30d6-d2eb-4c9a-b35f-1285108455ef', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:09.649935Z'),
('b4d43827-e609-4105-a3a4-7a1c4cbde59d', '91da2bf4-ffd7-43bf-a8e7-f0dd04311a58', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:09.991547Z'),
('38450cbc-76c1-442a-a4a8-c05b3bd22d1c', '4d6d85d1-c779-42ad-9cb1-3cafc0a0bcd5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:10.322939Z'),
('f63cbe18-8329-4f26-b204-32390333901b', 'd2a58327-b9bf-4862-915c-dcc82878ea05', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:10.661804Z'),
('bdef30a4-53c5-442d-bc7b-e3703af56a37', '269bc11e-2bf1-4d98-86b2-977f75e4e1fb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:11.012240Z'),
('0f144286-952f-4e0c-9914-98a441fe114f', '798a93e4-0fd7-402f-8c2b-65a8a0850ea2', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:11.360158Z'),
('7bb5c2b3-e228-41e6-9a7d-ee60cc709a8e', 'f2c66c58-3f34-4f92-a46d-848c62764499', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:11.696916Z'),
('aa6248fa-c93d-492d-a323-d92e942842bf', '91b3ca6b-8fab-40b3-8911-35beb879f823', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:12.115395Z'),
('259696a3-a9d7-421e-ad92-1eea66bab85d', 'aa4f72e5-1fea-4ae6-b7cb-8a7b05579120', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:12.453651Z'),
('7b719d77-1cc8-42ed-8fd8-6194d82f1342', 'fbcac5af-25a7-4e23-be64-ecf18c9fe23e', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:12.806459Z'),
('4f271d83-0207-45ec-ad9d-77b17d9b4ade', 'abd86a4b-db70-4aac-b397-96fa6fee821a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:13.167706Z'),
('e6d10972-ff51-4e5d-93cb-bf3c6d3fe0a1', 'c2013f5c-ac88-44cb-a33e-f0ed8d2d8109', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:13.537459Z'),
('2ac02dc6-2aae-4d81-8e8e-005840d45f61', 'cc927dcf-5a23-460b-b7ec-dc731374edaa', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:13.905506Z'),
('13d2b2f6-ab94-492e-8a0f-a72ba76bc27c', '87d871a1-3a41-4442-b61d-bd65e3ac7b8d', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:14.240268Z'),
('26ef2c97-ed2e-46c0-ba67-de2bb6b6e9c8', '47b3ed3c-bee7-427f-9174-4c0bc84a5685', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:20.831574Z'),
('19682b42-ac0a-4629-8b7d-b00309d296d4', '82fc7100-670d-4076-9292-5e4020aa9478', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', '2021-03-02T10:02:25.640795Z');

--INSERT TELEMETRY_GOES--COUNT:43
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b2b959d6-94ab-43ed-8e55-c24371f3eb5d', 'DE203122' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE203122');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a9ab7c07-45ac-4e3a-aa53-a794f9d03b35', '17ACC13E' where not exists (select 1 from telemetry_goes where nesdis_id = '17ACC13E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd3100aae-0725-4b17-a823-d658be5273bb', 'DE25E08A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE25E08A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd3695191-a29b-4d34-8abd-87d72da1f51c', 'CE8771CE' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8771CE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5d4d38a3-e69e-45bb-9934-fb59c3611365', 'DE15F58C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE15F58C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'db7ee394-3ba2-4338-8067-a43ed035fd5a', 'DD8D036E' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD8D036E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a98d6743-c3f1-48f2-81d5-04b0f16e0289', 'DE160206' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE160206');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8c82bc8a-e36d-4acf-b86c-a46ed4c3e2a2', '17E731CE' where not exists (select 1 from telemetry_goes where nesdis_id = '17E731CE');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a8f86f1a-399c-4ea0-b74e-a6e0a42405d9', 'CE8721B2' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8721B2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '26670a31-116b-429f-97fe-5dce9aa31d48', '17F2177C' where not exists (select 1 from telemetry_goes where nesdis_id = '17F2177C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'ee9e3d83-1597-4380-9be5-73b17ce55e96', 'DD22D5C2' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD22D5C2');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a15e367b-3af6-4f93-b412-155b5ce2ffed', '17E7475E' where not exists (select 1 from telemetry_goes where nesdis_id = '17E7475E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '425f953d-3f72-409e-b92f-f18a0ff28d80', 'DD64F03C' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD64F03C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'f93872b1-d17f-4ac5-a0a1-979ea66c21a0', 'DDB645BC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDB645BC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '99d8146d-4a18-4a95-a1ae-2f0458ba1146', 'DDB0A680' where not exists (select 1 from telemetry_goes where nesdis_id = 'DDB0A680');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'b94e9540-6dd7-4271-b272-3e61c39dc2b2', 'DD49F2B4' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD49F2B4');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '08089539-1354-4c87-8df5-85ae2e345e99', 'DD4A1648' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD4A1648');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5f46734e-7825-4c08-800d-6ad9eaecae10', '17B5F566' where not exists (select 1 from telemetry_goes where nesdis_id = '17B5F566');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '07c722dd-b899-4740-a468-6130b6876204', '17B6670A' where not exists (select 1 from telemetry_goes where nesdis_id = '17B6670A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '0dfb5a0e-cefe-4ab2-a4df-c56fd45f77f9', '17456210' where not exists (select 1 from telemetry_goes where nesdis_id = '17456210');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1dc1a065-be0e-4b6a-9ae4-bc9884e502dc', '17D33284' where not exists (select 1 from telemetry_goes where nesdis_id = '17D33284');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5ff571dc-0067-4259-a0af-2e20c89b4fa8', 'DD24C67A' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD24C67A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'edf5a0c4-dd6c-47e6-97b6-fe5aa9039432', '17B71360' where not exists (select 1 from telemetry_goes where nesdis_id = '17B71360');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a7c9e9dc-dc36-4c04-8ad7-9b057b1f0e68', '17CDA410' where not exists (select 1 from telemetry_goes where nesdis_id = '17CDA410');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '33165240-bb20-4653-afe7-a5af913f322b', '17B6E11E' where not exists (select 1 from telemetry_goes where nesdis_id = '17B6E11E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '5ebdb2ea-3558-46ee-a061-0c02f94ad9ad', '17CD610E' where not exists (select 1 from telemetry_goes where nesdis_id = '17CD610E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '95907a9c-c9e0-473d-82d4-12a6f865f447', 'CE87923C' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87923C');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '8bb80ee8-3002-464d-8699-f43100001480', '17055456' where not exists (select 1 from telemetry_goes where nesdis_id = '17055456');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'de142918-406a-4d71-b4a1-8ebc4692d332', 'CE87814A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87814A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '4e75f470-725d-4691-8211-5ae4965cf8b6', 'CE8762B8' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE8762B8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'd4a341d4-fa02-4afb-8d5e-58375fdcee69', 'CE875722' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE875722');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '72ad80e6-2131-4192-9c54-c0914d6bdb6b', 'CE874454' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE874454');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '702d6833-6478-4931-9496-fad55dad3633', 'CE871428' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE871428');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'c698d9af-fa82-4947-8e82-54fecfa7b1aa', 'CE87075E' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87075E');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '7e753661-8655-434b-8056-08e9379511ea', 'DD2CC2DC' where not exists (select 1 from telemetry_goes where nesdis_id = 'DD2CC2DC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a32928fc-40c6-45fb-9d93-d2a2a016e678', 'CE87E4AC' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87E4AC');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'fbc36b30-55e1-40e8-a84d-f579cd155f45', '17EAE4C0' where not exists (select 1 from telemetry_goes where nesdis_id = '17EAE4C0');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'a0bce16c-16d6-4073-95ba-feb0b17e6c78', 'CE87F7DA' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE87F7DA');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '1d318552-62eb-424f-9235-fb61711c860f', 'DE25D510' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE25D510');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '12bd1f12-33a2-45de-94ad-371c1aefebb2', 'CE86B62A' where not exists (select 1 from telemetry_goes where nesdis_id = 'CE86B62A');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '6ffe1a8b-70a3-44ea-9249-9301bb44d628', 'DE0CA432' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE0CA432');
INSERT INTO public.telemetry_goes (id, nesdis_id) select 'e2c9b442-f7e1-4e80-871c-66948f4747ba', 'D11C04F8' where not exists (select 1 from telemetry_goes where nesdis_id = 'D11C04F8');
INSERT INTO public.telemetry_goes (id, nesdis_id) select '969ee6ac-78b0-47e2-b205-3f2db200c5fb', 'DE0CD2A2' where not exists (select 1 from telemetry_goes where nesdis_id = 'DE0CD2A2');

--INSERT INSTRUMENT_TELEMETRY--COUNT:43
INSERT INTO public.instrument_telemetry (instrument_id, telemetry_type_id, telemetry_id) 
VALUES
('23655b91-27a8-4f46-90d2-fe92c12c321a', '10a32652-af43-4451-bd52-4980c5690cc9', 'b2b959d6-94ab-43ed-8e55-c24371f3eb5d'),
('d7e9e70e-4612-4904-891a-95899b1e9cbb', '10a32652-af43-4451-bd52-4980c5690cc9', 'a9ab7c07-45ac-4e3a-aa53-a794f9d03b35'),
('d8ee2a9b-baca-4ac2-840f-ee674abc4dd5', '10a32652-af43-4451-bd52-4980c5690cc9', 'd3100aae-0725-4b17-a823-d658be5273bb'),
('b31fcde3-71a6-4c96-9db4-f7b024a159a9', '10a32652-af43-4451-bd52-4980c5690cc9', 'd3695191-a29b-4d34-8abd-87d72da1f51c'),
('ea5fb959-5791-4150-a7db-b01691f9bae2', '10a32652-af43-4451-bd52-4980c5690cc9', '5d4d38a3-e69e-45bb-9934-fb59c3611365'),
('4043173a-17bb-47d8-b1b9-507075443cb0', '10a32652-af43-4451-bd52-4980c5690cc9', 'db7ee394-3ba2-4338-8067-a43ed035fd5a'),
('b360d2a9-7314-42a3-be5d-a69662d69e55', '10a32652-af43-4451-bd52-4980c5690cc9', 'a98d6743-c3f1-48f2-81d5-04b0f16e0289'),
('461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', '10a32652-af43-4451-bd52-4980c5690cc9', '8c82bc8a-e36d-4acf-b86c-a46ed4c3e2a2'),
('9f1e159e-0e94-4232-a61e-a426294aac5d', '10a32652-af43-4451-bd52-4980c5690cc9', 'a8f86f1a-399c-4ea0-b74e-a6e0a42405d9'),
('890a95de-c578-419d-849c-7ce1679ec6b8', '10a32652-af43-4451-bd52-4980c5690cc9', '26670a31-116b-429f-97fe-5dce9aa31d48'),
('b5b9f77f-94bc-4f48-b1c6-e47b208966fb', '10a32652-af43-4451-bd52-4980c5690cc9', 'ee9e3d83-1597-4380-9be5-73b17ce55e96'),
('69309767-e07c-474a-9ad6-d38d19599dae', '10a32652-af43-4451-bd52-4980c5690cc9', 'a15e367b-3af6-4f93-b412-155b5ce2ffed'),
('98c94154-b54a-44a0-b5ab-bc69127ff297', '10a32652-af43-4451-bd52-4980c5690cc9', '425f953d-3f72-409e-b92f-f18a0ff28d80'),
('0b941abe-45b7-470f-8c64-1a90f770dc08', '10a32652-af43-4451-bd52-4980c5690cc9', 'f93872b1-d17f-4ac5-a0a1-979ea66c21a0'),
('eaea76f8-0ef4-44f5-ad78-7c7286facf98', '10a32652-af43-4451-bd52-4980c5690cc9', '99d8146d-4a18-4a95-a1ae-2f0458ba1146'),
('9e1b87a5-2c84-45fb-8c52-d46acaef525d', '10a32652-af43-4451-bd52-4980c5690cc9', 'b94e9540-6dd7-4271-b272-3e61c39dc2b2'),
('8ed22f3a-2490-4a4b-afb9-7afd853ff491', '10a32652-af43-4451-bd52-4980c5690cc9', '08089539-1354-4c87-8df5-85ae2e345e99'),
('b6217f54-1741-4a6c-93d0-547f9cdec38c', '10a32652-af43-4451-bd52-4980c5690cc9', '5f46734e-7825-4c08-800d-6ad9eaecae10'),
('07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', '10a32652-af43-4451-bd52-4980c5690cc9', '07c722dd-b899-4740-a468-6130b6876204'),
('402a8b7e-688a-4fca-9b1d-31fc77391bd0', '10a32652-af43-4451-bd52-4980c5690cc9', '0dfb5a0e-cefe-4ab2-a4df-c56fd45f77f9'),
('62454df9-d25f-4d7a-8647-948f7c8393f4', '10a32652-af43-4451-bd52-4980c5690cc9', '1dc1a065-be0e-4b6a-9ae4-bc9884e502dc'),
('47be997a-168a-4f2d-a860-8b95476d0f09', '10a32652-af43-4451-bd52-4980c5690cc9', '5ff571dc-0067-4259-a0af-2e20c89b4fa8'),
('c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', '10a32652-af43-4451-bd52-4980c5690cc9', 'edf5a0c4-dd6c-47e6-97b6-fe5aa9039432'),
('a26f2105-23e4-4331-ba82-d3c5d9e777f5', '10a32652-af43-4451-bd52-4980c5690cc9', 'a7c9e9dc-dc36-4c04-8ad7-9b057b1f0e68'),
('f539acc2-c718-45e6-892a-d6d150317e6f', '10a32652-af43-4451-bd52-4980c5690cc9', '33165240-bb20-4653-afe7-a5af913f322b'),
('97fd6c25-af34-4579-b2a6-9141d03b13de', '10a32652-af43-4451-bd52-4980c5690cc9', '5ebdb2ea-3558-46ee-a061-0c02f94ad9ad'),
('f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', '10a32652-af43-4451-bd52-4980c5690cc9', '95907a9c-c9e0-473d-82d4-12a6f865f447'),
('2a6b30d6-d2eb-4c9a-b35f-1285108455ef', '10a32652-af43-4451-bd52-4980c5690cc9', '8bb80ee8-3002-464d-8699-f43100001480'),
('91da2bf4-ffd7-43bf-a8e7-f0dd04311a58', '10a32652-af43-4451-bd52-4980c5690cc9', 'de142918-406a-4d71-b4a1-8ebc4692d332'),
('4d6d85d1-c779-42ad-9cb1-3cafc0a0bcd5', '10a32652-af43-4451-bd52-4980c5690cc9', '4e75f470-725d-4691-8211-5ae4965cf8b6'),
('d2a58327-b9bf-4862-915c-dcc82878ea05', '10a32652-af43-4451-bd52-4980c5690cc9', 'd4a341d4-fa02-4afb-8d5e-58375fdcee69'),
('269bc11e-2bf1-4d98-86b2-977f75e4e1fb', '10a32652-af43-4451-bd52-4980c5690cc9', '72ad80e6-2131-4192-9c54-c0914d6bdb6b'),
('798a93e4-0fd7-402f-8c2b-65a8a0850ea2', '10a32652-af43-4451-bd52-4980c5690cc9', '702d6833-6478-4931-9496-fad55dad3633'),
('f2c66c58-3f34-4f92-a46d-848c62764499', '10a32652-af43-4451-bd52-4980c5690cc9', 'c698d9af-fa82-4947-8e82-54fecfa7b1aa'),
('91b3ca6b-8fab-40b3-8911-35beb879f823', '10a32652-af43-4451-bd52-4980c5690cc9', '7e753661-8655-434b-8056-08e9379511ea'),
('aa4f72e5-1fea-4ae6-b7cb-8a7b05579120', '10a32652-af43-4451-bd52-4980c5690cc9', 'a32928fc-40c6-45fb-9d93-d2a2a016e678'),
('fbcac5af-25a7-4e23-be64-ecf18c9fe23e', '10a32652-af43-4451-bd52-4980c5690cc9', 'fbc36b30-55e1-40e8-a84d-f579cd155f45'),
('abd86a4b-db70-4aac-b397-96fa6fee821a', '10a32652-af43-4451-bd52-4980c5690cc9', 'a0bce16c-16d6-4073-95ba-feb0b17e6c78'),
('c2013f5c-ac88-44cb-a33e-f0ed8d2d8109', '10a32652-af43-4451-bd52-4980c5690cc9', '1d318552-62eb-424f-9235-fb61711c860f'),
('cc927dcf-5a23-460b-b7ec-dc731374edaa', '10a32652-af43-4451-bd52-4980c5690cc9', '12bd1f12-33a2-45de-94ad-371c1aefebb2'),
('87d871a1-3a41-4442-b61d-bd65e3ac7b8d', '10a32652-af43-4451-bd52-4980c5690cc9', '6ffe1a8b-70a3-44ea-9249-9301bb44d628'),
('47b3ed3c-bee7-427f-9174-4c0bc84a5685', '10a32652-af43-4451-bd52-4980c5690cc9', 'e2c9b442-f7e1-4e80-871c-66948f4747ba'),
('82fc7100-670d-4076-9292-5e4020aa9478', '10a32652-af43-4451-bd52-4980c5690cc9', '969ee6ac-78b0-47e2-b205-3f2db200c5fb');

--INSERT TIMESERIES--COUNT:49
INSERT INTO public.timeseries(id, slug, name, instrument_id, parameter_id, unit_id) 
VALUES
('a37de59e-740a-40e3-9695-56eefd71e912','stage','Stage','23655b91-27a8-4f46-90d2-fe92c12c321a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f71a0d79-5378-4a70-9019-4cb64efacad4','precipitation','Precipitation','23655b91-27a8-4f46-90d2-fe92c12c321a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('68cbb8d7-532c-4b17-8334-9b55fc109f6d','voltage','Voltage','23655b91-27a8-4f46-90d2-fe92c12c321a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c7f78880-8935-4244-b3fc-754fa588c7d0','stage','Stage','d7e9e70e-4612-4904-891a-95899b1e9cbb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5ab04005-b222-441a-82f7-b23a7d14a552','precipitation','Precipitation','d7e9e70e-4612-4904-891a-95899b1e9cbb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('4358d75f-2ed8-4b38-8ded-561124d86956','voltage','Voltage','d7e9e70e-4612-4904-891a-95899b1e9cbb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('a7eb52bf-204e-4c0b-b2e2-2da3e528cd1e','stage','Stage','1d6ae3ce-7e9a-4cb3-8aff-efeefd4da85e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b09c2d02-80ef-4c27-ac93-e496b2a507a1','precipitation','Precipitation','1d6ae3ce-7e9a-4cb3-8aff-efeefd4da85e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('14713c8f-5303-4e0e-a9fc-cd28044c02e5','voltage','Voltage','1d6ae3ce-7e9a-4cb3-8aff-efeefd4da85e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('aa20e709-4dd3-481d-a812-1caee244fd54','stage','Stage','e2acbc5a-c694-4bf8-aa2d-4a6e5cd66a04', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('49f82c8c-3943-45e1-8d0b-d7701d47bcf7','precipitation','Precipitation','e2acbc5a-c694-4bf8-aa2d-4a6e5cd66a04', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('16d45c7e-47bd-4774-872b-ce4063f3306a','voltage','Voltage','e2acbc5a-c694-4bf8-aa2d-4a6e5cd66a04', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8836aecc-530f-42a7-b108-71c4c54da6ae','stage','Stage','d8ee2a9b-baca-4ac2-840f-ee674abc4dd5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('94c9ba1b-cb6d-4132-b3d8-7111fc00f807','precipitation','Precipitation','d8ee2a9b-baca-4ac2-840f-ee674abc4dd5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a77dc559-ee48-4b53-b7bf-17e93b53eea6','voltage','Voltage','d8ee2a9b-baca-4ac2-840f-ee674abc4dd5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('44654e02-a498-4599-9d8c-2b43fbe79ab7','stage','Stage','b31fcde3-71a6-4c96-9db4-f7b024a159a9', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f40b59fe-2f86-44fa-9ef9-2151a039aa10','precipitation','Precipitation','b31fcde3-71a6-4c96-9db4-f7b024a159a9', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('0f653658-30fc-46f3-854f-465779e8f4d2','voltage','Voltage','b31fcde3-71a6-4c96-9db4-f7b024a159a9', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('ca10b048-fd1d-4e44-b938-2836e0b10a7a','stage','Stage','ea5fb959-5791-4150-a7db-b01691f9bae2', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('0dd25973-4852-48f2-929d-ad163dd6c458','precipitation','Precipitation','ea5fb959-5791-4150-a7db-b01691f9bae2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('45b46905-d73d-4e5c-8b72-3ce07a7407a6','voltage','Voltage','ea5fb959-5791-4150-a7db-b01691f9bae2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('0b31ad67-f2d5-4ce1-8993-257f7bec0fb1','stage','Stage','4043173a-17bb-47d8-b1b9-507075443cb0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('5ee0b15a-1df1-4e4e-856d-febc312bb670','voltage','Voltage','4043173a-17bb-47d8-b1b9-507075443cb0', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4d9b4d6b-8189-4f0a-8a37-999754fc2586','stage','Stage','b360d2a9-7314-42a3-be5d-a69662d69e55', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6f6420c2-7e77-422f-9def-b672ba72d5fc','voltage','Voltage','b360d2a9-7314-42a3-be5d-a69662d69e55', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8ac3c343-593a-4378-ab80-022ac6bc4e26','stage','Stage','461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1a3f49ac-5764-4c4c-b9ff-3eb414883655','precipitation','Precipitation','461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('31234065-9908-43cb-a31c-6b74215bce76','voltage','Voltage','461dd041-66a1-422e-a0b6-ed9ae7f7c3a1', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5a73c8a7-bf41-4fda-a0e7-f1d017548675','stage','Stage','9f1e159e-0e94-4232-a61e-a426294aac5d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('4f8a21ea-c1af-4f9f-a7c3-891959a91870','precipitation','Precipitation','9f1e159e-0e94-4232-a61e-a426294aac5d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b06a9cc7-73fe-43ba-a508-3de4b6a7d5e3','voltage','Voltage','9f1e159e-0e94-4232-a61e-a426294aac5d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d2fb4148-4bda-41a8-a771-bf58aa79abe2','stage','Stage','890a95de-c578-419d-849c-7ce1679ec6b8', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b56353d5-9df6-4a9d-9dcc-0969fe5b47d3','precipitation','Precipitation','890a95de-c578-419d-849c-7ce1679ec6b8', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('edbdc991-929c-4381-bf29-408f77125256','voltage','Voltage','890a95de-c578-419d-849c-7ce1679ec6b8', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e17973a4-c65d-499d-b74e-d702de8d197e','precipitation','Precipitation','b5b9f77f-94bc-4f48-b1c6-e47b208966fb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('33985375-4983-43f4-b9ea-c857df97f0ae','voltage','Voltage','b5b9f77f-94bc-4f48-b1c6-e47b208966fb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('30a315b3-1078-4dbd-b735-766735c9ef8e','stage','Stage','69309767-e07c-474a-9ad6-d38d19599dae', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('bfcf57ea-53ae-4e5f-bba7-77e46becfc04','precipitation','Precipitation','69309767-e07c-474a-9ad6-d38d19599dae', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('52487ca6-dd78-4203-b5be-c12aab8904d7','voltage','Voltage','69309767-e07c-474a-9ad6-d38d19599dae', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4e36293b-f205-419a-bbf2-a99d163bd418','stage','Stage','98c94154-b54a-44a0-b5ab-bc69127ff297', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('2fb600b6-8331-414f-aed5-b260a06246ed','precipitation','Precipitation','98c94154-b54a-44a0-b5ab-bc69127ff297', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('8e08879c-ef36-4107-8bad-85b0a0f9ee1d','voltage','Voltage','98c94154-b54a-44a0-b5ab-bc69127ff297', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('810cd8fa-2707-41b7-9f5c-cea6d3a2d99b','stage','Stage','0b941abe-45b7-470f-8c64-1a90f770dc08', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a67439fb-3d64-43b9-abb5-94108183a995','precipitation','Precipitation','0b941abe-45b7-470f-8c64-1a90f770dc08', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ef2becbc-eccb-4e8c-8868-c0b5d0631da5','voltage','Voltage','0b941abe-45b7-470f-8c64-1a90f770dc08', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8ff10138-4362-4a23-b4dd-21d00560336b','stage','Stage','cce37bfe-98a0-4e16-8b15-2ad8831ca912', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ee6e6700-e59e-49e0-a32d-4346bb4929c6','precipitation','Precipitation','cce37bfe-98a0-4e16-8b15-2ad8831ca912', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b1f656f3-d69c-4caa-8bda-f5bb7d155479','voltage','Voltage','cce37bfe-98a0-4e16-8b15-2ad8831ca912', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('912a4b8a-dcad-4aa5-bd53-1292ac9afb0a','stage','Stage','eaea76f8-0ef4-44f5-ad78-7c7286facf98', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('64f3c9bf-509e-4bc5-bc55-182af8b474b7','precipitation','Precipitation','eaea76f8-0ef4-44f5-ad78-7c7286facf98', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('73af193c-1e96-4265-b82b-9cfe9276bb86','voltage','Voltage','eaea76f8-0ef4-44f5-ad78-7c7286facf98', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('90778d6b-c95c-4ebd-8aa9-934d874e6423','stage','Stage','9e1b87a5-2c84-45fb-8c52-d46acaef525d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('e0947169-47db-4e89-979f-4c76f783f058','precipitation','Precipitation','9e1b87a5-2c84-45fb-8c52-d46acaef525d', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('da3dcbb7-1487-4d3c-a8df-694aec1be19a','voltage','Voltage','9e1b87a5-2c84-45fb-8c52-d46acaef525d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('acbb3082-4262-470b-aa04-fa7fc3cd94ff','stage','Stage','8ed22f3a-2490-4a4b-afb9-7afd853ff491', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('11be0bdb-a78d-4cf8-b4fc-e68d3bb51b36','voltage','Voltage','8ed22f3a-2490-4a4b-afb9-7afd853ff491', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('9c92dbc9-c26d-4c98-856f-288dc1d783b6','stage','Stage','b6217f54-1741-4a6c-93d0-547f9cdec38c', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('33f82faa-7021-4f02-a19a-0852b6c15fa9','precipitation','Precipitation','b6217f54-1741-4a6c-93d0-547f9cdec38c', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('a6fd425c-39a8-4c01-84a2-a044c9b1a595','voltage','Voltage','b6217f54-1741-4a6c-93d0-547f9cdec38c', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('bf02a242-c926-44fb-b4b1-eade889363ff','stage','Stage','07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('98d4b757-7657-4444-a63a-2e64eed11fdd','precipitation','Precipitation','07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ce356001-3030-4a97-b50a-96273cdd26be','voltage','Voltage','07d31ac3-4db3-4ffb-a1fa-8097ed5ca9a6', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('31d93f2e-6b97-43b7-971a-099555f9f5f6','stage','Stage','402a8b7e-688a-4fca-9b1d-31fc77391bd0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a3bd0d60-297a-4a53-af3f-dbeabb191034','precipitation','Precipitation','402a8b7e-688a-4fca-9b1d-31fc77391bd0', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('b9f1460f-4026-44ef-aefe-4495a975ded8','voltage','Voltage','402a8b7e-688a-4fca-9b1d-31fc77391bd0', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4b27a777-e7a8-4e8e-bd3e-40bdd57ed90a','stage','Stage','62454df9-d25f-4d7a-8647-948f7c8393f4', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('6468c079-17ff-4594-a00c-1034e683f221','precipitation','Precipitation','62454df9-d25f-4d7a-8647-948f7c8393f4', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('85f71995-95af-41f3-b96e-303eaa3a28e6','voltage','Voltage','62454df9-d25f-4d7a-8647-948f7c8393f4', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('8982fb3e-aa88-49b1-9ce3-7c74ceb184a6','stage','Stage','47be997a-168a-4f2d-a860-8b95476d0f09', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a0332f1a-aead-47c3-8aef-f76990476ee1','precipitation','Precipitation','47be997a-168a-4f2d-a860-8b95476d0f09', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('88a73011-4d07-4b32-b7ab-83280d198733','voltage','Voltage','47be997a-168a-4f2d-a860-8b95476d0f09', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('1132a518-896a-45ec-b516-cabe38785f46','stage','Stage','c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('44570aa9-df6a-44f2-99f1-d5fdbb0711e8','precipitation','Precipitation','c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('faed230a-ef68-447c-87b0-32efedfde340','voltage','Voltage','c5658c8e-cc5c-4ee8-990a-154bb4b9aa79', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4190ac32-2028-4708-a2e3-4c60bae9c3a0','stage','Stage','a26f2105-23e4-4331-ba82-d3c5d9e777f5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f28b9d36-c532-4264-b221-a055f834e418','precipitation','Precipitation','a26f2105-23e4-4331-ba82-d3c5d9e777f5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('411c166c-9830-4f82-8ff7-752a0fb7e36c','dissolved-oxygen','Dissolved-Oxygen','a26f2105-23e4-4331-ba82-d3c5d9e777f5', '98007857-d027-4524-9a63-d07ae93e5fa2', '67d75ccd-6bf7-4086-a970-5ed65a5c30f3'),
('d48a0af5-2846-4f07-9f67-25acf27ec1cc','voltage','Voltage','a26f2105-23e4-4331-ba82-d3c5d9e777f5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('991b2650-9ab7-4d2f-ac8b-83f84962ddaf','stage','Stage','79b365ba-f543-456e-bf40-01dd7d515bcf', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('14950f03-242c-4c4d-bdba-d7e56f156472','precipitation','Precipitation','79b365ba-f543-456e-bf40-01dd7d515bcf', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('02ac183c-8738-4bcd-8c64-6f64c6ffe9b1','voltage','Voltage','79b365ba-f543-456e-bf40-01dd7d515bcf', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6393812b-1cd9-42c0-9c13-a322bbb83971','precipitation','Precipitation','c2034ba0-a672-42ce-87ef-1f12b2faa79a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('95987dc7-de34-4119-a56b-d732b95aaab9','voltage','Voltage','c2034ba0-a672-42ce-87ef-1f12b2faa79a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2141e506-b07a-485b-887d-0f5ea67ca0f9','stage','Stage','f539acc2-c718-45e6-892a-d6d150317e6f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b3a9f118-f8bb-4424-ab40-5ac55c31ef09','precipitation','Precipitation','f539acc2-c718-45e6-892a-d6d150317e6f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e354d2e1-bd6f-4bbf-9649-459b76ed2794','voltage','Voltage','f539acc2-c718-45e6-892a-d6d150317e6f', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6bd7e4cc-23ca-4592-b156-fd04adf96685','stage','Stage','d8e7ebe1-2e78-4f57-82b4-53a3432c81cc', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('1b1a5ff4-ad78-464b-bf99-a791567dbd35','precipitation','Precipitation','d8e7ebe1-2e78-4f57-82b4-53a3432c81cc', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e074d54e-639f-4c07-ac99-0298ea7083f8','voltage','Voltage','d8e7ebe1-2e78-4f57-82b4-53a3432c81cc', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('39f7ff23-e922-45d6-8da9-208e08de70b0','stage','Stage','97fd6c25-af34-4579-b2a6-9141d03b13de', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('a7401973-8076-4511-99c5-24346c648ade','precipitation','Precipitation','97fd6c25-af34-4579-b2a6-9141d03b13de', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9362e182-fd44-4fb2-a034-15464246f26e','voltage','Voltage','97fd6c25-af34-4579-b2a6-9141d03b13de', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5c818c84-560c-446f-a2a2-1325c40b086f','stage','Stage','f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('dc6a25fb-4fc8-4c79-8e19-589924e02e02','precipitation','Precipitation','f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('171b6601-5739-4e07-b70d-d86d9e6dab15','voltage','Voltage','f6d7ffe0-a7f4-4b5e-b720-1f935f4e1616', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c6d7b11e-7df9-44da-ab72-2248c1acbe35','stage','Stage','2a6b30d6-d2eb-4c9a-b35f-1285108455ef', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('9d22efcd-e1aa-4420-8712-cb3cfe569b74','voltage','Voltage','2a6b30d6-d2eb-4c9a-b35f-1285108455ef', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('837c1796-2d9b-4f59-a466-35920cec5ed7','precipitation','Precipitation','91da2bf4-ffd7-43bf-a8e7-f0dd04311a58', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('9d4c5552-82de-49b9-b70a-c9a952795cf8','voltage','Voltage','91da2bf4-ffd7-43bf-a8e7-f0dd04311a58', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '3254f483-5e66-405c-acf2-2a8add714bf5'),
('26b7bd50-b5bf-4ab3-8a6e-8e80bf65137b','stage','Stage','4d6d85d1-c779-42ad-9cb1-3cafc0a0bcd5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3f8dcb3a-4fc5-423d-916a-7c15d29ed2fa','voltage','Voltage','4d6d85d1-c779-42ad-9cb1-3cafc0a0bcd5', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('2dbfe482-bace-42de-9f15-815babee5aa2','precipitation','Precipitation','d2a58327-b9bf-4862-915c-dcc82878ea05', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('0aea24b3-8580-44f6-b49d-73059cbb77b4','voltage','Voltage','d2a58327-b9bf-4862-915c-dcc82878ea05', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('7d6917cb-451f-4799-899e-0f7a343a1ad4','precipitation','Precipitation','269bc11e-2bf1-4d98-86b2-977f75e4e1fb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('9465aa79-77a8-48f3-9ec1-045d1e3d4f12','stage','Stage','269bc11e-2bf1-4d98-86b2-977f75e4e1fb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d3fb2221-61bf-4d97-9bb3-b3a95068802c','voltage','Voltage','269bc11e-2bf1-4d98-86b2-977f75e4e1fb', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('c474d2c1-e9e1-4a0e-9524-58768322390a','precipitation','Precipitation','798a93e4-0fd7-402f-8c2b-65a8a0850ea2', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('e74c5308-7fc2-4d3b-8b89-d20917f0cc79','voltage','Voltage','798a93e4-0fd7-402f-8c2b-65a8a0850ea2', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5bca4933-c7ba-4da5-9193-9c2bb862a8f9','stage','Stage','f2c66c58-3f34-4f92-a46d-848c62764499', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('ba1f69f5-c978-470a-8d7d-ab507fad4085','voltage','Voltage','f2c66c58-3f34-4f92-a46d-848c62764499', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4f35c112-c8c5-4a84-bb65-0bff4ff87624','stage','Stage','91b3ca6b-8fab-40b3-8911-35beb879f823', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('29ea3867-d9cd-4885-8dd6-821b21f878a0','voltage','Voltage','91b3ca6b-8fab-40b3-8911-35beb879f823', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('d82a9071-d631-43f8-8194-9330059332cf','stage','Stage','aa4f72e5-1fea-4ae6-b7cb-8a7b05579120', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('f1f095ac-d610-4b62-b2d6-3dd13cbd123d','voltage','Voltage','aa4f72e5-1fea-4ae6-b7cb-8a7b05579120', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5416d7c8-ef18-4843-9c09-be00f8e9d0b8','stage','Stage','fbcac5af-25a7-4e23-be64-ecf18c9fe23e', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('c7beb772-13f9-4f5f-8ef7-28c8d2a8c9ef','voltage','Voltage','fbcac5af-25a7-4e23-be64-ecf18c9fe23e', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('e8076f82-7de2-4079-94ec-610d011121d0','stage','Stage','abd86a4b-db70-4aac-b397-96fa6fee821a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('19f33118-f745-4b06-a7aa-ff8109d67ed8','precipitation','Precipitation','abd86a4b-db70-4aac-b397-96fa6fee821a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('7a2a3947-054e-461d-9eb2-93ed39a7922b','voltage','Voltage','abd86a4b-db70-4aac-b397-96fa6fee821a', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5572b4eb-a228-40ef-bd50-0435fab9ba5e','stage','Stage','c2013f5c-ac88-44cb-a33e-f0ed8d2d8109', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('7e380e01-98de-468c-93e2-eb4d545120c1','voltage','Voltage','c2013f5c-ac88-44cb-a33e-f0ed8d2d8109', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('5b99094c-f707-48ea-bc5c-468c36017722','stage','Stage','cc927dcf-5a23-460b-b7ec-dc731374edaa', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('b750ac8e-9a32-42fe-9172-f2c46e68e831','precipitation','Precipitation','cc927dcf-5a23-460b-b7ec-dc731374edaa', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a'),
('ac039ec8-0560-4032-9f53-414fe8f083f2','voltage','Voltage','cc927dcf-5a23-460b-b7ec-dc731374edaa', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('4324785a-6ad5-4c87-916c-c9620bbba393','stage','Stage','87d871a1-3a41-4442-b61d-bd65e3ac7b8d', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('d4c5ba46-98c9-4a2b-88b0-d5fae2b328b6','voltage','Voltage','87d871a1-3a41-4442-b61d-bd65e3ac7b8d', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('6b492f75-4c8b-4eef-8f02-8c0243c223a3','stage','Stage','47b3ed3c-bee7-427f-9174-4c0bc84a5685', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('7ee9828b-7b5e-46e1-a035-f18fec4d3e0d','voltage','Voltage','47b3ed3c-bee7-427f-9174-4c0bc84a5685', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64'),
('3d33b28c-b3ae-4191-bf62-7b0dad24f259','stage','Stage','82fc7100-670d-4076-9292-5e4020aa9478', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce'),
('3c5e17f4-0af8-49fe-bebb-271137f69c9a','voltage','Voltage','82fc7100-670d-4076-9292-5e4020aa9478', '430e5edb-e2b5-4f86-b19f-cda26a27e151', '6b5bd788-8c78-43bb-b5a3-ad544b858a64');
