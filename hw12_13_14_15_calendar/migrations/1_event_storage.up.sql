CREATE SEQUENCE IF NOT EXISTS event_storage_seq_id
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;

create table if not exists event_storage (
    id numeric(19,0) DEFAULT nextval('event_storage_seq_id'::regclass) NOT NULL,
    title varchar(255),
    "start" timestamp,
    "end" timestamp,
    description varchar(1024),
    owner_id int,
    days_before_notify int
);

