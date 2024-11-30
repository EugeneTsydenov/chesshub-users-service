CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    display_name VARCHAR(10) NOT NULL,
    hash_password TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    games INT DEFAULT 0,
    rating INT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);
