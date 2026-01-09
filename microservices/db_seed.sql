--- Create main tables
CREATE TABLE
  IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

-- Seed test data for users table (all passwords are crypted version of the word 'password')
INSERT INTO
  users (
    first_name,
    last_name,
    email,
    password,
    active,
    created_at,
    updated_at
  )
VALUES
  (
    'John',
    'Doe',
    'test@example.com',
    '$2a$12$g5AO4c1Qk.vxI.z28OEX5OPugOqQ/nTxEiGhkNmH1xjlJlTQ9EpAK',
    true,
    NOW (),
    NOW ()
  );
