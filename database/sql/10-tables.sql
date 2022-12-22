-- extensions
CREATE extension IF NOT EXISTS "uuid-ossp";

-- Create MIDAS Schema
SET search_path TO midas,topology,public;
CREATE SCHEMA midas;

-- drop tables if they already exist
drop table if exists 
    profile_project_roles,
    role,
    timeseries_measurement,
    timeseries_notes,
    timeseries,
    instrument_telemetry,
    telemetry_goes,
    telemetry_iridium,
    telemetry_type,
    instrument_group_instruments,
    instrument_status,
    instrument_note,
    instrument,
    instrument_group,
    instrument_constants,
    parameter,
    unit_family,
    measure,
    unit,
    instrument_type,
    project_timeseries,
    project,
    status,
    profile,
    profile_token,
    email,
    alert,
    alert_read,
    alert_config,
    alert_profile_subscription,
    alert_email_subscription,
    heartbeat,
    collection_group_timeseries,
    collection_group,
    plot_configuration,
    plot_configuration_timeseries,
    config
	CASCADE;

-- config (application config variables)
CREATE TABLE IF NOT EXISTS config (
    static_host VARCHAR NOT NULL DEFAULT 'http://minio:9000',
    static_prefix VARCHAR NOT NULL DEFAULT '/instrumentation'
);

-- role
CREATE TABLE IF NOT EXISTS role (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    deleted boolean NOT NULL DEFAULT false
);

-- profile
CREATE TABLE IF NOT EXISTS profile (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    edipi BIGINT UNIQUE NOT NULL,
    username VARCHAR(240) UNIQUE NOT NULL,
    email VARCHAR(240) UNIQUE NOT NULL,
    is_admin boolean NOT NULL DEFAULT false
);

-- profile_token
CREATE TABLE IF NOT EXISTS profile_token (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    token_id VARCHAR NOT NULL,
    profile_id UUID NOT NULL REFERENCES profile(id),
    issued TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    hash VARCHAR(240) NOT NULL
);

-- email (user that will never login but still needs alerts; i.e. just an email)
CREATE TABLE IF NOT EXISTS email (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(240) UNIQUE NOT NULL
);

-- project
CREATE TABLE IF NOT EXISTS project (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    image VARCHAR,
    office_id UUID,
    federal_id VARCHAR,
    deleted boolean NOT NULL DEFAULT false,
    slug VARCHAR(240) UNIQUE NOT NULL,
    name VARCHAR(240) UNIQUE NOT NULL,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ
);

-- profile_project_roles
CREATE TABLE IF NOT EXISTS profile_project_roles (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    profile_id UUID NOT NULL REFERENCES profile(id),
    role_id UUID NOT NULL REFERENCES role(id),
    project_id UUID NOT NULL REFERENCES project(id),
    granted_by UUID REFERENCES profile(id),
    granted_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_profile_project_role UNIQUE(profile_id,project_id,role_id)
);

-- heartbeat
CREATE TABLE IF NOT EXISTS heartbeat (
    time TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- instrument_type
CREATE TABLE IF NOT EXISTS instrument_type (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- domain status
CREATE TABLE IF NOT EXISTS status (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(480)
);

-- measure
CREATE TABLE IF NOT EXISTS measure (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(240) UNIQUE NOT NULL
);

-- unit_family
CREATE TABLE IF NOT EXISTS unit_family (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- unit
CREATE TABLE IF NOT EXISTS unit (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL,
    abbreviation VARCHAR(120) UNIQUE NOT NULL,
    unit_family_id UUID REFERENCES unit_family (id),
    measure_id UUID REFERENCES measure (id)
);

-- parameter
CREATE TABLE IF NOT EXISTS parameter (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- instrument_group
CREATE TABLE IF NOT EXISTS instrument_group (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    deleted BOOLEAN NOT NULL DEFAULT false,
    slug VARCHAR(240) UNIQUE NOT NULL,
    name VARCHAR(120) NOT NULL,
    description VARCHAR(360),
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    project_id UUID REFERENCES project (id),
    CONSTRAINT project_unique_instrument_group_name UNIQUE(name,project_id)
	);

-- instrument
CREATE TABLE IF NOT EXISTS instrument (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    deleted BOOLEAN NOT NULL DEFAULT false,
    slug VARCHAR UNIQUE NOT NULL,
    name VARCHAR(360) NOT NULL,
    geometry geometry,
    station int,
    station_offset int,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    type_id UUID NOT NULL REFERENCES instrument_type (id),
    project_id UUID REFERENCES project (id),
    nid_id VARCHAR,
    usgs_id VARCHAR
);

-- alert_config
CREATE TABLE IF NOT EXISTS alert_config (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    name VARCHAR(480),
    body TEXT,
    formula TEXT,
    schedule TEXT,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    CONSTRAINT instrument_unique_alert_config_name UNIQUE(name,instrument_id)
);

-- alert
CREATE TABLE IF NOT EXISTS alert (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    create_date TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- alert_read
CREATE TABLE IF NOT EXISTS alert_read (
    alert_id UUID NOT NULL REFERENCES alert (id),
    profile_id UUID NOT NULL REFERENCES profile (id),
    CONSTRAINT profile_unique_alert_read UNIQUE(alert_id, profile_id)
);

-- profile alerts (subscribe profiles to alerts)
CREATE TABLE IF NOT EXISTS alert_profile_subscription (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    profile_id UUID NOT NULL REFERENCES profile (id),
    mute_ui boolean NOT NULL DEFAULT false,
    mute_notify boolean NOT NULL DEFAULT false,
    CONSTRAINT profile_unique_alert_config UNIQUE(profile_id, alert_config_id)
);

-- email alerts (subscribe emails to alerts)
CREATE TABLE IF NOT EXISTS alert_email_subscription (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    email_id UUID NOT NULL REFERENCES email (id),
    mute_notify boolean NOT NULL DEFAULT false,
    CONSTRAINT email_unique_alert_config UNIQUE(email_id, alert_config_id)
);

-- instrument_note
CREATE TABLE IF NOT EXISTS instrument_note (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    title VARCHAR(240) NOT NULL,
    body VARCHAR(65535) NOT NULL,
    time TIMESTAMPTZ NOT NULL DEFAULT now(),
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ
);

-- instrument_group_instruments
CREATE TABLE IF NOT EXISTS instrument_group_instruments (
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    instrument_group_id UUID NOT NULL REFERENCES instrument_group (id),
    UNIQUE (instrument_id, instrument_group_id)
);

-- instrument_status
CREATE TABLE IF NOT EXISTS instrument_status (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument (id),
    status_id UUID NOT NULL REFERENCES status (id),
    time TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT instrument_unique_status_in_time UNIQUE (instrument_id, time)
);

-- timeseries
CREATE TABLE IF NOT EXISTS timeseries (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR(240) NOT NULL,
    name VARCHAR(240) NOT NULL,
    instrument_id UUID REFERENCES instrument (id),
    parameter_id UUID NOT NULL REFERENCES parameter (id),
    unit_id UUID NOT NULL REFERENCES unit (id),
    CONSTRAINT instrument_unique_timeseries_name UNIQUE(instrument_id, name),
    CONSTRAINT instrument_unique_timeseries_slug UNIQUE(instrument_id, slug)
);


-- calculation
CREATE TABLE IF NOT EXISTS calculation (
	timeseries_id UUID UNIQUE NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
	contents VARCHAR
);

-- timeseries_measurement
CREATE TABLE IF NOT EXISTS timeseries_measurement (
    --id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    time TIMESTAMPTZ NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT timeseries_unique_time UNIQUE(timeseries_id,time),
    PRIMARY KEY (timeseries_id, time)
);

-- timeseries_notes
CREATE TABLE IF NOT EXISTS timeseries_notes (
    masked boolean NOT NULL DEFAULT false,
    validated boolean NOT NULL DEFAULT false,
    annotation varchar(400) NOT NULL DEFAULT '',
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    time TIMESTAMPTZ NOT NULL,
    CONSTRAINT notes_unique_time UNIQUE(timeseries_id, time),
    PRIMARY KEY (timeseries_id, time)
);

-- inclinometer_measurement
CREATE TABLE IF NOT EXISTS inclinometer_measurement (
    time TIMESTAMPTZ NOT NULL,
    values JSONB NOT NULL,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    timeseries_id UUID NOT NULL REFERENCES timeseries (id) ON DELETE CASCADE,
    CONSTRAINT inclinometer_unique_time UNIQUE(timeseries_id,time),
    PRIMARY KEY (timeseries_id, time)
);

-- instrument_constants
CREATE TABLE IF NOT EXISTS instrument_constants (
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    instrument_id UUID NOT NULL REFERENCES instrument(id) ON DELETE CASCADE,
    CONSTRAINT instrument_unique_timeseries UNIQUE(instrument_id, timeseries_id)
);

-- project_timeseries
CREATE TABLE IF NOT EXISTS project_timeseries (
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    project_id UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    CONSTRAINT project_unique_timeseries UNIQUE(project_id, timeseries_id)
);

-- collection_group
CREATE TABLE IF NOT EXISTS collection_group (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    CONSTRAINT project_unique_collection_group_name UNIQUE(project_id, name),
    CONSTRAINT project_unique_collection_group_slug UNIQUE(project_id, slug)
);

CREATE TABLE IF NOT EXISTS collection_group_timeseries (
    collection_group_id UUID NOT NULL REFERENCES collection_group(id) ON DELETE CASCADE,
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT collection_group_unique_timeseries UNIQUE(collection_group_id, timeseries_id)
);

-- plot_configuration
CREATE TABLE IF NOT EXISTS plot_configuration (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    project_id UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    creator UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater UUID,
    update_date TIMESTAMPTZ,
    CONSTRAINT project_unique_plot_configuration_name UNIQUE(project_id, name),
    CONSTRAINT project_unique_plot_configuration_slug UNIQUE(project_id, slug)
);


CREATE TABLE IF NOT EXISTS plot_configuration_timeseries (
    plot_configuration_id UUID NOT NULL REFERENCES plot_configuration(id) ON DELETE CASCADE,
    timeseries_id UUID NOT NULL REFERENCES timeseries(id) ON DELETE CASCADE,
    CONSTRAINT plot_configuration_unique_timeseries UNIQUE(plot_configuration_id, timeseries_id)
);


CREATE TABLE IF NOT EXISTS plot_configuration_settings (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    show_masked BOOLEAN DEFAULT 'true',
    show_nonvalidated BOOLEAN DEFAULT 'true',
    show_comments BOOLEAN DEFAULT 'true',
    auto_range BOOLEAN DEFAULT 'true',
    date_range VARCHAR(23) DEFAULT '1 year',
    FOREIGN KEY (id) REFERENCES plot_configuration (id) ON DELETE CASCADE
);


-- ---------
-- Telemetry
-- ---------

-- telemetry_type
CREATE TABLE IF NOT EXISTS telemetry_type (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    slug VARCHAR UNIQUE NOT NULL,
    name VARCHAR UNIQUE NOT NULL
);

-- instrument_telemetry
CREATE TABLE IF NOT EXISTS instrument_telemetry (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    instrument_id UUID NOT NULL REFERENCES instrument(id),
    telemetry_type_id UUID NOT NULL REFERENCES telemetry_type(id),
    telemetry_id UUID NOT NULL,
    CONSTRAINT instrument_unique_telemetry_id UNIQUE(instrument_id, telemetry_id)
);

-- GOES
CREATE TABLE IF NOT EXISTS telemetry_goes (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nesdis_id VARCHAR UNIQUE NOT NULL
);

-- IRIDIUM
CREATE TABLE IF NOT EXISTS telemetry_iridium (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    imei VARCHAR(15) UNIQUE NOT NULL
);

-- ------
-- Config
-- ------

-- Config for Local Development (minio replicating S3 Storage)
INSERT INTO config (static_host, static_prefix) VALUES
    ('http://localhost', '/instrumentation');


INSERT INTO role (id, name) VALUES
    ('37f14863-8f3b-44ca-8deb-4b74ce8a8a69', 'ADMIN'),
    ('2962bdde-7007-4ba0-943f-cb8e72e90704', 'MEMBER');

-- -------
-- Domains
-- -------

-- instrument_type
INSERT INTO instrument_type (id, name) VALUES
    ('2f340c02-8bac-4bf4-9d5e-08a6f395ee8f', 'Adas'),
    ('3350b1d1-a946-49a8-bf19-587d7163e0f7', 'Barometer'),
    ('43671cb6-4141-45e4-9d17-b75a0e0508e1', 'Crackmeter'),
    ('6b8ef914-79ec-41d1-93bc-ba44694baa71', 'Drain'),
    ('03b950c9-7b53-408f-b874-b42353dc1ba7', 'Earth Pressure Cell'),
    ('69c8fcc2-0bbc-43b0-bdc2-a44f6ce5c1de', 'Extensometer'),
    ('11ffb059-e3d2-4cd7-a076-4fba7828bfbb', 'Extraction Well'),
    ('d65f62d0-57d7-49af-8078-d8f3fdb75477', 'Flowmeter'),
    ('135460f1-033f-46e0-b204-ed23074c0817', 'Flume'),
    ('3c3dfc23-ed2a-4a4a-9ce0-683c7c1d4d20', 'Inclinometer'),
    ('98a61f29-18a8-430a-9d02-0f53486e0984', 'Instrument'),
    ('1c8e8b3b-3322-4175-8d3a-c4c08ca1a86c', 'Joint Meter'),
    ('cb6a616e-4953-4807-a312-50622a57570a', 'Liquid Level Gauge'),
    ('39d299ee-6cf6-4924-bba3-39f20d713b0f', 'Load Cell'),
    ('486da6b4-f7a8-4c7f-a0cc-ee13c39673fb', 'Observation Well'),
    ('d9cf24af-bfef-45ec-8812-c76faa2b8feb', 'Pendulum'),
    ('1bb4bf7c-f5f8-44eb-9805-43b07ffadbef', 'Piezometer'),
    ('835bb7b9-5c80-48db-8c7f-06dfcacfc5d8', 'Plumbline'),
    ('f735d3e5-3741-4946-9913-0b5f178f8052', 'Pore Pressure Cell'),
    ('f371126d-fbe9-494d-869c-d55d8c393e65', 'Precipitation'),
    ('466b7603-0763-4aae-a67b-9bc78e630934', 'Relief Well'),
    ('c7c6f90b-1621-4bd6-95ea-e87cd8c326b1', 'Seismic Monitoring Device'),
    ('21f37666-0fb2-401d-bd8e-b9bcb8240b8d', 'Settlement Plate'),
    ('0fd1f9ba-2731-4ff9-96dd-3c03215ab06f', 'Staff Gage'),
    ('3fd6186e-39bd-4ec0-b57e-91da0b542d79', 'Strain Gauge'),
    ('daf1e5f4-32bf-4a3d-bb31-94ffc5c95436', 'Stress Cell'),
    ('4a91ca85-86c7-4055-8185-2155f5e60fd8', 'Surface Monitoring Point'),
    ('e5ab6dfe-185c-4a72-a493-f7e3844df3fd', 'Survey Monitoring Point'),
    ('0dabc138-265f-4278-87d7-412e3ec469ec', 'Thermister'),
    ('d98aa744-ac14-4cb5-adbc-d275a7498951', 'Tiltmeter'),
    ('78ade434-d845-4195-b1c6-5ab5b376be97', 'Water Level'),
    ('02a518e8-12de-4a0b-a07d-441989e920c1', 'Water Quality'),
    ('7e63a703-aa68-444f-8cd4-9f7d09cbcb83', 'Weir');

-- status
INSERT INTO status (id, name, description) VALUES
    ('e26ba2ef-9b52-4c71-97df-9e4b6cf4174d', 'active', 'description for active instrument status'),
    ('c9ee4acb-9623-4fde-bf36-7668afe463d4', 'inactive', 'description for inactive instrument status'),
    ('3d9add10-4418-4cb2-bd31-20a639504539', 'abandoned', 'description for abandoned instrument status'),
    ('94578354-ffdf-4119-9663-6bd4323e58f5', 'destroyed', 'description for destroyed instrument status'),
    ('03a2bf9a-bbd8-4031-8f4e-13e8c77807f1', 'lost', 'description for lost instrument status');

-- unit_family
INSERT INTO unit_family (id, name) VALUES
    ('c9f3b6d2-3136-4330-a330-66e402b4ee04','univ'),
    ('19ad5455-4d6a-47d3-a28a-87bdfac2d75c','metric'),
    ('c4eccc63-4bfb-4dd2-9f73-920ec7b385a0','english');

-- measure
INSERT INTO measure (id, name) VALUES
    ('b25b75e1-5b3d-451c-afc6-3db18f7f5bf8','acidity'),
    ('f3f9c804-5379-492e-b08e-d13b31598cd8','angle'),
    ('8f49aea3-c580-4813-9b56-f159397f367e','angular speed'),
    ('71daec34-d255-4dd5-8075-70ff93411389','area'),
    ('de6b0e4b-6199-4515-8ec1-0656ed50be55','areal volume rate'),
    ('900e7711-bb92-4c42-9a78-852e3b1c93c0','concentration'),
    ('2febeb7d-325e-431d-b482-16d319826341','conductance'),
    ('73d0dfa4-a526-4cb7-ba26-535663a1ac1d','conductivity'),
    ('54101d51-6c6c-4436-9495-8f4abf0b2271','count'),
    ('0a1d1d80-ad12-482e-a91e-bd905748b282','currency'),
    ('0cf92974-f0aa-49f5-9b66-f9fe4af714f7','electromotive potential'),
    ('3ce398f2-985f-4ed4-93f6-23595d1849b7','energy'),
    ('700ebff9-4997-4462-8d20-7fb770557ee2','flow'),
    ('cd70a8e1-b914-42ca-8554-40e3c72e54c2','force'),
    ('1351e7ed-9e7e-4b24-a081-e436206b8a6d','intensity'),
    ('98c548e8-caea-4042-b083-7ba1d4ab57d5','irradiance'),
    ('0ca977f6-a74c-4bb5-86c9-0166dba80034','irradiation'),
    ('2c2b39d2-186d-46e9-8dc7-aca36f03aa23','length'),
    ('d6391c73-b550-4422-b24b-f5de07a6d824','linear speed'),
    ('5042b88e-ad61-454d-953e-f07224be4f07','mass'),
    ('d3a8cd51-6318-4a61-90fa-148e3731cc33','mass concentration'),
    ('a2ee4309-db08-48b8-97c9-d424bf393510','ph'),
    ('e8a958ec-9bdf-4d9a-95db-27ed6eac8fb9','phase change rate index'),
    ('513bafdb-48e3-44ff-8189-54e8130ec76a','power'),
    ('c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e','pressure'),
    ('4b5611c1-a395-4561-b816-abdeaf416e2a','ratio'),
    ('955f207d-9a48-45c5-9697-60a4ad7c9ca4','temperature'),
    ('b85b9367-f034-4783-bf5b-9220e32d4e6a','time'),
    ('631347f1-b12e-4d09-8d1c-9359148d8b22','turbidity'),
    ('43fefa8b-10e9-4b27-8ed4-36e36174fbeb','unknown'),
    ('c70e7392-0108-4a17-a99f-244895f12558','velocity'),
    ('a08f3cd5-233d-43f4-8f21-633a6aa63f0c','voltage'),
    ('4f18ff01-aeba-46bd-b5c8-05eec8cd8b43','volume'),
    ('68bd29ef-5908-4776-bcda-8e41056707fd','volume rate');

-- unit
INSERT INTO unit (id, unit_family_id, measure_id, name, abbreviation) VALUES
('31e65a46-2a67-48f9-9b51-c49f0c82a5dc', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Acre-feet', 'ac-ft'),
('4a2b0a3d-5687-49b5-a31f-0b94bb28eaeb', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Calories', 'cal'),
('2d379b61-fef6-4531-9e58-cace51909e44', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '955f207d-9a48-45c5-9697-60a4ad7c9ca4', 'Centigrade', 'C'),
('3ba8644b-46d6-46d2-88ad-8abfb8c1d89e', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'Centimeters', 'cm'),
('a0b5e08c-0e67-45f8-9bc6-c2cede798516', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'de6b0e4b-6199-4515-8ec1-0656ed50be55', 'Cfs per square mile', 'cfs/mi2'),
('ea3f83cf-64fc-4858-91e7-a44b719965fd', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'de6b0e4b-6199-4515-8ec1-0656ed50be55', 'Cms per square kilometer', 'cms/km2'),
('4484c18a-61aa-48b4-8cf5-63d3b8c6d200', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '54101d51-6c6c-4436-9495-8f4abf0b2271', 'Count', 'unit'),
('f505adf3-1f8c-4f05-ae51-06d8ed4e1d9a', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Cubic feet', 'ft3'),
('a24cdecc-a4ec-4ca7-bc0c-f8fb6bb39e04', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Cubic kilometers', 'km3'),
-- ('7bffd219-f3cb-4f27-b270-1d23a6569e4b', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Cubic meters', 'm3'),
('bf9a00d8-084b-4006-b71a-602616f1f59a', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Cubic meters', 'M^3'),
('67d3c3f0-ae76-4807-8cdd-4e29fa8d8b39', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '68bd29ef-5908-4776-bcda-8e41056707fd', 'Cubic meters per second', 'cms'),
('5a2922f1-3553-4b0f-88ff-1ce12c47bf5c', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Cubic miles', 'mile3'),
('28da1a36-4a7a-4f82-b65d-39ad543189ac', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'f3f9c804-5379-492e-b08e-d13b31598cd8', 'Degrees', 'deg'),
('be854f6e-e36e-4bba-9e06-6d5aa09485be', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '1351e7ed-9e7e-4b24-a081-e436206b8a6d', 'Decibels', 'dB'),
('f2b8e84d-be4d-44d2-b772-a87e4c04cc46', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '0a1d1d80-ad12-482e-a91e-bd905748b282', 'Dollars', '$'),
('10e05b5c-7e96-434b-9182-a547333e1c52', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '955f207d-9a48-45c5-9697-60a4ad7c9ca4', 'Fahrenheit', 'F'),
('7d8e5bb9-b9ea-4920-9def-0589160ea412', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '631347f1-b12e-4d09-8d1c-9359148d8b22', 'Formazin Nephelometric Unit', 'FNU'),
('8a7afb4f-e0ac-409d-98ea-b025dce9b777', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '68bd29ef-5908-4776-bcda-8e41056707fd', 'Gallons per minute', 'gpm'),
('d58bdb43-2c78-4237-bb52-6d44016d186b', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Gigawatt-hours', 'GWh'),
('89022ea4-71e5-4be7-a6ac-6a3d87b9223c', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'Gigawatts', 'GW'),
('fa7deb5c-783e-44f1-8622-3da4a33a5795', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd3a8cd51-6318-4a61-90fa-148e3731cc33', 'Grams per cubic centimeters', 'gm/cm3'),
('f0151377-4ce0-42e3-b47e-acd81ce5558e', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd3a8cd51-6318-4a61-90fa-148e3731cc33', 'Grams per liter', 'g/l'),
('42029141-61fc-4280-b1cf-d3ad8a3b210e', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'Hectares', 'ha'),
('39b025ed-488d-42ab-93a2-08ae5cdc8ae0', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'Inches of mercury', 'in-hg'),
('f5652580-4e5f-4fdb-b80c-ef1ad46c9242', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Inches per day', 'in/day'),
('7a0c33bd-8c18-4e88-9bad-bb18e2b7ed38', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'e8a958ec-9bdf-4d9a-95db-27ed6eac8fb9', 'Inches per degree-day', 'in/deg-day'),
('2210ac17-04f6-4b29-ad7c-0cd9c40ffa2f', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Inches per hour', 'in/hr'),
('55f22541-65b4-4c3e-9a36-72d7a0fe2b1e', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '631347f1-b12e-4d09-8d1c-9359148d8b22', 'Jackson Turbitiy Unit', 'JTU'),
('baa937d4-eeff-4a90-bdd5-cdba08ef21f1', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '0ca977f6-a74c-4bb5-86c9-0166dba80034', 'Joules per square meters', 'J/m2'),
('ebe3c23b-e25f-42b1-ad2a-db752fccc6fa', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '68bd29ef-5908-4776-bcda-8e41056707fd', 'Kilo-cubic feet per second', 'kcfs'),
('b4653aaf-cd50-4589-bf1d-c1ded1247d30', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Kiloacre-feet', 'kaf'),
('6df4c80f-7af9-4f91-94da-a1e24cc8e3f6', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Kilocalories', 'kcal'),
('4df68cc3-2065-412b-af6b-fac5ccf1a087', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Kilogallons', 'kgal'),
('2f50063c-121b-4926-8793-49a9a1baac51', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'Kilometers', 'km'),
('5f31367e-be0f-44b0-8618-53476db34944', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Kilometers per hour', 'kph'),
('50f18926-2953-4335-a96b-7212c4b927ed', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'Kilopascals', 'kPa'),
('d2805505-4615-4a2e-948a-e0f49025b0a7', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Kilowatt-hours', 'kWh'),
('aa2fc05d-162f-486b-9fe3-a1b175c207bf', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '0ca977f6-a74c-4bb5-86c9-0166dba80034', 'Langley', 'langley'),
('f18629fd-5c40-4d6c-a38b-05a15b7a47bd', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '98c548e8-caea-4042-b083-7ba1d4ab57d5', 'Langley per minute', 'langley/min'),
('d76034bc-3853-4326-a494-f9acb9f308fb', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Megawatt-hours', 'MWh'),
('6b2011ff-795e-4139-a6aa-afb593bf15ca', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'Megawatts', 'MW'),
('ae06a7db-1e18-4994-be41-9d5a408d6cad', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'Meters', 'm'),
('e142d705-9eb6-4965-91d0-af55739189b0', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Meters per second', 'm/s'),
('c0fb3e31-60eb-4203-8755-e88c533bc61c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '5042b88e-ad61-454d-953e-f07224be4f07', 'Metric ton', 'mt'),
('7885df76-fde6-4490-9bd2-0dce06660076', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '2febeb7d-325e-431d-b482-16d319826341', 'Mhos', 'mho'),
('64b913ae-90a9-45a9-96bc-6ef317e30c1a', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'Micro Siemens per Centimeter', 'uS/cm'),
('f9ec4d95-4367-4774-8cf7-5d21de82ed31', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '2febeb7d-325e-431d-b482-16d319826341', 'Micro-Siemens', 'uS'),
('bece84d1-13bd-4e99-9bb1-c79e5274d751', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '2febeb7d-325e-431d-b482-16d319826341', 'Micro-mhos', 'umho'),
('633bd96c-5bdb-436f-b464-f18d90b7d736', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '73d0dfa4-a526-4cb7-ba26-535663a1ac1d', 'Micro-mhos per centimeter', 'umho/cm'),
-- ('1094e563-86dd-449f-adff-1c408a20bc1c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'MicroMHOs per centimeter', 'uMHOs'),
--('580a1b73-fb06-4d2e-998f-1771194ce0c4', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'MicroMHOs per centimeter', 'uMHO'),
('ac6ecf6e-7cc1-48bc-ba61-f878cbe7b2b1', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'MicroMHOs per centimeter', 'uMHOs/cm'),
('55dda9ef-7ba6-4432-b64d-8ef0e65154f4', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'Millibars', 'mb'),
('67d75ccd-6bf7-4086-a970-5ed65a5c30f3', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd3a8cd51-6318-4a61-90fa-148e3731cc33', 'Milligrams per liter', 'mg/l'),
('612d6fa5-954a-40fd-86d6-c21a75fe6cff', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'Millimeters', 'mm'),
('fc5f41e1-f5c6-453b-bda9-97530f400cc7', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'Millimeters of mercury', 'mm-hg'),
('7a6dd39d-7662-43fa-bb1f-1ed28f266b5b', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Millimeters per day', 'mm/day'),
('4656ae18-596b-488c-94c5-c07f36e221cd', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'e8a958ec-9bdf-4d9a-95db-27ed6eac8fb9', 'Millimeters per degree-day', 'mm/deg-day'),
('c3beaa02-38aa-4925-b444-280ae847ded1', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'd6391c73-b550-4422-b24b-f5de07a6d824', 'Millimeters per hour', 'mm/hr'),
('cffd4714-3378-4c3d-96ed-2d68c11223d7', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Millions of gallons', 'mgal'),
('576d54f5-c90c-40e5-833f-7e2e9745c00d', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '68bd29ef-5908-4776-bcda-8e41056707fd', 'Millions of gallons per day', 'mgd'),
('e65274a5-3d42-4b96-8db6-696d65d92a8d', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '631347f1-b12e-4d09-8d1c-9359148d8b22', 'Nephelometric Turbidity Unit', 'NTU'),
('fc3e2065-3815-4078-827d-bee5cd714955', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'cd70a8e1-b914-42ca-8554-40e3c72e54c2', 'Pounds', 'lb'),
('3254f483-5e66-405c-acf2-2a8add714bf5', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '43fefa8b-10e9-4b27-8ed4-36e36174fbeb', 'Raw Undefined Units', 'raw'),
('565197e0-3bb4-4f7c-a55c-eabbc51c3a29', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'f3f9c804-5379-492e-b08e-d13b31598cd8', 'Revolution', 'rev'),
('191b914b-b4ad-43ca-9680-d4d44ddc12a6', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '8f49aea3-c580-4813-9b56-f159397f367e', 'Revolutions per minute', 'rpm'),
('cdf9a07a-99c9-4cfd-a87b-242407ce56c9', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '2febeb7d-325e-431d-b482-16d319826341', 'Siemens', 'S'),
('6b4727ea-dc34-4699-8a88-d7da04daa2e3', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'Square feet', 'ft2'),
('fe3715f0-bcaa-4b6e-ac68-5677096ab902', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'Square kilometers', 'km2'),
('bf91e9d0-f429-44f5-9774-ea7212ba8113', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'Square meters', 'm2'),
('0d179de1-188c-441c-9f04-f89008a8305b', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'Square miles', 'mile2'),
('5939c8bb-f924-45f0-9ec9-5673ae7f862c', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'a2ee4309-db08-48b8-97c9-d424bf393510', 'Standard pH units', 'su'),
('0971c924-d7ef-4a4a-b73b-8718a91ba415', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'f3f9c804-5379-492e-b08e-d13b31598cd8', 'Tens of Degrees', 'deg*10'),
('07f336b6-aa15-4d5e-af7a-5f2fcdc2001d', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Terawatt-hour', 'TWh'),
('c4a66be8-6bdf-4ded-8199-0233813fe777', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'Terawatts', 'TW'),
('c768c0c8-36e4-495a-bb9a-2065402db3c5', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'Thousands of cubic meters', '1000 m3'),
('3383d7d4-ffa1-4522-a8f5-f16561c5bd2f', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'Thousands of square meters', '1000 m2'),
('6b5bd788-8c78-43bb-b5a3-ad544b858a64', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'a08f3cd5-233d-43f4-8f21-633a6aa63f0c', 'Volts', 'V'),
('a0be2c0a-e6e7-41c1-9417-91f6a4d2f8ea', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'Watt-hours', 'Wh'),
('3008e1ff-338b-4072-865b-0ff68e0d68b6', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '98c548e8-caea-4042-b083-7ba1d4ab57d5', 'Watts per square meter', 'W/m2'),
('23aa81c1-74a0-4186-a481-1fd2f146986e', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'acre feet', 'acre*ft'),
('81451270-cce7-49b3-9d1c-cf7b1e557602', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'acres', 'acre'),
('bcc09b92-e21b-4be8-a76f-aa418764a83c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'atmospheres', 'atm'),
('3e6ebdd4-dc07-4854-8057-47a4b11da141', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'bars', 'bar'),
('76985877-e19f-49b5-87e2-982fc03d77d7', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'british thermal unit', 'btu'),
('2e596176-a176-4f1f-b9d4-7bd03f283eda', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'centimeters', 'cM'),
('aaa20163-ef13-4956-9ef4-76f6273954d5', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'centimeters per second', 'cM/s'),
('abc1c2e0-64a7-468c-9684-15569c33d56e', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '54101d51-6c6c-4436-9495-8f4abf0b2271', 'counts', 'count'),
('b8c28057-9751-4850-be45-d1ab921faf6a', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'cubic centimeter', 'cc'),
('71e15fb3-84a6-4aa2-a7d1-f8a6cfe2d9ef', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'cubic feet', 'ft^3'),
('d23b3b3b-69dc-4753-8b59-01848f027408', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '700ebff9-4997-4462-8d20-7fb770557ee2', 'cubic feet per second', 'cfs'),
-- ('02e3640d-7428-4383-b4b3-59ef97c747cb', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '700ebff9-4997-4462-8d20-7fb770557ee2', 'cubic feet per second', 'ft^3/s'),
('259c1ab3-4f70-4279-b698-b28bb1aaa009', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'cubic inches', 'in^3'),
('c4d57208-54ea-420f-802b-8de8698f0616', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'cubic meter', 'm^3'),
('66239345-fa21-48cd-922a-7304fc3cffa2', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '700ebff9-4997-4462-8d20-7fb770557ee2', 'cubic meters per second', 'm^3/s'),
('1c9e2b24-ec6f-49a6-b18a-7ca1a9a4a79d', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'day-second-foot', 'dsf'),
('3ac8944a-99d4-4ef2-874f-667983b63e6e', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'days', 'day'),
('6462733b-5b42-46a2-ad44-882a5332eafc', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '955f207d-9a48-45c5-9697-60a4ad7c9ca4', 'degrees Celsius', 'degC'),
('daeee256-c762-43a2-8369-2d295525023c', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '955f207d-9a48-45c5-9697-60a4ad7c9ca4', 'degrees Fahrenheit', 'degF'),
('959e4dcb-c871-4b85-bc07-1b5ab644962a', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '955f207d-9a48-45c5-9697-60a4ad7c9ca4', 'degrees Kelvin', 'degK'),
('877042fd-743e-43c0-b637-6c7c34f1875c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'cd70a8e1-b914-42ca-8554-40e3c72e54c2', 'dyn', 'dyn'),
('83ca3059-528d-497c-8944-4dff0ce6239b', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'ergs', 'erg'),
('f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'feet', 'ft'),
('1924c73e-591e-4d08-bd7d-cbd46d555b9b', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'feet per second', 'ft/s'),
('def3ff33-dae6-4a67-b24c-9a1b8bb463f8', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'fluid ounce', 'floz'),
('d84334e2-d99f-4708-b158-f80a40bf820a', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'foot-pounds per second', 'ft*lbf/s'),
('7c5b5dc4-8c71-43f1-8ae5-914ee21d1ccc', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'gallon', 'gal'),
('a063b3d6-c476-4dc6-8db1-2e4320bbfc02', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '5042b88e-ad61-454d-953e-f07224be4f07', 'grams', 'G'),
('6a8de16e-05ab-40ea-a51a-e8bc5fdb987b', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '900e7711-bb92-4c42-9a78-852e3b1c93c0', 'grams per liter', 'g/L'),
('82dbaa25-b2cc-400c-b258-fe1200c85141', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'horsepower', 'hp'),
('eed52efd-ec26-46ba-98de-582a22ec4ed4', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'hours', 'hr'),
('4ee79a3d-a053-41b8-85b5-bb2eea3c9d1a', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'inches', 'in'),
('3ccd658d-f656-4fb4-b2b9-913ab7279ada', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'inches of mercury', 'inHg'),
('063497fb-7693-47c5-b9a6-fdfac8dd7562', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'inches per second', 'in/s'),
('95a1e3e2-adf8-4658-85cc-fb86211be4bf', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'joules', 'j'),
('6f23f255-a323-401f-97da-2f2d831c1fa0', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '5042b88e-ad61-454d-953e-f07224be4f07', 'kilograms', 'kG'),
('ab8f267d-6541-43cb-9ee2-8f7ae4d2b056', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '3ce398f2-985f-4ed4-93f6-23595d1849b7', 'kilojoules', 'kj'),
('e21b6ee7-de8a-4f6b-ae18-aea85429914c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'kiloliter', 'kL'),
('aaf6631a-51a0-478a-be9c-997ffb5a3beb', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'kilometers', 'kM'),
-- ('7aa4f9cf-fc1e-48a3-9137-b7263e434626', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'kilometers per hour', 'kM/hr'),
-- ('8576c542-d217-4f40-8c5c-6661affb89ea', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'kilometers per hour', 'kph'),
('b3e51265-85df-44c3-8ad6-b04a902bd755', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'kilopascals', 'kpa'),
('6dbcab30-56b7-4dd1-b5e1-6da70f4df3da', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'kilowatts', 'kW'),
('7e2b9278-c291-4e08-b375-1fe9377ba535', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'liter', 'L'),
('4b503088-fab3-46cf-a5e1-44bf9fbf6ad9', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'meters', 'M'),
('c96294dc-f238-4d4d-8705-0a1b2d3f9b55', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'meters per second', 'M/s'),
('7ab9cb43-ac59-457d-8189-ee35fc0eed9e', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '5042b88e-ad61-454d-953e-f07224be4f07', 'micrograms', 'uG'),
('ff04a299-568f-4f7e-88f0-d926543bd800', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '900e7711-bb92-4c42-9a78-852e3b1c93c0', 'micrograms per liter', 'uG/L'),
('af03fe6f-b203-48fd-8fb0-a985fbe188a0', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'microliter', 'uL'),
('84d8a118-71e3-4541-a7a1-8d4cac9e7b35', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'micrometers', 'uM'),
('904c90a3-3e96-49fa-9b0d-778dadbe2bc0', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'miles', 'mi'),
('fb756f9f-0132-4e84-ac71-524d9a9c4164', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'miles per hour', 'mph'),
-- ('22e63820-677f-4d75-9db0-8c84eb037348', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'miles per hour', 'mi/hr'),
('4aebf158-1306-4b7d-9e35-9e5998414bba', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'millibars', 'mbar'),
('362bee5d-2157-4730-b2f4-bd8d06c374cd', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '5042b88e-ad61-454d-953e-f07224be4f07', 'milligrams', 'mG'),
('d4b90d1d-bffb-4b56-8de8-abcb5826b6c4', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '900e7711-bb92-4c42-9a78-852e3b1c93c0', 'milligrams per liter', 'mG/L'),
('3562948d-ff27-4604-8a24-0240eb946682', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'milliliter', 'mL'),
('22c1bb95-4e0c-4e53-a24d-c1327fd11575', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'millimeters', 'mM'),
('6407a23f-b5f8-4214-9343-50b6231e4bfe', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'millimeters of mercury', 'mmHg'),
('bdff5fd3-4b67-40b1-8ce3-8ff2733d74a1', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'millimeters per second', 'mM/s'),
('72a86938-dd9b-4590-9770-4d6fa76e1f4d', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'minutes', 'min'),
('cfcf408b-9193-4b80-99ce-1171554c9428', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'nautical miles', 'nmi'),
('9564a419-4cd7-4ca2-87e4-43cc299e6917', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c70e7392-0108-4a17-a99f-244895f12558', 'nautical miles per hour', 'knots'),
-- ('4415422f-d2e0-4e43-801f-893e2230e2fa', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'nautical miles per hour', 'nmi/hr'),
('5a9eba40-82ce-4b7a-b010-2258a691db17', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'cd70a8e1-b914-42ca-8554-40e3c72e54c2', 'newtons', 'N'),
('c6a6a797-ee49-46fe-b67a-727c718a51ad', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '5042b88e-ad61-454d-953e-f07224be4f07', 'ounces', 'oz'),
('cfac3e61-64e1-456d-890e-0655038e8218', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b25b75e1-5b3d-451c-afc6-3db18f7f5bf8', 'pH', 'pH'),
('64cf271d-ba2f-426f-b98d-63cb93fe72f3', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '4b5611c1-a395-4561-b816-abdeaf416e2a', 'parts per million', 'ppm'),
('b57b6f9a-29a3-426b-9c31-5fa37f0bd2ad', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '4b5611c1-a395-4561-b816-abdeaf416e2a', 'parts per thousand', 'ppt'),
('41d4f482-7b11-4761-998e-ffd5cc901e07', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'pascals', 'pa'),
('d43054bc-6cf7-440e-a262-f9f38da99841', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '4b5611c1-a395-4561-b816-abdeaf416e2a', 'percent', 'pct'),
-- ('e99f5559-ecb6-4efc-bde3-a0b9c8bf3711', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '4b5611c1-a395-4561-b816-abdeaf416e2a', 'percent', '%'),
('19cb5112-0393-47b7-bc38-cb59a5d925d1', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'pint', 'pt'),
('2b81c09e-818f-46a8-8ec6-dafbef690118', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'cd70a8e1-b914-42ca-8554-40e3c72e54c2', 'pound-force', 'lbf'),
-- ('93045059-9465-41ed-9015-7689327bd8e0', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '5042b88e-ad61-454d-953e-f07224be4f07', 'pounds', 'lb'),
('7c3f5f46-d7a1-4cf7-be9f-d82c11398a1d', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c5e6a255-b9c2-4f9b-9a5b-f1168f645e5e', 'pounds per square inch', 'psi'),
('f4367827-9622-40b9-9ff1-711e14252c05', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '4f18ff01-aeba-46bd-b5c8-05eec8cd8b43', 'quart', 'qt'),
('93229888-9513-4d66-9ab8-9fb6e37b65d8', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'second', 'sec'),
('81cb1c58-e050-46ba-b510-71dad374192a', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'square centimeters', 'cM^2'),
('49a95b0b-6688-4fd1-b22a-87a6bee39081', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'square feet', 'ft^2'),
('d4a9183d-b333-4b45-a23d-549a5609788e', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'square inches', 'in^2'),
('f3ec6179-8399-41b7-9ce3-049336e8a4fe', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'square kilometers', 'kM^2'),
('61b725e9-a741-4eb9-b34a-57c9dc4853de', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'square meters', 'M^2'),
('0fefe6f8-ab65-4ce4-86a6-b223980f06f2', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'square miles', 'mi^2'),
('290f1547-a3c5-4158-a088-b50f5cabe8be', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '71daec34-d255-4dd5-8075-70ff93411389', 'square millimeters', 'mM^2'),
('e0b63a59-6da5-43d8-8e1b-f9d50078eec6', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '71daec34-d255-4dd5-8075-70ff93411389', 'square yards', 'yd^2'),
('0b2ecf01-9034-4043-953e-ae20f0e8c935', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '5042b88e-ad61-454d-953e-f07224be4f07', 'tons', 'ton'),
('c0a116ef-058d-41b0-a845-557226ce557c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'watts', 'W'),
('5fa61c67-38e6-46ae-ac1f-114278706261', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'weeks', 'week'),
('cc83a42b-16a7-46a8-b3a6-966bad7ae2d7', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'yards', 'yd'),
('1292b2a5-b78e-4a7a-80e3-978d44cbff2b', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'yards per second', 'yd/s'),
('4a999277-4cf5-4282-93ce-23b33c65e2c8', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', '43fefa8b-10e9-4b27-8ed4-36e36174fbeb', 'unknown', 'unknown');

-- parameter
INSERT INTO parameter (id, name) VALUES
    ('b4ea8385-48a3-4e95-82fb-d102dfcbcb54', 'air-temperature'),
    ('377ecec0-f785-46ab-b0e2-5fd8c682dfea', 'conductivity'),
    ('98007857-d027-4524-9a63-d07ae93e5fa2', 'dissolved-oxygen'),
    ('83b5a1f7-948b-4373-a47c-d73ff622aafd', 'elevation'),
    ('a63a3202-3115-4ad4-9e5b-3d35f94647d2', 'flow'),
    ('068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'length'),
    ('5d0b2c85-6a4c-4d82-aed3-193b066349f1', 'ph'),
    ('0ce77a5a-8283-47cd-9126-c440bcec4ef6', 'precipitation'),
    ('1de79e29-fb70-45c3-ae7d-4695517ced90', 'pressure'),
    ('b23b141d-ce7b-4e0d-82ab-c8beb39c8325', 'signal-strength'),
    ('e46deb1d-e7e4-4d49-a874-18306991ecfe', 'speed'),
    ('b49f214e-f69f-43da-9ce3-ad96042268d0', 'stage'),   
    ('3676df6a-37c2-4a81-9072-ddcd4ab93702', 'turbidity'),
    ('2b7f96e1-820f-4f61-ba8f-861640af6232', 'unknown'),
    ('06189199-a25f-4101-b8bd-991c6a5a7ab3', 'velocity'),  
    ('430e5edb-e2b5-4f86-b19f-cda26a27e151', 'voltage'),
    ('de6112da-8489-4286-ae56-ec72aa09974d', 'water-temperature'),
    ('3ea5ed77-c926-4696-a580-a3fde0f9a556', 'inclinometer-constant');

-- Profile (MIDAS Automation Account)
INSERT INTO profile (edipi, username, email) VALUES (79, 'MIDAS Automation', 'midas@rsgis.dev');

-- -------------------------------------------------
-- basic seed data to demo the app and run API tests
-- -------------------------------------------------

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
INSERT INTO instrument_group (project_id, id, slug, name, description) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'd0916e8a-39a6-4f2f-bd31-879881f8b40c', 'sample-instrument-group', 'Sample Instrument Group 1', 'This is an example instrument group');

-- instrument
INSERT INTO instrument (project_id, id, slug, name, geometry, type_id) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'a7540f69-c41e-43b3-b655-6e44097edb7e', 'demo-piezometer-1', 'Demo Piezometer 1', ST_GeomFromText('POINT(-80.8 26.7)',4326),'1bb4bf7c-f5f8-44eb-9805-43b07ffadbef'),
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '9e8f2ca4-4037-45a4-aaca-d9e598877439', 'demo-staffgage-1', 'Demo Staffgage 1', ST_GeomFromText('POINT(-80.85 26.75)',4326),'0fd1f9ba-2731-4ff9-96dd-3c03215ab06f'),
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', 'inclinometer-1', 'inclinometer-1', ST_GeomFromText('POINT(-80.8 26.7)',4326),'98a61f29-18a8-430a-9d02-0f53486e0984');

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
('5b6f4f37-7755-4cf9-bd02-94f1e9bc5986', 'd8c66ef9-06f0-4d52-9233-f3778e0624f0', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', '4a999277-4cf5-4282-93ce-23b33c65e2c8', 'inclinometer-1.formula', 'inclinometer-1');

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
                            {
                              "depth": 106, 
                              "a0": 590,
                              "a180": -562,
                              "b0": -142,
                              "b180": 176
                            },
                            {
                              "depth": 108, 
                              "a0": 614,
                              "a180": -586,
                              "b0": 107,
                              "b180": -149
                            },
                            {
                              "depth": 110, 
                              "a0": 622,
                              "a180": -592,
                              "b0": -67,
                              "b180": 107
                            },
                            {
                              "depth": 112, 
                              "a0": 623,
                              "a180": -598,
                              "b0": 8,
                              "b180": -48
                            },
                            {
                              "depth": 114, 
                              "a0": 606,
                              "a180": -577,
                              "b0": 124,
                              "b180": -72
                            },
                            {
                              "depth": 116, 
                              "a0": 0,
                              "a180": 0,
                              "b0": 0,
                              "b180": 0
                            }
                      ]');
                      
-- alert_config
INSERT INTO alert_config (id, instrument_id, name, body, formula, schedule) VALUES
    ('1efd2d85-d3ee-4388-85a0-f824a761ff8b', '9e8f2ca4-4037-45a4-aaca-d9e598877439','Above Target Height', 'The demo staff gage has exceeded the target height. Sincerely, Midas', '[stage] >= 10', '0,10,20,30,40,50 * * * *'),
    ('243e9d32-2cba-4f12-9abe-63adc09fc5dd', 'a7540f69-c41e-43b3-b655-6e44097edb7e','Below Target Height', 'Distance to water is near artesian conditions. Sincerely, Midas', '[distance-to-water] <= 2', '0,10,20,30,40,50 * * * *'),
    ('6f3dfe9f-4664-4c78-931f-32ffac6d2d43', 'a7540f69-c41e-43b3-b655-6e44097edb7e','Sample Demo Alert', 'Sample Alert Condition Has Been Triggered. Sincerely, Midas', '1 == 1', '0,10,20,30,40,50 * * * *');

-- alert
INSERT INTO alert (id, alert_config_id) VALUES ('e070be13-ef17-40f3-99c8-fef3ee1b9fb5', '6f3dfe9f-4664-4c78-931f-32ffac6d2d43');

-- collection_group
INSERT INTO collection_group (id, project_id, name, slug) VALUES
    ('1519eaea-1799-4375-aa37-0e35aa654643', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'Manual Collection Route 1', 'manual-collection-route-1'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'High Water Inspection', 'high-water-inspection');

-- collection_group_timeseries
INSERT INTO collection_group_timeseries (collection_group_id, timeseries_id) VALUES
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '7ee902a3-56d0-4acf-8956-67ac82c03a96'),
    ('30b32cb1-0936-42c4-95d1-63a7832a57db', '9a3864a8-8766-4bfa-bad1-0328b166f6a8');

-- plot_configuration
INSERT INTO plot_configuration (project_id, id, slug, name) VALUES
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', 'cc28ca81-f125-46c6-a5cd-cc055a003c19', 'all-plots', 'All Plots'),
    ('5b6f4f37-7755-4cf9-bd02-94f1e9bc5984', '64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a', 'pz-1a-plot', 'PZ-1A PLOT');


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
