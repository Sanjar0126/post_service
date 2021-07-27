do $$
    begin
        ALTER TABLE shippers DROP COLUMN max_courier_orders;     
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;