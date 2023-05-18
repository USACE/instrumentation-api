DROP VIEW  IF EXISTS v_alert;
DROP TABLE IF EXISTS alert_config CASCADE;

CREATE TABLE evaluation (
    id              UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    project_id      UUID NOT NULL REFERENCES project (id) ON DELETE CASCADE,
    name 			VARCHAR(480) NOT NULL,
    body 			TEXT NOT NULL DEFAULT '',
    timerange       TSTZRANGE NOT NULL,
    creator 		UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date 	TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater 		UUID,
    update_date 	TIMESTAMPTZ
);

CREATE TABLE evaluation_instrument (
    evaluation_id   UUID REFERENCES evaluation (id),
    instrument_id   UUID REFERENCES instrument (id)
);

CREATE TABLE alert_status (
    id      SMALLINT PRIMARY KEY NOT NULL,
    name    TEXT UNIQUE NOT NULL
);

INSERT INTO alert_status (id, name) VALUES
    (1, 'green'),
    (2, 'yellow'),
    (3, 'red');

CREATE TABLE alert_type (
    id      SMALLINT PRIMARY KEY NOT NULL,
    name    TEXT UNIQUE NOT NULL
);

INSERT INTO alert_type (id, name) VALUES
    (1, 'Missing Time Series Measurements'),
    (2, 'Overdue Evaluation');

CREATE TABLE alert_config (
    id 				        UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    project_id              UUID NOT NULL REFERENCES project (id),
    name 			        VARCHAR(480) NOT NULL,
    body 			        TEXT NOT NULL DEFAULT '',
    creator 		        UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    create_date 	        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updater 		        UUID,
    update_date 	        TIMESTAMPTZ,
    alert_type_id           SMALLINT NOT NULL REFERENCES alert_type (id),
    start_date              TIMESTAMPTZ NOT NULL DEFAULT now(),
    schedule_interval 	    INTERVAL NOT NULL,
    n_missed_before_alert   INT NOT NULL DEFAULT 1,
    warning_interval        INTERVAL,
    remind_interval	        INTERVAL NOT NULL default '1 day',
    last_checked 	        TIMESTAMPTZ,
    last_reminded	        TIMESTAMPTZ,
    alert_status_id         SMALLINT NOT NULL REFERENCES alert_status (id) DEFAULT 1
);

CREATE TABLE alert_config_instrument (
    alert_config_id UUID NOT NULL REFERENCES alert_config (id),
    instrument_id   UUID NOT NULL REFERENCES instrument (id)
);
