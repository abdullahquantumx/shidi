-- Create the database if it doesn't exist
CREATE DATABASE accountuser;

-- Connect to the newly created database
\c accountuser;

-- Create a user with the appropriate password (replace 'yourpassword' with an actual password)
CREATE USER accountuser WITH ENCRYPTED PASSWORD 'yourpassword';

-- Grant privileges to the user on the newly created database
GRANT ALL PRIVILEGES ON DATABASE accountuser TO accountuser;

-- Create the necessary schema or tables for your application
CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert any initial data (optional)
-- Example: Inserting a default user (replace with your actual data)
INSERT INTO accounts (username, email, password) 
VALUES ('admin', 'admin@example.com', 'adminpassword');
