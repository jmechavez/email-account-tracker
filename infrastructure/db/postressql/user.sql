CREATE TABLE users (
    id_no VARCHAR(255) PRIMARY KEY,
    department VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    suffix VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    email_status VARCHAR(255),
    status VARCHAR(255),
    ticket_no VARCHAR(255),
    profile_picture VARCHAR(255),
    hashed_password VARCHAR(255),
    salt VARCHAR(255),
    smtp_email VARCHAR(255),
    smtp_password VARCHAR(255),
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP WITH TIME ZONE CURRENT_TIMESTAMP,
    date_deleted TIMESTAMP WITH TIME ZONE CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);

-- Enums for email_status (active or deleted only)
CREATE TYPE email_status AS ENUM ('active', 'deleted');
ALTER TABLE users ALTER COLUMN email_status TYPE email_status USING email_status::email_status;
