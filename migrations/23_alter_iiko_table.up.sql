do $$
    begin
        alter table iiko_credentials add column dispatcher_id uuid;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;