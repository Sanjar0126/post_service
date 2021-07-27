do $$
    begin
        ALTER TABLE system_users RENAME TO shipper_users;
        
        ALTER TABLE shipper_users alter column deleted_at type integer using date_part('epoch', deleted_at)::int;
        ALTER TABLE shipper_users alter column deleted_at set default 0;
        ALTER TABLE shipper_users add unique (username, deleted_at);
        ALTER TABLE shipper_users DROP CONSTRAINT system_users_username_key;
        ALTER TABLE shipper_users DROP COLUMN is_active;
        UPDATE shipper_users SET deleted_at=0 WHERE deleted_at is NULL;

        ALTER TABLE shippers DROP COLUMN access_token, DROP COLUMN username, DROP COLUMN password, DROP COLUMN user_role_id;
        ALTER TABLE shippers ADD COLUMN call_center_tg VARCHAR(255), ADD COLUMN work_hour_start TIME, ADD COLUMN work_hour_end TIME;
        ALTER TABLE shippers ALTER column deleted_at TYPE INTEGER USING date_part('epoch', deleted_at)::int;
        ALTER TABLE shippers ALTER column deleted_at SET default 0;
        UPDATE shippers SET deleted_at=0 WHERE deleted_at IS NULL;
        ALTER TABLE shippers DROP CONSTRAINT shippers_name_key, DROP CONSTRAINT shippers_phone_key;
        ALTER TABLE shippers ADD unique (name, deleted_at);
        ALTER TABLE shippers ALTER COLUMN work_hour_start set default '10:00', ALTER COLUMN work_hour_end set default '00:00';
        UPDATE shippers SET work_hour_start='10:00' WHERE work_hour_start is NULL;
		UPDATE shippers SET work_hour_end='00:00' WHERE work_hour_end is NULL;
                
        ALTER TABLE branches DROP COLUMN access_token, DROP COLUMN username, DROP COLUMN password;
        ALTER TABLE branches ALTER column deleted_at TYPE INTEGER USING date_part('epoch', deleted_at)::int;
        ALTER TABLE branches ALTER column deleted_at SET default 0;
        UPDATE branches SET deleted_at=0 WHERE deleted_at IS NULL;
        ALTER TABLE branches add unique (shipper_id, name, deleted_at); 
        ALTER TABLE branches ALTER column phone type text;
        ALTER TABLE branches ADD COLUMN work_hours varchar(255), ADD COLUMN image text;

        ALTER TABLE customers DROP COLUMN access_token;  
        ALTER TABLE customers RENAME COLUMN is_active to is_blocked;
        UPDATE customers SET is_blocked = false WHERE is_blocked = true;
        ALTER TABLE customers ADD COLUMN fcm_token TEXT;  
        ALTER TABLE customers ALTER COLUMN is_blocked SET default false;
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;


CREATE TABLE IF NOT EXISTS system_users (
    id uuid PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(20) NOT NULL,
    password VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_role_id uuid
);

ALTER TABLE system_users ADD unique(username, deleted_at);

CREATE TABLE IF NOT EXISTS branch_users (
    id uuid PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    shipper_id uuid NOT NULL,
    branch_id uuid NOT NULL, 
    user_role_id uuid NOT NULL,
    fcm_token TEXT
);

ALTER TABLE branch_users ADD unique(phone, shipper_id, deleted_at);