CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    phone VARCHAR,
    condition text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)