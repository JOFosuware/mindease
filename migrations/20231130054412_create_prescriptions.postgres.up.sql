CREATE TABLE prescriptions (
    id SERIAL PRIMARY KEY,
    form_id VARCHAR,
    name VARCHAR,
    institution VARCHAR,
    physician VARCHAR,
    image BYTEA,
    location VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)