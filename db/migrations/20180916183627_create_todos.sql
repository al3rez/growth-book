-- migrate:up
create table if not exists todos (
  id serial not null unique,
  title varchar(255),
  user_id integer,
	created_at timestamp not null default now(),
	updated_at timestamp
);


-- migrate:down
drop table if exists todos;
