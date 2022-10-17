--
-- NOTE:
--
-- File paths need to be edited. Search for $$PATH$$ and
-- replace it with the path to the directory containing
-- the extracted data files.
--
--
-- PostgreSQL database dump
--

-- Dumped from database version 13.7
-- Dumped by pg_dump version 13.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE midas_stable;
--
-- Name: midas_stable; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE midas_stable WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE midas_stable OWNER TO postgres;

\connect midas_stable

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: midas_stable; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE midas_stable SET search_path TO '$user', 'public', 'topology';


\connect midas_stable

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: midas; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA midas;


ALTER SCHEMA midas OWNER TO postgres;

--
-- Name: aware_create_timeseries(); Type: FUNCTION; Schema: midas; Owner: postgres
--

CREATE FUNCTION midas.aware_create_timeseries() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN

INSERT INTO timeseries(instrument_id, parameter_id, unit_id, slug, name) (
	SELECT a.instrument_id AS instrument_id,
		   ap.parameter_id AS parameter_id,
		   ap.unit_id AS unit_id,
		   ap.timeseries_slug AS slug,
		   ap.timeseries_name AS name
	FROM aware_parameter ap
	CROSS JOIN aware_platform a
	WHERE ap.id = NEW.aware_parameter_id AND a.id = NEW.aware_platform_id
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;


ALTER FUNCTION midas.aware_create_timeseries() OWNER TO postgres;

--
-- Name: aware_enable_params(); Type: FUNCTION; Schema: midas; Owner: postgres
--

CREATE FUNCTION midas.aware_enable_params() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN

INSERT INTO aware_platform_parameter_enabled (aware_platform_id, aware_parameter_id) (
	SELECT a.id AS aware_platform_id,
		   b.id AS aware_parameter_id
	FROM aware_platform a
	CROSS JOIN aware_parameter b	
	where a.id = NEW.id
	ORDER BY aware_platform_id
	
)
ON CONFLICT DO NOTHING;
RETURN NEW;
END;
$$;


ALTER FUNCTION midas.aware_enable_params() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: alert; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.alert (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    alert_config_id uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE midas.alert OWNER TO postgres;

--
-- Name: alert_config; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.alert_config (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    instrument_id uuid NOT NULL,
    name character varying(480),
    body text,
    formula text,
    schedule text,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone
);


ALTER TABLE midas.alert_config OWNER TO postgres;

--
-- Name: alert_email_subscription; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.alert_email_subscription (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    alert_config_id uuid NOT NULL,
    email_id uuid NOT NULL,
    mute_notify boolean DEFAULT false NOT NULL
);


ALTER TABLE midas.alert_email_subscription OWNER TO postgres;

--
-- Name: alert_profile_subscription; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.alert_profile_subscription (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    alert_config_id uuid NOT NULL,
    profile_id uuid NOT NULL,
    mute_ui boolean DEFAULT false NOT NULL,
    mute_notify boolean DEFAULT false NOT NULL
);


ALTER TABLE midas.alert_profile_subscription OWNER TO postgres;

--
-- Name: alert_read; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.alert_read (
    alert_id uuid NOT NULL,
    profile_id uuid NOT NULL
);


ALTER TABLE midas.alert_read OWNER TO postgres;

--
-- Name: aware_parameter; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.aware_parameter (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    key character varying NOT NULL,
    parameter_id uuid NOT NULL,
    unit_id uuid NOT NULL,
    timeseries_slug character varying NOT NULL,
    timeseries_name character varying NOT NULL
);


ALTER TABLE midas.aware_parameter OWNER TO postgres;

--
-- Name: aware_platform; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.aware_platform (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    aware_id uuid NOT NULL,
    instrument_id uuid
);


ALTER TABLE midas.aware_platform OWNER TO postgres;

--
-- Name: aware_platform_parameter_enabled; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.aware_platform_parameter_enabled (
    aware_platform_id uuid NOT NULL,
    aware_parameter_id uuid NOT NULL
);


ALTER TABLE midas.aware_platform_parameter_enabled OWNER TO postgres;

--
-- Name: collection_group; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.collection_group (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    project_id uuid NOT NULL,
    name character varying NOT NULL,
    slug character varying NOT NULL,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone
);


ALTER TABLE midas.collection_group OWNER TO postgres;

--
-- Name: collection_group_timeseries; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.collection_group_timeseries (
    collection_group_id uuid NOT NULL,
    timeseries_id uuid NOT NULL
);


ALTER TABLE midas.collection_group_timeseries OWNER TO postgres;

--
-- Name: config; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.config (
    static_host character varying DEFAULT 'http://minio:9000'::character varying NOT NULL,
    static_prefix character varying DEFAULT '/instrumentation'::character varying NOT NULL
);


ALTER TABLE midas.config OWNER TO postgres;

--
-- Name: email; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.email (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    email character varying(240) NOT NULL
);


ALTER TABLE midas.email OWNER TO postgres;

--
-- Name: heartbeat; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.heartbeat (
    "time" timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE midas.heartbeat OWNER TO postgres;

--
-- Name: inclinometer_measurement; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.inclinometer_measurement (
    "time" timestamp with time zone NOT NULL,
    "values" jsonb NOT NULL,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    timeseries_id uuid NOT NULL
);


ALTER TABLE midas.inclinometer_measurement OWNER TO postgres;

--
-- Name: instrument; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    slug character varying NOT NULL,
    name character varying(360) NOT NULL,
    formula_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    formula character varying,
    formula_parameter_id uuid,
    formula_unit_id uuid,
    geometry public.geometry,
    station integer,
    station_offset integer,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone,
    type_id uuid NOT NULL,
    project_id uuid,
    nid_id character varying,
    usgs_id character varying
);


ALTER TABLE midas.instrument OWNER TO postgres;

--
-- Name: instrument_constants; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_constants (
    timeseries_id uuid NOT NULL,
    instrument_id uuid NOT NULL
);


ALTER TABLE midas.instrument_constants OWNER TO postgres;

--
-- Name: instrument_group; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_group (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    slug character varying(240) NOT NULL,
    name character varying(120) NOT NULL,
    description character varying(360),
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone,
    project_id uuid
);


ALTER TABLE midas.instrument_group OWNER TO postgres;

--
-- Name: instrument_group_instruments; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_group_instruments (
    instrument_id uuid NOT NULL,
    instrument_group_id uuid NOT NULL
);


ALTER TABLE midas.instrument_group_instruments OWNER TO postgres;

--
-- Name: instrument_note; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_note (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    instrument_id uuid NOT NULL,
    title character varying(240) NOT NULL,
    body character varying(65535) NOT NULL,
    "time" timestamp with time zone DEFAULT now() NOT NULL,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone
);


ALTER TABLE midas.instrument_note OWNER TO postgres;

--
-- Name: instrument_status; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_status (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    instrument_id uuid NOT NULL,
    status_id uuid NOT NULL,
    "time" timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE midas.instrument_status OWNER TO postgres;

--
-- Name: instrument_telemetry; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_telemetry (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    instrument_id uuid NOT NULL,
    telemetry_type_id uuid NOT NULL,
    telemetry_id uuid NOT NULL
);


ALTER TABLE midas.instrument_telemetry OWNER TO postgres;

--
-- Name: instrument_type; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.instrument_type (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(120) NOT NULL
);


ALTER TABLE midas.instrument_type OWNER TO postgres;

--
-- Name: measure; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.measure (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(240) NOT NULL
);


ALTER TABLE midas.measure OWNER TO postgres;

--
-- Name: parameter; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.parameter (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(120) NOT NULL
);


ALTER TABLE midas.parameter OWNER TO postgres;

--
-- Name: plot_configuration; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.plot_configuration (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    slug character varying NOT NULL,
    name character varying NOT NULL,
    project_id uuid NOT NULL,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone
);


ALTER TABLE midas.plot_configuration OWNER TO postgres;

--
-- Name: plot_configuration_timeseries; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.plot_configuration_timeseries (
    plot_configuration_id uuid NOT NULL,
    timeseries_id uuid NOT NULL
);


ALTER TABLE midas.plot_configuration_timeseries OWNER TO postgres;

--
-- Name: profile; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.profile (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    edipi bigint NOT NULL,
    username character varying(240) NOT NULL,
    email character varying(240) NOT NULL,
    is_admin boolean DEFAULT false NOT NULL
);


ALTER TABLE midas.profile OWNER TO postgres;

--
-- Name: profile_project_roles; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.profile_project_roles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    role_id uuid NOT NULL,
    project_id uuid NOT NULL,
    granted_by uuid,
    granted_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE midas.profile_project_roles OWNER TO postgres;

--
-- Name: profile_token; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.profile_token (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    token_id character varying NOT NULL,
    profile_id uuid NOT NULL,
    issued timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    hash character varying(240) NOT NULL
);


ALTER TABLE midas.profile_token OWNER TO postgres;

--
-- Name: project; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.project (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image character varying,
    office_id uuid,
    federal_id character varying,
    deleted boolean DEFAULT false NOT NULL,
    slug character varying(240) NOT NULL,
    name character varying(240) NOT NULL,
    creator uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid NOT NULL,
    create_date timestamp with time zone DEFAULT now() NOT NULL,
    updater uuid,
    update_date timestamp with time zone
);


ALTER TABLE midas.project OWNER TO postgres;

--
-- Name: project_timeseries; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.project_timeseries (
    timeseries_id uuid NOT NULL,
    project_id uuid NOT NULL
);


ALTER TABLE midas.project_timeseries OWNER TO postgres;

--
-- Name: role; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.role (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE midas.role OWNER TO postgres;

--
-- Name: status; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.status (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(20) NOT NULL,
    description character varying(480)
);


ALTER TABLE midas.status OWNER TO postgres;

--
-- Name: telemetry_goes; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.telemetry_goes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    nesdis_id character varying NOT NULL
);


ALTER TABLE midas.telemetry_goes OWNER TO postgres;

--
-- Name: telemetry_iridium; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.telemetry_iridium (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    imei character varying(15) NOT NULL
);


ALTER TABLE midas.telemetry_iridium OWNER TO postgres;

--
-- Name: telemetry_type; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.telemetry_type (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    slug character varying NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE midas.telemetry_type OWNER TO postgres;

--
-- Name: timeseries; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.timeseries (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    slug character varying(240) NOT NULL,
    name character varying(240) NOT NULL,
    instrument_id uuid,
    parameter_id uuid NOT NULL,
    unit_id uuid NOT NULL
);


ALTER TABLE midas.timeseries OWNER TO postgres;

--
-- Name: timeseries_measurement; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.timeseries_measurement (
    "time" timestamp with time zone NOT NULL,
    value double precision NOT NULL,
    timeseries_id uuid NOT NULL
);


ALTER TABLE midas.timeseries_measurement OWNER TO postgres;

--
-- Name: timeseries_notes; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.timeseries_notes (
    masked boolean DEFAULT false NOT NULL,
    validated boolean DEFAULT false NOT NULL,
    annotation character varying(400) DEFAULT ''::character varying NOT NULL,
    timeseries_id uuid NOT NULL,
    "time" timestamp with time zone NOT NULL
);


ALTER TABLE midas.timeseries_notes OWNER TO postgres;

--
-- Name: unit; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.unit (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(120) NOT NULL,
    abbreviation character varying(120) NOT NULL,
    unit_family_id uuid,
    measure_id uuid
);


ALTER TABLE midas.unit OWNER TO postgres;

--
-- Name: unit_family; Type: TABLE; Schema: midas; Owner: postgres
--

CREATE TABLE midas.unit_family (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(120) NOT NULL
);


ALTER TABLE midas.unit_family OWNER TO postgres;

--
-- Name: v_alert; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_alert AS
 SELECT a.id,
    a.alert_config_id,
    a.create_date,
    p.id AS project_id,
    p.name AS project_name,
    i.id AS instrument_id,
    i.name AS instrument_name,
    ac.name,
    ac.body
   FROM (((midas.alert a
     JOIN midas.alert_config ac ON ((a.alert_config_id = ac.id)))
     JOIN midas.instrument i ON ((ac.instrument_id = i.id)))
     JOIN midas.project p ON ((i.project_id = p.id)));


ALTER TABLE midas.v_alert OWNER TO postgres;

--
-- Name: v_aware_platform_parameter_enabled; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_aware_platform_parameter_enabled AS
 SELECT i.project_id,
    i.id AS instrument_id,
    a.aware_id,
    b.key AS aware_parameter_key,
    t.id AS timeseries_id
   FROM ((((midas.aware_platform_parameter_enabled e
     JOIN midas.aware_platform a ON ((a.id = e.aware_platform_id)))
     JOIN midas.instrument i ON ((i.id = a.instrument_id)))
     JOIN midas.aware_parameter b ON ((b.id = e.aware_parameter_id)))
     LEFT JOIN midas.timeseries t ON (((t.instrument_id = i.id) AND (t.parameter_id = b.parameter_id) AND (t.unit_id = b.unit_id))))
  ORDER BY i.project_id, a.aware_id;


ALTER TABLE midas.v_aware_platform_parameter_enabled OWNER TO postgres;

--
-- Name: v_email_autocomplete; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_email_autocomplete AS
 SELECT email.id,
    'email'::text AS user_type,
    NULL::character varying AS username,
    email.email,
    email.email AS username_email
   FROM midas.email
UNION
 SELECT profile.id,
    'profile'::text AS user_type,
    profile.username,
    profile.email,
    ((profile.username)::text || (profile.email)::text) AS username_email
   FROM midas.profile;


ALTER TABLE midas.v_email_autocomplete OWNER TO postgres;

--
-- Name: v_instrument_telemetry; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_instrument_telemetry AS
 SELECT a.id,
    a.instrument_id,
    b.id AS telemetry_type_id,
    b.slug AS telemetry_type_slug,
    b.name AS telemetry_type_name
   FROM (((midas.instrument_telemetry a
     JOIN midas.telemetry_type b ON ((b.id = a.telemetry_type_id)))
     LEFT JOIN midas.telemetry_goes tg ON ((a.telemetry_id = tg.id)))
     LEFT JOIN midas.telemetry_iridium ti ON ((a.telemetry_id = ti.id)));


ALTER TABLE midas.v_instrument_telemetry OWNER TO postgres;

--
-- Name: v_instrument; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_instrument AS
 SELECT i.id,
    i.deleted,
    s.status_id,
    s.status,
    s.status_time,
    i.slug,
    i.name,
    i.type_id,
    i.formula,
    t.name AS type,
    public.st_asbinary(i.geometry) AS geometry,
    i.station,
    i.station_offset,
    i.creator,
    i.create_date,
    i.updater,
    i.update_date,
    i.project_id,
    i.nid_id,
    i.usgs_id,
    tel.telemetry,
    COALESCE(c.constants, '{}'::uuid[]) AS constants,
    COALESCE(g.groups, '{}'::uuid[]) AS groups,
    COALESCE(a.alert_configs, '{}'::uuid[]) AS alert_configs
   FROM ((((((midas.instrument i
     JOIN midas.instrument_type t ON ((t.id = i.type_id)))
     JOIN ( SELECT DISTINCT ON (a_1.instrument_id) a_1.instrument_id,
            a_1."time" AS status_time,
            a_1.status_id,
            d.name AS status
           FROM (midas.instrument_status a_1
             JOIN midas.status d ON ((d.id = a_1.status_id)))
          WHERE (a_1."time" <= now())
          ORDER BY a_1.instrument_id, a_1."time" DESC) s ON ((s.instrument_id = i.id)))
     LEFT JOIN ( SELECT array_agg(instrument_constants.timeseries_id) AS constants,
            instrument_constants.instrument_id
           FROM midas.instrument_constants
          GROUP BY instrument_constants.instrument_id) c ON ((c.instrument_id = i.id)))
     LEFT JOIN ( SELECT array_agg(instrument_group_instruments.instrument_group_id) AS groups,
            instrument_group_instruments.instrument_id
           FROM midas.instrument_group_instruments
          GROUP BY instrument_group_instruments.instrument_id) g ON ((g.instrument_id = i.id)))
     LEFT JOIN ( SELECT array_agg(alert_config.id) AS alert_configs,
            alert_config.instrument_id
           FROM midas.alert_config
          GROUP BY alert_config.instrument_id) a ON ((a.instrument_id = i.id)))
     LEFT JOIN ( SELECT v.instrument_id,
            json_agg(json_build_object('id', v.id, 'slug', v.telemetry_type_slug, 'name', v.telemetry_type_name)) AS telemetry
           FROM midas.v_instrument_telemetry v
          GROUP BY v.instrument_id) tel ON ((tel.instrument_id = i.id)));


ALTER TABLE midas.v_instrument OWNER TO postgres;

--
-- Name: v_instrument_group; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_instrument_group AS
SELECT
    NULL::uuid AS id,
    NULL::character varying(240) AS slug,
    NULL::character varying(120) AS name,
    NULL::character varying(360) AS description,
    NULL::uuid AS creator,
    NULL::timestamp with time zone AS create_date,
    NULL::uuid AS updater,
    NULL::timestamp with time zone AS update_date,
    NULL::uuid AS project_id,
    NULL::boolean AS deleted,
    NULL::bigint AS instrument_count,
    NULL::bigint AS timeseries_count;


ALTER TABLE midas.v_instrument_group OWNER TO postgres;

--
-- Name: v_plot_configuration; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_plot_configuration AS
 SELECT pc.id,
    pc.slug,
    pc.name,
    pc.project_id,
    t.timeseries_id,
    pc.creator,
    pc.create_date,
    pc.updater,
    pc.update_date
   FROM (midas.plot_configuration pc
     LEFT JOIN ( SELECT plot_configuration_timeseries.plot_configuration_id,
            array_agg(plot_configuration_timeseries.timeseries_id) AS timeseries_id
           FROM midas.plot_configuration_timeseries
          GROUP BY plot_configuration_timeseries.plot_configuration_id) t ON ((pc.id = t.plot_configuration_id)));


ALTER TABLE midas.v_plot_configuration OWNER TO postgres;

--
-- Name: v_profile; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_profile AS
 WITH roles_by_profile AS (
         SELECT a.profile_id,
            array_agg(upper((((b.slug)::text || '.'::text) || (c.name)::text))) AS roles
           FROM ((midas.profile_project_roles a
             LEFT JOIN midas.project b ON ((a.project_id = b.id)))
             LEFT JOIN midas.role c ON ((a.role_id = c.id)))
          GROUP BY a.profile_id
        )
 SELECT p.id,
    p.edipi,
    p.username,
    p.email,
    p.is_admin,
    COALESCE(r.roles, '{}'::text[]) AS roles
   FROM (midas.profile p
     LEFT JOIN roles_by_profile r ON ((r.profile_id = p.id)));


ALTER TABLE midas.v_profile OWNER TO postgres;

--
-- Name: v_profile_project_roles; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_profile_project_roles AS
 SELECT a.id,
    a.profile_id,
    b.edipi,
    b.username,
    b.email,
    b.is_admin,
    c.id AS project_id,
    r.id AS role_id,
    r.name AS role,
    upper((((c.slug)::text || '.'::text) || (r.name)::text)) AS rolename
   FROM (((midas.profile_project_roles a
     JOIN midas.profile b ON ((b.id = a.profile_id)))
     JOIN midas.project c ON ((c.id = a.project_id)))
     JOIN midas.role r ON ((r.id = a.role_id)))
  ORDER BY b.username, r.name;


ALTER TABLE midas.v_profile_project_roles OWNER TO postgres;

--
-- Name: v_project; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_project AS
 SELECT p.id,
    p.federal_id,
        CASE
            WHEN (p.image IS NOT NULL) THEN (((((cfg.static_host)::text || '/projects/'::text) || (p.slug)::text) || '/images/'::text) || (p.image)::text)
            ELSE NULL::text
        END AS image,
    p.office_id,
    p.deleted,
    p.slug,
    p.name,
    p.creator,
    p.create_date,
    p.updater,
    p.update_date,
    COALESCE(t.timeseries, '{}'::uuid[]) AS timeseries,
    COALESCE(i.count, (0)::bigint) AS instrument_count,
    COALESCE(g.count, (0)::bigint) AS instrument_group_count
   FROM ((((midas.project p
     LEFT JOIN ( SELECT instrument.project_id,
            count(instrument.*) AS count
           FROM midas.instrument
          WHERE (NOT instrument.deleted)
          GROUP BY instrument.project_id) i ON ((i.project_id = p.id)))
     LEFT JOIN ( SELECT instrument_group.project_id,
            count(instrument_group.*) AS count
           FROM midas.instrument_group
          WHERE (NOT instrument_group.deleted)
          GROUP BY instrument_group.project_id) g ON ((g.project_id = p.id)))
     LEFT JOIN ( SELECT array_agg(project_timeseries.timeseries_id) AS timeseries,
            project_timeseries.project_id
           FROM midas.project_timeseries
          GROUP BY project_timeseries.project_id) t ON ((t.project_id = p.id)))
     CROSS JOIN midas.config cfg);


ALTER TABLE midas.v_project OWNER TO postgres;

--
-- Name: v_timeseries; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_timeseries AS
 WITH ts_stored_and_computed AS (
         SELECT timeseries.id,
            timeseries.slug,
            timeseries.name,
            timeseries.instrument_id,
            timeseries.parameter_id,
            timeseries.unit_id,
            false AS is_computed
           FROM midas.timeseries
        UNION
         SELECT instrument.formula_id AS id,
            'formula'::character varying AS slug,
            'Formula'::character varying AS name,
            instrument.id AS instrument_id,
            instrument.formula_parameter_id AS parameter_id,
            instrument.formula_unit_id AS unit_id,
            true AS is_computed
           FROM midas.instrument
          WHERE ((NOT instrument.deleted) AND (instrument.formula IS NOT NULL))
        )
 SELECT t.id,
    t.slug,
    t.name,
    t.is_computed,
    (((i.slug)::text || '.'::text) || (t.slug)::text) AS variable,
    j.id AS project_id,
    j.slug AS project_slug,
    j.name AS project,
    i.id AS instrument_id,
    i.slug AS instrument_slug,
    i.name AS instrument,
    p.id AS parameter_id,
    p.name AS parameter,
    u.id AS unit_id,
    u.name AS unit
   FROM ((((ts_stored_and_computed t
     JOIN midas.instrument i ON ((i.id = t.instrument_id)))
     JOIN midas.project j ON ((j.id = i.project_id)))
     JOIN midas.parameter p ON ((p.id = t.parameter_id)))
     JOIN midas.unit u ON ((u.id = t.unit_id)));


ALTER TABLE midas.v_timeseries OWNER TO postgres;

--
-- Name: v_timeseries_dependency; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_timeseries_dependency AS
 WITH variable_tsid_map AS (
         SELECT a.id AS timeseries_id,
            (((b.slug)::text || '.'::text) || (a.slug)::text) AS variable
           FROM (midas.timeseries a
             LEFT JOIN midas.instrument b ON ((b.id = a.instrument_id)))
        )
 SELECT i.instrument_id,
    i.formula_id AS timeseries_id,
    i.parsed_variable,
    m.timeseries_id AS dependency_timeseries_id
   FROM (( SELECT instrument.id AS instrument_id,
            instrument.formula_id,
            (regexp_matches((instrument.formula)::text, '\[(.*?)\]'::text, 'g'::text))[1] AS parsed_variable
           FROM midas.instrument) i
     LEFT JOIN variable_tsid_map m ON ((m.variable = i.parsed_variable)));


ALTER TABLE midas.v_timeseries_dependency OWNER TO postgres;

--
-- Name: v_timeseries_latest; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_timeseries_latest AS
 SELECT t.id,
    t.slug,
    t.name,
    t.is_computed,
    t.variable,
    t.project_id,
    t.project_slug,
    t.project,
    t.instrument_id,
    t.instrument_slug,
    t.instrument,
    t.parameter_id,
    t.parameter,
    t.unit_id,
    t.unit,
    m."time" AS latest_time,
    m.value AS latest_value
   FROM (midas.v_timeseries t
     LEFT JOIN ( SELECT DISTINCT ON (timeseries_measurement.timeseries_id) timeseries_measurement.timeseries_id,
            timeseries_measurement."time",
            timeseries_measurement.value
           FROM midas.timeseries_measurement
          ORDER BY timeseries_measurement.timeseries_id, timeseries_measurement."time" DESC) m ON ((t.id = m.timeseries_id)));


ALTER TABLE midas.v_timeseries_latest OWNER TO postgres;

--
-- Name: v_timeseries_project_map; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_timeseries_project_map AS
 SELECT t.id AS timeseries_id,
    p.id AS project_id
   FROM ((midas.timeseries t
     LEFT JOIN midas.instrument n ON ((t.instrument_id = n.id)))
     LEFT JOIN midas.project p ON ((p.id = n.project_id)));


ALTER TABLE midas.v_timeseries_project_map OWNER TO postgres;

--
-- Name: v_unit; Type: VIEW; Schema: midas; Owner: postgres
--

CREATE VIEW midas.v_unit AS
 SELECT u.id,
    u.name,
    u.abbreviation,
    u.unit_family_id,
    f.name AS unit_family,
    u.measure_id,
    m.name AS measure
   FROM ((midas.unit u
     JOIN midas.unit_family f ON ((f.id = u.unit_family_id)))
     JOIN midas.measure m ON ((m.id = u.measure_id)));


ALTER TABLE midas.v_unit OWNER TO postgres;

--
-- Name: alert_config alert_config_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_config
    ADD CONSTRAINT alert_config_pkey PRIMARY KEY (id);


--
-- Name: alert_email_subscription alert_email_subscription_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_email_subscription
    ADD CONSTRAINT alert_email_subscription_pkey PRIMARY KEY (id);


--
-- Name: alert alert_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert
    ADD CONSTRAINT alert_pkey PRIMARY KEY (id);


--
-- Name: alert_profile_subscription alert_profile_subscription_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_profile_subscription
    ADD CONSTRAINT alert_profile_subscription_pkey PRIMARY KEY (id);


--
-- Name: aware_parameter aware_parameter_id_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_parameter
    ADD CONSTRAINT aware_parameter_id_key UNIQUE (id);


--
-- Name: aware_platform aware_platform_id_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_platform
    ADD CONSTRAINT aware_platform_id_key UNIQUE (id);


--
-- Name: aware_platform_parameter_enabled aware_platform_unique_parameter; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_platform_parameter_enabled
    ADD CONSTRAINT aware_platform_unique_parameter UNIQUE (aware_platform_id, aware_parameter_id);


--
-- Name: collection_group collection_group_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group
    ADD CONSTRAINT collection_group_pkey PRIMARY KEY (id);


--
-- Name: collection_group_timeseries collection_group_unique_timeseries; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group_timeseries
    ADD CONSTRAINT collection_group_unique_timeseries UNIQUE (collection_group_id, timeseries_id);


--
-- Name: email email_email_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.email
    ADD CONSTRAINT email_email_key UNIQUE (email);


--
-- Name: email email_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.email
    ADD CONSTRAINT email_pkey PRIMARY KEY (id);


--
-- Name: alert_email_subscription email_unique_alert_config; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_email_subscription
    ADD CONSTRAINT email_unique_alert_config UNIQUE (email_id, alert_config_id);


--
-- Name: inclinometer_measurement inclinometer_unique_time; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.inclinometer_measurement
    ADD CONSTRAINT inclinometer_unique_time PRIMARY KEY (timeseries_id, "time");


--
-- Name: instrument_group_instruments instrument_group_instruments_instrument_id_instrument_group_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group_instruments
    ADD CONSTRAINT instrument_group_instruments_instrument_id_instrument_group_key UNIQUE (instrument_id, instrument_group_id);


--
-- Name: instrument_group instrument_group_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group
    ADD CONSTRAINT instrument_group_pkey PRIMARY KEY (id);


--
-- Name: instrument_group instrument_group_slug_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group
    ADD CONSTRAINT instrument_group_slug_key UNIQUE (slug);


--
-- Name: instrument_note instrument_note_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_note
    ADD CONSTRAINT instrument_note_pkey PRIMARY KEY (id);


--
-- Name: instrument instrument_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_pkey PRIMARY KEY (id);


--
-- Name: instrument instrument_slug_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_slug_key UNIQUE (slug);


--
-- Name: instrument_status instrument_status_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_status
    ADD CONSTRAINT instrument_status_pkey PRIMARY KEY (id);


--
-- Name: instrument_telemetry instrument_telemetry_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_telemetry
    ADD CONSTRAINT instrument_telemetry_pkey PRIMARY KEY (id);


--
-- Name: instrument_type instrument_type_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_type
    ADD CONSTRAINT instrument_type_name_key UNIQUE (name);


--
-- Name: instrument_type instrument_type_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_type
    ADD CONSTRAINT instrument_type_pkey PRIMARY KEY (id);


--
-- Name: alert_config instrument_unique_alert_config_name; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_config
    ADD CONSTRAINT instrument_unique_alert_config_name UNIQUE (name, instrument_id);


--
-- Name: instrument_status instrument_unique_status_in_time; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_status
    ADD CONSTRAINT instrument_unique_status_in_time UNIQUE (instrument_id, "time");


--
-- Name: instrument_telemetry instrument_unique_telemetry_id; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_telemetry
    ADD CONSTRAINT instrument_unique_telemetry_id UNIQUE (instrument_id, telemetry_id);


--
-- Name: instrument_constants instrument_unique_timeseries; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_constants
    ADD CONSTRAINT instrument_unique_timeseries UNIQUE (instrument_id, timeseries_id);


--
-- Name: timeseries instrument_unique_timeseries_name; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT instrument_unique_timeseries_name UNIQUE (instrument_id, name);


--
-- Name: timeseries instrument_unique_timeseries_slug; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT instrument_unique_timeseries_slug UNIQUE (instrument_id, slug);


--
-- Name: measure measure_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.measure
    ADD CONSTRAINT measure_name_key UNIQUE (name);


--
-- Name: measure measure_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.measure
    ADD CONSTRAINT measure_pkey PRIMARY KEY (id);


--
-- Name: timeseries_notes notes_unique_time; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries_notes
    ADD CONSTRAINT notes_unique_time PRIMARY KEY (timeseries_id, "time");


--
-- Name: parameter parameter_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.parameter
    ADD CONSTRAINT parameter_name_key UNIQUE (name);


--
-- Name: parameter parameter_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.parameter
    ADD CONSTRAINT parameter_pkey PRIMARY KEY (id);


--
-- Name: plot_configuration plot_configuration_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration
    ADD CONSTRAINT plot_configuration_pkey PRIMARY KEY (id);


--
-- Name: plot_configuration_timeseries plot_configuration_unique_timeseries; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration_timeseries
    ADD CONSTRAINT plot_configuration_unique_timeseries UNIQUE (plot_configuration_id, timeseries_id);


--
-- Name: profile profile_edipi_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile
    ADD CONSTRAINT profile_edipi_key UNIQUE (edipi);


--
-- Name: profile profile_email_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile
    ADD CONSTRAINT profile_email_key UNIQUE (email);


--
-- Name: profile profile_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (id);


--
-- Name: profile_project_roles profile_project_roles_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT profile_project_roles_pkey PRIMARY KEY (id);


--
-- Name: profile_token profile_token_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_token
    ADD CONSTRAINT profile_token_pkey PRIMARY KEY (id);


--
-- Name: alert_profile_subscription profile_unique_alert_config; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_profile_subscription
    ADD CONSTRAINT profile_unique_alert_config UNIQUE (profile_id, alert_config_id);


--
-- Name: alert_read profile_unique_alert_read; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_read
    ADD CONSTRAINT profile_unique_alert_read UNIQUE (alert_id, profile_id);


--
-- Name: profile profile_username_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile
    ADD CONSTRAINT profile_username_key UNIQUE (username);


--
-- Name: project project_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project
    ADD CONSTRAINT project_name_key UNIQUE (name);


--
-- Name: project project_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);


--
-- Name: project project_slug_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project
    ADD CONSTRAINT project_slug_key UNIQUE (slug);


--
-- Name: collection_group project_unique_collection_group_name; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group
    ADD CONSTRAINT project_unique_collection_group_name UNIQUE (project_id, name);


--
-- Name: collection_group project_unique_collection_group_slug; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group
    ADD CONSTRAINT project_unique_collection_group_slug UNIQUE (project_id, slug);


--
-- Name: instrument_group project_unique_instrument_group_name; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group
    ADD CONSTRAINT project_unique_instrument_group_name UNIQUE (name, project_id);


--
-- Name: plot_configuration project_unique_plot_configuration_name; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration
    ADD CONSTRAINT project_unique_plot_configuration_name UNIQUE (project_id, name);


--
-- Name: plot_configuration project_unique_plot_configuration_slug; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration
    ADD CONSTRAINT project_unique_plot_configuration_slug UNIQUE (project_id, slug);


--
-- Name: project_timeseries project_unique_timeseries; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project_timeseries
    ADD CONSTRAINT project_unique_timeseries UNIQUE (project_id, timeseries_id);


--
-- Name: role role_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- Name: status status_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.status
    ADD CONSTRAINT status_name_key UNIQUE (name);


--
-- Name: status status_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.status
    ADD CONSTRAINT status_pkey PRIMARY KEY (id);


--
-- Name: telemetry_goes telemetry_goes_nesdis_id_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_goes
    ADD CONSTRAINT telemetry_goes_nesdis_id_key UNIQUE (nesdis_id);


--
-- Name: telemetry_goes telemetry_goes_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_goes
    ADD CONSTRAINT telemetry_goes_pkey PRIMARY KEY (id);


--
-- Name: telemetry_iridium telemetry_iridium_imei_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_iridium
    ADD CONSTRAINT telemetry_iridium_imei_key UNIQUE (imei);


--
-- Name: telemetry_iridium telemetry_iridium_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_iridium
    ADD CONSTRAINT telemetry_iridium_pkey PRIMARY KEY (id);


--
-- Name: telemetry_type telemetry_type_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_type
    ADD CONSTRAINT telemetry_type_name_key UNIQUE (name);


--
-- Name: telemetry_type telemetry_type_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_type
    ADD CONSTRAINT telemetry_type_pkey PRIMARY KEY (id);


--
-- Name: telemetry_type telemetry_type_slug_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.telemetry_type
    ADD CONSTRAINT telemetry_type_slug_key UNIQUE (slug);


--
-- Name: timeseries timeseries_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT timeseries_pkey PRIMARY KEY (id);


--
-- Name: timeseries_measurement timeseries_unique_time; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries_measurement
    ADD CONSTRAINT timeseries_unique_time PRIMARY KEY (timeseries_id, "time");


--
-- Name: profile_project_roles unique_profile_project_role; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT unique_profile_project_role UNIQUE (profile_id, project_id, role_id);


--
-- Name: unit unit_abbreviation_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit
    ADD CONSTRAINT unit_abbreviation_key UNIQUE (abbreviation);


--
-- Name: unit_family unit_family_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit_family
    ADD CONSTRAINT unit_family_name_key UNIQUE (name);


--
-- Name: unit_family unit_family_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit_family
    ADD CONSTRAINT unit_family_pkey PRIMARY KEY (id);


--
-- Name: unit unit_name_key; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit
    ADD CONSTRAINT unit_name_key UNIQUE (name);


--
-- Name: unit unit_pkey; Type: CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit
    ADD CONSTRAINT unit_pkey PRIMARY KEY (id);


--
-- Name: v_instrument_group _RETURN; Type: RULE; Schema: midas; Owner: postgres
--

CREATE OR REPLACE VIEW midas.v_instrument_group AS
 WITH instrument_count AS (
         SELECT igi.instrument_group_id,
            count(igi.instrument_group_id) AS i_count
           FROM (midas.instrument_group_instruments igi
             JOIN midas.instrument i ON (((igi.instrument_id = i.id) AND (NOT i.deleted))))
          GROUP BY igi.instrument_group_id
        ), timeseries_instruments AS (
         SELECT t.id,
            t.instrument_id,
            igi.instrument_group_id
           FROM ((midas.timeseries t
             JOIN midas.instrument i ON (((i.id = t.instrument_id) AND (NOT i.deleted))))
             JOIN midas.instrument_group_instruments igi ON ((igi.instrument_id = i.id)))
        )
 SELECT ig.id,
    ig.slug,
    ig.name,
    ig.description,
    ig.creator,
    ig.create_date,
    ig.updater,
    ig.update_date,
    ig.project_id,
    ig.deleted,
    COALESCE(ic.i_count, (0)::bigint) AS instrument_count,
    COALESCE(count(ti.id), (0)::bigint) AS timeseries_count
   FROM ((midas.instrument_group ig
     LEFT JOIN instrument_count ic ON ((ic.instrument_group_id = ig.id)))
     LEFT JOIN timeseries_instruments ti ON ((ig.id = ti.instrument_group_id)))
  GROUP BY ig.id, ic.i_count
  ORDER BY ig.name;


--
-- Name: aware_platform_parameter_enabled aware_create_timeseries; Type: TRIGGER; Schema: midas; Owner: postgres
--

CREATE TRIGGER aware_create_timeseries AFTER INSERT ON midas.aware_platform_parameter_enabled FOR EACH ROW EXECUTE FUNCTION midas.aware_create_timeseries();


--
-- Name: aware_platform aware_enable_params; Type: TRIGGER; Schema: midas; Owner: postgres
--

CREATE TRIGGER aware_enable_params AFTER INSERT ON midas.aware_platform FOR EACH ROW EXECUTE FUNCTION midas.aware_enable_params();


--
-- Name: alert alert_alert_config_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert
    ADD CONSTRAINT alert_alert_config_id_fkey FOREIGN KEY (alert_config_id) REFERENCES midas.alert_config(id);


--
-- Name: alert_config alert_config_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_config
    ADD CONSTRAINT alert_config_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: alert_email_subscription alert_email_subscription_alert_config_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_email_subscription
    ADD CONSTRAINT alert_email_subscription_alert_config_id_fkey FOREIGN KEY (alert_config_id) REFERENCES midas.alert_config(id);


--
-- Name: alert_email_subscription alert_email_subscription_email_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_email_subscription
    ADD CONSTRAINT alert_email_subscription_email_id_fkey FOREIGN KEY (email_id) REFERENCES midas.email(id);


--
-- Name: alert_profile_subscription alert_profile_subscription_alert_config_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_profile_subscription
    ADD CONSTRAINT alert_profile_subscription_alert_config_id_fkey FOREIGN KEY (alert_config_id) REFERENCES midas.alert_config(id);


--
-- Name: alert_profile_subscription alert_profile_subscription_profile_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_profile_subscription
    ADD CONSTRAINT alert_profile_subscription_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES midas.profile(id);


--
-- Name: alert_read alert_read_alert_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_read
    ADD CONSTRAINT alert_read_alert_id_fkey FOREIGN KEY (alert_id) REFERENCES midas.alert(id);


--
-- Name: alert_read alert_read_profile_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.alert_read
    ADD CONSTRAINT alert_read_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES midas.profile(id);


--
-- Name: aware_parameter aware_parameter_parameter_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_parameter
    ADD CONSTRAINT aware_parameter_parameter_id_fkey FOREIGN KEY (parameter_id) REFERENCES midas.parameter(id);


--
-- Name: aware_parameter aware_parameter_unit_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_parameter
    ADD CONSTRAINT aware_parameter_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES midas.unit(id);


--
-- Name: aware_platform aware_platform_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_platform
    ADD CONSTRAINT aware_platform_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: aware_platform_parameter_enabled aware_platform_parameter_enabled_aware_parameter_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_platform_parameter_enabled
    ADD CONSTRAINT aware_platform_parameter_enabled_aware_parameter_id_fkey FOREIGN KEY (aware_parameter_id) REFERENCES midas.aware_parameter(id);


--
-- Name: aware_platform_parameter_enabled aware_platform_parameter_enabled_aware_platform_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.aware_platform_parameter_enabled
    ADD CONSTRAINT aware_platform_parameter_enabled_aware_platform_id_fkey FOREIGN KEY (aware_platform_id) REFERENCES midas.aware_platform(id);


--
-- Name: collection_group collection_group_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group
    ADD CONSTRAINT collection_group_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id) ON DELETE CASCADE;


--
-- Name: collection_group_timeseries collection_group_timeseries_collection_group_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group_timeseries
    ADD CONSTRAINT collection_group_timeseries_collection_group_id_fkey FOREIGN KEY (collection_group_id) REFERENCES midas.collection_group(id) ON DELETE CASCADE;


--
-- Name: collection_group_timeseries collection_group_timeseries_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.collection_group_timeseries
    ADD CONSTRAINT collection_group_timeseries_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: inclinometer_measurement inclinometer_measurement_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.inclinometer_measurement
    ADD CONSTRAINT inclinometer_measurement_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: instrument_constants instrument_constants_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_constants
    ADD CONSTRAINT instrument_constants_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id) ON DELETE CASCADE;


--
-- Name: instrument_constants instrument_constants_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_constants
    ADD CONSTRAINT instrument_constants_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: instrument instrument_formula_parameter_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_formula_parameter_id_fkey FOREIGN KEY (formula_parameter_id) REFERENCES midas.parameter(id);


--
-- Name: instrument instrument_formula_unit_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_formula_unit_id_fkey FOREIGN KEY (formula_unit_id) REFERENCES midas.unit(id);


--
-- Name: instrument_group_instruments instrument_group_instruments_instrument_group_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group_instruments
    ADD CONSTRAINT instrument_group_instruments_instrument_group_id_fkey FOREIGN KEY (instrument_group_id) REFERENCES midas.instrument_group(id);


--
-- Name: instrument_group_instruments instrument_group_instruments_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group_instruments
    ADD CONSTRAINT instrument_group_instruments_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: instrument_group instrument_group_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_group
    ADD CONSTRAINT instrument_group_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id);


--
-- Name: instrument_note instrument_note_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_note
    ADD CONSTRAINT instrument_note_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: instrument instrument_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id);


--
-- Name: instrument_status instrument_status_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_status
    ADD CONSTRAINT instrument_status_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: instrument_status instrument_status_status_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_status
    ADD CONSTRAINT instrument_status_status_id_fkey FOREIGN KEY (status_id) REFERENCES midas.status(id);


--
-- Name: instrument_telemetry instrument_telemetry_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_telemetry
    ADD CONSTRAINT instrument_telemetry_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: instrument_telemetry instrument_telemetry_telemetry_type_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument_telemetry
    ADD CONSTRAINT instrument_telemetry_telemetry_type_id_fkey FOREIGN KEY (telemetry_type_id) REFERENCES midas.telemetry_type(id);


--
-- Name: instrument instrument_type_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.instrument
    ADD CONSTRAINT instrument_type_id_fkey FOREIGN KEY (type_id) REFERENCES midas.instrument_type(id);


--
-- Name: plot_configuration plot_configuration_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration
    ADD CONSTRAINT plot_configuration_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id) ON DELETE CASCADE;


--
-- Name: plot_configuration_timeseries plot_configuration_timeseries_plot_configuration_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration_timeseries
    ADD CONSTRAINT plot_configuration_timeseries_plot_configuration_id_fkey FOREIGN KEY (plot_configuration_id) REFERENCES midas.plot_configuration(id) ON DELETE CASCADE;


--
-- Name: plot_configuration_timeseries plot_configuration_timeseries_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.plot_configuration_timeseries
    ADD CONSTRAINT plot_configuration_timeseries_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: profile_project_roles profile_project_roles_granted_by_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT profile_project_roles_granted_by_fkey FOREIGN KEY (granted_by) REFERENCES midas.profile(id);


--
-- Name: profile_project_roles profile_project_roles_profile_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT profile_project_roles_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES midas.profile(id);


--
-- Name: profile_project_roles profile_project_roles_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT profile_project_roles_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id);


--
-- Name: profile_project_roles profile_project_roles_role_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_project_roles
    ADD CONSTRAINT profile_project_roles_role_id_fkey FOREIGN KEY (role_id) REFERENCES midas.role(id);


--
-- Name: profile_token profile_token_profile_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.profile_token
    ADD CONSTRAINT profile_token_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES midas.profile(id);


--
-- Name: project_timeseries project_timeseries_project_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project_timeseries
    ADD CONSTRAINT project_timeseries_project_id_fkey FOREIGN KEY (project_id) REFERENCES midas.project(id) ON DELETE CASCADE;


--
-- Name: project_timeseries project_timeseries_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.project_timeseries
    ADD CONSTRAINT project_timeseries_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: timeseries timeseries_instrument_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT timeseries_instrument_id_fkey FOREIGN KEY (instrument_id) REFERENCES midas.instrument(id);


--
-- Name: timeseries_measurement timeseries_measurement_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries_measurement
    ADD CONSTRAINT timeseries_measurement_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: timeseries_notes timeseries_notes_timeseries_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries_notes
    ADD CONSTRAINT timeseries_notes_timeseries_id_fkey FOREIGN KEY (timeseries_id) REFERENCES midas.timeseries(id) ON DELETE CASCADE;


--
-- Name: timeseries timeseries_parameter_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT timeseries_parameter_id_fkey FOREIGN KEY (parameter_id) REFERENCES midas.parameter(id);


--
-- Name: timeseries timeseries_unit_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.timeseries
    ADD CONSTRAINT timeseries_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES midas.unit(id);


--
-- Name: unit unit_measure_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit
    ADD CONSTRAINT unit_measure_id_fkey FOREIGN KEY (measure_id) REFERENCES midas.measure(id);


--
-- Name: unit unit_unit_family_id_fkey; Type: FK CONSTRAINT; Schema: midas; Owner: postgres
--

ALTER TABLE ONLY midas.unit
    ADD CONSTRAINT unit_unit_family_id_fkey FOREIGN KEY (unit_family_id) REFERENCES midas.unit_family(id);


--
-- Name: SCHEMA midas; Type: ACL; Schema: -; Owner: postgres
--

GRANT USAGE ON SCHEMA midas TO instrumentation_user;


--
-- Name: TABLE alert; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.alert TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.alert TO instrumentation_writer;


--
-- Name: TABLE alert_config; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.alert_config TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.alert_config TO instrumentation_writer;


--
-- Name: TABLE alert_email_subscription; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.alert_email_subscription TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.alert_email_subscription TO instrumentation_writer;


--
-- Name: TABLE alert_profile_subscription; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.alert_profile_subscription TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.alert_profile_subscription TO instrumentation_writer;


--
-- Name: TABLE alert_read; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.alert_read TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.alert_read TO instrumentation_writer;


--
-- Name: TABLE aware_parameter; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.aware_parameter TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.aware_parameter TO instrumentation_writer;


--
-- Name: TABLE aware_platform; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.aware_platform TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.aware_platform TO instrumentation_writer;


--
-- Name: TABLE aware_platform_parameter_enabled; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.aware_platform_parameter_enabled TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.aware_platform_parameter_enabled TO instrumentation_writer;


--
-- Name: TABLE collection_group; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.collection_group TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.collection_group TO instrumentation_writer;


--
-- Name: TABLE collection_group_timeseries; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.collection_group_timeseries TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.collection_group_timeseries TO instrumentation_writer;


--
-- Name: TABLE config; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.config TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.config TO instrumentation_writer;


--
-- Name: TABLE email; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.email TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.email TO instrumentation_writer;


--
-- Name: TABLE heartbeat; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.heartbeat TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.heartbeat TO instrumentation_writer;


--
-- Name: TABLE inclinometer_measurement; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.inclinometer_measurement TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.inclinometer_measurement TO instrumentation_writer;


--
-- Name: TABLE instrument; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument TO instrumentation_writer;


--
-- Name: TABLE instrument_constants; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_constants TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_constants TO instrumentation_writer;


--
-- Name: TABLE instrument_group; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_group TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_group TO instrumentation_writer;


--
-- Name: TABLE instrument_group_instruments; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_group_instruments TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_group_instruments TO instrumentation_writer;


--
-- Name: TABLE instrument_note; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_note TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_note TO instrumentation_writer;


--
-- Name: TABLE instrument_status; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_status TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_status TO instrumentation_writer;


--
-- Name: TABLE instrument_telemetry; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_telemetry TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_telemetry TO instrumentation_writer;


--
-- Name: TABLE instrument_type; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.instrument_type TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.instrument_type TO instrumentation_writer;


--
-- Name: TABLE measure; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.measure TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.measure TO instrumentation_writer;


--
-- Name: TABLE parameter; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.parameter TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.parameter TO instrumentation_writer;


--
-- Name: TABLE plot_configuration; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.plot_configuration TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.plot_configuration TO instrumentation_writer;


--
-- Name: TABLE plot_configuration_timeseries; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.plot_configuration_timeseries TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.plot_configuration_timeseries TO instrumentation_writer;


--
-- Name: TABLE profile; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.profile TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.profile TO instrumentation_writer;


--
-- Name: TABLE profile_project_roles; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.profile_project_roles TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.profile_project_roles TO instrumentation_writer;


--
-- Name: TABLE profile_token; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.profile_token TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.profile_token TO instrumentation_writer;


--
-- Name: TABLE project; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.project TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.project TO instrumentation_writer;


--
-- Name: TABLE project_timeseries; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.project_timeseries TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.project_timeseries TO instrumentation_writer;


--
-- Name: TABLE role; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.role TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.role TO instrumentation_writer;


--
-- Name: TABLE status; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.status TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.status TO instrumentation_writer;


--
-- Name: TABLE telemetry_goes; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.telemetry_goes TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.telemetry_goes TO instrumentation_writer;


--
-- Name: TABLE telemetry_iridium; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.telemetry_iridium TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.telemetry_iridium TO instrumentation_writer;


--
-- Name: TABLE telemetry_type; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.telemetry_type TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.telemetry_type TO instrumentation_writer;


--
-- Name: TABLE timeseries; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.timeseries TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.timeseries TO instrumentation_writer;


--
-- Name: TABLE timeseries_measurement; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.timeseries_measurement TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.timeseries_measurement TO instrumentation_writer;


--
-- Name: TABLE timeseries_notes; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.timeseries_notes TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.timeseries_notes TO instrumentation_writer;


--
-- Name: TABLE unit; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.unit TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.unit TO instrumentation_writer;


--
-- Name: TABLE unit_family; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.unit_family TO instrumentation_reader;
GRANT INSERT,DELETE,UPDATE ON TABLE midas.unit_family TO instrumentation_writer;


--
-- Name: TABLE v_alert; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_alert TO instrumentation_reader;


--
-- Name: TABLE v_aware_platform_parameter_enabled; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_aware_platform_parameter_enabled TO instrumentation_reader;


--
-- Name: TABLE v_email_autocomplete; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_email_autocomplete TO instrumentation_reader;


--
-- Name: TABLE v_instrument_telemetry; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_instrument_telemetry TO instrumentation_reader;


--
-- Name: TABLE v_instrument; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_instrument TO instrumentation_reader;


--
-- Name: TABLE v_instrument_group; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_instrument_group TO instrumentation_reader;


--
-- Name: TABLE v_plot_configuration; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_plot_configuration TO instrumentation_reader;


--
-- Name: TABLE v_profile; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_profile TO instrumentation_reader;


--
-- Name: TABLE v_profile_project_roles; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_profile_project_roles TO instrumentation_reader;


--
-- Name: TABLE v_project; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_project TO instrumentation_reader;


--
-- Name: TABLE v_timeseries; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_timeseries TO instrumentation_reader;


--
-- Name: TABLE v_timeseries_dependency; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_timeseries_dependency TO instrumentation_reader;


--
-- Name: TABLE v_timeseries_latest; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_timeseries_latest TO instrumentation_reader;


--
-- Name: TABLE v_timeseries_project_map; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_timeseries_project_map TO instrumentation_reader;


--
-- Name: TABLE v_unit; Type: ACL; Schema: midas; Owner: postgres
--

GRANT SELECT ON TABLE midas.v_unit TO instrumentation_reader;


--
-- PostgreSQL database dump complete
--

