--
-- PostgreSQL database dump
--

-- Dumped from database version 14.13 (Debian 14.13-1.pgdg120+1)
-- Dumped by pg_dump version 14.13 (Debian 14.13-1.pgdg120+1)

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
-- Name: keycloak; Type: SCHEMA; Schema: -; Owner: keycloak_user
--

CREATE SCHEMA keycloak;


ALTER SCHEMA keycloak OWNER TO keycloak_user;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: admin_event_entity; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.admin_event_entity (
    id character varying(36) NOT NULL,
    admin_event_time bigint,
    realm_id character varying(255),
    operation_type character varying(255),
    auth_realm_id character varying(255),
    auth_client_id character varying(255),
    auth_user_id character varying(255),
    ip_address character varying(255),
    resource_path character varying(2550),
    representation text,
    error character varying(255),
    resource_type character varying(64)
);


ALTER TABLE keycloak.admin_event_entity OWNER TO keycloak_user;

--
-- Name: associated_policy; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.associated_policy (
    policy_id character varying(36) NOT NULL,
    associated_policy_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.associated_policy OWNER TO keycloak_user;

--
-- Name: authentication_execution; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.authentication_execution (
    id character varying(36) NOT NULL,
    alias character varying(255),
    authenticator character varying(36),
    realm_id character varying(36),
    flow_id character varying(36),
    requirement integer,
    priority integer,
    authenticator_flow boolean DEFAULT false NOT NULL,
    auth_flow_id character varying(36),
    auth_config character varying(36)
);


ALTER TABLE keycloak.authentication_execution OWNER TO keycloak_user;

--
-- Name: authentication_flow; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.authentication_flow (
    id character varying(36) NOT NULL,
    alias character varying(255),
    description character varying(255),
    realm_id character varying(36),
    provider_id character varying(36) DEFAULT 'basic-flow'::character varying NOT NULL,
    top_level boolean DEFAULT false NOT NULL,
    built_in boolean DEFAULT false NOT NULL
);


ALTER TABLE keycloak.authentication_flow OWNER TO keycloak_user;

--
-- Name: authenticator_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.authenticator_config (
    id character varying(36) NOT NULL,
    alias character varying(255),
    realm_id character varying(36)
);


ALTER TABLE keycloak.authenticator_config OWNER TO keycloak_user;

--
-- Name: authenticator_config_entry; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.authenticator_config_entry (
    authenticator_id character varying(36) NOT NULL,
    value text,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.authenticator_config_entry OWNER TO keycloak_user;

--
-- Name: broker_link; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.broker_link (
    identity_provider character varying(255) NOT NULL,
    storage_provider_id character varying(255),
    realm_id character varying(36) NOT NULL,
    broker_user_id character varying(255),
    broker_username character varying(255),
    token text,
    user_id character varying(255) NOT NULL
);


ALTER TABLE keycloak.broker_link OWNER TO keycloak_user;

--
-- Name: client; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client (
    id character varying(36) NOT NULL,
    enabled boolean DEFAULT false NOT NULL,
    full_scope_allowed boolean DEFAULT false NOT NULL,
    client_id character varying(255),
    not_before integer,
    public_client boolean DEFAULT false NOT NULL,
    secret character varying(255),
    base_url character varying(255),
    bearer_only boolean DEFAULT false NOT NULL,
    management_url character varying(255),
    surrogate_auth_required boolean DEFAULT false NOT NULL,
    realm_id character varying(36),
    protocol character varying(255),
    node_rereg_timeout integer DEFAULT 0,
    frontchannel_logout boolean DEFAULT false NOT NULL,
    consent_required boolean DEFAULT false NOT NULL,
    name character varying(255),
    service_accounts_enabled boolean DEFAULT false NOT NULL,
    client_authenticator_type character varying(255),
    root_url character varying(255),
    description character varying(255),
    registration_token character varying(255),
    standard_flow_enabled boolean DEFAULT true NOT NULL,
    implicit_flow_enabled boolean DEFAULT false NOT NULL,
    direct_access_grants_enabled boolean DEFAULT false NOT NULL,
    always_display_in_console boolean DEFAULT false NOT NULL
);


ALTER TABLE keycloak.client OWNER TO keycloak_user;

--
-- Name: client_attributes; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_attributes (
    client_id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    value text
);


ALTER TABLE keycloak.client_attributes OWNER TO keycloak_user;

--
-- Name: client_auth_flow_bindings; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_auth_flow_bindings (
    client_id character varying(36) NOT NULL,
    flow_id character varying(36),
    binding_name character varying(255) NOT NULL
);


ALTER TABLE keycloak.client_auth_flow_bindings OWNER TO keycloak_user;

--
-- Name: client_initial_access; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_initial_access (
    id character varying(36) NOT NULL,
    realm_id character varying(36) NOT NULL,
    "timestamp" integer,
    expiration integer,
    count integer,
    remaining_count integer
);


ALTER TABLE keycloak.client_initial_access OWNER TO keycloak_user;

--
-- Name: client_node_registrations; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_node_registrations (
    client_id character varying(36) NOT NULL,
    value integer,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.client_node_registrations OWNER TO keycloak_user;

--
-- Name: client_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_scope (
    id character varying(36) NOT NULL,
    name character varying(255),
    realm_id character varying(36),
    description character varying(255),
    protocol character varying(255)
);


ALTER TABLE keycloak.client_scope OWNER TO keycloak_user;

--
-- Name: client_scope_attributes; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_scope_attributes (
    scope_id character varying(36) NOT NULL,
    value character varying(2048),
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.client_scope_attributes OWNER TO keycloak_user;

--
-- Name: client_scope_client; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_scope_client (
    client_id character varying(255) NOT NULL,
    scope_id character varying(255) NOT NULL,
    default_scope boolean DEFAULT false NOT NULL
);


ALTER TABLE keycloak.client_scope_client OWNER TO keycloak_user;

--
-- Name: client_scope_role_mapping; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_scope_role_mapping (
    scope_id character varying(36) NOT NULL,
    role_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_scope_role_mapping OWNER TO keycloak_user;

--
-- Name: client_session; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_session (
    id character varying(36) NOT NULL,
    client_id character varying(36),
    redirect_uri character varying(255),
    state character varying(255),
    "timestamp" integer,
    session_id character varying(36),
    auth_method character varying(255),
    realm_id character varying(255),
    auth_user_id character varying(36),
    current_action character varying(36)
);


ALTER TABLE keycloak.client_session OWNER TO keycloak_user;

--
-- Name: client_session_auth_status; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_session_auth_status (
    authenticator character varying(36) NOT NULL,
    status integer,
    client_session character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_session_auth_status OWNER TO keycloak_user;

--
-- Name: client_session_note; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_session_note (
    name character varying(255) NOT NULL,
    value character varying(255),
    client_session character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_session_note OWNER TO keycloak_user;

--
-- Name: client_session_prot_mapper; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_session_prot_mapper (
    protocol_mapper_id character varying(36) NOT NULL,
    client_session character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_session_prot_mapper OWNER TO keycloak_user;

--
-- Name: client_session_role; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_session_role (
    role_id character varying(255) NOT NULL,
    client_session character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_session_role OWNER TO keycloak_user;

--
-- Name: client_user_session_note; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.client_user_session_note (
    name character varying(255) NOT NULL,
    value character varying(2048),
    client_session character varying(36) NOT NULL
);


ALTER TABLE keycloak.client_user_session_note OWNER TO keycloak_user;

--
-- Name: component; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.component (
    id character varying(36) NOT NULL,
    name character varying(255),
    parent_id character varying(36),
    provider_id character varying(36),
    provider_type character varying(255),
    realm_id character varying(36),
    sub_type character varying(255)
);


ALTER TABLE keycloak.component OWNER TO keycloak_user;

--
-- Name: component_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.component_config (
    id character varying(36) NOT NULL,
    component_id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(4000)
);


ALTER TABLE keycloak.component_config OWNER TO keycloak_user;

--
-- Name: composite_role; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.composite_role (
    composite character varying(36) NOT NULL,
    child_role character varying(36) NOT NULL
);


ALTER TABLE keycloak.composite_role OWNER TO keycloak_user;

--
-- Name: credential; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.credential (
    id character varying(36) NOT NULL,
    salt bytea,
    type character varying(255),
    user_id character varying(36),
    created_date bigint,
    user_label character varying(255),
    secret_data text,
    credential_data text,
    priority integer
);


ALTER TABLE keycloak.credential OWNER TO keycloak_user;

--
-- Name: databasechangelog; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.databasechangelog (
    id character varying(255) NOT NULL,
    author character varying(255) NOT NULL,
    filename character varying(255) NOT NULL,
    dateexecuted timestamp without time zone NOT NULL,
    orderexecuted integer NOT NULL,
    exectype character varying(10) NOT NULL,
    md5sum character varying(35),
    description character varying(255),
    comments character varying(255),
    tag character varying(255),
    liquibase character varying(20),
    contexts character varying(255),
    labels character varying(255),
    deployment_id character varying(10)
);


ALTER TABLE keycloak.databasechangelog OWNER TO keycloak_user;

--
-- Name: databasechangeloglock; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.databasechangeloglock (
    id integer NOT NULL,
    locked boolean NOT NULL,
    lockgranted timestamp without time zone,
    lockedby character varying(255)
);


ALTER TABLE keycloak.databasechangeloglock OWNER TO keycloak_user;

--
-- Name: default_client_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.default_client_scope (
    realm_id character varying(36) NOT NULL,
    scope_id character varying(36) NOT NULL,
    default_scope boolean DEFAULT false NOT NULL
);


ALTER TABLE keycloak.default_client_scope OWNER TO keycloak_user;

--
-- Name: event_entity; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.event_entity (
    id character varying(36) NOT NULL,
    client_id character varying(255),
    details_json character varying(2550),
    error character varying(255),
    ip_address character varying(255),
    realm_id character varying(255),
    session_id character varying(255),
    event_time bigint,
    type character varying(255),
    user_id character varying(255)
);


ALTER TABLE keycloak.event_entity OWNER TO keycloak_user;

--
-- Name: fed_user_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_attribute (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36),
    value character varying(2024)
);


ALTER TABLE keycloak.fed_user_attribute OWNER TO keycloak_user;

--
-- Name: fed_user_consent; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_consent (
    id character varying(36) NOT NULL,
    client_id character varying(255),
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36),
    created_date bigint,
    last_updated_date bigint,
    client_storage_provider character varying(36),
    external_client_id character varying(255)
);


ALTER TABLE keycloak.fed_user_consent OWNER TO keycloak_user;

--
-- Name: fed_user_consent_cl_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_consent_cl_scope (
    user_consent_id character varying(36) NOT NULL,
    scope_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.fed_user_consent_cl_scope OWNER TO keycloak_user;

--
-- Name: fed_user_credential; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_credential (
    id character varying(36) NOT NULL,
    salt bytea,
    type character varying(255),
    created_date bigint,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36),
    user_label character varying(255),
    secret_data text,
    credential_data text,
    priority integer
);


ALTER TABLE keycloak.fed_user_credential OWNER TO keycloak_user;

--
-- Name: fed_user_group_membership; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_group_membership (
    group_id character varying(36) NOT NULL,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36)
);


ALTER TABLE keycloak.fed_user_group_membership OWNER TO keycloak_user;

--
-- Name: fed_user_required_action; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_required_action (
    required_action character varying(255) DEFAULT ' '::character varying NOT NULL,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36)
);


ALTER TABLE keycloak.fed_user_required_action OWNER TO keycloak_user;

--
-- Name: fed_user_role_mapping; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.fed_user_role_mapping (
    role_id character varying(36) NOT NULL,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    storage_provider_id character varying(36)
);


ALTER TABLE keycloak.fed_user_role_mapping OWNER TO keycloak_user;

--
-- Name: federated_identity; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.federated_identity (
    identity_provider character varying(255) NOT NULL,
    realm_id character varying(36),
    federated_user_id character varying(255),
    federated_username character varying(255),
    token text,
    user_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.federated_identity OWNER TO keycloak_user;

--
-- Name: federated_user; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.federated_user (
    id character varying(255) NOT NULL,
    storage_provider_id character varying(255),
    realm_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.federated_user OWNER TO keycloak_user;

--
-- Name: group_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.group_attribute (
    id character varying(36) DEFAULT 'sybase-needs-something-here'::character varying NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(255),
    group_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.group_attribute OWNER TO keycloak_user;

--
-- Name: group_role_mapping; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.group_role_mapping (
    role_id character varying(36) NOT NULL,
    group_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.group_role_mapping OWNER TO keycloak_user;

--
-- Name: identity_provider; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.identity_provider (
    internal_id character varying(36) NOT NULL,
    enabled boolean DEFAULT false NOT NULL,
    provider_alias character varying(255),
    provider_id character varying(255),
    store_token boolean DEFAULT false NOT NULL,
    authenticate_by_default boolean DEFAULT false NOT NULL,
    realm_id character varying(36),
    add_token_role boolean DEFAULT true NOT NULL,
    trust_email boolean DEFAULT false NOT NULL,
    first_broker_login_flow_id character varying(36),
    post_broker_login_flow_id character varying(36),
    provider_display_name character varying(255),
    link_only boolean DEFAULT false NOT NULL
);


ALTER TABLE keycloak.identity_provider OWNER TO keycloak_user;

--
-- Name: identity_provider_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.identity_provider_config (
    identity_provider_id character varying(36) NOT NULL,
    value text,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.identity_provider_config OWNER TO keycloak_user;

--
-- Name: identity_provider_mapper; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.identity_provider_mapper (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    idp_alias character varying(255) NOT NULL,
    idp_mapper_name character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.identity_provider_mapper OWNER TO keycloak_user;

--
-- Name: idp_mapper_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.idp_mapper_config (
    idp_mapper_id character varying(36) NOT NULL,
    value text,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.idp_mapper_config OWNER TO keycloak_user;

--
-- Name: keycloak_group; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.keycloak_group (
    id character varying(36) NOT NULL,
    name character varying(255),
    parent_group character varying(36) NOT NULL,
    realm_id character varying(36)
);


ALTER TABLE keycloak.keycloak_group OWNER TO keycloak_user;

--
-- Name: keycloak_role; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.keycloak_role (
    id character varying(36) NOT NULL,
    client_realm_constraint character varying(255),
    client_role boolean DEFAULT false NOT NULL,
    description character varying(255),
    name character varying(255),
    realm_id character varying(255),
    client character varying(36),
    realm character varying(36)
);


ALTER TABLE keycloak.keycloak_role OWNER TO keycloak_user;

--
-- Name: migration_model; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.migration_model (
    id character varying(36) NOT NULL,
    version character varying(36),
    update_time bigint DEFAULT 0 NOT NULL
);


ALTER TABLE keycloak.migration_model OWNER TO keycloak_user;

--
-- Name: offline_client_session; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.offline_client_session (
    user_session_id character varying(36) NOT NULL,
    client_id character varying(255) NOT NULL,
    offline_flag character varying(4) NOT NULL,
    "timestamp" integer,
    data text,
    client_storage_provider character varying(36) DEFAULT 'local'::character varying NOT NULL,
    external_client_id character varying(255) DEFAULT 'local'::character varying NOT NULL
);


ALTER TABLE keycloak.offline_client_session OWNER TO keycloak_user;

--
-- Name: offline_user_session; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.offline_user_session (
    user_session_id character varying(36) NOT NULL,
    user_id character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    created_on integer NOT NULL,
    offline_flag character varying(4) NOT NULL,
    data text,
    last_session_refresh integer DEFAULT 0 NOT NULL
);


ALTER TABLE keycloak.offline_user_session OWNER TO keycloak_user;

--
-- Name: policy_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.policy_config (
    policy_id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    value text
);


ALTER TABLE keycloak.policy_config OWNER TO keycloak_user;

--
-- Name: protocol_mapper; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.protocol_mapper (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    protocol character varying(255) NOT NULL,
    protocol_mapper_name character varying(255) NOT NULL,
    client_id character varying(36),
    client_scope_id character varying(36)
);


ALTER TABLE keycloak.protocol_mapper OWNER TO keycloak_user;

--
-- Name: protocol_mapper_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.protocol_mapper_config (
    protocol_mapper_id character varying(36) NOT NULL,
    value text,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.protocol_mapper_config OWNER TO keycloak_user;

--
-- Name: realm; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm (
    id character varying(36) NOT NULL,
    access_code_lifespan integer,
    user_action_lifespan integer,
    access_token_lifespan integer,
    account_theme character varying(255),
    admin_theme character varying(255),
    email_theme character varying(255),
    enabled boolean DEFAULT false NOT NULL,
    events_enabled boolean DEFAULT false NOT NULL,
    events_expiration bigint,
    login_theme character varying(255),
    name character varying(255),
    not_before integer,
    password_policy character varying(2550),
    registration_allowed boolean DEFAULT false NOT NULL,
    remember_me boolean DEFAULT false NOT NULL,
    reset_password_allowed boolean DEFAULT false NOT NULL,
    social boolean DEFAULT false NOT NULL,
    ssl_required character varying(255),
    sso_idle_timeout integer,
    sso_max_lifespan integer,
    update_profile_on_soc_login boolean DEFAULT false NOT NULL,
    verify_email boolean DEFAULT false NOT NULL,
    master_admin_client character varying(36),
    login_lifespan integer,
    internationalization_enabled boolean DEFAULT false NOT NULL,
    default_locale character varying(255),
    reg_email_as_username boolean DEFAULT false NOT NULL,
    admin_events_enabled boolean DEFAULT false NOT NULL,
    admin_events_details_enabled boolean DEFAULT false NOT NULL,
    edit_username_allowed boolean DEFAULT false NOT NULL,
    otp_policy_counter integer DEFAULT 0,
    otp_policy_window integer DEFAULT 1,
    otp_policy_period integer DEFAULT 30,
    otp_policy_digits integer DEFAULT 6,
    otp_policy_alg character varying(36) DEFAULT 'HmacSHA1'::character varying,
    otp_policy_type character varying(36) DEFAULT 'totp'::character varying,
    browser_flow character varying(36),
    registration_flow character varying(36),
    direct_grant_flow character varying(36),
    reset_credentials_flow character varying(36),
    client_auth_flow character varying(36),
    offline_session_idle_timeout integer DEFAULT 0,
    revoke_refresh_token boolean DEFAULT false NOT NULL,
    access_token_life_implicit integer DEFAULT 0,
    login_with_email_allowed boolean DEFAULT true NOT NULL,
    duplicate_emails_allowed boolean DEFAULT false NOT NULL,
    docker_auth_flow character varying(36),
    refresh_token_max_reuse integer DEFAULT 0,
    allow_user_managed_access boolean DEFAULT false NOT NULL,
    sso_max_lifespan_remember_me integer DEFAULT 0 NOT NULL,
    sso_idle_timeout_remember_me integer DEFAULT 0 NOT NULL,
    default_role character varying(255)
);


ALTER TABLE keycloak.realm OWNER TO keycloak_user;

--
-- Name: realm_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_attribute (
    name character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL,
    value text
);


ALTER TABLE keycloak.realm_attribute OWNER TO keycloak_user;

--
-- Name: realm_default_groups; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_default_groups (
    realm_id character varying(36) NOT NULL,
    group_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.realm_default_groups OWNER TO keycloak_user;

--
-- Name: realm_enabled_event_types; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_enabled_event_types (
    realm_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.realm_enabled_event_types OWNER TO keycloak_user;

--
-- Name: realm_events_listeners; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_events_listeners (
    realm_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.realm_events_listeners OWNER TO keycloak_user;

--
-- Name: realm_localizations; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_localizations (
    realm_id character varying(255) NOT NULL,
    locale character varying(255) NOT NULL,
    texts text NOT NULL
);


ALTER TABLE keycloak.realm_localizations OWNER TO keycloak_user;

--
-- Name: realm_required_credential; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_required_credential (
    type character varying(255) NOT NULL,
    form_label character varying(255),
    input boolean DEFAULT false NOT NULL,
    secret boolean DEFAULT false NOT NULL,
    realm_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.realm_required_credential OWNER TO keycloak_user;

--
-- Name: realm_smtp_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_smtp_config (
    realm_id character varying(36) NOT NULL,
    value character varying(255),
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.realm_smtp_config OWNER TO keycloak_user;

--
-- Name: realm_supported_locales; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.realm_supported_locales (
    realm_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.realm_supported_locales OWNER TO keycloak_user;

--
-- Name: redirect_uris; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.redirect_uris (
    client_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.redirect_uris OWNER TO keycloak_user;

--
-- Name: required_action_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.required_action_config (
    required_action_id character varying(36) NOT NULL,
    value text,
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.required_action_config OWNER TO keycloak_user;

--
-- Name: required_action_provider; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.required_action_provider (
    id character varying(36) NOT NULL,
    alias character varying(255),
    name character varying(255),
    realm_id character varying(36),
    enabled boolean DEFAULT false NOT NULL,
    default_action boolean DEFAULT false NOT NULL,
    provider_id character varying(255),
    priority integer
);


ALTER TABLE keycloak.required_action_provider OWNER TO keycloak_user;

--
-- Name: resource_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_attribute (
    id character varying(36) DEFAULT 'sybase-needs-something-here'::character varying NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(255),
    resource_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.resource_attribute OWNER TO keycloak_user;

--
-- Name: resource_policy; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_policy (
    resource_id character varying(36) NOT NULL,
    policy_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.resource_policy OWNER TO keycloak_user;

--
-- Name: resource_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_scope (
    resource_id character varying(36) NOT NULL,
    scope_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.resource_scope OWNER TO keycloak_user;

--
-- Name: resource_server; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_server (
    id character varying(36) NOT NULL,
    allow_rs_remote_mgmt boolean DEFAULT false NOT NULL,
    policy_enforce_mode character varying(15) NOT NULL,
    decision_strategy smallint DEFAULT 1 NOT NULL
);


ALTER TABLE keycloak.resource_server OWNER TO keycloak_user;

--
-- Name: resource_server_perm_ticket; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_server_perm_ticket (
    id character varying(36) NOT NULL,
    owner character varying(255) NOT NULL,
    requester character varying(255) NOT NULL,
    created_timestamp bigint NOT NULL,
    granted_timestamp bigint,
    resource_id character varying(36) NOT NULL,
    scope_id character varying(36),
    resource_server_id character varying(36) NOT NULL,
    policy_id character varying(36)
);


ALTER TABLE keycloak.resource_server_perm_ticket OWNER TO keycloak_user;

--
-- Name: resource_server_policy; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_server_policy (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255),
    type character varying(255) NOT NULL,
    decision_strategy character varying(20),
    logic character varying(20),
    resource_server_id character varying(36) NOT NULL,
    owner character varying(255)
);


ALTER TABLE keycloak.resource_server_policy OWNER TO keycloak_user;

--
-- Name: resource_server_resource; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_server_resource (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255),
    icon_uri character varying(255),
    owner character varying(255) NOT NULL,
    resource_server_id character varying(36) NOT NULL,
    owner_managed_access boolean DEFAULT false NOT NULL,
    display_name character varying(255)
);


ALTER TABLE keycloak.resource_server_resource OWNER TO keycloak_user;

--
-- Name: resource_server_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_server_scope (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    icon_uri character varying(255),
    resource_server_id character varying(36) NOT NULL,
    display_name character varying(255)
);


ALTER TABLE keycloak.resource_server_scope OWNER TO keycloak_user;

--
-- Name: resource_uris; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.resource_uris (
    resource_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.resource_uris OWNER TO keycloak_user;

--
-- Name: role_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.role_attribute (
    id character varying(36) NOT NULL,
    role_id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(255)
);


ALTER TABLE keycloak.role_attribute OWNER TO keycloak_user;

--
-- Name: scope_mapping; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.scope_mapping (
    client_id character varying(36) NOT NULL,
    role_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.scope_mapping OWNER TO keycloak_user;

--
-- Name: scope_policy; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.scope_policy (
    scope_id character varying(36) NOT NULL,
    policy_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.scope_policy OWNER TO keycloak_user;

--
-- Name: user_attribute; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_attribute (
    name character varying(255) NOT NULL,
    value character varying(255),
    user_id character varying(36) NOT NULL,
    id character varying(36) DEFAULT 'sybase-needs-something-here'::character varying NOT NULL
);


ALTER TABLE keycloak.user_attribute OWNER TO keycloak_user;

--
-- Name: user_consent; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_consent (
    id character varying(36) NOT NULL,
    client_id character varying(255),
    user_id character varying(36) NOT NULL,
    created_date bigint,
    last_updated_date bigint,
    client_storage_provider character varying(36),
    external_client_id character varying(255)
);


ALTER TABLE keycloak.user_consent OWNER TO keycloak_user;

--
-- Name: user_consent_client_scope; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_consent_client_scope (
    user_consent_id character varying(36) NOT NULL,
    scope_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.user_consent_client_scope OWNER TO keycloak_user;

--
-- Name: user_entity; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_entity (
    id character varying(36) NOT NULL,
    email character varying(255),
    email_constraint character varying(255),
    email_verified boolean DEFAULT false NOT NULL,
    enabled boolean DEFAULT false NOT NULL,
    federation_link character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    realm_id character varying(255),
    username character varying(255),
    created_timestamp bigint,
    service_account_client_link character varying(255),
    not_before integer DEFAULT 0 NOT NULL
);


ALTER TABLE keycloak.user_entity OWNER TO keycloak_user;

--
-- Name: user_federation_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_federation_config (
    user_federation_provider_id character varying(36) NOT NULL,
    value character varying(255),
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.user_federation_config OWNER TO keycloak_user;

--
-- Name: user_federation_mapper; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_federation_mapper (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    federation_provider_id character varying(36) NOT NULL,
    federation_mapper_type character varying(255) NOT NULL,
    realm_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.user_federation_mapper OWNER TO keycloak_user;

--
-- Name: user_federation_mapper_config; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_federation_mapper_config (
    user_federation_mapper_id character varying(36) NOT NULL,
    value character varying(255),
    name character varying(255) NOT NULL
);


ALTER TABLE keycloak.user_federation_mapper_config OWNER TO keycloak_user;

--
-- Name: user_federation_provider; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_federation_provider (
    id character varying(36) NOT NULL,
    changed_sync_period integer,
    display_name character varying(255),
    full_sync_period integer,
    last_sync integer,
    priority integer,
    provider_name character varying(255),
    realm_id character varying(36)
);


ALTER TABLE keycloak.user_federation_provider OWNER TO keycloak_user;

--
-- Name: user_group_membership; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_group_membership (
    group_id character varying(36) NOT NULL,
    user_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.user_group_membership OWNER TO keycloak_user;

--
-- Name: user_required_action; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_required_action (
    user_id character varying(36) NOT NULL,
    required_action character varying(255) DEFAULT ' '::character varying NOT NULL
);


ALTER TABLE keycloak.user_required_action OWNER TO keycloak_user;

--
-- Name: user_role_mapping; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_role_mapping (
    role_id character varying(255) NOT NULL,
    user_id character varying(36) NOT NULL
);


ALTER TABLE keycloak.user_role_mapping OWNER TO keycloak_user;

--
-- Name: user_session; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_session (
    id character varying(36) NOT NULL,
    auth_method character varying(255),
    ip_address character varying(255),
    last_session_refresh integer,
    login_username character varying(255),
    realm_id character varying(255),
    remember_me boolean DEFAULT false NOT NULL,
    started integer,
    user_id character varying(255),
    user_session_state integer,
    broker_session_id character varying(255),
    broker_user_id character varying(255)
);


ALTER TABLE keycloak.user_session OWNER TO keycloak_user;

--
-- Name: user_session_note; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.user_session_note (
    user_session character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(2048)
);


ALTER TABLE keycloak.user_session_note OWNER TO keycloak_user;

--
-- Name: username_login_failure; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.username_login_failure (
    realm_id character varying(36) NOT NULL,
    username character varying(255) NOT NULL,
    failed_login_not_before integer,
    last_failure bigint,
    last_ip_failure character varying(255),
    num_failures integer
);


ALTER TABLE keycloak.username_login_failure OWNER TO keycloak_user;

--
-- Name: web_origins; Type: TABLE; Schema: keycloak; Owner: keycloak_user
--

CREATE TABLE keycloak.web_origins (
    client_id character varying(36) NOT NULL,
    value character varying(255) NOT NULL
);


ALTER TABLE keycloak.web_origins OWNER TO keycloak_user;

--
-- Data for Name: admin_event_entity; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.admin_event_entity (id, admin_event_time, realm_id, operation_type, auth_realm_id, auth_client_id, auth_user_id, ip_address, resource_path, representation, error, resource_type) FROM stdin;
d8a86cfa-6cbb-4526-8d36-9daace636c2a	1719326158261	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.21.0.1	users/1c6decce-7042-4a2c-9880-ac7c5bc27b8c	\N	\N	USER
99ff61c2-96e1-4f01-bc43-f964b8bfb2e5	1719326170271	cwbi	ACTION	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.21.0.1	users/1c6decce-7042-4a2c-9880-ac7c5bc27b8c/reset-password	\N	\N	USER
abecd34e-b852-4786-9c7a-fc3e5bbb9192	1719327556360	cwbi	DELETE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.23.0.1	users/1c6decce-7042-4a2c-9880-ac7c5bc27b8c	\N	\N	USER
ad88d2ca-8bb4-4808-957f-39f233ee3537	1719327608110	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.23.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
92e80107-7868-49db-82b0-8e6781fc7dd0	1719327627811	cwbi	ACTION	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.23.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4/reset-password	\N	\N	USER
90c7e09a-f304-4b7b-80b1-5982aed23532	1719327657137	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.23.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
59e263a4-ac83-48eb-a7bb-670fd34e31cb	1719331815993	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.26.0.1	default-default-client-scopes/5286fee9-6cda-4a94-aba0-dffa0a5c2e8f	\N	\N	CLIENT_SCOPE
2386f610-be40-454a-bfda-967f96cb1314	1719331818597	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.26.0.1	default-default-client-scopes/9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0	\N	\N	CLIENT_SCOPE
320c6d73-3204-47d8-9135-3f93cdc728ab	1719331823951	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.26.0.1	default-default-client-scopes/b0a33b5f-7c9a-4d59-9602-855dfb2a0b92	\N	\N	CLIENT_SCOPE
2416f01e-d2b8-431a-9db5-b9fdd04a4d47	1719335640559	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
f035ca27-c79f-4041-a048-9376a630b5ff	1719335682908	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
e5d5157e-8dbd-402c-a59f-9a199b49e68f	1719336806686	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
ee843928-1968-432f-81c9-5cfceaa738de	1719336893542	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
11aca3bc-e5fc-4635-98a7-c43640856c83	1719337138964	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
6628f534-8452-4990-b34e-e1f5341ec981	1719337197474	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.27.0.1	clients/86b97bc5-1afd-40b2-ad62-bddaaaf321c7/protocol-mappers/models/d6680f62-2652-47d0-9b6c-6b7eae6d8c34	\N	\N	PROTOCOL_MAPPER
0e950623-a7fc-4071-90d0-7c59d9e885eb	1719428176784	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
f93b4615-888b-44d3-bd1c-032f2ff5961e	1719428396949	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
bb06a93a-77a8-45c0-b22b-204438ce8d2a	1719428401274	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
b58d1905-7272-4b7f-afca-ef7b897acbc5	1719428418986	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
30281760-87ca-44bb-a652-e52093eeba14	1719428445266	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
1283997b-4279-48a4-8280-2ce859fa06c2	1719428513022	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	clients/86b97bc5-1afd-40b2-ad62-bddaaaf321c7/protocol-mappers/models/1e3b6b8c-e208-4ef5-bc46-02ab17db4808	\N	\N	PROTOCOL_MAPPER
f675ebb2-8c9f-4d74-a5b9-ce32818865c4	1719428810190	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
00c1138b-f523-4985-8e65-ec91a12e25e3	1719429331334	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.28.0.1	users/f8dcafea-243e-4b89-8d7d-fa01918130f4	\N	\N	USER
19fb8d0c-baac-4d81-936a-4908633550af	1719438836121	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	192.168.32.1	users/127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a	\N	\N	USER
4380573e-54a8-46f5-a963-800ba182aa7d	1719438849825	cwbi	ACTION	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	192.168.32.1	users/127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a/reset-password	\N	\N	USER
76f3fc10-4508-4fb4-adb3-e309d34b41d6	1719438907833	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	192.168.32.1	users/127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a	\N	\N	USER
312690b1-fc76-493e-8f58-720216ca62ba	1720198846971	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.19.0.1	\N	\N	\N	REALM
0a22a727-bace-489d-93dd-92e846002e27	1725660663733	cwbi	ACTION	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f/reset-password	\N	\N	USER
c479b33e-903c-485e-9603-7a0e167f2102	1725660848446	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f	\N	\N	USER
4cf8efda-6550-4449-9604-054e448f6e49	1725660929672	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f	\N	\N	USER
566f91fe-d63c-4425-8c7a-45b5ba292660	1725660959411	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f	\N	\N	USER
bd0538bd-841f-4571-8f57-54d647d6ea73	1725661066497	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f	\N	\N	USER
2987d079-026a-482d-ab24-2827c0475574	1725661068753	cwbi	UPDATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	users/f9b33064-13d0-47d7-8294-fb8f0fac819f	\N	\N	USER
bf318ee8-4e38-4444-97a9-c20cbd7001cd	1725661279432	cwbi	CREATE	98749fe9-5c8f-4d46-b973-16664c916f0f	fca2fb0d-1434-4ba2-bd0a-699e623e79be	f3fc1dd9-af7b-498a-9435-31da080a37ad	172.18.0.1	clients/86b97bc5-1afd-40b2-ad62-bddaaaf321c7/protocol-mappers/models/cfe591a3-b08f-408b-90b7-a00eb96388d2	\N	\N	PROTOCOL_MAPPER
\.


--
-- Data for Name: associated_policy; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.associated_policy (policy_id, associated_policy_id) FROM stdin;
\.


--
-- Data for Name: authentication_execution; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.authentication_execution (id, alias, authenticator, realm_id, flow_id, requirement, priority, authenticator_flow, auth_flow_id, auth_config) FROM stdin;
6c1b2c0e-9861-42ba-83c9-2a415c6af823	\N	auth-cookie	98749fe9-5c8f-4d46-b973-16664c916f0f	ee2d0274-f49c-44f6-a979-7e7f86176265	2	10	f	\N	\N
cff8e04b-ec82-4070-bd8c-6f48b5240408	\N	auth-spnego	98749fe9-5c8f-4d46-b973-16664c916f0f	ee2d0274-f49c-44f6-a979-7e7f86176265	3	20	f	\N	\N
a30ff8d8-4126-4b29-ad47-ef7628aeb235	\N	identity-provider-redirector	98749fe9-5c8f-4d46-b973-16664c916f0f	ee2d0274-f49c-44f6-a979-7e7f86176265	2	25	f	\N	\N
54c1150c-7914-4549-ab9c-e384e9e28a35	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	ee2d0274-f49c-44f6-a979-7e7f86176265	2	30	t	9ff67418-36f6-4e75-8b00-07924e64fb79	\N
e5b51b7f-a8b6-495d-919c-bfb0d32d2848	\N	auth-username-password-form	98749fe9-5c8f-4d46-b973-16664c916f0f	9ff67418-36f6-4e75-8b00-07924e64fb79	0	10	f	\N	\N
2eb79d12-53fb-4406-ad82-5c5d4daeebe5	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	9ff67418-36f6-4e75-8b00-07924e64fb79	1	20	t	c6451566-012d-4d58-8e98-4e5fcdb9ff07	\N
e2560687-c3c5-44db-9c34-898e5c14fb88	\N	conditional-user-configured	98749fe9-5c8f-4d46-b973-16664c916f0f	c6451566-012d-4d58-8e98-4e5fcdb9ff07	0	10	f	\N	\N
6d0853ba-957e-4f90-aa8b-5896dfaa37b2	\N	auth-otp-form	98749fe9-5c8f-4d46-b973-16664c916f0f	c6451566-012d-4d58-8e98-4e5fcdb9ff07	0	20	f	\N	\N
8a065f5e-2931-46e0-b658-2193b746d454	\N	direct-grant-validate-username	98749fe9-5c8f-4d46-b973-16664c916f0f	ea02467e-0b93-485a-8bf2-e401083b818f	0	10	f	\N	\N
3db4d432-8da0-47af-8dfd-6c809c085835	\N	direct-grant-validate-password	98749fe9-5c8f-4d46-b973-16664c916f0f	ea02467e-0b93-485a-8bf2-e401083b818f	0	20	f	\N	\N
0711a713-e8b2-457d-96cc-f81ce76d69fc	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	ea02467e-0b93-485a-8bf2-e401083b818f	1	30	t	334b88f4-3255-4a09-96a5-727c8f8fe2bb	\N
23299bc7-5090-4f6a-b91c-aa231883c302	\N	conditional-user-configured	98749fe9-5c8f-4d46-b973-16664c916f0f	334b88f4-3255-4a09-96a5-727c8f8fe2bb	0	10	f	\N	\N
cdcf848b-5166-472d-8a9e-484dfc78ef94	\N	direct-grant-validate-otp	98749fe9-5c8f-4d46-b973-16664c916f0f	334b88f4-3255-4a09-96a5-727c8f8fe2bb	0	20	f	\N	\N
0ffd284b-6fba-4f73-94e4-7fbb96173335	\N	registration-page-form	98749fe9-5c8f-4d46-b973-16664c916f0f	6cda22fb-db28-4a17-89cf-fca653249993	0	10	t	a41c6a42-8ec0-42d9-890b-ab31dba5888c	\N
3a55898f-8c04-45f2-b715-6ae008d9487f	\N	registration-user-creation	98749fe9-5c8f-4d46-b973-16664c916f0f	a41c6a42-8ec0-42d9-890b-ab31dba5888c	0	20	f	\N	\N
d30f865d-4868-453d-bab8-0197fb4d0193	\N	registration-profile-action	98749fe9-5c8f-4d46-b973-16664c916f0f	a41c6a42-8ec0-42d9-890b-ab31dba5888c	0	40	f	\N	\N
df4efd2c-d1f3-4d34-ac0a-cf7b7d8023cb	\N	registration-password-action	98749fe9-5c8f-4d46-b973-16664c916f0f	a41c6a42-8ec0-42d9-890b-ab31dba5888c	0	50	f	\N	\N
7d3303a6-3c64-45f2-b4c8-96f8303f7a81	\N	registration-recaptcha-action	98749fe9-5c8f-4d46-b973-16664c916f0f	a41c6a42-8ec0-42d9-890b-ab31dba5888c	3	60	f	\N	\N
fa4344fe-f0c8-41fd-a08f-d24871d77847	\N	reset-credentials-choose-user	98749fe9-5c8f-4d46-b973-16664c916f0f	602d6901-12a3-4e29-8bf1-98757e0a2dac	0	10	f	\N	\N
2254fd6b-e18c-4886-9c99-c54542b889ed	\N	reset-credential-email	98749fe9-5c8f-4d46-b973-16664c916f0f	602d6901-12a3-4e29-8bf1-98757e0a2dac	0	20	f	\N	\N
8d282125-93b8-45a4-9a63-09068daf676e	\N	reset-password	98749fe9-5c8f-4d46-b973-16664c916f0f	602d6901-12a3-4e29-8bf1-98757e0a2dac	0	30	f	\N	\N
d7363287-4971-4dd7-8481-47d56fb9a0d3	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	602d6901-12a3-4e29-8bf1-98757e0a2dac	1	40	t	3580988e-c034-45c0-b6a3-25fca9d926a5	\N
fa397502-dbd5-43d0-8ac7-a163b639d8c9	\N	conditional-user-configured	98749fe9-5c8f-4d46-b973-16664c916f0f	3580988e-c034-45c0-b6a3-25fca9d926a5	0	10	f	\N	\N
072c815c-ad8d-41ab-b9d0-6b87fb1e3605	\N	reset-otp	98749fe9-5c8f-4d46-b973-16664c916f0f	3580988e-c034-45c0-b6a3-25fca9d926a5	0	20	f	\N	\N
4c76703e-dd0e-44a8-82a5-5861aeb893ea	\N	client-secret	98749fe9-5c8f-4d46-b973-16664c916f0f	e58dc8c5-51b0-4ce5-a7be-e4270afd426e	2	10	f	\N	\N
4b38a5c5-b147-486f-93a4-f5cc0834a942	\N	client-jwt	98749fe9-5c8f-4d46-b973-16664c916f0f	e58dc8c5-51b0-4ce5-a7be-e4270afd426e	2	20	f	\N	\N
051c4969-3936-47f6-8517-8eeb21ebc4ac	\N	client-secret-jwt	98749fe9-5c8f-4d46-b973-16664c916f0f	e58dc8c5-51b0-4ce5-a7be-e4270afd426e	2	30	f	\N	\N
11bedfc3-85ff-4309-aaf2-b4c66e69e66a	\N	client-x509	98749fe9-5c8f-4d46-b973-16664c916f0f	e58dc8c5-51b0-4ce5-a7be-e4270afd426e	2	40	f	\N	\N
60411a8a-9e9d-453b-8217-debbc1afb270	\N	idp-review-profile	98749fe9-5c8f-4d46-b973-16664c916f0f	3831e5bd-c4ad-4bbd-b4a0-0403d6bf04d5	0	10	f	\N	b572197f-4f16-4fe3-bf88-946e2c224bd1
93c72af4-2202-4502-8067-833192f878e6	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	3831e5bd-c4ad-4bbd-b4a0-0403d6bf04d5	0	20	t	5f8084d8-c868-4ab6-bc1f-7202cb13ea52	\N
df3f83a4-c1e9-4fc4-bd02-2dc05981708e	\N	idp-create-user-if-unique	98749fe9-5c8f-4d46-b973-16664c916f0f	5f8084d8-c868-4ab6-bc1f-7202cb13ea52	2	10	f	\N	6f9ec3b4-f0b4-46fc-bd49-765e5a6635cc
0538cc5c-7717-4a29-b95d-2aed9631b000	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	5f8084d8-c868-4ab6-bc1f-7202cb13ea52	2	20	t	46e76e1d-d6bc-4c58-993b-3f5a4b9849b6	\N
60d6a071-d53b-403e-8446-7fb96d485891	\N	idp-confirm-link	98749fe9-5c8f-4d46-b973-16664c916f0f	46e76e1d-d6bc-4c58-993b-3f5a4b9849b6	0	10	f	\N	\N
0fcb7115-25ff-4b49-bb5e-4e667f774e36	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	46e76e1d-d6bc-4c58-993b-3f5a4b9849b6	0	20	t	86940e2a-1c4e-4dab-9e3b-03e1998c7680	\N
9c9d3772-a97f-4bc3-911f-81e988f878af	\N	idp-email-verification	98749fe9-5c8f-4d46-b973-16664c916f0f	86940e2a-1c4e-4dab-9e3b-03e1998c7680	2	10	f	\N	\N
e3c38878-8b8f-4be7-bcd2-1e09901b3d9f	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	86940e2a-1c4e-4dab-9e3b-03e1998c7680	2	20	t	ccffa12a-e45b-48c8-93be-03c650658985	\N
3a5e6ee5-df07-4be9-8a86-ebdc0086ab2c	\N	idp-username-password-form	98749fe9-5c8f-4d46-b973-16664c916f0f	ccffa12a-e45b-48c8-93be-03c650658985	0	10	f	\N	\N
e44817ae-d49b-4d82-b2fd-3f9ab55887ca	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	ccffa12a-e45b-48c8-93be-03c650658985	1	20	t	ea9e03db-9bc2-4310-a29d-0c532df73c7f	\N
2457f37a-02bb-4f35-9613-0e280a65a631	\N	conditional-user-configured	98749fe9-5c8f-4d46-b973-16664c916f0f	ea9e03db-9bc2-4310-a29d-0c532df73c7f	0	10	f	\N	\N
e4f52552-455e-45e7-8b73-bda460bd2a60	\N	auth-otp-form	98749fe9-5c8f-4d46-b973-16664c916f0f	ea9e03db-9bc2-4310-a29d-0c532df73c7f	0	20	f	\N	\N
1724999e-1124-4fbb-a0a5-7f6564303a71	\N	http-basic-authenticator	98749fe9-5c8f-4d46-b973-16664c916f0f	67005fb1-bb0d-4f5b-baa1-39ba801af769	0	10	f	\N	\N
8ba92bf7-31b5-4ec2-b7db-a65f01e14bb1	\N	docker-http-basic-authenticator	98749fe9-5c8f-4d46-b973-16664c916f0f	b7d46802-4ce3-4643-bd2d-a5f250726a36	0	10	f	\N	\N
21ad3cef-c5d9-4ba5-a424-1dbb01307098	\N	no-cookie-redirect	98749fe9-5c8f-4d46-b973-16664c916f0f	78d2f3e5-b3a6-4310-8d08-892ac5ecb67f	0	10	f	\N	\N
c183c24b-e3ee-446b-958b-361eda0883e7	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	78d2f3e5-b3a6-4310-8d08-892ac5ecb67f	0	20	t	35737844-3a52-45bd-9a01-0e22a7eb14af	\N
f430294a-20a0-4266-a7f6-c6db16941bb4	\N	basic-auth	98749fe9-5c8f-4d46-b973-16664c916f0f	35737844-3a52-45bd-9a01-0e22a7eb14af	0	10	f	\N	\N
3725ca36-18d8-47d6-90cc-ec5f649987e0	\N	basic-auth-otp	98749fe9-5c8f-4d46-b973-16664c916f0f	35737844-3a52-45bd-9a01-0e22a7eb14af	3	20	f	\N	\N
56f2c7a4-c1b6-4a60-b54a-f417d4fc2302	\N	auth-spnego	98749fe9-5c8f-4d46-b973-16664c916f0f	35737844-3a52-45bd-9a01-0e22a7eb14af	3	30	f	\N	\N
e965de1f-8ae5-4c57-a6c0-35ffa30b6c9a	\N	idp-email-verification	cwbi	aa23b4d3-2163-4e10-b30d-4d65777c07b0	2	10	f	\N	\N
118584b9-88ca-4f5c-b87d-e537dbaa8219	\N	\N	cwbi	aa23b4d3-2163-4e10-b30d-4d65777c07b0	2	20	t	a1ee61b6-25d3-4038-8e01-6c0e95972f41	\N
e0693e83-f3ff-448a-8a40-6c9e983ce10c	\N	basic-auth	cwbi	d741433a-cee2-4ffe-b6c5-f6e389a439c1	0	10	f	\N	\N
14255375-67a3-4d4d-80b1-916f34d1df2c	\N	basic-auth-otp	cwbi	d741433a-cee2-4ffe-b6c5-f6e389a439c1	3	20	f	\N	\N
f0441ec3-0b14-4a83-969e-d591332e4c73	\N	auth-spnego	cwbi	d741433a-cee2-4ffe-b6c5-f6e389a439c1	3	30	f	\N	\N
dea05629-eb70-414c-8efc-13c1584cd4fd	\N	idp-create-user-if-unique	cwbi	771a70a5-cf3f-42f1-bb13-ece0b626389f	2	0	f	\N	\N
790b2048-8887-4457-bddb-518e021a77cb	\N	idp-auto-link	cwbi	771a70a5-cf3f-42f1-bb13-ece0b626389f	2	1	f	\N	\N
27194975-8b2a-4e9a-aeea-0634bf89e5ff	\N	conditional-user-configured	cwbi	d5ad35fd-d118-4869-bfc2-83702ac3076a	0	10	f	\N	\N
06068a48-8fbc-413a-a5b2-c4c002cc72ff	\N	auth-otp-form	cwbi	d5ad35fd-d118-4869-bfc2-83702ac3076a	0	20	f	\N	\N
4448e116-8b86-4c78-920a-0047b535c86a	\N	dls-x509-authentication-factory	cwbi	0679fb07-660f-4278-a965-820043ead874	0	0	f	\N	\N
fd44227d-d4d0-4e9b-b1d6-6a3577d0b8de	\N	conditional-user-configured	cwbi	c01b6492-7ac5-4f85-a085-e519e5bc0dea	0	10	f	\N	\N
6cd13b30-0c69-40cd-b00e-df148b2ee8b7	\N	direct-grant-validate-otp	cwbi	c01b6492-7ac5-4f85-a085-e519e5bc0dea	0	20	f	\N	\N
a4b8e00e-20eb-40f6-886e-5489b3ab786a	\N	conditional-user-configured	cwbi	3108dc5d-0b1c-427e-83ad-3ae677e1476a	0	10	f	\N	\N
67aa1607-ecff-4222-b67e-3a74d152d91e	\N	auth-otp-form	cwbi	3108dc5d-0b1c-427e-83ad-3ae677e1476a	0	20	f	\N	\N
1f9f7e7f-69c2-4299-a7a8-7701e9eb6245	\N	idp-confirm-link	cwbi	7aa06afe-f1f1-4c23-8b19-c534b99a7fd5	0	10	f	\N	\N
58fdd6fa-820d-4eae-a5f8-b2f3c28fc272	\N	\N	cwbi	7aa06afe-f1f1-4c23-8b19-c534b99a7fd5	0	20	t	aa23b4d3-2163-4e10-b30d-4d65777c07b0	\N
ca496d05-7869-483e-a034-d7ae73d303d1	\N	conditional-user-configured	cwbi	de71faee-a3f7-495a-a7ff-8fc8ac3a6ded	0	10	f	\N	\N
1942ff9d-323a-4ee0-b072-b293ccaae9b2	\N	reset-otp	cwbi	de71faee-a3f7-495a-a7ff-8fc8ac3a6ded	0	20	f	\N	\N
6460f960-983c-41fd-9edc-277d2b66e5ea	\N	idp-create-user-if-unique	cwbi	dbcabf07-be0e-457c-9fa7-7d091c42e2ce	2	10	f	\N	1a7024fc-ed33-41d1-a2bf-463b2e5bc744
9a7d148b-0a5d-4ca4-a0ec-374133066ff8	\N	\N	cwbi	dbcabf07-be0e-457c-9fa7-7d091c42e2ce	2	20	t	7aa06afe-f1f1-4c23-8b19-c534b99a7fd5	\N
de03f086-7d97-4256-94a5-2a398ba3016d	\N	idp-username-password-form	cwbi	a1ee61b6-25d3-4038-8e01-6c0e95972f41	0	10	f	\N	\N
f6bcb15e-2d5b-4800-aed8-6f3a891fbe9d	\N	\N	cwbi	a1ee61b6-25d3-4038-8e01-6c0e95972f41	1	20	t	3108dc5d-0b1c-427e-83ad-3ae677e1476a	\N
9bda4447-7ab4-4ce0-9eb3-f673924ab321	\N	auth-cookie	cwbi	fc78af35-3d57-4224-9087-7a4408ab4194	2	10	f	\N	\N
a3466ca7-239e-4801-954f-5931a9f54594	\N	auth-spnego	cwbi	fc78af35-3d57-4224-9087-7a4408ab4194	3	20	f	\N	\N
400441cb-7e44-467b-921e-093ee1ef0466	\N	identity-provider-redirector	cwbi	fc78af35-3d57-4224-9087-7a4408ab4194	2	25	f	\N	\N
3c6caaf1-e9cc-4208-8456-c935f14c306e	\N	\N	cwbi	fc78af35-3d57-4224-9087-7a4408ab4194	2	30	t	ff1f4e39-45e6-4ab5-b43d-74b527959e3d	\N
a6d4db91-4473-468e-8df2-07a6063c4acd	\N	auth-cookie	cwbi	0b8efdc6-045f-4e42-b086-730f8d817711	2	10	f	\N	\N
8722df72-6711-4d7d-a240-a04f5da63d37	\N	auth-spnego	cwbi	0b8efdc6-045f-4e42-b086-730f8d817711	3	20	f	\N	\N
4e773e11-9b89-4f56-ab8c-a4689eda35ba	\N	identity-provider-redirector	cwbi	0b8efdc6-045f-4e42-b086-730f8d817711	2	25	f	\N	ac7e58de-53b2-42d1-a708-65ec837458cc
beb52bfe-ddaa-4a16-85b3-88662e1ba870	\N	dls-x509-authentication-factory	cwbi	0b8efdc6-045f-4e42-b086-730f8d817711	3	30	f	\N	\N
18db52d3-78ad-4f17-889a-49ad0bfb6448	\N	\N	cwbi	0b8efdc6-045f-4e42-b086-730f8d817711	1	31	t	e6654661-fa03-4c49-9728-2c9df094c315	\N
6c8576ba-a39b-4db9-b70c-632a5c73c678	\N	conditional-user-configured	cwbi	0d25d6f7-a29e-4960-af91-9502967dc934	0	10	f	\N	\N
386ec97a-20ec-4891-9280-ed4f38d3eb18	\N	auth-otp-form	cwbi	0d25d6f7-a29e-4960-af91-9502967dc934	0	20	f	\N	\N
aaf885a3-5736-4e28-8f00-254a86a85b57	\N	auth-username-password-form	cwbi	e6654661-fa03-4c49-9728-2c9df094c315	0	10	f	\N	\N
9286a1de-a23c-4942-9086-6fd53793204b	\N	\N	cwbi	e6654661-fa03-4c49-9728-2c9df094c315	1	20	t	0d25d6f7-a29e-4960-af91-9502967dc934	\N
6eaa6748-f1f6-4e98-a618-01e9d120745a	\N	auth-cookie	cwbi	381d1702-a207-4abe-a834-6de732f058af	2	10	f	\N	\N
8cec3419-9aa9-4676-865d-f083ed4dfed0	\N	auth-spnego	cwbi	381d1702-a207-4abe-a834-6de732f058af	3	20	f	\N	\N
5c724120-d5ed-4a91-a950-92da7474208c	\N	identity-provider-redirector	cwbi	381d1702-a207-4abe-a834-6de732f058af	0	25	f	\N	ac7e58de-53b2-42d1-a708-65ec837458cc
1395b46e-5786-431f-8cba-6fd272211634	\N	\N	cwbi	381d1702-a207-4abe-a834-6de732f058af	2	30	t	7dde2057-fbeb-4dc3-85f4-04acbe9bf9bf	\N
216a2926-bcf0-48ae-8723-bfdfa5b2e474	\N	conditional-user-configured	cwbi	d86e6e71-d763-41ce-a68d-b83ccbe72f8d	0	10	f	\N	\N
b35e062c-7f93-4f8e-8846-9b8fb9c90004	\N	auth-otp-form	cwbi	d86e6e71-d763-41ce-a68d-b83ccbe72f8d	0	20	f	\N	\N
eb8d8cd2-2974-4804-98c7-985ae9d3ced0	\N	auth-username-password-form	cwbi	7dde2057-fbeb-4dc3-85f4-04acbe9bf9bf	0	10	f	\N	\N
c6b54006-7712-405f-afec-1779507ca40d	\N	\N	cwbi	7dde2057-fbeb-4dc3-85f4-04acbe9bf9bf	1	20	t	d86e6e71-d763-41ce-a68d-b83ccbe72f8d	\N
3e4020cc-584f-4e0f-ac5b-401d338440b0	\N	auth-cookie	cwbi	f423039a-26be-4d4e-89b5-bc5d821db555	2	10	f	\N	\N
a41ed9a1-d9ce-4395-be19-82c48a66f0ac	\N	auth-spnego	cwbi	f423039a-26be-4d4e-89b5-bc5d821db555	3	20	f	\N	\N
60575366-9b90-4594-9469-b88314fd8cad	\N	identity-provider-redirector	cwbi	f423039a-26be-4d4e-89b5-bc5d821db555	2	25	f	\N	ceb288ff-55ba-4e5e-b3db-0d8af4a44671
55708c7a-27cd-4673-8974-2a1010590d55	\N	\N	cwbi	f423039a-26be-4d4e-89b5-bc5d821db555	2	30	t	fc863779-cc81-43d2-b8df-b2c2e83289f5	\N
f3f6a47a-25fd-493e-9703-601673a4b314	\N	conditional-user-configured	cwbi	5b280719-b0f6-42ed-9dba-c3e197f873e8	0	10	f	\N	\N
ba1289b1-3e64-4749-beb3-6616d1a9ef7a	\N	auth-otp-form	cwbi	5b280719-b0f6-42ed-9dba-c3e197f873e8	0	20	f	\N	\N
63283865-f9b1-40d2-917e-17bf3567775e	\N	auth-username-password-form	cwbi	fc863779-cc81-43d2-b8df-b2c2e83289f5	0	10	f	\N	\N
f91f33ad-6e49-4346-b2f9-26ff494d9fb1	\N	\N	cwbi	fc863779-cc81-43d2-b8df-b2c2e83289f5	1	20	t	5b280719-b0f6-42ed-9dba-c3e197f873e8	\N
999abaaf-3cad-4ec2-bbd1-d0f4e04cca0c	\N	auth-cookie	cwbi	18cb7e12-5b9b-4f6a-8231-6b2725fea5c9	2	10	f	\N	\N
33a65cca-e5e0-43c4-a632-565311d131b3	\N	dls-x509-authentication-factory	cwbi	18cb7e12-5b9b-4f6a-8231-6b2725fea5c9	3	30	f	\N	\N
9dc9c6ea-5490-45f5-87ba-38a527d563af	\N	\N	cwbi	18cb7e12-5b9b-4f6a-8231-6b2725fea5c9	2	31	t	0ac897a7-a668-4207-bb5c-1f1fc6fdadd5	\N
9ed0b25f-cbb2-4c66-b95b-299190dcdae5	\N	auth-username-password-form	cwbi	0ac897a7-a668-4207-bb5c-1f1fc6fdadd5	0	10	f	\N	\N
177d303b-f76d-4126-93a2-1b2bead0d091	\N	client-secret	cwbi	eaf2aa1e-6359-43ca-b403-c169273acae9	2	10	f	\N	\N
07784987-7fba-4b5f-9202-4597213ce718	\N	client-jwt	cwbi	eaf2aa1e-6359-43ca-b403-c169273acae9	2	20	f	\N	\N
a1ead7d8-31bf-4beb-819e-4c804484454e	\N	client-secret-jwt	cwbi	eaf2aa1e-6359-43ca-b403-c169273acae9	2	30	f	\N	\N
c56a8fe0-a0d7-47f8-b29b-46225fb77e6a	\N	client-x509	cwbi	eaf2aa1e-6359-43ca-b403-c169273acae9	2	40	f	\N	\N
60cd9a21-4119-4c8c-8c65-e9eed25f9690	\N	idp-review-profile	cwbi	c8896e58-9c92-4390-ba29-47cd7f4522bf	0	10	f	\N	61e29fc3-c1dd-4f3c-aa03-455676504530
877263ad-0994-491e-84fd-ab49c683479a	\N	\N	cwbi	c8896e58-9c92-4390-ba29-47cd7f4522bf	0	20	t	027cf582-0a1f-429d-93e0-549cc194c933	\N
baee6451-3f92-40f2-99e8-e175dfd31175	\N	conditional-user-configured	cwbi	6cd4395e-ef43-47d5-86a1-c7e97cbde27d	0	10	f	\N	\N
fef2e32b-a2fa-463b-a33e-8a8dede0ac20	\N	auth-otp-form	cwbi	6cd4395e-ef43-47d5-86a1-c7e97cbde27d	0	20	f	\N	\N
bda9f77e-a246-47b7-a764-06dc9dfdc538	\N	idp-confirm-link	cwbi	de7c9cc6-efd9-40a7-9f02-c1727c8664d1	0	10	f	\N	\N
e1e3e9cf-13dd-4006-8c8a-78c09b11adaf	\N	idp-auto-link	cwbi	de7c9cc6-efd9-40a7-9f02-c1727c8664d1	0	20	f	\N	\N
e4cfa6e2-e9c6-45b6-9c3f-871059357b14	\N	idp-create-user-if-unique	cwbi	027cf582-0a1f-429d-93e0-549cc194c933	2	10	f	\N	1a7024fc-ed33-41d1-a2bf-463b2e5bc744
e1c81ddb-0de0-4d27-82af-40c503b4bbaa	\N	\N	cwbi	027cf582-0a1f-429d-93e0-549cc194c933	2	20	t	de7c9cc6-efd9-40a7-9f02-c1727c8664d1	\N
a71a3834-8512-4722-9c73-1c563f883876	\N	idp-username-password-form	cwbi	d0b7ebaa-ad7e-45b4-a126-006c6cd0be1b	0	10	f	\N	\N
e1d80ecb-c38b-4719-a6f5-549798b5aafc	\N	\N	cwbi	d0b7ebaa-ad7e-45b4-a126-006c6cd0be1b	1	20	t	6cd4395e-ef43-47d5-86a1-c7e97cbde27d	\N
7d67a6fc-7193-4ee9-a22c-815127939ce1	\N	direct-grant-validate-username	cwbi	a513d2fe-0ca7-4ef9-8246-33c176e80586	0	10	f	\N	\N
29c4ef92-d263-4a50-90e9-9cd695a168e2	\N	direct-grant-validate-password	cwbi	a513d2fe-0ca7-4ef9-8246-33c176e80586	0	20	f	\N	\N
63700e1e-1a5d-4550-b12a-278f0e288224	\N	\N	cwbi	a513d2fe-0ca7-4ef9-8246-33c176e80586	1	30	t	c01b6492-7ac5-4f85-a085-e519e5bc0dea	\N
6e76258a-c72c-40f5-8481-23eba399e6d7	\N	docker-http-basic-authenticator	cwbi	bb69c43d-7961-4f3e-abe3-793f283e0c68	0	10	f	\N	\N
d9348e97-5b10-45d6-aba8-133a98c65660	\N	idp-review-profile	cwbi	b9733c83-9bce-45c5-a160-35e54fb99e92	0	10	f	\N	61e29fc3-c1dd-4f3c-aa03-455676504530
48df7993-f5f2-4c52-b62e-5abd8bf4d739	\N	\N	cwbi	b9733c83-9bce-45c5-a160-35e54fb99e92	0	20	t	dbcabf07-be0e-457c-9fa7-7d091c42e2ce	\N
93044723-c9bd-474f-8feb-88cc3a3ae7fd	\N	auth-username-password-form	cwbi	ff1f4e39-45e6-4ab5-b43d-74b527959e3d	0	10	f	\N	\N
6a186af0-e765-4619-a963-21f1a01c6cda	\N	\N	cwbi	ff1f4e39-45e6-4ab5-b43d-74b527959e3d	1	20	t	d5ad35fd-d118-4869-bfc2-83702ac3076a	\N
6d3832a8-f854-4e79-8ccd-235d0681a5b4	\N	no-cookie-redirect	cwbi	7031c2b3-f4c0-4cf9-9975-465dcbcf3b3c	0	10	f	\N	\N
656b60f3-c5ad-4f31-9221-c120b5a80ac6	\N	\N	cwbi	7031c2b3-f4c0-4cf9-9975-465dcbcf3b3c	0	20	t	d741433a-cee2-4ffe-b6c5-f6e389a439c1	\N
1603938b-9bd1-444e-a49b-ecf4d6008ac9	\N	registration-page-form	cwbi	ee67ed51-4cbe-44b9-addf-8d1531025bfa	0	10	t	0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	\N
800759ba-62e0-4753-9d3f-7e34886da8d6	\N	registration-user-creation	cwbi	0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	0	20	f	\N	\N
b8df349d-9b65-480f-bc5f-2200cc81c8ab	\N	registration-profile-action	cwbi	0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	0	40	f	\N	\N
6563202e-e919-4a5e-98e7-465906c3b7bf	\N	registration-password-action	cwbi	0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	0	50	f	\N	\N
c492d009-5800-42c7-9363-5b99bf535e7e	\N	registration-recaptcha-action	cwbi	0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	3	60	f	\N	\N
f7a64b97-0568-4ef0-b423-6eb9ea8ea789	\N	reset-credentials-choose-user	cwbi	18c46dcb-0468-4180-b925-a529175a7b03	0	10	f	\N	\N
a3ca148f-31bc-4c14-bb67-845fc7841ec2	\N	reset-credential-email	cwbi	18c46dcb-0468-4180-b925-a529175a7b03	0	20	f	\N	\N
4baecb68-d3d9-43f6-b76c-e18887b1d70b	\N	reset-password	cwbi	18c46dcb-0468-4180-b925-a529175a7b03	0	30	f	\N	\N
6afa414f-0a40-401c-b9c8-00f3e3fe4134	\N	\N	cwbi	18c46dcb-0468-4180-b925-a529175a7b03	1	40	t	de71faee-a3f7-495a-a7ff-8fc8ac3a6ded	\N
8376857c-2b2d-4a50-a5fb-977599b7b2e0	\N	http-basic-authenticator	cwbi	cadcc071-78b6-4305-9fa6-dc212f15bfc2	0	10	f	\N	\N
\.


--
-- Data for Name: authentication_flow; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.authentication_flow (id, alias, description, realm_id, provider_id, top_level, built_in) FROM stdin;
ee2d0274-f49c-44f6-a979-7e7f86176265	browser	browser based authentication	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
9ff67418-36f6-4e75-8b00-07924e64fb79	forms	Username, password, otp and other auth forms.	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
c6451566-012d-4d58-8e98-4e5fcdb9ff07	Browser - Conditional OTP	Flow to determine if the OTP is required for the authentication	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
ea02467e-0b93-485a-8bf2-e401083b818f	direct grant	OpenID Connect Resource Owner Grant	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
334b88f4-3255-4a09-96a5-727c8f8fe2bb	Direct Grant - Conditional OTP	Flow to determine if the OTP is required for the authentication	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
6cda22fb-db28-4a17-89cf-fca653249993	registration	registration flow	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
a41c6a42-8ec0-42d9-890b-ab31dba5888c	registration form	registration form	98749fe9-5c8f-4d46-b973-16664c916f0f	form-flow	f	t
602d6901-12a3-4e29-8bf1-98757e0a2dac	reset credentials	Reset credentials for a user if they forgot their password or something	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
3580988e-c034-45c0-b6a3-25fca9d926a5	Reset - Conditional OTP	Flow to determine if the OTP should be reset or not. Set to REQUIRED to force.	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
e58dc8c5-51b0-4ce5-a7be-e4270afd426e	clients	Base authentication for clients	98749fe9-5c8f-4d46-b973-16664c916f0f	client-flow	t	t
3831e5bd-c4ad-4bbd-b4a0-0403d6bf04d5	first broker login	Actions taken after first broker login with identity provider account, which is not yet linked to any Keycloak account	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
5f8084d8-c868-4ab6-bc1f-7202cb13ea52	User creation or linking	Flow for the existing/non-existing user alternatives	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
46e76e1d-d6bc-4c58-993b-3f5a4b9849b6	Handle Existing Account	Handle what to do if there is existing account with same email/username like authenticated identity provider	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
86940e2a-1c4e-4dab-9e3b-03e1998c7680	Account verification options	Method with which to verity the existing account	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
ccffa12a-e45b-48c8-93be-03c650658985	Verify Existing Account by Re-authentication	Reauthentication of existing account	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
ea9e03db-9bc2-4310-a29d-0c532df73c7f	First broker login - Conditional OTP	Flow to determine if the OTP is required for the authentication	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
67005fb1-bb0d-4f5b-baa1-39ba801af769	saml ecp	SAML ECP Profile Authentication Flow	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
b7d46802-4ce3-4643-bd2d-a5f250726a36	docker auth	Used by Docker clients to authenticate against the IDP	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
78d2f3e5-b3a6-4310-8d08-892ac5ecb67f	http challenge	An authentication flow based on challenge-response HTTP Authentication Schemes	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	t	t
35737844-3a52-45bd-9a01-0e22a7eb14af	Authentication Options	Authentication options.	98749fe9-5c8f-4d46-b973-16664c916f0f	basic-flow	f	t
aa23b4d3-2163-4e10-b30d-4d65777c07b0	Account verification options	Method with which to verity the existing account	cwbi	basic-flow	f	t
d741433a-cee2-4ffe-b6c5-f6e389a439c1	Authentication Options	Authentication options.	cwbi	basic-flow	f	t
771a70a5-cf3f-42f1-bb13-ece0b626389f	Auto-link User		cwbi	basic-flow	t	f
d5ad35fd-d118-4869-bfc2-83702ac3076a	Browser - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	t
0679fb07-660f-4278-a965-820043ead874	CWBI x509 Direct Grant		cwbi	basic-flow	t	f
c01b6492-7ac5-4f85-a085-e519e5bc0dea	Direct Grant - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	t
3108dc5d-0b1c-427e-83ad-3ae677e1476a	First broker login - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	t
7aa06afe-f1f1-4c23-8b19-c534b99a7fd5	Handle Existing Account	Handle what to do if there is existing account with same email/username like authenticated identity provider	cwbi	basic-flow	f	t
de71faee-a3f7-495a-a7ff-8fc8ac3a6ded	Reset - Conditional OTP	Flow to determine if the OTP should be reset or not. Set to REQUIRED to force.	cwbi	basic-flow	f	t
dbcabf07-be0e-457c-9fa7-7d091c42e2ce	User creation or linking	Flow for the existing/non-existing user alternatives	cwbi	basic-flow	f	t
a1ee61b6-25d3-4038-8e01-6c0e95972f41	Verify Existing Account by Re-authentication	Reauthentication of existing account	cwbi	basic-flow	f	t
fc78af35-3d57-4224-9087-7a4408ab4194	browser	browser based authentication	cwbi	basic-flow	t	t
0b8efdc6-045f-4e42-b086-730f8d817711	browser_conditional_login	browser based authentication	cwbi	basic-flow	t	f
0d25d6f7-a29e-4960-af91-9502967dc934	browser_conditional_login browser_force_login_gov Browser - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	f
e6654661-fa03-4c49-9728-2c9df094c315	browser_conditional_login browser_force_login_gov forms	Username, password, otp and other auth forms.	cwbi	basic-flow	f	f
381d1702-a207-4abe-a834-6de732f058af	browser_force_login_gov	browser based authentication	cwbi	basic-flow	t	f
d86e6e71-d763-41ce-a68d-b83ccbe72f8d	browser_force_login_gov Browser - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	f
7dde2057-fbeb-4dc3-85f4-04acbe9bf9bf	browser_force_login_gov forms	Username, password, otp and other auth forms.	cwbi	basic-flow	f	f
f423039a-26be-4d4e-89b5-bc5d821db555	browser_saml_oidc	browser based authentication	cwbi	basic-flow	t	f
5b280719-b0f6-42ed-9dba-c3e197f873e8	browser_saml_oidc Browser - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	f
fc863779-cc81-43d2-b8df-b2c2e83289f5	browser_saml_oidc forms	Username, password, otp and other auth forms.	cwbi	basic-flow	f	f
18cb7e12-5b9b-4f6a-8231-6b2725fea5c9	browser_x509	browser based authentication	cwbi	basic-flow	t	f
0ac897a7-a668-4207-bb5c-1f1fc6fdadd5	browser_x509 forms	Username, password, otp and other auth forms.	cwbi	basic-flow	f	f
eaf2aa1e-6359-43ca-b403-c169273acae9	clients	Base authentication for clients	cwbi	client-flow	t	t
c8896e58-9c92-4390-ba29-47cd7f4522bf	custom first broker login	Actions taken after first broker login with identity provider account, which is not yet linked to any Keycloak account	cwbi	basic-flow	t	f
6cd4395e-ef43-47d5-86a1-c7e97cbde27d	custom first broker login First broker login - Conditional OTP	Flow to determine if the OTP is required for the authentication	cwbi	basic-flow	f	f
de7c9cc6-efd9-40a7-9f02-c1727c8664d1	custom first broker login Handle Existing Account	Handle what to do if there is existing account with same email/username like authenticated identity provider	cwbi	basic-flow	f	f
027cf582-0a1f-429d-93e0-549cc194c933	custom first broker login User creation or linking	Flow for the existing/non-existing user alternatives	cwbi	basic-flow	f	f
d0b7ebaa-ad7e-45b4-a126-006c6cd0be1b	custom first broker login Verify Existing Account by Re-authentication	Reauthentication of existing account	cwbi	basic-flow	f	f
a513d2fe-0ca7-4ef9-8246-33c176e80586	direct grant	OpenID Connect Resource Owner Grant	cwbi	basic-flow	t	t
bb69c43d-7961-4f3e-abe3-793f283e0c68	docker auth	Used by Docker clients to authenticate against the IDP	cwbi	basic-flow	t	t
b9733c83-9bce-45c5-a160-35e54fb99e92	first broker login	Actions taken after first broker login with identity provider account, which is not yet linked to any Keycloak account	cwbi	basic-flow	t	t
ff1f4e39-45e6-4ab5-b43d-74b527959e3d	forms	Username, password, otp and other auth forms.	cwbi	basic-flow	f	t
7031c2b3-f4c0-4cf9-9975-465dcbcf3b3c	http challenge	An authentication flow based on challenge-response HTTP Authentication Schemes	cwbi	basic-flow	t	t
ee67ed51-4cbe-44b9-addf-8d1531025bfa	registration	registration flow	cwbi	basic-flow	t	t
0a5d53e3-306a-4eb3-8c09-b1cb013c2a1a	registration form	registration form	cwbi	form-flow	f	t
18c46dcb-0468-4180-b925-a529175a7b03	reset credentials	Reset credentials for a user if they forgot their password or something	cwbi	basic-flow	t	t
cadcc071-78b6-4305-9fa6-dc212f15bfc2	saml ecp	SAML ECP Profile Authentication Flow	cwbi	basic-flow	t	t
\.


--
-- Data for Name: authenticator_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.authenticator_config (id, alias, realm_id) FROM stdin;
b572197f-4f16-4fe3-bf88-946e2c224bd1	review profile config	98749fe9-5c8f-4d46-b973-16664c916f0f
6f9ec3b4-f0b4-46fc-bd49-765e5a6635cc	create unique user config	98749fe9-5c8f-4d46-b973-16664c916f0f
1a7024fc-ed33-41d1-a2bf-463b2e5bc744	create unique user config	cwbi
ac7e58de-53b2-42d1-a708-65ec837458cc	login.gov	cwbi
61e29fc3-c1dd-4f3c-aa03-455676504530	review profile config	cwbi
ceb288ff-55ba-4e5e-b3db-0d8af4a44671	saml	cwbi
\.


--
-- Data for Name: authenticator_config_entry; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.authenticator_config_entry (authenticator_id, value, name) FROM stdin;
6f9ec3b4-f0b4-46fc-bd49-765e5a6635cc	false	require.password.update.after.registration
b572197f-4f16-4fe3-bf88-946e2c224bd1	missing	update.profile.on.first.login
1a7024fc-ed33-41d1-a2bf-463b2e5bc744	false	require.password.update.after.registration
61e29fc3-c1dd-4f3c-aa03-455676504530	missing	update.profile.on.first.login
ac7e58de-53b2-42d1-a708-65ec837458cc	login.gov	defaultProvider
ceb288ff-55ba-4e5e-b3db-0d8af4a44671	adfs-idp-alias	defaultProvider
\.


--
-- Data for Name: broker_link; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.broker_link (identity_provider, storage_provider_id, realm_id, broker_user_id, broker_username, token, user_id) FROM stdin;
\.


--
-- Data for Name: client; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client (id, enabled, full_scope_allowed, client_id, not_before, public_client, secret, base_url, bearer_only, management_url, surrogate_auth_required, realm_id, protocol, node_rereg_timeout, frontchannel_logout, consent_required, name, service_accounts_enabled, client_authenticator_type, root_url, description, registration_token, standard_flow_enabled, implicit_flow_enabled, direct_access_grants_enabled, always_display_in_console) FROM stdin;
ebe2670b-ba08-442e-9983-2807d8e8dbba	t	f	master-realm	0	f	\N	\N	t	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	0	f	f	master Realm	f	client-secret	\N	\N	\N	t	f	f	f
f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	f	account	0	t	\N	/realms/master/account/	f	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	openid-connect	0	f	f	${client_account}	f	client-secret	${authBaseUrl}	\N	\N	t	f	f	f
ff59471c-d8be-4e1b-a846-c01bc708fbe4	t	f	account-console	0	t	\N	/realms/master/account/	f	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	openid-connect	0	f	f	${client_account-console}	f	client-secret	${authBaseUrl}	\N	\N	t	f	f	f
84727cee-1c0d-426e-9f90-e5d126ef055b	t	f	broker	0	f	\N	\N	t	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	openid-connect	0	f	f	${client_broker}	f	client-secret	\N	\N	\N	t	f	f	f
fca2fb0d-1434-4ba2-bd0a-699e623e79be	t	f	security-admin-console	0	t	\N	/admin/master/console/	f	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	openid-connect	0	f	f	${client_security-admin-console}	f	client-secret	${authAdminUrl}	\N	\N	t	f	f	f
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	t	f	admin-cli	0	t	\N	\N	f	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	openid-connect	0	f	f	${client_admin-cli}	f	client-secret	\N	\N	\N	f	f	t	f
bc580584-cd30-4019-9ee4-0c162f4f9802	t	f	cwbi-realm	0	f	\N	\N	t	\N	f	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	0	f	f	cwbi Realm	f	client-secret	\N	\N	\N	t	f	f	f
46da91ba-9a49-40ae-a8b3-a9dca846129d	t	f	account	0	t	\N	/realms/cwbi/account/	f	\N	f	cwbi	openid-connect	0	f	f	${client_account}	f	client-secret	${authBaseUrl}	\N	\N	t	f	f	f
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	t	f	account-console	0	t	\N	/realms/cwbi/account/	f	\N	f	cwbi	openid-connect	0	f	f	${client_account-console}	f	client-secret	${authBaseUrl}	\N	\N	t	f	f	f
b07385e4-89f2-4207-8190-de6ff920c864	t	f	admin-cli	0	t	\N	\N	f	\N	f	cwbi	openid-connect	0	f	f	${client_admin-cli}	f	client-secret	\N	\N	\N	f	f	t	f
17db6de4-2231-432e-99a8-1432ac240dae	t	f	broker	0	f	\N	\N	t	\N	f	cwbi	openid-connect	0	f	f	${client_broker}	f	client-secret	\N	\N	\N	t	f	f	f
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	t	t	midas	0	t	\N		f		f	cwbi	openid-connect	-1	t	f	midas	f	client-secret	http://localhost:3000	Monitoring Instrumentation Data Acquisition Systems	\N	t	f	t	t
5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	f	realm-management	0	f	\N	\N	t	\N	f	cwbi	openid-connect	0	f	f	${client_realm-management}	f	client-secret	\N	\N	\N	t	f	f	f
38f6b360-c12e-4e69-8b63-43ab4910e344	t	f	security-admin-console	0	t	\N	/admin/cwbi/console/	f	\N	f	cwbi	openid-connect	0	f	f	${client_security-admin-console}	f	client-secret	${authAdminUrl}	\N	\N	t	f	f	f
\.


--
-- Data for Name: client_attributes; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_attributes (client_id, name, value) FROM stdin;
f656ab57-d2fe-4f82-a765-0357d7ef4a46	post.logout.redirect.uris	+
ff59471c-d8be-4e1b-a846-c01bc708fbe4	post.logout.redirect.uris	+
ff59471c-d8be-4e1b-a846-c01bc708fbe4	pkce.code.challenge.method	S256
fca2fb0d-1434-4ba2-bd0a-699e623e79be	post.logout.redirect.uris	+
fca2fb0d-1434-4ba2-bd0a-699e623e79be	pkce.code.challenge.method	S256
46da91ba-9a49-40ae-a8b3-a9dca846129d	post.logout.redirect.uris	+
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	post.logout.redirect.uris	+
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	pkce.code.challenge.method	S256
b07385e4-89f2-4207-8190-de6ff920c864	post.logout.redirect.uris	+
17db6de4-2231-432e-99a8-1432ac240dae	post.logout.redirect.uris	+
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	oidc.ciba.grant.enabled	false
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	display.on.consent.screen	false
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	oauth2.device.authorization.grant.enabled	false
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	backchannel.logout.session.required	true
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	backchannel.logout.revoke.offline.tokens	false
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	post.logout.redirect.uris	+
5fdc5f46-8594-4e73-a982-9138ff9a0f89	post.logout.redirect.uris	+
38f6b360-c12e-4e69-8b63-43ab4910e344	post.logout.redirect.uris	+
38f6b360-c12e-4e69-8b63-43ab4910e344	pkce.code.challenge.method	S256
\.


--
-- Data for Name: client_auth_flow_bindings; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_auth_flow_bindings (client_id, flow_id, binding_name) FROM stdin;
\.


--
-- Data for Name: client_initial_access; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_initial_access (id, realm_id, "timestamp", expiration, count, remaining_count) FROM stdin;
\.


--
-- Data for Name: client_node_registrations; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_node_registrations (client_id, value, name) FROM stdin;
\.


--
-- Data for Name: client_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_scope (id, name, realm_id, description, protocol) FROM stdin;
d3865983-63c5-4a61-8790-79a5f862056d	offline_access	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect built-in scope: offline_access	openid-connect
03ec7dfe-784f-4953-b4a5-7316c1193a66	role_list	98749fe9-5c8f-4d46-b973-16664c916f0f	SAML role list	saml
76365f1a-cc60-4b37-affe-c5c494cf2f47	profile	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect built-in scope: profile	openid-connect
4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	email	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect built-in scope: email	openid-connect
f1d89373-70ef-400b-a3bb-aafbcfa4326b	address	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect built-in scope: address	openid-connect
bdf9e8fc-7774-4216-ae5b-9c955ec11853	phone	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect built-in scope: phone	openid-connect
d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	roles	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect scope for add user roles to the access token	openid-connect
66485ea3-33cc-45e2-9db8-4b7632fcbcfd	web-origins	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect scope for add allowed web origins to the access token	openid-connect
7bab9854-2b99-4050-b410-81a6a09c7832	microprofile-jwt	98749fe9-5c8f-4d46-b973-16664c916f0f	Microprofile - JWT built-in scope	openid-connect
016e49db-14f9-4d10-be20-4502b2a84a27	acr	98749fe9-5c8f-4d46-b973-16664c916f0f	OpenID Connect scope for add acr (authentication context class reference) to the token	openid-connect
662e8870-a4c7-431a-9b40-56a3c4cbb6ea	role_list	cwbi	SAML role list	saml
b0a33b5f-7c9a-4d59-9602-855dfb2a0b92	x509_presented	cwbi	x509_presented	openid-connect
98992cfc-118d-4c64-976b-7ba01f0976a5	email	cwbi	OpenID Connect built-in scope: email	openid-connect
2b3b3db7-5772-4d81-a35a-742ea21a95e6	microprofile-jwt	cwbi	Microprofile - JWT built-in scope	openid-connect
0542ff9c-210f-4eb3-b62e-fa7272032823	address	cwbi	OpenID Connect built-in scope: address	openid-connect
c983959d-26d1-413e-9343-f5bad7dabc51	acr	cwbi	OpenID Connect scope for add acr (authentication context class reference) to the token	openid-connect
268fb6b2-58c0-44f2-ae22-a8baf2236e18	profile	cwbi	OpenID Connect built-in scope: profile	openid-connect
47da9728-25a4-4462-8897-67d5b9e56d92	groups	cwbi	group membership	openid-connect
7d5c5b91-9cc3-467b-aced-a25db29c2576	roles	cwbi	OpenID Connect scope for add user roles to the access token	openid-connect
cdc2e818-4856-4688-b54a-03a5d08e6a1d	web-origins	cwbi	OpenID Connect scope for add allowed web origins to the access token	openid-connect
8af9e1dc-6902-4244-a87e-b40749f6a92d	offline_access	cwbi	OpenID Connect built-in scope: offline_access	openid-connect
5286fee9-6cda-4a94-aba0-dffa0a5c2e8f	cacUID	cwbi	Unique identifier from CAC/PIV certificate	openid-connect
9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0	preferred_username	cwbi	\N	openid-connect
2648bc85-15fc-4968-b6b5-8b9743c8cfad	phone	cwbi	OpenID Connect built-in scope: phone	openid-connect
17e1b31a-5522-4e03-a85c-e476b919a19a	subjectDN	cwbi	CAC/PIV Designated Name value	openid-connect
\.


--
-- Data for Name: client_scope_attributes; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_scope_attributes (scope_id, value, name) FROM stdin;
d3865983-63c5-4a61-8790-79a5f862056d	true	display.on.consent.screen
d3865983-63c5-4a61-8790-79a5f862056d	${offlineAccessScopeConsentText}	consent.screen.text
03ec7dfe-784f-4953-b4a5-7316c1193a66	true	display.on.consent.screen
03ec7dfe-784f-4953-b4a5-7316c1193a66	${samlRoleListScopeConsentText}	consent.screen.text
76365f1a-cc60-4b37-affe-c5c494cf2f47	true	display.on.consent.screen
76365f1a-cc60-4b37-affe-c5c494cf2f47	${profileScopeConsentText}	consent.screen.text
76365f1a-cc60-4b37-affe-c5c494cf2f47	true	include.in.token.scope
4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	true	display.on.consent.screen
4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	${emailScopeConsentText}	consent.screen.text
4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	true	include.in.token.scope
f1d89373-70ef-400b-a3bb-aafbcfa4326b	true	display.on.consent.screen
f1d89373-70ef-400b-a3bb-aafbcfa4326b	${addressScopeConsentText}	consent.screen.text
f1d89373-70ef-400b-a3bb-aafbcfa4326b	true	include.in.token.scope
bdf9e8fc-7774-4216-ae5b-9c955ec11853	true	display.on.consent.screen
bdf9e8fc-7774-4216-ae5b-9c955ec11853	${phoneScopeConsentText}	consent.screen.text
bdf9e8fc-7774-4216-ae5b-9c955ec11853	true	include.in.token.scope
d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	true	display.on.consent.screen
d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	${rolesScopeConsentText}	consent.screen.text
d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	false	include.in.token.scope
66485ea3-33cc-45e2-9db8-4b7632fcbcfd	false	display.on.consent.screen
66485ea3-33cc-45e2-9db8-4b7632fcbcfd		consent.screen.text
66485ea3-33cc-45e2-9db8-4b7632fcbcfd	false	include.in.token.scope
7bab9854-2b99-4050-b410-81a6a09c7832	false	display.on.consent.screen
7bab9854-2b99-4050-b410-81a6a09c7832	true	include.in.token.scope
016e49db-14f9-4d10-be20-4502b2a84a27	false	display.on.consent.screen
016e49db-14f9-4d10-be20-4502b2a84a27	false	include.in.token.scope
662e8870-a4c7-431a-9b40-56a3c4cbb6ea	${samlRoleListScopeConsentText}	consent.screen.text
662e8870-a4c7-431a-9b40-56a3c4cbb6ea	true	display.on.consent.screen
b0a33b5f-7c9a-4d59-9602-855dfb2a0b92	true	include.in.token.scope
b0a33b5f-7c9a-4d59-9602-855dfb2a0b92	false	display.on.consent.screen
98992cfc-118d-4c64-976b-7ba01f0976a5	true	include.in.token.scope
98992cfc-118d-4c64-976b-7ba01f0976a5	true	display.on.consent.screen
98992cfc-118d-4c64-976b-7ba01f0976a5	${emailScopeConsentText}	consent.screen.text
2b3b3db7-5772-4d81-a35a-742ea21a95e6	true	include.in.token.scope
2b3b3db7-5772-4d81-a35a-742ea21a95e6	false	display.on.consent.screen
0542ff9c-210f-4eb3-b62e-fa7272032823	true	include.in.token.scope
0542ff9c-210f-4eb3-b62e-fa7272032823	true	display.on.consent.screen
0542ff9c-210f-4eb3-b62e-fa7272032823	${addressScopeConsentText}	consent.screen.text
c983959d-26d1-413e-9343-f5bad7dabc51	false	include.in.token.scope
c983959d-26d1-413e-9343-f5bad7dabc51	false	display.on.consent.screen
268fb6b2-58c0-44f2-ae22-a8baf2236e18	true	include.in.token.scope
268fb6b2-58c0-44f2-ae22-a8baf2236e18	true	display.on.consent.screen
268fb6b2-58c0-44f2-ae22-a8baf2236e18	${profileScopeConsentText}	consent.screen.text
47da9728-25a4-4462-8897-67d5b9e56d92	true	include.in.token.scope
47da9728-25a4-4462-8897-67d5b9e56d92	false	display.on.consent.screen
7d5c5b91-9cc3-467b-aced-a25db29c2576	false	include.in.token.scope
7d5c5b91-9cc3-467b-aced-a25db29c2576	true	display.on.consent.screen
7d5c5b91-9cc3-467b-aced-a25db29c2576	${rolesScopeConsentText}	consent.screen.text
cdc2e818-4856-4688-b54a-03a5d08e6a1d	false	include.in.token.scope
cdc2e818-4856-4688-b54a-03a5d08e6a1d	false	display.on.consent.screen
cdc2e818-4856-4688-b54a-03a5d08e6a1d		consent.screen.text
8af9e1dc-6902-4244-a87e-b40749f6a92d	${offlineAccessScopeConsentText}	consent.screen.text
8af9e1dc-6902-4244-a87e-b40749f6a92d	true	display.on.consent.screen
5286fee9-6cda-4a94-aba0-dffa0a5c2e8f	true	include.in.token.scope
5286fee9-6cda-4a94-aba0-dffa0a5c2e8f	false	display.on.consent.screen
9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0	true	include.in.token.scope
9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0	false	display.on.consent.screen
2648bc85-15fc-4968-b6b5-8b9743c8cfad	true	include.in.token.scope
2648bc85-15fc-4968-b6b5-8b9743c8cfad	true	display.on.consent.screen
2648bc85-15fc-4968-b6b5-8b9743c8cfad	${phoneScopeConsentText}	consent.screen.text
17e1b31a-5522-4e03-a85c-e476b919a19a	true	include.in.token.scope
17e1b31a-5522-4e03-a85c-e476b919a19a	false	display.on.consent.screen
\.


--
-- Data for Name: client_scope_client; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_scope_client (client_id, scope_id, default_scope) FROM stdin;
f656ab57-d2fe-4f82-a765-0357d7ef4a46	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
f656ab57-d2fe-4f82-a765-0357d7ef4a46	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
f656ab57-d2fe-4f82-a765-0357d7ef4a46	016e49db-14f9-4d10-be20-4502b2a84a27	t
f656ab57-d2fe-4f82-a765-0357d7ef4a46	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
f656ab57-d2fe-4f82-a765-0357d7ef4a46	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
f656ab57-d2fe-4f82-a765-0357d7ef4a46	d3865983-63c5-4a61-8790-79a5f862056d	f
f656ab57-d2fe-4f82-a765-0357d7ef4a46	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
f656ab57-d2fe-4f82-a765-0357d7ef4a46	7bab9854-2b99-4050-b410-81a6a09c7832	f
f656ab57-d2fe-4f82-a765-0357d7ef4a46	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
ff59471c-d8be-4e1b-a846-c01bc708fbe4	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
ff59471c-d8be-4e1b-a846-c01bc708fbe4	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
ff59471c-d8be-4e1b-a846-c01bc708fbe4	016e49db-14f9-4d10-be20-4502b2a84a27	t
ff59471c-d8be-4e1b-a846-c01bc708fbe4	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
ff59471c-d8be-4e1b-a846-c01bc708fbe4	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
ff59471c-d8be-4e1b-a846-c01bc708fbe4	d3865983-63c5-4a61-8790-79a5f862056d	f
ff59471c-d8be-4e1b-a846-c01bc708fbe4	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
ff59471c-d8be-4e1b-a846-c01bc708fbe4	7bab9854-2b99-4050-b410-81a6a09c7832	f
ff59471c-d8be-4e1b-a846-c01bc708fbe4	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	016e49db-14f9-4d10-be20-4502b2a84a27	t
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	d3865983-63c5-4a61-8790-79a5f862056d	f
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	7bab9854-2b99-4050-b410-81a6a09c7832	f
d8cbd378-a3de-4ce9-99e1-acbfe4edc35a	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
84727cee-1c0d-426e-9f90-e5d126ef055b	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
84727cee-1c0d-426e-9f90-e5d126ef055b	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
84727cee-1c0d-426e-9f90-e5d126ef055b	016e49db-14f9-4d10-be20-4502b2a84a27	t
84727cee-1c0d-426e-9f90-e5d126ef055b	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
84727cee-1c0d-426e-9f90-e5d126ef055b	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
84727cee-1c0d-426e-9f90-e5d126ef055b	d3865983-63c5-4a61-8790-79a5f862056d	f
84727cee-1c0d-426e-9f90-e5d126ef055b	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
84727cee-1c0d-426e-9f90-e5d126ef055b	7bab9854-2b99-4050-b410-81a6a09c7832	f
84727cee-1c0d-426e-9f90-e5d126ef055b	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
ebe2670b-ba08-442e-9983-2807d8e8dbba	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
ebe2670b-ba08-442e-9983-2807d8e8dbba	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
ebe2670b-ba08-442e-9983-2807d8e8dbba	016e49db-14f9-4d10-be20-4502b2a84a27	t
ebe2670b-ba08-442e-9983-2807d8e8dbba	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
ebe2670b-ba08-442e-9983-2807d8e8dbba	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
ebe2670b-ba08-442e-9983-2807d8e8dbba	d3865983-63c5-4a61-8790-79a5f862056d	f
ebe2670b-ba08-442e-9983-2807d8e8dbba	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
ebe2670b-ba08-442e-9983-2807d8e8dbba	7bab9854-2b99-4050-b410-81a6a09c7832	f
ebe2670b-ba08-442e-9983-2807d8e8dbba	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
fca2fb0d-1434-4ba2-bd0a-699e623e79be	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
fca2fb0d-1434-4ba2-bd0a-699e623e79be	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
fca2fb0d-1434-4ba2-bd0a-699e623e79be	016e49db-14f9-4d10-be20-4502b2a84a27	t
fca2fb0d-1434-4ba2-bd0a-699e623e79be	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
fca2fb0d-1434-4ba2-bd0a-699e623e79be	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
fca2fb0d-1434-4ba2-bd0a-699e623e79be	d3865983-63c5-4a61-8790-79a5f862056d	f
fca2fb0d-1434-4ba2-bd0a-699e623e79be	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
fca2fb0d-1434-4ba2-bd0a-699e623e79be	7bab9854-2b99-4050-b410-81a6a09c7832	f
fca2fb0d-1434-4ba2-bd0a-699e623e79be	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	cdc2e818-4856-4688-b54a-03a5d08e6a1d	t
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	c983959d-26d1-413e-9343-f5bad7dabc51	t
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	7d5c5b91-9cc3-467b-aced-a25db29c2576	t
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	268fb6b2-58c0-44f2-ae22-a8baf2236e18	t
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	98992cfc-118d-4c64-976b-7ba01f0976a5	t
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	0542ff9c-210f-4eb3-b62e-fa7272032823	f
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	2648bc85-15fc-4968-b6b5-8b9743c8cfad	f
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	8af9e1dc-6902-4244-a87e-b40749f6a92d	f
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	2b3b3db7-5772-4d81-a35a-742ea21a95e6	f
\.


--
-- Data for Name: client_scope_role_mapping; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_scope_role_mapping (scope_id, role_id) FROM stdin;
d3865983-63c5-4a61-8790-79a5f862056d	aafc97c2-f3e5-4566-b148-91289a4ed570
8af9e1dc-6902-4244-a87e-b40749f6a92d	51193cbf-31d0-4955-8374-b2bee6bff6c4
\.


--
-- Data for Name: client_session; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_session (id, client_id, redirect_uri, state, "timestamp", session_id, auth_method, realm_id, auth_user_id, current_action) FROM stdin;
\.


--
-- Data for Name: client_session_auth_status; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_session_auth_status (authenticator, status, client_session) FROM stdin;
\.


--
-- Data for Name: client_session_note; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_session_note (name, value, client_session) FROM stdin;
\.


--
-- Data for Name: client_session_prot_mapper; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_session_prot_mapper (protocol_mapper_id, client_session) FROM stdin;
\.


--
-- Data for Name: client_session_role; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_session_role (role_id, client_session) FROM stdin;
\.


--
-- Data for Name: client_user_session_note; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.client_user_session_note (name, value, client_session) FROM stdin;
\.


--
-- Data for Name: component; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.component (id, name, parent_id, provider_id, provider_type, realm_id, sub_type) FROM stdin;
89665e53-fba5-48cf-8bf9-75e86ab88753	Trusted Hosts	98749fe9-5c8f-4d46-b973-16664c916f0f	trusted-hosts	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
7f142bef-0878-45ca-91d9-fc70e0df4718	Consent Required	98749fe9-5c8f-4d46-b973-16664c916f0f	consent-required	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
9a161644-9a41-4373-acf2-a9d94b3e9499	Full Scope Disabled	98749fe9-5c8f-4d46-b973-16664c916f0f	scope	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
93883dc8-54c8-4748-8eff-0bc77858e7fc	Max Clients Limit	98749fe9-5c8f-4d46-b973-16664c916f0f	max-clients	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	Allowed Protocol Mapper Types	98749fe9-5c8f-4d46-b973-16664c916f0f	allowed-protocol-mappers	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
eac03604-0ef2-4ee1-b473-15213f1a4f3d	Allowed Client Scopes	98749fe9-5c8f-4d46-b973-16664c916f0f	allowed-client-templates	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	anonymous
a389800e-d1f2-44f2-809b-9e9ae12b39e0	Allowed Protocol Mapper Types	98749fe9-5c8f-4d46-b973-16664c916f0f	allowed-protocol-mappers	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	authenticated
8bdb078d-00ed-4ade-a0b8-a7e162c9881c	Allowed Client Scopes	98749fe9-5c8f-4d46-b973-16664c916f0f	allowed-client-templates	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	authenticated
f32d2ac2-c5c7-453e-bc81-f0a07c7b2a7a	rsa-generated	98749fe9-5c8f-4d46-b973-16664c916f0f	rsa-generated	org.keycloak.keys.KeyProvider	98749fe9-5c8f-4d46-b973-16664c916f0f	\N
95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	rsa-enc-generated	98749fe9-5c8f-4d46-b973-16664c916f0f	rsa-enc-generated	org.keycloak.keys.KeyProvider	98749fe9-5c8f-4d46-b973-16664c916f0f	\N
a32f6c8c-4c4d-4216-ae53-59abfe88daae	hmac-generated	98749fe9-5c8f-4d46-b973-16664c916f0f	hmac-generated	org.keycloak.keys.KeyProvider	98749fe9-5c8f-4d46-b973-16664c916f0f	\N
44a4b42e-59d6-4eb2-a8ca-1c5fbb7600cd	aes-generated	98749fe9-5c8f-4d46-b973-16664c916f0f	aes-generated	org.keycloak.keys.KeyProvider	98749fe9-5c8f-4d46-b973-16664c916f0f	\N
f29e0efc-3c9a-4ea7-81be-0ecd2f96de3f	Full Scope Disabled	cwbi	scope	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
ab83f002-f1b2-43be-ae34-3e5961c109da	Allowed Client Scopes	cwbi	allowed-client-templates	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	authenticated
54e8763e-93f3-4379-932f-b22012546868	Consent Required	cwbi	consent-required	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
9cd291a3-d7da-4228-9a0a-946e93cd525a	Allowed Protocol Mapper Types	cwbi	allowed-protocol-mappers	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	authenticated
c4425a9a-82e2-4402-8120-07a94382fc96	Allowed Protocol Mapper Types	cwbi	allowed-protocol-mappers	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
277e9d5a-396a-4efa-8f01-fdfd85dd6841	Allowed Client Scopes	cwbi	allowed-client-templates	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
1537b32f-40b9-4f93-b8db-507d22cdd27a	Trusted Hosts	cwbi	trusted-hosts	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
fd71d89d-9f31-4dbb-b37b-d7826e62004e	Max Clients Limit	cwbi	max-clients	org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy	cwbi	anonymous
9d5b10df-5236-446e-af9b-fb39dc4dc09f	hmac-generated	cwbi	hmac-generated	org.keycloak.keys.KeyProvider	cwbi	\N
cc0c3f19-6af7-4b58-9dca-fb0fb2e0db94	fallback-RS512	cwbi	rsa-generated	org.keycloak.keys.KeyProvider	cwbi	\N
f0a0b63f-84cf-4b56-a6e8-4fd9bc9415b0	rsa-generated	cwbi	rsa-generated	org.keycloak.keys.KeyProvider	cwbi	\N
1bbb3f1c-1cbc-4211-90ce-80d3bdd41320	aes-generated	cwbi	aes-generated	org.keycloak.keys.KeyProvider	cwbi	\N
b675b797-9c59-47da-b378-65e9ab174073	\N	cwbi	declarative-user-profile	org.keycloak.userprofile.UserProfileProvider	cwbi	\N
\.


--
-- Data for Name: component_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.component_config (id, component_id, name, value) FROM stdin;
acc5a2e2-dd42-4173-9d8b-82e439315007	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	oidc-sha256-pairwise-sub-mapper
1cd58b0c-170c-48da-acd1-62c0dd51d3e7	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	saml-user-property-mapper
7ae0b62d-c4ef-44d2-9d91-6f21b1d73959	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	oidc-usermodel-attribute-mapper
819480d0-a1d8-4633-b23d-0d13140cb870	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	oidc-full-name-mapper
fd86595a-9a0a-4c9e-8da1-7230bf074014	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	saml-user-attribute-mapper
6e00ad48-9ae6-47e6-a344-827f09ee1f6f	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	saml-role-list-mapper
2d064518-9221-4711-a739-8fa8650097ff	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	oidc-address-mapper
8ab566d2-3b0b-4ac4-8daf-b22d9d9e00b1	9a6bcf84-0b4b-44f1-93a7-16e3d3b8f683	allowed-protocol-mapper-types	oidc-usermodel-property-mapper
6670918f-7f55-4cd7-9626-7b222495d052	eac03604-0ef2-4ee1-b473-15213f1a4f3d	allow-default-scopes	true
3b0311e7-b740-4cff-8f1f-c57cd6a6740d	8bdb078d-00ed-4ade-a0b8-a7e162c9881c	allow-default-scopes	true
afe10e73-f33e-4753-866c-8ea5bdf7f093	89665e53-fba5-48cf-8bf9-75e86ab88753	client-uris-must-match	true
9d12d508-2c14-4fc1-905d-c79e7f483a28	89665e53-fba5-48cf-8bf9-75e86ab88753	host-sending-registration-request-must-match	true
bc91f9f1-68d0-4f81-88e7-7ca70bfbdd1f	93883dc8-54c8-4748-8eff-0bc77858e7fc	max-clients	200
3fdb5157-78d4-44d0-8cf6-dadb7f94fc26	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	saml-user-property-mapper
1a9d3a1c-08b0-4249-9b99-21d967d5aa57	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	oidc-usermodel-property-mapper
447f05e8-d18b-42d5-bd98-0a6bdc3efa86	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	saml-user-attribute-mapper
007e892d-6abf-4c07-ae81-8bfa71235ed7	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	oidc-sha256-pairwise-sub-mapper
31f76baa-6c5a-48d7-b489-52e537f63006	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	oidc-full-name-mapper
1ddf33b5-6df4-4694-bf0d-9ccc9f0965ed	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	oidc-address-mapper
1719eb91-3ba9-480e-ac54-ecbe71467511	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	oidc-usermodel-attribute-mapper
313ea0b7-6fa2-4345-bdb1-7086455ffea1	a389800e-d1f2-44f2-809b-9e9ae12b39e0	allowed-protocol-mapper-types	saml-role-list-mapper
b7dc5fbb-0ebd-4de9-b489-197c080ee903	f32d2ac2-c5c7-453e-bc81-f0a07c7b2a7a	certificate	MIICmzCCAYMCBgGQT9IzejANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjQwNjI1MTQzMzE0WhcNMzQwNjI1MTQzNDU0WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCo1MVap0Zp5flzIfEz06VEPE0tU53TwPYtHcfQOLq6P7W6hsgjrM1zptJ917OTgAQG3T2tl3s4v62YSnQN+ddaArskf5j03UYhXPLYQ3vxxu4ZgYudnbOkjZVyM7D+ty3qfyc/jfyCX8erLx36eY6TY4bPgrIo8FjUcxOirjAyZYWr+J4EQBHA8fOSlaQwQhkum1RgssIRAelblfs9Dy1zgE8CzVJPyk650n5UpLT7cTScPVhbkUVgHq8dDyCL8dLhCZ3WS5IjvGHN1U4CDl/dtUZQYqdGv51A+MLqBqQaWn1Tj1EUjPdToJSpzvYi/p/PIh3h1vQgpWNeWDWhwE5xAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAGZ8XneS6H8pUZVa3kdGz/xpFhapQKHNyzhHQvuhVfBG9yzJILBjzgw0cOfk944JXKb09xuq21NfvNDlrMOZPiquszCaRq9ap5Y+0cMaBxwYrCNF7wye4kzpaCcl1vSQVNpxyDgZOBmHDtPFSO0rZgrVfFK/DfBFv8NxDWtnm70SmZxszFyVqznkegeHIkCqIldiMegcn+34u2hXuQVHQ2OwWb9QrdzuEPDBwSxT6ZMrLSUnY16zvcLZ7t9vYHebz4jqgpjj968KQ0BYFGV4kDsDbjMcQZDY2hWeXYCSjPTP+476lfqaRbTIYMXu6MtvqzxTqKGKEhB5WZzhdw4pIbM=
df32d889-7529-4468-947a-29de9d0f5130	f32d2ac2-c5c7-453e-bc81-f0a07c7b2a7a	keyUse	SIG
e1b15a47-e621-4010-895f-13e456532343	f32d2ac2-c5c7-453e-bc81-f0a07c7b2a7a	priority	100
ab83b3bb-796e-4f64-96a6-07bd60fc60f7	f32d2ac2-c5c7-453e-bc81-f0a07c7b2a7a	privateKey	MIIEowIBAAKCAQEAqNTFWqdGaeX5cyHxM9OlRDxNLVOd08D2LR3H0Di6uj+1uobII6zNc6bSfdezk4AEBt09rZd7OL+tmEp0DfnXWgK7JH+Y9N1GIVzy2EN78cbuGYGLnZ2zpI2VcjOw/rct6n8nP438gl/Hqy8d+nmOk2OGz4KyKPBY1HMToq4wMmWFq/ieBEARwPHzkpWkMEIZLptUYLLCEQHpW5X7PQ8tc4BPAs1ST8pOudJ+VKS0+3E0nD1YW5FFYB6vHQ8gi/HS4Qmd1kuSI7xhzdVOAg5f3bVGUGKnRr+dQPjC6gakGlp9U49RFIz3U6CUqc72Iv6fzyId4db0IKVjXlg1ocBOcQIDAQABAoIBABF87K5tukz43eRvpSD1sN5HEsV5rlUDXVyiA5MNdUYamFPoZy3O1f8/TflsEPVb4s7lNuDW2pQvwqcOO0RBV23C76ihsPHYQ83r51nAb8PFE9+/e/tJHRUT92F7ej+AMPjjz+h06C2HB1Mzj7rkwYCB5DJ1esfj0Ye8HdIRkft+Po0do6TKqLxVGqPyU491uJeacpxT+k8dAjrF0EiA1HCt96ixZsZSVKbgVcFfHABunKCTfpV/qqCCe/Fh4GZuq6rpQx21MGaRszJicjkxNFPm1sUPZeUTYYpzAyKIlfMOo+eHMP5xGL2XCF1RHhfDgW/UxPo6cIbUUkRyAf/UAksCgYEA0+KP30E3FqaYBNiiP5MZqFHEeHkxLCzVk6/JOzFq5XXEW/kH/o8j4IfibiKCsDU9MHSc85ZdxaLcVTcl00o9Qu4QuOGT2bv4DPiIoYcMKj32KwynZfOCTjrQl0XgNjwDyrJGY9R40dAvJZwGcl7wGfYNawnCm3jkZRlJ7X/Y/TsCgYEAy/txurq6LtyC9oMB36wjLaF/8kGmjkoEpmFp3sll/CxUU3pCjBWIggXuAMh9t6KZkbun5F18pzwYayMQ9oKVipdRltbNLHg0fgKfXpo3kc3cTpyXOWNU5HWqdU6+4DiiRZBeMeEGboiPvHYERVvgeGkD7Kx4U7R8rSmcXAMhmEMCgYEAo7wsmcV3oJVUXEpb9uzhouStArwEd7KdyObKhmeFx5PeDYS/3MMYYmYfYCRjAW/ivRMgRkwoYpWb1m4rWL/B33rAiV4oNtW+tadJTiliCTKgjFvW8D3gsDta/csNnFt8QqLJKlQCmYBbLqEHilI7EJTHgbOlIyekt0t3iYSGVgkCgYBM/xLE62sP+UiuCHGSnoWA2e9T4JggKaxrqWWvJNKMR6dlA0dPXWlzuw1F2mgqAwR40B7lwqwk7DhDaK8kfdI4yDmR+G7mFFGsJw1FRm0nak94lP84OFh5DDlVom3GcVo/a+lCZrBo1L984gdmrvGiQfGrSsb4wa42JvjQepYZxQKBgByWhtN6flzYxqqdFDkHWDYfkfFe9KDPDbjzyATmpj9SQRlbtkzxwWDl9L5sGlm70Rj31DoGuDIokfstE7PrlzJ4Wm128v1GDzAY6eKvEzKT1DzbTosB3lw6fLkieVttQBswWU/9BoRa5LRJmnWH/RhuTPnG6u5PS60dZkoUpIUf
c10353b1-b05f-422c-b3bc-6a2b5bedc7dc	a32f6c8c-4c4d-4216-ae53-59abfe88daae	priority	100
19e0003e-f9e7-44c4-bc77-286779a91754	a32f6c8c-4c4d-4216-ae53-59abfe88daae	algorithm	HS256
3ef4d416-3e28-4950-8342-1033d2cd9795	a32f6c8c-4c4d-4216-ae53-59abfe88daae	kid	b7372e23-7e14-41f8-9b57-de16fd2e9a46
49d98941-29ff-4ef9-9ca9-8148c3fffd6a	a32f6c8c-4c4d-4216-ae53-59abfe88daae	secret	JWrCWAFmAfZklRO4CXlTnMt4dPjwDlqpodD00aRnqcS2eM5ZlcWJ7ZMt0FCh1EI4vjebpWaC_Ox4XpWpYjkxnQ
ee8f90c8-3638-4f38-881f-52c3bc1715ee	44a4b42e-59d6-4eb2-a8ca-1c5fbb7600cd	secret	jvZQRk05r4cityLWvqNIhQ
3b99131a-f95b-4cae-9775-d08859ade3f6	44a4b42e-59d6-4eb2-a8ca-1c5fbb7600cd	kid	3b772efe-c0ca-4808-acda-0de83287e2a9
63ec6acd-331a-439f-ac31-5234b60f87c9	44a4b42e-59d6-4eb2-a8ca-1c5fbb7600cd	priority	100
276fc1a4-24e7-4e9f-b9c6-3a9a5765a131	95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	algorithm	RSA-OAEP
04436082-20d7-41ce-87be-71acf4f4aa6f	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	saml-user-property-mapper
11920b22-8119-4f74-9c4e-cc9a61638509	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	saml-user-attribute-mapper
3f7be773-2dbb-4f7c-adf5-feeac081e6fb	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	oidc-usermodel-attribute-mapper
8216f113-e87d-493c-8265-03f123ab7ec8	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	oidc-address-mapper
b43c4d7a-dcd7-4a92-a1f4-f787b67e68b1	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	oidc-usermodel-property-mapper
f4a8fed8-0635-40f9-a94d-bbd95a70e5ed	95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	certificate	MIICmzCCAYMCBgGQT9I1wjANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjQwNjI1MTQzMzE0WhcNMzQwNjI1MTQzNDU0WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDhNwAIZBH2MBFFHancnJmBP9Of6aeYRbuGOVJMEWolSXh4mSqkL1cAxbu7+W4oFMyaFp7oVcu88O2ofgl0mJ5ytOQY8I1E3ytrHJKz4ZlPVls6jEEV4CRql2shBfHjgkYrk3kql5LoiUwnkTVlt3c0yj3iNDoH2EYTB5VsE3t/7M1fY8kSEYrJXiA988VRsAJfGDVlBlL3rz6HsX5K9RKmSWGEP+Y1yH9XRxOC7KdjHdQUEadQ1z3mYmrjxCtdGhvoqHEc0lPkSL0WKiG6GHFe6W+w5qN0gk9oKAyOPEEuKqxpSNYHvwZFfiA8HhDB5OWDI2SX7nPtsKsmM7VlUOP5AgMBAAEwDQYJKoZIhvcNAQELBQADggEBAA+zzEbdgc+2+XOe+03FPwAQChQO2mmQolBTGZDrXrWxAnGyCfEzuKIVFUcAWUFv2dqTCReIt9IGMBHuFZsNgBQbB8upsxmcRRDnwEOVIUZQj+w3SuNwG+lEYBocX1kWtSdwgcIUmZ9yZsQhlFlC6oteqLv6RNVBF4RvO0axpQ1s7qYtUG50rb/j4W2CkTBm6IPb/TlIusOEaMpKIzqBoqK2f8GWanaOVvehAhGT+IBKRtNcHYnDvnOKMXg8axzrY9/zZqfmgH/PKplFKLGb98aUZoubKKwBrXwgYsDOJmt8mDEj64J+7NaLedFYXB4qV6URI94jPmlHkC7ISyINIIU=
2b623e4c-3eb9-447d-8512-1b95dbe2e483	95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	keyUse	ENC
689faf33-1622-431a-83d5-07b1488a1e7f	95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	priority	100
e95d3775-4f93-4793-980e-f897ebcdb143	95e412cd-26e6-4b6f-9f51-5cb5d9b3f34b	privateKey	MIIEowIBAAKCAQEA4TcACGQR9jARRR2p3JyZgT/Tn+mnmEW7hjlSTBFqJUl4eJkqpC9XAMW7u/luKBTMmhae6FXLvPDtqH4JdJiecrTkGPCNRN8raxySs+GZT1ZbOoxBFeAkapdrIQXx44JGK5N5KpeS6IlMJ5E1Zbd3NMo94jQ6B9hGEweVbBN7f+zNX2PJEhGKyV4gPfPFUbACXxg1ZQZS968+h7F+SvUSpklhhD/mNch/V0cTguynYx3UFBGnUNc95mJq48QrXRob6KhxHNJT5Ei9FiohuhhxXulvsOajdIJPaCgMjjxBLiqsaUjWB78GRX4gPB4QweTlgyNkl+5z7bCrJjO1ZVDj+QIDAQABAoIBADYCv7JDdX9KFcoyi/sJX1L84J64JWZCSu+srYzqnD8m+IpLiUtowv+/a/9vmThpjIvjouZrNPox+XzwBQp+U3mE4jMM9YQv1TTR3GjhUKgTOLu4yR8a6wDZIWsBBvqd0oA+1M8fHK9Bxg6zJ1AmiKMTYXXvOn+JIX0M04vgvDx4NZ7Ya0V515PqG5R/M2U3IhGmKpYqPPPHmHxi4gtwgANJJjMzJ+TrQ1sTGt5DBoch63ydNw7LvkVekUu/kxjXMaN61UnGRRm7linfAnVmss5zcV92i3NPQDFdOsfsS97Rt8HY/sV4Be4crvyPMOOEXc4GE/tbu6T5ZO+CZSWA4UUCgYEA+nb92pmY75xrKO68M/1PGnmix5NQ6FfNMVI0mlM6PvDDauyYZkD3N2f6jzoFZ6v7xTTqwAIq89FN8Tpy29jjpaRiin8NyWeUAEdH9gVlKn2gp2SpUI1hYisJqNoq4paICX8HCr3gKWdFOyoy+Kl6bVIGWvVFzHWusmwOh/U7kkcCgYEA5jEoDRgLMrlXX8V9DfFrdLXgmc7UGfMGWnref39aFam+jnT4nO5+wtKsVy1UTRXXTUuWMGyldq1kgQ0RajjgJMAFtpbNur+r5RCaZmbGhqlpfeUz0CDLGLfLNCwX9Aih7TZPZp7dfPH8al4o9x3icT8sBcjHEGKawWG6Yc5Ct78CgYBbpiD71YqF2znHD5ykdYN1j57F4p0Pd9lou8jt05iToWwQeyTE/e64Qn8H29ZQejk1j0h7HA/1idg4dgfDdJjQatd4EEfOM+2PMIYfexfqtW+M8SOXizRgRJlTRQm+QhDjUK0V/CbDX7uQi799CB75U7Npoyh4SXO0bB/hFhZHTwKBgHCOBqRoZklWIOf+W8hujHlT2U+7tzA9CZVCUPsMnVLMXhWwEkRBxY+jjYtO6dKLZGwyEmz5Iurlm6gSaLqEtuyhS+nc0RmIURe/R4/cnyQHQELDyNLyxfv/GogXK1sywWKI1Mg709cdR6wHAIbcgPWYywDFDLjxTfRSwdzDCxJnAoGBALCdgVcCGA5+ztZMvNiqw7X8DAruSMB1x8AvFUCHYPAmOLp/7G9q8J7iBbswb6kuW1DMnasKUKiKobvOMKqDfX2lwLV5enfqn3Db2wfGUzXPT0UHmo2/bwoeu3cEmpYQE8lVuEDv9EH5GdzWD4ZqH4baUFT2YmoXdmViO0UnXOf+
968d1c83-df61-44d6-837e-c53f5525bc6f	ab83f002-f1b2-43be-ae34-3e5961c109da	allow-default-scopes	true
540dbec6-d48a-44dd-a6e3-ae9747fbdc7a	9d5b10df-5236-446e-af9b-fb39dc4dc09f	algorithm	HS256
fd423b41-2ad1-4c83-ac39-9d3506e2ea55	9d5b10df-5236-446e-af9b-fb39dc4dc09f	secret	EbRLWIBtg-ezs6DYCU7zDwNVcytR8efJAVad3O61fJXorU0KKUzR5ItgeGKvAc20M72UiinCzDTONjRIArifUw
31d40acc-47ca-4af8-9e4f-6117602a965a	9d5b10df-5236-446e-af9b-fb39dc4dc09f	kid	c2e313f0-4bef-45de-8217-eb4a6260f9ec
e2153a50-bf1f-478a-a5d7-abbe3ab4aed9	9d5b10df-5236-446e-af9b-fb39dc4dc09f	priority	100
3ed09b66-97e3-461a-8ea1-be6d807c3086	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	oidc-full-name-mapper
f57b4fab-12ea-4962-b812-426b56a45b9f	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	saml-role-list-mapper
23aedf03-ba82-4d73-a7ca-34d77172b3f7	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	oidc-address-mapper
2a146de4-4433-4405-8006-b67cf7b71a97	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	oidc-usermodel-attribute-mapper
60efd870-82c6-48e3-8e80-302f0365f188	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	saml-user-attribute-mapper
3e5a8a0f-1c67-4832-a968-e811b245e59c	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	oidc-sha256-pairwise-sub-mapper
bf7c2fad-4f1b-4cf1-92e2-998e65e26e45	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	saml-user-property-mapper
949862ec-2adc-4911-bcd1-1fd2896a443f	9cd291a3-d7da-4228-9a0a-946e93cd525a	allowed-protocol-mapper-types	oidc-usermodel-property-mapper
8fdec449-ee78-4c10-8dcb-e8590cdd35f4	cc0c3f19-6af7-4b58-9dca-fb0fb2e0db94	privateKey	MIIEogIBAAKCAQEA1mRPPHMPowbj3Ech0d1ij8zgp0SBDZaEEbYs2iOMFo2FFhUM+yVrvLBmDt33ifuVGpJ5leQBF+0/4kRMoVzWwpHazMLbj+j1bfxgtPXxv3n9++fiKe5c+1jcdqrCGSsCB9MArInFMLr9YoQzVWPzQa+0B8NZTs490ImnlvUE1qWqATxhck8IIuOTd5A0/SiSwhCdzZwoIXriqF/mtJaS4LNpYSIwzoX5NggHT8xziRsVs6R9RUFD74MLMRwa9PkV8TTxvHBrDb+hYOOhUNhY9u91yC6Py1W38F48rOBpofGi/XUqbI09kfAdYkfXAuhm4d9dAzxswlaKDOfRxRENwQIDAQABAoIBAAbbTm1wgJ+GKONyovJDUlOnCchPuJDmr3KhkO2pFWHjRM5f/fpKSBfQzHLNRo1zLmGbLahNknthaxmhdZHzlirC32yNDtibar4JxE4FT6YAEM9tqx4MMY0YWnSxIWQMrBPz+6GJBnV4hYIRGFMyzyTaqbdV2BVdIzz4KhP54h27P6e1Qb3E57Gz5LbStIFtSV/ZgTyefbe3PlJHSVCg+FckFYk4C2kulokGoFWQNo+GI6MQqe/eZm3BSxybjP93BYXFbXh6ipHL/rUQ7I3i0yinIcBVocj/CNVy1RUKSvHhjje/ogWC79d5/B9VAkRqRNi2pCgYgNlLx+JPQyinSQECgYEA7tw+vIisa/AsrcGWcwx81vOv7UHUCBxo3lHORpet4o0XzZPH+Gmi9ayUuTiVUp42OGSg6zfb53g7kiiCTBGpIp2bOGRR7t5NNNErs9oeo/PKguoj5yDrmLL94+j9N+rMyNGTpGNu+Kg/i+GqpapiBO3U7dCjIkNttKljjzjo3cUCgYEA5caW3ytEJC9YT+IEGV/XXGZ0GVMkqgKCn4LUmzzfxslL5G/YrX8GOVMUXDxHPE72EPKkLlK6k6cw5MVYJXs/tS8WZZHEODxTGGa5rxMjGy7VgpDtFkSyNDA57kblJnUg7of9pKyqUVIEKMq8Q7ycs6pVE4Est7YJ6Sn42A3QC80CgYA3asfva/I89LMY+RITzTDlmhIWBLDR1O2LrwUhoa1JI50DKCc7/h49y3WR54vVaDmCKe6fxAz0DhQcG+PnDC3mhxRtJ/FObysQdvshthhnx3cTmokL5bpjahu3leWx5HrwmJcdi6hCbp6XsJqr5vTo2dkN99rxZx17zdT4dKaqVQKBgHu8QarURdGmnUwHTmXLstHTalno6CmVSHpVneArG+aUqAXQJULo0JU2JBS3cTUM4H2n7Ln1WvwAYNgRXbJxeJE9VjZEFQKjmaveRcf96j3NLrUtDj+bpstr3QZvrx7SnHVXTkPLp7w7CnpdEpz3iPtHkqD6QvZ7VUL2k0blyU9RAoGAew/otrZysGSgffjam2z9UuBm6B0CxT7WYcJMUlLf5eJ2F+aBPYYUJ5Z8yHaMlrt5fFv7jPStpSfV3NMiADBxEAuNRBqRpfCuA8r7VVbrVThKbGSc9GLZEwbbi8S7OHkss8EBk3iTNOMLof2VhG8n9/uibpmSv3C14WZuXZaJIJo=
8a53b222-4d30-4045-aa4f-79d3256e9995	cc0c3f19-6af7-4b58-9dca-fb0fb2e0db94	algorithm	RS512
5cda289c-25e9-4c47-b9aa-0ff1ce96f6e2	cc0c3f19-6af7-4b58-9dca-fb0fb2e0db94	priority	-100
b5546873-42d5-4a4d-ac0b-c991ed83a28d	cc0c3f19-6af7-4b58-9dca-fb0fb2e0db94	certificate	MIIClzCCAX8CBgGQT9I56jANBgkqhkiG9w0BAQsFADAPMQ0wCwYDVQQDDARjd2JpMB4XDTI0MDYyNTE0MzMxNVoXDTM0MDYyNTE0MzQ1NVowDzENMAsGA1UEAwwEY3diaTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANZkTzxzD6MG49xHIdHdYo/M4KdEgQ2WhBG2LNojjBaNhRYVDPsla7ywZg7d94n7lRqSeZXkARftP+JETKFc1sKR2szC24/o9W38YLT18b95/fvn4inuXPtY3HaqwhkrAgfTAKyJxTC6/WKEM1Vj80GvtAfDWU7OPdCJp5b1BNalqgE8YXJPCCLjk3eQNP0oksIQnc2cKCF64qhf5rSWkuCzaWEiMM6F+TYIB0/Mc4kbFbOkfUVBQ++DCzEcGvT5FfE08bxwaw2/oWDjoVDYWPbvdcguj8tVt/BePKzgaaHxov11KmyNPZHwHWJH1wLoZuHfXQM8bMJWigzn0cURDcECAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAPOEgr1d5bpxeIAdBGow6fARO3LyJh2d1DGT1rvOOcVdUVUIXEiupyyBtQNLcSeW8p8lPynRlDf9FxnU+prlCjDlF8ZhjkgKPYjD5Ac4Z4dJuA14sLAbr8I2oQteJ4+OW2ZnLtzhH9doZ9pd0GtC2ooKGaejEAWjRe52NBFs6cOq0e5pWr0rAxUoeLXIYLbb+xjhSedBstDDINyFVyJXFUU3KOco8NSV106GpqoYtIeRA7GkCekXtXBKlUzk4mupgddhwgUX7k8OfSHL/AaSvlqH2CevwNKO5yhobqqIjtOu6LiuKOmWKFe607YS3u4IPCAhgW0fdAiaiviuNHjRCDA==
3fdb449a-6637-4c8c-bbc8-071368d5c8ed	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	oidc-full-name-mapper
75fa0119-3636-4cc1-bd71-f0389e6c63f8	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	oidc-sha256-pairwise-sub-mapper
9d403ed0-43a7-47ca-860f-f3bc1189a631	c4425a9a-82e2-4402-8120-07a94382fc96	allowed-protocol-mapper-types	saml-role-list-mapper
e0a75026-89b8-49a7-95e4-52078a2e3ef7	f0a0b63f-84cf-4b56-a6e8-4fd9bc9415b0	certificate	MIIClzCCAX8CBgGQT9I6HTANBgkqhkiG9w0BAQsFADAPMQ0wCwYDVQQDDARjd2JpMB4XDTI0MDYyNTE0MzMxNVoXDTM0MDYyNTE0MzQ1NVowDzENMAsGA1UEAwwEY3diaTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMswWHJdwA9WyADIGSHJDE0sn/kAQtlNknwjGXRlEB0nWML9RMaExXg2oGENw+3BH+CPdKKF0N/dMTehXKCkoinX4A5GcFUEPO9Z/2LZCpSjb1fgSQQuAUBqjyzTpCIYgGOQLT8OdgToYx9ku2mZH3Cx6ENawMWS3Ij++vKgBwx1N7X8mg2mzLsOpWugjvKRj7rnUMSITewVZbWk7KX8sfFAvrpHwwR8JZ8in5r7Z+PFsFQ0wiRtSUx4mmZzkcyRhgWlHiRXAAD+z7vtmk19gSXeMmEgxQ1dJyd4Tb6A2BSmOtLp3/rh1tqsa2oWJHAe0rvo0GG5HTtQQfsDH1hJSEsCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAMq/CwkbnF2cbDYdq6tAl1Oi1sQSGQ9GaqQ3u2zyNn9TzzeJjrBZT22WEFYHRGQTTxHaB2YIi8CbewxkC9p0u1bjE07vY3ZGIVvWz2BTuqVJRObJxOPo/ELdTLX9m4r/n99raPU+0mm8VIZV+yzxhR43cmMm/a8NWMVlEcrpVk40Qqw45o2/GZf+DPgMVD+cEIcuhALj6gL4uZ53zZ9R2OyWiqdxyCxH0WoIoJSXtr8n3uznHuiDv2sZkkCEM8+0xsXPefJNJXkrqjQt4H/CtL6xNY8Qo3Qsx+624cbBBqLgGg9SVukzSuAA27xM+szdAyhkX/2nIiM/jQF47NNZS7Q==
654d1a8a-0049-4288-b047-603e53d86db2	f0a0b63f-84cf-4b56-a6e8-4fd9bc9415b0	priority	100
5e7d5a2c-b1d6-4925-af17-8e8385beeb4d	f0a0b63f-84cf-4b56-a6e8-4fd9bc9415b0	privateKey	MIIEowIBAAKCAQEAyzBYcl3AD1bIAMgZIckMTSyf+QBC2U2SfCMZdGUQHSdYwv1ExoTFeDagYQ3D7cEf4I90ooXQ390xN6FcoKSiKdfgDkZwVQQ871n/YtkKlKNvV+BJBC4BQGqPLNOkIhiAY5AtPw52BOhjH2S7aZkfcLHoQ1rAxZLciP768qAHDHU3tfyaDabMuw6la6CO8pGPuudQxIhN7BVltaTspfyx8UC+ukfDBHwlnyKfmvtn48WwVDTCJG1JTHiaZnORzJGGBaUeJFcAAP7Pu+2aTX2BJd4yYSDFDV0nJ3hNvoDYFKY60unf+uHW2qxrahYkcB7Su+jQYbkdO1BB+wMfWElISwIDAQABAoIBAAePo90gL7WCasyhX8ciqumod4EkXZKmGahWHGaEqBJaBfuye7EcwovhQIRR4TlmO4d9U39hrVGoCgWZnqbWb1qW/cyxdvmpegeQJHoZsUL3jc6P108EvISNASB+DJtUv/2dOI+iXv3Xq81X8sazhJt+7j02ru6REC83/jQDbUzP/qsdAl7fKK8DNSMmQgbSSfPlm3SpBIAu55tueRrUb+ATXCjpMa342AQTUAfbUN/tHEx1gFTfqD05HYI1yorAU69MmdikVpH4W+maItX32/5uumW8D7ixj23YOimRL1vXKoReUNdphXBUbInSRthMgrsPd62FJ7AWOpbfQWUzVz0CgYEA8mYCxfhLXEdE5JXmDYbT3A5pqcSa98Gwia1sqvRRRCwOsVe2YwVGKn7sr4iswscQGdCxrQepUYSyskbMgE8EvBS+V/trHZ+SWRQmrjJuAxeBveiWGJqCDYEiS+NUiPUU+dV7HepOx5tGoAUdTX0C85VS9EYKgroN1L898YIBkbcCgYEA1pcZTWPEbZrYCZSSXUYg2y9zMJfZ8DGpRR7NzYa8qSANPDxRCs9Z66CTGkG8mfqZoe8Vz4yFvxvSgLUg071AnE70NtIOyyAzh3Wl+3akMnX1V//CC6sph3bINvvCOYjpnbfx9/oLQxSpCFoZLBtAY43upk/XalTRXnIjNuIqLg0CgYEAvqSsmbdmzfTfokii3xCjn/dV80fF+gZEKoRTa0EsiUl9ZM1vjQGg5dtdm9EKz0Zwy1zv4P6dlJehU8WLIX8EYkiOS/RZkrrmB1lp8qeHDrd2Oz6qjj82+hgOOVFaMz1OehAE/MpDm4nsSf67xS4FHD7dN3G+4oIiTqj6tu9g/JECgYAmVgATL0ucORl3PK+ZMjoUbjmp6LbqjjSrQIGLwhJaNHuo0y23PKvXyGv/ONc0uuxPXaML3RLXvWSx0an3qcutIP3H/WbfKvWJsZ7heaSDz0bxXaOQ6hcVOEc0a22bUbZKkt4LawQwC7TW5SGyG3w5TNXhqEnXmSd+M+3OlEDs0QKBgBiMU8OHAzblFJ0I5p41UWa8iQLclPGrTwxlHZWH9/aHWlzKtw+mp3pQrwQ9VizgGTCAeRhMc6K14CiwyxM/Andn6RLppiUlI6ASYTSH6eOqJOzC8djofRIfwNw7hX0q4DWj6FY3ovKSUmcQIWgO+lq9PcRtGWbHoETd01dwYKUe
84f78ac4-5e53-4fd2-95d8-3481ccee7842	1bbb3f1c-1cbc-4211-90ce-80d3bdd41320	priority	100
0ef132ba-b964-4e7a-90ea-f25e690c408b	1bbb3f1c-1cbc-4211-90ce-80d3bdd41320	kid	fe610332-154d-4187-b12c-a56c5b0b8db8
61c5adbb-d5fe-4741-877d-fa5e558404d8	1bbb3f1c-1cbc-4211-90ce-80d3bdd41320	secret	QYA77xkQ00RIQTZngeTOIg
68c521c3-0f29-4626-b141-2505c0fe1ff6	277e9d5a-396a-4efa-8f01-fdfd85dd6841	allow-default-scopes	true
8eafcd91-6164-48da-8c7e-442f2e32fba5	1537b32f-40b9-4f93-b8db-507d22cdd27a	host-sending-registration-request-must-match	true
8c9a5573-d622-4c91-9440-f47e3546523b	1537b32f-40b9-4f93-b8db-507d22cdd27a	client-uris-must-match	true
5f2140a9-f069-4f7c-834b-72559fd3d5b5	fd71d89d-9f31-4dbb-b37b-d7826e62004e	max-clients	200
\.


--
-- Data for Name: composite_role; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.composite_role (composite, child_role) FROM stdin;
8eb26a12-1853-44b8-bf59-4da9e0cb683c	7d9ddb34-3373-42d3-b09a-135a12a3d915
8eb26a12-1853-44b8-bf59-4da9e0cb683c	062a0cf2-00b3-45a3-b488-da81142c49fd
8eb26a12-1853-44b8-bf59-4da9e0cb683c	d594560c-cd04-4b97-9180-a5bd509c2697
8eb26a12-1853-44b8-bf59-4da9e0cb683c	74e9bad0-3106-45f8-80b3-d1a51355b046
8eb26a12-1853-44b8-bf59-4da9e0cb683c	c57caf0d-3402-464e-886f-7cc91131b4be
8eb26a12-1853-44b8-bf59-4da9e0cb683c	956b7000-9d33-4356-b2fe-5fb9eb8ace5c
8eb26a12-1853-44b8-bf59-4da9e0cb683c	63c6007d-8f4d-42b1-bae6-f2d57280ae4e
8eb26a12-1853-44b8-bf59-4da9e0cb683c	d6b070b9-5c18-4609-afed-966e7ab9e3d3
8eb26a12-1853-44b8-bf59-4da9e0cb683c	bd54f0dc-57f3-4648-b106-6c09ce9d9ef0
8eb26a12-1853-44b8-bf59-4da9e0cb683c	229c3773-8383-4e57-a4a2-da2a57d8e2be
8eb26a12-1853-44b8-bf59-4da9e0cb683c	18e3c446-5a2a-4d6b-9ff8-fcda1b265a6a
8eb26a12-1853-44b8-bf59-4da9e0cb683c	5d9f3e23-f626-46d5-b7a3-ea54ef829943
8eb26a12-1853-44b8-bf59-4da9e0cb683c	7fe30e28-ace6-492c-82e9-5abe06edc752
8eb26a12-1853-44b8-bf59-4da9e0cb683c	1ed30b08-d79b-492c-a49a-ceb1c2065cc7
8eb26a12-1853-44b8-bf59-4da9e0cb683c	a8088132-ed0e-4d2c-924d-76971d514b72
8eb26a12-1853-44b8-bf59-4da9e0cb683c	7891a41f-6a7c-4dcb-86a7-954be015b7ad
8eb26a12-1853-44b8-bf59-4da9e0cb683c	eb54652f-08d4-49ce-86ce-10626d6699c6
8eb26a12-1853-44b8-bf59-4da9e0cb683c	1205b878-4df1-4bb3-a644-e35ee220958a
74e9bad0-3106-45f8-80b3-d1a51355b046	a8088132-ed0e-4d2c-924d-76971d514b72
74e9bad0-3106-45f8-80b3-d1a51355b046	1205b878-4df1-4bb3-a644-e35ee220958a
c57caf0d-3402-464e-886f-7cc91131b4be	7891a41f-6a7c-4dcb-86a7-954be015b7ad
db5807fc-5e23-4f38-8243-bfdc9673bd9e	f0c7cd6e-d92c-47d8-aba6-fe2c7bc0d4c0
db5807fc-5e23-4f38-8243-bfdc9673bd9e	98e010f7-b9db-4a9b-872a-8eee15766444
98e010f7-b9db-4a9b-872a-8eee15766444	4426226e-0666-44dd-ad62-bbe162df3720
50bf3d94-3f5c-41ae-9007-640324a3c6f7	0c394c74-e2a3-41dc-a6d4-03b770b1cd0a
8eb26a12-1853-44b8-bf59-4da9e0cb683c	0849c1f4-5f15-4b30-b6a7-8f65aa40ee4e
db5807fc-5e23-4f38-8243-bfdc9673bd9e	aafc97c2-f3e5-4566-b148-91289a4ed570
db5807fc-5e23-4f38-8243-bfdc9673bd9e	c1638b4b-5d98-42ad-bb3f-85617b7b2bfa
8eb26a12-1853-44b8-bf59-4da9e0cb683c	4aa490af-405e-4ebb-b8ef-1cd3fbfa9c84
8eb26a12-1853-44b8-bf59-4da9e0cb683c	a0edf521-eb53-4487-94e6-b9036a5e208c
8eb26a12-1853-44b8-bf59-4da9e0cb683c	818ed298-e735-4251-875a-0778a78aaf81
8eb26a12-1853-44b8-bf59-4da9e0cb683c	e46df155-c26e-4253-9bce-998723e7ee41
8eb26a12-1853-44b8-bf59-4da9e0cb683c	85cb9aca-73a3-4c7d-b090-e8b9f9f6a627
8eb26a12-1853-44b8-bf59-4da9e0cb683c	06728afe-43b1-4807-9534-c1568181147c
8eb26a12-1853-44b8-bf59-4da9e0cb683c	8143ab94-1464-4cc8-9236-cb541a045bb2
8eb26a12-1853-44b8-bf59-4da9e0cb683c	a4df278e-8a69-4a4c-b53f-53325b4fbfe6
8eb26a12-1853-44b8-bf59-4da9e0cb683c	28379ac4-b0d7-45b5-89da-51d38484e289
8eb26a12-1853-44b8-bf59-4da9e0cb683c	245e9f9a-cc0b-4026-a922-3548120a11fc
8eb26a12-1853-44b8-bf59-4da9e0cb683c	6522efd1-7218-4939-bd7a-32ccf7727e89
8eb26a12-1853-44b8-bf59-4da9e0cb683c	095429cc-740d-40c5-9d1e-0f0ce3f3c789
8eb26a12-1853-44b8-bf59-4da9e0cb683c	8994378a-9b80-4fc1-b43f-90b34dd4bef9
8eb26a12-1853-44b8-bf59-4da9e0cb683c	97dd6d41-613e-4bf3-85af-982967e947b1
8eb26a12-1853-44b8-bf59-4da9e0cb683c	8ccf9449-b244-423c-b724-50f165af55ab
8eb26a12-1853-44b8-bf59-4da9e0cb683c	e278b6d3-25f1-4105-927a-e80a9e82a351
8eb26a12-1853-44b8-bf59-4da9e0cb683c	a0e29540-5e4f-428b-aebf-e5a2360570c5
818ed298-e735-4251-875a-0778a78aaf81	a0e29540-5e4f-428b-aebf-e5a2360570c5
818ed298-e735-4251-875a-0778a78aaf81	97dd6d41-613e-4bf3-85af-982967e947b1
e46df155-c26e-4253-9bce-998723e7ee41	8ccf9449-b244-423c-b724-50f165af55ab
16c70d55-bcfe-495b-95ae-1280ee7dca70	7ebc1499-41ab-40a5-8331-79248f1b0372
3422313a-546d-4e5d-9595-9af9a14532fe	d1220360-a04a-4765-a21e-767bf9848eb9
4d43ab7e-ce46-49f2-a2ab-9296a6a13aa6	b03e73e7-b952-4f18-bff0-fe7dd5cb72e2
4d43ab7e-ce46-49f2-a2ab-9296a6a13aa6	103043c1-8e66-4507-9dad-0ae3e02a0801
97b3fa9d-3031-4405-8f19-95b0a51fee30	f50d9a09-df21-4537-8d08-79c33fadf74f
97b3fa9d-3031-4405-8f19-95b0a51fee30	82d458ac-cb98-4cac-9bda-b0d2f0615238
97b3fa9d-3031-4405-8f19-95b0a51fee30	3c0bfa33-a92b-41e3-97e7-f38946afa659
97b3fa9d-3031-4405-8f19-95b0a51fee30	1a722e8c-695d-439a-b271-1aab67cbf3cb
97b3fa9d-3031-4405-8f19-95b0a51fee30	a0d3ca7a-9748-4dea-8c25-f417a0d60899
97b3fa9d-3031-4405-8f19-95b0a51fee30	b03e73e7-b952-4f18-bff0-fe7dd5cb72e2
97b3fa9d-3031-4405-8f19-95b0a51fee30	103043c1-8e66-4507-9dad-0ae3e02a0801
97b3fa9d-3031-4405-8f19-95b0a51fee30	4d43ab7e-ce46-49f2-a2ab-9296a6a13aa6
97b3fa9d-3031-4405-8f19-95b0a51fee30	f8483adf-20d3-4c60-89bf-423fb7127254
97b3fa9d-3031-4405-8f19-95b0a51fee30	755d6e23-cf4f-41b4-8b75-7e6eb1e56f76
97b3fa9d-3031-4405-8f19-95b0a51fee30	c60082f0-7d04-4c53-89e0-a18bdef5ccbd
97b3fa9d-3031-4405-8f19-95b0a51fee30	b67f2539-d402-4b27-844f-8a62c067893e
97b3fa9d-3031-4405-8f19-95b0a51fee30	259881c3-5ae3-48b2-a21e-5baa30805bef
97b3fa9d-3031-4405-8f19-95b0a51fee30	15672aea-3a89-42bd-99e5-e316c39e0ccf
97b3fa9d-3031-4405-8f19-95b0a51fee30	5c4876da-376b-4db3-b8e2-7077c4d38455
97b3fa9d-3031-4405-8f19-95b0a51fee30	429bd26e-2f7f-4150-bb87-a677fd7eac4a
97b3fa9d-3031-4405-8f19-95b0a51fee30	ccef0b3d-88cf-45e0-99a4-61b463442bcd
97b3fa9d-3031-4405-8f19-95b0a51fee30	1ff58d47-9fa7-4e20-a4c4-3072fae9b049
c883c84f-d797-4170-bcf1-888f3843ba15	f7b61645-41f7-4916-971d-bacb13088a1d
c883c84f-d797-4170-bcf1-888f3843ba15	51193cbf-31d0-4955-8374-b2bee6bff6c4
c883c84f-d797-4170-bcf1-888f3843ba15	49c34741-45f3-4edd-8259-19b827460e3b
c883c84f-d797-4170-bcf1-888f3843ba15	3422313a-546d-4e5d-9595-9af9a14532fe
f50d9a09-df21-4537-8d08-79c33fadf74f	259881c3-5ae3-48b2-a21e-5baa30805bef
8eb26a12-1853-44b8-bf59-4da9e0cb683c	9d546e45-4e4b-437d-88d4-328eece3d9ac
\.


--
-- Data for Name: credential; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.credential (id, salt, type, user_id, created_date, user_label, secret_data, credential_data, priority) FROM stdin;
08c1dcd8-c7da-4851-b18d-3fd2fe5905ef	\N	password	f3fc1dd9-af7b-498a-9435-31da080a37ad	1719326096351	\N	{"value":"fNSqW2YM3jIEBSKVk6zogGGIerlNBxlII/d0jCJmmIgahPKX3CmI0GKBbz2arwR5O6a4DulX0v6KGja33O3Y9w==","salt":"lrvDVTkV44mlKYR0H7ahKQ==","additionalParameters":{}}	{"hashIterations":27500,"algorithm":"pbkdf2-sha256","additionalParameters":{}}	10
eb68ec6c-b233-467a-83de-f8907db05bf8	\N	password	f8dcafea-243e-4b89-8d7d-fa01918130f4	1719327627809	My password	{"value":"rS7hFhgurbYTqf2xOl21TK+ma64Za4bGekInBymZql67Fw9onWH8ghYQrIXhI/Px7RtIdevaAwXtGEnxGBEtpg==","salt":"E4/Bwx3A3YV1wur9bAI4nQ==","additionalParameters":{}}	{"hashIterations":27500,"algorithm":"pbkdf2-sha256","additionalParameters":{}}	10
c0ddc8fb-b335-4908-979c-170956e5dfc1	\N	password	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a	1719438849822	My password	{"value":"cmmEQr11CDETApkFo92Ufi0h+qPfe2Rz34N6RNgdVYaOYZOcAwldY6d6M6spufpuGc62vGhZ0NBKlAR5/u8AHw==","salt":"oPBgKF5zunaSWRspE2bGVw==","additionalParameters":{}}	{"hashIterations":27500,"algorithm":"pbkdf2-sha256","additionalParameters":{}}	10
eea54938-e813-43ff-81b6-697354de1a11	\N	password	f9b33064-13d0-47d7-8294-fb8f0fac819f	1725660663731	My password	{"value":"ShYFD1RRiDgpcByhjWiOReh5+CQdQCt8yKGSwC7xKecM+A/dQ/gKwDZSndQjxIkJGRPmLXrIsS7N6CvTFUsHSQ==","salt":"iDVbMFzwwDYr4KkMfG74Vg==","additionalParameters":{}}	{"hashIterations":27500,"algorithm":"pbkdf2-sha256","additionalParameters":{}}	10
\.


--
-- Data for Name: databasechangelog; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.databasechangelog (id, author, filename, dateexecuted, orderexecuted, exectype, md5sum, description, comments, tag, liquibase, contexts, labels, deployment_id) FROM stdin;
1.0.0.Final-KEYCLOAK-5461	sthorger@redhat.com	META-INF/jpa-changelog-1.0.0.Final.xml	2024-06-25 14:34:53.032312	1	EXECUTED	8:bda77d94bf90182a1e30c24f1c155ec7	createTable tableName=APPLICATION_DEFAULT_ROLES; createTable tableName=CLIENT; createTable tableName=CLIENT_SESSION; createTable tableName=CLIENT_SESSION_ROLE; createTable tableName=COMPOSITE_ROLE; createTable tableName=CREDENTIAL; createTable tab...		\N	4.8.0	\N	\N	9326092871
1.0.0.Final-KEYCLOAK-5461	sthorger@redhat.com	META-INF/db2-jpa-changelog-1.0.0.Final.xml	2024-06-25 14:34:53.037349	2	MARK_RAN	8:1ecb330f30986693d1cba9ab579fa219	createTable tableName=APPLICATION_DEFAULT_ROLES; createTable tableName=CLIENT; createTable tableName=CLIENT_SESSION; createTable tableName=CLIENT_SESSION_ROLE; createTable tableName=COMPOSITE_ROLE; createTable tableName=CREDENTIAL; createTable tab...		\N	4.8.0	\N	\N	9326092871
1.1.0.Beta1	sthorger@redhat.com	META-INF/jpa-changelog-1.1.0.Beta1.xml	2024-06-25 14:34:53.059085	3	EXECUTED	8:cb7ace19bc6d959f305605d255d4c843	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION; createTable tableName=CLIENT_ATTRIBUTES; createTable tableName=CLIENT_SESSION_NOTE; createTable tableName=APP_NODE_REGISTRATIONS; addColumn table...		\N	4.8.0	\N	\N	9326092871
1.1.0.Final	sthorger@redhat.com	META-INF/jpa-changelog-1.1.0.Final.xml	2024-06-25 14:34:53.062364	4	EXECUTED	8:80230013e961310e6872e871be424a63	renameColumn newColumnName=EVENT_TIME, oldColumnName=TIME, tableName=EVENT_ENTITY		\N	4.8.0	\N	\N	9326092871
1.2.0.Beta1	psilva@redhat.com	META-INF/jpa-changelog-1.2.0.Beta1.xml	2024-06-25 14:34:53.141184	5	EXECUTED	8:67f4c20929126adc0c8e9bf48279d244	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION; createTable tableName=PROTOCOL_MAPPER; createTable tableName=PROTOCOL_MAPPER_CONFIG; createTable tableName=...		\N	4.8.0	\N	\N	9326092871
1.2.0.Beta1	psilva@redhat.com	META-INF/db2-jpa-changelog-1.2.0.Beta1.xml	2024-06-25 14:34:53.142035	6	MARK_RAN	8:7311018b0b8179ce14628ab412bb6783	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION; createTable tableName=PROTOCOL_MAPPER; createTable tableName=PROTOCOL_MAPPER_CONFIG; createTable tableName=...		\N	4.8.0	\N	\N	9326092871
1.2.0.RC1	bburke@redhat.com	META-INF/jpa-changelog-1.2.0.CR1.xml	2024-06-25 14:34:53.174278	7	EXECUTED	8:037ba1216c3640f8785ee6b8e7c8e3c1	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete tableName=USER_SESSION; createTable tableName=MIGRATION_MODEL; createTable tableName=IDENTITY_P...		\N	4.8.0	\N	\N	9326092871
1.2.0.RC1	bburke@redhat.com	META-INF/db2-jpa-changelog-1.2.0.CR1.xml	2024-06-25 14:34:53.175091	8	MARK_RAN	8:7fe6ffe4af4df289b3157de32c624263	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete tableName=USER_SESSION; createTable tableName=MIGRATION_MODEL; createTable tableName=IDENTITY_P...		\N	4.8.0	\N	\N	9326092871
1.2.0.Final	keycloak	META-INF/jpa-changelog-1.2.0.Final.xml	2024-06-25 14:34:53.176474	9	EXECUTED	8:9c136bc3187083a98745c7d03bc8a303	update tableName=CLIENT; update tableName=CLIENT; update tableName=CLIENT		\N	4.8.0	\N	\N	9326092871
1.3.0	bburke@redhat.com	META-INF/jpa-changelog-1.3.0.xml	2024-06-25 14:34:53.196704	10	EXECUTED	8:b5f09474dca81fb56a97cf5b6553d331	delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_PROT_MAPPER; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete tableName=USER_SESSION; createTable tableName=ADMI...		\N	4.8.0	\N	\N	9326092871
1.4.0	bburke@redhat.com	META-INF/jpa-changelog-1.4.0.xml	2024-06-25 14:34:53.210486	11	EXECUTED	8:ca924f31bd2a3b219fdcfe78c82dacf4	delete tableName=CLIENT_SESSION_AUTH_STATUS; delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_PROT_MAPPER; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete table...		\N	4.8.0	\N	\N	9326092871
1.4.0	bburke@redhat.com	META-INF/db2-jpa-changelog-1.4.0.xml	2024-06-25 14:34:53.211071	12	MARK_RAN	8:8acad7483e106416bcfa6f3b824a16cd	delete tableName=CLIENT_SESSION_AUTH_STATUS; delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_PROT_MAPPER; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete table...		\N	4.8.0	\N	\N	9326092871
1.5.0	bburke@redhat.com	META-INF/jpa-changelog-1.5.0.xml	2024-06-25 14:34:53.215386	13	EXECUTED	8:9b1266d17f4f87c78226f5055408fd5e	delete tableName=CLIENT_SESSION_AUTH_STATUS; delete tableName=CLIENT_SESSION_ROLE; delete tableName=CLIENT_SESSION_PROT_MAPPER; delete tableName=CLIENT_SESSION_NOTE; delete tableName=CLIENT_SESSION; delete tableName=USER_SESSION_NOTE; delete table...		\N	4.8.0	\N	\N	9326092871
1.6.1_from15	mposolda@redhat.com	META-INF/jpa-changelog-1.6.1.xml	2024-06-25 14:34:53.218587	14	EXECUTED	8:d80ec4ab6dbfe573550ff72396c7e910	addColumn tableName=REALM; addColumn tableName=KEYCLOAK_ROLE; addColumn tableName=CLIENT; createTable tableName=OFFLINE_USER_SESSION; createTable tableName=OFFLINE_CLIENT_SESSION; addPrimaryKey constraintName=CONSTRAINT_OFFL_US_SES_PK2, tableName=...		\N	4.8.0	\N	\N	9326092871
1.6.1_from16-pre	mposolda@redhat.com	META-INF/jpa-changelog-1.6.1.xml	2024-06-25 14:34:53.219006	15	MARK_RAN	8:d86eb172171e7c20b9c849b584d147b2	delete tableName=OFFLINE_CLIENT_SESSION; delete tableName=OFFLINE_USER_SESSION		\N	4.8.0	\N	\N	9326092871
1.6.1_from16	mposolda@redhat.com	META-INF/jpa-changelog-1.6.1.xml	2024-06-25 14:34:53.219407	16	MARK_RAN	8:5735f46f0fa60689deb0ecdc2a0dea22	dropPrimaryKey constraintName=CONSTRAINT_OFFLINE_US_SES_PK, tableName=OFFLINE_USER_SESSION; dropPrimaryKey constraintName=CONSTRAINT_OFFLINE_CL_SES_PK, tableName=OFFLINE_CLIENT_SESSION; addColumn tableName=OFFLINE_USER_SESSION; update tableName=OF...		\N	4.8.0	\N	\N	9326092871
1.6.1	mposolda@redhat.com	META-INF/jpa-changelog-1.6.1.xml	2024-06-25 14:34:53.219784	17	EXECUTED	8:d41d8cd98f00b204e9800998ecf8427e	empty		\N	4.8.0	\N	\N	9326092871
1.7.0	bburke@redhat.com	META-INF/jpa-changelog-1.7.0.xml	2024-06-25 14:34:53.227803	18	EXECUTED	8:5c1a8fd2014ac7fc43b90a700f117b23	createTable tableName=KEYCLOAK_GROUP; createTable tableName=GROUP_ROLE_MAPPING; createTable tableName=GROUP_ATTRIBUTE; createTable tableName=USER_GROUP_MEMBERSHIP; createTable tableName=REALM_DEFAULT_GROUPS; addColumn tableName=IDENTITY_PROVIDER; ...		\N	4.8.0	\N	\N	9326092871
1.8.0	mposolda@redhat.com	META-INF/jpa-changelog-1.8.0.xml	2024-06-25 14:34:53.235555	19	EXECUTED	8:1f6c2c2dfc362aff4ed75b3f0ef6b331	addColumn tableName=IDENTITY_PROVIDER; createTable tableName=CLIENT_TEMPLATE; createTable tableName=CLIENT_TEMPLATE_ATTRIBUTES; createTable tableName=TEMPLATE_SCOPE_MAPPING; dropNotNullConstraint columnName=CLIENT_ID, tableName=PROTOCOL_MAPPER; ad...		\N	4.8.0	\N	\N	9326092871
1.8.0-2	keycloak	META-INF/jpa-changelog-1.8.0.xml	2024-06-25 14:34:53.236452	20	EXECUTED	8:dee9246280915712591f83a127665107	dropDefaultValue columnName=ALGORITHM, tableName=CREDENTIAL; update tableName=CREDENTIAL		\N	4.8.0	\N	\N	9326092871
authz-3.4.0.CR1-resource-server-pk-change-part1	glavoie@gmail.com	META-INF/jpa-changelog-authz-3.4.0.CR1.xml	2024-06-25 14:34:53.336626	45	EXECUTED	8:a164ae073c56ffdbc98a615493609a52	addColumn tableName=RESOURCE_SERVER_POLICY; addColumn tableName=RESOURCE_SERVER_RESOURCE; addColumn tableName=RESOURCE_SERVER_SCOPE		\N	4.8.0	\N	\N	9326092871
1.8.0	mposolda@redhat.com	META-INF/db2-jpa-changelog-1.8.0.xml	2024-06-25 14:34:53.237366	21	MARK_RAN	8:9eb2ee1fa8ad1c5e426421a6f8fdfa6a	addColumn tableName=IDENTITY_PROVIDER; createTable tableName=CLIENT_TEMPLATE; createTable tableName=CLIENT_TEMPLATE_ATTRIBUTES; createTable tableName=TEMPLATE_SCOPE_MAPPING; dropNotNullConstraint columnName=CLIENT_ID, tableName=PROTOCOL_MAPPER; ad...		\N	4.8.0	\N	\N	9326092871
1.8.0-2	keycloak	META-INF/db2-jpa-changelog-1.8.0.xml	2024-06-25 14:34:53.237822	22	MARK_RAN	8:dee9246280915712591f83a127665107	dropDefaultValue columnName=ALGORITHM, tableName=CREDENTIAL; update tableName=CREDENTIAL		\N	4.8.0	\N	\N	9326092871
1.9.0	mposolda@redhat.com	META-INF/jpa-changelog-1.9.0.xml	2024-06-25 14:34:53.243468	23	EXECUTED	8:d9fa18ffa355320395b86270680dd4fe	update tableName=REALM; update tableName=REALM; update tableName=REALM; update tableName=REALM; update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=REALM; update tableName=REALM; customChange; dr...		\N	4.8.0	\N	\N	9326092871
1.9.1	keycloak	META-INF/jpa-changelog-1.9.1.xml	2024-06-25 14:34:53.244594	24	EXECUTED	8:90cff506fedb06141ffc1c71c4a1214c	modifyDataType columnName=PRIVATE_KEY, tableName=REALM; modifyDataType columnName=PUBLIC_KEY, tableName=REALM; modifyDataType columnName=CERTIFICATE, tableName=REALM		\N	4.8.0	\N	\N	9326092871
1.9.1	keycloak	META-INF/db2-jpa-changelog-1.9.1.xml	2024-06-25 14:34:53.244973	25	MARK_RAN	8:11a788aed4961d6d29c427c063af828c	modifyDataType columnName=PRIVATE_KEY, tableName=REALM; modifyDataType columnName=CERTIFICATE, tableName=REALM		\N	4.8.0	\N	\N	9326092871
1.9.2	keycloak	META-INF/jpa-changelog-1.9.2.xml	2024-06-25 14:34:53.248001	26	EXECUTED	8:a4218e51e1faf380518cce2af5d39b43	createIndex indexName=IDX_USER_EMAIL, tableName=USER_ENTITY; createIndex indexName=IDX_USER_ROLE_MAPPING, tableName=USER_ROLE_MAPPING; createIndex indexName=IDX_USER_GROUP_MAPPING, tableName=USER_GROUP_MEMBERSHIP; createIndex indexName=IDX_USER_CO...		\N	4.8.0	\N	\N	9326092871
authz-2.0.0	psilva@redhat.com	META-INF/jpa-changelog-authz-2.0.0.xml	2024-06-25 14:34:53.259937	27	EXECUTED	8:d9e9a1bfaa644da9952456050f07bbdc	createTable tableName=RESOURCE_SERVER; addPrimaryKey constraintName=CONSTRAINT_FARS, tableName=RESOURCE_SERVER; addUniqueConstraint constraintName=UK_AU8TT6T700S9V50BU18WS5HA6, tableName=RESOURCE_SERVER; createTable tableName=RESOURCE_SERVER_RESOU...		\N	4.8.0	\N	\N	9326092871
authz-2.5.1	psilva@redhat.com	META-INF/jpa-changelog-authz-2.5.1.xml	2024-06-25 14:34:53.260745	28	EXECUTED	8:d1bf991a6163c0acbfe664b615314505	update tableName=RESOURCE_SERVER_POLICY		\N	4.8.0	\N	\N	9326092871
2.1.0-KEYCLOAK-5461	bburke@redhat.com	META-INF/jpa-changelog-2.1.0.xml	2024-06-25 14:34:53.275315	29	EXECUTED	8:88a743a1e87ec5e30bf603da68058a8c	createTable tableName=BROKER_LINK; createTable tableName=FED_USER_ATTRIBUTE; createTable tableName=FED_USER_CONSENT; createTable tableName=FED_USER_CONSENT_ROLE; createTable tableName=FED_USER_CONSENT_PROT_MAPPER; createTable tableName=FED_USER_CR...		\N	4.8.0	\N	\N	9326092871
2.2.0	bburke@redhat.com	META-INF/jpa-changelog-2.2.0.xml	2024-06-25 14:34:53.281633	30	EXECUTED	8:c5517863c875d325dea463d00ec26d7a	addColumn tableName=ADMIN_EVENT_ENTITY; createTable tableName=CREDENTIAL_ATTRIBUTE; createTable tableName=FED_CREDENTIAL_ATTRIBUTE; modifyDataType columnName=VALUE, tableName=CREDENTIAL; addForeignKeyConstraint baseTableName=FED_CREDENTIAL_ATTRIBU...		\N	4.8.0	\N	\N	9326092871
2.3.0	bburke@redhat.com	META-INF/jpa-changelog-2.3.0.xml	2024-06-25 14:34:53.28916	31	EXECUTED	8:ada8b4833b74a498f376d7136bc7d327	createTable tableName=FEDERATED_USER; addPrimaryKey constraintName=CONSTR_FEDERATED_USER, tableName=FEDERATED_USER; dropDefaultValue columnName=TOTP, tableName=USER_ENTITY; dropColumn columnName=TOTP, tableName=USER_ENTITY; addColumn tableName=IDE...		\N	4.8.0	\N	\N	9326092871
2.4.0	bburke@redhat.com	META-INF/jpa-changelog-2.4.0.xml	2024-06-25 14:34:53.290321	32	EXECUTED	8:b9b73c8ea7299457f99fcbb825c263ba	customChange		\N	4.8.0	\N	\N	9326092871
2.5.0	bburke@redhat.com	META-INF/jpa-changelog-2.5.0.xml	2024-06-25 14:34:53.291497	33	EXECUTED	8:07724333e625ccfcfc5adc63d57314f3	customChange; modifyDataType columnName=USER_ID, tableName=OFFLINE_USER_SESSION		\N	4.8.0	\N	\N	9326092871
2.5.0-unicode-oracle	hmlnarik@redhat.com	META-INF/jpa-changelog-2.5.0.xml	2024-06-25 14:34:53.291909	34	MARK_RAN	8:8b6fd445958882efe55deb26fc541a7b	modifyDataType columnName=DESCRIPTION, tableName=AUTHENTICATION_FLOW; modifyDataType columnName=DESCRIPTION, tableName=CLIENT_TEMPLATE; modifyDataType columnName=DESCRIPTION, tableName=RESOURCE_SERVER_POLICY; modifyDataType columnName=DESCRIPTION,...		\N	4.8.0	\N	\N	9326092871
2.5.0-unicode-other-dbs	hmlnarik@redhat.com	META-INF/jpa-changelog-2.5.0.xml	2024-06-25 14:34:53.303622	35	EXECUTED	8:29b29cfebfd12600897680147277a9d7	modifyDataType columnName=DESCRIPTION, tableName=AUTHENTICATION_FLOW; modifyDataType columnName=DESCRIPTION, tableName=CLIENT_TEMPLATE; modifyDataType columnName=DESCRIPTION, tableName=RESOURCE_SERVER_POLICY; modifyDataType columnName=DESCRIPTION,...		\N	4.8.0	\N	\N	9326092871
2.5.0-duplicate-email-support	slawomir@dabek.name	META-INF/jpa-changelog-2.5.0.xml	2024-06-25 14:34:53.304779	36	EXECUTED	8:73ad77ca8fd0410c7f9f15a471fa52bc	addColumn tableName=REALM		\N	4.8.0	\N	\N	9326092871
2.5.0-unique-group-names	hmlnarik@redhat.com	META-INF/jpa-changelog-2.5.0.xml	2024-06-25 14:34:53.30559	37	EXECUTED	8:64f27a6fdcad57f6f9153210f2ec1bdb	addUniqueConstraint constraintName=SIBLING_NAMES, tableName=KEYCLOAK_GROUP		\N	4.8.0	\N	\N	9326092871
2.5.1	bburke@redhat.com	META-INF/jpa-changelog-2.5.1.xml	2024-06-25 14:34:53.306244	38	EXECUTED	8:27180251182e6c31846c2ddab4bc5781	addColumn tableName=FED_USER_CONSENT		\N	4.8.0	\N	\N	9326092871
3.0.0	bburke@redhat.com	META-INF/jpa-changelog-3.0.0.xml	2024-06-25 14:34:53.306847	39	EXECUTED	8:d56f201bfcfa7a1413eb3e9bc02978f9	addColumn tableName=IDENTITY_PROVIDER		\N	4.8.0	\N	\N	9326092871
3.2.0-fix	keycloak	META-INF/jpa-changelog-3.2.0.xml	2024-06-25 14:34:53.307206	40	MARK_RAN	8:91f5522bf6afdc2077dfab57fbd3455c	addNotNullConstraint columnName=REALM_ID, tableName=CLIENT_INITIAL_ACCESS		\N	4.8.0	\N	\N	9326092871
3.2.0-fix-with-keycloak-5416	keycloak	META-INF/jpa-changelog-3.2.0.xml	2024-06-25 14:34:53.307575	41	MARK_RAN	8:0f01b554f256c22caeb7d8aee3a1cdc8	dropIndex indexName=IDX_CLIENT_INIT_ACC_REALM, tableName=CLIENT_INITIAL_ACCESS; addNotNullConstraint columnName=REALM_ID, tableName=CLIENT_INITIAL_ACCESS; createIndex indexName=IDX_CLIENT_INIT_ACC_REALM, tableName=CLIENT_INITIAL_ACCESS		\N	4.8.0	\N	\N	9326092871
3.2.0-fix-offline-sessions	hmlnarik	META-INF/jpa-changelog-3.2.0.xml	2024-06-25 14:34:53.310356	42	EXECUTED	8:ab91cf9cee415867ade0e2df9651a947	customChange		\N	4.8.0	\N	\N	9326092871
3.2.0-fixed	keycloak	META-INF/jpa-changelog-3.2.0.xml	2024-06-25 14:34:53.334951	43	EXECUTED	8:ceac9b1889e97d602caf373eadb0d4b7	addColumn tableName=REALM; dropPrimaryKey constraintName=CONSTRAINT_OFFL_CL_SES_PK2, tableName=OFFLINE_CLIENT_SESSION; dropColumn columnName=CLIENT_SESSION_ID, tableName=OFFLINE_CLIENT_SESSION; addPrimaryKey constraintName=CONSTRAINT_OFFL_CL_SES_P...		\N	4.8.0	\N	\N	9326092871
3.3.0	keycloak	META-INF/jpa-changelog-3.3.0.xml	2024-06-25 14:34:53.33578	44	EXECUTED	8:84b986e628fe8f7fd8fd3c275c5259f2	addColumn tableName=USER_ENTITY		\N	4.8.0	\N	\N	9326092871
authz-3.4.0.CR1-resource-server-pk-change-part2-KEYCLOAK-6095	hmlnarik@redhat.com	META-INF/jpa-changelog-authz-3.4.0.CR1.xml	2024-06-25 14:34:53.337842	46	EXECUTED	8:70a2b4f1f4bd4dbf487114bdb1810e64	customChange		\N	4.8.0	\N	\N	9326092871
authz-3.4.0.CR1-resource-server-pk-change-part3-fixed	glavoie@gmail.com	META-INF/jpa-changelog-authz-3.4.0.CR1.xml	2024-06-25 14:34:53.33818	47	MARK_RAN	8:7be68b71d2f5b94b8df2e824f2860fa2	dropIndex indexName=IDX_RES_SERV_POL_RES_SERV, tableName=RESOURCE_SERVER_POLICY; dropIndex indexName=IDX_RES_SRV_RES_RES_SRV, tableName=RESOURCE_SERVER_RESOURCE; dropIndex indexName=IDX_RES_SRV_SCOPE_RES_SRV, tableName=RESOURCE_SERVER_SCOPE		\N	4.8.0	\N	\N	9326092871
authz-3.4.0.CR1-resource-server-pk-change-part3-fixed-nodropindex	glavoie@gmail.com	META-INF/jpa-changelog-authz-3.4.0.CR1.xml	2024-06-25 14:34:53.344882	48	EXECUTED	8:bab7c631093c3861d6cf6144cd944982	addNotNullConstraint columnName=RESOURCE_SERVER_CLIENT_ID, tableName=RESOURCE_SERVER_POLICY; addNotNullConstraint columnName=RESOURCE_SERVER_CLIENT_ID, tableName=RESOURCE_SERVER_RESOURCE; addNotNullConstraint columnName=RESOURCE_SERVER_CLIENT_ID, ...		\N	4.8.0	\N	\N	9326092871
authn-3.4.0.CR1-refresh-token-max-reuse	glavoie@gmail.com	META-INF/jpa-changelog-authz-3.4.0.CR1.xml	2024-06-25 14:34:53.345568	49	EXECUTED	8:fa809ac11877d74d76fe40869916daad	addColumn tableName=REALM		\N	4.8.0	\N	\N	9326092871
3.4.0	keycloak	META-INF/jpa-changelog-3.4.0.xml	2024-06-25 14:34:53.350381	50	EXECUTED	8:fac23540a40208f5f5e326f6ceb4d291	addPrimaryKey constraintName=CONSTRAINT_REALM_DEFAULT_ROLES, tableName=REALM_DEFAULT_ROLES; addPrimaryKey constraintName=CONSTRAINT_COMPOSITE_ROLE, tableName=COMPOSITE_ROLE; addPrimaryKey constraintName=CONSTR_REALM_DEFAULT_GROUPS, tableName=REALM...		\N	4.8.0	\N	\N	9326092871
3.4.0-KEYCLOAK-5230	hmlnarik@redhat.com	META-INF/jpa-changelog-3.4.0.xml	2024-06-25 14:34:53.353188	51	EXECUTED	8:2612d1b8a97e2b5588c346e817307593	createIndex indexName=IDX_FU_ATTRIBUTE, tableName=FED_USER_ATTRIBUTE; createIndex indexName=IDX_FU_CONSENT, tableName=FED_USER_CONSENT; createIndex indexName=IDX_FU_CONSENT_RU, tableName=FED_USER_CONSENT; createIndex indexName=IDX_FU_CREDENTIAL, t...		\N	4.8.0	\N	\N	9326092871
3.4.1	psilva@redhat.com	META-INF/jpa-changelog-3.4.1.xml	2024-06-25 14:34:53.353763	52	EXECUTED	8:9842f155c5db2206c88bcb5d1046e941	modifyDataType columnName=VALUE, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
3.4.2	keycloak	META-INF/jpa-changelog-3.4.2.xml	2024-06-25 14:34:53.354308	53	EXECUTED	8:2e12e06e45498406db72d5b3da5bbc76	update tableName=REALM		\N	4.8.0	\N	\N	9326092871
3.4.2-KEYCLOAK-5172	mkanis@redhat.com	META-INF/jpa-changelog-3.4.2.xml	2024-06-25 14:34:53.354792	54	EXECUTED	8:33560e7c7989250c40da3abdabdc75a4	update tableName=CLIENT		\N	4.8.0	\N	\N	9326092871
4.0.0-KEYCLOAK-6335	bburke@redhat.com	META-INF/jpa-changelog-4.0.0.xml	2024-06-25 14:34:53.355672	55	EXECUTED	8:87a8d8542046817a9107c7eb9cbad1cd	createTable tableName=CLIENT_AUTH_FLOW_BINDINGS; addPrimaryKey constraintName=C_CLI_FLOW_BIND, tableName=CLIENT_AUTH_FLOW_BINDINGS		\N	4.8.0	\N	\N	9326092871
4.0.0-CLEANUP-UNUSED-TABLE	bburke@redhat.com	META-INF/jpa-changelog-4.0.0.xml	2024-06-25 14:34:53.356611	56	EXECUTED	8:3ea08490a70215ed0088c273d776311e	dropTable tableName=CLIENT_IDENTITY_PROV_MAPPING		\N	4.8.0	\N	\N	9326092871
4.0.0-KEYCLOAK-6228	bburke@redhat.com	META-INF/jpa-changelog-4.0.0.xml	2024-06-25 14:34:53.359908	57	EXECUTED	8:2d56697c8723d4592ab608ce14b6ed68	dropUniqueConstraint constraintName=UK_JKUWUVD56ONTGSUHOGM8UEWRT, tableName=USER_CONSENT; dropNotNullConstraint columnName=CLIENT_ID, tableName=USER_CONSENT; addColumn tableName=USER_CONSENT; addUniqueConstraint constraintName=UK_JKUWUVD56ONTGSUHO...		\N	4.8.0	\N	\N	9326092871
4.0.0-KEYCLOAK-5579-fixed	mposolda@redhat.com	META-INF/jpa-changelog-4.0.0.xml	2024-06-25 14:34:53.379585	58	EXECUTED	8:3e423e249f6068ea2bbe48bf907f9d86	dropForeignKeyConstraint baseTableName=CLIENT_TEMPLATE_ATTRIBUTES, constraintName=FK_CL_TEMPL_ATTR_TEMPL; renameTable newTableName=CLIENT_SCOPE_ATTRIBUTES, oldTableName=CLIENT_TEMPLATE_ATTRIBUTES; renameColumn newColumnName=SCOPE_ID, oldColumnName...		\N	4.8.0	\N	\N	9326092871
authz-4.0.0.CR1	psilva@redhat.com	META-INF/jpa-changelog-authz-4.0.0.CR1.xml	2024-06-25 14:34:53.386341	59	EXECUTED	8:15cabee5e5df0ff099510a0fc03e4103	createTable tableName=RESOURCE_SERVER_PERM_TICKET; addPrimaryKey constraintName=CONSTRAINT_FAPMT, tableName=RESOURCE_SERVER_PERM_TICKET; addForeignKeyConstraint baseTableName=RESOURCE_SERVER_PERM_TICKET, constraintName=FK_FRSRHO213XCX4WNKOG82SSPMT...		\N	4.8.0	\N	\N	9326092871
authz-4.0.0.Beta3	psilva@redhat.com	META-INF/jpa-changelog-authz-4.0.0.Beta3.xml	2024-06-25 14:34:53.387831	60	EXECUTED	8:4b80200af916ac54d2ffbfc47918ab0e	addColumn tableName=RESOURCE_SERVER_POLICY; addColumn tableName=RESOURCE_SERVER_PERM_TICKET; addForeignKeyConstraint baseTableName=RESOURCE_SERVER_PERM_TICKET, constraintName=FK_FRSRPO2128CX4WNKOG82SSRFY, referencedTableName=RESOURCE_SERVER_POLICY		\N	4.8.0	\N	\N	9326092871
authz-4.2.0.Final	mhajas@redhat.com	META-INF/jpa-changelog-authz-4.2.0.Final.xml	2024-06-25 14:34:53.401539	61	EXECUTED	8:66564cd5e168045d52252c5027485bbb	createTable tableName=RESOURCE_URIS; addForeignKeyConstraint baseTableName=RESOURCE_URIS, constraintName=FK_RESOURCE_SERVER_URIS, referencedTableName=RESOURCE_SERVER_RESOURCE; customChange; dropColumn columnName=URI, tableName=RESOURCE_SERVER_RESO...		\N	4.8.0	\N	\N	9326092871
authz-4.2.0.Final-KEYCLOAK-9944	hmlnarik@redhat.com	META-INF/jpa-changelog-authz-4.2.0.Final.xml	2024-06-25 14:34:53.402422	62	EXECUTED	8:1c7064fafb030222be2bd16ccf690f6f	addPrimaryKey constraintName=CONSTRAINT_RESOUR_URIS_PK, tableName=RESOURCE_URIS		\N	4.8.0	\N	\N	9326092871
4.2.0-KEYCLOAK-6313	wadahiro@gmail.com	META-INF/jpa-changelog-4.2.0.xml	2024-06-25 14:34:53.403056	63	EXECUTED	8:2de18a0dce10cdda5c7e65c9b719b6e5	addColumn tableName=REQUIRED_ACTION_PROVIDER		\N	4.8.0	\N	\N	9326092871
4.3.0-KEYCLOAK-7984	wadahiro@gmail.com	META-INF/jpa-changelog-4.3.0.xml	2024-06-25 14:34:53.403666	64	EXECUTED	8:03e413dd182dcbd5c57e41c34d0ef682	update tableName=REQUIRED_ACTION_PROVIDER		\N	4.8.0	\N	\N	9326092871
4.6.0-KEYCLOAK-7950	psilva@redhat.com	META-INF/jpa-changelog-4.6.0.xml	2024-06-25 14:34:53.404558	65	EXECUTED	8:d27b42bb2571c18fbe3fe4e4fb7582a7	update tableName=RESOURCE_SERVER_RESOURCE		\N	4.8.0	\N	\N	9326092871
4.6.0-KEYCLOAK-8377	keycloak	META-INF/jpa-changelog-4.6.0.xml	2024-06-25 14:34:53.406428	66	EXECUTED	8:698baf84d9fd0027e9192717c2154fb8	createTable tableName=ROLE_ATTRIBUTE; addPrimaryKey constraintName=CONSTRAINT_ROLE_ATTRIBUTE_PK, tableName=ROLE_ATTRIBUTE; addForeignKeyConstraint baseTableName=ROLE_ATTRIBUTE, constraintName=FK_ROLE_ATTRIBUTE_ID, referencedTableName=KEYCLOAK_ROLE...		\N	4.8.0	\N	\N	9326092871
4.6.0-KEYCLOAK-8555	gideonray@gmail.com	META-INF/jpa-changelog-4.6.0.xml	2024-06-25 14:34:53.407061	67	EXECUTED	8:ced8822edf0f75ef26eb51582f9a821a	createIndex indexName=IDX_COMPONENT_PROVIDER_TYPE, tableName=COMPONENT		\N	4.8.0	\N	\N	9326092871
4.7.0-KEYCLOAK-1267	sguilhen@redhat.com	META-INF/jpa-changelog-4.7.0.xml	2024-06-25 14:34:53.40789	68	EXECUTED	8:f0abba004cf429e8afc43056df06487d	addColumn tableName=REALM		\N	4.8.0	\N	\N	9326092871
4.7.0-KEYCLOAK-7275	keycloak	META-INF/jpa-changelog-4.7.0.xml	2024-06-25 14:34:53.415918	69	EXECUTED	8:6662f8b0b611caa359fcf13bf63b4e24	renameColumn newColumnName=CREATED_ON, oldColumnName=LAST_SESSION_REFRESH, tableName=OFFLINE_USER_SESSION; addNotNullConstraint columnName=CREATED_ON, tableName=OFFLINE_USER_SESSION; addColumn tableName=OFFLINE_USER_SESSION; customChange; createIn...		\N	4.8.0	\N	\N	9326092871
4.8.0-KEYCLOAK-8835	sguilhen@redhat.com	META-INF/jpa-changelog-4.8.0.xml	2024-06-25 14:34:53.41781	70	EXECUTED	8:9e6b8009560f684250bdbdf97670d39e	addNotNullConstraint columnName=SSO_MAX_LIFESPAN_REMEMBER_ME, tableName=REALM; addNotNullConstraint columnName=SSO_IDLE_TIMEOUT_REMEMBER_ME, tableName=REALM		\N	4.8.0	\N	\N	9326092871
authz-7.0.0-KEYCLOAK-10443	psilva@redhat.com	META-INF/jpa-changelog-authz-7.0.0.xml	2024-06-25 14:34:53.419208	71	EXECUTED	8:4223f561f3b8dc655846562b57bb502e	addColumn tableName=RESOURCE_SERVER		\N	4.8.0	\N	\N	9326092871
8.0.0-adding-credential-columns	keycloak	META-INF/jpa-changelog-8.0.0.xml	2024-06-25 14:34:53.422893	72	EXECUTED	8:215a31c398b363ce383a2b301202f29e	addColumn tableName=CREDENTIAL; addColumn tableName=FED_USER_CREDENTIAL		\N	4.8.0	\N	\N	9326092871
8.0.0-updating-credential-data-not-oracle-fixed	keycloak	META-INF/jpa-changelog-8.0.0.xml	2024-06-25 14:34:53.427122	73	EXECUTED	8:83f7a671792ca98b3cbd3a1a34862d3d	update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=FED_USER_CREDENTIAL; update tableName=FED_USER_CREDENTIAL; update tableName=FED_USER_CREDENTIAL		\N	4.8.0	\N	\N	9326092871
8.0.0-updating-credential-data-oracle-fixed	keycloak	META-INF/jpa-changelog-8.0.0.xml	2024-06-25 14:34:53.427533	74	MARK_RAN	8:f58ad148698cf30707a6efbdf8061aa7	update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=CREDENTIAL; update tableName=FED_USER_CREDENTIAL; update tableName=FED_USER_CREDENTIAL; update tableName=FED_USER_CREDENTIAL		\N	4.8.0	\N	\N	9326092871
8.0.0-credential-cleanup-fixed	keycloak	META-INF/jpa-changelog-8.0.0.xml	2024-06-25 14:34:53.431977	75	EXECUTED	8:79e4fd6c6442980e58d52ffc3ee7b19c	dropDefaultValue columnName=COUNTER, tableName=CREDENTIAL; dropDefaultValue columnName=DIGITS, tableName=CREDENTIAL; dropDefaultValue columnName=PERIOD, tableName=CREDENTIAL; dropDefaultValue columnName=ALGORITHM, tableName=CREDENTIAL; dropColumn ...		\N	4.8.0	\N	\N	9326092871
8.0.0-resource-tag-support	keycloak	META-INF/jpa-changelog-8.0.0.xml	2024-06-25 14:34:53.432978	76	EXECUTED	8:87af6a1e6d241ca4b15801d1f86a297d	addColumn tableName=MIGRATION_MODEL; createIndex indexName=IDX_UPDATE_TIME, tableName=MIGRATION_MODEL		\N	4.8.0	\N	\N	9326092871
9.0.0-always-display-client	keycloak	META-INF/jpa-changelog-9.0.0.xml	2024-06-25 14:34:53.433611	77	EXECUTED	8:b44f8d9b7b6ea455305a6d72a200ed15	addColumn tableName=CLIENT		\N	4.8.0	\N	\N	9326092871
9.0.0-drop-constraints-for-column-increase	keycloak	META-INF/jpa-changelog-9.0.0.xml	2024-06-25 14:34:53.433932	78	MARK_RAN	8:2d8ed5aaaeffd0cb004c046b4a903ac5	dropUniqueConstraint constraintName=UK_FRSR6T700S9V50BU18WS5PMT, tableName=RESOURCE_SERVER_PERM_TICKET; dropUniqueConstraint constraintName=UK_FRSR6T700S9V50BU18WS5HA6, tableName=RESOURCE_SERVER_RESOURCE; dropPrimaryKey constraintName=CONSTRAINT_O...		\N	4.8.0	\N	\N	9326092871
9.0.0-increase-column-size-federated-fk	keycloak	META-INF/jpa-changelog-9.0.0.xml	2024-06-25 14:34:53.4373	79	EXECUTED	8:e290c01fcbc275326c511633f6e2acde	modifyDataType columnName=CLIENT_ID, tableName=FED_USER_CONSENT; modifyDataType columnName=CLIENT_REALM_CONSTRAINT, tableName=KEYCLOAK_ROLE; modifyDataType columnName=OWNER, tableName=RESOURCE_SERVER_POLICY; modifyDataType columnName=CLIENT_ID, ta...		\N	4.8.0	\N	\N	9326092871
9.0.0-recreate-constraints-after-column-increase	keycloak	META-INF/jpa-changelog-9.0.0.xml	2024-06-25 14:34:53.437685	80	MARK_RAN	8:c9db8784c33cea210872ac2d805439f8	addNotNullConstraint columnName=CLIENT_ID, tableName=OFFLINE_CLIENT_SESSION; addNotNullConstraint columnName=OWNER, tableName=RESOURCE_SERVER_PERM_TICKET; addNotNullConstraint columnName=REQUESTER, tableName=RESOURCE_SERVER_PERM_TICKET; addNotNull...		\N	4.8.0	\N	\N	9326092871
9.0.1-add-index-to-client.client_id	keycloak	META-INF/jpa-changelog-9.0.1.xml	2024-06-25 14:34:53.438281	81	EXECUTED	8:95b676ce8fc546a1fcfb4c92fae4add5	createIndex indexName=IDX_CLIENT_ID, tableName=CLIENT		\N	4.8.0	\N	\N	9326092871
9.0.1-KEYCLOAK-12579-drop-constraints	keycloak	META-INF/jpa-changelog-9.0.1.xml	2024-06-25 14:34:53.438573	82	MARK_RAN	8:38a6b2a41f5651018b1aca93a41401e5	dropUniqueConstraint constraintName=SIBLING_NAMES, tableName=KEYCLOAK_GROUP		\N	4.8.0	\N	\N	9326092871
9.0.1-KEYCLOAK-12579-add-not-null-constraint	keycloak	META-INF/jpa-changelog-9.0.1.xml	2024-06-25 14:34:53.439206	83	EXECUTED	8:3fb99bcad86a0229783123ac52f7609c	addNotNullConstraint columnName=PARENT_GROUP, tableName=KEYCLOAK_GROUP		\N	4.8.0	\N	\N	9326092871
9.0.1-KEYCLOAK-12579-recreate-constraints	keycloak	META-INF/jpa-changelog-9.0.1.xml	2024-06-25 14:34:53.439494	84	MARK_RAN	8:64f27a6fdcad57f6f9153210f2ec1bdb	addUniqueConstraint constraintName=SIBLING_NAMES, tableName=KEYCLOAK_GROUP		\N	4.8.0	\N	\N	9326092871
9.0.1-add-index-to-events	keycloak	META-INF/jpa-changelog-9.0.1.xml	2024-06-25 14:34:53.440045	85	EXECUTED	8:ab4f863f39adafd4c862f7ec01890abc	createIndex indexName=IDX_EVENT_TIME, tableName=EVENT_ENTITY		\N	4.8.0	\N	\N	9326092871
map-remove-ri	keycloak	META-INF/jpa-changelog-11.0.0.xml	2024-06-25 14:34:53.440949	86	EXECUTED	8:13c419a0eb336e91ee3a3bf8fda6e2a7	dropForeignKeyConstraint baseTableName=REALM, constraintName=FK_TRAF444KK6QRKMS7N56AIWQ5Y; dropForeignKeyConstraint baseTableName=KEYCLOAK_ROLE, constraintName=FK_KJHO5LE2C0RAL09FL8CM9WFW9		\N	4.8.0	\N	\N	9326092871
map-remove-ri	keycloak	META-INF/jpa-changelog-12.0.0.xml	2024-06-25 14:34:53.442409	87	EXECUTED	8:e3fb1e698e0471487f51af1ed80fe3ac	dropForeignKeyConstraint baseTableName=REALM_DEFAULT_GROUPS, constraintName=FK_DEF_GROUPS_GROUP; dropForeignKeyConstraint baseTableName=REALM_DEFAULT_ROLES, constraintName=FK_H4WPD7W4HSOOLNI3H0SW7BTJE; dropForeignKeyConstraint baseTableName=CLIENT...		\N	4.8.0	\N	\N	9326092871
12.1.0-add-realm-localization-table	keycloak	META-INF/jpa-changelog-12.0.0.xml	2024-06-25 14:34:53.443509	88	EXECUTED	8:babadb686aab7b56562817e60bf0abd0	createTable tableName=REALM_LOCALIZATIONS; addPrimaryKey tableName=REALM_LOCALIZATIONS		\N	4.8.0	\N	\N	9326092871
default-roles	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.444819	89	EXECUTED	8:72d03345fda8e2f17093d08801947773	addColumn tableName=REALM; customChange		\N	4.8.0	\N	\N	9326092871
default-roles-cleanup	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.446084	90	EXECUTED	8:61c9233951bd96ffecd9ba75f7d978a4	dropTable tableName=REALM_DEFAULT_ROLES; dropTable tableName=CLIENT_DEFAULT_ROLES		\N	4.8.0	\N	\N	9326092871
13.0.0-KEYCLOAK-16844	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.44672	91	EXECUTED	8:ea82e6ad945cec250af6372767b25525	createIndex indexName=IDX_OFFLINE_USS_PRELOAD, tableName=OFFLINE_USER_SESSION		\N	4.8.0	\N	\N	9326092871
map-remove-ri-13.0.0	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.448388	92	EXECUTED	8:d3f4a33f41d960ddacd7e2ef30d126b3	dropForeignKeyConstraint baseTableName=DEFAULT_CLIENT_SCOPE, constraintName=FK_R_DEF_CLI_SCOPE_SCOPE; dropForeignKeyConstraint baseTableName=CLIENT_SCOPE_CLIENT, constraintName=FK_C_CLI_SCOPE_SCOPE; dropForeignKeyConstraint baseTableName=CLIENT_SC...		\N	4.8.0	\N	\N	9326092871
13.0.0-KEYCLOAK-17992-drop-constraints	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.448708	93	MARK_RAN	8:1284a27fbd049d65831cb6fc07c8a783	dropPrimaryKey constraintName=C_CLI_SCOPE_BIND, tableName=CLIENT_SCOPE_CLIENT; dropIndex indexName=IDX_CLSCOPE_CL, tableName=CLIENT_SCOPE_CLIENT; dropIndex indexName=IDX_CL_CLSCOPE, tableName=CLIENT_SCOPE_CLIENT		\N	4.8.0	\N	\N	9326092871
13.0.0-increase-column-size-federated	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.449954	94	EXECUTED	8:9d11b619db2ae27c25853b8a37cd0dea	modifyDataType columnName=CLIENT_ID, tableName=CLIENT_SCOPE_CLIENT; modifyDataType columnName=SCOPE_ID, tableName=CLIENT_SCOPE_CLIENT		\N	4.8.0	\N	\N	9326092871
13.0.0-KEYCLOAK-17992-recreate-constraints	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.450285	95	MARK_RAN	8:3002bb3997451bb9e8bac5c5cd8d6327	addNotNullConstraint columnName=CLIENT_ID, tableName=CLIENT_SCOPE_CLIENT; addNotNullConstraint columnName=SCOPE_ID, tableName=CLIENT_SCOPE_CLIENT; addPrimaryKey constraintName=C_CLI_SCOPE_BIND, tableName=CLIENT_SCOPE_CLIENT; createIndex indexName=...		\N	4.8.0	\N	\N	9326092871
json-string-accomodation-fixed	keycloak	META-INF/jpa-changelog-13.0.0.xml	2024-06-25 14:34:53.451188	96	EXECUTED	8:dfbee0d6237a23ef4ccbb7a4e063c163	addColumn tableName=REALM_ATTRIBUTE; update tableName=REALM_ATTRIBUTE; dropColumn columnName=VALUE, tableName=REALM_ATTRIBUTE; renameColumn newColumnName=VALUE, oldColumnName=VALUE_NEW, tableName=REALM_ATTRIBUTE		\N	4.8.0	\N	\N	9326092871
14.0.0-KEYCLOAK-11019	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.452137	97	EXECUTED	8:75f3e372df18d38c62734eebb986b960	createIndex indexName=IDX_OFFLINE_CSS_PRELOAD, tableName=OFFLINE_CLIENT_SESSION; createIndex indexName=IDX_OFFLINE_USS_BY_USER, tableName=OFFLINE_USER_SESSION; createIndex indexName=IDX_OFFLINE_USS_BY_USERSESS, tableName=OFFLINE_USER_SESSION		\N	4.8.0	\N	\N	9326092871
14.0.0-KEYCLOAK-18286	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.452496	98	MARK_RAN	8:7fee73eddf84a6035691512c85637eef	createIndex indexName=IDX_CLIENT_ATT_BY_NAME_VALUE, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
14.0.0-KEYCLOAK-18286-revert	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.456304	99	MARK_RAN	8:7a11134ab12820f999fbf3bb13c3adc8	dropIndex indexName=IDX_CLIENT_ATT_BY_NAME_VALUE, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
14.0.0-KEYCLOAK-18286-supported-dbs	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.457	100	EXECUTED	8:c0f6eaac1f3be773ffe54cb5b8482b70	createIndex indexName=IDX_CLIENT_ATT_BY_NAME_VALUE, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
14.0.0-KEYCLOAK-18286-unsupported-dbs	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.457328	101	MARK_RAN	8:18186f0008b86e0f0f49b0c4d0e842ac	createIndex indexName=IDX_CLIENT_ATT_BY_NAME_VALUE, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
KEYCLOAK-17267-add-index-to-user-attributes	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.457907	102	EXECUTED	8:09c2780bcb23b310a7019d217dc7b433	createIndex indexName=IDX_USER_ATTRIBUTE_NAME, tableName=USER_ATTRIBUTE		\N	4.8.0	\N	\N	9326092871
KEYCLOAK-18146-add-saml-art-binding-identifier	keycloak	META-INF/jpa-changelog-14.0.0.xml	2024-06-25 14:34:53.458972	103	EXECUTED	8:276a44955eab693c970a42880197fff2	customChange		\N	4.8.0	\N	\N	9326092871
15.0.0-KEYCLOAK-18467	keycloak	META-INF/jpa-changelog-15.0.0.xml	2024-06-25 14:34:53.460136	104	EXECUTED	8:ba8ee3b694d043f2bfc1a1079d0760d7	addColumn tableName=REALM_LOCALIZATIONS; update tableName=REALM_LOCALIZATIONS; dropColumn columnName=TEXTS, tableName=REALM_LOCALIZATIONS; renameColumn newColumnName=TEXTS, oldColumnName=TEXTS_NEW, tableName=REALM_LOCALIZATIONS; addNotNullConstrai...		\N	4.8.0	\N	\N	9326092871
17.0.0-9562	keycloak	META-INF/jpa-changelog-17.0.0.xml	2024-06-25 14:34:53.460757	105	EXECUTED	8:5e06b1d75f5d17685485e610c2851b17	createIndex indexName=IDX_USER_SERVICE_ACCOUNT, tableName=USER_ENTITY		\N	4.8.0	\N	\N	9326092871
18.0.0-10625-IDX_ADMIN_EVENT_TIME	keycloak	META-INF/jpa-changelog-18.0.0.xml	2024-06-25 14:34:53.461286	106	EXECUTED	8:4b80546c1dc550ac552ee7b24a4ab7c0	createIndex indexName=IDX_ADMIN_EVENT_TIME, tableName=ADMIN_EVENT_ENTITY		\N	4.8.0	\N	\N	9326092871
19.0.0-10135	keycloak	META-INF/jpa-changelog-19.0.0.xml	2024-06-25 14:34:53.462172	107	EXECUTED	8:af510cd1bb2ab6339c45372f3e491696	customChange		\N	4.8.0	\N	\N	9326092871
20.0.0-12964-supported-dbs	keycloak	META-INF/jpa-changelog-20.0.0.xml	2024-06-25 14:34:53.462817	108	EXECUTED	8:05c99fc610845ef66ee812b7921af0ef	createIndex indexName=IDX_GROUP_ATT_BY_NAME_VALUE, tableName=GROUP_ATTRIBUTE		\N	4.8.0	\N	\N	9326092871
20.0.0-12964-unsupported-dbs	keycloak	META-INF/jpa-changelog-20.0.0.xml	2024-06-25 14:34:53.463126	109	MARK_RAN	8:314e803baf2f1ec315b3464e398b8247	createIndex indexName=IDX_GROUP_ATT_BY_NAME_VALUE, tableName=GROUP_ATTRIBUTE		\N	4.8.0	\N	\N	9326092871
client-attributes-string-accomodation-fixed	keycloak	META-INF/jpa-changelog-20.0.0.xml	2024-06-25 14:34:53.464175	110	EXECUTED	8:56e4677e7e12556f70b604c573840100	addColumn tableName=CLIENT_ATTRIBUTES; update tableName=CLIENT_ATTRIBUTES; dropColumn columnName=VALUE, tableName=CLIENT_ATTRIBUTES; renameColumn newColumnName=VALUE, oldColumnName=VALUE_NEW, tableName=CLIENT_ATTRIBUTES		\N	4.8.0	\N	\N	9326092871
\.


--
-- Data for Name: databasechangeloglock; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.databasechangeloglock (id, locked, lockgranted, lockedby) FROM stdin;
1	f	\N	\N
1000	f	\N	\N
1001	f	\N	\N
\.


--
-- Data for Name: default_client_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.default_client_scope (realm_id, scope_id, default_scope) FROM stdin;
98749fe9-5c8f-4d46-b973-16664c916f0f	d3865983-63c5-4a61-8790-79a5f862056d	f
98749fe9-5c8f-4d46-b973-16664c916f0f	03ec7dfe-784f-4953-b4a5-7316c1193a66	t
98749fe9-5c8f-4d46-b973-16664c916f0f	76365f1a-cc60-4b37-affe-c5c494cf2f47	t
98749fe9-5c8f-4d46-b973-16664c916f0f	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8	t
98749fe9-5c8f-4d46-b973-16664c916f0f	f1d89373-70ef-400b-a3bb-aafbcfa4326b	f
98749fe9-5c8f-4d46-b973-16664c916f0f	bdf9e8fc-7774-4216-ae5b-9c955ec11853	f
98749fe9-5c8f-4d46-b973-16664c916f0f	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795	t
98749fe9-5c8f-4d46-b973-16664c916f0f	66485ea3-33cc-45e2-9db8-4b7632fcbcfd	t
98749fe9-5c8f-4d46-b973-16664c916f0f	7bab9854-2b99-4050-b410-81a6a09c7832	f
98749fe9-5c8f-4d46-b973-16664c916f0f	016e49db-14f9-4d10-be20-4502b2a84a27	t
cwbi	662e8870-a4c7-431a-9b40-56a3c4cbb6ea	t
cwbi	268fb6b2-58c0-44f2-ae22-a8baf2236e18	t
cwbi	98992cfc-118d-4c64-976b-7ba01f0976a5	t
cwbi	7d5c5b91-9cc3-467b-aced-a25db29c2576	t
cwbi	cdc2e818-4856-4688-b54a-03a5d08e6a1d	t
cwbi	c983959d-26d1-413e-9343-f5bad7dabc51	t
cwbi	8af9e1dc-6902-4244-a87e-b40749f6a92d	f
cwbi	0542ff9c-210f-4eb3-b62e-fa7272032823	f
cwbi	2648bc85-15fc-4968-b6b5-8b9743c8cfad	f
cwbi	2b3b3db7-5772-4d81-a35a-742ea21a95e6	f
cwbi	5286fee9-6cda-4a94-aba0-dffa0a5c2e8f	t
cwbi	9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0	t
cwbi	b0a33b5f-7c9a-4d59-9602-855dfb2a0b92	t
\.


--
-- Data for Name: event_entity; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.event_entity (id, client_id, details_json, error, ip_address, realm_id, session_id, event_time, type, user_id) FROM stdin;
47c9859e-e12a-4812-8880-af386b3c0670	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","username":"smith_dennis@bah.com"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326269983	LOGIN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
4758dbfe-2aaa-498b-abd1-8f90946b396c	midas	{"token_id":"41ca7497-83bf-40b7-968d-7d76b1c3564d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"782def81-3de5-4c3c-bd65-94e8925fd786","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","client_auth_method":"client-secret"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326270292	CODE_TO_TOKEN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
d5135a26-a6eb-4069-920f-31a47ec62ad5	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","response_mode":"fragment","username":"smith_dennis@bah.com"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326335315	LOGIN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
f0b9d26f-454f-45ea-b213-bec42bbd5535	midas	{"token_id":"ae22e104-5e66-490f-bc28-3a7a68df6366","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"43345a33-f5f6-4110-b67a-fb236d78f0a4","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","client_auth_method":"client-secret"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326335337	CODE_TO_TOKEN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
d39bad0c-3397-4ea6-bfc3-8ddbf90aaeab	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","response_mode":"fragment","username":"smith_dennis@bah.com"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326385632	LOGIN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
cca63b92-098a-46fe-a91e-5fa70191d948	midas	{"token_id":"d4839301-7d3c-4c13-a160-08226454cf7d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"5116195b-7317-4bb6-810f-159d0c9c6b83","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","client_auth_method":"client-secret"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326385661	CODE_TO_TOKEN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
feec4233-466a-468e-966c-4571673f1956	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","response_mode":"fragment","username":"smith_dennis@bah.com"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326404018	LOGIN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
2ab2d731-92aa-4f5d-bc0e-7d40f0ce9b41	midas	{"token_id":"6932a4b0-7bcb-42a0-b300-8d10dbc1d2d9","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"985180bf-df49-4271-9169-3bf56dce2393","code_id":"55ca46c8-5e92-4c78-8cec-a8f753079c63","client_auth_method":"client-secret"}	\N	172.21.0.1	cwbi	55ca46c8-5e92-4c78-8cec-a8f753079c63	1719326404043	CODE_TO_TOKEN	1c6decce-7042-4a2c-9880-ac7c5bc27b8c
3aade133-5064-498e-9e33-4bad0fedcc14	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"3bebb32a-8533-4a52-8120-c6db5e44ec15","username":"test"}	\N	172.24.0.1	cwbi	3bebb32a-8533-4a52-8120-c6db5e44ec15	1719328122561	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7eb08eaf-f4cc-4f46-a059-32a8d1151afb	midas	{"token_id":"0c100ffa-0378-4d90-a508-1d8df758a233","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"2d10429d-f166-404a-b9f3-762fa7d35b2e","code_id":"3bebb32a-8533-4a52-8120-c6db5e44ec15","client_auth_method":"client-secret"}	\N	172.24.0.1	cwbi	3bebb32a-8533-4a52-8120-c6db5e44ec15	1719328123004	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ce5f8c29-9ca0-46a7-b419-397ddd5aefd5	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","username":"admin"}	user_not_found	172.27.0.1	cwbi	\N	1719335908711	LOGIN_ERROR	\N
11e77552-4da7-46cd-9bb2-e9b5bce85097	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","username":"test"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719335922065	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
97f5faf7-a04a-4255-89f3-d6c2583321e8	midas	{"token_id":"4d2fca61-410c-4f95-bc01-5fdde6c91c71","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"045fb238-02a3-4cd0-a358-f18910b2ad92","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719335922999	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ad313b07-941d-4173-b2ee-c7e7e5e719fd	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719335996334	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
16c96cec-feed-4510-9fba-c74467e4403e	midas	{"token_id":"365d1e8c-19c5-4263-8468-8b3c547954dd","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"d82156f7-8b66-4b04-b09b-f69596322ea0","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719335996367	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
38ef817d-18d5-4c4e-a0ce-6cfa9f301cd0	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336579477	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ec44b4de-8b4c-437b-a8bc-d7b0a17b1023	midas	{"token_id":"40d0e3b0-c104-4efb-b112-33f67e5652a4","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"8645fd13-d4e8-4d6b-af24-8cc65c526258","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336579518	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
aca63458-86a4-44ef-a35d-f81c358b6906	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","username":"test"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428820063	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
56553f2b-9e34-4a9c-a3bf-0d717dfac373	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336583718	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a76c3a3e-56c3-4ee0-a56c-644d5e934467	midas	{"token_id":"0b71088c-8c4d-4821-8a51-bddc9a95ed93","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"6cfcccef-66d6-41e6-a043-7e9b063f502c","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336583761	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7aa18916-8622-453f-b2fb-fc77d09aa499	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336594151	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
633c382f-37e3-4be6-a8ab-f2c06d559367	midas	{"token_id":"9a2aed71-45cf-4d88-bafb-15657cb11faa","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"c2642527-7126-497b-afd9-b445418810a6","code_id":"1b4bcefa-eed7-4798-8511-a09dfad7ace5","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	1b4bcefa-eed7-4798-8511-a09dfad7ace5	1719336594192	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
01da9dc0-62fe-4f3e-a8d1-686bc1cce3d8	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","username":"test"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719336683943	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ddd3318b-79e1-4aec-819b-4e7c3d51e03f	midas	{"token_id":"8825600b-6400-46d1-95c0-928b90a1429d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"050fee0a-4691-4657-be26-546d5cd32942","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719336684825	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
fe435fa8-bcde-4984-b428-e90f12bc9069	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337240864	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9b72090b-1591-421f-864f-306747e11415	midas	{"token_id":"266c4441-3afe-4fac-8464-ac7b8e3a978b","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"7b65bb9f-333e-4696-8080-7a40053e305c","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337240919	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7f3a80b4-7f37-403d-8801-8a6e9af924b1	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337281319	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
49a752f7-1ed0-4d6d-abc5-113dbe3d3017	midas	{"token_id":"35a4ce75-0fe9-48c0-ad0e-1d90419b521f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"69e00225-e2d8-4dbd-a977-fb58fc7150fe","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337281361	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
75840f89-c6bb-4e33-9f51-ed1a971eb421	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","response_mode":"fragment","username":"test"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337502911	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3d191c76-a085-4d9d-8933-dbb9aa8bbbf7	midas	{"token_id":"79113ad0-b2eb-4d26-adbc-3006e0ab172e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"cf7265b2-73ec-41a6-aaf6-c70173a3bce0","code_id":"17483e05-1fe0-49f0-adb3-d6af28344c6a","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	17483e05-1fe0-49f0-adb3-d6af28344c6a	1719337502954	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
62ed27ca-8573-42dc-a324-e0d9435d2630	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"cb7774ac-014d-4b8a-ab8c-c01247cce86f","username":"test"}	\N	172.27.0.1	cwbi	cb7774ac-014d-4b8a-ab8c-c01247cce86f	1719337577407	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
2c8e03dd-2171-4ccc-a8f6-817649536c1f	midas	{"token_id":"43630672-03e8-418b-8b0e-d227c1dc9fde","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"5954687a-7b09-4615-8af5-d7014adf96a4","code_id":"cb7774ac-014d-4b8a-ab8c-c01247cce86f","client_auth_method":"client-secret"}	\N	172.27.0.1	cwbi	cb7774ac-014d-4b8a-ab8c-c01247cce86f	1719337578269	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
0933e724-687b-4f6d-b0b0-f2ab77cbde54	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426708744	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
87a65f6a-1343-4208-98cd-8917b59b3e06	midas	{"token_id":"ed4ecb63-5e3a-40e1-8235-30d9f4ab80f9","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"b1cd24cc-e48b-4d1a-8e41-4bb2f88647f0","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426709194	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a7bb6228-5a1c-42d3-a8ff-b64021175850	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426718569	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f956a9fb-a146-43a8-8a09-958af66f4790	midas	{"token_id":"60e864c2-53c6-4b28-b913-4e570cc432d5","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"399a73d2-3742-47f0-b078-19e1dea30f19","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426718642	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
5fce2440-836a-4b1d-898f-61bf8ae7b45e	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426724933	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
b48fb9fb-2496-4b89-b820-42c45054b0af	midas	{"token_id":"7ac4ac7a-9c21-46a6-9314-85ce10cd267d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"384eee9f-5a5b-488c-9229-3b72a1e4de6b","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426724981	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
5e17979e-01af-4588-8516-fbe149a1d628	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426739487	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
8bbf7e30-33e6-4050-9706-cf6ec28649ce	midas	{"token_id":"5b40d434-b10b-409f-9460-c2681ff342d4","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"540edb49-ef44-4b15-9887-c5f1eef79a58","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426739533	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c93747a2-e03b-4332-a70b-03c82cd01290	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426746694	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
aea0cdc9-50ba-4e20-af58-e951d098f5c8	midas	{"token_id":"14309ef3-7f2f-4b12-a15f-2d9ce4c95f6d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"201d2123-4495-4c3f-ab88-e2824b66ca53","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719426746733	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
792c0407-fe21-413f-baae-be505efbba59	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719427679782	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3cfc8a69-eb17-4022-8579-271972fff6ed	midas	{"token_id":"f4370d49-1e0c-4eb0-9241-300d19ed5845","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"c55e9da6-9e95-455a-bd56-846d107ddd98","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719427679831	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3fe60b55-20d3-4f59-88f8-0e8505a7121c	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"lambert.anthony.m.2"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428545717	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
abc5dbd4-d985-42ff-b297-7ed55c0329e8	midas	{"token_id":"67d0d6a6-ff8e-4530-971f-ccb95f0700c1","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"9e806450-ee31-43f7-ba86-2b45f43a17de","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428545751	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
51ad30c6-df6b-4e25-83b0-3c510cae65fb	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"lambert.anthony.m.2"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428551878	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
795cd68b-2f57-4c60-94b3-ec660fb63513	midas	{"token_id":"54854a35-1b99-427b-bcdd-db53b899618e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"b0222f6c-6997-469f-91d7-a8d9964f0306","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428551922	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3894a710-f1ee-4935-bab3-7b5d145d7fac	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","response_mode":"fragment","username":"lambert.anthony.m.2"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428725626	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c1b4978f-db0b-42e0-9996-f7209ba515c9	midas	{"token_id":"e545a798-3473-466e-8df6-b470b9cbb259","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3a79ab28-8a78-44eb-85d7-93133ffd9bf7","code_id":"7392695a-acc0-4ab0-9744-f5ba1502c1de","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	7392695a-acc0-4ab0-9744-f5ba1502c1de	1719428725694	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
78dba2d6-ac91-49be-b43b-56443e790864	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","username":"test"}	user_not_found	172.28.0.1	cwbi	\N	1719428759822	LOGIN_ERROR	\N
e282fa0c-afc4-4666-a7f7-eb3c021b3519	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","username":"AnthonyLambert"}	user_not_found	172.28.0.1	cwbi	\N	1719428768390	LOGIN_ERROR	\N
e4dfe693-a048-42ad-b670-3b1e99e8c56d	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","username":"anthonylambert"}	user_not_found	172.28.0.1	cwbi	\N	1719428778880	LOGIN_ERROR	\N
174d770a-2d59-4086-8237-d477ad7abd1a	midas	{"token_id":"7fb976c2-0bfd-44a5-b4a8-9d46e4bca397","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"7f0e2bae-78a4-46d3-b74e-0266828a85ae","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428820965	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
79a933d8-0e94-4b9d-8a0d-3f2edb2070b0	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428832975	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
85712b3d-cf72-445a-ad31-e8d0de9f6512	midas	{"token_id":"050fce80-1214-46e8-89f6-e17254db6018","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"0bf36404-917f-47c1-a722-34473cafdf45","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428833032	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a612d22f-2340-48d8-9588-c9b8a1817467	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428905095	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
377c4f8f-5593-4e89-a4fe-1021b95671ba	midas	{"token_id":"7bfbe320-49c4-4e15-b37c-4272f372202e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"f1f68ad8-e789-46ed-bfb6-95239979791e","code_id":"db1641af-1aea-4402-8741-dea01f2cf470","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	db1641af-1aea-4402-8741-dea01f2cf470	1719428905148	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
4884def3-e5d6-4757-a407-5a73fdaebf23	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"f8ff6fb6-ba57-49b6-bfa4-52424e978971","username":"test"}	\N	172.28.0.1	cwbi	f8ff6fb6-ba57-49b6-bfa4-52424e978971	1719429338440	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
8d481359-4dcc-4398-9fff-b89228ce03be	midas	{"token_id":"6e73f969-924e-4aba-91f6-8d81d8249952","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"6178b7b2-9a8c-4c8f-b580-cfcc81c86528","code_id":"f8ff6fb6-ba57-49b6-bfa4-52424e978971","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	f8ff6fb6-ba57-49b6-bfa4-52424e978971	1719429339343	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
64f8bf30-7354-4e7a-9578-cb40798f0aac	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"f8ff6fb6-ba57-49b6-bfa4-52424e978971","response_mode":"fragment","username":"test"}	\N	172.28.0.1	cwbi	f8ff6fb6-ba57-49b6-bfa4-52424e978971	1719429536508	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
af718add-a959-4164-899d-c8ecb50a2337	midas	{"token_id":"2ce361ae-75b7-4275-9f8a-daaa0da94f97","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"32655748-bf9d-466e-8861-3516ff4263e5","code_id":"f8ff6fb6-ba57-49b6-bfa4-52424e978971","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	f8ff6fb6-ba57-49b6-bfa4-52424e978971	1719429536562	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
e2023d13-dfc8-456f-b521-d0d920c18cec	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"60caa530-9918-42e6-82c7-8b5f132e2086","username":"test"}	\N	172.28.0.1	cwbi	60caa530-9918-42e6-82c7-8b5f132e2086	1719429555434	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
00f193f1-e9b3-48a3-8286-4fd17c14ad67	midas	{"token_id":"fe1cf5c4-1676-4ac4-b576-8c1781469fac","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"57e58153-f972-4a2c-b5bb-1bc0c4fdd84c","code_id":"60caa530-9918-42e6-82c7-8b5f132e2086","client_auth_method":"client-secret"}	\N	172.28.0.1	cwbi	60caa530-9918-42e6-82c7-8b5f132e2086	1719429556306	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7f8886de-f561-472b-83ad-285fa27bacfa	\N	null	session_expired	172.29.0.1	cwbi	\N	1719431414675	LOGOUT_ERROR	\N
00db9379-9298-41a2-9d6c-04274ed88c49	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"9061ef37-a53e-4cb8-8491-4d805dc80282","username":"test"}	\N	172.29.0.1	cwbi	9061ef37-a53e-4cb8-8491-4d805dc80282	1719431543213	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f603e456-550c-4b3d-bb26-8a87aee59d08	midas	{"token_id":"a3920699-7eba-4f31-a578-8172cdd29ee7","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"6a262f6d-05a3-40a2-9b38-65a7db24e04e","code_id":"9061ef37-a53e-4cb8-8491-4d805dc80282","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	9061ef37-a53e-4cb8-8491-4d805dc80282	1719431544159	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ef5c8222-1610-41d9-9cd8-1bf19da56259	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"9061ef37-a53e-4cb8-8491-4d805dc80282","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	9061ef37-a53e-4cb8-8491-4d805dc80282	1719431576525	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a657fe78-6d0d-42e1-ad48-90cd1851a42a	midas	{"token_id":"50291355-58a7-4ad9-8029-083e71ccf8f7","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"9fbb00eb-80ce-4c8e-bfaf-d293a60cd11e","code_id":"9061ef37-a53e-4cb8-8491-4d805dc80282","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	9061ef37-a53e-4cb8-8491-4d805dc80282	1719431576585	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
682f3021-b079-40da-91df-46ca131ab2bf	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","username":"test"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431620152	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
769c7d0d-9f04-4b2d-842d-c940ac355548	midas	{"token_id":"69ad7d7f-07f9-4864-883f-b9bde16b1939","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"ecbc3756-dd5e-4ce7-b932-2e16d9c5cea6","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431621105	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c8fc7cd3-da13-4e81-84f7-b7747148f061	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431660988	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
2b24b242-7d0a-47e5-95d5-000c7166ed14	midas	{"token_id":"a793c465-0c0a-4fb2-9ce8-d848a69ddb3b","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"cb853d42-25bc-43a4-8239-57662c7aca49","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431661036	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
99e67057-6fba-464c-baa6-edc61444eb91	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431665861	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
1f62af81-3bb6-48c7-bce7-60690e727510	midas	{"token_id":"00c65b26-df3f-4381-b528-0efc2afef914","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"00f98d52-3dce-45dd-ab1b-3a486de78352","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431665925	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
e21b3a9e-f1df-4fa2-b4e4-ff0f1d1393a0	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431694480	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
8a380b0e-5918-4267-80c3-3d92812c5197	midas	{"token_id":"c0cf05ba-fe95-434b-8f49-2f497f11a44c","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"ac7931a9-858f-4f16-b81f-021c05c90ca3","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431694538	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
36dfaca1-9367-4958-b107-c5c5bf57bdbd	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431872347	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9f6308ff-5998-46df-a797-cd9cbc8dc9f4	midas	{"token_id":"40682f23-3a30-4d5b-924e-ae65ce133463","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"bbb555f4-dff9-49fa-82ad-6d019278262b","code_id":"ebdde03c-195a-4b4f-947a-429e15265dd2","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	ebdde03c-195a-4b4f-947a-429e15265dd2	1719431872412	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
08cbe03a-9450-4ef4-ab02-f04b0830c8a1	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"33560fda-4c80-4716-a0df-45b4e3d5606b","username":"test"}	\N	172.29.0.1	cwbi	33560fda-4c80-4716-a0df-45b4e3d5606b	1719431990734	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
97431830-0b96-4fe8-872f-3be121ede6aa	midas	{"token_id":"86299821-b27a-4a26-a274-053a514d2619","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"a547ab17-c77c-44d0-8064-813554880728","code_id":"33560fda-4c80-4716-a0df-45b4e3d5606b","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	33560fda-4c80-4716-a0df-45b4e3d5606b	1719431991647	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
0ec7e6ff-756e-48f3-a22b-2e0017b69060	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"d67fbc88-05f3-460b-bab3-a1df3b6c7939","username":"test"}	\N	172.29.0.1	cwbi	d67fbc88-05f3-460b-bab3-a1df3b6c7939	1719432078862	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
6b37eb7b-ede3-43f5-82f8-aae4e5522e4e	midas	{"token_id":"427429dc-5334-4575-baaa-5b7e238a04cc","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"89e81a25-54e0-44db-9761-5a3b4e00a2b9","code_id":"d67fbc88-05f3-460b-bab3-a1df3b6c7939","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d67fbc88-05f3-460b-bab3-a1df3b6c7939	1719432079796	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7d170853-d8e1-4b2c-be92-227e3edb80b9	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432367589	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
897a8610-4398-4eb0-b5ba-2cd588b5f537	midas	{"token_id":"d597b975-bda0-4932-9329-374ed1e64e51","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"6cc29a3d-3ef4-4529-8899-23d3a7add3a6","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432368514	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
b2998f2f-dfc4-40c2-976b-543f80520973	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432380008	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
27673131-bdfe-4cd4-b8f8-bbbae538a0f9	midas	{"token_id":"e29e8a48-7a42-41b9-9254-31cae7d0dc40","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"7392d55d-b933-49ee-bad8-105713013044","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432380055	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
b3631f03-8eff-42b7-8e1e-448d75454cc7	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432385551	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
558b3452-b49c-4768-af43-090b533d72de	midas	{"token_id":"c6608a65-2ee7-4794-aaff-a60d425ef479","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"c711222a-b7c8-419c-9ee1-485f77f85f44","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432385607	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
4889a6ec-06a9-4ee1-bb97-f1f0dece4dfd	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432398018	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ec03fea5-0dd6-4ced-82a1-59e8f057df98	midas	{"token_id":"a6a25af3-c7ba-483c-a580-7ebc3527cc4e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"34036d15-e6ea-4042-b695-630f0c32b27a","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432398069	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
4bd59fba-ff10-4e36-80a0-fc6bcf3b7188	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432401231	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
1f47e107-767d-481d-9354-610319098425	midas	{"token_id":"0b593308-c9a7-4a53-8fa6-64376e405004","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"fa5d9ddc-1f53-42d5-851d-0035904b8492","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432401303	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f7aebbbc-782e-4567-84b0-fdfd7bf4c673	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","response_mode":"fragment","username":"test"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432402838	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
0409d568-efe1-4b9e-b9c7-3540b4a0b49f	midas	{"token_id":"ac967f12-b925-4297-ba07-e058ed71f1c8","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"54d9ed37-90b3-4a37-944f-1e1b25ea7371","code_id":"d1451e6c-8187-4d2c-b06e-41be46e4f7cf","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d1451e6c-8187-4d2c-b06e-41be46e4f7cf	1719432402895	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
8bf13cd7-b2b4-4c17-9602-be020f0d9349	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"d99af11b-8606-4ddf-a8df-07acadb7b19a","username":"test"}	\N	172.29.0.1	cwbi	d99af11b-8606-4ddf-a8df-07acadb7b19a	1719432426611	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
b4614a34-12f4-4d23-8e74-0c29e985260c	midas	{"token_id":"b8d41880-586b-4d82-b9b1-33cdf996ded1","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"8c8a0e0c-c38a-4b7c-a4fa-c3eecc7fd601","code_id":"d99af11b-8606-4ddf-a8df-07acadb7b19a","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	d99af11b-8606-4ddf-a8df-07acadb7b19a	1719432427544	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
5b4e849c-b35e-4cb7-90f9-7d4ebf0f6856	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"fafcfc4d-5316-4501-a6e2-7b16b55555b4","username":"test"}	\N	172.29.0.1	cwbi	fafcfc4d-5316-4501-a6e2-7b16b55555b4	1719432594649	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
cf48c2a8-b9cd-4a5a-804d-0566ae358b37	midas	{"token_id":"a2a30d77-b6dc-403c-a1ce-fbdf5440a4db","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"35308cae-9157-4018-9407-d070cff9885a","code_id":"fafcfc4d-5316-4501-a6e2-7b16b55555b4","client_auth_method":"client-secret"}	\N	172.29.0.1	cwbi	fafcfc4d-5316-4501-a6e2-7b16b55555b4	1719432595543	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
d0b02f30-ac3e-4e2a-9afa-53e427c2f1a9	\N	null	session_expired	172.31.0.1	cwbi	\N	1719433495247	LOGOUT_ERROR	\N
7514c25c-d292-42e5-88ab-04865ab41aa3	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"c1d9083b-4df4-4644-aaaf-84f159abf663","username":"test"}	\N	172.31.0.1	cwbi	c1d9083b-4df4-4644-aaaf-84f159abf663	1719433641725	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
973c3eda-9df7-40c7-8bfc-6e19d58b8562	midas	{"token_id":"80c068da-5cc6-40eb-b349-a3c975d6919c","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"2127b86e-cbbf-4e06-8ccc-6887bb6b5c8e","code_id":"c1d9083b-4df4-4644-aaaf-84f159abf663","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	c1d9083b-4df4-4644-aaaf-84f159abf663	1719433642868	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
cdd1aa46-171f-4c9f-858a-96cb011ecdd4	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"c1d9083b-4df4-4644-aaaf-84f159abf663","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	c1d9083b-4df4-4644-aaaf-84f159abf663	1719433658666	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
1e17f743-b74d-4746-aeea-b9a7ccd14707	midas	{"token_id":"c464b7ab-bfba-4ebe-8d54-a00a61bbadcb","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"9640190f-07cd-4aab-badf-44e0c9e78222","code_id":"c1d9083b-4df4-4644-aaaf-84f159abf663","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	c1d9083b-4df4-4644-aaaf-84f159abf663	1719433658736	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
031bcd80-bd0a-4c7a-8f0e-471a06c79801	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"a1552826-617b-4d4a-9aab-7759bf42de70","username":"test"}	\N	172.31.0.1	cwbi	a1552826-617b-4d4a-9aab-7759bf42de70	1719433867486	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
2a622433-c852-42de-ad82-b7a44f651c22	midas	{"token_id":"4f8e0ede-fa35-434c-bd0a-8a05f75c67eb","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"d4eeda47-3fea-4b96-a70a-448020efcb01","code_id":"a1552826-617b-4d4a-9aab-7759bf42de70","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	a1552826-617b-4d4a-9aab-7759bf42de70	1719433868526	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
dc323c69-e177-47eb-9b16-ee7dfff5be39	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"a1552826-617b-4d4a-9aab-7759bf42de70","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	a1552826-617b-4d4a-9aab-7759bf42de70	1719434093170	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
1326f514-eeb7-4a09-bacb-7d7857355958	midas	{"token_id":"114ebebd-30c8-41d6-83d2-df06e1c57a50","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3ac184a0-f344-4d5d-a3d8-5fa791a6870c","code_id":"a1552826-617b-4d4a-9aab-7759bf42de70","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	a1552826-617b-4d4a-9aab-7759bf42de70	1719434093244	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
307251b7-d59e-4be2-afbd-31951a1805aa	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434222200	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3cc55d26-c8f3-49c0-a0db-344b3275cb2c	midas	{"token_id":"3bc351bd-a1c4-4665-a914-bc7c6fd030d5","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"8f57d3a2-64a5-4334-95a1-7564d3c9c2e3","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434223266	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a1f353d6-e7ad-464d-9881-18911d8fe090	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434452404	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a61970ed-cfb4-4ad9-b9aa-22b3e234a7d3	midas	{"token_id":"36ef319d-53ee-4a47-8966-3b1e925fd44d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"f7d0d5ce-9b45-4990-ad7e-84b60ea29403","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434452458	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
6fa0c11c-33cf-416c-98d9-27ce4fd30845	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434459453	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
ba18ba5c-a38a-42e7-ba44-7c9468a81eac	midas	{"token_id":"d55ddb47-f8e0-46d5-98b0-eb0336d1eacd","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"09c2c1c4-7e62-4bbe-aeb7-15a0df5a6827","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434459517	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
bb6918e1-5837-4a5e-b130-d8ebaae39f42	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434461096	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7c3b3b10-5d67-4042-923e-abe15f21dab4	midas	{"token_id":"a872d2d0-165e-40da-85be-b470f6bd6a96","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"94a513a2-4ade-400d-a6a0-978439ece29e","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434461168	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
e7a4d9fd-11e5-45c7-91cd-38c4db3fed87	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434683214	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
e35717f2-802f-4ca8-b691-17969bb17fd0	midas	{"token_id":"91ad450b-d77e-4a43-b9e5-eb5325c19769","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"e80bd7dc-58f2-4b04-afc7-196e4aaa29f5","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434683293	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
2a93e31c-66a4-4759-bbf9-af6d7cc9b26a	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434831025	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
0cbc185a-b88a-4ed5-b1ba-a18f30355632	midas	{"token_id":"95981868-a440-4e2c-abe1-e73e3bd5037d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"0655bc8b-ff93-4f9b-8e6a-d2fc11591ed8","code_id":"52f7a62c-889a-4fe9-8465-39ad9d91a054","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	52f7a62c-889a-4fe9-8465-39ad9d91a054	1719434831091	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a57d9ac6-747b-44b5-b8a0-3e1652bf54e4	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","username":"test"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719434866723	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9d451239-bb78-418e-bdc9-0c794b9ecad7	midas	{"token_id":"a2c2a1e8-fd36-42b2-9750-d5d510297a37","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"eb2fed7d-0dba-4e7c-801c-cf90998b8582","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719434867813	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
63162276-f27e-4176-b800-2994d0c94c37	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719436350976	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7cd28211-8ea6-44ac-940b-c5776ab06c65	midas	{"token_id":"3b1cd142-bf2f-401b-b41b-3e29510388db","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"285e30a9-8e5a-4739-80b5-fa3e2b7b70a8","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719436351031	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
cab6dbca-b8b2-4f2d-bdd7-a1d6ba0b6e41	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719436354587	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
187701f6-16cd-4eb9-b1ef-96966f9a4d1a	midas	{"token_id":"0e1b8054-5a13-466d-8913-40be07a2cdc3","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"719b998b-6f0f-41ad-90ad-0d96b3932673","code_id":"3c823133-5a53-48b1-8d23-4cf78222cf35","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	3c823133-5a53-48b1-8d23-4cf78222cf35	1719436354651	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9ed711b5-f081-4f1a-8076-c2c75fa89b02	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"4adcf531-f515-477a-a645-f987e257d775","username":"test"}	\N	172.31.0.1	cwbi	4adcf531-f515-477a-a645-f987e257d775	1719436368432	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
739cb124-1955-4674-9e99-8d2736e39a4a	midas	{"token_id":"95dd6b57-a9f9-4819-8aaa-caa242b728da","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"e84f614d-a255-48af-af2c-0d7b4c881fa6","code_id":"4adcf531-f515-477a-a645-f987e257d775","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	4adcf531-f515-477a-a645-f987e257d775	1719436369438	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
74859b41-c89f-4679-a584-163c5ab31415	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"4adcf531-f515-477a-a645-f987e257d775","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	4adcf531-f515-477a-a645-f987e257d775	1719436996678	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
e81a0aa5-3934-4a72-8810-698338a8040c	midas	{"token_id":"1cf8cc51-28a4-4d23-b8de-d9778d6ca60f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"d98eaf26-fe0c-48f1-b2f9-6ada06189d20","code_id":"4adcf531-f515-477a-a645-f987e257d775","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	4adcf531-f515-477a-a645-f987e257d775	1719436996731	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
d702ff34-caef-4e6c-a5cc-e083ba81ac9c	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"bc3838f3-1506-4744-8647-534e32890e7a","username":"test"}	\N	172.31.0.1	cwbi	bc3838f3-1506-4744-8647-534e32890e7a	1719437006381	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
900bc142-028b-4c7d-af0a-625305446270	midas	{"token_id":"cc6bc810-6b6e-4261-9529-a1ca42fd1753","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"b38fe3df-9ff3-41df-ba92-209913d44e51","code_id":"bc3838f3-1506-4744-8647-534e32890e7a","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	bc3838f3-1506-4744-8647-534e32890e7a	1719437007331	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a467488c-4123-43d7-883d-004cd1e2e115	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"bc3838f3-1506-4744-8647-534e32890e7a","response_mode":"fragment","username":"test"}	\N	172.31.0.1	cwbi	bc3838f3-1506-4744-8647-534e32890e7a	1719437566691	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f8b29962-81a4-4c08-a670-1924bfa02b43	midas	{"token_id":"3e69485e-6796-40b5-80e1-0b06a5f3cc3e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"a42a97ab-cc35-4b86-9089-4e8400b95e90","code_id":"bc3838f3-1506-4744-8647-534e32890e7a","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	bc3838f3-1506-4744-8647-534e32890e7a	1719437566760	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
11324515-9918-4a73-ba9b-f250cbd63ee6	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"0352165e-8393-433a-94f1-01999b5c3d4a","username":"test"}	\N	172.31.0.1	cwbi	0352165e-8393-433a-94f1-01999b5c3d4a	1719437589193	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9c1f7d72-e7e8-44c3-8761-3528c8fc2d5f	midas	{"token_id":"9b69ec16-5141-4de9-b4a4-5a9c3844a564","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"5f2b8bef-e9d4-4a3b-9adf-0b9f0dad6591","code_id":"0352165e-8393-433a-94f1-01999b5c3d4a","client_auth_method":"client-secret"}	\N	172.31.0.1	cwbi	0352165e-8393-433a-94f1-01999b5c3d4a	1719437590142	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f0b58df5-8873-4e1b-ad12-3893c7c28d08	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","username":"test"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719437773369	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
a1fbac92-7a85-4033-8775-3fa0775150a4	midas	{"token_id":"6d378d07-cd7f-4dcd-baa6-66a0eb5de1df","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"5a08d8d9-f81a-4ec9-934b-65923890ebd5","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719437774468	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
8c470e70-5f7a-4b26-a469-ba5f7197dd82	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","response_mode":"fragment","username":"test"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719437793667	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
63c6f31a-4baa-4d01-a91e-3d373361c7f7	midas	{"token_id":"71e606b2-dec0-4bd0-847e-91121cb48037","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"a1656069-e74d-4247-beb0-c821d8d7a08b","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719437793725	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c34458a2-90a9-49b0-b45b-6c8e1f1d2591	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","response_mode":"fragment","username":"test"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719438057939	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
f7a80588-70f4-4dfc-82aa-38e1079bf6d0	midas	{"token_id":"0e99a780-387c-484e-8a17-2142914a1383","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"55ac4c93-cf33-4bf0-b999-97849ebdfe2f","code_id":"a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	a5ab156c-cd73-493d-9cc6-ef1d9fe0f1b5	1719438057995	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
3eb14a61-7b6f-4a1c-aa74-268c63cc0c04	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"f5380ee8-21d5-40fe-ba2a-3eddcb2c76dd","username":"test"}	\N	192.168.32.1	cwbi	f5380ee8-21d5-40fe-ba2a-3eddcb2c76dd	1719438072045	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9504d569-1b17-437b-bce8-1368088a0db2	midas	{"token_id":"470dda11-96b3-4906-8383-f51038c6a891","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"ed84dc64-0864-469d-9242-af3985567519","code_id":"f5380ee8-21d5-40fe-ba2a-3eddcb2c76dd","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	f5380ee8-21d5-40fe-ba2a-3eddcb2c76dd	1719438073076	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
b9a15197-3517-40c4-81a6-be2d659146dd	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","username":"test"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438705139	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
bbbc978a-51c2-4956-b113-9c209b673970	midas	{"token_id":"816880ea-10d8-4439-b8c8-620e3acf7b33","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"78e54b46-59ca-4cf2-bea2-84daa9d5aafc","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438706147	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
81f907e2-6efe-47e6-a2ec-3f06a50de800	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","response_mode":"fragment","username":"test"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438916745	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
51b63317-443a-45f6-8bd1-78d2eb11ca4b	midas	{"token_id":"e53703ca-544f-442b-b451-553b57d88940","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3d3eeb88-302e-48fc-b221-1eb281338360","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438916798	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
dbe1745b-0091-433f-883b-156c698ad3ee	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","response_mode":"fragment","username":"test"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438937525	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
84cd79e5-865d-4af9-a546-24b0e674cda9	midas	{"token_id":"593cf23c-edb9-4bbf-afc7-a0b0fba0766f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"5a593a48-4006-4bec-89c3-343564577c85","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438937583	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
7013dcdb-681b-4cf1-9400-91cab855f7fd	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","response_mode":"fragment","username":"test"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438968989	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
df0ab200-9d54-4c85-93cf-26b8262fc37c	midas	{"token_id":"2d7d9dae-bae8-4fd8-939f-6e346b79841f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"7d08726a-2dfe-48bb-9a2a-564fe9b8ecaf","code_id":"8b9a63c5-b3ed-4837-8ada-5872ea8bac63","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	8b9a63c5-b3ed-4837-8ada-5872ea8bac63	1719438969040	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9271a41b-e1ad-49dc-8f8a-708b2b41c413	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","username":"nocactest"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719438983325	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
d50b5441-e25c-4351-a07e-c7ac9c4eb5bb	midas	{"token_id":"b23c56c4-5421-4c96-b1bc-52b90ed29e33","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"1eab8fb5-cc44-4b3f-aa8e-e578c3d46b1d","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719438984366	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
d464abd2-51ff-4a47-a7af-9894492c621d	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719439205080	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
6d01cb28-a145-49e2-98ba-6373fefc5bc0	midas	{"token_id":"bbaa849b-f0b9-4149-b3a9-28d7435d7b98","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"b35e8cfb-2615-4f10-b796-a947b63a4d7e","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719439205134	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
e68603c2-4d23-4b3d-a1ea-d67b665f09c0	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719439298012	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
48c3f06c-f8fe-4b10-9c26-33b5f906ef68	midas	{"token_id":"6999ccd3-d32e-40c5-9133-e79c65ffedab","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"7298c349-97c2-4159-a799-1e3ff949c0bf","code_id":"faf2f891-394c-47fe-bfe7-3cd1df2eacee","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	faf2f891-394c-47fe-bfe7-3cd1df2eacee	1719439298073	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
66d135f9-4343-4f96-8160-b5918a82dfd6	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"7d98231f-5bd3-4702-afa3-ed0522af4996","username":"nocactest"}	\N	192.168.32.1	cwbi	7d98231f-5bd3-4702-afa3-ed0522af4996	1719439310399	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
93561f49-3a8c-4e2d-a882-b91df4d58e16	midas	{"token_id":"e1fbb1a0-496a-4f30-be56-e0caccbf5979","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"709a5ab8-148f-469b-888e-6f9f3433b163","code_id":"7d98231f-5bd3-4702-afa3-ed0522af4996","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	7d98231f-5bd3-4702-afa3-ed0522af4996	1719439311361	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
c9e2b3d3-9e25-4cc0-83ff-f2946182864a	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"d07def14-7e28-4b30-9430-5c5ba23ef792","username":"nocactest"}	\N	192.168.32.1	cwbi	d07def14-7e28-4b30-9430-5c5ba23ef792	1719439490428	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
fd0f4bcb-8c8d-4725-8cbd-f21fe511009f	midas	{"token_id":"cbfc435a-a19e-4787-b944-ba09b304b371","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"9a994032-70d3-49c2-b358-7b01e41bdc3b","code_id":"d07def14-7e28-4b30-9430-5c5ba23ef792","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	d07def14-7e28-4b30-9430-5c5ba23ef792	1719439491434	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
7ca9a4b2-4b62-49a9-b837-9d218880b013	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d07def14-7e28-4b30-9430-5c5ba23ef792","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	d07def14-7e28-4b30-9430-5c5ba23ef792	1719440136697	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
4d578fe6-4a40-4741-b396-7597ea8a381e	midas	{"token_id":"1dbf013b-1fe6-4426-839b-308fd3f7f9f1","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"59998ff9-68bc-4882-a801-bea304d2221c","code_id":"d07def14-7e28-4b30-9430-5c5ba23ef792","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	d07def14-7e28-4b30-9430-5c5ba23ef792	1719440136764	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
c11403a5-73eb-4a39-b1c8-929920223a96	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","username":"nocactest"}	invalid_user_credentials	192.168.32.1	cwbi	\N	1719440147442	LOGIN_ERROR	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
237fce60-76d1-4b8a-b8af-dfa19095e41c	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","username":"nocactest"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440149153	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
a8a81a14-4966-4a02-ae52-11f523d1973b	midas	{"token_id":"08c954ef-8add-4f45-b834-30a1eee49f8c","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"ac79f8eb-9011-4908-a801-9696b3827c6d","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440150093	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
f6748340-f7af-42f3-8248-13d5cbd14da9	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440402985	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
cf164cd1-70ff-4d55-a985-2dfc5027e7ed	midas	{"token_id":"431ad36c-ca6d-45e9-a1e3-be4173c1b3ef","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"05db642b-8bb1-4ab1-a501-adf33d216522","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440403046	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
4128e2b3-0732-4ce4-a829-6339fe71b48c	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440471739	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
9f23f2d4-6e71-446f-9603-abd3dfc6b864	midas	{"token_id":"3c2dc02b-46f4-4aac-8a1d-5ab182e22d30","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"2c3ccaf8-b1be-4332-8b35-2fcbee4b8ffd","code_id":"f5c485da-4230-4d9c-8502-c53a5da0c1f3","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	f5c485da-4230-4d9c-8502-c53a5da0c1f3	1719440471791	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
ab131be6-75c0-4b68-89cc-4e2751e6ba18	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"937c5682-87fb-4aa3-a6e6-16c2d7d9d32b","username":"nocactest"}	\N	192.168.32.1	cwbi	937c5682-87fb-4aa3-a6e6-16c2d7d9d32b	1719440490024	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
b59bf466-29f7-46b3-8aa9-5c41935df748	midas	{"token_id":"6ea3c902-43d8-45e6-b181-ec63a1de070f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"89d40729-3ef4-4f68-9636-38d1e395404d","code_id":"937c5682-87fb-4aa3-a6e6-16c2d7d9d32b","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	937c5682-87fb-4aa3-a6e6-16c2d7d9d32b	1719440491044	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
60bdac60-0915-4c27-8f98-bf2751f76789	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"abdbc855-4cc8-42b8-a4e4-90b68763a586","username":"nocactest"}	\N	192.168.32.1	cwbi	abdbc855-4cc8-42b8-a4e4-90b68763a586	1719440621640	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
e100b394-a1c7-4383-b25a-604980ac0765	midas	{"token_id":"7a45c453-2130-4c17-883a-3d313d2d99e1","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"d4365a75-f128-4daa-8495-2453e87b944d","code_id":"abdbc855-4cc8-42b8-a4e4-90b68763a586","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	abdbc855-4cc8-42b8-a4e4-90b68763a586	1719440622649	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
40aa6774-0e51-4557-81fb-ab379549594b	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"abdbc855-4cc8-42b8-a4e4-90b68763a586","response_mode":"fragment","username":"nocactest"}	\N	192.168.32.1	cwbi	abdbc855-4cc8-42b8-a4e4-90b68763a586	1719441001096	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
23eb2cde-7bc5-4e11-aaa6-977040250f3b	midas	{"token_id":"e76fd558-81a1-499f-852b-7a5ad8d89bbc","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"24431144-aa2c-43da-8817-cfdcd37361f5","code_id":"abdbc855-4cc8-42b8-a4e4-90b68763a586","client_auth_method":"client-secret"}	\N	192.168.32.1	cwbi	abdbc855-4cc8-42b8-a4e4-90b68763a586	1719441001150	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
0244000f-7811-4786-b3d9-1ffe18484097	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"161dbceb-bf25-420a-85a8-f560df7beed2","username":"nocactest"}	\N	192.168.48.1	cwbi	161dbceb-bf25-420a-85a8-f560df7beed2	1719442142612	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
2b74a3a5-ddf3-47a3-ac7b-acbb82d4a3d1	midas	{"token_id":"e2a25466-80c3-4d0b-976e-cf61aedc7eaa","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"e9656700-00f2-4062-a882-358164a6e2bd","code_id":"161dbceb-bf25-420a-85a8-f560df7beed2","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	161dbceb-bf25-420a-85a8-f560df7beed2	1719442143675	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
02de765d-1fd0-4b01-aedf-1aa92ac4a408	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"161dbceb-bf25-420a-85a8-f560df7beed2","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	161dbceb-bf25-420a-85a8-f560df7beed2	1719442164568	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
18b99a91-023e-40e8-8100-c82818013135	midas	{"token_id":"60016ad3-5fcc-4c4b-a64f-a01215a80450","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"a99a8089-62c3-4d6b-89be-a304a22c7839","code_id":"161dbceb-bf25-420a-85a8-f560df7beed2","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	161dbceb-bf25-420a-85a8-f560df7beed2	1719442164634	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
e08a2b5e-7f4e-4c68-99ef-7f8922706248	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","username":"nocactest"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719442200917	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
649b66dd-96d9-4878-a254-b2721d84732b	midas	{"token_id":"0ea25384-8514-428c-9231-245753806182","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"742df027-e74a-49e7-b21e-9960d181706c","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719442201836	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
94a5689a-920e-4a7e-b35f-3950ff4ef271	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719442527905	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
2c9c8cdb-32f2-44c0-9f59-bda8ebf1f690	midas	{"token_id":"d77e8feb-8094-4a6c-b0e2-bd5a2d2ba695","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"14196da6-fa2d-4dd3-8c36-e29cd16fa3cc","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719442527973	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
710528c7-6700-4bf0-b787-0568b6d0fe2d	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719443563878	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
67b1c8e1-c490-4ddf-84cd-68a7185a3dc7	midas	{"token_id":"e9a6e3b6-ef9e-4219-85a1-fd2f3110707d","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"ec9f96f8-4721-4957-9e9e-9f2ba2a34ac3","code_id":"607d0edf-7cd3-4d25-8968-56d64c8267b3","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	607d0edf-7cd3-4d25-8968-56d64c8267b3	1719443563944	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
6db23675-b830-4b8d-938f-860e3f13fdac	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","username":"nocactest"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719443581103	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
01190335-c458-4085-b62b-491786bdc226	midas	{"token_id":"6eeba4a7-0c62-45d5-ac54-fc19d19246b7","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3674d926-edb9-4bb8-b034-d4c6904dba2d","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719443582013	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
584ed02e-e518-4594-b76a-25a5248828cf	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719443858101	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
2b4cd7f8-852b-40c1-b22c-85bdecd519f2	midas	{"token_id":"ef70fe47-944f-4556-9cd5-3e0bd9d16b17","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"989ade86-1aa8-4a87-b221-24de96a31b26","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719443858160	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
aa0c988c-2f44-4d30-96c4-ff3ab8121447	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719444286426	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
0576b819-8e4a-416f-a4cb-eb56ad134f0b	midas	{"token_id":"2f237e3a-9786-4f74-acab-fcae9bf28b55","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"43836862-d3c2-4b30-96f0-b0a97a9f0c07","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719444286491	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
258742d1-c314-41b5-99d2-fdb5f94164dc	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","response_mode":"fragment","username":"nocactest"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719444750717	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
a1b18336-d2b8-49ad-add4-aa3b36ab3cf1	midas	{"token_id":"294b8131-ef7e-436a-a064-ae9e04db0165","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"c5ed93a7-04fb-4125-9252-5b53c5b0bf64","code_id":"e2225ce6-1ab0-497e-96f0-d624fb51b142","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	e2225ce6-1ab0-497e-96f0-d624fb51b142	1719444750810	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
60d677f5-ca8e-46df-a6a2-de92c9360c54	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"a2b2f44d-d8a6-4dab-9b7f-dc15b0e3c70c","username":"nocactest"}	\N	192.168.48.1	cwbi	a2b2f44d-d8a6-4dab-9b7f-dc15b0e3c70c	1719444763894	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
f0755ce7-7c3c-4539-b7a1-b12422a5be50	midas	{"token_id":"9b50e160-4d0d-45bd-8a66-aeb3bae4879e","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"c0aca88d-d14b-40c7-b8c2-f3663ecf27d9","code_id":"a2b2f44d-d8a6-4dab-9b7f-dc15b0e3c70c","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	a2b2f44d-d8a6-4dab-9b7f-dc15b0e3c70c	1719444764801	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
f614e5c5-6842-4468-bfd4-d8a8de22c661	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"6f0de584-0f10-4264-93c6-3207cc57260e","username":"nocactest"}	\N	192.168.48.1	cwbi	6f0de584-0f10-4264-93c6-3207cc57260e	1719444923908	LOGIN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
ba8068ec-487a-445d-b426-2549fd4fedc4	midas	{"token_id":"1684605e-81c4-4d0c-a488-1e454ff8d2d5","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3416d07e-8df9-4f8e-ae10-f12c26ae04fc","code_id":"6f0de584-0f10-4264-93c6-3207cc57260e","client_auth_method":"client-secret"}	\N	192.168.48.1	cwbi	6f0de584-0f10-4264-93c6-3207cc57260e	1719444924883	CODE_TO_TOKEN	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
4cb95b11-e400-4d56-8392-d61f0c8e2849	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"55a38410-b834-4954-bddb-6f92ef622971","username":"test"}	\N	172.19.0.1	cwbi	55a38410-b834-4954-bddb-6f92ef622971	1720194792004	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
d1bb0c83-95c8-4150-8bcb-e629967cdb9f	midas	{"token_id":"4aa0bc62-3ad8-49ce-b975-be907ca5bd38","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"1ae79a08-d12e-48c5-9453-756933bdc2bc","code_id":"55a38410-b834-4954-bddb-6f92ef622971","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	55a38410-b834-4954-bddb-6f92ef622971	1720194793020	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
71fa4fad-8331-4e57-84b8-0d481617315c	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"55a38410-b834-4954-bddb-6f92ef622971","response_mode":"fragment","username":"test"}	\N	172.19.0.1	cwbi	55a38410-b834-4954-bddb-6f92ef622971	1720195165050	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
512738db-41fc-43e3-91b1-6c393b20aa7f	midas	{"token_id":"47c9ad36-a5e7-4b5f-89be-a8de36f6cc43","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"90c8acbb-63aa-49b1-8196-3b690c0cbeca","code_id":"55a38410-b834-4954-bddb-6f92ef622971","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	55a38410-b834-4954-bddb-6f92ef622971	1720195165105	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c641334f-1614-41c0-9cce-37b54cdd39f4	\N	{"redirect_uri":"http://localhost:3000/"}	\N	172.19.0.1	cwbi	55a38410-b834-4954-bddb-6f92ef622971	1720195168627	LOGOUT	f8dcafea-243e-4b89-8d7d-fa01918130f4
e64c75ad-f20a-4851-9675-b0215fd80ad5	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","username":"test"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720195179752	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
71ba83f5-bbe4-4062-9891-90dd98e47f8f	midas	{"token_id":"663fee1b-e5f9-4eba-a2b4-a613b238eb94","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"840fa03e-8f75-4c7c-b89b-59b864091d7d","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720195180635	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
55144ff0-4020-4c25-9af8-9efdb1102dd2	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","response_mode":"fragment","username":"test"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720197970655	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9952b20e-4d01-4dbd-ab71-3e6e0243f7c2	midas	{"token_id":"7e8c8d1f-5ce1-41fa-9e77-65018ae7ccf9","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3247e296-cbbb-4674-a6cc-a9b6d494f820","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720197970704	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
c0076dcf-bbc0-4849-9ce3-598734d943a6	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","response_mode":"fragment","username":"test"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720198420222	LOGIN	f8dcafea-243e-4b89-8d7d-fa01918130f4
064eafa9-7bee-491d-8b31-532705cd8aa5	midas	{"token_id":"a2ca212b-9ed5-4c51-bdb4-bd24a4a4158b","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"71ae5de9-b574-4ca9-9668-ac470f59f3f2","code_id":"d5ea5d16-fcfb-427a-9bfb-8be248c07ac2","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720198420259	CODE_TO_TOKEN	f8dcafea-243e-4b89-8d7d-fa01918130f4
9e8e4d4b-8d86-4bed-8609-2f6a54f9dc5a	\N	{"redirect_uri":"http://localhost:3000/blue-water-dam-example-project#dashboard"}	\N	172.19.0.1	cwbi	d5ea5d16-fcfb-427a-9bfb-8be248c07ac2	1720198481958	LOGOUT	f8dcafea-243e-4b89-8d7d-fa01918130f4
560311b7-1bf7-4053-9505-e98da6073c0c	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/blue-water-dam-example-project#dashboard","code_id":"b64bca9d-9370-4ba2-b29e-18cc8387ac70","username":"newuser"}	user_not_found	172.19.0.1	cwbi	\N	1720198497268	LOGIN_ERROR	\N
5d2740f0-42fe-449c-8f33-b6385979287f	midas	{"auth_method":"openid-connect","auth_type":"code","register_method":"form","last_name":"TestUser","redirect_uri":"http://localhost:3000/blue-water-dam-example-project#dashboard","first_name":"New","code_id":"deb471d7-7554-4cf6-8b9c-b04d87161ca9","email":"thisisatestemail@fake.usace.army.mil","username":"aaaaaaa"}	\N	172.19.0.1	cwbi	\N	1720198987621	REGISTER	f9b33064-13d0-47d7-8294-fb8f0fac819f
3720d91a-3c15-4480-acfc-84d0a274aa29	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/blue-water-dam-example-project#dashboard","consent":"no_consent_required","code_id":"deb471d7-7554-4cf6-8b9c-b04d87161ca9","username":"aaaaaaa"}	\N	172.19.0.1	cwbi	deb471d7-7554-4cf6-8b9c-b04d87161ca9	1720198987679	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
70ab02ce-946a-4b80-a479-9d4ff5fc48ac	midas	{"token_id":"00c0879f-dce7-464e-9872-bb55d1f5ecd4","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3cfa6bbd-62f6-4a1c-8a24-bb6e53869f5d","code_id":"deb471d7-7554-4cf6-8b9c-b04d87161ca9","client_auth_method":"client-secret"}	\N	172.19.0.1	cwbi	deb471d7-7554-4cf6-8b9c-b04d87161ca9	1720198988530	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
600f5a9b-ac3e-4173-8194-e8f107bb4879	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"45ae96a2-cc5e-4ea6-98fd-5e49b80fc732","username":"aaaaaaa"}	\N	172.18.0.1	cwbi	45ae96a2-cc5e-4ea6-98fd-5e49b80fc732	1725660672145	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
ff806a10-71c2-4643-a4b4-a7dee6a778d9	midas	{"token_id":"f40f432c-dd10-4e49-a2ce-1b5c97bc3ace","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"e774d295-8c4d-4fe2-b656-a06e20a40265","code_id":"45ae96a2-cc5e-4ea6-98fd-5e49b80fc732","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	45ae96a2-cc5e-4ea6-98fd-5e49b80fc732	1725660672854	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
e643e506-220c-406d-af63-fc9bf5f4efb8	\N	{"redirect_uri":"http://localhost:3000/"}	\N	172.18.0.1	cwbi	45ae96a2-cc5e-4ea6-98fd-5e49b80fc732	1725660711439	LOGOUT	f9b33064-13d0-47d7-8294-fb8f0fac819f
93492a81-6869-4550-9aba-7c3d3bb50258	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"a8082c75-6e89-48b3-b250-d59fc74706cf","username":"aaaaaaa"}	\N	172.18.0.1	cwbi	a8082c75-6e89-48b3-b250-d59fc74706cf	1725660734079	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
799a2f88-9f00-4e68-b5f3-e30ca94b61a5	midas	{"token_id":"8b70993b-3cc8-47de-b24d-1a2cd8053802","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"b996467b-9b2d-4c8d-a6ec-bc381fecd90a","code_id":"a8082c75-6e89-48b3-b250-d59fc74706cf","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	a8082c75-6e89-48b3-b250-d59fc74706cf	1725660735065	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
588caf08-9a52-48af-a4dd-2dac3fc9818f	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"a8082c75-6e89-48b3-b250-d59fc74706cf","response_mode":"fragment","username":"newcacuser"}	\N	172.18.0.1	cwbi	a8082c75-6e89-48b3-b250-d59fc74706cf	1725661086020	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
3deddde8-26f2-4fd3-a520-ea252e117aec	midas	{"token_id":"bee50885-69e6-445d-9187-d256d092e54a","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"bb3746ef-efde-466a-946e-61099f269bef","code_id":"a8082c75-6e89-48b3-b250-d59fc74706cf","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	a8082c75-6e89-48b3-b250-d59fc74706cf	1725661086052	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
1b039f6d-7999-47d3-9467-8c0391b5d9d7	\N	{"redirect_uri":"http://localhost:3000/"}	\N	172.18.0.1	cwbi	a8082c75-6e89-48b3-b250-d59fc74706cf	1725661092160	LOGOUT	f9b33064-13d0-47d7-8294-fb8f0fac819f
6b2038e5-9a02-4f7e-a2ee-c7ad122d562f	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"5bffeee7-eca2-48c7-884d-038c0f4fb430","username":"newcacuser"}	\N	172.18.0.1	cwbi	5bffeee7-eca2-48c7-884d-038c0f4fb430	1725661113341	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
85015ec3-ad27-4393-81de-23dbc37c6684	midas	{"token_id":"378c64d8-3c1d-4f41-8804-205c88dae06f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"a9509f19-d319-4776-a80c-dabded84568c","code_id":"5bffeee7-eca2-48c7-884d-038c0f4fb430","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	5bffeee7-eca2-48c7-884d-038c0f4fb430	1725661114328	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
48ea06f7-2daa-4259-8e59-3db4ed880a21	\N	{"redirect_uri":"http://localhost:3000/"}	\N	172.18.0.1	cwbi	5bffeee7-eca2-48c7-884d-038c0f4fb430	1725661324686	LOGOUT	f9b33064-13d0-47d7-8294-fb8f0fac819f
6d652666-6561-47a8-a20a-c0ef91f24d6d	midas	{"auth_method":"openid-connect","auth_type":"code","redirect_uri":"http://localhost:3000/","consent":"no_consent_required","code_id":"3accb510-2111-454c-b0fa-9a1543908453","username":"newcacuser"}	\N	172.18.0.1	cwbi	3accb510-2111-454c-b0fa-9a1543908453	1725661342555	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
4d3f6de2-7483-4316-a3a3-eaf4f3811d29	midas	{"token_id":"e281bb09-e205-4ff3-9d67-0e0879c43e1f","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"3a700787-fd9e-4870-ac86-bb6ec4c7cc8c","code_id":"3accb510-2111-454c-b0fa-9a1543908453","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	3accb510-2111-454c-b0fa-9a1543908453	1725661343568	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
d9ae2c13-7591-42fd-95b0-4fd375bfdfa0	midas	{"auth_method":"openid-connect","auth_type":"code","response_type":"code","redirect_uri":"http://localhost:3000/silent-check-sso.html","consent":"no_consent_required","code_id":"3accb510-2111-454c-b0fa-9a1543908453","response_mode":"fragment","username":"newcacuser"}	\N	172.18.0.1	cwbi	3accb510-2111-454c-b0fa-9a1543908453	1725666012850	LOGIN	f9b33064-13d0-47d7-8294-fb8f0fac819f
6e638b6e-6675-4407-b5af-65daed76ecc1	midas	{"token_id":"b8b2c09f-ed7b-44ef-b3c4-81a9d172a286","grant_type":"authorization_code","refresh_token_type":"Refresh","scope":"openid email profile","refresh_token_id":"0af5180f-834b-4ac7-9922-b4034a960631","code_id":"3accb510-2111-454c-b0fa-9a1543908453","client_auth_method":"client-secret"}	\N	172.18.0.1	cwbi	3accb510-2111-454c-b0fa-9a1543908453	1725666012886	CODE_TO_TOKEN	f9b33064-13d0-47d7-8294-fb8f0fac819f
\.


--
-- Data for Name: fed_user_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_attribute (id, name, user_id, realm_id, storage_provider_id, value) FROM stdin;
\.


--
-- Data for Name: fed_user_consent; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_consent (id, client_id, user_id, realm_id, storage_provider_id, created_date, last_updated_date, client_storage_provider, external_client_id) FROM stdin;
\.


--
-- Data for Name: fed_user_consent_cl_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_consent_cl_scope (user_consent_id, scope_id) FROM stdin;
\.


--
-- Data for Name: fed_user_credential; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_credential (id, salt, type, created_date, user_id, realm_id, storage_provider_id, user_label, secret_data, credential_data, priority) FROM stdin;
\.


--
-- Data for Name: fed_user_group_membership; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_group_membership (group_id, user_id, realm_id, storage_provider_id) FROM stdin;
\.


--
-- Data for Name: fed_user_required_action; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_required_action (required_action, user_id, realm_id, storage_provider_id) FROM stdin;
\.


--
-- Data for Name: fed_user_role_mapping; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.fed_user_role_mapping (role_id, user_id, realm_id, storage_provider_id) FROM stdin;
\.


--
-- Data for Name: federated_identity; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.federated_identity (identity_provider, realm_id, federated_user_id, federated_username, token, user_id) FROM stdin;
\.


--
-- Data for Name: federated_user; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.federated_user (id, storage_provider_id, realm_id) FROM stdin;
\.


--
-- Data for Name: group_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.group_attribute (id, name, value, group_id) FROM stdin;
\.


--
-- Data for Name: group_role_mapping; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.group_role_mapping (role_id, group_id) FROM stdin;
\.


--
-- Data for Name: identity_provider; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.identity_provider (internal_id, enabled, provider_alias, provider_id, store_token, authenticate_by_default, realm_id, add_token_role, trust_email, first_broker_login_flow_id, post_broker_login_flow_id, provider_display_name, link_only) FROM stdin;
\.


--
-- Data for Name: identity_provider_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.identity_provider_config (identity_provider_id, value, name) FROM stdin;
\.


--
-- Data for Name: identity_provider_mapper; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.identity_provider_mapper (id, name, idp_alias, idp_mapper_name, realm_id) FROM stdin;
\.


--
-- Data for Name: idp_mapper_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.idp_mapper_config (idp_mapper_id, value, name) FROM stdin;
\.


--
-- Data for Name: keycloak_group; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.keycloak_group (id, name, parent_group, realm_id) FROM stdin;
\.


--
-- Data for Name: keycloak_role; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.keycloak_role (id, client_realm_constraint, client_role, description, name, realm_id, client, realm) FROM stdin;
db5807fc-5e23-4f38-8243-bfdc9673bd9e	98749fe9-5c8f-4d46-b973-16664c916f0f	f	${role_default-roles}	default-roles-master	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	\N
7d9ddb34-3373-42d3-b09a-135a12a3d915	98749fe9-5c8f-4d46-b973-16664c916f0f	f	${role_create-realm}	create-realm	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	\N
062a0cf2-00b3-45a3-b488-da81142c49fd	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_create-client}	create-client	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
d594560c-cd04-4b97-9180-a5bd509c2697	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-realm}	view-realm	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
74e9bad0-3106-45f8-80b3-d1a51355b046	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-users}	view-users	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
c57caf0d-3402-464e-886f-7cc91131b4be	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-clients}	view-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
956b7000-9d33-4356-b2fe-5fb9eb8ace5c	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-events}	view-events	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
63c6007d-8f4d-42b1-bae6-f2d57280ae4e	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-identity-providers}	view-identity-providers	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
d6b070b9-5c18-4609-afed-966e7ab9e3d3	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_view-authorization}	view-authorization	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
bd54f0dc-57f3-4648-b106-6c09ce9d9ef0	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-realm}	manage-realm	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
229c3773-8383-4e57-a4a2-da2a57d8e2be	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-users}	manage-users	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
18e3c446-5a2a-4d6b-9ff8-fcda1b265a6a	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-clients}	manage-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
5d9f3e23-f626-46d5-b7a3-ea54ef829943	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-events}	manage-events	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
7fe30e28-ace6-492c-82e9-5abe06edc752	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-identity-providers}	manage-identity-providers	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
1ed30b08-d79b-492c-a49a-ceb1c2065cc7	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_manage-authorization}	manage-authorization	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
a8088132-ed0e-4d2c-924d-76971d514b72	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_query-users}	query-users	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
7891a41f-6a7c-4dcb-86a7-954be015b7ad	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_query-clients}	query-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
eb54652f-08d4-49ce-86ce-10626d6699c6	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_query-realms}	query-realms	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
8eb26a12-1853-44b8-bf59-4da9e0cb683c	98749fe9-5c8f-4d46-b973-16664c916f0f	f	${role_admin}	admin	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	\N
1205b878-4df1-4bb3-a644-e35ee220958a	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_query-groups}	query-groups	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
f0c7cd6e-d92c-47d8-aba6-fe2c7bc0d4c0	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_view-profile}	view-profile	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
98e010f7-b9db-4a9b-872a-8eee15766444	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_manage-account}	manage-account	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
4426226e-0666-44dd-ad62-bbe162df3720	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_manage-account-links}	manage-account-links	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
0e316da2-7fb7-445f-a13a-711d27fe1a22	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_view-applications}	view-applications	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
0c394c74-e2a3-41dc-a6d4-03b770b1cd0a	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_view-consent}	view-consent	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
50bf3d94-3f5c-41ae-9007-640324a3c6f7	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_manage-consent}	manage-consent	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
0a5420b4-1a6d-42b7-a07a-19aa82f6b26e	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_view-groups}	view-groups	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
b6fb955c-4133-446e-a704-d907da319eac	f656ab57-d2fe-4f82-a765-0357d7ef4a46	t	${role_delete-account}	delete-account	98749fe9-5c8f-4d46-b973-16664c916f0f	f656ab57-d2fe-4f82-a765-0357d7ef4a46	\N
d43a834a-5d41-4e3c-a545-4a1f17e2318c	84727cee-1c0d-426e-9f90-e5d126ef055b	t	${role_read-token}	read-token	98749fe9-5c8f-4d46-b973-16664c916f0f	84727cee-1c0d-426e-9f90-e5d126ef055b	\N
0849c1f4-5f15-4b30-b6a7-8f65aa40ee4e	ebe2670b-ba08-442e-9983-2807d8e8dbba	t	${role_impersonation}	impersonation	98749fe9-5c8f-4d46-b973-16664c916f0f	ebe2670b-ba08-442e-9983-2807d8e8dbba	\N
aafc97c2-f3e5-4566-b148-91289a4ed570	98749fe9-5c8f-4d46-b973-16664c916f0f	f	${role_offline-access}	offline_access	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	\N
c1638b4b-5d98-42ad-bb3f-85617b7b2bfa	98749fe9-5c8f-4d46-b973-16664c916f0f	f	${role_uma_authorization}	uma_authorization	98749fe9-5c8f-4d46-b973-16664c916f0f	\N	\N
c883c84f-d797-4170-bcf1-888f3843ba15	cwbi	f	${role_default-roles-cwbi}	default-roles-cwbi	cwbi	\N	\N
4aa490af-405e-4ebb-b8ef-1cd3fbfa9c84	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_create-client}	create-client	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
a0edf521-eb53-4487-94e6-b9036a5e208c	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-realm}	view-realm	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
818ed298-e735-4251-875a-0778a78aaf81	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-users}	view-users	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
e46df155-c26e-4253-9bce-998723e7ee41	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-clients}	view-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
85cb9aca-73a3-4c7d-b090-e8b9f9f6a627	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-events}	view-events	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
06728afe-43b1-4807-9534-c1568181147c	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-identity-providers}	view-identity-providers	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
8143ab94-1464-4cc8-9236-cb541a045bb2	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_view-authorization}	view-authorization	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
a4df278e-8a69-4a4c-b53f-53325b4fbfe6	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-realm}	manage-realm	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
28379ac4-b0d7-45b5-89da-51d38484e289	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-users}	manage-users	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
245e9f9a-cc0b-4026-a922-3548120a11fc	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-clients}	manage-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
6522efd1-7218-4939-bd7a-32ccf7727e89	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-events}	manage-events	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
095429cc-740d-40c5-9d1e-0f0ce3f3c789	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-identity-providers}	manage-identity-providers	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
8994378a-9b80-4fc1-b43f-90b34dd4bef9	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_manage-authorization}	manage-authorization	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
97dd6d41-613e-4bf3-85af-982967e947b1	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_query-users}	query-users	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
8ccf9449-b244-423c-b724-50f165af55ab	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_query-clients}	query-clients	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
e278b6d3-25f1-4105-927a-e80a9e82a351	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_query-realms}	query-realms	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
a0e29540-5e4f-428b-aebf-e5a2360570c5	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_query-groups}	query-groups	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
f7b61645-41f7-4916-971d-bacb13088a1d	cwbi	f	${role_uma_authorization}	uma_authorization	cwbi	\N	\N
51193cbf-31d0-4955-8374-b2bee6bff6c4	cwbi	f	${role_offline-access}	offline_access	cwbi	\N	\N
f50d9a09-df21-4537-8d08-79c33fadf74f	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-clients}	view-clients	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
82d458ac-cb98-4cac-9bda-b0d2f0615238	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_create-client}	create-client	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
3c0bfa33-a92b-41e3-97e7-f38946afa659	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-events}	view-events	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
a0d3ca7a-9748-4dea-8c25-f417a0d60899	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-authorization}	manage-authorization	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
1a722e8c-695d-439a-b271-1aab67cbf3cb	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-identity-providers}	view-identity-providers	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
b03e73e7-b952-4f18-bff0-fe7dd5cb72e2	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_query-groups}	query-groups	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
97b3fa9d-3031-4405-8f19-95b0a51fee30	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_realm-admin}	realm-admin	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
103043c1-8e66-4507-9dad-0ae3e02a0801	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_query-users}	query-users	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
4d43ab7e-ce46-49f2-a2ab-9296a6a13aa6	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-users}	view-users	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
f8483adf-20d3-4c60-89bf-423fb7127254	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-events}	manage-events	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
755d6e23-cf4f-41b4-8b75-7e6eb1e56f76	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-authorization}	view-authorization	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
c60082f0-7d04-4c53-89e0-a18bdef5ccbd	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_view-realm}	view-realm	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
b67f2539-d402-4b27-844f-8a62c067893e	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-realm}	manage-realm	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
259881c3-5ae3-48b2-a21e-5baa30805bef	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_query-clients}	query-clients	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
15672aea-3a89-42bd-99e5-e316c39e0ccf	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-clients}	manage-clients	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
5c4876da-376b-4db3-b8e2-7077c4d38455	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-users}	manage-users	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
429bd26e-2f7f-4150-bb87-a677fd7eac4a	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_query-realms}	query-realms	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
1ff58d47-9fa7-4e20-a4c4-3072fae9b049	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_impersonation}	impersonation	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
ccef0b3d-88cf-45e0-99a4-61b463442bcd	5fdc5f46-8594-4e73-a982-9138ff9a0f89	t	${role_manage-identity-providers}	manage-identity-providers	cwbi	5fdc5f46-8594-4e73-a982-9138ff9a0f89	\N
43b1813c-4dc3-4d9c-aade-e89b4d873b3e	17db6de4-2231-432e-99a8-1432ac240dae	t	${role_read-token}	read-token	cwbi	17db6de4-2231-432e-99a8-1432ac240dae	\N
a8248e52-3588-4856-86af-0cb9a9c4b2a8	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_delete-account}	delete-account	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
d1220360-a04a-4765-a21e-767bf9848eb9	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_manage-account-links}	manage-account-links	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
16c70d55-bcfe-495b-95ae-1280ee7dca70	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_manage-consent}	manage-consent	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
6fc29792-ef9a-4035-8fcf-73a93194cee4	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_view-groups}	view-groups	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
fc6fdfd9-a276-4370-8a5e-485381792a9e	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_view-applications}	view-applications	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
7ebc1499-41ab-40a5-8331-79248f1b0372	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_view-consent}	view-consent	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
3422313a-546d-4e5d-9595-9af9a14532fe	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_manage-account}	manage-account	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
49c34741-45f3-4edd-8259-19b827460e3b	46da91ba-9a49-40ae-a8b3-a9dca846129d	t	${role_view-profile}	view-profile	cwbi	46da91ba-9a49-40ae-a8b3-a9dca846129d	\N
9d546e45-4e4b-437d-88d4-328eece3d9ac	bc580584-cd30-4019-9ee4-0c162f4f9802	t	${role_impersonation}	impersonation	98749fe9-5c8f-4d46-b973-16664c916f0f	bc580584-cd30-4019-9ee4-0c162f4f9802	\N
\.


--
-- Data for Name: migration_model; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.migration_model (id, version, update_time) FROM stdin;
ebj68	20.0.3	1719326093
\.


--
-- Data for Name: offline_client_session; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.offline_client_session (user_session_id, client_id, offline_flag, "timestamp", data, client_storage_provider, external_client_id) FROM stdin;
\.


--
-- Data for Name: offline_user_session; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.offline_user_session (user_session_id, user_id, realm_id, created_on, offline_flag, data, last_session_refresh) FROM stdin;
\.


--
-- Data for Name: policy_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.policy_config (policy_id, name, value) FROM stdin;
\.


--
-- Data for Name: protocol_mapper; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.protocol_mapper (id, name, protocol, protocol_mapper_name, client_id, client_scope_id) FROM stdin;
375a5124-562d-4e7f-a5bd-1d979b5860e2	audience resolve	openid-connect	oidc-audience-resolve-mapper	ff59471c-d8be-4e1b-a846-c01bc708fbe4	\N
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	locale	openid-connect	oidc-usermodel-attribute-mapper	fca2fb0d-1434-4ba2-bd0a-699e623e79be	\N
d4115917-fa60-4d72-880c-f9c788dff012	role list	saml	saml-role-list-mapper	\N	03ec7dfe-784f-4953-b4a5-7316c1193a66
7ae0f458-95bf-49e4-a099-506743d4e63b	full name	openid-connect	oidc-full-name-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
1598aba9-9141-40cd-9173-4b521f56cb4a	family name	openid-connect	oidc-usermodel-property-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
74b09c06-a9d4-4455-808a-9931f40193c0	given name	openid-connect	oidc-usermodel-property-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
896115ac-8ad1-4098-92f9-0eb3bc46ba48	middle name	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
8e435086-cef9-4891-a350-3b4b48de2ce3	nickname	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
07869554-f3f6-492b-80ad-1736d6e0d5d0	username	openid-connect	oidc-usermodel-property-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
6efda899-c652-4930-abb2-fabdd3164e2e	profile	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	picture	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
2b62599e-37dd-4815-a476-8d4e5deb835a	website	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	gender	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
f25de3e6-3702-462c-b43f-6793663b63e3	birthdate	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
2a31806e-0839-4c32-8326-5d9fbbb37b67	zoneinfo	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
47c953a9-d78b-456e-9f61-2d40e6c98943	locale	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
0545db45-2b95-499a-b6ad-a6e28ba42776	updated at	openid-connect	oidc-usermodel-attribute-mapper	\N	76365f1a-cc60-4b37-affe-c5c494cf2f47
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	email	openid-connect	oidc-usermodel-property-mapper	\N	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8
e161b83a-6377-41d2-9274-523c3d7b0690	email verified	openid-connect	oidc-usermodel-property-mapper	\N	4aa65e8f-2d74-4a2e-9916-5cbc0ac2e2f8
85d062b8-2662-48a4-a8cc-8fbf14160441	address	openid-connect	oidc-address-mapper	\N	f1d89373-70ef-400b-a3bb-aafbcfa4326b
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	phone number	openid-connect	oidc-usermodel-attribute-mapper	\N	bdf9e8fc-7774-4216-ae5b-9c955ec11853
a680658f-348e-4bd0-aaba-17b443fd3977	phone number verified	openid-connect	oidc-usermodel-attribute-mapper	\N	bdf9e8fc-7774-4216-ae5b-9c955ec11853
f968dfcc-a897-44b5-9854-87a2a269c0af	realm roles	openid-connect	oidc-usermodel-realm-role-mapper	\N	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	client roles	openid-connect	oidc-usermodel-client-role-mapper	\N	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795
30d0b425-bcd7-4aa3-9fce-85b35333b362	audience resolve	openid-connect	oidc-audience-resolve-mapper	\N	d4bcfbdf-fba7-4bad-b54c-85ec3a32e795
edba4780-92ae-4f29-b7a3-b4dc5c8d1b3c	allowed web origins	openid-connect	oidc-allowed-origins-mapper	\N	66485ea3-33cc-45e2-9db8-4b7632fcbcfd
025bb0ed-5558-4e52-917c-41b7bfa5fd96	upn	openid-connect	oidc-usermodel-property-mapper	\N	7bab9854-2b99-4050-b410-81a6a09c7832
27746f9a-79d4-4d35-bbb5-6d523b956964	groups	openid-connect	oidc-usermodel-realm-role-mapper	\N	7bab9854-2b99-4050-b410-81a6a09c7832
a475d640-25ab-4b08-8576-f2db293f34dd	acr loa level	openid-connect	oidc-acr-mapper	\N	016e49db-14f9-4d10-be20-4502b2a84a27
b45224bb-c9c9-44b4-aa30-601409f8c980	role list	saml	saml-role-list-mapper	\N	662e8870-a4c7-431a-9b40-56a3c4cbb6ea
8e8b1862-83b7-410f-9bdf-216ceb711a51	x509_presented	openid-connect	oidc-usermodel-attribute-mapper	\N	b0a33b5f-7c9a-4d59-9602-855dfb2a0b92
6c9ea8bb-4448-4e3e-9964-85271ac97280	email	openid-connect	oidc-usermodel-property-mapper	\N	98992cfc-118d-4c64-976b-7ba01f0976a5
8e08c3c6-d7de-4480-8904-5ba1a241bd82	email verified	openid-connect	oidc-usermodel-property-mapper	\N	98992cfc-118d-4c64-976b-7ba01f0976a5
4406309c-52c4-43b6-9a21-e9cd98e56456	upn	openid-connect	oidc-usermodel-property-mapper	\N	2b3b3db7-5772-4d81-a35a-742ea21a95e6
ae128c2c-d060-4869-afd0-a5fc07858540	groups	openid-connect	oidc-usermodel-realm-role-mapper	\N	2b3b3db7-5772-4d81-a35a-742ea21a95e6
4d12cc90-1436-41e4-9238-0c53851f31ee	address	openid-connect	oidc-address-mapper	\N	0542ff9c-210f-4eb3-b62e-fa7272032823
6943ef7d-fadf-419d-a801-2d2c43e1b10c	acr loa level	openid-connect	oidc-acr-mapper	\N	c983959d-26d1-413e-9343-f5bad7dabc51
22cc4f9a-af34-4686-9778-cb597eaac561	locale	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
592205f5-802d-4695-b8d2-c3df4b1e15fd	username	openid-connect	oidc-usermodel-property-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
4c9f6543-5f73-4496-aac8-20d6d18b7355	family name	openid-connect	oidc-usermodel-property-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	birthdate	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
a1290958-233f-4643-9cd8-dabc24135572	website	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
4465a630-5190-4d58-bad2-438100952958	nickname	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
b0f49a96-f80f-42ec-985c-6ba7c87764c0	profile	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
452a54bb-b05d-4ad8-84e2-47c794055953	full name	openid-connect	oidc-full-name-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
54a94019-2ca4-4ba7-9cd6-b2be66befce2	middle name	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
a41252d3-cfb8-4572-80b1-fd744ca57f35	given name	openid-connect	oidc-usermodel-property-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
f1c2b587-ed25-4913-a549-c1974fabb90e	updated at	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
fb318369-1a6e-4015-a721-0d200eedfecb	gender	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
464409dd-6024-4eab-81b8-fa34e26cfb6f	zoneinfo	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
b1a3b180-cc3b-441c-a7ff-56557054978c	picture	openid-connect	oidc-usermodel-attribute-mapper	\N	268fb6b2-58c0-44f2-ae22-a8baf2236e18
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	groups	openid-connect	oidc-group-membership-mapper	\N	47da9728-25a4-4462-8897-67d5b9e56d92
33d0ff8d-edef-4991-8750-eea198771be4	realm roles	openid-connect	oidc-usermodel-realm-role-mapper	\N	7d5c5b91-9cc3-467b-aced-a25db29c2576
c15aac91-3391-4909-bc25-6791846e07f2	client roles	openid-connect	oidc-usermodel-client-role-mapper	\N	7d5c5b91-9cc3-467b-aced-a25db29c2576
38746245-955e-46e3-8d1c-02c9d5ecbea3	audience resolve	openid-connect	oidc-audience-resolve-mapper	\N	7d5c5b91-9cc3-467b-aced-a25db29c2576
5f10a819-822b-44bc-b25c-306449a1252e	allowed web origins	openid-connect	oidc-allowed-origins-mapper	\N	cdc2e818-4856-4688-b54a-03a5d08e6a1d
ab6ba770-fdab-462e-968a-41f8212f79e6	cacUID	openid-connect	oidc-usermodel-attribute-mapper	\N	5286fee9-6cda-4a94-aba0-dffa0a5c2e8f
927e430b-6855-4b4c-a135-44359b0ef28e	username	openid-connect	oidc-usermodel-property-mapper	\N	9cf08b6f-66b3-46ab-b59c-cd96e9f1b8c0
29e38983-d5d7-4683-8947-f2112f258ed8	phone number	openid-connect	oidc-usermodel-attribute-mapper	\N	2648bc85-15fc-4968-b6b5-8b9743c8cfad
2c1af343-bc8f-403c-8eea-74929dabd074	phone number verified	openid-connect	oidc-usermodel-attribute-mapper	\N	2648bc85-15fc-4968-b6b5-8b9743c8cfad
068560ec-501d-4184-b349-1d926a710793	subjectDN	openid-connect	oidc-usermodel-attribute-mapper	\N	17e1b31a-5522-4e03-a85c-e476b919a19a
8c06c640-aedc-469d-8051-56d7240e59bf	audience resolve	openid-connect	oidc-audience-resolve-mapper	bc4324a3-e1d1-4b25-bb23-1e46ed21709f	\N
e4e59817-cd1e-4b91-a6be-427e1e8c3093	locale	openid-connect	oidc-usermodel-attribute-mapper	38f6b360-c12e-4e69-8b63-43ab4910e344	\N
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	cacUID	openid-connect	oidc-usermodel-attribute-mapper	86b97bc5-1afd-40b2-ad62-bddaaaf321c7	\N
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	subjectDN	openid-connect	oidc-usermodel-attribute-mapper	86b97bc5-1afd-40b2-ad62-bddaaaf321c7	\N
cfe591a3-b08f-408b-90b7-a00eb96388d2	x509_presented	openid-connect	oidc-usermodel-attribute-mapper	86b97bc5-1afd-40b2-ad62-bddaaaf321c7	\N
\.


--
-- Data for Name: protocol_mapper_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.protocol_mapper_config (protocol_mapper_id, value, name) FROM stdin;
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	true	userinfo.token.claim
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	locale	user.attribute
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	true	id.token.claim
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	true	access.token.claim
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	locale	claim.name
7c85088c-35ce-4f2d-8f6f-e98ef4977af0	String	jsonType.label
d4115917-fa60-4d72-880c-f9c788dff012	false	single
d4115917-fa60-4d72-880c-f9c788dff012	Basic	attribute.nameformat
d4115917-fa60-4d72-880c-f9c788dff012	Role	attribute.name
0545db45-2b95-499a-b6ad-a6e28ba42776	true	userinfo.token.claim
0545db45-2b95-499a-b6ad-a6e28ba42776	updatedAt	user.attribute
0545db45-2b95-499a-b6ad-a6e28ba42776	true	id.token.claim
0545db45-2b95-499a-b6ad-a6e28ba42776	true	access.token.claim
0545db45-2b95-499a-b6ad-a6e28ba42776	updated_at	claim.name
0545db45-2b95-499a-b6ad-a6e28ba42776	long	jsonType.label
07869554-f3f6-492b-80ad-1736d6e0d5d0	true	userinfo.token.claim
07869554-f3f6-492b-80ad-1736d6e0d5d0	username	user.attribute
07869554-f3f6-492b-80ad-1736d6e0d5d0	true	id.token.claim
07869554-f3f6-492b-80ad-1736d6e0d5d0	true	access.token.claim
07869554-f3f6-492b-80ad-1736d6e0d5d0	preferred_username	claim.name
07869554-f3f6-492b-80ad-1736d6e0d5d0	String	jsonType.label
1598aba9-9141-40cd-9173-4b521f56cb4a	true	userinfo.token.claim
1598aba9-9141-40cd-9173-4b521f56cb4a	lastName	user.attribute
1598aba9-9141-40cd-9173-4b521f56cb4a	true	id.token.claim
1598aba9-9141-40cd-9173-4b521f56cb4a	true	access.token.claim
1598aba9-9141-40cd-9173-4b521f56cb4a	family_name	claim.name
1598aba9-9141-40cd-9173-4b521f56cb4a	String	jsonType.label
2a31806e-0839-4c32-8326-5d9fbbb37b67	true	userinfo.token.claim
2a31806e-0839-4c32-8326-5d9fbbb37b67	zoneinfo	user.attribute
2a31806e-0839-4c32-8326-5d9fbbb37b67	true	id.token.claim
2a31806e-0839-4c32-8326-5d9fbbb37b67	true	access.token.claim
2a31806e-0839-4c32-8326-5d9fbbb37b67	zoneinfo	claim.name
2a31806e-0839-4c32-8326-5d9fbbb37b67	String	jsonType.label
2b62599e-37dd-4815-a476-8d4e5deb835a	true	userinfo.token.claim
2b62599e-37dd-4815-a476-8d4e5deb835a	website	user.attribute
2b62599e-37dd-4815-a476-8d4e5deb835a	true	id.token.claim
2b62599e-37dd-4815-a476-8d4e5deb835a	true	access.token.claim
2b62599e-37dd-4815-a476-8d4e5deb835a	website	claim.name
2b62599e-37dd-4815-a476-8d4e5deb835a	String	jsonType.label
47c953a9-d78b-456e-9f61-2d40e6c98943	true	userinfo.token.claim
47c953a9-d78b-456e-9f61-2d40e6c98943	locale	user.attribute
47c953a9-d78b-456e-9f61-2d40e6c98943	true	id.token.claim
47c953a9-d78b-456e-9f61-2d40e6c98943	true	access.token.claim
47c953a9-d78b-456e-9f61-2d40e6c98943	locale	claim.name
47c953a9-d78b-456e-9f61-2d40e6c98943	String	jsonType.label
6efda899-c652-4930-abb2-fabdd3164e2e	true	userinfo.token.claim
6efda899-c652-4930-abb2-fabdd3164e2e	profile	user.attribute
6efda899-c652-4930-abb2-fabdd3164e2e	true	id.token.claim
6efda899-c652-4930-abb2-fabdd3164e2e	true	access.token.claim
6efda899-c652-4930-abb2-fabdd3164e2e	profile	claim.name
6efda899-c652-4930-abb2-fabdd3164e2e	String	jsonType.label
74b09c06-a9d4-4455-808a-9931f40193c0	true	userinfo.token.claim
74b09c06-a9d4-4455-808a-9931f40193c0	firstName	user.attribute
74b09c06-a9d4-4455-808a-9931f40193c0	true	id.token.claim
74b09c06-a9d4-4455-808a-9931f40193c0	true	access.token.claim
74b09c06-a9d4-4455-808a-9931f40193c0	given_name	claim.name
74b09c06-a9d4-4455-808a-9931f40193c0	String	jsonType.label
7ae0f458-95bf-49e4-a099-506743d4e63b	true	userinfo.token.claim
7ae0f458-95bf-49e4-a099-506743d4e63b	true	id.token.claim
7ae0f458-95bf-49e4-a099-506743d4e63b	true	access.token.claim
896115ac-8ad1-4098-92f9-0eb3bc46ba48	true	userinfo.token.claim
896115ac-8ad1-4098-92f9-0eb3bc46ba48	middleName	user.attribute
896115ac-8ad1-4098-92f9-0eb3bc46ba48	true	id.token.claim
896115ac-8ad1-4098-92f9-0eb3bc46ba48	true	access.token.claim
896115ac-8ad1-4098-92f9-0eb3bc46ba48	middle_name	claim.name
896115ac-8ad1-4098-92f9-0eb3bc46ba48	String	jsonType.label
8e435086-cef9-4891-a350-3b4b48de2ce3	true	userinfo.token.claim
8e435086-cef9-4891-a350-3b4b48de2ce3	nickname	user.attribute
8e435086-cef9-4891-a350-3b4b48de2ce3	true	id.token.claim
8e435086-cef9-4891-a350-3b4b48de2ce3	true	access.token.claim
8e435086-cef9-4891-a350-3b4b48de2ce3	nickname	claim.name
8e435086-cef9-4891-a350-3b4b48de2ce3	String	jsonType.label
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	true	userinfo.token.claim
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	gender	user.attribute
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	true	id.token.claim
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	true	access.token.claim
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	gender	claim.name
b1b7ebab-5323-41bb-a371-3bc2814bcdf3	String	jsonType.label
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	true	userinfo.token.claim
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	picture	user.attribute
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	true	id.token.claim
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	true	access.token.claim
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	picture	claim.name
bf6bdcb3-a6c2-417f-af8d-4253d9b2c257	String	jsonType.label
f25de3e6-3702-462c-b43f-6793663b63e3	true	userinfo.token.claim
f25de3e6-3702-462c-b43f-6793663b63e3	birthdate	user.attribute
f25de3e6-3702-462c-b43f-6793663b63e3	true	id.token.claim
f25de3e6-3702-462c-b43f-6793663b63e3	true	access.token.claim
f25de3e6-3702-462c-b43f-6793663b63e3	birthdate	claim.name
f25de3e6-3702-462c-b43f-6793663b63e3	String	jsonType.label
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	true	userinfo.token.claim
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	email	user.attribute
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	true	id.token.claim
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	true	access.token.claim
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	email	claim.name
6fc6a240-41cf-4b9b-9635-1d4ca00757a5	String	jsonType.label
e161b83a-6377-41d2-9274-523c3d7b0690	true	userinfo.token.claim
e161b83a-6377-41d2-9274-523c3d7b0690	emailVerified	user.attribute
e161b83a-6377-41d2-9274-523c3d7b0690	true	id.token.claim
e161b83a-6377-41d2-9274-523c3d7b0690	true	access.token.claim
e161b83a-6377-41d2-9274-523c3d7b0690	email_verified	claim.name
e161b83a-6377-41d2-9274-523c3d7b0690	boolean	jsonType.label
85d062b8-2662-48a4-a8cc-8fbf14160441	formatted	user.attribute.formatted
85d062b8-2662-48a4-a8cc-8fbf14160441	country	user.attribute.country
85d062b8-2662-48a4-a8cc-8fbf14160441	postal_code	user.attribute.postal_code
85d062b8-2662-48a4-a8cc-8fbf14160441	true	userinfo.token.claim
85d062b8-2662-48a4-a8cc-8fbf14160441	street	user.attribute.street
85d062b8-2662-48a4-a8cc-8fbf14160441	true	id.token.claim
85d062b8-2662-48a4-a8cc-8fbf14160441	region	user.attribute.region
85d062b8-2662-48a4-a8cc-8fbf14160441	true	access.token.claim
85d062b8-2662-48a4-a8cc-8fbf14160441	locality	user.attribute.locality
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	true	userinfo.token.claim
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	phoneNumber	user.attribute
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	true	id.token.claim
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	true	access.token.claim
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	phone_number	claim.name
84d6b0a4-29e2-48f1-addd-f3a0bf8ec925	String	jsonType.label
a680658f-348e-4bd0-aaba-17b443fd3977	true	userinfo.token.claim
a680658f-348e-4bd0-aaba-17b443fd3977	phoneNumberVerified	user.attribute
a680658f-348e-4bd0-aaba-17b443fd3977	true	id.token.claim
a680658f-348e-4bd0-aaba-17b443fd3977	true	access.token.claim
a680658f-348e-4bd0-aaba-17b443fd3977	phone_number_verified	claim.name
a680658f-348e-4bd0-aaba-17b443fd3977	boolean	jsonType.label
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	true	multivalued
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	foo	user.attribute
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	true	access.token.claim
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	resource_access.${client_id}.roles	claim.name
d79c3dbb-6c21-457d-a0b4-bd635a2ebd97	String	jsonType.label
f968dfcc-a897-44b5-9854-87a2a269c0af	true	multivalued
f968dfcc-a897-44b5-9854-87a2a269c0af	foo	user.attribute
f968dfcc-a897-44b5-9854-87a2a269c0af	true	access.token.claim
f968dfcc-a897-44b5-9854-87a2a269c0af	realm_access.roles	claim.name
f968dfcc-a897-44b5-9854-87a2a269c0af	String	jsonType.label
025bb0ed-5558-4e52-917c-41b7bfa5fd96	true	userinfo.token.claim
025bb0ed-5558-4e52-917c-41b7bfa5fd96	username	user.attribute
025bb0ed-5558-4e52-917c-41b7bfa5fd96	true	id.token.claim
025bb0ed-5558-4e52-917c-41b7bfa5fd96	true	access.token.claim
025bb0ed-5558-4e52-917c-41b7bfa5fd96	upn	claim.name
025bb0ed-5558-4e52-917c-41b7bfa5fd96	String	jsonType.label
27746f9a-79d4-4d35-bbb5-6d523b956964	true	multivalued
27746f9a-79d4-4d35-bbb5-6d523b956964	foo	user.attribute
27746f9a-79d4-4d35-bbb5-6d523b956964	true	id.token.claim
27746f9a-79d4-4d35-bbb5-6d523b956964	true	access.token.claim
27746f9a-79d4-4d35-bbb5-6d523b956964	groups	claim.name
27746f9a-79d4-4d35-bbb5-6d523b956964	String	jsonType.label
a475d640-25ab-4b08-8576-f2db293f34dd	true	id.token.claim
a475d640-25ab-4b08-8576-f2db293f34dd	true	access.token.claim
b45224bb-c9c9-44b4-aa30-601409f8c980	false	single
b45224bb-c9c9-44b4-aa30-601409f8c980	Basic	attribute.nameformat
b45224bb-c9c9-44b4-aa30-601409f8c980	Role	attribute.name
8e8b1862-83b7-410f-9bdf-216ceb711a51	false	userinfo.token.claim
8e8b1862-83b7-410f-9bdf-216ceb711a51	x509_presented	user.attribute
8e8b1862-83b7-410f-9bdf-216ceb711a51	true	id.token.claim
8e8b1862-83b7-410f-9bdf-216ceb711a51	true	access.token.claim
8e8b1862-83b7-410f-9bdf-216ceb711a51	x509_presented	claim.name
8e8b1862-83b7-410f-9bdf-216ceb711a51	String	jsonType.label
6c9ea8bb-4448-4e3e-9964-85271ac97280	true	userinfo.token.claim
6c9ea8bb-4448-4e3e-9964-85271ac97280	email	user.attribute
6c9ea8bb-4448-4e3e-9964-85271ac97280	true	id.token.claim
6c9ea8bb-4448-4e3e-9964-85271ac97280	true	access.token.claim
6c9ea8bb-4448-4e3e-9964-85271ac97280	email	claim.name
6c9ea8bb-4448-4e3e-9964-85271ac97280	String	jsonType.label
8e08c3c6-d7de-4480-8904-5ba1a241bd82	true	userinfo.token.claim
8e08c3c6-d7de-4480-8904-5ba1a241bd82	emailVerified	user.attribute
8e08c3c6-d7de-4480-8904-5ba1a241bd82	true	id.token.claim
8e08c3c6-d7de-4480-8904-5ba1a241bd82	true	access.token.claim
8e08c3c6-d7de-4480-8904-5ba1a241bd82	email_verified	claim.name
8e08c3c6-d7de-4480-8904-5ba1a241bd82	boolean	jsonType.label
4406309c-52c4-43b6-9a21-e9cd98e56456	true	userinfo.token.claim
4406309c-52c4-43b6-9a21-e9cd98e56456	username	user.attribute
4406309c-52c4-43b6-9a21-e9cd98e56456	true	id.token.claim
4406309c-52c4-43b6-9a21-e9cd98e56456	true	access.token.claim
4406309c-52c4-43b6-9a21-e9cd98e56456	upn	claim.name
4406309c-52c4-43b6-9a21-e9cd98e56456	String	jsonType.label
ae128c2c-d060-4869-afd0-a5fc07858540	true	multivalued
ae128c2c-d060-4869-afd0-a5fc07858540	true	userinfo.token.claim
ae128c2c-d060-4869-afd0-a5fc07858540	foo	user.attribute
ae128c2c-d060-4869-afd0-a5fc07858540	true	id.token.claim
ae128c2c-d060-4869-afd0-a5fc07858540	true	access.token.claim
ae128c2c-d060-4869-afd0-a5fc07858540	groups	claim.name
ae128c2c-d060-4869-afd0-a5fc07858540	String	jsonType.label
4d12cc90-1436-41e4-9238-0c53851f31ee	formatted	user.attribute.formatted
4d12cc90-1436-41e4-9238-0c53851f31ee	country	user.attribute.country
4d12cc90-1436-41e4-9238-0c53851f31ee	postal_code	user.attribute.postal_code
4d12cc90-1436-41e4-9238-0c53851f31ee	true	userinfo.token.claim
4d12cc90-1436-41e4-9238-0c53851f31ee	street	user.attribute.street
4d12cc90-1436-41e4-9238-0c53851f31ee	true	id.token.claim
4d12cc90-1436-41e4-9238-0c53851f31ee	region	user.attribute.region
4d12cc90-1436-41e4-9238-0c53851f31ee	true	access.token.claim
4d12cc90-1436-41e4-9238-0c53851f31ee	locality	user.attribute.locality
6943ef7d-fadf-419d-a801-2d2c43e1b10c	true	id.token.claim
6943ef7d-fadf-419d-a801-2d2c43e1b10c	true	access.token.claim
6943ef7d-fadf-419d-a801-2d2c43e1b10c	true	userinfo.token.claim
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	true	userinfo.token.claim
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	birthdate	user.attribute
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	true	id.token.claim
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	true	access.token.claim
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	birthdate	claim.name
07c8c4d5-7c47-4dae-a0f7-2eddb8dee8c3	String	jsonType.label
22cc4f9a-af34-4686-9778-cb597eaac561	true	userinfo.token.claim
22cc4f9a-af34-4686-9778-cb597eaac561	locale	user.attribute
22cc4f9a-af34-4686-9778-cb597eaac561	true	id.token.claim
22cc4f9a-af34-4686-9778-cb597eaac561	true	access.token.claim
22cc4f9a-af34-4686-9778-cb597eaac561	locale	claim.name
22cc4f9a-af34-4686-9778-cb597eaac561	String	jsonType.label
4465a630-5190-4d58-bad2-438100952958	true	userinfo.token.claim
4465a630-5190-4d58-bad2-438100952958	nickname	user.attribute
4465a630-5190-4d58-bad2-438100952958	true	id.token.claim
4465a630-5190-4d58-bad2-438100952958	true	access.token.claim
4465a630-5190-4d58-bad2-438100952958	nickname	claim.name
4465a630-5190-4d58-bad2-438100952958	String	jsonType.label
452a54bb-b05d-4ad8-84e2-47c794055953	true	id.token.claim
452a54bb-b05d-4ad8-84e2-47c794055953	true	access.token.claim
452a54bb-b05d-4ad8-84e2-47c794055953	true	userinfo.token.claim
464409dd-6024-4eab-81b8-fa34e26cfb6f	true	userinfo.token.claim
464409dd-6024-4eab-81b8-fa34e26cfb6f	zoneinfo	user.attribute
464409dd-6024-4eab-81b8-fa34e26cfb6f	true	id.token.claim
464409dd-6024-4eab-81b8-fa34e26cfb6f	true	access.token.claim
464409dd-6024-4eab-81b8-fa34e26cfb6f	zoneinfo	claim.name
464409dd-6024-4eab-81b8-fa34e26cfb6f	String	jsonType.label
4c9f6543-5f73-4496-aac8-20d6d18b7355	true	userinfo.token.claim
4c9f6543-5f73-4496-aac8-20d6d18b7355	lastName	user.attribute
4c9f6543-5f73-4496-aac8-20d6d18b7355	true	id.token.claim
4c9f6543-5f73-4496-aac8-20d6d18b7355	true	access.token.claim
4c9f6543-5f73-4496-aac8-20d6d18b7355	family_name	claim.name
4c9f6543-5f73-4496-aac8-20d6d18b7355	String	jsonType.label
54a94019-2ca4-4ba7-9cd6-b2be66befce2	true	userinfo.token.claim
54a94019-2ca4-4ba7-9cd6-b2be66befce2	middleName	user.attribute
54a94019-2ca4-4ba7-9cd6-b2be66befce2	true	id.token.claim
54a94019-2ca4-4ba7-9cd6-b2be66befce2	true	access.token.claim
54a94019-2ca4-4ba7-9cd6-b2be66befce2	middle_name	claim.name
54a94019-2ca4-4ba7-9cd6-b2be66befce2	String	jsonType.label
592205f5-802d-4695-b8d2-c3df4b1e15fd	true	userinfo.token.claim
592205f5-802d-4695-b8d2-c3df4b1e15fd	username	user.attribute
592205f5-802d-4695-b8d2-c3df4b1e15fd	true	id.token.claim
592205f5-802d-4695-b8d2-c3df4b1e15fd	true	access.token.claim
592205f5-802d-4695-b8d2-c3df4b1e15fd	preferred_username	claim.name
592205f5-802d-4695-b8d2-c3df4b1e15fd	String	jsonType.label
a1290958-233f-4643-9cd8-dabc24135572	true	userinfo.token.claim
a1290958-233f-4643-9cd8-dabc24135572	website	user.attribute
a1290958-233f-4643-9cd8-dabc24135572	true	id.token.claim
a1290958-233f-4643-9cd8-dabc24135572	true	access.token.claim
a1290958-233f-4643-9cd8-dabc24135572	website	claim.name
a1290958-233f-4643-9cd8-dabc24135572	String	jsonType.label
a41252d3-cfb8-4572-80b1-fd744ca57f35	true	userinfo.token.claim
a41252d3-cfb8-4572-80b1-fd744ca57f35	firstName	user.attribute
a41252d3-cfb8-4572-80b1-fd744ca57f35	true	id.token.claim
a41252d3-cfb8-4572-80b1-fd744ca57f35	true	access.token.claim
a41252d3-cfb8-4572-80b1-fd744ca57f35	given_name	claim.name
a41252d3-cfb8-4572-80b1-fd744ca57f35	String	jsonType.label
b0f49a96-f80f-42ec-985c-6ba7c87764c0	true	userinfo.token.claim
b0f49a96-f80f-42ec-985c-6ba7c87764c0	profile	user.attribute
b0f49a96-f80f-42ec-985c-6ba7c87764c0	true	id.token.claim
b0f49a96-f80f-42ec-985c-6ba7c87764c0	true	access.token.claim
b0f49a96-f80f-42ec-985c-6ba7c87764c0	profile	claim.name
b0f49a96-f80f-42ec-985c-6ba7c87764c0	String	jsonType.label
b1a3b180-cc3b-441c-a7ff-56557054978c	true	userinfo.token.claim
b1a3b180-cc3b-441c-a7ff-56557054978c	picture	user.attribute
b1a3b180-cc3b-441c-a7ff-56557054978c	true	id.token.claim
b1a3b180-cc3b-441c-a7ff-56557054978c	true	access.token.claim
b1a3b180-cc3b-441c-a7ff-56557054978c	picture	claim.name
b1a3b180-cc3b-441c-a7ff-56557054978c	String	jsonType.label
f1c2b587-ed25-4913-a549-c1974fabb90e	true	userinfo.token.claim
f1c2b587-ed25-4913-a549-c1974fabb90e	updatedAt	user.attribute
f1c2b587-ed25-4913-a549-c1974fabb90e	true	id.token.claim
f1c2b587-ed25-4913-a549-c1974fabb90e	true	access.token.claim
f1c2b587-ed25-4913-a549-c1974fabb90e	updated_at	claim.name
f1c2b587-ed25-4913-a549-c1974fabb90e	String	jsonType.label
fb318369-1a6e-4015-a721-0d200eedfecb	true	userinfo.token.claim
fb318369-1a6e-4015-a721-0d200eedfecb	gender	user.attribute
fb318369-1a6e-4015-a721-0d200eedfecb	true	id.token.claim
fb318369-1a6e-4015-a721-0d200eedfecb	true	access.token.claim
fb318369-1a6e-4015-a721-0d200eedfecb	gender	claim.name
fb318369-1a6e-4015-a721-0d200eedfecb	String	jsonType.label
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	true	full.path
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	true	id.token.claim
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	true	access.token.claim
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	groups	claim.name
507bbcd7-e7a3-48f1-afad-77e2ab3e4922	true	userinfo.token.claim
33d0ff8d-edef-4991-8750-eea198771be4	foo	user.attribute
33d0ff8d-edef-4991-8750-eea198771be4	true	access.token.claim
33d0ff8d-edef-4991-8750-eea198771be4	realm_access.roles	claim.name
33d0ff8d-edef-4991-8750-eea198771be4	String	jsonType.label
33d0ff8d-edef-4991-8750-eea198771be4	true	multivalued
c15aac91-3391-4909-bc25-6791846e07f2	foo	user.attribute
c15aac91-3391-4909-bc25-6791846e07f2	true	access.token.claim
c15aac91-3391-4909-bc25-6791846e07f2	resource_access.${client_id}.roles	claim.name
c15aac91-3391-4909-bc25-6791846e07f2	String	jsonType.label
c15aac91-3391-4909-bc25-6791846e07f2	true	multivalued
ab6ba770-fdab-462e-968a-41f8212f79e6	true	userinfo.token.claim
ab6ba770-fdab-462e-968a-41f8212f79e6	cacUID	user.attribute
ab6ba770-fdab-462e-968a-41f8212f79e6	true	id.token.claim
ab6ba770-fdab-462e-968a-41f8212f79e6	true	access.token.claim
ab6ba770-fdab-462e-968a-41f8212f79e6	cacUID	claim.name
ab6ba770-fdab-462e-968a-41f8212f79e6	String	jsonType.label
927e430b-6855-4b4c-a135-44359b0ef28e	true	userinfo.token.claim
927e430b-6855-4b4c-a135-44359b0ef28e	username	user.attribute
927e430b-6855-4b4c-a135-44359b0ef28e	true	id.token.claim
927e430b-6855-4b4c-a135-44359b0ef28e	true	access.token.claim
927e430b-6855-4b4c-a135-44359b0ef28e	preferred_username	claim.name
927e430b-6855-4b4c-a135-44359b0ef28e	String	jsonType.label
29e38983-d5d7-4683-8947-f2112f258ed8	true	userinfo.token.claim
29e38983-d5d7-4683-8947-f2112f258ed8	phoneNumber	user.attribute
29e38983-d5d7-4683-8947-f2112f258ed8	true	id.token.claim
29e38983-d5d7-4683-8947-f2112f258ed8	true	access.token.claim
29e38983-d5d7-4683-8947-f2112f258ed8	phone_number	claim.name
29e38983-d5d7-4683-8947-f2112f258ed8	String	jsonType.label
2c1af343-bc8f-403c-8eea-74929dabd074	true	userinfo.token.claim
2c1af343-bc8f-403c-8eea-74929dabd074	phoneNumberVerified	user.attribute
2c1af343-bc8f-403c-8eea-74929dabd074	true	id.token.claim
2c1af343-bc8f-403c-8eea-74929dabd074	true	access.token.claim
2c1af343-bc8f-403c-8eea-74929dabd074	phone_number_verified	claim.name
2c1af343-bc8f-403c-8eea-74929dabd074	boolean	jsonType.label
068560ec-501d-4184-b349-1d926a710793	true	userinfo.token.claim
068560ec-501d-4184-b349-1d926a710793	subjectDN	user.attribute
068560ec-501d-4184-b349-1d926a710793	true	id.token.claim
068560ec-501d-4184-b349-1d926a710793	true	access.token.claim
068560ec-501d-4184-b349-1d926a710793	subjectDN	claim.name
068560ec-501d-4184-b349-1d926a710793	String	jsonType.label
e4e59817-cd1e-4b91-a6be-427e1e8c3093	true	userinfo.token.claim
e4e59817-cd1e-4b91-a6be-427e1e8c3093	locale	user.attribute
e4e59817-cd1e-4b91-a6be-427e1e8c3093	true	id.token.claim
e4e59817-cd1e-4b91-a6be-427e1e8c3093	true	access.token.claim
e4e59817-cd1e-4b91-a6be-427e1e8c3093	locale	claim.name
e4e59817-cd1e-4b91-a6be-427e1e8c3093	String	jsonType.label
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	false	aggregate.attrs
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	true	userinfo.token.claim
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	false	multivalued
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	cacUID	user.attribute
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	true	id.token.claim
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	true	access.token.claim
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	cacUID	claim.name
d6680f62-2652-47d0-9b6c-6b7eae6d8c34	String	jsonType.label
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	false	aggregate.attrs
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	true	userinfo.token.claim
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	false	multivalued
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	subjectDN	user.attribute
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	true	id.token.claim
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	true	access.token.claim
1e3b6b8c-e208-4ef5-bc46-02ab17db4808	subjectDN	claim.name
cfe591a3-b08f-408b-90b7-a00eb96388d2	false	aggregate.attrs
cfe591a3-b08f-408b-90b7-a00eb96388d2	true	userinfo.token.claim
cfe591a3-b08f-408b-90b7-a00eb96388d2	false	multivalued
cfe591a3-b08f-408b-90b7-a00eb96388d2	x509_presented	user.attribute
cfe591a3-b08f-408b-90b7-a00eb96388d2	true	id.token.claim
cfe591a3-b08f-408b-90b7-a00eb96388d2	true	access.token.claim
cfe591a3-b08f-408b-90b7-a00eb96388d2	x509_presented	claim.name
\.


--
-- Data for Name: realm; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm (id, access_code_lifespan, user_action_lifespan, access_token_lifespan, account_theme, admin_theme, email_theme, enabled, events_enabled, events_expiration, login_theme, name, not_before, password_policy, registration_allowed, remember_me, reset_password_allowed, social, ssl_required, sso_idle_timeout, sso_max_lifespan, update_profile_on_soc_login, verify_email, master_admin_client, login_lifespan, internationalization_enabled, default_locale, reg_email_as_username, admin_events_enabled, admin_events_details_enabled, edit_username_allowed, otp_policy_counter, otp_policy_window, otp_policy_period, otp_policy_digits, otp_policy_alg, otp_policy_type, browser_flow, registration_flow, direct_grant_flow, reset_credentials_flow, client_auth_flow, offline_session_idle_timeout, revoke_refresh_token, access_token_life_implicit, login_with_email_allowed, duplicate_emails_allowed, docker_auth_flow, refresh_token_max_reuse, allow_user_managed_access, sso_max_lifespan_remember_me, sso_idle_timeout_remember_me, default_role) FROM stdin;
98749fe9-5c8f-4d46-b973-16664c916f0f	60	300	60	\N	\N	\N	t	f	0	\N	master	0	\N	f	f	f	f	EXTERNAL	1800	36000	f	f	ebe2670b-ba08-442e-9983-2807d8e8dbba	1800	f	\N	f	f	f	f	0	1	30	6	HmacSHA1	totp	ee2d0274-f49c-44f6-a979-7e7f86176265	6cda22fb-db28-4a17-89cf-fca653249993	ea02467e-0b93-485a-8bf2-e401083b818f	602d6901-12a3-4e29-8bf1-98757e0a2dac	e58dc8c5-51b0-4ce5-a7be-e4270afd426e	2592000	f	900	t	f	b7d46802-4ce3-4643-bd2d-a5f250726a36	0	f	0	0	db5807fc-5e23-4f38-8243-bfdc9673bd9e
cwbi	60	300	600	identity	\N	\N	t	t	0	identity	cwbi	0	\N	t	f	t	f	EXTERNAL	3600	36000	f	f	bc580584-cd30-4019-9ee4-0c162f4f9802	1800	f	\N	f	t	f	t	0	1	30	6	HmacSHA1	totp	fc78af35-3d57-4224-9087-7a4408ab4194	ee67ed51-4cbe-44b9-addf-8d1531025bfa	0679fb07-660f-4278-a965-820043ead874	18c46dcb-0468-4180-b925-a529175a7b03	eaf2aa1e-6359-43ca-b403-c169273acae9	86400	f	900	f	t	bb69c43d-7961-4f3e-abe3-793f283e0c68	0	t	0	0	c883c84f-d797-4170-bcf1-888f3843ba15
\.


--
-- Data for Name: realm_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_attribute (name, realm_id, value) FROM stdin;
_browser_header.contentSecurityPolicyReportOnly	98749fe9-5c8f-4d46-b973-16664c916f0f	
_browser_header.xContentTypeOptions	98749fe9-5c8f-4d46-b973-16664c916f0f	nosniff
_browser_header.xRobotsTag	98749fe9-5c8f-4d46-b973-16664c916f0f	none
_browser_header.xFrameOptions	98749fe9-5c8f-4d46-b973-16664c916f0f	SAMEORIGIN
_browser_header.contentSecurityPolicy	98749fe9-5c8f-4d46-b973-16664c916f0f	frame-src 'self'; frame-ancestors 'self'; object-src 'none';
_browser_header.xXSSProtection	98749fe9-5c8f-4d46-b973-16664c916f0f	1; mode=block
_browser_header.strictTransportSecurity	98749fe9-5c8f-4d46-b973-16664c916f0f	max-age=31536000; includeSubDomains
bruteForceProtected	98749fe9-5c8f-4d46-b973-16664c916f0f	false
permanentLockout	98749fe9-5c8f-4d46-b973-16664c916f0f	false
maxFailureWaitSeconds	98749fe9-5c8f-4d46-b973-16664c916f0f	900
minimumQuickLoginWaitSeconds	98749fe9-5c8f-4d46-b973-16664c916f0f	60
waitIncrementSeconds	98749fe9-5c8f-4d46-b973-16664c916f0f	60
quickLoginCheckMilliSeconds	98749fe9-5c8f-4d46-b973-16664c916f0f	1000
maxDeltaTimeSeconds	98749fe9-5c8f-4d46-b973-16664c916f0f	43200
failureFactor	98749fe9-5c8f-4d46-b973-16664c916f0f	30
realmReusableOtpCode	98749fe9-5c8f-4d46-b973-16664c916f0f	false
displayName	98749fe9-5c8f-4d46-b973-16664c916f0f	Keycloak
displayNameHtml	98749fe9-5c8f-4d46-b973-16664c916f0f	<div class="kc-logo-text"><span>Keycloak</span></div>
defaultSignatureAlgorithm	98749fe9-5c8f-4d46-b973-16664c916f0f	RS256
offlineSessionMaxLifespanEnabled	98749fe9-5c8f-4d46-b973-16664c916f0f	false
offlineSessionMaxLifespan	98749fe9-5c8f-4d46-b973-16664c916f0f	5184000
frontendUrl	cwbi	
userProfileEnabled	cwbi	false
displayName	cwbi	Civil Work Business Intelligence Development and Testing
bruteForceProtected	cwbi	false
permanentLockout	cwbi	false
maxFailureWaitSeconds	cwbi	900
minimumQuickLoginWaitSeconds	cwbi	60
waitIncrementSeconds	cwbi	60
quickLoginCheckMilliSeconds	cwbi	1000
maxDeltaTimeSeconds	cwbi	43200
failureFactor	cwbi	30
actionTokenGeneratedByAdminLifespan	cwbi	43200
actionTokenGeneratedByUserLifespan	cwbi	300
defaultSignatureAlgorithm	cwbi	RS256
oauth2DeviceCodeLifespan	cwbi	600
oauth2DevicePollingInterval	cwbi	600
offlineSessionMaxLifespanEnabled	cwbi	false
offlineSessionMaxLifespan	cwbi	5184000
clientOfflineSessionIdleTimeout	cwbi	0
clientOfflineSessionMaxLifespan	cwbi	0
clientSessionIdleTimeout	cwbi	3600
clientSessionMaxLifespan	cwbi	32400
realmReusableOtpCode	cwbi	false
webAuthnPolicyRpEntityName	cwbi	keycloak
webAuthnPolicySignatureAlgorithms	cwbi	ES256
webAuthnPolicyRpId	cwbi	
webAuthnPolicyAttestationConveyancePreference	cwbi	not specified
webAuthnPolicyAuthenticatorAttachment	cwbi	not specified
webAuthnPolicyRequireResidentKey	cwbi	not specified
webAuthnPolicyUserVerificationRequirement	cwbi	not specified
webAuthnPolicyCreateTimeout	cwbi	0
webAuthnPolicyAvoidSameAuthenticatorRegister	cwbi	false
webAuthnPolicyRpEntityNamePasswordless	cwbi	keycloak
webAuthnPolicySignatureAlgorithmsPasswordless	cwbi	ES256
webAuthnPolicyRpIdPasswordless	cwbi	
webAuthnPolicyAttestationConveyancePreferencePasswordless	cwbi	not specified
webAuthnPolicyAuthenticatorAttachmentPasswordless	cwbi	not specified
webAuthnPolicyRequireResidentKeyPasswordless	cwbi	not specified
webAuthnPolicyUserVerificationRequirementPasswordless	cwbi	not specified
webAuthnPolicyCreateTimeoutPasswordless	cwbi	0
webAuthnPolicyAvoidSameAuthenticatorRegisterPasswordless	cwbi	false
client-policies.profiles	cwbi	{"profiles":[]}
client-policies.policies	cwbi	{"policies":[]}
cibaAuthRequestedUserHint	cwbi	login_hint
cibaBackchannelTokenDeliveryMode	cwbi	poll
cibaExpiresIn	cwbi	120
cibaInterval	cwbi	5
parRequestUriLifespan	cwbi	60
_browser_header.contentSecurityPolicyReportOnly	cwbi	
_browser_header.xContentTypeOptions	cwbi	nosniff
_browser_header.xRobotsTag	cwbi	none
_browser_header.xFrameOptions	cwbi	SAMEORIGIN
_browser_header.contentSecurityPolicy	cwbi	frame-src 'self' *; frame-ancestors 'self' *
_browser_header.xXSSProtection	cwbi	1; mode=block
_browser_header.strictTransportSecurity	cwbi	max-age=31536000; includeSubDomains
\.


--
-- Data for Name: realm_default_groups; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_default_groups (realm_id, group_id) FROM stdin;
\.


--
-- Data for Name: realm_enabled_event_types; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_enabled_event_types (realm_id, value) FROM stdin;
cwbi	UPDATE_CONSENT_ERROR
cwbi	SEND_RESET_PASSWORD
cwbi	GRANT_CONSENT
cwbi	VERIFY_PROFILE_ERROR
cwbi	UPDATE_TOTP
cwbi	REMOVE_TOTP
cwbi	REVOKE_GRANT
cwbi	LOGIN_ERROR
cwbi	CLIENT_LOGIN
cwbi	RESET_PASSWORD_ERROR
cwbi	IMPERSONATE_ERROR
cwbi	CODE_TO_TOKEN_ERROR
cwbi	CUSTOM_REQUIRED_ACTION
cwbi	OAUTH2_DEVICE_CODE_TO_TOKEN_ERROR
cwbi	RESTART_AUTHENTICATION
cwbi	UPDATE_PROFILE_ERROR
cwbi	IMPERSONATE
cwbi	LOGIN
cwbi	UPDATE_PASSWORD_ERROR
cwbi	OAUTH2_DEVICE_VERIFY_USER_CODE
cwbi	CLIENT_INITIATED_ACCOUNT_LINKING
cwbi	TOKEN_EXCHANGE
cwbi	REGISTER
cwbi	LOGOUT
cwbi	AUTHREQID_TO_TOKEN
cwbi	DELETE_ACCOUNT_ERROR
cwbi	CLIENT_REGISTER
cwbi	IDENTITY_PROVIDER_LINK_ACCOUNT
cwbi	UPDATE_PASSWORD
cwbi	DELETE_ACCOUNT
cwbi	FEDERATED_IDENTITY_LINK_ERROR
cwbi	CLIENT_DELETE
cwbi	IDENTITY_PROVIDER_FIRST_LOGIN
cwbi	VERIFY_EMAIL
cwbi	CLIENT_DELETE_ERROR
cwbi	CLIENT_LOGIN_ERROR
cwbi	RESTART_AUTHENTICATION_ERROR
cwbi	REMOVE_FEDERATED_IDENTITY_ERROR
cwbi	EXECUTE_ACTIONS
cwbi	TOKEN_EXCHANGE_ERROR
cwbi	PERMISSION_TOKEN
cwbi	SEND_IDENTITY_PROVIDER_LINK_ERROR
cwbi	EXECUTE_ACTION_TOKEN_ERROR
cwbi	SEND_VERIFY_EMAIL
cwbi	OAUTH2_DEVICE_AUTH
cwbi	EXECUTE_ACTIONS_ERROR
cwbi	REMOVE_FEDERATED_IDENTITY
cwbi	OAUTH2_DEVICE_CODE_TO_TOKEN
cwbi	IDENTITY_PROVIDER_POST_LOGIN
cwbi	IDENTITY_PROVIDER_LINK_ACCOUNT_ERROR
cwbi	UPDATE_EMAIL
cwbi	OAUTH2_DEVICE_VERIFY_USER_CODE_ERROR
cwbi	REGISTER_ERROR
cwbi	REVOKE_GRANT_ERROR
cwbi	LOGOUT_ERROR
cwbi	UPDATE_EMAIL_ERROR
cwbi	EXECUTE_ACTION_TOKEN
cwbi	CLIENT_UPDATE_ERROR
cwbi	UPDATE_PROFILE
cwbi	AUTHREQID_TO_TOKEN_ERROR
cwbi	FEDERATED_IDENTITY_LINK
cwbi	CLIENT_REGISTER_ERROR
cwbi	SEND_VERIFY_EMAIL_ERROR
cwbi	SEND_IDENTITY_PROVIDER_LINK
cwbi	RESET_PASSWORD
cwbi	CLIENT_INITIATED_ACCOUNT_LINKING_ERROR
cwbi	OAUTH2_DEVICE_AUTH_ERROR
cwbi	UPDATE_CONSENT
cwbi	REMOVE_TOTP_ERROR
cwbi	VERIFY_EMAIL_ERROR
cwbi	SEND_RESET_PASSWORD_ERROR
cwbi	CLIENT_UPDATE
cwbi	IDENTITY_PROVIDER_POST_LOGIN_ERROR
cwbi	CUSTOM_REQUIRED_ACTION_ERROR
cwbi	UPDATE_TOTP_ERROR
cwbi	CODE_TO_TOKEN
cwbi	VERIFY_PROFILE
cwbi	GRANT_CONSENT_ERROR
cwbi	IDENTITY_PROVIDER_FIRST_LOGIN_ERROR
\.


--
-- Data for Name: realm_events_listeners; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_events_listeners (realm_id, value) FROM stdin;
98749fe9-5c8f-4d46-b973-16664c916f0f	jboss-logging
cwbi	jboss-logging
\.


--
-- Data for Name: realm_localizations; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_localizations (realm_id, locale, texts) FROM stdin;
\.


--
-- Data for Name: realm_required_credential; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_required_credential (type, form_label, input, secret, realm_id) FROM stdin;
password	password	t	t	98749fe9-5c8f-4d46-b973-16664c916f0f
password	password	t	t	cwbi
\.


--
-- Data for Name: realm_smtp_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_smtp_config (realm_id, value, name) FROM stdin;
\.


--
-- Data for Name: realm_supported_locales; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.realm_supported_locales (realm_id, value) FROM stdin;
cwbi	
\.


--
-- Data for Name: redirect_uris; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.redirect_uris (client_id, value) FROM stdin;
f656ab57-d2fe-4f82-a765-0357d7ef4a46	/realms/master/account/*
ff59471c-d8be-4e1b-a846-c01bc708fbe4	/realms/master/account/*
fca2fb0d-1434-4ba2-bd0a-699e623e79be	/admin/master/console/*
46da91ba-9a49-40ae-a8b3-a9dca846129d	/realms/cwbi/account/*
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	/realms/cwbi/account/*
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	http://localhost:3000/*
38f6b360-c12e-4e69-8b63-43ab4910e344	/admin/cwbi/console/*
\.


--
-- Data for Name: required_action_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.required_action_config (required_action_id, value, name) FROM stdin;
\.


--
-- Data for Name: required_action_provider; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.required_action_provider (id, alias, name, realm_id, enabled, default_action, provider_id, priority) FROM stdin;
1f0d913c-2188-4556-a13e-376e0320607d	VERIFY_EMAIL	Verify Email	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	VERIFY_EMAIL	50
2ec73962-d268-46b9-ac29-a796ba980237	UPDATE_PROFILE	Update Profile	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	UPDATE_PROFILE	40
b5f73b0c-af2b-4f1c-929b-78bf5139c193	CONFIGURE_TOTP	Configure OTP	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	CONFIGURE_TOTP	10
2e58ec55-0746-4c76-acc3-ddc9c65e4ecc	UPDATE_PASSWORD	Update Password	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	UPDATE_PASSWORD	30
a3fc6814-2d6a-4e7b-8bb8-44a5b9ed0c6d	terms_and_conditions	Terms and Conditions	98749fe9-5c8f-4d46-b973-16664c916f0f	f	f	terms_and_conditions	20
8c3f589d-c849-4acb-a076-ddcd8a5bd28e	delete_account	Delete Account	98749fe9-5c8f-4d46-b973-16664c916f0f	f	f	delete_account	60
801f5d98-8128-4efb-9e8f-d1afdc8a5133	update_user_locale	Update User Locale	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	update_user_locale	1000
5a5993bb-400f-4f71-94c9-75214dacda35	webauthn-register	Webauthn Register	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	webauthn-register	70
0811d3a3-0194-4174-8815-d19c7b741959	webauthn-register-passwordless	Webauthn Register Passwordless	98749fe9-5c8f-4d46-b973-16664c916f0f	t	f	webauthn-register-passwordless	80
0c11a730-ba89-4555-af74-8aa7949e1f10	CONFIGURE_TOTP	Configure OTP	cwbi	t	f	CONFIGURE_TOTP	10
926c88ee-13a2-4d70-9a48-1249cc9c5a10	terms_and_conditions	Terms and Conditions	cwbi	f	f	terms_and_conditions	20
909ccce6-4b44-468a-8d18-49d15ce7fe65	UPDATE_PASSWORD	Update Password	cwbi	t	f	UPDATE_PASSWORD	30
f6b6549a-03d0-41a3-a181-a223d0562864	UPDATE_PROFILE	Update Profile	cwbi	t	f	UPDATE_PROFILE	40
d17c3a54-0639-4a1f-83b0-29ea332905d6	VERIFY_EMAIL	Verify Email	cwbi	t	f	VERIFY_EMAIL	50
61bc6849-c455-41f1-92aa-52479f9b25f9	delete_account	Delete Account	cwbi	f	f	delete_account	60
9984b01f-2707-42d0-9741-cd9a495617ab	update_user_locale	Update User Locale	cwbi	t	f	update_user_locale	1000
\.


--
-- Data for Name: resource_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_attribute (id, name, value, resource_id) FROM stdin;
\.


--
-- Data for Name: resource_policy; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_policy (resource_id, policy_id) FROM stdin;
\.


--
-- Data for Name: resource_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_scope (resource_id, scope_id) FROM stdin;
\.


--
-- Data for Name: resource_server; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_server (id, allow_rs_remote_mgmt, policy_enforce_mode, decision_strategy) FROM stdin;
\.


--
-- Data for Name: resource_server_perm_ticket; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_server_perm_ticket (id, owner, requester, created_timestamp, granted_timestamp, resource_id, scope_id, resource_server_id, policy_id) FROM stdin;
\.


--
-- Data for Name: resource_server_policy; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_server_policy (id, name, description, type, decision_strategy, logic, resource_server_id, owner) FROM stdin;
\.


--
-- Data for Name: resource_server_resource; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_server_resource (id, name, type, icon_uri, owner, resource_server_id, owner_managed_access, display_name) FROM stdin;
\.


--
-- Data for Name: resource_server_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_server_scope (id, name, icon_uri, resource_server_id, display_name) FROM stdin;
\.


--
-- Data for Name: resource_uris; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.resource_uris (resource_id, value) FROM stdin;
\.


--
-- Data for Name: role_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.role_attribute (id, role_id, name, value) FROM stdin;
\.


--
-- Data for Name: scope_mapping; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.scope_mapping (client_id, role_id) FROM stdin;
ff59471c-d8be-4e1b-a846-c01bc708fbe4	0a5420b4-1a6d-42b7-a07a-19aa82f6b26e
ff59471c-d8be-4e1b-a846-c01bc708fbe4	98e010f7-b9db-4a9b-872a-8eee15766444
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	6fc29792-ef9a-4035-8fcf-73a93194cee4
bc4324a3-e1d1-4b25-bb23-1e46ed21709f	3422313a-546d-4e5d-9595-9af9a14532fe
\.


--
-- Data for Name: scope_policy; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.scope_policy (scope_id, policy_id) FROM stdin;
\.


--
-- Data for Name: user_attribute; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_attribute (name, value, user_id, id) FROM stdin;
cacUID	2	f8dcafea-243e-4b89-8d7d-fa01918130f4	4ec925d4-0a48-422b-9922-d764e3f2361f
subjectDN	lambert.anthony.m.2	f8dcafea-243e-4b89-8d7d-fa01918130f4	29137655-cad9-4d5f-b9a4-19427e42e43b
cacUID	54321	f9b33064-13d0-47d7-8294-fb8f0fac819f	4ec94757-c452-4065-9cad-1b052e285bf5
x509_presented	true	f9b33064-13d0-47d7-8294-fb8f0fac819f	d7f8c6eb-0eeb-4843-b2c2-a0ce9147df84
\.


--
-- Data for Name: user_consent; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_consent (id, client_id, user_id, created_date, last_updated_date, client_storage_provider, external_client_id) FROM stdin;
\.


--
-- Data for Name: user_consent_client_scope; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_consent_client_scope (user_consent_id, scope_id) FROM stdin;
\.


--
-- Data for Name: user_entity; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_entity (id, email, email_constraint, email_verified, enabled, federation_link, first_name, last_name, realm_id, username, created_timestamp, service_account_client_link, not_before) FROM stdin;
f3fc1dd9-af7b-498a-9435-31da080a37ad	\N	47fc5432-cc32-4471-ae60-5d638d6b7fb4	f	t	\N	\N	\N	98749fe9-5c8f-4d46-b973-16664c916f0f	admin	1719326096269	\N	0
f8dcafea-243e-4b89-8d7d-fa01918130f4	anthony.m.lambert@fake.usace.army.mil	e2b90b7a-7b44-4982-aa44-c4ca043efb03	t	t	\N	Anthony	Lambert	cwbi	test	1719327608101	\N	0
127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a	molly.rutherford@fake.usace.army.mil	7a01ff11-105d-40c5-9818-f70563fa8256	t	t	\N	Molly	Rutherford	cwbi	nocactest	1719438836112	\N	0
f9b33064-13d0-47d7-8294-fb8f0fac819f	testuser.new@fake.usace.army.mil	4cecef9f-c7b4-47f3-9f7d-2045c6bdb234	t	t	\N	New	TestUser	cwbi	newcacuser	1720198987614	\N	0
\.


--
-- Data for Name: user_federation_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_federation_config (user_federation_provider_id, value, name) FROM stdin;
\.


--
-- Data for Name: user_federation_mapper; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_federation_mapper (id, name, federation_provider_id, federation_mapper_type, realm_id) FROM stdin;
\.


--
-- Data for Name: user_federation_mapper_config; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_federation_mapper_config (user_federation_mapper_id, value, name) FROM stdin;
\.


--
-- Data for Name: user_federation_provider; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_federation_provider (id, changed_sync_period, display_name, full_sync_period, last_sync, priority, provider_name, realm_id) FROM stdin;
\.


--
-- Data for Name: user_group_membership; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_group_membership (group_id, user_id) FROM stdin;
\.


--
-- Data for Name: user_required_action; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_required_action (user_id, required_action) FROM stdin;
\.


--
-- Data for Name: user_role_mapping; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_role_mapping (role_id, user_id) FROM stdin;
db5807fc-5e23-4f38-8243-bfdc9673bd9e	f3fc1dd9-af7b-498a-9435-31da080a37ad
8eb26a12-1853-44b8-bf59-4da9e0cb683c	f3fc1dd9-af7b-498a-9435-31da080a37ad
c883c84f-d797-4170-bcf1-888f3843ba15	f8dcafea-243e-4b89-8d7d-fa01918130f4
c883c84f-d797-4170-bcf1-888f3843ba15	127cbaee-ee0c-4cd9-92a3-8e8a6f023e4a
c883c84f-d797-4170-bcf1-888f3843ba15	f9b33064-13d0-47d7-8294-fb8f0fac819f
\.


--
-- Data for Name: user_session; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_session (id, auth_method, ip_address, last_session_refresh, login_username, realm_id, remember_me, started, user_id, user_session_state, broker_session_id, broker_user_id) FROM stdin;
\.


--
-- Data for Name: user_session_note; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.user_session_note (user_session, name, value) FROM stdin;
\.


--
-- Data for Name: username_login_failure; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.username_login_failure (realm_id, username, failed_login_not_before, last_failure, last_ip_failure, num_failures) FROM stdin;
\.


--
-- Data for Name: web_origins; Type: TABLE DATA; Schema: keycloak; Owner: keycloak_user
--

COPY keycloak.web_origins (client_id, value) FROM stdin;
fca2fb0d-1434-4ba2-bd0a-699e623e79be	+
86b97bc5-1afd-40b2-ad62-bddaaaf321c7	+
38f6b360-c12e-4e69-8b63-43ab4910e344	+
\.


--
-- Name: username_login_failure CONSTRAINT_17-2; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.username_login_failure
    ADD CONSTRAINT "CONSTRAINT_17-2" PRIMARY KEY (realm_id, username);


--
-- Name: keycloak_role UK_J3RWUVD56ONTGSUHOGM184WW2-2; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.keycloak_role
    ADD CONSTRAINT "UK_J3RWUVD56ONTGSUHOGM184WW2-2" UNIQUE (name, client_realm_constraint);


--
-- Name: client_auth_flow_bindings c_cli_flow_bind; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_auth_flow_bindings
    ADD CONSTRAINT c_cli_flow_bind PRIMARY KEY (client_id, binding_name);


--
-- Name: client_scope_client c_cli_scope_bind; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope_client
    ADD CONSTRAINT c_cli_scope_bind PRIMARY KEY (client_id, scope_id);


--
-- Name: client_initial_access cnstr_client_init_acc_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_initial_access
    ADD CONSTRAINT cnstr_client_init_acc_pk PRIMARY KEY (id);


--
-- Name: realm_default_groups con_group_id_def_groups; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_default_groups
    ADD CONSTRAINT con_group_id_def_groups UNIQUE (group_id);


--
-- Name: broker_link constr_broker_link_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.broker_link
    ADD CONSTRAINT constr_broker_link_pk PRIMARY KEY (identity_provider, user_id);


--
-- Name: client_user_session_note constr_cl_usr_ses_note; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_user_session_note
    ADD CONSTRAINT constr_cl_usr_ses_note PRIMARY KEY (client_session, name);


--
-- Name: component_config constr_component_config_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.component_config
    ADD CONSTRAINT constr_component_config_pk PRIMARY KEY (id);


--
-- Name: component constr_component_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.component
    ADD CONSTRAINT constr_component_pk PRIMARY KEY (id);


--
-- Name: fed_user_required_action constr_fed_required_action; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_required_action
    ADD CONSTRAINT constr_fed_required_action PRIMARY KEY (required_action, user_id);


--
-- Name: fed_user_attribute constr_fed_user_attr_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_attribute
    ADD CONSTRAINT constr_fed_user_attr_pk PRIMARY KEY (id);


--
-- Name: fed_user_consent constr_fed_user_consent_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_consent
    ADD CONSTRAINT constr_fed_user_consent_pk PRIMARY KEY (id);


--
-- Name: fed_user_credential constr_fed_user_cred_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_credential
    ADD CONSTRAINT constr_fed_user_cred_pk PRIMARY KEY (id);


--
-- Name: fed_user_group_membership constr_fed_user_group; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_group_membership
    ADD CONSTRAINT constr_fed_user_group PRIMARY KEY (group_id, user_id);


--
-- Name: fed_user_role_mapping constr_fed_user_role; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_role_mapping
    ADD CONSTRAINT constr_fed_user_role PRIMARY KEY (role_id, user_id);


--
-- Name: federated_user constr_federated_user; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.federated_user
    ADD CONSTRAINT constr_federated_user PRIMARY KEY (id);


--
-- Name: realm_default_groups constr_realm_default_groups; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_default_groups
    ADD CONSTRAINT constr_realm_default_groups PRIMARY KEY (realm_id, group_id);


--
-- Name: realm_enabled_event_types constr_realm_enabl_event_types; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_enabled_event_types
    ADD CONSTRAINT constr_realm_enabl_event_types PRIMARY KEY (realm_id, value);


--
-- Name: realm_events_listeners constr_realm_events_listeners; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_events_listeners
    ADD CONSTRAINT constr_realm_events_listeners PRIMARY KEY (realm_id, value);


--
-- Name: realm_supported_locales constr_realm_supported_locales; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_supported_locales
    ADD CONSTRAINT constr_realm_supported_locales PRIMARY KEY (realm_id, value);


--
-- Name: identity_provider constraint_2b; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider
    ADD CONSTRAINT constraint_2b PRIMARY KEY (internal_id);


--
-- Name: client_attributes constraint_3c; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_attributes
    ADD CONSTRAINT constraint_3c PRIMARY KEY (client_id, name);


--
-- Name: event_entity constraint_4; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.event_entity
    ADD CONSTRAINT constraint_4 PRIMARY KEY (id);


--
-- Name: federated_identity constraint_40; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.federated_identity
    ADD CONSTRAINT constraint_40 PRIMARY KEY (identity_provider, user_id);


--
-- Name: realm constraint_4a; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm
    ADD CONSTRAINT constraint_4a PRIMARY KEY (id);


--
-- Name: client_session_role constraint_5; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_role
    ADD CONSTRAINT constraint_5 PRIMARY KEY (client_session, role_id);


--
-- Name: user_session constraint_57; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_session
    ADD CONSTRAINT constraint_57 PRIMARY KEY (id);


--
-- Name: user_federation_provider constraint_5c; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_provider
    ADD CONSTRAINT constraint_5c PRIMARY KEY (id);


--
-- Name: client_session_note constraint_5e; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_note
    ADD CONSTRAINT constraint_5e PRIMARY KEY (client_session, name);


--
-- Name: client constraint_7; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client
    ADD CONSTRAINT constraint_7 PRIMARY KEY (id);


--
-- Name: client_session constraint_8; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session
    ADD CONSTRAINT constraint_8 PRIMARY KEY (id);


--
-- Name: scope_mapping constraint_81; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.scope_mapping
    ADD CONSTRAINT constraint_81 PRIMARY KEY (client_id, role_id);


--
-- Name: client_node_registrations constraint_84; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_node_registrations
    ADD CONSTRAINT constraint_84 PRIMARY KEY (client_id, name);


--
-- Name: realm_attribute constraint_9; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_attribute
    ADD CONSTRAINT constraint_9 PRIMARY KEY (name, realm_id);


--
-- Name: realm_required_credential constraint_92; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_required_credential
    ADD CONSTRAINT constraint_92 PRIMARY KEY (realm_id, type);


--
-- Name: keycloak_role constraint_a; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.keycloak_role
    ADD CONSTRAINT constraint_a PRIMARY KEY (id);


--
-- Name: admin_event_entity constraint_admin_event_entity; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.admin_event_entity
    ADD CONSTRAINT constraint_admin_event_entity PRIMARY KEY (id);


--
-- Name: authenticator_config_entry constraint_auth_cfg_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authenticator_config_entry
    ADD CONSTRAINT constraint_auth_cfg_pk PRIMARY KEY (authenticator_id, name);


--
-- Name: authentication_execution constraint_auth_exec_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authentication_execution
    ADD CONSTRAINT constraint_auth_exec_pk PRIMARY KEY (id);


--
-- Name: authentication_flow constraint_auth_flow_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authentication_flow
    ADD CONSTRAINT constraint_auth_flow_pk PRIMARY KEY (id);


--
-- Name: authenticator_config constraint_auth_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authenticator_config
    ADD CONSTRAINT constraint_auth_pk PRIMARY KEY (id);


--
-- Name: client_session_auth_status constraint_auth_status_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_auth_status
    ADD CONSTRAINT constraint_auth_status_pk PRIMARY KEY (client_session, authenticator);


--
-- Name: user_role_mapping constraint_c; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_role_mapping
    ADD CONSTRAINT constraint_c PRIMARY KEY (role_id, user_id);


--
-- Name: composite_role constraint_composite_role; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.composite_role
    ADD CONSTRAINT constraint_composite_role PRIMARY KEY (composite, child_role);


--
-- Name: client_session_prot_mapper constraint_cs_pmp_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_prot_mapper
    ADD CONSTRAINT constraint_cs_pmp_pk PRIMARY KEY (client_session, protocol_mapper_id);


--
-- Name: identity_provider_config constraint_d; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider_config
    ADD CONSTRAINT constraint_d PRIMARY KEY (identity_provider_id, name);


--
-- Name: policy_config constraint_dpc; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.policy_config
    ADD CONSTRAINT constraint_dpc PRIMARY KEY (policy_id, name);


--
-- Name: realm_smtp_config constraint_e; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_smtp_config
    ADD CONSTRAINT constraint_e PRIMARY KEY (realm_id, name);


--
-- Name: credential constraint_f; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.credential
    ADD CONSTRAINT constraint_f PRIMARY KEY (id);


--
-- Name: user_federation_config constraint_f9; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_config
    ADD CONSTRAINT constraint_f9 PRIMARY KEY (user_federation_provider_id, name);


--
-- Name: resource_server_perm_ticket constraint_fapmt; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT constraint_fapmt PRIMARY KEY (id);


--
-- Name: resource_server_resource constraint_farsr; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_resource
    ADD CONSTRAINT constraint_farsr PRIMARY KEY (id);


--
-- Name: resource_server_policy constraint_farsrp; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_policy
    ADD CONSTRAINT constraint_farsrp PRIMARY KEY (id);


--
-- Name: associated_policy constraint_farsrpap; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.associated_policy
    ADD CONSTRAINT constraint_farsrpap PRIMARY KEY (policy_id, associated_policy_id);


--
-- Name: resource_policy constraint_farsrpp; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_policy
    ADD CONSTRAINT constraint_farsrpp PRIMARY KEY (resource_id, policy_id);


--
-- Name: resource_server_scope constraint_farsrs; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_scope
    ADD CONSTRAINT constraint_farsrs PRIMARY KEY (id);


--
-- Name: resource_scope constraint_farsrsp; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_scope
    ADD CONSTRAINT constraint_farsrsp PRIMARY KEY (resource_id, scope_id);


--
-- Name: scope_policy constraint_farsrsps; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.scope_policy
    ADD CONSTRAINT constraint_farsrsps PRIMARY KEY (scope_id, policy_id);


--
-- Name: user_entity constraint_fb; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_entity
    ADD CONSTRAINT constraint_fb PRIMARY KEY (id);


--
-- Name: user_federation_mapper_config constraint_fedmapper_cfg_pm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_mapper_config
    ADD CONSTRAINT constraint_fedmapper_cfg_pm PRIMARY KEY (user_federation_mapper_id, name);


--
-- Name: user_federation_mapper constraint_fedmapperpm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_mapper
    ADD CONSTRAINT constraint_fedmapperpm PRIMARY KEY (id);


--
-- Name: fed_user_consent_cl_scope constraint_fgrntcsnt_clsc_pm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.fed_user_consent_cl_scope
    ADD CONSTRAINT constraint_fgrntcsnt_clsc_pm PRIMARY KEY (user_consent_id, scope_id);


--
-- Name: user_consent_client_scope constraint_grntcsnt_clsc_pm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_consent_client_scope
    ADD CONSTRAINT constraint_grntcsnt_clsc_pm PRIMARY KEY (user_consent_id, scope_id);


--
-- Name: user_consent constraint_grntcsnt_pm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_consent
    ADD CONSTRAINT constraint_grntcsnt_pm PRIMARY KEY (id);


--
-- Name: keycloak_group constraint_group; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.keycloak_group
    ADD CONSTRAINT constraint_group PRIMARY KEY (id);


--
-- Name: group_attribute constraint_group_attribute_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.group_attribute
    ADD CONSTRAINT constraint_group_attribute_pk PRIMARY KEY (id);


--
-- Name: group_role_mapping constraint_group_role; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.group_role_mapping
    ADD CONSTRAINT constraint_group_role PRIMARY KEY (role_id, group_id);


--
-- Name: identity_provider_mapper constraint_idpm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider_mapper
    ADD CONSTRAINT constraint_idpm PRIMARY KEY (id);


--
-- Name: idp_mapper_config constraint_idpmconfig; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.idp_mapper_config
    ADD CONSTRAINT constraint_idpmconfig PRIMARY KEY (idp_mapper_id, name);


--
-- Name: migration_model constraint_migmod; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.migration_model
    ADD CONSTRAINT constraint_migmod PRIMARY KEY (id);


--
-- Name: offline_client_session constraint_offl_cl_ses_pk3; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.offline_client_session
    ADD CONSTRAINT constraint_offl_cl_ses_pk3 PRIMARY KEY (user_session_id, client_id, client_storage_provider, external_client_id, offline_flag);


--
-- Name: offline_user_session constraint_offl_us_ses_pk2; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.offline_user_session
    ADD CONSTRAINT constraint_offl_us_ses_pk2 PRIMARY KEY (user_session_id, offline_flag);


--
-- Name: protocol_mapper constraint_pcm; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.protocol_mapper
    ADD CONSTRAINT constraint_pcm PRIMARY KEY (id);


--
-- Name: protocol_mapper_config constraint_pmconfig; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.protocol_mapper_config
    ADD CONSTRAINT constraint_pmconfig PRIMARY KEY (protocol_mapper_id, name);


--
-- Name: redirect_uris constraint_redirect_uris; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.redirect_uris
    ADD CONSTRAINT constraint_redirect_uris PRIMARY KEY (client_id, value);


--
-- Name: required_action_config constraint_req_act_cfg_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.required_action_config
    ADD CONSTRAINT constraint_req_act_cfg_pk PRIMARY KEY (required_action_id, name);


--
-- Name: required_action_provider constraint_req_act_prv_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.required_action_provider
    ADD CONSTRAINT constraint_req_act_prv_pk PRIMARY KEY (id);


--
-- Name: user_required_action constraint_required_action; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_required_action
    ADD CONSTRAINT constraint_required_action PRIMARY KEY (required_action, user_id);


--
-- Name: resource_uris constraint_resour_uris_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_uris
    ADD CONSTRAINT constraint_resour_uris_pk PRIMARY KEY (resource_id, value);


--
-- Name: role_attribute constraint_role_attribute_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.role_attribute
    ADD CONSTRAINT constraint_role_attribute_pk PRIMARY KEY (id);


--
-- Name: user_attribute constraint_user_attribute_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_attribute
    ADD CONSTRAINT constraint_user_attribute_pk PRIMARY KEY (id);


--
-- Name: user_group_membership constraint_user_group; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_group_membership
    ADD CONSTRAINT constraint_user_group PRIMARY KEY (group_id, user_id);


--
-- Name: user_session_note constraint_usn_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_session_note
    ADD CONSTRAINT constraint_usn_pk PRIMARY KEY (user_session, name);


--
-- Name: web_origins constraint_web_origins; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.web_origins
    ADD CONSTRAINT constraint_web_origins PRIMARY KEY (client_id, value);


--
-- Name: databasechangeloglock databasechangeloglock_pkey; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.databasechangeloglock
    ADD CONSTRAINT databasechangeloglock_pkey PRIMARY KEY (id);


--
-- Name: client_scope_attributes pk_cl_tmpl_attr; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope_attributes
    ADD CONSTRAINT pk_cl_tmpl_attr PRIMARY KEY (scope_id, name);


--
-- Name: client_scope pk_cli_template; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope
    ADD CONSTRAINT pk_cli_template PRIMARY KEY (id);


--
-- Name: resource_server pk_resource_server; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server
    ADD CONSTRAINT pk_resource_server PRIMARY KEY (id);


--
-- Name: client_scope_role_mapping pk_template_scope; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope_role_mapping
    ADD CONSTRAINT pk_template_scope PRIMARY KEY (scope_id, role_id);


--
-- Name: default_client_scope r_def_cli_scope_bind; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.default_client_scope
    ADD CONSTRAINT r_def_cli_scope_bind PRIMARY KEY (realm_id, scope_id);


--
-- Name: realm_localizations realm_localizations_pkey; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_localizations
    ADD CONSTRAINT realm_localizations_pkey PRIMARY KEY (realm_id, locale);


--
-- Name: resource_attribute res_attr_pk; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_attribute
    ADD CONSTRAINT res_attr_pk PRIMARY KEY (id);


--
-- Name: keycloak_group sibling_names; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.keycloak_group
    ADD CONSTRAINT sibling_names UNIQUE (realm_id, parent_group, name);


--
-- Name: identity_provider uk_2daelwnibji49avxsrtuf6xj33; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider
    ADD CONSTRAINT uk_2daelwnibji49avxsrtuf6xj33 UNIQUE (provider_alias, realm_id);


--
-- Name: client uk_b71cjlbenv945rb6gcon438at; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client
    ADD CONSTRAINT uk_b71cjlbenv945rb6gcon438at UNIQUE (realm_id, client_id);


--
-- Name: client_scope uk_cli_scope; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope
    ADD CONSTRAINT uk_cli_scope UNIQUE (realm_id, name);


--
-- Name: user_entity uk_dykn684sl8up1crfei6eckhd7; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_entity
    ADD CONSTRAINT uk_dykn684sl8up1crfei6eckhd7 UNIQUE (realm_id, email_constraint);


--
-- Name: resource_server_resource uk_frsr6t700s9v50bu18ws5ha6; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_resource
    ADD CONSTRAINT uk_frsr6t700s9v50bu18ws5ha6 UNIQUE (name, owner, resource_server_id);


--
-- Name: resource_server_perm_ticket uk_frsr6t700s9v50bu18ws5pmt; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT uk_frsr6t700s9v50bu18ws5pmt UNIQUE (owner, requester, resource_server_id, resource_id, scope_id);


--
-- Name: resource_server_policy uk_frsrpt700s9v50bu18ws5ha6; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_policy
    ADD CONSTRAINT uk_frsrpt700s9v50bu18ws5ha6 UNIQUE (name, resource_server_id);


--
-- Name: resource_server_scope uk_frsrst700s9v50bu18ws5ha6; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_scope
    ADD CONSTRAINT uk_frsrst700s9v50bu18ws5ha6 UNIQUE (name, resource_server_id);


--
-- Name: user_consent uk_jkuwuvd56ontgsuhogm8uewrt; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_consent
    ADD CONSTRAINT uk_jkuwuvd56ontgsuhogm8uewrt UNIQUE (client_id, client_storage_provider, external_client_id, user_id);


--
-- Name: realm uk_orvsdmla56612eaefiq6wl5oi; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm
    ADD CONSTRAINT uk_orvsdmla56612eaefiq6wl5oi UNIQUE (name);


--
-- Name: user_entity uk_ru8tt6t700s9v50bu18ws5ha6; Type: CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_entity
    ADD CONSTRAINT uk_ru8tt6t700s9v50bu18ws5ha6 UNIQUE (realm_id, username);


--
-- Name: idx_admin_event_time; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_admin_event_time ON keycloak.admin_event_entity USING btree (realm_id, admin_event_time);


--
-- Name: idx_assoc_pol_assoc_pol_id; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_assoc_pol_assoc_pol_id ON keycloak.associated_policy USING btree (associated_policy_id);


--
-- Name: idx_auth_config_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_auth_config_realm ON keycloak.authenticator_config USING btree (realm_id);


--
-- Name: idx_auth_exec_flow; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_auth_exec_flow ON keycloak.authentication_execution USING btree (flow_id);


--
-- Name: idx_auth_exec_realm_flow; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_auth_exec_realm_flow ON keycloak.authentication_execution USING btree (realm_id, flow_id);


--
-- Name: idx_auth_flow_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_auth_flow_realm ON keycloak.authentication_flow USING btree (realm_id);


--
-- Name: idx_cl_clscope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_cl_clscope ON keycloak.client_scope_client USING btree (scope_id);


--
-- Name: idx_client_id; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_client_id ON keycloak.client USING btree (client_id);


--
-- Name: idx_client_init_acc_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_client_init_acc_realm ON keycloak.client_initial_access USING btree (realm_id);


--
-- Name: idx_client_session_session; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_client_session_session ON keycloak.client_session USING btree (session_id);


--
-- Name: idx_clscope_attrs; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_clscope_attrs ON keycloak.client_scope_attributes USING btree (scope_id);


--
-- Name: idx_clscope_cl; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_clscope_cl ON keycloak.client_scope_client USING btree (client_id);


--
-- Name: idx_clscope_protmap; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_clscope_protmap ON keycloak.protocol_mapper USING btree (client_scope_id);


--
-- Name: idx_clscope_role; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_clscope_role ON keycloak.client_scope_role_mapping USING btree (scope_id);


--
-- Name: idx_compo_config_compo; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_compo_config_compo ON keycloak.component_config USING btree (component_id);


--
-- Name: idx_component_provider_type; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_component_provider_type ON keycloak.component USING btree (provider_type);


--
-- Name: idx_component_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_component_realm ON keycloak.component USING btree (realm_id);


--
-- Name: idx_composite; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_composite ON keycloak.composite_role USING btree (composite);


--
-- Name: idx_composite_child; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_composite_child ON keycloak.composite_role USING btree (child_role);


--
-- Name: idx_defcls_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_defcls_realm ON keycloak.default_client_scope USING btree (realm_id);


--
-- Name: idx_defcls_scope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_defcls_scope ON keycloak.default_client_scope USING btree (scope_id);


--
-- Name: idx_event_time; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_event_time ON keycloak.event_entity USING btree (realm_id, event_time);


--
-- Name: idx_fedidentity_feduser; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fedidentity_feduser ON keycloak.federated_identity USING btree (federated_user_id);


--
-- Name: idx_fedidentity_user; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fedidentity_user ON keycloak.federated_identity USING btree (user_id);


--
-- Name: idx_fu_attribute; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_attribute ON keycloak.fed_user_attribute USING btree (user_id, realm_id, name);


--
-- Name: idx_fu_cnsnt_ext; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_cnsnt_ext ON keycloak.fed_user_consent USING btree (user_id, client_storage_provider, external_client_id);


--
-- Name: idx_fu_consent; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_consent ON keycloak.fed_user_consent USING btree (user_id, client_id);


--
-- Name: idx_fu_consent_ru; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_consent_ru ON keycloak.fed_user_consent USING btree (realm_id, user_id);


--
-- Name: idx_fu_credential; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_credential ON keycloak.fed_user_credential USING btree (user_id, type);


--
-- Name: idx_fu_credential_ru; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_credential_ru ON keycloak.fed_user_credential USING btree (realm_id, user_id);


--
-- Name: idx_fu_group_membership; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_group_membership ON keycloak.fed_user_group_membership USING btree (user_id, group_id);


--
-- Name: idx_fu_group_membership_ru; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_group_membership_ru ON keycloak.fed_user_group_membership USING btree (realm_id, user_id);


--
-- Name: idx_fu_required_action; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_required_action ON keycloak.fed_user_required_action USING btree (user_id, required_action);


--
-- Name: idx_fu_required_action_ru; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_required_action_ru ON keycloak.fed_user_required_action USING btree (realm_id, user_id);


--
-- Name: idx_fu_role_mapping; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_role_mapping ON keycloak.fed_user_role_mapping USING btree (user_id, role_id);


--
-- Name: idx_fu_role_mapping_ru; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_fu_role_mapping_ru ON keycloak.fed_user_role_mapping USING btree (realm_id, user_id);


--
-- Name: idx_group_att_by_name_value; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_group_att_by_name_value ON keycloak.group_attribute USING btree (name, ((value)::character varying(250)));


--
-- Name: idx_group_attr_group; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_group_attr_group ON keycloak.group_attribute USING btree (group_id);


--
-- Name: idx_group_role_mapp_group; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_group_role_mapp_group ON keycloak.group_role_mapping USING btree (group_id);


--
-- Name: idx_id_prov_mapp_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_id_prov_mapp_realm ON keycloak.identity_provider_mapper USING btree (realm_id);


--
-- Name: idx_ident_prov_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_ident_prov_realm ON keycloak.identity_provider USING btree (realm_id);


--
-- Name: idx_keycloak_role_client; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_keycloak_role_client ON keycloak.keycloak_role USING btree (client);


--
-- Name: idx_keycloak_role_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_keycloak_role_realm ON keycloak.keycloak_role USING btree (realm);


--
-- Name: idx_offline_css_preload; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_offline_css_preload ON keycloak.offline_client_session USING btree (client_id, offline_flag);


--
-- Name: idx_offline_uss_by_user; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_offline_uss_by_user ON keycloak.offline_user_session USING btree (user_id, realm_id, offline_flag);


--
-- Name: idx_offline_uss_by_usersess; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_offline_uss_by_usersess ON keycloak.offline_user_session USING btree (realm_id, offline_flag, user_session_id);


--
-- Name: idx_offline_uss_createdon; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_offline_uss_createdon ON keycloak.offline_user_session USING btree (created_on);


--
-- Name: idx_offline_uss_preload; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_offline_uss_preload ON keycloak.offline_user_session USING btree (offline_flag, created_on, user_session_id);


--
-- Name: idx_protocol_mapper_client; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_protocol_mapper_client ON keycloak.protocol_mapper USING btree (client_id);


--
-- Name: idx_realm_attr_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_attr_realm ON keycloak.realm_attribute USING btree (realm_id);


--
-- Name: idx_realm_clscope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_clscope ON keycloak.client_scope USING btree (realm_id);


--
-- Name: idx_realm_def_grp_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_def_grp_realm ON keycloak.realm_default_groups USING btree (realm_id);


--
-- Name: idx_realm_evt_list_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_evt_list_realm ON keycloak.realm_events_listeners USING btree (realm_id);


--
-- Name: idx_realm_evt_types_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_evt_types_realm ON keycloak.realm_enabled_event_types USING btree (realm_id);


--
-- Name: idx_realm_master_adm_cli; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_master_adm_cli ON keycloak.realm USING btree (master_admin_client);


--
-- Name: idx_realm_supp_local_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_realm_supp_local_realm ON keycloak.realm_supported_locales USING btree (realm_id);


--
-- Name: idx_redir_uri_client; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_redir_uri_client ON keycloak.redirect_uris USING btree (client_id);


--
-- Name: idx_req_act_prov_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_req_act_prov_realm ON keycloak.required_action_provider USING btree (realm_id);


--
-- Name: idx_res_policy_policy; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_res_policy_policy ON keycloak.resource_policy USING btree (policy_id);


--
-- Name: idx_res_scope_scope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_res_scope_scope ON keycloak.resource_scope USING btree (scope_id);


--
-- Name: idx_res_serv_pol_res_serv; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_res_serv_pol_res_serv ON keycloak.resource_server_policy USING btree (resource_server_id);


--
-- Name: idx_res_srv_res_res_srv; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_res_srv_res_res_srv ON keycloak.resource_server_resource USING btree (resource_server_id);


--
-- Name: idx_res_srv_scope_res_srv; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_res_srv_scope_res_srv ON keycloak.resource_server_scope USING btree (resource_server_id);


--
-- Name: idx_role_attribute; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_role_attribute ON keycloak.role_attribute USING btree (role_id);


--
-- Name: idx_role_clscope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_role_clscope ON keycloak.client_scope_role_mapping USING btree (role_id);


--
-- Name: idx_scope_mapping_role; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_scope_mapping_role ON keycloak.scope_mapping USING btree (role_id);


--
-- Name: idx_scope_policy_policy; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_scope_policy_policy ON keycloak.scope_policy USING btree (policy_id);


--
-- Name: idx_update_time; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_update_time ON keycloak.migration_model USING btree (update_time);


--
-- Name: idx_us_sess_id_on_cl_sess; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_us_sess_id_on_cl_sess ON keycloak.offline_client_session USING btree (user_session_id);


--
-- Name: idx_usconsent_clscope; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_usconsent_clscope ON keycloak.user_consent_client_scope USING btree (user_consent_id);


--
-- Name: idx_user_attribute; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_attribute ON keycloak.user_attribute USING btree (user_id);


--
-- Name: idx_user_attribute_name; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_attribute_name ON keycloak.user_attribute USING btree (name, value);


--
-- Name: idx_user_consent; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_consent ON keycloak.user_consent USING btree (user_id);


--
-- Name: idx_user_credential; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_credential ON keycloak.credential USING btree (user_id);


--
-- Name: idx_user_email; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_email ON keycloak.user_entity USING btree (email);


--
-- Name: idx_user_group_mapping; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_group_mapping ON keycloak.user_group_membership USING btree (user_id);


--
-- Name: idx_user_reqactions; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_reqactions ON keycloak.user_required_action USING btree (user_id);


--
-- Name: idx_user_role_mapping; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_role_mapping ON keycloak.user_role_mapping USING btree (user_id);


--
-- Name: idx_user_service_account; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_user_service_account ON keycloak.user_entity USING btree (realm_id, service_account_client_link);


--
-- Name: idx_usr_fed_map_fed_prv; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_usr_fed_map_fed_prv ON keycloak.user_federation_mapper USING btree (federation_provider_id);


--
-- Name: idx_usr_fed_map_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_usr_fed_map_realm ON keycloak.user_federation_mapper USING btree (realm_id);


--
-- Name: idx_usr_fed_prv_realm; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_usr_fed_prv_realm ON keycloak.user_federation_provider USING btree (realm_id);


--
-- Name: idx_web_orig_client; Type: INDEX; Schema: keycloak; Owner: keycloak_user
--

CREATE INDEX idx_web_orig_client ON keycloak.web_origins USING btree (client_id);


--
-- Name: client_session_auth_status auth_status_constraint; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_auth_status
    ADD CONSTRAINT auth_status_constraint FOREIGN KEY (client_session) REFERENCES keycloak.client_session(id);


--
-- Name: identity_provider fk2b4ebc52ae5c3b34; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider
    ADD CONSTRAINT fk2b4ebc52ae5c3b34 FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: client_attributes fk3c47c64beacca966; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_attributes
    ADD CONSTRAINT fk3c47c64beacca966 FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: federated_identity fk404288b92ef007a6; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.federated_identity
    ADD CONSTRAINT fk404288b92ef007a6 FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: client_node_registrations fk4129723ba992f594; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_node_registrations
    ADD CONSTRAINT fk4129723ba992f594 FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: client_session_note fk5edfb00ff51c2736; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_note
    ADD CONSTRAINT fk5edfb00ff51c2736 FOREIGN KEY (client_session) REFERENCES keycloak.client_session(id);


--
-- Name: user_session_note fk5edfb00ff51d3472; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_session_note
    ADD CONSTRAINT fk5edfb00ff51d3472 FOREIGN KEY (user_session) REFERENCES keycloak.user_session(id);


--
-- Name: client_session_role fk_11b7sgqw18i532811v7o2dv76; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_role
    ADD CONSTRAINT fk_11b7sgqw18i532811v7o2dv76 FOREIGN KEY (client_session) REFERENCES keycloak.client_session(id);


--
-- Name: redirect_uris fk_1burs8pb4ouj97h5wuppahv9f; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.redirect_uris
    ADD CONSTRAINT fk_1burs8pb4ouj97h5wuppahv9f FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: user_federation_provider fk_1fj32f6ptolw2qy60cd8n01e8; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_provider
    ADD CONSTRAINT fk_1fj32f6ptolw2qy60cd8n01e8 FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: client_session_prot_mapper fk_33a8sgqw18i532811v7o2dk89; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session_prot_mapper
    ADD CONSTRAINT fk_33a8sgqw18i532811v7o2dk89 FOREIGN KEY (client_session) REFERENCES keycloak.client_session(id);


--
-- Name: realm_required_credential fk_5hg65lybevavkqfki3kponh9v; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_required_credential
    ADD CONSTRAINT fk_5hg65lybevavkqfki3kponh9v FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: resource_attribute fk_5hrm2vlf9ql5fu022kqepovbr; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_attribute
    ADD CONSTRAINT fk_5hrm2vlf9ql5fu022kqepovbr FOREIGN KEY (resource_id) REFERENCES keycloak.resource_server_resource(id);


--
-- Name: user_attribute fk_5hrm2vlf9ql5fu043kqepovbr; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_attribute
    ADD CONSTRAINT fk_5hrm2vlf9ql5fu043kqepovbr FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: user_required_action fk_6qj3w1jw9cvafhe19bwsiuvmd; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_required_action
    ADD CONSTRAINT fk_6qj3w1jw9cvafhe19bwsiuvmd FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: keycloak_role fk_6vyqfe4cn4wlq8r6kt5vdsj5c; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.keycloak_role
    ADD CONSTRAINT fk_6vyqfe4cn4wlq8r6kt5vdsj5c FOREIGN KEY (realm) REFERENCES keycloak.realm(id);


--
-- Name: realm_smtp_config fk_70ej8xdxgxd0b9hh6180irr0o; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_smtp_config
    ADD CONSTRAINT fk_70ej8xdxgxd0b9hh6180irr0o FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: realm_attribute fk_8shxd6l3e9atqukacxgpffptw; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_attribute
    ADD CONSTRAINT fk_8shxd6l3e9atqukacxgpffptw FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: composite_role fk_a63wvekftu8jo1pnj81e7mce2; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.composite_role
    ADD CONSTRAINT fk_a63wvekftu8jo1pnj81e7mce2 FOREIGN KEY (composite) REFERENCES keycloak.keycloak_role(id);


--
-- Name: authentication_execution fk_auth_exec_flow; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authentication_execution
    ADD CONSTRAINT fk_auth_exec_flow FOREIGN KEY (flow_id) REFERENCES keycloak.authentication_flow(id);


--
-- Name: authentication_execution fk_auth_exec_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authentication_execution
    ADD CONSTRAINT fk_auth_exec_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: authentication_flow fk_auth_flow_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authentication_flow
    ADD CONSTRAINT fk_auth_flow_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: authenticator_config fk_auth_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.authenticator_config
    ADD CONSTRAINT fk_auth_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: client_session fk_b4ao2vcvat6ukau74wbwtfqo1; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_session
    ADD CONSTRAINT fk_b4ao2vcvat6ukau74wbwtfqo1 FOREIGN KEY (session_id) REFERENCES keycloak.user_session(id);


--
-- Name: user_role_mapping fk_c4fqv34p1mbylloxang7b1q3l; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_role_mapping
    ADD CONSTRAINT fk_c4fqv34p1mbylloxang7b1q3l FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: client_scope_attributes fk_cl_scope_attr_scope; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope_attributes
    ADD CONSTRAINT fk_cl_scope_attr_scope FOREIGN KEY (scope_id) REFERENCES keycloak.client_scope(id);


--
-- Name: client_scope_role_mapping fk_cl_scope_rm_scope; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_scope_role_mapping
    ADD CONSTRAINT fk_cl_scope_rm_scope FOREIGN KEY (scope_id) REFERENCES keycloak.client_scope(id);


--
-- Name: client_user_session_note fk_cl_usr_ses_note; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_user_session_note
    ADD CONSTRAINT fk_cl_usr_ses_note FOREIGN KEY (client_session) REFERENCES keycloak.client_session(id);


--
-- Name: protocol_mapper fk_cli_scope_mapper; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.protocol_mapper
    ADD CONSTRAINT fk_cli_scope_mapper FOREIGN KEY (client_scope_id) REFERENCES keycloak.client_scope(id);


--
-- Name: client_initial_access fk_client_init_acc_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.client_initial_access
    ADD CONSTRAINT fk_client_init_acc_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: component_config fk_component_config; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.component_config
    ADD CONSTRAINT fk_component_config FOREIGN KEY (component_id) REFERENCES keycloak.component(id);


--
-- Name: component fk_component_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.component
    ADD CONSTRAINT fk_component_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: realm_default_groups fk_def_groups_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_default_groups
    ADD CONSTRAINT fk_def_groups_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: user_federation_mapper_config fk_fedmapper_cfg; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_mapper_config
    ADD CONSTRAINT fk_fedmapper_cfg FOREIGN KEY (user_federation_mapper_id) REFERENCES keycloak.user_federation_mapper(id);


--
-- Name: user_federation_mapper fk_fedmapperpm_fedprv; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_mapper
    ADD CONSTRAINT fk_fedmapperpm_fedprv FOREIGN KEY (federation_provider_id) REFERENCES keycloak.user_federation_provider(id);


--
-- Name: user_federation_mapper fk_fedmapperpm_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_mapper
    ADD CONSTRAINT fk_fedmapperpm_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: associated_policy fk_frsr5s213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.associated_policy
    ADD CONSTRAINT fk_frsr5s213xcx4wnkog82ssrfy FOREIGN KEY (associated_policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: scope_policy fk_frsrasp13xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.scope_policy
    ADD CONSTRAINT fk_frsrasp13xcx4wnkog82ssrfy FOREIGN KEY (policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: resource_server_perm_ticket fk_frsrho213xcx4wnkog82sspmt; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT fk_frsrho213xcx4wnkog82sspmt FOREIGN KEY (resource_server_id) REFERENCES keycloak.resource_server(id);


--
-- Name: resource_server_resource fk_frsrho213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_resource
    ADD CONSTRAINT fk_frsrho213xcx4wnkog82ssrfy FOREIGN KEY (resource_server_id) REFERENCES keycloak.resource_server(id);


--
-- Name: resource_server_perm_ticket fk_frsrho213xcx4wnkog83sspmt; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT fk_frsrho213xcx4wnkog83sspmt FOREIGN KEY (resource_id) REFERENCES keycloak.resource_server_resource(id);


--
-- Name: resource_server_perm_ticket fk_frsrho213xcx4wnkog84sspmt; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT fk_frsrho213xcx4wnkog84sspmt FOREIGN KEY (scope_id) REFERENCES keycloak.resource_server_scope(id);


--
-- Name: associated_policy fk_frsrpas14xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.associated_policy
    ADD CONSTRAINT fk_frsrpas14xcx4wnkog82ssrfy FOREIGN KEY (policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: scope_policy fk_frsrpass3xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.scope_policy
    ADD CONSTRAINT fk_frsrpass3xcx4wnkog82ssrfy FOREIGN KEY (scope_id) REFERENCES keycloak.resource_server_scope(id);


--
-- Name: resource_server_perm_ticket fk_frsrpo2128cx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_perm_ticket
    ADD CONSTRAINT fk_frsrpo2128cx4wnkog82ssrfy FOREIGN KEY (policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: resource_server_policy fk_frsrpo213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_policy
    ADD CONSTRAINT fk_frsrpo213xcx4wnkog82ssrfy FOREIGN KEY (resource_server_id) REFERENCES keycloak.resource_server(id);


--
-- Name: resource_scope fk_frsrpos13xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_scope
    ADD CONSTRAINT fk_frsrpos13xcx4wnkog82ssrfy FOREIGN KEY (resource_id) REFERENCES keycloak.resource_server_resource(id);


--
-- Name: resource_policy fk_frsrpos53xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_policy
    ADD CONSTRAINT fk_frsrpos53xcx4wnkog82ssrfy FOREIGN KEY (resource_id) REFERENCES keycloak.resource_server_resource(id);


--
-- Name: resource_policy fk_frsrpp213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_policy
    ADD CONSTRAINT fk_frsrpp213xcx4wnkog82ssrfy FOREIGN KEY (policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: resource_scope fk_frsrps213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_scope
    ADD CONSTRAINT fk_frsrps213xcx4wnkog82ssrfy FOREIGN KEY (scope_id) REFERENCES keycloak.resource_server_scope(id);


--
-- Name: resource_server_scope fk_frsrso213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_server_scope
    ADD CONSTRAINT fk_frsrso213xcx4wnkog82ssrfy FOREIGN KEY (resource_server_id) REFERENCES keycloak.resource_server(id);


--
-- Name: composite_role fk_gr7thllb9lu8q4vqa4524jjy8; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.composite_role
    ADD CONSTRAINT fk_gr7thllb9lu8q4vqa4524jjy8 FOREIGN KEY (child_role) REFERENCES keycloak.keycloak_role(id);


--
-- Name: user_consent_client_scope fk_grntcsnt_clsc_usc; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_consent_client_scope
    ADD CONSTRAINT fk_grntcsnt_clsc_usc FOREIGN KEY (user_consent_id) REFERENCES keycloak.user_consent(id);


--
-- Name: user_consent fk_grntcsnt_user; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_consent
    ADD CONSTRAINT fk_grntcsnt_user FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: group_attribute fk_group_attribute_group; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.group_attribute
    ADD CONSTRAINT fk_group_attribute_group FOREIGN KEY (group_id) REFERENCES keycloak.keycloak_group(id);


--
-- Name: group_role_mapping fk_group_role_group; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.group_role_mapping
    ADD CONSTRAINT fk_group_role_group FOREIGN KEY (group_id) REFERENCES keycloak.keycloak_group(id);


--
-- Name: realm_enabled_event_types fk_h846o4h0w8epx5nwedrf5y69j; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_enabled_event_types
    ADD CONSTRAINT fk_h846o4h0w8epx5nwedrf5y69j FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: realm_events_listeners fk_h846o4h0w8epx5nxev9f5y69j; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_events_listeners
    ADD CONSTRAINT fk_h846o4h0w8epx5nxev9f5y69j FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: identity_provider_mapper fk_idpm_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider_mapper
    ADD CONSTRAINT fk_idpm_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: idp_mapper_config fk_idpmconfig; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.idp_mapper_config
    ADD CONSTRAINT fk_idpmconfig FOREIGN KEY (idp_mapper_id) REFERENCES keycloak.identity_provider_mapper(id);


--
-- Name: web_origins fk_lojpho213xcx4wnkog82ssrfy; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.web_origins
    ADD CONSTRAINT fk_lojpho213xcx4wnkog82ssrfy FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: scope_mapping fk_ouse064plmlr732lxjcn1q5f1; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.scope_mapping
    ADD CONSTRAINT fk_ouse064plmlr732lxjcn1q5f1 FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: protocol_mapper fk_pcm_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.protocol_mapper
    ADD CONSTRAINT fk_pcm_realm FOREIGN KEY (client_id) REFERENCES keycloak.client(id);


--
-- Name: credential fk_pfyr0glasqyl0dei3kl69r6v0; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.credential
    ADD CONSTRAINT fk_pfyr0glasqyl0dei3kl69r6v0 FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: protocol_mapper_config fk_pmconfig; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.protocol_mapper_config
    ADD CONSTRAINT fk_pmconfig FOREIGN KEY (protocol_mapper_id) REFERENCES keycloak.protocol_mapper(id);


--
-- Name: default_client_scope fk_r_def_cli_scope_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.default_client_scope
    ADD CONSTRAINT fk_r_def_cli_scope_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: required_action_provider fk_req_act_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.required_action_provider
    ADD CONSTRAINT fk_req_act_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: resource_uris fk_resource_server_uris; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.resource_uris
    ADD CONSTRAINT fk_resource_server_uris FOREIGN KEY (resource_id) REFERENCES keycloak.resource_server_resource(id);


--
-- Name: role_attribute fk_role_attribute_id; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.role_attribute
    ADD CONSTRAINT fk_role_attribute_id FOREIGN KEY (role_id) REFERENCES keycloak.keycloak_role(id);


--
-- Name: realm_supported_locales fk_supported_locales_realm; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.realm_supported_locales
    ADD CONSTRAINT fk_supported_locales_realm FOREIGN KEY (realm_id) REFERENCES keycloak.realm(id);


--
-- Name: user_federation_config fk_t13hpu1j94r2ebpekr39x5eu5; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_federation_config
    ADD CONSTRAINT fk_t13hpu1j94r2ebpekr39x5eu5 FOREIGN KEY (user_federation_provider_id) REFERENCES keycloak.user_federation_provider(id);


--
-- Name: user_group_membership fk_user_group_user; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.user_group_membership
    ADD CONSTRAINT fk_user_group_user FOREIGN KEY (user_id) REFERENCES keycloak.user_entity(id);


--
-- Name: policy_config fkdc34197cf864c4e43; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.policy_config
    ADD CONSTRAINT fkdc34197cf864c4e43 FOREIGN KEY (policy_id) REFERENCES keycloak.resource_server_policy(id);


--
-- Name: identity_provider_config fkdc4897cf864c4e43; Type: FK CONSTRAINT; Schema: keycloak; Owner: keycloak_user
--

ALTER TABLE ONLY keycloak.identity_provider_config
    ADD CONSTRAINT fkdc4897cf864c4e43 FOREIGN KEY (identity_provider_id) REFERENCES keycloak.identity_provider(internal_id);


--
-- PostgreSQL database dump complete
--

