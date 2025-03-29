CREATE TABLE users (
    id_no VARCHAR(255) PRIMARY KEY,
    department VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    suffix VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    email_status VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    ticket_no VARCHAR(255),
    updated_ticket_no VARCHAR(255),
    deleted_ticket_no VARCHAR(255),
    profile_picture VARCHAR(255),
    hashed_password VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    smtp_email VARCHAR(255),
    smtp_password VARCHAR(255),
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    date_deleted TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);


-- Enums for email_status (active or deleted only)
CREATE TYPE email_status AS ENUM ('active', 'deleted');
ALTER TABLE users ALTER COLUMN email_status TYPE email_status USING email_status::email_status;
