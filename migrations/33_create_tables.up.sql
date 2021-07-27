CREATE TABLE IF NOT EXISTS customer_types (
    id uuid primary key,
    name varchar(250) not null,
    phone_number varchar(20) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

ALTER TABLE customers
    ADD COLUMN customer_type_id uuid;