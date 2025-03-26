-- Create ENUM type for roles
DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('normal', 'admin', 'doctor');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    phone VARCHAR(20) UNIQUE,
    role user_role NOT NULL DEFAULT 'normal', -- Role column with ENUM
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);