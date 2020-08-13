-- extensions
CREATE extension IF NOT EXISTS "uuid-ossp";


-- drop tables if they already exist
drop table if exists 
    public.timeseries_measurement,
    public.timeseries,
    public.instrument_group_instruments,
    public.instrument_status,
    public.instrument_note,
    public.instrument,
    public.instrument_group,
    public.parameter,
    public.unit_family,
    public.measure,
    public.unit,
    public.instrument_type,
    public.project,
    public.status
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

-- measure
CREATE TABLE IF NOT EXISTS public.measure (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(240) UNIQUE NOT NULL
);

-- unit_family
CREATE TABLE IF NOT EXISTS public.unit_family (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL
);

-- unit
CREATE TABLE IF NOT EXISTS public.unit (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(120) UNIQUE NOT NULL,
    abbreviation VARCHAR(120) UNIQUE NOT NULL,
    unit_family_id UUID REFERENCES unit_family (id),
    measure_id UUID REFERENCES measure (id)
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

-- -------
-- Domains
-- -------

-- instrument_type
INSERT INTO instrument_type (id, name) VALUES
    ('0fd1f9ba-2731-4ff9-96dd-3c03215ab06f', 'Staff Gage'),
    ('1bb4bf7c-f5f8-44eb-9805-43b07ffadbef', 'Piezometer'),
    ('3350b1d1-a946-49a8-bf19-587d7163e0f7', 'Barometer');

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
    ('43fefa8b-10e9-4b27-8ed4-36e36174fbeb','undefined'),
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
('580a1b73-fb06-4d2e-998f-1771194ce0c4', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'MicroMHOs per centimeter', 'uMHO'),
-- ('ac6ecf6e-7cc1-48bc-ba61-f878cbe7b2b1', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '2febeb7d-325e-431d-b482-16d319826341', 'MicroMHOs per centimeter', 'uMHOs/cm'),
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
('6b5bd788-8c78-43bb-b5a3-ad544b858a64', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '0cf92974-f0aa-49f5-9b66-f9fe4af714f7', 'Volts', 'volt'),
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
('0fcf6088-7a05-43d3-bec6-5e825a68a2a9', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', 'a08f3cd5-233d-43f4-8f21-633a6aa63f0c', 'volts', 'V'),
('c0a116ef-058d-41b0-a845-557226ce557c', '19ad5455-4d6a-47d3-a28a-87bdfac2d75c', '513bafdb-48e3-44ff-8189-54e8130ec76a', 'watts', 'W'),
('5fa61c67-38e6-46ae-ac1f-114278706261', 'c9f3b6d2-3136-4330-a330-66e402b4ee04', 'b85b9367-f034-4783-bf5b-9220e32d4e6a', 'weeks', 'week'),
('cc83a42b-16a7-46a8-b3a6-966bad7ae2d7', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', '2c2b39d2-186d-46e9-8dc7-aca36f03aa23', 'yards', 'yd'),
('1292b2a5-b78e-4a7a-80e3-978d44cbff2b', 'c4eccc63-4bfb-4dd2-9f73-920ec7b385a0', 'c70e7392-0108-4a17-a99f-244895f12558', 'yards per second', 'yd/s');

-- parameter
INSERT INTO parameter (id, name) VALUES
    ('068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'length'),
    ('1de79e29-fb70-45c3-ae7d-4695517ced90', 'pressure'),
    ('b49f214e-f69f-43da-9ce3-ad96042268d0', 'stage'),
    ('de6112da-8489-4286-ae56-ec72aa09974d', 'temperature'),
    ('0ce77a5a-8283-47cd-9126-c440bcec4ef6', 'precipitation'),
    ('83b5a1f7-948b-4373-a47c-d73ff622aafd', 'elevation'),
    ('430e5edb-e2b5-4f86-b19f-cda26a27e151', 'voltage');

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
('7ee902a3-56d0-4acf-8956-67ac82c03a96', 'a7540f69-c41e-43b3-b655-6e44097edb7e', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'height', 'Height'),
('8f4ca3a3-5971-4597-bd6f-332d1cf5af7c', '9e8f2ca4-4037-45a4-aaca-d9e598877439', '068b59b0-aafb-4c98-ae4b-ed0365a6fbac', 'f777f2e2-5e32-424e-a1ca-19d16cd8abce', 'height-1', 'Height');

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
