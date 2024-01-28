CREATE TABLE providers (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    username VARCHAR,
    description text,
    password VARCHAR,
    image BYTEA,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)