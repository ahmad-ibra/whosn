CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
     id uuid NOT NULL DEFAULT uuid_generate_v4(),
     name text NOT NULL,
     user_name text NOT NULL UNIQUE,
     password text NOT NULL,
     email text NOT NULL UNIQUE,
     phone_number text NOT NULL,
     created_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
     updated_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
     PRIMARY KEY (id)
);
