CREATE TABLE IF NOT EXISTS customers (
     id uuid PRIMARY KEY,
     access_token VARCHAR NOT NULL UNIQUE,
     name  VARCHAR(50) NOT NULL,
     phone VARCHAR(20) NOT NULL,
     is_active BOOLEAN NOT NULL DEFAULT TRUE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS branches (
     id uuid PRIMARY KEY,
     shipper_id uuid NOT NULL,  
     access_token text NOT NULL UNIQUE,
     name VARCHAR(50) NOT NULL UNIQUE,
     username VARCHAR(20) NOT NULL UNIQUE,
     password  VARCHAR NOT NULL,
     phone  TEXT [] NOT NULL UNIQUE,
     is_active BOOLEAN NOT NULL DEFAULT TRUE,
     address TEXT not null,
     destination text,
     location geometry NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shippers (
     id uuid PRIMARY KEY,
     access_token text NOT NULL UNIQUE,
     name VARCHAR(50) NOT NULL UNIQUE,
     username VARCHAR(20) NOT NULL UNIQUE,
     logo VARCHAR,
     description text,
     password VARCHAR NOT NULL,
     phone TEXT [] NOT NULL UNIQUE,
     is_active BOOLEAN NOT NULL DEFAULT TRUE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);
