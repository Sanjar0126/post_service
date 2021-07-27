do $$
    begin
        alter table customers add column shipper_id uuid;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;

create unique index if not exists customers_phone_key on customers(shipper_id, phone);