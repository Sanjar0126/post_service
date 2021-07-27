do $$
    begin
        alter table system_users drop constraint system_users_phone_key;
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;