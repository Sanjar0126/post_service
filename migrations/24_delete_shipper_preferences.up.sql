DROP TABLE IF EXISTS shipper_preferences;

DROP TRIGGER IF EXISTS create_shipper_preferences
ON shippers;

do $$
    begin
        ALTER TABLE shippers
            ADD COLUMN courier_accepts_first BOOLEAN DEFAULT false,
            ADD COLUMN check_courier_action_radius BOOLEAN DEFAULT false,
            ADD COLUMN courier_action_radius INTEGER DEFAULT -1,
            ADD COLUMN max_deliver_time INTEGER DEFAULT 60;     
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;