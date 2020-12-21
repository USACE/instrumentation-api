-- Project
INSERT INTO project (id, slug, name) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'buffalo-district-streamgages', 'Buffalo District Streamgages');

-- Instrument Groups
INSERT INTO instrument_group (project_id, id, slug, name, description) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'f232edd7-5098-428b-af14-54a5dfd796a7', 'genesee', 'Genesee', 'Tributary of Lake Ontario');

-- Instruments
INSERT INTO instrument (project_id, id, slug, name, formula, geometry, type_id) VALUES
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'b695967d-c050-4428-ab1e-db9407fe9d2f', 'akln6', 'AKLN6', null, ST_GeomFromText('POINT(-77.72 42.4)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '089fbda3-5a6a-408d-aec2-4e108910d94b', 'avnn6', 'AVNN6', null, ST_GeomFromText('POINT(-77.76 42.92)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '08735594-14f0-4594-aa38-fb0805968918', 'blackcr-churchvl', 'BlackCr Churchvl', null, ST_GeomFromText('POINT(-77.88 43.1)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '209980f3-ab26-42d6-9fe7-13c0b6221f88', 'blbn6', 'BLBN6', null, ST_GeomFromText('POINT(-77.68 43.09)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'canaseragashaker', 'CanaseragaShaker', null, ST_GeomFromText('POINT(-77.75 42.92)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'chcn6', 'CHCN6', null, ST_GeomFromText('POINT(-77.88 43.1)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'dsvn6', 'DSVN6', null, ST_GeomFromText('POINT(-77.71 42.53)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '31be5853-f839-4cc2-80cb-071210651059', 'elkp1', 'ELKP1', null, ST_GeomFromText('POINT(-77.3 41.99)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '8df04db9-392f-4216-9776-66defd9d8d32', 'frkn6', 'FRKN6', null, ST_GeomFromText('POINT(-78.46 42.33)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'garn6', 'GARN6', null, ST_GeomFromText('POINT(-77.79 43.01)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'c559d145-3351-4e6e-bf8d-df05f69814ea', 'genr-avon', 'GenR Avon', null, ST_GeomFromText('POINT(-77.76 42.92)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '871485de-0116-4f4a-afcd-30f3f8ec02c5', 'genr-portagevill', 'GenR Portagevill', null, ST_GeomFromText('POINT(-78.04 42.57)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'genr-wellsville', 'GenR Wellsville', null, ST_GeomFromText('POINT(-77.96 42.12)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '1dc2a261-8b00-4039-8d3f-f48666e97729', 'hnyn6', 'HNYN6', null, ST_GeomFromText('POINT(-77.59 42.96)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'a6e1707c-08f9-4484-ba65-cd054e9561bd', 'hrln6', 'HRLN6', null, ST_GeomFromText('POINT(-77.7 42.35)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'jonn6', 'JONN6', null, ST_GeomFromText('POINT(-77.84 42.77)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'knvn6', 'KNVN6', null, ST_GeomFromText('POINT(-78.31 43.3)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '3b260872-3454-473b-bfb7-087ed0e809b0', 'mbyp1', 'MBYP1', null, ST_GeomFromText('POINT(-77.75 42.92)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'ef991379-a625-4de7-bc6a-2902fea1bc79', 'mount-morris', 'Mount Morris', null, ST_GeomFromText('POINT(-77.90690595378125 42.73327583350692)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '4e06209a-9e0b-4430-87fe-a7ac49d92160', 'oatkacr-garbutt', 'OatkaCr Garbutt', null, ST_GeomFromText('POINT(-77.79 43.01)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'oatkacr-warsaw', 'OatkaCr Warsaw', null, ST_GeomFromText('POINT(-78.14 42.74)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '606dcda7-8287-4435-9d66-deb4c03b4bf6', 'olnn6', 'OLNN6', null, ST_GeomFromText('POINT(-78.45 42.07)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '90664d51-0503-45cf-85bb-5d46fa579d0f', 'ptgn6', 'PTGN6', null, ST_GeomFromText('POINT(-78.04 42.57)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '15614ca3-c115-4cff-a021-c43ef352fda8', 'rcrn6', 'RCRN6', null, ST_GeomFromText('POINT(-77.6 43.26)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'rohn6', 'ROHN6', null, ST_GeomFromText('POINT(-77.62 43.14)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '2673f0e2-16a4-443a-ac86-7c337d960f4f', 'shnp1', 'SHNP1', null, ST_GeomFromText('POINT(-78.2 41.96)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'weln6', 'WELN6', null, ST_GeomFromText('POINT(-77.96 42.12)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984'),
    ('a012e753-9eff-426d-b0ee-090b430d1980', 'f063108f-9c3b-49fd-b981-1be021278ebf', 'wrsn6', 'WRSN6', null, ST_GeomFromText('POINT(-78.14 42.74)',4326), '98a61f29-18a8-430a-9d02-0f53486e0984');

-- Instrument Group Instruments
INSERT INTO instrument_group_instruments (instrument_id, instrument_group_id) VALUES
    ('b695967d-c050-4428-ab1e-db9407fe9d2f', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('089fbda3-5a6a-408d-aec2-4e108910d94b', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('209980f3-ab26-42d6-9fe7-13c0b6221f88', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('08735594-14f0-4594-aa38-fb0805968918', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('31be5853-f839-4cc2-80cb-071210651059', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('8df04db9-392f-4216-9776-66defd9d8d32', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('c559d145-3351-4e6e-bf8d-df05f69814ea', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('871485de-0116-4f4a-afcd-30f3f8ec02c5', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('1dc2a261-8b00-4039-8d3f-f48666e97729', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('a6e1707c-08f9-4484-ba65-cd054e9561bd', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('3b260872-3454-473b-bfb7-087ed0e809b0', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('ef991379-a625-4de7-bc6a-2902fea1bc79', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('606dcda7-8287-4435-9d66-deb4c03b4bf6', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('4e06209a-9e0b-4430-87fe-a7ac49d92160', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('90664d51-0503-45cf-85bb-5d46fa579d0f', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('15614ca3-c115-4cff-a021-c43ef352fda8', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('2673f0e2-16a4-443a-ac86-7c337d960f4f', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'f232edd7-5098-428b-af14-54a5dfd796a7'),
    ('f063108f-9c3b-49fd-b981-1be021278ebf', 'f232edd7-5098-428b-af14-54a5dfd796a7');


-- instrument_status
-- status 'active' for all instruments in January 01, 2000
INSERT INTO instrument_status (instrument_id, status_id) VALUES
    ('4e06209a-9e0b-4430-87fe-a7ac49d92160', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('a4fc3899-be8d-4055-b9e9-cd9ff986ba98', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('3b260872-3454-473b-bfb7-087ed0e809b0', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('606dcda7-8287-4435-9d66-deb4c03b4bf6', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('2673f0e2-16a4-443a-ac86-7c337d960f4f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('089fbda3-5a6a-408d-aec2-4e108910d94b', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('209980f3-ab26-42d6-9fe7-13c0b6221f88', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('1dc2a261-8b00-4039-8d3f-f48666e97729', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('08735594-14f0-4594-aa38-fb0805968918', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('c559d145-3351-4e6e-bf8d-df05f69814ea', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('871485de-0116-4f4a-afcd-30f3f8ec02c5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('b695967d-c050-4428-ab1e-db9407fe9d2f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('90664d51-0503-45cf-85bb-5d46fa579d0f', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('f063108f-9c3b-49fd-b981-1be021278ebf', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('15614ca3-c115-4cff-a021-c43ef352fda8', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('31be5853-f839-4cc2-80cb-071210651059', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('8df04db9-392f-4216-9776-66defd9d8d32', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('a6e1707c-08f9-4484-ba65-cd054e9561bd', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d'),
    ('ef991379-a625-4de7-bc6a-2902fea1bc79', 'e26ba2ef-9b52-4c71-97df-9e4b6cf4174d');

-- Time Series
INSERT INTO timeseries (id, instrument_id, parameter_id, unit_id, slug, name) VALUES
    ('e3539e7a-80d7-46df-a4e9-9701fa75fdf4', 'c05efa0c-fa2c-4776-a49b-dd68f70f5104', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('ad0dba35-c9a8-4c26-9e18-4fb4af67e0ed', '4e06209a-9e0b-4430-87fe-a7ac49d92160', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('977371a7-a8f2-4807-b527-6b581583c0f2', '08735594-14f0-4594-aa38-fb0805968918', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('9ce7d4a4-22da-4915-8910-45d78f754109', '871485de-0116-4f4a-afcd-30f3f8ec02c5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('e96402a2-918a-4d8b-86aa-3fd6f16c106b', '3b260872-3454-473b-bfb7-087ed0e809b0', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('35276612-3852-4064-b536-b7c37a032cfc', '089fbda3-5a6a-408d-aec2-4e108910d94b', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('61f7bc29-efc4-4eaa-b882-cce08117884a', '209980f3-ab26-42d6-9fe7-13c0b6221f88', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('5c073e76-3035-486a-8c35-34bfa95fb5b7', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('398129f7-5091-44b7-b6f2-22b952b8c586', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('6177acc9-3f98-4663-ac47-da74d931f78f', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('8a0b0e33-c7ae-494d-bb81-9e3fac025751', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('84d6f643-5616-46e1-9a42-5a1087d82c01', '90664d51-0503-45cf-85bb-5d46fa579d0f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('a6aa8c94-03db-475a-83ed-f302a87b539b', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('9f125882-9669-4bb4-9115-6f933b054477', '4f8c4f01-5e19-4a60-b55b-e7013e8977bb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('b21dee5b-698f-4dc4-b1e2-148448aacda6', 'c559d145-3351-4e6e-bf8d-df05f69814ea', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('be554569-9e27-4446-9d47-c6d665c09b9d', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('d8227f16-6e0d-42ff-a68d-e319fa832520', '606dcda7-8287-4435-9d66-deb4c03b4bf6', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('d118bf74-ce55-4a82-9deb-ef4e2d8b77e0', '0db970b5-3c46-4f54-b17b-8e861c0a5d65', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('f854d6c3-05c1-41c7-ab1c-5e17cccd2a07', '2673f0e2-16a4-443a-ac86-7c337d960f4f', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('63681eb3-c77c-4207-814a-f0338906b96c', '1dc2a261-8b00-4039-8d3f-f48666e97729', 'b49f214e-f69f-43da-9ce3-ad96042268d0', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'stage', 'Stage'),
    ('ef4c5de6-7cd5-4995-ae83-66fc88383736', 'a6e1707c-08f9-4484-ba65-cd054e9561bd', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('6d7c1bb9-4ac9-4453-8640-a304bf8a2599', 'b695967d-c050-4428-ab1e-db9407fe9d2f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('0cc95a8e-2fae-4ce0-8479-8937d51142ef', '3b260872-3454-473b-bfb7-087ed0e809b0', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('af3d5b4c-42c8-4fd8-942e-85eed81a81d7', '089fbda3-5a6a-408d-aec2-4e108910d94b', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('face6e5b-bf87-4fd7-b91c-a6e4d11fb1ed', '209980f3-ab26-42d6-9fe7-13c0b6221f88', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('84b38384-8e98-44df-aad5-379ba97f519f', '8d0307a9-7a78-4eae-919a-b38a6b1c9c97', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('eb43e42b-c4db-4c55-96f4-df9866768ae7', 'b4d043f5-82bc-4d1e-abdf-a43fe36f1695', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('8fb36b31-5d1e-487c-9cc8-190e7467032a', '891c7d48-fa1b-4f4a-a7dd-7393258bb575', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('87710d3b-b005-4df8-8df8-f17152ff615d', 'b9b0ec3f-3f84-4c41-b8ff-55bdf7b6138a', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('d119edc7-9d70-423f-8087-17854d110e0e', '90664d51-0503-45cf-85bb-5d46fa579d0f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('9bba44f4-c848-47ee-9f84-45f7b7513c25', '2ae81f1b-2bcb-40e8-9f12-5874d4269cc5', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('4b428c1f-9d71-4a4a-8a25-44ca781e17e7', 'ef991379-a625-4de7-bc6a-2902fea1bc79', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('639339c1-194a-404f-9d3a-54bf50c4d290', '8df04db9-392f-4216-9776-66defd9d8d32', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('0aac6480-4438-4ee6-b362-43ec82ae7b9d', 'abf85e4e-e9f5-4814-9ecd-d21de65a2feb', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('ec304413-dfdc-446c-af30-4eb45a2f6a91', '606dcda7-8287-4435-9d66-deb4c03b4bf6', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('e4e8ba0c-1cf6-4cf5-a3f3-b0925c3bd94a', '2673f0e2-16a4-443a-ac86-7c337d960f4f', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation'),
    ('cdff046e-2452-4749-8da8-d7093a0470f3', '1dc2a261-8b00-4039-8d3f-f48666e97729', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', '4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'precipitation', 'Precipitation');


-- Time Series Measurements
-- INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES
-- ('37bffc9f-0cd5-48ed-9b4a-b44a2a38bbc2', '2009-07-17 10:25:00', 12.090),