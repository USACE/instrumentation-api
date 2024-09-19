CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "btree_gist";
CREATE EXTENSION IF NOT EXISTS "unaccent";

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
    show_masked BOOLEAN NOT NULL DEFAULT true,
    show_nonvalidated BOOLEAN NOT NULL DEFAULT true,
    show_comments BOOLEAN NOT NULL DEFAULT true,
    auto_range BOOLEAN NOT NULL DEFAULT true,
    date_range VARCHAR NOT NULL DEFAULT '1 year',
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

CREATE TABLE IF NOT EXISTS aware_platform (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    aware_id UUID NOT NULL,
    instrument_id UUID REFERENCES instrument(id)
);

CREATE TABLE IF NOT EXISTS aware_parameter (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    key VARCHAR NOT NULL,
    parameter_id UUID NOT NULL REFERENCES parameter(id),
    unit_id UUID NOT NULL REFERENCES unit(id),
    timeseries_slug VARCHAR NOT NULL,
    timeseries_name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS aware_platform_parameter_enabled (
    aware_platform_id UUID NOT NULL REFERENCES aware_platform(id),
    aware_parameter_id UUID NOT NULL REFERENCES aware_parameter(id),
    CONSTRAINT aware_platform_unique_parameter UNIQUE(aware_platform_id, aware_parameter_id)
);
