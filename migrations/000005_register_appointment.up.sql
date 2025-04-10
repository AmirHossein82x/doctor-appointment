-- Create ENUM type for user_appointment status
DO $$ BEGIN
    CREATE TYPE user_appointment_status AS ENUM ('reserved', 'cancelled');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Create the user_appointment table
CREATE TABLE user_appointment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_appointment_id UUID NOT NULL REFERENCES doctor_appointment(id) ON DELETE CASCADE,
    status user_appointment_status NOT NULL DEFAULT 'reserved',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
