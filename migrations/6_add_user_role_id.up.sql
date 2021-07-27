do $$
    begin
        alter table shippers add column user_role_id uuid;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;