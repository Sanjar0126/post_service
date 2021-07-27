do $$
    begin
        alter table customers alter column access_token drop not null ;
    exception
        when others then
            RAISE NOTICE 'Already deleted';
    end $$;