CREATE TABLE IF NOT EXISTS iiko_credentials (
    shipper_id uuid PRIMARY KEY,
    api_login VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

DO $$
    BEGIN
        ALTER TABLE branches ADD COLUMN iiko_id uuid;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Already existed';
    END $$;