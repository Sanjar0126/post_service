CREATE TABLE IF NOT EXISTS aggregators (
    id uuid primary key,
    shipper_id uuid,
    name varchar(250),
    phone_number varchar(20)
);