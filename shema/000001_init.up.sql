CREATE TABLE
  users (
    id serial not nul unique,
    name varchar(255) not nul,
    username varchar(255) not nul unique,
    password_hash varchar(255) not nul
  );

CREATE TABLE
  users_lists (
    id serial not null,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
  );

CREATE TABLE
  todo_lists (
    id serial not null unique,
    title varchar(255) not nul,
    description varchar(255)
  );

CREATE TABLE
  lists_item(
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not nul,
    list_id int references todo_lists (id) on delete cascade not nul
  );

CREATE TABLE
  todo_items(
    id serial not null unique,
    title varchar(255) not nul,
    description varchar(255),
    done boolean not nul default false
  );