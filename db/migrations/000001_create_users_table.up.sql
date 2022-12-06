CREATE TABLE IF NOT EXISTS users (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR NOT NULL,
    password VARCHAR,
    email VARCHAR,
    confirmed boolean,
    created_at timestamp without time zone NOT NULL
);
