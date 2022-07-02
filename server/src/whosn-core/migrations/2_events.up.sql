CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS events (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL REFERENCES users (id),
    name text NOT NULL,
    time timestamp WITH TIME ZONE NOT NULL,
    location text NOT NULL,
    min_users smallint NOT NULL CHECK ( min_users > 0 ),
    max_users smallint NOT NULL CHECK ( max_users >= min_users ),
    price numeric(8,2) NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY (id)
);
