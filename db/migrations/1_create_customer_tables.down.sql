DROP INDEX IF EXISTS idx_customer_address_customer_id;
DROP INDEX IF EXISTS idx_customer_customer_id;

DROP TABLE IF EXISTS customer.address;
DROP TABLE IF EXISTS customer.customer;

DROP SCHEMA IF EXISTS customer CASCADE;