CREATE TABLE IF NOT EXISTS jowi_credentials (
    shipper_id uuid PRIMARY KEY,
    dispatcher_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

DO $$
    BEGIN
        ALTER TABLE branches ADD COLUMN jowi_id uuid;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Already existed';
    END $$;