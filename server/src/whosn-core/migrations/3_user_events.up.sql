CREATE TABLE IF NOT EXISTS event_users (
    event_id uuid NOT NULL REFERENCES events (id),
    user_id uuid NOT NULL REFERENCES users (id),
    created_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY (event_id, user_id)
);
