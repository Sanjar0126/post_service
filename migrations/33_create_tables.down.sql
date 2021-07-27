DROP TABLE IF EXISTS customer_types;

ALTER TABLE customers 
    DROP COLUMN customer_type_id;