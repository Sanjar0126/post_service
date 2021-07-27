do $$
    begin
        ALTER TABLE shippers DROP COLUMN process_only_paid_orders;    
        ALTER TABLE shippers DROP COLUMN show_location_before_accepting;     
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;