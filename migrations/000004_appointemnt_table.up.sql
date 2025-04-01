-- Create the doctor_appointment table with a default status of 'available'
CREATE TABLE doctor_appointment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doctor_profile_id UUID NOT NULL REFERENCES doctor_profile(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'booked')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);