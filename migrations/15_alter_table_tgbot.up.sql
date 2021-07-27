DO $$
    BEGIN
        ALTER TABLE tgbots ADD COLUMN url VARCHAR DEFAULT 'url_default';
        ALTER TABLE tgbots ADD COLUMN access_token TEXT DEFAULT 'url_access_token';
        ALTER TABLE shippers ADD COLUMN menu_image VARCHAR;
        ALTER TABLE branches ADD COLUMN work_hour_start TIME DEFAULT '10:00', ADD COLUMN work_hour_end TIME DEFAULT '00:00';
        ALTER TABLE branches DROP COLUMN work_hours;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Already existed';
    END $$;