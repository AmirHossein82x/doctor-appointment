-- Create a temporary ENUM type with both old and new values
DO $$ BEGIN
    CREATE TYPE user_role_temp AS ENUM ('normal', 'patient', 'admin', 'doctor');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Drop the default value before altering the type
ALTER TABLE users ALTER COLUMN role DROP DEFAULT;

-- Alter the column to use the temporary ENUM type
ALTER TABLE users 
    ALTER COLUMN role TYPE user_role_temp USING (role::text::user_role_temp);

-- Update 'patient' back to 'normal'
UPDATE users SET role = 'normal' WHERE role = 'patient';

-- Create the final ENUM type with original values
DO $$ BEGIN
    CREATE TYPE user_role_new AS ENUM ('normal', 'admin', 'doctor');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Alter the column to use the final ENUM type and restore the original default
ALTER TABLE users 
    ALTER COLUMN role TYPE user_role_new USING (role::text::user_role_new),
    ALTER COLUMN role SET DEFAULT 'normal';

-- Clean up: drop the old and temporary ENUM types
DROP TYPE user_role;
DROP TYPE user_role_temp;

-- Rename the new ENUM type to the original name
ALTER TYPE user_role_new RENAME TO user_role;