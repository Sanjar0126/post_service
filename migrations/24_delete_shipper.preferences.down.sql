do $$
    begin
        ALTER TABLE shippers
            DROP COLUMN courier_accepts_first,
            DROP COLUMN check_courier_action_radius,
            DROP COLUMN courier_action_radius,
            DROP COLUMN max_delivery_time;     
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;