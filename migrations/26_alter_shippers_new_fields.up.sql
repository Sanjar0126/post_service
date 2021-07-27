do $$
    begin
        alter table shippers add column process_only_paid_orders boolean default false;
        alter table shippers add column show_location_before_accepting boolean default false;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;