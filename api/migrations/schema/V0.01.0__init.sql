CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "btree_gist";
CREATE EXTENSION IF NOT EXISTS "unaccent";


CREATE TYPE job_status AS ENUM ('SUCCESS', 'FAIL', 'INIT');
CREATE TYPE line_style AS ENUM ('solid', 'dot', 'dash', 'longdash', 'dashdot', 'longdashdot');
CREATE TYPE plot_type AS ENUM ('scatter-line', 'profile', 'contour', 'bullseye');
CREATE TYPE timeseries_type AS ENUM ('standard', 'constant', 'computed', 'cwms');
CREATE TYPE trace_type AS ENUM ('bar', 'scattergl');
CREATE TYPE y_axis AS ENUM ('y1', 'y2');


CREATE TABLE telemetry_type (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  slug varchar NOT NULL,
  name varchar NOT NULL,
  CONSTRAINT telemetry_type_name_key UNIQUE (name),
  CONSTRAINT telemetry_type_pkey PRIMARY KEY (id),
  CONSTRAINT telemetry_type_slug_key UNIQUE (slug)
);


CREATE TABLE instrument_type (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(120) NOT NULL,
  icon text,
  CONSTRAINT instrument_type_name_key UNIQUE (name),
  CONSTRAINT instrument_type_pkey PRIMARY KEY (id)
);


CREATE TABLE status (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(20) NOT NULL,
  description varchar(480),
  CONSTRAINT status_name_key UNIQUE (name),
  CONSTRAINT status_pkey PRIMARY KEY (id)
);


CREATE TABLE measure (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(240) NOT NULL,
  CONSTRAINT measure_name_key UNIQUE (name),
  CONSTRAINT measure_pkey PRIMARY KEY (id)
);


CREATE TABLE unit_family (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(120) NOT NULL,
  CONSTRAINT unit_family_name_key UNIQUE (name),
  CONSTRAINT unit_family_pkey PRIMARY KEY (id)
);


CREATE TABLE parameter (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(120) NOT NULL,
  CONSTRAINT parameter_name_key UNIQUE (name),
  CONSTRAINT parameter_pkey PRIMARY KEY (id)
);


CREATE TABLE agency (
  id uuid NOT NULL,
  name text NOT NULL,
  CONSTRAINT agency_id_key UNIQUE (id),
  CONSTRAINT agency_name_key UNIQUE (name)
);


CREATE TABLE alert (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  alert_config_id uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  CONSTRAINT alert_pkey PRIMARY KEY (id)
);


CREATE TABLE config (
  static_host varchar DEFAULT 'http://minio:9000'::varchar NOT NULL,
  static_prefix varchar DEFAULT '/instrumentation'::varchar NOT NULL
);


CREATE TABLE role (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar NOT NULL,
  deleted boolean DEFAULT false NOT NULL,
  CONSTRAINT role_pkey PRIMARY KEY (id)
);


CREATE TABLE profile (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  edipi bigint NOT NULL,
  username varchar(240) NOT NULL,
  email varchar(240) NOT NULL,
  is_admin boolean DEFAULT false NOT NULL,
  display_name text NOT NULL,
  CONSTRAINT profile_edipi_key UNIQUE (edipi),
  CONSTRAINT profile_email_key UNIQUE (email),
  CONSTRAINT profile_pkey PRIMARY KEY (id),
  CONSTRAINT profile_username_key UNIQUE (username)
);


CREATE TABLE email (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  email varchar(240) NOT NULL,
  CONSTRAINT email_email_key UNIQUE (email),
  CONSTRAINT email_pkey PRIMARY KEY (id),
  CONSTRAINT unique_email UNIQUE (email)
);


CREATE TABLE office (
  id uuid NOT NULL,
  CONSTRAINT office_pkey PRIMARY KEY (id)
);


CREATE TABLE alert_type (
  id uuid NOT NULL,
  name text NOT NULL,
  CONSTRAINT alert_type_name_key UNIQUE (name),
  CONSTRAINT alert_type_pkey PRIMARY KEY (id)
);


CREATE TABLE submittal_status (
  id uuid NOT NULL,
  name text NOT NULL,
  CONSTRAINT alert_status_name_key UNIQUE (name),
  CONSTRAINT alert_status_pkey PRIMARY KEY (id)
);


CREATE TABLE telemetry_goes (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  nesdis_id varchar NOT NULL,
  CONSTRAINT telemetry_goes_nesdis_id_key UNIQUE (nesdis_id),
  CONSTRAINT telemetry_goes_pkey PRIMARY KEY (id)
);


CREATE TABLE telemetry_iridium (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  imei varchar(15) NOT NULL,
  CONSTRAINT telemetry_iridium_imei_key UNIQUE (imei),
  CONSTRAINT telemetry_iridium_pkey PRIMARY KEY (id)
);


CREATE TABLE datalogger_model (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  model text,
  CONSTRAINT datalogger_model_pkey PRIMARY KEY (id)
);


CREATE TABLE heartbeat (
  "time" timestamptz DEFAULT now() NOT NULL
);


CREATE TABLE unit (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name varchar(120) NOT NULL,
  abbreviation varchar(120) NOT NULL,
  unit_family_id uuid,
  measure_id uuid,
  CONSTRAINT unit_abbreviation_key UNIQUE (abbreviation),
  CONSTRAINT unit_name_key UNIQUE (name),
  CONSTRAINT unit_pkey PRIMARY KEY (id),
  CONSTRAINT unit_measure_id_fkey FOREIGN KEY (measure_id) REFERENCES measure(id),
  CONSTRAINT unit_unit_family_id_fkey FOREIGN KEY (unit_family_id) REFERENCES unit_family(id)
);


CREATE TABLE division (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  name text,
  initials varchar(3),
  agency_id uuid NOT NULL,
  CONSTRAINT division_pkey PRIMARY KEY (id),
  CONSTRAINT division_agency_id_fkey FOREIGN KEY (agency_id) REFERENCES agency(id)
);


CREATE TABLE instrument (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  deleted boolean DEFAULT false NOT NULL,
  slug varchar NOT NULL,
  name varchar(360) NOT NULL,
  geometry geometry(Geometry,4326),
  station integer,
  station_offset integer,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  type_id uuid NOT NULL,
  nid_id varchar,
  usgs_id varchar,
  show_cwms_tab boolean DEFAULT false NOT NULL,
  CONSTRAINT instrument_pkey PRIMARY KEY (id),
  CONSTRAINT instrument_slug_key UNIQUE (slug),
  CONSTRAINT instrument_type_id_fkey FOREIGN KEY (type_id) REFERENCES instrument_type(id)
);


CREATE TABLE profile_token (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  token_id varchar NOT NULL,
  profile_id uuid NOT NULL,
  issued timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  hash varchar(240) NOT NULL,
  CONSTRAINT profile_token_pkey PRIMARY KEY (id),
  CONSTRAINT profile_token_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profile(id)
);


CREATE TABLE alert_read (
  alert_id uuid NOT NULL,
  profile_id uuid NOT NULL,
  CONSTRAINT profile_unique_alert_read UNIQUE (alert_id, profile_id),
  CONSTRAINT alert_read_alert_id_fkey FOREIGN KEY (alert_id) REFERENCES alert(id),
  CONSTRAINT alert_read_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profile(id)
);


CREATE TABLE alert_profile_subscription (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  alert_config_id uuid NOT NULL,
  profile_id uuid NOT NULL,
  mute_ui boolean DEFAULT false NOT NULL,
  mute_notify boolean DEFAULT false NOT NULL,
  CONSTRAINT alert_profile_subscription_pkey PRIMARY KEY (id),
  CONSTRAINT profile_unique_alert_config UNIQUE (profile_id, alert_config_id),
  CONSTRAINT alert_profile_subscription_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profile(id)
);


CREATE TABLE alert_email_subscription (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  alert_config_id uuid NOT NULL,
  email_id uuid NOT NULL,
  mute_notify boolean DEFAULT false NOT NULL,
  CONSTRAINT alert_email_subscription_pkey PRIMARY KEY (id),
  CONSTRAINT email_unique_alert_config UNIQUE (email_id, alert_config_id),
  CONSTRAINT alert_email_subscription_email_id_fkey FOREIGN KEY (email_id) REFERENCES email(id)
);


CREATE TABLE district (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  division_id uuid NOT NULL,
  name text,
  initials varchar(3),
  office_id uuid,
  CONSTRAINT district_pkey PRIMARY KEY (id),
  CONSTRAINT district_division_id_fkey FOREIGN KEY (division_id) REFERENCES division(id),
  CONSTRAINT district_office_id_fkey FOREIGN KEY (office_id) REFERENCES office(id)
);


CREATE TABLE timeseries (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  slug varchar(240) NOT NULL,
  name varchar(240) NOT NULL,
  instrument_id uuid,
  parameter_id uuid NOT NULL,
  unit_id uuid NOT NULL,
  type timeseries_type,
  CONSTRAINT instrument_unique_timeseries_name UNIQUE (instrument_id, name),
  CONSTRAINT instrument_unique_timeseries_slug UNIQUE (instrument_id, slug),
  CONSTRAINT timeseries_pkey PRIMARY KEY (id),
  CONSTRAINT timeseries_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id),
  CONSTRAINT timeseries_parameter_id_fkey FOREIGN KEY (parameter_id) REFERENCES parameter(id),
  CONSTRAINT timeseries_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES unit(id)
);


CREATE TABLE instrument_status (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  instrument_id uuid NOT NULL,
  status_id uuid NOT NULL,
  "time" timestamptz DEFAULT now() NOT NULL,
  CONSTRAINT instrument_status_pkey PRIMARY KEY (id),
  CONSTRAINT instrument_unique_status_in_time UNIQUE (instrument_id, "time"),
  CONSTRAINT instrument_status_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id),
  CONSTRAINT instrument_status_status_id_fkey FOREIGN KEY (status_id) REFERENCES status(id)
);


CREATE TABLE aware_platform (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  aware_id uuid NOT NULL,
  instrument_id uuid,
  CONSTRAINT aware_platform_id_key UNIQUE (id),
  CONSTRAINT aware_platform_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id)
);


CREATE TABLE aware_parameter (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  key varchar NOT NULL,
  parameter_id uuid NOT NULL,
  unit_id uuid NOT NULL,
  timeseries_slug varchar NOT NULL,
  timeseries_name varchar NOT NULL,
  CONSTRAINT aware_parameter_id_key UNIQUE (id),
  CONSTRAINT aware_parameter_parameter_id_fkey FOREIGN KEY (parameter_id) REFERENCES parameter(id),
  CONSTRAINT aware_parameter_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES unit(id)
);


CREATE TABLE instrument_telemetry (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  instrument_id uuid NOT NULL,
  telemetry_type_id uuid NOT NULL,
  telemetry_id uuid NOT NULL,
  CONSTRAINT instrument_telemetry_pkey PRIMARY KEY (id),
  CONSTRAINT instrument_unique_telemetry_id UNIQUE (instrument_id, telemetry_id),
  CONSTRAINT instrument_telemetry_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id),
  CONSTRAINT instrument_telemetry_telemetry_type_id_fkey FOREIGN KEY (telemetry_type_id) REFERENCES telemetry_type(id)
);


CREATE TABLE instrument_note (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  instrument_id uuid NOT NULL,
  title varchar(240) NOT NULL,
  body varchar(65535) NOT NULL,
  "time" timestamptz DEFAULT now() NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  CONSTRAINT instrument_note_pkey PRIMARY KEY (id),
  CONSTRAINT instrument_note_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id)
);


CREATE TABLE timeseries_cwms (
  timeseries_id uuid NOT NULL,
  cwms_timeseries_id text NOT NULL,
  cwms_office_id text NOT NULL,
  cwms_extent_earliest_time timestamptz NOT NULL,
  cwms_extent_latest_time timestamptz,
  CONSTRAINT timeseries_cwms_check CHECK (((cwms_extent_latest_time IS NULL) OR (cwms_extent_earliest_time <= cwms_extent_latest_time))),
  CONSTRAINT timeseries_cwms_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE project (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  image varchar,
  federal_id varchar,
  deleted boolean DEFAULT false NOT NULL,
  slug varchar(240) NOT NULL,
  name varchar(240) NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  district_id uuid,
  CONSTRAINT project_name_key UNIQUE (name),
  CONSTRAINT project_pkey PRIMARY KEY (id),
  CONSTRAINT project_slug_key UNIQUE (slug),
  CONSTRAINT project_district_id_fkey FOREIGN KEY (district_id) REFERENCES district(id)
);


CREATE TABLE calculation (
  timeseries_id uuid NOT NULL,
  contents varchar,
  CONSTRAINT calculation_timeseries_id_key UNIQUE (timeseries_id),
  CONSTRAINT calculation_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE timeseries_measurement (
  "time" timestamptz NOT NULL,
  value double precision NOT NULL,
  timeseries_id uuid NOT NULL,
  CONSTRAINT timeseries_unique_time PRIMARY KEY (timeseries_id, "time"),
  CONSTRAINT timeseries_measurement_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE timeseries_notes (
  masked boolean DEFAULT false,
  validated boolean DEFAULT false,
  annotation varchar(400) DEFAULT ''::varchar,
  timeseries_id uuid NOT NULL,
  "time" timestamptz NOT NULL,
  CONSTRAINT notes_unique_time PRIMARY KEY (timeseries_id, "time"),
  CONSTRAINT timeseries_notes_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE inclinometer_measurement (
  "time" timestamptz NOT NULL,
  "values" jsonb NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  timeseries_id uuid NOT NULL,
  CONSTRAINT inclinometer_unique_time PRIMARY KEY (timeseries_id, "time"),
  CONSTRAINT inclinometer_measurement_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE instrument_constants (
  timeseries_id uuid NOT NULL,
  instrument_id uuid NOT NULL,
  CONSTRAINT instrument_unique_timeseries UNIQUE (instrument_id, timeseries_id),
  CONSTRAINT instrument_constants_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE,
  CONSTRAINT instrument_constants_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE aware_platform_parameter_enabled (
  aware_platform_id uuid NOT NULL,
  aware_parameter_id uuid NOT NULL,
  CONSTRAINT aware_platform_unique_parameter UNIQUE (aware_platform_id, aware_parameter_id),
  CONSTRAINT aware_platform_parameter_enabled_aware_parameter_id_fkey FOREIGN KEY (aware_parameter_id) REFERENCES aware_parameter(id),
  CONSTRAINT aware_platform_parameter_enabled_aware_platform_id_fkey FOREIGN KEY (aware_platform_id) REFERENCES aware_platform(id)
);


CREATE TABLE saa_opts (
  instrument_id uuid NOT NULL,
  num_segments integer NOT NULL,
  bottom_elevation_timeseries_id uuid,
  initial_time timestamptz,
  CONSTRAINT saa_opts_bottom_elevation_timeseries_id_fkey FOREIGN KEY (bottom_elevation_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT saa_opts_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE
);


CREATE TABLE saa_segment (
  instrument_id uuid NOT NULL,
  id integer NOT NULL,
  length_timeseries_id uuid,
  x_timeseries_id uuid,
  y_timeseries_id uuid,
  z_timeseries_id uuid,
  temp_timeseries_id uuid,
  CONSTRAINT saa_segment_pkey PRIMARY KEY (instrument_id, id),
  CONSTRAINT saa_segment_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE,
  CONSTRAINT saa_segment_length_timeseries_id_fkey FOREIGN KEY (length_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT saa_segment_temp_timeseries_id_fkey FOREIGN KEY (temp_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT saa_segment_x_timeseries_id_fkey FOREIGN KEY (x_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT saa_segment_y_timeseries_id_fkey FOREIGN KEY (y_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT saa_segment_z_timeseries_id_fkey FOREIGN KEY (z_timeseries_id) REFERENCES timeseries(id)
);


CREATE TABLE ipi_opts (
  instrument_id uuid NOT NULL,
  num_segments integer NOT NULL,
  bottom_elevation_timeseries_id uuid,
  initial_time timestamptz,
  CONSTRAINT ipi_opts_bottom_elevation_timeseries_id_fkey FOREIGN KEY (bottom_elevation_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT ipi_opts_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE
);


CREATE TABLE ipi_segment (
  instrument_id uuid NOT NULL,
  id integer NOT NULL,
  length_timeseries_id uuid,
  tilt_timeseries_id uuid,
  inc_dev_timeseries_id uuid,
  temp_timeseries_id uuid,
  CONSTRAINT ipi_segment_pkey PRIMARY KEY (instrument_id, id),
  CONSTRAINT ipi_segment_cum_dev_timeseries_id_fkey FOREIGN KEY (inc_dev_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT ipi_segment_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE,
  CONSTRAINT ipi_segment_length_timeseries_id_fkey FOREIGN KEY (length_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT ipi_segment_temp_timeseries_id_fkey FOREIGN KEY (temp_timeseries_id) REFERENCES timeseries(id),
  CONSTRAINT ipi_segment_tilt_timeseries_id_fkey FOREIGN KEY (tilt_timeseries_id) REFERENCES timeseries(id)
);


CREATE TABLE alert_config (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  project_id uuid NOT NULL,
  name varchar(480) NOT NULL,
  body text DEFAULT ''::text NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  alert_type_id uuid NOT NULL,
  start_date timestamptz DEFAULT now() NOT NULL,
  schedule_interval interval NOT NULL,
  n_missed_before_alert integer DEFAULT 1 NOT NULL,
  warning_interval interval DEFAULT '00:00:00'::interval NOT NULL,
  remind_interval interval DEFAULT '00:00:00'::interval NOT NULL,
  last_checked timestamptz,
  last_reminded timestamptz,
  deleted boolean DEFAULT false NOT NULL,
  mute_consecutive_alerts boolean DEFAULT false NOT NULL,
  CONSTRAINT alert_config_n_missed_before_alert_check CHECK ((n_missed_before_alert >= 1)),
  CONSTRAINT interval_not_negative CHECK (((schedule_interval >= '00:00:00'::interval) AND (warning_interval >= '00:00:00'::interval) AND ((remind_interval = '00:00:00'::interval) OR (remind_interval >= '1 day'::interval)))),
  CONSTRAINT warning_before_schedule CHECK ((warning_interval < schedule_interval)),
  CONSTRAINT alert_config_pkey PRIMARY KEY (id),
  CONSTRAINT alert_config_alert_type_id_fkey FOREIGN KEY (alert_type_id) REFERENCES alert_type(id),
  CONSTRAINT alert_config_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id)
);


CREATE TABLE project_instrument (
  project_id uuid NOT NULL,
  instrument_id uuid NOT NULL,
  CONSTRAINT project_instrument_project_id_instrument_id_key UNIQUE (project_id, instrument_id),
  CONSTRAINT project_instrument_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id),
  CONSTRAINT project_instrument_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id)
);


CREATE TABLE instrument_group (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  deleted boolean DEFAULT false NOT NULL,
  slug varchar(240) NOT NULL,
  name varchar(120) NOT NULL,
  description varchar(360),
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  project_id uuid,
  CONSTRAINT instrument_group_pkey PRIMARY KEY (id),
  CONSTRAINT instrument_group_slug_key UNIQUE (slug),
  CONSTRAINT project_unique_instrument_group_name UNIQUE (name, project_id),
  CONSTRAINT instrument_group_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id)
);


CREATE TABLE report_config (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  project_id uuid NOT NULL,
  slug text NOT NULL,
  name text NOT NULL,
  description text NOT NULL,
  creator uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  date_range text DEFAULT '1 year'::text,
  date_range_enabled boolean DEFAULT false,
  show_masked boolean DEFAULT false,
  show_masked_enabled boolean DEFAULT false,
  show_nonvalidated boolean DEFAULT false,
  show_nonvalidated_enabled boolean DEFAULT false,
  CONSTRAINT report_config_pkey PRIMARY KEY (id),
  CONSTRAINT report_config_slug_key UNIQUE (slug),
  CONSTRAINT report_config_creator_fkey FOREIGN KEY (creator) REFERENCES profile(id),
  CONSTRAINT report_config_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id),
  CONSTRAINT report_config_updater_fkey FOREIGN KEY (updater) REFERENCES profile(id)
);


CREATE TABLE plot_configuration (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  slug varchar NOT NULL,
  name varchar NOT NULL,
  project_id uuid NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  plot_type plot_type NOT NULL,
  CONSTRAINT plot_configuration_pkey PRIMARY KEY (id),
  CONSTRAINT project_unique_plot_configuration_name UNIQUE (project_id, name),
  CONSTRAINT project_unique_plot_configuration_slug UNIQUE (project_id, slug),
  CONSTRAINT plot_configuration_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);


CREATE TABLE datalogger (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  sn text NOT NULL,
  project_id uuid NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  update_date timestamptz DEFAULT now() NOT NULL,
  name text NOT NULL,
  slug text NOT NULL,
  model_id uuid NOT NULL,
  deleted boolean DEFAULT false NOT NULL,
  CONSTRAINT datalogger_pkey PRIMARY KEY (id),
  CONSTRAINT unique_datalogger_deleted UNIQUE (id, deleted),
  CONSTRAINT datalogger_model_id_fkey FOREIGN KEY (model_id) REFERENCES datalogger_model(id),
  CONSTRAINT datalogger_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id)
);


CREATE TABLE collection_group (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  project_id uuid NOT NULL,
  name varchar NOT NULL,
  slug varchar NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  CONSTRAINT collection_group_pkey PRIMARY KEY (id),
  CONSTRAINT project_unique_collection_group_name UNIQUE (project_id, name),
  CONSTRAINT project_unique_collection_group_slug UNIQUE (project_id, slug),
  CONSTRAINT collection_group_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);


CREATE TABLE profile_project_roles (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  profile_id uuid NOT NULL,
  role_id uuid NOT NULL,
  project_id uuid NOT NULL,
  granted_by uuid,
  granted_date timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT profile_project_roles_pkey PRIMARY KEY (id),
  CONSTRAINT unique_profile_project_role UNIQUE (profile_id, project_id, role_id),
  CONSTRAINT profile_project_roles_granted_by_fkey FOREIGN KEY (granted_by) REFERENCES profile(id),
  CONSTRAINT profile_project_roles_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profile(id),
  CONSTRAINT profile_project_roles_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id),
  CONSTRAINT profile_project_roles_role_id_fkey FOREIGN KEY (role_id) REFERENCES role(id)
);


CREATE TABLE plot_profile_config (
  plot_config_id uuid NOT NULL,
  instrument_id uuid NOT NULL,
  CONSTRAINT plot_profile_config_plot_config_id_key UNIQUE (plot_config_id),
  CONSTRAINT plot_profile_config_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE,
  CONSTRAINT plot_profile_config_plot_config_id_fkey FOREIGN KEY (plot_config_id) REFERENCES plot_configuration(id) ON DELETE CASCADE
);


CREATE TABLE plot_scatter_line_config (
  plot_config_id uuid NOT NULL,
  y_axis_title text,
  y2_axis_title text,
  CONSTRAINT plot_scatter_line_config_plot_config_id_key UNIQUE (plot_config_id),
  CONSTRAINT plot_scatter_line_config_plot_config_id_fkey FOREIGN KEY (plot_config_id) REFERENCES plot_configuration(id) ON DELETE CASCADE
);


CREATE TABLE datalogger_table (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  datalogger_id uuid NOT NULL,
  table_name text NOT NULL,
  CONSTRAINT datalogger_table_datalogger_id_table_name_key UNIQUE (datalogger_id, table_name),
  CONSTRAINT datalogger_table_pkey PRIMARY KEY (id),
  CONSTRAINT datalogger_table_datalogger_id_fkey FOREIGN KEY (datalogger_id) REFERENCES datalogger(id)
);


CREATE TABLE report_config_plot_config (
  report_config_id uuid NOT NULL,
  plot_config_id uuid NOT NULL,
  CONSTRAINT report_config_plot_config_report_config_id_plot_config_id_key UNIQUE (report_config_id, plot_config_id),
  CONSTRAINT report_config_plot_config_plot_config_id_fkey FOREIGN KEY (plot_config_id) REFERENCES plot_configuration(id) ON DELETE CASCADE,
  CONSTRAINT report_config_plot_config_report_config_id_fkey FOREIGN KEY (report_config_id) REFERENCES report_config(id) ON DELETE CASCADE
);


CREATE TABLE report_download_job (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  report_config_id uuid,
  creator uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  status job_status DEFAULT 'INIT'::midas.job_status NOT NULL,
  file_key text,
  file_expiry timestamptz,
  progress integer DEFAULT 0 NOT NULL,
  progress_update_date timestamptz DEFAULT now() NOT NULL,
  CONSTRAINT report_download_job_progress_check CHECK (((progress >= 0) AND (progress <= 100))),
  CONSTRAINT report_download_job_pkey PRIMARY KEY (id),
  CONSTRAINT report_download_job_creator_fkey FOREIGN KEY (creator) REFERENCES profile(id) ON DELETE CASCADE,
  CONSTRAINT report_download_job_report_config_id_fkey FOREIGN KEY (report_config_id) REFERENCES report_config(id) ON DELETE CASCADE
);


CREATE TABLE plot_configuration_timeseries_trace (
  plot_configuration_id uuid,
  timeseries_id uuid,
  trace_order integer NOT NULL,
  trace_type trace_type DEFAULT 'scattergl'::midas.trace_type NOT NULL,
  color text NOT NULL,
  line_style line_style DEFAULT 'solid'::midas.line_style NOT NULL,
  width real DEFAULT 1 NOT NULL,
  show_markers boolean DEFAULT false NOT NULL,
  y_axis y_axis DEFAULT 'y1'::midas.y_axis NOT NULL,
  CONSTRAINT plot_configuration_timeseries_plot_configuration_id_timeser_key UNIQUE (plot_configuration_id, timeseries_id),
  CONSTRAINT plot_configuration_timeseries_trace_plot_configuration_id_fkey FOREIGN KEY (plot_configuration_id) REFERENCES plot_configuration(id) ON DELETE CASCADE,
  CONSTRAINT plot_configuration_timeseries_trace_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE plot_contour_config (
  plot_config_id uuid NOT NULL,
  "time" timestamptz,
  locf_backfill interval NOT NULL,
  gradient_smoothing boolean DEFAULT false NOT NULL,
  contour_smoothing boolean DEFAULT false NOT NULL,
  show_labels boolean DEFAULT false NOT NULL,
  CONSTRAINT plot_contour_config_plot_config_id_key UNIQUE (plot_config_id),
  CONSTRAINT plot_contour_config_plot_config_id_fkey FOREIGN KEY (plot_config_id) REFERENCES plot_configuration(id) ON DELETE CASCADE
);


CREATE TABLE plot_configuration_custom_shape (
  plot_configuration_id uuid,
  enabled boolean DEFAULT false NOT NULL,
  name text NOT NULL,
  data_point real NOT NULL,
  color text NOT NULL,
  CONSTRAINT plot_configuration_custom_shape_plot_configuration_id_fkey FOREIGN KEY (plot_configuration_id) REFERENCES plot_configuration(id) ON DELETE CASCADE
);


CREATE TABLE alert_config_instrument (
  alert_config_id uuid NOT NULL,
  instrument_id uuid NOT NULL,
  CONSTRAINT alert_config_instrument_alert_config_id_fkey FOREIGN KEY (alert_config_id) REFERENCES alert_config(id) ON DELETE CASCADE,
  CONSTRAINT alert_config_instrument_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE
);


CREATE TABLE plot_configuration_settings (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  show_masked boolean DEFAULT true NOT NULL,
  show_nonvalidated boolean DEFAULT true NOT NULL,
  show_comments boolean DEFAULT true NOT NULL,
  auto_range boolean DEFAULT true NOT NULL,
  date_range varchar DEFAULT '1 year'::varchar NOT NULL,
  threshold integer DEFAULT 3000 NOT NULL,
  CONSTRAINT plot_configuration_settings_pkey PRIMARY KEY (id),
  CONSTRAINT plot_configuration_settings_id_fkey FOREIGN KEY (id) REFERENCES plot_configuration(id) ON DELETE CASCADE
);


CREATE TABLE datalogger_hash (
  datalogger_id uuid NOT NULL,
  hash text NOT NULL,
  CONSTRAINT unique_datalogger_hash UNIQUE (datalogger_id, hash),
  CONSTRAINT datalogger_hash_datalogger_id_fkey FOREIGN KEY (datalogger_id) REFERENCES datalogger(id)
);


CREATE TABLE submittal (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  alert_config_id uuid,
  submittal_status_id uuid DEFAULT '0c0d6487-3f71-4121-8575-19514c7b9f03'::uuid,
  completion_date timestamptz,
  create_date timestamptz DEFAULT now() NOT NULL,
  due_date timestamptz NOT NULL,
  marked_as_missing boolean DEFAULT false NOT NULL,
  warning_sent boolean DEFAULT false NOT NULL,
  CONSTRAINT submittal_check CHECK ((create_date < due_date)),
  CONSTRAINT submittal_alert_config_id_tstzrange_excl EXCLUDE USING gist (alert_config_id WITH =, tstzrange(create_date, due_date) WITH &&),
  CONSTRAINT submittal_pkey PRIMARY KEY (id),
  CONSTRAINT submittal_alert_config_id_fkey FOREIGN KEY (alert_config_id) REFERENCES alert_config(id) ON DELETE CASCADE,
  CONSTRAINT submittal_submittal_status_id_fkey FOREIGN KEY (submittal_status_id) REFERENCES submittal_status(id)
);


CREATE TABLE collection_group_timeseries (
  collection_group_id uuid NOT NULL,
  timeseries_id uuid NOT NULL,
  CONSTRAINT collection_group_unique_timeseries UNIQUE (collection_group_id, timeseries_id),
  CONSTRAINT collection_group_timeseries_collection_group_id_fkey FOREIGN KEY (collection_group_id) REFERENCES collection_group(id) ON DELETE CASCADE,
  CONSTRAINT collection_group_timeseries_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE instrument_group_instruments (
  instrument_id uuid NOT NULL,
  instrument_group_id uuid NOT NULL,
  CONSTRAINT instrument_group_instruments_instrument_id_instrument_group_key UNIQUE (instrument_id, instrument_group_id),
  CONSTRAINT instrument_group_instruments_instrument_group_id_fkey FOREIGN KEY (instrument_group_id) REFERENCES instrument_group(id),
  CONSTRAINT instrument_group_instruments_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id)
);


CREATE TABLE plot_bullseye_config (
  plot_config_id uuid NOT NULL,
  x_axis_timeseries_id uuid,
  y_axis_timeseries_id uuid,
  CONSTRAINT plot_bullseye_config_plot_config_id_key UNIQUE (plot_config_id),
  CONSTRAINT plot_bullseye_config_plot_config_id_fkey FOREIGN KEY (plot_config_id) REFERENCES plot_configuration(id) ON DELETE CASCADE,
  CONSTRAINT plot_bullseye_config_x_axis_timeseries_id_fkey FOREIGN KEY (x_axis_timeseries_id) REFERENCES timeseries(id) ON DELETE SET NULL,
  CONSTRAINT plot_bullseye_config_y_axis_timeseries_id_fkey FOREIGN KEY (y_axis_timeseries_id) REFERENCES timeseries(id) ON DELETE SET NULL
);


CREATE TABLE datalogger_equivalency_table (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  datalogger_id uuid NOT NULL,
  datalogger_deleted boolean DEFAULT false NOT NULL,
  field_name text NOT NULL,
  display_name text,
  instrument_id uuid,
  timeseries_id uuid,
  datalogger_table_id uuid,
  CONSTRAINT datalogger_equivalency_table_datalogger_table_id_field_name_key UNIQUE (datalogger_table_id, field_name),
  CONSTRAINT datalogger_equivalency_table_pkey PRIMARY KEY (id),
  CONSTRAINT datalogger_equivalency_table_datalogger_table_id_fkey FOREIGN KEY (datalogger_table_id) REFERENCES datalogger_table(id) ON DELETE CASCADE,
  CONSTRAINT datalogger_equivalency_table_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE,
  CONSTRAINT datalogger_equivalency_table_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE,
  CONSTRAINT unique_active_datalogger FOREIGN KEY (datalogger_id, datalogger_deleted) REFERENCES datalogger(id, deleted) ON UPDATE CASCADE
);


CREATE TABLE evaluation (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  project_id uuid NOT NULL,
  name varchar(480) NOT NULL,
  body text DEFAULT ''::text NOT NULL,
  start_date timestamptz NOT NULL,
  end_date timestamptz NOT NULL,
  creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  updater uuid,
  update_date timestamptz,
  submittal_id uuid,
  CONSTRAINT evaluation_pkey PRIMARY KEY (id),
  CONSTRAINT evaluation_project_id_fkey FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE,
  CONSTRAINT evaluation_submittal_id_fkey FOREIGN KEY (submittal_id) REFERENCES submittal(id) ON DELETE CASCADE
);


CREATE TABLE plot_contour_config_timeseries (
  plot_contour_config_id uuid NOT NULL,
  timeseries_id uuid NOT NULL,
  CONSTRAINT plot_contour_config_timeseries_plot_contour_config_id_timeserie UNIQUE (plot_contour_config_id, timeseries_id),
  CONSTRAINT plot_contour_config_timeseries_plot_contour_config_id_fkey FOREIGN KEY (plot_contour_config_id) REFERENCES plot_contour_config(plot_config_id) ON DELETE CASCADE,
  CONSTRAINT plot_contour_config_timeseries_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES timeseries(id) ON DELETE CASCADE
);


CREATE TABLE datalogger_error (
  datalogger_id uuid NOT NULL,
  error_message text,
  datalogger_table_id uuid,
  CONSTRAINT datalogger_error_datalogger_id_fkey FOREIGN KEY (datalogger_id) REFERENCES datalogger(id) ON DELETE CASCADE,
  CONSTRAINT datalogger_error_datalogger_table_id_fkey FOREIGN KEY (datalogger_table_id) REFERENCES datalogger_table(id) ON DELETE CASCADE
);


CREATE TABLE datalogger_preview (
  preview json,
  update_date timestamptz DEFAULT now() NOT NULL,
  datalogger_table_id uuid NOT NULL,
  CONSTRAINT datalogger_preview_datalogger_table_id_key UNIQUE (datalogger_table_id),
  CONSTRAINT datalogger_preview_datalogger_table_id_fkey FOREIGN KEY (datalogger_table_id) REFERENCES datalogger_table(id) ON DELETE CASCADE
);


CREATE TABLE evaluation_instrument (
  evaluation_id uuid,
  instrument_id uuid,
  CONSTRAINT evaluation_instrument_evaluation_id_fkey FOREIGN KEY (evaluation_id) REFERENCES evaluation(id),
  CONSTRAINT evaluation_instrument_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES instrument(id)
);


CREATE UNIQUE INDEX unique_idx_datalogger_sn_model ON datalogger (sn, model_id) WHERE NOT deleted;


CREATE UNIQUE INDEX unique_idx_datalogger_timeseries ON datalogger_equivalency_table (timeseries_id) WHERE NOT datalogger_deleted;


CREATE INDEX timeseries_instrument_id_idx ON timeseries (instrument_id);


CREATE INDEX timeseries_measurement_timeseries_id_idx ON timeseries_measurement (timeseries_id);


CREATE INDEX timeseries_measurement_time_idx ON timeseries_measurement (time);


CREATE UNIQUE INDEX unique_alert_config_id_submittal_date ON submittal (alert_config_id,completion_date) WHERE completion_date IS NOT NULL;


CREATE UNIQUE INDEX unique_district_office_id ON district (office_id) WHERE office_id IS NOT NULL;


CREATE INDEX project_instrument_instrument_id_idx ON project_instrument (instrument_id);
