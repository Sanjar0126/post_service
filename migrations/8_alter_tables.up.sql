drop index customers_phone_key;

do $$
    begin
        alter table customers alter column deleted_at type integer using date_part('epoch', deleted_at)::int;
        alter table customers alter column deleted_at set default 0;
        alter table customers add unique (phone, shipper_id, deleted_at);
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;

update customers set deleted_at=0 where deleted_at is null;