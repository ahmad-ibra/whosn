BEGIN;

SET client_encoding = 'UTF8';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    user_name text NOT NULL UNIQUE,
    password text NOT NULL,
    email text NOT NULL UNIQUE,
    phone_number text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE events (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL REFERENCES users (id),
    name text NOT NULL,
    time timestamp NOT NULL,
    location text NOT NULL,
    min_users smallint NOT NULL CHECK ( min_users > 0 ),
    max_users smallint NOT NULL CHECK ( max_users >= min_users ),
    price numeric(8,2) NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE event_users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL REFERENCES events (id),
    user_id uuid NOT NULL REFERENCES users (id),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);

COMMIT;
