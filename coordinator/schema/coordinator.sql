CREATE DATABASE coordinator;

\c coordinator;

CREATE TABLE IF NOT EXISTS operator
(
    id               serial primary key,
    operator_address varchar(42) not null unique,
    operator_id      varchar(64) not null,
    socket           varchar(20) not null,
    location         varchar(30),
    cpu_core         integer,
    memory           integer,
    status           operator_status default 'active'::operator_status,
    registered_at    timestamp with time zone,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP
);

alter table operator
    owner to postgres;


create type operator_status as enum ('active', 'inactive');

alter type operator_status owner to postgres;