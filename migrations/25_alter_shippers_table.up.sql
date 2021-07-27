do $$
    begin
        alter table shippers add column max_courier_orders integer default 3;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;