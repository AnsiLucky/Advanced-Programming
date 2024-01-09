CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR,
    email      VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
