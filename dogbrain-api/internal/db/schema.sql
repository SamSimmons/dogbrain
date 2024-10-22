CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    verification_token VARCHAR(255),
    verified_at TIMESTAMP,
    token_expiry TIMESTAMP WITH TIME ZONE
);