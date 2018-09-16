-- migrate:up
create table if not exists users (
	id serial not null unique,
	email varchar(255) not null unique,
	created_at timestamp not null default now(),
	updated_at timestamp
);


-- migrate:down
drop table if exists users;
