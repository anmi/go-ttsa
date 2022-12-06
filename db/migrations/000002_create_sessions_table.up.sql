CREATE TABLE IF NOT EXISTS sessions (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id bigint NOT NULL,
    token VARCHAR NOT NULL UNIQUE,
    created_at timestamp without time zone NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
