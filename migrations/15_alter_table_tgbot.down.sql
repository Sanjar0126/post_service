DO $$
    BEGIN
        ALTER TABLE tgbots DROP COLUMN url;
        ALTER TABLE tgbots DROP COLUMN access_token;
        ALTER TABLE shippers DROP COLUMN menu_image;
        ALTER TABLE branches DROP column work_hour_start;
        ALTER TABLE branches DROP column work_hour_end;
        ALTER TABLE branches ADD COLUMN work_hours VARCHAR;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Already dropped';
    END $$;