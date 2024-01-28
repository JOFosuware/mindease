CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR,
    phone INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)