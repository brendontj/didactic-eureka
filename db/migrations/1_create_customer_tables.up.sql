CREATE SCHEMA IF NOT EXISTS customer;

CREATE TABLE IF NOT EXISTS customer.customer (
    id UUID PRIMARY KEY,
    version UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    document VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customer.address (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    street VARCHAR(255) NOT NULL,
    number VARCHAR(255) NOT NULL,
    complement VARCHAR(255),
    neighborhood VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip_code VARCHAR(255) NOT NULL,
    UNIQUE (customer_id),
    FOREIGN KEY (customer_id) REFERENCES customer.customer(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_customer_customer_id ON customer.customer (id);
CREATE INDEX IF NOT EXISTS idx_customer_address_customer_id ON customer.address (customer_id);
