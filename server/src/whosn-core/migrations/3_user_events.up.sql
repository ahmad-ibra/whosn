CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS event_users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL REFERENCES events (id),
    user_id uuid NOT NULL REFERENCES users (id),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);