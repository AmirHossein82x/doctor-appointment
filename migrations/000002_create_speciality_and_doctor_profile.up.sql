-- Create the speciality table
CREATE TABLE speciality (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

-- Create the doctor_profile table
CREATE TABLE doctor_profile (
    id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    speciality_id INT NOT NULL REFERENCES speciality(id) ON DELETE RESTRICT,
    bio TEXT,
    experience_years INT CHECK (experience_years >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);