-- extensions
CREATE extension IF NOT EXISTS "uuid-ossp";


-- drop tables if they already exist
drop table if exists 
    public.timeseries_measurement,
    public.timeseries,
    public.instrument_group_instruments,
    public.instrument_status,
    public.instrument_zreference,
    public.instrument_note,
    public.instrument,
    public.instrument_group,
    public.parameter,
    public.unit,
    public.instrument_type,
    public.project,
    public.status,
    public.zreference_datum
	CASCADE;

-- project
CREATE TABLE IF NOT EXISTS public.project (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    deleted boolean NOT NULL DEFAULT false,
    slug VARCHAR(240) UNIQUE NOT NULL,
    federal_id VARCHAR(240),
    name VARCHAR(240) UNIQUE NOT NULL,
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- instrument_type
CREATE TABLE IF NOT EXISTS public.instrument_type (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- domain status
CREATE TABLE IF NOT EXISTS public.status (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(480)
);

-- domain zreference_datum
CREATE TABLE IF NOT EXISTS public.zreference_datum (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- unit
CREATE TABLE IF NOT EXISTS public.unit (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- parameter
CREATE TABLE IF NOT EXISTS public.parameter (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- instrument_group
CREATE TABLE IF NOT EXISTS public.instrument_group (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    deleted BOOLEAN NOT NULL DEFAULT false,
    slug VARCHAR(240) UNIQUE NOT NULL,
    name VARCHAR(120) NOT NULL,
    description VARCHAR(360),
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    project_id UUID REFERENCES project (id),
    CONSTRAINT project_unique_instrument_group_name UNIQUE(name,project_id)
	);

-- instrument
CREATE TABLE IF NOT EXISTS public.instrument (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    deleted BOOLEAN NOT NULL DEFAULT false,
    slug VARCHAR(240) UNIQUE NOT NULL,
    name VARCHAR(120) NOT NULL,
    geometry geometry,
    station int,
    station_offset int,
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    type_id UUID NOT NULL REFERENCES instrument_type (id),
    project_id UUID REFERENCES project (id),
    formula VARCHAR(360),
    CONSTRAINT project_unique_instrument_name UNIQUE(name,project_id)
);

-- instrument_note
CREATE TABLE IF NOT EXISTS public.instrument_note (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    title VARCHAR(240) NOT NULL,
    body VARCHAR(65535) NOT NULL,
    time TIMESTAMPTZ NOT NULL DEFAULT now(),
    creator BIGINT NOT NULL DEFAULT 0,
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater BIGINT NOT NULL DEFAULT 0,
    update_date TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- instrument_group_instruments
CREATE TABLE IF NOT EXISTS public.instrument_group_instruments (
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    instrument_group_id UUID NOT NULL REFERENCES instrument_group (id),
    UNIQUE (instrument_id, instrument_group_id)
);

-- instrument_status
CREATE TABLE IF NOT EXISTS public.instrument_status (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    status_id UUID NOT NULL REFERENCES status (id),
    time TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT instrument_unique_status_in_time UNIQUE (instrument_id, time)
);

-- instrument_zreference
CREATE TABLE IF NOT EXISTS public.instrument_zreference (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    time TIMESTAMPTZ NOT NULL DEFAULT '1776-08-02',
    zreference REAL NOT NULL,
    zreference_datum_id UUID NOT NULL REFERENCES zreference_datum (id),
    CONSTRAINT instrument_unique_zreference_in_time UNIQUE(instrument_id, time)
);

-- timeseries
CREATE TABLE IF NOT EXISTS public.timeseries (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR(240) UNIQUE NOT NULL,
    name VARCHAR(240) NOT NULL,
    instrument_id UUID REFERENCES instrument (id),
    parameter_id UUID NOT NULL REFERENCES parameter (id),
    unit_id UUID NOT NULL REFERENCES unit (id),
    CONSTRAINT instrument_unique_timeseries_name UNIQUE(instrument_id, name)
);

-- timeseries_measurement
CREATE TABLE IF NOT EXISTS public.timeseries_measurement (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    time TIMESTAMPTZ NOT NULL,
    value REAL NOT NULL,
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT timeseries_unique_time UNIQUE(timeseries_id,time)
);

-- constants
CREATE TABLE IF NOT EXISTS public.constants (
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    timeseries_id UUID NOT NULL REFERENCES timeseries (id),
)

-- alarms
CREATE TABLE IF NOT EXISTS public.alarms (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    formula VARCHAR(360),
    schedule VARCHAR(20),
    mute_notifications BOOLEAN,
    mute_ui BOOLEAN,
    to VARCHAR(360),
    body VARCHAR(1200)
)

-- annotations
CREATE TABLE IF NOT EXISTS public.annotations (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    x VARCHAR(100), -- we have to be able to store date-strings as well as numbers
    y REAL,
    body VARCHAR(100)
)

-- -------
-- Domains
-- -------

-- instrument_type
INSERT INTO instrument_type (id, name) VALUES
    ('0fd1f9ba-2731-4ff9-96dd-3c03215ab06f', 'Staff Gage'),
    ('1bb4bf7c-f5f8-44eb-9805-43b07ffadbef', 'Piezometer');
-- status
INSERT INTO status (id, name, description) VALUES
    ('e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', 'active', 'description for active instrument status'),
    ('c9ee4acb-9623-4fde-bf36-7668afe463d4', 'inactive', 'description for inactive instrument status'),
    ('3d9add10-4418-4cb2-bd31-20a639504539', 'abandoned', 'description for abandoned instrument status'),
    ('94578354-ffdf-4119-9663-6bd4323e58f5', 'destroyed', 'description for destroyed instrument status'),
    ('03a2bf9a-bbd8-4031-8f4e-13e8c77807f1', 'lost', 'description for lost instrument status');
-- parameter
INSERT INTO parameter (id, name) VALUES
    ('068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'length'),
    ('1de79e29-fb70-45c3-ae7d-4695517ced90', 'pressure'),
    ('de6112da-8489-4286-ae56-ec72aa09974d', 'temperature'),
    ('0ce77a5a-8283-47cd-9126-c440bcec4ef6', 'precipitation'),
    ('430e5edb-e2b5-4f86-b19f-cda26a27e151', 'voltage');
-- unit
INSERT INTO unit (id, name) VALUES
    ('7b924ec2-c488-401d-a503-ca734b1ab804', 'feet'),
    ('ac41edd9-b485-4e85-8f03-fe79b9a44305', 'millibar'),
    ('e0f8ecb4-02da-46f6-b10b-37b0bf09b43c', 'inches'),
    ('2ccf581b-d5cd-455a-a981-d4fc1efdc8a4', 'inches mercury (Hg)'),
    ('4b7773a4-54dc-4c92-a271-a68c2a77deb8', 'volts');

-- zreference_datum (https://www.ngs.noaa.gov/datums/vertical/)
INSERT INTO zreference_datum (id, name) VALUES
    ('85fb892d-7d55-41f1-95f6-addea9914264', 'National Geodetic Vertical Datum of 1929 (NGVD 29)'),
    ('72113f9a-982d-44e5-8fc1-8e595dafd344', 'North American Vertical Datum of 1988 (NAVD 88)');
-- -------------------------------------------------
-- basic seed data to demo the app and run API tests
-- -------------------------------------------------
-- project
INSERT INTO project (id, slug, name) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'blue-water-dam-example-project', 'Blue Water Dam Example Project');

-- instrument_group
INSERT INTO instrument_group (project_id, id, slug, name, description) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'd0916e8a-39a6-4f2f-bd31-879881f8b40c', 'sample-instrument-group', 'Sample Instrument Group 1', 'This is an example instrument group');

-- instrument
INSERT INTO instrument (project_id, id, slug, name, geometry, type_id) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'demo-piezometer-1', 'Demo Piezometer 1', ST_GeomFromText('POINT(-80.8 26.7)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '9e8f2ca4-4037-45a4-aaca-d9e598877439', 'demo-staffgage-1', 'Demo Staffgage 1', ST_GeomFromText('POINT(-80.85 26.75)',4326),'0fd1f9ba-2731-4ff9-96dd-3c03215ab06f');

-- instrument_group_instruments
INSERT INTO instrument_group_instruments (instrument_id, instrument_group_id) VALUES
    ('a7540f69-c41e-43b3-b655-6e44097edb7e', 'd0916e8a-39a6-4f2f-bd31-879881f8b40c');

-- instrument_zreference
-- Following simulates the described scenario
-- (1) Initial reference height (pz casing installed)
-- (2) PZ casing hit by mower, reducing height by 0.5 ft
-- (3) pz casing repaired/extended to be 4.0 ft higher
INSERT INTO instrument_zreference (id, instrument_id, time, zreference, zreference_datum_id) VALUES
    ('3f4718fb-897b-4840-8494-0f142cc11027', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '1975-01-01', 41.60, '85fb892d-7d55-41f1-95f6-addea9914264'),
    ('d18e5577-5b28-4722-8833-5cb14430e02a', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2000-01-01', 41.0, '72113f9a-982d-44e5-8fc1-8e595dafd344'),
    ('99841fcc-16e7-408b-ba53-126bd2e764d1', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2005-07-01', 40.5, '72113f9a-982d-44e5-8fc1-8e595dafd344'),
    ('9b16d6bf-81ab-488d-a650-996280c628dc', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '2006-06-01', 44.5, '72113f9a-982d-44e5-8fc1-8e595dafd344'),
    ('996b9650-24c7-4e43-b65c-abd767041ecd', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '2020-01-01', 10.5, '72113f9a-982d-44e5-8fc1-8e595dafd344');

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
('869465fc-dc1e-445e-81f4-9979b5fadda9', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '1de79e29-fb70-45c3-ae7d-4695517ced90', '2ccf581b-d5cd-455a-a981-d4fc1efdc8a4', 'atmospheric-pressure', 'Atmospheric Pressure'),
('9a3864a8-8766-4bfa-bad1-0328b166f6a8', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '0ce77a5a-8283-47cd-9126-c440bcec4ef6', 'e0f8ecb4-02da-46f6-b10b-37b0bf09b43c', 'precipitation', 'Precipitation'),
('7ee902a3-56d0-4acf-8956-67ac82c03a96', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '7b924ec2-c488-401d-a503-ca734b1ab804', 'height', 'Height'),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '7b924ec2-c488-401d-a503-ca734b1ab804', 'height-1', 'Height');

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
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '3/10/2020', 20.05);
