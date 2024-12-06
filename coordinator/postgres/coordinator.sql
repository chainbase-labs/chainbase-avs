CREATE DATABASE coordinator;

\c coordinator;

create type operator_status as enum ('active', 'inactive');

create type task_status as enum ('completed', 'failed');

CREATE TABLE IF NOT EXISTS operator
(
    id               serial primary key,
    operator_address varchar(42) not null unique,
    operator_id      varchar(64) not null unique,
    socket           varchar(20) not null,
    location         varchar(30),
    cpu_core         integer,
    memory           integer,
    status           operator_status          default 'active'::operator_status,
    registered_at    timestamp with time zone,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task
(
    id               serial primary key,
    task_id          integer      not null unique,
    task_detail      varchar(128) not null,
    task_response    varchar(66),
    create_task_tx   varchar(66),
    response_task_tx varchar(66),
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS operator_task
(
    id          serial primary key,
    operator_id integer not null references operator,
    task_id     integer not null,
    status      task_status              default 'completed'::task_status,
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone default CURRENT_TIMESTAMP
);